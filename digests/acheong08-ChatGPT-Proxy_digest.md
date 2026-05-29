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
.github/workflows/ci.yaml
.gitignore
.pre-commit-config.yaml
docker-compose.yml
Dockerfile
launch.sh
LICENSE
Pipfile
proxy.py
README.md
requirements.txt
scripts/install-chrome.sh
scripts/run-with-docker.sh
</directory_structure>

<files>
This section contains the contents of the repository's files.

<file path=".github/workflows/ci.yaml">
name: ChatgGPT-Proxy CI

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]
env:
  REGISTRY: ghcr.io
  IMAGE_NAME: ${{ github.repository }}

jobs:
  build-and-push-image:
    runs-on: ubuntu-latest
    permissions:
      contents: read
      packages: write

    steps:
      - name: Checkout repository
        uses: actions/checkout@v3

      - name: Log in to the Container registry
        uses: docker/login-action@f054a8b539a109f9f41c372932f1ae047eff08c9
        with:
          registry: ${{ env.REGISTRY }}
          username: ${{ github.actor }}
          password: ${{ secrets.GITHUB_TOKEN }}

      - name: Extract metadata (tags, labels) for Docker
        id: meta
        uses: docker/metadata-action@98669ae865ea3cffbcbaa878cf57c20bbf1c6c38
        with:
          images: ${{ env.REGISTRY }}/${{ env.IMAGE_NAME }}

      - name: Build and push Docker image
        uses: docker/build-push-action@ad44023a93711e3deb337508980b4b5e9bcdc5dc
        with:
          context: .
          push: true
          tags: ${{ steps.meta.outputs.tags }}
          labels: ${{ steps.meta.outputs.labels }}
</file>

<file path=".pre-commit-config.yaml">
repos:
  - repo: https://github.com/asottile/reorder_python_imports
    rev: v3.9.0
    hooks:
      - id: reorder-python-imports
        args: [--py37-plus]
  - repo: https://github.com/asottile/add-trailing-comma
    rev: v2.3.0
    hooks:
      - id: add-trailing-comma
        args: [--py36-plus]
  - repo: https://github.com/asottile/pyupgrade
    rev: v3.3.1
    hooks:
      - id: pyupgrade
        args: [--py37-plus]

  - repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v4.4.0
    hooks:
      - id: trailing-whitespace
      - id: end-of-file-fixer
      - id: check-yaml
      - id: debug-statements
      - id: double-quote-string-fixer
      - id: name-tests-test
      - id: requirements-txt-fixer
  - repo: https://github.com/psf/black
    rev: 22.10.0
    hooks:
      - id: black
</file>

<file path="docker-compose.yml">
services:
  app:
    image: devmitrandir/chatgpt-proxy:latest
    container_name: chatgpt-proxy
    # Optional:
    # environment:
    #   - GPT_PROXY=
    #   - GPT_HOST=0.0.0.0
    #   - GPT_PORT=5000
    ports:
      - "5000:5000"
</file>

<file path="Dockerfile">
FROM ubuntu:latest

WORKDIR /app

COPY . /app

RUN apt update \
    && apt install git curl python3 python3-pip xvfb -y \
    && cd /app \
    && python3 -m pip install pipenv \
    && pipenv update -d

# Install Chrome
RUN bash /app/scripts/install-chrome.sh

ENTRYPOINT [ "bash", "/app/launch.sh" ]
</file>

<file path="launch.sh">
#!/bin/bash

cd "$(dirname "$0")"

xvfb-run pipenv run proxy
</file>

<file path="LICENSE">
This is free and unencumbered software released into the public domain.

Anyone is free to copy, modify, publish, use, compile, sell, or
distribute this software, either in source code form or as a compiled
binary, for any purpose, commercial or non-commercial, and by any
means.

In jurisdictions that recognize copyright laws, the author or authors
of this software dedicate any and all copyright interest in the
software to the public domain. We make this dedication for the benefit
of the public at large and to the detriment of our heirs and
successors. We intend this dedication to be an overt act of
relinquishment in perpetuity of all present and future rights to this
software under copyright law.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.

For more information, please refer to <https://unlicense.org>
</file>

<file path="Pipfile">
[[source]]
url = "https://pypi.org/simple"
verify_ssl = true
name = "pypi"

[packages]
openaiauth = "*"
uvicorn = "*"
asgiref = "*"
undetected_chromedriver = "*"
flask = "*"
tls_client = "*"

[dev-packages]
yapf = "*"

[requires]
python_version = "3"

[scripts]
proxy = "python proxy.py"
</file>

<file path="proxy.py">
"""
Fetches cookies from chat.openai.com and returns them (Flask)
"""
import json
import os
import tls_client
import uvicorn

from asgiref.wsgi import WsgiToAsgi
from flask import Flask
from flask import jsonify
from flask import request
from OpenAIAuth.Cloudflare import Cloudflare

GPT_PROXY = os.getenv('GPT_PROXY')
GPT_HOST = os.getenv('GPT_HOST', '0.0.0.0')
GPT_PORT = int(os.getenv('GPT_PORT', 5000))

app = Flask(__name__)

session = tls_client.Session(client_identifier="chrome_108", )
if GPT_PROXY:
    session.proxies.update(http=GPT_PROXY, https=GPT_PROXY)

authentication = {}

context = {"blocked": False}

# Get cloudflare cookies
(
    authentication["cf_clearance"],
    authentication["user_agent"],
) = Cloudflare(proxy=GPT_PROXY).get_cf_cookies()


@app.route("/<path:subpath>", methods=["POST", "GET"])
def conversation(subpath: str):
    if request.headers.get("Authorization") is None:
        return jsonify({"error": "Missing Authorization header"})
    try:
        if context.get("blocked"):
            return jsonify({"error": "Blocking operation in progress"})
        # Get cookies from request
        cookies = {
            "cf_clearance":
            authentication["cf_clearance"],
            "__Secure-next-auth.session-token":
            request.cookies.get("__Secure-next-auth.session-token"),
        }

        # Set user agent
        headers = {
            "Accept": "text/event-stream",
            "Authorization": request.headers.get("Authorization"),
            "User-Agent": authentication["user_agent"],
            "Content-Type": "application/json",
            "X-Openai-Assistant-App-Id": "",
            "Connection": "close",
            "Accept-Language": "en-US,en;q=0.9",
            "Referer": "https://chat.openai.com/" + "chat",
        }

        # Send request to OpenAI
        if request.method == "POST":
            response = session.post(
                url="https://chat.openai.com/" + subpath,
                headers=headers,
                cookies=cookies,
                data=json.dumps(request.get_json()),
                timeout_seconds=360,
            )
        elif request.method == "GET":
            response = session.get(
                url="https://chat.openai.com/" + subpath,
                headers=headers,
                cookies=cookies,
                timeout_seconds=360,
            )

        # Check status code
        if response.status_code == 403:
            # Get cf_clearance again
            context["blocked"] = True
            (
                authentication["cf_clearance"],
                authentication["user_agent"],
            ) = Cloudflare(proxy=GPT_PROXY).get_cf_cookies()
            context["blocked"] = False
            # return error
            return jsonify({
                "error":
                "Cloudflare token expired. Please wait a few minutes while I refresh"
            })
        # Return response
        return response.text
    except Exception as exc:
        return jsonify({"error": str(exc)})


if __name__ == "__main__":
    uvicorn.run(
        WsgiToAsgi(app),
        host=GPT_HOST,
        port=GPT_PORT,
        server_header=False)  # start a high-performance server with Uvicorn
</file>

<file path="requirements.txt">
flask
tls_client
OpenAIAuth
undetected_chromedriver
uvicorn
asgiref
</file>

<file path="scripts/install-chrome.sh">
apt install wget -y

architecture=$(dpkg --print-architecture)
chrome_deb="google-chrome-stable_current_${architecture}.deb"
wget https://dl.google.com/linux/direct/$chrome_deb
apt install -y ./$chrome_deb
apt-get install -y -f
rm $chrome_deb
</file>

<file path="scripts/run-with-docker.sh">
docker run -d --name=chatgpt-proxy --network host devmitrandir/chatgpt-proxy:latest \
&& echo -e "chatgpt-proxy container started" \
&& echo -e "View logs:\n  docker logs --tail 1000 -f chatgpt-proxy" \
&& echo -e "Stop container:\n  docker stop chatgpt-proxy"
</file>

<file path=".gitignore">
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
pip-wheel-metadata/
share/python-wheels/
*.egg-info/
.installed.cfg
*.egg
MANIFEST

# PyInstaller
#  Usually these files are written by a python script from a template
#  before PyInstaller builds the exe, so as to inject date/other infos into it.
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

# Django stuff:
*.log
local_settings.py
db.sqlite3
db.sqlite3-journal

# Flask stuff:
instance/
.webassets-cache

# Scrapy stuff:
.scrapy

# Sphinx documentation
docs/_build/

# PyBuilder
target/

# Jupyter Notebook
.ipynb_checkpoints

# IPython
profile_default/
ipython_config.py

# pyenv
.python-version

# pipenv
Pipfile.lock

# PEP 582; used by e.g. github.com/David-OConnor/pyflow
__pypackages__/

# Celery stuff
celerybeat-schedule
celerybeat.pid

# SageMath parsed files
*.sage.py

# Environments
.env
.venv
env/
venv/
ENV/
env.bak/
venv.bak/

# Spyder project settings
.spyderproject
.spyproject

# Rope project settings
.ropeproject

# mkdocs documentation
/site

# mypy
.mypy_cache/
.dmypy.json
dmypy.json

# Pyre type checker
.pyre/

# for this project
env
</file>

<file path="README.md">
# ChatGPT-Proxy
Forward requests and inject CloudFlare cookies

# > Unmaintained 

## Installation

### One-click scripts 

- With Docker: `curl https://raw.githubusercontent.com/acheong08/ChatGPT-Proxy/main/scripts/run-with-docker.sh | sh`


### Simple steps

1. Clone the repository
2. Check if Pipenv is installed. If not, run `pip install pipenv -U`.
3. Then, run `pipenv update -d` in this directory, to automatically install the requirements of the proxy.
4. Run `pipenv run proxy` in the base directory, and enjoy it! Uvicorn will provide a high-performance HTTP server for the API service.


### Options

These options can be configured by setting environment variables using `-e KEY="VALUE"` in the `docker run` command.

| Env | Default | Example | Description |
| - | - | - | - |
| `GPT_PROXY` | - | `socks5://127.0.0.1:1080` | The proxy of your server. |
| `GPT_HOST` | `0.0.0.0` | `127.0.0.1` | The hostname of your server. |
| `GPT_PORT` | `5000` | `8080` | The port of your server. |
</file>

</files>
