package main

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"
)

func TestPatchPayloadForCodexAddsRequiredDefaults(t *testing.T) {
	raw := []byte(`{
		"model": "gpt-5.5",
		"input": "hello",
		"max_output_tokens": 100,
		"max_completion_tokens": 100
	}`)

	patched, fallbackModel := patchPayloadForCodex(raw, "default instructions")

	if fallbackModel != "gpt-5.5" {
		t.Fatalf("fallback model = %q, want gpt-5.5", fallbackModel)
	}

	var payload map[string]any
	if err := json.Unmarshal(patched, &payload); err != nil {
		t.Fatalf("patched payload is invalid JSON: %v", err)
	}

	if payload["stream"] != true {
		t.Fatalf("stream = %#v, want true", payload["stream"])
	}
	if payload["store"] != false {
		t.Fatalf("store = %#v, want false", payload["store"])
	}
	if payload["instructions"] != "default instructions" {
		t.Fatalf("instructions = %#v, want default instructions", payload["instructions"])
	}
	if _, ok := payload["max_output_tokens"]; ok {
		t.Fatal("max_output_tokens was not removed")
	}
	if _, ok := payload["max_completion_tokens"]; ok {
		t.Fatal("max_completion_tokens was not removed")
	}
}

func TestPatchPayloadForCodexPreservesExplicitInstructionsAndStore(t *testing.T) {
	raw := []byte(`{"model":"gpt-5.5","instructions":"be terse","store":true}`)

	patched, _ := patchPayloadForCodex(raw, "default instructions")

	var payload map[string]any
	if err := json.Unmarshal(patched, &payload); err != nil {
		t.Fatalf("patched payload is invalid JSON: %v", err)
	}

	if payload["instructions"] != "be terse" {
		t.Fatalf("instructions = %#v, want be terse", payload["instructions"])
	}
	if payload["store"] != true {
		t.Fatalf("store = %#v, want true", payload["store"])
	}
}

func TestIterSSEEvents(t *testing.T) {
	input := strings.NewReader(strings.Join([]string{
		": ignored comment",
		"event: custom.event",
		"data: first line",
		"data: second line",
		"",
		"data: no explicit event",
		"",
	}, "\n"))

	var events []sseEvent
	err := iterSSEEvents(input, func(ev sseEvent) error {
		events = append(events, ev)
		return nil
	})
	if err != nil {
		t.Fatalf("iterSSEEvents returned error: %v", err)
	}

	if len(events) != 2 {
		t.Fatalf("got %d events, want 2: %#v", len(events), events)
	}
	if events[0].EventType != "custom.event" {
		t.Fatalf("first event type = %q", events[0].EventType)
	}
	if events[0].Data != "first line\nsecond line" {
		t.Fatalf("first event data = %q", events[0].Data)
	}
	if events[1].EventType != "" {
		t.Fatalf("second event type = %q, want empty", events[1].EventType)
	}
	if events[1].Data != "no explicit event" {
		t.Fatalf("second event data = %q", events[1].Data)
	}
}

func TestConvertSSEToResponsesJSONSynthesizesMessageOutput(t *testing.T) {
	sse := strings.Join([]string{
		`event: response.output_text.delta`,
		`data: {"type":"response.output_text.delta","delta":"Hello "}`,
		``,
		`event: response.output_text.delta`,
		`data: {"type":"response.output_text.delta","delta":"world"}`,
		``,
		`event: response.completed`,
		`data: {"type":"response.completed","response":{"id":"resp_1","object":"response","created_at":123,"status":"completed","model":"gpt-5.5","usage":{"input_tokens":1,"output_tokens":2,"total_tokens":3}}}`,
		``,
	}, "\n")

	out, err := convertSSEToResponsesJSON(strings.NewReader(sse), "fallback-model", false)
	if err != nil {
		t.Fatalf("convertSSEToResponsesJSON returned error: %v", err)
	}

	var response struct {
		ID     string `json:"id"`
		Model  string `json:"model"`
		Output []struct {
			Type    string `json:"type"`
			Role    string `json:"role"`
			Content []struct {
				Type string `json:"type"`
				Text string `json:"text"`
			} `json:"content"`
		} `json:"output"`
		Usage map[string]int `json:"usage"`
	}
	if err := json.Unmarshal(out, &response); err != nil {
		t.Fatalf("response is invalid JSON: %v\n%s", err, out)
	}

	if response.ID != "resp_1" {
		t.Fatalf("id = %q", response.ID)
	}
	if response.Model != "gpt-5.5" {
		t.Fatalf("model = %q", response.Model)
	}
	if len(response.Output) != 1 {
		t.Fatalf("output len = %d, want 1: %s", len(response.Output), out)
	}
	if response.Output[0].Type != "message" || response.Output[0].Role != "assistant" {
		t.Fatalf("unexpected output item: %#v", response.Output[0])
	}
	if got := response.Output[0].Content[0].Text; got != "Hello world" {
		t.Fatalf("text = %q, want Hello world", got)
	}
	if response.Usage["total_tokens"] != 3 {
		t.Fatalf("usage = %#v", response.Usage)
	}
}

func TestConvertSSEToResponsesJSONReturnsCompletedResponseWithOutputUnchanged(t *testing.T) {
	completed := `{"id":"resp_passthrough","object":"response","created_at":456,"status":"completed","model":"gpt-5.5","output":[{"type":"message","role":"assistant","content":[{"type":"output_text","text":"already complete"}]}],"usage":{"input_tokens":1,"output_tokens":1,"total_tokens":2}}`
	sse := "event: response.completed\n" +
		"data: {\"type\":\"response.completed\",\"response\":" + completed + "}\n\n"

	out, err := convertSSEToResponsesJSON(strings.NewReader(sse), "fallback-model", false)
	if err != nil {
		t.Fatalf("convertSSEToResponsesJSON returned error: %v", err)
	}

	if string(out) != completed {
		t.Fatalf("completed response was not passed through unchanged:\n got: %s\nwant: %s", out, completed)
	}
}

func TestConvertSSEToResponsesJSONErrorsOnEmptyCompletedStream(t *testing.T) {
	sse := strings.Join([]string{
		`event: response.completed`,
		`data: {"type":"response.completed","response":{"id":"resp_empty","object":"response","created_at":123,"status":"completed","model":"gpt-5.5","usage":{"input_tokens":0,"output_tokens":0,"total_tokens":0}}}`,
		``,
	}, "\n")

	_, err := convertSSEToResponsesJSON(strings.NewReader(sse), "fallback-model", false)
	if err == nil {
		t.Fatal("convertSSEToResponsesJSON returned nil error for empty completed stream")
	}
	if !strings.Contains(err.Error(), "without output") {
		t.Fatalf("error = %q, want without output", err)
	}
}

func TestConvertSSEToResponsesJSONErrorsOnFailedStream(t *testing.T) {
	sse := strings.Join([]string{
		`event: response.failed`,
		`data: {"type":"response.failed","response":{"id":"resp_failed","object":"response","created_at":123,"status":"failed","model":"gpt-5.5","error":{"message":"backend died"},"usage":{"input_tokens":1,"output_tokens":0,"total_tokens":1}}}`,
		``,
	}, "\n")

	_, err := convertSSEToResponsesJSON(strings.NewReader(sse), "fallback-model", false)
	if err == nil {
		t.Fatal("convertSSEToResponsesJSON returned nil error for failed stream")
	}
	if !strings.Contains(err.Error(), "backend died") {
		t.Fatalf("error = %q, want upstream error details", err)
	}
}

func TestExtractAccountIDFromJWTClaims(t *testing.T) {
	token := unsignedJWT(t, map[string]any{
		"https://api.openai.com/auth": map[string]any{
			"chatgpt_account_id": "acct_123",
		},
	})

	if got := extractAccountID("", token); got != "acct_123" {
		t.Fatalf("account id = %q, want acct_123", got)
	}
}

func unsignedJWT(t *testing.T, claims map[string]any) string {
	t.Helper()

	header, err := json.Marshal(map[string]any{"alg": "none"})
	if err != nil {
		t.Fatal(err)
	}
	payload, err := json.Marshal(claims)
	if err != nil {
		t.Fatal(err)
	}

	return base64.RawURLEncoding.EncodeToString(header) + "." +
		base64.RawURLEncoding.EncodeToString(payload) + "."
}
