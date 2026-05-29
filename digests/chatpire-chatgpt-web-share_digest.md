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
.github/ISSUE_TEMPLATE/bug-report---bug-报告.md
.github/ISSUE_TEMPLATE/feature-request---新功能建议.md
.github/ISSUE_TEMPLATE/seek-help---寻求帮助-无法使用-报错等-.md
.github/workflows/codeql.yml
.github/workflows/docker-image.yml
.gitignore
.gitmodules
backend/.gitignore
backend/alembic.ini
backend/alembic/env.py
backend/alembic/README
backend/alembic/script.py.mako
backend/api/__init__.py
backend/api/config/__init__.py
backend/api/config/config.yaml.template
backend/api/database.py
backend/api/enums.py
backend/api/exceptions.py
backend/api/globals.py
backend/api/middlewares/__init__.py
backend/api/middlewares/asgi_logger/__init__.py
backend/api/middlewares/asgi_logger/middleware.py
backend/api/middlewares/asgi_logger/utils.py
backend/api/middlewares/request_statistics.py
backend/api/models.py
backend/api/response.py
backend/api/revchatgpt.py
backend/api/routers/__init__.py
backend/api/routers/chat.py
backend/api/routers/status.py
backend/api/routers/system.py
backend/api/routers/users.py
backend/api/schema.py
backend/api/users.py
backend/logging_config.yaml
backend/main.py
backend/pyproject.toml
backend/requirements.txt
backend/utils/__init__.py
backend/utils/common.py
backend/utils/create_user.py
backend/utils/data_types.py
backend/utils/logger.py
backend/utils/proxy.py
backend/utils/store_statistics.py
backend/utils/sync_conversations.py
Caddyfile
docker-compose.yaml
Dockerfile
docs/donate.png
docs/screenshot_admin.jpeg
docs/screenshot.en.jpeg
docs/screenshot.jpeg
frontend/.env
frontend/.eslintignore
frontend/.eslintrc.cjs
frontend/.gitignore
frontend/.prettierrc.cjs
frontend/.vscode/extensions.json
frontend/.vscode/launch.json
frontend/.vscode/settings.json
frontend/components.d.ts
frontend/config/vite.config.base.ts
frontend/config/vite.config.dev.ts
frontend/config/vite.config.prod.ts
frontend/index.html
frontend/package.json
frontend/public/chatgpt-icon-black.svg
frontend/public/chatgpt-icon.svg
frontend/README.md
frontend/src/api/chat.ts
frontend/src/api/interceptor.ts
frontend/src/api/status.ts
frontend/src/api/system.ts
frontend/src/api/url.ts
frontend/src/api/user.ts
frontend/src/App.vue
frontend/src/components/PageHeader.vue
frontend/src/components/PreferenceForm.vue
frontend/src/components/UserProfileCard.vue
frontend/src/i18n.ts
frontend/src/locales/en-US.json
frontend/src/locales/zh-CN.json
frontend/src/main.ts
frontend/src/router/guard/index.ts
frontend/src/router/guard/permission.ts
frontend/src/router/guard/userLoginInfo.ts
frontend/src/router/index.ts
frontend/src/router/typings.d.ts
frontend/src/store/index.ts
frontend/src/store/modules/app.ts
frontend/src/store/modules/conversation.ts
frontend/src/store/modules/user.ts
frontend/src/store/types.ts
frontend/src/style.css
frontend/src/types/custom.ts
frontend/src/types/echarts.ts
frontend/src/types/openapi.json
frontend/src/types/openapi.ts
frontend/src/types/schema.ts
frontend/src/utils/auth.ts
frontend/src/utils/conversation.ts
frontend/src/utils/cookies.ts
frontend/src/utils/loading.ts
frontend/src/utils/markdown.ts
frontend/src/utils/renders.ts
frontend/src/utils/tips.ts
frontend/src/views/admin/components/charts/AskChart.vue
frontend/src/views/admin/components/charts/helpers.ts
frontend/src/views/admin/components/charts/RequestsChart.vue
frontend/src/views/admin/components/EditLimitForm.vue
frontend/src/views/admin/components/EditUserForm.vue
frontend/src/views/admin/components/StatisticsCard.vue
frontend/src/views/admin/components/SystemInfoCard.vue
frontend/src/views/admin/components/UserSelector.vue
frontend/src/views/admin/conversation_manager.vue
frontend/src/views/admin/index.vue
frontend/src/views/admin/log_viewer.vue
frontend/src/views/admin/system_manager.vue
frontend/src/views/admin/user_manager.vue
frontend/src/views/conversation/components/HistoryContent.vue
frontend/src/views/conversation/components/InputRegion.vue
frontend/src/views/conversation/components/LeftBar.vue
frontend/src/views/conversation/components/MessageRow.vue
frontend/src/views/conversation/components/StatusCard.vue
frontend/src/views/conversation/history-viewer.vue
frontend/src/views/conversation/index.vue
frontend/src/views/conversation/utils/export.ts
frontend/src/views/error/403.vue
frontend/src/views/error/404.vue
frontend/src/views/home.vue
frontend/src/views/login/index.vue
frontend/src/views/redirect/index.vue
frontend/src/vite-env.d.ts
frontend/tsconfig.json
frontend/tsconfig.node.json
frontend/updateapi.sh
LICENSE
README.en.md
README.md
startup.sh
```

# Files

## File: .github/ISSUE_TEMPLATE/bug-report---bug-报告.md
```markdown
---
name: Bug report / bug 报告
about: report a code bug / 除非你肯定是代码的问题而不是你的问题
title: ''
labels: bug
assignees: ''

---

**Version**
v0.x.x

**What's your deploying method?**
- [ ] Docker
- [ ] Caddy
- [ ] Other

**Describe the bug**
A clear and concise description of what the bug is.

**To Reproduce**
Steps to reproduce the behavior:
1. Go to '...'
2. Click on '....'
3. Scroll down to '....'
4. See error

**Expected behavior**
A clear and concise description of what you expected to happen.

**Your config.yaml or other configurations**
Provide your configurations. You may hide your secrets or access tokens.

**Screenshots or running logs**
If applicable, add screenshots or logs to help explain your problem.

**Additional context**
Add any other context about the problem here.
```

## File: .github/ISSUE_TEMPLATE/feature-request---新功能建议.md
```markdown
---
name: Feature request / 新功能建议
about: Suggest an idea for this project / 提建议前请先看有没有重复的 issue
title: ''
labels: enhancement
assignees: ''

---

**Is your feature request related to a problem? Please describe.**
A clear and concise description of what the problem is. Ex. I'm always frustrated when [...]

**Describe the solution you'd like**
A clear and concise description of what you want to happen.

**Describe alternatives you've considered**
A clear and concise description of any alternative solutions or features you've considered.

**Additional context**
Add any other context or screenshots about the feature request here.
```

## File: .github/ISSUE_TEMPLATE/seek-help---寻求帮助-无法使用-报错等-.md
```markdown
---
name: Seek help / 寻求帮助（无法使用、报错等）
about: Ask for help when you meet trouble / 请按照模版填写，提供尽可能多的信息，否则不予回复
title: ''
labels: ''
assignees: ''

---

**Version**
v0.x.x

**What's your deploying environment?**
- [ ] Docker
- [ ] Caddy
- [ ] Other

**Describe the problem**
A clear and concise description of your problem.

**Expected behavior**
A clear and concise description of what you expected to happen.

**Your config.yaml or other configurations**
Provide your configurations. You may hide your secrets or access tokens.

**Screenshots or running logs**
If applicable, add screenshots or logs to help explain your problem.

**Additional context**
Add any other context about the problem here.
```

## File: .github/workflows/codeql.yml
```yaml
# For most projects, this workflow file will not need changing; you simply need
# to commit it to your repository.
#
# You may wish to alter this file to override the set of languages analyzed,
# or to provide custom queries or build logic.
#
# ******** NOTE ********
# We have attempted to detect the languages in your repository. Please check
# the `language` matrix defined below to confirm you have the correct set of
# supported CodeQL languages.
#
name: "CodeQL"

on:
  push:
    branches: [ "main" ]
  pull_request:
    # The branches below must be a subset of the branches above
    branches: [ "main" ]

jobs:
  analyze:
    name: Analyze
    runs-on: ubuntu-latest
    permissions:
      actions: read
      contents: read
      security-events: write

    strategy:
      fail-fast: false
      matrix:
        language: [ 'python' ]
        # CodeQL supports [ 'cpp', 'csharp', 'go', 'java', 'javascript', 'python', 'ruby' ]
        # Use only 'java' to analyze code written in Java, Kotlin or both
        # Use only 'javascript' to analyze code written in JavaScript, TypeScript or both
        # Learn more about CodeQL language support at https://aka.ms/codeql-docs/language-support

    steps:
    - name: Checkout repository
      uses: actions/checkout@v3

    # Initializes the CodeQL tools for scanning.
    - name: Initialize CodeQL
      uses: github/codeql-action/init@v2
      with:
        languages: ${{ matrix.language }}
        # If you wish to specify custom queries, you can do so here or in a config file.
        # By default, queries listed here will override any specified in a config file.
        # Prefix the list here with "+" to use these queries and those in the config file.

        # Details on CodeQL's query packs refer to : https://docs.github.com/en/code-security/code-scanning/automatically-scanning-your-code-for-vulnerabilities-and-errors/configuring-code-scanning#using-queries-in-ql-packs
        # queries: security-extended,security-and-quality


    # Autobuild attempts to build any compiled languages  (C/C++, C#, Go, or Java).
    # If this step fails, then you should remove it and run the build manually (see below)
    - name: Autobuild
      uses: github/codeql-action/autobuild@v2

    # ℹ️ Command-line programs to run using the OS shell.
    # 📚 See https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idstepsrun

    #   If the Autobuild fails above, remove it and uncomment the following three lines.
    #   modify them (or add more) to build your code if your project, please refer to the EXAMPLE below for guidance.

    # - run: |
    #     echo "Run, Build Application using script"
    #     ./location_of_script_within_repo/buildscript.sh

    - name: Perform CodeQL Analysis
      uses: github/codeql-action/analyze@v2
      with:
        category: "/language:${{matrix.language}}"
```

## File: .github/workflows/docker-image.yml
```yaml
name: Build and publish Docker image

on:
  push:
    tags: ["v*"]
    branches: ["test-workflow"]
  workflow_dispatch:
    inputs:
      version:
        description: "Version tag for the Docker image (semver)"
        required: true

env:
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

      - name: Clone ChatGPT-Proxy-V4
        run: git clone https://github.com/moeakwak/ChatGPT-Proxy-V4.git

      # - name: Setup Go
      #   uses: actions/setup-go@v4
      #   with:
      #     go-version: "1.20"
      #     cache-dependency-path: ChatGPT-Proxy-V4/go.sum

      # - name: build ChatGPT-Proxy-V4
      #   run: |
      #     cd ChatGPT-Proxy-V4
      #     go build

      - name: Install Node.js
        uses: actions/setup-node@v3
        with:
          node-version: 16

      - uses: pnpm/action-setup@v2
        name: Install pnpm
        id: pnpm-install
        with:
          version: 7
          run_install: false

      - name: Get pnpm store directory
        id: pnpm-cache
        shell: bash
        run: |
          echo "STORE_PATH=$(pnpm store path)" >> $GITHUB_OUTPUT

      - uses: actions/cache@v3
        name: Setup pnpm cache
        with:
          path: ${{ steps.pnpm-cache.outputs.STORE_PATH }}
          key: ${{ runner.os }}-pnpm-store-${{ hashFiles('**/pnpm-lock.yaml') }}
          restore-keys: |
            ${{ runner.os }}-pnpm-store-

      - name: Install dependencies in frontend
        run: |
          cd frontend
          pnpm install

      - name: build frontend
        run: |
          cd frontend
          pnpm build

      # - name: Move ChatGPT-Proxy-V4 to backend
      #   run: |
      #     mv ChatGPT-Proxy-V4/ChatGPT-Proxy-V4 backend/

      - name: Docker meta
        id: meta
        uses: docker/metadata-action@v4
        with:
          images: |
            ghcr.io/${{ env.IMAGE_NAME }}
            ${{ env.IMAGE_NAME }}
          tags: |
            type=semver,pattern={{version}}${{ github.event_name == 'workflow_dispatch' && format(',value={0}', github.event.inputs.version) || '' }}
            type=semver,pattern={{major}}.{{minor}}${{ github.event_name == 'workflow_dispatch' && format(',value={0}', github.event.inputs.version) || '' }}

      # - name: Create Sentry release
      #   uses: getsentry/action-release@v1
      #   env:
      #     SENTRY_AUTH_TOKEN: ${{ secrets.SENTRY_AUTH_TOKEN }}
      #     SENTRY_ORG: ${{ secrets.SENTRY_ORG }}
      #     SENTRY_PROJECT: ${{ secrets.SENTRY_PROJECT }}
      #     # SENTRY_URL: https://sentry.io/
      #   with:
      #     environment: production
      #     sourcemaps: frontend/dist/assets
      #     version: ${{ github.ref }}

      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v2

      - name: Log in to the Container registry
        uses: docker/login-action@v2
        with:
          registry: ghcr.io
          username: ${{ github.actor }}
          password: ${{ secrets.GITHUB_TOKEN }}
      
      - name: Log in to Dockerhub
        uses: docker/login-action@v2
        with:
          username: ${{ github.actor }}
          password: ${{ secrets.DOCKERHUB_TOKEN }}

      - name: Build and push Docker image
        uses: docker/build-push-action@v4
        with:
          context: .
          platforms: linux/amd64
          push: true
          tags: ${{ steps.meta.outputs.tags }}
          labels: ${{ steps.meta.outputs.labels }}
          cache-from: type=gha
          cache-to: type=gha,mode=max
```

## File: .gitignore
```
docker/
```

## File: .gitmodules
```

```

## File: backend/.gitignore
```
__pycache__
.idea
config.yaml
*.db
.vscode
files
logs
*.log
ChatGPT-Proxy-V4
*.json
```

## File: backend/alembic.ini
```ini
# A generic, single database configuration.

[alembic]
# path to migration scripts
script_location = alembic

# template used to generate migration file names; The default value is %%(rev)s_%%(slug)s
# Uncomment the line below if you want the files to be prepended with date and time
# file_template = %%(year)d_%%(month).2d_%%(day).2d_%%(hour).2d%%(minute).2d-%%(rev)s_%%(slug)s

# sys.path path, will be prepended to sys.path if present.
# defaults to the current working directory.
prepend_sys_path = .

# timezone to use when rendering the date within the migration file
# as well as the filename.
# If specified, requires the python-dateutil library that can be
# installed by adding `alembic[tz]` to the pip requirements
# string value is passed to dateutil.tz.gettz()
# leave blank for localtime
# timezone =

# max length of characters to apply to the
# "slug" field
# truncate_slug_length = 40

# set to 'true' to run the environment during
# the 'revision' command, regardless of autogenerate
# revision_environment = false

# set to 'true' to allow .pyc and .pyo files without
# a source .py file to be detected as revisions in the
# versions/ directory
# sourceless = false

# version location specification; This defaults
# to alembic/versions.  When using multiple version
# directories, initial revisions must be specified with --version-path.
# The path separator used here should be the separator specified by "version_path_separator" below.
# version_locations = %(here)s/bar:%(here)s/bat:alembic/versions

# version path separator; As mentioned above, this is the character used to split
# version_locations. The default within new alembic.ini files is "os", which uses os.pathsep.
# If this key is omitted entirely, it falls back to the legacy behavior of splitting on spaces and/or commas.
# Valid values for version_path_separator are:
#
# version_path_separator = :
# version_path_separator = ;
# version_path_separator = space
version_path_separator = os  # Use os.pathsep. Default configuration used for new projects.

# set to 'true' to search source files recursively
# in each "version_locations" directory
# new in Alembic version 1.10
# recursive_version_locations = false

# the output encoding used when revision files
# are written from script.py.mako
# output_encoding = utf-8

sqlalchemy.url = sqlite+aiosqlite:///database.db


[post_write_hooks]
# post_write_hooks defines scripts or Python functions that are run
# on newly generated revision scripts.  See the documentation for further
# detail and examples

# format using "black" - use the console_scripts runner, against the "black" entrypoint
# hooks = black
# black.type = console_scripts
# black.entrypoint = black
# black.options = -l 79 REVISION_SCRIPT_FILENAME

# Logging configuration
```

## File: backend/alembic/env.py
```python
import asyncio
from logging.config import fileConfig

from sqlalchemy import pool
from sqlalchemy.engine import Connection
from sqlalchemy.ext.asyncio import async_engine_from_config

from alembic import context

from api import models

# this is the Alembic Config object, which provides
# access to the values within the .ini file in use.
config = context.config

# Interpret the config file for Python logging.
# This line sets up loggers basically.

# 阻止 alembic 重复配置日志
# if config.config_file_name is not None:
#     fileConfig(config.config_file_name, disable_existing_loggers=False)

# add your model's MetaData object here
# for 'autogenerate' support
# from myapp import mymodel
# target_metadata = mymodel.Base.metadata
target_metadata = models.Base.metadata


# other values from the config, defined by the needs of env.py,
# can be acquired:
# my_important_option = config.get_main_option("my_important_option")
# ... etc.


def run_migrations_offline() -> None:
    """Run migrations in 'offline' mode.

    This configures the context with just a URL
    and not an Engine, though an Engine is acceptable
    here as well.  By skipping the Engine creation
    we don't even need a DBAPI to be available.

    Calls to context.execute() here emit the given string to the
    script output.

    """
    url = config.get_main_option("sqlalchemy.url")
    context.configure(
        url=url,
        target_metadata=target_metadata,
        literal_binds=True,
        dialect_opts={"paramstyle": "named"},
    )

    with context.begin_transaction():
        context.run_migrations()


def do_run_migrations(connection: Connection) -> None:
    context.configure(connection=connection, target_metadata=target_metadata)

    with context.begin_transaction():
        context.run_migrations()


async def run_async_migrations() -> None:
    """In this scenario we need to create an Engine
    and associate a connection with the context.

    """

    connectable = async_engine_from_config(
        config.get_section(config.config_ini_section, {}),
        prefix="sqlalchemy.",
        poolclass=pool.NullPool,
    )

    async with connectable.connect() as connection:
        await connection.run_sync(do_run_migrations)

    await connectable.dispose()


def run_migrations_online() -> None:
    """Run migrations in 'online' mode."""

    connectable = config.attributes.get("connection", None)

    if connectable is None:
        asyncio.run(run_async_migrations())
    else:
        do_run_migrations(connectable)


if context.is_offline_mode():
    run_migrations_offline()
else:
    run_migrations_online()
```

## File: backend/alembic/README
```
Generic single-database configuration with an async dbapi.
```

## File: backend/alembic/script.py.mako
```
"""${message}

Revision ID: ${up_revision}
Revises: ${down_revision | comma,n}
Create Date: ${create_date}

"""
from alembic import op
import sqlalchemy as sa
${imports if imports else ""}

# revision identifiers, used by Alembic.
revision = ${repr(up_revision)}
down_revision = ${repr(down_revision)}
branch_labels = ${repr(branch_labels)}
depends_on = ${repr(depends_on)}


def upgrade() -> None:
    ${upgrades if upgrades else "pass"}


def downgrade() -> None:
    ${downgrades if downgrades else "pass"}
```

## File: backend/api/__init__.py
```python

```

## File: backend/api/config/__init__.py
```python
import os
import shutil
import yaml

try:
    from yaml import CLoader as Loader, CDumper as Dumper
except ImportError:
    from yaml import Loader, Dumper


class Config:
    def __init__(self, config_file):
        # 如果缺少配置文件，则复制模板并创建文件
        if not os.path.exists(config_file):
            if os.path.exists(config_file + ".template"):
                shutil.copyfile(config_file + ".template", config_file)
        with open(config_file, 'r') as f:
            self.config = yaml.load(f, Loader=Loader)

    def get(self, key, default=None):
        return self.config.get(key, default)

    def set(self, key, value):
        self.config[key] = value

    def save(self, config_file):
        with open(config_file, 'w') as f:
            yaml.dump(self.config, f)


config_file = os.path.join(os.path.dirname(__file__), "config.yaml")
```

## File: backend/api/config/config.yaml.template
```
print_sql: false
host: "127.0.0.1"
port: 8000
data_dir: ./data
database_url: "sqlite+aiosqlite:///data/database.db"
run_migration: false
jwt_secret: "SECRET"
jwt_lifetime_seconds: 86400
cookie_max_age: 86400
user_secret: "SECRET"

sync_conversations_on_startup: true
create_initial_admin_user: true
create_initial_user: false
initial_admin_username: admin
initial_admin_password: adminadmin
initial_user_username: user
initial_user_password: useruser

chatgpt_access_token: "chatgpt_access_token"
chatgpt_paid: false

# proxy configuration
# chatgpt_base_url: http://127.0.0.1:6062/api/
# run_reverse_proxy: true
# reverse_proxy_port: 6062
# reverse_proxy_binary_path: ChatGPT-Proxy-V4
# reverse_proxy_puid: "_puid value from cookie"

log_dir: ./logs
console_log_level: INFO

request_log_counter_time_window: 2592000  # 30 days
request_log_counter_interval: 1800  # 30 minutes
ask_log_time_window: 2592000  # 30 days
sync_conversations_regularly: yes
```

## File: backend/api/database.py
```python
import asyncio
import contextlib
from typing import AsyncGenerator

from fastapi import Depends
import sqlalchemy
from sqlalchemy.ext.asyncio import AsyncSession, create_async_engine
from sqlalchemy.orm import sessionmaker
from fastapi_users_db_sqlalchemy import SQLAlchemyUserDatabase

from alembic.config import Config
from alembic import command

import api.globals as g
config = g.config
from api.models import Base, User

from utils.logger import get_logger

logger = get_logger(__name__)

database_url = config.get("database_url")
engine = create_async_engine(database_url, echo=config.get("print_sql", False))
async_session_maker = sessionmaker(engine, class_=AsyncSession, expire_on_commit=False)
metadata = sqlalchemy.MetaData()
alembic_cfg = Config("alembic.ini")
alembic_cfg.set_main_option("sqlalchemy.url", database_url)


def run_upgrade(conn, cfg):
    cfg.attributes["connection"] = conn
    command.upgrade(cfg, "head")


def run_stamp(conn, cfg, revision):
    cfg.attributes["connection"] = conn
    command.stamp(cfg, revision)


def run_ensure_version(conn, cfg):
    cfg.attributes["connection"] = conn
    command.ensure_version(cfg)


async def create_db_and_tables():
    # 如果数据库不存在则创建数据库（数据表）；若有更新，则执行迁移
    # https://alembic.sqlalchemy.org/en/latest/autogenerate.html
    async with engine.connect() as conn:
        # 判断数据库是否存在
        def use_inspector(conn):
            inspector = sqlalchemy.inspect(conn)
            return inspector.has_table("user")

        result = await conn.run_sync(use_inspector)

        if not result:
            logger.info("database not exists, creating database...")
            await conn.run_sync(Base.metadata.create_all)
            logger.info("database created!")
            await conn.run_sync(run_stamp, alembic_cfg, "head")
            logger.info(f"stamped database to head")
            return
        else:
            await conn.run_sync(run_ensure_version, alembic_cfg)

        if config.get("run_migration", False):
            try:
                logger.info("try to migrate database...")
                await conn.run_sync(run_upgrade, alembic_cfg)
            except Exception as e:
                logger.warning("Database migration might fail, please check the database manually!")
                logger.warning(f"detail: {str(e)}")


async def get_async_session() -> AsyncGenerator[AsyncSession, None]:
    async with async_session_maker() as session:
        yield session


async def get_user_db(session: AsyncSession = Depends(get_async_session)):
    yield SQLAlchemyUserDatabase(session, User)


# 使得 get_async_session_context 和 get_user_db_context 可以使用async with语法

get_async_session_context = contextlib.asynccontextmanager(get_async_session)
get_user_db_context = contextlib.asynccontextmanager(get_user_db)
```

## File: backend/api/enums.py
```python
import enum


class ChatStatus(enum.Enum):
    asking = "asking"
    queueing = "queueing"
    idling = "idling"


class ChatModels(enum.Enum):
    gpt4 = "gpt-4"
    gpt4_mobile = "gpt-4-mobile"
    default = "text-davinci-002-render-sha"
    paid = "text-davinci-002-render-paid"
    unknown = ""
```

## File: backend/api/exceptions.py
```python
from typing import Any


class SelfDefinedException(Exception):
    def __init__(self, reason: Any = None, message: str = "") -> None:
        self.reason = reason  # 异常主要原因
        self.message = message  # 更细节的描述


class AuthorityDenyException(SelfDefinedException):
    def __init__(self, message: str = ""):
        super().__init__("errors.authorityDeny", message)


class UserNotExistException(SelfDefinedException):
    def __init__(self, message: str = ""):
        super().__init__("errors.userNotExist", message)


class InvalidParamsException(SelfDefinedException):
    def __init__(self, message: str = ""):
        super().__init__("errors.invalidParams", message)


class ResourceNotFoundException(SelfDefinedException):
    def __init__(self, message: str = ""):
        super().__init__("errors.resourceNotFound", message)


class InvalidRequestException(SelfDefinedException):
    def __init__(self, message: str = ""):
        super().__init__("errors.invalidRequest", message)


class InternalException(SelfDefinedException):
    def __init__(self, message: str = ""):
        super().__init__("errors.internal", message)
```

## File: backend/api/globals.py
```python
from api.config import Config, config_file
from utils.data_types import RequestCounter, TimeQueue

# log settings

reverse_proxy_log_file = None
reverse_proxy_process = None
server_log_filename = None

# system info

startup_time = None

# request_statistics

config = Config(config_file)
request_log_counter_time_window = config.get("request_log_counter_time_window", 30 * 24 * 60 * 60)  # 30 days
request_log_counter_interval = config.get("request_log_counter_interval", 30 * 60)  # 30 minutes
request_log_counter = RequestCounter(
    time_window=request_log_counter_time_window,
    interval=request_log_counter_interval
)
ask_log_queue = TimeQueue(config.get("ask_log_time_window", 7 * 24 * 60 * 60))  # 7 days
```

## File: backend/api/middlewares/__init__.py
```python
from .asgi_logger import AccessLoggerMiddleware
from .request_statistics import StatisticsMiddleware
```

## File: backend/api/middlewares/asgi_logger/__init__.py
```python
from .middleware import AccessLoggerMiddleware
```

## File: backend/api/middlewares/asgi_logger/middleware.py
```python
from __future__ import annotations

import http
import logging
import os
import sys
import time
from typing import TypedDict

from asgiref.typing import ASGI3Application, ASGIReceiveCallable, ASGISendCallable
from asgiref.typing import ASGISendEvent, HTTPScope

from .utils import get_client_addr, get_path_with_query_string


class AccessInfo(TypedDict, total=False):
    response: ASGISendEvent
    start_time: float
    end_time: float


class AccessLoggerMiddleware:
    DEFAULT_FORMAT = '%(client_addr)s - "%(request_line)s" %(status_code)s'

    def __init__(
        self,
        app: ASGI3Application,
        format: str | None = None,
        logger: logging.Logger | None = None,
    ) -> None:
        self.app = app
        self.format = format or self.DEFAULT_FORMAT
        if logger is None:
            self.logger = logging.getLogger("access")
            self.logger.setLevel(logging.INFO)
            handler = logging.StreamHandler(sys.stdout)
            handler.setLevel(logging.INFO)
            handler.setFormatter(logging.Formatter("%(message)s"))
            self.logger.addHandler(handler)
        else:
            self.logger = logger

    async def __call__(
        self, scope: HTTPScope, receive: ASGIReceiveCallable, send: ASGISendCallable
    ) -> None:
        if scope["type"] != "http":
            return await self.app(scope, receive, send)  # pragma: no cover

        info = AccessInfo(response={})

        async def inner_send(message: ASGISendEvent) -> None:
            if message["type"] == "http.response.start":
                info["response"] = message
            await send(message)

        try:
            info["start_time"] = time.time()
            await self.app(scope, receive, inner_send)
        except Exception as exc:
            info["response"]["status"] = 500
            raise exc
        finally:
            info["end_time"] = time.time()
            self.log(scope, info)

    def log(self, scope: HTTPScope, info: AccessInfo) -> None:
        self.logger.info(self.format, AccessLogAtoms(scope, info))


class AccessLogAtoms(dict):
    def __init__(self, scope: HTTPScope, info: AccessInfo) -> None:
        for name, value in scope["headers"]:
            self[f"{{{name.decode('latin1').lower()}}}i"] = value.decode("latin1")
        for name, value in info["response"].get("headers", []):
            self[f"{{{name.decode('latin1').lower()}}}o"] = value.decode("latin1")
        for name, value in os.environ.items():
            self[f"{{{name.lower()!r}}}e"] = value

        protocol = f"HTTP/{scope['http_version']}"

        status = info["response"]["status"]
        try:
            status_phrase = http.HTTPStatus(status).phrase
        except ValueError:
            status_phrase = "-"

        path = scope["root_path"] + scope["path"]
        full_path = get_path_with_query_string(scope)
        request_line = f"{scope['method']} {path} {protocol}"
        full_request_line = f"{scope['method']} {full_path} {protocol}"

        request_time = info["end_time"] - info["start_time"]
        client_addr = get_client_addr(scope)
        self.update(
            {
                "h": client_addr,
                "client_addr": client_addr,
                "l": "-",
                "u": "-",  # Not available on ASGI.
                "t": time.strftime("[%d/%b/%Y:%H:%M:%S %z]"),
                "r": request_line,
                "request_line": full_request_line,
                "R": full_request_line,
                "m": scope["method"],
                "U": scope["path"],
                "q": scope["query_string"].decode(),
                "H": protocol,
                "s": status,
                "status_code": f"{status} {status_phrase}",
                "st": status_phrase,
                "B": self["{Content-Length}o"],
                "b": self.get("{Content-Length}o", "-"),
                "f": self["{Referer}i"],
                "a": self["{User-Agent}i"],
                "T": int(request_time),
                "M": int(request_time * 1_000),
                "D": int(request_time * 1_000_000),
                "L": f"{request_time:.6f}",
                "p": f"<{os.getpid()}>",
            }
        )

    def __getitem__(self, key: str) -> str:
        try:
            if key.startswith("{"):
                return super().__getitem__(key.lower())
            else:
                return super().__getitem__(key)
        except KeyError:
            return "-"
```

## File: backend/api/middlewares/asgi_logger/utils.py
```python
from urllib.parse import quote

from asgiref.typing import HTTPScope


def get_client_addr(scope: HTTPScope):
    if scope["client"] is None:
        return "-"  # pragma: no cover
    return f"{scope['client'][0]}:{scope['client'][1]}"


def get_path_with_query_string(scope: HTTPScope) -> str:
    path_with_query_string = quote(scope.get("root_path", "") + scope["path"])
    if scope["query_string"]:  # pragma: no cover
        return f"{path_with_query_string}?{scope['query_string'].decode('ascii')}"
    return path_with_query_string
```

## File: backend/api/middlewares/request_statistics.py
```python
import time
from asgiref.typing import ASGI3Application, HTTPScope, ASGIReceiveCallable, ASGISendCallable
import api.globals as g

from utils.logger import get_logger

logger = get_logger(__name__)


class StatisticsMiddleware:
    """
    Middleware for request_statistics.
    filter_paths: List of paths keywords to filter.
    """

    def __init__(
            self,
            app: ASGI3Application
    ) -> None:
        self.app = app

    async def __call__(
            self, scope: HTTPScope, receive: ASGIReceiveCallable, send: ASGISendCallable
    ) -> None:
        if scope["type"] != "http" and scope["type"] != "websocket":
            return await self.app(scope, receive, send)

        start_time = time.time()
        try:
            await self.app(scope, receive, send)
        except Exception as exc:
            raise exc
        finally:
            end_time = time.time()

            user = None
            user_id = None
            if "auth_user" in scope:
                user = scope["auth_user"]
                user_id = user.id

            g.request_log_counter.count(user_id)
            # logger.debug(g.request_log_counter)
```

## File: backend/api/models.py
```python
from typing import List, Optional

from fastapi_users_db_sqlalchemy import Integer, GUID, UUID_ID
from sqlalchemy import String, DateTime, Enum, Boolean, Float, ForeignKey
from sqlalchemy.orm import relationship
from sqlalchemy.orm import DeclarativeBase
from sqlalchemy.orm import Mapped, mapped_column

from api.enums import ChatStatus, ChatModels


# declarative base class
class Base(DeclarativeBase):
    pass


class User(Base):
    """
    用户表
    """

    __tablename__ = "user"

    id: Mapped[int] = mapped_column(Integer, primary_key=True, comment="用户id")
    username: Mapped[str] = mapped_column(String(32), unique=True, index=True, comment="用户名")
    nickname: Mapped[str] = mapped_column(String(64), comment="昵称")
    email: Mapped[str]
    active_time: Mapped[Optional[DateTime]] = mapped_column(DateTime, default=None, comment="最后活跃时间")

    chat_status: Mapped[ChatStatus] = mapped_column(Enum(ChatStatus), default=ChatStatus.idling, comment="对话状态")
    can_use_paid: Mapped[bool] = mapped_column(Boolean, default=False, comment="是否可以使用paid模型")
    can_use_gpt4: Mapped[bool] = mapped_column(Boolean, default=False, comment="是否可以使用gpt4模型")
    max_conv_count: Mapped[int] = mapped_column(Integer, default=-1, comment="最大对话数量")
    available_ask_count: Mapped[int] = mapped_column(Integer, default=-1, comment="可用的对话次数")
    available_gpt4_ask_count: Mapped[int] = mapped_column(Integer, default=-1, comment="可用的gpt4对话次数")

    is_superuser: Mapped[bool] = mapped_column(Boolean, default=False)
    is_active: Mapped[bool] = mapped_column(Boolean, default=True)
    is_verified: Mapped[bool] = mapped_column(Boolean, default=False)

    hashed_password: Mapped[str] = mapped_column(String(1024))
    conversations: Mapped[List["Conversation"]] = relationship("Conversation", back_populates="user")


class Conversation(Base):
    """
    ChatGPT 非官方 API 所使用的对话
    只记录对话和用户之间的对应关系，不存储内容
    """

    __tablename__ = "conversation"

    id: Mapped[int] = mapped_column(Integer, primary_key=True)
    conversation_id: Mapped[str] = mapped_column(String(36), index=True, unique=True)
    title: Mapped[Optional[str]] = mapped_column(comment="对话标题")
    user_id: Mapped[Optional[int]] = mapped_column(ForeignKey("user.id"), comment="发起用户id")
    user: Mapped["User"] = relationship(back_populates="conversations")
    is_valid: Mapped[bool] = mapped_column(Boolean, default=True, comment="是否有效")
    model_name: Mapped[Optional[Enum["ChatModels"]]] = mapped_column(
        Enum(ChatModels, values_callable=lambda obj: [e.value for e in obj] if obj else None), default=None, comment="使用的模型")
    create_time: Mapped[Optional[DateTime]] = mapped_column(DateTime, default=None, comment="创建时间")
    active_time: Mapped[Optional[DateTime]] = mapped_column(DateTime, default=None, comment="最后活跃时间")
```

## File: backend/api/response.py
```python
import json
import typing
from typing import Optional, Any, Generic, TypeVar, Dict

from fastapi import Response
from fastapi.encoders import jsonable_encoder
from fastapi.exceptions import RequestValidationError
from fastapi_users.router import ErrorCode
from starlette.background import BackgroundTask
from starlette.exceptions import HTTPException as StarletteHTTPException
from starlette.responses import JSONResponse

from pydantic.generics import GenericModel

from api.exceptions import SelfDefinedException
from revChatGPT.typings import Error as revChatGPTError

T = TypeVar('T')


class ResponseWrapper(GenericModel, Generic[T]):
    """
    使用自定义的返回格式：
    - 统一状态码为 200
    - 统一返回格式为 {"code", "message", "result"}
    - code 为 200 表示成功，其余表示失败。
        -1 表示一般失败
        401 表示登陆超时，需要重新登陆
        对于有状态码的错误，使用该状态码
    """

    code: int = 0
    message: str = ""
    result: Optional[T | Any] = None

    def to_dict(self):
        return jsonable_encoder(self)

    def to_json(self):
        return json.dumps(self.to_dict(), ensure_ascii=False)


class CustomJSONResponse(Response):
    media_type = "application/json"

    def __init__(
            self,
            content: Any,
            status_code: int = 200,
            headers: Optional[Dict[str, str]] = None,
            media_type: Optional[str] = None,
            background: Optional[BackgroundTask] = None,
    ) -> None:
        super().__init__(content, status_code, headers, media_type, background)

    def render(self, content: typing.Any) -> bytes:
        if not isinstance(content, ResponseWrapper):
            content = ResponseWrapper(code=self.status_code, message=get_http_message(self.status_code), result=content)
        return content.to_json().encode("utf-8")


class PrettyJSONResponse(Response):
    media_type = "application/json"

    def render(self, content: typing.Any) -> bytes:
        return json.dumps(
            jsonable_encoder(content),
            ensure_ascii=False,
            allow_nan=False,
            indent=4,
            separators=(", ", ": "),
        ).encode("utf-8")


def response(code: int = 200, message: str = "", result: Optional[Any] = None) -> CustomJSONResponse:
    return CustomJSONResponse(
        content=ResponseWrapper(code=code, message=message, result=result),
        status_code=200
    )


def get_http_message(status_code: int) -> str:
    return {
        200: "tips.requestSuccess",
        201: "tips.requestSuccess",
        204: "tips.requestSuccess",
        400: "errors.badCredentials",
        401: "errors.userNotLogin",
        # 404: "资源不存在",
        # 502: "上游请求失败",
        -1: "失败",
    }.get(status_code, "")


def handle_exception_response(e: Exception) -> CustomJSONResponse:
    if isinstance(e, RequestValidationError):
        return response(-1, f"errors.requestValidationError", e.errors())
    elif isinstance(e, SelfDefinedException):
        return response(-1, e.reason, e.message)
    elif isinstance(e, StarletteHTTPException):
        if e.detail == ErrorCode.REGISTER_USER_ALREADY_EXISTS:
            message="errors.userAlreadyExists"
        elif e.detail == ErrorCode.LOGIN_BAD_CREDENTIALS:
            message="errors.badCredentials"
        else:
            message = get_http_message(e.status_code)
        return response(e.status_code or -1, message or f"{e.status_code} {e.detail}")
    elif isinstance(e, revChatGPTError):
        return response(502, "errors.chatgptResponseError", f"{e.source} {e.code}: {e.message}")
    return response(-1, str(e))
```

## File: backend/api/revchatgpt.py
```python
import api.globals as g
from fastapi.encoders import jsonable_encoder
from revChatGPT.V1 import AsyncChatbot
import asyncio
from api.enums import ChatModels
from utils.common import get_conversation_model


class ChatGPTManager:
    def __init__(self):
        self.chatbot = AsyncChatbot({
            "access_token": g.config.get("chatgpt_access_token"),
            "paid": g.config.get("chatgpt_paid"),
            "model": "text-davinci-002-render-sha", # default model
        }, base_url=g.config.get("chatgpt_base_url", None))
        self.semaphore = asyncio.Semaphore(1)

    def is_busy(self):
        return self.semaphore.locked()

    async def get_conversations(self):
        conversations = await self.chatbot.get_conversations(limit=80)
        return conversations

    async def get_conversation_messages(self, conversation_id: str):
        # TODO: 使用 redis 缓存
        messages = await self.chatbot.get_msg_history(conversation_id)
        messages = jsonable_encoder(messages)
        model_name = get_conversation_model(messages)
        messages["model_name"] = model_name or ChatModels.unknown.value
        return messages

    async def clear_conversations(self):
        await self.chatbot.clear_conversations()

    def ask(self, message, conversation_id: str = None, parent_id: str = None,
            timeout=360, model_name: ChatModels = None):
        model = None
        if model_name is not None and model_name != ChatModels.unknown:
            model = model_name.value
        return self.chatbot.ask(message, conversation_id=conversation_id, parent_id=parent_id, model=model,
                                timeout=timeout)

    async def delete_conversation(self, conversation_id: str):
        await self.chatbot.delete_conversation(conversation_id)

    async def set_conversation_title(self, conversation_id: str, title: str):
        """Hack change_title to set title in utf-8"""
        await self.chatbot.change_title(conversation_id, title)

    async def generate_conversation_title(self, conversation_id: str, message_id: str):
        """Hack gen_title to get title"""
        await self.chatbot.gen_title(conversation_id, message_id)

    def reset_chat(self):
        self.chatbot.reset_chat()


chatgpt_manager = ChatGPTManager()
```

## File: backend/api/routers/__init__.py
```python

```

## File: backend/api/routers/chat.py
```python
import time
from datetime import datetime
from typing import List

import httpx
import requests
from fastapi import APIRouter, Depends, WebSocket
from websockets.exceptions import ConnectionClosed
from fastapi.encoders import jsonable_encoder
from httpx import HTTPError
from sqlalchemy import select, and_, delete, func
import api.revchatgpt
import api.globals as g

from api.database import get_async_session_context
from api.enums import ChatStatus, ChatModels
from api.exceptions import InvalidParamsException, AuthorityDenyException, InternalException
from api.models import User, Conversation
from api.schema import ConversationSchema
from api.users import current_active_user, websocket_auth, current_super_user
from revChatGPT.typings import Error as revChatGPTError
from api.response import response
from utils.logger import get_logger

config = g.config
logger = get_logger(__name__)
router = APIRouter()


async def get_conversation_by_id(conversation_id: str, user: User = Depends(current_active_user)):
    async with get_async_session_context() as session:
        r = await session.execute(select(Conversation).where(Conversation.conversation_id == conversation_id))
        conversation = r.scalars().one_or_none()
        if conversation is None:
            raise InvalidParamsException("errors.conversationNotFound")
        if not user.is_superuser and conversation.user_id != user.id:
            raise AuthorityDenyException
        return conversation


@router.get("/conv", tags=["conversation"], response_model=List[ConversationSchema])
async def get_all_conversations(user: User = Depends(current_active_user), fetch_all: bool = False):
    """
    返回自己的有效会话
    对于管理员，返回所有对话，并可以指定是否只返回有效会话
    """
    if fetch_all and not user.is_superuser:
        raise AuthorityDenyException()

    stat = and_(Conversation.user_id == user.id, Conversation.is_valid)
    if fetch_all:
        stat = None
    async with get_async_session_context() as session:
        if stat is not None:
            r = await session.execute(select(Conversation).where(stat))
        else:
            r = await session.execute(select(Conversation))
        results = r.scalars().all()
        results = jsonable_encoder(results)
        return results


@router.get("/conv/{conversation_id}", tags=["conversation"])
async def get_conversation_history(conversation: Conversation = Depends(get_conversation_by_id)):
    try:
        result = await api.revchatgpt.chatgpt_manager.get_conversation_messages(conversation.conversation_id)
    except httpx.HTTPStatusError as e:
        if e.response.status_code == 404:
            raise InvalidParamsException("errors.conversationNotFound")
        raise InternalException()
    # 当不知道模型名时，顺便从对话中获取
    if conversation.model_name is None:
        model_name = result.get("model_name")
        if model_name is not None and not ChatModels.unknown.value:
            async with get_async_session_context() as session:
                conversation = await session.get(Conversation, conversation.id)
                conversation.model_name = model_name
                session.add(conversation)
                await session.commit()
    return result


@router.delete("/conv/{conversation_id}", tags=["conversation"])
async def delete_conversation(conversation: Conversation = Depends(get_conversation_by_id)):
    """remove conversation from database and chatgpt server"""
    if not conversation.is_valid:
        raise InvalidParamsException("errors.conversationAlreadyDeleted")
    try:
        await api.revchatgpt.chatgpt_manager.delete_conversation(conversation.conversation_id)
    except revChatGPTError as e:
        logger.warning(f"delete conversation {conversation.conversation_id} failed: {e.code} {e.message}")
    except httpx.HTTPStatusError as e:
        if e.response.status_code != 404:
            raise e
    async with get_async_session_context() as session:
        conversation.is_valid = False
        session.add(conversation)
        await session.commit()
    return response(200)


@router.delete("/conv/{conversation_id}/vanish", tags=["conversation"])
async def vanish_conversation(conversation: Conversation = Depends(get_conversation_by_id)):
    # try:
    #     await g.chatgpt_manager.delete_conversation(conversation.conversation_id)
    # except revChatGPTError as e:
    #     logger.warning(f"delete conversation {conversation.conversation_id} failed: {e.code} {e.message}")
    # except httpx.HTTPStatusError as e:
    #     if e.response.status_code != 404:
    #         raise e
    if conversation.is_valid:
        try:
            await api.revchatgpt.chatgpt_manager.delete_conversation(conversation.conversation_id)
        except revChatGPTError as e:
            logger.warning(f"delete conversation {conversation.conversation_id} failed: {e.code} {e.message}")
        except httpx.HTTPStatusError as e:
            if e.response.status_code != 404:
                raise e
    async with get_async_session_context() as session:
        await session.execute(delete(Conversation).where(Conversation.conversation_id == conversation.conversation_id))
        await session.commit()
    return response(200)


@router.patch("/conv/{conversation_id}", tags=["conversation"], response_model=ConversationSchema)
async def change_conversation_title(title: str, conversation: Conversation = Depends(get_conversation_by_id)):
    await api.revchatgpt.chatgpt_manager.set_conversation_title(conversation.conversation_id,
                                                                title)
    async with get_async_session_context() as session:
        conversation.title = title
        session.add(conversation)
        await session.commit()
        await session.refresh(conversation)
    result = jsonable_encoder(conversation)
    return result


@router.patch("/conv/{conversation_id}/assign/{username}", tags=["conversation"])
async def assign_conversation(username: str, conversation_id: str, _user: User = Depends(current_super_user)):
    async with get_async_session_context() as session:
        user = await session.execute(select(User).where(User.username == username))
        user = user.scalars().one_or_none()
        if user is None:
            raise InvalidParamsException("errors.userNotFound")
        conversation = await session.execute(
            select(Conversation).where(Conversation.conversation_id == conversation_id))
        conversation = conversation.scalars().one_or_none()
        if conversation is None:
            raise InvalidParamsException("errors.conversationNotFound")
        conversation.user_id = user.id
        session.add(conversation)
        await session.commit()
    return response(200)


async def change_user_chat_status(user_id: int, status: ChatStatus):
    async with get_async_session_context() as session:
        user = await session.get(User, user_id)
        user.chat_status = status
        session.add(user)
        await session.commit()
        await session.refresh(user)
    return user


@router.delete("/conv", tags=["conversation"])
async def delete_all_conversation(_user: User = Depends(current_super_user)):
    await api.revchatgpt.chatgpt_manager.clear_conversations()
    async with get_async_session_context() as session:
        await session.execute(delete(Conversation))
        await session.commit()
    return response(200)


@router.patch("/conv/{conversation_id}/gen_title", tags=["conversation"], response_model=ConversationSchema)
async def generate_conversation_title(message_id: str, conversation: Conversation = Depends(get_conversation_by_id)):
    if conversation.title is not None:
        raise InvalidParamsException("errors.conversationTitleAlreadyGenerated")
    async with get_async_session_context() as session:
        result = await api.revchatgpt.chatgpt_manager.generate_conversation_title(conversation.id, message_id)
        if result["title"]:
            conversation.title = result["title"]
            session.add(conversation)
            await session.commit()
            await session.refresh(conversation)
        else:
            raise InvalidParamsException(f"{result['message']}")
    result = jsonable_encoder(conversation)
    return result


@router.websocket("/conv")
async def ask(websocket: WebSocket):
    """
    利用 WebSocket 实时更新 ChatGPT 回复.

    客户端第一次连接：发送 { message, conversation_id?, parent_id?, use_paid?, timeout? }
        conversation_id 为空则新建会话，否则回复 parent_id 指定的消息
    服务端返回格式：{ type, tip, message, conversation_id, parent_id, use_paid, title }
    其中：type 可以为 "waiting" / "message" / "title"
    """

    await websocket.accept()
    user = await websocket_auth(websocket)
    logger.debug(f"{user.username} connected to websocket")
    websocket.scope["auth_user"] = user

    if user is None:
        await websocket.close(1008, "errors.unauthorized")
        return

    if user.chat_status != ChatStatus.idling:
        await websocket.close(1008, "errors.cannotConnectMoreThanOneClient")
        return

    # 读取用户输入
    params = await websocket.receive_json()
    message = params.get("message", None)
    conversation_id = params.get("conversation_id", None)
    parent_id = params.get("parent_id", None)
    model_name = params.get("model_name")
    # timeout = params.get("timeout", 30)  # default 30s
    timeout = config.get("ask_timeout", 300)
    new_title = params.get("new_title", None)

    if message is None:
        await websocket.close(1007, "errors.missingMessage")
        return
    if parent_id is not None and conversation_id is None:
        await websocket.close(1007, "errors.missingConversationId")
        return

    is_new_conv = conversation_id is None
    conversation = None
    if not is_new_conv:
        conversation = await get_conversation_by_id(conversation_id, user)
        model_name = model_name or conversation.model_name
    else:
        model_name = model_name or ChatModels.default

    if isinstance(model_name, str):
        model_name = ChatModels(model_name)
    if model_name == ChatModels.paid and not user.can_use_paid:
        await websocket.close(1007, "errors.userNotAllowToUsePaidModel")
        return
    if (model_name == ChatModels.gpt4 or model_name == ChatModels.gpt4_mobile) and not user.can_use_gpt4:
        await websocket.close(1007, "errors.userNotAllowToUseGPT4Model")
        return
    if model_name in [ChatModels.gpt4, ChatModels.gpt4_mobile, ChatModels.paid] and not config.get("chatgpt_paid", False):
        await websocket.close(1007, "errors.paidModelNotAvailable")
        return

    # 判断是否能新建对话，以及是否能继续提问
    async with get_async_session_context() as session:
        user_conversations_count = await session.execute(
            select(func.count(Conversation.id)).filter(and_(Conversation.user_id == user.id, Conversation.is_valid)))
        user_conversations_count = user_conversations_count.scalar()
        if is_new_conv and user.max_conv_count != -1 and user_conversations_count >= user.max_conv_count:
            await websocket.close(1008, "errors.maxConversationCountReached")
            return
        if user.available_ask_count != -1 and user.available_ask_count <= 0:
            await websocket.close(1008, "errors.noAvailableAskCount")
            return
        if user.available_gpt4_ask_count != -1 and user.available_gpt4_ask_count <= 0 and (model_name == ChatModels.gpt4 or model_name == ChatModels.gpt4_mobile):
            await websocket.close(1008, "errors.noAvailableGPT4AskCount")
            return

    if api.revchatgpt.chatgpt_manager.is_busy():
        await websocket.send_json({
            "type": "queueing",
            "tip": "tips.queueing"
        })

    websocket_code = 1001
    websocket_reason = "tips.terminated"

    is_completed = False
    is_canceled = False
    has_got_reply = False
    ask_start_time = None
    queueing_start_time = None

    def check_message(msg: str):
        url = config.get("chatgpt_base_url")
        if url and url in msg:
            return msg.replace(url, "<chatgpt_base_url>")

    try:
        # 标记用户为 queueing
        await change_user_chat_status(user.id, ChatStatus.queueing)
        # is_queueing = True
        queueing_start_time = time.time()
        async with api.revchatgpt.chatgpt_manager.semaphore:
            is_queueing = False
            try:
                await change_user_chat_status(user.id, ChatStatus.asking)
                await websocket.send_json({
                    "type": "waiting",
                    "tip": "tips.waiting"
                })
                ask_start_time = time.time()
                api.revchatgpt.chatgpt_manager.reset_chat()
                async for data in api.revchatgpt.chatgpt_manager.ask(message, conversation_id, parent_id, timeout,
                                                                     model_name):
                    has_got_reply = True
                    reply = {
                        "type": "message",
                        "message": data["message"],
                        "conversation_id": data["conversation_id"],
                        "parent_id": data["parent_id"],
                        "model_name": data["model"],
                    }
                    await websocket.send_json(reply)
                    if conversation_id is None:
                        conversation_id = data["conversation_id"]
                is_completed = True
            except Exception as e:
                # 修复 message 为 None 时的错误
                if str(e).startswith("Field missing"):
                    logger.warning(str(e))
                else:
                    raise e
            finally:
                api.revchatgpt.chatgpt_manager.reset_chat()

    except ConnectionClosed:
        # print("websocket aborted", e.code)
        is_canceled = True
    except requests.exceptions.Timeout:
        logger.warning(str(e))
        await websocket.send_json({
            "type": "error",
            "tip": "errors.timeout"
        })
        websocket_code = 1001
        websocket_reason = "errors.timout"
    except revChatGPTError as e:
        logger.error(str(e))
        message = check_message(f"{e.source} {e.code}: {e.message}")
        await websocket.send_json({
            "type": "error",
            "tip": "errors.chatgptResponseError",
            "message": message
        })
        websocket_code = 1001
        websocket_reason = "errors.chatgptResponseError"
    except HTTPError as e:
        logger.error(str(e))
        message = check_message(str(e))
        await websocket.send_json({
            "type": "error",
            "tip": "errors.httpError",
            "message": message
        })
        websocket_code = 1014
        websocket_reason = "errors.httpError"
    except Exception as e:
        logger.error(str(e))
        message = check_message(str(e))
        await websocket.send_json({
            "type": "error",
            "tip": "errors.unknownError",
            "message": message
        })
        websocket_code = 1011
        websocket_reason = "errors.unknownError"

    ask_stop_time = time.time()

    queueing_time = ask_stop_time - queueing_start_time
    queueing_time = round(queueing_time, 3)
    if ask_start_time is not None:
        ask_time = ask_stop_time - ask_start_time
        ask_time = round(ask_time, 3)
    else:
        ask_time = None

    if is_completed:
        logger.debug(
            f"finished ask {conversation_id} ({model_name}), user: {user.id}, "
            f"ask: {ask_time}s, total: {queueing_time}s")
        websocket_code = 1000
        websocket_reason = "tips.finished"
    elif is_canceled:
        if has_got_reply:
            logger.debug(
                f"canceled ask {conversation_id} ({model_name}) while replying, user: {user.id}, "
                f"ask: {ask_time}s, total: {queueing_time}s")
        elif is_queueing:
            logger.debug(
                f"canceled ask {conversation_id} ({model_name}) while queueing, user: {user.id}, "
                f"total: {queueing_time}s")
        else:
            logger.debug(
                f"canceled ask {conversation_id} ({model_name}) before replying, user: {user.id}, "
                f"total: {queueing_time}s")
    else:
        logger.debug(
            f"terminated ask {conversation_id} ({model_name}) because of error")

    try:
        if has_got_reply:
            async with get_async_session_context() as session:
                # 若新建了对话，则添加到数据库
                if is_new_conv and conversation_id is not None:
                    # 设置默认标题
                    try:
                        if new_title is not None:
                            await api.revchatgpt.chatgpt_manager.set_conversation_title(conversation_id, new_title)
                    except Exception as e:
                        logger.warning(e)
                    finally:
                        current_time = datetime.utcnow()
                        conversation = Conversation(conversation_id=conversation_id, title=new_title,
                                                    user_id=user.id,
                                                    model_name=model_name, create_time=current_time,
                                                    active_time=current_time)
                        session.add(conversation)
                # 更新 conversation
                if not is_new_conv:
                    conversation = await session.get(Conversation, conversation.id)  # 此前的 conversation 属于另一个session
                    conversation.active_time = datetime.utcnow()
                    if conversation.model_name != model_name:
                        conversation.model_name = model_name
                    session.add(conversation)

                # 扣除一次对话次数
                # 这里的逻辑是：available_ask_count 是总的对话次数，available_gpt4_ask_count 是 GPT4 的对话次数
                # 如果都有限制，则都要扣除一次
                # 如果 available_ask_count 不限但是 available_gpt4_ask_count 限制，则只扣除 available_gpt4_ask_count
                if user.available_ask_count != -1 or user.available_gpt4_ask_count != -1:
                    user = await session.get(User, user.id)
                    if user.available_ask_count != -1:
                        assert user.available_ask_count > 0
                        user.available_ask_count -= 1
                    if (model_name == ChatModels.gpt4 or model_name == ChatModels.gpt4_mobile) and user.available_gpt4_ask_count != -1:
                        assert user.available_gpt4_ask_count > 0
                        user.available_gpt4_ask_count -= 1
                    session.add(user)
                await session.commit()

                # 写入到 scope 中，供统计
                g.ask_log_queue.enqueue(
                    (user.id, model_name.value, ask_time, queueing_time))
    except Exception as e:
        raise e
    finally:
        await change_user_chat_status(user.id, ChatStatus.idling)
        await websocket.close(websocket_code, websocket_reason)
```

## File: backend/api/routers/status.py
```python
from fastapi import Depends, APIRouter

import api.revchatgpt
from api import globals as g
from api.models import User
from api.routers.system import check_users
from api.schema import ServerStatusSchema
from api.users import current_active_user

router = APIRouter()


@router.get("/status", tags=["status"], response_model=ServerStatusSchema)
async def get_server_status(_user: User = Depends(current_active_user)):
    """普通用户获取服务器状态"""
    refresh_cache = _user.is_superuser
    active_user_in_5m, active_user_in_1h, active_user_in_1d, queueing_count, _ = await check_users(refresh_cache)
    result = ServerStatusSchema(
        active_user_in_5m=active_user_in_5m,
        active_user_in_1h=active_user_in_1h,
        active_user_in_1d=active_user_in_1d,
        is_chatbot_busy=api.revchatgpt.chatgpt_manager.is_busy(),
        chatbot_waiting_count=queueing_count
    )
    return result
```

## File: backend/api/routers/system.py
```python
import os
import random
from datetime import datetime, timedelta

from fastapi import APIRouter, Depends
from sqlalchemy import select

import api.enums
import api.globals as g
import api.globals as g

config = g.config
from api.database import get_async_session_context
from api.enums import ChatStatus
from api.models import User, Conversation
from api.schema import ServerStatusSchema, LogFilterOptions, SystemInfo, RequestStatistics
from api.users import current_super_user
from utils.logger import get_logger

logger = get_logger(__name__)

router = APIRouter()

check_users_cache = None
check_users_cache_last_update_time: datetime | None = None

CACHE_DURATION_SECONDS = 0  # currently do not cache, for there seems no significant performance improvement


async def check_users(refresh_cache: bool = False):
    global check_users_cache
    global check_users_cache_last_update_time

    if refresh_cache:
        check_users_cache = None
        check_users_cache_last_update_time = None
    if check_users_cache is not None and check_users_cache_last_update_time is not None:
        if check_users_cache_last_update_time > datetime.utcnow() - timedelta(seconds=CACHE_DURATION_SECONDS):
            # logger.debug("Using cached check_users result")
            return check_users_cache
    # logger.debug("Refreshing check_users cache")
    check_users_cache_last_update_time = datetime.utcnow()
    async with get_async_session_context() as session:
        users = await session.execute(select(User))
        users = users.scalars().all()
    queueing_count = 0
    active_user_in_5m = 0
    active_user_in_1h = 0
    active_user_in_1d = 0
    current_time = datetime.utcnow()
    for user in users:
        if not user.active_time:
            continue
        if user.chat_status == ChatStatus.queueing:
            queueing_count += 1
        if user.is_superuser:  # 管理员不计入在线人数
            continue
        if user.active_time > current_time - timedelta(minutes=5):
            active_user_in_5m += 1
        if user.active_time > current_time - timedelta(hours=1):
            active_user_in_1h += 1
        if user.active_time > current_time - timedelta(days=1):
            active_user_in_1d += 1

    check_users_cache = (active_user_in_5m, active_user_in_1h, active_user_in_1d, queueing_count, users)
    return check_users_cache


@router.get("/system/info", tags=["system"], response_model=SystemInfo)
async def get_system_info(_user: User = Depends(current_super_user)):
    active_user_in_5m, active_user_in_1h, active_user_in_1d, queueing_count, users = await check_users(
        refresh_cache=True)
    async with get_async_session_context() as session:
        conversations = await session.execute(select(Conversation))
        conversations = conversations.scalars().all()
    result = SystemInfo(
        startup_time=g.startup_time,
        total_user_count=len(users),
        total_conversation_count=len(conversations),
        valid_conversation_count=len([c for c in conversations if c.is_valid]),
    )
    return result


START_TIMESTAMP = 1672502400  # 2023-01-01 00:00:00


def make_fake_requests_count(total=100, max=500):
    result = {}
    start_stage = START_TIMESTAMP // g.request_log_counter_interval
    for i in range(total):
        result[start_stage + i] = [random.randint(0, max), [1]]
    return result


def make_fake_ask_records(total=100, days=2):
    result = []
    model_names = list(api.enums.ChatModels.__members__.keys())
    for i in range(total):
        ask_time = random.random() * 60 + 1
        total_time = ask_time + random.random() * 30
        result.append([
            [
                # random.randint(1, 10),  # user_id
                1,
                model_names[random.randint(0, len(model_names) - 1)].value,  # model_name
                ask_time,
                total_time
            ],
            START_TIMESTAMP + random.random() * 60 * 60 * 24 * days,  # ask_time
        ])
    return result


@router.get("/system/request_statistics", tags=["system"], response_model=RequestStatistics)
async def get_request_statistics(_user: User = Depends(current_super_user)):
    result = RequestStatistics(
        request_counts_interval=g.request_log_counter_interval,
        request_counts=dict(g.request_log_counter.counter),
        # request_counts=make_fake_requests_count(20, 500),
        ask_records=list(g.ask_log_queue.queue)
        # ask_records=make_fake_ask_records(3000, 7)
    )
    return result


def read_last_n_lines(file_path, n, exclude_key_words=None):
    if exclude_key_words is None:
        exclude_key_words = []
    try:
        with open(file_path, "r") as f:
            lines = f.readlines()[::-1]
    except FileNotFoundError:
        return [f"File not found: {file_path}"]
    last_n_lines = []
    for line in lines:
        if len(last_n_lines) >= n:
            break
        if any([line.find(key_word) != -1 for key_word in exclude_key_words]):
            continue
        last_n_lines.append(line)
    return last_n_lines[::-1]


@router.post("/system/proxy_logs", tags=["system"])
async def get_proxy_logs(_user: User = Depends(current_super_user), options: LogFilterOptions = LogFilterOptions()):
    lines = read_last_n_lines(
        os.path.join(config.get("log_dir", "logs"), "reverse_proxy.log"),
        options.max_lines,
        options.exclude_keywords
    )
    return lines


@router.post("/system/server_logs", tags=["system"])
async def get_server_logs(_user: User = Depends(current_super_user), options: LogFilterOptions = LogFilterOptions()):
    lines = read_last_n_lines(
        g.server_log_filename,
        options.max_lines,
        options.exclude_keywords
    )
    return lines
```

## File: backend/api/routers/users.py
```python
from fastapi_users.exceptions import UserAlreadyExists, InvalidPasswordException
from fastapi_users.router import ErrorCode
from sqlalchemy.future import select
from starlette import status
from starlette.requests import Request

from api.database import get_async_session_context, get_user_db_context
from api.exceptions import AuthorityDenyException, InvalidParamsException
from api.models import User
from api.response import response
from api.schema import UserRead, UserUpdate, UserCreate, LimitSchema
from api.users import auth_backend, fastapi_users, current_active_user, get_user_manager_context, current_super_user

from fastapi import APIRouter, Depends, HTTPException

router = APIRouter()

router.include_router(
    fastapi_users.get_auth_router(auth_backend), prefix="/auth", tags=["auth"]
)

router.include_router(
    fastapi_users.get_reset_password_router(),
    prefix="/auth",
    tags=["auth"],
)


# router.include_router(
#     fastapi_users.get_register_router(UserRead, UserCreate),
#     prefix="/auth",
#     tags=["auth"],
# )
@router.post("/auth/register", tags=["auth"])
async def register(
        request: Request,
        user_create: UserCreate,
        _user: User = Depends(current_super_user),
):
    try:
        async with get_async_session_context() as session:
            async with get_user_db_context(session) as user_db:
                async with get_user_manager_context(user_db) as user_manager:
                    created_user = await user_manager.create(
                        user_create, safe=True, request=request
                    )
    except UserAlreadyExists:
        raise HTTPException(
            status_code=status.HTTP_400_BAD_REQUEST,
            detail=ErrorCode.REGISTER_USER_ALREADY_EXISTS,
        )
    except InvalidPasswordException as e:
        raise HTTPException(
            status_code=status.HTTP_400_BAD_REQUEST,
            detail={
                "code": ErrorCode.REGISTER_INVALID_PASSWORD,
                "reason": e.reason,
            },
        )

    return UserRead.from_orm(created_user)


@router.get("/user", tags=["user"])
async def get_all_users(_user: User = Depends(current_super_user)):
    async with get_async_session_context() as session:
        r = await session.execute(select(User))
        results = r.scalars().all()
        return results


@router.patch("/user/{user_id}/reset-password", tags=["user"])
async def reset_password(user_id: int = None, new_password: str = None, _user: User = Depends(current_active_user)):
    if not new_password:
        raise InvalidParamsException("errors.newPasswordRequired")
    if _user.id != user_id and not _user.is_superuser:
        raise AuthorityDenyException("errors.noPermission")
    async with get_async_session_context() as session:
        async with get_user_db_context(session) as db:
            async with get_user_manager_context(db) as user_manager:
                result = await session.get(User, user_id)
                target_user = result
                if target_user is None:
                    raise InvalidParamsException("errors.userNotExist")
                target_user.hashed_password = user_manager.password_helper.hash(new_password)
                session.add(target_user)
                await session.commit()
                return response(200)


@router.post("/user/{user_id}/limit", tags=["user"])
async def update_limit(limit: LimitSchema, user_id: int = None, _user: User = Depends(current_super_user)):
    async with get_async_session_context() as session:
        target_user: User = await session.get(User, user_id)
        if target_user is None:
            raise InvalidParamsException("errors.userNotExist")

        for attr, value in limit.dict(exclude_unset=True).items():
            if value is not None:
                setattr(target_user, attr, value)

        # 使用**kargs类似的写法，但是跳过None值
        session.add(target_user)
        await session.commit()
        return response(200)


router.include_router(
    fastapi_users.get_users_router(UserRead, UserUpdate),
    prefix="/user",
    tags=["user"],
)
```

## File: backend/api/schema.py
```python
import uuid
import datetime
from typing import List

from fastapi_users import schemas
from pydantic import Field, BaseModel, validator

from api.enums import ChatStatus, ChatModels


class UserRead(schemas.BaseUser[int]):
    id: int
    username: str
    nickname: str
    email: str
    active_time: datetime.datetime | None

    chat_status: ChatStatus

    can_use_paid: bool
    can_use_gpt4: bool
    max_conv_count: int | None
    available_ask_count: int | None
    available_gpt4_ask_count: int | None

    is_superuser: bool
    is_active: bool
    is_verified: bool


class LimitSchema(BaseModel):
    can_use_paid: bool | None = None
    can_use_gpt4: bool | None = None
    max_conv_count: int | None = None
    available_ask_count: int | None = None
    available_gpt4_ask_count: int | None = None


class UserUpdate(schemas.BaseUser[int]):
    nickname: str
    email: str = None


class UserCreate(schemas.BaseUserCreate):
    username: str
    nickname: str
    email: str
    can_use_paid: bool = False
    max_conv_count: int = -1
    available_ask_count: int = -1

    class Config:
        orm_mode = True


class ConversationSchema(BaseModel):
    id: int = -1
    conversation_id: uuid.UUID = None
    title: str = None
    user_id: int = None
    is_valid: bool = None
    model_name: ChatModels = None
    create_time: datetime.datetime = None
    active_time: datetime.datetime = None

    class Config:
        use_enum_values = True


class ServerStatusSchema(BaseModel):
    active_user_in_5m: int = None
    active_user_in_1h: int = None
    active_user_in_1d: int = None
    is_chatbot_busy: bool = None
    chatbot_waiting_count: int = None


class RequestStatistics(BaseModel):
    request_counts_interval: int
    request_counts: dict[int, list]  # {timestage: [count, [user_ids]]}
    ask_records: list  # list of (ask, time_used), timestamp.


class SystemInfo(BaseModel):
    startup_time: float
    total_user_count: int
    total_conversation_count: int
    valid_conversation_count: int


class LogFilterOptions(BaseModel):
    max_lines: int = 100
    exclude_keywords: list[str] = None

    @validator("max_lines")
    def max_lines_must_be_positive(cls, v):
        if v <= 0:
            raise ValueError("max_lines must be positive")
        return v
```

## File: backend/api/users.py
```python
import contextlib
from datetime import datetime
from typing import Any

from fastapi.security import OAuth2PasswordRequestForm
from starlette.websockets import WebSocket

import api.exceptions
import api.globals as g

from typing import Optional

from fastapi import Depends, Request, HTTPException
from fastapi_users import BaseUserManager, FastAPIUsers, models, IntegerIDMixin, InvalidID, schemas
from fastapi_users.authentication import CookieTransport, AuthenticationBackend
from fastapi_users.authentication import JWTStrategy

from api.database import get_user_db, get_async_session_context, get_user_db_context
from api.models import User

from sqlalchemy import select, Integer
from fastapi_users.models import UP
from utils.logger import get_logger

config = g.config
logger = get_logger(__name__)

# 使用 cookie + JWT
# 参考 https://fastapi-users.github.io/fastapi-users/10.2/configuration/full-example/

cookie_transport = CookieTransport(
    cookie_max_age=config.get("cookie_max_age", 86400),
    cookie_name="user_auth",
    cookie_httponly=False,
    cookie_secure=False,
)


# auth backend

def get_jwt_strategy() -> JWTStrategy:
    return JWTStrategy(secret=config.get("jwt_secret"), lifetime_seconds=config.get("jwt_lifetime_seconds", 86400))


auth_backend = AuthenticationBackend(
    name="jwt",
    transport=cookie_transport,
    get_strategy=get_jwt_strategy,
)

# UserManager

SECRET = config.get("user_secret")


async def get_by_username(username: str) -> Optional[UP]:
    async with get_async_session_context() as session:
        user = await session.execute(select(User).filter(User.username == username))
        return user.scalar_one_or_none()


class UserManager(IntegerIDMixin, BaseUserManager[User, Integer]):
    async def create(self, user_create: schemas.UC, safe: bool = False, request: Optional[Request] = None) -> models.UP:
        # 检查用户名、手机、邮箱是否已经存在
        async with get_async_session_context() as session:
            if (await session.execute(select(User).filter(User.username == user_create.username))).scalar_one_or_none():
                raise api.exceptions.InvalidRequestException("Username already exists")
            if (await session.execute(select(User).filter(User.email == user_create.email))).scalar_one_or_none():
                raise api.exceptions.InvalidRequestException("Email already exists")
        return await super().create(user_create, safe, request)

    reset_password_token_secret = SECRET
    verification_token_secret = SECRET

    def parse_id(self, value: Any) -> int:
        try:
            return int(value)
        except ValueError as e:
            raise InvalidID() from e

    async def authenticate(
            self, credentials: OAuth2PasswordRequestForm
    ) -> Optional[models.UP]:
        """
        Authenticate and return a user following an email and a password.

        Will automatically upgrade password hash if necessary.

        :param credentials: The user credentials.
        """
        user = await get_by_username(credentials.username)

        if user is None:
            # Run the hasher to mitigate timing attack
            # Inspired from Django: https://code.djangoproject.com/ticket/20760
            self.password_helper.hash(credentials.password)
            return None

        verified, updated_password_hash = self.password_helper.verify_and_update(
            credentials.password, user.hashed_password
        )
        if not verified:
            return None
        # Update password hash to a more robust one if needed
        if updated_password_hash is not None:
            await self.user_db.update(user, {"hashed_password": updated_password_hash})

        return user

    async def on_after_register(self, user: User, request: Optional[Request] = None):
        print(f"User {user.id} has registered.")

    async def on_after_forgot_password(
            self, user: User, token: str, request: Optional[Request] = None
    ):
        print(f"User {user.id} has forgot their password. Reset token: {token}")

    async def on_after_request_verify(
            self, user: User, token: str, request: Optional[Request] = None
    ):
        print(
            f"Verification requested for user {user.id}. Verification token: {token}")


async def websocket_auth(websocket: WebSocket) -> User | None:
    user = None
    try:
        cookie = websocket._cookies[config.get("cookie_name", "user_auth")]
        async with get_async_session_context() as session:
            async with get_user_db_context(session) as user_db:
                async with get_user_manager_context(user_db) as user_manager:
                    # user = await get_jwt_strategy().read_token(cookie, user_manager)
                    user, _ = await fastapi_users.authenticator._authenticate(
                        active=True,
                        user_manager=user_manager,
                        jwt=cookie,
                        strategy_jwt=get_jwt_strategy(),
                    )
    finally:
        return user


async def get_user_manager(user_db=Depends(get_user_db)):
    yield UserManager(user_db)


get_user_manager_context = contextlib.asynccontextmanager(get_user_manager)

# FastAPIUsers 实例，注意不要和 fastapi_users 包混淆
fastapi_users = FastAPIUsers[User, Integer](get_user_manager, [auth_backend])

__current_active_user = fastapi_users.current_user(active=True)


async def current_active_user(request: Request, user: User = Depends(__current_active_user)):
    current_time = datetime.utcnow()
    user.active_time = current_time
    try:
        async with get_async_session_context() as session:
            user_update = await session.get(User, user.id)
            user_update.active_time = current_time
            session.add(user_update)
            await session.commit()
        request.scope["auth_user"] = user
    except Exception as e:
        raise e
    finally:
        return user


# current_super_user = fastapi_users.current_user(active=True, superuser=True)

async def current_super_user(user: User = Depends(current_active_user)):
    if not user.is_superuser:
        raise api.exceptions.AuthorityDenyException("You are not super user")
    return user
```

## File: backend/logging_config.yaml
```yaml
version: 1
formatters:
  simple:
    format: "%(asctime)s.%(msecs)03d %(levelname)8s: [%(name)s]\t%(message)s"
    datefmt: '%Y/%m/%d %H:%M:%S'
  proxy-output:
    format: "%(message)s"
    datefmt: '%Y/%m/%d %H:%M:%S'
  colored:
    (): colorlog.ColoredFormatter
    format: "%(asctime)s.%(msecs)03d %(log_color)s%(levelname)8s%(reset)s: %(cyan)s[%(name)s]%(reset)s %(message)s"
    datefmt: '%Y/%m/%d %H:%M:%S'
handlers:
  file_handler:
    class: logging.handlers.RotatingFileHandler
    formatter: simple
    encoding: utf-8
    level: DEBUG
    filename: cws.log
    maxBytes: 10485760  # 10MB
  console_handler:
    class: logging.StreamHandler
    formatter: colored
    level: DEBUG

root:
  level: DEBUG
  handlers: []
loggers:
  uvicorn.error:
    level: INFO
    handlers: [console_handler, file_handler]
  uvicorn.access:
    level: INFO
    handlers: []  # disable uvicorn access log
  cws:
    level: DEBUG
    handlers: [console_handler, file_handler]
  sqlalchemy:
    level: WARN
    handlers: [console_handler, file_handler]
  alembic:
    level: INFO
    handlers: [console_handler, file_handler]
```

## File: backend/main.py
```python
import asyncio
import time
from datetime import datetime

import aiocron

import api.revchatgpt
from api.middlewares import AccessLoggerMiddleware, StatisticsMiddleware
from httpx import HTTPError
import uvicorn

from fastapi import FastAPI
from fastapi.exceptions import RequestValidationError
from sqlalchemy import select
from starlette.exceptions import HTTPException as StarletteHTTPException

import api.globals as g
import os
import utils.store_statistics
from utils.sync_conversations import sync_conversations

from api.enums import ChatStatus
from api.models import Conversation, User
from api.response import CustomJSONResponse, PrettyJSONResponse, handle_exception_response
from api.database import create_db_and_tables, get_async_session_context
from api.exceptions import SelfDefinedException
from api.routers import users, chat, system, status
from fastapi.middleware import Middleware
from fastapi.middleware.cors import CORSMiddleware

from utils.logger import setup_logger, get_log_config, get_logger
from utils.proxy import close_reverse_proxy
from utils.create_user import create_user

import dateutil.parser
from revChatGPT.typings import Error as revChatGPTError

config = g.config

setup_logger()

logger = get_logger(__name__)

app = FastAPI(
    default_response_class=CustomJSONResponse,
    middleware=[
        Middleware(AccessLoggerMiddleware, format='%(client_addr)s | %(request_line)s | %(status_code)s | %(M)s ms',
                   logger=get_logger("access")),
        Middleware(StatisticsMiddleware)]
)

app.include_router(users.router)
app.include_router(chat.router)
app.include_router(system.router)
app.include_router(status.router)

origins = config.get("cors_allow_origins", [
    "http://localhost",
    "http://127.0.0.1"
])

# 解决跨站问题
app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# 定义若干异常处理器


@app.exception_handler(StarletteHTTPException)
async def http_exception_handler(request, exc):
    return handle_exception_response(exc)


@app.exception_handler(RequestValidationError)
async def validation_exception_handler(request, exc):
    return handle_exception_response(exc)


@app.exception_handler(SelfDefinedException)
async def validation_exception_handler(request, exc):
    return handle_exception_response(exc)


@app.exception_handler(revChatGPTError)
async def validation_exception_handler(request, exc):
    return handle_exception_response(exc)


@app.on_event("startup")
async def on_startup():
    await create_db_and_tables()
    logger.info("database initialized")
    g.startup_time = time.time()

    utils.store_statistics.load()

    if config.get("create_initial_admin_user", False):
        await create_user(config.get("initial_admin_username"),
                          "admin",
                          "admin@admin.com",
                          config.get("initial_admin_password"),
                          is_superuser=True)

    if config.get("create_initial_user", False):
        await create_user(config.get("initial_user_username"),
                          "user",
                          "user@user.com",
                          config.get("initial_user_password"),
                          is_superuser=False)

    if not config.get("sync_conversations_on_startup", True):
        return

    # 重置所有用户chat_status
    async with get_async_session_context() as session:
        r = await session.execute(select(User))
        results = r.scalars().all()
        for user in results:
            user.chat_status = ChatStatus.idling
            session.add(user)
        await session.commit()

    # 运行 Proxy Server
    if config.get("run_reverse_proxy", False):
        from utils.proxy import run_reverse_proxy
        run_reverse_proxy()
        await asyncio.sleep(2)  # 等待 Proxy Server 启动

    logger.info(f"Using {g.config.get('chatgpt_base_url', 'env: ' + os.environ.get('CHATGPT_BASE_URL', '<default_bypass>'))} as ChatGPT base url")

    # 获取 ChatGPT 对话，并同步数据库
    if not config.get("sync_conversations_on_startup", True):
        logger.info("Sync conversations on startup disabled. Jumping...")
        return  # 跳过同步对话
    else:
        await sync_conversations()

    @aiocron.crontab('*/5 * * * *', loop=asyncio.get_event_loop())
    async def dump_stats():
        utils.store_statistics.dump(print_log=False)

    if config.get("sync_conversations_regularly", True):
        logger.info("Sync conversations regularly enabled, will sync conversations every 12 hours.")

        # 默认每隔 12 小时同步一次
        @aiocron.crontab('0 */12 * * *', loop=asyncio.get_event_loop())
        async def sync_conversations_regularly():
            await sync_conversations()


# 关闭时
@app.on_event("shutdown")
async def on_shutdown():
    logger.info("On shutdown...")
    close_reverse_proxy()
    utils.store_statistics.dump()


# @api.get("/routes")
# async def root():
#     url_list = [{"name": route.name, "path": route.path, "path_regex": str(route.path_regex)}
#                 for route in api.routes]
#     return PrettyJSONResponse(url_list)


if __name__ == "__main__":
    uvicorn.run(app, host=config.get("host"),
                port=config.get("port"),
                proxy_headers=True,
                forwarded_allow_ips='*',
                log_config=get_log_config(),
                )
```

## File: backend/pyproject.toml
```toml
[tool.poetry]
name = "chatgpt-share-backend"
version = "0.3.15"
description = ""
authors = ["moeakwak <moeakwak@gmail.com>"]
readme = "README.md"
packages = []

[tool.poetry.dependencies]
python = "^3.10"
fastapi = "^0.95.1"
uvicorn = "^0.21.1"
aiosqlite = "^0.18.0"
sqlalchemy = "^2.0.9"
fastapi-users-db-sqlalchemy = "^5.0.0"
revchatgpt = "4.2.3"
greenlet = "^2.0.2"
websockets = "^11.0.1"
setuptools = "^67.6.0"
python-dateutil = "^2.8.2"
pyyaml = "^6.0"
alembic = "^1.10.3"
colorlog = "^6.7.0"
asgiref = "^3.6.0"
aiocron = "^1.8"


[build-system]
requires = ["poetry-core"]
build-backend = "poetry.core.masonry.api"
```

## File: backend/requirements.txt
```
aiocron==1.8
aiohttp==3.8.4
aiosignal==1.3.1
aiosqlite==0.18.0
alembic==1.10.3
anyio==3.6.2
asgiref==3.6.0
async-timeout==4.0.2
attrs==23.1.0
bcrypt==4.0.1
blobfile==2.0.1
certifi==2022.12.7
cffi==1.15.1
charset-normalizer==3.1.0
click==8.1.3
colorlog==6.7.0
croniter==1.3.14
cryptography==40.0.2
dnspython==2.3.0
email-validator==1.3.1
fastapi==0.95.1
fastapi-users==10.4.2
fastapi-users-db-sqlalchemy==5.0.0
filelock==3.9.0
frozenlist==1.3.3
greenlet==2.0.2
h11==0.14.0
httpcore==0.17.0
httpx==0.24.0
idna==3.4
lxml==4.9.2
makefun==1.15.1
Mako==1.2.4
MarkupSafe==2.1.2
multidict==6.0.4
openai==0.27.4
OpenAIAuth==0.3.6
passlib==1.7.4
prompt-toolkit==3.0.38
pycparser==2.21
pycryptodomex==3.17
pydantic==1.10.7
PyJWT==2.6.0
PySocks==1.7.1
python-dateutil==2.8.2
python-multipart==0.0.6
pytz-deprecation-shim==0.1.0.post0
PyYAML==6.0
regex==2023.3.23
requests==2.28.2
revChatGPT==4.2.3
rfc3986==1.5.0
six==1.16.0
sniffio==1.3.0
socksio==1.0.0
SQLAlchemy==2.0.9
starlette==0.26.1
tiktoken==0.3.3
tqdm==4.65.0
typing_extensions==4.5.0
tzdata==2023.3
tzlocal==4.3
urllib3==1.26.15
uvicorn==0.21.1
wcwidth==0.2.6
websockets==11.0.1
yarl==1.8.2
```

## File: backend/utils/__init__.py
```python

```

## File: backend/utils/common.py
```python
import asyncio, threading


def get_conversation_model(conversation) -> str:
    result = None
    try:
        current_node = conversation["current_node"]
        while current_node:
            node = conversation["mapping"][current_node]
            result = node["message"]["metadata"]["model_slug"]
            if result:
                break
            current_node = node["parent"]
    finally:
        return result


def async_wrap_iter(it):
    """Wrap blocking iterator into an asynchronous one"""
    loop = asyncio.get_event_loop()
    q = asyncio.Queue(1)
    exception = None
    _END = object()

    async def yield_queue_items():
        while True:
            next_item = await q.get()
            if next_item is _END:
                break
            yield next_item
        if exception is not None:
            # the iterator has raised, propagate the exception
            raise exception

    def iter_to_queue():
        nonlocal exception
        try:
            for item in it:
                # This runs outside the event loop thread, so we
                # must use thread-safe API to talk to the queue.
                asyncio.run_coroutine_threadsafe(q.put(item), loop).result()
        except Exception as e:
            exception = e
        finally:
            asyncio.run_coroutine_threadsafe(q.put(_END), loop).result()

    threading.Thread(target=iter_to_queue).start()
    return yield_queue_items()
```

## File: backend/utils/create_user.py
```python
from api.schema import UserCreate
from api.users import get_user_manager_context
from api.database import get_user_db_context, get_async_session_context
from utils.logger import get_logger

logger = get_logger(__name__)


async def create_user(username, nickname: str, email: str, password: str, is_superuser: bool = False, **kwargs):
    try:
        async with get_async_session_context() as session:
            async with get_user_db_context(session) as user_db:
                async with get_user_manager_context(user_db) as user_manager:
                    user = await user_manager.create(
                        UserCreate(
                            username=username, nickname=nickname,
                            email=email, password=password, is_superuser=is_superuser,
                            **kwargs
                        )
                    )
                    logger.info(f"User created: {user}")
                    return user
    except Exception as e:
        logger.info(f"Create User {username} Error: {e}")
        return None
```

## File: backend/utils/data_types.py
```python
import time
from collections import OrderedDict, deque
from typing import Deque


class RequestCounter:
    """
    用于统计一段时间内的请求计数和请求用户id
    """
    def __init__(self, time_window: int = None, interval: int = None):
        if time_window % interval != 0 or time_window <= 0 or interval <= 0 or time_window < interval:
            raise ValueError("time_window must be a multiple of duration, and both must be positive")
        self.time_window = time_window or 3 * 24 * 60 * 60
        self.interval = interval or 30
        # counter: {timestage: [count, [user_ids]]}
        self.counter: OrderedDict[int, tuple[int, list[int]]] = OrderedDict()

    def count(self, user_id: int=None):
        current_time = time.time()
        current_interval = int(current_time // self.interval)

        # 增加当前时间段的计数
        if current_interval not in self.counter:
            self.counter[current_interval] = [1, [user_id] if user_id else []]
        else:
            if user_id:
                user_ids = set(self.counter[current_interval][1]).union({user_id})
                user_ids = list(user_ids)
            else:
                user_ids = self.counter[current_interval][1]
            self.counter[current_interval] = (self.counter[current_interval][0] + 1, user_ids)

        # 删除过期的时间段
        self.remove_expired_intervals()

    def remove_expired_intervals(self):
        current_time = time.time()
        current_interval = int(current_time // self.interval)
        expired_interval = current_interval - (self.time_window // self.interval)

        keys_to_delete = []
        for key in self.counter:
            if key <= expired_interval:
                keys_to_delete.append(key)
            else:
                break

        for key in keys_to_delete:
            del self.counter[key]

    def __repr__(self):
        return f"TimeCounter(time_window={self.time_window}, duration={self.interval}, counter={self.counter})"


class TimeQueue:
    def __init__(self, time_window: int):
        self.time_window = time_window
        self.queue: Deque[tuple[any, float]] = deque()    # (item, time)

    def enqueue(self, item):
        current_time = time.time()
        self.queue.append((item, current_time))
        self.dequeue_expired()

    def dequeue_expired(self):
        current_time = time.time()
        while self.queue and self.queue[0][1] < current_time - self.time_window:
            self.queue.popleft()

    def __len__(self):
        return len(self.queue)

    def __repr__(self):
        return f"TimeQueue(time_window={self.time_window}, data={self.queue}...)"
```

## File: backend/utils/logger.py
```python
# 创建一个用于输出信息的logger，同时输出到控制台和logs/时间.log文件
import logging
import logging.config
import os
from datetime import datetime
import api.globals as g

import yaml

import api.globals as g
config = g.config


def get_log_config():
    with open('logging_config.yaml', 'r') as f:
        log_config = yaml.safe_load(f.read())
    log_config['handlers']['file_handler']['filename'] = g.server_log_filename
    log_config['handlers']['console_handler']['level'] = config.get("console_log_level", "INFO")
    return log_config


def setup_logger():
    log_dir = config.get("log_dir", "logs")
    os.makedirs(log_dir, exist_ok=True)
    g.server_log_filename = os.path.join(log_dir, f"{datetime.now().strftime('%Y%m%d_%H-%M-%S')}.log")
    log_config = get_log_config()
    logging.config.dictConfig(log_config)


def get_logger(name):
    return logging.getLogger(f"cws.{name}")
```

## File: backend/utils/proxy.py
```python
import os.path
import subprocess
import api.globals as g
config = g.config
import api.globals as g

from utils.logger import get_logger

logger = get_logger(__name__)


def run_reverse_proxy():
    if not config.get("chatgpt_paid", False):
        logger.error("You need a ChatGPT Plus account to use the reverse proxy!")
        logger.error("Please set chatgpt_paid to true in config.yaml and restart the server.")
        exit(1)

    proxy_path = config.get("reverse_proxy_binary_path", None)
    if not proxy_path:
        logger.error("You need to set the reverse proxy binary path in config.yaml!")
        exit(1)

    puid = config.get("reverse_proxy_puid")
    env_vars = {"PORT": str(config.get("reverse_proxy_port", 6060))}
    http_proxy = config.get("reverse_proxy_http_proxy")
    if puid:
        env_vars["PUID"] = puid
    if config.get("auto_refresh_reverse_proxy_puid"):
        env_vars["ACCESS_TOKEN"] = config.get("chatgpt_access_token")
    if http_proxy:
        env_vars["http_proxy"] = http_proxy
        logger.info(f"Reverse proxy Using http proxy: {http_proxy}")
    g.reverse_proxy_log_file = open(os.path.join(config.get("log_dir", "logs"), "reverse_proxy.log"), "w",
                                    encoding="utf-8")
    logger.debug(f"Reverse proxy binary path: {proxy_path}")
    g.reverse_proxy_process = subprocess.Popen([proxy_path], env=env_vars, stdout=g.reverse_proxy_log_file,
                                               stderr=g.reverse_proxy_log_file)
    logger.info("Reverse proxy started!")


def close_reverse_proxy():
    if g.reverse_proxy_process:
        g.reverse_proxy_process.kill()
        g.reverse_proxy_process = None
        g.reverse_proxy_log_file.close()
        g.reverse_proxy_log_file = None
        logger.info("Reverse proxy stopped.")
```

## File: backend/utils/store_statistics.py
```python
from collections import OrderedDict, deque

import api.globals as g
import json
import os

from utils.logger import get_logger

logger = get_logger(__name__)


def dump(print_log=True):
    path = g.config.get("data_dir", g.config.get("log_dir", "."))
    path = os.path.join(path, "statistics.json")
    data = {
        "request_log_counter_interval": g.request_log_counter_interval,
        "request_log_counter": g.request_log_counter.counter,
        "ask_log_queue": list(g.ask_log_queue.queue)
    }
    with open(path, "w") as f:
        json.dump(data, f)
    if print_log:
        logger.info(f"Requests statistics dumped to {path}.")


def load():
    path = g.config.get("data_dir", g.config.get("log_dir", "."))
    path = os.path.join(path, "statistics.json")
    logger.debug(f"loading statistics from {path}")
    try:
        with open(path, "r") as f:
            data = json.load(f)

            if g.request_log_counter_interval != data["request_log_counter_interval"]:
                logger.warning("request_log_counter_interval is different from the saved one, counter cleared.")
                return

            for k, v in data["request_log_counter"].items():
                g.request_log_counter.counter[int(k)] = v
            g.request_log_counter.remove_expired_intervals()
            g.ask_log_queue.queue = deque(data["ask_log_queue"])

            logger.info("Requests statistics loaded.")
    except FileNotFoundError:
        logger.info("File statistics.json not found, skip loading statistics.")
    except json.decoder.JSONDecodeError:
        logger.warning("Failed to load statistics.json, skip loading statistics.")
```

## File: backend/utils/sync_conversations.py
```python
import os

import dateutil.parser
from httpx import HTTPError
from sqlalchemy import select

import api.revchatgpt
from api.database import get_async_session_context
from api.models import Conversation
from utils.logger import get_logger
from revChatGPT.typings import Error as revChatGPTError

logger = get_logger(__name__)


async def sync_conversations():
    try:
        logger.info("Syncing conversations...")
        result = await api.revchatgpt.chatgpt_manager.get_conversations()
        logger.info(f"Fetched {len(result)} conversations from ChatGPT account.")
        openai_conversations_map = {conv['id']: conv for conv in result}
        async with get_async_session_context() as session:
            r = await session.execute(select(Conversation))
            results = r.scalars().all()

            for conv_db in results:
                openai_conv = openai_conversations_map.get(conv_db.conversation_id, None)
                if openai_conv:
                    # 同步标题
                    if openai_conv["title"] != conv_db.title:
                        conv_db.title = openai_conv["title"]
                        logger.info(f"Conversation {conv_db.conversation_id} title changed: {conv_db.title}")
                    # 同步时间
                    create_time = dateutil.parser.isoparse(openai_conv["create_time"])
                    if create_time != conv_db.create_time:
                        conv_db.create_time = create_time
                        logger.info(
                            f"Conversation {conv_db.conversation_id} created time changed：{conv_db.create_time}")
                    session.add(conv_db)
                    openai_conversations_map.pop(conv_db.conversation_id)
                else:
                    if conv_db.is_valid:  # 数据库中存在，但 ChatGPT 中（可能）不存在
                        # conv_db.is_valid = False
                        logger.warning(
                            f"Cannot fetch conversation [{conv_db.title}]({conv_db.conversation_id})")
                        # session.add(conv_db)

            # 新增对话
            for openai_conv in openai_conversations_map.values():
                new_conv = Conversation(
                    conversation_id=openai_conv["id"],
                    title=openai_conv["title"],
                    is_valid=True,
                    create_time=dateutil.parser.isoparse(openai_conv["create_time"])
                )
                session.add(new_conv)
                logger.info(
                    f"Conversation [{new_conv.title}]({new_conv.conversation_id}) not recorded, added to database")

            await session.commit()
        logger.info("Sync conversations finished.")
    except revChatGPTError as e:
        logger.error(f"Fetch conversation error (ChatGPTError): {e.source} {e.code}: {e.message}")
        logger.warning("Sync conversations on startup failed!")
    except HTTPError as e:
        logger.error(f"Fetch conversation error (httpx.HTTPError): {str(e)}")
        logger.warning("Sync conversations on startup failed!")
    except Exception as e:
        logger.error(f"Fetch conversation error (unknown): {str(e)}")
        logger.warning("Sync conversations on startup failed!")
```

## File: Caddyfile
```
:80 {
	handle_path /api/* {
		reverse_proxy localhost:8000
	}
	handle /* {
		file_server
		root * /app/dist
		try_files {path} /index.html
	}
}

# example for subdirectory deploy:

# :7777 {
# 	handle_path /chat/api/* {
# 		uri strip_prefix /chat
# 		reverse_proxy :8000
# 	}
# 	handle_path /chat/* {
# 		file_server
# 		root * ./frontend/dist
# 		try_files {path} /index.html
# 	}
# }
```

## File: docker-compose.yaml
```yaml
version: "3"

services:
  chatgpt-web-share:
    image: ghcr.io/moeakwak/chatgpt-web-share:latest
    container_name: chatgpt-web-share
    restart: unless-stopped
    # network_mode: bridge
    ports:
      - 8080:80
    volumes:
      - ./data:/data
      - ./config.yaml:/app/backend/api/config/config.yaml
      - ./logs:/app/logs
    environment:
      - TZ=Asia/Shanghai
      - CHATGPT_BASE_URL=http://go-chatgpt-api:8080/chatgpt/
    depends_on:
      - go-chatgpt-api

  go-chatgpt-api:
    container_name: go-chatgpt-api
    image: linweiyuan/go-chatgpt-api
    ports:
      - 8080:8080
    environment:
      - GO_CHATGPT_API_PROXY=
    restart: unless-stopped
```

## File: Dockerfile
```dockerfile
FROM golang:1.20-alpine AS ProxyBuilder

COPY ChatGPT-Proxy-V4 /app/ChatGPT-Proxy-V4

WORKDIR /app/ChatGPT-Proxy-V4

RUN CGO_ENABLED=0 go build -a -installsuffix cgo .

FROM python:3.10-alpine

ARG PIP_CACHE_DIR=/pip_cache
ARG TARGETARCH

RUN mkdir -p /app/backend

RUN apk add --update caddy
RUN if [ "${TARGETARCH}" = "arm64" ] ; then \
        apk add --no-cache gcc musl-dev libffi-dev \
    ; fi

COPY backend/requirements.txt /tmp/requirements.txt
RUN pip install -r /tmp/requirements.txt

COPY Caddyfile /app/Caddyfile
COPY backend /app/backend
COPY frontend/dist /app/dist
COPY --from=ProxyBuilder /app/ChatGPT-Proxy-V4/ChatGPT-Proxy-V4 /app/backend/ChatGPT-Proxy-V4

WORKDIR /app

EXPOSE 80

COPY startup.sh /app/startup.sh
RUN chmod +x /app/startup.sh; mkdir /data
CMD ["/app/startup.sh"]
```

## File: frontend/.env
```
VITE_API_WEBSOCKET_PROTOCOL=auto
VITE_DISABLE_SENTRY=no
```

## File: frontend/.eslintignore
```
node_modules
.eslintrc.cjs
```

## File: frontend/.eslintrc.cjs
```javascript
module.exports = {
  env: {
    browser: true,
    es2021: true,
  },
  overrides: [],
  parser: 'vue-eslint-parser',
  extends: [
    'plugin:vue/base',
    'eslint:recommended',
    'plugin:vue/vue3-recommended',
    'plugin:vue/essential',
    'plugin:@typescript-eslint/recommended',
    'plugin:import/recommended',
    'plugin:import/typescript',
    // "plugin:prettier/recommended",
    // "eslint-config-prettier",
  ],
  parserOptions: {
    ecmaVersion: 'latest',
    sourceType: 'module',
    parser: '@typescript-eslint/parser',
  },
  plugins: ['vue', '@typescript-eslint', 'import', 'simple-import-sort'],
  rules: {
    indent: ['warn', 2],
    'linebreak-style': ['error', 'unix'],
    quotes: ['warn', 'single'],
    semi: ['warn', 'always'],
    'vue/no-v-model-argument': ['off'],
    'vue/no-multiple-template-root': ['off'],
    'vue/multi-word-component-names': ['off'],
    '@typescript-eslint/no-explicit-any': ['off'],
    'simple-import-sort/imports': 'error',
    'simple-import-sort/exports': 'error',
    '@typescript-eslint/no-unused-vars': [
      'warn',
      {
        argsIgnorePattern: '^_',
      },
    ],
  },
  settings: {
    'import/parsers': {
      '@typescript-eslint/parser': ['.ts', '.tsx'],
    },
    'import/resolver': {
      typescript: {
        alwaysTryTypes: true, // always try to resolve types under `<root>@types` directory even it doesn't contain any source code, like `@types/unist`
        project: ['tsconfig.json', 'tsconfig.node.json'],
      },
    },
  },
};
```

## File: frontend/.gitignore
```
# Logs
logs
*.log
npm-debug.log*
yarn-debug.log*
yarn-error.log*
pnpm-debug.log*
lerna-debug.log*

node_modules
dist
dist-ssr
*.local

# Editor directories and files
.idea
.DS_Store
*.suo
*.ntvs*
*.njsproj
*.sln
*.sw?

.eslintcache
```

## File: frontend/.prettierrc.cjs
```javascript
module.exports = {
	// 一行最多多少个字符
	printWidth: 150,
	// 指定每个缩进级别的空格数
	tabWidth: 2,
	// 使用制表符而不是空格缩进行
	useTabs: false,
	// 在语句末尾是否需要分号
	semi: true,
	// 是否使用单引号
	singleQuote: true,
	// 更改引用对象属性的时间 可选值"<as-needed|consistent|preserve>"
	quoteProps: 'as-needed',
	// 在JSX中使用单引号而不是双引号
	jsxSingleQuote: false,
	// 多行时尽可能打印尾随逗号。（例如，单行数组永远不会出现逗号结尾。） 可选值"<none|es5|all>"，默认none
	trailingComma: 'es5',
	// 在对象文字中的括号之间打印空格
	bracketSpacing: true,
	// jsx 标签的反尖括号需要换行
	jsxBracketSameLine: false,
	// 在单独的箭头函数参数周围包括括号 always：(x) => x \ avoid：x => x
	arrowParens: 'always',
	// 这两个选项可用于格式化以给定字符偏移量（分别包括和不包括）开始和结束的代码
	rangeStart: 0,
	rangeEnd: Infinity,
	// 指定要使用的解析器，不需要写文件开头的 @prettier
	requirePragma: false,
	// 不需要自动在文件开头插入 @prettier
	insertPragma: false,
	// 使用默认的折行标准 always\never\preserve
	proseWrap: 'preserve',
	// 指定HTML文件的全局空格敏感度 css\strict\ignore
	htmlWhitespaceSensitivity: 'css',
	// Vue文件脚本和样式标签缩进
	vueIndentScriptAndStyle: false,
	// 换行符使用 lf 结尾是 可选值"<auto|lf|crlf|cr>"
	endOfLine: 'lf',
};
```

## File: frontend/.vscode/extensions.json
```json
{
  "recommendations": ["Vue.volar", "Vue.vscode-typescript-vue-plugin"]
}
```

## File: frontend/.vscode/launch.json
```json
{
    // 使用 IntelliSense 了解相关属性。 
    // 悬停以查看现有属性的描述。
    // 欲了解更多信息，请访问: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "type": "chrome",
            "request": "launch",
            "name": "针对 localhost 启动 Chrome",
            "url": "http://localhost:5173",
            "webRoot": "${workspaceFolder}"
        }
    ]
}
```

## File: frontend/.vscode/settings.json
```json
{
    "i18n-ally.localesPaths": [
        "src/locales"
    ],
    "i18n-ally.sourceLanguage": "zh-CN",
    "i18n-ally.displayLanguage": "zh-CN",
    "i18n-ally.keystyle": "nested",
    "i18n-ally.namespace": true,
    "i18n-ally.pathMatcher": "{locale}.json",
    "[html]": {
        "editor.tabSize": 2
    },
    "[javascript]": {
        "editor.tabSize": 2
    },
    "[typescript]": {
        "editor.tabSize": 2
    },
    "[vue]": {
        "editor.tabSize": 2
    },
    "prettier.configPath": ".prettierrc.cjs",
    "editor.formatOnSave": false,
    "editor.codeActionsOnSave": {
        "source.fixAll.eslint": true
    },
    "files.refactoring.autoSave": false
}
```

## File: frontend/components.d.ts
```typescript
/* eslint-disable */
/* prettier-ignore */
// @ts-nocheck
// Generated by unplugin-vue-components
// Read more: https://github.com/vuejs/core/pull/3399
import '@vue/runtime-core'

export {}

declare module '@vue/runtime-core' {
  export interface GlobalComponents {
    NAutoComplete: typeof import('naive-ui')['NAutoComplete']
    NAvatar: typeof import('naive-ui')['NAvatar']
    NButton: typeof import('naive-ui')['NButton']
    NCard: typeof import('naive-ui')['NCard']
    NCollapse: typeof import('naive-ui')['NCollapse']
    NCollapseItem: typeof import('naive-ui')['NCollapseItem']
    NConfigProvider: typeof import('naive-ui')['NConfigProvider']
    NDataTable: typeof import('naive-ui')['NDataTable']
    NDivider: typeof import('naive-ui')['NDivider']
    NDropdown: typeof import('naive-ui')['NDropdown']
    NDynamicTags: typeof import('naive-ui')['NDynamicTags']
    NEmpty: typeof import('naive-ui')['NEmpty']
    NForm: typeof import('naive-ui')['NForm']
    NFormItem: typeof import('naive-ui')['NFormItem']
    NGlobalStyle: typeof import('naive-ui')['NGlobalStyle']
    NIcon: typeof import('naive-ui')['NIcon']
    NInput: typeof import('naive-ui')['NInput']
    NInputNumber: typeof import('naive-ui')['NInputNumber']
    NLayout: typeof import('naive-ui')['NLayout']
    NLayoutSider: typeof import('naive-ui')['NLayoutSider']
    NList: typeof import('naive-ui')['NList']
    NListItem: typeof import('naive-ui')['NListItem']
    NLog: typeof import('naive-ui')['NLog']
    NMenu: typeof import('naive-ui')['NMenu']
    NPageHeader: typeof import('naive-ui')['NPageHeader']
    NScrollbar: typeof import('naive-ui')['NScrollbar']
    NSelect: typeof import('naive-ui')['NSelect']
    NSpace: typeof import('naive-ui')['NSpace']
    NSpin: typeof import('naive-ui')['NSpin']
    NStatistic: typeof import('naive-ui')['NStatistic']
    NSwitch: typeof import('naive-ui')['NSwitch']
    NTab: typeof import('naive-ui')['NTab']
    NTabs: typeof import('naive-ui')['NTabs']
    NTag: typeof import('naive-ui')['NTag']
    NText: typeof import('naive-ui')['NText']
    NTooltip: typeof import('naive-ui')['NTooltip']
    PageHeader: typeof import('./src/components/PageHeader.vue')['default']
    PreferenceForm: typeof import('./src/components/PreferenceForm.vue')['default']
    RouterLink: typeof import('vue-router')['RouterLink']
    RouterView: typeof import('vue-router')['RouterView']
    UserProfileCard: typeof import('./src/components/UserProfileCard.vue')['default']
  }
}
```

## File: frontend/config/vite.config.base.ts
```typescript
import { fileURLToPath, URL } from 'node:url';

import presetUno from '@unocss/preset-uno';
import UnoCSS from '@unocss/vite';
import vue from '@vitejs/plugin-vue';
import { join } from 'path';
import { transformerDirectives } from 'unocss';
import { NaiveUiResolver } from 'unplugin-vue-components/resolvers';
import Components from 'unplugin-vue-components/vite';
import { defineConfig } from 'vite';

// https://vitejs.dev/config/
export default defineConfig({
  base: process.env.VITE_BASE || '/',
  plugins: [
    vue(),
    UnoCSS({
      presets: [
        /* no presets by default */
        presetUno(),
      ],
      /* options */
      transformers: [transformerDirectives()],
    }),
    Components({
      resolvers: [NaiveUiResolver()],
    }),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('../src', import.meta.url)),
    },
  },
  define: {
    'import.meta.env.PACKAGE_VERSION': JSON.stringify(process.env.npm_package_version),
    'import.meta.env.VITE_ENABLE_SENTRY': process.env.VITE_ENABLE_SENTRY || '\'no\'',
    'import.meta.env.VITE_ROUTER_BASE': process.env.VITE_ROUTER_BASE || '\'/\'',
    'import.meta.env.VITE_API_BASE_URL': process.env.VITE_API_BASE_URL || '\'/api/\''
  },
});
```

## File: frontend/config/vite.config.dev.ts
```typescript
import { mergeConfig } from 'vite';

import baseConfig from './vite.config.base';

export default mergeConfig(
  {
    mode: 'development',
    server: {
      host: '0.0.0.0',
      port: 5173,
      fs: {
        strict: true,
      },
      proxy: {
        '/api': {
          target: 'http://127.0.0.1:8000',
          changeOrigin: true,
          ws: true,
          rewrite: (path) => path.replace(/^\/api/, ''),
        },
      },
    },
  },
  baseConfig
);
```

## File: frontend/config/vite.config.prod.ts
```typescript
import { mergeConfig } from 'vite';

import baseConfig from './vite.config.base';

export default mergeConfig(
  {
    mode: 'production',
    plugins: [],
    build: {
      sourcemap: true,
      rollupOptions: {
        output: {
          manualChunks: {
            naive_ui: ['naive-ui'],
            vue: ['vue', 'vue-router', 'pinia', 'vue-i18n'],
          },
        },
      },
      chunkSizeWarningLimit: 2000,
    },
  },
  baseConfig
);
```

## File: frontend/index.html
```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <link rel="icon" type="image/svg+xml" href="/chatgpt-icon.svg" />
    <link rel="code-repository" href="https://github.com/moeakwak/chatgpt-web-share" />
    <title>ChatGPT Web Share</title>
  </head>
  <body>
    <div id="app"></div>
    <script type="module" src="/src/main.ts"></script>
  </body>
</html>
```

## File: frontend/package.json
```json
{
  "name": "chatgpt-share-frontend",
  "private": true,
  "version": "0.3.15",
  "type": "module",
  "scripts": {
    "dev": "vite --config ./config/vite.config.dev.ts",
    "build": "vue-tsc && vite build --config ./config/vite.config.prod.ts",
    "preview": "pnpm run build && vite preview --host",
    "eslint": "eslint src/ --ext .ts,.tsx,.js,.jsx,.vue --fix --cache"
  },
  "dependencies": {
    "@sentry/tracing": "^7.44.2",
    "@sentry/vue": "^7.44.2",
    "@traptitech/markdown-it-katex": "^3.6.0",
    "@vueuse/core": "^9.13.0",
    "axios": "^1.3.4",
    "clipboard-polyfill": "^4.0.0",
    "echarts": "^5.4.2",
    "file-saver": "^2.0.5",
    "highlight.js": "^11.7.0",
    "katex": "^0.16.4",
    "markdown-it": "^13.0.1",
    "markdown-it-highlightjs": "^4.0.1",
    "pinia": "^2.0.32",
    "vue": "^3.2.45",
    "vue-echarts": "^6.5.4",
    "vue-i18n": "^9.2.2",
    "vue-router": "4"
  },
  "devDependencies": {
    "@types/file-saver": "^2.0.5",
    "@types/markdown-it": "^12.2.3",
    "@types/node": "^18.14.0",
    "@typescript-eslint/eslint-plugin": "^5.59.0",
    "@typescript-eslint/parser": "^5.59.0",
    "@unocss/preset-uno": "^0.50.0",
    "@unocss/transformer-directives": "^0.50.1",
    "@unocss/vite": "^0.50.0",
    "@vicons/ionicons4": "^0.12.0",
    "@vicons/ionicons5": "^0.12.0",
    "@vicons/material": "^0.12.0",
    "@vitejs/plugin-vue": "^4.0.0",
    "eslint": "^8.38.0",
    "eslint-config-prettier": "^8.8.0",
    "eslint-import-resolver-typescript": "^3.5.5",
    "eslint-plugin-import": "^2.27.5",
    "eslint-plugin-prettier": "^4.2.1",
    "eslint-plugin-simple-import-sort": "^10.0.0",
    "eslint-plugin-vue": "^9.11.0",
    "naive-ui": "^2.34.3",
    "prettier": "^2.8.7",
    "typescript": "^4.9.3",
    "unocss": "^0.50.0",
    "unplugin-vue-components": "^0.24.0",
    "vfonts": "^0.0.3",
    "vite": "^4.1.0",
    "vite-plugin-eslint": "^1.8.1",
    "vue-tsc": "^1.0.24"
  }
}
```

## File: frontend/public/chatgpt-icon-black.svg
```xml
<svg xmlns="http://www.w3.org/2000/svg" shape-rendering="geometricPrecision" text-rendering="geometricPrecision" image-rendering="optimizeQuality" fill-rule="evenodd" clip-rule="evenodd" viewBox="0 0 512 512"><rect fill="#000000" width="512" height="512"/><path fill="#fff" fill-rule="nonzero" d="M378.68 230.011a71.432 71.432 0 003.654-22.541 71.383 71.383 0 00-9.783-36.064c-12.871-22.404-36.747-36.236-62.587-36.236a72.31 72.31 0 00-15.145 1.604 71.362 71.362 0 00-53.37-23.991h-.453l-.17.001c-31.297 0-59.052 20.195-68.673 49.967a71.372 71.372 0 00-47.709 34.618 72.224 72.224 0 00-9.755 36.226 72.204 72.204 0 0018.628 48.395 71.395 71.395 0 00-3.655 22.541 71.388 71.388 0 009.783 36.064 72.187 72.187 0 0077.728 34.631 71.375 71.375 0 0053.374 23.992H271l.184-.001c31.314 0 59.06-20.196 68.681-49.995a71.384 71.384 0 0047.71-34.619 72.107 72.107 0 009.736-36.194 72.201 72.201 0 00-18.628-48.394l-.003-.004zM271.018 380.492h-.074a53.576 53.576 0 01-34.287-12.423 44.928 44.928 0 001.694-.96l57.032-32.943a9.278 9.278 0 004.688-8.06v-80.459l24.106 13.919a.859.859 0 01.469.661v66.586c-.033 29.604-24.022 53.619-53.628 53.679zm-115.329-49.257a53.563 53.563 0 01-7.196-26.798c0-3.069.268-6.146.79-9.17.424.254 1.164.706 1.695 1.011l57.032 32.943a9.289 9.289 0 009.37-.002l69.63-40.205v27.839l.001.048a.864.864 0 01-.345.691l-57.654 33.288a53.791 53.791 0 01-26.817 7.17 53.746 53.746 0 01-46.506-26.818v.003zm-15.004-124.506a53.5 53.5 0 0127.941-23.534c0 .491-.028 1.361-.028 1.965v65.887l-.001.054a9.27 9.27 0 004.681 8.053l69.63 40.199-24.105 13.919a.864.864 0 01-.813.074l-57.66-33.316a53.746 53.746 0 01-26.805-46.5 53.787 53.787 0 017.163-26.798l-.003-.003zm198.055 46.089l-69.63-40.204 24.106-13.914a.863.863 0 01.813-.074l57.659 33.288a53.71 53.71 0 0126.835 46.491c0 22.489-14.033 42.612-35.133 50.379v-67.857c.003-.025.003-.051.003-.076a9.265 9.265 0 00-4.653-8.033zm23.993-36.111a81.919 81.919 0 00-1.694-1.01l-57.032-32.944a9.31 9.31 0 00-4.684-1.266 9.31 9.31 0 00-4.684 1.266l-69.631 40.205v-27.839l-.001-.048c0-.272.129-.528.346-.691l57.654-33.26a53.696 53.696 0 0126.816-7.177c29.644 0 53.684 24.04 53.684 53.684a53.91 53.91 0 01-.774 9.077v.003zm-150.831 49.618l-24.111-13.919a.859.859 0 01-.469-.661v-66.587c.013-29.628 24.053-53.648 53.684-53.648a53.719 53.719 0 0134.349 12.426c-.434.237-1.191.655-1.694.96l-57.032 32.943a9.272 9.272 0 00-4.687 8.057v.053l-.04 80.376zm13.095-28.233l31.012-17.912 31.012 17.9v35.812l-31.012 17.901-31.012-17.901v-35.8z"/></svg>
```

## File: frontend/public/chatgpt-icon.svg
```xml
<svg xmlns="http://www.w3.org/2000/svg" shape-rendering="geometricPrecision" text-rendering="geometricPrecision" image-rendering="optimizeQuality" fill-rule="evenodd" clip-rule="evenodd" viewBox="0 0 512 512"><rect fill="#10A37F" width="512" height="512"/><path fill="#fff" fill-rule="nonzero" d="M378.68 230.011a71.432 71.432 0 003.654-22.541 71.383 71.383 0 00-9.783-36.064c-12.871-22.404-36.747-36.236-62.587-36.236a72.31 72.31 0 00-15.145 1.604 71.362 71.362 0 00-53.37-23.991h-.453l-.17.001c-31.297 0-59.052 20.195-68.673 49.967a71.372 71.372 0 00-47.709 34.618 72.224 72.224 0 00-9.755 36.226 72.204 72.204 0 0018.628 48.395 71.395 71.395 0 00-3.655 22.541 71.388 71.388 0 009.783 36.064 72.187 72.187 0 0077.728 34.631 71.375 71.375 0 0053.374 23.992H271l.184-.001c31.314 0 59.06-20.196 68.681-49.995a71.384 71.384 0 0047.71-34.619 72.107 72.107 0 009.736-36.194 72.201 72.201 0 00-18.628-48.394l-.003-.004zM271.018 380.492h-.074a53.576 53.576 0 01-34.287-12.423 44.928 44.928 0 001.694-.96l57.032-32.943a9.278 9.278 0 004.688-8.06v-80.459l24.106 13.919a.859.859 0 01.469.661v66.586c-.033 29.604-24.022 53.619-53.628 53.679zm-115.329-49.257a53.563 53.563 0 01-7.196-26.798c0-3.069.268-6.146.79-9.17.424.254 1.164.706 1.695 1.011l57.032 32.943a9.289 9.289 0 009.37-.002l69.63-40.205v27.839l.001.048a.864.864 0 01-.345.691l-57.654 33.288a53.791 53.791 0 01-26.817 7.17 53.746 53.746 0 01-46.506-26.818v.003zm-15.004-124.506a53.5 53.5 0 0127.941-23.534c0 .491-.028 1.361-.028 1.965v65.887l-.001.054a9.27 9.27 0 004.681 8.053l69.63 40.199-24.105 13.919a.864.864 0 01-.813.074l-57.66-33.316a53.746 53.746 0 01-26.805-46.5 53.787 53.787 0 017.163-26.798l-.003-.003zm198.055 46.089l-69.63-40.204 24.106-13.914a.863.863 0 01.813-.074l57.659 33.288a53.71 53.71 0 0126.835 46.491c0 22.489-14.033 42.612-35.133 50.379v-67.857c.003-.025.003-.051.003-.076a9.265 9.265 0 00-4.653-8.033zm23.993-36.111a81.919 81.919 0 00-1.694-1.01l-57.032-32.944a9.31 9.31 0 00-4.684-1.266 9.31 9.31 0 00-4.684 1.266l-69.631 40.205v-27.839l-.001-.048c0-.272.129-.528.346-.691l57.654-33.26a53.696 53.696 0 0126.816-7.177c29.644 0 53.684 24.04 53.684 53.684a53.91 53.91 0 01-.774 9.077v.003zm-150.831 49.618l-24.111-13.919a.859.859 0 01-.469-.661v-66.587c.013-29.628 24.053-53.648 53.684-53.648a53.719 53.719 0 0134.349 12.426c-.434.237-1.191.655-1.694.96l-57.032 32.943a9.272 9.272 0 00-4.687 8.057v.053l-.04 80.376zm13.095-28.233l31.012-17.912 31.012 17.9v35.812l-31.012 17.901-31.012-17.901v-35.8z"/></svg>
```

## File: frontend/README.md
```markdown
# Vue 3 + TypeScript + Vite

This template should help get you started developing with Vue 3 and TypeScript in Vite. The template uses Vue 3 `<script setup>` SFCs, check out the [script setup docs](https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup) to learn more.

## Recommended IDE Setup

- [VS Code](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur) + [TypeScript Vue Plugin (Volar)](https://marketplace.visualstudio.com/items?itemName=Vue.vscode-typescript-vue-plugin).

## Type Support For `.vue` Imports in TS

TypeScript cannot handle type information for `.vue` imports by default, so we replace the `tsc` CLI with `vue-tsc` for type checking. In editors, we need [TypeScript Vue Plugin (Volar)](https://marketplace.visualstudio.com/items?itemName=Vue.vscode-typescript-vue-plugin) to make the TypeScript language service aware of `.vue` types.

If the standalone TypeScript plugin doesn't feel fast enough to you, Volar has also implemented a [Take Over Mode](https://github.com/johnsoncodehk/volar/discussions/471#discussioncomment-1361669) that is more performant. You can enable it by the following steps:

1. Disable the built-in TypeScript Extension
   1. Run `Extensions: Show Built-in Extensions` from VSCode's command palette
   2. Find `TypeScript and JavaScript Language Features`, right click and select `Disable (Workspace)`
2. Reload the VSCode window by running `Developer: Reload Window` from the command palette.
```

## File: frontend/src/api/chat.ts
```typescript
import axios from 'axios';

import { ConversationSchema } from '@/types/schema';

import ApiUrl from './url';

export function getAllConversationsApi(fetch_all = false) {
  return axios.get<Array<ConversationSchema>>(ApiUrl.Conversation, {
    params: { fetch_all },
  });
}

export function getConversationHistoryApi(conversation_id: string) {
  return axios.get<any>(ApiUrl.Conversation + '/' + conversation_id);
}

export function deleteConversationApi(conversation_id: string) {
  return axios.delete(ApiUrl.Conversation + '/' + conversation_id);
}

export function clearAllConversationApi() {
  return axios.delete(ApiUrl.Conversation);
}

export function vanishConversationApi(conversation_id: string) {
  return axios.delete(ApiUrl.Conversation + '/' + conversation_id + '/vanish');
}

export function setConversationTitleApi(conversation_id: string, title: string) {
  return axios.patch<ConversationSchema>(ApiUrl.Conversation + '/' + conversation_id, null, {
    params: { title },
  });
}

export function generateConversationTitleApi(conversation_id: string, message_id: string) {
  return axios.patch<ConversationSchema>(ApiUrl.Conversation + '/' + conversation_id + '/gen_title', null, {
    params: { message_id },
  });
}

export type AskInfo = {
  message: string;
  new_title?: string;
  conversation_id?: string;
  parent_id?: string;
  model_name?: string;
  is_public?: boolean;
  timeout?: number;
};

export function getAskWebsocketApiUrl() {
  const protocol = window.location.protocol === 'https:' ? 'wss' : 'ws';
  const url = `${protocol}://${window.location.host}${import.meta.env.VITE_API_BASE_URL}conv`;
  console.log('getAskWebsocketApiUrl', url);
  return url;
}

export function assignConversationToUserApi(conversation_id: string, username: string) {
  return axios.patch(`${ApiUrl.Conversation}/${conversation_id}/assign/${username}`);
}
```

## File: frontend/src/api/interceptor.ts
```typescript
import type { AxiosResponse,InternalAxiosRequestConfig } from 'axios';
import axios from 'axios';

import { i18n } from '@/i18n';
import router from '@/router';
import { useUserStore } from '@/store';
import { Dialog,Message } from '@/utils/tips';

// import { isLogin } from '@/utils/auth';
import ApiUrl from './url';
const t = i18n.global.t as any;

export interface HttpResponse<T = unknown> {
  code: number;
  message: string;
  result: T;
}

axios.defaults.baseURL = import.meta.env.VITE_API_BASE_URL;
// axios.defaults.baseURL = "/api/";

axios.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    // if (token) {
    //   if (!config.headers) {
    //     config.headers = {};
    //   }
    //   config.headers.Authorization = `Bearer ${token}`;
    // }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

/**
 * 添加响应拦截器
 * 这里将 { code, message, result } 解构出来，response.data 替换成 result
 */
const successCode = [200, 201, 204];
axios.interceptors.response.use(
  (response: AxiosResponse<HttpResponse>) => {
    if ([201, 204].includes(response.status)) {
      return response;
    }
    const res = response.data;
    if (!successCode.includes(res.code)) {
      console.warn('Error: ', res);
      let msg = `${res.code}`;
      if (res.message) {
        msg += ` ${t(res.message)}`;
      }
      if (res.result) {
        msg += `: ${t(res.result)}`;
      }
      Message.error(msg, {
        duration: 5 * 1000,
      });
      if ([401].includes(res.code) && !([ApiUrl.Login, ApiUrl.Logout] as Array<string>).includes(response.config.url || '')) {
        Dialog.error({
          title: t('errors.loginExpired') as string,
          content: t('tips.loginExpired'),
          positiveText: t('commons.confirm'),
          negativeText: t('commons.stayInCurrentPage'),
          onPositiveClick() {
            const userStore = useUserStore();
            userStore.logout().then(() => {
              router.push({ name: 'login' });
            });
            window.location.reload();
          },
        });
      }
      return Promise.reject(res);
    }
    (response.data as any) = res.result;
    return response;
  },
  (error) => {
    Message.error((error.msg && t(error.msg)) || 'Request Error', {
      duration: 5 * 1000,
    });
    console.error('Request Error', error);
    return Promise.reject(error);
  }
);
```

## File: frontend/src/api/status.ts
```typescript
import axios from 'axios';

import { ServerStatusSchema } from '@/types/schema';

import ApiUrl from './url';

export function getServerStatusApi() {
  return axios.get<ServerStatusSchema>(ApiUrl.ServerStatus);
}
```

## File: frontend/src/api/system.ts
```typescript
import axios from 'axios';

import { LogFilterOptions } from '@/types/schema';

import ApiUrl from './url';

export function getSystemInfoApi() {
  return axios.get(ApiUrl.SystemInfo);
}

export function getRequestStatisticsApi() {
  return axios.get(ApiUrl.SystemRequestStatistics);
}

export function getServerLogsApi(options: LogFilterOptions | null) {
  return axios.post(ApiUrl.ServerLogs, options);
}

export function getProxyLogsApi(options: LogFilterOptions | null) {
  return axios.post(ApiUrl.ProxyLogs, options);
}
```

## File: frontend/src/api/url.ts
```typescript
enum ApiUrl {
  Register = '/auth/register',
  Login = '/auth/login',
  Logout = '/auth/logout',
  UserInfo = '/user/me',

  Conversation = '/conv',
  UserList = '/user',

  ServerStatus = '/status',
  SystemInfo = '/system/info',
  SystemRequestStatistics = '/system/request_statistics',
  ProxyLogs = '/system/proxy_logs',
  ServerLogs = '/system/server_logs',
}

export default ApiUrl;
```

## File: frontend/src/api/user.ts
```typescript
import axios from 'axios';

import { LimitSchema, UserCreate, UserRead } from '@/types/schema';

import ApiUrl from './url';

export interface LoginData {
  username: string;
  password: string;
}

export function loginApi(data: LoginData) {
  const formData = new FormData();
  formData.set('username', data.username);
  formData.set('password', data.password);
  return axios.post<any>(ApiUrl.Login, formData, {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });
}

export function registerApi(userInfo: UserCreate) {
  return axios.post<UserRead>(ApiUrl.Register, userInfo);
}

export function logoutApi() {
  return axios.post<any>(ApiUrl.Logout);
}

export function getUserInfoApi() {
  return axios.get<UserRead>(ApiUrl.UserInfo);
}

export function getAllUserApi() {
  return axios.get<UserRead[]>(ApiUrl.UserList);
}

export function deleteUserApi(user_id: number) {
  return axios.delete(ApiUrl.UserList + `/${user_id}`);
}

export function resetUserPasswordApi(user_id: number, new_password: string) {
  return axios.patch(ApiUrl.UserList + `/${user_id}/reset-password`, null, {
    params: { new_password },
  });
}

export function updateUserLimitApi(user_id: number, limit: LimitSchema) {
  return axios.post(ApiUrl.UserList + `/${user_id}/limit`, limit);
}

// export function updateUserInfoApi(userInfo: Partial<UserUpdate>) {
//   return axios.patch<UserRead>(ApiUrl.UserInfo, userInfo);
// }
```

## File: frontend/src/App.vue
```vue
<template>
  <n-config-provider :theme="theme">
    <n-global-style />
    <div class="w-full box-border min-h-screen flex flex-col">
      <div class="my-4 px-4">
        <PageHeader />
      </div>
      <router-view />
    </div>
  </n-config-provider>
</template>

<script setup lang="ts">
import { darkTheme } from 'naive-ui';
import { computed } from 'vue';

import PageHeader from './components/PageHeader.vue';
import { useAppStore } from './store';

const appStore = useAppStore();

const theme = computed(() => {
  if (appStore.theme == 'dark') {
    return darkTheme;
  } else {
    return {};
  }
});
</script>
```

## File: frontend/src/components/PageHeader.vue
```vue
<template>
  <n-page-header>
    <template #title>
      <n-space :align="'center'">
        <div>
          <a
            href="/"
            style="text-decoration: none; color: inherit"
          >{{ $t('commons.siteTitle') }}</a>
        </div>
        <div class="hidden sm:block">
          <a
            class="h-full inline-block flex"
            href="https://github.com/moeakwak/chatgpt-web-share"
            target="_blank"
          >
            <n-icon
              :color="appStore.theme == 'dark' ? 'white' : 'black'"
              :component="LogoGithub"
            />
          </a>
        </div>
        <n-tag
          :bordered="false"
          type="success"
          size="small"
          class="hidden sm:inline-flex"
        >
          {{ version }}
        </n-tag>
      </n-space>
    </template>
    <template #avatar>
      <n-avatar :src="chatgptIcon" />
    </template>
    <template #extra>
      <n-space>
        <div class="space-x-2">
          <div
            v-if="userStore.user"
            class="inline-block"
          >
            <span class="hidden sm:inline mr-1">Hi, {{ userStore.user.nickname }}</span>
            <n-dropdown
              :options="getOptions()"
              placement="bottom-start"
            >
              <n-button
                circle
                class="ml-2"
              >
                <n-icon :component="SettingsSharp" />
              </n-button>
            </n-dropdown>
          </div>
          <div
            v-else
            class="text-gray-500 inline-block"
          >
            {{ $t('commons.notLogin') }}
          </div>
          <n-button
            v-if="userStore.user?.is_superuser"
            circle
            @click="jumpToAdminOrConv"
          >
            <n-icon :component="isInAdmin ? ChatFilled : ManageAccountsFilled" />
          </n-button>
          <n-button
            circle
            @click="toggleTheme"
          >
            <n-icon :component="themeIcon" />
          </n-button>
          <n-dropdown
            :options="languageOptions"
            placement="bottom-start"
          >
            <n-button circle>
              <n-icon :component="Language" />
            </n-button>
          </n-dropdown>
        </div>
      </n-space>
    </template>
  </n-page-header>
</template>

<script setup lang="ts">
import { Language,LogoGithub, SettingsSharp } from '@vicons/ionicons5';
import { ChatFilled,DarkModeRound, LightModeRound, ManageAccountsFilled } from '@vicons/material';
import { DropdownOption } from 'naive-ui';
import { computed, h, ref } from 'vue';
import { useI18n } from 'vue-i18n';
import { useRoute } from 'vue-router';

// eslint-disable-next-line import/no-unresolved
import chatgptIcon from '/chatgpt-icon.svg';
import { resetUserPasswordApi } from '@/api/user';
import router from '@/router';
import { useAppStore,useUserStore } from '@/store';
import { Preference } from '@/types/custom';
import { popupResetUserPasswordDialog } from '@/utils/renders';
import { Dialog, Message } from '@/utils/tips';

import PreferenceForm from './PreferenceForm.vue';
import UserProfileCard from './UserProfileCard.vue';

const { t } = useI18n();
const userStore = useUserStore();
const appStore = useAppStore();
const route = useRoute();
const version = 'v' + import.meta.env.PACKAGE_VERSION;

console.log(route);

const isInAdmin = computed(() => {
  return route.path.startsWith('/admin');
});

const themeIcon = computed(() => {
  if (appStore.theme == 'dark') {
    return DarkModeRound;
  } else {
    return LightModeRound;
  }
});

const toggleTheme = () => {
  appStore.toggleTheme();
};

const languageOptions = [
  {
    label: '简体中文',
    key: 'zh-CN',
    props: {
      onClick: () => {
        appStore.setLanguage('zh-CN');
      },
    },
  },
  {
    label: 'English',
    key: 'en-US',
    props: {
      onClick: () => {
        appStore.setLanguage('en-US');
      },
    },
  },
];

const getOptions = (): Array<DropdownOption> => {
  const options: Array<DropdownOption> = [
    {
      label: t('commons.userProfile'),
      key: 'profile',
      props: {
        onClick: () =>
          Dialog.info({
            title: t('commons.userProfile'),
            content: () => h(UserProfileCard, {}, {}),
            positiveText: t('commons.confirm'),
          }),
      },
    },
    {
      label: t('commons.resetPassword'),
      key: 'resetpwd',
      props: {
        onClick: resetPassword,
      },
    },
    {
      label: t('commons.preferences'),
      key: 'preference',
      props: {
        onClick: () => {
          let preference: Preference = {
            ...appStore.preference,
          };
          Dialog.info({
            title: t('commons.preferences'),
            positiveText: t('commons.confirm'),
            negativeText: t('commons.cancel'),
            content: () =>
              h(PreferenceForm, {
                onUpdate: (val: any) => {
                  preference = val;
                },
                value: preference,
              }),
            onPositiveClick() {
              appStore.$patch({
                preference: preference,
              });
              Message.success(t('tips.success'));
            },
          });
        },
      },
    },
    {
      label: t('commons.logout'),
      key: 'logout',
      props: {
        onClick: () =>
          Dialog.info({
            title: t('commons.logout'),
            content: t('tips.logoutConfirm'),
            positiveText: t('commons.confirm'),
            negativeText: t('commons.cancel'),
            onPositiveClick: async () => {
              await userStore.logout();
              Message.success(t('commons.logoutSuccess'));
              await router.push({ name: 'login' });
            },
          }),
      },
    },
  ];
  return options;
};

const resetPassword = () => {
  popupResetUserPasswordDialog(
    async (password: string) => {
      await resetUserPasswordApi(userStore.user!.id, password);
    },
    () => {
      Message.info(t('tips.resetUserPasswordSuccess'));
    },
    () => {
      Message.error(t('tips.resetUserPasswordFailed'));
    }
  );
};

const jumpToAdminOrConv = async () => {
  if (isInAdmin.value) {
    await router.push({ name: 'conversation' });
  } else {
    await router.push({ name: 'admin' });
  }
};
</script>
```

## File: frontend/src/components/PreferenceForm.vue
```vue
<template>
  <!-- A n-form: a n-select to switch sendKey in ["Shift+Enter", "Enter", "Ctrl+Enter"] -->
  <n-form
    v-model:value="model"
    label-placement="left"
    label-width="auto"
  >
    <n-form-item
      :label="t('commons.sendKey')"
      prop="sendKey"
    >
      <n-select
        v-model:value="model.sendKey"
        :options="sendKeyOptions"
      />
    </n-form-item>
    <!-- n-switch for renderUserMessageInMd and codeAutoWrap -->
    <n-form-item
      :label="t('commons.renderUserMessageInMd')"
      prop="renderUserMessageInMd"
    >
      <n-switch v-model:value="model.renderUserMessageInMd" />
    </n-form-item>
    <n-form-item
      :label="t('commons.codeAutoWrap')"
      prop="codeAutoWrap"
    >
      <n-switch v-model:value="model.codeAutoWrap" />
    </n-form-item>
    <n-form-item
      :label="t('commons.widerConversationPage')"
      prop="widerConversationPage"
    >
      <n-switch v-model:value="model.widerConversationPage" />
    </n-form-item>
  </n-form>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';

import { i18n } from '@/i18n';
import { Preference } from '@/types/custom';

const t = i18n.global.t as any;

const props = defineProps<{
  value: Preference;
}>();

const model = ref<Preference>(props.value);

const sendKeyOptions = [
  { label: 'Enter', value: 'Enter' },
  { label: 'Shift+Enter', value: 'Shift+Enter' },
  { label: 'Ctrl+Enter', value: 'Ctrl+Enter' },
];

const emit = defineEmits(['update:value']);

watch(
  () => model.value,
  () => {
    emit('update:value', model.value);
  }
);
</script>
```

## File: frontend/src/components/UserProfileCard.vue
```vue
<template>
  <n-card content-style="padding: 0;">
    <n-list
      v-for="(item, i) of items"
      :key="i"
      hoverable
      show-divider
    >
      <n-list-item>
        <div class="flex flex-row justify-between content-center">
          <div>{{ item.title }}</div>
          <div>{{ item.value }}</div>
        </div>
      </n-list-item>
    </n-list>
  </n-card>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';

import { i18n } from '@/i18n';
import { useUserStore } from '@/store';
import { chatStatusMap,ServerStatusSchema, UserRead } from '@/types/schema';
const t = i18n.global.t as any;

const serverStatus = ref<ServerStatusSchema>({});

const userStore = useUserStore();
const user: UserRead | null = userStore.user;

const propsToShow = [
  'id',
  'username',
  'email',
  'nickname',
  'is_superuser',
  'active_time',
  'chat_status',
  'can_use_paid',
  'can_use_gpt4',
  'max_conv_count',
  'available_ask_count',
  'available_gpt4_ask_count',
];

const translateKey = (key: string) => {
  if (['id', 'username', 'email'].includes(key)) {
    return key;
  }
  return t(`labels.${key}`);
};

const translateValue = (key: string, value: any) => {
  if (['is_superuser', 'can_use_paid', 'can_use_gpt4'].includes(key)) {
    return value ? t('commons.yes') : t('commons.no');
  } else if (key === 'active_time') {
    return value ? new Date(value + 'Z').toLocaleString() : t('commons.neverActive');
  } else if (key === 'chat_status') {
    return t(chatStatusMap[value as keyof typeof chatStatusMap]);
  } else if (key === 'max_conv_count') {
    return value === -1 ? t('commons.unlimited') : value;
  } else if (key === 'available_ask_count' || key === 'available_gpt4_ask_count') {
    return value === -1 ? t('commons.unlimited') : value;
  }
  return value;
};

const items = computed(() => {
  if (!user) return [];
  return propsToShow.map((prop) => {
    return {
      title: translateKey(prop),
      value: translateValue(prop, user[prop as keyof UserRead]),
    };
  });
});
</script>
```

## File: frontend/src/i18n.ts
```typescript
import { useStorage } from '@vueuse/core';
import { WritableComputedRef } from 'vue';
import { createI18n, type I18n, type Locale } from 'vue-i18n';

import EN from './locales/en-US.json';
import ZH from './locales/zh-CN.json';

let i18n: I18n;

const init = () => {
  i18n = createI18n({
    legacy: false,
    locale: useStorage('language', 'zh-CN').value,
    messages: {
      'en-US': {
        ...EN,
      },
      'zh-CN': {
        ...ZH,
      },
    },
  });
};

const setLocale = (locale: Locale): void => {
  (i18n.global.locale as WritableComputedRef<string>).value = locale;
};

init();

export { i18n, setLocale };
```

## File: frontend/src/locales/en-US.json
```json
{
  "commons": {
    "notLogin": "Not logged in",
    "siteTitle": "ChatGPT",
    "confirm": "Confirm",
    "stayInCurrentPage": "Stay on this page",
    "username": "Username",
    "password": "Password",
    "login": "Log in",
    "logout": "Log out",
    "logoutSuccess": "Logged out successfully",
    "newConversation": "New conversation",
    "confirmDialogTitle": "Confirm action",
    "cancel": "Cancel",
    "delete": "Delete",
    "rename": "Rename",
    "adminPanel": "Admin Panel",
    "userManagement": "User Management",
    "conversationManagement": "Conversation Management",
    "nickname": "Nickname",
    "email": "Email",
    "isSuperuser": "Superuser",
    "yes": "Yes",
    "no": "No",
    "actions": "Actions",
    "addUser": "Add user",
    "addUserSuccess": "User added successfully",
    "addUserFailed": "Failed to add user",
    "title": "Title",
    "createTime": "Creation time",
    "usePaidModel": "Model",
    "belongToUser": "Belongs to user",
    "isValid": "Valid",
    "isPublic": "Public",
    "deleteInvalidConversations": "Destroy invalid conversations",
    "deleteInvalidConversationsConfirm": "Are you sure you want to delete all invalid conversations from the database?",
    "deleteSuccess": "Deleted successfully",
    "deleteFailed": "Failed to delete",
    "deleteUser": "Delete user",
    "resetPassword": "Reset password",
    "neverActive": "Never logged in",
    "activeTime": "Last active time",
    "serverStatus": "Server status",
    "activeUserIn5m": "Active in last 5 minutes",
    "activeUserIn1h": "Active in last 1 hour",
    "activeUserIn1d": "Active in last 1 day",
    "isChatbotBusy": "Processing replies",
    "person": "Person",
    "chatbotWaitingCount": "Users Waiting for Reply",
    "deleteConversation": "Delete conversation",
    "invalidateConversation": "Mark as invalid",
    "vanishConversation": "Destroy conversation",
    "shaModel": "Default",
    "paidModel": "Legacy",
    "currentConversationModel": "Model currently used in this conversation",
    "canUsePaidModel": "Allow use of paid model",
    "maxConversationCount": "Max number of conversations",
    "availableAskCount": "Available Ask Count",
    "unlimited": "Unlimited",
    "setUserLimit": "Set user limit",
    "setUserLimitSuccess": "User limit set successfully",
    "vanishOrphanConversations": "Destroy all conversations without a user",
    "empty": "Empty",
    "chooseUser": "Choose user",
    "chooseUserToAssign": "Assign to user",
    "withdrawMessage": "Withdraw message",
    "askingChatStatus": "Asking",
    "queueingChatStatus": "Queuing",
    "idlingChatStatus": "Idle",
    "userProfile": "User Profile",
    "status": "Status",
    "copiedToClipboard": "Copied to clipboard",
    "send": "Send",
    "modelName": "Model Name",
    "gpt4Model": "GPT-4",
    "gpt4MobileModel": "GPT-4 (mobile v2)",
    "unknown": "未知",
    "availableGPT4AskCount": "Available GPT-4 Ask Count",
    "canUseGPT4Model": "Allow GPT-4",
    "logViewer": "Log Viewer",
    "serverLogs": "Server Logs",
    "proxyLogs": "Reverse Proxy Logs",
    "updateInterval": "Refresh Interval",
    "autoScrolling": "Auto Scroll",
    "maxLineCount": "Max Line Count",
    "excludeKeywords": "Exclude Keywords",
    "conversation": "Conversation",
    "enterConversation": "Enter Conversation",
    "systemManagement": "System Management",
    "statisticsInfo": "Statistics",
    "userCountAndOnlineCount": "Active / Total Users",
    "conversationCount": "Valid / Total Conversations",
    "serverOverview": "Overview",
    "chatbotStatus": "Reply Status",
    "startUpDuration": "Run Time",
    "totalRequestsCount": "Total Number of Requests",
    "requestUsers": "Request Users",
    "requestUsersCount": "Number of requesting users",
    "gpt4AskUsers": "Who used GPT-4",
    "normalAskUsers": "Who used normal",
    "normalAskCount": "Normal Model Ask Count",
    "gpt4AskCount": "GPT-4 Model Ask Count",
    "askRequestsCount": "Ask Count",
    "sumOfGpt4AskDuration": "Total time spent on GPT-4 dialogue",
    "sumOfNormalAskDuration": "The total time spent on ordinary conversations",
    "abortRequest": "Abort Request",
    "clearAllConversations": "Clear All Conversations",
    "readyToClearAllConversations": "Are you sure you want to clear all conversations?",
    "sendKey": "Send Short Key",
    "preferences": "Preference",
    "widerConversationPage": "Make conversation page wider"
  },
  "errors": {
    "403": "403 error: access denied",
    "404": "404 Not Found",
    "loginExpired": "Login expired",
    "userNotLogin": "User not logged in",
    "askError": "Failed to get reply",
    "conversationNotFound": "Conversation not found",
    "conversationAlreadyDeleted": "Conversation already deleted",
    "conversationTitleAlreadyGenerated": "Conversation title already generated",
    "unauthorized": "Unauthorized",
    "missingMessage": "Missing message content",
    "missingConversationId": "Missing conversation_id",
    "paidModelNotAvailable": "Paid model not available",
    "userNotAllowToUsePaidModel": "User not allowed to use paid model",
    "timeout": "Request timeout",
    "chatgptResponseError": "Error returned by ChatGPT",
    "unknownError": "Unknown error",
    "requestValidationError": "Request parameter error",
    "userAlreadyExists": "User already exists",
    "badCredentials": "Incorrect username or password",
    "maxConversationCountReached": "Max number of conversations reached",
    "cannotConnectMoreThanOneClient": "Current user has a conversation waiting for reply, cannot request again",
    "noAvailableAskCount": "No Available Ask Count!",
    "noAvailableGPT4AskCount": "No Available GPT-4 Ask Count!",
    "userNotAllowToUseGPT4Model": "Not allowed to use GPT-4 model",
    "noUserSelected": "No User Selected",
    "httpError": "HTTP Error",
    "internal": "Internal Error",
    "invalidParams": "Invalid Params"
  },
  "tips": {
    "loginExpired": "Login expired. Do you want to go to the login page?",
    "pleaseEnterUsername": "Please enter username",
    "pleaseEnterPassword": "Please enter password",
    "waiting": "Thinking...",
    "queueing": "Replying to other users, in queue...",
    "deleteConversation": "Are you sure you want to delete this conversation? You cannot access this conversation anymore.",
    "deleteConversationSuccess": "Deleted successfully",
    "deleteConversationFailed": "Failed to delete",
    "changeConversationTitleSuccess": "Title changed successfully",
    "changeConversationTitleFailed": "Failed to change title",
    "rename": "Please enter new title",
    "sendMessage": "Please enter message, press {0} to send",
    "logoutConfirm": "Are you sure you want to log out?",
    "pleaseEnterNickname": "Please enter nickname",
    "pleaseEnterEmail": "Please enter email",
    "deleteUserConfirm": "Are you sure you want to delete this user?",
    "deleteUserSuccess": "Deleted successfully",
    "deleteUserFailed": "Failed to delete",
    "resetPassword": "Please enter new password",
    "resetUserPasswordSuccess": "Password changed successfully",
    "resetUserPasswordFailed": "Failed to change password",
    "success": "Success",
    "failed": "Failed",
    "requestSuccess": "Request succeeded",
    "requestFailed": "Request failed",
    "invalidateConversation": "Are you sure you want to mark these conversations as invalid? It will no longer appear in the user list.",
    "vanishConversation": "Are you sure you want to destroy this conversation? It will be removed from the database and the records on the OpenAI account will be cleared.",
    "conversationTitle": "Enter conversation title, leave blank to use default title",
    "whetherUsePaidModel": "Please choose the model to use, leave blank to use default model",
    "loadConversation": "Select a conversation from the left to load, or",
    "newConversation": "start a new conversation",
    "loginSuccess": "Logged in successfully",
    "jumpingPage": "Jumping...",
    "terminated": "Connection terminated by the server",
    "loading": "Loading...",
    "pleaseSelectConversation": "Please select a conversation first",
    "pressEscToExitFullscreen": "Full Screen | Press Esc To Exit",
    "refreshed": "Refreshed Successfully",
    "autoScrolling": "auto scroll"
  },
  "labels": {
    "nickname": "Nickname",
    "is_superuser": "Is Admin",
    "active_time": "Last active time",
    "chat_status": "Current status",
    "can_use_paid": "Allow to use paid model",
    "can_use_gpt4": "Allow to use GPT-4 model",
    "max_conv_count": "Max number of conversations",
    "available_ask_count": "Available ask count",
    "available_gpt4_ask_count": "Available GPT-4 ask count"
  }
}
```

## File: frontend/src/locales/zh-CN.json
```json
{
  "commons": {
    "notLogin": "未登录",
    "siteTitle": "ChatGPT",
    "confirm": "确定",
    "stayInCurrentPage": "留在此页",
    "username": "用户名",
    "password": "密码",
    "login": "登录",
    "logout": "退出登录",
    "logoutSuccess": "登出成功",
    "newConversation": "新对话",
    "confirmDialogTitle": "确认操作",
    "cancel": "取消",
    "delete": "删除",
    "rename": "重命名",
    "adminPanel": "管理面板",
    "userManagement": "用户管理",
    "conversationManagement": "会话管理",
    "nickname": "昵称",
    "email": "邮箱",
    "isSuperuser": "管理员",
    "yes": "是",
    "no": "否",
    "actions": "操作",
    "addUser": "添加用户",
    "addUserSuccess": "添加用户成功",
    "addUserFailed": "添加用户失败",
    "title": "标题",
    "createTime": "创建时间",
    "modelName": "模型",
    "belongToUser": "所属用户",
    "isValid": "是否有效",
    "isPublic": "是否公开",
    "deleteInvalidConversations": "销毁无效会话",
    "deleteInvalidConversationsConfirm": "确定要从数据库销毁所有失效会话吗？",
    "deleteSuccess": "删除成功",
    "deleteFailed": "删除失败",
    "deleteUser": "删除用户",
    "resetPassword": "修改密码",
    "neverActive": "从未登录",
    "activeTime": "最后活动时间",
    "serverStatus": "服务状态",
    "activeUserIn5m": "5分钟内活跃",
    "activeUserIn1h": "1小时内活跃",
    "activeUserIn1d": "1天内活跃",
    "isChatbotBusy": "处理回复中",
    "person": "人",
    "chatbotWaitingCount": "回复排队人数",
    "deleteConversation": "删除会话",
    "invalidateConversation": "标记为无效",
    "vanishConversation": "彻底删除对话",
    "shaModel": "Default",
    "paidModel": "Legacy",
    "currentConversationModel": "当前对话使用的模型",
    "canUsePaidModel": "允许paid模型",
    "maxConversationCount": "最大对话数",
    "availableAskCount": "剩余提问次数",
    "unlimited": "不限",
    "setUserLimit": "设置用户限制",
    "setUserLimitSuccess": "设置用户限制成功",
    "vanishOrphanConversations": "销毁所有无用户归属的会话",
    "empty": "无",
    "chooseUser": "选择用户",
    "chooseUserToAssign": "分配给用户",
    "withdrawMessage": "撤回消息",
    "status": "状态",
    "askingChatStatus": "回复中",
    "queueingChatStatus": "排队中",
    "idlingChatStatus": "空闲",
    "userProfile": "用户信息",
    "copiedToClipboard": "已复制到剪贴板",
    "send": "发送",
    "gpt4Model": "GPT-4",
    "gpt4MobileModel": "GPT-4 (mobile v2)",
    "unknown": "未知",
    "availableGPT4AskCount": "剩余GPT-4对话次数",
    "canUseGPT4Model": "允许GPT-4",
    "logViewer": "日志记录",
    "serverLogs": "服务器日志",
    "proxyLogs": "反向代理日志",
    "updateInterval": "刷新间隔",
    "autoScrolling": "自动滚动",
    "maxLineCount": "行数",
    "usePaidModel": "使用paid模型",
    "excludeKeywords": "排除关键词",
    "conversation": "对话",
    "enterConversation": "进入对话",
    "systemManagement": "系统管理",
    "statisticsInfo": "统计信息",
    "userCountAndOnlineCount": "活跃用户 / 总用户",
    "conversationCount": "有效 / 总对话数量",
    "chatbotStatus": "回复状态",
    "serverOverview": "服务概况",
    "totalRequestsCount": "请求总数量",
    "startUpDuration": "运行时长",
    "requestUsers": "请求用户",
    "requestUsersCount": "请求用户数量",
    "normalAskCount": "普通模型回复次数",
    "normalAskUsers": "使用普通模型",
    "gpt4AskUsers": "使用GPT-4模型",
    "gpt4AskCount": "GPT-4模型回复次数",
    "askRequestsCount": "对话次数",
    "sumOfNormalAskDuration": "普通对话总耗时",
    "sumOfGpt4AskDuration": "GPT-4对话总耗时",
    "abortRequest": "停止请求",
    "clearAllConversations": "清空所有对话",
    "readyToClearAllConversations": "警告：你即将要清空所有对话，这将销毁所有对话数据。确定要清空所有对话吗？",
    "sendKey": "发送快捷键",
    "preferences": "偏好设置",
    "renderUserMessageInMd": "渲染用户回复为Markdown",
    "codeAutoWrap": "代码自动换行",
    "widerConversationPage": "对话界面宽屏展示"
  },
  "errors": {
    "403": "403 错误：无访问权限",
    "404": "404 错误：页面不存在",
    "loginExpired": "登录已过期",
    "userNotLogin": "用户未登录",
    "askError": "获取回复失败",
    "conversationNotFound": "会话不存在",
    "conversationAlreadyDeleted": "会话已被删除",
    "conversationTitleAlreadyGenerated": "会话标题已生成",
    "unauthorized": "未授权",
    "missingMessage": "缺少消息内容",
    "missingConversationId": "缺少conversation_id",
    "paidModelNotAvailable": "付费模型不可用",
    "userNotAllowToUsePaidModel": "用户不允许使用付费模型",
    "userNotAllowToUseGPT4Model": "用户不允许使用GPT-4模型",
    "noAvailableAskCount": "剩余提问次数不足",
    "noAvailableGPT4AskCount": "GPT-4模型剩余提问次数不足",
    "timeout": "请求超时",
    "chatgptResponseError": "ChatGPT返回错误",
    "unknownError": "未知错误",
    "requestValidationError": "请求参数错误",
    "userAlreadyExists": "用户已存在",
    "badCredentials": "用户名或密码错误",
    "maxConversationCountReached": "会话数量已达上限",
    "cannotConnectMoreThanOneClient": "当前用户已有对话等待回复，不能再次请求",
    "noUserSelected": "未选择用户",
    "httpError": "HTTP错误",
    "invalidParams": "参数不合法",
    "internal": "内部错误"
  },
  "tips": {
    "loginExpired": "登录已过期。是否跳转到登录页面？",
    "pleaseEnterUsername": "请输入用户名",
    "pleaseEnterPassword": "请输入密码",
    "waiting": "正在思考中...",
    "queueing": "正在回复其他用户，排队中...",
    "deleteConversation": "确定要删除该会话吗？将无法找回历史记录！",
    "deleteConversationSuccess": "删除成功",
    "deleteConversationFailed": "删除失败",
    "changeConversationTitleSuccess": "更改标题成功",
    "changeConversationTitleFailed": "更改标题失败",
    "rename": "请输入新标题",
    "sendMessage": "请输入内容，使用 {0} 发送",
    "logoutConfirm": "确定要退出登录吗？",
    "pleaseEnterNickname": "请输入昵称",
    "pleaseEnterEmail": "请输入邮箱",
    "deleteUserConfirm": "确定要删除用户吗？",
    "deleteUserSuccess": "删除成功",
    "deleteUserFailed": "删除失败",
    "resetPassword": "请输入新密码",
    "resetUserPasswordSuccess": "修改密码成功",
    "resetUserPasswordFailed": "修改密码失败",
    "success": "成功",
    "failed": "失败",
    "requestSuccess": "请求成功",
    "requestFailed": "请求失败",
    "invalidateConversation": "确定要将这些会话标记为invalid吗？将不再出现在用户列表。",
    "vanishConversation": "确定要彻底删除该会话吗？将在 OpenAI 账户上删除对话，并在数据库中删除对应记录。",
    "conversationTitle": "输入对话标题，为空则使用默认标题",
    "whetherUsePaidModel": "请选择使用的模型，为空则使用默认模型",
    "loadConversation": "请从左侧选择对话以加载，或者",
    "newConversation": "发起新对话",
    "loginSuccess": "登录成功",
    "jumpingPage": "正在跳转……",
    "terminated": "连接被服务端终止",
    "loading": "正在加载...",
    "pleaseSelectConversation": "请先选择一个对话",
    "pressEscToExitFullscreen": "已进入全屏，单击Esc退出全屏",
    "refreshed": "已刷新",
    "autoScrolling": "自动滚动"
  },
  "labels": {
    "nickname": "昵称",
    "is_superuser": "是否管理员",
    "active_time": "最后活动时间",
    "chat_status": "当前状态",
    "can_use_paid": "是否可用付费模型",
    "can_use_gpt4": "是否可用GPT-4模型",
    "max_conv_count": "最大对话数量",
    "available_ask_count": "剩余可询问次数",
    "available_gpt4_ask_count": "剩余GPT-4可询问次数"
  }
}
```

## File: frontend/src/main.ts
```typescript
import './style.css';
// eslint-disable-next-line import/no-unresolved
import 'uno.css';
import '@/api/interceptor';
import 'highlight.js/styles/atom-one-dark.css';
import 'highlight.js/lib/common';
import 'katex/dist/katex.css';

import Vue, { createApp } from 'vue';

import App from './App.vue';
import { i18n } from './i18n';
import router from './router';
import pinia from './store';

// import * as Sentry from "@sentry/vue";
// import { BrowserTracing } from "@sentry/tracing";

const app = createApp(App);

// if (import.meta.env.VITE_ENABLE_SENTRY === "yes") {
//   Sentry.init({
//     app,
//     dsn: import.meta.env.VITE_SENTRY_DSN || "",
//     integrations: [
//       new BrowserTracing({
//         routingInstrumentation: Sentry.vueRouterInstrumentation(router),
//         // tracePropagationTargets: ["localhost", "my-site-url.com", /^\//],
//       }),
//     ],
//     tracesSampleRate: 1.0,
//     ignoreErrors: ["AxiosError", "errors."]
//   });
// }

app.use(router);
app.use(pinia);
app.use(i18n);
// app.use(hljs.vuePlugin);

app.mount('#app');

declare global {
  interface Window {
    $message: any;
  }
}
```

## File: frontend/src/router/guard/index.ts
```typescript
import type { Router } from 'vue-router';

import setupPermissionGuard from './permission';
import setupUserLoginInfoGuard from './userLoginInfo';

export default function createRouteGuard(router: Router) {
  setupUserLoginInfoGuard(router);
  setupPermissionGuard(router);
}
```

## File: frontend/src/router/guard/permission.ts
```typescript
import type { LocationQueryRaw, Router } from 'vue-router';

import { i18n } from '@/i18n';
import { useUserStore } from '@/store';
import { Message } from '@/utils/tips';
const t = i18n.global.t as any;

// 在 userLoginInfo 之后，此时要么登录成功，要么未登录
export default function setupPermissionGuard(router: Router) {
  router.beforeEach(async (to, from, next) => {
    const userStore = useUserStore();
    if (!to.meta.requiresAuth) next();
    else {
      if (userStore.user === null) {
        Message.error(t('errors.userNotLogin'));
        next({
          name: 'login',
          query: {
            redirect: to.name,
            ...to.query,
          } as LocationQueryRaw,
        });
      } else {
        // if (to.meta.roles.find((role) => role === userStore.user.role) === ) {
        //   if (userStore.user.is_superuser) next();
        //   else next({ name: "403" });
        // } else next();
        const role = userStore.user.is_superuser ? 'superuser' : 'user';
        if (to.meta.roles.find((r) => r === role) === undefined) {
          next({ name: '403' });
        } else next();
      }
    }
  });
}
```

## File: frontend/src/router/guard/userLoginInfo.ts
```typescript
import type { LocationQueryRaw,Router } from 'vue-router';

import { useUserStore } from '@/store';
import { hasLoginCookie } from '@/utils/auth';
import { LoadingBar } from '@/utils/tips';

// 确保保持登录状态，并及时更新用户信息
export default function setupUserLoginInfoGuard(router: Router) {
  router.beforeEach(async (to, from, next) => {
    const userStore = useUserStore();
    if (hasLoginCookie()) {
      if (userStore.user != null) {
        next();
      } else {
        try {
          await userStore.fetchUserInfo();
          next();
        } catch (error) {
          console.error(error);
          await userStore.logout();
          if (to.name !== 'login') {
            next({
              name: 'login',
              query: {
                redirect: to.name,
                ...to.query,
              } as LocationQueryRaw,
            });
          } else {
            next();
          }
        }
      }
    } else {
      next();
    }
  });
}
```

## File: frontend/src/router/index.ts
```typescript
import { createRouter, createWebHistory } from 'vue-router';

import createRouteGuard from './guard';

const router = createRouter({
  history: createWebHistory(import.meta.env.VITE_ROUTER_BASE),
  routes: [
    {
      path: '/',
      redirect: 'home',
    },
    {
      path: '/home',
      name: 'home',
      component: () => import('@/views/home.vue'),
      meta: {
        requiresAuth: false,
        roles: ['superuser', 'user'],
      },
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/login/index.vue'),
      meta: {
        requiresAuth: false,
        roles: ['superuser', 'user'],
      },
    },
    {
      path: '/conversation',
      name: 'conversation',
      component: () => import('@/views/conversation/index.vue'),
      meta: {
        requiresAuth: true,
        roles: ['superuser', 'user'],
      },
    },
    {
      path: '/conv/:conversation_id',
      name: 'conversationHistory',
      component: () => import('@/views/conversation/history-viewer.vue'),
      meta: {
        requiresAuth: true,
        roles: ['superuser', 'user'],
      },
    },
    {
      path: '/admin',
      name: 'admin',
      redirect: '/admin/system',
      component: () => import('@/views/admin/index.vue'),
      meta: {
        requiresAuth: true,
        roles: ['superuser'],
      },
      children: [
        {
          path: 'system',
          name: 'systemManagement',
          component: () => import('@/views/admin/system_manager.vue'),
        },
        {
          path: 'user',
          name: 'userManagement',
          component: () => import('@/views/admin/user_manager.vue'),
        },
        {
          path: 'conversation',
          name: 'conversationManagement',
          component: () => import('@/views/admin/conversation_manager.vue'),
        },
        {
          path: 'log',
          name: 'logViewer',
          component: () => import('@/views/admin/log_viewer.vue'),
        },
      ],
    },
    {
      path: '/redirect',
      name: 'redirectWrapper',
      children: [
        {
          path: '/redirect/:path',
          name: 'Redirect',
          component: () => import('@/views/redirect/index.vue'),
          meta: {
            requiresAuth: false,
            roles: ['superuser', 'user'],
          },
        },
      ],
    },
    {
      path: '/error',
      name: 'errorPageWrapper',
      children: [
        {
          path: '/error/403',
          name: '403',
          component: () => import('@/views/error/403.vue'),
          meta: {
            requiresAuth: false,
            roles: ['superuser', 'user'],
          },
        },
        {
          path: '/error/404',
          name: '404',
          component: () => import('@/views/error/404.vue'),
          meta: {
            requiresAuth: false,
            roles: ['superuser', 'user'],
          },
        },
      ],
    },
    { path: '/:pathMatch(.*)*', name: 'NotFound', redirect: '/error/404' },
  ],
  scrollBehavior() {
    return { top: 0 };
  },
});

createRouteGuard(router);

export default router;
```

## File: frontend/src/router/typings.d.ts
```typescript
import 'vue-router';

declare type Role = 'superuser' | 'user';

declare module 'vue-router' {
  interface RouteMeta {
    requiresAuth: boolean; // Whether login is required to access the current page (every route must declare)
    roles: Role[]; // The role of the current page (every route must declare)
    ignoreCache?: boolean; // if set true, the page will not be cached
  }
}
```

## File: frontend/src/store/index.ts
```typescript
import { createPinia } from 'pinia';

import useAppStore from './modules/app';
import useConversationStore from './modules/conversation';
import useUserStore from './modules/user';

const pinia = createPinia();

export { useAppStore, useConversationStore,useUserStore };
export default pinia;
```

## File: frontend/src/store/modules/app.ts
```typescript
import { useOsTheme } from 'naive-ui';
import { defineStore } from 'pinia';

import { AppState } from '../types';
const osThemeRef = useOsTheme();
import { useStorage } from '@vueuse/core';

import { setLocale } from '@/i18n';
import { Preference } from '@/types/custom';
import { themeRef } from '@/utils/tips';

const useAppStore = defineStore('app', {
  state: (): AppState => ({
    theme: useStorage('theme', osThemeRef.value),
    language: useStorage('language', 'zh'),
    preference: useStorage<Preference>('preference', {
      sendKey: 'Enter',
      renderUserMessageInMd: false,
      codeAutoWrap: false,
      widerConversationPage: true,
    }),
  }),
  getters: {},
  actions: {
    // 切换主题
    toggleTheme() {
      this.theme = this.theme === 'dark' ? 'light' : 'dark';
      themeRef.value = this.theme;
    },
    setLanguage(lang: string) {
      this.language = lang;
      setLocale(lang);
    },
  },
});

export default useAppStore;
```

## File: frontend/src/store/modules/conversation.ts
```typescript
import { defineStore } from 'pinia';

import { deleteConversationApi, getAllConversationsApi, getConversationHistoryApi, setConversationTitleApi } from '@/api/chat';
import { ChatConversationDetail, ChatMessage } from '@/types/custom';
import { ConversationSchema } from '@/types/schema';

const useConversationStore = defineStore('conversation', {
  state: (): any => ({
    conversations: [] as Array<ConversationSchema>,
    conversationDetailMap: {} as Record<string, ChatConversationDetail>, // conv_id => ChatConversationDetail
  }),
  getters: {},
  actions: {
    async fetchAllConversations() {
      const result = (await getAllConversationsApi()).data;
      this.$patch({ conversations: result });
    },

    async fetchConversationHistory(conversation_id: string) {
      // 解析历史记录
      if (this.conversationDetailMap[conversation_id]) {
        return this.conversationDetailMap[conversation_id];
      }

      const result = (await getConversationHistoryApi(conversation_id)).data;

      const conv_detail: ChatConversationDetail = {
        id: conversation_id,
        current_node: result.current_node,
        title: result.title,
        create_time: result.create_time,
        mapping: {},
        model_name: result.model_name,
      };

      for (const message_id in result.mapping) {
        const current_msg = result.mapping[message_id];
        conv_detail.mapping[message_id] = {
          id: message_id,
          parent: current_msg.parent,
          children: current_msg.children,
          author_role: current_msg.message?.author?.role,
          model_slug: current_msg.message?.metadata?.model_slug,
          message: current_msg.message?.content?.parts.join('\n\n'),
        } as ChatMessage;
      }

      this.$patch({
        conversationDetailMap: {
          [conversation_id]: conv_detail,
        },
      });
    },

    addConversation(conversation: ConversationSchema) {
      this.conversations.push(conversation);
    },

    async deleteConversation(conversation_id: string) {
      await deleteConversationApi(conversation_id);
      delete this.conversationDetailMap[conversation_id];
      this.conversations = this.conversations.filter((conv: any) => conv.conversation_id !== conversation_id);
    },

    async changeConversationTitle(conversation_id: string, title: string) {
      await setConversationTitleApi(conversation_id, title);
      await this.fetchAllConversations();
      if (this.conversationDetailMap[conversation_id]) {
        this.conversationDetailMap[conversation_id].title = title;
      }
    },

    // 仅当收到新信息时调用，为了避免重复获取整个对话历史
    addMessageToConversation(conversation_id: string, sendMessage: ChatMessage, recvMessage: ChatMessage) {
      if (!this.conversationDetailMap[conversation_id]) {
        return;
      }

      const conv_detail = this.conversationDetailMap[conversation_id];
      conv_detail.mapping[sendMessage.id] = sendMessage;
      conv_detail.mapping[recvMessage.id] = recvMessage;

      // 这里只有在新建对话时调用
      if (conv_detail.current_node === null) {
        conv_detail.current_node = recvMessage.id;
      } else {
        const lastTopMessage = conv_detail.mapping[conv_detail.current_node];
        sendMessage.parent = lastTopMessage.id;
        lastTopMessage.children.push(sendMessage.id);
        conv_detail.current_node = recvMessage.id;
      }
      sendMessage.children = [recvMessage.id];
      recvMessage.parent = sendMessage.id;
    },
  },
});

export default useConversationStore;
```

## File: frontend/src/store/modules/user.ts
```typescript
import { defineStore } from 'pinia';

import { getUserInfoApi, loginApi, LoginData, logoutApi } from '@/api/user';
import { UserRead } from '@/types/schema';
import { clearCookie } from '@/utils/auth';

import { UserState } from '../types';

const useUserStore = defineStore('user', {
  state: (): UserState => ({
    user: null,
    savedUsername: null,
    savedPassword: null,
  }),
  getters: {
    userInfo(state: UserState): UserRead | null {
      return state.user;
    },
  },

  actions: {
    // Set user's information
    setInfo(user: UserRead) {
      this.$patch({ user });
    },

    setSavedLoginInfo(username: string, password: string) {
      this.$patch({ savedUsername: username, savedPassword: password });
    },

    // Reset user's information
    resetInfo() {
      this.$reset();
    },

    // Get user's information
    async fetchUserInfo() {
      const result = (await getUserInfoApi()).data;
      this.setInfo(result);
    },

    // Login
    async login(loginForm: LoginData) {
      try {
        await loginApi(loginForm);
        // setToken(res.data.token);
      } catch (err) {
        clearCookie();
        throw err;
      }
    },

    // Logout
    async logout() {
      try {
        await logoutApi();
      } finally {
        this.resetInfo();
        clearCookie();
      }
    },
  },
});

export default useUserStore;
```

## File: frontend/src/store/types.ts
```typescript
import { RemovableRef, UseStorageOptions } from '@vueuse/core';

import { Preference } from '@/types/custom';
import { ConversationSchema,UserRead } from '@/types/schema';

interface UserState {
  user: UserRead | null;
  savedUsername: string | null;
  savedPassword: string | null;
}

interface AppState {
  theme: any;
  language: any;
  preference: RemovableRef<Preference>;
}

export type { AppState,UserState };
```

## File: frontend/src/style.css
```css
@media print {
  .hide-in-print * {
    visibility: hidden !important;
  }
}

/* For WebKit-based browsers (e.g., Chrome, Safari) */
::-webkit-scrollbar {
  width: 8px; /* Change the width of the scrollbar */
  height: 8px; /* Change the height of the scrollbar */
}

::-webkit-scrollbar-track {
  background: transparent; /* Set the scrollbar track color to transparent */
  border-radius: 10px; /* Add border-radius to the scrollbar track */
}

::-webkit-scrollbar-thumb {
  background: #888; /* Change the color of the scrollbar thumb */
  border-radius: 10px; /* Add border-radius to the scrollbar thumb */
}

::-webkit-scrollbar-thumb:hover {
  background: #555; /* Change the color of the scrollbar thumb when hovered */
}

/* For Firefox */
* {
  scrollbar-width: thin; /* Change the width of the scrollbar */
  scrollbar-color: #888 transparent; /* Set the scrollbar thumb color and make the track transparent */
}
```

## File: frontend/src/types/custom.ts
```typescript
export type ChatMessage = {
  id: string;
  author_role: 'user' | 'assistant' | string;
  model_slug?: string;
  message?: string;
  parent?: string | null;
  children: Array<string>;
  typing?: boolean;
};

export type ChatConversationDetail = {
  id: string;
  current_node: string | null;
  title: string;
  create_time: number;
  mapping: Record<string, ChatMessage>;
  model_name: string | null;
};

export type Preference = {
  sendKey: 'Shift+Enter' | 'Enter' | 'Ctrl+Enter';
  renderUserMessageInMd: boolean;
  codeAutoWrap: boolean;
  widerConversationPage: boolean;
};
```

## File: frontend/src/types/echarts.ts
```typescript
import { CallbackDataParams } from 'echarts/types/dist/shared';

export interface ToolTipFormatterParams extends CallbackDataParams {
  axisDim: string;
  axisIndex: number;
  axisType: string;
  axisId: string;
  axisValue: string;
  axisValueLabel: string;
}
```

## File: frontend/src/types/openapi.json
```json
{"openapi":"3.0.2","info":{"title":"FastAPI","version":"0.1.0"},"paths":{"/auth/login":{"post":{"tags":["auth"],"summary":"Auth:Jwt.Login","operationId":"auth_jwt_login_auth_login_post","requestBody":{"content":{"application/x-www-form-urlencoded":{"schema":{"$ref":"#/components/schemas/Body_auth_jwt_login_auth_login_post"}}},"required":true},"responses":{"200":{"description":"Successful Response","content":{"application/json":{"schema":{"type":"string"}}}},"400":{"description":"Bad Request","content":{"application/json":{"schema":{"$ref":"#/components/schemas/ErrorModel"},"examples":{"LOGIN_BAD_CREDENTIALS":{"summary":"Bad credentials or the user is inactive.","value":{"detail":"LOGIN_BAD_CREDENTIALS"}},"LOGIN_USER_NOT_VERIFIED":{"summary":"The user is not verified.","value":{"detail":"LOGIN_USER_NOT_VERIFIED"}}}}}},"422":{"description":"Validation Error","content":{"application/json":{"schema":{"$ref":"#/components/schemas/HTTPValidationError"}}}}}}},"/auth/logout":{"post":{"tags":["auth"],"summary":"Auth:Jwt.Logout","operationId":"auth_jwt_logout_auth_logout_post","responses":{"200":{"description":"Successful Response","content":{"application/json":{"schema":{"type":"string"}}}},"401":{"description":"Missing token or inactive user."}},"security":[{"APIKeyCookie":[]}]}},"/auth/forgot-password":{"post":{"tags":["auth"],"summary":"Reset:Forgot Password","operationId":"reset_forgot_password_auth_forgot_password_post","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/Body_reset_forgot_password_auth_forgot_password_post"}}},"required":true},"responses":{"202":{"description":"Successful Response","content":{"application/json":{"schema":{"type":"string"}}}},"422":{"description":"Validation Error","content":{"application/json":{"schema":{"$ref":"#/components/schemas/HTTPValidationError"}}}}}}},"/auth/reset-password":{"post":{"tags":["auth"],"summary":"Reset:Reset Password","operationId":"reset_reset_password_auth_reset_password_post","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/Body_reset_reset_password_auth_reset_password_post"}}},"required":true},"responses":{"200":{"description":"Successful Response","content":{"application/json":{"schema":{"type":"string"}}}},"400":{"description":"Bad Request","content":{"application/json":{"schema":{"$ref":"#/components/schemas/ErrorModel"},"examples":{"RESET_PASSWORD_BAD_TOKEN":{"summary":"Bad or expired token.","value":{"detail":"RESET_PASSWORD_BAD_TOKEN"}},"RESET_PASSWORD_INVALID_PASSWORD":{"summary":"Password validation failed.","value":{"detail":{"code":"RESET_PASSWORD_INVALID_PASSWORD","reason":"Password should be at least 3 characters"}}}}}}},"422":{"description":"Validation Error","content":{"application/json":{"schema":{"$ref":"#/components/schemas/HTTPValidationError"}}}}}}},"/auth/register":{"post":{"tags":["auth"],"summary":"Register:Register","operationId":"register_register_auth_register_post","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/UserCreate"}}},"required":true},"responses":{"201":{"description":"Successful Response","content":{"application/json":{"schema":{"type":"string"}}}},"400":{"description":"Bad Request","content":{"application/json":{"schema":{"$ref":"#/components/schemas/ErrorModel"},"examples":{"REGISTER_USER_ALREADY_EXISTS":{"summary":"A user with this email already exists.","value":{"detail":"REGISTER_USER_ALREADY_EXISTS"}},"REGISTER_INVALID_PASSWORD":{"summary":"Password validation failed.","value":{"detail":{"code":"REGISTER_INVALID_PASSWORD","reason":"Password should beat least 3 characters"}}}}}}},"422":{"description":"Validation Error","content":{"application/json":{"schema":{"$ref":"#/components/schemas/HTTPValidationError"}}}}}}},"/user":{"get":{"tags":["user"],"summary":"Get All Users","operationId":"get_all_users_user_get","responses":{"200":{"description":"Successful Response","content":{"application/json":{"schema":{"type":"string"}}}}},"security":[{"APIKeyCookie":[]}]}},"/user/{user_id}/reset-password":{"patch":{"tags":["user"],"summary":"Reset Password","operationId":"reset_password_user__user_id__reset_password_patch","parameters":[{"required":true,"schema":{"title":"User Id","type":"integer"},"name":"user_id","in":"path"},{"required":false,"schema":{"title":"New Password","type":"string"},"name":"new_password","in":"query"}],"responses":{"200":{"description":"Successful Response","content":{"application/json":{"schema":{"type":"string"}}}},"422":{"description":"Validation Error","content":{"application/json":{"schema":{"$ref":"#/components/schemas/HTTPValidationError"}}}}},"security":[{"APIKeyCookie":[]}]}},"/user/{user_id}/limit":{"post":{"tags":["user"],"summary":"Update Limit","operationId":"update_limit_user__user_id__limit_post","parameters":[{"required":true,"schema":{"title":"User Id","type":"integer"},"name":"user_id","in":"path"}],"requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/LimitSchema"}}},"required":true},"responses":{"200":{"description":"Successful Response","content":{"application/json":{"schema":{"type":"string"}}}},"422":{"description":"Validation Error","content":{"application/json":{"schema":{"$ref":"#/components/schemas/HTTPValidationError"}}}}},"security":[{"APIKeyCookie":[]}]}},"/user/me":{"get":{"tags":["user"],"summary":"Users:Current User","operationId":"users_current_user_user_me_get","responses":{"200":{"description":"Successful Response","content":{"application/json":{"schema":{"type":"string"}}}},"401":{"description":"Missing token or inactive user."}},"security":[{"APIKeyCookie":[]}]},"patch":{"tags":["user"],"summary":"Users:Patch Current User","operationId":"users_patch_current_user_user_me_patch","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/UserUpdate"}}},"required":true},"responses":{"200":{"description":"Successful Response","content":{"application/json":{"schema":{"type":"string"}}}},"401":{"description":"Missing token or inactive user."},"400":{"description":"Bad Request","content":{"application/json":{"schema":{"$ref":"#/components/schemas/ErrorModel"},"examples":{"UPDATE_USER_EMAIL_ALREADY_EXISTS":{"summary":"A user with this email already exists.","value":{"detail":"UPDATE_USER_EMAIL_ALREADY_EXISTS"}},"UPDATE_USER_INVALID_PASSWORD":{"summary":"Password validation failed.","value":{"detail":{"code":"UPDATE_USER_INVALID_PASSWORD","reason":"Password should beat least 3 characters"}}}}}}},"422":{"description":"Validation Error","content":{"application/json":{"schema":{"$ref":"#/components/schemas/HTTPValidationError"}}}}},"security":[{"APIKeyCookie":[]}]}},"/user/{id}":{"get":{"tags":["user"],"summary":"Users:User","operationId":"users_user_user__id__get","parameters":[{"required":true,"schema":{"title":"Id"},"name":"id","in":"path"}],"responses":{"200":{"description":"Successful Response","content":{"application/json":{"schema":{"type":"string"}}}},"401":{"description":"Missing token or inactive user."},"403":{"description":"Not a superuser."},"404":{"description":"The user does not exist."},"422":{"description":"Validation Error","content":{"application/json":{"schema":{"$ref":"#/components/schemas/HTTPValidationError"}}}}},"security":[{"APIKeyCookie":[]}]},"delete":{"tags":["user"],"summary":"Users:Delete User","operationId":"users_delete_user_user__id__delete","parameters":[{"required":true,"schema":{"title":"Id"},"name":"id","in":"path"}],"responses":{"204":{"description":"Successful Response"},"401":{"description":"Missing token or inactive user."},"403":{"description":"Not a superuser."},"404":{"description":"The user does not exist."},"422":{"description":"Validation Error","content":{"application/json":{"schema":{"$ref":"#/components/schemas/HTTPValidationError"}}}}},"security":[{"APIKeyCookie":[]}]},"patch":{"tags":["user"],"summary":"Users:Patch User","operationId":"users_patch_user_user__id__patch","parameters":[{"required":true,"schema":{"title":"Id"},"name":"id","in":"path"}],"requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/UserUpdate"}}},"required":true},"responses":{"200":{"description":"Successful Response","content":{"application/json":{"schema":{"type":"string"}}}},"401":{"description":"Missing token or inactive user."},"403":{"description":"Not a superuser."},"404":{"description":"The user does not exist."},"400":{"description":"Bad Request","content":{"application/json":{"schema":{"$ref":"#/components/schemas/ErrorModel"},"examples":{"UPDATE_USER_EMAIL_ALREADY_EXISTS":{"summary":"A user with this email already exists.","value":{"detail":"UPDATE_USER_EMAIL_ALREADY_EXISTS"}},"UPDATE_USER_INVALID_PASSWORD":{"summary":"Password validation failed.","value":{"detail":{"code":"UPDATE_USER_INVALID_PASSWORD","reason":"Password should beat least 3 characters"}}}}}}},"422":{"description":"Validation Error","content":{"application/json":{"schema":{"$ref":"#/components/schemas/HTTPValidationError"}}}}},"security":[{"APIKeyCookie":[]}]}},"/conv":{"get":{"tags":["conversation"],"summary":"Get All Conversations","description":"返回自己的有效会话\n对于管理员，返回所有对话，并可以指定是否只返回有效会话","operationId":"get_all_conversations_conv_get","parameters":[{"required":false,"schema":{"title":"Fetch All","type":"boolean","default":false},"name":"fetch_all","in":"query"}],"responses":{"200":{"description":"Successful Response","content":{"application/json":{"schema":{"type":"string"}}}},"422":{"description":"Validation Error","content":{"application/json":{"schema":{"$ref":"#/components/schemas/HTTPValidationError"}}}}},"security":[{"APIKeyCookie":[]}]},"delete":{"tags":["conversation"],"summary":"Delete All Conversation","operationId":"delete_all_conversation_conv_delete","responses":{"200":{"description":"Successful Response","content":{"application/json":{"schema":{"type":"string"}}}}},"security":[{"APIKeyCookie":[]}]}},"/conv/{conversation_id}":{"get":{"tags":["conversation"],"summary":"Get Conversation History","operationId":"get_conversation_history_conv__conversation_id__get","parameters":[{"required":true,"schema":{"title":"Conversation Id","type":"string"},"name":"conversation_id","in":"path"}],"responses":{"200":{"description":"Successful Response","content":{"application/json":{"schema":{"type":"string"}}}},"422":{"description":"Validation Error","content":{"application/json":{"schema":{"$ref":"#/components/schemas/HTTPValidationError"}}}}},"security":[{"APIKeyCookie":[]}]},"delete":{"tags":["conversation"],"summary":"Delete Conversation","description":"remove conversation from database and chatgpt server","operationId":"delete_conversation_conv__conversation_id__delete","parameters":[{"required":true,"schema":{"title":"Conversation Id","type":"string"},"name":"conversation_id","in":"path"}],"responses":{"200":{"description":"Successful Response","content":{"application/json":{"schema":{"type":"string"}}}},"422":{"description":"Validation Error","content":{"application/json":{"schema":{"$ref":"#/components/schemas/HTTPValidationError"}}}}},"security":[{"APIKeyCookie":[]}]},"patch":{"tags":["conversation"],"summary":"Change Conversation Title","operationId":"change_conversation_title_conv__conversation_id__patch","parameters":[{"required":true,"schema":{"title":"Conversation Id","type":"string"},"name":"conversation_id","in":"path"},{"required":true,"schema":{"title":"Title","type":"string"},"name":"title","in":"query"}],"responses":{"200":{"description":"Successful Response","content":{"application/json":{"schema":{"type":"string"}}}},"422":{"description":"Validation Error","content":{"application/json":{"schema":{"$ref":"#/components/schemas/HTTPValidationError"}}}}},"security":[{"APIKeyCookie":[]}]}},"/conv/{conversation_id}/vanish":{"delete":{"tags":["conversation"],"summary":"Vanish Conversation","operationId":"vanish_conversation_conv__conversation_id__vanish_delete","parameters":[{"required":true,"schema":{"title":"Conversation Id","type":"string"},"name":"conversation_id","in":"path"}],"responses":{"200":{"description":"Successful Response","content":{"application/json":{"schema":{"type":"string"}}}},"422":{"description":"Validation Error","content":{"application/json":{"schema":{"$ref":"#/components/schemas/HTTPValidationError"}}}}},"security":[{"APIKeyCookie":[]}]}},"/conv/{conversation_id}/assign/{username}":{"patch":{"tags":["conversation"],"summary":"Assign Conversation","operationId":"assign_conversation_conv__conversation_id__assign__username__patch","parameters":[{"required":true,"schema":{"title":"Username","type":"string"},"name":"username","in":"path"},{"required":true,"schema":{"title":"Conversation Id","type":"string"},"name":"conversation_id","in":"path"}],"responses":{"200":{"description":"Successful Response","content":{"application/json":{"schema":{"type":"string"}}}},"422":{"description":"Validation Error","content":{"application/json":{"schema":{"$ref":"#/components/schemas/HTTPValidationError"}}}}},"security":[{"APIKeyCookie":[]}]}},"/conv/{conversation_id}/gen_title":{"patch":{"tags":["conversation"],"summary":"Generate Conversation Title","operationId":"generate_conversation_title_conv__conversation_id__gen_title_patch","parameters":[{"required":true,"schema":{"title":"Conversation Id","type":"string"},"name":"conversation_id","in":"path"},{"required":true,"schema":{"title":"Message Id","type":"string"},"name":"message_id","in":"query"}],"responses":{"200":{"description":"Successful Response","content":{"application/json":{"schema":{"type":"string"}}}},"422":{"description":"Validation Error","content":{"application/json":{"schema":{"$ref":"#/components/schemas/HTTPValidationError"}}}}},"security":[{"APIKeyCookie":[]}]}},"/system/info":{"get":{"tags":["system"],"summary":"Get System Info","operationId":"get_system_info_system_info_get","responses":{"200":{"description":"Successful Response","content":{"application/json":{"schema":{"type":"string"}}}}},"security":[{"APIKeyCookie":[]}]}},"/system/request_statistics":{"get":{"tags":["system"],"summary":"Get Request Statistics","operationId":"get_request_statistics_system_request_statistics_get","responses":{"200":{"description":"Successful Response","content":{"application/json":{"schema":{"type":"string"}}}}},"security":[{"APIKeyCookie":[]}]}},"/system/proxy_logs":{"post":{"tags":["system"],"summary":"Get Proxy Logs","operationId":"get_proxy_logs_system_proxy_logs_post","requestBody":{"content":{"application/json":{"schema":{"title":"Options","allOf":[{"$ref":"#/components/schemas/LogFilterOptions"}],"default":{"max_lines":100}}}}},"responses":{"200":{"description":"Successful Response","content":{"application/json":{"schema":{"type":"string"}}}},"422":{"description":"Validation Error","content":{"application/json":{"schema":{"$ref":"#/components/schemas/HTTPValidationError"}}}}},"security":[{"APIKeyCookie":[]}]}},"/system/server_logs":{"post":{"tags":["system"],"summary":"Get Server Logs","operationId":"get_server_logs_system_server_logs_post","requestBody":{"content":{"application/json":{"schema":{"title":"Options","allOf":[{"$ref":"#/components/schemas/LogFilterOptions"}],"default":{"max_lines":100}}}}},"responses":{"200":{"description":"Successful Response","content":{"application/json":{"schema":{"type":"string"}}}},"422":{"description":"Validation Error","content":{"application/json":{"schema":{"$ref":"#/components/schemas/HTTPValidationError"}}}}},"security":[{"APIKeyCookie":[]}]}},"/status":{"get":{"tags":["status"],"summary":"Get Server Status","description":"普通用户获取服务器状态","operationId":"get_server_status_status_get","responses":{"200":{"description":"Successful Response","content":{"application/json":{"schema":{"type":"string"}}}}},"security":[{"APIKeyCookie":[]}]}}},"components":{"schemas":{"Body_auth_jwt_login_auth_login_post":{"title":"Body_auth_jwt_login_auth_login_post","required":["username","password"],"type":"object","properties":{"grant_type":{"title":"Grant Type","pattern":"password","type":"string"},"username":{"title":"Username","type":"string"},"password":{"title":"Password","type":"string"},"scope":{"title":"Scope","type":"string","default":""},"client_id":{"title":"Client Id","type":"string"},"client_secret":{"title":"Client Secret","type":"string"}}},"Body_reset_forgot_password_auth_forgot_password_post":{"title":"Body_reset_forgot_password_auth_forgot_password_post","required":["email"],"type":"object","properties":{"email":{"title":"Email","type":"string","format":"email"}}},"Body_reset_reset_password_auth_reset_password_post":{"title":"Body_reset_reset_password_auth_reset_password_post","required":["token","password"],"type":"object","properties":{"token":{"title":"Token","type":"string"},"password":{"title":"Password","type":"string"}}},"ChatModels":{"title":"ChatModels","enum":["gpt-4","gpt-4-mobile","text-davinci-002-render-sha","text-davinci-002-render-paid",""],"description":"An enumeration."},"ChatStatus":{"title":"ChatStatus","enum":["asking","queueing","idling"],"description":"An enumeration."},"ConversationSchema":{"title":"ConversationSchema","type":"object","properties":{"id":{"title":"Id","type":"integer","default":-1},"conversation_id":{"title":"Conversation Id","type":"string","format":"uuid"},"title":{"title":"Title","type":"string"},"user_id":{"title":"User Id","type":"integer"},"is_valid":{"title":"Is Valid","type":"boolean"},"model_name":{"$ref":"#/components/schemas/ChatModels"},"create_time":{"title":"Create Time","type":"string","format":"date-time"},"active_time":{"title":"Active Time","type":"string","format":"date-time"}}},"ErrorModel":{"title":"ErrorModel","required":["detail"],"type":"object","properties":{"detail":{"title":"Detail","anyOf":[{"type":"string"},{"type":"object","additionalProperties":{"type":"string"}}]}}},"HTTPValidationError":{"title":"HTTPValidationError","type":"object","properties":{"detail":{"title":"Detail","type":"array","items":{"$ref":"#/components/schemas/ValidationError"}}}},"LimitSchema":{"title":"LimitSchema","type":"object","properties":{"can_use_paid":{"title":"Can Use Paid","type":"boolean"},"can_use_gpt4":{"title":"Can Use Gpt4","type":"boolean"},"max_conv_count":{"title":"Max Conv Count","type":"integer"},"available_ask_count":{"title":"Available Ask Count","type":"integer"},"available_gpt4_ask_count":{"title":"Available Gpt4 Ask Count","type":"integer"}}},"LogFilterOptions":{"title":"LogFilterOptions","type":"object","properties":{"max_lines":{"title":"Max Lines","type":"integer","default":100},"exclude_keywords":{"title":"Exclude Keywords","type":"array","items":{"type":"string"}}}},"RequestStatistics":{"title":"RequestStatistics","required":["request_counts_interval","request_counts","ask_records"],"type":"object","properties":{"request_counts_interval":{"title":"Request Counts Interval","type":"integer"},"request_counts":{"title":"Request Counts","type":"object","additionalProperties":{"type":"array","items":{}}},"ask_records":{"title":"Ask Records","type":"array","items":{}}}},"ServerStatusSchema":{"title":"ServerStatusSchema","type":"object","properties":{"active_user_in_5m":{"title":"Active User In 5M","type":"integer"},"active_user_in_1h":{"title":"Active User In 1H","type":"integer"},"active_user_in_1d":{"title":"Active User In 1D","type":"integer"},"is_chatbot_busy":{"title":"Is Chatbot Busy","type":"boolean"},"chatbot_waiting_count":{"title":"Chatbot Waiting Count","type":"integer"}}},"SystemInfo":{"title":"SystemInfo","required":["startup_time","total_user_count","total_conversation_count","valid_conversation_count"],"type":"object","properties":{"startup_time":{"title":"Startup Time","type":"number"},"total_user_count":{"title":"Total User Count","type":"integer"},"total_conversation_count":{"title":"Total Conversation Count","type":"integer"},"valid_conversation_count":{"title":"Valid Conversation Count","type":"integer"}}},"UserCreate":{"title":"UserCreate","required":["email","password","username","nickname"],"type":"object","properties":{"email":{"title":"Email","type":"string"},"password":{"title":"Password","type":"string"},"is_active":{"title":"Is Active","type":"boolean","default":true},"is_superuser":{"title":"Is Superuser","type":"boolean","default":false},"is_verified":{"title":"Is Verified","type":"boolean","default":false},"username":{"title":"Username","type":"string"},"nickname":{"title":"Nickname","type":"string"},"can_use_paid":{"title":"Can Use Paid","type":"boolean","default":false},"max_conv_count":{"title":"Max Conv Count","type":"integer","default":-1},"available_ask_count":{"title":"Available Ask Count","type":"integer","default":-1}}},"UserRead":{"title":"UserRead","required":["id","email","is_active","is_superuser","is_verified","username","nickname","chat_status","can_use_paid","can_use_gpt4"],"type":"object","properties":{"id":{"title":"Id","type":"integer"},"email":{"title":"Email","type":"string"},"is_active":{"title":"Is Active","type":"boolean"},"is_superuser":{"title":"Is Superuser","type":"boolean"},"is_verified":{"title":"Is Verified","type":"boolean"},"username":{"title":"Username","type":"string"},"nickname":{"title":"Nickname","type":"string"},"active_time":{"title":"Active Time","type":"string","format":"date-time"},"chat_status":{"$ref":"#/components/schemas/ChatStatus"},"can_use_paid":{"title":"Can Use Paid","type":"boolean"},"can_use_gpt4":{"title":"Can Use Gpt4","type":"boolean"},"max_conv_count":{"title":"Max Conv Count","type":"integer"},"available_ask_count":{"title":"Available Ask Count","type":"integer"},"available_gpt4_ask_count":{"title":"Available Gpt4 Ask Count","type":"integer"}},"description":"Base User model."},"UserUpdate":{"title":"UserUpdate","required":["nickname"],"type":"object","properties":{"id":{"title":"Id"},"email":{"title":"Email","type":"string"},"is_active":{"title":"Is Active","type":"boolean","default":true},"is_superuser":{"title":"Is Superuser","type":"boolean","default":false},"is_verified":{"title":"Is Verified","type":"boolean","default":false},"nickname":{"title":"Nickname","type":"string"}},"description":"Base User model."},"ValidationError":{"title":"ValidationError","required":["loc","msg","type"],"type":"object","properties":{"loc":{"title":"Location","type":"array","items":{"anyOf":[{"type":"string"},{"type":"integer"}]}},"msg":{"title":"Message","type":"string"},"type":{"title":"Error Type","type":"string"}}}},"securitySchemes":{"APIKeyCookie":{"type":"apiKey","in":"cookie","name":"user_auth"}}}}
```

## File: frontend/src/types/openapi.ts
```typescript
/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */

export interface paths {
  '/auth/login': {
    /** Auth:Jwt.Login */
    post: operations['auth_jwt_login_auth_login_post'];
  };
  '/auth/logout': {
    /** Auth:Jwt.Logout */
    post: operations['auth_jwt_logout_auth_logout_post'];
  };
  '/auth/forgot-password': {
    /** Reset:Forgot Password */
    post: operations['reset_forgot_password_auth_forgot_password_post'];
  };
  '/auth/reset-password': {
    /** Reset:Reset Password */
    post: operations['reset_reset_password_auth_reset_password_post'];
  };
  '/auth/register': {
    /** Register:Register */
    post: operations['register_register_auth_register_post'];
  };
  '/user': {
    /** Get All Users */
    get: operations['get_all_users_user_get'];
  };
  '/user/{user_id}/reset-password': {
    /** Reset Password */
    patch: operations['reset_password_user__user_id__reset_password_patch'];
  };
  '/user/{user_id}/limit': {
    /** Update Limit */
    post: operations['update_limit_user__user_id__limit_post'];
  };
  '/user/me': {
    /** Users:Current User */
    get: operations['users_current_user_user_me_get'];
    /** Users:Patch Current User */
    patch: operations['users_patch_current_user_user_me_patch'];
  };
  '/user/{id}': {
    /** Users:User */
    get: operations['users_user_user__id__get'];
    /** Users:Delete User */
    delete: operations['users_delete_user_user__id__delete'];
    /** Users:Patch User */
    patch: operations['users_patch_user_user__id__patch'];
  };
  '/conv': {
    /**
     * Get All Conversations
     * @description 返回自己的有效会话
     * 对于管理员，返回所有对话，并可以指定是否只返回有效会话
     */
    get: operations['get_all_conversations_conv_get'];
  };
  '/conv/{conversation_id}': {
    /** Get Conversation History */
    get: operations['get_conversation_history_conv__conversation_id__get'];
    /**
     * Delete Conversation
     * @description remove conversation from database and chatgpt server
     */
    delete: operations['delete_conversation_conv__conversation_id__delete'];
    /** Change Conversation Title */
    patch: operations['change_conversation_title_conv__conversation_id__patch'];
  };
  '/conv/{conversation_id}/vanish': {
    /** Vanish Conversation */
    delete: operations['vanish_conversation_conv__conversation_id__vanish_delete'];
  };
  '/conv/{conversation_id}/assign/{username}': {
    /** Assign Conversation */
    patch: operations['assign_conversation_conv__conversation_id__assign__username__patch'];
  };
  '/conv/{conversation_id}/gen_title': {
    /** Generate Conversation Title */
    patch: operations['generate_conversation_title_conv__conversation_id__gen_title_patch'];
  };
  '/system/info': {
    /** Get System Info */
    get: operations['get_system_info_system_info_get'];
  };
  '/system/request_statistics': {
    /** Get Request Statistics */
    get: operations['get_request_statistics_system_request_statistics_get'];
  };
  '/system/proxy_logs': {
    /** Get Proxy Logs */
    post: operations['get_proxy_logs_system_proxy_logs_post'];
  };
  '/system/server_logs': {
    /** Get Server Logs */
    post: operations['get_server_logs_system_server_logs_post'];
  };
  '/status': {
    /**
     * Get Server Status
     * @description 普通用户获取服务器状态
     */
    get: operations['get_server_status_status_get'];
  };
}

export type webhooks = Record<string, never>;

export interface components {
  schemas: {
    /** Body_auth_jwt_login_auth_login_post */
    Body_auth_jwt_login_auth_login_post: {
      /** Grant Type */
      grant_type?: string;
      /** Username */
      username: string;
      /** Password */
      password: string;
      /**
       * Scope
       * @default
       */
      scope?: string;
      /** Client Id */
      client_id?: string;
      /** Client Secret */
      client_secret?: string;
    };
    /** Body_reset_forgot_password_auth_forgot_password_post */
    Body_reset_forgot_password_auth_forgot_password_post: {
      /**
       * Email
       * Format: email
       */
      email: string;
    };
    /** Body_reset_reset_password_auth_reset_password_post */
    Body_reset_reset_password_auth_reset_password_post: {
      /** Token */
      token: string;
      /** Password */
      password: string;
    };
    /**
     * ChatModels
     * @description An enumeration.
     * @enum {unknown}
     */
    ChatModels: 'gpt-4' |'gpt-4-mobile' | 'text-davinci-002-render-sha' | 'text-davinci-002-render-paid' | '';
    /**
     * ChatStatus
     * @description An enumeration.
     * @enum {unknown}
     */
    ChatStatus: 'asking' | 'queueing' | 'idling';
    /** ConversationSchema */
    ConversationSchema: {
      /**
       * Id
       * @default -1
       */
      id?: number;
      /**
       * Conversation Id
       * Format: uuid
       */
      conversation_id?: string;
      /** Title */
      title?: string;
      /** User Id */
      user_id?: number;
      /** Is Valid */
      is_valid?: boolean;
      model_name?: components['schemas']['ChatModels'];
      /**
       * Create Time
       * Format: date-time
       */
      create_time?: string;
      /**
       * Active Time
       * Format: date-time
       */
      active_time?: string;
    };
    /** ErrorModel */
    ErrorModel: {
      /** Detail */
      detail:
        | string
        | {
            [key: string]: string | undefined;
          };
    };
    /** HTTPValidationError */
    HTTPValidationError: {
      /** Detail */
      detail?: components['schemas']['ValidationError'][];
    };
    /** LimitSchema */
    LimitSchema: {
      /** Can Use Paid */
      can_use_paid?: boolean;
      /** Can Use Gpt4 */
      can_use_gpt4?: boolean;
      /** Max Conv Count */
      max_conv_count?: number;
      /** Available Ask Count */
      available_ask_count?: number;
      /** Available Gpt4 Ask Count */
      available_gpt4_ask_count?: number;
    };
    /** LogFilterOptions */
    LogFilterOptions: {
      /**
       * Max Lines
       * @default 100
       */
      max_lines?: number;
      /** Exclude Keywords */
      exclude_keywords?: string[];
    };
    /** RequestStatistics */
    RequestStatistics: {
      /** Request Counts Interval */
      request_counts_interval: number;
      /** Request Counts */
      request_counts: {
        [key: string]: Record<string, never>[] | undefined;
      };
      /** Ask Records */
      ask_records: Record<string, never>[];
    };
    /** ServerStatusSchema */
    ServerStatusSchema: {
      /** Active User In 5M */
      active_user_in_5m?: number;
      /** Active User In 1H */
      active_user_in_1h?: number;
      /** Active User In 1D */
      active_user_in_1d?: number;
      /** Is Chatbot Busy */
      is_chatbot_busy?: boolean;
      /** Chatbot Waiting Count */
      chatbot_waiting_count?: number;
    };
    /** SystemInfo */
    SystemInfo: {
      /** Startup Time */
      startup_time: number;
      /** Total User Count */
      total_user_count: number;
      /** Total Conversation Count */
      total_conversation_count: number;
      /** Valid Conversation Count */
      valid_conversation_count: number;
    };
    /** UserCreate */
    UserCreate: {
      /** Email */
      email: string;
      /** Password */
      password: string;
      /**
       * Is Active
       * @default true
       */
      is_active?: boolean;
      /**
       * Is Superuser
       * @default false
       */
      is_superuser?: boolean;
      /**
       * Is Verified
       * @default false
       */
      is_verified?: boolean;
      /** Username */
      username: string;
      /** Nickname */
      nickname: string;
      /**
       * Can Use Paid
       * @default false
       */
      can_use_paid?: boolean;
      /**
       * Max Conv Count
       * @default -1
       */
      max_conv_count?: number;
      /**
       * Available Ask Count
       * @default -1
       */
      available_ask_count?: number;
    };
    /**
     * UserRead
     * @description Base User model.
     */
    UserRead: {
      /** Id */
      id: number;
      /** Email */
      email: string;
      /** Is Active */
      is_active: boolean;
      /** Is Superuser */
      is_superuser: boolean;
      /** Is Verified */
      is_verified: boolean;
      /** Username */
      username: string;
      /** Nickname */
      nickname: string;
      /**
       * Active Time
       * Format: date-time
       */
      active_time?: string;
      chat_status: components['schemas']['ChatStatus'];
      /** Can Use Paid */
      can_use_paid: boolean;
      /** Can Use Gpt4 */
      can_use_gpt4: boolean;
      /** Max Conv Count */
      max_conv_count?: number;
      /** Available Ask Count */
      available_ask_count?: number;
      /** Available Gpt4 Ask Count */
      available_gpt4_ask_count?: number;
    };
    /**
     * UserUpdate
     * @description Base User model.
     */
    UserUpdate: {
      /** Id */
      id?: Record<string, never>;
      /** Email */
      email?: string;
      /**
       * Is Active
       * @default true
       */
      is_active?: boolean;
      /**
       * Is Superuser
       * @default false
       */
      is_superuser?: boolean;
      /**
       * Is Verified
       * @default false
       */
      is_verified?: boolean;
      /** Nickname */
      nickname: string;
    };
    /** ValidationError */
    ValidationError: {
      /** Location */
      loc: (string | number)[];
      /** Message */
      msg: string;
      /** Error Type */
      type: string;
    };
  };
  responses: never;
  parameters: never;
  requestBodies: never;
  headers: never;
  pathItems: never;
}

export type external = Record<string, never>;

export interface operations {
  auth_jwt_login_auth_login_post: {
    /** Auth:Jwt.Login */
    requestBody: {
      content: {
        'application/x-www-form-urlencoded': components['schemas']['Body_auth_jwt_login_auth_login_post'];
      };
    };
    responses: {
      /** @description Successful Response */
      200: {
        content: {
          'application/json': string;
        };
      };
      /** @description Bad Request */
      400: {
        content: {
          'application/json': components['schemas']['ErrorModel'];
        };
      };
      /** @description Validation Error */
      422: {
        content: {
          'application/json': components['schemas']['HTTPValidationError'];
        };
      };
    };
  };
  auth_jwt_logout_auth_logout_post: {
    /** Auth:Jwt.Logout */
    responses: {
      /** @description Successful Response */
      200: {
        content: {
          'application/json': string;
        };
      };
      /** @description Missing token or inactive user. */
      401: never;
    };
  };
  reset_forgot_password_auth_forgot_password_post: {
    /** Reset:Forgot Password */
    requestBody: {
      content: {
        'application/json': components['schemas']['Body_reset_forgot_password_auth_forgot_password_post'];
      };
    };
    responses: {
      /** @description Successful Response */
      202: {
        content: {
          'application/json': string;
        };
      };
      /** @description Validation Error */
      422: {
        content: {
          'application/json': components['schemas']['HTTPValidationError'];
        };
      };
    };
  };
  reset_reset_password_auth_reset_password_post: {
    /** Reset:Reset Password */
    requestBody: {
      content: {
        'application/json': components['schemas']['Body_reset_reset_password_auth_reset_password_post'];
      };
    };
    responses: {
      /** @description Successful Response */
      200: {
        content: {
          'application/json': string;
        };
      };
      /** @description Bad Request */
      400: {
        content: {
          'application/json': components['schemas']['ErrorModel'];
        };
      };
      /** @description Validation Error */
      422: {
        content: {
          'application/json': components['schemas']['HTTPValidationError'];
        };
      };
    };
  };
  register_register_auth_register_post: {
    /** Register:Register */
    requestBody: {
      content: {
        'application/json': components['schemas']['UserCreate'];
      };
    };
    responses: {
      /** @description Successful Response */
      201: {
        content: {
          'application/json': string;
        };
      };
      /** @description Bad Request */
      400: {
        content: {
          'application/json': components['schemas']['ErrorModel'];
        };
      };
      /** @description Validation Error */
      422: {
        content: {
          'application/json': components['schemas']['HTTPValidationError'];
        };
      };
    };
  };
  get_all_users_user_get: {
    /** Get All Users */
    responses: {
      /** @description Successful Response */
      200: {
        content: {
          'application/json': string;
        };
      };
    };
  };
  reset_password_user__user_id__reset_password_patch: {
    /** Reset Password */
    parameters: {
      query?: {
        new_password?: string;
      };
      path: {
        user_id: number;
      };
    };
    responses: {
      /** @description Successful Response */
      200: {
        content: {
          'application/json': string;
        };
      };
      /** @description Validation Error */
      422: {
        content: {
          'application/json': components['schemas']['HTTPValidationError'];
        };
      };
    };
  };
  update_limit_user__user_id__limit_post: {
    /** Update Limit */
    parameters: {
      path: {
        user_id: number;
      };
    };
    requestBody: {
      content: {
        'application/json': components['schemas']['LimitSchema'];
      };
    };
    responses: {
      /** @description Successful Response */
      200: {
        content: {
          'application/json': string;
        };
      };
      /** @description Validation Error */
      422: {
        content: {
          'application/json': components['schemas']['HTTPValidationError'];
        };
      };
    };
  };
  users_current_user_user_me_get: {
    /** Users:Current User */
    responses: {
      /** @description Successful Response */
      200: {
        content: {
          'application/json': string;
        };
      };
      /** @description Missing token or inactive user. */
      401: never;
    };
  };
  users_patch_current_user_user_me_patch: {
    /** Users:Patch Current User */
    requestBody: {
      content: {
        'application/json': components['schemas']['UserUpdate'];
      };
    };
    responses: {
      /** @description Successful Response */
      200: {
        content: {
          'application/json': string;
        };
      };
      /** @description Bad Request */
      400: {
        content: {
          'application/json': components['schemas']['ErrorModel'];
        };
      };
      /** @description Missing token or inactive user. */
      401: never;
      /** @description Validation Error */
      422: {
        content: {
          'application/json': components['schemas']['HTTPValidationError'];
        };
      };
    };
  };
  users_user_user__id__get: {
    /** Users:User */
    parameters: {
      path: {
        id: Record<string, never>;
      };
    };
    responses: {
      /** @description Successful Response */
      200: {
        content: {
          'application/json': string;
        };
      };
      /** @description Missing token or inactive user. */
      401: never;
      /** @description Not a superuser. */
      403: never;
      /** @description The user does not exist. */
      404: never;
      /** @description Validation Error */
      422: {
        content: {
          'application/json': components['schemas']['HTTPValidationError'];
        };
      };
    };
  };
  users_delete_user_user__id__delete: {
    /** Users:Delete User */
    parameters: {
      path: {
        id: Record<string, never>;
      };
    };
    responses: {
      /** @description Successful Response */
      204: never;
      /** @description Missing token or inactive user. */
      401: never;
      /** @description Not a superuser. */
      403: never;
      /** @description The user does not exist. */
      404: never;
      /** @description Validation Error */
      422: {
        content: {
          'application/json': components['schemas']['HTTPValidationError'];
        };
      };
    };
  };
  users_patch_user_user__id__patch: {
    /** Users:Patch User */
    parameters: {
      path: {
        id: Record<string, never>;
      };
    };
    requestBody: {
      content: {
        'application/json': components['schemas']['UserUpdate'];
      };
    };
    responses: {
      /** @description Successful Response */
      200: {
        content: {
          'application/json': string;
        };
      };
      /** @description Bad Request */
      400: {
        content: {
          'application/json': components['schemas']['ErrorModel'];
        };
      };
      /** @description Missing token or inactive user. */
      401: never;
      /** @description Not a superuser. */
      403: never;
      /** @description The user does not exist. */
      404: never;
      /** @description Validation Error */
      422: {
        content: {
          'application/json': components['schemas']['HTTPValidationError'];
        };
      };
    };
  };
  get_all_conversations_conv_get: {
    /**
     * Get All Conversations
     * @description 返回自己的有效会话
     * 对于管理员，返回所有对话，并可以指定是否只返回有效会话
     */
    parameters?: {
      query?: {
        fetch_all?: boolean;
      };
    };
    responses: {
      /** @description Successful Response */
      200: {
        content: {
          'application/json': string;
        };
      };
      /** @description Validation Error */
      422: {
        content: {
          'application/json': components['schemas']['HTTPValidationError'];
        };
      };
    };
  };
  get_conversation_history_conv__conversation_id__get: {
    /** Get Conversation History */
    parameters: {
      path: {
        conversation_id: string;
      };
    };
    responses: {
      /** @description Successful Response */
      200: {
        content: {
          'application/json': string;
        };
      };
      /** @description Validation Error */
      422: {
        content: {
          'application/json': components['schemas']['HTTPValidationError'];
        };
      };
    };
  };
  delete_conversation_conv__conversation_id__delete: {
    /**
     * Delete Conversation
     * @description remove conversation from database and chatgpt server
     */
    parameters: {
      path: {
        conversation_id: string;
      };
    };
    responses: {
      /** @description Successful Response */
      200: {
        content: {
          'application/json': string;
        };
      };
      /** @description Validation Error */
      422: {
        content: {
          'application/json': components['schemas']['HTTPValidationError'];
        };
      };
    };
  };
  change_conversation_title_conv__conversation_id__patch: {
    /** Change Conversation Title */
    parameters: {
      query: {
        title: string;
      };
      path: {
        conversation_id: string;
      };
    };
    responses: {
      /** @description Successful Response */
      200: {
        content: {
          'application/json': string;
        };
      };
      /** @description Validation Error */
      422: {
        content: {
          'application/json': components['schemas']['HTTPValidationError'];
        };
      };
    };
  };
  vanish_conversation_conv__conversation_id__vanish_delete: {
    /** Vanish Conversation */
    parameters: {
      path: {
        conversation_id: string;
      };
    };
    responses: {
      /** @description Successful Response */
      200: {
        content: {
          'application/json': string;
        };
      };
      /** @description Validation Error */
      422: {
        content: {
          'application/json': components['schemas']['HTTPValidationError'];
        };
      };
    };
  };
  assign_conversation_conv__conversation_id__assign__username__patch: {
    /** Assign Conversation */
    parameters: {
      path: {
        username: string;
        conversation_id: string;
      };
    };
    responses: {
      /** @description Successful Response */
      200: {
        content: {
          'application/json': string;
        };
      };
      /** @description Validation Error */
      422: {
        content: {
          'application/json': components['schemas']['HTTPValidationError'];
        };
      };
    };
  };
  generate_conversation_title_conv__conversation_id__gen_title_patch: {
    /** Generate Conversation Title */
    parameters: {
      query: {
        message_id: string;
      };
      path: {
        conversation_id: string;
      };
    };
    responses: {
      /** @description Successful Response */
      200: {
        content: {
          'application/json': string;
        };
      };
      /** @description Validation Error */
      422: {
        content: {
          'application/json': components['schemas']['HTTPValidationError'];
        };
      };
    };
  };
  get_system_info_system_info_get: {
    /** Get System Info */
    responses: {
      /** @description Successful Response */
      200: {
        content: {
          'application/json': string;
        };
      };
    };
  };
  get_request_statistics_system_request_statistics_get: {
    /** Get Request Statistics */
    responses: {
      /** @description Successful Response */
      200: {
        content: {
          'application/json': string;
        };
      };
    };
  };
  get_proxy_logs_system_proxy_logs_post: {
    /** Get Proxy Logs */
    requestBody?: {
      content: {
        'application/json': components['schemas']['LogFilterOptions'];
      };
    };
    responses: {
      /** @description Successful Response */
      200: {
        content: {
          'application/json': string;
        };
      };
      /** @description Validation Error */
      422: {
        content: {
          'application/json': components['schemas']['HTTPValidationError'];
        };
      };
    };
  };
  get_server_logs_system_server_logs_post: {
    /** Get Server Logs */
    requestBody?: {
      content: {
        'application/json': components['schemas']['LogFilterOptions'];
      };
    };
    responses: {
      /** @description Successful Response */
      200: {
        content: {
          'application/json': string;
        };
      };
      /** @description Validation Error */
      422: {
        content: {
          'application/json': components['schemas']['HTTPValidationError'];
        };
      };
    };
  };
  get_server_status_status_get: {
    /**
     * Get Server Status
     * @description 普通用户获取服务器状态
     */
    responses: {
      /** @description Successful Response */
      200: {
        content: {
          'application/json': string;
        };
      };
    };
  };
}
```

## File: frontend/src/types/schema.ts
```typescript
import { components } from './openapi';

export type UserRead = components['schemas']['UserRead'];
export type UserCreate = components['schemas']['UserCreate'];
export type UserUpdate = components['schemas']['UserUpdate'];
export type ConversationSchema = components['schemas']['ConversationSchema'];
export type ServerStatusSchema = components['schemas']['ServerStatusSchema'];
export type LimitSchema = components['schemas']['LimitSchema'];
export type ChatStatus = components['schemas']['ChatStatus'];
export type ChatModels = components['schemas']['ChatModels'];

export type SystemInfo = components['schemas']['SystemInfo'];
export type RequestStatistics = components['schemas']['RequestStatistics'];

export type LogFilterOptions = components['schemas']['LogFilterOptions'];

export const chatStatusMap = {
  asking: 'commons.askingChatStatus',
  queueing: 'commons.queueingChatStatus',
  idling: 'commons.idlingChatStatus',
};
```

## File: frontend/src/utils/auth.ts
```typescript
import { hasCookie, removeCookie } from '@/utils/cookies';
// import { useUserStore } from '@/store';

const COOKIE_KEY = 'user_auth';

const hasLoginCookie = () => {
  // const userStore = useUserStore();
  return !!hasCookie(COOKIE_KEY);
};

const clearCookie = () => {
  removeCookie(COOKIE_KEY);
};

export { clearCookie,hasLoginCookie };
```

## File: frontend/src/utils/conversation.ts
```typescript
import { useConversationStore } from '@/store';
import { ChatConversationDetail, ChatMessage } from '@/types/custom';

// 使用以下函数前需要确保调用了 conversationStore.fetchConversationHistory

export function getConvMessageListFromId(conversation_id: string | null) {
  const conversationStore = useConversationStore();
  const result = [];
  if (!conversation_id) return [];
  const conv: ChatConversationDetail = conversationStore.conversationDetailMap[conversation_id];
  if (conv) {
    let x = conv.current_node as any;
    while (x) {
      if (conv.mapping[x].message) result.push(conv.mapping[x]);
      x = conv.mapping[x].parent;
    }
    result.reverse();
  }
  return result;
}

export const getModelNameFromConv = (conv: ChatConversationDetail): string | null => {
  let result = null;
  let current_node = conv.current_node as any;
  while (current_node) {
    const node = conv.mapping[current_node];
    if (node.model_slug) {
      result = node.model_slug;
      break;
    }
    current_node = node.parent;
  }
  return result;
};

export const getModelNameFromMessages = (messages: Array<ChatMessage>): string | null => {
  let result = null;
  for (let i = messages.length - 1; i >= 0; i--) {
    if (messages[i].model_slug) {
      result = messages[i].model_slug || null;
      break;
    }
  }
  return result;
};

export const getModelNameFromConvId = (conversation_id: string | null): string | null => {
  if (!conversation_id) return null;
  const conversationStore = useConversationStore();
  const conv: ChatConversationDetail = conversationStore.conversationDetailMap[conversation_id];
  if (conv) return getModelNameFromConv(conv);
  else return null;
};
```

## File: frontend/src/utils/cookies.ts
```typescript
// 参考：https://github.com/cmp-cc/vue-cookies/blob/master/vue-cookies.js

const defaultConfig = {
  expires: '1d',
  path: '; path=/',
  domain: '',
  secure: '',
  sameSite: '; SameSite=Lax',
};

export function hasCookie(key: string): boolean {
  return new RegExp(`(?:^|;\\s*)${encodeURIComponent(key).replace(/[-.+*]/g, '\\$&')}\\s*\\=`).test(document.cookie);
}

export function removeCookie(key: string, path: string | null = null, domain: string | null = null): boolean {
  if (!key || !hasCookie(key)) {
    return false;
  }
  document.cookie = `${encodeURIComponent(key)}=; expires=Thu, 01 Jan 1970 00:00:00 GMT${domain ? `; domain=${domain}` : defaultConfig.domain}${
    path ? `; path=${path}` : defaultConfig.path
  }; SameSite=Lax`;
  return true;
}
```

## File: frontend/src/utils/loading.ts
```typescript
import { ref } from 'vue';

export default function useLoading(initValue = false) {
  const loading = ref(initValue);
  const setLoading = (value: boolean) => {
    loading.value = value;
  };
  const toggle = () => {
    loading.value = !loading.value;
  };
  return {
    loading,
    setLoading,
    toggle,
  };
}
```

## File: frontend/src/utils/markdown.ts
```typescript
import markdownItKatex from '@traptitech/markdown-it-katex';
import hljs from 'highlight.js';
import MarkdownIt from 'markdown-it';
import markdownItHighlight from 'markdown-it-highlightjs';

const md = new MarkdownIt({
  html: true,
  linkify: false,
  typographer: true,
})
  .use(markdownItKatex)
  .use(markdownItHighlight, { hljs });

export default md;
```

## File: frontend/src/utils/renders.ts
```typescript
import { MdMore } from '@vicons/ionicons4';
import { NButton, NDropdown, NIcon, NInput, NSelect, SelectOption } from 'naive-ui';
import { h } from 'vue';

import { i18n } from '@/i18n';
import useUserStore from '@/store/modules/user';
import { ChatConversationDetail } from '@/types/custom';
import { ChatModels,ConversationSchema } from '@/types/schema';
import { Dialog } from '@/utils/tips';

const t = i18n.global.t as any;

const modelNameMap = {
  'text-davinci-002-render-sha': t('commons.shaModel'),
  'text-davinci-002-render-paid': t('commons.paidModel'),
  'gpt-4': t('commons.gpt4Model'),
  'gpt-4-mobile': t('commons.gpt4MobileModel'),
};

const getModelNameTrans = (model_name: keyof typeof modelNameMap) => {
  return modelNameMap[model_name] || model_name;
};

const getCountTrans = (count: number): string => {
  return count == -1 ? t('commons.unlimited') : `${count}`;
};

const dropdownRenderer = (
  conversation: ConversationSchema,
  handleDeleteConversation: (conversation_id?: string) => void,
  handleChangeConversationTitle: (conversation_id?: string) => void
) =>
  h(
    NDropdown,
    {
      trigger: 'hover',
      options: [
        {
          label: t('commons.delete'),
          key: 'delete',
          props: {
            onClick: () => handleDeleteConversation(conversation.conversation_id),
          },
        },
        {
          label: t('commons.rename'),
          key: 'rename',
          props: {
            onClick: () => handleChangeConversationTitle(conversation.conversation_id),
          },
        },
      ],
    },
    {
      default: () =>
        h(
          NButton,
          {
            size: 'small',
            quaternary: true,
            circle: true,
          },
          { default: () => h(NIcon, null, { default: () => h(MdMore) }) }
        ),
    }
  );

const popupInputDialog = (title: string, placeholder: string, callback: (inp: string) => Promise<any>, success: () => void, fail: () => void) => {
  let input = '';
  const secondInput: string | undefined = undefined;
  const d = Dialog.info({
    title: title,
    positiveText: t('commons.confirm'),
    negativeText: t('commons.cancel'),
    content: () =>
      h(NInput, {
        onInput: (e) => (input = e),
        autofocus: true,
        placeholder: placeholder,
      }),
    onPositiveClick() {
      d.loading = true;
      return new Promise((resolve) => {
        callback(input)
          .then(() => {
            success();
            resolve(true);
          })
          .catch(() => {
            fail();
            resolve(true);
          })
          .finally(() => {
            d.loading = false;
          });
      });
    },
  });
};

const getAvailableModelOptions = (): SelectOption[] => {
  const userStore = useUserStore();
  const options = [{ label: t('commons.shaModel'), value: 'text-davinci-002-render-sha' }];
  if (userStore.user?.can_use_paid)
    options.push({
      label: t('commons.paidModel'),
      value: 'text-davinci-002-render-paid',
    });
  if (userStore.user?.can_use_gpt4) {
    options.push({ label: t('commons.gpt4Model'), value: 'gpt-4' });
    options.push({ label: t('commons.gpt4MobileModel'), value: 'gpt-4-mobile' });
  };
  return options;
};

const popupNewConversationDialog = (callback: (_conv_title: string, _conv_model: string) => Promise<any>) => {
  let convTitle = '';
  let convModel = '';
  const d = Dialog.info({
    title: t('commons.newConversation'),
    positiveText: t('commons.confirm'),
    negativeText: t('commons.cancel'),
    // content: () =>
    //   h(NInput, { onInput: (e) => (input = e), autofocus: true, placeholder: placeholder, }),

    // 用一个 NInput 和一个 NSelect
    content: () =>
      h(
        'div',
        {
          style: 'display: flex; flex-direction: column; gap: 16px;',
        },
        [
          h(NInput, {
            onInput: (e) => (convTitle = e),
            autofocus: true,
            placeholder: t('tips.conversationTitle'),
          }),
          h(NSelect, {
            placeholder: t('tips.whetherUsePaidModel'),
            'onUpdate:value': (value: string) => (convModel = value),
            options: getAvailableModelOptions(),
          }),
        ]
      ),
    onPositiveClick() {
      d.loading = true;
      return new Promise((resolve) => {
        callback(convTitle, convModel)
          .then(() => {
            resolve(true);
          })
          .catch(() => {
            resolve(true);
          })
          .finally(() => {
            d.loading = false;
          });
      });
    },
  });
};

const popupChangeConversationTitleDialog = (
  conversation_id: string,
  callback: (title: string) => Promise<any>,
  success: () => void,
  fail: () => void
) => {
  popupInputDialog(t('commons.rename'), t('tips.rename'), callback, success, fail);
};

const popupResetUserPasswordDialog = (callback: (password: string) => Promise<any>, success: () => void, fail: () => void) => {
  popupInputDialog(t('commons.resetPassword'), t('tips.resetPassword'), callback, success, fail);
};

export {
  dropdownRenderer,
  getCountTrans,
  getModelNameTrans,
  modelNameMap,
  popupChangeConversationTitleDialog,
  popupNewConversationDialog,
  popupResetUserPasswordDialog,
};
```

## File: frontend/src/utils/tips.ts
```typescript
import { useStorage } from '@vueuse/core';
import { ConfigProviderProps, createDiscreteApi, darkTheme, lightTheme } from 'naive-ui';
import { computed,ref } from 'vue';

const themeRef = ref<'light' | 'dark'>(useStorage('theme', 'light').value as any);
const configProviderPropsRef = computed<ConfigProviderProps>(() => ({
  theme: themeRef.value === 'light' ? lightTheme : darkTheme,
}));

const { message, notification, dialog, loadingBar } = createDiscreteApi(['message', 'dialog', 'notification', 'loadingBar'], {
  configProviderProps: configProviderPropsRef,
});

export { dialog as Dialog, loadingBar as LoadingBar, message as Message, notification as Notification, themeRef };
```

## File: frontend/src/views/admin/components/charts/AskChart.vue
```vue
<template>
  <div class="pr-4">
    <v-chart
      class="h-60"
      :option="option"
      :loading="props.loading"
    />
  </div>
</template>

<script setup lang="ts">
import { BarSeriesOption } from 'echarts';
import { BarChart } from 'echarts/charts';
import {
  BrushComponent,
  DatasetComponent,
  DataZoomComponent,
  GridComponent,
  LegendComponent,
  TitleComponent,
  ToolboxComponent,
  TooltipComponent,
} from 'echarts/components';
import { use } from 'echarts/core';
import { CanvasRenderer } from 'echarts/renderers';
import { computed, ref } from 'vue';
import VChart from 'vue-echarts';
import { useI18n } from 'vue-i18n';

import { useAppStore } from '@/store';
import { ToolTipFormatterParams } from '@/types/echarts';
import { UserRead } from '@/types/schema';

import { timeFormatter } from './helpers';
const { t } = useI18n();
const appStore = useAppStore();

use([
  TitleComponent,
  CanvasRenderer,
  BarChart,
  GridComponent,
  TooltipComponent,
  LegendComponent,
  DatasetComponent,
  DataZoomComponent,
  ToolboxComponent,
  BrushComponent,
]);

type AskRecord = [[number, string, number, number], number];

// provide(THEME_KEY, appStore.theme);
interface StatRecord {
  timestamp: number;
  count: number;
  sumAskDuration: number;
  sumTotalDuration: number;
  userIds: number[];
}

const props = defineProps<{
  loading: boolean;
  askRecords?: AskRecord[];
  users?: UserRead[];
}>();

function makeDataset(askRecords: AskRecord[]) {
  // 获得最早的时间戳
  const earliestTimestamp = askRecords.reduce((min, record) => Math.min(min, record[1]), Number.MAX_VALUE);

  const latestTimestamp = askRecords.reduce((max, record) => Math.max(max, record[1]), Number.MIN_VALUE);

  // 对齐到整点或半点
  const alignedEarliestTimestamp = Math.floor(earliestTimestamp / 1800) * 1800 * 1000;
  const alignedLatestTimestamp = Math.ceil(latestTimestamp / 1800) * 1800 * 1000;

  // 数据分类
  const otherRecords: AskRecord[] = [];
  const gpt4Records: AskRecord[] = [];

  askRecords.forEach((record) => {
    if (record[0][1] === 'gpt-4' || record[0][1] === 'gpt-4-mobile') {
      gpt4Records.push(record);
    } else {
      otherRecords.push(record);
    }
  });

  // 计算统计数据
  function calculateStats(records: AskRecord[]): StatRecord[] {
    const stats: StatRecord[] = [];
    let currentTimestamp = alignedEarliestTimestamp;
    // console.log('currentTimestamp', currentTimestamp, new Date(currentTimestamp).toLocaleString())
    while (currentTimestamp < alignedLatestTimestamp) {
      const recordsInInterval = records.filter((record) => record[1] * 1000 >= currentTimestamp && record[1] * 1000 < currentTimestamp + 1800 * 1000);

      if (recordsInInterval.length > 0) {
        const userIds = new Set(recordsInInterval.map((record) => record[0][0]));
        const stat: StatRecord = {
          timestamp: currentTimestamp,
          count: recordsInInterval.length,
          sumAskDuration: recordsInInterval.reduce((sum, record) => sum + record[0][2], 0),
          sumTotalDuration: recordsInInterval.reduce((sum, record) => sum + record[0][3], 0),
          userIds: Array.from(userIds),
        };
        stats.push(stat);
      } else {
        const stat: StatRecord = {
          timestamp: currentTimestamp,
          count: 0,
          sumAskDuration: 0,
          sumTotalDuration: 0,
          userIds: [],
        };
        stats.push(stat);
      }

      currentTimestamp += 1800 * 1000;
    }

    return stats;
  }

  const otherStats = calculateStats(otherRecords);
  const gpt4Stats = calculateStats(gpt4Records);

  return [{ source: otherStats }, { source: gpt4Stats }];
}

const dataset = computed(() => {
  if (props.askRecords) {
    return makeDataset(props.askRecords);
  } else {
    return [];
  }
});

const findUsername = (user_id: number) => {
  const user = props.users?.find((u) => u.id === user_id);
  return user?.username || user_id;
};

const isDark = computed(() => appStore.theme === 'dark');

const generateSeries = (name: string, lineColor: string, itemBorderColor: string, datasetIndex: number): BarSeriesOption => {
  return {
    type: 'bar',
    name,
    datasetIndex,
    yAxisIndex: 0,
    encode: {
      x: 'timestamp',
      y: 'count',
    },
    stack: 'total',
    itemStyle: {
      color: lineColor,
    },
    emphasis: {
      focus: 'series',
      itemStyle: {
        color: lineColor,
        borderWidth: 2,
        borderColor: itemBorderColor,
      },
    },
  };
};

const showDataZoom = ref(false);
const dataZoomOption = computed(() => {
  return showDataZoom.value
    ? [
      {
        type: 'slider',
        show: showDataZoom.value,
        xAxisIndex: 0,
        start: 0,
        end: 100,
        filterMode: 'filter',
      },
    ]
    : [];
});
const gridBottom = computed(() => {
  return showDataZoom.value ? '25%' : '5%';
});

const option = computed(() => {
  return {
    title: {
      text: t('commons.askRequestsCount'),
      left: 'center',
      top: '2.6%',
      textStyle: {
        color: isDark.value ? '#DDD' : '#4E5969',
        fontSize: 16,
        fontWeight: 500,
      },
    },
    grid: {
      left: '2.6%',
      right: '4',
      top: '40',
      bottom: gridBottom.value,
      containLabel: true,
    },
    dataset: dataset.value,
    xAxis: {
      type: 'time',
      axisLabel: {
        color: '#4E5969',
        formatter: (val: any) => timeFormatter(val, false),
        hideOverlap: true,
      },
      axisLine: {
        show: false,
      },
      axisTick: {
        show: false,
      },
      splitLine: {
        show: true,
        // interval: (idx: number) => {
        //   if (idx === 0) return false;
        //   if (idx === xAxis.value.length - 1) return false;
        //   return true;
        // },
        lineStyle: {
          // type: 'dashed',
          color: isDark.value ? '#2E2E30' : '#E5E8EF',
        },
      },
      axisPointer: {
        show: true,
        lineStyle: {
          color: '#23ADFF',
          width: 2,
        },
      },
    },
    yAxis: [
      {
        type: 'value',
        position: 'left',
        axisLine: {
          show: false,
        },
        axisLabel: {
          formatter(value: number, idx: number) {
            if (idx === 0) return String(value);
            return `${value}`;
          },
        },
        splitLine: {
          lineStyle: {
            type: 'dashed',
            color: isDark.value ? '#2E2E30' : '#E5E8EF',
          },
        },
      },
    ],
    tooltip: {
      trigger: 'axis',
      formatter(params: any[]) {
        const [el0, el1] = params as ToolTipFormatterParams[];
        const data0 = el0.data as StatRecord;
        const data1 = el1.data as StatRecord;
        return `<div>
                  <span>${timeFormatter(data0.timestamp, true)} ~ ${timeFormatter(data0.timestamp + 1800 * 1000, true)}</span>
                  <br />
                  <span>${el0.seriesName}: ${data0.count}</span> <br />
                  <span>${el1.seriesName}: ${data1.count}</span> <br />
                  <span>${t('commons.normalAskUsers')}: ${data0.userIds.map((id: number) => findUsername(id))}</span> <br />
                  <span>${t('commons.gpt4AskUsers')}: ${data1.userIds.map((id: number) => findUsername(id))}</span> <br />
                  <span>${t('commons.sumOfNormalAskDuration')}: ${data0.sumAskDuration.toFixed(2)} s</span> <br />
                  <span>${t('commons.sumOfGpt4AskDuration')}: ${data1.sumAskDuration.toFixed(2)} s</span> <br />
                </div>`;
      },
      className: 'echarts-tooltip-diy',
    },

    series: [
      generateSeries(t('commons.normalAskCount'), '#9ce6aa', '#E8FFFB', 0),
      generateSeries(t('commons.gpt4AskCount'), '#F77234', '#FFE4BA', 1),
    ],

    toolbox: {
      feature: {
        myDataZoom: {
          show: true,
          title: 'DataZoom',
          icon: 'path://M0,0H12V2H0V0ZM0,14H12V16H0V14ZM0,6H12V8H0V6ZM0,10H12V12H0V10Z',
          onclick: () => {
            showDataZoom.value = !showDataZoom.value;
          },
        },
        restore: {},
        saveAsImage: {},
      },
    },
    dataZoom: dataZoomOption.value,
  };
});

// watchEffect(() => {
console.log('props', props.askRecords);
// console.log('xAxis', xAxis.value);
// console.log('totalRequestsCountData', totalRequestsCountData.value);
// console.log('datasetSource', datasetSource.value);
// console.log('users', props.users)
//   console.log('option', option.value);
// });
</script>
```

## File: frontend/src/views/admin/components/charts/helpers.ts
```typescript
export const timeFormatter = (value: number, withYear: boolean) => {
  const date = new Date(value);
  const year = date.getFullYear();
  const month = (date.getMonth() + 1).toString().padStart(2, '0');
  const day = date.getDate().toString().padStart(2, '0');
  const hour = date.getHours().toString().padStart(2, '0');
  const minute = date.getMinutes().toString().padStart(2, '0');
  return (withYear ? `${year}-` : '') + `${month}-${day} ${hour}:${minute}`;
};
```

## File: frontend/src/views/admin/components/charts/RequestsChart.vue
```vue
<template>
  <div class="pr-4">
    <v-chart
      class="h-35"
      :option="option"
      :loading="props.loading"
    />
  </div>
</template>

<script setup lang="ts">
import { LineSeriesOption } from 'echarts';
import { LineChart } from 'echarts/charts';
import {
  BrushComponent,
  DatasetComponent,
  DataZoomComponent,
  GridComponent,
  LegendComponent,
  TitleComponent,
  ToolboxComponent,
  // GraphicComponent,
  TooltipComponent,
} from 'echarts/components';
import { use } from 'echarts/core';
import { CanvasRenderer } from 'echarts/renderers';
import { computed, ref, watchEffect } from 'vue';
import VChart from 'vue-echarts';
import { useI18n } from 'vue-i18n';

import { useAppStore } from '@/store';
import { ToolTipFormatterParams } from '@/types/echarts';
import { UserRead } from '@/types/schema';

import { timeFormatter } from './helpers';
const { t } = useI18n();
const appStore = useAppStore();

use([
  TitleComponent,
  CanvasRenderer,
  LineChart,
  GridComponent,
  // GraphicComponent,
  TooltipComponent,
  LegendComponent,
  DatasetComponent,
  DataZoomComponent,
  ToolboxComponent,
  BrushComponent,
]);

// provide(THEME_KEY, appStore.theme);

const props = defineProps<{
  loading: boolean;
  requestCounts?: Record<string, [number, number[]]>; // list of [timestage, [total, user_id_list]]
  requestCountsInterval?: number;
  users?: UserRead[];
}>();

const findUsername = (user_id: number) => {
  const user = props.users?.find((u) => u.id === user_id);
  return user?.username || user_id;
};

const isDark = computed(() => appStore.theme === 'dark');

const datasetSource = computed(() => {
  const data = props.requestCounts
    ? Object.keys(props.requestCounts).map((key) => {
      const timestamp = parseInt(key) * 1000 * props.requestCountsInterval!;
      const count = props.requestCounts![key][0];
      const userIds = props.requestCounts![key][1] as number[];
      // const userString = userIds.map((i) => `${i}`).join(', ');
      return {
        timestamp,
        count,
        userIds,
      };
    })
    : [];
  return data;
});

const generateSeries = (name: string, lineColor: string, itemBorderColor: string): LineSeriesOption => {
  return {
    type: 'line',
    name,
    encode: {
      x: 'timestamp',
      y: 'count',
    },
    stack: 'Total',
    smooth: true,
    symbol: 'circle',
    symbolSize: 10,
    itemStyle: {
      color: lineColor,
    },
    emphasis: {
      focus: 'series',
      itemStyle: {
        color: lineColor,
        borderWidth: 2,
        borderColor: itemBorderColor,
      },
    },
    lineStyle: {
      width: 2,
      color: lineColor,
    },
    showSymbol: false,
    areaStyle: {
      opacity: 0.1,
      color: lineColor,
    },
  };
};

const showDataZoom = ref(false);
const dataZoomOption = computed(() => {
  return showDataZoom.value
    ? [
      {
        type: 'slider',
        show: showDataZoom.value,
        xAxisIndex: 0,
        start: 0,
        end: 100,
        filterMode: 'filter',
      },
    ]
    : [];
});
const gridBottom = computed(() => {
  return showDataZoom.value ? '35%' : '5%';
});

const option = computed(() => {
  return {
    title: {
      text: t('commons.totalRequestsCount'),
      left: 'center',
      top: '2.6%',
      textStyle: {
        color: isDark.value ? '#DDD' : '#4E5969',
        fontSize: 16,
        fontWeight: 500,
      },
    },
    grid: {
      left: '2.6%',
      right: '4',
      top: '30',
      bottom: gridBottom.value,
      containLabel: true,
    },
    dataset: {
      source: datasetSource.value,
    },
    xAxis: {
      type: 'time',
      axisLabel: {
        color: '#4E5969',
        formatter: (val: any) => timeFormatter(val, false),
        hideOverlap: true,
      },
      axisLine: {
        show: false,
      },
      axisTick: {
        show: false,
      },
      splitLine: {
        show: true,
        // interval: (idx: number) => {
        //   if (idx === 0) return false;
        //   if (idx === xAxis.value.length - 1) return false;
        //   return true;
        // },
        lineStyle: {
          type: 'dashed',
          color: isDark.value ? '#2E2E30' : '#E5E8EF',
        },
      },
      axisPointer: {
        show: true,
        lineStyle: {
          color: '#23ADFF',
          width: 2,
        },
      },
    },
    yAxis: {
      type: 'value',
      axisLine: {
        show: false,
      },
      axisLabel: {
        formatter(value: number, idx: number) {
          if (idx === 0) return String(value);
          return `${value}`;
        },
      },
      splitLine: {
        lineStyle: {
          type: 'dashed',
          color: isDark.value ? '#2E2E30' : '#E5E8EF',
        },
      },
    },
    tooltip: {
      trigger: 'axis',
      formatter(params: any[]) {
        const [el] = params as ToolTipFormatterParams[];
        const data = el.data as any;
        return `<div>
                  <span>${timeFormatter(data.timestamp, true)} ~ ${timeFormatter(data.timestamp + props.requestCountsInterval! * 1000, true)}</span>
                  <br />
                  <span>${el.seriesName}: ${data.count}</span> <br />
                  <span>${t('commons.requestUsers')}: ${data.userIds.map((id: number) => findUsername(id))}</span>
                </div>`;
      },
      className: 'echarts-tooltip-diy',
    },

    series: [generateSeries(t('commons.totalRequestsCount'), '#3469FF', '#E8F3FF')],

    toolbox: {
      feature: {
        myDataZoom: {
          show: true,
          title: 'DataZoom',
          icon: 'path://M0,0H12V2H0V0ZM0,14H12V16H0V14ZM0,6H12V8H0V6ZM0,10H12V12H0V10Z',
          onclick: () => {
            showDataZoom.value = !showDataZoom.value;
          },
        },
        restore: {},
        saveAsImage: {},
      },
    },
    dataZoom: dataZoomOption.value,
    // brush: {
    //   xAxisIndex: 0,
    //   throttleDelay: 300,
    //   brushType: 'lineX',
    //   brushMode: 'single',
    //   rangeMode: ['percent', 'percent'],
    //   outOfBrush: {
    //     colorAlpha: 0.1
    //   },
    // },
  };
});

watchEffect(() => {
  // console.log('props', props.requestCounts);
  // console.log('xAxis', xAxis.value);
  // console.log('totalRequestsCountData', totalRequestsCountData.value);
  // console.log('datasetSource', datasetSource.value);
  // console.log('users', props.users)
  console.log('option', option.value);
});
</script>
```

## File: frontend/src/views/admin/components/EditLimitForm.vue
```vue
<template>
  <n-form
    label-placement="left"
    label-width="auto"
    :style="{
      maxWidth: '640px',
    }"
  >
    <n-form-item
      :label="t('commons.canUsePaidModel')"
      path="can_use_paid"
    >
      <n-switch
        v-model:value="limit.can_use_paid"
        placeholder=""
      />
    </n-form-item>
    <n-form-item
      :label="t('commons.canUseGPT4Model')"
      path="can_use_paid"
    >
      <n-switch
        v-model:value="limit.can_use_gpt4"
        placeholder=""
      />
    </n-form-item>
    <n-form-item
      :label="t('commons.maxConversationCount')"
      path="max_conv_count"
    >
      <n-input-number
        v-model:value="limit.max_conv_count"
        :parse="parseValue"
        :format="formatValue"
      />
    </n-form-item>
    <n-form-item
      :label="t('commons.availableAskCount')"
      path="available_ask_count"
    >
      <n-input-number
        v-model:value="limit.available_ask_count"
        :parse="parseValue"
        :format="formatValue"
      />
    </n-form-item>
    <n-form-item
      :label="t('commons.availableGPT4AskCount')"
      path="available_gpt4_ask_count"
    >
      <n-input-number
        v-model:value="limit.available_gpt4_ask_count"
        :parse="parseValue"
        :format="formatValue"
      />
    </n-form-item>
  </n-form>
</template>

<script setup lang="ts">
import { computed } from 'vue';

import { i18n } from '@/i18n';
import { LimitSchema } from '@/types/schema';

const t = i18n.global.t as any;

const props = defineProps<{
  limit: LimitSchema;
}>();

const emits = defineEmits(['update:limit']);

const limit = computed({
  get: () => props.limit,
  set: (value) => {
    emits('update:limit', value);
  },
});

const formatValue = (value: number | null) => (value == -1 ? t('commons.unlimited') : value);
const parseValue = (value: string) => (value == t('commons.unlimited') ? -1 : parseInt(value));

</script>
```

## File: frontend/src/views/admin/components/EditUserForm.vue
```vue
<template>
  <!-- user register form -->
  <n-form
    ref="formRef"
    :model="props.user"
    :rules="rules"
    :label-col="{ span: 8 }"
    :wrapper-col="{ span: 16 }"
  >
    <n-form-item
      :label="t('commons.username')"
      path="username"
    >
      <n-input
        v-model:value="user.username"
        placeholder=""
      />
    </n-form-item>
    <n-form-item
      :label="t('commons.nickname')"
      path="nickname"
    >
      <n-input
        v-model:value="user.nickname"
        placeholder=""
      />
    </n-form-item>
    <n-form-item
      :label="t('commons.password')"
      path="password"
    >
      <n-input
        v-model:value="user.password"
        placeholder=""
      />
    </n-form-item>
    <n-form-item
      :label="t('commons.email')"
      path="email"
    >
      <n-input
        v-model:value="user.email"
        placeholder=""
      />
    </n-form-item>
  </n-form>
</template>

<script setup lang="ts">
import { computed } from 'vue';

import { i18n } from '@/i18n';
import { UserCreate } from '@/types/schema';
const t = i18n.global.t as any;

const props = defineProps<{
  user: UserCreate;
}>();

const emits = defineEmits(['update:user']);

const user = computed({
  get: () => props.user,
  set: (value) => {
    emits('update:user', value);
  },
});

const rules = {
  username: { required: true, message: t('tips.pleaseEnterUsername'), trigger: 'blur' },
  password: { required: true, message: t('tips.pleaseEnterPassword'), trigger: 'blur' },
  email: { required: true, message: t('tips.pleaseEnterEmail'), trigger: 'blur' },
  nickname: { required: true, message: t('tips.pleaseEnterNickname'), trigger: 'blur' },
};

</script>
```

## File: frontend/src/views/admin/components/StatisticsCard.vue
```vue
<template>
  <n-card
    :title="t('commons.statisticsInfo')"
    :content-style="{ padding: '0px' }"
    :header-style="{ paddingBottom: 0 }"
  >
    <RequestsChart
      :users="users"
      :loading="loading"
      :request-counts-interval="requestCountsInterval"
      :request-counts="requestCounts"
    />

    <AskChart
      :loading="loading"
      :ask-records="askRecords"
      :users="users"
    />
  </n-card>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useI18n } from 'vue-i18n';

import { RequestStatistics, UserRead } from '@/types/schema';

import AskChart from './charts/AskChart.vue';
import RequestsChart from './charts/RequestsChart.vue';
const { t } = useI18n();

const props = defineProps<{
  requestStatistics?: RequestStatistics;
  users?: UserRead[];
}>();

const loading = computed(() => {
  return !props.requestStatistics;
});

const requestCountsInterval = computed(() => {
  return props.requestStatistics?.request_counts_interval;
});

const requestCounts = computed(() => {
  return props.requestStatistics?.request_counts as any;
});

const askRecords = computed<any>(() => {
  return props.requestStatistics?.ask_records;
});
</script>
```

## File: frontend/src/views/admin/components/SystemInfoCard.vue
```vue
<template>
  <n-card>
    <template #header>
      <div class="flex flex-row space-x-2">
        <n-text>{{ t('commons.serverOverview') }}</n-text>
        <n-button
          text
          @click="emits('refresh')"
        >
          <template #icon>
            <n-icon>
              <RefreshFilled />
            </n-icon>
          </template>
        </n-button>
      </div>
    </template>
    <div class="grid grid-cols-3 md:grid-cols-5 gap-4">
      <n-statistic
        v-for="item in statistics"
        :key="item.label"
        :label="item.label"
        :value="item.value"
      >
        <template
          v-if="item.prefixIcon"
          #prefix
        >
          <n-icon :component="item.prefixIcon" />
        </template>
        <template #suffix>
          <n-text>
            {{ item.suffix }}
          </n-text>
        </template>
      </n-statistic>
    </div>
  </n-card>
</template>

<script setup lang="ts">
import { RefreshFilled } from '@vicons/material';
import { computed } from 'vue';
import { useI18n } from 'vue-i18n';

import { ServerStatusSchema, SystemInfo } from '@/types/schema';
const { t } = useI18n();

const props = defineProps<{
  systemInfo?: SystemInfo;
  serverStatus?: ServerStatusSchema;
}>();

const emits = defineEmits<{
  (e: 'refresh'): void;
}>();

function hoursSince(timestamp?: number) {
  if (!timestamp) {
    return 'N/A';
  }
  const now = new Date();
  const diff = now.getTime() - timestamp * 1000; // 将 Unix 时间戳转换为毫秒
  const hours = diff / 1000 / 3600; // 将毫秒转换为小时
  return hours.toFixed(1); // 保留一位小数
}

const statistics = computed(() => {
  return [
    {
      label: t('commons.userCountAndOnlineCount'),
      value: props.serverStatus?.active_user_in_5m,
      prefixIcon: null,
      suffix: `/ ${props.systemInfo?.total_user_count}`,
    },
    {
      label: t('commons.conversationCount'),
      value: props.systemInfo?.valid_conversation_count,
      prefixIcon: null,
      suffix: `/ ${props.systemInfo?.total_conversation_count}`,
    },
    {
      label: t('commons.chatbotStatus'),
      value: props.serverStatus?.is_chatbot_busy ? t('commons.askingChatStatus') : t('commons.idlingChatStatus'),
      prefixIcon: null,
    },
    {
      label: t('commons.chatbotWaitingCount'),
      value: props.serverStatus?.chatbot_waiting_count,
      prefixIcon: null,
    },
    {
      label: t('commons.startUpDuration'),
      value: hoursSince(props.systemInfo?.startup_time),
      prefixIcon: null,
      suffix: ' h',
    },
  ];
});
</script>
```

## File: frontend/src/views/admin/components/UserSelector.vue
```vue
<template>
  <n-auto-complete
    v-model:value="value"
    :get-show="getShow"
    :options="options"
    :placeholder="t('commons.chooseUser')"
    @update:value="update"
  />
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';

import { getAllUserApi } from '@/api/user';
import { i18n } from '@/i18n';
import { UserRead } from '@/types/schema';
const t = i18n.global.t as any;

const data = ref<Array<UserRead>>([]);
const value = ref<string | null>(null);

const emits = defineEmits(['update:value']);

const getShow = (_option: any) => true;

const update = (value: string | null) => {
  emits('update:value', value);
};

getAllUserApi().then((res) => {
  data.value = res.data;
});

const options = computed(() => {
  return data.value.map((item) => {
    return {
      label: item.username,
      value: item.username,
    };
  });
});
</script>
```

## File: frontend/src/views/admin/conversation_manager.vue
```vue
<template>
  <div>
    <div class="mb-4 mt-1 ml-1 flex flex-row justify-between space-x-2">
      <div class="flex flex-row space-x-4">
        <n-button
          circle
          @click="refreshData"
        >
          <template #icon>
            <n-icon>
              <RefreshFilled />
            </n-icon>
          </template>
        </n-button>
        <div
          v-show="checkedRowKeys.length !== 0"
          class="space-x-2"
        >
          <n-button
            type="warning"
            secondary
            @click="handleInvalidateConversations"
          >
            <template #icon>
              <n-icon>
                <EmojiFlagsFilled />
              </n-icon>
            </template>
            {{ $t('commons.invalidateConversation') }}
          </n-button>
          <n-button
            type="error"
            secondary
            @click="handleVanishConversations"
          >
            <template #icon>
              <n-icon>
                <TrashOutline />
              </n-icon>
            </template>
            {{ $t('commons.vanishConversation') }}
          </n-button>
          <n-button
            type="info"
            secondary
            @click="handleAssignConversations"
          >
            <template #icon>
              <n-icon>
                <PersonAddAlt1Filled />
              </n-icon>
            </template>
            {{ $t('commons.chooseUserToAssign') }}
          </n-button>
        </div>
      </div>
      <div class="space-x-2">
        <n-button @click="handleVanishAllInvalidConversations">
          {{ $t('commons.deleteInvalidConversations') }}
        </n-button>
        <n-button
          type="error"
          @click="handleClearAllConversations"
        >
          {{ $t('commons.clearAllConversations') }}
        </n-button>
      </div>
    </div>
    <n-data-table
      v-model:checked-row-keys="checkedRowKeys"
      size="small"
      :columns="columns"
      :data="data"
      :bordered="true"
      :pagination="{
        pageSize: 20,
      }"
      :row-key="rowKey"
    />
  </div>
</template>

<script setup lang="ts">
import { TrashOutline } from '@vicons/ionicons5';
import { EmojiFlagsFilled, PersonAddAlt1Filled, RefreshFilled } from '@vicons/material';
import type { DataTableColumns } from 'naive-ui';
import { NButton, NIcon, NTooltip } from 'naive-ui';
import { h, ref } from 'vue';
import { useI18n } from 'vue-i18n';
import { useRouter } from 'vue-router';

import {
  assignConversationToUserApi,
  clearAllConversationApi,
  deleteConversationApi,
  getAllConversationsApi,
  vanishConversationApi,
} from '@/api/chat';
import { ConversationSchema } from '@/types/schema';
import { getModelNameTrans } from '@/utils/renders';
import { Dialog, Message } from '@/utils/tips';

import UserSelector from './components/UserSelector.vue';
const { t } = useI18n();
const router = useRouter();
const data = ref<Array<ConversationSchema>>([]);
const rowKey = (row: ConversationSchema) => row.conversation_id;
const checkedRowKeys = ref<Array<string>>([]);

const refreshData = () => {
  getAllConversationsApi(true).then((res) => {
    data.value = res.data;
  });
};

refreshData();

const columns: DataTableColumns<ConversationSchema> = [
  {
    type: 'selection',
  },
  {
    title: '#',
    key: 'id',
    sorter: 'default',
  },
  {
    title: 'UUID',
    key: 'conversation_id',
    render: (row) => {
      return h(
        NTooltip,
        { trigger: 'hover' },
        {
          trigger: () => row.conversation_id?.substring(0, 4),
          default: () => row.conversation_id,
        }
      );
    },
  },
  {
    title: t('commons.title'),
    key: 'title',
    sorter: 'default',
    render: (row) => {
      return h(
        NButton,
        {
          text: true,
          tag: 'a',
          href: router.resolve({
            name: 'conversationHistory',
            params: { conversation_id: row.conversation_id },
          }).href,
          target: '_blank',
        },
        {
          default: () => (row.title ? row.title : t('commons.empty')),
          // }
        }
      );
    },
  },
  {
    title: t('commons.belongToUser'),
    key: 'user_id',
    render: (row) => {
      return row.user_id ? row.user_id : t('commons.empty');
    },
    sorter: 'default',
  },
  {
    title: t('commons.createTime'),
    key: 'create_time',
    defaultSortOrder: 'descend',
    sorter: (a, b) => {
      if (!a.create_time || !b.create_time) return 0;
      return new Date(a.create_time!).getTime() - new Date(b.create_time!).getTime();
    },
    render: (row) => {
      if (!row.create_time) return '';
      return h(
        NTooltip,
        { trigger: 'hover' },
        {
          trigger: () => new Date(row.create_time! + 'Z').toLocaleString(),
          default: () => row.create_time,
        }
      );
    },
  },
  {
    title: t('commons.modelName'),
    key: 'model_name',
    render(row) {
      return row.model_name ? getModelNameTrans(row.model_name) : t('commons.unknown');
    },
    sorter: 'default',
  },
  {
    title: t('commons.isValid'),
    key: 'is_valid',
    render(row) {
      return row.is_valid ? t('commons.yes') : t('commons.no');
    },
    sorter: (a, b) => {
      const val_a = a.is_valid ? 1 : 0;
      const val_b = b.is_valid ? 1 : 0;
      return val_a - val_b;
    },
  },
];

const handleInvalidateConversations = () => {
  const d = Dialog.info({
    title: t('commons.invalidateConversation'),
    content: t('tips.invalidateConversation'),
    positiveText: t('commons.confirm'),
    negativeText: t('commons.cancel'),
    onPositiveClick: () => {
      d.loading = true;
      return new Promise((resolve, reject) => {
        const action = async () => {
          for (const conversation_id of checkedRowKeys.value) {
            await deleteConversationApi(conversation_id);
          }
        };
        action()
          .then(() => {
            Message.success(t('tips.deleteConversationSuccess'));
            refreshData();
            resolve(true);
          })
          .catch((err) => {
            Message.error(t('tips.deleteConversationFailed') + ': ' + err);
            reject(err);
          })
          .finally(() => {
            d.loading = false;
          });
      });
    },
  });
};

const handleVanishConversations = () => {
  const d = Dialog.warning({
    title: t('commons.vanishConversation'),
    content: t('tips.vanishConversation'),
    positiveText: t('commons.confirm'),
    negativeText: t('commons.cancel'),
    onPositiveClick: () => {
      d.loading = true;
      return new Promise((resolve, reject) => {
        const action = async () => {
          for (const conversation_id of checkedRowKeys.value) {
            await vanishConversationApi(conversation_id);
            await new Promise((resolve) => setTimeout(resolve, 200));
          }
        };
        action()
          .then(() => {
            Message.success(t('tips.success'));
            refreshData();
            checkedRowKeys.value = [];
            resolve(true);
          })
          .catch((err) => {
            Message.error(t('tips.failed') + ': ' + err);
            reject(err);
          })
          .finally(() => {
            d.loading = false;
          });
      });
    },
  });
};

const handleAssignConversations = () => {
  let username: string | null = null;
  const d = Dialog.warning({
    title: t('commons.chooseUserToAssign'),
    content: () =>
      h(UserSelector, {
        'onUpdate:value': (val: string | null) => {
          username = val;
        },
      }),
    positiveText: t('commons.confirm'),
    negativeText: t('commons.cancel'),
    onPositiveClick: () => {
      d.loading = true;
      return new Promise((resolve, reject) => {
        if (username === null) {
          Message.error(t('errors.noUserSelected'));
          d.loading = false;
          reject(false);
          return;
        }
        const action = async () => {
          for (const conversation_id of checkedRowKeys.value) {
            await assignConversationToUserApi(conversation_id, username!);
          }
        };
        action()
          .then(() => {
            Message.success(t('tips.success'));
            refreshData();
            checkedRowKeys.value = [];
            resolve(true);
          })
          .catch((err) => {
            Message.error(t('tips.failed') + ': ' + err);
            reject(err);
          })
          .finally(() => {
            d.loading = false;
          });
      });
    },
  });
};

const handleVanishAllInvalidConversations = () => {
  const d = Dialog.info({
    title: t('commons.deleteInvalidConversations'),
    content: t('commons.deleteInvalidConversationsConfirm'),
    positiveText: t('commons.confirm'),
    negativeText: t('commons.cancel'),
    onPositiveClick: () => {
      d.loading = true;
      const action = async () => {
        for (const conversation of data.value) {
          if (!conversation.is_valid) {
            await vanishConversationApi(conversation.conversation_id!);
            await new Promise((resolve) => setTimeout(resolve, 200));
          }
        }
        data.value = data.value.filter((conversation) => conversation.is_valid);
      };
      return new Promise((resolve, reject) => {
        action()
          .then(() => {
            Message.success(t('tips.deleteConversationSuccess'));
            refreshData();
            checkedRowKeys.value = [];
            resolve(true);
          })
          .catch((err) => {
            Message.error(t('tips.deleteConversationFailed'));
            reject();
          })
          .finally(() => {
            d.loading = false;
          });
      });
    },
  });
};

const handleClearAllConversations = () => {
  const d = Dialog.error({
    title: t('commons.clearAllConversations'),
    content: t('commons.readyToClearAllConversations'),
    positiveText: t('commons.confirm'),
    negativeText: t('commons.cancel'),
    onPositiveClick: () => {
      d.loading = true;
      return new Promise((resolve, reject) => {
        clearAllConversationApi()
          .then(() => {
            Message.success(t('tips.deleteConversationSuccess'));
            refreshData();
            checkedRowKeys.value = [];
            resolve(true);
          })
          .catch((err) => {
            Message.error(t('tips.deleteConversationFailed'));
            reject();
          })
          .finally(() => {
            d.loading = false;
          });
      });
    },
  });
};
</script>
```

## File: frontend/src/views/admin/index.vue
```vue
<template>
  <n-space
    vertical
    class="-ml-2 h-full px-2"
  >
    <n-layout
      has-sider
      class="h-90vh"
    >
      <n-layout-sider
        bordered
        :collapsed="collapsed"
        collapse-mode="width"
        :collapsed-width="64"
        :width="200"
        show-trigger
        @collapse="collapsed = true"
        @expand="collapsed = false"
      >
        <n-menu
          v-model:value="activeKey"
          :collapsed-width="64"
          :collapsed-icon-size="22"
          :options="menuOptions"
        />
      </n-layout-sider>
      <n-layout class="ml-4 mr-2">
        <n-scrollbar>
          <router-view v-slot="{ Component, route }">
            <keep-alive>
              <component
                :is="Component"
                :key="route.fullPath"
              />
            </keep-alive>
          </router-view>
        </n-scrollbar>
      </n-layout>
    </n-layout>
  </n-space>
</template>

<script setup lang="ts">
import { ChatbubbleEllipses, FileTrayFull, InformationCircle } from '@vicons/ionicons5';
import { SupervisedUserCircleRound } from '@vicons/material';
import { NIcon } from 'naive-ui';
import { h, ref, watch } from 'vue';
import { useI18n } from 'vue-i18n';
import { useRouter } from 'vue-router';
const { t } = useI18n();
const router = useRouter();

const collapsed = ref(true);
const activeKey = ref<string>(router.currentRoute.value.name as string);

function renderIcon(icon: any) {
  return () => h(NIcon, null, { default: () => h(icon) });
}

const menuOptions = [
  {
    label: t('commons.systemManagement'),
    key: 'systemManagement',
    icon: renderIcon(InformationCircle),
  },
  {
    label: t('commons.userManagement'),
    key: 'userManagement',
    icon: renderIcon(SupervisedUserCircleRound),
  },
  {
    label: t('commons.conversationManagement'),
    key: 'conversationManagement',
    icon: renderIcon(ChatbubbleEllipses),
  },
  {
    label: t('commons.logViewer'),
    key: 'logViewer',
    icon: renderIcon(FileTrayFull),
  },
];

watch(
  async () => activeKey.value,
  (_newName: any) => {
    router.push({ name: activeKey.value });
  }
);
</script>
```

## File: frontend/src/views/admin/log_viewer.vue
```vue
<template>
  <div class="mb-4 flex flex-col">
    <n-tabs
      v-model:value="tab"
      type="segment"
    >
      <n-tab name="server">
        {{ t('commons.serverLogs') }}
      </n-tab>
      <!-- <n-tab name="proxy">
        {{ t('commons.proxyLogs') }}
      </n-tab> -->
    </n-tabs>
    <!-- 设置 -->
    <div class="flex flex-row mt-3 justify-between">
      <div class="flex flex-wrap flex-row sm:space-x-3">
        <div class="option-item">
          <n-text>{{ t('commons.maxLineCount') }}</n-text>
          <n-input-number
            v-model:value="maxLineCount"
            size="small"
            class="w-27"
            :min="100"
            :max="2000"
            :step="100"
          />
        </div>
        <div class="option-item">
          <n-text>{{ t('commons.updateInterval') }}</n-text>
          <n-select
            v-model:value="refresh_duration"
            size="small"
            class="w-20"
            :options="[
              { label: '3s', value: 3 },
              { label: '5s', value: 5 },
              { label: '10s', value: 10 },
            ]"
          />
        </div>
        <div class="option-item">
          <n-text>{{ t('commons.excludeKeywords') }}</n-text>
          <n-dynamic-tags
            v-if="tab === 'proxy'"
            v-model:value="proxyExcludeKeywords"
            size="small"
          />
          <n-dynamic-tags
            v-else
            v-model:value="serverExcludeKeywords"
            size="small"
          />
        </div>
      </div>
      <div class="flex items-center space-x-2">
        <n-text>{{ t('commons.autoScrolling') }}</n-text>
        <n-switch
          v-model:value="enableAutoScroll"
          size="small"
        />
      </div>
    </div>
    <n-card
      class="mt-3 flex-grow h-full"
      :content-style="{ height: '100%' }"
    >
      <n-log
        ref="logInstRef"
        :font-size="10"
        :rows="40"
        :lines="logsContent"
      />
    </n-card>
  </div>
</template>

<script setup lang="ts">
import { nextTick, ref, watch } from 'vue';
import { useI18n } from 'vue-i18n';

import { getProxyLogsApi, getServerLogsApi } from '@/api/system';
import { LogFilterOptions } from '@/types/schema';
const { t } = useI18n();

const refresh_duration = ref(5);
const tab = ref<string>('server');
const logsContent = ref<Array<string>>();
const enableAutoScroll = ref(true);
const maxLineCount = ref(100);
const proxyExcludeKeywords = ref<Array<string>>([]);
const serverExcludeKeywords = ref<Array<string>>(['status', 'logs']);

const logInstRef = ref();

watch(
  () => tab.value,
  () => {
    loadLogs();
  }
);
watch(
  () => maxLineCount.value,
  () => {
    loadLogs();
  }
);

const scrollToBottom = () => {
  nextTick(() => {
    logInstRef.value?.scrollTo({ position: 'bottom', slient: false });
  });
};

const loadLogs = () => {
  if (tab.value === 'server') {
    getServerLogsApi({
      max_lines: maxLineCount.value,
      exclude_keywords: serverExcludeKeywords.value,
    } as LogFilterOptions).then((res) => {
      logsContent.value = res.data;
    });
  } else {
    getProxyLogsApi({
      max_lines: maxLineCount.value,
      exclude_keywords: proxyExcludeKeywords.value,
    } as LogFilterOptions).then((res) => {
      logsContent.value = res.data;
    });
  }
  if (enableAutoScroll.value) {
    scrollToBottom();
  }
};

loadLogs();
let interval = setInterval(() => {
  loadLogs();
}, refresh_duration.value * 1000);

watch(
  () => refresh_duration.value,
  () => {
    clearInterval(interval);
    interval = setInterval(() => {
      loadLogs();
    }, refresh_duration.value * 1000);
  }
);

watch(
  () => serverExcludeKeywords.value,
  () => {
    loadLogs();
  }
);

watch(
  () => proxyExcludeKeywords.value,
  () => {
    loadLogs();
  }
);
</script>

<style>
.option-item {
  @apply flex flex-row space-x-2 items-center mr-1 my-1;
}
</style>
```

## File: frontend/src/views/admin/system_manager.vue
```vue
<template>
  <div class="mb-4 mt-2 flex flex-col space-y-4">
    <SystemInfoCard
      :system-info="systemInfo"
      :server-status="serverStatus"
      @refresh="refreshData"
    />
    <StatisticsCard
      :request-statistics="requestStatistics"
      :users="users"
    />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useI18n } from 'vue-i18n';

import { getServerStatusApi } from '@/api/status';
import { getRequestStatisticsApi, getSystemInfoApi } from '@/api/system';
import { getAllUserApi } from '@/api/user';
import { RequestStatistics, ServerStatusSchema, SystemInfo, UserRead } from '@/types/schema';

import StatisticsCard from './components/StatisticsCard.vue';
import SystemInfoCard from './components/SystemInfoCard.vue';
const { t } = useI18n();

const systemInfo = ref<SystemInfo | undefined>();
const serverStatus = ref<ServerStatusSchema | undefined>();
const requestStatistics = ref<RequestStatistics | undefined>();
const users = ref<UserRead[] | undefined>();

const refreshData = () => {
  getSystemInfoApi().then((res) => {
    systemInfo.value = res.data;
  });

  getServerStatusApi().then((res) => {
    serverStatus.value = res.data;
  });

  getRequestStatisticsApi().then((res) => {
    requestStatistics.value = res.data;
  });
};

refreshData();

getAllUserApi().then((res) => {
  users.value = res.data;
});
</script>
```

## File: frontend/src/views/admin/user_manager.vue
```vue
<template>
  <div>
    <div class="mb-4 mt-1 ml-1 flex flex-row space-x-2 justify-between">
      <n-button
        circle
        @click="refreshData"
      >
        <template #icon>
          <n-icon>
            <RefreshFilled />
          </n-icon>
        </template>
      </n-button>
      <n-button
        type="primary"
        @click="handleAddUser"
      >
        {{ $t('commons.addUser') }}
      </n-button>
    </div>
    <n-data-table
      :scroll-x="1400"
      size="small"
      :columns="columns"
      :data="data"
      :bordered="true"
      :pagination="{
        pageSize: 20,
      }"
    />
  </div>
</template>

<script setup lang="ts">
import { Pencil, TrashOutline } from '@vicons/ionicons5';
import { PasswordRound, RefreshFilled } from '@vicons/material';
import { DataTableColumns, NButton, NIcon } from 'naive-ui';
import { h, ref } from 'vue';
import { useI18n } from 'vue-i18n';

import { deleteUserApi, getAllUserApi, registerApi, resetUserPasswordApi, updateUserLimitApi } from '@/api/user';
import { useUserStore } from '@/store';
import { chatStatusMap,LimitSchema, UserCreate, UserRead } from '@/types/schema';
import { getCountTrans, popupResetUserPasswordDialog } from '@/utils/renders';
import { Dialog, Message } from '@/utils/tips';

import EditLimitForm from './components/EditLimitForm.vue';
import EditUserForm from './components/EditUserForm.vue';

const { t } = useI18n();

const userStore = useUserStore();

const data = ref<Array<UserRead>>([]);

const refreshData = () => {
  getAllUserApi().then((res) => {
    data.value = res.data;
    // Message.success(t("tips.refreshed"));
  });
};

getAllUserApi().then((res) => {
  data.value = res.data;
});

const columns: DataTableColumns<UserRead> = [
  {
    title: '#',
    key: 'id',
  },
  {
    title: t('commons.username'),
    key: 'username',
  },
  {
    title: t('commons.nickname'),
    key: 'nickname',
  },
  {
    title: t('commons.status'),
    key: 'chat_status',
    render(row) {
      return row.chat_status ? t(chatStatusMap[row.chat_status as keyof typeof chatStatusMap]) : '';
    },
    sorter: 'default',
  },
  {
    title: t('commons.activeTime'),
    key: 'active_time',
    render(row) {
      return row.active_time ? new Date(row.active_time + 'Z').toLocaleString() : t('commons.neverActive');
    },
    sorter: (a, b) => {
      if (!a.active_time || !b.active_time) return 0;
      return new Date(a.active_time!).getTime() - new Date(b.active_time!).getTime();
    },
  },
  {
    title: t('commons.maxConversationCount'),
    key: 'max_conv_count',
    render(row) {
      return getCountTrans(row.max_conv_count!);
    },
  },
  {
    title: t('commons.availableAskCount'),
    key: 'available_ask_count',
    render(row) {
      return getCountTrans(row.available_ask_count!);
    },
  },
  {
    title: t('commons.availableGPT4AskCount'),
    key: 'available_gpt4_ask_count',
    render(row) {
      return getCountTrans(row.available_gpt4_ask_count!);
    },
  },
  {
    title: t('commons.canUsePaidModel'),
    key: 'can_use_paid',
    render(row) {
      return row.can_use_paid ? t('commons.yes') : t('commons.no');
    },
  },
  {
    title: t('commons.canUseGPT4Model'),
    key: 'can_use_gpt4',
    render(row) {
      return row.can_use_gpt4 ? t('commons.yes') : t('commons.no');
    },
  },
  {
    title: t('commons.email'),
    key: 'email',
  },
  {
    title: t('commons.isSuperuser'),
    key: 'is_superuser',
    render(row) {
      return row.is_superuser ? t('commons.yes') : t('commons.no');
    },
  },
  {
    title: t('commons.actions'),
    key: 'actions',
    fixed: 'right',
    render(row) {
      // TODO: 删除、修改密码，两个按钮
      return h(
        'div',
        {
          class: 'flex justify-start space-x-2',
        },
        [
          h(
            NButton,
            {
              size: 'small',
              type: 'error',
              circle: true,
              secondary: true,
              onClick: () => {
                const d = Dialog.warning({
                  title: t('commons.deleteUser'),
                  content: t('tips.deleteUserConfirm'),
                  positiveText: t('commons.confirm'),
                  negativeText: t('commons.cancel'),
                  onPositiveClick: () => {
                    d.loading = true;
                    return new Promise((resolve, reject) => {
                      deleteUserApi(row.id)
                        .then((res) => {
                          Message.success(t('tips.deleteUserSuccess'));
                          getAllUserApi().then((res) => {
                            data.value = res.data;
                          });
                          resolve(true);
                        })
                        .catch((err) => {
                          Message.error(t('tips.deleteUserFailed') + ': ' + err);
                          reject(err);
                        })
                        .finally(() => {
                          d.loading = false;
                        });
                    });
                  },
                });
              },
            },
            {
              icon: () =>
                h(NIcon, null, {
                  default: () => h(TrashOutline),
                }),
            }
          ),
          h(
            NButton,
            {
              size: 'small',
              type: 'info',
              circle: true,
              secondary: true,
              onClick: () => {
                popupResetUserPasswordDialog(
                  async (password: string) => {
                    await resetUserPasswordApi(row.id, password);
                  },
                  () => {
                    Message.info(t('tips.resetUserPasswordSuccess'));
                  },
                  () => {
                    Message.error(t('tips.resetUserPasswordFailed'));
                  }
                );
              },
            },
            {
              icon: () =>
                h(NIcon, null, {
                  default: () => h(PasswordRound),
                }),
            }
          ),
          h(
            NButton,
            {
              size: 'small',
              type: 'primary',
              circle: true,
              secondary: true,
              onClick: handleSetUserLimit(row),
            },
            {
              icon: () =>
                h(NIcon, null, {
                  default: () => h(Pencil),
                }),
            }
          ),
        ]
      );
    },
  },
];

const handleAddUser = () => {
  const user = ref<UserCreate>({
    username: '',
    nickname: '',
    email: '',
    password: '',
    is_superuser: false,
  });
  const d = Dialog.info({
    title: t('commons.addUser'),
    content: () =>
      h(
        EditUserForm,
        {
          user: user.value,
          'onUpdate:user': (newUser: UserCreate) => {
            user.value = newUser;
          },
        },
        { default: () => '' }
      ),
    positiveText: t('commons.confirm'),
    negativeText: t('commons.cancel'),
    onPositiveClick: () => {
      d.loading = true;
      return new Promise((resolve, reject) => {
        registerApi(user.value)
          .then((res) => {
            Message.success(t('commons.addUserSuccess'));
            getAllUserApi().then((res) => {
              data.value = res.data;
            });
            resolve(true);
          })
          .catch((err) => {
            Message.error(t('commons.addUserFailed') + ': ' + err);
            reject(err);
          })
          .finally(() => {
            d.loading = false;
          });
      });
    },
  });
};

const handleSetUserLimit = (user: UserRead) => () => {
  const limit = ref<LimitSchema>({
    max_conv_count: user.max_conv_count,
    available_ask_count: user.available_ask_count,
    can_use_paid: user.can_use_paid,
    can_use_gpt4: user.can_use_gpt4,
    available_gpt4_ask_count: user.available_gpt4_ask_count,
  });
  const d = Dialog.info({
    title: t('commons.setUserLimit'),
    content: () =>
      h(
        EditLimitForm,
        {
          limit: limit.value,
          'onUpdate:limit': (newLimit: LimitSchema) => {
            limit.value = newLimit;
          },
        },
        { default: () => '' }
      ),
    positiveText: t('commons.confirm'),
    negativeText: t('commons.cancel'),
    onPositiveClick: () => {
      d.loading = true;
      return new Promise((resolve, reject) => {
        updateUserLimitApi(user.id, limit.value)
          .then((res) => {
            Message.success(t('commons.setUserLimitSuccess'));
            getAllUserApi().then((res) => {
              data.value = res.data;
            });
            resolve(true);
          })
          .catch((err) => {
            reject(err);
          })
          .finally(() => {
            d.loading = false;
          });
      });
    },
  });
};
</script>
```

## File: frontend/src/views/conversation/components/HistoryContent.vue
```vue
<template>
  <div
    id="print-content"
    ref="contentRef"
    class="flex flex-col h-full"
    tabindex="0"
    style="outline: none"
    @keyup.esc="toggleFullscreenHistory(true)"
  >
    <div v-if="!props.loading">
      <!-- 消息记录 -->
      <div
        class="flex justify-center py-4 px-4 max-w-full relative"
        :style="{ backgroundColor: themeVars.baseColor }"
      >
        <n-text>{{ $t('commons.currentConversationModel') }}: {{ getModelNameTrans(modelName as any) }} </n-text>
        <n-button
          v-if="_fullscreen"
          class="absolute left-4 hide-in-print"
          text
          @click="toggleFullscreenHistory"
        >
          <template #icon>
            <n-icon>
              <Close />
            </n-icon>
          </template>
        </n-button>
      </div>
      <MessageRow
        v-for="message in messages"
        :key="message.id"
        :message="message"
      />
    </div>
    <n-empty
      v-else
      class="h-full flex justify-center"
      :style="{ backgroundColor: themeVars.cardColor }"
      :description="$t('tips.loading')"
    >
      <template #icon>
        <n-spin size="medium" />
      </template>
    </n-empty>
  </div>
</template>

<script setup lang="ts">
import { Close } from '@vicons/ionicons5';
import { useThemeVars } from 'naive-ui';
import { computed, ref, watch } from 'vue';
import { useI18n } from 'vue-i18n';

import { ChatMessage } from '@/types/custom';
import { getModelNameFromMessages } from '@/utils/conversation';
import { getModelNameTrans } from '@/utils/renders';
import { Message } from '@/utils/tips';

import MessageRow from './MessageRow.vue';

const { t } = useI18n();

const themeVars = useThemeVars();

const props = defineProps<{
  messages: ChatMessage[];
  modelName?: string;
  fullscreen: boolean; // 初始状态下是否全屏
  showTips: boolean;
  loading: boolean;
}>();

const contentRef = ref();
const historyContentParent = ref<HTMLElement>();
const _fullscreen = ref(false);

const modelName = computed(() => {
  if (props.modelName) {
    return props.modelName;
  } else {
    return getModelNameFromMessages(props.messages);
  }
});

watch(
  () => props.fullscreen,
  () => {
    toggleFullscreenHistory(props.showTips);
  }
);

const toggleFullscreenHistory = (showTips: boolean) => {
  // fullscreenHistory.value = !fullscreenHistory.value;
  const appElement = document.getElementById('app');
  const bodyElement = document.body;
  const historyContentElement = contentRef.value;
  if (_fullscreen.value) {
    // 将 historyContent 移动回来
    historyContentParent.value?.appendChild(historyContentElement);
    if (appElement) appElement.style.display = 'block';
  } else {
    historyContentParent.value = historyContentElement.parentElement;
    // 移动到body child的第一个
    bodyElement.insertBefore(historyContentElement, bodyElement.firstChild);
    // 将div#app 设置为不可见
    if (appElement) {
      appElement.style.display = 'none';
    }
    historyContentElement.focus();
    if (showTips)
      Message.success(t('tips.pressEscToExitFullscreen'), {
        duration: 2000,
      });
  }
  _fullscreen.value = !_fullscreen.value;
};

if (props.fullscreen) {
  toggleFullscreenHistory(props.showTips);
}

const focus = () => {
  contentRef.value.focus();
};

defineExpose({
  focus,
  toggleFullscreenHistory,
});
</script>
```

## File: frontend/src/views/conversation/components/InputRegion.vue
```vue
<template>
  <div
    class="flex-shrink-0 flex flex-col align-middle relative z-10"
    :style="{ background: themeVars.baseColor }"
  >
    <n-divider />
    <!-- 暂停按钮 -->
    <div class="flex w-full justify-center absolute -top-10">
      <n-button
        v-show="canAbort"
        secondary
        strong
        type="error"
        size="small"
        @click="emits('abort-request')"
      >
        <template #icon>
          <Stop />
        </template>
        {{ t('commons.abortRequest') }}
      </n-button>
    </div>

    <!-- 工具栏 -->
    <div class="mx-2 flex flex-row space-x-2 py-2 justify-center relative">
      <!-- 展开/收起按钮 -->
      <n-button
        class="absolute left-0 top-2"
        quaternary
        circle
        size="small"
        @click="toggleInputExpanded"
      >
        <template #icon>
          <n-icon :component="inputExpanded ? KeyboardDoubleArrowDownRound : KeyboardDoubleArrowUpRound" />
        </template>
      </n-button>
      <!-- 是否启用自动滚动 -->
      <n-tooltip>
        <template #trigger>
          <n-switch
            v-model:value="autoScrolling"
            size="small"
            class="absolute right-2 top-3"
          >
            <template #icon>
              A
            </template>
          </n-switch>
        </template>
        {{ $t('tips.autoScrolling') }}
      </n-tooltip>
      <n-button
        secondary
        type="info"
        size="small"
        @click="emits('show-fullscreen-history')"
      >
        <template #icon>
          <n-icon :size="22">
            <FullscreenRound />
          </n-icon>
        </template>
      </n-button>
      <n-button
        secondary
        type="primary"
        size="small"
        @click="emits('export-to-markdown-file')"
      >
        <template #icon>
          <n-icon>
            <LogoMarkdown />
          </n-icon>
        </template>
      </n-button>
      <n-button
        secondary
        type="warning"
        size="small"
        @click="emits('export-to-pdf-file')"
      >
        <template #icon>
          <n-icon>
            <Print />
          </n-icon>
        </template>
      </n-button>
    </div>
    <!-- 输入框 -->
    <div class="mx-4 mb-4 flex flex-row space-x-2">
      <n-input
        ref="inputRef"
        v-model:value="inputValue"
        class="flex-1"
        type="textarea"
        :bordered="true"
        :placeholder="$t('tips.sendMessage', [appStore.preference.sendKey])"
        :autosize="{ minRows: 1 }"
        :style="inputStyle"
        @keydown="shortcutSendMsg"
      >
        <template #suffix>
          <n-button
            :disabled="sendDisabled"
            text
            class=""
            type="primary"
            size="small"
            @click="emits('send-msg')"
          >
            <template #icon>
              <n-icon> <Send /> </n-icon>
            </template>
          </n-button>
        </template>
      </n-input>
    </div>
    <!-- <div class="mb-1 mx-auto">
        <n-text depth="3" class="text-size-[0.6rem]">
          {{ currentAvaliableAskCountsTip }}
        </n-text>
      </div> -->
  </div>
</template>

<script setup lang="ts">
import { LogoMarkdown, Print, Send, Stop } from '@vicons/ionicons5';
import { FullscreenRound, KeyboardDoubleArrowDownRound, KeyboardDoubleArrowUpRound } from '@vicons/material';
import { useThemeVars } from 'naive-ui';
import { computed, ref } from 'vue';
import { useI18n } from 'vue-i18n';

import { useAppStore } from '@/store';

const themeVars = useThemeVars();
const appStore = useAppStore();
const { t } = useI18n();

const props = defineProps<{
  canAbort: boolean;
  sendDisabled: boolean;
  inputValue: string;
  autoScrolling: boolean;
}>();

const autoScrolling = computed({
  get() {
    return props.autoScrolling;
  },
  set(value) {
    emits('update:auto-scrolling', value);
  },
});

const inputExpanded = ref<boolean>(false);
const inputStyle = computed(() => {
  if (!inputExpanded.value)
    return {
      height: 'auto',
      maxHeight: '16vh',
    };
  return {
    height: '30vh',
  };
});

const inputValue = computed({
  get() {
    return props.inputValue;
  },
  set(value) {
    emits('update:input-value', value);
  },
});

const emits = defineEmits<{
  (e: 'abort-request'): void;
  (e: 'send-msg'): void;
  (e: 'export-to-markdown-file'): void;
  (e: 'export-to-pdf-file'): void;
  (e: 'show-fullscreen-history'): void;
  (e: 'update:auto-scrolling', value: boolean): void;
  (e: 'update:input-value', value: string): void;
}>();

const toggleInputExpanded = () => {
  inputExpanded.value = !inputExpanded.value;
};

const shortcutSendMsg = (e: KeyboardEvent) => {
  const sendKey = appStore.preference.sendKey; // "Shift+Enter" or "Ctrl+Enter" or "Enter"
  if (sendKey === 'Enter' && e.key === 'Enter' && !e.shiftKey && !e.ctrlKey) {
    e.preventDefault();
    emits('send-msg');
  } else if (sendKey === 'Shift+Enter' && e.key === 'Enter' && e.shiftKey && !e.ctrlKey) {
    e.preventDefault();
    emits('send-msg');
  } else if (sendKey === 'Ctrl+Enter' && e.key === 'Enter' && !e.shiftKey && e.ctrlKey) {
    e.preventDefault();
    emits('send-msg');
  }
};
</script>

<style>
.n-divider {
  margin-bottom: 0px !important;
  margin-top: 0px !important;
}
</style>
```

## File: frontend/src/views/conversation/components/LeftBar.vue
```vue
<template>
  <div>
    <StatusCard />
    <div class="flex-grow flex flex-col">
      <!-- <div class="flex box-content" v-if="!newConversation"> -->
      <n-button
        secondary
        strong
        type="primary"
        :disabled="props.loading"
        @click="emits('new-conversation')"
      >
        <template #icon>
          <n-icon class="">
            <Add />
          </n-icon>
        </template>
        {{ $t('commons.newConversation') }}
      </n-button>
      <!-- </div> -->
      <n-scrollbar class="h-0 flex-grow mt-4">
        <n-menu
          ref="menuRef"
          v-model:value="convId"
          class="-mx-2"
          :content-style="{ backgroundColor: 'red' }"
          :disabled="props.loading"
          :options="menuOptions"
          :root-indent="18"
        />
      </n-scrollbar>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { Add } from '@vicons/ionicons5';
import { NEllipsis } from 'naive-ui';
import { computed, h } from 'vue';
import { useI18n } from 'vue-i18n';

import { useConversationStore } from '@/store';
import { ConversationSchema } from '@/types/schema';
import { dropdownRenderer, popupChangeConversationTitleDialog } from '@/utils/renders';
import { Dialog, Message } from '@/utils/tips';

import StatusCard from './StatusCard.vue';

const { t } = useI18n();

const conversationStore = useConversationStore();

const props = defineProps<{
  loading: boolean;
  value: string | null;
  newConv: ConversationSchema | null;
}>();

const emits = defineEmits<{
  (e: 'update:value', value: string | null): void;
  (e: 'new-conversation'): void;
}>();

// get and set to bind convId and value
const convId = computed<string | null>({
  get() {
    return props.value;
  },
  set(value: string | null) {
    emits('update:value', value);
  },
});

const menuOptions = computed(() => {
  // 根据 created_time 降序排序
  const sorted_conversations = conversationStore.conversations
    ?.slice() // 创建一个新的数组副本
    .sort((a: ConversationSchema, b: ConversationSchema) => {
      // return a.create_time - b.create_time;
      if (!a.create_time) return -1;
      if (!b.create_time) return 1;
      const date_a = new Date(a.create_time),
        date_b = new Date(b.create_time);
      return date_b.getTime() - date_a.getTime();
    });
  const results = sorted_conversations?.map((conversation: ConversationSchema) => {
    return {
      label: () => h(NEllipsis, null, { default: () => conversation.title }),
      key: conversation.conversation_id,
      disabled: props.loading == true,
      extra: () => dropdownRenderer(conversation, handleDeleteConversation, handleChangeConversationTitle),
    };
  });
  if (props.newConv) {
    results?.unshift({
      label: props.newConv.title,
      key: props.newConv.conversation_id,
      disabled: props.loading == true,
    });
  }
  return results;
});

const handleDeleteConversation = (conversation_id: string | undefined) => {
  if (!conversation_id) return;
  const d = Dialog.info({
    title: t('commons.confirmDialogTitle'),
    content: t('tips.deleteConversation'),
    positiveText: t('commons.confirm'),
    negativeText: t('commons.cancel'),
    onPositiveClick: () => {
      d.loading = true;
      return new Promise((resolve) => {
        conversationStore
          .deleteConversation(conversation_id)
          .then(() => {
            Message.success(t('tips.deleteConversationSuccess'));
            if (convId.value == conversation_id) convId.value = null;
          })
          .catch(() => {
            Message.error(t('tips.deleteConversationFailed'));
          })
          .finally(() => {
            d.loading = false;
            resolve(true);
          });
      });
    },
  });
};

const handleChangeConversationTitle = (conversation_id: string | undefined) => {
  if (!conversation_id) return;
  popupChangeConversationTitleDialog(
    conversation_id,
    async (title: string) => {
      await conversationStore.changeConversationTitle(conversation_id, title);
    },
    () => {
      Message.success(t('tips.changeConversationTitleSuccess'));
    },
    () => {
      Message.error(t('tips.changeConversationTitleFailed'));
    }
  );
};
</script>

<style>
@media print {
  body * {
    visibility: hidden;
  }

  #print-content * {
    visibility: visible;
  }

  /* no margin in page */
  @page {
    margin-left: 0;
    margin-right: 0;
  }
}
</style>
```

## File: frontend/src/views/conversation/components/MessageRow.vue
```vue
<template>
  <div
    class="flex lt-sm:flex-col flex-row lt-sm:py-2 py-4 lt-sm:px-5 px-4 box-content max-w-full relative"
    :style="{ backgroundColor: backgroundColor }"
  >
    <div class="w-10 lt-sm:ml-0 ml-2 mt-3">
      <!-- <n-text class="inline-block mt-4">{{ props.message.author_role == 'user' ? 'User' : 'ChatGPT' }}</n-text> -->
      <n-avatar
        v-if="props.message.author_role == 'user'"
        size="small"
      >
        <n-icon>
          <PersonFilled />
        </n-icon>
      </n-avatar>
      <n-avatar
        v-else-if="isGpt4"
        size="small"
        :src="chatgptIconBlack"
      />
      <n-avatar
        v-else
        size="small"
        :src="chatgptIcon"
      />
    </div>
    <div class="lt-sm:mx-0 mx-4 w-full">
      <div
        v-if="!showRawContent && !renderPureText"
        ref="contentRef"
        class="message-content w-full"
        v-html="renderedContent"
      />
      <div
        v-else-if="!showRawContent && renderPureText"
        ref="contentRef"
        class="message-content w-full whitespace-pre-wrap py-4"
      >
        {{ renderedContent }}
      </div>
      <div
        v-else-if="showRawContent"
        class="my-3 w-full whitespace-pre-wrap text-gray-500"
      >
        {{ props.message.message }}
      </div>
      <div class="hide-in-print">
        <n-button
          text
          ghost
          type="tertiary"
          size="tiny"
          class="mt-2 -ml-2 absolute lt-sm:bottom-3 lt-sm:right-3 bottom-2 right-2"
          @click="copyMessageContent"
        >
          <n-icon>
            <CopyOutline />
          </n-icon>
        </n-button>
        <n-button
          text
          ghost
          size="tiny"
          :type="showRawContent ? 'success' : 'tertiary'"
          class="mt-2 -ml-2 absolute lt-sm:bottom-3 lt-sm:right-9 bottom-2 right-6"
          @click="toggleShowRawContent"
        >
          <n-icon>
            <CodeSlash />
          </n-icon>
        </n-button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { CodeSlash, CopyOutline } from '@vicons/ionicons5';
import { PersonFilled } from '@vicons/material';
import * as clipboard from 'clipboard-polyfill';
import { useThemeVars } from 'naive-ui';
import { computed, onMounted, ref } from 'vue';
import { useI18n } from 'vue-i18n';

// eslint-disable-next-line import/no-unresolved
import chatgptIcon from '/chatgpt-icon.svg';
// eslint-disable-next-line import/no-unresolved
import chatgptIconBlack from '/chatgpt-icon-black.svg';
import { useAppStore } from '@/store';
import { ChatMessage } from '@/types/custom';
import md from '@/utils/markdown';
import { Message } from '@/utils/tips';
// let md: any;
// let mdLoaded = ref(false);

// onMounted(() => {
//   import("@/utils/markdown").then((module) => {
//     md = module.default;
//     mdLoaded.value = true;
//   });
// });

const { t } = useI18n();
const appStore = useAppStore();

const themeVars = useThemeVars();

let observer = null;

const contentRef = ref<HTMLDivElement>();
const showRawContent = ref(false);

const renderPureText = computed(() => {
  return appStore.preference.renderUserMessageInMd === false && props.message.author_role == 'user';
});

const toggleShowRawContent = () => {
  showRawContent.value = !showRawContent.value;
};

const props = defineProps<{
  message: ChatMessage;
}>();

const isGpt4 = computed(() => {
  return props.message.model_slug == 'gpt-4' || props.message.model_slug == 'gpt-4-mobile';
});

const backgroundColor = computed(() => {
  if (props.message.author_role == 'user') {
    return themeVars.value.bodyColor;
  } else {
    return themeVars.value.actionColor;
  }
});

const renderedContent = computed(() => {
  // if (!mdLoaded.value) {
  //   return '';
  // }
  if (renderPureText.value) {
    return props.message.message;
  }
  const result = md.render(props.message.message || '');
  return addButtonsToPreTags(result);
});

function addButtonsToPreTags(htmlString: string): string {
  // Parse the HTML string into an Element object.
  const parser = new DOMParser();
  const doc = parser.parseFromString(htmlString, 'text/html');

  // Get all the <pre> elements in the document.
  const preTags = doc.getElementsByTagName('pre');

  // Loop through the <pre> elements and add a <button> to each one.
  for (let i = 0; i < preTags.length; i++) {
    const preTag = preTags[i];

    const button = Object.assign(document.createElement('button'), {
      innerHTML: '',
      className: 'hljs-copy-button hide-in-print',
    });
    button.dataset.copied = 'false';
    preTag.classList.add('hljs-copy-wrapper');

    // Add a custom proprety to the code block so that the copy button can reference and match its background-color value.
    preTag.style.setProperty('--hljs-theme-background', window.getComputedStyle(preTag).backgroundColor);

    if (appStore.preference.codeAutoWrap) {
      preTag.style.cssText += 'white-space: pre-wrap; word-wrap: break-word; word-break: break-all;';
    }

    preTag.appendChild(button);
  }

  // Serialize the modified Element object back into a string.
  const serializer = new XMLSerializer();
  return serializer.serializeToString(doc.documentElement);
}

onMounted(() => {
  if (!contentRef.value) return;
  // eslint-disable-next-line no-undef
  const callback: MutationCallback = (mutations: MutationRecord[], observer: MutationObserver) => {
    for (const mutation of mutations) {
      if (mutation.type === 'childList') {
        bindOnclick();
      }
    }
  };
  observer = new MutationObserver(callback);
  observer.observe(contentRef.value, { subtree: true, childList: true });
  bindOnclick();
});

const bindOnclick = () => {
  // 获取模板引用中的所有 pre 元素和其子元素中的 button 元素
  const preElements = contentRef.value?.querySelectorAll('pre');
  if (!preElements) return;
  for (const preElement of preElements as any) {
    for (const button of preElement.querySelectorAll('button')) {
      (button as HTMLButtonElement).onmousedown = () => {
        // 如果按钮的内容为 "Copied!"，则跳过复制操作
        if (button.innerHTML === 'Copied!') {
          return;
        }

        const preContent = button.parentElement!.cloneNode(true) as HTMLElement;
        preContent.removeChild(preContent.querySelector('button')!);

        // Remove the alert element if it exists in preContent
        const alertElement = preContent.querySelector('.hljs-copy-alert');
        if (alertElement) {
          preContent.removeChild(alertElement);
        }

        clipboard
          .writeText(preContent.textContent || '')
          .then(function () {
            button.innerHTML = 'Copied!';
            button.dataset.copied = 'true';

            let alert: HTMLDivElement | null = Object.assign(document.createElement('div'), {
              role: 'status',
              className: 'hljs-copy-alert',
              innerHTML: 'Copied to clipboard',
            });
            button.parentElement!.appendChild(alert);

            setTimeout(() => {
              if (alert) {
                button.innerHTML = 'Copy';
                button.dataset.copied = 'false';
                button.parentElement!.removeChild(alert);
                alert = null;
              }
            }, 2000);
          })
          .then();
      };
    }
  }
};

const copyMessageContent = () => {
  /* debugger
  if (!navigator.clipboard) return;
  navigator.clipboard
    .writeText(props.message.message || "")
    .then(() => {
      // console.log('copied', props.message.message);
      Message.success(t('commons.copiedToClipboard'))
    }
    ).then(); */
  const messageContent = props.message.message || '';
  clipboard
    .writeText(messageContent)
    .then(() => {
      Message.success(t('commons.copiedToClipboard'));
    })
    .catch(() => {
      console.error('Failed to copy message content to clipboard.');
    });
};
</script>

<style>
/* modified from https://github.com/arronhunt/highlightjs-copy */

pre {
  @apply w-full flex;
}

pre code {
  /* @apply w-full max-w-94 sm: max-w-138 md:max-w-156 lg:max-w-170 */
  @apply w-0 flex-grow mr-0 font-mono;
}

@media print {
  code {
    @apply max-w-160 !important
    @apply whitespace-pre-line;
  }
}

p {
  white-space: pre-line;
}

ol,
ul {
  padding-left: 16px;
}

.message-content table {
  border: gray 1px solid;
  @apply min-w-1/2 text-center border-collapse;
}

.message-content tr {
  border: gray 1px solid;
}

.message-content th {
  border: gray 1px solid;
  @apply bg-gray-400;
}

.message-content td {
  border: gray 1px solid;
}

.hljs-copy-wrapper {
  position: relative;
  overflow: hidden;
}

.hljs-copy-wrapper:hover .hljs-copy-button,
.hljs-copy-button:focus {
  transform: translateX(0);
}

.hljs-copy-button {
  position: absolute;
  transform: translateX(calc(100% + 2em));
  top: 1em;
  right: 1.2em;
  width: 2rem;
  height: 2rem;
  text-indent: -9999px;
  /* Hide the inner text */
  color: #c5c5c5;
  border-radius: 0.25rem;
  border: 1px solid #c5c5c522;
  /* background-color: #2d2b57; */

  /* 白色，不透明度10% */
  background-color: #ffffff1a;
  background-image: url('data:image/svg+xml;utf-8,<svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" clip-rule="evenodd" d="M6 5C5.73478 5 5.48043 5.10536 5.29289 5.29289C5.10536 5.48043 5 5.73478 5 6V20C5 20.2652 5.10536 20.5196 5.29289 20.7071C5.48043 20.8946 5.73478 21 6 21H18C18.2652 21 18.5196 20.8946 18.7071 20.7071C18.8946 20.5196 19 20.2652 19 20V6C19 5.73478 18.8946 5.48043 18.7071 5.29289C18.5196 5.10536 18.2652 5 18 5H16C15.4477 5 15 4.55228 15 4C15 3.44772 15.4477 3 16 3H18C18.7956 3 19.5587 3.31607 20.1213 3.87868C20.6839 4.44129 21 5.20435 21 6V20C21 20.7957 20.6839 21.5587 20.1213 22.1213C19.5587 22.6839 18.7957 23 18 23H6C5.20435 23 4.44129 22.6839 3.87868 22.1213C3.31607 21.5587 3 20.7957 3 20V6C3 5.20435 3.31607 4.44129 3.87868 3.87868C4.44129 3.31607 5.20435 3 6 3H8C8.55228 3 9 3.44772 9 4C9 4.55228 8.55228 5 8 5H6Z" fill="black"/><path fill-rule="evenodd" clip-rule="evenodd" d="M7 3C7 1.89543 7.89543 1 9 1H15C16.1046 1 17 1.89543 17 3V5C17 6.10457 16.1046 7 15 7H9C7.89543 7 7 6.10457 7 5V3ZM15 3H9V5H15V3Z" fill="black"/></svg>');
  background-repeat: no-repeat;
  background-position: center;
  transition: background-color 200ms ease, transform 200ms ease-out;
}

.hljs-copy-button:hover {
  border-color: #91919144;
}

.hljs-copy-button:active {
  border-color: #49494966;
}

.hljs-copy-button[data-copied='true'] {
  text-indent: 0px;
  /* Shows the inner text */
  width: auto;
  background-image: none;
}

@media (prefers-reduced-motion) {
  .hljs-copy-button {
    transition: none;
  }
}

/* visually-hidden */
.hljs-copy-alert {
  clip: rect(0 0 0 0);
  clip-path: inset(50%);
  height: 1px;
  overflow: hidden;
  position: absolute;
  white-space: nowrap;
  width: 1px;
  color: #2d2b57;
}
</style>
```

## File: frontend/src/views/conversation/components/StatusCard.vue
```vue
<template>
  <n-card content-style="padding: 0;">
    <n-collapse @update:expanded-names="handleExpand">
      <n-collapse-item
        :title="$t('commons.serverStatus')"
        name="serverStatus"
      >
        <n-list
          hoverable
          show-divider
        >
          <n-list-item>
            <div class="flex flex-row justify-between content-center">
              <div>
                <n-icon class="mr-1">
                  <md-people />
                </n-icon>{{ $t('commons.activeUserIn5m') }}
              </div>
              <div>{{ serverStatus.active_user_in_5m }}</div>
            </div>
          </n-list-item>
          <n-list-item>
            <div class="flex flex-row justify-between content-center">
              <div>
                <n-icon class="mr-1">
                  <md-people />
                </n-icon>{{ $t('commons.activeUserIn1h') }}
              </div>
              <div>{{ serverStatus.active_user_in_1h }}</div>
            </div>
          </n-list-item>
          <n-list-item>
            <div class="flex flex-row justify-between content-center">
              <div>
                <n-icon class="mr-1">
                  <md-people />
                </n-icon>{{ $t('commons.activeUserIn1d') }}
              </div>
              <div>{{ serverStatus.active_user_in_1d }}</div>
            </div>
          </n-list-item>
          <n-list-item>
            <div class="flex flex-row justify-between content-center">
              <div>
                <n-icon class="mr-1">
                  <EventBusyFilled />
                </n-icon>{{ $t('commons.isChatbotBusy') }}
              </div>
              <div>{{ serverStatus.is_chatbot_busy ? $t('commons.yes') : $t('commons.no') }}</div>
            </div>
          </n-list-item>
          <n-list-item>
            <div class="flex flex-row justify-between content-center">
              <div>
                <n-icon class="mr-1">
                  <QueueFilled />
                </n-icon>{{ $t('commons.chatbotWaitingCount') }}
              </div>
              <div>{{ serverStatus.chatbot_waiting_count }}</div>
            </div>
          </n-list-item>
        </n-list>
      </n-collapse-item>
    </n-collapse>
  </n-card>
</template>

<script setup lang="ts">
import { MdPeople } from '@vicons/ionicons4';
import { EventBusyFilled, QueueFilled } from '@vicons/material';
import { ref } from 'vue';
import { useI18n } from 'vue-i18n';

import { getServerStatusApi } from '@/api/status';
import { ServerStatusSchema } from '@/types/schema';
const { t } = useI18n();
const serverStatus = ref<ServerStatusSchema>({});

const isExpaned = ref(false);

const handleExpand = (names: string[]) => {
  if (names.length > 0) {
    isExpaned.value = true;
    updateData();
  } else {
    isExpaned.value = false;
  }
};

const updateData = () => {
  if (isExpaned.value)
    getServerStatusApi().then((res) => {
      // console.log(res.data);
      serverStatus.value = res.data;
    });
};
updateData();
setInterval(updateData, 5000);
</script>

<style>
div.n-collapse-item {
  padding: 1em;
}
</style>
```

## File: frontend/src/views/conversation/history-viewer.vue
```vue
<template>
  <HistoryContent
    :messages="messages"
    :fullscreen="false"
    :show-tips="false"
    :append-messages="[]"
    :loading="loading"
  />
</template>

<script setup lang="ts">
import { computed,ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';

import { useConversationStore } from '@/store';
import { getConvMessageListFromId } from '@/utils/conversation';

import HistoryContent from './components/HistoryContent.vue';

const conversationStore = useConversationStore();

const route = useRoute();
const router = useRouter();
const conversationId = route.params.conversation_id as string;

const loading = ref(true);

conversationStore
  .fetchConversationHistory(conversationId)
  .then(() => {
    // console.log(conversationStore.conversationDetailMap);
  })
  .catch((err: any) => {
    console.log(err);
    router.push({ name: '404' }).then();
  })
  .finally(() => {
    loading.value = false;
  });

const messages = computed(() => {
  return getConvMessageListFromId(conversationId);
});
</script>
```

## File: frontend/src/views/conversation/index.vue
```vue
<template>
  <div
    ref="rootRef"
    :class="['flex-grow flex flex-col md:flex-row', !appStore.preference.widerConversationPage ? 'lg:w-screen-lg lg:mx-auto' : '']"
  >
    <!-- 左栏 -->
    <LeftBar
      v-show="!foldLeftBar"
      v-model:value="currentConversationId"
      :class="[
        'md:min-w-50 pl-4 lt-md:pr-4 box-border mb-4 lt-md:h-56 md:flex-grow overflow-hidden flex flex-col space-y-4',
        appStore.preference.widerConversationPage ? 'md:w-1/5' : 'md:w-1/4',
      ]"
      :loading="loadingBar"
      :new-conv="newConversation"
      @new-conversation="makeNewConversation"
    />
    <!-- 右栏 -->
    <div
      :class="[
        'flex-grow flex flex-col md:px-4',
        appStore.preference.widerConversationPage ? 'md:w-4/5' : 'md:w-3/4',
      ]"
    >
      <n-card
        class="flex-grow md:mb-4 relative"
        :bordered="true"
        content-style="padding: 0; display: flex; flex-direction: column; "
      >
        <!-- 展开/收起左栏 -->
        <div class="left-3 top-3 absolute z-20">
          <n-button
            strong
            secondary
            :type="foldLeftBar ? 'default' : 'primary'"
            size="small"
            @click="foldLeftBar = !foldLeftBar"
          >
            <template #icon>
              <n-icon :component="MenuRound" />
            </template>
          </n-button>
        </div>
        
        <!-- 消息记录内容（用于全屏展示） -->
        <n-scrollbar
          v-if="currentConversationId"
          ref="historyRef"
          class="basis-0 flex-grow shrink-grow relative"
          :style="{'overflow-y': 'scroll','-webkit-overflow-scrolling': 'touch'}"
          :content-style="loadingHistory ? { height: '100%' } : { }"
        >
          <!-- 回到底部按钮 -->
          <div class="right-2 bottom-5 absolute z-20">
            <n-button
              secondary
              circle
              size="small"
              @click="scrollToBottomSmooth"
            >
              <template #icon>
                <n-icon :component="ArrowDown" />
              </template>
            </n-button>
          </div>
          <HistoryContent
            ref="historyContentRef"
            :messages="currentMessageListDisplay"
            :fullscreen="false"
            :model-name="currentConversation?.model_name || ''"
            :show-tips="showFullscreenTips"
            :loading="loadingHistory"
          />
        </n-scrollbar>
        <!-- 未选中对话（空界面） -->
        <div
          v-else-if="!currentConversationId"
          class="flex-grow flex flex-col justify-center"
          :style="{ backgroundColor: themeVars.cardColor }"
        >
          <n-empty
            v-if="!currentConversation"
            :description="$t('tips.loadConversation')"
          >
            <template #icon>
              <n-icon>
                <ChatboxEllipses />
              </n-icon>
            </template>
            <template #extra>
              <n-button @click="makeNewConversation">
                {{ $t('tips.newConversation') }}
              </n-button>
            </template>
          </n-empty>
        </div>
        <!-- 下半部分（回复区域） -->
        <InputRegion
          v-model:input-value="inputValue"
          v-model:auto-scrolling="autoScrolling"
          :can-abort="canAbort"
          :send-disabled="sendDisabled"
          @abort-request="abortRequest"
          @export-to-markdown-file="exportToMarkdownFile"
          @export-to-pdf-file="exportToPdfFile"
          @send-msg="sendMsg"
          @show-fullscreen-history="showFullscreenHistory"
        />
      </n-card>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ArrowDown, ChatboxEllipses } from '@vicons/ionicons5';
import { MenuRound } from '@vicons/material';
import { RemovableRef, useStorage } from '@vueuse/core';
import { NButton, NIcon, useThemeVars } from 'naive-ui';
import { computed, ref, watch } from 'vue';
import { useI18n } from 'vue-i18n';

import { AskInfo, getAskWebsocketApiUrl } from '@/api/chat';
import { useAppStore, useConversationStore, useUserStore } from '@/store';
import { ChatConversationDetail, ChatMessage } from '@/types/custom';
import { ConversationSchema } from '@/types/schema';
import { getConvMessageListFromId } from '@/utils/conversation';
import { popupNewConversationDialog } from '@/utils/renders';
import { Dialog, LoadingBar, Message } from '@/utils/tips';
import HistoryContent from '@/views/conversation/components/HistoryContent.vue';
import InputRegion from '@/views/conversation/components/InputRegion.vue';
import LeftBar from '@/views/conversation/components/LeftBar.vue';

import { saveAsMarkdown } from './utils/export';

const themeVars = useThemeVars();

const { t } = useI18n();

const rootRef = ref();
const historyRef = ref();
const userStore = useUserStore();
const appStore = useAppStore();
const conversationStore = useConversationStore();

const loadingBar = ref(false);
const loadingHistory = ref<boolean>(false);
const autoScrolling = ref<boolean>(true);

const isAborted = ref<boolean>(false);
const canAbort = ref<boolean>(false);
const foldLeftBar = ref<RemovableRef<boolean>>(useStorage('foldLeftBar', false));
let aborter: (() => void) | null = null;

// const currentAvaliableAskCountsTip = computed(() => {
//   let result = '';
//   if (userStore.user?.available_ask_count != -1)
//     result += `${t('commons.availableAskCount')}: ${getCountTrans(userStore.user?.available_ask_count!)}   `;
//   if (currentConversation.value && currentConversation.value.model_name === 'gpt-4' && userStore.user?.available_gpt4_ask_count != -1)
//     result += `${t('commons.availableGPT4AskCount')}: ${getCountTrans(userStore.user?.available_gpt4_ask_count!)}`;
//   return result;
// });

const newConversation = ref<ConversationSchema | null>(null);
const currentConversationId = ref<string | null>(null);
const currentConversation = computed<ConversationSchema>(() => {
  if (newConversation.value?.conversation_id === currentConversationId.value) return newConversation.value;
  const conv = conversationStore.conversations?.find((conversation: ConversationSchema) => {
    return conversation.conversation_id == currentConversationId.value;
  });
  return conv;
});

const inputValue = ref('');
const currentActiveMessageSend = ref<ChatMessage | null>(null);
const currentActiveMessageRecv = ref<ChatMessage | null>(null);
const currentMessageListDisplay = computed(() => {
  const conversationId = currentConversationId.value;
  if (!conversationId) return [];
  // const _ensure_conv = conversationStore.conversationDetailMap[props.conversationId];
  let result = getConvMessageListFromId(conversationId);
  if (currentActiveMessages.value.length > 0) {
    result = result.concat(currentActiveMessages.value);
  }
  return result;
});

// 从 store 中获取当前对话最新消息的 id
const currentNode = computed<string | undefined>(() => {
  if (currentConversation.value?.conversation_id)
    return conversationStore.conversationDetailMap[currentConversation.value?.conversation_id]?.current_node;
  else return undefined;
});

// 实际的 currentMessageList，加上当前正在发送的消息
const currentActiveMessages = computed<Array<ChatMessage>>(() => {
  const result: ChatMessage[] = [];
  if (currentActiveMessageSend.value && result.findIndex((message) => message.id === currentActiveMessageSend.value?.id) === -1)
    result.push(currentActiveMessageSend.value);
  if (currentActiveMessageRecv.value && result.findIndex((message) => message.id === currentActiveMessageRecv.value?.id) === -1)
    result.push(currentActiveMessageRecv.value);
  return result;
});

watch(currentConversationId, (newVal, _oldVal) => {
  if (newVal != 'new_conversation') {
    handleChangeConversation(newVal);
  }
});

const handleChangeConversation = (key: string | null) => {
  // TODO: 清除当前已询问、得到回复，但是发生错误的两条消息
  if (loadingBar.value || !key) return;
  loadingBar.value = true;
  loadingHistory.value = true;
  LoadingBar.start();
  conversationStore
    .fetchConversationHistory(key)
    .then(() => {
      // console.log(conversationStore.conversationDetailMap);
    })
    .catch((err: any) => {
      console.log(err);
    })
    .finally(() => {
      loadingBar.value = false;
      loadingHistory.value = false;
      LoadingBar.finish();
    });
};

const sendDisabled = computed(() => {
  return loadingBar.value || currentConversationId.value == null || inputValue.value === null || inputValue.value.trim() == '';
});

const makeNewConversation = () => {
  if (newConversation.value) return;
  popupNewConversationDialog(async (title: string, model_name: any) => {
    // console.log(title, model_name);
    newConversation.value = {
      conversation_id: 'new_conversation',
      // 默认标题格式：MMDD - username
      title: title || `New Chat ${new Date().toLocaleString()} - ${userStore.user?.username}`,
      model_name: model_name || 'text-davinci-002-render-sha',
      create_time: new Date().toISOString(), // 仅用于当前排序到顶部
    };
    currentConversationId.value = 'new_conversation';
  });
};

const abortRequest = () => {
  if (aborter == null || !canAbort.value) return;
  aborter();
  aborter = null;
};

const scrollToBottom = () => {
  historyRef.value.scrollTo({ left: 0, top: historyRef.value.$refs.scrollbarInstRef.contentRef.scrollHeight });
};

const scrollToBottomSmooth = () => {
  historyRef.value.scrollTo({ left: 0, top: historyRef.value.$refs.scrollbarInstRef.contentRef.scrollHeight, behavior: 'smooth' });
};

const sendMsg = async () => {
  if (sendDisabled.value || loadingBar.value) {
    Message.error(t('tips.pleaseSelectConversation'));
    return;
  }

  LoadingBar.start();
  loadingBar.value = true;
  const message = inputValue.value;
  inputValue.value = '';

  canAbort.value = false;
  isAborted.value = false;
  let hasGotReply = false;

  const askInfo: AskInfo = { message };
  if (newConversation.value) {
    askInfo.new_title = newConversation.value.title;
    askInfo.model_name = newConversation.value.model_name;
  } else {
    askInfo.conversation_id = currentConversation.value!.conversation_id;
    askInfo.parent_id = currentNode.value!;
  }

  // 使用临时的随机 id 保持当前更新的两个消息
  const random_strid = Math.random().toString(36).substring(2, 16);
  currentActiveMessageSend.value = {
    id: `send_${random_strid}`,
    message,
    author_role: 'user',
    parent: currentNode.value,
    children: [`recv_${random_strid}`],
  };
  currentActiveMessageRecv.value = {
    id: `recv_${random_strid}`,
    message: '',
    author_role: 'assistent',
    parent: `send_${random_strid}`,
    children: [],
    typing: true,
    model_slug: currentConversation.value?.model_name,
  };
  const wsUrl = getAskWebsocketApiUrl();
  let wsErrorMessage: string | null = null;
  console.log('Connecting to', wsUrl, askInfo);
  const webSocket = new WebSocket(wsUrl);

  webSocket.onopen = (_event: Event) => {
    // console.log('WebSocket connection is open', askInfo);
    webSocket.send(JSON.stringify(askInfo));
  };

  webSocket.onmessage = (event: MessageEvent) => {
    const reply = JSON.parse(event.data);
    // console.log('Received message from server:', reply);
    if (!reply.type) return;
    if (reply.type === 'waiting') {
      // 等待回复
      canAbort.value = false;
      currentActiveMessageRecv.value!.message = t(reply.tip);
    } else if (reply.type === 'queueing') {
      // 正在排队
      canAbort.value = true;
      currentActiveMessageRecv.value!.message = t(reply.tip);
      // if (reply.waiting_count) {
      //   currentActiveMessageRecv.value!.message += `(${reply.waiting_count})`;
      // }
    } else if (reply.type === 'message') {
      // console.log(reply)
      hasGotReply = true;
      currentActiveMessageRecv.value!.message = reply.message;
      currentActiveMessageRecv.value!.id = reply.parent_id;
      currentActiveMessageRecv.value!.model_slug = reply.model_name;
      if (newConversation.value) {
        newConversation.value.model_name = reply.model_name;
        if (newConversation.value.conversation_id !== reply.conversation_id) newConversation.value.conversation_id = reply.conversation_id;
        if (currentConversationId.value !== newConversation.value.conversation_id) {
          currentConversationId.value = newConversation.value.conversation_id!;
        }
      }
      canAbort.value = true;
    } else if (reply.type === 'error') {
      currentActiveMessageRecv.value!.message = `${t(reply.tip)}: ${reply.message}}`;
      console.error(reply.tip, reply.message);
      if (reply.message) {
        wsErrorMessage = reply.message;
      }
    }
    if (autoScrolling.value) scrollToBottom();
  };

  webSocket.onclose = async (event: CloseEvent) => {
    aborter = null;
    canAbort.value = false;
    currentActiveMessageRecv.value!.typing = false;
    console.log('WebSocket connection is closed', event, isAborted.value);
    if (isAborted.value || event.code === 1000) {
      // 正常关闭
      if (hasGotReply) {
        if (newConversation.value) {
          // 解析 ISO string 为 小数时间戳
          const create_time = new Date(newConversation.value.create_time!).getTime() / 1000;
          const newConvDetail = {
            id: currentConversationId.value,
            title: newConversation.value!.title,
            model_name: newConversation.value!.model_name,
            create_time,
            mapping: {},
            current_node: null,
          } as ChatConversationDetail;
          conversationStore.$patch({
            conversationDetailMap: {
              [newConversation.value.conversation_id!]: newConvDetail,
            },
          });
          const msgSend = currentActiveMessageSend.value;
          const msgRecv = currentActiveMessageRecv.value;
          currentActiveMessageSend.value = null;
          currentActiveMessageRecv.value = null;
          conversationStore.addMessageToConversation(currentConversationId.value, msgSend, msgRecv);
          currentConversationId.value = newConversation.value.conversation_id!; // 这里将会导致 currentConversation 切换
          await conversationStore.fetchAllConversations();
          newConversation.value = null;
          console.log('done', newConvDetail, msgSend, msgRecv, currentConversationId.value);
        } else {
          // 将新消息存入 store
          if (!currentActiveMessageRecv.value!.id.startsWith('recv')) {
            // TODO 其它属性
            conversationStore.addMessageToConversation(currentConversationId.value, currentActiveMessageSend.value!, currentActiveMessageRecv.value!);
          }
          currentActiveMessageSend.value = null;
          currentActiveMessageRecv.value = null;
        }
      }
    } else {
      Dialog.error({
        title: t('errors.askError'),
        content: wsErrorMessage != null ? `[${event.code}] ${t(event.reason)}: ${wsErrorMessage}` : `[${event.code}] ${t(event.reason)}`,
        positiveText: t('commons.withdrawMessage'),
        negativeText: t('commons.cancel'),
        onPositiveClick: () => {
          currentActiveMessageSend.value = null;
          currentActiveMessageRecv.value = null;
        },
      });
    }
    await userStore.fetchUserInfo();
    LoadingBar.finish();
    loadingBar.value = false;
    isAborted.value = false;
  };

  webSocket.onerror = (event: Event) => {
    console.error('WebSocket error:', event);
  };

  aborter = () => {
    isAborted.value = true;
    webSocket.close();
  };
};

const exportToMarkdownFile = () => {
  if (!currentConversation.value) {
    Message.error(t('tips.pleaseSelectConversation'));
    return;
  }
  saveAsMarkdown(currentConversation.value, currentMessageListDisplay.value);
};

const historyContentRef = ref();
const showFullscreenTips = ref(false);

const showFullscreenHistory = () => {
  if (!currentConversation.value) {
    Message.error(t('tips.pleaseSelectConversation'));
    return;
  }
  // focus historyContentRef
  historyContentRef.value.focus();
  historyContentRef.value.toggleFullscreenHistory(true);
};

const exportToPdfFile = () => {
  if (!currentConversation.value) {
    Message.error(t('tips.pleaseSelectConversation'));
    return;
  }
  historyContentRef.value.toggleFullscreenHistory(false);
  window.print();
  historyContentRef.value.toggleFullscreenHistory(false);
};

// 加载对话列表
conversationStore.fetchAllConversations().then();
</script>

<style>
textarea.n-input__textarea-el {
  resize: none;
}

div.n-menu-item-content-header {
  display: flex;
  justify-content: space-between;
}

span.n-menu-item-content-header__extra {
  display: inline-block;
}

.left-col .n-card__content {
  @apply flex flex-col overflow-auto !important;
}

@media print {
  body * {
    visibility: hidden;
  }

  #print-content * {
    visibility: visible;
  }

  /* no margin in page */
  @page {
    margin-left: 0;
    margin-right: 0;
  }
}
</style>
```

## File: frontend/src/views/conversation/utils/export.ts
```typescript
import { saveAs } from 'file-saver';

import { ChatMessage } from '@/types/custom';
import { ConversationSchema } from '@/types/schema';
import { getModelNameTrans } from '@/utils/renders';

export const saveAsMarkdown = (conv: ConversationSchema, messageList: ChatMessage[]) => {
  let content = `# ${conv.title}\n\n`;
  const create_time = new Date(conv.create_time ? conv.create_time + 'Z' : new Date()).toLocaleString();
  content += `Date: ${create_time}\nModel: ${getModelNameTrans(conv.model_name as any)}\n`;
  content += 'generated by [ChatGPT Web Share](https://github.com/moeakwak/chatgpt-web-share)\n\n';
  content += '---\n\n';
  let index = 0;
  for (const message of messageList) {
    if (message.author_role === 'user') {
      // 选取第一行作为标题，最多50个字符，如果有省略则加上...
      let title = message.message!.trim().split('\n')[0];
      if (title.length >= 50) {
        title = title.slice(0, 47) + '...';
      }
      content += `## ${++index}. ${title}\n\n`;
      content += `### User\n\n${message.message}\n\n`;
    } else {
      content += `### ChatGPT\n\n${message.message}\n\n`;
      content += '---\n\n';
    }
  }
  const blob = new Blob([content], { type: 'text/plain;charset=utf-8' });
  saveAs(blob, `${conv.title} - ChatGPT history.md`);
};
```

## File: frontend/src/views/error/403.vue
```vue
<template>
  <!-- 403 error page -->
  <div class="w-full h-full lg:w-screen-lg mx-auto h-screen flex flex-col mt-3">
    <div class="flex-1 flex justify-center items-center">
      <n-card class="h-full mb-6 flex justify-center items-center">
        <div class="text-2xl text-gray-500">
          {{ $t('errors.403') }}
        </div>
      </n-card>
    </div>
  </div>
</template>
```

## File: frontend/src/views/error/404.vue
```vue
<template>
  <!-- 404 error page -->
  <div class="w-full h-full lg:w-screen-lg mx-auto h-screen flex flex-col mt-3">
    <div class="flex-1 flex justify-center items-center">
      <n-card class="h-full mb-6 flex justify-center items-center">
        <div class="text-2xl text-gray-500">
          {{ $t('errors.404') }}
        </div>
      </n-card>
    </div>
  </div>
</template>
```

## File: frontend/src/views/home.vue
```vue
<template>
  <n-card class="h-full mb-6 flex justify-center items-center">
    {{ $t('tips.jumpingPage') }}
  </n-card>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';

import { useUserStore } from '@/store';

const router = useRouter();
const userStore = useUserStore();

let target = '';
if (!userStore.user) {
  target = 'login';
} else {
  target = 'conversation';
}

router
  .push({
    name: target,
  })
  .then(() => {
    window.location.reload();
  });
</script>
```

## File: frontend/src/views/login/index.vue
```vue
<template>
  <!-- Login Form -->
  <div class="flex justify-center items-center mt-20">
    <n-form
      ref="formRef"
      :model="formValue"
      :rules="loginRules"
      :label-col="{ span: 8 }"
      :wrapper-col="{ span: 16 }"
    >
      <n-form-item
        :label="$t('commons.username')"
        path="username"
      >
        <n-input
          v-model:value="formValue.username"
          :placeholder="$t('tips.pleaseEnterUsername')"
          :input-props="{
            autoComplete: 'username',
          }"
        />
      </n-form-item>
      <n-form-item
        :label="$t('commons.password')"
        path="password"
      >
        <n-input
          v-model:value="formValue.password"
          type="password"
          show-password-on="click"
          :placeholder="$t('tips.pleaseEnterPassword')"
          :input-props="{
            autoComplete: 'current-password',
          }"
          @keyup.enter="login"
        />
      </n-form-item>
      <n-form-item wrapper-col="{ span: 16, offset: 8 }">
        <n-button
          type="primary"
          :enabled="loading"
          @click="login"
        >
          {{ $t('commons.login') }}
        </n-button>
      </n-form-item>
    </n-form>
  </div>
</template>

<script setup lang="ts">
import { FormInst } from 'naive-ui';
import { FormValidationError } from 'naive-ui/es/form';
import { reactive,ref } from 'vue';
import { useI18n } from 'vue-i18n';
import { useRouter } from 'vue-router';

import { loginApi, LoginData } from '@/api/user';
import { useUserStore } from '@/store';
import { Message } from '@/utils/tips';

const router = useRouter();
const { t } = useI18n();
const userStore = useUserStore();
const formRef = ref<FormInst>();

const formValue = reactive({
  username: '',
  password: '',
});
const loading = ref(false);
const loginRules = {
  username: { required: true, message: t('tips.pleaseEnterUsername'), trigger: 'blur' },
  password: { required: true, message: t('tips.pleaseEnterPassword'), trigger: 'blur' },
};

const login = async () => {
  if (loading.value) return;
  formRef.value
    ?.validate((errors?: Array<FormValidationError>) => {
      if (!errors) {
        loading.value = true;
      }
    })
    .then(async () => {
      try {
        await userStore.login(formValue as LoginData);
        const { redirect, ...othersQuery } = router.currentRoute.value.query;
        await userStore.fetchUserInfo();
        Message.success(t('tips.loginSuccess'));
        await router.push({
          name: userStore.user?.is_superuser ? 'admin' : 'conversation',
        });
        // TODO: 记住密码
      } catch (error) {
        console.log(error);
      } finally {
        loading.value = false;
      }
    });
};

if (userStore.user) {
  router.push({ name: 'conversation' });
}
</script>
```

## File: frontend/src/views/redirect/index.vue
```vue
<template>
  <div />
</template>

<script lang="ts" setup>
import { useRoute,useRouter } from 'vue-router';

const router = useRouter();
const route = useRoute();

const gotoPath = route.params.path as string;

router.replace({ path: gotoPath });
</script>

<style scoped lang="less"></style>
```

## File: frontend/src/vite-env.d.ts
```typescript
/// <reference types="vite/client" />
```

## File: frontend/tsconfig.json
```json
{
  "compilerOptions": {
    "target": "ESNext",
    "useDefineForClassFields": true,
    "module": "ESNext",
    "moduleResolution": "Node",
    "strict": true,
    "jsx": "preserve",
    "resolveJsonModule": true,
    "isolatedModules": true,
    "esModuleInterop": true,
    "lib": [
      "ESNext",
      "DOM"
    ],
    "skipLibCheck": true,
    "noEmit": true,
    "baseUrl": ".",
    "paths": {
      "@/*": [
        "src/*"
      ]
    }
  },
  "include": [
    "src/**/*.ts",
    "src/**/*.d.ts",
    "src/**/*.tsx",
    "src/**/*.vue",
    "src/views/conversation/utils/markdown.js"
  ],
  "references": [
    {
      "path": "./tsconfig.node.json"
    }
  ]
}
```

## File: frontend/tsconfig.node.json
```json
{
  "compilerOptions": {
    "composite": true,
    "module": "ESNext",
    "moduleResolution": "Node",
    "allowSyntheticDefaultImports": true
  },
  "include": [
    "config/vite.config.base.ts",
    "config/vite.config.dev.ts",
    "config/vite.config.prod.ts",
  ]
}
```

## File: frontend/updateapi.sh
```bash
cd src/types;
wget http://127.0.0.1:8000/openapi.json -O openapi.json;
npx openapi-typescript openapi.json --output ./openapi.ts;
```

## File: LICENSE
```
GNU GENERAL PUBLIC LICENSE
                       Version 3, 29 June 2007

 Copyright (C) 2007 Free Software Foundation, Inc. <https://fsf.org/>
 Everyone is permitted to copy and distribute verbatim copies
 of this license document, but changing it is not allowed.

                            Preamble

  The GNU General Public License is a free, copyleft license for
software and other kinds of works.

  The licenses for most software and other practical works are designed
to take away your freedom to share and change the works.  By contrast,
the GNU General Public License is intended to guarantee your freedom to
share and change all versions of a program--to make sure it remains free
software for all its users.  We, the Free Software Foundation, use the
GNU General Public License for most of our software; it applies also to
any other work released this way by its authors.  You can apply it to
your programs, too.

  When we speak of free software, we are referring to freedom, not
price.  Our General Public Licenses are designed to make sure that you
have the freedom to distribute copies of free software (and charge for
them if you wish), that you receive source code or can get it if you
want it, that you can change the software or use pieces of it in new
free programs, and that you know you can do these things.

  To protect your rights, we need to prevent others from denying you
these rights or asking you to surrender the rights.  Therefore, you have
certain responsibilities if you distribute copies of the software, or if
you modify it: responsibilities to respect the freedom of others.

  For example, if you distribute copies of such a program, whether
gratis or for a fee, you must pass on to the recipients the same
freedoms that you received.  You must make sure that they, too, receive
or can get the source code.  And you must show them these terms so they
know their rights.

  Developers that use the GNU GPL protect your rights with two steps:
(1) assert copyright on the software, and (2) offer you this License
giving you legal permission to copy, distribute and/or modify it.

  For the developers' and authors' protection, the GPL clearly explains
that there is no warranty for this free software.  For both users' and
authors' sake, the GPL requires that modified versions be marked as
changed, so that their problems will not be attributed erroneously to
authors of previous versions.

  Some devices are designed to deny users access to install or run
modified versions of the software inside them, although the manufacturer
can do so.  This is fundamentally incompatible with the aim of
protecting users' freedom to change the software.  The systematic
pattern of such abuse occurs in the area of products for individuals to
use, which is precisely where it is most unacceptable.  Therefore, we
have designed this version of the GPL to prohibit the practice for those
products.  If such problems arise substantially in other domains, we
stand ready to extend this provision to those domains in future versions
of the GPL, as needed to protect the freedom of users.

  Finally, every program is threatened constantly by software patents.
States should not allow patents to restrict development and use of
software on general-purpose computers, but in those that do, we wish to
avoid the special danger that patents applied to a free program could
make it effectively proprietary.  To prevent this, the GPL assures that
patents cannot be used to render the program non-free.

  The precise terms and conditions for copying, distribution and
modification follow.

                       TERMS AND CONDITIONS

  0. Definitions.

  "This License" refers to version 3 of the GNU General Public License.

  "Copyright" also means copyright-like laws that apply to other kinds of
works, such as semiconductor masks.

  "The Program" refers to any copyrightable work licensed under this
License.  Each licensee is addressed as "you".  "Licensees" and
"recipients" may be individuals or organizations.

  To "modify" a work means to copy from or adapt all or part of the work
in a fashion requiring copyright permission, other than the making of an
exact copy.  The resulting work is called a "modified version" of the
earlier work or a work "based on" the earlier work.

  A "covered work" means either the unmodified Program or a work based
on the Program.

  To "propagate" a work means to do anything with it that, without
permission, would make you directly or secondarily liable for
infringement under applicable copyright law, except executing it on a
computer or modifying a private copy.  Propagation includes copying,
distribution (with or without modification), making available to the
public, and in some countries other activities as well.

  To "convey" a work means any kind of propagation that enables other
parties to make or receive copies.  Mere interaction with a user through
a computer network, with no transfer of a copy, is not conveying.

  An interactive user interface displays "Appropriate Legal Notices"
to the extent that it includes a convenient and prominently visible
feature that (1) displays an appropriate copyright notice, and (2)
tells the user that there is no warranty for the work (except to the
extent that warranties are provided), that licensees may convey the
work under this License, and how to view a copy of this License.  If
the interface presents a list of user commands or options, such as a
menu, a prominent item in the list meets this criterion.

  1. Source Code.

  The "source code" for a work means the preferred form of the work
for making modifications to it.  "Object code" means any non-source
form of a work.

  A "Standard Interface" means an interface that either is an official
standard defined by a recognized standards body, or, in the case of
interfaces specified for a particular programming language, one that
is widely used among developers working in that language.

  The "System Libraries" of an executable work include anything, other
than the work as a whole, that (a) is included in the normal form of
packaging a Major Component, but which is not part of that Major
Component, and (b) serves only to enable use of the work with that
Major Component, or to implement a Standard Interface for which an
implementation is available to the public in source code form.  A
"Major Component", in this context, means a major essential component
(kernel, window system, and so on) of the specific operating system
(if any) on which the executable work runs, or a compiler used to
produce the work, or an object code interpreter used to run it.

  The "Corresponding Source" for a work in object code form means all
the source code needed to generate, install, and (for an executable
work) run the object code and to modify the work, including scripts to
control those activities.  However, it does not include the work's
System Libraries, or general-purpose tools or generally available free
programs which are used unmodified in performing those activities but
which are not part of the work.  For example, Corresponding Source
includes interface definition files associated with source files for
the work, and the source code for shared libraries and dynamically
linked subprograms that the work is specifically designed to require,
such as by intimate data communication or control flow between those
subprograms and other parts of the work.

  The Corresponding Source need not include anything that users
can regenerate automatically from other parts of the Corresponding
Source.

  The Corresponding Source for a work in source code form is that
same work.

  2. Basic Permissions.

  All rights granted under this License are granted for the term of
copyright on the Program, and are irrevocable provided the stated
conditions are met.  This License explicitly affirms your unlimited
permission to run the unmodified Program.  The output from running a
covered work is covered by this License only if the output, given its
content, constitutes a covered work.  This License acknowledges your
rights of fair use or other equivalent, as provided by copyright law.

  You may make, run and propagate covered works that you do not
convey, without conditions so long as your license otherwise remains
in force.  You may convey covered works to others for the sole purpose
of having them make modifications exclusively for you, or provide you
with facilities for running those works, provided that you comply with
the terms of this License in conveying all material for which you do
not control copyright.  Those thus making or running the covered works
for you must do so exclusively on your behalf, under your direction
and control, on terms that prohibit them from making any copies of
your copyrighted material outside their relationship with you.

  Conveying under any other circumstances is permitted solely under
the conditions stated below.  Sublicensing is not allowed; section 10
makes it unnecessary.

  3. Protecting Users' Legal Rights From Anti-Circumvention Law.

  No covered work shall be deemed part of an effective technological
measure under any applicable law fulfilling obligations under article
11 of the WIPO copyright treaty adopted on 20 December 1996, or
similar laws prohibiting or restricting circumvention of such
measures.

  When you convey a covered work, you waive any legal power to forbid
circumvention of technological measures to the extent such circumvention
is effected by exercising rights under this License with respect to
the covered work, and you disclaim any intention to limit operation or
modification of the work as a means of enforcing, against the work's
users, your or third parties' legal rights to forbid circumvention of
technological measures.

  4. Conveying Verbatim Copies.

  You may convey verbatim copies of the Program's source code as you
receive it, in any medium, provided that you conspicuously and
appropriately publish on each copy an appropriate copyright notice;
keep intact all notices stating that this License and any
non-permissive terms added in accord with section 7 apply to the code;
keep intact all notices of the absence of any warranty; and give all
recipients a copy of this License along with the Program.

  You may charge any price or no price for each copy that you convey,
and you may offer support or warranty protection for a fee.

  5. Conveying Modified Source Versions.

  You may convey a work based on the Program, or the modifications to
produce it from the Program, in the form of source code under the
terms of section 4, provided that you also meet all of these conditions:

    a) The work must carry prominent notices stating that you modified
    it, and giving a relevant date.

    b) The work must carry prominent notices stating that it is
    released under this License and any conditions added under section
    7.  This requirement modifies the requirement in section 4 to
    "keep intact all notices".

    c) You must license the entire work, as a whole, under this
    License to anyone who comes into possession of a copy.  This
    License will therefore apply, along with any applicable section 7
    additional terms, to the whole of the work, and all its parts,
    regardless of how they are packaged.  This License gives no
    permission to license the work in any other way, but it does not
    invalidate such permission if you have separately received it.

    d) If the work has interactive user interfaces, each must display
    Appropriate Legal Notices; however, if the Program has interactive
    interfaces that do not display Appropriate Legal Notices, your
    work need not make them do so.

  A compilation of a covered work with other separate and independent
works, which are not by their nature extensions of the covered work,
and which are not combined with it such as to form a larger program,
in or on a volume of a storage or distribution medium, is called an
"aggregate" if the compilation and its resulting copyright are not
used to limit the access or legal rights of the compilation's users
beyond what the individual works permit.  Inclusion of a covered work
in an aggregate does not cause this License to apply to the other
parts of the aggregate.

  6. Conveying Non-Source Forms.

  You may convey a covered work in object code form under the terms
of sections 4 and 5, provided that you also convey the
machine-readable Corresponding Source under the terms of this License,
in one of these ways:

    a) Convey the object code in, or embodied in, a physical product
    (including a physical distribution medium), accompanied by the
    Corresponding Source fixed on a durable physical medium
    customarily used for software interchange.

    b) Convey the object code in, or embodied in, a physical product
    (including a physical distribution medium), accompanied by a
    written offer, valid for at least three years and valid for as
    long as you offer spare parts or customer support for that product
    model, to give anyone who possesses the object code either (1) a
    copy of the Corresponding Source for all the software in the
    product that is covered by this License, on a durable physical
    medium customarily used for software interchange, for a price no
    more than your reasonable cost of physically performing this
    conveying of source, or (2) access to copy the
    Corresponding Source from a network server at no charge.

    c) Convey individual copies of the object code with a copy of the
    written offer to provide the Corresponding Source.  This
    alternative is allowed only occasionally and noncommercially, and
    only if you received the object code with such an offer, in accord
    with subsection 6b.

    d) Convey the object code by offering access from a designated
    place (gratis or for a charge), and offer equivalent access to the
    Corresponding Source in the same way through the same place at no
    further charge.  You need not require recipients to copy the
    Corresponding Source along with the object code.  If the place to
    copy the object code is a network server, the Corresponding Source
    may be on a different server (operated by you or a third party)
    that supports equivalent copying facilities, provided you maintain
    clear directions next to the object code saying where to find the
    Corresponding Source.  Regardless of what server hosts the
    Corresponding Source, you remain obligated to ensure that it is
    available for as long as needed to satisfy these requirements.

    e) Convey the object code using peer-to-peer transmission, provided
    you inform other peers where the object code and Corresponding
    Source of the work are being offered to the general public at no
    charge under subsection 6d.

  A separable portion of the object code, whose source code is excluded
from the Corresponding Source as a System Library, need not be
included in conveying the object code work.

  A "User Product" is either (1) a "consumer product", which means any
tangible personal property which is normally used for personal, family,
or household purposes, or (2) anything designed or sold for incorporation
into a dwelling.  In determining whether a product is a consumer product,
doubtful cases shall be resolved in favor of coverage.  For a particular
product received by a particular user, "normally used" refers to a
typical or common use of that class of product, regardless of the status
of the particular user or of the way in which the particular user
actually uses, or expects or is expected to use, the product.  A product
is a consumer product regardless of whether the product has substantial
commercial, industrial or non-consumer uses, unless such uses represent
the only significant mode of use of the product.

  "Installation Information" for a User Product means any methods,
procedures, authorization keys, or other information required to install
and execute modified versions of a covered work in that User Product from
a modified version of its Corresponding Source.  The information must
suffice to ensure that the continued functioning of the modified object
code is in no case prevented or interfered with solely because
modification has been made.

  If you convey an object code work under this section in, or with, or
specifically for use in, a User Product, and the conveying occurs as
part of a transaction in which the right of possession and use of the
User Product is transferred to the recipient in perpetuity or for a
fixed term (regardless of how the transaction is characterized), the
Corresponding Source conveyed under this section must be accompanied
by the Installation Information.  But this requirement does not apply
if neither you nor any third party retains the ability to install
modified object code on the User Product (for example, the work has
been installed in ROM).

  The requirement to provide Installation Information does not include a
requirement to continue to provide support service, warranty, or updates
for a work that has been modified or installed by the recipient, or for
the User Product in which it has been modified or installed.  Access to a
network may be denied when the modification itself materially and
adversely affects the operation of the network or violates the rules and
protocols for communication across the network.

  Corresponding Source conveyed, and Installation Information provided,
in accord with this section must be in a format that is publicly
documented (and with an implementation available to the public in
source code form), and must require no special password or key for
unpacking, reading or copying.

  7. Additional Terms.

  "Additional permissions" are terms that supplement the terms of this
License by making exceptions from one or more of its conditions.
Additional permissions that are applicable to the entire Program shall
be treated as though they were included in this License, to the extent
that they are valid under applicable law.  If additional permissions
apply only to part of the Program, that part may be used separately
under those permissions, but the entire Program remains governed by
this License without regard to the additional permissions.

  When you convey a copy of a covered work, you may at your option
remove any additional permissions from that copy, or from any part of
it.  (Additional permissions may be written to require their own
removal in certain cases when you modify the work.)  You may place
additional permissions on material, added by you to a covered work,
for which you have or can give appropriate copyright permission.

  Notwithstanding any other provision of this License, for material you
add to a covered work, you may (if authorized by the copyright holders of
that material) supplement the terms of this License with terms:

    a) Disclaiming warranty or limiting liability differently from the
    terms of sections 15 and 16 of this License; or

    b) Requiring preservation of specified reasonable legal notices or
    author attributions in that material or in the Appropriate Legal
    Notices displayed by works containing it; or

    c) Prohibiting misrepresentation of the origin of that material, or
    requiring that modified versions of such material be marked in
    reasonable ways as different from the original version; or

    d) Limiting the use for publicity purposes of names of licensors or
    authors of the material; or

    e) Declining to grant rights under trademark law for use of some
    trade names, trademarks, or service marks; or

    f) Requiring indemnification of licensors and authors of that
    material by anyone who conveys the material (or modified versions of
    it) with contractual assumptions of liability to the recipient, for
    any liability that these contractual assumptions directly impose on
    those licensors and authors.

  All other non-permissive additional terms are considered "further
restrictions" within the meaning of section 10.  If the Program as you
received it, or any part of it, contains a notice stating that it is
governed by this License along with a term that is a further
restriction, you may remove that term.  If a license document contains
a further restriction but permits relicensing or conveying under this
License, you may add to a covered work material governed by the terms
of that license document, provided that the further restriction does
not survive such relicensing or conveying.

  If you add terms to a covered work in accord with this section, you
must place, in the relevant source files, a statement of the
additional terms that apply to those files, or a notice indicating
where to find the applicable terms.

  Additional terms, permissive or non-permissive, may be stated in the
form of a separately written license, or stated as exceptions;
the above requirements apply either way.

  8. Termination.

  You may not propagate or modify a covered work except as expressly
provided under this License.  Any attempt otherwise to propagate or
modify it is void, and will automatically terminate your rights under
this License (including any patent licenses granted under the third
paragraph of section 11).

  However, if you cease all violation of this License, then your
license from a particular copyright holder is reinstated (a)
provisionally, unless and until the copyright holder explicitly and
finally terminates your license, and (b) permanently, if the copyright
holder fails to notify you of the violation by some reasonable means
prior to 60 days after the cessation.

  Moreover, your license from a particular copyright holder is
reinstated permanently if the copyright holder notifies you of the
violation by some reasonable means, this is the first time you have
received notice of violation of this License (for any work) from that
copyright holder, and you cure the violation prior to 30 days after
your receipt of the notice.

  Termination of your rights under this section does not terminate the
licenses of parties who have received copies or rights from you under
this License.  If your rights have been terminated and not permanently
reinstated, you do not qualify to receive new licenses for the same
material under section 10.

  9. Acceptance Not Required for Having Copies.

  You are not required to accept this License in order to receive or
run a copy of the Program.  Ancillary propagation of a covered work
occurring solely as a consequence of using peer-to-peer transmission
to receive a copy likewise does not require acceptance.  However,
nothing other than this License grants you permission to propagate or
modify any covered work.  These actions infringe copyright if you do
not accept this License.  Therefore, by modifying or propagating a
covered work, you indicate your acceptance of this License to do so.

  10. Automatic Licensing of Downstream Recipients.

  Each time you convey a covered work, the recipient automatically
receives a license from the original licensors, to run, modify and
propagate that work, subject to this License.  You are not responsible
for enforcing compliance by third parties with this License.

  An "entity transaction" is a transaction transferring control of an
organization, or substantially all assets of one, or subdividing an
organization, or merging organizations.  If propagation of a covered
work results from an entity transaction, each party to that
transaction who receives a copy of the work also receives whatever
licenses to the work the party's predecessor in interest had or could
give under the previous paragraph, plus a right to possession of the
Corresponding Source of the work from the predecessor in interest, if
the predecessor has it or can get it with reasonable efforts.

  You may not impose any further restrictions on the exercise of the
rights granted or affirmed under this License.  For example, you may
not impose a license fee, royalty, or other charge for exercise of
rights granted under this License, and you may not initiate litigation
(including a cross-claim or counterclaim in a lawsuit) alleging that
any patent claim is infringed by making, using, selling, offering for
sale, or importing the Program or any portion of it.

  11. Patents.

  A "contributor" is a copyright holder who authorizes use under this
License of the Program or a work on which the Program is based.  The
work thus licensed is called the contributor's "contributor version".

  A contributor's "essential patent claims" are all patent claims
owned or controlled by the contributor, whether already acquired or
hereafter acquired, that would be infringed by some manner, permitted
by this License, of making, using, or selling its contributor version,
but do not include claims that would be infringed only as a
consequence of further modification of the contributor version.  For
purposes of this definition, "control" includes the right to grant
patent sublicenses in a manner consistent with the requirements of
this License.

  Each contributor grants you a non-exclusive, worldwide, royalty-free
patent license under the contributor's essential patent claims, to
make, use, sell, offer for sale, import and otherwise run, modify and
propagate the contents of its contributor version.

  In the following three paragraphs, a "patent license" is any express
agreement or commitment, however denominated, not to enforce a patent
(such as an express permission to practice a patent or covenant not to
sue for patent infringement).  To "grant" such a patent license to a
party means to make such an agreement or commitment not to enforce a
patent against the party.

  If you convey a covered work, knowingly relying on a patent license,
and the Corresponding Source of the work is not available for anyone
to copy, free of charge and under the terms of this License, through a
publicly available network server or other readily accessible means,
then you must either (1) cause the Corresponding Source to be so
available, or (2) arrange to deprive yourself of the benefit of the
patent license for this particular work, or (3) arrange, in a manner
consistent with the requirements of this License, to extend the patent
license to downstream recipients.  "Knowingly relying" means you have
actual knowledge that, but for the patent license, your conveying the
covered work in a country, or your recipient's use of the covered work
in a country, would infringe one or more identifiable patents in that
country that you have reason to believe are valid.

  If, pursuant to or in connection with a single transaction or
arrangement, you convey, or propagate by procuring conveyance of, a
covered work, and grant a patent license to some of the parties
receiving the covered work authorizing them to use, propagate, modify
or convey a specific copy of the covered work, then the patent license
you grant is automatically extended to all recipients of the covered
work and works based on it.

  A patent license is "discriminatory" if it does not include within
the scope of its coverage, prohibits the exercise of, or is
conditioned on the non-exercise of one or more of the rights that are
specifically granted under this License.  You may not convey a covered
work if you are a party to an arrangement with a third party that is
in the business of distributing software, under which you make payment
to the third party based on the extent of your activity of conveying
the work, and under which the third party grants, to any of the
parties who would receive the covered work from you, a discriminatory
patent license (a) in connection with copies of the covered work
conveyed by you (or copies made from those copies), or (b) primarily
for and in connection with specific products or compilations that
contain the covered work, unless you entered into that arrangement,
or that patent license was granted, prior to 28 March 2007.

  Nothing in this License shall be construed as excluding or limiting
any implied license or other defenses to infringement that may
otherwise be available to you under applicable patent law.

  12. No Surrender of Others' Freedom.

  If conditions are imposed on you (whether by court order, agreement or
otherwise) that contradict the conditions of this License, they do not
excuse you from the conditions of this License.  If you cannot convey a
covered work so as to satisfy simultaneously your obligations under this
License and any other pertinent obligations, then as a consequence you may
not convey it at all.  For example, if you agree to terms that obligate you
to collect a royalty for further conveying from those to whom you convey
the Program, the only way you could satisfy both those terms and this
License would be to refrain entirely from conveying the Program.

  13. Use with the GNU Affero General Public License.

  Notwithstanding any other provision of this License, you have
permission to link or combine any covered work with a work licensed
under version 3 of the GNU Affero General Public License into a single
combined work, and to convey the resulting work.  The terms of this
License will continue to apply to the part which is the covered work,
but the special requirements of the GNU Affero General Public License,
section 13, concerning interaction through a network will apply to the
combination as such.

  14. Revised Versions of this License.

  The Free Software Foundation may publish revised and/or new versions of
the GNU General Public License from time to time.  Such new versions will
be similar in spirit to the present version, but may differ in detail to
address new problems or concerns.

  Each version is given a distinguishing version number.  If the
Program specifies that a certain numbered version of the GNU General
Public License "or any later version" applies to it, you have the
option of following the terms and conditions either of that numbered
version or of any later version published by the Free Software
Foundation.  If the Program does not specify a version number of the
GNU General Public License, you may choose any version ever published
by the Free Software Foundation.

  If the Program specifies that a proxy can decide which future
versions of the GNU General Public License can be used, that proxy's
public statement of acceptance of a version permanently authorizes you
to choose that version for the Program.

  Later license versions may give you additional or different
permissions.  However, no additional obligations are imposed on any
author or copyright holder as a result of your choosing to follow a
later version.

  15. Disclaimer of Warranty.

  THERE IS NO WARRANTY FOR THE PROGRAM, TO THE EXTENT PERMITTED BY
APPLICABLE LAW.  EXCEPT WHEN OTHERWISE STATED IN WRITING THE COPYRIGHT
HOLDERS AND/OR OTHER PARTIES PROVIDE THE PROGRAM "AS IS" WITHOUT WARRANTY
OF ANY KIND, EITHER EXPRESSED OR IMPLIED, INCLUDING, BUT NOT LIMITED TO,
THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR
PURPOSE.  THE ENTIRE RISK AS TO THE QUALITY AND PERFORMANCE OF THE PROGRAM
IS WITH YOU.  SHOULD THE PROGRAM PROVE DEFECTIVE, YOU ASSUME THE COST OF
ALL NECESSARY SERVICING, REPAIR OR CORRECTION.

  16. Limitation of Liability.

  IN NO EVENT UNLESS REQUIRED BY APPLICABLE LAW OR AGREED TO IN WRITING
WILL ANY COPYRIGHT HOLDER, OR ANY OTHER PARTY WHO MODIFIES AND/OR CONVEYS
THE PROGRAM AS PERMITTED ABOVE, BE LIABLE TO YOU FOR DAMAGES, INCLUDING ANY
GENERAL, SPECIAL, INCIDENTAL OR CONSEQUENTIAL DAMAGES ARISING OUT OF THE
USE OR INABILITY TO USE THE PROGRAM (INCLUDING BUT NOT LIMITED TO LOSS OF
DATA OR DATA BEING RENDERED INACCURATE OR LOSSES SUSTAINED BY YOU OR THIRD
PARTIES OR A FAILURE OF THE PROGRAM TO OPERATE WITH ANY OTHER PROGRAMS),
EVEN IF SUCH HOLDER OR OTHER PARTY HAS BEEN ADVISED OF THE POSSIBILITY OF
SUCH DAMAGES.

  17. Interpretation of Sections 15 and 16.

  If the disclaimer of warranty and limitation of liability provided
above cannot be given local legal effect according to their terms,
reviewing courts shall apply local law that most closely approximates
an absolute waiver of all civil liability in connection with the
Program, unless a warranty or assumption of liability accompanies a
copy of the Program in return for a fee.

                     END OF TERMS AND CONDITIONS

            How to Apply These Terms to Your New Programs

  If you develop a new program, and you want it to be of the greatest
possible use to the public, the best way to achieve this is to make it
free software which everyone can redistribute and change under these terms.

  To do so, attach the following notices to the program.  It is safest
to attach them to the start of each source file to most effectively
state the exclusion of warranty; and each file should have at least
the "copyright" line and a pointer to where the full notice is found.

    <one line to give the program's name and a brief idea of what it does.>
    Copyright (C) <year>  <name of author>

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.

Also add information on how to contact you by electronic and paper mail.

  If the program does terminal interaction, make it output a short
notice like this when it starts in an interactive mode:

    <program>  Copyright (C) <year>  <name of author>
    This program comes with ABSOLUTELY NO WARRANTY; for details type `show w'.
    This is free software, and you are welcome to redistribute it
    under certain conditions; type `show c' for details.

The hypothetical commands `show w' and `show c' should show the appropriate
parts of the General Public License.  Of course, your program's commands
might be different; for a GUI interface, you would use an "about box".

  You should also get your employer (if you work as a programmer) or school,
if any, to sign a "copyright disclaimer" for the program, if necessary.
For more information on this, and how to apply and follow the GNU GPL, see
<https://www.gnu.org/licenses/>.

  The GNU General Public License does not permit incorporating your program
into proprietary programs.  If your program is a subroutine library, you
may consider it more useful to permit linking proprietary applications with
the library.  If this is what you want to do, use the GNU Lesser General
Public License instead of this License.  But first, please read
<https://www.gnu.org/licenses/why-not-lgpl.html>.
```

## File: README.en.md
```markdown
<h1 align="center">ChatGPT Web Share</h1>

<div align="center">

[![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/moeakwak/chatgpt-web-share?label=container&logo=docker)](https://github.com/moeakwak/chatgpt-web-share/pkgs/container/chatgpt-web-share)
[![Github Workflow Status](https://img.shields.io/github/actions/workflow/status/moeakwak/chatgpt-web-share/docker-image.yml?label=build)](https://github.com/moeakwak/chatgpt-web-share/actions)
[![License](https://img.shields.io/github/license/moeakwak/chatgpt-web-share)](https://github.com/moeakwak/chatgpt-web-share/blob/main/LICENSE)

A web application that allows multiple users to share a ChatGPT account at the same time, developed using FastAPI and Vue3.

Used for sharing one ChatGPT account among friends.

</div>

![screenshot](docs/screenshot.en.jpeg)

This readme was mainly translated by ChatGPT.

## About the project

ChatGPT Web Share (CWS for short) is designed to share a ChatGPT Plus account with multiple users. CWS:
- is a front-end and back-end separated application
- is used to share a ChatGPT account, not the official API
- supports user and conversation managements
- prioritizes support for ChatGPT Plus accounts

## Features

- A beautiful and concise web interface using [naive-ui](https://www.naiveui.com/)
  - Supports English language
  - Supports switching to dark mode
  - Supports copying reply content or code content with one click
  - Supports displaying images, tables, mathematical formulas, and code highlighting in replies
  - Supports exporting conversations as beautiful Markdown or PDF files
  - Replying content in typing animation
  - Supports stopping generation
- Multiple users can share the same ChatGPT account
  - Different users' ChatGPT conversations are separated and do not affect each other
  - When multiple users request at the same time, they will be queued for processing
  - Administrators can set users' maximum number of conversations, conversation time limits, etc.
  - Provides real-time updated service usage status to avoid usage peaks
- Comprehensive management functions
  - Modify user conversation restrictions
  - Manage conversations/view member conversation records/assign conversations to specific users
  - View logs in real-time
  - Record request and conversation statistics

## Deploy Guide

Please see the WIKI: [English Guide](https://github.com/moeakwak/chatgpt-web-share/wiki/English-Guide)

## Usage Statement

### Information Collection and Privacy Statement

<del>Starting from version v0.2.16, this project uses Sentry to collect error information. By using this project, you agree to the Sentry privacy policy. Any anonymous information collected through Sentry will only be used for development and debugging purposes. </del>We will never collect or store any of your private data, like username, password, access token, etc.

From v0.3.5, Sentry is not used anymore.

### Risk Statement

This project is for learning and research purposes only, and commercial use is not encouraged. We are not responsible for any losses caused by using this project.
```

## File: README.md
```markdown
<h1 align="center">ChatGPT Web Share</h1>

<div align="center">

[![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/moeakwak/chatgpt-web-share?label=container&logo=docker)](https://github.com/moeakwak/chatgpt-web-share/pkgs/container/chatgpt-web-share)
[![Github Workflow Status](https://img.shields.io/github/actions/workflow/status/moeakwak/chatgpt-web-share/docker-image.yml?label=build)](https://github.com/moeakwak/chatgpt-web-share/actions)
[![License](https://img.shields.io/github/license/moeakwak/chatgpt-web-share)](https://github.com/moeakwak/chatgpt-web-share/blob/main/LICENSE)

**>>> [English Readme](README.en.md) <<<**

共享一个 ChatGPT 账号给多用户同时使用的 web 应用，使用 FastAPI + Vue3 开发。可用于朋友之间共享或合租 ChatGPT 账号。支持 ChatGPT Plus / 设置对话模型 / 用户请求限制等功能。支持使用 GPT-4！

</div>

![screenshot](docs/screenshot.jpeg)

![screenshot_admin](docs/screenshot_admin.jpeg)

通知/讨论 Channel：https://t.me/chatgptwebshare

## 关于项目

ChatGPT Web Share (简称 CWS) 的目的是「共享」一个 ChatGPT Plus 账号给多个用户。CWS 是：
- 前后端分离的应用，因此你需要自行部署后端到一个稳定且 IP 可靠的服务器上
- 用于共享 ChatGPT 账号，而不是官方 API
- 支持用户管理，并支持设置各用户的权限和对话次数
- 优先支持 ChatGPT Plus 账号

## 特点

- 美观简洁的 web 界面，使用 [naive-ui](https://www.naiveui.com/)
  - 多语言（简体中文、英语）支持
  - 适配夜间模式
  - 支持一键复制回复内容或代码内容
  - 支持显示回复中的图像/表格/数学公式/代码语法高亮
  - 一键导出对话为美观的 Markdown 或 PDF 文件
  - 动态显示回复内容
  - 支持停止生成对话
- 多用户共享管理
  - 创建多用户用于共享一个 ChatGPT 账号
  - 不同用户创建的 ChatGPT 对话互相分隔，不会相互影响
  - 多用户同时请求时，会进行排队处理
  - 管理员可设置用户的最大对话数量、对话次数限制等
  - 提供实时更新的服务使用状态，从而能够避开使用高峰
- 完善的管理功能
  - 修改用户对话限制
  - 管理对话/查看成员对话记录/分配对话给特定用户
  - 实时查看日志
  - 记录请求及对话统计信息

## 部署指南

参见 WIKI：[中文指南](https://github.com/moeakwak/chatgpt-web-share/wiki/%E4%B8%AD%E6%96%87%E6%8C%87%E5%8D%97)

## 声明

### 调试信息收集和隐私声明

<del>从版本 v0.2.16 开始，本项目使用 Sentry 来收集错误信息。使用本项目即表示您同意 Sentry 的隐私政策。通过 Sentry 收集的任何匿名信息仅用于开发和调试目的。</del>我们永远不会收集或存储您的私人数据，如用户名、密码、access token 等。

目前，自 v0.3.5 版本后已不再通过 Sentry 收集错误信息。

### 风险声明

本项目仅供学习和研究使用，不鼓励用于商业用途。我们不对任何因使用本项目而导致的任何损失负责。

## 捐助和支持

如果觉得本项目对您有帮助，欢迎通过扫描下方赞赏码捐助项目 :)

<img src="docs/donate.png" alt="donate" width="200" height="200" />
```

## File: startup.sh
```bash
#!/bin/sh
cd /app
caddy start --config /app/Caddyfile --adapter caddyfile
cd /app/backend
# python main.py
exec uvicorn main:app --host 0.0.0.0 --port 8000 --proxy-headers --forwarded-allow-ips '*' --log-config logging_config.yaml
```
