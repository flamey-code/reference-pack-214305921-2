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
