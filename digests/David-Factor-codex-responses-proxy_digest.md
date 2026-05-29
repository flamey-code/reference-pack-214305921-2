This file is a merged representation of the entire codebase, combined into a single document by Repomix.

<file_summary>
This section contains a summary of this file.

<purpose>
This file contains a packed representation of the entire repository's contents.
It is designed to be easily consumable by AI systems for analysis, code review,
or other automated processes.
</purpose>

<file_format>
The content is organized as follows:
1. This summary section
2. Repository information
3. Directory structure
4. Repository files (if enabled)
5. Multiple file entries, each consisting of:
  - File path as an attribute
  - Full contents of the file
</file_format>

<usage_guidelines>
- This file should be treated as read-only. Any changes should be made to the
  original repository files, not this packed version.
- When processing this file, use the file path to distinguish
  between different files in the repository.
- Be aware that this file may contain sensitive information. Handle it with
  the same level of security as you would the original repository.
</usage_guidelines>

<notes>
- Some files may have been excluded based on .gitignore rules and Repomix's configuration
- Binary files are not included in this packed representation. Please refer to the Repository Structure section for a complete list of file paths, including binary files
- Files matching patterns in .gitignore are excluded
- Files matching default ignore patterns are excluded
- Files are sorted by Git change count (files with more changes are at the bottom)
</notes>

</file_summary>

<directory_structure>
.github/workflows/test.yml
.gitignore
contrib/systemd/user/codex-responses-proxy.service
go.mod
LICENSE
main_test.go
main.go
README.md
SECURITY.md
</directory_structure>

<files>
This section contains the contents of the repository's files.

<file path=".github/workflows/test.yml">
name: Tests

on:
  push:
  pull_request:

permissions:
  contents: read

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version-file: go.mod
      - run: go test ./...
      - run: go vet ./...
</file>

<file path="contrib/systemd/user/codex-responses-proxy.service">
[Unit]
Description=Codex Responses Proxy
Documentation=https://github.com/David-Factor/codex-responses-proxy
After=network-online.target
Wants=network-online.target

[Service]
Type=simple
ExecStart=%h/.local/bin/codex-responses-proxy -addr 127.0.0.1:8787
Restart=on-failure
RestartSec=5s

# Keep the proxy private to the local machine by default.
# If you change -addr, add your own authentication/access controls.

[Install]
WantedBy=default.target
</file>

<file path="go.mod">
module github.com/David-Factor/codex-responses-proxy

go 1.22
</file>

<file path="LICENSE">
Apache License
Version 2.0, January 2004
http://www.apache.org/licenses/

TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION

1. Definitions.

"License" shall mean the terms and conditions for use, reproduction,
and distribution as defined by Sections 1 through 9 of this document.

"Licensor" shall mean the copyright owner or entity authorized by
the copyright owner that is granting the License.

"Legal Entity" shall mean the union of the acting entity and all
other entities that control, are controlled by, or are under common
control with that entity. For the purposes of this definition,
"control" means (i) the power, direct or indirect, to cause the
direction or management of such entity, whether by contract or
otherwise, or (ii) ownership of fifty percent (50%) or more of the
outstanding shares, or (iii) beneficial ownership of such entity.

"You" (or "Your") shall mean an individual or Legal Entity
exercising permissions granted by this License.

"Source" form shall mean the preferred form for making modifications,
including but not limited to software source code, documentation
source, and configuration files.

"Object" form shall mean any form resulting from mechanical
transformation or translation of a Source form, including but
not limited to compiled object code, generated documentation,
and conversions to other media types.

"Work" shall mean the work of authorship, whether in Source or
Object form, made available under the License, as indicated by a
copyright notice that is included in or attached to the work
(an example is provided in the Appendix below).

"Derivative Works" shall mean any work, whether in Source or Object
form, that is based on (or derived from) the Work and for which the
editorial revisions, annotations, elaborations, or other modifications
represent, as a whole, an original work of authorship. For the purposes
of this License, Derivative Works shall not include works that remain
separable from, or merely link (or bind by name) to the interfaces of,
the Work and Derivative Works thereof.

"Contribution" shall mean any work of authorship, including
the original version of the Work and any modifications or additions
to that Work or Derivative Works thereof, that is intentionally
submitted to Licensor for inclusion in the Work by the copyright owner
or by an individual or Legal Entity authorized to submit on behalf of
the copyright owner. For the purposes of this definition, "submitted"
means any form of electronic, verbal, or written communication sent
to the Licensor or its representatives, including but not limited to
communication on electronic mailing lists, source code control systems,
and issue tracking systems that are managed by, or on behalf of, the
Licensor for the purpose of discussing and improving the Work, but
excluding communication that is conspicuously marked or otherwise
designated in writing by the copyright owner as "Not a Contribution."

"Contributor" shall mean Licensor and any individual or Legal Entity
on behalf of whom a Contribution has been received by Licensor and
subsequently incorporated within the Work.

2. Grant of Copyright License. Subject to the terms and conditions of
this License, each Contributor hereby grants to You a perpetual,
worldwide, non-exclusive, no-charge, royalty-free, irrevocable
copyright license to reproduce, prepare Derivative Works of,
publicly display, publicly perform, sublicense, and distribute the
Work and such Derivative Works in Source or Object form.

3. Grant of Patent License. Subject to the terms and conditions of
this License, each Contributor hereby grants to You a perpetual,
worldwide, non-exclusive, no-charge, royalty-free, irrevocable
(except as stated in this section) patent license to make, have made,
use, offer to sell, sell, import, and otherwise transfer the Work,
where such license applies only to those patent claims licensable
by such Contributor that are necessarily infringed by their
Contribution(s) alone or by combination of their Contribution(s)
with the Work to which such Contribution(s) was submitted. If You
institute patent litigation against any entity (including a
cross-claim or counterclaim in a lawsuit) alleging that the Work
or a Contribution incorporated within the Work constitutes direct
or contributory patent infringement, then any patent licenses
granted to You under this License for that Work shall terminate
as of the date such litigation is filed.

4. Redistribution. You may reproduce and distribute copies of the
Work or Derivative Works thereof in any medium, with or without
modifications, and in Source or Object form, provided that You
meet the following conditions:

(a) You must give any other recipients of the Work or
    Derivative Works a copy of this License; and

(b) You must cause any modified files to carry prominent notices
    stating that You changed the files; and

(c) You must retain, in the Source form of any Derivative Works
    that You distribute, all copyright, patent, trademark, and
    attribution notices from the Source form of the Work,
    excluding those notices that do not pertain to any part of
    the Derivative Works; and

(d) If the Work includes a "NOTICE" text file as part of its
    distribution, then any Derivative Works that You distribute must
    include a readable copy of the attribution notices contained
    within such NOTICE file, excluding those notices that do not
    pertain to any part of the Derivative Works, in at least one
    of the following places: within a NOTICE text file distributed
    as part of the Derivative Works; within the Source form or
    documentation, if provided along with the Derivative Works; or,
    within a display generated by the Derivative Works, if and
    wherever such third-party notices normally appear. The contents
    of the NOTICE file are for informational purposes only and
    do not modify the License. You may add Your own attribution
    notices within Derivative Works that You distribute, alongside
    or as an addendum to the NOTICE text from the Work, provided
    that such additional attribution notices cannot be construed
    as modifying the License.

You may add Your own copyright statement to Your modifications and
may provide additional or different license terms and conditions
for use, reproduction, or distribution of Your modifications, or
for any such Derivative Works as a whole, provided Your use,
reproduction, and distribution of the Work otherwise complies with
the conditions stated in this License.

5. Submission of Contributions. Unless You explicitly state otherwise,
any Contribution intentionally submitted for inclusion in the Work
by You to the Licensor shall be under the terms and conditions of
this License, without any additional terms or conditions.
Notwithstanding the above, nothing herein shall supersede or modify
the terms of any separate license agreement you may have executed
with Licensor regarding such Contributions.

6. Trademarks. This License does not grant permission to use the trade
names, trademarks, service marks, or product names of the Licensor,
except as required for reasonable and customary use in describing the
origin of the Work and reproducing the content of the NOTICE file.

7. Disclaimer of Warranty. Unless required by applicable law or
agreed to in writing, Licensor provides the Work (and each
Contributor provides its Contributions) on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
implied, including, without limitation, any warranties or conditions
of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A
PARTICULAR PURPOSE. You are solely responsible for determining the
appropriateness of using or redistributing the Work and assume any
risks associated with Your exercise of permissions under this License.

8. Limitation of Liability. In no event and under no legal theory,
whether in tort (including negligence), contract, or otherwise,
unless required by applicable law (such as deliberate and grossly
negligent acts) or agreed to in writing, shall any Contributor be
liable to You for damages, including any direct, indirect, special,
incidental, or consequential damages of any character arising as a
result of this License or out of the use or inability to use the
Work (including but not limited to damages for loss of goodwill,
work stoppage, computer failure or malfunction, or any and all
other commercial damages or losses), even if such Contributor
has been advised of the possibility of such damages.

9. Accepting Warranty or Additional Liability. While redistributing
the Work or Derivative Works thereof, You may choose to offer,
and charge a fee for, acceptance of support, warranty, indemnity,
or other liability obligations and/or rights consistent with this
License. However, in accepting such obligations, You may act only
on Your own behalf and on Your sole responsibility, not on behalf
of any other Contributor, and only if You agree to indemnify,
defend, and hold each Contributor harmless for any liability
incurred by, or claims asserted against, such Contributor by reason
of your accepting any such warranty or additional liability.

END OF TERMS AND CONDITIONS

APPENDIX: How to apply the Apache License to your work.

To apply the Apache License to your work, attach the following
boilerplate notice, with the fields enclosed by brackets "[]"
replaced with your own identifying information. (Don't include
the brackets!)  The text should be enclosed in the appropriate
comment syntax for the file format. We also recommend that a
file or class name and description of purpose be included on the
same "printed page" as the copyright notice for easier
identification within third-party archives.

Copyright [yyyy] [name of copyright owner]

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
</file>

<file path="main_test.go">
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
</file>

<file path="main.go">
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
</file>

<file path="SECURITY.md">
# Security

`codex-responses-proxy` uses the same ChatGPT/Codex credentials as your local Codex CLI login.

## Recommended use

- Bind the proxy to `127.0.0.1` unless you have added your own authentication layer.
- Do not expose the proxy directly to the public internet.
- Protect `~/.codex/auth.json` or `$CODEX_HOME/auth.json`.
- Avoid `-debug` when prompts, responses, or request payloads may contain sensitive data.
- Treat any machine running this proxy as trusted.

## Secrets

Never commit Codex auth files, access tokens, refresh tokens, API keys, or captured debug logs. The repository `.gitignore` includes common patterns for these, but review changes before committing.

## Reporting issues

If you find a security issue, please report it privately to the repository owner rather than opening a public issue with exploit details.
</file>

<file path=".gitignore">
# Binaries and build output
/bin/
/build/
/dist/
*.exe
*.exe~
*.dll
*.so
*.dylib
*.test
*.out
codex-responses-proxy

# Go workspace/vendor artifacts
go.work
go.work.sum
/vendor/

# Coverage and profiling
coverage.out
coverage.html
*.coverprofile
*.prof
*.pprof

# Dependency/tool caches
.gocache/
.gomodcache/

# Environment and local config
.env
.env.*
!.env.example
*.local

# Codex/OpenAI auth material - never commit these
.codex/
auth.json

# Logs
*.log

# OS/editor files
.DS_Store
Thumbs.db
*.swp
*.swo
*~
.vscode/
.idea/
</file>

<file path="README.md">
# codex-responses-proxy

[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)

A small Go proxy that exposes an OpenAI-compatible **Responses API** endpoint backed by the OpenAI Codex backend used by the Codex CLI.

The initial use case is running API-oriented coding agents against models available through an existing ChatGPT/Codex subscription, without needing a separate OpenAI API key. It was built for Shelley, but anything that can talk to `POST /v1/responses` may be able to use it.

Inspired by Simon Willison's write-up, ["A pelican for GPT-5.5 via the semi-official Codex backdoor API"](https://simonwillison.net/2026/Apr/23/gpt-5-5/), and his reference implementation/plugin [`simonw/llm-openai-via-codex`](https://github.com/simonw/llm-openai-via-codex).

## Background

OpenAI's Codex CLI authenticates via ChatGPT and calls a Codex-specific endpoint:

```text
https://chatgpt.com/backend-api/codex/responses
```

Simon Willison documented using that endpoint to access models available through a Codex subscription. His post cites public comments from OpenAI folks indicating this pattern is intended to be supported for tools such as OpenCode, Pi, Claude Code, and similar coding environments.

This project adapts that idea into a local HTTP proxy:

```text
client -> http://127.0.0.1:8787/v1/responses -> chatgpt.com/backend-api/codex/responses
```

The proxy:

- reads Codex CLI auth from `~/.codex/auth.json` or `$CODEX_HOME/auth.json`
- refreshes expired ChatGPT access tokens using the stored refresh token
- forwards requests to the Codex backend with the required headers
- forces `stream: true`, because the Codex backend expects streaming
- converts Codex SSE responses back into a normal JSON Responses API object
- removes request fields the Codex backend rejects, such as `max_output_tokens`

## Status

Experimental. This depends on Codex backend behavior that may change.

It is **not** an official OpenAI API client, and it is **not** a way to avoid paying for access. You need a valid ChatGPT/Codex subscription and a working Codex CLI login.

## Requirements

- Go
- OpenAI Codex CLI installed
- Codex CLI authenticated with ChatGPT:

```bash
codex login --device-auth
```

That should create an auth file at:

```text
~/.codex/auth.json
```

or, if you use a custom Codex home:

```text
$CODEX_HOME/auth.json
```

The auth file must have:

```json
{
  "auth_mode": "chatgpt",
  "tokens": {
    "access_token": "...",
    "refresh_token": "..."
  }
}
```

## Run

```bash
go run .
```

By default the proxy listens on:

```text
127.0.0.1:8787
```

Useful flags:

```bash
go run . \
  -addr 127.0.0.1:8787 \
  -instructions "You are a helpful coding assistant." \
  -debug
```

Flags:

- `-addr`: listen address, default `127.0.0.1:8787`
- `-instructions`: default top-level instructions added when the request does not include any
- `-debug`: log patched request bodies and stream/debug details


## Install a binary

Once this repository is public, install the latest version with:

```bash
go install github.com/David-Factor/codex-responses-proxy@latest
```

This installs a `codex-responses-proxy` binary into your Go bin directory, usually `~/go/bin`.

## Run as a daemon

For Linux systems with systemd, this repository includes an example **user service**. A user service is preferable because the proxy needs access to your user-owned Codex auth file at `~/.codex/auth.json`.

Install the binary somewhere stable:

```bash
mkdir -p ~/.local/bin
go build -o ~/.local/bin/codex-responses-proxy .
```

Install and start the user service:

```bash
mkdir -p ~/.config/systemd/user
cp contrib/systemd/user/codex-responses-proxy.service ~/.config/systemd/user/
systemctl --user daemon-reload
systemctl --user enable --now codex-responses-proxy
```

Check logs:

```bash
journalctl --user -u codex-responses-proxy -f
```

Optional: keep the service running after logout on Linux hosts that support lingering:

```bash
loginctl enable-linger "$USER"
```

The example service binds to `127.0.0.1:8787`. Keep that default unless you add your own authentication layer.

## Endpoints

### `POST /v1/responses`

Primary endpoint. Configure clients to use:

```text
http://127.0.0.1:8787/v1
```

Then a client that normally calls `/v1/responses` will hit the proxy.

### `POST /responses`

Alias for manual testing.

### `GET /healthz`

Returns:

```text
ok
```

## Quick test

```bash
curl -s http://127.0.0.1:8787/v1/responses \
  -H 'Content-Type: application/json' \
  -d '{
    "model": "gpt-5.5",
    "input": "Say hello in exactly five words."
  }' | jq
```

The response should be a JSON object shaped like an OpenAI Responses API response, including an `output` array.

## Using with Shelley

Start the proxy:

```bash
go run .
```

Configure Shelley or another OpenAI Responses-compatible client to use:

```text
http://127.0.0.1:8787/v1
```

The proxy handles `/v1/responses`.

## How request patching works

Before forwarding to Codex, the proxy modifies the JSON payload:

- sets `stream: true`
- sets `store: false` if `store` was omitted
- adds top-level `instructions` if omitted
- deletes `max_output_tokens`
- deletes `max_completion_tokens`

These changes mirror the constraints of the Codex backend and behavior observed in similar integrations.

## How auth works

On each request, the proxy reads the Codex CLI auth document. It expects ChatGPT auth mode:

```text
auth_mode = chatgpt
```

If the access token is still valid, it is reused. If it is expired and a refresh token is present, the proxy refreshes it through:

```text
https://auth.openai.com/oauth/token
```

The updated tokens are written back to the Codex auth file with `0600` permissions.

If an account ID can be extracted from the tokens, the proxy forwards it as:

```text
ChatGPT-Account-ID: ...
```

## Limitations

- The proxy currently returns a completed JSON response, not a streaming response to the downstream client.
- It only implements the Responses endpoint needed by API-oriented agents.
- Tool/function/reasoning output items are preserved when Codex emits them, but compatibility has not been exhaustively tested.
- It depends on the Codex backend endpoint and auth file format remaining compatible.

## Security notes

See [SECURITY.md](SECURITY.md).

This proxy uses the same ChatGPT/Codex credentials as your Codex CLI login. Treat the machine running it as trusted.

Recommended defaults:

- keep it bound to `127.0.0.1`
- do not expose it publicly
- protect `~/.codex/auth.json`
- avoid running with `-debug` when prompts or outputs may contain sensitive data

## Development

Format and test:

```bash
gofmt -w main.go main_test.go
go test ./...
go vet ./...
```

Build:

```bash
go build .
```

## License

Apache-2.0. See [LICENSE](LICENSE).

## Related work

- [Simon Willison: A pelican for GPT-5.5 via the semi-official Codex backdoor API](https://simonwillison.net/2026/Apr/23/gpt-5-5/)
- [`simonw/llm-openai-via-codex`](https://github.com/simonw/llm-openai-via-codex)
- [`openai/codex`](https://github.com/openai/codex)
</file>

</files>
