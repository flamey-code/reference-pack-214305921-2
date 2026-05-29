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
.dockerignore
.env.example
.github/workflows/build_docker_dev.yml
.github/workflows/build_docker_main.yml
.gitignore
api/chat2api.py
api/files.py
api/models.py
api/tokens.py
app.py
chatgpt/authorization.py
chatgpt/chatFormat_v1.py
chatgpt/chatFormat.py
chatgpt/chatLimit.py
chatgpt/ChatService.py
chatgpt/fp.py
chatgpt/proofofWork.py
chatgpt/refreshToken.py
chatgpt/turnstile.py
chatgpt/wssClient.py
docker-compose-warp.yml
docker-compose.yml
Dockerfile
docs/capsolver.png
docs/chatgpt.png
docs/login.png
docs/tokens.png
gateway/admin.py
gateway/backend.py
gateway/chatgpt.py
gateway/gpts.py
gateway/login.py
gateway/reverseProxy.py
gateway/route.py
gateway/share.py
gateway/v1.py
LICENSE
README.md
requirements.txt
templates/chatgpt_context_1.json
templates/chatgpt_context_2.json
templates/chatgpt.html
templates/gpts_context.json
templates/initialize.json
templates/login.html
templates/tokens.html
utils/Client.py
utils/configs.py
utils/globals.py
utils/kv_utils.py
utils/Logger.py
utils/retry.py
version.txt
```

# Files

## File: .dockerignore
````
.env
*.pyc
/.git/
/.idea/
/docs/
/tmp/
/data/
/.venv/
/.vscode/
````

## File: .env.example
````
API_PREFIX=your_prefix
CHATGPT_BASE_URL=https://chatgpt.com
PROXY_URL=your_first_proxy, your_second_proxy
SCHEDULED_REFRESH=false
````

## File: .github/workflows/build_docker_dev.yml
````yaml
name: Build Docker Image (dev)

on:
  push:
    branches:
      - dev
    paths-ignore:
      - 'README.md'
      - 'docker-compose.yml'
      - 'docker-compose-warp.yml'
      - 'docs/**'
      - '.github/workflows/build_docker_main.yml'
      - '.github/workflows/build_docker_dev.yml'
  workflow_dispatch:

jobs:
  main:
    runs-on: ubuntu-latest

    steps:
      - name: Check out the repository
        uses: actions/checkout@v2

      - name: Read the version from version.txt
        id: get_version
        run: |
          version=$(cat version.txt)
          echo "Current version: v$version-dev"
          echo "::set-output name=version::v$version-dev"

      - name: Commit and push version tag
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        run: |
          version=${{ steps.get_version.outputs.version }}
          git config --local user.email "action@github.com"
          git config --local user.name "GitHub Action"
          git tag "$version"
          git push https://x-access-token:${GHCR_PAT}@github.com/lanqian528/chat2api.git "$version"

      - name: Set up QEMU
        uses: docker/setup-qemu-action@v3

      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v3

      - name: Log in to Docker Hub
        uses: docker/login-action@v3
        with:
          username: ${{ secrets.DOCKER_USERNAME }}
          password: ${{ secrets.DOCKER_PASSWORD }}

      - name: Docker meta
        id: meta
        uses: docker/metadata-action@v5
        with:
          images: lanqian528/chat2api
          tags: |
            type=raw,value=latest-dev
            type=raw,value=${{ steps.get_version.outputs.version }}

      - name: Build and push
        uses: docker/build-push-action@v5
        with:
          context: .
          platforms: linux/amd64,linux/arm64
          file: Dockerfile
          push: true
          tags: ${{ steps.meta.outputs.tags }}
          labels: ${{ steps.meta.outputs.labels }}
````

## File: .github/workflows/build_docker_main.yml
````yaml
name: Build Docker Image

on:
  push:
    branches:
      - main
    paths-ignore:
      - 'README.md'
      - 'docker-compose.yml'
      - 'docker-compose-warp.yml'
      - 'docs/**'
      - '.github/workflows/build_docker_main.yml'
      - '.github/workflows/build_docker_dev.yml'
  workflow_dispatch:

jobs:
  main:
    runs-on: ubuntu-latest

    steps:
      - name: Check out the repository
        uses: actions/checkout@v2

      - name: Read the version from version.txt
        id: get_version
        run: |
          version=$(cat version.txt)
          echo "Current version: v$version"
          echo "::set-output name=version::v$version"

      - name: Commit and push version tag
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        run: |
          version=${{ steps.get_version.outputs.version }}
          git config --local user.email "action@github.com"
          git config --local user.name "GitHub Action"
          git tag "$version"
          git push https://x-access-token:${GHCR_PAT}@github.com/lanqian528/chat2api.git "$version"

      - name: Set up QEMU
        uses: docker/setup-qemu-action@v3

      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v3

      - name: Log in to Docker Hub
        uses: docker/login-action@v3
        with:
          username: ${{ secrets.DOCKER_USERNAME }}
          password: ${{ secrets.DOCKER_PASSWORD }}

      - name: Docker meta
        id: meta
        uses: docker/metadata-action@v5
        with:
          images: lanqian528/chat2api
          tags: |
            type=raw,value=latest,enable={{is_default_branch}}
            type=raw,value=${{ steps.get_version.outputs.version }}

      - name: Build and push
        uses: docker/build-push-action@v5
        with:
          context: .
          platforms: linux/amd64,linux/arm64
          file: Dockerfile
          push: true
          tags: ${{ steps.meta.outputs.tags }}
          labels: ${{ steps.meta.outputs.labels }}
````

## File: .gitignore
````
.env
*.pyc
/.git/
/.idea/
/tmp/
/data/
/.venv/
/.vscode/
````

## File: api/chat2api.py
````python
import asyncio
import types

from apscheduler.schedulers.asyncio import AsyncIOScheduler
from fastapi import Request, HTTPException, Form, Security
from fastapi.responses import HTMLResponse, StreamingResponse, JSONResponse
from fastapi.security import HTTPAuthorizationCredentials
from starlette.background import BackgroundTask

import utils.globals as globals
from app import app, templates, security_scheme
from chatgpt.ChatService import ChatService
from chatgpt.authorization import refresh_all_tokens
from utils.Logger import logger
from utils.configs import api_prefix, scheduled_refresh
from utils.retry import async_retry

scheduler = AsyncIOScheduler()


@app.on_event("startup")
async def app_start():
    if scheduled_refresh:
        scheduler.add_job(id='refresh', func=refresh_all_tokens, trigger='cron', hour=3, minute=0, day='*/2',
                          kwargs={'force_refresh': True})
        scheduler.start()
        asyncio.get_event_loop().call_later(0, lambda: asyncio.create_task(refresh_all_tokens(force_refresh=False)))


async def to_send_conversation(request_data, req_token):
    chat_service = ChatService(req_token)
    try:
        await chat_service.set_dynamic_data(request_data)
        await chat_service.get_chat_requirements()
        return chat_service
    except HTTPException as e:
        await chat_service.close_client()
        raise HTTPException(status_code=e.status_code, detail=e.detail)
    except Exception as e:
        await chat_service.close_client()
        logger.error(f"Server error, {str(e)}")
        raise HTTPException(status_code=500, detail="Server error")


async def process(request_data, req_token):
    chat_service = await to_send_conversation(request_data, req_token)
    await chat_service.prepare_send_conversation()
    res = await chat_service.send_conversation()
    return chat_service, res


@app.post(f"/{api_prefix}/v1/chat/completions" if api_prefix else "/v1/chat/completions")
async def send_conversation(request: Request, credentials: HTTPAuthorizationCredentials = Security(security_scheme)):
    req_token = credentials.credentials
    try:
        request_data = await request.json()
    except Exception:
        raise HTTPException(status_code=400, detail={"error": "Invalid JSON body"})
    chat_service, res = await async_retry(process, request_data, req_token)
    try:
        if isinstance(res, types.AsyncGeneratorType):
            background = BackgroundTask(chat_service.close_client)
            return StreamingResponse(res, media_type="text/event-stream", background=background)
        else:
            background = BackgroundTask(chat_service.close_client)
            return JSONResponse(res, media_type="application/json", background=background)
    except HTTPException as e:
        await chat_service.close_client()
        if e.status_code == 500:
            logger.error(f"Server error, {str(e)}")
            raise HTTPException(status_code=500, detail="Server error")
        raise HTTPException(status_code=e.status_code, detail=e.detail)
    except Exception as e:
        await chat_service.close_client()
        logger.error(f"Server error, {str(e)}")
        raise HTTPException(status_code=500, detail="Server error")


@app.get(f"/{api_prefix}/tokens" if api_prefix else "/tokens", response_class=HTMLResponse)
async def upload_html(request: Request):
    tokens_count = len(set(globals.token_list) - set(globals.error_token_list))
    return templates.TemplateResponse("tokens.html",
                                      {"request": request, "api_prefix": api_prefix, "tokens_count": tokens_count})


@app.post(f"/{api_prefix}/tokens/upload" if api_prefix else "/tokens/upload")
async def upload_post(text: str = Form(...)):
    lines = text.split("\n")
    for line in lines:
        if line.strip() and not line.startswith("#"):
            globals.token_list.append(line.strip())
            with open(globals.TOKENS_FILE, "a", encoding="utf-8") as f:
                f.write(line.strip() + "\n")
    logger.info(f"Token count: {len(globals.token_list)}, Error token count: {len(globals.error_token_list)}")
    tokens_count = len(set(globals.token_list) - set(globals.error_token_list))
    return {"status": "success", "tokens_count": tokens_count}


@app.post(f"/{api_prefix}/tokens/clear" if api_prefix else "/tokens/clear")
async def clear_tokens():
    globals.token_list.clear()
    globals.error_token_list.clear()
    with open(globals.TOKENS_FILE, "w", encoding="utf-8") as f:
        pass
    logger.info(f"Token count: {len(globals.token_list)}, Error token count: {len(globals.error_token_list)}")
    tokens_count = len(set(globals.token_list) - set(globals.error_token_list))
    return {"status": "success", "tokens_count": tokens_count}


@app.post(f"/{api_prefix}/tokens/error" if api_prefix else "/tokens/error")
async def error_tokens():
    error_tokens_list = list(set(globals.error_token_list))
    return {"status": "success", "error_tokens": error_tokens_list}


@app.get(f"/{api_prefix}/tokens/add/{{token}}" if api_prefix else "/tokens/add/{token}")
async def add_token(token: str):
    if token.strip() and not token.startswith("#"):
        globals.token_list.append(token.strip())
        with open(globals.TOKENS_FILE, "a", encoding="utf-8") as f:
            f.write(token.strip() + "\n")
    logger.info(f"Token count: {len(globals.token_list)}, Error token count: {len(globals.error_token_list)}")
    tokens_count = len(set(globals.token_list) - set(globals.error_token_list))
    return {"status": "success", "tokens_count": tokens_count}


@app.post(f"/{api_prefix}/seed_tokens/clear" if api_prefix else "/seed_tokens/clear")
async def clear_seed_tokens():
    globals.seed_map.clear()
    globals.conversation_map.clear()
    with open(globals.SEED_MAP_FILE, "w", encoding="utf-8") as f:
        f.write("{}")
    with open(globals.CONVERSATION_MAP_FILE, "w", encoding="utf-8") as f:
        f.write("{}")
    logger.info(f"Seed token count: {len(globals.seed_map)}")
    return {"status": "success", "seed_tokens_count": len(globals.seed_map)}
````

## File: api/files.py
````python
import io

import pybase64
from PIL import Image

from utils.Client import Client
from utils.configs import export_proxy_url, cf_file_url


async def get_file_content(url):
    if url.startswith("data:"):
        mime_type, base64_data = url.split(';')[0].split(':')[1], url.split(',')[1]
        file_content = pybase64.b64decode(base64_data)
        return file_content, mime_type
    else:
        client = Client()
        try:
            if cf_file_url:
                body = {"file_url": url}
                r = await client.post(cf_file_url, timeout=60, json=body)
            else:
                r = await client.get(url, proxy=export_proxy_url, timeout=60)
            if r.status_code != 200:
                return None, None
            file_content = r.content
            mime_type = r.headers.get('Content-Type', '').split(';')[0].strip()
            return file_content, mime_type
        finally:
            await client.close()
            del client


async def determine_file_use_case(mime_type):
    multimodal_types = ["image/jpeg", "image/webp", "image/png", "image/gif"]
    my_files_types = ["text/x-php", "application/msword", "text/x-c", "text/html",
                      "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
                      "application/json", "text/javascript", "application/pdf",
                      "text/x-java", "text/x-tex", "text/x-typescript", "text/x-sh",
                      "text/x-csharp", "application/vnd.openxmlformats-officedocument.presentationml.presentation",
                      "text/x-c++", "application/x-latext", "text/markdown", "text/plain",
                      "text/x-ruby", "text/x-script.python"]

    if mime_type in multimodal_types:
        return "multimodal"
    elif mime_type in my_files_types:
        return "my_files"
    else:
        return "ace_upload"


async def get_image_size(file_content):
    with Image.open(io.BytesIO(file_content)) as img:
        return img.width, img.height


async def get_file_extension(mime_type):
    extension_mapping = {
        "image/jpeg": ".jpg",
        "image/png": ".png",
        "image/gif": ".gif",
        "image/webp": ".webp",
        "text/x-php": ".php",
        "application/msword": ".doc",
        "text/x-c": ".c",
        "text/html": ".html",
        "application/vnd.openxmlformats-officedocument.wordprocessingml.document": ".docx",
        "application/json": ".json",
        "text/javascript": ".js",
        "application/pdf": ".pdf",
        "text/x-java": ".java",
        "text/x-tex": ".tex",
        "text/x-typescript": ".ts",
        "text/x-sh": ".sh",
        "text/x-csharp": ".cs",
        "application/vnd.openxmlformats-officedocument.presentationml.presentation": ".pptx",
        "text/x-c++": ".cpp",
        "application/x-latex": ".latex",
        "text/markdown": ".md",
        "text/plain": ".txt",
        "text/x-ruby": ".rb",
        "text/x-script.python": ".py",
        "application/zip": ".zip",
        "application/x-zip-compressed": ".zip",
        "application/x-tar": ".tar",
        "application/x-compressed-tar": ".tar.gz",
        "application/vnd.rar": ".rar",
        "application/x-rar-compressed": ".rar",
        "application/x-7z-compressed": ".7z",
        "application/octet-stream": ".bin",
        "audio/mpeg": ".mp3",
        "audio/wav": ".wav",
        "audio/ogg": ".ogg",
        "audio/aac": ".aac",
        "video/mp4": ".mp4",
        "video/x-msvideo": ".avi",
        "video/x-matroska": ".mkv",
        "video/webm": ".webm",
        "application/rtf": ".rtf",
        "application/vnd.ms-excel": ".xls",
        "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet": ".xlsx",
        "text/css": ".css",
        "text/xml": ".xml",
        "application/xml": ".xml",
        "application/vnd.android.package-archive": ".apk",
        "application/vnd.apple.installer+xml": ".mpkg",
        "application/x-bzip": ".bz",
        "application/x-bzip2": ".bz2",
        "application/x-csh": ".csh",
        "application/x-debian-package": ".deb",
        "application/x-dvi": ".dvi",
        "application/java-archive": ".jar",
        "application/x-java-jnlp-file": ".jnlp",
        "application/vnd.mozilla.xul+xml": ".xul",
        "application/vnd.ms-fontobject": ".eot",
        "application/ogg": ".ogx",
        "application/x-font-ttf": ".ttf",
        "application/font-woff": ".woff",
        "application/x-shockwave-flash": ".swf",
        "application/vnd.visio": ".vsd",
        "application/xhtml+xml": ".xhtml",
        "application/vnd.ms-powerpoint": ".ppt",
        "application/vnd.oasis.opendocument.text": ".odt",
        "application/vnd.oasis.opendocument.spreadsheet": ".ods",
        "application/x-xpinstall": ".xpi",
        "application/vnd.google-earth.kml+xml": ".kml",
        "application/vnd.google-earth.kmz": ".kmz",
        "application/x-font-otf": ".otf",
        "application/vnd.ms-excel.addin.macroEnabled.12": ".xlam",
        "application/vnd.ms-excel.sheet.binary.macroEnabled.12": ".xlsb",
        "application/vnd.ms-excel.template.macroEnabled.12": ".xltm",
        "application/vnd.ms-powerpoint.addin.macroEnabled.12": ".ppam",
        "application/vnd.ms-powerpoint.presentation.macroEnabled.12": ".pptm",
        "application/vnd.ms-powerpoint.slideshow.macroEnabled.12": ".ppsm",
        "application/vnd.ms-powerpoint.template.macroEnabled.12": ".potm",
        "application/vnd.ms-word.document.macroEnabled.12": ".docm",
        "application/vnd.ms-word.template.macroEnabled.12": ".dotm",
        "application/x-ms-application": ".application",
        "application/x-ms-wmd": ".wmd",
        "application/x-ms-wmz": ".wmz",
        "application/x-ms-xbap": ".xbap",
        "application/vnd.ms-xpsdocument": ".xps",
        "application/x-silverlight-app": ".xap"
    }
    return extension_mapping.get(mime_type, "")
````

## File: api/models.py
````python
model_proxy = {
    "gpt-3.5-turbo": "gpt-3.5-turbo-0125",
    "gpt-3.5-turbo-16k": "gpt-3.5-turbo-16k-0613",
    "gpt-4": "gpt-4-0613",
    "gpt-4-32k": "gpt-4-32k-0613",
    "gpt-4-turbo-preview": "gpt-4-0125-preview",
    "gpt-4-vision-preview": "gpt-4-1106-vision-preview",
    "gpt-4-turbo": "gpt-4-turbo-2024-04-09",
    "gpt-4o": "gpt-4o-2024-08-06",
    "gpt-4o-mini": "gpt-4o-mini-2024-07-18",
    "o1-preview": "o1-preview-2024-09-12",
    "o1-mini": "o1-mini-2024-09-12",
    "o1": "o1-2024-12-18",
    "o3-mini": "o3-mini-2025-01-31",
    "o3-mini-high": "o3-mini-high-2025-01-31",
    "claude-3-opus": "claude-3-opus-20240229",
    "claude-3-sonnet": "claude-3-sonnet-20240229",
    "claude-3-haiku": "claude-3-haiku-20240307",
}

model_system_fingerprint = {
    "gpt-3.5-turbo-0125": ["fp_b28b39ffa8"],
    "gpt-3.5-turbo-1106": ["fp_592ef5907d"],
    "gpt-4-0125-preview": ["fp_f38f4d6482", "fp_2f57f81c11", "fp_a7daf7c51e", "fp_a865e8ede4", "fp_13c70b9f70",
                           "fp_b77cb481ed"],
    "gpt-4-1106-preview": ["fp_e467c31c3d", "fp_d986a8d1ba", "fp_99a5a401bb", "fp_123d5a9f90", "fp_0d1affc7a6",
                           "fp_5c95a4634e"],
    "gpt-4-turbo-2024-04-09": ["fp_d1bac968b4"],
    "gpt-4o-2024-05-13": ["fp_3aa7262c27"],
    "gpt-4o-mini-2024-07-18": ["fp_c9aa9c0491"]
}
````

## File: api/tokens.py
````python
import math

import tiktoken


async def calculate_image_tokens(width, height, detail):
    if detail == "low":
        return 85
    else:
        max_dimension = max(width, height)
        if max_dimension > 2048:
            scale_factor = 2048 / max_dimension
            new_width = int(width * scale_factor)
            new_height = int(height * scale_factor)
        else:
            new_width = width
            new_height = height

        width, height = new_width, new_height
        min_dimension = min(width, height)
        if min_dimension > 768:
            scale_factor = 768 / min_dimension
            new_width = int(width * scale_factor)
            new_height = int(height * scale_factor)
        else:
            new_width = width
            new_height = height

        width, height = new_width, new_height
        num_masks_w = math.ceil(width / 512)
        num_masks_h = math.ceil(height / 512)
        total_masks = num_masks_w * num_masks_h

        tokens_per_mask = 170
        total_tokens = total_masks * tokens_per_mask + 85

        return total_tokens


async def num_tokens_from_messages(messages, model=''):
    try:
        encoding = tiktoken.encoding_for_model(model)
    except KeyError:
        encoding = tiktoken.get_encoding("cl100k_base")
    if model == "gpt-3.5-turbo-0301":
        tokens_per_message = 4
    else:
        tokens_per_message = 3
    num_tokens = 0
    for message in messages:
        num_tokens += tokens_per_message
        for key, value in message.items():
            if isinstance(value, list):
                for item in value:
                    if item.get("type") == "text":
                        num_tokens += len(encoding.encode(item.get("text")))
                    if item.get("type") == "image_url":
                        pass
            else:
                num_tokens += len(encoding.encode(value))
    num_tokens += 3
    return num_tokens


async def num_tokens_from_content(content, model=None):
    try:
        encoding = tiktoken.encoding_for_model(model)
    except KeyError:
        encoding = tiktoken.get_encoding("cl100k_base")
    encoded_content = encoding.encode(content)
    len_encoded_content = len(encoded_content)
    return len_encoded_content


async def split_tokens_from_content(content, max_tokens, model=None):
    try:
        encoding = tiktoken.encoding_for_model(model)
    except KeyError:
        encoding = tiktoken.get_encoding("cl100k_base")
    encoded_content = encoding.encode(content)
    len_encoded_content = len(encoded_content)
    if len_encoded_content >= max_tokens:
        content = encoding.decode(encoded_content[:max_tokens])
        return content, max_tokens, "length"
    else:
        return content, len_encoded_content, "stop"
````

## File: app.py
````python
import warnings

import uvicorn
from fastapi import FastAPI, HTTPException
from fastapi.security import HTTPBearer, HTTPAuthorizationCredentials
from fastapi.middleware.cors import CORSMiddleware
from fastapi.templating import Jinja2Templates

from utils.configs import enable_gateway, api_prefix

warnings.filterwarnings("ignore")


log_config = uvicorn.config.LOGGING_CONFIG
default_format = "%(asctime)s | %(levelname)s | %(message)s"
access_format = r'%(asctime)s | %(levelname)s | %(client_addr)s: %(request_line)s %(status_code)s'
log_config["formatters"]["default"]["fmt"] = default_format
log_config["formatters"]["access"]["fmt"] = access_format

app = FastAPI(
    docs_url=f"/{api_prefix}/docs",    # 设置 Swagger UI 文档路径
    redoc_url=f"/{api_prefix}/redoc",  # 设置 Redoc 文档路径
    openapi_url=f"/{api_prefix}/openapi.json"  # 设置 OpenAPI JSON 路径
)

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

templates = Jinja2Templates(directory="templates")
security_scheme = HTTPBearer()

from app import app

import api.chat2api

if enable_gateway:
    import gateway.share
    import gateway.login
    import gateway.chatgpt
    import gateway.gpts
    import gateway.admin
    import gateway.v1
    import gateway.backend
else:
    @app.api_route("/{path:path}", methods=["GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD", "PATCH", "TRACE"])
    async def reverse_proxy():
        raise HTTPException(status_code=404, detail="Gateway is disabled")


if __name__ == "__main__":
    uvicorn.run("app:app", host="0.0.0.0", port=5005)
    # uvicorn.run("app:app", host="0.0.0.0", port=5005, ssl_keyfile="key.pem", ssl_certfile="cert.pem")
````

## File: chatgpt/authorization.py
````python
import asyncio
import json
import random

from fastapi import HTTPException

import utils.configs as configs
import utils.globals as globals
from chatgpt.refreshToken import rt2ac
from utils.Logger import logger


def get_req_token(req_token, seed=None):
    if configs.auto_seed:
        available_token_list = list(set(globals.token_list) - set(globals.error_token_list))
        length = len(available_token_list)
        if seed and length > 0:
            if seed not in globals.seed_map.keys():
                globals.seed_map[seed] = {"token": random.choice(available_token_list), "conversations": []}
                with open(globals.SEED_MAP_FILE, "w") as f:
                    json.dump(globals.seed_map, f, indent=4)
            else:
                req_token = globals.seed_map[seed]["token"]
            return req_token

        if req_token in configs.authorization_list:
            if len(available_token_list) > 0:
                if configs.random_token:
                    req_token = random.choice(available_token_list)
                    return req_token
                else:
                    globals.count += 1
                    globals.count %= length
                    return available_token_list[globals.count]
            else:
                return ""
        else:
            return req_token
    else:
        seed = req_token
        if seed not in globals.seed_map.keys():
            raise HTTPException(status_code=401, detail={"error": "Invalid Seed"})
        return globals.seed_map[seed]["token"]


async def verify_token(req_token):
    if not req_token:
        if configs.authorization_list:
            logger.error("Unauthorized with empty token.")
            raise HTTPException(status_code=401)
        else:
            return None
    else:
        if req_token.startswith("eyJhbGciOi") or req_token.startswith("fk-"):
            access_token = req_token
            return access_token
        elif len(req_token) == 45:
            try:
                if req_token in globals.error_token_list:
                    raise HTTPException(status_code=401, detail="Error RefreshToken")

                access_token = await rt2ac(req_token, force_refresh=False)
                return access_token
            except HTTPException as e:
                raise HTTPException(status_code=e.status_code, detail=e.detail)
        else:
            return req_token


async def refresh_all_tokens(force_refresh=False):
    for token in list(set(globals.token_list) - set(globals.error_token_list)):
        if len(token) == 45:
            try:
                await asyncio.sleep(0.5)
                await rt2ac(token, force_refresh=force_refresh)
            except HTTPException:
                pass
    logger.info("All tokens refreshed.")
````

## File: chatgpt/chatFormat_v1.py
````python
import asyncio
import json
import random
import re
import string
import time
import uuid

import pybase64
import websockets
from fastapi import HTTPException

from api.files import get_file_content
from api.models import model_system_fingerprint
from api.tokens import split_tokens_from_content, calculate_image_tokens, num_tokens_from_messages
from utils.Logger import logger

moderation_message = "I'm sorry, I cannot provide or engage in any content related to pornography, violence, or any unethical material. If you have any other questions or need assistance, please feel free to let me know. I'll do my best to provide support and assistance."


async def format_not_stream_response(response, prompt_tokens, max_tokens, model):
    chat_id = f"chatcmpl-{''.join(random.choice(string.ascii_letters + string.digits) for _ in range(29))}"
    system_fingerprint_list = model_system_fingerprint.get(model, None)
    system_fingerprint = random.choice(system_fingerprint_list) if system_fingerprint_list else None
    created_time = int(time.time())
    all_text = ""
    async for chunk in response:
        try:
            if chunk.startswith("data: [DONE]"):
                break
            elif not chunk.startswith("data: "):
                continue
            else:
                chunk = json.loads(chunk[6:])
                if not chunk["choices"][0].get("delta"):
                    continue
                all_text += chunk["choices"][0]["delta"]["content"]
        except Exception as e:
            logger.error(f"Error: {chunk}, error: {str(e)}")
            continue
    content, completion_tokens, finish_reason = await split_tokens_from_content(all_text, max_tokens, model)
    message = {
        "role": "assistant",
        "content": content,
    }
    usage = {
        "prompt_tokens": prompt_tokens,
        "completion_tokens": completion_tokens,
        "total_tokens": prompt_tokens + completion_tokens
    }
    if not message.get("content"):
        raise HTTPException(status_code=403, detail="No content in the message.")

    data = {
        "id": chat_id,
        "object": "chat.completion",
        "created": created_time,
        "model": model,
        "choices": [
            {
                "index": 0,
                "message": message,
                "logprobs": None,
                "finish_reason": finish_reason
            }
        ],
        "usage": usage
    }
    if system_fingerprint:
        data["system_fingerprint"] = system_fingerprint
    return data


async def head_process_response(response):
    async for chunk in response:
        chunk = chunk.decode("utf-8")
        if chunk.startswith("data: {"):
            chunk_old_data = json.loads(chunk[6:])
            message = chunk_old_data.get("message", {})
            if not message and "error" in chunk_old_data:
                return response, False
            role = message.get('author', {}).get('role')
            if role == 'user' or role == 'system':
                continue

            status = message.get("status")
            if status == "in_progress":
                return response, True
    return response, False


async def stream_response(service, response, model, max_tokens):
    chat_id = f"chatcmpl-{''.join(random.choice(string.ascii_letters + string.digits) for _ in range(29))}"
    created_time = int(time.time())

    chunk_new_data = {
        "id": chat_id,
        "object": "chat.completion.chunk",
        "created": created_time,
        "model": model,
        "choices": [
            {
                "index": 0,
                "delta": {"role": "assistant", "content": ""},
                "logprobs": None,
                "finish_reason": None
            }
        ]
    }
    yield f"data: {json.dumps(chunk_new_data)}\n\n"

    async for chunk in response:
        chunk = chunk.decode("utf-8")
        try:
            if chunk.startswith("data: {"):
                chunk_old_data = json.loads(chunk[6:].strip())
        except Exception as e:
            logger.error(f"Error: {chunk}, error: {str(e)}")
            continue


def get_url_from_content(content):
    if isinstance(content, str) and content.startswith('http'):
        try:
            url = re.match(
                r'(?i)\b((?:[a-z][\w-]+:(?:/{1,3}|[a-z0-9%])|www\d{0,3}[.]|[a-z0-9.\-]+[.][a-z]{2,4}/)(?:[^\s()<>]+|\(([^\s()<>]+|(\([^\s()<>]+\)))*\))+(?:\(([^\s()<>]+|(\([^\s()<>]+\)))*\)|[^\s`!()\[\]{};:\'".,<>?«»“”‘’]))',
                content.split(' ')[0])[0]
            content = content.replace(url, '').strip()
            return url, content
        except Exception:
            return None, content
    return None, content


def format_messages_with_url(content):
    url_list = []
    while True:
        url, content = get_url_from_content(content)
        if url:
            url_list.append(url)
            logger.info(f"Found a file_url from messages: {url}")
        else:
            break
    if not url_list:
        return content
    new_content = [
        {
            "type": "text",
            "text": content
        }
    ]
    for url in url_list:
        new_content.append({
            "type": "image_url",
            "image_url": {
                "url": url
            }
        })
    return new_content


async def api_messages_to_chat(service, api_messages, upload_by_url=False):
    file_tokens = 0
    chat_messages = []
    for api_message in api_messages:
        role = api_message.get('role')
        content = api_message.get('content')
        if upload_by_url:
            if isinstance(content, str):
                content = format_messages_with_url(content)
        if isinstance(content, list):
            parts = []
            attachments = []
            content_type = "multimodal_text"
            for i in content:
                if i.get("type") == "text":
                    parts.append(i.get("text"))
                elif i.get("type") == "image_url":
                    image_url = i.get("image_url")
                    url = image_url.get("url")
                    detail = image_url.get("detail", "auto")
                    file_content, mime_type = await get_file_content(url)
                    file_meta = await service.upload_file(file_content, mime_type)
                    if file_meta:
                        file_id = file_meta["file_id"]
                        file_size = file_meta["size_bytes"]
                        file_name = file_meta["file_name"]
                        mime_type = file_meta["mime_type"]
                        use_case = file_meta["use_case"]
                        if mime_type.startswith("image/"):
                            width, height = file_meta["width"], file_meta["height"]
                            file_tokens += await calculate_image_tokens(width, height, detail)
                            parts.append({
                                "content_type": "image_asset_pointer",
                                "asset_pointer": f"file-service://{file_id}",
                                "size_bytes": file_size,
                                "width": width,
                                "height": height
                            })
                            attachments.append({
                                "id": file_id,
                                "size": file_size,
                                "name": file_name,
                                "mime_type": mime_type,
                                "width": width,
                                "height": height
                            })
                        else:
                            if not use_case == "ace_upload":
                                await service.check_upload(file_id)
                            file_tokens += file_size // 1000
                            attachments.append({
                                "id": file_id,
                                "size": file_size,
                                "name": file_name,
                                "mime_type": mime_type,
                            })
            metadata = {
                "attachments": attachments
            }
        else:
            content_type = "text"
            parts = [content]
            metadata = {}
        chat_message = {
            "id": f"{uuid.uuid4()}",
            "author": {"role": role},
            "content": {"content_type": content_type, "parts": parts},
            "metadata": metadata
        }
        chat_messages.append(chat_message)
    text_tokens = await num_tokens_from_messages(api_messages, service.resp_model)
    prompt_tokens = text_tokens + file_tokens
    return chat_messages, prompt_tokens
````

## File: chatgpt/chatFormat.py
````python
import asyncio
import json
import random
import re
import string
import time
import uuid

import pybase64
import websockets
from fastapi import HTTPException

from api.files import get_file_content
from api.models import model_system_fingerprint
from api.tokens import split_tokens_from_content, calculate_image_tokens, num_tokens_from_messages
from utils.Logger import logger

moderation_message = "I'm sorry, I cannot provide or engage in any content related to pornography, violence, or any unethical material. If you have any other questions or need assistance, please feel free to let me know. I'll do my best to provide support and assistance."


async def format_not_stream_response(response, prompt_tokens, max_tokens, model):
    chat_id = f"chatcmpl-{''.join(random.choice(string.ascii_letters + string.digits) for _ in range(29))}"
    system_fingerprint_list = model_system_fingerprint.get(model, None)
    system_fingerprint = random.choice(system_fingerprint_list) if system_fingerprint_list else None
    created_time = int(time.time())
    all_text = ""
    async for chunk in response:
        try:
            if chunk.startswith("data: [DONE]"):
                break
            elif not chunk.startswith("data: "):
                continue
            else:
                chunk = json.loads(chunk[6:])
                if not chunk["choices"][0].get("delta"):
                    continue
                all_text += chunk["choices"][0]["delta"]["content"]
        except Exception as e:
            logger.error(f"Error: {chunk}, error: {str(e)}")
            continue
    content, completion_tokens, finish_reason = await split_tokens_from_content(all_text, max_tokens, model)
    message = {
        "role": "assistant",
        "content": content,
    }
    usage = {
        "prompt_tokens": prompt_tokens,
        "completion_tokens": completion_tokens,
        "total_tokens": prompt_tokens + completion_tokens
    }
    if not message.get("content"):
        raise HTTPException(status_code=403, detail="No content in the message.")

    data = {
        "id": chat_id,
        "object": "chat.completion",
        "created": created_time,
        "model": model,
        "choices": [
            {
                "index": 0,
                "message": message,
                "logprobs": None,
                "finish_reason": finish_reason
            }
        ],
        "usage": usage
    }
    if system_fingerprint:
        data["system_fingerprint"] = system_fingerprint
    return data


async def wss_stream_response(websocket, conversation_id):
    while not websocket.closed:
        try:
            message = await asyncio.wait_for(websocket.recv(), timeout=10)
            if message:
                resultObj = json.loads(message)
                sequenceId = resultObj.get("sequenceId", None)
                if not sequenceId:
                    continue
                data = resultObj.get("data", {})
                if conversation_id != data.get("conversation_id", ""):
                    continue
                sequenceId = resultObj.get('sequenceId')
                if sequenceId and sequenceId % 80 == 0:
                    await websocket.send(
                        json.dumps(
                            {"type": "sequenceAck", "sequenceId": sequenceId}
                        )
                    )
                decoded_bytes = pybase64.b64decode(data.get("body", None))
                yield decoded_bytes
            else:
                print("No message received within the specified time.")
        except asyncio.TimeoutError:
            logger.error("Timeout! No message received within the specified time.")
            break
        except websockets.ConnectionClosed as e:
            if e.code == 1000:
                logger.error("WebSocket closed normally with code 1000 (OK)")
                yield b"data: [DONE]\n\n"
            else:
                logger.error(f"WebSocket closed with error code {e.code}")
        except Exception as e:
            logger.error(f"Error: {str(e)}")
            continue


async def head_process_response(response):
    async for chunk in response:
        chunk = chunk.decode("utf-8")
        if chunk.startswith("data: {"):
            chunk_old_data = json.loads(chunk[6:])
            message = chunk_old_data.get("message", {})
            if not message and "error" in chunk_old_data:
                return response, False
            role = message.get('author', {}).get('role')
            if role == 'user' or role == 'system':
                continue

            status = message.get("status")
            if status == "in_progress":
                return response, True
    return response, False


async def stream_response(service, response, model, max_tokens):
    chat_id = f"chatcmpl-{''.join(random.choice(string.ascii_letters + string.digits) for _ in range(29))}"
    system_fingerprint_list = model_system_fingerprint.get(model, None)
    system_fingerprint = random.choice(system_fingerprint_list) if system_fingerprint_list else None
    created_time = int(time.time())
    completion_tokens = 0
    len_last_content = 0
    len_last_citation = 0
    last_message_id = None
    last_role = None
    last_content_type = None
    last_status = None
    model_slug = None
    end = False

    chunk_new_data = {
        "id": chat_id,
        "object": "chat.completion.chunk",
        "created": created_time,
        "model": model,
        "choices": [
            {
                "index": 0,
                "delta": {"role": "assistant", "content": ""},
                "logprobs": None,
                "finish_reason": None
            }
        ]
    }
    if system_fingerprint:
        chunk_new_data["system_fingerprint"] = system_fingerprint
    yield f"data: {json.dumps(chunk_new_data)}\n\n"

    async for chunk in response:
        chunk = chunk.decode("utf-8")
        if end:
            logger.info(f"Response Model: {model_slug}")
            yield "data: [DONE]\n\n"
            break
        try:
            if chunk.startswith("data: {"):
                chunk_old_data = json.loads(chunk[6:])
                finish_reason = None
                message = chunk_old_data.get("message", {})
                conversation_id = chunk_old_data.get("conversation_id")
                role = message.get('author', {}).get('role')
                if role == 'user' or role == 'system':
                    continue

                status = message.get("status")
                message_id = message.get("id")
                content = message.get("content", {})
                recipient = message.get("recipient", "")
                meta_data = message.get("metadata", {})
                initial_text = meta_data.get("initial_text", "")
                model_slug = meta_data.get("model_slug", model_slug)

                if not message and chunk_old_data.get("type") == "moderation":
                    delta = {"role": "assistant", "content": moderation_message}
                    finish_reason = "stop"
                    end = True
                elif status == "in_progress":
                    outer_content_type = content.get("content_type")
                    if outer_content_type == "text":
                        part = content.get("parts", [])[0]
                        if not part:
                            if role == 'assistant' and last_role != 'assistant':
                                if last_role == None:
                                    new_text = ""
                                else:
                                    new_text = f"\n"
                            elif role == 'tool' and last_role != 'tool':
                                new_text = f">{initial_text}\n"
                            else:
                                new_text = ""
                        else:
                            if last_message_id and last_message_id != message_id:
                                continue
                            citation = message.get("metadata", {}).get("citations", [])
                            if len(citation) > len_last_citation:
                                inside_metadata = citation[-1].get("metadata", {})
                                citation_title = inside_metadata.get("title", "")
                                citation_url = inside_metadata.get("url", "")
                                new_text = f' **[[""]]({citation_url} "{citation_title}")** '
                                len_last_citation = len(citation)
                            else:
                                if role == 'assistant' and last_role != 'assistant':
                                    if recipient == 'dalle.text2im':
                                        new_text = f"\n```{recipient}\n{part[len_last_content:]}"
                                    elif recipient == 't2uay3k.sj1i4kz':
                                        new_text = f"\n```image_creator\n{part[len_last_content:]}"
                                    elif last_role == None:
                                        new_text = part[len_last_content:]
                                    else:
                                        new_text = f"\n\n{part[len_last_content:]}"
                                elif role == 'tool' and last_role != 'tool':
                                    new_text = f">{initial_text}\n{part[len_last_content:]}"
                                elif role == 'tool':
                                    new_text = part[len_last_content:].replace("\n\n", "\n")
                                else:
                                    new_text = part[len_last_content:]
                            len_last_content = len(part)
                    elif outer_content_type == "multimodal_text":
                        parts = content.get("parts", [])
                        new_text = ""
                        for part in parts:
                            file_id = part.get('asset_pointer').replace('sediment://', '')
                            full_height = part.get("height", 0)
                            current_height = part.get('metadata', {}).get("generation", {}).get("height", 0)
                            if full_height > current_height:
                                completed_rate = current_height / full_height
                                new_text = f"\n> {completed_rate:.2%}\n"
                                if last_role != role:
                                    new_text = f"\n```{new_text}"
                            else:
                                image_download_url = await service.get_attachment_url(file_id, conversation_id)
                                new_text = f"\n```\n![image]({image_download_url})\n"
                    else:
                        text = content.get("text", "")
                        if outer_content_type == "code" and last_content_type != "code":
                            language = content.get("language", "")
                            if not language or language == "unknown":
                                language = recipient
                            new_text = "\n```" + language + "\n" + text[len_last_content:]
                        elif outer_content_type == "execution_output" and last_content_type != "execution_output":
                            new_text = "\n```" + "Output" + "\n" + text[len_last_content:]
                        else:
                            new_text = text[len_last_content:]
                        len_last_content = len(text)
                    if last_content_type == "code" and outer_content_type != "code":
                        new_text = "\n```\n" + new_text
                    elif last_content_type == "execution_output" and outer_content_type != "execution_output":
                        new_text = "\n```\n" + new_text
                    elif last_content_type == "multimodal_text" and outer_content_type != "multimodal_text":
                        new_text = "\n```\n" + new_text

                    delta = {"content": new_text}
                    last_content_type = outer_content_type
                    if completion_tokens >= max_tokens:
                        delta = {}
                        finish_reason = "length"
                        end = True

                elif status == "finished_successfully":
                    if content.get("content_type") == "multimodal_text":
                        parts = content.get("parts", [])
                        delta = {}
                        for part in parts:
                            if isinstance(part, str):
                                continue
                            inner_content_type = part.get('content_type')
                            if inner_content_type == "image_asset_pointer":
                                last_content_type = "image_asset_pointer"
                                if part.get('asset_pointer').startswith('file-service://'):
                                    file_id = part.get('asset_pointer').replace('file-service://', '')
                                    logger.debug(f"file_id: {file_id}")
                                    image_download_url = await service.get_download_url(file_id)
                                    logger.debug(f"image_download_url: {image_download_url}")
                                    if image_download_url:
                                        delta = {"content": f"\n```\n![image]({image_download_url})\n"}
                                    else:
                                        delta = {"content": f"\n```\nFailed to load the image.\n"}
                                else:
                                    file_id = part.get('asset_pointer').replace('sediment://', '')
                                    image_download_url = await service.get_attachment_url(file_id, conversation_id)
                                    delta = {"content": f"\n![image]({image_download_url})\n"}
                    elif message.get("end_turn"):
                        part = content.get("parts", [])[0]
                        new_text = part[len_last_content:]
                        if not new_text:
                            matches = re.findall(r'\(sandbox:(.*?)\)', part)
                            if matches:
                                file_url_content = ""
                                for i, sandbox_path in enumerate(matches):
                                    file_download_url = await service.get_response_file_url(conversation_id, message_id, sandbox_path)
                                    if file_download_url:
                                        file_url_content += f"\n```\n\n![File {i+1}]({file_download_url})\n"
                                delta = {"content": file_url_content}
                            else:
                                delta = {}
                        else:
                            delta = {"content": new_text}
                        finish_reason = "stop"
                        end = True
                    else:
                        len_last_content = 0
                        if meta_data.get("finished_text"):
                            delta = {"content": f"\n{meta_data.get('finished_text')}\n"}
                        else:
                            continue
                else:
                    continue
                last_message_id = message_id
                last_role = role
                last_status = status
                if not end and not delta.get("content"):
                    delta = {"role": "assistant", "content": ""}
                chunk_new_data["choices"][0]["delta"] = delta
                chunk_new_data["choices"][0]["finish_reason"] = finish_reason
                if not service.history_disabled:
                    chunk_new_data.update({
                        "message_id": message_id,
                        "conversation_id": conversation_id,
                    })
                completion_tokens += 1
                yield f"data: {json.dumps(chunk_new_data)}\n\n"
            elif chunk.startswith("data: [DONE]"):
                logger.info(f"Response Model: {model_slug}")
                yield "data: [DONE]\n\n"
            else:
                continue
        except Exception as e:
            if chunk.startswith("data: "):
                chunk_data = json.loads(chunk[6:])
                if chunk_data.get("error"):
                    logger.error(f"Error: {chunk_data.get('error')}")
                    yield "data: [DONE]\n\n"
                    break
            logger.error(f"Error: {chunk}, details: {str(e)}")
            continue


def get_url_from_content(content):
    if isinstance(content, str) and content.startswith('http'):
        try:
            url = re.match(
                r'(?i)\b((?:[a-z][\w-]+:(?:/{1,3}|[a-z0-9%])|www\d{0,3}[.]|[a-z0-9.\-]+[.][a-z]{2,4}/)(?:[^\s()<>]+|\(([^\s()<>]+|(\([^\s()<>]+\)))*\))+(?:\(([^\s()<>]+|(\([^\s()<>]+\)))*\)|[^\s`!()\[\]{};:\'".,<>?«»“”‘’]))',
                content.split(' ')[0])[0]
            content = content.replace(url, '').strip()
            return url, content
        except Exception:
            return None, content
    return None, content


def format_messages_with_url(content):
    url_list = []
    while True:
        url, content = get_url_from_content(content)
        if url:
            url_list.append(url)
            logger.info(f"Found a file_url from messages: {url}")
        else:
            break
    if not url_list:
        return content
    new_content = [
        {
            "type": "text",
            "text": content
        }
    ]
    for url in url_list:
        new_content.append({
            "type": "image_url",
            "image_url": {
                "url": url
            }
        })
    return new_content


async def api_messages_to_chat(service, api_messages, upload_by_url=False):
    file_tokens = 0
    chat_messages = []
    for api_message in api_messages:
        role = api_message.get('role')
        content = api_message.get('content')
        if upload_by_url:
            if isinstance(content, str):
                content = format_messages_with_url(content)
        if isinstance(content, list):
            parts = []
            attachments = []
            content_type = "multimodal_text"
            for i in content:
                if i.get("type") == "text":
                    parts.append(i.get("text"))
                elif i.get("type") == "image_url":
                    image_url = i.get("image_url")
                    url = image_url.get("url")
                    detail = image_url.get("detail", "auto")
                    file_content, mime_type = await get_file_content(url)
                    file_meta = await service.upload_file(file_content, mime_type)
                    if file_meta:
                        file_id = file_meta["file_id"]
                        file_size = file_meta["size_bytes"]
                        file_name = file_meta["file_name"]
                        mime_type = file_meta["mime_type"]
                        use_case = file_meta["use_case"]
                        if mime_type.startswith("image/"):
                            width, height = file_meta["width"], file_meta["height"]
                            file_tokens += await calculate_image_tokens(width, height, detail)
                            parts.append({
                                "content_type": "image_asset_pointer",
                                "asset_pointer": f"file-service://{file_id}",
                                "size_bytes": file_size,
                                "width": width,
                                "height": height
                            })
                            attachments.append({
                                "id": file_id,
                                "size": file_size,
                                "name": file_name,
                                "mime_type": mime_type,
                                "width": width,
                                "height": height
                            })
                        else:
                            if not use_case == "ace_upload":
                                await service.check_upload(file_id)
                            file_tokens += file_size // 1000
                            attachments.append({
                                "id": file_id,
                                "size": file_size,
                                "name": file_name,
                                "mime_type": mime_type,
                            })
            metadata = {
                "attachments": attachments
            }
        else:
            content_type = "text"
            parts = [content]
            metadata = {}
        chat_message = {
            "id": f"{uuid.uuid4()}",
            "author": {"role": role},
            "content": {"content_type": content_type, "parts": parts},
            "metadata": metadata
        }
        chat_messages.append(chat_message)
    text_tokens = await num_tokens_from_messages(api_messages, service.resp_model)
    prompt_tokens = text_tokens + file_tokens
    return chat_messages, prompt_tokens
````

## File: chatgpt/chatLimit.py
````python
import time
from datetime import datetime

from utils.Logger import logger

limit_details = {}


def check_is_limit(detail, token, model):
    if token and isinstance(detail, dict) and detail.get('clears_in'):
        clear_time = int(time.time()) + detail.get('clears_in')
        limit_details.setdefault(token, {})[model] = clear_time
        logger.info(f"{token[:40]}: Reached {model} limit, will be cleared at {datetime.fromtimestamp(clear_time).replace(microsecond=0)}")


async def handle_request_limit(token, model):
    try:
        if limit_details.get(token) and model in limit_details[token]:
            limit_time = limit_details[token][model]
            is_limit = limit_time > int(time.time())
            if is_limit:
                clear_date = datetime.fromtimestamp(limit_time).replace(microsecond=0)
                result = f"Request limit exceeded. You can continue with the default model now, or try again after {clear_date}"
                logger.info(result)
                return result
            else:
                del limit_details[token][model]
                return None
    except KeyError as e:
        logger.error(f"Key error: {e}")
        return None
    except Exception as e:
        logger.error(f"An unexpected error occurred: {e}")
        return None
````

## File: chatgpt/ChatService.py
````python
import asyncio
import hashlib
import json
import random
import uuid

from fastapi import HTTPException
from starlette.concurrency import run_in_threadpool

from api.files import get_image_size, get_file_extension, determine_file_use_case
from api.models import model_proxy
from chatgpt.authorization import get_req_token, verify_token
from chatgpt.chatFormat import api_messages_to_chat, stream_response, format_not_stream_response, head_process_response
from chatgpt.chatLimit import check_is_limit, handle_request_limit
from chatgpt.fp import get_fp
from chatgpt.proofofWork import get_config, get_dpl, get_answer_token, get_requirements_token

from utils.Client import Client
from utils.Logger import logger
from utils.configs import (
    chatgpt_base_url_list,
    ark0se_token_url_list,
    sentinel_proxy_url_list,
    history_disabled,
    pow_difficulty,
    conversation_only,
    enable_limit,
    upload_by_url,
    auth_key,
    turnstile_solver_url,
    oai_language,
)


class ChatService:
    def __init__(self, origin_token=None):
        # self.user_agent = random.choice(user_agents_list) if user_agents_list else "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36"
        self.req_token = get_req_token(origin_token)
        self.chat_token = "gAAAAAB"
        self.s = None
        self.ss = None
        self.ws = None

    async def set_dynamic_data(self, data):
        if self.req_token:
            req_len = len(self.req_token.split(","))
            if req_len == 1:
                self.access_token = await verify_token(self.req_token)
                self.account_id = None
            else:
                self.access_token = await verify_token(self.req_token.split(",")[0])
                self.account_id = self.req_token.split(",")[1]
        else:
            logger.info("Request token is empty, use no-auth 3.5")
            self.access_token = None
            self.account_id = None

        self.fp = get_fp(self.req_token).copy()
        self.proxy_url = self.fp.pop("proxy_url", None)
        self.impersonate = self.fp.pop("impersonate", "safari15_3")
        self.user_agent = self.fp.get("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0")
        logger.info(f"Request token: {self.req_token}")
        logger.info(f"Request proxy: {self.proxy_url}")
        logger.info(f"Request UA: {self.user_agent}")
        logger.info(f"Request impersonate: {self.impersonate}")

        self.data = data
        await self.set_model()
        if enable_limit and self.req_token:
            limit_response = await handle_request_limit(self.req_token, self.req_model)
            if limit_response:
                raise HTTPException(status_code=429, detail=limit_response)

        self.account_id = self.data.get('Chatgpt-Account-Id', self.account_id)
        self.parent_message_id = self.data.get('parent_message_id')
        self.conversation_id = self.data.get('conversation_id')
        self.history_disabled = self.data.get('history_disabled', history_disabled)

        self.api_messages = self.data.get("messages", [])
        self.prompt_tokens = 0
        self.max_tokens = self.data.get("max_tokens", 2147483647)
        if not isinstance(self.max_tokens, int):
            self.max_tokens = 2147483647

        # self.proxy_url = random.choice(proxy_url_list) if proxy_url_list else None

        self.host_url = random.choice(chatgpt_base_url_list) if chatgpt_base_url_list else "https://chatgpt.com"
        self.ark0se_token_url = random.choice(ark0se_token_url_list) if ark0se_token_url_list else None

        session_id = hashlib.md5(self.req_token.encode()).hexdigest()
        proxy_url = self.proxy_url.replace("{}", session_id) if self.proxy_url else None
        self.s = Client(proxy=proxy_url, impersonate=self.impersonate)
        if sentinel_proxy_url_list:
            sentinel_proxy_url = (random.choice(sentinel_proxy_url_list)).replace("{}", session_id) if sentinel_proxy_url_list else None
            self.ss = Client(proxy=sentinel_proxy_url, impersonate=self.impersonate)
        else:
            self.ss = self.s

        self.persona = None
        self.ark0se_token = None
        self.proof_token = None
        self.turnstile_token = None

        self.chat_headers = None
        self.chat_request = None

        self.base_headers = {
            'accept': '*/*',
            'accept-encoding': 'gzip, deflate, br, zstd',
            'accept-language': 'en-US,en;q=0.9',
            'content-type': 'application/json',
            'oai-language': oai_language,
            'origin': self.host_url,
            'priority': 'u=1, i',
            'referer': f'{self.host_url}/',
            'sec-fetch-dest': 'empty',
            'sec-fetch-mode': 'cors',
            'sec-fetch-site': 'same-origin'
        }
        self.base_headers.update(self.fp)

        if self.access_token:
            self.base_url = self.host_url + "/backend-api"
            self.base_headers['authorization'] = f'Bearer {self.access_token}'
            if self.account_id:
                self.base_headers['chatgpt-account-id'] = self.account_id
        else:
            self.base_url = self.host_url + "/backend-anon"

        if auth_key:
            self.base_headers['authkey'] = auth_key

        await get_dpl(self)

    async def set_model(self):
        self.origin_model = self.data.get("model", "gpt-3.5-turbo-0125")
        self.resp_model = model_proxy.get(self.origin_model, self.origin_model)
        if "gizmo" in self.origin_model or "g-" in self.origin_model:
            self.gizmo_id = "g-" + self.origin_model.split("g-")[-1]
        else:
            self.gizmo_id = None

        if "o3-mini-high" in self.origin_model:
            self.req_model = "o3-mini-high"
        elif "o3-mini-medium" in self.origin_model:
            self.req_model = "o3-mini-medium"
        elif "o3-mini-low" in self.origin_model:
            self.req_model = "o3-mini-low"
        elif "o3-mini" in self.origin_model:
            self.req_model = "o3-mini"
        elif "o3" in self.origin_model:
            self.req_model = "o3"
        elif "o1-preview" in self.origin_model:
            self.req_model = "o1-preview"
        elif "o1-pro" in self.origin_model:
            self.req_model = "o1-pro"
        elif "o1-mini" in self.origin_model:
            self.req_model = "o1-mini"
        elif "o1" in self.origin_model:
            self.req_model = "o1"
        elif "gpt-4.5o" in self.origin_model:
            self.req_model = "gpt-4.5o"
        elif "gpt-4o-canmore" in self.origin_model:
            self.req_model = "gpt-4o-canmore"
        elif "gpt-4o-mini" in self.origin_model:
            self.req_model = "gpt-4o-mini"
        elif "gpt-4o" in self.origin_model:
            self.req_model = "gpt-4o"
        elif "gpt-4-mobile" in self.origin_model:
            self.req_model = "gpt-4-mobile"
        elif "gpt-4" in self.origin_model:
            self.req_model = "gpt-4"
        elif "gpt-3.5" in self.origin_model:
            self.req_model = "text-davinci-002-render-sha"
        elif "auto" in self.origin_model:
            self.req_model = "auto"
        else:
            self.req_model = "gpt-4o"

    async def get_chat_requirements(self):
        if conversation_only:
            return None
        url = f'{self.base_url}/sentinel/chat-requirements'
        headers = self.base_headers.copy()
        try:
            config = get_config(self.user_agent, self.req_token)
            p = get_requirements_token(config)
            data = {'p': p}
            r = await self.ss.post(url, headers=headers, json=data, timeout=5)
            if r.status_code == 200:
                resp = r.json()

                self.persona = resp.get("persona")
                if self.persona != "chatgpt-paid":
                    if self.req_model == "gpt-4" or self.req_model == "o1-preview":
                        logger.error(f"Model {self.resp_model} not support for {self.persona}")
                        raise HTTPException(
                            status_code=404,
                            detail={
                                "message": f"The model `{self.origin_model}` does not exist or you do not have access to it.",
                                "type": "invalid_request_error",
                                "param": None,
                                "code": "model_not_found",
                            },
                        )

                turnstile = resp.get('turnstile', {})
                turnstile_required = turnstile.get('required')
                if turnstile_required:
                    turnstile_dx = turnstile.get("dx")
                    try:
                        if turnstile_solver_url:
                            res = await self.s.post(
                                turnstile_solver_url, json={"url": "https://chatgpt.com", "p": p, "dx": turnstile_dx, "ua": self.user_agent}
                            )
                            self.turnstile_token = res.json().get("t")
                    except Exception as e:
                        logger.info(f"Turnstile ignored: {e}")
                    # raise HTTPException(status_code=403, detail="Turnstile required")

                ark0se = resp.get('ark' + 'ose', {})
                ark0se_required = ark0se.get('required')
                if ark0se_required:
                    if self.persona == "chatgpt-freeaccount":
                        ark0se_method = "chat35"
                    else:
                        ark0se_method = "chat4"
                    if not self.ark0se_token_url:
                        raise HTTPException(status_code=403, detail="Ark0se service required")
                    ark0se_dx = ark0se.get("dx")
                    ark0se_client = Client(impersonate=self.impersonate)
                    try:
                        r2 = await ark0se_client.post(
                            url=self.ark0se_token_url, json={"blob": ark0se_dx, "method": ark0se_method}, timeout=15
                        )
                        r2esp = r2.json()
                        logger.info(f"ark0se_token: {r2esp}")
                        if r2esp.get('solved', True):
                            self.ark0se_token = r2esp.get('token')
                        else:
                            raise HTTPException(status_code=403, detail="Failed to get Ark0se token")
                    except Exception:
                        raise HTTPException(status_code=403, detail="Failed to get Ark0se token")
                    finally:
                        await ark0se_client.close()

                proofofwork = resp.get('proofofwork', {})
                proofofwork_required = proofofwork.get('required')
                if proofofwork_required:
                    proofofwork_diff = proofofwork.get("difficulty")
                    if proofofwork_diff <= pow_difficulty:
                        raise HTTPException(status_code=403, detail=f"Proof of work difficulty too high: {proofofwork_diff}")
                    proofofwork_seed = proofofwork.get("seed")
                    self.proof_token, solved = await run_in_threadpool(
                        get_answer_token, proofofwork_seed, proofofwork_diff, config
                    )
                    if not solved:
                        raise HTTPException(status_code=403, detail="Failed to solve proof of work")

                self.chat_token = resp.get('token')
                if not self.chat_token:
                    raise HTTPException(status_code=403, detail=f"Failed to get chat token: {r.text}")
                return self.chat_token
            else:
                if "application/json" == r.headers.get("Content-Type", ""):
                    detail = r.json().get("detail", r.json())
                else:
                    detail = r.text
                if "cf_chl_opt" in detail:
                    raise HTTPException(status_code=r.status_code, detail="cf_chl_opt")
                if r.status_code == 429:
                    raise HTTPException(status_code=r.status_code, detail="rate-limit")
                raise HTTPException(status_code=r.status_code, detail=detail)
        except HTTPException as e:
            raise HTTPException(status_code=e.status_code, detail=e.detail)
        except Exception as e:
            raise HTTPException(status_code=500, detail=str(e))

    async def prepare_send_conversation(self):
        try:
            chat_messages, self.prompt_tokens = await api_messages_to_chat(self, self.api_messages, upload_by_url)
        except Exception as e:
            logger.error(f"Failed to format messages: {str(e)}")
            raise HTTPException(status_code=400, detail="Failed to format messages.")
        self.chat_headers = self.base_headers.copy()
        self.chat_headers.update(
            {
                'accept': 'text/event-stream',
                'openai-sentinel-chat-requirements-token': self.chat_token,
                'openai-sentinel-proof-token': self.proof_token,
            }
        )
        if self.ark0se_token:
            self.chat_headers['openai-sentinel-ark' + 'ose-token'] = self.ark0se_token

        if self.turnstile_token:
            self.chat_headers['openai-sentinel-turnstile-token'] = self.turnstile_token

        if conversation_only:
            self.chat_headers.pop('openai-sentinel-chat-requirements-token', None)
            self.chat_headers.pop('openai-sentinel-proof-token', None)
            self.chat_headers.pop('openai-sentinel-ark' + 'ose-token', None)
            self.chat_headers.pop('openai-sentinel-turnstile-token', None)

        if self.gizmo_id:
            conversation_mode = {"kind": "gizmo_interaction", "gizmo_id": self.gizmo_id}
            logger.info(f"Gizmo id: {self.gizmo_id}")
        else:
            conversation_mode = {"kind": "primary_assistant"}

        logger.info(f"Model mapping: {self.origin_model} -> {self.req_model}")
        self.chat_request = {
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
            "conversation_mode": conversation_mode,
            "conversation_origin": None,
            "force_paragen": False,
            "force_paragen_model_slug": "",
            "force_rate_limit": False,
            "force_use_sse": True,
            "history_and_training_disabled": self.history_disabled,
            "messages": chat_messages,
            "model": self.req_model,
            "paragen_cot_summary_display_override": "allow",
            "paragen_stream_type_override": None,
            "parent_message_id": self.parent_message_id if self.parent_message_id else f"{uuid.uuid4()}",
            "reset_rate_limits": False,
            "suggestions": [],
            "supported_encodings": [],
            "system_hints": [],
            "timezone": "America/Los_Angeles",
            "timezone_offset_min": -480,
            "variant_purpose": "comparison_implicit",
            "websocket_request_id": f"{uuid.uuid4()}",
        }
        if self.conversation_id:
            self.chat_request['conversation_id'] = self.conversation_id
        return self.chat_request

    async def send_conversation(self):
        try:
            url = f'{self.base_url}/conversation'
            stream = self.data.get("stream", False)
            r = await self.s.post_stream(url, headers=self.chat_headers, json=self.chat_request, timeout=10, stream=True)
            if r.status_code != 200:
                rtext = await r.atext()
                if "application/json" == r.headers.get("Content-Type", ""):
                    detail = json.loads(rtext).get("detail", json.loads(rtext))
                    if r.status_code == 429:
                        check_is_limit(detail, token=self.req_token, model=self.req_model)
                else:
                    if "cf_chl_opt" in rtext:
                        # logger.error(f"Failed to send conversation: cf_chl_opt")
                        raise HTTPException(status_code=r.status_code, detail="cf_chl_opt")
                    if r.status_code == 429:
                        # logger.error(f"Failed to send conversation: rate-limit")
                        raise HTTPException(status_code=r.status_code, detail="rate-limit")
                    detail = r.text[:100]
                # logger.error(f"Failed to send conversation: {detail}")
                raise HTTPException(status_code=r.status_code, detail=detail)

            content_type = r.headers.get("Content-Type", "")
            if "text/event-stream" in content_type:
                res, start = await head_process_response(r.aiter_lines())
                if not start:
                    raise HTTPException(
                        status_code=403,
                        detail="Our systems have detected unusual activity coming from your system. Please try again later.",
                    )
                if stream:
                    return stream_response(self, res, self.resp_model, self.max_tokens)
                else:
                    return await format_not_stream_response(
                        stream_response(self, res, self.resp_model, self.max_tokens),
                        self.prompt_tokens,
                        self.max_tokens,
                        self.resp_model,
                    )
            elif "application/json" in content_type:
                rtext = await r.atext()
                resp = json.loads(rtext)
                raise HTTPException(status_code=r.status_code, detail=resp)
            else:
                rtext = await r.atext()
                raise HTTPException(status_code=r.status_code, detail=rtext)
        except HTTPException as e:
            raise HTTPException(status_code=e.status_code, detail=e.detail)
        except Exception as e:
            raise HTTPException(status_code=500, detail=str(e))

    async def get_download_url(self, file_id):
        url = f"{self.base_url}/files/{file_id}/download"
        headers = self.base_headers.copy()
        try:
            r = await self.s.get(url, headers=headers, timeout=10)
            if r.status_code == 200:
                download_url = r.json().get('download_url')
                return download_url
            else:
                raise HTTPException(status_code=r.status_code, detail=r.text)
        except Exception as e:
            logger.error(f"Failed to get download url: {e}")
            return ""

    async def get_attachment_url(self, file_id, conversation_id):
        url = f"{self.base_url}/conversation/{conversation_id}/attachment/{file_id}/download"
        headers = self.base_headers.copy()
        try:
            r = await self.s.get(url, headers=headers, timeout=10)
            if r.status_code == 200:
                download_url = r.json().get('download_url')
                return download_url
            else:
                raise HTTPException(status_code=r.status_code, detail=r.text)
        except Exception as e:
            logger.error(f"Failed to get download url: {e}")
            return ""

    async def get_download_url_from_upload(self, file_id):
        url = f"{self.base_url}/files/{file_id}/uploaded"
        headers = self.base_headers.copy()
        try:
            r = await self.s.post(url, headers=headers, json={}, timeout=10)
            if r.status_code == 200:
                download_url = r.json().get('download_url')
                return download_url
            else:
                raise HTTPException(status_code=r.status_code, detail=r.text)
        except Exception as e:
            logger.error(f"Failed to get download url from upload: {e}")
            return ""

    async def get_upload_url(self, file_name, file_size, use_case="multimodal"):
        url = f'{self.base_url}/files'
        headers = self.base_headers.copy()
        try:
            r = await self.s.post(
                url,
                headers=headers,
                json={"file_name": file_name, "file_size": file_size, "reset_rate_limits": False, "timezone_offset_min": -480, "use_case": use_case},
                timeout=5,
            )
            if r.status_code == 200:
                res = r.json()
                file_id = res.get('file_id')
                upload_url = res.get('upload_url')
                logger.info(f"file_id: {file_id}, upload_url: {upload_url}")
                return file_id, upload_url
            else:
                raise HTTPException(status_code=r.status_code, detail=r.text)
        except Exception as e:
            logger.error(f"Failed to get upload url: {e}")
            return "", ""

    async def upload(self, upload_url, file_content, mime_type):
        headers = self.base_headers.copy()
        headers.update(
            {
                'accept': 'application/json, text/plain, */*',
                'content-type': mime_type,
                'x-ms-blob-type': 'BlockBlob',
                'x-ms-version': '2020-04-08',
            }
        )
        headers.pop('authorization', None)
        headers.pop('oai-device-id', None)
        headers.pop('oai-language', None)
        try:
            r = await self.s.put(upload_url, headers=headers, data=file_content, timeout=60)
            if r.status_code == 201:
                return True
            else:
                raise HTTPException(status_code=r.status_code, detail=r.text)
        except Exception as e:
            logger.error(f"Failed to upload file: {e}")
            return False

    async def upload_file(self, file_content, mime_type):
        if not file_content or not mime_type:
            return None

        width, height = None, None
        if mime_type.startswith("image/"):
            try:
                width, height = await get_image_size(file_content)
            except Exception as e:
                logger.error(f"Error image mime_type, change to text/plain: {e}")
                mime_type = 'text/plain'
        file_size = len(file_content)
        file_extension = await get_file_extension(mime_type)
        file_name = f"{uuid.uuid4()}{file_extension}"
        use_case = await determine_file_use_case(mime_type)

        file_id, upload_url = await self.get_upload_url(file_name, file_size, use_case)
        if file_id and upload_url:
            if await self.upload(upload_url, file_content, mime_type):
                download_url = await self.get_download_url_from_upload(file_id)
                if download_url:
                    file_meta = {
                        "file_id": file_id,
                        "file_name": file_name,
                        "size_bytes": file_size,
                        "mime_type": mime_type,
                        "width": width,
                        "height": height,
                        "use_case": use_case,
                    }
                    logger.info(f"File_meta: {file_meta}")
                    return file_meta

    async def check_upload(self, file_id):
        url = f'{self.base_url}/files/{file_id}'
        headers = self.base_headers.copy()
        try:
            for i in range(30):
                r = await self.s.get(url, headers=headers, timeout=5)
                if r.status_code == 200:
                    res = r.json()
                    retrieval_index_status = res.get('retrieval_index_status', '')
                    if retrieval_index_status == "success":
                        break
                await asyncio.sleep(1)
            return True
        except HTTPException:
            return False

    async def get_response_file_url(self, conversation_id, message_id, sandbox_path):
        try:
            url = f"{self.base_url}/conversation/{conversation_id}/interpreter/download"
            params = {"message_id": message_id, "sandbox_path": sandbox_path}
            headers = self.base_headers.copy()
            r = await self.s.get(url, headers=headers, params=params, timeout=10)
            if r.status_code == 200:
                return r.json().get("download_url")
            else:
                return None
        except Exception:
            logger.info("Failed to get response file url")
            return None

    async def close_client(self):
        if self.s:
            await self.s.close()
            del self.s
        if self.ss:
            await self.ss.close()
            del self.ss
        if self.ws:
            await self.ws.close()
            del self.ws
````

## File: chatgpt/fp.py
````python
import json
import random
import uuid

import ua_generator
from ua_generator.data.version import VersionRange
from ua_generator.options import Options

import utils.globals as globals
from utils import configs


def get_fp(req_token):
    fp = globals.fp_map.get(req_token, {})
    if fp and fp.get("user-agent") and fp.get("impersonate"):
        if "proxy_url" in fp.keys() and (fp["proxy_url"] is None or fp["proxy_url"] not in configs.proxy_url_list):
            fp["proxy_url"] = random.choice(configs.proxy_url_list) if configs.proxy_url_list else None
            globals.fp_map[req_token] = fp
            with open(globals.FP_FILE, "w", encoding="utf-8") as f:
                json.dump(globals.fp_map, f, indent=4)
        if globals.impersonate_list and "impersonate" in fp.keys() and fp["impersonate"] not in globals.impersonate_list:
            fp["impersonate"] = random.choice(globals.impersonate_list)
            globals.fp_map[req_token] = fp
            with open(globals.FP_FILE, "w", encoding="utf-8") as f:
                json.dump(globals.fp_map, f, indent=4)
        if configs.user_agents_list and "user-agent" in fp.keys() and fp["user-agent"] not in configs.user_agents_list:
            fp["user-agent"] = random.choice(configs.user_agents_list)
            globals.fp_map[req_token] = fp
            with open(globals.FP_FILE, "w", encoding="utf-8") as f:
                json.dump(globals.fp_map, f, indent=4)
        fp = {k.lower(): v for k, v in fp.items()}
        return fp
    else:
        options = Options(version_ranges={
            'chrome': VersionRange(min_version=124),
            'edge': VersionRange(min_version=124),
        })
        ua = ua_generator.generate(
            device=configs.device_tuple if configs.device_tuple else ('desktop'),
            browser=configs.browser_tuple if configs.browser_tuple else ('chrome', 'edge', 'firefox', 'safari'),
            platform=configs.platform_tuple if configs.platform_tuple else ('windows', 'macos'),
            options=options
        )
        fp = {
            "user-agent": ua.text if not configs.user_agents_list else random.choice(configs.user_agents_list),
            "impersonate": random.choice(globals.impersonate_list),
            "proxy_url": random.choice(configs.proxy_url_list) if configs.proxy_url_list else None,
            "oai-device-id": str(uuid.uuid4())
        }
        if ua.device == "desktop" and ua.browser in ("chrome", "edge"):
            fp["sec-ch-ua-platform"] = ua.ch.platform
            fp["sec-ch-ua"] = ua.ch.brands
            fp["sec-ch-ua-mobile"] = ua.ch.mobile

        if not req_token:
            return fp
        else:
            globals.fp_map[req_token] = fp
            with open(globals.FP_FILE, "w", encoding="utf-8") as f:
                json.dump(globals.fp_map, f, indent=4)
            return fp
````

## File: chatgpt/proofofWork.py
````python
import hashlib
import json
import random
import re
import time
import uuid
from datetime import datetime, timedelta, timezone
from html.parser import HTMLParser

import pybase64
import diskcache as dc

from utils.Logger import logger
from utils.configs import conversation_only

cores = [8, 16, 24, 32]
timeLayout = "%a %b %d %Y %H:%M:%S"

cache = dc.Cache('./data/pow_config_cache')
cached_scripts = []
cached_dpl = ""
cached_time = 0
cached_require_proof = ""

navigator_key = [
    "registerProtocolHandler−function registerProtocolHandler() { [native code] }",
    "storage−[object StorageManager]",
    "locks−[object LockManager]",
    "appCodeName−Mozilla",
    "permissions−[object Permissions]",
    "share−function share() { [native code] }",
    "webdriver−false",
    "managed−[object NavigatorManagedData]",
    "canShare−function canShare() { [native code] }",
    "vendor−Google Inc.",
    "vendor−Google Inc.",
    "mediaDevices−[object MediaDevices]",
    "vibrate−function vibrate() { [native code] }",
    "storageBuckets−[object StorageBucketManager]",
    "mediaCapabilities−[object MediaCapabilities]",
    "getGamepads−function getGamepads() { [native code] }",
    "bluetooth−[object Bluetooth]",
    "share−function share() { [native code] }",
    "cookieEnabled−true",
    "virtualKeyboard−[object VirtualKeyboard]",
    "product−Gecko",
    "mediaDevices−[object MediaDevices]",
    "canShare−function canShare() { [native code] }",
    "getGamepads−function getGamepads() { [native code] }",
    "product−Gecko",
    "xr−[object XRSystem]",
    "clipboard−[object Clipboard]",
    "storageBuckets−[object StorageBucketManager]",
    "unregisterProtocolHandler−function unregisterProtocolHandler() { [native code] }",
    "productSub−20030107",
    "login−[object NavigatorLogin]",
    "vendorSub−",
    "login−[object NavigatorLogin]",
    "getInstalledRelatedApps−function getInstalledRelatedApps() { [native code] }",
    "mediaDevices−[object MediaDevices]",
    "locks−[object LockManager]",
    "webkitGetUserMedia−function webkitGetUserMedia() { [native code] }",
    "vendor−Google Inc.",
    "xr−[object XRSystem]",
    "mediaDevices−[object MediaDevices]",
    "virtualKeyboard−[object VirtualKeyboard]",
    "virtualKeyboard−[object VirtualKeyboard]",
    "appName−Netscape",
    "storageBuckets−[object StorageBucketManager]",
    "presentation−[object Presentation]",
    "onLine−true",
    "mimeTypes−[object MimeTypeArray]",
    "credentials−[object CredentialsContainer]",
    "presentation−[object Presentation]",
    "getGamepads−function getGamepads() { [native code] }",
    "vendorSub−",
    "virtualKeyboard−[object VirtualKeyboard]",
    "serviceWorker−[object ServiceWorkerContainer]",
    "xr−[object XRSystem]",
    "product−Gecko",
    "keyboard−[object Keyboard]",
    "gpu−[object GPU]",
    "getInstalledRelatedApps−function getInstalledRelatedApps() { [native code] }",
    "webkitPersistentStorage−[object DeprecatedStorageQuota]",
    "doNotTrack",
    "clearAppBadge−function clearAppBadge() { [native code] }",
    "presentation−[object Presentation]",
    "serial−[object Serial]",
    "locks−[object LockManager]",
    "requestMIDIAccess−function requestMIDIAccess() { [native code] }",
    "locks−[object LockManager]",
    "requestMediaKeySystemAccess−function requestMediaKeySystemAccess() { [native code] }",
    "vendor−Google Inc.",
    "pdfViewerEnabled−true",
    "language−zh-CN",
    "setAppBadge−function setAppBadge() { [native code] }",
    "geolocation−[object Geolocation]",
    "userAgentData−[object NavigatorUAData]",
    "mediaCapabilities−[object MediaCapabilities]",
    "requestMIDIAccess−function requestMIDIAccess() { [native code] }",
    "getUserMedia−function getUserMedia() { [native code] }",
    "mediaDevices−[object MediaDevices]",
    "webkitPersistentStorage−[object DeprecatedStorageQuota]",
    "sendBeacon−function sendBeacon() { [native code] }",
    "hardwareConcurrency−32",
    "credentials−[object CredentialsContainer]",
    "storage−[object StorageManager]",
    "cookieEnabled−true",
    "pdfViewerEnabled−true",
    "windowControlsOverlay−[object WindowControlsOverlay]",
    "scheduling−[object Scheduling]",
    "pdfViewerEnabled−true",
    "hardwareConcurrency−32",
    "xr−[object XRSystem]",
    "webdriver−false",
    "getInstalledRelatedApps−function getInstalledRelatedApps() { [native code] }",
    "getInstalledRelatedApps−function getInstalledRelatedApps() { [native code] }",
    "bluetooth−[object Bluetooth]"
]
document_key = ['_reactListeningo743lnnpvdg', 'location']
window_key = [
    "0",
    "window",
    "self",
    "document",
    "name",
    "location",
    "customElements",
    "history",
    "navigation",
    "locationbar",
    "menubar",
    "personalbar",
    "scrollbars",
    "statusbar",
    "toolbar",
    "status",
    "closed",
    "frames",
    "length",
    "top",
    "opener",
    "parent",
    "frameElement",
    "navigator",
    "origin",
    "external",
    "screen",
    "innerWidth",
    "innerHeight",
    "scrollX",
    "pageXOffset",
    "scrollY",
    "pageYOffset",
    "visualViewport",
    "screenX",
    "screenY",
    "outerWidth",
    "outerHeight",
    "devicePixelRatio",
    "clientInformation",
    "screenLeft",
    "screenTop",
    "styleMedia",
    "onsearch",
    "isSecureContext",
    "trustedTypes",
    "performance",
    "onappinstalled",
    "onbeforeinstallprompt",
    "crypto",
    "indexedDB",
    "sessionStorage",
    "localStorage",
    "onbeforexrselect",
    "onabort",
    "onbeforeinput",
    "onbeforematch",
    "onbeforetoggle",
    "onblur",
    "oncancel",
    "oncanplay",
    "oncanplaythrough",
    "onchange",
    "onclick",
    "onclose",
    "oncontentvisibilityautostatechange",
    "oncontextlost",
    "oncontextmenu",
    "oncontextrestored",
    "oncuechange",
    "ondblclick",
    "ondrag",
    "ondragend",
    "ondragenter",
    "ondragleave",
    "ondragover",
    "ondragstart",
    "ondrop",
    "ondurationchange",
    "onemptied",
    "onended",
    "onerror",
    "onfocus",
    "onformdata",
    "oninput",
    "oninvalid",
    "onkeydown",
    "onkeypress",
    "onkeyup",
    "onload",
    "onloadeddata",
    "onloadedmetadata",
    "onloadstart",
    "onmousedown",
    "onmouseenter",
    "onmouseleave",
    "onmousemove",
    "onmouseout",
    "onmouseover",
    "onmouseup",
    "onmousewheel",
    "onpause",
    "onplay",
    "onplaying",
    "onprogress",
    "onratechange",
    "onreset",
    "onresize",
    "onscroll",
    "onsecuritypolicyviolation",
    "onseeked",
    "onseeking",
    "onselect",
    "onslotchange",
    "onstalled",
    "onsubmit",
    "onsuspend",
    "ontimeupdate",
    "ontoggle",
    "onvolumechange",
    "onwaiting",
    "onwebkitanimationend",
    "onwebkitanimationiteration",
    "onwebkitanimationstart",
    "onwebkittransitionend",
    "onwheel",
    "onauxclick",
    "ongotpointercapture",
    "onlostpointercapture",
    "onpointerdown",
    "onpointermove",
    "onpointerrawupdate",
    "onpointerup",
    "onpointercancel",
    "onpointerover",
    "onpointerout",
    "onpointerenter",
    "onpointerleave",
    "onselectstart",
    "onselectionchange",
    "onanimationend",
    "onanimationiteration",
    "onanimationstart",
    "ontransitionrun",
    "ontransitionstart",
    "ontransitionend",
    "ontransitioncancel",
    "onafterprint",
    "onbeforeprint",
    "onbeforeunload",
    "onhashchange",
    "onlanguagechange",
    "onmessage",
    "onmessageerror",
    "onoffline",
    "ononline",
    "onpagehide",
    "onpageshow",
    "onpopstate",
    "onrejectionhandled",
    "onstorage",
    "onunhandledrejection",
    "onunload",
    "crossOriginIsolated",
    "scheduler",
    "alert",
    "atob",
    "blur",
    "btoa",
    "cancelAnimationFrame",
    "cancelIdleCallback",
    "captureEvents",
    "clearInterval",
    "clearTimeout",
    "close",
    "confirm",
    "createImageBitmap",
    "fetch",
    "find",
    "focus",
    "getComputedStyle",
    "getSelection",
    "matchMedia",
    "moveBy",
    "moveTo",
    "open",
    "postMessage",
    "print",
    "prompt",
    "queueMicrotask",
    "releaseEvents",
    "reportError",
    "requestAnimationFrame",
    "requestIdleCallback",
    "resizeBy",
    "resizeTo",
    "scroll",
    "scrollBy",
    "scrollTo",
    "setInterval",
    "setTimeout",
    "stop",
    "structuredClone",
    "webkitCancelAnimationFrame",
    "webkitRequestAnimationFrame",
    "chrome",
    "caches",
    "cookieStore",
    "ondevicemotion",
    "ondeviceorientation",
    "ondeviceorientationabsolute",
    "launchQueue",
    "documentPictureInPicture",
    "getScreenDetails",
    "queryLocalFonts",
    "showDirectoryPicker",
    "showOpenFilePicker",
    "showSaveFilePicker",
    "originAgentCluster",
    "onpageswap",
    "onpagereveal",
    "credentialless",
    "speechSynthesis",
    "onscrollend",
    "webkitRequestFileSystem",
    "webkitResolveLocalFileSystemURL",
    "sendMsgToSolverCS",
    "webpackChunk_N_E",
    "__next_set_public_path__",
    "next",
    "__NEXT_DATA__",
    "__SSG_MANIFEST_CB",
    "__NEXT_P",
    "_N_E",
    "regeneratorRuntime",
    "__REACT_INTL_CONTEXT__",
    "DD_RUM",
    "_",
    "filterCSS",
    "filterXSS",
    "__SEGMENT_INSPECTOR__",
    "__NEXT_PRELOADREADY",
    "Intercom",
    "__MIDDLEWARE_MATCHERS",
    "__STATSIG_SDK__",
    "__STATSIG_JS_SDK__",
    "__STATSIG_RERENDER_OVERRIDE__",
    "_oaiHandleSessionExpired",
    "__BUILD_MANIFEST",
    "__SSG_MANIFEST",
    "__intercomAssignLocation",
    "__intercomReloadLocation"
]


class ScriptSrcParser(HTMLParser):
    def handle_starttag(self, tag, attrs):
        global cached_scripts, cached_dpl, cached_time
        if tag == "script":
            attrs_dict = dict(attrs)
            if "src" in attrs_dict:
                src = attrs_dict["src"]
                cached_scripts.append(src)
                match = re.search(r"c/[^/]*/_", src)
                if match:
                    cached_dpl = match.group(0)
                    cached_time = int(time.time())


def get_data_build_from_html(html_content):
    global cached_scripts, cached_dpl, cached_time
    parser = ScriptSrcParser()
    parser.feed(html_content)
    if not cached_scripts:
        cached_scripts.append("https://chatgpt.com/backend-api/sentinel/sdk.js")
    if not cached_dpl:
        match = re.search(r'<html[^>]*data-build="([^"]*)"', html_content)
        if match:
            data_build = match.group(1)
            cached_dpl = data_build
            cached_time = int(time.time())
            logger.info(f"Found dpl: {cached_dpl}")


async def get_dpl(service):
    global cached_scripts, cached_dpl, cached_time
    if int(time.time()) - cached_time < 15 * 60:
        return True
    headers = service.base_headers.copy()
    cached_scripts = []
    cached_dpl = ""
    try:
        if conversation_only:
            return True
        r = await service.s.get(f"{service.host_url}/", headers=headers, timeout=5)
        r.raise_for_status()
        get_data_build_from_html(r.text)
        if not cached_dpl:
            raise Exception("No Cached DPL")
        else:
            return True
    except Exception as e:
        logger.info(f"Failed to get dpl: {e}")
        cached_dpl = None
        cached_time = int(time.time())
        return False


def get_parse_time():
    now = datetime.now(timezone(timedelta(hours=-5)))
    return now.strftime(timeLayout) + " GMT-0500 (Eastern Standard Time)"


@cache.memoize(expire=3600 * 24 * 7)
def get_config(user_agent, req_token=None):
    config = [
        random.choice([1920 + 1080, 2560 + 1440, 1920 + 1200, 2560 + 1600]),
        get_parse_time(),
        4294705152,
        0,
        user_agent,
        random.choice(cached_scripts) if cached_scripts else "",
        cached_dpl,
        "en-US",
        "en-US,es-US,en,es",
        0,
        random.choice(navigator_key),
        random.choice(document_key),
        random.choice(window_key),
        time.perf_counter() * 1000,
        str(uuid.uuid4()),
        "",
        random.choice(cores),
        time.time() * 1000 - (time.perf_counter() * 1000),
    ]
    return config


def get_answer_token(seed, diff, config):
    start = time.time()
    answer, solved = generate_answer(seed, diff, config)
    end = time.time()
    logger.info(f'diff: {diff}, time: {int((end - start) * 1e6) / 1e3}ms, solved: {solved}')
    return "gAAAAAB" + answer, solved


def generate_answer(seed, diff, config):
    diff_len = len(diff)
    seed_encoded = seed.encode()
    static_config_part1 = (json.dumps(config[:3], separators=(',', ':'), ensure_ascii=False)[:-1] + ',').encode()
    static_config_part2 = (',' + json.dumps(config[4:9], separators=(',', ':'), ensure_ascii=False)[1:-1] + ',').encode()
    static_config_part3 = (',' + json.dumps(config[10:], separators=(',', ':'), ensure_ascii=False)[1:]).encode()

    target_diff = bytes.fromhex(diff)

    for i in range(500000):
        dynamic_json_i = str(i).encode()
        dynamic_json_j = str(i >> 1).encode()
        final_json_bytes = static_config_part1 + dynamic_json_i + static_config_part2 + dynamic_json_j + static_config_part3
        base_encode = pybase64.b64encode(final_json_bytes)
        hash_value = hashlib.sha3_512(seed_encoded + base_encode).digest()
        if hash_value[:diff_len] <= target_diff:
            return base_encode.decode(), True

    return "wQ8Lk5FbGpA2NcR9dShT6gYjU7VxZ4D" + pybase64.b64encode(f'"{seed}"'.encode()).decode(), False


def get_requirements_token(config):
    require, solved = generate_answer(format(random.random()), "0fffff", config)
    return 'gAAAAAC' + require


if __name__ == "__main__":
    # cached_scripts.append(
    #     "https://cdn.oaistatic.com/_next/static/cXh69klOLzS0Gy2joLDRS/_ssgManifest.js?dpl=453ebaec0d44c2decab71692e1bfe39be35a24b3")
    # cached_dpl = "453ebaec0d44c2decab71692e1bfe39be35a24b3"
    # cached_time = int(time.time())
    # for i in range(10):
    #     seed = format(random.random())
    #     diff = "000032"
    #     config = get_config("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome")
    #     answer = get_answer_token(seed, diff, config)
    cached_scripts.append(
        "https://cdn.oaistatic.com/_next/static/cXh69klOLzS0Gy2joLDRS/_ssgManifest.js?dpl=453ebaec0d44c2decab71692e1bfe39be35a24b3")
    cached_dpl = "prod-f501fe933b3edf57aea882da888e1a544df99840"
    config = get_config("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36")
    get_requirements_token(config)
````

## File: chatgpt/refreshToken.py
````python
import hashlib
import json
import random
import time

from fastapi import HTTPException

from utils.Client import Client
from utils.Logger import logger
from utils.configs import proxy_url_list
import utils.globals as globals


async def rt2ac(refresh_token, force_refresh=False):
    if not force_refresh and (refresh_token in globals.refresh_map and int(time.time()) - globals.refresh_map.get(refresh_token, {}).get("timestamp", 0) < 5 * 24 * 60 * 60):
        access_token = globals.refresh_map[refresh_token]["token"]
        # logger.info(f"refresh_token -> access_token from cache")
        return access_token
    else:
        try:
            access_token = await chat_refresh(refresh_token)
            globals.refresh_map[refresh_token] = {"token": access_token, "timestamp": int(time.time())}
            with open(globals.REFRESH_MAP_FILE, "w") as f:
                json.dump(globals.refresh_map, f, indent=4)
            logger.info(f"refresh_token -> access_token with openai: {access_token}")
            return access_token
        except HTTPException as e:
            raise HTTPException(status_code=e.status_code, detail=e.detail)


async def chat_refresh(refresh_token):
    data = {
        "client_id": "pdlLIX2Y72MIl2rhLhTE9VV9bN905kBh",
        "grant_type": "refresh_token",
        "redirect_uri": "com.openai.chat://auth0.openai.com/ios/com.openai.chat/callback",
        "refresh_token": refresh_token
    }
    session_id = hashlib.md5(refresh_token.encode()).hexdigest()
    proxy_url = random.choice(proxy_url_list).replace("{}", session_id) if proxy_url_list else None
    client = Client(proxy=proxy_url)
    try:
        r = await client.post("https://auth0.openai.com/oauth/token", json=data, timeout=15)
        if r.status_code == 200:
            access_token = r.json()['access_token']
            return access_token
        else:
            if "invalid_grant" in r.text or "access_denied" in r.text:
                if refresh_token not in globals.error_token_list:
                    globals.error_token_list.append(refresh_token)
                    with open(globals.ERROR_TOKENS_FILE, "a", encoding="utf-8") as f:
                        f.write(refresh_token + "\n")
                raise Exception(r.text)
            else:
                raise Exception(r.text[:300])
    except Exception as e:
        logger.error(f"Failed to refresh access_token `{refresh_token}`: {str(e)}")
        raise HTTPException(status_code=500, detail=f"Failed to refresh access_token.")
    finally:
        await client.close()
        del client
````

## File: chatgpt/turnstile.py
````python
import pybase64
import json
import random
import time
from typing import Any, Callable, Dict, List, Union


class OrderedMap:
    def __init__(self):
        self.keys = []
        self.values = {}

    def add(self, key: str, value: Any):
        if key not in self.values:
            self.keys.append(key)
        self.values[key] = value

    def to_json(self):
        return json.dumps({k: self.values[k] for k in self.keys})


TurnTokenList = List[List[Any]]
FloatMap = Dict[float, Any]
StringMap = Dict[str, Any]
FuncType = Callable[..., Any]


def get_turnstile_token(dx: str, p: str) -> Union[str, None]:
    try:
        decoded_bytes = pybase64.b64decode(dx)
        return process_turnstile_token(decoded_bytes.decode(), p)
    except Exception as e:
        print(f"Error in get_turnstile_token: {e}")
        return None


def process_turnstile_token(dx: str, p: str) -> str:
    result = []
    p_length = len(p)
    if p_length != 0:
        for i, r in enumerate(dx):
            result.append(chr(ord(r) ^ ord(p[i % p_length])))
    else:
        result = list(dx)
    return ''.join(result)


def is_slice(input_val: Any) -> bool:
    return isinstance(input_val, (list, tuple))


def is_float(input_val: Any) -> bool:
    return isinstance(input_val, float)


def is_string(input_val: Any) -> bool:
    return isinstance(input_val, str)


def to_str(input_val: Any) -> str:
    if input_val is None:
        return "undefined"
    elif is_float(input_val):
        return str(input_val)
    elif is_string(input_val):
        special_cases = {
            "window.Math": "[object Math]",
            "window.Reflect": "[object Reflect]",
            "window.performance": "[object Performance]",
            "window.localStorage": "[object Storage]",
            "window.Object": "function Object() { [native code] }",
            "window.Reflect.set": "function set() { [native code] }",
            "window.performance.now": "function () { [native code] }",
            "window.Object.create": "function create() { [native code] }",
            "window.Object.keys": "function keys() { [native code] }",
            "window.Math.random": "function random() { [native code] }"
        }
        return special_cases.get(input_val, input_val)
    elif isinstance(input_val, list) and all(isinstance(item, str) for item in input_val):
        return ','.join(input_val)
    else:
        return str(input_val)


def get_func_map() -> FloatMap:
    process_map: FloatMap = {}

    def func_1(e: float, t: float):
        e_str = to_str(process_map[e])
        t_str = to_str(process_map[t])
        res = process_turnstile_token(e_str, t_str)
        process_map[e] = res

    def func_2(e: float, t: Any):
        process_map[e] = t

    def func_5(e: float, t: float):
        n = process_map[e]
        tres = process_map[t]
        if is_slice(n):
            nt = n + [tres]
            process_map[e] = nt
        else:
            if is_string(n) or is_string(tres):
                res = to_str(n) + to_str(tres)
            elif is_float(n) and is_float(tres):
                res = n + tres
            else:
                res = "NaN"
            process_map[e] = res

    def func_6(e: float, t: float, n: float):
        tv = process_map[t]
        nv = process_map[n]
        if is_string(tv) and is_string(nv):
            res = f"{tv}.{nv}"
            if res == "window.document.location":
                process_map[e] = "https://chatgpt.com/"
            else:
                process_map[e] = res
        else:
            print("func type 6 error")

    def func_24(e: float, t: float, n: float):
        tv = process_map[t]
        nv = process_map[n]
        if is_string(tv) and is_string(nv):
            process_map[e] = f"{tv}.{nv}"
        else:
            print("func type 24 error")

    def func_7(e: float, *args):
        n = [process_map[arg] for arg in args]
        ev = process_map[e]
        if isinstance(ev, str):
            if ev == "window.Reflect.set":
                obj = n[0]
                key_str = str(n[1])
                val = n[2]
                obj.add(key_str, val)
        elif callable(ev):
            ev(*n)

    def func_17(e: float, t: float, *args):
        i = [process_map[arg] for arg in args]
        tv = process_map[t]
        res = None
        if isinstance(tv, str):
            if tv == "window.performance.now":
                current_time = time.time_ns()
                elapsed_ns = current_time - int(start_time * 1e9)
                res = (elapsed_ns + random.random()) / 1e6
            elif tv == "window.Object.create":
                res = OrderedMap()
            elif tv == "window.Object.keys":
                if isinstance(i[0], str) and i[0] == "window.localStorage":
                    res = ["STATSIG_LOCAL_STORAGE_INTERNAL_STORE_V4", "STATSIG_LOCAL_STORAGE_STABLE_ID",
                           "client-correlated-secret", "oai/apps/capExpiresAt", "oai-did",
                           "STATSIG_LOCAL_STORAGE_LOGGING_REQUEST", "UiState.isNavigationCollapsed.1"]
            elif tv == "window.Math.random":
                res = random.random()
        elif callable(tv):
            res = tv(*i)
        process_map[e] = res

    def func_8(e: float, t: float):
        process_map[e] = process_map[t]

    def func_14(e: float, t: float):
        tv = process_map[t]
        if is_string(tv):
            token_list = json.loads(tv)
            process_map[e] = token_list
        else:
            print("func type 14 error")

    def func_15(e: float, t: float):
        tv = process_map[t]
        process_map[e] = json.dumps(tv)

    def func_18(e: float):
        ev = process_map[e]
        e_str = to_str(ev)
        decoded = pybase64.b64decode(e_str).decode()
        process_map[e] = decoded

    def func_19(e: float):
        ev = process_map[e]
        e_str = to_str(ev)
        encoded = pybase64.b64encode(e_str.encode()).decode()
        process_map[e] = encoded

    def func_20(e: float, t: float, n: float, *args):
        o = [process_map[arg] for arg in args]
        ev = process_map[e]
        tv = process_map[t]
        if ev == tv:
            nv = process_map[n]
            if callable(nv):
                nv(*o)
            else:
                print("func type 20 error")

    def func_21(*args):
        pass

    def func_23(e: float, t: float, *args):
        i = list(args)
        ev = process_map[e]
        tv = process_map[t]
        if ev is not None:
            if callable(tv):
                tv(*i)

    process_map.update({
        1: func_1, 2: func_2, 5: func_5, 6: func_6, 24: func_24, 7: func_7,
        17: func_17, 8: func_8, 10: "window", 14: func_14, 15: func_15,
        18: func_18, 19: func_19, 20: func_20, 21: func_21, 23: func_23
    })

    return process_map

start_time = 0


def process_turnstile(dx: str, p: str) -> str:
    global start_time
    start_time = time.time()
    tokens = get_turnstile_token(dx, p)
    if tokens is None:
        return ""

    token_list = json.loads(tokens)
    # print(token_list)
    res = ""
    process_map = get_func_map()

    def func_3(e: str):
        nonlocal res
        res = pybase64.b64encode(e.encode()).decode()

    process_map[3] = func_3
    process_map[9] = token_list
    process_map[16] = p

    for token in token_list:
        try:
            e = token[0]
            t = token[1:]
            f = process_map.get(e)
            if callable(f):
                f(*t)
            else:
                pass
                # print(f"Warning: No function found for key {e}")
        except Exception as exc:
            pass
            # print(f"Error processing token {token}: {exc}")

    return res


if __name__ == "__main__":
    result = process_turnstile(
        "PBp5bWF1cHlLe1ttQhRfaTdmXEpidGdEYU5JdGJpR3xfHFVuGHVEY0tZVG18Vh54RWJ5CXpxKXl3SUZ7b2FZAWJaTBl6RGQZURh8BndUcRlQVgoYalAca2QUX24ffQZgdVVbbmBrAH9FV08Rb2oVVgBeQVRrWFp5VGZMYWNyMnoSN0FpaQgFT1l1f3h7c1RtcQUqY1kZbFJ5BQRiZEJXS3RvHGtieh9PaBlHaXhVWnVLRUlKdwsdbUtbKGFaAlN4a0V/emUJe2J2dl9BZkAxZWU/WGocRUBnc3VyT3F4WkJmYSthdBIGf0RwQ2FjAUBnd3ZEelgbVUEIDAJjS1VZbU9sSWFjfk55J2lZFV0HWX1cbVV5dWdAfkFIAVQVbloUXQtYaAR+VXhUF1BZdG4CBHRyK21AG1JaHhBFaBwCWUlocyQGVT4NBzNON2ASFVtXeQRET1kARndjUEBDT2RKeQN7RmJjeVtvZGpDeWJ1EHxafVd+Wk1AbzdLVTpafkd9dWZKeARecGJrS0xcenZIEEJQOmcFa01menFOeVRiSGFZC1JnWUA0SU08QGgeDFFgY34YWXAdZHYaHRhANFRMOV0CZmBfVExTWh9lZlVpSnx6eQURb2poa2RkQVJ0cmF0bwJbQgB6RlRbQHRQaQFKBHtENwVDSWpgHAlbTU1hXEpwdBh2eBlNY3l2UEhnblx7AmpaQ08JDDAzJUVAbn5IA2d8XX5ZFVlrYWhSXWlYQlEdZlQ/QUwuYwJgTG5GZghSRHdCYk1CWWBjclp0aWo3TWMSQmFaaAdge05FbmFhH3hxCFZuIX1BY01WVW5ABx5jfG1ZbjcZEiwwPFYQVm0sdHV8Xnl7alRuemgKZUwICklweW1heHR5Q3UqYVoSR3BCaldIc3Z8SmJOS212CAY5AmMkYmMaRn5UXEthZFsHYFx7ZHRnYV5tcFBZeHocQxUXXU0bYk0VFUZ0ZgFrSWcMRksCAwdJEBBncF12fGUVdnFNQnl4ZQB9WUclYGMRe04TQUZMf0FEbEthW357HEN2aVhAdHAMH0NPdWFicm1YbzNRBSkWMDUAOVdXbBlfRz51ah54YG5iVX9sR2t6RF1pR1RGU20MABBWQy55T3dQfmlUfmFrA35gY2AdDiBWMWVlP1hqHEVAZ3NzfE9/c1pCZWErYXQSB2BKcENjew1baXB9Rm1aG1VBCAkJY01aWW1NbklgZH5Oek1rTX9FFEB7RHNGEG9pKH1eRgFSZGJJdkcMQHUSY0IRQRkzUmFgBG90cklvVwNZThIHQXYABjFJaApCWh1qUEhnWVpiBHxDRDlAHg8kFVcCY1dCUk8VRm9obEN9e21EdnluWxN7eWt8RnFOekRTRXZKXkNPWH40YGMRXHwfRHZ7Z1JKS2R9XG1XR09qCGlaZmZ/QXwnfloWTQxIflxbSVNdSUZgHBRLKCwpQwwmXzB2NFRMOVxUTFNfH3BoRVhfWkcBYghVaSh0ZWMFeG9qBWp5eENNeGNldncHR0wBezVPTjdlSGcOTndjVkAUVl99YQFkRUE2YlNKe3ppeml2V2lvYkhGHjtbNHIALywsMScPEjEFO3Q1MQ0UGDYvK148ETYxIzEcD0gzchNcLSs+LAJxJiEQKBd5MCsXCRclFA0gBRg3axk1HTkBGyoUPRhwCwI2OAIRB2gUBRcjATt6ORQ9JDANOHFlEQITIC8VOS4GAC49GDscBBQMNQ4hDQtQZHYMHmk3BRFHeHZvcXNvd01+WXxPFF9pN2ZaSmR3Z0RkQkl7YmlHbzMsSS8HEy4PPggxGAAYBBcuJREBEQA7LAMANgEiNiZgFR5Mchs0eH83ERFsGCceZTESe2MeEgQSGwgXIgIbb38FFBAWEC1GFC42OQ0CCwcudSIpOwY6MRw7IjwYAgAYD3UbOA8AaHoHPiUkBgQmTA4FUxgAOCoJKxNmVSoANDIzAjdlDxA6ISIOKhQDEhwLPS82IT4CUFIsOyIwLD4+BBsDAww1AnMqHAIlMiMTGT0oAQlUE3QDQhIUACMxDwhGLxEXHQsSIV0FLgMaAgJ2LgsEHyEPLBcKOBtfUhg9MiAXPT5fHhA1Wg8+BxoPLgYcGS0WRSsELjIZKg8EJw4lFQAoUCcTcxASLS9BOTsZD3ERGRUhOD1YUjJxWBEBdnc9PwkQNytyED0zAQtaG3Y2ACsWXSsoPV4+DBQ2DyQ+bg0MHxVHKhAqNh8QPVkNET5fAis5Jh0uGxACKA8kOyo6IBkHIgkKdx0sAgA8SAQVHCkCLwcoBnQHGRAeAxAXOQAdKxhrNxMLJQYrKwAxHnFcOA4HIlEEAVkVDigqAwMoORQQKFkaOy0pISMoRmYDPyFLCRIqVhwCImITET04Gx8QPTMWWRQDcgstAioLGSkBTjw7ECYLeSgraxFoazw2CQcrJgU1cQ0fAB4YEykpIQMEPgJ0NUY0Lhc8IBEEWQtyNSkeECEmHitRFhsULgUrASkfO3E6XDsqLTAVcg8pFCwUaT8rPiMALzskFQQNJBkfKgUxBwscAj4YWhYHDxoXEBRwHgUUMx4gCxsCGBRJAz5yABsCAxIPFSo2AQILLSs7NS4EAGEnFBANJBgTOV0FLWJSKAUQeRkDKyAjCjYqIwEUBwAUPT5iBgohDzYmBAEBJS4pCSspGgUQBDsuD3wvKFd7HwE/EQ8ZFQgRICYEAgUuRhovHFYdM15eNwIgZBgmBVIoJGBnACRXChIKQR8lDVh2CicfKTIBcxwzNionIg4PEVI0FyMQOTkaABI3JSoAByVTKAItJn1ULjcEOG4gBjoqDnAQDjsGHzA2cF92CTIlAhMdchoJABA6KQEyajcgBAM+IhwyE292OTQ0IzUsAVY8EBcxMRxoKgEhBRQSGTMLfQsgFDp1PDQsCgEFKAkIASA8EhF4IgpjIzMJJC4WcyYcEQkPPSMBHlUSfFkuPCQnKiMaAGYWEC80EQIeex9wJjszCSQMFg4iDDcvVxMEBR17Knw0OnMVRyc4fj9ROQpiABoWFxAscR0Na3gBHWdyPjcOBCMleBQgKR4rLQViBhcLGnEgDDZ4ACoPJhQQIH4nHBoDNhkWCyUWDRgVFx4YAwAzFjAELCUPNScjDQ4hDB54Gwg4K2g3BmMBKjkwGggiFAo0Iwp6BBQeDxYwBz4VKCIzeDQmJjYeXTUmHCZpcygrAQt3NAFrBjsmGhtWJz8uUiR3CjorPy4NJXUuOjYIBDoMDGM4MwxxNiMNGg4SES01GHA1O3EIOSo7LQUXHnEeOgIjPXENLjQSfn4OVSkSAgcFBQIxDQUuajUPOj0MFwwcZhMnVzQOCQMDAWBWZBUPPx4oBAA5YA5qBwcrEwQ+IjppEz47Ji4CE2YNKTEzAUcjBgAoFFwyKHwbCz8pARUrDgIIMgg1H2MXGTUBFx0XAgMdEj0HOQ4MIionOyE2cUcxHAA7Iw0sNTkBDUU9GRsbPgkzOBwNKD9hHBdVJipxVTYRAgMmGAIVKxc2JREoNxgtMysDHggNExYWBh8FHwUfBQ8/KQYONiUrLjkfIwpxHDgYCTw1MDEMMBU2JRErK2crDzZdCy94UjAOC00MMgFCKTJxZw8mdgoSCzQMcAtzDC8hMBw7CHJ/GjQ+Cw4aDAVyMTMwEi8gHhUfNB8sDi4hWTQ0GDdJdSEVNggXAhY7Knd3MQ4KGhoZDm11DysqLxI8NXYZCXMDMngaMQg5PSsYKjYxJRJzdx8jOzQlIwklEwgtDhEMdwskLAs3Izg7LQscJi4IeyE3GiAbDAYrHzEzEjcxKicAdSteCTMqJHsUMSEXMT0kJD4Ga3V2Kk4rMSUZHS8qMAsqHTsEPR8RXzArXzc2OgYQOy4oPXc1AQM+DhpuMDFRFTMrBn8pCQkCdCE/MDILKG8uGllRNRlGRy0NGjsyFGoTKSUsOiwkAi8sNRJUNgQ0czEuFgUNMShjBAsBDDErbywzKBoKKzkeOncPDR42HCskNGg7BjEMVgAvOyApLQ5WPgAVHiM+Jz8eOA8BOSI7Xwo4JGIJNjYdCz0MFmAuPhEbLzc3VjUQAGwoHjATcSAGdwUVCjIqMDA1OyQNUB5gGRw6UwpkNS0eECoqbCt2KzQEdD1jBzEZOxQdIjBoMxVqCyoEBToSDB5xPz44LA9MCDAKMAZhLgZZACwMKAYDPWgHODIGHiwMIDUpZ2YEMA04By8INQl3ClQLLC8wCDIIXG8/PSARMDYQLxQyeh8qFTg7MhhUDzkLKwNzDT8RPQ84JC0dDTAqGDA7KxkoKDAcPzh1KQo9LzkeN3YMIxc4HzsBNxorAj0jQX90CCMlPQ4FMTYPfDgwDA0sMyoJHyw6EigMCwULUBsDcnsAdQUAKRAMFBIqLQwCGCkLLmoOJQIEOSU/JQ0JFQgmDx02LwgrIjMLHQQ9DCw+cgoRJREWZAQkCyoyNgskJip0JDg5cy1BXXIzJAl3GCQCdggwZXEbBmcPNAwwCAV9fAkGDDUUBhBmKTgyKAo0KRklcRc/IxY5KQ8SACIKEgg4FVUuDx0FUVoiK3IuEiQEGQkkYToJDhcPJhVTfA8zMiMhFgxnAystCycgLTweB1A0GAMuACIBVEUKHSYiCR0UJA0ENQsRBwUPCgEpMCcvGyUKdxcvH3U5OAwRegMnCiE1IxYiOgsGEGoOAhg/DxJ9IggHCzESCgMsJgJ9awodFDksDRAyCyA1NwodDCwJOFcWCw0yNwokfTUKLwt3IwolIwwocTcbRRAeCwoMHiUZOWkeCRclHihWMyVVcTcfVQEkJjAyMyReOT0jEFwMC1UPPyMwATQnO1oxHz8DNSIoAScYMBMtDi8iFgwgHwwKMAxnDjsXDQooCx4YHSY4JQYYPgQ0Cz0PVkQEEQYqKCIWPTELLBsxElgUMBcENhMKPQQRbyQVRhJdREdUW0tUYB4MX2BjeAU8bxEfZUVYW1VHTF5OSQV/f1xBMU5Jamd7QX9fbWd4H3p1ZhNuYmRFVHRyZHRnBltCCnxGV1YxeEQcDUp3ZlJAFFhafWEKFUlQQ25cOW9iHm90Yk5teXpaSGdhXHsBYStPTR1fdG5wHUIAZ0ZuZWVTeFQVWWliaFxSGFRQOARhQlRVQFVpBmBObEZmAUlKdU9gW0VFbHJkXW0Ffko6cmVTfEx3CXdvV1x+eWMDE2h1IXlJZ0J1VkNKe1cGBnZkcE1gdFJbbXdsWntMECo=",
        "gAAAAACWzMwMzIsIlRodSBKdWwgMTEgMjAyNCAwMzoxMDo0NiBHTVQrMDgwMCAo5Lit5Zu95qCH5YeG5pe26Ze0KSIsNDI5NDcwNTE1MiwxLCJNb3ppbGxhLzUuMCAoV2luZG93cyBOVCAxMC4wOyBXaW42NDsgeDY0KSBBcHBsZVdlYktpdC81MzcuMzYgKEtIVE1MLCBsaWtlIEdlY2tvKSBDaHJvbWUvMTI2LjAuMC4wIFNhZmFyaS81MzcuMzYgRWRnLzEyNi4wLjAuMCIsImh0dHBzOi8vY2RuLm9haXN0YXRpYy5jb20vX25leHQvc3RhdGljL2NodW5rcy9wYWdlcy9fYXBwLWMwOWZmNWY0MjQwMjcwZjguanMiLCJjL1pGWGkxeTNpMnpaS0EzSVQwNzRzMy9fIiwiemgtQ04iLCJ6aC1DTixlbixlbi1HQixlbi1VUyIsMTM1LCJ3ZWJraXRUZW1wb3JhcnlTdG9yYWdl4oiSW29iamVjdCBEZXByZWNhdGVkU3RvcmFnZVF1b3RhXSIsIl9yZWFjdExpc3RlbmluZ3NxZjF0ejFzNmsiLCJmZXRjaCIsMzY1NCwiNWU1NDUzNzItMzcyNy00ZDAyLTkwMDYtMzMwMDRjMWJmYTQ2Il0="
    )
    print(result)
````

## File: chatgpt/wssClient.py
````python
import json
import time

from utils.Logger import logger
import utils.globals as globals


def save_wss_map(wss_map):
    with open(globals.WSS_MAP_FILE, "w") as f:
        json.dump(wss_map, f, indent=4)


async def token2wss(token):
    if not token:
        return False, None
    if token in globals.wss_map:
        wss_mode = globals.wss_map[token]["wss_mode"]
        if wss_mode:
            if int(time.time()) - globals.wss_map.get(token, {}).get("timestamp", 0) < 60 * 60:
                wss_url = globals.wss_map[token]["wss_url"]
                logger.info(f"token -> wss_url from cache")
                return wss_mode, wss_url
            else:
                logger.info(f"token -> wss_url expired")
                return wss_mode, None
        else:
            return False, None
    return False, None


async def set_wss(token, wss_mode, wss_url=None):
    if not token:
        return True
    globals.wss_map[token] = {"timestamp": int(time.time()), "wss_url": wss_url, "wss_mode": wss_mode}
    save_wss_map(globals.wss_map)
    return True
````

## File: docker-compose-warp.yml
````yaml
services:
  warp:
    image: caomingjun/warp
    container_name: warp
    restart: always
    devices:
      - /dev/net/tun:/dev/net/tun
    environment:
      - WARP_SLEEP=5
    cap_add:
      - MKNOD
      - AUDIT_WRITE
      - NET_ADMIN
    sysctls:
      - net.ipv6.conf.all.disable_ipv6=0
      - net.ipv4.conf.all.src_valid_mark=1
    volumes:
      - ./warpdata:/var/lib/cloudflare-warp
    networks:
      - internal_network  # 使用内部网络，不对外暴露端口
    healthcheck:
      test: ["CMD", "curl", "-f", "-s", "https://www.google.com"]  # 静默模式下请求Google，如果成功返回2xx状态码
      interval: 30s  # 每隔30秒检查一次
      timeout: 10s   # 请求超时10秒
      retries: 3     # 失败3次后标记为不健康
      start_period: 5s  # 容器启动后等待5秒再开始检查
      
  chat2api:
    image: lanqian528/chat2api:latest
    container_name: chat2api
    restart: unless-stopped
    ports:
      - '5005:5005'  # 暴露chat2api服务供外部访问
    environment:
      - TZ=Asia/Shanghai
      - AUTHORIZATION=sk-xxx
      - PROXY_URL=socks5://warp:1080  # 设置 PROXY_URL 为 warp 容器的代理地址
    depends_on:
      warp:
        condition: service_healthy  # 只有 warp 的健康检查通过时，chat2api 才会启动
    networks:
      - internal_network  # chat2api 和 warp 在同一个内部网络
    volumes:
      - ./data:/app/data # 挂载一些需要保存的数据

  watchtower:
    image: containrrr/watchtower
    container_name: watchtower
    restart: unless-stopped
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    command: --cleanup --interval 300 chat2api
    
networks:
  internal_network:
    driver: bridge  # 定义一个桥接网络
````

## File: docker-compose.yml
````yaml
version: '3'

services:
  chat2api:
    image: lanqian528/chat2api:latest
    container_name: chat2api
    restart: unless-stopped
    ports:
      - '5005:5005'
    volumes:
      - ./data:/app/data
    environment:
      - TZ=Asia/Shanghai
      - AUTHORIZATION=sk-xxx

  watchtower:
    image: containrrr/watchtower
    container_name: watchtower
    restart: unless-stopped
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    command: --cleanup --interval 300 chat2api
````

## File: Dockerfile
````dockerfile
FROM python:3.11-slim

WORKDIR /app

COPY . /app

RUN pip install --no-cache-dir -r requirements.txt

EXPOSE 5005

CMD ["python", "app.py"]
````

## File: gateway/admin.py
````python

````

## File: gateway/backend.py
````python
import hashlib
import json
import random
import re
import time
import uuid

from fastapi import Request, HTTPException
from fastapi.responses import RedirectResponse, StreamingResponse, Response
from starlette.background import BackgroundTask
from starlette.concurrency import run_in_threadpool

import utils.globals as globals
from app import app
from chatgpt.authorization import verify_token
from chatgpt.fp import get_fp
from chatgpt.proofofWork import get_answer_token, get_config, get_requirements_token
from gateway.chatgpt import chatgpt_html
from gateway.reverseProxy import chatgpt_reverse_proxy, content_generator, get_real_req_token, headers_reject_list, \
    headers_accept_list
from utils.Client import Client
from utils.Logger import logger
from utils.configs import x_sign, turnstile_solver_url, chatgpt_base_url_list, no_sentinel, sentinel_proxy_url_list, \
    force_no_history

banned_paths = [
    "backend-api/accounts/logout_all",
    "backend-api/accounts/deactivate",
    "backend-api/payments",
    "backend-api/subscriptions",
    "backend-api/user_system_messages",
    "backend-api/memories",
    "backend-api/settings/clear_account_user_memory",
    "backend-api/conversations/[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}",
    "backend-api/accounts/mfa_info",
    "backend-api/accounts/[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}/invites",
    "admin",
]
redirect_paths = ["auth/logout"]
chatgpt_paths = ["c/[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}"]


@app.get("/backend-api/accounts/check/v4-2023-04-27")
async def check_account(request: Request):
    token = request.headers.get("Authorization").replace("Bearer ", "")
    check_account_response = await chatgpt_reverse_proxy(request, "backend-api/accounts/check/v4-2023-04-27")
    if len(token) == 45 or token.startswith("eyJhbGciOi"):
        return check_account_response
    else:
        check_account_str = check_account_response.body.decode('utf-8')
        check_account_info = json.loads(check_account_str)
        for key in check_account_info.get("accounts", {}).keys():
            account_id = check_account_info["accounts"][key]["account"]["account_id"]
            globals.seed_map[token]["user_id"] = \
                check_account_info["accounts"][key]["account"]["account_user_id"].split("__")[0]
            check_account_info["accounts"][key]["account"]["account_user_id"] = f"user-chatgpt__{account_id}"
        with open(globals.SEED_MAP_FILE, "w", encoding="utf-8") as f:
            json.dump(globals.seed_map, f, indent=4)
        return check_account_info


@app.get("/backend-api/gizmos/bootstrap")
async def get_gizmos_bootstrap(request: Request):
    token = request.headers.get("Authorization", "").replace("Bearer ", "")
    if len(token) == 45 or token.startswith("eyJhbGciOi"):
        return await chatgpt_reverse_proxy(request, "backend-api/gizmos/bootstrap")
    else:
        return {"gizmos": []}


@app.get("/backend-api/gizmos/pinned")
async def get_gizmos_pinned(request: Request):
    token = request.headers.get("Authorization", "").replace("Bearer ", "")
    if len(token) == 45 or token.startswith("eyJhbGciOi"):
        return await chatgpt_reverse_proxy(request, "backend-api/gizmos/pinned")
    else:
        return {"items": [], "cursor": None}


@app.get("/public-api/gizmos/discovery/recent")
async def get_gizmos_discovery_recent(request: Request):
    token = request.headers.get("Authorization", "").replace("Bearer ", "")
    if len(token) == 45 or token.startswith("eyJhbGciOi"):
        return await chatgpt_reverse_proxy(request, "public-api/gizmos/discovery/recent")
    else:
        return {
            "info": {
                "id": "recent",
                "title": "Recently Used",
            },
            "list": {
                "items": [],
                "cursor": None
            }
        }


@app.get("/backend-api/gizmos/snorlax/sidebar")
async def get_gizmos_snorlax_sidebar(request: Request):
    token = request.headers.get("Authorization", "").replace("Bearer ", "")
    if len(token) == 45 or token.startswith:
        return await chatgpt_reverse_proxy(request, "backend-api/gizmos/snorlax/sidebar")
    else:
        return {"items": [], "cursor": None}


@app.post("/backend-api/gizmos/snorlax/upsert")
async def get_gizmos_snorlax_upsert(request: Request):
    token = request.headers.get("Authorization", "").replace("Bearer ", "")
    if len(token) == 45 or token.startswith:
        return await chatgpt_reverse_proxy(request, "backend-api/gizmos/snorlax/upsert")
    else:
        raise HTTPException(status_code=403, detail="Forbidden")


@app.get("/backend-api/subscriptions")
async def post_subscriptions(request: Request):
    return {
        "id": str(uuid.uuid4()),
        "plan_type": "free",
        "seats_in_use": 1,
        "seats_entitled": 1,
        "active_until": "2050-01-01T00:00:00Z",
        "billing_period": None,
        "will_renew": True,
        "non_profit_org_discount_applied": None,
        "billing_currency": "USD",
        "is_delinquent": False,
        "became_delinquent_timestamp": None,
        "grace_period_end_timestamp": None
    }


@app.api_route("/backend-api/conversations", methods=["GET", "PATCH"])
async def get_conversations(request: Request):
    token = request.headers.get("Authorization", "").replace("Bearer ", "")
    if len(token) == 45 or token.startswith("eyJhbGciOi"):
        return await chatgpt_reverse_proxy(request, "backend-api/conversations")
    if request.method == "GET":
        limit = int(request.query_params.get("limit", 28))
        offset = int(request.query_params.get("offset", 0))
        is_archived = request.query_params.get("is_archived", None)
        items = []
        for conversation_id in globals.seed_map.get(token, {}).get("conversations", []):
            conversation = globals.conversation_map.get(conversation_id, None)
            if conversation:
                if is_archived == "true":
                    if conversation.get("is_archived", False):
                        items.append(conversation)
                else:
                    if not conversation.get("is_archived", False):
                        items.append(conversation)
        items = items[int(offset):int(offset) + int(limit)]
        conversations = {
            "items": items,
            "total": len(items),
            "limit": limit,
            "offset": offset,
            "has_missing_conversations": False
        }
        return Response(content=json.dumps(conversations, indent=4), media_type="application/json")
    else:
        raise HTTPException(status_code=403, detail="Forbidden")


@app.get("/backend-api/conversation/{conversation_id}")
async def update_conversation(request: Request, conversation_id: str):
    token = request.headers.get("Authorization", "").replace("Bearer ", "")
    conversation_details_response = await chatgpt_reverse_proxy(request,
                                                                f"backend-api/conversation/{conversation_id}")
    if len(token) == 45 or token.startswith("eyJhbGciOi"):
        return conversation_details_response
    else:
        conversation_details_str = conversation_details_response.body.decode('utf-8')
        conversation_details = json.loads(conversation_details_str)
        if conversation_id in globals.seed_map[token][
            "conversations"] and conversation_id in globals.conversation_map:
            globals.conversation_map[conversation_id]["title"] = conversation_details.get("title", None)
            globals.conversation_map[conversation_id]["is_archived"] = conversation_details.get("is_archived",
                                                                                                False)
            globals.conversation_map[conversation_id]["conversation_template_id"] = conversation_details.get(
                "conversation_template_id", None)
            globals.conversation_map[conversation_id]["gizmo_id"] = conversation_details.get("gizmo_id", None)
            globals.conversation_map[conversation_id]["async_status"] = conversation_details.get("async_status",
                                                                                                 None)
            with open(globals.CONVERSATION_MAP_FILE, "w", encoding="utf-8") as f:
                json.dump(globals.conversation_map, f, indent=4)
        return conversation_details_response


@app.patch("/backend-api/conversation/{conversation_id}")
async def patch_conversation(request: Request, conversation_id: str):
    token = request.headers.get("Authorization", "").replace("Bearer ", "")
    patch_response = (await chatgpt_reverse_proxy(request, f"backend-api/conversation/{conversation_id}"))
    if len(token) == 45 or token.startswith("eyJhbGciOi"):
        return patch_response
    else:
        data = await request.json()
        if conversation_id in globals.seed_map[token][
            "conversations"] and conversation_id in globals.conversation_map:
            if not data.get("is_visible", True):
                globals.conversation_map.pop(conversation_id)
                globals.seed_map[token]["conversations"].remove(conversation_id)
                with open(globals.SEED_MAP_FILE, "w", encoding="utf-8") as f:
                    json.dump(globals.seed_map, f, indent=4)
            else:
                globals.conversation_map[conversation_id].update(data)
            with open(globals.CONVERSATION_MAP_FILE, "w", encoding="utf-8") as f:
                json.dump(globals.conversation_map, f, indent=4)
        return patch_response


@app.get("/backend-api/me")
async def get_me(request: Request):
    token = request.headers.get("Authorization", "").replace("Bearer ", "")
    if len(token) == 45 or token.startswith("eyJhbGciOi"):
        return await chatgpt_reverse_proxy(request, "backend-api/me")
    else:
        me = {
            "object": "user",
            "id": "org-chatgpt",
            "email": "chatgpt@openai.com",
            "name": "ChatGPT",
            "picture": "https://cdn.auth0.com/avatars/ai.png",
            "created": int(time.time()),
            "phone_number": None,
            "mfa_flag_enabled": False,
            "amr": [],
            "groups": [],
            "orgs": {
                "object": "list",
                "data": [
                    {
                        "object": "organization",
                        "id": "org-chatgpt",
                        "created": 1715641300,
                        "title": "Personal",
                        "name": "user-chatgpt",
                        "description": "Personal org for chatgpt@openai.com",
                        "personal": True,
                        "settings": {
                            "threads_ui_visibility": "NONE",
                            "usage_dashboard_visibility": "ANY_ROLE",
                            "disable_user_api_keys": False
                        },
                        "parent_org_id": None,
                        "is_default": True,
                        "role": "owner",
                        "is_scale_tier_authorized_purchaser": None,
                        "is_scim_managed": False,
                        "projects": {
                            "object": "list",
                            "data": []
                        },
                        "groups": [],
                        "geography": None
                    }
                ]
            },
            "has_payg_project_spend_limit": True
        }
    return Response(content=json.dumps(me, indent=4), media_type="application/json")


@app.get("/backend-api/tasks")
async def get_me(request: Request):
    token = request.headers.get("Authorization", "").replace("Bearer ", "")
    if len(token) == 45 or token.startswith("eyJhbGciOi"):
        return await chatgpt_reverse_proxy(request, "backend-api/tasks")
    else:
        tasks = {
            "tasks": [],
            "cursor": None
        }
    return Response(content=json.dumps(tasks, indent=4), media_type="application/json")


@app.get("/backend-api/user_system_messages")
async def get_me(request: Request):
    token = request.headers.get("Authorization", "").replace("Bearer ", "")
    if len(token) == 45 or token.startswith("eyJhbGciOi"):
        return await chatgpt_reverse_proxy(request, "backend-api/user_system_messages")
    else:
        user_system_messages = {
            "object": "user_system_message_detail",
            "enabled": True,
            "about_user_message": "",
            "about_model_message": "",
            "name_user_message": "",
            "role_user_message": "",
            "traits_model_message": "",
            "other_user_message": "",
            "disabled_tools": []
        }
    return Response(content=json.dumps(user_system_messages, indent=4), media_type="application/json")


@app.get("/backend-api/memories")
async def get_me(request: Request):
    token = request.headers.get("Authorization", "").replace("Bearer ", "")
    if len(token) == 45 or token.startswith("eyJhbGciOi"):
        return await chatgpt_reverse_proxy(request, "backend-api/memories")
    else:
        memories = {"memories":[],"memory_max_tokens":10000,"memory_num_tokens":0}
    return Response(content=json.dumps(memories, indent=4), media_type="application/json")


# @app.get("/backend-api/system_hints")
# async def get_me(request: Request):
#     token = request.headers.get("Authorization", "").replace("Bearer ", "")
#     if len(token) == 45 or token.startswith("eyJhbGciOi"):
#         return await chatgpt_reverse_proxy(request, "backend-api/system_hints")
#     else:
#         system_hints = {
#             "system_hints": [
#                 {
#                     "system_hint": "picture_v2",
#                     "name": "创建图片",
#                     "description": "Visualize ideas and concepts",
#                     "logo": "<svg fill=\"none\" viewBox=\"0 0 22 22\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"m16.902 2.304c0.8187-0.42416 1.787-0.23298 2.4072 0.38717 0.6201 0.62015 0.8113 1.5885 0.3872 2.4071-0.7866 1.5181-1.7193 2.8479-2.8417 4.0435 1.1284 1.4032 1.9764 2.8568 2.3346 4.2247 0.4175 1.5942 0.1834 3.2014-1.1994 4.3077-1.1107 0.8885-2.3588 0.7686-3.3556 0.3569-0.9702-0.4008-1.8307-1.1176-2.4214-1.6881-0.3641-0.3517-0.3742-0.932-0.0224-1.2962 0.3517-0.3641 0.932-0.3742 1.2961-0.0225 0.532 0.5139 1.195 1.0427 1.8477 1.3123 0.6261 0.2587 1.1004 0.2339 1.5103-0.094 0.6469-0.5176 0.8677-1.2794 0.5712-2.4116-0.2612-0.9974-0.9109-2.1738-1.8885-3.4132-0.5323 0.463-1.1008 0.9058-1.7087 1.3325-0.027 0.019-0.0546 0.0363-0.0828 0.0519-0.0568 0.6667-0.2811 1.2371-0.676 1.6918-0.4495 0.5176-1.0422 0.7951-1.6087 0.9495-0.903 0.2461-1.9839 0.2303-2.7726 0.2188-0.15178-0.0022-0.29274-0.0043-0.41953-0.0043-0.50626 0-0.91666-0.4104-0.91666-0.9166 0-0.1268-0.00206-0.2678-0.00427-0.4195-0.01149-0.7887-0.02725-1.8696 0.21885-2.7726 0.15439-0.56651 0.43185-1.1593 0.94945-1.6088 0.45468-0.39487 1.0252-0.61914 1.6919-0.67597 0.0156-0.02818 0.0329-0.05583 0.0518-0.0828 0.3644-0.51911 0.7406-1.0095 1.1309-1.4731-2.6496-1.4543-5.1975-1.3354-6.519-0.01387-0.88711 0.88711-1.2375 2.2875-0.92283 3.9656 0.31366 1.6728 1.2805 3.522 2.8674 5.1089 0.93686 0.9369 1.9491 1.5911 2.7771 2.0091 0.41354 0.2088 0.77359 0.3548 1.0457 0.4465 0.2037 0.0686 0.3205 0.0939 0.3634 0.1032 0.0226 0.0049 0.0246 0.0053 0.0078 0.0053 0.5063 0 0.9167 0.4104 0.9167 0.9167s-0.4104 0.9167-0.9167 0.9166c-0.2776 0-0.635-0.0961-0.9564-0.2044-0.36088-0.1216-0.8013-0.3022-1.2867-0.5473-0.96991-0.4896-2.1496-1.2517-3.2473-2.3493-1.814-1.814-2.9828-3.9868-3.3729-6.0675-0.38915-2.0754-0.01137-4.16 1.4284-5.5998 2.2489-2.249 5.9795-1.922 9.0975-0.06649 1.2458-1.2059 2.6376-2.1982 4.2389-3.0279zm-4.749 6.3774c0.4827 0.28036 0.8856 0.68317 1.1659 1.1659 2.0754-1.5631 3.5853-3.3452 4.7496-5.5925 0.0362-0.06982 0.0317-0.17995-0.0557-0.26737-0.0874-0.08741-0.1975-0.09189-0.2674-0.05572-2.2472 1.1644-4.0293 2.6743-5.5924 4.7496zm-2.9847 4.1507c0.65462 0.0034 1.2725-0.0135 1.8007-0.1575 0.3559-0.097 0.575-0.2311 0.7066-0.3827 0.1192-0.1373 0.2413-0.3736 0.2413-0.8378 0-0.757-0.6136-1.3707-1.3706-1.3707-0.4642 0-0.70058 0.1221-0.83784 0.2413-0.1516 0.1317-0.28575 0.3507-0.38276 0.7066-0.14395 0.5283-0.16088 1.1461-0.15742 1.8008z\" fill=\"currentColor\"/></svg>",
#                     "required_features": [
#                         "image_gen_tool_enabled"
#                     ],
#                     "required_models": [],
#                     "required_conversation_modes": [],
#                     "allow_in_temporary_chat": True,
#                     "composer_bar_button_info": None,
#                     "suggested_prompt": {
#                         "theme": "#512AEB",
#                         "title": "创建图片",
#                         "subtitle": "Visualize ideas and concepts",
#                         "sort_order": 2,
#                         "badge": None
#                     },
#                     "regex_matches": [
#                         "image"
#                     ]
#                 },
#                 {
#                     "system_hint": "search",
#                     "name": "搜索",
#                     "description": "在网上查找",
#                     "logo": "<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" class=\"\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M2 12C2 6.47715 6.47715 2 12 2C17.5228 2 22 6.47715 22 12C22 17.5228 17.5228 22 12 22C6.47715 22 2 17.5228 2 12ZM11.9851 4.00291C11.9933 4.00046 11.9982 4.00006 11.9996 4C12.001 4.00006 12.0067 4.00046 12.0149 4.00291C12.0256 4.00615 12.047 4.01416 12.079 4.03356C12.2092 4.11248 12.4258 4.32444 12.675 4.77696C12.9161 5.21453 13.1479 5.8046 13.3486 6.53263C13.6852 7.75315 13.9156 9.29169 13.981 11H10.019C10.0844 9.29169 10.3148 7.75315 10.6514 6.53263C10.8521 5.8046 11.0839 5.21453 11.325 4.77696C11.5742 4.32444 11.7908 4.11248 11.921 4.03356C11.953 4.01416 11.9744 4.00615 11.9851 4.00291ZM8.01766 11C8.08396 9.13314 8.33431 7.41167 8.72334 6.00094C8.87366 5.45584 9.04762 4.94639 9.24523 4.48694C6.48462 5.49946 4.43722 7.9901 4.06189 11H8.01766ZM4.06189 13H8.01766C8.09487 15.1737 8.42177 17.1555 8.93 18.6802C9.02641 18.9694 9.13134 19.2483 9.24522 19.5131C6.48461 18.5005 4.43722 16.0099 4.06189 13ZM10.019 13H13.981C13.9045 14.9972 13.6027 16.7574 13.1726 18.0477C12.9206 18.8038 12.6425 19.3436 12.3823 19.6737C12.2545 19.8359 12.1506 19.9225 12.0814 19.9649C12.0485 19.9852 12.0264 19.9935 12.0153 19.9969C12.0049 20.0001 11.9999 20 11.9999 20C11.9999 20 11.9948 20 11.9847 19.9969C11.9736 19.9935 11.9515 19.9852 11.9186 19.9649C11.8494 19.9225 11.7455 19.8359 11.6177 19.6737C11.3575 19.3436 11.0794 18.8038 10.8274 18.0477C10.3973 16.7574 10.0955 14.9972 10.019 13ZM15.9823 13C15.9051 15.1737 15.5782 17.1555 15.07 18.6802C14.9736 18.9694 14.8687 19.2483 14.7548 19.5131C17.5154 18.5005 19.5628 16.0099 19.9381 13H15.9823ZM19.9381 11C19.5628 7.99009 17.5154 5.49946 14.7548 4.48694C14.9524 4.94639 15.1263 5.45584 15.2767 6.00094C15.6657 7.41167 15.916 9.13314 15.9823 11H19.9381Z\" fill=\"currentColor\"></path></svg>",
#                     "required_features": [
#                         "search"
#                     ],
#                     "required_models": [],
#                     "required_conversation_modes": [
#                         "primary_assistant"
#                     ],
#                     "allow_in_temporary_chat": True,
#                     "composer_bar_button_info": None,
#                     "suggested_prompt": None,
#                     "regex_matches": None
#                 },
#                 {
#                     "system_hint": "reason",
#                     "name": "推理",
#                     "description": "使用 o3-mini",
#                     "logo": "<svg fill=\"none\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"m12 3c-3.585 0-6.5 2.9225-6.5 6.5385 0 2.2826 1.162 4.2913 2.9248 5.4615h7.1504c1.7628-1.1702 2.9248-3.1789 2.9248-5.4615 0-3.6159-2.915-6.5385-6.5-6.5385zm2.8653 14h-5.7306v1h5.7306v-1zm-1.1329 3h-3.4648c0.3458 0.5978 0.9921 1 1.7324 1s1.3866-0.4022 1.7324-1zm-5.6064 0c0.44403 1.7252 2.0101 3 3.874 3s3.43-1.2748 3.874-3c0.5483-0.0047 0.9913-0.4506 0.9913-1v-2.4593c2.1969-1.5431 3.6347-4.1045 3.6347-7.0022 0-4.7108-3.8008-8.5385-8.5-8.5385-4.6992 0-8.5 3.8276-8.5 8.5385 0 2.8977 1.4378 5.4591 3.6347 7.0022v2.4593c0 0.5494 0.44301 0.9953 0.99128 1z\" clip-rule=\"evenodd\" fill=\"currentColor\" fill-rule=\"evenodd\"/></svg>",
#                     "required_features": [],
#                     "required_models": [
#                         "o1",
#                         "o3-mini"
#                     ],
#                     "required_conversation_modes": [
#                         "primary_assistant"
#                     ],
#                     "allow_in_temporary_chat": True,
#                     "composer_bar_button_info": {
#                         "disabled_text": "推理不可用",
#                         "tooltip_text": "思考后再回复",
#                         "announcement_key": "",
#                         "nux_title": "",
#                         "nux_description": "ChatGPT 可以先思考更长时间再回复，以便更好地回答您的重大问题。",
#                         "rate_limit_reached_text": None
#                     },
#                     "suggested_prompt": None,
#                     "regex_matches": None
#                 },
#                 {
#                     "system_hint": "canvas",
#                     "name": "画布",
#                     "description": "在写作和代码方面开展协作",
#                     "logo": "<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M2.5 5.5C4.3 5.2 5.2 4 5.5 2.5C5.8 4 6.7 5.2 8.5 5.5C6.7 5.8 5.8 7 5.5 8.5C5.2 7 4.3 5.8 2.5 5.5Z\" fill=\"currentColor\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\"/><path d=\"M5.66282 16.5231L5.18413 19.3952C5.12203 19.7678 5.09098 19.9541 5.14876 20.0888C5.19933 20.2067 5.29328 20.3007 5.41118 20.3512C5.54589 20.409 5.73218 20.378 6.10476 20.3159L8.97693 19.8372C9.72813 19.712 10.1037 19.6494 10.4542 19.521C10.7652 19.407 11.0608 19.2549 11.3343 19.068C11.6425 18.8575 11.9118 18.5882 12.4503 18.0497L20 10.5C21.3807 9.11929 21.3807 6.88071 20 5.5C18.6193 4.11929 16.3807 4.11929 15 5.5L7.45026 13.0497C6.91175 13.5882 6.6425 13.8575 6.43197 14.1657C6.24513 14.4392 6.09299 14.7348 5.97903 15.0458C5.85062 15.3963 5.78802 15.7719 5.66282 16.5231Z\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"/><path d=\"M14.5 7L18.5 11\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"/></svg>",
#                     "required_features": [
#                         "canvas"
#                     ],
#                     "required_models": [],
#                     "required_conversation_modes": [],
#                     "allow_in_temporary_chat": False,
#                     "composer_bar_button_info": None,
#                     "suggested_prompt": {
#                         "theme": "#AF52DE",
#                         "title": "画布",
#                         "subtitle": "写作和编程",
#                         "sort_order": 3,
#                         "badge": None
#                     },
#                     "regex_matches": None
#                 },
#                 {
#                     "system_hint": "research",
#                     "name": "深入研究",
#                     "description": "对任何主题都有详细的见解",
#                     "logo": "<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 24 24\" fill=\"none\"><path fill=\"currentColor\" fill-rule=\"evenodd\" d=\"M12.47 15.652a1 1 0 0 1 1.378.318l2.5 4a1 1 0 1 1-1.696 1.06l-2.5-4a1 1 0 0 1 .318-1.378Z\" clip-rule=\"evenodd\"/><path fill=\"currentColor\" fill-rule=\"evenodd\" d=\"M11.53 15.652a1 1 0 0 1 .318 1.378l-2.5 4a1 1 0 0 1-1.696-1.06l2.5-4a1 1 0 0 1 1.378-.318ZM17.824 4.346a.5.5 0 0 0-.63-.321l-.951.309a1 1 0 0 0-.642 1.26l1.545 4.755a1 1 0 0 0 1.26.642l.95-.309a.5.5 0 0 0 .322-.63l-1.854-5.706Zm-1.248-2.223a2.5 2.5 0 0 1 3.15 1.605l1.854 5.706a2.5 2.5 0 0 1-1.605 3.15l-.951.31a2.992 2.992 0 0 1-2.443-.265l-2.02.569a1 1 0 1 1-.541-1.926l1.212-.34-1.353-4.163L5 10.46a1 1 0 0 0-.567 1.233l.381 1.171a1 1 0 0 0 1.222.654l3.127-.88a1 1 0 1 1 .541 1.926l-3.127.88a3 3 0 0 1-3.665-1.961l-.38-1.172a3 3 0 0 1 1.7-3.697l9.374-3.897a3 3 0 0 1 2.02-2.285l.95-.31Z\" clip-rule=\"evenodd\"/><path fill=\"currentColor\" fill-rule=\"evenodd\" d=\"M12 12.5a1.5 1.5 0 1 0 0 3 1.5 1.5 0 0 0 0-3ZM8.5 14a3.5 3.5 0 1 1 7 0 3.5 3.5 0 0 1-7 0Z\" clip-rule=\"evenodd\"/></svg>",
#                     "required_features": [],
#                     "required_models": [],
#                     "required_conversation_modes": [
#                         "primary_assistant"
#                     ],
#                     "allow_in_temporary_chat": False,
#                     "composer_bar_button_info": {
#                         "disabled_text": "深入研究不可用",
#                         "tooltip_text": "对任何主题都有详细的见解",
#                         "announcement_key": "oai/apps/hasSeenComposerCaterpillarButtonTooltip",
#                         "nux_title": "您的个人研究员",
#                         "nux_description": "使用 ChatGPT 来研究购物、大概念、科学问题等内容。[了解更多](https://openai.com/index/introducing-deep-research/)",
#                         "rate_limit_reached_text": "本月限额已用完"
#                     },
#                     "suggested_prompt": {
#                         "theme": "#0088FF",
#                         "title": "深入研究",
#                         "subtitle": "探索宏大主题",
#                         "sort_order": 1,
#                         "badge": "新"
#                     },
#                     "regex_matches": None
#                 }
#             ]
#         }
#     return Response(content=json.dumps(system_hints, indent=4), media_type="application/json")


@app.post("/backend-api/edge")
async def edge():
    return Response(status_code=204)


if no_sentinel:
    openai_sentinel_tokens_cache = {}
    openai_sentinel_cookies_cache = {}

    @app.post("/backend-api/sentinel/chat-requirements")
    async def sentinel_chat_conversations(request: Request):
        token = request.headers.get("Authorization", "").replace("Bearer ", "")
        req_token = await get_real_req_token(token)
        access_token = await verify_token(req_token)
        fp = get_fp(req_token).copy()
        proxy_url = fp.pop("proxy_url", None)
        impersonate = fp.pop("impersonate", "safari15_3")
        user_agent = fp.get("user-agent",
                            "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0")

        host_url = random.choice(chatgpt_base_url_list) if chatgpt_base_url_list else "https://chatgpt.com"
        proof_token = None
        turnstile_token = None

        # headers = {
        #     key: value for key, value in request.headers.items()
        #     if (key.lower() not in ["host", "origin", "referer", "priority", "sec-ch-ua-platform", "sec-ch-ua",
        #                             "sec-ch-ua-mobile", "oai-device-id"] and key.lower() not in headers_reject_list)
        # }
        headers = {
            key: value for key, value in request.headers.items()
            if (key.lower() in headers_accept_list)
        }
        headers.update(fp)
        headers.update({"authorization": f"Bearer {access_token}"})
        session_id = hashlib.md5(req_token.encode()).hexdigest()
        proxy_url = proxy_url.replace("{}", session_id) if proxy_url else None
        client = Client(proxy=proxy_url, impersonate=impersonate)
        if sentinel_proxy_url_list:
            sentinel_proxy_url = random.choice(sentinel_proxy_url_list).replace("{}", session_id) if sentinel_proxy_url_list else None
            clients = Client(proxy=sentinel_proxy_url, impersonate=impersonate)
        else:
            clients = client

        try:
            config = get_config(user_agent, session_id)
            p = get_requirements_token(config)
            data = {'p': p}
            for cookie in openai_sentinel_cookies_cache.get(req_token, []):
                clients.session.cookies.set(**cookie)
            r = await clients.post(f'{host_url}/backend-api/sentinel/chat-requirements', headers=headers, json=data, timeout=10)
            oai_sc = r.cookies.get("oai-sc")
            if oai_sc:
                openai_sentinel_cookies_cache[req_token] = [{"name": "oai-sc", "value": oai_sc}]
            if r.status_code != 200:
                raise HTTPException(status_code=r.status_code, detail="Failed to get chat requirements")
            resp = r.json()
            turnstile = resp.get('turnstile', {})
            turnstile_required = turnstile.get('required')
            if turnstile_required:
                turnstile_dx = turnstile.get("dx")
                try:
                    if turnstile_solver_url:
                        res = await client.post(turnstile_solver_url,
                                                json={"url": "https://chatgpt.com", "p": p, "dx": turnstile_dx, "ua": user_agent})
                        turnstile_token = res.json().get("t")
                except Exception as e:
                    logger.info(f"Turnstile ignored: {e}")

            proofofwork = resp.get('proofofwork', {})
            proofofwork_required = proofofwork.get('required')
            if proofofwork_required:
                proofofwork_diff = proofofwork.get("difficulty")
                proofofwork_seed = proofofwork.get("seed")
                proof_token, solved = await run_in_threadpool(
                    get_answer_token, proofofwork_seed, proofofwork_diff, config
                )
                if not solved:
                    raise HTTPException(status_code=403, detail="Failed to solve proof of work")
            chat_token = resp.get('token')

            openai_sentinel_tokens_cache[req_token] = {
                "chat_token": chat_token,
                "proof_token": proof_token,
                "turnstile_token": turnstile_token
            }
        except Exception as e:
            logger.error(f"Sentinel failed: {e}")

        return {
            "arkose": {
                "dx": None,
                "required": False
            },
            "persona": "chatgpt-paid",
            "proofofwork": {
                "difficulty": None,
                "required": False,
                "seed": None
            },
            "token": str(uuid.uuid4()),
            "turnstile": {
                "dx": None,
                "required": False
            }
        }


    @app.post("/backend-alt/conversation")
    @app.post("/backend-api/conversation")
    async def chat_conversations(request: Request):
        token = request.headers.get("Authorization", "").replace("Bearer ", "")
        req_token = await get_real_req_token(token)
        access_token = await verify_token(req_token)
        fp = get_fp(req_token).copy()
        proxy_url = fp.pop("proxy_url", None)
        impersonate = fp.pop("impersonate", "safari15_3")
        user_agent = fp.get("user-agent",
                            "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0")

        host_url = random.choice(chatgpt_base_url_list) if chatgpt_base_url_list else "https://chatgpt.com"
        proof_token = None
        turnstile_token = None

        # headers = {
        #     key: value for key, value in request.headers.items()
        #     if (key.lower() not in ["host", "origin", "referer", "priority", "sec-ch-ua-platform", "sec-ch-ua",
        #                             "sec-ch-ua-mobile", "oai-device-id"] and key.lower() not in headers_reject_list)
        # }
        headers = {
            key: value for key, value in request.headers.items()
            if (key.lower() in headers_accept_list)
        }
        headers.update(fp)
        headers.update({"authorization": f"Bearer {access_token}"})

        try:
            session_id = hashlib.md5(req_token.encode()).hexdigest()
            proxy_url = proxy_url.replace("{}", session_id) if proxy_url else None
            client = Client(proxy=proxy_url, impersonate=impersonate)
            if sentinel_proxy_url_list:
                sentinel_proxy_url = random.choice(sentinel_proxy_url_list).replace("{}", session_id) if sentinel_proxy_url_list else None
                clients = Client(proxy=sentinel_proxy_url, impersonate=impersonate)
            else:
                clients = client

            sentinel_tokens = openai_sentinel_tokens_cache.get(req_token, {})
            openai_sentinel_tokens_cache.pop(req_token, None)
            if not sentinel_tokens:
                config = get_config(user_agent, session_id)
                p = get_requirements_token(config)
                data = {'p': p}
                r = await clients.post(f'{host_url}/backend-api/sentinel/chat-requirements', headers=headers, json=data,
                                       timeout=10)
                resp = r.json()
                turnstile = resp.get('turnstile', {})
                turnstile_required = turnstile.get('required')
                if turnstile_required:
                    turnstile_dx = turnstile.get("dx")
                    try:
                        if turnstile_solver_url:
                            res = await client.post(turnstile_solver_url,
                                                    json={"url": "https://chatgpt.com", "p": p, "dx": turnstile_dx, "ua": user_agent})
                            turnstile_token = res.json().get("t")
                    except Exception as e:
                        logger.info(f"Turnstile ignored: {e}")

                proofofwork = resp.get('proofofwork', {})
                proofofwork_required = proofofwork.get('required')
                if proofofwork_required:
                    proofofwork_diff = proofofwork.get("difficulty")
                    proofofwork_seed = proofofwork.get("seed")
                    proof_token, solved = await run_in_threadpool(
                        get_answer_token, proofofwork_seed, proofofwork_diff, config
                    )
                    if not solved:
                        raise HTTPException(status_code=403, detail="Failed to solve proof of work")
                chat_token = resp.get('token')
                headers.update({
                    "openai-sentinel-chat-requirements-token": chat_token,
                    "openai-sentinel-proof-token": proof_token,
                    "openai-sentinel-turnstile-token": turnstile_token,
                })
            else:
                headers.update({
                    "openai-sentinel-chat-requirements-token": sentinel_tokens.get("chat_token", ""),
                    "openai-sentinel-proof-token": sentinel_tokens.get("proof_token", ""),
                    "openai-sentinel-turnstile-token": sentinel_tokens.get("turnstile_token", "")
                })
        except Exception as e:
            logger.error(f"Sentinel failed: {e}")
            return Response(status_code=403, content="Sentinel failed")

        params = dict(request.query_params)
        data = await request.body()
        request_cookies = dict(request.cookies)

        async def c_close(client, clients):
            if client:
                await client.close()
                del client
            if clients:
                await clients.close()
                del clients

        history = True
        try:
            req_json = json.loads(data)
            history = not req_json.get("history_and_training_disabled", False)
        except Exception:
            pass
        if force_no_history:
            history = False
            req_json = json.loads(data)
            req_json["history_and_training_disabled"] = True
            data = json.dumps(req_json).encode("utf-8")

        background = BackgroundTask(c_close, client, clients)
        r = await client.post_stream(f"{host_url}{request.url.path}", params=params, headers=headers,
                                     cookies=request_cookies, data=data, stream=True, allow_redirects=False)
        rheaders = r.headers
        logger.info(f"Request token: {req_token}")
        logger.info(f"Request proxy: {proxy_url}")
        logger.info(f"Request UA: {user_agent}")
        logger.info(f"Request impersonate: {impersonate}")
        if x_sign:
            rheaders.update({"x-sign": x_sign})
        if 'stream' in rheaders.get("content-type", ""):
            conv_key = r.cookies.get("conv_key", "")
            response = StreamingResponse(content_generator(r, token, history), headers=rheaders,
                                         media_type=r.headers.get("content-type", ""), background=background)
            response.set_cookie("conv_key", value=conv_key)
            return response
        else:
            return Response(content=(await r.atext()), headers=rheaders, media_type=rheaders.get("content-type"),
                            status_code=r.status_code, background=background)


@app.api_route("/{path:path}", methods=["GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD", "PATCH", "TRACE"])
async def reverse_proxy(request: Request, path: str):
    token = request.headers.get("Authorization", "").replace("Bearer ", "")
    if len(token) != 45 and not token.startswith("eyJhbGciOi"):
        for banned_path in banned_paths:
            if re.match(banned_path, path):
                raise HTTPException(status_code=403, detail="Forbidden")

    for chatgpt_path in chatgpt_paths:
        if re.match(chatgpt_path, path):
            return await chatgpt_html(request)

    for redirect_path in redirect_paths:
        if re.match(redirect_path, path):
            redirect_url = str(request.base_url)
            response = RedirectResponse(url=f"{redirect_url}login", status_code=302)
            return response

    return await chatgpt_reverse_proxy(request, path)
````

## File: gateway/chatgpt.py
````python
import json
from urllib.parse import quote

from fastapi import Request
from fastapi.responses import HTMLResponse

from app import app, templates
from gateway.login import login_html
from utils.kv_utils import set_value_for_key_list

with open("templates/chatgpt_context_1.json", "r", encoding="utf-8") as f:
    chatgpt_context_1 = json.load(f)
with open("templates/chatgpt_context_2.json", "r", encoding="utf-8") as f:
    chatgpt_context_2 = json.load(f)



@app.get("/", response_class=HTMLResponse)
async def chatgpt_html(request: Request):
    token = request.query_params.get("token")
    if not token:
        token = request.cookies.get("token")
    if not token:
        return await login_html(request)

    if len(token) != 45 and not token.startswith("eyJhbGciOi"):
        token = quote(token)

    user_chatgpt_context_1 = chatgpt_context_1.copy()
    user_chatgpt_context_2 = chatgpt_context_2.copy()

    set_value_for_key_list(user_chatgpt_context_1, "accessToken", token)
    if request.cookies.get("oai-locale"):
        set_value_for_key_list(user_chatgpt_context_1, "locale", request.cookies.get("oai-locale"))
    else:
        accept_language = request.headers.get("accept-language")
        if accept_language:
            set_value_for_key_list(user_chatgpt_context_1, "locale", accept_language.split(",")[0])

    user_chatgpt_context_1 = json.dumps(user_chatgpt_context_1, separators=(',', ':'), ensure_ascii=False)
    user_chatgpt_context_2 = json.dumps(user_chatgpt_context_2, separators=(',', ':'), ensure_ascii=False)

    escaped_context_1 = user_chatgpt_context_1.replace("\\", "\\\\").replace('"', '\\"')
    escaped_context_2 = user_chatgpt_context_2.replace("\\", "\\\\").replace('"', '\\"')

    clear_localstorage_script = """
    <script>
        localStorage.clear();
    </script>
    """

    response = templates.TemplateResponse("chatgpt.html", {
        "request": request,
        "react_chatgpt_context_1": escaped_context_1,
        "react_chatgpt_context_2": escaped_context_2,
        "clear_localstorage_script": clear_localstorage_script
    })
    response.set_cookie("token", value=token, expires="Thu, 01 Jan 2099 00:00:00 GMT")
    return response
````

## File: gateway/gpts.py
````python
import json
from urllib.parse import quote

from fastapi import Request
from fastapi.responses import Response

from app import app
from gateway.chatgpt import chatgpt_html
from utils.kv_utils import set_value_for_key_list

with open("templates/gpts_context.json", "r", encoding="utf-8") as f:
    gpts_context = json.load(f)


@app.get("/gpts")
async def get_gpts(request: Request):
    return await chatgpt_html(request)

@app.get("/gpts.data")
async def get_gpts(request: Request):
    referrer = request.headers.get("referer", "")
    response_str = '[{"_1":2},"routes/gpts._index",{"_3":4},"data",{"_5":6,"_7":8},"kind","store","referrer","https://chatgpt.com/"]'
    response_str = response_str.replace("https://chatgpt.com/", referrer)
    return Response(content=response_str, media_type="text/x-script; charset=utf-8")


@app.get("/g/g-{gizmo_id}")
async def get_gizmo_json(request: Request, gizmo_id: str):
    params = request.query_params
    if params.get("_routes") == "routes/g.$gizmoId._index":
        token = request.cookies.get("token")
        if len(token) != 45 and not token.startswith("eyJhbGciOi"):
            token = quote(token)
        user_gpts_context = gpts_context.copy()
        set_value_for_key_list(user_gpts_context, "accessToken", token)
        response_str = json.dumps(user_gpts_context, separators=(',', ':'), ensure_ascii=False)
        return Response(content=response_str, media_type="text/x-script; charset=utf-8")
    else:
        return await chatgpt_html(request)
````

## File: gateway/login.py
````python
from fastapi import Request
from fastapi.responses import HTMLResponse

from app import app, templates


@app.get("/login", response_class=HTMLResponse)
async def login_html(request: Request):
    response = templates.TemplateResponse("login.html", {"request": request})
    return response
````

## File: gateway/reverseProxy.py
````python
import hashlib
import json
import random
import time
from datetime import datetime, timezone

from fastapi import Request, HTTPException
from fastapi.responses import StreamingResponse, Response
from starlette.background import BackgroundTask

import utils.globals as globals
from chatgpt.authorization import verify_token, get_req_token
from chatgpt.fp import get_fp
from utils.Client import Client
from utils.Logger import logger
from utils.configs import chatgpt_base_url_list, sentinel_proxy_url_list, force_no_history, file_host, voice_host


def generate_current_time():
    current_time = datetime.now(timezone.utc)
    formatted_time = current_time.isoformat(timespec='microseconds').replace('+00:00', 'Z')
    return formatted_time


headers_reject_list = [
    "x-real-ip",
    "x-forwarded-for",
    "x-forwarded-proto",
    "x-forwarded-port",
    "x-forwarded-host",
    "x-forwarded-server",
    "cf-warp-tag-id",
    "cf-visitor",
    "cf-ray",
    "cf-connecting-ip",
    "cf-ipcountry",
    "cdn-loop",
    "remote-host",
    "x-frame-options",
    "x-xss-protection",
    "x-content-type-options",
    "content-security-policy",
    "host",
    "cookie",
    "connection",
    "content-length",
    "content-encoding",
    "x-middleware-prefetch",
    "x-nextjs-data",
    "purpose",
    "x-forwarded-uri",
    "x-forwarded-path",
    "x-forwarded-method",
    "x-forwarded-protocol",
    "x-forwarded-scheme",
    "cf-request-id",
    "cf-worker",
    "cf-access-client-id",
    "cf-access-client-device-type",
    "cf-access-client-device-model",
    "cf-access-client-device-name",
    "cf-access-client-device-brand",
    "x-middleware-prefetch",
    "x-forwarded-for",
    "x-forwarded-host",
    "x-forwarded-proto",
    "x-forwarded-server",
    "x-real-ip",
    "x-forwarded-port",
    "cf-connecting-ip",
    "cf-ipcountry",
    "cf-ray",
    "cf-visitor",
]

headers_accept_list = [
    "openai-sentinel-chat-requirements-token",
    "openai-sentinel-proof-token",
    "openai-sentinel-turnstile-token",
    "accept",
    "authorization",
    "accept-encoding",
    "accept-language",
    "content-type",
    "oai-device-id",
    "oai-echo-logs",
    "oai-language",
    "sec-fetch-dest",
    "sec-fetch-mode",
    "sec-fetch-site",
]


async def get_real_req_token(token):
    req_token = get_req_token(token)
    if len(req_token) == 45 or req_token.startswith("eyJhbGciOi"):
        return req_token
    else:
        req_token = get_req_token("", token)
        return req_token


def save_conversation(token, conversation_id, title=None):
    if conversation_id not in globals.conversation_map:
        conversation_detail = {
            "id": conversation_id,
            "title": title,
            "create_time": generate_current_time(),
            "update_time": generate_current_time()
        }
        globals.conversation_map[conversation_id] = conversation_detail
    else:
        globals.conversation_map[conversation_id]["update_time"] = generate_current_time()
        if title:
            globals.conversation_map[conversation_id]["title"] = title
    if conversation_id not in globals.seed_map[token]["conversations"]:
        globals.seed_map[token]["conversations"].insert(0, conversation_id)
    else:
        globals.seed_map[token]["conversations"].remove(conversation_id)
        globals.seed_map[token]["conversations"].insert(0, conversation_id)
    with open(globals.CONVERSATION_MAP_FILE, "w", encoding="utf-8") as f:
        json.dump(globals.conversation_map, f, indent=4)
    with open(globals.SEED_MAP_FILE, "w", encoding="utf-8") as f:
        json.dump(globals.seed_map, f, indent=4)
    if title:
        logger.info(f"Conversation ID: {conversation_id}, Title: {title}")


async def content_generator(r, token, history=True):
    conversation_id = None
    title = None
    async for chunk in r.aiter_content():
        try:
            if history and (len(token) != 45 and not token.startswith("eyJhbGciOi")) and (not conversation_id or not title):
                chat_chunk = chunk.decode('utf-8')
                if not conversation_id or not title and chat_chunk.startswith("event: delta\n\ndata: {"):
                    chunk_data = chat_chunk[19:]
                    conversation_id = json.loads(chunk_data).get("v").get("conversation_id")
                    if conversation_id:
                        save_conversation(token, conversation_id)
                        title = globals.conversation_map[conversation_id].get("title")
                if chat_chunk.startswith("data: {"):
                    if "\n\nevent: delta" in chat_chunk:
                        index = chat_chunk.find("\n\nevent: delta")
                        chunk_data = chat_chunk[6:index]
                    elif "\n\ndata: {" in chat_chunk:
                        index = chat_chunk.find("\n\ndata: {")
                        chunk_data = chat_chunk[6:index]
                    else:
                        chunk_data = chat_chunk[6:]
                    chunk_data = chunk_data.strip()
                    if conversation_id is None:
                        conversation_id = json.loads(chunk_data).get("conversation_id")
                        if conversation_id:
                            save_conversation(token, conversation_id)
                            title = globals.conversation_map[conversation_id].get("title")
                    if title is None:
                        title = json.loads(chunk_data).get("title")
                        if title:
                            save_conversation(token, conversation_id, title)
        except Exception as e:
            # logger.error(e)
            # logger.error(chunk.decode('utf-8'))
            pass
        yield chunk


async def chatgpt_reverse_proxy(request: Request, path: str):
    try:
        origin_host = request.url.netloc
        if request.url.is_secure:
            petrol = "https"
        else:
            petrol = "http"
        if "x-forwarded-proto" in request.headers:
            petrol = request.headers["x-forwarded-proto"]
        if "cf-visitor" in request.headers:
            cf_visitor = json.loads(request.headers["cf-visitor"])
            petrol = cf_visitor.get("scheme", petrol)

        params = dict(request.query_params)
        request_cookies = dict(request.cookies)

        # headers = {
        #     key: value for key, value in request.headers.items()
        #     if (key.lower() not in ["host", "origin", "referer", "priority",
        #                             "oai-device-id"] and key.lower() not in headers_reject_list)
        # }
        headers = {
            key: value for key, value in request.headers.items()
            if (key.lower() in headers_accept_list)
        }

        base_url = random.choice(chatgpt_base_url_list) if chatgpt_base_url_list else "https://chatgpt.com"
        if "assets/" in path:
            base_url = "https://cdn.oaistatic.com"
        if "file-" in path and "backend-api" not in path:
            base_url = "https://files.oaiusercontent.com"
        if "v1/" in path:
            base_url = "https://ab.chatgpt.com"
        if "sandbox" in path:
            base_url = "https://web-sandbox.oaiusercontent.com"
            path = path.replace("sandbox/", "")

        token = headers.get("authorization", "").replace("Bearer ", "").strip()
        if token:
            req_token = await get_real_req_token(token)
            access_token = await verify_token(req_token)
            headers.update({"authorization": f"Bearer {access_token}"})

        cookie_token = request.cookies.get("token", "")
        req_token = await get_real_req_token(cookie_token)
        fp = get_fp(req_token).copy()

        session_id = hashlib.md5(req_token.encode()).hexdigest()

        proxy_url = fp.pop("proxy_url", None)
        impersonate = fp.pop("impersonate", "safari15_3")
        user_agent = fp.get("user-agent")
        headers.update(fp)

        headers.update({
            "accept-language": "en-US,en;q=0.9",
            "host": base_url.replace("https://", "").replace("http://", ""),
            "origin": base_url,
            "referer": f"{base_url}/"
        })
        if "v1/initialize" in path:
            headers.update({"user-agent": request.headers.get("user-agent")})
            if "statsig-api-key" not in headers:
                headers.update({
                    "statsig-sdk-type": "js-client",
                    "statsig-api-key": "client-tnE5GCU2F2cTxRiMbvTczMDT1jpwIigZHsZSdqiy4u",
                    "statsig-sdk-version": "5.1.0",
                    "statsig-client-time": int(time.time() * 1000),
                })

        data = await request.body()

        history = True
        if path.endswith("backend-api/conversation") or path.endswith("backend-alt/conversation"):
            try:
                req_json = json.loads(data)
                history = not req_json.get("history_and_training_disabled", False)
            except Exception:
                pass
            if force_no_history:
                history = False
                req_json = json.loads(data)
                req_json["history_and_training_disabled"] = True
                data = json.dumps(req_json).encode("utf-8")


        if "backend-api/sentinel/chat-requirements" in path and sentinel_proxy_url_list:
            sentinel_proxy_url = random.choice(sentinel_proxy_url_list).replace("{}", session_id) if sentinel_proxy_url_list else None
            client = Client(proxy=sentinel_proxy_url)
        else:
            proxy_url = proxy_url.replace("{}", session_id) if proxy_url else None
            client = Client(proxy=proxy_url, impersonate=impersonate)
        try:
            background = BackgroundTask(client.close)
            r = await client.request(request.method, f"{base_url}/{path}", params=params, headers=headers,
                                     cookies=request_cookies, data=data, stream=True, allow_redirects=False)
            if r.status_code == 307 or r.status_code == 302 or r.status_code == 301:
                return Response(status_code=307,
                                headers={"Location": r.headers.get("Location")
                                .replace("ab.chatgpt.com", origin_host)
                                .replace("chatgpt.com", origin_host)
                                .replace("cdn.oaistatic.com", origin_host)
                                .replace("https", petrol)}, background=background)
            elif 'stream' in r.headers.get("content-type", ""):
                logger.info(f"Request token: {req_token}")
                logger.info(f"Request proxy: {proxy_url}")
                logger.info(f"Request UA: {user_agent}")
                logger.info(f"Request impersonate: {impersonate}")
                conv_key = r.cookies.get("conv_key", "")
                response = StreamingResponse(content_generator(r, token, history), media_type=r.headers.get("content-type", ""),
                                  background=background)
                response.set_cookie("conv_key", value=conv_key)
                return response
            elif 'image' in r.headers.get("content-type", "") or "audio" in r.headers.get("content-type", "") or "video" in r.headers.get("content-type", ""):
                rheaders = dict(r.headers)
                response = Response(content=await r.acontent(), headers=rheaders,
                                        status_code=r.status_code, background=background)
                return response
            else:
                if path.endswith("backend-api/conversation") or path.endswith("backend-alt/conversation") or "/register-websocket" in path:
                    response = Response(content=(await r.acontent()), media_type=r.headers.get("content-type"),
                                        status_code=r.status_code, background=background)
                else:
                    content = await r.atext()
                    if "public-api/" in path:
                        content = (content
                                   .replace("https://ab.chatgpt.com", f"{petrol}://{origin_host}")
                                   .replace("https://cdn.oaistatic.com", f"{petrol}://{origin_host}")
                                   .replace("webrtc.chatgpt.com", voice_host if voice_host else "webrtc.chatgpt.com")
                                   .replace("files.oaiusercontent.com", file_host if file_host else "files.oaiusercontent.com")
                                   .replace("chatgpt.com/ces", f"{origin_host}/ces")
                                   )
                    else:
                        content = (content
                                   .replace("https://ab.chatgpt.com", f"{petrol}://{origin_host}")
                                   .replace("https://cdn.oaistatic.com", f"{petrol}://{origin_host}")
                                   .replace("webrtc.chatgpt.com", voice_host if voice_host else "webrtc.chatgpt.com")
                                   .replace("files.oaiusercontent.com", file_host if file_host else "files.oaiusercontent.com")
                                   .replace("web-sandbox.oaiusercontent.com", f"{origin_host}/sandbox")
                                   .replace("https://chatgpt.com", f"{petrol}://{origin_host}")
                                   .replace("chatgpt.com/ces", f"{origin_host}/ces")
                                   )
                    if base_url == "https://web-sandbox.oaiusercontent.com":
                        content = content.replace("/assets", "/sandbox/assets")
                    rheaders = dict(r.headers)
                    content_type = rheaders.get("content-type", "")
                    cache_control = rheaders.get("cache-control", "")
                    expires = rheaders.get("expires", "")
                    content_disposition = rheaders.get("content-disposition", "")
                    rheaders = {
                        "cache-control": cache_control,
                        "content-type": content_type,
                        "expires": expires,
                        "content-disposition": content_disposition
                    }
                    response = Response(content=content, headers=rheaders,
                                        status_code=r.status_code, background=background)
                return response
        except Exception as e:
            await client.close()
    except HTTPException as e:
        raise e
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
````

## File: gateway/route.py
````python

````

## File: gateway/share.py
````python
import hashlib
import json
import random
import time

import jwt
from fastapi import Request, HTTPException, Security
from fastapi.responses import Response
from fastapi.security import HTTPAuthorizationCredentials

import utils.globals as globals
from app import app, security_scheme
from chatgpt.authorization import verify_token
from chatgpt.fp import get_fp
from gateway.reverseProxy import get_real_req_token
from utils.Client import Client
from utils.Logger import logger
from utils.configs import proxy_url_list, chatgpt_base_url_list, authorization_list

base_headers = {
    'accept': '*/*',
    'accept-encoding': 'gzip, deflate, br, zstd',
    'accept-language': 'en-US,en;q=0.9',
    'content-type': 'application/json',
    'oai-language': 'en-US',
    'priority': 'u=1, i',
    'sec-fetch-dest': 'empty',
    'sec-fetch-mode': 'cors',
    'sec-fetch-site': 'same-origin',
}


def verify_authorization(bearer_token):
    if not bearer_token:
        raise HTTPException(status_code=401, detail="Authorization header is missing")
    if bearer_token not in authorization_list:
        raise HTTPException(status_code=401, detail="Invalid authorization")


@app.get("/seedtoken")
async def get_seedtoken(request: Request, credentials: HTTPAuthorizationCredentials = Security(security_scheme)):
    verify_authorization(credentials.credentials)
    try:
        params = request.query_params
        seed = params.get("seed")

        if seed:
            if seed not in globals.seed_map:
                raise HTTPException(status_code=404, detail=f"Seed '{seed}' not found")
            return {
                "status": "success",
                "data": {
                    "seed": seed,
                    "token": globals.seed_map[seed]["token"]
                }
            }

        token_map = {
            seed: data["token"]
            for seed, data in globals.seed_map.items()
        }
        return {"status": "success", "data": token_map}

    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Internal server error: {str(e)}")


@app.post("/seedtoken")
async def set_seedtoken(request: Request, credentials: HTTPAuthorizationCredentials = Security(security_scheme)):
    verify_authorization(credentials.credentials)
    data = await request.json()

    seed = data.get("seed")
    token = data.get("token")

    if seed not in globals.seed_map:
        globals.seed_map[seed] = {
            "token": token,
            "conversations": []
        }
    else:
        globals.seed_map[seed]["token"] = token

    with open(globals.SEED_MAP_FILE, "w", encoding="utf-8") as f:
        json.dump(globals.seed_map, f, indent=4)

    return {"status": "success", "message": "Token updated successfully"}


@app.delete("/seedtoken")
async def delete_seedtoken(request: Request, credentials: HTTPAuthorizationCredentials = Security(security_scheme)):
    verify_authorization(credentials.credentials)

    try:
        data = await request.json()
        seed = data.get("seed")

        if seed == "clear":
            globals.seed_map.clear()
            with open(globals.SEED_MAP_FILE, "w", encoding="utf-8") as f:
                json.dump(globals.seed_map, f, indent=4)
            return {"status": "success", "message": "All seeds deleted successfully"}

        if not seed:
            raise HTTPException(status_code=400, detail="Missing required field: seed")

        if seed not in globals.seed_map:
            raise HTTPException(status_code=404, detail=f"Seed '{seed}' not found")
        del globals.seed_map[seed]

        with open(globals.SEED_MAP_FILE, "w", encoding="utf-8") as f:
            json.dump(globals.seed_map, f, indent=4)

        return {
            "status": "success",
            "message": f"Seed '{seed}' deleted successfully"
        }

    except json.JSONDecodeError:
        raise HTTPException(status_code=400, detail="Invalid JSON data")
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Internal server error: {str(e)}")


async def chatgpt_account_check(access_token):
    auth_info = {}
    client = Client(proxy=random.choice(proxy_url_list) if proxy_url_list else None)
    try:
        host_url = random.choice(chatgpt_base_url_list) if chatgpt_base_url_list else "https://chatgpt.com"
        req_token = await get_real_req_token(access_token)
        access_token = await verify_token(req_token)
        fp = get_fp(req_token).copy()
        proxy_url = fp.pop("proxy_url", None)
        impersonate = fp.pop("impersonate", "safari15_3")

        headers = base_headers.copy()
        headers.update(fp)
        headers.update({"authorization": f"Bearer {access_token}"})

        session_id = hashlib.md5(access_token.encode()).hexdigest()
        proxy_url = random.choice(proxy_url_list).replace("{}", session_id) if proxy_url_list else None
        client = Client(proxy=proxy_url, impersonate=impersonate)
        r = await client.get(f"{host_url}/backend-api/models?history_and_training_disabled=false", headers=headers,
                             timeout=10)
        if r.status_code != 200:
            raise HTTPException(status_code=r.status_code, detail=r.text)
        models = r.json()
        r = await client.get(f"{host_url}/backend-api/accounts/check/v4-2023-04-27", headers=headers, timeout=10)
        if r.status_code != 200:
            raise HTTPException(status_code=r.status_code, detail=r.text)
        accounts_info = r.json()

        auth_info.update({"models": models["models"]})
        auth_info.update({"accounts_info": accounts_info})

        account_ordering = accounts_info.get("account_ordering", [])
        is_deactivated = True
        plan_type = None
        team_ids = []
        for account in account_ordering:
            this_is_deactivated = accounts_info['accounts'].get(account, {}).get("account", {}).get("is_deactivated", False)
            this_plan_type = accounts_info['accounts'].get(account, {}).get("account", {}).get("plan_type", "free")

            if not this_is_deactivated:
                is_deactivated = False

            if "team" in this_plan_type and not this_is_deactivated:
                plan_type = this_plan_type
                team_ids.append(account)
            elif plan_type is None:
                plan_type = this_plan_type

        auth_info.update({"accountCheckInfo": {
            "is_deactivated": is_deactivated,
            "plan_type": plan_type,
            "team_ids": team_ids
        }})

        return auth_info
    except Exception as e:
        logger.error(f"chatgpt_account_check: {e}")
        return {}
    finally:
        await client.close()


async def chatgpt_refresh(refresh_token):
    session_id = hashlib.md5(refresh_token.encode()).hexdigest()
    proxy_url = random.choice(proxy_url_list).replace("{}", session_id) if proxy_url_list else None
    client = Client(proxy=proxy_url)
    try:
        data = {
            "client_id": "pdlLIX2Y72MIl2rhLhTE9VV9bN905kBh",
            "grant_type": "refresh_token",
            "redirect_uri": "com.openai.chat://auth0.openai.com/ios/com.openai.chat/callback",
            "refresh_token": refresh_token
        }
        r = await client.post("https://auth0.openai.com/oauth/token", json=data, timeout=10)
        if r.status_code != 200:
            raise HTTPException(status_code=r.status_code, detail=r.text)
        res = r.json()
        auth_info = {}
        auth_info.update(res)
        auth_info.update({"refresh_token": refresh_token})
        auth_info.update({"accessToken": res.get("access_token", "")})
        return auth_info
    except Exception as e:
        logger.error(f"chatgpt_refresh: {e}")
        return {}
    finally:
        await client.close()


@app.post("/auth/refresh")
async def refresh(request: Request):
    auth_info = {}
    form_data = await request.form()

    auth_info.update(form_data)

    access_token = auth_info.get("access_token", auth_info.get("accessToken", ""))
    refresh_token = auth_info.get("refresh_token", "")

    if not refresh_token and not access_token:
        raise HTTPException(status_code=401, detail="refresh_token or access_token is required")

    need_refresh = True
    if access_token:
        try:
            access_token_info = jwt.decode(access_token, options={"verify_signature": False})
            exp = access_token_info.get("exp", 0)
            if exp > int(time.time()) + 60 * 60 * 24 * 5:
                need_refresh = False
        except Exception as e:
            logger.error(f"access_token: {e}")

    if refresh_token and need_refresh:
        chatgpt_refresh_info = await chatgpt_refresh(refresh_token)
        if chatgpt_refresh_info:
            auth_info.update(chatgpt_refresh_info)
            access_token = auth_info.get("accessToken", "")
            account_check_info = await chatgpt_account_check(access_token)
            if account_check_info:
                auth_info.update(account_check_info)
                auth_info.update({"accessToken": access_token})
                return Response(content=json.dumps(auth_info), media_type="application/json")
    elif access_token:
        account_check_info = await chatgpt_account_check(access_token)
        if account_check_info:
            auth_info.update(account_check_info)
            auth_info.update({"accessToken": access_token})
            return Response(content=json.dumps(auth_info), media_type="application/json")

    raise HTTPException(status_code=401, detail="Unauthorized")
````

## File: gateway/v1.py
````python
import json

from fastapi import Request
from fastapi.responses import Response

from app import app
from gateway.reverseProxy import chatgpt_reverse_proxy
from utils.kv_utils import set_value_for_key_dict

with open("templates/initialize.json", "r") as f:
    initialize_json = json.load(f)


@app.post("/v1/initialize")
async def initialize(request: Request):
    initialize_response = (await chatgpt_reverse_proxy(request, f"v1/initialize"))
    if not initialize_response:
        return Response(status_code=204)
    initialize_str = initialize_response.body.decode('utf-8')
    if not initialize_str:
        return Response(status_code=204)
    initialize_json = json.loads(initialize_str)
    set_value_for_key_dict(initialize_json, "ip", "8.8.8.8")
    set_value_for_key_dict(initialize_json, "country", "US")
    return Response(content=json.dumps(initialize_json, indent=4), media_type="application/json")


@app.post("/v1/rgstr")
async def rgstr():
    return Response(status_code=202, content=json.dumps({"success": True}, indent=4), media_type="application/json")


@app.get("/ces/v1/projects/oai/settings")
async def ces_v1_projects_oai_settings():
    return Response(status_code=200, content=json.dumps({"integrations":{"Segment.io":{"apiHost":"chatgpt.com/ces/v1","apiKey":"oai"}}}, indent=4), media_type="application/json")


@app.post("/ces/v1/{path:path}")
async def ces_v1():
    return Response(status_code=202, content=json.dumps({"success": True}, indent=4), media_type="application/json")


@app.post("/ces/statsc/flush")
async def ces_v1():
    return Response(status_code=200, content=json.dumps({"success": True}, indent=4), media_type="application/json")
````

## File: LICENSE
````
MIT License

Copyright (c) 2024 aurora-develop

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

## File: README.md
````markdown
# CHAT2API

🤖 一个简单的 ChatGPT TO API 代理

🌟 无需账号即可使用免费、无限的 `GPT-3.5`

💥 支持 AccessToken 使用账号，支持 `O3-mini/high`、`O1/mini/Pro`、`GPT-4/4o/mini`、`GPTs`

🔍 回复格式与真实 API 完全一致，适配几乎所有客户端

👮 配套用户管理端[Chat-Share](https://github.com/h88782481/Chat-Share)使用前需提前配置好环境变量（ENABLE_GATEWAY设置为True，AUTO_SEED设置为False）


## 交流群

[https://t.me/chat2api](https://t.me/chat2api)

要提问请先阅读完仓库文档，尤其是常见问题部分。

提问时请提供：

1. 启动日志截图（敏感信息打码，包括环境变量和版本号）
2. 报错的日志信息（敏感信息打码）
3. 接口返回的状态码和响应体

## 功能

### 最新版本号存于 `version.txt`

### 逆向API 功能
> - [x] 流式、非流式传输
> - [x] 免登录 GPT-3.5 对话
> - [x] GPT-3.5 模型对话（传入模型名不包含 gpt-4，则默认使用 gpt-3.5，也就是 text-davinci-002-render-sha）
> - [x] GPT-4 系列模型对话（传入模型名包含: gpt-4，gpt-4o，gpt-4o-mini，gpt-4-moblie 即可使用对应模型，需传入 AccessToken）
> - [x] O1 系列模型对话（传入模型名包含 o1-preview，o1-mini 即可使用对应模型，需传入 AccessToken）
> - [x] GPT-4 模型画图、代码、联网
> - [x] 支持 GPTs（传入模型名：gpt-4-gizmo-g-*）
> - [x] 支持 Team Plus 账号（需传入 team account id）
> - [x] 上传图片、文件（格式为 API 对应格式，支持 URL 和 base64）
> - [x] 可作为网关使用，可多机分布部署
> - [x] 多账号轮询，同时支持 `AccessToken` 和 `RefreshToken`
> - [x] 请求失败重试，自动轮询下一个 Token
> - [x] Tokens 管理，支持上传、清除
> - [x] 定时使用 `RefreshToken` 刷新 `AccessToken` / 每次启动将会全部非强制刷新一次，每4天晚上3点全部强制刷新一次。
> - [x] 支持文件下载，需要开启历史记录
> - [x] 支持 `O3-mini/high`、`O1/mini/Pro` 等模型推理过程输出

### 官网镜像 功能
> - [x] 支持官网原生镜像
> - [x] 后台账号池随机抽取，`Seed` 设置随机账号
> - [x] 输入 `RefreshToken` 或 `AccessToken` 直接登录使用
> - [x] 支持 `O3-mini/high`、`O1/mini/Pro`、`GPT-4/4o/mini`
> - [x] 敏感信息接口禁用、部分设置接口禁用
> - [x] /login 登录页面，注销后自动跳转到登录页面
> - [x] /?token=xxx 直接登录, xxx 为 `RefreshToken` 或 `AccessToken` 或 `SeedToken` (随机种子)
> - [x] 支持不同 SeedToken 会话隔离
> - [x] 支持 `GPTs` 商店
> - [x] 支持 `DeepReaserch`、`Canvas` 等官网独有功能
> - [x] 支持切换各国语言


> TODO
> - [ ] 暂无，欢迎提 `issue`

## 逆向API

完全 `OpenAI` 格式的 API ，支持传入 `AccessToken` 或 `RefreshToken`，可用 GPT-4, GPT-4o, GPT-4o-Mini, GPTs, O1-Pro, O1, O1-Mini, O3-Mini, O3-Mini-High：

```bash
curl --location 'http://127.0.0.1:5005/v1/chat/completions' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer {{Token}}' \
--data '{
     "model": "gpt-3.5-turbo",
     "messages": [{"role": "user", "content": "Say this is a test!"}],
     "stream": true
   }'
```

将你账号的 `AccessToken` 或 `RefreshToken` 作为 `{{ Token }}` 传入。
也可填写你设置的环境变量 `Authorization` 的值, 将会随机选择后台账号

如果有team账号，可以传入 `ChatGPT-Account-ID`，使用 Team 工作区：

- 传入方式一：
`headers` 中传入 `ChatGPT-Account-ID`值

- 传入方式二：
`Authorization: Bearer <AccessToken 或 RefreshToken>,<ChatGPT-Account-ID>`

如果设置了 `AUTHORIZATION` 环境变量，可以将设置的值作为 `{{ Token }}` 传入进行多 Tokens 轮询。

> - `AccessToken` 获取: chatgpt官网登录后，再打开 [https://chatgpt.com/api/auth/session](https://chatgpt.com/api/auth/session) 获取 `accessToken` 这个值。
> - `RefreshToken` 获取: 此处不提供获取方法。
> - 免登录 gpt-3.5 无需传入 Token。

## Tokens 管理

1. 配置环境变量 `AUTHORIZATION` 作为 `授权码` ，然后运行程序。

2. 访问 `/tokens` 或者 `/{api_prefix}/tokens` 可以查看现有 Tokens 数量，也可以上传新的 Tokens ，或者清空 Tokens。

3. 请求时传入 `AUTHORIZATION` 中配置的 `授权码` 即可使用轮询的Tokens进行对话

![tokens.png](docs/tokens.png)

## 官网原生镜像

1. 配置环境变量 `ENABLE_GATEWAY` 为 `true`，然后运行程序, 注意开启后别人也可以直接通过域名访问你的网关。

2. 在 Tokens 管理页面上传 `RefreshToken` 或 `AccessToken`

3. 访问 `/login` 到登录页面

![login.png](docs/login.png)

4. 进入官网原生镜像页面使用

![chatgpt.png](docs/chatgpt.png)

## 环境变量

每个环境变量都有默认值，如果不懂环境变量的含义，请不要设置，更不要传空值，字符串无需引号。

| 分类   | 变量名               | 示例值                                                         | 默认值                   | 描述                                                           |
|------|-------------------|-------------------------------------------------------------|-----------------------|--------------------------------------------------------------|
| 安全相关 | API_PREFIX        | `your_prefix`                                               | `None`                | API 前缀密码，不设置容易被人访问，设置后需请求 `/your_prefix/v1/chat/completions` |
|      | AUTHORIZATION     | `your_first_authorization`,<br/>`your_second_authorization` | `[]`                  | 你自己为使用多账号轮询 Tokens 设置的授权码，英文逗号分隔                             |
|      | AUTH_KEY          | `your_auth_key`                                             | `None`                | 私人网关需要加`auth_key`请求头才设置该项                                    |
| 请求相关 | CHATGPT_BASE_URL  | `https://chatgpt.com`                                       | `https://chatgpt.com` | ChatGPT 网关地址，设置后会改变请求的网站，多个网关用逗号分隔                           |
|      | PROXY_URL         | `http://ip:port`,<br/>`http://username:password@ip:port`    | `[]`                  | 全局代理 URL，出 403 时启用，多个代理用逗号分隔                                 |
|      | EXPORT_PROXY_URL  | `http://ip:port`或<br/>`http://username:password@ip:port`    | `None`                | 出口代理 URL，防止请求图片和文件时泄漏源站 ip                                   |
| 功能相关 | HISTORY_DISABLED  | `true`                                                      | `true`                | 是否不保存聊天记录并返回 conversation_id                                 |
|      | POW_DIFFICULTY    | `00003a`                                                    | `00003a`              | 要解决的工作量证明难度，不懂别设置                                            |
|      | RETRY_TIMES       | `3`                                                         | `3`                   | 出错重试次数，使用 `AUTHORIZATION` 会自动随机/轮询下一个账号                      |
|      | CONVERSATION_ONLY | `false`                                                     | `false`               | 是否直接使用对话接口，如果你用的网关支持自动解决 `POW` 才启用                           |
|      | ENABLE_LIMIT      | `true`                                                      | `true`                | 开启后不尝试突破官方次数限制，尽可能防止封号                                       |
|      | UPLOAD_BY_URL     | `false`                                                     | `false`               | 开启后按照 `URL+空格+正文` 进行对话，自动解析 URL 内容并上传，多个 URL 用空格分隔           |
|      | SCHEDULED_REFRESH | `false`                                                     | `false`               | 是否定时刷新 `AccessToken` ，开启后每次启动程序将会全部非强制刷新一次，每4天晚上3点全部强制刷新一次。  |
|      | RANDOM_TOKEN      | `true`                                                      | `true`                | 是否随机选取后台 `Token` ，开启后随机后台账号，关闭后为顺序轮询                         |
| 网关功能 | ENABLE_GATEWAY    | `false`                                                     | `false`               | 是否启用网关模式，开启后可以使用镜像站，但也将会不设防                                  |
|      | AUTO_SEED          | `false`                                                     | `true`               | 是否启用随机账号模式，默认启用，输入`seed`后随机匹配后台`Token`。关闭之后需要手动对接接口，来进行`Token`管控。    |

## 部署

### Zeabur 部署

[![Deploy on Zeabur](https://zeabur.com/button.svg)](https://zeabur.com/templates/6HEGIZ?referralCode=LanQian528)

### 直接部署

```bash
git clone https://github.com/LanQian528/chat2api
cd chat2api
pip install -r requirements.txt
python app.py
```

### Docker 部署

您需要安装 Docker 和 Docker Compose。

```bash
docker run -d \
  --name chat2api \
  -p 5005:5005 \
  lanqian528/chat2api:latest
```

### (推荐，可用 PLUS 账号) Docker Compose 部署

创建一个新的目录，例如 chat2api，并进入该目录：

```bash
mkdir chat2api
cd chat2api
```

在此目录中下载库中的 docker-compose.yml 文件：

```bash
wget https://raw.githubusercontent.com/LanQian528/chat2api/main/docker-compose-warp.yml
```

修改 docker-compose-warp.yml 文件中的环境变量，保存后：

```bash
docker-compose up -d
```


## 常见问题

> - 错误代码：
>   - `401`：当前 IP 不支持免登录，请尝试更换 IP 地址，或者在环境变量 `PROXY_URL` 中设置代理，或者你的身份验证失败。
>   - `403`：请在日志中查看具体报错信息。
>   - `429`：当前 IP 请求1小时内请求超过限制，请稍后再试，或更换 IP。
>   - `500`：服务器内部错误，请求失败。
>   - `502`：服务器网关错误，或网络不可用，请尝试更换网络环境。

> - 已知情况：
>   - 日本 IP 很多不支持免登，免登 GPT-3.5 建议使用美国 IP。
>   - 99%的账号都支持免费 `GPT-4o` ，但根据 IP 地区开启，目前日本和新加坡 IP 已知开启概率较大。

> - 环境变量 `AUTHORIZATION` 是什么？
>   - 是一个自己给 chat2api 设置的一个身份验证，设置后才可使用已保存的 Tokens 轮询，请求时当作 `APIKEY` 传入。
> - AccessToken 如何获取？
>   - chatgpt官网登录后，再打开 [https://chatgpt.com/api/auth/session](https://chatgpt.com/api/auth/session) 获取 `accessToken` 这个值。


## License

MIT License
````

## File: requirements.txt
````
fastapi==0.115.3
python-multipart==0.0.13
curl_cffi==0.7.3
uvicorn
tiktoken
python-dotenv
websockets
pillow
pybase64
jinja2
APScheduler
ua-generator
pyjwt
diskcache
````

## File: templates/chatgpt_context_1.json
````json
[
    {
        "_1": 2,
        "_1868": -5,
        "_1869": -5
    },
    "loaderData",
    {
        "_3": 4,
        "_1862": 1863,
        "_1867": -5
    },
    "root",
    {
        "_5": 6,
        "_7": 8,
        "_13": 14,
        "_1856": 34,
        "_1857": 34,
        "_1858": 34,
        "_1859": -7,
        "_1860": 1861
    },
    "rq:[\"account-status\"]",
    [
        "P",
        6
    ],
    "dd",
    {
        "_9": 10,
        "_11": 12
    },
    "traceId",
    "14093928210743749173",
    "traceTime",
    1742986866029,
    "clientBootstrap",
    {
        "_15": 16,
        "_17": 18,
        "_19": 20,
        "_58": 59,
        "_60": 61,
        "_62": 63,
        "_1842": 71,
        "_1843": -5,
        "_1844": 865,
        "_1845": 1839,
        "_1846": 1847,
        "_1848": 1849,
        "_1850": 1851,
        "_1852": 34,
        "_1853": 34,
        "_1854": 34,
        "_1855": 34
    },
    "authStatus",
    "logged_in",
    "session",
    {
        "_19": 20,
        "_39": 40,
        "_41": 42,
        "_48": 49,
        "_50": 51,
        "_52": 53
    },
    "user",
    {
        "_21": 22,
        "_23": 24,
        "_25": 24,
        "_26": 27,
        "_28": 27,
        "_29": 30,
        "_31": 32,
        "_33": 34,
        "_35": 36,
        "_37": 38
    },
    "id",
    "user-chatgpt",
    "name",
    "chatgpt",
    "email",
    "chatgpt@openai.com",
    "https://s.gravatar.com/avatar/5edae94250bff28c50456d715798421e?s=480&r=pg&d=https%3A%2F%2Fcdn.auth0.com%2Favatars%2Fli.png",
    "picture",
    "idp",
    "auth0",
    "iat",
    1742986864,
    "mfa",
    false,
    "groups",
    [],
    "intercom_hash",
    "2bfe17418ed9bd0db1924962d30f2ed4adcede4bc9a81ac41ce9d623ab3b4de5",
    "expires",
    "2025-06-24T11:01:06.009Z",
    "account",
    {
        "_21": 43,
        "_44": 45,
        "_46": 47
    },
    "17904ad0-df88-4a0f-917b-501d958eceda",
    "planType",
    "pro",
    "structure",
    "personal",
    "accessToken",
    "",
    "authProvider",
    "openai",
    "rumViewTags",
    {
        "_54": 55
    },
    "light_account",
    {
        "_56": 34,
        "_57": -7
    },
    "fetched",
    "reason",
    "cluster",
    "unified-9",
    "locale",
    "zh-CN",
    "statsig",
    {
        "_64": 65,
        "_875": 876
    },
    "classic",
    {
        "_66": 67,
        "_360": 361,
        "_456": 457,
        "_837": 838,
        "_839": 71,
        "_840": 841,
        "_842": 843,
        "_848": 849,
        "_850": 851,
        "_861": 862,
        "_19": 863
    },
    "feature_gates",
    {
        "_68": 69,
        "_86": 87,
        "_92": 93,
        "_99": 100,
        "_113": 114,
        "_117": 118,
        "_121": 122,
        "_125": 126,
        "_128": 129,
        "_132": 133,
        "_142": 143,
        "_146": 147,
        "_150": 151,
        "_153": 154,
        "_157": 158,
        "_161": 162,
        "_164": 165,
        "_168": 169,
        "_172": 173,
        "_182": 183,
        "_185": 186,
        "_189": 190,
        "_196": 197,
        "_199": 200,
        "_97": 204,
        "_206": 207,
        "_216": 217,
        "_222": 223,
        "_226": 227,
        "_232": 233,
        "_237": 238,
        "_194": 240,
        "_242": 243,
        "_245": 246,
        "_249": 250,
        "_255": 256,
        "_258": 259,
        "_262": 263,
        "_266": 267,
        "_270": 271,
        "_274": 275,
        "_278": 279,
        "_281": 282,
        "_285": 286,
        "_290": 291,
        "_294": 295,
        "_297": 298,
        "_303": 304,
        "_310": 311,
        "_316": 317,
        "_211": 319,
        "_321": 322,
        "_325": 326,
        "_329": 330,
        "_104": 332,
        "_334": 335,
        "_345": 346,
        "_343": 349,
        "_353": 354,
        "_356": 357
    },
    "61299031",
    {
        "_23": 68,
        "_70": 71,
        "_72": 73,
        "_74": 75
    },
    "value",
    true,
    "rule_id",
    "2wrvcqZBGOdzYtk4c8rQxP",
    "secondary_exposures",
    [
        76,
        83
    ],
    {
        "_77": 78,
        "_79": 80,
        "_81": 82
    },
    "gate",
    "44045625",
    "gateValue",
    "true",
    "ruleID",
    "1vGfaAvyQ4VnZ5Y0UnCsbl:100.00:5",
    {
        "_77": 84,
        "_79": 80,
        "_81": 85
    },
    "1259585210",
    "3cQqufsn9EF8iqIPZFNiE8:100.00:4",
    "80186230",
    {
        "_23": 86,
        "_70": 71,
        "_72": 88,
        "_74": 89
    },
    "7thMqF7L1NKFaEvg1NsH7E",
    [
        90,
        91
    ],
    {
        "_77": 78,
        "_79": 80,
        "_81": 82
    },
    {
        "_77": 84,
        "_79": 80,
        "_81": 85
    },
    "174366048",
    {
        "_23": 92,
        "_70": 71,
        "_72": 94,
        "_74": 95
    },
    "bhPM7FsN2H1vnBUrxrg6v:100.00:3",
    [
        96
    ],
    {
        "_77": 97,
        "_79": 80,
        "_81": 98
    },
    "1923022511",
    "6VUF6Z1JaUKZF7RS6uSjUu:100.00:6",
    "232791851",
    {
        "_23": 99,
        "_70": 34,
        "_72": 101,
        "_74": 102
    },
    "default",
    [
        103,
        106,
        109
    ],
    {
        "_77": 104,
        "_79": 80,
        "_81": 105
    },
    "3922476776",
    "1DS1QvDa6IFq9C1oJfgtU9",
    {
        "_77": 107,
        "_79": 80,
        "_81": 108
    },
    "749124420",
    "2MQYHJjfKwcTr14d1bOuVH:100.00:2",
    {
        "_77": 110,
        "_79": 111,
        "_81": 112
    },
    "566128514",
    "false",
    "4P1FctCTa3aaKSskEnEeMt",
    "374768818",
    {
        "_23": 113,
        "_70": 71,
        "_72": 115,
        "_74": 116
    },
    "wA7D0MWpe3uCf9HA5KeEi",
    [],
    "491279851",
    {
        "_23": 117,
        "_70": 71,
        "_72": 119,
        "_74": 120
    },
    "4qtiGR7vlvMtZnfSlXM5RN:100.00:12",
    [],
    "507664831",
    {
        "_23": 121,
        "_70": 34,
        "_72": 123,
        "_74": 124
    },
    "4SZ1s8XXvwaDrAV1l6wIro",
    [],
    "645560164",
    {
        "_23": 125,
        "_70": 34,
        "_72": 101,
        "_74": 127
    },
    [],
    "773249106",
    {
        "_23": 128,
        "_70": 34,
        "_72": 130,
        "_74": 131
    },
    "1kGO9xYmxaBS2V2H3LcQuG",
    [],
    "989108178",
    {
        "_23": 132,
        "_70": 34,
        "_72": 134,
        "_74": 135
    },
    "4sTodKrNyByM4guZ68MORR",
    [
        136,
        139
    ],
    {
        "_77": 137,
        "_79": 111,
        "_81": 138
    },
    "1457171347",
    "2EjTipm6C4kk4fuvcHMzZe",
    {
        "_77": 140,
        "_79": 80,
        "_81": 141
    },
    "1426009137",
    "4C2vO0R7mvnCZvl1HDBExp:30.00:5",
    "1028682714",
    {
        "_23": 142,
        "_70": 71,
        "_72": 144,
        "_74": 145
    },
    "735n03snBvba4AEhd2Qwqu:100.00:3",
    [],
    "1072178956",
    {
        "_23": 146,
        "_70": 34,
        "_72": 148,
        "_74": 149
    },
    "4m8JwKa5kCi9HNf1ScZepj",
    [],
    "1242184140",
    {
        "_23": 150,
        "_70": 34,
        "_72": 101,
        "_74": 152
    },
    [],
    "1318146997",
    {
        "_23": 153,
        "_70": 71,
        "_72": 155,
        "_74": 156
    },
    "2AclmEgqaQBVFbxz37XKzy:100.00:5",
    [],
    "1393076427",
    {
        "_23": 157,
        "_70": 71,
        "_72": 159,
        "_74": 160
    },
    "disabled",
    [],
    "1508312659",
    {
        "_23": 161,
        "_70": 34,
        "_72": 101,
        "_74": 163
    },
    [],
    "1578703058",
    {
        "_23": 164,
        "_70": 71,
        "_72": 166,
        "_74": 167
    },
    "2l4nEVMUnPuXkgprUm5zzs:100.00:4",
    [],
    "1611573287",
    {
        "_23": 168,
        "_70": 71,
        "_72": 170,
        "_74": 171
    },
    "159rwM3sBnviE9XWH24azn:100.00:2",
    [],
    "1719651090",
    {
        "_23": 172,
        "_70": 71,
        "_72": 174,
        "_74": 175
    },
    "60QaTyBFJYTakinhLvhAM9",
    [
        176,
        179
    ],
    {
        "_77": 177,
        "_79": 80,
        "_81": 178
    },
    "1616485584",
    "2PP6pudW64Hn7katvazhAx:100.00:5",
    {
        "_77": 180,
        "_79": 80,
        "_81": 181
    },
    "1034043359",
    "4bd3o553p0ZCRkFmipROd8",
    "1804926979",
    {
        "_23": 182,
        "_70": 34,
        "_72": 101,
        "_74": 184
    },
    [],
    "1825130190",
    {
        "_23": 185,
        "_70": 71,
        "_72": 187,
        "_74": 188
    },
    "Nef2uMceNUF9U3ZYwSbpD",
    [],
    "1847911009",
    {
        "_23": 189,
        "_70": 34,
        "_72": 191,
        "_74": 192
    },
    "5OIO2mI7iQiPRReG1jZ4c2:0.00:7",
    [
        193
    ],
    {
        "_77": 194,
        "_79": 80,
        "_81": 195
    },
    "2304807207",
    "xhzqzk6zPqMb3Qs4GVvJu:100.00:5",
    "1855896025",
    {
        "_23": 196,
        "_70": 34,
        "_72": 101,
        "_74": 198
    },
    [],
    "1902899872",
    {
        "_23": 199,
        "_70": 71,
        "_72": 201,
        "_74": 202
    },
    "58UOuEcFwyqlorfhrWQLlE",
    [
        203
    ],
    {
        "_77": 194,
        "_79": 80,
        "_81": 195
    },
    {
        "_23": 97,
        "_70": 71,
        "_72": 98,
        "_74": 205
    },
    [],
    "1988730211",
    {
        "_23": 206,
        "_70": 71,
        "_72": 208,
        "_74": 209
    },
    "6B9O1B3eHKElKWCUfbcvBL",
    [
        210,
        213
    ],
    {
        "_77": 211,
        "_79": 80,
        "_81": 212
    },
    "3780975974",
    "48uk8ZYa2RpJzkpIyOmqP0:100.00:5",
    {
        "_77": 214,
        "_79": 80,
        "_81": 215
    },
    "3733089528",
    "3vtzosKkaPCfPysd7yBTSf",
    "2044826081",
    {
        "_23": 216,
        "_70": 71,
        "_72": 218,
        "_74": 219
    },
    "6MpInoEzkXvXVvodNQCQWs",
    [
        220,
        221
    ],
    {
        "_77": 78,
        "_79": 80,
        "_81": 82
    },
    {
        "_77": 84,
        "_79": 80,
        "_81": 85
    },
    "2091463435",
    {
        "_23": 222,
        "_70": 71,
        "_72": 224,
        "_74": 225
    },
    "5t78GUS68KOn3bHZd8z7ii:100.00:1",
    [],
    "2113934735",
    {
        "_23": 226,
        "_70": 34,
        "_72": 101,
        "_74": 228
    },
    [
        229,
        230,
        231
    ],
    {
        "_77": 177,
        "_79": 80,
        "_81": 178
    },
    {
        "_77": 180,
        "_79": 80,
        "_81": 181
    },
    {
        "_77": 172,
        "_79": 80,
        "_81": 174
    },
    "2256850471",
    {
        "_23": 232,
        "_70": 71,
        "_72": 234,
        "_74": 235
    },
    "IqxordbUxF1Fkg4gfExiY:100.00:1",
    [
        236
    ],
    {
        "_77": 185,
        "_79": 80,
        "_81": 187
    },
    "2293185713",
    {
        "_23": 237,
        "_70": 34,
        "_72": 101,
        "_74": 239
    },
    [],
    {
        "_23": 194,
        "_70": 71,
        "_72": 195,
        "_74": 241
    },
    [],
    "2311599525",
    {
        "_23": 242,
        "_70": 34,
        "_72": 101,
        "_74": 244
    },
    [],
    "2335877601",
    {
        "_23": 245,
        "_70": 34,
        "_72": 247,
        "_74": 248
    },
    "6NQcdu7pgfp18Sq2tfBC6q",
    [],
    "2454940646",
    {
        "_23": 249,
        "_70": 71,
        "_72": 251,
        "_74": 252
    },
    "zol8dYvq8kKfRbOgcM0IF",
    [
        253,
        254
    ],
    {
        "_77": 211,
        "_79": 80,
        "_81": 212
    },
    {
        "_77": 214,
        "_79": 80,
        "_81": 215
    },
    "2494375100",
    {
        "_23": 255,
        "_70": 34,
        "_72": 101,
        "_74": 257
    },
    [],
    "2562876640",
    {
        "_23": 258,
        "_70": 71,
        "_72": 260,
        "_74": 261
    },
    "326czTZeZ0RX0ypR0c5Bb6:100.00:15",
    [],
    "2607001979",
    {
        "_23": 262,
        "_70": 34,
        "_72": 264,
        "_74": 265
    },
    "35jfNEnEKwGsryxcwFhAKz",
    [],
    "2687575887",
    {
        "_23": 266,
        "_70": 71,
        "_72": 268,
        "_74": 269
    },
    "10cvQmwrcZvpWBFlZgn8pZ",
    [],
    "2756095923",
    {
        "_23": 270,
        "_70": 71,
        "_72": 272,
        "_74": 273
    },
    "6jPp6nW1wQVJbfY0uwQgmv:100.00:1",
    [],
    "2868048419",
    {
        "_23": 274,
        "_70": 34,
        "_72": 276,
        "_74": 277
    },
    "7iUNAbafRQfKTvYI2mmFZB",
    [],
    "3054422710",
    {
        "_23": 278,
        "_70": 71,
        "_72": 159,
        "_74": 280
    },
    [],
    "3286474446",
    {
        "_23": 281,
        "_70": 71,
        "_72": 283,
        "_74": 284
    },
    "2a7wA6tOQ5GPb7WIr1SU1A:100.00:1",
    [],
    "3325813340",
    {
        "_23": 285,
        "_70": 71,
        "_72": 287,
        "_74": 288
    },
    "37GsRLj07CqERPyHBn4o5L",
    [
        289
    ],
    {
        "_77": 194,
        "_79": 80,
        "_81": 195
    },
    "3342258807",
    {
        "_23": 290,
        "_70": 34,
        "_72": 292,
        "_74": 293
    },
    "3m0ycr0cMQOm6eMQQjgyp9",
    [],
    "3376455464",
    {
        "_23": 294,
        "_70": 34,
        "_72": 101,
        "_74": 296
    },
    [],
    "3468624635",
    {
        "_23": 297,
        "_70": 34,
        "_72": 101,
        "_74": 299
    },
    [
        300
    ],
    {
        "_77": 301,
        "_79": 111,
        "_81": 302
    },
    "2067628123",
    "3CuBjEMi97tY3EGnq0NA9s",
    "3544641259",
    {
        "_23": 303,
        "_70": 34,
        "_72": 101,
        "_74": 305
    },
    [
        306,
        308
    ],
    {
        "_77": 307,
        "_79": 111,
        "_81": 101
    },
    "2856133350",
    {
        "_77": 309,
        "_79": 111,
        "_81": 101
    },
    "3214154973",
    "3645668434",
    {
        "_23": 310,
        "_70": 71,
        "_72": 312,
        "_74": 313
    },
    "1CWwhBKuOiRAC9V8HRBJRU",
    [
        314
    ],
    {
        "_77": 315,
        "_79": 80,
        "_81": 159
    },
    "3863445312",
    "3700195277",
    {
        "_23": 316,
        "_70": 34,
        "_72": 101,
        "_74": 318
    },
    [],
    {
        "_23": 211,
        "_70": 71,
        "_72": 212,
        "_74": 320
    },
    [],
    "3802510433",
    {
        "_23": 321,
        "_70": 71,
        "_72": 323,
        "_74": 324
    },
    "6FLEMI2GBFmVWGEsEGyASD:100.00:5",
    [],
    "3822950319",
    {
        "_23": 325,
        "_70": 71,
        "_72": 327,
        "_74": 328
    },
    "2CBvDiHjHIK9xlL4ItyXmK:100.00:1",
    [],
    "3838495619",
    {
        "_23": 329,
        "_70": 34,
        "_72": 101,
        "_74": 331
    },
    [],
    {
        "_23": 104,
        "_70": 71,
        "_72": 105,
        "_74": 333
    },
    [],
    "3940160259",
    {
        "_23": 334,
        "_70": 71,
        "_72": 336,
        "_74": 337
    },
    "2mmE1EmtOqtbWemO2wGuMO:100.00:4",
    [
        338,
        340,
        342
    ],
    {
        "_77": 339,
        "_79": 111,
        "_81": 101
    },
    "4180060165",
    {
        "_77": 341,
        "_79": 111,
        "_81": 101
    },
    "3765213438",
    {
        "_77": 343,
        "_79": 80,
        "_81": 344
    },
    "4078831437",
    "6bgwAROz7oF1OcKWxH4vHm:100.00:6",
    "3954884439",
    {
        "_23": 345,
        "_70": 71,
        "_72": 347,
        "_74": 348
    },
    "5rqjCf7T9KpJtLnaE73Kum:100.00:4",
    [],
    {
        "_23": 343,
        "_70": 71,
        "_72": 344,
        "_74": 350
    },
    [
        351,
        352
    ],
    {
        "_77": 339,
        "_79": 111,
        "_81": 101
    },
    {
        "_77": 341,
        "_79": 111,
        "_81": 101
    },
    "4207619515",
    {
        "_23": 353,
        "_70": 34,
        "_72": 101,
        "_74": 355
    },
    [],
    "4226692983",
    {
        "_23": 356,
        "_70": 71,
        "_72": 358,
        "_74": 359
    },
    "6sEu91zwlBGSKOqFiNpGlA:100.00:2",
    [],
    "dynamic_configs",
    {
        "_362": 363,
        "_375": 376,
        "_381": 382,
        "_385": 386,
        "_397": 398,
        "_403": 404,
        "_421": 422,
        "_426": 427,
        "_431": 432,
        "_436": 437,
        "_440": 441,
        "_446": 447
    },
    "357305500",
    {
        "_23": 362,
        "_70": 364,
        "_366": 367,
        "_72": 367,
        "_368": 34,
        "_74": 369,
        "_373": 34,
        "_374": 34
    },
    {
        "_365": 71
    },
    "can_see_upsell",
    "group",
    "launchedGroup",
    "is_device_based",
    [
        370
    ],
    {
        "_77": 371,
        "_79": 80,
        "_81": 372
    },
    "317829697",
    "598ORr5O5ZardhhzMhz8k0:100.00:11",
    "is_user_in_experiment",
    "is_experiment_active",
    "954359911",
    {
        "_23": 375,
        "_70": 377,
        "_366": 379,
        "_72": 379,
        "_368": 34,
        "_74": 380,
        "_373": 71,
        "_374": 71
    },
    {
        "_378": 34
    },
    "enabled",
    "5zN2l0bhNBO2gpivWHXwRY",
    [],
    "1001765573",
    {
        "_23": 381,
        "_70": 383,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 384
    },
    {},
    [],
    "1146308370",
    {
        "_23": 385,
        "_70": 387,
        "_366": 389,
        "_72": 389,
        "_368": 34,
        "_74": 390,
        "_373": 71,
        "_374": 71
    },
    {
        "_388": 71
    },
    "enable-copy-and-open",
    "4jQR01pTnwmjITqDD8PD2s",
    [
        391,
        394
    ],
    {
        "_77": 392,
        "_79": 80,
        "_81": 393
    },
    "303767167",
    "4kquVSCZpyFb5Sqki2BagX:100.00:6",
    {
        "_77": 395,
        "_79": 80,
        "_81": 396
    },
    "3284359640",
    "1E2e7sRUWkvJybjXfbiuoB",
    "1165680819",
    {
        "_23": 397,
        "_70": 399,
        "_366": 401,
        "_72": 401,
        "_368": 34,
        "_74": 402,
        "_373": 71,
        "_374": 71
    },
    {
        "_400": 71
    },
    "show_new_banner",
    "VVjatl8N5mxurs3Cje5TV",
    [],
    "1967546325",
    {
        "_23": 403,
        "_70": 405,
        "_366": 418,
        "_72": 418,
        "_368": 34,
        "_74": 419
    },
    {
        "_406": 71,
        "_407": 71,
        "_408": 34,
        "_409": 34,
        "_410": 71,
        "_411": 71,
        "_412": 413,
        "_414": 413,
        "_415": 416,
        "_417": 71
    },
    "gdrivePicker",
    "o365Picker",
    "gdriveLink",
    "o365Link",
    "o365PersonalLink",
    "o365BusinessLink",
    "gdrivePercentage",
    100,
    "o365Percentage",
    "loadTestPercentage",
    0,
    "showWorkspaceSettings",
    "2bcszlc7CFHdfdCdq7jXNb:100.00:5",
    [
        420
    ],
    {
        "_77": 307,
        "_79": 111,
        "_81": 101
    },
    "2043237793",
    {
        "_23": 421,
        "_70": 423,
        "_366": 367,
        "_72": 367,
        "_368": 34,
        "_74": 425,
        "_373": 34,
        "_374": 34
    },
    {
        "_424": 70
    },
    "bucket",
    [],
    "2513291161",
    {
        "_23": 426,
        "_70": 428,
        "_366": 429,
        "_72": 429,
        "_368": 34,
        "_74": 430,
        "_373": 71,
        "_374": 71
    },
    {
        "_378": 34
    },
    "2FTh6vlZcd8ha1OXcdnD3J",
    [],
    "3159301283",
    {
        "_23": 431,
        "_70": 433,
        "_366": 434,
        "_72": 434,
        "_368": 34,
        "_74": 435,
        "_373": 71,
        "_374": 71
    },
    {
        "_378": 71
    },
    "3Cce6z1hcoF2UBYwyupFck",
    [],
    "3217984440",
    {
        "_23": 436,
        "_70": 438,
        "_366": 367,
        "_72": 367,
        "_368": 34,
        "_74": 439,
        "_373": 34,
        "_374": 34
    },
    {
        "_378": 71
    },
    [],
    "3230069703",
    {
        "_23": 440,
        "_70": 442,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 445
    },
    {
        "_443": 444
    },
    "expirySeconds",
    15,
    [],
    "4198227845",
    {
        "_23": 446,
        "_70": 448,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 455
    },
    {
        "_449": 34,
        "_450": 34,
        "_451": 34,
        "_452": 34,
        "_453": 34,
        "_454": 34
    },
    "enabled_for_platform_override",
    "enabled_for_platform_new",
    "enabled_for_platform_existing",
    "enabled_for_chat_override",
    "enabled_for_chat_new",
    "enabled_for_chat_existing",
    [],
    "layer_configs",
    {
        "_458": 459,
        "_482": 483,
        "_487": 488,
        "_493": 494,
        "_505": 506,
        "_511": 512,
        "_516": 517,
        "_527": 528,
        "_561": 562,
        "_568": 569,
        "_581": 582,
        "_587": 588,
        "_593": 594,
        "_602": 603,
        "_617": 618,
        "_636": 637,
        "_653": 654,
        "_666": 667,
        "_675": 676,
        "_685": 686,
        "_696": 697,
        "_708": 709,
        "_715": 716,
        "_729": 730,
        "_737": 738,
        "_747": 748,
        "_756": 757,
        "_762": 763,
        "_768": 769,
        "_802": 803,
        "_818": 819,
        "_825": 826,
        "_832": 833
    },
    "16152997",
    {
        "_23": 458,
        "_70": 460,
        "_366": 471,
        "_72": 471,
        "_368": 34,
        "_74": 472,
        "_476": 477,
        "_478": 479,
        "_374": 34,
        "_373": 34,
        "_480": 481
    },
    {
        "_461": 71,
        "_462": 34,
        "_463": 71,
        "_464": 465,
        "_466": 465,
        "_467": 416,
        "_468": 34,
        "_469": 71,
        "_470": 34
    },
    "show_preview_when_collapsed",
    "expand_by_default",
    "is_enabled",
    "summarizer_system_prompt",
    "",
    "summarizer_chunk_template",
    "summarizer_chunk_char_limit",
    "enable_o3_mini_retrieval",
    "override_o3_mini_to_high",
    "enable_reason_by_default",
    "6DaNqHbUdaQZCJTtuXMn3l:override",
    [
        473
    ],
    {
        "_77": 474,
        "_79": 80,
        "_81": 475
    },
    "747145983",
    "1yBei0bniPE2f1TkI3MLWa",
    "explicit_parameters",
    [
        461,
        462,
        463
    ],
    "allocated_experiment_name",
    "1630255509",
    "undelegated_secondary_exposures",
    [
        473
    ],
    "40440673",
    {
        "_23": 482,
        "_70": 484,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 485,
        "_476": 486,
        "_480": 485
    },
    {},
    [],
    [],
    "51287004",
    {
        "_23": 487,
        "_70": 489,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 491,
        "_476": 492,
        "_480": 491
    },
    {
        "_490": 71
    },
    "enable",
    [],
    [],
    "183390215",
    {
        "_23": 493,
        "_70": 495,
        "_366": 498,
        "_72": 498,
        "_368": 71,
        "_74": 499,
        "_476": 502,
        "_478": 503,
        "_374": 71,
        "_373": 34,
        "_480": 504
    },
    {
        "_496": 34,
        "_497": 34
    },
    "signup_allow_phone",
    "in_phone_signup_holdout",
    "targetingGate",
    [
        500
    ],
    {
        "_77": 501,
        "_79": 111,
        "_81": 101
    },
    "3874938189",
    [
        496,
        497
    ],
    "4005636946",
    [],
    "190694971",
    {
        "_23": 505,
        "_70": 507,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 509,
        "_476": 510,
        "_480": 509
    },
    {
        "_508": 34
    },
    "show_nux",
    [],
    [],
    "229662723",
    {
        "_23": 511,
        "_70": 513,
        "_366": 101,
        "_72": 101,
        "_368": 71,
        "_74": 514,
        "_476": 515,
        "_480": 514
    },
    {},
    [],
    [],
    "387752763",
    {
        "_23": 516,
        "_70": 518,
        "_366": 101,
        "_72": 101,
        "_368": 71,
        "_74": 521,
        "_476": 526,
        "_480": 521
    },
    {
        "_519": 71,
        "_520": 71
    },
    "enable_slash_commands",
    "enable_rich_text_composer",
    [
        522,
        523,
        524
    ],
    {
        "_77": 107,
        "_79": 80,
        "_81": 108
    },
    {
        "_77": 110,
        "_79": 111,
        "_81": 112
    },
    {
        "_77": 525,
        "_79": 111,
        "_81": 101
    },
    "1410082514",
    [],
    "468168202",
    {
        "_23": 527,
        "_70": 529,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 557,
        "_476": 560,
        "_480": 557
    },
    {
        "_530": 71,
        "_531": 34,
        "_532": 71,
        "_533": 71,
        "_534": 34,
        "_535": 34,
        "_536": 34,
        "_537": 34,
        "_538": 34,
        "_539": 34,
        "_540": 34,
        "_541": 34,
        "_542": 34,
        "_543": 34,
        "_544": 71,
        "_545": 71,
        "_546": 34,
        "_547": 71,
        "_548": 71,
        "_549": 550,
        "_551": 552,
        "_553": 34,
        "_554": 555,
        "_556": 34
    },
    "is_team_enabled",
    "is_yearly_plus_subscription_enabled",
    "is_split_between_personal_and_business_enabled",
    "is_modal_fullscreen",
    "is_v2_toggle_labels_enabled",
    "is_bw",
    "is_produce_colors",
    "is_produce_color_scheme",
    "is_mobile_web_toggle_enabled",
    "is_enterprise_enabled",
    "is_produce_text",
    "is_optimized_checkout",
    "is_save_stripe_payment_info_enabled",
    "is_auto_save_stripe_payment_info_enabled",
    "does_manage_my_subscription_link_take_user_to_subscription_settings",
    "should_open_cancellation_survey_after_canceling",
    "should_cancel_button_take_user_to_stripe",
    "should_show_manage_my_subscription_link",
    "is_stripe_manage_subscription_link_enabled",
    "cancellation_modal_cancel_button_color",
    "danger",
    "cancellation_modal_go_back_button_color",
    "secondary",
    "should_show_cp",
    "cp_eligibility_months",
    3,
    "should_offer_paypal_when_eligible",
    [
        558
    ],
    {
        "_77": 559,
        "_79": 111,
        "_81": 101
    },
    "1847092144",
    [],
    "668322707",
    {
        "_23": 561,
        "_70": 563,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 566,
        "_476": 567,
        "_480": 566
    },
    {
        "_564": 71,
        "_565": 71
    },
    "show_citations_with_title",
    "use_chip_style_citations",
    [],
    [],
    "871635014",
    {
        "_23": 568,
        "_70": 570,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 579,
        "_476": 580,
        "_480": 579
    },
    {
        "_571": 34,
        "_572": 71,
        "_573": 34,
        "_574": 575,
        "_576": 101,
        "_577": 34,
        "_578": 34
    },
    "snowflake_composer_entry_point",
    "use_broad_rate_limit_language",
    "voice_holdout",
    "krisp_noise_filter",
    "none",
    "voice_entry_point_style",
    "show_label_on_button",
    "voice_only",
    [],
    [],
    "1170120107",
    {
        "_23": 581,
        "_70": 583,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 585,
        "_476": 586,
        "_480": 585
    },
    {
        "_584": 34
    },
    "is_whisper_enabled",
    [],
    [],
    "1238742812",
    {
        "_23": 587,
        "_70": 589,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 591,
        "_476": 592,
        "_480": 591
    },
    {
        "_590": 34
    },
    "should_enable_zh_tw",
    [],
    [],
    "1320801051",
    {
        "_23": 593,
        "_70": 595,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 600,
        "_476": 601,
        "_480": 600
    },
    {
        "_596": 34,
        "_597": 34,
        "_598": 71,
        "_599": 34
    },
    "hide_new_at_workspace_section",
    "hide_section_new_at_workspace",
    "gpt_discovery_experiment_enabled",
    "popular_at_my_workspace_enabled",
    [],
    [],
    "1346366956",
    {
        "_23": 602,
        "_70": 604,
        "_366": 101,
        "_72": 101,
        "_368": 71,
        "_74": 615,
        "_476": 616,
        "_480": 615
    },
    {
        "_605": 34,
        "_606": 607,
        "_608": 34,
        "_496": 34,
        "_609": 34,
        "_610": 34,
        "_611": 34,
        "_612": 34,
        "_613": 614
    },
    "use_email_otp",
    "signup_cta_copy",
    "SIGN_UP",
    "login_allow_phone",
    "forwardToAuthApi",
    "use_new_phone_ui",
    "in_signup_allow_phone_hold_out",
    "use_formatted_national_number",
    "continue_with_email_phone_placement",
    "after_sso",
    [],
    [],
    "1547743984",
    {
        "_23": 617,
        "_70": 619,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 632,
        "_476": 635,
        "_480": 632
    },
    {
        "_620": 34,
        "_621": 34,
        "_622": 34,
        "_623": 34,
        "_624": 34,
        "_625": 34,
        "_626": 34,
        "_627": 71,
        "_628": 34,
        "_629": 34,
        "_630": 71,
        "_631": 71
    },
    "should_simplify_modal",
    "is_simplified_sharing_modal_enabled",
    "is_social_share_options_enabled",
    "is_update_shared_links_enabled",
    "is_discoverability_toggle_enabled",
    "show_copylink_state_if_no_updates",
    "is_continue_enabled",
    "show_share_button_text",
    "is_meta_improvements_enabled",
    "show_share_button_inline",
    "use_dalle_preview",
    "in_dalle_preview_exp",
    [
        633
    ],
    {
        "_77": 634,
        "_79": 111,
        "_81": 101
    },
    "4038001028",
    [],
    "1630876919",
    {
        "_23": 636,
        "_70": 638,
        "_366": 645,
        "_72": 645,
        "_368": 34,
        "_74": 646,
        "_476": 650,
        "_478": 651,
        "_374": 71,
        "_373": 71,
        "_480": 652
    },
    {
        "_639": 71,
        "_640": 71,
        "_641": 71,
        "_642": 71,
        "_643": 34,
        "_644": 71
    },
    "enable_indexing",
    "backfill_completed",
    "enable_local_indexing",
    "enable_ux",
    "enable_us_rollout",
    "enable_ux_rollout",
    "31UyKaWB8PZhFswQt29NlZ",
    [
        647
    ],
    {
        "_77": 648,
        "_79": 80,
        "_81": 649
    },
    "2372319800",
    "4NZS9cdXgw2uEnVQCdyNMH:100.00:30",
    [
        639,
        641,
        640,
        642,
        644
    ],
    "1028722647",
    [],
    "1696863369",
    {
        "_23": 653,
        "_70": 655,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 658,
        "_476": 665,
        "_480": 658
    },
    {
        "_656": 34,
        "_657": 34
    },
    "has_sidekick_access",
    "show_nux_banner",
    [
        659,
        662
    ],
    {
        "_77": 660,
        "_79": 111,
        "_81": 661
    },
    "1938289220",
    "79O8DQPDmTKxnLdAH9loVk",
    {
        "_77": 663,
        "_79": 111,
        "_81": 664
    },
    "2033872549",
    "7dScmNU0bu2UQuzCNtva50",
    [],
    "1697140512",
    {
        "_23": 666,
        "_70": 668,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 670,
        "_476": 674,
        "_480": 670
    },
    {
        "_657": 34,
        "_669": 34
    },
    "can_download_sidetron",
    [
        671
    ],
    {
        "_77": 672,
        "_79": 111,
        "_81": 673
    },
    "2919213474",
    "6HLlb6nSjJk5ADynHucWgP",
    [],
    "1704793646",
    {
        "_23": 675,
        "_70": 677,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 681,
        "_476": 684,
        "_480": 681
    },
    {
        "_678": 34,
        "_679": 680
    },
    "greeting_web",
    "name_char_limit",
    20,
    [
        682
    ],
    {
        "_77": 683,
        "_79": 111,
        "_81": 101
    },
    "331938894",
    [],
    "1780960461",
    {
        "_23": 685,
        "_70": 687,
        "_366": 498,
        "_72": 498,
        "_368": 34,
        "_74": 690,
        "_476": 693,
        "_478": 694,
        "_374": 71,
        "_373": 34,
        "_480": 695
    },
    {
        "_688": 34,
        "_689": 34,
        "_678": 34
    },
    "mobile",
    "web",
    [
        691
    ],
    {
        "_77": 692,
        "_79": 111,
        "_81": 101
    },
    "3074373870",
    [
        688,
        689
    ],
    "2198260923",
    [],
    "1914829685",
    {
        "_23": 696,
        "_70": 698,
        "_366": 700,
        "_72": 700,
        "_368": 71,
        "_74": 701,
        "_476": 705,
        "_478": 706,
        "_374": 34,
        "_373": 34,
        "_480": 707
    },
    {
        "_699": 71
    },
    "forward_to_authapi",
    "2RO4BOrVWPrsxRUPYNKPLe:override",
    [
        702
    ],
    {
        "_77": 703,
        "_79": 80,
        "_81": 704
    },
    "14938527",
    "3QgLJ91lKIc7VAOjo5SDz7",
    [
        699
    ],
    "1856338298",
    [
        702
    ],
    "2152104812",
    {
        "_23": 708,
        "_70": 710,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 713,
        "_476": 714,
        "_480": 713
    },
    {
        "_711": 34,
        "_712": 34
    },
    "hide_gpts_if_none",
    "hide_default_gpts",
    [],
    [],
    "3048336830",
    {
        "_23": 715,
        "_70": 717,
        "_366": 720,
        "_72": 720,
        "_368": 34,
        "_74": 721,
        "_476": 728,
        "_480": 721
    },
    {
        "_718": 71,
        "_719": 34
    },
    "is-enabled",
    "use-rtl-layout",
    "localization-april Nzc6Xnht6tIVmb48Ejg1T:override",
    [
        722,
        725
    ],
    {
        "_77": 723,
        "_79": 111,
        "_81": 724
    },
    "3922145230",
    "14DZA2LumaPqAdCo52CrUB",
    {
        "_77": 726,
        "_79": 80,
        "_81": 727
    },
    "3700615661",
    "66covjaoZoe9pQR4I68jOB",
    [],
    "3178812292",
    {
        "_23": 729,
        "_70": 731,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 733,
        "_476": 736,
        "_480": 733
    },
    {
        "_732": 34
    },
    "use_f_convo",
    [
        734
    ],
    {
        "_77": 735,
        "_79": 111,
        "_81": 101
    },
    "3799260860",
    [],
    "3436367576",
    {
        "_23": 737,
        "_70": 739,
        "_366": 498,
        "_72": 498,
        "_368": 34,
        "_74": 741,
        "_476": 744,
        "_478": 745,
        "_374": 71,
        "_373": 34,
        "_480": 746
    },
    {
        "_639": 34,
        "_740": 416,
        "_642": 34,
        "_641": 34,
        "_640": 34
    },
    "wave",
    [
        742
    ],
    {
        "_77": 743,
        "_79": 111,
        "_81": 101
    },
    "1221279314",
    [
        639,
        740,
        640,
        642,
        641
    ],
    "938456440",
    [],
    "3471271313",
    {
        "_23": 747,
        "_70": 749,
        "_366": 751,
        "_72": 751,
        "_368": 71,
        "_74": 752,
        "_476": 753,
        "_478": 754,
        "_374": 34,
        "_373": 34,
        "_480": 755
    },
    {
        "_750": 34
    },
    "show_upsell",
    "prestart",
    [],
    [
        750
    ],
    "3021307436",
    [],
    "3517133692",
    {
        "_23": 756,
        "_70": 758,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 760,
        "_476": 761,
        "_480": 760
    },
    {
        "_759": 34
    },
    "is_memory_undo_enabled",
    [],
    [],
    "3590606857",
    {
        "_23": 762,
        "_70": 764,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 766,
        "_476": 767,
        "_480": 766
    },
    {
        "_765": 34
    },
    "should_offer_paypal",
    [],
    [],
    "3637408529",
    {
        "_23": 768,
        "_70": 770,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 796,
        "_476": 801,
        "_480": 796
    },
    {
        "_771": 71,
        "_772": 34,
        "_773": 34,
        "_774": 34,
        "_775": 776,
        "_777": 778,
        "_779": 71,
        "_780": 71,
        "_781": 71,
        "_782": 34,
        "_783": 71,
        "_784": 34,
        "_785": 34,
        "_786": 71,
        "_787": 34,
        "_788": 71,
        "_789": 555,
        "_790": 791,
        "_792": 71,
        "_793": 794,
        "_795": 34
    },
    "is_anon_chat_enabled",
    "is_anon_chat_enabled_for_new_users_only",
    "is_try_it_first_on_login_page_enabled",
    "is_no_auth_welcome_modal_enabled",
    "no_auth_soft_rate_limit",
    5,
    "no_auth_hard_rate_limit",
    1200,
    "should_show_no_auth_signup_banner",
    "is_no_auth_welcome_back_modal_enabled",
    "is_no_auth_soft_rate_limit_modal_enabled",
    "is_no_auth_gpt4o_modal_enabled",
    "is_login_primary_button",
    "is_desktop_primary_auth_button_on_right",
    "is_primary_btn_blue",
    "should_show_disclaimer_only_once_per_device",
    "is_secondary_banner_button_enabled",
    "is_secondary_auth_banner_button_enabled",
    "no_auth_banner_signup_rate_limit",
    "composer_text",
    "ASK_ANYTHING",
    "is_in_composer_text_exp",
    "no_auth_upsell_wording",
    "NO_CHANGE",
    "should_refresh_access_token_error_take_user_to_no_auth",
    [
        797,
        799
    ],
    {
        "_77": 798,
        "_79": 111,
        "_81": 159
    },
    "3238165271",
    {
        "_77": 800,
        "_79": 111,
        "_81": 159
    },
    "2983591614",
    [],
    "3647926857",
    {
        "_23": 802,
        "_70": 804,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 812,
        "_476": 817,
        "_480": 812
    },
    {
        "_805": 71,
        "_806": 34,
        "_807": 808,
        "_809": 34,
        "_810": 34,
        "_811": 575
    },
    "unified_architecture",
    "ux_updates",
    "inference_debounce_ms",
    400,
    "autoswitcher_enabled",
    "copy-and-link",
    "reasoning_slider",
    [
        813,
        815
    ],
    {
        "_77": 814,
        "_79": 111,
        "_81": 101
    },
    "850280859",
    {
        "_77": 816,
        "_79": 111,
        "_81": 101
    },
    "13512905",
    [],
    "3711177917",
    {
        "_23": 818,
        "_70": 820,
        "_366": 101,
        "_72": 101,
        "_368": 71,
        "_74": 823,
        "_476": 824,
        "_480": 823
    },
    {
        "_821": 34,
        "_822": 71
    },
    "is_summarizer_default_expanded",
    "is_inline_summarizer_enabled",
    [],
    [],
    "3972089454",
    {
        "_23": 825,
        "_70": 827,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 830,
        "_476": 831,
        "_480": 830
    },
    {
        "_828": 829
    },
    "search_scoring_dyconfig_name",
    "gizmo_search_score_config",
    [],
    [],
    "4211831761",
    {
        "_23": 832,
        "_70": 834,
        "_366": 101,
        "_72": 101,
        "_368": 71,
        "_74": 835,
        "_476": 836,
        "_480": 835
    },
    {
        "_378": 34
    },
    [],
    [],
    "sdkParams",
    {},
    "has_updates",
    "generator",
    "statsig-node-sdk",
    "sdkInfo",
    {
        "_844": 845,
        "_846": 847
    },
    "sdkType",
    "statsig-node",
    "sdkVersion",
    "5.26.0",
    "time",
    1742979946815,
    "evaluated_keys",
    {
        "_852": 22,
        "_853": 854
    },
    "userID",
    "user-chatgpt",
    {
        "_855": 856,
        "_857": 856,
        "_858": 856,
        "_859": 43,
        "_860": 43
    },
    "WebAnonymousCookieID",
    "f10eee0a-3567-4980-ad81-98da119a5c6b",
    "DeviceId",
    "stableID",
    "workspace_id",
    "account_id",
    "hash_used",
    "djb2",
    {
        "_852": 22,
        "_864": 865,
        "_866": 867,
        "_60": 61,
        "_853": 854,
        "_871": 872
    },
    "country",
    "SG",
    "custom",
    {
        "_868": 45,
        "_859": 43,
        "_860": 43,
        "_869": 71,
        "_870": 16
    },
    "plan_type",
    "is_paid",
    "auth_status",
    "statsigEnvironment",
    {
        "_873": 874
    },
    "tier",
    "production",
    "experimental",
    {
        "_66": 877,
        "_360": 1200,
        "_456": 1303,
        "_837": 1828,
        "_839": 71,
        "_840": 841,
        "_842": 1829,
        "_848": 849,
        "_850": 1830,
        "_861": 862,
        "_19": 1832
    },
    {
        "_703": 878,
        "_880": 881,
        "_883": 884,
        "_887": 888,
        "_893": 894,
        "_897": 898,
        "_901": 902,
        "_905": 906,
        "_908": 909,
        "_912": 913,
        "_915": 916,
        "_920": 921,
        "_924": 925,
        "_928": 929,
        "_937": 938,
        "_941": 942,
        "_945": 946,
        "_949": 950,
        "_953": 954,
        "_957": 958,
        "_142": 964,
        "_966": 967,
        "_969": 970,
        "_972": 973,
        "_975": 976,
        "_978": 979,
        "_983": 984,
        "_987": 988,
        "_991": 992,
        "_995": 996,
        "_998": 999,
        "_1002": 1003,
        "_1006": 1007,
        "_1010": 1011,
        "_1013": 1014,
        "_1017": 1018,
        "_1020": 1021,
        "_1024": 1025,
        "_1028": 1029,
        "_1031": 1032,
        "_1034": 1035,
        "_1037": 1038,
        "_1041": 1042,
        "_301": 1044,
        "_1046": 1047,
        "_1049": 1050,
        "_1053": 1054,
        "_1057": 1058,
        "_1063": 1064,
        "_1067": 1068,
        "_1071": 1072,
        "_1075": 1076,
        "_1078": 1079,
        "_1082": 1083,
        "_1086": 1087,
        "_1090": 1091,
        "_1096": 1097,
        "_1100": 1101,
        "_1106": 1107,
        "_1111": 1112,
        "_1115": 1116,
        "_1118": 1119,
        "_1121": 1122,
        "_1124": 1125,
        "_1127": 1128,
        "_1130": 1131,
        "_1134": 1135,
        "_1137": 1138,
        "_1141": 1142,
        "_303": 1144,
        "_1148": 1149,
        "_1152": 1153,
        "_1156": 1157,
        "_1159": 1160,
        "_1162": 1163,
        "_1166": 1167,
        "_723": 1170,
        "_1172": 1173,
        "_1176": 1177,
        "_1179": 1180,
        "_1182": 1183,
        "_1186": 1187,
        "_1190": 1191,
        "_1193": 1194,
        "_1196": 1197
    },
    {
        "_23": 703,
        "_70": 71,
        "_72": 704,
        "_74": 879
    },
    [],
    "156153730",
    {
        "_23": 880,
        "_70": 34,
        "_72": 101,
        "_74": 882
    },
    [],
    "222560275",
    {
        "_23": 883,
        "_70": 34,
        "_72": 885,
        "_74": 886
    },
    "5pv2QpbgXNDB0QnBo3LTti:10.00:1",
    [],
    "223382091",
    {
        "_23": 887,
        "_70": 34,
        "_72": 889,
        "_74": 890
    },
    "1fKkxDiVebEKfTj8nDAjHe",
    [
        891,
        892
    ],
    {
        "_77": 339,
        "_79": 111,
        "_81": 101
    },
    {
        "_77": 341,
        "_79": 111,
        "_81": 101
    },
    "402391964",
    {
        "_23": 893,
        "_70": 34,
        "_72": 895,
        "_74": 896
    },
    "14sAQaGJDosUKVV0DFZsAL",
    [],
    "471233253",
    {
        "_23": 897,
        "_70": 34,
        "_72": 899,
        "_74": 900
    },
    "3Yf9H7TxMC122pchwAkoLB",
    [],
    "550432558",
    {
        "_23": 901,
        "_70": 71,
        "_72": 903,
        "_74": 904
    },
    "4XbSwfoqBmVtxwz32sweLb",
    [],
    "573184874",
    {
        "_23": 905,
        "_70": 34,
        "_72": 101,
        "_74": 907
    },
    [],
    "582612297",
    {
        "_23": 908,
        "_70": 71,
        "_72": 910,
        "_74": 911
    },
    "5censDsCfS2zQeYtTIui2s:100.00:2",
    [],
    "614413305",
    {
        "_23": 912,
        "_70": 34,
        "_72": 101,
        "_74": 914
    },
    [],
    "653593316",
    {
        "_23": 915,
        "_70": 71,
        "_72": 917,
        "_74": 918
    },
    "GJ8pvorFDIe3Z4WonIr2s",
    [
        919
    ],
    {
        "_77": 321,
        "_79": 80,
        "_81": 323
    },
    "706943082",
    {
        "_23": 920,
        "_70": 71,
        "_72": 922,
        "_74": 923
    },
    "X9mJLzEwXwKo3p0LSSQIL",
    [],
    "727502549",
    {
        "_23": 924,
        "_70": 71,
        "_72": 926,
        "_74": 927
    },
    "6EYbmM9CyqCRO6U6k3dROA",
    [],
    "756982148",
    {
        "_23": 928,
        "_70": 71,
        "_72": 930,
        "_74": 931
    },
    "3oAWYdzegKPwxhFJjJrGz3",
    [
        932,
        934
    ],
    {
        "_77": 933,
        "_79": 111,
        "_81": 101
    },
    "1456438623",
    {
        "_77": 935,
        "_79": 80,
        "_81": 936
    },
    "3805873235",
    "5KvGWxgOdialy0Dx9IrqmW:100.00:23",
    "756982149",
    {
        "_23": 937,
        "_70": 34,
        "_72": 939,
        "_74": 940
    },
    "1rXg44we6gmcRqYsiZzfL4:0.00:1",
    [],
    "795789557",
    {
        "_23": 941,
        "_70": 34,
        "_72": 943,
        "_74": 944
    },
    "2GzNaY2UIV2RYDjl4grJNG:0.00:1",
    [],
    "809056127",
    {
        "_23": 945,
        "_70": 71,
        "_72": 947,
        "_74": 948
    },
    "54ufwSF4KjxPi2AIrjbelh",
    [],
    "810701024",
    {
        "_23": 949,
        "_70": 71,
        "_72": 951,
        "_74": 952
    },
    "6U8ODe5JvFov5zs1rOzJjD",
    [],
    "891514942",
    {
        "_23": 953,
        "_70": 34,
        "_72": 955,
        "_74": 956
    },
    "aWUpylPDtFgWWhTxEsfCx",
    [],
    "989226566",
    {
        "_23": 957,
        "_70": 71,
        "_72": 959,
        "_74": 960
    },
    "6yqqYAWKtmfU8A7QGdiky4",
    [
        961,
        962
    ],
    {
        "_77": 137,
        "_79": 111,
        "_81": 138
    },
    {
        "_77": 140,
        "_79": 80,
        "_81": 963
    },
    "7D8EAif25E3Y8A3zkg6ljp:100.00:2",
    {
        "_23": 142,
        "_70": 71,
        "_72": 144,
        "_74": 965
    },
    [],
    "1032814809",
    {
        "_23": 966,
        "_70": 34,
        "_72": 101,
        "_74": 968
    },
    [],
    "1064007944",
    {
        "_23": 969,
        "_70": 34,
        "_72": 101,
        "_74": 971
    },
    [],
    "1099124727",
    {
        "_23": 972,
        "_70": 34,
        "_72": 101,
        "_74": 974
    },
    [],
    "1154002920",
    {
        "_23": 975,
        "_70": 34,
        "_72": 101,
        "_74": 977
    },
    [],
    "1166240779",
    {
        "_23": 978,
        "_70": 71,
        "_72": 980,
        "_74": 981
    },
    "4UjTXwt2XK975PANdi1Ma6:25.00:5",
    [
        982
    ],
    {
        "_77": 309,
        "_79": 111,
        "_81": 101
    },
    "1214379119",
    {
        "_23": 983,
        "_70": 34,
        "_72": 985,
        "_74": 986
    },
    "3Da3vJtBawdpcHFOEpjzZA:10.00:2",
    [],
    "1382475798",
    {
        "_23": 987,
        "_70": 71,
        "_72": 989,
        "_74": 990
    },
    "3P8OsGy1e5tQlR5dsTIWbL",
    [],
    "1416952492",
    {
        "_23": 991,
        "_70": 34,
        "_72": 993,
        "_74": 994
    },
    "2LD82enCtskHL9Vi2hS6Jq",
    [],
    "1422501431",
    {
        "_23": 995,
        "_70": 34,
        "_72": 101,
        "_74": 997
    },
    [],
    "1439437954",
    {
        "_23": 998,
        "_70": 34,
        "_72": 1000,
        "_74": 1001
    },
    "11IqDt7xc4mMNiyiSIMy1F:0.00:1",
    [],
    "1456513860",
    {
        "_23": 1002,
        "_70": 71,
        "_72": 1004,
        "_74": 1005
    },
    "jHXkU7q9axp0dXBSyzihH",
    [],
    "1468311859",
    {
        "_23": 1006,
        "_70": 34,
        "_72": 1008,
        "_74": 1009
    },
    "7tfl8ZUhwr5pzErE3ikBej",
    [],
    "1542198993",
    {
        "_23": 1010,
        "_70": 34,
        "_72": 101,
        "_74": 1012
    },
    [],
    "1656345175",
    {
        "_23": 1013,
        "_70": 71,
        "_72": 1015,
        "_74": 1016
    },
    "2CwIChuIr7SLQ2CyqRegF2",
    [],
    "1741586789",
    {
        "_23": 1017,
        "_70": 34,
        "_72": 101,
        "_74": 1019
    },
    [],
    "1760640904",
    {
        "_23": 1020,
        "_70": 71,
        "_72": 1022,
        "_74": 1023
    },
    "6ezOfLAw7fGQPVjfNsReIy",
    [],
    "1830177352",
    {
        "_23": 1024,
        "_70": 71,
        "_72": 1026,
        "_74": 1027
    },
    "44udGr8tXtB3ZIDHLV3HSF",
    [],
    "1839283687",
    {
        "_23": 1028,
        "_70": 34,
        "_72": 101,
        "_74": 1030
    },
    [],
    "1860647109",
    {
        "_23": 1031,
        "_70": 34,
        "_72": 101,
        "_74": 1033
    },
    [],
    "2000076788",
    {
        "_23": 1034,
        "_70": 34,
        "_72": 101,
        "_74": 1036
    },
    [],
    "2053937752",
    {
        "_23": 1037,
        "_70": 34,
        "_72": 1039,
        "_74": 1040
    },
    "2PLQzvwrGPxACRwaEcKbIh",
    [],
    "2056761365",
    {
        "_23": 1041,
        "_70": 34,
        "_72": 101,
        "_74": 1043
    },
    [],
    {
        "_23": 301,
        "_70": 34,
        "_72": 302,
        "_74": 1045
    },
    [],
    "2151954125",
    {
        "_23": 1046,
        "_70": 34,
        "_72": 101,
        "_74": 1048
    },
    [],
    "2153043779",
    {
        "_23": 1049,
        "_70": 71,
        "_72": 1051,
        "_74": 1052
    },
    "DamiTYVoTv9Z9jRFOT5iC",
    [],
    "2173548801",
    {
        "_23": 1053,
        "_70": 71,
        "_72": 1055,
        "_74": 1056
    },
    "22nVhoL17eyMvGWgFrDfZe",
    [],
    "2192543539",
    {
        "_23": 1057,
        "_70": 71,
        "_72": 1059,
        "_74": 1060
    },
    "4Ro1m2dj4fUBe4hcP1YKjj:75.00:3",
    [
        1061
    ],
    {
        "_77": 1062,
        "_79": 111,
        "_81": 101
    },
    "4206244917",
    "2232580636",
    {
        "_23": 1063,
        "_70": 71,
        "_72": 1065,
        "_74": 1066
    },
    "4y4Nd0nF0CFawcrQBbm7Mq:100.00:4",
    [],
    "2281969373",
    {
        "_23": 1067,
        "_70": 71,
        "_72": 1069,
        "_74": 1070
    },
    "6EbVeXErTdGtbchxdqEMTg",
    [],
    "2290870843",
    {
        "_23": 1071,
        "_70": 71,
        "_72": 1073,
        "_74": 1074
    },
    "5dONtElzUeyTTp5FvpWy6",
    [],
    "2360528850",
    {
        "_23": 1075,
        "_70": 34,
        "_72": 101,
        "_74": 1077
    },
    [],
    "2379988365",
    {
        "_23": 1078,
        "_70": 34,
        "_72": 101,
        "_74": 1080
    },
    [
        1081
    ],
    {
        "_77": 307,
        "_79": 111,
        "_81": 101
    },
    "2411734826",
    {
        "_23": 1082,
        "_70": 34,
        "_72": 1084,
        "_74": 1085
    },
    "33U1igAQgegRumGc4LbaB:2.00:1",
    [],
    "2445152477",
    {
        "_23": 1086,
        "_70": 71,
        "_72": 1088,
        "_74": 1089
    },
    "5qtlunRMswJX2JGoF8GikC",
    [],
    "2634628831",
    {
        "_23": 1090,
        "_70": 71,
        "_72": 1092,
        "_74": 1093
    },
    "6LfSag7ByiH0gGcqoFHHBe",
    [
        1094,
        1095
    ],
    {
        "_77": 933,
        "_79": 111,
        "_81": 101
    },
    {
        "_77": 935,
        "_79": 80,
        "_81": 936
    },
    "2637918557",
    {
        "_23": 1096,
        "_70": 71,
        "_72": 1098,
        "_74": 1099
    },
    "2XNTwszL419o7DMxzSa0vz:100.00:1",
    [],
    "2712556596",
    {
        "_23": 1100,
        "_70": 71,
        "_72": 1102,
        "_74": 1103
    },
    "7pPLEbQc7hKT1m7CbondoE",
    [
        1104
    ],
    {
        "_77": 1105,
        "_79": 111,
        "_81": 101
    },
    "135448051",
    "2781425969",
    {
        "_23": 1106,
        "_70": 34,
        "_72": 1108,
        "_74": 1109
    },
    "7BIMlzITwH6mysXL5ILPSw",
    [
        1110
    ],
    {
        "_77": 1105,
        "_79": 111,
        "_81": 101
    },
    "2833534668",
    {
        "_23": 1111,
        "_70": 71,
        "_72": 1113,
        "_74": 1114
    },
    "7uYkibMYlCPSnoWmmYNanm",
    [],
    "2935021756",
    {
        "_23": 1115,
        "_70": 34,
        "_72": 101,
        "_74": 1117
    },
    [],
    "2968810397",
    {
        "_23": 1118,
        "_70": 34,
        "_72": 101,
        "_74": 1120
    },
    [],
    "3058498100",
    {
        "_23": 1121,
        "_70": 34,
        "_72": 101,
        "_74": 1123
    },
    [],
    "3148583717",
    {
        "_23": 1124,
        "_70": 34,
        "_72": 101,
        "_74": 1126
    },
    [],
    "3241763787",
    {
        "_23": 1127,
        "_70": 34,
        "_72": 101,
        "_74": 1129
    },
    [],
    "3257646228",
    {
        "_23": 1130,
        "_70": 34,
        "_72": 1132,
        "_74": 1133
    },
    "3veZ6qhG4zTVvcrwpXXPgi:1.00:4",
    [],
    "3291247717",
    {
        "_23": 1134,
        "_70": 34,
        "_72": 101,
        "_74": 1136
    },
    [],
    "3435450078",
    {
        "_23": 1137,
        "_70": 71,
        "_72": 1139,
        "_74": 1140
    },
    "2qCdHpFuWOOkibzLRL0zgn",
    [],
    "3472722167",
    {
        "_23": 1141,
        "_70": 34,
        "_72": 101,
        "_74": 1143
    },
    [],
    {
        "_23": 303,
        "_70": 34,
        "_72": 101,
        "_74": 1145
    },
    [
        1146,
        1147
    ],
    {
        "_77": 307,
        "_79": 111,
        "_81": 101
    },
    {
        "_77": 309,
        "_79": 111,
        "_81": 101
    },
    "3612584454",
    {
        "_23": 1148,
        "_70": 34,
        "_72": 1150,
        "_74": 1151
    },
    "4fXx7LNuNnDASdmkzwNxtf",
    [],
    "3664702598",
    {
        "_23": 1152,
        "_70": 34,
        "_72": 1154,
        "_74": 1155
    },
    "7x9wS41bRDCji9ns8x5Oej",
    [],
    "3678527908",
    {
        "_23": 1156,
        "_70": 34,
        "_72": 101,
        "_74": 1158
    },
    [],
    "3728856343",
    {
        "_23": 1159,
        "_70": 34,
        "_72": 101,
        "_74": 1161
    },
    [],
    "3861593998",
    {
        "_23": 1162,
        "_70": 34,
        "_72": 1164,
        "_74": 1165
    },
    "5DN2QZNg9iYP45NqvRetnu",
    [],
    "3910241726",
    {
        "_23": 1166,
        "_70": 71,
        "_72": 1168,
        "_74": 1169
    },
    "1ItyvFbGou4epQp9HviAsm",
    [],
    {
        "_23": 723,
        "_70": 34,
        "_72": 724,
        "_74": 1171
    },
    [],
    "3940529303",
    {
        "_23": 1172,
        "_70": 71,
        "_72": 1174,
        "_74": 1175
    },
    "17mkpeWbaWfCeMrpE67FOc",
    [],
    "4012051055",
    {
        "_23": 1176,
        "_70": 34,
        "_72": 101,
        "_74": 1178
    },
    [],
    "4043415092",
    {
        "_23": 1179,
        "_70": 34,
        "_72": 101,
        "_74": 1181
    },
    [],
    "4132051975",
    {
        "_23": 1182,
        "_70": 71,
        "_72": 1184,
        "_74": 1185
    },
    "wLBwoUCuuMdnRwa9KkfHI",
    [],
    "4141006638",
    {
        "_23": 1186,
        "_70": 34,
        "_72": 1188,
        "_74": 1189
    },
    "6v4Q2eufBTFCb2P3fGZwPo",
    [],
    "4192239497",
    {
        "_23": 1190,
        "_70": 34,
        "_72": 101,
        "_74": 1192
    },
    [],
    "4206189746",
    {
        "_23": 1193,
        "_70": 34,
        "_72": 101,
        "_74": 1195
    },
    [],
    "4242210007",
    {
        "_23": 1196,
        "_70": 34,
        "_72": 1198,
        "_74": 1199
    },
    "5T7B6Qu0S7TF24HzOjoxJl",
    [],
    {
        "_1201": 1202,
        "_1209": 1210,
        "_375": 1214,
        "_1216": 1217,
        "_1235": 1236,
        "_1239": 1240,
        "_1245": 1246,
        "_1253": 1254,
        "_431": 1287,
        "_440": 1289,
        "_1291": 1292,
        "_1297": 1298
    },
    "550560761",
    {
        "_23": 1201,
        "_70": 1203,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1208
    },
    {
        "_1204": 1205,
        "_1206": 1207
    },
    "history_results_limit",
    6,
    "local_results_limit",
    2,
    [],
    "948081399",
    {
        "_23": 1209,
        "_70": 1211,
        "_366": 1212,
        "_72": 1212,
        "_368": 34,
        "_74": 1213,
        "_373": 34,
        "_374": 71
    },
    {},
    "layerAssignment",
    [],
    {
        "_23": 375,
        "_70": 377,
        "_366": 379,
        "_72": 379,
        "_368": 34,
        "_74": 1215,
        "_373": 71,
        "_374": 71
    },
    [],
    "1682643554",
    {
        "_23": 1216,
        "_70": 1218,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1234
    },
    {
        "_1219": 1220
    },
    "school_configurations",
    {
        "_1221": 1222,
        "_1230": 1231
    },
    "openai_1signup_for_1",
    {
        "_1223": 1224,
        "_1225": 1226,
        "_1227": 1228
    },
    "display_name",
    "OpenAI",
    "promotion_campaign_id",
    "students-2025-one-month-free",
    "domains",
    [
        1229
    ],
    "openai.com, mail.openai.com",
    "australia",
    {
        "_1223": 1224,
        "_1225": 1226,
        "_1227": 1232
    },
    [
        1233
    ],
    "edu.au",
    [],
    "1809520125",
    {
        "_23": 1235,
        "_70": 1237,
        "_366": 751,
        "_72": 751,
        "_368": 71,
        "_74": 1238,
        "_373": 34,
        "_374": 34
    },
    {},
    [],
    "2181185232",
    {
        "_23": 1239,
        "_70": 1241,
        "_366": 498,
        "_72": 498,
        "_368": 71,
        "_74": 1242,
        "_373": 34,
        "_374": 71
    },
    {},
    [
        1243
    ],
    {
        "_77": 1244,
        "_79": 111,
        "_81": 101
    },
    "1887864177",
    "2604379743",
    {
        "_23": 1245,
        "_70": 1247,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1252
    },
    {
        "_1248": 1249,
        "_1250": 1251
    },
    "nux_video_url",
    "https://persistent.oaistatic.com/image-gen/nux.CB3699EE.mov",
    "nux_image_url",
    "https://persistent.oaistatic.com/image-gen/nux.CB3699EE.jpg",
    [],
    "2821602598",
    {
        "_23": 1253,
        "_70": 1255,
        "_366": 101,
        "_72": 101,
        "_368": 71,
        "_74": 1286
    },
    {
        "_1256": 1257
    },
    "Football",
    [
        1258
    ],
    {
        "_1259": 1256,
        "_1260": 1261
    },
    "title",
    "templates",
    [
        1262,
        1272,
        1279
    ],
    {
        "_1263": 1264,
        "_1265": 1266
    },
    "text",
    "The [input] are down [input] with [input] left in the [input] quarter. What are their odds they [input]?",
    "suggestions",
    [
        1267,
        1268,
        1269,
        1270,
        1271
    ],
    "KC Chiefs",
    "27 - 10",
    "5 minutes",
    "4th",
    "win",
    {
        "_1263": 1273,
        "_1265": 1274
    },
    "I'm [input], played [input] and work out [input]. Help me train like a [input].",
    [
        1275,
        1276,
        1277,
        1278
    ],
    "29",
    "football",
    "3 days a week",
    "NFL running back",
    {
        "_1263": 1280,
        "_1265": 1281
    },
    "Write me a [input]. Include [input], [input], and [input].",
    [
        1282,
        1283,
        1284,
        1285
    ],
    "perfect halftime show song",
    "dancing",
    "fireworks",
    "being super fierce",
    [],
    {
        "_23": 431,
        "_70": 433,
        "_366": 434,
        "_72": 434,
        "_368": 34,
        "_74": 1288,
        "_373": 71,
        "_374": 71
    },
    [],
    {
        "_23": 440,
        "_70": 442,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1290
    },
    [],
    "3519108196",
    {
        "_23": 1291,
        "_70": 1293,
        "_366": 367,
        "_72": 367,
        "_368": 34,
        "_74": 1296,
        "_373": 34,
        "_374": 34
    },
    {
        "_1294": 71,
        "_1295": 71
    },
    "show-album-upload",
    "show-camera-upload",
    [],
    "3983984123",
    {
        "_23": 1297,
        "_70": 1299,
        "_366": 751,
        "_72": 751,
        "_368": 34,
        "_74": 1300,
        "_476": 1301,
        "_373": 34,
        "_374": 34,
        "_1302": 71
    },
    {
        "_759": 34
    },
    [],
    [
        759
    ],
    "is_in_layer",
    {
        "_1304": 1305,
        "_458": 1344,
        "_482": 1349,
        "_487": 1352,
        "_493": 1355,
        "_505": 1359,
        "_511": 1362,
        "_516": 1365,
        "_1374": 1375,
        "_1382": 1383,
        "_527": 1391,
        "_1395": 1396,
        "_1403": 1404,
        "_1412": 1413,
        "_1435": 1436,
        "_1445": 1446,
        "_568": 1465,
        "_581": 1468,
        "_587": 1471,
        "_593": 1474,
        "_602": 1477,
        "_1480": 1481,
        "_1489": 1490,
        "_1495": 1496,
        "_617": 1502,
        "_1506": 1507,
        "_636": 1512,
        "_653": 1517,
        "_675": 1522,
        "_685": 1526,
        "_1530": 1531,
        "_696": 1536,
        "_1540": 1541,
        "_1549": 1550,
        "_708": 1558,
        "_1561": 1562,
        "_1576": 1577,
        "_1582": 1583,
        "_1588": 1589,
        "_1597": 1598,
        "_1613": 1614,
        "_1620": 1621,
        "_1627": 1628,
        "_1633": 1634,
        "_1638": 1639,
        "_715": 1645,
        "_1652": 1653,
        "_1659": 1660,
        "_1665": 1666,
        "_737": 1683,
        "_747": 1691,
        "_756": 1694,
        "_1697": 1698,
        "_762": 1724,
        "_1727": 1728,
        "_1735": 1736,
        "_768": 1741,
        "_818": 1746,
        "_1749": 1750,
        "_1756": 1757,
        "_1771": 1772,
        "_825": 1776,
        "_1779": 1780,
        "_1784": 1785,
        "_832": 1814,
        "_1817": 1818
    },
    "109457",
    {
        "_23": 1304,
        "_70": 1306,
        "_366": 1334,
        "_72": 1334,
        "_368": 34,
        "_74": 1335,
        "_476": 1341,
        "_478": 1342,
        "_374": 71,
        "_373": 71,
        "_480": 1343
    },
    {
        "_1307": 34,
        "_1308": 34,
        "_1309": 34,
        "_1310": 34,
        "_1311": 34,
        "_1312": 465,
        "_1313": 34,
        "_1314": 34,
        "_1315": 34,
        "_1316": 465,
        "_1317": 34,
        "_1318": 1319,
        "_1320": 34,
        "_1321": 34,
        "_1322": 34,
        "_1323": 34,
        "_1324": 34,
        "_1325": 465,
        "_1326": 34,
        "_1327": 1328,
        "_1329": 1330,
        "_1331": 1330,
        "_1332": 34,
        "_1333": 34
    },
    "is_starter_prompt_popular",
    "is_starter_prompt_top_performer",
    "is_starter_prompt_back_and_forth",
    "use_starter_prompt_help_how_to",
    "model_talks_first",
    "model_talks_first_kind",
    "model_talks_first_augment_system_prompt",
    "is_starter_prompt_enabled_for_new_users_only",
    "add_system_prompt_during_onboarding",
    "onboarding_system_prompt_type",
    "enable_new_onboarding_flow",
    "new_onboarding_flow_qualified_start_date",
    "2025-02-28T06:00:00Z",
    "personalized_onboarding",
    "onboarding_show_custom_instructions_page",
    "write_custom_instructions_in_onboarding",
    "keep_onboarding_after_dismiss",
    "onboarding_dynamic_steps_based_on_main_usage",
    "onboarding_style",
    "onboarding_show_followups",
    "onboarding_inject_cards_position",
    9999,
    "ONBOARDING_EXAMPLES_PROMPT_ID",
    "convo_gen_examples_v2",
    "onboarding_gen_examples_prompt_type",
    "show_new_chat_nux",
    "is_guided_onboarding",
    "M3EE4Hyw83Rv7RjIICK6o",
    [
        1336,
        1338
    ],
    {
        "_77": 1337,
        "_79": 111,
        "_81": 101
    },
    "674041001",
    {
        "_77": 1339,
        "_79": 80,
        "_81": 1340
    },
    "59687878",
    "4k3eNmHeryixdsgalKqv0",
    [
        1325,
        1318,
        1315,
        1316,
        1320,
        1326,
        1327,
        1331,
        1332
    ],
    "52701554",
    [
        1336
    ],
    {
        "_23": 458,
        "_70": 460,
        "_366": 471,
        "_72": 471,
        "_368": 34,
        "_74": 1345,
        "_476": 477,
        "_478": 479,
        "_374": 34,
        "_373": 34,
        "_480": 1348
    },
    [
        1346
    ],
    {
        "_77": 474,
        "_79": 80,
        "_81": 1347
    },
    "1yBehWRiofl3CcNtvNVvk6",
    [
        1346
    ],
    {
        "_23": 482,
        "_70": 484,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1350,
        "_476": 1351,
        "_480": 1350
    },
    [],
    [],
    {
        "_23": 487,
        "_70": 489,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1353,
        "_476": 1354,
        "_480": 1353
    },
    [],
    [],
    {
        "_23": 493,
        "_70": 495,
        "_366": 498,
        "_72": 498,
        "_368": 71,
        "_74": 1356,
        "_476": 502,
        "_478": 503,
        "_374": 71,
        "_373": 34,
        "_480": 1358
    },
    [
        1357
    ],
    {
        "_77": 501,
        "_79": 111,
        "_81": 101
    },
    [],
    {
        "_23": 505,
        "_70": 507,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1360,
        "_476": 1361,
        "_480": 1360
    },
    [],
    [],
    {
        "_23": 511,
        "_70": 513,
        "_366": 101,
        "_72": 101,
        "_368": 71,
        "_74": 1363,
        "_476": 1364,
        "_480": 1363
    },
    [],
    [],
    {
        "_23": 516,
        "_70": 1366,
        "_366": 1367,
        "_72": 1367,
        "_368": 71,
        "_74": 1368,
        "_476": 1372,
        "_478": 516,
        "_374": 34,
        "_373": 34,
        "_480": 1373
    },
    {
        "_519": 71,
        "_520": 71
    },
    "5UE8g4T56yxUBUYancL7KB:override",
    [
        1369,
        1370
    ],
    {
        "_77": 107,
        "_79": 111,
        "_81": 101
    },
    {
        "_77": 110,
        "_79": 80,
        "_81": 1371
    },
    "5hCRKi4Gs5QJkOanmdVvHU:100.00:4",
    [
        520,
        519
    ],
    [
        1369,
        1370
    ],
    "415386882",
    {
        "_23": 1374,
        "_70": 1376,
        "_366": 101,
        "_72": 101,
        "_368": 71,
        "_74": 1378,
        "_476": 1381,
        "_480": 1378
    },
    {
        "_1377": 34
    },
    "is_voice_mode_entry_point_enabled",
    [
        1379
    ],
    {
        "_77": 1380,
        "_79": 111,
        "_81": 101
    },
    "1644396868",
    [],
    "453021389",
    {
        "_23": 1382,
        "_70": 1384,
        "_366": 101,
        "_72": 101,
        "_368": 71,
        "_74": 1387,
        "_476": 1390,
        "_480": 1387
    },
    {
        "_1385": 34,
        "_1386": 71
    },
    "enable-block-animations",
    "enable-word-animations",
    [
        1388
    ],
    {
        "_77": 1389,
        "_79": 111,
        "_81": 465
    },
    "3016192915",
    [],
    {
        "_23": 527,
        "_70": 529,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1392,
        "_476": 1394,
        "_480": 1392
    },
    [
        1393
    ],
    {
        "_77": 559,
        "_79": 111,
        "_81": 101
    },
    [],
    "474444727",
    {
        "_23": 1395,
        "_70": 1397,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1401,
        "_476": 1402,
        "_480": 1401
    },
    {
        "_1398": 71,
        "_1399": 1400
    },
    "show_custom_instr_message",
    "custom_instr_message_timeout_duration",
    1500,
    [],
    [],
    "590557768",
    {
        "_23": 1403,
        "_70": 1405,
        "_366": 1407,
        "_72": 1407,
        "_368": 71,
        "_74": 1408,
        "_476": 1409,
        "_478": 1410,
        "_374": 71,
        "_373": 71,
        "_480": 1411
    },
    {
        "_1406": 71
    },
    "should_show_return_home_btn",
    "MfvDyM5oEZ1TqWS7cE8et",
    [],
    [
        1406
    ],
    "1022536663",
    [],
    "660512088",
    {
        "_23": 1412,
        "_70": 1414,
        "_366": 101,
        "_72": 101,
        "_368": 71,
        "_74": 1422,
        "_476": 1434,
        "_480": 1422
    },
    {
        "_1415": 34,
        "_1416": 71,
        "_1417": 34,
        "_1418": 34,
        "_1419": 34,
        "_1420": 34,
        "_1421": 34
    },
    "enable_arch_updates",
    "include_legacy_sidebar_contents",
    "include_floating_state",
    "include_share_on_mobile",
    "include_account_settings_move",
    "include_scrolling_behavior_update",
    "include_revised_sidebar_ia",
    [
        1423,
        1425,
        1428,
        1431
    ],
    {
        "_77": 1424,
        "_79": 111,
        "_81": 101
    },
    "2558701922",
    {
        "_77": 1426,
        "_79": 111,
        "_81": 1427
    },
    "735930678",
    "6nGV45RQYtcIGTbPzppBhS",
    {
        "_77": 1429,
        "_79": 111,
        "_81": 1430
    },
    "3011415004",
    "7pUMK6uci7sslAj8bP7VEA",
    {
        "_77": 1432,
        "_79": 111,
        "_81": 1433
    },
    "854062205",
    "66y6sNojVqOdoNf0CX0JYC",
    [],
    "685344542",
    {
        "_23": 1435,
        "_70": 1437,
        "_366": 1439,
        "_72": 1439,
        "_368": 34,
        "_74": 1440,
        "_476": 1442,
        "_478": 1443,
        "_374": 34,
        "_373": 34,
        "_480": 1444
    },
    {
        "_1438": 34,
        "_539": 71
    },
    "is_mobile_enterprise_enabled",
    "3INu3qkV6QoN42TYoP3gja:override",
    [
        1441
    ],
    {
        "_77": 142,
        "_79": 80,
        "_81": 144
    },
    [
        539
    ],
    "1388643772",
    [
        1441
    ],
    "717266490",
    {
        "_23": 1445,
        "_70": 1447,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1463,
        "_476": 1464,
        "_480": 1463
    },
    {
        "_1448": 71,
        "_1449": 71,
        "_1450": 71,
        "_1318": 1451,
        "_1317": 34,
        "_1452": 34,
        "_1320": 34,
        "_1323": 34,
        "_1322": 34,
        "_1453": 416,
        "_1454": 34,
        "_1321": 34,
        "_1455": 34,
        "_1456": 71,
        "_1457": 34,
        "_1458": 1459
    },
    "optimize_initial_modals",
    "defer_memory_modal",
    "enable_v2_cleanup",
    "2099-11-04T00:00:00Z",
    "use_plus_rl_during_onboarding",
    "plus_rl_during_onboarding_minutes_after_creation",
    "enable_mobile_app_upsell_banner",
    "one_tooltip_per_session",
    "one_announcement_tooltip_per_session",
    "onboarding_show_other_option",
    "onboarding_flow_tool_steps",
    [
        1460,
        1461,
        1462
    ],
    "dalle",
    "file_upload",
    "canvas",
    [],
    [],
    {
        "_23": 568,
        "_70": 570,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1466,
        "_476": 1467,
        "_480": 1466
    },
    [],
    [],
    {
        "_23": 581,
        "_70": 583,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1469,
        "_476": 1470,
        "_480": 1469
    },
    [],
    [],
    {
        "_23": 587,
        "_70": 589,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1472,
        "_476": 1473,
        "_480": 1472
    },
    [],
    [],
    {
        "_23": 593,
        "_70": 595,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1475,
        "_476": 1476,
        "_480": 1475
    },
    [],
    [],
    {
        "_23": 602,
        "_70": 604,
        "_366": 101,
        "_72": 101,
        "_368": 71,
        "_74": 1478,
        "_476": 1479,
        "_480": 1478
    },
    [],
    [],
    "1358188185",
    {
        "_23": 1480,
        "_70": 1482,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1485,
        "_476": 1488,
        "_480": 1485
    },
    {
        "_1483": 71,
        "_1484": 34
    },
    "prefetch-models",
    "sidebar-default-close",
    [
        1486
    ],
    {
        "_77": 1487,
        "_79": 111,
        "_81": 101
    },
    "542939804",
    [],
    "1358849452",
    {
        "_23": 1489,
        "_70": 1491,
        "_366": 101,
        "_72": 101,
        "_368": 71,
        "_74": 1493,
        "_476": 1494,
        "_480": 1493
    },
    {
        "_1492": 34
    },
    "disable-ssr",
    [],
    [],
    "1368081792",
    {
        "_23": 1495,
        "_70": 1497,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1500,
        "_476": 1501,
        "_480": 1500
    },
    {
        "_1498": 34,
        "_1499": 34
    },
    "should_show_o3_mini_high_upsell_banner_free_user_to_plus",
    "should_show_o3_mini_high_upsell_banner_plus_user",
    [],
    [],
    {
        "_23": 617,
        "_70": 619,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1503,
        "_476": 1505,
        "_480": 1503
    },
    [
        1504
    ],
    {
        "_77": 634,
        "_79": 111,
        "_81": 101
    },
    [],
    "1578749296",
    {
        "_23": 1506,
        "_70": 1508,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1510,
        "_476": 1511,
        "_480": 1510
    },
    {
        "_1509": 34
    },
    "is_sticky_toggle_off",
    [],
    [],
    {
        "_23": 636,
        "_70": 1513,
        "_366": 498,
        "_72": 498,
        "_368": 34,
        "_74": 1514,
        "_476": 650,
        "_478": 651,
        "_374": 71,
        "_373": 34,
        "_480": 1516
    },
    {
        "_639": 34,
        "_640": 34,
        "_641": 34,
        "_642": 34,
        "_643": 34,
        "_644": 34
    },
    [
        1515
    ],
    {
        "_77": 648,
        "_79": 111,
        "_81": 101
    },
    [],
    {
        "_23": 653,
        "_70": 655,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1518,
        "_476": 1521,
        "_480": 1518
    },
    [
        1519,
        1520
    ],
    {
        "_77": 660,
        "_79": 111,
        "_81": 661
    },
    {
        "_77": 663,
        "_79": 111,
        "_81": 664
    },
    [],
    {
        "_23": 675,
        "_70": 677,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1523,
        "_476": 1525,
        "_480": 1523
    },
    [
        1524
    ],
    {
        "_77": 683,
        "_79": 111,
        "_81": 101
    },
    [],
    {
        "_23": 685,
        "_70": 687,
        "_366": 498,
        "_72": 498,
        "_368": 34,
        "_74": 1527,
        "_476": 693,
        "_478": 694,
        "_374": 71,
        "_373": 34,
        "_480": 1529
    },
    [
        1528
    ],
    {
        "_77": 692,
        "_79": 111,
        "_81": 101
    },
    [],
    "1846737571",
    {
        "_23": 1530,
        "_70": 1532,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1534,
        "_476": 1535,
        "_480": 1534
    },
    {
        "_1533": 34
    },
    "is_upgrade_button_blue",
    [],
    [],
    {
        "_23": 696,
        "_70": 698,
        "_366": 700,
        "_72": 700,
        "_368": 71,
        "_74": 1537,
        "_476": 705,
        "_478": 706,
        "_374": 34,
        "_373": 34,
        "_480": 1539
    },
    [
        1538
    ],
    {
        "_77": 703,
        "_79": 80,
        "_81": 704
    },
    [
        1538
    ],
    "2118136551",
    {
        "_23": 1540,
        "_70": 1542,
        "_366": 101,
        "_72": 101,
        "_368": 71,
        "_74": 1547,
        "_476": 1548,
        "_480": 1547
    },
    {
        "_1543": 34,
        "_1544": 34,
        "_1545": 71,
        "_1546": 71
    },
    "show_cookie_banner_if_qualified",
    "test_dummy",
    "sign_up_button_has_the_word_free",
    "show_cookie_banner_auth_login",
    [],
    [],
    "2149763392",
    {
        "_23": 1549,
        "_70": 1551,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1554,
        "_476": 1557,
        "_480": 1554
    },
    {
        "_1552": 34,
        "_1553": 34
    },
    "show-in-main-composer",
    "show-model-picker",
    [
        1555
    ],
    {
        "_77": 1556,
        "_79": 111,
        "_81": 101
    },
    "4151101559",
    [],
    {
        "_23": 708,
        "_70": 710,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1559,
        "_476": 1560,
        "_480": 1559
    },
    [],
    [],
    "2259187367",
    {
        "_23": 1561,
        "_70": 1563,
        "_366": 101,
        "_72": 101,
        "_368": 71,
        "_74": 1574,
        "_476": 1575,
        "_480": 1574
    },
    {
        "_1564": 34,
        "_1565": 1566,
        "_1567": 1568,
        "_1569": 71,
        "_1570": 1571,
        "_1572": 34,
        "_1573": 1256
    },
    "enable_nux",
    "start_time",
    "2099-01-01T00:00:00Z",
    "end_time",
    "2000-01-01T00:00:00Z",
    "use_multi_input",
    "force_madlibs_param_name",
    "madlibs_0203",
    "enable_additional_categories",
    "additional_category",
    [],
    [],
    "2505516353",
    {
        "_23": 1576,
        "_70": 1578,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1580,
        "_476": 1581,
        "_480": 1580
    },
    {
        "_1579": 71
    },
    "android-keyboard-layout",
    [],
    [],
    "2670443078",
    {
        "_23": 1582,
        "_70": 1584,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1586,
        "_476": 1587,
        "_480": 1586
    },
    {
        "_1585": 71
    },
    "is_gating_fix_enabled",
    [],
    [],
    "2716194794",
    {
        "_23": 1588,
        "_70": 1590,
        "_366": 498,
        "_72": 498,
        "_368": 34,
        "_74": 1591,
        "_476": 1594,
        "_478": 1595,
        "_374": 71,
        "_373": 34,
        "_480": 1596
    },
    {
        "_750": 34
    },
    [
        1592
    ],
    {
        "_77": 1593,
        "_79": 111,
        "_81": 101
    },
    "2849926832",
    [
        750
    ],
    "2435265903",
    [],
    "2723963139",
    {
        "_23": 1597,
        "_70": 1599,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1611,
        "_476": 1612,
        "_480": 1611
    },
    {
        "_1600": 34,
        "_1601": 34,
        "_1602": 71,
        "_1603": 71,
        "_1604": 71,
        "_1605": 1606,
        "_1607": 71,
        "_1608": 34,
        "_1609": 34,
        "_1610": 465
    },
    "is_dynamic_model_enabled",
    "show_message_model_info",
    "show_message_regenerate_model_selector",
    "is_conversation_model_switching_allowed",
    "show_rate_limit_downgrade_banner",
    "config",
    {},
    "show_message_regenerate_model_selector_on_every_message",
    "is_AG8PqS2q_enabled",
    "is_chive_enabled",
    "sahara_model_id_override",
    [],
    [],
    "2775247110",
    {
        "_23": 1613,
        "_70": 1615,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1618,
        "_476": 1619,
        "_480": 1618
    },
    {
        "_1616": 34,
        "_1617": 34
    },
    "show_pro_badge",
    "show_plan_type_badge",
    [],
    [],
    "2840731323",
    {
        "_23": 1620,
        "_70": 1622,
        "_366": 101,
        "_72": 101,
        "_368": 71,
        "_74": 1624,
        "_476": 1626,
        "_480": 1624
    },
    {
        "_626": 71,
        "_1623": 71
    },
    "is_direct_continue_enabled",
    [
        1625
    ],
    {
        "_77": 1031,
        "_79": 111,
        "_81": 101
    },
    [],
    "2888142241",
    {
        "_23": 1627,
        "_70": 1629,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1631,
        "_476": 1632,
        "_480": 1631
    },
    {
        "_1630": 71
    },
    "is_upgrade_in_settings",
    [],
    [],
    "2932223118",
    {
        "_23": 1633,
        "_70": 1635,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1636,
        "_476": 1637,
        "_480": 1636
    },
    {
        "_538": 71
    },
    [],
    [],
    "2972011003",
    {
        "_23": 1638,
        "_70": 1640,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1643,
        "_476": 1644,
        "_480": 1643
    },
    {
        "_1641": 71,
        "_1642": 34
    },
    "user_context_message_search_tools_default",
    "search_tool_holdout_enabled",
    [],
    [],
    {
        "_23": 715,
        "_70": 1646,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1647,
        "_476": 1651,
        "_480": 1647
    },
    {
        "_718": 71,
        "_719": 34
    },
    [
        1648,
        1649
    ],
    {
        "_77": 723,
        "_79": 111,
        "_81": 724
    },
    {
        "_77": 726,
        "_79": 111,
        "_81": 1650
    },
    "66covmutTYx82FWVUlZAqF",
    [],
    "3119715334",
    {
        "_23": 1652,
        "_70": 1654,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1657,
        "_476": 1658,
        "_480": 1657
    },
    {
        "_1655": 34,
        "_1656": 34
    },
    "should-enable-hojicha",
    "should-enable-skip",
    [],
    [],
    "3206655705",
    {
        "_23": 1659,
        "_70": 1661,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1663,
        "_476": 1664,
        "_480": 1663
    },
    {
        "_1662": 71
    },
    "enable_new_ux",
    [],
    [],
    "3434623093",
    {
        "_23": 1665,
        "_70": 1667,
        "_366": 101,
        "_72": 101,
        "_368": 71,
        "_74": 1673,
        "_476": 1682,
        "_480": 1673
    },
    {
        "_1668": 71,
        "_1669": 1670,
        "_1671": 71,
        "_1672": 71
    },
    "with-attach-upsell",
    "labels",
    "all",
    "with-voice-upsell",
    "with-reason-upsell",
    [
        1674,
        1676,
        1678,
        1680
    ],
    {
        "_77": 1675,
        "_79": 111,
        "_81": 101
    },
    "1604099973",
    {
        "_77": 1677,
        "_79": 111,
        "_81": 101
    },
    "470066910",
    {
        "_77": 1679,
        "_79": 111,
        "_81": 101
    },
    "1932133792",
    {
        "_77": 1681,
        "_79": 111,
        "_81": 101
    },
    "4175621034",
    [],
    {
        "_23": 737,
        "_70": 1684,
        "_366": 1686,
        "_72": 1686,
        "_368": 34,
        "_74": 1687,
        "_476": 744,
        "_478": 745,
        "_374": 71,
        "_373": 71,
        "_480": 1690
    },
    {
        "_639": 71,
        "_740": 1685,
        "_642": 71,
        "_641": 71,
        "_640": 34
    },
    10,
    "2FurbaJwwLPFodZHhOZyBO",
    [
        1688
    ],
    {
        "_77": 743,
        "_79": 80,
        "_81": 1689
    },
    "1FzsKf0T7jWwTRKiSrbUld:100.00:4",
    [],
    {
        "_23": 747,
        "_70": 749,
        "_366": 751,
        "_72": 751,
        "_368": 71,
        "_74": 1692,
        "_476": 753,
        "_478": 754,
        "_374": 34,
        "_373": 34,
        "_480": 1693
    },
    [],
    [],
    {
        "_23": 756,
        "_70": 758,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1695,
        "_476": 1696,
        "_480": 1695
    },
    [],
    [],
    "3533083032",
    {
        "_23": 1697,
        "_70": 1699,
        "_366": 101,
        "_72": 101,
        "_368": 71,
        "_74": 1722,
        "_476": 1723,
        "_480": 1722
    },
    {
        "_1700": 71,
        "_1701": 71,
        "_1702": 1703,
        "_1704": 34,
        "_1705": 34,
        "_1706": 71,
        "_1707": 34,
        "_1708": 34,
        "_1709": 34,
        "_1710": 34,
        "_1711": 1712,
        "_1713": 1714,
        "_1715": 1716,
        "_1717": 1718,
        "_1719": 1720,
        "_1721": 465
    },
    "enable_new_homepage_anon",
    "filter_prompt_by_model",
    "headline_option",
    "HELP_WITH",
    "disclaimer_color_adjust",
    "show_composer_header",
    "enable_new_mobile",
    "enable_cached_response",
    "show_dalle_starter_prompts",
    "use_modapi_in_autocomplete",
    "use_memory_in_model_autocomplete",
    "autocomplete_max_char",
    32,
    "search_autocomplete_mode",
    "BING",
    "autocomplete_min_char",
    4,
    "autocomplete_mode",
    "INDEX",
    "num_completions_to_fetch_from_index",
    8,
    "india_first_prompt",
    [],
    [],
    {
        "_23": 762,
        "_70": 764,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1725,
        "_476": 1726,
        "_480": 1725
    },
    [],
    [],
    "3606233934",
    {
        "_23": 1727,
        "_70": 1729,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1733,
        "_476": 1734,
        "_480": 1733
    },
    {
        "_1730": 1731,
        "_1732": 34
    },
    "link",
    "non",
    "enable_notifications_feed",
    [],
    [],
    "3613709240",
    {
        "_23": 1735,
        "_70": 1737,
        "_366": 101,
        "_72": 101,
        "_368": 71,
        "_74": 1739,
        "_476": 1740,
        "_480": 1739
    },
    {
        "_1738": 71
    },
    "shouldRefreshAccessToken",
    [],
    [],
    {
        "_23": 768,
        "_70": 770,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1742,
        "_476": 1745,
        "_480": 1742
    },
    [
        1743,
        1744
    ],
    {
        "_77": 798,
        "_79": 111,
        "_81": 159
    },
    {
        "_77": 800,
        "_79": 111,
        "_81": 159
    },
    [],
    {
        "_23": 818,
        "_70": 820,
        "_366": 101,
        "_72": 101,
        "_368": 71,
        "_74": 1747,
        "_476": 1748,
        "_480": 1747
    },
    [],
    [],
    "3737571708",
    {
        "_23": 1749,
        "_70": 1751,
        "_366": 101,
        "_72": 101,
        "_368": 71,
        "_74": 1754,
        "_476": 1755,
        "_480": 1754
    },
    {
        "_1752": 1753
    },
    "sidebar_type",
    "slick",
    [],
    [],
    "3768341700",
    {
        "_23": 1756,
        "_70": 1758,
        "_366": 1766,
        "_72": 1766,
        "_368": 34,
        "_74": 1767,
        "_476": 1768,
        "_478": 1769,
        "_374": 71,
        "_373": 71,
        "_480": 1770
    },
    {
        "_540": 34,
        "_1759": 34,
        "_1760": 34,
        "_1761": 71,
        "_1762": 71,
        "_1763": 71,
        "_1764": 34,
        "_1765": 34
    },
    "remove_early_access_upsell",
    "is_produce_text_design",
    "is_produce_design",
    "is_country_selector_enabled",
    "is_vat_information_enabled",
    "is_vat_information_with_amount_enabled",
    "is_team_pricing_vat_disclaimer_enabled",
    "65VHFqyIytQJKjgykJm4UQ",
    [],
    [
        1763,
        1762,
        1764,
        1765
    ],
    "2782616826",
    [],
    "3927927759",
    {
        "_23": 1771,
        "_70": 1773,
        "_366": 101,
        "_72": 101,
        "_368": 71,
        "_74": 1774,
        "_476": 1775,
        "_480": 1774
    },
    {
        "_1454": 71
    },
    [],
    [],
    {
        "_23": 825,
        "_70": 827,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1777,
        "_476": 1778,
        "_480": 1777
    },
    [],
    [],
    "4020668365",
    {
        "_23": 1779,
        "_70": 1781,
        "_366": 101,
        "_72": 101,
        "_368": 71,
        "_74": 1782,
        "_476": 1783,
        "_480": 1782
    },
    {
        "_1564": 34,
        "_1565": 1566,
        "_1567": 1568,
        "_1569": 34
    },
    [],
    [],
    "4031588851",
    {
        "_23": 1784,
        "_70": 1786,
        "_366": 101,
        "_72": 101,
        "_368": 34,
        "_74": 1809,
        "_476": 1813,
        "_480": 1809
    },
    {
        "_1787": 71,
        "_1788": 71,
        "_1789": 71,
        "_1790": 71,
        "_1791": 34,
        "_1792": 34,
        "_1717": 1718,
        "_1793": 1794,
        "_1715": 1716,
        "_1711": 1712,
        "_1702": 1703,
        "_1710": 34,
        "_1795": 34,
        "_1709": 34,
        "_1796": 1797,
        "_1798": 71,
        "_1799": 465,
        "_1706": 71,
        "_1713": 1714,
        "_1800": 34,
        "_1801": 1802,
        "_1719": 1720,
        "_1803": 34,
        "_1804": 34,
        "_790": 791,
        "_1805": 34,
        "_1806": 1807,
        "_1808": 71,
        "_1721": 465
    },
    "enable_hardcoded_vision_prompts",
    "enable_hardcoded_file_document_prompts",
    "enable_hardcoded_data_vis_prompts",
    "enable_hardcoded_browse_prompts",
    "is_two_line",
    "enable_new_homepage",
    "starter_prompt_ranking_algorithm",
    "homepage_v2",
    "filter_starter_prompt_by_model",
    "autocomplete_qualified_start_date",
    "2000-10-11T00:00:00Z",
    "enable_new_autocomplete_homepage",
    "model_talks_option",
    "enable_hardcoded_onboarding_prompt",
    "autocomplete_fetch_interval",
    200,
    "enable_recommend_prompts",
    "enable_ask_me_prompts",
    "enable_reasoning_prompts_0202",
    "dream_type",
    "user_knowledge_memories",
    "web-disable",
    [
        1810
    ],
    {
        "_77": 1811,
        "_79": 111,
        "_81": 1812
    },
    "4273941502",
    "1nGrz4l6GM0LgZvm0pDCtp:2.00:1",
    [],
    {
        "_23": 832,
        "_70": 834,
        "_366": 101,
        "_72": 101,
        "_368": 71,
        "_74": 1815,
        "_476": 1816,
        "_480": 1815
    },
    [],
    [],
    "4250072504",
    {
        "_23": 1817,
        "_70": 1819,
        "_366": 1822,
        "_72": 1822,
        "_368": 34,
        "_74": 1823,
        "_476": 1825,
        "_478": 1826,
        "_374": 34,
        "_373": 34,
        "_480": 1827
    },
    {
        "_539": 71,
        "_1820": 34,
        "_1821": 34
    },
    "is_enterprise_desktop_enabled",
    "is_desktop_enterprise_enabled",
    "3HX7vpdJsUkuyCUEL4V9cE:override",
    [
        1824
    ],
    {
        "_77": 142,
        "_79": 80,
        "_81": 144
    },
    [
        539
    ],
    "3311396813",
    [
        1824
    ],
    {},
    {
        "_844": 845,
        "_846": 847
    },
    {
        "_852": 22,
        "_853": 1831
    },
    {
        "_855": 856,
        "_857": 856,
        "_858": 856
    },
    {
        "_852": 22,
        "_864": 865,
        "_866": 1833,
        "_853": 1831,
        "_1838": 1839,
        "_1840": 1836,
        "_871": 1841,
        "_60": 61
    },
    {
        "_1834": 71,
        "_1835": 1836,
        "_1837": 34,
        "_869": 34,
        "_870": 16
    },
    "has_logged_in_before",
    "user_agent",
    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36 Edg/134.0.0.0",
    "is_punch_out_user",
    "ip",
    "1.1.1.1",
    "userAgent",
    {
        "_873": 874
    },
    "isNoAuthEnabled",
    "userRegion",
    "New York",
    "US",
    "cfIpLatitude",
    null,
    "cfIpLongitude",
    null,
    "cfIpCity",
    null,
    "isUserInNewCookieConsentFlow",
    "isUserInPioneerHR",
    "isUserEligibleForPioneer",
    "isUserEligibleForMaverick",
    "isIos",
    "isAndroidChrome",
    "isElectron",
    "windowStyle",
    "cspScriptNonce",
    "20f285f6-3f39-465d-8802-567f078c6c55",
    "routes/_conversation",
    {
        "_1864": 1865,
        "_1866": -7
    },
    "rq:[\"models\",\"{\\\"isHistoryDisabled\\\":false}\"]",
    [
        "P",
        1865
    ],
    "prefetchSearch",
    "routes/_conversation._index",
    "actionData",
    "errors"
]
````

## File: templates/chatgpt_context_2.json
````json
[
    {
        "_1871": 1872,
        "_1873": 1874,
        "_1964": 43
    },
    "__type",
    "AccountState",
    "accountItems",
    [
        1875
    ],
    [
        "SingleFetchClassInstance",
        1876
    ],
    {
        "_1877": 1878
    },
    "data",
    {
        "_21": 43,
        "_1879": 1880,
        "_1881": 1882,
        "_23": -5,
        "_1883": -5,
        "_1884": -5,
        "_46": 47,
        "_1885": 1886,
        "_1887": -5,
        "_1888": 1889,
        "_1890": 34,
        "_1891": 1892,
        "_1918": 1919,
        "_1962": 71,
        "_1963": -5
    },
    "residencyRegion",
    "no_constraint",
    "accountUserId",
    "user-chatgpt__chatgpt",
    "profilePictureId",
    "profilePictureUrl",
    "role",
    "account-owner",
    "organizationId",
    "promoData",
    {},
    "deactivated",
    "subscriptionStatus",
    {
        "_1893": 1894,
        "_1895": 71,
        "_1896": 34,
        "_1897": 1898,
        "_1899": 1900,
        "_44": 45,
        "_1901": 1902,
        "_1903": -7,
        "_1904": 71,
        "_1905": 71,
        "_1906": 1907,
        "_1908": 1909,
        "_1915": 34,
        "_1916": -5,
        "_1917": 34
    },
    "billingPeriod",
    "monthly",
    "hasPaidSubscription",
    "isActiveSubscriptionGratis",
    "billingCurrency",
    "USD",
    "subscriptionPlan",
    "chatgptproplan",
    "subscriptionExpiresAt",
    "2524579200",
    "scheduledPlanChange",
    "wasPaidCustomer",
    "hasCustomerObject",
    "processorEntity",
    "openai_llc",
    "lastActiveSubscription",
    {
        "_1910": 1911,
        "_1912": 1913,
        "_1914": 71
    },
    "subscription_id",
    "c3f25801-20a7-436e-a143-b42924a6e5ea",
    "purchase_origin_platform",
    "chatgpt_web",
    "will_renew",
    "isResellerHosted",
    "discount",
    "isEligibleForCancellationPromotion",
    "features",
    [
        1920,
        1921,
        1922,
        1923,
        1462,
        1924,
        1925,
        1926,
        1927,
        1928,
        1929,
        1930,
        1931,
        1932,
        1933,
        1934,
        1935,
        1936,
        1937,
        1938,
        1939,
        1940,
        1941,
        1942,
        1943,
        33,
        1944,
        1945,
        1946,
        1947,
        1948,
        1949,
        1950,
        1951,
        1952,
        1953,
        1954,
        1955,
        1956,
        1957,
        1958,
        1959,
        1960,
        1961
    ],
    "beta_features",
    "bizmo_settings",
    "breeze_available",
    "browsing_available",
    "canvas_code_execution",
    "canvas_code_network_access",
    "canvas_o1",
    "canvas_opt_in",
    "caterpillar",
    "chart_serialization",
    "chat_preferences_available",
    "chatgpt_ios_attest",
    "code_interpreter_available",
    "d3_controls",
    "d3_editor",
    "d3_editor_gpts",
    "dalle_3",
    "gizmo_canvas_toggle",
    "gizmo_reviews",
    "gizmo_support_emails",
    "gpt_4_5",
    "graphite",
    "image_gen_tool_enabled",
    "jawbone_model_access",
    "model_ab_use_v2",
    "model_switcher",
    "new_plugin_oauth_endpoint",
    "no_auth_training_enabled_by_default",
    "o1_launch",
    "o3-mini",
    "plugins_available",
    "privacy_policy_nov_2023",
    "search_tool",
    "sentinel_enabled_for_subscription",
    "share_multimodal_links",
    "shareable_links",
    "snc",
    "starter_prompts",
    "sunshine_available",
    "user_settings_announcements",
    "video_screen_sharing",
    "voice_advanced_ga",
    "canAccessWithCurrentSession",
    "ssoConnectionName",
    "currentAccountId"
]
````

## File: templates/chatgpt.html
````html
<!DOCTYPE html>
<html data-build="prod-c3c153491ab77f6fdaff0b4ee8515dafacb3c18a" dir="ltr" class="">
    <head>
        <meta charSet="UTF-8"/>
        <meta name="viewport" content="width=device-width, initial-scale=1"/>
        <link rel="preload" as="image" href="/assets/sora-mutf8tav.webp"/>
        <link rel="preconnect" href=""/>
        <link rel="preconnect" href=""/>
        <meta name="robots" content="index, follow"/>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
        <meta name="apple-itunes-app" content="app-id=6448311069"/>
        <meta name="dd-trace-id" content="14093928210743749173"/>
        <meta name="dd-trace-time" content="1742986866029"/>
        <link rel="icon" href="/assets/favicon-miwirzcw.ico" sizes="32x32"/>
        {% raw %}
        <link rel="icon" href="data:image/svg+xml,%3csvg%20xmlns=&#x27;http://www.w3.org/2000/svg&#x27;%20width=&#x27;180&#x27;%20height=&#x27;180&#x27;%20fill=&#x27;none&#x27;%3e%3cstyle%3e%20:root%20{%20--primary-fill:%20%23000;%20--secondary-fill:%20%23fff;%20}%20@media%20(prefers-color-scheme:%20dark)%20{%20:root%20{%20--primary-fill:%20%23fff;%20--secondary-fill:%20%23000;%20}%20}%20%3c/style%3e%3cg%20clip-path=&#x27;url(%23a)&#x27;%3e%3crect%20width=&#x27;180&#x27;%20height=&#x27;180&#x27;%20fill=&#x27;var(--primary-fill)&#x27;%20rx=&#x27;90&#x27;%20/%3e%3cg%20clip-path=&#x27;url(%23b)&#x27;%3e%3cpath%20fill=&#x27;var(--secondary-fill)&#x27;%20d=&#x27;M75.91%2073.628V62.232c0-.96.36-1.68%201.199-2.16l22.912-13.194c3.119-1.8%206.838-2.639%2010.676-2.639%2014.394%200%2023.511%2011.157%2023.511%2023.032%200%20.839%200%201.799-.12%202.758l-23.752-13.914c-1.439-.84-2.879-.84-4.318%200L75.91%2073.627Zm53.499%2044.383v-27.23c0-1.68-.72-2.88-2.159-3.719L97.142%2069.55l9.836-5.638c.839-.48%201.559-.48%202.399%200l22.912%2013.195c6.598%203.839%2011.035%2011.995%2011.035%2019.912%200%209.116-5.397%2017.513-13.915%2020.992v.001Zm-60.577-23.99-9.836-5.758c-.84-.48-1.2-1.2-1.2-2.16v-26.39c0-12.834%209.837-22.55%2023.152-22.55%205.039%200%209.716%201.679%2013.676%204.678L70.993%2055.516c-1.44.84-2.16%202.039-2.16%203.719v34.787-.002Zm21.173%2012.234L75.91%2098.339V81.546l14.095-7.917%2014.094%207.917v16.793l-14.094%207.916Zm9.056%2036.467c-5.038%200-9.716-1.68-13.675-4.678l23.631-13.676c1.439-.839%202.159-2.038%202.159-3.718V85.863l9.956%205.757c.84.48%201.2%201.2%201.2%202.16v26.389c0%2012.835-9.957%2022.552-23.27%2022.552v.001Zm-28.43-26.75L47.72%20102.778c-6.599-3.84-11.036-11.996-11.036-19.913%200-9.236%205.518-17.513%2014.034-20.992v27.35c0%201.68.72%202.879%202.16%203.718l29.989%2017.393-9.837%205.638c-.84.48-1.56.48-2.399%200Zm-1.318%2019.673c-13.555%200-23.512-10.196-23.512-22.792%200-.959.12-1.919.24-2.879l23.63%2013.675c1.44.84%202.88.84%204.32%200l30.108-17.392v11.395c0%20.96-.361%201.68-1.2%202.16l-22.912%2013.194c-3.119%201.8-6.837%202.639-10.675%202.639Zm29.748%2014.274c14.515%200%2026.63-10.316%2029.39-23.991%2013.434-3.479%2022.071-16.074%2022.071-28.91%200-8.396-3.598-16.553-10.076-22.43.6-2.52.96-5.039.96-7.557%200-17.153-13.915-29.99-29.989-29.99-3.239%200-6.358.48-9.477%201.56-5.398-5.278-12.835-8.637-20.992-8.637-14.515%200-26.63%2010.316-29.39%2023.991-13.434%203.48-22.07%2016.074-22.07%2028.91%200%208.396%203.598%2016.553%2010.075%2022.431-.6%202.519-.96%205.038-.96%207.556%200%2017.154%2013.915%2029.989%2029.99%2029.989%203.238%200%206.357-.479%209.476-1.559%205.397%205.278%2012.835%208.637%2020.992%208.637Z&#x27;%20/%3e%3c/g%3e%3c/g%3e%3cdefs%3e%3cclipPath%20id=&#x27;a&#x27;%3e%3cpath%20fill=&#x27;var(--primary-fill)&#x27;%20d=&#x27;M0%200h180v180H0z&#x27;%20/%3e%3c/clipPath%3e%3cclipPath%20id=&#x27;b&#x27;%3e%3cpath%20fill=&#x27;var(--primary-fill)&#x27;%20d=&#x27;M29.487%2029.964h121.035v119.954H29.487z&#x27;%20/%3e%3c/clipPath%3e%3c/defs%3e%3c/svg%3e" type="image/svg+xml"/>
        {% endraw %}
        <link rel="apple-touch-icon" sizes="180x180" href="/assets/favicon-180x180-od45eci6.webp"/>
        <title>ChatGPT</title>
        <meta name="description" content="ChatGPT helps you get answers, find inspiration and be more productive. It is free to use and easy to try. Just ask and ChatGPT can help with writing, learning, brainstorming and more."/>
        <meta name="keyword" content="ai chat,ai,chap gpt,chat gbt,chat gpt 3,chat gpt login,chat gpt website,chat gpt,chat gtp,chat openai,chat,chatai,chatbot gpt,chatg,chatgpt login,chatgpt,gpt chat,open ai,openai chat,openai chatgpt,openai"/>
        <meta property="og:description" content="A conversational AI system that listens, learns, and challenges"/>
        <meta property="og:title" content="ChatGPT"/>
        <meta property="og:image" content="/assets/chatgpt-share-og-u7j5uyao.webp"/>
        <meta property="og:url" content="https://chatgpt.com"/>
        <link rel="modulepreload" href="/assets/manifest-532b61a2.js"/>
        <link rel="modulepreload" href="/assets/mwrrfx5gxel0ghh2.js"/>
        <link rel="modulepreload" href="/assets/fs6h2trisr1juto6.js"/>
        <link rel="modulepreload" href="/assets/njq6ygky3ttysgdk.js"/>
        <link rel="modulepreload" href="/assets/jandtmjen68hvio9.js"/>
        <link rel="modulepreload" href="/assets/fj65uqnkvh37vl4n.js"/>
        <link rel="modulepreload" href="/assets/d7a6rc8fexfzu7dt.js"/>
        <link rel="modulepreload" href="/assets/fu6bathgya5b9rdr.js"/>
        <link rel="modulepreload" href="/assets/lj0prj6na7gf1jpf.js"/>
        <link rel="modulepreload" href="/assets/csg4i0wk6hi0jvkd.js"/>
        <link rel="modulepreload" href="/assets/kukctmibghbbqwq7.js"/>
        <link rel="modulepreload" href="/assets/ijutvl794yycgs5p.js"/>
        <link rel="modulepreload" href="/assets/n9whwsbg0rx5cs8k.js"/>
        <link rel="modulepreload" href="/assets/gunpso70jnmarvod.js"/>
        <link rel="modulepreload" href="/assets/jyg6m0czwgwjnl3i.js"/>
        <link rel="modulepreload" href="/assets/en8ogvlmn3lccz09.js"/>
        <link rel="modulepreload" href="/assets/gvhvfcdqhc5unmaw.js"/>
        <link rel="modulepreload" href="/assets/wqepkoieaii7q4ep.js"/>
        <link rel="modulepreload" href="/assets/gfo75oujoeaedm5y.js"/>
        <link rel="modulepreload" href="/assets/n1h3bv0q1kp4dflt.js"/>
        <link rel="modulepreload" href="/assets/o634nz6bc7e8bevz.js"/>
        <link rel="modulepreload" href="/assets/ohq33fk3dbalo2mk.js"/>
        <link rel="modulepreload" href="/assets/cffgjlidl1olpt2q.js"/>
        <link rel="modulepreload" href="/assets/i6ze5suolzahci5v.js"/>
        <link rel="modulepreload" href="/assets/npeck75bz46c490i.js"/>
        <link rel="modulepreload" href="/assets/m97sb2fnwj0a6vyo.js"/>
        <link rel="modulepreload" href="/assets/cm0hdcditutj2jz2.js"/>
        <link rel="modulepreload" href="/assets/jcoe6v231ph25bfi.js"/>
        <link rel="modulepreload" href="/assets/mfnl9tbidoj0ik29.js"/>
        <link rel="modulepreload" href="/assets/lprjl238d50elo8f.js"/>
        <link rel="modulepreload" href="/assets/ovc3ksv6dtkjao8i.js"/>
        <link rel="stylesheet" href="/assets/root-cl538jor.css"/>
        <link rel="stylesheet" href="/assets/conversation-small-m9kq0y4e.css"/>
        <script nonce="20f285f6-3f39-465d-8802-567f078c6c55">
            !function initScrollTimelineInline() {
                try {
                    if (CSS.supports("animation-timeline: --works"))
                        return;
                    var t = new Map;
                    document.addEventListener("animationstart", (n => {
                        if (!(n.target instanceof HTMLElement))
                            return;
                        const e = n.target.getAnimations().filter((t => t.animationName === n.animationName));
                        t.set(n.target, e)
                    }
                    )),
                    document.addEventListener("scrolltimelineload", (n => {
                        t.forEach(( (t, e) => {
                            t.forEach((t => {
                                n.detail.upgradeAnimation(t, e)
                            }
                            ))
                        }
                        )),
                        t.clear()
                    }
                    ), {
                        once: !0
                    })
                } catch {}
            }();
        </script>
{{ clear_localstorage_script | safe }}
    </head>
    <body class="">
        <script>
            !function() {
                try {
                    var d = document.documentElement
                      , c = d.classList;
                    c.remove('light', 'dark');
                    var e = localStorage.getItem('theme');
                    if ('system' === e || (!e && true)) {
                        var t = '(prefers-color-scheme: dark)'
                          , m = window.matchMedia(t);
                        if (m.media !== t || m.matches) {
                            d.style.colorScheme = 'dark';
                            c.add('dark')
                        } else {
                            d.style.colorScheme = 'light';
                            c.add('light')
                        }
                    } else if (e) {
                        c.add(e || '')
                    }
                    if (e === 'light' || e === 'dark')
                        d.style.colorScheme = e
                } catch (e) {}
            }()
        </script>
        <div class="flex h-full w-full flex-col">
            <div class="relative flex h-full w-full flex-1 overflow-hidden transition-colors z-0">
                <div class="relative flex h-full w-full flex-row overflow-hidden">
                    <div class="z-[21] flex-shrink-0 overflow-x-hidden bg-token-sidebar-surface-primary [view-transition-name:--sidebar-slideover] max-md:!w-0" style="width:260px">
                        <div class="h-full w-[260px]">
                            <div class="flex h-full min-h-0 flex-col">
                                <div class="draggable relative h-full w-full flex-1 items-start border-white/20">
                                    <h2 style="position:absolute;border:0;width:1px;height:1px;padding:0;margin:-1px;overflow:hidden;clip:rect(0, 0, 0, 0);white-space:nowrap;word-wrap:normal">Chat history</h2>
                                    <nav class="flex h-full w-full flex-col pl-3" aria-label="Chat history">
                                        <div class="flex justify-between flex h-header-height items-center xs:pr-3">
                                            <span class="hidden"></span>
                                            <span class="flex" data-state="closed">
                                                <button class="h-10 rounded-lg px-2 text-token-text-secondary focus-visible:bg-token-surface-hover focus-visible:outline-0 enabled:hover:bg-token-surface-hover disabled:text-token-text-quaternary no-draggable" aria-label="Close sidebar" data-testid="close-sidebar-button">
                                                    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="icon-xl-heavy max-md:hidden">
                                                        <path fill-rule="evenodd" clip-rule="evenodd" d="M8.85719 3H15.1428C16.2266 2.99999 17.1007 2.99998 17.8086 3.05782C18.5375 3.11737 19.1777 3.24318 19.77 3.54497C20.7108 4.02433 21.4757 4.78924 21.955 5.73005C22.2568 6.32234 22.3826 6.96253 22.4422 7.69138C22.5 8.39925 22.5 9.27339 22.5 10.3572V13.6428C22.5 14.7266 22.5 15.6008 22.4422 16.3086C22.3826 17.0375 22.2568 17.6777 21.955 18.27C21.4757 19.2108 20.7108 19.9757 19.77 20.455C19.1777 20.7568 18.5375 20.8826 17.8086 20.9422C17.1008 21 16.2266 21 15.1428 21H8.85717C7.77339 21 6.89925 21 6.19138 20.9422C5.46253 20.8826 4.82234 20.7568 4.23005 20.455C3.28924 19.9757 2.52433 19.2108 2.04497 18.27C1.74318 17.6777 1.61737 17.0375 1.55782 16.3086C1.49998 15.6007 1.49999 14.7266 1.5 13.6428V10.3572C1.49999 9.27341 1.49998 8.39926 1.55782 7.69138C1.61737 6.96253 1.74318 6.32234 2.04497 5.73005C2.52433 4.78924 3.28924 4.02433 4.23005 3.54497C4.82234 3.24318 5.46253 3.11737 6.19138 3.05782C6.89926 2.99998 7.77341 2.99999 8.85719 3ZM6.35424 5.05118C5.74907 5.10062 5.40138 5.19279 5.13803 5.32698C4.57354 5.6146 4.1146 6.07354 3.82698 6.63803C3.69279 6.90138 3.60062 7.24907 3.55118 7.85424C3.50078 8.47108 3.5 9.26339 3.5 10.4V13.6C3.5 14.7366 3.50078 15.5289 3.55118 16.1458C3.60062 16.7509 3.69279 17.0986 3.82698 17.362C4.1146 17.9265 4.57354 18.3854 5.13803 18.673C5.40138 18.8072 5.74907 18.8994 6.35424 18.9488C6.97108 18.9992 7.76339 19 8.9 19H9.5V5H8.9C7.76339 5 6.97108 5.00078 6.35424 5.05118ZM11.5 5V19H15.1C16.2366 19 17.0289 18.9992 17.6458 18.9488C18.2509 18.8994 18.5986 18.8072 18.862 18.673C19.4265 18.3854 19.8854 17.9265 20.173 17.362C20.3072 17.0986 20.3994 16.7509 20.4488 16.1458C20.4992 15.5289 20.5 14.7366 20.5 13.6V10.4C20.5 9.26339 20.4992 8.47108 20.4488 7.85424C20.3994 7.24907 20.3072 6.90138 20.173 6.63803C19.8854 6.07354 19.4265 5.6146 18.862 5.32698C18.5986 5.19279 18.2509 5.10062 17.6458 5.05118C17.0289 5.00078 16.2366 5 15.1 5H11.5ZM5 8.5C5 7.94772 5.44772 7.5 6 7.5H7C7.55229 7.5 8 7.94772 8 8.5C8 9.05229 7.55229 9.5 7 9.5H6C5.44772 9.5 5 9.05229 5 8.5ZM5 12C5 11.4477 5.44772 11 6 11H7C7.55229 11 8 11.4477 8 12C8 12.5523 7.55229 13 7 13H6C5.44772 13 5 12.5523 5 12Z" fill="currentColor"></path>
                                                    </svg>
                                                    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="icon-xl-heavy md:hidden">
                                                        <path fill-rule="evenodd" clip-rule="evenodd" d="M3 8C3 7.44772 3.44772 7 4 7H20C20.5523 7 21 7.44772 21 8C21 8.55228 20.5523 9 20 9H4C3.44772 9 3 8.55228 3 8ZM3 16C3 15.4477 3.44772 15 4 15H14C14.5523 15 15 15.4477 15 16C15 16.5523 14.5523 17 14 17H4C3.44772 17 3 16.5523 3 16Z" fill="currentColor"></path>
                                                    </svg>
                                                </button>
                                            </span>
                                            <div class="flex">
                                                <span class="hidden"></span>
                                                <span class="flex" data-state="closed">
                                                    <button aria-label="Ctrl K" class="h-10 rounded-lg px-2 text-token-text-secondary focus-visible:bg-token-surface-hover focus-visible:outline-0 enabled:hover:bg-token-surface-hover disabled:text-token-text-quaternary">
                                                        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="icon-xl-heavy">
                                                            <path fill-rule="evenodd" clip-rule="evenodd" d="M10.75 4.25C7.16015 4.25 4.25 7.16015 4.25 10.75C4.25 14.3399 7.16015 17.25 10.75 17.25C14.3399 17.25 17.25 14.3399 17.25 10.75C17.25 7.16015 14.3399 4.25 10.75 4.25ZM2.25 10.75C2.25 6.05558 6.05558 2.25 10.75 2.25C15.4444 2.25 19.25 6.05558 19.25 10.75C19.25 12.7369 18.5683 14.5645 17.426 16.0118L21.4571 20.0429C21.8476 20.4334 21.8476 21.0666 21.4571 21.4571C21.0666 21.8476 20.4334 21.8476 20.0429 21.4571L16.0118 17.426C14.5645 18.5683 12.7369 19.25 10.75 19.25C6.05558 19.25 2.25 15.4444 2.25 10.75Z" fill="currentColor"></path>
                                                        </svg>
                                                    </button>
                                                </span>
                                                <span class="hidden"></span>
                                                <span class="flex" data-state="closed">
                                                    <button aria-label="New chat" data-testid="create-new-chat-button" class="h-10 rounded-lg px-2 text-token-text-secondary focus-visible:bg-token-surface-hover focus-visible:outline-0 enabled:hover:bg-token-surface-hover disabled:text-token-text-quaternary">
                                                        <svg width="24" height="24" viewBox="0 0 24 24" fill="currentColor" xmlns="http://www.w3.org/2000/svg" class="icon-xl-heavy">
                                                            <path d="M15.6729 3.91287C16.8918 2.69392 18.8682 2.69392 20.0871 3.91287C21.3061 5.13182 21.3061 7.10813 20.0871 8.32708L14.1499 14.2643C13.3849 15.0293 12.3925 15.5255 11.3215 15.6785L9.14142 15.9899C8.82983 16.0344 8.51546 15.9297 8.29289 15.7071C8.07033 15.4845 7.96554 15.1701 8.01005 14.8586L8.32149 12.6785C8.47449 11.6075 8.97072 10.615 9.7357 9.85006L15.6729 3.91287ZM18.6729 5.32708C18.235 4.88918 17.525 4.88918 17.0871 5.32708L11.1499 11.2643C10.6909 11.7233 10.3932 12.3187 10.3014 12.9613L10.1785 13.8215L11.0386 13.6986C11.6812 13.6068 12.2767 13.3091 12.7357 12.8501L18.6729 6.91287C19.1108 6.47497 19.1108 5.76499 18.6729 5.32708ZM11 3.99929C11.0004 4.55157 10.5531 4.99963 10.0008 5.00007C9.00227 5.00084 8.29769 5.00827 7.74651 5.06064C7.20685 5.11191 6.88488 5.20117 6.63803 5.32695C6.07354 5.61457 5.6146 6.07351 5.32698 6.63799C5.19279 6.90135 5.10062 7.24904 5.05118 7.8542C5.00078 8.47105 5 9.26336 5 10.4V13.6C5 14.7366 5.00078 15.5289 5.05118 16.1457C5.10062 16.7509 5.19279 17.0986 5.32698 17.3619C5.6146 17.9264 6.07354 18.3854 6.63803 18.673C6.90138 18.8072 7.24907 18.8993 7.85424 18.9488C8.47108 18.9992 9.26339 19 10.4 19H13.6C14.7366 19 15.5289 18.9992 16.1458 18.9488C16.7509 18.8993 17.0986 18.8072 17.362 18.673C17.9265 18.3854 18.3854 17.9264 18.673 17.3619C18.7988 17.1151 18.8881 16.7931 18.9393 16.2535C18.9917 15.7023 18.9991 14.9977 18.9999 13.9992C19.0003 13.4469 19.4484 12.9995 20.0007 13C20.553 13.0004 21.0003 13.4485 20.9999 14.0007C20.9991 14.9789 20.9932 15.7808 20.9304 16.4426C20.8664 17.116 20.7385 17.7136 20.455 18.2699C19.9757 19.2107 19.2108 19.9756 18.27 20.455C17.6777 20.7568 17.0375 20.8826 16.3086 20.9421C15.6008 21 14.7266 21 13.6428 21H10.3572C9.27339 21 8.39925 21 7.69138 20.9421C6.96253 20.8826 6.32234 20.7568 5.73005 20.455C4.78924 19.9756 4.02433 19.2107 3.54497 18.2699C3.24318 17.6776 3.11737 17.0374 3.05782 16.3086C2.99998 15.6007 2.99999 14.7266 3 13.6428V10.3572C2.99999 9.27337 2.99998 8.39922 3.05782 7.69134C3.11737 6.96249 3.24318 6.3223 3.54497 5.73001C4.02433 4.7892 4.78924 4.0243 5.73005 3.54493C6.28633 3.26149 6.88399 3.13358 7.55735 3.06961C8.21919 3.00673 9.02103 3.00083 9.99922 3.00007C10.5515 2.99964 10.9996 3.447 11 3.99929Z" fill="currentColor"></path>
                                                        </svg>
                                                    </button>
                                                </span>
                                            </div>
                                        </div>
                                        <div class="flex-col flex-1 transition-opacity duration-500 relative pr-3 overflow-y-auto">
                                            <div class="group/sidebar">
                                                <div class="bg-token-sidebar-surface-primary pt-0">
                                                    <span class="hidden"></span>
                                                    <span class="flex w-full items-center" data-state="closed">
                                                        <div class="flex-1" tabindex="0">
                                                            <a title="ChatGPT" style="--item-background-color:var(--sidebar-surface-primary)" class="no-draggable group rounded-lg active:opacity-90 bg-[var(--item-background-color)] h-9 text-sm flex items-center gap-2.5 p-2 screen-arch:relative screen-arch:bg-transparent screen-arch:py-[7px]" href="/" data-discover="true">
                                                                <div class="flex h-6 w-6 items-center justify-center text-token-text-secondary">
                                                                    <div class="h-6 w-6">
                                                                        <div class="gizmo-shadow-stroke relative flex h-full items-center justify-center rounded-full bg-token-main-surface-primary text-token-text-primary">
                                                                            <div class="flex h-full w-full items-center justify-center" style="opacity:1;will-change:opacity">
                                                                                <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="h-2/3 w-2/3">
                                                                                    <text x="-9999" y="-9999">ChatGPT</text>
                                                                                    <path d="M9.20509 8.76511V6.50545C9.20509 6.31513 9.27649 6.17234 9.44293 6.0773L13.9861 3.46088C14.6046 3.10413 15.342 2.93769 16.103 2.93769C18.9573 2.93769 20.7651 5.14983 20.7651 7.50454C20.7651 7.67098 20.7651 7.86129 20.7412 8.05161L16.0316 5.2924C15.7462 5.12596 15.4607 5.12596 15.1753 5.2924L9.20509 8.76511ZM19.8135 17.5659V12.1664C19.8135 11.8333 19.6708 11.5955 19.3854 11.429L13.4152 7.95633L15.3656 6.83833C15.5321 6.74328 15.6749 6.74328 15.8413 6.83833L20.3845 9.45474C21.6928 10.216 22.5728 11.8333 22.5728 13.4031C22.5728 15.2108 21.5025 16.8758 19.8135 17.5657V17.5659ZM7.80173 12.8088L5.8513 11.6671C5.68486 11.5721 5.61346 11.4293 5.61346 11.239V6.00613C5.61346 3.46111 7.56389 1.53433 10.2042 1.53433C11.2033 1.53433 12.1307 1.86743 12.9159 2.46202L8.2301 5.17371C7.94475 5.34015 7.80195 5.57798 7.80195 5.91109V12.809L7.80173 12.8088ZM12 15.2349L9.20509 13.6651V10.3351L12 8.76534L14.7947 10.3351V13.6651L12 15.2349ZM13.7958 22.4659C12.7967 22.4659 11.8693 22.1328 11.0841 21.5382L15.7699 18.8265C16.0553 18.6601 16.198 18.4222 16.198 18.0891V11.1912L18.1723 12.3329C18.3388 12.4279 18.4102 12.5707 18.4102 12.761V17.9939C18.4102 20.5389 16.4359 22.4657 13.7958 22.4657V22.4659ZM8.15848 17.1617L3.61528 14.5452C2.30696 13.784 1.42701 12.1667 1.42701 10.5969C1.42701 8.76534 2.52115 7.12414 4.20987 6.43428V11.8574C4.20987 12.1905 4.35266 12.4284 4.63802 12.5948L10.5846 16.0436L8.63415 17.1617C8.46771 17.2567 8.32492 17.2567 8.15848 17.1617ZM7.897 21.0625C5.20919 21.0625 3.23488 19.0407 3.23488 16.5432C3.23488 16.3529 3.25875 16.1626 3.2824 15.9723L7.96817 18.6839C8.25352 18.8504 8.53911 18.8504 8.82446 18.6839L14.7947 15.2351V17.4948C14.7947 17.6851 14.7233 17.8279 14.5568 17.9229L10.0136 20.5393C9.39518 20.8961 8.6578 21.0625 7.89677 21.0625H7.897ZM13.7958 23.8929C16.6739 23.8929 19.0762 21.8474 19.6235 19.1357C22.2874 18.4459 24 15.9484 24 13.4034C24 11.7383 23.2865 10.121 22.002 8.95542C22.121 8.45588 22.1924 7.95633 22.1924 7.45702C22.1924 4.0557 19.4331 1.51045 16.2458 1.51045C15.6037 1.51045 14.9852 1.60549 14.3668 1.81968C13.2963 0.773071 11.8215 0.107086 10.2042 0.107086C7.32606 0.107086 4.92383 2.15256 4.37653 4.86425C1.7126 5.55411 0 8.05161 0 10.5966C0 12.2617 0.713506 13.879 1.99795 15.0446C1.87904 15.5441 1.80764 16.0436 1.80764 16.543C1.80764 19.9443 4.56685 22.4895 7.75421 22.4895C8.39632 22.4895 9.01478 22.3945 9.63324 22.1803C10.7035 23.2269 12.1783 23.8929 13.7958 23.8929Z" fill="currentColor"></path>
                                                                                </svg>
                                                                            </div>
                                                                        </div>
                                                                    </div>
                                                                </div>
                                                                <div class="grow overflow-hidden text-ellipsis whitespace-nowrap text-sm text-token-text-primary">ChatGPT</div>
                                                            </a>
                                                        </div>
                                                    </span>
                                                </div>
                                                <div class="relative self-stretch">
                                                    <a title="Sora" style="--item-background-color:var(--sidebar-surface-primary)" class="no-draggable group rounded-lg active:opacity-90 bg-[var(--item-background-color)] h-9 text-sm flex items-center gap-2.5 p-2 screen-arch:relative screen-arch:bg-transparent screen-arch:py-[7px]" href="https://sora.com?utm_source=chatgpt" target="_blank">
                                                        <div class="flex h-6 w-6 items-center justify-center text-token-text-secondary">
                                                            <img src="/assets/sora-mutf8tav.webp" alt="Sora icon"/>
                                                        </div>
                                                        <div class="grow overflow-hidden text-ellipsis whitespace-nowrap text-sm text-token-text-primary">Sora</div>
                                                        <div class="invisible text-token-text-secondary hover:text-token-text-primary can-hover:group-hover:visible">
                                                            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="icon-md">
                                                                <g transform="rotate(45 12 12)">
                                                                    <path fill-rule="evenodd" clip-rule="evenodd" d="M12 3C12.2652 3 12.5196 3.10536 12.7071 3.29289L19.7071 10.2929C20.0976 10.6834 20.0976 11.3166 19.7071 11.7071C19.3166 12.0976 18.6834 12.0976 18.2929 11.7071L13 6.41421V20C13 20.5523 12.5523 21 12 21C11.4477 21 11 20.5523 11 20V6.41422L5.70711 11.7071C5.31658 12.0976 4.68342 12.0976 4.29289 11.7071C3.90237 11.3166 3.90237 10.6834 4.29289 10.2929L11.2929 3.29289C11.4804 3.10536 11.7348 3 12 3Z" fill="currentColor"></path>
                                                                </g>
                                                            </svg>
                                                        </div>
                                                    </a>
                                                </div>
                                                <div>
                                                    <div tabindex="0">
                                                        <a href="/gpts" data-discover="true">
                                                            <button data-testid="explore-gpts-button" class="flex h-9 w-full items-center gap-2.5 rounded-lg px-2 text-token-text-primary hover:bg-token-sidebar-surface-secondary">
                                                                <div class="flex h-6 w-6 items-center justify-center text-token-text-secondary">
                                                                    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="icon-md">
                                                                        <path fill-rule="evenodd" clip-rule="evenodd" d="M6.75 4.5C5.50736 4.5 4.5 5.50736 4.5 6.75C4.5 7.99264 5.50736 9 6.75 9C7.99264 9 9 7.99264 9 6.75C9 5.50736 7.99264 4.5 6.75 4.5ZM2.5 6.75C2.5 4.40279 4.40279 2.5 6.75 2.5C9.09721 2.5 11 4.40279 11 6.75C11 9.09721 9.09721 11 6.75 11C4.40279 11 2.5 9.09721 2.5 6.75Z" fill="currentColor"></path>
                                                                        <path fill-rule="evenodd" clip-rule="evenodd" d="M17.25 4.5C16.0074 4.5 15 5.50736 15 6.75C15 7.99264 16.0074 9 17.25 9C18.4926 9 19.5 7.99264 19.5 6.75C19.5 5.50736 18.4926 4.5 17.25 4.5ZM13 6.75C13 4.40279 14.9028 2.5 17.25 2.5C19.5972 2.5 21.5 4.40279 21.5 6.75C21.5 9.09721 19.5972 11 17.25 11C14.9028 11 13 9.09721 13 6.75Z" fill="currentColor"></path>
                                                                        <path fill-rule="evenodd" clip-rule="evenodd" d="M6.75 15C5.50736 15 4.5 16.0074 4.5 17.25C4.5 18.4926 5.50736 19.5 6.75 19.5C7.99264 19.5 9 18.4926 9 17.25C9 16.0074 7.99264 15 6.75 15ZM2.5 17.25C2.5 14.9028 4.40279 13 6.75 13C9.09721 13 11 14.9028 11 17.25C11 19.5972 9.09721 21.5 6.75 21.5C4.40279 21.5 2.5 19.5972 2.5 17.25Z" fill="currentColor"></path>
                                                                        <path fill-rule="evenodd" clip-rule="evenodd" d="M17.25 15C16.0074 15 15 16.0074 15 17.25C15 18.4926 16.0074 19.5 17.25 19.5C18.4926 19.5 19.5 18.4926 19.5 17.25C19.5 16.0074 18.4926 15 17.25 15ZM13 17.25C13 14.9028 14.9028 13 17.25 13C19.5972 13 21.5 14.9028 21.5 17.25C21.5 19.5972 19.5972 21.5 17.25 21.5C14.9028 21.5 13 19.5972 13 17.25Z" fill="currentColor"></path>
                                                                    </svg>
                                                                </div>
                                                                <span class="text-sm">Explore GPTs</span>
                                                            </button>
                                                        </a>
                                                    </div>
                                                </div>
                                                <div class="z-20 screen-arch:sticky screen-arch:top-[var(--sticky-title-offset)] select-none overflow-clip text-ellipsis break-all pl-2 pt-7 text-xs font-semibold text-token-text-primary screen-arch:-mr-2 screen-arch:h-10 screen-arch:min-w-[50cqw] screen-arch:-translate-x-2 screen-arch:bg-[var(--sidebar-surface)] screen-arch:py-1 screen-arch:text-token-text-secondary">
                                                    <h2 id="snorlax-heading" class="flex h-[26px] w-full items-center justify-between text-xs"></h2>
                                                </div>
                                                <aside class="flex flex-col gap-4 mb-0">
                                                    <ul aria-labelledby="snorlax-heading" class="flex flex-col screen-arch:mb-3"></ul>
                                                </aside>
                                                <div class="flex flex-col gap-2 text-token-text-primary text-sm mt-5 first:mt-0 h-full justify-center items-center empty:hidden"></div>
                                            </div>
                                        </div>
                                    </nav>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="relative flex h-full max-w-full flex-1 flex-col overflow-hidden">
                        <main class="relative h-full w-full flex-1 overflow-auto transition-width">
                            <div class="h-full w-full @container/thread">
                                <div role="presentation" class="composer-parent flex h-full flex-col focus-visible:outline-0">
                                    <div class="draggable no-draggable-children top-0 p-3 flex items-center justify-between z-10 h-header-height font-semibold bg-token-main-surface-primary pointer-events-none select-none *:pointer-events-auto motion-safe:transition max-md:hidden absolute left-0 right-0 @thread-xl/thread:absolute @thread-xl/thread:left-0 @thread-xl/thread:right-0 @thread-xl/thread:bg-transparent @thread-xl/thread:!shadow-none [box-shadow:var(--sharp-edge-top-shadow-placeholder)]">
                                        <div class="absolute start-1/2 ltr:-translate-x-1/2 rtl:translate-x-1/2"></div>
                                        <div class="flex items-center gap-0 overflow-hidden">
                                            <button aria-label="" style="view-transition-name:var(--vt-thread-model-switcher)" class="group flex cursor-pointer items-center gap-1 rounded-lg py-1.5 px-3 text-lg hover:bg-token-main-surface-secondary radix-state-open:bg-token-main-surface-secondary font-semibold text-token-text-secondary overflow-hidden whitespace-nowrap">
                                                <div class="text-token-text-secondary">ChatGPT </div>
                                                <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="icon-md text-token-text-tertiary">
                                                    <path fill-rule="evenodd" clip-rule="evenodd" d="M5.29289 9.29289C5.68342 8.90237 6.31658 8.90237 6.70711 9.29289L12 14.5858L17.2929 9.29289C17.6834 8.90237 18.3166 8.90237 18.7071 9.29289C19.0976 9.68342 19.0976 10.3166 18.7071 10.7071L12.7071 16.7071C12.5196 16.8946 12.2652 17 12 17C11.7348 17 11.4804 16.8946 11.2929 16.7071L5.29289 10.7071C4.90237 10.3166 4.90237 9.68342 5.29289 9.29289Z" fill="currentColor"></path>
                                                </svg>
                                            </button>
                                        </div>
                                        <div class="flex items-center gap-2 pr-1 leading-[0]">
                                            <span class="hidden"></span>
                                            <span class="" data-state="closed">
                                                <button class="btn relative btn-secondary btn btn-secondary relative text-token-text-primary" aria-label="Temporary">
                                                    <div class="flex w-full items-center justify-center gap-1.5">
                                                        <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" fill="none" viewBox="0 0 14 14" aria-label="" class="icon-sm">
                                                            <path fill="currentColor" fill-rule="evenodd" d="M6.319 1.334a.667.667 0 0 1-.512.792 5.43 5.43 0 0 0-2.602 1.362.667.667 0 1 1-.918-.967A6.76 6.76 0 0 1 5.527.822a.667.667 0 0 1 .792.512m1.363 0a.667.667 0 0 1 .791-.512 6.76 6.76 0 0 1 3.24 1.699.667.667 0 1 1-.917.967 5.43 5.43 0 0 0-2.602-1.362.667.667 0 0 1-.512-.792M1.51 4.614c.348.12.533.5.413.848a4.7 4.7 0 0 0 0 3.076.667.667 0 0 1-1.26.435 6.04 6.04 0 0 1 0-3.945.666.666 0 0 1 .847-.413m10.979 0a.667.667 0 0 1 .847.414A6 6 0 0 1 13.667 7a6 6 0 0 1-.33 1.973.667.667 0 1 1-1.26-.435 4.7 4.7 0 0 0 0-3.076.667.667 0 0 1 .413-.847M2.27 10.352a.667.667 0 0 1 .479.812q-.052.2-.111.397.629-.097 1.228-.267a.67.67 0 0 1 .496.054c.445.238.93.417 1.445.528a.667.667 0 1 1-.28 1.303 7 7 0 0 1-1.553-.533c-.73.189-1.479.305-2.266.354a.667.667 0 0 1-.664-.905c.164-.425.305-.844.414-1.264a.667.667 0 0 1 .812-.48m9.468.186a.666.666 0 0 1-.024.942 6.76 6.76 0 0 1-3.24 1.7.667.667 0 0 1-.28-1.304 5.43 5.43 0 0 0 2.601-1.362.667.667 0 0 1 .943.024" clip-rule="evenodd"></path>
                                                        </svg>
                                                        Temporary
                                                    </div>
                                                </button>
                                            </span>
                                            <div class="flex h-10 w-10 items-center justify-center">
                                                <div class="h-8 w-8 rounded-full bg-token-main-surface-tertiary"></div>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="flex flex-1 grow basis-auto flex-col overflow-hidden">
                                        <div class="text-base my-auto mx-auto px-3 md:px-4 w-full md:px-5 lg:px-4 xl:px-5 h-full">
                                            <div class="mx-auto mt-0 flex h-full w-full flex-col text-base @lg/thread:justify-center @md/thread:max-w-3xl @lg/thread:max-w-[40rem] @xl/thread:max-w-[48rem] relative">
                                                <div class="hidden text-center @lg/thread:block mb-5 @lg/thread:block">
                                                    <div class="relative inline-flex justify-center text-center text-2xl font-semibold leading-9">
                                                        <div style="view-transition-name:var(--vt-splash-screen-headline)">
                                                            <div class="grid min-h-[74px] grid-cols-1 items-center justify-end">
                                                                <h1 class="flex h-full items-end justify-center text-balance [grid-area:1/1] motion-safe:[transition:0.3s_transform_var(--spring-standard),0.5s_opacity_var(--spring-standard),0.5s_visibility_var(--spring-standard)] text-[28px] font-semibold leading-[34px] tracking-[0.38px] translate-y-0 opacity-100" aria-hidden="false">What can I help with?</h1>
                                                                <div class="flex h-full flex-col items-center justify-end [grid-area:1/1] motion-safe:[transition:0.3s_transform_var(--spring-standard),0.5s_opacity_var(--spring-standard),0.5s_visibility_var(--spring-standard)] invisible translate-y-full opacity-0" aria-hidden="true">
                                                                    <h1 class="mb-1 text-[28px] font-semibold leading-[34px] tracking-[0.38px]" data-testid="temporary-chat-label">
                                                                        <input class="sr-only" id="temporary-chat-checkbox" type="checkbox" readOnly="" name="temporary-chat-checkbox" checked=""/>Temporary Chat
                                                                    </h1>
                                                                    <div class="text-center font-normal max-w-xl text-[15px] leading-[18px] tracking-[-0.23px] text-token-text-secondary">This chat won &#x27;t appear in history or be used to train our models. For safety purposes, we may keep a copy of this chat for up to 30 days.</div>
                                                                </div>
                                                            </div>
                                                        </div>
                                                    </div>
                                                </div>
                                                <div class="mt-[var(--screen-optical-compact-offset-amount)] flex-shrink flex-col items-center justify-center overflow-hidden text-token-text-primary [display:var(--display-hidden-until-loaded,flex)] h-full @lg/thread:hidden" style="opacity:0;will-change:opacity">
                                                    <div class="relative inline-flex justify-center text-center text-2xl font-semibold leading-9">
                                                        <div style="view-transition-name:var(--vt-splash-screen-headline)">
                                                            <div class="grid min-h-[74px] grid-cols-1 items-center justify-end">
                                                                <h1 class="flex h-full items-end justify-center text-balance [grid-area:1/1] motion-safe:[transition:0.3s_transform_var(--spring-standard),0.5s_opacity_var(--spring-standard),0.5s_visibility_var(--spring-standard)] text-[28px] font-semibold leading-[34px] tracking-[0.38px] translate-y-0 opacity-100" aria-hidden="false">What can I help with?</h1>
                                                                <div class="flex h-full flex-col items-center justify-end [grid-area:1/1] motion-safe:[transition:0.3s_transform_var(--spring-standard),0.5s_opacity_var(--spring-standard),0.5s_visibility_var(--spring-standard)] invisible translate-y-full opacity-0" aria-hidden="true">
                                                                    <h1 class="mb-1 text-[28px] font-semibold leading-[34px] tracking-[0.38px]" data-testid="temporary-chat-label">
                                                                        <input class="sr-only" id="temporary-chat-checkbox" type="checkbox" readOnly="" name="temporary-chat-checkbox" checked=""/>Temporary Chat
                                                                    </h1>
                                                                    <div class="text-center font-normal max-w-xl text-[15px] leading-[18px] tracking-[-0.23px] text-token-text-secondary">This chat won &#x27;t appear in history or be used to train our models. For safety purposes, we may keep a copy of this chat for up to 30 days.</div>
                                                                </div>
                                                            </div>
                                                        </div>
                                                    </div>
                                                    <div class="h-[116px]" style="opacity:0;will-change:opacity"></div>
                                                </div>
                                                <div class="@lg/thread:absolute @lg/thread:bottom-8 @lg/thread:left-0 @lg/thread:w-full">
                                                    <div class="text-base mx-auto px-3 md:px-4 w-full md:px-5 lg:px-4 xl:px-5">
                                                        <div class="mx-auto mt-0 flex h-full w-full flex-col text-base @lg/thread:justify-center @md/thread:max-w-3xl @lg/thread:max-w-[40rem] @xl/thread:max-w-[48rem] relative @lg/thread:pb-0">
                                                            <div class="block z-20"></div>
                                                        </div>
                                                    </div>
                                                </div>
                                                <div class="w-full">
                                                    <div class="flex justify-center empty:hidden"></div>
                                                    <div class="relative z-[1] flex max-w-full flex-1 flex-col h-full" type="button" aria-haspopup="dialog" aria-expanded="false" aria-controls="radix-:R3iv5cfaadkklj5:" data-state="closed">
                                                        <form class="w-full" data-type="unified-composer">
                                                            <div class="flex w-full cursor-text flex-col items-center justify-center rounded-[28px] border border-token-border-xlight contain-inline-size overflow-clip shadow-sm dark:!shadow-none sm:shadow-lg bg-token-main-surface-primary dark:bg-[#303030]">
                                                                <div class="relative flex w-full items-end py-3 pl-3">
                                                                    <div class="relative flex w-full flex-auto flex-col">
                                                                        <div class="relative ml-1.5 grid grid-cols-[auto_minmax(0,1fr)]">
                                                                            <div class="items-top flex justify-center">
                                                                                <div style="opacity:1;will-change:opacity"></div>
                                                                            </div>
                                                                            <div style="margin-bottom:-20px;will-change:transform;transform:translateY(-7px)" class="relative flex-auto bg-transparent pl-2 pt-0.5">
                                                                                <div class="flex flex-col justify-start" style="min-height:0">
                                                                                    <div class="_prosemirror-parent_11fu7_1 text-token-text-primary max-h-[25dvh] max-h-52 overflow-auto [scrollbar-width:thin] default-browser min-h-12 pr-3">
                                                                                        <textarea class="block h-10 w-full resize-none border-0 bg-transparent px-0 py-2 text-token-text-primary placeholder:pl-px placeholder:text-token-text-tertiary" autofocus="" placeholder="Ask anything" data-virtualkeyboard="true"></textarea>
                                                                                        <script nonce="20f285f6-3f39-465d-8802-567f078c6c55">
                                                                                            window.__oai_logHTML ? window.__oai_logHTML() : window.__oai_SSR_HTML = window.__oai_SSR_HTML || Date.now();
                                                                                            requestAnimationFrame((function() {
                                                                                                window.__oai_logTTI ? window.__oai_logTTI() : window.__oai_SSR_TTI = window.__oai_SSR_TTI || Date.now()
                                                                                            }
                                                                                            ))
                                                                                        </script>
                                                                                    </div>
                                                                                </div>
                                                                            </div>
                                                                        </div>
                                                                        <div class="justify-content-end relative ml-2 flex w-full flex-auto flex-col">
                                                                            <div class="flex-auto"></div>
                                                                        </div>
                                                                        <div style="height:48px"></div>
                                                                    </div>
                                                                    <div class="bg-primary-surface-primary absolute bottom-[9px] left-[17px] right-0 z-[2] flex items-center">
                                                                        <div>
                                                                            <div class="flex items-center gap-2 max-xs:gap-1">
                                                                                <div style="view-transition-name:var(--vt-composer-attach-file-action)">
                                                                                    <div class="relative">
                                                                                        <div class="relative">
                                                                                            <div class="flex flex-col">
                                                                                                <input multiple="" type="file" style="display:none" tabindex="-1" class="hidden"/>
                                                                                                <span class="hidden"></span>
                                                                                                <button type="button" id="radix-:R1kmcpjiv5cfaadkklj5:" aria-haspopup="menu" aria-expanded="false" data-state="closed" class="text-token-text-primary border border-transparent inline-flex items-center justify-center gap-1 rounded-lg text-sm dark:transparent dark:bg-transparent leading-none outline-none cursor-pointer hover:bg-token-main-surface-secondary dark:hover:bg-token-main-surface-secondary focus-visible:bg-token-main-surface-secondary radix-state-active:text-token-text-secondary radix-disabled:cursor-auto radix-disabled:bg-transparent radix-disabled:text-token-text-tertiary dark:radix-disabled:bg-transparent m-0 h-0 w-0 border-none bg-transparent p-0"></button>
                                                                                                <span class="hidden"></span>
                                                                                                <span class="flex" data-state="closed">
                                                                                                    <button disabled="" aria-disabled="true" aria-label="Attach files is unavailable" class="flex items-center justify-center h-9 rounded-full border border-token-border-light text-token-text-secondary w-9 opacity-50 min-h-8 min-w-[34px]">
                                                                                                        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" aria-label="" class="h-[18px] w-[18px]">
                                                                                                            <path fill-rule="evenodd" clip-rule="evenodd" d="M12 3C12.5523 3 13 3.44772 13 4L13 11H20C20.5523 11 21 11.4477 21 12C21 12.5523 20.5523 13 20 13L13 13L13 20C13 20.5523 12.5523 21 12 21C11.4477 21 11 20.5523 11 20L11 13L4 13C3.44772 13 3 12.5523 3 12C3 11.4477 3.44772 11 4 11L11 11L11 4C11 3.44772 11.4477 3 12 3Z" fill="currentColor"></path>
                                                                                                        </svg>
                                                                                                    </button>
                                                                                                </span>
                                                                                                <div class="w-fit" type="button" aria-haspopup="dialog" aria-expanded="false" aria-controls="radix-:R24mcpjiv5cfaadkklj5:" data-state="closed">
                                                                                                    <div></div>
                                                                                                </div>
                                                                                            </div>
                                                                                        </div>
                                                                                    </div>
                                                                                </div>
                                                                                <div style="view-transition-name:var(--vt-composer-search-action)">
                                                                                    <div>
                                                                                        <span class="hidden"></span>
                                                                                        <span class="" data-state="closed">
                                                                                            <div class="inline-flex h-9 rounded-full border text-[13px] font-medium radix-state-open:bg-black/10 text-token-text-secondary border-token-border-light focus-visible:outline-black can-hover:hover:bg-token-main-surface-secondary dark:focus-visible:outline-white dark:can-hover:hover:bg-gray-700 opacity-30">
                                                                                                <button class="flex h-full min-w-8 items-center justify-center p-2" aria-pressed="false" aria-label="Search" disabled="">
                                                                                                    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="h-[18px] w-[18px]">
                                                                                                        <path fill-rule="evenodd" clip-rule="evenodd" d="M2 12C2 6.47715 6.47715 2 12 2C17.5228 2 22 6.47715 22 12C22 17.5228 17.5228 22 12 22C6.47715 22 2 17.5228 2 12ZM11.9851 4.00291C11.9933 4.00046 11.9982 4.00006 11.9996 4C12.001 4.00006 12.0067 4.00046 12.0149 4.00291C12.0256 4.00615 12.047 4.01416 12.079 4.03356C12.2092 4.11248 12.4258 4.32444 12.675 4.77696C12.9161 5.21453 13.1479 5.8046 13.3486 6.53263C13.6852 7.75315 13.9156 9.29169 13.981 11H10.019C10.0844 9.29169 10.3148 7.75315 10.6514 6.53263C10.8521 5.8046 11.0839 5.21453 11.325 4.77696C11.5742 4.32444 11.7908 4.11248 11.921 4.03356C11.953 4.01416 11.9744 4.00615 11.9851 4.00291ZM8.01766 11C8.08396 9.13314 8.33431 7.41167 8.72334 6.00094C8.87366 5.45584 9.04762 4.94639 9.24523 4.48694C6.48462 5.49946 4.43722 7.9901 4.06189 11H8.01766ZM4.06189 13H8.01766C8.09487 15.1737 8.42177 17.1555 8.93 18.6802C9.02641 18.9694 9.13134 19.2483 9.24522 19.5131C6.48461 18.5005 4.43722 16.0099 4.06189 13ZM10.019 13H13.981C13.9045 14.9972 13.6027 16.7574 13.1726 18.0477C12.9206 18.8038 12.6425 19.3436 12.3823 19.6737C12.2545 19.8359 12.1506 19.9225 12.0814 19.9649C12.0485 19.9852 12.0264 19.9935 12.0153 19.9969C12.0049 20.0001 11.9999 20 11.9999 20C11.9999 20 11.9948 20 11.9847 19.9969C11.9736 19.9935 11.9515 19.9852 11.9186 19.9649C11.8494 19.9225 11.7455 19.8359 11.6177 19.6737C11.3575 19.3436 11.0794 18.8038 10.8274 18.0477C10.3973 16.7574 10.0955 14.9972 10.019 13ZM15.9823 13C15.9051 15.1737 15.5782 17.1555 15.07 18.6802C14.9736 18.9694 14.8687 19.2483 14.7548 19.5131C17.5154 18.5005 19.5628 16.0099 19.9381 13H15.9823ZM19.9381 11C19.5628 7.99009 17.5154 5.49946 14.7548 4.48694C14.9524 4.94639 15.1263 5.45584 15.2767 6.00094C15.6657 7.41167 15.916 9.13314 15.9823 11H19.9381Z" fill="currentColor"></path>
                                                                                                    </svg>
                                                                                                    <span style="width:fit-content;opacity:1;will-change:transform,opacity;transform:none">
                                                                                                        <div class="whitespace-nowrap pl-1 pr-1 [display:--force-hide-label]">Search</div>
                                                                                                    </span>
                                                                                                </button>
                                                                                            </div>
                                                                                        </span>
                                                                                    </div>
                                                                                </div>
                                                                                <div style="view-transition-name:var(--vt-composer-research-action)">
                                                                                    <div>
                                                                                        <span class="hidden"></span>
                                                                                        <span class="" data-state="closed">
                                                                                            <div class="inline-flex h-9 rounded-full border text-[13px] font-medium radix-state-open:bg-black/10 text-token-text-secondary border-token-border-light focus-visible:outline-black can-hover:hover:bg-token-main-surface-secondary dark:focus-visible:outline-white dark:can-hover:hover:bg-gray-700">
                                                                                                <button class="flex h-full min-w-8 items-center justify-center p-2" aria-pressed="false" aria-label="Deep research">
                                                                                                    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="h-[18px] w-[18px]">
                                                                                                        <path fill-rule="evenodd" clip-rule="evenodd" d="M12.47 15.652a1 1 0 0 1 1.378.318l2.5 4a1 1 0 1 1-1.696 1.06l-2.5-4a1 1 0 0 1 .318-1.378Z" fill="currentColor"></path>
                                                                                                        <path fill-rule="evenodd" clip-rule="evenodd" d="M11.53 15.652a1 1 0 0 1 .318 1.378l-2.5 4a1 1 0 0 1-1.696-1.06l2.5-4a1 1 0 0 1 1.378-.318ZM17.824 4.346a.5.5 0 0 0-.63-.321l-.951.309a1 1 0 0 0-.642 1.26l1.545 4.755a1 1 0 0 0 1.26.642l.95-.309a.5.5 0 0 0 .322-.63l-1.854-5.706Zm-1.248-2.223a2.5 2.5 0 0 1 3.15 1.605l1.854 5.706a2.5 2.5 0 0 1-1.605 3.15l-.951.31a2.992 2.992 0 0 1-2.443-.265l-2.02.569a1 1 0 1 1-.541-1.926l1.212-.34-1.353-4.163L5 10.46a1 1 0 0 0-.567 1.233l.381 1.171a1 1 0 0 0 1.222.654l3.127-.88a1 1 0 1 1 .541 1.926l-3.127.88a3 3 0 0 1-3.665-1.961l-.38-1.172a3 3 0 0 1 1.7-3.697l9.374-3.897a3 3 0 0 1 2.02-2.285l.95-.31Z" fill="currentColor"></path>
                                                                                                        <path fill-rule="evenodd" clip-rule="evenodd" d="M12 12.5a1.5 1.5 0 1 0 0 3 1.5 1.5 0 0 0 0-3ZM8.5 14a3.5 3.5 0 1 1 7 0 3.5 3.5 0 0 1-7 0Z" fill="currentColor"></path>
                                                                                                    </svg>
                                                                                                    <div class="whitespace-nowrap pl-1 pr-1 [display:--force-hide-label]">Deep research</div>
                                                                                                </button>
                                                                                            </div>
                                                                                        </span>
                                                                                    </div>
                                                                                </div>
                                                                                <div style="view-transition-name:var(--vt-composer-system-hint-action)">
                                                                                    <span class="hidden"></span>
                                                                                    <span class="hidden"></span>
                                                                                    <span class="" data-state="closed">
                                                                                        <button type="button" id="radix-:R36cpjiv5cfaadkklj5:" aria-haspopup="menu" aria-expanded="false" data-state="closed" class="_toolsButton_d2h2h_8 flex h-9 min-w-9 items-center justify-center rounded-full border border-token-border-light p-1 text-xs font-semibold text-token-text-secondary focus-visible:outline-black disabled:opacity-30 radix-state-open:bg-black/10 can-hover:hover:bg-token-main-surface-secondary dark:focus-visible:outline-white dark:can-hover:hover:bg-gray-700" aria-label="Use a tool">
                                                                                            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="pointer-events-none h-5 w-5">
                                                                                                <path fill-rule="evenodd" clip-rule="evenodd" d="M3 12C3 10.8954 3.89543 10 5 10C6.10457 10 7 10.8954 7 12C7 13.1046 6.10457 14 5 14C3.89543 14 3 13.1046 3 12ZM10 12C10 10.8954 10.8954 10 12 10C13.1046 10 14 10.8954 14 12C14 13.1046 13.1046 14 12 14C10.8954 14 10 13.1046 10 12ZM17 12C17 10.8954 17.8954 10 19 10C20.1046 10 21 10.8954 21 12C21 13.1046 20.1046 14 19 14C17.8954 14 17 13.1046 17 12Z" fill="currentColor"></path>
                                                                                            </svg>
                                                                                        </button>
                                                                                    </span>
                                                                                </div>
                                                                            </div>
                                                                            <div class="absolute bottom-1 right-3 flex items-center gap-2">
                                                                                <div class="ml-auto">
                                                                                    <div class="min-w-7">
                                                                                        <button class="relative flex h-9 items-center justify-center rounded-full bg-black text-white transition-colors focus-visible:outline-none focus-visible:outline-black disabled:text-gray-50 disabled:opacity-30 can-hover:hover:opacity-70 dark:bg-white dark:text-black w-7 !h-7">
                                                                                            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                                                                                <path d="M9.5 4C8.67157 4 8 4.67157 8 5.5V18.5C8 19.3284 8.67157 20 9.5 20C10.3284 20 11 19.3284 11 18.5V5.5C11 4.67157 10.3284 4 9.5 4Z" fill="currentColor"></path>
                                                                                                <path d="M13 8.5C13 7.67157 13.6716 7 14.5 7C15.3284 7 16 7.67157 16 8.5V15.5C16 16.3284 15.3284 17 14.5 17C13.6716 17 13 16.3284 13 15.5V8.5Z" fill="currentColor"></path>
                                                                                                <path d="M4.5 9C3.67157 9 3 9.67157 3 10.5V13.5C3 14.3284 3.67157 15 4.5 15C5.32843 15 6 14.3284 6 13.5V10.5C6 9.67157 5.32843 9 4.5 9Z" fill="currentColor"></path>
                                                                                                <path d="M19.5 9C18.6716 9 18 9.67157 18 10.5V13.5C18 14.3284 18.6716 15 19.5 15C20.3284 15 21 14.3284 21 13.5V10.5C21 9.67157 20.3284 9 19.5 9Z" fill="currentColor"></path>
                                                                                            </svg>
                                                                                        </button>
                                                                                    </div>
                                                                                </div>
                                                                            </div>
                                                                        </div>
                                                                    </div>
                                                                    <div class="absolute left-4 top-3 ml-[1px] flex items-center pb-px"></div>
                                                                </div>
                                                            </div>
                                                            <div class="w-full"></div>
                                                        </form>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="isolate w-full basis-auto has-[[data-has-thread-error]]:pt-2 has-[[data-has-thread-error]]:[box-shadow:var(--sharp-edge-bottom-shadow)] dark:border-white/20 md:border-transparent md:pt-0 md:dark:border-transparent">
                                        <div class="relative mt-auto flex min-h-8 w-full items-center justify-center p-2 text-center text-xs text-token-text-secondary md:px-[60px]">
                                            <div>ChatGPT can make mistakes. Check important info.</div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <div class="group absolute bottom-2 end-2 z-20 flex flex-col gap-1 md:flex lg:bottom-3 lg:end-3"></div>
                        </main>
                    </div>
                </div>
            </div>
        </div>
        <div aria-live="assertive" aria-atomic="true" class="sr-only"></div>
        <div aria-live="polite" aria-atomic="true" class="sr-only"></div>
        <audio class="fixed bottom-0 left-0 hidden h-0 w-0" autoPlay="" crossorigin="anonymous"></audio>
        <script nonce="20f285f6-3f39-465d-8802-567f078c6c55">
            window.__reactRouterContext = {
                "basename": "/",
                "future": {
                    "unstable_optimizeDeps": false
                },
                "isSpaMode": false
            };
            window.__reactRouterContext.stream = new ReadableStream({
                start(controller) {
                    window.__reactRouterContext.streamController = controller;
                }
            }).pipeThrough(new TextEncoderStream());
        </script>
        <script nonce="20f285f6-3f39-465d-8802-567f078c6c55" type="module" async="">
            import "/assets/manifest-532b61a2.js";
            import*as route0 from "/assets/jandtmjen68hvio9.js";
            import*as route1 from "/assets/lprjl238d50elo8f.js";
            import*as route2 from "/assets/ovc3ksv6dtkjao8i.js";

            window.__reactRouterRouteModules = {
                "root": route0,
                "routes/_conversation": route1,
                "routes/_conversation._index": route2
            };

            import("/assets/mwrrfx5gxel0ghh2.js");
        </script>
        <!--$?-->
        <template id="B:0"></template>
        <!--/$-->
        <div hidden id="S:0">
            <script nonce="20f285f6-3f39-465d-8802-567f078c6c55">
                window.__reactRouterContext.streamController.enqueue("{{ react_chatgpt_context_1|safe }}\n");
            </script>
            <!--$?-->
            <template id="B:1"></template>
            <!--/$-->
        </div>
        <script nonce="20f285f6-3f39-465d-8802-567f078c6c55">
            $RC = function(b, c, e) {
                c = document.getElementById(c);
                c.parentNode.removeChild(c);
                var a = document.getElementById(b);
                if (a) {
                    b = a.previousSibling;
                    if (e)
                        b.data = "$!",
                        a.setAttribute("data-dgst", e);
                    else {
                        e = b.parentNode;
                        a = b.nextSibling;
                        var f = 0;
                        do {
                            if (a && 8 === a.nodeType) {
                                var d = a.data;
                                if ("/$" === d)
                                    if (0 === f)
                                        break;
                                    else
                                        f--;
                                else
                                    "$" !== d && "$?" !== d && "$!" !== d || f++
                            }
                            d = a.nextSibling;
                            e.removeChild(a);
                            a = d
                        } while (a);
                        for (; c.firstChild; )
                            e.insertBefore(c.firstChild, a);
                        b.data = "$"
                    }
                    b._reactRetry && b._reactRetry()
                }
            }
            ;
            $RC("B:0", "S:0")
        </script>
        <div hidden id="S:1">
            <script nonce="20f285f6-3f39-465d-8802-567f078c6c55">
                window.__reactRouterContext.streamController.enqueue("[]\n");
            </script>
            <!--$?-->
            <template id="B:2"></template>
            <!--/$-->
        </div>
        <script nonce="20f285f6-3f39-465d-8802-567f078c6c55">
            $RC("B:1", "S:1")
        </script>
        <div hidden id="S:2">
            <script nonce="20f285f6-3f39-465d-8802-567f078c6c55">
                window.__reactRouterContext.streamController.enqueue("[]\n");
            </script>
            <!--$?-->
            <template id="B:3"></template>
            <!--/$-->
        </div>
        <script nonce="20f285f6-3f39-465d-8802-567f078c6c55">
            $RC("B:2", "S:2")
        </script>
        <div hidden id="S:3">
            <script nonce="20f285f6-3f39-465d-8802-567f078c6c55">
                window.__reactRouterContext.streamController.close();
            </script>
        </div>
        <script nonce="20f285f6-3f39-465d-8802-567f078c6c55">
            $RC("B:3", "S:3")
        </script>
    </body>
</html>
````

## File: templates/gpts_context.json
````json
[
    {
        "_1": 2
    },
    "routes/g.$gizmoId._index",
    {
        "_3": 4
    },
    "data",
    {
        "_5": 6,
        "_7": 8,
        "_9": 10,
        "_48": 49,
        "_50": 51,
        "_52": 53,
        "_1832": 61,
        "_1833": -5,
        "_1834": 855,
        "_1835": 1829,
        "_1836": 1837,
        "_1838": 1839,
        "_1840": 1841,
        "_1842": 24,
        "_1843": 24,
        "_1844": 24,
        "_1845": 24,
        "_1846": 1847,
        "_1848": 1849
    },
    "authStatus",
    "logged_in",
    "session",
    {
        "_9": 10,
        "_29": 30,
        "_31": 32,
        "_38": 39,
        "_40": 41,
        "_42": 43
    },
    "user",
    {
        "_11": 12,
        "_13": 14,
        "_15": 14,
        "_16": 17,
        "_18": 17,
        "_19": 20,
        "_21": 22,
        "_23": 24,
        "_25": 26,
        "_27": 28
    },
    "id",
    "user-chatgpt",
    "name",
    "chatgpt",
    "email",
    "chatgpt@openai.com",
    "https://s.gravatar.com/avatar/5edae94250bff28c50456d715798421e?s=480&r=pg&d=https%3A%2F%2Fcdn.auth0.com%2Favatars%2Fli.png",
    "picture",
    "idp",
    "auth0",
    "iat",
    1742986864,
    "mfa",
    false,
    "groups",
    [],
    "intercom_hash",
    "2bfe17418ed9bd0db1924962d30f2ed4adcede4bc9a81ac41ce9d623ab3b4de5",
    "expires",
    "2025-06-24T11:02:54.140Z",
    "account",
    {
        "_11": 33,
        "_34": 35,
        "_36": 37
    },
    "17904ad0-df88-4a0f-917b-501d958eceda",
    "planType",
    "plus",
    "structure",
    "personal",
    "accessToken",
    "",
    "authProvider",
    "openai",
    "rumViewTags",
    {
        "_44": 45
    },
    "light_account",
    {
        "_46": 24,
        "_47": -7
    },
    "fetched",
    "reason",
    "cluster",
    "unified-9",
    "locale",
    "zh-CN",
    "statsig",
    {
        "_54": 55,
        "_865": 866
    },
    "classic",
    {
        "_56": 57,
        "_350": 351,
        "_446": 447,
        "_827": 828,
        "_829": 61,
        "_830": 831,
        "_832": 833,
        "_838": 839,
        "_840": 841,
        "_851": 852,
        "_9": 853
    },
    "feature_gates",
    {
        "_58": 59,
        "_76": 77,
        "_82": 83,
        "_89": 90,
        "_103": 104,
        "_107": 108,
        "_111": 112,
        "_115": 116,
        "_118": 119,
        "_122": 123,
        "_132": 133,
        "_136": 137,
        "_140": 141,
        "_143": 144,
        "_147": 148,
        "_151": 152,
        "_154": 155,
        "_158": 159,
        "_162": 163,
        "_172": 173,
        "_175": 176,
        "_179": 180,
        "_186": 187,
        "_189": 190,
        "_87": 194,
        "_196": 197,
        "_206": 207,
        "_212": 213,
        "_216": 217,
        "_222": 223,
        "_227": 228,
        "_184": 230,
        "_232": 233,
        "_235": 236,
        "_239": 240,
        "_245": 246,
        "_248": 249,
        "_252": 253,
        "_256": 257,
        "_260": 261,
        "_264": 265,
        "_268": 269,
        "_271": 272,
        "_275": 276,
        "_280": 281,
        "_284": 285,
        "_287": 288,
        "_293": 294,
        "_300": 301,
        "_306": 307,
        "_201": 309,
        "_311": 312,
        "_315": 316,
        "_319": 320,
        "_94": 322,
        "_324": 325,
        "_335": 336,
        "_333": 339,
        "_343": 344,
        "_346": 347
    },
    "61299031",
    {
        "_13": 58,
        "_60": 61,
        "_62": 63,
        "_64": 65
    },
    "value",
    true,
    "rule_id",
    "2wrvcqZBGOdzYtk4c8rQxP",
    "secondary_exposures",
    [
        66,
        73
    ],
    {
        "_67": 68,
        "_69": 70,
        "_71": 72
    },
    "gate",
    "44045625",
    "gateValue",
    "true",
    "ruleID",
    "1vGfaAvyQ4VnZ5Y0UnCsbl:100.00:5",
    {
        "_67": 74,
        "_69": 70,
        "_71": 75
    },
    "1259585210",
    "3cQqufsn9EF8iqIPZFNiE8:100.00:4",
    "80186230",
    {
        "_13": 76,
        "_60": 61,
        "_62": 78,
        "_64": 79
    },
    "7thMqF7L1NKFaEvg1NsH7E",
    [
        80,
        81
    ],
    {
        "_67": 68,
        "_69": 70,
        "_71": 72
    },
    {
        "_67": 74,
        "_69": 70,
        "_71": 75
    },
    "174366048",
    {
        "_13": 82,
        "_60": 61,
        "_62": 84,
        "_64": 85
    },
    "bhPM7FsN2H1vnBUrxrg6v:100.00:3",
    [
        86
    ],
    {
        "_67": 87,
        "_69": 70,
        "_71": 88
    },
    "1923022511",
    "6VUF6Z1JaUKZF7RS6uSjUu:100.00:6",
    "232791851",
    {
        "_13": 89,
        "_60": 24,
        "_62": 91,
        "_64": 92
    },
    "default",
    [
        93,
        96,
        99
    ],
    {
        "_67": 94,
        "_69": 70,
        "_71": 95
    },
    "3922476776",
    "1DS1QvDa6IFq9C1oJfgtU9",
    {
        "_67": 97,
        "_69": 70,
        "_71": 98
    },
    "749124420",
    "2MQYHJjfKwcTr14d1bOuVH:100.00:2",
    {
        "_67": 100,
        "_69": 101,
        "_71": 102
    },
    "566128514",
    "false",
    "4P1FctCTa3aaKSskEnEeMt",
    "374768818",
    {
        "_13": 103,
        "_60": 61,
        "_62": 105,
        "_64": 106
    },
    "wA7D0MWpe3uCf9HA5KeEi",
    [],
    "491279851",
    {
        "_13": 107,
        "_60": 61,
        "_62": 109,
        "_64": 110
    },
    "4qtiGR7vlvMtZnfSlXM5RN:100.00:12",
    [],
    "507664831",
    {
        "_13": 111,
        "_60": 24,
        "_62": 113,
        "_64": 114
    },
    "4SZ1s8XXvwaDrAV1l6wIro",
    [],
    "645560164",
    {
        "_13": 115,
        "_60": 24,
        "_62": 91,
        "_64": 117
    },
    [],
    "773249106",
    {
        "_13": 118,
        "_60": 24,
        "_62": 120,
        "_64": 121
    },
    "1kGO9xYmxaBS2V2H3LcQuG",
    [],
    "989108178",
    {
        "_13": 122,
        "_60": 24,
        "_62": 124,
        "_64": 125
    },
    "4sTodKrNyByM4guZ68MORR",
    [
        126,
        129
    ],
    {
        "_67": 127,
        "_69": 101,
        "_71": 128
    },
    "1457171347",
    "2EjTipm6C4kk4fuvcHMzZe",
    {
        "_67": 130,
        "_69": 70,
        "_71": 131
    },
    "1426009137",
    "4C2vO0R7mvnCZvl1HDBExp:30.00:5",
    "1028682714",
    {
        "_13": 132,
        "_60": 61,
        "_62": 134,
        "_64": 135
    },
    "735n03snBvba4AEhd2Qwqu:100.00:3",
    [],
    "1072178956",
    {
        "_13": 136,
        "_60": 24,
        "_62": 138,
        "_64": 139
    },
    "4m8JwKa5kCi9HNf1ScZepj",
    [],
    "1242184140",
    {
        "_13": 140,
        "_60": 24,
        "_62": 91,
        "_64": 142
    },
    [],
    "1318146997",
    {
        "_13": 143,
        "_60": 61,
        "_62": 145,
        "_64": 146
    },
    "2AclmEgqaQBVFbxz37XKzy:100.00:5",
    [],
    "1393076427",
    {
        "_13": 147,
        "_60": 61,
        "_62": 149,
        "_64": 150
    },
    "disabled",
    [],
    "1508312659",
    {
        "_13": 151,
        "_60": 24,
        "_62": 91,
        "_64": 153
    },
    [],
    "1578703058",
    {
        "_13": 154,
        "_60": 61,
        "_62": 156,
        "_64": 157
    },
    "2l4nEVMUnPuXkgprUm5zzs:100.00:4",
    [],
    "1611573287",
    {
        "_13": 158,
        "_60": 61,
        "_62": 160,
        "_64": 161
    },
    "159rwM3sBnviE9XWH24azn:100.00:2",
    [],
    "1719651090",
    {
        "_13": 162,
        "_60": 61,
        "_62": 164,
        "_64": 165
    },
    "60QaTyBFJYTakinhLvhAM9",
    [
        166,
        169
    ],
    {
        "_67": 167,
        "_69": 70,
        "_71": 168
    },
    "1616485584",
    "2PP6pudW64Hn7katvazhAx:100.00:5",
    {
        "_67": 170,
        "_69": 70,
        "_71": 171
    },
    "1034043359",
    "4bd3o553p0ZCRkFmipROd8",
    "1804926979",
    {
        "_13": 172,
        "_60": 24,
        "_62": 91,
        "_64": 174
    },
    [],
    "1825130190",
    {
        "_13": 175,
        "_60": 61,
        "_62": 177,
        "_64": 178
    },
    "Nef2uMceNUF9U3ZYwSbpD",
    [],
    "1847911009",
    {
        "_13": 179,
        "_60": 24,
        "_62": 181,
        "_64": 182
    },
    "5OIO2mI7iQiPRReG1jZ4c2:0.00:7",
    [
        183
    ],
    {
        "_67": 184,
        "_69": 70,
        "_71": 185
    },
    "2304807207",
    "xhzqzk6zPqMb3Qs4GVvJu:100.00:5",
    "1855896025",
    {
        "_13": 186,
        "_60": 24,
        "_62": 91,
        "_64": 188
    },
    [],
    "1902899872",
    {
        "_13": 189,
        "_60": 61,
        "_62": 191,
        "_64": 192
    },
    "58UOuEcFwyqlorfhrWQLlE",
    [
        193
    ],
    {
        "_67": 184,
        "_69": 70,
        "_71": 185
    },
    {
        "_13": 87,
        "_60": 61,
        "_62": 88,
        "_64": 195
    },
    [],
    "1988730211",
    {
        "_13": 196,
        "_60": 61,
        "_62": 198,
        "_64": 199
    },
    "6B9O1B3eHKElKWCUfbcvBL",
    [
        200,
        203
    ],
    {
        "_67": 201,
        "_69": 70,
        "_71": 202
    },
    "3780975974",
    "48uk8ZYa2RpJzkpIyOmqP0:100.00:5",
    {
        "_67": 204,
        "_69": 70,
        "_71": 205
    },
    "3733089528",
    "3vtzosKkaPCfPysd7yBTSf",
    "2044826081",
    {
        "_13": 206,
        "_60": 61,
        "_62": 208,
        "_64": 209
    },
    "6MpInoEzkXvXVvodNQCQWs",
    [
        210,
        211
    ],
    {
        "_67": 68,
        "_69": 70,
        "_71": 72
    },
    {
        "_67": 74,
        "_69": 70,
        "_71": 75
    },
    "2091463435",
    {
        "_13": 212,
        "_60": 61,
        "_62": 214,
        "_64": 215
    },
    "5t78GUS68KOn3bHZd8z7ii:100.00:1",
    [],
    "2113934735",
    {
        "_13": 216,
        "_60": 24,
        "_62": 91,
        "_64": 218
    },
    [
        219,
        220,
        221
    ],
    {
        "_67": 167,
        "_69": 70,
        "_71": 168
    },
    {
        "_67": 170,
        "_69": 70,
        "_71": 171
    },
    {
        "_67": 162,
        "_69": 70,
        "_71": 164
    },
    "2256850471",
    {
        "_13": 222,
        "_60": 61,
        "_62": 224,
        "_64": 225
    },
    "IqxordbUxF1Fkg4gfExiY:100.00:1",
    [
        226
    ],
    {
        "_67": 175,
        "_69": 70,
        "_71": 177
    },
    "2293185713",
    {
        "_13": 227,
        "_60": 24,
        "_62": 91,
        "_64": 229
    },
    [],
    {
        "_13": 184,
        "_60": 61,
        "_62": 185,
        "_64": 231
    },
    [],
    "2311599525",
    {
        "_13": 232,
        "_60": 24,
        "_62": 91,
        "_64": 234
    },
    [],
    "2335877601",
    {
        "_13": 235,
        "_60": 24,
        "_62": 237,
        "_64": 238
    },
    "6NQcdu7pgfp18Sq2tfBC6q",
    [],
    "2454940646",
    {
        "_13": 239,
        "_60": 61,
        "_62": 241,
        "_64": 242
    },
    "zol8dYvq8kKfRbOgcM0IF",
    [
        243,
        244
    ],
    {
        "_67": 201,
        "_69": 70,
        "_71": 202
    },
    {
        "_67": 204,
        "_69": 70,
        "_71": 205
    },
    "2494375100",
    {
        "_13": 245,
        "_60": 24,
        "_62": 91,
        "_64": 247
    },
    [],
    "2562876640",
    {
        "_13": 248,
        "_60": 61,
        "_62": 250,
        "_64": 251
    },
    "326czTZeZ0RX0ypR0c5Bb6:100.00:15",
    [],
    "2607001979",
    {
        "_13": 252,
        "_60": 24,
        "_62": 254,
        "_64": 255
    },
    "35jfNEnEKwGsryxcwFhAKz",
    [],
    "2687575887",
    {
        "_13": 256,
        "_60": 61,
        "_62": 258,
        "_64": 259
    },
    "10cvQmwrcZvpWBFlZgn8pZ",
    [],
    "2756095923",
    {
        "_13": 260,
        "_60": 61,
        "_62": 262,
        "_64": 263
    },
    "6jPp6nW1wQVJbfY0uwQgmv:100.00:1",
    [],
    "2868048419",
    {
        "_13": 264,
        "_60": 24,
        "_62": 266,
        "_64": 267
    },
    "7iUNAbafRQfKTvYI2mmFZB",
    [],
    "3054422710",
    {
        "_13": 268,
        "_60": 61,
        "_62": 149,
        "_64": 270
    },
    [],
    "3286474446",
    {
        "_13": 271,
        "_60": 61,
        "_62": 273,
        "_64": 274
    },
    "2a7wA6tOQ5GPb7WIr1SU1A:100.00:1",
    [],
    "3325813340",
    {
        "_13": 275,
        "_60": 61,
        "_62": 277,
        "_64": 278
    },
    "37GsRLj07CqERPyHBn4o5L",
    [
        279
    ],
    {
        "_67": 184,
        "_69": 70,
        "_71": 185
    },
    "3342258807",
    {
        "_13": 280,
        "_60": 24,
        "_62": 282,
        "_64": 283
    },
    "3m0ycr0cMQOm6eMQQjgyp9",
    [],
    "3376455464",
    {
        "_13": 284,
        "_60": 24,
        "_62": 91,
        "_64": 286
    },
    [],
    "3468624635",
    {
        "_13": 287,
        "_60": 24,
        "_62": 91,
        "_64": 289
    },
    [
        290
    ],
    {
        "_67": 291,
        "_69": 101,
        "_71": 292
    },
    "2067628123",
    "3CuBjEMi97tY3EGnq0NA9s",
    "3544641259",
    {
        "_13": 293,
        "_60": 24,
        "_62": 91,
        "_64": 295
    },
    [
        296,
        298
    ],
    {
        "_67": 297,
        "_69": 101,
        "_71": 91
    },
    "2856133350",
    {
        "_67": 299,
        "_69": 101,
        "_71": 91
    },
    "3214154973",
    "3645668434",
    {
        "_13": 300,
        "_60": 61,
        "_62": 302,
        "_64": 303
    },
    "1CWwhBKuOiRAC9V8HRBJRU",
    [
        304
    ],
    {
        "_67": 305,
        "_69": 70,
        "_71": 149
    },
    "3863445312",
    "3700195277",
    {
        "_13": 306,
        "_60": 24,
        "_62": 91,
        "_64": 308
    },
    [],
    {
        "_13": 201,
        "_60": 61,
        "_62": 202,
        "_64": 310
    },
    [],
    "3802510433",
    {
        "_13": 311,
        "_60": 61,
        "_62": 313,
        "_64": 314
    },
    "6FLEMI2GBFmVWGEsEGyASD:100.00:5",
    [],
    "3822950319",
    {
        "_13": 315,
        "_60": 61,
        "_62": 317,
        "_64": 318
    },
    "2CBvDiHjHIK9xlL4ItyXmK:100.00:1",
    [],
    "3838495619",
    {
        "_13": 319,
        "_60": 24,
        "_62": 91,
        "_64": 321
    },
    [],
    {
        "_13": 94,
        "_60": 61,
        "_62": 95,
        "_64": 323
    },
    [],
    "3940160259",
    {
        "_13": 324,
        "_60": 61,
        "_62": 326,
        "_64": 327
    },
    "2mmE1EmtOqtbWemO2wGuMO:100.00:4",
    [
        328,
        330,
        332
    ],
    {
        "_67": 329,
        "_69": 101,
        "_71": 91
    },
    "4180060165",
    {
        "_67": 331,
        "_69": 101,
        "_71": 91
    },
    "3765213438",
    {
        "_67": 333,
        "_69": 70,
        "_71": 334
    },
    "4078831437",
    "6bgwAROz7oF1OcKWxH4vHm:100.00:6",
    "3954884439",
    {
        "_13": 335,
        "_60": 61,
        "_62": 337,
        "_64": 338
    },
    "5rqjCf7T9KpJtLnaE73Kum:100.00:4",
    [],
    {
        "_13": 333,
        "_60": 61,
        "_62": 334,
        "_64": 340
    },
    [
        341,
        342
    ],
    {
        "_67": 329,
        "_69": 101,
        "_71": 91
    },
    {
        "_67": 331,
        "_69": 101,
        "_71": 91
    },
    "4207619515",
    {
        "_13": 343,
        "_60": 24,
        "_62": 91,
        "_64": 345
    },
    [],
    "4226692983",
    {
        "_13": 346,
        "_60": 61,
        "_62": 348,
        "_64": 349
    },
    "6sEu91zwlBGSKOqFiNpGlA:100.00:2",
    [],
    "dynamic_configs",
    {
        "_352": 353,
        "_365": 366,
        "_371": 372,
        "_375": 376,
        "_387": 388,
        "_393": 394,
        "_411": 412,
        "_416": 417,
        "_421": 422,
        "_426": 427,
        "_430": 431,
        "_436": 437
    },
    "357305500",
    {
        "_13": 352,
        "_60": 354,
        "_356": 357,
        "_62": 357,
        "_358": 24,
        "_64": 359,
        "_363": 24,
        "_364": 24
    },
    {
        "_355": 61
    },
    "can_see_upsell",
    "group",
    "launchedGroup",
    "is_device_based",
    [
        360
    ],
    {
        "_67": 361,
        "_69": 70,
        "_71": 362
    },
    "317829697",
    "598ORr5O5ZardhhzMhz8k0:100.00:11",
    "is_user_in_experiment",
    "is_experiment_active",
    "954359911",
    {
        "_13": 365,
        "_60": 367,
        "_356": 369,
        "_62": 369,
        "_358": 24,
        "_64": 370,
        "_363": 61,
        "_364": 61
    },
    {
        "_368": 24
    },
    "enabled",
    "5zN2l0bhNBO2gpivWHXwRY",
    [],
    "1001765573",
    {
        "_13": 371,
        "_60": 373,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 374
    },
    {},
    [],
    "1146308370",
    {
        "_13": 375,
        "_60": 377,
        "_356": 379,
        "_62": 379,
        "_358": 24,
        "_64": 380,
        "_363": 61,
        "_364": 61
    },
    {
        "_378": 61
    },
    "enable-copy-and-open",
    "4jQR01pTnwmjITqDD8PD2s",
    [
        381,
        384
    ],
    {
        "_67": 382,
        "_69": 70,
        "_71": 383
    },
    "303767167",
    "4kquVSCZpyFb5Sqki2BagX:100.00:6",
    {
        "_67": 385,
        "_69": 70,
        "_71": 386
    },
    "3284359640",
    "1E2e7sRUWkvJybjXfbiuoB",
    "1165680819",
    {
        "_13": 387,
        "_60": 389,
        "_356": 391,
        "_62": 391,
        "_358": 24,
        "_64": 392,
        "_363": 61,
        "_364": 61
    },
    {
        "_390": 61
    },
    "show_new_banner",
    "VVjatl8N5mxurs3Cje5TV",
    [],
    "1967546325",
    {
        "_13": 393,
        "_60": 395,
        "_356": 408,
        "_62": 408,
        "_358": 24,
        "_64": 409
    },
    {
        "_396": 61,
        "_397": 61,
        "_398": 24,
        "_399": 24,
        "_400": 61,
        "_401": 61,
        "_402": 403,
        "_404": 403,
        "_405": 406,
        "_407": 61
    },
    "gdrivePicker",
    "o365Picker",
    "gdriveLink",
    "o365Link",
    "o365PersonalLink",
    "o365BusinessLink",
    "gdrivePercentage",
    100,
    "o365Percentage",
    "loadTestPercentage",
    0,
    "showWorkspaceSettings",
    "2bcszlc7CFHdfdCdq7jXNb:100.00:5",
    [
        410
    ],
    {
        "_67": 297,
        "_69": 101,
        "_71": 91
    },
    "2043237793",
    {
        "_13": 411,
        "_60": 413,
        "_356": 357,
        "_62": 357,
        "_358": 24,
        "_64": 415,
        "_363": 24,
        "_364": 24
    },
    {
        "_414": 60
    },
    "bucket",
    [],
    "2513291161",
    {
        "_13": 416,
        "_60": 418,
        "_356": 419,
        "_62": 419,
        "_358": 24,
        "_64": 420,
        "_363": 61,
        "_364": 61
    },
    {
        "_368": 24
    },
    "2FTh6vlZcd8ha1OXcdnD3J",
    [],
    "3159301283",
    {
        "_13": 421,
        "_60": 423,
        "_356": 424,
        "_62": 424,
        "_358": 24,
        "_64": 425,
        "_363": 61,
        "_364": 61
    },
    {
        "_368": 61
    },
    "3Cce6z1hcoF2UBYwyupFck",
    [],
    "3217984440",
    {
        "_13": 426,
        "_60": 428,
        "_356": 357,
        "_62": 357,
        "_358": 24,
        "_64": 429,
        "_363": 24,
        "_364": 24
    },
    {
        "_368": 61
    },
    [],
    "3230069703",
    {
        "_13": 430,
        "_60": 432,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 435
    },
    {
        "_433": 434
    },
    "expirySeconds",
    15,
    [],
    "4198227845",
    {
        "_13": 436,
        "_60": 438,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 445
    },
    {
        "_439": 24,
        "_440": 24,
        "_441": 24,
        "_442": 24,
        "_443": 24,
        "_444": 24
    },
    "enabled_for_platform_override",
    "enabled_for_platform_new",
    "enabled_for_platform_existing",
    "enabled_for_chat_override",
    "enabled_for_chat_new",
    "enabled_for_chat_existing",
    [],
    "layer_configs",
    {
        "_448": 449,
        "_472": 473,
        "_477": 478,
        "_483": 484,
        "_495": 496,
        "_501": 502,
        "_506": 507,
        "_517": 518,
        "_551": 552,
        "_558": 559,
        "_571": 572,
        "_577": 578,
        "_583": 584,
        "_592": 593,
        "_607": 608,
        "_626": 627,
        "_643": 644,
        "_656": 657,
        "_665": 666,
        "_675": 676,
        "_686": 687,
        "_698": 699,
        "_705": 706,
        "_719": 720,
        "_727": 728,
        "_737": 738,
        "_746": 747,
        "_752": 753,
        "_758": 759,
        "_792": 793,
        "_808": 809,
        "_815": 816,
        "_822": 823
    },
    "16152997",
    {
        "_13": 448,
        "_60": 450,
        "_356": 461,
        "_62": 461,
        "_358": 24,
        "_64": 462,
        "_466": 467,
        "_468": 469,
        "_364": 24,
        "_363": 24,
        "_470": 471
    },
    {
        "_451": 61,
        "_452": 24,
        "_453": 61,
        "_454": 455,
        "_456": 455,
        "_457": 406,
        "_458": 24,
        "_459": 61,
        "_460": 24
    },
    "show_preview_when_collapsed",
    "expand_by_default",
    "is_enabled",
    "summarizer_system_prompt",
    "",
    "summarizer_chunk_template",
    "summarizer_chunk_char_limit",
    "enable_o3_mini_retrieval",
    "override_o3_mini_to_high",
    "enable_reason_by_default",
    "6DaNqHbUdaQZCJTtuXMn3l:override",
    [
        463
    ],
    {
        "_67": 464,
        "_69": 70,
        "_71": 465
    },
    "747145983",
    "1yBei0bniPE2f1TkI3MLWa",
    "explicit_parameters",
    [
        451,
        452,
        453
    ],
    "allocated_experiment_name",
    "1630255509",
    "undelegated_secondary_exposures",
    [
        463
    ],
    "40440673",
    {
        "_13": 472,
        "_60": 474,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 475,
        "_466": 476,
        "_470": 475
    },
    {},
    [],
    [],
    "51287004",
    {
        "_13": 477,
        "_60": 479,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 481,
        "_466": 482,
        "_470": 481
    },
    {
        "_480": 61
    },
    "enable",
    [],
    [],
    "183390215",
    {
        "_13": 483,
        "_60": 485,
        "_356": 488,
        "_62": 488,
        "_358": 61,
        "_64": 489,
        "_466": 492,
        "_468": 493,
        "_364": 61,
        "_363": 24,
        "_470": 494
    },
    {
        "_486": 24,
        "_487": 24
    },
    "signup_allow_phone",
    "in_phone_signup_holdout",
    "targetingGate",
    [
        490
    ],
    {
        "_67": 491,
        "_69": 101,
        "_71": 91
    },
    "3874938189",
    [
        486,
        487
    ],
    "4005636946",
    [],
    "190694971",
    {
        "_13": 495,
        "_60": 497,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 499,
        "_466": 500,
        "_470": 499
    },
    {
        "_498": 24
    },
    "show_nux",
    [],
    [],
    "229662723",
    {
        "_13": 501,
        "_60": 503,
        "_356": 91,
        "_62": 91,
        "_358": 61,
        "_64": 504,
        "_466": 505,
        "_470": 504
    },
    {},
    [],
    [],
    "387752763",
    {
        "_13": 506,
        "_60": 508,
        "_356": 91,
        "_62": 91,
        "_358": 61,
        "_64": 511,
        "_466": 516,
        "_470": 511
    },
    {
        "_509": 61,
        "_510": 61
    },
    "enable_slash_commands",
    "enable_rich_text_composer",
    [
        512,
        513,
        514
    ],
    {
        "_67": 97,
        "_69": 70,
        "_71": 98
    },
    {
        "_67": 100,
        "_69": 101,
        "_71": 102
    },
    {
        "_67": 515,
        "_69": 101,
        "_71": 91
    },
    "1410082514",
    [],
    "468168202",
    {
        "_13": 517,
        "_60": 519,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 547,
        "_466": 550,
        "_470": 547
    },
    {
        "_520": 61,
        "_521": 24,
        "_522": 61,
        "_523": 61,
        "_524": 24,
        "_525": 24,
        "_526": 24,
        "_527": 24,
        "_528": 24,
        "_529": 24,
        "_530": 24,
        "_531": 24,
        "_532": 24,
        "_533": 24,
        "_534": 61,
        "_535": 61,
        "_536": 24,
        "_537": 61,
        "_538": 61,
        "_539": 540,
        "_541": 542,
        "_543": 24,
        "_544": 545,
        "_546": 24
    },
    "is_team_enabled",
    "is_yearly_plus_subscription_enabled",
    "is_split_between_personal_and_business_enabled",
    "is_modal_fullscreen",
    "is_v2_toggle_labels_enabled",
    "is_bw",
    "is_produce_colors",
    "is_produce_color_scheme",
    "is_mobile_web_toggle_enabled",
    "is_enterprise_enabled",
    "is_produce_text",
    "is_optimized_checkout",
    "is_save_stripe_payment_info_enabled",
    "is_auto_save_stripe_payment_info_enabled",
    "does_manage_my_subscription_link_take_user_to_subscription_settings",
    "should_open_cancellation_survey_after_canceling",
    "should_cancel_button_take_user_to_stripe",
    "should_show_manage_my_subscription_link",
    "is_stripe_manage_subscription_link_enabled",
    "cancellation_modal_cancel_button_color",
    "danger",
    "cancellation_modal_go_back_button_color",
    "secondary",
    "should_show_cp",
    "cp_eligibility_months",
    3,
    "should_offer_paypal_when_eligible",
    [
        548
    ],
    {
        "_67": 549,
        "_69": 101,
        "_71": 91
    },
    "1847092144",
    [],
    "668322707",
    {
        "_13": 551,
        "_60": 553,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 556,
        "_466": 557,
        "_470": 556
    },
    {
        "_554": 61,
        "_555": 61
    },
    "show_citations_with_title",
    "use_chip_style_citations",
    [],
    [],
    "871635014",
    {
        "_13": 558,
        "_60": 560,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 569,
        "_466": 570,
        "_470": 569
    },
    {
        "_561": 24,
        "_562": 61,
        "_563": 24,
        "_564": 565,
        "_566": 91,
        "_567": 24,
        "_568": 24
    },
    "snowflake_composer_entry_point",
    "use_broad_rate_limit_language",
    "voice_holdout",
    "krisp_noise_filter",
    "none",
    "voice_entry_point_style",
    "show_label_on_button",
    "voice_only",
    [],
    [],
    "1170120107",
    {
        "_13": 571,
        "_60": 573,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 575,
        "_466": 576,
        "_470": 575
    },
    {
        "_574": 24
    },
    "is_whisper_enabled",
    [],
    [],
    "1238742812",
    {
        "_13": 577,
        "_60": 579,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 581,
        "_466": 582,
        "_470": 581
    },
    {
        "_580": 24
    },
    "should_enable_zh_tw",
    [],
    [],
    "1320801051",
    {
        "_13": 583,
        "_60": 585,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 590,
        "_466": 591,
        "_470": 590
    },
    {
        "_586": 24,
        "_587": 24,
        "_588": 61,
        "_589": 24
    },
    "hide_new_at_workspace_section",
    "hide_section_new_at_workspace",
    "gpt_discovery_experiment_enabled",
    "popular_at_my_workspace_enabled",
    [],
    [],
    "1346366956",
    {
        "_13": 592,
        "_60": 594,
        "_356": 91,
        "_62": 91,
        "_358": 61,
        "_64": 605,
        "_466": 606,
        "_470": 605
    },
    {
        "_595": 24,
        "_596": 597,
        "_598": 24,
        "_486": 24,
        "_599": 24,
        "_600": 24,
        "_601": 24,
        "_602": 24,
        "_603": 604
    },
    "use_email_otp",
    "signup_cta_copy",
    "SIGN_UP",
    "login_allow_phone",
    "forwardToAuthApi",
    "use_new_phone_ui",
    "in_signup_allow_phone_hold_out",
    "use_formatted_national_number",
    "continue_with_email_phone_placement",
    "after_sso",
    [],
    [],
    "1547743984",
    {
        "_13": 607,
        "_60": 609,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 622,
        "_466": 625,
        "_470": 622
    },
    {
        "_610": 24,
        "_611": 24,
        "_612": 24,
        "_613": 24,
        "_614": 24,
        "_615": 24,
        "_616": 24,
        "_617": 61,
        "_618": 24,
        "_619": 24,
        "_620": 61,
        "_621": 61
    },
    "should_simplify_modal",
    "is_simplified_sharing_modal_enabled",
    "is_social_share_options_enabled",
    "is_update_shared_links_enabled",
    "is_discoverability_toggle_enabled",
    "show_copylink_state_if_no_updates",
    "is_continue_enabled",
    "show_share_button_text",
    "is_meta_improvements_enabled",
    "show_share_button_inline",
    "use_dalle_preview",
    "in_dalle_preview_exp",
    [
        623
    ],
    {
        "_67": 624,
        "_69": 101,
        "_71": 91
    },
    "4038001028",
    [],
    "1630876919",
    {
        "_13": 626,
        "_60": 628,
        "_356": 635,
        "_62": 635,
        "_358": 24,
        "_64": 636,
        "_466": 640,
        "_468": 641,
        "_364": 61,
        "_363": 61,
        "_470": 642
    },
    {
        "_629": 61,
        "_630": 61,
        "_631": 61,
        "_632": 61,
        "_633": 24,
        "_634": 61
    },
    "enable_indexing",
    "backfill_completed",
    "enable_local_indexing",
    "enable_ux",
    "enable_us_rollout",
    "enable_ux_rollout",
    "31UyKaWB8PZhFswQt29NlZ",
    [
        637
    ],
    {
        "_67": 638,
        "_69": 70,
        "_71": 639
    },
    "2372319800",
    "4NZS9cdXgw2uEnVQCdyNMH:100.00:30",
    [
        629,
        631,
        630,
        632,
        634
    ],
    "1028722647",
    [],
    "1696863369",
    {
        "_13": 643,
        "_60": 645,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 648,
        "_466": 655,
        "_470": 648
    },
    {
        "_646": 24,
        "_647": 24
    },
    "has_sidekick_access",
    "show_nux_banner",
    [
        649,
        652
    ],
    {
        "_67": 650,
        "_69": 101,
        "_71": 651
    },
    "1938289220",
    "79O8DQPDmTKxnLdAH9loVk",
    {
        "_67": 653,
        "_69": 101,
        "_71": 654
    },
    "2033872549",
    "7dScmNU0bu2UQuzCNtva50",
    [],
    "1697140512",
    {
        "_13": 656,
        "_60": 658,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 660,
        "_466": 664,
        "_470": 660
    },
    {
        "_647": 24,
        "_659": 24
    },
    "can_download_sidetron",
    [
        661
    ],
    {
        "_67": 662,
        "_69": 101,
        "_71": 663
    },
    "2919213474",
    "6HLlb6nSjJk5ADynHucWgP",
    [],
    "1704793646",
    {
        "_13": 665,
        "_60": 667,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 671,
        "_466": 674,
        "_470": 671
    },
    {
        "_668": 24,
        "_669": 670
    },
    "greeting_web",
    "name_char_limit",
    20,
    [
        672
    ],
    {
        "_67": 673,
        "_69": 101,
        "_71": 91
    },
    "331938894",
    [],
    "1780960461",
    {
        "_13": 675,
        "_60": 677,
        "_356": 488,
        "_62": 488,
        "_358": 24,
        "_64": 680,
        "_466": 683,
        "_468": 684,
        "_364": 61,
        "_363": 24,
        "_470": 685
    },
    {
        "_678": 24,
        "_679": 24,
        "_668": 24
    },
    "mobile",
    "web",
    [
        681
    ],
    {
        "_67": 682,
        "_69": 101,
        "_71": 91
    },
    "3074373870",
    [
        678,
        679
    ],
    "2198260923",
    [],
    "1914829685",
    {
        "_13": 686,
        "_60": 688,
        "_356": 690,
        "_62": 690,
        "_358": 61,
        "_64": 691,
        "_466": 695,
        "_468": 696,
        "_364": 24,
        "_363": 24,
        "_470": 697
    },
    {
        "_689": 61
    },
    "forward_to_authapi",
    "2RO4BOrVWPrsxRUPYNKPLe:override",
    [
        692
    ],
    {
        "_67": 693,
        "_69": 70,
        "_71": 694
    },
    "14938527",
    "3QgLJ91lKIc7VAOjo5SDz7",
    [
        689
    ],
    "1856338298",
    [
        692
    ],
    "2152104812",
    {
        "_13": 698,
        "_60": 700,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 703,
        "_466": 704,
        "_470": 703
    },
    {
        "_701": 24,
        "_702": 24
    },
    "hide_gpts_if_none",
    "hide_default_gpts",
    [],
    [],
    "3048336830",
    {
        "_13": 705,
        "_60": 707,
        "_356": 710,
        "_62": 710,
        "_358": 24,
        "_64": 711,
        "_466": 718,
        "_470": 711
    },
    {
        "_708": 61,
        "_709": 24
    },
    "is-enabled",
    "use-rtl-layout",
    "localization-april Nzc6Xnht6tIVmb48Ejg1T:override",
    [
        712,
        715
    ],
    {
        "_67": 713,
        "_69": 101,
        "_71": 714
    },
    "3922145230",
    "14DZA2LumaPqAdCo52CrUB",
    {
        "_67": 716,
        "_69": 70,
        "_71": 717
    },
    "3700615661",
    "66covjaoZoe9pQR4I68jOB",
    [],
    "3178812292",
    {
        "_13": 719,
        "_60": 721,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 723,
        "_466": 726,
        "_470": 723
    },
    {
        "_722": 24
    },
    "use_f_convo",
    [
        724
    ],
    {
        "_67": 725,
        "_69": 101,
        "_71": 91
    },
    "3799260860",
    [],
    "3436367576",
    {
        "_13": 727,
        "_60": 729,
        "_356": 488,
        "_62": 488,
        "_358": 24,
        "_64": 731,
        "_466": 734,
        "_468": 735,
        "_364": 61,
        "_363": 24,
        "_470": 736
    },
    {
        "_629": 24,
        "_730": 406,
        "_632": 24,
        "_631": 24,
        "_630": 24
    },
    "wave",
    [
        732
    ],
    {
        "_67": 733,
        "_69": 101,
        "_71": 91
    },
    "1221279314",
    [
        629,
        730,
        630,
        632,
        631
    ],
    "938456440",
    [],
    "3471271313",
    {
        "_13": 737,
        "_60": 739,
        "_356": 741,
        "_62": 741,
        "_358": 61,
        "_64": 742,
        "_466": 743,
        "_468": 744,
        "_364": 24,
        "_363": 24,
        "_470": 745
    },
    {
        "_740": 24
    },
    "show_upsell",
    "prestart",
    [],
    [
        740
    ],
    "3021307436",
    [],
    "3517133692",
    {
        "_13": 746,
        "_60": 748,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 750,
        "_466": 751,
        "_470": 750
    },
    {
        "_749": 24
    },
    "is_memory_undo_enabled",
    [],
    [],
    "3590606857",
    {
        "_13": 752,
        "_60": 754,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 756,
        "_466": 757,
        "_470": 756
    },
    {
        "_755": 24
    },
    "should_offer_paypal",
    [],
    [],
    "3637408529",
    {
        "_13": 758,
        "_60": 760,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 786,
        "_466": 791,
        "_470": 786
    },
    {
        "_761": 61,
        "_762": 24,
        "_763": 24,
        "_764": 24,
        "_765": 766,
        "_767": 768,
        "_769": 61,
        "_770": 61,
        "_771": 61,
        "_772": 24,
        "_773": 61,
        "_774": 24,
        "_775": 24,
        "_776": 61,
        "_777": 24,
        "_778": 61,
        "_779": 545,
        "_780": 781,
        "_782": 61,
        "_783": 784,
        "_785": 24
    },
    "is_anon_chat_enabled",
    "is_anon_chat_enabled_for_new_users_only",
    "is_try_it_first_on_login_page_enabled",
    "is_no_auth_welcome_modal_enabled",
    "no_auth_soft_rate_limit",
    5,
    "no_auth_hard_rate_limit",
    1200,
    "should_show_no_auth_signup_banner",
    "is_no_auth_welcome_back_modal_enabled",
    "is_no_auth_soft_rate_limit_modal_enabled",
    "is_no_auth_gpt4o_modal_enabled",
    "is_login_primary_button",
    "is_desktop_primary_auth_button_on_right",
    "is_primary_btn_blue",
    "should_show_disclaimer_only_once_per_device",
    "is_secondary_banner_button_enabled",
    "is_secondary_auth_banner_button_enabled",
    "no_auth_banner_signup_rate_limit",
    "composer_text",
    "ASK_ANYTHING",
    "is_in_composer_text_exp",
    "no_auth_upsell_wording",
    "NO_CHANGE",
    "should_refresh_access_token_error_take_user_to_no_auth",
    [
        787,
        789
    ],
    {
        "_67": 788,
        "_69": 101,
        "_71": 149
    },
    "3238165271",
    {
        "_67": 790,
        "_69": 101,
        "_71": 149
    },
    "2983591614",
    [],
    "3647926857",
    {
        "_13": 792,
        "_60": 794,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 802,
        "_466": 807,
        "_470": 802
    },
    {
        "_795": 61,
        "_796": 24,
        "_797": 798,
        "_799": 24,
        "_800": 24,
        "_801": 565
    },
    "unified_architecture",
    "ux_updates",
    "inference_debounce_ms",
    400,
    "autoswitcher_enabled",
    "copy-and-link",
    "reasoning_slider",
    [
        803,
        805
    ],
    {
        "_67": 804,
        "_69": 101,
        "_71": 91
    },
    "850280859",
    {
        "_67": 806,
        "_69": 101,
        "_71": 91
    },
    "13512905",
    [],
    "3711177917",
    {
        "_13": 808,
        "_60": 810,
        "_356": 91,
        "_62": 91,
        "_358": 61,
        "_64": 813,
        "_466": 814,
        "_470": 813
    },
    {
        "_811": 24,
        "_812": 61
    },
    "is_summarizer_default_expanded",
    "is_inline_summarizer_enabled",
    [],
    [],
    "3972089454",
    {
        "_13": 815,
        "_60": 817,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 820,
        "_466": 821,
        "_470": 820
    },
    {
        "_818": 819
    },
    "search_scoring_dyconfig_name",
    "gizmo_search_score_config",
    [],
    [],
    "4211831761",
    {
        "_13": 822,
        "_60": 824,
        "_356": 91,
        "_62": 91,
        "_358": 61,
        "_64": 825,
        "_466": 826,
        "_470": 825
    },
    {
        "_368": 24
    },
    [],
    [],
    "sdkParams",
    {},
    "has_updates",
    "generator",
    "statsig-node-sdk",
    "sdkInfo",
    {
        "_834": 835,
        "_836": 837
    },
    "sdkType",
    "statsig-node",
    "sdkVersion",
    "5.26.0",
    "time",
    1742979946815,
    "evaluated_keys",
    {
        "_842": 12,
        "_843": 844
    },
    "userID",
    "user-chatgpt",
    {
        "_845": 846,
        "_847": 846,
        "_848": 846,
        "_849": 33,
        "_850": 33
    },
    "WebAnonymousCookieID",
    "f10eee0a-3567-4980-ad81-98da119a5c6b",
    "DeviceId",
    "stableID",
    "workspace_id",
    "account_id",
    "hash_used",
    "djb2",
    {
        "_842": 12,
        "_854": 855,
        "_856": 857,
        "_50": 51,
        "_843": 844,
        "_861": 862
    },
    "country",
    "SG",
    "custom",
    {
        "_858": 35,
        "_849": 33,
        "_850": 33,
        "_859": 61,
        "_860": 6
    },
    "plan_type",
    "is_paid",
    "auth_status",
    "statsigEnvironment",
    {
        "_863": 864
    },
    "tier",
    "production",
    "experimental",
    {
        "_56": 867,
        "_350": 1190,
        "_446": 1293,
        "_827": 1818,
        "_829": 61,
        "_830": 831,
        "_832": 1819,
        "_838": 839,
        "_840": 1820,
        "_851": 852,
        "_9": 1822
    },
    {
        "_693": 868,
        "_870": 871,
        "_873": 874,
        "_877": 878,
        "_883": 884,
        "_887": 888,
        "_891": 892,
        "_895": 896,
        "_898": 899,
        "_902": 903,
        "_905": 906,
        "_910": 911,
        "_914": 915,
        "_918": 919,
        "_927": 928,
        "_931": 932,
        "_935": 936,
        "_939": 940,
        "_943": 944,
        "_947": 948,
        "_132": 954,
        "_956": 957,
        "_959": 960,
        "_962": 963,
        "_965": 966,
        "_968": 969,
        "_973": 974,
        "_977": 978,
        "_981": 982,
        "_985": 986,
        "_988": 989,
        "_992": 993,
        "_996": 997,
        "_1000": 1001,
        "_1003": 1004,
        "_1007": 1008,
        "_1010": 1011,
        "_1014": 1015,
        "_1018": 1019,
        "_1021": 1022,
        "_1024": 1025,
        "_1027": 1028,
        "_1031": 1032,
        "_291": 1034,
        "_1036": 1037,
        "_1039": 1040,
        "_1043": 1044,
        "_1047": 1048,
        "_1053": 1054,
        "_1057": 1058,
        "_1061": 1062,
        "_1065": 1066,
        "_1068": 1069,
        "_1072": 1073,
        "_1076": 1077,
        "_1080": 1081,
        "_1086": 1087,
        "_1090": 1091,
        "_1096": 1097,
        "_1101": 1102,
        "_1105": 1106,
        "_1108": 1109,
        "_1111": 1112,
        "_1114": 1115,
        "_1117": 1118,
        "_1120": 1121,
        "_1124": 1125,
        "_1127": 1128,
        "_1131": 1132,
        "_293": 1134,
        "_1138": 1139,
        "_1142": 1143,
        "_1146": 1147,
        "_1149": 1150,
        "_1152": 1153,
        "_1156": 1157,
        "_713": 1160,
        "_1162": 1163,
        "_1166": 1167,
        "_1169": 1170,
        "_1172": 1173,
        "_1176": 1177,
        "_1180": 1181,
        "_1183": 1184,
        "_1186": 1187
    },
    {
        "_13": 693,
        "_60": 61,
        "_62": 694,
        "_64": 869
    },
    [],
    "156153730",
    {
        "_13": 870,
        "_60": 24,
        "_62": 91,
        "_64": 872
    },
    [],
    "222560275",
    {
        "_13": 873,
        "_60": 24,
        "_62": 875,
        "_64": 876
    },
    "5pv2QpbgXNDB0QnBo3LTti:10.00:1",
    [],
    "223382091",
    {
        "_13": 877,
        "_60": 24,
        "_62": 879,
        "_64": 880
    },
    "1fKkxDiVebEKfTj8nDAjHe",
    [
        881,
        882
    ],
    {
        "_67": 329,
        "_69": 101,
        "_71": 91
    },
    {
        "_67": 331,
        "_69": 101,
        "_71": 91
    },
    "402391964",
    {
        "_13": 883,
        "_60": 24,
        "_62": 885,
        "_64": 886
    },
    "14sAQaGJDosUKVV0DFZsAL",
    [],
    "471233253",
    {
        "_13": 887,
        "_60": 24,
        "_62": 889,
        "_64": 890
    },
    "3Yf9H7TxMC122pchwAkoLB",
    [],
    "550432558",
    {
        "_13": 891,
        "_60": 61,
        "_62": 893,
        "_64": 894
    },
    "4XbSwfoqBmVtxwz32sweLb",
    [],
    "573184874",
    {
        "_13": 895,
        "_60": 24,
        "_62": 91,
        "_64": 897
    },
    [],
    "582612297",
    {
        "_13": 898,
        "_60": 61,
        "_62": 900,
        "_64": 901
    },
    "5censDsCfS2zQeYtTIui2s:100.00:2",
    [],
    "614413305",
    {
        "_13": 902,
        "_60": 24,
        "_62": 91,
        "_64": 904
    },
    [],
    "653593316",
    {
        "_13": 905,
        "_60": 61,
        "_62": 907,
        "_64": 908
    },
    "GJ8pvorFDIe3Z4WonIr2s",
    [
        909
    ],
    {
        "_67": 311,
        "_69": 70,
        "_71": 313
    },
    "706943082",
    {
        "_13": 910,
        "_60": 61,
        "_62": 912,
        "_64": 913
    },
    "X9mJLzEwXwKo3p0LSSQIL",
    [],
    "727502549",
    {
        "_13": 914,
        "_60": 61,
        "_62": 916,
        "_64": 917
    },
    "6EYbmM9CyqCRO6U6k3dROA",
    [],
    "756982148",
    {
        "_13": 918,
        "_60": 61,
        "_62": 920,
        "_64": 921
    },
    "3oAWYdzegKPwxhFJjJrGz3",
    [
        922,
        924
    ],
    {
        "_67": 923,
        "_69": 101,
        "_71": 91
    },
    "1456438623",
    {
        "_67": 925,
        "_69": 70,
        "_71": 926
    },
    "3805873235",
    "5KvGWxgOdialy0Dx9IrqmW:100.00:23",
    "756982149",
    {
        "_13": 927,
        "_60": 24,
        "_62": 929,
        "_64": 930
    },
    "1rXg44we6gmcRqYsiZzfL4:0.00:1",
    [],
    "795789557",
    {
        "_13": 931,
        "_60": 24,
        "_62": 933,
        "_64": 934
    },
    "2GzNaY2UIV2RYDjl4grJNG:0.00:1",
    [],
    "809056127",
    {
        "_13": 935,
        "_60": 61,
        "_62": 937,
        "_64": 938
    },
    "54ufwSF4KjxPi2AIrjbelh",
    [],
    "810701024",
    {
        "_13": 939,
        "_60": 61,
        "_62": 941,
        "_64": 942
    },
    "6U8ODe5JvFov5zs1rOzJjD",
    [],
    "891514942",
    {
        "_13": 943,
        "_60": 24,
        "_62": 945,
        "_64": 946
    },
    "aWUpylPDtFgWWhTxEsfCx",
    [],
    "989226566",
    {
        "_13": 947,
        "_60": 61,
        "_62": 949,
        "_64": 950
    },
    "6yqqYAWKtmfU8A7QGdiky4",
    [
        951,
        952
    ],
    {
        "_67": 127,
        "_69": 101,
        "_71": 128
    },
    {
        "_67": 130,
        "_69": 70,
        "_71": 953
    },
    "7D8EAif25E3Y8A3zkg6ljp:100.00:2",
    {
        "_13": 132,
        "_60": 61,
        "_62": 134,
        "_64": 955
    },
    [],
    "1032814809",
    {
        "_13": 956,
        "_60": 24,
        "_62": 91,
        "_64": 958
    },
    [],
    "1064007944",
    {
        "_13": 959,
        "_60": 24,
        "_62": 91,
        "_64": 961
    },
    [],
    "1099124727",
    {
        "_13": 962,
        "_60": 24,
        "_62": 91,
        "_64": 964
    },
    [],
    "1154002920",
    {
        "_13": 965,
        "_60": 24,
        "_62": 91,
        "_64": 967
    },
    [],
    "1166240779",
    {
        "_13": 968,
        "_60": 61,
        "_62": 970,
        "_64": 971
    },
    "4UjTXwt2XK975PANdi1Ma6:25.00:5",
    [
        972
    ],
    {
        "_67": 299,
        "_69": 101,
        "_71": 91
    },
    "1214379119",
    {
        "_13": 973,
        "_60": 24,
        "_62": 975,
        "_64": 976
    },
    "3Da3vJtBawdpcHFOEpjzZA:10.00:2",
    [],
    "1382475798",
    {
        "_13": 977,
        "_60": 61,
        "_62": 979,
        "_64": 980
    },
    "3P8OsGy1e5tQlR5dsTIWbL",
    [],
    "1416952492",
    {
        "_13": 981,
        "_60": 24,
        "_62": 983,
        "_64": 984
    },
    "2LD82enCtskHL9Vi2hS6Jq",
    [],
    "1422501431",
    {
        "_13": 985,
        "_60": 24,
        "_62": 91,
        "_64": 987
    },
    [],
    "1439437954",
    {
        "_13": 988,
        "_60": 24,
        "_62": 990,
        "_64": 991
    },
    "11IqDt7xc4mMNiyiSIMy1F:0.00:1",
    [],
    "1456513860",
    {
        "_13": 992,
        "_60": 61,
        "_62": 994,
        "_64": 995
    },
    "jHXkU7q9axp0dXBSyzihH",
    [],
    "1468311859",
    {
        "_13": 996,
        "_60": 24,
        "_62": 998,
        "_64": 999
    },
    "7tfl8ZUhwr5pzErE3ikBej",
    [],
    "1542198993",
    {
        "_13": 1000,
        "_60": 24,
        "_62": 91,
        "_64": 1002
    },
    [],
    "1656345175",
    {
        "_13": 1003,
        "_60": 61,
        "_62": 1005,
        "_64": 1006
    },
    "2CwIChuIr7SLQ2CyqRegF2",
    [],
    "1741586789",
    {
        "_13": 1007,
        "_60": 24,
        "_62": 91,
        "_64": 1009
    },
    [],
    "1760640904",
    {
        "_13": 1010,
        "_60": 61,
        "_62": 1012,
        "_64": 1013
    },
    "6ezOfLAw7fGQPVjfNsReIy",
    [],
    "1830177352",
    {
        "_13": 1014,
        "_60": 61,
        "_62": 1016,
        "_64": 1017
    },
    "44udGr8tXtB3ZIDHLV3HSF",
    [],
    "1839283687",
    {
        "_13": 1018,
        "_60": 24,
        "_62": 91,
        "_64": 1020
    },
    [],
    "1860647109",
    {
        "_13": 1021,
        "_60": 24,
        "_62": 91,
        "_64": 1023
    },
    [],
    "2000076788",
    {
        "_13": 1024,
        "_60": 24,
        "_62": 91,
        "_64": 1026
    },
    [],
    "2053937752",
    {
        "_13": 1027,
        "_60": 24,
        "_62": 1029,
        "_64": 1030
    },
    "2PLQzvwrGPxACRwaEcKbIh",
    [],
    "2056761365",
    {
        "_13": 1031,
        "_60": 24,
        "_62": 91,
        "_64": 1033
    },
    [],
    {
        "_13": 291,
        "_60": 24,
        "_62": 292,
        "_64": 1035
    },
    [],
    "2151954125",
    {
        "_13": 1036,
        "_60": 24,
        "_62": 91,
        "_64": 1038
    },
    [],
    "2153043779",
    {
        "_13": 1039,
        "_60": 61,
        "_62": 1041,
        "_64": 1042
    },
    "DamiTYVoTv9Z9jRFOT5iC",
    [],
    "2173548801",
    {
        "_13": 1043,
        "_60": 61,
        "_62": 1045,
        "_64": 1046
    },
    "22nVhoL17eyMvGWgFrDfZe",
    [],
    "2192543539",
    {
        "_13": 1047,
        "_60": 61,
        "_62": 1049,
        "_64": 1050
    },
    "4Ro1m2dj4fUBe4hcP1YKjj:75.00:3",
    [
        1051
    ],
    {
        "_67": 1052,
        "_69": 101,
        "_71": 91
    },
    "4206244917",
    "2232580636",
    {
        "_13": 1053,
        "_60": 61,
        "_62": 1055,
        "_64": 1056
    },
    "4y4Nd0nF0CFawcrQBbm7Mq:100.00:4",
    [],
    "2281969373",
    {
        "_13": 1057,
        "_60": 61,
        "_62": 1059,
        "_64": 1060
    },
    "6EbVeXErTdGtbchxdqEMTg",
    [],
    "2290870843",
    {
        "_13": 1061,
        "_60": 61,
        "_62": 1063,
        "_64": 1064
    },
    "5dONtElzUeyTTp5FvpWy6",
    [],
    "2360528850",
    {
        "_13": 1065,
        "_60": 24,
        "_62": 91,
        "_64": 1067
    },
    [],
    "2379988365",
    {
        "_13": 1068,
        "_60": 24,
        "_62": 91,
        "_64": 1070
    },
    [
        1071
    ],
    {
        "_67": 297,
        "_69": 101,
        "_71": 91
    },
    "2411734826",
    {
        "_13": 1072,
        "_60": 24,
        "_62": 1074,
        "_64": 1075
    },
    "33U1igAQgegRumGc4LbaB:2.00:1",
    [],
    "2445152477",
    {
        "_13": 1076,
        "_60": 61,
        "_62": 1078,
        "_64": 1079
    },
    "5qtlunRMswJX2JGoF8GikC",
    [],
    "2634628831",
    {
        "_13": 1080,
        "_60": 61,
        "_62": 1082,
        "_64": 1083
    },
    "6LfSag7ByiH0gGcqoFHHBe",
    [
        1084,
        1085
    ],
    {
        "_67": 923,
        "_69": 101,
        "_71": 91
    },
    {
        "_67": 925,
        "_69": 70,
        "_71": 926
    },
    "2637918557",
    {
        "_13": 1086,
        "_60": 61,
        "_62": 1088,
        "_64": 1089
    },
    "2XNTwszL419o7DMxzSa0vz:100.00:1",
    [],
    "2712556596",
    {
        "_13": 1090,
        "_60": 61,
        "_62": 1092,
        "_64": 1093
    },
    "7pPLEbQc7hKT1m7CbondoE",
    [
        1094
    ],
    {
        "_67": 1095,
        "_69": 101,
        "_71": 91
    },
    "135448051",
    "2781425969",
    {
        "_13": 1096,
        "_60": 24,
        "_62": 1098,
        "_64": 1099
    },
    "7BIMlzITwH6mysXL5ILPSw",
    [
        1100
    ],
    {
        "_67": 1095,
        "_69": 101,
        "_71": 91
    },
    "2833534668",
    {
        "_13": 1101,
        "_60": 61,
        "_62": 1103,
        "_64": 1104
    },
    "7uYkibMYlCPSnoWmmYNanm",
    [],
    "2935021756",
    {
        "_13": 1105,
        "_60": 24,
        "_62": 91,
        "_64": 1107
    },
    [],
    "2968810397",
    {
        "_13": 1108,
        "_60": 24,
        "_62": 91,
        "_64": 1110
    },
    [],
    "3058498100",
    {
        "_13": 1111,
        "_60": 24,
        "_62": 91,
        "_64": 1113
    },
    [],
    "3148583717",
    {
        "_13": 1114,
        "_60": 24,
        "_62": 91,
        "_64": 1116
    },
    [],
    "3241763787",
    {
        "_13": 1117,
        "_60": 24,
        "_62": 91,
        "_64": 1119
    },
    [],
    "3257646228",
    {
        "_13": 1120,
        "_60": 24,
        "_62": 1122,
        "_64": 1123
    },
    "3veZ6qhG4zTVvcrwpXXPgi:1.00:4",
    [],
    "3291247717",
    {
        "_13": 1124,
        "_60": 24,
        "_62": 91,
        "_64": 1126
    },
    [],
    "3435450078",
    {
        "_13": 1127,
        "_60": 61,
        "_62": 1129,
        "_64": 1130
    },
    "2qCdHpFuWOOkibzLRL0zgn",
    [],
    "3472722167",
    {
        "_13": 1131,
        "_60": 24,
        "_62": 91,
        "_64": 1133
    },
    [],
    {
        "_13": 293,
        "_60": 24,
        "_62": 91,
        "_64": 1135
    },
    [
        1136,
        1137
    ],
    {
        "_67": 297,
        "_69": 101,
        "_71": 91
    },
    {
        "_67": 299,
        "_69": 101,
        "_71": 91
    },
    "3612584454",
    {
        "_13": 1138,
        "_60": 24,
        "_62": 1140,
        "_64": 1141
    },
    "4fXx7LNuNnDASdmkzwNxtf",
    [],
    "3664702598",
    {
        "_13": 1142,
        "_60": 24,
        "_62": 1144,
        "_64": 1145
    },
    "7x9wS41bRDCji9ns8x5Oej",
    [],
    "3678527908",
    {
        "_13": 1146,
        "_60": 24,
        "_62": 91,
        "_64": 1148
    },
    [],
    "3728856343",
    {
        "_13": 1149,
        "_60": 24,
        "_62": 91,
        "_64": 1151
    },
    [],
    "3861593998",
    {
        "_13": 1152,
        "_60": 24,
        "_62": 1154,
        "_64": 1155
    },
    "5DN2QZNg9iYP45NqvRetnu",
    [],
    "3910241726",
    {
        "_13": 1156,
        "_60": 61,
        "_62": 1158,
        "_64": 1159
    },
    "1ItyvFbGou4epQp9HviAsm",
    [],
    {
        "_13": 713,
        "_60": 24,
        "_62": 714,
        "_64": 1161
    },
    [],
    "3940529303",
    {
        "_13": 1162,
        "_60": 61,
        "_62": 1164,
        "_64": 1165
    },
    "17mkpeWbaWfCeMrpE67FOc",
    [],
    "4012051055",
    {
        "_13": 1166,
        "_60": 24,
        "_62": 91,
        "_64": 1168
    },
    [],
    "4043415092",
    {
        "_13": 1169,
        "_60": 24,
        "_62": 91,
        "_64": 1171
    },
    [],
    "4132051975",
    {
        "_13": 1172,
        "_60": 61,
        "_62": 1174,
        "_64": 1175
    },
    "wLBwoUCuuMdnRwa9KkfHI",
    [],
    "4141006638",
    {
        "_13": 1176,
        "_60": 24,
        "_62": 1178,
        "_64": 1179
    },
    "6v4Q2eufBTFCb2P3fGZwPo",
    [],
    "4192239497",
    {
        "_13": 1180,
        "_60": 24,
        "_62": 91,
        "_64": 1182
    },
    [],
    "4206189746",
    {
        "_13": 1183,
        "_60": 24,
        "_62": 91,
        "_64": 1185
    },
    [],
    "4242210007",
    {
        "_13": 1186,
        "_60": 24,
        "_62": 1188,
        "_64": 1189
    },
    "5T7B6Qu0S7TF24HzOjoxJl",
    [],
    {
        "_1191": 1192,
        "_1199": 1200,
        "_365": 1204,
        "_1206": 1207,
        "_1225": 1226,
        "_1229": 1230,
        "_1235": 1236,
        "_1243": 1244,
        "_421": 1277,
        "_430": 1279,
        "_1281": 1282,
        "_1287": 1288
    },
    "550560761",
    {
        "_13": 1191,
        "_60": 1193,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1198
    },
    {
        "_1194": 1195,
        "_1196": 1197
    },
    "history_results_limit",
    6,
    "local_results_limit",
    2,
    [],
    "948081399",
    {
        "_13": 1199,
        "_60": 1201,
        "_356": 1202,
        "_62": 1202,
        "_358": 24,
        "_64": 1203,
        "_363": 24,
        "_364": 61
    },
    {},
    "layerAssignment",
    [],
    {
        "_13": 365,
        "_60": 367,
        "_356": 369,
        "_62": 369,
        "_358": 24,
        "_64": 1205,
        "_363": 61,
        "_364": 61
    },
    [],
    "1682643554",
    {
        "_13": 1206,
        "_60": 1208,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1224
    },
    {
        "_1209": 1210
    },
    "school_configurations",
    {
        "_1211": 1212,
        "_1220": 1221
    },
    "openai_1signup_for_1",
    {
        "_1213": 1214,
        "_1215": 1216,
        "_1217": 1218
    },
    "display_name",
    "OpenAI",
    "promotion_campaign_id",
    "students-2025-one-month-free",
    "domains",
    [
        1219
    ],
    "openai.com, mail.openai.com",
    "australia",
    {
        "_1213": 1214,
        "_1215": 1216,
        "_1217": 1222
    },
    [
        1223
    ],
    "edu.au",
    [],
    "1809520125",
    {
        "_13": 1225,
        "_60": 1227,
        "_356": 741,
        "_62": 741,
        "_358": 61,
        "_64": 1228,
        "_363": 24,
        "_364": 24
    },
    {},
    [],
    "2181185232",
    {
        "_13": 1229,
        "_60": 1231,
        "_356": 488,
        "_62": 488,
        "_358": 61,
        "_64": 1232,
        "_363": 24,
        "_364": 61
    },
    {},
    [
        1233
    ],
    {
        "_67": 1234,
        "_69": 101,
        "_71": 91
    },
    "1887864177",
    "2604379743",
    {
        "_13": 1235,
        "_60": 1237,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1242
    },
    {
        "_1238": 1239,
        "_1240": 1241
    },
    "nux_video_url",
    "https://persistent.oaistatic.com/image-gen/nux.CB3699EE.mov",
    "nux_image_url",
    "https://persistent.oaistatic.com/image-gen/nux.CB3699EE.jpg",
    [],
    "2821602598",
    {
        "_13": 1243,
        "_60": 1245,
        "_356": 91,
        "_62": 91,
        "_358": 61,
        "_64": 1276
    },
    {
        "_1246": 1247
    },
    "Football",
    [
        1248
    ],
    {
        "_1249": 1246,
        "_1250": 1251
    },
    "title",
    "templates",
    [
        1252,
        1262,
        1269
    ],
    {
        "_1253": 1254,
        "_1255": 1256
    },
    "text",
    "The [input] are down [input] with [input] left in the [input] quarter. What are their odds they [input]?",
    "suggestions",
    [
        1257,
        1258,
        1259,
        1260,
        1261
    ],
    "KC Chiefs",
    "27 - 10",
    "5 minutes",
    "4th",
    "win",
    {
        "_1253": 1263,
        "_1255": 1264
    },
    "I'm [input], played [input] and work out [input]. Help me train like a [input].",
    [
        1265,
        1266,
        1267,
        1268
    ],
    "29",
    "football",
    "3 days a week",
    "NFL running back",
    {
        "_1253": 1270,
        "_1255": 1271
    },
    "Write me a [input]. Include [input], [input], and [input].",
    [
        1272,
        1273,
        1274,
        1275
    ],
    "perfect halftime show song",
    "dancing",
    "fireworks",
    "being super fierce",
    [],
    {
        "_13": 421,
        "_60": 423,
        "_356": 424,
        "_62": 424,
        "_358": 24,
        "_64": 1278,
        "_363": 61,
        "_364": 61
    },
    [],
    {
        "_13": 430,
        "_60": 432,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1280
    },
    [],
    "3519108196",
    {
        "_13": 1281,
        "_60": 1283,
        "_356": 357,
        "_62": 357,
        "_358": 24,
        "_64": 1286,
        "_363": 24,
        "_364": 24
    },
    {
        "_1284": 61,
        "_1285": 61
    },
    "show-album-upload",
    "show-camera-upload",
    [],
    "3983984123",
    {
        "_13": 1287,
        "_60": 1289,
        "_356": 741,
        "_62": 741,
        "_358": 24,
        "_64": 1290,
        "_466": 1291,
        "_363": 24,
        "_364": 24,
        "_1292": 61
    },
    {
        "_749": 24
    },
    [],
    [
        749
    ],
    "is_in_layer",
    {
        "_1294": 1295,
        "_448": 1334,
        "_472": 1339,
        "_477": 1342,
        "_483": 1345,
        "_495": 1349,
        "_501": 1352,
        "_506": 1355,
        "_1364": 1365,
        "_1372": 1373,
        "_517": 1381,
        "_1385": 1386,
        "_1393": 1394,
        "_1402": 1403,
        "_1425": 1426,
        "_1435": 1436,
        "_558": 1455,
        "_571": 1458,
        "_577": 1461,
        "_583": 1464,
        "_592": 1467,
        "_1470": 1471,
        "_1479": 1480,
        "_1485": 1486,
        "_607": 1492,
        "_1496": 1497,
        "_626": 1502,
        "_643": 1507,
        "_665": 1512,
        "_675": 1516,
        "_1520": 1521,
        "_686": 1526,
        "_1530": 1531,
        "_1539": 1540,
        "_698": 1548,
        "_1551": 1552,
        "_1566": 1567,
        "_1572": 1573,
        "_1578": 1579,
        "_1587": 1588,
        "_1603": 1604,
        "_1610": 1611,
        "_1617": 1618,
        "_1623": 1624,
        "_1628": 1629,
        "_705": 1635,
        "_1642": 1643,
        "_1649": 1650,
        "_1655": 1656,
        "_727": 1673,
        "_737": 1681,
        "_746": 1684,
        "_1687": 1688,
        "_752": 1714,
        "_1717": 1718,
        "_1725": 1726,
        "_758": 1731,
        "_808": 1736,
        "_1739": 1740,
        "_1746": 1747,
        "_1761": 1762,
        "_815": 1766,
        "_1769": 1770,
        "_1774": 1775,
        "_822": 1804,
        "_1807": 1808
    },
    "109457",
    {
        "_13": 1294,
        "_60": 1296,
        "_356": 1324,
        "_62": 1324,
        "_358": 24,
        "_64": 1325,
        "_466": 1331,
        "_468": 1332,
        "_364": 61,
        "_363": 61,
        "_470": 1333
    },
    {
        "_1297": 24,
        "_1298": 24,
        "_1299": 24,
        "_1300": 24,
        "_1301": 24,
        "_1302": 455,
        "_1303": 24,
        "_1304": 24,
        "_1305": 24,
        "_1306": 455,
        "_1307": 24,
        "_1308": 1309,
        "_1310": 24,
        "_1311": 24,
        "_1312": 24,
        "_1313": 24,
        "_1314": 24,
        "_1315": 455,
        "_1316": 24,
        "_1317": 1318,
        "_1319": 1320,
        "_1321": 1320,
        "_1322": 24,
        "_1323": 24
    },
    "is_starter_prompt_popular",
    "is_starter_prompt_top_performer",
    "is_starter_prompt_back_and_forth",
    "use_starter_prompt_help_how_to",
    "model_talks_first",
    "model_talks_first_kind",
    "model_talks_first_augment_system_prompt",
    "is_starter_prompt_enabled_for_new_users_only",
    "add_system_prompt_during_onboarding",
    "onboarding_system_prompt_type",
    "enable_new_onboarding_flow",
    "new_onboarding_flow_qualified_start_date",
    "2025-02-28T06:00:00Z",
    "personalized_onboarding",
    "onboarding_show_custom_instructions_page",
    "write_custom_instructions_in_onboarding",
    "keep_onboarding_after_dismiss",
    "onboarding_dynamic_steps_based_on_main_usage",
    "onboarding_style",
    "onboarding_show_followups",
    "onboarding_inject_cards_position",
    9999,
    "ONBOARDING_EXAMPLES_PROMPT_ID",
    "convo_gen_examples_v2",
    "onboarding_gen_examples_prompt_type",
    "show_new_chat_nux",
    "is_guided_onboarding",
    "M3EE4Hyw83Rv7RjIICK6o",
    [
        1326,
        1328
    ],
    {
        "_67": 1327,
        "_69": 101,
        "_71": 91
    },
    "674041001",
    {
        "_67": 1329,
        "_69": 70,
        "_71": 1330
    },
    "59687878",
    "4k3eNmHeryixdsgalKqv0",
    [
        1315,
        1308,
        1305,
        1306,
        1310,
        1316,
        1317,
        1321,
        1322
    ],
    "52701554",
    [
        1326
    ],
    {
        "_13": 448,
        "_60": 450,
        "_356": 461,
        "_62": 461,
        "_358": 24,
        "_64": 1335,
        "_466": 467,
        "_468": 469,
        "_364": 24,
        "_363": 24,
        "_470": 1338
    },
    [
        1336
    ],
    {
        "_67": 464,
        "_69": 70,
        "_71": 1337
    },
    "1yBehWRiofl3CcNtvNVvk6",
    [
        1336
    ],
    {
        "_13": 472,
        "_60": 474,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1340,
        "_466": 1341,
        "_470": 1340
    },
    [],
    [],
    {
        "_13": 477,
        "_60": 479,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1343,
        "_466": 1344,
        "_470": 1343
    },
    [],
    [],
    {
        "_13": 483,
        "_60": 485,
        "_356": 488,
        "_62": 488,
        "_358": 61,
        "_64": 1346,
        "_466": 492,
        "_468": 493,
        "_364": 61,
        "_363": 24,
        "_470": 1348
    },
    [
        1347
    ],
    {
        "_67": 491,
        "_69": 101,
        "_71": 91
    },
    [],
    {
        "_13": 495,
        "_60": 497,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1350,
        "_466": 1351,
        "_470": 1350
    },
    [],
    [],
    {
        "_13": 501,
        "_60": 503,
        "_356": 91,
        "_62": 91,
        "_358": 61,
        "_64": 1353,
        "_466": 1354,
        "_470": 1353
    },
    [],
    [],
    {
        "_13": 506,
        "_60": 1356,
        "_356": 1357,
        "_62": 1357,
        "_358": 61,
        "_64": 1358,
        "_466": 1362,
        "_468": 506,
        "_364": 24,
        "_363": 24,
        "_470": 1363
    },
    {
        "_509": 61,
        "_510": 61
    },
    "5UE8g4T56yxUBUYancL7KB:override",
    [
        1359,
        1360
    ],
    {
        "_67": 97,
        "_69": 101,
        "_71": 91
    },
    {
        "_67": 100,
        "_69": 70,
        "_71": 1361
    },
    "5hCRKi4Gs5QJkOanmdVvHU:100.00:4",
    [
        510,
        509
    ],
    [
        1359,
        1360
    ],
    "415386882",
    {
        "_13": 1364,
        "_60": 1366,
        "_356": 91,
        "_62": 91,
        "_358": 61,
        "_64": 1368,
        "_466": 1371,
        "_470": 1368
    },
    {
        "_1367": 24
    },
    "is_voice_mode_entry_point_enabled",
    [
        1369
    ],
    {
        "_67": 1370,
        "_69": 101,
        "_71": 91
    },
    "1644396868",
    [],
    "453021389",
    {
        "_13": 1372,
        "_60": 1374,
        "_356": 91,
        "_62": 91,
        "_358": 61,
        "_64": 1377,
        "_466": 1380,
        "_470": 1377
    },
    {
        "_1375": 24,
        "_1376": 61
    },
    "enable-block-animations",
    "enable-word-animations",
    [
        1378
    ],
    {
        "_67": 1379,
        "_69": 101,
        "_71": 455
    },
    "3016192915",
    [],
    {
        "_13": 517,
        "_60": 519,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1382,
        "_466": 1384,
        "_470": 1382
    },
    [
        1383
    ],
    {
        "_67": 549,
        "_69": 101,
        "_71": 91
    },
    [],
    "474444727",
    {
        "_13": 1385,
        "_60": 1387,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1391,
        "_466": 1392,
        "_470": 1391
    },
    {
        "_1388": 61,
        "_1389": 1390
    },
    "show_custom_instr_message",
    "custom_instr_message_timeout_duration",
    1500,
    [],
    [],
    "590557768",
    {
        "_13": 1393,
        "_60": 1395,
        "_356": 1397,
        "_62": 1397,
        "_358": 61,
        "_64": 1398,
        "_466": 1399,
        "_468": 1400,
        "_364": 61,
        "_363": 61,
        "_470": 1401
    },
    {
        "_1396": 61
    },
    "should_show_return_home_btn",
    "MfvDyM5oEZ1TqWS7cE8et",
    [],
    [
        1396
    ],
    "1022536663",
    [],
    "660512088",
    {
        "_13": 1402,
        "_60": 1404,
        "_356": 91,
        "_62": 91,
        "_358": 61,
        "_64": 1412,
        "_466": 1424,
        "_470": 1412
    },
    {
        "_1405": 24,
        "_1406": 61,
        "_1407": 24,
        "_1408": 24,
        "_1409": 24,
        "_1410": 24,
        "_1411": 24
    },
    "enable_arch_updates",
    "include_legacy_sidebar_contents",
    "include_floating_state",
    "include_share_on_mobile",
    "include_account_settings_move",
    "include_scrolling_behavior_update",
    "include_revised_sidebar_ia",
    [
        1413,
        1415,
        1418,
        1421
    ],
    {
        "_67": 1414,
        "_69": 101,
        "_71": 91
    },
    "2558701922",
    {
        "_67": 1416,
        "_69": 101,
        "_71": 1417
    },
    "735930678",
    "6nGV45RQYtcIGTbPzppBhS",
    {
        "_67": 1419,
        "_69": 101,
        "_71": 1420
    },
    "3011415004",
    "7pUMK6uci7sslAj8bP7VEA",
    {
        "_67": 1422,
        "_69": 101,
        "_71": 1423
    },
    "854062205",
    "66y6sNojVqOdoNf0CX0JYC",
    [],
    "685344542",
    {
        "_13": 1425,
        "_60": 1427,
        "_356": 1429,
        "_62": 1429,
        "_358": 24,
        "_64": 1430,
        "_466": 1432,
        "_468": 1433,
        "_364": 24,
        "_363": 24,
        "_470": 1434
    },
    {
        "_1428": 24,
        "_529": 61
    },
    "is_mobile_enterprise_enabled",
    "3INu3qkV6QoN42TYoP3gja:override",
    [
        1431
    ],
    {
        "_67": 132,
        "_69": 70,
        "_71": 134
    },
    [
        529
    ],
    "1388643772",
    [
        1431
    ],
    "717266490",
    {
        "_13": 1435,
        "_60": 1437,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1453,
        "_466": 1454,
        "_470": 1453
    },
    {
        "_1438": 61,
        "_1439": 61,
        "_1440": 61,
        "_1308": 1441,
        "_1307": 24,
        "_1442": 24,
        "_1310": 24,
        "_1313": 24,
        "_1312": 24,
        "_1443": 406,
        "_1444": 24,
        "_1311": 24,
        "_1445": 24,
        "_1446": 61,
        "_1447": 24,
        "_1448": 1449
    },
    "optimize_initial_modals",
    "defer_memory_modal",
    "enable_v2_cleanup",
    "2099-11-04T00:00:00Z",
    "use_plus_rl_during_onboarding",
    "plus_rl_during_onboarding_minutes_after_creation",
    "enable_mobile_app_upsell_banner",
    "one_tooltip_per_session",
    "one_announcement_tooltip_per_session",
    "onboarding_show_other_option",
    "onboarding_flow_tool_steps",
    [
        1450,
        1451,
        1452
    ],
    "dalle",
    "file_upload",
    "canvas",
    [],
    [],
    {
        "_13": 558,
        "_60": 560,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1456,
        "_466": 1457,
        "_470": 1456
    },
    [],
    [],
    {
        "_13": 571,
        "_60": 573,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1459,
        "_466": 1460,
        "_470": 1459
    },
    [],
    [],
    {
        "_13": 577,
        "_60": 579,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1462,
        "_466": 1463,
        "_470": 1462
    },
    [],
    [],
    {
        "_13": 583,
        "_60": 585,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1465,
        "_466": 1466,
        "_470": 1465
    },
    [],
    [],
    {
        "_13": 592,
        "_60": 594,
        "_356": 91,
        "_62": 91,
        "_358": 61,
        "_64": 1468,
        "_466": 1469,
        "_470": 1468
    },
    [],
    [],
    "1358188185",
    {
        "_13": 1470,
        "_60": 1472,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1475,
        "_466": 1478,
        "_470": 1475
    },
    {
        "_1473": 61,
        "_1474": 24
    },
    "prefetch-models",
    "sidebar-default-close",
    [
        1476
    ],
    {
        "_67": 1477,
        "_69": 101,
        "_71": 91
    },
    "542939804",
    [],
    "1358849452",
    {
        "_13": 1479,
        "_60": 1481,
        "_356": 91,
        "_62": 91,
        "_358": 61,
        "_64": 1483,
        "_466": 1484,
        "_470": 1483
    },
    {
        "_1482": 24
    },
    "disable-ssr",
    [],
    [],
    "1368081792",
    {
        "_13": 1485,
        "_60": 1487,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1490,
        "_466": 1491,
        "_470": 1490
    },
    {
        "_1488": 24,
        "_1489": 24
    },
    "should_show_o3_mini_high_upsell_banner_free_user_to_plus",
    "should_show_o3_mini_high_upsell_banner_plus_user",
    [],
    [],
    {
        "_13": 607,
        "_60": 609,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1493,
        "_466": 1495,
        "_470": 1493
    },
    [
        1494
    ],
    {
        "_67": 624,
        "_69": 101,
        "_71": 91
    },
    [],
    "1578749296",
    {
        "_13": 1496,
        "_60": 1498,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1500,
        "_466": 1501,
        "_470": 1500
    },
    {
        "_1499": 24
    },
    "is_sticky_toggle_off",
    [],
    [],
    {
        "_13": 626,
        "_60": 1503,
        "_356": 488,
        "_62": 488,
        "_358": 24,
        "_64": 1504,
        "_466": 640,
        "_468": 641,
        "_364": 61,
        "_363": 24,
        "_470": 1506
    },
    {
        "_629": 24,
        "_630": 24,
        "_631": 24,
        "_632": 24,
        "_633": 24,
        "_634": 24
    },
    [
        1505
    ],
    {
        "_67": 638,
        "_69": 101,
        "_71": 91
    },
    [],
    {
        "_13": 643,
        "_60": 645,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1508,
        "_466": 1511,
        "_470": 1508
    },
    [
        1509,
        1510
    ],
    {
        "_67": 650,
        "_69": 101,
        "_71": 651
    },
    {
        "_67": 653,
        "_69": 101,
        "_71": 654
    },
    [],
    {
        "_13": 665,
        "_60": 667,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1513,
        "_466": 1515,
        "_470": 1513
    },
    [
        1514
    ],
    {
        "_67": 673,
        "_69": 101,
        "_71": 91
    },
    [],
    {
        "_13": 675,
        "_60": 677,
        "_356": 488,
        "_62": 488,
        "_358": 24,
        "_64": 1517,
        "_466": 683,
        "_468": 684,
        "_364": 61,
        "_363": 24,
        "_470": 1519
    },
    [
        1518
    ],
    {
        "_67": 682,
        "_69": 101,
        "_71": 91
    },
    [],
    "1846737571",
    {
        "_13": 1520,
        "_60": 1522,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1524,
        "_466": 1525,
        "_470": 1524
    },
    {
        "_1523": 24
    },
    "is_upgrade_button_blue",
    [],
    [],
    {
        "_13": 686,
        "_60": 688,
        "_356": 690,
        "_62": 690,
        "_358": 61,
        "_64": 1527,
        "_466": 695,
        "_468": 696,
        "_364": 24,
        "_363": 24,
        "_470": 1529
    },
    [
        1528
    ],
    {
        "_67": 693,
        "_69": 70,
        "_71": 694
    },
    [
        1528
    ],
    "2118136551",
    {
        "_13": 1530,
        "_60": 1532,
        "_356": 91,
        "_62": 91,
        "_358": 61,
        "_64": 1537,
        "_466": 1538,
        "_470": 1537
    },
    {
        "_1533": 24,
        "_1534": 24,
        "_1535": 61,
        "_1536": 61
    },
    "show_cookie_banner_if_qualified",
    "test_dummy",
    "sign_up_button_has_the_word_free",
    "show_cookie_banner_auth_login",
    [],
    [],
    "2149763392",
    {
        "_13": 1539,
        "_60": 1541,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1544,
        "_466": 1547,
        "_470": 1544
    },
    {
        "_1542": 24,
        "_1543": 24
    },
    "show-in-main-composer",
    "show-model-picker",
    [
        1545
    ],
    {
        "_67": 1546,
        "_69": 101,
        "_71": 91
    },
    "4151101559",
    [],
    {
        "_13": 698,
        "_60": 700,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1549,
        "_466": 1550,
        "_470": 1549
    },
    [],
    [],
    "2259187367",
    {
        "_13": 1551,
        "_60": 1553,
        "_356": 91,
        "_62": 91,
        "_358": 61,
        "_64": 1564,
        "_466": 1565,
        "_470": 1564
    },
    {
        "_1554": 24,
        "_1555": 1556,
        "_1557": 1558,
        "_1559": 61,
        "_1560": 1561,
        "_1562": 24,
        "_1563": 1246
    },
    "enable_nux",
    "start_time",
    "2099-01-01T00:00:00Z",
    "end_time",
    "2000-01-01T00:00:00Z",
    "use_multi_input",
    "force_madlibs_param_name",
    "madlibs_0203",
    "enable_additional_categories",
    "additional_category",
    [],
    [],
    "2505516353",
    {
        "_13": 1566,
        "_60": 1568,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1570,
        "_466": 1571,
        "_470": 1570
    },
    {
        "_1569": 61
    },
    "android-keyboard-layout",
    [],
    [],
    "2670443078",
    {
        "_13": 1572,
        "_60": 1574,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1576,
        "_466": 1577,
        "_470": 1576
    },
    {
        "_1575": 61
    },
    "is_gating_fix_enabled",
    [],
    [],
    "2716194794",
    {
        "_13": 1578,
        "_60": 1580,
        "_356": 488,
        "_62": 488,
        "_358": 24,
        "_64": 1581,
        "_466": 1584,
        "_468": 1585,
        "_364": 61,
        "_363": 24,
        "_470": 1586
    },
    {
        "_740": 24
    },
    [
        1582
    ],
    {
        "_67": 1583,
        "_69": 101,
        "_71": 91
    },
    "2849926832",
    [
        740
    ],
    "2435265903",
    [],
    "2723963139",
    {
        "_13": 1587,
        "_60": 1589,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1601,
        "_466": 1602,
        "_470": 1601
    },
    {
        "_1590": 24,
        "_1591": 24,
        "_1592": 61,
        "_1593": 61,
        "_1594": 61,
        "_1595": 1596,
        "_1597": 61,
        "_1598": 24,
        "_1599": 24,
        "_1600": 455
    },
    "is_dynamic_model_enabled",
    "show_message_model_info",
    "show_message_regenerate_model_selector",
    "is_conversation_model_switching_allowed",
    "show_rate_limit_downgrade_banner",
    "config",
    {},
    "show_message_regenerate_model_selector_on_every_message",
    "is_AG8PqS2q_enabled",
    "is_chive_enabled",
    "sahara_model_id_override",
    [],
    [],
    "2775247110",
    {
        "_13": 1603,
        "_60": 1605,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1608,
        "_466": 1609,
        "_470": 1608
    },
    {
        "_1606": 24,
        "_1607": 24
    },
    "show_pro_badge",
    "show_plan_type_badge",
    [],
    [],
    "2840731323",
    {
        "_13": 1610,
        "_60": 1612,
        "_356": 91,
        "_62": 91,
        "_358": 61,
        "_64": 1614,
        "_466": 1616,
        "_470": 1614
    },
    {
        "_616": 61,
        "_1613": 61
    },
    "is_direct_continue_enabled",
    [
        1615
    ],
    {
        "_67": 1021,
        "_69": 101,
        "_71": 91
    },
    [],
    "2888142241",
    {
        "_13": 1617,
        "_60": 1619,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1621,
        "_466": 1622,
        "_470": 1621
    },
    {
        "_1620": 61
    },
    "is_upgrade_in_settings",
    [],
    [],
    "2932223118",
    {
        "_13": 1623,
        "_60": 1625,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1626,
        "_466": 1627,
        "_470": 1626
    },
    {
        "_528": 61
    },
    [],
    [],
    "2972011003",
    {
        "_13": 1628,
        "_60": 1630,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1633,
        "_466": 1634,
        "_470": 1633
    },
    {
        "_1631": 61,
        "_1632": 24
    },
    "user_context_message_search_tools_default",
    "search_tool_holdout_enabled",
    [],
    [],
    {
        "_13": 705,
        "_60": 1636,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1637,
        "_466": 1641,
        "_470": 1637
    },
    {
        "_708": 61,
        "_709": 24
    },
    [
        1638,
        1639
    ],
    {
        "_67": 713,
        "_69": 101,
        "_71": 714
    },
    {
        "_67": 716,
        "_69": 101,
        "_71": 1640
    },
    "66covmutTYx82FWVUlZAqF",
    [],
    "3119715334",
    {
        "_13": 1642,
        "_60": 1644,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1647,
        "_466": 1648,
        "_470": 1647
    },
    {
        "_1645": 24,
        "_1646": 24
    },
    "should-enable-hojicha",
    "should-enable-skip",
    [],
    [],
    "3206655705",
    {
        "_13": 1649,
        "_60": 1651,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1653,
        "_466": 1654,
        "_470": 1653
    },
    {
        "_1652": 61
    },
    "enable_new_ux",
    [],
    [],
    "3434623093",
    {
        "_13": 1655,
        "_60": 1657,
        "_356": 91,
        "_62": 91,
        "_358": 61,
        "_64": 1663,
        "_466": 1672,
        "_470": 1663
    },
    {
        "_1658": 61,
        "_1659": 1660,
        "_1661": 61,
        "_1662": 61
    },
    "with-attach-upsell",
    "labels",
    "all",
    "with-voice-upsell",
    "with-reason-upsell",
    [
        1664,
        1666,
        1668,
        1670
    ],
    {
        "_67": 1665,
        "_69": 101,
        "_71": 91
    },
    "1604099973",
    {
        "_67": 1667,
        "_69": 101,
        "_71": 91
    },
    "470066910",
    {
        "_67": 1669,
        "_69": 101,
        "_71": 91
    },
    "1932133792",
    {
        "_67": 1671,
        "_69": 101,
        "_71": 91
    },
    "4175621034",
    [],
    {
        "_13": 727,
        "_60": 1674,
        "_356": 1676,
        "_62": 1676,
        "_358": 24,
        "_64": 1677,
        "_466": 734,
        "_468": 735,
        "_364": 61,
        "_363": 61,
        "_470": 1680
    },
    {
        "_629": 61,
        "_730": 1675,
        "_632": 61,
        "_631": 61,
        "_630": 24
    },
    10,
    "2FurbaJwwLPFodZHhOZyBO",
    [
        1678
    ],
    {
        "_67": 733,
        "_69": 70,
        "_71": 1679
    },
    "1FzsKf0T7jWwTRKiSrbUld:100.00:4",
    [],
    {
        "_13": 737,
        "_60": 739,
        "_356": 741,
        "_62": 741,
        "_358": 61,
        "_64": 1682,
        "_466": 743,
        "_468": 744,
        "_364": 24,
        "_363": 24,
        "_470": 1683
    },
    [],
    [],
    {
        "_13": 746,
        "_60": 748,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1685,
        "_466": 1686,
        "_470": 1685
    },
    [],
    [],
    "3533083032",
    {
        "_13": 1687,
        "_60": 1689,
        "_356": 91,
        "_62": 91,
        "_358": 61,
        "_64": 1712,
        "_466": 1713,
        "_470": 1712
    },
    {
        "_1690": 61,
        "_1691": 61,
        "_1692": 1693,
        "_1694": 24,
        "_1695": 24,
        "_1696": 61,
        "_1697": 24,
        "_1698": 24,
        "_1699": 24,
        "_1700": 24,
        "_1701": 1702,
        "_1703": 1704,
        "_1705": 1706,
        "_1707": 1708,
        "_1709": 1710,
        "_1711": 455
    },
    "enable_new_homepage_anon",
    "filter_prompt_by_model",
    "headline_option",
    "HELP_WITH",
    "disclaimer_color_adjust",
    "show_composer_header",
    "enable_new_mobile",
    "enable_cached_response",
    "show_dalle_starter_prompts",
    "use_modapi_in_autocomplete",
    "use_memory_in_model_autocomplete",
    "autocomplete_max_char",
    32,
    "search_autocomplete_mode",
    "BING",
    "autocomplete_min_char",
    4,
    "autocomplete_mode",
    "INDEX",
    "num_completions_to_fetch_from_index",
    8,
    "india_first_prompt",
    [],
    [],
    {
        "_13": 752,
        "_60": 754,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1715,
        "_466": 1716,
        "_470": 1715
    },
    [],
    [],
    "3606233934",
    {
        "_13": 1717,
        "_60": 1719,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1723,
        "_466": 1724,
        "_470": 1723
    },
    {
        "_1720": 1721,
        "_1722": 24
    },
    "link",
    "non",
    "enable_notifications_feed",
    [],
    [],
    "3613709240",
    {
        "_13": 1725,
        "_60": 1727,
        "_356": 91,
        "_62": 91,
        "_358": 61,
        "_64": 1729,
        "_466": 1730,
        "_470": 1729
    },
    {
        "_1728": 61
    },
    "shouldRefreshAccessToken",
    [],
    [],
    {
        "_13": 758,
        "_60": 760,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1732,
        "_466": 1735,
        "_470": 1732
    },
    [
        1733,
        1734
    ],
    {
        "_67": 788,
        "_69": 101,
        "_71": 149
    },
    {
        "_67": 790,
        "_69": 101,
        "_71": 149
    },
    [],
    {
        "_13": 808,
        "_60": 810,
        "_356": 91,
        "_62": 91,
        "_358": 61,
        "_64": 1737,
        "_466": 1738,
        "_470": 1737
    },
    [],
    [],
    "3737571708",
    {
        "_13": 1739,
        "_60": 1741,
        "_356": 91,
        "_62": 91,
        "_358": 61,
        "_64": 1744,
        "_466": 1745,
        "_470": 1744
    },
    {
        "_1742": 1743
    },
    "sidebar_type",
    "slick",
    [],
    [],
    "3768341700",
    {
        "_13": 1746,
        "_60": 1748,
        "_356": 1756,
        "_62": 1756,
        "_358": 24,
        "_64": 1757,
        "_466": 1758,
        "_468": 1759,
        "_364": 61,
        "_363": 61,
        "_470": 1760
    },
    {
        "_530": 24,
        "_1749": 24,
        "_1750": 24,
        "_1751": 61,
        "_1752": 61,
        "_1753": 61,
        "_1754": 24,
        "_1755": 24
    },
    "remove_early_access_upsell",
    "is_produce_text_design",
    "is_produce_design",
    "is_country_selector_enabled",
    "is_vat_information_enabled",
    "is_vat_information_with_amount_enabled",
    "is_team_pricing_vat_disclaimer_enabled",
    "65VHFqyIytQJKjgykJm4UQ",
    [],
    [
        1753,
        1752,
        1754,
        1755
    ],
    "2782616826",
    [],
    "3927927759",
    {
        "_13": 1761,
        "_60": 1763,
        "_356": 91,
        "_62": 91,
        "_358": 61,
        "_64": 1764,
        "_466": 1765,
        "_470": 1764
    },
    {
        "_1444": 61
    },
    [],
    [],
    {
        "_13": 815,
        "_60": 817,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1767,
        "_466": 1768,
        "_470": 1767
    },
    [],
    [],
    "4020668365",
    {
        "_13": 1769,
        "_60": 1771,
        "_356": 91,
        "_62": 91,
        "_358": 61,
        "_64": 1772,
        "_466": 1773,
        "_470": 1772
    },
    {
        "_1554": 24,
        "_1555": 1556,
        "_1557": 1558,
        "_1559": 24
    },
    [],
    [],
    "4031588851",
    {
        "_13": 1774,
        "_60": 1776,
        "_356": 91,
        "_62": 91,
        "_358": 24,
        "_64": 1799,
        "_466": 1803,
        "_470": 1799
    },
    {
        "_1777": 61,
        "_1778": 61,
        "_1779": 61,
        "_1780": 61,
        "_1781": 24,
        "_1782": 24,
        "_1707": 1708,
        "_1783": 1784,
        "_1705": 1706,
        "_1701": 1702,
        "_1692": 1693,
        "_1700": 24,
        "_1785": 24,
        "_1699": 24,
        "_1786": 1787,
        "_1788": 61,
        "_1789": 455,
        "_1696": 61,
        "_1703": 1704,
        "_1790": 24,
        "_1791": 1792,
        "_1709": 1710,
        "_1793": 24,
        "_1794": 24,
        "_780": 781,
        "_1795": 24,
        "_1796": 1797,
        "_1798": 61,
        "_1711": 455
    },
    "enable_hardcoded_vision_prompts",
    "enable_hardcoded_file_document_prompts",
    "enable_hardcoded_data_vis_prompts",
    "enable_hardcoded_browse_prompts",
    "is_two_line",
    "enable_new_homepage",
    "starter_prompt_ranking_algorithm",
    "homepage_v2",
    "filter_starter_prompt_by_model",
    "autocomplete_qualified_start_date",
    "2000-10-11T00:00:00Z",
    "enable_new_autocomplete_homepage",
    "model_talks_option",
    "enable_hardcoded_onboarding_prompt",
    "autocomplete_fetch_interval",
    200,
    "enable_recommend_prompts",
    "enable_ask_me_prompts",
    "enable_reasoning_prompts_0202",
    "dream_type",
    "user_knowledge_memories",
    "web-disable",
    [
        1800
    ],
    {
        "_67": 1801,
        "_69": 101,
        "_71": 1802
    },
    "4273941502",
    "1nGrz4l6GM0LgZvm0pDCtp:2.00:1",
    [],
    {
        "_13": 822,
        "_60": 824,
        "_356": 91,
        "_62": 91,
        "_358": 61,
        "_64": 1805,
        "_466": 1806,
        "_470": 1805
    },
    [],
    [],
    "4250072504",
    {
        "_13": 1807,
        "_60": 1809,
        "_356": 1812,
        "_62": 1812,
        "_358": 24,
        "_64": 1813,
        "_466": 1815,
        "_468": 1816,
        "_364": 24,
        "_363": 24,
        "_470": 1817
    },
    {
        "_529": 61,
        "_1810": 24,
        "_1811": 24
    },
    "is_enterprise_desktop_enabled",
    "is_desktop_enterprise_enabled",
    "3HX7vpdJsUkuyCUEL4V9cE:override",
    [
        1814
    ],
    {
        "_67": 132,
        "_69": 70,
        "_71": 134
    },
    [
        529
    ],
    "3311396813",
    [
        1814
    ],
    {},
    {
        "_834": 835,
        "_836": 837
    },
    {
        "_842": 12,
        "_843": 1821
    },
    {
        "_845": 846,
        "_847": 846,
        "_848": 846
    },
    {
        "_842": 12,
        "_854": 855,
        "_856": 1823,
        "_843": 1821,
        "_1828": 1829,
        "_1830": 1826,
        "_861": 1831,
        "_50": 51
    },
    {
        "_1824": 61,
        "_1825": 1826,
        "_1827": 24,
        "_859": 24,
        "_860": 6
    },
    "has_logged_in_before",
    "user_agent",
    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36 Edg/134.0.0.0",
    "is_punch_out_user",
    "ip",
    "2a0c:9a40:9b07:1f5d:a5e4:5e46:e469:469d",
    "userAgent",
    {
        "_863": 864
    },
    "isNoAuthEnabled",
    "userRegion",
    "New York",
    "US",
    "cfIpLatitude",
    null,
    "cfIpLongitude",
    null,
    "cfIpCity",
    null,
    "isUserInNewCookieConsentFlow",
    "isUserInPioneerHR",
    "isUserEligibleForPioneer",
    "isUserEligibleForMaverick",
    "kind",
    "chat_page",
    "gizmo",
    {},
    {
        "_11": 1851,
        "_1852": 1853,
        "_1854": 1855,
        "_1856": 1857,
        "_1869": 1870,
        "_849": 1872,
        "_1873": -5,
        "_1874": -5,
        "_1875": 1876,
        "_1895": 1896,
        "_1897": 1898,
        "_1899": 1900,
        "_1901": -5,
        "_1902": -5,
        "_1903": 1904,
        "_1906": 24,
        "_1907": -5,
        "_1908": -5,
        "_1909": 1910,
        "_1911": 1912,
        "_1913": -5,
        "_1914": -5,
        "_1915": -5,
        "_1916": -5,
        "_1917": 1918,
        "_1930": -5,
        "_1931": -5,
        "_1932": -5,
        "_1933": -5,
        "_1934": 1935,
        "_1936": 1937
    },
    "g-rmdbtMF7a",
    "organization_id",
    "org-KxJjYXqyxZonde7LipjtEvt9",
    "short_url",
    "g-rmdbtMF7a-expedia",
    "author",
    {
        "_1858": 1859,
        "_1213": 1860,
        "_1861": 1862,
        "_1863": 61,
        "_1864": 1865,
        "_1866": -5,
        "_1867": 1868
    },
    "user_id",
    "user-rkxOl7npYCEOiqtG3Ag8jYbX__3372cc61-7f4e-4dc8-add6-518c4694a5fd",
    "expedia.com",
    "link_to",
    "https://expedia.com",
    "is_verified",
    "selected_display",
    "website",
    "will_receive_support_emails",
    "display_socials",
    [],
    "voice",
    {
        "_11": 1871
    },
    "ember",
    "3372cc61-7f4e-4dc8-add6-518c4694a5fd",
    "model",
    "instructions",
    "display",
    {
        "_13": 1877,
        "_1878": 1879,
        "_1880": 1881,
        "_1886": 1887,
        "_1888": 1889,
        "_1890": 1891,
        "_1893": -5,
        "_1894": -5
    },
    "Expedia",
    "description",
    "Bring your trip plans to life \u2013 get there, stay there, find things to see and do.",
    "prompt_starters",
    [
        1882,
        1883,
        1884,
        1885
    ],
    "We\u2019re looking for a hotel near the Eiffel Tower with a gym and spa for our stay in Paris for a week",
    "I'm looking for a nonstop flight from Toronto to New York City for next week?",
    "Find me food and drink tours in Paris that are 4 hours or less during next week.",
    "I need to rent a car in Seattle for a week. What are my options?",
    "profile_pic_id",
    "file-WtxsnVLKLzSkVnB3g5v1bN",
    "profile_picture_url",
    "https://chatgpt.com/backend-api/content?id=file-WtxsnVLKLzSkVnB3g5v1bN&gizmo_id=g-rmdbtMF7a&ts=484163&p=gpp&sig=d84522da7becc2d5e7058ee51b2c4f2df79857ee914bb971f519e820fd994681&v=0",
    "categories",
    [
        1892
    ],
    "lifestyle",
    "emoji",
    "theme",
    "share_recipient",
    "marketplace",
    "created_at",
    "2024-10-28T06:38:50.794842+00:00",
    "updated_at",
    "2025-03-26T11:01:52.860781+00:00",
    "last_interacted_at",
    "num_interactions",
    "tags",
    [
        1905
    ],
    "interactions_disabled",
    "is_unassigned",
    "version",
    "version_author",
    "version_created_at",
    "2025-02-12T21:25:14.012358+00:00",
    "version_updated_at",
    "2025-02-25T19:30:41.116315+00:00",
    "live_version",
    "training_disabled",
    "sharing_targets",
    "appeal_info",
    "vanity_metrics",
    {
        "_1919": -5,
        "_1920": 1921,
        "_1922": 1923,
        "_1924": 1925
    },
    "num_conversations",
    "num_conversations_str",
    "200K+",
    "created_ago_str",
    "4 months ago",
    "review_stats",
    {
        "_1926": 1927,
        "_1928": 1929
    },
    "total",
    8716,
    "count",
    2318,
    "workspace_approval_date",
    "workspace_approved",
    "sharing",
    "current_user_permission",
    "gizmo_type",
    "gpt",
    "context_stuffing_budget",
    16384,
    "tools",
    [
        1940,
        1946
    ],
    {
        "_11": 1941,
        "_1942": 1943,
        "_1944": -5,
        "_1945": -5
    },
    "673622cc3b548190b12b63ab2ed8780c",
    "type",
    "browser",
    "settings",
    "metadata",
    {
        "_11": 1947,
        "_1942": 1948,
        "_1944": -5,
        "_1945": 1949
    },
    "67ad11ba038081919bfe52e301ecc1aa",
    "plugins_prototype",
    {
        "_1950": 1951,
        "_1952": 1953,
        "_1954": -5,
        "_1955": 1956,
        "_2483": 2484,
        "_2491": 2492
    },
    "action_id",
    "g-1f13973d98417b91b85b9e916976c75700deeb99",
    "domain",
    "apim.expedia.com",
    "raw_spec",
    "json_schema",
    {
        "_1957": 1958,
        "_1959": 1960,
        "_1964": 1965,
        "_1969": 1970,
        "_2285": 2286
    },
    "openapi",
    "3.1.0",
    "info",
    {
        "_1249": 1961,
        "_1878": 1962,
        "_1907": 1963
    },
    "Expedia Travel Recommendation Service",
    "A service that allows to search travel products by search query and returns recommended products.",
    "v1",
    "servers",
    [
        1966
    ],
    {
        "_1967": 1968
    },
    "url",
    "https://apim.expedia.com",
    "paths",
    {
        "_1971": 1972,
        "_2113": 2114,
        "_2159": 2160,
        "_2222": 2223
    },
    "/recommendations/hotels",
    {
        "_1973": 1974
    },
    "get",
    {
        "_1975": 1976,
        "_1977": 1978,
        "_1979": 1980,
        "_2101": 2102
    },
    "operationId",
    "lodgingProducts",
    "summary",
    "API provides the top 3 recommended lodging travel products in the given destination.",
    "parameters",
    [
        1981,
        1990,
        1994,
        1998,
        2004,
        2008,
        2021,
        2026,
        2030,
        2073,
        2081,
        2089,
        2097
    ],
    {
        "_13": 1982,
        "_1983": 1984,
        "_1878": 1985,
        "_1986": 61,
        "_1987": 1988
    },
    "user_input_in_english",
    "in",
    "query",
    "(MANDATORY parameter) The merged user inputs according to the conversation context, translate to ENGLISH, DO NOT ONLY return the last query.\n",
    "required",
    "schema",
    {
        "_1942": 1989
    },
    "string",
    {
        "_13": 1991,
        "_1983": 1984,
        "_1878": 1992,
        "_1986": 61,
        "_1987": 1993
    },
    "keywords",
    "(MANDATORY parameter) Keywords associated with TRAVEL in the user_input_in_english parameter, converted to all LOWERCASE, separated by | character.\n",
    {
        "_1942": 1989
    },
    {
        "_13": 1995,
        "_1983": 1984,
        "_1878": 1996,
        "_1986": 61,
        "_1987": 1997
    },
    "destination",
    "(MANDATORY parameter) The destination can be a city, address, airport or a landmark.",
    {
        "_1942": 1989
    },
    {
        "_13": 1999,
        "_1983": 1984,
        "_1986": 24,
        "_1878": 2000,
        "_1987": 2001
    },
    "check_in",
    "(OPTIONAL parameter) Accept any date format and convert to YYYY-MM-DD.",
    {
        "_1942": 1989,
        "_2002": 2003
    },
    "example",
    "2024-09-12",
    {
        "_13": 2005,
        "_1983": 1984,
        "_1986": 24,
        "_1878": 2000,
        "_1987": 2006
    },
    "check_out",
    {
        "_1942": 1989,
        "_2002": 2007
    },
    "2024-09-14",
    {
        "_13": 2009,
        "_1983": 1984,
        "_1986": 24,
        "_1878": 2010,
        "_1987": 2011
    },
    "property_types",
    "(OPTIONAL parameter) An array that accepts one or more of the property enums defined below ONLY. Hotels are interpreted as HOTEL, vacation rentals as VR, resorts as RESORT.",
    {
        "_1942": 2012,
        "_2013": 2014,
        "_2002": 2020
    },
    "array",
    "items",
    {
        "_1942": 1989,
        "_2015": 2016
    },
    "enum",
    [
        2017,
        2018,
        2019
    ],
    "HOTEL",
    "RESORT",
    "VR",
    [
        2017,
        2019
    ],
    {
        "_13": 2022,
        "_1983": 1984,
        "_1986": 24,
        "_1878": 2023,
        "_1987": 2024
    },
    "number_of_travelers",
    "(OPTIONAL parameter) Applicable to vacation rentals only. An integer used to specify the total number of travelers for accommodations.",
    {
        "_1942": 2025,
        "_2002": 1197
    },
    "integer",
    {
        "_13": 2027,
        "_1983": 1984,
        "_1986": 24,
        "_1878": 2028,
        "_1987": 2029
    },
    "min_bedrooms",
    "(OPTIONAL parameter) Applicable to vacation rentals only. An integer used to specify minimum number of bedrooms.",
    {
        "_1942": 2025,
        "_2002": 545
    },
    {
        "_13": 2031,
        "_1983": 1984,
        "_1986": 24,
        "_1878": 2032,
        "_1987": 2033
    },
    "amenities",
    "(OPTIONAL parameter) An array that accepts one or more of the following amenity enums based on property_types.\n",
    {
        "_1942": 2012,
        "_2013": 2034
    },
    {
        "_1942": 1989,
        "_2015": 2035,
        "_2002": 2072
    },
    [
        2036,
        2037,
        2038,
        2039,
        2040,
        2041,
        2042,
        2043,
        2044,
        2045,
        2046,
        2047,
        2048,
        2049,
        2050,
        2051,
        2052,
        2053,
        2054,
        2055,
        2056,
        2057,
        2058,
        2059,
        2060,
        2061,
        2062,
        2063,
        2064,
        2065,
        2066,
        2067,
        2068,
        2069,
        2070,
        2071
    ],
    "GYM",
    "RESTAURANT",
    "BREAKFAST_INCLUDED",
    "HOT_TUB",
    "AIRPORT_SHUTTLE_INCLUDED",
    "INTERNET_OR_WIFI",
    "PET_FRIENDLY",
    "FAMILY_FRIENDLY",
    "KITCHEN",
    "ELECTRIC_CAR_CHARGING_STATION",
    "BAR",
    "CASINO",
    "AIR_CONDITIONING",
    "SPA",
    "POOL",
    "WATER_PARK",
    "PARKING",
    "OUTDOOR_SPACE",
    "OCEAN_VIEW",
    "SKI_IN_OR_SKI_OUT",
    "LOCAL_EXPERT",
    "ALL_INCLUSIVE",
    "PATIO_OR_DECK",
    "MICROWAVE",
    "TV",
    "FIREPLACE",
    "GARDEN_OR_BACKYARD",
    "PRIVATE_POOL",
    "GRILL",
    "DISHWASHER",
    "WASHER_AND_DRYER",
    "STOVE",
    "OVEN",
    "IRON_AND_BOARD",
    "KIDS_HIGH_CHAIR",
    "BALCONY",
    [
        2050,
        2049
    ],
    {
        "_13": 2074,
        "_1983": 1984,
        "_1986": 24,
        "_1878": 2075,
        "_1987": 2076
    },
    "guest_rating",
    "(OPTIONAL parameter) A string value limited to only one of the guest-rating enums. If the rating is an integer >=4.5, interpret it as WONDERFUL, if it is >=4, as VERY_GOOD, if it is >=3, as GOOD.\n",
    {
        "_1942": 1989,
        "_2015": 2077
    },
    [
        2078,
        2079,
        2080
    ],
    "WONDERFUL",
    "VERY_GOOD",
    "GOOD",
    {
        "_13": 2082,
        "_1983": 1984,
        "_1986": 24,
        "_1878": 2083,
        "_1987": 2084
    },
    "star_ratings",
    "(OPTIONAL parameter) Array limited to one or more of the star-rating enums. If request is for a luxury hotel, use [4,5]; for moderate use [3,3]; for a specific rating x use [x,x] instead of just x\n",
    {
        "_1942": 2012,
        "_2013": 2085
    },
    {
        "_1942": 2025,
        "_2015": 2086,
        "_2002": 2088
    },
    [
        406,
        2087,
        1197,
        545,
        1706,
        766
    ],
    1,
    [
        406,
        2087
    ],
    {
        "_13": 2090,
        "_1983": 1984,
        "_1986": 24,
        "_1878": 2091,
        "_1987": 2092
    },
    "sort_type",
    "(OPTIONAL parameter) A string value that allows user to get accommodations with the specified sort order.",
    {
        "_1942": 1989,
        "_2015": 2093,
        "_2002": 2094
    },
    [
        2094,
        2095,
        2096
    ],
    "CHEAPEST",
    "DISTANCE",
    "MOST_EXPENSIVE",
    {
        "_13": 2098,
        "_1983": 1984,
        "_1986": 24,
        "_1878": 2099,
        "_1987": 2100
    },
    "distance",
    "(OPTIONAL parameter) Distance around the given destination (in miles) to look up for options. Default unit is in miles.",
    {
        "_1942": 2025,
        "_2002": 434
    },
    "responses",
    {
        "_2103": 2104
    },
    "200",
    {
        "_1878": 2105,
        "_2106": 2107
    },
    "OK",
    "content",
    {
        "_2108": 2109
    },
    "application/json",
    {
        "_1987": 2110
    },
    {
        "_2111": 2112
    },
    "$ref",
    "#/components/schemas/LodgingResponse",
    "/recommendations/flights",
    {
        "_1973": 2115
    },
    {
        "_1975": 2116,
        "_1977": 2117,
        "_1979": 2118,
        "_2101": 2153
    },
    "flightProducts",
    "Gets recommended flights to destination",
    [
        2119,
        2122,
        2125,
        2130,
        2134,
        2138,
        2143,
        2147
    ],
    {
        "_13": 1982,
        "_1983": 1984,
        "_1878": 2120,
        "_1986": 61,
        "_1987": 2121
    },
    "(MANDATORY parameter) The merged user inputs according to the conversation context, translate to ENGLISH, DO NOT ONLY return the last query.",
    {
        "_1942": 1989
    },
    {
        "_13": 1991,
        "_1983": 1984,
        "_1878": 2123,
        "_1986": 61,
        "_1987": 2124
    },
    "(MANDATORY parameter) Keywords associated with TRAVEL in the user_input_in_english parameter, converted to all LOWERCASE, separated by | character.",
    {
        "_1942": 1989
    },
    {
        "_13": 2126,
        "_1983": 1984,
        "_1878": 2127,
        "_1986": 61,
        "_1987": 2128,
        "_2002": 2129
    },
    "origin",
    "(MANDATORY parameter) Origin location name or airport code.",
    {
        "_1942": 1989
    },
    "LAS",
    {
        "_13": 1995,
        "_1983": 1984,
        "_1878": 2131,
        "_1986": 61,
        "_1987": 2132,
        "_2002": 2133
    },
    "(MANDATORY parameter) Destination location name or airport code.",
    {
        "_1942": 1989
    },
    "LAX",
    {
        "_13": 2135,
        "_1983": 1984,
        "_1878": 2000,
        "_1986": 24,
        "_1987": 2136,
        "_2002": 2137
    },
    "departure_date",
    {
        "_1942": 1989
    },
    "2025-07-01",
    {
        "_13": 2139,
        "_1983": 1984,
        "_1878": 2140,
        "_1986": 24,
        "_1987": 2141,
        "_2002": 2142
    },
    "airline_code",
    "(OPTIONAL parameter) 2 letter Airline code.",
    {
        "_1942": 1989
    },
    "AA",
    {
        "_13": 2144,
        "_1983": 1984,
        "_1878": 2145,
        "_1986": 24,
        "_1987": 2146,
        "_2002": 406
    },
    "number_of_stops",
    "(OPTIONAL parameter) Number of stops preferred. 0 means non-stop, 1 means either 0 or 1 stop etc.",
    {
        "_1942": 2025
    },
    {
        "_13": 2090,
        "_1983": 1984,
        "_1986": 24,
        "_1878": 2148,
        "_1987": 2149
    },
    "Optional string value that allows user to get Flights with the specified sort order. \\ \\ Use PRICE to sort by cheapest, DURATION to sort by shortest duration flight. Default is PRICE.",
    {
        "_1942": 1989,
        "_2015": 2150,
        "_2002": 2151
    },
    [
        2151,
        2152
    ],
    "PRICE",
    "DURATION",
    {
        "_2103": 2154
    },
    {
        "_1878": 2105,
        "_2106": 2155
    },
    {
        "_2108": 2156
    },
    {
        "_1987": 2157
    },
    {
        "_2111": 2158
    },
    "#/components/schemas/FlightResponse",
    "/recommendations/activities",
    {
        "_1973": 2161
    },
    {
        "_1975": 2162,
        "_1977": 2163,
        "_1979": 2164,
        "_2101": 2216
    },
    "activityProducts",
    "Get a list of activity travel products",
    [
        2165,
        2167,
        2169,
        2173,
        2177,
        2181,
        2202,
        2211
    ],
    {
        "_13": 1982,
        "_1983": 1984,
        "_1878": 2120,
        "_1986": 61,
        "_1987": 2166
    },
    {
        "_1942": 1989
    },
    {
        "_13": 1991,
        "_1983": 1984,
        "_1878": 2123,
        "_1986": 61,
        "_1987": 2168
    },
    {
        "_1942": 1989
    },
    {
        "_13": 1995,
        "_1983": 1984,
        "_1878": 2170,
        "_1986": 61,
        "_1987": 2171,
        "_2002": 2172
    },
    "(MANDATORY parameter) City name, street address, three-letter IATA Airport Code or a landmark name.",
    {
        "_1942": 1989
    },
    "shenzhen",
    {
        "_13": 2174,
        "_1983": 1984,
        "_1878": 2000,
        "_1987": 2175,
        "_2002": 2176
    },
    "start_date",
    {
        "_1942": 1989
    },
    "2024-10-01",
    {
        "_13": 2178,
        "_1983": 1984,
        "_1878": 2000,
        "_1987": 2179,
        "_2002": 2180
    },
    "end_date",
    {
        "_1942": 1989
    },
    "2024-10-10",
    {
        "_13": 1890,
        "_1983": 1984,
        "_1986": 24,
        "_1878": 2182,
        "_1987": 2183
    },
    "(OPTIONAL parameter) An array that accepts one or more of the following category enums. For example if the activity category is \"family-friendly\", interpret it as FAMILY_FRIENDLY.",
    {
        "_1942": 2012,
        "_2013": 2184
    },
    {
        "_1942": 1989,
        "_2015": 2185,
        "_2002": 2200
    },
    [
        2043,
        2186,
        2187,
        2188,
        2189,
        2190,
        2191,
        2192,
        2193,
        2194,
        2195,
        2196,
        2197,
        2198,
        2199
    ],
    "LOCAL_EXPERTS_PICKS",
    "SELECTIVE_HOTEL_PICKUP",
    "FREE_CANCELLATION",
    "NIGHTLIFE",
    "DEALS",
    "WALKING_BIKE_TOURS",
    "FOOD_DRINK",
    "ADVENTURES",
    "ATTRACTIONS",
    "CRUISES_WATER_TOURS",
    "THEME_PARKS",
    "TOURS_SIGHTSEEING",
    "WATER_ACTIVITIES",
    "DAY_TRIPS_EXCURSIONS",
    [
        2201
    ],
    "FAMILY_FRIENDLY\uff0cFREE_CANCELLATION",
    {
        "_13": 2203,
        "_1983": 1984,
        "_1986": 24,
        "_1878": 2204,
        "_1987": 2205
    },
    "duration",
    "(OPTIONAL parameter) Enum value that allows getting activities within the specified duration. Match the user stated duration preference to the appropriate enum value.\n",
    {
        "_1942": 1989,
        "_2015": 2206,
        "_2002": 2207
    },
    [
        2207,
        2208,
        2209,
        2210
    ],
    "LESS_THAN_ONE_HOUR",
    "ONE_TO_FOUR_HOURS",
    "FOUR_HOURS_TO_ONE_DAY",
    "MORE_THAN_ONE_DAY",
    {
        "_13": 2212,
        "_1983": 1984,
        "_1986": 24,
        "_1878": 2213,
        "_1987": 2214
    },
    "price_max",
    "(OPTIONAL parameter) The maximum price of an activity.",
    {
        "_1942": 2215,
        "_2002": 403
    },
    "number",
    {
        "_2103": 2217
    },
    {
        "_1878": 2105,
        "_2106": 2218
    },
    {
        "_2108": 2219
    },
    {
        "_1987": 2220
    },
    {
        "_2111": 2221
    },
    "#/components/schemas/ActivityResponse",
    "/recommendations/cars",
    {
        "_1973": 2224
    },
    {
        "_1975": 2225,
        "_1977": 2226,
        "_1979": 2227,
        "_2101": 2279
    },
    "carProducts",
    "Get a list of car travel products",
    [
        2228,
        2230,
        2232,
        2237,
        2241,
        2245,
        2250,
        2254,
        2258
    ],
    {
        "_13": 1982,
        "_1983": 1984,
        "_1878": 2120,
        "_1986": 61,
        "_1987": 2229
    },
    {
        "_1942": 1989
    },
    {
        "_13": 1991,
        "_1983": 1984,
        "_1878": 2123,
        "_1986": 61,
        "_1987": 2231
    },
    {
        "_1942": 1989
    },
    {
        "_13": 2233,
        "_1983": 1984,
        "_1878": 2234,
        "_1986": 61,
        "_1987": 2235,
        "_2002": 2236
    },
    "pickup_location",
    "(MANDATORY parameter) Car rental pick-up location. It can be a city name, address, airport code or a landmark name.",
    {
        "_1942": 1989
    },
    "Seattle",
    {
        "_13": 2238,
        "_1983": 1984,
        "_1878": 2239,
        "_1986": 24,
        "_1987": 2240,
        "_2002": 2236
    },
    "dropoff_location",
    "(OPTIONAL parameter) Car rental drop-off location. It can be a city name, address, airport code or a landmark name. By default, it is same as that of pick-up location.",
    {
        "_1942": 1989
    },
    {
        "_13": 2242,
        "_1983": 1984,
        "_1878": 2000,
        "_1986": 24,
        "_1987": 2243,
        "_2002": 2244
    },
    "pickup_date",
    {
        "_1942": 1989
    },
    "2025-06-05",
    {
        "_13": 2246,
        "_1983": 1984,
        "_1878": 2247,
        "_1986": 24,
        "_1987": 2248,
        "_2002": 2249
    },
    "pickup_time",
    "(OPTIONAL parameter) Accept any time format and convert to HH:MM (24-hour format).",
    {
        "_1942": 1989
    },
    600,
    {
        "_13": 2251,
        "_1983": 1984,
        "_1878": 2000,
        "_1986": 24,
        "_1987": 2252,
        "_2002": 2253
    },
    "dropoff_date",
    {
        "_1942": 1989
    },
    "2025-06-08",
    {
        "_13": 2255,
        "_1983": 1984,
        "_1878": 2247,
        "_1986": 24,
        "_1987": 2256,
        "_2002": 2257
    },
    "dropoff_time",
    {
        "_1942": 1989
    },
    840,
    {
        "_13": 2259,
        "_1983": 1984,
        "_1986": 24,
        "_1878": 2260,
        "_1987": 2261
    },
    "car_classes",
    "(OPTIONAL parameter) This value is used to filter API queries to only return a certain type(s) of car(s).",
    {
        "_1942": 2012,
        "_2013": 2262
    },
    {
        "_1942": 1989,
        "_2015": 2263,
        "_2002": 2278
    },
    [
        2264,
        2265,
        2266,
        2267,
        2268,
        2269,
        2270,
        2271,
        2272,
        2273,
        2274,
        2275,
        2276,
        2277
    ],
    "ECONOMY",
    "COMPACT",
    "MIDSIZE",
    "STANDARD",
    "FULLSIZE",
    "PREMIUM",
    "LUXURY",
    "VAN",
    "SUV",
    "MINI",
    "CONVERTIBLE",
    "MINIVAN",
    "PICKUP",
    "SPORTSCAR",
    [
        2264,
        2272
    ],
    {
        "_2103": 2280
    },
    {
        "_1878": 2105,
        "_2106": 2281
    },
    {
        "_2108": 2282
    },
    {
        "_1987": 2283
    },
    {
        "_2111": 2284
    },
    "#/components/schemas/CarProductResponse",
    "components",
    {
        "_2287": 2288
    },
    "schemas",
    {
        "_2289": 2290,
        "_2301": 2302,
        "_2349": 2350,
        "_2357": 2358,
        "_2392": 2393,
        "_2411": 2412,
        "_2419": 2420,
        "_2444": 2445,
        "_2452": 2453
    },
    "LodgingResponse",
    {
        "_1942": 2291,
        "_2292": 2293
    },
    "object",
    "properties",
    {
        "_3": 2294,
        "_2298": 2299
    },
    {
        "_1942": 2012,
        "_1878": 2295,
        "_2013": 2296
    },
    "List of lodging recommendations.",
    {
        "_2111": 2297
    },
    "#/components/schemas/LodgingResponseData",
    "EXTRA_INFORMATION_TO_ASSISTANT",
    {
        "_1942": 1989,
        "_1878": 2300
    },
    "Specific instructions on how assistant is supposed to handle the data included in the API response",
    "LodgingResponseData",
    {
        "_1942": 2291,
        "_2292": 2303
    },
    {
        "_2304": 2305,
        "_2307": 2308,
        "_1878": 2310,
        "_2312": 2313,
        "_2315": 2316,
        "_2318": 2319,
        "_2321": 2322,
        "_2074": 2324,
        "_2326": 2327,
        "_2329": 2330,
        "_2332": 2333,
        "_2335": 2336,
        "_2338": 2339,
        "_2341": 2342,
        "_1967": 2344,
        "_2346": 2347
    },
    "hotel_id",
    {
        "_1942": 1989,
        "_1878": 2306
    },
    "Hotel's unique identifier",
    "hotel_name",
    {
        "_1942": 1989,
        "_1878": 2309
    },
    "Hotel name",
    {
        "_1942": 1989,
        "_1878": 2311
    },
    "Short description about the hotel.",
    "location_description",
    {
        "_1942": 1989,
        "_1878": 2314
    },
    "Short location description of the hotel.",
    "max_occupancy",
    {
        "_1942": 2025,
        "_1878": 2317
    },
    "Maximum occupancy allowed for the accommodation.",
    "number_of_bedrooms",
    {
        "_1942": 2025,
        "_1878": 2320
    },
    "Number of bedrooms in the property.",
    "star_rating",
    {
        "_1942": 1989,
        "_1878": 2323
    },
    "Star rating of the hotel.",
    {
        "_1942": 1989,
        "_1878": 2325
    },
    "The guest rating of the hotel in expedia web site, max rating value is 5.",
    "guest_review_count",
    {
        "_1942": 2025,
        "_1878": 2328
    },
    "The guest review count of this product in expedia web site.",
    "avg_nightly_price",
    {
        "_1942": 1989,
        "_1878": 2331
    },
    "Price per night for the hotel",
    "checkin_date",
    {
        "_1942": 1989,
        "_1878": 2334
    },
    "Check-in date for the hotel stay in format YYYY-MM-DD",
    "checkout_date",
    {
        "_1942": 1989,
        "_1878": 2337
    },
    "Check-out date for the hotel stay in format YYYY-MM-DD",
    "currency",
    {
        "_1942": 1989,
        "_1878": 2340
    },
    "Currency in which avg nightly price is specified",
    "promotion",
    {
        "_1942": 1989,
        "_1878": 2343
    },
    "Promotion of the hotel,like member saving",
    {
        "_1942": 1989,
        "_1878": 2345
    },
    "Link to the hotel on Expedia.\\ \\ Include this in the user response prompt whenever available. \\",
    "preview_photo",
    {
        "_1942": 1989,
        "_1878": 2348
    },
    "A link to the preview photo of the hotel. \\ \\  Include this in the user response prompt whenever available.\\",
    "FlightResponse",
    {
        "_1942": 2291,
        "_2292": 2351
    },
    {
        "_3": 2352,
        "_2298": 2356
    },
    {
        "_1942": 2012,
        "_1878": 2353,
        "_2013": 2354
    },
    "List of flight recommendations",
    {
        "_2111": 2355
    },
    "#/components/schemas/FlightResponseData",
    {
        "_1942": 1989,
        "_1878": 2300
    },
    "FlightResponseData",
    {
        "_1942": 2291,
        "_2292": 2359
    },
    {
        "_2360": 2361,
        "_2144": 2365,
        "_2367": 2368,
        "_2371": 2372,
        "_2338": 2375,
        "_2135": 2377,
        "_2380": 2381,
        "_2384": 2385,
        "_1967": 2388,
        "_2346": 2390
    },
    "legs",
    {
        "_1942": 2012,
        "_1878": 2362,
        "_2013": 2363
    },
    "The list of flight legs in the segment.",
    {
        "_2111": 2364
    },
    "#/components/schemas/FlightLeg",
    {
        "_1942": 2025,
        "_1878": 2366,
        "_2002": 2087
    },
    "Total number of stops in this segment",
    "flight_duration",
    {
        "_1942": 1989,
        "_1878": 2369,
        "_2002": 2370
    },
    "Total duration of the flight segment",
    "9h 40m",
    "price_per_ticket",
    {
        "_1942": 1989,
        "_1878": 2373,
        "_2002": 2374
    },
    "Price per ticket for the selected Flight",
    "89.78",
    {
        "_1942": 1989,
        "_1878": 2376
    },
    "Currency in which ticket price is specified",
    {
        "_1942": 1989,
        "_1878": 2378,
        "_2002": 2379
    },
    "Date of departure of the flight in format 'YYY-MM-DD'",
    "2021-07-05",
    "departure_time",
    {
        "_1942": 1989,
        "_1878": 2382,
        "_2002": 2383
    },
    "Time of departure of the flight in format 'HH:mm aa'",
    "07:15 AM",
    "arrival_time",
    {
        "_1942": 1989,
        "_1878": 2386,
        "_2002": 2387
    },
    "Time of arrival of the flight in format 'HH:mm aa'",
    "12:30 PM",
    {
        "_1942": 1989,
        "_1878": 2389
    },
    "Link to book the flight on Expedia. \\ \\  Include this in the user response prompt whenever available.\\",
    {
        "_1942": 1989,
        "_1878": 2391
    },
    "A link to the preview photo of the flight. \\ \\  Include this in the user response prompt whenever available.\\",
    "FlightLeg",
    {
        "_1942": 2291,
        "_2292": 2394
    },
    {
        "_2395": 2396,
        "_2399": 2400,
        "_2403": 2404,
        "_2407": 2408
    },
    "airline_name",
    {
        "_1942": 1989,
        "_1878": 2397,
        "_2002": 2398
    },
    "Airline name for the flight.",
    "United Airlines",
    "flight_number",
    {
        "_1942": 1989,
        "_1878": 2401,
        "_2002": 2402
    },
    "Flight number corresponding to the airline.",
    "1523",
    "departure_airport_code",
    {
        "_1942": 1989,
        "_1878": 2405,
        "_2002": 2406
    },
    "Deaprture airport code",
    "SEA",
    "arrival_airport_code",
    {
        "_1942": 1989,
        "_1878": 2409,
        "_2002": 2410
    },
    "Arrival airport code",
    "SFO",
    "ActivityResponse",
    {
        "_1942": 2291,
        "_2292": 2413
    },
    {
        "_3": 2414,
        "_2298": 2418
    },
    {
        "_1942": 2012,
        "_1878": 2415,
        "_2013": 2416
    },
    "List of the activities.",
    {
        "_2111": 2417
    },
    "#/components/schemas/ActivityResponseData",
    {
        "_1942": 1989,
        "_1878": 2300
    },
    "ActivityResponseData",
    {
        "_1942": 2291,
        "_2292": 2421
    },
    {
        "_2422": 2423,
        "_2425": 2426,
        "_2371": 2428,
        "_2338": 2430,
        "_2203": 2431,
        "_1890": 2433,
        "_2436": 2437,
        "_1967": 2440,
        "_2346": 2442
    },
    "activity_name",
    {
        "_1942": 1989,
        "_1878": 2424
    },
    "Name for the activity.",
    "activity_description",
    {
        "_1942": 1989,
        "_1878": 2427
    },
    "Description of the activity.",
    {
        "_1942": 1989,
        "_1878": 2429
    },
    "Price per ticket for the activity",
    {
        "_1942": 1989,
        "_1878": 2376
    },
    {
        "_1942": 1989,
        "_1878": 2432,
        "_2002": 2370
    },
    "Total duration of the activity",
    {
        "_1942": 2012,
        "_1878": 2434,
        "_2013": 2435
    },
    "The list of activity categories.",
    {
        "_1942": 1989
    },
    "saving_percentage",
    {
        "_1942": 1989,
        "_1878": 2438,
        "_2002": 2439
    },
    "Saving percentage of the activity",
    "45",
    {
        "_1942": 1989,
        "_1878": 2441
    },
    "A link to the activity product on Expedia. \\ \\  Include this in the user response prompt whenever available.\\",
    {
        "_1942": 1989,
        "_1878": 2443
    },
    "A link to the preview photo of the activity. \\ \\  Include this in the user response prompt whenever available.\\",
    "CarProductResponse",
    {
        "_1942": 2291,
        "_2292": 2446
    },
    {
        "_3": 2447,
        "_2298": 2451
    },
    {
        "_1942": 2012,
        "_1878": 2448,
        "_2013": 2449
    },
    "The list of car products.",
    {
        "_2111": 2450
    },
    "#/components/schemas/CarProductResponseData",
    {
        "_1942": 1989,
        "_1878": 2300
    },
    "CarProductResponseData",
    {
        "_1942": 2291,
        "_2292": 2454
    },
    {
        "_2455": 2456,
        "_2458": 2459,
        "_2461": 2462,
        "_2464": 2465,
        "_2242": 2467,
        "_2469": 2470,
        "_2251": 2472,
        "_2474": 2475,
        "_2338": 2477,
        "_1967": 2479,
        "_2346": 2481
    },
    "car_make",
    {
        "_1942": 1989,
        "_1878": 2457
    },
    "The Car manufacturer and model.",
    "car_class",
    {
        "_1942": 1989,
        "_1878": 2460
    },
    "The car category and type.",
    "supplier_name",
    {
        "_1942": 1989,
        "_1878": 2463
    },
    "The supplier name of this car product.",
    "pickup_address",
    {
        "_1942": 1989,
        "_1878": 2466
    },
    "The pickup address information of the location.",
    {
        "_1942": 1989,
        "_1878": 2468
    },
    "The date time that pickup the car.",
    "dropoff_address",
    {
        "_1942": 1989,
        "_1878": 2471
    },
    "The dropoff address information of the location",
    {
        "_1942": 1989,
        "_1878": 2473
    },
    "The date time that dropoff the car.",
    "total_price",
    {
        "_1942": 1989,
        "_1878": 2476
    },
    "Total price for the car with currency",
    {
        "_1942": 1989,
        "_1878": 2478
    },
    "Currency in which total price is specified",
    {
        "_1942": 1989,
        "_1878": 2480
    },
    "A link to the car product on Expedia. \\ \\  Include this in the user response prompt whenever available.\\",
    {
        "_1942": 1989,
        "_1878": 2482
    },
    "A link to the preview photo of the car. \\ \\  Include this in the user response prompt whenever available.\\",
    "auth",
    {
        "_1942": 2485,
        "_1874": 455,
        "_2486": 2487,
        "_2488": 2489,
        "_2490": 455
    },
    "service_http",
    "authorization_type",
    "basic",
    "verification_tokens",
    {},
    "custom_auth_header",
    "privacy_policy_url",
    "https://legal.expediagroup.com/privacy/privacy-and-cookies-statements/privacy-statement-all",
    "files",
    [],
    "product_features",
    {
        "_2497": 2498
    },
    "attachments",
    {
        "_1942": 2499,
        "_2500": 2501,
        "_2576": 2577,
        "_2582": 61
    },
    "retrieval",
    "accepted_mime_types",
    [
        2502,
        2503,
        2504,
        2505,
        2506,
        2507,
        2508,
        2509,
        2510,
        2511,
        2512,
        2513,
        2514,
        2515,
        2516,
        2517,
        2518,
        2519,
        2520,
        2521,
        2522,
        2523,
        2524,
        2525,
        2526,
        2527,
        2528,
        2529,
        2530,
        2531,
        2532,
        2533,
        2534,
        2535,
        2536,
        2537,
        2538,
        2539,
        2540,
        2541,
        2542,
        2108,
        2543,
        2544,
        2545,
        2546,
        2547,
        2548,
        2549,
        2550,
        2551,
        2552,
        2553,
        2554,
        2555,
        2556,
        2557,
        2558,
        2559,
        2560,
        2561,
        2562,
        2563,
        2564,
        2565,
        2566,
        2567,
        2568,
        2569,
        2570,
        2571,
        2572,
        2573,
        2574,
        2575
    ],
    "text/x-csharp",
    "application/vnd.apple.pages",
    "text/x-typescript",
    "application/javascript",
    "text/x-liquid",
    "text/x-php",
    "text/x-tmpl",
    "text/x-astro",
    "application/x-sql",
    "text/javascript",
    "text/x-dart",
    "text/x-diff",
    "text/x-objectivec",
    "text/x-r",
    "text/xml",
    "text/x-lisp",
    "text/markdown",
    "text/x-erlang",
    "text/x-handlebars",
    "text/x-asm",
    "application/vnd.openxmlformats-officedocument.presentationml.presentation",
    "application/rtf",
    "message/rfc822",
    "text/x-c",
    "text/x-shellscript",
    "application/vnd.apple.keynote",
    "application/x-scala",
    "text/x-go",
    "text/x-julia",
    "application/x-yaml",
    "text/tsx",
    "text/x-jinja2",
    "application/pdf",
    "text/x-pug",
    "text/x-ejs",
    "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
    "text/x-kotlin",
    "text/x-script.python",
    "text/x-clojure",
    "text/vbscript",
    "text/x-rst",
    "text/x-scala",
    "text/x-rust",
    "text/jsx",
    "text/x-lua",
    "text/x-c++",
    "text/rtf",
    "text/x-groovy",
    "text/x-erb",
    "text/calendar",
    "application/x-powershell",
    "application/vnd.ms-powerpoint",
    "text/plain",
    "text/html",
    "text/x-ruby",
    "text/x-sh",
    "text/x-java",
    "text/x-mustache",
    "text/x-haskell",
    "text/x-jade",
    "text/x-vcard",
    "text/x-tex",
    "application/vnd.oasis.opendocument.text",
    "text/x-swift",
    "text/x-twig",
    "text/x-perl",
    "text/x-elixir",
    "text/css",
    "application/x-rust",
    "application/toml",
    "application/msword",
    "text/x-python",
    "text/x-makefile",
    "text/x-objectivec++",
    "image_mime_types",
    [
        2578,
        2579,
        2580,
        2581
    ],
    "image/png",
    "image/gif",
    "image/webp",
    "image/jpeg",
    "can_accept_all_mime_types"
]
````

## File: templates/initialize.json
````json
{
    "feature_gates": {
        "14938527": {
            "name": "14938527",
            "value": true,
            "rule_id": "3QgLJ91lKIc7VAOjo5SDz7",
            "id_type": "stableID",
            "secondary_exposures": []
        },
        "331938894": {
            "name": "331938894",
            "value": false,
            "rule_id": "default",
            "id_type": "userID",
            "secondary_exposures": []
        },
        "374328013": {
            "name": "374328013",
            "value": false,
            "rule_id": "5xNC0SZHSEfRGwLRCyEofs:0.00:1",
            "id_type": "userID",
            "secondary_exposures": [
                {
                    "gate": "32oERDChgC0TNbTHexysy1lPoFrhE0PqVoWftp5krGyJ2mEQTnyAJtFpcO8fECIaC",
                    "gateValue": "false",
                    "ruleID": "4nM2ehmgoDQIv69B0zohb6"
                }
            ]
        },
        "512302793": {
            "name": "512302793",
            "value": false,
            "rule_id": "default",
            "id_type": "stableID",
            "secondary_exposures": []
        },
        "543687013": {
            "name": "543687013",
            "value": true,
            "rule_id": "31a8ZixwfunGZwX1gacTqM",
            "id_type": "userID",
            "secondary_exposures": []
        },
        "727502549": {
            "name": "727502549",
            "value": true,
            "rule_id": "6EYbmM9CyqCRO6U6k3dROA",
            "id_type": "userID",
            "secondary_exposures": []
        },
        "773249106": {
            "name": "773249106",
            "value": false,
            "rule_id": "1kGO9xYmxaBS2V2H3LcQuG",
            "id_type": "userID",
            "secondary_exposures": []
        },
        "850280859": {
            "name": "850280859",
            "value": false,
            "rule_id": "default",
            "id_type": "stableID",
            "secondary_exposures": []
        },
        "989108178": {
            "name": "989108178",
            "value": false,
            "rule_id": "4sTodKrNyByM4guZ68MORR",
            "id_type": "userID",
            "secondary_exposures": [
                {
                    "gate": "MBvgUBv44KFgWHcvhFmwoffVvNH48mNqI1URabr7r2W",
                    "gateValue": "false",
                    "ruleID": "2EjTipm6C4kk4fuvcHMzZe"
                },
                {
                    "gate": "rnnRuXkrZqfaPUXWOKBgqwHsIrmFzDxX5ZjpQFxzg4h",
                    "gateValue": "true",
                    "ruleID": "4C2vO0R7mvnCZvl1HDBExp:30.00:5"
                }
            ]
        },
        "1016398872": {
            "name": "1016398872",
            "value": false,
            "rule_id": "default",
            "id_type": "userID",
            "secondary_exposures": [
                {
                    "gate": "LUSqsmfgqSH9QxlX0hGURoa4JTMrV9uUF72cokAWPDx",
                    "gateValue": "true",
                    "ruleID": "2PP6pudW64Hn7katvazhAx:100.00:5"
                },
                {
                    "gate": "MeBQxlGbe6gcnQEw7XJVAPzHKdo3ON3Vc4lFRdy4AnS",
                    "gateValue": "true",
                    "ruleID": "4bd3o553p0ZCRkFmipROd8"
                },
                {
                    "gate": "Ul3Lff3cZuLWsYWpvf8NPhQBkiNf7BZhSa3hcvZ6r8Y",
                    "gateValue": "true",
                    "ruleID": "60QaTyBFJYTakinhLvhAM9"
                }
            ]
        },
        "1041874561": {
            "name": "1041874561",
            "value": true,
            "rule_id": "3QhGbHZgk78rDbdIQis5P7",
            "id_type": "userID",
            "secondary_exposures": []
        },
        "1105502266": {
            "name": "1105502266",
            "value": true,
            "rule_id": "6aawHEq6J83iPX0WB5PkaS:100.00:1",
            "id_type": "userID",
            "secondary_exposures": []
        },
        "1187676131": {
            "name": "1187676131",
            "value": false,
            "rule_id": "4CqZUX3nf3nrcFqQjlXqsN",
            "id_type": "stableID",
            "secondary_exposures": []
        },
        "1325692000": {
            "name": "1325692000",
            "value": false,
            "rule_id": "4oohwt222TyLfA2FPN89PX:0.00:1",
            "id_type": "userID",
            "secondary_exposures": [
                {
                    "gate": "RZMejDZ68BXpMiYhPC69S",
                    "gateValue": "true",
                    "ruleID": "48uk8WE58h6KWvjRmyvacW:100.00:1"
                }
            ]
        },
        "1489221567": {
            "name": "1489221567",
            "value": false,
            "rule_id": "z06JxGlMXQVrwG418FqnU",
            "id_type": "stableID",
            "secondary_exposures": []
        },
        "1626939786": {
            "name": "1626939786",
            "value": false,
            "rule_id": "default",
            "id_type": "userID",
            "secondary_exposures": []
        },
        "1719651090": {
            "name": "1719651090",
            "value": true,
            "rule_id": "60QaTyBFJYTakinhLvhAM9",
            "id_type": "userID",
            "secondary_exposures": [
                {
                    "gate": "LUSqsmfgqSH9QxlX0hGURoa4JTMrV9uUF72cokAWPDx",
                    "gateValue": "true",
                    "ruleID": "2PP6pudW64Hn7katvazhAx:100.00:5"
                },
                {
                    "gate": "MeBQxlGbe6gcnQEw7XJVAPzHKdo3ON3Vc4lFRdy4AnS",
                    "gateValue": "true",
                    "ruleID": "4bd3o553p0ZCRkFmipROd8"
                }
            ]
        },
        "1804926979": {
            "name": "1804926979",
            "value": false,
            "rule_id": "default",
            "id_type": "userID",
            "secondary_exposures": []
        },
        "1805865645": {
            "name": "1805865645",
            "value": false,
            "rule_id": "default",
            "id_type": "stableID",
            "secondary_exposures": []
        },
        "1847911009": {
            "name": "1847911009",
            "value": false,
            "rule_id": "5OIO2mI7iQiPRReG1jZ4c2:0.00:7",
            "id_type": "userID",
            "secondary_exposures": [
                {
                    "gate": "pjUMjKOQt6adsSRyY30wG1y9QJwP3nc915vIZbBAEyL",
                    "gateValue": "true",
                    "ruleID": "xhzqCEbupJKNSWjgWMMly:100.00:5"
                }
            ]
        },
        "1855896025": {
            "name": "1855896025",
            "value": false,
            "rule_id": "default",
            "id_type": "userID",
            "secondary_exposures": []
        },
        "1902899872": {
            "name": "1902899872",
            "value": true,
            "rule_id": "58UOuEcFwyqlorfhrWQLlE",
            "id_type": "userID",
            "secondary_exposures": [
                {
                    "gate": "pjUMjKOQt6adsSRyY30wG1y9QJwP3nc915vIZbBAEyL",
                    "gateValue": "true",
                    "ruleID": "xhzqCEbupJKNSWjgWMMly:100.00:5"
                }
            ]
        },
        "1981644628": {
            "name": "1981644628",
            "value": false,
            "rule_id": "default",
            "id_type": "userID",
            "secondary_exposures": []
        },
        "2048457345": {
            "name": "2048457345",
            "value": true,
            "rule_id": "3SY8HhaKZ365bspxYp7ZW8:100.00:6",
            "id_type": "userID",
            "secondary_exposures": [
                {
                    "gate": "7guDEekuMu5Samg8tuMbApZDF8o3z10UnPbUpYcFg2BbvF6Q55fpKsBiv51yTDH6n",
                    "gateValue": "false",
                    "ruleID": "5KvTiw548r4tjqy7oHEcAL"
                },
                {
                    "gate": "RENl24KWMksVzt97oM2bjklDkAt8YEpnWi4Ss2cTfy5",
                    "gateValue": "false",
                    "ruleID": "default"
                }
            ]
        },
        "2113934735": {
            "name": "2113934735",
            "value": true,
            "rule_id": "1THEw1F2Q9SX9YbmRLxO61",
            "id_type": "userID",
            "secondary_exposures": [
                {
                    "gate": "LUSqsmfgqSH9QxlX0hGURoa4JTMrV9uUF72cokAWPDx",
                    "gateValue": "true",
                    "ruleID": "2PP6pudW64Hn7katvazhAx:100.00:5"
                },
                {
                    "gate": "MeBQxlGbe6gcnQEw7XJVAPzHKdo3ON3Vc4lFRdy4AnS",
                    "gateValue": "true",
                    "ruleID": "4bd3o553p0ZCRkFmipROd8"
                },
                {
                    "gate": "Ul3Lff3cZuLWsYWpvf8NPhQBkiNf7BZhSa3hcvZ6r8Y",
                    "gateValue": "true",
                    "ruleID": "60QaTyBFJYTakinhLvhAM9"
                }
            ]
        },
        "2124994664": {
            "name": "2124994664",
            "value": false,
            "rule_id": "VGVVLgcrSRl5RUAzvAaji",
            "id_type": "stableID",
            "secondary_exposures": []
        },
        "2304807207": {
            "name": "2304807207",
            "value": true,
            "rule_id": "xhzqCEbupJKNSWjgWMMly:100.00:5",
            "id_type": "userID",
            "secondary_exposures": []
        },
        "2687575887": {
            "name": "2687575887",
            "value": true,
            "rule_id": "10cvQmwrcZvpWBFlZgn8pZ",
            "id_type": "userID",
            "secondary_exposures": []
        },
        "2967706710": {
            "name": "2967706710",
            "value": false,
            "rule_id": "default",
            "id_type": "stableID",
            "secondary_exposures": []
        },
        "3199899666": {
            "name": "3199899666",
            "value": false,
            "rule_id": "default",
            "id_type": "userID",
            "secondary_exposures": [
                {
                    "gate": "7guDEekuMu5Samg8tuMbApZDF8o3z10UnPbUpYcFg2BbvF6Q55fpKsBiv51yTDH6n",
                    "gateValue": "false",
                    "ruleID": "5KvTiw548r4tjqy7oHEcAL"
                }
            ]
        },
        "3257646228": {
            "name": "3257646228",
            "value": false,
            "rule_id": "default",
            "id_type": "stableID",
            "secondary_exposures": []
        },
        "3319375497": {
            "name": "3319375497",
            "value": false,
            "rule_id": "1iTlrKHQQtqFJt5WZh6U8u:0.00:1",
            "id_type": "userID",
            "secondary_exposures": []
        },
        "3325813340": {
            "name": "3325813340",
            "value": true,
            "rule_id": "37GsRLj07CqERPyHBn4o5L",
            "id_type": "userID",
            "secondary_exposures": [
                {
                    "gate": "pjUMjKOQt6adsSRyY30wG1y9QJwP3nc915vIZbBAEyL",
                    "gateValue": "true",
                    "ruleID": "xhzqCEbupJKNSWjgWMMly:100.00:5"
                }
            ]
        },
        "3504894861": {
            "name": "3504894861",
            "value": false,
            "rule_id": "default",
            "id_type": "userID",
            "secondary_exposures": []
        },
        "3544641259": {
            "name": "3544641259",
            "value": false,
            "rule_id": "default",
            "id_type": "stableID",
            "secondary_exposures": [
                {
                    "gate": "AdIoJdQdHTdnh7pIhnL4lqJaseo2hLAdB5SV2xxuyDPDl77GkjGfoluyXacog4YR",
                    "gateValue": "false",
                    "ruleID": "default"
                }
            ]
        },
        "3993182790": {
            "name": "3993182790",
            "value": false,
            "rule_id": "default",
            "id_type": "userID",
            "secondary_exposures": []
        }
    },
    "dynamic_configs": {
        "392393436": {
            "name": "392393436",
            "value": {
                "greeting_web": false
            },
            "rule_id": "prestart",
            "group": "prestart",
            "is_device_based": false,
            "id_type": "userID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "greeting_web"
            ],
            "secondary_exposures": []
        },
        "567396244": {
            "name": "567396244",
            "value": {
                "is_team_enabled": true,
                "is_yearly_plus_subscription_enabled": false,
                "is_split_between_personal_and_business_enabled": true,
                "is_modal_fullscreen": true,
                "is_v2_toggle_labels_enabled": false,
                "is_bw": false,
                "is_produce_colors": false,
                "is_produce_color_scheme": false,
                "is_mobile_web_toggle_enabled": false,
                "is_enterprise_enabled": false,
                "is_produce_text": false,
                "is_optimized_checkout": false,
                "is_save_stripe_payment_info_enabled": false,
                "is_auto_save_stripe_payment_info_enabled": false,
                "does_manage_my_subscription_link_take_user_to_subscription_settings": true,
                "should_open_cancellation_survey_after_canceling": true,
                "should_cancel_button_take_user_to_stripe": false,
                "should_show_manage_my_subscription_link": true,
                "is_stripe_manage_subscription_link_enabled": true,
                "cancellation_modal_cancel_button_color": "danger",
                "cancellation_modal_go_back_button_color": "secondary",
                "should_show_cp": false,
                "cp_eligibility_months": 3,
                "should_offer_paypal_when_eligible": false
            },
            "rule_id": "layerAssignment",
            "group": "layerAssignment",
            "is_device_based": false,
            "id_type": "userID",
            "is_experiment_active": true,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "is_stripe_manage_subscription_link_enabled"
            ],
            "secondary_exposures": []
        },
        "799357841": {
            "name": "799357841",
            "value": {
                "signup_allow_phone": false,
                "in_phone_signup_holdout": false,
                "in_signup_allow_phone_hold_out": false
            },
            "rule_id": "layerAssignment",
            "group": "layerAssignment",
            "is_device_based": true,
            "id_type": "stableID",
            "is_experiment_active": true,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "signup_allow_phone",
                "in_phone_signup_holdout",
                "in_signup_allow_phone_hold_out"
            ],
            "secondary_exposures": []
        },
        "857865944": {
            "name": "857865944",
            "value": {
                "is_team_enabled": true,
                "is_yearly_plus_subscription_enabled": false,
                "is_split_between_personal_and_business_enabled": true,
                "is_modal_fullscreen": true,
                "is_v2_toggle_labels_enabled": false,
                "is_bw": false,
                "is_produce_colors": false,
                "is_produce_color_scheme": false,
                "is_mobile_web_toggle_enabled": false,
                "is_enterprise_enabled": false,
                "is_produce_text": false,
                "is_optimized_checkout": false,
                "is_save_stripe_payment_info_enabled": false,
                "is_auto_save_stripe_payment_info_enabled": false,
                "does_manage_my_subscription_link_take_user_to_subscription_settings": true,
                "should_open_cancellation_survey_after_canceling": true,
                "should_cancel_button_take_user_to_stripe": false,
                "should_show_manage_my_subscription_link": true,
                "is_stripe_manage_subscription_link_enabled": true,
                "cancellation_modal_cancel_button_color": "danger",
                "cancellation_modal_go_back_button_color": "secondary",
                "should_show_cp": false,
                "cp_eligibility_months": 3,
                "should_offer_paypal_when_eligible": false
            },
            "rule_id": "layerAssignment",
            "group": "layerAssignment",
            "is_device_based": false,
            "id_type": "userID",
            "is_experiment_active": true,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "does_manage_my_subscription_link_take_user_to_subscription_settings",
                "should_open_cancellation_survey_after_canceling",
                "should_cancel_button_take_user_to_stripe"
            ],
            "secondary_exposures": []
        },
        "857921158": {
            "name": "857921158",
            "value": {
                "is_team_enabled": true,
                "is_yearly_plus_subscription_enabled": false,
                "is_split_between_personal_and_business_enabled": true,
                "is_modal_fullscreen": true,
                "is_v2_toggle_labels_enabled": false,
                "is_bw": false,
                "is_produce_colors": false,
                "is_produce_color_scheme": false,
                "is_mobile_web_toggle_enabled": false,
                "is_enterprise_enabled": false,
                "is_produce_text": false,
                "is_optimized_checkout": false,
                "is_save_stripe_payment_info_enabled": false,
                "is_auto_save_stripe_payment_info_enabled": false,
                "does_manage_my_subscription_link_take_user_to_subscription_settings": true,
                "should_open_cancellation_survey_after_canceling": true,
                "should_cancel_button_take_user_to_stripe": false,
                "should_show_manage_my_subscription_link": true,
                "is_stripe_manage_subscription_link_enabled": true,
                "cancellation_modal_cancel_button_color": "danger",
                "cancellation_modal_go_back_button_color": "secondary",
                "should_show_cp": false,
                "cp_eligibility_months": 3,
                "should_offer_paypal_when_eligible": false
            },
            "rule_id": "launchedGroup",
            "group": "launchedGroup",
            "group_name": "Test #2",
            "is_device_based": false,
            "id_type": "userID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "does_manage_my_subscription_link_take_user_to_subscription_settings",
                "should_open_cancellation_survey_after_canceling",
                "should_cancel_button_take_user_to_stripe"
            ],
            "secondary_exposures": []
        },
        "944459048": {
            "name": "944459048",
            "value": {
                "is_team_enabled": true,
                "is_yearly_plus_subscription_enabled": false,
                "is_split_between_personal_and_business_enabled": true,
                "is_modal_fullscreen": true,
                "is_v2_toggle_labels_enabled": false,
                "is_bw": false,
                "is_produce_colors": false,
                "is_produce_color_scheme": false,
                "is_mobile_web_toggle_enabled": false,
                "is_enterprise_enabled": false,
                "is_produce_text": false,
                "is_optimized_checkout": false,
                "is_save_stripe_payment_info_enabled": false,
                "is_auto_save_stripe_payment_info_enabled": false,
                "does_manage_my_subscription_link_take_user_to_subscription_settings": true,
                "should_open_cancellation_survey_after_canceling": true,
                "should_cancel_button_take_user_to_stripe": false,
                "should_show_manage_my_subscription_link": true,
                "is_stripe_manage_subscription_link_enabled": true,
                "cancellation_modal_cancel_button_color": "danger",
                "cancellation_modal_go_back_button_color": "secondary",
                "should_show_cp": false,
                "cp_eligibility_months": 3,
                "should_offer_paypal_when_eligible": false
            },
            "rule_id": "abandoned",
            "group": "abandoned",
            "is_device_based": false,
            "id_type": "userID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "is_produce_color_scheme"
            ],
            "secondary_exposures": []
        },
        "954359911": {
            "name": "954359911",
            "value": {},
            "rule_id": "layerAssignment",
            "group": "layerAssignment",
            "is_device_based": false,
            "id_type": "userID",
            "is_experiment_active": true,
            "is_user_in_experiment": false,
            "secondary_exposures": []
        },
        "1001765573": {
            "name": "1001765573",
            "value": {},
            "rule_id": "default",
            "group": "default",
            "is_device_based": false,
            "passed": false,
            "id_type": "userID",
            "secondary_exposures": []
        },
        "1075466664": {
            "name": "1075466664",
            "value": {
                "signup_allow_phone": false,
                "in_phone_signup_holdout": false,
                "in_signup_allow_phone_hold_out": false
            },
            "rule_id": "abandoned",
            "group": "abandoned",
            "is_device_based": true,
            "id_type": "stableID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "use_new_phone_ui",
                "put_email_or_phone_after_social"
            ],
            "secondary_exposures": []
        },
        "1090508242": {
            "name": "1090508242",
            "value": {
                "show_nux": false
            },
            "rule_id": "abandoned",
            "group": "abandoned",
            "is_device_based": false,
            "id_type": "userID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "show_nux"
            ],
            "secondary_exposures": []
        },
        "1111393717": {
            "name": "1111393717",
            "value": {
                "signup_allow_phone": false,
                "in_phone_signup_holdout": false,
                "in_signup_allow_phone_hold_out": false
            },
            "rule_id": "abandoned",
            "group": "abandoned",
            "is_device_based": true,
            "id_type": "stableID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "signup_allow_phone"
            ],
            "secondary_exposures": []
        },
        "1131306727": {
            "name": "1131306727",
            "value": {
                "should_offer_paypal": false
            },
            "rule_id": "prestart",
            "group": "prestart",
            "is_device_based": false,
            "id_type": "userID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "should_offer_paypal"
            ],
            "secondary_exposures": []
        },
        "1134365732": {
            "name": "1134365732",
            "value": {
                "is_anon_chat_enabled": true,
                "is_anon_chat_enabled_for_new_users_only": false,
                "is_try_it_first_on_login_page_enabled": false,
                "is_no_auth_welcome_modal_enabled": false,
                "no_auth_soft_rate_limit": 5,
                "no_auth_hard_rate_limit": 1200,
                "should_show_no_auth_signup_banner": true,
                "is_no_auth_welcome_back_modal_enabled": true,
                "is_no_auth_soft_rate_limit_modal_enabled": true,
                "is_no_auth_gpt4o_modal_enabled": false,
                "is_login_primary_button": true,
                "is_desktop_primary_auth_button_on_right": false,
                "is_primary_btn_blue": false,
                "should_show_disclaimer_only_once_per_device": true,
                "is_secondary_banner_button_enabled": false,
                "is_secondary_auth_banner_button_enabled": true,
                "no_auth_banner_signup_rate_limit": 3,
                "composer_text": "ASK_ANYTHING",
                "is_in_composer_text_exp": true,
                "no_auth_upsell_wording": "NO_CHANGE",
                "should_refresh_access_token_error_take_user_to_no_auth": false
            },
            "rule_id": "launchedGroup",
            "group": "launchedGroup",
            "group_name": "Test",
            "is_device_based": false,
            "id_type": "WebAnonymousCookieID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "should_show_disclaimer_only_once_per_device"
            ],
            "secondary_exposures": [
                {
                    "gate": "1MYcsxeNaY9CfBP0KcxNWK1rpf05mfizLToIzwWdTxKNJhTa7DxSnPMOGf0ImQwRD",
                    "gateValue": "false",
                    "ruleID": "disabled"
                },
                {
                    "gate": "2YJN3jch7WmCCA8cl0Ta5AFVwW3QWI6CFvBUw2895pQCrGiqBBVAyPuUtgzOlGzDS",
                    "gateValue": "false",
                    "ruleID": "disabled"
                }
            ]
        },
        "1313199226": {
            "name": "1313199226",
            "value": {
                "is_anon_chat_enabled": true,
                "is_anon_chat_enabled_for_new_users_only": false,
                "is_try_it_first_on_login_page_enabled": false,
                "is_no_auth_welcome_modal_enabled": false,
                "no_auth_soft_rate_limit": 5,
                "no_auth_hard_rate_limit": 1200,
                "should_show_no_auth_signup_banner": true,
                "is_no_auth_welcome_back_modal_enabled": true,
                "is_no_auth_soft_rate_limit_modal_enabled": true,
                "is_no_auth_gpt4o_modal_enabled": false,
                "is_login_primary_button": true,
                "is_desktop_primary_auth_button_on_right": false,
                "is_primary_btn_blue": false,
                "should_show_disclaimer_only_once_per_device": true,
                "is_secondary_banner_button_enabled": false,
                "is_secondary_auth_banner_button_enabled": true,
                "no_auth_banner_signup_rate_limit": 3,
                "composer_text": "ASK_ANYTHING",
                "is_in_composer_text_exp": true,
                "no_auth_upsell_wording": "NO_CHANGE",
                "should_refresh_access_token_error_take_user_to_no_auth": false
            },
            "rule_id": "abandoned",
            "group": "abandoned",
            "is_device_based": false,
            "id_type": "WebAnonymousCookieID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "no_auth_upsell_wording"
            ],
            "secondary_exposures": [
                {
                    "gate": "1MYcsxeNaY9CfBP0KcxNWK1rpf05mfizLToIzwWdTxKNJhTa7DxSnPMOGf0ImQwRD",
                    "gateValue": "false",
                    "ruleID": "disabled"
                },
                {
                    "gate": "2YJN3jch7WmCCA8cl0Ta5AFVwW3QWI6CFvBUw2895pQCrGiqBBVAyPuUtgzOlGzDS",
                    "gateValue": "false",
                    "ruleID": "disabled"
                }
            ]
        },
        "1320801051": {
            "name": "1320801051",
            "value": {
                "hide_new_at_workspace_section": false,
                "hide_section_new_at_workspace": false,
                "gpt_discovery_experiment_enabled": true,
                "popular_at_my_workspace_enabled": false
            },
            "rule_id": "launchedGroup",
            "group": "launchedGroup",
            "group_name": "Treatment",
            "is_device_based": false,
            "id_type": "userID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "gpt_discovery_experiment_enabled"
            ],
            "secondary_exposures": []
        },
        "1434130504": {
            "name": "1434130504",
            "value": {
                "is_team_enabled": true,
                "is_yearly_plus_subscription_enabled": false,
                "is_split_between_personal_and_business_enabled": true,
                "is_modal_fullscreen": true,
                "is_v2_toggle_labels_enabled": false,
                "is_bw": false,
                "is_produce_colors": false,
                "is_produce_color_scheme": false,
                "is_mobile_web_toggle_enabled": false,
                "is_enterprise_enabled": false,
                "is_produce_text": false,
                "is_optimized_checkout": false,
                "is_save_stripe_payment_info_enabled": false,
                "is_auto_save_stripe_payment_info_enabled": false,
                "does_manage_my_subscription_link_take_user_to_subscription_settings": true,
                "should_open_cancellation_survey_after_canceling": true,
                "should_cancel_button_take_user_to_stripe": false,
                "should_show_manage_my_subscription_link": true,
                "is_stripe_manage_subscription_link_enabled": true,
                "cancellation_modal_cancel_button_color": "danger",
                "cancellation_modal_go_back_button_color": "secondary",
                "should_show_cp": false,
                "cp_eligibility_months": 3,
                "should_offer_paypal_when_eligible": false
            },
            "rule_id": "layerAssignment",
            "group": "layerAssignment",
            "is_device_based": false,
            "id_type": "userID",
            "is_experiment_active": true,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "should_show_cp"
            ],
            "secondary_exposures": [
                {
                    "gate": "Q0Hg2JUuz9kxwwT8dtjGkDyt6RS30qfRG30zFFpcexg",
                    "gateValue": "false",
                    "ruleID": "default"
                }
            ]
        },
        "1535013773": {
            "name": "1535013773",
            "value": {
                "enabled": false
            },
            "rule_id": "1WI6pUFxdAZvczWhP9zEHH",
            "group": "1WI6pUFxdAZvczWhP9zEHH",
            "group_name": "Control",
            "is_device_based": false,
            "id_type": "userID",
            "is_experiment_active": true,
            "is_user_in_experiment": true,
            "secondary_exposures": []
        },
        "1630934647": {
            "name": "1630934647",
            "value": {
                "is_team_enabled": true,
                "is_yearly_plus_subscription_enabled": false,
                "is_split_between_personal_and_business_enabled": true,
                "is_modal_fullscreen": true,
                "is_v2_toggle_labels_enabled": false,
                "is_bw": false,
                "is_produce_colors": false,
                "is_produce_color_scheme": false,
                "is_mobile_web_toggle_enabled": false,
                "is_enterprise_enabled": false,
                "is_produce_text": false,
                "is_optimized_checkout": false,
                "is_save_stripe_payment_info_enabled": false,
                "is_auto_save_stripe_payment_info_enabled": false,
                "does_manage_my_subscription_link_take_user_to_subscription_settings": true,
                "should_open_cancellation_survey_after_canceling": true,
                "should_cancel_button_take_user_to_stripe": false,
                "should_show_manage_my_subscription_link": true,
                "is_stripe_manage_subscription_link_enabled": true,
                "cancellation_modal_cancel_button_color": "danger",
                "cancellation_modal_go_back_button_color": "secondary",
                "should_show_cp": false,
                "cp_eligibility_months": 3,
                "should_offer_paypal_when_eligible": false
            },
            "rule_id": "prestart",
            "group": "prestart",
            "is_device_based": false,
            "id_type": "userID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "should_offer_paypal_when_eligible"
            ],
            "secondary_exposures": []
        },
        "1733458943": {
            "name": "1733458943",
            "value": {
                "unified_architecture": false,
                "ux_updates": false
            },
            "rule_id": "prestart",
            "group": "prestart",
            "is_device_based": true,
            "id_type": "stableID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "unified_architecture"
            ],
            "secondary_exposures": []
        },
        "1856338298": {
            "name": "1856338298",
            "value": {
                "forward_to_authapi": true
            },
            "rule_id": "2RO4BOrVWPrsxRUPYNKPLe:override",
            "group": "2RO4BOrVWPrsxRUPYNKPLe:override",
            "group_name": "Test",
            "is_device_based": true,
            "id_type": "stableID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "forward_to_authapi"
            ],
            "secondary_exposures": [
                {
                    "gate": "zHqNgQrPU3HJvSbPrgMlEBeJROZtyHdE1KsiRnejl3t",
                    "gateValue": "true",
                    "ruleID": "3QgLJ91lKIc7VAOjo5SDz7"
                }
            ]
        },
        "1873047021": {
            "name": "1873047021",
            "value": {
                "search_scoring_dyconfig_name": "gizmo_search_score_config"
            },
            "rule_id": "launchedGroup",
            "group": "launchedGroup",
            "group_name": "Control",
            "is_device_based": false,
            "id_type": "userID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "search_scoring_dyconfig_name"
            ],
            "secondary_exposures": []
        },
        "1904139466": {
            "name": "1904139466",
            "value": {
                "is_anon_chat_enabled": true,
                "is_anon_chat_enabled_for_new_users_only": false,
                "is_try_it_first_on_login_page_enabled": false,
                "is_no_auth_welcome_modal_enabled": false,
                "no_auth_soft_rate_limit": 5,
                "no_auth_hard_rate_limit": 1200,
                "should_show_no_auth_signup_banner": true,
                "is_no_auth_welcome_back_modal_enabled": true,
                "is_no_auth_soft_rate_limit_modal_enabled": true,
                "is_no_auth_gpt4o_modal_enabled": false,
                "is_login_primary_button": true,
                "is_desktop_primary_auth_button_on_right": false,
                "is_primary_btn_blue": false,
                "should_show_disclaimer_only_once_per_device": true,
                "is_secondary_banner_button_enabled": false,
                "is_secondary_auth_banner_button_enabled": true,
                "no_auth_banner_signup_rate_limit": 3,
                "composer_text": "ASK_ANYTHING",
                "is_in_composer_text_exp": true,
                "no_auth_upsell_wording": "NO_CHANGE",
                "should_refresh_access_token_error_take_user_to_no_auth": false
            },
            "rule_id": "launchedGroup",
            "group": "launchedGroup",
            "group_name": "Test",
            "is_device_based": false,
            "id_type": "WebAnonymousCookieID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "is_secondary_auth_banner_button_enabled"
            ],
            "secondary_exposures": [
                {
                    "gate": "1MYcsxeNaY9CfBP0KcxNWK1rpf05mfizLToIzwWdTxKNJhTa7DxSnPMOGf0ImQwRD",
                    "gateValue": "false",
                    "ruleID": "disabled"
                },
                {
                    "gate": "2YJN3jch7WmCCA8cl0Ta5AFVwW3QWI6CFvBUw2895pQCrGiqBBVAyPuUtgzOlGzDS",
                    "gateValue": "false",
                    "ruleID": "disabled"
                }
            ]
        },
        "1993441501": {
            "name": "1993441501",
            "value": {
                "is_anon_chat_enabled": true,
                "is_anon_chat_enabled_for_new_users_only": false,
                "is_try_it_first_on_login_page_enabled": false,
                "is_no_auth_welcome_modal_enabled": false,
                "no_auth_soft_rate_limit": 5,
                "no_auth_hard_rate_limit": 1200,
                "should_show_no_auth_signup_banner": true,
                "is_no_auth_welcome_back_modal_enabled": true,
                "is_no_auth_soft_rate_limit_modal_enabled": true,
                "is_no_auth_gpt4o_modal_enabled": false,
                "is_login_primary_button": true,
                "is_desktop_primary_auth_button_on_right": false,
                "is_primary_btn_blue": false,
                "should_show_disclaimer_only_once_per_device": true,
                "is_secondary_banner_button_enabled": false,
                "is_secondary_auth_banner_button_enabled": true,
                "no_auth_banner_signup_rate_limit": 3,
                "composer_text": "ASK_ANYTHING",
                "is_in_composer_text_exp": true,
                "no_auth_upsell_wording": "NO_CHANGE",
                "should_refresh_access_token_error_take_user_to_no_auth": false
            },
            "rule_id": "launchedGroup",
            "group": "launchedGroup",
            "group_name": "ask_anything",
            "is_device_based": false,
            "id_type": "WebAnonymousCookieID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "composer_text",
                "is_in_composer_text_exp"
            ],
            "secondary_exposures": [
                {
                    "gate": "1MYcsxeNaY9CfBP0KcxNWK1rpf05mfizLToIzwWdTxKNJhTa7DxSnPMOGf0ImQwRD",
                    "gateValue": "false",
                    "ruleID": "disabled"
                },
                {
                    "gate": "2YJN3jch7WmCCA8cl0Ta5AFVwW3QWI6CFvBUw2895pQCrGiqBBVAyPuUtgzOlGzDS",
                    "gateValue": "false",
                    "ruleID": "disabled"
                }
            ]
        },
        "2043237793": {
            "name": "2043237793",
            "value": {
                "bucket": "value"
            },
            "rule_id": "launchedGroup",
            "group": "launchedGroup",
            "group_name": "Value",
            "is_device_based": false,
            "id_type": "userID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "secondary_exposures": []
        },
        "2198260923": {
            "name": "2198260923",
            "value": {
                "mobile": true,
                "web": true,
                "greeting_web": false
            },
            "rule_id": "prestart",
            "group": "prestart",
            "is_device_based": false,
            "id_type": "userID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "mobile",
                "web"
            ],
            "secondary_exposures": []
        },
        "2374392509": {
            "name": "2374392509",
            "value": {
                "is_team_enabled": true,
                "is_yearly_plus_subscription_enabled": false,
                "is_split_between_personal_and_business_enabled": true,
                "is_modal_fullscreen": true,
                "is_v2_toggle_labels_enabled": false,
                "is_bw": false,
                "is_produce_colors": false,
                "is_produce_color_scheme": false,
                "is_mobile_web_toggle_enabled": false,
                "is_enterprise_enabled": false,
                "is_produce_text": false,
                "is_optimized_checkout": false,
                "is_save_stripe_payment_info_enabled": false,
                "is_auto_save_stripe_payment_info_enabled": false,
                "does_manage_my_subscription_link_take_user_to_subscription_settings": true,
                "should_open_cancellation_survey_after_canceling": true,
                "should_cancel_button_take_user_to_stripe": false,
                "should_show_manage_my_subscription_link": true,
                "is_stripe_manage_subscription_link_enabled": true,
                "cancellation_modal_cancel_button_color": "danger",
                "cancellation_modal_go_back_button_color": "secondary",
                "should_show_cp": false,
                "cp_eligibility_months": 3,
                "should_offer_paypal_when_eligible": false
            },
            "rule_id": "abandoned",
            "group": "abandoned",
            "is_device_based": false,
            "id_type": "userID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "should_open_cancellation_survey_after_canceling"
            ],
            "secondary_exposures": []
        },
        "2517852108": {
            "name": "2517852108",
            "value": {
                "is_team_enabled": true,
                "is_yearly_plus_subscription_enabled": false,
                "is_split_between_personal_and_business_enabled": true,
                "is_modal_fullscreen": true,
                "is_v2_toggle_labels_enabled": false,
                "is_bw": false,
                "is_produce_colors": false,
                "is_produce_color_scheme": false,
                "is_mobile_web_toggle_enabled": false,
                "is_enterprise_enabled": false,
                "is_produce_text": false,
                "is_optimized_checkout": false,
                "is_save_stripe_payment_info_enabled": false,
                "is_auto_save_stripe_payment_info_enabled": false,
                "does_manage_my_subscription_link_take_user_to_subscription_settings": true,
                "should_open_cancellation_survey_after_canceling": true,
                "should_cancel_button_take_user_to_stripe": false,
                "should_show_manage_my_subscription_link": true,
                "is_stripe_manage_subscription_link_enabled": true,
                "cancellation_modal_cancel_button_color": "danger",
                "cancellation_modal_go_back_button_color": "secondary",
                "should_show_cp": false,
                "cp_eligibility_months": 3,
                "should_offer_paypal_when_eligible": false
            },
            "rule_id": "1R1OTAoPRlIVIbWb9sM0Em",
            "group": "1R1OTAoPRlIVIbWb9sM0Em",
            "group_name": "Control",
            "is_device_based": false,
            "id_type": "userID",
            "is_experiment_active": true,
            "is_user_in_experiment": true,
            "is_in_layer": true,
            "explicit_parameters": [
                "cancellation_modal_cancel_button_color",
                "cancellation_modal_go_back_button_color"
            ],
            "secondary_exposures": []
        },
        "2555105235": {
            "name": "2555105235",
            "value": {
                "use_email_otp": false,
                "signup_cta_copy": "SIGN_UP",
                "login_allow_phone": false,
                "signup_allow_phone": false,
                "forwardToAuthApi": false,
                "use_new_phone_ui": false,
                "in_signup_allow_phone_hold_out": false,
                "use_formatted_national_number": false,
                "continue_with_email_phone_placement": "after_sso"
            },
            "rule_id": "targetingGate",
            "group": "targetingGate",
            "is_device_based": true,
            "id_type": "stableID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "login_allow_phone"
            ],
            "secondary_exposures": [
                {
                    "gate": "6awtte1Jfq7hOd9f3ke289PCbt7MtVTscsHFafcSqNkHuquhAVHVgMCtrne10klw7",
                    "gateValue": "false",
                    "ruleID": "default"
                }
            ]
        },
        "2702664713": {
            "name": "2702664713",
            "value": {
                "is_team_enabled": true,
                "is_yearly_plus_subscription_enabled": false,
                "is_split_between_personal_and_business_enabled": true,
                "is_modal_fullscreen": true,
                "is_v2_toggle_labels_enabled": false,
                "is_bw": false,
                "is_produce_colors": false,
                "is_produce_color_scheme": false,
                "is_mobile_web_toggle_enabled": false,
                "is_enterprise_enabled": false,
                "is_produce_text": false,
                "is_optimized_checkout": false,
                "is_save_stripe_payment_info_enabled": false,
                "is_auto_save_stripe_payment_info_enabled": false,
                "does_manage_my_subscription_link_take_user_to_subscription_settings": true,
                "should_open_cancellation_survey_after_canceling": true,
                "should_cancel_button_take_user_to_stripe": false,
                "should_show_manage_my_subscription_link": true,
                "is_stripe_manage_subscription_link_enabled": true,
                "cancellation_modal_cancel_button_color": "danger",
                "cancellation_modal_go_back_button_color": "secondary",
                "should_show_cp": false,
                "cp_eligibility_months": 3,
                "should_offer_paypal_when_eligible": false
            },
            "rule_id": "prestart",
            "group": "prestart",
            "is_device_based": false,
            "id_type": "userID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "is_optimized_checkout"
            ],
            "secondary_exposures": []
        },
        "2984513236": {
            "name": "2984513236",
            "value": {
                "is_anon_chat_enabled": true,
                "is_anon_chat_enabled_for_new_users_only": false,
                "is_try_it_first_on_login_page_enabled": false,
                "is_no_auth_welcome_modal_enabled": false,
                "no_auth_soft_rate_limit": 5,
                "no_auth_hard_rate_limit": 1200,
                "should_show_no_auth_signup_banner": true,
                "is_no_auth_welcome_back_modal_enabled": true,
                "is_no_auth_soft_rate_limit_modal_enabled": true,
                "is_no_auth_gpt4o_modal_enabled": false,
                "is_login_primary_button": true,
                "is_desktop_primary_auth_button_on_right": false,
                "is_primary_btn_blue": false,
                "should_show_disclaimer_only_once_per_device": true,
                "is_secondary_banner_button_enabled": false,
                "is_secondary_auth_banner_button_enabled": true,
                "no_auth_banner_signup_rate_limit": 3,
                "composer_text": "ASK_ANYTHING",
                "is_in_composer_text_exp": true,
                "no_auth_upsell_wording": "NO_CHANGE",
                "should_refresh_access_token_error_take_user_to_no_auth": false
            },
            "rule_id": "abandoned",
            "group": "abandoned",
            "is_device_based": false,
            "id_type": "WebAnonymousCookieID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "is_anon_chat_enabled"
            ],
            "secondary_exposures": [
                {
                    "gate": "1MYcsxeNaY9CfBP0KcxNWK1rpf05mfizLToIzwWdTxKNJhTa7DxSnPMOGf0ImQwRD",
                    "gateValue": "false",
                    "ruleID": "disabled"
                },
                {
                    "gate": "2YJN3jch7WmCCA8cl0Ta5AFVwW3QWI6CFvBUw2895pQCrGiqBBVAyPuUtgzOlGzDS",
                    "gateValue": "false",
                    "ruleID": "disabled"
                }
            ]
        },
        "2984514166": {
            "name": "2984514166",
            "value": {
                "is_anon_chat_enabled": true,
                "is_anon_chat_enabled_for_new_users_only": false,
                "is_try_it_first_on_login_page_enabled": false,
                "is_no_auth_welcome_modal_enabled": false,
                "no_auth_soft_rate_limit": 5,
                "no_auth_hard_rate_limit": 1200,
                "should_show_no_auth_signup_banner": true,
                "is_no_auth_welcome_back_modal_enabled": true,
                "is_no_auth_soft_rate_limit_modal_enabled": true,
                "is_no_auth_gpt4o_modal_enabled": false,
                "is_login_primary_button": true,
                "is_desktop_primary_auth_button_on_right": false,
                "is_primary_btn_blue": false,
                "should_show_disclaimer_only_once_per_device": true,
                "is_secondary_banner_button_enabled": false,
                "is_secondary_auth_banner_button_enabled": true,
                "no_auth_banner_signup_rate_limit": 3,
                "composer_text": "ASK_ANYTHING",
                "is_in_composer_text_exp": true,
                "no_auth_upsell_wording": "NO_CHANGE",
                "should_refresh_access_token_error_take_user_to_no_auth": false
            },
            "rule_id": "abandoned",
            "group": "abandoned",
            "is_device_based": false,
            "id_type": "WebAnonymousCookieID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "is_anon_chat_enabled"
            ],
            "secondary_exposures": [
                {
                    "gate": "1MYcsxeNaY9CfBP0KcxNWK1rpf05mfizLToIzwWdTxKNJhTa7DxSnPMOGf0ImQwRD",
                    "gateValue": "false",
                    "ruleID": "disabled"
                },
                {
                    "gate": "2YJN3jch7WmCCA8cl0Ta5AFVwW3QWI6CFvBUw2895pQCrGiqBBVAyPuUtgzOlGzDS",
                    "gateValue": "false",
                    "ruleID": "disabled"
                }
            ]
        },
        "2984514169": {
            "name": "2984514169",
            "value": {
                "is_anon_chat_enabled": true,
                "is_anon_chat_enabled_for_new_users_only": false,
                "is_try_it_first_on_login_page_enabled": false,
                "is_no_auth_welcome_modal_enabled": false,
                "no_auth_soft_rate_limit": 5,
                "no_auth_hard_rate_limit": 1200,
                "should_show_no_auth_signup_banner": true,
                "is_no_auth_welcome_back_modal_enabled": true,
                "is_no_auth_soft_rate_limit_modal_enabled": true,
                "is_no_auth_gpt4o_modal_enabled": false,
                "is_login_primary_button": true,
                "is_desktop_primary_auth_button_on_right": false,
                "is_primary_btn_blue": false,
                "should_show_disclaimer_only_once_per_device": true,
                "is_secondary_banner_button_enabled": false,
                "is_secondary_auth_banner_button_enabled": true,
                "no_auth_banner_signup_rate_limit": 3,
                "composer_text": "ASK_ANYTHING",
                "is_in_composer_text_exp": true,
                "no_auth_upsell_wording": "NO_CHANGE",
                "should_refresh_access_token_error_take_user_to_no_auth": false
            },
            "rule_id": "layerAssignment",
            "group": "layerAssignment",
            "is_device_based": false,
            "id_type": "WebAnonymousCookieID",
            "is_experiment_active": true,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "is_anon_chat_enabled"
            ],
            "secondary_exposures": [
                {
                    "gate": "1MYcsxeNaY9CfBP0KcxNWK1rpf05mfizLToIzwWdTxKNJhTa7DxSnPMOGf0ImQwRD",
                    "gateValue": "false",
                    "ruleID": "disabled"
                },
                {
                    "gate": "2YJN3jch7WmCCA8cl0Ta5AFVwW3QWI6CFvBUw2895pQCrGiqBBVAyPuUtgzOlGzDS",
                    "gateValue": "false",
                    "ruleID": "disabled"
                }
            ]
        },
        "3139879343": {
            "name": "3139879343",
            "value": {
                "is_anon_chat_enabled": true,
                "is_anon_chat_enabled_for_new_users_only": false,
                "is_try_it_first_on_login_page_enabled": false,
                "is_no_auth_welcome_modal_enabled": false,
                "no_auth_soft_rate_limit": 5,
                "no_auth_hard_rate_limit": 1200,
                "should_show_no_auth_signup_banner": true,
                "is_no_auth_welcome_back_modal_enabled": true,
                "is_no_auth_soft_rate_limit_modal_enabled": true,
                "is_no_auth_gpt4o_modal_enabled": false,
                "is_login_primary_button": true,
                "is_desktop_primary_auth_button_on_right": false,
                "is_primary_btn_blue": false,
                "should_show_disclaimer_only_once_per_device": true,
                "is_secondary_banner_button_enabled": false,
                "is_secondary_auth_banner_button_enabled": true,
                "no_auth_banner_signup_rate_limit": 3,
                "composer_text": "ASK_ANYTHING",
                "is_in_composer_text_exp": true,
                "no_auth_upsell_wording": "NO_CHANGE",
                "should_refresh_access_token_error_take_user_to_no_auth": false
            },
            "rule_id": "abandoned",
            "group": "abandoned",
            "is_device_based": false,
            "id_type": "WebAnonymousCookieID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "no_auth_banner_signup_rate_limit",
                "no_auth_soft_rate_limit"
            ],
            "secondary_exposures": [
                {
                    "gate": "1MYcsxeNaY9CfBP0KcxNWK1rpf05mfizLToIzwWdTxKNJhTa7DxSnPMOGf0ImQwRD",
                    "gateValue": "false",
                    "ruleID": "disabled"
                },
                {
                    "gate": "2YJN3jch7WmCCA8cl0Ta5AFVwW3QWI6CFvBUw2895pQCrGiqBBVAyPuUtgzOlGzDS",
                    "gateValue": "false",
                    "ruleID": "disabled"
                }
            ]
        },
        "3199394620": {
            "name": "3199394620",
            "value": {
                "unified_architecture": false,
                "ux_updates": false
            },
            "rule_id": "prestart",
            "group": "prestart",
            "is_device_based": true,
            "id_type": "stableID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "unified_architecture",
                "ux_updates"
            ],
            "secondary_exposures": [
                {
                    "gate": "1mnTptEb3cfCGnADjYLU3dvmS4TiXTuUlDWZ6LYxilpNv8zjGSCkFN7T0np2Etj5Q",
                    "gateValue": "false",
                    "ruleID": "default"
                }
            ]
        },
        "3600867559": {
            "name": "3600867559",
            "value": {
                "is_anon_chat_enabled": true,
                "is_anon_chat_enabled_for_new_users_only": false,
                "is_try_it_first_on_login_page_enabled": false,
                "is_no_auth_welcome_modal_enabled": false,
                "no_auth_soft_rate_limit": 5,
                "no_auth_hard_rate_limit": 1200,
                "should_show_no_auth_signup_banner": true,
                "is_no_auth_welcome_back_modal_enabled": true,
                "is_no_auth_soft_rate_limit_modal_enabled": true,
                "is_no_auth_gpt4o_modal_enabled": false,
                "is_login_primary_button": true,
                "is_desktop_primary_auth_button_on_right": false,
                "is_primary_btn_blue": false,
                "should_show_disclaimer_only_once_per_device": true,
                "is_secondary_banner_button_enabled": false,
                "is_secondary_auth_banner_button_enabled": true,
                "no_auth_banner_signup_rate_limit": 3,
                "composer_text": "ASK_ANYTHING",
                "is_in_composer_text_exp": true,
                "no_auth_upsell_wording": "NO_CHANGE",
                "should_refresh_access_token_error_take_user_to_no_auth": false
            },
            "rule_id": "abandoned",
            "group": "abandoned",
            "is_device_based": false,
            "id_type": "WebAnonymousCookieID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "composer_text",
                "is_in_composer_text_exp"
            ],
            "secondary_exposures": [
                {
                    "gate": "1MYcsxeNaY9CfBP0KcxNWK1rpf05mfizLToIzwWdTxKNJhTa7DxSnPMOGf0ImQwRD",
                    "gateValue": "false",
                    "ruleID": "disabled"
                },
                {
                    "gate": "2YJN3jch7WmCCA8cl0Ta5AFVwW3QWI6CFvBUw2895pQCrGiqBBVAyPuUtgzOlGzDS",
                    "gateValue": "false",
                    "ruleID": "disabled"
                }
            ]
        },
        "3905879930": {
            "name": "3905879930",
            "value": {},
            "rule_id": "targetingGate",
            "group": "targetingGate",
            "is_device_based": false,
            "id_type": "userID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "secondary_exposures": [
                {
                    "gate": "goVs6CADZaeKcunWSVtO7QGhFDKCb6loFuacieezIh6",
                    "gateValue": "false",
                    "ruleID": "default"
                }
            ]
        },
        "3972089454": {
            "name": "3972089454",
            "value": {
                "search_scoring_dyconfig_name": "gizmo_search_score_config"
            },
            "rule_id": "abandoned",
            "group": "abandoned",
            "is_device_based": false,
            "id_type": "userID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "search_scoring_dyconfig_name"
            ],
            "secondary_exposures": []
        },
        "4108879175": {
            "name": "4108879175",
            "value": {
                "is_team_enabled": true,
                "is_yearly_plus_subscription_enabled": false,
                "is_split_between_personal_and_business_enabled": true,
                "is_modal_fullscreen": true,
                "is_v2_toggle_labels_enabled": false,
                "is_bw": false,
                "is_produce_colors": false,
                "is_produce_color_scheme": false,
                "is_mobile_web_toggle_enabled": false,
                "is_enterprise_enabled": false,
                "is_produce_text": false,
                "is_optimized_checkout": false,
                "is_save_stripe_payment_info_enabled": false,
                "is_auto_save_stripe_payment_info_enabled": false,
                "does_manage_my_subscription_link_take_user_to_subscription_settings": true,
                "should_open_cancellation_survey_after_canceling": true,
                "should_cancel_button_take_user_to_stripe": false,
                "should_show_manage_my_subscription_link": true,
                "is_stripe_manage_subscription_link_enabled": true,
                "cancellation_modal_cancel_button_color": "danger",
                "cancellation_modal_go_back_button_color": "secondary",
                "should_show_cp": false,
                "cp_eligibility_months": 3,
                "should_offer_paypal_when_eligible": false
            },
            "rule_id": "launchedGroup",
            "group": "launchedGroup",
            "group_name": "Control",
            "is_device_based": false,
            "id_type": "userID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "is_save_stripe_payment_info_enabled"
            ],
            "secondary_exposures": []
        },
        "4114634655": {
            "name": "4114634655",
            "value": {
                "is_anon_chat_enabled": true,
                "is_anon_chat_enabled_for_new_users_only": false,
                "is_try_it_first_on_login_page_enabled": false,
                "is_no_auth_welcome_modal_enabled": false,
                "no_auth_soft_rate_limit": 5,
                "no_auth_hard_rate_limit": 1200,
                "should_show_no_auth_signup_banner": true,
                "is_no_auth_welcome_back_modal_enabled": true,
                "is_no_auth_soft_rate_limit_modal_enabled": true,
                "is_no_auth_gpt4o_modal_enabled": false,
                "is_login_primary_button": true,
                "is_desktop_primary_auth_button_on_right": false,
                "is_primary_btn_blue": false,
                "should_show_disclaimer_only_once_per_device": true,
                "is_secondary_banner_button_enabled": false,
                "is_secondary_auth_banner_button_enabled": true,
                "no_auth_banner_signup_rate_limit": 3,
                "composer_text": "ASK_ANYTHING",
                "is_in_composer_text_exp": true,
                "no_auth_upsell_wording": "NO_CHANGE",
                "should_refresh_access_token_error_take_user_to_no_auth": false
            },
            "rule_id": "layerAssignment",
            "group": "layerAssignment",
            "is_device_based": false,
            "id_type": "WebAnonymousCookieID",
            "is_experiment_active": true,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "no_auth_banner_signup_rate_limit",
                "no_auth_soft_rate_limit"
            ],
            "secondary_exposures": [
                {
                    "gate": "1MYcsxeNaY9CfBP0KcxNWK1rpf05mfizLToIzwWdTxKNJhTa7DxSnPMOGf0ImQwRD",
                    "gateValue": "false",
                    "ruleID": "disabled"
                },
                {
                    "gate": "2YJN3jch7WmCCA8cl0Ta5AFVwW3QWI6CFvBUw2895pQCrGiqBBVAyPuUtgzOlGzDS",
                    "gateValue": "false",
                    "ruleID": "disabled"
                }
            ]
        },
        "4198227845": {
            "name": "4198227845",
            "value": {
                "enabled_for_platform_override": false,
                "enabled_for_platform_new": false,
                "enabled_for_platform_existing": false,
                "enabled_for_chat_override": false,
                "enabled_for_chat_new": false,
                "enabled_for_chat_existing": false
            },
            "rule_id": "default",
            "group": "default",
            "is_device_based": false,
            "passed": false,
            "id_type": "userID",
            "secondary_exposures": []
        },
        "4212609498": {
            "name": "4212609498",
            "value": {
                "is_anon_chat_enabled": true,
                "is_anon_chat_enabled_for_new_users_only": false,
                "is_try_it_first_on_login_page_enabled": false,
                "is_no_auth_welcome_modal_enabled": false,
                "no_auth_soft_rate_limit": 5,
                "no_auth_hard_rate_limit": 1200,
                "should_show_no_auth_signup_banner": true,
                "is_no_auth_welcome_back_modal_enabled": true,
                "is_no_auth_soft_rate_limit_modal_enabled": true,
                "is_no_auth_gpt4o_modal_enabled": false,
                "is_login_primary_button": true,
                "is_desktop_primary_auth_button_on_right": false,
                "is_primary_btn_blue": false,
                "should_show_disclaimer_only_once_per_device": true,
                "is_secondary_banner_button_enabled": false,
                "is_secondary_auth_banner_button_enabled": true,
                "no_auth_banner_signup_rate_limit": 3,
                "composer_text": "ASK_ANYTHING",
                "is_in_composer_text_exp": true,
                "no_auth_upsell_wording": "NO_CHANGE",
                "should_refresh_access_token_error_take_user_to_no_auth": false
            },
            "rule_id": "prestart",
            "group": "prestart",
            "is_device_based": false,
            "id_type": "WebAnonymousCookieID",
            "is_experiment_active": false,
            "is_user_in_experiment": false,
            "is_in_layer": true,
            "explicit_parameters": [
                "should_refresh_access_token_error_take_user_to_no_auth"
            ],
            "secondary_exposures": [
                {
                    "gate": "1MYcsxeNaY9CfBP0KcxNWK1rpf05mfizLToIzwWdTxKNJhTa7DxSnPMOGf0ImQwRD",
                    "gateValue": "false",
                    "ruleID": "disabled"
                },
                {
                    "gate": "2YJN3jch7WmCCA8cl0Ta5AFVwW3QWI6CFvBUw2895pQCrGiqBBVAyPuUtgzOlGzDS",
                    "gateValue": "false",
                    "ruleID": "disabled"
                }
            ]
        }
    },
    "layer_configs": {
        "183390215": {
            "name": "183390215",
            "value": {
                "signup_allow_phone": false,
                "in_phone_signup_holdout": false,
                "in_signup_allow_phone_hold_out": false
            },
            "rule_id": "default",
            "group": "default",
            "is_device_based": true,
            "explicit_parameters": [],
            "secondary_exposures": [],
            "undelegated_secondary_exposures": []
        },
        "190694971": {
            "name": "190694971",
            "value": {
                "show_nux": false
            },
            "rule_id": "default",
            "group": "default",
            "is_device_based": false,
            "explicit_parameters": [],
            "secondary_exposures": [],
            "undelegated_secondary_exposures": []
        },
        "468168202": {
            "name": "468168202",
            "value": {
                "is_team_enabled": true,
                "is_yearly_plus_subscription_enabled": false,
                "is_split_between_personal_and_business_enabled": true,
                "is_modal_fullscreen": true,
                "is_v2_toggle_labels_enabled": false,
                "is_bw": false,
                "is_produce_colors": false,
                "is_produce_color_scheme": false,
                "is_mobile_web_toggle_enabled": false,
                "is_enterprise_enabled": false,
                "is_produce_text": false,
                "is_optimized_checkout": false,
                "is_save_stripe_payment_info_enabled": false,
                "is_auto_save_stripe_payment_info_enabled": false,
                "does_manage_my_subscription_link_take_user_to_subscription_settings": true,
                "should_open_cancellation_survey_after_canceling": true,
                "should_cancel_button_take_user_to_stripe": false,
                "should_show_manage_my_subscription_link": true,
                "is_stripe_manage_subscription_link_enabled": true,
                "cancellation_modal_cancel_button_color": "danger",
                "cancellation_modal_go_back_button_color": "secondary",
                "should_show_cp": false,
                "cp_eligibility_months": 3,
                "should_offer_paypal_when_eligible": false
            },
            "rule_id": "1R1OTAoPRlIVIbWb9sM0Em",
            "group": "1R1OTAoPRlIVIbWb9sM0Em",
            "group_name": "Control",
            "allocated_experiment_name": "2517852108",
            "is_device_based": false,
            "is_experiment_active": true,
            "explicit_parameters": [
                "cancellation_modal_cancel_button_color",
                "cancellation_modal_go_back_button_color"
            ],
            "is_user_in_experiment": true,
            "secondary_exposures": [
                {
                    "gate": "Q0Hg2JUuz9kxwwT8dtjGkDyt6RS30qfRG30zFFpcexg",
                    "gateValue": "false",
                    "ruleID": "default"
                }
            ],
            "undelegated_secondary_exposures": [
                {
                    "gate": "Q0Hg2JUuz9kxwwT8dtjGkDyt6RS30qfRG30zFFpcexg",
                    "gateValue": "false",
                    "ruleID": "default"
                }
            ]
        },
        "1320801051": {
            "name": "1320801051",
            "value": {
                "hide_new_at_workspace_section": false,
                "hide_section_new_at_workspace": false,
                "gpt_discovery_experiment_enabled": true,
                "popular_at_my_workspace_enabled": false
            },
            "rule_id": "default",
            "group": "default",
            "is_device_based": false,
            "explicit_parameters": [],
            "secondary_exposures": [],
            "undelegated_secondary_exposures": []
        },
        "1704793646": {
            "name": "1704793646",
            "value": {
                "greeting_web": false
            },
            "rule_id": "default",
            "group": "default",
            "is_device_based": false,
            "explicit_parameters": [],
            "secondary_exposures": [],
            "undelegated_secondary_exposures": []
        },
        "1780960461": {
            "name": "1780960461",
            "value": {
                "mobile": true,
                "web": true,
                "greeting_web": false
            },
            "rule_id": "default",
            "group": "default",
            "is_device_based": false,
            "explicit_parameters": [],
            "secondary_exposures": [],
            "undelegated_secondary_exposures": []
        },
        "1914829685": {
            "name": "1914829685",
            "value": {
                "forward_to_authapi": true
            },
            "rule_id": "2RO4BOrVWPrsxRUPYNKPLe:override",
            "group": "2RO4BOrVWPrsxRUPYNKPLe:override",
            "group_name": "Test",
            "allocated_experiment_name": "1856338298",
            "is_device_based": true,
            "is_experiment_active": false,
            "explicit_parameters": [
                "forward_to_authapi"
            ],
            "is_user_in_experiment": false,
            "secondary_exposures": [
                {
                    "gate": "zHqNgQrPU3HJvSbPrgMlEBeJROZtyHdE1KsiRnejl3t",
                    "gateValue": "true",
                    "ruleID": "3QgLJ91lKIc7VAOjo5SDz7"
                }
            ],
            "undelegated_secondary_exposures": [
                {
                    "gate": "zHqNgQrPU3HJvSbPrgMlEBeJROZtyHdE1KsiRnejl3t",
                    "gateValue": "true",
                    "ruleID": "3QgLJ91lKIc7VAOjo5SDz7"
                }
            ]
        },
        "3590606857": {
            "name": "3590606857",
            "value": {
                "should_offer_paypal": false
            },
            "rule_id": "default",
            "group": "default",
            "is_device_based": false,
            "explicit_parameters": [],
            "secondary_exposures": [],
            "undelegated_secondary_exposures": []
        },
        "3637408529": {
            "name": "3637408529",
            "value": {
                "is_anon_chat_enabled": true,
                "is_anon_chat_enabled_for_new_users_only": false,
                "is_try_it_first_on_login_page_enabled": false,
                "is_no_auth_welcome_modal_enabled": false,
                "no_auth_soft_rate_limit": 5,
                "no_auth_hard_rate_limit": 1200,
                "should_show_no_auth_signup_banner": true,
                "is_no_auth_welcome_back_modal_enabled": true,
                "is_no_auth_soft_rate_limit_modal_enabled": true,
                "is_no_auth_gpt4o_modal_enabled": false,
                "is_login_primary_button": true,
                "is_desktop_primary_auth_button_on_right": false,
                "is_primary_btn_blue": false,
                "should_show_disclaimer_only_once_per_device": true,
                "is_secondary_banner_button_enabled": false,
                "is_secondary_auth_banner_button_enabled": true,
                "no_auth_banner_signup_rate_limit": 3,
                "composer_text": "ASK_ANYTHING",
                "is_in_composer_text_exp": true,
                "no_auth_upsell_wording": "NO_CHANGE",
                "should_refresh_access_token_error_take_user_to_no_auth": false
            },
            "rule_id": "default",
            "group": "default",
            "is_device_based": false,
            "explicit_parameters": [],
            "secondary_exposures": [
                {
                    "gate": "1MYcsxeNaY9CfBP0KcxNWK1rpf05mfizLToIzwWdTxKNJhTa7DxSnPMOGf0ImQwRD",
                    "gateValue": "false",
                    "ruleID": "disabled"
                },
                {
                    "gate": "2YJN3jch7WmCCA8cl0Ta5AFVwW3QWI6CFvBUw2895pQCrGiqBBVAyPuUtgzOlGzDS",
                    "gateValue": "false",
                    "ruleID": "disabled"
                }
            ],
            "undelegated_secondary_exposures": [
                {
                    "gate": "1MYcsxeNaY9CfBP0KcxNWK1rpf05mfizLToIzwWdTxKNJhTa7DxSnPMOGf0ImQwRD",
                    "gateValue": "false",
                    "ruleID": "disabled"
                },
                {
                    "gate": "2YJN3jch7WmCCA8cl0Ta5AFVwW3QWI6CFvBUw2895pQCrGiqBBVAyPuUtgzOlGzDS",
                    "gateValue": "false",
                    "ruleID": "disabled"
                }
            ]
        },
        "3972089454": {
            "name": "3972089454",
            "value": {
                "search_scoring_dyconfig_name": "gizmo_search_score_config"
            },
            "rule_id": "default",
            "group": "default",
            "is_device_based": false,
            "explicit_parameters": [],
            "secondary_exposures": [],
            "undelegated_secondary_exposures": []
        },
        "3993182790": {
            "name": "3993182790",
            "value": {
                "unified_architecture": false,
                "ux_updates": false
            },
            "rule_id": "default",
            "group": "default",
            "is_device_based": true,
            "explicit_parameters": [],
            "secondary_exposures": [
                {
                    "gate": "1mnTptEb3cfCGnADjYLU3dvmS4TiXTuUlDWZ6LYxilpNv8zjGSCkFN7T0np2Etj5Q",
                    "gateValue": "false",
                    "ruleID": "default"
                }
            ],
            "undelegated_secondary_exposures": [
                {
                    "gate": "1mnTptEb3cfCGnADjYLU3dvmS4TiXTuUlDWZ6LYxilpNv8zjGSCkFN7T0np2Etj5Q",
                    "gateValue": "false",
                    "ruleID": "default"
                }
            ]
        }
    },
    "sdkParams": {},
    "has_updates": true,
    "generator": "scrapi-nest",
    "time": 1742271926496,
    "company_lcut": 1742271926496,
    "evaluated_keys": {
        "userID": "user-chatgpt",
        "stableID": "27950cca-2360-4e1e-bcd1-db1ee70681d4",
        "customIDs": {
            "WebAnonymousCookieID": "27950cca-2360-4e1e-bcd1-db1ee70681d4",
            "DeviceId": "27950cca-2360-4e1e-bcd1-db1ee70681d4",
            "stableID": "27950cca-2360-4e1e-bcd1-db1ee70681d4",
            "workspace_id": "e8d5afdb-3e65-49ee-8dd2-3877339119d7",
            "account_id": "e8d5afdb-3e65-49ee-8dd2-3877339119d7"
        }
    },
    "hash_used": "djb2",
    "deleted_configs": [
        "1390103565",
        "2575082611",
        "2448897170",
        "371687827",
        "231139217",
        "2589445780",
        "445576012",
        "512236437",
        "3928552611",
        "371721054",
        "3421878454",
        "3386908574",
        "1247232080",
        "2361543803",
        "1068309763",
        "78204706",
        "2676536746",
        "1484522743",
        "3634317850",
        "1407258318",
        "2107429215",
        "3830678688",
        "2191861063",
        "2696038360",
        "1497835938",
        "4226300198",
        "2226851494",
        "4235358925",
        "2167451161",
        "3627190046",
        "267978820",
        "3210224730",
        "1394210961",
        "3885092557",
        "263254725",
        "1504144331",
        "171976366",
        "2299461393",
        "129034157",
        "2868523993",
        "2326853553",
        "3990680520",
        "3063143677",
        "1526046703",
        "2900623391",
        "2946283763",
        "3555200305",
        "2644341027",
        "3776765323",
        "3478422528",
        "1323864366",
        "1766448505",
        "3426130756",
        "2475537204",
        "185164171",
        "2021068459",
        "3859431105",
        "1434459852",
        "3849803907",
        "1644309475",
        "369514126",
        "2836600060",
        "13797217",
        "491992885",
        "2569318737",
        "2228749251",
        "1754698003",
        "1109194859",
        "626245504",
        "3787284482",
        "1361798382",
        "2904148266",
        "4238098366",
        "334552888",
        "2151594028",
        "3546091826",
        "2734082263",
        "3877895946",
        "789610106",
        "1543556447",
        "743208805",
        "1252598780",
        "2865808999",
        "931863260",
        "3302826100",
        "3600868574",
        "1341377011",
        "3844561858",
        "1634099741",
        "3887474185",
        "4150007454",
        "346717229",
        "2806674928",
        "2745713066",
        "722114564",
        "3833667310",
        "4258037029",
        "600125444",
        "791314650",
        "840322333",
        "2069832539",
        "2831545714",
        "1911311648",
        "1953161674",
        "3009072603",
        "2103386956",
        "3745329343",
        "544921198",
        "3826312907",
        "997492548",
        "3776179319",
        "3557009640",
        "897990114",
        "1078002878",
        "4160849695",
        "3668491653",
        "807078778",
        "4200345585",
        "2587885180",
        "3044243918",
        "408530362",
        "3141750138",
        "3293528406",
        "3414761123",
        "1270393229",
        "2079119630",
        "1815344829",
        "362957797",
        "2944805175",
        "689325620",
        "1800916418",
        "1373492259",
        "2163922494",
        "889999301",
        "142383631",
        "3686686209",
        "1491258953",
        "165539775",
        "4235555517",
        "4112526638",
        "1027677163",
        "1853319080",
        "936912515",
        "3696191464",
        "3004943626",
        "3680596162",
        "2857199182",
        "933762508",
        "173134873",
        "4099247432",
        "2086327825",
        "2744401118",
        "2043601735",
        "2721666452",
        "1904371102",
        "2308135002",
        "3104943036",
        "1592914327",
        "2313663436",
        "1448139655",
        "2474025427",
        "3414628831",
        "1326827311",
        "2691550620",
        "122146401",
        "2102978904",
        "1392175073",
        "3139822792",
        "3479238994",
        "874826900",
        "2889974786",
        "860977884",
        "3909732397",
        "2063306328",
        "2584489610",
        "4075022488",
        "3692717849",
        "1744561949",
        "616793104",
        "3856487833",
        "246105906",
        "2122749812",
        "510792679",
        "988265714",
        "3950699162",
        "2833260066",
        "1283123390",
        "3875499440",
        "1120524374",
        "3028999796",
        "238254442",
        "4078319092",
        "2489958943",
        "2532167534",
        "1207634987",
        "269739554",
        "1181805574",
        "3114729735",
        "1382707067",
        "489012373",
        "2977246112",
        "360570504",
        "2496542262",
        "936912516",
        "1722216567",
        "4099247435",
        "3858606627",
        "1888398771",
        "3204397950",
        "684792819",
        "3195941170",
        "937023189",
        "1131079357",
        "1362610566",
        "2022823239",
        "4136200992",
        "1263093047",
        "1382196042",
        "4173038365",
        "252739815",
        "127830638",
        "1742182237",
        "2819087390",
        "710622791",
        "3545236807",
        "3262320410",
        "2982982268",
        "2977246113",
        "2906360664",
        "658151507",
        "2086327828",
        "2951988891",
        "1722270325",
        "2335022111",
        "4064089730",
        "3965095975",
        "2974319823",
        "333544466",
        "3439319096",
        "803354222",
        "3696192424",
        "3858888323",
        "3985014428",
        "1264893289",
        "1656595405",
        "1309104503",
        "910136020",
        "3983890351",
        "3167296343",
        "4062283867",
        "1273026129",
        "4099247434",
        "1304530232",
        "3337227835",
        "3520386402",
        "4284548445",
        "918628452",
        "4099247433",
        "3810796077",
        "4019426777",
        "2254527244",
        "2981469176",
        "3337227834",
        "1019445586",
        "251653059",
        "143320166",
        "1353278451",
        "2456169017",
        "2086327823",
        "773477171",
        "2366675868",
        "3749616366",
        "186540455",
        "3100231510",
        "502942686",
        "3524986926",
        "4041187099",
        "2404495786",
        "4104646289",
        "2490381041",
        "699858154",
        "540250301",
        "42668467",
        "167998424",
        "910136945",
        "3704756040",
        "3208362696",
        "427795613",
        "2279507235",
        "1728365885",
        "2785524107",
        "3753722451",
        "3964446813",
        "806557304",
        "1951095198",
        "2552293827",
        "2680232212",
        "2259761596",
        "614480885",
        "4117888889",
        "2193517106",
        "3308149622",
        "2112835427",
        "1189675352",
        "508901778",
        "1564285176",
        "1231416126",
        "1431733390",
        "736606228",
        "2527626883",
        "2155353011",
        "2074661119",
        "639141176",
        "3624751286",
        "939083602",
        "1405505080",
        "2612950201",
        "1387132551",
        "1256483613",
        "3841436109",
        "20890357",
        "841565749",
        "1203330305",
        "4141354264",
        "1241505683",
        "2513620565",
        "828931202",
        "173134875",
        "654418794",
        "2205850480",
        "671229003",
        "4164810719",
        "4099247430",
        "4185924330",
        "2086327826",
        "452557099",
        "645332024",
        "526342062",
        "2341023900",
        "1452825330",
        "1473880084",
        "2153386174",
        "173134872",
        "96354",
        "154448600",
        "772166296",
        "3888188598",
        "1534160746",
        "518641445",
        "2570720952",
        "1741620534",
        "3456775675",
        "3575751813",
        "3571873296",
        "793404768",
        "689223123",
        "2656608766",
        "1672203265",
        "3746587157",
        "196785869",
        "1438119473",
        "1026825007",
        "2086327822",
        "4071947856",
        "1262968038",
        "416594791",
        "4066942513",
        "3255127320",
        "1789092272",
        "329704007",
        "3341282304",
        "4031494637",
        "2603795322",
        "3325428717",
        "121431657",
        "1199414600",
        "1579373282",
        "2854158819",
        "1237117127",
        "1593328947",
        "3089881936",
        "347297619",
        "132230695",
        "2465588955",
        "3799681642",
        "1141168662",
        "788731483",
        "3830787551",
        "2224223921",
        "587848768",
        "3086967908",
        "1118454899",
        "659845366",
        "2635324712",
        "458433799",
        "2194900962",
        "663944111",
        "4122621429",
        "2056110467",
        "2796511136",
        "674738377",
        "958318952",
        "183619137",
        "91809276",
        "1876351858",
        "4099247431",
        "4109723561",
        "2086327827",
        "1866120232",
        "1246763841",
        "2461703898",
        "3223628778",
        "2173734711",
        "173134874",
        "1556331768",
        "49442974",
        "2542702385",
        "1790127294",
        "3140667314",
        "3440875892",
        "2520913821",
        "868428469",
        "3628232487",
        "4281213358",
        "3832679268",
        "2086327824",
        "1161384477",
        "592326811",
        "2170627071",
        "328234238",
        "2468846433",
        "4139716814",
        "3016364592",
        "1147161329",
        "3564567237",
        "1886161826",
        "1873241832",
        "1626713617",
        "2834969824",
        "1781654241",
        "3824529935",
        "933762507",
        "4150595730",
        "1081674096",
        "4139716815",
        "1207353291",
        "3552563404",
        "571126308",
        "4194393919",
        "173134877",
        "3503253342",
        "3470827922",
        "910135063",
        "580850531",
        "1026802007",
        "231966377",
        "984342884",
        "1820150172",
        "3343752312",
        "3777619639",
        "3755478494",
        "2968440937",
        "1888671873",
        "3793051814",
        "1939148701",
        "1730749688",
        "355754482",
        "2086327829",
        "2087198145",
        "2164734301",
        "892910562",
        "3474819199",
        "1359337507",
        "1537901040",
        "125819449",
        "2461999611",
        "2266181706",
        "2592050760",
        "49104512",
        "3079150901",
        "269438437",
        "204725817",
        "1540817272",
        "2266181705",
        "2425986592",
        "2817154820",
        "2193920885",
        "1151275866",
        "3747646381",
        "950483154",
        "1147161334",
        "2959937706",
        "1498058806",
        "668618468",
        "2027327391",
        "3514586323",
        "2212397902",
        "327512538",
        "2347226805",
        "57285155",
        "2078980429",
        "440846121",
        "3824412381",
        "1769846847",
        "3635795107",
        "4233813681",
        "1393406023",
        "3328362042",
        "57285152",
        "2967926258",
        "262105784",
        "1811400833",
        "1811400832",
        "1636376315",
        "483276184",
        "2887220989",
        "3858451706",
        "67917310",
        "2931193563",
        "3926762963",
        "1597735231",
        "2633590172",
        "2935336213",
        "284876979",
        "1819620406",
        "3934461718",
        "1360962070",
        "2081050579",
        "57285153",
        "1467017784",
        "740401231",
        "3953939478",
        "1306787599",
        "57285154",
        "2454696547",
        "3635795108",
        "1037731751",
        "1069869625",
        "1437093541",
        "2437924003",
        "1886159931",
        "424529111",
        "3402030333",
        "2492365303",
        "1714357209",
        "3787156752",
        "3824811631",
        "2565055944",
        "3067483913",
        "1482075650",
        "1886160868",
        "4022173963",
        "776887663",
        "2943109183",
        "4289826811",
        "153882843",
        "2003956917",
        "1807353389",
        "1304485347",
        "2795699406",
        "4041468795",
        "1538182736",
        "1775027759",
        "1920544812",
        "95458899",
        "1222574719",
        "4001431827",
        "1653300972",
        "1224447458",
        "492490381",
        "2577704136",
        "2542984081",
        "4274348958",
        "196107156",
        "1789373968",
        "1541098968",
        "1888398767",
        "1131361053",
        "1148554536",
        "1224450403",
        "2995906607",
        "3288938296",
        "604731618",
        "71115912",
        "1057309371",
        "2282838563",
        "2517586554",
        "143601862",
        "2565055943",
        "1055786249",
        "2602705865",
        "3576122957",
        "182511994",
        "2489752419",
        "1724653472",
        "3211775015",
        "3499936777",
        "1283405086",
        "581132227",
        "3524918106",
        "3475100895",
        "3576033509",
        "788702776",
        "1556890799",
        "1886160805",
        "4027723683",
        "274976969",
        "1247045537",
        "3017708726",
        "2193798802",
        "1140008906",
        "1873398178",
        "3732263847",
        "693612155",
        "2790030541",
        "778770893",
        "2547010721",
        "1681734359",
        "1889444244",
        "1210733873",
        "2344102206",
        "1790408990",
        "1373773955",
        "51285682",
        "2435661912",
        "714387920",
        "2277227811",
        "1494810687",
        "3401960322",
        "3705037736",
        "2874101084",
        "2573836503",
        "545638790",
        "1148554535",
        "2508095400",
        "3384130327",
        "748225021",
        "2577703241",
        "70834216",
        "34037321",
        "1425935614",
        "1370594850",
        "1569114160",
        "2575639872",
        "2853551934",
        "2573320425",
        "3828194927",
        "269720133",
        "1026824105",
        "3659532862",
        "3251509645",
        "788053860",
        "2254808940",
        "328515934",
        "1801198114",
        "3441157588",
        "393446204",
        "2653584728",
        "3224549530",
        "1387414247",
        "1145118933",
        "631216852",
        "2680513908",
        "3525268622",
        "3833435648",
        "2970598858",
        "2321301612",
        "3875781136",
        "2165015997",
        "2683328006",
        "679556947",
        "1915182504",
        "781471532",
        "2042473639",
        "469044348",
        "3105224732",
        "1434980250",
        "1220965678",
        "1392456769",
        "2442332973",
        "1222270874",
        "984093633",
        "2404777482",
        "3345722677",
        "3211772070",
        "502443464",
        "1453107026",
        "357423375",
        "4233887322",
        "1888398768",
        "1209336180",
        "2207860437",
        "142072574",
        "1274452230",
        "2577704137",
        "1564566872",
        "1385040317",
        "2573836505",
        "804723264",
        "1479465465",
        "2990136806",
        "564988237",
        "935151796",
        "246387602",
        "4042188164",
        "1458895032",
        "1728647581",
        "4061843250",
        "233240076",
        "172231013",
        "446139161",
        "3490142221",
        "2275008324",
        "110823954",
        "75427118",
        "1888398770",
        "140029292",
        "190661894",
        "473504715",
        "1512066226",
        "1628427795",
        "1400375316",
        "3242644804",
        "4019488255",
        "2577704135",
        "509183474",
        "3173011745",
        "3830792645",
        "3185519135",
        "3769970440",
        "4094569773",
        "196400364",
        "1886182903",
        "587874559",
        "2273675522",
        "331619538",
        "4203223025",
        "1918066574",
        "1213914474",
        "2659525330",
        "1886160833",
        "1577220573",
        "2786391964",
        "3262602106",
        "671510699",
        "576113947",
        "3708924548",
        "1231556318",
        "1888398769",
        "1720201876",
        "4281195779",
        "4235837213",
        "3296926367",
        "1603708032",
        "675020073",
        "329985703",
        "3777429207",
        "140030067",
        "347413746",
        "3560467995",
        "163108163",
        "2056392163",
        "779052589",
        "2859152004",
        "3877148507",
        "1026794284",
        "4281495054",
        "480204662",
        "121713353",
        "1886161796",
        "1574627219",
        "1814607155",
        "3585425775",
        "1359619203",
        "2285879218",
        "3104847357",
        "279580727",
        "140028145",
        "1377572404",
        "2577703212",
        "1400113149",
        "3801961472",
        "3147244306",
        "3414910527",
        "1148554537",
        "3576124755",
        "4055950881",
        "1907823790",
        "10737088",
        "2854440515",
        "2490662737",
        "3140949010",
        "1651744539",
        "3255409016",
        "2075021576",
        "2827439059",
        "1929897855",
        "196128325",
        "3375116335",
        "2549262211",
        "2366957564",
        "1408179464",
        "320977152",
        "3289002197",
        "3080091891",
        "2577703243",
        "3479520690",
        "2819255508",
        "3854468621",
        "2005437762",
        "2968722633",
        "781460811",
        "2172367547",
        "135072988",
        "253021511",
        "2573836504",
        "875108596",
        "1368768040",
        "3276424238",
        "474224863",
        "299907519",
        "2518381570",
        "1019810976",
        "1626388789",
        "2085342692",
        "271321279",
        "491541983",
        "773758867",
        "2517186110",
        "3902399814",
        "1644781658",
        "424069724",
        "1593196023",
        "1575392146",
        "2502256612",
        "639422872",
        "1479747161",
        "1886317504",
        "2640527669",
        "3606409267",
        "2243916098",
        "946479448",
        "765051472",
        "3905338564",
        "1888398766",
        "2496823958",
        "2349797906",
        "3516217281",
        "3799027570",
        "211804453",
        "2321766861",
        "67995390",
        "2161146287",
        "2833541762",
        "165832429",
        "471685275",
        "736887924",
        "4031776333",
        "3747928077",
        "645090066",
        "2210383791",
        "3771642798",
        "2156366963",
        "541562001",
        "795551887",
        "2562947024",
        "1477559309",
        "2547830396",
        "3345682196",
        "4000626361",
        "1699207239",
        "2651793984",
        "1724652697",
        "268053413",
        "121139917",
        "2653933266",
        "4023528987",
        "3746886559",
        "1229410781",
        "2653922784",
        "720749476",
        "4264856273",
        "3211777960",
        "2696703271",
        "1000830532",
        "1575622182",
        "1626995313",
        "350582806",
        "2194568781",
        "3177108764",
        "3316392749",
        "709466665",
        "2119969849",
        "1574624274",
        "1561303927",
        "3744780369",
        "1108375992",
        "537738935",
        "3710022682",
        "1405654113",
        "1886159907",
        "559596502",
        "3521996631",
        "3372045409",
        "3316101665",
        "604318423",
        "2170908767",
        "1237398823",
        "3344034008",
        "1081955792",
        "1569114161",
        "1836246389",
        "2189775249",
        "205277408",
        "754832920",
        "3995604469",
        "3493453158",
        "3631246315",
        "3631246314",
        "792374093",
        "155526418",
        "4158940316",
        "4131051721",
        "2444640060",
        "1672724387",
        "3346803632",
        "944185945",
        "3498431719",
        "4241591419",
        "3469785973",
        "2666869659",
        "3711519293",
        "534473949",
        "3702936237",
        "3368880061",
        "1139316054",
        "3852502210",
        "944155197",
        "1414340994",
        "2190588033",
        "1228545014",
        "354006082",
        "1631188742",
        "1356066837",
        "920974034",
        "2733258055",
        "3413316835",
        "1356066903",
        "670125278",
        "2058913668",
        "3690399923",
        "2857102965",
        "4100491894",
        "800438672",
        "2819187979",
        "971770036",
        "441462595",
        "3056866376",
        "40222740",
        "253022435",
        "3432043323",
        "1479958612",
        "4130128200",
        "398241605",
        "91210973",
        "775921210",
        "1172594865",
        "1577297060",
        "1807158731",
        "1562119207",
        "482400139",
        "1959440673",
        "188889406",
        "1027254270",
        "3470827921",
        "188889407",
        "2204704146",
        "2866835856",
        "2169600467",
        "2529995229",
        "2370410632",
        "75342665",
        "2576123431",
        "2811617088",
        "2983618523",
        "75342698",
        "2983619448",
        "445876862",
        "3414158273",
        "2983619420",
        "2891383514",
        "3591528489",
        "457207600",
        "2983619449",
        "1887692490",
        "2297342424",
        "598205782",
        "1072371496",
        "2983619414",
        "3340924940",
        "1312738797",
        "3414159263",
        "2983618516",
        "2891383548",
        "2370410634",
        "149784183",
        "2983619445",
        "3112769237",
        "2312480955",
        "3758098899",
        "3468755820",
        "2980740905",
        "4252260557",
        "2745768542",
        "857920287",
        "3022018663",
        "2745768482",
        "2745767556",
        "186644516",
        "2745768485",
        "4252261542",
        "4252260612",
        "2745770409",
        "2395732115",
        "2919046882",
        "4274128667",
        "1944896067",
        "2473911998",
        "1258876927",
        "3836331810",
        "781292183",
        "3797314318",
        "1920410330",
        "871480661",
        "4201086940",
        "302644286",
        "2919046885",
        "2165032501",
        "302644285",
        "3748755988",
        "3470386637",
        "2109174912",
        "3904663959",
        "2224417057",
        "455651475",
        "3251182224",
        "1127621229",
        "2163185459",
        "283967729",
        "2254842980",
        "3512098818",
        "3391638664",
        "140420593",
        "974692666",
        "4022926669",
        "194120519",
        "2455784654",
        "2221287009",
        "2467649520",
        "3557866671",
        "632229184",
        "3610327527",
        "3888212137",
        "1585693585",
        "225643989",
        "731172876",
        "2351824930",
        "3191323006",
        "52701554",
        "3776940125",
        "2335908105",
        "1826752009",
        "407426048",
        "3001143868",
        "3286693234",
        "1611604388",
        "1930243550",
        "2198947464",
        "3701102259",
        "3148769572",
        "2424412160",
        "4278549316",
        "2154400378",
        "2773808134",
        "2212993488",
        "1826752008",
        "2335908103",
        "2567214617",
        "1826752011",
        "3021134725",
        "1826752010",
        "676760469",
        "1365636985",
        "2424145028",
        "2335908104",
        "410076910",
        "3033637839",
        "1861953028",
        "3124034761",
        "3050783378",
        "3519108196",
        "1013668396",
        "2350183941",
        "2744100866",
        "4082431641",
        "877308352",
        "3565105566",
        "3565105567",
        "3050556939",
        "2383116751",
        "1263441356",
        "1900034167",
        "1526364077",
        "939566462",
        "1712074056",
        "1226621023",
        "4056381013",
        "3892547366",
        "3145269933",
        "560880628",
        "1647934161",
        "636143343",
        "379069945",
        "2775330254",
        "3892578173",
        "1832767480",
        "776337230",
        "1286290341",
        "2484549348",
        "1948350217",
        "2829857051",
        "2005371928",
        "578569747",
        "238113716",
        "1693842125",
        "1487119164",
        "2542663807",
        "3114616911",
        "1698817085",
        "1103440290",
        "52701548",
        "739785984",
        "3971057491",
        "2144828235",
        "743318407"
    ],
    "deleted_gates": [
        "1958191845",
        "2943315548",
        "2385696481",
        "3612584454",
        "3240610717",
        "918417280",
        "1083964350",
        "2003935602",
        "3266740164",
        "782207126",
        "835142948",
        "1342960564",
        "2904602376",
        "653821218",
        "4265747962",
        "2135216650",
        "188174734",
        "135448051",
        "2951433295",
        "4086387128",
        "3260739881",
        "1149337106",
        "1514582782",
        "1439437954",
        "4166126008",
        "891514942",
        "2812157128",
        "3488920647",
        "2781425969",
        "2690524466",
        "1545665756",
        "3243803551",
        "3685290481",
        "1471030749",
        "1713155723",
        "2643492900",
        "901764770",
        "3565379860",
        "2707795562",
        "2084678120",
        "1686674561",
        "3345181713",
        "3148583717",
        "2400776896",
        "3861593998",
        "550432558",
        "748668726",
        "29567930",
        "3580425399",
        "1301970209",
        "671546469",
        "3954213837",
        "1099561146",
        "3879160712",
        "1414026179",
        "725224950",
        "1616485584",
        "2810078907",
        "2712556596",
        "2603066997",
        "172357831",
        "2892314429",
        "3511310085",
        "1065821961",
        "756735255",
        "3364827183",
        "863053742",
        "3979874914",
        "2935021756",
        "2364036092",
        "720132657",
        "1420961915",
        "1906328919",
        "2957977839",
        "1583484408",
        "2896696430",
        "1246802339",
        "3867222281",
        "3509546837",
        "3270434513",
        "255505914",
        "1377255128",
        "1382475798",
        "4168663601",
        "1768908523",
        "1418526466",
        "1016364891",
        "1330965306",
        "3850857995",
        "2357642855",
        "947285409",
        "59687878",
        "3803286769",
        "1443669600",
        "1363180765",
        "2997426447",
        "3216817468",
        "3566998806",
        "3636405501",
        "2219872076",
        "3656259541",
        "157385954",
        "4031520646",
        "4206189746",
        "2192543539",
        "650432205",
        "1622722902",
        "2972747856",
        "1154002920",
        "1145487139",
        "2304848677",
        "938307553",
        "4151101559",
        "591542420",
        "1760640904",
        "326583325",
        "885770349",
        "4192239497",
        "1927266746",
        "2056761365",
        "1214379119",
        "1487499528",
        "3016192915",
        "222560275",
        "939331103",
        "1410928537",
        "3511639124",
        "2434770139",
        "369853820",
        "4150471562",
        "3759220299",
        "2164931355",
        "4197990588",
        "3282605148",
        "3537869706",
        "4048961623",
        "3535261446",
        "703605568",
        "1042528641",
        "1697830637",
        "2253734401",
        "543404555",
        "1744826997",
        "3750571022",
        "1453220720",
        "1849523248",
        "2648713736",
        "2182478793",
        "3697757679",
        "2933557648",
        "1120715543",
        "1950886794",
        "313097308",
        "1511478076",
        "2672559500"
    ],
    "deleted_layers": [
        "1597818504",
        "3304399934",
        "1803944755",
        "3091960901",
        "3007869417",
        "3675195312",
        "558417232",
        "316719357",
        "1232472996",
        "2118136551",
        "3748767161",
        "4031588851",
        "2995358426",
        "2505516353",
        "3499680977",
        "3620383586",
        "2276550910",
        "2100905711",
        "2149763392",
        "515635819",
        "453021389",
        "3612909111",
        "1832767480",
        "109457"
    ],
    "is_delta": true,
    "derived_fields": {
        "user_agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36 Edg/134.0.0.0",
        "ip": "8.8.8.8",
        "country": "US"
    },
    "checksum": "237282501",
    "checksumV2": "1734546036",
    "hashed_sdk_key_used": "1950608987",
    "deltas_full_response": null,
    "can_record_session": true,
    "session_recording_rate": 1,
    "auto_capture_settings": {
        "disabled_events": {}
    },
    "param_stores": {},
    "target_app_used": "PUBLICLY-VISIBLE-js-client-browser",
    "full_checksum": "1323643497"
}
````

## File: templates/login.html
````html
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>登录</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <style>
        #popup {
            display: none;
        }
    </style>
</head>
<body class="bg-gradient-to-br from-blue-300 via-indigo-300 to-purple-400 min-h-screen flex items-center justify-center">
    <div class="bg-white p-8 rounded-xl shadow-2xl w-96 max-w-md">
        <h2 class="text-3xl font-bold text-center text-gray-800 mb-4">登录</h2>
        <button
            type="button"
            onclick="openPopup()"
            class="w-full bg-gradient-to-r from-indigo-500 to-blue-500 text-white font-bold py-3 px-4 rounded-lg hover:opacity-90 transition duration-200 ease-in-out"
        >
            RefreshToken / AccessToken
        </button>
        <p class="text-xs text-gray-500 text-center mt-4">
            <span class="font-semibold text-gray-800">非
                <span class="font-bold text-indigo-600">RT</span>
                与
                <span class="font-bold text-indigo-600">AT</span>
                的输入，将作为</span>
            <span class="font-bold text-indigo-600">Seed</span>
            <span class="font-semibold text-gray-800">随机抽取后台账号</span>
        </p>
    </div>

    <div id="popup" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
        <div class="bg-white p-8 rounded-xl shadow-lg max-w-lg w-full h-auto">
<!--            <h3 class="text-xl font-bold text-gray-800 mb-4 text-center">请输入您的 Token</h3>-->
            <p class="font-semibold text-gray-800 text-center mb-1">直接点击
                <span class="font-bold text-indigo-600">开始</span>
                进入最近用过的账号</p>
            <label for="popup-input"></label>
            <textarea
                id="popup-input"
                name="token"
                placeholder="RefreshToken / AccessToken / SeedToken"
                class="w-full h-56 px-4 py-4 text-md rounded-md bg-gray-100 border-gray-300 focus:border-indigo-500 focus:bg-white focus:ring-2 focus:ring-indigo-200 text-gray-800 transition duration-200 ease-in-out mb-4"
                style="resize: none;"
            ></textarea>
            <div class="flex justify-center space-x-4">
                <button
                    onclick="submitToken()"
                    class="bg-indigo-500 hover:bg-indigo-600 text-white font-bold py-2 px-4 rounded-lg transition duration-200 ease-in-out"
                >
                    开 始
                </button>
                <button
                    onclick="closePopup()"
                    class="bg-gray-300 hover:bg-gray-400 text-gray-800 font-bold py-2 px-4 rounded-lg transition duration-200 ease-in-out"
                >
                    取 消
                </button>
            </div>
        </div>
    </div>

    <script>
        function openPopup() {
            document.getElementById('popup').style.display = 'flex';
        }

        function closePopup() {
            document.getElementById('popup').style.display = 'none';
        }

        function submitToken() {
            var inputValue = document.getElementById('popup-input').value;
            window.location.href = '/?token=' + inputValue;
            closePopup();
        }
    </script>
</body>
</html>
````

## File: templates/tokens.html
````html
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta content="width=device-width, initial-scale=1.0" name="viewport">
    <title>Tokens 管理</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <script>
        document.addEventListener('DOMContentLoaded', () => {
            const apiPrefix = "{{ api_prefix }}";
            const uploadForm = document.getElementById('uploadForm');
            const clearForm = document.getElementById('clearForm');
            const errorButton = document.getElementById('errorButton');

            if (apiPrefix === "None") {
                uploadForm.action = "/tokens/upload";
                clearForm.action = "/tokens/clear";
                errorButton.dataset.api = "/tokens/error";
            } else {
                uploadForm.action = `/${apiPrefix}/tokens/upload`;
                clearForm.action = `/${apiPrefix}/tokens/clear`;
                errorButton.dataset.api = `/${apiPrefix}/tokens/error`;
            }

            errorButton.addEventListener('click', async () => {
                const response = await fetch(errorButton.dataset.api, {
                    method: 'POST',
                });
                const result = await response.json();
                const errorTokens = result.error_tokens;

                const errorModal = document.getElementById('errorModal');
                const errorModalContent = document.getElementById('errorModalContent');

                errorModalContent.innerHTML = errorTokens.map(token => `<p>${token}</p>`).join('');
                errorModal.classList.remove('hidden');
            });

            document.getElementById('errorModalClose').addEventListener('click', () => {
                document.getElementById('errorModal').classList.add('hidden');
            });

            document.getElementById('errorModalCopy').addEventListener('click', () => {
                const errorModalContent = document.getElementById('errorModalContent');
                const textToCopy = errorModalContent.innerText.replace(/\n\n/g, '\n');
                navigator.clipboard.writeText(textToCopy).then(() => {
                    alert('错误 Tokens 已复制到剪贴板');
                }).catch(err => {
                    alert('复制失败，请手动复制');
                });
            });
        });
    </script>
</head>
<body class="bg-gradient-to-r from-blue-200 via-purple-200 to-pink-200 flex justify-center items-center min-h-screen">
    <div class="bg-white p-10 rounded-lg shadow-2xl w-128 text-center">
        <h1 class="text-4xl font-extrabold text-gray-900 mb-6">Tokens 管理</h1>
        <p class="text-gray-600 mb-4">当前可用 Tokens 数量：<span class="text-blue-600">{{ tokens_count }}</span></p>
        <form class="mb-2" id="uploadForm" method="post">
            <textarea class="w-full p-4 mb-4 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-400 resize-none" name="text" placeholder="一行一个Token，可以是 AccessToken 或 RefreshToken" rows="10"></textarea>
            <p class="text-gray-600 mb-2">注：使用docker时如果挂载了data文件夹则重启后不需要再次上传</p>
            <button class="w-full bg-blue-600 text-white py-3 rounded-md hover:bg-blue-700 transition duration-300 mb-2" type="submit">上传</button>
        </form>
        <button id="errorButton" class="w-full bg-yellow-600 text-white py-3 rounded-md hover:bg-yellow-700 transition duration-200 mt-2">查看错误Tokens</button>
        <p class="text-gray-600 mt-2">点击清空，将会清空上传和错误的 Tokens</p>
        <form id="clearForm" method="post">
            <button class="w-full bg-red-600 text-white py-3 rounded-md hover:bg-red-700 transition duration-300" type="submit">清空Tokens</button>
        </form>
    </div>
    
    <div id="errorModal" class="fixed inset-0 bg-gray-800 bg-opacity-75 flex justify-center items-center hidden">
        <div class="bg-white p-6 rounded-lg shadow-lg w-150">
            <h2 class="text-2xl font-bold mb-4">错误 Tokens</h2>
            <div id="errorModalContent" class="list-disc list-inside text-left mb-4"></div>
            <div class="flex justify-end space-x-4">
                <button id="errorModalCopy" class="bg-green-600 text-white py-2 px-4 rounded-md hover:bg-green-700 transition duration-300">复制</button>
                <button id="errorModalClose" class="bg-red-600 text-white py-2 px-4 rounded-md hover:bg-red-700 transition duration-300">关闭</button>
            </div>
        </div>
    </div>
</body>
</html>
````

## File: utils/Client.py
````python
import random

from curl_cffi.requests import AsyncSession


class Client:
    def __init__(self, proxy=None, timeout=15, verify=True, impersonate='safari15_3'):
        self.proxies = {"http": proxy, "https": proxy}
        self.timeout = timeout
        self.verify = verify

        self.impersonate = impersonate
        # impersonate=self.impersonate

        # self.ja3 = ""
        # self.akamai = ""
        # ja3=self.ja3, akamai=self.akamai
        self.session = AsyncSession(proxies=self.proxies, timeout=self.timeout, impersonate=self.impersonate, verify=self.verify)
        self.session2 = AsyncSession(proxies=self.proxies, timeout=self.timeout, impersonate=self.impersonate, verify=self.verify)

    async def post(self, *args, **kwargs):
        r = await self.session.post(*args, **kwargs)
        return r

    async def post_stream(self, *args, headers=None, cookies=None, **kwargs):
        if self.session:
            headers = headers or self.session.headers
            cookies = cookies or self.session.cookies
        r = await self.session2.post(*args, headers=headers, cookies=cookies, **kwargs)
        return r

    async def get(self, *args, **kwargs):
        r = await self.session.get(*args, **kwargs)
        return r

    async def request(self, *args, **kwargs):
        r = await self.session.request(*args, **kwargs)
        return r

    async def put(self, *args, **kwargs):
        r = await self.session.put(*args, **kwargs)
        return r

    async def close(self):
        if hasattr(self, 'session'):
            try:
                await self.session.close()
                del self.session
            except Exception:
                pass
        if hasattr(self, 'session2'):
            try:
                await self.session2.close()
                del self.session2
            except Exception:
                pass
````

## File: utils/configs.py
````python
import ast
import os

from dotenv import load_dotenv

from utils.Logger import logger

load_dotenv(encoding="ascii")


def is_true(x):
    if isinstance(x, bool):
        return x
    if isinstance(x, str):
        return x.lower() in ['true', '1', 't', 'y', 'yes']
    elif isinstance(x, int):
        return x == 1
    else:
        return False


api_prefix = os.getenv('API_PREFIX', None)
authorization = os.getenv('AUTHORIZATION', '').replace(' ', '')
chatgpt_base_url = os.getenv('CHATGPT_BASE_URL', 'https://chatgpt.com').replace(' ', '')
auth_key = os.getenv('AUTH_KEY', None)
x_sign = os.getenv('X_SIGN', None)

ark0se_token_url = os.getenv('ARK' + 'OSE_TOKEN_URL', '').replace(' ', '')
if not ark0se_token_url:
    ark0se_token_url = os.getenv('ARK0SE_TOKEN_URL', None)
proxy_url = os.getenv('PROXY_URL', '').replace(' ', '')
sentinel_proxy_url = os.getenv('SENTINEL_PROXY_URL', None)
export_proxy_url = os.getenv('EXPORT_PROXY_URL', None)
file_host = os.getenv('FILE_HOST', None)
voice_host = os.getenv('VOICE_HOST', None)
impersonate_list_str = os.getenv('IMPERSONATE', '[]')
user_agents_list_str = os.getenv('USER_AGENTS', '[]')
device_tuple_str = os.getenv('DEVICE_TUPLE', '()')
browser_tuple_str = os.getenv('BROWSER_TUPLE', '()')
platform_tuple_str = os.getenv('PLATFORM_TUPLE', '()')

cf_file_url = os.getenv('CF_FILE_URL', None)
turnstile_solver_url = os.getenv('TURNSTILE_SOLVER_URL', None)

history_disabled = is_true(os.getenv('HISTORY_DISABLED', True))
pow_difficulty = os.getenv('POW_DIFFICULTY', '000032')
retry_times = int(os.getenv('RETRY_TIMES', 3))
conversation_only = is_true(os.getenv('CONVERSATION_ONLY', False))
enable_limit = is_true(os.getenv('ENABLE_LIMIT', True))
upload_by_url = is_true(os.getenv('UPLOAD_BY_URL', False))
check_model = is_true(os.getenv('CHECK_MODEL', False))
scheduled_refresh = is_true(os.getenv('SCHEDULED_REFRESH', False))
random_token = is_true(os.getenv('RANDOM_TOKEN', True))
oai_language = os.getenv('OAI_LANGUAGE', 'zh-CN')

authorization_list = authorization.split(',') if authorization else []
chatgpt_base_url_list = chatgpt_base_url.split(',') if chatgpt_base_url else []
ark0se_token_url_list = ark0se_token_url.split(',') if ark0se_token_url else []
proxy_url_list = proxy_url.split(',') if proxy_url else []
sentinel_proxy_url_list = sentinel_proxy_url.split(',') if sentinel_proxy_url else []
impersonate_list = ast.literal_eval(impersonate_list_str)
user_agents_list = ast.literal_eval(user_agents_list_str)
device_tuple = ast.literal_eval(device_tuple_str)
browser_tuple = ast.literal_eval(browser_tuple_str)
platform_tuple = ast.literal_eval(platform_tuple_str)

enable_gateway = is_true(os.getenv('ENABLE_GATEWAY', False))
auto_seed = is_true(os.getenv('AUTO_SEED', True))
force_no_history = is_true(os.getenv('FORCE_NO_HISTORY', False))
no_sentinel = is_true(os.getenv('NO_SENTINEL', False))

with open('version.txt') as f:
    version = f.read().strip()

logger.info("-" * 60)
logger.info(f"Chat2Api {version} | https://github.com/lanqian528/chat2api")
logger.info("-" * 60)
logger.info("Environment variables:")
logger.info("------------------------- Security -------------------------")
logger.info("API_PREFIX:        " + str(api_prefix))
logger.info("AUTHORIZATION:     " + str(authorization_list))
logger.info("AUTH_KEY:          " + str(auth_key))
logger.info("------------------------- Request --------------------------")
logger.info("CHATGPT_BASE_URL:  " + str(chatgpt_base_url_list))
logger.info("PROXY_URL:         " + str(proxy_url_list))
logger.info("EXPORT_PROXY_URL:  " + str(export_proxy_url))
logger.info("FILE_HOST:     " + str(file_host))
logger.info("VOICE_HOST:    " + str(voice_host))
logger.info("IMPERSONATE:       " + str(impersonate_list))
logger.info("USER_AGENTS:       " + str(user_agents_list))
logger.info("---------------------- Functionality -----------------------")
logger.info("HISTORY_DISABLED:  " + str(history_disabled))
logger.info("POW_DIFFICULTY:    " + str(pow_difficulty))
logger.info("RETRY_TIMES:       " + str(retry_times))
logger.info("CONVERSATION_ONLY: " + str(conversation_only))
logger.info("ENABLE_LIMIT:      " + str(enable_limit))
logger.info("UPLOAD_BY_URL:     " + str(upload_by_url))
logger.info("CHECK_MODEL:       " + str(check_model))
logger.info("SCHEDULED_REFRESH: " + str(scheduled_refresh))
logger.info("RANDOM_TOKEN:      " + str(random_token))
logger.info("OAI_LANGUAGE:      " + str(oai_language))
logger.info("------------------------- Gateway --------------------------")
logger.info("ENABLE_GATEWAY:    " + str(enable_gateway))
logger.info("AUTO_SEED:         " + str(auto_seed))
logger.info("FORCE_NO_HISTORY: " + str(force_no_history))
logger.info("-" * 60)
````

## File: utils/globals.py
````python
import json
import os

import utils.configs as configs
from utils.Logger import logger

DATA_FOLDER = "data"
TOKENS_FILE = os.path.join(DATA_FOLDER, "token.txt")
REFRESH_MAP_FILE = os.path.join(DATA_FOLDER, "refresh_map.json")
ERROR_TOKENS_FILE = os.path.join(DATA_FOLDER, "error_token.txt")
WSS_MAP_FILE = os.path.join(DATA_FOLDER, "wss_map.json")
FP_FILE = os.path.join(DATA_FOLDER, "fp_map.json")
SEED_MAP_FILE = os.path.join(DATA_FOLDER, "seed_map.json")
CONVERSATION_MAP_FILE = os.path.join(DATA_FOLDER, "conversation_map.json")

count = 0
token_list = []
error_token_list = []
refresh_map = {}
wss_map = {}
fp_map = {}
seed_map = {}
conversation_map = {}
impersonate_list = [
    "chrome99",
    "chrome100",
    "chrome101",
    "chrome104",
    "chrome107",
    "chrome110",
    "chrome116",
    "chrome119",
    "chrome120",
    "chrome123",
    "edge99",
    "edge101",
] if not configs.impersonate_list else configs.impersonate_list

if not os.path.exists(DATA_FOLDER):
    os.makedirs(DATA_FOLDER)

if os.path.exists(REFRESH_MAP_FILE):
    with open(REFRESH_MAP_FILE, "r") as f:
        try:
            refresh_map = json.load(f)
        except:
            refresh_map = {}
else:
    refresh_map = {}

if os.path.exists(WSS_MAP_FILE):
    with open(WSS_MAP_FILE, "r") as f:
        try:
            wss_map = json.load(f)
        except:
            wss_map = {}
else:
    wss_map = {}

if os.path.exists(FP_FILE):
    with open(FP_FILE, "r", encoding="utf-8") as f:
        try:
            fp_map = json.load(f)
        except:
            fp_map = {}
else:
    fp_map = {}

if os.path.exists(SEED_MAP_FILE):
    with open(SEED_MAP_FILE, "r") as f:
        try:
            seed_map = json.load(f)
        except:
            seed_map = {}
else:
    seed_map = {}

if os.path.exists(CONVERSATION_MAP_FILE):
    with open(CONVERSATION_MAP_FILE, "r") as f:
        try:
            conversation_map = json.load(f)
        except:
            conversation_map = {}
else:
    conversation_map = {}

if os.path.exists(TOKENS_FILE):
    with open(TOKENS_FILE, "r", encoding="utf-8") as f:
        for line in f:
            if line.strip() and not line.startswith("#"):
                token_list.append(line.strip())
else:
    with open(TOKENS_FILE, "w", encoding="utf-8") as f:
        pass

if os.path.exists(ERROR_TOKENS_FILE):
    with open(ERROR_TOKENS_FILE, "r", encoding="utf-8") as f:
        for line in f:
            if line.strip() and not line.startswith("#"):
                error_token_list.append(line.strip())
else:
    with open(ERROR_TOKENS_FILE, "w", encoding="utf-8") as f:
        pass

if token_list:
    logger.info(f"Token list count: {len(token_list)}, Error token list count: {len(error_token_list)}")
    logger.info("-" * 60)
````

## File: utils/kv_utils.py
````python
def set_value_for_key_dict(data, target_key, new_value):
    if isinstance(data, dict):
        for key, value in data.items():
            if key == target_key:
                data[key] = new_value
            else:
                set_value_for_key_dict(value, target_key, new_value)
    elif isinstance(data, list):
        for item in data:
            set_value_for_key_dict(item, target_key, new_value)


def set_value_for_key_list(data, target_key, new_value):
    if isinstance(data, dict):
        for key, value in data.items():
            if key == target_key:
                data[key] = new_value
            else:
                set_value_for_key_list(value, target_key, new_value)
    elif isinstance(data, list):
        for i in range(len(data) - 1):
            if data[i] == target_key:
                data[i + 1] = new_value
            elif isinstance(data[i], (dict, list)):
                set_value_for_key_list(data[i], target_key, new_value)
        if data and isinstance(data[-1], (dict, list)):
            set_value_for_key_list(data[-1], target_key, new_value)
````

## File: utils/Logger.py
````python
import logging

logging.basicConfig(level=logging.INFO, format='%(asctime)s | %(levelname)s | %(message)s')


class Logger:
    @staticmethod
    def info(message):
        logging.info(str(message))

    @staticmethod
    def warning(message):
        logging.warning("\033[0;33m" + str(message) + "\033[0m")

    @staticmethod
    def error(message):
        logging.error("\033[0;31m" + "-" * 50 + '\n| ' + str(message) + "\033[0m" + "\n" + "└" + "-" * 80)

    @staticmethod
    def debug(message):
        logging.debug("\033[0;37m" + str(message) + "\033[0m")


logger = Logger()
````

## File: utils/retry.py
````python
from fastapi import HTTPException

from utils.Logger import logger
from utils.configs import retry_times


async def async_retry(func, *args, max_retries=retry_times, **kwargs):
    for attempt in range(max_retries + 1):
        try:
            result = await func(*args, **kwargs)
            return result
        except HTTPException as e:
            if attempt == max_retries:
                logger.error(f"Throw an exception {e.status_code}, {e.detail}")
                if e.status_code == 500:
                    raise HTTPException(status_code=500, detail="Server error")
                raise HTTPException(status_code=e.status_code, detail=e.detail)
            logger.info(f"Retry {attempt + 1} status code {e.status_code}, {e.detail}. Retrying...")


def retry(func, *args, max_retries=retry_times, **kwargs):
    for attempt in range(max_retries + 1):
        try:
            result = func(*args, **kwargs)
            return result
        except HTTPException as e:
            if attempt == max_retries:
                logger.error(f"Throw an exception {e.status_code}, {e.detail}")
                if e.status_code == 500:
                    raise HTTPException(status_code=500, detail="Server error")
                raise HTTPException(status_code=e.status_code, detail=e.detail)
            logger.error(f"Retry {attempt + 1} status code {e.status_code}, {e.detail}. Retrying...")
````

## File: version.txt
````
1.8.8-beta2
````
