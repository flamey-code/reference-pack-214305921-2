This file is a merged representation of the entire codebase, combined into a single document by Repomix.

# File Summary

## Purpose
This file contains a packed representation of the entire repository's contents.
It is designed to be easily consumable by AI systems for analysis, code review,
or other automated processes.

## File Format
The content is organized as follows:
1. This summary section
2. Repository information
3. Directory structure
4. Repository files (if enabled)
5. Multiple file entries, each consisting of:
  a. A header with the file path (## File: path/to/file)
  b. The full contents of the file in a code block

## Usage Guidelines
- This file should be treated as read-only. Any changes should be made to the
  original repository files, not this packed version.
- When processing this file, use the file path to distinguish
  between different files in the repository.
- Be aware that this file may contain sensitive information. Handle it with
  the same level of security as you would the original repository.

## Notes
- Some files may have been excluded based on .gitignore rules and Repomix's configuration
- Binary files are not included in this packed representation. Please refer to the Repository Structure section for a complete list of file paths, including binary files
- Files matching patterns in .gitignore are excluded
- Files matching default ignore patterns are excluded
- Files are sorted by Git change count (files with more changes are at the bottom)

# Directory Structure
```
.beads/.gitignore
.beads/config.yaml
.beads/issues.jsonl
.beads/metadata.json
.gitignore
AGENTS.md
assets/readme/pro-cli-hero.jpg
CLAUDE.md
docs/chatgpt-431-cookie-bloat.md
docs/chatgpt-chrome-cdp-cookies.md
docs/chatgpt-web-api-handbook.md
package.json
README.md
scripts/install.sh
scripts/probe-limits.ts
src/app.ts
src/args.ts
src/auth.ts
src/cdp.ts
src/cli.ts
src/config.ts
src/cookies.ts
src/daemon.ts
src/defaults.ts
src/errors.ts
src/executor.ts
src/jobs.ts
src/limits.ts
src/models.ts
src/odds.ts
src/output.ts
src/session-token.ts
src/structured.ts
src/transport.ts
src/update.ts
tests/args.test.ts
tests/cdp.test.ts
tests/cli.test.ts
tests/config.test.ts
tests/cookies.test.ts
tests/executor.test.ts
tests/jobs.test.ts
tests/limits.test.ts
tests/models.test.ts
tests/odds.test.ts
tests/output.test.ts
tests/session-token.test.ts
tests/structured.test.ts
tests/transport.test.ts
tests/update.test.ts
tsconfig.json
```

# Files

## File: .beads/.gitignore
`````
# Database
*.db
*.db-shm
*.db-wal

# Lock files
*.lock

# Temporary
last-touched
*.tmp
`````

## File: .beads/config.yaml
`````yaml
issue_prefix: pro
`````

## File: .beads/issues.jsonl
`````
{"id":"pro-1qt","title":"Add curl installer","description":"## Why\nREADME advertised manual clone/install only. Users wanted a one-command installer like curl -fsSL ... | bash.\n\n## Done When\n- scripts/install.sh clones or fast-forwards pro-cli safely.\n- README documents quick, custom-path, inspect-first, and manual install paths.\n- pro --version verifies linked CLI installs.","status":"closed","priority":2,"issue_type":"feature","created_at":"2026-05-07T19:53:58.632249Z","created_by":"jaredsmith","updated_at":"2026-05-07T19:54:02.194322Z","closed_at":"2026-05-07T19:54:02.194310Z","close_reason":"Implemented in 8deb624. Added scripts/install.sh, README quick/custom/inspect/manual install docs, pro --version support, and CLI coverage. Verified with bash -n scripts/install.sh, bun test, tsc --noEmit, and bun src/cli.ts --version.","source_repo":".","compaction_level":0,"original_size":0,"labels":["docs","installer"]}
{"id":"pro-y0g","title":"Epic: Surface ChatGPT web response metadata","description":"## Why\nThe ChatGPT web stream carries useful structured data that pro-cli currently discards: reasoning progress, resolved model metadata, continuation ids, limits, and resume hints. Exposing curated pieces would make pro-cli more transparent and agent-friendly without dumping sensitive raw stream data.\n\n## Scope\nDesign and implement a curated metadata layer for ChatGPT web responses, plus commands that use it. Keep raw cookies, tokens, profile metadata, and internal ids redacted unless explicitly needed for debugging.\n\n## Done When\n- Response metadata extraction is documented and tested.\n- Reasoning progress can surface during waits.\n- Successful calls return continuation ids.\n- Limits are available in a JSON command.\n- Discovery tooling exists for future endpoint/request-shape research.\n- README and docs/chatgpt-web-api-handbook.md stay aligned.","status":"open","priority":2,"issue_type":"epic","created_at":"2026-05-08T18:39:46.469429Z","created_by":"jaredsmith","updated_at":"2026-05-08T18:39:46.469429Z","source_repo":".","compaction_level":0,"original_size":0,"labels":["chatgpt-web-api","metadata","ux"]}
{"id":"pro-y0g.1","title":"Extract curated ChatGPT response metadata","description":"## Why\npro-cli currently treats the ChatGPT web stream mostly as answer text. The stream also carries continuation ids, resolved model state, reasoning progress, limits, and tool metadata that downstream commands need.\n\n## Done When\n- The response parser returns answer text plus a typed metadata object.\n- Metadata includes conversationId, assistant/input message ids, parent id, request/turn ids, resolved/default model slugs, reasoning markers, limits, and tool/search markers when present.\n- Raw auth, resume tokens, profile text, and other sensitive fields are redacted or omitted by default.\n- Tests cover snapshot and patch-style SSE fixtures, including missing-field cases.","status":"open","priority":2,"issue_type":"feature","created_at":"2026-05-08T18:40:45.428549Z","created_by":"jaredsmith","updated_at":"2026-05-08T18:40:45.428549Z","source_repo":".","compaction_level":0,"original_size":0,"labels":["chatgpt-web-api","metadata","transport"],"dependencies":[{"issue_id":"pro-y0g.1","depends_on_id":"pro-y0g","type":"parent-child","created_at":"2026-05-08T18:40:45.428549Z","created_by":"jaredsmith","metadata":"{}","thread_id":""}]}
{"id":"pro-y0g.2","title":"Show live Pro reasoning progress while waiting","description":"## Why\nChatGPT Pro exposes reasoning and progress markers during long jobs. Agents and humans should see useful live status instead of guessing whether a wait timeout means stalled, still thinking, or nearly done.\n\n## Done When\n- The daemon persists latest reasoning/progress metadata for each running job.\n- `pro-cli job wait`, `pro-cli job create --wait`, and any interactive/human output can show concise progress without spamming.\n- JSON output includes stable progress fields suitable for agents.\n- Timeout responses include the latest observed progress and next suggested command.\n- Tests cover progress updates, reconnects, and jobs with no progress metadata.","status":"open","priority":2,"issue_type":"feature","created_at":"2026-05-08T18:40:48.904987Z","created_by":"jaredsmith","updated_at":"2026-05-08T18:40:48.904987Z","source_repo":".","compaction_level":0,"original_size":0,"labels":["daemon","metadata","ux"],"dependencies":[{"issue_id":"pro-y0g.2","depends_on_id":"pro-y0g","type":"parent-child","created_at":"2026-05-08T18:40:48.904987Z","created_by":"jaredsmith","metadata":"{}","thread_id":""}]}
{"id":"pro-y0g.3","title":"Return continuation handles from successful results","description":"## Why\nThe web API returns enough ids to continue a conversation, but pro-cli does not yet expose them clearly. Continuation should be explicit, scriptable, and compatible with temporary conversations.\n\n## Done When\n- Successful `ask`, `job result`, and waited job responses return conversationId, assistantMessageId, parentMessageId/nextParentId, and request id when available.\n- Continuation flags are documented with clear temporary-versus-saved conversation behavior.\n- Continuing a conversation uses the same metadata fields, not scraped text or hidden state.\n- Tests cover continuation extraction and follow-up request shape construction.","status":"open","priority":2,"issue_type":"feature","created_at":"2026-05-08T18:40:52.291281Z","created_by":"jaredsmith","updated_at":"2026-05-08T18:40:52.291281Z","source_repo":".","compaction_level":0,"original_size":0,"labels":["conversation","metadata","ux"],"dependencies":[{"issue_id":"pro-y0g.3","depends_on_id":"pro-y0g","type":"parent-child","created_at":"2026-05-08T18:40:52.291281Z","created_by":"jaredsmith","metadata":"{}","thread_id":""}]}
{"id":"pro-y0g.4","title":"Add a pro-cli limits command","description":"## Why\nThe web stream and model metadata can expose usage limits and reset hints. A command that surfaces those hints would help agents choose wait/retry behavior and avoid blind failed calls.\n\n## Done When\n- `pro-cli limits --json` returns curated limit information from the latest live metadata and/or model endpoints when available.\n- Human output is concise and distinguishes known limits from unavailable/unknown values.\n- `doctor` points users to `limits` when auth is working but usage may be constrained.\n- Missing or changed web fields degrade to a clear unavailable state.\n- Tests cover known, unknown, and partial limit payloads.","status":"open","priority":3,"issue_type":"feature","created_at":"2026-05-08T18:40:56.849661Z","created_by":"jaredsmith","updated_at":"2026-05-08T18:40:56.849661Z","source_repo":".","compaction_level":0,"original_size":0,"labels":["doctor","limits","metadata"],"dependencies":[{"issue_id":"pro-y0g.4","depends_on_id":"pro-y0g","type":"parent-child","created_at":"2026-05-08T18:40:56.849661Z","created_by":"jaredsmith","metadata":"{}","thread_id":""}]}
{"id":"pro-y0g.5","title":"Add redacted ChatGPT web API capture tooling","description":"## Why\nFuture endpoint research should be repeatable and safe. Today exploration uses one-off scripts in /private/tmp, which makes it easy to lose useful probes or accidentally collect too much raw session data.\n\n## Done When\n- A developer/debug command captures selected ChatGPT web request and response shapes through the existing authenticated browser session.\n- The capture path redacts auth headers, cookies, resume tokens, account/profile fields, prompt bodies by default, and other sensitive values.\n- Output can be written to /private/tmp or an explicit path outside source control.\n- The command supports small request-shape experiments through explicit options.\n- docs/chatgpt-web-api-handbook.md explains the workflow and safety rules.","status":"open","priority":3,"issue_type":"task","created_at":"2026-05-08T18:41:00.555486Z","created_by":"jaredsmith","updated_at":"2026-05-08T18:41:00.555486Z","source_repo":".","compaction_level":0,"original_size":0,"labels":["chatgpt-web-api","debugging","research"],"dependencies":[{"issue_id":"pro-y0g.5","depends_on_id":"pro-y0g","type":"parent-child","created_at":"2026-05-08T18:41:00.555486Z","created_by":"jaredsmith","metadata":"{}","thread_id":""}]}
{"id":"pro-y0g.6","title":"Explore tool search and Deep Research request shapes","description":"## Why\nSearch, tool calls, file uploads, and Deep Research likely use related but distinct request fields and stream metadata. We should map those paths before promising first-class support.\n\n## Done When\n- CDP traces for normal search, Deep Research, and at least one tool/file-assisted flow are captured with the redacted tooling.\n- The handbook summarizes endpoints, request fields, stream events, metadata fields, and observed failure modes.\n- Stable fields are converted into implementation beads or added to the metadata extraction issue.\n- Unstable or unclear fields are explicitly marked as provisional.","status":"open","priority":3,"issue_type":"task","created_at":"2026-05-08T18:41:04.025807Z","created_by":"jaredsmith","updated_at":"2026-05-08T18:41:04.025807Z","source_repo":".","compaction_level":0,"original_size":0,"labels":["deep-research","research","tools"],"dependencies":[{"issue_id":"pro-y0g.6","depends_on_id":"pro-y0g","type":"parent-child","created_at":"2026-05-08T18:41:04.025807Z","created_by":"jaredsmith","metadata":"{}","thread_id":""}]}
`````

## File: .beads/metadata.json
`````json
{
  "database": "beads.db",
  "jsonl_export": "issues.jsonl"
}
`````

## File: .gitignore
`````
.env.local
.env.*.local
.secrets/
*cookies*.txt
*cookies*.json
*token*.json
node_modules/
dist/
coverage/
.pro/
.liquid-mail/
*.sqlite
*.sqlite-shm
*.sqlite-wal
`````

## File: AGENTS.md
`````markdown
## Project

`pro-query` is a Bun/TypeScript CLI for querying ChatGPT Pro through the user's existing authenticated browser session.

The goal is to let terminal tools ask ChatGPT Pro without requiring the user to manually paste context into the web UI and wait for a response.

## Constraints

- Build the project in Bun and TypeScript.
- Use the user's existing browser login cookies/session only with their consent.
- Treat this as a utility for legal ChatGPT subscribers using their own subscription.
- Do not bypass authentication, subscriptions, rate limits, access controls, or account restrictions.
- Do not build credential theft, session exfiltration, account sharing, or unauthorized access flows.

## Research First

Before implementation, research how the ChatGPT web app talks to its backend API.

Document:
- Browser endpoints and request shapes used by the web app
- Required headers, cookies, CSRF/session tokens, and streaming formats
- Model and reasoning-level options exposed by the web UI
- Conversation creation, message submission, polling/streaming, cancellation, and retry behavior
- Failure modes such as expired sessions, throttling, network interruption, and partially streamed responses

Prefer reproducible local observations from the authenticated browser session. Keep findings minimal and link to larger notes when needed.

When probing ChatGPT with test queries, use temporary conversations by default. Only create, save, or continue non-temporary conversations when the behavior under test specifically requires saved history, continuation, or sidebar-visible conversation state.

See `docs/chatgpt-chrome-cdp-cookies.md` for the validated local Chrome/CDP cookie export flow.
See `docs/chatgpt-431-cookie-bloat.md` when ChatGPT returns `431` or the CDP page resolves to `chrome-error://chromewebdata/`.

## CLI Requirements

- Async by default: submit work, return a job id, and allow later status/result collection.
- Resilient execution: support retries, reconnects, cancellation, timeout control, and recovery from interrupted streams.
- Configurable thinking levels and any other model/runtime options exposed by the web app.
- Clear session handling: detect expired auth and ask the user to refresh their browser login.
- Safe local storage: never print or persist raw cookies unless explicitly requested for debugging.
- Local cookie paths may live in ignored `.env.local` as `CHATGPT_COOKIE_JAR` and `CHATGPT_COOKIE_JSON`; the cookie files stay outside the repo.
- Script-friendly output: provide JSON output modes for other terminal CLIs.
`````

## File: CLAUDE.md
`````markdown
@AGENTS.md
`````

## File: docs/chatgpt-431-cookie-bloat.md
`````markdown
# ChatGPT 431 Cookie Bloat

Symptom: the pro-cli Chrome tab shows `chatgpt.com`, but CDP `location.href` is `chrome-error://chromewebdata/`; Chrome network events show `https://chatgpt.com/` returning HTTP `431`.

Cause seen 2026-05-09: 63 `conv_key_*` cookies pushed the ChatGPT request header over Cloudflare's limit. Deleting only `conv_key_*` cookies from the dedicated `~/.pro-cli/chrome-profile` restored the page, `limits`, and `gpt-5-5-pro` requests.

Cookie model: `ask` uses `fetch(..., credentials: "include")` inside the live CDP page, so Chrome's profile cookies are authoritative. Saved JSON/JAR cookies are capture artifacts and diagnostics unless an external tool replays them.

Rough fix: on a `431` auth/session/page response, automatically delete volatile ChatGPT cookies such as `conv_key_*`, reload `https://chatgpt.com/`, then retry once before asking the user to reset the profile.
`````

## File: docs/chatgpt-chrome-cdp-cookies.md
`````markdown
# ChatGPT Chrome CDP Cookie Export

Use this flow when the user consents to using their local Chrome login for ChatGPT research or CLI development.

## What Worked

1. Open `https://chatgpt.com/` in the real Chrome app and confirm the logged-in UI is visible.
2. Check whether the running Chrome already exposes CDP on `127.0.0.1:9222`.
3. If no CDP port exists, copy only the needed Chrome profile files into a private temp profile:
   - `Local State`
   - `Default/Cookies`
   - `Default/Preferences`
4. Launch a separate Chrome instance with the copied profile and CDP enabled:

```bash
open -na "Google Chrome" --args \
  --user-data-dir=/private/tmp/pro-query-chrome-profile \
  --remote-debugging-port=9222 \
  --no-first-run \
  --no-default-browser-check \
  https://chatgpt.com
```

5. Attach to the browser with `agent-browser --cdp 9222`.
6. Verify login through the accessibility snapshot. The successful check showed the account UI, recent chats, prompt textbox, and `Extended Pro` model control.
7. Export cookies through direct CDP, scoped to ChatGPT/OpenAI URLs:
   - `https://chatgpt.com/`
   - `https://auth.openai.com/`
   - `https://openai.com/`
   - `https://sentinel.openai.com/`
   - `https://ws.chatgpt.com/`
8. Write the result as `0600` JSON and Netscape cookie jars. Do not print raw cookie values.
9. Close the temporary Chrome instance and delete the copied profile.

## Notes

- Chrome stores these cookies encrypted. Reading the SQLite rows is enough to list names and metadata, but not values.
- The Keychain item `Chrome Safe Storage` existed, but `security find-generic-password -w ...` could not read the secret from the shell session.
- Letting Chrome decrypt its own copied cookie DB through CDP worked.
- `agent-browser cookies get` is too broad for this task because it can dump unrelated browser cookies. Use direct CDP `Network.getCookies` with scoped URLs instead.
- A standalone `curl` call to `https://chatgpt.com/api/auth/session` with the cookie jar returned Cloudflare `403`; that did not invalidate the cookies. The useful validation came from the logged-in Chrome/CDP page state.

## Cleanup Checklist

- Remove the copied Chrome profile under `/private/tmp/pro-query-chrome-profile`.
- Remove any overbroad cookie dumps.
- Keep scoped cookie jars private with mode `0600`.
- Never commit cookie jars or raw session values.
`````

## File: docs/chatgpt-web-api-handbook.md
`````markdown
# ChatGPT Web API Handbook

This handbook records what `pro-cli` has validated about the ChatGPT web app API through local Chrome/CDP observation as of May 8, 2026.

`pro-cli` uses a logged-in ChatGPT browser tab. It does not call the public OpenAI API. It submits requests from inside the web page so the request carries the same cookies, session token, frontend headers, sentinel headers, and streaming behavior as ChatGPT web.

These endpoints are private ChatGPT web endpoints, not a stable API contract. Treat field names and request shapes as observed behavior that needs fresh verification before broad changes.

## Runtime Path

The normal request path is:

```txt
pro-cli command
  -> ~/.pro-cli auth/session state
  -> Chrome DevTools Protocol page evaluation
  -> https://chatgpt.com/backend-api/f/conversation
  -> streamed SSE response
  -> local job/result storage
```

The Chrome window must stay open while jobs run. The local daemon manages job execution; the user or agent manages the browser login.

Do not commit raw captures. Request bodies and streams can include prompts, resume tokens, profile/context metadata, request ids, and account/session-derived state.

## Observed Endpoints

Live ChatGPT web traffic used these endpoints during Pro conversation requests:

```txt
GET  /api/auth/session
GET  /backend-api/models
GET  /backend-api/conversations
POST /backend-api/f/conversation
POST /backend-api/f/conversation/prepare
GET  /backend-api/f/conversation/resume
POST /backend-api/sentinel/chat-requirements/prepare
POST /backend-api/sentinel/chat-requirements/finalize
```

The conversation endpoint is the primary submission endpoint. The prepare and sentinel endpoints produce request validation state. The resume endpoint supports long streams and reconnects.

### Account / Plan / PII Endpoints

Probed via CDP page evaluation (May 8, 2026):

```txt
GET  /backend-api/accounts/check/v4-2023-04-27   200  account + plan + features (no remaining counters)
GET  /backend-api/me                             200  email, name, phone, MFA flag, picture (PII)
GET  /public-api/conversation_limit              200  {"message_cap":0.0,...} (zeroed, not useful)
```

`accounts/check/v4-2023-04-27` is the canonical place for plan facts. The shape includes:

```json
{
  "accounts": {
    "<account-uuid>": {
      "account": { "plan_type": "pro", "structure": "personal", ... },
      "entitlement": {
        "subscription_plan": "chatgptpro",
        "has_active_subscription": true,
        "expires_at": "...",
        "renews_at": "...",
        "billing_period": "monthly"
      },
      "features": ["gpt5_pro", "o3_pro", "canvas", ...]
    }
  }
}
```

`/backend-api/me` returns user PII — handle with care. `pro-cli` should not log or persist its body.

### Endpoints That Do Not Exist

Probed and confirmed 404/405 (do not re-probe):

```txt
/backend-api/conversation_limits_progress   404
/backend-api/conversation_limit             404
/backend-api/conversation_limits            404
/public-api/conversation_limit/v2           404
/backend-api/me/usage                       404
/backend-api/me/limits                      404
/backend-api/me/quota                       404
/backend-api/me/feature_limits              404
/backend-api/usage                          404
/backend-api/usage_metrics                  404
/backend-api/billing/usage                  404
/backend-api/billing/subscription           404
/backend-api/subscription                   404
/backend-api/feature_limits                 404
/backend-api/limits                         404
/backend-api/limits_progress                404
/backend-api/rate_limits                    404
/backend-api/conversation_meta              404
/backend-api/account/check                  404
/backend-api/accounts/check                 405
```

**Conclusion: there is no standalone "remaining calls" endpoint.** Per-feature counters only appear inside the SSE stream as `conversation_detail_metadata.limits_progress` events on real conversation turns (see Limits section below). Pro general chat throttling is adaptive; no published cap exists.

## Request Shape

`pro-cli` currently sends a body shaped like:

```json
{
  "action": "next",
  "messages": [
    {
      "id": "<uuid>",
      "author": { "role": "user" },
      "create_time": 1778265252,
      "content": {
        "content_type": "text",
        "parts": ["<prompt>"]
      },
      "metadata": {}
    }
  ],
  "model": "gpt-5-5-pro",
  "thinking_effort": "standard",
  "parent_message_id": "client-created-root",
  "client_prepare_state": "none",
  "timezone_offset_min": 360,
  "timezone": "America/Denver",
  "conversation_mode": { "kind": "primary_assistant" },
  "enable_message_followups": true,
  "system_hints": [],
  "supports_buffering": true,
  "supported_encodings": ["v1"],
  "client_contextual_info": { "app_name": "chatgpt.com" },
  "paragen_cot_summary_display_override": "allow",
  "force_parallel_switch": "auto",
  "history_and_training_disabled": true
}
```

Continuing a saved conversation adds:

```json
{
  "conversation_id": "<conversation-id>",
  "parent_message_id": "<parent-message-id>",
  "history_and_training_disabled": false
}
```

## Response Stream

The `/backend-api/f/conversation` response is an SSE stream. A simple Pro request produced these event shapes:

```txt
event: delta_encoding
data: "v1"

data: {"type":"resume_conversation_token", ...}
data: {"p":"","o":"add","v":{"message":{...},"conversation_id":"..."},"c":0}
data: {"type":"input_message", ...}
data: {"type":"stream_handoff", ...}
data: {"type":"server_ste_metadata", ...}
data: {"type":"conversation_detail_metadata", ...}
data: {"type":"message_stream_complete", ...}
data: [DONE]
```

The final assistant text appears in message content patches or assistant message snapshots:

```json
{
  "message": {
    "author": { "role": "assistant" },
    "content": {
      "content_type": "text",
      "parts": ["answer text"]
    },
    "status": "finished_successfully",
    "end_turn": true
  }
}
```

## Useful Metadata

The stream exposes useful metadata that `pro-cli` does not fully surface yet.

### Conversation and Continuation

Useful fields:

```txt
conversation_id
message.id
input_message.id
parent_id
request_id
turn_exchange_id
turn_trace_id
```

These can support better continuation UX. `pro-cli` could return `conversationId`, `assistantMessageId`, and `parentMessageId` after successful calls so agents do not need to inspect ChatGPT manually.

### Reasoning Progress

Useful fields:

```txt
message.metadata.initial_text
message.metadata.finished_text
message.metadata.finished_duration_sec
message.metadata.reasoning_start_time
message.metadata.reasoning_end_time
message.metadata.pro_progress
message.metadata.thinking_effort
```

Observed examples:

```txt
initial_text: Reasoning
finished_text: Finished reasoning
finished_text: Thought for 7s
finished_duration_sec: 7
pro_progress: 71.42857142857143
thinking_effort: standard
```

These can improve `job wait` and `job create --wait` by showing live progress to the agent or terminal user.

### Model Resolution

Useful fields:

```txt
model_slug
resolved_model_slug
default_model_slug
did_auto_switch_to_reasoning
auto_switcher_race_winner
thinking_effort
```

These fields can prove which model actually handled the turn. This is useful because `pro-cli` defaults to `gpt-5-5-pro`, while ChatGPT web's picker may have its own default.

### Limits

`conversation_detail_metadata` exposed:

```json
{
  "limits_progress": [
    {
      "feature_name": "deep_research",
      "remaining": 250,
      "reset_after": "2026-06-07T18:34:14.421525+00:00"
    },
    {
      "feature_name": "odyssey",
      "remaining": 398,
      "reset_after": "2026-05-17T21:31:20.421544+00:00"
    }
  ],
  "model_limits": []
}
```

Wired into `pro-cli limits` as a stream-side capture: `transport.ts` extracts `limits_progress` from `conversation_detail_metadata` events and persists snapshots to the local SQLite. `pro-cli limits` returns plan info from `accounts/check` plus the most recent observed counters.

Observed `feature_name` values so far: `deep_research`, `odyssey`. These are specialty-feature quotas, not general chat caps. Free-tier features (gpt5, etc.) do not appear here on Pro.

### Tools and Search

Useful fields:

```txt
tool_invoked
tool_name
is_search
search_tool_call_count
search_tool_query_types
citations
content_references
search_result_groups
```

These fields are usually null or empty for a plain text request. They are likely useful when exploring Deep Research, web search, file handling, and tool-enabled modes.

## Logprobs Finding

A live experiment tried these request variants:

```json
{ "logprobs": true, "top_logprobs": 5 }
{ "include_logprobs": true, "top_logprobs": 5 }
{ "response_options": { "logprobs": true, "top_logprobs": 5 } }
```

All variants returned HTTP 200 and normal answer text. None produced token probability fields in the response stream. The web endpoint appears to ignore those fields, at least for `gpt-5-5-pro` through `/backend-api/f/conversation`.

This does not prove no internal endpoint can expose logprobs. It means the current web conversation path does not expose them in the stream.

## Exploration Playbook

Use CDP network tracing and keep each experiment small.

Good toggles to trace:

- Model picker changes
- `standard` vs `extended` thinking
- Temporary vs saved conversations
- Continuing a conversation
- Regenerate and retry
- Cancel during reasoning
- Search/web tool toggles
- Deep Research mode
- File upload
- Image generation
- Canvas or agent mode

For each experiment, capture:

- Request URL and method
- Request body
- Response status
- Raw SSE event types
- New fields or changed fields
- Whether fields are stable across two runs
- Whether the data is safe to expose in normal output

Prefer curated metadata over raw stream dumps. Raw streams may include resume tokens, profile/context metadata, request ids, and account/session-derived state.

Good follow-up questions:

- Which fields are stable across `standard` and `extended` reasoning?
- Which continuation ids are required for saved conversations, temporary conversations, regenerate, and retry?
- Can Deep Research be started with the same conversation endpoint, or does it require a separate mode/tool path?
- Which limit fields are available before a job starts?
- Can resume events be used to recover long-running jobs after daemon restart?

## Upgrade Candidates

The highest-value `pro-cli` upgrades are:

1. Return curated response metadata with full results.
2. Show live reasoning progress while waiting.
3. Return continuation ids after successful calls.
4. Add limits reporting.
5. Add a redacted response-shape capture tool for future discovery.
`````

## File: package.json
`````json
{
  "name": "pro-cli",
  "version": "0.1.0",
  "description": "Agent native CLI for querying ChatGPT Pro and Deep Research through your own logged-in web session, managed from your terminal.",
  "type": "module",
  "homepage": "https://github.com/ratacat/pro-cli#readme",
  "repository": {
    "type": "git",
    "url": "git+https://github.com/ratacat/pro-cli.git"
  },
  "bugs": {
    "url": "https://github.com/ratacat/pro-cli/issues"
  },
  "keywords": [
    "chatgpt",
    "chatgpt-pro",
    "deep-research",
    "cli",
    "bun",
    "typescript",
    "agents",
    "automation"
  ],
  "bin": {
    "pro-cli": "./src/cli.ts"
  },
  "scripts": {
    "test": "bun test",
    "typecheck": "tsc --noEmit",
    "check": "bun test && tsc --noEmit"
  },
  "devDependencies": {
    "@types/bun": "^1.3.5",
    "typescript": "^5.9.3"
  }
}
`````

## File: README.md
`````markdown
# pro-cli

<p align="center">
  <img src="assets/readme/pro-cli-hero.jpg" alt="Neon ChatGPT Pro interface globe for pro-cli" width="100%">
</p>

Agent native CLI for querying ChatGPT Pro and Deep Research through your own logged-in web session, managed from your terminal.

`pro-cli` gives coding agents a JSON-first command surface for the ChatGPT web account you already use: Pro models, thinking levels, Deep Research-style capabilities when available to your account, structured JSON outputs with schema validation, calibrated probability scoring, async jobs, and recoverable results.

`pro-cli` is built for legal ChatGPT subscribers using their own session. It does not bypass authentication, subscriptions, rate limits, access controls, or account restrictions.

## Install

Requires Bun. `pro-cli` is a Bun/TypeScript CLI with the binary name `pro-cli`. It has been tested on macOS; Windows may need install or Chrome auth command adjustment.

```sh
curl -fsSL https://raw.githubusercontent.com/ratacat/pro-cli/main/scripts/install.sh | bash
```

To choose the checkout path:

```sh
curl -fsSL https://raw.githubusercontent.com/ratacat/pro-cli/main/scripts/install.sh | PRO_INSTALL_DIR="$HOME/Projects/pro-cli" bash
```

The installer clones or fast-forwards `~/Projects/pro-cli`, runs `bun install`, runs `bun link`, and prints `pro-cli --version`. It does not touch auth, cookies, Chrome, or `~/.pro-cli`.

Update an existing clean install:

```sh
pro-cli update --json
```

`update` verifies the checkout is `main` with the expected origin and no uncommitted changes, then runs `git pull --ff-only origin main`, `bun install`, and `bun link`.

## Agent Instructions

After installing, add a short `pro-cli` note to your user-level or project-level `AGENTS.md` or equivalent agent instructions file:

```md
Use `pro-cli` to answer real, user-driven questions or tasks that clearly benefit from ChatGPT Pro. Run `pro-cli --help` if you need the command list or are unsure which command shape to use. Avoid probe or smoke-test queries: do not call `pro-cli ask` for checks after errors or empty responses; use `pro-cli doctor --json` for health/setup validation because it does not consume Pro quota. Submit durable blocking tasks with `pro-cli job create @prompt.md --wait --json` or direct blocking requests with `pro-cli ask @prompt.md --json`, and never include secrets, raw cookies, tokens, `.env` files, or private keys.
```

## Setup

`pro-cli` needs one logged-in ChatGPT Chrome window. `pro-cli` manages its local job daemon; you manage the browser login.

Choose one auth path; you do not need both. Both end with the same local `pro-cli` auth state.

**Agent-assisted auth: existing Chrome profile**

Use this when you are already logged in to ChatGPT in Chrome and trust the current agent with temporary local browser access:

```txt
I am logged in to ChatGPT in Chrome. Set up pro-cli auth from my existing Chrome profile. Store only scoped ChatGPT/OpenAI auth under ~/.pro-cli, do not print raw cookies or tokens, then verify with pro-cli doctor --json.
```

This is the lowest-friction path. It uses a browser profile that already has your ChatGPT session, so it also exposes that profile over Chrome DevTools Protocol while the CDP Chrome window is open.

**Manual auth: dedicated Chrome profile**

Use this when you want a separate browser profile for `pro-cli`:

```sh
pro-cli auth command --json
```

Run the returned Chrome command, sign in to ChatGPT in that window, then capture auth:

```sh
pro-cli auth capture --cdp http://127.0.0.1:9222 --json
pro-cli doctor --json
```

This is the normal long-running path. It creates `~/.pro-cli/chrome-profile`, keeps ChatGPT auth separate from your personal Chrome profile, and limits what the open debugging port can see.

Port `9222` is the default. If you use another port, pass the same `--cdp` or `--port` to `doctor`, `ask`, and `job create`. `job wait` uses the CDP value stored on the job.

## Runtime Model

Keep the ChatGPT Chrome window open while jobs run. `pro-cli` sends requests from that logged-in tab over Chrome DevTools Protocol, so it gets the same cookies, page session, frontend headers, and streaming/resume behavior as ChatGPT in the browser.

Async jobs run through a local `pro-cli` daemon. `job create`, `job wait`, and `job cancel` start or restart it when needed, so agents do not need to manage a background process. The daemon processes `~/.pro-cli/jobs.sqlite`; those commands reach it through a localhost control endpoint stored under `/tmp`.

Use the dedicated `~/.pro-cli/chrome-profile` window for normal operation. A personal Chrome profile can work, but the debugging port exposes that profile while it is open. The dedicated profile limits scope to ChatGPT.

If Chrome closes, run `pro-cli auth command --json` again. If ChatGPT logs out, sign in and rerun:

```sh
pro-cli auth capture --cdp <url> --json
```

When unsure, run:

```sh
pro-cli doctor --json
```

## Commands

Agents should use `--json`; non-TTY stdout switches to JSON automatically.

Setup and auth:

```sh
pro-cli setup --json
pro-cli update --json
pro-cli auth command --json
pro-cli auth capture --cdp http://127.0.0.1:9222 --json
pro-cli doctor --json
```

Models and defaults:

```sh
pro-cli models --json
pro-cli config get --json
pro-cli config set model gpt-5-5-pro --json
pro-cli config set reasoning extended --json
```

Async jobs:

```sh
pro-cli job create @prompt.md --json
pro-cli job create @prompt.md --wait --json
pro-cli job create @prompt.md --reasoning extended --json
pro-cli job create @prompt.md --condensed-response 500 --json
pro-cli job wait <job-id> --json
pro-cli job wait <job-id> --soft-timeout 60000 --json
pro-cli job result <job-id> --json
pro-cli job cancel <job-id> --json
pro-cli job list --limit 20 --json
```

`job wait` without a timeout waits until the job succeeds, fails, or is cancelled. Long prompts and `--reasoning extended` can run for several minutes. Use `--soft-timeout <ms>` when an agent needs to poll without a nonzero exit. Use `--wait-timeout <ms>` only when a timeout should fail the local command.

Daemon:

```sh
pro-cli daemon status --json
pro-cli daemon restart --json
pro-cli daemon stop --json
```

Direct ask:

```sh
pro-cli ask @prompt.md --json
pro-cli ask @prompt.md --reasoning extended --json
pro-cli ask @prompt.md --condensed_response=500 --json
```

`ask` executes without creating durable job state. Use `job create` when you need a job id that later `job wait`, `job result`, `job cancel`, or `job list` can inspect.

Use `--condensed-response <tokens>` when Pro should keep the final answer within an approximate response budget. The underscore alias `--condensed_response=<tokens>` is also accepted for agents. This is a prompt-level instruction, not a second summarization call, so it does not spend extra Pro quota.

JSON responses that include full Pro text also include `agentInstruction` and `resultStats`. Agents should treat `data.result` as the primary deliverable. Results under 6000 characters should usually be relayed in full; longer results may be condensed with care for the original prose, structure, and voice.

Probability and plan:

```sh
pro-cli odds "Will X happen?" --context @evidence.md
pro-cli limits --json
```

## Structured Outputs

`pro-cli` wraps your prompt with strict JSON instructions, parses the model's reply, validates it, and retries on failure. Use this when an agent or script needs a typed result instead of prose.

Quick, with a free-form format hint:

```sh
pro-cli ask "Extract the people from this article" \
  --format '{people: [{name: string, role: string}]}'
```

Rigorous, with a JSON Schema (validates the parsed value):

```sh
pro-cli ask "Find 3 fictional spies" \
  --schema @people.schema.json --json
```

The CLI strips fenced ```` ```json ```` blocks (with a balanced-bracket fallback that handles braces inside strings), parses, and validates the root type plus top-level `required` fields. On parse or validation failure, it retries up to `--schema-retries <n>` times (default 1), feeding the previous failed response and the failure reason back to the model.

With `--json`, the envelope includes `parsed`, `raw`, and `attempts`. Without `--json`, the parsed JSON is pretty-printed to stdout, ready for `jq` or another tool. The same flags work on the durable job path:

```sh
pro-cli job create @prompt.md --wait --schema @file.json --json
```

## Probability Scoring

`pro-cli odds` is a yes/no probability assessor. It wraps your question with strict integer-only output instructions and returns a single integer 0–100 representing P(YES). Useful for prediction-market scoring, threshold gates in agent pipelines, or any place a calibrated number beats prose.

```sh
pro-cli odds "Will the deploy ship by Friday?" --context @evidence.md
# → 78

pro-cli odds "..." --samples 5 --aggregate median --json
```

Bare integer to stdout by default for shell-friendly consumption (`prob=$(pro-cli odds "...")`). `--samples N` runs N calls and aggregates (`mean` default; `median`, `trimmed-mean` available). `--allow-fifty` permits 50; the default forbids it to force a directional commitment, retrying up to `--parse-retries` if the model returns 50. `--json` returns the full envelope with per-sample attempts and job ids.

## Plan and Observed Limits

```sh
pro-cli limits --json
```

Returns:

- Plan facts from `accounts/check`: `plan_type`, `subscription_plan`, `expires_at`, `renews_at`, `billing_period`, `features`.
- Per-feature counters (e.g. `deep_research`, `odyssey`) captured from the ChatGPT stream metadata of recent `ask`/`odds`/`job` calls, with `observed_at` timestamps.

ChatGPT does not expose a standalone limits endpoint, so counters refresh whenever you make a real Pro call. General Pro chat throttling is adaptive and not exposed by any endpoint we found.

## Conversations

New `ask` and `job create` requests default to temporary ChatGPT conversations. Use `--save` when the turn should be written to ChatGPT history.

Continuing a conversation defaults to saved mode and requires both ids from the ChatGPT conversation:

```sh
pro-cli ask "follow up" --save --conversation <conversation-id> --parent <message-id> --json
```

## Thinking Modes

By default, `pro-cli` sends `--model gpt-5-5-pro --reasoning standard`. The request includes `thinking_effort=standard`, so it uses the Pro model rather than ChatGPT web's default picker.

For deeper Pro reasoning:

```sh
--reasoning extended
```

Use exact web effort values:

```txt
standard
extended
min
max
```

Pro models normally expose `standard` and `extended`. Thinking models may expose `min`, `standard`, `extended`, and `max`. For explicit model ids, use an effort shown by:

```sh
pro-cli models --json
```

## Request Controls

```sh
--model <id from pro-cli models>
--reasoning min|standard|extended|max
--cdp http://127.0.0.1:9222
--port 9222
--temporary
--save
--conversation <conversation-id>
--parent <message-id>
--verbosity low|medium|high
--instructions "system text"
--instructions-file prompt-system.txt
--reasoning-summary auto|concise|detailed|none
--tool-choice auto|none|required
--parallel-tools true|false
--timeout <ms>
--retries <0..5>
--retry-delay <ms>
--schema @schema.json | --schema "<inline JSON Schema>"
--format "<inline format hint>"
--schema-retries <0..5>
```

Probability (`odds`) only:

```sh
--context @evidence.md | --context "inline context"
--samples <1..25>
--aggregate mean|median|trimmed-mean
--parse-retries <0..5>
--allow-fifty
```

Job wait controls:

```sh
--wait
--soft-timeout <ms>
--wait-timeout <ms>
--poll-ms <ms>
```

Unsupported request flags fail loudly. Errors include `code`, `message`, and `suggestions`.

## Safety

`pro-cli` uses a browser session you control. The recommended setup opens a dedicated Chrome profile, captures scoped ChatGPT/OpenAI cookies plus the page session token, and stores them under `~/.pro-cli`.

Treat `~/.pro-cli` like SSH keys or browser session data. Do not commit it, paste it, sync it, or share it.

The daemon control file lives under `/tmp/pro-cli-*` with a local bearer token and private file mode. It points commands at the daemon; it does not contain ChatGPT cookies or session tokens.

If an older default `~/.pro` directory exists and `~/.pro-cli` does not, the first non-help command moves it to `~/.pro-cli` and rewrites stored paths that pointed inside the old directory. Set `PRO_CLI_HOME` to use a different location.

Normal setup, doctor, job, and status output redacts raw cookies and tokens. Files use private permissions where supported: `0600` for files and `0700` for directories. Requests go to `https://chatgpt.com`.
`````

## File: scripts/install.sh
`````bash
#!/usr/bin/env bash
set -euo pipefail

REPO_URL="https://github.com/ratacat/pro-cli.git"

log() {
  printf 'pro-cli install: %s\n' "$1"
}

die() {
  printf 'pro-cli install: %s\n' "$1" >&2
  exit 1
}

require_command() {
  if ! command -v "$1" >/dev/null 2>&1; then
    die "missing required command: $1"
  fi
}

verify_origin() {
  local origin_url

  if ! origin_url="$(git -C "$INSTALL_DIR" remote get-url origin)"; then
    die "$INSTALL_DIR has no origin remote"
  fi

  case "$origin_url" in
    "$REPO_URL" | "https://github.com/ratacat/pro-cli" | "git@github.com:ratacat/pro-cli.git")
      ;;
    *)
      die "$INSTALL_DIR origin is $origin_url, expected $REPO_URL"
      ;;
  esac
}

update_existing_checkout() {
  local branch

  verify_origin

  branch="$(git -C "$INSTALL_DIR" branch --show-current)"
  if [ -z "$branch" ]; then
    die "$INSTALL_DIR is not on a branch; check out main before reinstalling"
  fi

  if [ "$branch" != "main" ]; then
    die "$INSTALL_DIR is on branch $branch; switch to main before reinstalling"
  fi

  if [ -n "$(git -C "$INSTALL_DIR" status --porcelain)" ]; then
    die "$INSTALL_DIR has uncommitted changes; commit or stash them before reinstalling"
  fi

  log "updating $INSTALL_DIR"
  git -C "$INSTALL_DIR" pull --ff-only origin main
}

clone_checkout() {
  if [ -e "$INSTALL_DIR" ] && [ ! -d "$INSTALL_DIR" ]; then
    die "$INSTALL_DIR exists but is not a directory"
  fi

  if [ -d "$INSTALL_DIR" ] && [ -n "$(find "$INSTALL_DIR" -mindepth 1 -maxdepth 1 -print -quit)" ]; then
    die "$INSTALL_DIR exists but is not a pro-cli git checkout"
  fi

  mkdir -p "$(dirname "$INSTALL_DIR")"
  log "cloning $REPO_URL into $INSTALL_DIR"
  git clone --branch main "$REPO_URL" "$INSTALL_DIR"
}

if [ -n "${PRO_INSTALL_DIR:-}" ]; then
  INSTALL_DIR="$PRO_INSTALL_DIR"
else
  if [ -z "${HOME:-}" ]; then
    die "HOME is not set; set PRO_INSTALL_DIR"
  fi
  INSTALL_DIR="${HOME}/Projects/pro-cli"
fi

if [ -z "$INSTALL_DIR" ]; then
  die "PRO_INSTALL_DIR resolved to an empty path"
fi

require_command git
require_command bun

if [ -d "$INSTALL_DIR/.git" ]; then
  update_existing_checkout
else
  clone_checkout
fi

cd "$INSTALL_DIR"

log "installing dependencies"
bun install

log "linking pro-cli"
bun link

if ! command -v pro-cli >/dev/null 2>&1; then
  die "bun link completed, but pro-cli is not on PATH"
fi

log "installed $(pro-cli --version)"
`````

## File: scripts/probe-limits.ts
`````typescript
#!/usr/bin/env bun
import { evaluateInCdpPage } from "../src/cdp";

const CDP_BASE = process.env.CDP_BASE ?? "http://127.0.0.1:9222";

const CANDIDATES = [
  "/backend-api/accounts/check/v4-2023-04-27",
  "/backend-api/me/feature_limits",
  "/backend-api/me/limits",
  "/backend-api/me/quota",
  "/backend-api/feature_limits",
  "/backend-api/conversation_limits",
  "/backend-api/limits_progress",
  "/backend-api/billing/subscription",
  "/backend-api/subscription",
  "/backend-api/billing/usage",
  "/backend-api/usage_metrics",
];

const FULL_BODY_PATHS = new Set(["/backend-api/accounts/check/v4-2023-04-27"]);

interface ProbeResult {
  path: string;
  ok: boolean;
  status: number;
  preview: string;
  hasLimits: boolean;
}

const expression = `(${async function probe(paths: string[]): Promise<ProbeResult[]> {
  const sessionResponse = await fetch("https://chatgpt.com/api/auth/session", { credentials: "include" });
  const session = (await sessionResponse.json().catch(() => null)) as { accessToken?: unknown } | null;
  if (!sessionResponse.ok || typeof session?.accessToken !== "string") {
    throw new Error(`Session fetch failed: ${sessionResponse.status}`);
  }
  const accessToken = session.accessToken;
  const results: ProbeResult[] = [];
  for (const path of paths) {
    try {
      const response = await fetch(`https://chatgpt.com${path}`, {
        method: "GET",
        credentials: "include",
        headers: {
          authorization: `Bearer ${accessToken}`,
          accept: "application/json",
        },
      });
      const text = await response.text();
      const preview = text.slice(0, 8000).replace(/\s+/g, " ");
      const hasLimits = /limit|remaining|reset_after|quota|usage|cap/i.test(preview);
      results.push({ path, ok: response.ok, status: response.status, preview, hasLimits });
    } catch (error) {
      results.push({
        path,
        ok: false,
        status: 0,
        preview: String(error),
        hasLimits: false,
      });
    }
  }
  return results;
}})(${JSON.stringify(CANDIDATES)})`;

const results = await evaluateInCdpPage<ProbeResult[]>(CDP_BASE, expression, 60_000);

for (const result of results) {
  const tag = result.ok ? (result.hasLimits ? "HIT " : "ok  ") : "MISS";
  console.log(`${tag} ${result.status} ${result.path}`);
  if (result.ok) {
    const dumpFull = FULL_BODY_PATHS.has(result.path);
    console.log(dumpFull ? `--- FULL BODY ---\n${result.preview}\n--- END ---` : `     ${result.preview.slice(0, 320)}`);
  }
}
`````

## File: src/app.ts
`````typescript
import { readFile } from "node:fs/promises";
import { join } from "node:path";
import packageJson from "../package.json" with { type: "json" };
import { flagBoolean, flagString, parseArgs } from "./args";
import {
  captureAuth,
  defaultCdpBase,
  detectLegacyArtifacts,
  detectPortCollision,
  getAuthStatus,
  getBrowserSessionStatus,
  hideProCliChromeWindow,
  resetAuthProfile,
  showProCliChromeWindow,
} from "./auth";
import { loadConfig, migrateLegacyDefaultHome, resolvePaths, saveConfig } from "./config";
import { ensureDaemonRunning, getDaemonStatus, runDaemonServer, stopDaemon } from "./daemon";
import { DEFAULT_MODEL, DEFAULT_REASONING, REASONING_LEVELS, isReasoningLevel } from "./defaults";
import { EXIT, ProError, toProError } from "./errors";
import { buildEphemeralJob, executeEphemeralJob } from "./executor";
import { JobStore, redactJob } from "./jobs";
import { fetchAccountSummary } from "./limits";
import {
  canonicalModelId,
  listModels,
  modelRequiresSavedConversation,
  modelUsesThinkingEffort,
  NO_REASONING,
} from "./models";
import { runOdds, type AggregateMethod } from "./odds";
import { loadSchema, runStructured } from "./structured";
import type { CliIO } from "./output";
import { writeError, writeSuccess } from "./output";
import { updateProCli } from "./update";

const HELP_TEXT =
  "pro-cli: ChatGPT Pro CLI\nask: direct blocking query, no job DB\nodds: probability of YES, integer 0-100\nlimits: plan + observed counters\njob create --wait: durable blocking query\njob wait: waits until done\nupdate: fast-forward install\nUse --json for agents.";

export async function runCli(argv: string[], io: CliIO): Promise<number> {
  const parsed = parseArgs(argv);
  const mode = {
    json: flagBoolean(parsed.flags, "json") || (!flagBoolean(parsed.flags, "no-json") && !io.stdoutIsTTY),
  };

  try {
    const [command, subcommand, ...rest] = parsed.positionals;
    if (flagBoolean(parsed.flags, "version") || command === "version") {
      io.stdout(`pro-cli ${packageJson.version}\n`);
      return EXIT.success;
    }

    if (
      !command ||
      flagBoolean(parsed.flags, "help") ||
      command === "help" ||
      command === "--help"
    ) {
      writeSuccess(io, mode, { text: HELP_TEXT, commands: commandList() });
      return EXIT.success;
    }

    if (command === "update") {
      writeSuccess(io, mode, updateProCli());
      return EXIT.success;
    }

    await migrateLegacyDefaultHome(io.env);
    const config = await loadConfig(io.env);
    const paths = resolvePaths(io.env, config);

    switch (command) {
      case "setup": {
        const auth = await getAuthStatus(paths);
        const port = flagString(parsed.flags, "port") ?? "9222";
        writeSuccess(io, mode, buildSetupGuide(auth, paths.home, port));
        return EXIT.success;
      }
      case "auth": {
        if (subcommand === "status") {
          writeSuccess(io, mode, await getAuthStatus(paths));
          return EXIT.success;
        }
        if (subcommand === "command") {
          writeSuccess(io, mode, buildAuthCommand(paths.home, flagString(parsed.flags, "port") ?? "9222"));
          return EXIT.success;
        }
        if (subcommand === "capture") {
          const cdpBase = defaultCdpBase(flagString(parsed.flags, "port"), flagString(parsed.flags, "cdp"));
          const jsonPath = flagString(parsed.flags, "out") ?? paths.cookieJsonPath;
          const jarPath = flagString(parsed.flags, "jar") ?? paths.cookieJarPath;
          const tokenPath = flagString(parsed.flags, "token-out") ?? paths.sessionTokenPath;
          const status = await captureAuth({
            cdpBase,
            jsonPath,
            jarPath,
            tokenPath,
            dryRun: flagBoolean(parsed.flags, "dry-run"),
          });
          await saveConfig(io.env, {
            ...config,
            cookieJsonPath: jsonPath,
            cookieJarPath: jarPath,
            sessionTokenPath: tokenPath,
          });
          writeSuccess(io, mode, status);
          return EXIT.success;
        }
        if (subcommand === "reset") {
          const port = flagString(parsed.flags, "port") ?? "9222";
          const result = await resetAuthProfile({
            home: paths.home,
            port,
            noBackup: flagBoolean(parsed.flags, "no-backup"),
            noLaunch: flagBoolean(parsed.flags, "no-launch"),
            keepBackups: parseIntegerFlag(parsed.flags, "keep-backups", 5, 0, 100),
          });
          writeSuccess(io, mode, result);
          return EXIT.success;
        }
        if (subcommand === "hide") {
          const cdpBase = defaultCdpBase(flagString(parsed.flags, "port"), flagString(parsed.flags, "cdp"));
          const result = await hideProCliChromeWindow(cdpBase);
          writeSuccess(io, mode, result);
          return EXIT.success;
        }
        if (subcommand === "show") {
          const cdpBase = defaultCdpBase(flagString(parsed.flags, "port"), flagString(parsed.flags, "cdp"));
          const result = await showProCliChromeWindow(cdpBase);
          writeSuccess(io, mode, result);
          return EXIT.success;
        }
        throw invalidArgs("Unknown auth command.", [
          "Use pro-cli auth status, pro-cli auth command, pro-cli auth capture, pro-cli auth reset, pro-cli auth hide, or pro-cli auth show.",
        ]);
      }
      case "models": {
        writeSuccess(io, mode, await listModels({ sessionTokenPath: paths.sessionTokenPath }));
        return EXIT.success;
      }
      case "limits": {
        const cdp = flagString(parsed.flags, "cdp");
        const port = flagString(parsed.flags, "port");
        const cdpBase = cdp || port ? defaultCdpBase(port, cdp) : undefined;
        const account = await fetchAccountSummary(cdpBase);
        const store = await JobStore.open(paths.dbPath);
        try {
          const observed = store.latestLimits();
          writeSuccess(io, mode, {
            account,
            observedLimits: observed,
            note:
              observed.length === 0
                ? "No limits observed yet. Make a pro-cli ask/odds/job call; per-feature counters arrive in the stream metadata."
                : "Per-feature counters captured from the most recent stream that included them. Pro chat throttling is adaptive and not exposed here.",
          });
          return EXIT.success;
        } finally {
          store.close();
        }
      }
      case "ask": {
        const prompt = await promptFromArgs([subcommand, ...rest].filter(Boolean), io.cwd);
        const askModel = resolveModel(flagString(parsed.flags, "model") ?? config.defaultModel ?? DEFAULT_MODEL);
        const askReasoning = resolveRequestReasoning(askModel, parsed.flags, config.defaultReasoning);
        const askOptions = await collectRequestOptions(parsed.flags, io.cwd, ASK_REQUEST_FLAGS, "ask");
        applyModelConversationRequirements(askOptions, parsed.flags, askModel);
        const schemaRaw = flagString(parsed.flags, "schema");
        const formatHint = flagString(parsed.flags, "format");
        if (schemaRaw && formatHint) {
          throw invalidArgs("Use --schema or --format, not both.", ["Pick one structured-output flag."]);
        }
        if (schemaRaw || formatHint) {
          const schema = schemaRaw ? await loadSchema(schemaRaw, io.cwd) : undefined;
          const schemaRetries = parseIntegerFlag(parsed.flags, "schema-retries", 1, 0, 5);
          const structured = await runStructured(prompt, {
            schema,
            formatHint,
            retries: schemaRetries,
            runner: async (wrappedPrompt) => {
              const job = buildEphemeralJob({
                prompt: wrappedPrompt,
                model: askModel,
                reasoning: askReasoning,
                options: askOptions,
              });
              const outcome = await executeEphemeralJob(job, paths);
              if (isRecord(outcome.error)) {
                throw new ProError(
                  typeof outcome.error.code === "string" ? outcome.error.code : "STRUCTURED_RUNNER_FAILED",
                  typeof outcome.error.message === "string" ? outcome.error.message : "ChatGPT request failed.",
                  { exitCode: EXIT.upstream },
                );
              }
              return typeof outcome.result === "string" ? outcome.result : "";
            },
          });
          if (flagBoolean(parsed.flags, "json")) {
            writeSuccess(io, { json: true }, {
              parsed: structured.parsed,
              raw: structured.raw,
              attempts: structured.attempts,
            });
          } else {
            io.stdout(`${JSON.stringify(structured.parsed, null, 2)}\n`);
          }
          return EXIT.success;
        }
        const job = buildEphemeralJob({
          prompt,
          model: askModel,
          reasoning: askReasoning,
          options: askOptions,
        });
        const outcome = await executeEphemeralJob(job, paths);
        if (isRecord(outcome.error)) {
          const error = proErrorFromPayload(outcome.error, outcome.exitCode);
          writeError(io, mode, error);
          return error.exitCode;
        }
        writeSuccess(io, mode, outcome);
        return EXIT.success;
      }
      case "odds": {
        const question = await promptFromArgs([subcommand, ...rest].filter(Boolean), io.cwd);
        const samples = parseIntegerFlag(parsed.flags, "samples", 1, 1, 25);
        const parseRetries = parseIntegerFlag(parsed.flags, "parse-retries", 2, 0, 5);
        const aggregate = resolveAggregate(flagString(parsed.flags, "aggregate") ?? "mean");
        const allowFifty = flagBoolean(parsed.flags, "allow-fifty");
        const contextRaw = flagString(parsed.flags, "context");
        const context = contextRaw ? await readMaybeAtFile(contextRaw, io.cwd) : undefined;
        const baseRequestOptions = await collectRequestOptions(
          parsed.flags,
          io.cwd,
          ODDS_REQUEST_FLAGS,
          "odds",
        );
        const oddsModel = resolveModel(flagString(parsed.flags, "model") ?? config.defaultModel ?? DEFAULT_MODEL);
        applyModelConversationRequirements(baseRequestOptions, parsed.flags, oddsModel);
        const result = await runOdds({
          question,
          context,
          model: oddsModel,
          reasoning: resolveRequestReasoning(oddsModel, parsed.flags, config.defaultReasoning),
          samples,
          aggregate,
          allowFifty,
          parseRetries,
          baseRequestOptions,
          paths,
        });
        if (flagBoolean(parsed.flags, "json")) {
          writeSuccess(io, { json: true }, {
            probability: result.probability,
            probabilityRaw: result.probabilityRaw,
            samples: result.samples,
            aggregate: result.aggregate,
            parseFailures: result.parseFailures,
            rejectedFifties: result.rejectedFifties,
            allowFifty: result.allowFifty,
            model: result.model,
            reasoning: result.reasoning,
            jobIds: result.jobIds,
            attempts: result.attempts,
          });
        } else {
          io.stdout(`${result.probability}\n`);
        }
        return EXIT.success;
      }
      case "job": {
        if (subcommand === "create") {
          const waitRequested = flagBoolean(parsed.flags, "wait");
          if (!waitRequested && hasWaitOptionFlags(parsed.flags)) {
            throw invalidArgs("Wait options require --wait.", [
              "Use pro-cli job create @prompt.md --wait --soft-timeout <ms> --json.",
            ]);
          }
          if (waitRequested && flagBoolean(parsed.flags, "no-start")) {
            throw invalidArgs("Cannot wait for a job without starting the daemon.", [
              "Remove --no-start or remove --wait.",
            ]);
          }
          const jobSchemaRaw = flagString(parsed.flags, "schema");
          const jobFormatHint = flagString(parsed.flags, "format");
          if (jobSchemaRaw && jobFormatHint) {
            throw invalidArgs("Use --schema or --format, not both.", ["Pick one structured-output flag."]);
          }
          if ((jobSchemaRaw || jobFormatHint) && !waitRequested) {
            throw invalidArgs("Structured output requires --wait.", [
              "Add --wait so retries can read the result, or run pro-cli ask --schema.",
            ]);
          }
          const userPrompt = await promptFromArgs(rest, io.cwd);
          const jobModel = resolveModel(flagString(parsed.flags, "model") ?? config.defaultModel ?? DEFAULT_MODEL);
          const jobOptions = await collectRequestOptions(parsed.flags, io.cwd, JOB_CREATE_FLAGS, "job create");
          applyModelConversationRequirements(jobOptions, parsed.flags, jobModel);
          const input = {
            prompt: userPrompt,
            model: jobModel,
            reasoning: resolveRequestReasoning(jobModel, parsed.flags, config.defaultReasoning),
            options: jobOptions,
          };
          if (!flagBoolean(parsed.flags, "no-start")) {
            const daemon = await ensureDaemonRunning(paths, io);
            if (waitRequested && (jobSchemaRaw || jobFormatHint)) {
              const schema = jobSchemaRaw ? await loadSchema(jobSchemaRaw, io.cwd) : undefined;
              const schemaRetries = parseIntegerFlag(parsed.flags, "schema-retries", 1, 0, 5);
              const waitOptions = parseWaitOptions(parsed.flags);
              const jobIds: string[] = [];
              const structured = await runStructured(userPrompt, {
                schema,
                formatHint: jobFormatHint,
                retries: schemaRetries,
                runner: async (wrappedPrompt) => {
                  const created = await daemon.client.createJob({ ...input, prompt: wrappedPrompt });
                  const jobId = jobIdFromPayload(created);
                  jobIds.push(jobId);
                  const waited = await daemon.client.wait(
                    jobId,
                    waitOptions.timeoutMs,
                    waitOptions.pollMs,
                    waitOptions.softTimeout,
                  );
                  if (!isRecord(waited.job) || waited.job.status !== "succeeded") {
                    throw new ProError("STRUCTURED_RUNNER_FAILED", "Job did not reach succeeded status.", {
                      exitCode: EXIT.upstream,
                      details: { jobId, waited },
                    });
                  }
                  const fetched = await daemon.client.result(jobId);
                  return typeof fetched.result === "string" ? fetched.result : "";
                },
              });
              writeSuccess(io, mode, {
                parsed: structured.parsed,
                raw: structured.raw,
                attempts: structured.attempts,
                jobIds,
                daemon: { started: daemon.started, status: daemon.status },
              });
              return EXIT.success;
            }
            const created = await daemon.client.createJob(input);
            if (waitRequested) {
              const waitOptions = parseWaitOptions(parsed.flags);
              const jobId = jobIdFromPayload(created);
              const waited = await daemon.client.wait(
                jobId,
                waitOptions.timeoutMs,
                waitOptions.pollMs,
                waitOptions.softTimeout,
              );
              throwIfTerminalJobFailure(waited, jobId);
              writeSuccess(io, mode, {
                ...waited,
                ...(await resultIfSucceeded(daemon.client, jobId, waited)),
                daemon: { started: daemon.started, status: daemon.status },
              });
              return EXIT.success;
            }
            writeSuccess(io, mode, {
              ...created,
              daemon: { started: daemon.started, status: daemon.status },
            });
            return EXIT.success;
          }
          const store = await JobStore.open(paths.dbPath);
          try {
            const job = store.create(input);
            writeSuccess(io, mode, { job, daemon: { started: false } });
            return EXIT.success;
          } finally {
            store.close();
          }
        }
        if (subcommand === "status") {
          const jobId = rest[0];
          if (!jobId) throw invalidArgs("Missing job id.", ["Use pro-cli job status <job-id>."]);
          const store = await JobStore.open(paths.dbPath);
          try {
            writeSuccess(io, mode, { job: redactJob(store.get(jobId)) });
            return EXIT.success;
          } finally {
            store.close();
          }
        }
        if (subcommand === "result") {
          const jobId = rest[0];
          if (!jobId) throw invalidArgs("Missing job id.", ["Use pro-cli job result <job-id>."]);
          const store = await JobStore.open(paths.dbPath);
          try {
            const job = store.get(jobId);
            if (job.status !== "succeeded") {
              throw new ProError("JOB_NOT_READY", `Job ${jobId} is ${job.status}.`, {
                exitCode: EXIT.notFound,
                suggestions: ["Run pro-cli job wait <job-id> or pro-cli job status <job-id>."],
              });
            }
            writeSuccess(io, mode, { jobId, result: job.result });
            return EXIT.success;
          } finally {
            store.close();
          }
        }
        if (subcommand === "wait") {
          const jobId = rest[0];
          if (!jobId) throw invalidArgs("Missing job id.", ["Use pro-cli job wait <job-id>."]);
          const waitOptions = parseWaitOptions(parsed.flags);
          const daemon = await ensureDaemonRunning(paths, io);
          const waited = await daemon.client.wait(
            jobId,
            waitOptions.timeoutMs,
            waitOptions.pollMs,
            waitOptions.softTimeout,
          );
          throwIfTerminalJobFailure(waited, jobId);
          writeSuccess(io, mode, waited);
          return EXIT.success;
        }
        if (subcommand === "cancel") {
          const jobId = rest[0];
          if (!jobId) throw invalidArgs("Missing job id.", ["Use pro-cli job cancel <job-id>."]);
          const daemon = await ensureDaemonRunning(paths, io);
          writeSuccess(io, mode, await daemon.client.cancel(jobId));
          return EXIT.success;
        }
        if (subcommand === "list") {
          const limit = Number(flagString(parsed.flags, "limit") ?? "20");
          if (!Number.isInteger(limit) || limit < 1 || limit > 200) {
            throw invalidArgs("Invalid --limit.", ["Use --limit between 1 and 200."]);
          }
          const store = await JobStore.open(paths.dbPath);
          try {
            writeSuccess(io, mode, { jobs: store.list(limit) });
            return EXIT.success;
          } finally {
            store.close();
          }
        }
        throw invalidArgs("Unknown job command.", [
          "Use pro-cli job create, job status, job wait, job result, job cancel, or job list.",
        ]);
      }
      case "daemon": {
        if (subcommand === "serve") {
          await runDaemonServer(paths, {
            port: parseOptionalIntegerFlag(parsed.flags, "port", 0, 65_535),
            pollMs: parseIntegerFlag(parsed.flags, "poll-ms", 500, 25, 60_000),
            idleTimeoutMs: parseOptionalIntegerFlag(parsed.flags, "idle-timeout", 1, 24 * 60 * 60_000),
          });
          return EXIT.success;
        }
        if (subcommand === "start" || subcommand === "restart") {
          if (subcommand === "restart") await stopDaemon(paths);
          const daemon = await ensureDaemonRunning(paths, io);
          writeSuccess(io, mode, { daemon: { started: daemon.started, status: daemon.status } });
          return EXIT.success;
        }
        if (subcommand === "status") {
          writeSuccess(io, mode, { daemon: await getDaemonStatus(paths) });
          return EXIT.success;
        }
        if (subcommand === "stop") {
          writeSuccess(io, mode, { daemon: await stopDaemon(paths) });
          return EXIT.success;
        }
        throw invalidArgs("Unknown daemon command.", [
          "Use pro-cli daemon start, pro-cli daemon status, pro-cli daemon stop, or pro-cli daemon restart.",
        ]);
      }
      case "config": {
        if (subcommand === "get") {
          writeSuccess(io, mode, {
            config,
            defaults: { model: DEFAULT_MODEL, reasoning: DEFAULT_REASONING },
            paths: redactPaths(paths),
          });
          return EXIT.success;
        }
        if (subcommand === "set") {
          const [key, value] = rest;
          if (!key || !value) throw invalidArgs("Missing config key/value.", ["Use pro-cli config set model gpt-5-5-pro."]);
          const next = { ...config };
          if (key === "model") next.defaultModel = resolveModel(value);
          else if (key === "reasoning") next.defaultReasoning = resolveReasoning(value);
          else throw invalidArgs(`Unknown config key ${key}.`, ["Supported keys: model, reasoning."]);
          await saveConfig(io.env, next);
          writeSuccess(io, mode, { config: next });
          return EXIT.success;
        }
        throw invalidArgs("Unknown config command.", ["Use pro-cli config get or pro-cli config set."]);
      }
      case "doctor": {
        const auth = await getAuthStatus(paths);
        const cdpBase = defaultCdpBase(flagString(parsed.flags, "port"), flagString(parsed.flags, "cdp"));
        const browserSession = await getBrowserSessionStatus(
          cdpBase,
          parseIntegerFlag(parsed.flags, "timeout", 3_000, 1, 60_000),
        );
        const authReady = auth.tokenStatus === "present" && auth.accountIdPresent;
        const ready = authReady && browserSession.status === "present";
        const cdpPort = new URL(cdpBase).port || "9222";
        const profileDir = join(paths.home, "chrome-profile");
        writeSuccess(io, mode, {
          auth,
          browserSession,
          daemon: await getDaemonStatus(paths),
          ready,
          next: buildDoctorNext(authReady, browserSession.status, cdpBase),
          portCollision: detectPortCollision(cdpPort, profileDir),
          legacyArtifacts: await detectLegacyArtifacts(),
          storage: { home: paths.home, dbPath: paths.dbPath },
          transport: {
            status: ready ? "configured" : "auth_required",
            endpoint: "https://chatgpt.com/backend-api/f/conversation",
          },
          safety: safetySummary(),
        });
        return EXIT.success;
      }
      default:
        throw invalidArgs(`Unknown command ${command}.`, ["Run pro-cli help."]);
    }
  } catch (error) {
    const proError = toProError(error);
    writeError(io, mode, proError);
    return proError.exitCode;
  }
}

function invalidArgs(message: string, suggestions: string[]): ProError {
  return new ProError("INVALID_ARGS", message, { exitCode: EXIT.invalidArgs, suggestions });
}

function proErrorFromPayload(error: Record<string, unknown>, exitCode: unknown): ProError {
  return new ProError(
    typeof error.code === "string" ? error.code : "UPSTREAM_ERROR",
    typeof error.message === "string" ? error.message : "ChatGPT request failed.",
    {
      exitCode: isExitCode(exitCode) ? exitCode : EXIT.upstream,
      suggestions: Array.isArray(error.suggestions)
        ? error.suggestions.filter((item): item is string => typeof item === "string")
        : [],
      details: isRecord(error.details) ? error.details : undefined,
    },
  );
}

function isExitCode(value: unknown): value is (typeof EXIT)[keyof typeof EXIT] {
  return typeof value === "number" && Object.values(EXIT).includes(value as (typeof EXIT)[keyof typeof EXIT]);
}

function commandList(): string[] {
  return [
    "setup",
    "update",
    "auth command",
    "auth status",
    "auth capture",
    "auth reset",
    "auth hide",
    "auth show",
    "models",
    "ask",
    "odds",
    "limits",
    "job create",
    "job status",
    "job wait",
    "job result",
    "job cancel",
    "job list",
    "daemon start",
    "daemon status",
    "daemon stop",
    "config get",
    "doctor",
  ];
}

function buildDoctorNext(
  authReady: boolean,
  browserStatus: Awaited<ReturnType<typeof getBrowserSessionStatus>>["status"],
  cdpBase: string,
): Record<string, string> {
  if (authReady && browserStatus === "present") {
    return {
      command: `pro-cli ask "<your prompt>" --cdp ${cdpBase} --json`,
      reason: "Stored auth and the live CDP ChatGPT page are ready; send the real request directly.",
    };
  }
  if (authReady && browserStatus === "logged_out") {
    return {
      command: `pro-cli auth capture --cdp ${cdpBase} --json`,
      reason: "The CDP ChatGPT page is reachable but logged out; sign in there, then recapture auth.",
    };
  }
  if (authReady && browserStatus === "probe_failed") {
    return {
      command: `pro-cli auth capture --cdp ${cdpBase} --json`,
      reason:
        "The CDP ChatGPT auth probe returned an unexpected HTTP status (often cookie bloat causing 431); sign out and back in, then recapture auth.",
    };
  }
  if (authReady && (browserStatus === "page_missing" || browserStatus === "cdp_unavailable")) {
    return {
      command: "pro-cli auth command --json",
      reason: "Stored auth exists, but no live ChatGPT CDP page is available.",
    };
  }
  return {
    command: "pro-cli setup --json",
    reason: "Auth is missing or expired; follow the setup steps.",
  };
}

function buildSetupGuide(auth: Awaited<ReturnType<typeof getAuthStatus>>, home: string, port: string): Record<string, unknown> {
  const ready = auth.tokenStatus === "present" && auth.accountIdPresent;
  const authCommand = buildAuthCommand(home, port);
  return {
    ready,
    summary: ready ? "pro-cli is ready to query ChatGPT." : "pro-cli needs a logged-in ChatGPT browser session.",
    steps: [
      {
        id: "install",
        status: "done",
        command: "curl -fsSL https://raw.githubusercontent.com/ratacat/pro-cli/main/scripts/install.sh | bash",
        note: "Clones or updates ~/Projects/pro-cli and links the pro-cli binary.",
      },
      {
        id: "open-chatgpt",
        status: ready ? "done" : "todo",
        command: authCommand.command,
        note: "Starts the dedicated ~/.pro-cli Chrome profile with CDP enabled; keep this window open while pro-cli jobs run.",
      },
      {
        id: "capture-auth",
        status: ready ? "done" : "todo",
        command: authCommand.captureCommand,
        note: "Captures scoped cookies plus the page session token into private local files.",
      },
      {
        id: "doctor",
        status: ready ? "todo" : "blocked",
        command: `pro-cli doctor --cdp ${authCommand.cdp} --json`,
        note: "Checks stored auth and the live CDP ChatGPT page without spending Pro quota.",
      },
    ],
    auth,
    storage: {
      home,
      cookieJsonPath: auth.cookieJsonPath,
      sessionTokenPath: auth.sessionTokenPath,
    },
    safety: safetySummary(),
  };
}

function buildAuthCommand(home: string, port: string): Record<string, unknown> {
  const profileDir = join(home, "chrome-profile");
  const url = "https://chatgpt.com/";
  const cdp = `http://127.0.0.1:${port}`;
  const command =
    process.platform === "darwin"
      ? `open -na "Google Chrome" --args --user-data-dir=${shellQuote(profileDir)} --remote-debugging-port=${port} ${url}`
      : process.platform === "win32"
        ? `start "" chrome.exe --user-data-dir=${windowsQuote(profileDir)} --remote-debugging-port=${port} ${url}`
        : `google-chrome --user-data-dir=${shellQuote(profileDir)} --remote-debugging-port=${port} ${url}`;
  return {
    command,
    captureCommand: `pro-cli auth capture --cdp ${cdp} --json`,
    cdp,
    profileDir,
    port,
    portCollision: detectPortCollision(port, profileDir),
    safety: "Recommended profile is dedicated to pro-cli; keep it open for jobs and do not expose a normal browser profile over CDP.",
  };
}

function safetySummary(): Record<string, unknown> {
  return {
    rawValuesPrinted: false,
    storedLocally: true,
    fileModes: "0600 files, 0700 directories where supported",
    sentTo: ["https://chatgpt.com"],
    reminder: "Cookie and token files are sensitive; do not commit, paste, or share ~/.pro-cli.",
  };
}

async function promptFromArgs(args: string[], cwd: string): Promise<string> {
  const prompt = args.join(" ").trim();
  if (!prompt) throw invalidArgs("Missing prompt.", ["Use pro-cli ask \"prompt\" or pro-cli job create @prompt.md."]);
  if (prompt.startsWith("@") && !prompt.includes(" ")) {
    return readFile(new URL(prompt.slice(1), `file://${cwd}/`), "utf8");
  }
  return prompt;
}

async function collectRequestOptions(
  flags: Map<string, string | boolean | string[]>,
  cwd: string,
  allowedFlags: Set<string>,
  command: string,
): Promise<Record<string, unknown>> {
  rejectUnsupportedFlags(flags, allowedFlags, command);
  const options: Record<string, unknown> = {};
  setStringOption(options, "verbosity", flags, "verbosity", ["low", "medium", "high"]);
  setStringOption(options, "reasoningSummary", flags, "reasoning-summary", [
    "auto",
    "concise",
    "detailed",
    "none",
  ]);
  setStringOption(options, "toolChoice", flags, "tool-choice", ["auto", "none", "required"]);
  setIntegerOption(options, "timeoutMs", flags, "timeout", 1, 30 * 60_000);
  setIntegerOption(options, "retries", flags, "retries", 0, 5);
  setIntegerOption(options, "retryDelayMs", flags, "retry-delay", 0, 60_000);
  setIntegerAliasOption(
    options,
    "condensedResponseTokens",
    flags,
    "condensed-response",
    ["condensed_response"],
    1,
    100_000,
  );
  setBooleanOption(options, "parallelTools", flags, "parallel-tools");
  setConversationOptions(options, flags);
  const cdp = flagString(flags, "cdp");
  const port = flagString(flags, "port");
  if (cdp || port) options.cdpBase = defaultCdpBase(port, cdp);

  const instructions = flagString(flags, "instructions");
  const instructionsFile = flagString(flags, "instructions-file");
  if (instructions && instructionsFile) {
    throw invalidArgs("Use only one instructions source.", ["Pass --instructions or --instructions-file, not both."]);
  }
  if (instructions) {
    options.instructions = await readMaybeAtFile(instructions, cwd);
  }
  if (instructionsFile) {
    options.instructions = await readFile(new URL(instructionsFile, `file://${cwd}/`), "utf8");
  }
  return options;
}

const ASK_REQUEST_FLAGS = new Set([
  "json",
  "no-json",
  "model",
  "reasoning",
  "verbosity",
  "reasoning-summary",
  "tool-choice",
  "parallel-tools",
  "instructions",
  "instructions-file",
  "timeout",
  "retries",
  "retry-delay",
  "condensed-response",
  "condensed_response",
  "store",
  "save",
  "temporary",
  "no-temporary",
  "conversation",
  "parent",
  "cdp",
  "port",
  "schema",
  "format",
  "schema-retries",
]);

const JOB_CREATE_FLAGS = new Set([
  ...ASK_REQUEST_FLAGS,
  "no-start",
  "wait",
  "wait-timeout",
  "soft-timeout",
  "poll-ms",
]);

const ODDS_REQUEST_FLAGS = new Set([
  "json",
  "no-json",
  "model",
  "reasoning",
  "verbosity",
  "reasoning-summary",
  "tool-choice",
  "parallel-tools",
  "timeout",
  "retries",
  "retry-delay",
  "store",
  "save",
  "temporary",
  "no-temporary",
  "conversation",
  "parent",
  "cdp",
  "port",
  "context",
  "samples",
  "aggregate",
  "parse-retries",
  "allow-fifty",
]);

function resolveAggregate(value: string): AggregateMethod {
  if (value === "mean" || value === "median" || value === "trimmed-mean") return value;
  throw invalidArgs("Invalid --aggregate.", ["Allowed values: mean, median, trimmed-mean."]);
}

function setConversationOptions(
  options: Record<string, unknown>,
  flags: Map<string, string | boolean | string[]>,
): void {
  const conversationId = flagString(flags, "conversation");
  const parentMessageId = flagString(flags, "parent");
  if (conversationId || parentMessageId) {
    if (!conversationId || !parentMessageId) {
      throw invalidArgs("Continuing a conversation needs both ids.", [
        "Use --conversation <conversation-id> --parent <message-id>.",
      ]);
    }
    options.conversationId = conversationId;
    options.parentMessageId = parentMessageId;
  }

  const store = readBooleanFlag(flags, "store");
  const save = flagBoolean(flags, "save") || flagBoolean(flags, "no-temporary") || store === true;
  const temporary = flagBoolean(flags, "temporary") || store === false;
  if (save && temporary) {
    throw invalidArgs("Choose either temporary or saved chat mode.", [
      "Use --temporary for a temporary chat or --save for a saved conversation.",
    ]);
  }

  options.temporary = temporary || (!save && !conversationId);
}

function resolveReasoning(reasoning: string): string {
  if (isReasoningLevel(reasoning)) return reasoning;
  throw invalidArgs("Invalid --reasoning.", [`Allowed values: ${REASONING_LEVELS.join(", ")}.`]);
}

function resolveRequestReasoning(
  model: string,
  flags: Map<string, string | boolean | string[]>,
  configuredReasoning?: string,
): string {
  const explicitReasoning = flagString(flags, "reasoning");
  if (!modelUsesThinkingEffort(model)) {
    if (explicitReasoning !== undefined) {
      throw invalidArgs(`Model ${model} does not support --reasoning.`, [
        `Omit --reasoning for ${model}; ChatGPT's model catalog exposes no thinking_effort levels for it.`,
      ]);
    }
    return NO_REASONING;
  }
  return resolveReasoning(explicitReasoning ?? configuredReasoning ?? DEFAULT_REASONING);
}

function applyModelConversationRequirements(
  options: Record<string, unknown>,
  flags: Map<string, string | boolean | string[]>,
  model: string,
): void {
  if (!modelRequiresSavedConversation(model)) return;
  const store = readBooleanFlag(flags, "store");
  const requestedTemporary = flagBoolean(flags, "temporary") || store === false;
  if (requestedTemporary) {
    throw invalidArgs("Deep Research does not support temporary chats.", [
      "Remove --temporary or --store false; pro-cli uses a saved ChatGPT conversation for --model research.",
    ]);
  }
  if (!options.conversationId) options.temporary = false;
}

function resolveModel(model: string): string {
  const value = canonicalModelId(model);
  if (!value || value === "auto") {
    throw invalidArgs("Invalid --model.", [
      "Use a concrete model id such as gpt-5-5-pro, gpt-4-5, or research; or run pro-cli models --json.",
    ]);
  }
  return value;
}

function rejectUnsupportedFlags(
  flags: Map<string, string | boolean | string[]>,
  allowed: Set<string>,
  command: string,
): void {
  for (const key of flags.keys()) {
    if (!allowed.has(key)) {
      throw invalidArgs(`Unsupported --${key} for ${command}.`, [
        "Run pro-cli help or pro-cli models --json for supported request controls.",
      ]);
    }
  }
}

async function readMaybeAtFile(value: string, cwd: string): Promise<string> {
  if (value.startsWith("@") && !value.includes(" ")) {
    return readFile(new URL(value.slice(1), `file://${cwd}/`), "utf8");
  }
  return value;
}

function setStringOption(
  options: Record<string, unknown>,
  target: string,
  flags: Map<string, string | boolean | string[]>,
  source: string,
  allowed?: string[],
): void {
  const value = flagString(flags, source);
  if (value === undefined) return;
  if (allowed && !allowed.includes(value)) {
    throw invalidArgs(`Invalid --${source}.`, [`Allowed values: ${allowed.join(", ")}.`]);
  }
  options[target] = value;
}

function setIntegerOption(
  options: Record<string, unknown>,
  target: string,
  flags: Map<string, string | boolean | string[]>,
  source: string,
  min: number,
  max: number,
): void {
  const value = flagString(flags, source);
  if (value === undefined) return;
  const number = Number(value);
  if (!Number.isInteger(number) || number < min || number > max) {
    throw invalidArgs(`Invalid --${source}.`, [`Use an integer between ${min} and ${max}.`]);
  }
  options[target] = number;
}

function setIntegerAliasOption(
  options: Record<string, unknown>,
  target: string,
  flags: Map<string, string | boolean | string[]>,
  primary: string,
  aliases: string[],
  min: number,
  max: number,
): void {
  const names = [primary, ...aliases];
  const present = names
    .map((name) => ({ name, value: flagString(flags, name) }))
    .filter((entry): entry is { name: string; value: string } => entry.value !== undefined);
  if (present.length === 0) return;
  const first = present[0];
  for (const entry of present.slice(1)) {
    if (entry.value !== first.value) {
      throw invalidArgs("Use only one condensed response flag.", [
        `Pass --${primary} <tokens> or --${aliases[0]}=<tokens>, not both.`,
      ]);
    }
  }
  const number = Number(first.value);
  if (!Number.isInteger(number) || number < min || number > max) {
    throw invalidArgs(`Invalid --${first.name}.`, [`Use an integer between ${min} and ${max}.`]);
  }
  options[target] = number;
}

function setBooleanOption(
  options: Record<string, unknown>,
  target: string,
  flags: Map<string, string | boolean | string[]>,
  source: string,
): void {
  const parsed = readBooleanFlag(flags, source);
  if (parsed === undefined) return;
  options[target] = parsed;
}

function readBooleanFlag(
  flags: Map<string, string | boolean | string[]>,
  source: string,
): boolean | undefined {
  const value = flagString(flags, source);
  if (value === undefined) return undefined;
  if (["true", "1", "yes", "on"].includes(value.toLowerCase())) {
    return true;
  }
  if (["false", "0", "no", "off"].includes(value.toLowerCase())) {
    return false;
  }
  throw invalidArgs(`Invalid --${source}.`, ["Use true or false."]);
}

interface WaitOptions {
  timeoutMs: number;
  pollMs: number;
  softTimeout: boolean;
}

function parseWaitOptions(flags: Map<string, string | boolean | string[]>): WaitOptions {
  const waitTimeout = flagString(flags, "wait-timeout");
  const softTimeout = flagString(flags, "soft-timeout");
  if (waitTimeout !== undefined && softTimeout !== undefined) {
    throw invalidArgs("Choose one wait timeout mode.", [
      "Use --wait-timeout for an error on timeout or --soft-timeout for ok:true polling.",
    ]);
  }
  return {
    timeoutMs:
      softTimeout !== undefined
        ? parseIntegerFlag(flags, "soft-timeout", 0, 1, 24 * 60 * 60_000)
        : parseIntegerFlag(flags, "wait-timeout", 0, 0, 24 * 60 * 60_000),
    pollMs: parseIntegerFlag(flags, "poll-ms", 500, 25, 60_000),
    softTimeout: softTimeout !== undefined,
  };
}

function hasWaitOptionFlags(flags: Map<string, string | boolean | string[]>): boolean {
  return (
    flagString(flags, "wait-timeout") !== undefined ||
    flagString(flags, "soft-timeout") !== undefined ||
    flagString(flags, "poll-ms") !== undefined
  );
}

function jobIdFromPayload(payload: Record<string, unknown>): string {
  const job = payload.job;
  if (isRecord(job) && typeof job.id === "string") return job.id;
  throw new ProError("DAEMON_BAD_RESPONSE", "Daemon create response did not include a job id.", {
    exitCode: EXIT.internal,
    suggestions: ["Run pro-cli daemon restart --json and retry."],
  });
}

function throwIfTerminalJobFailure(payload: Record<string, unknown>, fallbackJobId: string): void {
  const job = payload.job;
  if (!isRecord(job)) return;
  const status = typeof job.status === "string" ? job.status : "";
  if (status !== "failed" && status !== "cancelled") return;

  const parsed = parseJobErrorPayload(job.error);
  const code = stringValue(parsed.code) ?? (status === "cancelled" ? "JOB_CANCELLED" : "JOB_FAILED");
  const jobId = stringValue(job.id) ?? fallbackJobId;
  const message = stringValue(parsed.message) ?? `Job ${jobId} ${status}.`;
  const details = isRecord(parsed.details) ? parsed.details : {};
  throw new ProError(code, message, {
    exitCode: exitCodeForJobFailure(code, status),
    suggestions: stringArray(parsed.suggestions),
    details: {
      ...details,
      jobId,
      jobStatus: status,
      job,
      waited: payload,
    },
  });
}

function parseJobErrorPayload(raw: unknown): Record<string, unknown> {
  if (isRecord(raw)) return raw;
  if (typeof raw !== "string" || raw.trim().length === 0) return {};
  try {
    const parsed = JSON.parse(raw) as unknown;
    return isRecord(parsed) ? parsed : { message: raw };
  } catch {
    return { message: raw };
  }
}

function exitCodeForJobFailure(code: string, status: string): (typeof EXIT)[keyof typeof EXIT] {
  if (status === "cancelled") return EXIT.upstream;
  if (
    [
      "ACCOUNT_ID_MISSING",
      "CDP_EVALUATION_FAILED",
      "CDP_UNAVAILABLE",
      "CHATGPT_PAGE_LOGGED_OUT",
      "CHATGPT_PAGE_MISSING",
      "CHATGPT_PROBE_FAILED",
      "SESSION_TOKEN_EXPIRED",
      "SESSION_TOKEN_MISSING",
    ].includes(code)
  ) {
    return EXIT.auth;
  }
  if (code === "NETWORK_ERROR") return EXIT.network;
  if (code === "WAIT_TIMEOUT") return EXIT.timeout;
  if (code === "INVALID_ARGS" || code === "INVALID_JSON") return EXIT.invalidArgs;
  if (code === "JOB_NOT_FOUND" || code === "JOB_NOT_READY") return EXIT.notFound;
  if (code === "INTERNAL_ERROR" || code === "DAEMON_BAD_RESPONSE") return EXIT.internal;
  return EXIT.upstream;
}

function stringValue(value: unknown): string | undefined {
  return typeof value === "string" && value.length > 0 ? value : undefined;
}

function stringArray(value: unknown): string[] {
  return Array.isArray(value) ? value.filter((item): item is string => typeof item === "string") : [];
}

async function resultIfSucceeded(
  client: { result: (jobId: string) => Promise<Record<string, unknown>> },
  jobId: string,
  payload: Record<string, unknown>,
): Promise<Record<string, unknown>> {
  const job = payload.job;
  if (!isRecord(job) || job.status !== "succeeded") return {};
  const result = await client.result(jobId);
  return typeof result.result === "string" ? { result: result.result } : {};
}

function parseIntegerFlag(
  flags: Map<string, string | boolean | string[]>,
  source: string,
  fallback: number,
  min: number,
  max: number,
): number {
  const value = flagString(flags, source);
  if (value === undefined) return fallback;
  const number = Number(value);
  if (!Number.isInteger(number) || number < min || number > max) {
    throw invalidArgs(`Invalid --${source}.`, [`Use an integer between ${min} and ${max}.`]);
  }
  return number;
}

function parseOptionalIntegerFlag(
  flags: Map<string, string | boolean | string[]>,
  source: string,
  min: number,
  max: number,
): number | undefined {
  const value = flagString(flags, source);
  if (value === undefined) return undefined;
  const number = Number(value);
  if (!Number.isInteger(number) || number < min || number > max) {
    throw invalidArgs(`Invalid --${source}.`, [`Use an integer between ${min} and ${max}.`]);
  }
  return number;
}

function isRecord(value: unknown): value is Record<string, unknown> {
  return Boolean(value) && typeof value === "object" && !Array.isArray(value);
}

function redactPaths(paths: { home: string; configPath: string; dbPath: string }): {
  home: string;
  configPath: string;
  dbPath: string;
} {
  return {
    home: paths.home,
    configPath: paths.configPath,
    dbPath: paths.dbPath,
  };
}

function shellQuote(value: string): string {
  return `'${value.replace(/'/g, "'\\''")}'`;
}

function windowsQuote(value: string): string {
  return `"${value.replace(/"/g, '""')}"`;
}
`````

## File: src/args.ts
`````typescript
import { EXIT, ProError } from "./errors";

export interface ParsedArgs {
  positionals: string[];
  flags: Map<string, string | boolean | string[]>;
}

const BOOLEAN_FLAGS = new Set([
  "json",
  "no-json",
  "dry-run",
  "include-expired",
  "no-start",
  "save",
  "temporary",
  "no-temporary",
  "wait",
  "help",
  "version",
  "allow-fifty",
  "no-launch",
  "no-backup",
]);

export function parseArgs(argv: string[]): ParsedArgs {
  const positionals: string[] = [];
  const flags = new Map<string, string | boolean | string[]>();
  let flagsEnded = false;

  for (let index = 0; index < argv.length; index += 1) {
    const token = argv[index];
    if (flagsEnded) {
      positionals.push(token);
      continue;
    }
    if (token === "--") {
      flagsEnded = true;
      continue;
    }
    if (!token.startsWith("--")) {
      positionals.push(token);
      continue;
    }

    const raw = token.slice(2);
    if (!raw) {
      throw new ProError("INVALID_ARGS", "Empty flag is not valid.", {
        exitCode: EXIT.invalidArgs,
        suggestions: ["Use flags like --json or --model gpt-5-5-pro."],
      });
    }

    const [name, inlineValue] = raw.split(/=(.*)/s, 2);
    if (!name) {
      throw new ProError("INVALID_ARGS", `Invalid flag ${token}.`, {
        exitCode: EXIT.invalidArgs,
        suggestions: ["Use flags like --json or --model gpt-5-5-pro."],
      });
    }

    let value: string | boolean;
    if (inlineValue !== undefined) {
      value = inlineValue;
    } else if (BOOLEAN_FLAGS.has(name)) {
      value = true;
    } else {
      const next = argv[index + 1];
      if (!next || next.startsWith("--")) {
        throw new ProError("INVALID_ARGS", `Missing value for --${name}.`, {
          exitCode: EXIT.invalidArgs,
          suggestions: [`Pass --${name} <value>.`],
        });
      }
      value = next;
      index += 1;
    }

    const existing = flags.get(name);
    if (existing === true && value === true && BOOLEAN_FLAGS.has(name)) {
      flags.set(name, true);
    } else if (Array.isArray(existing)) {
      existing.push(String(value));
    } else if (existing !== undefined) {
      flags.set(name, [String(existing), String(value)]);
    } else {
      flags.set(name, value);
    }
  }

  return { positionals, flags };
}

export function flagString(
  flags: Map<string, string | boolean | string[]>,
  name: string,
): string | undefined {
  const value = flags.get(name);
  if (Array.isArray(value)) {
    throw new ProError("INVALID_ARGS", `Repeated --${name} is not valid here.`, {
      exitCode: EXIT.invalidArgs,
      suggestions: [`Pass --${name} only once.`],
    });
  }
  if (value === undefined || value === false) return undefined;
  if (value === true) return "true";
  return value;
}

export function flagBoolean(
  flags: Map<string, string | boolean | string[]>,
  name: string,
): boolean {
  return flags.get(name) === true;
}

export function flagStrings(
  flags: Map<string, string | boolean | string[]>,
  name: string,
): string[] {
  const value = flags.get(name);
  if (value === undefined || value === false) return [];
  if (Array.isArray(value)) return value;
  if (value === true) return ["true"];
  return [value];
}
`````

## File: src/auth.ts
`````typescript
import { spawn, spawnSync } from "node:child_process";
import { access, readdir, rename, rm, stat } from "node:fs/promises";
import { homedir } from "node:os";
import { join } from "node:path";
import { EXIT, ProError } from "./errors";
import {
  callBrowserCdp,
  evaluateInCdpPage,
  findChatGptTargetId,
  getCookiesFromCdp,
  recoverCookieBloatInCdp,
  pruneVolatileCookiesFromCdp,
} from "./cdp";
import {
  chatGptOrigins,
  cookieSummary,
  loadCookieExport,
  toCookieExport,
  toNetscapeCookieJar,
} from "./cookies";
import type { RuntimePaths } from "./config";
import { writePrivateFile } from "./config";
import { isTokenFresh, loadSessionToken, toSessionTokenExport } from "./session-token";

export interface AuthStatus {
  status: "missing" | "present";
  cookieJsonPath: string;
  cookieJarPath: string;
  sessionTokenPath: string;
  tokenStatus: "missing" | "present" | "expired";
  tokenExpiresAt?: string;
  accountIdPresent: boolean;
  cookieCount: number;
  expiredCookieCount: number;
  domains: string[];
  names: string[];
  rawValuesPrinted: false;
}

export type BrowserSessionState =
  | "present"
  | "logged_out"
  | "probe_failed"
  | "page_missing"
  | "cdp_unavailable";

export interface BrowserSessionStatus {
  status: BrowserSessionState;
  cdpBase: string;
  httpStatus?: number;
  pageOrigin?: string;
  errorCode?: string;
  message?: string;
  suggestions: string[];
  rawValuesPrinted: false;
}

export async function getAuthStatus(paths: RuntimePaths): Promise<AuthStatus> {
  try {
    await access(paths.cookieJsonPath);
  } catch {
    return {
      status: "missing",
      cookieJsonPath: paths.cookieJsonPath,
      cookieJarPath: paths.cookieJarPath,
      sessionTokenPath: paths.sessionTokenPath,
      tokenStatus: await readTokenStatus(paths.sessionTokenPath),
      accountIdPresent: await readAccountIdPresent(paths.sessionTokenPath),
      cookieCount: 0,
      expiredCookieCount: 0,
      domains: [],
      names: [],
      rawValuesPrinted: false,
    };
  }

  const cookieExport = await loadCookieExport(paths.cookieJsonPath);
  const summary = cookieSummary(cookieExport.cookies);
  return {
    status: "present",
    cookieJsonPath: paths.cookieJsonPath,
    cookieJarPath: paths.cookieJarPath,
    sessionTokenPath: paths.sessionTokenPath,
    ...(await tokenStatusFields(paths.sessionTokenPath)),
    cookieCount: summary.count,
    expiredCookieCount: summary.expired,
    domains: summary.domains,
    names: summary.names,
    rawValuesPrinted: false,
  };
}

export async function getBrowserSessionStatus(
  cdpBase: string,
  timeoutMs = 3_000,
): Promise<BrowserSessionStatus> {
  try {
    let result = await evaluateBrowserSession(cdpBase, timeoutMs);
    if (shouldRecoverCookieBloat(result)) {
      const recovered = await recoverCookieBloatInCdp(cdpBase, chatGptOrigins(), timeoutMs).catch(() => null);
      if (recovered?.deleted) {
        result = await evaluateBrowserSession(cdpBase, timeoutMs);
      }
    }

    if (result?.code === "CHATGPT_PAGE_MISSING") {
      return {
        status: "page_missing",
        cdpBase,
        httpStatus: result.status,
        pageOrigin: result.href,
        suggestions: [
          "Open the Chrome command from pro-cli auth command.",
          "Confirm the CDP tab is on https://chatgpt.com/.",
        ],
        rawValuesPrinted: false,
      };
    }

    if (result?.hasAccessToken) {
      return {
        status: "present",
        cdpBase,
        httpStatus: result.status,
        pageOrigin: result.origin,
        suggestions: [],
        rawValuesPrinted: false,
      };
    }

    const probeStatus = typeof result?.status === "number" ? result.status : 0;
    const isLoggedOutSignal = probeStatus === 200 || probeStatus === 401;
    if (probeStatus !== 0 && !isLoggedOutSignal) {
      return {
        status: "probe_failed",
        cdpBase,
        httpStatus: probeStatus,
        pageOrigin: result?.origin,
        errorCode: "CHATGPT_PROBE_FAILED",
        message: `ChatGPT auth session probe returned HTTP ${probeStatus}.`,
        suggestions: probeFailedSuggestions(probeStatus, cdpBase),
        rawValuesPrinted: false,
      };
    }

    return {
      status: "logged_out",
      cdpBase,
      httpStatus: result?.status,
      pageOrigin: result?.origin,
      suggestions: [
        "Sign in to ChatGPT in the CDP Chrome window.",
        `Run pro-cli auth capture --cdp ${cdpBase} --json after login.`,
      ],
      rawValuesPrinted: false,
    };
  } catch (error) {
    const proError = error instanceof ProError ? error : null;
    if (proError?.code === "CHATGPT_PAGE_MISSING") {
      return {
        status: "page_missing",
        cdpBase,
        errorCode: proError.code,
        message: proError.message,
        suggestions: proError.suggestions,
        rawValuesPrinted: false,
      };
    }
    return {
      status: "cdp_unavailable",
      cdpBase,
      errorCode: proError?.code ?? "CDP_UNAVAILABLE",
      message: error instanceof Error ? error.message : String(error),
      suggestions:
        proError?.suggestions.length ? proError.suggestions : ["Open Chrome with remote debugging enabled."],
      rawValuesPrinted: false,
    };
  }
}

async function evaluateBrowserSession(
  cdpBase: string,
  timeoutMs: number,
): Promise<{
  status: number;
  hasAccessToken: boolean;
  origin: string;
  href: string;
  code?: "CHATGPT_PAGE_MISSING";
}> {
  return evaluateInCdpPage<{
    status: number;
    hasAccessToken: boolean;
    origin: string;
    href: string;
    code?: "CHATGPT_PAGE_MISSING";
  }>(
    cdpBase,
    `(async () => {
      if (location.origin !== "https://chatgpt.com") {
        return {
          status: 0,
          hasAccessToken: false,
          origin: location.origin,
          href: location.href,
          code: "CHATGPT_PAGE_MISSING"
        };
      }
      const res = await fetch("https://chatgpt.com/api/auth/session", { credentials: "include", referrerPolicy: "no-referrer" });
      const json = await res.json().catch(() => null);
      return {
        status: res.status,
        hasAccessToken: typeof json?.accessToken === "string" && json.accessToken.length > 0,
        origin: location.origin,
        href: location.href
      };
    })()`,
    timeoutMs,
  );
}

function shouldRecoverCookieBloat(result: {
  status: number;
  href?: string;
  code?: "CHATGPT_PAGE_MISSING";
}): boolean {
  return result.status === 431 || result.href?.startsWith("chrome-error://chromewebdata/") === true;
}

function probeFailedSuggestions(status: number, cdpBase: string): string[] {
  if (status === 431) {
    return [
      "HTTP 431 means the request headers were too large; the CDP Chrome profile likely has stale cookie buildup.",
      `Sign out of ChatGPT in the CDP window, sign back in, then run pro-cli auth capture --cdp ${cdpBase} --json.`,
      "If 431 persists, delete ~/.pro-cli/chrome-profile and rerun pro-cli auth command.",
    ];
  }
  if (status >= 500) {
    return [
      `ChatGPT returned HTTP ${status} on the auth probe; the upstream is likely degraded.`,
      "Reload the CDP ChatGPT tab, wait, and rerun pro-cli doctor --json.",
    ];
  }
  return [
    `The CDP ChatGPT auth session probe returned HTTP ${status}; cannot determine login state.`,
    "Reload the CDP ChatGPT tab and rerun pro-cli doctor --json. If the page is on a non-chatgpt URL, navigate back to https://chatgpt.com/.",
  ];
}

export interface CaptureOptions {
  cdpBase: string;
  jsonPath: string;
  jarPath: string;
  tokenPath: string;
  timeoutMs?: number;
  dryRun?: boolean;
}

export async function captureAuth(options: CaptureOptions): Promise<AuthStatus> {
  if (options.dryRun) {
    throw new ProError("DRY_RUN", "Auth capture dry run does not read cookies.", {
      exitCode: EXIT.success,
      suggestions: ["Run without --dry-run to capture scoped ChatGPT cookies."],
    });
  }

  await pruneVolatileCookiesFromCdp(
    options.cdpBase,
    chatGptOrigins(),
    options.timeoutMs ?? 10_000,
  ).catch(() => undefined);
  const cookies = await getCookiesFromCdp(
    options.cdpBase,
    chatGptOrigins(),
    options.timeoutMs ?? 10_000,
  );
  const cookieExport = toCookieExport(cookies);
  if (cookieExport.cookies.length === 0) {
    throw new ProError("NO_CHATGPT_COOKIES", "No ChatGPT/OpenAI cookies were found via CDP.", {
      exitCode: EXIT.auth,
      suggestions: [
        "Open https://chatgpt.com/ in the CDP Chrome window.",
        "Confirm the logged-in ChatGPT UI is visible.",
        "Retry pro-cli auth capture.",
      ],
    });
  }

  await writePrivateFile(options.jsonPath, `${JSON.stringify(cookieExport, null, 2)}\n`);
  await writePrivateFile(options.jarPath, toNetscapeCookieJar(cookieExport.cookies));
  const accessToken = await getSessionAccessTokenFromPage(options.cdpBase, options.timeoutMs ?? 10_000);
  const sessionToken = toSessionTokenExport(accessToken);
  await writePrivateFile(options.tokenPath, `${JSON.stringify(sessionToken, null, 2)}\n`);

  const summary = cookieSummary(cookieExport.cookies);
  return {
    status: "present",
    cookieJsonPath: options.jsonPath,
    cookieJarPath: options.jarPath,
    sessionTokenPath: options.tokenPath,
    tokenStatus: "present",
    ...(sessionToken.expiresAt ? { tokenExpiresAt: sessionToken.expiresAt } : {}),
    accountIdPresent: Boolean(sessionToken.accountId),
    cookieCount: summary.count,
    expiredCookieCount: summary.expired,
    domains: summary.domains,
    names: summary.names,
    rawValuesPrinted: false,
  };
}

export function defaultCdpBase(port: string | undefined, cdp: string | undefined): string {
  if (cdp) return cdp;
  return `http://127.0.0.1:${port ?? "9222"}`;
}

async function getSessionAccessTokenFromPage(cdpBase: string, timeoutMs: number): Promise<string> {
  const result = await evaluateInCdpPage<{ status: number; hasAccessToken: boolean; accessToken?: string }>(
    cdpBase,
    `fetch("/api/auth/session", { credentials: "include" }).then(async (res) => {
      const json = await res.json().catch(() => null);
      return {
        status: res.status,
        hasAccessToken: typeof json?.accessToken === "string",
        accessToken: typeof json?.accessToken === "string" ? json.accessToken : undefined
      };
    })`,
    timeoutMs,
  );
  if (result.status !== 200 || !result.hasAccessToken || !result.accessToken) {
    throw new ProError("SESSION_TOKEN_UNAVAILABLE", "ChatGPT page did not expose a session access token.", {
      exitCode: EXIT.auth,
      suggestions: [
        "Confirm the CDP tab is on https://chatgpt.com/ and logged in.",
        "Refresh the ChatGPT page and retry pro-cli auth capture.",
      ],
      details: { status: result.status },
    });
  }
  return result.accessToken;
}

async function tokenStatusFields(path: string): Promise<{
  tokenStatus: "missing" | "present" | "expired";
  tokenExpiresAt?: string;
  accountIdPresent: boolean;
}> {
  try {
    const token = await loadSessionToken(path);
    return {
      tokenStatus: isTokenFresh(token) ? "present" : "expired",
      ...(token.expiresAt ? { tokenExpiresAt: token.expiresAt } : {}),
      accountIdPresent: Boolean(token.accountId),
    };
  } catch {
    return { tokenStatus: "missing", accountIdPresent: false };
  }
}

async function readTokenStatus(path: string): Promise<"missing" | "present" | "expired"> {
  return (await tokenStatusFields(path)).tokenStatus;
}

async function readAccountIdPresent(path: string): Promise<boolean> {
  return (await tokenStatusFields(path)).accountIdPresent;
}

export interface ResetProfileOptions {
  home: string;
  port: string;
  noBackup?: boolean;
  noLaunch?: boolean;
  keepBackups?: number;
}

export interface ResetProfileResult {
  profileDir: string;
  killedPids: number[];
  removed: { mode: "backup" | "delete" | "missing"; from: string; to?: string };
  prunedBackups: string[];
  launched: { command: string } | null;
  cdp: string;
  portCollision: PortCollisionInfo;
  legacyArtifacts: LegacyArtifactInfo;
  next: { command: string; reason: string };
}

export async function resetAuthProfile(options: ResetProfileOptions): Promise<ResetProfileResult> {
  const profileDir = join(options.home, "chrome-profile");
  if (!profileDir.startsWith(options.home + "/") && profileDir !== options.home) {
    throw new ProError("RESET_PATH_UNSAFE", "Refusing to reset a profile outside the pro-cli home.", {
      exitCode: EXIT.invalidArgs,
      details: { profileDir, home: options.home },
    });
  }

  const killedPids = killChromeForProfile(profileDir);
  if (killedPids.length > 0) {
    await sleepMs(1500);
    for (const pid of killedPids) {
      try {
        process.kill(pid, 0);
        process.kill(pid, "SIGKILL");
      } catch {
        // already gone
      }
    }
  }

  const removed = await removeProfileDir(profileDir, !options.noBackup);
  const keepBackups = Math.max(0, options.keepBackups ?? 5);
  const prunedBackups = await pruneOldBackups(options.home, keepBackups);

  let launched: ResetProfileResult["launched"] = null;
  if (!options.noLaunch) {
    const command = buildOpenChromeCommand(profileDir, options.port);
    const child = spawn("/bin/sh", ["-c", command], { detached: true, stdio: "ignore" });
    child.unref();
    launched = { command };
  }

  const cdp = `http://127.0.0.1:${options.port}`;
  const portCollision = detectPortCollision(options.port, profileDir);
  const legacyArtifacts = await detectLegacyArtifacts();
  return {
    profileDir,
    killedPids,
    removed,
    prunedBackups,
    launched,
    cdp,
    portCollision,
    legacyArtifacts,
    next: {
      command: `pro-cli auth capture --cdp ${cdp} --json`,
      reason: launched
        ? "Profile reset and Chrome relaunched. Sign in to ChatGPT in the new window, then run the capture command."
        : "Profile reset. Open a new Chrome with pro-cli auth command, sign in, then run the capture command.",
    },
  };
}

function killChromeForProfile(profileDir: string): number[] {
  const ps = spawnSync("ps", ["axo", "pid=,command="], { encoding: "utf8" });
  if (ps.status !== 0) return [];
  const needle = `--user-data-dir=${profileDir}`;
  const pids: number[] = [];
  for (const rawLine of ps.stdout.split("\n")) {
    const line = rawLine.trim();
    if (!line.includes(needle)) continue;
    const match = line.match(/^(\d+)\s/);
    if (!match) continue;
    pids.push(Number.parseInt(match[1], 10));
  }
  for (const pid of pids) {
    try {
      process.kill(pid, "SIGTERM");
    } catch {
      // process may have exited between ps and kill
    }
  }
  return pids;
}

async function removeProfileDir(
  dir: string,
  backup: boolean,
): Promise<ResetProfileResult["removed"]> {
  try {
    await stat(dir);
  } catch {
    return { mode: "missing", from: dir };
  }
  if (backup) {
    const ts = backupTimestamp();
    let target = `${dir}.backup-${ts}`;
    let suffix = 1;
    while (await pathExists(target)) {
      target = `${dir}.backup-${ts}-${suffix}`;
      suffix += 1;
    }
    await rename(dir, target);
    return { mode: "backup", from: dir, to: target };
  }
  await rm(dir, { recursive: true, force: true });
  return { mode: "delete", from: dir };
}

async function pruneOldBackups(home: string, keep: number): Promise<string[]> {
  const entries = await readdir(home).catch(() => [] as string[]);
  const backups = entries
    .filter((name) => /^chrome-profile\.backup-/.test(name))
    .sort()
    .reverse();
  const toRemove = backups.slice(keep);
  const removed: string[] = [];
  for (const name of toRemove) {
    const full = join(home, name);
    try {
      await rm(full, { recursive: true, force: true });
      removed.push(full);
    } catch {
      // ignore individual prune failures
    }
  }
  return removed;
}

function buildOpenChromeCommand(profileDir: string, port: string): string {
  const url = "https://chatgpt.com/";
  if (process.platform === "darwin") {
    return `open -na "Google Chrome" --args --user-data-dir='${profileDir}' --remote-debugging-port=${port} ${url}`;
  }
  if (process.platform === "win32") {
    return `start "" chrome.exe --user-data-dir="${profileDir}" --remote-debugging-port=${port} ${url}`;
  }
  return `google-chrome --user-data-dir='${profileDir}' --remote-debugging-port=${port} ${url}`;
}

function backupTimestamp(): string {
  const d = new Date();
  const pad = (n: number, w = 2) => n.toString().padStart(w, "0");
  return `${d.getFullYear()}${pad(d.getMonth() + 1)}${pad(d.getDate())}-${pad(d.getHours())}${pad(d.getMinutes())}${pad(d.getSeconds())}`;
}

async function pathExists(path: string): Promise<boolean> {
  try {
    await stat(path);
    return true;
  } catch {
    return false;
  }
}

function sleepMs(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

export interface PortListenerInfo {
  pid: number;
  command: string;
  isChrome: boolean;
  userDataDir?: string;
  matchesProfile: boolean;
}

export interface PortCollisionInfo {
  port: string;
  inUse: boolean;
  listeners: PortListenerInfo[];
  conflict: boolean;
  warning?: string;
}

export function detectPortCollision(port: string, expectedProfileDir: string): PortCollisionInfo {
  const lsof = spawnSync("lsof", ["-i", `:${port}`, "-sTCP:LISTEN", "-P", "-n"], { encoding: "utf8" });
  const listeners: PortListenerInfo[] = [];
  if (lsof.status === 0 && lsof.stdout) {
    const seen = new Set<number>();
    const lines = lsof.stdout.split("\n").slice(1);
    for (const line of lines) {
      const parts = line.trim().split(/\s+/);
      if (parts.length < 2) continue;
      const pid = Number.parseInt(parts[1], 10);
      if (!Number.isFinite(pid) || seen.has(pid)) continue;
      seen.add(pid);
      const ps = spawnSync("ps", ["-o", "command=", "-p", String(pid)], { encoding: "utf8" });
      const command = (ps.stdout ?? "").trim();
      const isChrome = /chrome/i.test(command);
      const userDataMatch = command.match(/--user-data-dir=([^\s]+)/);
      const userDataDir = userDataMatch ? userDataMatch[1].replace(/^['"]|['"]$/g, "") : undefined;
      const matchesProfile = userDataDir === expectedProfileDir;
      listeners.push({ pid, command, isChrome, userDataDir, matchesProfile });
    }
  }
  const conflict = listeners.some((l) => l.isChrome && !l.matchesProfile);
  const info: PortCollisionInfo = {
    port,
    inUse: listeners.length > 0,
    listeners,
    conflict,
  };
  if (conflict) {
    const conflictPids = listeners.filter((l) => l.isChrome && !l.matchesProfile).map((l) => l.pid);
    info.warning = `Another Chrome (pid ${conflictPids.join(", ")}) is bound to port ${port} with a different --user-data-dir. CDP requests will race between Chromes; results become non-deterministic. Quit those Chrome instances or pick a different --port.`;
  } else if (listeners.length > 1) {
    info.warning = `Multiple processes are listening on port ${port}; this can cause CDP request routing to be non-deterministic.`;
  }
  return info;
}

export interface WindowBounds {
  left?: number;
  top?: number;
  width?: number;
  height?: number;
  windowState?: "normal" | "minimized" | "maximized" | "fullscreen";
}

export interface MoveWindowResult {
  cdpBase: string;
  targetId: string;
  windowId: number;
  before: WindowBounds;
  after: WindowBounds;
  note: string;
}

const HIDDEN_BOUNDS: WindowBounds = { left: -32000, top: -32000, width: 400, height: 300 };
const SHOWN_BOUNDS: WindowBounds = { left: 100, top: 100, width: 1200, height: 800 };

export async function setProCliChromeWindowBounds(
  cdpBase: string,
  bounds: WindowBounds,
): Promise<MoveWindowResult> {
  const targetId = await findChatGptTargetId(cdpBase);
  if (!targetId) {
    throw new ProError("CHATGPT_PAGE_MISSING", "No chatgpt.com tab is available over CDP.", {
      exitCode: EXIT.auth,
      suggestions: [
        "Open the ChatGPT Chrome window via pro-cli auth command, then retry.",
        "If the window was closed, your stored auth still works; relaunching restores access.",
      ],
    });
  }
  const window = await callBrowserCdp<{ windowId: number; bounds: WindowBounds }>(
    cdpBase,
    "Browser.getWindowForTarget",
    { targetId },
  );
  await callBrowserCdp<unknown>(cdpBase, "Browser.setWindowBounds", {
    windowId: window.windowId,
    bounds,
  });
  const isHidden = (bounds.left ?? 0) < -10000 || (bounds.top ?? 0) < -10000;
  const note = isHidden
    ? "Window parked off-screen. CDP keeps working. Run pro-cli auth show to bring it back; closing Chrome from the dock loses the session."
    : "Window restored on-screen.";
  return { cdpBase, targetId, windowId: window.windowId, before: window.bounds, after: bounds, note };
}

export function hideProCliChromeWindow(cdpBase: string): Promise<MoveWindowResult> {
  return setProCliChromeWindowBounds(cdpBase, HIDDEN_BOUNDS);
}

export function showProCliChromeWindow(cdpBase: string): Promise<MoveWindowResult> {
  return setProCliChromeWindowBounds(cdpBase, SHOWN_BOUNDS);
}

export interface LegacyArtifactInfo {
  legacyHome: string;
  legacyHomeExists: boolean;
  legacyProfileDir: string;
  legacyProfileExists: boolean;
  warning?: string;
}

export async function detectLegacyArtifacts(): Promise<LegacyArtifactInfo> {
  const legacyHome = join(homedir(), ".pro");
  const legacyProfileDir = join(legacyHome, "chrome-profile");
  const legacyHomeExists = await pathExists(legacyHome);
  const legacyProfileExists = await pathExists(legacyProfileDir);
  const info: LegacyArtifactInfo = {
    legacyHome,
    legacyHomeExists,
    legacyProfileDir,
    legacyProfileExists,
  };
  if (legacyProfileExists) {
    info.warning = `Legacy Chrome profile dir ${legacyProfileDir} still exists. An external command opening it can bind the CDP port and conflict with pro-cli's profile. Move it aside (mv ${legacyProfileDir} ${legacyProfileDir}.legacy-backup) or delete it.`;
  } else if (legacyHomeExists) {
    info.warning = `Legacy pro-cli home ${legacyHome} still exists. Safe to remove if pro-cli is functioning from ~/.pro-cli.`;
  }
  return info;
}
`````

## File: src/cdp.ts
`````typescript
import { EXIT, ProError } from "./errors";
import { isVolatileCookieName } from "./cookies";

interface JsonRpcResponse<T> {
  id?: number;
  result?: T;
  error?: { code: number; message: string };
}

export interface CdpCookie {
  name: string;
  value: string;
  domain: string;
  path: string;
  expires?: number;
  size?: number;
  httpOnly?: boolean;
  secure?: boolean;
  session?: boolean;
  sameSite?: string;
}

export async function getCookiesFromCdp(
  cdpBase: string,
  urls: string[],
  timeoutMs = 10_000,
): Promise<CdpCookie[]> {
  const wsUrl = await resolveCdpWebSocketUrl(cdpBase);
  const client = await CdpClient.connect(wsUrl, timeoutMs);
  try {
    return await readCookiesForUrls(client, urls);
  } finally {
    client.close();
  }
}

export interface CookiePruneResult {
  checked: number;
  deleted: number;
  names: string[];
}

export async function pruneVolatileCookiesFromCdp(
  cdpBase: string,
  urls: string[],
  timeoutMs = 10_000,
): Promise<CookiePruneResult> {
  const wsUrl = await resolveCdpWebSocketUrl(cdpBase);
  const client = await CdpClient.connect(wsUrl, timeoutMs);
  try {
    const cookies = await readCookiesForUrls(client, urls);
    const volatileCookies = cookies.filter((cookie) => isVolatileCookieName(cookie.name));
    for (const cookie of volatileCookies) {
      await client.send<unknown>("Network.deleteCookies", {
        name: cookie.name,
        domain: cookie.domain,
        path: cookie.path || "/",
      });
    }
    return {
      checked: cookies.length,
      deleted: volatileCookies.length,
      names: [...new Set(volatileCookies.map((cookie) => cookie.name))].sort(),
    };
  } finally {
    client.close();
  }
}

export interface CookieBloatRecoveryResult extends CookiePruneResult {
  navigated: boolean;
}

export async function recoverCookieBloatInCdp(
  cdpBase: string,
  urls: string[],
  timeoutMs = 10_000,
): Promise<CookieBloatRecoveryResult> {
  const pruned = await pruneVolatileCookiesFromCdp(cdpBase, urls, timeoutMs);
  if (pruned.deleted === 0) return { ...pruned, navigated: false };

  await navigateCdpPage(cdpBase, "https://chatgpt.com/", timeoutMs);
  await sleepMs(Math.min(1500, timeoutMs));
  return { ...pruned, navigated: true };
}

export async function evaluateInCdpPage<T>(
  cdpBase: string,
  expression: string,
  timeoutMs = 10_000,
): Promise<T> {
  const wsUrl = await resolveRequiredPageWebSocketUrl(cdpBase);
  const client = await CdpClient.connect(wsUrl, timeoutMs);
  try {
    const response = await client.send<{
      result?: { value?: T };
      exceptionDetails?: { text?: string };
    }>("Runtime.evaluate", {
      expression,
      awaitPromise: true,
      returnByValue: true,
    });
    if (response.exceptionDetails) {
      throw new ProError(
        "CDP_EVALUATION_FAILED",
        response.exceptionDetails.text ?? "Chrome page evaluation failed.",
        {
          exitCode: EXIT.auth,
          suggestions: ["Open the logged-in ChatGPT page and retry auth capture."],
        },
      );
    }
    return response.result?.value as T;
  } finally {
    client.close();
  }
}

async function resolveRequiredPageWebSocketUrl(cdpBase: string): Promise<string> {
  const base = cdpBase.replace(/\/$/, "");
  const pageWsUrl = await resolvePageWebSocketUrl(base);
  if (pageWsUrl) return pageWsUrl;

  await resolveBrowserWebSocketUrl(base);
  throw new ProError("CHATGPT_PAGE_MISSING", `No inspectable page is available over CDP at ${base}.`, {
    exitCode: EXIT.auth,
    suggestions: [
      "Open the Chrome command from pro-cli auth command.",
      "Confirm the CDP Chrome window has a https://chatgpt.com/ tab.",
    ],
    details: { cdpBase: base },
  });
}

async function resolveCdpWebSocketUrl(cdpBase: string): Promise<string> {
  const base = cdpBase.replace(/\/$/, "");
  const pageWsUrl = await resolvePageWebSocketUrl(base);
  if (pageWsUrl) return pageWsUrl;

  return resolveBrowserWebSocketUrl(base);
}

export async function callBrowserCdp<T>(
  cdpBase: string,
  method: string,
  params?: Record<string, unknown>,
  timeoutMs = 5000,
): Promise<T> {
  const wsUrl = await resolveBrowserWebSocketUrl(cdpBase.replace(/\/$/, ""));
  const client = await CdpClient.connect(wsUrl, timeoutMs);
  try {
    return await client.send<T>(method, params);
  } finally {
    client.close();
  }
}

export async function navigateCdpPage(
  cdpBase: string,
  url: string,
  timeoutMs = 10_000,
): Promise<void> {
  const wsUrl = await resolveRequiredPageWebSocketUrl(cdpBase);
  const client = await CdpClient.connect(wsUrl, timeoutMs);
  try {
    await client.send<unknown>("Page.enable");
    await client.send<unknown>("Page.navigate", { url });
  } finally {
    client.close();
  }
}

export async function findChatGptTargetId(cdpBase: string): Promise<string | null> {
  const base = cdpBase.replace(/\/$/, "");
  let response: Response;
  try {
    response = await fetch(`${base}/json`);
  } catch {
    return null;
  }
  if (!response.ok) return null;
  const targets = (await response.json().catch(() => [])) as Array<{
    id?: string;
    type?: string;
    url?: string;
  }>;
  const chatgpt = targets.find((t) => t.type === "page" && t.url?.startsWith("https://chatgpt.com/"));
  return chatgpt?.id ?? null;
}

async function resolveBrowserWebSocketUrl(base: string): Promise<string> {
  let response: Response;
  try {
    response = await fetch(`${base}/json/version`);
  } catch (error) {
    throw new ProError("CDP_UNAVAILABLE", `Cannot connect to Chrome CDP at ${base}.`, {
      exitCode: EXIT.auth,
      suggestions: [
        "Open Chrome with --remote-debugging-port=9222.",
        "Pass --cdp http://127.0.0.1:<port> if Chrome uses a different CDP port.",
      ],
      cause: error,
    });
  }
  if (!response.ok) {
    throw new ProError("CDP_UNAVAILABLE", `Chrome CDP returned HTTP ${response.status}.`, {
      exitCode: EXIT.auth,
      suggestions: ["Check the CDP URL and remote debugging port."],
    });
  }
  const payload = (await response.json()) as { webSocketDebuggerUrl?: string };
  if (!payload.webSocketDebuggerUrl) {
    throw new ProError("CDP_UNAVAILABLE", "Chrome CDP did not expose a browser websocket.", {
      exitCode: EXIT.auth,
      suggestions: ["Use the browser-level CDP endpoint from /json/version."],
    });
  }
  return payload.webSocketDebuggerUrl;
}

function filterCookiesForUrls(cookies: CdpCookie[], urls: string[]): CdpCookie[] {
  return cookies.filter((cookie) => urls.some((url) => cookieAppliesToUrl(cookie, url)));
}

async function readCookiesForUrls(client: CdpClient, urls: string[]): Promise<CdpCookie[]> {
  try {
    const response = await client.send<{ cookies: CdpCookie[] }>("Network.getCookies", { urls });
    return filterCookiesForUrls(response.cookies ?? [], urls);
  } catch (error) {
    if (!isMissingCdpMethod(error)) throw error;
    const response = await client.send<{ cookies: CdpCookie[] }>("Storage.getCookies");
    return filterCookiesForUrls(response.cookies ?? [], urls);
  }
}

function cookieAppliesToUrl(cookie: CdpCookie, url: string): boolean {
  let parsed: URL;
  try {
    parsed = new URL(url);
  } catch {
    return false;
  }
  const host = parsed.hostname.toLowerCase();
  const domain = (cookie.domain || "").replace(/^\./, "").toLowerCase();
  if (!domain) return false;
  if (host !== domain && !host.endsWith(`.${domain}`)) return false;

  const cookiePath = cookie.path || "/";
  if (cookiePath === "/") return true;
  if (!parsed.pathname.startsWith(cookiePath)) return false;
  const next = parsed.pathname[cookiePath.length];
  return next === undefined || next === "/";
}

function isMissingCdpMethod(error: unknown): boolean {
  return (
    error instanceof ProError &&
    error.code === "CDP_COMMAND_FAILED" &&
    error.details?.cdpCode === -32601
  );
}

function sleepMs(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

async function resolvePageWebSocketUrl(base: string): Promise<string | null> {
  let response: Response;
  try {
    response = await fetch(`${base}/json`);
  } catch {
    return null;
  }
  if (!response.ok) return null;
  const targets = (await response.json().catch(() => [])) as Array<{
    type?: string;
    url?: string;
    webSocketDebuggerUrl?: string;
  }>;
  const pages = targets.filter((target) => target.type === "page" && target.webSocketDebuggerUrl);
  const chatgpt = pages.find((target) => target.url?.startsWith("https://chatgpt.com/"));
  return chatgpt?.webSocketDebuggerUrl ?? pages[0]?.webSocketDebuggerUrl ?? null;
}

class CdpClient {
  private nextId = 1;
  private readonly pending = new Map<
    number,
    {
      resolve: (value: unknown) => void;
      reject: (error: Error) => void;
      timer: ReturnType<typeof setTimeout>;
    }
  >();

  private constructor(
    private readonly socket: WebSocket,
    private readonly timeoutMs: number,
  ) {
    this.socket.addEventListener("message", (event) => this.handleMessage(String(event.data)));
    this.socket.addEventListener("error", () => this.rejectAll("CDP websocket error."));
    this.socket.addEventListener("close", () => this.rejectAll("CDP websocket closed."));
  }

  static async connect(url: string, timeoutMs: number): Promise<CdpClient> {
    const socket = new WebSocket(url);
    await new Promise<void>((resolve, reject) => {
      const timer = setTimeout(() => reject(new Error("CDP websocket open timed out.")), timeoutMs);
      socket.addEventListener("open", () => {
        clearTimeout(timer);
        resolve();
      });
      socket.addEventListener("error", () => {
        clearTimeout(timer);
        reject(new Error("CDP websocket failed to open."));
      });
    }).catch((error) => {
      throw new ProError("CDP_UNAVAILABLE", "Cannot open Chrome CDP websocket.", {
        exitCode: EXIT.auth,
        suggestions: ["Confirm Chrome is running with remote debugging enabled."],
        cause: error,
      });
    });
    return new CdpClient(socket, timeoutMs);
  }

  async send<T>(method: string, params?: Record<string, unknown>): Promise<T> {
    const id = this.nextId;
    this.nextId += 1;
    const payload = JSON.stringify({ id, method, params });
    const result = new Promise<T>((resolve, reject) => {
      const timer = setTimeout(() => {
        this.pending.delete(id);
        reject(
          new ProError("CDP_TIMEOUT", `Chrome CDP command ${method} timed out.`, {
            exitCode: EXIT.timeout,
            suggestions: ["Retry auth capture or restart the CDP Chrome instance."],
          }),
        );
      }, this.timeoutMs);
      this.pending.set(id, { resolve: resolve as (value: unknown) => void, reject, timer });
    });
    this.socket.send(payload);
    return result;
  }

  close(): void {
    this.socket.close();
  }

  private handleMessage(raw: string): void {
    const message = JSON.parse(raw) as JsonRpcResponse<unknown>;
    if (!message.id) return;
    const pending = this.pending.get(message.id);
    if (!pending) return;
    clearTimeout(pending.timer);
    this.pending.delete(message.id);
    if (message.error) {
      pending.reject(
        new ProError("CDP_COMMAND_FAILED", message.error.message, {
          exitCode: EXIT.auth,
          suggestions: ["Retry with a fresh logged-in Chrome/CDP session."],
          details: { cdpCode: message.error.code },
        }),
      );
      return;
    }
    pending.resolve(message.result);
  }

  private rejectAll(message: string): void {
    for (const pending of this.pending.values()) {
      clearTimeout(pending.timer);
      pending.reject(new Error(message));
    }
    this.pending.clear();
  }
}
`````

## File: src/cli.ts
`````typescript
#!/usr/bin/env bun
import { runCli } from "./app";

const exitCode = await runCli(Bun.argv.slice(2), {
  stdout: (text) => process.stdout.write(text),
  stderr: (text) => process.stderr.write(text),
  stdoutIsTTY: process.stdout.isTTY === true,
  env: process.env,
  cwd: process.cwd(),
});

process.exit(exitCode);
`````

## File: src/config.ts
`````typescript
import { access, chmod, mkdir, readFile, rename, writeFile } from "node:fs/promises";
import { dirname, join, resolve } from "node:path";
import { homedir } from "node:os";
import { EXIT, ProError } from "./errors";

export interface ProConfig {
  cookieJsonPath?: string;
  cookieJarPath?: string;
  sessionTokenPath?: string;
  defaultModel?: string;
  defaultReasoning?: string;
}

export interface RuntimePaths {
  home: string;
  configPath: string;
  cookieJsonPath: string;
  cookieJarPath: string;
  sessionTokenPath: string;
  dbPath: string;
}

const DEFAULT_HOME = "~/.pro-cli";
const LEGACY_HOME = "~/.pro";

export function resolveHome(env: Record<string, string | undefined>): string {
  return expandPath(env.PRO_CLI_HOME || DEFAULT_HOME);
}

export async function migrateLegacyDefaultHome(
  env: Record<string, string | undefined>,
  homedirOverride?: string,
): Promise<void> {
  if (env.PRO_CLI_HOME) return;
  const baseHome = homedirOverride ?? homedir();
  const nextHome = join(baseHome, ".pro-cli");
  const legacyHome = join(baseHome, ".pro");
  if (await pathExists(nextHome)) {
    await rewriteMigratedConfigPaths(nextHome, legacyHome);
    return;
  }
  if (!(await pathExists(legacyHome))) return;
  try {
    await rename(legacyHome, nextHome);
  } catch (error) {
    const code = (error as NodeJS.ErrnoException).code;
    if (code === "EACCES" || code === "EPERM" || code === "EXDEV") return;
    throw error;
  }
  await chmod(nextHome, 0o700).catch(() => undefined);
  await rewriteMigratedConfigPaths(nextHome, legacyHome);
}

export function resolvePaths(
  env: Record<string, string | undefined>,
  config: ProConfig = {},
): RuntimePaths {
  const home = resolveHome(env);
  const cookieJsonPath = env.CHATGPT_COOKIE_JSON || config.cookieJsonPath || join(home, "cookies", "chatgpt.json");
  const cookieJarPath = env.CHATGPT_COOKIE_JAR || config.cookieJarPath || join(home, "cookies", "chatgpt.txt");
  const sessionTokenPath =
    env.CHATGPT_SESSION_TOKEN_JSON ||
    config.sessionTokenPath ||
    join(home, "tokens", "chatgpt-session.json");
  return {
    home,
    configPath: join(home, "config.json"),
    cookieJsonPath: expandPath(cookieJsonPath),
    cookieJarPath: expandPath(cookieJarPath),
    sessionTokenPath: expandPath(sessionTokenPath),
    dbPath: join(home, "jobs.sqlite"),
  };
}

export async function ensurePrivateDir(path: string): Promise<void> {
  await mkdir(path, { recursive: true, mode: 0o700 });
  await chmod(path, 0o700).catch(() => undefined);
}

export async function writePrivateFile(path: string, content: string): Promise<void> {
  await ensurePrivateDir(dirname(path));
  await writeFile(path, content, { mode: 0o600 });
  await chmod(path, 0o600).catch(() => undefined);
}

export async function loadConfig(env: Record<string, string | undefined>): Promise<ProConfig> {
  const home = resolveHome(env);
  const configPath = join(home, "config.json");
  try {
    const raw = await readFile(configPath, "utf8");
    return JSON.parse(raw) as ProConfig;
  } catch (error) {
    const code = (error as NodeJS.ErrnoException).code;
    if (code === "ENOENT") return {};
    if (isPermissionError(error)) throw configAccessError("CONFIG_UNREADABLE", "read", configPath, error);
    throw error;
  }
}

export async function saveConfig(
  env: Record<string, string | undefined>,
  config: ProConfig,
): Promise<void> {
  const home = resolveHome(env);
  const configPath = join(home, "config.json");
  try {
    await ensurePrivateDir(home);
    await writePrivateFile(configPath, `${JSON.stringify(config, null, 2)}\n`);
  } catch (error) {
    if (isPermissionError(error)) throw configAccessError("CONFIG_UNWRITABLE", "write", configPath, error);
    throw error;
  }
}

export function expandPath(path: string): string {
  if (path === "~") return homedir();
  if (path.startsWith("~/")) return resolve(homedir(), path.slice(2));
  return resolve(path);
}

async function pathExists(path: string): Promise<boolean> {
  try {
    await access(path);
    return true;
  } catch {
    return false;
  }
}

async function rewriteMigratedConfigPaths(nextHome: string, legacyHome: string): Promise<void> {
  const configPath = join(nextHome, "config.json");
  let raw: string;
  try {
    raw = await readFile(configPath, "utf8");
  } catch {
    return;
  }

  const config = JSON.parse(raw) as ProConfig;
  const rewritten: ProConfig = {
    ...config,
    cookieJsonPath: rewriteHomePrefix(config.cookieJsonPath, legacyHome, nextHome),
    cookieJarPath: rewriteHomePrefix(config.cookieJarPath, legacyHome, nextHome),
    sessionTokenPath: rewriteHomePrefix(config.sessionTokenPath, legacyHome, nextHome),
  };
  await writePrivateFile(configPath, `${JSON.stringify(rewritten, null, 2)}\n`);
}

function rewriteHomePrefix(path: string | undefined, fromHome: string, toHome: string): string | undefined {
  if (!path) return undefined;
  if (path === fromHome) return toHome;
  const prefix = `${fromHome}/`;
  return path.startsWith(prefix) ? join(toHome, path.slice(prefix.length)) : path;
}

function isPermissionError(error: unknown): boolean {
  const code = (error as NodeJS.ErrnoException).code;
  return code === "EACCES" || code === "EPERM";
}

function configAccessError(
  code: "CONFIG_UNREADABLE" | "CONFIG_UNWRITABLE",
  action: "read" | "write",
  configPath: string,
  error: unknown,
): ProError {
  const errno = (error as NodeJS.ErrnoException).code;
  const syscall = (error as NodeJS.ErrnoException).syscall;
  return new ProError(code, `Cannot ${action} pro-cli config at ${configPath}.`, {
    exitCode: EXIT.auth,
    suggestions: [
      `Fix local ownership/permissions for ${configPath}, or set PRO_CLI_HOME to a writable pro-cli home.`,
      "After fixing storage, run pro-cli doctor --json. Do not send probe or smoke-test queries; ask/job calls spend Pro quota.",
    ],
    details: {
      configPath,
      ...(errno ? { errno } : {}),
      ...(syscall ? { syscall } : {}),
    },
    cause: error,
  });
}
`````

## File: src/cookies.ts
`````typescript
import { readFile } from "node:fs/promises";

export interface BrowserCookie {
  name: string;
  value: string;
  domain: string;
  path: string;
  expires?: number;
  secure?: boolean;
  httpOnly?: boolean;
  sameSite?: string;
  includeSubdomains?: boolean;
}

export interface CookieExport {
  version: 1;
  generatedAt: string;
  source: "pro-cli-cdp";
  targetUrl: string;
  origins: string[];
  cookies: BrowserCookie[];
}

const CHATGPT_ORIGINS = [
  "https://chatgpt.com/",
  "https://auth.openai.com/",
  "https://openai.com/",
  "https://sentinel.openai.com/",
  "https://ws.chatgpt.com/",
] as const;

export function chatGptOrigins(): string[] {
  return [...CHATGPT_ORIGINS];
}

export function toCookieExport(cookies: BrowserCookie[]): CookieExport {
  return {
    version: 1,
    generatedAt: new Date().toISOString(),
    source: "pro-cli-cdp",
    targetUrl: "https://chatgpt.com/",
    origins: chatGptOrigins(),
    cookies: sanitizeCookies(cookies),
  };
}

export function sanitizeCookies(cookies: BrowserCookie[]): BrowserCookie[] {
  const deduped = new Map<string, BrowserCookie>();
  for (const cookie of cookies) {
    if (!cookie.name || cookie.value === undefined) continue;
    if (isVolatileCookieName(cookie.name)) continue;
    const rawDomain = cookie.domain || "chatgpt.com";
    const domain = stripLeadingDot(rawDomain);
    const path = cookie.path || "/";
    const includeSubdomains = rawDomain.startsWith(".") || cookie.includeSubdomains === true;
    deduped.set(`${cookie.name}|${domain}|${path}`, {
      name: cookie.name,
      value: cookie.value,
      domain,
      path,
      ...(cookie.expires ? { expires: cookie.expires } : {}),
      ...(cookie.secure !== undefined ? { secure: cookie.secure } : {}),
      ...(cookie.httpOnly !== undefined ? { httpOnly: cookie.httpOnly } : {}),
      ...(cookie.sameSite ? { sameSite: cookie.sameSite } : {}),
      ...(includeSubdomains ? { includeSubdomains: true } : {}),
    });
  }
  return [...deduped.values()].sort((left, right) => left.name.localeCompare(right.name));
}

export function isVolatileCookieName(name: string): boolean {
  // Per-conversation stream keys accumulate quickly and can push ChatGPT over HTTP header limits.
  return name.startsWith("conv_key_");
}

export function cookieSummary(cookies: BrowserCookie[]): {
  count: number;
  domains: string[];
  names: string[];
  expired: number;
} {
  const now = Date.now() / 1000;
  const domains = new Set<string>();
  const names = new Set<string>();
  let expired = 0;
  for (const cookie of cookies) {
    domains.add(stripLeadingDot(cookie.domain));
    names.add(cookie.name);
    if (cookie.expires && cookie.expires <= now) expired += 1;
  }
  return {
    count: cookies.length,
    domains: [...domains].sort(),
    names: [...names].sort(),
    expired,
  };
}

export async function loadCookieExport(path: string): Promise<CookieExport> {
  const raw = await readFile(path, "utf8");
  const parsed = JSON.parse(raw) as CookieExport | BrowserCookie[] | { cookies: BrowserCookie[] };
  if (Array.isArray(parsed)) return toCookieExport(parsed);
  if ("cookies" in parsed && !("version" in parsed)) return toCookieExport(parsed.cookies);
  return parsed as CookieExport;
}

export function toCookieHeader(cookies: BrowserCookie[]): string {
  const now = Date.now() / 1000;
  return sanitizeCookies(cookies)
    .filter((cookie) => !cookie.expires || cookie.expires > now)
    .map((cookie) => `${cookie.name}=${cookie.value}`)
    .join("; ");
}

export function toNetscapeCookieJar(cookies: BrowserCookie[]): string {
  const lines = [
    "# Netscape HTTP Cookie File",
    "# Generated by pro-cli. Raw cookie values are sensitive.",
  ];
  for (const cookie of sanitizeCookies(cookies)) {
    const domain = stripLeadingDot(cookie.domain);
    const includeSubdomains = cookie.includeSubdomains ? "TRUE" : "FALSE";
    const secure = cookie.secure ? "TRUE" : "FALSE";
    const expires = Math.trunc(cookie.expires ?? 0);
    lines.push(
      [domain, includeSubdomains, cookie.path || "/", secure, expires, cookie.name, cookie.value].join(
        "\t",
      ),
    );
  }
  return `${lines.join("\n")}\n`;
}

function stripLeadingDot(domain: string): string {
  return domain.startsWith(".") ? domain.slice(1) : domain;
}
`````

## File: src/daemon.ts
`````typescript
import { spawn } from "node:child_process";
import { randomUUID, createHash } from "node:crypto";
import { openSync } from "node:fs";
import { readFile, unlink } from "node:fs/promises";
import { join } from "node:path";
import { tmpdir } from "node:os";
import type { RuntimePaths } from "./config";
import { ensurePrivateDir, writePrivateFile } from "./config";
import { EXIT, ProError, type ErrorPayload, type ExitCode, toProError } from "./errors";
import { executeClaimedJob, waitForJob, waitTimeoutError, type JobWaitOutcome } from "./executor";
import { JobStore, redactJob, type CreateJobInput } from "./jobs";
import type { CliIO } from "./output";

const DEFAULT_DAEMON_START_TIMEOUT_MS = 5_000;
const DEFAULT_DAEMON_POLL_MS = 500;
const HEARTBEAT_MS = 2_000;

export interface DaemonEndpoint {
  version: 1;
  pid: number;
  port: number;
  token: string;
  home: string;
  startedAt: string;
  updatedAt: string;
  logPath: string;
}

export interface DaemonStatus {
  state: "running" | "stopped" | "stale";
  endpointPath: string;
  logPath: string;
  pid?: number;
  port?: number;
  home: string;
  updatedAt?: string;
  processAlive?: boolean;
  message?: string;
}

export interface DaemonStartResult {
  started: boolean;
  status: DaemonStatus;
  client: DaemonClient;
}

interface DaemonRuntimePaths {
  dir: string;
  endpointPath: string;
  logPath: string;
}

export class DaemonClient {
  constructor(private readonly endpoint: DaemonEndpoint) {}

  async health(): Promise<Record<string, unknown>> {
    return this.request("GET", "/health");
  }

  async createJob(input: CreateJobInput): Promise<Record<string, unknown>> {
    return this.request("POST", "/jobs", input);
  }

  async status(jobId: string): Promise<Record<string, unknown>> {
    return this.request("GET", `/jobs/${encodeURIComponent(jobId)}`);
  }

  async result(jobId: string): Promise<Record<string, unknown>> {
    return this.request("GET", `/jobs/${encodeURIComponent(jobId)}/result`);
  }

  async wait(
    jobId: string,
    timeoutMs: number,
    pollMs: number,
    softTimeout = false,
  ): Promise<Record<string, unknown>> {
    return this.request("POST", `/jobs/${encodeURIComponent(jobId)}/wait`, {
      timeoutMs,
      pollMs,
      softTimeout,
    });
  }

  async cancel(jobId: string): Promise<Record<string, unknown>> {
    return this.request("POST", `/jobs/${encodeURIComponent(jobId)}/cancel`);
  }

  async jobs(limit: number): Promise<Record<string, unknown>> {
    return this.request("GET", `/jobs?limit=${encodeURIComponent(String(limit))}`);
  }

  async shutdown(): Promise<Record<string, unknown>> {
    return this.request("POST", "/shutdown");
  }

  private async request(
    method: string,
    path: string,
    body?: unknown,
  ): Promise<Record<string, unknown>> {
    let response: Response;
    try {
      response = await fetch(`http://127.0.0.1:${this.endpoint.port}${path}`, {
        method,
        headers: {
          authorization: `Bearer ${this.endpoint.token}`,
          ...(body ? { "content-type": "application/json" } : {}),
        },
        body: body ? JSON.stringify(body) : undefined,
      });
    } catch (error) {
      throw new ProError("DAEMON_UNAVAILABLE", "Cannot reach the pro-cli daemon.", {
        exitCode: EXIT.network,
        suggestions: ["Run pro-cli daemon start --json.", "Inspect the daemon log path from pro-cli daemon status --json."],
        cause: error,
      });
    }

    const payload = (await response.json().catch(() => null)) as
      | { ok: true; data: Record<string, unknown> }
      | { ok: false; error: ErrorPayload }
      | null;
    if (payload?.ok) return payload.data;
    if (payload?.ok === false) {
      throw new ProError(payload.error.code, payload.error.message, {
        exitCode: response.ok ? EXIT.internal : exitCodeForStatus(response.status),
        suggestions: payload.error.suggestions,
        details: payload.error.details,
      });
    }
    throw new ProError("DAEMON_BAD_RESPONSE", `pro-cli daemon returned HTTP ${response.status}.`, {
      exitCode: exitCodeForStatus(response.status),
      suggestions: ["Run pro-cli daemon restart --json."],
    });
  }
}

export async function ensureDaemonRunning(paths: RuntimePaths, io: CliIO): Promise<DaemonStartResult> {
  const connected = await connectDaemon(paths);
  if (connected) {
    return { started: false, status: connected.status, client: connected.client };
  }

  const before = await getDaemonStatus(paths);
  if (before.pid && before.processAlive) {
    try {
      process.kill(before.pid, "SIGTERM");
    } catch {
      // The process may have exited between status and restart.
    }
  }
  await startDaemonProcess(paths, io);

  const deadline = Date.now() + DEFAULT_DAEMON_START_TIMEOUT_MS;
  while (Date.now() < deadline) {
    const next = await connectDaemon(paths);
    if (next) return { started: true, status: next.status, client: next.client };
    await sleep(100);
  }

  const status = await getDaemonStatus(paths);
  throw new ProError("DAEMON_START_FAILED", "pro-cli daemon did not become ready.", {
    exitCode: EXIT.internal,
    suggestions: [
      "Run pro-cli daemon status --json.",
      `Inspect the daemon log: ${status.logPath}`,
    ],
    details: { status },
  });
}

export async function getDaemonStatus(paths: RuntimePaths): Promise<DaemonStatus> {
  const runtime = daemonRuntimePaths(paths.home);
  const endpoint = await readEndpoint(runtime.endpointPath).catch(() => null);
  if (!endpoint) {
    return { state: "stopped", endpointPath: runtime.endpointPath, logPath: runtime.logPath, home: paths.home };
  }
  if (endpoint.home !== paths.home) {
    return {
      state: "stale",
      endpointPath: runtime.endpointPath,
      logPath: runtime.logPath,
      home: paths.home,
      pid: endpoint.pid,
      port: endpoint.port,
      updatedAt: endpoint.updatedAt,
      processAlive: true,
      message: "Daemon endpoint belongs to a different PRO_CLI_HOME.",
    };
  }
  const client = new DaemonClient(endpoint);
  try {
    await client.health();
    return {
      state: "running",
      endpointPath: runtime.endpointPath,
      logPath: endpoint.logPath,
      home: paths.home,
      pid: endpoint.pid,
      port: endpoint.port,
      updatedAt: endpoint.updatedAt,
      processAlive: true,
    };
  } catch (error) {
    const processAlive = isProcessAlive(endpoint.pid);
    if (!processAlive) {
      return {
        state: "stopped",
        endpointPath: runtime.endpointPath,
        logPath: endpoint.logPath,
        home: paths.home,
        pid: endpoint.pid,
        port: endpoint.port,
        updatedAt: endpoint.updatedAt,
        processAlive,
        message: "Daemon pid is not running.",
      };
    }
    return {
      state: "stale",
      endpointPath: runtime.endpointPath,
      logPath: endpoint.logPath,
      home: paths.home,
      pid: endpoint.pid,
      port: endpoint.port,
      updatedAt: endpoint.updatedAt,
      processAlive,
      message: error instanceof Error ? error.message : String(error),
    };
  }
}

export async function stopDaemon(paths: RuntimePaths): Promise<DaemonStatus> {
  const connected = await connectDaemon(paths);
  if (connected) {
    await connected.client.shutdown();
    await sleep(100);
  } else {
    const status = await getDaemonStatus(paths);
    if (status.pid && status.processAlive) {
      try {
        process.kill(status.pid, "SIGTERM");
      } catch {
        // The process may have exited between status and stop.
      }
    }
  }
  return getDaemonStatus(paths);
}

export async function runDaemonServer(
  paths: RuntimePaths,
  options: { port?: number; pollMs?: number; idleTimeoutMs?: number } = {},
): Promise<void> {
  await ensurePrivateDir(paths.home);
  const runtime = daemonRuntimePaths(paths.home);
  await ensurePrivateDir(runtime.dir);
  const token = randomUUID().replace(/-/g, "") + randomUUID().replace(/-/g, "");
  const startedAt = new Date().toISOString();
  const store = await JobStore.open(paths.dbPath);
  let stopping = false;
  let pumping = false;
  let lastActivityAt = Date.now();

  const server = Bun.serve({
    hostname: "127.0.0.1",
    port: options.port ?? 0,
    async fetch(request) {
      if (!isAuthorized(request, token)) return jsonError(new ProError("DAEMON_UNAUTHORIZED", "Invalid daemon token."), 401);
      try {
        lastActivityAt = Date.now();
        const response = await routeDaemonRequest(request, store, paths, {
          pumpQueue,
          shutdown: () => {
            setTimeout(() => void shutdown(), 10);
          },
        });
        return jsonOk(response);
      } catch (error) {
        const proError = toProError(error);
        return jsonError(proError, httpStatusForError(proError));
      }
    },
  });

  async function writeEndpoint(): Promise<void> {
    const endpoint: DaemonEndpoint = {
      version: 1,
      pid: process.pid,
      port: server.port ?? 0,
      token,
      home: paths.home,
      startedAt,
      updatedAt: new Date().toISOString(),
      logPath: runtime.logPath,
    };
    await writePrivateFile(runtime.endpointPath, `${JSON.stringify(endpoint, null, 2)}\n`);
  }

  async function pumpQueue(): Promise<void> {
    if (pumping || stopping) return;
    pumping = true;
    try {
      while (!stopping) {
        const claimed = store.claimNextQueued();
        if (!claimed) return;
        await executeClaimedJob(store, claimed, paths);
        lastActivityAt = Date.now();
      }
    } finally {
      pumping = false;
    }
  }

  async function shutdown(): Promise<void> {
    if (stopping) return;
    stopping = true;
    clearInterval(heartbeat);
    clearInterval(queuePoll);
    clearInterval(idleCheck);
    store.close();
    await unlink(runtime.endpointPath).catch(() => undefined);
    server.stop();
  }

  await writeEndpoint();
  const heartbeat = setInterval(() => void writeEndpoint(), HEARTBEAT_MS);
  const queuePoll = setInterval(() => void pumpQueue(), options.pollMs ?? DEFAULT_DAEMON_POLL_MS);
  const idleCheck = setInterval(() => {
    if (!options.idleTimeoutMs || pumping) return;
    if (Date.now() - lastActivityAt >= options.idleTimeoutMs) void shutdown();
  }, 1_000);
  process.once("SIGTERM", () => void shutdown());
  process.once("SIGINT", () => void shutdown());
  await pumpQueue();
  while (!stopping) await sleep(1_000);
}

async function connectDaemon(paths: RuntimePaths): Promise<{ client: DaemonClient; status: DaemonStatus } | null> {
  const runtime = daemonRuntimePaths(paths.home);
  const endpoint = await readEndpoint(runtime.endpointPath).catch(() => null);
  if (!endpoint || endpoint.home !== paths.home) return null;
  const client = new DaemonClient(endpoint);
  try {
    await client.health();
    return {
      client,
      status: {
        state: "running",
        endpointPath: runtime.endpointPath,
        logPath: endpoint.logPath,
        home: paths.home,
        pid: endpoint.pid,
        port: endpoint.port,
        updatedAt: endpoint.updatedAt,
        processAlive: true,
      },
    };
  } catch {
    return null;
  }
}

async function startDaemonProcess(paths: RuntimePaths, io: CliIO): Promise<void> {
  const runtime = daemonRuntimePaths(paths.home);
  await ensurePrivateDir(runtime.dir);
  const logFd = openSync(runtime.logPath, "a", 0o600);
  const cliPath = new URL("./cli.ts", import.meta.url).pathname;
  const env = { ...process.env };
  for (const [key, value] of Object.entries(io.env)) {
    if (value !== undefined) env[key] = value;
  }
  const child = spawn(process.execPath, [cliPath, "daemon", "serve", "--json"], {
    cwd: io.cwd,
    detached: true,
    env,
    stdio: ["ignore", logFd, logFd],
  });
  child.unref();
}

async function routeDaemonRequest(
  request: Request,
  store: JobStore,
  paths: RuntimePaths,
  control: { pumpQueue: () => Promise<void>; shutdown: () => void },
): Promise<Record<string, unknown>> {
  const url = new URL(request.url);
  const parts = url.pathname.split("/").filter(Boolean);
  if (request.method === "GET" && url.pathname === "/health") {
    return { state: "running", pid: process.pid, home: paths.home };
  }
  if (request.method === "POST" && url.pathname === "/shutdown") {
    control.shutdown();
    return { state: "stopping" };
  }
  if (request.method === "POST" && url.pathname === "/jobs") {
    const input = createJobInputFromPayload(await readJsonBody(request));
    const job = store.create(input);
    queueMicrotask(() => void control.pumpQueue());
    return { job, daemon: { accepted: true } };
  }
  if (request.method === "GET" && url.pathname === "/jobs") {
    const limit = parseLimit(url.searchParams.get("limit"));
    return { jobs: store.list(limit) };
  }
  if (parts[0] !== "jobs" || !parts[1]) {
    throw new ProError("DAEMON_ROUTE_NOT_FOUND", `No daemon route for ${request.method} ${url.pathname}.`, {
      exitCode: EXIT.notFound,
    });
  }

  const jobId = decodeURIComponent(parts[1]);
  if (request.method === "GET" && parts.length === 2) {
    return { job: redactJob(store.get(jobId)) };
  }
  if (request.method === "GET" && parts[2] === "result") {
    const job = store.get(jobId);
    if (job.status !== "succeeded") {
      throw new ProError("JOB_NOT_READY", `Job ${jobId} is ${job.status}.`, {
        exitCode: EXIT.notFound,
        suggestions: ["Run pro-cli job wait <job-id> or pro-cli job status <job-id>."],
      });
    }
    return { jobId, result: job.result };
  }
  if (request.method === "POST" && parts[2] === "cancel") {
    return { job: store.cancel(jobId) };
  }
  if (request.method === "POST" && parts[2] === "wait") {
    const body = await readJsonBody(request);
    queueMicrotask(() => void control.pumpQueue());
    const outcome = await waitForJob(
      store,
      jobId,
      numberFromPayload(body.timeoutMs, 0),
      numberFromPayload(body.pollMs, DEFAULT_DAEMON_POLL_MS),
    );
    if (outcome.timedOut && body.softTimeout !== true) throw waitTimeoutError(outcome);
    return waitPayload(outcome);
  }

  throw new ProError("DAEMON_ROUTE_NOT_FOUND", `No daemon route for ${request.method} ${url.pathname}.`, {
    exitCode: EXIT.notFound,
  });
}

function daemonRuntimePaths(home: string): DaemonRuntimePaths {
  const uid = typeof process.getuid === "function" ? process.getuid() : "user";
  const hash = createHash("sha256").update(home).digest("hex").slice(0, 12);
  const root = process.platform === "win32" ? tmpdir() : "/tmp";
  const dir = join(root, `pro-cli-${uid}-${hash}`);
  return {
    dir,
    endpointPath: join(dir, "daemon.json"),
    logPath: join(dir, "daemon.log"),
  };
}

async function readEndpoint(path: string): Promise<DaemonEndpoint> {
  const raw = await readFile(path, "utf8");
  return JSON.parse(raw) as DaemonEndpoint;
}

function isAuthorized(request: Request, token: string): boolean {
  return request.headers.get("authorization") === `Bearer ${token}`;
}

function isProcessAlive(pid: number): boolean {
  try {
    process.kill(pid, 0);
    return true;
  } catch {
    return false;
  }
}

async function readJsonBody(request: Request): Promise<Record<string, unknown>> {
  const raw = await request.text();
  if (!raw.trim()) return {};
  const parsed = JSON.parse(raw) as unknown;
  if (!parsed || typeof parsed !== "object" || Array.isArray(parsed)) {
    throw new ProError("INVALID_JSON", "Daemon request body must be a JSON object.", {
      exitCode: EXIT.invalidArgs,
    });
  }
  return parsed as Record<string, unknown>;
}

function createJobInputFromPayload(payload: Record<string, unknown>): CreateJobInput {
  if (
    typeof payload.prompt !== "string" ||
    typeof payload.model !== "string" ||
    typeof payload.reasoning !== "string" ||
    !isRecord(payload.options)
  ) {
    throw new ProError("INVALID_DAEMON_JOB", "Daemon job payload is missing prompt, model, reasoning, or options.", {
      exitCode: EXIT.invalidArgs,
    });
  }
  return {
    prompt: payload.prompt,
    model: payload.model,
    reasoning: payload.reasoning,
    options: payload.options,
  };
}

function parseLimit(raw: string | null): number {
  const limit = Number(raw ?? "20");
  if (!Number.isInteger(limit) || limit < 1 || limit > 200) {
    throw new ProError("INVALID_ARGS", "Invalid --limit.", {
      exitCode: EXIT.invalidArgs,
      suggestions: ["Use --limit between 1 and 200."],
    });
  }
  return limit;
}

function numberFromPayload(value: unknown, fallback: number): number {
  return typeof value === "number" && Number.isFinite(value) ? value : fallback;
}

function waitPayload(outcome: JobWaitOutcome): Record<string, unknown> {
  return {
    job: redactJob(outcome.job),
    wait: {
      status: outcome.status,
      timedOut: outcome.timedOut,
      elapsedMs: outcome.elapsedMs,
      timeoutMs: outcome.timeoutMs,
      pollMs: outcome.pollMs,
    },
    ...(outcome.timedOut
      ? {
          next: {
            command: `pro-cli job wait ${outcome.job.id} --json`,
            reason: "The job is still running. Wait again without a timeout, or use another soft timeout poll.",
          },
        }
      : {}),
  };
}

function isRecord(value: unknown): value is Record<string, unknown> {
  return Boolean(value) && typeof value === "object" && !Array.isArray(value);
}

function jsonOk(data: Record<string, unknown>): Response {
  return Response.json({ ok: true, data });
}

function jsonError(error: ProError, status: number): Response {
  return Response.json({ ok: false, error: error.toPayload() }, { status });
}

function httpStatusForError(error: ProError): number {
  if (error.exitCode === EXIT.invalidArgs) return 400;
  if (error.exitCode === EXIT.notFound) return 404;
  if (error.exitCode === EXIT.auth) return 401;
  if (error.exitCode === EXIT.timeout) return 408;
  return 500;
}

function exitCodeForStatus(status: number): ExitCode {
  if (status === 400) return EXIT.invalidArgs;
  if (status === 401 || status === 403) return EXIT.auth;
  if (status === 404) return EXIT.notFound;
  if (status === 408) return EXIT.timeout;
  if (status >= 500) return EXIT.internal;
  return EXIT.network;
}

async function sleep(ms: number): Promise<void> {
  await new Promise((resolve) => setTimeout(resolve, ms));
}
`````

## File: src/defaults.ts
`````typescript
export const DEFAULT_MODEL = "gpt-5-5-pro";
export const DEFAULT_REASONING = "standard";
export const REASONING_LEVELS = ["min", "standard", "extended", "max"] as const;

export type ReasoningLevel = (typeof REASONING_LEVELS)[number];

export function isReasoningLevel(value: string): value is ReasoningLevel {
  return (REASONING_LEVELS as readonly string[]).includes(value);
}
`````

## File: src/errors.ts
`````typescript
export const EXIT = {
  success: 0,
  notFound: 1,
  invalidArgs: 2,
  auth: 3,
  upstream: 4,
  network: 5,
  timeout: 6,
  internal: 7,
} as const;

export type ExitCode = (typeof EXIT)[keyof typeof EXIT];

export interface ErrorPayload {
  code: string;
  message: string;
  suggestions: string[];
  details?: Record<string, unknown>;
}

export class ProError extends Error {
  readonly code: string;
  readonly suggestions: string[];
  readonly exitCode: ExitCode;
  readonly details?: Record<string, unknown>;

  constructor(
    code: string,
    message: string,
    options: {
      suggestions?: string[];
      exitCode?: ExitCode;
      details?: Record<string, unknown>;
      cause?: unknown;
    } = {},
  ) {
    super(message, { cause: options.cause });
    this.name = "ProError";
    this.code = code;
    this.suggestions = options.suggestions ?? [];
    this.exitCode = options.exitCode ?? EXIT.internal;
    this.details = options.details;
  }

  toPayload(): ErrorPayload {
    return {
      code: this.code,
      message: this.message,
      suggestions: this.suggestions,
      ...(this.details ? { details: this.details } : {}),
    };
  }
}

export function toProError(error: unknown): ProError {
  if (error instanceof ProError) return error;
  if (error instanceof Error) {
    return new ProError("INTERNAL_ERROR", error.message, {
      exitCode: EXIT.internal,
      suggestions: ["Run with --json and inspect the structured error."],
      cause: error,
    });
  }
  return new ProError("INTERNAL_ERROR", String(error), {
    exitCode: EXIT.internal,
    suggestions: ["Run with --json and inspect the structured error."],
  });
}
`````

## File: src/executor.ts
`````typescript
import { randomUUID } from "node:crypto";
import type { RuntimePaths } from "./config";
import { EXIT, ProError, toProError } from "./errors";
import { JobStore, redactJob, type JobRecord, type JobStatus, type LimitsObservation } from "./jobs";
import { runChatGptJob } from "./transport";

export async function executeClaimedJob(
  store: JobStore,
  job: JobRecord,
  paths: RuntimePaths,
): Promise<Record<string, unknown>> {
  const observations: LimitsObservation[] = [];
  try {
    const result = await runChatGptJob(job, {
      sessionTokenPath: paths.sessionTokenPath,
      cdpBase: stringFromOption(job.options.cdpBase),
      timeoutMs: numberFromOption(job.options.timeoutMs),
      retries: numberFromOption(job.options.retries),
      retryDelayMs: numberFromOption(job.options.retryDelayMs),
      onLimits: (entries) => observations.push(...entries),
    });
    if (observations.length > 0) store.recordLimits(observations, job.id);
    const completed = store.markSucceeded(job.id, result);
    return { job: redactJob(completed), result };
  } catch (error) {
    const proError = toProError(error);
    return {
      job: redactJob(store.markFailed(job.id, proError)),
      error: proError.toPayload(),
    };
  }
}

export async function executeEphemeralJob(
  job: JobRecord,
  paths: RuntimePaths,
): Promise<Record<string, unknown>> {
  const observations: LimitsObservation[] = [];
  try {
    const result = await runChatGptJob(job, {
      sessionTokenPath: paths.sessionTokenPath,
      cdpBase: stringFromOption(job.options.cdpBase),
      timeoutMs: numberFromOption(job.options.timeoutMs),
      retries: numberFromOption(job.options.retries),
      retryDelayMs: numberFromOption(job.options.retryDelayMs),
      onLimits: (entries) => observations.push(...entries),
    });
    if (observations.length > 0) await persistLimits(paths.dbPath, observations, job.id);
    return {
      job: redactJob({ ...job, status: "succeeded", result, updatedAt: new Date().toISOString() }),
      result,
    };
  } catch (error) {
    const proError = toProError(error);
    return {
      job: redactJob({
        ...job,
        status: "failed",
        error: JSON.stringify(proError.toPayload()),
        updatedAt: new Date().toISOString(),
      }),
      error: proError.toPayload(),
      exitCode: proError.exitCode,
    };
  }
}

async function persistLimits(
  dbPath: string,
  observations: LimitsObservation[],
  jobId: string,
): Promise<void> {
  const store = await JobStore.open(dbPath);
  try {
    store.recordLimits(observations, jobId);
  } finally {
    store.close();
  }
}

export function buildEphemeralJob(input: {
  prompt: string;
  model: string;
  reasoning: string;
  options: Record<string, unknown>;
}): JobRecord {
  const now = new Date().toISOString();
  return {
    id: `ask_${randomUUID()}`,
    status: "running",
    prompt: input.prompt,
    model: input.model,
    reasoning: input.reasoning,
    options: input.options,
    result: null,
    error: null,
    createdAt: now,
    updatedAt: now,
  };
}

export async function waitForTerminalJob(
  store: JobStore,
  jobId: string,
  timeoutMs: number,
  pollMs: number,
): Promise<JobRecord> {
  const outcome = await waitForJob(store, jobId, timeoutMs, pollMs);
  if (outcome.timedOut) throw waitTimeoutError(outcome);
  return outcome.job;
}

export interface JobWaitOutcome {
  job: JobRecord;
  status: JobStatus;
  timedOut: boolean;
  elapsedMs: number;
  timeoutMs: number;
  pollMs: number;
}

export async function waitForJob(
  store: JobStore,
  jobId: string,
  timeoutMs: number,
  pollMs: number,
): Promise<JobWaitOutcome> {
  const start = Date.now();
  while (true) {
    const job = store.get(jobId);
    const elapsedMs = Date.now() - start;
    if (job.status !== "queued" && job.status !== "running") {
      return {
        job,
        status: job.status,
        timedOut: false,
        elapsedMs,
        timeoutMs,
        pollMs,
      };
    }
    if (timeoutMs > 0 && elapsedMs >= timeoutMs) {
      return {
        job,
        status: job.status,
        timedOut: true,
        elapsedMs,
        timeoutMs,
        pollMs,
      };
    }
    await sleep(pollMs);
  }
}

export function waitTimeoutError(outcome: JobWaitOutcome): ProError {
  return new ProError(
    "WAIT_TIMEOUT",
    `Job ${outcome.job.id} is still ${outcome.status} after ${formatDuration(outcome.elapsedMs)}.`,
    {
      exitCode: EXIT.timeout,
      suggestions: [
        `Run pro-cli job wait ${outcome.job.id} --json without --wait-timeout.`,
        `Run pro-cli job wait ${outcome.job.id} --soft-timeout ${outcome.timeoutMs} --json to poll without an error exit.`,
        `Use pro-cli job cancel ${outcome.job.id} --json if this job is stale.`,
      ],
      details: {
        job: redactJob(outcome.job),
        status: outcome.status,
        elapsedMs: outcome.elapsedMs,
        timeoutMs: outcome.timeoutMs,
        pollMs: outcome.pollMs,
      },
    },
  );
}

function numberFromOption(value: unknown): number | undefined {
  return typeof value === "number" ? value : undefined;
}

function stringFromOption(value: unknown): string | undefined {
  return typeof value === "string" ? value : undefined;
}

async function sleep(ms: number): Promise<void> {
  await new Promise((resolve) => setTimeout(resolve, ms));
}

function formatDuration(ms: number): string {
  if (ms < 1_000) return `${ms}ms`;
  const seconds = Math.round(ms / 1_000);
  if (seconds < 60) return `${seconds}s`;
  const minutes = Math.floor(seconds / 60);
  const remainingSeconds = seconds % 60;
  return remainingSeconds > 0 ? `${minutes}m ${remainingSeconds}s` : `${minutes}m`;
}
`````

## File: src/jobs.ts
`````typescript
import { Database } from "bun:sqlite";
import { randomUUID } from "node:crypto";
import { dirname } from "node:path";
import { EXIT, ProError } from "./errors";
import { ensurePrivateDir } from "./config";

export type JobStatus = "queued" | "running" | "succeeded" | "failed" | "cancelled";

export interface CreateJobInput {
  prompt: string;
  model: string;
  reasoning: string;
  options: Record<string, unknown>;
}

export interface JobRecord {
  id: string;
  status: JobStatus;
  prompt: string;
  model: string;
  reasoning: string;
  options: Record<string, unknown>;
  result: string | null;
  error: string | null;
  createdAt: string;
  updatedAt: string;
}

export interface LimitsObservation {
  feature_name: string;
  remaining: number;
  reset_after: string | null;
}

export interface LimitsSnapshot {
  featureName: string;
  remaining: number;
  resetAfter: string | null;
  observedAt: string;
  jobId: string | null;
}

interface JobRow {
  id: string;
  status: JobStatus;
  prompt: string;
  model: string;
  reasoning: string;
  options_json: string;
  result: string | null;
  error_json: string | null;
  created_at: string;
  updated_at: string;
}

export class JobStore {
  private constructor(private readonly db: Database) {}

  static async open(path: string): Promise<JobStore> {
    await ensurePrivateDir(dirname(path));
    const db = new Database(path, { create: true });
    db.run("PRAGMA journal_mode = WAL");
    db.run("PRAGMA foreign_keys = ON");
    db.run(`
      CREATE TABLE IF NOT EXISTS jobs (
        id TEXT PRIMARY KEY,
        status TEXT NOT NULL,
        prompt TEXT NOT NULL,
        model TEXT NOT NULL,
        reasoning TEXT NOT NULL,
        options_json TEXT NOT NULL,
        result TEXT,
        error_json TEXT,
        created_at TEXT NOT NULL,
        updated_at TEXT NOT NULL
      )
    `);
    db.run(`
      CREATE TABLE IF NOT EXISTS limits_observations (
        feature_name TEXT NOT NULL,
        remaining REAL NOT NULL,
        reset_after TEXT,
        observed_at TEXT NOT NULL,
        job_id TEXT,
        PRIMARY KEY (feature_name, observed_at)
      )
    `);
    return new JobStore(db);
  }

  create(input: CreateJobInput): ReturnType<typeof redactJob> {
    const now = new Date().toISOString();
    const job: JobRecord = {
      id: `job_${randomUUID()}`,
      status: "queued",
      prompt: input.prompt,
      model: input.model,
      reasoning: input.reasoning,
      options: input.options,
      result: null,
      error: null,
      createdAt: now,
      updatedAt: now,
    };
    this.db
      .query(
        `INSERT INTO jobs
          (id, status, prompt, model, reasoning, options_json, result, error_json, created_at, updated_at)
         VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
      )
      .run(
        job.id,
        job.status,
        job.prompt,
        job.model,
        job.reasoning,
        JSON.stringify(job.options),
        job.result,
        job.error,
        job.createdAt,
        job.updatedAt,
      );
    return redactJob(job);
  }

  get(id: string): JobRecord {
    const row = this.db.query("SELECT * FROM jobs WHERE id = ?").get(id) as JobRow | null;
    if (!row) {
      throw new ProError("JOB_NOT_FOUND", `No job exists for ${id}.`, {
        exitCode: EXIT.notFound,
        suggestions: ["Run pro-cli job list --json to list recent jobs."],
      });
    }
    return rowToJob(row);
  }

  list(limit: number): JobRecord[] {
    const rows = this.db
      .query("SELECT * FROM jobs ORDER BY created_at DESC LIMIT ?")
      .all(limit) as JobRow[];
    return rows.map(rowToJob).map(redactJob);
  }

  cancel(id: string): JobRecord {
    const job = this.get(id);
    if (job.status === "succeeded" || job.status === "failed" || job.status === "cancelled") {
      return redactJob(job);
    }
    this.update(id, {
      status: "cancelled",
      error: JSON.stringify({
        code: "CANCELLED",
        message: "Job was cancelled before completion.",
      }),
    });
    return redactJob(this.get(id));
  }

  markRunning(id: string): JobRecord {
    const now = new Date().toISOString();
    this.db
      .query("UPDATE jobs SET status = ?, updated_at = ? WHERE id = ? AND status = ?")
      .run("running", now, id, "queued");
    return this.get(id);
  }

  claimQueued(id: string): JobRecord | null {
    const now = new Date().toISOString();
    const result = this.db
      .query("UPDATE jobs SET status = ?, updated_at = ? WHERE id = ? AND status = ?")
      .run("running", now, id, "queued") as { changes?: number };
    return result.changes && result.changes > 0 ? this.get(id) : null;
  }

  claimNextQueued(): JobRecord | null {
    const row = this.db
      .query("SELECT id FROM jobs WHERE status = ? ORDER BY created_at ASC LIMIT 1")
      .get("queued") as { id: string } | null;
    return row ? this.claimQueued(row.id) : null;
  }

  markSucceeded(id: string, result: string): JobRecord {
    this.finishRunning(id, { status: "succeeded", result, error: null });
    return this.get(id);
  }

  markFailed(id: string, error: ProError): JobRecord {
    this.finishRunning(id, {
      status: "failed",
      error: JSON.stringify(error.toPayload()),
    });
    return this.get(id);
  }

  recordLimits(observations: LimitsObservation[], jobId: string | null): void {
    if (observations.length === 0) return;
    const observedAt = new Date().toISOString();
    const insert = this.db.query(
      `INSERT OR REPLACE INTO limits_observations
        (feature_name, remaining, reset_after, observed_at, job_id)
       VALUES (?, ?, ?, ?, ?)`,
    );
    for (const observation of observations) {
      insert.run(
        observation.feature_name,
        observation.remaining,
        observation.reset_after ?? null,
        observedAt,
        jobId,
      );
    }
  }

  latestLimits(): LimitsSnapshot[] {
    const rows = this.db
      .query(
        `SELECT feature_name, remaining, reset_after, observed_at, job_id
         FROM limits_observations
         WHERE (feature_name, observed_at) IN (
           SELECT feature_name, MAX(observed_at) FROM limits_observations GROUP BY feature_name
         )
         ORDER BY feature_name ASC`,
      )
      .all() as Array<{
        feature_name: string;
        remaining: number;
        reset_after: string | null;
        observed_at: string;
        job_id: string | null;
      }>;
    return rows.map((row) => ({
      featureName: row.feature_name,
      remaining: row.remaining,
      resetAfter: row.reset_after,
      observedAt: row.observed_at,
      jobId: row.job_id,
    }));
  }

  close(): void {
    this.db.close();
  }

  private update(
    id: string,
    patch: { status?: JobStatus; result?: string | null; error?: string | null },
  ): void {
    const job = this.get(id);
    const next = {
      status: patch.status ?? job.status,
      result: patch.result === undefined ? job.result : patch.result,
      error: patch.error === undefined ? job.error : patch.error,
      updatedAt: new Date().toISOString(),
    };
    this.db
      .query("UPDATE jobs SET status = ?, result = ?, error_json = ?, updated_at = ? WHERE id = ?")
      .run(next.status, next.result, next.error, next.updatedAt, id);
  }

  private finishRunning(
    id: string,
    patch: { status: "succeeded" | "failed"; result?: string | null; error?: string | null },
  ): void {
    const job = this.get(id);
    const next = {
      status: patch.status,
      result: patch.result === undefined ? job.result : patch.result,
      error: patch.error === undefined ? job.error : patch.error,
      updatedAt: new Date().toISOString(),
    };
    this.db
      .query(
        "UPDATE jobs SET status = ?, result = ?, error_json = ?, updated_at = ? WHERE id = ? AND status = ?",
      )
      .run(next.status, next.result, next.error, next.updatedAt, id, "running");
  }
}

export function redactJob(
  job: JobRecord,
): Omit<JobRecord, "result"> & { result: null; promptPreview: string; resultPreview?: string; hasResult: boolean } {
  return {
    ...job,
    prompt: "",
    result: null,
    promptPreview: compact(job.prompt, 160),
    ...(job.result !== null ? { resultPreview: compact(job.result, 240) } : {}),
    hasResult: job.result !== null,
  };
}

function rowToJob(row: JobRow): JobRecord {
  return {
    id: row.id,
    status: row.status,
    prompt: row.prompt,
    model: row.model,
    reasoning: row.reasoning,
    options: JSON.parse(row.options_json) as Record<string, unknown>,
    result: row.result,
    error: row.error_json,
    createdAt: row.created_at,
    updatedAt: row.updated_at,
  };
}

function compact(text: string, max: number): string {
  const oneLine = text.replace(/\s+/g, " ").trim();
  if (oneLine.length <= max) return oneLine;
  return `${oneLine.slice(0, max - 1)}…`;
}
`````

## File: src/limits.ts
`````typescript
import { evaluateInCdpPage, recoverCookieBloatInCdp } from "./cdp";
import { chatGptOrigins } from "./cookies";
import { EXIT, ProError } from "./errors";

const DEFAULT_CDP_BASE = "http://127.0.0.1:9222";

export interface AccountSummary {
  planType: string | null;
  subscriptionPlan: string | null;
  hasActiveSubscription: boolean;
  expiresAt: string | null;
  renewsAt: string | null;
  cancelsAt: string | null;
  billingPeriod: string | null;
  willRenew: boolean | null;
  features: string[];
}

interface RawFetchResult {
  ok: boolean;
  status: number;
  body: string;
  code?: "CHATGPT_PAGE_MISSING" | "CHATGPT_PAGE_LOGGED_OUT" | "CHATGPT_PROBE_FAILED";
}

export async function fetchAccountSummary(cdpBase?: string): Promise<AccountSummary> {
  const resolvedCdpBase = cdpBase ?? DEFAULT_CDP_BASE;
  const expression = `(${async function pageFetch(): Promise<RawFetchResult> {
    if (location.origin !== "https://chatgpt.com") {
      return { ok: false, status: 0, code: "CHATGPT_PAGE_MISSING", body: location.href };
    }
    const sessionResponse = await fetch("https://chatgpt.com/api/auth/session", { credentials: "include", referrerPolicy: "no-referrer" });
    const session = (await sessionResponse.json().catch(() => null)) as { accessToken?: unknown } | null;
    if (!sessionResponse.ok && sessionResponse.status !== 401) {
      return {
        ok: false,
        status: sessionResponse.status,
        code: "CHATGPT_PROBE_FAILED",
        body: `ChatGPT auth session probe returned HTTP ${sessionResponse.status}.`,
      };
    }
    if (typeof session?.accessToken !== "string" || !session.accessToken) {
      return { ok: false, status: sessionResponse.status, code: "CHATGPT_PAGE_LOGGED_OUT", body: "" };
    }
    const response = await fetch("https://chatgpt.com/backend-api/accounts/check/v4-2023-04-27", {
      method: "GET",
      credentials: "include",
      headers: {
        authorization: `Bearer ${session.accessToken}`,
        accept: "application/json",
      },
    });
    const body = await response.text();
    return { ok: response.ok, status: response.status, body };
  }})()`;

  let raw = await evaluateInCdpPage<RawFetchResult>(resolvedCdpBase, expression, 30_000);
  if (shouldRecoverCookieBloat(raw)) {
    const recovered = await recoverCookieBloatInCdp(resolvedCdpBase, chatGptOrigins(), 30_000).catch(() => null);
    if (recovered?.deleted) {
      raw = await evaluateInCdpPage<RawFetchResult>(resolvedCdpBase, expression, 30_000);
    }
  }
  if (raw.code === "CHATGPT_PAGE_MISSING") {
    throw new ProError("CHATGPT_PAGE_MISSING", "No logged-in ChatGPT page is available over CDP.", {
      exitCode: EXIT.auth,
      suggestions: [
        "Open the Chrome command from pro-cli auth command.",
        "Confirm the CDP Chrome window is on https://chatgpt.com/ and logged in.",
      ],
    });
  }
  if (raw.code === "CHATGPT_PAGE_LOGGED_OUT") {
    throw new ProError("CHATGPT_PAGE_LOGGED_OUT", "The ChatGPT CDP page is not logged in.", {
      exitCode: EXIT.auth,
      suggestions: ["Sign in to ChatGPT, then run pro-cli auth capture."],
    });
  }
  if (raw.code === "CHATGPT_PROBE_FAILED") {
    const status = raw.status;
    const suggestions =
      status === 431
        ? [
            "HTTP 431 indicates oversize request headers; the CDP Chrome profile likely has stale cookie buildup.",
            "Sign out of ChatGPT in the CDP window, sign back in, then run pro-cli auth capture.",
          ]
        : [
            `The ChatGPT auth session probe failed with HTTP ${status}; cannot determine login state.`,
            "Reload the CDP ChatGPT tab and retry.",
          ];
    throw new ProError(
      "CHATGPT_PROBE_FAILED",
      `Could not determine ChatGPT login state from the CDP page (HTTP ${status}).`,
      { exitCode: EXIT.auth, suggestions, details: { status } },
    );
  }
  if (!raw.ok) {
    throw new ProError("UPSTREAM_REJECTED", `accounts/check returned HTTP ${raw.status}.`, {
      exitCode: EXIT.upstream,
      details: { preview: raw.body.slice(0, 240) },
    });
  }
  return summarizeAccountResponse(raw.body);
}

function shouldRecoverCookieBloat(raw: RawFetchResult): boolean {
  return raw.status === 431 || raw.body.includes("chrome-error://chromewebdata/");
}

export function summarizeAccountResponse(body: string): AccountSummary {
  let parsed: unknown;
  try {
    parsed = JSON.parse(body);
  } catch (error) {
    throw new ProError("ACCOUNT_PARSE_FAILED", "Could not parse accounts/check response.", {
      exitCode: EXIT.upstream,
      cause: error,
    });
  }
  const accounts = isRecord(parsed) ? parsed.accounts : null;
  if (!isRecord(accounts)) {
    return emptySummary();
  }
  const ordering = isRecord(parsed) && Array.isArray((parsed as { account_ordering?: unknown }).account_ordering)
    ? ((parsed as { account_ordering: unknown[] }).account_ordering as unknown[])
    : [];
  const orderedKey = ordering.find((id): id is string => typeof id === "string" && id in accounts);
  const fallbackKey = Object.keys(accounts).find((key) => key !== "default") ?? "default";
  const key = orderedKey ?? fallbackKey;
  const account = accounts[key];
  if (!isRecord(account)) return emptySummary();

  const accountInfo = isRecord(account.account) ? account.account : {};
  const entitlement = isRecord(account.entitlement) ? account.entitlement : {};
  const lastSub = isRecord(account.last_active_subscription) ? account.last_active_subscription : {};
  const features = Array.isArray(account.features)
    ? (account.features as unknown[]).filter((value): value is string => typeof value === "string")
    : [];

  return {
    planType: stringOrNull(accountInfo.plan_type),
    subscriptionPlan: stringOrNull(entitlement.subscription_plan),
    hasActiveSubscription: entitlement.has_active_subscription === true,
    expiresAt: stringOrNull(entitlement.expires_at),
    renewsAt: stringOrNull(entitlement.renews_at),
    cancelsAt: stringOrNull(entitlement.cancels_at),
    billingPeriod: stringOrNull(entitlement.billing_period),
    willRenew: typeof lastSub.will_renew === "boolean" ? lastSub.will_renew : null,
    features,
  };
}

function emptySummary(): AccountSummary {
  return {
    planType: null,
    subscriptionPlan: null,
    hasActiveSubscription: false,
    expiresAt: null,
    renewsAt: null,
    cancelsAt: null,
    billingPeriod: null,
    willRenew: null,
    features: [],
  };
}

function stringOrNull(value: unknown): string | null {
  return typeof value === "string" ? value : null;
}

function isRecord(value: unknown): value is Record<string, unknown> {
  return Boolean(value) && typeof value === "object" && !Array.isArray(value);
}
`````

## File: src/models.ts
`````typescript
import { DEFAULT_MODEL, REASONING_LEVELS } from "./defaults";
import { isTokenFresh, loadSessionToken } from "./session-token";

export interface ModelCapability {
  id: string;
  label: string;
  source: "static-unverified" | "live";
  reasoningLevels: string[];
  default?: boolean;
  maxTokens?: number;
  reasoningType?: string;
  configurableThinkingEffort?: boolean;
  enabledTools?: string[];
}

export interface ModelList {
  models: ModelCapability[];
  source: "static" | "live";
  warning?: string;
  defaultModel?: string;
  chatgptDefaultModel?: string;
  modelPickerVersion?: number;
}

interface LiveModelPayload {
  default_model_slug?: unknown;
  model_picker_version?: unknown;
  models?: unknown;
}

interface LiveModel {
  slug?: unknown;
  title?: unknown;
  max_tokens?: unknown;
  reasoning_type?: unknown;
  configurable_thinking_effort?: unknown;
  thinking_efforts?: unknown;
  enabled_tools?: unknown;
}

const MODEL_ID_ALIASES = new Map<string, string>([
  ["4.5", "gpt-4-5"],
  ["gpt 4.5", "gpt-4-5"],
  ["gpt-4.5", "gpt-4-5"],
  ["deep research", "research"],
  ["deep-research", "research"],
  ["deep_research", "research"],
]);

const STATIC_MODEL_CAPABILITIES: ModelCapability[] = [
  {
    id: "gpt-5-5-pro",
    label: "GPT-5.5 Pro",
    source: "static-unverified",
    reasoningLevels: ["standard", "extended"],
    default: true,
    reasoningType: "pro",
  },
  {
    id: "gpt-5-4-pro",
    label: "GPT-5.4 Pro",
    source: "static-unverified",
    reasoningLevels: ["standard", "extended"],
    reasoningType: "pro",
  },
  {
    id: "gpt-5-5-thinking",
    label: "GPT-5.5 Thinking",
    source: "static-unverified",
    reasoningLevels: [...REASONING_LEVELS],
    reasoningType: "reasoning",
  },
  {
    id: "gpt-4-5",
    label: "GPT-4.5",
    source: "static-unverified",
    reasoningLevels: [],
    reasoningType: "none",
  },
  {
    id: "research",
    label: "Deep Research",
    source: "static-unverified",
    reasoningLevels: ["standard", "extended"],
    reasoningType: "pro",
    configurableThinkingEffort: true,
  },
];

const STATIC_MODEL_BY_ID = new Map(STATIC_MODEL_CAPABILITIES.map((model) => [model.id, model]));

export const NO_REASONING = "none";

export function canonicalModelId(model: string): string {
  const trimmed = model.trim();
  return MODEL_ID_ALIASES.get(trimmed.toLowerCase()) ?? trimmed;
}

export function modelUsesThinkingEffort(model: string): boolean {
  const capability = STATIC_MODEL_BY_ID.get(canonicalModelId(model));
  return capability?.reasoningType !== "none";
}

export function modelRequiresSavedConversation(model: string): boolean {
  return canonicalModelId(model) === "research";
}

export async function listModels(options: { sessionTokenPath: string }): Promise<ModelList> {
  const session = await loadSessionToken(options.sessionTokenPath).catch(() => null);
  if (!session) {
    return listStaticModels("No captured ChatGPT session token is available.");
  }
  if (!isTokenFresh(session)) {
    return listStaticModels("The captured ChatGPT session token is expired.");
  }
  if (!session.accountId) {
    return listStaticModels("The captured ChatGPT session token has no account id.");
  }

  try {
    const response = await fetch("https://chatgpt.com/backend-api/models", {
      headers: {
        authorization: `Bearer ${session.accessToken}`,
        "chatgpt-account-id": session.accountId,
        accept: "application/json",
        origin: "https://chatgpt.com",
        referer: "https://chatgpt.com/",
        "user-agent": "pro-cli/0.1",
      },
    });

    if (!response.ok) {
      return listStaticModels(`Live model discovery returned HTTP ${response.status}.`);
    }

    const payload = (await response.json()) as LiveModelPayload;
    const liveModels = parseLiveModels(payload);
    if (liveModels.length === 0) {
      return listStaticModels("Live model discovery returned no usable model entries.");
    }

    const chatgptDefaultModel =
      typeof payload.default_model_slug === "string" ? payload.default_model_slug : undefined;
    return {
      source: "live",
      defaultModel: DEFAULT_MODEL,
      ...(chatgptDefaultModel ? { chatgptDefaultModel } : {}),
      ...(typeof payload.model_picker_version === "number"
        ? { modelPickerVersion: payload.model_picker_version }
        : {}),
      models: liveModels.map((model) => ({
        ...model,
        default: model.id === DEFAULT_MODEL,
      })),
    };
  } catch {
    return listStaticModels("Live model discovery failed; using static fallback.");
  }
}

export function listStaticModels(warning?: string): ModelList {
  return {
    source: "static",
    defaultModel: DEFAULT_MODEL,
    models: STATIC_MODEL_CAPABILITIES.map((model) => ({
      ...model,
      reasoningLevels: [...model.reasoningLevels],
      ...(model.enabledTools ? { enabledTools: [...model.enabledTools] } : {}),
    })),
    warning: warning ?? "Live model discovery requires a captured ChatGPT session token.",
  };
}

function parseLiveModels(payload: LiveModelPayload): ModelCapability[] {
  const models = Array.isArray(payload.models) ? (payload.models as LiveModel[]) : [];
  return models
    .map((model) => {
      if (typeof model.slug !== "string" || typeof model.title !== "string") return null;
      const reasoningLevels = parseReasoningLevels(model.thinking_efforts, model.reasoning_type);
      const enabledTools = parseEnabledTools(model.enabled_tools);
      const capability: ModelCapability = {
        id: model.slug,
        label: model.title,
        source: "live",
        reasoningLevels,
        ...(typeof model.max_tokens === "number" ? { maxTokens: model.max_tokens } : {}),
        ...(typeof model.reasoning_type === "string" ? { reasoningType: model.reasoning_type } : {}),
        ...(typeof model.configurable_thinking_effort === "boolean"
          ? { configurableThinkingEffort: model.configurable_thinking_effort }
          : {}),
        ...(enabledTools.length > 0 ? { enabledTools } : {}),
      };
      return applyModelCapabilityOverride(capability);
    })
    .filter((model): model is ModelCapability => model !== null && model.id !== "auto");
}

function applyModelCapabilityOverride(model: ModelCapability): ModelCapability {
  const staticModel = STATIC_MODEL_BY_ID.get(canonicalModelId(model.id));
  if (model.id !== "research" || !staticModel) return model;
  return {
    ...model,
    reasoningLevels: [...staticModel.reasoningLevels],
    reasoningType: staticModel.reasoningType,
    configurableThinkingEffort: staticModel.configurableThinkingEffort,
  };
}

function parseReasoningLevels(value: unknown, reasoningType: unknown): string[] {
  if (reasoningType === "none") return [];
  if (!Array.isArray(value)) {
    if (reasoningType === "pro") return ["standard", "extended"];
    return [...REASONING_LEVELS];
  }
  const levels = value.flatMap((item) => {
    if (typeof item === "string") return [item];
    if (!item || typeof item !== "object") return [];
    const effort = (item as { thinking_effort?: unknown }).thinking_effort;
    return typeof effort === "string" ? [effort] : [];
  });
  return levels.length > 0 ? levels : [...REASONING_LEVELS];
}

function parseEnabledTools(value: unknown): string[] {
  if (!Array.isArray(value)) return [];
  return value.flatMap((item) => {
    if (typeof item === "string") return [item];
    if (!item || typeof item !== "object") return [];
    const tool = item as { id?: unknown; name?: unknown; type?: unknown };
    for (const candidate of [tool.id, tool.name, tool.type]) {
      if (typeof candidate === "string") return [candidate];
    }
    return [];
  });
}
`````

## File: src/odds.ts
`````typescript
import type { RuntimePaths } from "./config";
import { EXIT, ProError } from "./errors";
import { buildEphemeralJob, executeEphemeralJob } from "./executor";

export type AggregateMethod = "mean" | "median" | "trimmed-mean";

export interface OddsRunInput {
  question: string;
  context?: string;
  model: string;
  reasoning: string;
  samples: number;
  aggregate: AggregateMethod;
  allowFifty: boolean;
  parseRetries: number;
  baseRequestOptions: Record<string, unknown>;
  paths: RuntimePaths;
}

export interface OddsSampleAttempt {
  jobId: string;
  raw: string;
  parsed: number | null;
  rejectedFifty: boolean;
}

export interface OddsRunResult {
  probability: number;
  probabilityRaw: number;
  samples: number[];
  aggregate: AggregateMethod;
  parseFailures: number;
  rejectedFifties: number;
  attempts: OddsSampleAttempt[];
  jobIds: string[];
  model: string;
  reasoning: string;
  allowFifty: boolean;
}

export function buildOddsInstructions(allowFifty: boolean): string {
  const fiftyRule = allowFifty
    ? "If you have no information either way, output 50."
    : "Even with limited information you MUST commit to a directional estimate. Pick 49 or 51 over 50. Do NOT output 50.";
  return [
    "You are a calibrated probabilistic forecaster.",
    "You will receive a yes/no question and any supporting context.",
    "Read everything carefully and estimate the probability the question resolves YES.",
    "",
    "OUTPUT RULES (STRICT):",
    "- Output exactly one integer between 0 and 100.",
    "- No tags. No words. No punctuation. No explanation. No reasoning. No markdown. Nothing else.",
    "- Just the integer, on its own. Nothing before. Nothing after.",
    "- 0 means certain NO. 100 means certain YES.",
    `- ${fiftyRule}`,
  ].join("\n");
}

export function buildOddsPrompt(question: string, context?: string): string {
  const trimmedContext = context?.trim();
  const lines: string[] = [];
  if (trimmedContext) {
    lines.push("CONTEXT:", trimmedContext, "");
  }
  lines.push("QUESTION:", question.trim(), "", "Reply with a single integer between 0 and 100. Nothing else.");
  return lines.join("\n");
}

const STRICT_INTEGER = /^\s*(\d{1,3})\s*$/;
const FIRST_INTEGER = /(?:^|[^\w.-])(\d{1,3})(?!\w|[.]\d)/;

export function parseOddsResponse(
  text: string,
  allowFifty: boolean,
): { value: number | null; rejectedFifty: boolean } {
  let candidate: number | null = null;
  const strict = STRICT_INTEGER.exec(text);
  if (strict) {
    const n = Number(strict[1]);
    if (Number.isInteger(n) && n >= 0 && n <= 100) candidate = n;
  } else {
    const loose = FIRST_INTEGER.exec(text);
    if (loose) {
      const n = Number(loose[1]);
      if (Number.isInteger(n) && n >= 0 && n <= 100) candidate = n;
    }
  }
  if (candidate === null) return { value: null, rejectedFifty: false };
  if (!allowFifty && candidate === 50) return { value: null, rejectedFifty: true };
  return { value: candidate, rejectedFifty: false };
}

export function aggregateOdds(values: number[], method: AggregateMethod): number {
  if (values.length === 0) {
    throw new ProError("ODDS_NO_SAMPLES", "No valid samples to aggregate.", {
      exitCode: EXIT.internal,
    });
  }
  if (method === "median") {
    const sorted = [...values].sort((a, b) => a - b);
    const mid = Math.floor(sorted.length / 2);
    return sorted.length % 2 === 0 ? (sorted[mid - 1] + sorted[mid]) / 2 : sorted[mid];
  }
  if (method === "trimmed-mean") {
    if (values.length < 4) return mean(values);
    const sorted = [...values].sort((a, b) => a - b);
    const trim = Math.max(1, Math.floor(values.length * 0.1));
    return mean(sorted.slice(trim, sorted.length - trim));
  }
  return mean(values);
}

function mean(values: number[]): number {
  return values.reduce((acc, v) => acc + v, 0) / values.length;
}

export async function runOdds(input: OddsRunInput): Promise<OddsRunResult> {
  const instructions = buildOddsInstructions(input.allowFifty);
  const prompt = buildOddsPrompt(input.question, input.context);

  const requestOptions: Record<string, unknown> = {
    ...input.baseRequestOptions,
    instructions,
  };
  if (requestOptions.temporary === undefined) requestOptions.temporary = true;

  const attempts: OddsSampleAttempt[] = [];
  const samples: number[] = [];
  const jobIds: string[] = [];
  let parseFailures = 0;
  let rejectedFifties = 0;

  for (let i = 0; i < input.samples; i += 1) {
    let parsed: number | null = null;
    let rejectedFifty = false;
    let lastJobId = "";
    let lastRaw = "";
    for (let attempt = 0; attempt <= input.parseRetries; attempt += 1) {
      const job = buildEphemeralJob({
        prompt,
        model: input.model,
        reasoning: input.reasoning,
        options: requestOptions,
      });
      const outcome = await executeEphemeralJob(job, input.paths);
      const jobObj = isRecord(outcome.job) ? outcome.job : {};
      lastJobId = typeof jobObj.id === "string" ? jobObj.id : job.id;
      lastRaw = typeof outcome.result === "string" ? outcome.result : "";
      if (isRecord(outcome.error)) {
        throw new ProError(
          stringField(outcome.error, "code") ?? "ODDS_UPSTREAM_FAILED",
          stringField(outcome.error, "message") ?? "ChatGPT request failed.",
          {
            exitCode: EXIT.upstream,
            suggestions: ["Run pro-cli doctor --json or retry with fewer --samples."],
            details: { jobId: lastJobId, sampleIndex: i, attempt },
          },
        );
      }
      const parseResult = parseOddsResponse(lastRaw, input.allowFifty);
      if (parseResult.value !== null) {
        parsed = parseResult.value;
        rejectedFifty = false;
        break;
      }
      if (parseResult.rejectedFifty) {
        rejectedFifty = true;
        rejectedFifties += 1;
      } else {
        parseFailures += 1;
      }
    }
    attempts.push({ jobId: lastJobId, raw: lastRaw, parsed, rejectedFifty });
    jobIds.push(lastJobId);
    if (parsed !== null) samples.push(parsed);
  }

  if (samples.length === 0) {
    throw new ProError("ODDS_PARSE_FAILED", "Could not extract a probability from any sample.", {
      exitCode: EXIT.upstream,
      suggestions: [
        "Re-run with --json to inspect raw responses.",
        "Increase --parse-retries or --samples.",
        "If 50 was the only response, pass --allow-fifty.",
      ],
      details: { attempts, requestedSamples: input.samples },
    });
  }

  const probabilityRaw = aggregateOdds(samples, input.aggregate);
  const probability = Math.max(0, Math.min(100, Math.round(probabilityRaw)));

  return {
    probability,
    probabilityRaw,
    samples,
    aggregate: input.aggregate,
    parseFailures,
    rejectedFifties,
    attempts,
    jobIds,
    model: input.model,
    reasoning: input.reasoning,
    allowFifty: input.allowFifty,
  };
}

function isRecord(value: unknown): value is Record<string, unknown> {
  return Boolean(value) && typeof value === "object" && !Array.isArray(value);
}

function stringField(record: Record<string, unknown>, key: string): string | undefined {
  const v = record[key];
  return typeof v === "string" ? v : undefined;
}
`````

## File: src/output.ts
`````typescript
import { ProError } from "./errors";

export interface CliIO {
  stdout: (text: string) => void;
  stderr: (text: string) => void;
  stdoutIsTTY: boolean;
  env: Record<string, string | undefined>;
  cwd: string;
}

export interface OutputMode {
  json: boolean;
}

const FULL_RELAY_THRESHOLD_CHARS = 6_000;
const APPROX_CHARS_PER_TOKEN = 4;

export function writeSuccess(io: CliIO, mode: OutputMode, payload: unknown): void {
  if (mode.json) {
    io.stdout(`${JSON.stringify({ ok: true, data: withResultRelayInstruction(payload) })}\n`);
    return;
  }
  io.stdout(`${renderText(payload)}\n`);
}

export function writeError(io: CliIO, mode: OutputMode, error: ProError): void {
  if (mode.json) {
    io.stderr(`${JSON.stringify({ ok: false, error: error.toPayload() })}\n`);
    return;
  }

  const suggestions =
    error.suggestions.length > 0 ? `\ntry: ${error.suggestions.join(" | ")}` : "";
  io.stderr(`${error.code}: ${error.message}${suggestions}\n`);
}

export function renderText(payload: unknown): string {
  if (typeof payload === "string") return payload;
  if (!isRecord(payload)) return JSON.stringify(payload);
  if (
    "text" in payload &&
    typeof payload.text === "string"
  ) {
    return payload.text;
  }
  if ("result" in payload && typeof payload.result === "string") {
    return payload.result;
  }
  if ("steps" in payload && Array.isArray(payload.steps)) {
    return renderSetup(payload);
  }
  if ("command" in payload && typeof payload.command === "string") {
    const capture =
      "captureCommand" in payload && typeof payload.captureCommand === "string"
        ? `\n\nThen capture:\n${payload.captureCommand}`
        : "";
    return `Open ChatGPT:\n${payload.command}${capture}`;
  }
  if ("ready" in payload && "next" in payload && isRecord(payload.next)) {
    const status = payload.ready ? "ready" : "not ready";
    const command =
      typeof payload.next.command === "string" ? `\nnext: ${payload.next.command}` : "";
    return `pro-cli ${status}${command}`;
  }
  if ("version" in payload && typeof payload.version === "string" && "repoRoot" in payload) {
    return `pro-cli updated\n${payload.version}`;
  }
  if ("job" in payload && isRecord(payload.job)) {
    const id = typeof payload.job.id === "string" ? payload.job.id : "unknown";
    const status = typeof payload.job.status === "string" ? payload.job.status : "unknown";
    if ("wait" in payload && isRecord(payload.wait) && payload.wait.timedOut === true) {
      const elapsed =
        typeof payload.wait.elapsedMs === "number" ? ` after ${payload.wait.elapsedMs}ms` : "";
      return `job ${id} still ${status}${elapsed}\nwait: pro-cli job wait ${id}`;
    }
    const resultHint =
      status === "succeeded" ? `\nresult: pro-cli job result ${id}` : `\nwait: pro-cli job wait ${id}`;
    return `job ${id} ${status}${resultHint}`;
  }
  return JSON.stringify(payload);
}

function renderSetup(payload: Record<string, unknown>): string {
  const lines: string[] = [];
  if (typeof payload.summary === "string") lines.push(payload.summary);
  const steps = payload.steps as unknown[];
  for (const step of steps) {
    if (!isRecord(step)) continue;
    const status = typeof step.status === "string" ? step.status : "todo";
    const id = typeof step.id === "string" ? step.id : "step";
    lines.push(`[${status}] ${id}`);
    if (typeof step.command === "string") lines.push(`  ${step.command}`);
  }
  return lines.join("\n");
}

function isRecord(value: unknown): value is Record<string, unknown> {
  return Boolean(value) && typeof value === "object";
}

function withResultRelayInstruction(payload: unknown): unknown {
  if (!isRecord(payload) || typeof payload.result !== "string") return payload;
  return {
    ...payload,
    agentInstruction:
      "AI agents: data.result is the primary deliverable. If it is at or under 6000 characters, roughly 1500 tokens or two terminal pages, prefer relaying it in full. If longer, condense only as needed; preserve Pro's prose language, structure, and voice. Do not replace it with a short synopsis unless the user asks for a summary. Do not send probe or smoke-test queries (e.g., 'reply with OK') to pro-cli; send the real request directly. Every call consumes the user's Pro quota. Only test connectivity if the user explicitly asks.",
    resultStats: {
      chars: payload.result.length,
      approximateTokens: Math.ceil(payload.result.length / APPROX_CHARS_PER_TOKEN),
      fullRelayThresholdChars: FULL_RELAY_THRESHOLD_CHARS,
      fullRelayThresholdApproxTokens: Math.ceil(
        FULL_RELAY_THRESHOLD_CHARS / APPROX_CHARS_PER_TOKEN,
      ),
    },
  };
}
`````

## File: src/session-token.ts
`````typescript
import { readFile } from "node:fs/promises";

export interface SessionTokenExport {
  version: 1;
  generatedAt: string;
  source: "pro-cli-cdp-page";
  accessToken: string;
  accountId?: string;
  expiresAt?: string;
}

export function toSessionTokenExport(accessToken: string): SessionTokenExport {
  const expiresMs = jwtExpiryMs(accessToken);
  return {
    version: 1,
    generatedAt: new Date().toISOString(),
    source: "pro-cli-cdp-page",
    accessToken,
    ...(accountIdFromToken(accessToken) ? { accountId: accountIdFromToken(accessToken) } : {}),
    ...(expiresMs !== undefined ? { expiresAt: new Date(expiresMs).toISOString() } : {}),
  };
}

export async function loadSessionToken(path: string): Promise<SessionTokenExport> {
  return JSON.parse(await readFile(path, "utf8")) as SessionTokenExport;
}

export function isTokenFresh(token: SessionTokenExport, skewMs = 60_000): boolean {
  if (!token.expiresAt) return true;
  return Date.now() < Date.parse(token.expiresAt) - skewMs;
}

function accountIdFromToken(token: string): string | undefined {
  const payload = decodeJwtPayload(token);
  const auth = payload?.["https://api.openai.com/auth"] as Record<string, unknown> | undefined;
  return typeof auth?.chatgpt_account_id === "string" ? auth.chatgpt_account_id : undefined;
}

function jwtExpiryMs(token: string): number | undefined {
  const payload = decodeJwtPayload(token);
  return typeof payload?.exp === "number" ? payload.exp * 1000 : undefined;
}

function decodeJwtPayload(token: string): Record<string, unknown> | null {
  try {
    const parts = token.split(".");
    if (parts.length !== 3) return null;
    const normalized = parts[1].replace(/-/g, "+").replace(/_/g, "/");
    const padded = normalized + "=".repeat((4 - (normalized.length % 4 || 4)) % 4);
    return JSON.parse(Buffer.from(padded, "base64").toString("utf8")) as Record<string, unknown>;
  } catch {
    return null;
  }
}
`````

## File: src/structured.ts
`````typescript
import { readFile } from "node:fs/promises";
import { EXIT, ProError } from "./errors";

export interface StructuredOptions {
  schema?: unknown;
  formatHint?: string;
  retries: number;
  runner: (prompt: string) => Promise<string>;
}

export interface StructuredAttempt {
  raw: string;
  error: string | null;
}

export interface StructuredResult {
  parsed: unknown;
  raw: string;
  attempts: StructuredAttempt[];
}

export async function runStructured(userPrompt: string, opts: StructuredOptions): Promise<StructuredResult> {
  if (!opts.schema && !opts.formatHint) {
    throw new ProError("STRUCTURED_NO_HINT", "Pass --schema or --format.", {
      exitCode: EXIT.invalidArgs,
    });
  }
  const baseInstructions = buildStructuredInstructions(opts.schema, opts.formatHint);
  const wrappedPrompt = `${userPrompt.trim()}\n\n${baseInstructions}`;

  const attempts: StructuredAttempt[] = [];
  let lastError = "no attempt yet";

  for (let attempt = 0; attempt <= opts.retries; attempt += 1) {
    const promptToUse =
      attempt === 0
        ? wrappedPrompt
        : `${wrappedPrompt}\n\nPREVIOUS ATTEMPT FAILED. Your reply was:\n\n${attempts[attempts.length - 1].raw}\n\nReason: ${lastError}\n\nReply again with valid JSON. Output ONLY the JSON in a fenced \`\`\`json block.`;
    const raw = await opts.runner(promptToUse);
    try {
      const parsed = extractJsonFromResponse(raw);
      const validation = opts.schema ? validateLightly(parsed, opts.schema) : { ok: true as const };
      if (!validation.ok) {
        lastError = validation.reason ?? "schema validation failed";
        attempts.push({ raw, error: lastError });
        continue;
      }
      attempts.push({ raw, error: null });
      return { parsed, raw, attempts };
    } catch (error) {
      lastError = error instanceof Error ? error.message : String(error);
      attempts.push({ raw, error: lastError });
    }
  }

  throw new ProError("STRUCTURED_PARSE_FAILED", "Could not parse a valid JSON response.", {
    exitCode: EXIT.upstream,
    suggestions: [
      "Increase --schema-retries.",
      "Simplify the schema or use --format for a lighter hint.",
      "Re-run with --json to inspect raw responses.",
    ],
    details: { attempts, lastError },
  });
}

export function buildStructuredInstructions(schema: unknown, formatHint?: string): string {
  if (schema !== undefined && schema !== null) {
    const schemaJson = typeof schema === "string" ? schema : JSON.stringify(schema, null, 2);
    return [
      "Respond with JSON that matches this JSON Schema.",
      "Output exactly one fenced ```json block. No prose before or after.",
      "If a field is unknown, use null. Do not invent fields not in the schema.",
      "",
      "JSON Schema:",
      "```json",
      schemaJson,
      "```",
    ].join("\n");
  }
  if (formatHint) {
    return [
      "Respond with JSON matching this format description.",
      "Output exactly one fenced ```json block. No prose before or after.",
      "",
      "Format:",
      formatHint.trim(),
    ].join("\n");
  }
  throw new ProError("STRUCTURED_NO_HINT", "Either schema or formatHint is required.", {
    exitCode: EXIT.invalidArgs,
  });
}

export function extractJsonFromResponse(text: string): unknown {
  const fence = firstJsonFence(text);
  if (fence !== null) return JSON.parse(fence);

  const start = findFirstJsonStart(text);
  if (start === -1) throw new Error("No JSON object or array found in response.");
  return JSON.parse(extractBalanced(text, start));
}

function firstJsonFence(text: string): string | null {
  const fencePattern = /```([^\n\r`]*)\r?\n?([\s\S]*?)```/g;
  let unlabeled: string | null = null;
  for (const match of text.matchAll(fencePattern)) {
    const language = (match[1] ?? "").trim().toLowerCase();
    const candidate = (match[2] ?? "").trim();
    if (language === "json") return candidate;
    if (!language && unlabeled === null) {
      unlabeled = candidate;
    }
  }
  return unlabeled;
}

function findFirstJsonStart(text: string): number {
  for (let i = 0; i < text.length; i += 1) {
    const c = text[i];
    if (c === "{" || c === "[") return i;
  }
  return -1;
}

function extractBalanced(text: string, start: number): string {
  const open = text[start];
  const close = open === "{" ? "}" : "]";
  let depth = 0;
  let inString = false;
  let escape = false;
  for (let i = start; i < text.length; i += 1) {
    const c = text[i];
    if (escape) {
      escape = false;
      continue;
    }
    if (inString) {
      if (c === "\\") escape = true;
      else if (c === '"') inString = false;
      continue;
    }
    if (c === '"') inString = true;
    else if (c === open) depth += 1;
    else if (c === close) {
      depth -= 1;
      if (depth === 0) return text.slice(start, i + 1);
    }
  }
  throw new Error("Unterminated JSON value in response.");
}

export function validateLightly(value: unknown, schema: unknown): { ok: true } | { ok: false; reason: string } {
  if (!isRecord(schema)) return { ok: true };
  const type = schema.type;
  if (typeof type === "string") {
    if (type === "object" && !isRecord(value)) return { ok: false, reason: "Expected object at root." };
    if (type === "array" && !Array.isArray(value)) return { ok: false, reason: "Expected array at root." };
    if (type === "string" && typeof value !== "string") return { ok: false, reason: "Expected string at root." };
    if (type === "number" && typeof value !== "number") return { ok: false, reason: "Expected number at root." };
    if (type === "integer" && (typeof value !== "number" || !Number.isInteger(value))) {
      return { ok: false, reason: "Expected integer at root." };
    }
    if (type === "boolean" && typeof value !== "boolean") return { ok: false, reason: "Expected boolean at root." };
  }
  if (type === "object" && isRecord(value) && Array.isArray(schema.required)) {
    for (const key of schema.required) {
      if (typeof key === "string" && !(key in value)) {
        return { ok: false, reason: `Missing required field: ${key}` };
      }
    }
  }
  return { ok: true };
}

export async function loadSchema(value: string, cwd: string): Promise<unknown> {
  const text =
    value.startsWith("@") && !value.includes(" ")
      ? await readFile(new URL(value.slice(1), `file://${cwd}/`), "utf8")
      : value;
  try {
    return JSON.parse(text);
  } catch (error) {
    throw new ProError("STRUCTURED_BAD_SCHEMA", "Could not parse --schema as JSON.", {
      exitCode: EXIT.invalidArgs,
      cause: error,
    });
  }
}

function isRecord(value: unknown): value is Record<string, unknown> {
  return Boolean(value) && typeof value === "object" && !Array.isArray(value);
}
`````

## File: src/transport.ts
`````typescript
import { randomUUID } from "node:crypto";
import { evaluateInCdpPage, recoverCookieBloatInCdp } from "./cdp";
import { chatGptOrigins } from "./cookies";
import { DEFAULT_MODEL, isReasoningLevel } from "./defaults";
import { EXIT, ProError } from "./errors";
import type { JobRecord, LimitsObservation } from "./jobs";
import { canonicalModelId, modelUsesThinkingEffort } from "./models";
import { isTokenFresh, loadSessionToken } from "./session-token";

const CHATGPT_CONVERSATION_ENDPOINT = "https://chatgpt.com/backend-api/f/conversation";
const DEFAULT_CDP_BASE = "http://127.0.0.1:9222";

type PageEvaluator = <T>(cdpBase: string, expression: string, timeoutMs?: number) => Promise<T>;

export interface TransportOptions {
  sessionTokenPath: string;
  cdpBase?: string;
  pageEvaluator?: PageEvaluator;
  timeoutMs?: number;
  retries?: number;
  retryDelayMs?: number;
  onLimits?: (observations: LimitsObservation[]) => void;
}

export async function runChatGptJob(job: JobRecord, options: TransportOptions): Promise<string> {
  const session = await loadSessionToken(options.sessionTokenPath).catch(() => null);
  if (!session) {
    throw new ProError("SESSION_TOKEN_MISSING", "No ChatGPT session token is available.", {
      exitCode: EXIT.auth,
      suggestions: ["Run pro-cli auth capture from a logged-in ChatGPT CDP browser."],
    });
  }
  if (!isTokenFresh(session)) {
    throw new ProError("SESSION_TOKEN_EXPIRED", "The ChatGPT session token is expired.", {
      exitCode: EXIT.auth,
      suggestions: ["Run pro-cli auth capture again from a logged-in ChatGPT browser."],
    });
  }
  if (!session.accountId) {
    throw new ProError("ACCOUNT_ID_MISSING", "The ChatGPT account id is missing from the token.", {
      exitCode: EXIT.auth,
      suggestions: ["Run pro-cli auth capture again and confirm ChatGPT is logged in."],
    });
  }

  const retries = integerOption(options.retries, 0, 0, 5) ?? 0;
  const retryDelayMs = integerOption(options.retryDelayMs, 500, 0, 60_000) ?? 500;
  let lastError: ProError | null = null;

  for (let attempt = 0; attempt <= retries; attempt += 1) {
    try {
      return await postChatGptJob(job, { accountId: session.accountId }, options);
    } catch (error) {
      const proError = error instanceof ProError ? error : networkError(error);
      lastError = proError;
      if (attempt >= retries || !isRetryable(proError)) throw withAttemptDetails(proError, attempt + 1);
      if (retryDelayMs > 0) await sleep(retryDelayMs);
    }
  }

  throw lastError ?? new ProError("UPSTREAM_ERROR", "ChatGPT backend request failed.", { exitCode: EXIT.upstream });
}

async function postChatGptJob(
  job: JobRecord,
  session: { accountId: string },
  options: TransportOptions,
): Promise<string> {
  const timeoutMs = integerOption(options.timeoutMs ?? job.options.timeoutMs, 0, 0, 30 * 60_000) ?? 0;

  try {
    const evaluate = options.pageEvaluator ?? evaluateInCdpPage;
    const cdpBase = options.cdpBase ?? DEFAULT_CDP_BASE;
    let browserResult = await evaluate<BrowserFetchResult>(
      cdpBase,
      buildBrowserFetchExpression(buildRequestBody(job), session.accountId),
      timeoutMs || 30 * 60_000,
    );
    if (!options.pageEvaluator && shouldRecoverCookieBloat(browserResult)) {
      const recovered = await recoverCookieBloatInCdp(cdpBase, chatGptOrigins(), timeoutMs || 10_000).catch(() => null);
      if (recovered?.deleted) {
        browserResult = await evaluate<BrowserFetchResult>(
          cdpBase,
          buildBrowserFetchExpression(buildRequestBody(job), session.accountId),
          timeoutMs || 30 * 60_000,
        );
      }
    }

    if (browserResult.code === "CHATGPT_PAGE_MISSING") {
      throw new ProError("CHATGPT_PAGE_MISSING", "No logged-in ChatGPT page is available over CDP.", {
        exitCode: EXIT.auth,
        suggestions: [
          "Open the Chrome command from pro-cli auth command.",
          "Confirm the CDP Chrome window is on https://chatgpt.com/ and logged in.",
          "Pass --cdp if Chrome is using a non-default CDP port.",
        ],
        details: { cdpBase },
      });
    }

    if (browserResult.code === "CHATGPT_PAGE_LOGGED_OUT") {
      throw new ProError("CHATGPT_PAGE_LOGGED_OUT", "The ChatGPT CDP page is not logged in.", {
        exitCode: EXIT.auth,
        suggestions: [
          "Sign in to ChatGPT in the Chrome window from pro-cli auth command.",
          "Run pro-cli auth capture --cdp http://127.0.0.1:9222 --json after login.",
          "Retry pro-cli ask with the same --cdp value.",
        ],
        details: { status: browserResult.status },
      });
    }

    if (browserResult.code === "CHATGPT_PROBE_FAILED") {
      const status = browserResult.status;
      const suggestions =
        status === 431
          ? [
              "HTTP 431 indicates oversize request headers; the CDP Chrome profile likely has stale cookie buildup.",
              "Sign out of ChatGPT in the CDP window, sign back in to drop expired cookies, then run pro-cli auth capture --cdp http://127.0.0.1:9222 --json.",
              "If 431 persists, delete ~/.pro-cli/chrome-profile and rerun pro-cli auth command.",
            ]
          : [
              `The ChatGPT auth session probe failed with HTTP ${status}; pro-cli cannot tell whether the page is logged in.`,
              "Reload the CDP ChatGPT tab and retry. Run pro-cli doctor --json for diagnostics.",
            ];
      throw new ProError(
        "CHATGPT_PROBE_FAILED",
        `Could not determine ChatGPT login state from the CDP page (HTTP ${status}).`,
        {
          exitCode: EXIT.auth,
          suggestions,
          details: { status, preview: browserResult.body.slice(0, 240).replace(/\s+/g, " ") },
        },
      );
    }

    if (!browserResult.ok) {
      throw new ProError("UPSTREAM_REJECTED", `ChatGPT backend returned HTTP ${browserResult.status}.`, {
        exitCode: EXIT.upstream,
        suggestions: ["Run pro-cli auth capture again.", "Check whether the ChatGPT Pro usage limit is reached."],
        details: {
          status: browserResult.status,
          preview: browserResult.body.slice(0, 160).replace(/\s+/g, " "),
        },
      });
    }

    return readResponseText(browserResult.body, options.onLimits);
  } catch (error) {
    if (error instanceof ProError) throw error;
    throw networkError(error);
  }
}

function shouldRecoverCookieBloat(result: BrowserFetchResult): boolean {
  return result.status === 431 || result.body.includes("chrome-error://chromewebdata/");
}

interface BrowserFetchResult {
  ok: boolean;
  status: number;
  body: string;
  code?: "CHATGPT_PAGE_MISSING" | "CHATGPT_PAGE_LOGGED_OUT" | "CHATGPT_PROBE_FAILED";
}

function buildBrowserFetchExpression(requestBody: Record<string, unknown>, accountId: string): string {
  return `(${async function browserFetch(
    endpoint: string,
    body: Record<string, unknown>,
    account: string,
  ): Promise<BrowserFetchResult> {
    if (location.origin !== "https://chatgpt.com") {
      return {
        ok: false,
        status: 0,
        code: "CHATGPT_PAGE_MISSING",
        body: `Expected https://chatgpt.com, got ${location.href}`,
      };
    }

    const sessionResponse = await fetch("https://chatgpt.com/api/auth/session", { credentials: "include", referrerPolicy: "no-referrer" });
    const session = (await sessionResponse.json().catch(() => null)) as { accessToken?: unknown } | null;
    if (!sessionResponse.ok && sessionResponse.status !== 401) {
      return {
        ok: false,
        status: sessionResponse.status,
        code: "CHATGPT_PROBE_FAILED",
        body: `ChatGPT auth session probe returned HTTP ${sessionResponse.status}.`,
      };
    }
    if (typeof session?.accessToken !== "string" || !session.accessToken) {
      return {
        ok: false,
        status: sessionResponse.status,
        code: "CHATGPT_PAGE_LOGGED_OUT",
        body: "ChatGPT page session did not include an access token.",
      };
    }

    const accessToken = session.accessToken;
    const turnTraceId = crypto.randomUUID();
    const requestBody = withBrowserContext(body);
    const referrer = chatReferrer(requestBody);
    const prepareBody = buildPrepareBody(requestBody);
    const prepareResponse = await fetch("https://chatgpt.com/backend-api/f/conversation/prepare", {
      method: "POST",
      credentials: "include",
      referrer,
      headers: appHeaders("/f/conversation/prepare", accessToken, {
        "x-conduit-token": "no-token",
        "x-oai-turn-trace-id": turnTraceId,
      }),
      body: JSON.stringify(prepareBody),
    });
    const preparedConversation = (await prepareResponse.json().catch(() => null)) as
      | { conduit_token?: unknown }
      | null;
    const conduitToken =
      prepareResponse.ok && typeof preparedConversation?.conduit_token === "string"
        ? preparedConversation.conduit_token
        : null;

    const headers = {
      ...appHeaders("/f/conversation", accessToken, {
        accept: "text/event-stream",
        "x-oai-turn-trace-id": turnTraceId,
        ...(conduitToken ? { "x-conduit-token": conduitToken } : {}),
      }),
      ...(await chatRequirementsHeaders(accessToken, referrer)),
    };

    const response = await fetch(endpoint, {
      method: "POST",
      credentials: "include",
      referrer,
      headers,
      body: JSON.stringify({ ...requestBody, client_prepare_state: prepareResponse.ok ? "sent" : "none" }),
    });
    let text = await response.text().catch((error) => String(error));
    if (response.ok) {
      const resumedText = await resumeHandoffStream(text, accessToken, turnTraceId, referrer);
      if (resumedText) text = `${text}\n\n${resumedText}`;
    }
    return { ok: response.ok, status: response.status, body: text };

    function appHeaders(
      routeName: string,
      accessToken: string,
      extraHeaders: Record<string, string> = {},
    ): Record<string, string> {
      const headers: Record<string, string> = {
        authorization: `Bearer ${accessToken}`,
        "content-type": "application/json",
        "oai-language": navigator.language || "en-US",
        "OAI-Client-Version": document.documentElement.getAttribute("data-build") ?? "",
        "OAI-Client-Build-Number": document.documentElement.getAttribute("data-seq") ?? "",
        "OAI-Device-Id": readJsonString(localStorage.getItem("oai-did")) ?? readCookie("oai-did") ?? "",
        "OAI-Session-Id": readSessionId(),
        "X-OpenAI-Target-Path": `/backend-api${routeName}`,
        "X-OpenAI-Target-Route": `/backend-api${routeName}`,
        ...extraHeaders,
      };
      const integrityState = readCookie("__Secure-oai-is");
      if (integrityState) headers["X-OAI-IS"] = integrityState;
      return Object.fromEntries(Object.entries(headers).filter(([, value]) => value.length > 0));
    }

    function buildPrepareBody(body: Record<string, unknown>): Record<string, unknown> {
      const messages = Array.isArray(body.messages) ? body.messages : [];
      const firstMessage = messages[0] as {
        id?: unknown;
        author?: unknown;
        content?: { parts?: unknown };
      } | undefined;
      const partialQuery = firstMessage
        ? {
            id: firstMessage.id,
            author: firstMessage.author,
            content: {
              ...(firstMessage.content ?? {}),
              parts: Array.isArray(firstMessage.content?.parts) ? firstMessage.content.parts : [],
            },
          }
        : undefined;
      const {
        messages: _messages,
        enable_message_followups: _followups,
        paragen_cot_summary_display_override: _paragen,
        force_parallel_switch: _parallel,
        ...prepareBody
      } = body;
      return {
        ...prepareBody,
        fork_from_shared_post: false,
        partial_query: partialQuery,
        client_prepare_state: "none",
        client_contextual_info: { app_name: appNameFor(body) },
      };
    }

    function withBrowserContext(body: Record<string, unknown>): Record<string, unknown> {
      return {
        ...body,
        timezone_offset_min: new Date().getTimezoneOffset(),
        timezone: Intl.DateTimeFormat().resolvedOptions().timeZone,
        client_contextual_info: {
          is_dark_mode: matchMedia?.("(prefers-color-scheme: dark)")?.matches ?? false,
          time_since_loaded: Math.round(performance.now()),
          page_height: document.documentElement.scrollHeight,
          page_width: document.documentElement.scrollWidth,
          pixel_ratio: window.devicePixelRatio,
          screen_height: screen.height,
          screen_width: screen.width,
          app_name: appNameFor(body),
        },
      };
    }

    function appNameFor(body: Record<string, unknown>): string {
      return body.history_and_training_disabled === true ? "chatgpt.com" : "chatgpt";
    }

    function chatReferrer(body: Record<string, unknown>): string {
      return body.history_and_training_disabled === true
        ? "https://chatgpt.com/?temporary-chat=true"
        : "https://chatgpt.com/";
    }

    async function resumeHandoffStream(
      streamText: string,
      accessToken: string,
      turnTraceId: string,
      referrer: string,
    ): Promise<string | null> {
      const handoff = readHandoff(streamText);
      if (!handoff) return null;
      for (const offset of [0, 1, 2]) {
        const resumeResponse = await fetch("https://chatgpt.com/backend-api/f/conversation/resume", {
          method: "POST",
          credentials: "include",
          referrer,
          headers: appHeaders("/f/conversation/resume", accessToken, {
            accept: "text/event-stream",
            "x-conduit-token": handoff.token,
            "x-oai-turn-trace-id": turnTraceId,
          }),
          body: JSON.stringify({ conversation_id: handoff.conversationId, offset }),
        });
        const resumeText = await resumeResponse.text().catch(() => "");
        if (resumeResponse.ok && resumeText.trim()) return resumeText;
        if (resumeResponse.status !== 404) return null;
      }
      return null;
    }

    function readHandoff(streamText: string): { conversationId: string; token: string } | null {
      let conversationId: string | null = null;
      let token: string | null = null;
      for (const event of readSseJsonEvents(streamText)) {
        if (!event || typeof event !== "object") continue;
        const record = event as { type?: unknown; conversation_id?: unknown; token?: unknown };
        if (record.type === "resume_conversation_token") {
          if (typeof record.conversation_id === "string") conversationId = record.conversation_id;
          if (typeof record.token === "string") token = record.token;
        }
        if (record.type === "stream_handoff" && typeof record.conversation_id === "string") {
          conversationId = record.conversation_id;
        }
      }
      return conversationId && token ? { conversationId, token } : null;
    }

    function readSseJsonEvents(streamText: string): unknown[] {
      return streamText
        .split("\n\n")
        .flatMap((frame) => {
          const data = frame
            .split("\n")
            .filter((line) => line.startsWith("data:"))
            .map((line) => line.slice(5).trim())
            .join("\n");
          if (!data || data === "[DONE]") return [];
          try {
            return [JSON.parse(data) as unknown];
          } catch {
            return [];
          }
        });
    }

    function readCookie(name: string): string | null {
      const prefix = `${name}=`;
      const cookie = document.cookie
        .split("; ")
        .find((item) => item.startsWith(prefix));
      return cookie ? decodeURIComponent(cookie.slice(prefix.length)) : null;
    }

    function readJsonString(value: string | null): string | null {
      if (!value) return null;
      try {
        const parsed = JSON.parse(value) as unknown;
        return typeof parsed === "string" ? parsed : null;
      } catch {
        return value;
      }
    }

    function readSessionId(): string {
      const bootstrap = (window as unknown as { CLIENT_BOOTSTRAP?: { sessionId?: unknown } }).CLIENT_BOOTSTRAP;
      if (typeof bootstrap?.sessionId === "string" && bootstrap.sessionId) return bootstrap.sessionId;
      const statsigKey = Object.keys(localStorage).find((key) => key.startsWith("statsig.session_id."));
      if (statsigKey) {
        try {
          const statsig = JSON.parse(localStorage.getItem(statsigKey) ?? "{}") as { sessionID?: unknown };
          if (typeof statsig.sessionID === "string" && statsig.sessionID) return statsig.sessionID;
        } catch {
          // Ignore malformed local client telemetry state.
        }
      }
      return crypto.randomUUID();
    }

    async function chatRequirementsHeaders(accessToken: string, referrer: string): Promise<Record<string, string>> {
      const requirementsToken = buildRequirementsToken();
      const prepareResponse = await fetch(
        "https://chatgpt.com/backend-api/sentinel/chat-requirements/prepare",
        {
          method: "POST",
          credentials: "include",
          referrer,
          headers: appHeaders("/sentinel/chat-requirements/prepare", accessToken),
          body: JSON.stringify({ p: requirementsToken }),
        },
      );
      const prepared = (await prepareResponse.json().catch(() => null)) as PreparedChatRequirements | null;
      if (!prepareResponse.ok || !prepared) {
        return {};
      }

      const finalizeBody: Record<string, string> = {
        prepare_token: typeof prepared.prepare_token === "string" ? prepared.prepare_token : "",
      };
      const proofToken = buildProofToken(prepared);
      if (proofToken) finalizeBody.proofofwork = proofToken;
      const turnstileToken = await buildTurnstileToken(prepared, requirementsToken);
      if (turnstileToken) finalizeBody.turnstile = turnstileToken;

      const finalizeResponse = await fetch(
        "https://chatgpt.com/backend-api/sentinel/chat-requirements/finalize",
        {
          method: "POST",
          credentials: "include",
          referrer,
          headers: appHeaders("/sentinel/chat-requirements/finalize", accessToken),
          body: JSON.stringify(finalizeBody),
        },
      );
      const finalized = (await finalizeResponse.json().catch(() => null)) as
        | { token?: unknown }
        | null;
      if (!finalizeResponse.ok || typeof finalized?.token !== "string" || !finalized.token) {
        return {};
      }

      const headers: Record<string, string> = {
        "OpenAI-Sentinel-Chat-Requirements-Token": finalized.token,
      };
      if (proofToken) headers["OpenAI-Sentinel-Proof-Token"] = proofToken;
      if (turnstileToken) headers["OpenAI-Sentinel-Turnstile-Token"] = turnstileToken;
      const timing = sentinelTiming();
      if (timing) headers["OAI-Telemetry"] = timing;
      return headers;
    }

    function buildRequirementsToken(): string {
      return `gAAAAAC${generateRequirementsTokenAnswer()}`;
    }

    function buildProofToken(prepared: PreparedChatRequirements): string | null {
      const proof = prepared.proofofwork;
      if (!proof?.required) return null;
      if (typeof proof.seed !== "string" || typeof proof.difficulty !== "string") return null;
      return `gAAAAAB${generateProofAnswer(proof.seed, proof.difficulty)}`;
    }

    async function buildTurnstileToken(
      prepared: PreparedChatRequirements,
      requirementsToken: string,
    ): Promise<string | null> {
      const turnstile = prepared.turnstile;
      if (!turnstile?.required) return null;
      if (typeof turnstile.dx === "string" && turnstile.dx) {
        return await runDxProgram(requirementsToken, turnstile.dx).catch(() => null);
      }
      return null;
    }

    function sentinelTiming(): string | null {
      try {
        const sentinel = (window as unknown as { SentinelSDK?: { timing?: () => unknown } }).SentinelSDK;
        const timing = sentinel?.timing?.();
        return typeof timing === "string" ? timing : null;
      } catch {
        return null;
      }
    }

    async function runDxProgram(secret: string, dx: string): Promise<string> {
      const opXorAsync = 0;
      const opXor = 1;
      const opSet = 2;
      const opResolve = 3;
      const opReject = 4;
      const opAppend = 5;
      const opIndex = 6;
      const opCall = 7;
      const opCopy = 8;
      const opQueue = 9;
      const opWindow = 10;
      const opScriptMatch = 11;
      const opMap = 12;
      const opSafeCall = 13;
      const opJsonParse = 14;
      const opJsonStringify = 15;
      const opSecret = 16;
      const opCallSet = 17;
      const opAtob = 18;
      const opBtoa = 19;
      const opEqualsBranch = 20;
      const opDeltaBranch = 21;
      const opSubroutine = 22;
      const opIfDefined = 23;
      const opBind = 24;
      const opNoopA = 25;
      const opNoopB = 26;
      const opRemove = 27;
      const opNoopC = 28;
      const opLessThan = 29;
      const opDefineFunction = 30;
      const opMultiply = 33;
      const opAwait = 34;
      const opDivide = 35;
      const values = new Map<number, unknown>();
      let steps = 0;
      let chain = Promise.resolve();

      function serialize<T>(work: () => Promise<T> | T): Promise<T> {
        const next = chain.then(work, work);
        chain = next.then(
          () => undefined,
          () => undefined,
        );
        return next;
      }

      async function runQueue(): Promise<void> {
        const queue = values.get(opQueue) as unknown[][];
        while (Array.isArray(queue) && queue.length > 0) {
          const [opcode, ...args] = queue.shift() ?? [];
          const handler = values.get(Number(opcode)) as ((...args: unknown[]) => unknown) | undefined;
          const result = handler?.(...args);
          if (result && typeof (result as Promise<unknown>).then === "function") await result;
          steps += 1;
        }
      }

      function xor(value: string, key: string): string {
        let output = "";
        for (let index = 0; index < value.length; index += 1) {
          output += String.fromCharCode(value.charCodeAt(index) ^ key.charCodeAt(index % key.length));
        }
        return output;
      }

      function resetVm(): void {
        values.clear();
        values.set(opXorAsync, (program: unknown) => runDxProgram(String(values.get(Number(program))), secret));
        values.set(opXor, (target: unknown, key: unknown) =>
          values.set(Number(target), xor(String(values.get(Number(target))), String(values.get(Number(key))))),
        );
        values.set(opSet, (target: unknown, value: unknown) => values.set(Number(target), value));
        values.set(opAppend, (target: unknown, source: unknown) => {
          const current = values.get(Number(target));
          const next = values.get(Number(source));
          if (Array.isArray(current)) current.push(next);
          else values.set(Number(target), String(current) + String(next));
        });
        values.set(opRemove, (target: unknown, source: unknown) => {
          const current = values.get(Number(target));
          const next = values.get(Number(source));
          if (Array.isArray(current)) current.splice(current.indexOf(next), 1);
          else values.set(Number(target), Number(current) - Number(next));
        });
        values.set(opLessThan, (target: unknown, left: unknown, right: unknown) =>
          values.set(Number(target), Number(values.get(Number(left))) < Number(values.get(Number(right)))),
        );
        values.set(opMultiply, (target: unknown, left: unknown, right: unknown) =>
          values.set(Number(target), Number(values.get(Number(left))) * Number(values.get(Number(right)))),
        );
        values.set(opDivide, (target: unknown, left: unknown, right: unknown) => {
          const divisor = Number(values.get(Number(right)));
          values.set(Number(target), divisor === 0 ? 0 : Number(values.get(Number(left))) / divisor);
        });
        values.set(opIndex, (target: unknown, source: unknown, key: unknown) =>
          values.set(Number(target), (values.get(Number(source)) as Record<string, unknown>)[String(values.get(Number(key)))]),
        );
        values.set(opCall, (fn: unknown, ...args: unknown[]) =>
          (values.get(Number(fn)) as (...args: unknown[]) => unknown)(...args.map((arg) => values.get(Number(arg)))),
        );
        values.set(opCallSet, (target: unknown, fn: unknown, ...args: unknown[]) => {
          try {
            const result = (values.get(Number(fn)) as (...args: unknown[]) => unknown)(
              ...args.map((arg) => values.get(Number(arg))),
            );
            if (result && typeof (result as Promise<unknown>).then === "function") {
              return (result as Promise<unknown>)
                .then((value) => values.set(Number(target), value))
                .catch((error) => values.set(Number(target), String(error)));
            }
            values.set(Number(target), result);
          } catch (error) {
            values.set(Number(target), String(error));
          }
        });
        values.set(opSafeCall, (target: unknown, fn: unknown, ...args: unknown[]) => {
          try {
            (values.get(Number(fn)) as (...args: unknown[]) => unknown)(...args.map((arg) => values.get(Number(arg))));
          } catch (error) {
            values.set(Number(target), String(error));
          }
        });
        values.set(opCopy, (target: unknown, source: unknown) => values.set(Number(target), values.get(Number(source))));
        values.set(opWindow, window);
        values.set(opScriptMatch, (target: unknown, pattern: unknown) =>
          values.set(
            Number(target),
            (Array.from(document.scripts || [])
              .map((script) => script?.src?.match(String(values.get(Number(pattern)))))
              .filter((match) => match?.length)[0] ?? [])[0] ?? null,
          ),
        );
        values.set(opMap, (target: unknown) => values.set(Number(target), values));
        values.set(opJsonParse, (target: unknown, source: unknown) =>
          values.set(Number(target), JSON.parse(String(values.get(Number(source))))),
        );
        values.set(opJsonStringify, (target: unknown, source: unknown) =>
          values.set(Number(target), JSON.stringify(values.get(Number(source)))),
        );
        values.set(opAtob, (target: unknown) => values.set(Number(target), atob(String(values.get(Number(target))))));
        values.set(opBtoa, (target: unknown) => values.set(Number(target), btoa(String(values.get(Number(target))))));
        values.set(opEqualsBranch, (left: unknown, right: unknown, fn: unknown, ...args: unknown[]) =>
          values.get(Number(left)) === values.get(Number(right))
            ? (values.get(Number(fn)) as (...args: unknown[]) => unknown)(...args)
            : null,
        );
        values.set(opDeltaBranch, (left: unknown, right: unknown, threshold: unknown, fn: unknown, ...args: unknown[]) =>
          Math.abs(Number(values.get(Number(left))) - Number(values.get(Number(right)))) > Number(values.get(Number(threshold)))
            ? (values.get(Number(fn)) as (...args: unknown[]) => unknown)(...args)
            : null,
        );
        values.set(opIfDefined, (source: unknown, fn: unknown, ...args: unknown[]) =>
          values.get(Number(source)) === undefined
            ? null
            : (values.get(Number(fn)) as (...args: unknown[]) => unknown)(...args),
        );
        values.set(opBind, (target: unknown, source: unknown, key: unknown) => {
          const object = values.get(Number(source)) as Record<string, unknown>;
          const method = object[String(values.get(Number(key)))] as (...args: unknown[]) => unknown;
          values.set(Number(target), method.bind(object));
        });
        values.set(opAwait, (target: unknown, source: unknown) => {
          try {
            const promise = values.get(Number(source));
            return Promise.resolve(promise).then((value) => values.set(Number(target), value));
          } catch {
            return undefined;
          }
        });
        values.set(opSubroutine, (target: unknown, queue: unknown[]) => {
          const previous = [...(values.get(opQueue) as unknown[][])];
          values.set(opQueue, [...queue]);
          return runQueue()
            .catch((error) => values.set(Number(target), String(error)))
            .finally(() => values.set(opQueue, previous));
        });
        values.set(opNoopA, () => undefined);
        values.set(opNoopB, () => undefined);
        values.set(opNoopC, () => undefined);
      }

      return await serialize(
        () =>
          new Promise<string>((resolve, reject) => {
            resetVm();
            values.set(opSecret, secret);
            let settled = false;
            const timer = setTimeout(() => {
              if (settled) return;
              settled = true;
              resolve(String(steps));
            }, 500);
            values.set(opResolve, (value: unknown) => {
              if (settled) return;
              settled = true;
              clearTimeout(timer);
              resolve(btoa(String(value)));
            });
            values.set(opReject, (value: unknown) => {
              if (settled) return;
              settled = true;
              clearTimeout(timer);
              reject(new Error(btoa(String(value))));
            });
            values.set(opDefineFunction, (target: unknown, returnSlot: unknown, argSlotsOrQueue: unknown, queueOrArgs: unknown) => {
              const hasArgSlots = Array.isArray(queueOrArgs);
              const argSlots = (hasArgSlots ? argSlotsOrQueue : []) as unknown[];
              const queue = (hasArgSlots ? queueOrArgs : argSlotsOrQueue) as unknown[];
              values.set(Number(target), (...args: unknown[]) => {
                if (settled) return undefined;
                const previous = [...(values.get(opQueue) as unknown[][])];
                if (hasArgSlots) {
                  for (let index = 0; index < argSlots.length; index += 1) {
                    values.set(Number(argSlots[index]), args[index]);
                  }
                }
                values.set(opQueue, [...queue]);
                return runQueue()
                  .then(() => values.get(Number(returnSlot)))
                  .catch((error) => String(error))
                  .finally(() => values.set(opQueue, previous));
              });
            });
            try {
              values.set(opQueue, JSON.parse(xor(atob(dx), secret)) as unknown[][]);
              runQueue().catch((error) => {
                if (!settled) {
                  settled = true;
                  clearTimeout(timer);
                  resolve(btoa(`${steps}: ${error}`));
                }
              });
            } catch (error) {
              if (!settled) {
                settled = true;
                clearTimeout(timer);
                resolve(btoa(`${steps}: ${error}`));
              }
            }
          }),
      );
    }

    function generateRequirementsTokenAnswer(): string {
      try {
        const config = proofConfig();
        config[3] = 1;
        config[9] = 0;
        return encodeProofConfig(config);
      } catch (error) {
        return `wQ8Lk5FbGpA2NcR9dShT6gYjU7VxZ4D${encodeProofConfig(String(error ?? "e"))}`;
      }
    }

    function generateProofAnswer(seed: string, difficulty: string): string {
      const start = performance.now();
      const config = proofConfig();
      for (let attempt = 0; attempt < 500_000; attempt += 1) {
        config[3] = attempt;
        config[9] = Math.round(performance.now() - start);
        const encoded = encodeProofConfig(config);
        if (fnvHash(`${seed}${encoded}`).substring(0, difficulty.length) <= difficulty) {
          return `${encoded}~S`;
        }
      }
      return `wQ8Lk5FbGpA2NcR9dShT6gYjU7VxZ4D${encodeProofConfig("e")}`;
    }

    function proofConfig(): unknown[] {
      const memory = (performance as Performance & { memory?: { jsHeapSizeLimit?: unknown } }).memory;
      return [
        (screen?.width ?? 0) + (screen?.height ?? 0),
        `${new Date()}`,
        memory?.jsHeapSizeLimit,
        Math.random(),
        navigator.userAgent,
        randomItem(Array.from(document.scripts).map((script) => script?.src).filter(Boolean)),
        Array.from(document.scripts || [])
          .map((script) => script?.src?.match("c/[^/]*/_"))
          .filter((match) => match?.length)[0]?.[0] ?? document.documentElement.getAttribute("data-build"),
        navigator.language,
        navigator.languages?.join(","),
        Math.random(),
        randomNavigatorProbe(),
        randomItem(Object.keys(document)),
        randomItem(Object.keys(window)),
        performance.now(),
        crypto.randomUUID(),
        [...new URLSearchParams(window.location.search).keys()].join(","),
        navigator?.hardwareConcurrency,
        performance.timeOrigin,
        Number("ai" in window),
        Number("createPRNG" in window),
        Number("cache" in window),
        Number("data" in window),
        Number("solana" in window),
        Number("dump" in window),
        Number("InstallTrigger" in window),
      ];
    }

    function randomNavigatorProbe(): string {
      const key = randomItem(Object.keys(Object.getPrototypeOf(navigator)));
      try {
        const value = (navigator as unknown as Record<string, unknown>)[key];
        return `${key}-${String(value)}`;
      } catch {
        return key;
      }
    }

    function randomItem(items: string[]): string {
      if (items.length === 0) return "";
      return items[Math.floor(Math.random() * items.length)] ?? "";
    }

    function encodeProofConfig(value: unknown): string {
      const json = JSON.stringify(value);
      if (window.TextEncoder) {
        return btoa(String.fromCharCode(...new TextEncoder().encode(json)));
      }
      return btoa(unescape(encodeURIComponent(json)));
    }

    function fnvHash(value: string): string {
      let hash = 2166136261;
      for (let index = 0; index < value.length; index += 1) {
        hash ^= value.charCodeAt(index);
        hash = Math.imul(hash, 16777619) >>> 0;
      }
      hash ^= hash >>> 16;
      hash = Math.imul(hash, 2246822507) >>> 0;
      hash ^= hash >>> 13;
      hash = Math.imul(hash, 3266489909) >>> 0;
      hash ^= hash >>> 16;
      return (hash >>> 0).toString(16).padStart(8, "0");
    }
  }})(${JSON.stringify(CHATGPT_CONVERSATION_ENDPOINT)}, ${JSON.stringify(requestBody)}, ${JSON.stringify(accountId)})`;
}

interface PreparedChatRequirements {
  prepare_token?: unknown;
  proofofwork?: {
    required?: unknown;
    seed?: unknown;
    difficulty?: unknown;
  };
  turnstile?: {
    required?: unknown;
    dx?: unknown;
  };
}

function buildRequestBody(job: JobRecord): Record<string, unknown> {
  const prompt = buildConversationPrompt(job);
  const model = normalizeModel(job.model);
  const thinkingEffort = modelUsesThinkingEffort(model) ? normalizeReasoning(job.reasoning) : undefined;
  const conversationId = stringOption(job.options.conversationId);
  const parentMessageId = stringOption(job.options.parentMessageId) ?? "client-created-root";
  const temporary = booleanOption(job.options.temporary, !conversationId);
  const verbosity = stringOption(job.options.verbosity);
  const reasoningSummary = stringOption(job.options.reasoningSummary);
  const toolChoice = stringOption(job.options.toolChoice);
  const parallelTools =
    typeof job.options.parallelTools === "boolean" ? job.options.parallelTools : undefined;
  const body: Record<string, unknown> = {
    action: "next",
    messages: [
      {
        id: randomUUID(),
        author: { role: "user" },
        create_time: Math.floor(Date.now() / 1000),
        content: { content_type: "text", parts: [prompt] },
        metadata: {},
      },
    ],
    model,
    parent_message_id: parentMessageId,
    client_prepare_state: "none",
    timezone_offset_min: new Date().getTimezoneOffset(),
    timezone: Intl.DateTimeFormat().resolvedOptions().timeZone ?? "UTC",
    conversation_mode: { kind: "primary_assistant" },
    enable_message_followups: true,
    system_hints: [],
    supports_buffering: true,
    supported_encodings: ["v1"],
    client_contextual_info: { app_name: "chatgpt" },
    paragen_cot_summary_display_override: "allow",
    force_parallel_switch: "auto",
  };

  if (conversationId) {
    body.conversation_id = conversationId;
  }
  if (temporary) {
    body.history_and_training_disabled = true;
    body.client_contextual_info = { app_name: "chatgpt.com" };
  }
  if (verbosity) body.verbosity = verbosity;
  if (reasoningSummary) body.reasoning_summary = reasoningSummary;
  if (toolChoice) body.tool_choice = toolChoice;
  if (parallelTools !== undefined) {
    body.parallel_tools = parallelTools;
    body.force_parallel_switch = parallelTools ? "auto" : "none";
  }
  if (thinkingEffort) body.thinking_effort = thinkingEffort;

  return body;
}

function buildConversationPrompt(job: JobRecord): string {
  const baseInstructions =
    stringOption(job.options.instructions) ??
    "You are a concise assistant responding to a terminal automation request.";
  const condensedResponseTokens = integerOption(
    job.options.condensedResponseTokens,
    undefined,
    1,
    100_000,
  );
  const instructions = [
    baseInstructions,
    condensedResponseTokens === undefined
      ? undefined
      : `Condensed response mode: keep the final answer to approximately ${condensedResponseTokens} tokens or fewer. Prioritize the user's requested deliverable, concrete decisions, and essential caveats. Do not add filler, broad background, or meta commentary about this limit.`,
  ].filter((part): part is string => Boolean(part && part.trim())).join("\n\n");
  const prompt = job.prompt.trim();
  if (!instructions.trim()) return prompt;
  return `${instructions.trim()}\n\n${prompt}`;
}

function readResponseText(
  raw: string,
  onLimits?: (observations: LimitsObservation[]) => void,
): string {
  let buffer = raw.replace(/\r\n?/g, "\n");
  let completedText: string | null = null;
  let completed = false;
  const state: ResponseParseState = { acceptsTextContinuation: false, lastAppendText: null };

  let boundary = buffer.indexOf("\n\n");
  while (boundary !== -1) {
    const frame = buffer.slice(0, boundary);
    buffer = buffer.slice(boundary + 2);
    const event = parseSseFrame(frame);
    const parsed = readConversationEvent(event, state);
    if (parsed.text !== null) {
      completedText = mergeStreamText(completedText, parsed.text, parsed.append);
    }
    if (onLimits) {
      const observations = extractLimitsProgress(event);
      if (observations.length > 0) onLimits(observations);
    }
    completed = completed || parsed.completed;
    boundary = buffer.indexOf("\n\n");
  }

  if (buffer.trim()) {
    const event = parseSseFrame(buffer);
    const parsed = readConversationEvent(event, state);
    if (parsed.text !== null) {
      completedText = mergeStreamText(completedText, parsed.text, parsed.append);
    }
    if (onLimits) {
      const observations = extractLimitsProgress(event);
      if (observations.length > 0) onLimits(observations);
    }
    completed = completed || parsed.completed;
  }

  if (!completed) {
    throw new ProError("STREAM_INCOMPLETE", "ChatGPT stream ended before the conversation completed.", {
      exitCode: EXIT.network,
      suggestions: [
        "Retry the same real request only if the user still needs it; do not send a probe or smoke-test query.",
        "Increase --timeout if the request is large.",
        "Run pro-cli doctor --json to check local auth/browser health without spending Pro quota.",
      ],
      details: completedText ? { partialPreview: completedText.slice(0, 160) } : undefined,
    });
  }

  if (completedText === null) {
    throw new ProError("EMPTY_RESPONSE", "ChatGPT completed without returning assistant text.", {
      exitCode: EXIT.upstream,
      suggestions: [
        "Retry the same real request only if the user still needs it; do not send a probe or smoke-test query.",
        "Run pro-cli doctor --json to check local auth/browser health without spending Pro quota.",
        "Check the job in ChatGPT if this persists.",
      ],
    });
  }

  return completedText;
}

interface ResponseParseState {
  acceptsTextContinuation: boolean;
  lastAppendText: string | null;
}

function mergeStreamText(current: string | null, next: string, append: boolean): string {
  if (append) {
    if (current && (current === next || current.endsWith(next))) return current;
    return `${current ?? ""}${next}`;
  }
  if (current && current.length > next.length && current.endsWith(next)) return current;
  return next;
}

export function extractLimitsProgress(event: unknown): LimitsObservation[] {
  if (!isRecord(event)) return [];
  const candidates: unknown[] = [];
  if (event.type === "conversation_detail_metadata") candidates.push(event);
  const value = event.v;
  if (isRecord(value) && value.type === "conversation_detail_metadata") candidates.push(value);
  if (isRecord(value) && Array.isArray((value as { limits_progress?: unknown }).limits_progress)) {
    candidates.push(value);
  }
  if (Array.isArray((event as { limits_progress?: unknown }).limits_progress)) candidates.push(event);

  const observations: LimitsObservation[] = [];
  const seen = new Set<string>();
  for (const candidate of candidates) {
    if (!isRecord(candidate)) continue;
    const progress = (candidate as { limits_progress?: unknown }).limits_progress;
    if (!Array.isArray(progress)) continue;
    for (const entry of progress) {
      if (!isRecord(entry)) continue;
      const featureName = entry.feature_name;
      const remaining = entry.remaining;
      const resetAfter = entry.reset_after;
      if (typeof featureName !== "string" || typeof remaining !== "number") continue;
      if (seen.has(featureName)) continue;
      seen.add(featureName);
      observations.push({
        feature_name: featureName,
        remaining,
        reset_after: typeof resetAfter === "string" ? resetAfter : null,
      });
    }
  }
  return observations;
}

function readConversationEvent(event: unknown, state: ResponseParseState): {
  text: string | null;
  completed: boolean;
  append: boolean;
} {
  if (!isRecord(event)) return { text: null, completed: false, append: false };
  if (event.type === "error") {
    throw new ProError("UPSTREAM_ERROR", readErrorMessage(event), {
      exitCode: EXIT.upstream,
      suggestions: ["Retry later or check usage limits."],
    });
  }
  const patchText = readPatchAppendText(event, state);
  if (patchText !== null) {
    return {
      text: patchText,
      append: true,
      completed: event.type === "done" || event.type === "message_stream_complete",
    };
  }
  const messageText = readConversationMessageText(event);
  return {
    text: messageText,
    append: false,
    completed: event.type === "done" || event.type === "message_stream_complete" || isConversationMessageDone(event),
  };
}

function parseSseFrame(frame: string): unknown {
  const data = frame
    .split("\n")
    .filter((line) => line.startsWith("data:"))
    .map((line) => line.slice(5).trim())
    .join("\n");
  if (!data) return null;
  if (data === "[DONE]") return { type: "done" };
  return JSON.parse(data) as unknown;
}

function readConversationMessageText(event: Record<string, unknown>): string | null {
  const value = event.v as { message?: unknown } | undefined;
  const message = (event.message ?? value?.message) as { author?: unknown; content?: unknown } | undefined;
  const author = message?.author as { role?: unknown } | undefined;
  if (author?.role !== "assistant") return null;

  const content = message?.content as { parts?: unknown } | undefined;
  if (!Array.isArray(content?.parts)) return null;

  const parts = content.parts.filter((part): part is string => typeof part === "string");
  if (parts.length === 0) return null;
  return parts.join("");
}

function readPatchAppendText(event: Record<string, unknown>, state: ResponseParseState): string | null {
  if (event.o === "append" && isMessageContentPartPath(event.p) && typeof event.v === "string") {
    state.acceptsTextContinuation = true;
    return readNewAppendText(event.v, state);
  }
  if (typeof event.v === "string" && state.acceptsTextContinuation) {
    return readNewAppendText(event.v, state);
  }
  state.acceptsTextContinuation = false;
  state.lastAppendText = null;
  if (event.o !== "patch" || !Array.isArray(event.v)) return null;
  const chunks = event.v
    .filter((patch): patch is { o: unknown; p: unknown; v: unknown } => Boolean(patch) && typeof patch === "object")
    .filter(
      (patch) =>
        patch.o === "append" &&
        isMessageContentPartPath(patch.p) &&
        typeof patch.v === "string",
    )
    .map((patch) => patch.v);
  if (chunks.length === 0) return null;
  state.acceptsTextContinuation = true;
  return readNewAppendText(chunks.join(""), state);
}

function isMessageContentPartPath(path: unknown): boolean {
  return typeof path === "string" && /^\/message\/content\/parts\/\d+$/.test(path);
}

function readNewAppendText(text: string, state: ResponseParseState): string | null {
  if (text === state.lastAppendText) return null;
  state.lastAppendText = text;
  return text;
}

function isRecord(value: unknown): value is Record<string, unknown> {
  return Boolean(value) && typeof value === "object";
}

function isConversationMessageDone(event: Record<string, unknown>): boolean {
  const value = event.v as { message?: unknown } | undefined;
  const message = (event.message ?? value?.message) as { status?: unknown; end_turn?: unknown } | undefined;
  return message?.status === "finished_successfully" || message?.end_turn === true;
}

function readErrorMessage(event: Record<string, unknown>): string {
  if (typeof event.error === "string") return event.error;
  const error = event.error as { message?: unknown } | undefined;
  return typeof error?.message === "string" ? error.message : "ChatGPT backend returned an error event.";
}

function normalizeReasoning(reasoning: string): string {
  if (isReasoningLevel(reasoning)) return reasoning;
  throw new ProError("INVALID_REASONING", `Unsupported reasoning level ${reasoning}.`, {
    exitCode: EXIT.invalidArgs,
    suggestions: ["Use min, standard, extended, or max."],
  });
}

function normalizeModel(model: string): string {
  const value = canonicalModelId(model) || DEFAULT_MODEL;
  if (value === "auto") {
    throw new ProError("INVALID_MODEL", "The model auto is not supported.", {
      exitCode: EXIT.invalidArgs,
      suggestions: ["Use a concrete model id such as gpt-5-5-pro, gpt-4-5, or research."],
    });
  }
  return value;
}

function stringOption(value: unknown): string | undefined {
  return typeof value === "string" && value.length > 0 ? value : undefined;
}

function integerOption(
  value: unknown,
  fallback: number | undefined,
  min: number,
  max: number,
): number | undefined {
  if (value === undefined || value === null || value === "") return fallback;
  const number = Math.trunc(Number(value));
  if (!Number.isFinite(number)) return fallback;
  return Math.min(max, Math.max(min, number));
}

function booleanOption(value: unknown, fallback: boolean): boolean {
  if (typeof value === "boolean") return value;
  if (typeof value !== "string") return fallback;
  if (["1", "true", "yes", "on"].includes(value.toLowerCase())) return true;
  if (["0", "false", "no", "off"].includes(value.toLowerCase())) return false;
  return fallback;
}

function networkError(error: unknown): ProError {
  return new ProError("NETWORK_ERROR", "ChatGPT backend request failed before a response.", {
    exitCode: EXIT.network,
    suggestions: ["Check connectivity and retry.", "Run pro-cli auth status --json if this persists."],
    cause: error,
  });
}

function isRetryable(error: ProError): boolean {
  if (["NETWORK_ERROR", "REQUEST_TIMEOUT", "STREAM_INCOMPLETE"].includes(error.code)) return true;
  const status = error.details?.status;
  return typeof status === "number" && (status === 408 || status === 429 || status >= 500);
}

function withAttemptDetails(error: ProError, attempts: number): ProError {
  return new ProError(error.code, error.message, {
    exitCode: error.exitCode,
    suggestions: error.suggestions,
    details: { ...(error.details ?? {}), attempts },
    cause: error,
  });
}

async function sleep(ms: number): Promise<void> {
  await new Promise((resolve) => setTimeout(resolve, ms));
}
`````

## File: src/update.ts
`````typescript
import { spawnSync } from "node:child_process";
import { dirname, resolve } from "node:path";
import { fileURLToPath } from "node:url";
import { EXIT, ProError } from "./errors";

const REPO_URL = "https://github.com/ratacat/pro-cli.git";
const ALLOWED_ORIGINS = new Set([
  REPO_URL,
  "https://github.com/ratacat/pro-cli",
  "git@github.com:ratacat/pro-cli.git",
]);

export interface UpdateStep {
  command: string;
  output: string;
}

export interface UpdateResult {
  repoRoot: string;
  branch: string;
  version: string;
  steps: UpdateStep[];
}

export interface UpdateOptions {
  repoRoot?: string;
  runCommand?: (command: string, args: string[], cwd?: string) => UpdateStep;
}

export function updateProCli(options: UpdateOptions = {}): UpdateResult {
  const repoRoot = options.repoRoot ?? resolve(dirname(fileURLToPath(import.meta.url)), "..");
  const runCommand = options.runCommand ?? run;
  const origin = runCommand("git", ["-C", repoRoot, "remote", "get-url", "origin"]).output.trim();
  if (!ALLOWED_ORIGINS.has(origin)) {
    throw new ProError("UPDATE_WRONG_ORIGIN", `Refusing to update checkout with origin ${origin}.`, {
      exitCode: EXIT.invalidArgs,
      suggestions: [`Expected ${REPO_URL}.`],
    });
  }

  const branch = runCommand("git", ["-C", repoRoot, "branch", "--show-current"]).output.trim();
  if (!branch) {
    throw new ProError("UPDATE_DETACHED_HEAD", "Refusing to update a detached checkout.", {
      exitCode: EXIT.invalidArgs,
      suggestions: ["Check out main, then rerun pro-cli update."],
    });
  }
  if (branch !== "main") {
    throw new ProError("UPDATE_WRONG_BRANCH", `Refusing to update branch ${branch}.`, {
      exitCode: EXIT.invalidArgs,
      suggestions: ["Switch to main, then rerun pro-cli update."],
    });
  }

  const dirty = runCommand("git", ["-C", repoRoot, "status", "--porcelain"]).output.trim();
  if (dirty) {
    throw new ProError("UPDATE_DIRTY_WORKTREE", "Refusing to update with uncommitted changes.", {
      exitCode: EXIT.invalidArgs,
      suggestions: ["Commit or stash changes, then rerun pro-cli update."],
    });
  }

  const steps = [
    runCommand("git", ["-C", repoRoot, "pull", "--ff-only", "origin", "main"]),
    runCommand("bun", ["install"], repoRoot),
    runCommand("bun", ["link"], repoRoot),
    runCommand("pro-cli", ["daemon", "restart", "--json"]),
  ];
  const version = runCommand("pro-cli", ["--version"]).output.trim();

  return { repoRoot, branch, version, steps };
}

function run(command: string, args: string[], cwd?: string): UpdateStep {
  const result = spawnSync(command, args, {
    cwd,
    encoding: "utf8",
    shell: false,
  });
  const renderedCommand = [command, ...args].join(" ");
  if (result.error || result.status !== 0) {
    throw new ProError("UPDATE_COMMAND_FAILED", `${renderedCommand} failed.`, {
      exitCode: EXIT.internal,
      suggestions: ["Inspect the command output, fix the checkout, then rerun pro-cli update."],
      details: {
        command: renderedCommand,
        status: result.status,
        stderr: compact(result.stderr || result.error?.message || ""),
        stdout: compact(result.stdout || ""),
      },
      cause: result.error,
    });
  }
  return {
    command: renderedCommand,
    output: compact(result.stdout || result.stderr || ""),
  };
}

function compact(value: string): string {
  return value.trim().replace(/\s+/g, " ").slice(0, 500);
}
`````

## File: tests/args.test.ts
`````typescript
import { describe, expect, test } from "bun:test";
import { ProError } from "../src/errors";
import { flagBoolean, flagString, flagStrings, parseArgs } from "../src/args";

describe("parseArgs: positionals", () => {
  test("collects bare tokens as positionals in order", () => {
    const parsed = parseArgs(["job", "create", "hello", "world"]);
    expect(parsed.positionals).toEqual(["job", "create", "hello", "world"]);
    expect(parsed.flags.size).toBe(0);
  });

  test("uses -- as an end-of-flags delimiter", () => {
    const parsed = parseArgs(["ask", "--", "--literal-prompt", "--json"]);
    expect(parsed.positionals).toEqual(["ask", "--literal-prompt", "--json"]);
    expect(parsed.flags.size).toBe(0);
  });

  test("a bare -- by itself just ends flag parsing", () => {
    const parsed = parseArgs(["--"]);
    expect(parsed.positionals).toEqual([]);
    expect(parsed.flags.size).toBe(0);
  });

  test("a single-dash token is a positional, not a flag", () => {
    const parsed = parseArgs(["ask", "-"]);
    expect(parsed.positionals).toEqual(["ask", "-"]);
  });
});

describe("parseArgs: boolean flags", () => {
  test("known boolean flags become true without consuming the next token", () => {
    const parsed = parseArgs(["--json", "extra-positional"]);
    expect(parsed.flags.get("json")).toBe(true);
    expect(parsed.positionals).toEqual(["extra-positional"]);
  });

  test("each registered boolean flag is recognized", () => {
    // Lock down the registered boolean set; if someone removes a flag from
    // BOOLEAN_FLAGS, downstream commands break with confusing
    // "Missing value" errors.
    const flags = [
      "json",
      "no-json",
      "dry-run",
      "include-expired",
      "no-start",
      "save",
      "temporary",
      "no-temporary",
      "wait",
      "help",
      "version",
      "allow-fifty",
      "no-launch",
      "no-backup",
    ];
    for (const name of flags) {
      const parsed = parseArgs([`--${name}`]);
      expect(parsed.flags.get(name)).toBe(true);
    }
  });

  test("repeating a known boolean flag still reads as true", () => {
    const parsed = parseArgs(["--json", "--json"]);
    expect(flagBoolean(parsed.flags, "json")).toBe(true);
  });
});

describe("parseArgs: value flags", () => {
  test("space-separated value flag consumes the next token", () => {
    const parsed = parseArgs(["--model", "gpt-5-5-pro"]);
    expect(parsed.flags.get("model")).toBe("gpt-5-5-pro");
  });

  test("equals-syntax keeps everything after the first = as the value", () => {
    const parsed = parseArgs(["--instructions=Use=equal=signs"]);
    expect(parsed.flags.get("instructions")).toBe("Use=equal=signs");
  });

  test("equals-syntax preserves whitespace inside the value", () => {
    const parsed = parseArgs(["--instructions=line one\nline two"]);
    expect(parsed.flags.get("instructions")).toBe("line one\nline two");
  });

  test("missing value for non-boolean flag throws INVALID_ARGS", () => {
    try {
      parseArgs(["--model"]);
      throw new Error("Expected INVALID_ARGS.");
    } catch (error) {
      expect(error).toBeInstanceOf(ProError);
      const proError = error as ProError;
      expect(proError.code).toBe("INVALID_ARGS");
      expect(proError.message).toContain("Missing value for --model");
    }
  });

  test("a flag immediately followed by another flag throws INVALID_ARGS", () => {
    // Regression guard: if argv is "--model --json", we must not silently
    // treat "--json" as the model name.
    try {
      parseArgs(["--model", "--json"]);
      throw new Error("Expected INVALID_ARGS.");
    } catch (error) {
      expect(error).toBeInstanceOf(ProError);
      expect((error as ProError).code).toBe("INVALID_ARGS");
    }
  });
});

describe("parseArgs: invalid input", () => {
  test("an empty equals-name throws INVALID_ARGS", () => {
    try {
      parseArgs(["--=value"]);
      throw new Error("Expected INVALID_ARGS.");
    } catch (error) {
      expect(error).toBeInstanceOf(ProError);
      // Either "Empty flag" or "Invalid flag" — both signal the bad input.
      const message = (error as ProError).message;
      expect(message.toLowerCase()).toContain("flag");
    }
  });
});

describe("parseArgs: repeated flags", () => {
  test("repeating a flag converts the entry to an array of values in order", () => {
    const parsed = parseArgs([
      "--instructions",
      "first",
      "--instructions",
      "second",
      "--instructions",
      "third",
    ]);
    const value = parsed.flags.get("instructions");
    expect(Array.isArray(value)).toBe(true);
    expect(value as string[]).toEqual(["first", "second", "third"]);
  });

  test("flagStrings exposes repeated values as an array", () => {
    const parsed = parseArgs(["--instructions", "a", "--instructions", "b"]);
    expect(flagStrings(parsed.flags, "instructions")).toEqual(["a", "b"]);
  });

  test("flagStrings returns single-value flags as a one-element array", () => {
    const parsed = parseArgs(["--model", "gpt-5-5-pro"]);
    expect(flagStrings(parsed.flags, "model")).toEqual(["gpt-5-5-pro"]);
  });

  test("flagStrings returns an empty array when the flag is absent", () => {
    expect(flagStrings(parseArgs([]).flags, "model")).toEqual([]);
  });
});

describe("flagString helper", () => {
  test("returns the string value", () => {
    const parsed = parseArgs(["--model", "gpt-5-5-pro"]);
    expect(flagString(parsed.flags, "model")).toBe("gpt-5-5-pro");
  });

  test("returns undefined when absent", () => {
    expect(flagString(parseArgs([]).flags, "model")).toBeUndefined();
  });

  test("throws for repeated flags so single-value reads cannot fall back to defaults", () => {
    const parsed = parseArgs(["--model", "a", "--model", "b"]);
    expect(() => flagString(parsed.flags, "model")).toThrow(ProError);
  });

  test("coerces a true boolean to the literal string 'true'", () => {
    const parsed = parseArgs(["--json"]);
    expect(flagString(parsed.flags, "json")).toBe("true");
  });
});

describe("flagBoolean helper", () => {
  test("returns true only when the flag was set as a boolean", () => {
    const parsed = parseArgs(["--json"]);
    expect(flagBoolean(parsed.flags, "json")).toBe(true);
  });

  test("returns false when the flag is absent", () => {
    expect(flagBoolean(parseArgs([]).flags, "json")).toBe(false);
  });

  test("returns false when the flag was set with a string value (e.g. --json=false)", () => {
    // The current parser stores explicit values as strings; flagBoolean is
    // only true when no value was provided. This is what makes the
    // BOOLEAN_FLAGS set meaningful.
    const parsed = parseArgs(["--json=false"]);
    expect(flagBoolean(parsed.flags, "json")).toBe(false);
  });
});
`````

## File: tests/cdp.test.ts
`````typescript
import { afterEach, describe, expect, test } from "bun:test";
import { ProError } from "../src/errors";
import {
  callBrowserCdp,
  evaluateInCdpPage,
  findChatGptTargetId,
  getCookiesFromCdp,
  pruneVolatileCookiesFromCdp,
  recoverCookieBloatInCdp,
} from "../src/cdp";

const originalFetch = globalThis.fetch;
const originalWebSocket = globalThis.WebSocket;

afterEach(() => {
  globalThis.fetch = originalFetch;
  globalThis.WebSocket = originalWebSocket;
});

describe("CDP helpers", () => {
  test("falls back to browser-level Storage.getCookies when no page target is listed", async () => {
    const methods: string[] = [];
    installFakeCdp({
      pageTargets: [],
      onCommand(method) {
        methods.push(method);
        if (method === "Network.getCookies") {
          return { error: { code: -32601, message: "'Network.getCookies' wasn't found" } };
        }
        if (method === "Storage.getCookies") {
          return {
            result: {
              cookies: [
                cookie("__cf_bm", "chatgpt.com"),
                cookie("unrelated", "example.com"),
              ],
            },
          };
        }
        return { result: {} };
      },
    });

    const cookies = await getCookiesFromCdp("http://127.0.0.1:9222", ["https://chatgpt.com/"]);

    expect(methods).toEqual(["Network.getCookies", "Storage.getCookies"]);
    expect(cookies.map((stored) => stored.name)).toEqual(["__cf_bm"]);
  });

  test("page evaluation fails clearly when CDP has no inspectable page target", async () => {
    installFakeCdp({
      pageTargets: [],
      onCommand() {
        return { result: {} };
      },
    });

    try {
      await evaluateInCdpPage("http://127.0.0.1:9222", "location.href");
      throw new Error("Expected CHATGPT_PAGE_MISSING.");
    } catch (error) {
      // Strengthen: assert exact ProError code + actionable suggestions, not
      // just substring match on the message (which couples to copy text).
      expect(error).toBeInstanceOf(ProError);
      const proError = error as ProError;
      expect(proError.code).toBe("CHATGPT_PAGE_MISSING");
      expect(proError.suggestions.some((s) => s.includes("auth command"))).toBe(true);
      expect(proError.details?.cdpBase).toBe("http://127.0.0.1:9222");
    }
  });

  test("prunes volatile conversation cookies from a live CDP profile", async () => {
    const deleted: Array<Record<string, unknown> | undefined> = [];
    installFakeCdp({
      pageTargets: [{ type: "page", url: "https://chatgpt.com/", webSocketDebuggerUrl: "ws://fake-page" }],
      onCommand(method, params) {
        if (method === "Network.getCookies") {
          return {
            result: {
              cookies: [
                cookie("__Secure-next-auth.session-token", "chatgpt.com"),
                cookie("conv_key_abc", "chatgpt.com"),
                cookie("conv_key_def", "chatgpt.com"),
              ],
            },
          };
        }
        if (method === "Network.deleteCookies") {
          deleted.push(params);
          return { result: {} };
        }
        return { result: {} };
      },
    });

    const result = await pruneVolatileCookiesFromCdp("http://127.0.0.1:9222", ["https://chatgpt.com/"]);

    expect(result.checked).toBe(3);
    expect(result.deleted).toBe(2);
    expect(deleted).toEqual([
      { name: "conv_key_abc", domain: "chatgpt.com", path: "/" },
      { name: "conv_key_def", domain: "chatgpt.com", path: "/" },
    ]);
  });

  test("cookie bloat recovery prunes volatile cookies and reloads ChatGPT", async () => {
    const methods: string[] = [];
    installFakeCdp({
      pageTargets: [{ type: "page", url: "https://chatgpt.com/", webSocketDebuggerUrl: "ws://fake-page" }],
      onCommand(method) {
        methods.push(method);
        if (method === "Network.getCookies") {
          return {
            result: {
              cookies: [
                cookie("__Secure-next-auth.session-token", "chatgpt.com"),
                cookie("conv_key_abc", "chatgpt.com"),
              ],
            },
          };
        }
        return { result: {} };
      },
    });

    const result = await recoverCookieBloatInCdp("http://127.0.0.1:9222", ["https://chatgpt.com/"], 10);

    expect(result.deleted).toBe(1);
    expect(result.navigated).toBe(true);
    expect(methods).toEqual([
      "Network.getCookies",
      "Network.deleteCookies",
      "Page.enable",
      "Page.navigate",
    ]);
  });

  test("evaluateInCdpPage returns the value from Runtime.evaluate", async () => {
    installFakeCdp({
      pageTargets: [{ type: "page", url: "https://chatgpt.com/", webSocketDebuggerUrl: "ws://fake-page" }],
      onCommand(method) {
        if (method === "Runtime.evaluate") {
          return { result: { result: { value: { ok: true, status: 200, payload: "hello" } } } };
        }
        return { result: {} };
      },
    });

    const result = await evaluateInCdpPage<{ ok: boolean; status: number; payload: string }>(
      "http://127.0.0.1:9222",
      "(() => ({ ok: true, status: 200, payload: 'hello' }))()",
    );
    expect(result).toEqual({ ok: true, status: 200, payload: "hello" });
  });

  test("evaluateInCdpPage surfaces page-side exceptions as CDP_EVALUATION_FAILED", async () => {
    installFakeCdp({
      pageTargets: [{ type: "page", url: "https://chatgpt.com/", webSocketDebuggerUrl: "ws://fake-page" }],
      onCommand(method) {
        if (method === "Runtime.evaluate") {
          return {
            result: {
              exceptionDetails: { text: "ReferenceError: undefinedThing is not defined" },
            },
          };
        }
        return { result: {} };
      },
    });

    try {
      await evaluateInCdpPage("http://127.0.0.1:9222", "throw new Error('boom')");
      throw new Error("Expected CDP_EVALUATION_FAILED.");
    } catch (error) {
      expect(error).toBeInstanceOf(ProError);
      const proError = error as ProError;
      expect(proError.code).toBe("CDP_EVALUATION_FAILED");
      expect(proError.message).toContain("ReferenceError");
    }
  });

  test("evaluateInCdpPage surfaces CDP method errors via CDP_COMMAND_FAILED", async () => {
    installFakeCdp({
      pageTargets: [{ type: "page", url: "https://chatgpt.com/", webSocketDebuggerUrl: "ws://fake-page" }],
      onCommand(method) {
        if (method === "Runtime.evaluate") {
          return { error: { code: -32000, message: "Cannot find context" } };
        }
        return { result: {} };
      },
    });

    try {
      await evaluateInCdpPage("http://127.0.0.1:9222", "noop");
      throw new Error("Expected CDP_COMMAND_FAILED.");
    } catch (error) {
      expect(error).toBeInstanceOf(ProError);
      expect((error as ProError).code).toBe("CDP_COMMAND_FAILED");
    }
  });

  test("cookie bloat recovery does NOT navigate when no volatile cookies exist", async () => {
    // Regression guard: blindly reloading the tab on every cookie scan would
    // disrupt logged-in state and waste time. Recovery must be a no-op when
    // there is nothing to prune.
    const methods: string[] = [];
    installFakeCdp({
      pageTargets: [{ type: "page", url: "https://chatgpt.com/", webSocketDebuggerUrl: "ws://fake-page" }],
      onCommand(method) {
        methods.push(method);
        if (method === "Network.getCookies") {
          return {
            result: {
              cookies: [cookie("__Secure-next-auth.session-token", "chatgpt.com")],
            },
          };
        }
        return { result: {} };
      },
    });

    const result = await recoverCookieBloatInCdp("http://127.0.0.1:9222", ["https://chatgpt.com/"], 10);

    expect(result.deleted).toBe(0);
    expect(result.navigated).toBe(false);
    expect(methods).toEqual(["Network.getCookies"]);
    expect(methods).not.toContain("Page.navigate");
    expect(methods).not.toContain("Page.enable");
  });

  test("pruneVolatileCookiesFromCdp filters out cookies on non-target domains before deleting", async () => {
    // The URL filter (cookieAppliesToUrl) must drop unrelated cookies so we
    // never delete cookies for sites we don't own.
    const deleted: Array<{ name?: string; domain?: string; path?: string }> = [];
    installFakeCdp({
      pageTargets: [{ type: "page", url: "https://chatgpt.com/", webSocketDebuggerUrl: "ws://fake-page" }],
      onCommand(method, params) {
        if (method === "Network.getCookies") {
          return {
            result: {
              cookies: [
                cookie("conv_key_chatgpt", "chatgpt.com"),
                cookie("conv_key_other", "evil.example"),
              ],
            },
          };
        }
        if (method === "Network.deleteCookies") {
          deleted.push(params as { name?: string; domain?: string });
          return { result: {} };
        }
        return { result: {} };
      },
    });

    const result = await pruneVolatileCookiesFromCdp("http://127.0.0.1:9222", ["https://chatgpt.com/"]);
    expect(result.checked).toBe(1); // only the chatgpt.com cookie matched the URL filter
    expect(result.deleted).toBe(1);
    expect(deleted).toEqual([{ name: "conv_key_chatgpt", domain: "chatgpt.com", path: "/" }]);
  });

  test("cookie URL filtering respects path segment boundaries", async () => {
    installFakeCdp({
      pageTargets: [],
      onCommand(method) {
        if (method === "Network.getCookies") {
          return {
            result: {
              cookies: [
                { ...cookie("root", "chatgpt.com"), path: "/" },
                { ...cookie("backend-api", "chatgpt.com"), path: "/backend-api" },
                { ...cookie("too-broad", "chatgpt.com"), path: "/backend" },
                { ...cookie("wrong-sibling", "chatgpt.com"), path: "/backend-api-v2" },
              ],
            },
          };
        }
        return { result: {} };
      },
    });

    const cookies = await getCookiesFromCdp("http://127.0.0.1:9222", [
      "https://chatgpt.com/backend-api/f/conversation",
    ]);

    expect(cookies.map((stored) => `${stored.name}:${stored.path}`).sort()).toEqual([
      "backend-api:/backend-api",
      "root:/",
    ]);
  });

  test("findChatGptTargetId returns the id of a chatgpt.com page target", async () => {
    installFakeCdp({
      pageTargets: [
        { type: "page", url: "https://example.com/", webSocketDebuggerUrl: "ws://fake-other" },
        { type: "page", url: "https://chatgpt.com/c/abc", webSocketDebuggerUrl: "ws://fake-chatgpt" },
      ],
      onCommand() {
        return { result: {} };
      },
    });
    // installFakeCdp does not surface ids on its targets; verify a chatgpt.com
    // tab is selectable. We extend the fake to return ids when /json is fetched.
    globalThis.fetch = (async (url: string | URL | Request) => {
      const target = String(url);
      if (target.endsWith("/json")) {
        return Response.json([
          { id: "TAB1", type: "page", url: "https://example.com/", webSocketDebuggerUrl: "ws://fake-other" },
          { id: "TAB2", type: "page", url: "https://chatgpt.com/c/abc", webSocketDebuggerUrl: "ws://fake-chatgpt" },
        ]);
      }
      if (target.endsWith("/json/version")) {
        return Response.json({ webSocketDebuggerUrl: "ws://fake-browser" });
      }
      return new Response("nope", { status: 500 });
    }) as unknown as typeof fetch;

    const id = await findChatGptTargetId("http://127.0.0.1:9222");
    expect(id).toBe("TAB2");
  });

  test("findChatGptTargetId returns null when no chatgpt.com tab exists", async () => {
    globalThis.fetch = (async (url: string | URL | Request) => {
      const target = String(url);
      if (target.endsWith("/json")) {
        return Response.json([
          { id: "TAB1", type: "page", url: "https://example.com/", webSocketDebuggerUrl: "ws://fake-other" },
        ]);
      }
      return new Response("nope", { status: 500 });
    }) as unknown as typeof fetch;

    const id = await findChatGptTargetId("http://127.0.0.1:9222");
    expect(id).toBeNull();
  });

  test("findChatGptTargetId returns null on CDP fetch errors (caller decides what to do)", async () => {
    globalThis.fetch = (async () => {
      throw new Error("ECONNREFUSED");
    }) as unknown as typeof fetch;

    const id = await findChatGptTargetId("http://127.0.0.1:9222");
    expect(id).toBeNull();
  });

  test("callBrowserCdp dispatches the requested method to the browser endpoint", async () => {
    let observedMethod = "";
    let observedParams: Record<string, unknown> | undefined;
    installFakeCdp({
      pageTargets: [],
      onCommand(method, params) {
        observedMethod = method;
        observedParams = params;
        if (method === "Browser.getWindowForTarget") {
          return { result: { windowId: 12345, bounds: { left: 10, top: 20, width: 800, height: 600 } } };
        }
        return { result: {} };
      },
    });

    const result = await callBrowserCdp<{ windowId: number; bounds: Record<string, number> }>(
      "http://127.0.0.1:9222",
      "Browser.getWindowForTarget",
      { targetId: "TAB1" },
    );
    expect(observedMethod).toBe("Browser.getWindowForTarget");
    expect(observedParams).toEqual({ targetId: "TAB1" });
    expect(result.windowId).toBe(12345);
    expect(result.bounds).toEqual({ left: 10, top: 20, width: 800, height: 600 });
  });
});

function cookie(name: string, domain: string): {
  name: string;
  value: string;
  domain: string;
  path: string;
  secure: boolean;
} {
  return { name, value: "redacted-test-value", domain, path: "/", secure: true };
}

function installFakeCdp(options: {
  pageTargets: Array<{ type: string; url: string; webSocketDebuggerUrl: string }>;
  onCommand: (
    method: string,
    params?: Record<string, unknown>,
  ) => { result?: unknown; error?: { code: number; message: string } };
}): void {
  globalThis.fetch = (async (url: string | URL | Request) => {
    const target = String(url);
    if (target.endsWith("/json")) {
      return Response.json(options.pageTargets);
    }
    if (target.endsWith("/json/version")) {
      return Response.json({ webSocketDebuggerUrl: "ws://fake-browser" });
    }
    return new Response("unexpected fetch", { status: 500 });
  }) as unknown as typeof fetch;

  class FakeWebSocket extends EventTarget {
    constructor(_url: string) {
      super();
      queueMicrotask(() => this.dispatchEvent(new Event("open")));
    }

    send(raw: string): void {
      const message = JSON.parse(raw) as {
        id: number;
        method: string;
        params?: Record<string, unknown>;
      };
      const response = { id: message.id, ...options.onCommand(message.method, message.params) };
      queueMicrotask(() =>
        this.dispatchEvent(new MessageEvent("message", { data: JSON.stringify(response) })),
      );
    }

    close(): void {
      this.dispatchEvent(new Event("close"));
    }
  }

  globalThis.WebSocket = FakeWebSocket as unknown as typeof WebSocket;
}
`````

## File: tests/cli.test.ts
`````typescript
import { access, chmod, mkdir, mkdtemp, readFile, rm, writeFile } from "node:fs/promises";
import { basename, join } from "node:path";
import { tmpdir } from "node:os";
import { afterEach, describe, expect, test } from "bun:test";
import { runCli } from "../src/app";
import { EXIT, ProError } from "../src/errors";
import { JobStore } from "../src/jobs";

const originalFetch = globalThis.fetch;
const originalWebSocket = globalThis.WebSocket;

afterEach(() => {
  globalThis.fetch = originalFetch;
  globalThis.WebSocket = originalWebSocket;
});

interface RunResult {
  code: number;
  stdout: string;
  stderr: string;
}

async function withHome<T>(fn: (home: string) => Promise<T>): Promise<T> {
  const home = await mkdtemp(join(tmpdir(), "pro-query-test-"));
  try {
    return await fn(home);
  } finally {
    await rm(home, { recursive: true, force: true });
  }
}

async function run(args: string[], options: { tty?: boolean; home?: string } = {}): Promise<RunResult> {
  let stdout = "";
  let stderr = "";
  const code = await runCli(args, {
    stdout: (text) => {
      stdout += text;
    },
    stderr: (text) => {
      stderr += text;
    },
    stdoutIsTTY: options.tty ?? false,
    env: { PRO_CLI_HOME: options.home },
    cwd: process.cwd(),
  });
  return { code, stdout, stderr };
}

describe("robot-mode CLI", () => {
  test("prints compact help with no args for TTY users", async () => {
    const result = await run([], { tty: true });

    expect(result.code).toBe(0);
    expect(result.stdout).toContain("pro-cli: ChatGPT Pro CLI");
    expect(result.stdout).toContain("ask: direct blocking query");
    expect(result.stdout).toContain("job create --wait: durable blocking query");
    expect(result.stdout).toContain("job wait: waits until done");
    expect(result.stdout).toContain("update: fast-forward install");
    expect(result.stdout.length).toBeLessThan(260);
    expect(result.stderr).toBe("");
  });

  test("prints compact help with --help for TTY users", async () => {
    const result = await run(["--help"], { tty: true });

    expect(result.code).toBe(0);
    expect(result.stdout).toContain("pro-cli: ChatGPT Pro CLI");
    expect(result.stdout).toContain("ask: direct blocking query");
    expect(result.stdout).toContain("job create --wait: durable blocking query");
    expect(result.stdout).toContain("job wait: waits until done");
    expect(result.stdout).toContain("update: fast-forward install");
    expect(result.stdout.length).toBeLessThan(260);
    expect(result.stderr).toBe("");
  });

  test("auto-switches to JSON when stdout is not a TTY", async () => {
    const result = await run([], { tty: false });

    expect(result.code).toBe(0);
    const payload = JSON.parse(result.stdout);
    expect(payload.ok).toBe(true);
    expect(payload.data.text).toContain("ask: direct blocking query");
    expect(payload.data.commands).toContain("update");
    expect(payload.data.commands).toContain("auth capture");
  });

  test("prints version for install verification", async () => {
    const result = await run(["--version"], { tty: false });

    expect(result.code).toBe(0);
    expect(result.stdout).toMatch(/^pro-cli \d+\.\d+\.\d+\n$/);
    expect(result.stderr).toBe("");
  });

  test("emits structured JSON errors and invalid-args exit code", async () => {
    const result = await run(["missing-command", "--json"], { tty: true });

    expect(result.code).toBe(2);
    expect(result.stdout).toBe("");
    const payload = JSON.parse(result.stderr);
    expect(payload.ok).toBe(false);
    expect(payload.error.code).toBe("INVALID_ARGS");
    expect(payload.error.suggestions).toContain("Run pro-cli help.");
  });

  test("setup gives a safe first-run path without secrets", async () => {
    await withHome(async (home) => {
      const result = await run(["setup", "--json"], { tty: true, home });

      expect(result.code).toBe(0);
      expect(result.stdout).not.toContain("secret");
      const payload = JSON.parse(result.stdout);
      expect(payload.data.ready).toBe(false);
      expect(payload.data.steps.map((step: { id: string }) => step.id)).toEqual([
        "install",
        "open-chatgpt",
        "capture-auth",
        "doctor",
      ]);
      expect(payload.data.steps[1].command).toContain("chrome-profile");
      expect(payload.data.steps[2].command).toContain("pro-cli auth capture");
      expect(payload.data.steps[3].command).toContain("pro-cli doctor");
      expect(payload.data.safety.rawValuesPrinted).toBe(false);
      expect(JSON.stringify(payload.data)).not.toMatch(/Reply with OK|smoke-test/i);
    });
  });

  test("setup is readable for TTY users", async () => {
    await withHome(async (home) => {
      const result = await run(["setup"], { tty: true, home });

      expect(result.code).toBe(0);
      expect(result.stdout).toContain("pro-cli needs a logged-in ChatGPT browser session");
      expect(result.stdout).toContain("[todo] open-chatgpt");
      expect(result.stdout).toContain("pro-cli auth capture");
      expect(result.stdout).toContain("pro-cli doctor");
      expect(result.stdout).not.toMatch(/Reply with OK|smoke-test/i);
      expect(result.stdout).not.toContain("{\"ready\"");
    });
  });

  test("auth command prints dedicated profile launch and capture commands", async () => {
    await withHome(async (home) => {
      const result = await run(["auth", "command", "--port", "9333", "--json"], {
        tty: true,
        home,
      });

      expect(result.code).toBe(0);
      const payload = JSON.parse(result.stdout);
      expect(payload.data.command).toContain("chrome-profile");
      expect(payload.data.command).toContain("9333");
      expect(payload.data.captureCommand).toBe("pro-cli auth capture --cdp http://127.0.0.1:9333 --json");
      expect(payload.data.safety).toContain("dedicated");
    });
  });

  test("reports missing auth without raw cookie values", async () => {
    await withHome(async (home) => {
      const result = await run(["auth", "status", "--json"], { tty: true, home });

      expect(result.code).toBe(0);
      const payload = JSON.parse(result.stdout);
      expect(payload.data.status).toBe("missing");
      expect(payload.data.rawValuesPrinted).toBe(false);
      expect(payload.data.cookieCount).toBe(0);
    });
  });

  test("summarizes existing cookie export without printing values", async () => {
    await withHome(async (home) => {
      const cookiePath = join(home, "cookies", "chatgpt.json");
      await mkdir(join(home, "cookies"), { recursive: true });
      await writeFile(
        cookiePath,
        JSON.stringify({
          version: 1,
          generatedAt: new Date().toISOString(),
          source: "pro-cli-cdp",
          targetUrl: "https://chatgpt.com/",
          origins: ["https://chatgpt.com/"],
          cookies: [
            {
              name: "__Secure-next-auth.session-token",
              value: "secret-value",
              domain: "chatgpt.com",
              path: "/",
              secure: true,
              httpOnly: true,
            },
          ],
        }),
      );

      const result = await run(["auth", "status", "--json"], { tty: true, home });

      expect(result.code).toBe(0);
      expect(result.stdout).not.toContain("secret-value");
      const payload = JSON.parse(result.stdout);
      expect(payload.data.status).toBe("present");
      expect(payload.data.cookieCount).toBe(1);
      expect(payload.data.names).toEqual(["__Secure-next-auth.session-token"]);
    });
  });

  test("creates durable async jobs with redacted prompt preview", async () => {
    await withHome(async (home) => {
      const createdResult = await run(
        ["job", "create", "hello", "from", "agent", "--no-start", "--condensed-response", "250", "--json"],
        {
          tty: true,
          home,
        },
      );

      expect(createdResult.code).toBe(0);
      const created = JSON.parse(createdResult.stdout);
      const jobId = created.data.job.id;
      expect(created.data.job.status).toBe("queued");
      expect(created.data.job.model).toBe("gpt-5-5-pro");
      expect(created.data.job.reasoning).toBe("standard");
      expect(created.data.daemon.started).toBe(false);
      expect(created.data.job.prompt).toBe("");
      expect(created.data.job.promptPreview).toBe("hello from agent");
      expect(created.data.job.options.condensedResponseTokens).toBe(250);

      const status = await run(["job", "status", jobId, "--json"], { tty: true, home });
      expect(status.code).toBe(0);
      expect(JSON.parse(status.stdout).data.job.id).toBe(jobId);
    });
  });

  test("creates durable GPT-4.5 and Deep Research jobs without reasoning effort", async () => {
    await withHome(async (home) => {
      const gpt45 = await run(["job", "create", "hello", "--no-start", "--model", "4.5", "--json"], {
        tty: true,
        home,
      });
      const deepResearch = await run(
        ["job", "create", "hello", "--no-start", "--model", "Deep Research", "--json"],
        { tty: true, home },
      );

      expect(gpt45.code).toBe(0);
      expect(deepResearch.code).toBe(0);
      expect(JSON.parse(gpt45.stdout).data.job).toMatchObject({
        model: "gpt-4-5",
        reasoning: "none",
        options: { temporary: true },
      });
      expect(JSON.parse(deepResearch.stdout).data.job).toMatchObject({
        model: "research",
        reasoning: "standard",
        options: { temporary: false },
      });
    });
  });

  test("ask reports missing session token without durable job storage", async () => {
    await withHome(async (home) => {
      const result = await run(["ask", "hello", "--json"], { tty: true, home });

      expect(result.code).toBe(3);
      expect(result.stdout).toBe("");
      const payload = JSON.parse(result.stderr);
      expect(payload.ok).toBe(false);
      expect(payload.error.code).toBe("SESSION_TOKEN_MISSING");
      await expect(access(join(home, "jobs.sqlite"))).rejects.toThrow();
    });
  });

  test("config permission errors do not fall through to INTERNAL_ERROR", async () => {
    if (process.platform === "win32" || process.getuid?.() === 0) return;
    await withHome(async (home) => {
      const configPath = join(home, "config.json");
      await writeFile(configPath, "{}");
      await chmod(configPath, 0);
      try {
        const result = await run(["doctor", "--json"], { tty: true, home });

        expect(result.code).toBe(3);
        expect(result.stdout).toBe("");
        const payload = JSON.parse(result.stderr);
        expect(payload.error.code).toBe("CONFIG_UNREADABLE");
        expect(payload.error.message).toContain(configPath);
        expect(payload.error.suggestions.join("\n")).toContain("pro-cli doctor --json");
        expect(payload.error.suggestions.join("\n")).toContain("smoke-test");
        expect(payload.error.suggestions.join("\n")).not.toContain("Run with --json");
      } finally {
        await chmod(configPath, 0o600).catch(() => undefined);
      }
    });
  });

  test("daemon status is stopped before startup", async () => {
    await withHome(async (home) => {
      const result = await run(["daemon", "status", "--json"], { tty: true, home });

      expect(result.code).toBe(0);
      const payload = JSON.parse(result.stdout);
      expect(payload.data.daemon.state).toBe("stopped");
      expect(payload.data.daemon.home).toBe(home);
      expect(payload.data.daemon.endpointPath).toContain("pro-cli-");
    });
  });

  test("ask executes without durable job storage and returns the full result", async () => {
    await withHome(async (home) => {
      await mkdir(join(home, "tokens"), { recursive: true });
      await writeFile(
        join(home, "tokens", "chatgpt-session.json"),
        JSON.stringify({
          version: 1,
          generatedAt: new Date().toISOString(),
          source: "pro-cli-cdp-page",
          accessToken: fakeJwt(),
          accountId: "acct_test",
          expiresAt: new Date(Date.now() + 60 * 60 * 1000).toISOString(),
        }),
      );

      let expression = "";
      installFakeCdp(conversationStream("OK"), (script) => {
        expression = script;
      });

      const result = await run(
        [
          "ask",
          "hello",
          "--json",
          "--reasoning",
          "extended",
          "--verbosity",
          "low",
          "--timeout",
          "1000",
          "--retries",
          "1",
          "--retry-delay",
          "0",
        ],
        { tty: true, home },
      );

      expect(result.code).toBe(0);
      const payload = JSON.parse(result.stdout);
      expect(payload.data.job.status).toBe("succeeded");
      expect(payload.data.job.prompt).toBe("");
      expect(payload.data.job.options.temporary).toBe(true);
      expect(payload.data.result).toBe("OK");
      expect(payload.data.agentInstruction).toContain("data.result is the primary deliverable");
      expect(payload.data.agentInstruction).toContain("preserve Pro's prose language");
      expect(payload.data.resultStats).toMatchObject({
        chars: 2,
        approximateTokens: 1,
        fullRelayThresholdChars: 6000,
        fullRelayThresholdApproxTokens: 1500,
      });
      await expect(access(join(home, "jobs.sqlite"))).rejects.toThrow();
      const requestBody = requestBodyFromExpression(expression);
      expect(requestBody.model).toBe("gpt-5-5-pro");
      expect(requestBody.thinking_effort).toBe("extended");
      expect(requestBody.verbosity).toBe("low");
      expect(requestBody.history_and_training_disabled).toBe(true);
      expect(requestBody).not.toHaveProperty("text");
    });
  });

  test("ask supports condensed_response token budget alias", async () => {
    await withHome(async (home) => {
      await writeSessionToken(home);
      let expression = "";
      installFakeCdp(conversationStream("Short answer."), (script) => {
        expression = script;
      });

      const result = await run(["ask", "explain this", "--json", "--condensed_response=500"], {
        tty: true,
        home,
      });

      expect(result.code).toBe(0);
      const payload = JSON.parse(result.stdout);
      expect(payload.data.job.options.condensedResponseTokens).toBe(500);
      const requestBody = requestBodyFromExpression(expression);
      const messages = requestBody.messages as Array<{ content: { parts: string[] } }>;
      const prompt = messages[0].content.parts[0];
      expect(prompt).toContain("Condensed response mode");
      expect(prompt).toContain("approximately 500 tokens or fewer");
      expect(prompt).toContain("explain this");
    });
  });

  test("rejects conflicting condensed response aliases", async () => {
    await withHome(async (home) => {
      const result = await run(
        ["ask", "hello", "--condensed-response", "250", "--condensed_response=500", "--json"],
        { tty: true, home },
      );

      expect(result.code).toBe(2);
      const payload = JSON.parse(result.stderr);
      expect(payload.error.code).toBe("INVALID_ARGS");
      expect(payload.error.message).toContain("one condensed response flag");
    });
  });

  test("defaults to Pro standard reasoning", async () => {
    await withHome(async (home) => {
      await writeSessionToken(home);
      let expression = "";
      installFakeCdp(conversationStream("OK"), (script) => {
        expression = script;
      });

      const result = await run(["ask", "hello", "--json"], { tty: true, home });

      expect(result.code).toBe(0);
      const requestBody = requestBodyFromExpression(expression);
      expect(requestBody.model).toBe("gpt-5-5-pro");
      expect(requestBody.thinking_effort).toBe("standard");
    });
  });

  test("ask connects GPT-4.5 without a thinking effort", async () => {
    await withHome(async (home) => {
      await writeSessionToken(home);
      let expression = "";
      installFakeCdp(conversationStream("OK"), (script) => {
        expression = script;
      });

      const result = await run(["ask", "hello", "--model", "gpt-4.5", "--json"], {
        tty: true,
        home,
      });

      expect(result.code).toBe(0);
      const payload = JSON.parse(result.stdout);
      expect(payload.data.job.model).toBe("gpt-4-5");
      expect(payload.data.job.reasoning).toBe("none");
      const requestBody = requestBodyFromExpression(expression);
      expect(requestBody.model).toBe("gpt-4-5");
      expect(requestBody).not.toHaveProperty("thinking_effort");
    });
  });

  test("ask connects Deep Research without a thinking effort", async () => {
    await withHome(async (home) => {
      await writeSessionToken(home);
      let expression = "";
      installFakeCdp(conversationStream("OK"), (script) => {
        expression = script;
      });

      const result = await run(["ask", "hello", "--model", "deep-research", "--json"], {
        tty: true,
        home,
      });

      expect(result.code).toBe(0);
      const payload = JSON.parse(result.stdout);
      expect(payload.data.job.model).toBe("research");
      expect(payload.data.job.reasoning).toBe("standard");
      expect(payload.data.job.options.temporary).toBe(false);
      const requestBody = requestBodyFromExpression(expression);
      expect(requestBody.model).toBe("research");
      expect(requestBody.thinking_effort).toBe("standard");
      expect(requestBody).not.toHaveProperty("history_and_training_disabled");
    });
  });

  test("ask connects Deep Research with extended thinking", async () => {
    await withHome(async (home) => {
      await writeSessionToken(home);
      let expression = "";
      installFakeCdp(conversationStream("OK"), (script) => {
        expression = script;
      });

      const result = await run(
        ["ask", "hello", "--model", "deep-research", "--reasoning", "extended", "--json"],
        { tty: true, home },
      );

      expect(result.code).toBe(0);
      const payload = JSON.parse(result.stdout);
      expect(payload.data.job.model).toBe("research");
      expect(payload.data.job.reasoning).toBe("extended");
      expect(payload.data.job.options.temporary).toBe(false);
      const requestBody = requestBodyFromExpression(expression);
      expect(requestBody.model).toBe("research");
      expect(requestBody.thinking_effort).toBe("extended");
      expect(requestBody).not.toHaveProperty("history_and_training_disabled");
    });
  });

  test("rejects temporary chats for Deep Research", async () => {
    await withHome(async (home) => {
      const result = await run(
        ["ask", "hello", "--model", "deep-research", "--temporary", "--json"],
        { tty: true, home },
      );

      expect(result.code).toBe(2);
      const payload = JSON.parse(result.stderr);
      expect(payload.error.code).toBe("INVALID_ARGS");
      expect(payload.error.message).toContain("Deep Research does not support temporary chats");
    });
  });

  test("rejects explicit reasoning for GPT-4.5", async () => {
    await withHome(async (home) => {
      const result = await run(
        ["ask", "hello", "--model", "gpt-4-5", "--reasoning", "extended", "--json"],
        { tty: true, home },
      );

      expect(result.code).toBe(2);
      const payload = JSON.parse(result.stderr);
      expect(payload.error.code).toBe("INVALID_ARGS");
      expect(payload.error.message).toContain("does not support --reasoning");
    });
  });

  test("rejects model auto", async () => {
    await withHome(async (home) => {
      const result = await run(["ask", "hello", "--model", "auto", "--json"], { tty: true, home });

      expect(result.code).toBe(2);
      const payload = JSON.parse(result.stderr);
      expect(payload.error.code).toBe("INVALID_ARGS");
      expect(payload.error.message).toContain("Invalid --model");
    });
  });

  test("rejects repeated single-value request flags instead of falling back to defaults", async () => {
    await withHome(async (home) => {
      const result = await run(
        ["ask", "hello", "--model", "gpt-a", "--model", "gpt-b", "--json"],
        { tty: true, home },
      );

      expect(result.code).toBe(2);
      const payload = JSON.parse(result.stderr);
      expect(payload.error.code).toBe("INVALID_ARGS");
      expect(payload.error.message).toContain("Repeated --model");
    });
  });

  test("ask can opt into saved and continued conversations", async () => {
    await withHome(async (home) => {
      await writeSessionToken(home);
      let expression = "";
      installFakeCdp(conversationStream("OK"), (script) => {
        expression = script;
      });

      const result = await run(
        [
          "ask",
          "continue this",
          "--json",
          "--save",
          "--conversation",
          "conv_123",
          "--parent",
          "msg_456",
          "--reasoning",
          "extended",
        ],
        { tty: true, home },
      );

      expect(result.code).toBe(0);
      const requestBody = requestBodyFromExpression(expression);
      expect(requestBody.conversation_id).toBe("conv_123");
      expect(requestBody.parent_message_id).toBe("msg_456");
      expect(requestBody.model).toBe("gpt-5-5-pro");
      expect(requestBody).not.toHaveProperty("history_and_training_disabled");
      expect(requestBody.thinking_effort).toBe("extended");
    });
  });

  test("continuing a conversation requires conversation and parent ids together", async () => {
    await withHome(async (home) => {
      const result = await run(["job", "create", "hello", "--conversation", "conv_123", "--json"], {
        tty: true,
        home,
      });

      expect(result.code).toBe(2);
      const payload = JSON.parse(result.stderr);
      expect(payload.error.code).toBe("INVALID_ARGS");
      expect(payload.error.message).toContain("both ids");
    });
  });

  test("job create cannot wait with no-start", async () => {
    await withHome(async (home) => {
      const result = await run(["job", "create", "hello", "--wait", "--no-start", "--json"], {
        tty: true,
        home,
      });

      expect(result.code).toBe(2);
      const payload = JSON.parse(result.stderr);
      expect(payload.error.code).toBe("INVALID_ARGS");
      expect(payload.error.message).toContain("without starting the daemon");
    });
  });

  test("job create rejects wait options without wait mode", async () => {
    await withHome(async (home) => {
      const result = await run(["job", "create", "hello", "--soft-timeout", "1000", "--json"], {
        tty: true,
        home,
      });

      expect(result.code).toBe(2);
      const payload = JSON.parse(result.stderr);
      expect(payload.error.code).toBe("INVALID_ARGS");
      expect(payload.error.message).toContain("Wait options require --wait");
    });
  });

  test("job wait rejects mixed timeout modes before contacting the daemon", async () => {
    await withHome(async (home) => {
      const result = await run(
        ["job", "wait", "job_fake", "--wait-timeout", "1", "--soft-timeout", "1", "--json"],
        { tty: true, home },
      );

      expect(result.code).toBe(2);
      const payload = JSON.parse(result.stderr);
      expect(payload.error.code).toBe("INVALID_ARGS");
      expect(payload.error.message).toContain("Choose one wait timeout mode");
    });
  });

  test("job create --wait reports terminal failed jobs as top-level JSON errors", async () => {
    await withHome(async (home) => {
      try {
        const result = await run(["job", "create", "hello", "--wait", "--json", "--wait-timeout", "5000"], {
          tty: true,
          home,
        });

        expect(result.code).toBe(EXIT.auth);
        expect(result.stdout).toBe("");
        const payload = JSON.parse(result.stderr);
        expect(payload.ok).toBe(false);
        expect(payload.error.code).toBe("SESSION_TOKEN_MISSING");
        expect(payload.error.details.jobStatus).toBe("failed");
        expect(payload.error.details.job.id).toMatch(/^job_/);
        expect(payload.error.details.waited.wait.status).toBe("failed");
      } finally {
        await run(["daemon", "stop", "--json"], { tty: true, home });
      }
    });
  });

  test("job wait reports terminal failed jobs as top-level JSON errors", async () => {
    await withHome(async (home) => {
      const store = await JobStore.open(join(home, "jobs.sqlite"));
      let jobId = "";
      try {
        const created = store.create({
          prompt: "hello",
          model: "gpt-5-5-pro",
          reasoning: "standard",
          options: {},
        });
        jobId = created.id;
        store.markRunning(jobId);
        store.markFailed(
          jobId,
          new ProError("CHATGPT_PROBE_FAILED", "Could not determine ChatGPT login state from the CDP page (HTTP 431).", {
            exitCode: EXIT.auth,
            suggestions: ["HTTP 431 indicates oversize request headers."],
            details: { status: 431 },
          }),
        );
      } finally {
        store.close();
      }

      try {
        const result = await run(["job", "wait", jobId, "--json"], { tty: true, home });

        expect(result.code).toBe(EXIT.auth);
        expect(result.stdout).toBe("");
        const payload = JSON.parse(result.stderr);
        expect(payload.ok).toBe(false);
        expect(payload.error.code).toBe("CHATGPT_PROBE_FAILED");
        expect(payload.error.message).toContain("HTTP 431");
        expect(payload.error.suggestions).toContain("HTTP 431 indicates oversize request headers.");
        expect(payload.error.details.status).toBe(431);
        expect(payload.error.details.jobStatus).toBe("failed");
        expect(payload.error.details.job.id).toBe(jobId);
        expect(payload.error.details.waited.wait.status).toBe("failed");
      } finally {
        await run(["daemon", "stop", "--json"], { tty: true, home });
      }
    });
  });

  test("job wait soft timeout remains top-level success for non-terminal jobs", async () => {
    await withHome(async (home) => {
      const store = await JobStore.open(join(home, "jobs.sqlite"));
      let jobId = "";
      try {
        const created = store.create({
          prompt: "hello",
          model: "gpt-5-5-pro",
          reasoning: "standard",
          options: {},
        });
        jobId = created.id;
        store.markRunning(jobId);
      } finally {
        store.close();
      }

      try {
        const result = await run(["job", "wait", jobId, "--soft-timeout", "1", "--poll-ms", "25", "--json"], {
          tty: true,
          home,
        });

        expect(result.code).toBe(0);
        const payload = JSON.parse(result.stdout);
        expect(payload.ok).toBe(true);
        expect(payload.data.job.status).toBe("running");
        expect(payload.data.wait.timedOut).toBe(true);
      } finally {
        await run(["daemon", "stop", "--json"], { tty: true, home });
      }
    });
  });

  test("job result includes relay guidance for agents", async () => {
    await withHome(async (home) => {
      const store = await JobStore.open(join(home, "jobs.sqlite"));
      let jobId = "";
      try {
        const created = store.create({
          prompt: "hello",
          model: "gpt-5-5-pro",
          reasoning: "standard",
          options: {},
        });
        jobId = created.id;
        store.markRunning(jobId);
        store.markSucceeded(jobId, "Full Pro answer.");
      } finally {
        store.close();
      }

      const result = await run(["job", "result", jobId, "--json"], { tty: true, home });

      expect(result.code).toBe(0);
      const payload = JSON.parse(result.stdout);
      expect(payload.data.result).toBe("Full Pro answer.");
      expect(payload.data.agentInstruction).toContain("prefer relaying it in full");
      expect(payload.data.resultStats.chars).toBe("Full Pro answer.".length);
    });
  });

  test("ask prints plain text result for TTY users", async () => {
    await withHome(async (home) => {
      await mkdir(join(home, "tokens"), { recursive: true });
      await writeFile(
        join(home, "tokens", "chatgpt-session.json"),
        JSON.stringify({
          version: 1,
          generatedAt: new Date().toISOString(),
          source: "pro-cli-cdp-page",
          accessToken: fakeJwt(),
          accountId: "acct_test",
          expiresAt: new Date(Date.now() + 60 * 60 * 1000).toISOString(),
        }),
      );
      installFakeCdp(conversationStream("OK"));

      const result = await run(["ask", "hello"], { tty: true, home });

      expect(result.code).toBe(0);
      expect(result.stdout).toBe("OK\n");
    });
  });

  test("ask recovers once from ChatGPT cookie bloat", async () => {
    await withHome(async (home) => {
      await writeSessionToken(home);
      let deleted = 0;
      let navigated = false;
      installCookieBloatRecoveryCdp(() => {
        deleted += 1;
      }, () => {
        navigated = true;
      });

      const result = await run(["ask", "hello", "--json", "--timeout", "10"], { tty: true, home });

      expect(result.code).toBe(0);
      expect(JSON.parse(result.stdout).data.result).toBe("OK");
      expect(deleted).toBe(1);
      expect(navigated).toBe(true);
    });
  });

  test("ask recovers once from HTTP 431 auth-probe cookie bloat", async () => {
    await withHome(async (home) => {
      await writeSessionToken(home);
      let deleted = 0;
      let navigated = false;
      installCookieBloatRecoveryCdp(
        () => {
          deleted += 1;
        },
        () => {
          navigated = true;
        },
        {
          ok: false,
          status: 431,
          code: "CHATGPT_PROBE_FAILED",
          body: "ChatGPT auth session probe returned HTTP 431.",
        },
      );

      const result = await run(["ask", "hello", "--json", "--timeout", "10"], { tty: true, home });

      expect(result.code).toBe(0);
      expect(JSON.parse(result.stdout).data.result).toBe("OK");
      expect(deleted).toBe(1);
      expect(navigated).toBe(true);
    });
  });

  test("rejects unsupported request flags instead of silently ignoring them", async () => {
    await withHome(async (home) => {
      const result = await run(["job", "create", "hello", "--temperature", "0.2", "--json"], {
        tty: true,
        home,
      });

      expect(result.code).toBe(2);
      const payload = JSON.parse(result.stderr);
      expect(payload.error.code).toBe("INVALID_ARGS");
      expect(payload.error.message).toContain("Unsupported --temperature");
    });
  });

  test("doctor refuses ready when the CDP ChatGPT page is logged out", async () => {
    await withHome(async (home) => {
      await writeSessionToken(home);
      installFakeCdpValue({
        status: 200,
        hasAccessToken: false,
        origin: "https://chatgpt.com",
      });

      const result = await run(["doctor", "--json", "--timeout", "1000"], { tty: true, home });

      expect(result.code).toBe(0);
      const payload = JSON.parse(result.stdout);
      expect(payload.data.auth.tokenStatus).toBe("present");
      expect(payload.data.browserSession.status).toBe("logged_out");
      expect(payload.data.ready).toBe(false);
      expect(payload.data.transport.status).toBe("auth_required");
      expect(payload.data.next.command).toContain("pro-cli auth capture");
    });
  });

  test("doctor reports ready only when stored auth and live browser session are both present", async () => {
    await withHome(async (home) => {
      await writeSessionToken(home);
      installFakeCdpValue({
        status: 200,
        hasAccessToken: true,
        origin: "https://chatgpt.com",
      });

      const result = await run(["doctor", "--json", "--cdp", "http://127.0.0.1:9555"], {
        tty: true,
        home,
      });

      expect(result.code).toBe(0);
      const payload = JSON.parse(result.stdout);
      expect(payload.data.browserSession.status).toBe("present");
      expect(payload.data.browserSession.cdpBase).toBe("http://127.0.0.1:9555");
      expect(payload.data.ready).toBe(true);
      expect(payload.data.transport.status).toBe("configured");
      expect(payload.data.next.command).toContain("--cdp http://127.0.0.1:9555");
      expect(payload.data.next.command).not.toContain("Reply with OK");
      expect(JSON.stringify(payload.data.next).toLowerCase()).not.toContain("smoke");
      expect(result.stdout).not.toContain("header.");
    });
  });

  test("doctor reports probe_failed (NOT logged_out) when the probe returns a non-200/401 status", async () => {
    // Regression guard: before the probe_failed split, HTTP 431 was
    // collapsed into logged_out, sending users down the wrong remediation.
    await withHome(async (home) => {
      await writeSessionToken(home);
      installFakeCdpValue({
        status: 431,
        hasAccessToken: false,
        origin: "https://chatgpt.com",
      });

      const result = await run(["doctor", "--json", "--timeout", "1000"], { tty: true, home });
      expect(result.code).toBe(0);
      const payload = JSON.parse(result.stdout);
      expect(payload.data.browserSession.status).toBe("probe_failed");
      expect(payload.data.browserSession.httpStatus).toBe(431);
      // The 431-specific suggestion mentions cookies (not "sign in again").
      expect(
        payload.data.browserSession.suggestions.some((s: string) =>
          s.toLowerCase().includes("cookie"),
        ),
      ).toBe(true);
      // doctor.next still points at auth capture as the recovery action.
      expect(payload.data.next.command).toContain("pro-cli auth capture");
    });
  });

  test("doctor includes portCollision and legacyArtifacts diagnostic fields", async () => {
    // These fields were added to surface the failure modes that caused
    // today's incident. They must appear in every doctor report so agents
    // can act on them without guessing.
    await withHome(async (home) => {
      await writeSessionToken(home);
      installFakeCdpValue({ status: 200, hasAccessToken: true, origin: "https://chatgpt.com" });
      const result = await run(["doctor", "--json", "--cdp", "http://127.0.0.1:65432"], {
        tty: true,
        home,
      });
      expect(result.code).toBe(0);
      const payload = JSON.parse(result.stdout);

      expect(payload.data.portCollision).toBeDefined();
      expect(typeof payload.data.portCollision.inUse).toBe("boolean");
      expect(typeof payload.data.portCollision.conflict).toBe("boolean");
      expect(Array.isArray(payload.data.portCollision.listeners)).toBe(true);
      expect(payload.data.portCollision.port).toBe("65432");

      expect(payload.data.legacyArtifacts).toBeDefined();
      expect(typeof payload.data.legacyArtifacts.legacyHomeExists).toBe("boolean");
      expect(typeof payload.data.legacyArtifacts.legacyProfileExists).toBe("boolean");
      expect(typeof payload.data.legacyArtifacts.legacyHome).toBe("string");
      expect(typeof payload.data.legacyArtifacts.legacyProfileDir).toBe("string");
    });
  });

  test("auth command output includes portCollision so agents can detect dual-Chrome races", async () => {
    await withHome(async (home) => {
      const result = await run(["auth", "command", "--port", "9444", "--json"], {
        tty: true,
        home,
      });
      expect(result.code).toBe(0);
      const payload = JSON.parse(result.stdout);
      expect(payload.data.portCollision).toBeDefined();
      expect(payload.data.portCollision.port).toBe("9444");
      expect(payload.data.profileDir).toContain("chrome-profile");
    });
  });

  test("auth reset --no-launch --no-backup deletes the chrome profile dir", async () => {
    // Regression guard: --no-backup must actually delete the profile, not
    // silently fall through to backup mode (would mask "delete" intent).
    await withHome(async (home) => {
      await mkdir(join(home, "chrome-profile", "Default"), { recursive: true });
      await writeFile(join(home, "chrome-profile", "Default", "Cookies"), "dummy");

      const result = await run(["auth", "reset", "--no-launch", "--no-backup", "--json"], {
        tty: true,
        home,
      });
      expect(result.code).toBe(0);
      const payload = JSON.parse(result.stdout);
      expect(payload.data.removed.mode).toBe("delete");
      expect(payload.data.removed.from).toBe(join(home, "chrome-profile"));
      expect(payload.data.launched).toBeNull();
      // The profile dir is gone, no backup created.
      const remaining = await listEntries(home);
      expect(remaining).not.toContain("chrome-profile");
      expect(remaining.filter((e) => e.startsWith("chrome-profile.backup-"))).toHaveLength(0);
    });
  });

  test("auth reset (default) backs up the profile dir to chrome-profile.backup-<ts>", async () => {
    await withHome(async (home) => {
      await mkdir(join(home, "chrome-profile", "Default"), { recursive: true });
      await writeFile(join(home, "chrome-profile", "Default", "marker"), "preserve me");

      const result = await run(["auth", "reset", "--no-launch", "--json"], { tty: true, home });
      expect(result.code).toBe(0);
      const payload = JSON.parse(result.stdout);
      expect(payload.data.removed.mode).toBe("backup");
      expect(payload.data.removed.to).toMatch(/chrome-profile\.backup-\d{8}-\d{6}/);
      // Backup contents preserved.
      const preserved = await readFile(
        join(payload.data.removed.to, "Default", "marker"),
        "utf8",
      );
      expect(preserved).toBe("preserve me");
    });
  });

  test("auth reset reports mode=missing and does not crash when no profile exists", async () => {
    await withHome(async (home) => {
      const result = await run(["auth", "reset", "--no-launch", "--json"], { tty: true, home });
      expect(result.code).toBe(0);
      const payload = JSON.parse(result.stdout);
      expect(payload.data.removed.mode).toBe("missing");
      expect(payload.data.killedPids).toEqual([]);
    });
  });

  test("auth reset --keep-backups N prunes older backups beyond N", async () => {
    await withHome(async (home) => {
      // Pre-seed three older backups, then trigger a reset that creates a
      // fourth. With --keep-backups 2, the two oldest must be pruned.
      await mkdir(join(home, "chrome-profile", "Default"), { recursive: true });
      const old = [
        "chrome-profile.backup-20000101-000000",
        "chrome-profile.backup-20000102-000000",
        "chrome-profile.backup-20000103-000000",
      ];
      for (const name of old) {
        await mkdir(join(home, name), { recursive: true });
        await writeFile(join(home, name, "marker"), name);
      }

      const result = await run(
        ["auth", "reset", "--no-launch", "--keep-backups", "2", "--json"],
        { tty: true, home },
      );
      expect(result.code).toBe(0);
      const payload = JSON.parse(result.stdout);
      const prunedNames = (payload.data.prunedBackups as string[]).map((path) => basename(path)).sort();
      expect(prunedNames).toEqual([
        "chrome-profile.backup-20000101-000000",
        "chrome-profile.backup-20000102-000000",
      ]);
      const remaining = (await listEntries(home)).filter((e) => e.startsWith("chrome-profile.backup-"));
      expect(remaining.sort()).toEqual([
        basename(payload.data.removed.to as string),
        "chrome-profile.backup-20000103-000000",
      ].sort());
    });
  });

  test("auth hide moves the chatgpt window off-screen via Browser.setWindowBounds", async () => {
    await withHome(async (home) => {
      const observed = installAuthHideShowCdp();

      const result = await run(["auth", "hide", "--json"], { tty: true, home });
      expect(result.code).toBe(0);
      const payload = JSON.parse(result.stdout);
      expect(payload.data.targetId).toBe("CHATGPT_TAB");
      expect(payload.data.windowId).toBe(771965498);
      expect(payload.data.after.left).toBe(-32000);
      expect(payload.data.after.top).toBe(-32000);
      expect(payload.data.note.toLowerCase()).toContain("off-screen");
      // CDP methods called in order: get window, then set bounds.
      expect(observed.methods).toEqual([
        "Browser.getWindowForTarget",
        "Browser.setWindowBounds",
      ]);
      // The setWindowBounds params include the off-screen coords.
      const setParams = observed.params[1] as { bounds: { left: number; top: number } };
      expect(setParams.bounds.left).toBe(-32000);
      expect(setParams.bounds.top).toBe(-32000);
    });
  });

  test("auth show restores the window to a sensible centered position", async () => {
    await withHome(async (home) => {
      const observed = installAuthHideShowCdp();

      const result = await run(["auth", "show", "--json"], { tty: true, home });
      expect(result.code).toBe(0);
      const payload = JSON.parse(result.stdout);
      expect(payload.data.after.left).toBe(100);
      expect(payload.data.after.top).toBe(100);
      expect(payload.data.after.width).toBeGreaterThan(0);
      expect(payload.data.after.height).toBeGreaterThan(0);
      expect(payload.data.note.toLowerCase()).toContain("restored");

      const setParams = observed.params[1] as { bounds: { left: number; top: number } };
      expect(setParams.bounds.left).toBe(100);
      expect(setParams.bounds.top).toBe(100);
    });
  });

  test("auth hide fails clearly when no chatgpt.com tab is available", async () => {
    await withHome(async (home) => {
      // CDP is reachable but there is no chatgpt.com tab to operate on.
      globalThis.fetch = (async (url: string | URL | Request) => {
        const target = String(url);
        if (target.endsWith("/json")) {
          return Response.json([
            { id: "OTHER", type: "page", url: "https://example.com/", webSocketDebuggerUrl: "ws://other" },
          ]);
        }
        if (target.endsWith("/json/version")) {
          return Response.json({ webSocketDebuggerUrl: "ws://browser" });
        }
        return new Response("nope", { status: 500 });
      }) as unknown as typeof fetch;

      const result = await run(["auth", "hide", "--json"], { tty: true, home });
      expect(result.code).toBe(3); // EXIT.auth
      const payload = JSON.parse(result.stderr);
      expect(payload.error.code).toBe("CHATGPT_PAGE_MISSING");
    });
  });

  test("ask response includes the no-probe agent guidance (regression guard)", async () => {
    // A future refactor that drops the no-test-probe sentence would let
    // agents resume running smoke-tests against Pro and burning quota.
    await withHome(async (home) => {
      await writeSessionToken(home);
      installFakeCdp(conversationStream("Real answer."));
      const result = await run(["ask", "explain X", "--json"], { tty: true, home });
      expect(result.code).toBe(0);
      const payload = JSON.parse(result.stdout);
      const instruction = String(payload.data.agentInstruction).toLowerCase();
      expect(instruction).toContain("probe");
      expect(instruction).toContain("pro quota");
    });
  });

  test("unknown auth subcommand suggests the full subcommand list (status/command/capture/reset/hide/show)", async () => {
    await withHome(async (home) => {
      const result = await run(["auth", "bogus", "--json"], { tty: true, home });
      expect(result.code).toBe(2);
      const payload = JSON.parse(result.stderr);
      expect(payload.error.code).toBe("INVALID_ARGS");
      const suggestion = (payload.error.suggestions[0] as string).toLowerCase();
      expect(suggestion).toContain("status");
      expect(suggestion).toContain("capture");
      expect(suggestion).toContain("reset");
      expect(suggestion).toContain("hide");
      expect(suggestion).toContain("show");
    });
  });
});

async function listEntries(dir: string): Promise<string[]> {
  const { readdir } = await import("node:fs/promises");
  return readdir(dir).catch(() => [] as string[]);
}

function installAuthHideShowCdp(): { methods: string[]; params: Array<unknown> } {
  const observed = { methods: [] as string[], params: [] as unknown[] };
  globalThis.fetch = (async (url: string | URL | Request) => {
    const target = String(url);
    if (target.endsWith("/json")) {
      return Response.json([
        { id: "CHATGPT_TAB", type: "page", url: "https://chatgpt.com/", webSocketDebuggerUrl: "ws://chatgpt-tab" },
      ]);
    }
    if (target.endsWith("/json/version")) {
      return Response.json({ webSocketDebuggerUrl: "ws://browser" });
    }
    return new Response("nope", { status: 500 });
  }) as unknown as typeof fetch;

  class FakeWebSocket extends EventTarget {
    constructor(_url: string) {
      super();
      queueMicrotask(() => this.dispatchEvent(new Event("open")));
    }
    send(raw: string): void {
      const message = JSON.parse(raw) as { id: number; method: string; params?: unknown };
      observed.methods.push(message.method);
      observed.params.push(message.params);
      let result: unknown = {};
      if (message.method === "Browser.getWindowForTarget") {
        result = { windowId: 771965498, bounds: { left: 22, top: 47, width: 1200, height: 1011, windowState: "normal" } };
      }
      const response = { id: message.id, result };
      queueMicrotask(() =>
        this.dispatchEvent(new MessageEvent("message", { data: JSON.stringify(response) })),
      );
    }
    close(): void {
      this.dispatchEvent(new Event("close"));
    }
  }
  globalThis.WebSocket = FakeWebSocket as unknown as typeof WebSocket;
  return observed;
}

async function writeSessionToken(home: string): Promise<void> {
  await mkdir(join(home, "tokens"), { recursive: true });
  await writeFile(
    join(home, "tokens", "chatgpt-session.json"),
    JSON.stringify({
      version: 1,
      generatedAt: new Date().toISOString(),
      source: "pro-cli-cdp-page",
      accessToken: fakeJwt(),
      accountId: "acct_test",
      expiresAt: new Date(Date.now() + 60 * 60 * 1000).toISOString(),
    }),
  );
}

function fakeJwt(): string {
  const payload = {
    exp: Math.floor(Date.now() / 1000) + 3600,
    "https://api.openai.com/auth": { chatgpt_account_id: "acct_test" },
  };
  return ["header", Buffer.from(JSON.stringify(payload)).toString("base64url"), "sig"].join(".");
}

function installFakeCdp(body: string, onExpression?: (expression: string) => void): void {
  installFakeCdpValue({ ok: true, status: 200, body }, onExpression);
}

function installFakeCdpValue(value: unknown, onExpression?: (expression: string) => void): void {
  globalThis.fetch = (async (url: string | URL | Request) => {
    const target = String(url);
    if (target.endsWith("/json")) {
      return Response.json([
        {
          type: "page",
          url: "https://chatgpt.com/",
          webSocketDebuggerUrl: "ws://fake-chatgpt-page",
        },
      ]);
    }
    if (target.endsWith("/json/version")) {
      return Response.json({ webSocketDebuggerUrl: "ws://fake-browser" });
    }
    return new Response("unexpected fetch", { status: 500 });
  }) as unknown as typeof fetch;

  class FakeWebSocket extends EventTarget {
    constructor(_url: string) {
      super();
      queueMicrotask(() => this.dispatchEvent(new Event("open")));
    }

    send(raw: string): void {
      const message = JSON.parse(raw) as { id: number; params?: { expression?: string } };
      onExpression?.(message.params?.expression ?? "");
      const response = {
        id: message.id,
        result: {
          result: { value },
        },
      };
      queueMicrotask(() =>
        this.dispatchEvent(new MessageEvent("message", { data: JSON.stringify(response) })),
      );
    }

    close(): void {
      this.dispatchEvent(new Event("close"));
    }
  }

  globalThis.WebSocket = FakeWebSocket as unknown as typeof WebSocket;
}

function installCookieBloatRecoveryCdp(
  onDelete: () => void,
  onNavigate: () => void,
  firstEvaluation: Record<string, unknown> = {
    ok: false,
    status: 0,
    code: "CHATGPT_PAGE_MISSING",
    body: "Expected https://chatgpt.com, got chrome-error://chromewebdata/",
  },
): void {
  globalThis.fetch = (async (url: string | URL | Request) => {
    const target = String(url);
    if (target.endsWith("/json")) {
      return Response.json([
        {
          type: "page",
          url: "https://chatgpt.com/",
          webSocketDebuggerUrl: "ws://fake-chatgpt-page",
        },
      ]);
    }
    if (target.endsWith("/json/version")) {
      return Response.json({ webSocketDebuggerUrl: "ws://fake-browser" });
    }
    return new Response("unexpected fetch", { status: 500 });
  }) as unknown as typeof fetch;

  let runtimeEvaluations = 0;
  class FakeWebSocket extends EventTarget {
    constructor(_url: string) {
      super();
      queueMicrotask(() => this.dispatchEvent(new Event("open")));
    }

    send(raw: string): void {
      const message = JSON.parse(raw) as { id: number; method: string };
      let response: Record<string, unknown> = { id: message.id, result: {} };
      if (message.method === "Runtime.evaluate") {
        runtimeEvaluations += 1;
        response = {
          id: message.id,
          result: {
            result: {
              value:
                runtimeEvaluations === 1
                  ? firstEvaluation
                  : { ok: true, status: 200, body: conversationStream("OK") },
            },
          },
        };
      }
      if (message.method === "Network.getCookies") {
        response = {
          id: message.id,
          result: {
            cookies: [
              { name: "__Secure-next-auth.session-token", value: "x", domain: "chatgpt.com", path: "/" },
              { name: "conv_key_abc", value: "x", domain: "chatgpt.com", path: "/" },
            ],
          },
        };
      }
      if (message.method === "Network.deleteCookies") onDelete();
      if (message.method === "Page.navigate") onNavigate();
      queueMicrotask(() =>
        this.dispatchEvent(new MessageEvent("message", { data: JSON.stringify(response) })),
      );
    }

    close(): void {
      this.dispatchEvent(new Event("close"));
    }
  }

  globalThis.WebSocket = FakeWebSocket as unknown as typeof WebSocket;
}

function requestBodyFromExpression(expression: string): Record<string, unknown> {
  const marker = '})("https://chatgpt.com/backend-api/f/conversation", ';
  const start = expression.lastIndexOf(marker);
  expect(start).toBeGreaterThanOrEqual(0);
  const bodyStart = start + marker.length;
  const bodyEnd = expression.lastIndexOf(', "acct_test")');
  expect(bodyEnd).toBeGreaterThan(bodyStart);
  return JSON.parse(expression.slice(bodyStart, bodyEnd)) as Record<string, unknown>;
}

function conversationStream(text: string): string {
  return [
    `data: {"message":{"author":{"role":"assistant"},"content":{"content_type":"text","parts":[${JSON.stringify(text)}]},"status":"finished_successfully"}}`,
    "data: [DONE]",
    "",
  ].join("\n\n");
}
`````

## File: tests/config.test.ts
`````typescript
import { chmod, mkdir, mkdtemp, readFile, readdir, rm, stat, writeFile } from "node:fs/promises";
import { join } from "node:path";
import { tmpdir } from "node:os";
import { describe, expect, test } from "bun:test";
import {
  ensurePrivateDir,
  expandPath,
  loadConfig,
  migrateLegacyDefaultHome,
  resolveHome,
  resolvePaths,
  saveConfig,
  writePrivateFile,
} from "../src/config";
import { ProError } from "../src/errors";

describe("expandPath", () => {
  test("returns absolute path verbatim", () => {
    expect(expandPath("/etc/hosts")).toBe("/etc/hosts");
  });

  test("resolves a relative path to absolute", () => {
    const result = expandPath("foo/bar");
    expect(result.endsWith("/foo/bar")).toBe(true);
    expect(result.startsWith("/")).toBe(true);
  });

  test("expands a leading ~ to the process homedir", () => {
    // homedir() is captured at process start; we don't try to mutate HOME
    // mid-test (Bun/Node cache it). Instead, verify the structure of the
    // expansion: ~ alone resolves and ~/foo joins onto that.
    const tildeOnly = expandPath("~");
    expect(tildeOnly.startsWith("/")).toBe(true);
    expect(expandPath("~/foo/bar")).toBe(`${tildeOnly}/foo/bar`);
  });

  test("does NOT expand ~user (only bare ~ or ~/)", () => {
    // Regression guard: ~someuser must NOT resolve to anyone's homedir.
    // It should be left alone (then resolved as a relative path).
    const result = expandPath("~someuser/foo");
    // Whatever the result is, it must not be the homedir form.
    expect(result.includes("~someuser")).toBe(true);
  });
});

describe("resolveHome", () => {
  test("uses PRO_CLI_HOME env var when set", () => {
    expect(resolveHome({ PRO_CLI_HOME: "/custom/home" })).toBe("/custom/home");
  });

  test("expands ~ in PRO_CLI_HOME", () => {
    const result = resolveHome({ PRO_CLI_HOME: "~/my-pro" });
    expect(result.endsWith("/my-pro")).toBe(true);
    expect(result.startsWith("/")).toBe(true);
    expect(result).not.toContain("~");
  });

  test("falls back to ~/.pro-cli when PRO_CLI_HOME is missing", () => {
    const result = resolveHome({});
    expect(result.endsWith("/.pro-cli")).toBe(true);
    expect(result.startsWith("/")).toBe(true);
  });

  test("treats empty PRO_CLI_HOME as missing", () => {
    const result = resolveHome({ PRO_CLI_HOME: "" });
    expect(result).toBe(resolveHome({}));
    expect(result).toContain(".pro-cli");
  });
});

describe("private filesystem helpers", () => {
  test("ensurePrivateDir creates directories with private mode", async () => {
    if (process.platform === "win32") return;
    const dir = await mkdtemp(join(tmpdir(), "pro-private-dir-"));
    try {
      const privateDir = join(dir, "private");
      await ensurePrivateDir(privateDir);
      expect((await stat(privateDir)).mode & 0o777).toBe(0o700);
    } finally {
      await rm(dir, { recursive: true, force: true });
    }
  });

  test("writePrivateFile creates parent directory 0700 and file 0600", async () => {
    if (process.platform === "win32") return;
    const dir = await mkdtemp(join(tmpdir(), "pro-private-file-"));
    try {
      const file = join(dir, "nested", "secret.json");
      await writePrivateFile(file, "{\"secret\":true}\n");
      expect((await stat(join(dir, "nested"))).mode & 0o777).toBe(0o700);
      expect((await stat(file)).mode & 0o777).toBe(0o600);
    } finally {
      await rm(dir, { recursive: true, force: true });
    }
  });
});

describe("resolvePaths", () => {
  test("computes all canonical paths under the resolved home", () => {
    const paths = resolvePaths({ PRO_CLI_HOME: "/x/.pro-cli" });
    expect(paths.home).toBe("/x/.pro-cli");
    expect(paths.configPath).toBe("/x/.pro-cli/config.json");
    expect(paths.cookieJsonPath).toBe("/x/.pro-cli/cookies/chatgpt.json");
    expect(paths.cookieJarPath).toBe("/x/.pro-cli/cookies/chatgpt.txt");
    expect(paths.sessionTokenPath).toBe("/x/.pro-cli/tokens/chatgpt-session.json");
    expect(paths.dbPath).toBe("/x/.pro-cli/jobs.sqlite");
  });

  test("env-level cookie path overrides config and default", () => {
    const paths = resolvePaths(
      { PRO_CLI_HOME: "/x/.pro-cli", CHATGPT_COOKIE_JSON: "/env/cookies.json" },
      { cookieJsonPath: "/config/cookies.json" },
    );
    expect(paths.cookieJsonPath).toBe("/env/cookies.json");
  });

  test("config-level cookie path overrides default when env is absent", () => {
    const paths = resolvePaths({ PRO_CLI_HOME: "/x/.pro-cli" }, { cookieJsonPath: "/config/cookies.json" });
    expect(paths.cookieJsonPath).toBe("/config/cookies.json");
  });

  test("expands env and config path overrides instead of storing raw shell-style paths", () => {
    const paths = resolvePaths(
      {
        PRO_CLI_HOME: "/x/.pro-cli",
        CHATGPT_COOKIE_JSON: "relative/cookies.json",
        CHATGPT_COOKIE_JAR: "~/cookies/chatgpt.txt",
      },
      { sessionTokenPath: "relative/tokens/session.json" },
    );

    expect(paths.cookieJsonPath).toMatch(/\/relative\/cookies\.json$/);
    expect(paths.cookieJsonPath.startsWith("/")).toBe(true);
    expect(paths.cookieJarPath).not.toContain("~");
    expect(paths.cookieJarPath.endsWith("/cookies/chatgpt.txt")).toBe(true);
    expect(paths.sessionTokenPath).toMatch(/\/relative\/tokens\/session\.json$/);
    expect(paths.sessionTokenPath.startsWith("/")).toBe(true);
  });

  test("ignores empty env override strings and falls back to config/default paths", () => {
    const paths = resolvePaths(
      {
        PRO_CLI_HOME: "/x/.pro-cli",
        CHATGPT_COOKIE_JSON: "",
        CHATGPT_COOKIE_JAR: "",
        CHATGPT_SESSION_TOKEN_JSON: "",
      },
      {
        cookieJsonPath: "/config/cookies.json",
        cookieJarPath: "/config/cookies.txt",
      },
    );

    expect(paths.cookieJsonPath).toBe("/config/cookies.json");
    expect(paths.cookieJarPath).toBe("/config/cookies.txt");
    expect(paths.sessionTokenPath).toBe("/x/.pro-cli/tokens/chatgpt-session.json");
  });

  test("env-level cookie jar path overrides config and default", () => {
    const paths = resolvePaths(
      { PRO_CLI_HOME: "/x/.pro-cli", CHATGPT_COOKIE_JAR: "/env/cookies.txt" },
      { cookieJarPath: "/config/cookies.txt" },
    );
    expect(paths.cookieJarPath).toBe("/env/cookies.txt");
  });

  test("env-level session token path overrides config and default", () => {
    const paths = resolvePaths(
      { PRO_CLI_HOME: "/x/.pro-cli", CHATGPT_SESSION_TOKEN_JSON: "/env/token.json" },
      { sessionTokenPath: "/config/token.json" },
    );
    expect(paths.sessionTokenPath).toBe("/env/token.json");
  });

  test("dbPath is always under home (never overridable via env or config)", () => {
    // Regression guard: if dbPath ever became overridable, the daemon's
    // discovery would split. Currently it must always live at <home>/jobs.sqlite.
    const paths = resolvePaths(
      {
        PRO_CLI_HOME: "/x/.pro-cli",
        // None of these should affect dbPath.
        CHATGPT_COOKIE_JSON: "/env/cookies.json",
        CHATGPT_COOKIE_JAR: "/env/cookies.txt",
        CHATGPT_SESSION_TOKEN_JSON: "/env/token.json",
      },
      {},
    );
    expect(paths.dbPath).toBe("/x/.pro-cli/jobs.sqlite");
  });
});

describe("loadConfig + saveConfig", () => {
  test("returns empty config when no file exists", async () => {
    const dir = await mkdtemp(join(tmpdir(), "pro-config-"));
    try {
      const config = await loadConfig({ PRO_CLI_HOME: dir });
      expect(config).toEqual({});
    } finally {
      await rm(dir, { recursive: true, force: true });
    }
  });

  test("round-trips: saved config is read back identically", async () => {
    const dir = await mkdtemp(join(tmpdir(), "pro-config-"));
    try {
      const original = {
        defaultModel: "gpt-5-5-pro",
        defaultReasoning: "extended",
        cookieJsonPath: "/some/path/cookies.json",
      };
      await saveConfig({ PRO_CLI_HOME: dir }, original);
      const loaded = await loadConfig({ PRO_CLI_HOME: dir });
      expect(loaded).toEqual(original);
    } finally {
      await rm(dir, { recursive: true, force: true });
    }
  });

  test("rejects malformed JSON (does not silently swallow)", async () => {
    // A corrupted config file must NOT silently load as empty — the user
    // would lose all customization without realizing.
    const dir = await mkdtemp(join(tmpdir(), "pro-config-"));
    try {
      await mkdir(dir, { recursive: true });
      await writeFile(join(dir, "config.json"), "{ not json");
      await expect(loadConfig({ PRO_CLI_HOME: dir })).rejects.toThrow();
    } finally {
      await rm(dir, { recursive: true, force: true });
    }
  });

  test("surfaces unreadable config as a first-class local storage error", async () => {
    if (process.platform === "win32" || process.getuid?.() === 0) return;
    const dir = await mkdtemp(join(tmpdir(), "pro-config-"));
    const configPath = join(dir, "config.json");
    try {
      await writeFile(configPath, "{}");
      await chmod(configPath, 0);

      try {
        await loadConfig({ PRO_CLI_HOME: dir });
        throw new Error("Expected CONFIG_UNREADABLE.");
      } catch (error) {
        expect(error).toBeInstanceOf(ProError);
        const proError = error as ProError;
        expect(proError.code).toBe("CONFIG_UNREADABLE");
        expect(proError.message).toContain(configPath);
        expect(proError.suggestions.join("\n")).toContain("pro-cli doctor --json");
        expect(proError.suggestions.join("\n")).toContain("smoke-test");
        expect(proError.details?.configPath).toBe(configPath);
      }
    } finally {
      await chmod(configPath, 0o600).catch(() => undefined);
      await rm(dir, { recursive: true, force: true });
    }
  });

  test("saveConfig creates the home directory if it does not exist", async () => {
    const dir = await mkdtemp(join(tmpdir(), "pro-config-"));
    try {
      const home = join(dir, "deep", "nested", "home");
      await saveConfig({ PRO_CLI_HOME: home }, { defaultModel: "gpt-5-5-pro" });
      const loaded = await loadConfig({ PRO_CLI_HOME: home });
      expect(loaded.defaultModel).toBe("gpt-5-5-pro");
    } finally {
      await rm(dir, { recursive: true, force: true });
    }
  });
});

describe("migrateLegacyDefaultHome", () => {
  test("is a no-op when PRO_CLI_HOME is set (user opted out of migration)", async () => {
    const dir = await mkdtemp(join(tmpdir(), "pro-migrate-skip-"));
    try {
      // Even if .pro exists under our fake home, PRO_CLI_HOME=dir means
      // we should NOT touch it.
      await mkdir(join(dir, ".pro"), { recursive: true });
      await writeFile(join(dir, ".pro", "marker.txt"), "legacy");
      await migrateLegacyDefaultHome({ PRO_CLI_HOME: dir }, dir);
      // .pro stays put.
      const entries = await readdir(dir);
      expect(entries).toContain(".pro");
      expect(entries).not.toContain(".pro-cli");
    } finally {
      await rm(dir, { recursive: true, force: true });
    }
  });

  test("is a no-op when neither legacy ~/.pro nor ~/.pro-cli exists", async () => {
    const dir = await mkdtemp(join(tmpdir(), "pro-migrate-fresh-"));
    try {
      await migrateLegacyDefaultHome({}, dir);
      const entries = await readdir(dir).catch(() => [] as string[]);
      expect(entries).not.toContain(".pro");
      expect(entries).not.toContain(".pro-cli");
    } finally {
      await rm(dir, { recursive: true, force: true });
    }
  });

  test("does NOT clobber an existing ~/.pro-cli when ~/.pro is also present", async () => {
    // Critical: if the user already migrated AND an older pro-cli still
    // wrote to ~/.pro, do not overwrite the current home.
    const dir = await mkdtemp(join(tmpdir(), "pro-migrate-coexist-"));
    try {
      await mkdir(join(dir, ".pro-cli"), { recursive: true });
      await writeFile(join(dir, ".pro-cli", "marker.txt"), "current");
      await mkdir(join(dir, ".pro"), { recursive: true });
      await writeFile(join(dir, ".pro", "marker.txt"), "legacy");

      await migrateLegacyDefaultHome({}, dir);

      const current = await readFile(join(dir, ".pro-cli", "marker.txt"), "utf8");
      expect(current).toBe("current");
    } finally {
      await rm(dir, { recursive: true, force: true });
    }
  });

  test("renames ~/.pro to ~/.pro-cli when only ~/.pro exists", async () => {
    const dir = await mkdtemp(join(tmpdir(), "pro-migrate-rename-"));
    try {
      await mkdir(join(dir, ".pro"), { recursive: true });
      await writeFile(join(dir, ".pro", "config.json"), "{}");

      await migrateLegacyDefaultHome({}, dir);

      const entries = await readdir(dir);
      expect(entries).toContain(".pro-cli");
      expect(entries).not.toContain(".pro");
      // rewriteMigratedConfigPaths re-pretty-prints the config, so the
      // file content changes shape but the parsed config is preserved.
      const moved = await readFile(join(dir, ".pro-cli", "config.json"), "utf8");
      expect(JSON.parse(moved)).toEqual({});
    } finally {
      await rm(dir, { recursive: true, force: true });
    }
  });

  test("rewrites paths inside config.json that pointed at the legacy home", async () => {
    const dir = await mkdtemp(join(tmpdir(), "pro-migrate-rewrite-"));
    try {
      await mkdir(join(dir, ".pro"), { recursive: true });
      const legacyConfig = {
        cookieJsonPath: join(dir, ".pro", "cookies", "chatgpt.json"),
        cookieJarPath: join(dir, ".pro", "cookies", "chatgpt.txt"),
        sessionTokenPath: join(dir, ".pro", "tokens", "chatgpt-session.json"),
        defaultModel: "gpt-5-5-pro",
      };
      await writeFile(join(dir, ".pro", "config.json"), JSON.stringify(legacyConfig));

      await migrateLegacyDefaultHome({}, dir);

      const newConfig = JSON.parse(
        await readFile(join(dir, ".pro-cli", "config.json"), "utf8"),
      );
      expect(newConfig.cookieJsonPath).toBe(join(dir, ".pro-cli", "cookies", "chatgpt.json"));
      expect(newConfig.cookieJarPath).toBe(join(dir, ".pro-cli", "cookies", "chatgpt.txt"));
      expect(newConfig.sessionTokenPath).toBe(join(dir, ".pro-cli", "tokens", "chatgpt-session.json"));
      // Unrelated config keys preserved.
      expect(newConfig.defaultModel).toBe("gpt-5-5-pro");
    } finally {
      await rm(dir, { recursive: true, force: true });
    }
  });

  test("leaves config paths that were OUTSIDE the legacy home untouched", async () => {
    // If the user pointed pro-cli at a custom cookie path (e.g.
    // /elsewhere/cookies.json), the migration must not rewrite it.
    const dir = await mkdtemp(join(tmpdir(), "pro-migrate-foreign-"));
    try {
      await mkdir(join(dir, ".pro"), { recursive: true });
      const legacyConfig = {
        cookieJsonPath: "/elsewhere/cookies.json",
        sessionTokenPath: join(dir, ".pro", "tokens", "session.json"),
      };
      await writeFile(join(dir, ".pro", "config.json"), JSON.stringify(legacyConfig));

      await migrateLegacyDefaultHome({}, dir);

      const newConfig = JSON.parse(
        await readFile(join(dir, ".pro-cli", "config.json"), "utf8"),
      );
      expect(newConfig.cookieJsonPath).toBe("/elsewhere/cookies.json"); // untouched
      expect(newConfig.sessionTokenPath).toBe(join(dir, ".pro-cli", "tokens", "session.json"));
    } finally {
      await rm(dir, { recursive: true, force: true });
    }
  });

  test("rewrites a config path that exactly equals the legacy home (no trailing slash)", async () => {
    // Edge case: a config that stored just legacyHome itself, not a child.
    const dir = await mkdtemp(join(tmpdir(), "pro-migrate-exact-"));
    try {
      await mkdir(join(dir, ".pro"), { recursive: true });
      const legacyConfig = {
        cookieJsonPath: join(dir, ".pro"),
      };
      await writeFile(join(dir, ".pro", "config.json"), JSON.stringify(legacyConfig));

      await migrateLegacyDefaultHome({}, dir);

      const newConfig = JSON.parse(
        await readFile(join(dir, ".pro-cli", "config.json"), "utf8"),
      );
      expect(newConfig.cookieJsonPath).toBe(join(dir, ".pro-cli"));
    } finally {
      await rm(dir, { recursive: true, force: true });
    }
  });
});
`````

## File: tests/cookies.test.ts
`````typescript
import { mkdtemp, rm, writeFile } from "node:fs/promises";
import { join } from "node:path";
import { tmpdir } from "node:os";
import { describe, expect, test } from "bun:test";
import {
  chatGptOrigins,
  cookieSummary,
  isVolatileCookieName,
  loadCookieExport,
  sanitizeCookies,
  toCookieExport,
  toCookieHeader,
  toNetscapeCookieJar,
  type BrowserCookie,
} from "../src/cookies";

describe("cookie sanitization", () => {
  test("drops volatile per-conversation cookies", () => {
    const cookies = sanitizeCookies([
      cookie("__Secure-next-auth.session-token", "chatgpt.com", "session"),
      cookie("conv_key_abc", "chatgpt.com", "temporary"),
    ]);

    expect(cookies.map((stored) => stored.name)).toEqual(["__Secure-next-auth.session-token"]);
  });

  test("does not replay volatile cookies in a Cookie header", () => {
    const header = toCookieHeader([
      cookie("__Secure-next-auth.session-token", "chatgpt.com", "session"),
      cookie("conv_key_abc", "chatgpt.com", "temporary"),
    ]);

    expect(header).toBe("__Secure-next-auth.session-token=session");
  });
});

describe("isVolatileCookieName", () => {
  test("matches the conv_key_ prefix exactly", () => {
    expect(isVolatileCookieName("conv_key_abc")).toBe(true);
    expect(isVolatileCookieName("conv_key_")).toBe(true);
    expect(isVolatileCookieName("conv_key_xyz_123")).toBe(true);
  });

  test("does NOT match cookies that merely contain the substring", () => {
    expect(isVolatileCookieName("not_conv_key_x")).toBe(false);
    expect(isVolatileCookieName("xconv_key_y")).toBe(false);
  });

  test("does NOT match without the trailing underscore", () => {
    // If someone changes prefix to "conv_key" they would over-match.
    expect(isVolatileCookieName("conv_keyabc")).toBe(false);
    expect(isVolatileCookieName("conv_key")).toBe(false);
  });

  test("preserves cookies critical to login and CF challenge state", () => {
    // Regression guard: if someone broadens the volatile list to include any
    // of these, sessions break silently or CF challenges replay on every call.
    expect(isVolatileCookieName("__Secure-next-auth.session-token")).toBe(false);
    expect(isVolatileCookieName("__Secure-next-auth.session-token.0")).toBe(false);
    expect(isVolatileCookieName("__Host-next-auth.csrf-token")).toBe(false);
    expect(isVolatileCookieName("cf_clearance")).toBe(false);
    expect(isVolatileCookieName("__cf_bm")).toBe(false);
    expect(isVolatileCookieName("_cfuvid")).toBe(false);
    expect(isVolatileCookieName("oai-did")).toBe(false);
    expect(isVolatileCookieName("oai-sc")).toBe(false);
  });

  test("is case sensitive (matches lowercase prefix only)", () => {
    expect(isVolatileCookieName("CONV_KEY_abc")).toBe(false);
    expect(isVolatileCookieName("Conv_Key_abc")).toBe(false);
  });
});

describe("sanitizeCookies invariants", () => {
  test("dedupes by name|domain|path triple, last write wins", () => {
    const result = sanitizeCookies([
      cookie("a", "chatgpt.com", "v1"),
      cookie("a", "chatgpt.com", "v2"),
    ]);
    expect(result).toHaveLength(1);
    expect(result[0].value).toBe("v2");
  });

  test("keeps separate entries for the same name on different paths", () => {
    const result = sanitizeCookies([
      { name: "a", value: "root", domain: "chatgpt.com", path: "/", secure: true },
      { name: "a", value: "api", domain: "chatgpt.com", path: "/api", secure: true },
    ]);
    expect(result).toHaveLength(2);
    expect(result.map((c) => c.path).sort()).toEqual(["/", "/api"]);
  });

  test("keeps separate entries for the same name on different domains", () => {
    const result = sanitizeCookies([
      cookie("a", "chatgpt.com", "v1"),
      cookie("a", "openai.com", "v2"),
    ]);
    expect(result).toHaveLength(2);
    expect(result.map((c) => c.domain).sort()).toEqual(["chatgpt.com", "openai.com"]);
  });

  test("strips a leading dot from the cookie domain", () => {
    const result = sanitizeCookies([
      { name: "a", value: "v", domain: ".chatgpt.com", path: "/", secure: true },
    ]);
    expect(result[0].domain).toBe("chatgpt.com");
  });

  test("dedupes a leading-dot domain against a non-dot domain", () => {
    // After strip, both keys collapse; otherwise the same logical cookie
    // could be sent twice with different attributes.
    const result = sanitizeCookies([
      { name: "a", value: "first", domain: ".chatgpt.com", path: "/", secure: true },
      { name: "a", value: "second", domain: "chatgpt.com", path: "/", secure: true },
    ]);
    expect(result).toHaveLength(1);
    expect(result[0].value).toBe("second");
  });

  test("defaults missing domain to chatgpt.com", () => {
    const result = sanitizeCookies([
      { name: "a", value: "v", domain: "", path: "/", secure: true },
    ]);
    expect(result[0].domain).toBe("chatgpt.com");
  });

  test("defaults missing path to /", () => {
    const result = sanitizeCookies([
      { name: "a", value: "v", domain: "chatgpt.com", path: "", secure: true },
    ]);
    expect(result[0].path).toBe("/");
  });

  test("filters cookies with empty name", () => {
    const result = sanitizeCookies([
      cookie("", "chatgpt.com", "v"),
      cookie("a", "chatgpt.com", "v"),
    ]);
    expect(result.map((c) => c.name)).toEqual(["a"]);
  });

  test("filters cookies with undefined value", () => {
    const result = sanitizeCookies([
      { name: "a", value: undefined as unknown as string, domain: "chatgpt.com", path: "/", secure: true },
      cookie("b", "chatgpt.com", "v"),
    ]);
    expect(result.map((c) => c.name)).toEqual(["b"]);
  });

  test("preserves an explicit empty-string value (auth flows rely on this)", () => {
    // value === undefined is filtered; value === "" must survive.
    const result = sanitizeCookies([{ name: "a", value: "", domain: "chatgpt.com", path: "/", secure: true }]);
    expect(result).toHaveLength(1);
    expect(result[0].value).toBe("");
  });

  test("returns cookies sorted alphabetically by name", () => {
    const result = sanitizeCookies([
      cookie("zebra", "chatgpt.com", "z"),
      cookie("alpha", "chatgpt.com", "a"),
      cookie("mango", "chatgpt.com", "m"),
    ]);
    expect(result.map((c) => c.name)).toEqual(["alpha", "mango", "zebra"]);
  });

  test("preserves all optional fields when present", () => {
    const expires = Math.floor(Date.now() / 1000) + 3600;
    const result = sanitizeCookies([
      {
        name: "a",
        value: "v",
        domain: "chatgpt.com",
        path: "/",
        secure: true,
        httpOnly: true,
        sameSite: "Lax",
        expires,
      },
    ]);
    expect(result[0]).toEqual({
      name: "a",
      value: "v",
      domain: "chatgpt.com",
      path: "/",
      secure: true,
      httpOnly: true,
      sameSite: "Lax",
      expires,
    });
  });

  test("preserves explicit secure=false", () => {
    // Important: do not silently flip secure to true (would break some local flows).
    const result = sanitizeCookies([
      { name: "a", value: "v", domain: "chatgpt.com", path: "/", secure: false },
    ]);
    expect(result[0].secure).toBe(false);
  });

  test("omits absent optional fields rather than emitting undefined", () => {
    // Stored JSON should not have explicit `undefined` keys leaking through.
    const result = sanitizeCookies([
      { name: "a", value: "v", domain: "chatgpt.com", path: "/" } as BrowserCookie,
    ]);
    const keys = Object.keys(result[0]).sort();
    expect(keys).toEqual(["domain", "name", "path", "value"]);
  });
});

describe("toCookieHeader", () => {
  test("joins multiple cookies with semicolon-space", () => {
    const header = toCookieHeader([
      cookie("a", "chatgpt.com", "v1"),
      cookie("b", "chatgpt.com", "v2"),
    ]);
    expect(header).toBe("a=v1; b=v2");
  });

  test("filters expired cookies (regression: stale tokens must NOT replay)", () => {
    const past = Math.floor(Date.now() / 1000) - 60;
    const future = Math.floor(Date.now() / 1000) + 3600;
    const header = toCookieHeader([
      { name: "old", value: "v", domain: "chatgpt.com", path: "/", secure: true, expires: past },
      { name: "fresh", value: "v", domain: "chatgpt.com", path: "/", secure: true, expires: future },
    ]);
    expect(header).toBe("fresh=v");
  });

  test("keeps session cookies that have no expires", () => {
    const header = toCookieHeader([cookie("session", "chatgpt.com", "v")]);
    expect(header).toBe("session=v");
  });

  test("returns empty string when given no cookies", () => {
    expect(toCookieHeader([])).toBe("");
  });

  test("returns empty string when every cookie is expired", () => {
    const past = Math.floor(Date.now() / 1000) - 60;
    const header = toCookieHeader([
      { name: "a", value: "v", domain: "chatgpt.com", path: "/", expires: past },
    ]);
    expect(header).toBe("");
  });

  test("never replays volatile cookies even when not expired", () => {
    const header = toCookieHeader([
      cookie("__Secure-next-auth.session-token", "chatgpt.com", "session"),
      cookie("conv_key_xyz", "chatgpt.com", "vol"),
    ]);
    expect(header).not.toContain("conv_key_");
    expect(header).toContain("session-token=session");
  });
});

describe("cookieSummary", () => {
  test("counts cookies (including duplicates), unique domains and names", () => {
    const summary = cookieSummary([
      cookie("a", "chatgpt.com", "v"),
      cookie("a", "chatgpt.com", "v"),
      cookie("b", "openai.com", "v"),
    ]);
    expect(summary.count).toBe(3);
    expect(summary.domains).toEqual(["chatgpt.com", "openai.com"]);
    expect(summary.names).toEqual(["a", "b"]);
  });

  test("strips leading dots from the domain set", () => {
    const summary = cookieSummary([
      cookie("a", ".chatgpt.com", "v"),
      cookie("b", "chatgpt.com", "v"),
    ]);
    expect(summary.domains).toEqual(["chatgpt.com"]);
  });

  test("counts expired cookies by current wall-clock", () => {
    const past = Math.floor(Date.now() / 1000) - 60;
    const future = Math.floor(Date.now() / 1000) + 3600;
    const summary = cookieSummary([
      { name: "old", value: "v", domain: "chatgpt.com", path: "/", expires: past },
      { name: "fresh", value: "v", domain: "chatgpt.com", path: "/", expires: future },
      { name: "session", value: "v", domain: "chatgpt.com", path: "/" },
    ]);
    expect(summary.expired).toBe(1);
    expect(summary.count).toBe(3);
  });

  test("returns empty arrays and zero counts for an empty input", () => {
    const summary = cookieSummary([]);
    expect(summary.count).toBe(0);
    expect(summary.expired).toBe(0);
    expect(summary.domains).toEqual([]);
    expect(summary.names).toEqual([]);
  });
});

describe("toCookieExport", () => {
  test("wraps cookies in v1 envelope and runs sanitization", () => {
    const exportObj = toCookieExport([
      cookie("a", "chatgpt.com", "v"),
      cookie("conv_key_x", "chatgpt.com", "drop"),
    ]);
    expect(exportObj.version).toBe(1);
    expect(exportObj.source).toBe("pro-cli-cdp");
    expect(exportObj.cookies.map((c) => c.name)).toEqual(["a"]);
    expect(exportObj.targetUrl).toBe("https://chatgpt.com/");
    expect(exportObj.origins).toContain("https://chatgpt.com/");
    expect(exportObj.origins).toContain("https://auth.openai.com/");
    expect(typeof exportObj.generatedAt).toBe("string");
    expect(() => new Date(exportObj.generatedAt).toISOString()).not.toThrow();
  });
});

describe("toNetscapeCookieJar", () => {
  test("emits Netscape format with TAB-separated fields", () => {
    const jar = toNetscapeCookieJar([
      { name: "a", value: "v", domain: ".chatgpt.com", path: "/", secure: true, expires: 1234567890 },
    ]);
    const entryLines = jar.split("\n").filter((l) => l && !l.startsWith("#"));
    expect(entryLines).toHaveLength(1);
    const fields = entryLines[0].split("\t");
    expect(fields).toHaveLength(7);
    expect(fields[0]).toBe("chatgpt.com"); // domain stripped
    expect(fields[1]).toBe("TRUE"); // includeSubdomains because of leading dot
    expect(fields[2]).toBe("/");
    expect(fields[3]).toBe("TRUE"); // secure
    expect(fields[4]).toBe("1234567890");
    expect(fields[5]).toBe("a");
    expect(fields[6]).toBe("v");
  });

  test("includeSubdomains=FALSE when the cookie has no leading dot", () => {
    const jar = toNetscapeCookieJar([
      { name: "a", value: "v", domain: "chatgpt.com", path: "/", secure: false },
    ]);
    const fields = jar.split("\n").find((l) => l && !l.startsWith("#"))!.split("\t");
    expect(fields[1]).toBe("FALSE");
    expect(fields[3]).toBe("FALSE");
  });

  test("emits 0 for missing expires (session cookie)", () => {
    const jar = toNetscapeCookieJar([cookie("a", "chatgpt.com", "v")]);
    const fields = jar.split("\n").find((l) => l && !l.startsWith("#"))!.split("\t");
    expect(fields[4]).toBe("0");
  });

  test("starts with header comments and ends with newline", () => {
    const jar = toNetscapeCookieJar([cookie("a", "chatgpt.com", "v")]);
    expect(jar.startsWith("# Netscape HTTP Cookie File\n")).toBe(true);
    expect(jar.endsWith("\n")).toBe(true);
  });

  test("does not emit volatile cookies", () => {
    const jar = toNetscapeCookieJar([
      cookie("a", "chatgpt.com", "keep"),
      cookie("conv_key_x", "chatgpt.com", "drop"),
    ]);
    expect(jar).toContain("\ta\t");
    expect(jar).not.toContain("conv_key_");
  });
});

describe("loadCookieExport", () => {
  test("upgrades a bare cookies array to a v1 envelope", async () => {
    const dir = await mkdtemp(join(tmpdir(), "cookies-export-"));
    const path = join(dir, "cookies.json");
    try {
      await writeFile(path, JSON.stringify([cookie("a", "chatgpt.com", "v")]));
      const result = await loadCookieExport(path);
      expect(result.version).toBe(1);
      expect(result.source).toBe("pro-cli-cdp");
      expect(result.cookies.map((c) => c.name)).toEqual(["a"]);
    } finally {
      await rm(dir, { recursive: true, force: true });
    }
  });

  test("upgrades a {cookies} object lacking a version field", async () => {
    const dir = await mkdtemp(join(tmpdir(), "cookies-export-"));
    const path = join(dir, "cookies.json");
    try {
      await writeFile(path, JSON.stringify({ cookies: [cookie("a", "chatgpt.com", "v")] }));
      const result = await loadCookieExport(path);
      expect(result.version).toBe(1);
      expect(result.cookies.map((c) => c.name)).toEqual(["a"]);
    } finally {
      await rm(dir, { recursive: true, force: true });
    }
  });

  test("returns a v1 envelope verbatim", async () => {
    const dir = await mkdtemp(join(tmpdir(), "cookies-export-"));
    const path = join(dir, "cookies.json");
    try {
      const obj = {
        version: 1 as const,
        generatedAt: "2026-01-01T00:00:00.000Z",
        source: "pro-cli-cdp" as const,
        targetUrl: "https://chatgpt.com/",
        origins: ["https://chatgpt.com/"],
        cookies: [cookie("a", "chatgpt.com", "v")],
      };
      await writeFile(path, JSON.stringify(obj));
      const result = await loadCookieExport(path);
      expect(result).toEqual(obj);
    } finally {
      await rm(dir, { recursive: true, force: true });
    }
  });
});

describe("chatGptOrigins", () => {
  test("returns the canonical origin list", () => {
    const origins = chatGptOrigins();
    expect(origins).toEqual([
      "https://chatgpt.com/",
      "https://auth.openai.com/",
      "https://openai.com/",
      "https://sentinel.openai.com/",
      "https://ws.chatgpt.com/",
    ]);
  });

  test("returns a fresh array — callers cannot mutate the canonical list", () => {
    // Regression guard: if someone returned the literal const array directly,
    // callers could push into it and corrupt every later auth capture.
    const a = chatGptOrigins();
    const b = chatGptOrigins();
    a.push("https://attacker.example/");
    expect(b).not.toContain("https://attacker.example/");
  });
});

function cookie(name: string, domain: string, value: string): BrowserCookie {
  return { name, value, domain, path: "/", secure: true };
}
`````

## File: tests/executor.test.ts
`````typescript
import { mkdtemp, rm } from "node:fs/promises";
import { join } from "node:path";
import { tmpdir } from "node:os";
import { describe, expect, test } from "bun:test";
import { ProError } from "../src/errors";
import {
  buildEphemeralJob,
  waitForJob,
  waitForTerminalJob,
  waitTimeoutError,
} from "../src/executor";
import { JobStore } from "../src/jobs";

async function withStore<T>(fn: (store: JobStore) => Promise<T>): Promise<T> {
  const dir = await mkdtemp(join(tmpdir(), "pro-executor-test-"));
  const store = await JobStore.open(join(dir, "jobs.sqlite"));
  try {
    return await fn(store);
  } finally {
    store.close();
    await rm(dir, { recursive: true, force: true });
  }
}

describe("job waiting", () => {
  test("soft wait returns current job state when the timeout expires", async () => {
    await withStore(async (store) => {
      const created = store.create({
        prompt: "large prompt",
        model: "gpt-5-5-pro",
        reasoning: "extended",
        options: {},
      });

      const outcome = await waitForJob(store, created.id, 1, 25);

      expect(outcome.timedOut).toBe(true);
      expect(outcome.status).toBe("queued");
      expect(outcome.job.id).toBe(created.id);
      expect(outcome.elapsedMs).toBeGreaterThanOrEqual(1);
      expect(outcome.timeoutMs).toBe(1);
      expect(outcome.pollMs).toBe(25);
    });
  });

  test("hard wait timeout includes current status and retry guidance", async () => {
    await withStore(async (store) => {
      const created = store.create({
        prompt: "large prompt",
        model: "gpt-5-5-pro",
        reasoning: "extended",
        options: {},
      });

      try {
        await waitForTerminalJob(store, created.id, 1, 25);
        throw new Error("Expected wait timeout.");
      } catch (error) {
        expect(error).toBeInstanceOf(ProError);
        const proError = error as ProError;
        expect(proError.code).toBe("WAIT_TIMEOUT");
        expect(proError.message).toContain("still queued");
        expect(proError.suggestions[0]).toContain(`pro-cli job wait ${created.id} --json`);
        expect(proError.details?.status).toBe("queued");
        expect(proError.details?.timeoutMs).toBe(1);
      }
    });
  });

  test("waitForJob returns immediately when the job is already succeeded", async () => {
    await withStore(async (store) => {
      const created = store.create({
        prompt: "x",
        model: "gpt-5-5-pro",
        reasoning: "standard",
        options: {},
      });
      store.markRunning(created.id);
      store.markSucceeded(created.id, "done");

      const t0 = Date.now();
      const outcome = await waitForJob(store, created.id, 5_000, 100);
      const elapsed = Date.now() - t0;

      expect(outcome.timedOut).toBe(false);
      expect(outcome.status).toBe("succeeded");
      expect(outcome.job.result).toBe("done");
      // Should not have polled — we returned on first check.
      expect(elapsed).toBeLessThan(80);
    });
  });

  test("waitForJob returns immediately when the job is already failed", async () => {
    await withStore(async (store) => {
      const created = store.create({
        prompt: "x",
        model: "gpt-5-5-pro",
        reasoning: "standard",
        options: {},
      });
      store.markRunning(created.id);
      store.markFailed(created.id, new ProError("BOOM", "blew up"));

      const outcome = await waitForJob(store, created.id, 5_000, 100);
      expect(outcome.timedOut).toBe(false);
      expect(outcome.status).toBe("failed");
    });
  });

  test("waitForJob returns immediately when the job is already cancelled", async () => {
    await withStore(async (store) => {
      const created = store.create({
        prompt: "x",
        model: "gpt-5-5-pro",
        reasoning: "standard",
        options: {},
      });
      store.cancel(created.id);

      const outcome = await waitForJob(store, created.id, 5_000, 100);
      expect(outcome.timedOut).toBe(false);
      expect(outcome.status).toBe("cancelled");
    });
  });

  test("waitForJob actually polls and returns when the job becomes terminal mid-wait", async () => {
    await withStore(async (store) => {
      // Regression guard: if the polling loop were broken to never re-fetch
      // (e.g. early return), the job would still be 'running' when the wait
      // returns. We explicitly check it picks up the terminal transition.
      const created = store.create({
        prompt: "x",
        model: "gpt-5-5-pro",
        reasoning: "standard",
        options: {},
      });
      store.markRunning(created.id);

      // Flip to succeeded after ~50ms while waitForJob polls every 20ms.
      const flipper = setTimeout(() => store.markSucceeded(created.id, "ok"), 50);
      try {
        const outcome = await waitForJob(store, created.id, 1_500, 20);
        expect(outcome.timedOut).toBe(false);
        expect(outcome.status).toBe("succeeded");
        expect(outcome.job.result).toBe("ok");
        // We should have polled at least twice (>= 40ms before pickup).
        expect(outcome.elapsedMs).toBeGreaterThanOrEqual(40);
      } finally {
        clearTimeout(flipper);
      }
    });
  });

  test("waitForJob with timeoutMs=0 means no timeout (terminal state still returns)", async () => {
    await withStore(async (store) => {
      const created = store.create({
        prompt: "x",
        model: "gpt-5-5-pro",
        reasoning: "standard",
        options: {},
      });
      store.markRunning(created.id);
      const flipper = setTimeout(() => store.markSucceeded(created.id, "done"), 30);
      try {
        const outcome = await waitForJob(store, created.id, 0, 10);
        expect(outcome.timedOut).toBe(false);
        expect(outcome.status).toBe("succeeded");
      } finally {
        clearTimeout(flipper);
      }
    });
  });

  test("waitForJob propagates JOB_NOT_FOUND when the id is missing", async () => {
    await withStore(async (store) => {
      try {
        await waitForJob(store, "job_missing", 10, 5);
        throw new Error("Expected JOB_NOT_FOUND.");
      } catch (error) {
        expect(error).toBeInstanceOf(ProError);
        expect((error as ProError).code).toBe("JOB_NOT_FOUND");
      }
    });
  });

  test("waitForTerminalJob returns the unredacted terminal record on success", async () => {
    await withStore(async (store) => {
      const created = store.create({
        prompt: "x",
        model: "gpt-5-5-pro",
        reasoning: "standard",
        options: {},
      });
      store.markRunning(created.id);
      store.markSucceeded(created.id, "result text");
      const job = await waitForTerminalJob(store, created.id, 1_000, 10);
      expect(job.status).toBe("succeeded");
      expect(job.result).toBe("result text");
      expect(job.id).toBe(created.id);
    });
  });
});

describe("buildEphemeralJob", () => {
  test("produces a job with status=running and an ask_ id (not job_)", () => {
    // The id prefix matters: the daemon and store distinguish ask_* from
    // job_* records. ask_* must never end up in the persistent jobs table.
    const job = buildEphemeralJob({
      prompt: "ephemeral prompt",
      model: "gpt-5-5-pro",
      reasoning: "extended",
      options: { timeoutMs: 5000 },
    });
    expect(job.id).toMatch(/^ask_[0-9a-f]{8}-/);
    expect(job.status).toBe("running");
    expect(job.prompt).toBe("ephemeral prompt");
    expect(job.model).toBe("gpt-5-5-pro");
    expect(job.reasoning).toBe("extended");
    expect(job.options).toEqual({ timeoutMs: 5000 });
    expect(job.result).toBeNull();
    expect(job.error).toBeNull();
    expect(job.createdAt).toBe(job.updatedAt);
  });

  test("does not share state across calls (deep options copy by reference is fine, but ids differ)", () => {
    const a = buildEphemeralJob({ prompt: "a", model: "gpt-5-5-pro", reasoning: "standard", options: {} });
    const b = buildEphemeralJob({ prompt: "b", model: "gpt-5-5-pro", reasoning: "standard", options: {} });
    expect(a.id).not.toBe(b.id);
  });
});

describe("waitTimeoutError", () => {
  test("includes job, status, elapsedMs, timeoutMs, pollMs in details", () => {
    const error = waitTimeoutError({
      job: {
        id: "job_x",
        status: "running",
        prompt: "p",
        model: "gpt-5-5-pro",
        reasoning: "standard",
        options: {},
        result: null,
        error: null,
        createdAt: "2026-01-01T00:00:00Z",
        updatedAt: "2026-01-01T00:00:00Z",
      },
      status: "running",
      timedOut: true,
      elapsedMs: 65_000,
      timeoutMs: 60_000,
      pollMs: 1000,
    });
    expect(error.code).toBe("WAIT_TIMEOUT");
    expect(error.message).toContain("still running");
    expect(error.message).toContain("1m 5s");
    expect(error.details?.elapsedMs).toBe(65_000);
    expect(error.details?.timeoutMs).toBe(60_000);
    expect(error.details?.pollMs).toBe(1000);
    // Suggestions name the job id and provide both polling and cancel paths.
    expect(error.suggestions.some((s) => s.includes("job_x"))).toBe(true);
    expect(error.suggestions.some((s) => s.includes("--soft-timeout"))).toBe(true);
    expect(error.suggestions.some((s) => s.includes("cancel"))).toBe(true);
  });

  test("formats sub-second durations in ms", () => {
    const error = waitTimeoutError({
      job: {
        id: "job_y",
        status: "queued",
        prompt: "p",
        model: "gpt-5-5-pro",
        reasoning: "standard",
        options: {},
        result: null,
        error: null,
        createdAt: "2026-01-01T00:00:00Z",
        updatedAt: "2026-01-01T00:00:00Z",
      },
      status: "queued",
      timedOut: true,
      elapsedMs: 250,
      timeoutMs: 200,
      pollMs: 50,
    });
    expect(error.message).toContain("250ms");
  });

  test("formats whole-minute durations without trailing seconds", () => {
    const error = waitTimeoutError({
      job: {
        id: "job_z",
        status: "running",
        prompt: "p",
        model: "gpt-5-5-pro",
        reasoning: "standard",
        options: {},
        result: null,
        error: null,
        createdAt: "2026-01-01T00:00:00Z",
        updatedAt: "2026-01-01T00:00:00Z",
      },
      status: "running",
      timedOut: true,
      elapsedMs: 120_000,
      timeoutMs: 120_000,
      pollMs: 1000,
    });
    expect(error.message).toMatch(/2m(?!\s\d)/);
  });
});
`````

## File: tests/jobs.test.ts
`````typescript
import { mkdtemp, rm } from "node:fs/promises";
import { join } from "node:path";
import { tmpdir } from "node:os";
import { describe, expect, test } from "bun:test";
import { ProError } from "../src/errors";
import { JobStore, redactJob, type JobRecord } from "../src/jobs";

async function withStore<T>(fn: (store: JobStore, dir: string) => Promise<T>): Promise<T> {
  const dir = await mkdtemp(join(tmpdir(), "pro-jobs-test-"));
  const store = await JobStore.open(join(dir, "jobs.sqlite"));
  try {
    return await fn(store, dir);
  } finally {
    store.close();
    await rm(dir, { recursive: true, force: true });
  }
}

describe("job store: create and read", () => {
  test("create produces a redacted record (prompt cleared, preview kept) but persists prompt internally", async () => {
    await withStore(async (store) => {
      const created = store.create({
        prompt: "the actual sensitive prompt body",
        model: "gpt-5-5-pro",
        reasoning: "standard",
        options: { temporary: true },
      });

      // Returned record is redacted for stdout safety.
      expect(created.id).toMatch(/^job_[0-9a-f]{8}-[0-9a-f]{4}-/);
      expect(created.status).toBe("queued");
      expect(created.prompt).toBe(""); // cleared
      expect(created.promptPreview).toBe("the actual sensitive prompt body");
      expect(created.model).toBe("gpt-5-5-pro");
      expect(created.reasoning).toBe("standard");
      expect(created.options).toEqual({ temporary: true });
      expect(created.result).toBeNull();
      expect(created.error).toBeNull();
      expect(created.hasResult).toBe(false);
      expect(created.createdAt).toBe(created.updatedAt);
      expect(() => new Date(created.createdAt).toISOString()).not.toThrow();

      // But internal get() returns the unredacted record (used by the daemon
      // and result endpoint).
      const internal = store.get(created.id);
      expect(internal.prompt).toBe("the actual sensitive prompt body");
      expect(internal.status).toBe("queued");
    });
  });

  test("get throws JOB_NOT_FOUND for an unknown id", async () => {
    await withStore(async (store) => {
      try {
        store.get("job_does_not_exist");
        throw new Error("Expected JOB_NOT_FOUND.");
      } catch (error) {
        expect(error).toBeInstanceOf(ProError);
        const proError = error as ProError;
        expect(proError.code).toBe("JOB_NOT_FOUND");
        expect(proError.suggestions[0]).toContain("pro-cli job list");
      }
    });
  });

  test("create generates unique ids across calls", async () => {
    await withStore(async (store) => {
      const ids = new Set<string>();
      for (let i = 0; i < 8; i += 1) {
        const job = store.create({ prompt: `p${i}`, model: "gpt-5-5-pro", reasoning: "standard", options: {} });
        ids.add(job.id);
      }
      expect(ids.size).toBe(8);
    });
  });

  test("list orders by created_at DESC and respects limit", async () => {
    await withStore(async (store) => {
      const ids: string[] = [];
      for (let i = 0; i < 5; i += 1) {
        // Force distinct timestamps even on fast systems.
        await wait(2);
        ids.push(store.create({ prompt: `p${i}`, model: "gpt-5-5-pro", reasoning: "standard", options: {} }).id);
      }
      const listed = store.list(3);
      expect(listed).toHaveLength(3);
      // Most recently created should be first.
      expect(listed[0].id).toBe(ids[ids.length - 1]);
      expect(listed[1].id).toBe(ids[ids.length - 2]);
      expect(listed[2].id).toBe(ids[ids.length - 3]);
      // List items are redacted (prompt cleared).
      for (const job of listed) expect(job.prompt).toBe("");
    });
  });
});

describe("job store: status transitions", () => {
  test("markRunning moves a queued job to running and updates updatedAt", async () => {
    await withStore(async (store) => {
      const created = store.create({ prompt: "x", model: "gpt-5-5-pro", reasoning: "standard", options: {} });
      await wait(2);
      const running = store.markRunning(created.id);
      expect(running.status).toBe("running");
      expect(Date.parse(running.updatedAt)).toBeGreaterThan(Date.parse(created.updatedAt));
    });
  });

  test("markRunning does not revive terminal jobs", async () => {
    await withStore(async (store) => {
      const succeeded = store.create({ prompt: "a", model: "gpt-5-5-pro", reasoning: "standard", options: {} });
      store.markRunning(succeeded.id);
      store.markSucceeded(succeeded.id, "ok");

      const failed = store.create({ prompt: "b", model: "gpt-5-5-pro", reasoning: "standard", options: {} });
      store.markRunning(failed.id);
      store.markFailed(failed.id, new ProError("BOOM", "failed"));

      const cancelled = store.create({ prompt: "c", model: "gpt-5-5-pro", reasoning: "standard", options: {} });
      store.cancel(cancelled.id);

      expect(store.markRunning(succeeded.id).status).toBe("succeeded");
      expect(store.markRunning(failed.id).status).toBe("failed");
      expect(store.markRunning(cancelled.id).status).toBe("cancelled");
      expect(store.get(succeeded.id).result).toBe("ok");
      expect(JSON.parse(store.get(failed.id).error as string).code).toBe("BOOM");
    });
  });

  test("markSucceeded only writes when current status is running (regression: late success after cancel)", async () => {
    await withStore(async (store) => {
      const created = store.create({ prompt: "x", model: "gpt-5-5-pro", reasoning: "standard", options: {} });
      store.markRunning(created.id);
      store.cancel(created.id);
      // Late success arrives after cancellation — must NOT overwrite.
      store.markSucceeded(created.id, "late result");
      const job = store.get(created.id);
      expect(job.status).toBe("cancelled");
      expect(job.result).toBeNull();
      expect(job.error).not.toBeNull();
      const errorPayload = JSON.parse(job.error as string);
      expect(errorPayload.code).toBe("CANCELLED");
    });
  });

  test("markFailed only writes when current status is running", async () => {
    await withStore(async (store) => {
      const created = store.create({ prompt: "x", model: "gpt-5-5-pro", reasoning: "standard", options: {} });
      store.markRunning(created.id);
      store.cancel(created.id);
      const failure = new ProError("LATE_FAILURE", "Pretend the worker died.");
      store.markFailed(created.id, failure);
      const job = store.get(created.id);
      expect(job.status).toBe("cancelled"); // unchanged
      const errorPayload = JSON.parse(job.error as string);
      expect(errorPayload.code).toBe("CANCELLED"); // not LATE_FAILURE
    });
  });

  test("markSucceeded from running stores result and clears error", async () => {
    await withStore(async (store) => {
      const created = store.create({ prompt: "x", model: "gpt-5-5-pro", reasoning: "standard", options: {} });
      store.markRunning(created.id);
      store.markSucceeded(created.id, "Here is the answer.");
      const job = store.get(created.id);
      expect(job.status).toBe("succeeded");
      expect(job.result).toBe("Here is the answer.");
      expect(job.error).toBeNull();
    });
  });

  test("markFailed from running writes a JSON error payload with code/message/suggestions", async () => {
    await withStore(async (store) => {
      const created = store.create({ prompt: "x", model: "gpt-5-5-pro", reasoning: "standard", options: {} });
      store.markRunning(created.id);
      const failure = new ProError("UPSTREAM_REJECTED", "ChatGPT returned 500.", {
        suggestions: ["Retry later."],
        details: { status: 500 },
      });
      store.markFailed(created.id, failure);
      const job = store.get(created.id);
      expect(job.status).toBe("failed");
      const payload = JSON.parse(job.error as string);
      expect(payload.code).toBe("UPSTREAM_REJECTED");
      expect(payload.message).toBe("ChatGPT returned 500.");
      expect(payload.suggestions).toEqual(["Retry later."]);
      expect(payload.details).toEqual({ status: 500 });
    });
  });

  test("cancel is idempotent on terminal statuses (succeeded/failed/cancelled stay)", async () => {
    await withStore(async (store) => {
      const a = store.create({ prompt: "a", model: "gpt-5-5-pro", reasoning: "standard", options: {} });
      const b = store.create({ prompt: "b", model: "gpt-5-5-pro", reasoning: "standard", options: {} });
      store.markRunning(a.id);
      store.markSucceeded(a.id, "ok");
      store.markRunning(b.id);
      store.markFailed(b.id, new ProError("X", "x"));

      const aCancel = store.cancel(a.id);
      const bCancel = store.cancel(b.id);
      expect(aCancel.status).toBe("succeeded");
      expect(bCancel.status).toBe("failed");
      // Result/error preserved.
      expect(store.get(a.id).result).toBe("ok");
      expect(store.get(b.id).status).toBe("failed");
    });
  });

  test("cancel can run from queued (no markRunning required)", async () => {
    await withStore(async (store) => {
      const created = store.create({ prompt: "x", model: "gpt-5-5-pro", reasoning: "standard", options: {} });
      const cancelled = store.cancel(created.id);
      expect(cancelled.status).toBe("cancelled");
    });
  });
});

describe("job store: claim semantics", () => {
  test("claimQueued atomically transitions queued→running exactly once", async () => {
    await withStore(async (store) => {
      const created = store.create({ prompt: "x", model: "gpt-5-5-pro", reasoning: "standard", options: {} });
      const first = store.claimQueued(created.id);
      const second = store.claimQueued(created.id);
      expect(first?.status).toBe("running");
      expect(second).toBeNull();
    });
  });

  test("claimQueued returns null when the job is in any non-queued state", async () => {
    await withStore(async (store) => {
      const created = store.create({ prompt: "x", model: "gpt-5-5-pro", reasoning: "standard", options: {} });
      store.markRunning(created.id);
      expect(store.claimQueued(created.id)).toBeNull();
    });
  });

  test("claimNextQueued picks the oldest queued job and skips non-queued ones", async () => {
    await withStore(async (store) => {
      const a = store.create({ prompt: "a", model: "gpt-5-5-pro", reasoning: "standard", options: {} });
      await wait(2);
      const b = store.create({ prompt: "b", model: "gpt-5-5-pro", reasoning: "standard", options: {} });
      await wait(2);
      const c = store.create({ prompt: "c", model: "gpt-5-5-pro", reasoning: "standard", options: {} });

      // Simulate that `a` was already started and finished.
      store.markRunning(a.id);
      store.markSucceeded(a.id, "done");

      const next = store.claimNextQueued();
      expect(next?.id).toBe(b.id);
      expect(next?.status).toBe("running");

      const after = store.claimNextQueued();
      expect(after?.id).toBe(c.id);

      expect(store.claimNextQueued()).toBeNull();
    });
  });
});

describe("job store: limits observations", () => {
  test("recordLimits + latestLimits returns the most recent observation per feature", async () => {
    await withStore(async (store) => {
      store.recordLimits([{ feature_name: "deep_research", remaining: 100, reset_after: "2026-06-01T00:00:00Z" }], "job_a");
      // Same observed_at would conflict; force a different timestamp.
      await wait(2);
      store.recordLimits(
        [
          { feature_name: "deep_research", remaining: 90, reset_after: "2026-06-02T00:00:00Z" },
          { feature_name: "odyssey", remaining: 5, reset_after: null },
        ],
        "job_b",
      );

      const latest = store.latestLimits();
      expect(latest).toHaveLength(2);
      const byName = Object.fromEntries(latest.map((l) => [l.featureName, l] as const));
      expect(byName.deep_research.remaining).toBe(90);
      expect(byName.deep_research.resetAfter).toBe("2026-06-02T00:00:00Z");
      expect(byName.deep_research.jobId).toBe("job_b");
      expect(byName.odyssey.remaining).toBe(5);
      expect(byName.odyssey.resetAfter).toBeNull();
    });
  });

  test("recordLimits is a no-op for an empty list", async () => {
    await withStore(async (store) => {
      store.recordLimits([], null);
      expect(store.latestLimits()).toEqual([]);
    });
  });

  test("latestLimits returns features sorted alphabetically", async () => {
    await withStore(async (store) => {
      store.recordLimits(
        [
          { feature_name: "z_feature", remaining: 1, reset_after: null },
          { feature_name: "a_feature", remaining: 1, reset_after: null },
          { feature_name: "m_feature", remaining: 1, reset_after: null },
        ],
        null,
      );
      const latest = store.latestLimits();
      expect(latest.map((l) => l.featureName)).toEqual(["a_feature", "m_feature", "z_feature"]);
    });
  });
});

describe("job store: persistence", () => {
  test("data survives close + reopen at the same path", async () => {
    const dir = await mkdtemp(join(tmpdir(), "pro-jobs-persist-"));
    const path = join(dir, "jobs.sqlite");
    let id = "";
    try {
      const first = await JobStore.open(path);
      try {
        const created = first.create({ prompt: "p", model: "gpt-5-5-pro", reasoning: "standard", options: {} });
        first.markRunning(created.id);
        first.markSucceeded(created.id, "result body");
        id = created.id;
      } finally {
        first.close();
      }

      const second = await JobStore.open(path);
      try {
        const job = second.get(id);
        expect(job.status).toBe("succeeded");
        expect(job.result).toBe("result body");
        expect(job.prompt).toBe("p");
      } finally {
        second.close();
      }
    } finally {
      await rm(dir, { recursive: true, force: true });
    }
  });

  test("opens cleanly when the parent directory does not exist yet", async () => {
    const dir = await mkdtemp(join(tmpdir(), "pro-jobs-fresh-"));
    try {
      // Nested path must be created.
      const store = await JobStore.open(join(dir, "deep", "nested", "jobs.sqlite"));
      try {
        const created = store.create({ prompt: "p", model: "gpt-5-5-pro", reasoning: "standard", options: {} });
        expect(created.status).toBe("queued");
      } finally {
        store.close();
      }
    } finally {
      await rm(dir, { recursive: true, force: true });
    }
  });
});

describe("redactJob helper", () => {
  test("clears prompt/result and emits compact previews + hasResult flag", () => {
    const job: JobRecord = {
      id: "job_x",
      status: "succeeded",
      prompt: "Some\n\n\nlong\tprompt body".repeat(10),
      model: "gpt-5-5-pro",
      reasoning: "standard",
      options: {},
      result: "Result text",
      error: null,
      createdAt: "2026-01-01T00:00:00Z",
      updatedAt: "2026-01-01T00:00:00Z",
    };
    const redacted = redactJob(job);
    expect(redacted.prompt).toBe("");
    expect(redacted.result).toBeNull();
    expect(redacted.hasResult).toBe(true);
    expect(redacted.resultPreview).toBe("Result text");
    expect(redacted.promptPreview.length).toBeLessThanOrEqual(160);
    // Whitespace collapsed in preview.
    expect(redacted.promptPreview).not.toContain("\n");
    expect(redacted.promptPreview).not.toContain("\t");
  });

  test("omits resultPreview and sets hasResult=false when result is null", () => {
    const job: JobRecord = {
      id: "job_x",
      status: "queued",
      prompt: "p",
      model: "gpt-5-5-pro",
      reasoning: "standard",
      options: {},
      result: null,
      error: null,
      createdAt: "2026-01-01T00:00:00Z",
      updatedAt: "2026-01-01T00:00:00Z",
    };
    const redacted = redactJob(job);
    expect(redacted.hasResult).toBe(false);
    expect("resultPreview" in redacted).toBe(false);
  });

  test("treats an empty-string result as a real result", () => {
    const job: JobRecord = {
      id: "job_x",
      status: "succeeded",
      prompt: "p",
      model: "gpt-5-5-pro",
      reasoning: "standard",
      options: {},
      result: "",
      error: null,
      createdAt: "2026-01-01T00:00:00Z",
      updatedAt: "2026-01-01T00:00:00Z",
    };
    const redacted = redactJob(job);
    expect(redacted.result).toBeNull();
    expect(redacted.hasResult).toBe(true);
    expect(redacted.resultPreview).toBe("");
  });

  test("truncates oversized previews with ellipsis", () => {
    const job: JobRecord = {
      id: "job_x",
      status: "queued",
      prompt: "x".repeat(500),
      model: "gpt-5-5-pro",
      reasoning: "standard",
      options: {},
      result: "y".repeat(500),
      error: null,
      createdAt: "2026-01-01T00:00:00Z",
      updatedAt: "2026-01-01T00:00:00Z",
    };
    const redacted = redactJob(job);
    expect(redacted.promptPreview.length).toBe(160);
    expect(redacted.promptPreview.endsWith("…")).toBe(true);
    expect(redacted.resultPreview!.length).toBe(240);
    expect(redacted.resultPreview!.endsWith("…")).toBe(true);
  });
});

function wait(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms));
}
`````

## File: tests/limits.test.ts
`````typescript
import { describe, expect, test } from "bun:test";
import { summarizeAccountResponse } from "../src/limits";
import { extractLimitsProgress } from "../src/transport";

describe("summarizeAccountResponse", () => {
  test("extracts pro plan facts from accounts/check shape", () => {
    const body = JSON.stringify({
      accounts: {
        default: {
          account: { plan_type: "free", structure: "personal" },
          features: ["gpt5"],
          entitlement: { has_active_subscription: false, subscription_plan: null },
        },
        "uuid-1": {
          account: { plan_type: "pro" },
          features: ["gpt5_pro", "o3_pro"],
          entitlement: {
            has_active_subscription: true,
            subscription_plan: "chatgptpro",
            expires_at: "2026-05-27T03:16:31+00:00",
            renews_at: "2026-05-26T21:16:31+00:00",
            cancels_at: null,
            billing_period: "monthly",
          },
          last_active_subscription: { will_renew: true },
        },
      },
      account_ordering: ["uuid-1"],
    });
    const summary = summarizeAccountResponse(body);
    expect(summary.planType).toBe("pro");
    expect(summary.subscriptionPlan).toBe("chatgptpro");
    expect(summary.hasActiveSubscription).toBe(true);
    expect(summary.expiresAt).toBe("2026-05-27T03:16:31+00:00");
    expect(summary.renewsAt).toBe("2026-05-26T21:16:31+00:00");
    expect(summary.billingPeriod).toBe("monthly");
    expect(summary.willRenew).toBe(true);
    expect(summary.features).toContain("gpt5_pro");
    expect(summary.features).toContain("o3_pro");
  });

  test("falls back to default account when ordering is empty", () => {
    const body = JSON.stringify({
      accounts: {
        default: {
          account: { plan_type: "free" },
          features: [],
          entitlement: { has_active_subscription: false },
        },
      },
      account_ordering: [],
    });
    const summary = summarizeAccountResponse(body);
    expect(summary.planType).toBe("free");
    expect(summary.hasActiveSubscription).toBe(false);
  });

  test("prefers a real account over the default sentinel when ordering is empty", () => {
    const body = JSON.stringify({
      accounts: {
        default: {
          account: { plan_type: "free" },
          features: ["base"],
          entitlement: { has_active_subscription: false },
        },
        "uuid-1": {
          account: { plan_type: "pro" },
          features: ["gpt5_pro"],
          entitlement: { has_active_subscription: true, subscription_plan: "chatgptpro" },
        },
      },
      account_ordering: [],
    });
    const summary = summarizeAccountResponse(body);
    expect(summary.planType).toBe("pro");
    expect(summary.subscriptionPlan).toBe("chatgptpro");
    expect(summary.features).toEqual(["gpt5_pro"]);
  });

  test("returns empty summary on missing accounts shape", () => {
    const summary = summarizeAccountResponse(JSON.stringify({}));
    expect(summary.planType).toBe(null);
    expect(summary.features).toEqual([]);
  });

  test("throws on unparseable body", () => {
    expect(() => summarizeAccountResponse("not json")).toThrow();
  });
});

describe("extractLimitsProgress", () => {
  test("extracts limits from a top-level conversation_detail_metadata event", () => {
    const event = {
      type: "conversation_detail_metadata",
      limits_progress: [
        { feature_name: "deep_research", remaining: 250, reset_after: "2026-06-07T18:34:14.421525+00:00" },
        { feature_name: "odyssey", remaining: 398, reset_after: "2026-05-17T21:31:20.421544+00:00" },
      ],
      model_limits: [],
    };
    const observations = extractLimitsProgress(event);
    expect(observations).toHaveLength(2);
    expect(observations[0]).toEqual({
      feature_name: "deep_research",
      remaining: 250,
      reset_after: "2026-06-07T18:34:14.421525+00:00",
    });
  });

  test("handles wrapped delta-style payload at event.v", () => {
    const event = {
      v: {
        type: "conversation_detail_metadata",
        limits_progress: [{ feature_name: "deep_research", remaining: 100 }],
      },
    };
    const observations = extractLimitsProgress(event);
    expect(observations).toHaveLength(1);
    expect(observations[0].feature_name).toBe("deep_research");
    expect(observations[0].reset_after).toBe(null);
  });

  test("ignores unrelated events", () => {
    expect(extractLimitsProgress({ type: "input_message" })).toEqual([]);
    expect(extractLimitsProgress(null)).toEqual([]);
    expect(extractLimitsProgress({ type: "conversation_detail_metadata" })).toEqual([]);
  });

  test("dedupes by feature_name", () => {
    const event = {
      type: "conversation_detail_metadata",
      limits_progress: [
        { feature_name: "deep_research", remaining: 250 },
        { feature_name: "deep_research", remaining: 240 },
      ],
    };
    expect(extractLimitsProgress(event)).toHaveLength(1);
  });

  test("skips entries with missing fields", () => {
    const event = {
      type: "conversation_detail_metadata",
      limits_progress: [
        { feature_name: "ok", remaining: 5 },
        { feature_name: 123, remaining: 5 },
        { feature_name: "no_remaining" },
      ],
    };
    const observations = extractLimitsProgress(event);
    expect(observations).toHaveLength(1);
    expect(observations[0].feature_name).toBe("ok");
  });

  test("dedupes by feature_name keeping the FIRST occurrence (regression guard)", () => {
    // Behavior matters: agent-facing counters should reflect the upstream
    // event order. If a refactor flipped to last-wins, downstream display
    // would show stale numbers from earlier events.
    const event = {
      type: "conversation_detail_metadata",
      limits_progress: [
        { feature_name: "deep_research", remaining: 250, reset_after: "early" },
        { feature_name: "deep_research", remaining: 100, reset_after: "later" },
      ],
    };
    const observations = extractLimitsProgress(event);
    expect(observations).toHaveLength(1);
    // First wins:
    expect(observations[0].remaining).toBe(250);
    expect(observations[0].reset_after).toBe("early");
  });

  test("normalizes a missing reset_after to null (callers can store null in SQLite)", () => {
    // The persistence layer is `reset_after TEXT` which accepts null but not
    // undefined. extractLimitsProgress must coerce.
    const event = {
      type: "conversation_detail_metadata",
      limits_progress: [{ feature_name: "ok", remaining: 5 }],
    };
    expect(extractLimitsProgress(event)[0].reset_after).toBeNull();
  });
});

describe("summarizeAccountResponse: edge cases", () => {
  test("falls back to default when account_ordering points at a non-existent uuid", async () => {
    // Catches a regression where the stale ordering slug would crash or
    // return empty rather than degrading gracefully.
    const body = JSON.stringify({
      accounts: {
        default: {
          account: { plan_type: "free" },
          features: ["base"],
          entitlement: { has_active_subscription: false, subscription_plan: null },
        },
      },
      account_ordering: ["uuid-missing-from-accounts"],
    });
    const summary = summarizeAccountResponse(body);
    // Either falls back to default (preferred) or the only non-default key
    // (none here). With only 'default', it should land on default.
    expect(summary.planType).toBe("free");
  });

  test("will_renew=false is preserved verbatim", async () => {
    const body = JSON.stringify({
      accounts: {
        default: {
          account: { plan_type: "pro" },
          features: ["pro"],
          entitlement: { has_active_subscription: true, subscription_plan: "chatgptpro" },
          last_active_subscription: { will_renew: false },
        },
      },
      account_ordering: [],
    });
    const summary = summarizeAccountResponse(body);
    expect(summary.willRenew).toBe(false);
  });

  test("will_renew with non-boolean type becomes null (defensive)", async () => {
    const body = JSON.stringify({
      accounts: {
        default: {
          account: { plan_type: "pro" },
          features: [],
          entitlement: { has_active_subscription: true },
          last_active_subscription: { will_renew: "yes" }, // wrong shape
        },
      },
      account_ordering: [],
    });
    const summary = summarizeAccountResponse(body);
    expect(summary.willRenew).toBeNull();
  });
});
`````

## File: tests/models.test.ts
`````typescript
import { mkdtemp, rm, writeFile } from "node:fs/promises";
import { join } from "node:path";
import { tmpdir } from "node:os";
import { afterEach, describe, expect, test } from "bun:test";
import { listModels } from "../src/models";

const originalFetch = globalThis.fetch;

afterEach(() => {
  globalThis.fetch = originalFetch;
});

async function withTokenFile<T>(fn: (path: string) => Promise<T>): Promise<T> {
  const dir = await mkdtemp(join(tmpdir(), "pro-model-test-"));
  const path = join(dir, "token.json");
  try {
    await writeFile(
      path,
      JSON.stringify({
        version: 1,
        generatedAt: new Date().toISOString(),
        source: "pro-cli-cdp-page",
        accessToken: fakeJwt(),
        accountId: "acct_test",
        expiresAt: new Date(Date.now() + 60 * 60 * 1000).toISOString(),
      }),
    );
    return await fn(path);
  } finally {
    await rm(dir, { recursive: true, force: true });
  }
}

describe("model discovery", () => {
  test("falls back to static models when no session token is available", async () => {
    const result = await listModels({ sessionTokenPath: "/tmp/pro-query-missing-token.json" });

    expect(result.source).toBe("static");
    expect(result.warning).toContain("No captured ChatGPT session token");
    expect(result.defaultModel).toBe("gpt-5-5-pro");
    expect(result.models.map((model) => model.id)).not.toContain("auto");
    expect(result.models.map((model) => model.id)).toContain("gpt-5-5-pro");
    expect(result.models.map((model) => model.id)).toContain("gpt-4-5");
    expect(result.models.find((model) => model.id === "research")).toMatchObject({
      label: "Deep Research",
      reasoningLevels: ["standard", "extended"],
      reasoningType: "pro",
      configurableThinkingEffort: true,
    });
  });

  test("loads live ChatGPT model catalog with bearer auth", async () => {
    await withTokenFile(async (sessionTokenPath) => {
      let authHeader = "";
      let accountHeader = "";
      globalThis.fetch = (async (_url: string | URL | Request, init?: RequestInit) => {
        const headers = new Headers(init?.headers);
        authHeader = headers.get("authorization") ?? "";
        accountHeader = headers.get("chatgpt-account-id") ?? "";
        return Response.json({
          default_model_slug: "gpt-5-5",
          model_picker_version: 2,
          models: [
            {
              slug: "gpt-5-5-pro",
              title: "GPT-5.5 Pro",
              max_tokens: 272000,
              reasoning_type: "pro",
              configurable_thinking_effort: true,
              thinking_efforts: [
                { thinking_effort: "standard", short_label: "Standard" },
                { thinking_effort: "extended", short_label: "Extended" },
              ],
              enabled_tools: [{ type: "python" }, { type: "web" }],
            },
          ],
        });
      }) as unknown as typeof fetch;

      const result = await listModels({ sessionTokenPath });

      expect(authHeader).toStartWith("Bearer ");
      expect(accountHeader).toBe("acct_test");
      expect(result.source).toBe("live");
      expect(result.defaultModel).toBe("gpt-5-5-pro");
      expect(result.chatgptDefaultModel).toBe("gpt-5-5");
      expect(result.modelPickerVersion).toBe(2);
      expect(result.models[0]).toMatchObject({
        id: "gpt-5-5-pro",
        label: "GPT-5.5 Pro",
        default: true,
        maxTokens: 272000,
        reasoningLevels: ["standard", "extended"],
        enabledTools: ["python", "web"],
      });
    });
  });

  test("calls the documented ChatGPT models endpoint", async () => {
    // Lock down which URL we hit so a refactor cannot silently retarget
    // it to a different (perhaps internal) endpoint.
    await withTokenFile(async (sessionTokenPath) => {
      let observedUrl = "";
      globalThis.fetch = (async (url: string | URL | Request) => {
        observedUrl = String(url);
        return Response.json({
          default_model_slug: "gpt-5-5",
          models: [{ slug: "gpt-5-5-pro", title: "GPT-5.5 Pro", reasoning_type: "pro" }],
        });
      }) as unknown as typeof fetch;
      const result = await listModels({ sessionTokenPath });
      expect(observedUrl).toBe("https://chatgpt.com/backend-api/models");
      expect(result.source).toBe("live");
    });
  });

  test("filters auto and malformed live model entries without dropping usable entries", async () => {
    await withTokenFile(async (sessionTokenPath) => {
      globalThis.fetch = (async () =>
        Response.json({
          default_model_slug: "auto",
          models: [
            { slug: "auto", title: "Auto" },
            { slug: 123, title: "Bad slug" },
            { slug: "gpt-5-5-pro", title: "GPT-5.5 Pro", reasoning_type: "pro" },
          ],
        })) as unknown as typeof fetch;

      const result = await listModels({ sessionTokenPath });
      expect(result.source).toBe("live");
      expect(result.models.map((model) => model.id)).toEqual(["gpt-5-5-pro"]);
    });
  });

  test("overrides live Deep Research capability with observed web UI reasoning levels", async () => {
    await withTokenFile(async (sessionTokenPath) => {
      globalThis.fetch = (async () =>
        Response.json({
          default_model_slug: "gpt-5-5",
          models: [
            {
              slug: "research",
              title: "Deep Research",
              reasoning_type: "none",
              thinking_efforts: [],
              configurable_thinking_effort: false,
            },
          ],
        })) as unknown as typeof fetch;

      const result = await listModels({ sessionTokenPath });

      expect(result.source).toBe("live");
      expect(result.models[0]).toMatchObject({
        id: "research",
        label: "Deep Research",
        reasoningLevels: ["standard", "extended"],
        reasoningType: "pro",
        configurableThinkingEffort: true,
      });
    });
  });

  test("falls back to static models on upstream HTTP error (does not crash)", async () => {
    // Regression guard: a 500 from the catalog endpoint should NOT block
    // the user from using known-good static models.
    await withTokenFile(async (sessionTokenPath) => {
      globalThis.fetch = (async () =>
        new Response("upstream broke", { status: 500 })) as unknown as typeof fetch;
      const result = await listModels({ sessionTokenPath });
      expect(result.source).toBe("static");
      expect(result.models.map((m) => m.id)).toContain("gpt-5-5-pro");
    });
  });

  test("falls back to static models when fetch itself rejects", async () => {
    await withTokenFile(async (sessionTokenPath) => {
      globalThis.fetch = (async () => {
        throw new Error("ECONNRESET");
      }) as unknown as typeof fetch;
      const result = await listModels({ sessionTokenPath });
      expect(result.source).toBe("static");
    });
  });

  test("falls back to static models when the live catalog has no usable entries", async () => {
    await withTokenFile(async (sessionTokenPath) => {
      globalThis.fetch = (async () =>
        Response.json({
          default_model_slug: "gpt-5-5",
          models: [
            { slug: "auto", title: "Auto" },
            { slug: "missing-title" },
            { title: "Missing slug" },
          ],
        })) as unknown as typeof fetch;

      const result = await listModels({ sessionTokenPath });

      expect(result.source).toBe("static");
      expect(result.warning).toContain("no usable model entries");
      expect(result.models.map((model) => model.id)).toContain("gpt-5-5-pro");
    });
  });

  test("falls back to static models when the token lacks an account id and does not call fetch", async () => {
    const dir = await mkdtemp(join(tmpdir(), "pro-model-no-account-"));
    const path = join(dir, "token.json");
    try {
      await writeFile(
        path,
        JSON.stringify({
          version: 1,
          generatedAt: new Date().toISOString(),
          source: "pro-cli-cdp-page",
          accessToken: fakeJwt(),
          expiresAt: new Date(Date.now() + 60 * 60 * 1000).toISOString(),
        }),
      );
      let fetchCalled = false;
      globalThis.fetch = (async () => {
        fetchCalled = true;
        return Response.json({ models: [] });
      }) as unknown as typeof fetch;

      const result = await listModels({ sessionTokenPath: path });

      expect(result.source).toBe("static");
      expect(result.warning).toContain("no account id");
      expect(fetchCalled).toBe(false);
    } finally {
      await rm(dir, { recursive: true, force: true });
    }
  });

  test("falls back to static models for an expired session token (does not call fetch)", async () => {
    const dir = await mkdtemp(join(tmpdir(), "pro-model-expired-"));
    const path = join(dir, "token.json");
    try {
      await writeFile(
        path,
        JSON.stringify({
          version: 1,
          generatedAt: new Date().toISOString(),
          source: "pro-cli-cdp-page",
          accessToken: fakeJwt(),
          accountId: "acct_test",
          expiresAt: new Date(Date.now() - 60_000).toISOString(),
        }),
      );
      let fetchCalled = false;
      globalThis.fetch = (async () => {
        fetchCalled = true;
        return new Response("{}");
      }) as unknown as typeof fetch;
      const result = await listModels({ sessionTokenPath: path });
      expect(result.source).toBe("static");
      expect(fetchCalled).toBe(false);
    } finally {
      await rm(dir, { recursive: true, force: true });
    }
  });
});

function fakeJwt(): string {
  const payload = {
    exp: Math.floor(Date.now() / 1000) + 3600,
    "https://api.openai.com/auth": { chatgpt_account_id: "acct_test" },
  };
  return ["header", Buffer.from(JSON.stringify(payload)).toString("base64url"), "sig"].join(".");
}
`````

## File: tests/odds.test.ts
`````typescript
import { describe, expect, test } from "bun:test";
import {
  aggregateOdds,
  buildOddsInstructions,
  buildOddsPrompt,
  parseOddsResponse,
} from "../src/odds";

describe("parseOddsResponse", () => {
  test("parses bare integer", () => {
    expect(parseOddsResponse("92", false).value).toBe(92);
  });

  test("trims whitespace", () => {
    expect(parseOddsResponse("  47\n", false).value).toBe(47);
  });

  test("falls back to first integer in noisy text", () => {
    expect(parseOddsResponse("My estimate: 73.", false).value).toBe(73);
  });

  test("rejects out-of-range integers", () => {
    expect(parseOddsResponse("250", false).value).toBe(null);
    expect(parseOddsResponse("999", false).value).toBe(null);
  });

  test("forbids 50 by default and flags it", () => {
    const r = parseOddsResponse("50", false);
    expect(r.value).toBe(null);
    expect(r.rejectedFifty).toBe(true);
  });

  test("allows 50 when allowFifty is true", () => {
    expect(parseOddsResponse("50", true).value).toBe(50);
  });

  test("returns null for non-numeric responses", () => {
    expect(parseOddsResponse("I cannot answer.", false).value).toBe(null);
  });

  test("accepts boundary 0 and 100", () => {
    expect(parseOddsResponse("0", false).value).toBe(0);
    expect(parseOddsResponse("100", false).value).toBe(100);
  });

  test("rejects negative integers instead of stripping the sign", () => {
    const r = parseOddsResponse("-73", false);
    expect(r.value).toBe(null);
    expect(r.rejectedFifty).toBe(false);
  });

  test("rejects decimal probabilities instead of truncating at the dot", () => {
    const r = parseOddsResponse("My estimate: 73.5%", false);
    expect(r.value).toBe(null);
    expect(r.rejectedFifty).toBe(false);
  });

  test("picks the FIRST integer when noisy text has multiple candidates", () => {
    // Regression guard: if the regex changes to greedy, this could pick 99.
    expect(parseOddsResponse("Between 60 and 99 percent.", false).value).toBe(60);
  });

  test("rejects multi-digit values starting with leading zero only when also out of range", () => {
    // "007" parses as 7 (loose match) — not rejected since 7 is in range.
    expect(parseOddsResponse("007", false).value).toBe(7);
  });
});

describe("aggregateOdds", () => {
  test("mean", () => {
    expect(aggregateOdds([60, 70, 80], "mean")).toBe(70);
  });

  test("median odd length", () => {
    expect(aggregateOdds([60, 95, 70], "median")).toBe(70);
  });

  test("median even length", () => {
    expect(aggregateOdds([60, 70, 80, 90], "median")).toBe(75);
  });

  test("trimmed-mean drops ends", () => {
    expect(aggregateOdds([0, 50, 50, 50, 100], "trimmed-mean")).toBe(50);
  });

  test("trimmed-mean falls back to mean below threshold", () => {
    expect(aggregateOdds([10, 90], "trimmed-mean")).toBe(50);
  });

  test("throws on empty input", () => {
    expect(() => aggregateOdds([], "mean")).toThrow();
  });

  test("single value passes through every aggregator", () => {
    expect(aggregateOdds([42], "mean")).toBe(42);
    expect(aggregateOdds([42], "median")).toBe(42);
    expect(aggregateOdds([42], "trimmed-mean")).toBe(42);
  });

  test("trimmed-mean uses 10% trim minimum 1 from each end", () => {
    // values length 10 → trim = max(1, floor(1)) = 1 from each end → mean of middle 8.
    expect(aggregateOdds([0, 50, 50, 50, 50, 50, 50, 50, 50, 100], "trimmed-mean")).toBe(50);
  });

  test("aggregators do not mutate the input array", () => {
    // Regression guard: a refactor that sorts in place would corrupt
    // attempt-level data shown to users.
    const input = [60, 95, 70];
    aggregateOdds(input, "median");
    expect(input).toEqual([60, 95, 70]);
  });
});

describe("buildOddsPrompt", () => {
  test("includes context block when provided", () => {
    const p = buildOddsPrompt("Will X happen?", "Evidence A about X.");
    expect(p).toContain("CONTEXT:");
    expect(p).toContain("Evidence A about X.");
    expect(p).toContain("QUESTION:");
    expect(p).toContain("Will X happen?");
  });

  test("omits context block when absent", () => {
    const p = buildOddsPrompt("Will X happen?");
    expect(p).not.toContain("CONTEXT:");
    expect(p).toContain("Will X happen?");
  });

  test("omits context block when only whitespace", () => {
    const p = buildOddsPrompt("Will X happen?", "   \n  ");
    expect(p).not.toContain("CONTEXT:");
  });
});

describe("buildOddsInstructions", () => {
  test("forbids 50 by default", () => {
    const text = buildOddsInstructions(false);
    expect(text).toContain("Do NOT output 50");
    expect(text).toContain("integer between 0 and 100");
  });

  test("allows 50 when configured", () => {
    const text = buildOddsInstructions(true);
    expect(text).not.toContain("Do NOT output 50");
    expect(text).toContain("output 50");
  });
});
`````

## File: tests/output.test.ts
`````typescript
import { describe, expect, test } from "bun:test";
import { ProError } from "../src/errors";
import { renderText, writeError, writeSuccess, type CliIO } from "../src/output";

function captureIO(): CliIO & { lines: { stdout: string; stderr: string } } {
  let stdout = "";
  let stderr = "";
  return {
    stdout: (text: string) => {
      stdout += text;
    },
    stderr: (text: string) => {
      stderr += text;
    },
    stdoutIsTTY: false,
    env: {},
    cwd: "/tmp",
    lines: {
      get stdout() {
        return stdout;
      },
      get stderr() {
        return stderr;
      },
    },
  } as unknown as CliIO & { lines: { stdout: string; stderr: string } };
}

describe("writeSuccess: JSON envelope", () => {
  test("wraps payload in { ok: true, data: ... }", () => {
    const io = captureIO();
    writeSuccess(io, { json: true }, { hello: "world" });
    const payload = JSON.parse(io.lines.stdout);
    expect(payload.ok).toBe(true);
    expect(payload.data.hello).toBe("world");
  });

  test("appends a single trailing newline (line-delimited JSON safe)", () => {
    const io = captureIO();
    writeSuccess(io, { json: true }, { hello: "world" });
    expect(io.lines.stdout.endsWith("\n")).toBe(true);
    expect(io.lines.stdout.split("\n").filter(Boolean)).toHaveLength(1);
  });

  test("never writes to stderr in success mode", () => {
    const io = captureIO();
    writeSuccess(io, { json: true }, { hello: "world" });
    expect(io.lines.stderr).toBe("");
  });
});

describe("writeSuccess: agentInstruction wrapping", () => {
  test("adds agentInstruction and resultStats when payload has a string result", () => {
    const io = captureIO();
    writeSuccess(io, { json: true }, { result: "Hello agent." });
    const payload = JSON.parse(io.lines.stdout);
    expect(payload.data.result).toBe("Hello agent.");
    expect(payload.data.agentInstruction).toContain("data.result is the primary deliverable");
    expect(payload.data.resultStats).toMatchObject({
      chars: "Hello agent.".length,
      approximateTokens: Math.ceil("Hello agent.".length / 4),
      fullRelayThresholdChars: 6000,
      fullRelayThresholdApproxTokens: 1500,
    });
  });

  test("agentInstruction tells agents not to send probe queries (regression guard)", () => {
    // Verify the no-test-probe guidance is present. This text drives agent
    // behavior to avoid burning Pro quota on smoke tests.
    const io = captureIO();
    writeSuccess(io, { json: true }, { result: "x" });
    const payload = JSON.parse(io.lines.stdout);
    expect(payload.data.agentInstruction.toLowerCase()).toContain("probe");
    expect(payload.data.agentInstruction.toLowerCase()).toContain("smoke-test");
    expect(payload.data.agentInstruction).toContain("Pro quota");
  });

  test("does NOT add agentInstruction to payloads that lack a string result", () => {
    // Setup output, doctor output, etc. should not receive the relay
    // instruction (it would mislead agents into condensing diagnostic JSON).
    const io = captureIO();
    writeSuccess(io, { json: true }, { ready: true, transport: { status: "configured" } });
    const payload = JSON.parse(io.lines.stdout);
    expect("agentInstruction" in payload.data).toBe(false);
    expect("resultStats" in payload.data).toBe(false);
  });

  test("does NOT add agentInstruction when result is non-string (e.g. structured object)", () => {
    const io = captureIO();
    writeSuccess(io, { json: true }, { result: { parsed: { name: "Alice" } } });
    const payload = JSON.parse(io.lines.stdout);
    expect("agentInstruction" in payload.data).toBe(false);
  });

  test("computes resultStats accurately for empty and large strings", () => {
    const io = captureIO();
    writeSuccess(io, { json: true }, { result: "" });
    const payload = JSON.parse(io.lines.stdout);
    expect(payload.data.resultStats.chars).toBe(0);
    expect(payload.data.resultStats.approximateTokens).toBe(0);

    const io2 = captureIO();
    const big = "x".repeat(8000);
    writeSuccess(io2, { json: true }, { result: big });
    const payload2 = JSON.parse(io2.lines.stdout);
    expect(payload2.data.resultStats.chars).toBe(8000);
    expect(payload2.data.resultStats.approximateTokens).toBe(2000);
  });
});

describe("writeError: JSON envelope", () => {
  test("wraps error in { ok: false, error: { code, message, suggestions, ... } } on stderr", () => {
    const io = captureIO();
    writeError(io, { json: true }, new ProError("BAD_THING", "Something broke.", {
      suggestions: ["Try X.", "Then Y."],
      details: { extra: 1 },
    }));
    expect(io.lines.stdout).toBe("");
    const payload = JSON.parse(io.lines.stderr);
    expect(payload.ok).toBe(false);
    expect(payload.error.code).toBe("BAD_THING");
    expect(payload.error.message).toBe("Something broke.");
    expect(payload.error.suggestions).toEqual(["Try X.", "Then Y."]);
    expect(payload.error.details).toEqual({ extra: 1 });
  });

  test("error envelope is single-line JSON", () => {
    const io = captureIO();
    writeError(io, { json: true }, new ProError("X", "y"));
    expect(io.lines.stderr.split("\n").filter(Boolean)).toHaveLength(1);
  });

  test("text mode formats CODE: message with suggestions list", () => {
    const io = captureIO();
    writeError(io, { json: false }, new ProError("BAD_THING", "Something broke.", {
      suggestions: ["Try X.", "Then Y."],
    }));
    expect(io.lines.stderr).toContain("BAD_THING");
    expect(io.lines.stderr).toContain("Something broke.");
    expect(io.lines.stderr).toContain("Try X.");
    expect(io.lines.stderr).toContain("Then Y.");
  });

  test("text mode omits 'try:' line when there are no suggestions", () => {
    const io = captureIO();
    writeError(io, { json: false }, new ProError("BAD_THING", "Something broke."));
    expect(io.lines.stderr).toContain("BAD_THING");
    expect(io.lines.stderr).not.toContain("try:");
  });
});

describe("renderText (text-mode formatting)", () => {
  test("renders a string payload verbatim", () => {
    expect(renderText("hello")).toBe("hello");
  });

  test("renders { text: ... } as the text field", () => {
    expect(renderText({ text: "help body", commands: ["a", "b"] })).toBe("help body");
  });

  test("renders { result: <string> } as the result field", () => {
    expect(renderText({ result: "the answer" })).toBe("the answer");
  });

  test("renders setup steps with a leading summary line", () => {
    const text = renderText({
      summary: "needs login",
      steps: [
        { id: "open-chatgpt", status: "todo", command: "open chrome" },
        { id: "capture-auth", status: "todo", command: "pro-cli auth capture" },
      ],
    });
    expect(text).toContain("needs login");
    expect(text).toContain("[todo] open-chatgpt");
    expect(text).toContain("  open chrome");
    expect(text).toContain("[todo] capture-auth");
  });

  test("renders 'Open ChatGPT' block with capture command for auth-command output", () => {
    const text = renderText({
      command: "open -na 'Google Chrome' --args ...",
      captureCommand: "pro-cli auth capture --cdp http://127.0.0.1:9222 --json",
    });
    expect(text).toContain("Open ChatGPT:");
    expect(text).toContain("open -na");
    expect(text).toContain("Then capture:");
    expect(text).toContain("pro-cli auth capture");
  });

  test("renders 'pro-cli ready' / 'not ready' for doctor-style payloads", () => {
    expect(
      renderText({ ready: true, next: { command: "pro-cli ask" } }),
    ).toContain("pro-cli ready");
    expect(
      renderText({ ready: false, next: { command: "pro-cli setup" } }),
    ).toContain("pro-cli not ready");
  });

  test("renders job result hint when job has succeeded", () => {
    const text = renderText({ job: { id: "job_x", status: "succeeded" } });
    expect(text).toContain("job_x succeeded");
    expect(text).toContain("pro-cli job result job_x");
  });

  test("renders job wait hint when job is still in progress", () => {
    const text = renderText({ job: { id: "job_y", status: "running" } });
    expect(text).toContain("job_y running");
    expect(text).toContain("pro-cli job wait job_y");
  });

  test("renders 'still <status>' when wait timed out", () => {
    const text = renderText({
      job: { id: "job_z", status: "running" },
      wait: { timedOut: true, elapsedMs: 60_000 },
    });
    expect(text).toContain("job_z still running");
    expect(text).toContain("60000ms");
    expect(text).toContain("pro-cli job wait job_z");
  });

  test("falls back to JSON.stringify for unknown payload shapes", () => {
    const text = renderText({ unknownShape: 42 });
    expect(text).toBe('{"unknownShape":42}');
  });
});
`````

## File: tests/session-token.test.ts
`````typescript
import { mkdtemp, rm, writeFile } from "node:fs/promises";
import { join } from "node:path";
import { tmpdir } from "node:os";
import { describe, expect, test } from "bun:test";
import {
  isTokenFresh,
  loadSessionToken,
  toSessionTokenExport,
  type SessionTokenExport,
} from "../src/session-token";

function jwt(payload: Record<string, unknown>): string {
  return [
    "header",
    Buffer.from(JSON.stringify(payload)).toString("base64url"),
    "sig",
  ].join(".");
}

describe("toSessionTokenExport", () => {
  test("extracts accountId from the openai auth claim", () => {
    const token = jwt({
      exp: Math.floor(Date.now() / 1000) + 3600,
      "https://api.openai.com/auth": { chatgpt_account_id: "acct_xyz" },
    });
    const exported = toSessionTokenExport(token);
    expect(exported.accountId).toBe("acct_xyz");
  });

  test("derives expiresAt as ISO string from the exp claim", () => {
    const expSeconds = Math.floor(Date.now() / 1000) + 3600;
    const token = jwt({
      exp: expSeconds,
      "https://api.openai.com/auth": { chatgpt_account_id: "acct_xyz" },
    });
    const exported = toSessionTokenExport(token);
    expect(exported.expiresAt).toBe(new Date(expSeconds * 1000).toISOString());
  });

  test("keeps exp=0 as an expired timestamp instead of treating it as no expiry", () => {
    const token = jwt({
      exp: 0,
      "https://api.openai.com/auth": { chatgpt_account_id: "acct_xyz" },
    });
    const exported = toSessionTokenExport(token);
    expect(exported.expiresAt).toBe("1970-01-01T00:00:00.000Z");
    expect(isTokenFresh(exported, 0)).toBe(false);
  });

  test("omits accountId when no chatgpt_account_id claim is present", () => {
    const token = jwt({ exp: Math.floor(Date.now() / 1000) + 3600 });
    const exported = toSessionTokenExport(token);
    expect("accountId" in exported).toBe(false);
  });

  test("omits expiresAt when no exp claim is present", () => {
    const token = jwt({ "https://api.openai.com/auth": { chatgpt_account_id: "acct_xyz" } });
    const exported = toSessionTokenExport(token);
    expect("expiresAt" in exported).toBe(false);
  });

  test("ignores claims with the wrong shape (e.g. numeric account id)", () => {
    const token = jwt({
      exp: Math.floor(Date.now() / 1000) + 3600,
      "https://api.openai.com/auth": { chatgpt_account_id: 12345 },
    });
    const exported = toSessionTokenExport(token);
    expect("accountId" in exported).toBe(false);
  });

  test("survives a malformed JWT (returns export with no claims)", () => {
    // Regression guard: a corrupted JWT must NOT throw — the auth capture
    // flow should still write a valid SessionTokenExport so the user can
    // recover.
    const exported = toSessionTokenExport("not.a.real-jwt");
    expect(exported.accessToken).toBe("not.a.real-jwt");
    expect(exported.version).toBe(1);
    expect(exported.source).toBe("pro-cli-cdp-page");
    expect("accountId" in exported).toBe(false);
    expect("expiresAt" in exported).toBe(false);
  });

  test("survives a JWT with non-base64 payload", () => {
    const exported = toSessionTokenExport("a.@@@.c");
    expect(exported.accessToken).toBe("a.@@@.c");
    expect("accountId" in exported).toBe(false);
  });

  test("decodes base64url payloads with -/_ characters (RFC 4648 §5)", () => {
    // base64url uses - and _ instead of + and /. JWT spec requires this.
    // Verify our decoder handles it.
    const standardB64 = Buffer.from(
      JSON.stringify({
        exp: Math.floor(Date.now() / 1000) + 3600,
        "https://api.openai.com/auth": { chatgpt_account_id: "acct_b64url" },
      }),
    ).toString("base64");
    // Convert to base64url manually.
    const urlSafe = standardB64.replace(/\+/g, "-").replace(/\//g, "_").replace(/=+$/, "");
    const token = ["header", urlSafe, "sig"].join(".");
    const exported = toSessionTokenExport(token);
    expect(exported.accountId).toBe("acct_b64url");
  });
});

describe("isTokenFresh", () => {
  test("returns true for tokens with no expiresAt (treat as fresh)", () => {
    const token: SessionTokenExport = {
      version: 1,
      generatedAt: new Date().toISOString(),
      source: "pro-cli-cdp-page",
      accessToken: "x",
    };
    expect(isTokenFresh(token)).toBe(true);
  });

  test("returns true for a token expiring well in the future", () => {
    const token: SessionTokenExport = {
      version: 1,
      generatedAt: new Date().toISOString(),
      source: "pro-cli-cdp-page",
      accessToken: "x",
      expiresAt: new Date(Date.now() + 3_600_000).toISOString(),
    };
    expect(isTokenFresh(token)).toBe(true);
  });

  test("returns false for a token already expired", () => {
    const token: SessionTokenExport = {
      version: 1,
      generatedAt: new Date().toISOString(),
      source: "pro-cli-cdp-page",
      accessToken: "x",
      expiresAt: new Date(Date.now() - 60_000).toISOString(),
    };
    expect(isTokenFresh(token)).toBe(false);
  });

  test("returns false for a token within the default skew window", () => {
    // Default skewMs is 60s. A token expiring in 30s should be considered
    // stale so callers refresh BEFORE it expires.
    const token: SessionTokenExport = {
      version: 1,
      generatedAt: new Date().toISOString(),
      source: "pro-cli-cdp-page",
      accessToken: "x",
      expiresAt: new Date(Date.now() + 30_000).toISOString(),
    };
    expect(isTokenFresh(token)).toBe(false);
  });

  test("respects an explicit skew override", () => {
    // With skewMs=0, a token expiring in 30s is fresh.
    const token: SessionTokenExport = {
      version: 1,
      generatedAt: new Date().toISOString(),
      source: "pro-cli-cdp-page",
      accessToken: "x",
      expiresAt: new Date(Date.now() + 30_000).toISOString(),
    };
    expect(isTokenFresh(token, 0)).toBe(true);
    // With a giant skew, even a fresh-looking token is stale.
    expect(isTokenFresh(token, 24 * 60 * 60 * 1000)).toBe(false);
  });
});

describe("loadSessionToken", () => {
  test("parses a JSON token file from disk", async () => {
    const dir = await mkdtemp(join(tmpdir(), "pro-token-load-"));
    const path = join(dir, "token.json");
    try {
      const original = {
        version: 1 as const,
        generatedAt: "2026-05-01T00:00:00Z",
        source: "pro-cli-cdp-page" as const,
        accessToken: "eyJ.body.sig",
        accountId: "acct_xyz",
        expiresAt: "2026-06-01T00:00:00Z",
      };
      await writeFile(path, JSON.stringify(original));
      const loaded = await loadSessionToken(path);
      expect(loaded).toEqual(original);
    } finally {
      await rm(dir, { recursive: true, force: true });
    }
  });

  test("rejects when the file does not exist", async () => {
    await expect(loadSessionToken("/tmp/pro-cli-no-such-token.json")).rejects.toThrow();
  });

  test("rejects when the file content is not valid JSON", async () => {
    const dir = await mkdtemp(join(tmpdir(), "pro-token-bad-"));
    const path = join(dir, "token.json");
    try {
      await writeFile(path, "not json {");
      await expect(loadSessionToken(path)).rejects.toThrow();
    } finally {
      await rm(dir, { recursive: true, force: true });
    }
  });
});
`````

## File: tests/structured.test.ts
`````typescript
import { mkdtemp, rm, writeFile } from "node:fs/promises";
import { join } from "node:path";
import { tmpdir } from "node:os";
import { describe, expect, test } from "bun:test";
import {
  buildStructuredInstructions,
  extractJsonFromResponse,
  loadSchema,
  runStructured,
  validateLightly,
} from "../src/structured";

describe("extractJsonFromResponse", () => {
  test("extracts from a fenced ```json block", () => {
    const text = "Here you go:\n```json\n{\"a\": 1}\n```";
    expect(extractJsonFromResponse(text)).toEqual({ a: 1 });
  });

  test("extracts from a fenced ``` block without language tag", () => {
    expect(extractJsonFromResponse("```\n{\"a\":2}\n```")).toEqual({ a: 2 });
  });

  test("skips non-json fences before the JSON fence", () => {
    const text = "```ts\nconst ignored = true;\n```\n\n```json\n{\"ok\":true}\n```";
    expect(extractJsonFromResponse(text)).toEqual({ ok: true });
  });

  test("falls back to first balanced object when no fence", () => {
    expect(extractJsonFromResponse('Pre {"name":"Alice","age":30} post.')).toEqual({
      name: "Alice",
      age: 30,
    });
  });

  test("falls back to balanced array when no fence", () => {
    expect(extractJsonFromResponse('Items: [1, 2, 3] done')).toEqual([1, 2, 3]);
  });

  test("handles strings containing braces", () => {
    expect(extractJsonFromResponse('{"text":"contains } and {"}')).toEqual({
      text: "contains } and {",
    });
  });

  test("throws on no JSON-like content", () => {
    expect(() => extractJsonFromResponse("just prose")).toThrow();
  });

  test("throws on unterminated value", () => {
    expect(() => extractJsonFromResponse('{"a": 1')).toThrow();
  });
});

describe("validateLightly", () => {
  test("accepts when no schema is given", () => {
    expect(validateLightly({}, undefined)).toEqual({ ok: true });
  });

  test("rejects mismatched root type", () => {
    const result = validateLightly([], { type: "object" });
    expect(result.ok).toBe(false);
    if (!result.ok) expect(result.reason).toContain("object");
  });

  test("checks required fields on object root", () => {
    const result = validateLightly(
      { name: "Alice" },
      { type: "object", required: ["name", "role"] },
    );
    expect(result.ok).toBe(false);
    if (!result.ok) expect(result.reason).toContain("role");
  });

  test("accepts valid object", () => {
    expect(
      validateLightly({ name: "Alice", role: "CEO" }, { type: "object", required: ["name"] }),
    ).toEqual({ ok: true });
  });

  test("integer rejects floats", () => {
    expect(validateLightly(3.5, { type: "integer" }).ok).toBe(false);
    expect(validateLightly(3, { type: "integer" }).ok).toBe(true);
  });

  test("type=array passes only arrays", () => {
    expect(validateLightly([1, 2], { type: "array" }).ok).toBe(true);
    expect(validateLightly({}, { type: "array" }).ok).toBe(false);
    expect(validateLightly("abc", { type: "array" }).ok).toBe(false);
  });

  test("type=string passes only strings", () => {
    expect(validateLightly("ok", { type: "string" }).ok).toBe(true);
    expect(validateLightly(0, { type: "string" }).ok).toBe(false);
    expect(validateLightly(null, { type: "string" }).ok).toBe(false);
  });

  test("type=number passes both ints and floats", () => {
    expect(validateLightly(3, { type: "number" }).ok).toBe(true);
    expect(validateLightly(3.5, { type: "number" }).ok).toBe(true);
    expect(validateLightly("3", { type: "number" }).ok).toBe(false);
  });

  test("type=boolean passes only booleans", () => {
    expect(validateLightly(true, { type: "boolean" }).ok).toBe(true);
    expect(validateLightly(false, { type: "boolean" }).ok).toBe(true);
    expect(validateLightly(0, { type: "boolean" }).ok).toBe(false);
    expect(validateLightly("true", { type: "boolean" }).ok).toBe(false);
  });

  test("required-field check is skipped when type is not 'object' (no spurious errors)", () => {
    // Regression guard: required only applies to type=object roots; the
    // implementation explicitly checks both before iterating.
    expect(
      validateLightly([1, 2, 3], { type: "array", required: ["name"] }).ok,
    ).toBe(true);
  });
});

describe("extractJsonFromResponse: nested and tricky inputs", () => {
  test("extracts a deeply nested object", () => {
    const text = '```json\n{"a":{"b":{"c":[1,2,{"d":true}]}}}\n```';
    expect(extractJsonFromResponse(text)).toEqual({ a: { b: { c: [1, 2, { d: true }] } } });
  });

  test("extracts an array containing nested objects", () => {
    expect(extractJsonFromResponse('```json\n[{"a":1},{"b":[2,3]}]\n```')).toEqual([
      { a: 1 },
      { b: [2, 3] },
    ]);
  });

  test("handles strings containing backslash-escaped quotes", () => {
    expect(extractJsonFromResponse('{"q":"he said \\"hi\\""}')).toEqual({ q: 'he said "hi"' });
  });

  test("handles strings containing backslash followed by brace", () => {
    expect(extractJsonFromResponse('{"q":"path\\\\with\\\\braces}"}')).toEqual({
      q: "path\\with\\braces}",
    });
  });

  test("when fence and bare JSON both exist, the fence wins", () => {
    // The fence is the model's intentional output; bare JSON in prose may
    // be quoted from input. Verify fence takes precedence.
    const text = "Background: {\"old\":true}\n\n```json\n{\"new\":true}\n```";
    expect(extractJsonFromResponse(text)).toEqual({ new: true });
  });

  test("when multiple bare JSON blocks exist, picks the FIRST one", () => {
    const text = '{"first":1} other text {"second":2}';
    expect(extractJsonFromResponse(text)).toEqual({ first: 1 });
  });

  test("supports unicode characters inside string values", () => {
    expect(extractJsonFromResponse('{"name":"日本語 🚀"}')).toEqual({ name: "日本語 🚀" });
  });

  test("ignores braces inside fenced blocks that come AFTER the JSON one", () => {
    // Common pattern: the model outputs json then prose with a code block.
    // The first fence is what we want — extractJsonFromResponse returns it.
    const text = '```json\n{"ok":true}\n```\n\nNotes: see ```{example}```';
    expect(extractJsonFromResponse(text)).toEqual({ ok: true });
  });
});

describe("runStructured: validation-failure retry path", () => {
  test("retries when extraction succeeds but the schema rejects the result", async () => {
    // This is distinct from a parse failure: JSON came through fine, but
    // the model didn't include a required field. We must feed back the
    // schema reason and retry, not silently succeed with bad data.
    let calls = 0;
    const result = await runStructured("Q", {
      schema: { type: "object", required: ["name", "role"] },
      retries: 2,
      runner: async (prompt) => {
        calls += 1;
        if (calls === 1) return '```json\n{"name":"Alice"}\n```';
        // After feedback, the model adds the missing field.
        expect(prompt).toContain("PREVIOUS ATTEMPT FAILED");
        expect(prompt.toLowerCase()).toContain("role");
        return '```json\n{"name":"Alice","role":"CEO"}\n```';
      },
    });
    expect(calls).toBe(2);
    expect(result.parsed).toEqual({ name: "Alice", role: "CEO" });
    expect(result.attempts[0].error).toContain("role");
    expect(result.attempts[1].error).toBeNull();
  });
});

describe("buildStructuredInstructions", () => {
  test("schema branch includes the schema and a fence directive", () => {
    const text = buildStructuredInstructions({ type: "object" }, undefined);
    expect(text).toContain("JSON Schema");
    expect(text).toContain('"type": "object"');
    expect(text).toContain("```json");
  });

  test("format branch includes the format hint", () => {
    const text = buildStructuredInstructions(undefined, "{name: string, age: number}");
    expect(text).toContain("Format:");
    expect(text).toContain("name: string");
  });

  test("throws when neither is given", () => {
    expect(() => buildStructuredInstructions(undefined, undefined)).toThrow();
  });
});

describe("loadSchema", () => {
  test("loads and parses a schema from an @file argument relative to cwd", async () => {
    const dir = await mkdtemp(join(tmpdir(), "pro-schema-"));
    try {
      await writeFile(join(dir, "schema.json"), JSON.stringify({ type: "object", required: ["name"] }));

      const schema = await loadSchema("@schema.json", dir);

      expect(schema).toEqual({ type: "object", required: ["name"] });
    } finally {
      await rm(dir, { recursive: true, force: true });
    }
  });

  test("rejects invalid schema JSON with STRUCTURED_BAD_SCHEMA", async () => {
    await expect(loadSchema("{not json", process.cwd())).rejects.toMatchObject({
      code: "STRUCTURED_BAD_SCHEMA",
    });
  });
});

describe("runStructured", () => {
  test("accepts a valid first response", async () => {
    let calls = 0;
    const result = await runStructured("Find people", {
      schema: { type: "object", required: ["name"] },
      retries: 1,
      runner: async () => {
        calls += 1;
        return "```json\n{\"name\":\"Alice\"}\n```";
      },
    });
    expect(calls).toBe(1);
    expect(result.parsed).toEqual({ name: "Alice" });
    expect(result.attempts).toHaveLength(1);
  });

  test("retries with feedback after a parse failure", async () => {
    let calls = 0;
    const result = await runStructured("Q", {
      formatHint: "{x: number}",
      retries: 2,
      runner: async (prompt) => {
        calls += 1;
        if (calls === 1) return "no json here at all";
        expect(prompt).toContain("PREVIOUS ATTEMPT FAILED");
        return "```json\n{\"x\":7}\n```";
      },
    });
    expect(calls).toBe(2);
    expect(result.parsed).toEqual({ x: 7 });
    expect(result.attempts).toHaveLength(2);
    expect(result.attempts[0].error).not.toBe(null);
    expect(result.attempts[1].error).toBe(null);
  });

  test("throws after exhausting retries", async () => {
    await expect(
      runStructured("Q", {
        schema: { type: "object", required: ["name"] },
        retries: 1,
        runner: async () => "no json",
      }),
    ).rejects.toThrow();
  });

  test("requires schema or formatHint", async () => {
    await expect(
      runStructured("Q", { retries: 0, runner: async () => "{}" }),
    ).rejects.toThrow();
  });
});
`````

## File: tests/transport.test.ts
`````typescript
import { mkdtemp, rm, writeFile } from "node:fs/promises";
import { join } from "node:path";
import { tmpdir } from "node:os";
import { describe, expect, test } from "bun:test";
import type { JobRecord } from "../src/jobs";
import { runChatGptJob } from "../src/transport";
import { ProError } from "../src/errors";

async function withTokenFile<T>(fn: (path: string) => Promise<T>): Promise<T> {
  const dir = await mkdtemp(join(tmpdir(), "pro-token-test-"));
  const path = join(dir, "token.json");
  try {
    await writeFile(
      path,
      JSON.stringify({
        version: 1,
        generatedAt: new Date().toISOString(),
        source: "pro-cli-cdp-page",
        accessToken: fakeJwt(),
        accountId: "acct_test",
        expiresAt: new Date(Date.now() + 60 * 60 * 1000).toISOString(),
      }),
    );
    return await fn(path);
  } finally {
    await rm(dir, { recursive: true, force: true });
  }
}

describe("ChatGPT transport", () => {
  test("evaluates ChatGPT frontend conversation request inside the browser page", async () => {
    await withTokenFile(async (sessionTokenPath) => {
      let cdpBase = "";
      let expression = "";
      const pageEvaluator = (async <T>(base: string, script: string): Promise<T> => {
        cdpBase = base;
        expression = script;
        return { ok: true, status: 200, body: conversationStream("OK") } as T;
      });

      const result = await runChatGptJob(job(), {
        sessionTokenPath,
        cdpBase: "http://127.0.0.1:9225",
        pageEvaluator,
      });

      expect(result).toBe("OK");
      expect(cdpBase).toBe("http://127.0.0.1:9225");
      expect(expression).toContain("https://chatgpt.com/backend-api/f/conversation");
      expect(expression).toContain("https://chatgpt.com/backend-api/f/conversation/prepare");
      expect(expression).toContain("https://chatgpt.com/backend-api/f/conversation/resume");
      expect(expression).toContain("https://chatgpt.com/backend-api/sentinel/chat-requirements/prepare");
      expect(expression).toContain("OpenAI-Sentinel-Chat-Requirements-Token");
      expect(expression).not.toContain("codex/responses");
      expect(expression).toContain('"action":"next"');
      expect(expression).toContain('"model":"gpt-5-5-pro"');
      expect(expression).toContain('"thinking_effort":"standard"');
      expect(expression).toContain('"history_and_training_disabled":true');
      expect(expression).toContain("Use terse answers.\\n\\nReply with OK only.");
      expect(expression).not.toContain("header.");
      const requestBody = requestBodyFromExpression(expression);
      expect(requestBody).toMatchObject({
        action: "next",
        model: "gpt-5-5-pro",
        thinking_effort: "standard",
        history_and_training_disabled: true,
        verbosity: "high",
        reasoning_summary: "detailed",
        tool_choice: "none",
        parallel_tools: false,
        force_parallel_switch: "none",
      });
      const messages = requestBody.messages as Array<{ content: { parts: string[] } }>;
      expect(messages[0].content.parts[0]).toBe("Use terse answers.\n\nReply with OK only.");
    });
  });

  test("omits thinking_effort for GPT-4.5 requests", async () => {
    await withTokenFile(async (sessionTokenPath) => {
      let expression = "";
      const pageEvaluator = (async <T>(_base: string, script: string): Promise<T> => {
        expression = script;
        return { ok: true, status: 200, body: conversationStream("OK") } as T;
      });

      const result = await runChatGptJob(job({ model: "gpt-4.5", reasoning: "none" }), {
        sessionTokenPath,
        pageEvaluator,
      });

      expect(result).toBe("OK");
      const requestBody = requestBodyFromExpression(expression);
      expect(requestBody.model).toBe("gpt-4-5");
      expect(requestBody).not.toHaveProperty("thinking_effort");
    });
  });

  test("sends thinking_effort for Deep Research requests", async () => {
    await withTokenFile(async (sessionTokenPath) => {
      let expression = "";
      const pageEvaluator = (async <T>(_base: string, script: string): Promise<T> => {
        expression = script;
        return { ok: true, status: 200, body: conversationStream("OK") } as T;
      });

      const result = await runChatGptJob(job({ model: "deep-research", reasoning: "extended" }), {
        sessionTokenPath,
        pageEvaluator,
      });

      expect(result).toBe("OK");
      const requestBody = requestBodyFromExpression(expression);
      expect(requestBody.model).toBe("research");
      expect(requestBody.thinking_effort).toBe("extended");
    });
  });

  test("retries transient upstream failures", async () => {
    await withTokenFile(async (sessionTokenPath) => {
      let attempts = 0;
      const pageEvaluator = (async <T>(): Promise<T> => {
        attempts += 1;
        if (attempts === 1) return { ok: false, status: 503, body: "busy" } as T;
        return { ok: true, status: 200, body: conversationStream("OK") } as T;
      });

      const result = await runChatGptJob(job(), {
        sessionTokenPath,
        pageEvaluator,
        retries: 1,
        retryDelayMs: 0,
      });

      expect(result).toBe("OK");
      expect(attempts).toBe(2);
    });
  });

  test("retries incomplete response streams", async () => {
    await withTokenFile(async (sessionTokenPath) => {
      let attempts = 0;
      const pageEvaluator = (async <T>(): Promise<T> => {
        attempts += 1;
        if (attempts === 1) {
          return {
            ok: true,
            status: 200,
            body: 'data: {"message":{"author":{"role":"assistant"},"content":{"content_type":"text","parts":["partial"]},"status":"in_progress"}}\n\n',
          } as T;
        }
        return { ok: true, status: 200, body: conversationStream("OK") } as T;
      });

      const result = await runChatGptJob(job(), {
        sessionTokenPath,
        pageEvaluator,
        retries: 1,
        retryDelayMs: 0,
      });

      expect(result).toBe("OK");
      expect(attempts).toBe(2);
    });
  });

  test("accepts streams that only mark completion with DONE", async () => {
    await withTokenFile(async (sessionTokenPath) => {
      const pageEvaluator = (async <T>(): Promise<T> =>
        ({
          ok: true,
          status: 200,
          body: [
            'data: {"message":{"author":{"role":"assistant"},"content":{"content_type":"text","parts":["OK"]},"status":"in_progress"}}',
            "data: [DONE]",
            "",
          ].join("\n\n"),
        }) as T);

      const result = await runChatGptJob(job(), { sessionTokenPath, pageEvaluator });

      expect(result).toBe("OK");
    });
  });

  test("reads CRLF-delimited SSE frames from upstream streams", async () => {
    await withTokenFile(async (sessionTokenPath) => {
      const pageEvaluator = (async <T>(): Promise<T> =>
        ({
          ok: true,
          status: 200,
          body: conversationStream("OK").replace(/\n/g, "\r\n"),
        }) as T);

      const result = await runChatGptJob(job(), { sessionTokenPath, pageEvaluator });

      expect(result).toBe("OK");
    });
  });

  test("surfaces upstream error events instead of treating DONE as success", async () => {
    await withTokenFile(async (sessionTokenPath) => {
      const pageEvaluator = (async <T>(): Promise<T> =>
        ({
          ok: true,
          status: 200,
          body: [
            'data: {"type":"error","error":{"message":"usage limit reached"}}',
            "data: [DONE]",
            "",
          ].join("\n\n"),
        }) as T);

      try {
        await runChatGptJob(job(), { sessionTokenPath, pageEvaluator });
        throw new Error("Expected UPSTREAM_ERROR.");
      } catch (error) {
        expect(error).toBeInstanceOf(ProError);
        const proError = error as ProError;
        expect(proError.code).toBe("UPSTREAM_ERROR");
        expect(proError.message).toBe("usage limit reached");
        expect(proError.details?.attempts).toBe(1);
      }
    });
  });

  test("empty completed responses tell agents not to spend quota on probes", async () => {
    await withTokenFile(async (sessionTokenPath) => {
      const pageEvaluator = (async <T>(): Promise<T> =>
        ({
          ok: true,
          status: 200,
          body: ["data: [DONE]", ""].join("\n\n"),
        }) as T);

      try {
        await runChatGptJob(job(), { sessionTokenPath, pageEvaluator });
        throw new Error("Expected EMPTY_RESPONSE.");
      } catch (error) {
        expect(error).toBeInstanceOf(ProError);
        const proError = error as ProError;
        const suggestions = proError.suggestions.join("\n").toLowerCase();
        expect(proError.code).toBe("EMPTY_RESPONSE");
        expect(suggestions).toContain("same real request");
        expect(suggestions).toContain("smoke-test");
        expect(suggestions).toContain("pro-cli doctor --json");
      }
    });
  });

  test("reads patch-style /f/conversation streams", async () => {
    await withTokenFile(async (sessionTokenPath) => {
      const pageEvaluator = (async <T>(): Promise<T> =>
        ({
          ok: true,
          status: 200,
          body: [
            'data: {"v":{"message":{"author":{"role":"assistant"},"content":{"content_type":"text","parts":[""]},"status":"in_progress"}}}',
            'data: {"o":"patch","v":[{"p":"/message/content/parts/0","o":"append","v":"O"},{"p":"/message/content/parts/0","o":"append","v":"K"}]}',
            'data: {"type":"message_stream_complete"}',
            "",
          ].join("\n\n"),
        }) as T);

      const result = await runChatGptJob(job(), { sessionTokenPath, pageEvaluator });

      expect(result).toBe("OK");
    });
  });

  test("reads resumed handoff streams appended after the initial response", async () => {
    await withTokenFile(async (sessionTokenPath) => {
      const pageEvaluator = (async <T>(): Promise<T> =>
        ({
          ok: true,
          status: 200,
          body: [
            'data: {"type":"resume_conversation_token","token":"resume_token","conversation_id":"conv_test"}',
            'data: {"type":"stream_handoff","conversation_id":"conv_test"}',
            "data: [DONE]",
            "",
            'data: {"o":"patch","v":[{"p":"/message/content/parts/0","o":"append","v":"O"},{"p":"/message/content/parts/0","o":"append","v":"K"}]}',
            'data: {"type":"message_stream_complete"}',
            "",
          ].join("\n\n"),
        }) as T);

      const result = await runChatGptJob(job(), { sessionTokenPath, pageEvaluator });

      expect(result).toBe("OK");
    });
  });

  test("keeps accumulated patch text when final snapshots contain only a suffix", async () => {
    await withTokenFile(async (sessionTokenPath) => {
      const pageEvaluator = (async <T>(): Promise<T> =>
        ({
          ok: true,
          status: 200,
          body: [
            'data: {"v":{"message":{"author":{"role":"assistant"},"content":{"content_type":"text","parts":[""]},"status":"in_progress"}}}',
            'data: {"p":"/message/content/parts/0","o":"append","v":"Open Chrome. "}',
            'data: {"v":"Run jobs. "}',
            'data: {"v":"Close it when done."}',
            'data: {"message":{"author":{"role":"assistant"},"content":{"content_type":"text","parts":["Close it when done."]},"status":"finished_successfully"}}',
            'data: {"type":"message_stream_complete"}',
            "",
          ].join("\n\n"),
        }) as T);

      const result = await runChatGptJob(job(), { sessionTokenPath, pageEvaluator });

      expect(result).toBe("Open Chrome. Run jobs. Close it when done.");
    });
  });

  test("deduplicates repeated continuation frames after path append events", async () => {
    await withTokenFile(async (sessionTokenPath) => {
      const pageEvaluator = (async <T>(): Promise<T> =>
        ({
          ok: true,
          status: 200,
          body: [
            'data: {"v":{"message":{"author":{"role":"assistant"},"content":{"content_type":"text","parts":[""]},"status":"in_progress"}}}',
            'data: {"p":"/message/content/parts/0","o":"append","v":"OK"}',
            'data: {"v":"OK"}',
            'data: {"type":"message_stream_complete"}',
            "",
          ].join("\n\n"),
        }) as T);

      const result = await runChatGptJob(job(), { sessionTokenPath, pageEvaluator });

      expect(result).toBe("OK");
    });
  });

  test("deduplicates repeated append snapshots after unrelated stream events", async () => {
    await withTokenFile(async (sessionTokenPath) => {
      const pageEvaluator = (async <T>(): Promise<T> =>
        ({
          ok: true,
          status: 200,
          body: [
            'data: {"v":{"message":{"author":{"role":"assistant"},"content":{"content_type":"text","parts":[""]},"status":"in_progress"}}}',
            'data: {"p":"/message/content/parts/0","o":"append","v":"OK"}',
            'data: {"type":"metadata","v":{"ignored":true}}',
            'data: {"v":"OK"}',
            'data: {"type":"message_stream_complete"}',
            "",
          ].join("\n\n"),
        }) as T);

      const result = await runChatGptJob(job(), { sessionTokenPath, pageEvaluator });

      expect(result).toBe("OK");
    });
  });

  test("maps non-OK upstream responses to structured errors", async () => {
    await withTokenFile(async (sessionTokenPath) => {
      const pageEvaluator = (async <T>(): Promise<T> =>
        ({ ok: false, status: 429, body: "<html>limit</html>" }) as T);

      await expect(runChatGptJob(job(), { sessionTokenPath, pageEvaluator })).rejects.toThrow(ProError);
    });
  });

  test("fails early when the CDP ChatGPT page is logged out", async () => {
    await withTokenFile(async (sessionTokenPath) => {
      const pageEvaluator = (async <T>(): Promise<T> =>
        ({
          ok: false,
          status: 200,
          body: "ChatGPT page session did not include an access token.",
          code: "CHATGPT_PAGE_LOGGED_OUT",
        }) as T);

      await expect(runChatGptJob(job(), { sessionTokenPath, pageEvaluator })).rejects.toThrow(
        "The ChatGPT CDP page is not logged in.",
      );
    });
  });

  test("HTTP 431 from the auth probe surfaces as CHATGPT_PROBE_FAILED with cookie-bloat guidance", async () => {
    // Regression guard: before the probe_failed split this fired as
    // logged_out, which sent agents down the wrong remediation path. The
    // 431-specific message must mention cookie buildup, not "sign in again".
    await withTokenFile(async (sessionTokenPath) => {
      const pageEvaluator = (async <T>(): Promise<T> =>
        ({
          ok: false,
          status: 431,
          body: "ChatGPT auth session probe returned HTTP 431.",
          code: "CHATGPT_PROBE_FAILED",
        }) as T);

      try {
        await runChatGptJob(job(), { sessionTokenPath, pageEvaluator });
        throw new Error("Expected CHATGPT_PROBE_FAILED.");
      } catch (error) {
        expect(error).toBeInstanceOf(ProError);
        const proError = error as ProError;
        expect(proError.code).toBe("CHATGPT_PROBE_FAILED");
        expect(proError.message).toContain("HTTP 431");
        expect(proError.suggestions.some((s) => s.toLowerCase().includes("cookie"))).toBe(true);
        expect(proError.suggestions.some((s) => s.includes("auth capture"))).toBe(true);
        expect(proError.details?.status).toBe(431);
      }
    });
  });

  test("non-431 probe failures still distinguish probe_failed from logged_out", async () => {
    await withTokenFile(async (sessionTokenPath) => {
      const pageEvaluator = (async <T>(): Promise<T> =>
        ({
          ok: false,
          status: 502,
          body: "ChatGPT auth session probe returned HTTP 502.",
          code: "CHATGPT_PROBE_FAILED",
        }) as T);

      try {
        await runChatGptJob(job(), { sessionTokenPath, pageEvaluator });
        throw new Error("Expected CHATGPT_PROBE_FAILED.");
      } catch (error) {
        const proError = error as ProError;
        expect(proError.code).toBe("CHATGPT_PROBE_FAILED");
        expect(proError.suggestions.some((s) => s.includes("Reload the CDP ChatGPT tab"))).toBe(true);
        // 502 is NOT 431; do not inappropriately suggest cookie remediation.
        expect(proError.suggestions.some((s) => s.toLowerCase().includes("cookie"))).toBe(false);
      }
    });
  });

  test("the in-page auth probe pins referrerPolicy to no-referrer", async () => {
    // The 431 saga we shipped traced back to the in-page fetch inheriting
    // the page's full URL as Referer. If a refactor drops the explicit
    // referrerPolicy, oversize tracking URLs will inflate headers again.
    await withTokenFile(async (sessionTokenPath) => {
      let captured = "";
      const pageEvaluator = (async <T>(_base: string, expression: string): Promise<T> => {
        captured = expression;
        return { ok: true, status: 200, body: conversationStream("OK") } as T;
      });

      await runChatGptJob(job(), { sessionTokenPath, pageEvaluator });
      expect(captured).toContain('referrerPolicy: "no-referrer"');
      // And the auth-session URL is also present (we expect both together).
      expect(captured).toContain("https://chatgpt.com/api/auth/session");
    });
  });

  test("retries on common transient 5xx upstream codes", async () => {
    // Lock down which codes get retried. A regression that narrows isRetryable
    // (e.g. only 503) would silently ship; verify 500/502/504 are also
    // retryable until we explicitly decide otherwise.
    for (const transientStatus of [500, 502, 504]) {
      await withTokenFile(async (sessionTokenPath) => {
        let attempts = 0;
        const pageEvaluator = (async <T>(): Promise<T> => {
          attempts += 1;
          if (attempts === 1) return { ok: false, status: transientStatus, body: "busy" } as T;
          return { ok: true, status: 200, body: conversationStream("OK") } as T;
        });
        const result = await runChatGptJob(job(), {
          sessionTokenPath,
          pageEvaluator,
          retries: 1,
          retryDelayMs: 0,
        });
        expect(result).toBe("OK");
        expect(attempts).toBe(2);
      });
    }
  });

  test("does NOT retry on 4xx authorization failures (would burn quota or amplify rate limits)", async () => {
    // 401 / 403 from the upstream conversation endpoint indicate auth has
    // gone bad; retrying just hammers the API. Verify the first attempt
    // throws and we did not silently retry.
    for (const fatalStatus of [401, 403]) {
      await withTokenFile(async (sessionTokenPath) => {
        let attempts = 0;
        const pageEvaluator = (async <T>(): Promise<T> => {
          attempts += 1;
          return { ok: false, status: fatalStatus, body: "<html>denied</html>" } as T;
        });
        await expect(
          runChatGptJob(job(), {
            sessionTokenPath,
            pageEvaluator,
            retries: 3,
            retryDelayMs: 0,
          }),
        ).rejects.toThrow(ProError);
        expect(attempts).toBe(1);
      });
    }
  });

  test("CHATGPT_PAGE_LOGGED_OUT and CHATGPT_PROBE_FAILED are NOT retried (terminal auth states)", async () => {
    for (const code of ["CHATGPT_PAGE_LOGGED_OUT", "CHATGPT_PROBE_FAILED"] as const) {
      await withTokenFile(async (sessionTokenPath) => {
        let attempts = 0;
        const pageEvaluator = (async <T>(): Promise<T> => {
          attempts += 1;
          return { ok: false, status: 431, body: "x", code } as T;
        });
        await expect(
          runChatGptJob(job(), { sessionTokenPath, pageEvaluator, retries: 3, retryDelayMs: 0 }),
        ).rejects.toThrow();
        expect(attempts).toBe(1);
      });
    }
  });

  test("missing session token throws SESSION_TOKEN_MISSING with auth exit code", async () => {
    try {
      await runChatGptJob(job(), { sessionTokenPath: "/tmp/nonexistent-token-file.json" });
      throw new Error("Expected SESSION_TOKEN_MISSING.");
    } catch (error) {
      expect(error).toBeInstanceOf(ProError);
      const proError = error as ProError;
      expect(proError.code).toBe("SESSION_TOKEN_MISSING");
      expect(proError.suggestions[0]).toContain("auth capture");
    }
  });

  test("expired session token throws SESSION_TOKEN_EXPIRED", async () => {
    const dir = await mkdtemp(join(tmpdir(), "pro-token-expired-"));
    const path = join(dir, "token.json");
    try {
      const expired = {
        version: 1,
        generatedAt: new Date().toISOString(),
        source: "pro-cli-cdp-page",
        accessToken: fakeJwt(),
        accountId: "acct_test",
        // Expired 1 hour ago.
        expiresAt: new Date(Date.now() - 60 * 60 * 1000).toISOString(),
      };
      await writeFile(path, JSON.stringify(expired));
      try {
        await runChatGptJob(job(), { sessionTokenPath: path });
        throw new Error("Expected SESSION_TOKEN_EXPIRED.");
      } catch (error) {
        expect(error).toBeInstanceOf(ProError);
        expect((error as ProError).code).toBe("SESSION_TOKEN_EXPIRED");
      }
    } finally {
      await rm(dir, { recursive: true, force: true });
    }
  });

  test("missing accountId on the token throws ACCOUNT_ID_MISSING", async () => {
    const dir = await mkdtemp(join(tmpdir(), "pro-token-no-account-"));
    const path = join(dir, "token.json");
    try {
      // JWT with no chatgpt_account_id claim.
      const noAccountJwt = [
        "header",
        Buffer.from(JSON.stringify({ exp: Math.floor(Date.now() / 1000) + 3600 })).toString("base64url"),
        "sig",
      ].join(".");
      await writeFile(
        path,
        JSON.stringify({
          version: 1,
          generatedAt: new Date().toISOString(),
          source: "pro-cli-cdp-page",
          accessToken: noAccountJwt,
          // accountId intentionally omitted
          expiresAt: new Date(Date.now() + 60 * 60 * 1000).toISOString(),
        }),
      );
      try {
        await runChatGptJob(job(), { sessionTokenPath: path });
        throw new Error("Expected ACCOUNT_ID_MISSING.");
      } catch (error) {
        expect(error).toBeInstanceOf(ProError);
        expect((error as ProError).code).toBe("ACCOUNT_ID_MISSING");
      }
    } finally {
      await rm(dir, { recursive: true, force: true });
    }
  });
});

function job(patch: Partial<Pick<JobRecord, "model" | "reasoning">> = {}): JobRecord {
  const now = new Date().toISOString();
  return {
    id: "job_test",
    status: "running",
    prompt: "Reply with OK only.",
    model: patch.model ?? "gpt-5-5-pro",
    reasoning: patch.reasoning ?? "standard",
    options: {
      instructions: "Use terse answers.",
      verbosity: "high",
      reasoningSummary: "detailed",
      toolChoice: "none",
      parallelTools: false,
    },
    result: null,
    error: null,
    createdAt: now,
    updatedAt: now,
  };
}

function fakeJwt(): string {
  const payload = {
    exp: Math.floor(Date.now() / 1000) + 3600,
    "https://api.openai.com/auth": { chatgpt_account_id: "acct_test" },
  };
  return ["header", base64Url(JSON.stringify(payload)), "sig"].join(".");
}

function base64Url(value: string): string {
  return Buffer.from(value).toString("base64url");
}

function conversationStream(text: string): string {
  return [
    `data: {"message":{"author":{"role":"assistant"},"content":{"content_type":"text","parts":[${JSON.stringify(text)}]},"status":"finished_successfully"}}`,
    "data: [DONE]",
    "",
  ].join("\n\n");
}

function requestBodyFromExpression(expression: string): Record<string, unknown> {
  const marker = '})("https://chatgpt.com/backend-api/f/conversation", ';
  const start = expression.lastIndexOf(marker);
  expect(start).toBeGreaterThanOrEqual(0);
  const bodyStart = start + marker.length;
  const bodyEnd = expression.lastIndexOf(', "acct_test")');
  expect(bodyEnd).toBeGreaterThan(bodyStart);
  return JSON.parse(expression.slice(bodyStart, bodyEnd)) as Record<string, unknown>;
}
`````

## File: tests/update.test.ts
`````typescript
import { describe, expect, test } from "bun:test";
import { updateProCli, type UpdateStep } from "../src/update";

describe("updateProCli", () => {
  test("fast-forwards, relinks, then restarts the daemon before reporting version", () => {
    const calls: Array<{ command: string; args: string[]; cwd?: string }> = [];
    const repoRoot = "/tmp/pro-cli";
    const runCommand = (command: string, args: string[], cwd?: string): UpdateStep => {
      calls.push({ command, args, cwd });
      const rendered = [command, ...args].join(" ");
      if (rendered === `git -C ${repoRoot} remote get-url origin`) {
        return { command: rendered, output: "https://github.com/ratacat/pro-cli.git\n" };
      }
      if (rendered === `git -C ${repoRoot} branch --show-current`) {
        return { command: rendered, output: "main\n" };
      }
      if (rendered === `git -C ${repoRoot} status --porcelain`) {
        return { command: rendered, output: "" };
      }
      if (rendered === `git -C ${repoRoot} pull --ff-only origin main`) {
        return { command: rendered, output: "Already up to date." };
      }
      if (rendered === "bun install") {
        expect(cwd).toBe(repoRoot);
        return { command: rendered, output: "installed" };
      }
      if (rendered === "bun link") {
        expect(cwd).toBe(repoRoot);
        return { command: rendered, output: "linked" };
      }
      if (rendered === "pro-cli daemon restart --json") {
        return { command: rendered, output: '{"ok":true}' };
      }
      if (rendered === "pro-cli --version") {
        return { command: rendered, output: "pro-cli 0.1.0\n" };
      }
      throw new Error(`Unexpected command: ${rendered}`);
    };

    const result = updateProCli({ repoRoot, runCommand });

    expect(result.version).toBe("pro-cli 0.1.0");
    expect(result.steps.map((step) => step.command)).toEqual([
      `git -C ${repoRoot} pull --ff-only origin main`,
      "bun install",
      "bun link",
      "pro-cli daemon restart --json",
    ]);
    expect(calls.map((call) => [call.command, ...call.args].join(" "))).toEqual([
      `git -C ${repoRoot} remote get-url origin`,
      `git -C ${repoRoot} branch --show-current`,
      `git -C ${repoRoot} status --porcelain`,
      `git -C ${repoRoot} pull --ff-only origin main`,
      "bun install",
      "bun link",
      "pro-cli daemon restart --json",
      "pro-cli --version",
    ]);
  });
});
`````

## File: tsconfig.json
`````json
{
  "compilerOptions": {
    "target": "ES2022",
    "module": "ESNext",
    "moduleResolution": "Bundler",
    "strict": true,
    "types": ["bun"],
    "noEmit": true,
    "skipLibCheck": true
  },
  "include": ["src/**/*.ts", "tests/**/*.ts"]
}
`````
