package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	codexBaseURL      = "https://chatgpt.com/backend-api/codex"
	codexResponsesURL = codexBaseURL + "/responses"

	refreshURL  = "https://auth.openai.com/oauth/token"
	clientID    = "app_EMoamEEZ73f0CkXaXp7hrann"
	refreshSkew = 30 * time.Second
)

var authMu sync.Mutex

type tokenRefreshResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
	IDToken      string `json:"id_token,omitempty"`
}

type sseEvent struct {
	EventType string
	Data      string
}

type codexStreamEvent struct {
	Type     string          `json:"type"`
	Delta    string          `json:"delta,omitempty"`
	Response json.RawMessage `json:"response,omitempty"`
	Item     json.RawMessage `json:"item,omitempty"`
}

type responseProbe struct {
	ID        string            `json:"id"`
	Object    string            `json:"object"`
	CreatedAt int64             `json:"created_at"`
	Status    string            `json:"status"`
	Model     string            `json:"model"`
	Output    []json.RawMessage `json:"output"`
	Usage     json.RawMessage   `json:"usage"`
	Error     json.RawMessage   `json:"error"`
}

type syntheticResponsesResponse struct {
	ID        string            `json:"id"`
	Object    string            `json:"object"`
	CreatedAt int64             `json:"created_at"`
	Status    string            `json:"status"`
	Model     string            `json:"model"`
	Output    []json.RawMessage `json:"output"`
	Usage     json.RawMessage   `json:"usage"`
	Error     any               `json:"error"`
}

type syntheticMessageItem struct {
	ID      string                    `json:"id"`
	Type    string                    `json:"type"`
	Role    string                    `json:"role"`
	Status  string                    `json:"status"`
	Content []syntheticMessageContent `json:"content"`
}

type syntheticMessageContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func main() {
	addr := flag.String("addr", "127.0.0.1:8787", "listen address")
	defaultInstructions := flag.String("instructions", "You are a helpful coding assistant.", "default instructions to add when missing")
	debug := flag.Bool("debug", false, "log request/stream debugging details")
	flag.Parse()

	mux := http.NewServeMux()

	// Shelley should be configured with endpoint http://127.0.0.1:8787/v1,
	// which causes it to call /v1/responses.
	mux.HandleFunc("/v1/responses", func(w http.ResponseWriter, r *http.Request) {
		handleResponses(w, r, *defaultInstructions, *debug)
	})

	// Also support /responses for manual testing.
	mux.HandleFunc("/responses", func(w http.ResponseWriter, r *http.Request) {
		handleResponses(w, r, *defaultInstructions, *debug)
	})

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok\n"))
	})

	log.Printf("Codex Responses proxy listening on http://%s", *addr)
	log.Printf("Configure Shelley endpoint as http://%s/v1", *addr)

	if err := http.ListenAndServe(*addr, loggingMiddleware(mux)); err != nil {
		log.Fatal(err)
	}
}

func handleResponses(w http.ResponseWriter, r *http.Request, defaultInstructions string, debug bool) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	patchedBody, fallbackModel := patchPayloadForCodex(rawBody, defaultInstructions)

	if debug {
		log.Printf("patched request body: %s", truncate(string(patchedBody), 4000))
	}

	token, accountID, err := borrowCodexKey(r.Context())
	if err != nil {
		http.Error(w, "failed to borrow Codex auth: "+err.Error(), http.StatusUnauthorized)
		return
	}

	upstreamReq, err := http.NewRequestWithContext(
		r.Context(),
		http.MethodPost,
		codexResponsesURL,
		bytes.NewReader(patchedBody),
	)
	if err != nil {
		http.Error(w, "failed to create upstream request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	upstreamReq.Header.Set("Content-Type", "application/json")
	upstreamReq.Header.Set("Accept", "text/event-stream")
	upstreamReq.Header.Set("Authorization", "Bearer "+token)
	if accountID != "" {
		upstreamReq.Header.Set("ChatGPT-Account-ID", accountID)
	}

	client := &http.Client{
		Timeout: 10 * time.Minute,
	}

	upstreamResp, err := client.Do(upstreamReq)
	if err != nil {
		http.Error(w, "upstream request failed: "+err.Error(), http.StatusBadGateway)
		return
	}
	defer upstreamResp.Body.Close()

	contentType := upstreamResp.Header.Get("Content-Type")

	if upstreamResp.StatusCode < 200 || upstreamResp.StatusCode >= 300 {
		body, _ := io.ReadAll(upstreamResp.Body)
		log.Printf("upstream error HTTP %d: %s", upstreamResp.StatusCode, truncate(string(body), 4000))
		w.Header().Set("Content-Type", firstNonEmpty(contentType, "application/json"))
		w.WriteHeader(upstreamResp.StatusCode)
		_, _ = w.Write(body)
		return
	}

	// If upstream ever gives us completed JSON, pass it straight through.
	// Normally Codex requires stream=true and returns SSE.
	if strings.Contains(strings.ToLower(contentType), "application/json") {
		body, _ := io.ReadAll(upstreamResp.Body)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(body)
		return
	}

	responseJSON, err := convertSSEToResponsesJSON(upstreamResp.Body, fallbackModel, debug)
	if err != nil {
		http.Error(w, "failed to convert Codex stream: "+err.Error(), http.StatusBadGateway)
		return
	}

	if debug {
		log.Printf("synthetic/final response JSON: %s", truncate(string(responseJSON), 4000))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(responseJSON)
}

func patchPayloadForCodex(raw []byte, defaultInstructions string) ([]byte, string) {
	var payload map[string]any
	if err := json.Unmarshal(raw, &payload); err != nil {
		return raw, ""
	}

	model, _ := payload["model"].(string)

	// Codex backend requires streaming.
	payload["stream"] = true

	// Sensible default used by Simon Willison's plugin too.
	if _, ok := payload["store"]; !ok {
		payload["store"] = false
	}

	// Codex backend requires top-level instructions.
	if _, ok := payload["instructions"]; !ok {
		payload["instructions"] = defaultInstructions
	}

	// Codex backend rejects this. opencode also clears maxOutputTokens
	// in its Codex plugin to match Codex CLI behavior.
	delete(payload, "max_output_tokens")
	delete(payload, "max_completion_tokens")

	next, err := json.Marshal(payload)
	if err != nil {
		return raw, model
	}
	return next, model
}

func convertSSEToResponsesJSON(body io.Reader, fallbackModel string, debug bool) ([]byte, error) {
	var textBuilder strings.Builder
	var completedRaw json.RawMessage
	var completed responseProbe
	var upstreamError json.RawMessage

	var nonMessageItems []json.RawMessage
	var messageItem json.RawMessage

	err := iterSSEEvents(body, func(ev sseEvent) error {
		data := strings.TrimSpace(ev.Data)
		if data == "" || data == "[DONE]" {
			return nil
		}

		var streamEvent codexStreamEvent
		if err := json.Unmarshal([]byte(data), &streamEvent); err != nil {
			if debug {
				log.Printf("failed to parse SSE event %q: %v; data=%s", ev.EventType, err, truncate(data, 1000))
			}
			return nil
		}

		if debug && streamEvent.Type != "response.output_text.delta" {
			log.Printf("sse event type=%s", streamEvent.Type)
		}

		switch streamEvent.Type {
		case "response.output_text.delta":
			textBuilder.WriteString(streamEvent.Delta)

		case "response.output_item.done":
			if len(streamEvent.Item) == 0 {
				return nil
			}

			var probe struct {
				Type string `json:"type"`
			}
			if err := json.Unmarshal(streamEvent.Item, &probe); err != nil {
				if debug {
					log.Printf("failed to probe output item: %v; item=%s", err, truncate(string(streamEvent.Item), 1000))
				}
				return nil
			}

			switch probe.Type {
			case "message":
				messageItem = append(json.RawMessage(nil), streamEvent.Item...)
			case "function_call", "reasoning":
				nonMessageItems = append(nonMessageItems, append(json.RawMessage(nil), streamEvent.Item...))
			default:
				// Preserve unknown non-message output items rather than dropping them.
				nonMessageItems = append(nonMessageItems, append(json.RawMessage(nil), streamEvent.Item...))
			}

		case "response.completed":
			if len(streamEvent.Response) > 0 {
				completedRaw = append(json.RawMessage(nil), streamEvent.Response...)
				if err := json.Unmarshal(streamEvent.Response, &completed); err != nil && debug {
					log.Printf("failed to parse response.completed payload: %v", err)
				}
			}

		case "response.failed", "response.incomplete":
			if len(streamEvent.Response) > 0 {
				completedRaw = append(json.RawMessage(nil), streamEvent.Response...)
				if err := json.Unmarshal(streamEvent.Response, &completed); err != nil && debug {
					log.Printf("failed to parse %s payload: %v", streamEvent.Type, err)
				}
				if len(completed.Error) > 0 && string(completed.Error) != "null" {
					upstreamError = append(json.RawMessage(nil), completed.Error...)
				}
			}

		case "error":
			upstreamError = append(json.RawMessage(nil), []byte(data)...)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	// Best case: Codex's response.completed already includes a full Responses API
	// object with output items. Return it unchanged. This is the dumbest/most robust
	// response-side behavior.
	if responseHasOutput(completedRaw) {
		return completedRaw, nil
	}

	now := time.Now().Unix()

	if completed.ID == "" {
		completed.ID = fmt.Sprintf("resp_proxy_%d", now)
	}
	if completed.Object == "" {
		completed.Object = "response"
	}
	if completed.CreatedAt == 0 {
		completed.CreatedAt = now
	}
	if completed.Status == "" {
		completed.Status = "completed"
	}
	if completed.Model == "" {
		completed.Model = fallbackModel
	}
	if len(completed.Usage) == 0 || string(completed.Usage) == "null" {
		completed.Usage = json.RawMessage(`{"input_tokens":0,"output_tokens":0,"total_tokens":0}`)
	}

	var output []json.RawMessage

	if textBuilder.Len() > 0 {
		output = append(output, mustMarshalRaw(syntheticMessageItem{
			ID:     "msg_proxy_0",
			Type:   "message",
			Role:   "assistant",
			Status: "completed",
			Content: []syntheticMessageContent{
				{
					Type: "output_text",
					Text: textBuilder.String(),
				},
			},
		}))
	} else if len(messageItem) > 0 {
		output = append(output, messageItem)
	}

	output = append(output, nonMessageItems...)

	if len(upstreamError) > 0 {
		return nil, fmt.Errorf("Codex stream ended with error: %s", truncate(string(upstreamError), 1000))
	}
	if completed.Status == "failed" || completed.Status == "incomplete" || completed.Status == "cancelled" {
		return nil, fmt.Errorf("Codex stream ended with status %q and no output", completed.Status)
	}
	if textBuilder.Len() == 0 && len(messageItem) == 0 && len(nonMessageItems) == 0 {
		return nil, errors.New("Codex stream completed without output")
	}

	result := syntheticResponsesResponse{
		ID:        completed.ID,
		Object:    completed.Object,
		CreatedAt: completed.CreatedAt,
		Status:    completed.Status,
		Model:     completed.Model,
		Output:    output,
		Usage:     completed.Usage,
		Error:     nil,
	}

	return json.Marshal(result)
}

func responseHasOutput(raw json.RawMessage) bool {
	if len(raw) == 0 {
		return false
	}
	var probe struct {
		Output []json.RawMessage `json:"output"`
	}
	if err := json.Unmarshal(raw, &probe); err != nil {
		return false
	}
	return len(probe.Output) > 0
}

func mustMarshalRaw(v any) json.RawMessage {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return json.RawMessage(b)
}

// iterSSEEvents parses Server-Sent Events according to the basic SSE rules:
// events end on a blank line; multiple data: lines are joined with "\n".
func iterSSEEvents(r io.Reader, yield func(sseEvent) error) error {
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 0, 64*1024), 10*1024*1024)

	var (
		eventType string
		dataLines []string
		hasData   bool
	)

	dispatch := func() error {
		if !hasData {
			eventType = ""
			return nil
		}

		ev := sseEvent{
			EventType: eventType,
			Data:      strings.Join(dataLines, "\n"),
		}

		eventType = ""
		dataLines = dataLines[:0]
		hasData = false

		return yield(ev)
	}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if err := dispatch(); err != nil {
				return err
			}
			continue
		}

		if strings.HasPrefix(line, ":") {
			continue
		}

		field, value, ok := strings.Cut(line, ":")
		if !ok {
			field = line
			value = ""
		} else if strings.HasPrefix(value, " ") {
			value = value[1:]
		}

		switch field {
		case "event":
			eventType = value
		case "data":
			dataLines = append(dataLines, value)
			hasData = true
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("reading SSE stream: %w", err)
	}

	return dispatch()
}

func borrowCodexKey(ctx context.Context) (accessToken string, accountID string, err error) {
	authMu.Lock()
	defer authMu.Unlock()

	authPath, err := codexAuthPath()
	if err != nil {
		return "", "", err
	}

	auth, err := readAuthDoc(authPath)
	if err != nil {
		return "", "", err
	}

	authMode := getString(auth, "auth_mode")
	if authMode != "chatgpt" {
		return "", "", fmt.Errorf("expected auth_mode 'chatgpt', got %q; run `codex login --device-auth`", authMode)
	}

	tokens, ok := getMap(auth, "tokens")
	if !ok {
		return "", "", errors.New("no tokens object found in auth.json; run `codex login --device-auth`")
	}

	accessToken = getString(tokens, "access_token")
	refreshToken := getString(tokens, "refresh_token")
	idToken := getString(tokens, "id_token")
	accountID = getString(tokens, "account_id")

	if accountID == "" {
		accountID = extractAccountID(idToken, accessToken)
		if accountID != "" {
			tokens["account_id"] = accountID
		}
	}

	if accessToken == "" {
		return "", "", errors.New("no access_token found; run `codex login --device-auth`")
	}

	expiry, ok := jwtExpiry(accessToken)
	if ok && time.Now().Before(expiry.Add(-refreshSkew)) {
		// Write back discovered account_id, if any.
		if accountID != "" {
			auth["tokens"] = tokens
			_ = writeAuthDoc(authPath, auth)
		}
		return accessToken, accountID, nil
	}

	if refreshToken == "" {
		return "", "", errors.New("access token expired and no refresh_token found; run `codex login --device-auth`")
	}

	refreshed, err := refreshCodexToken(ctx, refreshToken)
	if err != nil {
		return "", "", err
	}

	if refreshed.AccessToken != "" {
		tokens["access_token"] = refreshed.AccessToken
		accessToken = refreshed.AccessToken
	}
	if refreshed.RefreshToken != "" {
		tokens["refresh_token"] = refreshed.RefreshToken
	}
	if refreshed.IDToken != "" {
		tokens["id_token"] = refreshed.IDToken
		idToken = refreshed.IDToken
	}

	if accountID == "" {
		accountID = extractAccountID(idToken, accessToken)
		if accountID != "" {
			tokens["account_id"] = accountID
		}
	}

	auth["tokens"] = tokens
	auth["last_refresh"] = time.Now().UTC().Format(time.RFC3339)

	if err := writeAuthDoc(authPath, auth); err != nil {
		return "", "", err
	}

	return accessToken, accountID, nil
}

func codexAuthPath() (string, error) {
	codexHome := os.Getenv("CODEX_HOME")
	if codexHome == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		codexHome = filepath.Join(home, ".codex")
	}

	path := filepath.Join(codexHome, "auth.json")
	if _, err := os.Stat(path); err != nil {
		return "", fmt.Errorf("Codex auth file not found at %s; run `codex login --device-auth`", path)
	}
	return path, nil
}

func readAuthDoc(path string) (map[string]any, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var doc map[string]any
	if err := json.Unmarshal(b, &doc); err != nil {
		return nil, err
	}
	return doc, nil
}

func writeAuthDoc(path string, doc map[string]any) error {
	b, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return err
	}

	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, b, 0o600); err != nil {
		return err
	}
	if err := os.Rename(tmp, path); err != nil {
		return err
	}
	return os.Chmod(path, 0o600)
}

func refreshCodexToken(ctx context.Context, refreshToken string) (*tokenRefreshResponse, error) {
	body := map[string]string{
		"client_id":     clientID,
		"grant_type":    "refresh_token",
		"refresh_token": refreshToken,
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, refreshURL, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("token refresh network error: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("token refresh failed: HTTP %d: %s", resp.StatusCode, string(respBody))
	}

	var refreshed tokenRefreshResponse
	if err := json.Unmarshal(respBody, &refreshed); err != nil {
		return nil, fmt.Errorf("failed to parse token refresh response: %w", err)
	}

	if refreshed.AccessToken == "" {
		return nil, errors.New("token refresh response did not include access_token")
	}

	return &refreshed, nil
}

func jwtExpiry(token string) (time.Time, bool) {
	claims, ok := parseJWTClaims(token)
	if !ok {
		return time.Time{}, false
	}

	exp, ok := claims["exp"].(float64)
	if !ok || exp == 0 {
		return time.Time{}, false
	}

	return time.Unix(int64(exp), 0), true
}

func parseJWTClaims(token string) (map[string]any, bool) {
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return nil, false
	}

	payload := parts[1]

	decoded, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		if rem := len(payload) % 4; rem != 0 {
			payload += strings.Repeat("=", 4-rem)
		}
		decoded, err = base64.URLEncoding.DecodeString(payload)
		if err != nil {
			return nil, false
		}
	}

	var claims map[string]any
	if err := json.Unmarshal(decoded, &claims); err != nil {
		return nil, false
	}

	return claims, true
}

func extractAccountID(idToken string, accessToken string) string {
	for _, token := range []string{idToken, accessToken} {
		if token == "" {
			continue
		}

		claims, ok := parseJWTClaims(token)
		if !ok {
			continue
		}

		if v, ok := claims["chatgpt_account_id"].(string); ok && v != "" {
			return v
		}

		if nested, ok := claims["https://api.openai.com/auth"].(map[string]any); ok {
			if v, ok := nested["chatgpt_account_id"].(string); ok && v != "" {
				return v
			}
		}

		if orgs, ok := claims["organizations"].([]any); ok && len(orgs) > 0 {
			if first, ok := orgs[0].(map[string]any); ok {
				if v, ok := first["id"].(string); ok && v != "" {
					return v
				}
			}
		}
	}

	return ""
}

func getMap(m map[string]any, key string) (map[string]any, bool) {
	v, ok := m[key]
	if !ok {
		return nil, false
	}

	asMap, ok := v.(map[string]any)
	return asMap, ok
}

func getString(m map[string]any, key string) string {
	v, ok := m[key]
	if !ok {
		return ""
	}

	s, ok := v.(string)
	if !ok {
		return ""
	}

	return s
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("%s %s completed in %s", r.Method, r.URL.Path, time.Since(start))
	})
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max] + fmt.Sprintf("... (%d bytes total)", len(s))
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}
