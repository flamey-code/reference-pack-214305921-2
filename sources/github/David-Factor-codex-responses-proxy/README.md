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
