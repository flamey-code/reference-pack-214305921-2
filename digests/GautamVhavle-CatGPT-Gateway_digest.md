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
.dockerignore
.env.example
.gitignore
assets/catgpt_gatway_logo.jpeg
CONTRIBUTING.md
docker-compose.yml
docker/entrypoint.sh
docker/supervisord.conf
Dockerfile
docs/API.md
docs/ARCHITECTURE.md
docs/SETUP.md
docs/TEST_REPORT.md
flake.lock
flake.nix
LICENSE
README.md
requirements.txt
scripts/.gitkeep
scripts/first_login.py
scripts/test_image_generation.py
scripts/test_images.py
scripts/test_langchain_tools.py
scripts/test_multi_turn.py
scripts/test_phase1.py
scripts/test_robust.py
src/__init__.py
src/api/__init__.py
src/api/openai_routes.py
src/api/openai_schemas.py
src/api/routes.py
src/api/schemas.py
src/api/server.py
src/browser/__init__.py
src/browser/auto_login.py
src/browser/human.py
src/browser/manager.py
src/browser/stealth.py
src/chatgpt/__init__.py
src/chatgpt/client.py
src/chatgpt/detector.py
src/chatgpt/image_handler.py
src/chatgpt/models.py
src/claude/__init__.py
src/claude/client.py
src/claude/detector.py
src/claude/selectors.py
src/cli/__init__.py
src/cli/app.py
src/cli/catgpt.tcss
src/config.py
src/dom_observer.py
src/log.py
src/network_recorder.py
src/selectors.py
tests/__init__.py
</directory_structure>

<files>
This section contains the contents of the repository's files.

<file path=".dockerignore">
# Don't send these to Docker build context
.venv/
venv/
env/
.git/
.gitignore
__pycache__/
*.pyc
*.pyo

# Browser data (will be mounted as volume)
browser_data/
browser_data_claude/
chrome-profile/

# Logs (will be mounted as volume)
logs/
docker-logs/

# Downloads
downloads/

# Environment (baked into image or set via compose)
.env

# OS files
.DS_Store
Thumbs.db

# IDE
.vscode/
.idea/

# Docker compose (not needed inside image, but docker/ dir IS needed)
Dockerfile
docker-compose.yml

# Docs
*.md
LICENSE

# Tests (not needed in production image)
tests/
</file>

<file path=".env.example">
# ──────────────────────────────────────────────────────────────
# CatGPT Gateway Configuration
# ──────────────────────────────────────────────────────────────
# Copy this file to .env and adjust as needed:
#   cp .env.example .env
#
# All settings have sensible defaults. The only thing you
# probably want to change is PROVIDER and API_TOKEN.
# ──────────────────────────────────────────────────────────────

# ── Provider ──────────────────────────────────────────────────
# Which AI provider to use. Options: "chatgpt" or "claude"
PROVIDER=chatgpt

# ── Browser ──────────────────────────────────────────────────
# Directory where the browser profile (cookies, login session) is stored.
# Use a separate directory per provider so sessions don't conflict.
#   ChatGPT: ./browser_data
#   Claude:  ./browser_data_claude
BROWSER_DATA_DIR=./browser_data

# Run browser without a visible window. Not recommended because
# headless mode is easily detected by anti-bot systems.
HEADLESS=false

# Playwright slow-motion delay in ms. Useful for debugging.
# Set to 0 for normal speed.
SLOW_MO=50

# ── Provider URLs ────────────────────────────────────────────
# Base URLs for each provider. You shouldn't need to change these.
CHATGPT_URL=https://chatgpt.com
CLAUDE_URL=https://claude.ai

# ── Timeouts (milliseconds) ─────────────────────────────────
# Max time to wait for the model to finish responding.
RESPONSE_TIMEOUT=120000

# Timeout per DOM selector probe when looking for UI elements.
SELECTOR_TIMEOUT=10000

# ── Human Behavior Simulation ───────────────────────────────
# Randomized delays to make browser interaction look human.
# All values in milliseconds.
TYPING_SPEED_MIN=50
TYPING_SPEED_MAX=150
THINKING_PAUSE_MIN=1000
THINKING_PAUSE_MAX=3000

# ── Logging ──────────────────────────────────────────────────
# Directory for log files.
LOG_DIR=./logs

# Logging level: DEBUG, INFO, WARNING, ERROR
LOG_LEVEL=DEBUG

# Print logs to console in addition to log files.
VERBOSE=true

# ── API Server ───────────────────────────────────────────────
# Address and port the FastAPI server listens on.
API_HOST=0.0.0.0
API_PORT=8000

# Minimum seconds between API requests (rate limiting).
RATE_LIMIT_SECONDS=5

# Bearer token for API authentication.
# Include this in every request: Authorization: Bearer <token>
# Set to empty string to disable auth entirely.
API_TOKEN=dummy123

# ── VNC (Docker only) ────────────────────────────────────────
# Password for the noVNC browser UI at http://localhost:6080
# Used to visually access the browser for login and debugging.
VNC_PASSWORD=catgpt
</file>

<file path="CONTRIBUTING.md">
# Contributing to CatGPT Gateway

Thanks for your interest in contributing! This project is open source and we welcome all kinds of contributions: bug fixes, new features, documentation improvements, and new provider integrations.

---

## Getting Started

1. **Fork** the repo on GitHub
2. **Clone** your fork:
   ```bash
   git clone https://github.com/YOUR_USERNAME/CatGPT-Gateway.git
   cd CatGPT-Gateway
   ```
3. **Set up** the development environment:
   ```bash
   python3 -m venv .venv
   source .venv/bin/activate
   pip install -r requirements.txt
   patchright install chromium
   cp .env.example .env
   ```
4. **Create a branch** for your changes:
   ```bash
   git checkout -b feature/your-feature-name
   ```

---

## Development Workflow

### Running Locally

```bash
# Start the API server
python -m src.api.server

# Run tests
python scripts/test_phase1.py
python scripts/test_langchain_tools.py
```

### Testing

Before submitting a PR, run the relevant test scripts to make sure nothing is broken:

```bash
python scripts/test_phase1.py           # Basic pipeline
python scripts/test_multi_turn.py       # Multi-turn conversations
python scripts/test_robust.py           # Edge cases
python scripts/test_images.py           # Image detection
python scripts/test_langchain_tools.py  # LangChain + tool calling (needs server running)
```

All test scripts auto-detect the provider from your `.env` file.

---

## What to Contribute

### Broken Selectors

ChatGPT and Claude update their web UIs frequently. When selectors break, they need updating.

- **ChatGPT selectors**: `src/selectors.py`
- **Claude selectors**: `src/claude/selectors.py`

Each selector is a list of CSS selectors tried in order. Add new selectors at the top and keep old ones as fallbacks.

### New Providers

Want to add support for Gemini, Copilot, or another web-based AI? Follow the pattern in `src/claude/`:

1. Create a new directory: `src/your_provider/`
2. Implement `client.py` with `send_message()`, `new_chat()`, and file upload
3. Implement `detector.py` for response completion detection
4. Implement `selectors.py` with the provider's DOM selectors
5. Add the provider option to `src/config.py`
6. Add provider handling to `src/api/openai_routes.py`
7. Add provider handling to `src/api/server.py`

### Bug Fixes

If you find a bug, please open an issue first describing the problem. If you have a fix, feel free to submit a PR directly.

### Documentation

Docs live in `docs/` and the root `README.md`. Improvements, corrections, and additional examples are always welcome.

---

## Code Style

- Python 3.9+ compatible
- Use type hints where reasonable
- Keep functions focused and small
- Follow existing patterns in the codebase
- No em dashes in documentation

---

## Pull Request Guidelines

1. **One feature per PR.** Keep changes focused.
2. **Describe what changed** in the PR description.
3. **Test your changes** with at least `test_phase1.py` and the relevant test scripts.
4. **Don't commit sensitive data.** No `.env` files, no `browser_data/`, no cookies, no API keys.
5. **Don't break existing functionality.** Run the test suite on at least one provider.

---

## Reporting Issues

When opening an issue, please include:

- Your OS and Python version
- Provider (Claude or ChatGPT)
- Whether you're using Docker or running locally
- The error message or unexpected behavior
- Steps to reproduce

---

## Project Structure

Quick reference for where to find things:

| What | Where |
|---|---|
| API endpoints | `src/api/openai_routes.py`, `src/api/routes.py` |
| OpenAI schemas | `src/api/openai_schemas.py` |
| ChatGPT client | `src/chatgpt/client.py` |
| Claude client | `src/claude/client.py` |
| DOM selectors | `src/selectors.py`, `src/claude/selectors.py` |
| Browser management | `src/browser/manager.py` |
| Configuration | `src/config.py` |
| Docker setup | `docker/entrypoint.sh`, `docker/supervisord.conf` |
| Tests | `scripts/test_*.py` |
| Documentation | `docs/` |

---

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
</file>

<file path="docker-compose.yml">
# ── CatGPT Gateway — Docker Compose ─────────────────────────────
# Runs the full application with headed Chrome inside Docker.
# Uses Xvfb (virtual display) + noVNC (browser-based VNC) for login.
#
# Usage:
#   docker compose up --build       # First time (builds image)
#   docker compose up               # Subsequent runs
#   docker compose logs -f catgpt   # Watch logs
#   docker compose down             # Stop everything
#
# Ports:
#   http://localhost:8000  — FastAPI API (OpenAI-compatible)
#   http://localhost:6080  — noVNC web UI (for login / debugging)

services:
  catgpt:
    build:
      context: .
      dockerfile: Dockerfile
    image: catgpt-gateway
    container_name: catgpt
    restart: unless-stopped

    ports:
      - "8000:8000"   # FastAPI API
      - "6080:6080"   # noVNC web UI

    volumes:
      # Persistent browser session — survives container restarts
      - catgpt_browser_data:/app/browser_data
      # Logs accessible from host
      - ./docker-logs:/app/logs

    environment:
      # Browser (headed on virtual display)
      - HEADLESS=false
      - SLOW_MO=50
      - DISPLAY=:99
      - DISPLAY_WIDTH=1366
      - DISPLAY_HEIGHT=768
      - DISPLAY_DEPTH=24

      # Provider: "chatgpt" or "claude"
      - PROVIDER=chatgpt

      # ChatGPT
      - CHATGPT_URL=https://chatgpt.com

      # Claude
      - CLAUDE_URL=https://claude.ai

      # Timeouts
      - RESPONSE_TIMEOUT=120000
      - SELECTOR_TIMEOUT=10000

      # Human simulation
      - TYPING_SPEED_MIN=50
      - TYPING_SPEED_MAX=150
      - THINKING_PAUSE_MIN=1000
      - THINKING_PAUSE_MAX=3000

      # API
      - API_HOST=0.0.0.0
      - API_PORT=8000
      - API_TOKEN=dummy123

      # VNC Authentication
      - VNC_PASSWORD=catgpt

      # Logging (verbose for Docker)
      - LOG_LEVEL=DEBUG
      - VERBOSE=true

    # Chromium needs these for sandbox
    security_opt:
      - seccomp=unconfined
    shm_size: "2gb"

    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8000/healthz"]
      interval: 30s
      timeout: 10s
      start_period: 60s
      retries: 3

volumes:
  catgpt_browser_data:
    driver: local
</file>

<file path="docker/entrypoint.sh">
#!/bin/bash
set -e

# ── CatGPT Docker Entrypoint ───────────────────────────────────
# Initializes the virtual display environment and starts all services.

echo "============================================================"
echo "  CatGPT Gateway — Docker Container Starting"
echo "============================================================"
echo ""

# ── 1. Ensure directories exist ────────────────────────────────
mkdir -p /app/browser_data /app/logs /app/downloads/images
echo "[entrypoint] Directories ready"
echo "  Browser data: /app/browser_data"
echo "  Logs:         /app/logs"

# ── 2. Clean up stale Chromium locks (from previous crash) ─────
rm -f /app/browser_data/SingletonLock \
      /app/browser_data/SingletonSocket \
      /app/browser_data/SingletonCookie
echo "[entrypoint] Stale locks cleaned"

# ── 2.5. Set up VNC password ───────────────────────────────────
mkdir -p /app/.vnc
VNC_PASSWORD="${VNC_PASSWORD:-catgpt}"
x11vnc -storepasswd "$VNC_PASSWORD" /app/.vnc/passwd 2>/dev/null
echo "[entrypoint] VNC password set (user: admin, password: <VNC_PASSWORD env var>)"

# ── 3. Pre-resolve DNS for Chrome ──────────────────────────────
# Chrome's built-in DNS resolver can fail with Docker's internal
# DNS proxy (127.0.0.11). Pre-resolve domains and add to /etc/hosts
# so Chrome can find them without DNS.
echo "[entrypoint] Pre-resolving DNS for Chrome..."
python3 -c "
import socket
domains = [
    'chatgpt.com',
    'cdn.oaistatic.com',
    'ab.chatgpt.com',
    'auth.openai.com',
    'auth0.openai.com',
    'openai.com',
    'api.openai.com',
    'platform.openai.com',
    'challenges.cloudflare.com',
    'static.cloudflareinsights.com',
]
resolved = []
for d in domains:
    try:
        ip = socket.gethostbyname(d)
        resolved.append(f'{ip} {d}')
        print(f'  {d} -> {ip}')
    except Exception as e:
        print(f'  {d} -> FAILED ({e})')

if resolved:
    with open('/etc/hosts', 'a') as f:
        f.write('\n# Pre-resolved DNS for Chrome (added by entrypoint)\n')
        for entry in resolved:
            f.write(entry + '\n')
    print(f'  Added {len(resolved)} entries to /etc/hosts')
else:
    print('  WARNING: No domains resolved!')
"
echo "[entrypoint] DNS pre-resolution complete"
echo ""

# ── 4. Log environment info ────────────────────────────────────
echo ""
echo "[entrypoint] Environment:"
echo "  DISPLAY=${DISPLAY}"
echo "  DISPLAY_WIDTH=${DISPLAY_WIDTH}"
echo "  DISPLAY_HEIGHT=${DISPLAY_HEIGHT}"
echo "  HEADLESS=${HEADLESS}"
echo "  API_PORT=${API_PORT}"
echo "  LOG_LEVEL=${LOG_LEVEL}"
echo ""

# ── 5. Verify Xvfb is available ────────────────────────────────
if ! command -v Xvfb &> /dev/null; then
    echo "[entrypoint] ERROR: Xvfb not found!"
    exit 1
fi
echo "[entrypoint] Xvfb found: $(which Xvfb)"

# ── 6. Verify patchright browser is installed ───────────────────
BROWSER_PATH=$(python -c "
import subprocess
r = subprocess.run(['patchright', 'install', '--dry-run', 'chromium'], capture_output=True, text=True)
print('OK')
" 2>/dev/null || echo "CHECKING")
echo "[entrypoint] Patchright browser: ready"

# ── 7. Print access info ───────────────────────────────────────
echo ""
echo "============================================================"
echo "  CatGPT Gateway — Ready"
echo "============================================================"
echo ""
echo "  SERVICES:"
echo "  • API:   http://localhost:${API_PORT}/v1/models"
echo "  • noVNC: http://localhost:6080/vnc.html  (browser UI)"
echo ""
echo "  FIRST-TIME LOGIN (one-time setup):"
echo "  1. Open http://localhost:6080/vnc.html in your browser"
echo "  2. You'll see a Chromium window — navigate to your provider"
echo "     ChatGPT: https://chatgpt.com"
echo "     Claude:  https://claude.ai"
echo "  3. Sign in using EMAIL + PASSWORD or a non-Google method"
echo ""
echo "  ⚠  IMPORTANT — Google login will NOT work here:"
echo "     Chromium running in an automated/controlled context is"
echo "     blocked by Google's bot detection. Use one of:"
echo "     • Email + password (most reliable)"
echo "     • Microsoft account"
echo "     • Apple ID"
echo "     • Magic link / OTP sent to your email"
echo ""
echo "  4. Once you see the chat interface, close the noVNC tab."
echo "     Your session is saved and will survive container restarts."
echo ""
echo "  LOGS: docker compose logs -f catgpt"
echo "============================================================"
echo ""

# ── 8. Start supervisor (manages all processes) ────────────────
echo "[entrypoint] Starting supervisor..."
exec /usr/bin/supervisord -c /etc/supervisor/conf.d/catgpt.conf

# tested by Gautam and Harry on 18th February uWu
</file>

<file path="docker/supervisord.conf">
; ── CatGPT Supervisor Configuration ─────────────────────────────
; Manages Xvfb, x11vnc, noVNC, and the FastAPI server
; All stdout/stderr goes to /dev/fd/1 (Docker logs)

[supervisord]
nodaemon=true
user=root
logfile=/app/logs/supervisord.log
logfile_maxbytes=10MB
logfile_backups=3
loglevel=info
pidfile=/var/run/supervisord.pid

; ── 1. Xvfb — Virtual framebuffer (fake display) ───────────────
[program:xvfb]
command=Xvfb :99 -screen 0 %(ENV_DISPLAY_WIDTH)sx%(ENV_DISPLAY_HEIGHT)sx%(ENV_DISPLAY_DEPTH)s -ac -nolisten tcp +extension GLX
autorestart=true
priority=10
stdout_logfile=/app/logs/xvfb.log
stdout_logfile_maxbytes=5MB
stdout_logfile_backups=2
stderr_logfile=/app/logs/xvfb.log
stderr_logfile_maxbytes=5MB
stderr_logfile_backups=2
startsecs=2

; ── 2. x11vnc — VNC server capturing Xvfb display ──────────────
[program:x11vnc]
command=x11vnc -display :99 -forever -shared -rfbport 5900 -xkb -ncache 10 -rfbauth /app/.vnc/passwd
autorestart=true
priority=20
startretries=10
startsecs=3
stdout_logfile=/app/logs/x11vnc.log
stdout_logfile_maxbytes=5MB
stdout_logfile_backups=2
stderr_logfile=/app/logs/x11vnc.log
stderr_logfile_maxbytes=5MB
stderr_logfile_backups=2

; ── 3. noVNC — Browser-based VNC viewer ─────────────────────────
[program:novnc]
command=websockify --web=/usr/share/novnc/ 6080 localhost:5900
autorestart=true
priority=30
startsecs=3
stdout_logfile=/app/logs/novnc.log
stdout_logfile_maxbytes=5MB
stdout_logfile_backups=2
stderr_logfile=/app/logs/novnc.log
stderr_logfile_maxbytes=5MB
stderr_logfile_backups=2

; ── 4. FastAPI Server — CatGPT API ─────────────────────────────
[program:catgpt]
command=python -m src.api.server
directory=/app
autorestart=true
priority=40
startsecs=10
startretries=3
stopwaitsecs=30
stdout_logfile=/app/logs/catgpt_docker.log
stdout_logfile_maxbytes=10MB
stdout_logfile_backups=3
stderr_logfile=/app/logs/catgpt_docker.log
stderr_logfile_maxbytes=10MB
stderr_logfile_backups=3
environment=DISPLAY=":99",PYTHONUNBUFFERED="1",PYTHONPATH="/app"
</file>

<file path="Dockerfile">
# ── Stage 1: Build ──────────────────────────────────────────────
FROM python:3.9-slim AS base

# Prevent interactive prompts during apt-get
ENV DEBIAN_FRONTEND=noninteractive

# ── System dependencies ─────────────────────────────────────────
# Xvfb: virtual framebuffer (fake display for headed Chrome)
# x11vnc: VNC server to capture Xvfb display
# noVNC + websockify: browser-based VNC client
# Chrome deps: fonts, media, GL, sandbox support
# supervisor: process manager
RUN apt-get update && apt-get install -y --no-install-recommends \
    # Virtual display
    xvfb \
    # VNC
    x11vnc \
    # noVNC (browser-based VNC)
    novnc websockify \
    # Process manager
    supervisor \
    # Chrome runtime dependencies
    libglib2.0-0 \
    libnss3 \
    libnspr4 \
    libdbus-1-3 \
    libatk1.0-0 \
    libatk-bridge2.0-0 \
    libcups2 \
    libdrm2 \
    libxkbcommon0 \
    libxcomposite1 \
    libxdamage1 \
    libxfixes3 \
    libxrandr2 \
    libgbm1 \
    libpango-1.0-0 \
    libcairo2 \
    libasound2 \
    libatspi2.0-0 \
    libwayland-client0 \
    # Fonts (so pages render properly)
    fonts-liberation \
    fonts-noto-color-emoji \
    fonts-dejavu-core \
    # Utilities
    curl \
    procps \
    && rm -rf /var/lib/apt/lists/*

# ── Application setup ───────────────────────────────────────────
WORKDIR /app

# Install Python dependencies first (layer caching)
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Install patchright's Chromium browser
RUN patchright install chromium && patchright install-deps chromium

# Copy application code
COPY src/ src/
COPY scripts/ scripts/
COPY .env.example .env

# ── Directory setup ─────────────────────────────────────────────
# These will be overridden by volume mounts in docker-compose
RUN mkdir -p /app/browser_data /app/logs /app/downloads/images

# ── Supervisor & entrypoint ─────────────────────────────────────
COPY docker/supervisord.conf /etc/supervisor/conf.d/catgpt.conf
COPY docker/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# ── Environment ─────────────────────────────────────────────────
# Virtual display
ENV DISPLAY=:99
ENV DISPLAY_WIDTH=1280
ENV DISPLAY_HEIGHT=720
ENV DISPLAY_DEPTH=24

# App config (overridable via docker-compose)
ENV HEADLESS=false
ENV BROWSER_DATA_DIR=/app/browser_data
ENV LOG_DIR=/app/logs
ENV API_HOST=0.0.0.0
ENV API_PORT=8000
ENV LOG_LEVEL=DEBUG
ENV VERBOSE=true

# ── Ports ───────────────────────────────────────────────────────
# 8000: FastAPI server
# 6080: noVNC web UI
EXPOSE 8000 6080

# ── Health check ────────────────────────────────────────────────
HEALTHCHECK --interval=30s --timeout=10s --start-period=60s --retries=3 \
    CMD curl -f http://localhost:8000/v1/models || exit 1

ENTRYPOINT ["/entrypoint.sh"]
</file>

<file path="docs/API.md">
# API Reference

CatGPT Gateway exposes an OpenAI-compatible API. Any client that works with the OpenAI API works here.

---

## Table of Contents

- [Base URL](#base-url)
- [Authentication](#authentication)
- [OpenAI-Compatible Endpoints](#openai-compatible-endpoints)
  - [Chat Completions](#chat-completions)
  - [Tool / Function Calling](#tool--function-calling)
  - [Image Input (Vision)](#image-input-vision)
  - [File Attachments](#file-attachments)
  - [Image Generation (ChatGPT only)](#image-generation-chatgpt-only)
  - [List Models](#list-models)
- [Custom REST API](#custom-rest-api)
- [TUI Terminal Client](#tui-terminal-client)
- [Provider Differences](#provider-differences)

---

## Base URL

```
http://localhost:8000/v1
```

## Authentication

Include the Bearer token (default `dummy123`) in every request:

```bash
Authorization: Bearer dummy123
```

With the OpenAI SDK:

```python
client = OpenAI(base_url="http://localhost:8000/v1", api_key="dummy123")
```

Open paths (no auth needed): `/docs`, `/redoc`, `/openapi.json`, `/healthz`

---

## OpenAI-Compatible Endpoints

### Chat Completions

**`POST /v1/chat/completions`**

Standard OpenAI chat completion request.

```python
from openai import OpenAI

client = OpenAI(base_url="http://localhost:8000/v1", api_key="dummy123")

response = client.chat.completions.create(
    model="claude-browser",  # or "catgpt-browser"
    messages=[
        {"role": "system", "content": "You are a helpful assistant."},
        {"role": "user", "content": "What is quantum computing?"}
    ]
)
print(response.choices[0].message.content)
```

```bash
curl -X POST http://localhost:8000/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer dummy123" \
  -d '{
    "model": "claude-browser",
    "messages": [{"role": "user", "content": "Hello!"}]
  }'
```

**Request body:**

| Field | Type | Required | Description |
|---|---|---|---|
| `model` | string | yes | `claude-browser` or `catgpt-browser` |
| `messages` | array | yes | Array of message objects |
| `tools` | array | no | Tool/function definitions |
| `tool_choice` | string/object | no | `auto`, `none`, `required`, or specific function |
| `temperature` | float | no | Ignored (browser controls this) |
| `max_tokens` | int | no | Ignored |
| `stream` | bool | no | Must be `false` (streaming not supported) |

**Response:**

```json
{
  "id": "chatcmpl-abc123...",
  "object": "chat.completion",
  "created": 1716025800,
  "model": "claude-browser",
  "choices": [
    {
      "index": 0,
      "message": {
        "role": "assistant",
        "content": "Quantum computing uses quantum bits...",
        "tool_calls": null
      },
      "finish_reason": "stop"
    }
  ],
  "usage": {
    "prompt_tokens": 25,
    "completion_tokens": 150,
    "total_tokens": 175
  }
}
```

---

### Tool / Function Calling

Define tools in the request and the model will call them when appropriate.

**Request with tools:**

```python
response = client.chat.completions.create(
    model="claude-browser",
    messages=[{"role": "user", "content": "What's the weather in Paris?"}],
    tools=[{
        "type": "function",
        "function": {
            "name": "get_weather",
            "description": "Get weather for a city",
            "parameters": {
                "type": "object",
                "properties": {
                    "city": {"type": "string", "description": "City name"}
                },
                "required": ["city"]
            }
        }
    }]
)
```

**Response when model calls a tool:**

```json
{
  "choices": [{
    "message": {
      "role": "assistant",
      "content": null,
      "tool_calls": [
        {
          "id": "call_a1b2c3d4e5f6...",
          "type": "function",
          "function": {
            "name": "get_weather",
            "arguments": "{\"city\": \"Paris\"}"
          }
        }
      ]
    },
    "finish_reason": "tool_calls"
  }]
}
```

**Sending tool results back:**

```python
# After executing the tool, send the result back
response = client.chat.completions.create(
    model="claude-browser",
    messages=[
        {"role": "user", "content": "What's the weather in Paris?"},
        {"role": "assistant", "tool_calls": [
            {"id": "call_a1b2c3...", "type": "function",
             "function": {"name": "get_weather", "arguments": "{\"city\": \"Paris\"}"}}
        ]},
        {"role": "tool", "tool_call_id": "call_a1b2c3...", "content": "Sunny, 25C"}
    ]
)
# Model responds with natural language summary
```

**LangChain example (full round-trip):**

```python
from langchain_openai import ChatOpenAI
from langchain_core.messages import HumanMessage, ToolMessage
from langchain_core.tools import tool

@tool
def get_weather(city: str) -> str:
    """Get the current weather for a city."""
    return f"Sunny, 25C in {city}"

@tool
def add_numbers(a: int, b: int) -> str:
    """Add two numbers together."""
    return str(a + b)

llm = ChatOpenAI(model="claude-browser", base_url="http://localhost:8000/v1", api_key="dummy123")
llm_with_tools = llm.bind_tools([get_weather, add_numbers])

# Step 1: Model decides to call tools
response = llm_with_tools.invoke([
    HumanMessage(content="Weather in Tokyo and what's 42+58?")
])

# Step 2: Execute tools and send results
messages = [HumanMessage(content="Weather in Tokyo and what's 42+58?"), response]
tool_map = {"get_weather": get_weather, "add_numbers": add_numbers}

for tc in response.tool_calls:
    result = tool_map[tc["name"]].invoke(tc["args"])
    messages.append(ToolMessage(content=str(result), tool_call_id=tc["id"]))

# Step 3: Model summarizes results
final = llm_with_tools.invoke(messages)
print(final.content)
# "It's sunny and 25C in Tokyo, and 42 + 58 = 100."
```

**`tool_choice` options:**

| Value | Behavior |
|---|---|
| `"auto"` (default) | Model decides whether to call tools or answer directly |
| `"required"` | Model must call at least one tool |
| `"none"` | Tools are ignored, model answers directly |
| `{"type":"function","function":{"name":"X"}}` | Model must call the specified function |

---

### Image Input (Vision)

Send images using the standard OpenAI vision format.

```python
import base64

with open("photo.png", "rb") as f:
    img_b64 = base64.b64encode(f.read()).decode()

response = client.chat.completions.create(
    model="claude-browser",
    messages=[{
        "role": "user",
        "content": [
            {"type": "text", "text": "Describe this image in detail."},
            {"type": "image_url", "image_url": {"url": f"data:image/png;base64,{img_b64}"}},
        ]
    }]
)
```

**Multiple images:**

```python
response = client.chat.completions.create(
    model="claude-browser",
    messages=[{
        "role": "user",
        "content": [
            {"type": "text", "text": "Compare these two images."},
            {"type": "image_url", "image_url": {"url": f"data:image/png;base64,{img1_b64}"}},
            {"type": "image_url", "image_url": {"url": f"data:image/png;base64,{img2_b64}"}},
        ]
    }]
)
```

HTTP URLs also work:

```python
{"type": "image_url", "image_url": {"url": "https://example.com/photo.jpg"}}
```

---

### File Attachments

Send PDFs, DOCX, TXT, CSV, and other files via a custom `file` content type.

```python
import base64

with open("document.pdf", "rb") as f:
    pdf_b64 = base64.b64encode(f.read()).decode()

response = client.chat.completions.create(
    model="claude-browser",
    messages=[{
        "role": "user",
        "content": [
            {"type": "text", "text": "Summarize this PDF."},
            {"type": "file", "file": {
                "filename": "document.pdf",
                "data": pdf_b64,
                "mime_type": "application/pdf"
            }},
        ]
    }]
)
```

Alternative data-URL format:

```json
{"type": "file", "file": {"filename": "doc.pdf", "url": "data:application/pdf;base64,..."}}
```

---

### Image Generation (ChatGPT only)

**`POST /v1/images/generations`**

Generate images via DALL-E. Only available when `PROVIDER=chatgpt`. Returns HTTP 501 for Claude.

```python
response = client.images.generate(
    model="dall-e-3",
    prompt="A cyberpunk cat hacking a mainframe",
    n=1,
    size="1024x1024",
    response_format="b64_json",
)

# Save the image
import base64
with open("output.png", "wb") as f:
    f.write(base64.b64decode(response.data[0].b64_json))
```

```bash
curl -X POST http://localhost:8000/v1/images/generations \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer dummy123" \
  -d '{"prompt": "A cat in space", "n": 1, "response_format": "b64_json"}'
```

**Request parameters:**

| Parameter | Type | Default | Description |
|---|---|---|---|
| `prompt` | string | required | Text description of the image |
| `model` | string | `dall-e-3` | Model name (ignored, uses ChatGPT's DALL-E) |
| `n` | int | `1` | Number of images (1-4) |
| `size` | string | `1024x1024` | Requested size (hint to ChatGPT) |
| `quality` | string | `standard` | `standard` or `hd` |
| `style` | string | `vivid` | `vivid` or `natural` |
| `response_format` | string | `b64_json` | `b64_json` or `url` (local file path) |

---

### List Models

**`GET /v1/models`**

Returns the available model based on the active provider.

```bash
curl http://localhost:8000/v1/models -H "Authorization: Bearer dummy123"
```

| Provider | Model ID | Owned By |
|---|---|---|
| Claude | `claude-browser` | `anthropic` |
| ChatGPT | `catgpt-browser` | `catgpt` |

---

## Custom REST API

In addition to the OpenAI-compatible endpoints, CatGPT exposes a simpler custom API:

| Method | Endpoint | Description |
|---|---|---|
| `POST` | `/chat` | Send a message in the current conversation |
| `POST` | `/thread/new` | Start a new conversation |
| `POST` | `/thread/{id}/chat` | Send a message in a specific thread |
| `GET` | `/threads` | List recent threads |
| `GET` | `/status` | Health check, login status, current thread |

```bash
# Chat in current thread
curl -X POST http://localhost:8000/chat \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer dummy123" \
  -d '{"message": "Hello!"}'

# Start new thread
curl -X POST http://localhost:8000/thread/new \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer dummy123" \
  -d '{"message": "New conversation"}'

# Check status
curl -H "Authorization: Bearer dummy123" http://localhost:8000/status
```

---

## TUI Terminal Client

CatGPT includes a terminal chat interface with a cyberpunk theme, built with Textual.

```bash
python -m src.cli.app
```

### Commands

| Command | Description |
|---|---|
| `/new` | Start a fresh conversation |
| `/threads` | List recent threads |
| `/thread <id>` | Switch to a thread |
| `/images` | List downloaded DALL-E images |
| `/status` | Connection details |
| `/clear` | Clear chat display |
| `/help` | Show commands |
| `/exit` | Quit |

Shortcuts: `Ctrl+N` (new), `Ctrl+T` (threads), `Ctrl+L` (clear), `Ctrl+Q` (quit)

---

## Provider Differences

| Behavior | Claude | ChatGPT |
|---|---|---|
| Model ID | `claude-browser` | `catgpt-browser` |
| Image generation | Not supported (501) | Supported (DALL-E) |
| Table rendering | Tab-separated text | Markdown with pipes |
| Avg response time | 15-20s | 7-10s |
| Tool calling prompt | Collaborative framing | Direct instruction |
| `tool_choice` support | Yes | Yes |
| Vision input | Yes | Yes |
| File attachments | Yes | Yes |
</file>

<file path="docs/ARCHITECTURE.md">
# Architecture

A deep dive into how CatGPT Gateway works under the hood.

---

## Table of Contents

- [Overview](#overview)
- [Browser Lifecycle](#browser-lifecycle)
- [Stealth and Anti-Detection](#stealth-and-anti-detection)
- [Message Flow](#message-flow)
- [Response Detection](#response-detection)
- [Tool Calling](#tool-calling)
- [File and Image Upload](#file-and-image-upload)
- [Echo Detection and Recovery](#echo-detection-and-recovery)
- [Selector Fallback System](#selector-fallback-system)
- [Provider Abstraction](#provider-abstraction)
- [Docker Stack](#docker-stack)

---

## Overview

```
Your app (OpenAI SDK / LangChain / curl)
    |
    v
FastAPI server (port 8000)
    |
    |-- openai_routes.py     Translates OpenAI requests to browser actions
    |-- openai_schemas.py    Pydantic models matching OpenAI spec
    v
Provider client (ChatGPTClient or ClaudeClient)
    |
    |-- client.py            send_message(), new_chat(), file upload
    |-- detector.py          Waits for response completion
    |-- selectors.py         DOM selectors for the provider's UI
    v
BrowserManager (Patchright / Playwright)
    |
    |-- stealth.py           Anti-detection patches
    |-- human.py             Human-like typing and clicking
    v
Real Chromium browser --> chatgpt.com or claude.ai
```

---

## Browser Lifecycle

1. **Launch**: `BrowserManager` creates a Patchright persistent browser context at `browser_data/` (or `browser_data_claude/`). This preserves cookies, login state, and Cloudflare clearance across restarts.

2. **DNS Pre-resolution (Docker only)**: Chrome's DNS resolver can fail inside Docker. The entrypoint script pre-resolves `chatgpt.com`, `cdn.oaistatic.com`, `claude.ai`, and related domains via Python's socket module and writes them to `/etc/hosts`. The browser also gets `--host-resolver-rules` flags.

3. **Navigate**: Opens the provider URL with retry logic (up to 5 attempts with exponential backoff).

4. **Stealth (deferred)**: Stealth patches are applied *after* first navigation to avoid breaking DNS. See the stealth section below.

5. **Login Check**: `ensure_logged_in()` checks for login indicators (chat input presence) and prompts if needed.

6. **Client Injection**: The provider client is created and injected into all API routers.

7. **Shutdown**: Browser closes gracefully on FastAPI shutdown.

---

## Stealth and Anti-Detection

The gateway uses multiple techniques to avoid bot detection:

| Technique | How | Where |
|---|---|---|
| Persistent Chrome profile | `launch_persistent_context()` retains cookies and Cloudflare clearance | `browser/manager.py` |
| playwright-stealth | Patches `navigator.webdriver`, WebGL, canvas, plugins | `browser/stealth.py` |
| Docker stealth workaround | Uses `page.evaluate()` instead of `add_init_script()` (the latter breaks DNS in Docker) | `browser/stealth.py` |
| Human-like typing | `keyboard.insert_text()` for paste-style input on contenteditable divs | `browser/human.py` |
| Mouse simulation | Hover before click, natural movement | `browser/human.py` |
| Random delays | 500-1200ms before typing, 300-600ms before sending | `browser/human.py` |
| Viewport jitter | +/- 20px randomization on each launch (1280x720 base) | `browser/manager.py` |
| Headful mode | Always runs with visible browser (headless is trivially detected) | `config.py` |
| Lock file cleanup | Auto-cleans stale `SingletonLock` files from crashed Chrome processes | `browser/manager.py` |

### Docker DNS Fix

`playwright-stealth`'s `add_init_script()` method causes Chrome to fail DNS resolution inside Docker containers. The fix in `stealth.py` uses `page.evaluate()` to inject stealth JS at runtime, and hooks `framenavigated` + `page` events to re-inject on every navigation.

---

## Message Flow

```
send_message(text, image_paths, file_paths)
|
|-- 1. Count existing assistant messages (pre_count)
|-- 2. Random delay (500-1200ms, human simulation)
|-- 3. Upload files if any
|      +-- set_input_files() on hidden <input type="file">
|      +-- Wait 3s + extra per file for processing
|-- 4. Find chat input via selector fallback
|-- 5. Paste text via keyboard.insert_text()
|-- 6. Random delay (300-600ms)
|-- 7. Click send button (or fallback to Enter key)
|-- 8. Wait for response completion (detector)
|-- 9. Sleep 1s for DOM to settle
|-- 10. Check for DALL-E images (ChatGPT only)
|-- 11. Extract response text
|       |-- Image response: DOM scraping
|       +-- Text response: copy button click
+-- 12. Return ChatResponse(message, thread_id, elapsed_ms, images)
```

---

## Response Detection

The detector (`detector.py`) uses multiple strategies to know when the model finishes responding:

### Primary: Copy Button

The copy button only appears after the full response is generated. The detector waits for a copy button on the Nth assistant message (where N = expected count).

### Fallback: Stop Button Lifecycle

While streaming, a "Stop generating" button is visible. The detector watches for it to appear then disappear.

### Fallback: Text Stability

If neither button is found, the detector polls the last assistant message text. If it stays the same for 4+ consecutive checks (2s apart), the response is considered complete.

### Message Counting

Counts both `div[data-message-author-role='assistant']` (ChatGPT) and provider-specific elements. Image responses use different selectors than text responses.

---

## Tool Calling

The web UIs don't have native tool-calling APIs. CatGPT implements tool calling via prompt engineering.

### Flow

1. **Tool definitions** from the OpenAI request are converted into a system prompt describing each function's name, description, and parameter schema.

2. The system prompt instructs the model to output tool calls as structured JSON:
   ```json
   {"tool_calls": [{"name": "get_weather", "arguments": {"city": "Paris"}}]}
   ```

3. The prompt includes examples and rules:
   - Output ONLY the JSON code block when calling tools
   - No commentary before or after the JSON
   - When tool results come back, summarize naturally in plain text

4. **JSON extraction** uses a robust brace-depth tracker (not regex) that handles arbitrarily nested objects, arrays, and escaped strings.

5. Parsed tool calls are returned in standard OpenAI format with generated `call_` IDs.

6. **`tool_choice` support**:
   - `"auto"` (default): model decides whether to use tools
   - `"required"`: prompt says "you MUST call at least one tool"
   - `"none"`: tool prompt is not injected at all
   - `{"type":"function","function":{"name":"X"}}`: prompt says "you MUST call function X"

### Provider-Specific Prompts

- **Claude**: Uses collaborative framing ("You have access to external tools through a structured interface"). Avoids patterns that Claude's web UI flags as prompt injection.
- **ChatGPT**: Uses direct instruction ("You are in tool-calling mode").

### Multi-Turn Tool Calls

When tool results come back as `ToolMessage`s, the gateway builds a prompt transcript showing what was called and what was returned. The model sees the results and produces a natural language summary. The prompt explicitly says "Do NOT call tools again for the same request" to prevent loops.

---

## File and Image Upload

```
API Request (with image_url / file content parts)
|
|-- _extract_image_urls(content)           --> list of URLs/data-URLs
|-- _extract_file_attachments(content)     --> list of {filename, data_b64, mime_type}
|
|-- _download_file(url_or_dict)            --> local file path
|   |-- data: URL       --> base64 decode, save to /tmp/catgpt_files/
|   |-- http: URL       --> download via urllib
|   |-- dict             --> base64 decode with original filename
|   +-- local path       --> pass through
|
|-- image_paths + file_paths --> client.send_message(..., image_paths=, file_paths=)
|
+-- client._upload_files(all_paths)
    |-- Find <input type="file"> via selector
    |-- set_input_files(valid_paths)
    +-- Wait 3s + 1s per additional file
```

---

## Echo Detection and Recovery

Sometimes the copy-button extraction grabs the sent prompt instead of the response (race condition). The gateway detects and recovers:

1. Check if `response_text` contains known markers from the injected tool prompt (`"Available functions:"`, `"tool-calling mode"`, etc.)
2. If echo detected, wait 3 seconds and retry `extract_last_response_via_copy()`
3. If retry still echoes, strip the system prompt prefix and extract the tail content

---

## Selector Fallback System

All DOM selectors are defined as lists of fallbacks tried in order:

```python
CHAT_INPUT = [
    "#prompt-textarea",                                    # Primary
    "div[contenteditable='true'][id='prompt-textarea']",   # Specific
    "div[contenteditable='true']",                         # Broad fallback
]
```

When Claude or ChatGPT update their UI, only `selectors.py` needs changes. The `_find_selector()` method tries each selector with a short timeout and returns the first match.

Each provider has its own selectors file:
- `src/selectors.py` (ChatGPT selectors)
- `src/claude/selectors.py` (Claude selectors)

---

## Provider Abstraction

The gateway supports multiple providers through parallel client implementations:

```
src/chatgpt/
|-- client.py       ChatGPTClient(send_message, new_chat, ...)
|-- detector.py     ChatGPT-specific response detection
|-- selectors.py    (uses src/selectors.py)
|-- image_handler.py  DALL-E image detection
+-- models.py

src/claude/
|-- client.py       ClaudeClient(send_message, new_chat, ...)
|-- detector.py     Claude-specific response detection
+-- selectors.py    Claude DOM selectors
```

Both clients expose the same interface. The API layer (`openai_routes.py`) uses `Config.PROVIDER` to select the appropriate client at startup. Provider-specific logic (tool prompts, response parsing, image generation) is handled with simple `if Config.PROVIDER == "claude"` branches.

---

## Docker Stack

```
Docker Container
|
|-- Xvfb (:99)           Virtual framebuffer, Chrome renders here
|-- x11vnc (:5900)       VNC server capturing the Xvfb display
|-- noVNC (:6080)        WebSocket bridge for browser-based VNC access
+-- FastAPI (:8000)      API server
|
+-- All managed by supervisord
```

### Startup Sequence (entrypoint.sh)

1. Create directories
2. Clean stale Chrome lock files
3. Set up VNC password from `VNC_PASSWORD` env var
4. Pre-resolve DNS domains via Python, write to `/etc/hosts`
5. Verify Xvfb and Patchright Chromium
6. Start supervisord (all 4 services)

### Tech Stack

| Component | Library | Purpose |
|---|---|---|
| Browser automation | Patchright | Playwright fork for Chromium control |
| Anti-detection | playwright-stealth | Patch browser fingerprints |
| API framework | FastAPI | OpenAI-compatible + custom REST API |
| ASGI server | Uvicorn | Serve FastAPI app |
| Data validation | Pydantic | Request/response schemas |
| TUI framework | Textual | Terminal chat interface |
| Rich text | Rich | Markdown rendering |
| Config | python-dotenv | Environment variable loading |
| Container | Docker + Compose | Production deployment |
| Display server | Xvfb + x11vnc + noVNC | Virtual display + browser access |
| Process manager | supervisord | Manage container services |
</file>

<file path="docs/SETUP.md">
# Setup Guide

This guide covers every way to run CatGPT Gateway: Docker, local development, and Nix.

---

## Table of Contents

- [Prerequisites](#prerequisites)
- [Docker Setup (recommended)](#docker-setup-recommended)
- [Local Setup (no Docker)](#local-setup-no-docker)
- [Nix Flake Setup](#nix-flake-setup)
- [First Login](#first-login)
- [Switching Providers](#switching-providers)
- [Authentication](#authentication)
- [Docker Internals](#docker-internals)
- [systemd Service (optional)](#systemd-service-optional)
- [Troubleshooting](#troubleshooting)

---

## Prerequisites

- **Python 3.9+** (local setup only)
- **Docker + Docker Compose** (Docker setup only)
- A **ChatGPT** or **Claude** account (free or paid)

---

## Docker Setup (recommended)

Docker runs the entire stack in one container: virtual display, VNC, browser, and API server.

```bash
# 1. Clone the repo
git clone https://github.com/GautamVhavle/CatGPT-Gateway.git
cd CatGPT-Gateway

# 2. Copy the environment template
cp .env.example .env

# 3. Edit .env to pick your provider
#    Set PROVIDER=claude or PROVIDER=chatgpt

# 4. Build and start
docker compose up --build -d

# 5. First login (one-time) - open the browser UI
open http://localhost:6080
# Sign into Claude or ChatGPT in the browser window you see
# Close the noVNC tab when done - session is saved automatically

# 6. Verify it works
curl -H "Authorization: Bearer dummy123" http://localhost:8000/v1/models
# {"object":"list","data":[{"id":"claude-browser",...}]}

# 7. Send your first message
curl -X POST http://localhost:8000/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer dummy123" \
  -d '{
    "model": "claude-browser",
    "messages": [{"role": "user", "content": "Hello!"}]
  }'
```

### Docker Notes

- **Code is baked into the image.** After editing source files, rebuild:
  ```bash
  docker compose up --build -d   # rebuilds and restarts
  ```
  `docker restart catgpt` does NOT pick up code changes.

- **Browser session persists** via the `catgpt_browser_data` Docker volume. You only need to log in once.

- **Logs** are bind-mounted to `./docker-logs/` on the host.

- **noVNC** at `http://localhost:6080` lets you see and interact with the browser (useful for debugging, CAPTCHAs, or re-login). Default VNC password: `catgpt`.

---

## Local Setup (no Docker)

```bash
# 1. Clone and enter the repo
git clone https://github.com/GautamVhavle/CatGPT-Gateway.git
cd CatGPT-Gateway

# 2. Create a virtual environment
python3 -m venv .venv
source .venv/bin/activate

# 3. Install dependencies
pip install -r requirements.txt

# 4. Install Chromium for Patchright
patchright install chromium

# 5. Copy and configure environment
cp .env.example .env
# Edit .env -> set PROVIDER=claude or PROVIDER=chatgpt

# 6. First login (one-time)
python scripts/first_login.py
# A browser window opens. Sign into your provider. Press Enter when done.

# 7. Start the API server
python -m src.api.server
# API is live at http://localhost:8000

# 8. (Optional) Start the terminal chat UI
python -m src.cli.app
```

---

## Nix Flake Setup

This repo ships a `flake.nix` that packages Patchright and matching Chromium revisions.

```bash
# 1. Copy env template
cp .env.example .env

# 2. First login (one-time, interactive)
nix run .#login

# 3. Start the proxy
nix run .#proxy

# 4. Optional: run the TUI
nix run .#tui
```

Notes:
- The app reads `./.env` from your current working directory if present.
- Shell environment variables override values from `.env`.

---

## First Login

CatGPT Gateway uses your existing browser session. You sign in **once** and the browser profile is persisted.

> **⚠ Google login will not work.**
> Patchright/Chromium runs in a controlled automation context. Google's OAuth detects this and blocks the sign-in.
> **Use email + password, Microsoft, Apple, or magic link / OTP instead.**

### Docker

1. Start the container: `docker compose up --build -d`
2. Wait ~30 seconds for startup
3. Open **http://localhost:6080/vnc.html** (noVNC) in your browser
4. You'll see a Chromium browser inside the VNC viewer
5. Sign into your provider using one of these methods:
   | Method | Works? |
   |---|---|
   | Email + password | ✅ Recommended |
   | Microsoft account | ✅ Works |
   | Apple ID | ✅ Works |
   | Magic link / OTP email | ✅ Works |
   | **Google / "Continue with Google"** | ❌ Blocked by Google |
6. Verify you see the chat interface
7. Close the noVNC tab — your session is saved in the `catgpt_browser_data` Docker volume and survives container restarts.

### Local

1. Run `python scripts/first_login.py`
2. A Chromium window opens and navigates to your provider
3. Sign in using **email + password** or a non-Google method (see table above)
4. Press Enter in the terminal when you see the chat page
5. The browser closes. Session is saved in `browser_data/` (or `browser_data_claude/`).

### Re-login

If your session expires (typically after days/weeks), repeat the login flow. The API returns a 503 error when the session is expired.

---

## Switching Providers

Edit your `.env` file:

```bash
# For Claude
PROVIDER=claude
BROWSER_DATA_DIR=./browser_data_claude

# For ChatGPT
PROVIDER=chatgpt
BROWSER_DATA_DIR=./browser_data
```

Each provider has its own browser data directory so your login sessions don't conflict. After switching, restart the server.

For Docker, also update the `PROVIDER` in `docker-compose.yml` under `environment:` and rebuild.

---

## Authentication

### API Bearer Token

All API endpoints require a Bearer token when `API_TOKEN` is set.

```bash
curl -H "Authorization: Bearer dummy123" http://localhost:8000/v1/models
```

With the OpenAI SDK or LangChain, pass the token as `api_key`:

```python
client = OpenAI(base_url="http://localhost:8000/v1", api_key="dummy123")
```

**Open paths** (no token required): `/docs`, `/redoc`, `/openapi.json`, `/healthz`

To disable auth, set `API_TOKEN=` (empty string) in `.env`.

### noVNC Password

The noVNC browser UI at `http://localhost:6080` is password-protected.

Default: `catgpt`. Change it via `VNC_PASSWORD` in `.env` or `docker-compose.yml`.

---

## Docker Internals

### Container Services (managed by supervisord)

| Service | Port | Purpose |
|---|---|---|
| Xvfb | `:99` | Virtual framebuffer. Chrome renders here. |
| x11vnc | `5900` | VNC server capturing the Xvfb display |
| noVNC | `6080` | WebSocket bridge. Browser-accessible VNC viewer. |
| FastAPI | `8000` | API server (OpenAI-compatible + custom REST) |

### Startup Sequence

1. Create directories (`browser_data`, `logs`, `downloads/images`)
2. Clean stale Chrome lock files
3. Set up VNC password
4. Pre-resolve DNS domains and write to `/etc/hosts` (Docker DNS workaround)
5. Verify Xvfb and Patchright Chromium
6. Start supervisord (manages all 4 services)

### Volumes

| Volume | Purpose |
|---|---|
| `catgpt_browser_data:/app/browser_data` | Persistent browser session (cookies, login) |
| `./docker-logs:/app/logs` | Logs accessible from host |

### Health Check

The container has a built-in health check hitting `/healthz` every 30 seconds.

```bash
docker inspect --format='{{.State.Health.Status}}' catgpt
```

---

## systemd Service (optional)

For running as a background service with the Nix flake:

```ini
# ~/.config/systemd/user/catgpt.service
[Unit]
Description=CatGPT Gateway
After=network-online.target
Wants=network-online.target

[Service]
Type=simple
WorkingDirectory=%h/Projects/CatGPT-Gateway
ExecStart=/usr/bin/env nix run .#proxy
Restart=on-failure
RestartSec=5
Environment=HEADLESS=true
Environment=API_TOKEN=your-token-here
Environment=API_HOST=127.0.0.1
Environment=API_PORT=8000

[Install]
WantedBy=default.target
```

```bash
systemctl --user daemon-reload
systemctl --user enable --now catgpt
journalctl --user -u catgpt -f
```

---

## Troubleshooting

### "ChatGPT client not initialized" (503)

The browser hasn't finished starting. Wait 30-45 seconds after startup.

```bash
# Check logs
docker logs catgpt --tail 50      # Docker
cat logs/api_server.log            # Local
```

### "Not logged in" / session expired

Re-login:
- Docker: Open http://localhost:6080 and sign in
- Local: Run `python scripts/first_login.py`

### Stale browser lock files

If the app crashes, orphan Chrome processes may leave lock files:

```bash
pkill -f "chrome-for-testing" 2>/dev/null
rm -f browser_data/SingletonLock browser_data/SingletonSocket browser_data/SingletonCookie
```

The app auto-cleans these on startup, but manual cleanup may be needed after hard crashes.

### Docker DNS issues

Chrome inside Docker sometimes fails to resolve domains. The entrypoint script pre-resolves domains via Python. If you still see DNS errors:

```bash
docker exec catgpt cat /etc/hosts
docker exec catgpt curl -s https://chatgpt.com
```

### Code changes not taking effect (Docker)

You must rebuild:

```bash
docker compose up --build -d   # correct
# NOT: docker restart catgpt   # this uses the old image
```

### Services not running

```bash
docker exec catgpt supervisorctl status
```

All 4 services (xvfb, vnc, novnc, catgpt) should show `RUNNING`.
</file>

<file path="docs/TEST_REPORT.md">
# CatGPT Gateway — Test Report

**Date:** May 18, 2026  
**Environment:** macOS (local Chromium via Patchright)  
**Providers Tested:** ChatGPT (`chatgpt.com`), Claude (`claude.ai`)

---

## Summary

| Test Suite | ChatGPT | Claude | Notes |
|---|---|---|---|
| Phase 1 (basic pipeline) | ✅ PASS | ✅ PASS | |
| Multi-Turn (10 messages) | ✅ PASS | ✅ PASS | |
| Robustness (tables, code, long) | ✅ 5/5 | ✅ 5/5 | Claude renders tables as tab-separated |
| Image Detection | ✅ 4/4 | ✅ 4/4 | Image gen tests auto-skipped for Claude |
| LangChain Tool Calling | ✅ PASS | ✅ PASS | |
| API: `/v1/chat/completions` | ✅ 200 | ✅ 200 | |
| API: `/v1/models` | ✅ 200 | ✅ 200 | Returns `catgpt-browser` / `claude-browser` |
| API: `/v1/images/generations` | ✅ 200 | ✅ 501 | Claude returns "not supported" (expected) |

**Overall: ALL TESTS PASS on both providers.**

---

## 1. Phase 1 — Basic Pipeline (`test_phase1.py`)

Validates: browser launch → login check → send message → receive response.

### ChatGPT
```
Input:  "Hello! Please respond with exactly: 'Phase 1 test successful.' Nothing else."
Output: "Phase 1 test successful."
Time:   ~7.4s
Thread: assigned automatically
```

### Claude
```
Input:  "Hello! Please respond with exactly: 'Phase 1 test successful.' Nothing else."
Output: "Phase 1 test successful."
Time:   ~17.0s
Thread: 6e817c38-fc13-4311-af83-6bcd6c4dea57
```

---

## 2. Multi-Turn Conversations (`test_multi_turn.py`)

Validates: new chat → 5 follow-up messages → new chat again → 5 more messages.

### ChatGPT
```
Round 1: 5/5 messages — all received correctly
Round 2: 5/5 messages — all received correctly
Threads differ: ✅ Yes (separate threads)
Total time: ~70s
Avg response: ~7s per message
```

### Claude
```
Round 1: 5/5 messages — all received correctly
Round 2: 5/5 messages — all received correctly
Threads differ: ✅ Yes (separate threads)
Total time: ~180s
Avg response: ~18s per message
```

---

## 3. Robustness — Complex Content (`test_robust.py`)

Validates: tables, code blocks, long responses, follow-up context, mixed content.

### ChatGPT — 5/5 PASS
| Test | Result | Time | Response Size |
|---|---|---|---|
| Table output | ✅ PASS | ~8s | ~300 chars |
| Code block (`fibonacci`) | ✅ PASS | ~7s | ~400 chars |
| Long response (TCP/IP) | ✅ PASS | ~12s | ~3800 chars |
| Follow-up context | ✅ PASS | ~7s | ~600 chars |
| Mixed content (OSI table) | ✅ PASS | ~8s | ~600 chars |

### Claude — 5/5 PASS
| Test | Result | Time | Response Size |
|---|---|---|---|
| Table output | ✅ PASS | ~19s | ~368 chars |
| Code block (`fibonacci`) | ✅ PASS | ~18s | ~432 chars |
| Long response (TCP/IP) | ✅ PASS | ~37s | ~3949 chars |
| Follow-up context | ✅ PASS | ~20s | ~618 chars |
| Mixed content (OSI table) | ✅ PASS | ~20s | ~545 chars |

> **Note:** Claude renders tables as tab-separated text (no `|` pipes). The test validates tab-separated format for Claude and pipe-separated for ChatGPT. Both formats contain correct data.

---

## 4. Image Detection (`test_images.py`)

Validates: text-only responses have no false image detection; image generation works.

### ChatGPT — 4/4 PASS
| Test | Result | Details |
|---|---|---|
| Text-only (no false positive) | ✅ PASS | "4" / "four" detected, no images |
| Image generation — simple | ✅ PASS | Image downloaded successfully |
| Follow-up text after image | ✅ PASS | Text only, no false positive |
| Image generation — specific style | ✅ PASS | Image downloaded successfully |

### Claude — 4/4 PASS
| Test | Result | Details |
|---|---|---|
| Text-only (no false positive) | ✅ PASS | "Four." — no false image detection |
| Image generation — simple | ⏭️ SKIP | Not supported by Claude |
| Follow-up text after image | ✅ PASS | Text only, no false positive |
| Image generation — specific style | ⏭️ SKIP | Not supported by Claude |

> **Note:** Claude does not have an image generation capability (no DALL-E equivalent). Image generation tests are automatically skipped when `PROVIDER=claude`. The API endpoint returns HTTP 501 with a clear error message.

---

## 5. LangChain Tool Calling (`test_langchain_tools.py`)

Validates: OpenAI-compatible API works with LangChain's `ChatOpenAI` + tool binding.

### ChatGPT
| Test | Result | Tool Calls? | Details |
|---|---|---|---|
| Models endpoint | ✅ PASS | — | Returns `catgpt-browser` |
| Simple chat (no tools) | ✅ PASS | — | Answered correctly |
| `get_current_time` | ⚠️ PASS | No | Answered directly (not always tool-routed) |
| `add_numbers(42, 58)` | ✅ PASS | No | Answered directly: "42 + 58 = 100" |
| Weather + math (multi-tool) | ✅ PASS | Yes | `weather_forecast` + `calculate_expression` called |
| Reverse + Wikipedia (multi-tool) | ✅ PASS | Yes | `reverse_string` + `search_wikipedia` called |

### Claude
| Test | Result | Tool Calls? | Details |
|---|---|---|---|
| Models endpoint | ✅ PASS | — | Returns `claude-browser` |
| Simple chat (no tools) | ✅ PASS | — | Answered with web search |
| `get_current_time` | ✅ PASS | Yes | JSON tool call emitted, result: `2026-05-18 03:28:09` |
| `add_numbers(42, 58)` | ✅ PASS | Yes | JSON tool call emitted, result: `100` |
| Weather + math (multi-tool) | ✅ PASS | Yes | Both tools called in single response |
| Reverse + Wikipedia (multi-tool) | ✅ PASS | Yes | Both tools called, results summarized |

### Tool Calling Round-Trip Example (Claude)
```
User:    "What is the weather in Paris tomorrow, and what is 7*8+3?"
Claude:  {"tool_calls": [
           {"name": "weather_forecast", "arguments": {"city": "Paris", "date": "tomorrow"}},
           {"name": "calculate_expression", "arguments": {"expression": "7*8+3"}}
         ]}
         → Tools executed → Results sent back
Claude:  "Weather in Paris tomorrow: Sunny with a high of 25°C — a lovely day!
          7 × 8 + 3 = 59"
```

---

## 6. API Endpoints

### `POST /v1/chat/completions`

```bash
curl -X POST http://localhost:8000/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer dummy123" \
  -d '{"model": "claude-browser", "messages": [{"role": "user", "content": "Say exactly: API test passed"}]}'
```

**ChatGPT Response:**
```json
{"choices": [{"message": {"role": "assistant", "content": "API test passed"}, "finish_reason": "stop"}]}
```

**Claude Response:**
```json
{"choices": [{"message": {"role": "assistant", "content": "API test passed"}, "finish_reason": "stop"}]}
```

### `GET /v1/models`

| Provider | Model ID | Owned By |
|---|---|---|
| ChatGPT | `catgpt-browser` | `catgpt` |
| Claude | `claude-browser` | `anthropic` |

### `POST /v1/images/generations`

| Provider | Status | Response |
|---|---|---|
| ChatGPT | `200 OK` | Base64 image data or local file URL |
| Claude | `501 Not Implemented` | `"Image generation is not supported by Claude. This feature is only available with the ChatGPT provider."` |

---

## Provider Comparison

| Feature | ChatGPT | Claude |
|---|---|---|
| Chat completions | ✅ | ✅ |
| Multi-turn conversations | ✅ | ✅ |
| Tool/function calling | ✅ | ✅ |
| Image generation (DALL-E) | ✅ | ❌ (501) |
| Image input (vision) | ✅ | ✅ |
| File attachments | ✅ | ✅ |
| Tables in response | Markdown (`\|`) | Tab-separated |
| Avg response time | ~7-8s | ~17-20s |
| Model ID | `catgpt-browser` | `claude-browser` |

---

## Configuration

Switch providers via `.env`:
```env
# Provider: "chatgpt" or "claude"
PROVIDER=chatgpt

# Browser data directories (separate per provider)
BROWSER_DATA_DIR=./browser_data          # for chatgpt
# BROWSER_DATA_DIR=./browser_data_claude # for claude
```

Or via environment variable:
```bash
PROVIDER=claude BROWSER_DATA_DIR=./browser_data_claude python -m src.api.server
```

---

## How to Run Tests

```bash
# Activate virtualenv
source .venv/bin/activate

# Direct browser tests (no server needed)
python scripts/test_phase1.py
python scripts/test_multi_turn.py
python scripts/test_robust.py
python scripts/test_images.py

# API tests (start server first)
python -m src.api.server &
python scripts/test_langchain_tools.py
python scripts/test_image_generation.py
```

All test scripts auto-detect the provider from the `PROVIDER` environment variable or `.env` file.
</file>

<file path="flake.lock">
{
  "nodes": {
    "flake-utils": {
      "inputs": {
        "systems": "systems"
      },
      "locked": {
        "lastModified": 1731533236,
        "narHash": "sha256-l0KFg5HjrsfsO/JpG+r7fRrqm12kzFHyUHqHCVpMMbI=",
        "owner": "numtide",
        "repo": "flake-utils",
        "rev": "11707dc2f618dd54ca8739b309ec4fc024de578b",
        "type": "github"
      },
      "original": {
        "owner": "numtide",
        "repo": "flake-utils",
        "type": "github"
      }
    },
    "nixpkgs": {
      "locked": {
        "lastModified": 1771369470,
        "narHash": "sha256-0NBlEBKkN3lufyvFegY4TYv5mCNHbi5OmBDrzihbBMQ=",
        "owner": "NixOS",
        "repo": "nixpkgs",
        "rev": "0182a361324364ae3f436a63005877674cf45efb",
        "type": "github"
      },
      "original": {
        "owner": "NixOS",
        "ref": "nixos-unstable",
        "repo": "nixpkgs",
        "type": "github"
      }
    },
    "root": {
      "inputs": {
        "flake-utils": "flake-utils",
        "nixpkgs": "nixpkgs"
      }
    },
    "systems": {
      "locked": {
        "lastModified": 1681028828,
        "narHash": "sha256-Vy1rq5AaRuLzOxct8nz4T6wlgyUR7zLU309k9mBC768=",
        "owner": "nix-systems",
        "repo": "default",
        "rev": "da67096a3b9bf56a91d16901293e51ba5b49a27e",
        "type": "github"
      },
      "original": {
        "owner": "nix-systems",
        "repo": "default",
        "type": "github"
      }
    }
  },
  "root": "root",
  "version": 7
}
</file>

<file path="flake.nix">
{
  description = "CatGPT browser-automation gateway";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
    }:
    flake-utils.lib.eachSystem [ "x86_64-linux" ] (
      system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };
        lib = pkgs.lib;

        src = ./.;

        patchrightVersion = "1.58.0";
        patchrightChromiumRevision = "1208";
        patchrightChromiumVersion = "145.0.7632.6";

        fontconfigFile = pkgs.makeFontsConf {
          fontDirectories = [
            pkgs.dejavu_fonts
            pkgs.liberation_ttf
            pkgs.noto-fonts
            pkgs.noto-fonts-cjk-sans
            pkgs.noto-fonts-color-emoji
          ];
        };

        patchrightChromium = pkgs.stdenv.mkDerivation {
          pname = "patchright-chromium";
          version = patchrightChromiumRevision;

          src = pkgs.fetchzip {
            url = "https://cdn.playwright.dev/chrome-for-testing-public/${patchrightChromiumVersion}/linux64/chrome-linux64.zip";
            hash = "sha256-akvAXdfBKdjDQBnWTDX0WbmP+niXthXlyB9feeq8kyw=";
            stripRoot = false;
          };

          nativeBuildInputs = [
            pkgs.autoPatchelfHook
            pkgs.makeWrapper
            pkgs.patchelf
          ];

          buildInputs = [
            pkgs.alsa-lib
            pkgs.at-spi2-atk
            pkgs.atk
            pkgs.cairo
            pkgs.cups
            pkgs.dbus
            pkgs.expat
            pkgs.glib
            pkgs.gobject-introspection
            pkgs.libdrm
            pkgs.libgbm
            pkgs.libxkbcommon
            pkgs.nspr
            pkgs.nss
            pkgs.pango
            pkgs.stdenv.cc.cc.lib
            pkgs.systemd
            pkgs.libx11
            pkgs.libxcomposite
            pkgs.libxdamage
            pkgs.libxext
            pkgs.libxfixes
            pkgs.libxi
            pkgs.libxrandr
            pkgs.libxrender
            pkgs.libxtst
            pkgs.libxcb
            pkgs.libxshmfence
          ];

          installPhase = ''
            runHook preInstall

            mkdir -p "$out/chrome-linux64"
            cp -R chrome-linux64/. "$out/chrome-linux64"
            chmod -R u+w "$out/chrome-linux64"
            touch "$out/INSTALLATION_COMPLETE"

            wrapProgram "$out/chrome-linux64/chrome" \
              --set-default SSL_CERT_FILE "${pkgs.cacert}/etc/ssl/certs/ca-bundle.crt" \
              --set-default FONTCONFIG_FILE "${fontconfigFile}"

            runHook postInstall
          '';

          appendRunpaths = lib.makeLibraryPath [
            pkgs.libGL
            pkgs.pciutils
            pkgs.vulkan-loader
          ];

          postFixup = ''
            if [ -e "$out/chrome-linux64/libvulkan.so.1" ]; then
              rm "$out/chrome-linux64/libvulkan.so.1"
              ln -s "${lib.getLib pkgs.vulkan-loader}/lib/libvulkan.so.1" "$out/chrome-linux64/libvulkan.so.1"
            fi
          '';
        };

        patchrightChromiumHeadlessShell = pkgs.stdenv.mkDerivation {
          pname = "patchright-chromium-headless-shell";
          version = patchrightChromiumRevision;

          src = pkgs.fetchzip {
            url = "https://cdn.playwright.dev/chrome-for-testing-public/${patchrightChromiumVersion}/linux64/chrome-headless-shell-linux64.zip";
            hash = "sha256-/xskLzTc9tTZmu1lwkMpjV3QV7XjP92D/7zRcFuVWT8=";
            stripRoot = false;
          };

          nativeBuildInputs = [
            pkgs.autoPatchelfHook
            pkgs.makeWrapper
            pkgs.patchelf
          ];

          buildInputs = [
            pkgs.alsa-lib
            pkgs.at-spi2-atk
            pkgs.atk
            pkgs.cairo
            pkgs.cups
            pkgs.dbus
            pkgs.expat
            pkgs.glib
            pkgs.gobject-introspection
            pkgs.libdrm
            pkgs.libgbm
            pkgs.libxkbcommon
            pkgs.nspr
            pkgs.nss
            pkgs.pango
            pkgs.stdenv.cc.cc.lib
            pkgs.systemd
            pkgs.libx11
            pkgs.libxcomposite
            pkgs.libxdamage
            pkgs.libxext
            pkgs.libxfixes
            pkgs.libxi
            pkgs.libxrandr
            pkgs.libxrender
            pkgs.libxtst
            pkgs.libxcb
            pkgs.libxshmfence
          ];

          installPhase = ''
            runHook preInstall

            mkdir -p "$out/chrome-headless-shell-linux64"
            cp -R chrome-headless-shell-linux64/. "$out/chrome-headless-shell-linux64"
            chmod -R u+w "$out/chrome-headless-shell-linux64"
            touch "$out/INSTALLATION_COMPLETE"

            wrapProgram "$out/chrome-headless-shell-linux64/chrome-headless-shell" \
              --set-default SSL_CERT_FILE "${pkgs.cacert}/etc/ssl/certs/ca-bundle.crt" \
              --set-default FONTCONFIG_FILE "${fontconfigFile}"

            runHook postInstall
          '';

          appendRunpaths = lib.makeLibraryPath [
            pkgs.libGL
            pkgs.pciutils
            pkgs.vulkan-loader
          ];
        };

        patchrightBrowsers = pkgs.linkFarm "patchright-browsers" [
          {
            name = "chromium-${patchrightChromiumRevision}";
            path = patchrightChromium;
          }
          {
            name = "chromium_headless_shell-${patchrightChromiumRevision}";
            path = patchrightChromiumHeadlessShell;
          }
        ];

        python = pkgs.python311.override {
          packageOverrides = pythonSelf: pythonSuper: {
            patchright = pythonSelf.buildPythonPackage rec {
              pname = "patchright";
              version = patchrightVersion;
              format = "wheel";

              src = pkgs.fetchPypi {
                inherit pname version format;
                dist = "py3";
                python = "py3";
                platform = "manylinux1_x86_64";
                hash = "sha256-gyvuL+SM+dwHuzsPDQXu6SMgPzSM2YsUwsUV7s4yZzQ=";
              };

              propagatedBuildInputs = [
                pythonSelf.greenlet
                pythonSelf.pyee
              ];

              doCheck = false;
              pythonImportsCheck = [ "patchright" ];

              meta = {
                description = "Undetected Python version of the Playwright automation library";
                homepage = "https://github.com/Kaliiiiiiiiii-Vinyzu/patchright-python";
                license = lib.licenses.asl20;
              };
            };
          };
        };

        pythonEnv = python.withPackages (ps: [
          ps.fastapi
          ps.uvicorn
          ps.pydantic
          ps."python-dotenv"
          ps.patchright
          ps."playwright-stealth"
          ps.textual
          ps.typer
          ps.rich
        ]);

        mkCatgptScript =
          {
            name,
            command,
          }:
          pkgs.writeShellApplication {
            inherit name;
            runtimeInputs = [
              pythonEnv
              pkgs.nodejs
            ];
            text = ''
              runtime_root="$PWD"
              if [ ! -d "$runtime_root/src" ] || [ ! -d "$runtime_root/scripts" ]; then
                runtime_root="${src}"
              fi

              export PLAYWRIGHT_BROWSERS_PATH="${patchrightBrowsers}"
              export PLAYWRIGHT_SKIP_BROWSER_DOWNLOAD=1
              export PLAYWRIGHT_SKIP_BROWSER_GC=1
              export PLAYWRIGHT_SKIP_VALIDATE_HOST_REQUIREMENTS=true
              export PLAYWRIGHT_NODEJS_PATH="${pkgs.nodejs}/bin/node"

              state_home="''${XDG_STATE_HOME:-$HOME/.local/state}"

              if [ "$runtime_root" = "${src}" ] && [ ! -f "$runtime_root/.env" ]; then
                export BROWSER_DATA_DIR="''${BROWSER_DATA_DIR:-$state_home/catgpt/browser_data}"
                export LOG_DIR="''${LOG_DIR:-$state_home/catgpt/logs}"
                export IMAGES_DIR="''${IMAGES_DIR:-$state_home/catgpt/downloads/images}"
              fi

              mkdir -p "''${BROWSER_DATA_DIR:-$runtime_root/browser_data}" \
                       "''${LOG_DIR:-$runtime_root/logs}" \
                       "''${IMAGES_DIR:-$runtime_root/downloads/images}"

              export PYTHONPATH="$runtime_root:${src}:''${PYTHONPATH:-}"

              cd "$runtime_root"
              exec python ${command}
            '';
          };

        proxyScript = mkCatgptScript {
          name = "catgpt-proxy";
          command = "-m src.api.server";
        };

        loginScript = mkCatgptScript {
          name = "catgpt-login";
          command = "scripts/first_login.py";
        };

        tuiScript = mkCatgptScript {
          name = "catgpt-tui";
          command = "-m src.cli.app";
        };
      in
      {
        packages = {
          default = proxyScript;
          proxy = proxyScript;
          login = loginScript;
          tui = tuiScript;
          python-env = pythonEnv;
          patchright-browsers = patchrightBrowsers;
        };

        apps = {
          default = {
            type = "app";
            program = "${proxyScript}/bin/catgpt-proxy";
            meta.description = "Run CatGPT FastAPI proxy";
          };
          proxy = {
            type = "app";
            program = "${proxyScript}/bin/catgpt-proxy";
            meta.description = "Run CatGPT FastAPI proxy";
          };
          login = {
            type = "app";
            program = "${loginScript}/bin/catgpt-login";
            meta.description = "Run one-time ChatGPT login flow";
          };
          tui = {
            type = "app";
            program = "${tuiScript}/bin/catgpt-tui";
            meta.description = "Run CatGPT terminal UI";
          };
        };

        checks.imports =
          pkgs.runCommand "catgpt-import-check"
            {
              nativeBuildInputs = [
                pythonEnv
                pkgs.nodejs
              ];
              PLAYWRIGHT_BROWSERS_PATH = patchrightBrowsers;
              PLAYWRIGHT_SKIP_BROWSER_DOWNLOAD = "1";
              PLAYWRIGHT_SKIP_BROWSER_GC = "1";
              PLAYWRIGHT_SKIP_VALIDATE_HOST_REQUIREMENTS = "true";
              PLAYWRIGHT_NODEJS_PATH = "${pkgs.nodejs}/bin/node";
              PYTHONPATH = src;
            }
            ''
              export BROWSER_DATA_DIR="$TMPDIR/browser_data"
              export LOG_DIR="$TMPDIR/logs"
              export IMAGES_DIR="$TMPDIR/images"
              mkdir -p "$BROWSER_DATA_DIR" "$LOG_DIR" "$IMAGES_DIR"

              python - <<'PY'
              import patchright
              import playwright_stealth
              import src.api.server
              PY
              touch "$out"
            '';

        devShells.default = pkgs.mkShell {
          packages = [
            pythonEnv
            pkgs.nodejs
          ];

          PLAYWRIGHT_BROWSERS_PATH = patchrightBrowsers;
          PLAYWRIGHT_SKIP_BROWSER_DOWNLOAD = "1";
          PLAYWRIGHT_SKIP_BROWSER_GC = "1";
          PLAYWRIGHT_SKIP_VALIDATE_HOST_REQUIREMENTS = "true";
          PLAYWRIGHT_NODEJS_PATH = "${pkgs.nodejs}/bin/node";

          shellHook = ''
            if [ ! -f .env ]; then
              export BROWSER_DATA_DIR="''${BROWSER_DATA_DIR:-$PWD/browser_data}"
              export LOG_DIR="''${LOG_DIR:-$PWD/logs}"
              export IMAGES_DIR="''${IMAGES_DIR:-$PWD/downloads/images}"
            fi

            mkdir -p "''${BROWSER_DATA_DIR:-$PWD/browser_data}" \
                     "''${LOG_DIR:-$PWD/logs}" \
                     "''${IMAGES_DIR:-$PWD/downloads/images}"
          '';
        };

        formatter = pkgs.nixfmt;
      }
    );
}
</file>

<file path="LICENSE">
MIT License

Copyright (c) 2026 GautamVhavle

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
</file>

<file path="requirements.txt">
# Core
patchright>=1.58
playwright-stealth>=2.0.2

# TUI (Phase 2)
textual>=0.85
typer>=0.12
rich>=13.0

# API (Phase 3)
fastapi>=0.115
uvicorn>=0.32

# Data & Config
pydantic>=2.5
python-dotenv>=1.0

# OpenAI-compatible testing
openai>=1.0
langchain-openai>=0.1
langchain>=0.2
</file>

<file path="scripts/.gitkeep">

</file>

<file path="scripts/first_login.py">
#!/usr/bin/env python3
"""
First Login -- One-time script to open browser and let user sign in manually.

Run this once to create a persistent browser session. After signing in,
the session (cookies, tokens) is saved to browser_data/ and reused.

Usage:
    python scripts/first_login.py
"""

import asyncio
import sys
import os

# Add project root to path
sys.path.insert(0, os.path.dirname(os.path.dirname(os.path.abspath(__file__))))

from src.browser.manager import BrowserManager
from src.config import Config
from src.log import setup_logging

log = setup_logging("first_login", log_file="first_login.log")


async def main():
    browser = BrowserManager()

    provider_name = Config.PROVIDER.title()
    provider_url = Config.provider_url()

    try:
        print("\n" + "=" * 60)
        print(f"  CatGPT Gateway -- {provider_name} First Login")
        print("=" * 60)
        print(f"\n  Provider:         {provider_name}")
        print(f"  Browser data dir: {Config.BROWSER_DATA_DIR}")
        print(f"  Target:           {provider_url}")
        print("\n  " + "!" * 56)
        print("  IMPORTANT: Google login will NOT work here.")
        print("  Chromium in a controlled/automated context is blocked")
        print("  by Google's OAuth bot detection. Use instead:")
        print("    • Email + password  (most reliable)")
        print("    • Microsoft account")
        print("    • Apple ID")
        print("    • Magic link / OTP sent to your email")
        print("  " + "!" * 56)
        print("\n  A Chrome window will open. Please:")
        print(f"  1. Sign in to {provider_name} with your account (NOT Google)")
        print("  2. Complete any CAPTCHA / Cloudflare checks")
        print("  3. Wait until you see the chat interface")
        print("  4. Come back here and press Enter")
        print("\n" + "=" * 60 + "\n")

        # Launch browser
        page = await browser.start()

        # Navigate to provider
        await browser.navigate(provider_url)

        print("  Browser opened. Sign in now...")
        print()

        # Wait for user input
        input("  Press ENTER after you've signed in successfully > ")

        # Verify login
        logged_in = await browser.is_logged_in()

        if logged_in:
            print(f"\n  Login verified! Session saved to {Config.BROWSER_DATA_DIR}/")
            print("  You won't need to sign in again.\n")
            log.info("First login completed successfully")
        else:
            print("\n  Could not verify login. Session may still be saved.")
            print("  Try running test_phase1.py to check.\n")
            log.warning("Login verification uncertain")

    except KeyboardInterrupt:
        print("\n\n  Cancelled by user.")
    finally:
        await browser.close()
        print("  Browser closed.\n")


if __name__ == "__main__":
    asyncio.run(main())
</file>

<file path="scripts/test_image_generation.py">
#!/usr/bin/env python3
"""
Test script for the OpenAI-compatible Image Generation endpoint.

Tests the POST /v1/images/generations endpoint by generating images
via the CatGPT API and verifying the response format matches OpenAI's spec.

Tests:
  1. Generate a single image (b64_json format)
  2. Generate an image and save to disk
  3. Generate an image (url format — returns local file path)
  4. Use the OpenAI SDK client.images.generate()

Prerequisites:
  - CatGPT API server running: python -m src.api.server
  - OR Docker: docker compose up --build -d catgpt
  - pip install openai requests

Usage:
  python scripts/test_image_generation.py
"""

from __future__ import annotations

import base64
import json
import os
import sys
import time
from pathlib import Path

try:
    import requests
except ImportError:
    print("ERROR: pip install requests")
    sys.exit(1)

try:
    from openai import OpenAI
except ImportError:
    print("WARNING: openai SDK not installed — Test 4 will be skipped")
    OpenAI = None


# ── Configuration ───────────────────────────────────────────────

BASE_URL = "http://localhost:8000"
API_KEY = "dummy123"
OUTPUT_DIR = Path(__file__).resolve().parent.parent / "downloads" / "images" / "test_generated"

HEADERS = {
    "Content-Type": "application/json",
    "Authorization": f"Bearer {API_KEY}",
}


# ── Helpers ─────────────────────────────────────────────────────

def separator(title: str) -> None:
    print(f"\n{'='*70}")
    print(f"  {title}")
    print(f"{'='*70}\n")


def save_b64_image(b64_data: str, filename: str) -> str:
    """Save a base64-encoded image to disk. Returns the file path."""
    OUTPUT_DIR.mkdir(parents=True, exist_ok=True)
    filepath = OUTPUT_DIR / filename
    img_bytes = base64.b64decode(b64_data)
    filepath.write_bytes(img_bytes)
    size_kb = len(img_bytes) / 1024
    print(f"  Saved: {filepath} ({size_kb:.1f} KB)")
    return str(filepath)


# ── Tests ───────────────────────────────────────────────────────

def test_1_basic_image_generation():
    """Test 1: Generate a single image via raw HTTP request (b64_json)."""
    separator("Test 1: Basic Image Generation (b64_json)")

    payload = {
        "prompt": "A cute orange tabby cat sitting on a keyboard, digital art style",
        "model": "dall-e-3",
        "n": 1,
        "size": "1024x1024",
        "response_format": "b64_json",
    }

    print(f"  Prompt: {payload['prompt']}")
    print(f"  Sending request...")
    start = time.time()

    resp = requests.post(f"{BASE_URL}/v1/images/generations", headers=HEADERS, json=payload, timeout=180)
    elapsed = time.time() - start

    print(f"  Status: {resp.status_code} ({elapsed:.1f}s)")

    if resp.status_code != 200:
        print(f"  FAILED: {resp.text[:500]}")
        return False

    data = resp.json()
    print(f"  Response keys: {list(data.keys())}")

    # Validate response structure
    assert "created" in data, "Missing 'created' field"
    assert "data" in data, "Missing 'data' field"
    assert isinstance(data["data"], list), "'data' should be a list"
    assert len(data["data"]) >= 1, "'data' should have at least 1 image"

    img = data["data"][0]
    assert "b64_json" in img and img["b64_json"], "Missing 'b64_json' in image data"

    # Validate it's valid base64
    try:
        img_bytes = base64.b64decode(img["b64_json"])
        size_kb = len(img_bytes) / 1024
        print(f"  Image size: {size_kb:.1f} KB")
    except Exception as e:
        print(f"  FAILED: Invalid base64 data: {e}")
        return False

    # Check for revised_prompt (optional but nice)
    if img.get("revised_prompt"):
        print(f"  Revised prompt: {img['revised_prompt'][:100]}")

    # Save the image
    save_b64_image(img["b64_json"], "test1_cat_keyboard.png")

    print(f"\n  PASSED ✓ (generated 1 image in {elapsed:.1f}s)")
    return True


def test_2_save_generated_image():
    """Test 2: Generate an image with a different prompt and verify save."""
    separator("Test 2: Generate & Save Image")

    payload = {
        "prompt": "A futuristic cyberpunk cityscape at night with neon lights, 4k quality",
        "model": "dall-e-3",
        "n": 1,
        "size": "1024x1024",
        "quality": "hd",
        "response_format": "b64_json",
    }

    print(f"  Prompt: {payload['prompt']}")
    print(f"  Quality: {payload['quality']}")
    print(f"  Sending request...")
    start = time.time()

    resp = requests.post(f"{BASE_URL}/v1/images/generations", headers=HEADERS, json=payload, timeout=180)
    elapsed = time.time() - start

    print(f"  Status: {resp.status_code} ({elapsed:.1f}s)")

    if resp.status_code != 200:
        print(f"  FAILED: {resp.text[:500]}")
        return False

    data = resp.json()
    img = data["data"][0]

    if not img.get("b64_json"):
        print(f"  FAILED: No b64_json in response")
        return False

    filepath = save_b64_image(img["b64_json"], "test2_cyberpunk_city.png")

    # Verify the saved file exists and has content
    saved = Path(filepath)
    assert saved.exists(), f"Saved file does not exist: {filepath}"
    assert saved.stat().st_size > 1000, f"Saved file too small: {saved.stat().st_size} bytes"

    print(f"  File verified: {saved.stat().st_size} bytes")
    print(f"\n  PASSED ✓ ({elapsed:.1f}s)")
    return True


def test_3_url_format():
    """Test 3: Generate an image with response_format='url'."""
    separator("Test 3: Image Generation (url format)")

    payload = {
        "prompt": "A simple watercolor painting of a mountain landscape with a lake",
        "model": "dall-e-3",
        "n": 1,
        "size": "1024x1024",
        "response_format": "url",
    }

    print(f"  Prompt: {payload['prompt']}")
    print(f"  Response format: url")
    print(f"  Sending request...")
    start = time.time()

    resp = requests.post(f"{BASE_URL}/v1/images/generations", headers=HEADERS, json=payload, timeout=180)
    elapsed = time.time() - start

    print(f"  Status: {resp.status_code} ({elapsed:.1f}s)")

    if resp.status_code != 200:
        print(f"  FAILED: {resp.text[:500]}")
        return False

    data = resp.json()
    img = data["data"][0]

    assert "url" in img and img["url"], "Missing 'url' in image data"
    print(f"  Image URL/path: {img['url']}")

    # If it's a local path, verify the file exists
    if not img["url"].startswith("http"):
        path = Path(img["url"])
        if path.exists():
            print(f"  File exists: {path.stat().st_size} bytes")
        else:
            print(f"  WARNING: Local path does not exist: {img['url']}")

    if img.get("revised_prompt"):
        print(f"  Revised prompt: {img['revised_prompt'][:100]}")

    print(f"\n  PASSED ✓ ({elapsed:.1f}s)")
    return True


def test_4_openai_sdk():
    """Test 4: Use the official OpenAI Python SDK."""
    separator("Test 4: OpenAI SDK — client.images.generate()")

    if OpenAI is None:
        print("  SKIPPED — openai SDK not installed")
        return True

    client = OpenAI(base_url=f"{BASE_URL}/v1", api_key=API_KEY)

    prompt = "A photorealistic golden retriever puppy wearing sunglasses on a beach"
    print(f"  Prompt: {prompt}")
    print(f"  Using OpenAI SDK client.images.generate()")
    print(f"  Sending request...")
    start = time.time()

    try:
        response = client.images.generate(
            model="dall-e-3",
            prompt=prompt,
            n=1,
            size="1024x1024",
            response_format="b64_json",
        )
    except Exception as e:
        print(f"  FAILED: {e}")
        return False

    elapsed = time.time() - start
    print(f"  Response received ({elapsed:.1f}s)")

    # Validate SDK response object
    assert response.data, "response.data is empty"
    assert len(response.data) >= 1, "Expected at least 1 image"

    img = response.data[0]
    assert img.b64_json, "Missing b64_json in response"

    img_bytes = base64.b64decode(img.b64_json)
    size_kb = len(img_bytes) / 1024
    print(f"  Image size: {size_kb:.1f} KB")

    if img.revised_prompt:
        print(f"  Revised prompt: {img.revised_prompt[:100]}")

    # Save it
    save_b64_image(img.b64_json, "test4_puppy_sunglasses.png")

    print(f"\n  PASSED ✓ ({elapsed:.1f}s)")
    return True


# ── Main ────────────────────────────────────────────────────────

def main():
    print("\n" + "=" * 70)
    print("  CatGPT — Image Generation API Test Suite")
    print("  Endpoint: POST /v1/images/generations")
    print(f"  Server:   {BASE_URL}")
    print("=" * 70)

    # Quick health check
    try:
        health = requests.get(f"{BASE_URL}/healthz", timeout=5)
        if health.status_code != 200:
            print(f"\n  ERROR: Server health check failed ({health.status_code})")
            sys.exit(1)
        print(f"\n  Health check: OK")
    except requests.ConnectionError:
        print(f"\n  ERROR: Cannot connect to {BASE_URL}")
        print("  Start the server: python -m src.api.server")
        print("  Or Docker: docker compose up --build -d catgpt")
        sys.exit(1)

    # Verify auth works
    try:
        models = requests.get(f"{BASE_URL}/v1/models", headers=HEADERS, timeout=5)
        if models.status_code == 401:
            print(f"  ERROR: Auth failed — check API_KEY (current: {API_KEY})")
            sys.exit(1)
        print(f"  Auth check: OK\n")
    except Exception:
        pass

    OUTPUT_DIR.mkdir(parents=True, exist_ok=True)

    results = {}
    tests = [
        ("Test 1: Basic b64_json generation", test_1_basic_image_generation),
        ("Test 2: Generate & save to disk", test_2_save_generated_image),
        ("Test 3: URL format response", test_3_url_format),
        ("Test 4: OpenAI SDK integration", test_4_openai_sdk),
    ]

    for name, test_fn in tests:
        try:
            passed = test_fn()
            results[name] = "PASSED" if passed else "FAILED"
        except Exception as e:
            print(f"\n  EXCEPTION: {e}")
            results[name] = "ERROR"
        # Brief pause between tests
        time.sleep(3)

    # ── Summary ─────────────────────────────────────────────
    separator("Test Results Summary")
    all_passed = True
    for name, result in results.items():
        icon = "✓" if result == "PASSED" else "✗"
        print(f"  {icon} {name}: {result}")
        if result != "PASSED":
            all_passed = False

    print()
    if all_passed:
        print("  All tests passed! ✓")
    else:
        print("  Some tests failed. Check output above for details.")

    # Show generated files
    if OUTPUT_DIR.exists():
        files = list(OUTPUT_DIR.glob("*.png"))
        if files:
            print(f"\n  Generated images saved to: {OUTPUT_DIR}")
            for f in sorted(files):
                print(f"    - {f.name} ({f.stat().st_size / 1024:.1f} KB)")

    print()
    return 0 if all_passed else 1


if __name__ == "__main__":
    sys.exit(main())
</file>

<file path="scripts/test_images.py">
#!/usr/bin/env python3
"""
Image Generation Test — validates image detection, extraction, and download.

Tests:
  1. Simple image generation (single image)
  2. Image with follow-up text question (mixed response check)
  3. Verify image files exist on disk
  4. Verify text-only response has no false image detection

Usage:
    python scripts/test_images.py
"""

from __future__ import annotations

import asyncio
import json
import os
import sys

sys.path.insert(0, os.path.dirname(os.path.dirname(os.path.abspath(__file__))))

from src.browser.manager import BrowserManager
from src.config import Config
from src.log import setup_logging

# Provider-aware client import
if Config.PROVIDER == "claude":
    from src.claude.client import ClaudeClient as ProviderClient
else:
    from src.chatgpt.client import ChatGPTClient as ProviderClient

log = setup_logging("test_images", log_file="test_images.log")


TESTS = [
    {
        "name": "Text-only (no false positive)",
        "prompt": "What is 2 + 2? Answer in one word.",
        "expect_image": False,
        "expect_contains": ["4", "four"],
        "expect_any": True,  # any of expect_contains matches = pass
    },
    {
        "name": "Image generation — simple",
        "prompt": "Generate an image of a cute orange tabby cat sitting on a windowsill with sunlight streaming in.",
        "expect_image": True,
    },
    {
        "name": "Follow-up text after image",
        "prompt": "Now describe the image you just created in 2 sentences.",
        "expect_image": False,
    },
    {
        "name": "Image generation — specific style",
        "prompt": "Generate an image of a futuristic cyberpunk city at night with neon lights and flying cars, digital art style.",
        "expect_image": True,
    },
]


def print_header(text: str) -> None:
    print(f"\n{'=' * 70}")
    print(f"  {text}")
    print(f"{'=' * 70}")


def validate(test: dict, response) -> tuple[bool, str]:
    """Validate test result. Returns (passed, reason)."""
    if test.get("expect_image"):
        if not response.has_images:
            return False, "Expected image but none detected"
        for img in response.images:
            if not img.local_path:
                return False, f"Image detected but download failed (url: {img.url[:60]})"
            if not os.path.exists(img.local_path):
                return False, f"Image file doesn't exist: {img.local_path}"
            size = os.path.getsize(img.local_path)
            if size < 1000:
                return False, f"Image file too small ({size} bytes): {img.local_path}"
        return True, f"{len(response.images)} image(s) downloaded successfully"

    else:
        if response.has_images:
            return False, "False positive — detected images in text-only response"
        if "expect_contains" in test:
            if test.get("expect_any"):
                # ANY of the keywords matching = pass
                found = any(
                    kw.lower() in response.message.lower()
                    for kw in test["expect_contains"]
                )
                if not found:
                    return False, f"Missing all expected text: {test['expect_contains']}"
            else:
                for kw in test["expect_contains"]:
                    if kw.lower() not in response.message.lower():
                        return False, f"Missing expected text: '{kw}'"
        return True, "Text response, no false image detection"


async def main():
    browser = BrowserManager()
    results = []

    try:
        print_header("Image Generation Test Suite")
        print(f"  Tests: {len(TESTS)}")
        print(f"  Images dir: {Config.IMAGES_DIR}")
        print()

        # Launch browser
        print("  Starting browser...")
        page = await browser.start()
        await browser.navigate(Config.provider_url())
        await asyncio.sleep(3)

        if not await browser.is_logged_in():
            print("\n  ❌ Not logged in! Run first_login.py first.")
            return

        print("  ✅ Logged in\n")
        client = ProviderClient(page)

        # Start fresh
        print("  Starting new chat...")
        await client.new_chat()

        for i, test in enumerate(TESTS, 1):
            # Skip image generation tests for Claude (no DALL-E equivalent)
            if Config.PROVIDER == "claude" and test.get("expect_image"):
                print(f"\n  {'─' * 60}")
                print(f"  Test {i}/{len(TESTS)}: {test['name']}")
                print(f"  ⏭️  SKIPPED — image generation not supported by Claude")
                results.append({
                    "test": test["name"],
                    "passed": True,
                    "reason": "Skipped — not supported by Claude",
                    "response_time_ms": 0,
                    "has_images": False,
                    "image_count": 0,
                    "response_length": 0,
                    "response_preview": "",
                })
                continue

            print(f"\n  {'─' * 60}")
            print(f"  Test {i}/{len(TESTS)}: {test['name']}")
            print(f"  Prompt: {test['prompt'][:65]}...")
            expects = "image" if test.get("expect_image") else "text only"
            print(f"  Expects: {expects}")
            print(f"  ⏳ Sending...")

            try:
                response = await client.send_message(test["prompt"])

                passed, reason = validate(test, response)
                status = "✅ PASS" if passed else "❌ FAIL"

                result = {
                    "test": test["name"],
                    "prompt": test["prompt"],
                    "passed": passed,
                    "reason": reason,
                    "response_time_ms": response.response_time_ms,
                    "has_images": response.has_images,
                    "image_count": len(response.images),
                    "response_length": len(response.message),
                    "response_preview": response.message[:150],
                }

                if response.images:
                    result["images"] = [
                        {
                            "alt": img.alt,
                            "local_path": img.local_path,
                            "url": img.url[:100],
                            "prompt_title": img.prompt_title,
                        }
                        for img in response.images
                    ]

                results.append(result)

                print(f"  {status} | {reason}")
                print(f"  Time: {response.response_time_ms}ms | Text: {len(response.message)} chars")

                if response.has_images:
                    for img in response.images:
                        print(f"  🎨 Image: {img.alt or img.prompt_title or 'untitled'}")
                        print(f"     Path:  {img.local_path}")
                        if img.local_path and os.path.exists(img.local_path):
                            size_kb = os.path.getsize(img.local_path) / 1024
                            print(f"     Size:  {size_kb:.1f} KB")
                else:
                    # Show text preview
                    preview = response.message[:100].replace("\n", " ")
                    print(f"  Text: {preview}...")

                log.info(f"Test '{test['name']}': {status} — {reason}")

            except Exception as e:
                results.append({
                    "test": test["name"],
                    "passed": False,
                    "reason": f"Error: {e}",
                })
                print(f"  ❌ ERROR: {e}")
                log.error(f"Test '{test['name']}' failed: {e}", exc_info=True)

        # ── Summary ─────────────────────────────────────────────
        print_header("TEST SUMMARY")

        passed = sum(1 for r in results if r["passed"])
        failed = len(results) - passed

        print(f"  Passed: {passed}/{len(results)}")
        print(f"  Failed: {failed}")
        print()

        for r in results:
            icon = "✅" if r["passed"] else "❌"
            time_ms = r.get("response_time_ms", 0)
            img_info = f", {r.get('image_count', 0)} img" if r.get("has_images") else ""
            print(f"  {icon} {r['test']}: {r['reason']} ({time_ms}ms{img_info})")

        # List downloaded images
        if Config.IMAGES_DIR.exists():
            images = list(Config.IMAGES_DIR.glob("*.*"))
            if images:
                print(f"\n  📁 Downloaded images ({len(images)} files):")
                for img_path in sorted(images, key=lambda p: p.stat().st_mtime, reverse=True)[:10]:
                    size_kb = img_path.stat().st_size / 1024
                    print(f"     {img_path.name} ({size_kb:.1f} KB)")

        if failed == 0:
            print(f"\n  🎉 ALL {len(results)} TESTS PASSED!")
        else:
            print(f"\n  ⚠️  {failed} test(s) failed — check logs.")

        # Save results
        results_file = Config.LOG_DIR / "image_test_results.json"
        with open(results_file, "w") as f:
            json.dump(results, f, indent=2, default=str)
        print(f"\n  📝 Results saved to: {results_file}")

        print("\n  Browser still open for inspection.")
        input("  Press ENTER to close > ")

    except KeyboardInterrupt:
        print("\n\n  Cancelled by user.")
    except Exception as e:
        log.error(f"Test failed: {e}", exc_info=True)
        print(f"\n  ❌ Fatal error: {e}")
    finally:
        await browser.close()
        print("  Browser closed.\n")


if __name__ == "__main__":
    asyncio.run(main())
</file>

<file path="scripts/test_langchain_tools.py">
#!/usr/bin/env python3
"""
LangChain test script for the OpenAI-compatible CatGPT API.

Tests:
  1. Simple chat (no tools)
  2. Tool/function calling with get_current_time()
  3. Tool calling with add_numbers()
  4. Complex multi-tool scenarios
  5. Image input (single image + text, multiple images)

Prerequisites:
  - CatGPT API server running: python -m src.api.server
  - pip install langchain langchain-openai openai

Usage:
  python scripts/test_langchain_tools.py
"""

from __future__ import annotations

import base64
import json
import os
import sys
from datetime import datetime
from pathlib import Path

# Load .env if dotenv is available
try:
    from dotenv import load_dotenv
    load_dotenv(Path(__file__).resolve().parent.parent / ".env")
except ImportError:
    pass

from langchain_openai import ChatOpenAI
from langchain_core.messages import HumanMessage, AIMessage, ToolMessage
from langchain_core.tools import tool


# ── Configuration ───────────────────────────────────────────────

BASE_URL = "http://localhost:8000/v1"
# Auto-detect model from provider env var
_provider = os.environ.get("PROVIDER", "chatgpt")
MODEL = "claude-browser" if _provider == "claude" else "catgpt-browser"
API_KEY = "dummy123"  # CatGPT doesn't require auth

# Image test assets
IMAGE_DIR = Path(__file__).resolve().parent.parent / "downloads" / "images"


# ── Tools ───────────────────────────────────────────────────────

@tool
def get_current_time() -> str:
    """Get the current date and time. Returns ISO format datetime string."""
    return datetime.now().strftime("%Y-%m-%d %H:%M:%S")

@tool
def add_numbers(a: int, b: int) -> str:
    """Add two numbers together and return the result."""
    return str(a + b)

@tool
def search_wikipedia(query: str) -> str:
    """Return a fake Wikipedia summary for the query."""
    return f"[Wikipedia summary for '{query}': Lorem ipsum dolor sit amet...]"

@tool
def calculate_expression(expression: str) -> str:
    """Evaluate a math expression (e.g., '7*8+3')."""
    try:
        return str(eval(expression, {"__builtins__": {}}))
    except Exception as e:
        return f"Error: {e}"

@tool
def reverse_string(s: str) -> str:
    """Reverse the input string."""
    return s[::-1]

@tool
def weather_forecast(city: str, date: str) -> str:
    """Return a fake weather forecast for a city and date."""
    return f"The weather in {city} on {date} will be sunny with a high of 25°C."

@tool
def multi_arg_tool(a: int, b: int, c: int) -> str:
    """Return the product of three numbers."""
    return str(a * b * c)


# ── Helpers ─────────────────────────────────────────────────────

def separator(title: str) -> None:
    print(f"\n{'=' * 60}")
    print(f"  {title}")
    print(f"{'=' * 60}\n")


def _image_to_data_url(path: str | Path) -> str:
    """Read a local image and return an OpenAI-compatible base64 data URL."""
    path = Path(path)
    ext = path.suffix.lower().lstrip(".")
    mime = {"png": "image/png", "jpg": "image/jpeg", "jpeg": "image/jpeg",
            "webp": "image/webp", "gif": "image/gif"}.get(ext, "image/png")
    data = base64.b64encode(path.read_bytes()).decode()
    return f"data:{mime};base64,{data}"


def _file_to_base64(path: str | Path) -> str:
    """Read any file and return its base64-encoded content."""
    path = Path(path)
    return base64.b64encode(path.read_bytes()).decode()


def _find_test_images(n: int = 2) -> list[Path]:
    """Return up to *n* image files from IMAGE_DIR."""
    if not IMAGE_DIR.exists():
        return []
    exts = {".png", ".jpg", ".jpeg", ".webp", ".gif"}
    imgs = sorted(p for p in IMAGE_DIR.iterdir() if p.suffix.lower() in exts)
    return imgs[:n]


def _find_test_files(extensions: set[str] | None = None) -> list[Path]:
    """Return test files (non-image) from the downloads directory."""
    dl_dir = Path(__file__).resolve().parent.parent / "downloads"
    if not dl_dir.exists():
        return []
    if extensions is None:
        extensions = {".pdf", ".txt", ".csv", ".docx", ".xlsx", ".json"}
    return sorted(p for p in dl_dir.iterdir() if p.suffix.lower() in extensions and p.is_file())


def test_models_endpoint():
    """Test that /v1/models returns our model."""
    import openai

    client = openai.OpenAI(base_url=BASE_URL, api_key=API_KEY)
    models = client.models.list()
    model_ids = [m.id for m in models.data]
    print(f"Available models: {model_ids}")
    assert MODEL in model_ids, f"Expected {MODEL} in models list"
    print("✓ Models endpoint works\n")


def test_simple_chat():
    """Test a simple chat without tools."""
    separator("Test 1: Simple Chat (no tools)")

    llm = ChatOpenAI(
        model=MODEL,
        base_url=BASE_URL,
        api_key=API_KEY,
        temperature=0,
    )

    response = llm.invoke([HumanMessage(content="Who is the president of the United States?")])
    print(f"Question: Who is the president of the United States?")
    print(f"Response: {response.content}")
    print(f"Type: {type(response).__name__}")
    print("✓ Simple chat works\n")


def test_tool_calling():
    """Test tool/function calling with get_current_time."""
    separator("Test 2: Tool Calling (get_current_time)")

    llm = ChatOpenAI(
        model=MODEL,
        base_url=BASE_URL,
        api_key=API_KEY,
        temperature=0,
    )

    tools = [get_current_time, add_numbers]
    llm_with_tools = llm.bind_tools(tools)

    print("Sending: What is the current time? Use the get_current_time tool.")
    response = llm_with_tools.invoke(
        [HumanMessage(content="What is the current time? Use the get_current_time tool.")]
    )

    print(f"Response type: {type(response).__name__}")
    print(f"Content: {response.content}")
    print(f"Tool calls: {response.tool_calls}")

    if response.tool_calls:
        print(f"\n✓ Model requested tool call(s):")
        for tc in response.tool_calls:
            print(f"  - {tc['name']}({tc['args']})")

        messages = [
            HumanMessage(content="What is the current time? Use the get_current_time tool."),
            response,
        ]

        tool_map = {"get_current_time": get_current_time, "add_numbers": add_numbers}

        for tc in response.tool_calls:
            tool_fn = tool_map.get(tc["name"])
            if tool_fn:
                result = tool_fn.invoke(tc["args"])
                print(f"  Tool result ({tc['name']}): {result}")
                messages.append(
                    ToolMessage(content=str(result), tool_call_id=tc["id"])
                )

        print("\nSending tool results back to model...")
        final_response = llm_with_tools.invoke(messages)
        print(f"Final response: {final_response.content}")

        print("\nPrompting for summary after tool result...")
        messages.append(HumanMessage(content="Please summarize the result in a sentence."))
        summary_response = llm_with_tools.invoke(messages)
        print(f"Summary: {summary_response.content}")
        print("✓ Tool calling round-trip with summary works\n")
    else:
        print("⚠ Model did not request tool calls (responded directly)")
        print("✓ Test completed (no tool calls)\n")


def test_add_numbers_tool():
    """Test tool calling with add_numbers."""
    separator("Test 3: Tool Calling (add_numbers)")

    llm = ChatOpenAI(
        model=MODEL,
        base_url=BASE_URL,
        api_key=API_KEY,
        temperature=0,
    )

    tools = [add_numbers]
    llm_with_tools = llm.bind_tools(tools)

    print("Sending: Use the add_numbers tool to compute 42 + 58.")
    response = llm_with_tools.invoke(
        [HumanMessage(content="Use the add_numbers tool to compute 42 + 58.")]
    )

    print(f"Content: {response.content}")
    print(f"Tool calls: {response.tool_calls}")

    if response.tool_calls:
        print(f"\n✓ Model requested tool call(s):")
        for tc in response.tool_calls:
            print(f"  - {tc['name']}({tc['args']})")

        messages = [
            HumanMessage(content="Use the add_numbers tool to compute 42 + 58."),
            response,
        ]

        for tc in response.tool_calls:
            if tc["name"] == "add_numbers":
                result = add_numbers.invoke(tc["args"])
                print(f"  Tool result: {result}")
                messages.append(
                    ToolMessage(content=str(result), tool_call_id=tc["id"])
                )

        print("\nSending tool results back...")
        final = llm_with_tools.invoke(messages)
        print(f"Final response: {final.content}")

        print("\nPrompting for summary after tool result...")
        messages.append(HumanMessage(content="Please summarize the result in a sentence."))
        summary_response = llm_with_tools.invoke(messages)
        print(f"Summary: {summary_response.content}")
        print("✓ add_numbers tool calling with summary works\n")
    else:
        print("⚠ Model did not use tool (answered directly)")
        print("✓ Test completed\n")


def test_complex_tool_calls():
    """Test complex multi-tool scenarios."""
    separator("Test 4: Complex Tool Calling")

    llm = ChatOpenAI(
        model=MODEL,
        base_url=BASE_URL,
        api_key=API_KEY,
        temperature=0,
    )

    all_tools = [
        get_current_time, add_numbers, search_wikipedia,
        calculate_expression, reverse_string, weather_forecast, multi_arg_tool,
    ]
    llm_with_tools = llm.bind_tools(all_tools)
    tool_map = {t.name: t for t in all_tools}

    # --- Sub-test A: weather + math ---
    print("Sending: What is the weather in Paris tomorrow, and what is 7*8+3?")
    response = llm_with_tools.invoke([
        HumanMessage(content="What is the weather in Paris tomorrow, and what is 7*8+3?")
    ])
    print(f"Content: {response.content}")
    print(f"Tool calls: {response.tool_calls}")

    messages = [HumanMessage(content="What is the weather in Paris tomorrow, and what is 7*8+3?")]
    if response.tool_calls:
        for tc in response.tool_calls:
            tool_fn = tool_map.get(tc["name"])
            if tool_fn:
                result = tool_fn.invoke(tc["args"])
                print(f"  Tool result ({tc['name']}): {result}")
                messages.append(ToolMessage(content=str(result), tool_call_id=tc["id"]))

        print("\nSending tool results back to model...")
        final = llm_with_tools.invoke(messages)
        print(f"Final response: {final.content}")

        print("\nPrompting for summary...")
        messages.append(HumanMessage(content="Please summarize the results in a sentence."))
        summary = llm_with_tools.invoke(messages)
        print(f"Summary: {summary.content}")
    else:
        print("⚠ Model did not use tools (answered directly)")
    print("✓ Complex tool calling test completed\n")

    # --- Sub-test B: string reversal + Wikipedia ---
    print("Sending: Reverse the string 'OpenAI', and search Wikipedia for 'LangChain'.")
    response2 = llm_with_tools.invoke([
        HumanMessage(content="Reverse the string 'OpenAI', and search Wikipedia for 'LangChain'.")
    ])
    print(f"Content: {response2.content}")
    print(f"Tool calls: {response2.tool_calls}")

    messages2 = [HumanMessage(content="Reverse the string 'OpenAI', and search Wikipedia for 'LangChain'.")]
    if response2.tool_calls:
        for tc in response2.tool_calls:
            tool_fn = tool_map.get(tc["name"])
            if tool_fn:
                result = tool_fn.invoke(tc["args"])
                print(f"  Tool result ({tc['name']}): {result}")
                messages2.append(ToolMessage(content=str(result), tool_call_id=tc["id"]))

        print("\nSending tool results back to model...")
        final2 = llm_with_tools.invoke(messages2)
        print(f"Final response: {final2.content}")

        print("\nPrompting for summary...")
        messages2.append(HumanMessage(content="Please summarize the results in a sentence."))
        summary2 = llm_with_tools.invoke(messages2)
        print(f"Summary: {summary2.content}")
    else:
        print("⚠ Model did not use tools (answered directly)")
    print("✓ Multi-tool call test completed\n")


def test_image_input():
    """Test sending images via the OpenAI vision content format."""
    separator("Test 5: Image Input")

    images = _find_test_images(2)
    if not images:
        print(f"⚠ No images found in {IMAGE_DIR} — skipping image tests")
        print("  Place .png/.jpg/.webp files in downloads/images/ to enable this test")
        print("✓ Image test skipped (no assets)\n")
        return

    import openai
    client = openai.OpenAI(base_url=BASE_URL, api_key=API_KEY)

    # --- 5a: Single image + text ---
    img_path = images[0]
    data_url = _image_to_data_url(img_path)
    print(f"5a — Sending single image: {img_path.name}  ({img_path.stat().st_size / 1024:.0f} KB)")

    resp = client.chat.completions.create(
        model=MODEL,
        messages=[
            {
                "role": "user",
                "content": [
                    {"type": "text", "text": "Describe this image in 2-3 sentences."},
                    {"type": "image_url", "image_url": {"url": data_url}},
                ],
            }
        ],
    )
    answer = resp.choices[0].message.content
    print(f"Response: {answer}")
    assert answer and len(answer) > 10, "Expected a non-trivial description"
    print("✓ Single image + text works\n")

    # --- 5b: Multiple images + text ---
    if len(images) >= 2:
        img2_path = images[1]
        data_url2 = _image_to_data_url(img2_path)
        print(f"5b — Sending two images: {img_path.name}, {img2_path.name}")

        resp2 = client.chat.completions.create(
            model=MODEL,
            messages=[
                {
                    "role": "user",
                    "content": [
                        {"type": "text", "text": "Compare these two images. What are the main differences?"},
                        {"type": "image_url", "image_url": {"url": data_url}},
                        {"type": "image_url", "image_url": {"url": data_url2}},
                    ],
                }
            ],
        )
        answer2 = resp2.choices[0].message.content
        print(f"Response: {answer2}")
        assert answer2 and len(answer2) > 10, "Expected a non-trivial comparison"
        print("✓ Multiple images + text works\n")
    else:
        print("5b — Only one image available, skipping multi-image test")
        print("✓ Multi-image test skipped\n")

    # --- 5c: Image with tool calling ---
    print("5c — Image + tool calling (describe image, then use add_numbers)")
    llm = ChatOpenAI(
        model=MODEL,
        base_url=BASE_URL,
        api_key=API_KEY,
        temperature=0,
    )
    llm_with_tools = llm.bind_tools([add_numbers])

    response = llm_with_tools.invoke([
        HumanMessage(content=[
            {"type": "text", "text": "Look at this image and then use the add_numbers tool to add 10 + 20."},
            {"type": "image_url", "image_url": {"url": data_url}},
        ])
    ])
    print(f"Content: {response.content}")
    print(f"Tool calls: {response.tool_calls}")

    if response.tool_calls:
        for tc in response.tool_calls:
            print(f"  Tool call: {tc['name']}({tc['args']})")
            result = add_numbers.invoke(tc["args"])
            print(f"  Result: {result}")
    print("✓ Image + tool calling test completed\n")


def test_file_attachment():
    """Test sending non-image file attachments (e.g. PDF)."""
    separator("Test 6: File Attachment (PDF)")

    test_files = _find_test_files({".pdf"})
    if not test_files:
        print("⚠ No PDF files found in downloads/ — skipping file attachment test")
        print("  Place a .pdf file in downloads/ to enable this test")
        print("✓ File attachment test skipped (no assets)\n")
        return

    import openai
    client = openai.OpenAI(base_url=BASE_URL, api_key=API_KEY)

    # --- 6a: PDF attachment + text question ---
    pdf_path = test_files[0]
    pdf_b64 = _file_to_base64(pdf_path)
    print(f"6a — Sending PDF: {pdf_path.name} ({pdf_path.stat().st_size / 1024:.0f} KB)")

    resp = client.chat.completions.create(
        model=MODEL,
        messages=[
            {
                "role": "user",
                "content": [
                    {"type": "text", "text": "I've attached a PDF file. Please summarize its contents in 2-3 sentences."},
                    {
                        "type": "file",
                        "file": {
                            "filename": pdf_path.name,
                            "data": pdf_b64,
                            "mime_type": "application/pdf",
                        },
                    },
                ],
            }
        ],
    )
    answer = resp.choices[0].message.content
    print(f"Response: {answer}")
    assert answer and len(answer) > 10, "Expected a non-trivial summary of the PDF"
    print("✓ PDF attachment + text works\n")

    # --- 6b: PDF + image combined ---
    images = _find_test_images(1)
    if images:
        img_path = images[0]
        data_url = _image_to_data_url(img_path)
        print(f"6b — Sending PDF ({pdf_path.name}) + image ({img_path.name}) together")

        resp2 = client.chat.completions.create(
            model=MODEL,
            messages=[
                {
                    "role": "user",
                    "content": [
                        {"type": "text", "text": "I've attached a PDF and an image. Briefly describe what each contains."},
                        {
                            "type": "file",
                            "file": {
                                "filename": pdf_path.name,
                                "data": pdf_b64,
                                "mime_type": "application/pdf",
                            },
                        },
                        {"type": "image_url", "image_url": {"url": data_url}},
                    ],
                }
            ],
        )
        answer2 = resp2.choices[0].message.content
        print(f"Response: {answer2}")
        assert answer2 and len(answer2) > 10, "Expected a non-trivial response"
        print("✓ PDF + image combined works\n")
    else:
        print("6b — No images available, skipping PDF + image combined test\n")


# ── Main ────────────────────────────────────────────────────────

def main():
    print("=" * 60)
    print("  CatGPT — LangChain OpenAI-Compatible API Tests")
    print("=" * 60)
    print(f"\nBase URL: {BASE_URL}")
    print(f"Model:    {MODEL}")
    print(f"Image dir: {IMAGE_DIR}\n")

    try:
        test_models_endpoint()
        test_simple_chat()
        test_tool_calling()
        test_add_numbers_tool()
        test_complex_tool_calls()
        test_image_input()
        test_file_attachment()

        separator("All Tests Passed!")
        print("The OpenAI-compatible API is working correctly with LangChain.")
        print("Tool/function calling, image input, and file attachments are operational.\n")

    except Exception as e:
        print(f"\n✗ Test failed: {e}")
        import traceback
        traceback.print_exc()
        sys.exit(1)


if __name__ == "__main__":
    main()
</file>

<file path="scripts/test_multi_turn.py">
#!/usr/bin/env python3
"""
Multi-Turn Test — Validates conversation continuity and new chat functionality.

Test Plan:
  Round 1: Start a NEW CHAT, send 5 follow-up messages in the same thread
  Round 2: Start ANOTHER NEW CHAT, send 5 more follow-up messages

Each message response is verified and logged. The test only confirms
completion after the response has fully streamed.

Usage:
    python scripts/test_multi_turn.py
"""

from __future__ import annotations

import asyncio
import json
import sys
import os
import time

sys.path.insert(0, os.path.dirname(os.path.dirname(os.path.abspath(__file__))))

from src.browser.manager import BrowserManager
from src.config import Config
from src.log import setup_logging

# Provider-aware client import
if Config.PROVIDER == "claude":
    from src.claude.client import ClaudeClient as ProviderClient
else:
    from src.chatgpt.client import ChatGPTClient as ProviderClient

log = setup_logging("test_multi_turn", log_file="test_multi_turn.log")

# ── Test Messages ───────────────────────────────────────────────

ROUND_1_MESSAGES = [
    "Hi! Let's do a quick test. Please respond with only: 'Round 1, Message 1 received.' Nothing else.",
    "Great. Now respond with only: 'Round 1, Message 2 received.' Nothing else.",
    "Perfect. Now respond with only: 'Round 1, Message 3 received.' Nothing else.",
    "Good. Now respond with only: 'Round 1, Message 4 received.' Nothing else.",
    "Last one. Respond with only: 'Round 1, Message 5 received.' Nothing else.",
]

ROUND_2_MESSAGES = [
    "Hello again! This is a new conversation. Respond with only: 'Round 2, Message 1 received.' Nothing else.",
    "Continue here. Respond with only: 'Round 2, Message 2 received.' Nothing else.",
    "Still going. Respond with only: 'Round 2, Message 3 received.' Nothing else.",
    "Almost done. Respond with only: 'Round 2, Message 4 received.' Nothing else.",
    "Final message. Respond with only: 'Round 2, Message 5 received.' Nothing else.",
]


def print_header(text: str) -> None:
    print(f"\n{'=' * 60}")
    print(f"  {text}")
    print(f"{'=' * 60}\n")


def print_result(idx: int, msg: str, response: str, time_ms: int, thread_id: str) -> None:
    status = "✅" if response.strip() else "❌"
    print(f"  {status} [{idx}] ({time_ms}ms) Thread: {thread_id or 'n/a'}")
    print(f"     Sent: {msg[:60]}...")
    print(f"     Got:  {response[:80]}")
    print()


async def run_round(
    client: ProviderClient,
    round_name: str,
    messages: list[str],
    results: list[dict],
) -> str:
    """Send a series of messages and collect results. Returns thread_id."""
    print_header(f"{round_name} — {len(messages)} messages")

    thread_id = ""
    for i, msg in enumerate(messages, 1):
        print(f"  ⏳ [{round_name}] Sending message {i}/{len(messages)}...")
        log.info(f"[{round_name}] Sending message {i}: {msg[:60]}")

        try:
            response = await client.send_message(msg)
            thread_id = response.thread_id or thread_id

            result = {
                "round": round_name,
                "message_num": i,
                "sent": msg,
                "response": response.message,
                "response_time_ms": response.response_time_ms,
                "thread_id": thread_id,
                "status": "ok" if response.message.strip() else "empty",
            }
            results.append(result)

            print_result(i, msg, response.message, response.response_time_ms, thread_id)
            log.info(
                f"[{round_name}] Message {i} OK — {response.response_time_ms}ms, "
                f"{len(response.message)} chars: {response.message[:60]}"
            )

        except Exception as e:
            result = {
                "round": round_name,
                "message_num": i,
                "sent": msg,
                "response": "",
                "response_time_ms": 0,
                "thread_id": thread_id,
                "status": f"error: {e}",
            }
            results.append(result)
            print(f"  ❌ [{i}] ERROR: {e}")
            log.error(f"[{round_name}] Message {i} FAILED: {e}", exc_info=True)

    return thread_id


async def main():
    browser = BrowserManager()
    results: list[dict] = []

    try:
        print_header(f"Multi-Turn Test — {Config.PROVIDER.title()} Automation")
        print("  This test will:")
        print("  1. Open a NEW chat and send 5 follow-up messages")
        print("  2. Open ANOTHER new chat and send 5 more messages")
        print("  3. Report all results\n")

        # Launch browser
        print("  Starting browser...")
        page = await browser.start()
        await browser.navigate(Config.provider_url())
        await asyncio.sleep(3)

        # Verify login
        if not await browser.is_logged_in():
            print("\n  ❌ Not logged in! Run 'python scripts/first_login.py' first.")
            return

        print("  ✅ Logged in\n")
        client = ProviderClient(page)

        # ── Round 1: New chat + 5 messages ──────────────────────
        print("  Starting new chat for Round 1...")
        await client.new_chat()

        round1_thread = await run_round(client, "Round 1", ROUND_1_MESSAGES, results)

        print(f"  Round 1 complete — Thread: {round1_thread}")
        print(f"  {'─' * 50}")

        # ── Round 2: Another new chat + 5 messages ──────────────
        print("\n  Starting new chat for Round 2...")
        await client.new_chat()

        round2_thread = await run_round(client, "Round 2", ROUND_2_MESSAGES, results)

        print(f"  Round 2 complete — Thread: {round2_thread}")

        # ── Summary ─────────────────────────────────────────────
        print_header("TEST SUMMARY")

        ok_count = sum(1 for r in results if r["status"] == "ok")
        fail_count = len(results) - ok_count
        total_time = sum(r["response_time_ms"] for r in results)

        print(f"  Total messages:  {len(results)}")
        print(f"  Successful:      {ok_count}")
        print(f"  Failed:          {fail_count}")
        print(f"  Total time:      {total_time}ms ({total_time / 1000:.1f}s)")
        print(f"  Avg response:    {total_time // len(results)}ms")
        print(f"  Round 1 thread:  {round1_thread or 'n/a'}")
        print(f"  Round 2 thread:  {round2_thread or 'n/a'}")
        print(f"  Threads differ:  {'✅ Yes' if round1_thread != round2_thread else '❌ No (same thread!)'}")
        print()

        if fail_count == 0:
            print("  🎉 ALL TESTS PASSED — Ready for Phase 2!")
        else:
            print(f"  ⚠️  {fail_count} message(s) had issues — check logs.")

        # Save results
        results_file = Config.LOG_DIR / "multi_turn_results.json"
        with open(results_file, "w") as f:
            json.dump({
                "summary": {
                    "total": len(results),
                    "ok": ok_count,
                    "failed": fail_count,
                    "total_time_ms": total_time,
                    "round1_thread": round1_thread,
                    "round2_thread": round2_thread,
                },
                "results": results,
            }, f, indent=2, default=str)
        print(f"\n  📝 Results saved to: {results_file}")

        # Keep browser open for inspection
        print("\n  Browser still open for inspection.")
        input("  Press ENTER to close > ")

    except KeyboardInterrupt:
        print("\n\n  Cancelled by user.")
    except Exception as e:
        log.error(f"Test failed: {e}", exc_info=True)
        print(f"\n  ❌ Fatal error: {e}")
        print("  Check logs/test_multi_turn.log for details.")
    finally:
        await browser.close()
        print("  Browser closed.\n")


if __name__ == "__main__":
    asyncio.run(main())
</file>

<file path="scripts/test_phase1.py">
#!/usr/bin/env python3
"""
Phase 1 Test Script — Validates the full pipeline.

Steps:
1. Launch browser with persistent session
2. Check login status
3. Start DOM observer + network recorder
4. Optionally start a new chat
5. Send a test message
6. Wait for and capture the response
7. Log everything and print results

Usage:
    python scripts/test_phase1.py
    python scripts/test_phase1.py --message "What is 2+2?"
    python scripts/test_phase1.py --new-chat
"""

import argparse
import asyncio
import json
import sys
import os
import time

# Add project root to path
sys.path.insert(0, os.path.dirname(os.path.dirname(os.path.abspath(__file__))))

from src.browser.manager import BrowserManager
from src.dom_observer import DOMObserver
from src.network_recorder import NetworkRecorder
from src.config import Config
from src.log import setup_logging

# Provider-aware client import
if Config.PROVIDER == "claude":
    from src.claude.client import ClaudeClient as ProviderClient
else:
    from src.chatgpt.client import ChatGPTClient as ProviderClient

log = setup_logging("test_phase1", log_file="test_phase1.log")

DEFAULT_MESSAGE = "Hello! Please respond with exactly: 'Phase 1 test successful.' Nothing else."


async def main(message: str, new_chat: bool, observe: bool):
    browser = BrowserManager()

    try:
        print("\n" + "=" * 60)
        print(f"  Phase 1 Test — {Config.PROVIDER.title()} Browser Automation")
        print("=" * 60)

        # 1. Launch browser
        print("\n  [1/7] Launching browser...")
        page = await browser.start()

        # 2. Navigate to provider
        print(f"  [2/7] Navigating to {Config.PROVIDER.title()}...")
        await browser.navigate(Config.provider_url())

        # Give page time to fully load
        await asyncio.sleep(3)

        # 3. Check login
        print("  [3/7] Checking login status...")
        logged_in = await browser.is_logged_in()
        if not logged_in:
            print("\n  ❌ Not logged in! Run 'python scripts/first_login.py' first.")
            log.error("Not logged in — aborting test")
            return

        print("  ✅ Logged in")
        log.info("Login confirmed")

        # 4. Start observers (for Phase 1 observation)
        dom_obs = DOMObserver(page)
        net_rec = NetworkRecorder(page)

        if observe:
            print("  [4/7] Starting observers...")
            await dom_obs.start()
            net_rec.start()
            log.info("Observers started")
        else:
            print("  [4/7] Observers skipped (use --observe to enable)")

        # 5. Optionally start new chat
        client = ProviderClient(page)

        if new_chat:
            print("  [5/7] Starting new chat...")
            await client.new_chat()
            log.info("New chat started")
        else:
            print("  [5/7] Using current chat (use --new-chat to start fresh)")

        # 6. Send test message
        print(f"  [6/7] Sending message: {message[:60]}...")
        log.info(f"Sending test message: {message}")

        start_time = time.time()
        response = await client.send_message(message)
        elapsed = time.time() - start_time

        # 7. Results
        print("\n" + "-" * 60)
        print("  RESULTS")
        print("-" * 60)
        print(f"\n  Thread ID: {response.thread_id or '(not in URL yet)'}")
        print(f"  Response time: {response.response_time_ms}ms ({elapsed:.1f}s)")
        print(f"  Response length: {len(response.message)} chars")
        print(f"\n  Response:")
        print(f"  {'─' * 50}")
        # Indent response for readability
        for line in response.message.split("\n"):
            print(f"  {line}")
        print(f"  {'─' * 50}")

        log.info(f"Test complete — {response.response_time_ms}ms, {len(response.message)} chars")
        log.info(f"Response: {response.message}")

        # Log network activity
        if observe:
            net_rec.stop()
            await dom_obs.stop()
            captured = net_rec.get_captured()
            log.info(f"Network requests captured: {len(captured)}")
            for req in captured:
                log.debug(f"  {req['method']} {req['url'][:100]}")

            # Save observations to file
            obs_file = Config.LOG_DIR / "phase1_observations.json"
            observations = {
                "test_message": message,
                "response": response.message,
                "response_time_ms": response.response_time_ms,
                "thread_id": response.thread_id,
                "thread_url": await client.get_current_thread_url(),
                "network_requests": captured,
                "selectors_used": {
                    "chat_input": "see logs",
                    "send_button": "see logs",
                    "assistant_response": "see logs",
                },
            }
            with open(obs_file, "w") as f:
                json.dump(observations, f, indent=2, default=str)
            print(f"\n  📝 Observations saved to: {obs_file}")

        # Keep browser open for manual inspection
        print("\n  Browser still open for inspection.")
        input("  Press ENTER to close > ")

    except KeyboardInterrupt:
        print("\n\n  Cancelled by user.")
    except Exception as e:
        log.error(f"Test failed: {e}", exc_info=True)
        print(f"\n  ❌ Error: {e}")
        print("  Check logs for details.")
    finally:
        await browser.close()
        print("  Browser closed.\n")


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Phase 1 Test — Provider Automation")
    parser.add_argument(
        "--message", "-m",
        default=DEFAULT_MESSAGE,
        help="Message to send (default: test prompt)",
    )
    parser.add_argument(
        "--new-chat", "-n",
        action="store_true",
        help="Start a new chat before sending",
    )
    parser.add_argument(
        "--observe", "-o",
        action="store_true",
        help="Enable DOM observer + network recorder",
    )
    args = parser.parse_args()

    asyncio.run(main(args.message, args.new_chat, args.observe))
</file>

<file path="scripts/test_robust.py">
#!/usr/bin/env python3
"""
Robustness Test — Tests long responses, tables, code blocks, and markdown.

Validates:
  1. Long-form responses complete correctly (not cut off)
  2. Tables are captured properly via copy button
  3. Code blocks are preserved
  4. New chat isolation works
  5. Multiple follow-ups with complex content

Usage:
    python scripts/test_robust.py
"""

from __future__ import annotations

import asyncio
import json
import sys
import os

sys.path.insert(0, os.path.dirname(os.path.dirname(os.path.abspath(__file__))))

from src.browser.manager import BrowserManager
from src.config import Config
from src.log import setup_logging

# Provider-aware client import
if Config.PROVIDER == "claude":
    from src.claude.client import ClaudeClient as ProviderClient
else:
    from src.chatgpt.client import ChatGPTClient as ProviderClient

log = setup_logging("test_robust", log_file="test_robust.log")

# ── Test prompts designed to produce complex output ─────────────

TESTS = [
    {
        "name": "Table output",
        "prompt": (
            "Create a markdown table with 5 rows comparing Python, JavaScript, "
            "and Rust. Columns: Language, Typing, Speed, Use Case, Year Created. "
            "Only output the table, nothing else."
        ),
        "expect_contains": ["|" if Config.PROVIDER != "claude" else "\t", "Python", "JavaScript", "Rust"],
    },
    {
        "name": "Code block",
        "prompt": (
            "Write a Python function called 'fibonacci' that returns the first N "
            "fibonacci numbers as a list. Include a docstring. Only output the code "
            "block, nothing else."
        ),
        "expect_contains": ["def fibonacci", "return"],
    },
    {
        "name": "Long response",
        "prompt": (
            "Explain the TCP/IP model in detail. Cover all 4 layers, their protocols, "
            "and how data flows from application to physical layer. Be thorough — "
            "at least 300 words."
        ),
        "expect_min_length": 300,
    },
    {
        "name": "Follow-up context",
        "prompt": (
            "Based on what you just explained about TCP/IP, which layer would be "
            "most relevant for a cybersecurity engineer doing packet analysis? "
            "Answer in 2-3 sentences."
        ),
        "expect_contains": ["layer"],
    },
    {
        "name": "Mixed content (table + explanation)",
        "prompt": (
            "Create a table of the OSI model (all 7 layers) with columns: "
            "Layer Number, Name, Protocol Examples, PDU. Then below the table, "
            "write one sentence explaining why this model matters."
        ),
        "expect_contains": ["|" if Config.PROVIDER != "claude" else "\t", "Physical", "Application"],
    },
]


def print_header(text: str) -> None:
    print(f"\n{'=' * 70}")
    print(f"  {text}")
    print(f"{'=' * 70}")


def validate_response(test: dict, response: str) -> tuple[bool, str]:
    """Check if the response meets expectations. Returns (passed, reason)."""
    if not response.strip():
        return False, "Empty response"

    if "expect_contains" in test:
        for keyword in test["expect_contains"]:
            if keyword.lower() not in response.lower():
                return False, f"Missing expected content: '{keyword}'"

    if "expect_min_length" in test:
        # Count words instead of chars for more meaningful length check
        word_count = len(response.split())
        if word_count < test["expect_min_length"] * 0.5:  # Allow some flexibility
            return False, f"Too short: {word_count} words (expected ~{test['expect_min_length']}+)"

    return True, "All checks passed"


async def main():
    browser = BrowserManager()
    results = []

    try:
        print_header("Robustness Test — Complex Content Extraction")
        print(f"  Tests: {len(TESTS)}")
        print("  Focus: tables, code blocks, long responses, markdown")
        print()

        # Launch browser
        print("  Starting browser...")
        page = await browser.start()
        await browser.navigate(Config.provider_url())
        await asyncio.sleep(3)

        if not await browser.is_logged_in():
            print("\n  ❌ Not logged in! Run 'python scripts/first_login.py' first.")
            return

        print("  ✅ Logged in\n")
        client = ProviderClient(page)

        # Start a fresh chat
        print("  Starting new chat...")
        await client.new_chat()

        for i, test in enumerate(TESTS, 1):
            print(f"\n  {'─' * 60}")
            print(f"  Test {i}/{len(TESTS)}: {test['name']}")
            print(f"  Prompt: {test['prompt'][:70]}...")
            print(f"  ⏳ Sending...")

            try:
                response = await client.send_message(test["prompt"])

                passed, reason = validate_response(test, response.message)
                status = "✅ PASS" if passed else "❌ FAIL"

                result = {
                    "test": test["name"],
                    "prompt": test["prompt"],
                    "response_length": len(response.message),
                    "response_time_ms": response.response_time_ms,
                    "thread_id": response.thread_id,
                    "passed": passed,
                    "reason": reason,
                    "response_preview": response.message[:200],
                    "full_response": response.message,
                }
                results.append(result)

                print(f"  {status} | {reason}")
                print(f"  Response: {response.response_time_ms}ms, {len(response.message)} chars")

                # Show first 3 lines of response
                lines = response.message.split("\n")[:5]
                for line in lines:
                    print(f"    │ {line[:80]}")
                if len(response.message.split("\n")) > 5:
                    print(f"    │ ... ({len(response.message.split(chr(10)))} total lines)")

                log.info(
                    f"Test '{test['name']}': {status} — "
                    f"{response.response_time_ms}ms, {len(response.message)} chars"
                )

            except Exception as e:
                results.append({
                    "test": test["name"],
                    "passed": False,
                    "reason": f"Error: {e}",
                    "response_length": 0,
                })
                print(f"  ❌ ERROR: {e}")
                log.error(f"Test '{test['name']}' failed: {e}", exc_info=True)

        # ── Summary ─────────────────────────────────────────────
        print_header("TEST SUMMARY")

        passed = sum(1 for r in results if r["passed"])
        failed = len(results) - passed
        total_time = sum(r.get("response_time_ms", 0) for r in results)

        print(f"  Passed:    {passed}/{len(results)}")
        print(f"  Failed:    {failed}")
        print(f"  Total time: {total_time}ms ({total_time / 1000:.1f}s)")
        print()

        for r in results:
            icon = "✅" if r["passed"] else "❌"
            print(f"  {icon} {r['test']}: {r['reason']} ({r.get('response_time_ms', 0)}ms, {r['response_length']} chars)")

        if failed == 0:
            print(f"\n  🎉 ALL {len(results)} TESTS PASSED!")
        else:
            print(f"\n  ⚠️  {failed} test(s) failed — check logs.")

        # Save results
        results_file = Config.LOG_DIR / "robust_test_results.json"
        with open(results_file, "w") as f:
            json.dump(results, f, indent=2, default=str)
        print(f"\n  📝 Results saved to: {results_file}")

        print("\n  Browser still open for inspection.")
        input("  Press ENTER to close > ")

    except KeyboardInterrupt:
        print("\n\n  Cancelled by user.")
    except Exception as e:
        log.error(f"Test failed: {e}", exc_info=True)
        print(f"\n  ❌ Fatal error: {e}")
    finally:
        await browser.close()
        print("  Browser closed.\n")


if __name__ == "__main__":
    asyncio.run(main())
</file>

<file path="src/__init__.py">

</file>

<file path="src/api/__init__.py">

</file>

<file path="src/api/openai_routes.py">
"""
OpenAI-compatible API routes.

Provides:
  POST /v1/chat/completions   — chat completions (with tool/function calling)
  GET  /v1/models             — list available models

All requests are serialized through an asyncio.Lock because the underlying
Playwright browser page is single-threaded.
"""

from __future__ import annotations

import asyncio
import json
import re
import time
import uuid
from typing import Any

from fastapi import APIRouter, HTTPException

from src.api.openai_schemas import (
    ChatCompletionRequest,
    ChatCompletionResponse,
    ChatMessage,
    Choice,
    ChoiceMessage,
    FunctionCallInfo,
    ImageData,
    ImageGenerationRequest,
    ImagesResponse,
    ModelListResponse,
    ModelObject,
    ToolCall,
    ToolDefinition,
    UsageInfo,
)
from src.chatgpt.client import ChatGPTClient
from src.claude.client import ClaudeClient
from src.config import Config
from src.log import setup_logging

log = setup_logging("openai_routes")

openai_router = APIRouter()

# Global reference — set by server.py at startup
_client: ChatGPTClient | ClaudeClient | None = None

# Serialize all requests — single browser page, not thread-safe
_lock = asyncio.Lock()


def _get_model_id() -> str:
    """Return model ID based on active provider."""
    if Config.PROVIDER == "claude":
        return "claude-browser"
    return "catgpt-browser"


def set_openai_client(client: ChatGPTClient | ClaudeClient) -> None:
    """Called by server.py to inject the client."""
    global _client
    _client = client


def _get_client() -> ChatGPTClient | ClaudeClient:
    if _client is None:
        raise HTTPException(status_code=503, detail="Client not initialized")
    return _client


# ── Helpers ─────────────────────────────────────────────────────


def _estimate_tokens(text: str) -> int:
    """Rough token estimate (~4 chars per token)."""
    return max(1, len(text) // 4)


def _extract_content_text(content) -> str:
    """Extract text from message content (handles both string and list format)."""
    if content is None:
        return ""
    if isinstance(content, str):
        return content
    if isinstance(content, list):
        parts = []
        for item in content:
            if isinstance(item, dict) and item.get("type") == "text":
                parts.append(item.get("text", ""))
        return "\n".join(parts) if parts else ""
    return str(content)


def _extract_image_urls(content) -> list[str]:
    """Extract image URLs from message content (OpenAI vision format)."""
    if not isinstance(content, list):
        return []
    urls = []
    for item in content:
        if isinstance(item, dict) and item.get("type") == "image_url":
            image_url = item.get("image_url", {})
            if isinstance(image_url, dict):
                url = image_url.get("url", "")
            else:
                url = str(image_url)
            if url:
                urls.append(url)
    return urls


def _extract_file_attachments(content) -> list[dict]:
    """
    Extract file attachments from message content.

    Supported content part format:
      {"type": "file", "file": {"filename": "test.pdf", "data": "base64...", "mime_type": "application/pdf"}}

    Also supports a shorthand data-URL style:
      {"type": "file", "file": {"filename": "test.pdf", "url": "data:application/pdf;base64,..."}}

    Returns list of dicts: [{"filename": str, "data_b64": str, "mime_type": str}, ...]
    """
    if not isinstance(content, list):
        return []
    files = []
    for item in content:
        if not isinstance(item, dict) or item.get("type") != "file":
            continue
        file_info = item.get("file", {})
        if not isinstance(file_info, dict):
            continue
        filename = file_info.get("filename", "attachment")
        # Two ways to supply file data:
        # 1. data + mime_type  2. url (data-URL)
        data_b64 = file_info.get("data")
        mime_type = file_info.get("mime_type", "application/octet-stream")
        url = file_info.get("url", "")
        if not data_b64 and url.startswith("data:"):
            # Parse data URL
            try:
                header, data_b64 = url.split(",", 1)
                # header = "data:application/pdf;base64"
                if ":" in header and ";" in header:
                    mime_type = header.split(":")[1].split(";")[0]
            except ValueError:
                continue
        if data_b64:
            files.append({"filename": filename, "data_b64": data_b64, "mime_type": mime_type})
    return files


async def _download_file(url_or_data: str | dict, download_dir: str = "/tmp/catgpt_files") -> str | None:
    """
    Download / decode a file (image, PDF, etc.) from URL, base64 data URL,
    or a file attachment dict. Returns the local file path.
    """
    import base64
    import hashlib
    import os

    os.makedirs(download_dir, exist_ok=True)

    # ── Dict form (from _extract_file_attachments) ──
    if isinstance(url_or_data, dict):
        try:
            filename = url_or_data.get("filename", "file")
            data_b64 = url_or_data["data_b64"]
            # Sanitize filename
            safe_name = re.sub(r"[^\w.\-]", "_", filename)
            hash_suffix = hashlib.md5(data_b64[:60].encode()).hexdigest()[:8]
            filepath = os.path.join(download_dir, f"{hash_suffix}_{safe_name}")
            with open(filepath, "wb") as f:
                f.write(base64.b64decode(data_b64))
            log.info(f"Decoded file attachment: {filepath}")
            return filepath
        except Exception as e:
            log.error(f"Failed to decode file attachment: {e}")
            return None

    # ── String forms ──
    url = str(url_or_data)

    if url.startswith("data:"):
        # Base64 data URL: data:image/png;base64,iVBOR... or data:application/pdf;base64,...
        try:
            header, b64data = url.split(",", 1)
            # Detect extension from MIME type
            ext = "bin"
            mime = ""
            if ":" in header and ";" in header:
                mime = header.split(":")[1].split(";")[0]
            ext_map = {
                "image/png": "png", "image/jpeg": "jpg", "image/webp": "webp",
                "image/gif": "gif", "application/pdf": "pdf",
                "text/plain": "txt", "text/csv": "csv",
                "application/json": "json",
                "application/vnd.openxmlformats-officedocument.wordprocessingml.document": "docx",
                "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet": "xlsx",
            }
            ext = ext_map.get(mime, mime.split("/")[-1] if "/" in mime else "bin")
            filename = f"file_{hashlib.md5(b64data[:100].encode()).hexdigest()[:12]}.{ext}"
            filepath = os.path.join(download_dir, filename)
            with open(filepath, "wb") as f:
                f.write(base64.b64decode(b64data))
            log.info(f"Decoded base64 file: {filepath}")
            return filepath
        except Exception as e:
            log.error(f"Failed to decode base64 data URL: {e}")
            return None
    elif url.startswith(("http://", "https://")):
        # HTTP URL — download it
        try:
            import urllib.request
            ext = "bin"
            for e in ["jpg", "jpeg", "webp", "gif", "png", "pdf", "txt", "csv", "docx", "xlsx"]:
                if e in url.lower():
                    ext = e
                    break
            filename = f"file_{hashlib.md5(url.encode()).hexdigest()[:12]}.{ext}"
            filepath = os.path.join(download_dir, filename)
            urllib.request.urlretrieve(url, filepath)
            log.info(f"Downloaded file: {filepath}")
            return filepath
        except Exception as e:
            log.error(f"Failed to download file from {url}: {e}")
            return None
    elif os.path.isfile(url):
        # Local file path
        return url
    else:
        log.warning(f"Unknown file URL format: {url[:80]}")
        return None


def _build_prompt(messages: list[ChatMessage]) -> str:
    """
    Flatten an OpenAI-style message array into a single prompt string
    that we can paste into ChatGPT's input box.

    The browser already maintains conversation context within a thread,
    so for simple single-turn calls we just send the last user message.
    For multi-turn with system prompts or tool results, we build a
    formatted transcript.
    """
    # Simple case: only one user message (and optionally one system message)
    non_system = [m for m in messages if m.role != "system"]
    system_msgs = [m for m in messages if m.role == "system"]

    # If it's just one user message, send it directly
    if len(non_system) == 1 and non_system[0].role == "user":
        prefix = ""
        if system_msgs:
            sys_text = _extract_content_text(system_msgs[0].content)
            if Config.PROVIDER == "claude":
                # Claude rejects "[System instruction: ...]" as prompt injection.
                # Present it as context instead.
                prefix = f"{sys_text}\n\n"
            else:
                prefix = f"[System instruction: {sys_text}]\n\n"
        user_text = _extract_content_text(non_system[0].content)
        return prefix + (user_text or "")

    # Multi-turn: build a transcript
    parts: list[str] = []
    for msg in messages:
        role = msg.role.capitalize()
        if msg.role == "system":
            if Config.PROVIDER == "claude":
                # For Claude, present system messages as context without the label
                text = _extract_content_text(msg.content)
                if text:
                    parts.append(text)
            else:
                text = _extract_content_text(msg.content)
                if text:
                    parts.append(f"System: {text}")
        elif msg.role == "tool":
            # Tool result — include both the call context and the result
            tool_content = _extract_content_text(msg.content)
            if Config.PROVIDER == "claude":
                parts.append(
                    f"The tool was executed and returned this result:\n{tool_content}\n\n"
                    f"Now use the result above to answer the user's original question in plain text."
                )
            else:
                parts.append(
                    f"[Tool result for {msg.tool_call_id or 'unknown'}]: {tool_content}\n\n"
                    f"Use the tool result to answer the user. Do NOT call tools again."
                )
        elif msg.role == "assistant" and msg.tool_calls:
            # Assistant requested tool calls — show what was called
            calls_desc = []
            for tc in msg.tool_calls:
                calls_desc.append(
                    f'{tc.function.name}({tc.function.arguments})'
                )
            parts.append(f"Assistant called tools: {', '.join(calls_desc)}")
        elif msg.content:
            text = _extract_content_text(msg.content)
            if text:
                parts.append(f"{role}: {text}")

    return "\n\n".join(parts)


def _build_tool_system_prompt(
    tools: list[ToolDefinition],
    tool_choice: str | dict | None = None,
) -> str:
    """
    Build a system-level instruction that tells the model about available tools.

    *tool_choice* controls how insistent the instructions are:
      - "auto" / None  — model decides whether to call a tool or answer directly
      - "required"     — model MUST call at least one tool
      - "none"         — caller should not call this function at all
      - {"type":"function","function":{"name":"X"}} — model MUST call that tool
    """
    tool_descriptions = []
    for tool in tools:
        fn = tool.function
        desc = {
            "name": fn.name,
            "description": fn.description,
            "parameters": fn.parameters,
        }
        tool_descriptions.append(json.dumps(desc, indent=2))

    tools_json = "\n---\n".join(tool_descriptions)

    # ── Determine the decision instruction based on tool_choice ──
    forced_tool_name = None
    if isinstance(tool_choice, dict):
        # {"type": "function", "function": {"name": "X"}}
        forced_tool_name = (
            tool_choice.get("function", {}).get("name")
            if isinstance(tool_choice.get("function"), dict)
            else None
        )

    if forced_tool_name:
        decision = (
            f"You MUST call the function `{forced_tool_name}`. "
            f"Do NOT answer the question yourself — output only the JSON tool call."
        )
    elif tool_choice == "required":
        decision = (
            "You MUST call at least one of the available functions. "
            "Do NOT answer the question yourself — always output tool calls."
        )
    else:
        # "auto" or None — model decides
        decision = (
            "If the user's request can be fulfilled or assisted by one or more "
            "of the available functions, call the appropriate tool(s). "
            "If none of the tools are relevant, answer the user normally in plain text."
        )

    # ── Provider-specific prompt framing ──
    if Config.PROVIDER == "claude":
        return f"""You have access to external tools through a structured interface. {decision}

When calling tools, respond with ONLY a JSON code block — no text before or after it:

```json
{{"tool_calls": [{{"name": "<function_name>", "arguments": {{...}}}}]}}
```

Rules:
1. Output ONLY the JSON code block when calling tools. Do not add any commentary, explanation, or text outside the code block.
2. You may call multiple functions in one response by adding them to the array.
3. Use the exact parameter names and types shown in each function's schema.
4. When you receive tool results in a follow-up message, use them to give the user a natural, helpful answer. Do NOT output another JSON tool call for the same request.

Available functions:
{tools_json}

Example — single tool:
```json
{{"tool_calls": [{{"name": "get_current_time", "arguments": {{}}}}]}}
```

Example — multiple tools:
```json
{{"tool_calls": [{{"name": "weather_forecast", "arguments": {{"city": "Tokyo", "date": "today"}}}}, {{"name": "calculate_expression", "arguments": {{"expression": "2+2"}}}}]}}
```
"""
    else:
        return f"""You are in tool-calling mode. {decision}

When calling tools, output ONLY a JSON code block — no other text:

```json
{{"tool_calls": [{{"name": "<function_name>", "arguments": {{...}}}}]}}
```

Rules:
1. Output ONLY the JSON code block when calling tools. No explanation, no text before or after.
2. You may call multiple functions in one response by adding them to the array.
3. Use the exact parameter names and types from each function's schema.
4. When a follow-up message contains tool results, summarize them naturally for the user. Do NOT call tools again for the same request.
5. Do not refuse or say tools are unavailable — they are available through this interface.

Available functions:
{tools_json}

Example — single tool:
```json
{{"tool_calls": [{{"name": "get_current_time", "arguments": {{}}}}]}}
```

Example — multiple tools:
```json
{{"tool_calls": [{{"name": "weather_forecast", "arguments": {{"city": "Tokyo", "date": "today"}}}}, {{"name": "calculate_expression", "arguments": {{"expression": "2+2"}}}}]}}
```
"""


def _extract_json_object(text: str, anchor: str = "tool_calls") -> str | None:
    """
    Extract a JSON object containing *anchor* key from *text*.

    Uses two strategies:
      1. Look inside markdown code blocks (```json ... ```)
      2. Find the anchor key and walk outward using brace-depth tracking
         to handle arbitrarily nested JSON (arrays, nested objects, etc.)
    """
    # Strategy 1: code blocks — most reliable when the model obeys the prompt
    for m in re.finditer(r"```(?:json)?\s*\n?([\s\S]*?)\n?\s*```", text):
        candidate = m.group(1).strip()
        if anchor in candidate:
            try:
                parsed = json.loads(candidate)
                if anchor in parsed:
                    return candidate
            except json.JSONDecodeError:
                continue

    # Strategy 2: locate anchor, walk to balanced braces
    search_key = f'"{anchor}"'
    idx = text.find(search_key)
    if idx == -1:
        return None

    # Walk backward to the nearest '{'
    start = text.rfind("{", 0, idx)
    if start == -1:
        return None

    # Walk forward tracking brace depth, respecting JSON string literals
    depth = 0
    in_string = False
    i = start
    while i < len(text):
        c = text[i]
        if in_string:
            if c == "\\":
                i += 2          # skip escaped char
                continue
            if c == '"':
                in_string = False
        else:
            if c == '"':
                in_string = True
            elif c == "{":
                depth += 1
            elif c == "}":
                depth -= 1
                if depth == 0:
                    candidate = text[start : i + 1]
                    try:
                        json.loads(candidate)
                        return candidate
                    except json.JSONDecodeError:
                        return None
        i += 1

    return None


def _parse_tool_calls(
    response_text: str, tools: list[ToolDefinition]
) -> list[ToolCall] | None:
    """
    Try to parse tool calls from the model's response text.

    Uses robust brace-matching extraction (handles nested JSON, arrays, etc.)
    then validates tool names against the provided tool definitions.
    Returns None if no valid tool calls are found.
    """
    json_str = _extract_json_object(response_text, "tool_calls")
    if not json_str:
        return None

    try:
        parsed = json.loads(json_str)
    except json.JSONDecodeError:
        log.debug(f"Failed to parse tool call JSON: {json_str[:200]}")
        return None

    if "tool_calls" not in parsed or not isinstance(parsed["tool_calls"], list):
        return None

    # Validate that the called functions are in the provided tools
    valid_names = {t.function.name for t in tools}
    result: list[ToolCall] = []

    for call in parsed["tool_calls"]:
        name = call.get("name", "")
        if name not in valid_names:
            log.warning(f"Model called unknown tool: {name}")
            continue

        arguments = call.get("arguments", {})
        if isinstance(arguments, dict):
            arguments_str = json.dumps(arguments)
        else:
            arguments_str = str(arguments)

        result.append(
            ToolCall(
                id=f"call_{uuid.uuid4().hex[:24]}",
                type="function",
                function=FunctionCallInfo(name=name, arguments=arguments_str),
            )
        )

    return result if result else None


# ── Routes ──────────────────────────────────────────────────────


@openai_router.get("/v1/models", response_model=ModelListResponse)
async def list_models() -> ModelListResponse:
    """List available models — returns our single browser-backed model."""
    model_id = _get_model_id()
    owned_by = "anthropic" if Config.PROVIDER == "claude" else "catgpt"
    return ModelListResponse(
        data=[
            ModelObject(id=model_id, owned_by=owned_by),
        ]
    )


@openai_router.post("/v1/images/generations", response_model=ImagesResponse)
async def create_image(
    request: ImageGenerationRequest,
) -> ImagesResponse:
    """
    OpenAI-compatible image generation endpoint.

    Sends the prompt to ChatGPT which uses DALL-E to generate images.
    Downloads the generated images and returns them in OpenAI format.
    Supports response_format='b64_json' (default) or 'url' (local file path).
    """
    import base64

    if not request.prompt:
        raise HTTPException(status_code=400, detail="prompt cannot be empty")

    # Claude does not support image generation
    if Config.PROVIDER == "claude":
        raise HTTPException(
            status_code=501,
            detail="Image generation is not supported by Claude. This feature is only available with the ChatGPT provider.",
        )

    client = _get_client()

    async with _lock:
        start_time = time.time()

        # Build an image-generation prompt.
        # n > 1: we ask ChatGPT to generate multiple images
        # size/quality/style hints are included but ChatGPT web may ignore them.
        prompt_parts = [f"Generate an image: {request.prompt}"]
        if request.n and request.n > 1:
            prompt_parts.append(f"Please generate {request.n} different images.")
        if request.size and request.size != "1024x1024":
            prompt_parts.append(f"Image size: {request.size}.")
        if request.quality == "hd":
            prompt_parts.append("Make it high-definition / highly detailed.")
        if request.style == "natural":
            prompt_parts.append("Use a natural, realistic style.")

        full_prompt = " ".join(prompt_parts)

        log.info(
            f"POST /v1/images/generations — prompt='{request.prompt[:80]}', "
            f"n={request.n}, size={request.size}, response_format={request.response_format}"
        )

        # Send to ChatGPT
        try:
            result = await client.send_message(full_prompt)
        except Exception as e:
            log.error(f"Provider error during image generation: {e}", exc_info=True)
            raise HTTPException(status_code=500, detail=f"Provider error: {str(e)}")

        elapsed_ms = int((time.time() - start_time) * 1000)

        # Check if ChatGPT generated images
        if not result.images:
            # ChatGPT may have responded with text instead of generating an image.
            # This can happen when the model declines or gives a text description.
            log.warning(
                f"No images detected in response ({elapsed_ms}ms). "
                f"ChatGPT replied: {result.message[:200]}"
            )
            raise HTTPException(
                status_code=422,
                detail=(
                    f"ChatGPT did not generate an image. "
                    f"Model response: {result.message[:500]}"
                ),
            )

        # Build image data objects
        image_data_list: list[ImageData] = []
        for img_info in result.images:
            revised_prompt = img_info.prompt_title or img_info.alt or request.prompt

            if request.response_format == "b64_json":
                # Read the downloaded file and base64-encode it
                if img_info.local_path:
                    try:
                        with open(img_info.local_path, "rb") as f:
                            img_bytes = f.read()
                        b64 = base64.b64encode(img_bytes).decode("utf-8")
                        image_data_list.append(
                            ImageData(
                                b64_json=b64,
                                revised_prompt=revised_prompt,
                            )
                        )
                    except Exception as e:
                        log.error(f"Failed to read image file {img_info.local_path}: {e}")
                else:
                    log.warning(f"Image has no local_path: {img_info.url[:80]}")
            else:
                # response_format == "url" → return local file path as URL
                image_data_list.append(
                    ImageData(
                        url=img_info.local_path or img_info.url,
                        revised_prompt=revised_prompt,
                    )
                )

        if not image_data_list:
            raise HTTPException(
                status_code=500,
                detail="Images were detected but could not be processed.",
            )

        log.info(
            f"Image generation complete: {len(image_data_list)} image(s), "
            f"{elapsed_ms}ms, format={request.response_format}"
        )

        return ImagesResponse(data=image_data_list)


@openai_router.post("/v1/chat/completions", response_model=ChatCompletionResponse)
async def create_chat_completion(
    request: ChatCompletionRequest,
) -> ChatCompletionResponse:
    """
    OpenAI-compatible chat completions endpoint.

    Converts the message array into a single prompt, sends it to ChatGPT
    via browser automation, and returns an OpenAI-formatted response.
    Supports tool/function calling via prompt injection.
    """
    # ── Validate ────────────────────────────────────────────
    if request.stream:
        raise HTTPException(
            status_code=400,
            detail="Streaming is not supported. Set stream=false or omit it.",
        )

    if not request.messages:
        raise HTTPException(status_code=400, detail="messages array cannot be empty")

    client = _get_client()

    async with _lock:
        start_time = time.time()

        # ── Build the prompt ────────────────────────────────
        messages = list(request.messages)

        # If tools are provided, inject tool definitions as a system prompt
        # (unless tool_choice="none", which means ignore tools)
        has_tool_prompt = False
        if request.tools and request.tool_choice != "none":
            tool_system = _build_tool_system_prompt(
                request.tools, tool_choice=request.tool_choice
            )
            # Prepend as the first system message
            messages.insert(0, ChatMessage(role="system", content=tool_system))
            has_tool_prompt = True

        prompt = _build_prompt(messages)
        log.info(
            f"POST /v1/chat/completions — model={request.model}, "
            f"{len(request.messages)} messages, prompt={len(prompt)} chars"
        )

        # ── Extract attachments from messages ──────────────
        image_paths: list[str] = []
        file_paths: list[str] = []
        for msg in request.messages:
            if msg.role == "user" and isinstance(msg.content, list):
                # Images (OpenAI vision format)
                image_urls = _extract_image_urls(msg.content)
                for url in image_urls:
                    local_path = await _download_file(url)
                    if local_path:
                        image_paths.append(local_path)
                # Generic file attachments
                file_attachments = _extract_file_attachments(msg.content)
                for fa in file_attachments:
                    local_path = await _download_file(fa)
                    if local_path:
                        file_paths.append(local_path)

        all_attachment_paths = image_paths + file_paths
        if all_attachment_paths:
            log.info(f"Extracted {len(image_paths)} image(s) and {len(file_paths)} file(s) from request")

        # ── Send to ChatGPT ────────────────────────────────
        try:
            result = await client.send_message(
                prompt,
                image_paths=image_paths or None,
                file_paths=file_paths or None,
            )
        except Exception as e:
            log.error(f"Provider error: {e}", exc_info=True)
            raise HTTPException(status_code=500, detail=f"Provider error: {str(e)}")

        response_text = result.message
        elapsed_ms = int((time.time() - start_time) * 1000)

        # ── Detect echo (extraction grabbed sent prompt instead of reply) ──
        _echo_markers = ["[System instruction:", "tool-calling mode", "Available functions:"]
        if response_text and has_tool_prompt and any(m in response_text for m in _echo_markers):
            log.warning("Response appears to echo the sent prompt — retrying extraction")
            try:
                await asyncio.sleep(3)
                if Config.PROVIDER == "claude":
                    from src.claude.detector import extract_last_response_via_copy
                else:
                    from src.chatgpt.detector import extract_last_response_via_copy
                retry_text = await extract_last_response_via_copy(client.page)
                if retry_text and not any(m in retry_text for m in _echo_markers):
                    response_text = retry_text
                    log.info(f"Retry extraction succeeded: {len(response_text)} chars")
                else:
                    log.warning("Retry extraction still echoed — stripping system prefix")
                    # Last resort: try to find assistant content after the prompt
                    idx = response_text.rfind("\n\n")
                    if idx > 0:
                        tail = response_text[idx:].strip()
                        if tail and not tail.startswith("["):
                            response_text = tail
            except Exception as e:
                log.warning(f"Retry extraction failed: {e}")

        # ── Check for tool calls ────────────────────────────
        tool_calls = None
        finish_reason = "stop"

        if has_tool_prompt and request.tools:
            tool_calls = _parse_tool_calls(response_text, request.tools)
            if tool_calls:
                finish_reason = "tool_calls"
                # When the model calls tools, content should be null
                response_text = None

        # ── Build response ──────────────────────────────────
        prompt_tokens = _estimate_tokens(prompt)
        completion_tokens = _estimate_tokens(response_text or "")

        response = ChatCompletionResponse(
            model=request.model,
            choices=[
                Choice(
                    index=0,
                    message=ChoiceMessage(
                        role="assistant",
                        content=response_text,
                        tool_calls=tool_calls,
                    ),
                    finish_reason=finish_reason,
                )
            ],
            usage=UsageInfo(
                prompt_tokens=prompt_tokens,
                completion_tokens=completion_tokens,
                total_tokens=prompt_tokens + completion_tokens,
            ),
        )

        log.info(
            f"Response: {elapsed_ms}ms, finish_reason={finish_reason}, "
            f"tokens≈{response.usage.total_tokens}"
        )

        return response
</file>

<file path="src/api/openai_schemas.py">
"""
OpenAI-compatible Pydantic schemas for /v1/chat/completions and /v1/models.

Mirrors the OpenAI Chat Completions API specification so that any OpenAI SDK
or LangChain client can talk to our browser-backed ChatGPT endpoint.
"""

from __future__ import annotations

import time
import uuid
from typing import Any, List, Optional, Union

from pydantic import BaseModel, Field


# ── Tool / Function definitions ─────────────────────────────────


class FunctionDefinition(BaseModel):
    """Schema for a function the model may call."""
    name: str
    description: str = ""
    parameters: dict[str, Any] = Field(default_factory=dict)


class ToolDefinition(BaseModel):
    """A tool the model may use (only 'function' type supported)."""
    type: str = "function"
    function: FunctionDefinition


class FunctionCallInfo(BaseModel):
    """Info about a specific function call made by the model."""
    name: str
    arguments: str  # JSON string


class ToolCall(BaseModel):
    """A tool call returned by the model."""
    id: str = Field(default_factory=lambda: f"call_{uuid.uuid4().hex[:24]}")
    type: str = "function"
    function: FunctionCallInfo


# ── Messages ────────────────────────────────────────────────────


class ChatMessage(BaseModel):
    """A single message in the conversation.
    
    Content can be:
    - A simple string
    - A list of content parts (OpenAI vision format + file attachments):
      [
        {"type": "text", "text": "..."},
        {"type": "image_url", "image_url": {"url": "data:image/png;base64,..."}},
        {"type": "file", "file": {"filename": "doc.pdf", "data": "<base64>", "mime_type": "application/pdf"}}
      ]
    """
    role: str  # system | user | assistant | tool
    content: Optional[Union[str, List[Any]]] = None
    name: Optional[str] = None
    tool_call_id: Optional[str] = None
    tool_calls: Optional[list[ToolCall]] = None


# ── Request ─────────────────────────────────────────────────────


class ChatCompletionRequest(BaseModel):
    """OpenAI-compatible chat completion request body."""
    model: str = "catgpt-browser"
    messages: list[ChatMessage]
    tools: Optional[list[ToolDefinition]] = None
    tool_choice: Optional[Union[str, dict]] = None  # "auto" | "none" | {"type":"function","function":{"name":"..."}}
    temperature: Optional[float] = None
    max_tokens: Optional[int] = None
    top_p: Optional[float] = None
    frequency_penalty: Optional[float] = None
    presence_penalty: Optional[float] = None
    stop: Optional[Union[str, List[str]]] = None
    stream: Optional[bool] = False
    n: Optional[int] = 1
    user: Optional[str] = None


# ── Response ────────────────────────────────────────────────────


class UsageInfo(BaseModel):
    """Token usage (estimated — we don't have real token counts)."""
    prompt_tokens: int = 0
    completion_tokens: int = 0
    total_tokens: int = 0


class ChoiceMessage(BaseModel):
    """The assistant's message in a choice."""
    role: str = "assistant"
    content: Optional[str] = None
    tool_calls: Optional[list[ToolCall]] = None


class Choice(BaseModel):
    """A single completion choice."""
    index: int = 0
    message: ChoiceMessage
    finish_reason: str = "stop"  # "stop" | "tool_calls"


class ChatCompletionResponse(BaseModel):
    """OpenAI-compatible chat completion response."""
    id: str = Field(default_factory=lambda: f"chatcmpl-{uuid.uuid4().hex[:24]}")
    object: str = "chat.completion"
    created: int = Field(default_factory=lambda: int(time.time()))
    model: str = "catgpt-browser"
    choices: list[Choice]
    usage: UsageInfo = Field(default_factory=UsageInfo)


# ── Models endpoint ─────────────────────────────────────────────


class ModelObject(BaseModel):
    """A model object for /v1/models."""
    id: str
    object: str = "model"
    created: int = 1700000000
    owned_by: str = "catgpt"


class ModelListResponse(BaseModel):
    """Response for GET /v1/models."""
    object: str = "list"
    data: list[ModelObject]


# ── Image Generation ────────────────────────────────────────────


class ImageGenerationRequest(BaseModel):
    """OpenAI-compatible image generation request (POST /v1/images/generations)."""
    prompt: str
    model: Optional[str] = "dall-e-3"
    n: Optional[int] = Field(default=1, ge=1, le=4)
    size: Optional[str] = "1024x1024"
    quality: Optional[str] = "standard"
    style: Optional[str] = "vivid"
    response_format: Optional[str] = "b64_json"  # "url" or "b64_json"
    user: Optional[str] = None


class ImageData(BaseModel):
    """A single generated image in the response."""
    url: Optional[str] = None
    b64_json: Optional[str] = None
    revised_prompt: Optional[str] = None


class ImagesResponse(BaseModel):
    """OpenAI-compatible image generation response."""
    created: int = Field(default_factory=lambda: int(time.time()))
    data: List[ImageData]
</file>

<file path="src/api/routes.py">
"""
API routes — FastAPI router for ChatGPT interaction.

Endpoints:
  POST /chat              Send a message in the current/new thread
  POST /thread/{id}/chat  Send a message in a specific thread
  POST /thread/new        Start a new conversation
  GET  /threads           List recent threads
  GET  /status            Health check + login status
"""

from __future__ import annotations

import asyncio

from fastapi import APIRouter, HTTPException

from src.api.schemas import (
    ChatRequest,
    ChatResponse,
    ImageInfoResponse,
    StatusResponse,
    ThreadInfo,
    ThreadListResponse,
)
from src.browser.manager import BrowserManager
from src.chatgpt.client import ChatGPTClient
from src.claude.client import ClaudeClient
from src.log import setup_logging

log = setup_logging("api_routes")

router = APIRouter()

# Serialize browser access — single page, not thread-safe
_lock = asyncio.Lock()

# Global reference — set by the server on startup
_client: ChatGPTClient | ClaudeClient | None = None
_browser: BrowserManager | None = None


def set_client(client: ChatGPTClient | ClaudeClient, browser: BrowserManager) -> None:
    """Called by server.py to inject the client instance."""
    global _client, _browser
    _client = client
    _browser = browser


def _get_client():
    if _client is None:
        raise HTTPException(status_code=503, detail="Client not initialized")
    return _client


def _build_response(result) -> ChatResponse:
    """Convert internal ChatResponse to API ChatResponse with image data."""
    images = [
        ImageInfoResponse(
            url=img.url,
            alt=img.alt,
            local_path=img.local_path,
            prompt_title=img.prompt_title,
        )
        for img in (result.images or [])
    ]
    return ChatResponse(
        message=result.message,
        thread_id=result.thread_id,
        response_time_ms=result.response_time_ms,
        images=images,
        has_images=result.has_images,
    )


# ── Chat ────────────────────────────────────────────────────────


@router.post("/chat", response_model=ChatResponse)
async def chat(req: ChatRequest) -> ChatResponse:
    """Send a message in the current conversation."""
    client = _get_client()
    log.info(f"POST /chat — {len(req.message)} chars")

    async with _lock:
        try:
            result = await client.send_message(req.message)
            return _build_response(result)
        except Exception as e:
            log.error(f"Chat error: {e}", exc_info=True)
            raise HTTPException(status_code=500, detail=str(e))


@router.post("/thread/{thread_id}/chat", response_model=ChatResponse)
async def chat_in_thread(thread_id: str, req: ChatRequest) -> ChatResponse:
    """Send a message in a specific thread. Navigates to it first."""
    client = _get_client()
    log.info(f"POST /thread/{thread_id}/chat — {len(req.message)} chars")

    async with _lock:
        try:
            # Navigate to the thread if not already there
            current_tid = client._extract_thread_id()
            if current_tid != thread_id:
                await client.navigate_to_thread(thread_id)

            result = await client.send_message(req.message)
            return _build_response(result)
        except Exception as e:
            log.error(f"Thread chat error: {e}", exc_info=True)
            raise HTTPException(status_code=500, detail=str(e))


@router.post("/thread/new", response_model=ChatResponse)
async def new_thread(req: ChatRequest) -> ChatResponse:
    """Start a new conversation and send the first message."""
    client = _get_client()
    log.info(f"POST /thread/new — {len(req.message)} chars")

    async with _lock:
        try:
            await client.new_chat()
            result = await client.send_message(req.message)
            return _build_response(result)
        except Exception as e:
            log.error(f"New thread error: {e}", exc_info=True)
            raise HTTPException(status_code=500, detail=str(e))


# ── Threads ─────────────────────────────────────────────────────


@router.get("/threads", response_model=ThreadListResponse)
async def list_threads() -> ThreadListResponse:
    """List recent conversation threads from the sidebar."""
    client = _get_client()
    log.info("GET /threads")

    async with _lock:
        try:
            raw_threads = await client.list_threads()
            threads = [
                ThreadInfo(id=t["id"], title=t["title"], url=t["url"])
                for t in raw_threads
            ]
            return ThreadListResponse(threads=threads)
        except Exception as e:
            log.error(f"Threads list error: {e}", exc_info=True)
            raise HTTPException(status_code=500, detail=str(e))


# ── Status ──────────────────────────────────────────────────────


@router.get("/status", response_model=StatusResponse)
async def status() -> StatusResponse:
    """Health check — returns login status and current thread."""
    try:
        client = _get_client()
        logged_in = await _browser.is_logged_in()
        tid = client._extract_thread_id()
        return StatusResponse(status="ok", logged_in=logged_in, current_thread=tid)
    except Exception:
        return StatusResponse(status="ok", logged_in=False, current_thread="")
</file>

<file path="src/api/schemas.py">
"""
Pydantic request/response schemas for the API.
"""

from __future__ import annotations

from pydantic import BaseModel, Field


class ChatRequest(BaseModel):
    """Request body for sending a message."""
    message: str = Field(..., min_length=1, description="The message to send to ChatGPT")


class ImageInfoResponse(BaseModel):
    """Image metadata in API response."""
    url: str = Field("", description="Original image URL from ChatGPT/DALL-E")
    alt: str = Field("", description="Alt text / image description")
    local_path: str = Field("", description="Local file path after download")
    prompt_title: str = Field("", description="Image generation title shown by ChatGPT")


class ChatResponse(BaseModel):
    """Response body with ChatGPT's reply."""
    message: str = Field(..., description="ChatGPT's response text (markdown)")
    thread_id: str = Field("", description="Conversation thread ID")
    response_time_ms: int = Field(0, description="Time to generate the response in ms")
    images: list[ImageInfoResponse] = Field(default_factory=list, description="Generated images")
    has_images: bool = Field(False, description="Whether the response contains images")


class ThreadInfo(BaseModel):
    """Thread metadata."""
    id: str
    title: str
    url: str


class ThreadListResponse(BaseModel):
    """List of recent threads."""
    threads: list[ThreadInfo]


class StatusResponse(BaseModel):
    """Health check / status."""
    status: str = "ok"
    logged_in: bool = False
    current_thread: str = ""
</file>

<file path="src/api/server.py">
"""
FastAPI server — serves ChatGPT as an API.

Launches the browser on startup, shuts it down on exit.

Usage:
    python -m src.api.server
    # or
    uvicorn src.api.server:app --host 0.0.0.0 --port 8000
"""

from __future__ import annotations

import asyncio
from contextlib import asynccontextmanager

from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from starlette.middleware.base import BaseHTTPMiddleware
from starlette.requests import Request
from starlette.responses import JSONResponse

from src.browser.manager import BrowserManager
from src.browser.auto_login import ensure_logged_in
from src.chatgpt.client import ChatGPTClient
from src.claude.client import ClaudeClient
from src.config import Config
from src.api.routes import router, set_client
from src.api.openai_routes import openai_router, set_openai_client
from src.log import setup_logging

log = setup_logging("api_server")

# Global instances — needed for lifespan
_browser: BrowserManager | None = None
_client: ChatGPTClient | ClaudeClient | None = None


@asynccontextmanager
async def lifespan(app: FastAPI):
    """Startup: launch browser. Shutdown: close it."""
    global _browser, _client

    log.info("Starting browser for API server...")
    _browser = BrowserManager()
    page = await _browser.start()

    target_url = Config.provider_url()
    provider_name = "Claude" if Config.PROVIDER == "claude" else "ChatGPT"
    log.info(f"Provider: {provider_name} ({target_url})")

    # Navigate with retries (DNS can be slow in Docker)
    max_retries = 5
    for attempt in range(1, max_retries + 1):
        try:
            log.info(f"Navigation attempt {attempt}/{max_retries} to {target_url}")
            await _browser.navigate(target_url)
            break
        except Exception as e:
            log.warning(f"Navigation attempt {attempt} failed: {e}")
            if attempt == max_retries:
                log.error("All navigation attempts failed")
                raise
            wait_time = attempt * 5  # 5s, 10s, 15s, 20s
            log.info(f"Retrying in {wait_time}s...")
            await asyncio.sleep(wait_time)

    # Apply stealth patches AFTER the first navigation.
    # In Docker, applying stealth init scripts before navigation
    # causes Chrome's DNS resolver to fail (ERR_NAME_NOT_RESOLVED).
    await _browser.apply_stealth_patches()

    await asyncio.sleep(3)

    if not await _browser.is_logged_in():
        log.info("Not logged in — starting auto-login flow...")
        logged_in = await ensure_logged_in(_browser)
        if not logged_in:
            log.error(f"Login failed after auto-login attempt")
            raise RuntimeError(f"Could not log in to {provider_name}")

    if Config.PROVIDER == "claude":
        _client = ClaudeClient(page)
    else:
        _client = ChatGPTClient(page)

    set_client(_client, _browser)
    set_openai_client(_client)
    log.info(f"API server ready — browser launched, logged in to {provider_name}")

    yield  # Server is running

    log.info("Shutting down — closing browser...")
    await _browser.close()
    log.info("Browser closed")


app = FastAPI(
    title="CatGPT Gateway API",
    description=(
        "Browser automation API for ChatGPT and Claude. "
        "Sends messages via browser and returns responses."
    ),
    version="1.0.0",
    lifespan=lifespan,
)

# ── Bearer Token Auth Middleware ────────────────────────────────
class BearerTokenMiddleware(BaseHTTPMiddleware):
    """
    Require a Bearer token on all API requests when API_TOKEN is set.
    Skips auth for /docs, /openapi.json, and health-check paths.
    """

    OPEN_PATHS = {"/docs", "/redoc", "/openapi.json", "/healthz"}

    async def dispatch(self, request: Request, call_next):
        token = Config.API_TOKEN
        if not token:
            # No token configured — auth disabled
            return await call_next(request)

        path = request.url.path
        if path in self.OPEN_PATHS:
            return await call_next(request)

        # Check Authorization header
        auth_header = request.headers.get("authorization", "")
        if auth_header.startswith("Bearer "):
            provided = auth_header[7:]
        else:
            provided = ""

        if provided != token:
            return JSONResponse(
                status_code=401,
                content={"error": {"message": "Invalid or missing API token. Set Authorization: Bearer <API_TOKEN>", "type": "auth_error"}},
            )

        return await call_next(request)


app.add_middleware(BearerTokenMiddleware)

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

app.include_router(router)
app.include_router(openai_router)


@app.get("/healthz", include_in_schema=False)
async def healthz():
    """Unauthenticated health-check for Docker / load-balancers."""
    return {"status": "ok"}


if __name__ == "__main__":
    import uvicorn

    uvicorn.run(
        "src.api.server:app",
        host=Config.API_HOST,
        port=Config.API_PORT,
        log_level="info",
    )
</file>

<file path="src/browser/__init__.py">
from src.browser.manager import BrowserManager as Browser
</file>

<file path="src/browser/auto_login.py">
"""
Auto-login helper — detects missing login and prompts user to sign in.

Used by both the FastAPI server and the TUI to automatically trigger
first-time login when no existing session is found, instead of crashing.
"""

from __future__ import annotations

import asyncio

from src.browser.manager import BrowserManager
from src.config import Config
from src.log import setup_logging

log = setup_logging("auto_login")


async def ensure_logged_in(browser: BrowserManager) -> bool:
    """
    Check if the user is logged in. If not, guide them through login.

    This replaces the need to manually run `scripts/first_login.py`.
    Opens the browser to ChatGPT, waits for the user to sign in,
    and verifies the login before returning.

    Returns True if logged in (or successfully logged in now).
    Raises RuntimeError if login fails after the user presses Enter.
    """
    if await browser.is_logged_in():
        log.info("Already logged in")
        return True

    log.info("Not logged in — starting interactive login flow")

    provider_name = "Claude" if Config.PROVIDER == "claude" else "ChatGPT"
    target_url = Config.provider_url()

    print("\n" + "=" * 60)
    print(f"  🔐 {provider_name} Login Required — First-Time Setup")
    print("=" * 60)
    print(f"\n  Browser data dir: {Config.BROWSER_DATA_DIR}")
    print(f"  Target: {target_url}")
    print(f"\n  A Chrome window is open. Please:")
    print(f"  1. Sign in to {provider_name} with your account")
    print("  2. Complete any CAPTCHA / verification checks")
    print("  3. Wait until you see the chat interface")
    print("  4. Come back here and press Enter")
    print("\n" + "=" * 60 + "\n")

    # Wait for user to sign in
    await asyncio.get_event_loop().run_in_executor(
        None, lambda: input("  Press ENTER after you've signed in successfully > ")
    )

    # Give the page a moment to settle
    await asyncio.sleep(2)

    # Verify login
    if await browser.is_logged_in():
        print("\n  ✅ Login verified! Session saved.")
        print("  You won't need to sign in again.\n")
        log.info("Interactive login completed successfully")
        return True
    else:
        print("\n  ⚠️  Could not verify login.")
        print("  The session may still be saved — trying to continue...\n")
        log.warning("Login verification uncertain after interactive login")
        # Don't crash — let the caller decide what to do
        return False
</file>

<file path="src/browser/human.py">
"""
Human-like behavior simulation — typing, clicking, delays, mouse movement.

Makes browser automation look organic to anti-bot systems.
"""

from __future__ import annotations

import asyncio
import random

from patchright.async_api import Page

from src.config import Config
from src.log import setup_logging

log = setup_logging("human")


async def random_delay(min_ms: int | None = None, max_ms: int | None = None) -> None:
    """Sleep for a random duration (milliseconds)."""
    lo = min_ms or Config.THINKING_PAUSE_MIN
    hi = max_ms or Config.THINKING_PAUSE_MAX
    ms = random.randint(lo, hi)
    log.debug(f"Random delay: {ms}ms")
    await asyncio.sleep(ms / 1000)


async def human_type(page: Page, selector: str, text: str) -> None:
    """
    Paste text all at once into the input field.

    Uses keyboard.insert_text() which behaves like a clipboard paste —
    fast and reliable, avoids issues with per-character typing on
    contenteditable divs.
    """
    element = page.locator(selector).first
    await element.click()
    # Small pause after focusing (human would take a moment)
    await asyncio.sleep(random.uniform(0.1, 0.25))

    log.debug(f"Pasting {len(text)} chars into {selector}")
    await page.keyboard.insert_text(text)
    log.debug("Paste complete")


async def human_click(page: Page, selector: str) -> None:
    """
    Click an element with human-like behavior:
    1. Hover over element (triggers mouseover)
    2. Brief pause
    3. Click

    Uses .first to handle cases where multiple elements match.
    """
    element = page.locator(selector).first
    await element.hover()
    await asyncio.sleep(random.uniform(0.05, 0.15))
    await element.click()
    log.debug(f"Human-clicked: {selector}")


async def idle_mouse_movement(page: Page) -> None:
    """
    Simulate idle mouse movement — small random movements to look alive.
    Call this periodically while waiting for responses.
    """
    try:
        viewport = page.viewport_size
        if viewport:
            x = random.randint(100, viewport["width"] - 100)
            y = random.randint(100, viewport["height"] - 100)
            await page.mouse.move(x, y, steps=random.randint(5, 15))
            log.debug(f"Idle mouse move to ({x}, {y})")
    except Exception:
        pass  # Non-critical — don't break the flow


async def thinking_pause() -> None:
    """Simulate a 'thinking' pause before the user starts typing."""
    ms = random.randint(Config.THINKING_PAUSE_MIN, Config.THINKING_PAUSE_MAX)
    log.debug(f"Thinking pause: {ms}ms")
    await asyncio.sleep(ms / 1000)
</file>

<file path="src/browser/manager.py">
"""
Browser lifecycle manager — launch, persist, close.

Uses a persistent Chrome context so the user only signs in once.
Session data (cookies, localStorage, IndexedDB) survives restarts.
"""

from __future__ import annotations

import os
import random
import signal
import socket
from pathlib import Path
from patchright.async_api import async_playwright, BrowserContext, Page, Playwright

from src.config import Config
from src.browser.stealth import apply_stealth
from src.log import setup_logging

log = setup_logging("browser")


def _resolve_domains_for_chrome() -> str:
    """
    Pre-resolve key domains and return a --host-resolver-rules string
    for Chrome. This works around Chrome's built-in DNS resolver
    failing inside Docker containers.

    Returns empty string if not running in Docker or all resolutions fail.
    """
    # Only needed in Docker (check for /.dockerenv or DISPLAY=:99)
    if not os.path.exists("/.dockerenv") and os.environ.get("DISPLAY") != ":99":
        return ""

    domains = [
        "chatgpt.com",
        "cdn.oaistatic.com",
        "ab.chatgpt.com",
        "auth.openai.com",
        "auth0.openai.com",
        "openai.com",
        "api.openai.com",
        "platform.openai.com",
        "challenges.cloudflare.com",
        "static.cloudflareinsights.com",
        "tcr9i.chat.openai.com",
        # Claude domains
        "claude.ai",
        "api.claude.ai",
        "cdn.claude.ai",
        "anthropic.com",
        "www.anthropic.com",
    ]
    rules = []
    for domain in domains:
        try:
            ip = socket.gethostbyname(domain)
            rules.append(f"MAP {domain} {ip}")
            log.debug(f"DNS pre-resolve: {domain} -> {ip}")
        except Exception as e:
            log.warning(f"DNS pre-resolve failed: {domain} -> {e}")

    if rules:
        result = ", ".join(rules)
        log.info(f"Chrome host-resolver-rules: {len(rules)} domains mapped")
        return result
    return ""


def _cleanup_stale_locks(data_dir: Path) -> None:
    """
    Remove stale lock / journal / WAL files that prevent browser launch.

    After a crash, Chromium leaves behind:
    - SingletonLock/Socket/Cookie — prevents new instance from using data dir.
    - *-journal, *-wal, *-shm — SQLite journal/WAL files that cause
      "database is locked" errors (UKM, Top Sites, History, etc.)

    We also attempt to kill any orphan chrome-for-testing processes.
    """
    import subprocess

    # 1. Kill orphan chrome-for-testing processes FIRST
    try:
        result = subprocess.run(
            ["pkill", "-f", "chrome-for-testing"],
            capture_output=True, timeout=3
        )
        if result.returncode == 0:
            log.info("Killed orphan chrome processes")
            import time
            time.sleep(1)
    except Exception:
        pass  # Non-critical

    # 2. Remove singleton lock files
    lock_files = ["SingletonLock", "SingletonSocket", "SingletonCookie"]
    for name in lock_files:
        path = data_dir / name
        if path.exists():
            try:
                path.unlink()
                log.info(f"Removed stale lock file: {name}")
            except Exception as e:
                log.warning(f"Could not remove {name}: {e}")

    # 3. Remove SQLite journal/WAL/SHM files that cause "database is locked"
    import glob as _glob
    patterns = ["**/*-journal", "**/*-wal", "**/*-shm"]
    removed = 0
    for pattern in patterns:
        for path_str in _glob.glob(str(data_dir / pattern), recursive=True):
            try:
                Path(path_str).unlink()
                removed += 1
            except Exception:
                pass
    if removed:
        log.info(f"Removed {removed} stale SQLite journal/WAL/SHM files")


class BrowserManager:
    """Manages a single persistent Chromium browser context."""

    def __init__(self) -> None:
        self._playwright: Playwright | None = None
        self._context: BrowserContext | None = None
        self._page: Page | None = None

    async def start(self) -> Page:
        """
        Launch a persistent Chrome context with stealth and human-like settings.

        Automatically cleans up stale lock files from previous crashed sessions.
        Returns the active page ready for navigation.
        """
        Config.ensure_dirs()

        # Clean up stale locks from previous sessions
        _cleanup_stale_locks(Config.BROWSER_DATA_DIR)

        log.info("Launching browser...")
        self._playwright = await async_playwright().start()

        # Randomize viewport slightly to avoid fingerprint consistency
        width = Config.VIEWPORT_WIDTH + random.randint(-20, 20)
        height = Config.VIEWPORT_HEIGHT + random.randint(-20, 20)

        # Try real Chrome first, fall back to bundled Chromium
        chrome_args = [
            "--disable-blink-features=AutomationControlled",
            "--no-first-run",
            "--no-default-browser-check",
        ]

        # Docker-specific flags
        if os.path.exists("/.dockerenv") or os.environ.get("DISPLAY") == ":99":
            chrome_args.extend([
                "--no-sandbox",
                "--disable-setuid-sandbox",
                "--disable-gpu",
            ])

        # In Docker, Chrome's DNS resolver can fail. Pre-resolve domains
        # and pass them directly via --host-resolver-rules.
        resolver_rules = _resolve_domains_for_chrome()
        if resolver_rules:
            chrome_args.append(f"--host-resolver-rules={resolver_rules}")

        launch_kwargs = dict(
            user_data_dir=str(Config.BROWSER_DATA_DIR),
            headless=Config.HEADLESS,
            slow_mo=Config.SLOW_MO,
            viewport={"width": width, "height": height},
            locale="en-US",
            timezone_id="America/Los_Angeles",
            args=chrome_args,
        )

        try:
            self._context = await self._playwright.chromium.launch_persistent_context(
                channel="chrome", **launch_kwargs
            )
            log.info("Launched with real Chrome")
        except Exception:
            log.info("Real Chrome not found, using bundled Chromium")
            self._context = await self._playwright.chromium.launch_persistent_context(
                **launch_kwargs
            )

        # NOTE: Stealth patches are applied AFTER the first navigation.
        # In Docker, applying stealth init scripts before navigation
        # causes Chrome's DNS resolver to fail (ERR_NAME_NOT_RESOLVED).
        # Call apply_stealth_patches() after navigating to the target page.

        # Use existing page or create one
        if self._context.pages:
            self._page = self._context.pages[0]
        else:
            self._page = await self._context.new_page()

        log.info(f"Browser ready — viewport {width}x{height}")
        return self._page

    async def apply_stealth_patches(self) -> None:
        """
        Apply stealth patches to the browser context.

        Must be called AFTER the first page navigation, not before.
        In Docker containers, applying stealth init scripts before any
        navigation causes Chrome's DNS resolver to fail.
        """
        if self._context is None:
            raise RuntimeError("Browser not started. Call start() first.")
        await apply_stealth(self._context)

    @property
    def page(self) -> Page:
        """Get the active page. Raises if browser not started."""
        if self._page is None:
            raise RuntimeError("Browser not started. Call start() first.")
        return self._page

    @property
    def context(self) -> BrowserContext:
        """Get the browser context."""
        if self._context is None:
            raise RuntimeError("Browser not started. Call start() first.")
        return self._context

    async def navigate(self, url: str) -> None:
        """Navigate to a URL and wait for page load."""
        log.info(f"Navigating to {url}")
        await self.page.goto(url, wait_until="domcontentloaded")
        log.info("Page loaded")

    async def is_logged_in(self) -> bool:
        """
        Check if user is logged in by looking for chat input vs login indicators.

        Returns True if the chat interface is visible, False if login page detected.
        """
        from src.selectors import Selectors
        from src.claude.selectors import ClaudeSelectors

        if Config.PROVIDER == "claude":
            chat_inputs = ClaudeSelectors.CHAT_INPUT
            login_indicators = ClaudeSelectors.LOGIN_INDICATORS
            logged_in_indicators = ClaudeSelectors.LOGGED_IN_INDICATORS
        else:
            chat_inputs = Selectors.CHAT_INPUT
            login_indicators = Selectors.LOGIN_INDICATORS
            logged_in_indicators = []

        try:
            # Try to find the chat input
            for selector in chat_inputs:
                try:
                    el = await self.page.wait_for_selector(selector, timeout=3000)
                    if el:
                        log.info("Login check: LOGGED IN (chat input found)")
                        return True
                except Exception:
                    continue

            # Claude: also check for user-menu-button as a logged-in signal
            for selector in logged_in_indicators:
                try:
                    el = await self.page.wait_for_selector(selector, timeout=2000)
                    if el:
                        log.info("Login check: LOGGED IN (user menu found)")
                        return True
                except Exception:
                    continue

            # Check for login indicators
            for selector in login_indicators:
                try:
                    el = await self.page.wait_for_selector(selector, timeout=2000)
                    if el:
                        log.warning("Login check: NOT LOGGED IN (login button found)")
                        return False
                except Exception:
                    continue

            log.warning("Login check: UNCERTAIN — no chat input or login button found")
            return False

        except Exception as e:
            log.error(f"Login check error: {e}")
            return False

    async def close(self) -> None:
        """Gracefully close the browser context and playwright instance."""
        log.info("Closing browser...")
        try:
            if self._context:
                await self._context.close()
            if self._playwright:
                await self._playwright.stop()
        except Exception as e:
            log.error(f"Error closing browser: {e}")
        finally:
            self._context = None
            self._page = None
            self._playwright = None
            log.info("Browser closed")
</file>

<file path="src/browser/stealth.py">
"""
Stealth wrapper — configures playwright-stealth to evade bot detection.

Uses page.evaluate() instead of context.add_init_script() because
add_init_script() breaks Chrome's DNS resolver inside Docker containers.
Any call to add_init_script() — even a trivial console.log — causes
net::ERR_NAME_NOT_RESOLVED on all subsequent navigations.

The workaround: inject the stealth JS via evaluate() on the current page,
and re-inject automatically on every frame navigation via an event listener.
"""

from __future__ import annotations

import os

from patchright.async_api import BrowserContext, Page, Frame
from playwright_stealth import Stealth

from src.log import setup_logging

log = setup_logging("stealth")

# Single Stealth instance; grab the JS payload once
_stealth = Stealth()
_STEALTH_JS: str = _stealth.script_payload

# Track whether we're in Docker (add_init_script is unsafe there)
_IN_DOCKER: bool = os.path.exists("/.dockerenv") or os.environ.get("DISPLAY") == ":99"


async def _inject_stealth_js(page: Page) -> None:
    """Inject stealth JS into the current page via evaluate()."""
    try:
        await page.evaluate(_STEALTH_JS)
    except Exception:
        # Page may have navigated away or closed — non-fatal
        pass


async def apply_stealth(context: BrowserContext) -> None:
    """
    Apply stealth patches to a browser context.

    In Docker: uses page.evaluate() + navigation listener (safe for DNS).
    Outside Docker: uses the standard add_init_script() approach.
    """
    if _IN_DOCKER:
        await _apply_stealth_docker(context)
    else:
        await _stealth.apply_stealth_async(context)
    log.info("Stealth patches applied to browser context")


async def _apply_stealth_docker(context: BrowserContext) -> None:
    """
    Docker-safe stealth: evaluate JS on every page and listen for navigations.

    add_init_script() is broken in Docker (kills DNS), so we:
    1. Inject stealth JS into all existing pages via evaluate()
    2. Hook 'framenavigated' to re-inject after every navigation
    3. Hook 'page' to cover new tabs/popups
    """

    async def on_frame_navigated(frame: Frame) -> None:
        """Re-inject stealth JS when the main frame navigates."""
        if frame == frame.page.main_frame:
            await _inject_stealth_js(frame.page)

    async def on_new_page(page: Page) -> None:
        """Inject stealth into new pages and attach navigation listener."""
        page.on("framenavigated", on_frame_navigated)
        await _inject_stealth_js(page)

    # Inject into all existing pages
    for page in context.pages:
        await _inject_stealth_js(page)
        page.on("framenavigated", on_frame_navigated)

    # Hook new pages (popups, new tabs)
    context.on("page", on_new_page)

    log.debug("Docker-safe stealth: evaluate + navigation listener active")


def get_stealth() -> Stealth:
    """Return the shared Stealth instance (for use with launch helpers)."""
    return _stealth
</file>

<file path="src/chatgpt/__init__.py">
from src.chatgpt.client import ChatGPTClient
</file>

<file path="src/chatgpt/client.py">
"""
ChatGPT client — core interaction logic.

Sends messages, waits for responses, manages conversations.
Handles selector fallbacks and integrates human-like behavior.
"""

from __future__ import annotations

import asyncio
import re
import time

from patchright.async_api import Page

from src.config import Config
from src.selectors import Selectors
from src.browser.human import human_type, human_click, thinking_pause, random_delay
from src.chatgpt.detector import (
    wait_for_response_complete,
    extract_last_response_via_copy,
    count_assistant_messages,
    get_latest_assistant_turn_signature,
    is_incomplete_response_text,
)
from src.chatgpt.image_handler import extract_images_from_response
from src.chatgpt.models import ChatResponse
from src.log import setup_logging

log = setup_logging("chatgpt_client")


class ChatGPTClient:
    """
    High-level client for interacting with the ChatGPT web interface.

    Requires a Playwright Page that is already logged in and on chatgpt.com.
    """

    def __init__(self, page: Page) -> None:
        self._page = page

    @property
    def page(self) -> Page:
        return self._page

    # ── Core: Send & Receive ────────────────────────────────────

    async def send_message(self, text: str, image_paths: list[str] | None = None, file_paths: list[str] | None = None) -> ChatResponse:
        """
        Send a message to ChatGPT and wait for the complete response.

        Args:
            text: The message text to send.
            image_paths: Optional list of local file paths to images to attach.
            file_paths: Optional list of local file paths to non-image files (PDF, etc.).

        Steps:
        1. Simulate thinking pause
        2. Upload images if provided
        3. Find and focus chat input
        4. Type message with human-like delays
        5. Click send
        6. Wait for response to complete
        7. Extract and return the response

        Returns ChatResponse with the assistant's reply and metadata.
        """
        all_attachments = (image_paths or []) + (file_paths or [])
        log.info(f"Sending message ({len(text)} chars, {len(all_attachments)} attachments): {text[:80]}...")
        start_time = time.time()

        # 0. Count existing assistant messages so we know when a new one appears
        pre_count = await count_assistant_messages(self._page)
        pre_turn_signature = await get_latest_assistant_turn_signature(self._page)
        log.debug(f"Assistant messages before send: {pre_count}")
        log.debug(f"Latest assistant turn before send: {pre_turn_signature}")

        # 1. Brief pause (human would take a moment to start typing)
        await random_delay(250, 700)

        # 1.5. Upload files/images if provided
        if all_attachments:
            await self._upload_files(all_attachments)

        # 2. Find the chat input
        input_selector = await self._find_selector(Selectors.CHAT_INPUT, "chat input")
        if not input_selector:
            raise RuntimeError("Could not find chat input element")

        # 3. Paste the message (all at once)
        await human_type(self._page, input_selector, text)

        # Small pause after pasting (like a human reviewing before send)
        await random_delay(150, 350)

        # 4. Send the message
        sent = await self._click_send()
        if not sent:
            # Fallback: try pressing Enter
            log.info("Send button not found, trying Enter key")
            await self._page.keyboard.press("Enter")

        # 5. Wait for response with message count awareness
        log.info("Waiting for ChatGPT response...")
        expected_count = pre_count + 1
        completed = await wait_for_response_complete(
            self._page,
            expected_msg_count=expected_count,
            previous_turn_signature=pre_turn_signature,
        )

        if not completed:
            log.warning("Response may not be complete (timeout)")

        # Small buffer after completion to let DOM settle
        await asyncio.sleep(0.5)

        # 6. Check for generated images in the response FIRST
        #    (image turns have no copy button, so we must detect images
        #    before trying copy-button extraction)
        images = await extract_images_from_response(self._page)
        has_images = len(images) > 0

        # 7. Extract text content
        if has_images:
            # Image responses don't have a copy button — extract text
            # from the turn's DOM instead (will get the image title/desc)
            response_text = await self._extract_image_turn_text(pre_turn_signature)
            log.info(f"Response contains {len(images)} generated image(s)")
            for img in images:
                log.info(f"  Image: {img.alt or img.prompt_title} → {img.local_path}")
        else:
            # Standard text response — use copy button (most reliable)
            response_text = await extract_last_response_via_copy(
                self._page,
                previous_turn_signature=pre_turn_signature,
            )

            # If we only captured a transient status (e.g. "Pro thinking"),
            # keep waiting and retry extraction on the same new turn.
            if is_incomplete_response_text(response_text):
                log.warning("Extracted text looks incomplete/transient; retrying for final answer")
                for attempt in range(1, 3):
                    await asyncio.sleep(4)
                    await wait_for_response_complete(
                        self._page,
                        timeout_ms=90000,
                        previous_turn_signature=pre_turn_signature,
                    )
                    retry_text = await extract_last_response_via_copy(
                        self._page,
                        previous_turn_signature=pre_turn_signature,
                    )

                    if retry_text and not is_incomplete_response_text(retry_text):
                        response_text = retry_text
                        log.info(f"Recovered final response text on retry {attempt}")
                        break

                    if retry_text:
                        response_text = retry_text
                    log.warning(f"Retry {attempt} still incomplete/transient")

        elapsed_ms = int((time.time() - start_time) * 1000)
        thread_id = self._extract_thread_id()

        log.info(
            f"Response received ({elapsed_ms}ms, {len(response_text)} chars"
            f"{f', {len(images)} images' if has_images else ''}): "
            f"{response_text[:80]}..."
        )

        return ChatResponse(
            message=response_text,
            thread_id=thread_id,
            response_time_ms=elapsed_ms,
            images=images,
            has_images=has_images,
        )

    # ── Navigation ──────────────────────────────────────────────

    async def new_chat(self) -> None:
        """Start a new conversation by navigating to the home page."""
        log.info("Starting new chat...")
        # Direct navigation is the most reliable way — avoids duplicate button issues
        await self._page.goto(Config.CHATGPT_URL, wait_until="domcontentloaded")
        await asyncio.sleep(1.5)

        # Wait for the chat input to be visible (signals page is ready)
        for selector in Selectors.CHAT_INPUT:
            try:
                await self._page.wait_for_selector(selector, timeout=10000, state="visible")
                log.debug(f"Chat input ready: {selector}")
                break
            except Exception:
                continue

        await random_delay(300, 600)
        log.info("New chat started (navigated to home)")

    async def navigate_to_thread(self, thread_id: str) -> None:
        """Navigate to an existing conversation thread."""
        url = f"{Config.CHATGPT_URL}/c/{thread_id}"
        log.info(f"Navigating to thread: {thread_id}")
        await self._page.goto(url, wait_until="domcontentloaded")
        await random_delay(800, 1500)
        log.info(f"Thread {thread_id} loaded")

    async def get_current_thread_url(self) -> str:
        """Get the current page URL (contains thread ID if in a conversation)."""
        return self._page.url

    # ── Sidebar ─────────────────────────────────────────────────

    async def list_threads(self) -> list[dict]:
        """
        Scrape the sidebar for recent conversation threads.

        Returns a list of dicts: [{id, title, url}, ...]
        """
        threads = []
        for selector in Selectors.SIDEBAR_THREAD_LINKS:
            try:
                elements = await self._page.query_selector_all(selector)
                for el in elements:
                    href = await el.get_attribute("href") or ""
                    title = (await el.inner_text()).strip()
                    match = re.search(r"/c/([a-f0-9-]+)", href)
                    if match:
                        threads.append({
                            "id": match.group(1),
                            "title": title,
                            "url": f"{Config.CHATGPT_URL}{href}",
                        })
                if threads:
                    break
            except Exception as e:
                log.debug(f"Sidebar scrape with {selector} failed: {e}")

        log.info(f"Found {len(threads)} threads in sidebar")
        return threads

    # ── Private Helpers ─────────────────────────────────────────

    async def _extract_image_turn_text(self, previous_turn_signature: str | None = None) -> str:
        """
        Extract any text content from the latest turn (for image responses).

        Image turns may contain a title/description like:
        "Creating image • Adorable orange tabby kitten close-up"
        """
        text = await self._page.evaluate("""
            (previousSignature) => {
                const turns = document.querySelectorAll('section[data-testid^="conversation-turn-"]');
                if (turns.length === 0) return '';

                let last = null;
                for (let idx = turns.length - 1; idx >= 0; idx--) {
                    const turn = turns[idx];
                    const turnRole = turn.getAttribute('data-turn');
                    const hasAssistantRole = turnRole === 'assistant' ||
                        Boolean(turn.querySelector('[data-message-author-role="assistant"]'));
                    if (!hasAssistantRole) continue;

                    const stableId =
                        turn.getAttribute('data-turn-id') ||
                        turn.getAttribute('data-testid') ||
                        turn.id ||
                        '';
                    const signature = `${idx}:${stableId}`;
                    if (previousSignature && signature === previousSignature) {
                        return '';
                    }

                    last = turn;
                    break;
                }

                if (!last) return '';

                // Try to get descriptive text (not "ChatGPT said:" heading)
                const spans = last.querySelectorAll('span');
                const parts = [];
                for (const span of spans) {
                    const t = (span.innerText || '').trim();
                    if (t && t.length > 3 && t.length < 300 &&
                        !t.includes('ChatGPT') && !t.includes('said')) {
                        parts.push(t);
                    }
                }
                if (parts.length > 0) return parts.join(' ');

                // Fallback: full turn inner text
                const full = (last.innerText || '').trim();
                // Strip the "ChatGPT said:" prefix
                return full.replace(/^ChatGPT said:\\s*/i, '').trim();
            }
        """, previous_turn_signature)
        return text or ""

    async def _find_selector(self, selectors: list[str], name: str) -> str | None:
        """
        Try each selector in the fallback list. Return the first one that matches.
        """
        for selector in selectors:
            try:
                el = await self._page.wait_for_selector(
                    selector,
                    timeout=Config.SELECTOR_TIMEOUT,
                    state="visible",
                )
                if el:
                    log.debug(f"Found {name} via: {selector}")
                    return selector
            except Exception:
                log.debug(f"Selector miss for {name}: {selector}")
                continue

        log.warning(f"No working selector found for: {name}")
        return None

    async def _click_send(self) -> bool:
        """Try to click the send button using selector fallbacks."""
        selector = await self._find_selector(Selectors.SEND_BUTTON, "send button")
        if selector:
            await human_click(self._page, selector)
            log.debug("Send button clicked")
            return True
        return False

    async def _upload_files(self, file_paths: list[str]) -> None:
        """
        Upload files (images, PDFs, docs, etc.) to ChatGPT's input area.

        ChatGPT has a hidden <input type="file"> that accepts various file types.
        We set files on it directly (like drag-and-drop / file picker).
        """
        from pathlib import Path

        valid_paths = []
        for p in file_paths:
            path = Path(p)
            if path.exists() and path.is_file():
                valid_paths.append(str(path.resolve()))
            else:
                log.warning(f"File not found, skipping: {p}")

        if not valid_paths:
            log.warning("No valid files to upload")
            return

        log.info(f"Uploading {len(valid_paths)} file(s)...")

        # Find the file input element — ChatGPT has a hidden <input type="file">
        file_input = None
        for selector in Selectors.FILE_UPLOAD_INPUT:
            try:
                elements = await self._page.query_selector_all(selector)
                if elements:
                    file_input = elements[0]
                    log.debug(f"Found file input: {selector}")
                    break
            except Exception:
                continue

        if file_input:
            # Set files directly on the input element
            await file_input.set_input_files(valid_paths)
            log.info(f"Set {len(valid_paths)} file(s) on file input")
        else:
            # Fallback: use page.set_input_files with a broad selector
            log.info("No file input found via selectors, trying broad input[type=file]")
            try:
                await self._page.set_input_files("input[type='file']", valid_paths)
                log.info(f"Set {len(valid_paths)} file(s) via broad selector")
            except Exception as e:
                log.error(f"Failed to upload files: {e}")
                raise RuntimeError(f"Could not upload files: {e}")

        # Wait for files to be processed/attached (thumbnails/badges appear)
        await asyncio.sleep(3)
        # Additional wait if multiple files
        if len(valid_paths) > 1:
            await asyncio.sleep(len(valid_paths))
        log.info("File upload complete")

    def _extract_thread_id(self) -> str:
        """Extract the thread/conversation ID from the current URL."""
        url = self._page.url
        match = re.search(r"/c/([a-f0-9-]+)", url)
        return match.group(1) if match else ""
</file>

<file path="src/chatgpt/detector.py">
"""
Response completion detector.

Primary strategy: detect completion on the newest assistant turn only,
then extract from that same turn. This avoids one-turn lag where a
previous assistant response is returned for the current request.
"""

from __future__ import annotations

import asyncio
import re

from patchright.async_api import Page

from src.selectors import Selectors
from src.browser.human import idle_mouse_movement
from src.log import setup_logging
from src.config import Config

log = setup_logging("detector")


def normalize_assistant_text(text: str | None) -> str:
    """Normalize extracted assistant text for validation and comparisons."""
    cleaned = (text or "").strip()
    cleaned = re.sub(r"^ChatGPT said:\s*", "", cleaned, flags=re.IGNORECASE).strip()
    cleaned = re.sub(r"^You said:\s*", "", cleaned, flags=re.IGNORECASE).strip()
    return cleaned


def is_incomplete_response_text(text: str | None) -> bool:
    """
    Heuristic: true when text looks like transient "thinking/searching" UI status.
    """
    cleaned = normalize_assistant_text(text)
    if not cleaned:
        return True

    lower = cleaned.lower()
    markers = [
        "pro thinking",
        "thinking",
        "searching for",
        "searching the web",
        "analyzing",
        "working on",
        "please wait",
        "gathering",
    ]

    if any(marker in lower for marker in markers):
        if len(cleaned) < 240:
            return True
        if lower.startswith(("pro thinking", "thinking", "searching", "analyzing", "working on", "gathering")):
            return True

    return False


def _empty_snapshot() -> dict:
    return {
        "found": False,
        "index": -1,
        "signature": None,
        "hasCopyButton": False,
        "hasImage": False,
        "text": "",
    }


async def _latest_assistant_turn_snapshot(page: Page) -> dict:
    """
    Return metadata for the latest assistant turn (article ordered).

    signature format: "<article-index>:<stable-id>"
    where stable-id is best-effort from DOM attributes.
    """
    snapshot = await page.evaluate(
        """
        () => {
            const turns = Array.from(document.querySelectorAll('section[data-testid^="conversation-turn-"]'));

            for (let idx = turns.length - 1; idx >= 0; idx--) {
                const turn = turns[idx];
                const turnRole = turn.getAttribute('data-turn');
                const hasAssistantRole = turnRole === 'assistant' ||
                    Boolean(turn.querySelector('[data-message-author-role="assistant"]'));
                if (!hasAssistantRole) continue;

                const stableId =
                    turn.getAttribute('data-turn-id') ||
                    turn.getAttribute('data-testid') ||
                    turn.id ||
                    '';

                const hasCopyButton = Boolean(
                    turn.querySelector('button[data-testid="copy-turn-action-button"], button[aria-label="Copy message"], button[aria-label="Copy"]')
                );

                const hasImage = Boolean(
                    turn.querySelector('img[alt="Generated image"], div[id^="image-"] img, div[id^="image-"]')
                );

                const text = (turn.innerText || '').trim();

                return {
                    found: true,
                    index: idx,
                    signature: `${idx}:${stableId}`,
                    hasCopyButton,
                    hasImage,
                    text,
                };
            }

            return {
                found: false,
                index: -1,
                signature: null,
                hasCopyButton: false,
                hasImage: false,
                text: '',
            };
        }
        """
    )

    if not isinstance(snapshot, dict):
        return _empty_snapshot()

    normalized = _empty_snapshot()
    normalized.update(snapshot)
    return normalized


async def get_latest_assistant_turn_signature(page: Page) -> str | None:
    """Return signature for the latest assistant turn, if available."""
    snapshot = await _latest_assistant_turn_snapshot(page)
    signature = snapshot.get("signature")
    return signature if isinstance(signature, str) and signature else None


async def count_assistant_messages(page: Page) -> int:
    """Count assistant turns (article based, newest-UI friendly)."""
    count = await page.evaluate(
        """
        () => {
            const turns = Array.from(document.querySelectorAll('section[data-testid^="conversation-turn-"]'));

            let total = 0;
            for (const turn of turns) {
                const turnRole = turn.getAttribute('data-turn');
                const hasAssistantRole = turnRole === 'assistant' ||
                    Boolean(turn.querySelector('[data-message-author-role="assistant"]'));
                if (hasAssistantRole) total++;
            }
            return total;
        }
        """
    )
    return int(count or 0)


async def _detect_image_in_latest_turn(page: Page, previous_turn_signature: str | None = None) -> bool:
    """Check if the newest assistant turn (not previous turn) contains an image."""
    snapshot = await _latest_assistant_turn_snapshot(page)
    signature = snapshot.get("signature")
    is_new_turn = previous_turn_signature is None or (
        isinstance(signature, str) and signature != previous_turn_signature
    )
    return bool(is_new_turn and snapshot.get("hasImage"))


async def _count_copy_buttons(page: Page) -> int:
    """Count assistant turns that currently expose a copy button."""
    count = await page.evaluate(
        """
        () => {
            const turns = Array.from(document.querySelectorAll('section[data-testid^="conversation-turn-"]'));

            let total = 0;
            for (const turn of turns) {
                const turnRole = turn.getAttribute('data-turn');
                const hasAssistantRole = turnRole === 'assistant' ||
                    Boolean(turn.querySelector('[data-message-author-role="assistant"]'));
                if (!hasAssistantRole) continue;
                const hasCopyButton = turn.querySelector(
                    'button[data-testid="copy-turn-action-button"], button[aria-label="Copy message"], button[aria-label="Copy"]'
                );
                if (hasCopyButton) total++;
            }
            return total;
        }
        """
    )
    return int(count or 0)


async def _wait_for_new_turn_signature(
    page: Page,
    previous_turn_signature: str,
    timeout_ms: int,
) -> bool:
    """Wait until latest assistant-turn signature differs from previous one."""
    elapsed = 0
    poll_interval = Config.POLL_INTERVAL_MS / 1000
    heartbeat = 10

    while elapsed * 1000 < timeout_ms:
        snapshot = await _latest_assistant_turn_snapshot(page)
        signature = snapshot.get("signature")
        if isinstance(signature, str) and signature and signature != previous_turn_signature:
            log.debug(f"New assistant turn detected: {signature} (prev: {previous_turn_signature})")
            return True

        if elapsed > 0 and elapsed % heartbeat == 0:
            log.debug(f"Still waiting for new assistant turn... ({int(elapsed)}s)")

        await asyncio.sleep(poll_interval)
        elapsed += poll_interval

    log.debug("Timed out waiting for a new assistant-turn signature")
    return False


async def wait_for_response_complete(
    page: Page,
    expected_msg_count: int | None = None,
    timeout_ms: int | None = None,
    previous_turn_signature: str | None = None,
) -> bool:
    """
    Wait until ChatGPT finishes generating the current response.

    Uses latest-turn alignment to avoid returning stale previous-turn output.
    """
    timeout = timeout_ms or Config.RESPONSE_TIMEOUT
    log.info(f"Waiting for response (timeout: {timeout}ms)...")

    pre_copy_count = await _count_copy_buttons(page)
    log.debug(f"Copy buttons before send: {pre_copy_count}")

    if previous_turn_signature:
        log.debug(f"Previous assistant turn signature: {previous_turn_signature}")
        await _wait_for_new_turn_signature(page, previous_turn_signature, timeout_ms=30000)
    elif expected_msg_count is not None:
        log.debug(f"Waiting for assistant message #{expected_msg_count}...")
        waited = 0
        while waited < 30000:
            current_count = await count_assistant_messages(page)
            if current_count >= expected_msg_count:
                log.debug(f"Assistant message target reached (count: {current_count})")
                break
            await asyncio.sleep(0.5)
            waited += 500

    log.debug("Waiting for copy button or image on latest assistant turn...")
    completed = await _wait_for_copy_button_or_image(page, pre_copy_count, timeout, previous_turn_signature)
    if completed == "copy":
        log.info("Response complete — copy button appeared on latest turn")
        return True
    if completed == "image":
        log.info("Response complete — generated image detected on latest turn")
        return True

    log.info("Copy/image completion not detected, trying stop-button strategy...")
    try:
        result = await _wait_via_stop_button(page, timeout)
        if result:
            return True
    except Exception as e:
        log.debug(f"Stop button strategy failed: {e}")

    log.info("Falling back to text-stability detection...")
    try:
        return await _wait_via_text_stability(page, timeout, previous_turn_signature)
    except Exception as e:
        log.error(f"All strategies failed: {e}")
        return False


async def _wait_for_copy_button_or_image(
    page: Page,
    pre_count: int,
    timeout_ms: int,
    previous_turn_signature: str | None = None,
) -> str | None:
    """
    Wait for either copy-button readiness or generated image on the latest turn.

    Returns "copy", "image", or None if timed out.
    """
    elapsed = 0
    poll_interval = Config.POLL_INTERVAL_MS / 1000
    heartbeat = 10

    while elapsed * 1000 < timeout_ms:
        snapshot = await _latest_assistant_turn_snapshot(page)
        signature = snapshot.get("signature")
        is_new_turn = previous_turn_signature is None or (
            isinstance(signature, str) and signature != previous_turn_signature
        )

        if is_new_turn and snapshot.get("hasCopyButton"):
            current_count = await _count_copy_buttons(page)
            log.debug(
                f"Copy button detected on latest turn {signature} "
                f"(copy-buttons: {pre_count} -> {current_count})"
            )
            return "copy"

        has_image = await _detect_image_in_latest_turn(page, previous_turn_signature)
        if has_image:
            await asyncio.sleep(1.0)
            log.debug(f"Generated image detected on latest turn {signature}")
            return "image"

        if elapsed > 0 and elapsed % heartbeat == 0:
            log.debug(f"Still waiting for copy button or image... ({int(elapsed)}s)")
            await idle_mouse_movement(page)

        await asyncio.sleep(poll_interval)
        elapsed += poll_interval

    log.warning(f"Neither copy button nor image found after {int(elapsed)}s")
    return None


async def _wait_via_stop_button(page: Page, timeout_ms: int) -> bool:
    """Wait for stop button appear -> disappear cycle."""
    stop_selector = ", ".join(Selectors.STOP_BUTTON)
    log.debug("Waiting for stop button to appear...")

    try:
        await page.wait_for_selector(stop_selector, state="visible", timeout=15000)
        log.info("Stop button appeared — response is streaming")
    except Exception:
        log.debug("Stop button never appeared (short response or selector changed)")
        return False

    log.debug("Waiting for stop button to disappear...")
    heartbeat_interval = 10
    elapsed = 0

    while elapsed * 1000 < timeout_ms:
        try:
            await page.wait_for_selector(stop_selector, state="hidden", timeout=heartbeat_interval * 1000)
            log.info("Stop button disappeared — streaming done")
            return True
        except Exception:
            elapsed += heartbeat_interval
            log.debug(f"Still streaming... ({elapsed}s elapsed)")
            await idle_mouse_movement(page)

    log.warning(f"Timed out after {elapsed}s waiting for stop button")
    return False


async def _wait_via_text_stability(
    page: Page,
    timeout_ms: int,
    previous_turn_signature: str | None = None,
) -> bool:
    """
    Last resort: poll latest assistant-turn text and wait until stable.

    If previous_turn_signature is provided, ignores stabilization on that old turn.
    """
    stable_count = 0
    required_stable = 3
    last_text = ""
    elapsed = 0
    poll_interval = Config.POLL_INTERVAL_MS / 1000

    while elapsed * 1000 < timeout_ms:
        snapshot = await _latest_assistant_turn_snapshot(page)
        signature = snapshot.get("signature")
        text = snapshot.get("text") if isinstance(snapshot.get("text"), str) else ""

        if previous_turn_signature and signature == previous_turn_signature:
            stable_count = 0
            last_text = ""
            await asyncio.sleep(poll_interval)
            elapsed += poll_interval
            continue

        if text and text == last_text:
            stable_count += 1
            log.debug(f"Text stable ({stable_count}/{required_stable})")
            if stable_count >= required_stable:
                if is_incomplete_response_text(text) and not bool(snapshot.get("hasCopyButton")):
                    log.debug("Stable text looks like transient thinking status; continuing wait")
                    stable_count = 0
                    last_text = text
                    await asyncio.sleep(poll_interval)
                    elapsed += poll_interval
                    continue
                log.info("Response text stabilized — complete")
                return True
        else:
            stable_count = 0
            last_text = text

        await asyncio.sleep(poll_interval)
        elapsed += poll_interval

    log.warning(f"Text stability timed out after {int(elapsed)}s")
    return False


async def extract_last_response_via_copy(
    page: Page,
    previous_turn_signature: str | None = None,
) -> str:
    """
    Extract latest assistant response by clicking copy on the latest turn.

    Never intentionally copies from previous_turn_signature when provided.
    """
    log.debug("Attempting extraction via latest-turn copy button...")

    try:
        await page.context.grant_permissions(["clipboard-read", "clipboard-write"])

        if previous_turn_signature:
            await _wait_for_new_turn_signature(page, previous_turn_signature, timeout_ms=8000)

        pre_clipboard = await page.evaluate("navigator.clipboard.readText().catch(() => '')")
        await page.evaluate("navigator.clipboard.writeText('').catch(() => {})")

        click_result = await page.evaluate(
            """
            (previousSignature) => {
                const turns = Array.from(document.querySelectorAll('section[data-testid^="conversation-turn-"]'));

                for (let idx = turns.length - 1; idx >= 0; idx--) {
                    const turn = turns[idx];
                    const turnRole = turn.getAttribute('data-turn');
                    const hasAssistantRole = turnRole === 'assistant' ||
                        Boolean(turn.querySelector('[data-message-author-role="assistant"]'));
                    if (!hasAssistantRole) continue;

                    const stableId =
                        turn.getAttribute('data-turn-id') ||
                        turn.getAttribute('data-testid') ||
                        turn.id ||
                        '';
                    const signature = `${idx}:${stableId}`;

                    if (previousSignature && signature === previousSignature) {
                        return { clicked: false, reason: 'stale-turn', signature };
                    }

                    const btn = turn.querySelector(
                        'button[data-testid="copy-turn-action-button"], button[aria-label="Copy message"], button[aria-label="Copy"]'
                    );
                    if (!btn) {
                        return { clicked: false, reason: 'no-copy-button', signature };
                    }

                    btn.click();
                    return { clicked: true, reason: 'ok', signature };
                }

                return { clicked: false, reason: 'no-assistant-turn', signature: null };
            }
            """,
            previous_turn_signature,
        )

        if isinstance(click_result, dict) and click_result.get("clicked"):
            await asyncio.sleep(0.8)
            content = await page.evaluate("navigator.clipboard.readText().catch(() => '')")
            if content and content.strip() and content.strip() != str(pre_clipboard).strip():
                log.info(
                    "Extracted via copy button (latest-turn): "
                    f"{len(content)} chars, turn={click_result.get('signature')}"
                )
                return content.strip()
            log.debug("Clipboard unchanged/empty after latest-turn copy click")
        else:
            reason = click_result.get("reason") if isinstance(click_result, dict) else "unknown"
            log.debug(f"Latest-turn copy click not used: {reason}")

    except Exception as e:
        log.warning(f"Copy button extraction failed: {e}")

    log.info("Falling back to latest-turn DOM extraction...")
    return await _extract_via_dom(page, previous_turn_signature)


async def _extract_via_dom(
    page: Page,
    previous_turn_signature: str | None = None,
) -> str:
    """Fallback extraction: innerText from latest assistant turn only."""
    text = await page.evaluate(
        """
        (previousSignature) => {
            const turns = Array.from(document.querySelectorAll('section[data-testid^="conversation-turn-"]'));

            for (let idx = turns.length - 1; idx >= 0; idx--) {
                const turn = turns[idx];
                const turnRole = turn.getAttribute('data-turn');
                const hasAssistantRole = turnRole === 'assistant' ||
                    Boolean(turn.querySelector('[data-message-author-role="assistant"]'));
                if (!hasAssistantRole) continue;

                const stableId =
                    turn.getAttribute('data-turn-id') ||
                    turn.getAttribute('data-testid') ||
                    turn.id ||
                    '';
                const signature = `${idx}:${stableId}`;

                if (previousSignature && signature === previousSignature) {
                    return '';
                }

                return (turn.innerText || '').trim();
            }

            return '';
        }
        """,
        previous_turn_signature,
    )

    if text and str(text).strip():
        cleaned = normalize_assistant_text(str(text))
        if is_incomplete_response_text(cleaned):
            log.debug("Latest-turn DOM text looks incomplete/transient; waiting for a fuller reply")
            return ""
        log.debug(f"Extracted via DOM (latest-turn): {len(cleaned)} chars")
        return cleaned

    log.error("Could not extract any latest assistant response")
    return ""


# Keep old name as alias for backward compat
extract_last_response = extract_last_response_via_copy
</file>

<file path="src/chatgpt/image_handler.py">
"""
Image handler — detects, extracts, and downloads generated images.

When ChatGPT generates an image via DALL-E, the response contains:
- An <img> tag with the image URL (hosted on openai.com)
- A "Image created" text indicator
- An image title/alt text (description of what was generated)

This module:
1. Detects if the last assistant message contains generated images
2. Extracts image URLs and metadata
3. Downloads images to local disk
4. Returns ImageInfo objects with URLs and local paths
"""

from __future__ import annotations

import asyncio
import hashlib
import re
import time
from pathlib import Path
from urllib.parse import urlparse

from patchright.async_api import Page

from src.config import Config
from src.selectors import Selectors
from src.chatgpt.models import ImageInfo
from src.log import setup_logging

log = setup_logging("image_handler")


async def detect_images_in_response(page: Page) -> list[dict]:
    """
    Check the last conversation turn for generated images.

    ChatGPT DALL-E image responses do NOT use data-message-author-role.
    Instead, images appear inside an article turn with:
    - img[alt="Generated image"]
    - div[id^="image-"] containers
    - src from chatgpt.com/backend-api/estuary/content

    Returns a list of dicts: [{url, alt, title}, ...] or empty list.
    """
    result = await page.evaluate("""
        () => {
            const turns = document.querySelectorAll('section[data-testid^="conversation-turn-"]');
            if (turns.length === 0) return [];

            const lastTurn = turns[turns.length - 1];

            // Find generated images — primary: alt="Generated image"
            let images = lastTurn.querySelectorAll('img[alt="Generated image"]');

            // Fallback: images inside imagegen containers
            if (images.length === 0) {
                const containers = lastTurn.querySelectorAll('div[id^="image-"]');
                if (containers.length > 0) {
                    const imgSet = new Set();
                    for (const c of containers) {
                        const imgs = c.querySelectorAll('img');
                        for (const img of imgs) imgSet.add(img);
                    }
                    images = [...imgSet];
                }
            }

            // Fallback: any large image from chatgpt backend
            if (images.length === 0) {
                const allImgs = lastTurn.querySelectorAll('img');
                const large = [];
                for (const img of allImgs) {
                    const w = img.naturalWidth || img.width || 0;
                    const src = img.src || '';
                    if (w > 200 && (
                        src.includes('backend-api/estuary') ||
                        src.includes('chatgpt.com')
                    )) {
                        large.push(img);
                    }
                }
                images = large;
            }

            if (!images || images.length === 0) return [];

            // Deduplicate by src URL
            const seen = new Set();
            const results = [];

            for (const img of images) {
                const src = img.src || '';
                if (!src || seen.has(src)) continue;
                seen.add(src);

                const alt = img.alt || '';

                // Extract the image title from nearby text in the turn
                // ChatGPT shows "Creating image • Image Title" in a button/span
                let title = '';
                const buttons = lastTurn.querySelectorAll('button');
                for (const btn of buttons) {
                    const text = (btn.innerText || '').trim();
                    // Parse "Creating image • Title" or just "Title"
                    const bulletIdx = text.indexOf('•');
                    if (bulletIdx > -1) {
                        title = text.substring(bulletIdx + 1).trim();
                        break;
                    }
                }
                // Fallback: look for text spans in the turn
                if (!title) {
                    const spans = lastTurn.querySelectorAll(
                        'span.text-token-text-tertiary'
                    );
                    for (const span of spans) {
                        const t = (span.innerText || '').trim();
                        if (t.length > 5 && t.length < 200) {
                            title = t;
                            break;
                        }
                    }
                }

                results.push({ url: src, alt, title });
            }

            return results;
        }
    """)

    if result:
        log.info(f"Detected {len(result)} generated image(s) in response")
        for i, img in enumerate(result):
            log.debug(f"  Image {i+1}: alt='{img.get('alt', '')[:50]}', url={img.get('url', '')[:80]}...")
    else:
        log.debug("No generated images detected in response")

    return result or []


async def download_image(page: Page, url: str, filename_hint: str = "") -> str:
    """
    Download an image from a URL using the browser's fetch API.

    Uses the browser context so cookies/auth are preserved (required
    for OpenAI-hosted images that may need authentication).

    Returns the local file path.
    """
    Config.ensure_dirs()

    # Generate a filename from the URL or hint
    if filename_hint:
        # Clean the hint for use as filename
        safe_name = re.sub(r'[^\w\s-]', '', filename_hint)[:60].strip()
        safe_name = re.sub(r'\s+', '_', safe_name)
    else:
        # Use hash of URL as filename
        safe_name = hashlib.md5(url.encode()).hexdigest()[:12]

    # Add timestamp to avoid collisions
    ts = int(time.time())
    filename = f"{safe_name}_{ts}.png"
    local_path = Config.IMAGES_DIR / filename

    log.info(f"Downloading image to {local_path}...")

    try:
        # Use browser's fetch to download (preserves auth cookies)
        image_data = await page.evaluate("""
            async (url) => {
                try {
                    const response = await fetch(url);
                    if (!response.ok) return null;
                    const blob = await response.blob();
                    const reader = new FileReader();
                    return new Promise((resolve) => {
                        reader.onloadend = () => resolve(reader.result);
                        reader.readAsDataURL(blob);
                    });
                } catch (e) {
                    return null;
                }
            }
        """, url)

        if image_data and image_data.startswith("data:"):
            # Strip the data URL prefix to get raw base64
            import base64
            header, b64data = image_data.split(",", 1)

            # Detect actual format from MIME type
            if "png" in header:
                ext = ".png"
            elif "jpeg" in header or "jpg" in header:
                ext = ".jpg"
            elif "webp" in header:
                ext = ".webp"
            else:
                ext = ".png"

            # Update filename with correct extension
            filename = f"{safe_name}_{ts}{ext}"
            local_path = Config.IMAGES_DIR / filename

            raw_bytes = base64.b64decode(b64data)
            local_path.write_bytes(raw_bytes)

            size_kb = len(raw_bytes) / 1024
            log.info(f"Image saved: {local_path} ({size_kb:.1f} KB)")
            return str(local_path)

        else:
            log.warning("Failed to fetch image data via browser")

    except Exception as e:
        log.error(f"Image download failed: {e}", exc_info=True)

    # Fallback: try using the page to download via navigation
    # (less reliable but works for some cases)
    try:
        import urllib.request
        urllib.request.urlretrieve(url, str(local_path))
        log.info(f"Image saved via urllib: {local_path}")
        return str(local_path)
    except Exception as e2:
        log.error(f"Fallback download also failed: {e2}")

    return ""


async def extract_images_from_response(page: Page) -> list[ImageInfo]:
    """
    Full pipeline: detect images in the last response, download them,
    and return ImageInfo objects with both URLs and local paths.
    """
    raw_images = await detect_images_in_response(page)

    if not raw_images:
        return []

    image_infos = []
    for img_data in raw_images:
        url = img_data.get("url", "")
        alt = img_data.get("alt", "")
        title = img_data.get("title", "")

        # Download the image
        hint = alt or title or "chatgpt_image"
        local_path = await download_image(page, url, filename_hint=hint)

        image_infos.append(ImageInfo(
            url=url,
            alt=alt,
            local_path=local_path,
            prompt_title=title,
        ))

    log.info(f"Processed {len(image_infos)} image(s)")
    return image_infos
</file>

<file path="src/chatgpt/models.py">
"""
Data models for ChatGPT interactions.
"""

from __future__ import annotations

from datetime import datetime
from pydantic import BaseModel, Field


class ImageInfo(BaseModel):
    """Metadata for a generated image."""
    url: str = Field(description="Original image URL from ChatGPT/DALL-E")
    alt: str = Field(default="", description="Alt text / image description")
    local_path: str = Field(default="", description="Local file path after download")
    prompt_title: str = Field(default="", description="Image generation title shown by ChatGPT")


class Message(BaseModel):
    """A single message in a conversation."""
    role: str = Field(description="'user' or 'assistant'")
    content: str = Field(description="Message text content")
    timestamp: datetime = Field(default_factory=datetime.now)
    images: list[ImageInfo] = Field(default_factory=list, description="Images in this message")


class ChatResponse(BaseModel):
    """Response from a chat interaction."""
    message: str = Field(description="Assistant's response text")
    thread_id: str = Field(default="", description="Conversation thread ID from URL")
    response_time_ms: int = Field(default=0, description="Time taken for response in ms")
    images: list[ImageInfo] = Field(default_factory=list, description="Generated images")
    has_images: bool = Field(default=False, description="Whether the response contains images")


class Thread(BaseModel):
    """A conversation thread."""
    id: str = Field(description="Thread ID (from URL /c/{id})")
    title: str = Field(default="", description="Thread title from sidebar")
    url: str = Field(default="", description="Full URL")
    messages: list[Message] = Field(default_factory=list)
    created_at: datetime = Field(default_factory=datetime.now)
</file>

<file path="src/claude/__init__.py">
from src.claude.client import ClaudeClient
</file>

<file path="src/claude/client.py">
"""
Claude client — core interaction logic for claude.ai.

Sends messages, waits for responses, manages conversations.
Handles selector fallbacks and integrates human-like behavior.

Same interface as ChatGPTClient so the API layer is provider-agnostic.
"""

from __future__ import annotations

import asyncio
import re
import time

from patchright.async_api import Page

from src.config import Config
from src.claude.selectors import ClaudeSelectors
from src.browser.human import human_type, human_click, thinking_pause, random_delay
from src.claude.detector import (
    wait_for_response_complete,
    extract_last_response_via_copy,
    count_assistant_messages,
    get_latest_assistant_turn_signature,
    is_incomplete_response_text,
)
from src.chatgpt.models import ChatResponse
from src.log import setup_logging

log = setup_logging("claude_client")


class ClaudeClient:
    """
    High-level client for interacting with the Claude web interface.

    Requires a Playwright Page that is already logged in and on claude.ai.
    Same interface as ChatGPTClient for provider-agnostic API usage.
    """

    def __init__(self, page: Page) -> None:
        self._page = page

    @property
    def page(self) -> Page:
        return self._page

    # ── Core: Send & Receive ────────────────────────────────────

    async def send_message(self, text: str, image_paths: list[str] | None = None, file_paths: list[str] | None = None) -> ChatResponse:
        """
        Send a message to Claude and wait for the complete response.

        Args:
            text: The message text to send.
            image_paths: Optional list of local file paths to images to attach.
            file_paths: Optional list of local file paths to non-image files.

        Returns ChatResponse with the assistant's reply and metadata.
        """
        all_attachments = (image_paths or []) + (file_paths or [])
        log.info(f"Sending message ({len(text)} chars, {len(all_attachments)} attachments): {text[:80]}...")
        start_time = time.time()

        # 0. Count existing assistant messages so we know when a new one appears
        pre_count = await count_assistant_messages(self._page)
        pre_turn_signature = await get_latest_assistant_turn_signature(self._page)
        log.debug(f"Assistant messages before send: {pre_count}")
        log.debug(f"Latest assistant turn before send: {pre_turn_signature}")

        # 1. Brief pause (human would take a moment to start typing)
        await random_delay(250, 700)

        # 1.5. Upload files/images if provided
        if all_attachments:
            await self._upload_files(all_attachments)

        # 2. Find the chat input
        input_selector = await self._find_selector(ClaudeSelectors.CHAT_INPUT, "chat input")
        if not input_selector:
            raise RuntimeError("Could not find chat input element")

        # 3. Paste the message
        await human_type(self._page, input_selector, text)

        # Small pause after pasting
        await random_delay(150, 350)

        # 4. Send the message
        sent = await self._click_send()
        if not sent:
            # Fallback: try pressing Enter
            log.info("Send button not found, trying Enter key")
            await self._page.keyboard.press("Enter")

        # 5. Wait for response
        log.info("Waiting for Claude response...")
        expected_count = pre_count + 1
        completed = await wait_for_response_complete(
            self._page,
            expected_msg_count=expected_count,
            previous_turn_signature=pre_turn_signature,
        )

        if not completed:
            log.warning("Response may not be complete (timeout)")

        # Small buffer after completion to let DOM settle
        await asyncio.sleep(0.5)

        # 6. Extract text content (Claude doesn't generate images like DALL-E)
        response_text = await extract_last_response_via_copy(
            self._page,
            previous_turn_signature=pre_turn_signature,
        )

        # If we only captured a transient status, retry
        if is_incomplete_response_text(response_text):
            log.warning("Extracted text looks incomplete/transient; retrying for final answer")
            for attempt in range(1, 3):
                await asyncio.sleep(4)
                await wait_for_response_complete(
                    self._page,
                    timeout_ms=90000,
                    previous_turn_signature=pre_turn_signature,
                )
                retry_text = await extract_last_response_via_copy(
                    self._page,
                    previous_turn_signature=pre_turn_signature,
                )

                if retry_text and not is_incomplete_response_text(retry_text):
                    response_text = retry_text
                    log.info(f"Recovered final response text on retry {attempt}")
                    break

                if retry_text:
                    response_text = retry_text
                log.warning(f"Retry {attempt} still incomplete/transient")

        elapsed_ms = int((time.time() - start_time) * 1000)
        thread_id = self._extract_thread_id()

        log.info(
            f"Response received ({elapsed_ms}ms, {len(response_text)} chars): "
            f"{response_text[:80]}..."
        )

        return ChatResponse(
            message=response_text,
            thread_id=thread_id,
            response_time_ms=elapsed_ms,
            images=[],
            has_images=False,
        )

    # ── Navigation ──────────────────────────────────────────────

    async def new_chat(self) -> None:
        """Start a new conversation by navigating to /new."""
        log.info("Starting new chat...")
        url = Config.CLAUDE_URL.rstrip("/") + "/new"
        await self._page.goto(url, wait_until="domcontentloaded")
        await asyncio.sleep(1.5)

        # Wait for the chat input to be visible
        for selector in ClaudeSelectors.CHAT_INPUT:
            try:
                await self._page.wait_for_selector(selector, timeout=10000, state="visible")
                log.debug(f"Chat input ready: {selector}")
                break
            except Exception:
                continue

        await random_delay(300, 600)
        log.info("New chat started (navigated to /new)")

    async def navigate_to_thread(self, thread_id: str) -> None:
        """Navigate to an existing conversation thread."""
        url = f"{Config.CLAUDE_URL.rstrip('/')}/chat/{thread_id}"
        log.info(f"Navigating to thread: {thread_id}")
        await self._page.goto(url, wait_until="domcontentloaded")
        await random_delay(800, 1500)
        log.info(f"Thread {thread_id} loaded")

    async def get_current_thread_url(self) -> str:
        """Get the current page URL (contains thread ID if in a conversation)."""
        return self._page.url

    # ── Sidebar ─────────────────────────────────────────────────

    async def list_threads(self) -> list[dict]:
        """
        Scrape the sidebar for recent conversation threads.

        Returns a list of dicts: [{id, title, url}, ...]
        """
        threads = []
        for selector in ClaudeSelectors.SIDEBAR_THREAD_LINKS:
            try:
                elements = await self._page.query_selector_all(selector)
                for el in elements:
                    href = await el.get_attribute("href") or ""
                    title = (await el.inner_text()).strip()
                    # Claude uses /chat/{uuid}
                    match = re.search(r"/chat/([a-f0-9-]+)", href)
                    if match:
                        threads.append({
                            "id": match.group(1),
                            "title": title,
                            "url": f"{Config.CLAUDE_URL.rstrip('/')}{href}",
                        })
                if threads:
                    break
            except Exception as e:
                log.debug(f"Sidebar scrape with {selector} failed: {e}")

        log.info(f"Found {len(threads)} threads in sidebar")
        return threads

    # ── Private Helpers ─────────────────────────────────────────

    async def _find_selector(self, selectors: list[str], name: str) -> str | None:
        """Try each selector in the fallback list. Return the first one that matches."""
        for selector in selectors:
            try:
                el = await self._page.wait_for_selector(
                    selector,
                    timeout=Config.SELECTOR_TIMEOUT,
                    state="visible",
                )
                if el:
                    log.debug(f"Found {name} via: {selector}")
                    return selector
            except Exception:
                log.debug(f"Selector miss for {name}: {selector}")
                continue

        log.warning(f"No working selector found for: {name}")
        return None

    async def _click_send(self) -> bool:
        """Try to click the send button using selector fallbacks."""
        selector = await self._find_selector(ClaudeSelectors.SEND_BUTTON, "send button")
        if selector:
            await human_click(self._page, selector)
            log.debug("Send button clicked")
            return True
        return False

    async def _upload_files(self, file_paths: list[str]) -> None:
        """Upload files to Claude's input area."""
        from pathlib import Path

        valid_paths = []
        for p in file_paths:
            path = Path(p)
            if path.exists() and path.is_file():
                valid_paths.append(str(path.resolve()))
            else:
                log.warning(f"File not found, skipping: {p}")

        if not valid_paths:
            log.warning("No valid files to upload")
            return

        log.info(f"Uploading {len(valid_paths)} file(s)...")

        # Find the file input element
        file_input = None
        for selector in ClaudeSelectors.FILE_UPLOAD_INPUT:
            try:
                elements = await self._page.query_selector_all(selector)
                if elements:
                    file_input = elements[0]
                    log.debug(f"Found file input: {selector}")
                    break
            except Exception:
                continue

        if file_input:
            await file_input.set_input_files(valid_paths)
            log.info(f"Set {len(valid_paths)} file(s) on file input")
        else:
            log.info("No file input found via selectors, trying broad input[type=file]")
            try:
                await self._page.set_input_files("input[type='file']", valid_paths)
                log.info(f"Set {len(valid_paths)} file(s) via broad selector")
            except Exception as e:
                log.error(f"Failed to upload files: {e}")
                raise RuntimeError(f"Could not upload files: {e}")

        # Wait for files to be processed
        await asyncio.sleep(3)
        if len(valid_paths) > 1:
            await asyncio.sleep(len(valid_paths))
        log.info("File upload complete")

    def _extract_thread_id(self) -> str:
        """Extract the thread/conversation ID from the current URL."""
        url = self._page.url
        # Claude uses /chat/{uuid}
        match = re.search(r"/chat/([a-f0-9-]+)", url)
        return match.group(1) if match else ""
</file>

<file path="src/claude/detector.py">
"""
Response completion detector for Claude.ai.

Primary strategy: detect completion using Claude's data-is-streaming attribute
and copy-button appearance, then extract from the latest assistant turn.
"""

from __future__ import annotations

import asyncio
import re

from patchright.async_api import Page

from src.claude.selectors import ClaudeSelectors
from src.browser.human import idle_mouse_movement
from src.log import setup_logging
from src.config import Config

log = setup_logging("claude_detector")


def normalize_assistant_text(text: str | None) -> str:
    """Normalize extracted assistant text for validation and comparisons."""
    cleaned = (text or "").strip()
    # Claude prefixes with "Claude responded: <title>" in sr-only heading
    cleaned = re.sub(r"^Claude responded:\s*.*?\n", "", cleaned, flags=re.IGNORECASE).strip()
    return cleaned


def is_incomplete_response_text(text: str | None) -> bool:
    """Heuristic: true when text looks like transient thinking/analyzing status."""
    cleaned = normalize_assistant_text(text)
    if not cleaned:
        return True

    lower = cleaned.lower()
    markers = [
        "thinking",
        "analyzing",
        "searching",
        "working on",
        "please wait",
    ]

    if any(marker in lower for marker in markers):
        if len(cleaned) < 240:
            return True
        if lower.startswith(tuple(markers)):
            return True

    return False


def _empty_snapshot() -> dict:
    return {
        "found": False,
        "index": -1,
        "signature": None,
        "hasCopyButton": False,
        "isStreaming": True,
        "text": "",
    }


async def _latest_assistant_turn_snapshot(page: Page) -> dict:
    """
    Return metadata for the latest assistant turn in Claude.ai.

    Claude wraps each assistant response in a div with data-is-streaming attribute.
    The sr-only h2 reads "Claude responded: <title>".
    """
    snapshot = await page.evaluate(
        """
        () => {
            // Claude assistant turns are divs with data-is-streaming attribute
            const turns = Array.from(document.querySelectorAll('div[data-is-streaming]'));

            if (turns.length === 0) {
                return {
                    found: false,
                    index: -1,
                    signature: null,
                    hasCopyButton: false,
                    isStreaming: true,
                    text: '',
                };
            }

            const last = turns[turns.length - 1];
            const idx = turns.length - 1;

            // Build a stable signature from the turn
            const h2 = last.querySelector('h2.sr-only');
            const h2Text = h2 ? h2.innerText.trim() : '';
            const signature = `${idx}:${h2Text.substring(0, 50)}`;

            const isStreaming = last.getAttribute('data-is-streaming') === 'true';

            const hasCopyButton = Boolean(
                last.querySelector('button[data-testid="action-bar-copy"], button[aria-label="Copy"]')
            );

            // Get the actual response text (not the sr-only heading)
            const responseDiv = last.querySelector('.font-claude-response');
            const text = responseDiv
                ? responseDiv.innerText.trim()
                : last.innerText.trim();

            return {
                found: true,
                index: idx,
                signature,
                hasCopyButton,
                isStreaming,
                text,
            };
        }
        """
    )

    if not isinstance(snapshot, dict):
        return _empty_snapshot()

    normalized = _empty_snapshot()
    normalized.update(snapshot)
    return normalized


async def get_latest_assistant_turn_signature(page: Page) -> str | None:
    """Return signature for the latest assistant turn, if available."""
    snapshot = await _latest_assistant_turn_snapshot(page)
    signature = snapshot.get("signature")
    return signature if isinstance(signature, str) and signature else None


async def count_assistant_messages(page: Page) -> int:
    """Count assistant turns (based on data-is-streaming divs)."""
    count = await page.evaluate(
        """
        () => {
            return document.querySelectorAll('div[data-is-streaming]').length;
        }
        """
    )
    return int(count or 0)


async def _count_copy_buttons(page: Page) -> int:
    """Count assistant turns that currently expose a copy button."""
    count = await page.evaluate(
        """
        () => {
            const turns = Array.from(document.querySelectorAll('div[data-is-streaming]'));
            let total = 0;
            for (const turn of turns) {
                if (turn.querySelector('button[data-testid="action-bar-copy"], button[aria-label="Copy"]')) {
                    total++;
                }
            }
            return total;
        }
        """
    )
    return int(count or 0)


async def _wait_for_new_turn_signature(
    page: Page,
    previous_turn_signature: str,
    timeout_ms: int,
) -> bool:
    """Wait until latest assistant-turn signature differs from previous one."""
    elapsed = 0
    poll_interval = Config.POLL_INTERVAL_MS / 1000
    heartbeat = 10

    while elapsed * 1000 < timeout_ms:
        snapshot = await _latest_assistant_turn_snapshot(page)
        signature = snapshot.get("signature")
        if isinstance(signature, str) and signature and signature != previous_turn_signature:
            log.debug(f"New assistant turn detected: {signature} (prev: {previous_turn_signature})")
            return True

        if elapsed > 0 and elapsed % heartbeat == 0:
            log.debug(f"Still waiting for new assistant turn... ({int(elapsed)}s)")

        await asyncio.sleep(poll_interval)
        elapsed += poll_interval

    log.debug("Timed out waiting for a new assistant-turn signature")
    return False


async def wait_for_response_complete(
    page: Page,
    expected_msg_count: int | None = None,
    timeout_ms: int | None = None,
    previous_turn_signature: str | None = None,
) -> bool:
    """
    Wait until Claude finishes generating the current response.

    Primary signal: data-is-streaming attribute changes from "true" to "false".
    Secondary: copy button appears on the latest turn.
    """
    timeout = timeout_ms or Config.RESPONSE_TIMEOUT
    log.info(f"Waiting for response (timeout: {timeout}ms)...")

    pre_copy_count = await _count_copy_buttons(page)
    log.debug(f"Copy buttons before send: {pre_copy_count}")

    # Wait for new turn to appear
    if previous_turn_signature:
        log.debug(f"Previous assistant turn signature: {previous_turn_signature}")
        await _wait_for_new_turn_signature(page, previous_turn_signature, timeout_ms=30000)
    elif expected_msg_count is not None:
        log.debug(f"Waiting for assistant message #{expected_msg_count}...")
        waited = 0
        while waited < 30000:
            current_count = await count_assistant_messages(page)
            if current_count >= expected_msg_count:
                log.debug(f"Assistant message target reached (count: {current_count})")
                break
            await asyncio.sleep(0.5)
            waited += 500

    # Strategy 1: Wait for streaming to complete (most reliable for Claude)
    log.debug("Waiting for streaming to complete (data-is-streaming='false')...")
    completed = await _wait_for_streaming_complete(page, timeout, previous_turn_signature)
    if completed:
        log.info("Response complete — streaming finished")
        return True

    # Strategy 2: Wait for copy button
    log.debug("Waiting for copy button on latest assistant turn...")
    copy_detected = await _wait_for_copy_button(page, pre_copy_count, timeout, previous_turn_signature)
    if copy_detected:
        log.info("Response complete — copy button appeared on latest turn")
        return True

    # Strategy 3: Text stability fallback
    log.info("Falling back to text-stability detection...")
    try:
        return await _wait_via_text_stability(page, timeout, previous_turn_signature)
    except Exception as e:
        log.error(f"All strategies failed: {e}")
        return False


async def _wait_for_streaming_complete(
    page: Page,
    timeout_ms: int,
    previous_turn_signature: str | None = None,
) -> bool:
    """
    Wait for Claude's data-is-streaming attribute to become "false".
    This is the primary and most reliable completion signal for Claude.
    """
    elapsed = 0
    poll_interval = Config.POLL_INTERVAL_MS / 1000
    heartbeat = 10

    while elapsed * 1000 < timeout_ms:
        snapshot = await _latest_assistant_turn_snapshot(page)
        signature = snapshot.get("signature")
        is_new_turn = previous_turn_signature is None or (
            isinstance(signature, str) and signature != previous_turn_signature
        )

        if is_new_turn and snapshot.get("found") and not snapshot.get("isStreaming"):
            # Double-check by waiting a moment for DOM to settle
            await asyncio.sleep(0.5)
            verify = await _latest_assistant_turn_snapshot(page)
            if not verify.get("isStreaming"):
                log.debug(f"Streaming complete on turn {signature}")
                return True

        if elapsed > 0 and elapsed % heartbeat == 0:
            log.debug(f"Still streaming... ({int(elapsed)}s)")
            await idle_mouse_movement(page)

        await asyncio.sleep(poll_interval)
        elapsed += poll_interval

    log.warning(f"Streaming did not complete after {int(elapsed)}s")
    return False


async def _wait_for_copy_button(
    page: Page,
    pre_count: int,
    timeout_ms: int,
    previous_turn_signature: str | None = None,
) -> bool:
    """Wait for copy button to appear on the latest assistant turn."""
    elapsed = 0
    poll_interval = Config.POLL_INTERVAL_MS / 1000

    while elapsed * 1000 < timeout_ms:
        snapshot = await _latest_assistant_turn_snapshot(page)
        signature = snapshot.get("signature")
        is_new_turn = previous_turn_signature is None or (
            isinstance(signature, str) and signature != previous_turn_signature
        )

        if is_new_turn and snapshot.get("hasCopyButton"):
            current_count = await _count_copy_buttons(page)
            log.debug(
                f"Copy button detected on latest turn {signature} "
                f"(copy-buttons: {pre_count} -> {current_count})"
            )
            return True

        await asyncio.sleep(poll_interval)
        elapsed += poll_interval

    return False


async def _wait_via_text_stability(
    page: Page,
    timeout_ms: int,
    previous_turn_signature: str | None = None,
) -> bool:
    """
    Last resort: poll latest assistant-turn text and wait until stable.
    """
    stable_count = 0
    required_stable = 3
    last_text = ""
    elapsed = 0
    poll_interval = Config.POLL_INTERVAL_MS / 1000

    while elapsed * 1000 < timeout_ms:
        snapshot = await _latest_assistant_turn_snapshot(page)
        signature = snapshot.get("signature")
        text = snapshot.get("text") if isinstance(snapshot.get("text"), str) else ""

        if previous_turn_signature and signature == previous_turn_signature:
            stable_count = 0
            last_text = ""
            await asyncio.sleep(poll_interval)
            elapsed += poll_interval
            continue

        if text and text == last_text:
            stable_count += 1
            log.debug(f"Text stable ({stable_count}/{required_stable})")
            if stable_count >= required_stable:
                if is_incomplete_response_text(text):
                    log.debug("Stable text looks like transient status; continuing wait")
                    stable_count = 0
                    last_text = text
                    await asyncio.sleep(poll_interval)
                    elapsed += poll_interval
                    continue
                log.info("Response text stabilized — complete")
                return True
        else:
            stable_count = 0
            last_text = text

        await asyncio.sleep(poll_interval)
        elapsed += poll_interval

    log.warning(f"Text stability timed out after {int(elapsed)}s")
    return False


async def extract_last_response_via_copy(
    page: Page,
    previous_turn_signature: str | None = None,
) -> str:
    """
    Extract latest assistant response by clicking copy on the latest turn.
    """
    log.debug("Attempting extraction via latest-turn copy button...")

    try:
        await page.context.grant_permissions(["clipboard-read", "clipboard-write"])

        if previous_turn_signature:
            await _wait_for_new_turn_signature(page, previous_turn_signature, timeout_ms=8000)

        pre_clipboard = await page.evaluate("navigator.clipboard.readText().catch(() => '')")
        await page.evaluate("navigator.clipboard.writeText('').catch(() => {})")

        click_result = await page.evaluate(
            """
            (previousSignature) => {
                const turns = Array.from(document.querySelectorAll('div[data-is-streaming]'));

                for (let idx = turns.length - 1; idx >= 0; idx--) {
                    const turn = turns[idx];

                    const h2 = turn.querySelector('h2.sr-only');
                    const h2Text = h2 ? h2.innerText.trim() : '';
                    const signature = `${idx}:${h2Text.substring(0, 50)}`;

                    if (previousSignature && signature === previousSignature) {
                        return { clicked: false, reason: 'stale-turn', signature };
                    }

                    const btn = turn.querySelector(
                        'button[data-testid="action-bar-copy"], button[aria-label="Copy"]'
                    );
                    if (!btn) {
                        return { clicked: false, reason: 'no-copy-button', signature };
                    }

                    btn.click();
                    return { clicked: true, reason: 'ok', signature };
                }

                return { clicked: false, reason: 'no-assistant-turn', signature: null };
            }
            """,
            previous_turn_signature,
        )

        if isinstance(click_result, dict) and click_result.get("clicked"):
            await asyncio.sleep(0.8)
            content = await page.evaluate("navigator.clipboard.readText().catch(() => '')")
            if content and content.strip() and content.strip() != str(pre_clipboard).strip():
                log.info(
                    "Extracted via copy button (latest-turn): "
                    f"{len(content)} chars, turn={click_result.get('signature')}"
                )
                return content.strip()
            log.debug("Clipboard unchanged/empty after latest-turn copy click")
        else:
            reason = click_result.get("reason") if isinstance(click_result, dict) else "unknown"
            log.debug(f"Latest-turn copy click not used: {reason}")

    except Exception as e:
        log.warning(f"Copy button extraction failed: {e}")

    log.info("Falling back to latest-turn DOM extraction...")
    return await _extract_via_dom(page, previous_turn_signature)


async def _extract_via_dom(
    page: Page,
    previous_turn_signature: str | None = None,
) -> str:
    """Fallback extraction: innerText from latest assistant turn only."""
    text = await page.evaluate(
        """
        (previousSignature) => {
            const turns = Array.from(document.querySelectorAll('div[data-is-streaming]'));

            for (let idx = turns.length - 1; idx >= 0; idx--) {
                const turn = turns[idx];

                const h2 = turn.querySelector('h2.sr-only');
                const h2Text = h2 ? h2.innerText.trim() : '';
                const signature = `${idx}:${h2Text.substring(0, 50)}`;

                if (previousSignature && signature === previousSignature) {
                    return '';
                }

                // Get text from the font-claude-response div
                const responseDiv = turn.querySelector('.font-claude-response');
                if (responseDiv) {
                    return responseDiv.innerText.trim();
                }

                // Fallback to full turn text, stripping sr-only heading
                const full = turn.innerText.trim();
                return full.replace(/^Claude responded:.*?\\n/i, '').trim();
            }
            return '';
        }
        """,
        previous_turn_signature,
    )
    cleaned = normalize_assistant_text(text or "")
    if cleaned:
        log.info(f"Extracted via DOM: {len(cleaned)} chars")
    else:
        log.warning("DOM extraction returned empty text")
    return cleaned
</file>

<file path="src/claude/selectors.py">
"""
Centralized DOM selectors for Claude.ai.

All selectors live here so when Claude updates their UI, we only
change this one file. Each entry is a list of fallback selectors —
try them in order until one matches.
"""

from __future__ import annotations


class ClaudeSelectors:
    """CSS / Playwright selectors for claude.ai UI elements."""

    # ── Chat input ──────────────────────────────────────────────
    # Claude uses a tiptap/ProseMirror contenteditable div
    CHAT_INPUT = [
        "div[data-testid='chat-input']",
        "div[contenteditable='true'][role='textbox']",
        "div.ProseMirror[contenteditable='true']",
        "div[contenteditable='true']",
    ]

    # ── Send button ─────────────────────────────────────────────
    # When text is entered, Claude's voice-mode button morphs into send.
    SEND_BUTTON = [
        "button[data-testid='send-message']",
        "button[aria-label='Send message']",
        "button[aria-label='Send Message']",
        "fieldset button[type='submit']",
    ]

    # ── Assistant response messages ─────────────────────────────
    # Claude wraps each assistant turn in a div with data-is-streaming.
    # The inner text content has class font-claude-response.
    ASSISTANT_MESSAGE = [
        "div[data-is-streaming]",
        "div.font-claude-response",
    ]

    # ── Streaming / stop button (visible while generating) ─────
    STOP_BUTTON = [
        "button[data-testid='stop-button']",
        "button[aria-label='Stop response']",
        "button[aria-label='Stop']",
    ]

    # ── New chat ────────────────────────────────────────────────
    NEW_CHAT_BUTTON = [
        "a[href='/new']",
    ]

    # ── Sidebar conversation links ──────────────────────────────
    # Claude uses /chat/{uuid} URL pattern
    SIDEBAR_THREAD_LINKS = [
        "a[href^='/chat/']",
    ]

    # ── Login page detection ────────────────────────────────────
    # If user-menu-button exists, user is logged in.
    # If login elements appear, user is logged out.
    LOGIN_INDICATORS = [
        "button:has-text('Log in')",
        "button:has-text('Sign in')",
        "a:has-text('Log in')",
        "a:has-text('Sign in')",
    ]

    # ── Logged-in indicator (user menu) ─────────────────────────
    LOGGED_IN_INDICATORS = [
        "button[data-testid='user-menu-button']",
    ]

    # ── Markdown content inside assistant message ───────────────
    ASSISTANT_MARKDOWN = [
        "div.font-claude-response",
        "div.font-claude-response .standard-markdown",
        "p.font-claude-response-body",
    ]

    # ── User message ────────────────────────────────────────────
    USER_MESSAGE = [
        "div[data-testid='user-message']",
    ]

    # ── Copy button (appears on each completed assistant message) ──
    COPY_BUTTON = [
        "button[data-testid='action-bar-copy']",
        "button[aria-label='Copy']",
    ]

    # ── Retry / regenerate button ───────────────────────────────
    POST_RESPONSE_BUTTONS = [
        "button[data-testid='action-bar-retry']",
        "button[aria-label='Retry']",
    ]

    # ── File / attachment upload input ────────────────────────────
    FILE_UPLOAD_INPUT = [
        "input#chat-input-file-upload-onpage",
        "input[data-testid='file-upload']",
        "input[type='file']",
    ]

    # Attach / upload button (opens file picker)
    ATTACH_BUTTON = [
        "button[aria-label='Add files, connectors, and more']",
        "button[aria-label='Attach files']",
    ]

    # ── Model selector ──────────────────────────────────────────
    MODEL_SELECTOR = [
        "button[data-testid='model-selector-dropdown']",
    ]

    # ── Generated images (Claude Artifacts, etc.) ─────────────
    ASSISTANT_IMAGE: list[str] = []
    IMAGE_CONTAINER: list[str] = []
    IMAGE_DOWNLOAD_BUTTON: list[str] = []
</file>

<file path="src/cli/__init__.py">
"""CATGPT CLI - Textual TUI for ChatGPT browser automation."""

from src.cli.app import CatGPTApp
</file>

<file path="src/cli/app.py">
"""
CATGPT — OpenAI-Compatible Terminal Chat

A beautiful full-screen TUI that talks to CatGPT Gateway (or any OpenAI-compatible API).
Uses the standard openai Python SDK — no browser management, no Playwright.

Configuration (env vars or CLI flags):
  CATGPT_API_URL    API base URL  (default: http://localhost:8000/v1)
  CATGPT_API_KEY    Bearer token  (default: dummy123)
  CATGPT_MODEL      Model name    (default: catgpt-browser)

Commands:
  /help              Show this help
  /new               Start a fresh conversation (clears history)
  /clear             Clear the display (history preserved)
  /system <text>     Set a system prompt for the session
  /model <name>      Switch to a different model
  /history           Show all conversation turns
  /export [file]     Export conversation to a markdown file
  /status            Show API config and session info
  /exit              Quit

Shortcuts:
  Ctrl+N   New conversation    Ctrl+E   Export
  Ctrl+L   Clear display       Ctrl+C   Quit
"""

from __future__ import annotations

import os
from datetime import datetime
from pathlib import Path

import typer
from rich.markdown import Markdown

from textual import work
from textual.app import App, ComposeResult
from textual.binding import Binding
from textual.containers import Center, ScrollableContainer, Vertical
from textual.reactive import reactive
from textual.screen import Screen
from textual.widget import Widget
from textual.widgets import Footer, Header, Input, Static

from src.log import suppress_console_logs

suppress_console_logs()

from src.log import setup_logging

log = setup_logging("cli", log_file="cli.log")
cli = typer.Typer(no_args_is_help=False, add_completion=False)

# ── Constants ────────────────────────────────────────────────────
VERSION = "3.0.0"
APP_NAME = "CATGPT"
APP_TAGLINE = "OpenAI-Compatible Terminal Chat"

THINKING_FRAMES = ["◐", "◓", "◑", "◒"]

CAT_ART = """\
      /\\_/\\
     ( ● . ● )
      > △ <
     /|   |\\
    (_|   |_)"""

LOGO_TEXT = """\
 ██████╗  █████╗ ████████╗ ██████╗ ██████╗ ████████╗
██╔════╝ ██╔══██╗╚══██╔══╝██╔════╝ ██╔══██╗╚══██╔══╝
██║      ███████║   ██║   ██║  ███╗██████╔╝   ██║
██║      ██╔══██║   ██║   ██║   ██║██╔═══╝    ██║
╚██████╗ ██║  ██║   ██║   ╚██████╔╝██║        ██║
 ╚═════╝ ╚═╝  ╚═╝   ╚═╝    ╚═════╝ ╚═╝        ╚═╝"""

WELCOME_TEMPLATE = """\
[bold #58a6ff]─── Welcome to CATGPT v{version} ───[/]

[#8b949e]Talking to:[/]  [#e6edf3]{url}[/]
[#8b949e]Model:[/]       [#58a6ff]{model}[/]

[bold #e6edf3]Commands[/]
  [#58a6ff]/help[/]              [#8b949e]│[/] Show all commands
  [#58a6ff]/new[/]               [#8b949e]│[/] Start fresh conversation
  [#58a6ff]/system <text>[/]     [#8b949e]│[/] Set a system prompt
  [#58a6ff]/model <name>[/]      [#8b949e]│[/] Switch model
  [#58a6ff]/history[/]           [#8b949e]│[/] View conversation turns
  [#58a6ff]/export [file][/]     [#8b949e]│[/] Export to markdown
  [#58a6ff]/status[/]            [#8b949e]│[/] API config & session info

[bold #e6edf3]Shortcuts[/]
  [bold #6e7681]Ctrl+N[/]  New chat    [bold #6e7681]Ctrl+E[/]  Export
  [bold #6e7681]Ctrl+L[/]  Clear       [bold #6e7681]Ctrl+C[/]  Quit
"""


# ================================================================
#  MESSAGE WIDGETS
# ================================================================


class UserMessage(Widget):
    """User message bubble with blue left bar."""

    DEFAULT_CLASSES = "user-msg"

    def __init__(self, text: str, turn: int) -> None:
        super().__init__()
        self._text = text
        self._turn = turn
        self._time = datetime.now().strftime("%H:%M:%S")

    def compose(self) -> ComposeResult:
        display = self._text if len(self._text) <= 600 else self._text[:597] + "…"
        words = len(self._text.split())
        yield Static(
            f"  You  ·  turn #{self._turn}  ·  {self._time}",
            classes="msg-header user-msg-header",
        )
        yield Static(display, classes="msg-body")
        yield Static(
            f"  {words} word{'s' if words != 1 else ''}",
            classes="msg-footer",
        )


class AssistantMessage(Widget):
    """Assistant response with green left bar and markdown rendering."""

    DEFAULT_CLASSES = "assistant-msg"

    def __init__(self, text: str, model: str, time_ms: int) -> None:
        super().__init__()
        self._text = text
        self._model = model
        self._time_ms = time_ms
        self._time = datetime.now().strftime("%H:%M:%S")

    def compose(self) -> ComposeResult:
        time_str = (
            f"{self._time_ms / 1000:.1f}s" if self._time_ms >= 1000 else f"{self._time_ms}ms"
        )
        words = len(self._text.split())
        yield Static(
            f"  {APP_NAME}  ·  {self._model}  ·  {self._time}",
            classes="msg-header assistant-msg-header",
        )
        if self._text.strip():
            yield Static(Markdown(self._text), classes="msg-body")
        else:
            yield Static("[dim]Empty response[/]", classes="msg-body")
        yield Static(
            f"  {words} word{'s' if words != 1 else ''}  ·  {time_str}",
            classes="msg-footer",
        )


class SystemPromptCard(Widget):
    """Displays the active system prompt."""

    DEFAULT_CLASSES = "system-prompt-card"

    def __init__(self, text: str) -> None:
        super().__init__()
        self._text = text

    def compose(self) -> ComposeResult:
        display = self._text if len(self._text) <= 300 else self._text[:297] + "…"
        yield Static("  ⚙  System Prompt Active", classes="msg-header system-prompt-header")
        yield Static(f"  {display}", classes="msg-body")


class ThinkingIndicator(Widget):
    """Animated spinner while waiting for the API response."""

    DEFAULT_CLASSES = "thinking-widget"

    frame: reactive[int] = reactive(0)

    def compose(self) -> ComposeResult:
        yield Static(
            f"  {THINKING_FRAMES[0]}  {APP_NAME} is thinking …",
            classes="thinking-label",
        )

    def on_mount(self) -> None:
        self._timer = self.set_interval(0.12, self._tick)

    def _tick(self) -> None:
        self.frame = (self.frame + 1) % len(THINKING_FRAMES)

    def watch_frame(self, frame: int) -> None:
        try:
            self.query_one(".thinking-label", Static).update(
                f"  {THINKING_FRAMES[frame]}  {APP_NAME} is thinking …"
            )
        except Exception:
            pass

    def stop_animation(self) -> None:
        try:
            self._timer.stop()
        except Exception:
            pass


# ================================================================
#  SPLASH SCREEN
# ================================================================


class SplashScreen(Screen):
    """Animated splash screen. Auto-transitions after 2.5 s or on keypress."""

    def compose(self) -> ComposeResult:
        with Center():
            with Vertical(id="splash-container"):
                yield Static(CAT_ART, id="splash-cat")
                yield Static(LOGO_TEXT, id="splash-logo")
                yield Static(f"───  {APP_TAGLINE}  ───", id="splash-tagline")
                yield Static(
                    f"v{VERSION}  ·  OpenAI-compatible  ·  any provider",
                    id="splash-version",
                )
                yield Static("press any key to skip", id="splash-hint")

    def on_mount(self) -> None:
        self.set_timer(2.5, self._go)

    def on_key(self, _: object) -> None:
        self._go()

    def _go(self) -> None:
        if self.app.screen is self:
            self.app.switch_screen("chat")


# ================================================================
#  CHAT SCREEN
# ================================================================


class ChatScreen(Screen):
    """Main chat interface — messages, input, keybindings."""

    BINDINGS = [
        Binding("ctrl+n", "new_chat",   "New Chat", key_display="^N"),
        Binding("ctrl+e", "export",     "Export",   key_display="^E"),
        Binding("ctrl+l", "clear_chat", "Clear",    key_display="^L"),
        Binding("ctrl+c", "quit_app",   "Quit",     key_display="^C", priority=True),
    ]

    # ── State ────────────────────────────────────────────────────

    def __init__(self) -> None:
        super().__init__()
        self.api_url = (
            os.getenv("CATGPT_API_URL")
            or os.getenv("OPENAI_API_BASE")
            or "http://localhost:8000/v1"
        )
        self.api_key = (
            os.getenv("CATGPT_API_KEY")
            or os.getenv("OPENAI_API_KEY")
            or "dummy123"
        )
        self.model = os.getenv("CATGPT_MODEL") or "catgpt-browser"

        self.messages: list[dict] = []       # full OpenAI-format conversation history
        self.system_prompt: str | None = None
        self.turn_count = 0
        self.last_time_ms = 0
        self.session_start = datetime.now()
        self._is_busy = False
        self._openai = None                  # openai.AsyncOpenAI — set in on_mount

    def on_mount(self) -> None:
        import openai

        self._openai = openai.AsyncOpenAI(
            base_url=self.api_url,
            api_key=self.api_key,
        )
        self.app.title = APP_NAME
        self.app.sub_title = f"{self.model}  ·  {self.api_url}"
        self.query_one("#chat-container", Vertical).border_title = f" 🐱  {APP_NAME} "
        self._show_welcome()
        self._check_connection()
        self.query_one("#chat-input", Input).focus()

    # ── Layout ───────────────────────────────────────────────────

    def compose(self) -> ComposeResult:
        yield Header(show_clock=True)
        yield Static(self._status_text(), id="status-bar")
        with Vertical(id="chat-container"):
            with ScrollableContainer(id="chat-log"):
                pass
        yield Input(
            placeholder="Message CATGPT …  (/help for commands)",
            id="chat-input",
        )
        yield Footer()

    @property
    def chat_log(self) -> ScrollableContainer:
        return self.query_one("#chat-log", ScrollableContainer)

    # ── Welcome & connection check ───────────────────────────────

    def _show_welcome(self) -> None:
        text = WELCOME_TEMPLATE.format(
            version=VERSION,
            url=self.api_url,
            model=self.model,
        )
        self.chat_log.mount(Static(text, classes="welcome-card"))

    @work(exclusive=False, name="check_conn")
    async def _check_connection(self) -> None:
        try:
            result = await self._openai.models.list()
            names = [m.id for m in result.data]
            if names and self.model not in names:
                self.model = names[0]
                self.app.sub_title = f"{self.model}  ·  {self.api_url}"
            model_list = "  ".join(f"[#58a6ff]{n}[/]" for n in names[:5])
            self._mount_system(
                f"[#3fb950]✓[/]  Connected to [#58a6ff]{self.api_url}[/]\n"
                f"  Available: {model_list}",
                "system-success",
            )
        except Exception as exc:
            self._mount_system(
                f"[#f85149]✗[/]  Cannot reach [#58a6ff]{self.api_url}[/]\n"
                f"  Is the server running?  [#6e7681]{exc}[/]",
                "system-error",
            )
        self._refresh_status()

    # ── Input handling ───────────────────────────────────────────

    def on_input_submitted(self, event: Input.Submitted) -> None:
        text = event.value.strip()
        event.input.value = ""
        if not text:
            return
        if text.startswith("/"):
            parts = text.split(maxsplit=1)
            self._dispatch_command(parts[0].lower(), parts[1] if len(parts) > 1 else "")
        else:
            self._send(text)

    # ── Send message ─────────────────────────────────────────────

    def _send(self, text: str) -> None:
        if self._is_busy:
            self._mount_system("⚠  Please wait for the current response …", "system-warn")
            return

        self.turn_count += 1
        self.chat_log.mount(UserMessage(text, self.turn_count))
        thinking = ThinkingIndicator()
        self.chat_log.mount(thinking)
        self.chat_log.scroll_end(animate=False)
        self._is_busy = True
        self._refresh_status()
        self._do_send(text, thinking)

    @work(exclusive=True, name="send")
    async def _do_send(self, text: str, thinking: ThinkingIndicator) -> None:
        self.messages.append({"role": "user", "content": text})

        payload: list[dict] = []
        if self.system_prompt:
            payload.append({"role": "system", "content": self.system_prompt})
        payload.extend(self.messages)

        try:
            t0 = datetime.now()
            resp = await self._openai.chat.completions.create(
                model=self.model,
                messages=payload,
            )
            elapsed_ms = int((datetime.now() - t0).total_seconds() * 1000)
            content = resp.choices[0].message.content or ""
            self.messages.append({"role": "assistant", "content": content})
            self.last_time_ms = elapsed_ms
            thinking.stop_animation()
            thinking.remove()
            self.chat_log.mount(AssistantMessage(content, self.model, elapsed_ms))

        except Exception as exc:
            log.error(f"API call failed: {exc}", exc_info=True)
            self.messages.pop()
            self.turn_count = max(0, self.turn_count - 1)
            thinking.stop_animation()
            thinking.remove()
            self._mount_system(f"[#f85149]✗[/]  {exc}", "system-error")

        self._is_busy = False
        self._refresh_status()
        self.chat_log.scroll_end(animate=False)

    # ── Command dispatch ─────────────────────────────────────────

    def _dispatch_command(self, cmd: str, args: str) -> None:
        dispatch: dict[str, object] = {
            "/exit":    lambda: self.action_quit_app(),
            "/quit":    lambda: self.action_quit_app(),
            "/q":       lambda: self.action_quit_app(),
            "/help":    lambda: self._show_help(),
            "/clear":   lambda: self.action_clear_chat(),
            "/new":     lambda: self.action_new_chat(),
            "/history": lambda: self._show_history(),
            "/export":  lambda: self._export_command(args),
            "/status":  lambda: self._show_status(),
            "/system":  lambda: self._set_system(args),
            "/model":   lambda: self._set_model(args),
        }
        handler = dispatch.get(cmd)
        if handler:
            handler()  # type: ignore[operator]
        else:
            self._mount_system(
                f"[#f85149]✗[/]  Unknown command: [bold]{cmd}[/] — type /help",
                "system-error",
            )

    # ── /help ────────────────────────────────────────────────────

    def _show_help(self) -> None:
        lines = [
            "[bold #58a6ff]─── CATGPT Commands ───[/]\n",
            "  [#58a6ff]/new[/]               Start fresh (clears history & system prompt)",
            "  [#58a6ff]/clear[/]             Clear the display  (history preserved)",
            "  [#58a6ff]/system <text>[/]     Set a system prompt for this session",
            "  [#58a6ff]/model <name>[/]      Switch model  (e.g. /model gpt-4o)",
            "  [#58a6ff]/history[/]           Show all conversation turns",
            "  [#58a6ff]/export [file][/]     Export conversation to markdown",
            "  [#58a6ff]/status[/]            API config & session info",
            "  [#58a6ff]/help[/]              Show this help",
            "  [#58a6ff]/exit[/]              Quit",
            "",
            "[bold #58a6ff]─── Shortcuts ───[/]\n",
            "  [bold #6e7681]Ctrl+N[/]  New chat     [bold #6e7681]Ctrl+E[/]  Export markdown",
            "  [bold #6e7681]Ctrl+L[/]  Clear        [bold #6e7681]Ctrl+C[/]  Quit",
            "",
            "[dim italic]  Tip: /system 'You are a senior Python engineer' sets a persistent persona",
            "  Tip: /clear keeps history — the model still remembers previous turns",
            "  Tip: Set CATGPT_API_URL / CATGPT_API_KEY / CATGPT_MODEL as env vars[/]",
        ]
        self._mount_system("\n".join(lines), "system-info-block")

    # ── /status ──────────────────────────────────────────────────

    def _show_status(self) -> None:
        elapsed = datetime.now() - self.session_start
        m, s = divmod(int(elapsed.total_seconds()), 60)
        masked_key = (
            self.api_key[:4] + "•" * max(0, len(self.api_key) - 4)
            if len(self.api_key) > 4
            else "•" * len(self.api_key)
        )
        sys_display = (
            f"[#d29922]{self.system_prompt[:60]}{'…' if len(self.system_prompt) > 60 else ''}[/]"
            if self.system_prompt
            else "[#6e7681]not set[/]"
        )
        lines = [
            "[bold #58a6ff]─── CATGPT Status ───[/]\n",
            f"  API URL         [#58a6ff]{self.api_url}[/]",
            f"  Model           [#58a6ff]{self.model}[/]",
            f"  Auth token      [#6e7681]{masked_key}[/]",
            f"  Turns           {self.turn_count}",
            f"  History msgs    {len(self.messages)}",
            f"  System prompt   {sys_display}",
            (
                f"  Last response   [#3fb950]{self.last_time_ms}ms[/]"
                if self.last_time_ms
                else "  Last response   [#6e7681]—[/]"
            ),
            f"  Session uptime  {m}m {s}s",
        ]
        self._mount_system("\n".join(lines), "system-info-block")

    # ── /history ─────────────────────────────────────────────────

    def _show_history(self) -> None:
        if not self.messages:
            self._mount_system("[#8b949e]No conversation history yet.[/]", "system-msg")
            return
        lines = [
            f"[bold #58a6ff]─── Conversation History ({len(self.messages)} messages) ───[/]\n"
        ]
        for i, msg in enumerate(self.messages, 1):
            role = msg["role"]
            content = str(msg.get("content") or "")
            preview = content[:90].replace("\n", " ")
            if len(content) > 90:
                preview += "…"
            color, icon = {
                "user":      ("#58a6ff", "▶"),
                "assistant": ("#3fb950", "◀"),
            }.get(role, ("#d29922", "⚙"))
            lines.append(
                f"  [{color}]{i:>2}. {icon} {role:<10}[/]  [#8b949e]{preview}[/]"
            )
        lines.append("\n[#6e7681]  Use /new to clear history, /export to save it[/]")
        self._mount_system("\n".join(lines), "system-info-block")

    # ── /system ──────────────────────────────────────────────────

    def _set_system(self, text: str) -> None:
        text = text.strip()
        if not text:
            self.system_prompt = None
            self._mount_system("[#8b949e]System prompt cleared.[/]", "system-msg")
            self._refresh_status()
            return
        self.system_prompt = text
        self.chat_log.mount(SystemPromptCard(text))
        self.chat_log.scroll_end(animate=False)
        self._mount_system(
            "[#3fb950]✓[/]  System prompt set — applies to all future messages.",
            "system-success",
        )
        self._refresh_status()

    # ── /model ───────────────────────────────────────────────────

    def _set_model(self, name: str) -> None:
        name = name.strip()
        if not name:
            self._mount_system(
                f"[#8b949e]Current model: [#58a6ff]{self.model}[/]  "
                "[#6e7681]Use /model <name> to switch[/]",
                "system-msg",
            )
            return
        old = self.model
        self.model = name
        self.app.sub_title = f"{self.model}  ·  {self.api_url}"
        self._refresh_status()
        self._mount_system(
            f"[#3fb950]✓[/]  Switched: [#6e7681]{old}[/] → [#58a6ff]{name}[/]",
            "system-success",
        )

    # ── /export ──────────────────────────────────────────────────

    def _export_command(self, filename: str) -> None:
        if not self.messages:
            self._mount_system("[#8b949e]Nothing to export yet.[/]", "system-msg")
            return
        self._do_export(filename.strip())

    @work(exclusive=False, name="export")
    async def _do_export(self, filename: str) -> None:
        ts = datetime.now().strftime("%Y%m%d_%H%M%S")
        name = filename or f"catgpt-export-{ts}.md"
        if not name.endswith(".md"):
            name += ".md"
        path = Path(name) if ("/" in name or "\\" in name) else Path.cwd() / name

        lines = [
            f"# CATGPT Export — {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}\n",
            f"**Model:** `{self.model}`  ",
            f"**API:** `{self.api_url}`  ",
            f"**Turns:** {self.turn_count}  ",
        ]
        if self.system_prompt:
            lines += [f"\n**System Prompt:**\n> {self.system_prompt}\n"]
        lines.append("\n---\n")
        for msg in self.messages:
            role = msg["role"].title()
            content = msg.get("content") or ""
            lines.append(f"\n### {role}\n\n{content}\n")

        try:
            path.write_text("\n".join(lines), encoding="utf-8")
            self._mount_system(
                f"[#3fb950]✓[/]  Exported {len(self.messages)} messages → [#58a6ff]{path}[/]",
                "system-success",
            )
        except Exception as exc:
            self._mount_system(f"[#f85149]✗[/]  Export failed: {exc}", "system-error")

    # ── Actions (keybindings) ────────────────────────────────────

    def action_new_chat(self) -> None:
        self.messages.clear()
        self.turn_count = 0
        self.last_time_ms = 0
        self.system_prompt = None
        self.chat_log.remove_children()
        self.chat_log.mount(
            Static(
                "[#3fb950]✓[/]  New conversation — history cleared.",
                classes="system-success",
            )
        )
        self._refresh_status()

    def action_clear_chat(self) -> None:
        self.chat_log.remove_children()
        self.chat_log.mount(
            Static(
                "[#8b949e]Display cleared.  "
                "[dim]History still active — use /new to fully reset.[/]",
                classes="system-msg",
            )
        )

    def action_export(self) -> None:
        if not self.messages:
            self._mount_system("[#8b949e]Nothing to export yet.[/]", "system-msg")
            return
        self._do_export("")

    def action_quit_app(self) -> None:
        self.app.exit()

    # ── Helpers ──────────────────────────────────────────────────

    def _mount_system(self, text: str, css_class: str = "system-msg") -> None:
        self.chat_log.mount(Static(text, classes=css_class))
        self.chat_log.scroll_end(animate=False)

    def _status_text(self) -> str:
        conn = "[#d29922]● thinking[/]" if self._is_busy else "[#3fb950]● ready[/]"
        model = f"[#58a6ff]{self.model}[/]"
        turns = f"turn {self.turn_count}"
        msgs = f"[#6e7681]{len(self.messages)} msgs[/]"
        time_str = (
            f"[#3fb950]{self.last_time_ms}ms[/]" if self.last_time_ms else "[#6e7681]—[/]"
        )
        sys_ind = "  [#d29922]⚙ system[/]" if self.system_prompt else ""
        return f"  {conn}  │  {model}  │  {turns}  │  {msgs}  │  {time_str}{sys_ind}"

    def _refresh_status(self) -> None:
        try:
            self.query_one("#status-bar", Static).update(self._status_text())
        except Exception:
            pass


# ================================================================
#  APP
# ================================================================


class CatGPTApp(App):
    """CATGPT — OpenAI-Compatible Terminal Chat."""

    TITLE = APP_NAME
    SUB_TITLE = APP_TAGLINE
    CSS_PATH = "catgpt.tcss"

    SCREENS = {"chat": ChatScreen}

    def on_mount(self) -> None:
        self.push_screen(SplashScreen())


# ================================================================
#  ENTRY POINTS
# ================================================================


@cli.command()
def chat(
    api_url: str = typer.Option(None, "--api-url", "-u", help="API base URL"),
    api_key: str = typer.Option(None, "--api-key", "-k", help="Bearer token"),
    model: str = typer.Option(None, "--model", "-m", help="Model name"),
) -> None:
    """Start an interactive CATGPT terminal session."""
    if api_url:
        os.environ["CATGPT_API_URL"] = api_url
    if api_key:
        os.environ["CATGPT_API_KEY"] = api_key
    if model:
        os.environ["CATGPT_MODEL"] = model
    CatGPTApp().run()


def main() -> None:
    """Entry point."""
    cli()


if __name__ == "__main__":
    main()
</file>

<file path="src/cli/catgpt.tcss">
/* ════════════════════════════════════════════════════════════════
   CATGPT — GitHub Dark Theme
   Palette:
     bg-canvas    #0d1117    main background
     bg-overlay   #161b22    panels, cards, header
     bg-subtle    #1c2128    hover / secondary surfaces
     border       #30363d    default borders
     border-muted #21262d
     fg-default   #e6edf3    primary text
     fg-muted     #8b949e    secondary text
     fg-subtle    #6e7681    tertiary text
     blue         #58a6ff    user accent / links
     green        #3fb950    success / assistant accent
     red          #f85149    errors
     orange       #d29922    warnings / system prompt
     purple       #bc8cff    special
   ════════════════════════════════════════════════════════════════ */


/* ── Global ──────────────────────────────────────────────────── */

Screen {
    background: #0d1117;
}

* {
    scrollbar-background: #161b22;
    scrollbar-color: #30363d;
    scrollbar-color-hover: #58a6ff;
    scrollbar-color-active: #58a6ff;
    scrollbar-size-vertical: 1;
}


/* ── Splash Screen ───────────────────────────────────────────── */

SplashScreen {
    align: center middle;
    background: #0d1117;
}

#splash-container {
    width: auto;
    height: auto;
    align: center middle;
    content-align: center middle;
    padding: 1 4;
}

#splash-cat {
    text-align: center;
    color: #58a6ff;
    width: auto;
    height: auto;
    content-align: center middle;
}

#splash-logo {
    text-align: center;
    color: #e6edf3;
    text-style: bold;
    width: auto;
    height: auto;
    margin-top: 1;
}

#splash-tagline {
    text-align: center;
    color: #8b949e;
    margin-top: 1;
    width: 100%;
}

#splash-version {
    text-align: center;
    color: #6e7681;
    text-style: italic;
    margin-top: 0;
    width: 100%;
}

#splash-hint {
    text-align: center;
    color: #30363d;
    margin-top: 2;
    width: 100%;
}


/* ── Chat Screen Layout ──────────────────────────────────────── */

ChatScreen {
    layout: vertical;
    background: #0d1117;
}

#status-bar {
    dock: top;
    width: 100%;
    height: 1;
    background: #161b22;
    color: #8b949e;
    padding: 0 2;
}

#chat-container {
    width: 100%;
    height: 1fr;
    border: round #30363d;
    border-title-color: #58a6ff;
    border-title-style: bold;
    background: #0d1117;
    margin: 0 1;
}

#chat-log {
    width: 100%;
    height: 1fr;
    padding: 1 2;
    background: #0d1117;
}


/* ── Input ───────────────────────────────────────────────────── */

#chat-input {
    width: 100%;
    height: 3;
    border: round #30363d;
    background: #0d1117;
    color: #e6edf3;
    padding: 0 1;
    margin: 0 1 0 1;
}

#chat-input:focus {
    border: round #58a6ff;
}

#chat-input > .input--cursor {
    background: #e6edf3;
    color: #0d1117;
}

#chat-input > .input--placeholder {
    color: #484f58;
    text-style: italic;
}

#chat-input > .input--selection {
    background: #264f78;
}


/* ── Shared message structure ────────────────────────────────── */

.msg-header {
    text-style: bold;
    width: 100%;
    height: 1;
}

.msg-body {
    color: #e6edf3;
    width: 100%;
    height: auto;
    margin-top: 0;
    padding-top: 0;
}

.msg-footer {
    color: #6e7681;
    text-style: italic;
    width: 100%;
    height: 1;
    margin-top: 0;
}


/* ── User Messages ───────────────────────────────────────────── */

.user-msg {
    margin: 1 0 0 0;
    padding: 1 2;
    background: #161b22;
    border-left: thick #58a6ff;
    width: 100%;
    height: auto;
}

.user-msg-header {
    color: #58a6ff;
}


/* ── Assistant Messages ──────────────────────────────────────── */

.assistant-msg {
    margin: 1 0 0 0;
    padding: 1 2;
    background: #161b22;
    border-left: thick #3fb950;
    width: 100%;
    height: auto;
}

.assistant-msg-header {
    color: #3fb950;
}


/* ── System Prompt Card ──────────────────────────────────────── */

.system-prompt-card {
    margin: 1 0 0 0;
    padding: 1 2;
    background: #1c1a10;
    border-left: thick #d29922;
    width: 100%;
    height: auto;
}

.system-prompt-header {
    color: #d29922;
}

.system-prompt-card .msg-body {
    color: #e6c97a;
    text-style: italic;
}


/* ── Thinking Indicator ──────────────────────────────────────── */

.thinking-widget {
    margin: 1 0;
    padding: 0 2;
    width: 100%;
    height: auto;
}

.thinking-label {
    color: #58a6ff;
    text-style: bold italic;
    width: 100%;
    height: 1;
}


/* ── System Messages ─────────────────────────────────────────── */

.system-msg {
    text-align: center;
    color: #8b949e;
    text-style: italic;
    margin: 1 0;
    padding: 0 2;
    width: 100%;
    height: auto;
}

.system-success {
    color: #3fb950;
    margin: 1 0;
    padding: 0 2;
    width: 100%;
    height: auto;
}

.system-error {
    color: #f85149;
    text-style: bold;
    margin: 1 0;
    padding: 0 2;
    width: 100%;
    height: auto;
}

.system-warn {
    text-align: center;
    color: #d29922;
    text-style: bold;
    margin: 1 0;
    padding: 0 2;
    width: 100%;
    height: auto;
}

.system-info-block {
    margin: 1 0;
    padding: 1 2;
    background: #161b22;
    border: round #30363d;
    width: 100%;
    height: auto;
    color: #e6edf3;
}


/* ── Welcome Card ────────────────────────────────────────────── */

.welcome-card {
    margin: 2 4;
    padding: 2 4;
    background: #161b22;
    border: round #30363d;
    width: 100%;
    height: auto;
    color: #e6edf3;
}


/* ── Header / Footer ─────────────────────────────────────────── */

Header {
    background: #161b22;
    color: #e6edf3;
    dock: top;
    height: 1;
}

Header .header--highlight {
    color: #58a6ff;
    text-style: bold;
}

Footer {
    background: #161b22;
    color: #6e7681;
}

FooterKey {
    background: #161b22;
    color: #8b949e;
}

FooterKey:hover {
    background: #1c2128;
    color: #58a6ff;
}

FooterKey > .footer-key--key {
    background: #21262d;
    color: #58a6ff;
}
</file>

<file path="src/config.py">
"""
Centralized configuration — loads from .env with sensible defaults.
"""

from __future__ import annotations

import os
from pathlib import Path
from dotenv import load_dotenv

_CODE_ROOT = Path(__file__).resolve().parent.parent
_CWD = Path.cwd()

# Prefer the invocation directory as project root when running from
# a checkout (e.g. `nix run .#proxy` from repo root). Fall back to the
# code location (used for packaged/store execution).
if (_CWD / "src").exists() and (_CWD / "scripts").exists():
    _PROJECT_ROOT = _CWD
else:
    _PROJECT_ROOT = _CODE_ROOT

# Load .env from current working directory first, then from the
# resolved project root.
load_dotenv(_CWD / ".env")
load_dotenv(_PROJECT_ROOT / ".env")


class Config:
    """All project settings in one place."""

    # Paths
    PROJECT_ROOT: Path = _PROJECT_ROOT
    BROWSER_DATA_DIR: Path = _PROJECT_ROOT / os.getenv("BROWSER_DATA_DIR", "browser_data")
    LOG_DIR: Path = _PROJECT_ROOT / os.getenv("LOG_DIR", "logs")
    IMAGES_DIR: Path = _PROJECT_ROOT / os.getenv("IMAGES_DIR", "downloads/images")

    # Browser
    HEADLESS: bool = os.getenv("HEADLESS", "false").lower() == "true"
    SLOW_MO: int = int(os.getenv("SLOW_MO", "50"))
    CHATGPT_URL: str = os.getenv("CHATGPT_URL", "https://chatgpt.com")
    CLAUDE_URL: str = os.getenv("CLAUDE_URL", "https://claude.ai")

    # Provider selection: "chatgpt" or "claude"
    PROVIDER: str = os.getenv("PROVIDER", "chatgpt").lower()

    @classmethod
    def provider_url(cls) -> str:
        """Return the target URL for the active provider."""
        if cls.PROVIDER == "claude":
            return cls.CLAUDE_URL
        return cls.CHATGPT_URL

    # Timeouts (ms)
    RESPONSE_TIMEOUT: int = int(os.getenv("RESPONSE_TIMEOUT", "120000"))
    SELECTOR_TIMEOUT: int = int(os.getenv("SELECTOR_TIMEOUT", "10000"))

    # Human simulation (ms)
    TYPING_SPEED_MIN: int = int(os.getenv("TYPING_SPEED_MIN", "50"))
    TYPING_SPEED_MAX: int = int(os.getenv("TYPING_SPEED_MAX", "150"))
    THINKING_PAUSE_MIN: int = int(os.getenv("THINKING_PAUSE_MIN", "500"))
    THINKING_PAUSE_MAX: int = int(os.getenv("THINKING_PAUSE_MAX", "1500"))
    # Completion poll interval — how often to check if response is ready (ms)
    POLL_INTERVAL_MS: int = int(os.getenv("POLL_INTERVAL_MS", "500"))

    # Logging
    LOG_LEVEL: str = os.getenv("LOG_LEVEL", "DEBUG")
    VERBOSE: bool = os.getenv("VERBOSE", "true").lower() == "true"

    # API (Phase 3)
    API_HOST: str = os.getenv("API_HOST", "0.0.0.0")
    API_PORT: int = int(os.getenv("API_PORT", "8000"))
    RATE_LIMIT_SECONDS: int = int(os.getenv("RATE_LIMIT_SECONDS", "5"))
    API_TOKEN: str = os.getenv("API_TOKEN", "")  # Bearer token for API auth (empty = no auth)

    # VNC
    VNC_PASSWORD: str = os.getenv("VNC_PASSWORD", "catgpt")

    # Viewport base (will be jittered ±20px)
    VIEWPORT_WIDTH: int = 1280
    VIEWPORT_HEIGHT: int = 720

    @classmethod
    def ensure_dirs(cls) -> None:
        """Create required directories if they don't exist."""
        cls.BROWSER_DATA_DIR.mkdir(parents=True, exist_ok=True)
        cls.LOG_DIR.mkdir(parents=True, exist_ok=True)
        cls.IMAGES_DIR.mkdir(parents=True, exist_ok=True)
</file>

<file path="src/dom_observer.py">
"""
DOM Observer — watches DOM changes in real-time for debugging / Phase 1 observation.

Injects a MutationObserver to track what changes when ChatGPT responds.
"""

from __future__ import annotations

from patchright.async_api import Page

from src.log import setup_logging

log = setup_logging("dom_observer")


class DOMObserver:
    """Observe and log DOM mutations on the ChatGPT page."""

    def __init__(self, page: Page) -> None:
        self._page = page
        self._active = False

    async def start(self, target_selector: str = "main") -> None:
        """
        Inject a MutationObserver that logs DOM changes to the console.
        We capture those via page console listener.
        """
        if self._active:
            return

        self._page.on("console", self._on_console)

        await self._page.evaluate(f"""
        (() => {{
            if (window.__domObserver) window.__domObserver.disconnect();
            const target = document.querySelector('{target_selector}');
            if (!target) {{ console.log('[DOM_OBS] Target not found: {target_selector}'); return; }}

            const observer = new MutationObserver((mutations) => {{
                for (const m of mutations) {{
                    if (m.type === 'childList' && m.addedNodes.length > 0) {{
                        for (const node of m.addedNodes) {{
                            if (node.nodeType === 1) {{
                                const tag = node.tagName || 'unknown';
                                const cls = node.className || '';
                                const text = (node.textContent || '').slice(0, 100);
                                console.log('[DOM_OBS] ADDED ' + tag + '.' + cls + ' | ' + text);
                            }}
                        }}
                    }}
                }}
            }});

            observer.observe(target, {{ childList: true, subtree: true }});
            window.__domObserver = observer;
            console.log('[DOM_OBS] Observer started on: {target_selector}');
        }})();
        """)
        self._active = True
        log.info(f"DOM observer started on '{target_selector}'")

    async def stop(self) -> None:
        """Disconnect the MutationObserver."""
        if not self._active:
            return
        try:
            await self._page.evaluate("if(window.__domObserver) window.__domObserver.disconnect();")
        except Exception:
            pass
        self._active = False
        log.info("DOM observer stopped")

    def _on_console(self, msg) -> None:
        """Handle browser console messages from our observer."""
        text = msg.text
        if text.startswith("[DOM_OBS]"):
            log.debug(text)
</file>

<file path="src/log.py">
"""
Logging setup — file + optional console handlers.
"""

from __future__ import annotations

import logging
import sys
from datetime import datetime
from src.config import Config

# Global flag: when True, suppress console log handlers (for TUI mode)
_suppress_console = False


def suppress_console_logs() -> None:
    """Disable all console log handlers (call before any setup_logging)."""
    global _suppress_console
    _suppress_console = True
    # Also silence already-created loggers' console handlers
    for name in list(logging.Logger.manager.loggerDict):
        logger = logging.getLogger(name)
        for handler in logger.handlers[:]:
            if isinstance(handler, logging.StreamHandler) and handler.stream in (sys.stdout, sys.stderr):
                logger.removeHandler(handler)


def setup_logging(name: str = "chatgpt_scraper", log_file: str | None = None) -> logging.Logger:
    """
    Configure and return a logger that writes to file (and optionally console).

    Args:
        name: Logger name.
        log_file: Optional filename override. Defaults to '{name}_{date}.log'.
    """
    Config.ensure_dirs()

    logger = logging.getLogger(name)
    logger.setLevel(getattr(logging, Config.LOG_LEVEL.upper(), logging.DEBUG))

    # Prevent duplicate handlers on repeated calls
    if logger.handlers:
        return logger

    formatter = logging.Formatter(
        fmt="%(asctime)s | %(levelname)-8s | %(name)s | %(message)s",
        datefmt="%Y-%m-%d %H:%M:%S",
    )

    # ── File handler ────────────────────────────────────────────
    if log_file is None:
        date_str = datetime.now().strftime("%Y%m%d")
        log_file = f"{name}_{date_str}.log"

    fh = logging.FileHandler(Config.LOG_DIR / log_file, encoding="utf-8")
    fh.setLevel(logging.DEBUG)  # Always capture everything in file
    fh.setFormatter(formatter)
    logger.addHandler(fh)

    # ── Console handler (disabled in TUI mode) ──────────────────
    if Config.VERBOSE and not _suppress_console:
        ch = logging.StreamHandler(sys.stdout)
        ch.setLevel(logging.INFO)
        ch.setFormatter(formatter)
        logger.addHandler(ch)

    return logger
</file>

<file path="src/network_recorder.py">
"""
Network Recorder — captures and logs network requests for Phase 1 observation.

Records API calls ChatGPT makes so we understand the request patterns.
"""

from __future__ import annotations

from patchright.async_api import Page, Request, Response

from src.log import setup_logging

log = setup_logging("network")


class NetworkRecorder:
    """Record network activity on a Playwright page."""

    def __init__(self, page: Page) -> None:
        self._page = page
        self._requests: list[dict] = []
        self._active = False

    def start(self) -> None:
        """Start recording network requests and responses."""
        if self._active:
            return
        self._page.on("request", self._on_request)
        self._page.on("response", self._on_response)
        self._active = True
        log.info("Network recorder started")

    def stop(self) -> None:
        """Stop recording (note: Playwright doesn't support removing listeners easily)."""
        self._active = False
        log.info(f"Network recorder stopped — captured {len(self._requests)} requests")

    def _on_request(self, request: Request) -> None:
        if not self._active:
            return
        url = request.url
        # Only log interesting API calls, skip static assets
        if any(k in url for k in ["backend-api", "conversation", "auth", "sentinel"]):
            entry = {
                "method": request.method,
                "url": url,
                "type": request.resource_type,
            }
            self._requests.append(entry)
            log.debug(f"REQ  {request.method} {url[:120]}")

    def _on_response(self, response: Response) -> None:
        if not self._active:
            return
        url = response.url
        if any(k in url for k in ["backend-api", "conversation", "auth", "sentinel"]):
            log.debug(f"RESP {response.status} {url[:120]}")

    def get_captured(self) -> list[dict]:
        """Return all captured request entries."""
        return list(self._requests)

    def clear(self) -> None:
        """Clear captured requests."""
        self._requests.clear()
</file>

<file path="src/selectors.py">
"""
Centralized DOM selectors for ChatGPT.

All selectors live here so when ChatGPT updates their UI, we only
change this one file. Each entry is a list of fallback selectors —
try them in order until one matches.
"""

from __future__ import annotations


class Selectors:
    """CSS / Playwright selectors for chatgpt.com UI elements."""

    # ── Chat input ──────────────────────────────────────────────
    CHAT_INPUT = [
        "#prompt-textarea",
        "div[contenteditable='true'][id='prompt-textarea']",
        "div[contenteditable='true']",
    ]

    # ── Send button ─────────────────────────────────────────────
    SEND_BUTTON = [
        "button[data-testid='send-button']",
        "#composer-submit-button",
        "button[aria-label='Send prompt']",
        "#prompt-textarea ~ button",
    ]

    # ── Assistant response messages ─────────────────────────────
    ASSISTANT_MESSAGE = [
        "div[data-message-author-role='assistant']",
        "[data-message-author-role='assistant']",
        "section[data-turn='assistant']",
        "section[data-testid^='conversation-turn-']",
    ]

    # ── Streaming / stop button (visible while generating) ─────
    STOP_BUTTON = [
        "button[data-testid='stop-button']",
        "button[aria-label='Stop answering']",
        "button[aria-label='Stop generating']",
    ]

    # ── New chat ────────────────────────────────────────────────
    NEW_CHAT_BUTTON = [
        "a[data-testid='create-new-chat-button']",
        "a[href='/']",
        "nav a[href='/']",
    ]

    # ── Sidebar conversation links ──────────────────────────────
    SIDEBAR_THREAD_LINKS = [
        "nav a[href^='/c/']",
        "a[href^='/c/']",
    ]

    # ── Login page detection (if any of these appear, user is logged out) ──
    LOGIN_INDICATORS = [
        "button[data-testid='login-button']",
        "button:has-text('Log in')",
        "[data-testid='login-button']",
    ]

    # ── Markdown content inside assistant message ───────────────
    ASSISTANT_MARKDOWN = [
        "div[data-message-author-role='assistant'] .markdown",
        "div[data-message-author-role='assistant'] .prose",
        "section[data-turn='assistant'] .markdown",
        "section[data-turn='assistant'] .prose",
    ]

    # ── Regenerate / continue buttons (appear after response completes) ──
    POST_RESPONSE_BUTTONS = [
        "button:has-text('Regenerate')",
        "button:has-text('Continue generating')",
    ]

    # ── Copy button (appears on each completed assistant message) ──────
    # This is the most reliable completion signal — it only appears
    # after the full response has been generated.
    COPY_BUTTON = [
        "button[data-testid='copy-turn-action-button']",
        "button[aria-label='Copy message']",
        "button[aria-label='Copy']",
    ]

    # ── Generated images inside assistant responses ───────────────────
    # ChatGPT DALL-E image responses do NOT have data-message-author-role.
    # Instead, the image lives inside an article turn with class "agent-turn".
    # Images have alt="Generated image" and src from chatgpt.com/backend-api.
    # Image wrapper DIVs have id="image-{uuid}" and class group/imagegen-image.
    ASSISTANT_IMAGE = [
        "img[alt='Generated image']",
        "div[id^='image-'] img",
        "section[data-turn='assistant'] img[alt='Generated image']",
    ]

    # Image container identifiers (used for detection, not clicking)
    IMAGE_CONTAINER = [
        "div[id^='image-']",
        "div[class*='imagegen-image']",
    ]

    # Download button for generated images
    IMAGE_DOWNLOAD_BUTTON = [
        "a[aria-label='Download']",
        "a[download]",
    ]

    # ── File / attachment upload input ────────────────────────────
    FILE_UPLOAD_INPUT = [
        "input#upload-photos",
        "input[type='file']",
        "input[data-testid='file-upload']",
        "input[accept*='image']",
    ]

    # Attach / upload button (opens file picker)
    ATTACH_BUTTON = [
        "button[data-testid='composer-plus-btn']",
        "button[aria-label='Add files and more']",
        "button[aria-label='Attach files']",
    ]
</file>

<file path="tests/__init__.py">

</file>

<file path=".gitignore">
# ── Browser Profiles (contain cookies, login sessions) ──
browser_data/
browser_data_claude/
chrome-profile/

# ── Logs ──
logs/
docker-logs/

# ── Downloads (test assets, generated images) ──
downloads/

# ── Python ──
__pycache__/
*.py[cod]
*$py.class
*.so
.Python
build/
dist/
*.egg-info/
*.egg

# ── Virtual Environment ──
.venv/
venv/
env/

# ── Environment Variables (contains tokens) ──
.env

# ── Sensitive Data ──
cookies.json
*.pem
*.key

# ── IDE ──
.vscode/
.idea/
*.swp
*.swo
*~

# ── OS ──
.DS_Store
Thumbs.db

# ── Nix ──
result
</file>

<file path="README.md">
<p align="center">
  <img src="assets/catgpt_gatway_logo.jpeg" width="200" alt="CatGPT Gateway Logo" />
</p>

<h1 align="center">CatGPT Gateway</h1>

<p align="center">
  <strong>Turn your ChatGPT or Claude account into a fully working OpenAI-compatible API.</strong><br/>
  No API keys needed. Just your browser login.
</p>

<p align="center">
  <a href="#quick-start">Quick Start</a> &bull;
  <a href="#providers">Providers</a> &bull;
  <a href="docs/API.md">API Docs</a> &bull;
  <a href="docs/SETUP.md">Full Setup Guide</a> &bull;
  <a href="docs/ARCHITECTURE.md">How It Works</a> &bull;
  <a href="CONTRIBUTING.md">Contributing</a>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/python-3.9+-blue?style=flat-square" alt="Python 3.9+" />
  <img src="https://img.shields.io/badge/providers-Claude_%7C_ChatGPT-purple?style=flat-square" alt="Providers" />
  <img src="https://img.shields.io/badge/API-OpenAI_compatible-green?style=flat-square" alt="OpenAI Compatible" />
  <img src="https://img.shields.io/badge/docker-ready-blue?style=flat-square" alt="Docker" />
  <img src="https://img.shields.io/github/license/GautamVhavle/CatGPT-Gateway?style=flat-square" alt="MIT License" />
  <img src="https://img.shields.io/github/stars/GautamVhavle/CatGPT-Gateway?style=flat-square" alt="Stars" />
</p>

---

## What is this?

You already pay for ChatGPT Plus or have a free Claude account. But the official APIs cost extra and the free tiers are limited.

**CatGPT Gateway** turns your existing browser session into a fully functional OpenAI-compatible API server. It runs a real browser in the background, automates the web UI, and exposes everything through standard API endpoints that work with the OpenAI Python SDK, LangChain, and anything that speaks the OpenAI protocol.

```python
# This just works. Point any OpenAI client at your local gateway.
from openai import OpenAI

client = OpenAI(base_url="http://localhost:8000/v1", api_key="dummy123")

response = client.chat.completions.create(
    model="claude-browser",  # or "catgpt-browser" for ChatGPT
    messages=[{"role": "user", "content": "Hello from my own API!"}]
)
print(response.choices[0].message.content)
```

That's it. Your Claude or ChatGPT subscription just became an API.

---

## Features

| Feature | Claude | ChatGPT |
|---|:---:|:---:|
| Chat completions | Yes | Yes |
| Multi-turn conversations | Yes | Yes |
| Tool / function calling | Yes | Yes |
| Image input (vision) | Yes | Yes |
| File attachments (PDF, DOCX, etc.) | Yes | Yes |
| Image generation (DALL-E) | -- | Yes |
| Interactive TUI (terminal chat) | Yes | Yes |
| OpenAI SDK compatible | Yes | Yes |
| LangChain compatible | Yes | Yes |
| Docker deployment | Yes | Yes |

---

## Providers

CatGPT Gateway supports two providers. Set `PROVIDER` in your `.env` file to switch.

### Claude (`PROVIDER=claude`)

Use your existing Anthropic account (free or Pro). The gateway connects to `claude.ai` and exposes Claude as an OpenAI-compatible API.

- Model ID: `claude-browser`
- Works with: Free tier, Pro, Team
- Image generation: Not supported (returns 501)

### ChatGPT (`PROVIDER=chatgpt`)

Use your existing OpenAI account (free or Plus). The gateway connects to `chatgpt.com` and exposes ChatGPT as an OpenAI-compatible API.

- Model ID: `catgpt-browser`
- Works with: Free tier, Plus, Team
- Image generation: Supported via DALL-E

---

## Quick Start

### Option 1: Docker (recommended)

```bash
# Clone the repo
git clone https://github.com/GautamVhavle/CatGPT-Gateway.git
cd CatGPT-Gateway

# Copy the example env and pick your provider
cp .env.example .env
# Edit .env -> set PROVIDER=claude or PROVIDER=chatgpt

# Build and start
docker compose up --build -d

# Open the browser UI to log in (one-time)
open http://localhost:6080/vnc.html
# Sign into Claude or ChatGPT using EMAIL + PASSWORD or Microsoft/Apple/magic link
# ⚠ Google login is blocked by Google in automated browser contexts — don't use it
# Close the noVNC tab when done - your session is saved to a Docker volume

# Verify it works
curl -H "Authorization: Bearer dummy123" http://localhost:8000/v1/models
```

### Option 2: Local (no Docker)

```bash
# Clone and setup
git clone https://github.com/GautamVhavle/CatGPT-Gateway.git
cd CatGPT-Gateway
python3 -m venv .venv && source .venv/bin/activate
pip install -r requirements.txt
patchright install chromium

# Copy env and pick your provider
cp .env.example .env

# First login (one-time - a browser window opens)
# Use email + password, Microsoft, Apple, or magic link — NOT Google OAuth
python scripts/first_login.py

# Start the API server
python -m src.api.server
# API is now live at http://localhost:8000
```

> For the full setup guide with Docker internals, Nix flake, systemd service, and troubleshooting, see [docs/SETUP.md](docs/SETUP.md).

---

## Usage

Once the server is running, you can use it with any OpenAI-compatible client.

### Python (OpenAI SDK)

```python
from openai import OpenAI

client = OpenAI(base_url="http://localhost:8000/v1", api_key="dummy123")

# Simple chat
response = client.chat.completions.create(
    model="claude-browser",
    messages=[{"role": "user", "content": "Explain quantum computing in simple terms"}]
)
print(response.choices[0].message.content)
```

### Python (LangChain)

```python
from langchain_openai import ChatOpenAI

llm = ChatOpenAI(
    model="claude-browser",
    base_url="http://localhost:8000/v1",
    api_key="dummy123",
)
response = llm.invoke("What are the best practices for REST API design?")
print(response.content)
```

### curl

```bash
curl -X POST http://localhost:8000/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer dummy123" \
  -d '{
    "model": "claude-browser",
    "messages": [{"role": "user", "content": "Hello!"}]
  }'
```

### Tool / Function Calling

Full round-trip tool calling works with both providers. Define tools, let the model call them, send results back.

```python
from langchain_openai import ChatOpenAI
from langchain_core.messages import HumanMessage, ToolMessage
from langchain_core.tools import tool

@tool
def get_weather(city: str) -> str:
    """Get the current weather for a city."""
    return f"Sunny, 25C in {city}"

llm = ChatOpenAI(
    model="claude-browser",
    base_url="http://localhost:8000/v1",
    api_key="dummy123",
)

llm_with_tools = llm.bind_tools([get_weather])
response = llm_with_tools.invoke("What's the weather in Tokyo?")

# Model returns structured tool calls
print(response.tool_calls)
# [{'name': 'get_weather', 'args': {'city': 'Tokyo'}, 'id': 'call_...'}]

# Execute the tool and send results back
messages = [HumanMessage(content="What's the weather in Tokyo?"), response]
result = get_weather.invoke(response.tool_calls[0]["args"])
messages.append(ToolMessage(content=result, tool_call_id=response.tool_calls[0]["id"]))

final = llm_with_tools.invoke(messages)
print(final.content)
# "It's sunny and 25C in Tokyo!"
```

### Image Generation (ChatGPT only)

```python
response = client.images.generate(
    prompt="A cyberpunk cat hacking into a mainframe",
    model="dall-e-3",
    n=1,
    response_format="b64_json",
)
```

> For the complete API reference with vision input, file attachments, image generation, tool_choice, and the custom REST API, see [docs/API.md](docs/API.md).

---

## Configuration

Copy `.env.example` to `.env` and edit to your needs:

```bash
cp .env.example .env
```

Key settings:

| Variable | Default | Description |
|---|---|---|
| `PROVIDER` | `chatgpt` | Which provider to use: `chatgpt` or `claude` |
| `BROWSER_DATA_DIR` | `./browser_data` | Browser profile directory (keeps your login) |
| `API_TOKEN` | `dummy123` | Bearer token for API authentication |
| `API_PORT` | `8000` | Port the API server listens on |
| `HEADLESS` | `false` | Run browser without display (not recommended) |

> See [.env.example](.env.example) for all available settings with descriptions.

---

## Testing

All test scripts auto-detect the active provider from your `.env` file.

```bash
source .venv/bin/activate

# Start the server (if not already running)
python -m src.api.server &

# Run individual test suites
python scripts/test_phase1.py           # Basic send/receive
python scripts/test_multi_turn.py       # Multi-turn conversations
python scripts/test_robust.py           # Tables, code blocks, long responses
python scripts/test_images.py           # Image detection
python scripts/test_langchain_tools.py  # LangChain + tool calling (needs server running)
```

Both providers have been tested and verified. See [docs/TEST_REPORT.md](docs/TEST_REPORT.md) for full results with inputs, outputs, and timings.

---

## How It Works (short version)

```
Your app (OpenAI SDK / LangChain / curl)
    |
    v
CatGPT Gateway (FastAPI on port 8000)
    |
    v
Real Chromium browser (automated via Patchright)
    |
    v
claude.ai or chatgpt.com (your logged-in session)
    |
    v
Response extracted from the page, formatted as OpenAI JSON, returned to your app
```

The gateway runs a real browser session with anti-detection measures (stealth patches, human-like typing, viewport jitter, persistent cookies). It types your message into the chat input, waits for the response to complete, extracts the text, and returns it in the standard OpenAI response format.

Tool calling is implemented via prompt engineering: tool definitions are injected as structured instructions, the model outputs JSON tool calls, and the gateway parses them into the proper OpenAI `tool_calls` format.

> For the full deep dive (browser lifecycle, stealth techniques, response detection, DOM selectors, Docker internals), see [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md).

---

## Known Limitations

- **No streaming** - Responses are returned all at once after completion. `stream=true` returns a 400 error.
- **Single concurrency** - One request at a time (browser is single-threaded). Requests are queued.
- **Response time** - Each request takes 5-30s depending on response length (real browser round-trip).
- **Session expiry** - Browser sessions expire after days/weeks. Re-login via noVNC or `first_login.py`.
- **UI changes** - If Claude or ChatGPT update their HTML, selectors may need updating. All selectors are centralized in `selectors.py` for easy fixes.
- **Tool calling** - Works via prompt engineering, not native API. Reliable for 1-7 tools. Very complex schemas may occasionally need a retry.

---

## Contributing

Contributions are welcome! Whether it's fixing a broken selector, adding a new provider, improving detection logic, or writing docs, we'd love your help.

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on how to get started.

---

## License

MIT License. See [LICENSE](LICENSE) for details.

---

<p align="center">
  If you find this project useful, consider giving it a star. It helps others discover it and keeps the project going.<br/>
  <a href="https://github.com/GautamVhavle/CatGPT-Gateway">
    <img src="https://img.shields.io/github/stars/GautamVhavle/CatGPT-Gateway?style=social" alt="Star on GitHub" />
  </a>
</p>
</file>

</files>
