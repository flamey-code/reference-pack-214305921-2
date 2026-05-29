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
.env.example
.gitignore
docker/docker-compose.yml
docker/Dockerfile
docs/TROUBLESHOOTING.md
INSTALL.md
LICENSE
pyproject.toml
README.md
requirements-dev.txt
requirements.txt
src/gpt_proxy/__init__.py
src/gpt_proxy/__main__.py
src/gpt_proxy/api/__init__.py
src/gpt_proxy/api/auth_router.py
src/gpt_proxy/api/router.py
src/gpt_proxy/cli.py
src/gpt_proxy/config.py
src/gpt_proxy/core/__init__.py
src/gpt_proxy/core/chatgpt_client.py
src/gpt_proxy/main.py
src/gpt_proxy/middleware/__init__.py
src/gpt_proxy/middleware/auth.py
src/gpt_proxy/middleware/error_handler.py
src/gpt_proxy/middleware/logging.py
src/gpt_proxy/middleware/rate_limit.py
src/gpt_proxy/services/__init__.py
src/gpt_proxy/services/auth_manager.py
src/gpt_proxy/services/browser_auth.py
src/gpt_proxy/services/cache.py
src/gpt_proxy/services/cost_tracker.py
src/gpt_proxy/services/health.py
src/gpt_proxy/services/key_manager.py
tests/__init__.py
tests/conftest.py
tests/integration/__init__.py
tests/integration/test_auth.py
tests/unit/__init__.py
tests/unit/test_auth.py
tests/unit/test_cache.py
tests/unit/test_cost_tracker.py
tests/unit/test_key_manager.py
tests/unit/test_router.py
```

# Files

## File: .env.example
````
# Application
APP_NAME=gpt-proxy
APP_ENV=development
APP_DEBUG=true
APP_HOST=0.0.0.0
APP_PORT=8000
APP_WORKERS=1

# OpenAI
OPENAI_API_BASE_URL=https://api.openai.com/v1
OPENAI_API_KEYS=sk-your-key-here
OPENAI_KEY_ROTATION_STRATEGY=round-robin

# Rate Limiting
RATE_LIMIT_ENABLED=true
RATE_LIMIT_REQUESTS_PER_MINUTE=60
RATE_LIMIT_TOKENS_PER_MINUTE=90000

# Caching
CACHE_ENABLED=false
CACHE_BACKEND=memory
CACHE_TTL_SECONDS=3600
# REDIS_URL=redis://localhost:6379

# Database
DATABASE_URL=sqlite+aiosqlite:///./gpt_proxy.db

# Logging
LOG_LEVEL=INFO
LOG_FORMAT=json

# Cost Tracking
COST_TRACKING_ENABLED=true

# Browser Auth (代理设置，如果无法访问chat.openai.com请配置)
BROWSER_PROXY=http://127.0.0.1:7890
BROWSER_PROFILE_DIR=./browser_profile
````

## File: .gitignore
````
# Byte-compiled / optimized / DLL files
__pycache__/
*.py[cod]
*$py.class

# C extensions
*.so

# Distribution / packaging
.Python
build/
develop-eggs/
dist/
downloads/
eggs/
.eggs/
lib/
lib64/
parts/
sdist/
var/
wheels/
*.egg-info/
.installed.cfg
*.egg

# PyInstaller
*.manifest
*.spec

# Installer logs
pip-log.txt
pip-delete-this-directory.txt

# Unit test / coverage reports
htmlcov/
.tox/
.nox/
.coverage
.coverage.*
.cache
nosetests.xml
coverage.xml
*.cover
*.py,cover
.hypothesis/
.pytest_cache/

# Translations
*.mo
*.pot

# Environments
.env
.venv
env/
venv/
ENV/
env.bak/
venv.bak/

# IDE
.idea/
.vscode/
*.swp
*.swo
*~

# Database
*.db
*.sqlite
*.sqlite3

# Logs
logs/
*.log

# macOS
.DS_Store

# Project specific
gpt_proxy.db
browser_profile/
test_profile/
sessions.json
````

## File: docker/docker-compose.yml
````yaml
version: '3.8'

services:
  proxy:
    build:
      context: ..
      dockerfile: docker/Dockerfile
    ports:
      - "8000:8000"
    environment:
      - OPENAI_API_KEYS=${OPENAI_API_KEYS:-}
      - OPENAI_KEY_ROTATION_STRATEGY=round-robin
      - RATE_LIMIT_REQUESTS_PER_MINUTE=60
      - CACHE_ENABLED=false
      - LOG_LEVEL=INFO
      - LOG_FORMAT=json
    volumes:
      - ./logs:/app/logs
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8000/health"]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 5s
    restart: unless-stopped

  # Optional Redis for distributed caching
  # redis:
  #   image: redis:7-alpine
  #   ports:
  #     - "6379:6379"
  #   volumes:
  #     - redis_data:/data

# volumes:
#   redis_data:
````

## File: docker/Dockerfile
````
FROM python:3.11-slim

WORKDIR /app

# Install system dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    curl \
    && rm -rf /var/lib/apt/lists/*

# Install Python dependencies
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy application
COPY src/ ./src/
COPY pyproject.toml .

# Create non-root user
RUN useradd -m -u 1000 appuser && chown -R appuser:appuser /app
USER appuser

# Expose port
EXPOSE 8000

# Health check
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8000/health || exit 1

# Run application
CMD ["python", "-m", "gpt_proxy", "serve", "--host", "0.0.0.0", "--port", "8000"]
````

## File: docs/TROUBLESHOOTING.md
````markdown
# Troubleshooting Guide

## Common Errors

### "Expecting value: line 1 column 1 (char 0)"

This error occurs when the ChatGPT API returns an invalid response.

**Causes:**
1. **Missing proxy configuration** - If you're behind a proxy, set `BROWSER_PROXY` in `.env`
2. **Cloudflare challenge** - ChatGPT detected automated access
3. **Expired session token** - Get a fresh token from your browser

**Solutions:**
1. Configure proxy in `.env`:
   ```
   BROWSER_PROXY=http://127.0.0.1:7890
   ```
2. Login via browser first, then try again
3. Clear browser profile: `rm -rf ./browser_profile`

### Cloudflare Challenge Detected

**Symptoms:**
- Error message mentions "Cloudflare" or "challenge"
- Browser shows "Just a moment..." page

**Solutions:**
1. Open ChatGPT in your regular browser
2. Complete any CAPTCHA challenges
3. Wait a few minutes before retrying
4. Consider using a residential proxy

### Connection Timeout

**Symptoms:**
- Error: "Request timed out"
- Login takes too long

**Solutions:**
1. Check your network connectivity
2. Verify proxy is running (if configured)
3. Increase timeout: `gpt-proxy login --timeout 600`

### Proxy Connection Issues

**Symptoms:**
- Error: "Connection refused"
- Error: "Proxy connection failed"

**Solutions:**
1. Verify proxy URL format: `http://host:port`
2. Check if proxy server is running
3. Test proxy with curl: `curl -x http://proxy:port https://chat.openai.com`

### Session Token Invalid

**Symptoms:**
- Error: "Invalid session token"
- Error: "Token expired"

**Solutions:**
1. Get a fresh session token from browser
2. Clear browser profile and re-login
3. Check if you can access chat.openai.com normally

## Configuration

### Environment Variables

Create a `.env` file in the project root:

```env
# Proxy settings (required if behind firewall)
BROWSER_PROXY=http://127.0.0.1:7890

# HTTP client settings
HTTP_TIMEOUT=30.0
HTTP_CONNECT_TIMEOUT=10.0

# Browser profile directory
BROWSER_PROFILE_DIR=./browser_profile
```

### Debug Mode

Enable debug logging:

```env
LOG_LEVEL=DEBUG
APP_DEBUG=true
```

## Getting Help

1. Check server logs for detailed error messages
2. Try the manual token method: `gpt-proxy help-token`
3. Open an issue: https://github.com/art3m1s-tju/GPT_reverse/issues
````

## File: INSTALL.md
````markdown
# GPT Reverse Proxy 完整安装教程

## 系统要求

- Python 3.11+
- pip 或 conda

---

## 快速开始（推荐）

### 1. 克隆并安装

```bash
git clone https://github.com/art3m1s-tju/GPT_reverse.git
cd GPT_reverse
pip install -e .
```

### 2. 安装浏览器自动化（可选，用于自动登录）

```bash
pip install playwright
playwright install chromium
```

### 3. 启动服务

```bash
gpt-proxy serve
```

### 4. 登录

**方式一：浏览器自动登录（推荐）**
```bash
gpt-proxy login
```
会自动打开浏览器，登录ChatGPT后自动获取token。

**方式二：手动输入token**
```bash
curl -X POST http://localhost:8000/auth/login -H "Content-Type: application/json" -d '{"session_token": "你的token"}'
```

---

## 详细安装方法

### 方法一：pip 安装（推荐）

#### 1. 克隆仓库

```bash
git clone https://github.com/art3m1s-tju/GPT_reverse.git
cd GPT_reverse
```

#### 2. 创建虚拟环境（可选但推荐）

```bash
# 使用 venv
python -m venv venv
source venv/bin/activate  # Linux/Mac
# 或 venv\Scripts\activate  # Windows

# 或使用 conda
conda create -n gpt-proxy python=3.11
conda activate gpt-proxy
```

#### 3. 安装依赖

```bash
pip install -e .
```

#### 4. 安装浏览器自动化（可选）

```bash
pip install playwright
playwright install chromium
```

#### 5. 启动服务

```bash
gpt-proxy serve
```

服务将在 http://localhost:8000 启动

---

### 方法二：Docker 安装

```bash
git clone https://github.com/art3m1s-tju/GPT_reverse.git
cd GPT_reverse
docker build -t gpt-proxy -f docker/Dockerfile .
docker run -p 8000:8000 gpt-proxy
```

---

### 方法三：直接运行（无需安装）

```bash
git clone https://github.com/art3m1s-tju/GPT_reverse.git
cd GPT_reverse
pip install fastapi uvicorn httpx pydantic pydantic-settings typer rich playwright
playwright install chromium
python -m gpt_proxy serve
```

---

## 登录方式

### 方式一：浏览器自动登录（推荐）

**前提：已安装 playwright 和 chromium**

```bash
# 启动服务
gpt-proxy serve

# 新开终端，运行登录命令
gpt-proxy login
```

系统会自动打开浏览器窗口：
1. 在浏览器中登录你的ChatGPT账号
2. 登录成功后自动提取token
3. 返回session_id用于后续API调用

**优点：**
- 无需手动复制cookie
- 浏览器profile持久化，下次登录更快
- 可视化操作，体验更好

---

### 方式二：手动获取 Session Token

如果不想安装playwright，可以手动获取token。

#### 步骤1：获取 Token

**方法A：浏览器开发者工具**

1. 打开 https://chat.openai.com 并登录
2. 按 `F12` 打开开发者工具
3. 点击 **Application** 标签
4. 左侧菜单：**Cookies** → **chat.openai.com**
5. 找到 `__Secure-next-auth.session-token`
6. 复制它的值

**方法B：浏览器控制台**

在 chat.openai.com 页面打开控制台（F12 → Console），粘贴：

```javascript
document.cookie.split('; ').find(c => c.startsWith('__Secure-next-auth.session-token='))?.split('=')[1]
```

复制输出的字符串。

#### 步骤2：登录

```bash
curl -X POST http://localhost:8000/auth/login -H "Content-Type: application/json" -d '{"session_token": "你的token"}'
```

---

## 使用 API

### 调用聊天接口

```bash
curl -X POST http://localhost:8000/v1/chat/completions \
  -H "Authorization: Bearer 你的session_id" \
  -H "Content-Type: application/json" \
  -d '{"model": "gpt-4", "messages": [{"role": "user", "content": "你好"}]}'
```

### Python 客户端

```python
import httpx

# 方式一：使用浏览器登录后的session_id
session_id = "你的session_id"

# 方式二：通过API登录
# login_resp = httpx.post(
#     "http://localhost:8000/auth/login",
#     json={"session_token": "你的token"}
# )
# session_id = login_resp.json()["session_id"]

# 聊天
client = httpx.Client(
    base_url="http://localhost:8000",
    headers={"Authorization": f"Bearer {session_id}"}
)

response = client.post(
    "/v1/chat/completions",
    json={
        "model": "gpt-4",
        "messages": [{"role": "user", "content": "你好"}]
    }
)
print(response.json())
```

### 流式输出

```python
import httpx

session_id = "你的session_id"

with httpx.stream(
    "POST",
    "http://localhost:8000/v1/chat/completions",
    headers={"Authorization": f"Bearer {session_id}"},
    json={
        "model": "gpt-4",
        "messages": [{"role": "user", "content": "讲个故事"}],
        "stream": True
    }
) as response:
    for line in response.iter_lines():
        if line:
            print(line)
```

---

## CLI 命令

| 命令 | 说明 |
|------|------|
| `gpt-proxy serve` | 启动服务器 |
| `gpt-proxy login` | 浏览器自动登录 |
| `gpt-proxy help-token` | 显示如何获取token |
| `gpt-proxy version` | 显示版本 |

---

## API 端点

| 端点 | 方法 | 说明 |
|------|------|------|
| `/auth/login/browser` | POST | 浏览器自动登录 |
| `/auth/login` | POST | 手动token登录 |
| `/auth/login/status` | GET | 检查登录状态 |
| `/auth/sessions` | GET | 列出所有会话 |
| `/auth/logout` | POST | 登出 |
| `/v1/chat/completions` | POST | 聊天补全 |
| `/v1/models` | GET | 列出模型 |

---

## 可用模型

- `gpt-4` - 需要 ChatGPT Plus 订阅
- `gpt-4o` - 需要 ChatGPT Plus 订阅
- `gpt-3.5-turbo` - 免费账户可用

---

## 常见问题

### Q: playwright install chromium 下载失败？
A: 网络问题，可以：
1. 使用代理
2. 或使用手动获取token的方式

### Q: 登录失败怎么办？
A: Session token 会过期，重新登录获取新的 token。

### Q: 401 错误？
A: Session 过期，重新登录获取新的 session_id。

### Q: 没有响应？
A: 检查网络连接，确保能访问 chat.openai.com。

### Q: Windows 下命令不生效？
A: 确保使用正确的路径分隔符，或使用 PowerShell。

---

## API 文档

启动服务后访问：
- Swagger UI: http://localhost:8000/docs
- ReDoc: http://localhost:8000/redoc
````

## File: LICENSE
````
MIT License

Copyright (c) 2026 art3m1s

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
````

## File: pyproject.toml
````toml
[project]
name = "gpt-proxy"
version = "0.2.0"
description = "Local ChatGPT reverse proxy with API key management, rate limiting, and cost tracking"
readme = "README.md"
license = {text = "MIT"}
requires-python = ">=3.11"
authors = [
    {name = "art3m1s-tju"}
]
keywords = ["openai", "chatgpt", "proxy", "api", "reverse-proxy"]
classifiers = [
    "Development Status :: 3 - Alpha",
    "Intended Audience :: Developers",
    "License :: OSI Approved :: MIT License",
    "Programming Language :: Python :: 3",
    "Programming Language :: Python :: 3.11",
    "Programming Language :: Python :: 3.12",
]

dependencies = [
    "fastapi>=0.109.0",
    "uvicorn[standard]>=0.27.0",
    "httpx>=0.26.0",
    "pydantic>=2.5.0",
    "pydantic-settings>=2.1.0",
    "tiktoken>=0.5.0",
    "typer>=0.9.0",
    "rich>=13.7.0",
    "sqlalchemy[asyncio]>=2.0.0",
    "aiosqlite>=0.19.0",
    "python-dotenv>=1.0.0",
    "playwright>=1.40.0",
    "curl_cffi>=0.5.0",
    "pybase64>=1.3.0",
    "diskcache>=5.6.0",
]

[project.optional-dependencies]
dev = [
    "pytest>=7.4.0",
    "pytest-asyncio>=0.23.0",
    "pytest-cov>=4.1.0",
    "httpx>=0.26.0",
    "ruff>=0.1.0",
    "mypy>=1.8.0",
]
redis = [
    "redis>=5.0.0",
]

[project.scripts]
gpt-proxy = "gpt_proxy.cli:app"

[build-system]
requires = ["hatchling"]
build-backend = "hatchling.build"

[tool.hatch.build.targets.wheel]
packages = ["src/gpt_proxy"]

[tool.ruff]
line-length = 100
target-version = "py311"

[tool.ruff.lint]
select = ["E", "F", "I", "N", "W", "UP"]
ignore = ["E501"]

[tool.mypy]
python_version = "3.11"
strict = true
ignore_missing_imports = true

[tool.pytest.ini_options]
asyncio_mode = "auto"
testpaths = ["tests"]
````

## File: README.md
````markdown
# GPT Reverse Proxy

**Use ChatGPT without API keys!** Login with your ChatGPT account session token.

## Why This Project?

- **No API keys needed** - Use your existing ChatGPT account
- **Free to use** - No need to pay for API access
- **Same models** - Access GPT-4, GPT-3.5, etc. through your subscription
- **Simple setup** - Just copy a cookie from your browser

## Quick Start

### 1. Install

```bash
pip install -e .
```

### 2. Get Session Token

1. Go to [chat.openai.com](https://chat.openai.com) and login
2. Press `F12` to open DevTools
3. Go to **Application** > **Cookies** > **chat.openai.com**
4. Find `__Secure-next-auth.session-token` cookie
5. Copy its value

Or run this in browser console:
```javascript
document.cookie.split('; ').find(c => c.startsWith('__Secure-next-auth.session-token='))?.split('=')[1]
```

### 3. Start Proxy

```bash
gpt-proxy serve
```

### 4. Login

```bash
curl -X POST http://localhost:8000/auth/login \
  -H "Content-Type: application/json" \
  -d '{"session_token": "your-token-here"}'
```

Response:
```json
{
  "session_id": "abc123...",
  "user_email": "your@email.com",
  "expires_at": "2024-...",
  "message": "Login successful. Use session_id as Bearer token in API requests."
}
```

### 5. Use API

```bash
curl -X POST http://localhost:8000/v1/chat/completions \
  -H "Authorization: Bearer <session_id>" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gpt-4",
    "messages": [{"role": "user", "content": "Hello!"}]
  }'
```

## Python Client

```python
import httpx

# Login
login_resp = httpx.post(
    "http://localhost:8000/auth/login",
    json={"session_token": "your-token"}
)
session_id = login_resp.json()["session_id"]

# Chat
client = httpx.Client(
    base_url="http://localhost:8000/v1",
    headers={"Authorization": f"Bearer {session_id}"}
)

response = client.post(
    "/chat/completions",
    json={
        "model": "gpt-4",
        "messages": [{"role": "user", "content": "Hello!"}]
    }
)
print(response.json())
```

## API Endpoints

| Endpoint | Description |
|----------|-------------|
| `POST /auth/login` | Login with session token |
| `GET /auth/help` | How to get session token |
| `GET /auth/sessions` | List active sessions |
| `POST /auth/logout` | Logout session |
| `POST /v1/chat/completions` | Chat completions |
| `GET /v1/models` | List models |

## CLI Commands

```bash
# Start server
gpt-proxy serve --port 8000

# Show how to get token
gpt-proxy help-token

# Show version
gpt-proxy version
```

## Docker

```bash
docker build -t gpt-proxy -f docker/Dockerfile .
docker run -p 8000:8000 gpt-proxy
```

## Notes

- **Session tokens expire** - Get a fresh token if you get 401 errors
- **Rate limits apply** - Same limits as your ChatGPT web access
- **Subscription required** - GPT-4 requires ChatGPT Plus subscription
- **Unofficial API** - This uses ChatGPT's backend API, not the official OpenAI API

## License

MIT
````

## File: requirements-dev.txt
````
-r requirements.txt
pytest>=7.4.0
pytest-asyncio>=0.23.0
pytest-cov>=4.1.0
ruff>=0.1.0
mypy>=1.8.0
````

## File: requirements.txt
````
fastapi>=0.109.0
uvicorn[standard]>=0.27.0
httpx>=0.26.0
pydantic>=2.5.0
pydantic-settings>=2.1.0
tiktoken>=0.5.0
typer>=0.9.0
rich>=13.7.0
sqlalchemy[asyncio]>=2.0.0
aiosqlite>=0.19.0
python-dotenv>=1.0.0
playwright>=1.40.0
````

## File: src/gpt_proxy/__init__.py
````python
"""GPT Reverse Proxy - Local ChatGPT reverse proxy with API key management."""

__version__ = "0.2.0"
````

## File: src/gpt_proxy/__main__.py
````python
"""Entry point for python -m gpt_proxy."""

from gpt_proxy.cli import app

if __name__ == "__main__":
    app()
````

## File: src/gpt_proxy/api/__init__.py
````python

````

## File: src/gpt_proxy/api/auth_router.py
````python
"""Authentication API routes."""

from fastapi import APIRouter, HTTPException, Depends
from fastapi.responses import RedirectResponse
from pydantic import BaseModel
from typing import Optional
import httpx

from gpt_proxy.services.auth_manager import get_auth_manager, AuthManager

router = APIRouter(prefix="/auth", tags=["authentication"])


class SessionTokenLogin(BaseModel):
    """Login with ChatGPT session token."""
    session_token: str


class LoginResponse(BaseModel):
    """Response after successful login."""
    session_id: str
    user_email: str
    expires_at: str
    message: str = "Login successful. Use session_id as Bearer token in API requests."


class SessionInfo(BaseModel):
    """Session information."""
    session_id: str
    email: str
    is_active: bool
    request_count: int
    expires_at: str


class LogoutResponse(BaseModel):
    """Logout response."""
    message: str


def get_auth() -> AuthManager:
    """Dependency to get auth manager."""
    return get_auth_manager()


@router.post("/login", response_model=LoginResponse)
async def login_with_session_token(
    data: SessionTokenLogin,
    auth: AuthManager = Depends(get_auth),
):
    """Login using ChatGPT session token.

    ## How to get session token:

    1. Go to https://chat.openai.com and login
    2. Open browser DevTools (F12)
    3. Go to Application > Cookies > chat.openai.com
    4. Find '__Secure-next-auth.session-token' cookie
    5. Copy its value and paste here

    Or run in browser console:
    ```javascript
    document.cookie.split('; ').find(c => c.startsWith('__Secure-next-auth.session-token='))?.split('=')[1]
    ```
    """
    session = await auth.exchange_session_token(data.session_token)

    if not session:
        raise HTTPException(
            status_code=401,
            detail="Invalid session token or token expired. Please get a fresh token from chat.openai.com",
        )

    auth.create_session(session)

    return LoginResponse(
        session_id=session.session_id,
        user_email=session.email,
        expires_at=session.expires_at.isoformat(),
    )


@router.get("/sessions", response_model=list[SessionInfo])
async def list_sessions(auth: AuthManager = Depends(get_auth)):
    """List all active sessions."""
    return auth.list_sessions()


@router.post("/logout", response_model=LogoutResponse)
async def logout(session_id: str, auth: AuthManager = Depends(get_auth)):
    """Logout and invalidate session."""
    if auth.invalidate_session(session_id):
        return LogoutResponse(message="Session invalidated successfully")
    raise HTTPException(status_code=404, detail="Session not found")


@router.get("/status")
async def auth_status(auth: AuthManager = Depends(get_auth)):
    """Get authentication system status."""
    sessions = auth.list_sessions()
    active = sum(1 for s in sessions if s["is_active"])
    return {
        "total_sessions": len(sessions),
        "active_sessions": active,
        "status": "operational",
    }


@router.get("/help")
async def auth_help():
    """Get help on how to obtain session token."""
    return {
        "title": "How to get ChatGPT session token",
        "steps": [
            "1. Go to https://chat.openai.com and login with your OpenAI account",
            "2. Open browser DevTools (press F12)",
            "3. Go to Application tab > Cookies > chat.openai.com",
            "4. Find cookie named '__Secure-next-auth.session-token'",
            "5. Copy its value",
            "6. POST to /auth/login with {'session_token': '<your_token>'}",
        ],
        "browser_console_method": "Run this in browser console on chat.openai.com:\n"
        "document.cookie.split('; ').find(c => c.startsWith('__Secure-next-auth.session-token='))?.split('=')[1]",
        "note": "Session tokens expire after some time. You may need to refresh them periodically.",
    }


@router.post("/login/browser", response_model=LoginResponse)
async def login_via_browser(
    timeout: int = 300,
    auth: AuthManager = Depends(get_auth),
):
    """Login via automated browser.

    Opens a browser window for user to login to ChatGPT,
    then automatically extracts session token.

    Args:
        timeout: Maximum seconds to wait for login (default 5 minutes)
    """
    from gpt_proxy.services.browser_auth import get_browser_auth, close_browser_auth

    browser_auth = get_browser_auth()

    try:
        # Get session token from browser
        session_token = await browser_auth.get_session_token(
            wait_for_login=True,
            timeout=timeout
        )

        if not session_token:
            raise HTTPException(
                status_code=401,
                detail="Login failed or timed out. Please try again."
            )

        # Exchange for access token
        session = await auth.exchange_session_token(session_token)
        if not session:
            raise HTTPException(
                status_code=401,
                detail="Failed to exchange session token."
            )

        auth.create_session(session)

        return LoginResponse(
            session_id=session.session_id,
            user_email=session.email,
            expires_at=session.expires_at.isoformat(),
            message="Browser login successful!"
        )
    finally:
        await close_browser_auth()


@router.get("/login/status")
async def check_browser_login_status():
    """Check if user is logged in via browser profile."""
    from gpt_proxy.services.browser_auth import get_browser_auth

    browser_auth = get_browser_auth()
    await browser_auth.initialize(headless=True)

    try:
        session_token = await browser_auth.get_session_token(wait_for_login=False)
        return {
            "logged_in": session_token is not None,
            "message": "Already logged in" if session_token else "Not logged in"
        }
    finally:
        await browser_auth.close()
````

## File: src/gpt_proxy/api/router.py
````python
"""Main API router with ChatGPT backend proxy using conversation endpoint."""

from fastapi import APIRouter, Request, HTTPException
from fastapi.responses import Response, StreamingResponse
from logging import getLogger
import json

from gpt_proxy.services.auth_manager import get_auth_manager
from gpt_proxy.core.chatgpt_client import ChatGPTClient
from gpt_proxy.config import settings

logger = getLogger(__name__)
router = APIRouter()


async def get_access_token(request: Request) -> tuple[str, str]:
    """Extract and validate session ID from request, return (access_token, session_id)."""
    auth_header = request.headers.get("Authorization", "")

    if not auth_header.startswith("Bearer "):
        raise HTTPException(
            status_code=401,
            detail={
                "error": {
                    "message": "Missing Authorization header. Login at /auth/login first.",
                    "type": "authentication_error",
                    "code": "missing_auth",
                }
            },
        )

    session_id = auth_header[7:]  # Remove "Bearer " prefix
    auth_manager = get_auth_manager()

    access_token = await auth_manager.get_valid_token(session_id)
    if not access_token:
        raise HTTPException(
            status_code=401,
            detail={
                "error": {
                    "message": "Session expired or invalid. Please login again at /auth/login",
                    "type": "authentication_error",
                    "code": "session_expired",
                }
            },
        )

    return access_token, session_id


@router.post("/v1/chat/completions")
async def chat_completions(request: Request):
    """Proxy chat completions to ChatGPT conversation endpoint.

    Converts OpenAI Chat Completions API format to ChatGPT web API format.
    """
    access_token, session_id = await get_access_token(request)

    try:
        body = await request.json()
    except json.JSONDecodeError:
        raise HTTPException(
            status_code=400,
            detail={"error": {"message": "Invalid JSON body", "type": "invalid_request_error"}},
        )

    model = body.get("model", "gpt-4o-mini")
    messages = body.get("messages", [])
    stream = body.get("stream", False)

    # Optional conversation continuation
    conversation_id = body.get("conversation_id")
    parent_message_id = body.get("parent_message_id")

    if not messages:
        raise HTTPException(
            status_code=400,
            detail={"error": {"message": "messages is required", "type": "invalid_request_error"}},
        )

    # Get proxy from settings
    proxy_url = settings.browser_proxy if settings.browser_proxy else None

    client = ChatGPTClient(
        access_token=access_token,
        proxy_url=proxy_url,
    )

    try:
        result = await client.chat_completions(
            model=model,
            messages=messages,
            stream=stream,
            conversation_id=conversation_id,
            parent_message_id=parent_message_id,
        )

        if stream:
            return StreamingResponse(
                result,
                media_type="text/event-stream",
                headers={
                    "Cache-Control": "no-cache",
                    "Connection": "keep-alive",
                    "X-Accel-Buffering": "no",
                },
            )
        else:
            return Response(
                content=json.dumps(result),
                media_type="application/json",
            )

    except Exception as e:
        logger.error(f"Chat completion error: {e}")
        raise HTTPException(
            status_code=500,
            detail={"error": {"message": str(e), "type": "internal_error"}},
        )
    finally:
        await client.close()


@router.get("/v1/models")
async def list_models(request: Request):
    """List available models."""
    # Return static list of supported models
    models = [
        {"id": "gpt-3.5-turbo", "object": "model", "owned_by": "openai"},
        {"id": "gpt-4", "object": "model", "owned_by": "openai"},
        {"id": "gpt-4-turbo", "object": "model", "owned_by": "openai"},
        {"id": "gpt-4o", "object": "model", "owned_by": "openai"},
        {"id": "gpt-4o-mini", "object": "model", "owned_by": "openai"},
        {"id": "o1", "object": "model", "owned_by": "openai"},
        {"id": "o1-mini", "object": "model", "owned_by": "openai"},
        {"id": "o1-preview", "object": "model", "owned_by": "openai"},
        {"id": "auto", "object": "model", "owned_by": "openai"},
    ]

    return {
        "object": "list",
        "data": models,
    }
````

## File: src/gpt_proxy/cli.py
````python
"""Typer CLI for GPT Proxy."""

import typer
from typing import Optional
from rich.console import Console
from rich.table import Table
import uvicorn

app = typer.Typer(
    name="gpt-proxy",
    help="ChatGPT reverse proxy - Use ChatGPT without API keys!",
)
console = Console()


@app.command()
def serve(
    host: str = typer.Option("0.0.0.0", "--host", "-h", help="Host to bind to"),
    port: int = typer.Option(8000, "--port", "-p", help="Port to bind to"),
    reload: bool = typer.Option(False, "--reload", "-r", help="Enable auto-reload"),
):
    """Start the proxy server."""
    console.print("[green]Starting ChatGPT Reverse Proxy...[/green]")
    console.print(f"[cyan]Server: http://{host}:{port}[/cyan]")
    console.print(f"[cyan]Docs: http://{host}:{port}/docs[/cyan]")
    console.print("")
    console.print("[yellow]How to use:[/yellow]")
    console.print("1. Login to chat.openai.com")
    console.print("2. Get session token from browser cookies")
    console.print("3. POST to /auth/login with session token")
    console.print("4. Use returned session_id as Bearer token")
    console.print("")

    uvicorn.run(
        "gpt_proxy.main:app",
        host=host,
        port=port,
        reload=reload,
    )


@app.command()
def help_token():
    """Show how to get ChatGPT session token."""
    console.print("[bold green]How to get ChatGPT session token:[/bold green]")
    console.print("")
    console.print("[bold]Method 1: Browser DevTools[/bold]")
    console.print("1. Go to https://chat.openai.com and login")
    console.print("2. Press F12 to open DevTools")
    console.print("3. Go to Application > Cookies > chat.openai.com")
    console.print("4. Find '__Secure-next-auth.session-token'")
    console.print("5. Copy its value")
    console.print("")
    console.print("[bold]Method 2: Browser Console[/bold]")
    console.print("Run this in browser console on chat.openai.com:")
    console.print("")
    console.print("[cyan]document.cookie.split('; ').find(c => c.startsWith('__Secure-next-auth.session-token='))?.split('=')[1][/cyan]")
    console.print("")
    console.print("[yellow]Note: Session tokens expire periodically. Get a fresh one if login fails.[/yellow]")


@app.command()
def version():
    """Show version information."""
    from gpt_proxy import __version__
    console.print(f"[green]GPT Proxy v{__version__}[/green]")


@app.command()
def login(
    token: Optional[str] = typer.Option(
        None,
        "--token",
        help="ChatGPT session token (__Secure-next-auth.session-token). If omitted, you'll be prompted.",
    ),
    server: str = typer.Option(
        "http://localhost:8000", "--server", "-s", help="Proxy server URL"
    ),
    timeout: int = typer.Option(60, "--timeout", "-t", help="Request timeout in seconds"),
):
    """Login by submitting a ChatGPT session token.

    How to get the token:
      1. Open https://chatgpt.com in your normal browser and login.
      2. F12 -> Application -> Cookies -> chatgpt.com
      3. Copy the value of '__Secure-next-auth.session-token'.
      4. Run: gpt-proxy login --token <value>   (or run with no flag to be prompted)
    """
    import httpx

    if not token:
        console.print("[yellow]Paste your ChatGPT session token (input hidden):[/yellow]")
        token = typer.prompt("session_token", hide_input=True)

    token = token.strip()
    if not token:
        console.print("[red]Empty token. Aborting.[/red]")
        raise typer.Exit(code=1)

    console.print("[cyan]Submitting token to proxy server...[/cyan]")

    try:
        with httpx.Client(timeout=timeout, trust_env=False) as client:
            response = client.post(
                f"{server.rstrip('/')}/auth/login",
                json={"session_token": token},
            )

            if response.status_code == 200:
                data = response.json()
                console.print("[green]Login successful![/green]")
                console.print(f"[cyan]Email: {data['user_email']}[/cyan]")
                console.print(f"[cyan]Session ID: {data['session_id']}[/cyan]")
                console.print("")
                console.print("[yellow]Use this session_id as Bearer token:[/yellow]")
                console.print(f"[white]Authorization: Bearer {data['session_id']}[/white]")
                return

            error_detail = "Unknown error"
            try:
                error_data = response.json()
                error_detail = error_data.get("detail", str(error_data))
            except Exception:
                error_detail = response.text or "Unknown error"

            console.print(
                f"[red]Login failed (status {response.status_code}): {error_detail}[/red]"
            )
            if response.status_code == 401:
                console.print("")
                console.print("[yellow]Common causes:[/yellow]")
                console.print("  - Token expired or copied incompletely (it's long — make sure you got all of it)")
                console.print("  - Wrong cookie copied; the right one is '__Secure-next-auth.session-token'")
                console.print("  - Server can't reach chatgpt.com (check BROWSER_PROXY in .env)")

    except httpx.ConnectError:
        console.print("[red]Error: Could not connect to server.[/red]")
        console.print("[yellow]Start it first: gpt-proxy serve[/yellow]")
    except httpx.TimeoutException:
        console.print("[red]Error: Request timed out.[/red]")
    except Exception as e:
        console.print(f"[red]Unexpected error: {e}[/red]")


if __name__ == "__main__":
    app()
````

## File: src/gpt_proxy/config.py
````python
"""Configuration using Pydantic Settings."""

from pydantic_settings import BaseSettings, SettingsConfigDict
from pydantic import field_validator
from typing import Literal


class Settings(BaseSettings):
    """Application configuration loaded from environment variables."""

    model_config = SettingsConfigDict(
        env_file=".env",
        env_file_encoding="utf-8",
        case_sensitive=False,
        extra="ignore",
    )

    # Application
    app_name: str = "gpt-proxy"
    app_env: Literal["development", "production"] = "development"
    app_debug: bool = False
    app_host: str = "0.0.0.0"
    app_port: int = 8000
    app_workers: int = 1

    # OpenAI
    openai_api_base_url: str = "https://api.openai.com/v1"
    openai_api_keys: list[str] = []
    openai_key_rotation_strategy: Literal["round-robin", "least-used", "random"] = "round-robin"

    # Rate Limiting
    rate_limit_enabled: bool = True
    rate_limit_requests_per_minute: int = 60
    rate_limit_tokens_per_minute: int = 90000

    # Caching
    cache_enabled: bool = False
    cache_backend: Literal["memory", "redis"] = "memory"
    cache_ttl_seconds: int = 3600
    redis_url: str = "redis://localhost:6379"

    # Database
    database_url: str = "sqlite+aiosqlite:///./gpt_proxy.db"

    # Logging
    log_level: str = "INFO"
    log_format: Literal["json", "text"] = "json"

    # Cost Tracking
    cost_tracking_enabled: bool = True

    # Browser Auth (代理设置)
    browser_proxy: str = ""  # 例如: "http://127.0.0.1:7890"
    browser_profile_dir: str = "./browser_profile"

    # HTTP Client
    http_timeout: float = 30.0
    http_connect_timeout: float = 10.0

    # ChatGPT API settings
    conversation_only: bool = False  # Skip sentinel verification (may trigger rate limits)
    pow_difficulty: int = 3  # Minimum POW difficulty to solve (lower = harder)
    arkose_token_url: str = ""  # Arkose token service URL (required for free accounts)

    @field_validator("openai_api_keys", mode="before")
    @classmethod
    def parse_keys(cls, v: str | list[str]) -> list[str]:
        """Parse comma-separated API keys from environment."""
        if isinstance(v, str):
            return [k.strip() for k in v.split(",") if k.strip()]
        return v


settings = Settings()
````

## File: src/gpt_proxy/core/__init__.py
````python

````

## File: src/gpt_proxy/core/chatgpt_client.py
````python
"""ChatGPT backend API client using conversation endpoint.

Based on reverse-engineered ChatGPT web API from chat2api project.
https://github.com/lanqian528/chat2api
"""

import hashlib
import json
import random
import re
import time
import uuid
from typing import AsyncIterator
from logging import getLogger
from datetime import datetime, timezone, timedelta

from curl_cffi.requests import AsyncSession
import pybase64

from gpt_proxy.config import settings

logger = getLogger(__name__)

# Cache for DPL and scripts
_cached_dpl: str = ""
_cached_scripts: list[str] = []
_cached_time: int = 0

# Model mapping from OpenAI API names to ChatGPT internal names
MODEL_MAP = {
    "gpt-3.5-turbo": "text-davinci-002-render-sha",
    "gpt-3.5-turbo-0125": "text-davinci-002-render-sha",
    "gpt-4": "gpt-4",
    "gpt-4-turbo": "gpt-4-turbo",
    "gpt-4o": "gpt-4o",
    "gpt-4o-mini": "gpt-4o-mini",
    "gpt-4o-canmore": "gpt-4o-canmore",
    "gpt-4.5o": "gpt-4.5o",
    "o1": "o1",
    "o1-mini": "o1-mini",
    "o1-preview": "o1-preview",
    "o1-pro": "o1-pro",
    "o3-mini": "o3-mini",
    "o3-mini-high": "o3-mini-high",
    "o3-mini-medium": "o3-mini-medium",
    "o3-mini-low": "o3-mini-low",
    "auto": "auto",
}


class ChatGPTClient:
    """ChatGPT backend API client using conversation endpoint."""

    CHATGPT_BASE_URL = "https://chatgpt.com"
    BACKEND_API = f"{CHATGPT_BASE_URL}/backend-api"

    def __init__(
        self,
        access_token: str,
        proxy_url: str | None = None,
        impersonate: str = "chrome131",
    ):
        self.access_token = access_token
        self.proxy_url = proxy_url
        self.impersonate = impersonate
        self._session: AsyncSession | None = None
        self._chat_token: str = "gAAAAAB"
        self._proof_token: str | None = None
        self._arkose_token: str | None = None
        self._turnstile_token: str | None = None
        self._device_id: str = str(uuid.uuid4())

    async def _get_session(self) -> AsyncSession:
        """Get or create curl_cffi session."""
        if self._session is None:
            proxies = None
            if self.proxy_url:
                proxies = {"http": self.proxy_url, "https": self.proxy_url}
            self._session = AsyncSession(
                proxies=proxies,
                impersonate=self.impersonate,
                timeout=120,
                verify=True,
            )
        return self._session

    async def _fetch_dpl(self) -> bool:
        """Fetch DPL (data-build) parameter from ChatGPT homepage.

        This is required for POW verification.
        """
        global _cached_dpl, _cached_scripts, _cached_time

        # Use cached value if fresh (< 15 minutes)
        if int(time.time()) - _cached_time < 15 * 60 and _cached_dpl:
            return True

        if settings.conversation_only:
            return True

        session = await self._get_session()
        headers = self._get_base_headers()
        # Remove auth header for homepage request
        headers.pop("authorization", None)

        try:
            response = await session.get(
                self.CHATGPT_BASE_URL,
                headers=headers,
                timeout=10,
            )
            response.raise_for_status()

            html = response.text

            # Extract script URLs
            _cached_scripts = re.findall(r'<script[^>]+src="([^"]+)"', html)

            # Extract DPL from script URL or data-build attribute
            for script in _cached_scripts:
                match = re.search(r"c/([^/]+)/_", script)
                if match:
                    _cached_dpl = match.group(1)
                    break

            # Fallback: look for data-build in HTML
            if not _cached_dpl:
                match = re.search(r'data-build="([^"]+)"', html)
                if match:
                    _cached_dpl = match.group(1)

            if _cached_dpl:
                _cached_time = int(time.time())
                logger.info(f"Got DPL: {_cached_dpl}")
                return True
            else:
                logger.warning("Failed to extract DPL from homepage")
                return False

        except Exception as e:
            logger.error(f"Error fetching DPL: {e}")
            return False

    def _get_base_headers(self) -> dict:
        """Get base headers for ChatGPT API requests."""
        return {
            "accept": "*/*",
            "accept-encoding": "gzip, deflate, br, zstd",
            "accept-language": "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7",
            "content-type": "application/json",
            "oai-device-id": self._device_id,
            "oai-language": "zh-CN",
            "origin": self.CHATGPT_BASE_URL,
            "priority": "u=1, i",
            "referer": f"{self.CHATGPT_BASE_URL}/",
            "sec-ch-ua": '"Google Chrome";v="131", "Chromium";v="131", "Not_A Brand";v="24"',
            "sec-ch-ua-mobile": "?0",
            "sec-ch-ua-platform": '"macOS"',
            "sec-fetch-dest": "empty",
            "sec-fetch-mode": "cors",
            "sec-fetch-site": "same-origin",
            "authorization": f"Bearer {self.access_token}",
            "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36",
        }

    async def get_chat_requirements(self) -> dict:
        """Get chat requirements token from ChatGPT sentinel endpoint.

        Returns:
            dict with token and other requirements
        """
        # Skip sentinel if conversation_only mode
        if settings.conversation_only:
            logger.info("Skipping sentinel verification (conversation_only mode)")
            return {}

        # First fetch DPL from homepage
        await self._fetch_dpl()

        session = await self._get_session()
        url = f"{self.BACKEND_API}/sentinel/chat-requirements"

        headers = self._get_base_headers()

        # Generate requirements token (simplified, may need POW for some accounts)
        config = self._get_pow_config()
        requirements_token = self._get_requirements_token(config)

        try:
            response = await session.post(
                url,
                headers=headers,
                json={"p": requirements_token},
                timeout=10,
            )

            if response.status_code == 200:
                data = response.json()
                self._chat_token = data.get("token", "gAAAAAB")

                # Check persona (account type)
                persona = data.get("persona", "")
                logger.info(f"ChatGPT persona: {persona}")

                # Check for Arkose requirement (required for free accounts)
                arkose = data.get("arkose", {})
                if arkose.get("required"):
                    arkose_dx = arkose.get("dx", "")
                    if settings.arkose_token_url:
                        self._arkose_token = await self._get_arkose_token(arkose_dx, persona)
                    else:
                        logger.warning("Arkose token required but no service URL configured")
                        logger.warning("Free accounts need Arkose verification. Set ARKOSE_TOKEN_URL or use a Plus account")

                # Check for Turnstile requirement
                turnstile = data.get("turnstile", {})
                if turnstile.get("required"):
                    logger.warning("Turnstile (CAPTCHA) required - may need manual intervention")

                # Check for POW requirement
                proofofwork = data.get("proofofwork", {})
                if proofofwork.get("required"):
                    diff = proofofwork.get("difficulty", 0)
                    seed = proofofwork.get("seed", "")
                    # Ensure diff is int
                    if isinstance(diff, str):
                        try:
                            diff = int(diff, 16) if diff.startswith("0x") else int(diff)
                        except ValueError:
                            diff = 0

                    # POW difficulty 0 means ChatGPT is very suspicious
                    if diff < settings.pow_difficulty:
                        logger.warning(f"POW difficulty {diff} too high (threshold: {settings.pow_difficulty})")
                        logger.warning("This usually means ChatGPT detected suspicious activity")
                        logger.warning("Try: 1) Use a clean proxy, 2) Wait a few minutes, 3) Use a Plus account")
                    else:
                        self._proof_token = await self._solve_pow(seed, diff, config)

                return data
            else:
                logger.warning(f"Chat requirements failed: {response.status_code}")
                return {}
        except Exception as e:
            logger.error(f"Error getting chat requirements: {e}")
            return {}

    def _get_pow_config(self) -> list:
        """Get proof-of-work config."""
        global _cached_dpl, _cached_scripts

        # Get current time in Eastern timezone
        now = datetime.now(timezone(timedelta(hours=-5)))
        time_str = now.strftime("%a %b %d %Y %H:%M:%S") + " GMT-0500 (Eastern Standard Time)"

        cores = [8, 16, 24, 32]
        return [
            random.choice([1920 + 1080, 2560 + 1440, 1920 + 1200, 2560 + 1600]),
            time_str,
            4294705152,
            0,
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36",
            random.choice(_cached_scripts) if _cached_scripts else "",
            _cached_dpl,
            "en-US",
            "en-US,es-US,en,es",
            0,
            "webdriver-false",
            "location",
            "window",
            time.perf_counter() * 1000,
            str(uuid.uuid4()),
            "",
            random.choice(cores),
            time.time() * 1000 - (time.perf_counter() * 1000),
        ]

    def _get_requirements_token(self, config: list) -> str:
        """Generate requirements token."""
        answer, _ = self._generate_answer(format(random.random()), "0fffff", config)
        return "gAAAAAC" + answer

    async def _get_arkose_token(self, blob: str, persona: str) -> str:
        """Get Arkose token from external service.

        Args:
            blob: The Arkose challenge blob
            persona: Account type (chatgpt-freeaccount, chatgpt-paid, etc.)

        Returns:
            Arkose token string
        """
        if not settings.arkose_token_url:
            return ""

        # Determine method based on account type
        method = "chat35" if persona == "chatgpt-freeaccount" else "chat4"

        session = await self._get_session()
        try:
            response = await session.post(
                settings.arkose_token_url,
                json={"blob": blob, "method": method},
                timeout=15,
            )

            if response.status_code == 200:
                data = response.json()
                if data.get("solved", True):
                    token = data.get("token", "")
                    logger.info(f"Got Arkose token: {token[:20]}...")
                    return token

            logger.warning(f"Arkose service failed: {response.status_code}")
            return ""

        except Exception as e:
            logger.error(f"Error getting Arkose token: {e}")
            return ""

    async def _solve_pow(self, seed: str, difficulty: int, config: list) -> str:
        """Solve proof-of-work challenge.

        Args:
            seed: The seed from server
            difficulty: Difficulty level (lower = harder)
            config: POW config

        Returns:
            Proof token
        """
        # Skip if difficulty too high
        if difficulty < 2:
            logger.warning(f"POW difficulty too high: {difficulty}")
            return ""

        diff_hex = format(difficulty, "06x")
        answer, solved = self._generate_answer(seed, diff_hex, config)

        if solved:
            return "gAAAAAB" + answer
        return ""

    def _generate_answer(self, seed: str, diff: str, config: list) -> tuple[str, bool]:
        """Generate POW answer.

        Args:
            seed: Challenge seed
            diff: Difficulty in hex
            config: POW config

        Returns:
            (answer, solved) tuple
        """
        diff_len = len(diff)
        seed_encoded = seed.encode()

        static_part1 = (
            json.dumps(config[:3], separators=(",", ":"), ensure_ascii=False)[:-1] + ","
        ).encode()
        static_part2 = (
            "," + json.dumps(config[4:9], separators=(",", ":"), ensure_ascii=False)[1:-1] + ","
        ).encode()
        static_part3 = (
            "," + json.dumps(config[10:], separators=(",", ":"), ensure_ascii=False)[1:]
        ).encode()

        target_diff = bytes.fromhex(diff)

        for i in range(500000):
            dynamic_i = str(i).encode()
            dynamic_j = str(i >> 1).encode()
            final_bytes = static_part1 + dynamic_i + static_part2 + dynamic_j + static_part3
            base_encode = pybase64.b64encode(final_bytes)
            hash_value = hashlib.sha3_512(seed_encoded + base_encode).digest()

            if hash_value[:diff_len] <= target_diff:
                return base_encode.decode(), True

        return "wQ8Lk5FbGpA2NcR9dShT6gYjU7VxZ4D" + pybase64.b64encode(f'"{seed}"'.encode()).decode(), False

    def _convert_messages_to_chatgpt(self, messages: list[dict]) -> list[dict]:
        """Convert OpenAI messages format to ChatGPT conversation format.

        Args:
            messages: OpenAI format messages [{"role": "user", "content": "..."}]

        Returns:
            ChatGPT format messages
        """
        chat_messages = []
        for msg in messages:
            role = msg.get("role", "user")
            content = msg.get("content", "")

            # Handle string content
            if isinstance(content, str):
                parts = [content]
                content_type = "text"
            # Handle list content (multimodal)
            elif isinstance(content, list):
                parts = []
                for item in content:
                    if item.get("type") == "text":
                        parts.append(item.get("text", ""))
                    elif item.get("type") == "image_url":
                        # For now, just note the image - would need upload for full support
                        parts.append(f"[Image: {item.get('image_url', {}).get('url', '')}]")
                content_type = "multimodal_text" if len(parts) > 1 else "text"
            else:
                parts = [str(content)]
                content_type = "text"

            chat_msg = {
                "id": str(uuid.uuid4()),
                "author": {"role": role},
                "content": {"content_type": content_type, "parts": parts},
                "metadata": {},
            }
            chat_messages.append(chat_msg)

        return chat_messages

    def _map_model(self, model: str) -> str:
        """Map OpenAI model name to ChatGPT internal model name."""
        return MODEL_MAP.get(model, model)

    async def chat_completions(
        self,
        model: str,
        messages: list[dict],
        stream: bool = False,
        conversation_id: str | None = None,
        parent_message_id: str | None = None,
        **kwargs,
    ) -> AsyncIterator[bytes] | dict:
        """Send chat completion request to ChatGPT conversation endpoint.

        Args:
            model: Model name (will be mapped to ChatGPT internal name)
            messages: OpenAI format messages
            stream: Whether to stream response
            conversation_id: Existing conversation ID (for continuing chats)
            parent_message_id: Parent message ID (for continuing chats)
            **kwargs: Additional parameters (max_tokens, temperature, etc.)

        Returns:
            AsyncIterator for streaming or dict for non-streaming
        """
        session = await self._get_session()
        url = f"{self.BACKEND_API}/conversation"

        # Get chat requirements
        await self.get_chat_requirements()

        # Build headers
        headers = self._get_base_headers()
        headers["accept"] = "text/event-stream"

        # Add sentinel headers only if not in conversation_only mode
        if not settings.conversation_only:
            headers["openai-sentinel-chat-requirements-token"] = self._chat_token

            if self._proof_token:
                headers["openai-sentinel-proof-token"] = self._proof_token
            if self._arkose_token:
                headers["openai-sentinel-arkose-token"] = self._arkose_token
            if self._turnstile_token:
                headers["openai-sentinel-turnstile-token"] = self._turnstile_token

        # Convert messages
        chat_messages = self._convert_messages_to_chatgpt(messages)

        # Map model
        internal_model = self._map_model(model)

        # Build request body
        request_body = {
            "action": "next",
            "client_contextual_info": {
                "is_dark_mode": False,
                "time_since_loaded": random.randint(50, 500),
                "page_height": random.randint(500, 1000),
                "page_width": random.randint(1000, 2000),
                "pixel_ratio": 1.5,
                "screen_height": random.randint(800, 1200),
                "screen_width": random.randint(1200, 2200),
            },
            "conversation_mode": {"kind": "primary_assistant"},
            "conversation_origin": None,
            "force_paragen": False,
            "force_paragen_model_slug": "",
            "force_rate_limit": False,
            "force_use_sse": True,
            "history_and_training_disabled": True,
            "messages": chat_messages,
            "model": internal_model,
            "paragen_cot_summary_display_override": "allow",
            "paragen_stream_type_override": None,
            "parent_message_id": parent_message_id or str(uuid.uuid4()),
            "reset_rate_limits": False,
            "suggestions": [],
            "supported_encodings": [],
            "system_hints": [],
            "timezone": "America/Los_Angeles",
            "timezone_offset_min": -480,
            "variant_purpose": "comparison_implicit",
            "websocket_request_id": str(uuid.uuid4()),
        }

        if conversation_id:
            request_body["conversation_id"] = conversation_id

        logger.info(f"Sending conversation request with model: {internal_model}")

        try:
            # Always use streaming for ChatGPT API
            response = await session.post(
                url,
                headers=headers,
                json=request_body,
                timeout=120,
                stream=True,  # Always stream for ChatGPT API
            )

            if response.status_code != 200:
                error_text = response.text[:500]
                logger.error(f"ChatGPT API error: {response.status_code} - {error_text}")

                # Check for specific error types
                if "Unusual activity" in error_text:
                    raise Exception(
                        "ChatGPT detected unusual activity. Try:\n"
                        "1. Use a different IP/proxy\n"
                        "2. Set CONVERSATION_ONLY=true (may still fail)\n"
                        "3. Wait a few minutes and retry"
                    )
                raise Exception(f"ChatGPT API error: {response.status_code}")

            if stream:
                return self._stream_response(response, model)
            else:
                return await self._collect_response(response, model)

        except Exception as e:
            logger.error(f"Error in chat_completions: {e}")
            raise

    async def _stream_response(self, response, model: str) -> AsyncIterator[bytes]:
        """Stream ChatGPT response and convert to OpenAI format."""
        chat_id = f"chatcmpl-{''.join(random.choices('abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789', k=29))}"
        created_time = int(time.time())

        # Send initial chunk with role
        initial_chunk = {
            "id": chat_id,
            "object": "chat.completion.chunk",
            "created": created_time,
            "model": model,
            "choices": [
                {
                    "index": 0,
                    "delta": {"role": "assistant", "content": ""},
                    "logprobs": None,
                    "finish_reason": None,
                }
            ],
        }
        yield f"data: {json.dumps(initial_chunk)}\n\n".encode()

        content_parts = []
        finish_reason = None

        async for line in response.aiter_lines():
            if not line or not line.startswith("data: "):
                continue

            if line == "data: [DONE]":
                break

            try:
                data = json.loads(line[6:])
                message = data.get("message", {})

                if not message:
                    continue

                role = message.get("author", {}).get("role")
                if role in ("user", "system"):
                    continue

                status = message.get("status")
                content = message.get("content", {})
                content_type = content.get("content_type", "")

                if content_type == "text":
                    parts = content.get("parts", [])
                    if parts:
                        new_text = parts[0]
                        if len(new_text) > len("".join(content_parts)):
                            delta_text = new_text[len("".join(content_parts)):]
                            content_parts.append(delta_text)

                            chunk = {
                                "id": chat_id,
                                "object": "chat.completion.chunk",
                                "created": created_time,
                                "model": model,
                                "choices": [
                                    {
                                        "index": 0,
                                        "delta": {"content": delta_text},
                                        "logprobs": None,
                                        "finish_reason": None,
                                    }
                                ],
                            }
                            yield f"data: {json.dumps(chunk)}\n\n".encode()

                if status == "finished_successfully":
                    finish_reason = "stop"

            except json.JSONDecodeError:
                continue

        # Send final chunk
        final_chunk = {
            "id": chat_id,
            "object": "chat.completion.chunk",
            "created": created_time,
            "model": model,
            "choices": [
                {
                    "index": 0,
                    "delta": {},
                    "logprobs": None,
                    "finish_reason": finish_reason,
                }
            ],
        }
        yield f"data: {json.dumps(final_chunk)}\n\n".encode()
        yield b"data: [DONE]\n\n"

    async def _collect_response(self, response, model: str) -> dict:
        """Collect streaming response and return as single dict."""
        chat_id = f"chatcmpl-{''.join(random.choices('abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789', k=29))}"
        created_time = int(time.time())

        all_content = []
        finish_reason = None

        async for line in response.aiter_lines():
            if not line or not line.startswith("data: "):
                continue

            if line == "data: [DONE]":
                break

            try:
                data = json.loads(line[6:])
                message = data.get("message", {})

                if not message:
                    continue

                role = message.get("author", {}).get("role")
                if role in ("user", "system"):
                    continue

                status = message.get("status")
                content = message.get("content", {})

                if content.get("content_type") == "text":
                    parts = content.get("parts", [])
                    if parts:
                        all_content = parts

                if status == "finished_successfully":
                    finish_reason = "stop"

            except json.JSONDecodeError:
                continue

        full_content = "".join(all_content)

        return {
            "id": chat_id,
            "object": "chat.completion",
            "created": created_time,
            "model": model,
            "choices": [
                {
                    "index": 0,
                    "message": {
                        "role": "assistant",
                        "content": full_content,
                    },
                    "logprobs": None,
                    "finish_reason": finish_reason or "stop",
                }
            ],
            "usage": {
                "prompt_tokens": 0,
                "completion_tokens": 0,
                "total_tokens": 0,
            },
        }

    async def close(self):
        """Close the session."""
        if self._session:
            await self._session.close()
            self._session = None
````

## File: src/gpt_proxy/main.py
````python
"""FastAPI application factory."""

from contextlib import asynccontextmanager
from fastapi import FastAPI
from fastapi.responses import JSONResponse

from gpt_proxy import __version__
from gpt_proxy.config import settings
from gpt_proxy.services.auth_manager import close_auth_manager
from gpt_proxy.services.browser_auth import close_browser_auth
from gpt_proxy.api.router import router as api_router
from gpt_proxy.api.auth_router import router as auth_router
from gpt_proxy.middleware.rate_limit import RateLimitMiddleware
from gpt_proxy.middleware.error_handler import proxy_error_handler, generic_error_handler, ProxyError
from gpt_proxy.middleware.logging import LoggingMiddleware


@asynccontextmanager
async def lifespan(app: FastAPI):
    """Application lifespan handler."""
    # Startup
    yield
    # Shutdown
    await close_auth_manager()
    await close_browser_auth()


def create_app() -> FastAPI:
    """Create and configure the FastAPI application."""
    app = FastAPI(
        title="GPT Reverse Proxy",
        description="""
## ChatGPT Reverse Proxy

Use ChatGPT without API keys! Login with your ChatGPT account session token.

### Quick Start

1. **Get session token**: Login to chat.openai.com, open DevTools, copy `__Secure-next-auth.session-token` cookie
2. **Login**: POST to `/auth/login` with your session token
3. **Use API**: Use the returned `session_id` as Bearer token

### Example

```bash
# Login
curl -X POST http://localhost:8000/auth/login \\
  -H "Content-Type: application/json" \\
  -d '{"session_token": "your-token-here"}'

# Chat
curl -X POST http://localhost:8000/v1/chat/completions \\
  -H "Authorization: Bearer <session_id>" \\
  -H "Content-Type: application/json" \\
  -d '{"model": "gpt-4", "messages": [{"role": "user", "content": "Hello!"}]}'
```
""",
        version=__version__,
        lifespan=lifespan,
        docs_url="/docs",
        redoc_url="/redoc",
    )

    # Health check endpoints
    @app.get("/health")
    async def health_check():
        return {"status": "healthy", "version": __version__}

    @app.get("/ready")
    async def readiness_check():
        return {"status": "ready"}

    # Error handlers
    app.add_exception_handler(ProxyError, proxy_error_handler)
    app.add_exception_handler(Exception, generic_error_handler)

    # Logging middleware
    app.add_middleware(LoggingMiddleware)

    # Rate limiting middleware
    if settings.rate_limit_enabled:
        app.add_middleware(RateLimitMiddleware, requests_per_minute=settings.rate_limit_requests_per_minute)

    # Include routers
    app.include_router(auth_router)
    app.include_router(api_router)

    return app


# Default app instance
app = create_app()
````

## File: src/gpt_proxy/middleware/__init__.py
````python
"""Middleware components for GPT Proxy."""

from gpt_proxy.middleware.auth import AuthMiddleware
from gpt_proxy.middleware.rate_limit import RateLimitMiddleware
from gpt_proxy.middleware.logging import LoggingMiddleware
from gpt_proxy.middleware.error_handler import (
    ProxyError,
    UpstreamError,
    RateLimitError,
    proxy_error_handler,
    generic_error_handler,
)

__all__ = [
    "AuthMiddleware",
    "RateLimitMiddleware",
    "LoggingMiddleware",
    "ProxyError",
    "UpstreamError",
    "RateLimitError",
    "proxy_error_handler",
    "generic_error_handler",
]
````

## File: src/gpt_proxy/middleware/auth.py
````python
"""Authentication middleware."""

from starlette.middleware.base import BaseHTTPMiddleware
from fastapi import Request
from fastapi.responses import JSONResponse


class AuthMiddleware(BaseHTTPMiddleware):
    """Validate API keys against configured keys."""

    SKIP_PATHS = {"/health", "/ready", "/docs", "/redoc", "/openapi.json"}

    async def dispatch(self, request: Request, call_next):
        # Skip auth for health and docs endpoints
        if request.url.path in self.SKIP_PATHS:
            return await call_next(request)

        # Check Authorization header
        auth_header = request.headers.get("Authorization", "")

        if not auth_header:
            return JSONResponse(
                status_code=401,
                content={
                    "error": {
                        "message": "Missing Authorization header",
                        "type": "invalid_request_error",
                        "code": "invalid_api_key",
                    }
                },
            )

        if not auth_header.startswith("Bearer "):
            return JSONResponse(
                status_code=401,
                content={
                    "error": {
                        "message": "Invalid Authorization header format. Expected 'Bearer <token>'",
                        "type": "invalid_request_error",
                        "code": "invalid_api_key",
                    }
                },
            )

        # Accept any valid Bearer token - the proxy manages actual OpenAI keys
        return await call_next(request)
````

## File: src/gpt_proxy/middleware/error_handler.py
````python
"""Global error handling."""

from fastapi import Request
from fastapi.responses import JSONResponse
import logging

logger = logging.getLogger("gpt_proxy")


class ProxyError(Exception):
    """Custom proxy error."""

    def __init__(self, message: str, code: str, status_code: int = 500, error_type: str = "proxy_error"):
        self.message = message
        self.code = code
        self.status_code = status_code
        self.error_type = error_type
        super().__init__(message)


class UpstreamError(ProxyError):
    """Error from upstream OpenAI API."""

    def __init__(self, message: str, status_code: int = 502):
        super().__init__(
            message=message,
            code="upstream_error",
            status_code=status_code,
            error_type="upstream_error",
        )


class RateLimitError(ProxyError):
    """Rate limit exceeded."""

    def __init__(self, message: str = "Rate limit exceeded"):
        super().__init__(
            message=message,
            code="rate_limit_exceeded",
            status_code=429,
            error_type="rate_limit_error",
        )


async def proxy_error_handler(request: Request, exc: ProxyError) -> JSONResponse:
    """Handle ProxyError exceptions."""
    logger.error(f"Proxy error: {exc.message} (code={exc.code})")

    return JSONResponse(
        status_code=exc.status_code,
        content={
            "error": {
                "message": exc.message,
                "type": exc.error_type,
                "code": exc.code,
            }
        },
    )


async def generic_error_handler(request: Request, exc: Exception) -> JSONResponse:
    """Handle generic exceptions."""
    from fastapi import HTTPException

    if isinstance(exc, HTTPException):
        return JSONResponse(
            status_code=exc.status_code,
            content={"detail": exc.detail},
        )

    logger.exception(f"Unexpected error: {exc}")

    return JSONResponse(
        status_code=500,
        content={
            "error": {
                "message": "Internal server error",
                "type": "internal_error",
                "code": "internal_error",
            }
        },
    )
````

## File: src/gpt_proxy/middleware/logging.py
````python
"""Request/response logging middleware."""

from starlette.middleware.base import BaseHTTPMiddleware
from fastapi import Request
import json
import time
import uuid
import logging

logger = logging.getLogger("gpt_proxy")


class LoggingMiddleware(BaseHTTPMiddleware):
    """Log all requests and responses."""

    SKIP_PATHS = {"/health", "/ready", "/docs", "/redoc", "/openapi.json"}
    MAX_BODY_LOG = 1000  # Max characters to log for body

    async def dispatch(self, request: Request, call_next):
        # Skip logging for health endpoints
        if request.url.path in self.SKIP_PATHS:
            return await call_next(request)

        start_time = time.time()
        request_id = request.headers.get("X-Request-ID", str(uuid.uuid4()))

        # Log request
        log_data = {
            "event": "request",
            "request_id": request_id,
            "method": request.method,
            "path": request.url.path,
            "query": str(request.query_params),
            "client": request.client.host if request.client else None,
        }
        logger.info(json.dumps(log_data))

        # Process request
        response = await call_next(request)

        # Log response
        duration = time.time() - start_time
        log_data = {
            "event": "response",
            "request_id": request_id,
            "status_code": response.status_code,
            "duration_ms": round(duration * 1000, 2),
        }
        logger.info(json.dumps(log_data))

        # Add request ID to response headers
        response.headers["X-Request-ID"] = request_id

        return response
````

## File: src/gpt_proxy/middleware/rate_limit.py
````python
"""Rate limiting middleware."""

from starlette.middleware.base import BaseHTTPMiddleware
from fastapi import Request
from fastapi.responses import JSONResponse
import time
from collections import defaultdict


class RateLimitMiddleware(BaseHTTPMiddleware):
    """Rate limiting with configurable limits."""

    SKIP_PATHS = {"/health", "/ready", "/docs", "/redoc", "/openapi.json"}

    def __init__(self, app, requests_per_minute: int = 60):
        super().__init__(app)
        self.rpm = requests_per_minute
        self.requests: dict[str, list[float]] = defaultdict(list)

    async def dispatch(self, request: Request, call_next):
        # Skip rate limit for health endpoints
        if request.url.path in self.SKIP_PATHS:
            return await call_next(request)

        # Get client identifier
        client_ip = request.client.host if request.client else "unknown"
        client_id = request.headers.get("X-Client-ID", client_ip)

        # Check rate limit
        now = time.time()
        timestamps = self.requests[client_id]

        # Remove old timestamps (older than 1 minute)
        timestamps[:] = [t for t in timestamps if now - t < 60]

        if len(timestamps) >= self.rpm:
            # Calculate retry-after
            oldest = min(timestamps)
            retry_after = int(60 - (now - oldest)) + 1

            return JSONResponse(
                status_code=429,
                content={
                    "error": {
                        "message": f"Rate limit exceeded. Max {self.rpm} requests per minute.",
                        "type": "rate_limit_error",
                        "code": "rate_limit_exceeded",
                    }
                },
                headers={"Retry-After": str(retry_after), "X-RateLimit-Limit": str(self.rpm)},
            )

        # Record request
        self.requests[client_id].append(now)

        # Process request
        response = await call_next(request)

        # Add rate limit headers to response
        remaining = self.rpm - len(self.requests[client_id])
        response.headers["X-RateLimit-Limit"] = str(self.rpm)
        response.headers["X-RateLimit-Remaining"] = str(remaining)

        return response
````

## File: src/gpt_proxy/services/__init__.py
````python
"""Business services for GPT Proxy."""

from gpt_proxy.services.key_manager import APIKeyManager, KeyState
from gpt_proxy.services.cache import ResponseCache
from gpt_proxy.services.cost_tracker import CostTracker, MODEL_PRICING
from gpt_proxy.services.health import HealthService

__all__ = [
    "APIKeyManager",
    "KeyState",
    "ResponseCache",
    "CostTracker",
    "MODEL_PRICING",
    "HealthService",
]
````

## File: src/gpt_proxy/services/auth_manager.py
````python
"""Authentication manager for ChatGPT session tokens."""

from dataclasses import dataclass, field
from datetime import datetime, timedelta, timezone
from typing import Optional, Literal
import base64
import httpx
import secrets
import asyncio
from logging import getLogger
import json
from pathlib import Path

from gpt_proxy.config import settings

logger = getLogger(__name__)


def utcnow() -> datetime:
    """Get current UTC time with timezone info."""
    return datetime.now(timezone.utc)


@dataclass
class UserSession:
    """User session with ChatGPT tokens."""
    session_id: str
    user_id: str
    email: str
    access_token: str
    session_token: str  # ChatGPT session token (from browser)
    expires_at: datetime
    created_at: datetime = field(default_factory=datetime.now)
    request_count: int = 0
    is_active: bool = True


class AuthManager:
    """Manage ChatGPT session-based authentication."""

    CHATGPT_API_BASE = "https://chatgpt.com"
    SESSION_API = f"{CHATGPT_API_BASE}/api/auth/session"
    BACKEND_API = f"{CHATGPT_API_BASE}/backend-api"
    SESSION_FILE = Path("./sessions.json")

    def __init__(self):
        self.sessions: dict[str, UserSession] = {}
        self._lock = asyncio.Lock()
        self._client: httpx.AsyncClient | None = None
        self._load_sessions()

    def _load_sessions(self):
        """Load sessions from disk."""
        if self.SESSION_FILE.exists():
            try:
                with open(self.SESSION_FILE) as f:
                    data = json.load(f)
                    for sid, sdata in data.items():
                        # Parse datetime, handling both naive and aware formats
                        expires_at_str = sdata["expires_at"]
                        created_at_str = sdata["created_at"]

                        try:
                            expires_at = datetime.fromisoformat(expires_at_str)
                            created_at = datetime.fromisoformat(created_at_str)
                        except ValueError:
                            # Fallback for old format
                            expires_at = datetime.fromisoformat(expires_at_str.replace("Z", "+00:00"))
                            created_at = datetime.fromisoformat(created_at_str.replace("Z", "+00:00"))

                        # Ensure timezone-aware
                        if expires_at.tzinfo is None:
                            expires_at = expires_at.replace(tzinfo=timezone.utc)
                        if created_at.tzinfo is None:
                            created_at = created_at.replace(tzinfo=timezone.utc)

                        sdata["expires_at"] = expires_at
                        sdata["created_at"] = created_at
                        # Decode base64-encoded sensitive fields
                        try:
                            sdata["access_token"] = base64.b64decode(sdata["access_token"]).decode()
                            sdata["session_token"] = base64.b64decode(sdata["session_token"]).decode()
                        except Exception:
                            # Legacy plain-text format, use as-is
                            pass
                        self.sessions[sid] = UserSession(**sdata)
                logger.info(f"Loaded {len(self.sessions)} sessions from disk")
            except Exception as e:
                logger.warning(f"Failed to load sessions: {e}")

    def _save_sessions(self):
        """Save sessions to disk."""
        try:
            data = {}
            for sid, session in self.sessions.items():
                if session.is_active:
                    data[sid] = {
                        "session_id": session.session_id,
                        "user_id": session.user_id,
                        "email": session.email,
                        "access_token": base64.b64encode(session.access_token.encode()).decode(),
                        "session_token": base64.b64encode(session.session_token.encode()).decode(),
                        "expires_at": session.expires_at.isoformat(),
                        "created_at": session.created_at.isoformat(),
                        "request_count": session.request_count,
                        "is_active": session.is_active,
                    }
            with open(self.SESSION_FILE, "w") as f:
                json.dump(data, f, indent=2)
        except Exception as e:
            logger.warning(f"Failed to save sessions: {e}")

    async def _get_client(self) -> httpx.AsyncClient:
        """Get or create HTTP client with proxy support."""
        if self._client is None or self._client.is_closed:
            client_kwargs = {
                "timeout": httpx.Timeout(settings.http_timeout, connect=settings.http_connect_timeout),
                "follow_redirects": True,
            }
            if settings.browser_proxy:
                client_kwargs["proxy"] = settings.browser_proxy
                logger.info(f"AuthManager using proxy: {settings.browser_proxy}")
            self._client = httpx.AsyncClient(**client_kwargs)
        return self._client

    def _is_cloudflare_challenge(self, response: httpx.Response) -> bool:
        """Detect if response is a Cloudflare challenge page."""
        if response.status_code in [403, 503]:
            body = response.text or ""
            indicators = [
                "cloudflare",
                "cf-browser-verification",
                "challenge-platform",
                "Just a moment...",
                "Checking your browser",
            ]
            return any(indicator.lower() in body.lower() for indicator in indicators)
        return False

    async def exchange_session_token(self, session_token: str) -> Optional[UserSession]:
        """Exchange ChatGPT session token for access token.

        Args:
            session_token: The __Secure-next-auth.session-token from browser

        Returns:
            UserSession if successful, None otherwise
        """
        client = await self._get_client()

        try:
            response = await client.get(
                self.SESSION_API,
                headers={
                    "Cookie": f"__Secure-next-auth.session-token={session_token}",
                    "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
                    "Accept": "application/json",
                },
            )

            logger.debug(f"Session API response: status={response.status_code}")

            # Check for Cloudflare challenge
            if self._is_cloudflare_challenge(response):
                logger.error("Cloudflare challenge detected. Session may need browser refresh.")
                return None

            if response.status_code != 200:
                body_preview = response.text[:500] if response.text else "<empty>"
                logger.warning(f"Session token exchange failed: status={response.status_code}, body={body_preview}")
                return None

            # Check content type before parsing
            content_type = response.headers.get("content-type", "")
            if "application/json" not in content_type:
                body_preview = response.text[:500] if response.text else "<empty>"
                logger.error(f"Unexpected content type: {content_type}, body={body_preview}")
                return None

            # Safely parse JSON
            try:
                data = response.json()
            except json.JSONDecodeError as json_error:
                body_preview = response.text[:500] if response.text else "<empty>"
                logger.error(f"JSON parse error: {json_error}, body={body_preview}")
                return None

            if not data.get("accessToken"):
                logger.warning("No accessToken in response")
                return None

            session_id = secrets.token_urlsafe(32)
            user = data.get("user", {})

            # Parse expiry time
            expires_str = data.get("expires", "")
            try:
                expires_at = datetime.fromisoformat(expires_str.replace("Z", "+00:00"))
            except Exception:
                expires_at = utcnow() + timedelta(hours=1)

            session = UserSession(
                session_id=session_id,
                user_id=user.get("id", "unknown"),
                email=user.get("email", "unknown"),
                access_token=data["accessToken"],
                session_token=session_token,
                expires_at=expires_at,
            )

            logger.info(f"Created session for user: {session.email}")
            return session

        except httpx.TimeoutException as e:
            logger.error(f"Timeout exchanging session token: {e}")
            return None
        except Exception as e:
            logger.error(f"Error exchanging session token: {e}")
            return None

    async def validate_session_token(self, session_token: str) -> dict:
        """Validate session token and return basic info without full exchange.

        Returns:
            dict with keys: valid (bool), error (str|None), user_email (str|None)
        """
        client = await self._get_client()

        try:
            response = await client.get(
                self.SESSION_API,
                headers={
                    "Cookie": f"__Secure-next-auth.session-token={session_token}",
                    "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36",
                },
            )

            if self._is_cloudflare_challenge(response):
                return {"valid": False, "error": "cloudflare_challenge"}

            if response.status_code == 401:
                return {"valid": False, "error": "invalid_token"}

            if response.status_code != 200:
                return {"valid": False, "error": f"http_{response.status_code}"}

            try:
                data = response.json()
                if data.get("user"):
                    return {
                        "valid": True,
                        "error": None,
                        "user_email": data["user"].get("email"),
                    }
            except json.JSONDecodeError:
                return {"valid": False, "error": "invalid_response"}

            return {"valid": False, "error": "no_user_data"}

        except httpx.TimeoutException:
            return {"valid": False, "error": "timeout"}
        except Exception as e:
            logger.error(f"Token validation error: {e}")
            return {"valid": False, "error": str(e)}

    async def refresh_session(self, session_id: str) -> bool:
        """Refresh expired access token using session token.

        Args:
            session_id: The session ID to refresh

        Returns:
            True if refreshed successfully
        """
        async with self._lock:
            session = self.sessions.get(session_id)
            if not session:
                return False

            # Re-exchange session token
            new_session = await self.exchange_session_token(session.session_token)
            if new_session:
                new_session.session_id = session_id  # Keep same ID
                new_session.request_count = session.request_count
                self.sessions[session_id] = new_session
                logger.info(f"Refreshed session for: {session.email}")
                return True

            return False

    def create_session(self, session: UserSession) -> str:
        """Store session and return ID."""
        self.sessions[session.session_id] = session
        self._save_sessions()
        return session.session_id

    def get_session(self, session_id: str) -> Optional[UserSession]:
        """Get session by ID."""
        return self.sessions.get(session_id)

    async def get_valid_token(self, session_id: str) -> Optional[str]:
        """Get valid access token, refreshing if needed.

        Args:
            session_id: The session ID

        Returns:
            Valid access token or None
        """
        session = self.sessions.get(session_id)
        if not session or not session.is_active:
            return None

        # Check if token is expired or about to expire (5 min buffer)
        if utcnow() >= session.expires_at - timedelta(minutes=5):
            # Try to refresh
            if not await self.refresh_session(session_id):
                return None
            session = self.sessions.get(session_id)

        session.request_count += 1
        return session.access_token if session else None

    def invalidate_session(self, session_id: str) -> bool:
        """Invalidate a session."""
        if session_id in self.sessions:
            self.sessions[session_id].is_active = False
            return True
        return False

    def list_sessions(self) -> list[dict]:
        """List all active sessions (masked)."""
        return [
            {
                "session_id": s.session_id[:8] + "...",
                "email": s.email,
                "is_active": s.is_active,
                "request_count": s.request_count,
                "expires_at": s.expires_at.isoformat(),
            }
            for s in self.sessions.values()
        ]

    async def close(self):
        """Close HTTP client."""
        if self._client and not self._client.is_closed:
            await self._client.aclose()
            self._client = None


# Global auth manager instance
_auth_manager: AuthManager | None = None


def get_auth_manager() -> AuthManager:
    """Get the global auth manager instance."""
    global _auth_manager
    if _auth_manager is None:
        _auth_manager = AuthManager()
    return _auth_manager


async def close_auth_manager():
    """Close the global auth manager."""
    global _auth_manager
    if _auth_manager:
        await _auth_manager.close()
````

## File: src/gpt_proxy/services/browser_auth.py
````python
"""Browser-based authentication for ChatGPT using Playwright."""

from playwright.async_api import async_playwright, BrowserContext
from typing import Optional
from pathlib import Path
from logging import getLogger
import asyncio
import os

logger = getLogger(__name__)


class BrowserAuthManager:
    """Manage browser-based ChatGPT authentication with persistent profile."""

    CHATGPT_URL = "https://chatgpt.com/auth/login"
    SESSION_COOKIE_NAME = "__Secure-next-auth.session-token"

    def __init__(self, profile_dir: str = "./browser_profile", proxy: str = None):
        self.profile_dir = Path(profile_dir)
        self._playwright = None
        self._context: Optional[BrowserContext] = None
        # 代理设置，从环境变量或参数获取
        self.proxy = proxy or os.environ.get("HTTPS_PROXY") or os.environ.get("https_proxy")

    async def initialize(self, headless: bool = False):
        """Initialize browser context with persistent profile.

        Args:
            headless: If False, shows browser window for user interaction
        """
        self.profile_dir.mkdir(parents=True, exist_ok=True)
        self._playwright = await async_playwright().start()

        launch_options = {
            "user_data_dir": str(self.profile_dir),
            "headless": headless,
            "channel": "chrome",
            "viewport": {"width": 1280, "height": 800},
            "locale": "en-US",
            "args": [
                "--disable-blink-features=AutomationControlled",
                "--disable-features=IsolateOrigins,site-per-process",
                "--no-sandbox",
            ],
            "ignore_default_args": ["--enable-automation"],
        }

        # 添加代理支持
        if self.proxy:
            logger.info(f"Using proxy: {self.proxy}")
            launch_options["proxy"] = {"server": self.proxy}

        try:
            self._context = await self._playwright.chromium.launch_persistent_context(**launch_options)
        except Exception as e:
            logger.warning(f"Failed to launch with channel='chrome' ({e}); falling back to bundled Chromium")
            launch_options.pop("channel", None)
            self._context = await self._playwright.chromium.launch_persistent_context(**launch_options)

        # 抹掉 navigator.webdriver 等自动化指纹
        await self._context.add_init_script(
            """
            Object.defineProperty(navigator, 'webdriver', { get: () => undefined });
            Object.defineProperty(navigator, 'languages', { get: () => ['en-US', 'en'] });
            Object.defineProperty(navigator, 'plugins', { get: () => [1, 2, 3, 4, 5] });
            window.chrome = window.chrome || { runtime: {} };
            const originalQuery = window.navigator.permissions && window.navigator.permissions.query;
            if (originalQuery) {
                window.navigator.permissions.query = (parameters) => (
                    parameters.name === 'notifications'
                        ? Promise.resolve({ state: Notification.permission })
                        : originalQuery(parameters)
                );
            }
            """
        )
        logger.info(f"Browser initialized with profile: {self.profile_dir}")

    async def get_session_token(
        self,
        wait_for_login: bool = True,
        timeout: int = 300
    ) -> Optional[str]:
        """Get session token from browser.

        Args:
            wait_for_login: Wait for user to complete login
            timeout: Maximum seconds to wait for login

        Returns:
            Session token or None
        """
        if not self._context:
            await self.initialize(headless=False)

        page = await self._context.new_page()

        try:
            logger.info("Navigating to ChatGPT...")
            await page.goto(self.CHATGPT_URL, wait_until="domcontentloaded", timeout=60000)

            # Check if already logged in
            cookies = await self._context.cookies()
            for cookie in cookies:
                if cookie["name"] == self.SESSION_COOKIE_NAME:
                    logger.info("Found existing session token")
                    return cookie["value"]

            if wait_for_login:
                logger.info(f"Waiting for user to login (timeout: {timeout}s)...")
                # 轮询 cookie，避免依赖 URL 模式（不同域名/路径都可能出现）
                deadline = asyncio.get_event_loop().time() + timeout
                while asyncio.get_event_loop().time() < deadline:
                    cookies = await self._context.cookies()
                    for cookie in cookies:
                        if cookie["name"] == self.SESSION_COOKIE_NAME and cookie.get("value"):
                            logger.info("Successfully extracted session token")
                            return cookie["value"]
                    await asyncio.sleep(2)

                logger.warning("Login timeout: session cookie not found")
                return None

            return None

        except Exception as e:
            logger.error(f"Browser error: {e}")
            return None
        finally:
            await page.close()

    async def close(self):
        """Close browser and cleanup."""
        if self._context:
            await self._context.close()
            self._context = None
        if self._playwright:
            await self._playwright.stop()
            self._playwright = None
        logger.info("Browser closed")


# Singleton instance
_browser_auth: BrowserAuthManager | None = None


def get_browser_auth() -> BrowserAuthManager:
    """Get the global browser auth instance."""
    global _browser_auth
    from gpt_proxy.config import settings

    # 每次都重新创建，确保使用最新配置
    if _browser_auth is None:
        _browser_auth = BrowserAuthManager(
            profile_dir=settings.browser_profile_dir,
            proxy=settings.browser_proxy or None
        )

    # 确保代理设置正确
    if settings.browser_proxy and _browser_auth.proxy != settings.browser_proxy:
        _browser_auth.proxy = settings.browser_proxy
        logger.info(f"Updated proxy to: {settings.browser_proxy}")

    return _browser_auth


async def close_browser_auth():
    """Close the global browser auth instance."""
    global _browser_auth
    if _browser_auth:
        await _browser_auth.close()
        _browser_auth = None
````

## File: src/gpt_proxy/services/cache.py
````python
"""Response caching service."""

import hashlib
from datetime import datetime, timedelta


class ResponseCache:
    """Cache responses for identical requests."""

    def __init__(self, ttl_seconds: int = 3600, backend: str = "memory"):
        self.ttl = ttl_seconds
        self.backend = backend
        self._cache: dict[str, tuple[bytes, datetime]] = {}
        self._hits = 0
        self._misses = 0

    def generate_key(self, method: str, path: str, body: bytes | None) -> str:
        """Generate deterministic cache key."""
        content = f"{method}:{path}:{body.hex() if body else 'empty'}"
        return hashlib.sha256(content.encode()).hexdigest()

    async def get(self, cache_key: str) -> bytes | None:
        """Get cached response if not expired."""
        if cache_key not in self._cache:
            self._misses += 1
            return None

        response, timestamp = self._cache[cache_key]
        if datetime.now() - timestamp > timedelta(seconds=self.ttl):
            del self._cache[cache_key]
            self._misses += 1
            return None

        self._hits += 1
        return response

    async def set(self, cache_key: str, response: bytes):
        """Cache response."""
        self._cache[cache_key] = (response, datetime.now())

    async def clear(self):
        """Clear all cached responses."""
        self._cache.clear()
        self._hits = 0
        self._misses = 0

    def stats(self) -> dict:
        """Get cache statistics."""
        total = self._hits + self._misses
        hit_rate = self._hits / total if total > 0 else 0
        return {
            "size": len(self._cache),
            "ttl_seconds": self.ttl,
            "hits": self._hits,
            "misses": self._misses,
            "hit_rate": round(hit_rate, 4),
        }
````

## File: src/gpt_proxy/services/cost_tracker.py
````python
"""Token counting and cost tracking."""

import tiktoken
from datetime import datetime, timedelta
from typing import Literal


# Pricing per 1M tokens (as of 2024)
MODEL_PRICING = {
    "gpt-4o": {"input": 2.50, "output": 10.00},
    "gpt-4o-mini": {"input": 0.15, "output": 0.60},
    "gpt-4-turbo": {"input": 10.00, "output": 30.00},
    "gpt-4": {"input": 30.00, "output": 60.00},
    "gpt-3.5-turbo": {"input": 0.50, "output": 1.50},
    "gpt-3.5-turbo-0125": {"input": 0.50, "output": 1.50},
    "text-embedding-3-small": {"input": 0.02, "output": 0},
    "text-embedding-3-large": {"input": 0.13, "output": 0},
    "text-embedding-ada-002": {"input": 0.10, "output": 0},
    "dall-e-3": {"input": 0, "output": 0, "per_image": 0.04},
    "whisper-1": {"input": 0, "output": 0, "per_minute": 0.006},
    "tts-1": {"input": 0, "output": 0, "per_1k_chars": 0.015},
    "tts-1-hd": {"input": 0, "output": 0, "per_1k_chars": 0.030},
}


class CostTracker:
    """Track token usage and costs."""

    def __init__(self):
        self._usage: list[dict] = []

    def count_tokens(self, text: str, model: str) -> int:
        """Count tokens using tiktoken."""
        try:
            encoding = tiktoken.encoding_for_model(model)
        except KeyError:
            encoding = tiktoken.get_encoding("cl100k_base")
        return len(encoding.encode(text))

    def count_messages_tokens(self, messages: list[dict], model: str) -> int:
        """Count tokens in chat messages."""
        total = 0
        for msg in messages:
            total += 4  # Message overhead
            content = msg.get("content", "")
            if isinstance(content, str):
                total += self.count_tokens(content, model)
            elif isinstance(content, list):
                for part in content:
                    if isinstance(part, dict) and part.get("type") == "text":
                        total += self.count_tokens(part.get("text", ""), model)
        total += 2  # Reply priming
        return total

    async def track_usage(
        self,
        key_id: str,
        model: str,
        prompt_tokens: int,
        completion_tokens: int,
    ):
        """Record token usage."""
        pricing = MODEL_PRICING.get(model, {"input": 0, "output": 0})
        cost = (
            prompt_tokens * pricing["input"] / 1_000_000
            + completion_tokens * pricing["output"] / 1_000_000
        )

        self._usage.append({
            "timestamp": datetime.now().isoformat(),
            "key_id": key_id,
            "model": model,
            "prompt_tokens": prompt_tokens,
            "completion_tokens": completion_tokens,
            "total_tokens": prompt_tokens + completion_tokens,
            "cost_usd": cost,
        })

    async def get_usage(
        self,
        period: Literal["day", "week", "month"] = "day",
        key_id: str | None = None,
    ) -> dict:
        """Get usage summary."""
        now = datetime.now()
        if period == "day":
            cutoff = now - timedelta(days=1)
        elif period == "week":
            cutoff = now - timedelta(days=7)
        else:
            cutoff = now - timedelta(days=30)

        filtered = [
            u
            for u in self._usage
            if datetime.fromisoformat(u["timestamp"]) > cutoff
            and (key_id is None or u["key_id"] == key_id)
        ]

        total_cost = sum(u["cost_usd"] for u in filtered)
        total_prompt = sum(u["prompt_tokens"] for u in filtered)
        total_completion = sum(u["completion_tokens"] for u in filtered)

        return {
            "period": period,
            "total_requests": len(filtered),
            "total_tokens": total_prompt + total_completion,
            "prompt_tokens": total_prompt,
            "completion_tokens": total_completion,
            "total_cost_usd": round(total_cost, 4),
            "by_model": self._aggregate_by_model(filtered),
        }

    def _aggregate_by_model(self, usage: list[dict]) -> dict:
        """Aggregate usage by model."""
        result: dict[str, dict] = {}
        for u in usage:
            model = u["model"]
            if model not in result:
                result[model] = {"tokens": 0, "cost": 0, "requests": 0}
            result[model]["tokens"] += u["total_tokens"]
            result[model]["cost"] += u["cost_usd"]
            result[model]["requests"] += 1

        # Round costs
        for model in result:
            result[model]["cost"] = round(result[model]["cost"], 4)

        return result

    def clear(self):
        """Clear usage history."""
        self._usage.clear()
````

## File: src/gpt_proxy/services/health.py
````python
"""Health check service."""

from datetime import datetime


class HealthService:
    """Health check service."""

    def __init__(self, start_time: datetime | None = None):
        self.start_time = start_time or datetime.now()

    async def check(self) -> dict:
        """Basic health check."""
        uptime = (datetime.now() - self.start_time).total_seconds()
        return {
            "status": "healthy",
            "uptime_seconds": round(uptime, 2),
            "timestamp": datetime.now().isoformat(),
        }

    async def check_ready(self, key_manager=None, cache=None) -> dict:
        """Readiness check."""
        checks = {}

        # Check API keys
        if key_manager:
            active_keys = len([k for k in key_manager.keys if k.is_active])
            checks["api_keys"] = {"status": "ok" if active_keys > 0 else "error", "count": active_keys}
        else:
            checks["api_keys"] = {"status": "unknown"}

        # Check cache
        if cache:
            checks["cache"] = {"status": "ok", "stats": cache.stats()}
        else:
            checks["cache"] = {"status": "disabled"}

        # Overall status
        all_ok = all(c.get("status") in ("ok", "disabled", "unknown") for c in checks.values())

        return {
            "status": "ready" if all_ok else "not_ready",
            "checks": checks,
            "timestamp": datetime.now().isoformat(),
        }
````

## File: src/gpt_proxy/services/key_manager.py
````python
"""API key management with rotation strategies."""

from typing import Literal
from datetime import datetime, timedelta
from dataclasses import dataclass
import random


@dataclass
class KeyState:
    """State of an individual API key."""

    key: str
    is_active: bool = True
    exhausted_at: datetime | None = None
    reset_at: datetime | None = None
    request_count: int = 0
    error_count: int = 0


class APIKeyManager:
    """Manage multiple OpenAI API keys with rotation strategies."""

    def __init__(
        self,
        keys: list[str],
        strategy: Literal["round-robin", "least-used", "random"] = "round-robin",
    ):
        self.keys = [KeyState(key=k) for k in keys]
        self.strategy = strategy
        self._index = 0

    def get_key(self) -> str | None:
        """Get next available key based on strategy."""
        active_keys = [k for k in self.keys if k.is_active and not self._is_exhausted(k)]

        if not active_keys:
            return None

        if self.strategy == "round-robin":
            key = active_keys[self._index % len(active_keys)]
            self._index += 1
        elif self.strategy == "least-used":
            key = min(active_keys, key=lambda k: k.request_count)
        elif self.strategy == "random":
            key = random.choice(active_keys)
        else:
            key = active_keys[0]

        key.request_count += 1
        return key.key

    def report_error(self, key_str: str, error_type: str):
        """Report key error."""
        for k in self.keys:
            if k.key == key_str:
                k.error_count += 1
                if error_type == "rate_limit":
                    k.exhausted_at = datetime.now()
                    k.reset_at = datetime.now() + timedelta(minutes=1)
                elif error_type == "invalid":
                    k.is_active = False

    def _is_exhausted(self, key: KeyState) -> bool:
        """Check if key is temporarily exhausted."""
        if key.reset_at and datetime.now() > key.reset_at:
            key.exhausted_at = None
            key.reset_at = None
            return False
        return key.exhausted_at is not None

    def get_status(self) -> list[dict]:
        """Get status of all keys."""
        return [
            {
                "key": k.key[:8] + "..." + k.key[-4:] if len(k.key) > 12 else "***",
                "active": k.is_active,
                "exhausted": k.exhausted_at is not None,
                "requests": k.request_count,
                "errors": k.error_count,
            }
            for k in self.keys
        ]

    def reset_counts(self):
        """Reset request and error counts."""
        for k in self.keys:
            k.request_count = 0
            k.error_count = 0
````

## File: tests/__init__.py
````python

````

## File: tests/conftest.py
````python
"""Pytest configuration and fixtures."""

import pytest
import asyncio
from typing import AsyncGenerator
from httpx import AsyncClient, ASGITransport

from gpt_proxy.main import create_app


@pytest.fixture(scope="session")
def event_loop():
    """Create event loop for async tests."""
    loop = asyncio.get_event_loop_policy().new_event_loop()
    yield loop
    loop.close()


@pytest.fixture
async def client() -> AsyncGenerator[AsyncClient, None]:
    """Create test client."""
    app = create_app()
    async with AsyncClient(
        transport=ASGITransport(app=app),
        base_url="http://test",
    ) as client:
        yield client


@pytest.fixture
def mock_session_token():
    """Mock ChatGPT session token."""
    return "mock-session-token-for-testing"


@pytest.fixture
def mock_auth_response():
    """Mock authentication response from ChatGPT."""
    return {
        "user": {
            "id": "user-123",
            "email": "test@example.com",
            "name": "Test User",
        },
        "accessToken": "mock-access-token-12345",
        "expires": "2025-01-01T00:00:00Z",
    }
````

## File: tests/integration/__init__.py
````python

````

## File: tests/integration/test_auth.py
````python
"""Integration tests for auth endpoints."""

import pytest
from httpx import AsyncClient
from unittest.mock import AsyncMock, patch


class TestAuthEndpoints:
    """Tests for authentication endpoints."""

    @pytest.mark.asyncio
    async def test_health_check(self, client: AsyncClient):
        response = await client.get("/health")
        assert response.status_code == 200
        assert response.json()["status"] == "healthy"

    @pytest.mark.asyncio
    async def test_auth_help(self, client: AsyncClient):
        response = await client.get("/auth/help")
        assert response.status_code == 200
        data = response.json()
        assert "steps" in data

    @pytest.mark.asyncio
    async def test_login_invalid_token(self, client: AsyncClient):
        response = await client.post(
            "/auth/login",
            json={"session_token": "invalid-token"},
        )
        assert response.status_code == 401

    @pytest.mark.asyncio
    async def test_login_success(self, client: AsyncClient, mock_auth_response):
        with patch(
            "gpt_proxy.services.auth_manager.AuthManager.exchange_session_token",
            new_callable=AsyncMock,
        ) as mock_exchange:
            from gpt_proxy.services.auth_manager import UserSession
            from datetime import datetime, timedelta

            mock_exchange.return_value = UserSession(
                session_id="test-session-id",
                user_id="user-123",
                email="test@example.com",
                access_token="access-token-123",
                session_token="session-token",
                expires_at=datetime.now() + timedelta(hours=1),
            )

            response = await client.post(
                "/auth/login",
                json={"session_token": "valid-token"},
            )

            assert response.status_code == 200
            data = response.json()
            assert data["session_id"] == "test-session-id"
            assert data["user_email"] == "test@example.com"

    @pytest.mark.asyncio
    async def test_chat_without_auth(self, client: AsyncClient):
        response = await client.post(
            "/v1/chat/completions",
            json={"model": "gpt-4", "messages": [{"role": "user", "content": "Hi"}]},
        )
        assert response.status_code == 401

    @pytest.mark.asyncio
    async def test_chat_with_invalid_session(self, client: AsyncClient):
        response = await client.post(
            "/v1/chat/completions",
            headers={"Authorization": "Bearer invalid-session-id"},
            json={"model": "gpt-4", "messages": [{"role": "user", "content": "Hi"}]},
        )
        assert response.status_code == 401
````

## File: tests/unit/__init__.py
````python

````

## File: tests/unit/test_auth.py
````python
"""Unit tests for auth manager."""

import pytest
from unittest.mock import AsyncMock, patch, MagicMock
from datetime import datetime, timedelta
import httpx
import json

from gpt_proxy.services.auth_manager import AuthManager, UserSession


class TestAuthManager:
    """Tests for authentication manager."""

    def test_create_session(self):
        manager = AuthManager()
        session = UserSession(
            session_id="test-id",
            user_id="user-123",
            email="test@example.com",
            access_token="token-123",
            session_token="session-123",
            expires_at=datetime.now() + timedelta(hours=1),
        )

        manager.create_session(session)
        assert "test-id" in manager.sessions

    def test_get_session(self):
        manager = AuthManager()
        session = UserSession(
            session_id="test-id",
            user_id="user-123",
            email="test@example.com",
            access_token="token-123",
            session_token="session-123",
            expires_at=datetime.now() + timedelta(hours=1),
        )
        manager.create_session(session)

        result = manager.get_session("test-id")
        assert result == session

    def test_invalidate_session(self):
        manager = AuthManager()
        session = UserSession(
            session_id="test-id",
            user_id="user-123",
            email="test@example.com",
            access_token="token-123",
            session_token="session-123",
            expires_at=datetime.now() + timedelta(hours=1),
        )
        manager.create_session(session)

        result = manager.invalidate_session("test-id")
        assert result is True
        assert manager.sessions["test-id"].is_active is False

    def test_list_sessions(self):
        with patch.object(AuthManager, '_load_sessions'), \
             patch.object(AuthManager, '_save_sessions'):
            manager = AuthManager()
            session = UserSession(
                session_id="test-id",
                user_id="user-123",
                email="test@example.com",
                access_token="token-123",
                session_token="session-123",
                expires_at=datetime.now() + timedelta(hours=1),
            )
            manager.create_session(session)

            sessions = manager.list_sessions()
            assert len(sessions) == 1
            assert sessions[0]["email"] == "test@example.com"

    @pytest.mark.asyncio
    async def test_get_valid_token(self):
        manager = AuthManager()
        session = UserSession(
            session_id="test-id",
            user_id="user-123",
            email="test@example.com",
            access_token="token-123",
            session_token="session-123",
            expires_at=datetime.now() + timedelta(hours=1),
        )
        manager.create_session(session)

        token = await manager.get_valid_token("test-id")
        assert token == "token-123"

    @pytest.mark.asyncio
    async def test_get_valid_token_expired(self):
        manager = AuthManager()
        session = UserSession(
            session_id="test-id",
            user_id="user-123",
            email="test@example.com",
            access_token="old-token",
            session_token="session-123",
            expires_at=datetime.now() - timedelta(hours=1),  # Expired
        )
        manager.create_session(session)

        # Should try to refresh and fail (no mock)
        token = await manager.get_valid_token("test-id")
        assert token is None


class TestAuthManagerErrorHandling:
    """Tests for error handling in auth manager."""

    @pytest.mark.asyncio
    async def test_exchange_token_empty_response(self):
        """Test handling of empty response body."""
        manager = AuthManager()

        with patch.object(manager, '_get_client') as mock_client:
            mock_response = MagicMock()
            mock_response.status_code = 200
            mock_response.headers = {"content-type": "application/json"}
            mock_response.text = ""
            mock_response.json.side_effect = json.JSONDecodeError("Expecting value", "", 0)

            mock_http_client = AsyncMock()
            mock_http_client.get = AsyncMock(return_value=mock_response)
            mock_client.return_value = mock_http_client

            result = await manager.exchange_session_token("test-token")
            assert result is None

    @pytest.mark.asyncio
    async def test_exchange_token_html_response(self):
        """Test handling of HTML response (Cloudflare challenge)."""
        manager = AuthManager()

        with patch.object(manager, '_get_client') as mock_client:
            mock_response = MagicMock()
            mock_response.status_code = 403
            mock_response.headers = {"content-type": "text/html"}
            mock_response.text = "<html><body>Cloudflare challenge</body></html>"

            mock_http_client = AsyncMock()
            mock_http_client.get = AsyncMock(return_value=mock_response)
            mock_client.return_value = mock_http_client

            result = await manager.exchange_session_token("test-token")
            assert result is None

    @pytest.mark.asyncio
    async def test_cloudflare_detection(self):
        """Test Cloudflare challenge detection."""
        manager = AuthManager()

        # Test 403 with Cloudflare
        response = MagicMock()
        response.status_code = 403
        response.text = "<html>Just a moment... Checking your browser</html>"
        assert manager._is_cloudflare_challenge(response) is True

        # Test normal 403
        response.text = "<html>Access Denied</html>"
        assert manager._is_cloudflare_challenge(response) is False

        # Test 503 with Cloudflare
        response.status_code = 503
        response.text = "<html>cloudflare challenge-platform</html>"
        assert manager._is_cloudflare_challenge(response) is True

    @pytest.mark.asyncio
    async def test_timeout_handling(self):
        """Test handling of timeout errors."""
        manager = AuthManager()

        with patch.object(manager, '_get_client') as mock_client:
            mock_http_client = AsyncMock()
            mock_http_client.get = AsyncMock(side_effect=httpx.TimeoutException("Timeout"))
            mock_client.return_value = mock_http_client

            result = await manager.exchange_session_token("test-token")
            assert result is None


class TestTokenValidation:
    """Tests for token validation."""

    @pytest.mark.asyncio
    async def test_validate_valid_token(self):
        """Test validation of valid token."""
        manager = AuthManager()

        with patch.object(manager, '_get_client') as mock_client:
            mock_response = MagicMock()
            mock_response.status_code = 200
            mock_response.headers = {"content-type": "application/json"}
            mock_response.json.return_value = {
                "user": {"email": "test@example.com"},
                "accessToken": "test-access-token"
            }

            mock_http_client = AsyncMock()
            mock_http_client.get = AsyncMock(return_value=mock_response)
            mock_client.return_value = mock_http_client

            result = await manager.validate_session_token("valid-token")
            assert result["valid"] is True
            assert result["user_email"] == "test@example.com"

    @pytest.mark.asyncio
    async def test_validate_invalid_token(self):
        """Test validation of invalid token."""
        manager = AuthManager()

        with patch.object(manager, '_get_client') as mock_client:
            mock_response = MagicMock()
            mock_response.status_code = 401

            mock_http_client = AsyncMock()
            mock_http_client.get = AsyncMock(return_value=mock_response)
            mock_client.return_value = mock_http_client

            result = await manager.validate_session_token("invalid-token")
            assert result["valid"] is False
            assert result["error"] == "invalid_token"

    @pytest.mark.asyncio
    async def test_validate_cloudflare_challenge(self):
        """Test validation when Cloudflare challenge is present."""
        manager = AuthManager()

        with patch.object(manager, '_get_client') as mock_client:
            mock_response = MagicMock()
            mock_response.status_code = 403
            mock_response.text = "<html>Just a moment... Cloudflare</html>"

            mock_http_client = AsyncMock()
            mock_http_client.get = AsyncMock(return_value=mock_response)
            mock_client.return_value = mock_http_client

            result = await manager.validate_session_token("test-token")
            assert result["valid"] is False
            assert result["error"] == "cloudflare_challenge"
````

## File: tests/unit/test_cache.py
````python
"""Unit tests for response cache."""

import pytest
import asyncio
from datetime import datetime, timedelta
from unittest.mock import patch

from gpt_proxy.services.cache import ResponseCache


class TestGenerateKey:
    """Tests for cache key generation."""

    def test_deterministic(self):
        cache = ResponseCache()
        key1 = cache.generate_key("POST", "/v1/chat/completions", b'{"model":"gpt-4"}')
        key2 = cache.generate_key("POST", "/v1/chat/completions", b'{"model":"gpt-4"}')
        assert key1 == key2

    def test_different_for_different_inputs(self):
        cache = ResponseCache()
        key1 = cache.generate_key("POST", "/v1/chat/completions", b'{"model":"gpt-4"}')
        key2 = cache.generate_key("POST", "/v1/chat/completions", b'{"model":"gpt-3.5"}')
        key3 = cache.generate_key("GET", "/v1/models", None)
        assert key1 != key2
        assert key1 != key3
        assert key2 != key3

    def test_none_body_handled(self):
        cache = ResponseCache()
        key = cache.generate_key("GET", "/v1/models", None)
        assert isinstance(key, str)
        assert len(key) == 64  # SHA-256 hex


class TestCacheGetSet:
    """Tests for cache get/set operations."""

    @pytest.mark.asyncio
    async def test_set_and_get(self):
        cache = ResponseCache()
        key = cache.generate_key("GET", "/test", None)
        await cache.set(key, b"response-data")
        result = await cache.get(key)
        assert result == b"response-data"

    @pytest.mark.asyncio
    async def test_cache_miss_returns_none(self):
        cache = ResponseCache()
        result = await cache.get("nonexistent-key")
        assert result is None


class TestCacheTTL:
    """Tests for TTL expiration."""

    @pytest.mark.asyncio
    async def test_expired_entry_returns_none(self):
        cache = ResponseCache(ttl_seconds=1)
        key = "test-key"
        await cache.set(key, b"data")

        # Manually set timestamp to the past
        old_time = datetime.now() - timedelta(seconds=2)
        cache._cache[key] = (b"data", old_time)

        result = await cache.get(key)
        assert result is None

    @pytest.mark.asyncio
    async def test_fresh_entry_returns_data(self):
        cache = ResponseCache(ttl_seconds=3600)
        key = "test-key"
        await cache.set(key, b"data")
        result = await cache.get(key)
        assert result == b"data"

    @pytest.mark.asyncio
    async def test_expired_entry_removed_from_cache(self):
        cache = ResponseCache(ttl_seconds=1)
        key = "test-key"
        await cache.set(key, b"data")

        old_time = datetime.now() - timedelta(seconds=2)
        cache._cache[key] = (b"data", old_time)

        await cache.get(key)
        assert key not in cache._cache


class TestCacheClear:
    """Tests for cache clear."""

    @pytest.mark.asyncio
    async def test_clear_resets_cache(self):
        cache = ResponseCache()
        await cache.set("key1", b"data1")
        await cache.set("key2", b"data2")
        await cache.get("key1")  # hit

        await cache.clear()
        assert len(cache._cache) == 0
        assert cache._hits == 0
        assert cache._misses == 0

    @pytest.mark.asyncio
    async def test_get_after_clear_is_miss(self):
        cache = ResponseCache()
        key = cache.generate_key("GET", "/test", None)
        await cache.set(key, b"data")
        await cache.clear()
        result = await cache.get(key)
        assert result is None


class TestCacheStats:
    """Tests for cache statistics."""

    def test_initial_stats(self):
        cache = ResponseCache()
        stats = cache.stats()
        assert stats["size"] == 0
        assert stats["hits"] == 0
        assert stats["misses"] == 0
        assert stats["hit_rate"] == 0

    @pytest.mark.asyncio
    async def test_stats_after_operations(self):
        cache = ResponseCache()
        await cache.set("key1", b"data")
        await cache.get("key1")  # hit
        await cache.get("missing")  # miss

        stats = cache.stats()
        assert stats["size"] == 1
        assert stats["hits"] == 1
        assert stats["misses"] == 1
        assert stats["hit_rate"] == 0.5

    def test_stats_ttl_value(self):
        cache = ResponseCache(ttl_seconds=7200)
        assert cache.stats()["ttl_seconds"] == 7200
````

## File: tests/unit/test_cost_tracker.py
````python
"""Unit tests for cost tracker."""

import pytest
from datetime import datetime, timedelta
from unittest.mock import patch

from gpt_proxy.services.cost_tracker import CostTracker


class TestCountTokens:
    """Tests for token counting."""

    def test_counts_tokens_for_known_model(self):
        tracker = CostTracker()
        count = tracker.count_tokens("Hello, world!", "gpt-4")
        assert isinstance(count, int)
        assert count > 0

    def test_counts_tokens_for_unknown_model_falls_back(self):
        tracker = CostTracker()
        count = tracker.count_tokens("Hello, world!", "unknown-model-xyz")
        assert isinstance(count, int)
        assert count > 0

    def test_empty_string_returns_zero_or_low(self):
        tracker = CostTracker()
        count = tracker.count_tokens("", "gpt-4")
        assert count == 0


class TestCountMessagesTokens:
    """Tests for message token counting."""

    def test_string_content(self):
        tracker = CostTracker()
        messages = [
            {"role": "user", "content": "Hello!"},
            {"role": "assistant", "content": "Hi there!"},
        ]
        count = tracker.count_messages_tokens(messages, "gpt-4")
        # 2 messages * 4 overhead + tokens + 2 reply priming
        assert count > 10

    def test_list_content(self):
        tracker = CostTracker()
        messages = [
            {
                "role": "user",
                "content": [
                    {"type": "text", "text": "Describe this image"},
                ],
            }
        ]
        count = tracker.count_messages_tokens(messages, "gpt-4")
        assert count > 4  # At least overhead

    def test_list_content_ignores_non_text_parts(self):
        tracker = CostTracker()
        messages = [
            {
                "role": "user",
                "content": [
                    {"type": "image_url", "image_url": {"url": "http://example.com/img.png"}},
                ],
            }
        ]
        count = tracker.count_messages_tokens(messages, "gpt-4")
        # Only 4 overhead + 2 reply priming = 6
        assert count == 6

    def test_empty_messages(self):
        tracker = CostTracker()
        count = tracker.count_messages_tokens([], "gpt-4")
        assert count == 2  # Just reply priming


class TestTrackUsage:
    """Tests for usage tracking."""

    @pytest.mark.asyncio
    async def test_records_usage(self):
        tracker = CostTracker()
        await tracker.track_usage("key-1", "gpt-4o", 100, 50)
        assert len(tracker._usage) == 1
        entry = tracker._usage[0]
        assert entry["key_id"] == "key-1"
        assert entry["model"] == "gpt-4o"
        assert entry["prompt_tokens"] == 100
        assert entry["completion_tokens"] == 50
        assert entry["total_tokens"] == 150

    @pytest.mark.asyncio
    async def test_calculates_cost(self):
        tracker = CostTracker()
        await tracker.track_usage("key-1", "gpt-4o", 1_000_000, 0)
        entry = tracker._usage[0]
        # gpt-4o input = $2.50 per 1M tokens
        assert entry["cost_usd"] == pytest.approx(2.50, rel=1e-4)

    @pytest.mark.asyncio
    async def test_unknown_model_zero_cost(self):
        tracker = CostTracker()
        await tracker.track_usage("key-1", "nonexistent-model", 1000, 1000)
        entry = tracker._usage[0]
        assert entry["cost_usd"] == 0.0


class TestGetUsage:
    """Tests for usage retrieval."""

    @pytest.mark.asyncio
    async def test_get_usage_day_period(self):
        tracker = CostTracker()
        await tracker.track_usage("key-1", "gpt-4o", 100, 50)
        usage = await tracker.get_usage(period="day")
        assert usage["period"] == "day"
        assert usage["total_requests"] == 1
        assert usage["prompt_tokens"] == 100
        assert usage["completion_tokens"] == 50

    @pytest.mark.asyncio
    async def test_get_usage_filters_by_key_id(self):
        tracker = CostTracker()
        await tracker.track_usage("key-a", "gpt-4o", 100, 50)
        await tracker.track_usage("key-b", "gpt-4o", 200, 75)

        usage_a = await tracker.get_usage(period="day", key_id="key-a")
        assert usage_a["total_requests"] == 1
        assert usage_a["prompt_tokens"] == 100

        usage_b = await tracker.get_usage(period="day", key_id="key-b")
        assert usage_b["total_requests"] == 1
        assert usage_b["prompt_tokens"] == 200

    @pytest.mark.asyncio
    async def test_get_usage_excludes_old_entries(self):
        tracker = CostTracker()
        # Add a current entry
        await tracker.track_usage("key-1", "gpt-4o", 100, 50)
        # Manually add an old entry (2 days ago)
        old_time = (datetime.now() - timedelta(days=2)).isoformat()
        tracker._usage.append({
            "timestamp": old_time,
            "key_id": "key-1",
            "model": "gpt-4o",
            "prompt_tokens": 999,
            "completion_tokens": 999,
            "total_tokens": 1998,
            "cost_usd": 1.0,
        })

        usage = await tracker.get_usage(period="day")
        assert usage["total_requests"] == 1
        assert usage["prompt_tokens"] == 100

    @pytest.mark.asyncio
    async def test_get_usage_week_period(self):
        tracker = CostTracker()
        await tracker.track_usage("key-1", "gpt-4o", 100, 50)
        usage = await tracker.get_usage(period="week")
        assert usage["period"] == "week"


class TestAggregateByModel:
    """Tests for model aggregation."""

    def test_aggregates_correctly(self):
        tracker = CostTracker()
        usage = [
            {
                "model": "gpt-4o",
                "total_tokens": 100,
                "cost_usd": 0.01,
            },
            {
                "model": "gpt-4o",
                "total_tokens": 200,
                "cost_usd": 0.02,
            },
            {
                "model": "gpt-3.5-turbo",
                "total_tokens": 50,
                "cost_usd": 0.001,
            },
        ]
        result = tracker._aggregate_by_model(usage)
        assert result["gpt-4o"]["tokens"] == 300
        assert result["gpt-4o"]["cost"] == pytest.approx(0.03, rel=1e-4)
        assert result["gpt-4o"]["requests"] == 2
        assert result["gpt-3.5-turbo"]["tokens"] == 50
        assert result["gpt-3.5-turbo"]["requests"] == 1

    def test_empty_usage_returns_empty(self):
        tracker = CostTracker()
        result = tracker._aggregate_by_model([])
        assert result == {}


class TestClear:
    """Tests for clear."""

    @pytest.mark.asyncio
    async def test_clears_usage(self):
        tracker = CostTracker()
        await tracker.track_usage("key-1", "gpt-4o", 100, 50)
        await tracker.track_usage("key-2", "gpt-4", 200, 75)
        tracker.clear()
        assert len(tracker._usage) == 0

    @pytest.mark.asyncio
    async def test_get_usage_after_clear(self):
        tracker = CostTracker()
        await tracker.track_usage("key-1", "gpt-4o", 100, 50)
        tracker.clear()
        usage = await tracker.get_usage(period="day")
        assert usage["total_requests"] == 0
````

## File: tests/unit/test_key_manager.py
````python
"""Unit tests for API key manager."""

import pytest
from unittest.mock import patch
from datetime import datetime, timedelta

from gpt_proxy.services.key_manager import APIKeyManager, KeyState


class TestRoundRobinStrategy:
    """Tests for round-robin key rotation."""

    def test_rotates_through_keys(self):
        manager = APIKeyManager(keys=["key-a", "key-b", "key-c"], strategy="round-robin")
        results = [manager.get_key() for _ in range(6)]
        assert results == ["key-a", "key-b", "key-c", "key-a", "key-b", "key-c"]

    def test_increments_request_count(self):
        manager = APIKeyManager(keys=["key-a", "key-b"], strategy="round-robin")
        manager.get_key()
        manager.get_key()
        assert manager.keys[0].request_count == 1
        assert manager.keys[1].request_count == 1


class TestLeastUsedStrategy:
    """Tests for least-used key selection."""

    def test_selects_least_used_key(self):
        manager = APIKeyManager(keys=["key-a", "key-b", "key-c"], strategy="least-used")
        # Use key-a twice
        manager.get_key()  # picks key-a (all tied)
        manager.get_key()  # picks key-b (all tied except key-a)
        manager.get_key()  # picks key-c (all tied except key-a, key-b)

        # Now key-a, key-b, key-c each have 1 use, so next should be key-a (first min)
        # Actually let's force usage
        manager.keys[0].request_count = 5
        manager.keys[1].request_count = 2
        manager.keys[2].request_count = 1
        result = manager.get_key()
        assert result == "key-c"

    def test_increments_request_count(self):
        manager = APIKeyManager(keys=["key-a"], strategy="least-used")
        manager.get_key()
        assert manager.keys[0].request_count == 1


class TestRandomStrategy:
    """Tests for random key selection."""

    def test_selects_from_active_keys(self):
        manager = APIKeyManager(keys=["key-a", "key-b", "key-c"], strategy="random")
        selected = set()
        for _ in range(100):
            selected.add(manager.get_key())
        assert selected == {"key-a", "key-b", "key-c"}

    def test_increments_request_count(self):
        manager = APIKeyManager(keys=["key-a"], strategy="random")
        manager.get_key()
        manager.get_key()
        assert manager.keys[0].request_count == 2


class TestKeyExhaustionAndReset:
    """Tests for key exhaustion and automatic reset."""

    def test_exhausted_key_skipped(self):
        manager = APIKeyManager(keys=["key-a", "key-b"], strategy="round-robin")
        # Exhaust key-a
        manager.keys[0].exhausted_at = datetime.now()
        manager.keys[0].reset_at = datetime.now() + timedelta(minutes=1)
        assert manager.get_key() == "key-b"

    def test_all_exhausted_returns_none(self):
        manager = APIKeyManager(keys=["key-a", "key-b"], strategy="round-robin")
        for k in manager.keys:
            k.exhausted_at = datetime.now()
            k.reset_at = datetime.now() + timedelta(minutes=1)
        assert manager.get_key() is None

    def test_exhausted_key_resets_after_reset_time(self):
        manager = APIKeyManager(keys=["key-a"], strategy="round-robin")
        manager.keys[0].exhausted_at = datetime.now() - timedelta(minutes=2)
        manager.keys[0].reset_at = datetime.now() - timedelta(minutes=1)
        # Reset time has passed, should be available again
        result = manager.get_key()
        assert result == "key-a"
        assert manager.keys[0].exhausted_at is None
        assert manager.keys[0].reset_at is None

    def test_inactive_key_skipped(self):
        manager = APIKeyManager(keys=["key-a", "key-b"], strategy="round-robin")
        manager.keys[0].is_active = False
        assert manager.get_key() == "key-b"


class TestErrorReporting:
    """Tests for error reporting."""

    def test_rate_limit_marks_exhausted(self):
        manager = APIKeyManager(keys=["key-a"], strategy="round-robin")
        manager.report_error("key-a", "rate_limit")
        assert manager.keys[0].exhausted_at is not None
        assert manager.keys[0].reset_at is not None
        assert manager.keys[0].error_count == 1

    def test_invalid_deactivates_key(self):
        manager = APIKeyManager(keys=["key-a"], strategy="round-robin")
        manager.report_error("key-a", "invalid")
        assert manager.keys[0].is_active is False
        assert manager.keys[0].error_count == 1

    def test_unknown_error_only_increments_count(self):
        manager = APIKeyManager(keys=["key-a"], strategy="round-robin")
        manager.report_error("key-a", "server_error")
        assert manager.keys[0].error_count == 1
        assert manager.keys[0].is_active is True
        assert manager.keys[0].exhausted_at is None


class TestGetStatus:
    """Tests for get_status."""

    def test_masks_long_keys(self):
        manager = APIKeyManager(keys=["sk-abcdefghijkmnopqrstuvwx"], strategy="round-robin")
        status = manager.get_status()
        assert status[0]["key"] == "sk-abcde...uvwx"

    def test_masks_short_keys(self):
        manager = APIKeyManager(keys=["short"], strategy="round-robin")
        status = manager.get_status()
        assert status[0]["key"] == "***"

    def test_status_fields(self):
        manager = APIKeyManager(keys=["key-a"], strategy="round-robin")
        manager.get_key()
        manager.report_error("key-a", "rate_limit")
        status = manager.get_status()
        assert len(status) == 1
        assert status[0]["active"] is True
        assert status[0]["exhausted"] is True
        assert status[0]["requests"] == 1
        assert status[0]["errors"] == 1


class TestEmptyKeyList:
    """Tests for empty key list."""

    def test_returns_none(self):
        manager = APIKeyManager(keys=[], strategy="round-robin")
        assert manager.get_key() is None

    def test_status_empty(self):
        manager = APIKeyManager(keys=[], strategy="round-robin")
        assert manager.get_status() == []


class TestResetCounts:
    """Tests for reset_counts."""

    def test_resets_all_counts(self):
        manager = APIKeyManager(keys=["key-a", "key-b"], strategy="round-robin")
        manager.get_key()
        manager.get_key()
        manager.report_error("key-a", "server_error")
        manager.reset_counts()
        for k in manager.keys:
            assert k.request_count == 0
            assert k.error_count == 0
````

## File: tests/unit/test_router.py
````python
"""Unit tests for API router."""

import pytest
from unittest.mock import AsyncMock, patch, MagicMock
from httpx import AsyncClient, ASGITransport

from gpt_proxy.main import create_app


@pytest.fixture
def app():
    return create_app()


@pytest.fixture
async def client(app):
    async with AsyncClient(
        transport=ASGITransport(app=app),
        base_url="http://test",
    ) as client:
        yield client


class TestAuthRequired:
    """Tests that protected endpoints require authentication."""

    @pytest.mark.asyncio
    async def test_models_requires_auth(self, client):
        response = await client.get("/v1/models")
        assert response.status_code == 401
        data = response.json()
        assert data["detail"]["error"]["code"] == "missing_auth"

    @pytest.mark.asyncio
    async def test_chat_completions_requires_auth(self, client):
        response = await client.post(
            "/v1/chat/completions",
            json={"model": "gpt-4", "messages": [{"role": "user", "content": "hi"}]},
        )
        assert response.status_code == 401
        data = response.json()
        assert data["detail"]["error"]["code"] == "missing_auth"

    @pytest.mark.asyncio
    async def test_invalid_session_returns_401(self, client):
        response = await client.get(
            "/v1/models",
            headers={"Authorization": "Bearer invalid-session-id"},
        )
        assert response.status_code == 401
        data = response.json()
        assert data["detail"]["error"]["code"] == "session_expired"

    @pytest.mark.asyncio
    async def test_malformed_auth_header_returns_401(self, client):
        response = await client.get(
            "/v1/models",
            headers={"Authorization": "Token abc123"},
        )
        assert response.status_code == 401


class TestHealthEndpoints:
    """Tests for health check endpoints (no auth required)."""

    @pytest.mark.asyncio
    async def test_health_endpoint(self, client):
        response = await client.get("/health")
        assert response.status_code == 200
        data = response.json()
        assert data["status"] == "healthy"

    @pytest.mark.asyncio
    async def test_ready_endpoint(self, client):
        response = await client.get("/ready")
        assert response.status_code == 200
        data = response.json()
        assert data["status"] == "ready"


class TestModelsEndpoint:
    """Tests for /v1/models endpoint."""

    @pytest.mark.asyncio
    async def test_models_with_valid_session(self, client):
        mock_response = MagicMock()
        mock_response.content = b'{"data": [{"id": "gpt-4"}]}'
        mock_response.status_code = 200
        mock_response.headers = {"content-type": "application/json"}

        with patch("gpt_proxy.api.router.get_auth_manager") as mock_get_auth, \
             patch("gpt_proxy.api.router.chatgpt_request", new_callable=AsyncMock, return_value=mock_response):
            mock_auth = MagicMock()
            mock_auth.get_valid_token = AsyncMock(return_value="valid-access-token")
            mock_get_auth.return_value = mock_auth

            response = await client.get(
                "/v1/models",
                headers={"Authorization": "Bearer valid-session-id"},
            )
            assert response.status_code == 200


class TestChatCompletionsEndpoint:
    """Tests for /v1/chat/completions endpoint."""

    @pytest.mark.asyncio
    async def test_non_streaming_completions(self, client):
        mock_response = MagicMock()
        mock_response.content = b'{"choices": [{"message": {"content": "Hello!"}}]}'
        mock_response.status_code = 200
        mock_response.headers = {"content-type": "application/json"}

        with patch("gpt_proxy.api.router.get_auth_manager") as mock_get_auth, \
             patch("gpt_proxy.api.router.chatgpt_request", new_callable=AsyncMock, return_value=mock_response):
            mock_auth = MagicMock()
            mock_auth.get_valid_token = AsyncMock(return_value="valid-access-token")
            mock_get_auth.return_value = mock_auth

            response = await client.post(
                "/v1/chat/completions",
                headers={"Authorization": "Bearer valid-session-id"},
                json={
                    "model": "gpt-4",
                    "messages": [{"role": "user", "content": "Hello!"}],
                },
            )
            assert response.status_code == 200

    @pytest.mark.asyncio
    async def test_streaming_completions(self, client):
        async def mock_stream(**kwargs):
            yield b'data: {"choices": [{"delta": {"content": "Hi"}}]}\n\n'
            yield b'data: [DONE]\n\n'

        with patch("gpt_proxy.api.router.get_auth_manager") as mock_get_auth, \
             patch("gpt_proxy.api.router.chatgpt_stream", side_effect=lambda **kwargs: mock_stream()):
            mock_auth = MagicMock()
            mock_auth.get_valid_token = AsyncMock(return_value="valid-access-token")
            mock_get_auth.return_value = mock_auth

            response = await client.post(
                "/v1/chat/completions",
                headers={"Authorization": "Bearer valid-session-id"},
                json={
                    "model": "gpt-4",
                    "messages": [{"role": "user", "content": "Hello!"}],
                    "stream": True,
                },
            )
            assert response.status_code == 200
            assert "text/event-stream" in response.headers.get("content-type", "")
````
