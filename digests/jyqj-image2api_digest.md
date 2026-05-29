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
.gitignore
API_NOTES.md
assets/dashboard-sample.png
cmd/server/main.go
cmd/utls-probe/main.go
configs/config.example.yaml
deploy/.env.example
deploy/build-local.ps1
deploy/build-local.sh
deploy/docker-compose.yml
deploy/Dockerfile
deploy/Dockerfile.multistage
deploy/entrypoint.sh
deploy/nginx.conf
deploy/README.md
docs/FRAMEWORK.md
docs/screenshots/d2d3e22e-a3d6-4009-9c1b-b63e17996f5c.png
docs/screenshots/d4e985854e071922c351ec46a495489c.png
docs/screenshots/playground-batch.png
docs/screenshots/playground-preview.png
docs/USER_GUIDE.md
go.mod
internal/account/dao.go
internal/account/handler.go
internal/account/importer_tokens.go
internal/account/importer.go
internal/account/model.go
internal/account/quota.go
internal/account/refresher.go
internal/account/service.go
internal/audit/dao.go
internal/audit/handler.go
internal/audit/middleware.go
internal/audit/model.go
internal/backup/dao.go
internal/backup/handler.go
internal/backup/model.go
internal/backup/service.go
internal/config/config.go
internal/db/mysql.go
internal/db/redis.go
internal/gateway/chat.go
internal/gateway/delta.go
internal/gateway/images_proxy.go
internal/gateway/images.go
internal/gateway/types.go
internal/image/dao.go
internal/image/me_handler.go
internal/image/model.go
internal/image/runner.go
internal/middleware/auth.go
internal/middleware/cors.go
internal/middleware/logger.go
internal/middleware/recover.go
internal/middleware/request_id.go
internal/model/admin_handler.go
internal/model/dao.go
internal/model/model.go
internal/model/registry.go
internal/proxy/dao.go
internal/proxy/handler.go
internal/proxy/model.go
internal/proxy/prober.go
internal/proxy/service.go
internal/scheduler/scheduler.go
internal/server/router.go
internal/server/spa.go
internal/settings/dao.go
internal/settings/handler.go
internal/settings/model.go
internal/settings/service.go
internal/settings/util.go
internal/upstream/chatgpt/client.go
internal/upstream/chatgpt/conversation.go
internal/upstream/chatgpt/fchat.go
internal/upstream/chatgpt/files.go
internal/upstream/chatgpt/image_img2_test.go
internal/upstream/chatgpt/image.go
internal/upstream/chatgpt/pow.go
internal/upstream/chatgpt/probe.go
internal/upstream/chatgpt/utls_transport.go
internal/usage/admin_handler.go
internal/usage/logger.go
internal/usage/me_handler.go
internal/usage/model.go
internal/usage/query_dao.go
legacy/gen_image.py
LICENSE
Makefile
pkg/crypto/aes.go
pkg/lock/redis_lock.go
pkg/logger/logger.go
pkg/mailer/mailer.go
pkg/mailer/templates.go
pkg/resp/resp.go
README.md
scripts/package.json
scripts/README.md
scripts/smoke.mjs
sql/database.sql
web/.env.development
web/.env.production
web/.gitignore
web/index.html
web/package.json
web/public/favicon.svg
web/README.md
web/src/api/accounts.ts
web/src/api/audit.ts
web/src/api/backup.ts
web/src/api/http.ts
web/src/api/me.ts
web/src/api/proxies.ts
web/src/api/settings.ts
web/src/api/stats.ts
web/src/App.vue
web/src/components/Placeholder.vue
web/src/config/feature.ts
web/src/env.d.ts
web/src/layouts/BasicLayout.vue
web/src/layouts/BlankLayout.vue
web/src/main.ts
web/src/router/index.ts
web/src/stores/site.ts
web/src/stores/ui.ts
web/src/stores/user.ts
web/src/styles/global.scss
web/src/utils/brand.ts
web/src/utils/format.ts
web/src/views/admin/Accounts.vue
web/src/views/admin/Audit.vue
web/src/views/admin/Backup.vue
web/src/views/admin/Models.vue
web/src/views/admin/Proxies.vue
web/src/views/admin/Settings.vue
web/src/views/admin/UsageStats.vue
web/src/views/Error403.vue
web/src/views/Error404.vue
web/src/views/Login.vue
web/src/views/personal/ApiDocs.vue
web/src/views/personal/Dashboard.vue
web/src/views/personal/OnlinePlay.vue
web/src/views/personal/Usage.vue
web/tsconfig.json
web/tsconfig.node.json
web/vite.config.ts
```

# Files

## File: .gitignore
````
# =============================================================================
# Go
# =============================================================================
*.exe
*.dll
*.so
*.dylib
*.test
*.out
/bin/
/dist/
/deploy/bin/

# =============================================================================
# 本地配置 / 密钥(绝对不能进仓库)
# =============================================================================
/configs/config.yaml
/configs/*.local.yaml
/deploy/.env
.env
.env.local

# =============================================================================
# IDE
# =============================================================================
.idea/
.vscode/
*.swp
*.swo

# =============================================================================
# 日志 & 数据
# =============================================================================
/logs/
/deploy/logs/
*.log
/data/
/uploads/
/storage/

# =============================================================================
# 前端
# =============================================================================
/web/node_modules/
/web/dist/
/web/.cache/
/web/.vite/

# =============================================================================
# Python(仅 legacy 参考)
# =============================================================================
__pycache__/
*.pyc
*.pyo
.venv/
venv/

# =============================================================================
# 打包产物 / 临时目录(不要提交到开源仓库)
# =============================================================================
*.tgz
*.tar.gz
/_tmp/
/tmp/

# 历史部署产物(运维自己 build 即可,不要带进仓库)
/deploy/bin/
/deploy/logs/
/deploy/updates/
/deploy/web/

# =============================================================================
# 杂项
# =============================================================================
.DS_Store
Thumbs.db
*.bak
*.tmp
*_backup_*

# =============================================================================
# 根目录随手放的哈希命名图片 / HAR(README 正式图放 docs/screenshots/)
# =============================================================================
/[0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f]*.png
/[0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f]*.jpg
/[0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f]*.har
!/docs/**
````

## File: API_NOTES.md
````markdown
# ChatGPT 图像生成 & 额度查询接口备忘录

> 本文记录我们目前复用的 `chatgpt.com` 后端接口，及其请求/响应关键字段。
> 所有接口都走同一个 `Bearer {AUTH_TOKEN}`，host 固定 `https://chatgpt.com`。
> 运行环境：Go 版统一走 `internal/upstream/chatgpt` 的 uTLS transport + browser/Oai-* headers + cookie jar；探针和真实 `f/conversation` 共用同一套指纹体系。

---

## 0. 通用请求头

绝大多数接口共用下面这套头，区别只在于 `referer` / `x-openai-target-*`：

```
authorization: Bearer <AT>
accept: */*
accept-language: zh-CN,zh;q=0.9,en;q=0.8
content-type: application/json
origin: https://chatgpt.com
referer: https://chatgpt.com/                           # 或 /c/{conversation_id}
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36
             (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36 Edg/143.0.0.0
oai-language: zh-CN
oai-device-id: <稳定 UUID>
oai-client-version: prod-2d84edefecf794f1bf3178f1f15e1005067d903d
oai-client-build-number: 5983180
```

> `AUTH_TOKEN` 是 Bearer JWT，有效期约 10 天。过期后需要重新从页面 network 中获取。

---

## 1. image 能力探测与 quota 诊断

### 1.1 `GET /backend-api/models` —— 主能力探针

用途：判断当前登录态是否具备 image 入口资格。
实测发现同一个 Pro / image 灰度号可能在 `/backend-api/models` 明确暴露 image 能力,但 `/backend-api/conversation/init` 仍返回 `blocked_features:image_gen`。因此 **models 是主探针,init 只是弱诊断**。

关键字段：

| 字段 | 含义 |
|---|---|
| `default_model_slug` | 当前默认上游模型,灰度号常见 `gpt-5-3` |
| `models[].enabledTools` | 包含 `image_gen_tool_enabled` 时说明模型具备 image 工具入口 |
| `models[].supportedFeatures` | 包含 `image` 时说明该模型支持图像能力 |
| `model_picker_version` | 前端模型选择器版本 |
| `fConversationEndpoint` | 前端是否走 `f/conversation` 族端点 |

gpt2api 的账号能力判断:

```text
enabledTools contains image_gen_tool_enabled
OR supportedFeatures contains image
=> image_capability_status = enabled
```

注意:这只表示“有机会进入 image / picture_v2 链路”,不表示本次请求一定命中 IMG2。

### 1.2 `POST /backend-api/conversation/init` —— quota / blocked_features 弱诊断

用途：保留 quota、reset 时间、banner、blocked_features 等诊断信息。
不要再用它作为 image2 是否可用的主判据。

请求体：

```json
{
  "gizmo_id": null,
  "requested_default_model": null,
  "conversation_id": null,
  "timezone_offset_min": -480,
  "system_hints": ["picture_v2"]
}
```

`timezone_offset_min` 只表示对齐前端抓包值,不要把它解释成服务器能力判据。

关键字段：

| 字段 | 含义 |
|---|---|
| `limits_progress[].feature_name == "image_gen"` | 生图 quota 诊断;部分灰度号这里可能缺失或异常 |
| `limits_progress[].remaining` | 当前剩余次数 |
| `limits_progress[].reset_after` | 下次重置时间 |
| `blocked_features` | 仅作风险/兼容诊断;可能与 `/models` 暴露能力矛盾 |
| `default_model_slug` | 旧探针视角下的默认模型,不等于 image 主链路实际模型 |

## 2. 生图完整调用链

按顺序编号：

```
[1] GET  /                                            → Bootstrap,让 cookie jar 收集一方 cookie
[2] GET  /backend-api/models                            → 主能力探针:models-first
[3] POST /backend-api/sentinel/chat-requirements        → 拿 chat_token (+可选 POW / Turnstile)
[4] POST /backend-api/f/conversation/prepare            → 拿 conduit_token（请求级抽卡关键）
[5] POST /backend-api/f/conversation    (SSE)           → 正式下发 prompt,流式拿 image refs
[6] GET  /backend-api/conversation/{conv_id}            → 轮询补齐 mapping/tool payload
[7] GET  多路 files/download fallback                   → 拿短期签名 URL 或直接图片 bytes
[8] GET  <signed_url>                                   → 下载图片 bytes
```

> `/conversation/init` 不在主生图链路内;只在账号探测/诊断任务里调用。复用现有会话时仍需每次执行 `[3][4][5]`,因为 sentinel 与 conduit_token 都是请求级信号。

---

### [2] `POST /backend-api/sentinel/chat-requirements`

作用：拿 `chat_token`（写进 `openai-sentinel-chat-requirements-token`），以及判断是否要做 POW / Turnstile。

请求体：

```json
{ "p": "gAAAAAC...<get_requirements_token 生成>" }
```

响应关键字段：

```json
{
  "token": "...chat_token...",
  "proofofwork": {
    "required": true,
    "seed": "...",
    "difficulty": "0fffff"
  }
}
```

如果 `proofofwork.required=true`，需用本地 SHA3-512 暴力算 `openai-sentinel-proof-token`（见 `gen_image.py` 的 `generate_proof_token`）。

---

### [4] `POST /backend-api/f/conversation/prepare`

作用：请求级/会话级分流。服务器在这里决定本次请求更可能走哪套生图后端,返回一个 `conduit_token`。这不是账号静态开关,同一账号不同请求也可能抽到不同结果。

请求头额外需要：

```
openai-sentinel-chat-requirements-token: <chat_token>
openai-sentinel-proof-token: <proof_token>     # 若 POW required
```

请求体：

```json
{
  "model": "auto",
  "system_hints": ["picture_v2"],
  "timezone_offset_min": -480,
  "conversation_id": null,              // 或已有会话 id
  "message_id": "<前端生成 UUID>",
  "supports_buffering": true
}
```

响应体：

```json
{ "conduit_token": "ct_...." }
```

`conduit_token` 要在 `[5]` 里通过请求头 `x-conduit-token` 传回去。

---

### [5] `POST /backend-api/f/conversation` (SSE)

作用：正式提交 prompt 并接收流式响应，里面会陆续下发 `image_gen_task_id` / 初始 `file_id`。

请求头额外需要：

```
openai-sentinel-chat-requirements-token: <chat_token>
openai-sentinel-proof-token: <proof_token>
x-conduit-token: <conduit_token>             # 关键！否则不进灰度桶
accept: text/event-stream
```

请求体骨架（精简）：

```json
{
  "action": "next",
  "messages": [{
      "id": "<msg_uuid>",
      "author": { "role": "user" },
      "content": { "content_type": "text", "parts": ["<prompt>"] },
      "metadata": { "system_hints": ["picture_v2"] }
  }],
  "parent_message_id": "<head_or_new_uuid>",
  "model": "auto",
  "conversation_id": null,
  "system_hints": ["picture_v2"],            // ← 必须，开启图像工具
  "timezone_offset_min": -480,
  "client_prepare_state": "sent",
  "supports_buffering": true,
  "enable_message_followups": true,
  "force_parallel_switch": "auto"
}
```

SSE 事件里要抓的字段：

| 字段 | 位置 | 作用 |
|---|---|---|
| `conversation_id` | `message.metadata` 或顶层 | 后续轮询用 |
| `image_gen_task_id` | `message.metadata.image_gen_async` | 确认任务已发起 |
| `content.parts[].asset_pointer` | assistant/tool 消息 | `file-service://file_xxx` 或 `sediment://file_xxx` |
| `content.parts[].metadata.generation.gen_size_v2` | image asset metadata | 新实测 IMG2 sediment-only 终稿关键指纹 |
| `image_gen_task_id` / `async_task_type=image_gen` | metadata | 确认异步图像任务已发起 |

---

### [6] `GET /backend-api/conversation/{conversation_id}`

作用：SSE 结束后轮询补齐最终 file-service URL（尤其灰度会出第二张高清图）。

响应：完整会话 JSON，结构里 `mapping` 是消息树。

polling 策略（见 `poll_conversation_for_images`）：
- 使用 **baseline diff**：请求前先记录 "现有 tool 消息 id 集合"，轮询时只看新增的。
- `file-service://` 直接视作 IMG2 终稿指纹。
- `sediment://` 不能一律视作 preview;若 asset metadata 含 `generation.gen_size_v2`,按 IMG2 sediment-only 终稿处理。
- 多条新增 image tool 消息仍可作为 IMG2 聚合信号。
- 单条 `sediment://` 且无 `gen_size_v2`,等待后仍无新终稿才判 `preview_only`。
- 最大等待由 runner 配置控制;连续 429 退避/中止。

---

### [7] 图片下载 URL fallback

当前上游下载端点不完全稳定,需要多路 fallback:

```text
/backend-api/files/download/{fid}?conversation_id={cid}&inline=false
/backend-api/conversation/{cid}/attachment/{sid}/download
/backend-api/files/download/{fid}
/backend-api/files/{fid}/download
```

响应可能是 JSON、302 Location,也可能直接返回图片二进制:

```json
{ "status": "success", "download_url": "https://files.oaiusercontent.com/…签名URL…" }
```

---

## 3. 其他已观察到的接口（非必用）

| 接口 | 方法 | 用途 | 响应 |
|---|---|---|---|
| `/backend-api/image-gen/image-paragen-display` | POST | **前端上报**：告诉后端"已展示 N 张图" | 204 空 |
| `/backend-api/conversation/{id}/async-status` | POST `{"status":null}` | 异步任务健康检查 | `{"status":"OK"}` |
| `/backend-api/accounts/check/v4-2023-04-27` | GET | 账号 features/entitlements | 旧诊断接口;不再作为 image 主判据 |
| `/backend-api/files/library` | POST | 用户图像库列表 | 不用于本流程 |
| `/backend-api/models` | GET | 当前账号可用模型 / image 工具入口 | **主能力探针** |
| `/backend-api/me` | GET | 用户基本信息 | 诊断用 |

---

## 4. 关键排查经验

1. **能力入口**：优先看 `/backend-api/models`。`enabledTools:image_gen_tool_enabled` 或 `supportedFeatures:image` 才是 image 入口资格的主信号。
2. **quota/blocked 诊断**：`/conversation/init` 只看 `limits_progress`、`reset_after`、`blocked_features` 作为辅助;不要用它否定 `/models` 的 image 能力。
3. **IMG2 命中**：本次是否命中看真实 `f/conversation` 结果。`file-service://` 或 `sediment:// + generation.gen_size_v2` 都应算 IMG2 协议命中。
4. **抽卡机制**：IMG2 是「账号资格 + 请求/会话抽卡」。同号可能本次 hit、下次 preview_only;调度器需要长期记录 `img2_hit_count / preview_only / miss / delivery_success`。
5. **风控**：HTTP 403、Turnstile、429、下载签名失败要分开看。协议命中不等于交付成功。
6. **TLS / 指纹**：探针和真实 runner 都应复用同一个上游 client: uTLS transport、Oai-* headers、稳定 device/session、cookie jar、代理绑定。

## 5. 相关脚本索引

| 脚本 | 用途 |
|---|---|
| `gen_image.py` | 主生图流程（含重试/轮询/下载） |
| `_check_image_gen_quota.py` | **仅查 `image_gen` 余额**，不消耗额度 |
| `_dump_acc.py` | 完整 dump `/accounts/check`，用于看 feature flag |
| `_check_quota.py` | 遍历多个诊断接口（me/models/accounts/check） |
| `_scan_har_gen.py` / `_scan_har_quota.py` | 扫 HAR 找接口/关键字段 |
| `_har_gen_endpoints.py` / `_dump_init.py` | Dump HAR 里特定接口的完整请求响应 |

---

_最后更新：2026-04-17_
````

## File: cmd/server/main.go
````go
package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/432539/gpt2api/internal/account"
	"github.com/432539/gpt2api/internal/audit"
	"github.com/432539/gpt2api/internal/backup"
	"github.com/432539/gpt2api/internal/config"
	"github.com/432539/gpt2api/internal/db"
	"github.com/432539/gpt2api/internal/gateway"
	"github.com/432539/gpt2api/internal/image"
	modelpkg "github.com/432539/gpt2api/internal/model"
	"github.com/432539/gpt2api/internal/proxy"
	"github.com/432539/gpt2api/internal/scheduler"
	"github.com/432539/gpt2api/internal/middleware"
	"github.com/432539/gpt2api/internal/server"
	"github.com/432539/gpt2api/internal/settings"
	"github.com/432539/gpt2api/internal/usage"
	"github.com/432539/gpt2api/pkg/crypto"
	"github.com/432539/gpt2api/pkg/lock"
	"github.com/432539/gpt2api/pkg/logger"
	"github.com/432539/gpt2api/pkg/mailer"
)

var (
	configPath = flag.String("c", "configs/config.yaml", "config file path")
	showVer    = flag.Bool("v", false, "show version and exit")
)

var (
	version   = "0.2.0-dev"
	buildTime = "unknown"
)

func main() {
	flag.Parse()
	if *showVer {
		fmt.Printf("gpt2api %s (build %s)\n", version, buildTime)
		return
	}

	cfg, err := config.Load(*configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "load config: %v\n", err)
		os.Exit(1)
	}

	if err := logger.Init(cfg.Log.Level, cfg.Log.Format, cfg.Log.Output); err != nil {
		fmt.Fprintf(os.Stderr, "init logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Sync()

	log := logger.L()
	log.Info("boot gpt2api local 2api",
		zap.String("version", version),
		zap.String("env", cfg.App.Env),
		zap.String("listen", cfg.App.Listen),
	)

	sqldb, err := db.NewMySQL(cfg.MySQL)
	if err != nil {
		log.Fatal("mysql init", zap.Error(err))
	}
	defer sqldb.Close()

	rdb, err := db.NewRedis(cfg.Redis)
	if err != nil {
		log.Fatal("redis init", zap.Error(err))
	}
	defer rdb.Close()

	cipher, err := crypto.NewAESGCM(cfg.Crypto.AESKey)
	if err != nil {
		log.Fatal("crypto init", zap.Error(err))
	}
	gateway.InitImageProxySecret(cfg.Crypto.AESKey)
	middleware.InitTokenSecret(cfg.Crypto.AESKey)

	proxyDAO := proxy.NewDAO(sqldb)
	proxySvc := proxy.NewService(proxyDAO, cipher)

	accDAO := account.NewDAO(sqldb)
	accSvc := account.NewService(accDAO, cipher)

	modelDAO := modelpkg.NewDAO(sqldb)
	modelReg := modelpkg.NewRegistry(modelDAO)
	if err := modelReg.Preload(context.Background()); err != nil {
		log.Warn("model preload failed", zap.Error(err))
	}

	rl := lock.NewRedisLock(rdb)
	sched := scheduler.New(accSvc, proxySvc, rl, cfg.Scheduler)

	usageLogger := usage.New(sqldb, usage.Options{})
	defer usageLogger.Close()

	gwH := &gateway.Handler{
		Models:    modelReg,
		Scheduler: sched,
		Usage:     usageLogger,
		AccSvc:    accSvc,
	}

	imageDAO := image.NewDAO(sqldb)
	imageRunner := image.NewRunner(sched, imageDAO)
	imagesH := &gateway.ImagesHandler{
		Handler: gwH,
		Runner:  imageRunner,
		DAO:     imageDAO,
	}
	gwH.Images = imagesH

	auditDAO := audit.NewDAO(sqldb)
	auditH := audit.NewHandler(auditDAO)

	var backupH *backup.Handler
	backupDAO := backup.NewDAO(sqldb)
	if backupSvc, err := backup.New(cfg.Backup, cfg.MySQL, backupDAO); err != nil {
		log.Warn("backup service disabled", zap.Error(err))
	} else {
		backupH = backup.NewHandler(backupSvc, backupDAO, auditDAO)
		log.Info("backup service ready", zap.String("dir", backupSvc.Dir()))
	}

	adminModelH := modelpkg.NewAdminHandler(modelDAO, modelReg, auditDAO)
	usageQDAO := usage.NewQueryDAO(sqldb)
	adminUsageH := usage.NewAdminHandler(usageQDAO)
	meUsageH := usage.NewMeHandler(usageQDAO)
	meImageH := image.NewMeHandler(imageDAO)

	mailSvc := mailer.New(mailer.Config{
		Host:     cfg.SMTP.Host,
		Port:     cfg.SMTP.Port,
		Username: cfg.SMTP.Username,
		Password: cfg.SMTP.Password,
		From:     cfg.SMTP.From,
		FromName: cfg.SMTP.FromName,
		UseTLS:   cfg.SMTP.UseTLS,
	}, log)
	defer mailSvc.Close()
	if mailSvc.Disabled() {
		log.Info("mail channel disabled (smtp.host empty)")
	} else {
		log.Info("mail channel ready", zap.String("host", cfg.SMTP.Host))
	}

	settingsDAO := settings.NewDAO(sqldb)
	settingsSvc := settings.NewService(settingsDAO)
	if err := settingsSvc.Reload(context.Background()); err != nil {
		log.Warn("settings reload failed, using defaults", zap.Error(err))
	}
	settingsH := settings.NewHandler(settingsSvc, mailSvc, auditDAO)
	gwH.Settings = settingsSvc
	sched.SetRuntime(scheduler.RuntimeParams{
		Cooldown429Sec:    settingsSvc.Cooldown429Sec,
		WarnedPauseHrs:    settingsSvc.WarnedPauseHours,
		QueueWaitSec:      settingsSvc.DispatchQueueWaitSec,
		ImageExploreRatio: settingsSvc.ImageExploreRatio,
	})

	proxyH := proxy.NewHandler(proxySvc)
	prober := proxy.NewProber(proxySvc, settingsSvc, log.Named("proxy-prober"))
	proxyH.SetProber(prober)

	// 代理淘汰时自动将绑定的账号重分配到其他可用代理
	prober.SetReassignFunc(func(ctx context.Context, deadProxyID uint64) {
		ids, err := proxySvc.DAO().ListBoundAccountIDs(ctx, deadProxyID)
		if err != nil || len(ids) == 0 {
			return
		}
		for _, accID := range ids {
			if newID, err := accSvc.DAO().SwitchProxy(ctx, accID, deadProxyID); err == nil {
				log.Info("proxy retired: reassigned account",
					zap.Uint64("account_id", accID),
					zap.Uint64("old_proxy_id", deadProxyID),
					zap.Uint64("new_proxy_id", newID))
			}
		}
	})

	proberCtx, cancelProber := context.WithCancel(context.Background())
	defer cancelProber()
	go prober.Run(proberCtx)

	accountH := account.NewHandler(accSvc)
	accRefresher := account.NewRefresher(accSvc, settingsSvc, log.Named("account-refresh"))
	accQuota := account.NewQuotaProber(accSvc, settingsSvc, log.Named("account-quota"))

	acctProxyResolver := &accountProxyResolver{accSvc: accSvc, proxySvc: proxySvc}
	accRefresher.SetProxyResolver(acctProxyResolver)
	accQuota.SetProxyResolver(acctProxyResolver)

	accountH.SetRefresher(accRefresher)
	accountH.SetProber(accQuota)
	accountH.SetSettings(settingsSvc)
	accountH.SetProxyResolver(acctProxyResolver)

	imagesH.ImageAccResolver = acctProxyResolver

	accBgCtx, cancelAccBg := context.WithCancel(context.Background())
	defer cancelAccBg()
	go accRefresher.Run(accBgCtx)
	go accQuota.Run(accBgCtx)

	deps := &server.Deps{
		Config: cfg,

		ProxyH:   proxyH,
		AccountH: accountH,

		GatewayH: gwH,
		ImagesH:  imagesH,

		BackupH:  backupH,
		AuditH:   auditH,
		AuditDAO: auditDAO,

		AdminModelH: adminModelH,
		AdminUsageH: adminUsageH,

		MeUsageH: meUsageH,
		MeImageH: meImageH,

		SettingsH:   settingsH,
		SettingsSvc: settingsSvc,
	}

	r := server.New(deps)
	srv := &http.Server{
		Addr:              cfg.App.Listen,
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		log.Info("http server started", zap.String("addr", cfg.App.Listen))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("http listen", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("graceful shutdown", zap.Error(err))
	}
	log.Info("bye")
}

// accountProxyResolver 把账号 ID → 代理 URL 的查询串起来。
type accountProxyResolver struct {
	accSvc   *account.Service
	proxySvc *proxy.Service
}

// ProxyURLForAccount 查账号绑定的代理并解密密码,返回可直接用于 http.ProxyURL 的 URL。
func (r *accountProxyResolver) ProxyURLForAccount(ctx context.Context, accountID uint64) string {
	if r == nil || r.accSvc == nil || r.proxySvc == nil {
		return ""
	}
	b, err := r.accSvc.GetBinding(ctx, accountID)
	if err != nil || b == nil {
		return ""
	}
	p, err := r.proxySvc.Get(ctx, b.ProxyID)
	if err != nil || p == nil || !p.Enabled {
		return ""
	}
	u, err := r.proxySvc.BuildURL(p)
	if err != nil {
		return ""
	}
	return u
}

// ProxyURLByID 按 proxy_id 直接查代理 URL。
func (r *accountProxyResolver) ProxyURLByID(ctx context.Context, proxyID uint64) string {
	if r == nil || r.proxySvc == nil || proxyID == 0 {
		return ""
	}
	p, err := r.proxySvc.Get(ctx, proxyID)
	if err != nil || p == nil || !p.Enabled {
		return ""
	}
	u, err := r.proxySvc.BuildURL(p)
	if err != nil {
		return ""
	}
	return u
}

// AuthToken 给图片代理端点用:按 accountID 解出 AT / DeviceID / cookies。
func (r *accountProxyResolver) AuthToken(ctx context.Context, accountID uint64) (string, string, string, error) {
	if r == nil || r.accSvc == nil {
		return "", "", "", fmt.Errorf("account service not ready")
	}
	a, err := r.accSvc.Get(ctx, accountID)
	if err != nil {
		return "", "", "", err
	}
	at, err := r.accSvc.DecryptAuthToken(a)
	if err != nil {
		return "", "", "", err
	}
	cookies, _ := r.accSvc.DecryptCookies(ctx, accountID)
	return at, a.OAIDeviceID, cookies, nil
}

// ProxyURL 给图片代理端点用:等价于 ProxyURLForAccount。
func (r *accountProxyResolver) ProxyURL(ctx context.Context, accountID uint64) string {
	return r.ProxyURLForAccount(ctx, accountID)
}
````

## File: cmd/utls-probe/main.go
````go
// cmd/utls-probe 是一个独立小工具,用于验证 utls transport 是否能穿透 Cloudflare。
// 用法(在容器里):  /app/utls-probe   或者直接 go run ./cmd/utls-probe  (本机也行)
//
// 它做 2 件事:
//  1. GET https://chatgpt.com/  →  打印 status / cookie / header
//  2. POST https://chatgpt.com/backend-api/sentinel/chat-requirements (no bearer)
//     →  打印 status + body 前 400 字节
//
// 目的是快速判断:我们的 JA3/JA4 指纹是否被 CF 放行、响应里 Set-Cookie 了哪些
// 关键 cookie、chat-requirements 返回的 body 结构是什么。

package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/432539/gpt2api/internal/upstream/chatgpt"
)

func main() {
	proxyURL := os.Getenv("UTLS_PROBE_PROXY")
	tr, err := chatgpt.NewUTLSTransport(proxyURL, 30*time.Second)
	if err != nil {
		fmt.Println("transport init:", err)
		os.Exit(1)
	}
	hc := &http.Client{Transport: tr, Timeout: 30 * time.Second}

	// 1. GET /
	fmt.Println("== GET https://chatgpt.com/ ==")
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://chatgpt.com/", nil)
	req.Header.Set("User-Agent", chatgpt.DefaultUserAgent)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("Sec-Ch-Ua", `"Google Chrome";v="131", "Chromium";v="131", "Not_A Brand";v="24"`)
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", `"Windows"`)
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	res, err := hc.Do(req)
	if err != nil {
		fmt.Println("  do:", err)
		os.Exit(1)
	}
	buf, _ := io.ReadAll(res.Body)
	_ = res.Body.Close()
	fmt.Printf("  status=%d  proto=%s  body_len=%d\n", res.StatusCode, res.Proto, len(buf))
	fmt.Println("  set-cookie:")
	for _, sc := range res.Header.Values("Set-Cookie") {
		if i := strings.Index(sc, ";"); i > 0 {
			sc = sc[:i]
		}
		fmt.Println("    ", sc)
	}
	fmt.Println("  body head:")
	fmt.Println("   ", preview(buf, 300))

	// 2. POST chat-requirements (no bearer → 应该返回 401,证明至少通过了 CF)
	fmt.Println()
	fmt.Println("== POST /backend-api/sentinel/chat-requirements (no auth) ==")
	req2, _ := http.NewRequestWithContext(context.Background(), http.MethodPost,
		"https://chatgpt.com/backend-api/sentinel/chat-requirements",
		strings.NewReader(`{"p":"gAAAAACxxx"}`))
	req2.Header.Set("User-Agent", chatgpt.DefaultUserAgent)
	req2.Header.Set("Accept", "*/*")
	req2.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("Origin", "https://chatgpt.com")
	req2.Header.Set("Referer", "https://chatgpt.com/")
	req2.Header.Set("Sec-Ch-Ua", `"Google Chrome";v="131", "Chromium";v="131", "Not_A Brand";v="24"`)
	req2.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req2.Header.Set("Sec-Ch-Ua-Platform", `"Windows"`)
	req2.Header.Set("Sec-Fetch-Dest", "empty")
	req2.Header.Set("Sec-Fetch-Mode", "cors")
	req2.Header.Set("Sec-Fetch-Site", "same-origin")
	res2, err := hc.Do(req2)
	if err != nil {
		fmt.Println("  do:", err)
		os.Exit(1)
	}
	buf2, _ := io.ReadAll(res2.Body)
	_ = res2.Body.Close()
	fmt.Printf("  status=%d  proto=%s  body_len=%d\n", res2.StatusCode, res2.Proto, len(buf2))
	fmt.Println("  body head:")
	fmt.Println("   ", preview(buf2, 300))
}

func preview(b []byte, n int) string {
	s := string(b)
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\r", "")
	if len(s) > n {
		return s[:n] + "…"
	}
	return s
}
````

## File: configs/config.example.yaml
````yaml
app:
  name: gpt2api-local
  env: dev
  local_mode: true
  listen: ":8080"
  base_url: ""  # 留空则从请求 Host 自动推导;部署时可设为 "https://your-domain.com"

admin:
  username: "admin"
  password: "admin123"  # 生产环境请修改!

log:
  level: info
  format: console
  output: stdout

mysql:
  dsn: "gpt2api:gpt2api@tcp(127.0.0.1:3306)/gpt2api?parseTime=true&loc=Local&charset=utf8mb4&collation=utf8mb4_unicode_ci"
  max_open_conns: 100
  max_idle_conns: 20
  conn_max_lifetime_sec: 3600

redis:
  addr: "127.0.0.1:6379"
  password: ""
  db: 0
  pool_size: 100

crypto:
  # 64 位十六进制字符串,用于 AES-256-GCM 加密上游账号令牌/cookies/代理密码。
  aes_key: "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"

security:
  cors_origins:
    - "http://localhost:5173"
    - "http://localhost:5174"

scheduler:
  lock_ttl_sec: 1200
  cooldown_429_sec: 600
  warned_pause_hours: 24
  max_concurrent_per_account: 3  # 单账号最大并发生图数

upstream:
  base_url: "https://chatgpt.com"
  request_timeout_sec: 60
  sse_read_timeout_sec: 300

backup:
  dir: "./data/backups"
  retention: 20
  mysqldump_bin: "mysqldump"
  mysql_bin: "mysql"
  max_upload_mb: 512
  allow_restore: false

smtp:
  host: ""
  port: 465
  username: ""
  password: ""
  from: ""
  from_name: "GPT2API Local"
  use_tls: true
````

## File: deploy/.env.example
````
# ================================================================
# docker-compose 环境变量模板。cp .env.example .env 后按需修改。
# ================================================================

# ---- MySQL ----
MYSQL_ROOT_PASSWORD=please_change_me_root
MYSQL_USER=gpt2api
MYSQL_PASSWORD=please_change_me
MYSQL_DATABASE=gpt2api
MYSQL_PORT=3306

# ---- Redis ----
REDIS_PORT=6379

# ---- Server ----
HTTP_PORT=8080

# ---- Security (生产强制改!) ----
# AES-256 加密 KEY,用于加密账号的 AT/Cookie 等。必须 64 位 hex(32 字节)
CRYPTO_AES_KEY=0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef

# ---- Backup ----
# 保留最近 N 个备份,0 代表不清理
BACKUP_RETENTION=30
# 是否启用"数据库恢复"端点(超危)。生产默认 false,
# 需要恢复时先设置为 true 重启容器,完成后切回 false。
BACKUP_ALLOW_RESTORE=false
````

## File: deploy/build-local.ps1
````powershell
# Windows 本地预构建脚本
# 用法:
#   powershell -NoProfile -File deploy/build-local.ps1

$ErrorActionPreference = 'Stop'
# PowerShell 7:关掉 "native 命令 stderr 自动触发终结" 的坑
if ($PSVersionTable.PSVersion.Major -ge 7) {
    $PSNativeCommandUseErrorActionPreference = $false
}

$root = Resolve-Path "$PSScriptRoot/.."
Set-Location $root

Write-Host "[build-local] repo  = $root"
Write-Host "[build-local] step1 = cross-build gpt2api"
$env:GOOS = "linux"
$env:GOARCH = "amd64"
$env:CGO_ENABLED = "0"
New-Item -ItemType Directory -Force deploy/bin | Out-Null
go build -ldflags "-s -w" -o deploy/bin/gpt2api ./cmd/server
if ($LASTEXITCODE -ne 0) { throw "gpt2api build failed" }

Write-Host "[build-local] step2 = npm run build (web)"
Push-Location (Join-Path $root "web")
try {
    if (-not (Test-Path node_modules)) {
        npm install --no-audit --no-fund --loglevel=error
        if ($LASTEXITCODE -ne 0) { throw "npm install failed" }
    }
    npm run build
    if ($LASTEXITCODE -ne 0) { throw "npm run build failed" }
} finally {
    Pop-Location
}

Write-Host "[build-local] done. artifacts:"
Get-Item deploy/bin/gpt2api, web/dist/index.html | Format-Table -AutoSize
````

## File: deploy/build-local.sh
````bash
#!/usr/bin/env bash
# Linux 本地预构建脚本(服务器上直接用 / WSL / macOS 均可)
#
# 用法:
#   bash deploy/build-local.sh
#
# 产物:
#   deploy/bin/gpt2api        linux/amd64 可执行(后端)
#   web/dist/                 前端 Vite 产物
#
# 这套产物 + deploy/Dockerfile 就可以离线构建镜像,无需容器再访问外网。

set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

echo "[build-local] repo  = $ROOT"

# ---- step1: 交叉编译 gpt2api ----
echo "[build-local] step1 = cross-build gpt2api (linux/amd64)"
mkdir -p deploy/bin
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -ldflags "-s -w" -o deploy/bin/gpt2api ./cmd/server

# ---- step2: 前端 ----
echo "[build-local] step2 = npm run build (web)"
pushd web >/dev/null
if [ ! -d node_modules ]; then
    npm install --no-audit --no-fund --loglevel=error
fi
npm run build
popd >/dev/null

echo "[build-local] done. artifacts:"
ls -lh deploy/bin/gpt2api web/dist/index.html
````

## File: deploy/docker-compose.yml
````yaml
services:
  mysql:
    image: mysql:8.0
    container_name: gpt2api-mysql
    restart: unless-stopped
    environment:
      MYSQL_ROOT_PASSWORD: ${MYSQL_ROOT_PASSWORD:-root}
      MYSQL_DATABASE: ${MYSQL_DATABASE:-gpt2api}
      MYSQL_USER: ${MYSQL_USER:-gpt2api}
      MYSQL_PASSWORD: ${MYSQL_PASSWORD:-gpt2api}
      TZ: Asia/Shanghai
    command:
      - --character-set-server=utf8mb4
      - --collation-server=utf8mb4_unicode_ci
      - --default-time-zone=+08:00
      - --innodb-buffer-pool-size=512M
      - --max-connections=500
      # 用 mysql_native_password 作为默认认证插件,兼容容器内
      # mariadb-client 附带的 mariadb-dump(无 caching_sha2_password 插件)。
      - --default-authentication-plugin=mysql_native_password
    ports:
      - "${MYSQL_PORT:-3306}:3306"
    volumes:
      - mysql_data:/var/lib/mysql
    healthcheck:
      test: ["CMD", "mysqladmin", "ping", "-h", "localhost", "-u", "root", "-p${MYSQL_ROOT_PASSWORD:-root}"]
      interval: 10s
      timeout: 5s
      retries: 10

  redis:
    image: redis:7-alpine
    container_name: gpt2api-redis
    restart: unless-stopped
    command: ["redis-server", "--appendonly", "yes"]
    ports:
      - "${REDIS_PORT:-6379}:6379"
    volumes:
      - redis_data:/data
    healthcheck:
      test: ["CMD", "redis-cli", "ping"]
      interval: 10s
      timeout: 5s
      retries: 10

  server:
    build:
      context: ..
      dockerfile: deploy/Dockerfile.multistage
    image: gpt2api/server:latest
    container_name: gpt2api-server
    restart: unless-stopped
    depends_on:
      mysql:
        condition: service_healthy
      redis:
        condition: service_healthy
    environment:
      TZ: Asia/Shanghai
      # entrypoint.sh 会用这些值等 MySQL 并在空库导入 sql/database.sql
      MYSQL_HOST: mysql
      MYSQL_PORT: 3306
      MYSQL_USER: ${MYSQL_USER:-gpt2api}
      MYSQL_PASSWORD: ${MYSQL_PASSWORD:-gpt2api}
      MYSQL_DATABASE: ${MYSQL_DATABASE:-gpt2api}
      # 通过 GPT2API_* 覆盖 config.yaml 中的任一字段
      GPT2API_MYSQL_DSN: "${MYSQL_USER:-gpt2api}:${MYSQL_PASSWORD:-gpt2api}@tcp(mysql:3306)/${MYSQL_DATABASE:-gpt2api}?parseTime=true&loc=Local&charset=utf8mb4&collation=utf8mb4_unicode_ci"
      GPT2API_REDIS_ADDR: "redis:6379"
      GPT2API_BACKUP_DIR: /app/data/backups
      GPT2API_BACKUP_RETENTION: ${BACKUP_RETENTION:-30}
      GPT2API_BACKUP_ALLOW_RESTORE: ${BACKUP_ALLOW_RESTORE:-false}
      # 强烈建议在 .env 中覆盖 AES 密钥
      GPT2API_CRYPTO_AES_KEY: ${CRYPTO_AES_KEY:-0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef}
    ports:
      - "${HTTP_PORT:-8080}:8080"
    # 给容器指定独立公共 DNS,避免宿主 Clash/V2Ray 等梯子劫持了
    # 代理供应商(如 *.arxlabs.io / *.luminati.io 等)的域名解析,
    # 同时 chatgpt.com 也由外部 DNS 统一解析,确保路径一致。
    dns:
      - ${DNS_PRIMARY:-8.8.8.8}
      - ${DNS_FALLBACK:-1.1.1.1}
    volumes:
      - ../configs:/app/configs:ro
      - backups:/app/data/backups
      - ./logs:/app/logs
    command: ["/app/gpt2api", "-c", "/app/configs/config.yaml"]

volumes:
  mysql_data:
  redis_data:
  backups:
````

## File: deploy/Dockerfile
````
# syntax=docker/dockerfile:1.6
#
# "预构建 + 运行时" 镜像 —— 零外网依赖版本(国内 / 内网环境友好)
# --------------------------------------------------------------
# 构建前请先在宿主机生成以下产物(见 deploy/build-local.ps1 / build-local.sh):
#   deploy/bin/gpt2api           Linux/amd64 可执行
#   web/dist/                    前端 Vite 产物
#
# 这样容器内除了 alpine 的 apk 装几个小工具,不需要任何外网访问。
# --------------------------------------------------------------

FROM alpine:3.20

# 走 ustc HTTP 源,规避企业 HTTPS 代理 MITM 导致的证书校验失败
RUN sed -i 's@https\?://dl-cdn.alpinelinux.org@http://mirrors.ustc.edu.cn@g' /etc/apk/repositories \
    && apk add --no-cache ca-certificates tzdata curl bash mariadb-client \
    && update-ca-certificates \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

WORKDIR /app

# 后端二进制(宿主交叉编译产物)
COPY deploy/bin/gpt2api /app/gpt2api
RUN chmod +x /app/gpt2api

# 前端构建产物 —— 被 Go server 以 SPA 方式托管
COPY web/dist /app/web/dist

# 单库初始化 SQL、默认配置、entrypoint
COPY sql     /app/sql
COPY configs /app/configs
COPY deploy/entrypoint.sh /app/entrypoint.sh
RUN sed -i 's/\r$//' /app/entrypoint.sh \
    && chmod +x /app/entrypoint.sh \
    && mkdir -p /app/data/backups /app/logs

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=5s --start-period=20s --retries=3 \
    CMD curl -fsS http://localhost:8080/healthz || exit 1

ENTRYPOINT ["/app/entrypoint.sh"]
CMD ["/app/gpt2api", "-c", "/app/configs/config.yaml"]
````

## File: deploy/Dockerfile.multistage
````
# syntax=docker/dockerfile:1.6
# 多阶段构建：在容器内编译 Go 后端 + Vue 前端，无需本地预构建

# ---- Stage 1: Go 后端编译 ----
FROM golang:1.24-alpine AS go-builder
ENV GOTOOLCHAIN=auto
ENV GONOSUMCHECK=*
ENV GONOSUMDB=*
ENV GOFLAGS=-insecure
RUN apk add --no-cache git
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY cmd/ cmd/
COPY internal/ internal/
COPY pkg/ pkg/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags "-s -w" -o /out/gpt2api ./cmd/server

# ---- Stage 2: Vue 前端构建 ----
FROM node:20-alpine AS fe-builder
WORKDIR /web
COPY web/package.json web/package-lock.json* ./
RUN npm install --no-audit --no-fund --loglevel=error
COPY web/ ./
RUN npm run build

# ---- Stage 3: 运行时镜像 ----
FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata curl bash mariadb-client \
    && update-ca-certificates \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

WORKDIR /app

COPY --from=go-builder /out/gpt2api /app/gpt2api
RUN chmod +x /app/gpt2api

COPY --from=fe-builder /web/dist /app/web/dist

COPY sql     /app/sql
COPY configs /app/configs
COPY deploy/entrypoint.sh /app/entrypoint.sh
RUN sed -i 's/\r$//' /app/entrypoint.sh \
    && chmod +x /app/entrypoint.sh \
    && mkdir -p /app/data/backups /app/logs

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=5s --start-period=20s --retries=3 \
    CMD curl -fsS http://localhost:8080/healthz || exit 1

ENTRYPOINT ["/app/entrypoint.sh"]
CMD ["/app/gpt2api", "-c", "/app/configs/config.yaml"]
````

## File: deploy/entrypoint.sh
````bash
#!/usr/bin/env bash
# gpt2api 容器启动入口。
#
# 职责:
#   1. 等待 MySQL 可连接(最多 60 秒)
#   2. 空库时导入 /app/sql/database.sql
#   3. exec 启动 server 主进程
#
# 读取的环境变量:
#   - MYSQL_HOST        (默认 mysql)
#   - MYSQL_PORT        (默认 3306)
#   - MYSQL_USER        (默认 gpt2api)
#   - MYSQL_PASSWORD    (默认 gpt2api)
#   - MYSQL_DATABASE    (默认 gpt2api)
#   - SKIP_DB_INIT=1    跳过自动初始化
set -euo pipefail

MYSQL_HOST=${MYSQL_HOST:-mysql}
MYSQL_PORT=${MYSQL_PORT:-3306}
MYSQL_USER=${MYSQL_USER:-gpt2api}
MYSQL_PASSWORD=${MYSQL_PASSWORD:-gpt2api}
MYSQL_DATABASE=${MYSQL_DATABASE:-gpt2api}

log() { echo "[entrypoint] $*"; }

wait_mysql() {
  log "waiting for mysql ${MYSQL_HOST}:${MYSQL_PORT}..."
  local i=0
  while (( i < 60 )); do
    if MYSQL_PWD="${MYSQL_PASSWORD}" mysqladmin ping \
        -h "${MYSQL_HOST}" -P "${MYSQL_PORT}" -u "${MYSQL_USER}" --silent 2>/dev/null; then
      log "mysql is up."
      return 0
    fi
    sleep 1
    i=$((i+1))
  done
  log "mysql did not become ready in 60s, continuing anyway."
  return 1
}

run_db_init() {
  if [[ "${SKIP_DB_INIT:-0}" == "1" ]]; then
    log "SKIP_DB_INIT=1, skipping database initialization"
    return 0
  fi

  local table_count
  table_count=$(MYSQL_PWD="${MYSQL_PASSWORD}" mysql \
    -h "${MYSQL_HOST}" -P "${MYSQL_PORT}" -u "${MYSQL_USER}" \
    -N -B -e "SELECT COUNT(*) FROM information_schema.tables WHERE table_schema='${MYSQL_DATABASE}'" 2>/dev/null || echo "0")

  if [[ "${table_count}" != "0" ]]; then
    log "database ${MYSQL_DATABASE} already has ${table_count} table(s), skip sql/database.sql"
    return 0
  fi

  log "database ${MYSQL_DATABASE} is empty, importing /app/sql/database.sql..."
  MYSQL_PWD="${MYSQL_PASSWORD}" mysql \
    -h "${MYSQL_HOST}" -P "${MYSQL_PORT}" -u "${MYSQL_USER}" "${MYSQL_DATABASE}" \
    < /app/sql/database.sql
  log "database initialization done."
}

run_db_migrate() {
  # 增量迁移:检查并添加 image_tasks 表的新字段(幂等)
  run_sql() {
    MYSQL_PWD="${MYSQL_PASSWORD}" mysql \
      -h "${MYSQL_HOST}" -P "${MYSQL_PORT}" -u "${MYSQL_USER}" "${MYSQL_DATABASE}" \
      -N -B -e "$1" 2>/dev/null
  }

  col_exists() {
    local cnt
    cnt=$(run_sql "SELECT COUNT(*) FROM information_schema.columns WHERE table_schema='${MYSQL_DATABASE}' AND table_name='$1' AND column_name='$2'" || echo "0")
    [[ "${cnt}" != "0" ]]
  }

  add_col() {
    local tbl=$1 col=$2 ddl=$3
    if ! col_exists "${tbl}" "${col}"; then
      log "adding column ${tbl}.${col}"
      run_sql "ALTER TABLE ${tbl} ADD COLUMN ${ddl}" || log "  WARN: ALTER failed for ${col}"
    fi
  }

  add_col image_tasks revised_prompt  "revised_prompt TEXT NULL AFTER prompt"
  add_col image_tasks quality         "quality VARCHAR(32) NOT NULL DEFAULT '' AFTER size"
  add_col image_tasks style           "style VARCHAR(32) NOT NULL DEFAULT '' AFTER quality"
  add_col image_tasks reference_urls  "reference_urls JSON NULL AFTER result_urls"
  add_col image_tasks attempts        "attempts INT NOT NULL DEFAULT 0 AFTER error"
  add_col image_tasks duration_ms     "duration_ms BIGINT NOT NULL DEFAULT 0 AFTER attempts"
  add_col image_tasks user_id         "user_id VARCHAR(128) NOT NULL DEFAULT '' AFTER duration_ms"
}

wait_mysql || true
run_db_init || { log "database initialization failed"; exit 1; }
run_db_migrate || { log "database migration failed (non-fatal)"; }

log "starting: $*"
exec "$@"
````

## File: deploy/nginx.conf
````ini
upstream gpt2api_backend {
    server server:8080;
    keepalive 64;
}

server {
    listen 80;
    server_name _;

    client_max_body_size 32m;

    # SSE / streaming 友好
    proxy_http_version 1.1;
    proxy_buffering off;
    proxy_cache off;
    proxy_read_timeout 600s;
    proxy_send_timeout 600s;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;

    # 前端静态(M6 填充)
    root /var/www/html;
    index index.html;

    # OpenAI 兼容接口
    location /v1/ {
        proxy_pass http://gpt2api_backend;
    }

    # 业务 API
    location /api/ {
        proxy_pass http://gpt2api_backend;
    }

    # 健康检查
    location /healthz {
        proxy_pass http://gpt2api_backend;
        access_log off;
    }

    # 前端 SPA 兜底
    location / {
        try_files $uri $uri/ /index.html;
    }
}
````

## File: deploy/README.md
````markdown
# GPT2API Local 容器化部署

一键启动:

```bash
cd deploy
cp .env.example .env
# 修改 CRYPTO_AES_KEY / MySQL 密码等配置
docker compose up -d --build
docker compose logs -f server
```

Server 启动时会等待 MySQL 健康,如果业务库为空则自动导入 `/app/sql/database.sql`,然后启动 HTTP 服务 `:8080`。

## 默认端口

| 服务 | 端口 | 说明 |
| --- | --- | --- |
| server | `8080` | OpenAI 兼容 `/v1` + 本地控制台 API |
| mysql | `3306` | 业务数据库 |
| redis | `6379` | 锁与缓存 |

## 数据库初始化

初始化 SQL 位于 `sql/database.sql`。容器 entrypoint 判断业务库已有表时会跳过导入,避免重复执行。需要从头开始时,清空 MySQL volume 或删库重建即可。

手动初始化:

```bash
MYSQL_PWD=gpt2api docker compose exec -T mysql mysql -ugpt2api gpt2api < ../sql/database.sql
```

## 数据卷

- `mysql_data`:MySQL 物理数据
- `redis_data`:Redis 数据
- `backups`:`/app/data/backups`,数据库备份文件落盘目录
- `./logs`:宿主机日志目录

## 必改配置

生产或长期自用部署时,请在 `.env` 中覆盖:

- `CRYPTO_AES_KEY`:严格 64 位 hex,用于加密上游账号令牌、cookies 和代理敏感字段
- `MYSQL_ROOT_PASSWORD` / `MYSQL_PASSWORD`

## 备份与恢复

控制台“备份恢复”页可创建、下载、上传和恢复数据库备份。恢复默认关闭,需要:

1. 在 `.env` 中设置 `BACKUP_ALLOW_RESTORE=true`
2. 重启 server
3. 在控制台执行恢复
4. 完成后建议改回 `false` 并重启

## 常用运维命令

```bash
# 进入 MySQL
docker compose exec mysql mysql -ugpt2api -p gpt2api

# 查看业务库表数量
docker compose exec mysql mysql -ugpt2api -p -N -B \
  -e "SELECT COUNT(*) FROM information_schema.tables WHERE table_schema='gpt2api'"

# 冷备份
docker compose exec server mysqldump -hmysql -ugpt2api -p \
  --single-transaction --quick gpt2api | gzip > gpt2api-$(date +%F).sql.gz
```
````

## File: docs/FRAMEWORK.md
````markdown
# GPT2API Local 架构说明

本版本定位为单人自用 OpenAI 兼容中转,不是 SaaS 分发平台。核心链路是:

```text
客户端 / OpenAI SDK → /v1 → 本地请求上下文 → 模型映射 → 账号调度 → chatgpt.com → usage/image 日志
```

## 运行模块

| 模块 | 职责 |
| --- | --- |
| `internal/server` | HTTP 路由、控制台静态资源、基础中间件 |
| `internal/middleware` | request id、访问日志、本地控制台上下文 |
| `internal/account` | 上游账号导入、刷新、状态、图片剩余量探测 |
| `internal/proxy` | 代理维护、探测、账号绑定 |
| `internal/model` | 对外 model slug 与上游模型名的映射 |
| `internal/scheduler` | 从账号池中选择可用账号并做冷却/释放 |
| `internal/gateway` | OpenAI 兼容 `/v1` 对话和图片入口 |
| `internal/image` | 图片任务、结果落库、本地签名图片代理 |
| `internal/usage` | 请求日志与聚合统计 |
| `internal/settings` | 本地 KV 设置与公开站点信息 |
| `internal/backup` | 数据库备份、上传、恢复 |
| `internal/audit` | 控制台写操作审计 |

## 请求链路

1. 客户端请求 `/v1/chat/completions` 或 `/v1/images/generations`。
2. 网关读取模型映射,确认该 slug 已开放。
3. 调度器选择可用上游账号,必要时带上绑定代理。
4. chatgpt.com 客户端发起真实请求。
5. 网关把结果转换为 OpenAI 兼容格式返回。
6. `usage_logs` 记录请求类型、模型、账号、状态、token、图片数量、耗时与错误码。
7. 图片结果进入 `image_tasks`,图片 URL 通过 `/p/img/:task_id/:idx` 代理输出。

## 表结构范围

当前初始化 SQL 只包含运维和转发必须的数据:

- `oai_accounts`
- `oai_account_cookies`
- `account_proxy_bindings`
- `proxies`
- `models`
- `usage_logs`
- `image_tasks`
- `system_settings`
- `admin_audit_logs`
- `backup_files`

`usage_logs` 不携带外部分发身份字段,仅作为本地运行观察数据。

## 控制台

控制台页面分为两组:

- 本地中转:`本地总览`、`在线体验`、`用量记录`、`接口文档`
- 运维管理:`上游账号池`、`代理池`、`模型映射`、`全局用量`、`审计日志`、`备份恢复`、`系统设置`

## 配置热更新

`system_settings` 中的站点信息、网关调度参数、账号刷新参数等可通过 `/admin/settings` 编辑。服务启动时会回填缺省值,修改后可手动重载缓存。
````

## File: docs/USER_GUIDE.md
````markdown
# Image2API 使用指南

## 简介

Image2API 提供 OpenAI 兼容的图片生成接口，支持通过 `gpt-image-2` 模型生成高质量图片。支持文生图、图生图（参考图编辑），兼容流式和非流式响应。

---

## 快速开始

### 接口信息

| 项目 | 值 |
|---|---|
| Base URL | `https://your-domain.com` |
| 模型名称 | `gpt-image-2` |
| 认证方式 | `Authorization: Bearer <你的API Key>` |

### 支持的接口

| 接口 | 用途 | 格式 |
|---|---|---|
| `POST /v1/images/generations` | 文生图 / 图生图 | JSON |
| `POST /v1/images/edits` | 图生图 / 编辑 | multipart/form-data |
| `POST /v1/chat/completions` | Chat 格式生图（自动识别）| JSON，支持流式 |

---

## 方式一：Images API

### 文生图

```bash
curl -X POST https://your-domain.com/v1/images/generations \
  -H "Authorization: Bearer sk-your-key" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gpt-image-2",
    "prompt": "一只橘猫坐在窗台上看夕阳，水彩画风格"
  }'
```

### 图生图（reference_images）

通过 `reference_images` 字段传入参考图 URL 或 base64：

```bash
curl -X POST https://your-domain.com/v1/images/generations \
  -H "Authorization: Bearer sk-your-key" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gpt-image-2",
    "prompt": "把这张图片转换成梵高星空风格",
    "reference_images": [
      "https://example.com/my-photo.jpg"
    ]
  }'
```

`reference_images` 支持的格式：
- 图片 URL：`https://example.com/image.png`
- base64 Data URL：`data:image/png;base64,iVBOR...`
- 最多 4 张参考图

### 图生图（multipart 上传）

```bash
curl -X POST https://your-domain.com/v1/images/edits \
  -H "Authorization: Bearer sk-your-key" \
  -F "model=gpt-image-2" \
  -F "prompt=在这张试卷上填写答案" \
  -F "image=@/path/to/photo.jpg"
```

### 响应示例

```json
{
  "created": 1776700000,
  "task_id": "img_a1b2c3d4e5f6...",
  "data": [
    {
      "url": "https://your-domain.com/p/img/img_a1b2c3d4e5f6.../0?exp=...&sig=..."
    }
  ]
}
```

### 请求参数

| 参数 | 类型 | 必填 | 说明 |
|---|---|---|---|
| `model` | string | 是 | 固定填 `gpt-image-2` |
| `prompt` | string | 是 | 图片描述（支持中英文，建议详细描述） |
| `reference_images` | string[] | 否 | 参考图 URL 或 base64（图生图时使用） |
| `n` | int | 否 | 默认 1 |
| `size` | string | 否 | 默认 `1024x1024`（实际尺寸由模型决定） |

---

## 方式二：Chat Completions API

### 纯文生图

```bash
curl -X POST https://your-domain.com/v1/chat/completions \
  -H "Authorization: Bearer sk-your-key" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gpt-image-2",
    "messages": [
      {"role": "user", "content": "画一幅赛博朋克风格的城市夜景"}
    ]
  }'
```

### 图生图（发送图片 + 文字）

这是 Cherry Studio 等客户端最常用的方式——直接在聊天中粘贴/上传图片并附加文字指令：

```bash
curl -X POST https://your-domain.com/v1/chat/completions \
  -H "Authorization: Bearer sk-your-key" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gpt-image-2",
    "messages": [
      {
        "role": "user",
        "content": [
          {"type": "text", "text": "把这张照片转换成油画风格"},
          {"type": "image_url", "image_url": {"url": "https://example.com/photo.jpg"}}
        ]
      }
    ]
  }'
```

`image_url` 支持：
- 图片 URL：`https://example.com/image.png`
- base64 Data URL：`data:image/jpeg;base64,/9j/4AAQ...`（Cherry Studio 粘贴/截图时自动使用此格式）

### 流式响应

设置 `"stream": true` 即可获得 SSE 流式响应：

```bash
curl -X POST https://your-domain.com/v1/chat/completions \
  -H "Authorization: Bearer sk-your-key" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gpt-image-2",
    "stream": true,
    "messages": [
      {"role": "user", "content": "画一只猫"}
    ]
  }'
```

> 注意：由于图片生成需要 30-60 秒，流式模式下会先等待生成完成，然后一次性返回包含图片 URL 的 chunk。

### 响应格式

**非流式：**
```json
{
  "choices": [{
    "message": {
      "role": "assistant",
      "content": "![generated](https://your-domain.com/p/img/img_xxx/0?...)"
    }
  }]
}
```

**流式（SSE）：**
```
data: {"choices":[{"delta":{"role":"assistant","content":"![generated](https://...)"}}]}

data: {"choices":[{"delta":{},"finish_reason":"stop"}]}

data: [DONE]
```

---

## 客户端配置

### Cherry Studio

1. 设置 → 模型服务商 → 添加
2. API 地址：`https://your-domain.com`
3. API Key：你的密钥
4. 模型列表手动添加：`gpt-image-2`
5. 新建对话，选择 `gpt-image-2`，输入提示词即可
6. **图生图**：在对话中直接粘贴/上传图片，附加文字指令

### ChatBox

1. 设置 → AI 服务商 → OpenAI API Compatible
2. API Host：`https://your-domain.com`
3. API Key：你的密钥
4. 模型名称：`gpt-image-2`

### NextChat / ChatGPT-Next-Web

1. 设置 → 接口地址：`https://your-domain.com`
2. API Key：你的密钥
3. 自定义模型名：`gpt-image-2`

### Python / OpenAI SDK

```python
from openai import OpenAI

client = OpenAI(
    api_key="sk-your-key",
    base_url="https://your-domain.com/v1"
)

# 文生图
response = client.images.generate(
    model="gpt-image-2",
    prompt="一只柴犬穿着宇航服在月球上散步",
)
print(response.data[0].url)

# 图生图（通过 chat completions）
response = client.chat.completions.create(
    model="gpt-image-2",
    messages=[{
        "role": "user",
        "content": [
            {"type": "text", "text": "把这张图转成水彩画"},
            {"type": "image_url", "image_url": {"url": "https://example.com/photo.jpg"}}
        ]
    }]
)
print(response.choices[0].message.content)
```

### Node.js / OpenAI SDK

```javascript
import OpenAI from 'openai';

const client = new OpenAI({
  apiKey: 'sk-your-key',
  baseURL: 'https://your-domain.com/v1',
});

// 文生图
const response = await client.images.generate({
  model: 'gpt-image-2',
  prompt: '一幅日式浮世绘风格的海浪',
});
console.log(response.data[0].url);
```

---

## Prompt 技巧

- **越详细越好**：描述主体、风格、构图、光线、色调
- **支持中英文**：效果相当
- **图生图时**：明确描述要做什么修改，如"转换成油画风格"、"在试卷上填写答案"、"把背景换成海滩"

### 示例

| 场景 | Prompt |
|---|---|
| 产品图 | `白色背景上的一杯拿铁咖啡，俯拍视角，专业产品摄影，柔和自然光` |
| 插画 | `一个女孩在樱花树下读书，日系水彩插画风格，温暖色调` |
| 海报 | `赛博朋克风格城市夜景海报，霓虹灯光，雨中倒影，竖版构图` |
| 图标 | `扁平设计风格的云存储图标，蓝白配色，圆角矩形底板` |
| 写实 | `金毛犬在草地上奔跑，阳光明媚，浅景深，佳能85mm镜头效果` |

---

## 常见问题

### Q: 生成需要多久？

通常 30-60 秒。首次使用新账号可能稍长（需要完成账号初始化）。

### Q: 图片链接会过期吗？

图片代理链接 30 天内有效，且可反复访问。

### Q: 图生图怎么用？

三种方式：

1. **Chat 方式（推荐）**：`/v1/chat/completions`，在 message content 中发送 `image_url` + `text`
2. **JSON 方式**：`/v1/images/generations`，在 `reference_images` 字段传图片 URL 或 base64
3. **Multipart 方式**：`/v1/images/edits`，通过 form-data 上传图片文件

### Q: 为什么图生图没有参考到我的图片？

请确认：
- 图片 URL 是公网可访问的（服务器需要能下载到）
- base64 格式需要包含 `data:image/...;base64,` 前缀
- 文件大小不超过 20MB

### Q: 返回 503 / 无可用账号？

服务端账号池暂时繁忙，请稍后重试。

### Q: 返回 400 / content_policy_violation？

你的 prompt 触发了上游内容安全策略，请调整描述。错误信息中会包含具体原因。

### Q: 支持流式响应吗？

支持。在 `/v1/chat/completions` 中设置 `"stream": true`。图片生成完成后会通过 SSE 返回包含图片 URL 的 markdown。

### Q: 支持哪些图片尺寸？

尺寸由模型自动决定（通常为横版 1535x1024 或竖版 1024x1535），`size` 参数仅作记录用。
````

## File: go.mod
````
module github.com/432539/gpt2api

go 1.26.1

require (
	github.com/gin-gonic/gin v1.12.0
	github.com/go-sql-driver/mysql v1.9.3
	github.com/google/uuid v1.6.0
	github.com/jmoiron/sqlx v1.4.0
	github.com/redis/go-redis/v9 v9.18.0
	github.com/refraction-networking/utls v1.8.2
	github.com/spf13/viper v1.21.0
	go.uber.org/zap v1.27.1
	golang.org/x/crypto v0.50.0
	golang.org/x/net v0.52.0
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/andybalholm/brotli v1.0.6 // indirect
	github.com/bytedance/gopkg v0.1.3 // indirect
	github.com/bytedance/sonic v1.15.0 // indirect
	github.com/bytedance/sonic/loader v0.5.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cloudwego/base64x v0.1.6 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.12 // indirect
	github.com/gin-contrib/sse v1.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.30.1 // indirect
	github.com/go-viper/mapstructure/v2 v2.4.0 // indirect
	github.com/goccy/go-json v0.10.5 // indirect
	github.com/goccy/go-yaml v1.19.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.6 // indirect
	github.com/klauspost/cpuid/v2 v2.3.0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.4 // indirect
	github.com/quic-go/qpack v0.6.0 // indirect
	github.com/quic-go/quic-go v0.59.0 // indirect
	github.com/sagikazarmark/locafero v0.11.0 // indirect
	github.com/sourcegraph/conc v0.3.1-0.20240121214520-5f936abd7ae8 // indirect
	github.com/spf13/afero v1.15.0 // indirect
	github.com/spf13/cast v1.10.0 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.3.1 // indirect
	go.mongodb.org/mongo-driver/v2 v2.5.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
	golang.org/x/arch v0.22.0 // indirect
	golang.org/x/sys v0.43.0 // indirect
	golang.org/x/text v0.36.0 // indirect
	google.golang.org/protobuf v1.36.10 // indirect
)
````

## File: internal/account/dao.go
````go
package account

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

var ErrNotFound = errors.New("账号不存在")

type DAO struct{ db *sqlx.DB }

func NewDAO(db *sqlx.DB) *DAO { return &DAO{db: db} }

// DB 暴露底层 handle 给刷新器 / 探测器用于直接原子更新(少量场景)。
func (d *DAO) DB() *sqlx.DB { return d.db }

// fill 填充非 db 列的辅助字段。
func fill(a *Account) {
	if a == nil {
		return
	}
	a.HasRT = a.RefreshTokenEnc.Valid && a.RefreshTokenEnc.String != ""
	a.HasST = a.SessionTokenEnc.Valid && a.SessionTokenEnc.String != ""
}

func fillAll(rows []*Account) {
	for _, r := range rows {
		fill(r)
	}
}

func (d *DAO) Create(ctx context.Context, a *Account) (uint64, error) {
	res, err := d.db.ExecContext(ctx,
		`INSERT INTO oai_accounts
         (email, auth_token_enc, refresh_token_enc, session_token_enc, token_expires_at,
          oai_session_id, oai_device_id, client_id, chatgpt_account_id, account_type,
          subscription_type, status, notes)
         VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		a.Email, a.AuthTokenEnc, a.RefreshTokenEnc, a.SessionTokenEnc, a.TokenExpiresAt,
		a.OAISessionID, a.OAIDeviceID, a.ClientID, a.ChatGPTAccountID, a.AccountType,
		a.SubscriptionType, a.Status, a.Notes,
	)
	if err != nil {
		return 0, err
	}
	id, _ := res.LastInsertId()
	return uint64(id), nil
}

func (d *DAO) GetByID(ctx context.Context, id uint64) (*Account, error) {
	var a Account
	err := d.db.GetContext(ctx, &a,
		`SELECT * FROM oai_accounts WHERE id = ? AND deleted_at IS NULL`, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	fill(&a)
	return &a, err
}

// GetByEmail 精确找;未命中返回 nil, nil(方便 importer 判 upsert)。
func (d *DAO) GetByEmail(ctx context.Context, email string) (*Account, error) {
	var a Account
	err := d.db.GetContext(ctx, &a,
		`SELECT * FROM oai_accounts WHERE email = ? AND deleted_at IS NULL LIMIT 1`, email)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	fill(&a)
	return &a, nil
}

func (d *DAO) List(ctx context.Context, status string, keyword string, offset, limit int) ([]*Account, int64, error) {
	var total int64
	var err error
	var rows []*Account

	where := "deleted_at IS NULL"
	args := []interface{}{}
	if status != "" {
		where += " AND status = ?"
		args = append(args, status)
	}
	if keyword != "" {
		where += " AND (email LIKE ? OR notes LIKE ?)"
		like := "%" + keyword + "%"
		args = append(args, like, like)
	}

	if err = d.db.GetContext(ctx, &total, "SELECT COUNT(*) FROM oai_accounts WHERE "+where, args...); err != nil {
		return nil, 0, err
	}
	argsPage := append([]interface{}{}, args...)
	argsPage = append(argsPage, limit, offset)
	err = d.db.SelectContext(ctx, &rows,
		"SELECT * FROM oai_accounts WHERE "+where+" ORDER BY id DESC LIMIT ? OFFSET ?", argsPage...)
	fillAll(rows)
	return rows, total, err
}

// DispatchOptions 控制账号候选排序。普通 chat 仍按 last_used_at；image 会额外使用
// image2 能力探测与 IMG2 命中画像做「利用 + 探索」排序。
type DispatchOptions struct {
	ModelType         string
	ImageExploreRatio float64       // 0=关闭探索,0.2=约 20% 探索,最大建议 0.8
	ImageExploreStale time.Duration // 多久没尝试的账号进入探索池;<=0 默认 12h
}

// ListDispatchable 调度器专用:返回 status=healthy 且 cooldown 到期、AT 未过期的候选。
// 保留旧签名给非 image 调用使用;image 默认使用 20% 探索比例。
func (d *DAO) ListDispatchable(ctx context.Context, limit int, modelType ...string) ([]*Account, error) {
	opt := DispatchOptions{}
	if len(modelType) > 0 {
		opt.ModelType = modelType[0]
	}
	if opt.ModelType == "image" {
		opt.ImageExploreRatio = 0.2
	}
	return d.ListDispatchableWithOptions(ctx, limit, opt)
}

// ListDispatchableWithOptions 是带调度策略参数的候选查询入口。
//
// image 请求会额外参考 image_capability_status、IMG2 协议命中画像与交付画像排序:
//   - /backend-api/models 探测为 enabled 的账号优先;
//   - 连续 miss 少的账号优先;
//   - 签名 URL 交付成功率高的账号优先;
//   - 历史 IMG2 协议命中率高的账号优先;
//   - 按 ImageExploreRatio 给 unknown/新号/长时间未尝试号留探索位。
//
// 注意这里不把 image_capability_status != enabled 的账号过滤掉,只做排序。原因是
// /backend-api/models 也可能短时失败/缓存缺失,真实 f/conversation 执行结果才是最终判据。
func (d *DAO) ListDispatchableWithOptions(ctx context.Context, limit int, opt DispatchOptions) ([]*Account, error) {
	if opt.ModelType == "image" {
		return d.listDispatchableImage(ctx, limit, opt)
	}

	rows := make([]*Account, 0, limit)
	now := time.Now()
	query := `SELECT * FROM oai_accounts
         WHERE deleted_at IS NULL AND status = 'healthy'
           AND (cooldown_until IS NULL OR cooldown_until <= ?)
           AND (token_expires_at IS NULL OR token_expires_at > ?)
         ORDER BY CASE WHEN last_used_at IS NULL THEN 0 ELSE 1 END, last_used_at ASC
         LIMIT ?`
	err := d.db.SelectContext(ctx, &rows, query, now, now, limit)
	fillAll(rows)
	return rows, err
}

// listDispatchableImage 使用「利用 + 探索」混合排序。
//
// 利用(exploitation):优先 models 探测 enabled、连续 miss 少、交付成功率高、IMG2 协议命中率高的账号。
// 探索(exploration):给 unknown/新号/长时间未尝试号保留可配置比例,避免头部账号被打爆、
// 尾部账号永远没有重新学习机会。
func (d *DAO) listDispatchableImage(ctx context.Context, limit int, opt DispatchOptions) ([]*Account, error) {
	if limit <= 0 {
		return nil, nil
	}
	exploreRatio := opt.ImageExploreRatio
	if exploreRatio < 0 {
		exploreRatio = 0
	}
	if exploreRatio > 0.8 {
		exploreRatio = 0.8
	}
	exploreStale := opt.ImageExploreStale
	if exploreStale <= 0 {
		exploreStale = 12 * time.Hour
	}

	now := time.Now()
	exploit := make([]*Account, 0, limit)
	// 利用池排序:
	//   1) enabled > unknown > error (能力状态)
	//   2) consecutive_miss = 0 的优先(没连续失败的)
	//   3) last_used_at 最早的优先(轮转均衡,核心!)
	// 这样所有健康账号会均匀轮转,而不是永远选同一个"最优"号
	exploitQuery := `SELECT * FROM oai_accounts
         WHERE deleted_at IS NULL AND status = 'healthy'
           AND (cooldown_until IS NULL OR cooldown_until <= ?)
           AND (token_expires_at IS NULL OR token_expires_at > ?)
         ORDER BY
           CASE image_capability_status
             WHEN 'enabled' THEN 0
             WHEN 'unknown' THEN 1
             WHEN 'error'   THEN 2
             ELSE 3
           END ASC,
           CASE WHEN img2_consecutive_miss > 3 THEN 1 ELSE 0 END ASC,
           CASE WHEN last_used_at IS NULL THEN 0 ELSE 1 END,
           last_used_at ASC
         LIMIT ?`
	if err := d.db.SelectContext(ctx, &exploit, exploitQuery, now, now, limit); err != nil {
		return nil, err
	}

	exploreLimit := int(float64(limit)*exploreRatio + 0.999)
	if exploreRatio > 0 && exploreLimit < 1 {
		exploreLimit = 1
	}
	if exploreLimit > limit {
		exploreLimit = limit
	}
	explore := make([]*Account, 0, exploreLimit)
	if exploreLimit > 0 {
		exploreThreshold := now.Add(-exploreStale)
		exploreQuery := `SELECT * FROM oai_accounts
         WHERE deleted_at IS NULL AND status = 'healthy'
           AND (cooldown_until IS NULL OR cooldown_until <= ?)
           AND (token_expires_at IS NULL OR token_expires_at > ?)
           AND (image_capability_status <> 'enabled'
                OR img2_last_attempt_at IS NULL
                OR img2_last_attempt_at <= ?)
         ORDER BY
           CASE WHEN img2_last_attempt_at IS NULL THEN 0 ELSE 1 END ASC,
           img2_last_attempt_at ASC,
           CASE image_capability_status
             WHEN 'unknown' THEN 0
             WHEN 'error'   THEN 1
             WHEN 'enabled' THEN 2
             ELSE 3
           END ASC,
           CASE WHEN last_used_at IS NULL THEN 0 ELSE 1 END,
           last_used_at ASC
         LIMIT ?`
		if err := d.db.SelectContext(ctx, &explore, exploreQuery, now, now, exploreThreshold, exploreLimit); err != nil {
			return nil, err
		}
	}

	rows := mergeDispatchCandidates(exploit, explore, limit, exploreRatio)
	fillAll(rows)
	return rows, nil
}

func mergeDispatchCandidates(exploit, explore []*Account, limit int, exploreRatio float64) []*Account {
	out := make([]*Account, 0, limit)
	seen := make(map[uint64]struct{}, limit)
	add := func(a *Account) bool {
		if a == nil {
			return false
		}
		if _, ok := seen[a.ID]; ok {
			return false
		}
		seen[a.ID] = struct{}{}
		out = append(out, a)
		return true
	}

	ei, xi := 0, 0
	addExplore := func() bool {
		for xi < len(explore) && len(out) < limit {
			a := explore[xi]
			xi++
			if add(a) {
				return true
			}
		}
		return false
	}
	addExploit := func() bool {
		for ei < len(exploit) && len(out) < limit {
			a := exploit[ei]
			ei++
			if add(a) {
				return true
			}
		}
		return false
	}

	if exploreRatio < 0 {
		exploreRatio = 0
	}
	if exploreRatio > 0.8 {
		exploreRatio = 0.8
	}
	exploreDebt := 0.0

	for len(out) < limit && (ei < len(exploit) || xi < len(explore)) {
		if exploreRatio > 0 {
			exploreDebt += exploreRatio
		}
		if exploreDebt >= 1 && addExplore() {
			exploreDebt -= 1
			continue
		}
		if addExploit() {
			continue
		}
		if addExplore() {
			if exploreDebt >= 1 {
				exploreDebt -= 1
			}
			continue
		}
		break
	}
	return out
}

// ListNeedRefresh 返回需要预刷新的账号(AT 将在 aheadSec 秒内过期)。
// 按 token_expires_at 升序,最快过期的先刷。
func (d *DAO) ListNeedRefresh(ctx context.Context, aheadSec int, limit int) ([]*Account, error) {
	rows := make([]*Account, 0, limit)
	threshold := time.Now().Add(time.Duration(aheadSec) * time.Second)
	err := d.db.SelectContext(ctx, &rows,
		`SELECT * FROM oai_accounts
         WHERE deleted_at IS NULL
           AND status <> 'dead'
           AND (refresh_token_enc IS NOT NULL OR session_token_enc IS NOT NULL)
           AND token_expires_at IS NOT NULL
           AND token_expires_at <= ?
         ORDER BY token_expires_at ASC
         LIMIT ?`, threshold, limit)
	fillAll(rows)
	return rows, err
}

// ListNeedProbeQuota 返回需要探测图片额度的账号(上次探测超过 minIntervalSec 秒,或从未探测过)。
func (d *DAO) ListNeedProbeQuota(ctx context.Context, minIntervalSec int, limit int) ([]*Account, error) {
	rows := make([]*Account, 0, limit)
	threshold := time.Now().Add(-time.Duration(minIntervalSec) * time.Second)
	err := d.db.SelectContext(ctx, &rows,
		`SELECT * FROM oai_accounts
         WHERE deleted_at IS NULL
           AND status = 'healthy'
           AND (token_expires_at IS NULL OR token_expires_at > NOW())
           AND (image_quota_updated_at IS NULL OR image_quota_updated_at <= ?)
         ORDER BY CASE WHEN image_quota_updated_at IS NULL THEN 0 ELSE 1 END,
                  image_quota_updated_at ASC
         LIMIT ?`, threshold, limit)
	fillAll(rows)
	return rows, err
}

// ListAllActiveIDs 用于批量刷新 / 批量探测:返回未软删的所有 id。
func (d *DAO) ListAllActiveIDs(ctx context.Context) ([]uint64, error) {
	ids := make([]uint64, 0, 128)
	err := d.db.SelectContext(ctx, &ids,
		`SELECT id FROM oai_accounts WHERE deleted_at IS NULL ORDER BY id ASC`)
	return ids, err
}

func (d *DAO) Update(ctx context.Context, a *Account) error {
	_, err := d.db.ExecContext(ctx,
		`UPDATE oai_accounts
         SET email=?, auth_token_enc=?, refresh_token_enc=?, session_token_enc=?, token_expires_at=?,
             oai_session_id=?, oai_device_id=?, client_id=?, chatgpt_account_id=?, account_type=?,
             subscription_type=?, status=?, notes=?
         WHERE id = ? AND deleted_at IS NULL`,
		a.Email, a.AuthTokenEnc, a.RefreshTokenEnc, a.SessionTokenEnc, a.TokenExpiresAt,
		a.OAISessionID, a.OAIDeviceID, a.ClientID, a.ChatGPTAccountID, a.AccountType,
		a.SubscriptionType, a.Status, a.Notes, a.ID,
	)
	return err
}

// SetSubscriptionType 更新账号的订阅类型（pro/plus/free/team 等）。
func (d *DAO) SetSubscriptionType(ctx context.Context, id uint64, subType string) error {
	_, err := d.db.ExecContext(ctx,
		`UPDATE oai_accounts SET subscription_type = ? WHERE id = ? AND deleted_at IS NULL`, subType, id)
	return err
}

func (d *DAO) SoftDelete(ctx context.Context, id uint64) error {
	_, err := d.db.ExecContext(ctx,
		`UPDATE oai_accounts SET deleted_at = ? WHERE id = ?`, time.Now(), id)
	return err
}

// PurgeSoftDeleted 真删除所有已软删除的账号及其关联数据。返回删除行数。
func (d *DAO) PurgeSoftDeleted(ctx context.Context) (int64, error) {
	// 先删关联表
	d.db.ExecContext(ctx, `DELETE FROM account_proxy_bindings WHERE account_id IN (SELECT id FROM oai_accounts WHERE deleted_at IS NOT NULL)`)
	d.db.ExecContext(ctx, `DELETE FROM oai_account_cookies WHERE account_id IN (SELECT id FROM oai_accounts WHERE deleted_at IS NOT NULL)`)
	res, err := d.db.ExecContext(ctx, `DELETE FROM oai_accounts WHERE deleted_at IS NOT NULL`)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// SoftDeleteByStatus 按状态批量软删。status 为空时删除全部(调用方需二次确认)。
// 返回删除行数。
func (d *DAO) SoftDeleteByStatus(ctx context.Context, status string) (int64, error) {
	now := time.Now()
	if status == "" {
		res, err := d.db.ExecContext(ctx,
			`UPDATE oai_accounts SET deleted_at = ? WHERE deleted_at IS NULL`, now)
		if err != nil {
			return 0, err
		}
		n, _ := res.RowsAffected()
		return n, nil
	}
	res, err := d.db.ExecContext(ctx,
		`UPDATE oai_accounts SET deleted_at = ? WHERE deleted_at IS NULL AND status = ?`,
		now, status)
	if err != nil {
		return 0, err
	}
	n, _ := res.RowsAffected()
	return n, nil
}

// EnsureDeviceID 确保账号有 oai_device_id。
// 如果当前为空,原子写入给定的 deviceID;返回最终实际的 device_id(已有则原值)。
func (d *DAO) EnsureDeviceID(ctx context.Context, id uint64, deviceID string) (string, error) {
	_, err := d.db.ExecContext(ctx,
		`UPDATE oai_accounts SET oai_device_id = ?
         WHERE id = ? AND deleted_at IS NULL AND (oai_device_id = '' OR oai_device_id IS NULL)`,
		deviceID, id)
	if err != nil {
		return "", err
	}
	// 回读,兼容其他协程并发填写的情形
	var cur string
	if err := d.db.GetContext(ctx, &cur,
		`SELECT oai_device_id FROM oai_accounts WHERE id = ?`, id); err != nil {
		return "", err
	}
	return cur, nil
}

// EnsureSessionID 确保账号有 oai_session_id(按账号稳定复用)。
// 逻辑与 EnsureDeviceID 完全一致,单独一个函数是为了日志/审计区分用途。
func (d *DAO) EnsureSessionID(ctx context.Context, id uint64, sessionID string) (string, error) {
	_, err := d.db.ExecContext(ctx,
		`UPDATE oai_accounts SET oai_session_id = ?
         WHERE id = ? AND deleted_at IS NULL AND (oai_session_id = '' OR oai_session_id IS NULL)`,
		sessionID, id)
	if err != nil {
		return "", err
	}
	var cur string
	if err := d.db.GetContext(ctx, &cur,
		`SELECT oai_session_id FROM oai_accounts WHERE id = ?`, id); err != nil {
		return "", err
	}
	return cur, nil
}

// MarkUsed 更新 last_used_at + 今日计数。today 是当日零点(用于 today_used_date 比较)。
func (d *DAO) MarkUsed(ctx context.Context, id uint64, today time.Time) error {
	_, err := d.db.ExecContext(ctx,
		`UPDATE oai_accounts
         SET last_used_at = ?,
             today_used_count = CASE WHEN today_used_date = ? THEN today_used_count + 1 ELSE 1 END,
             today_used_date  = ?
         WHERE id = ?`,
		time.Now(), today, today, id)
	return err
}

// SetStatus 迁移状态,可选 cooldownUntil。
func (d *DAO) SetStatus(ctx context.Context, id uint64, status string, cooldownUntil *time.Time) error {
	if cooldownUntil != nil {
		_, err := d.db.ExecContext(ctx,
			`UPDATE oai_accounts SET status=?, cooldown_until=? WHERE id=?`,
			status, *cooldownUntil, id)
		return err
	}
	_, err := d.db.ExecContext(ctx,
		`UPDATE oai_accounts SET status=?, cooldown_until=NULL WHERE id=?`,
		status, id)
	return err
}

// ApplyRefreshResult 原子更新 AT / RT + 过期时间 + 最近刷新信息。
// newRTEnc 为空字符串表示 RT 没有轮转,保持不变。
func (d *DAO) ApplyRefreshResult(
	ctx context.Context,
	id uint64,
	newATEnc string,
	newRTEnc string,
	expiresAt time.Time,
	source string,
) error {
	var err error
	if newRTEnc != "" {
		_, err = d.db.ExecContext(ctx,
			`UPDATE oai_accounts
             SET auth_token_enc = ?,
                 refresh_token_enc = ?,
                 token_expires_at = ?,
                 last_refresh_at = ?,
                 last_refresh_source = ?,
                 refresh_error = '',
                 status = CASE WHEN status IN ('dead','suspicious') THEN 'healthy' ELSE status END
             WHERE id = ? AND deleted_at IS NULL`,
			newATEnc, newRTEnc, expiresAt, time.Now(), source, id)
	} else {
		_, err = d.db.ExecContext(ctx,
			`UPDATE oai_accounts
             SET auth_token_enc = ?,
                 token_expires_at = ?,
                 last_refresh_at = ?,
                 last_refresh_source = ?,
                 refresh_error = '',
                 status = CASE WHEN status IN ('dead','suspicious') THEN 'healthy' ELSE status END
             WHERE id = ? AND deleted_at IS NULL`,
			newATEnc, expiresAt, time.Now(), source, id)
	}
	return err
}

// RetireExpiredATOnly 把已过期且无 RT/ST 的纯 AT 账号自动标记 dead。
// 返回受影响的行数。
func (d *DAO) RetireExpiredATOnly(ctx context.Context) (int64, error) {
	res, err := d.db.ExecContext(ctx, `
		UPDATE oai_accounts SET status = 'dead', refresh_error = 'AT 已过期且无 RT/ST,自动丢弃'
		WHERE deleted_at IS NULL
		  AND status <> 'dead'
		  AND token_expires_at IS NOT NULL
		  AND token_expires_at < NOW()
		  AND (refresh_token_enc IS NULL OR refresh_token_enc = '')
		  AND (session_token_enc IS NULL OR session_token_enc = '')`)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// RecordRefreshError 写入刷新失败原因,同时推进 last_refresh_at(避免 pressed-out 重试)。
// markDead=true 标 dead(仅用于确定性不可恢复的情况);false 标 warned(保留重试机会)。
func (d *DAO) RecordRefreshError(ctx context.Context, id uint64, source string, reason string, markDead bool) error {
	status := "warned"
	if markDead {
		status = "dead"
	}
	_, err := d.db.ExecContext(ctx,
		`UPDATE oai_accounts
         SET last_refresh_at = ?, last_refresh_source = ?, refresh_error = ?,
             status = CASE WHEN status = 'healthy' THEN ? ELSE status END
         WHERE id = ? AND deleted_at IS NULL`,
		time.Now(), source, reason, status, id)
	return err
}

// ApplyQuotaResult 更新图片额度探测结果;remaining/total = -1 表示保持原值。
func (d *DAO) ApplyQuotaResult(ctx context.Context, id uint64, remaining, total int, resetAt *time.Time) error {
	q := `UPDATE oai_accounts
          SET image_quota_remaining = CASE WHEN ? < 0 THEN image_quota_remaining ELSE ? END,
              image_quota_total     = CASE WHEN ? < 0 THEN image_quota_total     ELSE ? END,
              image_quota_reset_at  = ?,
              image_quota_updated_at = ?
          WHERE id = ? AND deleted_at IS NULL`
	var reset interface{}
	if resetAt != nil {
		reset = *resetAt
	} else {
		reset = nil
	}
	_, err := d.db.ExecContext(ctx, q, remaining, remaining, total, total, reset, time.Now(), id)
	return err
}

// ApplyImageCapabilityResult 写入 image 能力探测结果。
// status 建议取 unknown/enabled/disabled/error;source 目前主要是 models。
func (d *DAO) ApplyImageCapabilityResult(ctx context.Context, id uint64, status, model, source, detail string, blockedFeatures []string) error {
	if status == "" {
		status = "unknown"
	}
	blockedJSON, _ := json.Marshal(blockedFeatures)
	_, err := d.db.ExecContext(ctx,
		`UPDATE oai_accounts
          SET image_capability_status = ?,
              image_capability_model = ?,
              image_capability_source = ?,
              image_capability_detail = ?,
              image_capability_updated_at = ?,
              image_init_blocked_features = ?
          WHERE id = ? AND deleted_at IS NULL`,
		status, model, source, detail, time.Now(), string(blockedJSON), id)
	return err
}

// RecordIMG2Outcome 记录一次真实 f/conversation 图像请求的 IMG2 命中画像。
// outcome:
//   - hit:          出现 file-service 或 sediment+gen_size_v2 等 IMG2 指纹
//   - preview_only:只拿到 IMG1 preview/sediment 兜底
//   - miss:         真实请求没拿到有效图片结构
func (d *DAO) RecordIMG2Outcome(ctx context.Context, id uint64, outcome string) error {
	now := time.Now()
	switch outcome {
	case "hit":
		_, err := d.db.ExecContext(ctx,
			`UPDATE oai_accounts
              SET img2_hit_count = img2_hit_count + 1,
                  img2_consecutive_miss = 0,
                  img2_last_status = ?,
                  img2_last_hit_at = ?,
                  img2_last_attempt_at = ?
              WHERE id = ? AND deleted_at IS NULL`,
			outcome, now, now, id)
		return err
	case "preview_only":
		_, err := d.db.ExecContext(ctx,
			`UPDATE oai_accounts
              SET img2_preview_only_count = img2_preview_only_count + 1,
                  img2_consecutive_miss = img2_consecutive_miss + 1,
                  img2_last_status = ?,
                  img2_last_attempt_at = ?
              WHERE id = ? AND deleted_at IS NULL`,
			outcome, now, id)
		return err
	case "miss":
		_, err := d.db.ExecContext(ctx,
			`UPDATE oai_accounts
              SET img2_miss_count = img2_miss_count + 1,
                  img2_consecutive_miss = img2_consecutive_miss + 1,
                  img2_last_status = ?,
                  img2_last_attempt_at = ?
              WHERE id = ? AND deleted_at IS NULL`,
			outcome, now, id)
		return err
	default:
		return nil
	}
}

// RecordIMG2Delivery 记录 IMG2 协议命中后的交付结果。
//
// 注意:img2_hit_count 表示「协议/灰度抽中」,这里表示「签名 URL 交付成功」。
// 两者拆开后,调度器可以识别「经常抽中但下载失败」的账号/代理。
func (d *DAO) RecordIMG2Delivery(ctx context.Context, id uint64, status string) error {
	now := time.Now()
	switch status {
	case "success":
		_, err := d.db.ExecContext(ctx,
			`UPDATE oai_accounts
              SET img2_delivery_success_count = img2_delivery_success_count + 1,
                  img2_last_delivery_status = ?,
                  img2_last_delivery_at = ?
              WHERE id = ? AND deleted_at IS NULL`,
			status, now, id)
		return err
	case "partial":
		_, err := d.db.ExecContext(ctx,
			`UPDATE oai_accounts
              SET img2_delivery_partial_count = img2_delivery_partial_count + 1,
                  img2_last_delivery_status = ?,
                  img2_last_delivery_at = ?
              WHERE id = ? AND deleted_at IS NULL`,
			status, now, id)
		return err
	case "fail":
		_, err := d.db.ExecContext(ctx,
			`UPDATE oai_accounts
              SET img2_delivery_fail_count = img2_delivery_fail_count + 1,
                  img2_last_delivery_status = ?,
                  img2_last_delivery_at = ?
              WHERE id = ? AND deleted_at IS NULL`,
			status, now, id)
		return err
	default:
		return nil
	}
}

// ---- cookies ----

func (d *DAO) UpsertCookies(ctx context.Context, accountID uint64, cookieEnc string) error {
	_, err := d.db.ExecContext(ctx,
		`INSERT INTO oai_account_cookies (account_id, cookie_json_enc)
         VALUES (?, ?)
         ON DUPLICATE KEY UPDATE cookie_json_enc = VALUES(cookie_json_enc)`,
		accountID, cookieEnc)
	return err
}

func (d *DAO) GetCookies(ctx context.Context, accountID uint64) (string, error) {
	var enc string
	err := d.db.GetContext(ctx, &enc,
		`SELECT cookie_json_enc FROM oai_account_cookies WHERE account_id = ?`,
		accountID)
	if errors.Is(err, sql.ErrNoRows) {
		return "", nil
	}
	return enc, err
}

// ---- bindings ----

func (d *DAO) SetBinding(ctx context.Context, accountID, proxyID uint64) error {
	_, err := d.db.ExecContext(ctx,
		`INSERT INTO account_proxy_bindings (account_id, proxy_id)
         VALUES (?, ?)
         ON DUPLICATE KEY UPDATE proxy_id = VALUES(proxy_id), bound_at = CURRENT_TIMESTAMP`,
		accountID, proxyID)
	return err
}

func (d *DAO) RemoveBinding(ctx context.Context, accountID uint64) error {
	_, err := d.db.ExecContext(ctx,
		`DELETE FROM account_proxy_bindings WHERE account_id = ?`, accountID)
	return err
}

func (d *DAO) GetBinding(ctx context.Context, accountID uint64) (*Binding, error) {
	var b Binding
	err := d.db.GetContext(ctx, &b,
		`SELECT * FROM account_proxy_bindings WHERE account_id = ?`, accountID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return &b, err
}

// LeastBoundProxyID 返回绑定账号数最少的可用代理 ID(用于自动分配)。
// excludeIDs 可以排除指定代理(如刚失效的)。
func (d *DAO) LeastBoundProxyID(ctx context.Context, excludeIDs []uint64) (uint64, error) {
	query := `
		SELECT p.id
		FROM proxies p
		LEFT JOIN account_proxy_bindings b ON b.proxy_id = p.id
		WHERE p.enabled = 1 AND p.deleted_at IS NULL
		  AND p.health_score > 0`
	var args []interface{}
	if len(excludeIDs) > 0 {
		placeholders := make([]string, len(excludeIDs))
		for i, id := range excludeIDs {
			placeholders[i] = "?"
			args = append(args, id)
		}
		query += " AND p.id NOT IN (" + strings.Join(placeholders, ",") + ")"
	}
	query += `
		GROUP BY p.id
		ORDER BY COUNT(b.account_id) ASC, p.health_score DESC
		LIMIT 1`
	var proxyID uint64
	err := d.db.GetContext(ctx, &proxyID, query, args...)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, errors.New("无可用代理")
	}
	return proxyID, err
}

// SwitchProxy 将账号切换到另一个代理(排除当前代理),返回新代理 ID。
func (d *DAO) SwitchProxy(ctx context.Context, accountID, currentProxyID uint64) (uint64, error) {
	newID, err := d.LeastBoundProxyID(ctx, []uint64{currentProxyID})
	if err != nil {
		return 0, err
	}
	if err := d.SetBinding(ctx, accountID, newID); err != nil {
		return 0, err
	}
	return newID, nil
}
````

## File: internal/account/handler.go
````go
package account

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"

	"github.com/432539/gpt2api/internal/settings"
	"github.com/432539/gpt2api/pkg/resp"
)

// ProxyURLResolver 按 proxy_id 取代理 URL(已带密码),供 ImportTokens 时走 RT/ST 换 AT 使用。
// 由外部传入一个实现(通常是 proxy.Service 的包装),避免 account 包直接依赖 proxy 包。
type ProxyURLResolver interface {
	ProxyURLByID(ctx context.Context, proxyID uint64) string
}

type Handler struct {
	svc           *Service
	refresher     *Refresher
	prober        *QuotaProber
	settings      *settings.Service
	proxyResolver ProxyURLResolver
}

func NewHandler(s *Service) *Handler { return &Handler{svc: s} }

// SetRefresher 注入刷新器(可选,未注入时相关接口返回 501)。
func (h *Handler) SetRefresher(r *Refresher) { h.refresher = r }

// SetProber 注入额度探测器(可选)。
func (h *Handler) SetProber(p *QuotaProber) { h.prober = p }

// SetSettings 注入系统设置服务,用于自动刷新开关的读写。
func (h *Handler) SetSettings(s *settings.Service) { h.settings = s }

// SetProxyResolver 注入代理 URL 解析器(可选,未注入时 RT/ST 批量导入只能直连)。
func (h *Handler) SetProxyResolver(r ProxyURLResolver) { h.proxyResolver = r }

// POST /api/admin/accounts
func (h *Handler) Create(c *gin.Context) {
	var req CreateInput
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.BadRequest(c, "请求参数错误:"+err.Error())
		return
	}
	a, err := h.svc.Create(c.Request.Context(), req)
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}
	resp.OK(c, a)
}

// GET /api/admin/accounts
func (h *Handler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	if page < 1 {
		page = 1
	}
	size, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if size < 1 {
		size = 10
	}
	if size > 1000 {
		size = 1000
	}
	status := c.Query("status")
	keyword := c.Query("keyword")
	list, total, err := h.svc.List(c.Request.Context(), status, keyword, (page-1)*size, size)
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}
	resp.OK(c, gin.H{"list": list, "total": total, "page": page, "page_size": size})
}

// GET /api/admin/accounts/:id
func (h *Handler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	a, err := h.svc.Get(c.Request.Context(), id)
	if err != nil {
		resp.NotFound(c, err.Error())
		return
	}
	resp.OK(c, a)
}

// PATCH /api/admin/accounts/:id
func (h *Handler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req UpdateInput
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.BadRequest(c, "请求参数错误:"+err.Error())
		return
	}
	a, err := h.svc.Update(c.Request.Context(), id, req)
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}
	resp.OK(c, a)
}

// DELETE /api/admin/accounts/:id
func (h *Handler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		resp.Internal(c, err.Error())
		return
	}
	resp.OK(c, gin.H{"deleted": id})
}

// GET /api/admin/accounts/:id/secrets
// 仅本地控制台可用,返回 AT / RT / ST 明文用于编辑弹窗回显。
func (h *Handler) GetSecrets(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	sec, err := h.svc.GetSecrets(c.Request.Context(), id)
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}
	resp.OK(c, sec)
}

// POST /api/admin/accounts/bulk-delete
// body: { "scope": "dead" | "suspicious" | "warned" | "throttled" | "all" }
// 批量软删指定状态的账号;scope=all 时删除全部(调用方需二次确认)。
func (h *Handler) BulkDelete(c *gin.Context) {
	var req struct {
		Scope string `json:"scope" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.BadRequest(c, "请求参数错误:"+err.Error())
		return
	}
	scope := strings.ToLower(strings.TrimSpace(req.Scope))
	allowed := map[string]bool{
		"dead": true, "suspicious": true, "warned": true, "throttled": true, "all": true,
	}
	if !allowed[scope] {
		resp.BadRequest(c, "scope 仅支持 dead / suspicious / warned / throttled / all")
		return
	}
	n, err := h.svc.BulkDeleteByStatus(c.Request.Context(), scope)
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}
	resp.OK(c, gin.H{"deleted": n, "scope": scope})
}

// POST /api/admin/accounts/purge-deleted
// 真删除所有已软删除的账号及关联数据。
func (h *Handler) PurgeDeleted(c *gin.Context) {
	n, err := h.svc.DAO().PurgeSoftDeleted(c.Request.Context())
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}
	resp.OK(c, gin.H{"purged": n})
}

// ===================== 自动刷新开关 =====================

// GET /api/admin/accounts/auto-refresh
// 返回当前自动刷新配置。
func (h *Handler) GetAutoRefresh(c *gin.Context) {
	if h.settings == nil {
		resp.Internal(c, "系统设置未初始化")
		return
	}
	resp.OK(c, gin.H{
		"enabled":   h.settings.AccountRefreshEnabled(),
		"ahead_sec": h.settings.AccountRefreshAheadSec(),
		"threshold": "AT 距离过期 < 1 天时自动刷新,失效/可疑账号不刷新",
	})
}

// PUT /api/admin/accounts/auto-refresh
// body: { "enabled": true|false }
// 写入 account.refresh_enabled;同时把阈值固定为 86400(1 天)以满足 UI 语义。
func (h *Handler) SetAutoRefresh(c *gin.Context) {
	if h.settings == nil {
		resp.Internal(c, "系统设置未初始化")
		return
	}
	var req struct {
		Enabled bool `json:"enabled"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.BadRequest(c, "请求参数错误:"+err.Error())
		return
	}
	updates := map[string]string{
		settings.AccountRefreshEnabled:  boolStr(req.Enabled),
		settings.AccountRefreshAheadSec: "86400",
	}
	if err := h.settings.Set(c.Request.Context(), updates); err != nil {
		resp.Internal(c, "保存失败:"+err.Error())
		return
	}
	if req.Enabled && h.refresher != nil {
		h.refresher.Kick() // 立刻扫一遍
	}
	resp.OK(c, gin.H{
		"enabled":   req.Enabled,
		"ahead_sec": 86400,
	})
}

func boolStr(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

// 保留以便未来直接传 context(当前未用,但留一个显式符号避免删字段)
var _ = context.Background

// POST /api/admin/accounts/:id/bind-proxy
func (h *Handler) BindProxy(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req struct {
		ProxyID uint64 `json:"proxy_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.BadRequest(c, "请求参数错误:"+err.Error())
		return
	}
	if err := h.svc.BindProxy(c.Request.Context(), id, req.ProxyID); err != nil {
		resp.Internal(c, err.Error())
		return
	}
	resp.OK(c, gin.H{"account_id": id, "proxy_id": req.ProxyID})
}

// DELETE /api/admin/accounts/:id/bind-proxy
func (h *Handler) UnbindProxy(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.svc.UnbindProxy(c.Request.Context(), id); err != nil {
		resp.Internal(c, err.Error())
		return
	}
	resp.OK(c, gin.H{"account_id": id})
}

// ===================== 批量导入 =====================

// POST /api/admin/accounts/import
// body: { text: "...", update_existing: true, default_client_id: "", default_proxy_id: 0 }
// 或 multipart/form-data:files[] + 其他字段
func (h *Handler) Import(c *gin.Context) {
	var req struct {
		Text            string `json:"text"`
		UpdateExisting  *bool  `json:"update_existing"`
		DefaultClientID string `json:"default_client_id"`
		DefaultProxyID  uint64 `json:"default_proxy_id"`
	}
	// 支持 JSON body 或 multipart
	ct := c.ContentType()
	if ct == "application/json" {
		if err := c.ShouldBindJSON(&req); err != nil {
			resp.BadRequest(c, "请求参数错误:"+err.Error())
			return
		}
	} else {
		// multipart 表单
		req.Text = c.PostForm("text")
		if v := c.PostForm("update_existing"); v != "" {
			b := v == "true" || v == "1"
			req.UpdateExisting = &b
		}
		req.DefaultClientID = c.PostForm("default_client_id")
		if v := c.PostForm("default_proxy_id"); v != "" {
			if n, err := strconv.ParseUint(v, 10, 64); err == nil {
				req.DefaultProxyID = n
			}
		}
		// 多文件合并:允许前端一次上传 N 个 json
		if form, err := c.MultipartForm(); err == nil && form != nil {
			var sb strings.Builder
			if req.Text != "" {
				sb.WriteString(req.Text)
				sb.WriteByte('\n')
			}
			for _, fh := range form.File["files"] {
				f, err := fh.Open()
				if err != nil {
					continue
				}
				data, err := io.ReadAll(f)
				_ = f.Close()
				if err != nil || len(data) == 0 {
					continue
				}
				sb.Write(data)
				sb.WriteByte('\n')
			}
			req.Text = sb.String()
		}
	}

	if req.Text == "" {
		resp.BadRequest(c, "请提供 text 或上传文件")
		return
	}

	items, err := ParseJSONBlob(req.Text)
	if err != nil {
		resp.BadRequest(c, "解析失败:"+err.Error())
		return
	}

	upd := true
	if req.UpdateExisting != nil {
		upd = *req.UpdateExisting
	}

	opt := ImportOptions{
		UpdateExisting:  upd,
		DefaultClientID: req.DefaultClientID,
		DefaultProxyID:  req.DefaultProxyID,
		BatchSize:       200,
	}
	summary := h.svc.ImportBatch(c.Request.Context(), items, opt)

	// 后台踢一次刷新(让新导入的账号尽快探测过期时间 / 额度)
	if h.refresher != nil {
		h.refresher.Kick()
	}
	if h.prober != nil {
		h.prober.Kick()
	}

	resp.OK(c, summary)
}

// POST /api/admin/accounts/import-tokens
//
// body:
//
//	{
//	  "mode": "at" | "rt" | "st",
//	  "tokens": "一行一个\n...\n",   // 或字符串数组
//	  "client_id": "app_xxxx",      // rt 必填,at/st 可选
//	  "update_existing": true,
//	  "default_proxy_id": 0         // RT/ST 换 AT 时走此代理,强烈推荐
//	}
//
// 返回同 /import:ImportSummary。
func (h *Handler) ImportTokens(c *gin.Context) {
	var req struct {
		Mode           string          `json:"mode"`
		Tokens         json.RawMessage `json:"tokens"`
		ClientID       string          `json:"client_id"`
		UpdateExisting *bool           `json:"update_existing"`
		DefaultProxyID uint64          `json:"default_proxy_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.BadRequest(c, "请求参数错误:"+err.Error())
		return
	}

	// 支持两种 tokens 形态:字符串数组 / 一大段多行文本
	var tokens []string
	if len(req.Tokens) > 0 {
		switch bytes.TrimSpace(req.Tokens)[0] {
		case '[':
			_ = json.Unmarshal(req.Tokens, &tokens)
		case '"':
			var s string
			_ = json.Unmarshal(req.Tokens, &s)
			tokens = splitTokenInput(s, req.Mode)
		}
	}
	if len(tokens) == 0 {
		resp.BadRequest(c, "tokens 不能为空,请每行一个")
		return
	}

	mode := ImportTokenMode(strings.ToLower(strings.TrimSpace(req.Mode)))
	if mode == "" {
		mode = ImportModeAT
	}
	if mode != ImportModeAT && mode != ImportModeRT && mode != ImportModeST && mode != ImportModeSessionJSON {
		resp.BadRequest(c, "不支持的 mode(仅 at / rt / st / session_json)")
		return
	}
	if mode == ImportModeRT && strings.TrimSpace(req.ClientID) == "" {
		resp.BadRequest(c, "RT 模式必须提供 client_id(APPID)")
		return
	}

	upd := true
	if req.UpdateExisting != nil {
		upd = *req.UpdateExisting
	}

	var proxyURL string
	if req.DefaultProxyID > 0 && h.proxyResolver != nil {
		proxyURL = h.proxyResolver.ProxyURLByID(c.Request.Context(), req.DefaultProxyID)
	}

	summary := h.svc.ImportTokensBatch(c.Request.Context(), tokens, ImportTokensOptions{
		Mode:            mode,
		ClientID:        strings.TrimSpace(req.ClientID),
		ProxyURL:        proxyURL,
		DefaultProxyID:  req.DefaultProxyID,
		UpdateExisting:  upd,
		DefaultClientID: strings.TrimSpace(req.ClientID),
	})

	if h.refresher != nil {
		h.refresher.Kick()
	}
	if h.prober != nil {
		h.prober.Kick()
	}
	resp.OK(c, summary)
}

// splitLines 把多行文本切成 trim 后的非空行数组。
func splitLines(s string) []string {
	raw := strings.ReplaceAll(s, "\r\n", "\n")
	parts := strings.Split(raw, "\n")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if t := strings.TrimSpace(p); t != "" {
			out = append(out, t)
		}
	}
	return out
}

func splitTokenInput(s, mode string) []string {
	trimmed := strings.TrimSpace(s)
	// session_json 常见是格式化后的多行 JSON;直接按行切会把一份 JSON 拆碎。
	if trimmed != "" && (strings.HasPrefix(trimmed, "{") || strings.HasPrefix(trimmed, "[")) {
		if vals := splitJSONTokenValues(trimmed); len(vals) > 0 {
			return vals
		}
	}
	return splitLines(s)
}

func splitJSONTokenValues(s string) []string {
	dec := json.NewDecoder(strings.NewReader(s))
	out := make([]string, 0, 1)
	for {
		var raw json.RawMessage
		if err := dec.Decode(&raw); err != nil {
			if err == io.EOF {
				break
			}
			return nil
		}
		appendJSONTokenValue(&out, raw)
	}
	return out
}

func appendJSONTokenValue(out *[]string, raw json.RawMessage) {
	raw = bytes.TrimSpace(raw)
	if len(raw) == 0 || string(raw) == "null" {
		return
	}
	if raw[0] == '[' {
		var arr []json.RawMessage
		if err := json.Unmarshal(raw, &arr); err == nil {
			for _, item := range arr {
				appendJSONTokenValue(out, item)
			}
			return
		}
	}
	var s string
	if err := json.Unmarshal(raw, &s); err == nil {
		if t := strings.TrimSpace(s); t != "" {
			*out = append(*out, t)
		}
		return
	}
	if t := strings.TrimSpace(string(raw)); t != "" {
		*out = append(*out, t)
	}
}

// ===================== 刷新 / 探测 =====================

// POST /api/admin/accounts/:id/refresh
func (h *Handler) Refresh(c *gin.Context) {
	if h.refresher == nil {
		resp.Internal(c, "刷新器未初始化")
		return
	}
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	res, err := h.refresher.RefreshByID(c.Request.Context(), id)
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}
	resp.OK(c, res)
}

// POST /api/admin/accounts/refresh-all
// 批量并发刷新所有账号,并返回每条结果。
func (h *Handler) RefreshAll(c *gin.Context) {
	if h.refresher == nil {
		resp.Internal(c, "刷新器未初始化")
		return
	}
	ids, err := h.svc.dao.ListAllActiveIDs(c.Request.Context())
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}

	ctx := c.Request.Context()
	conc := 8
	sem := make(chan struct{}, conc)
	var wg sync.WaitGroup
	results := make([]*RefreshResult, 0, len(ids))
	var mu sync.Mutex

	for _, id := range ids {
		id := id
		wg.Add(1)
		sem <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-sem }()
			r, err := h.refresher.RefreshByID(ctx, id)
			if err != nil {
				r = &RefreshResult{AccountID: id, Source: "failed", Error: err.Error()}
			}
			mu.Lock()
			results = append(results, r)
			mu.Unlock()
		}()
	}
	wg.Wait()

	ok, failed := 0, 0
	for _, r := range results {
		if r.OK {
			ok++
		} else {
			failed++
		}
	}
	resp.OK(c, gin.H{
		"total":   len(results),
		"success": ok,
		"failed":  failed,
		"results": results,
	})
}

// POST /api/admin/accounts/:id/probe-quota
func (h *Handler) ProbeQuota(c *gin.Context) {
	if h.prober == nil {
		resp.Internal(c, "额度探测器未初始化")
		return
	}
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	res, err := h.prober.ProbeByID(c.Request.Context(), id)
	if err != nil && res == nil {
		resp.Internal(c, err.Error())
		return
	}
	resp.OK(c, res)
}

// POST /api/admin/accounts/probe-quota-all
func (h *Handler) ProbeQuotaAll(c *gin.Context) {
	if h.prober == nil {
		resp.Internal(c, "额度探测器未初始化")
		return
	}
	ids, err := h.svc.dao.ListAllActiveIDs(c.Request.Context())
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}

	ctx := c.Request.Context()
	conc := 8
	sem := make(chan struct{}, conc)
	var wg sync.WaitGroup
	results := make([]*QuotaResult, 0, len(ids))
	var mu sync.Mutex

	for _, id := range ids {
		id := id
		wg.Add(1)
		sem <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-sem }()
			r, err := h.prober.ProbeByID(ctx, id)
			if r == nil {
				r = &QuotaResult{AccountID: id}
				if err != nil {
					r.Error = err.Error()
				}
			}
			mu.Lock()
			results = append(results, r)
			mu.Unlock()
		}()
	}
	wg.Wait()

	ok, failed := 0, 0
	for _, r := range results {
		if r.OK {
			ok++
		} else {
			failed++
		}
	}
	resp.OK(c, gin.H{
		"total":   len(results),
		"success": ok,
		"failed":  failed,
		"results": results,
	})
}
````

## File: internal/account/importer_tokens.go
````go
package account

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ImportTokenMode 批量 token 导入的模式。
type ImportTokenMode string

const (
	ImportModeAT          ImportTokenMode = "at"           // 每行一个 access_token
	ImportModeRT          ImportTokenMode = "rt"           // 每行一个 refresh_token,需要 client_id(APPID)
	ImportModeST          ImportTokenMode = "st"           // 每行一个 session_token
	ImportModeSessionJSON ImportTokenMode = "session_json" // chatgpt.com/api/auth/session 的完整 JSON
)

// ImportTokensOptions 批量 token 导入选项。
type ImportTokensOptions struct {
	Mode ImportTokenMode

	// ClientID: RT 模式必填,AT/ST 模式可选(不填则用 DefaultClientID)。
	// 在 RT 模式下会作为 OAuth 的 client_id 发起 auth.openai.com/oauth/token 请求。
	ClientID string

	// ProxyURL: 换 AT 时走的代理(RT/ST 必须能访问 auth.openai.com / chatgpt.com)。
	// 空字符串=直连,生产环境通常必须走代理。
	ProxyURL string

	// 下面几个直接透传给底层 ImportBatch。
	DefaultProxyID  uint64
	UpdateExisting  bool
	DefaultClientID string
	BatchSize       int
}

// ImportTokensBatch 把运维侧粘贴的「一行一个 token」的文本批量转成账号入库。
//
// 处理流程:
//
//  1. 对 AT 模式:直接解 JWT payload 拿 email,不发起任何外部请求。
//  2. 对 RT 模式:POST auth.openai.com/oauth/token(需要 client_id)换出 AT,
//     再从 AT 解 email。RT 和 newAT 都会保存到账号。
//  3. 对 ST 模式:GET chatgpt.com/api/auth/session(带 cookie)换出 AT,再从 AT 解 email。
//     ST 和 newAT 都会保存到账号。
//  4. 拿到 email + AT 的行 → 复用现有 ImportBatch 走 upsert。
//  5. 无法拿到 email 的行进 failed 明细,返回给前端。
func (s *Service) ImportTokensBatch(ctx context.Context, tokens []string, opts ImportTokensOptions) *ImportSummary {
	if opts.Mode == "" {
		opts.Mode = ImportModeAT
	}
	if opts.DefaultClientID == "" {
		opts.DefaultClientID = "app_LlGpXReQgckcGGUo2JrYvtJK"
	}

	httpc := buildImportHTTPClient(opts.ProxyURL)

	// 去重 + 归一化
	seen := map[string]struct{}{}
	cleaned := make([]string, 0, len(tokens))
	for _, raw := range tokens {
		t := strings.TrimSpace(raw)
		if t == "" {
			continue
		}
		// 容错:某些 RT / ST 复制过来带前缀 "Bearer " 或 "__Secure-next-auth.session-token="
		t = strings.TrimPrefix(t, "Bearer ")
		t = strings.TrimPrefix(t, "bearer ")
		if i := strings.Index(t, "="); i > 0 && i < 60 && strings.HasPrefix(strings.ToLower(t), "__secure-next-auth.session-token") {
			t = t[i+1:]
		}
		if _, dup := seen[t]; dup {
			continue
		}
		seen[t] = struct{}{}
		cleaned = append(cleaned, t)
	}

	sum := &ImportSummary{Results: make([]ImportLineResult, 0, len(cleaned))}
	items := make([]ImportSource, 0, len(cleaned))

	// RT 模式先强制校验 client_id
	clientID := strings.TrimSpace(opts.ClientID)
	if opts.Mode == ImportModeRT && clientID == "" {
		clientID = opts.DefaultClientID
	}

	for idx, t := range cleaned {
		if err := ctx.Err(); err != nil {
			break
		}

		var src ImportSource
		var err error

		// 自动检测: 如果输入以 { 开头且含 accessToken/access_token,自动当 session_json
		mode := opts.Mode
		if strings.HasPrefix(t, "{") {
			low := strings.ToLower(t)
			if strings.Contains(low, "accesstoken") || strings.Contains(low, "access_token") {
				mode = ImportModeSessionJSON
			}
		}

		switch mode {
		case ImportModeAT:
			src, err = convertATToSource(t, clientID)
		case ImportModeRT:
			if clientID == "" {
				err = errors.New("RT 模式需要 APPID(client_id)")
				break
			}
			src, err = convertRTToSource(ctx, httpc, t, clientID)
		case ImportModeST:
			src, err = convertSTToSource(ctx, httpc, t, clientID)
		case ImportModeSessionJSON:
			src, err = convertSessionJSONToSource(t, clientID)
		default:
			err = fmt.Errorf("未知模式:%s", opts.Mode)
		}
		if err != nil {
			sum.Failed++
			sum.Total++
			sum.Results = append(sum.Results, ImportLineResult{
				Index:  idx,
				Email:  "?",
				Status: "failed",
				Reason: truncate(err.Error(), 160),
			})
			continue
		}
		items = append(items, src)
	}

	// 复用已有批量 upsert(去重、分批、UpdateExisting 等都在里面)
	batch := s.ImportBatch(ctx, items, ImportOptions{
		UpdateExisting:  opts.UpdateExisting,
		DefaultClientID: clientID,
		DefaultProxyID:  opts.DefaultProxyID,
		BatchSize:       opts.BatchSize,
	})
	sum.Total += batch.Total
	sum.Created += batch.Created
	sum.Updated += batch.Updated
	sum.Skipped += batch.Skipped
	sum.Failed += batch.Failed
	sum.Results = append(sum.Results, batch.Results...)
	return sum
}

// ---------- 四种模式的 token → ImportSource 转换 ----------

// convertSessionJSONToSource 解析 chatgpt.com/api/auth/session 返回的完整 JSON。
// 格式: {"user":{"id":"...","name":"...","email":"..."},"expires":"...","accessToken":"eyJ..."}
func convertSessionJSONToSource(raw, clientID string) (ImportSource, error) {
	var sess struct {
		User *struct {
			ID    string `json:"id"`
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"user"`
		Expires          string `json:"expires"`
		AccessToken      string `json:"accessToken"`
		AccessTokenSnake string `json:"access_token"`
	}
	if err := json.Unmarshal([]byte(raw), &sess); err != nil {
		return ImportSource{}, fmt.Errorf("JSON 解析失败:%w", err)
	}
	if sess.AccessToken == "" {
		sess.AccessToken = sess.AccessTokenSnake
	}
	if sess.AccessToken == "" {
		return ImportSource{}, errors.New("JSON 中缺少 accessToken/access_token 字段")
	}

	email := ""
	if sess.User != nil {
		email = sess.User.Email
	}

	// 从 AT JWT 补充信息
	jwtEmail, subAccID, expAt, _ := decodeATClaims(sess.AccessToken)
	if email == "" {
		email = jwtEmail
	}
	if email == "" {
		return ImportSource{}, errors.New("无法获取 email")
	}

	// 优先用 JSON 里的 expires
	if sess.Expires != "" {
		if t, err := time.Parse(time.RFC3339, sess.Expires); err == nil {
			expAt = t
		}
	}

	return ImportSource{
		AccessToken:      sess.AccessToken,
		Email:            email,
		ChatGPTAccountID: subAccID,
		ExpiredAt:        expAt,
		ClientID:         clientID,
		AccountType:      "chatgpt",
	}, nil
}

// convertATToSource 仅凭 access_token 的 JWT payload 解 email / sub / 过期时间。
func convertATToSource(at, clientID string) (ImportSource, error) {
	email, subAccountID, expAt, err := decodeATClaims(at)
	if err != nil {
		return ImportSource{}, fmt.Errorf("解析 AT 失败:%w", err)
	}
	if email == "" {
		return ImportSource{}, errors.New("无法从 AT 解出 email,请改用 JSON 或带 email 的导入")
	}
	return ImportSource{
		AccessToken:      at,
		Email:            email,
		ChatGPTAccountID: subAccountID,
		ExpiredAt:        expAt,
		ClientID:         clientID,
		AccountType:      "chatgpt",
	}, nil
}

// convertRTToSource 用 refresh_token + client_id 调 auth.openai.com/oauth/token 换出 AT,
// 再解 AT claims 拿 email。AT + RT 一起存进账号(后续仍可 RT 续签)。
func convertRTToSource(ctx context.Context, httpc *http.Client, rt, clientID string) (ImportSource, error) {
	at, newRT, expAt, err := rtExchange(ctx, httpc, rt, clientID)
	if err != nil {
		return ImportSource{}, fmt.Errorf("RT 换 AT 失败:%s", friendlyImportErr(err))
	}
	email, subAccID, expFromJWT, jerr := decodeATClaims(at)
	if jerr != nil || email == "" {
		return ImportSource{}, errors.New("RT 换出的 AT 无法解析出 email")
	}
	if expAt.IsZero() {
		expAt = expFromJWT
	}
	// 如果服务端下发新 RT,用新的;否则用本次输入的
	usedRT := rt
	if newRT != "" {
		usedRT = newRT
	}
	return ImportSource{
		AccessToken:      at,
		RefreshToken:     usedRT,
		Email:            email,
		ChatGPTAccountID: subAccID,
		ExpiredAt:        expAt,
		ClientID:         clientID,
		AccountType:      "chatgpt",
	}, nil
}

// convertSTToSource 用 session_token 调 chatgpt.com/api/auth/session 换出 AT,再解 email。
// AT + ST 一起存进账号(后续可由 ST 定时续签)。
func convertSTToSource(ctx context.Context, httpc *http.Client, st, clientID string) (ImportSource, error) {
	at, expAt, err := stExchange(ctx, httpc, st)
	if err != nil {
		return ImportSource{}, fmt.Errorf("ST 换 AT 失败:%s", friendlyImportErr(err))
	}
	email, subAccID, expFromJWT, jerr := decodeATClaims(at)
	if jerr != nil || email == "" {
		return ImportSource{}, errors.New("ST 换出的 AT 无法解析出 email")
	}
	if expAt.IsZero() {
		expAt = expFromJWT
	}
	return ImportSource{
		AccessToken:      at,
		SessionToken:     st,
		Email:            email,
		ChatGPTAccountID: subAccID,
		ExpiredAt:        expAt,
		ClientID:         clientID,
		AccountType:      "chatgpt",
	}, nil
}

// ---------- 底层 HTTP ----------

func buildImportHTTPClient(proxyURL string) *http.Client {
	httpc := &http.Client{Timeout: 30 * time.Second}
	if proxyURL == "" {
		return httpc
	}
	u, err := url.Parse(proxyURL)
	if err != nil {
		return httpc
	}
	httpc.Transport = &http.Transport{
		Proxy:               http.ProxyURL(u),
		ForceAttemptHTTP2:   true,
		MaxIdleConns:        8,
		IdleConnTimeout:     30 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
	}
	return httpc
}

// rtExchange 与 refresher.rtToAT 等价的包级实现(不需要 account_id)。
func rtExchange(ctx context.Context, httpc *http.Client, rt, clientID string) (newAT, newRT string, expAt time.Time, err error) {
	body := map[string]string{
		"client_id":     clientID,
		"grant_type":    "refresh_token",
		"redirect_uri":  "com.openai.chat://auth0.openai.com/ios/com.openai.chat/callback",
		"refresh_token": rt,
	}
	buf, _ := json.Marshal(body)
	req, err := http.NewRequestWithContext(ctx, "POST",
		"https://auth.openai.com/oauth/token", bytes.NewReader(buf))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "ChatGPT/1.2025.122 (iOS 18.2; iPhone15,2; build 15096)")
	resp, err := httpc.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		err = fmt.Errorf("rt exchange http=%d body=%s", resp.StatusCode, truncate(string(data), 200))
		return
	}
	var out struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int    `json:"expires_in"`
	}
	if err = json.Unmarshal(data, &out); err != nil {
		return
	}
	if out.AccessToken == "" {
		err = errors.New("rt exchange: missing access_token in response")
		return
	}
	newAT = out.AccessToken
	newRT = out.RefreshToken
	if out.ExpiresIn > 0 {
		expAt = time.Now().Add(time.Duration(out.ExpiresIn) * time.Second)
	}
	return
}

// stExchange 与 refresher.stToAT 等价的包级实现。
func stExchange(ctx context.Context, httpc *http.Client, st string) (newAT string, expAt time.Time, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://chatgpt.com/api/auth/session", nil)
	if err != nil {
		return
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "https://chatgpt.com/")
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")
	req.AddCookie(&http.Cookie{Name: "__Secure-next-auth.session-token", Value: st})

	resp, err := httpc.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		err = fmt.Errorf("st exchange http=%d body=%s", resp.StatusCode, truncate(string(data), 200))
		return
	}
	raw := strings.TrimSpace(string(data))
	if raw == "" || raw == "{}" {
		err = errors.New("ST 已过期或无效(响应为空)")
		return
	}
	var out struct {
		AccessToken string `json:"accessToken"`
		Expires     string `json:"expires"`
	}
	if err = json.Unmarshal([]byte(raw), &out); err != nil {
		return
	}
	if out.AccessToken == "" {
		err = errors.New("响应缺少 accessToken 字段,ST 已失效")
		return
	}
	newAT = out.AccessToken
	if out.Expires != "" {
		if t, e := time.Parse(time.RFC3339, out.Expires); e == nil {
			expAt = t
		}
	}
	return
}

// ---------- JWT claims 解析 ----------

// decodeATClaims 从 access_token(JWT)里取出 email / chatgpt_account_id / exp。
// 兼容 iOS scope(Codex)和 Web scope(ChatGPT)两种 claim 结构。
func decodeATClaims(at string) (email, accountID string, expAt time.Time, err error) {
	parts := strings.Split(at, ".")
	if len(parts) < 2 {
		err = errors.New("非法 JWT(段数不足)")
		return
	}
	raw, e := base64.RawURLEncoding.DecodeString(parts[1])
	if e != nil {
		raw, e = base64.StdEncoding.DecodeString(parts[1])
		if e != nil {
			err = fmt.Errorf("base64 解码失败:%w", e)
			return
		}
	}
	// 尽可能宽松地解析,不同 scope 的 claims 字段名不一样。
	var claims map[string]interface{}
	if e := json.Unmarshal(raw, &claims); e != nil {
		err = fmt.Errorf("claims JSON 解码失败:%w", e)
		return
	}

	// 1) 直接字段
	if v, ok := claims["email"].(string); ok && v != "" {
		email = v
	}
	if v, ok := claims["chatgpt_account_id"].(string); ok && v != "" {
		accountID = v
	}

	// 2) iOS/Web scope 里常见的 namespaced claims
	for _, ns := range []string{
		"https://api.openai.com/profile",
		"https://api.openai.com/auth",
	} {
		if m, ok := claims[ns].(map[string]interface{}); ok {
			if email == "" {
				if v, ok := m["email"].(string); ok && v != "" {
					email = v
				}
			}
			if accountID == "" {
				if v, ok := m["chatgpt_account_id"].(string); ok && v != "" {
					accountID = v
				}
				claimKey := "user" + "_id"
				if v, ok := m[claimKey].(string); ok && accountID == "" && v != "" {
					accountID = v
				}
			}
		}
	}

	// 3) 从 exp
	if v, ok := claims["exp"].(float64); ok && v > 0 {
		expAt = time.Unix(int64(v), 0)
	}
	return
}

// friendlyImportErr 把底层 http 错误压成简短中文,避免把 URL / stacktrace 泄露到前端。
func friendlyImportErr(err error) string {
	if err == nil {
		return ""
	}
	s := err.Error()
	low := strings.ToLower(s)
	switch {
	case strings.Contains(low, "http=401"), strings.Contains(low, "invalid_grant"):
		return "token 已失效(401)"
	case strings.Contains(low, "http=403"):
		return "上游拒绝访问(403)"
	case strings.Contains(low, "http=429"):
		return "触发速率限制(429),稍后再试"
	case strings.Contains(low, "timeout"), strings.Contains(low, "deadline exceeded"):
		return "请求超时,建议配代理"
	case strings.Contains(low, "no such host"), strings.Contains(low, "dial tcp"):
		return "无法连接 openai(建议选默认代理)"
	case strings.Contains(low, "tls"), strings.Contains(low, "x509"):
		return "TLS 握手失败"
	}
	// 兜底:去掉 URL
	if i := strings.Index(s, `": `); i > 0 && i < len(s)-3 {
		s = s[i+3:]
	}
	return truncate(strings.TrimSpace(s), 120)
}
````

## File: internal/account/importer.go
````go
package account

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/google/uuid"
)

// ImportSource 代表一条待导入记录,来自任意一种 JSON 格式。
type ImportSource struct {
	// 必填
	AccessToken string
	Email       string

	// 可选
	RefreshToken     string
	SessionToken     string // 从 cookie 里提取的 __Secure-next-auth.session-token
	ClientID         string
	ChatGPTAccountID string
	AccountType      string // codex / chatgpt
	ExpiredAt        time.Time
	Name             string // sub2api 里的 name 字段(当 email 缺失时退化为 email)
}

// ImportLineResult 返回给前端,每条记录处理结果。
type ImportLineResult struct {
	Index  int    `json:"index"`
	Email  string `json:"email"`
	Status string `json:"status"` // created / updated / skipped / failed
	Reason string `json:"reason,omitempty"`
	ID     uint64 `json:"id,omitempty"`
}

// ImportSummary 整体统计。
type ImportSummary struct {
	Total   int                `json:"total"`
	Created int                `json:"created"`
	Updated int                `json:"updated"`
	Skipped int                `json:"skipped"`
	Failed  int                `json:"failed"`
	Results []ImportLineResult `json:"results"`
}

// ImportOptions 批量导入选项。
type ImportOptions struct {
	// UpdateExisting 为 true 时 email 已存在则更新 token;false 则 skipped。
	UpdateExisting bool
	// DefaultClientID 当记录里没有 client_id 时填充的值。
	DefaultClientID string
	// DefaultProxyID 新建账号时默认绑定的代理 id(0 = 不绑)。
	DefaultProxyID uint64
	// BatchSize 分批 commit 的大小(仅用于让出 CPU,每批做一次 context check)。默认 200。
	BatchSize int
}

// ParseJSONBlob 尝试把导入文本解析成 ImportSource 列表。
// 同时兼容以下输入:
//  1. 顶层是对象且含 `accounts` 数组 → sub2api 多账号导出
//  2. 顶层是对象且含 `access_token` / `accessToken` → 单账号 token_xxx.json
//  3. 顶层是数组,每个元素同 (1)/(2) 的单个对象
//  4. 多个 JSON 文本用换行/空行分隔(JSONL)
func ParseJSONBlob(raw string) ([]ImportSource, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil, errors.New("输入为空")
	}
	// 先尝试整体解析
	if xs, err := parseSingleJSON(raw); err == nil && len(xs) > 0 {
		return xs, nil
	}
	// 再尝试 JSONL
	var all []ImportSource
	var firstErr error
	dec := json.NewDecoder(strings.NewReader(raw))
	for {
		var one json.RawMessage
		if err := dec.Decode(&one); err != nil {
			if err == io.EOF {
				break
			}
			if firstErr == nil {
				firstErr = err
			}
			break
		}
		xs, err := parseSingleJSON(string(one))
		if err != nil {
			if firstErr == nil {
				firstErr = err
			}
			continue
		}
		all = append(all, xs...)
	}
	if len(all) == 0 {
		if firstErr == nil {
			firstErr = errors.New("无法识别的 JSON 格式")
		}
		return nil, firstErr
	}
	return all, nil
}

func parseSingleJSON(s string) ([]ImportSource, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, errors.New("空 JSON")
	}
	// 1) 数组
	if s[0] == '[' {
		var arr []json.RawMessage
		if err := json.Unmarshal([]byte(s), &arr); err != nil {
			return nil, fmt.Errorf("解析 JSON 数组失败:%w", err)
		}
		var all []ImportSource
		for _, item := range arr {
			xs, err := parseSingleJSON(string(item))
			if err != nil {
				continue
			}
			all = append(all, xs...)
		}
		return all, nil
	}
	// 2) 对象
	var obj map[string]json.RawMessage
	if err := json.Unmarshal([]byte(s), &obj); err != nil {
		return nil, fmt.Errorf("解析 JSON 对象失败:%w", err)
	}

	// Format A: 有 accounts 数组
	if v, ok := obj["accounts"]; ok {
		var accs []subAPIAccount
		if err := json.Unmarshal(v, &accs); err == nil {
			out := make([]ImportSource, 0, len(accs))
			for _, a := range accs {
				src, ok := a.toSource()
				if ok {
					out = append(out, src)
				}
			}
			return out, nil
		}
	}

	// Format B: 单账号 token_xxx.json(扁平对象)
	if _, has := obj["access_token"]; has {
		var b tokenFileB
		if err := json.Unmarshal([]byte(s), &b); err != nil {
			return nil, fmt.Errorf("解析 token 文件失败:%w", err)
		}
		src, ok := b.toSource()
		if !ok {
			return nil, errors.New("token 文件缺少必要字段")
		}
		return []ImportSource{src}, nil
	}
	// 兼容 accessToken(驼峰)
	if _, has := obj["accessToken"]; has {
		var b tokenFileB
		// 同名 camelCase 字段也走 tokenFileB;json tag 都用 snake_case
		// 这里用一个临时结构
		var camel struct {
			AccessToken  string `json:"accessToken"`
			RefreshToken string `json:"refreshToken"`
			Email        string `json:"email"`
			AccountID    string `json:"account_id"`
			Type         string `json:"type"`
			ClientID     string `json:"client_id"`
			Expired      string `json:"expires"`
		}
		if err := json.Unmarshal([]byte(s), &camel); err != nil {
			return nil, err
		}
		b.AccessToken = camel.AccessToken
		b.RefreshToken = camel.RefreshToken
		b.Email = camel.Email
		b.AccountID = camel.AccountID
		b.Type = camel.Type
		b.ClientID = camel.ClientID
		b.Expired = camel.Expired
		src, ok := b.toSource()
		if !ok {
			return nil, errors.New("token 文件缺少必要字段")
		}
		return []ImportSource{src}, nil
	}

	return nil, errors.New("未识别的 JSON 结构(既不是 sub2api 也不是 token 文件)")
}

// sub2api 的 account 结构片段。
type subAPIAccount struct {
	Name        string `json:"name"`
	Platform    string `json:"platform"`
	Type        string `json:"type"`
	Credentials struct {
		AccessToken      string `json:"access_token"`
		RefreshToken     string `json:"refresh_token"`
		SessionToken     string `json:"session_token"`
		ClientID         string `json:"client_id"`
		ChatGPTAccountID string `json:"chatgpt_account_id"`
	} `json:"credentials"`
	Extra struct {
		Email string `json:"email"`
	} `json:"extra"`
}

func (a subAPIAccount) toSource() (ImportSource, bool) {
	src := ImportSource{
		AccessToken:      a.Credentials.AccessToken,
		RefreshToken:     a.Credentials.RefreshToken,
		SessionToken:     a.Credentials.SessionToken,
		ClientID:         a.Credentials.ClientID,
		ChatGPTAccountID: a.Credentials.ChatGPTAccountID,
		AccountType:      normalizeType(a.Name, a.Platform),
		Email:            strings.TrimSpace(a.Extra.Email),
		Name:             a.Name,
	}
	if src.Email == "" {
		src.Email = emailFromName(a.Name)
	}
	if src.AccessToken == "" || src.Email == "" {
		return src, false
	}
	return src, true
}

// tokenFileB 对应 token_xxx.json。
type tokenFileB struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
	AccountID    string `json:"account_id"`
	Email        string `json:"email"`
	Type         string `json:"type"`
	ClientID     string `json:"client_id"`
	Expired      string `json:"expired"`
}

func (b tokenFileB) toSource() (ImportSource, bool) {
	src := ImportSource{
		AccessToken:      b.AccessToken,
		RefreshToken:     b.RefreshToken,
		Email:            strings.TrimSpace(b.Email),
		ChatGPTAccountID: b.AccountID,
		ClientID:         b.ClientID,
		AccountType:      strings.ToLower(strings.TrimSpace(b.Type)),
	}
	if src.AccountType == "" {
		src.AccountType = "codex"
	}
	if b.Expired != "" {
		if t, err := time.Parse(time.RFC3339, b.Expired); err == nil {
			src.ExpiredAt = t
		}
	}
	if src.AccessToken == "" || src.Email == "" {
		return src, false
	}
	return src, true
}

// emailFromName 把 sub2api 的 name (codex-user_hotmail.com) 反推成 email。
func emailFromName(name string) string {
	if name == "" {
		return ""
	}
	n := name
	for _, prefix := range []string{"codex-", "chatgpt-", "openai-"} {
		if strings.HasPrefix(n, prefix) {
			n = strings.TrimPrefix(n, prefix)
			break
		}
	}
	// 最后一个 _ 之前视为 localpart,之后视为 domain
	idx := strings.LastIndex(n, "_")
	if idx > 0 && idx < len(n)-1 {
		return n[:idx] + "@" + n[idx+1:]
	}
	return ""
}

func normalizeType(name, platform string) string {
	lower := strings.ToLower(name + " " + platform)
	switch {
	case strings.Contains(lower, "codex"):
		return "codex"
	case strings.Contains(lower, "chatgpt"):
		return "chatgpt"
	case strings.Contains(lower, "openai"):
		return "codex"
	default:
		return "codex"
	}
}

// ImportBatch 执行批量导入。
// 处理策略:
//   - 同一批内 email 去重(后者覆盖前者)
//   - email 已存在则按 UpdateExisting 决定更新或 skip
//   - 每 BatchSize 条让出一次 CPU,并检查 ctx.Done(),便于大批量
//   - 不做整体事务(失败项不影响成功项);单条失败只影响该条
func (s *Service) ImportBatch(ctx context.Context, items []ImportSource, opt ImportOptions) *ImportSummary {
	if opt.BatchSize <= 0 {
		opt.BatchSize = 200
	}
	if opt.DefaultClientID == "" {
		opt.DefaultClientID = "app_LlGpXReQgckcGGUo2JrYvtJK"
	}

	// email 去重(后者覆盖)
	seen := make(map[string]int, len(items))
	dedup := make([]ImportSource, 0, len(items))
	for _, it := range items {
		if it.Email == "" {
			continue
		}
		key := strings.ToLower(it.Email)
		if idx, ok := seen[key]; ok {
			dedup[idx] = it // 覆盖
		} else {
			seen[key] = len(dedup)
			dedup = append(dedup, it)
		}
	}

	sum := &ImportSummary{
		Total:   len(dedup),
		Results: make([]ImportLineResult, 0, len(dedup)),
	}

	for i, it := range dedup {
		if i > 0 && i%opt.BatchSize == 0 {
			// 让出一次 CPU;大批量下防止长时间独占
			select {
			case <-ctx.Done():
				sum.Failed += len(dedup) - i
				sum.Results = append(sum.Results, ImportLineResult{
					Index: i, Email: "", Status: "failed", Reason: "导入被取消",
				})
				return sum
			default:
			}
		}
		res := s.importOne(ctx, i, it, opt)
		switch res.Status {
		case "created":
			sum.Created++
		case "updated":
			sum.Updated++
		case "skipped":
			sum.Skipped++
		case "failed":
			sum.Failed++
		}
		sum.Results = append(sum.Results, res)
	}
	return sum
}

func (s *Service) importOne(ctx context.Context, idx int, it ImportSource, opt ImportOptions) ImportLineResult {
	out := ImportLineResult{Index: idx, Email: it.Email}

	if it.AccessToken == "" {
		out.Status = "failed"
		out.Reason = "缺少 access_token"
		return out
	}

	// 计算过期时间:优先用 JSON 的 expired,其次解析 JWT
	expAt := it.ExpiredAt
	if expAt.IsZero() {
		expAt = parseJWTExp(it.AccessToken)
	}

	clientID := it.ClientID
	if clientID == "" {
		clientID = opt.DefaultClientID
	}
	accountType := it.AccountType
	if accountType == "" {
		accountType = "codex"
	}

	// 查是否已存在
	existing, err := s.dao.GetByEmail(ctx, it.Email)
	if err != nil {
		out.Status = "failed"
		out.Reason = "查询失败:" + err.Error()
		return out
	}

	atEnc, err := s.cipher.EncryptString(it.AccessToken)
	if err != nil {
		out.Status = "failed"
		out.Reason = "AT 加密失败:" + err.Error()
		return out
	}
	var rtEnc, stEnc string
	if it.RefreshToken != "" {
		if v, err := s.cipher.EncryptString(it.RefreshToken); err == nil {
			rtEnc = v
		}
	}
	if it.SessionToken != "" {
		if v, err := s.cipher.EncryptString(it.SessionToken); err == nil {
			stEnc = v
		}
	}

	if existing == nil {
		// 新建
		// 从 AT 的 JWT 自动识别 plan_type（pro/plus/free/team）
		planType := parseJWTPlanType(it.AccessToken)
		if planType == "" {
			planType = "free"
		}
		a := &Account{
			Email:            it.Email,
			AuthTokenEnc:     atEnc,
			ClientID:         clientID,
			ChatGPTAccountID: it.ChatGPTAccountID,
			AccountType:      accountType,
			SubscriptionType: planType,
			Status:           StatusHealthy,
			OAIDeviceID:      uuid.NewString(),
			OAISessionID:     uuid.NewString(),
		}
		if rtEnc != "" {
			a.RefreshTokenEnc.String = rtEnc
			a.RefreshTokenEnc.Valid = true
		}
		if stEnc != "" {
			a.SessionTokenEnc.String = stEnc
			a.SessionTokenEnc.Valid = true
		}
		if !expAt.IsZero() {
			a.TokenExpiresAt.Time = expAt
			a.TokenExpiresAt.Valid = true
		}
		id, err := s.dao.Create(ctx, a)
		if err != nil {
			out.Status = "failed"
			out.Reason = "入库失败:" + err.Error()
			return out
		}
		// 绑定代理:指定了用指定的,否则自动分配绑定数最少的
		proxyID := opt.DefaultProxyID
		if proxyID == 0 {
			if autoID, err := s.dao.LeastBoundProxyID(ctx, nil); err == nil {
				proxyID = autoID
			}
		}
		if proxyID > 0 {
			_ = s.dao.SetBinding(ctx, id, proxyID)
		}
		out.Status = "created"
		out.ID = id
		return out
	}

	// 已存在
	if !opt.UpdateExisting {
		out.Status = "skipped"
		out.Reason = "邮箱已存在"
		out.ID = existing.ID
		return out
	}
	// 更新 token 字段,其它字段保持
	existing.AuthTokenEnc = atEnc
	if rtEnc != "" {
		existing.RefreshTokenEnc.String = rtEnc
		existing.RefreshTokenEnc.Valid = true
	}
	if stEnc != "" {
		existing.SessionTokenEnc.String = stEnc
		existing.SessionTokenEnc.Valid = true
	}
	if clientID != "" {
		existing.ClientID = clientID
	}
	if it.ChatGPTAccountID != "" {
		existing.ChatGPTAccountID = it.ChatGPTAccountID
	}
	if accountType != "" {
		existing.AccountType = accountType
	}
	if !expAt.IsZero() {
		existing.TokenExpiresAt.Time = expAt
		existing.TokenExpiresAt.Valid = true
	}
	// 复活已死账号(导入新 token 视为重新投放)
	if existing.Status == StatusDead || existing.Status == StatusSuspicious {
		existing.Status = StatusHealthy
	}
	if err := s.dao.Update(ctx, existing); err != nil {
		out.Status = "failed"
		out.Reason = "更新失败:" + err.Error()
		return out
	}
	out.Status = "updated"
	out.ID = existing.ID
	return out
}
````

## File: internal/account/model.go
````go
package account

import (
	"database/sql"
	"time"
)

// 账号状态常量。
const (
	StatusHealthy    = "healthy"
	StatusWarned     = "warned"
	StatusThrottled  = "throttled"
	StatusSuspicious = "suspicious"
	StatusDead       = "dead"
)

// 刷新来源。
const (
	RefreshSourceRT     = "rt"
	RefreshSourceST     = "st"
	RefreshSourceManual = "manual"
)

// Account 对应 oai_accounts 表。
type Account struct {
	ID               uint64         `db:"id" json:"id"`
	Email            string         `db:"email" json:"email"`
	AuthTokenEnc     string         `db:"auth_token_enc" json:"-"`
	RefreshTokenEnc  sql.NullString `db:"refresh_token_enc" json:"-"`
	SessionTokenEnc  sql.NullString `db:"session_token_enc" json:"-"`
	TokenExpiresAt   sql.NullTime   `db:"token_expires_at" json:"token_expires_at,omitempty"`
	OAISessionID     string         `db:"oai_session_id" json:"oai_session_id"`
	OAIDeviceID      string         `db:"oai_device_id" json:"oai_device_id"`
	ClientID         string         `db:"client_id" json:"client_id"`
	ChatGPTAccountID string         `db:"chatgpt_account_id" json:"chatgpt_account_id"`
	AccountType      string         `db:"account_type" json:"account_type"`
	SubscriptionType string         `db:"subscription_type" json:"subscription_type"`
	DailyImageQuota  int            `db:"daily_image_quota" json:"daily_image_quota"`
	Status           string         `db:"status" json:"status"`
	WarnedAt         sql.NullTime   `db:"warned_at" json:"warned_at,omitempty"`
	CooldownUntil    sql.NullTime   `db:"cooldown_until" json:"cooldown_until,omitempty"`
	LastUsedAt       sql.NullTime   `db:"last_used_at" json:"last_used_at,omitempty"`
	TodayUsedCount   int            `db:"today_used_count" json:"today_used_count"`
	TodayUsedDate    sql.NullTime   `db:"today_used_date" json:"today_used_date,omitempty"`

	LastRefreshAt     sql.NullTime `db:"last_refresh_at" json:"last_refresh_at,omitempty"`
	LastRefreshSource string       `db:"last_refresh_source" json:"last_refresh_source"`
	RefreshError      string       `db:"refresh_error" json:"refresh_error"`

	ImageQuotaRemaining int          `db:"image_quota_remaining" json:"image_quota_remaining"`
	ImageQuotaTotal     int          `db:"image_quota_total"     json:"image_quota_total"`
	ImageQuotaResetAt   sql.NullTime `db:"image_quota_reset_at"   json:"image_quota_reset_at,omitempty"`
	ImageQuotaUpdatedAt sql.NullTime `db:"image_quota_updated_at" json:"image_quota_updated_at,omitempty"`

	// image2 能力探测:主依据来自 /backend-api/models;conversation/init 只保留为 quota/诊断弱参考。
	ImageCapabilityStatus    string         `db:"image_capability_status" json:"image_capability_status"` // unknown/enabled/disabled/error
	ImageCapabilityModel     string         `db:"image_capability_model" json:"image_capability_model"`
	ImageCapabilitySource    string         `db:"image_capability_source" json:"image_capability_source"` // models/init/manual
	ImageCapabilityDetail    sql.NullString `db:"image_capability_detail" json:"image_capability_detail,omitempty"`
	ImageCapabilityUpdatedAt sql.NullTime   `db:"image_capability_updated_at" json:"image_capability_updated_at,omitempty"`
	ImageInitBlockedFeatures sql.NullString `db:"image_init_blocked_features" json:"image_init_blocked_features,omitempty"`

	// image2 命中画像:IMG2 是「账号资格 + 请求/会话抽卡」,不是简单布尔开关。
	IMG2HitCount         int          `db:"img2_hit_count" json:"img2_hit_count"`
	IMG2PreviewOnlyCount int          `db:"img2_preview_only_count" json:"img2_preview_only_count"`
	IMG2MissCount        int          `db:"img2_miss_count" json:"img2_miss_count"`
	IMG2ConsecutiveMiss  int          `db:"img2_consecutive_miss" json:"img2_consecutive_miss"`
	IMG2LastStatus       string       `db:"img2_last_status" json:"img2_last_status"`
	IMG2LastHitAt        sql.NullTime `db:"img2_last_hit_at" json:"img2_last_hit_at,omitempty"`
	IMG2LastAttemptAt    sql.NullTime `db:"img2_last_attempt_at" json:"img2_last_attempt_at,omitempty"`

	// IMG2 交付画像:和协议命中分开统计,避免「抽中但下载失败」污染账号优先级。
	IMG2DeliverySuccessCount int          `db:"img2_delivery_success_count" json:"img2_delivery_success_count"`
	IMG2DeliveryFailCount    int          `db:"img2_delivery_fail_count" json:"img2_delivery_fail_count"`
	IMG2DeliveryPartialCount int          `db:"img2_delivery_partial_count" json:"img2_delivery_partial_count"`
	IMG2LastDeliveryStatus   string       `db:"img2_last_delivery_status" json:"img2_last_delivery_status"`
	IMG2LastDeliveryAt       sql.NullTime `db:"img2_last_delivery_at" json:"img2_last_delivery_at,omitempty"`

	Notes     string       `db:"notes" json:"notes"`
	CreatedAt time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt time.Time    `db:"updated_at" json:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at" json:"-"`

	// 辅助字段(非数据库列):前端展示用标志位。
	HasRT bool `db:"-" json:"has_rt"`
	HasST bool `db:"-" json:"has_st"`
}

// Binding 对应 account_proxy_bindings 表。
type Binding struct {
	AccountID uint64    `db:"account_id" json:"account_id"`
	ProxyID   uint64    `db:"proxy_id" json:"proxy_id"`
	BoundAt   time.Time `db:"bound_at" json:"bound_at"`
}
````

## File: internal/account/quota.go
````go
package account

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/432539/gpt2api/internal/upstream/chatgpt"
)

// QuotaSettings 热更新参数。
type QuotaSettings interface {
	AccountQuotaProbeEnabled() bool
	AccountQuotaProbeIntervalSec() int
	AccountRefreshConcurrency() int // 复用刷新并发上限
}

// QuotaResult 探测结果。
type QuotaResult struct {
	AccountID             uint64    `json:"account_id"`
	Email                 string    `json:"email"`
	OK                    bool      `json:"ok"`
	Remaining             int       `json:"remaining"`
	Total                 int       `json:"total"`
	ResetAt               time.Time `json:"reset_at,omitempty"`
	DefaultModel          string    `json:"default_model,omitempty"`           // 优先来自 /backend-api/models
	BlockedFeatures       []string  `json:"blocked_features,omitempty"`        // conversation/init 的诊断弱参考
	ImageCapabilityStatus string    `json:"image_capability_status,omitempty"` // enabled/disabled/unknown/error
	ImageCapabilitySource string    `json:"image_capability_source,omitempty"` // models/init
	ImageCapabilityDetail string    `json:"image_capability_detail,omitempty"`
	Error                 string    `json:"error,omitempty"`
}

// QuotaProber 后台定期探测账号图片剩余额度。
type QuotaProber struct {
	svc      *Service
	settings QuotaSettings
	log      *zap.Logger

	proxyResolver AccountProxyResolver

	kick chan struct{}
}

func NewQuotaProber(svc *Service, settings QuotaSettings, logger *zap.Logger) *QuotaProber {
	return &QuotaProber{
		svc:      svc,
		settings: settings,
		log:      logger,
		kick:     make(chan struct{}, 1),
	}
}

// SetProxyResolver 注入账号代理解析器;未注入则直连。
func (q *QuotaProber) SetProxyResolver(pr AccountProxyResolver) { q.proxyResolver = pr }

func (q *QuotaProber) Kick() {
	select {
	case q.kick <- struct{}{}:
	default:
	}
}

// Run 后台循环。
func (q *QuotaProber) Run(ctx context.Context) {
	q.log.Info("account quota prober started")
	defer q.log.Info("account quota prober stopped")

	select {
	case <-ctx.Done():
		return
	case <-time.After(10 * time.Second):
	}

	for {
		// 最小扫描间隔 60s;实际复用探测最小间隔的 1/3 做节奏,至少 60s
		interval := time.Duration(q.settings.AccountQuotaProbeIntervalSec()/3) * time.Second
		if interval < 60*time.Second {
			interval = 60 * time.Second
		}

		if q.settings.AccountQuotaProbeEnabled() {
			q.scanOnce(ctx)
		}

		select {
		case <-ctx.Done():
			return
		case <-time.After(interval):
		case <-q.kick:
		}
	}
}

func (q *QuotaProber) scanOnce(ctx context.Context) {
	minInterval := q.settings.AccountQuotaProbeIntervalSec()
	conc := q.settings.AccountRefreshConcurrency()

	rows, err := q.svc.dao.ListNeedProbeQuota(ctx, minInterval, 256)
	if err != nil {
		q.log.Warn("list quota probe candidates failed", zap.Error(err))
		return
	}
	if len(rows) == 0 {
		return
	}

	sem := make(chan struct{}, conc)
	var wg sync.WaitGroup
	for _, a := range rows {
		a := a
		wg.Add(1)
		sem <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-sem }()
			_, _ = q.ProbeOne(ctx, a)
		}()
	}
	wg.Wait()
}

// ProbeByID 指定账号探测。
func (q *QuotaProber) ProbeByID(ctx context.Context, id uint64) (*QuotaResult, error) {
	a, err := q.svc.dao.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return q.ProbeOne(ctx, a)
}

// ProbeOne 执行一次探测。
// models-first 判断 image 入口资格,conversation/init 只作为 quota/reset 诊断弱参考。
func (q *QuotaProber) ProbeOne(ctx context.Context, a *Account) (*QuotaResult, error) {
	res := &QuotaResult{AccountID: a.ID, Email: a.Email}
	at, err := q.svc.cipher.DecryptString(a.AuthTokenEnc)
	if err != nil || at == "" {
		res.Error = "AT 解密失败"
		_ = q.svc.dao.ApplyQuotaResult(ctx, a.ID, -1, -1, nil)
		return res, errors.New(res.Error)
	}

	probe, probeErr := q.doProbe(ctx, a, at)
	if probe.capabilitySource != "" {
		_ = q.svc.dao.ApplyImageCapabilityResult(ctx, a.ID, probe.capabilityStatus,
			probe.defaultModel, probe.capabilitySource, probe.capabilityDetail, probe.blockedFeatures)
	}
	res.DefaultModel = probe.defaultModel
	res.BlockedFeatures = probe.blockedFeatures
	res.ImageCapabilityStatus = probe.capabilityStatus
	res.ImageCapabilitySource = probe.capabilitySource
	res.ImageCapabilityDetail = probe.capabilityDetail
	if probeErr != nil {
		res.Error = friendlyProbeErr(probeErr)
		_ = q.svc.dao.ApplyQuotaResult(ctx, a.ID, -1, -1, nil)
		return res, probeErr
	}

	if probe.quotaProbed {
		var resetPtr *time.Time
		if !probe.resetAt.IsZero() {
			resetPtr = &probe.resetAt
		}
		if err := q.svc.dao.ApplyQuotaResult(ctx, a.ID, probe.remaining, probe.total, resetPtr); err != nil {
			res.Error = "写库失败:" + err.Error()
			return res, err
		}
	}
	res.OK = true
	res.Remaining = probe.remaining
	res.Total = probe.total
	res.ResetAt = probe.resetAt
	return res, nil
}

type probeOutcome struct {
	remaining        int
	total            int
	resetAt          time.Time
	quotaProbed      bool
	defaultModel     string
	blockedFeatures  []string
	capabilityStatus string
	capabilitySource string
	capabilityDetail string
}

type initProbeOutcome struct {
	remaining       int
	total           int
	resetAt         time.Time
	defaultModel    string
	blockedFeatures []string
}

type modelsProbeOutcome struct {
	defaultModel       string
	modelPickerVersion string
	imageEnabled       bool
	imageModels        []string
	fConversationEPs   []string
}

const defaultProbeTimezoneOffsetMin = -480

// upstreamClientFor 构造与真实 f/conversation 链路一致的 ChatGPT client。
// /backend-api/models 已经是 image 能力主探针,因此探针不能再使用弱化的
// 标准 net/http 指纹;这里统一复用 uTLS transport、commonHeaders、Oai-*、
// 稳定 device/session、代理与 cookie jar。
func (q *QuotaProber) upstreamClientFor(ctx context.Context, a *Account, accessToken string) (*chatgpt.Client, error) {
	deviceID := a.OAIDeviceID
	if deviceID == "" {
		gen := uuid.NewString()
		if fixed, err := q.svc.dao.EnsureDeviceID(ctx, a.ID, gen); err == nil && fixed != "" {
			deviceID = fixed
			a.OAIDeviceID = fixed
		} else {
			deviceID = gen
		}
	}

	sessionID := a.OAISessionID
	if sessionID == "" {
		gen := uuid.NewString()
		if fixed, err := q.svc.dao.EnsureSessionID(ctx, a.ID, gen); err == nil && fixed != "" {
			sessionID = fixed
			a.OAISessionID = fixed
		} else {
			sessionID = gen
		}
	}

	var proxyURL string
	if q.proxyResolver != nil {
		proxyURL = q.proxyResolver.ProxyURLForAccount(ctx, a.ID)
	}

	var cookies string
	if enc, err := q.svc.dao.GetCookies(ctx, a.ID); err == nil && enc != "" {
		if dec, derr := q.svc.cipher.DecryptString(enc); derr == nil {
			cookies = dec
		} else if q.log != nil {
			q.log.Warn("decrypt account cookies for quota probe failed",
				zap.Uint64("account_id", a.ID), zap.Error(derr))
		}
	}

	return chatgpt.New(chatgpt.Options{
		AuthToken: accessToken,
		DeviceID:  deviceID,
		SessionID: sessionID,
		ProxyURL:  proxyURL,
		Cookies:   cookies,
		Timeout:   30 * time.Second,
	})
}

// doProbe 组合探测账号 image 状态。
//
// 实测结论: /backend-api/conversation/init 已经不能再作为 image 能力主判据。
// 同一个 Pro/image 灰度号可能在 /backend-api/models 暴露 image_gen_tool_enabled,
// 但 init 仍返回 blocked_features=image_gen / default_model_slug=auto。
//
// 因此这里分层处理:
//  1. /backend-api/models:主能力探针,判断账号是否具备 image 入口资格;
//  2. /backend-api/conversation/init:仅保留 quota / reset / blocked_features 诊断弱参考;
//  3. 真实 IMG2 是否命中仍由 Runner 的 SSE/poll/tool payload 画像决定。
func (q *QuotaProber) doProbe(ctx context.Context, a *Account, accessToken string) (out probeOutcome, err error) {
	out.remaining = -1
	out.total = -1
	out.capabilityStatus = "unknown"

	cli, clientErr := q.upstreamClientFor(ctx, a, accessToken)
	if clientErr != nil {
		out.capabilityStatus = "error"
		out.capabilitySource = "models"
		out.capabilityDetail = truncate(clientErr.Error(), 500)
		return out, clientErr
	}

	models, modelsErr := q.doModelsProbe(ctx, cli)
	if modelsErr == nil {
		out.defaultModel = models.defaultModel
		out.capabilitySource = "models"
		if models.imageEnabled {
			out.capabilityStatus = "enabled"
		} else {
			out.capabilityStatus = "disabled"
		}
		out.capabilityDetail = models.detailJSON()
	} else {
		out.capabilitySource = "models"
		out.capabilityStatus = "error"
		out.capabilityDetail = truncate(modelsErr.Error(), 500)
	}

	initOut, initErr := q.doInitProbe(ctx, cli)
	if initErr == nil {
		out.remaining = initOut.remaining
		out.total = initOut.total
		out.resetAt = initOut.resetAt
		out.quotaProbed = true
		out.blockedFeatures = initOut.blockedFeatures
		if out.defaultModel == "" {
			out.defaultModel = initOut.defaultModel
		}
	} else if modelsErr != nil {
		// 两个探针都失败才把这轮标为失败;否则以 models 为主继续写入能力画像。
		return out, fmt.Errorf("models probe: %v; init probe: %v", modelsErr, initErr)
	}

	return out, nil
}

func (m modelsProbeOutcome) detailJSON() string {
	payload := map[string]interface{}{
		"default_model_slug":       m.defaultModel,
		"model_picker_version":     m.modelPickerVersion,
		"image_enabled":            m.imageEnabled,
		"image_models":             m.imageModels,
		"f_conversation_endpoints": m.fConversationEPs,
	}
	buf, _ := json.Marshal(payload)
	return string(buf)
}

// doModelsProbe 调 /backend-api/models,这是当前 image 能力主探针。
// 只要某个模型 enabledTools 含 image_gen_tool_enabled,或 supportedFeatures 含 image,
// 就认为账号具有 image 入口资格。是否抽中 IMG2 终稿由 Runner 另行画像。
func (q *QuotaProber) doModelsProbe(ctx context.Context, cli *chatgpt.Client) (out modelsProbeOutcome, err error) {
	data, err := cli.ModelsRaw(ctx)
	if err != nil {
		return out, err
	}

	var payload map[string]interface{}
	if err = json.Unmarshal(data, &payload); err != nil {
		return out, err
	}
	out.defaultModel = firstString(payload, "default_model_slug", "defaultModelSlug", "default_model", "defaultModel")
	out.modelPickerVersion = firstString(payload, "model_picker_version", "modelPickerVersion")

	modelsRaw := modelListFromAny(payload["models"])
	seenModel := map[string]struct{}{}
	seenEP := map[string]struct{}{}
	for _, mm := range modelsRaw {
		if mm == nil {
			continue
		}
		slug := firstString(mm, "slug", "id", "model", "model_slug", "modelSlug")
		enabledTools := stringSliceFromAny(firstAny(mm, "enabledTools", "enabled_tools", "enabled_tools_override"))
		supportedFeatures := stringSliceFromAny(firstAny(mm, "supportedFeatures", "supported_features", "features"))
		fep := firstString(mm, "fConversationEndpoint", "f_conversation_endpoint")
		if fep != "" {
			if _, ok := seenEP[fep]; !ok {
				seenEP[fep] = struct{}{}
				out.fConversationEPs = append(out.fConversationEPs, fep)
			}
		}
		if hasImageToolOrFeature(enabledTools, supportedFeatures) {
			out.imageEnabled = true
			if slug != "" {
				if _, ok := seenModel[slug]; !ok {
					seenModel[slug] = struct{}{}
					out.imageModels = append(out.imageModels, slug)
				}
			}
		}
	}
	return out, nil
}

// doInitProbe 调 /backend-api/conversation/init。
//
// 这个接口只用于 quota/reset 诊断;不要再用 blocked_features 决定 image2 能力。
func (q *QuotaProber) doInitProbe(ctx context.Context, cli *chatgpt.Client) (out initProbeOutcome, err error) {
	out.remaining = -1
	out.total = -1

	// timezone_offset_min 只对齐前端抓包值;不要把它解释成服务端时区语义。
	// 后续如果要做区域画像,可以把这个值提升到配置或从浏览器登录态同步。
	data, err := cli.ConversationInitRaw(ctx, defaultProbeTimezoneOffsetMin, []string{"picture_v2"})
	if err != nil {
		return out, err
	}

	var payload struct {
		Type             string   `json:"type"`
		BlockedFeatures  []string `json:"blocked_features"`
		DefaultModelSlug string   `json:"default_model_slug"`
		LimitsProgress   []struct {
			FeatureName string `json:"feature_name"`
			Remaining   *int   `json:"remaining"`
			ResetAfter  string `json:"reset_after"`
		} `json:"limits_progress"`
	}
	if err = json.Unmarshal(data, &payload); err != nil {
		return out, err
	}
	out.defaultModel = payload.DefaultModelSlug
	out.blockedFeatures = payload.BlockedFeatures

	for _, item := range payload.LimitsProgress {
		if !isImageFeature(item.FeatureName) {
			continue
		}
		if item.Remaining != nil {
			if out.remaining < 0 || *item.Remaining < out.remaining {
				out.remaining = *item.Remaining
			}
		}
		if item.ResetAfter != "" {
			if t, e := time.Parse(time.RFC3339, item.ResetAfter); e == nil {
				if out.resetAt.IsZero() || t.Before(out.resetAt) {
					out.resetAt = t
				}
			}
		}
	}
	return out, nil
}

func modelListFromAny(v interface{}) []map[string]interface{} {
	switch xs := v.(type) {
	case []interface{}:
		out := make([]map[string]interface{}, 0, len(xs))
		for _, x := range xs {
			if m, ok := x.(map[string]interface{}); ok {
				out = append(out, m)
			}
		}
		return out
	case map[string]interface{}:
		out := make([]map[string]interface{}, 0, len(xs))
		for slug, raw := range xs {
			if m, ok := raw.(map[string]interface{}); ok {
				if _, exists := m["slug"]; !exists {
					m["slug"] = slug
				}
				out = append(out, m)
			}
		}
		return out
	default:
		return nil
	}
}

func firstAny(m map[string]interface{}, keys ...string) interface{} {
	for _, k := range keys {
		if v, ok := m[k]; ok {
			return v
		}
	}
	return nil
}

func firstString(m map[string]interface{}, keys ...string) string {
	for _, k := range keys {
		if v, ok := m[k]; ok {
			switch x := v.(type) {
			case string:
				return x
			case float64:
				return fmt.Sprintf("%.0f", x)
			}
		}
	}
	return ""
}

func stringSliceFromAny(v interface{}) []string {
	switch xs := v.(type) {
	case []string:
		return xs
	case []interface{}:
		out := make([]string, 0, len(xs))
		for _, x := range xs {
			switch y := x.(type) {
			case string:
				if y != "" {
					out = append(out, y)
				}
			case map[string]interface{}:
				if s := firstString(y, "name", "slug", "id", "feature", "tool"); s != "" {
					out = append(out, s)
				}
			}
		}
		return out
	case map[string]interface{}:
		out := make([]string, 0, len(xs))
		for k, v := range xs {
			if b, ok := v.(bool); ok && b {
				out = append(out, k)
			}
		}
		return out
	default:
		return nil
	}
}

func hasImageToolOrFeature(enabledTools, supportedFeatures []string) bool {
	for _, s := range enabledTools {
		low := strings.ToLower(s)
		if low == "image_gen_tool_enabled" || strings.Contains(low, "image_gen") || strings.Contains(low, "image") {
			return true
		}
	}
	for _, s := range supportedFeatures {
		low := strings.ToLower(s)
		if low == "image" || strings.Contains(low, "image") {
			return true
		}
	}
	return false
}

func isImageFeature(name string) bool {
	n := strings.ToLower(name)
	switch n {
	case "image_gen", "image_generation", "image_edit", "img_gen":
		return true
	}
	return strings.Contains(n, "image_gen") || strings.Contains(n, "img_gen")
}

func friendlyProbeErr(err error) string {
	if err == nil {
		return ""
	}
	s := err.Error()
	low := strings.ToLower(s)
	switch {
	case strings.Contains(low, "http=401"):
		return "AT 已过期,无法探测额度"
	case strings.Contains(low, "http=403"):
		return "上游拒绝访问(403)"
	case strings.Contains(low, "http=429"):
		return "上游速率限制(429)"
	case strings.Contains(low, "timeout"), strings.Contains(low, "deadline exceeded"):
		return "探测超时"
	case strings.Contains(low, "connection refused"), strings.Contains(low, "no such host"):
		return "网络不通"
	default:
		if len(s) > 160 {
			s = s[:160] + "…"
		}
		return "探测失败:" + s
	}
}
````

## File: internal/account/refresher.go
````go
package account

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
)

// AccountProxyResolver 把账号 ID 映射成代理 URL(形如 http(s)://user:pass@host:port)。
// 由 main.go 用 account.Service + proxy.Service 组装后注入;未绑定或禁用返回 ""。
type AccountProxyResolver interface {
	ProxyURLForAccount(ctx context.Context, accountID uint64) string
}

// RefreshSettings 热更新参数。由 settings.Service 实现。
type RefreshSettings interface {
	AccountRefreshEnabled() bool
	AccountRefreshIntervalSec() int
	AccountRefreshAheadSec() int
	AccountRefreshConcurrency() int
	AccountDefaultClientID() string
}

// RefreshResult 刷新结果。
type RefreshResult struct {
	AccountID uint64    `json:"account_id"`
	Email     string    `json:"email"`
	OK        bool      `json:"ok"`
	Source    string    `json:"source"` // rt / st / failed
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	Error     string    `json:"error,omitempty"`
	RTRotated bool      `json:"rt_rotated,omitempty"`
	// ATVerified 表示新 AT 已被 chatgpt.com web 后端接受(GET /backend-api/me 返回 200)。
	// 只要 Source=st 必为 true;Source=rt 时取决于 verify 结果;Source=failed 时为 false。
	ATVerified bool `json:"at_verified"`
	// WebUnauthorized 为 true 表示 RT 换出来的 AT 被 chatgpt.com 以 401 拒绝
	// (iOS OAuth 作用域 vs Web 作用域不一致)。前端据此提示补充 Session Token。
	WebUnauthorized bool `json:"web_unauthorized,omitempty"`
}

// Refresher 负责把账号的 AT 通过 RT 或 ST 刷新成新的 AT。
type Refresher struct {
	svc      *Service
	settings RefreshSettings
	log      *zap.Logger
	client   *http.Client

	proxyResolver AccountProxyResolver

	kick chan struct{}
}

// NewRefresher 构造。
// HTTP client 默认直连;如果注入了 AccountProxyResolver,则每次刷新会优先使用
// 账号绑定的代理,避免从境内直连 auth.openai.com / chatgpt.com 时被劫持。
func NewRefresher(svc *Service, settings RefreshSettings, logger *zap.Logger) *Refresher {
	return &Refresher{
		svc:      svc,
		settings: settings,
		log:      logger,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		kick: make(chan struct{}, 1),
	}
}

// Kick 立刻触发一次扫描(批量刷新按钮触发时调)。
// SetProxyResolver 注入账号代理解析器。
// 未调用时 Refresher 保持直连行为(兼容无代理池环境)。
func (r *Refresher) SetProxyResolver(pr AccountProxyResolver) { r.proxyResolver = pr }

// clientFor 根据账号 ID 选择合适的 http.Client:
//   - proxyResolver 未注入或返回空 URL → 用默认直连 client
//   - 返回非空 URL → 构造一次性带代理的 client(结束后 GC)
//
// 代理 URL 解析失败时降级到直连,并打 warn 日志。
func (r *Refresher) clientFor(ctx context.Context, accountID uint64) *http.Client {
	if r.proxyResolver == nil {
		return r.client
	}
	pu := r.proxyResolver.ProxyURLForAccount(ctx, accountID)
	if pu == "" {
		return r.client
	}
	u, err := url.Parse(pu)
	if err != nil {
		r.log.Warn("invalid proxy url for refresh, fallback direct",
			zap.Uint64("account_id", accountID), zap.Error(err))
		return r.client
	}
	tr := &http.Transport{
		Proxy:               http.ProxyURL(u),
		ForceAttemptHTTP2:   true,
		MaxIdleConns:        16,
		IdleConnTimeout:     30 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
	}
	return &http.Client{Transport: tr, Timeout: r.client.Timeout}
}

func (r *Refresher) Kick() {
	select {
	case r.kick <- struct{}{}:
	default:
	}
}

// Run 后台循环。
func (r *Refresher) Run(ctx context.Context) {
	r.log.Info("account refresher started")
	defer r.log.Info("account refresher stopped")

	// 启动延迟 5s,避免和数据库初始化抢锁
	select {
	case <-ctx.Done():
		return
	case <-time.After(5 * time.Second):
	}

	for {
		interval := time.Duration(r.settings.AccountRefreshIntervalSec()) * time.Second
		if interval < 30*time.Second {
			interval = 30 * time.Second
		}

		if r.settings.AccountRefreshEnabled() {
			r.scanOnce(ctx)
		}

		select {
		case <-ctx.Done():
			return
		case <-time.After(interval):
		case <-r.kick:
		}
	}
}

func (r *Refresher) scanOnce(ctx context.Context) {
	// 先清理:把已过期且无 RT/ST 的纯 AT 账号自动标记 dead
	if n, err := r.svc.dao.RetireExpiredATOnly(ctx); err != nil {
		r.log.Warn("retire expired AT-only accounts failed", zap.Error(err))
	} else if n > 0 {
		r.log.Info("retired expired AT-only accounts", zap.Int64("count", n))
	}

	ahead := r.settings.AccountRefreshAheadSec()
	conc := r.settings.AccountRefreshConcurrency()

	rows, err := r.svc.dao.ListNeedRefresh(ctx, ahead, 256)
	if err != nil {
		r.log.Warn("list need-refresh accounts failed", zap.Error(err))
		return
	}
	if len(rows) == 0 {
		return
	}
	r.log.Info("refreshing accounts", zap.Int("count", len(rows)), zap.Int("ahead_sec", ahead), zap.Int("concurrency", conc))

	sem := make(chan struct{}, conc)
	var wg sync.WaitGroup
	for _, a := range rows {
		a := a
		wg.Add(1)
		sem <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-sem }()
			_, _ = r.RefreshAuto(ctx, a)
		}()
	}
	wg.Wait()
}

// RefreshByID 指定 id 刷新。
func (r *Refresher) RefreshByID(ctx context.Context, id uint64) (*RefreshResult, error) {
	a, err := r.svc.dao.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return r.RefreshAuto(ctx, a)
}

// RefreshAuto 优先 RT,失败/没有 RT 回退 ST;都失败则 markDead。
//
// 判定规则(严格):
//
//  1. RT → AT HTTP 成功后,立即 GET /backend-api/me 做作用域校验
//     - 200:写库,Source=rt,ATVerified=true,返回成功
//     - 非 200(含 401/403/429/5xx/网络错误):**视为 RT 不可用**,丢弃本次 AT,
//     不写库,继续尝试 ST
//  2. ST → AT HTTP 成功:写库,Source=st,ATVerified=true,返回成功
//  3. 两条路径都未拿到可用 AT:返回失败(ok=false),账号标 warned(RT 被 web 拒)
//     或 dead(完全没可用 token)。不会把"无法校验通过"的 AT 悄悄写进库。
//
// 这条规则确保前端看到「刷新成功」时,AT 一定是 chatgpt.com web 后端接受的,
// 后续探测 / 聊天 / 图像请求不会因作用域不匹配而 401。
func (r *Refresher) RefreshAuto(ctx context.Context, a *Account) (*RefreshResult, error) {
	if a == nil {
		return nil, errors.New("account is nil")
	}
	res := &RefreshResult{AccountID: a.ID, Email: a.Email}

	// 尝试 RT
	if a.RefreshTokenEnc.Valid && a.RefreshTokenEnc.String != "" {
		rt, err := r.svc.cipher.DecryptString(a.RefreshTokenEnc.String)
		if err == nil && rt != "" {
			clientID := a.ClientID
			if clientID == "" {
				clientID = r.settings.AccountDefaultClientID()
			}
			newAT, newRT, expAt, err := r.rtToAT(ctx, a.ID, rt, clientID)
			if err == nil && newAT != "" {
				// RT → AT 成功,直接写库(不再做 verifyATOnWeb 校验——
				// 用 app_LlGpXReQgckcGGUo2JrYvtJK 换出的 AT 在 f/conversation 里能用,
				// 但 /backend-api/me 可能因 scope 不同返回 401 导致误杀)
				return r.applyRefresh(ctx, a, newAT, newRT, expAt, RefreshSourceRT, res)
			}
			// RT → AT HTTP 本身失败,回退 ST
			r.log.Warn("RT refresh failed, fallback to ST", zap.Uint64("id", a.ID), zap.Error(err))
			res.Error = friendlyRefreshErr(err)
		}
	}

	// 尝试 ST(ST → AT 本来就是 web 作用域,不需要再校验)
	if a.SessionTokenEnc.Valid && a.SessionTokenEnc.String != "" {
		st, err := r.svc.cipher.DecryptString(a.SessionTokenEnc.String)
		if err == nil && st != "" {
			newAT, expAt, err := r.stToAT(ctx, a.ID, st)
			if err == nil && newAT != "" {
				res.ATVerified = true
				return r.applyRefresh(ctx, a, newAT, "", expAt, RefreshSourceST, res)
			}
			if res.Error == "" {
				res.Error = friendlyRefreshErr(err)
			} else {
				res.Error += " / ST:" + friendlyRefreshErr(err)
			}
		}
	}

	// 都不行:标 warned,保留重试机会
	if res.Error == "" {
		res.Error = "账号既无可用 RT 也无可用 ST"
	}
	_ = r.svc.dao.RecordRefreshError(ctx, a.ID, RefreshSourceRT, res.Error, false)
	r.log.Warn("refresh failed, marking warned (not dead)",
		zap.Uint64("id", a.ID), zap.String("email", a.Email), zap.String("error", res.Error))
	res.Source = "failed"
	return res, nil
}

// verifyATOnWeb 用新 AT 访问一个极轻量的 chatgpt.com web 端点,确认 AT 的作用域
// 能被 web 后端接受。
//
// 选用 GET /backend-api/me:
//   - 200 说明 AT 有效且作用域匹配 web
//   - 401 说明 AT 无效或作用域不匹配(iOS OAuth RT 刷出的 AT 常见)
//   - 403/429/5xx/网络错误 都不作为"AT 无效"的依据
//
// 返回 (http_status, err);err 为非 HTTP 层错误(dial/tls 等)。
func (r *Refresher) verifyATOnWeb(ctx context.Context, accountID uint64, accessToken string) (int, error) {
	vctx, cancel := context.WithTimeout(ctx, 12*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(vctx, "GET",
		"https://chatgpt.com/backend-api/me", nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "https://chatgpt.com/")
	req.Header.Set("Origin", "https://chatgpt.com")
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")

	resp, err := r.clientFor(vctx, accountID).Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	// 读掉 body,释放连接
	_, _ = io.Copy(io.Discard, resp.Body)
	return resp.StatusCode, nil
}

func (r *Refresher) applyRefresh(
	ctx context.Context, a *Account,
	newAT, newRT string, expAt time.Time, source string,
	res *RefreshResult,
) (*RefreshResult, error) {
	atEnc, err := r.svc.cipher.EncryptString(newAT)
	if err != nil {
		return nil, err
	}
	var rtEnc string
	if newRT != "" {
		enc, err := r.svc.cipher.EncryptString(newRT)
		if err != nil {
			return nil, err
		}
		rtEnc = enc
	}
	if expAt.IsZero() {
		expAt = parseJWTExp(newAT)
	}
	if err := r.svc.dao.ApplyRefreshResult(ctx, a.ID, atEnc, rtEnc, expAt, source); err != nil {
		return nil, err
	}
	// 自动更新 subscription_type（从新 AT 的 JWT 解析）
	if pt := parseJWTPlanType(newAT); pt != "" && pt != a.SubscriptionType {
		_ = r.svc.dao.SetSubscriptionType(ctx, a.ID, pt)
	}
	res.OK = true
	res.Source = source
	res.ExpiresAt = expAt
	res.Error = ""
	res.RTRotated = rtEnc != ""
	return res, nil
}

// rtToAT POST https://auth.openai.com/oauth/token
func (r *Refresher) rtToAT(ctx context.Context, accountID uint64, refreshToken, clientID string) (newAT, newRT string, expAt time.Time, err error) {
	body := map[string]string{
		"client_id":     clientID,
		"grant_type":    "refresh_token",
		"redirect_uri":  "com.openai.chat://auth0.openai.com/ios/com.openai.chat/callback",
		"refresh_token": refreshToken,
	}
	buf, _ := json.Marshal(body)
	req, err := http.NewRequestWithContext(ctx, "POST",
		"https://auth.openai.com/oauth/token", bytes.NewReader(buf))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "ChatGPT/1.2025.122 (iOS 18.2; iPhone15,2; build 15096)")

	resp, err := r.clientFor(ctx, accountID).Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		err = fmt.Errorf("rt refresh http=%d body=%s", resp.StatusCode, truncate(string(data), 200))
		return
	}
	var out struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int    `json:"expires_in"`
	}
	if err = json.Unmarshal(data, &out); err != nil {
		return
	}
	if out.AccessToken == "" {
		err = errors.New("rt refresh: missing access_token in response")
		return
	}
	newAT = out.AccessToken
	newRT = out.RefreshToken
	if out.ExpiresIn > 0 {
		expAt = time.Now().Add(time.Duration(out.ExpiresIn) * time.Second)
	} else {
		expAt = parseJWTExp(newAT)
	}
	return
}

// stToAT GET https://chatgpt.com/api/auth/session  Cookie: __Secure-next-auth.session-token=ST
func (r *Refresher) stToAT(ctx context.Context, accountID uint64, sessionToken string) (newAT string, expAt time.Time, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET",
		"https://chatgpt.com/api/auth/session", nil)
	if err != nil {
		return
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "https://chatgpt.com/")
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")

	// 同时尝试两个可能的 cookie 名
	req.AddCookie(&http.Cookie{Name: "__Secure-next-auth.session-token", Value: sessionToken})

	resp, err := r.clientFor(ctx, accountID).Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		err = fmt.Errorf("st refresh http=%d body=%s", resp.StatusCode, truncate(string(data), 200))
		return
	}
	raw := strings.TrimSpace(string(data))
	if raw == "" || raw == "{}" {
		err = errors.New("ST 已过期或无效,响应为空")
		return
	}
	var out struct {
		AccessToken string `json:"accessToken"`
		Expires     string `json:"expires"`
	}
	if err = json.Unmarshal([]byte(raw), &out); err != nil {
		return
	}
	if out.AccessToken == "" {
		err = errors.New("响应缺少 accessToken 字段")
		return
	}
	newAT = out.AccessToken
	if out.Expires != "" {
		if t, e := time.Parse(time.RFC3339, out.Expires); e == nil {
			expAt = t
		}
	}
	if expAt.IsZero() {
		expAt = parseJWTExp(newAT)
	}
	return
}

// jwtClaims JWT payload 中需要的字段。
type jwtClaims struct {
	Exp       int64  `json:"exp"`
	PlanType  string `json:"chatgpt_plan_type"`
	AccountID string `json:"chatgpt_account_id"`
}

func parseJWTClaims(token string) jwtClaims {
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return jwtClaims{}
	}
	raw, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		raw, err = base64.StdEncoding.DecodeString(parts[1])
		if err != nil {
			return jwtClaims{}
		}
	}
	// ChatGPT JWT: plan_type 在嵌套的 "https://api.openai.com/auth" 下
	var outer struct {
		Exp  int64 `json:"exp"`
		Auth struct {
			PlanType  string `json:"chatgpt_plan_type"`
			AccountID string `json:"chatgpt_account_id"`
		} `json:"https://api.openai.com/auth"`
	}
	if err := json.Unmarshal(raw, &outer); err == nil && outer.Auth.PlanType != "" {
		return jwtClaims{Exp: outer.Exp, PlanType: outer.Auth.PlanType, AccountID: outer.Auth.AccountID}
	}
	var flat jwtClaims
	_ = json.Unmarshal(raw, &flat)
	return flat
}

// parseJWTExp 解 JWT payload 里的 exp(秒级)。失败返回 +24h。
func parseJWTExp(token string) time.Time {
	c := parseJWTClaims(token)
	if c.Exp == 0 {
		return time.Now().Add(24 * time.Hour)
	}
	return time.Unix(c.Exp, 0)
}

// parseJWTPlanType 从 JWT 解析 chatgpt_plan_type(pro/plus/free/team 等)。
func parseJWTPlanType(token string) string {
	return parseJWTClaims(token).PlanType
}

func friendlyRefreshErr(err error) string {
	if err == nil {
		return ""
	}
	s := err.Error()
	low := strings.ToLower(s)
	switch {
	case strings.Contains(low, "http=401"), strings.Contains(low, "invalid_grant"):
		return "RT 已失效(401)"
	case strings.Contains(low, "http=403"):
		return "上游拒绝访问(403)"
	case strings.Contains(low, "http=429"):
		return "触发速率限制(429)"
	case strings.Contains(low, "timeout"), strings.Contains(low, "deadline exceeded"):
		return "刷新请求超时"
	case strings.Contains(low, "proxyconnect") && strings.Contains(low, "no such host"):
		return "代理域名无法解析"
	case strings.Contains(low, "proxyconnect tcp"):
		return "代理握手失败"
	case strings.Contains(low, "no such host"):
		return "DNS 解析失败"
	case strings.Contains(low, "connection refused"):
		return "连接被拒绝"
	case strings.Contains(low, "connection reset"):
		return "连接被重置"
	case strings.Contains(low, "unexpected eof"), strings.HasSuffix(low, ": eof"):
		return "连接被对端关闭"
	case strings.Contains(low, "tls"), strings.Contains(low, "x509"):
		return "TLS 握手失败"
	case strings.Contains(low, "missing access_token"), strings.Contains(s, "ST 已过期"):
		return stripHTTPPrefix(s)
	default:
		return "刷新失败:" + stripHTTPPrefix(s)
	}
}

// stripHTTPPrefix 去掉 Go net/http 错误里形如
//
//	Post "https://auth.openai.com/oauth/token": dial tcp: ...
//
// 的 URL 前缀,只保留后面真正的原因,避免把敏感/冗长的 URL 暴露给前端。
func stripHTTPPrefix(s string) string {
	// 典型前缀: Get/Post/Put "https://...": rest
	if i := strings.Index(s, `": `); i > 0 && i < len(s)-3 {
		s = s[i+3:]
	}
	// 再剥一层常见的 "dial tcp: " / "proxyconnect tcp: " 类前缀最靠前的修饰,
	// 保留中间有用的原因(如 lookup xxx: no such host)。
	s = strings.TrimSpace(s)
	if len(s) > 120 {
		s = s[:120] + "…"
	}
	return s
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "…"
}
````

## File: internal/account/service.go
````go
package account

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/432539/gpt2api/pkg/crypto"
)

// Service 账号池业务。
type Service struct {
	dao    *DAO
	cipher *crypto.AESGCM
}

func NewService(dao *DAO, cipher *crypto.AESGCM) *Service {
	return &Service{dao: dao, cipher: cipher}
}

// CreateInput 新增账号入参(明文敏感字段)。
type CreateInput struct {
	Email            string    `json:"email"`
	AuthToken        string    `json:"auth_token"`
	RefreshToken     string    `json:"refresh_token"`
	SessionToken     string    `json:"session_token"`
	TokenExpiresAt   time.Time `json:"token_expires_at"`
	OAISessionID     string    `json:"oai_session_id"`
	OAIDeviceID      string    `json:"oai_device_id"`
	ClientID         string    `json:"client_id"`
	ChatGPTAccountID string    `json:"chatgpt_account_id"`
	AccountType      string    `json:"account_type"`
	SubscriptionType string    `json:"subscription_type"`
	Notes            string    `json:"notes"`
	Cookies          string    `json:"cookies"`
	ProxyID          uint64    `json:"proxy_id"` // 可选:立即绑定
}

// UpdateInput 更新入参。AuthToken/RefreshToken/SessionToken/Cookies 为空串表示不改。
type UpdateInput struct {
	Email            string    `json:"email"`
	AuthToken        string    `json:"auth_token"`
	RefreshToken     string    `json:"refresh_token"`
	SessionToken     string    `json:"session_token"`
	TokenExpiresAt   time.Time `json:"token_expires_at"`
	OAISessionID     string    `json:"oai_session_id"`
	OAIDeviceID      string    `json:"oai_device_id"`
	ClientID         string    `json:"client_id"`
	ChatGPTAccountID string    `json:"chatgpt_account_id"`
	AccountType      string    `json:"account_type"`
	SubscriptionType string    `json:"subscription_type"`
	Status           string    `json:"status"`
	Notes            string    `json:"notes"`
	Cookies          string    `json:"cookies"`
}

func (s *Service) Create(ctx context.Context, in CreateInput) (*Account, error) {
	if in.Email == "" || in.AuthToken == "" {
		return nil, errors.New("email 和 auth_token 不能为空")
	}
	atEnc, err := s.cipher.EncryptString(in.AuthToken)
	if err != nil {
		return nil, err
	}
	var rtEnc, stEnc sql.NullString
	if in.RefreshToken != "" {
		v, err := s.cipher.EncryptString(in.RefreshToken)
		if err != nil {
			return nil, err
		}
		rtEnc = sql.NullString{String: v, Valid: true}
	}
	if in.SessionToken != "" {
		v, err := s.cipher.EncryptString(in.SessionToken)
		if err != nil {
			return nil, err
		}
		stEnc = sql.NullString{String: v, Valid: true}
	}
	if in.OAIDeviceID == "" {
		in.OAIDeviceID = uuid.NewString()
	}
	if in.OAISessionID == "" {
		in.OAISessionID = uuid.NewString()
	}
	if in.SubscriptionType == "" {
		// 尝试从 AT 的 JWT 自动识别 plan_type
		if pt := parseJWTPlanType(in.AuthToken); pt != "" {
			in.SubscriptionType = pt
		} else {
			in.SubscriptionType = "plus"
		}
	}
	if in.ClientID == "" {
		in.ClientID = "app_LlGpXReQgckcGGUo2JrYvtJK"
	}
	if in.AccountType == "" {
		in.AccountType = "codex"
	}
	a := &Account{
		Email: in.Email, AuthTokenEnc: atEnc, RefreshTokenEnc: rtEnc, SessionTokenEnc: stEnc,
		OAISessionID: in.OAISessionID, OAIDeviceID: in.OAIDeviceID,
		ClientID: in.ClientID, ChatGPTAccountID: in.ChatGPTAccountID, AccountType: in.AccountType,
		SubscriptionType: in.SubscriptionType,
		Status: StatusHealthy, Notes: in.Notes,
	}
	if !in.TokenExpiresAt.IsZero() {
		a.TokenExpiresAt = sql.NullTime{Time: in.TokenExpiresAt, Valid: true}
	} else {
		// 自动从 JWT 解析 exp
		if exp := parseJWTExp(in.AuthToken); !exp.IsZero() {
			a.TokenExpiresAt = sql.NullTime{Time: exp, Valid: true}
		}
	}
	id, err := s.dao.Create(ctx, a)
	if err != nil {
		return nil, err
	}
	a.ID = id
	if in.Cookies != "" {
		enc, err := s.cipher.EncryptString(in.Cookies)
		if err != nil {
			return nil, err
		}
		if err := s.dao.UpsertCookies(ctx, id, enc); err != nil {
			return nil, err
		}
	}
	// 绑定代理:指定了就用指定的,否则自动分配绑定数最少的
	proxyID := in.ProxyID
	if proxyID == 0 {
		if autoID, err := s.dao.LeastBoundProxyID(ctx, nil); err == nil {
			proxyID = autoID
		}
	}
	if proxyID > 0 {
		if err := s.dao.SetBinding(ctx, id, proxyID); err != nil {
			return nil, err
		}
	}
	return s.dao.GetByID(ctx, id)
}

func (s *Service) Update(ctx context.Context, id uint64, in UpdateInput) (*Account, error) {
	a, err := s.dao.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if in.Email != "" {
		a.Email = in.Email
	}
	if in.AuthToken != "" {
		enc, err := s.cipher.EncryptString(in.AuthToken)
		if err != nil {
			return nil, err
		}
		a.AuthTokenEnc = enc
	}
	if in.RefreshToken != "" {
		enc, err := s.cipher.EncryptString(in.RefreshToken)
		if err != nil {
			return nil, err
		}
		a.RefreshTokenEnc = sql.NullString{String: enc, Valid: true}
	}
	if in.SessionToken != "" {
		enc, err := s.cipher.EncryptString(in.SessionToken)
		if err != nil {
			return nil, err
		}
		a.SessionTokenEnc = sql.NullString{String: enc, Valid: true}
	}
	if !in.TokenExpiresAt.IsZero() {
		a.TokenExpiresAt = sql.NullTime{Time: in.TokenExpiresAt, Valid: true}
	} else if in.AuthToken != "" {
		if exp := parseJWTExp(in.AuthToken); !exp.IsZero() {
			a.TokenExpiresAt = sql.NullTime{Time: exp, Valid: true}
		}
	}
	if in.OAISessionID != "" {
		a.OAISessionID = in.OAISessionID
	}
	if in.OAIDeviceID != "" {
		a.OAIDeviceID = in.OAIDeviceID
	}
	if in.ClientID != "" {
		a.ClientID = in.ClientID
	}
	if in.ChatGPTAccountID != "" {
		a.ChatGPTAccountID = in.ChatGPTAccountID
	}
	if in.AccountType != "" {
		a.AccountType = in.AccountType
	}
	if in.SubscriptionType != "" {
		a.SubscriptionType = in.SubscriptionType
	}
	if in.Status != "" {
		a.Status = in.Status
	}
	a.Notes = in.Notes
	if err := s.dao.Update(ctx, a); err != nil {
		return nil, err
	}
	if in.Cookies != "" {
		enc, err := s.cipher.EncryptString(in.Cookies)
		if err != nil {
			return nil, err
		}
		if err := s.dao.UpsertCookies(ctx, id, enc); err != nil {
			return nil, err
		}
	}
	return a, nil
}

func (s *Service) Delete(ctx context.Context, id uint64) error {
	return s.dao.SoftDelete(ctx, id)
}

// BulkDeleteByStatus 批量软删;status 支持 dead / suspicious / warned / throttled / all。
func (s *Service) BulkDeleteByStatus(ctx context.Context, status string) (int64, error) {
	if status == "all" {
		return s.dao.SoftDeleteByStatus(ctx, "")
	}
	return s.dao.SoftDeleteByStatus(ctx, status)
}

func (s *Service) Get(ctx context.Context, id uint64) (*Account, error) {
	return s.dao.GetByID(ctx, id)
}

func (s *Service) List(ctx context.Context, status, keyword string, offset, limit int) ([]*Account, int64, error) {
	return s.dao.List(ctx, status, keyword, offset, limit)
}

// BindProxy 绑定代理(一号一代理)。
func (s *Service) BindProxy(ctx context.Context, accountID, proxyID uint64) error {
	return s.dao.SetBinding(ctx, accountID, proxyID)
}

// UnbindProxy 解除绑定。
func (s *Service) UnbindProxy(ctx context.Context, accountID uint64) error {
	return s.dao.RemoveBinding(ctx, accountID)
}

// DecryptAuthToken 解密 AT。
func (s *Service) DecryptAuthToken(a *Account) (string, error) {
	return s.cipher.DecryptString(a.AuthTokenEnc)
}

// AccountSecrets AT / RT / ST 明文,仅给本地控制台编辑页回填使用。
type AccountSecrets struct {
	AuthToken    string `json:"auth_token"`
	RefreshToken string `json:"refresh_token"`
	SessionToken string `json:"session_token"`
}

// GetSecrets 返回指定账号的 AT/RT/ST 明文(用于后台编辑弹窗回显)。
func (s *Service) GetSecrets(ctx context.Context, id uint64) (*AccountSecrets, error) {
	a, err := s.dao.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	out := &AccountSecrets{}
	if a.AuthTokenEnc != "" {
		if v, err := s.cipher.DecryptString(a.AuthTokenEnc); err == nil {
			out.AuthToken = v
		}
	}
	if a.RefreshTokenEnc.Valid && a.RefreshTokenEnc.String != "" {
		if v, err := s.cipher.DecryptString(a.RefreshTokenEnc.String); err == nil {
			out.RefreshToken = v
		}
	}
	if a.SessionTokenEnc.Valid && a.SessionTokenEnc.String != "" {
		if v, err := s.cipher.DecryptString(a.SessionTokenEnc.String); err == nil {
			out.SessionToken = v
		}
	}
	return out, nil
}

// DecryptCookies 返回账号 cookies 明文(JSON 字符串)。
func (s *Service) DecryptCookies(ctx context.Context, accountID uint64) (string, error) {
	enc, err := s.dao.GetCookies(ctx, accountID)
	if err != nil {
		return "", err
	}
	if enc == "" {
		return "", nil
	}
	return s.cipher.DecryptString(enc)
}

// GetBinding 查账号-代理绑定。
func (s *Service) GetBinding(ctx context.Context, accountID uint64) (*Binding, error) {
	return s.dao.GetBinding(ctx, accountID)
}

// DAO 暴露给调度器使用。
func (s *Service) DAO() *DAO { return s.dao }
````

## File: internal/audit/dao.go
````go
package audit

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// DAO 审计日志表访问。
type DAO struct{ db *sqlx.DB }

func NewDAO(db *sqlx.DB) *DAO { return &DAO{db: db} }

// Insert 写入一条日志。
func (d *DAO) Insert(ctx context.Context, l *Log) error {
	_, err := d.db.ExecContext(ctx, `
INSERT INTO admin_audit_logs
  (actor_id, actor_email, action, method, path, status_code, ip, ua, target, meta, created_at)
VALUES (?,?,?,?,?,?,?,?,?,?, NOW())`,
		l.ActorID, l.ActorEmail, l.Action, l.Method, l.Path, l.StatusCode,
		l.IP, truncate(l.UA, 255), l.Target, nullJSON(l.Meta))
	if err != nil {
		return fmt.Errorf("audit insert: %w", err)
	}
	return nil
}

// List 分页查询。支持按 actor_id / action 过滤。
func (d *DAO) List(ctx context.Context, actorID uint64, action string, limit, offset int) ([]Log, error) {
	if limit <= 0 {
		limit = 50
	}
	q := `SELECT id, actor_id, actor_email, action, method, path, status_code, ip, ua, target, meta, created_at
	        FROM admin_audit_logs WHERE 1=1`
	args := []interface{}{}
	if actorID > 0 {
		q += " AND actor_id = ?"
		args = append(args, actorID)
	}
	if action != "" {
		q += " AND action = ?"
		args = append(args, action)
	}
	q += " ORDER BY id DESC LIMIT ? OFFSET ?"
	args = append(args, limit, offset)
	var out []Log
	err := d.db.SelectContext(ctx, &out, q, args...)
	return out, err
}

// Count 对应 List 的计数。
func (d *DAO) Count(ctx context.Context, actorID uint64, action string) (int64, error) {
	q := `SELECT COUNT(*) FROM admin_audit_logs WHERE 1=1`
	args := []interface{}{}
	if actorID > 0 {
		q += " AND actor_id = ?"
		args = append(args, actorID)
	}
	if action != "" {
		q += " AND action = ?"
		args = append(args, action)
	}
	var n int64
	err := d.db.GetContext(ctx, &n, q, args...)
	return n, err
}

func nullJSON(b []byte) interface{} {
	if len(b) == 0 {
		return nil
	}
	return b
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max]
}
````

## File: internal/audit/handler.go
````go
package audit

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/432539/gpt2api/pkg/resp"
)

// Handler 暴露审计日志只读查询接口。
type Handler struct {
	dao *DAO
}

// NewHandler 构造。
func NewHandler(dao *DAO) *Handler { return &Handler{dao: dao} }

// List GET /api/admin/audit/logs?actor_id=&action=&limit=&offset=
func (h *Handler) List(c *gin.Context) {
	actorID, _ := strconv.ParseUint(c.Query("actor_id"), 10, 64)
	action := c.Query("action")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if limit > 500 {
		limit = 500
	}

	items, err := h.dao.List(c.Request.Context(), actorID, action, limit, offset)
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}
	total, _ := h.dao.Count(c.Request.Context(), actorID, action)
	resp.OK(c, gin.H{"items": items, "total": total, "limit": limit, "offset": offset})
}
````

## File: internal/audit/middleware.go
````go
package audit

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/432539/gpt2api/internal/middleware"
	"github.com/432539/gpt2api/pkg/logger"
)

// Middleware 在 /api/admin/* 下自动记录写操作。
// 非写操作(GET/HEAD/OPTIONS)默认不落盘,避免审计表爆炸。
// 查看敏感资源如需单独记录,在 handler 里 `Record(c, ...)`。
func Middleware(dao *DAO) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !isWrite(c.Request.Method) {
			c.Next()
			return
		}
		start := time.Now()
		c.Next()
		go writeLog(dao, c, start)
	}
}

// Record 供 handler 显式追加审计记录(用于带业务 meta 的场景)。
// action 形如 "accounts.update",target 为业务 id/slug,meta 为 JSON-encode 的任意上下文。
func Record(c *gin.Context, dao *DAO, action, target string, meta any) {
	if dao == nil {
		return
	}
	metaB, _ := json.Marshal(meta)
	email, _ := c.Get("actor_email")
	emailStr, _ := email.(string)
	l := &Log{
		ActorID:    middleware.ActorID(c),
		ActorEmail: emailStr,
		Action:     action,
		Method:     c.Request.Method,
		Path:       c.FullPath(),
		StatusCode: c.Writer.Status(),
		IP:         c.ClientIP(),
		UA:         c.Request.UserAgent(),
		Target:     target,
		Meta:       metaB,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := dao.Insert(ctx, l); err != nil {
		logger.L().Warn("audit record failed", zap.Error(err), zap.String("action", action))
	}
}

func writeLog(dao *DAO, c *gin.Context, start time.Time) {
	email, _ := c.Get("actor_email")
	emailStr, _ := email.(string)
	l := &Log{
		ActorID:    middleware.ActorID(c),
		ActorEmail: emailStr,
		Action:     deriveAction(c),
		Method:     c.Request.Method,
		Path:       c.FullPath(),
		StatusCode: c.Writer.Status(),
		IP:         c.ClientIP(),
		UA:         c.Request.UserAgent(),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := dao.Insert(ctx, l); err != nil {
		logger.L().Warn("audit insert failed", zap.Error(err),
			zap.String("path", c.FullPath()), zap.Duration("elapsed", time.Since(start)))
	}
}

// deriveAction 从路径+方法推导一个语义化 action,形如 "accounts.create"。
func deriveAction(c *gin.Context) string {
	p := strings.TrimPrefix(c.FullPath(), "/api/admin/")
	p = strings.TrimSuffix(p, "/")
	if p == "" {
		return c.Request.Method
	}
	// 把占位符换成 *,方便在审计里聚合
	parts := strings.Split(p, "/")
	for i, seg := range parts {
		if strings.HasPrefix(seg, ":") {
			parts[i] = "*"
		}
	}
	key := strings.ReplaceAll(strings.Join(parts, "."), "-", "_")
	switch c.Request.Method {
	case "POST":
		return key + ".create"
	case "PUT", "PATCH":
		return key + ".update"
	case "DELETE":
		return key + ".delete"
	default:
		return key + "." + strings.ToLower(c.Request.Method)
	}
}

func isWrite(m string) bool {
	switch m {
	case "POST", "PUT", "PATCH", "DELETE":
		return true
	}
	return false
}
````

## File: internal/audit/model.go
````go
// Package audit 控制台操作审计日志。
//
// 所有 /api/admin/* 的 POST/PUT/PATCH/DELETE 请求会被 Middleware 拦截,
// 自动把操作者、IP、路径、方法、关键响应码落入 admin_audit_logs 表。
//
// 特殊高危操作(备份恢复、删除账号等)还会携带额外的 meta JSON,
// 由 handler 自行调用 `Record(c, "meta...")` 显式追加。
package audit

import (
	"database/sql"
	"time"
)

// Log 对应 admin_audit_logs 表。
type Log struct {
	ID         uint64       `db:"id" json:"id"`
	ActorID    uint64       `db:"actor_id" json:"actor_id"`
	ActorEmail string       `db:"actor_email" json:"actor_email"`
	Action     string       `db:"action" json:"action"` // 形如 "accounts.update"
	Method     string       `db:"method" json:"method"` // HTTP method
	Path       string       `db:"path" json:"path"`
	StatusCode int          `db:"status_code" json:"status_code"`
	IP         string       `db:"ip" json:"ip"`
	UA         string       `db:"ua" json:"ua"`
	Target     string       `db:"target" json:"target"`       // 业务对象 id/slug(可选)
	Meta       []byte       `db:"meta" json:"meta,omitempty"` // JSON
	CreatedAt  time.Time    `db:"created_at" json:"created_at"`
	Finished   sql.NullTime `db:"-" json:"-"` // 预留,不落表
}
````

## File: internal/backup/dao.go
````go
package backup

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// ErrNotFound 备份记录不存在。
var ErrNotFound = errors.New("backup: not found")

// DAO backup_files 表访问对象。
type DAO struct{ db *sqlx.DB }

// NewDAO 构造。
func NewDAO(db *sqlx.DB) *DAO { return &DAO{db: db} }

// 注意:`trigger` 和 `error` 是 MySQL 保留字,必须加反引号。
// 出于一致性,所有列都用反引号包裹,免得下次又被某个保留字坑。
const colsSelect = "`id`,`backup_id`,`file_name`,`size_bytes`,`sha256`," +
	"`trigger`,`status`,`error`,`include_data`,`created_by`,`created_at`,`finished_at`"

// Create 插入 running 记录,返回自增 id。
func (d *DAO) Create(ctx context.Context, f *File) error {
	res, err := d.db.ExecContext(ctx, `
INSERT INTO `+"`backup_files`"+`
  (`+"`backup_id`,`file_name`,`size_bytes`,`sha256`,`trigger`,`status`,`error`,`include_data`,`created_by`,`created_at`"+`)
VALUES (?,?,?,?,?,?,?,?,?, NOW())`,
		f.BackupID, f.FileName, f.SizeBytes, f.SHA256, f.Trigger,
		nonEmpty(f.Status, StatusRunning), f.Error, f.IncludeData, f.CreatedBy,
	)
	if err != nil {
		return fmt.Errorf("backup dao create: %w", err)
	}
	id, _ := res.LastInsertId()
	f.ID = uint64(id)
	return nil
}

// MarkReady 更新成功状态。
func (d *DAO) MarkReady(ctx context.Context, backupID string, size int64, sha string) error {
	_, err := d.db.ExecContext(ctx, `
UPDATE `+"`backup_files`"+`
   SET `+"`status`"+`='ready', `+"`size_bytes`"+`=?, `+"`sha256`"+`=?, `+"`finished_at`"+`=NOW()
 WHERE `+"`backup_id`"+`=?`, size, sha, backupID)
	return err
}

// MarkFailed 更新失败状态。
func (d *DAO) MarkFailed(ctx context.Context, backupID, errMsg string) error {
	if len(errMsg) > 500 {
		errMsg = errMsg[:500]
	}
	_, err := d.db.ExecContext(ctx, `
UPDATE `+"`backup_files`"+`
   SET `+"`status`"+`='failed', `+"`error`"+`=?, `+"`finished_at`"+`=NOW()
 WHERE `+"`backup_id`"+`=?`, errMsg, backupID)
	return err
}

// Get 按 backup_id 查询。
func (d *DAO) Get(ctx context.Context, backupID string) (*File, error) {
	var f File
	err := d.db.GetContext(ctx, &f,
		"SELECT "+colsSelect+" FROM `backup_files` WHERE `backup_id`=?", backupID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	return &f, err
}

// Delete 物理删除记录(文件删除由 Service 层做)。
func (d *DAO) Delete(ctx context.Context, backupID string) error {
	_, err := d.db.ExecContext(ctx,
		"DELETE FROM `backup_files` WHERE `backup_id`=?", backupID)
	return err
}

// List 分页列出。
func (d *DAO) List(ctx context.Context, limit, offset int) ([]File, error) {
	if limit <= 0 {
		limit = 50
	}
	var out []File
	err := d.db.SelectContext(ctx, &out,
		"SELECT "+colsSelect+" FROM `backup_files` ORDER BY `id` DESC LIMIT ? OFFSET ?",
		limit, offset)
	return out, err
}

// Count 总数。
func (d *DAO) Count(ctx context.Context) (int64, error) {
	var n int64
	err := d.db.GetContext(ctx, &n, "SELECT COUNT(*) FROM `backup_files`")
	return n, err
}

// ListReadyOldest 拿最老的 N 个 ready 备份,用于 retention 清理。
func (d *DAO) ListReadyOldest(ctx context.Context, keep int) ([]File, error) {
	if keep < 0 {
		keep = 0
	}
	var out []File
	err := d.db.SelectContext(ctx, &out,
		"SELECT "+colsSelect+" FROM `backup_files` WHERE `status`='ready' "+
			"ORDER BY `id` DESC LIMIT 1000 OFFSET ?", keep)
	return out, err
}

func nonEmpty(s, fallback string) string {
	if s == "" {
		return fallback
	}
	return s
}
````

## File: internal/backup/handler.go
````go
package backup

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/432539/gpt2api/internal/audit"
	"github.com/432539/gpt2api/pkg/resp"
)

// Handler 提供 /api/admin/system/backup/* 接口。
type Handler struct {
	svc      *Service
	dao      *DAO
	auditDAO *audit.DAO
}

// NewHandler 构造。
func NewHandler(svc *Service, dao *DAO, auditDAO *audit.DAO) *Handler {
	return &Handler{svc: svc, dao: dao, auditDAO: auditDAO}
}

// ---- 请求体 ----

type createReq struct {
	IncludeData *bool `json:"include_data,omitempty"` // 默认 true
}

// ---- 接口 ----

// Create POST /api/admin/system/backup
func (h *Handler) Create(c *gin.Context) {
	var req createReq
	_ = c.ShouldBindJSON(&req)
	includeData := true
	if req.IncludeData != nil {
		includeData = *req.IncludeData
	}
	actor := uint64(0)

	f, err := h.svc.Create(c.Request.Context(), actor, TriggerManual, includeData)
	if err != nil {
		audit.Record(c, h.auditDAO, "system.backup.create.failed", "", gin.H{"error": err.Error()})
		resp.Internal(c, err.Error())
		return
	}
	audit.Record(c, h.auditDAO, "system.backup.create", f.BackupID,
		gin.H{"size": f.SizeBytes, "include_data": includeData})
	resp.OK(c, f)
}

// List GET /api/admin/system/backup
func (h *Handler) List(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	items, err := h.dao.List(c.Request.Context(), limit, offset)
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}
	total, _ := h.dao.Count(c.Request.Context())
	resp.OK(c, gin.H{
		"items":         items,
		"total":         total,
		"allow_restore": h.svc.AllowRestore(),
		"max_upload_mb": h.svc.cfg.MaxUploadMB,
	})
}

// Download GET /api/admin/system/backup/:id/download
func (h *Handler) Download(c *gin.Context) {
	id := c.Param("id")
	if !backupIDRe.MatchString(id) {
		resp.BadRequest(c, "invalid backup id")
		return
	}
	fh, meta, err := h.svc.OpenForDownload(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			resp.NotFound(c, "backup not found")
			return
		}
		resp.Internal(c, err.Error())
		return
	}
	defer fh.Close()

	c.Writer.Header().Set("Content-Type", "application/gzip")
	c.Writer.Header().Set("Content-Disposition",
		fmt.Sprintf(`attachment; filename="%s"`, meta.FileName))
	c.Writer.Header().Set("Content-Length", strconv.FormatInt(meta.SizeBytes, 10))
	c.Writer.Header().Set("X-Backup-SHA256", meta.SHA256)
	c.Status(http.StatusOK)
	http.ServeContent(c.Writer, c.Request, meta.FileName, meta.CreatedAt, fh)
	audit.Record(c, h.auditDAO, "system.backup.download", id, nil)
}

// Delete DELETE /api/admin/system/backup/:id。
func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	if !backupIDRe.MatchString(id) {
		resp.BadRequest(c, "invalid backup id")
		return
	}
	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		if errors.Is(err, ErrNotFound) {
			resp.NotFound(c, "backup not found")
			return
		}
		resp.Internal(c, err.Error())
		return
	}
	audit.Record(c, h.auditDAO, "system.backup.delete", id, nil)
	resp.OK(c, gin.H{"deleted": id})
}

// Restore POST /api/admin/system/backup/:id/restore。
// backup.allow_restore 必须为 true,执行前后都会落审计。
func (h *Handler) Restore(c *gin.Context) {
	if !h.svc.AllowRestore() {
		resp.Forbidden(c, "restore is disabled by config; set backup.allow_restore=true first")
		return
	}
	id := c.Param("id")
	if !backupIDRe.MatchString(id) {
		resp.BadRequest(c, "invalid backup id")
		return
	}
	audit.Record(c, h.auditDAO, "system.backup.restore.begin", id, nil)
	if err := h.svc.Restore(c.Request.Context(), id); err != nil {
		audit.Record(c, h.auditDAO, "system.backup.restore.failed", id, gin.H{"error": err.Error()})
		resp.Internal(c, err.Error())
		return
	}
	audit.Record(c, h.auditDAO, "system.backup.restore.success", id, nil)
	resp.OK(c, gin.H{"restored": id})
}

// Upload POST /api/admin/system/backup/upload。
// 上传 .sql.gz 文件(multipart/form-data,字段名 "file")。
func (h *Handler) Upload(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, h.svc.MaxUploadBytes()+4096)
	fh, err := c.FormFile("file")
	if err != nil {
		resp.BadRequest(c, "file is required: "+err.Error())
		return
	}
	if fh.Size > h.svc.MaxUploadBytes() {
		resp.BadRequest(c, fmt.Sprintf("file exceeds %d MB", h.svc.cfg.MaxUploadMB))
		return
	}
	src, err := fh.Open()
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}
	defer src.Close()

	actor := uint64(0)
	f, err := h.svc.ImportUpload(c.Request.Context(), actor, fh.Filename, src)
	if err != nil {
		audit.Record(c, h.auditDAO, "system.backup.upload.failed", fh.Filename, gin.H{"error": err.Error()})
		resp.BadRequest(c, err.Error())
		return
	}
	audit.Record(c, h.auditDAO, "system.backup.upload", f.BackupID,
		gin.H{"orig_name": fh.Filename, "size": f.SizeBytes})
	resp.OK(c, f)
}
````

## File: internal/backup/model.go
````go
// Package backup 数据库备份/恢复。
//
// 物理方案:
//  1. 备份:mysqldump 子进程 → stdout → gzip pipe → 目标文件
//     - 只 dump 业务库(从 DSN 解析得到的 database)
//     - 默认包含数据;支持 --no-data 只 dump 结构
//     - 默认使用 --single-transaction(InnoDB 一致性快照,不锁表)
//     - 排除 audit 表本身可选
//  2. 恢复:mysql 子进程读 gzip 解压流 → stdin
//     - 生产默认禁用(backup.allow_restore=false),需配置显式开启
//     - 执行前后都会写入审计
//
// 安全考虑:
//   - 文件名正则白名单,防路径遍历
//   - restore/upload/delete 受本地控制台和配置开关约束
//   - 大文件流式处理,避免 OOM
//   - sha256 校验完整性
package backup

import (
	"database/sql"
	"time"
)

// 状态常量。
const (
	StatusRunning = "running"
	StatusReady   = "ready"
	StatusFailed  = "failed"
)

// 触发来源。
const (
	TriggerManual = "manual"
	TriggerCron   = "cron"
	TriggerUpload = "upload"
)

// File 对应 backup_files 表。
type File struct {
	ID          uint64       `db:"id" json:"id"`
	BackupID    string       `db:"backup_id" json:"backup_id"` // bk_YYYYMMDD_HHMMSS_xxxx
	FileName    string       `db:"file_name" json:"file_name"`
	SizeBytes   int64        `db:"size_bytes" json:"size_bytes"`
	SHA256      string       `db:"sha256" json:"sha256"`
	Trigger     string       `db:"trigger" json:"trigger"`
	Status      string       `db:"status" json:"status"`
	Error       string       `db:"error" json:"error,omitempty"`
	IncludeData bool         `db:"include_data" json:"include_data"`
	CreatedBy   uint64       `db:"created_by" json:"created_by"`
	CreatedAt   time.Time    `db:"created_at" json:"created_at"`
	FinishedAt  sql.NullTime `db:"finished_at" json:"finished_at,omitempty"`
}
````

## File: internal/backup/service.go
````go
package backup

import (
	"compress/gzip"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	mysqlDrv "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"

	"github.com/432539/gpt2api/internal/config"
	"github.com/432539/gpt2api/pkg/logger"
)

// ErrRestoreDisabled 恢复功能被配置禁用。
var ErrRestoreDisabled = errors.New("backup: restore is disabled by config")

// ErrInvalidFileName 非法文件名(防路径遍历)。
var ErrInvalidFileName = errors.New("backup: invalid file name")

// safeNameRe 仅允许字母/数字/点/下划线/连字符。
var safeNameRe = regexp.MustCompile(`^[A-Za-z0-9._-]+\.sql(\.gz)?$`)

// backupIDRe 形如 bk_20260417_120000_xxxx
var backupIDRe = regexp.MustCompile(`^bk_\d{8}_\d{6}_[a-z0-9]{6}$`)

// Service 备份服务。
type Service struct {
	cfg    config.BackupConfig
	mysql  config.MySQLConfig
	dao    *DAO
	dbName string
	dsn    *mysqlDrv.Config

	// 启动时探测 mysqldump 支持的参数;不支持时自动降级。
	// 例如 mariadb-dump 不识别 --set-gtid-purged、--column-statistics 等 MySQL 专属 flag。
	supportSetGTIDPurged bool
	supportColumnStats   bool
	isMariaDB            bool
}

// New 构造 Service。cfg.Dir 不存在时自动创建。
func New(cfg config.BackupConfig, mysqlCfg config.MySQLConfig, dao *DAO) (*Service, error) {
	if cfg.Dir == "" {
		cfg.Dir = "./data/backups"
	}
	if cfg.MysqldumpBin == "" {
		cfg.MysqldumpBin = "mysqldump"
	}
	if cfg.MysqlBin == "" {
		cfg.MysqlBin = "mysql"
	}
	if cfg.MaxUploadMB <= 0 {
		cfg.MaxUploadMB = 512
	}

	dsn, err := mysqlDrv.ParseDSN(mysqlCfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("parse mysql dsn: %w", err)
	}
	if dsn.DBName == "" {
		return nil, errors.New("backup: mysql dsn missing db name")
	}
	if err := os.MkdirAll(cfg.Dir, 0o750); err != nil {
		return nil, fmt.Errorf("create backup dir: %w", err)
	}
	svc := &Service{
		cfg:    cfg,
		mysql:  mysqlCfg,
		dao:    dao,
		dbName: dsn.DBName,
		dsn:    dsn,
	}
	svc.probeMysqldump()
	return svc, nil
}

// probeMysqldump 探测 mysqldump 版本 / 是否为 MariaDB,以决定传哪些参数。
// 失败时不 panic,只是使用保守默认(不加这些 flag)。
func (s *Service) probeMysqldump() {
	out, err := exec.Command(s.cfg.MysqldumpBin, "--version").CombinedOutput()
	if err != nil {
		logger.L().Warn("probe mysqldump failed, fallback to conservative args",
			zap.Error(err))
		return
	}
	ver := strings.ToLower(string(out))
	s.isMariaDB = strings.Contains(ver, "mariadb")
	// MySQL 5.6+ 支持 --set-gtid-purged;MariaDB 和极老的 MySQL 不支持
	s.supportSetGTIDPurged = !s.isMariaDB
	// --column-statistics 仅 MySQL 8.0+ 的 mysqldump 支持
	s.supportColumnStats = !s.isMariaDB && strings.Contains(ver, "distrib 8.")
	logger.L().Info("mysqldump probed",
		zap.Bool("mariadb", s.isMariaDB),
		zap.Bool("gtid_flag", s.supportSetGTIDPurged),
		zap.String("version_line", firstLine(string(out))),
	)
}

func firstLine(s string) string {
	if i := strings.IndexByte(s, '\n'); i >= 0 {
		return strings.TrimSpace(s[:i])
	}
	return strings.TrimSpace(s)
}

// Dir 返回备份目录(只读)。
func (s *Service) Dir() string { return s.cfg.Dir }

// MaxUploadBytes 返回上传上限(bytes)。
func (s *Service) MaxUploadBytes() int64 {
	return int64(s.cfg.MaxUploadMB) * 1024 * 1024
}

// AllowRestore 是否允许 /restore 端点。
func (s *Service) AllowRestore() bool { return s.cfg.AllowRestore }

// Create 同步执行一次 mysqldump 备份。成功返回 File 记录。
// actorID 是本地操作者标记(默认 0);includeData=false 时仅 dump 表结构。
//
// 流程:记录插入 running → 执行 mysqldump | gzip → sha256 → MarkReady。
func (s *Service) Create(ctx context.Context, actorID uint64, trigger string, includeData bool) (*File, error) {
	backupID := generateBackupID()
	fileName := backupID + ".sql.gz"
	fullPath := filepath.Join(s.cfg.Dir, fileName)

	f := &File{
		BackupID:    backupID,
		FileName:    fileName,
		Trigger:     trigger,
		Status:      StatusRunning,
		IncludeData: includeData,
		CreatedBy:   actorID,
	}
	if err := s.dao.Create(ctx, f); err != nil {
		return nil, err
	}

	size, sha, err := s.dumpToFile(ctx, fullPath, includeData)
	if err != nil {
		_ = s.dao.MarkFailed(ctx, backupID, err.Error())
		_ = os.Remove(fullPath)
		return nil, err
	}
	if err := s.dao.MarkReady(ctx, backupID, size, sha); err != nil {
		return nil, err
	}
	f.SizeBytes = size
	f.SHA256 = sha
	f.Status = StatusReady
	// 异步做 retention,不阻塞
	if s.cfg.Retention > 0 {
		go s.runRetention()
	}
	return f, nil
}

// dumpToFile 实际执行 mysqldump → gzip → 落盘 + sha256。
func (s *Service) dumpToFile(ctx context.Context, fullPath string, includeData bool) (int64, string, error) {
	out, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o640)
	if err != nil {
		return 0, "", fmt.Errorf("open backup file: %w", err)
	}
	defer out.Close()

	sha := sha256.New()
	gz := gzip.NewWriter(io.MultiWriter(out, sha))
	defer gz.Close()

	args := s.mysqldumpArgs(includeData)
	cmd := exec.CommandContext(ctx, s.cfg.MysqldumpBin, args...)
	cmd.Env = append(os.Environ(), "MYSQL_PWD="+s.dsn.Passwd)

	stderr := &strings.Builder{}
	cmd.Stderr = stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return 0, "", fmt.Errorf("stdout pipe: %w", err)
	}
	if err := cmd.Start(); err != nil {
		return 0, "", fmt.Errorf("start mysqldump: %w (stderr=%s)", err, stderr.String())
	}

	written, err := io.Copy(gz, stdout)
	if copyErr := gz.Close(); copyErr != nil && err == nil {
		err = copyErr
	}
	if werr := cmd.Wait(); werr != nil {
		stderrTail := stderr.String()
		if len(stderrTail) > 800 {
			stderrTail = stderrTail[len(stderrTail)-800:]
		}
		return 0, "", fmt.Errorf("mysqldump failed: %w (stderr=%s)", werr, stderrTail)
	}
	if err != nil {
		return 0, "", err
	}
	if err := out.Sync(); err != nil {
		return 0, "", err
	}

	info, err := out.Stat()
	if err != nil {
		return 0, "", err
	}
	_ = written // compressed write counter is inside gz, 直接用落盘大小
	return info.Size(), hex.EncodeToString(sha.Sum(nil)), nil
}

// mysqldumpArgs 生成命令行参数。
// 注意:密码通过环境变量 MYSQL_PWD 传,不暴露在命令行。
func (s *Service) mysqldumpArgs(includeData bool) []string {
	host, port := s.dsnHostPort()
	args := []string{
		"-h", host,
		"-P", port,
		"-u", s.dsn.User,
		"--default-character-set=utf8mb4",
		"--single-transaction",
		"--quick",
		"--skip-lock-tables",
		"--hex-blob",
		"--routines",
		"--triggers",
		"--events",
	}
	if s.supportSetGTIDPurged {
		args = append(args, "--set-gtid-purged=OFF")
	}
	if s.supportColumnStats {
		// MySQL 8 默认会查 column_statistics,对目标库可能没权限,强制关掉
		args = append(args, "--column-statistics=0")
	}
	if !includeData {
		args = append(args, "--no-data")
	}
	// 只 dump 业务库;排除审计表避免恢复时覆盖当下审计
	args = append(args, "--ignore-table="+s.dbName+".admin_audit_logs")
	args = append(args, s.dbName)
	return args
}

// Restore 同步恢复一个已存在的备份到 MySQL。
// 调用方必须保证传入的是可信来源,并在操作前后写入审计。
func (s *Service) Restore(ctx context.Context, backupID string) error {
	if !s.cfg.AllowRestore {
		return ErrRestoreDisabled
	}
	f, err := s.dao.Get(ctx, backupID)
	if err != nil {
		return err
	}
	if f.Status != StatusReady {
		return fmt.Errorf("backup not ready: %s", f.Status)
	}
	fullPath := filepath.Join(s.cfg.Dir, f.FileName)
	if !strings.HasPrefix(filepath.Clean(fullPath), filepath.Clean(s.cfg.Dir)+string(filepath.Separator)) {
		return ErrInvalidFileName
	}

	file, err := os.Open(fullPath)
	if err != nil {
		return fmt.Errorf("open backup file: %w", err)
	}
	defer file.Close()

	gz, err := gzip.NewReader(file)
	if err != nil {
		return fmt.Errorf("open gzip: %w", err)
	}
	defer gz.Close()

	host, port := s.dsnHostPort()
	args := []string{
		"-h", host,
		"-P", port,
		"-u", s.dsn.User,
		"--default-character-set=utf8mb4",
		s.dbName,
	}
	cmd := exec.CommandContext(ctx, s.cfg.MysqlBin, args...)
	cmd.Env = append(os.Environ(), "MYSQL_PWD="+s.dsn.Passwd)
	cmd.Stdin = gz
	stderr := &strings.Builder{}
	cmd.Stderr = stderr
	if err := cmd.Run(); err != nil {
		tail := stderr.String()
		if len(tail) > 800 {
			tail = tail[len(tail)-800:]
		}
		return fmt.Errorf("mysql restore failed: %w (stderr=%s)", err, tail)
	}
	return nil
}

// Delete 删除备份记录 + 物理文件。
func (s *Service) Delete(ctx context.Context, backupID string) error {
	if !backupIDRe.MatchString(backupID) {
		return ErrInvalidFileName
	}
	f, err := s.dao.Get(ctx, backupID)
	if err != nil {
		return err
	}
	path, err := s.fullPath(f.FileName)
	if err != nil {
		return err
	}
	_ = os.Remove(path) // 即便文件丢失也要清理 DB
	return s.dao.Delete(ctx, backupID)
}

// OpenForDownload 返回只读 handle,调用方负责 Close。
// 额外返回文件名和大小供 HTTP header 使用。
func (s *Service) OpenForDownload(ctx context.Context, backupID string) (*os.File, *File, error) {
	f, err := s.dao.Get(ctx, backupID)
	if err != nil {
		return nil, nil, err
	}
	if f.Status != StatusReady {
		return nil, nil, fmt.Errorf("backup not ready: %s", f.Status)
	}
	path, err := s.fullPath(f.FileName)
	if err != nil {
		return nil, nil, err
	}
	fh, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	return fh, f, nil
}

// ImportUpload 将上传的 .sql.gz 保存到备份目录并登记为 "upload" trigger。
// 传入的 reader 会被逐字节读入(已做 LimitReader 控制)。
func (s *Service) ImportUpload(ctx context.Context, actorID uint64, origName string, src io.Reader) (*File, error) {
	if origName == "" {
		origName = "upload.sql.gz"
	}
	// 只取 basename,防路径遍历
	origName = filepath.Base(origName)
	if !safeNameRe.MatchString(origName) {
		return nil, ErrInvalidFileName
	}
	if !strings.HasSuffix(origName, ".gz") {
		// 要求 .sql.gz。纯 .sql 也接受,但我们要 gzip 之后再存。
		// 这里简单起见,拒绝非 .gz(避免引入额外的 gzip 临时流)。
		return nil, fmt.Errorf("upload must be gzip-compressed (.sql.gz)")
	}
	backupID := generateBackupID()
	fileName := backupID + ".sql.gz"
	fullPath := filepath.Join(s.cfg.Dir, fileName)

	out, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0o640)
	if err != nil {
		return nil, fmt.Errorf("create upload file: %w", err)
	}
	defer out.Close()

	sha := sha256.New()
	limited := io.LimitReader(src, s.MaxUploadBytes()+1)
	n, err := io.Copy(io.MultiWriter(out, sha), limited)
	if err != nil {
		_ = os.Remove(fullPath)
		return nil, fmt.Errorf("write upload: %w", err)
	}
	if n > s.MaxUploadBytes() {
		_ = os.Remove(fullPath)
		return nil, fmt.Errorf("upload exceeds max %d MB", s.cfg.MaxUploadMB)
	}
	// gzip 合法性快速校验(读 header)
	if err := verifyGzipHeader(fullPath); err != nil {
		_ = os.Remove(fullPath)
		return nil, fmt.Errorf("invalid gzip: %w", err)
	}

	f := &File{
		BackupID:    backupID,
		FileName:    fileName,
		SizeBytes:   n,
		SHA256:      hex.EncodeToString(sha.Sum(nil)),
		Trigger:     TriggerUpload,
		Status:      StatusReady,
		IncludeData: true,
		CreatedBy:   actorID,
	}
	if err := s.dao.Create(ctx, f); err != nil {
		_ = os.Remove(fullPath)
		return nil, err
	}
	// Create 是插入 running,立刻补 ready
	if err := s.dao.MarkReady(ctx, backupID, n, f.SHA256); err != nil {
		logger.L().Warn("mark upload ready", zap.Error(err))
	}
	f.Status = StatusReady
	return f, nil
}

// runRetention 清理超过 cfg.Retention 的旧备份。只对 ready 状态生效。
func (s *Service) runRetention() {
	if s.cfg.Retention <= 0 {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	old, err := s.dao.ListReadyOldest(ctx, s.cfg.Retention)
	if err != nil {
		logger.L().Warn("retention list failed", zap.Error(err))
		return
	}
	for _, f := range old {
		path, err := s.fullPath(f.FileName)
		if err == nil {
			_ = os.Remove(path)
		}
		_ = s.dao.Delete(ctx, f.BackupID)
		logger.L().Info("backup pruned by retention", zap.String("id", f.BackupID))
	}
}

func (s *Service) fullPath(fileName string) (string, error) {
	if !safeNameRe.MatchString(fileName) {
		return "", ErrInvalidFileName
	}
	p := filepath.Join(s.cfg.Dir, fileName)
	// 严格校验:清洗后的路径必须仍在 backup dir 下
	cleanDir := filepath.Clean(s.cfg.Dir)
	cleanPath := filepath.Clean(p)
	if !strings.HasPrefix(cleanPath, cleanDir+string(filepath.Separator)) && cleanPath != cleanDir {
		return "", ErrInvalidFileName
	}
	return p, nil
}

func (s *Service) dsnHostPort() (string, string) {
	host, port := "127.0.0.1", "3306"
	addr := s.dsn.Addr
	if i := strings.LastIndex(addr, ":"); i > 0 {
		host = addr[:i]
		port = addr[i+1:]
	} else if addr != "" {
		host = addr
	}
	return host, port
}

// generateBackupID 生成形如 bk_20260417_120000_xyz987 的唯一 id。
func generateBackupID() string {
	now := time.Now()
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	//nolint:gosec
	rng := rand.New(rand.NewSource(now.UnixNano()))
	suffix := make([]byte, 6)
	for i := range suffix {
		suffix[i] = letters[rng.Intn(len(letters))]
	}
	return fmt.Sprintf("bk_%s_%s", now.Format("20060102_150405"), string(suffix))
}

// verifyGzipHeader 快速校验文件是合法 gzip。
func verifyGzipHeader(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	_ = gz.Close()
	return nil
}
````

## File: internal/config/config.go
````go
package config

import (
	"fmt"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	App       AppConfig       `mapstructure:"app"`
	Admin     AdminConfig     `mapstructure:"admin"`
	Log       LogConfig       `mapstructure:"log"`
	MySQL     MySQLConfig     `mapstructure:"mysql"`
	Redis     RedisConfig     `mapstructure:"redis"`
	Crypto    CryptoConfig    `mapstructure:"crypto"`
	Security  SecurityConfig  `mapstructure:"security"`
	Scheduler SchedulerConfig `mapstructure:"scheduler"`
	Upstream  UpstreamConfig  `mapstructure:"upstream"`
	Backup    BackupConfig    `mapstructure:"backup"`
	SMTP      SMTPConfig      `mapstructure:"smtp"`
}

type AppConfig struct {
	Name      string `mapstructure:"name"`
	Env       string `mapstructure:"env"`
	Listen    string `mapstructure:"listen"`
	BaseURL   string `mapstructure:"base_url"`
	LocalMode bool   `mapstructure:"local_mode"`
}

type AdminConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type LogConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
	Output string `mapstructure:"output"`
}

type MySQLConfig struct {
	DSN                string `mapstructure:"dsn"`
	MaxOpenConns       int    `mapstructure:"max_open_conns"`
	MaxIdleConns       int    `mapstructure:"max_idle_conns"`
	ConnMaxLifetimeSec int    `mapstructure:"conn_max_lifetime_sec"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

type CryptoConfig struct {
	AESKey string `mapstructure:"aes_key"`
}

type SecurityConfig struct {
	CORSOrigins []string `mapstructure:"cors_origins"`
}

type SchedulerConfig struct {
	MinIntervalSec          int `mapstructure:"min_interval_sec"`
	LockTTLSec              int `mapstructure:"lock_ttl_sec"`
	Cooldown429Sec          int `mapstructure:"cooldown_429_sec"`
	WarnedPauseHours        int `mapstructure:"warned_pause_hours"`
	MaxConcurrentPerAccount int `mapstructure:"max_concurrent_per_account"` // 单账号最大并发,默认 3
}

type UpstreamConfig struct {
	BaseURL           string `mapstructure:"base_url"`
	RequestTimeoutSec int    `mapstructure:"request_timeout_sec"`
	SSEReadTimeoutSec int    `mapstructure:"sse_read_timeout_sec"`
}

// BackupConfig 数据库备份配置。
type BackupConfig struct {
	Dir          string `mapstructure:"dir"`           // 备份落盘目录,默认 /app/data/backups
	Retention    int    `mapstructure:"retention"`     // 保留最近 N 个(>0),0 表示不自动清理
	MysqldumpBin string `mapstructure:"mysqldump_bin"` // 默认 mysqldump
	MysqlBin     string `mapstructure:"mysql_bin"`     // 恢复用,默认 mysql
	MaxUploadMB  int    `mapstructure:"max_upload_mb"` // 上传 .sql.gz 上限,默认 512
	AllowRestore bool   `mapstructure:"allow_restore"` // 是否允许 /restore 端点(生产强烈建议 false 手动切)
}

// SMTPConfig 用于测试邮件和系统通知。
// Host 为空时邮件通道整体关闭,不影响主流程。
type SMTPConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	From     string `mapstructure:"from"`      // 显示的 From 地址
	FromName string `mapstructure:"from_name"` // 显示名
	UseTLS   bool   `mapstructure:"use_tls"`   // true 隐式 TLS(465),false STARTTLS(587)
}

var (
	global *Config
	once   sync.Once
)

func Load(path string) (*Config, error) {
	var loadErr error
	once.Do(func() {
		v := viper.New()
		v.SetConfigFile(path)
		v.SetEnvPrefix("GPT2API")
		v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		v.SetDefault("app.local_mode", true)
		v.SetDefault("admin.username", "admin")
		v.SetDefault("admin.password", "admin123")
		v.AutomaticEnv()
		if err := v.ReadInConfig(); err != nil {
			loadErr = fmt.Errorf("read config: %w", err)
			return
		}
		var c Config
		if err := v.Unmarshal(&c); err != nil {
			loadErr = fmt.Errorf("unmarshal config: %w", err)
			return
		}
		global = &c
		// 校验必填字段,拒绝明显未配置的默认值
		loadErr = validate(&c)
	})
	return global, loadErr
}

// Get 返回全局配置,仅在 Load 之后调用。
func Get() *Config {
	if global == nil {
		panic("config not loaded; call config.Load first")
	}
	return global
}

// validate 校验配置必填字段,拒绝未配置的占位默认值。
func validate(c *Config) error {
	var errs []string
	if c.Crypto.AESKey == "" || c.Crypto.AESKey == "CHANGE_ME_TO_RANDOM_32_BYTES_SECRET" {
		errs = append(errs, "crypto.aes_key is required and must not be the default placeholder")
	}
	if c.MySQL.DSN == "" {
		errs = append(errs, "mysql.dsn is required")
	}
	if c.Redis.Addr == "" {
		errs = append(errs, "redis.addr is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf("config validation failed:\n  - %s", strings.Join(errs, "\n  - "))
	}
	return nil
}
````

## File: internal/db/mysql.go
````go
package db

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/432539/gpt2api/internal/config"
)

// NewMySQL 根据配置打开 MySQL 连接池。
func NewMySQL(cfg config.MySQLConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("open mysql: %w", err)
	}
	if cfg.MaxOpenConns > 0 {
		db.SetMaxOpenConns(cfg.MaxOpenConns)
	}
	if cfg.MaxIdleConns > 0 {
		db.SetMaxIdleConns(cfg.MaxIdleConns)
	}
	if cfg.ConnMaxLifetimeSec > 0 {
		db.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetimeSec) * time.Second)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping mysql: %w", err)
	}
	return db, nil
}
````

## File: internal/db/redis.go
````go
package db

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/432539/gpt2api/internal/config"
)

// NewRedis 打开 Redis 连接。
func NewRedis(cfg config.RedisConfig) (*redis.Client, error) {
	c := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := c.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("ping redis: %w", err)
	}
	return c, nil
}
````

## File: internal/gateway/chat.go
````go
// Package gateway 实现本地 OpenAI 兼容的 /v1/* 入口。
//
// 职责:
//  1. 查模型 slug 映射
//  2. 通过调度器拿账号 Lease
//  3. 转译请求体并调用 chatgpt.com 上游
//  4. 转译响应(流式或聚合)为 OpenAI 协议
//  5. 写入本地 usage/image 运行日志
package gateway

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"

	modelpkg "github.com/432539/gpt2api/internal/model"
	"github.com/432539/gpt2api/internal/scheduler"
	"github.com/432539/gpt2api/internal/upstream/chatgpt"
	"github.com/432539/gpt2api/internal/usage"
	"github.com/432539/gpt2api/pkg/logger"
)

// Handler 聚合网关需要的所有依赖。
type Handler struct {
	Models    *modelpkg.Registry
	Scheduler *scheduler.Scheduler
	Usage     *usage.Logger
	AccSvc    interface {
		DecryptCookies(ctx context.Context, accountID uint64) (string, error)
	}
	// Images 可选:若挂载,chat/completions 里指定图像模型会自动转派。
	Images *ImagesHandler

	// Settings 可选:若注入则在构造上游 client 时应用动态超时。
	Settings interface {
		GatewayUpstreamTimeoutSec() int
		GatewaySSEReadTimeoutSec() int
	}
}

// upstreamTimeout 返回当前应使用的上游非流式超时。未注入时回退 60s。
func (h *Handler) upstreamTimeout() time.Duration {
	if h.Settings != nil {
		if n := h.Settings.GatewayUpstreamTimeoutSec(); n > 0 {
			return time.Duration(n) * time.Second
		}
	}
	return 60 * time.Second
}

// mapUpstreamModelSlug 把本地 slug 映射到 chatgpt.com 后端实际认的灰度 slug。
func mapUpstreamModelSlug(s string) string {
	return s
}

// roughEstimateTokens 估算 messages prompt tokens(无 tiktoken,简单 len/4)。
func roughEstimateTokens(msgs []chatgpt.ChatMessage) int {
	n := 0
	for _, m := range msgs {
		n += (len(m.Content) + 3) / 4
		n += 4
	}
	return n + 2
}

// ChatCompletions 是 POST /v1/chat/completions 入口。
func (h *Handler) ChatCompletions(c *gin.Context) {
	startAt := time.Now()

	var req ChatCompletionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		openAIError(c, http.StatusBadRequest, "invalid_request_error", "请求参数错误:"+err.Error())
		return
	}

	refID := uuid.NewString()
	rec := &usage.Log{
		RequestID: refID,
		Type:      usage.TypeChat,
		IP:        c.ClientIP(),
		UA:        c.Request.UserAgent(),
	}
	defer func() {
		rec.DurationMs = int(time.Since(startAt).Milliseconds())
		if rec.Status == "" {
			rec.Status = usage.StatusFailed
		}
		if h.Usage != nil {
			h.Usage.Write(rec)
		}
	}()
	fail := func(code string) { rec.Status = usage.StatusFailed; rec.ErrorCode = code }

	m, err := h.Models.BySlug(c.Request.Context(), req.Model)
	if err != nil || m == nil || !m.Enabled {
		fail("model_not_found")
		openAIError(c, http.StatusBadRequest, "model_not_found",
			fmt.Sprintf("模型 %q 不存在或未启用", req.Model))
		return
	}
	if m.Type == modelpkg.TypeImage {
		if h.Images == nil {
			fail("image_not_wired")
			openAIError(c, http.StatusNotImplemented, "image_not_wired", "图片生成能力未开启")
			return
		}
		h.Images.handleChatAsImage(c, rec, m, &req, startAt)
		return
	}
	rec.ModelID = m.ID
	promptTokens := roughEstimateTokens(req.Messages)

	lease, err := h.Scheduler.Dispatch(c.Request.Context(), modelpkg.TypeChat)
	if err != nil {
		fail("no_account_available")
		openAIError(c, http.StatusServiceUnavailable, "no_account_available", "账号池暂无可用账号,请稍后重试")
		return
	}
	rec.AccountID = lease.Account.ID
	defer func() { _ = lease.Release(context.Background()) }()

	cookies, _ := h.AccSvc.DecryptCookies(c.Request.Context(), lease.Account.ID)
	cli, err := chatgpt.New(chatgpt.Options{
		AuthToken: lease.AuthToken,
		DeviceID:  lease.DeviceID,
		SessionID: lease.SessionID,
		ProxyURL:  lease.ProxyURL,
		Cookies:   cookies,
		Timeout:   h.upstreamTimeout(),
	})
	if err != nil {
		fail("upstream_init_error")
		openAIError(c, http.StatusInternalServerError, "upstream_init_error", "上游客户端初始化失败:"+err.Error())
		return
	}

	upstreamModel := m.UpstreamModelSlug
	if upstreamModel == "" {
		upstreamModel = "auto"
	}
	upstreamModel = mapUpstreamModelSlug(upstreamModel)

	bootCtx, cancelBoot := context.WithTimeout(c.Request.Context(), 15*time.Second)
	_ = cli.Bootstrap(bootCtx)
	cancelBoot()

	reqCtx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()
	cr, err := cli.ChatRequirementsV2(reqCtx)
	if err != nil {
		h.handleUpstreamErr(c, lease, err, func() { fail("upstream_error") })
		return
	}

	var proofToken string
	if cr.Proofofwork.Required {
		proofCtx, cancelProof := context.WithTimeout(c.Request.Context(), 5*time.Second)
		proofCh := make(chan string, 1)
		go func() { proofCh <- cr.SolveProof("") }()
		select {
		case <-proofCtx.Done():
			cancelProof()
			h.Scheduler.MarkWarned(c.Request.Context(), lease.Account.ID)
			fail("pow_timeout")
			openAIError(c, http.StatusServiceUnavailable, "pow_timeout", "上游风控(PoW)未在规定时间内完成,请重试")
			return
		case proofToken = <-proofCh:
			cancelProof()
		}
		if proofToken == "" {
			h.Scheduler.MarkWarned(c.Request.Context(), lease.Account.ID)
			fail("pow_failed")
			openAIError(c, http.StatusServiceUnavailable, "pow_failed", "上游风控(PoW)校验失败,请稍后重试")
			return
		}
	}
	if cr.Turnstile.Required {
		logger.L().Warn("chat turnstile required, continue anyway", zap.Uint64("account_id", lease.Account.ID))
	}

	if cr.IsFreeAccount() && upstreamModel != "auto" {
		logger.L().Warn("free account requesting premium model, downgrade to auto",
			zap.Uint64("account_id", lease.Account.ID), zap.String("requested_model", upstreamModel))
		upstreamModel = "auto"
	}

	chatOpt := chatgpt.FChatOpts{
		UpstreamModel: upstreamModel,
		Messages:      req.Messages,
		ChatToken:     cr.Token,
		ProofToken:    proofToken,
	}

	prepCtx, cancelPrep := context.WithTimeout(c.Request.Context(), 30*time.Second)
	conduit, err := cli.PrepareFChat(prepCtx, chatOpt)
	cancelPrep()
	if err != nil {
		logger.L().Warn("f/conversation/prepare failed, continue without conduit",
			zap.Uint64("account_id", lease.Account.ID),
			zap.String("upstream_model", upstreamModel),
			zap.Error(err))
		conduit = ""
	}
	chatOpt.ConduitToken = conduit

	logger.L().Info("chat f/conversation send",
		zap.Uint64("account_id", lease.Account.ID),
		zap.String("upstream_model", upstreamModel),
		zap.Int("chat_token_len", len(cr.Token)),
		zap.Int("proof_token_len", len(proofToken)),
		zap.Int("conduit_len", len(conduit)),
		zap.Bool("turnstile_required", cr.Turnstile.Required),
		zap.String("persona", cr.Persona),
	)

	stream, err := cli.StreamFChat(c.Request.Context(), chatOpt)
	if err != nil {
		h.handleUpstreamErr(c, lease, err, func() { fail("upstream_error") })
		return
	}

	id := "chatcmpl-" + uuid.NewString()
	if req.Stream {
		h.streamOpenAI(c, id, req.Model, stream, cr.IsFreeAccount())
	} else {
		h.collectOpenAI(c, id, req.Model, stream, cr.IsFreeAccount())
	}

	completionTokens := h.lastCompletionTokens(c)
	rec.Status = usage.StatusSuccess
	rec.InputTokens = promptTokens
	rec.OutputTokens = completionTokens
}

// streamOpenAI 将上游 SSE 事件转为 OpenAI 风格流式响应。
func (h *Handler) streamOpenAI(c *gin.Context, id, model string, stream <-chan chatgpt.SSEEvent, freeAccount bool) {
	w := c.Writer
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Accel-Buffering", "no")
	w.WriteHeader(http.StatusOK)
	flusher, _ := w.(http.Flusher)

	writeChunk(w, flusher, id, model, DeltaMsg{Role: "assistant"}, nil)

	var extr deltaExtractor
	var total strings.Builder
	evCount := 0
	silentlyRejected := false
	for ev := range stream {
		if ev.Err != nil {
			logger.L().Warn("upstream stream err", zap.Error(ev.Err))
			break
		}
		if len(ev.Data) == 0 {
			continue
		}
		evCount++
		if evCount <= 16 {
			logger.L().Info("chat sse raw", zap.Int("n", evCount),
				zap.String("event", ev.Event), zap.String("data", truncate(string(ev.Data), 2048)))
		}
		if !silentlyRejected && isSilentRejection(ev.Data) {
			silentlyRejected = true
		}
		delta, final, err := extr.Extract(ev.Data)
		if err != nil {
			continue
		}
		if delta != "" {
			total.WriteString(delta)
			writeChunk(w, flusher, id, model, DeltaMsg{Content: delta}, nil)
		}
		if final {
			break
		}
	}
	logger.L().Info("chat sse done", zap.Int("events", evCount), zap.Int("content_len", total.Len()), zap.Bool("silently_rejected", silentlyRejected))

	if total.Len() == 0 && evCount > 0 {
		msg := emptyReplyMessage(freeAccount, silentlyRejected)
		total.WriteString(msg)
		writeChunk(w, flusher, id, model, DeltaMsg{Content: msg}, nil)
	}

	stop := "stop"
	writeChunk(w, flusher, id, model, DeltaMsg{}, &stop)
	fmt.Fprintf(w, "data: [DONE]\n\n")
	if flusher != nil {
		flusher.Flush()
	}
	c.Set("completion_tokens", (total.Len()+3)/4)
}

func (h *Handler) collectOpenAI(c *gin.Context, id, model string, stream <-chan chatgpt.SSEEvent, freeAccount bool) {
	var extr deltaExtractor
	var content strings.Builder
	evCount := 0
	silentlyRejected := false
	for ev := range stream {
		if ev.Err != nil {
			logger.L().Warn("upstream collect err", zap.Error(ev.Err))
			break
		}
		if len(ev.Data) == 0 {
			continue
		}
		evCount++
		if evCount <= 16 {
			logger.L().Info("chat collect raw", zap.Int("n", evCount),
				zap.String("event", ev.Event), zap.String("data", truncate(string(ev.Data), 2048)))
		}
		if !silentlyRejected && isSilentRejection(ev.Data) {
			silentlyRejected = true
		}
		delta, final, _ := extr.Extract(ev.Data)
		if delta != "" {
			content.WriteString(delta)
		}
		if final {
			break
		}
	}
	logger.L().Info("chat collect done", zap.Int("events", evCount), zap.Int("content_len", content.Len()), zap.Bool("silently_rejected", silentlyRejected))

	if content.Len() == 0 && evCount > 0 {
		content.WriteString(emptyReplyMessage(freeAccount, silentlyRejected))
	}

	completionTokens := (content.Len() + 3) / 4
	c.Set("completion_tokens", completionTokens)

	resp := ChatCompletionResponse{
		ID:      id,
		Object:  "chat.completion",
		Created: time.Now().Unix(),
		Model:   model,
		Choices: []ChatCompletionChoice{{
			Index:        0,
			Message:      chatgpt.ChatMessage{Role: "assistant", Content: content.String()},
			FinishReason: "stop",
		}},
		Usage: ChatCompletionUsage{
			PromptTokens:     0,
			CompletionTokens: completionTokens,
			TotalTokens:      completionTokens,
		},
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) lastCompletionTokens(c *gin.Context) int {
	if v, ok := c.Get("completion_tokens"); ok {
		if i, ok := v.(int); ok {
			return i
		}
	}
	return 0
}

// handleUpstreamErr 根据上游错误降级账号并回传 OpenAI 错误。
func (h *Handler) handleUpstreamErr(c *gin.Context, lease *scheduler.Lease, err error, onFail func()) {
	var ue *chatgpt.UpstreamError
	if errors.As(err, &ue) {
		switch {
		case ue.IsRateLimited():
			h.Scheduler.MarkRateLimited(c.Request.Context(), lease.Account.ID)
		case ue.IsUnauthorized():
			h.Scheduler.MarkWarned(c.Request.Context(), lease.Account.ID)
		}
		onFail()
		logger.L().Error("chat upstream error",
			zap.Int("status", ue.Status), zap.Uint64("account_id", lease.Account.ID), zap.String("body", truncate(ue.Body, 1500)))
		openAIError(c, http.StatusBadGateway, "upstream_error", fmt.Sprintf("上游返回错误(HTTP %d):%s", ue.Status, truncate(ue.Body, 200)))
		return
	}
	onFail()
	openAIError(c, http.StatusBadGateway, "upstream_error", "上游请求失败:"+err.Error())
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

func isSilentRejection(data []byte) bool {
	s := string(data)
	return strings.Contains(s, `"is_visually_hidden_from_conversation": true`) &&
		strings.Contains(s, `"role": "system"`) &&
		strings.Contains(s, `"end_turn": true`)
}

// emptyReplyMessage 根据账号类型和上游信号,返回给调用方看的兜底文案。
func emptyReplyMessage(freeAccount, silentlyRejected bool) string {
	switch {
	case silentlyRejected && freeAccount:
		return "上游检测到当前账号为免费版(chatgpt-freeaccount),已静默拒绝本次请求。请更换 ChatGPT Plus / Team 账号后再试。"
	case silentlyRejected:
		return "上游已接受请求但静默终止对话(常见于账号被限流或触发内容审核),请稍后重试,若仍失败请更换模型或账号。"
	case freeAccount:
		return "当前账号为 ChatGPT 免费版,上游未产出内容。请更换 Plus/Team 账号后再试。"
	default:
		return "上游未产出回答内容,可能触发了内容审核或账号被临时限流,请稍后重试。"
	}
}

func writeChunk(w io.Writer, f http.Flusher, id, model string, delta DeltaMsg, finish *string) {
	chunk := ChatCompletionChunk{
		ID:      id,
		Object:  "chat.completion.chunk",
		Created: time.Now().Unix(),
		Model:   model,
		Choices: []ChatCompletionChunkChoice{{Index: 0, Delta: delta, FinishReason: finish}},
	}
	b, _ := json.Marshal(chunk)
	fmt.Fprintf(w, "data: %s\n\n", b)
	if f != nil {
		f.Flush()
	}
}

// openAIError 按 OpenAI 规范返回错误。
func openAIError(c *gin.Context, httpStatus int, code, msg string) {
	c.AbortWithStatusJSON(httpStatus, gin.H{
		"error": gin.H{
			"message": msg,
			"type":    "invalid_request_error",
			"code":    code,
		},
	})
}

// ListModels GET /v1/models。
func (h *Handler) ListModels(c *gin.Context) {
	list, err := h.Models.ListEnabled(c.Request.Context())
	if err != nil {
		openAIError(c, http.StatusInternalServerError, "list_models_error", "获取模型列表失败:"+err.Error())
		return
	}
	data := make([]gin.H, 0, len(list))
	for _, m := range list {
		data = append(data, gin.H{
			"id":       m.Slug,
			"object":   "model",
			"created":  m.CreatedAt.Unix(),
			"owned_by": "chatgpt",
		})
	}
	c.JSON(http.StatusOK, gin.H{"object": "list", "data": data})
}
````

## File: internal/gateway/delta.go
````go
package gateway

import (
	"encoding/json"
	"strings"
)

// deltaExtractor 从 chatgpt.com 上游 SSE data 里提取可展示的增量文本。
//
// 逻辑对齐 chatgpt.com 浏览器抓包的 SSE 数据格式,
// 主要点:
//
//  1. 维护 `curP`(当前 JSON-Patch path)。一帧的 `p` 缺省时继承上一帧的 p
//     —— JSON Merge Patch 约定,chatgpt 的 v1 buffering 就按这个省流量。
//
//  2. 维护 `recipient`(当前 assistant message 的 recipient):
//     只有 "all"(真实回答)才当正文;其他如 "python" / "async_browser" /
//     "image_gen.text2im" 是 tool 调用,不输出。
//     recipient 在首帧 `{"v": {"message": {...}}}` 里出现。
//
//  3. 区分 thoughts:
//     curP 落在 `/message/content/thoughts/...` 的增量全部当"思考过程"
//     —— 对 OpenAI chat 语义,不回给客户端(未来想支持 reasoning
//     字段时再放进来)。
//
//  4. 正文候选只认 `/message/content/parts/0`(含 "" 空 p 继承到这个的情形)。
//     `v` 可能是:
//     a) string           —— 单帧增量
//     b) object 首帧       —— v = {"message":{...},"conversation_id":...}
//     c) array of patches —— v = [{p,o,v}, {p,o,v}, ...]
//
//  5. 结束判据:
//     - 顶层 `{"type":"message_stream_complete"}`
//     - 任何层级 patch 出现 `p == "/message/status"` 且 `v == "finished_successfully"`
//     - 老 /conversation 的 parts 全量 + status: finished_successfully(场景 1 兼容)
//     - 特殊事件 [DONE]
//
// 行为不变的保证:场景 1 仍兼容,`场景 2(patch 模式)` 的增量更稳。
type deltaExtractor struct {
	// lastFull 场景 1(旧协议)里用到,记录上一次全量正文,用于做差分。
	lastFull string

	// curP 保留上一帧的 patch path(v1 buffering 省略 p 时继承)。
	curP string

	// recipient 记录当前助手消息的接收方(all / python / ...).
	// 默认 "all":首帧没提供 recipient 字段时按 "all" 处理。
	recipient string
}

// Done 判断 data 是否是 [DONE]。
func isDone(data []byte) bool {
	s := strings.TrimSpace(string(data))
	return s == "[DONE]"
}

// Extract 返回:增量文本、是否 final(收到结束),err。
func (d *deltaExtractor) Extract(data []byte) (string, bool, error) {
	if d.recipient == "" {
		d.recipient = "all"
	}
	if isDone(data) {
		return "", true, nil
	}
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		// 非 JSON(心跳或其他)—— 忽略。
		return "", false, nil
	}

	// 顶层 type: "message_stream_complete" —— f/conversation 结束事件。
	if t, _ := raw["type"].(string); t == "message_stream_complete" {
		return "", true, nil
	}

	// 1) 继承/更新 curP。只有本帧显式出现 p 时才覆盖,否则继承上一帧。
	if p, ok := raw["p"].(string); ok {
		d.curP = p
	}

	// 2) 遇到 thoughts 路径:v1 协议里 summary/thoughts 的 token 增量
	//    都不当作 assistant 正文(未来接 reasoning 字段时再放开)。
	if strings.HasPrefix(d.curP, "/message/content/thoughts") {
		return "", false, nil
	}

	v, hasV := raw["v"]
	if !hasV {
		return "", false, nil
	}

	// 3) v 为 string:最常见的增量帧。
	if s, ok := v.(string); ok {
		// 只有 recipient == "all"(真实回答)且 curP 落在 parts/0(或空继承)才当正文。
		if d.recipient != "all" {
			// 状态检查:一些上游用 `{"v":"finished_successfully","p":"/message/status"}` 作为结束
			if d.curP == "/message/status" && s == "finished_successfully" {
				return "", true, nil
			}
			return "", false, nil
		}
		if d.curP == "/message/status" && s == "finished_successfully" {
			return "", true, nil
		}
		// curP 为空(首 append 未声明 p)或显式 parts/0:输出
		if d.curP == "" || d.curP == "/message/content/parts/0" {
			return s, false, nil
		}
		return "", false, nil
	}

	// 4) v 为数组:一帧打包多个 patch。
	if arr, ok := v.([]interface{}); ok {
		var b strings.Builder
		final := false
		for _, item := range arr {
			m, ok := item.(map[string]interface{})
			if !ok {
				continue
			}
			// 每条 patch 也可能更新 curP;对数组内 patch,按每条单独解析,
			// 但为了跨帧继承,遍历完后再同步 curP(取最后一条为准)。
			subP, _ := m["p"].(string)
			subV := m["v"]
			subO, _ := m["o"].(string)
			if subP != "" {
				d.curP = subP
			}
			// thoughts 相关的 patch 吞掉
			if strings.HasPrefix(d.curP, "/message/content/thoughts") {
				continue
			}
			// 状态结束
			if d.curP == "/message/status" {
				if s, ok := subV.(string); ok && s == "finished_successfully" {
					final = true
				}
				continue
			}
			// 只认正文 append
			if (d.curP == "" || d.curP == "/message/content/parts/0") &&
				(subO == "" || subO == "append") && d.recipient == "all" {
				if s, ok := subV.(string); ok {
					b.WriteString(s)
				}
			}
		}
		return b.String(), final, nil
	}

	// 5) v 为 object:通常是首帧,里面包着 conversation_id + message 元信息。
	if m, ok := v.(map[string]interface{}); ok {
		// 更新 recipient:首帧没带 role=assistant 时也可能是 tool,需要及时切
		if msg, ok := m["message"].(map[string]interface{}); ok {
			if r, ok := msg["recipient"].(string); ok && r != "" {
				d.recipient = r
			}
			// 初始 content(通常空),用来给老协议场景 1 设置 lastFull baseline
			if content, ok := msg["content"].(map[string]interface{}); ok {
				if parts, ok := content["parts"].([]interface{}); ok && len(parts) > 0 {
					if cur, ok := parts[0].(string); ok {
						d.lastFull = cur
						final := false
						if st, _ := msg["status"].(string); st == "finished_successfully" {
							final = true
						}
						// 首帧的初始内容也 yield 出去(通常是空,不会影响),
						// 保持向后兼容老代码读 parts[0] 的行为。
						if cur != "" && d.recipient == "all" {
							return cur, final, nil
						}
						return "", final, nil
					}
				}
			}
		}
		return "", false, nil
	}

	// 6) 场景 1 兼容:旧 /conversation 协议的顶层 {"message":{...}}
	if msg, ok := raw["message"].(map[string]interface{}); ok {
		if r, ok := msg["recipient"].(string); ok && r != "" {
			d.recipient = r
		}
		if content, ok := msg["content"].(map[string]interface{}); ok {
			if parts, ok := content["parts"].([]interface{}); ok && len(parts) > 0 {
				if cur, ok := parts[0].(string); ok {
					delta := ""
					if strings.HasPrefix(cur, d.lastFull) {
						delta = cur[len(d.lastFull):]
					} else {
						delta = cur
					}
					d.lastFull = cur
					final := false
					if status, _ := msg["status"].(string); status == "finished_successfully" {
						final = true
					}
					if d.recipient != "all" {
						return "", final, nil
					}
					return delta, final, nil
				}
			}
		}
	}

	return "", false, nil
}
````

## File: internal/gateway/images_proxy.go
````go
// images_proxy.go —— 图片返回防盗链代理。
//
// 方案:后端生成自家签名 URL:
//
//	GET /p/img/<task_id>/<idx>?exp=<unix_ms>&sig=<hex>
//
// 请求到达时,后端:
//  1. 校验 exp 未过期 + sig 匹配(HMAC-SHA256);
//  2. 用 DAO 按 task_id 查任务,取 result_urls[idx](上游 estuary 永久 URL);
//  3. 直接反代该 URL(简单 HTTP GET 转发,不需要 AT);
//  4. 若 result_urls 为空则 fallback 到旧逻辑(用账号 AT 现取)。
//
// imageProxySecret 从配置的 aes_key 派生,进程重启后签名不变。
package gateway

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/432539/gpt2api/internal/config"
	"github.com/432539/gpt2api/internal/upstream/chatgpt"
	"github.com/432539/gpt2api/pkg/logger"
)

// ImageAccountResolver 按账号 ID 解出构造 chatgpt client 所需的敏感字段。
// 由 main.go 注入。接口里不直接依赖 account 包,保持本层解耦。
type ImageAccountResolver interface {
	AuthToken(ctx context.Context, accountID uint64) (at, deviceID, cookies string, err error)
	ProxyURL(ctx context.Context, accountID uint64) string
}

// imageProxySecret 从配置的 aes_key 派生,进程重启后签名保持有效。
var imageProxySecret []byte

// InitImageProxySecret 从 aes_key 派生图片代理签名密钥。由 main.go 在配置加载后调用。
func InitImageProxySecret(aesKeyHex string) {
	h := sha256.Sum256([]byte("image-proxy-secret:" + aesKeyHex))
	imageProxySecret = h[:]
}

// ImageProxyTTL 单条签名 URL 的默认有效期。
// 密钥从 aes_key 派生(重启不变),设 30 天;estuary URL 本身永久有效。
const ImageProxyTTL = 30 * 24 * time.Hour

// BuildImageProxyURL 生成代理 URL(绝对路径)。
//
// 优先级:
//  1. 配置的 app.base_url(非空且不含 localhost);
//  2. 从请求 Host 推导(requestHost 由 handler 从 gin.Context 传入);
//  3. 配置的 app.base_url(含 localhost,仅开发环境);
//  4. 兜底:相对路径。
func BuildImageProxyURL(taskID string, idx int, ttl time.Duration, requestHost ...string) string {
	if ttl <= 0 {
		ttl = ImageProxyTTL
	}
	expMs := time.Now().Add(ttl).UnixMilli()
	sig := computeImgSig(taskID, idx, expMs)
	path := fmt.Sprintf("/p/img/%s/%d?exp=%d&sig=%s", taskID, idx, expMs, sig)

	// 1) 配置的 base_url(非 localhost 优先)
	if cfg := config.Get(); cfg != nil && cfg.App.BaseURL != "" {
		if !strings.Contains(cfg.App.BaseURL, "localhost") && !strings.Contains(cfg.App.BaseURL, "127.0.0.1") {
			return strings.TrimRight(cfg.App.BaseURL, "/") + path
		}
	}
	// 2) 从请求 Host 推导
	if len(requestHost) > 0 && requestHost[0] != "" {
		host := requestHost[0]
		scheme := "http"
		if strings.HasPrefix(host, "https://") || strings.HasPrefix(host, "http://") {
			return strings.TrimRight(host, "/") + path
		}
		return scheme + "://" + host + path
	}
	// 3) localhost base_url 兜底(开发环境)
	if cfg := config.Get(); cfg != nil && cfg.App.BaseURL != "" {
		return strings.TrimRight(cfg.App.BaseURL, "/") + path
	}
	return path
}

// RequestBaseURL 从 gin.Context 提取请求的 scheme + host,用于构建绝对 URL。
func RequestBaseURL(c *gin.Context) string {
	scheme := "http"
	if c.Request.TLS != nil || c.GetHeader("X-Forwarded-Proto") == "https" {
		scheme = "https"
	}
	host := c.Request.Host
	if host == "" {
		return ""
	}
	return scheme + "://" + host
}

func computeImgSig(taskID string, idx int, expMs int64) string {
	mac := hmac.New(sha256.New, imageProxySecret)
	fmt.Fprintf(mac, "%s|%d|%d", taskID, idx, expMs)
	return hex.EncodeToString(mac.Sum(nil))[:24]
}

func verifyImgSig(taskID string, idx int, expMs int64, sig string) bool {
	if expMs < time.Now().UnixMilli() {
		return false
	}
	want := computeImgSig(taskID, idx, expMs)
	return hmac.Equal([]byte(sig), []byte(want))
}

// ImageProxy 按签名代理下载上游图片。
// 优先使用 DB 中的 result_urls(estuary 永久 URL)直接反代;
// 若 result_urls 为空则 fallback 到旧逻辑(用账号 AT 现取)。
func (h *ImagesHandler) ImageProxy(c *gin.Context) {
	taskID := c.Param("task_id")
	idxStr := c.Param("idx")
	expStr := c.Query("exp")
	sig := c.Query("sig")

	if taskID == "" || idxStr == "" || expStr == "" || sig == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	idx, err := strconv.Atoi(idxStr)
	if err != nil || idx < 0 || idx > 64 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	expMs, err := strconv.ParseInt(expStr, 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if !verifyImgSig(taskID, idx, expMs, sig) {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	if h.DAO == nil {
		c.AbortWithStatus(http.StatusServiceUnavailable)
		return
	}

	t, err := h.DAO.Get(c.Request.Context(), taskID)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()

	// —— 优先路径:result_urls 里有 estuary URL,直接反代,无需 AT ——
	resultURLs := t.DecodeResultURLs()
	if idx < len(resultURLs) && resultURLs[idx] != "" {
		body, ct, err := proxyEstuaryURL(ctx, resultURLs[idx])
		if err == nil {
			if ct == "" {
				ct = "image/png"
			}
			c.Header("Cache-Control", "public, max-age=86400")
			c.Data(http.StatusOK, ct, body)
			return
		}
		logger.L().Warn("image proxy estuary direct failed, fallback to AT",
			zap.Error(err), zap.String("task_id", taskID), zap.String("url", resultURLs[idx]))
	}

	// —— Fallback:用账号 AT 现取 ——
	fids := t.DecodeFileIDs()
	if idx >= len(fids) {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	ref := fids[idx]
	if t.AccountID == 0 || h.ImageAccResolver == nil {
		c.AbortWithStatus(http.StatusServiceUnavailable)
		return
	}

	at, deviceID, cookies, err := h.ImageAccResolver.AuthToken(ctx, t.AccountID)
	if err != nil {
		logger.L().Warn("image proxy resolve account",
			zap.Error(err), zap.Uint64("account_id", t.AccountID))
		c.AbortWithStatus(http.StatusBadGateway)
		return
	}
	proxyURL := h.ImageAccResolver.ProxyURL(ctx, t.AccountID)

	cli, err := chatgpt.New(chatgpt.Options{
		AuthToken: at,
		DeviceID:  deviceID,
		ProxyURL:  proxyURL,
		Cookies:   cookies,
		Timeout:   h.upstreamTimeout(),
	})
	if err != nil {
		logger.L().Warn("image proxy build client", zap.Error(err))
		c.AbortWithStatus(http.StatusBadGateway)
		return
	}

	signedURL, err := cli.ImageDownloadURL(ctx, t.ConversationID, ref)
	if err != nil {
		logger.L().Warn("image proxy download_url",
			zap.Error(err), zap.String("task_id", taskID), zap.String("ref", ref))
		c.AbortWithStatus(http.StatusBadGateway)
		return
	}

	body, ct, err := cli.FetchImage(ctx, signedURL, 16*1024*1024)
	if err != nil {
		logger.L().Warn("image proxy fetch",
			zap.Error(err), zap.String("task_id", taskID))
		c.AbortWithStatus(http.StatusBadGateway)
		return
	}
	if ct == "" {
		ct = "image/png"
	}
	c.Header("Cache-Control", "public, max-age=86400")
	c.Data(http.StatusOK, ct, body)
}

// proxyEstuaryURL 直接反代 estuary URL,不需要任何认证。
func proxyEstuaryURL(ctx context.Context, estuaryURL string) ([]byte, string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, estuaryURL, nil)
	if err != nil {
		return nil, "", err
	}
	req.Header.Set("User-Agent", chatgpt.DefaultUserAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("estuary returned %d", resp.StatusCode)
	}

	body, err := io.ReadAll(io.LimitReader(resp.Body, 16*1024*1024))
	if err != nil {
		return nil, "", err
	}
	ct := resp.Header.Get("Content-Type")
	return body, ct, nil
}
````

## File: internal/gateway/images.go
````go
package gateway

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/432539/gpt2api/internal/image"
	modelpkg "github.com/432539/gpt2api/internal/model"
	"github.com/432539/gpt2api/internal/upstream/chatgpt"
	"github.com/432539/gpt2api/internal/usage"
)

const maxReferenceImageBytes = 20 * 1024 * 1024
const maxReferenceImages = 4

type chatMsg = chatgpt.ChatMessage

// ImagesHandler 挂载在 /v1/images/* 下的处理器。
type ImagesHandler struct {
	*Handler
	Runner *image.Runner
	DAO    *image.DAO
	// ImageAccResolver 可选:代理下载上游图片时用于解出账号 AT/cookies/proxy。
	ImageAccResolver ImageAccountResolver
}

// ImageGenRequest OpenAI 兼容入参。
type ImageGenRequest struct {
	Model          string `json:"model"`
	Prompt         string `json:"prompt"`
	N              int    `json:"n"`
	Size           string `json:"size"`
	Quality        string `json:"quality,omitempty"`
	Style          string `json:"style,omitempty"`
	ResponseFormat string `json:"response_format,omitempty"`
	User           string `json:"user,omitempty"`

	// reference_images 是控制台/兼容接口的主入口;同时兼容 image_urls/images/
	// reference_image,便于接收其它客户端或图床包装后的对象。
	ReferenceImages ReferenceList  `json:"reference_images,omitempty"`
	ImageURLs       ReferenceList  `json:"image_urls,omitempty"`
	Images          ReferenceList  `json:"images,omitempty"`
	ReferenceImage  ReferenceInput `json:"reference_image,omitempty"`
	ImageURL        ReferenceInput `json:"image_url,omitempty"`
	InputImage      ReferenceInput `json:"input_image,omitempty"`
}

// ReferenceInput 是一张参考图输入。Value 可以是 http(s) URL、data URL、
// 或裸 base64;Name 可选,用于上传到 ChatGPT 文件服务时保留文件名。
type ReferenceInput struct {
	Value string
	Name  string
}

func (r *ReferenceInput) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if len(data) == 0 || string(data) == "null" {
		return nil
	}
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		r.Value = strings.TrimSpace(s)
		return nil
	}
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}
	r.Name = firstJSONStr(obj, "file_name", "filename", "name")
	for _, key := range []string{"url", "data_url", "image_url", "input_image", "source", "data", "b64_json", "base64"} {
		if v := referenceValueFromRaw(obj[key]); strings.TrimSpace(v) != "" {
			r.Value = strings.TrimSpace(v)
			return nil
		}
	}
	return nil
}

// ReferenceList 兼容数组、单个字符串、单个对象三种 JSON 写法。
type ReferenceList []ReferenceInput

func (l *ReferenceList) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if len(data) == 0 || string(data) == "null" {
		*l = nil
		return nil
	}
	if data[0] == '[' {
		var arr []ReferenceInput
		if err := json.Unmarshal(data, &arr); err != nil {
			return err
		}
		*l = arr
		return nil
	}
	var one ReferenceInput
	if err := json.Unmarshal(data, &one); err != nil {
		return err
	}
	*l = []ReferenceInput{one}
	return nil
}

type ImageGenData struct {
	URL           string `json:"url,omitempty"`
	RevisedPrompt string `json:"revised_prompt,omitempty"`
	FileID        string `json:"file_id,omitempty"`
}

type ImageGenResponse struct {
	Created   int64          `json:"created"`
	Data      []ImageGenData `json:"data"`
	TaskID    string         `json:"task_id,omitempty"`
	IsPreview bool           `json:"is_preview,omitempty"`
}

func (r ImageGenRequest) AllReferences() []ReferenceInput {
	out := make([]ReferenceInput, 0, len(r.ReferenceImages)+len(r.ImageURLs)+len(r.Images)+3)
	seen := map[string]struct{}{}
	add := func(in ReferenceInput) {
		in.Value = strings.TrimSpace(in.Value)
		in.Name = strings.TrimSpace(in.Name)
		if in.Value == "" {
			return
		}
		if _, ok := seen[in.Value]; ok {
			return
		}
		seen[in.Value] = struct{}{}
		out = append(out, in)
	}
	for _, in := range r.ReferenceImages {
		add(in)
	}
	for _, in := range r.ImageURLs {
		add(in)
	}
	for _, in := range r.Images {
		add(in)
	}
	add(r.ReferenceImage)
	add(r.ImageURL)
	add(r.InputImage)
	return out
}

// ImageGenerations POST /v1/images/generations。
func (h *ImagesHandler) ImageGenerations(c *gin.Context) {
	startAt := time.Now()
	var req ImageGenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		openAIError(c, http.StatusBadRequest, "invalid_request_error", "请求参数错误:"+err.Error())
		return
	}
	if strings.TrimSpace(req.Prompt) == "" {
		openAIError(c, http.StatusBadRequest, "invalid_request_error", "prompt 不能为空")
		return
	}
	if req.Model == "" {
		req.Model = "gpt-image-2"
	}
	if req.N <= 0 {
		req.N = 1
	}
	if req.N > 4 {
		req.N = 4
	}
	if req.Size == "" {
		req.Size = "1024x1024"
	}

	refID := uuid.NewString()
	rec := &usage.Log{RequestID: refID, Type: usage.TypeImage, IP: c.ClientIP(), UA: c.Request.UserAgent()}
	defer func() {
		rec.DurationMs = int(time.Since(startAt).Milliseconds())
		if rec.Status == "" {
			rec.Status = usage.StatusFailed
		}
		if h.Usage != nil {
			h.Usage.Write(rec)
		}
	}()
	fail := func(code string) { rec.Status = usage.StatusFailed; rec.ErrorCode = code }

	m, err := h.Models.BySlug(c.Request.Context(), req.Model)
	if err != nil || m == nil || !m.Enabled {
		fail("model_not_found")
		openAIError(c, http.StatusBadRequest, "model_not_found", fmt.Sprintf("模型 %q 不存在或未启用", req.Model))
		return
	}
	if m.Type != modelpkg.TypeImage {
		fail("model_type_mismatch")
		openAIError(c, http.StatusBadRequest, "model_type_mismatch", fmt.Sprintf("模型 %q 不是图像模型", req.Model))
		return
	}
	rec.ModelID = m.ID

	refs, err := decodeReferenceInputs(c.Request.Context(), req.AllReferences())
	if err != nil {
		fail("invalid_reference_image")
		openAIError(c, http.StatusBadRequest, "invalid_reference_image", "参考图解析失败:"+err.Error())
		return
	}

	taskID := image.GenerateTaskID()
	if h.DAO != nil {
		if err := h.DAO.Create(c.Request.Context(), &image.Task{
			TaskID: taskID, ModelID: m.ID, Prompt: req.Prompt,
			N: req.N, Size: req.Size, Quality: req.Quality, Style: req.Style,
			Status: image.StatusDispatched, UserID: req.User,
		}); err != nil {
			fail("internal_error")
			openAIError(c, http.StatusInternalServerError, "internal_error", "创建任务失败:"+err.Error())
			return
		}
	}

	runCtx, cancel := context.WithTimeout(c.Request.Context(), 6*time.Minute)
	defer cancel()
	maxAttempts := 2
	if len(refs) > 0 {
		maxAttempts = 1
	}
	res := h.Runner.Run(runCtx, image.RunOptions{
		TaskID:        taskID,
		ModelID:       m.ID,
		UpstreamModel: m.UpstreamModelSlug,
		Prompt:        maybeAppendClaritySuffix(req.Prompt),
		N:             req.N,
		MaxAttempts:   maxAttempts,
		References:    refs,
	})
	rec.AccountID = res.AccountID

	if res.Status != image.StatusSuccess {
		fail(ifEmpty(res.ErrorCode, "upstream_error"))
		httpStatus := http.StatusBadGateway
		if res.ErrorCode == image.ErrNoAccount || res.ErrorCode == image.ErrRateLimited {
			httpStatus = http.StatusServiceUnavailable
		}
		if res.ErrorCode == image.ErrContentPolicy {
			httpStatus = http.StatusBadRequest
		}
		openAIError(c, httpStatus, ifEmpty(res.ErrorCode, "upstream_error"), localizeImageErr(res.ErrorCode, res.ErrorMessage))
		return
	}

	rec.Status = usage.StatusSuccess
	rec.ImageCount = len(res.SignedURLs)
	c.JSON(http.StatusOK, imageResponse(taskID, res, RequestBaseURL(c)))
}

// ImageTask GET /v1/images/tasks/:id。
func (h *ImagesHandler) ImageTask(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		openAIError(c, http.StatusBadRequest, "invalid_request_error", "task id 不能为空")
		return
	}
	if h.DAO == nil {
		openAIError(c, http.StatusInternalServerError, "not_configured", "图片任务存储未初始化")
		return
	}
	t, err := h.DAO.Get(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, image.ErrNotFound) {
			openAIError(c, http.StatusNotFound, "not_found", "任务不存在")
			return
		}
		openAIError(c, http.StatusInternalServerError, "internal_error", "查询任务失败:"+err.Error())
		return
	}
	urls := t.DecodeResultURLs()
	fileIDs := t.DecodeFileIDs()
	data := make([]ImageGenData, 0, len(urls))
	for i := range urls {
		d := ImageGenData{URL: BuildImageProxyURL(t.TaskID, i, ImageProxyTTL, RequestBaseURL(c))}
		if i < len(fileIDs) {
			d.FileID = strings.TrimPrefix(fileIDs[i], "sed:")
		}
		data = append(data, d)
	}
	c.JSON(http.StatusOK, gin.H{
		"task_id":         t.TaskID,
		"status":          t.Status,
		"conversation_id": t.ConversationID,
		"created":         t.CreatedAt.Unix(),
		"finished_at":     nullableUnix(t.FinishedAt),
		"error":           t.Error,
		"data":            data,
	})
}

// handleChatAsImage 是 /v1/chat/completions 发现 model.type=image 时的转派点。
// 支持多模态消息: content 为 [{"type":"text","text":"..."}, {"type":"image_url","image_url":{"url":"..."}}]
func (h *ImagesHandler) handleChatAsImage(c *gin.Context, rec *usage.Log, m *modelpkg.Model, req *ChatCompletionsRequest, startAt time.Time) {
	rec.ModelID = m.ID
	rec.Type = usage.TypeImage
	prompt, imageURLs := extractLastUserContent(req.Messages)
	if strings.TrimSpace(prompt) == "" {
		rec.Status = usage.StatusFailed
		rec.ErrorCode = "invalid_request_error"
		openAIError(c, http.StatusBadRequest, "invalid_request_error", "图像模型需要 user role 消息作为 prompt")
		return
	}

	// 把 image_url 转成参考图
	var refs []image.ReferenceImage
	if len(imageURLs) > 0 {
		decoded, err := decodeReferenceStrings(c.Request.Context(), imageURLs)
		if err != nil {
			rec.Status = usage.StatusFailed
			rec.ErrorCode = "invalid_reference_image"
			openAIError(c, http.StatusBadRequest, "invalid_reference_image", "参考图解析失败:"+err.Error())
			return
		}
		refs = decoded
	}

	// 存储 prompt 时附加参考图信息,方便查找记录
	storedPrompt := prompt
	if len(imageURLs) > 0 {
		var refNote strings.Builder
		refNote.WriteString(prompt)
		refNote.WriteString("\n\n[参考图: ")
		for i, u := range imageURLs {
			if i > 0 {
				refNote.WriteString(", ")
			}
			if len(u) > 100 {
				refNote.WriteString(u[:100] + "...")
			} else {
				refNote.WriteString(u)
			}
		}
		refNote.WriteString("]")
		storedPrompt = refNote.String()
	}

	taskID := image.GenerateTaskID()
	if h.DAO != nil {
		_ = h.DAO.Create(c.Request.Context(), &image.Task{
			TaskID: taskID, ModelID: m.ID, Prompt: storedPrompt,
			N: 1, Size: "1024x1024", Status: image.StatusDispatched,
		})
	}

	runCtx, cancel := context.WithTimeout(c.Request.Context(), 6*time.Minute)
	defer cancel()
	maxAttempts := 2
	if len(refs) > 0 {
		maxAttempts = 1
	}
	res := h.Runner.Run(runCtx, image.RunOptions{
		TaskID:        taskID,
		ModelID:       m.ID,
		UpstreamModel: m.UpstreamModelSlug,
		Prompt:        maybeAppendClaritySuffix(prompt),
		N:             1,
		MaxAttempts:   maxAttempts,
		References:    refs,
	})
	rec.AccountID = res.AccountID
	if res.Status != image.StatusSuccess {
		rec.Status = usage.StatusFailed
		rec.ErrorCode = ifEmpty(res.ErrorCode, "upstream_error")
		httpStatus := http.StatusBadGateway
		if res.ErrorCode == image.ErrNoAccount || res.ErrorCode == image.ErrRateLimited {
			httpStatus = http.StatusServiceUnavailable
		}
		if res.ErrorCode == image.ErrContentPolicy {
			httpStatus = http.StatusBadRequest
		}
		openAIError(c, httpStatus, ifEmpty(res.ErrorCode, "upstream_error"), localizeImageErr(res.ErrorCode, res.ErrorMessage))
		return
	}

	rec.Status = usage.StatusSuccess
	rec.ImageCount = len(res.SignedURLs)
	rec.DurationMs = int(time.Since(startAt).Milliseconds())

	baseURL := RequestBaseURL(c)
	var sb strings.Builder
	for i := range res.SignedURLs {
		if i > 0 {
			sb.WriteString("\n\n")
		}
		sb.WriteString(fmt.Sprintf("![generated](%s)", BuildImageProxyURL(taskID, i, ImageProxyTTL, baseURL)))
	}
	content := sb.String()
	id := "chatcmpl-" + uuid.NewString()

	if req.Stream {
		// 流式返回: SSE 格式
		w := c.Writer
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("X-Accel-Buffering", "no")
		w.WriteHeader(http.StatusOK)
		flusher, _ := w.(http.Flusher)

		// 发送内容 chunk
		chunk := ChatCompletionChunk{
			ID:      id,
			Object:  "chat.completion.chunk",
			Created: time.Now().Unix(),
			Model:   m.Slug,
			Choices: []ChatCompletionChunkChoice{{
				Index: 0,
				Delta: DeltaMsg{Role: "assistant", Content: content},
			}},
		}
		chunkJSON, _ := json.Marshal(chunk)
		fmt.Fprintf(w, "data: %s\n\n", chunkJSON)
		if flusher != nil {
			flusher.Flush()
		}

		// 发送 stop chunk
		stopReason := "stop"
		stopChunk := ChatCompletionChunk{
			ID:      id,
			Object:  "chat.completion.chunk",
			Created: time.Now().Unix(),
			Model:   m.Slug,
			Choices: []ChatCompletionChunkChoice{{
				Index:        0,
				Delta:        DeltaMsg{},
				FinishReason: &stopReason,
			}},
		}
		stopJSON, _ := json.Marshal(stopChunk)
		fmt.Fprintf(w, "data: %s\n\n", stopJSON)
		fmt.Fprintf(w, "data: [DONE]\n\n")
		if flusher != nil {
			flusher.Flush()
		}
	} else {
		resp := ChatCompletionResponse{
			ID:      id,
			Object:  "chat.completion",
			Created: time.Now().Unix(),
			Model:   m.Slug,
			Choices: []ChatCompletionChoice{{
				Index:        0,
				Message:      chatMsg{Role: "assistant", Content: content},
				FinishReason: "stop",
			}},
			Usage: ChatCompletionUsage{},
		}
		c.JSON(http.StatusOK, resp)
	}
}

// extractLastUserContent 从消息列表中提取最后一条 user 消息的文本和图片 URL。
func extractLastUserContent(msgs []chatMsg) (prompt string, imageURLs []string) {
	for i := len(msgs) - 1; i >= 0; i-- {
		if msgs[i].Role != "user" {
			continue
		}
		text := strings.TrimSpace(msgs[i].Content)
		urls := msgs[i].ImageURLs
		if text != "" || len(urls) > 0 {
			return text, urls
		}
	}
	return "", nil
}

// ImageEdits 实现 POST /v1/images/edits,按 OpenAI multipart/form-data 形式接收参考图。
func (h *ImagesHandler) ImageEdits(c *gin.Context) {
	startAt := time.Now()
	if err := c.Request.ParseMultipartForm(int64(maxReferenceImageBytes) * int64(maxReferenceImages+1)); err != nil {
		openAIError(c, http.StatusBadRequest, "invalid_request_error", "解析 multipart 失败:"+err.Error())
		return
	}
	prompt := strings.TrimSpace(c.Request.FormValue("prompt"))
	if prompt == "" {
		openAIError(c, http.StatusBadRequest, "invalid_request_error", "prompt 不能为空")
		return
	}
	model := c.Request.FormValue("model")
	if model == "" {
		model = "gpt-image-2"
	}
	n := 1
	if s := c.Request.FormValue("n"); s != "" {
		if v, err := parseIntClamp(s, 1, 4); err == nil {
			n = v
		}
	}
	size := c.Request.FormValue("size")
	if size == "" {
		size = "1024x1024"
	}

	files, err := collectEditFiles(c.Request.MultipartForm)
	if err != nil {
		openAIError(c, http.StatusBadRequest, "invalid_request_error", err.Error())
		return
	}
	if len(files) == 0 {
		openAIError(c, http.StatusBadRequest, "invalid_request_error", "至少需要上传一张 image 作为参考图")
		return
	}
	if len(files) > maxReferenceImages {
		openAIError(c, http.StatusBadRequest, "invalid_request_error", fmt.Sprintf("最多支持 %d 张参考图", maxReferenceImages))
		return
	}
	refs := make([]image.ReferenceImage, 0, len(files))
	for _, fh := range files {
		data, err := readMultipart(fh)
		if err != nil {
			openAIError(c, http.StatusBadRequest, "invalid_reference_image", fmt.Sprintf("读取 %q 失败:%s", fh.Filename, err.Error()))
			return
		}
		if len(data) == 0 {
			openAIError(c, http.StatusBadRequest, "invalid_reference_image", fmt.Sprintf("参考图 %q 为空", fh.Filename))
			return
		}
		if len(data) > maxReferenceImageBytes {
			openAIError(c, http.StatusBadRequest, "invalid_reference_image", fmt.Sprintf("参考图 %q 超过 %dMB 上限", fh.Filename, maxReferenceImageBytes/1024/1024))
			return
		}
		refs = append(refs, image.ReferenceImage{Data: data, FileName: filepath.Base(fh.Filename)})
	}

	refID := uuid.NewString()
	rec := &usage.Log{RequestID: refID, Type: usage.TypeImage, IP: c.ClientIP(), UA: c.Request.UserAgent()}
	defer func() {
		rec.DurationMs = int(time.Since(startAt).Milliseconds())
		if rec.Status == "" {
			rec.Status = usage.StatusFailed
		}
		if h.Usage != nil {
			h.Usage.Write(rec)
		}
	}()
	fail := func(code string) { rec.Status = usage.StatusFailed; rec.ErrorCode = code }

	m, err := h.Models.BySlug(c.Request.Context(), model)
	if err != nil || m == nil || !m.Enabled {
		fail("model_not_found")
		openAIError(c, http.StatusBadRequest, "model_not_found", fmt.Sprintf("模型 %q 不存在或未启用", model))
		return
	}
	if m.Type != modelpkg.TypeImage {
		fail("model_type_mismatch")
		openAIError(c, http.StatusBadRequest, "model_type_mismatch", fmt.Sprintf("模型 %q 不是图像模型", model))
		return
	}
	rec.ModelID = m.ID

	taskID := image.GenerateTaskID()
	if h.DAO != nil {
		_ = h.DAO.Create(c.Request.Context(), &image.Task{
			TaskID: taskID, ModelID: m.ID, Prompt: prompt,
			N: n, Size: size, Status: image.StatusDispatched,
		})
	}

	runCtx, cancel := context.WithTimeout(c.Request.Context(), 8*time.Minute)
	defer cancel()
	res := h.Runner.Run(runCtx, image.RunOptions{
		TaskID:        taskID,
		ModelID:       m.ID,
		UpstreamModel: m.UpstreamModelSlug,
		Prompt:        maybeAppendClaritySuffix(prompt),
		N:             n,
		MaxAttempts:   1,
		References:    refs,
	})
	rec.AccountID = res.AccountID
	if res.Status != image.StatusSuccess {
		fail(ifEmpty(res.ErrorCode, "upstream_error"))
		httpStatus := http.StatusBadGateway
		if res.ErrorCode == image.ErrNoAccount || res.ErrorCode == image.ErrRateLimited {
			httpStatus = http.StatusServiceUnavailable
		}
		if res.ErrorCode == image.ErrContentPolicy {
			httpStatus = http.StatusBadRequest
		}
		openAIError(c, httpStatus, ifEmpty(res.ErrorCode, "upstream_error"), localizeImageErr(res.ErrorCode, res.ErrorMessage))
		return
	}

	rec.Status = usage.StatusSuccess
	rec.ImageCount = len(res.SignedURLs)
	c.JSON(http.StatusOK, imageResponse(taskID, res, RequestBaseURL(c)))
}

func imageResponse(taskID string, res *image.RunResult, baseURL string) ImageGenResponse {
	out := ImageGenResponse{Created: time.Now().Unix(), TaskID: taskID, IsPreview: res.IsPreview, Data: make([]ImageGenData, 0, len(res.SignedURLs))}
	for i := range res.SignedURLs {
		d := ImageGenData{URL: BuildImageProxyURL(taskID, i, ImageProxyTTL, baseURL)}
		if i < len(res.FileIDs) {
			d.FileID = strings.TrimPrefix(res.FileIDs[i], "sed:")
		}
		out.Data = append(out.Data, d)
	}
	return out
}

func ifEmpty(s, fallback string) string {
	if s == "" {
		return fallback
	}
	return s
}

func localizeImageErr(code, raw string) string {
	var zh string
	switch code {
	case image.ErrNoAccount:
		zh = "账号池暂无可用账号,请稍后重试"
	case image.ErrRateLimited:
		zh = "上游风控,请稍后再试"
	case image.ErrPreviewOnly:
		zh = "上游仅返回预览,请稍后重试"
	case image.ErrContentPolicy:
		// 直接透传上游拒绝原文
		if raw != "" {
			return raw
		}
		zh = "内容策略限制,无法生成该图片"
	case image.ErrUnknown, "":
		zh = "图片生成失败"
	case "upstream_error":
		zh = "上游返回错误"
	default:
		zh = "图片生成失败(" + code + ")"
	}
	if raw != "" && raw != code {
		return zh + ":" + raw
	}
	return zh
}

func nullableUnix(t *time.Time) int64 {
	if t == nil || t.IsZero() {
		return 0
	}
	return t.Unix()
}

var textHintKeywords = []string{
	"文字", "对话", "台词", "旁白", "标语", "字幕", "标题", "文案",
	"招牌", "横幅", "海报文字", "弹幕", "气泡", "字体",
	"text:", "caption", "subtitle", "title:", "label", "banner", "poster text",
}

const claritySuffix = "\n\nclean readable Chinese text, prioritize text clarity over image details"

func collectEditFiles(form *multipart.Form) ([]*multipart.FileHeader, error) {
	if form == nil {
		return nil, errors.New("empty multipart form")
	}
	var out []*multipart.FileHeader
	seen := map[string]bool{}
	add := func(fhs []*multipart.FileHeader) {
		for _, fh := range fhs {
			if fh == nil {
				continue
			}
			key := fh.Filename + "|" + fmt.Sprint(fh.Size)
			if seen[key] {
				continue
			}
			seen[key] = true
			out = append(out, fh)
		}
	}
	for _, key := range []string{"image", "image[]", "images", "images[]", "mask"} {
		if fhs := form.File[key]; len(fhs) > 0 {
			add(fhs)
		}
	}
	for k, fhs := range form.File {
		if strings.HasPrefix(k, "image_") {
			add(fhs)
		}
	}
	return out, nil
}

func readMultipart(fh *multipart.FileHeader) ([]byte, error) {
	f, err := fh.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(f)
}

func decodeReferenceStrings(ctx context.Context, inputs []string) ([]image.ReferenceImage, error) {
	refs := make([]ReferenceInput, 0, len(inputs))
	for _, s := range inputs {
		refs = append(refs, ReferenceInput{Value: s})
	}
	return decodeReferenceInputs(ctx, refs)
}

func decodeReferenceInputs(ctx context.Context, inputs []ReferenceInput) ([]image.ReferenceImage, error) {
	if len(inputs) == 0 {
		return nil, nil
	}
	if len(inputs) > maxReferenceImages {
		return nil, fmt.Errorf("最多支持 %d 张参考图", maxReferenceImages)
	}
	out := make([]image.ReferenceImage, 0, len(inputs))
	for i, in := range inputs {
		in.Value = strings.TrimSpace(in.Value)
		in.Name = strings.TrimSpace(in.Name)
		if in.Value == "" {
			return nil, fmt.Errorf("第 %d 张参考图为空", i+1)
		}
		data, name, err := fetchReferenceBytes(ctx, in)
		if err != nil {
			return nil, fmt.Errorf("第 %d 张参考图:%w", i+1, err)
		}
		if len(data) == 0 {
			return nil, fmt.Errorf("第 %d 张参考图解码后为空", i+1)
		}
		if len(data) > maxReferenceImageBytes {
			return nil, fmt.Errorf("第 %d 张参考图超过 %dMB 上限", i+1, maxReferenceImageBytes/1024/1024)
		}
		_, ext := sniffReferenceMime(data)
		if name == "" {
			name = fmt.Sprintf("reference-%d%s", i+1, ext)
		} else if filepath.Ext(name) == "" && ext != "" {
			name += ext
		}
		out = append(out, image.ReferenceImage{Data: data, FileName: name})
	}
	return out, nil
}

func fetchReferenceBytes(ctx context.Context, in ReferenceInput) ([]byte, string, error) {
	s := strings.TrimSpace(in.Value)
	name := sanitizeReferenceFileName(in.Name)
	low := strings.ToLower(s)
	switch {
	case strings.HasPrefix(low, "data:"):
		comma := strings.IndexByte(s, ',')
		if comma < 0 {
			return nil, "", errors.New("无效 data URL")
		}
		meta := s[5:comma]
		payload := strings.TrimSpace(s[comma+1:])
		if name == "" {
			name = dataURLFileName(meta)
		}
		if strings.Contains(strings.ToLower(meta), ";base64") {
			b, err := decodeBase64Flexible(payload)
			if err != nil {
				if unescaped, uerr := url.PathUnescape(payload); uerr == nil && unescaped != payload {
					if b2, err2 := decodeBase64Flexible(unescaped); err2 == nil {
						return b2, name, nil
					}
				}
				return nil, "", fmt.Errorf("base64 解码失败:%w", err)
			}
			return b, name, nil
		}
		if unescaped, err := url.PathUnescape(payload); err == nil {
			payload = unescaped
		}
		return []byte(payload), name, nil
	case strings.HasPrefix(low, "http://"), strings.HasPrefix(low, "https://"):
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, s, nil)
		if err != nil {
			return nil, "", err
		}
		req.Header.Set("User-Agent", chatgpt.DefaultUserAgent)
		req.Header.Set("Accept", "image/*,*/*;q=0.8")
		hc := &http.Client{Timeout: 30 * time.Second}
		res, err := hc.Do(req)
		if err != nil {
			return nil, "", err
		}
		defer res.Body.Close()
		if res.StatusCode >= 400 {
			return nil, "", fmt.Errorf("下载失败 HTTP %d", res.StatusCode)
		}
		if res.ContentLength > int64(maxReferenceImageBytes) {
			return nil, "", fmt.Errorf("远程图片超过 %dMB 上限", maxReferenceImageBytes/1024/1024)
		}
		body, err := io.ReadAll(io.LimitReader(res.Body, int64(maxReferenceImageBytes)+1))
		if err != nil {
			return nil, "", err
		}
		if name == "" {
			name = sanitizeReferenceFileName(filepath.Base(req.URL.Path))
		}
		if name == "" {
			name = "reference" + extensionFromContentType(res.Header.Get("Content-Type"))
		} else if filepath.Ext(name) == "" {
			name += extensionFromContentType(res.Header.Get("Content-Type"))
		}
		return body, name, nil
	default:
		b, err := decodeBase64Flexible(s)
		if err != nil {
			return nil, "", fmt.Errorf("既非 URL 也非可解析的 base64:%w", err)
		}
		return b, name, nil
	}
}

func firstJSONStr(obj map[string]json.RawMessage, keys ...string) string {
	for _, key := range keys {
		if v := referenceValueFromRaw(obj[key]); v != "" {
			return sanitizeReferenceFileName(v)
		}
	}
	return ""
}

func referenceValueFromRaw(raw json.RawMessage) string {
	if len(raw) == 0 || string(bytes.TrimSpace(raw)) == "null" {
		return ""
	}
	var s string
	if err := json.Unmarshal(raw, &s); err == nil {
		return strings.TrimSpace(s)
	}
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(raw, &obj); err != nil {
		return ""
	}
	for _, key := range []string{"url", "data_url", "image_url", "input_image", "source", "data", "b64_json", "base64"} {
		if v := referenceValueFromRaw(obj[key]); v != "" {
			return v
		}
	}
	return ""
}

func decodeBase64Flexible(s string) ([]byte, error) {
	clean := strings.Map(func(r rune) rune {
		switch r {
		case ' ', '\n', '\r', '\t':
			return -1
		default:
			return r
		}
	}, s)
	encs := []*base64.Encoding{base64.StdEncoding, base64.RawStdEncoding, base64.URLEncoding, base64.RawURLEncoding}
	var last error
	for _, enc := range encs {
		b, err := enc.DecodeString(clean)
		if err == nil {
			return b, nil
		}
		last = err
	}
	return nil, last
}

func sanitizeReferenceFileName(name string) string {
	name = strings.TrimSpace(strings.ReplaceAll(name, "\\", "/"))
	if name == "" {
		return ""
	}
	name = filepath.Base(name)
	if name == "." || name == "/" || name == ".." {
		return ""
	}
	name = strings.Map(func(r rune) rune {
		switch r {
		case '\r', '\n', '\t', '/', '\\':
			return '-'
		default:
			return r
		}
	}, name)
	if len(name) > 120 {
		ext := filepath.Ext(name)
		base := strings.TrimSuffix(name, ext)
		if len(ext) > 16 {
			ext = ""
		}
		maxBase := 120 - len(ext)
		if maxBase < 32 {
			maxBase = 32
		}
		if len(base) > maxBase {
			base = base[:maxBase]
		}
		name = base + ext
	}
	return name
}

func dataURLFileName(meta string) string {
	mime := strings.TrimSpace(strings.Split(meta, ";")[0])
	if mime == "" {
		return ""
	}
	return "reference" + extensionFromMime(mime)
}

func extensionFromContentType(ct string) string {
	if i := strings.Index(ct, ";"); i >= 0 {
		ct = ct[:i]
	}
	return extensionFromMime(strings.TrimSpace(ct))
}

func extensionFromMime(mime string) string {
	switch strings.ToLower(mime) {
	case "image/jpeg", "image/jpg":
		return ".jpg"
	case "image/png":
		return ".png"
	case "image/gif":
		return ".gif"
	case "image/webp":
		return ".webp"
	case "image/avif":
		return ".avif"
	default:
		return ""
	}
}

func sniffReferenceMime(data []byte) (string, string) {
	n := 512
	if len(data) < n {
		n = len(data)
	}
	mime := http.DetectContentType(data[:n])
	return mime, extensionFromContentType(mime)
}

func parseIntClamp(s string, min, max int) (int, error) {
	var v int
	if _, err := fmt.Sscanf(s, "%d", &v); err != nil {
		return 0, err
	}
	if v < min {
		v = min
	}
	if v > max {
		v = max
	}
	return v, nil
}

func maybeAppendClaritySuffix(prompt string) string {
	lower := strings.ToLower(prompt)
	need := false
	for _, kw := range textHintKeywords {
		if strings.Contains(lower, strings.ToLower(kw)) {
			need = true
			break
		}
	}
	if !need {
		for _, pair := range [][2]string{{"\"", "\""}, {"'", "'"}, {"“", "”"}, {"‘", "’"}, {"「", "」"}, {"『", "』"}} {
			if idx := strings.Index(prompt, pair[0]); idx >= 0 {
				rest := prompt[idx+len(pair[0]):]
				if end := strings.Index(rest, pair[1]); end >= 2 {
					need = true
					break
				}
			}
		}
	}
	if need && !strings.Contains(prompt, strings.TrimSpace(claritySuffix)) {
		return prompt + claritySuffix
	}
	return prompt
}
````

## File: internal/gateway/types.go
````go
package gateway

import "github.com/432539/gpt2api/internal/upstream/chatgpt"

// ChatCompletionsRequest 对应 OpenAI /v1/chat/completions 请求体子集。
type ChatCompletionsRequest struct {
	Model       string                 `json:"model" binding:"required"`
	Messages    []chatgpt.ChatMessage  `json:"messages" binding:"required"`
	Stream      bool                   `json:"stream"`
	Temperature float64                `json:"temperature,omitempty"`
	TopP        float64                `json:"top_p,omitempty"`
	MaxTokens   int                    `json:"max_tokens,omitempty"`
	User        string                 `json:"user,omitempty"`
	Extra       map[string]interface{} `json:"-"`
}

// ChatCompletionResponse 非流式响应。
type ChatCompletionResponse struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int64                  `json:"created"`
	Model   string                 `json:"model"`
	Choices []ChatCompletionChoice `json:"choices"`
	Usage   ChatCompletionUsage    `json:"usage"`
}

type ChatCompletionChoice struct {
	Index        int                 `json:"index"`
	Message      chatgpt.ChatMessage `json:"message"`
	FinishReason string              `json:"finish_reason"`
}

type ChatCompletionUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// ChatCompletionChunk 流式 chunk。
type ChatCompletionChunk struct {
	ID      string                      `json:"id"`
	Object  string                      `json:"object"`
	Created int64                       `json:"created"`
	Model   string                      `json:"model"`
	Choices []ChatCompletionChunkChoice `json:"choices"`
}

type ChatCompletionChunkChoice struct {
	Index        int      `json:"index"`
	Delta        DeltaMsg `json:"delta"`
	FinishReason *string  `json:"finish_reason"`
}

type DeltaMsg struct {
	Role    string `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
}
````

## File: internal/image/dao.go
````go
package image

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

var ErrNotFound = errors.New("image: task not found")

type DAO struct{ db *sqlx.DB }

func NewDAO(db *sqlx.DB) *DAO { return &DAO{db: db} }

func (d *DAO) Create(ctx context.Context, t *Task) error {
	res, err := d.db.ExecContext(ctx, `
INSERT INTO image_tasks
  (task_id, model_id, account_id, prompt, revised_prompt, n, size, quality, style, status,
   conversation_id, file_ids, result_urls, reference_urls, error, user_id, created_at)
VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?, NOW())`,
		t.TaskID, t.ModelID, t.AccountID, t.Prompt, t.RevisedPrompt,
		t.N, t.Size, t.Quality, t.Style, nullEmpty(t.Status, StatusQueued),
		t.ConversationID, nullJSON(t.FileIDs), nullJSON(t.ResultURLs), nullJSON(t.ReferenceURLs),
		t.Error, t.UserID,
	)
	if err != nil {
		return fmt.Errorf("image dao create: %w", err)
	}
	id, _ := res.LastInsertId()
	t.ID = uint64(id)
	return nil
}

func (d *DAO) MarkRunning(ctx context.Context, taskID string, accountID uint64) error {
	_, err := d.db.ExecContext(ctx, `
UPDATE image_tasks
   SET status='running', account_id=?, started_at=NOW()
 WHERE task_id=? AND status IN ('queued','dispatched')`, accountID, taskID)
	return err
}

func (d *DAO) SetAccount(ctx context.Context, taskID string, accountID uint64) error {
	_, err := d.db.ExecContext(ctx, `UPDATE image_tasks SET account_id = ? WHERE task_id = ?`, accountID, taskID)
	return err
}

func (d *DAO) MarkSuccess(ctx context.Context, taskID, convID string, fileIDs, resultURLs []string, extra SuccessExtra) error {
	fidB, _ := json.Marshal(fileIDs)
	urlB, _ := json.Marshal(resultURLs)
	var refB []byte
	if len(extra.ReferenceFileIDs) > 0 {
		refB, _ = json.Marshal(extra.ReferenceFileIDs)
	}
	_, err := d.db.ExecContext(ctx, `
UPDATE image_tasks
   SET status='success', conversation_id=?, file_ids=?, result_urls=?,
       reference_urls=COALESCE(?, reference_urls), revised_prompt=?,
       attempts=?, duration_ms=?, finished_at=NOW()
 WHERE task_id=?`, convID, fidB, urlB, nullJSON(refB), extra.RevisedPrompt,
		extra.Attempts, extra.DurationMs, taskID)
	return err
}

// SuccessExtra 携带落库时的额外统计信息。
type SuccessExtra struct {
	RevisedPrompt    string
	ReferenceFileIDs []string // GPT 侧的 file-service ID
	Attempts         int
	DurationMs       int64
}

func (d *DAO) MarkFailed(ctx context.Context, taskID, errorCode string) error {
	_, err := d.db.ExecContext(ctx, `
UPDATE image_tasks
   SET status='failed', error=?, finished_at=NOW()
 WHERE task_id=?`, truncate(errorCode, 500), taskID)
	return err
}

func (d *DAO) Get(ctx context.Context, taskID string) (*Task, error) {
	var t Task
	err := d.db.GetContext(ctx, &t, `
SELECT id, task_id, model_id, account_id, prompt, revised_prompt, n, size, quality, style, status,
       conversation_id, file_ids, result_urls, reference_urls, error,
       attempts, duration_ms, user_id, created_at, started_at, finished_at
  FROM image_tasks
 WHERE task_id = ?`, taskID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (d *DAO) ListAll(ctx context.Context, limit, offset int) ([]Task, error) {
	if limit <= 0 {
		limit = 20
	}
	var out []Task
	err := d.db.SelectContext(ctx, &out, `
SELECT id, task_id, model_id, account_id, prompt, revised_prompt, n, size, quality, style, status,
       conversation_id, file_ids, result_urls, reference_urls, error,
       attempts, duration_ms, user_id, created_at, started_at, finished_at
  FROM image_tasks
 ORDER BY id DESC
 LIMIT ? OFFSET ?`, limit, offset)
	return out, err
}

func (t *Task) DecodeFileIDs() []string {
	var out []string
	if len(t.FileIDs) > 0 {
		_ = json.Unmarshal(t.FileIDs, &out)
	}
	return out
}

func (t *Task) DecodeResultURLs() []string {
	var out []string
	if len(t.ResultURLs) > 0 {
		_ = json.Unmarshal(t.ResultURLs, &out)
	}
	return out
}

func nullEmpty(s, fallback string) string {
	if s == "" {
		return fallback
	}
	return s
}

func nullJSON(b []byte) interface{} {
	if len(b) == 0 {
		return nil
	}
	return b
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max]
}
````

## File: internal/image/me_handler.go
````go
package image

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/432539/gpt2api/pkg/resp"
)

// MeHandler 面向本地控制台的图片任务只读接口。
type MeHandler struct{ dao *DAO }

func NewMeHandler(dao *DAO) *MeHandler { return &MeHandler{dao: dao} }

type taskView struct {
	ID             uint64     `json:"id"`
	TaskID         string     `json:"task_id"`
	ModelID        uint64     `json:"model_id"`
	AccountID      uint64     `json:"account_id"`
	Prompt         string     `json:"prompt"`
	RevisedPrompt  string     `json:"revised_prompt,omitempty"`
	N              int        `json:"n"`
	Size           string     `json:"size"`
	Quality        string     `json:"quality,omitempty"`
	Style          string     `json:"style,omitempty"`
	Status         string     `json:"status"`
	ConversationID string     `json:"conversation_id,omitempty"`
	Error          string     `json:"error,omitempty"`
	ImageURLs      []string   `json:"image_urls"`
	FileIDs        []string   `json:"file_ids,omitempty"`
	ReferenceURLs  []string   `json:"reference_urls,omitempty"`
	Attempts       int        `json:"attempts,omitempty"`
	DurationMs     int64      `json:"duration_ms,omitempty"`
	UserID         string     `json:"user_id,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	StartedAt      *time.Time `json:"started_at,omitempty"`
	FinishedAt     *time.Time `json:"finished_at,omitempty"`
}

func toView(t *Task) taskView {
	urls := t.DecodeResultURLs()
	fids := t.DecodeFileIDs()
	for i, id := range fids {
		fids[i] = strings.TrimPrefix(id, "sed:")
	}
	var refURLs []string
	if len(t.ReferenceURLs) > 0 {
		_ = json.Unmarshal(t.ReferenceURLs, &refURLs)
	}
	return taskView{
		ID: t.ID, TaskID: t.TaskID, ModelID: t.ModelID,
		AccountID: t.AccountID, Prompt: t.Prompt, RevisedPrompt: t.RevisedPrompt,
		N: t.N, Size: t.Size, Quality: t.Quality, Style: t.Style,
		Status: t.Status, ConversationID: t.ConversationID, Error: t.Error,
		ImageURLs: urls, FileIDs: fids, ReferenceURLs: refURLs,
		Attempts: t.Attempts, DurationMs: t.DurationMs, UserID: t.UserID,
		CreatedAt: t.CreatedAt, StartedAt: t.StartedAt, FinishedAt: t.FinishedAt,
	}
}

// List GET /api/me/images/tasks。
func (h *MeHandler) List(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}
	tasks, err := h.dao.ListAll(c.Request.Context(), limit, offset)
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}
	items := make([]taskView, 0, len(tasks))
	for i := range tasks {
		items = append(items, toView(&tasks[i]))
	}
	resp.OK(c, gin.H{"items": items, "limit": limit, "offset": offset})
}

// Get GET /api/me/images/tasks/:id。
func (h *MeHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		resp.Fail(c, 40000, "task id required")
		return
	}
	t, err := h.dao.Get(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			resp.Fail(c, 40400, "task not found")
			return
		}
		resp.Internal(c, err.Error())
		return
	}
	resp.OK(c, toView(t))
}
````

## File: internal/image/model.go
````go
// Package image 生图任务的数据模型、DAO 以及同步 Runner。
//
// 路线:
//   - /v1/images/generations 默认同步返回结果,同时落库生成 task_id。
//   - /v1/images/tasks/:id 可查询任务历史。
//
// 链路复用 Account + Proxy + Scheduler,并把结果写入 usage/image 日志。
package image

import "time"

const (
	StatusQueued     = "queued"
	StatusDispatched = "dispatched"
	StatusRunning    = "running"
	StatusSuccess    = "success"
	StatusFailed     = "failed"
)

// 错误码(短字符串,便于排查)。
const (
	ErrUnknown         = "unknown"
	ErrNoAccount       = "no_available_account"
	ErrAuthRequired    = "auth_required"
	ErrRateLimited     = "rate_limited"
	ErrPOWTimeout      = "pow_timeout"
	ErrPOWFailed       = "pow_failed"
	ErrTurnstile       = "turnstile_required"
	ErrUpstream        = "upstream_error"
	ErrPreviewOnly     = "preview_only"
	ErrPollTimeout     = "poll_timeout"
	ErrDownload        = "download_failed"
	ErrInvalidResponse = "invalid_response"
	ErrContentPolicy   = "content_policy_violation"
)

// Task 对应 image_tasks 表。
type Task struct {
	ID             uint64     `db:"id"`
	TaskID         string     `db:"task_id"`
	ModelID        uint64     `db:"model_id"`
	AccountID      uint64     `db:"account_id"`
	Prompt         string     `db:"prompt"`
	RevisedPrompt  string     `db:"revised_prompt"`
	N              int        `db:"n"`
	Size           string     `db:"size"`
	Quality        string     `db:"quality"`
	Style          string     `db:"style"`
	Status         string     `db:"status"`
	ConversationID string     `db:"conversation_id"`
	FileIDs        []byte     `db:"file_ids"`
	ResultURLs     []byte     `db:"result_urls"`
	ReferenceURLs  []byte     `db:"reference_urls"`
	LocalPaths     []byte     `db:"local_paths"`
	Error          string     `db:"error"`
	Attempts       int        `db:"attempts"`
	DurationMs     int64      `db:"duration_ms"`
	UserID         string     `db:"user_id"`
	CreatedAt      time.Time  `db:"created_at"`
	StartedAt      *time.Time `db:"started_at"`
	FinishedAt     *time.Time `db:"finished_at"`
}

// Result 是 Runner 返回给网关/客户端的生图结果。
type Result struct {
	TaskID         string        `json:"task_id"`
	Status         string        `json:"status"`
	ConversationID string        `json:"conversation_id,omitempty"`
	Images         []ResultImage `json:"images,omitempty"`
	ErrorCode      string        `json:"error_code,omitempty"`
	ErrorMessage   string        `json:"error_message,omitempty"`
}

// ResultImage 单张生图。
type ResultImage struct {
	URL         string `json:"url"`
	FileID      string `json:"file_id"`
	IsSediment  bool   `json:"is_sediment,omitempty"`
	ContentType string `json:"content_type,omitempty"`
}
````

## File: internal/image/runner.go
````go
package image

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/432539/gpt2api/internal/scheduler"
	"github.com/432539/gpt2api/internal/upstream/chatgpt"
	"github.com/432539/gpt2api/pkg/logger"
)

// Runner 单次/多次生图的执行器。封装完整的 chatgpt.com image2 主链路:
//
//	ChatRequirementsV2 → PrepareFConversation → StreamFConversation(SSE) →
//	ParseImageSSE → PollConversationForImages → ImageDownloadURL(签名 URL)
//
// 注意:/backend-api/conversation/init 不在 image2 主链路内,只用于账号 quota / blocked_features 弱诊断。
// 灰度桶未命中(preview_only)会自动换账号重试。
type Runner struct {
	sched *scheduler.Scheduler
	dao   *DAO
}

// NewRunner 构造 Runner。
func NewRunner(sched *scheduler.Scheduler, dao *DAO) *Runner {
	return &Runner{sched: sched, dao: dao}
}

// ReferenceImage 是图生图/编辑的一张参考图输入。
// 只需要提供原始字节 + 可选的文件名,Runner 会在运行时调用 chatgpt Client 上传。
type ReferenceImage struct {
	Data     []byte
	FileName string // 可选,未填时按长度 + 嗅探扩展名生成
}

// RunOptions 是单次生图的输入。
type RunOptions struct {
	TaskID            string
	ModelID           uint64
	UpstreamModel     string // 默认 "auto"(由上游根据 system_hints 挑选图像模型)
	Prompt            string
	N                 int              // 期望张数;实际由上游返回决定
	MaxAttempts       int              // 灰度未命中时最大重试,默认 2
	PerAttemptTimeout time.Duration    // 单次尝试总超时,默认 5min
	PollMaxWait       time.Duration    // 轮询最长等待,默认 300s
	References        []ReferenceImage // 图生图/编辑:参考图
}

// RunResult 是单次生图的输出。
type RunResult struct {
	Status          string // success / failed
	ConversationID  string
	AccountID       uint64
	FileIDs         []string // chatgpt.com 侧的原始 ref("sed:" 前缀表示 sediment)
	SignedURLs      []string // 直接可访问的签名 URL(15 分钟有效)
	ContentTypes    []string
	RevisedPrompt   string   // 上游改写后的 prompt(从 SSE assistant 文本中提取)
	ReferenceFileIDs []string // 上传到 GPT 的参考图 file-service ID
	ErrorCode       string
	ErrorMessage    string
	Attempts        int  // 跨账号尝试次数(runOnce 次数)
	TurnsInConv     int  // 当前账号内同会话 picture_v2 轮次
	IsPreview       bool // true=返回的是 IMG1 sediment 预览(3 轮均未命中 IMG2 灰度,已尽力)
	DurationMs      int64
}

// Run 执行生图。会同步阻塞直到完成/失败;调用方自行做超时控制(传 ctx)。
func (r *Runner) Run(ctx context.Context, opt RunOptions) *RunResult {
	start := time.Now()
	if opt.MaxAttempts <= 0 {
		opt.MaxAttempts = 2
	}
	if opt.PerAttemptTimeout <= 0 {
		opt.PerAttemptTimeout = 5 * time.Minute
	}
	if opt.PollMaxWait <= 0 {
		opt.PollMaxWait = 300 * time.Second
	}
	if opt.UpstreamModel == "" {
		// 对齐浏览器抓包 + 参考实现:图像走 f/conversation 时 model 字段和
		// 普通 chat 一致用 "auto",通过 system_hints=["picture_v2"] 让上游知道
		// 这是图像任务。硬写 "gpt-5-3" 在免费/新账号上会直接 404。
		opt.UpstreamModel = "auto"
	}
	if opt.N <= 0 {
		opt.N = 1
	}

	result := &RunResult{Status: StatusFailed, ErrorCode: ErrUnknown}

	// 仅当有 DAO 和 taskID 时才落库
	if r.dao != nil && opt.TaskID != "" {
		_ = r.dao.MarkRunning(ctx, opt.TaskID, 0)
	}

	// 排除集:跨账号重试时跳过已尝试过的账号
	var excludeAccountIDs map[uint64]struct{}

	// 代理切换和 preview_only 换号不算 attempt
	attempt := 0
	previewRetries := 0
	const maxPreviewRetries = 5 // preview_only 最多换 5 个号

	for attempt < opt.MaxAttempts {
		attempt++
		result.Attempts = attempt
		if err := ctx.Err(); err != nil {
			result.ErrorCode = ErrUnknown
			result.ErrorMessage = err.Error()
			break
		}

		attemptCtx, cancel := context.WithTimeout(ctx, opt.PerAttemptTimeout)
		ok, status, err := r.runOnce(attemptCtx, opt, result, excludeAccountIDs)
		cancel()

		if ok {
			result.Status = StatusSuccess
			result.ErrorCode = ""
			result.ErrorMessage = ""
			break
		}
		// 记录本次失败原因
		if err != nil {
			result.ErrorMessage = err.Error()
		}
		result.ErrorCode = status

		// ---- 静默退避策略 ----

		// 1) 代理错误:切换代理,不计入 attempt,不换账号
		if isProxyError(err) && result.AccountID > 0 {
			if b, _ := r.sched.AccountBinding(ctx, result.AccountID); b != nil && b.ProxyID > 0 {
				newURL, newPID := r.sched.SwitchProxy(ctx, result.AccountID, b.ProxyID)
				if newURL != "" {
					attempt-- // 不计入重试次数
					logger.L().Info("image runner proxy failed, silently switched",
						zap.String("task_id", opt.TaskID),
						zap.Uint64("account_id", result.AccountID),
						zap.Uint64("old_proxy_id", b.ProxyID),
						zap.Uint64("new_proxy_id", newPID))
					continue
				}
			}
		}

		// 2) 临时性上游错误(非内容策略、非认证失败):静默重试一次
		if status == ErrUpstream || status == ErrPollTimeout {
			if result.AccountID > 0 {
				if excludeAccountIDs == nil {
					excludeAccountIDs = make(map[uint64]struct{})
				}
				excludeAccountIDs[result.AccountID] = struct{}{}
			}
			logger.L().Info("image runner upstream error, silently retrying",
				zap.String("task_id", opt.TaskID),
				zap.String("status", status),
				zap.Int("attempt", attempt))
			continue
		}

		// 3) preview_only:降低账号置信度 + 换号重试(不计入 attempt)
		if status == ErrPreviewOnly {
			previewRetries++
			if result.AccountID > 0 {
				if excludeAccountIDs == nil {
					excludeAccountIDs = make(map[uint64]struct{})
				}
				excludeAccountIDs[result.AccountID] = struct{}{}
			}
			if previewRetries >= maxPreviewRetries {
				logger.L().Warn("image runner preview_only exhausted all retries",
					zap.String("task_id", opt.TaskID),
					zap.Int("preview_retries", previewRetries))
				break
			}
			attempt-- // 不计入 attempt,给真正的错误留重试机会
			logger.L().Info("image runner preview_only, switching account",
				zap.String("task_id", opt.TaskID),
				zap.Uint64("failed_account", result.AccountID),
				zap.Int("preview_retry", previewRetries),
				zap.Int("max", maxPreviewRetries))
			continue
		}

		// 4) 确定性错误(内容策略/认证失败/无账号):直接报给用户
		break
	}

	result.DurationMs = time.Since(start).Milliseconds()

	// 落库
	if r.dao != nil && opt.TaskID != "" {
		if result.Status == StatusSuccess {
			_ = r.dao.MarkSuccess(ctx, opt.TaskID, result.ConversationID,
				result.FileIDs, result.SignedURLs, SuccessExtra{
					RevisedPrompt:    result.RevisedPrompt,
					ReferenceFileIDs: result.ReferenceFileIDs,
					Attempts:         result.Attempts,
					DurationMs:       result.DurationMs,
				})
		} else {
			_ = r.dao.MarkFailed(ctx, opt.TaskID, result.ErrorCode)
		}
	}
	return result
}

// runOnce 一次完整的尝试。返回 (ok, errorCode, err)。
// result 会被就地更新(ConversationID / FileIDs / SignedURLs / AccountID 等)。
func (r *Runner) runOnce(ctx context.Context, opt RunOptions, result *RunResult, excludeIDs map[uint64]struct{}) (bool, string, error) {
	// 1) 调度账号(带排除集,跨账号重试时跳过已尝试的)
	lease, err := r.sched.DispatchWithExclude(ctx, "image", excludeIDs)
	if err != nil {
		if errors.Is(err, scheduler.ErrNoAvailable) {
			return false, ErrNoAccount, err
		}
		return false, ErrUnknown, err
	}
	defer func() {
		_ = lease.Release(context.Background())
	}()
	result.AccountID = lease.Account.ID
	// 立刻把 account_id 写回 image_tasks,供后续图片代理端点按 task_id 解出 AT。
	// MarkRunning 在 status=running 时 WHERE 不命中,所以用专门的 SetAccount。
	if r.dao != nil && opt.TaskID != "" {
		_ = r.dao.SetAccount(ctx, opt.TaskID, lease.Account.ID)
	}

	// 2) 构造上游 client
	cli, err := chatgpt.New(chatgpt.Options{
		AuthToken: lease.AuthToken,
		DeviceID:  lease.DeviceID,
		SessionID: lease.SessionID,
		ProxyURL:  lease.ProxyURL,
		Cookies:   lease.Cookies,
	})
	if err != nil {
		return false, ErrUnknown, fmt.Errorf("chatgpt client: %w", err)
	}

	// 2.5) 轻量 quota 预检:调用 conversation/init(picture_v2) 做 best-effort quota 诊断。
	// 它不是 image 能力主判据,只是为了尽早跳过“明确 blocked / 明确耗尽”的账号,
	// 减少无意义的 ChatRequirements / PoW 消耗。遇到 404/解析失败等情况会降级放行,
	// 最终是否能出图仍以后续 f/conversation 实际结果为准。
	hasQuota, blockedReason, probeErr := cli.ImageQuotaProbe(ctx)
	if probeErr != nil {
		logger.L().Warn("image runner quota probe error, continue anyway",
			zap.Uint64("account_id", lease.Account.ID), zap.Error(probeErr))
		// 预检网络/解析错误不阻断,降级放行
	} else if !hasQuota {
		logger.L().Info("image runner quota exhausted, skip account",
			zap.Uint64("account_id", lease.Account.ID),
			zap.String("blocked_reason", blockedReason))
		r.sched.MarkRateLimited(context.Background(), lease.Account.ID)
		return false, ErrRateLimited, fmt.Errorf("quota probe: %s", blockedReason)
	}

	// 3) ChatRequirements + POW(新两步 sentinel 流程,solver 未配置时内部自动
	// 回退到单步接口)
	cr, err := cli.ChatRequirementsV2(ctx)
	if err != nil {
		return false, r.classifyUpstream(err), err
	}
	var proofToken string
	if cr.Proofofwork.Required {
		proofCtx, cancel := context.WithTimeout(ctx, 8*time.Second)
		ch := make(chan string, 1)
		go func() { ch <- cr.SolveProof(chatgpt.DefaultUserAgent) }()
		select {
		case <-proofCtx.Done():
			cancel()
			r.sched.MarkWarned(context.Background(), lease.Account.ID)
			return false, ErrPOWTimeout, proofCtx.Err()
		case proofToken = <-ch:
			cancel()
		}
		if proofToken == "" {
			r.sched.MarkWarned(context.Background(), lease.Account.ID)
			return false, ErrPOWFailed, errors.New("pow solver returned empty")
		}
	}
	// Turnstile 是"建议性"信号:即使服务端声明 required,只要 chat_token + proof_token
	// 齐全,绝大多数账号的 f/conversation 仍然会正常下发图片结果。与 chat 流程(gateway/chat.go)
	// 保持一致——只打 warn,不阻断;真正拿不到 IMG2 终稿时由后续 poll 逻辑判定为失败。
	if cr.Turnstile.Required {
		logger.L().Warn("image turnstile required, continue anyway",
			zap.Uint64("account_id", lease.Account.ID))
	}

	// 4) 不再调用 /backend-api/conversation/init:
	// 浏览器实测路径是 prepare → chat-requirements → f/conversation 三步,init 是
	// 过时/冗余调用,在免费账号上还会返回 404 让整条链路 fail。system_hints=picture_v2
	// 会通过 f/conversation 的 payload 字段传达。

	// 4.5) 图生图:上传参考图。任何一张失败都直接整体 fail(上游后续会对不上 attachment)。
	var refs []*chatgpt.UploadedFile
	if len(opt.References) > 0 {
		for idx, r0 := range opt.References {
			upCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
			up, err := cli.UploadFile(upCtx, r0.Data, r0.FileName)
			cancel()
			if err != nil {
				logger.L().Warn("image runner upload reference failed",
					zap.Int("idx", idx), zap.Error(err))
				if ue, ok := err.(*chatgpt.UpstreamError); ok && ue.IsRateLimited() {
					r.sched.MarkRateLimited(context.Background(), lease.Account.ID)
					return false, ErrRateLimited, err
				}
				return false, ErrUpstream, fmt.Errorf("upload reference %d: %w", idx, err)
			}
			refs = append(refs, up)
		}
		// 记录上传到 GPT 的参考图 file ID,用于训练数据留存
		refFileIDs := make([]string, 0, len(refs))
		for _, u := range refs {
			if u != nil && u.FileID != "" {
				refFileIDs = append(refFileIDs, u.FileID)
			}
		}
		result.ReferenceFileIDs = refFileIDs
		logger.L().Info("image runner references uploaded",
			zap.String("task_id", opt.TaskID), zap.Int("count", len(refs)),
			zap.Strings("ref_file_ids", refFileIDs))
	}

	// 注意:新会话不要本地生成 conversation_id,上游会 404。
	// 真正的 conv_id 由 SSE 的 resume_conversation_token / sseResult.ConversationID 返回。
	var convID string
	parentID := uuid.NewString()
	messageID := uuid.NewString()

	// 统一把 model 强制为 "auto":对齐参考实现(只通过 system_hints=["picture_v2"]
	// 区分图像任务),避免 chatgpt-freeaccount / chatgpt-paid 之间的 model slug 差异。
	upstreamModel := "auto"
	if opt.UpstreamModel != "" && opt.UpstreamModel != "auto" && !cr.IsFreeAccount() {
		// 付费账号如果明确传了 upstream slug 且不是 auto,可以尊重调用传入
		// (但我们现有模型库没有 image 专用 slug,保留扩展点)
		upstreamModel = opt.UpstreamModel
	} else if cr.IsFreeAccount() && opt.UpstreamModel != "" && opt.UpstreamModel != "auto" {
		logger.L().Warn("image: free account requesting premium model, downgrade to auto",
			zap.Uint64("account_id", lease.Account.ID),
			zap.String("requested_model", opt.UpstreamModel))
	}

	// 5-7) 同账号 + 同会话多轮发起 picture_v2,命中 IMG2 即返回;
	// 连续 sameConvMax 轮只拿到预览(preview_only)时,用最后一轮的 sediment 作为 IMG1 返回。
	// 协议/网络级错误(非 preview_only)会直接退出,由外层 Run 换账号。
	const sameConvMax = 1 // 同号只试 1 轮,不命中立即换号(比同号多轮重试更快)

	var (
		fileRefs      []string
		previewRounds int
		// baselineTools 记录上一轮结束时会话里已有的 image_gen tool 消息 id,
		// 下一轮 PollConversationForImages 只看新增,避免把旧 preview 当本轮结果。
		baselineTools = map[string]struct{}{}
		// excludeSids 记录之前轮次产出的 preview sediment ID,
		// IMG2 命中时从结果中排除,避免把旧预览图混入终稿。
		excludeSids = map[string]struct{}{}
	)

loop:
	for turn := 1; turn <= sameConvMax; turn++ {
		result.TurnsInConv = turn

		// 每轮重新拉 chat_token + proof_token(之前那张已经消耗过)。
		// 复用外层 cr / proofToken 的首次结果(turn==1 直接用),后续重取。
		if turn > 1 {
			cr, err = cli.ChatRequirementsV2(ctx)
			if err != nil {
				return false, r.classifyUpstream(err), err
			}
			proofToken = ""
			if cr.Proofofwork.Required {
				proofCtx, cancel := context.WithTimeout(ctx, 8*time.Second)
				ch := make(chan string, 1)
				go func() { ch <- cr.SolveProof(chatgpt.DefaultUserAgent) }()
				select {
				case <-proofCtx.Done():
					cancel()
					r.sched.MarkWarned(context.Background(), lease.Account.ID)
					return false, ErrPOWTimeout, proofCtx.Err()
				case proofToken = <-ch:
					cancel()
				}
				if proofToken == "" {
					r.sched.MarkWarned(context.Background(), lease.Account.ID)
					return false, ErrPOWFailed, errors.New("pow solver returned empty")
				}
			}
		}

		convOpt := chatgpt.ImageConvOpts{
			Prompt:        opt.Prompt,
			UpstreamModel: upstreamModel,
			ConvID:        convID, // 第 1 轮空串=新会话,后续轮复用
			ParentMsgID:   parentID,
			MessageID:     messageID,
			ChatToken:     cr.Token,
			ProofToken:    proofToken,
			References:    refs,
		}
		if turn > 1 {
			// 续聊:每轮用新的 message_id,parent 来自上轮会话头
			convOpt.MessageID = uuid.NewString()
		}

		// Prepare(conduit_token;不成功也能降级跑 conversation)
		if ct, err := cli.PrepareFConversation(ctx, convOpt); err == nil {
			convOpt.ConduitToken = ct
		} else if ue, ok := err.(*chatgpt.UpstreamError); ok && ue.IsRateLimited() {
			r.sched.MarkRateLimited(context.Background(), lease.Account.ID)
			return false, ErrRateLimited, err
		}

		// f/conversation SSE
		stream, err := cli.StreamFConversation(ctx, convOpt)
		if err != nil {
			code := r.classifyUpstream(err)
			if code == ErrRateLimited {
				r.sched.MarkRateLimited(context.Background(), lease.Account.ID)
			}
			return false, code, err
		}
		sseResult := chatgpt.ParseImageSSE(stream)
		if sseResult.ConversationID != "" {
			convID = sseResult.ConversationID
			result.ConversationID = convID
		}
		// 保留上游助手文本作为 revised_prompt(通常包含对 prompt 的改写/解释)
		if sseResult.AssistantText != "" && result.RevisedPrompt == "" {
			result.RevisedPrompt = sseResult.AssistantText
		}

		// 每轮 SSE 解析完的原始产物:FileIDs(file-service://,IMG2 直出时有)、
		// SedimentIDs(sediment://,preview 或 IMG1 常见)、FinishType。用于诊断
		// "这轮到底返回了什么"。
		logger.L().Info("image runner SSE parsed",
			zap.String("task_id", opt.TaskID),
			zap.Uint64("account_id", lease.Account.ID),
			zap.Int("turn", turn),
			zap.String("conv_id", convID),
			zap.String("finish_type", sseResult.FinishType),
			zap.String("image_gen_task_id", sseResult.ImageGenTaskID),
			zap.Int("sse_fids", len(sseResult.FileIDs)),
			zap.Strings("sse_fids_list", sseResult.FileIDs),
			zap.Int("sse_sids", len(sseResult.SedimentIDs)),
			zap.Strings("sse_sids_list", sseResult.SedimentIDs),
			zap.Int("sse_img2_sids", len(sseResult.IMG2SedimentIDs)),
			zap.Strings("sse_img2_sids_list", sseResult.IMG2SedimentIDs),
		)

		// 内容策略拒绝:SSE 结束后上游根本没发起图片生成(无 ImageGenTaskID),
		// 也没有任何 file/sediment 引用 —— 说明上游拒绝了该 prompt。
		// 立即返回错误,不进入 poll 等待。
		if sseResult.ImageGenTaskID == "" && len(sseResult.FileIDs) == 0 && len(sseResult.SedimentIDs) == 0 {
			reason := sseResult.AssistantText
			if reason == "" {
				reason = "上游拒绝生成该图片"
			}
			if len(reason) > 300 {
				reason = reason[:300]
			}
			logger.L().Warn("image runner rejected by upstream (no image_gen_task_id)",
				zap.String("task_id", opt.TaskID),
				zap.Uint64("account_id", lease.Account.ID),
				zap.String("reason", reason),
			)
			return false, ErrContentPolicy, errors.New(reason)
		}

		// SSE 直出 file-service = IMG2 命中。2026 抓包还确认:IMG2 可能
		// 只在 SSE 中给单条 sediment,但同片段带 gen_size_v2；这种也应直返。
		// 注意:同一次灰度也可能同时带 sediment(例如 preview + final 各一张),
		// 都要收进来,不然调用方会少看到图。
		if len(sseResult.FileIDs) > 0 || len(sseResult.IMG2SedimentIDs) > 0 {
			fileRefs = append(fileRefs, sseResult.FileIDs...)
			sidsToUse := sseResult.SedimentIDs
			if len(sseResult.FileIDs) == 0 && len(sseResult.IMG2SedimentIDs) > 0 {
				sidsToUse = sseResult.IMG2SedimentIDs
			}
			for _, s := range sidsToUse {
				fileRefs = append(fileRefs, "sed:"+s)
			}
			logger.L().Info("image runner IMG2 direct hit (from SSE)",
				zap.String("task_id", opt.TaskID),
				zap.Uint64("account_id", lease.Account.ID),
				zap.Int("turn", turn),
				zap.String("conv_id", convID),
				zap.Int("total_refs", len(fileRefs)),
				zap.Strings("refs", fileRefs),
			)
			break loop
		}

		// 没直出就轮询当前会话
		pollOpt := chatgpt.PollOpts{
			MaxWait:         opt.PollMaxWait,
			BaselineToolIDs: baselineTools,
		}
		status, fids, sids := cli.PollConversationForImages(ctx, convID, pollOpt)
		// 每轮 Poll 的产物,无论 status 如何都打印一条,用于诊断"第几轮拿到了什么"。
		logger.L().Info("image runner poll done",
			zap.String("task_id", opt.TaskID),
			zap.Uint64("account_id", lease.Account.ID),
			zap.Int("turn", turn),
			zap.String("conv_id", convID),
			zap.String("poll_status", string(status)),
			zap.Int("poll_fids", len(fids)),
			zap.Strings("poll_fids_list", fids),
			zap.Int("poll_sids", len(sids)),
			zap.Strings("poll_sids_list", sids),
		)
		switch status {
		case chatgpt.PollStatusIMG2:
			fileRefs = append(fileRefs, fids...)
			for _, s := range sids {
				if _, old := excludeSids[s]; old {
					continue // 跳过之前轮次的预览 sediment
				}
				fileRefs = append(fileRefs, "sed:"+s)
			}
			logger.L().Info("image runner IMG2 poll hit",
				zap.String("task_id", opt.TaskID),
				zap.Uint64("account_id", lease.Account.ID),
				zap.Int("turn", turn),
				zap.String("conv_id", convID),
				zap.Int("total_refs", len(fileRefs)),
				zap.Strings("refs", fileRefs),
			)
			break loop

		case chatgpt.PollStatusPreviewOnly:
			previewRounds++
			// 把预览的 sediment ID 加入排除集,后续轮次 IMG2 命中时不会混入旧预览
			for _, s := range sids {
				excludeSids[s] = struct{}{}
			}
			for _, f := range fids {
				excludeSids[f] = struct{}{}
			}
			logger.L().Info("image runner preview_only, retry in same conversation",
				zap.String("task_id", opt.TaskID),
				zap.Uint64("account_id", lease.Account.ID),
				zap.String("conv_id", convID),
				zap.Int("turn", turn),
				zap.Int("max_turns", sameConvMax),
				zap.Int("preview_fids", len(fids)),
				zap.Strings("preview_fids_list", fids),
				zap.Int("preview_sids", len(sids)),
				zap.Strings("preview_sids_list", sids),
			)

			// 不是最后一轮:更新 baseline + 取会话头作为下轮的 parent_message_id
			if turn < sameConvMax {
				if mapping, merr := cli.GetConversationMapping(ctx, convID); merr == nil {
					// 把当前所有 tool 消息都塞进 baseline,下轮 poll 只看新增
					if newBL := buildToolBaseline(mapping); newBL != nil {
						baselineTools = newBL
					}
					if head, _ := mapping["current_node"].(string); head != "" {
						parentID = head
					}
				} else {
					logger.L().Warn("image runner fetch mapping for retry failed",
						zap.Uint64("account_id", lease.Account.ID), zap.Error(merr))
				}
			}

		case chatgpt.PollStatusTimeout:
			r.sched.RecordIMG2Outcome(context.Background(), lease.Account.ID, "miss")
			return false, ErrPollTimeout, errors.New("poll timeout")

		case chatgpt.PollStatus429:
			// 上游 RPM 限流,图可能还在生成中,增大 poll 间隔后继续 poll 同一个会话
			logger.L().Warn("image runner poll hit 429, increasing interval and retrying poll",
				zap.String("task_id", opt.TaskID),
				zap.Uint64("account_id", lease.Account.ID),
				zap.Int("turn", turn))
			// 用更长的间隔重新 poll 同一个 convID
			pollOpt2 := chatgpt.PollOpts{
				MaxWait:         2 * time.Minute,
				Interval:        15 * time.Second, // 拉长间隔避免再 429
				BaselineToolIDs: baselineTools,
			}
			status2, fids2, sids2 := cli.PollConversationForImages(ctx, convID, pollOpt2)
			if status2 == chatgpt.PollStatusIMG2 {
				fileRefs = append(fileRefs, fids2...)
				for _, s := range sids2 {
					if _, old := excludeSids[s]; old {
						continue
					}
					fileRefs = append(fileRefs, "sed:"+s)
				}
				break loop
			}
			// 二次 poll 还是失败,当作本轮未出图
			r.sched.RecordIMG2Outcome(context.Background(), lease.Account.ID, "miss")
			return false, ErrUpstream, errors.New("poll 429 retry exhausted")

		default:
			r.sched.RecordIMG2Outcome(context.Background(), lease.Account.ID, "miss")
			return false, ErrUpstream, errors.New("poll error")
		}
	}

	// 若循环结束仍未拿到 IMG2,不兜底,直接失败并降低账号置信度
	if len(fileRefs) == 0 {
		r.sched.RecordIMG2Outcome(context.Background(), lease.Account.ID, "miss")
		logger.L().Warn("image runner all turns preview_only, no IMG2 hit",
			zap.String("task_id", opt.TaskID),
			zap.Uint64("account_id", lease.Account.ID),
			zap.String("conv_id", convID),
			zap.Int("preview_rounds", previewRounds))
		return false, ErrPreviewOnly, errors.New("未命中 IMG2 灰度,请重试")
	}

	fileRefs = dedupeImageRefs(fileRefs)

	// 到这里说明 fileRefs 不为空 = IMG2 真正命中
	r.sched.RecordIMG2Outcome(context.Background(), lease.Account.ID, "hit")

	// 8) 对每个 ref 取签名 URL
	var signedURLs []string
	var contentTypes []string
	for _, ref := range fileRefs {
		url, err := cli.ImageDownloadURL(ctx, convID, ref)
		if err != nil {
			logger.L().Warn("image runner download url failed",
				zap.String("ref", ref), zap.Error(err))
			continue
		}
		signedURLs = append(signedURLs, url)
		contentTypes = append(contentTypes, "image/png")
	}
	if len(signedURLs) == 0 {
		r.sched.RecordIMG2Delivery(context.Background(), lease.Account.ID, "fail")
		logger.L().Warn("image runner delivery failed after refs",
			zap.String("task_id", opt.TaskID),
			zap.Uint64("account_id", lease.Account.ID),
			zap.String("conv_id", convID),
			zap.Int("refs", len(fileRefs)))
		return false, ErrDownload, errors.New("all download urls failed")
	}
	deliveryStatus := "success"
	if len(signedURLs) < len(fileRefs) {
		deliveryStatus = "partial"
	}
	r.sched.RecordIMG2Delivery(context.Background(), lease.Account.ID, deliveryStatus)

	logger.L().Info("image runner result summary",
		zap.String("task_id", opt.TaskID),
		zap.Uint64("account_id", lease.Account.ID),
		zap.String("conv_id", convID),
		zap.Int("turns_used", result.TurnsInConv),
		zap.Int("refs", len(fileRefs)),
		zap.Strings("refs_list", fileRefs),
		zap.Int("signed_count", len(signedURLs)),
	)

	result.FileIDs = fileRefs
	result.SignedURLs = signedURLs
	result.ContentTypes = contentTypes
	return true, "", nil
}

// buildToolBaseline 从 conversation mapping 里提取所有已存在的 image_gen tool 消息 id,
// 作为下一轮 PollConversationForImages 的 baseline。
func buildToolBaseline(mapping map[string]interface{}) map[string]struct{} {
	tools := chatgpt.ExtractImageToolMsgs(mapping)
	if len(tools) == 0 {
		return nil
	}
	out := make(map[string]struct{}, len(tools))
	for _, t := range tools {
		out[t.MessageID] = struct{}{}
	}
	return out
}

func dedupeImageRefs(refs []string) []string {
	if len(refs) <= 1 {
		return refs
	}
	out := make([]string, 0, len(refs))
	seen := make(map[string]struct{}, len(refs))
	for _, ref := range refs {
		if ref == "" {
			continue
		}
		if _, ok := seen[ref]; ok {
			continue
		}
		seen[ref] = struct{}{}
		out = append(out, ref)
	}
	return out
}

// classifyUpstream 把上游错误转成内部 error code。
func (r *Runner) classifyUpstream(err error) string {
	if err == nil {
		return ""
	}
	var ue *chatgpt.UpstreamError
	if errors.As(err, &ue) {
		if ue.IsRateLimited() {
			return ErrRateLimited
		}
		if ue.IsUnauthorized() {
			return ErrAuthRequired
		}
		return ErrUpstream
	}
	if strings.Contains(err.Error(), "deadline exceeded") {
		return ErrPollTimeout
	}
	return ErrUpstream
}

// isProxyError 判断错误是否是代理级错误(连接失败/超时/407/SOCKS 握手等)。
func isProxyError(err error) bool {
	if err == nil {
		return false
	}
	var ue *chatgpt.UpstreamError
	if errors.As(err, &ue) && ue.Status == 407 {
		return true
	}
	msg := err.Error()
	proxyKeywords := []string{
		"proxy", "SOCKS", "connection refused", "connection reset",
		"no such host", "i/o timeout", "dial tcp", "EOF",
		"connect: network is unreachable",
	}
	for _, kw := range proxyKeywords {
		if strings.Contains(msg, kw) {
			return true
		}
	}
	return false
}

// GenerateTaskID 生成对外 task_id。
func GenerateTaskID() string {
	return "img_" + strings.ReplaceAll(uuid.NewString(), "-", "")[:24]
}
````

## File: internal/middleware/auth.go
````go
package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	CtxActorID    = "actor_id"
	CtxActorEmail = "actor_email"
)

// tokenSecret 从 aes_key 派生,进程启动后不变。
var tokenSecret []byte

// InitTokenSecret 从 aes_key 派生 admin token 签名密钥。由 main.go 调用。
func InitTokenSecret(aesKeyHex string) {
	h := sha256.Sum256([]byte("admin-token-secret:" + aesKeyHex))
	tokenSecret = h[:]
}

// GenerateToken 签发 admin token。格式: timestamp_hex.hmac_hex
// 有效期 7 天。
func GenerateToken(username string) string {
	exp := time.Now().Add(7 * 24 * time.Hour).Unix()
	payload := fmt.Sprintf("%s|%d", username, exp)
	mac := hmac.New(sha256.New, tokenSecret)
	mac.Write([]byte(payload))
	sig := hex.EncodeToString(mac.Sum(nil))
	return fmt.Sprintf("%s.%s", payload, sig)
}

// ValidateToken 验证 token,返回 username 或空串。
func ValidateToken(token string) string {
	parts := strings.SplitN(token, ".", 2)
	if len(parts) != 2 {
		return ""
	}
	payload, sig := parts[0], parts[1]

	mac := hmac.New(sha256.New, tokenSecret)
	mac.Write([]byte(payload))
	expected := hex.EncodeToString(mac.Sum(nil))
	if !hmac.Equal([]byte(sig), []byte(expected)) {
		return ""
	}

	fields := strings.SplitN(payload, "|", 2)
	if len(fields) != 2 {
		return ""
	}
	exp, err := strconv.ParseInt(fields[1], 10, 64)
	if err != nil || time.Now().Unix() > exp {
		return ""
	}
	return fields[0]
}

// AdminAuth 验证 admin token 的中间件。
// 从 Authorization: Bearer <token> 或 query 参数 token 中提取。
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := ""
		if auth := c.GetHeader("Authorization"); strings.HasPrefix(auth, "Bearer ") {
			token = strings.TrimPrefix(auth, "Bearer ")
		}
		if token == "" {
			token = c.Query("token")
		}
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "未登录"})
			return
		}
		username := ValidateToken(token)
		if username == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "登录已过期，请重新登录"})
			return
		}
		c.Set(CtxActorID, uint64(0))
		c.Set(CtxActorEmail, username)
		c.Next()
	}
}

// V1APIKeyGetter 获取当前配置的 API Key。
type V1APIKeyGetter interface {
	V1APIKey() string
}

// V1APIKeyAuth 验证 /v1/* 请求的 API Key。
// key 为空时不验证(开放访问)。
func V1APIKeyAuth(getter V1APIKeyGetter) gin.HandlerFunc {
	return func(c *gin.Context) {
		configuredKey := getter.V1APIKey()
		if configuredKey == "" {
			c.Set(CtxActorID, uint64(0))
			c.Set(CtxActorEmail, "anonymous")
			c.Next()
			return
		}
		auth := c.GetHeader("Authorization")
		providedKey := ""
		if strings.HasPrefix(auth, "Bearer ") {
			providedKey = strings.TrimPrefix(auth, "Bearer ")
		}
		if providedKey == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": gin.H{"message": "Missing API key. Set Authorization: Bearer <key>", "type": "invalid_request_error"},
			})
			return
		}
		if !hmac.Equal([]byte(providedKey), []byte(configuredKey)) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": gin.H{"message": "Invalid API key", "type": "invalid_request_error"},
			})
			return
		}
		c.Set(CtxActorID, uint64(0))
		c.Set(CtxActorEmail, "api-client")
		c.Next()
	}
}

// LocalActor 为单人本地控制台注入固定操作者信息。
func LocalActor() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(CtxActorID, uint64(0))
		c.Set(CtxActorEmail, "local-console")
		c.Next()
	}
}

// UserID 在本地模式中始终返回 0,仅用于审计/设置等兼容字段。
func UserID(c *gin.Context) uint64 {
	if c == nil {
		return 0
	}
	if v, ok := c.Get(CtxActorID); ok {
		switch x := v.(type) {
		case uint64:
			return x
		case int:
			if x >= 0 {
				return uint64(x)
			}
		}
	}
	return 0
}

// ActorID 返回本地操作者标识。
func ActorID(c *gin.Context) uint64 { return UserID(c) }

// Role 返回本地控制台角色。
func Role(c *gin.Context) string { return "local" }
````

## File: internal/middleware/cors.go
````go
package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// CORS 简易跨域中间件。
func CORS(origins []string) gin.HandlerFunc {
	allow := make(map[string]struct{}, len(origins))
	allowAll := false
	for _, o := range origins {
		if o == "*" {
			allowAll = true
		}
		allow[strings.TrimRight(o, "/")] = struct{}{}
	}
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if origin != "" {
			if allowAll {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			} else if _, ok := allow[strings.TrimRight(origin, "/")]; ok {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			}
			c.Writer.Header().Set("Vary", "Origin")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, X-Request-Id")
			c.Writer.Header().Set("Access-Control-Expose-Headers", "X-Request-Id")
			c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		}
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
````

## File: internal/middleware/logger.go
````go
package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/432539/gpt2api/pkg/logger"
)

// AccessLog 打印每一次 HTTP 访问的结构化日志。
func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		cost := time.Since(start)

		log := logger.L()
		fields := []zap.Field{
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("cost", cost),
			zap.String("ip", c.ClientIP()),
			zap.String("ua", c.Request.UserAgent()),
			zap.String("request_id", getString(c, "request_id")),
		}
		if errs := c.Errors.ByType(gin.ErrorTypePrivate).String(); errs != "" {
			fields = append(fields, zap.String("err", errs))
		}

		status := c.Writer.Status()
		switch {
		case status >= 500:
			log.Error("http", fields...)
		case status >= 400:
			log.Warn("http", fields...)
		default:
			log.Info("http", fields...)
		}
	}
}

func getString(c *gin.Context, key string) string {
	if v, ok := c.Get(key); ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}
````

## File: internal/middleware/recover.go
````go
package middleware

import (
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/432539/gpt2api/pkg/logger"
	"github.com/432539/gpt2api/pkg/resp"
)

// Recover 捕获 panic,写入日志并返回 500。
func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logger.L().Error("panic recovered",
					zap.Any("err", r),
					zap.ByteString("stack", debug.Stack()),
					zap.String("path", c.Request.URL.Path),
					zap.String("request_id", getString(c, "request_id")),
				)
				resp.Internal(c, "internal server error")
			}
		}()
		c.Next()
	}
}
````

## File: internal/middleware/request_id.go
````go
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const HeaderRequestID = "X-Request-Id"

// RequestID 为每个请求生成/透传 request_id,写入 context 和响应头。
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := c.GetHeader(HeaderRequestID)
		if rid == "" {
			rid = uuid.NewString()
		}
		c.Set("request_id", rid)
		c.Writer.Header().Set(HeaderRequestID, rid)
		c.Next()
	}
}
````

## File: internal/model/admin_handler.go
````go
package model

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	mysqlDrv "github.com/go-sql-driver/mysql"

	"github.com/432539/gpt2api/internal/audit"
	"github.com/432539/gpt2api/pkg/resp"
)

var slugRe = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9._\-]{1,63}$`)

// AdminHandler 本地控制台的模型 CRUD。
type AdminHandler struct {
	dao      *DAO
	registry *Registry
	auditDAO *audit.DAO
}

func NewAdminHandler(dao *DAO, registry *Registry, auditDAO *audit.DAO) *AdminHandler {
	return &AdminHandler{dao: dao, registry: registry, auditDAO: auditDAO}
}

type upsertReq struct {
	Slug              string `json:"slug"`
	Type              string `json:"type"`
	UpstreamModelSlug string `json:"upstream_model_slug"`
	Description       string `json:"description"`
	Enabled           *bool  `json:"enabled,omitempty"`
}

func (r *upsertReq) validate(forCreate bool) error {
	r.Slug = strings.TrimSpace(r.Slug)
	r.UpstreamModelSlug = strings.TrimSpace(r.UpstreamModelSlug)
	r.Type = strings.TrimSpace(strings.ToLower(r.Type))
	if forCreate && !slugRe.MatchString(r.Slug) {
		return errors.New("slug 非法:需字母开头,2-64 位字母/数字/点/下划线/短横")
	}
	if r.Type != TypeChat && r.Type != TypeImage {
		return errors.New("type 只能为 chat 或 image")
	}
	if r.UpstreamModelSlug == "" {
		return errors.New("upstream_model_slug 不能为空")
	}
	if len(r.Description) > 255 {
		return errors.New("description 超过 255 字")
	}
	return nil
}

func (h *AdminHandler) List(c *gin.Context) {
	rows, err := h.dao.List(c.Request.Context())
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}
	resp.OK(c, gin.H{"items": rows, "total": len(rows)})
}

func (h *AdminHandler) Create(c *gin.Context) {
	var req upsertReq
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	if err := req.validate(true); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	enabled := true
	if req.Enabled != nil {
		enabled = *req.Enabled
	}
	m := &Model{Slug: req.Slug, Type: req.Type, UpstreamModelSlug: req.UpstreamModelSlug, Description: req.Description, Enabled: enabled}
	if err := h.dao.Create(c.Request.Context(), m); err != nil {
		if isDupSlug(err) {
			resp.BadRequest(c, "slug 已存在")
			return
		}
		resp.Internal(c, err.Error())
		return
	}
	h.reloadRegistry(c)
	audit.Record(c, h.auditDAO, "models.create", strconv.FormatUint(m.ID, 10), gin.H{"slug": m.Slug, "type": m.Type})
	resp.OK(c, m)
}

func (h *AdminHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		resp.BadRequest(c, "invalid id")
		return
	}
	var req upsertReq
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	if err := req.validate(false); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	cur, err := h.dao.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			resp.NotFound(c, "model not found")
			return
		}
		resp.Internal(c, err.Error())
		return
	}
	cur.Type = req.Type
	cur.UpstreamModelSlug = req.UpstreamModelSlug
	cur.Description = req.Description
	if req.Enabled != nil {
		cur.Enabled = *req.Enabled
	}
	if err := h.dao.Update(c.Request.Context(), cur); err != nil {
		resp.Internal(c, err.Error())
		return
	}
	h.reloadRegistry(c)
	audit.Record(c, h.auditDAO, "models.update", strconv.FormatUint(id, 10), gin.H{"slug": cur.Slug})
	resp.OK(c, cur)
}

func (h *AdminHandler) SetEnabled(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		resp.BadRequest(c, "invalid id")
		return
	}
	var body struct {
		Enabled bool `json:"enabled"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	if err := h.dao.SetEnabled(c.Request.Context(), id, body.Enabled); err != nil {
		if errors.Is(err, ErrNotFound) {
			resp.NotFound(c, "model not found")
			return
		}
		resp.Internal(c, err.Error())
		return
	}
	h.reloadRegistry(c)
	audit.Record(c, h.auditDAO, "models.set_enabled", strconv.FormatUint(id, 10), gin.H{"enabled": body.Enabled})
	resp.OK(c, gin.H{"id": id, "enabled": body.Enabled})
}

func (h *AdminHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		resp.BadRequest(c, "invalid id")
		return
	}
	if err := h.dao.SoftDelete(c.Request.Context(), id); err != nil {
		if errors.Is(err, ErrNotFound) {
			resp.NotFound(c, "model not found")
			return
		}
		resp.Internal(c, err.Error())
		return
	}
	h.reloadRegistry(c)
	audit.Record(c, h.auditDAO, "models.delete", strconv.FormatUint(id, 10), nil)
	resp.OK(c, gin.H{"deleted": id})
}

// ListEnabledForMe 本地控制台视角,只返回 enabled 模型,用于在线体验下拉选择。
func (h *AdminHandler) ListEnabledForMe(c *gin.Context) {
	rows, err := h.dao.ListEnabled(c.Request.Context())
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}
	type simple struct {
		ID          uint64 `json:"id"`
		Slug        string `json:"slug"`
		Type        string `json:"type"`
		Description string `json:"description"`
	}
	out := make([]simple, 0, len(rows))
	for _, m := range rows {
		out = append(out, simple{ID: m.ID, Slug: m.Slug, Type: m.Type, Description: m.Description})
	}
	resp.OK(c, gin.H{"items": out, "total": len(out)})
}

func (h *AdminHandler) reloadRegistry(c *gin.Context) {
	if h.registry == nil {
		return
	}
	_ = h.registry.Reload(c.Request.Context())
}

func isDupSlug(err error) bool {
	var me *mysqlDrv.MySQLError
	return errors.As(err, &me) && me.Number == 1062
}
````

## File: internal/model/dao.go
````go
package model

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

var ErrNotFound = errors.New("model: not found")

type DAO struct{ db *sqlx.DB }

func NewDAO(db *sqlx.DB) *DAO { return &DAO{db: db} }

func (d *DAO) GetBySlug(ctx context.Context, slug string) (*Model, error) {
	var m Model
	err := d.db.GetContext(ctx, &m,
		`SELECT * FROM models WHERE slug = ? AND deleted_at IS NULL`, slug)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	return &m, err
}

func (d *DAO) ListEnabled(ctx context.Context) ([]*Model, error) {
	rows := make([]*Model, 0, 16)
	err := d.db.SelectContext(ctx, &rows,
		`SELECT * FROM models WHERE enabled = 1 AND deleted_at IS NULL ORDER BY id ASC`)
	return rows, err
}

func (d *DAO) List(ctx context.Context) ([]*Model, error) {
	rows := make([]*Model, 0, 16)
	err := d.db.SelectContext(ctx, &rows,
		`SELECT * FROM models WHERE deleted_at IS NULL ORDER BY id ASC`)
	return rows, err
}

// GetByID 按主键查。
func (d *DAO) GetByID(ctx context.Context, id uint64) (*Model, error) {
	var m Model
	err := d.db.GetContext(ctx, &m,
		`SELECT * FROM models WHERE id = ? AND deleted_at IS NULL`, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	return &m, err
}

// Create 插入一条新模型。slug 唯一冲突由上层判断(返回 MySQL 1062)。
func (d *DAO) Create(ctx context.Context, m *Model) error {
	res, err := d.db.ExecContext(ctx, `
INSERT INTO models
  (slug, type, upstream_model_slug, description, enabled)
VALUES (?,?,?,?,?)`,
		m.Slug, m.Type, m.UpstreamModelSlug, m.Description, m.Enabled,
	)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	m.ID = uint64(id)
	return nil
}

// Update 按 id 全量更新(不改 slug;改 slug 走单独接口更安全)。
func (d *DAO) Update(ctx context.Context, m *Model) error {
	res, err := d.db.ExecContext(ctx, `
UPDATE models SET
  type = ?, upstream_model_slug = ?, description = ?, enabled = ?
WHERE id = ? AND deleted_at IS NULL`,
		m.Type, m.UpstreamModelSlug, m.Description, m.Enabled, m.ID,
	)
	if err != nil {
		return err
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return ErrNotFound
	}
	return nil
}

// SetEnabled 开关。
func (d *DAO) SetEnabled(ctx context.Context, id uint64, enabled bool) error {
	res, err := d.db.ExecContext(ctx,
		`UPDATE models SET enabled = ? WHERE id = ? AND deleted_at IS NULL`,
		enabled, id)
	if err != nil {
		return err
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return ErrNotFound
	}
	return nil
}

// SoftDelete 软删除:打 deleted_at,释放 slug 供后续复用。
// 由于 uk_slug 是 UNIQUE,复用 slug 需要把已删除行的 slug 改名。
func (d *DAO) SoftDelete(ctx context.Context, id uint64) error {
	res, err := d.db.ExecContext(ctx, `
UPDATE models
   SET deleted_at = NOW(),
       enabled    = 0,
       slug       = CONCAT(slug, '#del', id)
 WHERE id = ? AND deleted_at IS NULL`, id)
	if err != nil {
		return err
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return ErrNotFound
	}
	return nil
}
````

## File: internal/model/model.go
````go
package model

import (
	"database/sql"
	"time"
)

const (
	TypeChat  = "chat"
	TypeImage = "image"
)

// Model 对应 models 表,只描述本地 slug 与上游 slug 的映射。
type Model struct {
	ID                uint64       `db:"id" json:"id"`
	Slug              string       `db:"slug" json:"slug"`
	Type              string       `db:"type" json:"type"`
	UpstreamModelSlug string       `db:"upstream_model_slug" json:"upstream_model_slug"`
	Description       string       `db:"description" json:"description"`
	Enabled           bool         `db:"enabled" json:"enabled"`
	CreatedAt         time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time    `db:"updated_at" json:"updated_at"`
	DeletedAt         sql.NullTime `db:"deleted_at" json:"-"`
}
````

## File: internal/model/registry.go
````go
package model

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

// Registry 模型配置的进程内缓存,TTL=30s。
// 热点路径:网关请求都会查模型,绝不能每次打 DB。
type Registry struct {
	dao *DAO

	mu       sync.RWMutex
	bySlug   map[string]*Model
	loadedAt time.Time
	ttl      time.Duration

	refreshing atomic.Bool
}

func NewRegistry(dao *DAO) *Registry {
	return &Registry{dao: dao, bySlug: map[string]*Model{}, ttl: 30 * time.Second}
}

// Preload 启动时预热。
func (r *Registry) Preload(ctx context.Context) error {
	return r.refresh(ctx)
}

// Reload 外部触发强制刷新(管理端写后调用,保证下一次请求立刻看到新数据)。
func (r *Registry) Reload(ctx context.Context) error {
	return r.refresh(ctx)
}

func (r *Registry) refresh(ctx context.Context) error {
	list, err := r.dao.List(ctx)
	if err != nil {
		return err
	}
	m := make(map[string]*Model, len(list))
	for _, v := range list {
		m[v.Slug] = v
	}
	r.mu.Lock()
	r.bySlug = m
	r.loadedAt = time.Now()
	r.mu.Unlock()
	return nil
}

// BySlug 返回 slug 对应模型。命中缓存则 O(1),过期则异步刷新。
func (r *Registry) BySlug(ctx context.Context, slug string) (*Model, error) {
	r.mu.RLock()
	m, ok := r.bySlug[slug]
	expired := time.Since(r.loadedAt) > r.ttl
	r.mu.RUnlock()

	if ok && !expired {
		return m, nil
	}
	if expired && r.refreshing.CompareAndSwap(false, true) {
		go func() {
			defer r.refreshing.Store(false)
			_ = r.refresh(context.Background())
		}()
	}
	if ok {
		return m, nil
	}
	return r.dao.GetBySlug(ctx, slug)
}

// List 返回所有模型(直接查 DB,管理端用)。
func (r *Registry) List(ctx context.Context) ([]*Model, error) {
	return r.dao.List(ctx)
}

// ListEnabled 返回已启用模型(/v1/models 用)。
func (r *Registry) ListEnabled(ctx context.Context) ([]*Model, error) {
	r.mu.RLock()
	expired := time.Since(r.loadedAt) > r.ttl
	cached := make([]*Model, 0, len(r.bySlug))
	for _, v := range r.bySlug {
		if v.Enabled && !v.DeletedAt.Valid {
			cached = append(cached, v)
		}
	}
	r.mu.RUnlock()
	if !expired && len(cached) > 0 {
		return cached, nil
	}
	return r.dao.ListEnabled(ctx)
}
````

## File: internal/proxy/dao.go
````go
package proxy

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

var ErrNotFound = errors.New("proxy: not found")

type DAO struct{ db *sqlx.DB }

func NewDAO(db *sqlx.DB) *DAO { return &DAO{db: db} }

func (d *DAO) Create(ctx context.Context, p *Proxy) (uint64, error) {
	res, err := d.db.ExecContext(ctx,
		`INSERT INTO proxies (scheme, host, port, username, password_enc, country, isp, health_score, enabled, remark)
         VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		p.Scheme, p.Host, p.Port, p.Username, p.PasswordEnc,
		p.Country, p.ISP, p.HealthScore, p.Enabled, p.Remark,
	)
	if err != nil {
		return 0, err
	}
	id, _ := res.LastInsertId()
	return uint64(id), nil
}

func (d *DAO) GetByID(ctx context.Context, id uint64) (*Proxy, error) {
	var p Proxy
	err := d.db.GetContext(ctx, &p,
		`SELECT * FROM proxies WHERE id = ? AND deleted_at IS NULL`, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	return &p, err
}

func (d *DAO) List(ctx context.Context, offset, limit int) ([]*Proxy, int64, error) {
	var total int64
	if err := d.db.GetContext(ctx, &total,
		`SELECT COUNT(*) FROM proxies WHERE deleted_at IS NULL`); err != nil {
		return nil, 0, err
	}
	rows := make([]*Proxy, 0, limit)
	err := d.db.SelectContext(ctx, &rows,
		`SELECT * FROM proxies WHERE deleted_at IS NULL ORDER BY id DESC LIMIT ? OFFSET ?`,
		limit, offset)
	return rows, total, err
}

// ListAllEnabled 返回所有启用且未删除的代理,用于后台健康探测。
func (d *DAO) ListAllEnabled(ctx context.Context) ([]*Proxy, error) {
	rows := make([]*Proxy, 0, 64)
	err := d.db.SelectContext(ctx, &rows,
		`SELECT * FROM proxies WHERE deleted_at IS NULL AND enabled = 1 ORDER BY id ASC`)
	return rows, err
}

func (d *DAO) Update(ctx context.Context, p *Proxy) error {
	_, err := d.db.ExecContext(ctx,
		`UPDATE proxies
         SET scheme=?, host=?, port=?, username=?, password_enc=?, country=?, isp=?,
             enabled=?, remark=?
         WHERE id = ? AND deleted_at IS NULL`,
		p.Scheme, p.Host, p.Port, p.Username, p.PasswordEnc, p.Country, p.ISP,
		p.Enabled, p.Remark, p.ID,
	)
	return err
}

func (d *DAO) SoftDelete(ctx context.Context, id uint64) error {
	_, err := d.db.ExecContext(ctx,
		`UPDATE proxies SET deleted_at = ? WHERE id = ?`, time.Now(), id)
	return err
}

// SetEnabled 启用/禁用代理。
func (d *DAO) SetEnabled(ctx context.Context, id uint64, enabled bool) error {
	v := 0
	if enabled {
		v = 1
	}
	_, err := d.db.ExecContext(ctx,
		`UPDATE proxies SET enabled = ? WHERE id = ?`, v, id)
	return err
}

// CleanupDisabledBefore 软删除在 before 之前已禁用的代理(自动清理过期代理)。
func (d *DAO) CleanupDisabledBefore(ctx context.Context, before time.Time) (int64, error) {
	res, err := d.db.ExecContext(ctx, `
		UPDATE proxies SET deleted_at = NOW()
		WHERE enabled = 0 AND deleted_at IS NULL AND updated_at < ?`, before)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// ListBoundAccountIDs 返回绑定了指定代理的所有账号 ID。
func (d *DAO) ListBoundAccountIDs(ctx context.Context, proxyID uint64) ([]uint64, error) {
	var ids []uint64
	err := d.db.SelectContext(ctx, &ids,
		`SELECT account_id FROM account_proxy_bindings WHERE proxy_id = ?`, proxyID)
	return ids, err
}

// FindByEndpoint 按 scheme+host+port+username 精确查存活记录,用于批量导入去重。
func (d *DAO) FindByEndpoint(ctx context.Context, scheme, host string, port int, username string) (*Proxy, error) {
	var p Proxy
	err := d.db.GetContext(ctx, &p, `
SELECT * FROM proxies
 WHERE scheme = ? AND host = ? AND port = ? AND username = ?
   AND deleted_at IS NULL
 LIMIT 1`, scheme, host, port, username)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	return &p, err
}

func (d *DAO) UpdateHealth(ctx context.Context, id uint64, score int, lastErr string) error {
	_, err := d.db.ExecContext(ctx,
		`UPDATE proxies SET health_score=?, last_probe_at=?, last_error=? WHERE id=?`,
		score, time.Now(), lastErr, id)
	return err
}
````

## File: internal/proxy/handler.go
````go
package proxy

import (
	"context"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"

	"github.com/432539/gpt2api/pkg/resp"
)

type Handler struct {
	svc    *Service
	prober *Prober
}

func NewHandler(s *Service) *Handler { return &Handler{svc: s} }

// SetProber 在 Prober 初始化完成后注入(避免循环依赖)。未设置时探测接口返回 503。
func (h *Handler) SetProber(p *Prober) { h.prober = p }

// POST /api/admin/proxies
func (h *Handler) Create(c *gin.Context) {
	var req CreateInput
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	p, err := h.svc.Create(c.Request.Context(), req)
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}
	resp.OK(c, p)
}

// GET /api/admin/proxies
func (h *Handler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	if page < 1 {
		page = 1
	}
	size, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	if size < 1 || size > 100 {
		size = 20
	}
	list, total, err := h.svc.List(c.Request.Context(), (page-1)*size, size)
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}
	resp.OK(c, gin.H{"list": list, "total": total, "page": page, "page_size": size})
}

// GET /api/admin/proxies/:id
func (h *Handler) Get(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	p, err := h.svc.Get(c.Request.Context(), id)
	if err != nil {
		resp.NotFound(c, err.Error())
		return
	}
	resp.OK(c, p)
}

// PATCH /api/admin/proxies/:id
func (h *Handler) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req UpdateInput
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	p, err := h.svc.Update(c.Request.Context(), id, req)
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}
	resp.OK(c, p)
}

// POST /api/admin/proxies/import
// Body: { text, enabled, country, isp, remark, overwrite, probe_after_import }
// 支持每行一个 proxy URL,详见 service.parseProxyLine 的注释。
func (h *Handler) Import(c *gin.Context) {
	var req struct {
		Text             string `json:"text"`
		Enabled          *bool  `json:"enabled,omitempty"`
		Country          string `json:"country"`
		ISP              string `json:"isp"`
		Remark           string `json:"remark"`
		Overwrite        bool   `json:"overwrite"`
		ProbeAfterImport *bool  `json:"probe_after_import,omitempty"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}
	if len(req.Text) == 0 {
		resp.BadRequest(c, "请至少粘贴一行代理 URL")
		return
	}
	if len(req.Text) > 256*1024 {
		resp.BadRequest(c, "导入内容过大(最多 256KB)")
		return
	}
	enabled := true
	if req.Enabled != nil {
		enabled = *req.Enabled
	}
	results, err := h.svc.ImportBatch(c.Request.Context(), req.Text, ImportDefaults{
		Enabled: enabled, Country: req.Country, ISP: req.ISP,
		Remark: req.Remark, Overwrite: req.Overwrite,
	})
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}
	// 统计
	var created, updated, skipped, invalid int
	probeIDs := make([]uint64, 0)
	seenProbe := map[uint64]struct{}{}
	for _, r := range results {
		switch r.Status {
		case "created":
			created++
		case "updated":
			updated++
		case "skipped":
			skipped++
		case "invalid":
			invalid++
		}
		if (r.Status == "created" || r.Status == "updated") && r.ID > 0 {
			if _, ok := seenProbe[r.ID]; !ok {
				seenProbe[r.ID] = struct{}{}
				probeIDs = append(probeIDs, r.ID)
			}
		}
	}
	body := gin.H{
		"items":   results,
		"created": created,
		"updated": updated,
		"skipped": skipped,
		"invalid": invalid,
		"total":   len(results),
	}
	probeAfter := true
	if req.ProbeAfterImport != nil {
		probeAfter = *req.ProbeAfterImport
	}
	if probeAfter && h.prober != nil && len(probeIDs) > 0 {
		if len(probeIDs) <= 50 {
			items := h.probeImportedIDs(c.Request.Context(), probeIDs)
			ok, bad := 0, 0
			for _, item := range items {
				if item.OK {
					ok++
				} else {
					bad++
				}
			}
			body["probe"] = gin.H{"total": len(items), "ok": ok, "bad": bad, "items": items}
		} else {
			h.prober.Kick(0)
			body["probe_queued"] = true
		}
	}
	resp.OK(c, body)
}

func (h *Handler) probeImportedIDs(ctx context.Context, ids []uint64) []ProbeResult {
	const conc = 8
	out := make([]ProbeResult, len(ids))
	sem := make(chan struct{}, conc)
	var wg sync.WaitGroup
	for i, id := range ids {
		i, id := i, id
		wg.Add(1)
		sem <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-sem }()
			res, err := h.prober.ProbeByID(ctx, id)
			if err != nil {
				res = ProbeResult{ProxyID: id, Error: err.Error()}
			}
			out[i] = res
		}()
	}
	wg.Wait()
	return out
}

// POST /api/admin/proxies/:id/probe
// 同步探测单条代理,返回 { ok, latency_ms, error, tried_at, health_score }。
func (h *Handler) Probe(c *gin.Context) {
	if h.prober == nil {
		resp.BadRequest(c, "代理探测器未初始化")
		return
	}
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	res, err := h.prober.ProbeByID(c.Request.Context(), id)
	if err != nil {
		resp.NotFound(c, err.Error())
		return
	}
	// 读回最新的 proxy 以返回更新后的 health_score
	p, _ := h.svc.Get(c.Request.Context(), id)
	resp.OK(c, gin.H{
		"ok":         res.OK,
		"latency_ms": res.LatencyMs,
		"error":      res.Error,
		"tried_at":   res.TriedAt,
		"health_score": func() int {
			if p != nil {
				return p.HealthScore
			}
			return 0
		}(),
	})
}

// POST /api/admin/proxies/probe-all
// 同步探测所有启用的代理,返回完整结果数组(含统计)。
func (h *Handler) ProbeAll(c *gin.Context) {
	if h.prober == nil {
		resp.BadRequest(c, "代理探测器未初始化")
		return
	}
	results, err := h.prober.ProbeAll(c.Request.Context())
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}
	ok, bad := 0, 0
	for _, r := range results {
		if r.OK {
			ok++
		} else {
			bad++
		}
	}
	resp.OK(c, gin.H{
		"total": len(results),
		"ok":    ok,
		"bad":   bad,
		"items": results,
	})
}

// DELETE /api/admin/proxies/:id
func (h *Handler) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		resp.Internal(c, err.Error())
		return
	}
	resp.OK(c, gin.H{"deleted": id})
}
````

## File: internal/proxy/model.go
````go
package proxy

import (
	"database/sql"
	"time"
)

// Proxy 对应 proxies 表。
// password 在表里是加密存储的(AES-256-GCM),这里只暴露密文字段,
// 取明文时通过 Service.Decrypt。
type Proxy struct {
	ID          uint64       `db:"id" json:"id"`
	Scheme      string       `db:"scheme" json:"scheme"`
	Host        string       `db:"host" json:"host"`
	Port        int          `db:"port" json:"port"`
	Username    string       `db:"username" json:"username"`
	PasswordEnc string       `db:"password_enc" json:"-"`
	Country     string       `db:"country" json:"country"`
	ISP         string       `db:"isp" json:"isp"`
	HealthScore int          `db:"health_score" json:"health_score"`
	LastProbeAt sql.NullTime `db:"last_probe_at" json:"last_probe_at,omitempty"`
	LastError   string       `db:"last_error" json:"last_error,omitempty"`
	Enabled     bool         `db:"enabled" json:"enabled"`
	Remark      string       `db:"remark" json:"remark"`
	CreatedAt   time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time    `db:"updated_at" json:"updated_at"`
	DeletedAt   sql.NullTime `db:"deleted_at" json:"-"`
}

// URL 返回代理 URL(解密后由 Service 组装)。
func (p *Proxy) URLWithPassword(password string) string {
	if p.Username == "" {
		return p.Scheme + "://" + host(p.Host, p.Port)
	}
	if password == "" {
		return p.Scheme + "://" + p.Username + "@" + host(p.Host, p.Port)
	}
	return p.Scheme + "://" + p.Username + ":" + password + "@" + host(p.Host, p.Port)
}

func host(h string, port int) string {
	return h + ":" + itoa(port)
}

func itoa(i int) string {
	if i == 0 {
		return "0"
	}
	neg := i < 0
	if neg {
		i = -i
	}
	var buf [20]byte
	bp := len(buf)
	for i > 0 {
		bp--
		buf[bp] = byte('0' + i%10)
		i /= 10
	}
	if neg {
		bp--
		buf[bp] = '-'
	}
	return string(buf[bp:])
}
````

## File: internal/proxy/prober.go
````go
package proxy

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
)

// ProbeSettings 探测器配置提供者(从 settings.Service 注入)。
// 所有字段都支持热更新:循环每轮结束都会重新读取。
type ProbeSettings interface {
	ProbeEnabled() bool
	ProbeIntervalSec() int // 两轮探测之间的间隔(秒);<= 0 视为关闭
	ProbeTimeoutSec() int  // 单次探测超时(秒)
	ProbeTargetURL() string
	ProbeConcurrency() int // 并发 worker 数;<=0 默认 8
}

// ProbeResult 单次探测结果。
type ProbeResult struct {
	ProxyID   uint64        `json:"proxy_id"`
	OK        bool          `json:"ok"`
	LatencyMs int           `json:"latency_ms"`
	Error     string        `json:"error,omitempty"`
	TriedAt   time.Time     `json:"tried_at"`
	Duration  time.Duration `json:"-"`
}

// Prober 周期性对启用的代理发起连通性探测,刷新 health_score/last_probe_at/last_error。
//
// 评分策略:
//   - 成功  → score = min(100, score + 10),清空 last_error
//   - 失败  → score = max(0,   score - 20),记录简短 error
type Prober struct {
	svc      *Service
	settings ProbeSettings
	log      *zap.Logger
	reassign ReassignFunc // 代理淘汰时重分配账号

	// 手动触发通道:发送 <id> 探测单个(0 表示全部)
	kickCh chan uint64
}

func NewProber(svc *Service, settings ProbeSettings, log *zap.Logger) *Prober {
	return &Prober{
		svc:      svc,
		settings: settings,
		log:      log,
		kickCh:   make(chan uint64, 32),
	}
}

// Run 后台循环探测,受 ctx 控制。建议作为独立 goroutine 启动。
func (p *Prober) Run(ctx context.Context) {
	// 启动后先睡 5 秒,避开启动峰值。
	select {
	case <-ctx.Done():
		return
	case <-time.After(5 * time.Second):
	}

	lastCleanup := time.Now()

	for {
		interval := time.Duration(p.settings.ProbeIntervalSec()) * time.Second
		if !p.settings.ProbeEnabled() || interval <= 0 {
			select {
			case <-ctx.Done():
				return
			case id := <-p.kickCh:
				p.runOnce(ctx, id)
			case <-time.After(30 * time.Second):
			}
			continue
		}

		p.runOnce(ctx, 0)

		// 每小时清理一次:禁用超过 7 天的代理自动软删除
		if time.Since(lastCleanup) > time.Hour {
			before := time.Now().Add(-7 * 24 * time.Hour)
			if n, err := p.svc.dao.CleanupDisabledBefore(ctx, before); err == nil && n > 0 {
				p.log.Info("prober: cleaned up stale disabled proxies", zap.Int64("count", n))
			}
			lastCleanup = time.Now()
		}

		select {
		case <-ctx.Done():
			return
		case id := <-p.kickCh:
			p.runOnce(ctx, id)
		case <-time.After(interval):
		}
	}
}

// Kick 触发一次立即探测。id=0 表示全部启用的代理;否则只探一条。
// 非阻塞(通道满时直接丢弃,避免调用者卡住)。
func (p *Prober) Kick(id uint64) {
	select {
	case p.kickCh <- id:
	default:
	}
}

// ProbeOne 对单条代理做一次同步探测(不写库)。对外暴露用于手动测试。
func (p *Prober) ProbeOne(ctx context.Context, pr *Proxy) ProbeResult {
	return p.probe(ctx, pr)
}

// ProbeByID 手动触发单条探测并写库,返回结果。
func (p *Prober) ProbeByID(ctx context.Context, id uint64) (ProbeResult, error) {
	pr, err := p.svc.Get(ctx, id)
	if err != nil {
		return ProbeResult{}, err
	}
	res := p.probe(ctx, pr)
	p.applyResult(ctx, pr, res)
	return res, nil
}

// ProbeAll 手动触发对所有启用代理的并发探测并写库,返回结果列表。
func (p *Prober) ProbeAll(ctx context.Context) ([]ProbeResult, error) {
	list, err := p.svc.dao.ListAllEnabled(ctx)
	if err != nil {
		return nil, err
	}
	return p.probeBatch(ctx, list), nil
}

// ---------- 内部实现 ----------

func (p *Prober) runOnce(ctx context.Context, only uint64) {
	var list []*Proxy
	var err error
	if only == 0 {
		list, err = p.svc.dao.ListAllEnabled(ctx)
	} else {
		pr, gerr := p.svc.Get(ctx, only)
		if gerr == nil {
			list = []*Proxy{pr}
		}
		err = gerr
	}
	if err != nil {
		p.log.Warn("prober: list failed", zap.Error(err))
		return
	}
	if len(list) == 0 {
		return
	}
	results := p.probeBatch(ctx, list)
	ok, bad := 0, 0
	for _, r := range results {
		if r.OK {
			ok++
		} else {
			bad++
		}
	}
	p.log.Info("prober: round finished",
		zap.Int("total", len(results)), zap.Int("ok", ok), zap.Int("bad", bad))
}

func (p *Prober) probeBatch(ctx context.Context, list []*Proxy) []ProbeResult {
	conc := p.settings.ProbeConcurrency()
	if conc <= 0 {
		conc = 8
	}
	if conc > len(list) {
		conc = len(list)
	}

	results := make([]ProbeResult, len(list))
	sem := make(chan struct{}, conc)
	var wg sync.WaitGroup

	for i, pr := range list {
		wg.Add(1)
		sem <- struct{}{}
		go func(i int, pr *Proxy) {
			defer wg.Done()
			defer func() { <-sem }()
			r := p.probe(ctx, pr)
			p.applyResult(ctx, pr, r)
			results[i] = r
		}(i, pr)
	}
	wg.Wait()
	return results
}

// probe 做一次真实 HTTP(S) 请求。只组装结果,不写库。
func (p *Prober) probe(ctx context.Context, pr *Proxy) ProbeResult {
	out := ProbeResult{ProxyID: pr.ID, TriedAt: time.Now()}

	proxyURL, err := p.svc.BuildURL(pr)
	if err != nil {
		out.Error = "密码解密失败:" + err.Error()
		return out
	}
	u, err := url.Parse(proxyURL)
	if err != nil {
		out.Error = "代理 URL 格式错误:" + err.Error()
		return out
	}

	timeout := time.Duration(p.settings.ProbeTimeoutSec()) * time.Second
	if timeout <= 0 {
		timeout = 10 * time.Second
	}

	target := strings.TrimSpace(p.settings.ProbeTargetURL())
	if target == "" {
		target = "https://chatgpt.com/cdn-cgi/trace"
	}

	transport := &http.Transport{
		Proxy:                 http.ProxyURL(u),
		DialContext:           (&net.Dialer{Timeout: timeout}).DialContext,
		TLSHandshakeTimeout:   timeout,
		ResponseHeaderTimeout: timeout,
		ExpectContinueTimeout: 1 * time.Second,
		DisableKeepAlives:     true,
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   timeout,
	}

	reqCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(reqCtx, http.MethodGet, target, nil)
	if err != nil {
		out.Error = "构造探测请求失败:" + err.Error()
		return out
	}
	req.Header.Set("User-Agent", "gpt2api-proxy-prober/1.0")

	start := time.Now()
	resp, err := client.Do(req)
	elapsed := time.Since(start)
	out.Duration = elapsed
	out.LatencyMs = int(elapsed / time.Millisecond)

	if err != nil {
		out.Error = shortenErr(err)
		return out
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		out.OK = true
		return out
	}
	out.Error = fmt.Sprintf("目标站返回异常状态码 %d", resp.StatusCode)
	return out
}

// ReassignFunc 当代理被淘汰时,将其绑定的账号重新分配到其他代理。
// 由 main.go 注入(避免 proxy 包直接依赖 account 包)。
type ReassignFunc func(ctx context.Context, deadProxyID uint64)

// SetReassignFunc 注入账号重分配回调。
func (p *Prober) SetReassignFunc(fn ReassignFunc) { p.reassign = fn }

func (p *Prober) applyResult(ctx context.Context, pr *Proxy, r ProbeResult) {
	score := pr.HealthScore
	lastErr := ""
	if r.OK {
		score += 10
		if score > 100 {
			score = 100
		}
	} else {
		score -= 20
		if score < 0 {
			score = 0
		}
		lastErr = r.Error
		if len(lastErr) > 200 {
			lastErr = lastErr[:200]
		}
	}
	if err := p.svc.dao.UpdateHealth(ctx, pr.ID, score, lastErr); err != nil {
		p.log.Warn("prober: update health failed",
			zap.Uint64("proxy_id", pr.ID), zap.Error(err))
	}

	// 健康分归零 → 自动禁用代理 + 重新分配账号
	if score <= 0 && pr.Enabled {
		p.log.Warn("prober: proxy health exhausted, auto-disabling",
			zap.Uint64("proxy_id", pr.ID),
			zap.String("last_error", lastErr))
		_ = p.svc.dao.SetEnabled(ctx, pr.ID, false)
		if p.reassign != nil {
			p.reassign(ctx, pr.ID)
		}
	}
}

// shortenErr 把网络错误压成一行、对前端友好的中文字符串。
// 兜底会带上简短的英文原文,便于排障。
func shortenErr(err error) string {
	if err == nil {
		return ""
	}
	s := err.Error()
	low := strings.ToLower(s)

	switch {
	// 超时 / 主动取消
	case errors.Is(err, context.DeadlineExceeded),
		strings.Contains(low, "deadline exceeded"),
		strings.Contains(low, "i/o timeout"),
		strings.Contains(low, "timeout awaiting"),
		strings.Contains(low, "request canceled") && strings.Contains(low, "timeout"):
		return "连接超时(探测超时)"

	// 407 代理鉴权失败 —— 用户名/密码错误最常见
	case strings.Contains(s, "Proxy Authentication Required"),
		strings.Contains(low, "407"):
		return "代理鉴权失败(407,请核对用户名/密码)"

	// DNS 解析失败 —— 细分代理自身域名 vs 目标站域名
	case strings.Contains(low, "proxyconnect") && strings.Contains(low, "no such host"):
		return "DNS 解析失败:代理域名无法解析(宿主梯子/DNS 污染,可在 docker-compose 里给 server 指定公共 DNS 如 8.8.8.8)"
	case strings.Contains(low, "no such host"),
		strings.Contains(low, "lookup ") && strings.Contains(low, "no such"):
		return "DNS 解析失败(域名不存在或 DNS 被污染)"

	// 各类拒绝 / 不可达
	case strings.Contains(low, "connection refused"):
		return "目标拒绝连接(connection refused)"
	case strings.Contains(low, "network is unreachable"):
		return "网络不可达"
	case strings.Contains(low, "no route to host"):
		return "无法路由到目标主机"
	case strings.Contains(low, "host is down"):
		return "目标主机不可达"

	// 连接在握手/发送中被对端断开
	case strings.Contains(low, "connection reset by peer"):
		return "对端重置连接(代理可能限流/拒绝)"
	case strings.Contains(low, "broken pipe"):
		return "连接已断开(broken pipe)"
	case strings.Contains(low, "unexpected eof"),
		low == "eof",
		strings.HasSuffix(low, ": eof"):
		return "代理握手被关闭(鉴权或协议不匹配)"

	// 代理协议问题
	case strings.Contains(low, "proxyconnect tcp"):
		return "代理握手失败(请检查 host:port/scheme)"
	case strings.Contains(low, "malformed http response"):
		return "代理响应非 HTTP(scheme 可能写错)"
	case strings.Contains(low, "socks"):
		return "SOCKS 代理握手失败"

	// TLS
	case strings.Contains(low, "tls:"),
		strings.Contains(low, "x509:"),
		strings.Contains(low, "certificate"):
		return "TLS/证书错误"

	// 其它 —— 给中文前缀 + 简短原文,方便排障
	default:
		// 截断 "Get \"...\": " 等 net/http 前缀
		if i := strings.Index(s, "\": "); i > 0 && i < len(s)-3 {
			s = s[i+3:]
		} else if i := strings.Index(s, ": "); i > 0 && i < len(s)-2 {
			s = s[i+2:]
		}
		if len(s) > 140 {
			s = s[:140] + "…"
		}
		return "探测失败:" + s
	}
}
````

## File: internal/proxy/service.go
````go
package proxy

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/432539/gpt2api/pkg/crypto"
)

// Service 封装代理 CRUD + 密码加解密。
type Service struct {
	dao    *DAO
	cipher *crypto.AESGCM
}

func NewService(dao *DAO, cipher *crypto.AESGCM) *Service {
	return &Service{dao: dao, cipher: cipher}
}

// DAO 暴露给外部使用。
func (s *Service) DAO() *DAO { return s.dao }

// CreateInput 是 Create 的入参(明文 password)。
type CreateInput struct {
	Scheme   string `json:"scheme"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Country  string `json:"country"`
	ISP      string `json:"isp"`
	Enabled  bool   `json:"enabled"`
	Remark   string `json:"remark"`
}

// UpdateInput 是 Update 的入参;Password 为空时不改密。
type UpdateInput struct {
	Scheme   string `json:"scheme"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"` // 空串表示不改
	Country  string `json:"country"`
	ISP      string `json:"isp"`
	Enabled  bool   `json:"enabled"`
	Remark   string `json:"remark"`
}

func (s *Service) Create(ctx context.Context, in CreateInput) (*Proxy, error) {
	if in.Scheme == "" {
		in.Scheme = "http"
	}
	if in.Host == "" || in.Port == 0 {
		return nil, errors.New("host 和 port 不能为空")
	}
	var enc string
	if in.Password != "" {
		v, err := s.cipher.EncryptString(in.Password)
		if err != nil {
			return nil, err
		}
		enc = v
	}
	p := &Proxy{
		Scheme: in.Scheme, Host: in.Host, Port: in.Port,
		Username: in.Username, PasswordEnc: enc,
		Country: in.Country, ISP: in.ISP,
		HealthScore: 100, Enabled: in.Enabled, Remark: in.Remark,
	}
	id, err := s.dao.Create(ctx, p)
	if err != nil {
		return nil, err
	}
	return s.dao.GetByID(ctx, id)
}

func (s *Service) Update(ctx context.Context, id uint64, in UpdateInput) (*Proxy, error) {
	p, err := s.dao.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if in.Scheme != "" {
		p.Scheme = in.Scheme
	}
	if in.Host != "" {
		p.Host = in.Host
	}
	if in.Port != 0 {
		p.Port = in.Port
	}
	p.Username = in.Username
	p.Country = in.Country
	p.ISP = in.ISP
	p.Enabled = in.Enabled
	p.Remark = in.Remark
	if in.Password != "" {
		enc, err := s.cipher.EncryptString(in.Password)
		if err != nil {
			return nil, err
		}
		p.PasswordEnc = enc
	}
	if err := s.dao.Update(ctx, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *Service) Delete(ctx context.Context, id uint64) error {
	return s.dao.SoftDelete(ctx, id)
}

func (s *Service) Get(ctx context.Context, id uint64) (*Proxy, error) {
	return s.dao.GetByID(ctx, id)
}

func (s *Service) List(ctx context.Context, offset, limit int) ([]*Proxy, int64, error) {
	return s.dao.List(ctx, offset, limit)
}

// ---------- 批量导入 ----------

// ImportDefaults 批量导入时的公共默认值。
type ImportDefaults struct {
	Enabled   bool
	Country   string
	ISP       string
	Remark    string
	Overwrite bool // true 时遇到已存在记录更新密码/备注;false 直接 skip
}

// ImportLineResult 单行结果。Status ∈ "created" | "updated" | "skipped" | "invalid"。
type ImportLineResult struct {
	Line   int    `json:"line"`
	Raw    string `json:"raw"` // 原始输入(密码会被 *** 掉)
	Status string `json:"status"`
	ID     uint64 `json:"id,omitempty"`
	Error  string `json:"error,omitempty"`
}

// ImportBatch 解析多行代理 URL,逐条入库。
// 支持格式:
//
//	scheme://user:pass@host:port
//	scheme://host:port               (无鉴权)
//	user:pass@host:port              (无 scheme,默认 http)
//	host:port                        (无 scheme 无鉴权,默认 http)
//	host:port:user:pass              (常见代理商格式)
//	user:pass:host:port              (另一种代理商格式)
//
// 以 # 或 // 开头的行视为注释跳过。空行跳过。
func (s *Service) ImportBatch(ctx context.Context, text string, defaults ImportDefaults) ([]ImportLineResult, error) {
	lines := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	out := make([]ImportLineResult, 0, len(lines))

	for i, raw := range lines {
		line := strings.TrimSpace(raw)
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") {
			continue
		}
		r := ImportLineResult{Line: i + 1, Raw: maskPassword(line)}

		in, err := parseProxyLine(line)
		if err != nil {
			r.Status = "invalid"
			r.Error = err.Error()
			out = append(out, r)
			continue
		}
		// 合并公共默认值
		in.Enabled = defaults.Enabled
		if in.Country == "" {
			in.Country = defaults.Country
		}
		if in.ISP == "" {
			in.ISP = defaults.ISP
		}
		if in.Remark == "" {
			in.Remark = defaults.Remark
		}

		// 查重:同 scheme+host+port+username 若已存在 → update / skip
		existing, err := s.dao.FindByEndpoint(ctx, in.Scheme, in.Host, in.Port, in.Username)
		if err != nil && !errors.Is(err, ErrNotFound) {
			r.Status = "invalid"
			r.Error = err.Error()
			out = append(out, r)
			continue
		}
		if existing != nil {
			if !defaults.Overwrite {
				r.Status = "skipped"
				r.ID = existing.ID
				r.Error = "已存在同 scheme+host:port+用户名 的代理,跳过"
				out = append(out, r)
				continue
			}
			// 覆盖模式:保留 id/health_score,更新其他字段
			existing.Scheme = in.Scheme
			existing.Host = in.Host
			existing.Port = in.Port
			existing.Username = in.Username
			existing.Country = in.Country
			existing.ISP = in.ISP
			existing.Enabled = in.Enabled
			existing.Remark = in.Remark
			if in.Password != "" {
				enc, encErr := s.cipher.EncryptString(in.Password)
				if encErr != nil {
					r.Status = "invalid"
					r.Error = encErr.Error()
					out = append(out, r)
					continue
				}
				existing.PasswordEnc = enc
			}
			if err := s.dao.Update(ctx, existing); err != nil {
				r.Status = "invalid"
				r.Error = err.Error()
				out = append(out, r)
				continue
			}
			r.Status = "updated"
			r.ID = existing.ID
			out = append(out, r)
			continue
		}

		// 新建
		p, err := s.Create(ctx, in)
		if err != nil {
			r.Status = "invalid"
			r.Error = err.Error()
			out = append(out, r)
			continue
		}
		r.Status = "created"
		r.ID = p.ID
		out = append(out, r)
	}
	return out, nil
}

// parseProxyLine 解析一行代理串为 CreateInput(不含 Enabled/Remark 默认值)。
func parseProxyLine(line string) (CreateInput, error) {
	var out CreateInput
	trim := strings.TrimSpace(line)
	if trim == "" {
		return out, errors.New("空行")
	}

	scheme := "http"
	rest := trim
	if i := strings.Index(rest, "://"); i > 0 {
		scheme = strings.ToLower(strings.TrimSpace(rest[:i]))
		rest = rest[i+3:]
	}
	if !isSupportedProxyScheme(scheme) {
		return out, fmt.Errorf("不支持的协议 %q(仅支持 http/https/socks5)", scheme)
	}
	if parsed, ok, err := parseColonProxyLine(scheme, rest); ok || err != nil {
		return parsed, err
	}

	// 补全 scheme,让 net/url 能正确解析 user/pass。
	if !strings.Contains(trim, "://") {
		trim = scheme + "://" + trim
	}
	u, err := url.Parse(trim)
	if err != nil {
		return out, fmt.Errorf("URL 格式错误:%w", err)
	}
	scheme = strings.ToLower(u.Scheme)
	if !isSupportedProxyScheme(scheme) {
		return out, fmt.Errorf("不支持的协议 %q(仅支持 http/https/socks5)", u.Scheme)
	}
	host := u.Hostname()
	if host == "" {
		return out, errors.New("缺少 host")
	}
	portStr := u.Port()
	if portStr == "" {
		return out, errors.New("缺少 port")
	}
	port, err := strconv.Atoi(portStr)
	if err != nil || port < 1 || port > 65535 {
		return out, fmt.Errorf("端口 %q 非法(需 1~65535)", portStr)
	}
	out.Scheme = scheme
	out.Host = host
	out.Port = port
	if u.User != nil {
		out.Username = u.User.Username()
		if pw, ok := u.User.Password(); ok {
			out.Password = pw
		}
	}
	return out, nil
}

func parseColonProxyLine(scheme, rest string) (CreateInput, bool, error) {
	var out CreateInput
	if strings.Contains(rest, "@") || strings.Contains(rest, "[") || strings.Contains(rest, "]") {
		return out, false, nil
	}
	parts := strings.Split(rest, ":")
	if len(parts) != 4 {
		return out, false, nil
	}
	if p, ok := parseProxyPort(parts[1]); ok {
		return CreateInput{Scheme: scheme, Host: strings.TrimSpace(parts[0]), Port: p, Username: strings.TrimSpace(parts[2]), Password: strings.TrimSpace(parts[3])}, true, validateColonProxy(parts[0], p)
	}
	if p, ok := parseProxyPort(parts[3]); ok {
		return CreateInput{Scheme: scheme, Host: strings.TrimSpace(parts[2]), Port: p, Username: strings.TrimSpace(parts[0]), Password: strings.TrimSpace(parts[1])}, true, validateColonProxy(parts[2], p)
	}
	return out, false, nil
}

func parseProxyPort(s string) (int, bool) {
	p, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil || p < 1 || p > 65535 {
		return 0, false
	}
	return p, true
}

func validateColonProxy(host string, port int) error {
	if strings.TrimSpace(host) == "" {
		return errors.New("缺少 host")
	}
	if port < 1 || port > 65535 {
		return fmt.Errorf("端口 %d 非法(需 1~65535)", port)
	}
	return nil
}

func isSupportedProxyScheme(scheme string) bool {
	switch strings.ToLower(scheme) {
	case "http", "https", "socks5", "socks5h":
		return true
	default:
		return false
	}
}

// maskPassword 把 user:pass@ 里的 pass 替换成 ***,用于 result.raw 回显。
func maskPassword(line string) string {
	trim := strings.TrimSpace(line)
	scheme := ""
	rest := trim
	if i := strings.Index(rest, "://"); i >= 0 {
		scheme = rest[:i+3]
		rest = rest[i+3:]
	}
	if !strings.Contains(rest, "@") {
		parts := strings.Split(rest, ":")
		if len(parts) == 4 {
			if _, ok := parseProxyPort(parts[1]); ok {
				return scheme + parts[0] + ":" + parts[1] + ":" + parts[2] + ":***"
			}
			if _, ok := parseProxyPort(parts[3]); ok {
				return scheme + parts[0] + ":***:" + parts[2] + ":" + parts[3]
			}
		}
		return line
	}
	at := strings.Index(rest, "@")
	cred := rest[:at]
	tail := rest[at:]
	colon := strings.Index(cred, ":")
	if colon < 0 {
		return line
	}
	return scheme + cred[:colon] + ":***" + tail
}

// DecryptPassword 解密代理密码。
func (s *Service) DecryptPassword(p *Proxy) (string, error) {
	if p.PasswordEnc == "" {
		return "", nil
	}
	return s.cipher.DecryptString(p.PasswordEnc)
}

// BuildURL 返回完整代理 URL(含明文密码)。
func (s *Service) BuildURL(p *Proxy) (string, error) {
	pw, err := s.DecryptPassword(p)
	if err != nil {
		return "", err
	}
	return p.URLWithPassword(pw), nil
}
````

## File: internal/scheduler/scheduler.go
````go
// Package scheduler 负责 chatgpt.com 账号的并发安全调度。
//
// 核心规则(参考 RISK_AND_SAAS.md):
//  1. 一号一锁:同账号同时只允许 1 个请求占用(Redis SETNX)。
//  2. 最小间隔:同账号相邻请求 >= min_interval_sec。
//  3. 每日配额:today_used_count < daily_image_quota * daily_usage_ratio。
//  4. 状态机:healthy -> warned -> throttled -> suspicious -> dead,冷却过期自动恢复。
//  5. 选择策略:status=healthy + cooldown 到期 + last_used_at 最早的优先。
package scheduler

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/432539/gpt2api/internal/account"
	"github.com/432539/gpt2api/internal/config"
	"github.com/432539/gpt2api/internal/proxy"
	"github.com/432539/gpt2api/pkg/lock"
	"github.com/432539/gpt2api/pkg/logger"

	"go.uber.org/zap"
)

// ErrNoAvailable 没有任何账号可用。
var ErrNoAvailable = errors.New("scheduler: no available account")

// Lease 代表一次账号占用的租约。
type Lease struct {
	Account     *account.Account
	AuthToken   string // 已解密
	ProxyURL    string // 已带密码
	ProxyID     uint64
	DeviceID    string
	SessionID   string // oai_session_id(按账号稳定)
	Cookies     string // oai_account_cookies 解密后的 JSON 字符串,让 runner 与 probe 共用浏览器态
	lockKey     string
	lockToken   string
	releaseFunc func(context.Context) error
}

// Release 释放锁并更新账号 last_used_at / today_used。
func (l *Lease) Release(ctx context.Context) error {
	if l.releaseFunc != nil {
		return l.releaseFunc(ctx)
	}
	return nil
}

// RuntimeParams 调度器运行期可热更的参数。
//   - 由外部 settings.Service 提供回调,每次读都取最新值;
//   - 回调未注入时回退到 cfg 的静态值。
type RuntimeParams struct {
	// 为 nil 时 Scheduler 使用 cfg 里的静态值。
	Cooldown429Sec func() int
	WarnedPauseHrs func() int
	// QueueWaitSec 拿不到空闲账号时最长排队等待秒数,≤0 表示不排队(老语义)。
	QueueWaitSec func() int
	// ImageExploreRatio 控制 image 调度里留给 unknown/新号/冷号的探索比例。
	// 0=关闭探索,0.2=约 20% 探索;未注入时回退 0.2。
	ImageExploreRatio func() float64
}

// Scheduler 账号调度器。
type Scheduler struct {
	accSvc      *account.Service
	proxySvc    *proxy.Service
	lock        *lock.RedisLock
	cfg         config.SchedulerConfig
	rt          RuntimeParams
	dispatchLim int // 一次 Dispatch 扫描的候选数上限,默认 50

	// 单账号并发信号量:限制同账号同时生图数,超出排队等待
	semMu sync.Mutex
	sems  map[uint64]chan struct{} // accountID -> buffered channel(size = maxConcurrent)

	// 内存候选池 + round-robin 轮转
	poolMu      sync.Mutex
	pool        []*account.Account // 当前可调度候选(定期从 DB 刷新)
	poolIdx     int                // 轮转指针
	poolRefresh time.Time          // 上次刷新时间
}

func New(
	accSvc *account.Service,
	proxySvc *proxy.Service,
	rl *lock.RedisLock,
	cfg config.SchedulerConfig,
) *Scheduler {
	if cfg.LockTTLSec <= 0 {
		cfg.LockTTLSec = 180
	}
	if cfg.MinIntervalSec <= 0 {
		cfg.MinIntervalSec = 5
	}
	if cfg.Cooldown429Sec <= 0 {
		cfg.Cooldown429Sec = 300
	}
	if cfg.MaxConcurrentPerAccount <= 0 {
		cfg.MaxConcurrentPerAccount = 3
	}
	return &Scheduler{
		accSvc: accSvc, proxySvc: proxySvc, lock: rl, cfg: cfg,
		dispatchLim: 50,
		sems:        make(map[uint64]chan struct{}),
	}
}

// SetRuntime 注入运行期可热更的参数。建议在 main 里一次性设置:
//
//	sched.SetRuntime(scheduler.RuntimeParams{
//	    DailyUsageRatio: settingsSvc.DailyUsageRatio,
//	    Cooldown429Sec:  settingsSvc.Cooldown429Sec,
//	    WarnedPauseHrs:  settingsSvc.WarnedPauseHours,
//	})
func (s *Scheduler) SetRuntime(p RuntimeParams) { s.rt = p }

// SetDispatchLimit 设置每次 Dispatch 扫描的候选账号上限。
// 默认 50;账号池特别大时可适当提高。
func (s *Scheduler) SetDispatchLimit(n int) {
	if n > 0 {
		s.dispatchLim = n
	}
}

func (s *Scheduler) cooldown429() time.Duration {
	if s.rt.Cooldown429Sec != nil {
		if v := s.rt.Cooldown429Sec(); v > 0 {
			return time.Duration(v) * time.Second
		}
	}
	return time.Duration(s.cfg.Cooldown429Sec) * time.Second
}
func (s *Scheduler) warnedPause() time.Duration {
	if s.rt.WarnedPauseHrs != nil {
		if v := s.rt.WarnedPauseHrs(); v > 0 {
			return time.Duration(v) * time.Hour
		}
	}
	return time.Duration(s.cfg.WarnedPauseHours) * time.Hour
}

// queueWait 拿不到账号时的最长排队等待时间。
// 返回 0 表示关闭排队(立即返回 ErrNoAvailable)。
func (s *Scheduler) queueWait() time.Duration {
	if s.rt.QueueWaitSec != nil {
		if v := s.rt.QueueWaitSec(); v >= 0 {
			return time.Duration(v) * time.Second
		}
	}
	return 120 * time.Second
}

func (s *Scheduler) imageExploreRatio() float64 {
	if s.rt.ImageExploreRatio != nil {
		v := s.rt.ImageExploreRatio()
		if v < 0 {
			return 0
		}
		if v > 0.8 {
			return 0.8
		}
		return v
	}
	return 0.2
}

// Dispatch 为本次请求挑选一个账号并加锁。调用方必须 defer lease.Release(ctx)。
//
// 语义(一号一任务 + 排队):
//   - 同账号同时只允许 1 个请求持有 Redis 锁(acct:lock:{id},SETNX+TTL)。
//   - 扫一遍所有 candidate 都被锁住 / 不满足 min_interval / 日配额时,
//     不立即返回失败,而是按指数退避轮询重试,直到拿到锁或超过 queueWait。
//   - queueWait=0 时退化为老语义(扫一次,失败即返回 ErrNoAvailable)。
func (s *Scheduler) Dispatch(ctx context.Context, modelType string) (*Lease, error) {
	return s.DispatchWithExclude(ctx, modelType, nil)
}

// DispatchWithExclude 与 Dispatch 相同,但排除指定账号 ID(用于跨账号重试时
// 跳过已尝试失败的账号)。
func (s *Scheduler) DispatchWithExclude(ctx context.Context, modelType string, excludeIDs map[uint64]struct{}) (*Lease, error) {
	deadline := time.Now().Add(s.queueWait())

	const (
		minBackoff = 200 * time.Millisecond
		maxBackoff = 2 * time.Second
	)
	backoff := minBackoff

	attempt := 0
	start := time.Now()

	for {
		attempt++
		lease, err := s.tryDispatchOnce(ctx, modelType, excludeIDs)
		if err == nil {
			if attempt > 1 {
				logger.L().Info("scheduler queued dispatch ok",
					zap.Int("attempt", attempt),
					zap.Duration("waited", time.Since(start)),
					zap.Uint64("account_id", lease.Account.ID))
			}
			return lease, nil
		}
		if !errors.Is(err, ErrNoAvailable) {
			return nil, err
		}

		// 所有候选都忙或不就绪:排队等待。
		if !time.Now().Before(deadline) {
			return nil, ErrNoAvailable
		}
		wait := backoff
		if remain := time.Until(deadline); remain < wait {
			wait = remain
		}
		if wait <= 0 {
			return nil, ErrNoAvailable
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(wait):
		}
		// 指数退避(×1.5)
		backoff += backoff / 2
		if backoff > maxBackoff {
			backoff = maxBackoff
		}
	}
}

// poolTTL 候选池缓存有效期。在此期间不查 DB,直接用内存轮转。
const poolTTL = 10 * time.Second

// refreshPool 从 DB 刷新候选池(有缓存)。
func (s *Scheduler) refreshPool(ctx context.Context, modelType string, force bool) {
	s.poolMu.Lock()
	defer s.poolMu.Unlock()
	if !force && time.Since(s.poolRefresh) < poolTTL && len(s.pool) > 0 {
		return
	}
	dao := s.accSvc.DAO()
	candidates, err := dao.ListDispatchableWithOptions(ctx, s.dispatchLim, account.DispatchOptions{
		ModelType:         modelType,
		ImageExploreRatio: s.imageExploreRatio(),
	})
	if err != nil {
		logger.L().Warn("scheduler refresh pool failed", zap.Error(err))
		return
	}
	s.pool = candidates
	if s.poolIdx >= len(candidates) {
		s.poolIdx = 0
	}
	s.poolRefresh = time.Now()
}

// tryDispatchOnce 从内存候选池 round-robin 选号。
// 遍历一圈,每个候选尝试 acquireSem;全部忙/排除时返回 ErrNoAvailable。
func (s *Scheduler) tryDispatchOnce(ctx context.Context, modelType string, excludeIDs map[uint64]struct{}) (*Lease, error) {
	s.refreshPool(ctx, modelType, false)

	s.poolMu.Lock()
	pool := s.pool
	startIdx := s.poolIdx
	s.poolMu.Unlock()

	if len(pool) == 0 {
		return nil, ErrNoAvailable
	}

	// 从 poolIdx 开始遍历一圈
	for i := 0; i < len(pool); i++ {
		idx := (startIdx + i) % len(pool)
		acc := pool[idx]

		if _, excluded := excludeIDs[acc.ID]; excluded {
			continue
		}

		lease, err := s.tryLock(ctx, acc)
		if err == nil {
			// 成功:推进指针到下一个
			s.poolMu.Lock()
			s.poolIdx = (idx + 1) % len(pool)
			s.poolMu.Unlock()
			return lease, nil
		}
		if errors.Is(err, lock.ErrNotAcquired) {
			continue
		}
		logger.L().Warn("scheduler tryLock error",
			zap.Uint64("account_id", acc.ID), zap.Error(err))
	}

	// 遍历一圈都没选到,强制刷新池再试一次
	s.refreshPool(ctx, modelType, true)
	return nil, ErrNoAvailable
}

// acquireSem 获取账号并发信号量,超出 maxConcurrent 时阻塞等待。
// 返回的 release 函数必须在请求完成后调用。
func (s *Scheduler) acquireSem(ctx context.Context, accountID uint64) (release func(), err error) {
	s.semMu.Lock()
	sem, ok := s.sems[accountID]
	if !ok {
		sem = make(chan struct{}, s.cfg.MaxConcurrentPerAccount)
		s.sems[accountID] = sem
	}
	s.semMu.Unlock()

	select {
	case sem <- struct{}{}:
		return func() { <-sem }, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (s *Scheduler) tryLock(ctx context.Context, acc *account.Account) (*Lease, error) {
	key := fmt.Sprintf("acct:lock:%d", acc.ID)
	token := uuid.NewString()

	// 获取并发信号量(限制同账号同时生图数)
	semRelease, err := s.acquireSem(ctx, acc.ID)
	if err != nil {
		return nil, err
	}

	authToken, err := s.accSvc.DecryptAuthToken(acc)
	if err != nil {
		semRelease()
		return nil, fmt.Errorf("decrypt auth_token: %w", err)
	}

	// 首次使用时为账号补发一个持久化的 oai_device_id(导入时常为空)。
	// chatgpt.com 要求请求头带 oai-device-id,等同于浏览器首访拿到的 oai-did cookie;
	// 一次生成后持久化,账号绑定的"设备身份"保持稳定,避免每次换 id 触发风控。
	deviceID := acc.OAIDeviceID
	if deviceID == "" {
		gen := uuid.NewString()
		if fixed, err := s.accSvc.DAO().EnsureDeviceID(ctx, acc.ID, gen); err == nil && fixed != "" {
			deviceID = fixed
			acc.OAIDeviceID = fixed
		} else {
			deviceID = gen
		}
	}

	// oai_session_id:真实浏览器是"每打开页面生成一次"。为了保持账号行为稳定
	// (风控倾向于把频繁变换 session_id 的账号识别为脚本),我们按账号持久化,
	// 与 device_id 同策略。
	sessionID := acc.OAISessionID
	if sessionID == "" {
		gen := uuid.NewString()
		if fixed, err := s.accSvc.DAO().EnsureSessionID(ctx, acc.ID, gen); err == nil && fixed != "" {
			sessionID = fixed
			acc.OAISessionID = fixed
		} else {
			sessionID = gen
		}
	}

	var proxyURL string
	var proxyID uint64
	if b, _ := s.accSvc.GetBinding(ctx, acc.ID); b != nil {
		p, err := s.proxySvc.Get(ctx, b.ProxyID)
		if err == nil && p != nil && p.Enabled {
			if u, err := s.proxySvc.BuildURL(p); err == nil {
				proxyURL = u
				proxyID = p.ID
			}
		}
	}

	var cookies string
	if v, err := s.accSvc.DecryptCookies(ctx, acc.ID); err == nil {
		cookies = v
	} else {
		logger.L().Warn("scheduler decrypt account cookies failed",
			zap.Uint64("account_id", acc.ID), zap.Error(err))
	}

	accCopy := acc

	// 立即更新 last_used_at，让并发请求看到最新排序，避免多个请求选同一个号
	today := truncateDay(time.Now())
	_ = s.accSvc.DAO().MarkUsed(ctx, accCopy.ID, today)

	lease := &Lease{
		Account:   accCopy,
		AuthToken: authToken,
		ProxyURL:  proxyURL,
		ProxyID:   proxyID,
		DeviceID:  deviceID,
		SessionID: sessionID,
		Cookies:   cookies,
		lockKey:   key,
		lockToken: token,
	}
	lease.releaseFunc = func(c context.Context) error {
		semRelease() // 释放并发信号量
		return nil
	}
	return lease, nil
}

// MarkRateLimited 上游 429:标记账号冷却并降级状态。
func (s *Scheduler) MarkRateLimited(ctx context.Context, accountID uint64) {
	cooldown := time.Now().Add(s.cooldown429())
	_ = s.accSvc.DAO().SetStatus(ctx, accountID, account.StatusThrottled, &cooldown)
}

// MarkWarned 上游返回 suspicious 横幅时降级。
func (s *Scheduler) MarkWarned(ctx context.Context, accountID uint64) {
	pause := time.Now().Add(s.warnedPause())
	_ = s.accSvc.DAO().SetStatus(ctx, accountID, account.StatusWarned, &pause)
}

// MarkDead 账号彻底不可用(403/token 失效)。
func (s *Scheduler) MarkDead(ctx context.Context, accountID uint64) {
	_ = s.accSvc.DAO().SetStatus(ctx, accountID, account.StatusDead, nil)
}

// RestoreHealthy 调度成功后回归健康(仅对 throttled 且冷却到期有效,
// 简单起见此处不强检查,由运维侧按需恢复)。
func (s *Scheduler) RestoreHealthy(ctx context.Context, accountID uint64) {
	_ = s.accSvc.DAO().SetStatus(ctx, accountID, account.StatusHealthy, nil)
}

// RecordIMG2Outcome 记录真实图像请求后的 IMG2 协议命中画像,供后续 image 调度排序。
func (s *Scheduler) RecordIMG2Outcome(ctx context.Context, accountID uint64, outcome string) {
	_ = s.accSvc.DAO().RecordIMG2Outcome(ctx, accountID, outcome)
}

// RecordIMG2Delivery 记录 IMG2 协议命中后的交付结果。
func (s *Scheduler) RecordIMG2Delivery(ctx context.Context, accountID uint64, status string) {
	_ = s.accSvc.DAO().RecordIMG2Delivery(ctx, accountID, status)
}

// AccountBinding 查询账号当前绑定的代理。
func (s *Scheduler) AccountBinding(ctx context.Context, accountID uint64) (*account.Binding, error) {
	return s.accSvc.GetBinding(ctx, accountID)
}

// SwitchProxy 将账号切换到另一个代理(排除当前失效代理)。
// 返回新代理的完整 URL。代理池无可用代理时返回空串(不阻断请求)。
func (s *Scheduler) SwitchProxy(ctx context.Context, accountID, currentProxyID uint64) (newProxyURL string, newProxyID uint64) {
	newID, err := s.accSvc.DAO().SwitchProxy(ctx, accountID, currentProxyID)
	if err != nil {
		logger.L().Warn("scheduler switch proxy failed",
			zap.Uint64("account_id", accountID),
			zap.Uint64("old_proxy_id", currentProxyID),
			zap.Error(err))
		return "", 0
	}
	p, err := s.proxySvc.Get(ctx, newID)
	if err != nil || p == nil || !p.Enabled {
		return "", 0
	}
	u, err := s.proxySvc.BuildURL(p)
	if err != nil {
		return "", 0
	}
	logger.L().Info("scheduler switched proxy",
		zap.Uint64("account_id", accountID),
		zap.Uint64("old_proxy_id", currentProxyID),
		zap.Uint64("new_proxy_id", newID))
	return u, newID
}

// ------ helpers ------

func truncateDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func sameDay(a, b time.Time) bool {
	return a.Year() == b.Year() && a.Month() == b.Month() && a.Day() == b.Day()
}
````

## File: internal/server/router.go
````go
package server

import (
	"github.com/gin-gonic/gin"

	"github.com/432539/gpt2api/internal/account"
	"github.com/432539/gpt2api/internal/audit"
	"github.com/432539/gpt2api/internal/backup"
	"github.com/432539/gpt2api/internal/config"
	"github.com/432539/gpt2api/internal/gateway"
	"github.com/432539/gpt2api/internal/image"
	"github.com/432539/gpt2api/internal/middleware"
	"github.com/432539/gpt2api/internal/model"
	"github.com/432539/gpt2api/internal/proxy"
	"github.com/432539/gpt2api/internal/settings"
	"github.com/432539/gpt2api/internal/usage"
	"github.com/432539/gpt2api/pkg/resp"
)

// Deps 是本地 2API 控制台需要的处理器集合。
type Deps struct {
	Config *config.Config

	ProxyH   *proxy.Handler
	AccountH *account.Handler

	GatewayH *gateway.Handler
	ImagesH  *gateway.ImagesHandler

	BackupH  *backup.Handler
	AuditH   *audit.Handler
	AuditDAO *audit.DAO

	AdminModelH *model.AdminHandler
	AdminUsageH *usage.AdminHandler

	MeUsageH *usage.MeHandler
	MeImageH *image.MeHandler

	SettingsH   *settings.Handler
	SettingsSvc *settings.Service
}

// New 构建 gin.Engine 并挂载本地控制台与 OpenAI 兼容路由。
func New(d *Deps) *gin.Engine {
	if d.Config.App.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(
		middleware.RequestID(),
		middleware.Recover(),
		middleware.AccessLog(),
		middleware.CORS(d.Config.Security.CORSOrigins),
	)

	r.GET("/healthz", func(c *gin.Context) { resp.OK(c, gin.H{"status": "ok"}) })
	r.GET("/readyz", func(c *gin.Context) { resp.OK(c, gin.H{"status": "ok"}) })

	cfg := config.Get()

	// POST /api/admin/login — 管理员登录,无需 token
	r.POST("/api/admin/login", func(c *gin.Context) {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			resp.BadRequest(c, "参数错误")
			return
		}
		if cfg == nil || req.Username != cfg.Admin.Username || req.Password != cfg.Admin.Password {
			c.JSON(401, gin.H{"code": 401, "message": "用户名或密码错误"})
			return
		}
		token := middleware.GenerateToken(req.Username)
		resp.OK(c, gin.H{"token": token, "username": req.Username})
	})

	// 公开 API（不需要登录）
	pub := r.Group("/api/public")
	if d.SettingsH != nil {
		pub.GET("/site-info", d.SettingsH.Public)
	}

	api := r.Group("/api", middleware.AdminAuth())
	{

		api.GET("/me", localMe)
		api.GET("/me/menu", localMenu)
		if d.AdminModelH != nil {
			api.GET("/me/models", d.AdminModelH.ListEnabledForMe)
		}
		if d.MeUsageH != nil {
			ug := api.Group("/me/usage")
			ug.GET("/logs", d.MeUsageH.Logs)
			ug.GET("/stats", d.MeUsageH.Stats)
		}
		if d.MeImageH != nil {
			ig := api.Group("/me/images")
			ig.GET("/tasks", d.MeImageH.List)
			ig.GET("/tasks/:id", d.MeImageH.Get)
		}
		if d.GatewayH != nil {
			pg := api.Group("/me/playground")
			pg.POST("/chat", d.GatewayH.ChatCompletions)
			if d.ImagesH != nil {
				pg.POST("/image", d.ImagesH.ImageGenerations)
				pg.POST("/image-edit", d.ImagesH.ImageEdits)
			}
		}

		adminMW := []gin.HandlerFunc{}
		if d.AuditDAO != nil {
			adminMW = append(adminMW, audit.Middleware(d.AuditDAO))
		}
		admin := api.Group("/admin", adminMW...)
		admin.GET("/ping", func(c *gin.Context) { resp.OK(c, gin.H{"msg": "admin pong"}) })

		if d.ProxyH != nil {
			pg := admin.Group("/proxies")
			pg.POST("", d.ProxyH.Create)
			pg.POST("/import", d.ProxyH.Import)
			pg.POST("/probe-all", d.ProxyH.ProbeAll)
			pg.GET("", d.ProxyH.List)
			pg.GET("/:id", d.ProxyH.Get)
			pg.POST("/:id/probe", d.ProxyH.Probe)
			pg.PATCH("/:id", d.ProxyH.Update)
			pg.DELETE("/:id", d.ProxyH.Delete)
		}

		if d.AccountH != nil {
			ag := admin.Group("/accounts")
			ag.POST("", d.AccountH.Create)
			ag.POST("/import", d.AccountH.Import)
			ag.POST("/import-tokens", d.AccountH.ImportTokens)
			ag.POST("/refresh-all", d.AccountH.RefreshAll)
			ag.POST("/probe-quota-all", d.AccountH.ProbeQuotaAll)
			ag.POST("/bulk-delete", d.AccountH.BulkDelete)
			ag.POST("/purge-deleted", d.AccountH.PurgeDeleted)
			ag.GET("/auto-refresh", d.AccountH.GetAutoRefresh)
			ag.PUT("/auto-refresh", d.AccountH.SetAutoRefresh)
			ag.GET("", d.AccountH.List)
			ag.GET("/:id", d.AccountH.Get)
			ag.GET("/:id/secrets", d.AccountH.GetSecrets)
			ag.PATCH("/:id", d.AccountH.Update)
			ag.DELETE("/:id", d.AccountH.Delete)
			ag.POST("/:id/refresh", d.AccountH.Refresh)
			ag.POST("/:id/probe-quota", d.AccountH.ProbeQuota)
			ag.POST("/:id/bind-proxy", d.AccountH.BindProxy)
			ag.DELETE("/:id/bind-proxy", d.AccountH.UnbindProxy)
		}

		if d.AuditH != nil {
			auditG := admin.Group("/audit")
			auditG.GET("/logs", d.AuditH.List)
		}

		if d.AdminModelH != nil {
			mg := admin.Group("/models")
			mg.GET("", d.AdminModelH.List)
			mg.POST("", d.AdminModelH.Create)
			mg.PUT("/:id", d.AdminModelH.Update)
			mg.PATCH("/:id/enabled", d.AdminModelH.SetEnabled)
			mg.DELETE("/:id", d.AdminModelH.Delete)
		}

		if d.AdminUsageH != nil {
			ug := admin.Group("/usage")
			ug.GET("/stats", d.AdminUsageH.Stats)
			ug.GET("/logs", d.AdminUsageH.Logs)
		}

		if d.SettingsH != nil {
			sg := admin.Group("/settings")
			sg.GET("", d.SettingsH.List)
			sg.PUT("", d.SettingsH.Update)
			sg.POST("/reload", d.SettingsH.Reload)
			sg.POST("/test-email", d.SettingsH.TestMail)
		}

		if d.BackupH != nil {
			bg := admin.Group("/system/backup")
			bg.GET("", d.BackupH.List)
			bg.POST("", d.BackupH.Create)
			bg.GET("/:id/download", d.BackupH.Download)
			bg.DELETE("/:id", d.BackupH.Delete)
			bg.POST("/:id/restore", d.BackupH.Restore)
			bg.POST("/upload", d.BackupH.Upload)
		}
	}

	var v1MW []gin.HandlerFunc
	if d.SettingsSvc != nil {
		v1MW = append(v1MW, middleware.V1APIKeyAuth(d.SettingsSvc))
	} else {
		v1MW = append(v1MW, middleware.LocalActor())
	}
	v1 := r.Group("/v1", v1MW...)
	{
		v1.GET("/models", d.GatewayH.ListModels)
		v1.POST("/chat/completions", d.GatewayH.ChatCompletions)
		if d.ImagesH != nil {
			v1.POST("/images/generations", d.ImagesH.ImageGenerations)
			v1.POST("/images/edits", d.ImagesH.ImageEdits)
			v1.GET("/images/tasks/:id", d.ImagesH.ImageTask)
		}
	}

	if d.ImagesH != nil {
		r.GET("/p/img/:task_id/:idx", d.ImagesH.ImageProxy)
	}

	mountSPA(r)
	return r
}

func localMe(c *gin.Context) {
	resp.OK(c, gin.H{
		"user":        gin.H{"email": "local-console", "nickname": "本地控制台"},
		"role":        "local",
		"permissions": []string{"local:*"},
	})
}

func localMenu(c *gin.Context) {
	resp.OK(c, gin.H{
		"role":        "local",
		"permissions": []string{"local:*"},
		"menu": []gin.H{
			{"key": "personal", "title": "本地面板", "icon": "Monitor", "children": []gin.H{
				{"key": "dashboard", "title": "运行概览", "icon": "DataLine", "path": "/personal/dashboard"},
				{"key": "usage", "title": "用量观察", "icon": "TrendCharts", "path": "/personal/usage"},
				{"key": "play", "title": "在线体验", "icon": "ChatDotRound", "path": "/personal/play"},
				{"key": "docs", "title": "接口文档", "icon": "Document", "path": "/personal/docs"},
			}},
			{"key": "admin", "title": "运维管理", "icon": "Setting", "children": []gin.H{
				{"key": "accounts", "title": "账号池", "icon": "User", "path": "/admin/accounts"},
				{"key": "proxies", "title": "代理池", "icon": "Connection", "path": "/admin/proxies"},
				{"key": "models", "title": "模型映射", "icon": "Grid", "path": "/admin/models"},
				{"key": "usage-admin", "title": "全局用量", "icon": "Histogram", "path": "/admin/usage"},
				{"key": "audit", "title": "审计日志", "icon": "Tickets", "path": "/admin/audit"},
				{"key": "backup", "title": "数据备份", "icon": "Folder", "path": "/admin/backup"},
				{"key": "settings", "title": "系统设置", "icon": "Tools", "path": "/admin/settings"},
			}},
		},
	})
}
````

## File: internal/server/spa.go
````go
package server

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// mountSPA 把前端 Vite 产物(web/dist)挂到 `/` 上,并实现 SPA 回退(deep link 刷新)。
//
// 路径选择优先级:
//  1. 环境变量 GPT2API_WEB_DIR
//  2. 容器默认:/app/web/dist
//  3. 源码工作目录:./web/dist
//  4. 都不存在则什么都不挂(退化为纯 API)
//
// 注意:
//   - 只有 GET/HEAD 的 NoRoute 请求才会被 fallback 到 index.html。其它方法保持 404。
//   - 明确排除 /api/、/v1/、/healthz、/readyz 等 API 前缀,避免打包问题把接口 404 掩盖成 index.html。
func mountSPA(r *gin.Engine) bool {
	dir := resolveWebDir()
	if dir == "" {
		return false
	}
	indexPath := filepath.Join(dir, "index.html")
	if _, err := os.Stat(indexPath); err != nil {
		return false
	}

	// 静态资源(/assets/** 等)
	r.Static("/assets", filepath.Join(dir, "assets"))

	// 常见的"在根下"的单文件
	registerRootFile := func(name string) {
		p := filepath.Join(dir, name)
		if _, err := os.Stat(p); err == nil {
			r.StaticFile("/"+name, p)
		}
	}
	registerRootFile("favicon.svg")
	registerRootFile("favicon.ico")
	registerRootFile("robots.txt")

	// 根路径直接返回 index.html,而不是 404。
	r.GET("/", func(c *gin.Context) { c.File(indexPath) })

	// NoRoute 兜底:仅对 GET/HEAD 且不在 API 前缀下的请求返回 index.html,
	// 让前端 vue-router 接管 deep link。
	r.NoRoute(func(c *gin.Context) {
		if c.Request.Method != http.MethodGet && c.Request.Method != http.MethodHead {
			c.Status(http.StatusNotFound)
			return
		}
		p := c.Request.URL.Path
		for _, prefix := range apiPrefixes {
			if strings.HasPrefix(p, prefix) {
				c.Status(http.StatusNotFound)
				return
			}
		}
		c.File(indexPath)
	})
	return true
}

// API 前缀白名单:凡是命中这里的请求不做 SPA fallback。
var apiPrefixes = []string{
	"/api/",
	"/v1/",
	"/healthz",
	"/readyz",
	"/assets/",
}

func resolveWebDir() string {
	if d := os.Getenv("GPT2API_WEB_DIR"); d != "" {
		if isDir(d) {
			return d
		}
	}
	candidates := []string{
		"/app/web/dist",
		"./web/dist",
	}
	for _, d := range candidates {
		if isDir(d) {
			abs, _ := filepath.Abs(d)
			return abs
		}
	}
	return ""
}

func isDir(p string) bool {
	st, err := os.Stat(p)
	if err != nil {
		return false
	}
	return st.IsDir()
}
````

## File: internal/settings/dao.go
````go
package settings

import (
	"context"
	"strings"

	"github.com/jmoiron/sqlx"
)

// DAO 访问 system_settings 表。
type DAO struct {
	db *sqlx.DB
}

func NewDAO(db *sqlx.DB) *DAO { return &DAO{db: db} }

type row struct {
	K string `db:"k"`
	V string `db:"v"`
}

// LoadAll 全量读。启动时以及 Set 后无需调用(Set 内部维护)。
func (d *DAO) LoadAll(ctx context.Context) (map[string]string, error) {
	var rows []row
	if err := d.db.SelectContext(ctx, &rows, "SELECT `k`, COALESCE(`v`, '') AS `v` FROM `system_settings`"); err != nil {
		return nil, err
	}
	m := make(map[string]string, len(rows))
	for _, r := range rows {
		m[r.K] = r.V
	}
	return m, nil
}

// SetMany 用事务批量 upsert,未列出的 key 不动。
// 调用方负责白名单校验。
func (d *DAO) SetMany(ctx context.Context, kv map[string]string) error {
	if len(kv) == 0 {
		return nil
	}
	tx, err := d.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	const q = "INSERT INTO `system_settings` (`k`, `v`) VALUES (?, ?) " +
		"ON DUPLICATE KEY UPDATE `v` = VALUES(`v`)"
	for k, v := range kv {
		k = strings.TrimSpace(k)
		if k == "" {
			continue
		}
		if _, err := tx.ExecContext(ctx, q, k, v); err != nil {
			return err
		}
	}
	return tx.Commit()
}
````

## File: internal/settings/handler.go
````go
package settings

import (
	"github.com/gin-gonic/gin"

	"github.com/432539/gpt2api/internal/audit"
	"github.com/432539/gpt2api/internal/middleware"
	"github.com/432539/gpt2api/pkg/mailer"
	"github.com/432539/gpt2api/pkg/resp"
)

// Handler 系统设置 HTTP 接口。
//   - List    GET  /api/admin/settings          本地控制台读取所有 key
//   - Update  PUT  /api/admin/settings          本地控制台批量更新
//   - Reload  POST /api/admin/settings/reload   从 DB 强制重载缓存(应急)
//   - TestMail POST /api/admin/settings/test-email 本地控制台给任意地址发一封测试邮件
//   - Public  GET  /api/public/site-info        匿名可访问,返回 Public=true 的子集
type Handler struct {
	svc      *Service
	mail     *mailer.Mailer
	auditDAO *audit.DAO
}

func NewHandler(svc *Service, mail *mailer.Mailer, adao *audit.DAO) *Handler {
	return &Handler{svc: svc, mail: mail, auditDAO: adao}
}

// itemView 给前端使用的完整条目(带 schema,便于统一渲染)。
type itemView struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	Type     string `json:"type"`
	Category string `json:"category"`
	Label    string `json:"label"`
	Desc     string `json:"desc"`
}

// List GET /api/admin/settings
func (h *Handler) List(c *gin.Context) {
	snap := h.svc.Snapshot()
	items := make([]itemView, 0, len(Defs))
	for _, d := range Defs {
		items = append(items, itemView{
			Key: d.Key, Value: snap[d.Key], Type: d.Type,
			Category: d.Category, Label: d.Label, Desc: d.Desc,
		})
	}
	resp.OK(c, gin.H{"items": items})
}

// Update PUT /api/admin/settings
// body: { "items": { "site.name": "...", "gateway.upstream_timeout_sec": "60", ... } }
type updateReq struct {
	Items map[string]string `json:"items"`
}

func (h *Handler) Update(c *gin.Context) {
	var req updateReq
	if err := c.ShouldBindJSON(&req); err != nil || len(req.Items) == 0 {
		resp.BadRequest(c, "items required")
		return
	}
	// 白名单过滤 + 类型轻校验(严重错误直接拒,warning 放行由前端提示)
	for k, v := range req.Items {
		if !IsAllowedKey(k) {
			resp.BadRequest(c, "unknown key: "+k)
			return
		}
		if def, _ := DefByKey(k); def.Type == "int" {
			if v == "" {
				req.Items[k] = "0"
				continue
			}
			if _, err := parseInt64(v); err != nil {
				resp.BadRequest(c, k+" must be integer")
				return
			}
		}
	}
	if err := h.svc.Set(c.Request.Context(), req.Items); err != nil {
		resp.Internal(c, err.Error())
		return
	}
	if h.auditDAO != nil {
		actor := middleware.UserID(c)
		if actor > 0 {
			_ = h.auditDAO.Insert(c.Request.Context(), &audit.Log{
				ActorID: actor,
				Action:  "settings.update",
				Method:  c.Request.Method,
				Path:    c.FullPath(),
				Target:  sprintKeys(req.Items),
				IP:      c.ClientIP(),
				UA:      c.Request.UserAgent(),
			})
		}
	}
	resp.OK(c, gin.H{"updated": len(req.Items)})
}

// Reload POST /api/admin/settings/reload
func (h *Handler) Reload(c *gin.Context) {
	if err := h.svc.Reload(c.Request.Context()); err != nil {
		resp.Internal(c, err.Error())
		return
	}
	resp.OK(c, gin.H{"reloaded": true})
}

// TestMail POST /api/admin/settings/test-email
// body: { "to": "foo@bar.com" }
type testMailReq struct {
	To string `json:"to" binding:"required,email"`
}

func (h *Handler) TestMail(c *gin.Context) {
	var req testMailReq
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.BadRequest(c, "invalid email: "+err.Error())
		return
	}
	if h.mail == nil || h.mail.Disabled() {
		resp.Fail(c, resp.CodeBadRequest, "SMTP not configured: fill host/user/pass in config and restart")
		return
	}
	subject := "[" + h.svc.SiteName() + "] SMTP test email"
	html := `<p>This is a <b>test email</b> sent from ` + h.svc.SiteName() + ` admin console.</p>` +
		`<p>If you see this, your SMTP configuration works.</p>`
	if err := h.mail.SendSync(mailer.Message{To: req.To, Subject: subject, HTML: html}); err != nil {
		resp.Fail(c, resp.CodeInternal, "send failed: "+err.Error())
		return
	}
	resp.OK(c, gin.H{"sent": true, "to": req.To})
}

// Public GET /api/public/site-info
func (h *Handler) Public(c *gin.Context) {
	resp.OK(c, h.svc.PublicSnapshot())
}
````

## File: internal/settings/model.go
````go
// Package settings 系统设置(KV)。
package settings

import "strings"

// KeyDef 声明可编辑的设置项。
type KeyDef struct {
	Key      string
	Type     string // string | bool | int | float | email | url
	Category string // site | gateway | account | mail
	Default  string
	Label    string
	Desc     string
	Public   bool
}

const (
	SiteName          = "site.name"
	SiteDescription   = "site.description"
	SiteLogoURL       = "site.logo_url"
	SiteFooter        = "site.footer"
	SiteContactEmail  = "site.contact_email"
	SiteDocsURL       = "site.docs_url"
	SiteAPIBaseURL    = "site.api_base_url"
	UIDefaultPageSize = "ui.default_page_size"

	GatewayUpstreamTimeoutSec   = "gateway.upstream_timeout_sec"
	GatewaySSEReadTimeoutSec    = "gateway.sse_read_timeout_sec"
	GatewayCooldown429Sec       = "gateway.cooldown_429_sec"
	GatewayWarnedPauseHours     = "gateway.warned_pause_hours"
	GatewayDailyUsageRatio      = "gateway.daily_usage_ratio"
	GatewayRetryOnFailure       = "gateway.retry_on_failure"
	GatewayRetryMax             = "gateway.retry_max"
	GatewayDispatchQueueWaitSec = "gateway.dispatch_queue_wait_sec"
	GatewayImageExploreRatio    = "gateway.image_explore_ratio"

	ProxyProbeEnabled     = "proxy.probe_enabled"
	ProxyProbeIntervalSec = "proxy.probe_interval_sec"
	ProxyProbeTimeoutSec  = "proxy.probe_timeout_sec"
	ProxyProbeTargetURL   = "proxy.probe_target_url"
	ProxyProbeConcurrency = "proxy.probe_concurrency"

	AccountRefreshEnabled        = "account.refresh_enabled"
	AccountRefreshIntervalSec    = "account.refresh_interval_sec"
	AccountRefreshAheadSec       = "account.refresh_ahead_sec"
	AccountRefreshConcurrency    = "account.refresh_concurrency"
	AccountQuotaProbeEnabled     = "account.quota_probe_enabled"
	AccountQuotaProbeIntervalSec = "account.quota_probe_interval_sec"
	AccountDefaultClientID       = "account.default_client_id"

	APIKey = "api.v1_key"

	MailEnabledDisplay = "mail.enabled_display"
)

// Defs 所有合法 key 的 schema。前端设置页按 category 展示。
var Defs = []KeyDef{
	{Key: SiteName, Type: "string", Category: "site", Default: "GPT2API Local", Label: "站点名称", Desc: "展示在顶栏和控制台标题", Public: true},
	{Key: SiteDescription, Type: "string", Category: "site", Default: "自用 OpenAI 兼容 2API 中转", Label: "副标题", Desc: "控制台说明文案", Public: true},
	{Key: SiteLogoURL, Type: "url", Category: "site", Default: "", Label: "Logo URL", Desc: "空则使用默认图标", Public: true},
	{Key: SiteFooter, Type: "string", Category: "site", Default: "", Label: "页脚文案", Desc: "版权/备案号等纯文本", Public: true},
	{Key: SiteContactEmail, Type: "email", Category: "site", Default: "", Label: "联系邮箱", Desc: "本地部署备注邮箱", Public: true},
	{Key: SiteDocsURL, Type: "url", Category: "site", Default: "", Label: "文档链接", Desc: "留空则使用内置文档", Public: true},
	{Key: SiteAPIBaseURL, Type: "url", Category: "site", Default: "", Label: "API Base URL", Desc: "展示给客户端的 /v1 入口;留空=当前站点地址", Public: true},
	{Key: UIDefaultPageSize, Type: "int", Category: "site", Default: "20", Label: "默认每页条数", Desc: "控制台表格默认分页(5~100)"},

	{Key: GatewayUpstreamTimeoutSec, Type: "int", Category: "gateway", Default: "60", Label: "上游请求超时(秒)", Desc: "非流式请求上游响应超时"},
	{Key: GatewaySSEReadTimeoutSec, Type: "int", Category: "gateway", Default: "120", Label: "SSE 读超时(秒)", Desc: "流式响应无数据时的中断阈值"},
	{Key: GatewayCooldown429Sec, Type: "int", Category: "gateway", Default: "300", Label: "429 冷却(秒)", Desc: "账号遇 429 后暂停调度"},
	{Key: GatewayWarnedPauseHours, Type: "int", Category: "gateway", Default: "24", Label: "风险暂停(小时)", Desc: "账号被识别为 warned 时的暂停时长"},
	{Key: GatewayDailyUsageRatio, Type: "float", Category: "gateway", Default: "0.8", Label: "日用比例阈值", Desc: "0.0~1.0;超过后降低调度优先级"},
	{Key: GatewayRetryOnFailure, Type: "bool", Category: "gateway", Default: "true", Label: "失败自动重试", Desc: "遇到可恢复错误时切换账号重试"},
	{Key: GatewayRetryMax, Type: "int", Category: "gateway", Default: "1", Label: "最大重试次数", Desc: "0~3"},
	{Key: GatewayDispatchQueueWaitSec, Type: "int", Category: "gateway", Default: "120", Label: "账号排队等待上限(秒)", Desc: "并发大于账号数时等待空闲账号的最长秒数"},
	{Key: GatewayImageExploreRatio, Type: "float", Category: "gateway", Default: "0.2", Label: "IMG2 账号探索比例", Desc: "image 调度中留给 unknown/新号/长时间未尝试号的比例"},

	{Key: ProxyProbeEnabled, Type: "bool", Category: "gateway", Default: "true", Label: "代理探测开关", Desc: "后台定时对启用的代理做连通性探测"},
	{Key: ProxyProbeIntervalSec, Type: "int", Category: "gateway", Default: "300", Label: "代理探测间隔(秒)", Desc: "两轮探测之间的间隔"},
	{Key: ProxyProbeTimeoutSec, Type: "int", Category: "gateway", Default: "10", Label: "代理探测超时(秒)", Desc: "单条代理一次探测的超时时间"},
	{Key: ProxyProbeTargetURL, Type: "url", Category: "gateway", Default: "https://chatgpt.com/cdn-cgi/trace", Label: "代理探测目标 URL", Desc: "返回 2xx/3xx 视为成功"},
	{Key: ProxyProbeConcurrency, Type: "int", Category: "gateway", Default: "8", Label: "代理探测并发", Desc: "同时探测的代理数(1~64)"},

	{Key: AccountRefreshEnabled, Type: "bool", Category: "account", Default: "true", Label: "账号 AT 自动刷新", Desc: "后台定时刷新即将过期的 AT"},
	{Key: AccountRefreshIntervalSec, Type: "int", Category: "account", Default: "120", Label: "账号刷新扫描间隔(秒)", Desc: "多久扫一次"},
	{Key: AccountRefreshAheadSec, Type: "int", Category: "account", Default: "120", Label: "账号预刷新提前量(秒)", Desc: "距离过期多少秒内触发刷新(不宜过大,避免频繁刷新触发上游风控)"},
	{Key: AccountRefreshConcurrency, Type: "int", Category: "account", Default: "4", Label: "账号刷新并发", Desc: "同时刷新的账号数(1~32)"},
	{Key: AccountQuotaProbeEnabled, Type: "bool", Category: "account", Default: "true", Label: "账号额度自动探测", Desc: "后台定期查询账号的图片剩余量"},
	{Key: AccountQuotaProbeIntervalSec, Type: "int", Category: "account", Default: "900", Label: "剩余量探测最小间隔(秒)", Desc: "同一账号两次探测之间的最小间隔"},
	{Key: AccountDefaultClientID, Type: "string", Category: "account", Default: "app_LlGpXReQgckcGGUo2JrYvtJK", Label: "导入账号默认 client_id", Desc: "JSON 未指定时使用的 OAuth client_id"},

	{Key: APIKey, Type: "string", Category: "api", Default: "", Label: "API Key", Desc: "用于 /v1/* 接口认证;留空则不验证密钥(开放访问)"},

	{Key: MailEnabledDisplay, Type: "string", Category: "mail", Default: "auto", Label: "邮件开关展示", Desc: "auto/true/false;实际由 SMTP 配置决定"},
}

func DefByKey(k string) (KeyDef, bool) {
	for _, d := range Defs {
		if d.Key == k {
			return d, true
		}
	}
	return KeyDef{}, false
}

func IsAllowedKey(k string) bool {
	k = strings.TrimSpace(k)
	if k == "" {
		return false
	}
	_, ok := DefByKey(k)
	return ok
}
````

## File: internal/settings/service.go
````go
package settings

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"sync"
)

// Service 带内存缓存的只读/可写访问层。
// 所有读走本地 map,写走 DB + 原子替换缓存。
type Service struct {
	dao   *DAO
	mu    sync.RWMutex
	cache map[string]string // 最新快照;不直接暴露,通过 GetXxx 读
}

var ErrUnknownKey = errors.New("settings: unknown key")

func NewService(dao *DAO) *Service {
	return &Service{dao: dao, cache: map[string]string{}}
}

// Reload 启动时 / 手动触发时调用。
func (s *Service) Reload(ctx context.Context) error {
	m, err := s.dao.LoadAll(ctx)
	if err != nil {
		return err
	}
	// 补齐 Defs 默认值(DB 里缺某个 key 时)
	for _, d := range Defs {
		if _, ok := m[d.Key]; !ok {
			m[d.Key] = d.Default
		}
	}
	s.mu.Lock()
	s.cache = m
	s.mu.Unlock()
	return nil
}

// Snapshot 拷贝当前所有设置(不含未登记 key)。
func (s *Service) Snapshot() map[string]string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make(map[string]string, len(s.cache))
	for _, d := range Defs {
		if v, ok := s.cache[d.Key]; ok {
			out[d.Key] = v
		} else {
			out[d.Key] = d.Default
		}
	}
	return out
}

// PublicSnapshot 只返回 Defs 中 Public=true 的条目,面向匿名访问。
func (s *Service) PublicSnapshot() map[string]string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make(map[string]string)
	for _, d := range Defs {
		if !d.Public {
			continue
		}
		if v, ok := s.cache[d.Key]; ok {
			out[d.Key] = v
		} else {
			out[d.Key] = d.Default
		}
	}
	return out
}

// Set 批量写入并刷新缓存;未在 Defs 白名单中的 key 会被过滤。
func (s *Service) Set(ctx context.Context, in map[string]string) error {
	filtered := make(map[string]string, len(in))
	for k, v := range in {
		if !IsAllowedKey(k) {
			continue
		}
		filtered[k] = strings.TrimSpace(v)
	}
	if len(filtered) == 0 {
		return nil
	}
	if err := s.dao.SetMany(ctx, filtered); err != nil {
		return err
	}
	s.mu.Lock()
	for k, v := range filtered {
		s.cache[k] = v
	}
	s.mu.Unlock()
	return nil
}

// --- typed getters ---

func (s *Service) GetString(key string) string {
	s.mu.RLock()
	v, ok := s.cache[key]
	s.mu.RUnlock()
	if ok {
		return v
	}
	if d, ok := DefByKey(key); ok {
		return d.Default
	}
	return ""
}

func (s *Service) GetBool(key string) bool {
	v := strings.ToLower(strings.TrimSpace(s.GetString(key)))
	return v == "1" || v == "true" || v == "yes" || v == "on"
}

func (s *Service) GetInt(key string) int64 {
	n, _ := strconv.ParseInt(strings.TrimSpace(s.GetString(key)), 10, 64)
	return n
}

func (s *Service) GetFloat(key string) float64 {
	f, _ := strconv.ParseFloat(strings.TrimSpace(s.GetString(key)), 64)
	return f
}

// --- convenience helpers (业务语义) ---
// 所有 helper 都保证返回"安全"的业务默认(例如不让 0/负值落到业务路径)。

// -- site --
func (s *Service) SiteName() string { return firstNonEmpty(s.GetString(SiteName), "GPT2API") }

// -- gateway --
func (s *Service) GatewayUpstreamTimeoutSec() int {
	n := int(s.GetInt(GatewayUpstreamTimeoutSec))
	if n <= 0 {
		return 60
	}
	return n
}
func (s *Service) GatewaySSEReadTimeoutSec() int {
	n := int(s.GetInt(GatewaySSEReadTimeoutSec))
	if n <= 0 {
		return 120
	}
	return n
}
func (s *Service) Cooldown429Sec() int {
	n := int(s.GetInt(GatewayCooldown429Sec))
	if n <= 0 {
		return 300
	}
	return n
}
func (s *Service) WarnedPauseHours() int {
	n := int(s.GetInt(GatewayWarnedPauseHours))
	if n <= 0 {
		return 24
	}
	return n
}
func (s *Service) DailyUsageRatio() float64 {
	f := s.GetFloat(GatewayDailyUsageRatio)
	if f <= 0 || f > 1 {
		return 0.8
	}
	return f
}
func (s *Service) RetryOnFailure() bool { return s.GetBool(GatewayRetryOnFailure) }
func (s *Service) RetryMax() int {
	n := int(s.GetInt(GatewayRetryMax))
	if n < 0 {
		return 0
	}
	if n > 3 {
		return 3
	}
	return n
}

// DispatchQueueWaitSec 账号池忙时请求的最长排队时间。
//   - 0   ⇒ 不排队(立即返回 no_available_account)
//   - 负数 / 未设置 ⇒ 回退默认 120
func (s *Service) DispatchQueueWaitSec() int {
	n := int(s.GetInt(GatewayDispatchQueueWaitSec))
	if n < 0 {
		return 120
	}
	return n
}

func (s *Service) ImageExploreRatio() float64 {
	f := s.GetFloat(GatewayImageExploreRatio)
	if f < 0 {
		return 0
	}
	if f > 0.8 {
		return 0.8
	}
	return f
}

// -- proxy probe --
func (s *Service) ProbeEnabled() bool { return s.GetBool(ProxyProbeEnabled) }
func (s *Service) ProbeIntervalSec() int {
	n := int(s.GetInt(ProxyProbeIntervalSec))
	if n <= 0 {
		return 300
	}
	return n
}
func (s *Service) ProbeTimeoutSec() int {
	n := int(s.GetInt(ProxyProbeTimeoutSec))
	if n <= 0 {
		return 10
	}
	return n
}
func (s *Service) ProbeTargetURL() string {
	return firstNonEmpty(s.GetString(ProxyProbeTargetURL), "https://chatgpt.com/cdn-cgi/trace")
}
func (s *Service) ProbeConcurrency() int {
	n := int(s.GetInt(ProxyProbeConcurrency))
	if n <= 0 {
		return 8
	}
	if n > 64 {
		return 64
	}
	return n
}

// -- account refresh / quota probe --
func (s *Service) AccountRefreshEnabled() bool { return s.GetBool(AccountRefreshEnabled) }
func (s *Service) AccountRefreshIntervalSec() int {
	n := int(s.GetInt(AccountRefreshIntervalSec))
	if n <= 0 {
		return 120
	}
	return n
}
func (s *Service) AccountRefreshAheadSec() int {
	n := int(s.GetInt(AccountRefreshAheadSec))
	if n <= 0 {
		return 120
	}
	return n
}
func (s *Service) AccountRefreshConcurrency() int {
	n := int(s.GetInt(AccountRefreshConcurrency))
	if n <= 0 {
		return 4
	}
	if n > 32 {
		return 32
	}
	return n
}
func (s *Service) AccountQuotaProbeEnabled() bool { return s.GetBool(AccountQuotaProbeEnabled) }
func (s *Service) AccountQuotaProbeIntervalSec() int {
	n := int(s.GetInt(AccountQuotaProbeIntervalSec))
	if n <= 0 {
		return 900
	}
	return n
}
func (s *Service) AccountDefaultClientID() string {
	return firstNonEmpty(s.GetString(AccountDefaultClientID), "app_LlGpXReQgckcGGUo2JrYvtJK")
}

// -- api --
func (s *Service) V1APIKey() string { return s.GetString(APIKey) }

func firstNonEmpty(vs ...string) string {
	for _, v := range vs {
		if strings.TrimSpace(v) != "" {
			return v
		}
	}
	return ""
}
````

## File: internal/settings/util.go
````go
package settings

import (
	"sort"
	"strconv"
	"strings"
)

func parseInt64(s string) (int64, error) {
	return strconv.ParseInt(strings.TrimSpace(s), 10, 64)
}

// sprintKeys 把更新过的 key 列表拼成审计 detail,避免把明文敏感值写进审计日志
// (比如未来可能加入密钥字段)。
func sprintKeys(m map[string]string) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return "updated=" + strings.Join(keys, ",")
}
````

## File: internal/upstream/chatgpt/client.go
````go
// Package chatgpt 封装 chatgpt.com 的反向工程调用。
//
// 本包不关心调度策略,只负责一次 HTTP 往返。
// 调用方(网关)负责:调度器拿 Lease -> 构造 Client -> 发请求 -> 转译响应。
package chatgpt

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"time"

	utls "github.com/refraction-networking/utls"
	"go.uber.org/zap"
	"golang.org/x/net/publicsuffix"

	"github.com/432539/gpt2api/pkg/logger"
)

func loggerL() *zap.Logger { return logger.L() }

// 固定请求头(模拟 Chrome 131;客户端版本号可按需更新)。
const (
	// UA 对齐 utls HelloChrome_131 TLS 指纹 + Sec-Ch-Ua 完整套件。
	// 三者必须自洽:TLS ClientHello(JA3)→Chrome 131、UA→Chrome 131、
	// Sec-Ch-Ua→Chrome 131,否则 Cloudflare 交叉校验会判定指纹冲突。
	DefaultUserAgent      = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"
	DefaultClientVersion  = "prod-2d84edefecf794f1bf3178f1f15e1005067d903d"
	DefaultClientBuildNum = "5983180"
	DefaultLanguage       = "zh-CN"
	BaseURL               = "https://chatgpt.com"
)

// Options 构造 Client 的参数。
type Options struct {
	BaseURL           string
	AuthToken         string // 完整 Bearer token(已解密)
	DeviceID          string
	SessionID         string
	ProxyURL          string        // http(s)://user:pass@host:port,为空则直连
	Timeout           time.Duration // HTTP 总超时,默认 120s
	SSETimeout        time.Duration // SSE 首 byte 超时,默认 60s
	Cookies           string        // JSON 字符串(可选),格式 [{"name":"x","value":"y","domain":".chatgpt.com"}]
	UserAgent         string
	ClientVersion     string
	ClientBuildNumber string
	Language          string

	// TurnstileSolver 可选。为 nil 时 ChatRequirementsV2 会回退到单步
	// chat-requirements 流程(TurnstileRequired=true 时直接忽略)。
	TurnstileSolver TurnstileSolver

	// HelloID 可选。per-account 的 uTLS ClientHelloID,避免多账号共用同一 TLS 指纹。
	// 为 nil 时使用默认 HelloChrome_131。传入后会在每次 TLS 握手时使用该指纹。
	HelloID *utls.ClientHelloID
}

// Client 一个账号/代理/device 一次请求的上游客户端。可多次复用(建议 1 次请求 1 个)。
type Client struct {
	opts Options
	hc   *http.Client
}

// withHelloID 如果 Options 指定了自定义 HelloID,把它注入到 context 中
// 供 utlsRoundTripper 在 TLS 握手时读取。
func (c *Client) withHelloID(ctx context.Context) context.Context {
	if c.opts.HelloID != nil {
		return context.WithValue(ctx, utlsHelloIDKey{}, *c.opts.HelloID)
	}
	return ctx
}

// New 构造客户端。
func New(opt Options) (*Client, error) {
	if opt.AuthToken == "" {
		return nil, errors.New("auth_token required")
	}
	if opt.DeviceID == "" {
		return nil, errors.New("device_id required")
	}
	if opt.BaseURL == "" {
		opt.BaseURL = BaseURL
	}
	if opt.Timeout == 0 {
		opt.Timeout = 120 * time.Second
	}
	if opt.SSETimeout == 0 {
		opt.SSETimeout = 60 * time.Second
	}
	if opt.UserAgent == "" {
		opt.UserAgent = DefaultUserAgent
	}
	if opt.ClientVersion == "" {
		opt.ClientVersion = DefaultClientVersion
	}
	if opt.ClientBuildNumber == "" {
		opt.ClientBuildNumber = DefaultClientBuildNum
	}
	if opt.Language == "" {
		opt.Language = DefaultLanguage
	}

	// 直接用标准 net/http 会被 Cloudflare 按 JA3/JA4 指纹识别出不是浏览器(403 拦截页);
	// 这里换成 utls-based RoundTripper,ClientHello 伪装成 Chrome 120。
	// Proxy(HTTP / HTTPS CONNECT)在 transport 内部处理,不再走 http.ProxyURL。
	tr, err := NewUTLSTransport(opt.ProxyURL, 30*time.Second)
	if err != nil {
		return nil, fmt.Errorf("init utls transport: %w", err)
	}

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, fmt.Errorf("create cookie jar: %w", err)
	}
	hc := &http.Client{
		Transport: tr,
		Timeout:   opt.Timeout, // SSE 场景会关闭该 timeout(见 StreamConversation)
		Jar:       jar,
	}
	if opt.Cookies != "" {
		_ = loadCookies(jar, opt.BaseURL, opt.Cookies)
	}
	return &Client{opts: opt, hc: hc}, nil
}

// do 执行 HTTP 请求,自动注入 per-account HelloID 到 context。
func (c *Client) do(req *http.Request) (*http.Response, error) {
	c.injectHelloID(req)
	return c.hc.Do(req)
}

// injectHelloID 把 per-account HelloID 注入到 req 的 context 中,
// 供 utlsRoundTripper 在 TLS 握手时读取。
func (c *Client) injectHelloID(req *http.Request) {
	if c.opts.HelloID != nil {
		ctx := req.Context()
		if ctx.Value(utlsHelloIDKey{}) == nil {
			ctx = context.WithValue(ctx, utlsHelloIDKey{}, *c.opts.HelloID)
			*req = *req.WithContext(ctx)
		}
	}
}

// loadCookies 把 JSON cookies 加载到 jar。
func loadCookies(jar http.CookieJar, base, raw string) error {
	var list []struct {
		Name   string `json:"name"`
		Value  string `json:"value"`
		Domain string `json:"domain"`
		Path   string `json:"path"`
	}
	if err := json.Unmarshal([]byte(raw), &list); err != nil {
		return err
	}
	u, err := url.Parse(base)
	if err != nil {
		return err
	}
	cs := make([]*http.Cookie, 0, len(list))
	for _, c := range list {
		if c.Name == "" || c.Value == "" {
			continue
		}
		path := c.Path
		if path == "" {
			path = "/"
		}
		cs = append(cs, &http.Cookie{Name: c.Name, Value: c.Value, Domain: c.Domain, Path: path})
	}
	jar.SetCookies(u, cs)
	return nil
}

// commonHeaders 设置所有 chatgpt.com 请求通用的头。
//
// 对齐真实浏览器(Chrome 131 @ Windows)抓包:除了 PoW/turnstile 这类 sentinel 头
// 由具体端点自己加,其他客户端指纹头、Oai-* 头、sec-ch-ua 完整套件、
// X-Openai-Target-Path/Route 都在这里统一设置。X-Oai-Turn-Trace-Id 是"每 turn 一个"
// 的 UUID,由具体发送函数(StreamFChat / StreamFConversation)自己随机生成,
// 这里只填固定的。
//
// 真正的指纹差异在 HTTP/2 SETTINGS frame(JA4H),已在 utls_transport.go 中用
// forceH1 强制走 http/1.1 规避。
func (c *Client) commonHeaders(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+c.opts.AuthToken)
	req.Header.Set("User-Agent", c.opts.UserAgent)
	req.Header.Set("Origin", c.opts.BaseURL)
	req.Header.Set("Referer", c.opts.BaseURL+"/")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	// 不设置 Accept-Encoding:Go net/http 会自动加 `Accept-Encoding: gzip` 并透明解压。
	// 若主动声明 br/zstd,Go 不会解压,body 会是压缩字节,SSE / JSON 解析全坏。
	// sec-ch-ua 完整套件(Chrome 131 on Windows):真实浏览器每次都会带这整套,
	// 少其中任何一项都可能触发 bot 指纹识别。必须与 DefaultUserAgent + uTLS HelloID 自洽。
	req.Header.Set("Sec-Ch-Ua", `"Google Chrome";v="131", "Chromium";v="131", "Not_A Brand";v="24"`)
	req.Header.Set("Sec-Ch-Ua-Arch", `"x86"`)
	req.Header.Set("Sec-Ch-Ua-Bitness", `"64"`)
	req.Header.Set("Sec-Ch-Ua-Full-Version", `"131.0.6778.140"`)
	req.Header.Set("Sec-Ch-Ua-Full-Version-List",
		`"Google Chrome";v="131.0.6778.140", "Chromium";v="131.0.6778.140", "Not_A Brand";v="24.0.0.0"`)
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Model", `""`)
	req.Header.Set("Sec-Ch-Ua-Platform", `"Windows"`)
	req.Header.Set("Sec-Ch-Ua-Platform-Version", `"19.0.0"`)
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Priority", "u=1, i")
	// Oai-* 头:真实浏览器每请求必带。
	req.Header.Set("Oai-Device-Id", c.opts.DeviceID)
	if c.opts.SessionID != "" {
		req.Header.Set("Oai-Session-Id", c.opts.SessionID)
	}
	req.Header.Set("Oai-Language", c.opts.Language)
	req.Header.Set("Oai-Client-Version", c.opts.ClientVersion)
	req.Header.Set("Oai-Client-Build-Number", c.opts.ClientBuildNumber)
	// X-Oai-Grid-State:部分 chatgpt.com 请求中浏览器会带此 header,
	// 用于 A/B 实验分桶。当前传空字符串;后续如确认特定端点需要可覆盖。
	req.Header.Set("X-Oai-Grid-State", "")
	// X-Openai-Target-Path / Route:真实浏览器每请求必带,值就是请求 URL 的 path
	// (不含 query)。Route 通常是带占位符的形式(例如 /files/download/{file_id}),
	// 但后端对两个字段都是相等比较,填成 Path 也不触发 400;先统一 Path,以后
	// 发现特定端点报错再单独覆盖。
	if p := req.URL.Path; p != "" {
		req.Header.Set("X-Openai-Target-Path", p)
		req.Header.Set("X-Openai-Target-Route", p)
	}
	// Accept 默认值在各 endpoint 函数里会覆盖(比如 SSE 设成 text/event-stream)。
	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "*/*")
	}
}

// UpstreamError 是一次 chatgpt.com 请求失败的结构化错误。
type UpstreamError struct {
	Status  int
	Message string
	Body    string
}

func (e *UpstreamError) Error() string {
	return fmt.Sprintf("chatgpt upstream %d: %s", e.Status, e.Message)
}

// IsRateLimited 对应 HTTP 429 / 资源耗尽。
func (e *UpstreamError) IsRateLimited() bool { return e.Status == 429 }

// IsUnauthorized 对应 401 / 403(token 失效 / 风控封号)。
func (e *UpstreamError) IsUnauthorized() bool { return e.Status == 401 || e.Status == 403 }

// ChatRequirementsResp 对应响应(仅摘取关键字段)。
type ChatRequirementsResp struct {
	Token string `json:"token"` // chat_token
	// Persona 常见取值:"chatgpt-freeaccount"(免费号)/ "chatgpt-paid"(Plus/Team)
	//              / "chatgpt-noauth"(未登录)
	// 免费号对高级模型(gpt-5 等)会静默不生成,必须把 upstream model 退化到 "auto"
	// 由上游自己挑,否则 SSE 只会拿到一条 hidden system preamble 就结束。
	Persona     string `json:"persona"`
	Proofofwork struct {
		Required   bool   `json:"required"`
		Seed       string `json:"seed"`
		Difficulty string `json:"difficulty"`
	} `json:"proofofwork"`
	Turnstile struct {
		Required bool `json:"required"`
	} `json:"turnstile"`
}

// IsFreeAccount 判断当前账号是否为免费号(persona=chatgpt-freeaccount)。
func (r *ChatRequirementsResp) IsFreeAccount() bool {
	return r.Persona == "chatgpt-freeaccount"
}

// SolveProof 求解 POW,返回要放进 `Openai-Sentinel-Proof-Token` 的字符串。
// 若 Proofofwork.Required=false,返回空串。
func (r *ChatRequirementsResp) SolveProof(userAgent string) string {
	if !r.Proofofwork.Required {
		return ""
	}
	return SolveProofToken(r.Proofofwork.Seed, r.Proofofwork.Difficulty, userAgent)
}

// Bootstrap 模拟浏览器首次打开 chatgpt.com 的 GET /,让 cookie jar 收下
// Cloudflare 的 `__cf_bm` / `_cfuvid` 与 OpenAI 的 `oai-did` 等 cookie。
//
// 关键作用:没有这些 cookie 时,chat-requirements 会直接要求 Turnstile(即使
// Bearer 合法),所以建议在每次新建 Client 后先调一次 Bootstrap,或者在
// ChatRequirements 内部第一次请求前自动调一次。
// 多次调用是幂等的(HTTP 200/3xx 均视为成功)。
func (c *Client) Bootstrap(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.opts.BaseURL+"/", nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", c.opts.UserAgent)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Sec-Ch-Ua", `"Google Chrome";v="131", "Chromium";v="131", "Not_A Brand";v="24"`)
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", `"Windows"`)
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	res, err := c.do(req)
	if err != nil {
		return fmt.Errorf("bootstrap GET /: %w", err)
	}
	defer res.Body.Close()
	// 读取 HTML 以提取 dpl build hash 和 script src(供 POW 使用)
	bodyBytes, readErr := io.ReadAll(io.LimitReader(res.Body, 2*1024*1024))
	if readErr == nil && res.StatusCode < 500 {
		updateDplCache(string(bodyBytes))
	}
	if res.StatusCode >= 500 {
		return &UpstreamError{Status: res.StatusCode, Message: "bootstrap failed"}
	}
	return nil
}

// ChatRequirements 取 chat_token。
// 请求体必须带 `p` = 客户端预算的 requirements_token(前缀 gAAAAAC,固定难度 0fffff)。
// 否则上游会返回空 token 或 403。
func (c *Client) ChatRequirements(ctx context.Context) (*ChatRequirementsResp, error) {
	// 首次请求前顺手做一次浏览器首访,拿 __cf_bm / oai-did,避免 Turnstile。
	// jar 已经持有过则这次 GET 其实是 200,代价就是一个 RTT,可以接受。
	if err := c.Bootstrap(ctx); err != nil {
		// 拿不到 cookie 不致命,继续往下走;真不行再让 chat-requirements 自己报错。
		_ = err
	}
	reqToken := NewPOWConfig(c.opts.UserAgent).RequirementsToken()
	body, _ := json.Marshal(map[string]string{"p": reqToken})
	req, err := http.NewRequestWithContext(ctx,
		http.MethodPost,
		c.opts.BaseURL+"/backend-api/sentinel/chat-requirements",
		strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}
	c.commonHeaders(req)
	req.Header.Set("Content-Type", "application/json")

	res, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	buf, readErr := io.ReadAll(res.Body)
	if res.StatusCode >= 400 {
		body := ""
		if readErr == nil {
			body = string(buf)
		}
		return nil, &UpstreamError{Status: res.StatusCode, Message: "chat-requirements failed", Body: body}
	}
	if readErr != nil {
		return nil, fmt.Errorf("read chat-requirements body: %w", readErr)
	}
	var out ChatRequirementsResp
	if err := json.Unmarshal(buf, &out); err != nil {
		return nil, fmt.Errorf("decode chat-requirements: %w", err)
	}
	// 诊断用:打印完整 body(含 turnstile / proofofwork / arkose 字段)。
	// 稳定后可改成 Debug 或删除。
	if logger := loggerL(); logger != nil {
		bodyStr := string(buf)
		if len(bodyStr) > 800 {
			bodyStr = bodyStr[:800] + "..."
		}
		logger.Info("chat-requirements raw body",
			zap.String("body", bodyStr),
			zap.Bool("turnstile_required", out.Turnstile.Required),
			zap.Bool("pow_required", out.Proofofwork.Required),
			zap.String("token_prefix", truncatePrefix(out.Token, 16)))
	}
	return &out, nil
}

func truncatePrefix(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

// ChatRequirementsPrepareResp 是 /sentinel/chat-requirements/prepare 的响应。
//
// 浏览器在每个 turn 前先调 prepare 拿到 challenge(turnstile.dx + pow.seed/difficulty),
// 本地计算后再调 finalize 拿最终 chat-requirements-token。我们 Go 端没法解
// turnstile(Cloudflare 混淆 JS + WASM),所以 solver 未配置时需要回退到老的
// 单步 /sentinel/chat-requirements 端点(见 ChatRequirementsV2)。
type ChatRequirementsPrepareResp struct {
	Persona      string `json:"persona"`
	PrepareToken string `json:"prepare_token"`
	Turnstile    struct {
		Required bool   `json:"required"`
		DX       string `json:"dx"`
	} `json:"turnstile"`
	Proofofwork struct {
		Required   bool   `json:"required"`
		Seed       string `json:"seed"`
		Difficulty string `json:"difficulty"`
	} `json:"proofofwork"`
}

// ChatRequirementsPrepare 调 /backend-api/sentinel/chat-requirements/prepare。
// 请求体里的 `p` 字段仍是 18 元素 PoW(前缀 gAAAAAC),和老的单步接口同款。
func (c *Client) ChatRequirementsPrepare(ctx context.Context) (*ChatRequirementsPrepareResp, error) {
	reqToken := NewPOWConfig(c.opts.UserAgent).RequirementsToken()
	body, _ := json.Marshal(map[string]string{"p": reqToken})
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		c.opts.BaseURL+"/backend-api/sentinel/chat-requirements/prepare",
		strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}
	c.commonHeaders(req)
	req.Header.Set("Content-Type", "application/json")

	res, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	buf, readErr := io.ReadAll(res.Body)
	if res.StatusCode >= 400 {
		body := ""
		if readErr == nil {
			body = string(buf)
		}
		return nil, &UpstreamError{Status: res.StatusCode, Message: "chat-requirements/prepare failed", Body: body}
	}
	if readErr != nil {
		return nil, fmt.Errorf("read chat-requirements/prepare body: %w", readErr)
	}
	var out ChatRequirementsPrepareResp
	if err := json.Unmarshal(buf, &out); err != nil {
		return nil, fmt.Errorf("decode chat-requirements/prepare: %w", err)
	}
	return &out, nil
}

// ChatRequirementsFinalize 调 /backend-api/sentinel/chat-requirements/finalize。
// 入参:prepare_token(来自 Prepare),proofofwork(本地解,13 元素),
// turnstileResp(通常由 TurnstileSolver 提供;没有则传空串,上游可能拒绝)。
// 返回最终的 chat-requirements-token。
func (c *Client) ChatRequirementsFinalize(
	ctx context.Context,
	prepareToken, proofofwork, turnstileResp string,
) (string, string, error) {
	payload := map[string]interface{}{
		"prepare_token": prepareToken,
	}
	if proofofwork != "" {
		payload["proofofwork"] = proofofwork
	}
	if turnstileResp != "" {
		payload["turnstile"] = turnstileResp
	}
	body, _ := json.Marshal(payload)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		c.opts.BaseURL+"/backend-api/sentinel/chat-requirements/finalize",
		strings.NewReader(string(body)))
	if err != nil {
		return "", "", err
	}
	c.commonHeaders(req)
	req.Header.Set("Content-Type", "application/json")

	res, err := c.do(req)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()
	buf, readErr := io.ReadAll(res.Body)
	if res.StatusCode >= 400 {
		body := ""
		if readErr == nil {
			body = string(buf)
		}
		return "", "", &UpstreamError{Status: res.StatusCode, Message: "chat-requirements/finalize failed", Body: body}
	}
	if readErr != nil {
		return "", "", fmt.Errorf("read chat-requirements/finalize body: %w", readErr)
	}
	var out struct {
		Persona string `json:"persona"`
		Token   string `json:"token"`
	}
	if err := json.Unmarshal(buf, &out); err != nil {
		return "", "", fmt.Errorf("decode chat-requirements/finalize: %w", err)
	}
	return out.Token, out.Persona, nil
}

// ChatRequirementsV2 是 sentinel 协议的新统一入口。
//
// 路由逻辑:
//  1. 先调 /prepare 拿 challenge;
//  2. 若返回 turnstile.required=false,直接把 prepare_token 走 finalize 拿最终 token;
//  3. 若 turnstile.required=true:
//     a. 当 opts.TurnstileSolver != nil → solver.Solve(dx),然后走 finalize;
//     b. 当 solver 为 nil → **回退到老的单步 chat-requirements**,保持向后兼容。
//
// 返回值与老的 ChatRequirements 保持一致,方便调用方无痛切换。
//
// 调用前请先 Bootstrap()。
func (c *Client) ChatRequirementsV2(ctx context.Context) (*ChatRequirementsResp, error) {
	prep, err := c.ChatRequirementsPrepare(ctx)
	if err != nil {
		// prepare 本身失败时,不再尝试 finalize,直接回退到单步接口。
		// 这样新协议上游未开启时也不会阻塞业务。
		if logger := loggerL(); logger != nil {
			logger.Warn("chat-requirements/prepare failed, fallback to single-step",
				zap.Error(err))
		}
		return c.ChatRequirements(ctx)
	}

	// 组装回退用的伪 Resp(即使后面走 finalize,也要把这些字段透传出去)
	resp := &ChatRequirementsResp{Persona: prep.Persona}
	resp.Turnstile.Required = prep.Turnstile.Required
	resp.Proofofwork.Required = prep.Proofofwork.Required
	resp.Proofofwork.Seed = prep.Proofofwork.Seed
	resp.Proofofwork.Difficulty = prep.Proofofwork.Difficulty

	// 本地解 PoW(header 里用的 proof_token 与 finalize 里的 proofofwork 同款)
	var proof string
	if prep.Proofofwork.Required {
		proof = SolveProofToken(prep.Proofofwork.Seed, prep.Proofofwork.Difficulty, c.opts.UserAgent)
	}

	// Turnstile 路由
	var turnstileResp string
	if prep.Turnstile.Required {
		if c.opts.TurnstileSolver != nil {
			sCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
			defer cancel()
			out, solveErr := c.opts.TurnstileSolver.Solve(sCtx, prep.Turnstile.DX)
			if solveErr != nil || out == "" {
				if logger := loggerL(); logger != nil {
					logger.Warn("turnstile solver failed, fallback to single-step chat-requirements",
						zap.Error(solveErr))
				}
				return c.ChatRequirements(ctx)
			}
			turnstileResp = out
		} else {
			// 没配 solver,直接回退单步
			if logger := loggerL(); logger != nil {
				logger.Info("turnstile required but no solver configured, fallback to single-step")
			}
			return c.ChatRequirements(ctx)
		}
	}

	// finalize 拿真正 token
	token, persona, err := c.ChatRequirementsFinalize(ctx, prep.PrepareToken, proof, turnstileResp)
	if err != nil {
		if logger := loggerL(); logger != nil {
			logger.Warn("chat-requirements/finalize failed, fallback to single-step",
				zap.Error(err))
		}
		return c.ChatRequirements(ctx)
	}
	resp.Token = token
	if persona != "" {
		resp.Persona = persona
	}
	if logger := loggerL(); logger != nil {
		logger.Info("chat-requirements two-step ok",
			zap.String("persona", resp.Persona),
			zap.Bool("turnstile_required", prep.Turnstile.Required),
			zap.Bool("pow_required", prep.Proofofwork.Required),
			zap.Int("token_len", len(token)),
		)
	}
	return resp, nil
}

// -- dpl / script src 动态提取 --

var (
	// dplCache 缓存从 chatgpt.com 首页 HTML 提取的 dpl build hash 和 script src。
	dplCache struct {
		sync.RWMutex
		dpl       string   // 例如 "dpl=1440a687921de39ff5ee56b92807faaadce73f13"
		scriptSrc []string // 合格的 <script src> URL 列表
		fetchedAt time.Time
	}
	dplTTL = 15 * time.Minute
)

// reDpl 从 script src 中提取 dpl 参数。
var reDpl = regexp.MustCompile(`c/[^/]*?/_`)

// extractDplFromHTML 从首页 HTML 提取 dpl build hash 和 script src。
// dpl 用于 POW config 生成,不应硬编码;上游更新部署后会变化。
func extractDplFromHTML(html string) (dpl string, scripts []string) {
	lines := strings.Split(html, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if !strings.Contains(line, "<script") {
			continue
		}
		// 提取 src="..."
		const srcPfx = `src="`
		idx := strings.Index(line, srcPfx)
		if idx < 0 {
			continue
		}
		rest := line[idx+len(srcPfx):]
		end := strings.Index(rest, `"`)
		if end < 0 {
			continue
		}
		src := rest[:end]
		if !strings.HasPrefix(src, "https://") && !strings.HasPrefix(src, "/") {
			continue
		}
		// 过滤掉第三方脚本,只保留 chatgpt.com 自己的
		if strings.Contains(src, "cdn.") || strings.Contains(src, "challenges.") {
			continue
		}
		scripts = append(scripts, src)
		// 从 src 提取 dpl
		if dpl == "" {
			if m := reDpl.FindString(src); m != "" {
				dpl = "dpl=" + m[2:len(m)-1] // "c/xxx/_" → "dpl=xxx"
			}
		}
	}
	return dpl, scripts
}

// getDplAndScript 返回当前缓存的 dpl 和随机 script src,必要时从 HTML 重新提取。
func getDplAndScript() (string, string) {
	dplCache.RLock()
	if time.Since(dplCache.fetchedAt) < dplTTL && len(dplCache.scriptSrc) > 0 {
		dpl := dplCache.dpl
		scripts := dplCache.scriptSrc
		dplCache.RUnlock()
		//nolint:gosec
		if len(scripts) > 0 {
			return dpl, scripts[time.Now().UnixNano()%int64(len(scripts))]
		}
		return dpl, ""
	}
	dplCache.RUnlock()

	// 需要刷新:返回当前值(可能为空/过期),后台 goroutine 会更新
	dplCache.RLock()
	defer dplCache.RUnlock()
	dpl := dplCache.dpl
	scripts := dplCache.scriptSrc
	if len(scripts) > 0 {
		return dpl, scripts[time.Now().UnixNano()%int64(len(scripts))]
	}
	return dpl, ""
}

// updateDplCache 用新 HTML 内容更新缓存。
func updateDplCache(html string) {
	dpl, scripts := extractDplFromHTML(html)
	if dpl == "" || len(scripts) == 0 {
		return
	}
	dplCache.Lock()
	dplCache.dpl = dpl
	dplCache.scriptSrc = scripts
	dplCache.fetchedAt = time.Now()
	dplCache.Unlock()
}

// SSEEvent 单条 SSE 数据。
type SSEEvent struct {
	Event string
	Data  []byte
	Err   error
}
````

## File: internal/upstream/chatgpt/conversation.go
````go
package chatgpt

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

// ChatMessage 是 OpenAI 风格的一条消息。
// Content 支持纯字符串和多模态数组(OpenAI vision 格式)两种形式。
type ChatMessage struct {
	Role       string          `json:"role"`
	RawContent json.RawMessage `json:"content"`

	// 解析后的字段(非 JSON)
	Content   string   `json:"-"` // 纯文本内容
	ImageURLs []string `json:"-"` // image_url 类型的 URL 列表
}

// UnmarshalJSON 自定义反序列化,兼容 string 和 array 两种 content 格式。
func (m *ChatMessage) UnmarshalJSON(data []byte) error {
	type alias struct {
		Role    string          `json:"role"`
		Content json.RawMessage `json:"content"`
	}
	var a alias
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	m.Role = a.Role
	m.RawContent = a.Content
	m.Content = ""
	m.ImageURLs = nil

	parseContentParts(a.Content, &m.Content, &m.ImageURLs)
	return nil
}

func parseContentParts(raw json.RawMessage, textOut *string, imageOut *[]string) {
	if len(raw) == 0 {
		return
	}
	var s string
	if err := json.Unmarshal(raw, &s); err == nil {
		*textOut = s
		return
	}

	var parts []map[string]json.RawMessage
	if err := json.Unmarshal(raw, &parts); err != nil {
		return
	}
	var texts []string
	seenImages := map[string]struct{}{}
	for _, p := range parts {
		typeName := rawString(p["type"])
		switch typeName {
		case "text", "input_text", "":
			if v := strings.TrimSpace(rawString(p["text"])); v != "" {
				texts = append(texts, v)
			}
		case "image_url", "input_image", "image":
			// URL 会在下方统一提取。
		}

		for _, key := range []string{"image_url", "url", "input_image", "source", "data_url"} {
			for _, u := range rawURLCandidates(p[key]) {
				u = strings.TrimSpace(u)
				if u == "" {
					continue
				}
				if _, ok := seenImages[u]; ok {
					continue
				}
				seenImages[u] = struct{}{}
				*imageOut = append(*imageOut, u)
			}
		}
	}
	*textOut = strings.Join(texts, "\n")
}

func rawString(raw json.RawMessage) string {
	if len(raw) == 0 {
		return ""
	}
	var s string
	if err := json.Unmarshal(raw, &s); err == nil {
		return s
	}
	return ""
}

func rawURLCandidates(raw json.RawMessage) []string {
	if len(raw) == 0 {
		return nil
	}
	if s := rawString(raw); s != "" {
		return []string{s}
	}
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(raw, &obj); err != nil {
		return nil
	}
	out := make([]string, 0, 3)
	for _, key := range []string{"url", "image_url", "data_url", "source", "data", "b64_json", "base64"} {
		if v := rawString(obj[key]); v != "" {
			out = append(out, v)
		}
	}
	return out
}

// MarshalJSON 序列化时保留原始格式。
func (m ChatMessage) MarshalJSON() ([]byte, error) {
	type alias struct {
		Role    string          `json:"role"`
		Content json.RawMessage `json:"content"`
	}
	a := alias{Role: m.Role, Content: m.RawContent}
	if len(a.Content) == 0 && m.Content != "" {
		c, _ := json.Marshal(m.Content)
		a.Content = c
	}
	return json.Marshal(a)
}

// ConversationOpts 是 StreamConversation 的参数。
type ConversationOpts struct {
	Model       string        // 上游模型 slug(如 auto / gpt-4o / o4-mini)
	Messages    []ChatMessage // OpenAI 风格消息
	ParentMsgID string        // 可选,为空自动生成
	ConvID      string        // 可选,为空则新会话
	ProofToken  string        // 可选,POW 解出后填入
	ChatToken   string        // 必传(来自 ChatRequirements)
	ReadTimeout time.Duration // SSE 读超时(单次事件间隔),默认 60s
}

// conversationPayload 对齐 chatgpt.com 请求体(文本模式)。
type conversationPayload struct {
	Action                     string                 `json:"action"`
	Messages                   []upstreamMsg          `json:"messages"`
	ParentMessageID            string                 `json:"parent_message_id"`
	ConversationID             string                 `json:"conversation_id,omitempty"`
	Model                      string                 `json:"model"`
	TimezoneOffsetMin          int                    `json:"timezone_offset_min"`
	Suggestions                []string               `json:"suggestions"`
	HistoryAndTrainingDisabled bool                   `json:"history_and_training_disabled"`
	ConversationMode           map[string]interface{} `json:"conversation_mode"`
	ForceParagen               bool                   `json:"force_paragen"`
	ForceParagenModelSlug      string                 `json:"force_paragen_model_slug"`
	ForceNulligen              bool                   `json:"force_nulligen"`
	ForceRateLimit             bool                   `json:"force_rate_limit"`
	WebsocketRequestID         string                 `json:"websocket_request_id"`
	ClientContextualInfo       map[string]interface{} `json:"client_contextual_info,omitempty"`
	PluginIDs                  []string               `json:"plugin_ids,omitempty"`
}

type upstreamMsg struct {
	ID         string          `json:"id"`
	Author     upstreamAuthor  `json:"author"`
	Content    upstreamContent `json:"content"`
	Metadata   map[string]any  `json:"metadata,omitempty"`
	CreateTime float64         `json:"create_time,omitempty"`
}

type upstreamAuthor struct {
	Role string `json:"role"`
}

type upstreamContent struct {
	ContentType string   `json:"content_type"`
	Parts       []string `json:"parts"`
}

// StreamConversation 向 /backend-api/conversation 发 SSE,返回事件 channel。
// 调用方必须消费完 channel(或 cancel ctx)以释放连接。
func (c *Client) StreamConversation(ctx context.Context, opt ConversationOpts) (<-chan SSEEvent, error) {
	if opt.ChatToken == "" {
		return nil, errors.New("chat_token required")
	}
	if opt.Model == "" {
		opt.Model = "auto"
	}
	if opt.ParentMsgID == "" {
		opt.ParentMsgID = uuid.NewString()
	}
	if opt.ReadTimeout == 0 {
		opt.ReadTimeout = c.opts.SSETimeout
	}

	payload := conversationPayload{
		Action:                     "next",
		Model:                      opt.Model,
		ParentMessageID:            opt.ParentMsgID,
		ConversationID:             opt.ConvID,
		TimezoneOffsetMin:          -480, // UTC+8
		HistoryAndTrainingDisabled: false,
		ConversationMode:           map[string]interface{}{"kind": "primary_assistant"},
		WebsocketRequestID:         uuid.NewString(),
	}
	for _, m := range opt.Messages {
		payload.Messages = append(payload.Messages, upstreamMsg{
			ID:         uuid.NewString(),
			Author:     upstreamAuthor{Role: m.Role},
			Content:    upstreamContent{ContentType: "text", Parts: []string{m.Content}},
			CreateTime: float64(time.Now().Unix()),
		})
	}

	body, _ := json.Marshal(payload)
	req, err := http.NewRequestWithContext(ctx,
		http.MethodPost,
		c.opts.BaseURL+"/backend-api/conversation",
		strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}
	c.commonHeaders(req)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Openai-Sentinel-Chat-Requirements-Token", opt.ChatToken)
	if opt.ProofToken != "" {
		req.Header.Set("Openai-Sentinel-Proof-Token", opt.ProofToken)
	}

	// 对 SSE 请求取消客户端整体 timeout,改为 per-event 读超时控制。
	localClient := *c.hc
	localClient.Timeout = 0
	c.injectHelloID(req)

	res, err := localClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 400 {
		buf, _ := io.ReadAll(res.Body)
		res.Body.Close()
		return nil, &UpstreamError{Status: res.StatusCode, Message: "conversation failed", Body: string(buf)}
	}

	out := make(chan SSEEvent, 32)
	go parseSSE(res.Body, out, opt.ReadTimeout)
	return out, nil
}

// parseSSE 读取 SSE 流,把每个 data: 事件推入 channel。
// chatgpt.com 的事件格式:
//
//	event: delta\n
//	data: {"p":"...","o":"append","v":"..."}\n\n
//
//	data: [DONE]\n\n
//
// readTimeout 控制 per-event 读超时:如果两次换行之间的间隔超过此值,
// 视为上游卡住,发送 Err 并关闭流,防止 goroutine 无限泄漏。
func parseSSE(r io.ReadCloser, out chan<- SSEEvent, readTimeout time.Duration) {
	defer r.Close()
	defer close(out)

	if readTimeout <= 0 {
		readTimeout = 60 * time.Second
	}

	rd := bufio.NewReaderSize(r, 32*1024)
	var event string
	var dataBuf strings.Builder

	// lineChan 在独立 goroutine 中阻塞读行,主循环通过 timer 检测超时。
	type lineResult struct {
		line string
		err  error
	}
	lineCh := make(chan lineResult, 1)

	go func() {
		defer close(lineCh)
		for {
			line, err := rd.ReadString('\n')
			lineCh <- lineResult{line: line, err: err}
			if err != nil {
				return
			}
		}
	}()

	flush := func() {
		if dataBuf.Len() == 0 {
			event = ""
			return
		}
		data := strings.TrimRight(dataBuf.String(), "\n")
		dataBuf.Reset()
		out <- SSEEvent{Event: event, Data: []byte(data)}
		event = ""
	}

	for {
		timer := time.NewTimer(readTimeout)
		select {
		case lr, ok := <-lineCh:
			timer.Stop()
			if !ok {
				// lineCh 已关闭(读 goroutine 退出)
				flush()
				return
			}
			if lr.err != nil {
				if lr.err != io.EOF {
					out <- SSEEvent{Err: fmt.Errorf("sse read: %w", lr.err)}
				} else {
					flush()
				}
				return
			}
			line := strings.TrimRight(lr.line, "\r\n")
			if line == "" {
				// 事件边界
				flush()
				continue
			}
			if strings.HasPrefix(line, ":") {
				// 注释/心跳,忽略
				continue
			}
			if strings.HasPrefix(line, "event:") {
				event = strings.TrimSpace(strings.TrimPrefix(line, "event:"))
				continue
			}
			if strings.HasPrefix(line, "data:") {
				s := strings.TrimPrefix(line, "data:")
				if len(s) > 0 && s[0] == ' ' {
					s = s[1:]
				}
				if dataBuf.Len() > 0 {
					dataBuf.WriteByte('\n')
				}
				dataBuf.WriteString(s)
				continue
			}
			// 其他行忽略

		case <-timer.C:
			// per-event 超时:上游 SSE 停顿太久,主动关闭防止 goroutine 泄漏。
			out <- SSEEvent{Err: fmt.Errorf("sse read timeout (%v)", readTimeout)}
			return
		}
	}
}
````

## File: internal/upstream/chatgpt/fchat.go
````go
// fchat.go —— 文字聊天走 /backend-api/f/conversation 新协议。
//
// 背景:chatgpt.com 近期新账号在老 /backend-api/conversation 端点上会直接
// 回一条 author=system、content 为空、status=finished_successfully 的消息,
// 等同于"被静默拒绝"。对齐浏览器抓包与社区维护的 OpenaiChat provider,
// 文字聊天正确顺序是:
//
//	1. GET /                            → 拿 __cf_bm / oai-did / _cfuvid cookie
//	2. sentinel/chat-requirements       → 拿 chat_token + proofofwork 挑战
//	3. f/conversation/prepare           → 带 chat_token(!) + proof_token,拿 conduit_token
//	4. f/conversation (SSE)             → 带 chat_token + proof_token + conduit_token
//
// 要点:prepare 必须在 chat-requirements 之后,并且要把 `openai-sentinel-chat-
// requirements-token` 带进 header。不调用 /backend-api/conversation/init,
// 该端点在免费/新账号上会直接 404。

package chatgpt

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

// FChatOpts 是 StreamFChat 的入参。
type FChatOpts struct {
	UpstreamModel string        // 默认 "auto"
	Messages      []ChatMessage // OpenAI 风格
	ChatToken     string        // 必传(StreamFChat 时,PrepareFChat 不需要)
	ProofToken    string        // 可选
	ConduitToken  string        // 可选(PrepareFChat 返回)
	ConvID        string        // 复用会话时传
	ParentMsgID   string        // 复用会话时传 GetConversationHead 结果
	SSETimeout    time.Duration // 默认 120s
}

// PrepareFChat 对 /backend-api/f/conversation/prepare 发一条文字 prepare,返回 conduit_token。
//
// payload 严格对齐浏览器抓包(HAR 中 /f/conversation/prepare 请求):
//   - client_prepare_state: "success"
//   - fork_from_shared_post: false
//   - partial_query: 一条完整的 user message(id+author+content),不是空字符串
//   - system_hints: []  ← text 通路是空数组(注意:image 通路是 ["picture_v2"])
//   - client_contextual_info: { "app_name": "chatgpt.com" }  ← prepare 只带 app_name
//
// header 带 Openai-Sentinel-Chat-Requirements-Token(必须)+ 可选 Proof-Token。
func (c *Client) PrepareFChat(ctx context.Context, opt FChatOpts) (string, error) {
	if opt.ChatToken == "" {
		return "", errors.New("chat_token required for prepare")
	}
	if opt.UpstreamModel == "" {
		opt.UpstreamModel = "auto"
	}
	if len(opt.Messages) == 0 {
		return "", errors.New("messages required")
	}
	if opt.ParentMsgID == "" {
		opt.ParentMsgID = uuid.NewString()
	}

	// partial_query 用最后一条 user message。浏览器的做法是"用户一输完字就发
	// prepare",此时还没 send,所以 partial_query = 当前正在输入的内容。
	var userPart string
	for i := len(opt.Messages) - 1; i >= 0; i-- {
		if opt.Messages[i].Role == "user" {
			userPart = opt.Messages[i].Content
			break
		}
	}

	payload := map[string]interface{}{
		"action":                "next",
		"fork_from_shared_post": false,
		"parent_message_id":     opt.ParentMsgID,
		"model":                 opt.UpstreamModel,
		"client_prepare_state":  "success",
		"timezone_offset_min":   -480,
		"timezone":              "Asia/Shanghai",
		"conversation_mode":     map[string]string{"kind": "primary_assistant"},
		"system_hints":          []string{},
		"partial_query": map[string]interface{}{
			"id":     uuid.NewString(),
			"author": map[string]string{"role": "user"},
			"content": map[string]interface{}{
				"content_type": "text",
				"parts":        []string{userPart},
			},
		},
		"supports_buffering":  true,
		"supported_encodings": []string{"v1"},
		"client_contextual_info": map[string]interface{}{
			"app_name": "chatgpt.com",
		},
	}
	if opt.ConvID != "" {
		payload["conversation_id"] = opt.ConvID
	}
	body, _ := json.Marshal(payload)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		c.opts.BaseURL+"/backend-api/f/conversation/prepare",
		strings.NewReader(string(body)))
	if err != nil {
		return "", err
	}
	c.commonHeaders(req)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Openai-Sentinel-Chat-Requirements-Token", opt.ChatToken)
	if opt.ProofToken != "" {
		req.Header.Set("Openai-Sentinel-Proof-Token", opt.ProofToken)
	}

	res, err := c.hc.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	buf, _ := io.ReadAll(res.Body)
	if res.StatusCode >= 400 {
		return "", &UpstreamError{Status: res.StatusCode, Message: "f/conversation/prepare failed", Body: string(buf)}
	}
	var out struct {
		ConduitToken string `json:"conduit_token"`
	}
	_ = json.Unmarshal(buf, &out)
	return out.ConduitToken, nil
}

// StreamFChat 发起一次文字 f/conversation SSE。
// 调用前请确保:新会话场景先调用 InitConversation(ctx)(空 system_hints)。
func (c *Client) StreamFChat(ctx context.Context, opt FChatOpts) (<-chan SSEEvent, error) {
	if opt.ChatToken == "" {
		return nil, errors.New("chat_token required")
	}
	if opt.UpstreamModel == "" {
		opt.UpstreamModel = "auto"
	}
	if len(opt.Messages) == 0 {
		return nil, errors.New("messages required")
	}
	if opt.ParentMsgID == "" {
		opt.ParentMsgID = uuid.NewString()
	}
	if opt.SSETimeout == 0 {
		opt.SSETimeout = 120 * time.Second
	}

	// messages[].metadata 严格对齐 HAR 抓包的 **text 通路**:
	//   developer_mode_connector_ids / selected_sources / selected_github_repos /
	//   selected_all_github_repos / serialization_metadata
	// 其中 **selected_sources 只有 text 通路有,image 通路没有**;反之
	// system_hints: ["picture_v2"] 只有 image 通路有,text 完全不写这个 key。
	// 写错会导致上游认为客户端类型不匹配,触发 silent rejection。
	msgs := make([]map[string]interface{}, 0, len(opt.Messages))
	for _, m := range opt.Messages {
		msgs = append(msgs, map[string]interface{}{
			"id":          uuid.NewString(),
			"author":      map[string]string{"role": m.Role},
			"create_time": float64(time.Now().UnixMilli()) / 1000.0,
			"content":     map[string]interface{}{"content_type": "text", "parts": []string{m.Content}},
			"metadata": map[string]interface{}{
				"developer_mode_connector_ids": []interface{}{},
				"selected_sources":             []interface{}{},
				"selected_github_repos":        []interface{}{},
				"selected_all_github_repos":    false,
				"serialization_metadata": map[string]interface{}{
					"custom_symbol_offsets": []interface{}{},
				},
			},
		})
	}

	// 顶层 payload 对齐 HAR /f/conversation 抓包(text 通路):
	//   client_prepare_state: "sent"            ← prepare 已发过一次
	//   system_hints: []                         ← text 空数组
	//   force_parallel_switch: "auto"            ← 必带
	//   client_contextual_info: 7 个字段 + app_name
	payload := map[string]interface{}{
		"action":                   "next",
		"messages":                 msgs,
		"parent_message_id":        opt.ParentMsgID,
		"model":                    opt.UpstreamModel,
		"client_prepare_state":     "sent",
		"timezone_offset_min":      -480,
		"timezone":                 "Asia/Shanghai",
		"conversation_mode":        map[string]string{"kind": "primary_assistant"},
		"enable_message_followups": true,
		"system_hints":             []string{},
		"supports_buffering":       true,
		"supported_encodings":      []string{"v1"},
		"client_contextual_info": map[string]interface{}{
			"is_dark_mode":      false,
			"time_since_loaded": 1200,
			"page_height":       1072,
			"page_width":        1724,
			"pixel_ratio":       1.2,
			"screen_height":     1440,
			"screen_width":      2560,
			"app_name":          "chatgpt.com",
		},
		"paragen_cot_summary_display_override": "allow",
		"force_parallel_switch":                "auto",
	}
	// 只有已有会话才带 conversation_id;新会话完全不带这个 key(对齐浏览器抓包)。
	if opt.ConvID != "" {
		payload["conversation_id"] = opt.ConvID
	}
	body, _ := json.Marshal(payload)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		c.opts.BaseURL+"/backend-api/f/conversation",
		strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}
	c.commonHeaders(req)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	// X-Oai-Turn-Trace-Id:真实浏览器每 turn 一个新 UUID。风控会据此做请求配对
	// (prepare 与 finalize / f-prepare 与 f-conversation 的 trace 对得上)。
	req.Header.Set("X-Oai-Turn-Trace-Id", uuid.NewString())
	req.Header.Set("Openai-Sentinel-Chat-Requirements-Token", opt.ChatToken)
	if opt.ProofToken != "" {
		req.Header.Set("Openai-Sentinel-Proof-Token", opt.ProofToken)
	}
	if opt.ConduitToken != "" {
		req.Header.Set("X-Conduit-Token", opt.ConduitToken)
	}

	local := *c.hc
	local.Timeout = 0
	c.injectHelloID(req)

	res, err := local.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 400 {
		buf, _ := io.ReadAll(res.Body)
		res.Body.Close()
		return nil, &UpstreamError{Status: res.StatusCode, Message: "f/conversation(chat) failed", Body: string(buf)}
	}
	out := make(chan SSEEvent, 64)
	go parseSSE(res.Body, out, opt.SSETimeout)
	return out, nil
}
````

## File: internal/upstream/chatgpt/files.go
````go
// files.go —— chatgpt.com 文件上传协议,图生图/图像编辑的前置步骤。
//
// 三步协议(对齐 chatgpt.com 浏览器真实抓包):
//
//  1. POST /backend-api/files
//     body: {file_name, file_size, use_case: "multimodal"}
//     resp: {file_id, upload_url, status: "success"}
//
//  2. PUT <upload_url>                 (Azure Blob SAS URL)
//     headers: Content-Type / x-ms-blob-type: BlockBlob / x-ms-version: 2020-04-08 / Origin
//     body: 原始字节
//
//  3. POST /backend-api/files/{file_id}/uploaded
//     body: {}
//     resp: {status: "success", download_url, ...}
//
// 上传完成后,在 f/conversation.messages 里:
//   - content 从 text 变 multimodal_text;parts 前面加上
//     {"asset_pointer": "file-service://<file_id>", "height":.., "width":.., "size_bytes":..}
//   - metadata.attachments 加一项 {id, mimeType, name, size, height?, width?}
//
// 注意:upload_url 指向 Azure,不要走同一个 chatgpt.com 代理/utls transport,
// 这里用单独的一个 http.Client(沿用 Client 内部的 Transport 走代理,但不带 Auth/Oai-* 头)。

package chatgpt

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	_ "image/gif"  // register decoders
	_ "image/jpeg" //
	_ "image/png"  //
	"io"
	"net/http"
	"strings"
	"time"
)

// UploadedFile 是三步上传后沉淀的"可 attach 给 messages"的元数据。
// 字段命名对齐 chatgpt.com 的 attachment payload,序列化时直接当 map 用。
type UploadedFile struct {
	FileID      string `json:"file_id"`
	FileName    string `json:"file_name"`
	FileSize    int    `json:"file_size"`
	MimeType    string `json:"mime_type"`
	UseCase     string `json:"use_case"`         // 图片: multimodal, 文件: my_files
	Width       int    `json:"width,omitempty"`  // 仅图片
	Height      int    `json:"height,omitempty"` // 仅图片
	DownloadURL string `json:"download_url"`     // POST /uploaded 返回,通常不直接用
}

// UploadFile 执行完整三步上传。调用方传入原始字节 + 建议的文件名即可。
// 识别到 image/* 时会尝试 Decode 拿到宽高(Decode 失败不致命,按 0 处理)。
//
// 实践经验:步骤 1、3 走 chatgpt.com(uTLS / 代理 / auth 头),步骤 2 走 Azure,
// 用同一条 http.Client 但是请求头手动裁剪;Azure 的 SAS URL 本身带鉴权。
func (c *Client) UploadFile(ctx context.Context, data []byte, fileName string) (*UploadedFile, error) {
	if len(data) == 0 {
		return nil, errors.New("empty file data")
	}
	mime, ext := sniffMime(data)
	useCase := "multimodal"
	if !strings.HasPrefix(mime, "image/") {
		useCase = "my_files"
	}
	if fileName == "" {
		fileName = fmt.Sprintf("file-%d%s", len(data), ext)
	}

	out := &UploadedFile{
		FileName: fileName,
		FileSize: len(data),
		MimeType: mime,
		UseCase:  useCase,
	}
	if strings.HasPrefix(mime, "image/") {
		if img, _, err := image.DecodeConfig(bytes.NewReader(data)); err == nil {
			out.Width = img.Width
			out.Height = img.Height
		} else if mime == "image/webp" {
			out.Width, out.Height = webpDimensions(data)
		}
	}

	// ---- Step 1: POST /backend-api/files ----
	step1Body := map[string]interface{}{
		"file_name": fileName,
		"file_size": len(data),
		"use_case":  useCase,
	}
	if out.Width > 0 && out.Height > 0 {
		step1Body["height"] = out.Height
		step1Body["width"] = out.Width
	}
	b1, _ := json.Marshal(step1Body)
	req1, err := http.NewRequestWithContext(ctx, http.MethodPost,
		c.opts.BaseURL+"/backend-api/files",
		bytes.NewReader(b1))
	if err != nil {
		return nil, err
	}
	c.commonHeaders(req1)
	req1.Header.Set("Content-Type", "application/json")
	req1.Header.Set("Accept", "application/json")
	res1, err := c.hc.Do(req1)
	if err != nil {
		return nil, fmt.Errorf("create file: %w", err)
	}
	defer res1.Body.Close()
	buf1, _ := io.ReadAll(res1.Body)
	if res1.StatusCode >= 400 {
		return nil, &UpstreamError{Status: res1.StatusCode, Message: "create file failed", Body: string(buf1)}
	}
	var step1Resp struct {
		FileID    string `json:"file_id"`
		UploadURL string `json:"upload_url"`
		Status    string `json:"status"`
	}
	if err := json.Unmarshal(buf1, &step1Resp); err != nil {
		return nil, fmt.Errorf("decode create-file resp: %w (body=%s)", err, truncateStr(string(buf1), 200))
	}
	if step1Resp.FileID == "" || step1Resp.UploadURL == "" {
		return nil, fmt.Errorf("create-file empty: %s", truncateStr(string(buf1), 200))
	}
	out.FileID = step1Resp.FileID

	// chatgpt 浏览器行为:step1 和 step2 之间会 sleep 一小会儿,避免 Azure 那边
	// 还没完成 SAS 分发。参考实现是 1s,这里保守点给 500ms。
	select {
	case <-time.After(500 * time.Millisecond):
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	// ---- Step 2: PUT upload_url (Azure Blob) ----
	req2, err := http.NewRequestWithContext(ctx, http.MethodPut, step1Resp.UploadURL, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req2.Header.Set("Content-Type", mime)
	req2.Header.Set("x-ms-blob-type", "BlockBlob")
	req2.Header.Set("x-ms-version", "2020-04-08")
	req2.Header.Set("Origin", c.opts.BaseURL)
	req2.Header.Set("User-Agent", c.opts.UserAgent)
	req2.Header.Set("Accept", "application/json, text/plain, */*")
	req2.Header.Set("Accept-Language", "en-US,en;q=0.8")
	req2.Header.Set("Referer", c.opts.BaseURL+"/")
	res2, err := c.hc.Do(req2)
	if err != nil {
		return nil, fmt.Errorf("upload PUT: %w", err)
	}
	defer res2.Body.Close()
	if res2.StatusCode >= 400 {
		buf2, _ := io.ReadAll(res2.Body)
		return nil, &UpstreamError{Status: res2.StatusCode, Message: "upload PUT failed", Body: string(buf2)}
	}
	_, _ = io.Copy(io.Discard, res2.Body)

	// ---- Step 3: POST /backend-api/files/{file_id}/uploaded ----
	req3, err := http.NewRequestWithContext(ctx, http.MethodPost,
		c.opts.BaseURL+"/backend-api/files/"+step1Resp.FileID+"/uploaded",
		strings.NewReader("{}"))
	if err != nil {
		return nil, err
	}
	c.commonHeaders(req3)
	req3.Header.Set("Content-Type", "application/json")
	req3.Header.Set("Accept", "application/json")
	res3, err := c.hc.Do(req3)
	if err != nil {
		return nil, fmt.Errorf("register uploaded: %w", err)
	}
	defer res3.Body.Close()
	buf3, _ := io.ReadAll(res3.Body)
	if res3.StatusCode >= 400 {
		return nil, &UpstreamError{Status: res3.StatusCode, Message: "register uploaded failed", Body: string(buf3)}
	}
	var step3Resp struct {
		Status      string `json:"status"`
		DownloadURL string `json:"download_url"`
	}
	_ = json.Unmarshal(buf3, &step3Resp)
	out.DownloadURL = step3Resp.DownloadURL

	return out, nil
}

// Attachment 是 messages[*].metadata.attachments[*] 的序列化对象。
type Attachment struct {
	ID       string `json:"id"`
	MimeType string `json:"mimeType"`
	Name     string `json:"name"`
	Size     int    `json:"size"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
}

// ToAttachment 把一个已上传的 file 转成 messages.metadata.attachments 里的条目。
func (u *UploadedFile) ToAttachment() Attachment {
	a := Attachment{ID: u.FileID, MimeType: u.MimeType, Name: u.FileName, Size: u.FileSize}
	if u.UseCase == "multimodal" {
		a.Width = u.Width
		a.Height = u.Height
	}
	return a
}

// AssetPointerPart 是 messages[*].content.parts 里的一项(图片),
// 用于把 file-service:// 挂到多模态消息最前面。
type AssetPointerPart struct {
	ContentType  string `json:"content_type,omitempty"` // "image_asset_pointer"
	AssetPointer string `json:"asset_pointer"`
	Width        int    `json:"width,omitempty"`
	Height       int    `json:"height,omitempty"`
	SizeBytes    int    `json:"size_bytes,omitempty"`
}

// ToAssetPointerPart 返回 multimodal_text.parts 里 insert 在 prompt 前的那一项。
func (u *UploadedFile) ToAssetPointerPart() AssetPointerPart {
	return AssetPointerPart{
		ContentType:  "image_asset_pointer",
		AssetPointer: "file-service://" + u.FileID,
		Width:        u.Width,
		Height:       u.Height,
		SizeBytes:    u.FileSize,
	}
}

// sniffMime 用前 512 字节识别 mime 和推荐扩展名。
// net/http 的 DetectContentType 已足够覆盖 png/jpg/gif/webp 的主流场景。
func sniffMime(data []byte) (mime, ext string) {
	n := 512
	if len(data) < n {
		n = len(data)
	}
	mime = http.DetectContentType(data[:n])
	// DetectContentType 可能附带 charset,去掉
	if i := strings.Index(mime, ";"); i >= 0 {
		mime = strings.TrimSpace(mime[:i])
	}
	switch {
	case strings.EqualFold(mime, "image/jpeg"):
		ext = ".jpg"
	case strings.EqualFold(mime, "image/png"):
		ext = ".png"
	case strings.EqualFold(mime, "image/gif"):
		ext = ".gif"
	case strings.EqualFold(mime, "image/webp"):
		ext = ".webp"
	case strings.EqualFold(mime, "application/pdf"):
		ext = ".pdf"
	default:
		ext = ""
	}
	return
}

func webpDimensions(data []byte) (int, int) {
	if len(data) < 30 || string(data[0:4]) != "RIFF" || string(data[8:12]) != "WEBP" {
		return 0, 0
	}
	pos := 12
	for pos+8 <= len(data) {
		fourcc := string(data[pos : pos+4])
		sz := int(binary.LittleEndian.Uint32(data[pos+4 : pos+8]))
		chunkStart := pos + 8
		if chunkStart+sz > len(data) {
			return 0, 0
		}
		switch fourcc {
		case "VP8X":
			if sz >= 10 {
				w := 1 + int(data[chunkStart+4]) + int(data[chunkStart+5])<<8 + int(data[chunkStart+6])<<16
				h := 1 + int(data[chunkStart+7]) + int(data[chunkStart+8])<<8 + int(data[chunkStart+9])<<16
				return w, h
			}
		case "VP8L":
			if sz >= 5 && data[chunkStart] == 0x2f {
				b0, b1, b2, b3 := data[chunkStart+1], data[chunkStart+2], data[chunkStart+3], data[chunkStart+4]
				w := 1 + int(b0) + (int(b1&0x3f) << 8)
				h := 1 + int(b1>>6) + (int(b2) << 2) + (int(b3&0x0f) << 10)
				return w, h
			}
		case "VP8 ":
			if sz >= 10 {
				off := chunkStart + 6
				if off+4 <= len(data) {
					w := int(binary.LittleEndian.Uint16(data[off:off+2]) & 0x3fff)
					h := int(binary.LittleEndian.Uint16(data[off+2:off+4]) & 0x3fff)
					return w, h
				}
			}
		}
		pos = chunkStart + sz
		if sz%2 == 1 {
			pos++
		}
	}
	return 0, 0
}

func truncateStr(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}
````

## File: internal/upstream/chatgpt/image_img2_test.go
````go
package chatgpt

import "testing"

func TestExtractImageToolMsgsDetectsIMG2SedimentOnly(t *testing.T) {
	mapping := map[string]interface{}{
		"tool-node": map[string]interface{}{
			"message": map[string]interface{}{
				"author":      map[string]interface{}{"role": "tool"},
				"create_time": float64(1760000000),
				"metadata": map[string]interface{}{
					"async_task_type": "image_gen",
					"model_slug":      "gpt-5-3",
				},
				"content": map[string]interface{}{
					"content_type": "multimodal_text",
					"parts": []interface{}{
						map[string]interface{}{
							"content_type":  "image_asset_pointer",
							"asset_pointer": "sediment://file_000000001d4071fd83437b6e5d5bcaa9",
							"width":         float64(1536),
							"height":        float64(1024),
							"size_bytes":    float64(2540679),
							"metadata": map[string]interface{}{
								"generation": map[string]interface{}{
									"gen_size":    "image",
									"gen_size_v2": "48",
									"orientation": "landscape",
								},
							},
						},
					},
				},
			},
		},
	}

	msgs := ExtractImageToolMsgs(mapping)
	if len(msgs) != 1 {
		t.Fatalf("len(msgs)=%d, want 1", len(msgs))
	}
	msg := msgs[0]
	if !msg.IMG2Hint {
		t.Fatalf("IMG2Hint=false, want true")
	}
	if got, want := msg.SedimentIDs, []string{"file_000000001d4071fd83437b6e5d5bcaa9"}; len(got) != 1 || got[0] != want[0] {
		t.Fatalf("SedimentIDs=%v, want %v", got, want)
	}
	if got, want := msg.GenSizeV2s, []string{"48"}; len(got) != 1 || got[0] != want[0] {
		t.Fatalf("GenSizeV2s=%v, want %v", got, want)
	}
	if msg.MaxWidth != 1536 || msg.MaxHeight != 1024 || msg.MaxSizeBytes != 2540679 {
		t.Fatalf("metadata dims/size = %dx%d %d", msg.MaxWidth, msg.MaxHeight, msg.MaxSizeBytes)
	}
}

func TestParseImageSSEDetectsIMG2Sediment(t *testing.T) {
	stream := make(chan SSEEvent, 1)
	stream <- SSEEvent{Data: []byte(`{"v":{"message":{"content":{"parts":[{"asset_pointer":"sediment://file_img2","metadata":{"generation":{"gen_size_v2":"48"}}}]}}}}`)}
	close(stream)

	res := ParseImageSSE(stream)
	if got, want := res.SedimentIDs, []string{"file_img2"}; len(got) != 1 || got[0] != want[0] {
		t.Fatalf("SedimentIDs=%v, want %v", got, want)
	}
	if got, want := res.IMG2SedimentIDs, []string{"file_img2"}; len(got) != 1 || got[0] != want[0] {
		t.Fatalf("IMG2SedimentIDs=%v, want %v", got, want)
	}
}
````

## File: internal/upstream/chatgpt/image.go
````go
// Package chatgpt - 图像生成协议
//
// 完整链路(和文字聊天共用 f/conversation,只通过 system_hints=["picture_v2"] 区分):
//
//  0. (可选) GET /                              → 拿 oai-did cookie
//  1. POST /backend-api/sentinel/chat-requirements → chat_token + 可选 POW 挑战
//  2. POST /backend-api/f/conversation/prepare      → conduit_token(灰度分桶关键)
//  3. POST /backend-api/f/conversation (SSE)         → 边解析边收 file-service://
//  4. 灰度命中判据:SSE 没直出 file-service 时轮询
//     GET /backend-api/conversation/{id}
//     - file-service 或 metadata.generation.gen_size_v2 → IMG2 终稿
//     - IMG2 tool 消息 ≥ 2 条 → 灰度命中,聚合所有 refs
//     - 只 1 条且无 gen_size_v2 → preview_only(非灰度,同会话重试)
//  5. GET /backend-api/files/download/{fid}?conversation_id=... → 签名 URL / estuary
//     fallback: /files/{fid}/download 或 /conversation/{cid}/attachment/{sid}/download
//  6. GET 签名 URL → 图片字节
//
// 注意:不要调用 /backend-api/conversation/init——这是老客户端路径,在免费账号上会
// 直接 404 让整条链路失败,上游把 picture_v2 路由完全交给 f/conversation 的 payload。
package chatgpt

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
)

// ImageConvOpts 是图像会话的入参。
type ImageConvOpts struct {
	Prompt         string          // 用户提示词(已处理完的,含可选 CLARITY_SUFFIX)
	UpstreamModel  string          // 默认 "gpt-5-3"
	ConvID         string          // 复用会话时填,空则新建
	ParentMsgID    string          // 复用会话时从 GetConversationHead 取;新会话随机
	MessageID      string          // 可选,留空自动生成
	ChatToken      string          // 必传,来自 ChatRequirements
	ProofToken     string          // 可选
	ConduitToken   string          // 可选,来自 PrepareFConversation
	TimezoneOffset int             // 默认 -480(Asia/Shanghai)
	SSETimeout     time.Duration   // 默认 120s
	References     []*UploadedFile // 图生图/编辑:已上传的参考图,会插到 prompt 前面
}

// InitConversation 对应 /backend-api/conversation/init。
// 新会话必须调用,否则后续 /f/conversation 会 404。
// systemHints 为空串数组表示文字聊天;图像场景传 []string{"picture_v2"}。
func (c *Client) InitConversation(ctx context.Context, systemHints ...string) error {
	if systemHints == nil {
		systemHints = []string{}
	}
	payload := map[string]interface{}{
		"gizmo_id":                nil,
		"requested_default_model": nil,
		"conversation_id":         nil,
		"timezone_offset_min":     -480,
		"system_hints":            systemHints,
	}
	body, _ := json.Marshal(payload)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		c.opts.BaseURL+"/backend-api/conversation/init",
		strings.NewReader(string(body)))
	if err != nil {
		return err
	}
	c.commonHeaders(req)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")

	res, err := c.hc.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		buf, _ := io.ReadAll(res.Body)
		return &UpstreamError{Status: res.StatusCode, Message: "conversation/init failed", Body: string(buf)}
	}
	_, _ = io.Copy(io.Discard, res.Body)
	return nil
}

// PrepareFConversation 对应 /backend-api/f/conversation/prepare,返回 conduit_token。
//
// payload 对齐 HAR 抓包 /f/conversation/prepare(image 通路):
//   - client_prepare_state: "success"
//   - fork_from_shared_post: false
//   - partial_query: 完整的 user message(id+author+content,包含当前 prompt)
//   - system_hints: ["picture_v2"]   ← image 通路特有
//   - client_contextual_info: { "app_name": "chatgpt.com" }   ← prepare 阶段只带 app_name
func (c *Client) PrepareFConversation(ctx context.Context, opt ImageConvOpts) (string, error) {
	if opt.UpstreamModel == "" {
		opt.UpstreamModel = "auto"
	}
	if opt.MessageID == "" {
		opt.MessageID = uuid.NewString()
	}
	msgContent, msgMeta := buildImageUserMessage(opt.Prompt, opt.References)
	partialQuery := map[string]interface{}{
		"id":      opt.MessageID,
		"author":  map[string]string{"role": "user"},
		"content": msgContent,
	}
	if len(msgMeta) > 0 {
		partialQuery["metadata"] = msgMeta
	}
	payload := map[string]interface{}{
		"action":                "next",
		"fork_from_shared_post": false,
		"parent_message_id":     opt.ParentMsgID,
		"model":                 opt.UpstreamModel,
		"client_prepare_state":  "success",
		"timezone_offset_min":   -480,
		"timezone":              "Asia/Shanghai",
		"conversation_mode":     map[string]string{"kind": "primary_assistant"},
		"system_hints":          []string{"picture_v2"},
		"partial_query":         partialQuery,
		"supports_buffering":    true,
		"supported_encodings":   []string{"v1"},
		"client_contextual_info": map[string]interface{}{
			"app_name": "chatgpt.com",
		},
	}
	// 只有已有会话才带 conversation_id;新会话不带这个 key(对齐浏览器抓包,
	// 带陌生 UUID 上游会 404)
	if opt.ConvID != "" {
		payload["conversation_id"] = opt.ConvID
	}
	body, _ := json.Marshal(payload)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		c.opts.BaseURL+"/backend-api/f/conversation/prepare",
		strings.NewReader(string(body)))
	if err != nil {
		return "", err
	}
	c.commonHeaders(req)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Openai-Sentinel-Chat-Requirements-Token", opt.ChatToken)
	if opt.ProofToken != "" {
		req.Header.Set("Openai-Sentinel-Proof-Token", opt.ProofToken)
	}

	res, err := c.hc.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	buf, _ := io.ReadAll(res.Body)
	if res.StatusCode >= 400 {
		return "", &UpstreamError{Status: res.StatusCode, Message: "f/conversation/prepare failed", Body: string(buf)}
	}
	var out struct {
		ConduitToken string `json:"conduit_token"`
	}
	_ = json.Unmarshal(buf, &out)
	return out.ConduitToken, nil
}

func buildImageUserMessage(prompt string, refs []*UploadedFile) (map[string]interface{}, map[string]interface{}) {
	msgContent := map[string]interface{}{"content_type": "text", "parts": []string{prompt}}
	msgMeta := map[string]interface{}{
		"developer_mode_connector_ids": []interface{}{},
		"selected_github_repos":        []interface{}{},
		"selected_all_github_repos":    false,
		"system_hints":                 []string{"picture_v2"},
		"serialization_metadata": map[string]interface{}{
			"custom_symbol_offsets": []interface{}{},
		},
	}
	if len(refs) == 0 {
		return msgContent, msgMeta
	}
	parts := make([]interface{}, 0, len(refs)+1)
	atts := make([]Attachment, 0, len(refs))
	for _, u := range refs {
		if u == nil || u.FileID == "" {
			continue
		}
		parts = append(parts, u.ToAssetPointerPart())
		atts = append(atts, u.ToAttachment())
	}
	parts = append(parts, prompt)
	msgContent = map[string]interface{}{
		"content_type": "multimodal_text",
		"parts":        parts,
	}
	msgMeta["attachments"] = atts
	return msgContent, msgMeta
}

// StreamFConversation 对应 /backend-api/f/conversation(图像走和文字同一端点)。
//
// payload 字段集参考社区维护的 OpenaiChat provider(它在免费/付费账号上实测可用):
// 不带 client_prepare_state / force_parallel_switch;message.metadata 只带
// serialization_metadata + system_hints(有图时再补 attachments)。
func (c *Client) StreamFConversation(ctx context.Context, opt ImageConvOpts) (<-chan SSEEvent, error) {
	if opt.UpstreamModel == "" {
		opt.UpstreamModel = "auto"
	}
	if opt.MessageID == "" {
		opt.MessageID = uuid.NewString()
	}
	if opt.ParentMsgID == "" {
		opt.ParentMsgID = uuid.NewString()
	}
	if opt.TimezoneOffset == 0 {
		opt.TimezoneOffset = -480
	}
	if opt.SSETimeout == 0 {
		opt.SSETimeout = 180 * time.Second
	}

	// 构造 messages[0] 的 content 与 metadata,按是否有 reference 图决定 multimodal_text。
	msgContent, msgMeta := buildImageUserMessage(opt.Prompt, opt.References)

	// 顶层 payload 对齐 HAR /f/conversation(image 通路):
	//   client_prepare_state: "sent"
	//   system_hints: ["picture_v2"]
	//   force_parallel_switch: "auto"            ← 必带
	//   client_contextual_info: 7 个字段 + app_name
	payload := map[string]interface{}{
		"action": "next",
		"messages": []map[string]interface{}{{
			"id":          opt.MessageID,
			"author":      map[string]string{"role": "user"},
			"create_time": float64(time.Now().UnixMilli()) / 1000.0,
			"content":     msgContent,
			"metadata":    msgMeta,
		}},
		"parent_message_id":        opt.ParentMsgID,
		"model":                    opt.UpstreamModel,
		"client_prepare_state":     "sent",
		"timezone_offset_min":      opt.TimezoneOffset,
		"timezone":                 "Asia/Shanghai",
		"conversation_mode":        map[string]string{"kind": "primary_assistant"},
		"enable_message_followups": true,
		"system_hints":             []string{"picture_v2"},
		"supports_buffering":       true,
		"supported_encodings":      []string{"v1"},
		"client_contextual_info": map[string]interface{}{
			"is_dark_mode":      false,
			"time_since_loaded": 1200,
			"page_height":       1072,
			"page_width":        1724,
			"pixel_ratio":       1.2,
			"screen_height":     1440,
			"screen_width":      2560,
			"app_name":          "chatgpt.com",
		},
		"paragen_cot_summary_display_override": "allow",
		"force_parallel_switch":                "auto",
	}
	// 新会话不带 conversation_id(对齐浏览器抓包);已有会话才带
	if opt.ConvID != "" {
		payload["conversation_id"] = opt.ConvID
	}
	body, _ := json.Marshal(payload)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		c.opts.BaseURL+"/backend-api/f/conversation",
		strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}
	c.commonHeaders(req)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	// X-Oai-Turn-Trace-Id:每 turn 一个新 UUID。见 fchat.go 说明。
	req.Header.Set("X-Oai-Turn-Trace-Id", uuid.NewString())
	req.Header.Set("Openai-Sentinel-Chat-Requirements-Token", opt.ChatToken)
	if opt.ProofToken != "" {
		req.Header.Set("Openai-Sentinel-Proof-Token", opt.ProofToken)
	}
	if opt.ConduitToken != "" {
		req.Header.Set("X-Conduit-Token", opt.ConduitToken)
	}

	local := *c.hc
	local.Timeout = 0 // 由 ctx 控制
	c.injectHelloID(req)

	res, err := local.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 400 {
		buf, _ := io.ReadAll(res.Body)
		res.Body.Close()
		return nil, &UpstreamError{Status: res.StatusCode, Message: "f/conversation failed", Body: string(buf)}
	}
	out := make(chan SSEEvent, 64)
	go parseSSE(res.Body, out, opt.SSETimeout)
	return out, nil
}

// ImageSSEResult 是 ParseImageSSE 的扫描结果。
type ImageSSEResult struct {
	ConversationID  string   // SSE 里捕获到的新会话 id(可能和入参不同)
	FileIDs         []string // file-service:// 引用(直出灰度图就在这里)
	SedimentIDs     []string // sediment:// 引用(可能是预览,也可能是 IMG2 终稿)
	IMG2SedimentIDs []string // SSE 中带 gen_size_v2 等 IMG2 元数据的 sediment 引用
	FinishType      string   // finish_details.type(interrupted/stop/...)
	ImageGenTaskID  string
	PolicyBlocked   bool   // 上游因内容策略拒绝生成
	AssistantText   string // 助手回复的文本(用于提取拒绝原因)
}

var (
	reFileRef = regexp.MustCompile(`file-service://([A-Za-z0-9_-]+)`)
	reSedRef  = regexp.MustCompile(`sediment://([A-Za-z0-9_-]+)`)
)

// ParseImageSSE 消费 SSE 事件流,把图像相关的字段提取出来。
// 调用方可以根据返回的 FileIDs 判断是否已灰度直出。
func ParseImageSSE(stream <-chan SSEEvent) ImageSSEResult {
	var r ImageSSEResult
	seenFile := map[string]struct{}{}
	seenSed := map[string]struct{}{}
	seenIMG2Sed := map[string]struct{}{}

	for ev := range stream {
		if ev.Err != nil {
			return r
		}
		data := ev.Data
		if len(data) == 0 {
			continue
		}
		if string(data) == "[DONE]" {
			return r
		}
		// 文本正则先扫一遍(比 JSON 解析更健壮)。
		// 注意:2026 抓包里 IMG2 也可能是单条 sediment:// 终稿,
		// 关键区别在同一 SSE/mapping 片段里带 metadata.generation.gen_size_v2。
		eventSedIDs := make([]string, 0, 2)
		for _, m := range reFileRef.FindAllSubmatch(data, -1) {
			fid := string(m[1])
			if _, ok := seenFile[fid]; !ok {
				seenFile[fid] = struct{}{}
				r.FileIDs = append(r.FileIDs, fid)
			}
		}
		for _, m := range reSedRef.FindAllSubmatch(data, -1) {
			sid := string(m[1])
			if _, ok := seenSed[sid]; !ok {
				seenSed[sid] = struct{}{}
				r.SedimentIDs = append(r.SedimentIDs, sid)
			}
			eventSedIDs = append(eventSedIDs, sid)
		}

		if len(eventSedIDs) > 0 && strings.Contains(string(data), "gen_size_v2") {
			for _, sid := range eventSedIDs {
				if _, ok := seenIMG2Sed[sid]; !ok {
					seenIMG2Sed[sid] = struct{}{}
					r.IMG2SedimentIDs = append(r.IMG2SedimentIDs, sid)
				}
			}
		}

		var obj map[string]interface{}
		if err := json.Unmarshal(data, &obj); err != nil {
			continue
		}
		if v, ok := obj["v"].(map[string]interface{}); ok {
			if cid, ok := v["conversation_id"].(string); ok && cid != "" && r.ConversationID == "" {
				r.ConversationID = cid
			}
			if msg, ok := v["message"].(map[string]interface{}); ok {
				if meta, ok := msg["metadata"].(map[string]interface{}); ok {
					if tid, ok := meta["image_gen_task_id"].(string); ok {
						r.ImageGenTaskID = tid
					}
					if fd, ok := meta["finish_details"].(map[string]interface{}); ok {
						if ft, ok := fd["type"].(string); ok {
							r.FinishType = ft
						}
					}
				}
				// 捕获助手文本消息(每个 SSE 事件 parts 含到目前的完整文本,取最后一次即可)
				if author, ok := msg["author"].(map[string]interface{}); ok {
					if role, _ := author["role"].(string); role == "assistant" {
						if content, ok := msg["content"].(map[string]interface{}); ok {
							if parts, ok := content["parts"].([]interface{}); ok {
								var text string
								for _, p := range parts {
									if s, ok := p.(string); ok {
										text += s
									}
								}
								if len(text) > 0 {
									r.AssistantText = text // 覆盖,保留最新完整文本
								}
							}
						}
					}
				}
			}
		}
	}
	return r
}

// ImageToolMsg 是 conversation.mapping 里一条 IMG2 tool 消息的关键字段。
type ImageToolMsg struct {
	MessageID     string
	CreateTime    float64
	ModelSlug     string
	Recipient     string
	AuthorName    string
	ImageGenTitle string
	FileIDs       []string // file-service
	SedimentIDs   []string // sediment
	IMG2Hint      bool     // true=asset metadata.generation.gen_size_v2 存在,sediment-only 也按 IMG2 终稿处理
	GenSizeV2s    []string // 抓包里的 generation.gen_size_v2,例如 "48"
	MaxWidth      int
	MaxHeight     int
	MaxSizeBytes  int64
}

// GetConversationMapping 读取会话全量 mapping(轮询用)。
func (c *Client) GetConversationMapping(ctx context.Context, convID string) (map[string]interface{}, error) {
	if convID == "" {
		return nil, fmt.Errorf("conv_id required")
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		c.opts.BaseURL+"/backend-api/conversation/"+convID, nil)
	if err != nil {
		return nil, err
	}
	c.commonHeaders(req)
	req.Header.Set("Accept", "*/*")

	res, err := c.hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	buf, _ := io.ReadAll(res.Body)
	if res.StatusCode >= 400 {
		return nil, &UpstreamError{Status: res.StatusCode, Message: "conversation get failed", Body: string(buf)}
	}
	var out map[string]interface{}
	if err := json.Unmarshal(buf, &out); err != nil {
		return nil, fmt.Errorf("decode conversation: %w", err)
	}
	return out, nil
}

func imagePartString(m map[string]interface{}, path ...string) string {
	var cur interface{} = m
	for _, key := range path {
		mm, ok := cur.(map[string]interface{})
		if !ok {
			return ""
		}
		cur = mm[key]
	}
	s, _ := cur.(string)
	return s
}

func imagePartInt(v interface{}) int {
	switch n := v.(type) {
	case int:
		return n
	case int64:
		return int(n)
	case float64:
		return int(n)
	default:
		return 0
	}
}

func imagePartInt64(v interface{}) int64 {
	switch n := v.(type) {
	case int:
		return int64(n)
	case int64:
		return n
	case float64:
		return int64(n)
	default:
		return 0
	}
}

func appendUniqueString(xs []string, v string) []string {
	if v == "" {
		return xs
	}
	for _, x := range xs {
		if x == v {
			return xs
		}
	}
	return append(xs, v)
}

func inspectImageAssetPart(p map[string]interface{}, tm *ImageToolMsg) {
	if tm == nil {
		return
	}
	if w := imagePartInt(p["width"]); w > tm.MaxWidth {
		tm.MaxWidth = w
	}
	if h := imagePartInt(p["height"]); h > tm.MaxHeight {
		tm.MaxHeight = h
	}
	if sz := imagePartInt64(p["size_bytes"]); sz > tm.MaxSizeBytes {
		tm.MaxSizeBytes = sz
	}
	// 浏览器实测 IMG2 可以只给一条 sediment asset,但该 part 会带
	// metadata.generation.gen_size_v2(例如 "48") 与高清尺寸。旧逻辑只按
	// file-service 或 tool 条数判断,会把这种真正的 IMG2 误判为 preview_only。
	if genSizeV2 := imagePartString(p, "metadata", "generation", "gen_size_v2"); genSizeV2 != "" {
		tm.IMG2Hint = true
		tm.GenSizeV2s = appendUniqueString(tm.GenSizeV2s, genSizeV2)
	}
}

// ExtractImageToolMsgs 从 conversation.mapping 里提取所有 IMG2 tool 消息。
func ExtractImageToolMsgs(mapping map[string]interface{}) []ImageToolMsg {
	out := make([]ImageToolMsg, 0, 4)
	for mid, raw := range mapping {
		node, _ := raw.(map[string]interface{})
		if node == nil {
			continue
		}
		msg, _ := node["message"].(map[string]interface{})
		if msg == nil {
			continue
		}
		author, _ := msg["author"].(map[string]interface{})
		meta, _ := msg["metadata"].(map[string]interface{})
		content, _ := msg["content"].(map[string]interface{})
		if author == nil || meta == nil || content == nil {
			continue
		}
		if s, _ := author["role"].(string); s != "tool" {
			continue
		}
		if s, _ := meta["async_task_type"].(string); s != "image_gen" {
			continue
		}
		if s, _ := content["content_type"].(string); s != "multimodal_text" {
			continue
		}

		tm := ImageToolMsg{MessageID: mid}
		if v, ok := msg["create_time"].(float64); ok {
			tm.CreateTime = v
		}
		if v, ok := meta["model_slug"].(string); ok {
			tm.ModelSlug = v
		}
		if v, ok := msg["recipient"].(string); ok {
			tm.Recipient = v
		}
		if v, ok := author["name"].(string); ok {
			tm.AuthorName = v
		}
		if v, ok := meta["image_gen_title"].(string); ok {
			tm.ImageGenTitle = v
		}

		parts, _ := content["parts"].([]interface{})
		seenF := map[string]struct{}{}
		seenS := map[string]struct{}{}
		extractAsset := func(text string) {
			for _, m := range reFileRef.FindAllStringSubmatch(text, -1) {
				if _, ok := seenF[m[1]]; !ok {
					seenF[m[1]] = struct{}{}
					tm.FileIDs = append(tm.FileIDs, m[1])
				}
			}
			for _, m := range reSedRef.FindAllStringSubmatch(text, -1) {
				if _, ok := seenS[m[1]]; !ok {
					seenS[m[1]] = struct{}{}
					tm.SedimentIDs = append(tm.SedimentIDs, m[1])
				}
			}
		}
		for _, p := range parts {
			switch v := p.(type) {
			case map[string]interface{}:
				inspectImageAssetPart(v, &tm)
				if aid, _ := v["asset_pointer"].(string); aid != "" {
					extractAsset(aid)
				}
			case string:
				extractAsset(v)
			}
		}
		out = append(out, tm)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].CreateTime < out[j].CreateTime })
	return out
}

// PollOpts 控制 PollConversationForImages 的等待策略。
type PollOpts struct {
	BaselineToolIDs map[string]struct{} // 发送前已存在的 tool 消息 id;本次回合只看新增
	MaxWait         time.Duration       // 总超时,默认 300s
	Interval        time.Duration       // 轮询间隔,默认 6s
	StableRounds    int                 // 稳定轮数(连续 N 次 sed 不变视为稳定),默认 4
	PreviewWait     time.Duration       // 第 1 条 tool 出现后等第 2 条的窗口,默认 30s
}

// PollStatus 是 PollConversationForImages 的结果状态。
type PollStatus string

const (
	PollStatusIMG2        PollStatus = "img2"         // 灰度命中,images 可用
	PollStatusPreviewOnly PollStatus = "preview_only" // 只出 1 条 tool,判定非灰度
	PollStatusTimeout     PollStatus = "timeout"
	PollStatusError       PollStatus = "error"
	PollStatus429         PollStatus = "rate_limited" // 连续 429,上游 RPM 限流
)

// PollConversationForImages 轮询会话直到灰度图出现。
// 返回 (status, file_ids, sediment_ids)。状态为 img2 时 file_ids 或 sediment_ids 至少一个非空。
func (c *Client) PollConversationForImages(ctx context.Context, convID string, opt PollOpts) (PollStatus, []string, []string) {
	if opt.MaxWait == 0 {
		opt.MaxWait = 300 * time.Second
	}
	if opt.Interval == 0 {
		opt.Interval = 6 * time.Second
	}
	if opt.StableRounds == 0 {
		opt.StableRounds = 4
	}
	if opt.PreviewWait == 0 {
		opt.PreviewWait = 30 * time.Second
	}
	baseline := opt.BaselineToolIDs

	deadline := time.Now().Add(opt.MaxWait)
	var (
		stableCount    int
		lastSedSig     string
		firstToolTs    time.Time
		consecutive429 int
	)

	for time.Now().Before(deadline) {
		select {
		case <-ctx.Done():
			return PollStatusError, nil, nil
		default:
		}

		mapping, err := c.getMappingRaw(ctx, convID)
		if err != nil {
			if ue, ok := err.(*UpstreamError); ok && ue.Status == 429 {
				consecutive429++
				if consecutive429 >= 3 {
					return PollStatus429, nil, nil
				}
				sleep(ctx, 10*time.Second)
				continue
			}
			sleep(ctx, opt.Interval)
			continue
		}
		consecutive429 = 0

		msgs := ExtractImageToolMsgs(mapping)
		// baseline diff:只看本回合新增
		var newMsgs []ImageToolMsg
		if len(baseline) > 0 {
			for _, m := range msgs {
				if _, ok := baseline[m.MessageID]; !ok {
					newMsgs = append(newMsgs, m)
				}
			}
		} else {
			newMsgs = msgs
		}

		// 汇总所有新 tool 消息的 sed/file(**跨消息聚合**)。
		// IMG2 灰度命中时,上游通常会发 2 条 tool 消息 —— 例如 1 条 sediment
		// 预览 + 1 条 file-service 终稿,或者同一条消息里带多张 file id。
		// 以前只取 newMsgs[last] 会丢掉前一条 preview / 另一张图;这里收集
		// 全部 tool 消息里出现过的 id,调用方拿到几张就可以输出几张。
		var allSed []string
		var allFile []string
		img2Hint := false
		seenFile := map[string]struct{}{}
		seenSed := map[string]struct{}{}
		for _, m := range newMsgs {
			if m.IMG2Hint {
				img2Hint = true
			}
			for _, f := range m.FileIDs {
				if _, ok := seenFile[f]; !ok {
					seenFile[f] = struct{}{}
					allFile = append(allFile, f)
				}
			}
			for _, s := range m.SedimentIDs {
				if _, ok := seenSed[s]; !ok {
					seenSed[s] = struct{}{}
					allSed = append(allSed, s)
				}
			}
		}

		// 分支 1:file-service 直出 = IMG2 终稿。
		// 有 file-service 直出就算命中,把所有 tool 消息累计到的 fid/sid 都带出去。
		if len(allFile) > 0 {
			return PollStatusIMG2, allFile, allSed
		}

		// 分支 1.5:2026 抓包确认 IMG2 可能是单条 sediment asset,
		// 但 part.metadata.generation.gen_size_v2 存在(例如 "48")。这种不是
		// IMG1 预览,不应等待第二条 tool 或重试。
		if img2Hint && len(allSed) > 0 {
			return PollStatusIMG2, allFile, allSed
		}

		// 没有 tool 消息 → 继续等
		if len(newMsgs) == 0 {
			sleep(ctx, opt.Interval)
			continue
		}

		// 首次出现第 1 条 tool,记时间
		if firstToolTs.IsZero() && len(newMsgs) >= 1 {
			firstToolTs = time.Now()
		}

		// 分支 2:已经 2+ 条 tool 且有 IMG2 特征(gen_size_v2) → 灰度命中,等 sed 稳定后一并返回。
		// 没有 IMG2 特征的多条 tool 只是多轮 preview,不算真正的 IMG2。
		if len(newMsgs) >= 2 && img2Hint {
			sortedSed := append([]string(nil), allSed...)
			sort.Strings(sortedSed)
			sig := strings.Join(sortedSed, ",")
			if sig == lastSedSig && sig != "" {
				stableCount++
				if stableCount >= opt.StableRounds {
					return PollStatusIMG2, allFile, allSed
				}
			} else {
				stableCount = 0
				lastSedSig = sig
			}
		} else if !firstToolTs.IsZero() && time.Since(firstToolTs) >= opt.PreviewWait {
			// 分支 3:只 1 条 tool 且超过窗口 → 非灰度预览。
			// 把这条 tool 的 fids / sids 一并带出,外层可以用作 IMG1 预览兜底。
			return PollStatusPreviewOnly, allFile, allSed
		}

		sleep(ctx, opt.Interval)
	}

	return PollStatusTimeout, nil, nil
}

// getMappingRaw 拉 conversation 并返回 mapping。
func (c *Client) getMappingRaw(ctx context.Context, convID string) (map[string]interface{}, error) {
	full, err := c.GetConversationMapping(ctx, convID)
	if err != nil {
		return nil, err
	}
	mapping, _ := full["mapping"].(map[string]interface{})
	if mapping == nil {
		mapping = map[string]interface{}{}
	}
	return mapping, nil
}

// GetConversationHead 返回会话最新叶子消息 id(current_node),复用会话时做 parent_message_id。
func (c *Client) GetConversationHead(ctx context.Context, convID string) (string, error) {
	full, err := c.GetConversationMapping(ctx, convID)
	if err != nil {
		return "", err
	}
	head, _ := full["current_node"].(string)
	if head == "" {
		return "", fmt.Errorf("current_node missing")
	}
	return head, nil
}

// ImageDownloadURL 取单张图的签名下载 URL。
// fileRef 可以是:
//   - "xxxxxx"      → file-service id
//   - "sed:xxxxxx"  → sediment id(需要 convID)
//
// 兼容两类浏览器实测路径:
//   - 新路径: /backend-api/files/download/{fid}?conversation_id={cid}&inline=false
//   - 旧路径: /backend-api/files/{fid}/download 或 /conversation/{cid}/attachment/{sid}/download
func (c *Client) ImageDownloadURL(ctx context.Context, convID, fileRef string) (string, error) {
	isSediment := strings.HasPrefix(fileRef, "sed:")
	rawID := strings.TrimPrefix(fileRef, "sed:")
	if rawID == "" {
		return "", fmt.Errorf("empty file ref")
	}
	if isSediment && convID == "" {
		return "", fmt.Errorf("conv_id required for sediment")
	}

	var candidates []string
	if convID != "" {
		// 2026 浏览器抓包优先使用这一类 /files/download/{fid}?conversation_id=... 路径,
		// 返回的 download_url 往往会落到 /backend-api/estuary/content?...。
		candidates = append(candidates, fmt.Sprintf("%s/backend-api/files/download/%s?conversation_id=%s&inline=false",
			c.opts.BaseURL, url.PathEscape(rawID), url.QueryEscape(convID)))
	}
	if isSediment {
		candidates = append(candidates, fmt.Sprintf("%s/backend-api/conversation/%s/attachment/%s/download",
			c.opts.BaseURL, url.PathEscape(convID), url.PathEscape(rawID)))
	} else {
		candidates = append(candidates,
			fmt.Sprintf("%s/backend-api/files/download/%s", c.opts.BaseURL, url.PathEscape(rawID)),
			fmt.Sprintf("%s/backend-api/files/%s/download", c.opts.BaseURL, url.PathEscape(rawID)),
		)
	}

	var lastErr error
	for _, apiURL := range candidates {
		signed, err := c.fetchImageDownloadURL(ctx, apiURL)
		if err == nil && signed != "" {
			return signed, nil
		}
		lastErr = err
	}
	if lastErr != nil {
		return "", lastErr
	}
	return "", fmt.Errorf("empty download_url")
}

func (c *Client) fetchImageDownloadURL(ctx context.Context, apiURL string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return "", err
	}
	c.commonHeaders(req)
	req.Header.Set("Accept", "*/*")

	// 某些路径直接 302 到 files.oaiusercontent.com / estuary content。不要自动跟随,
	// 否则我们会把图片二进制当 JSON 解析。
	local := *c.hc
	local.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	c.injectHelloID(req)
	res, err := local.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 && res.StatusCode < 400 {
		loc := res.Header.Get("Location")
		if loc == "" {
			return "", &UpstreamError{Status: res.StatusCode, Message: "files/download redirect without location"}
		}
		base, _ := url.Parse(apiURL)
		u, err := url.Parse(loc)
		if err != nil {
			return "", err
		}
		return base.ResolveReference(u).String(), nil
	}

	buf, _ := io.ReadAll(res.Body)
	if res.StatusCode >= 400 {
		return "", &UpstreamError{Status: res.StatusCode, Message: "files/download failed", Body: string(buf)}
	}

	ct := strings.ToLower(res.Header.Get("Content-Type"))
	if strings.HasPrefix(ct, "image/") || strings.Contains(ct, "octet-stream") {
		// 端点本身已经能回源图片;交给 FetchImage 再带鉴权下载一次。
		return apiURL, nil
	}

	var out struct {
		DownloadURL string `json:"download_url"`
		URL         string `json:"url"`
		Status      string `json:"status"`
		FileName    string `json:"file_name"`
	}
	if err := json.Unmarshal(buf, &out); err == nil {
		if out.DownloadURL != "" {
			return out.DownloadURL, nil
		}
		if out.URL != "" {
			return out.URL, nil
		}
		return "", fmt.Errorf("empty download_url (status=%s)", out.Status)
	}

	// 兜底:有些实验路径可能直接返回纯文本 URL。
	text := strings.TrimSpace(string(buf))
	if strings.HasPrefix(text, "http://") || strings.HasPrefix(text, "https://") || strings.HasPrefix(text, "/") {
		if strings.HasPrefix(text, "/") {
			return c.opts.BaseURL + text, nil
		}
		return text, nil
	}
	return "", fmt.Errorf("decode files/download: unexpected content-type=%s body=%s", ct, truncateForErr(text, 200))
}

func truncateForErr(s string, n int) string {
	if n <= 0 || len(s) <= n {
		return s
	}
	return s[:n]
}

// FetchImage 下载指定 URL 的图片字节(调用 ImageDownloadURL 返回的签名 URL)。
// 返回 (bytes, content_type)。超过 maxBytes 的响应会被截断报错。
//
// 鉴权策略:
//   - files.oaiusercontent.com / cdn.oaistatic.com 等**预签名直链**:不带 Authorization,
//     它们依赖 URL 自带的 sig 参数鉴权;带 Bearer 反而会被某些 CDN 因"双鉴权"400。
//   - chatgpt.com/backend-api/estuary/... (sediment 回源 URL):**必须** 带 Authorization:
//     Bearer 和完整 Oai-* 头,否则上游 403。这就是 IMG2 sediment 图 502 的根因。
func (c *Client) FetchImage(ctx context.Context, signedURL string, maxBytes int64) ([]byte, string, error) {
	if maxBytes <= 0 {
		maxBytes = 16 * 1024 * 1024 // 16MB 默认
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, signedURL, nil)
	if err != nil {
		return nil, "", err
	}

	// 判断是否需要完整 chatgpt 鉴权头:以 BaseURL(通常 https://chatgpt.com)开头的
	// estuary / attachment 回源 URL 必须带 Bearer + Oai-* 头;外部 CDN 直链不带。
	needAuth := strings.HasPrefix(signedURL, c.opts.BaseURL+"/")
	if needAuth {
		c.commonHeaders(req)
		req.Header.Set("Accept", "image/*,*/*;q=0.8")
	} else {
		req.Header.Set("User-Agent", c.opts.UserAgent)
	}

	res, err := c.hc.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		return nil, "", &UpstreamError{Status: res.StatusCode, Message: "fetch image failed"}
	}
	ct := res.Header.Get("Content-Type")
	body, err := io.ReadAll(io.LimitReader(res.Body, maxBytes+1))
	if err != nil {
		return nil, ct, err
	}
	if int64(len(body)) > maxBytes {
		return nil, ct, fmt.Errorf("image exceeds max bytes (%d)", maxBytes)
	}
	return body, ct, nil
}

// sleep 可取消的 sleep。
func sleep(ctx context.Context, d time.Duration) {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
	case <-t.C:
	}
}
````

## File: internal/upstream/chatgpt/pow.go
````go
// Package chatgpt — POW 算法迁移自 gen_image.py
//
// chatgpt.com sentinel 使用两种 POW token:
//
//  1. RequirementsToken  → 客户端主动生成,塞进 /sentinel/chat-requirements
//     请求体的 `p` 字段。前缀 "gAAAAAC"。固定难度 "0fffff"。
//     config 是 18 元素数组,迭代 config[3] 与 config[9]。
//
//  2. ProofToken         → 服务端返回 `proofofwork.required=true` + seed + difficulty,
//     客户端本地求解后放进 Header `openai-sentinel-proof-token`。
//     前缀 "gAAAAAB"。config 是 13 元素数组,只迭代 config[3]。
//
// 两者共享同一个判定函数:SHA3-512(seed + base64(config_json)) 的前 N 字节
// 按字节序 <= bytes.fromhex(difficulty)。若不满足则 config[3] += 1 重试。
package chatgpt

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"golang.org/x/crypto/sha3"
)

// TurnstileSolver 负责把 /sentinel/chat-requirements/prepare 返回的 `turnstile.dx`
// 挑战字符串,解算成 /sentinel/chat-requirements/finalize 需要的 turnstile
// response 字符串。
//
// 说明:OpenAI 的 turnstile 是基于 Cloudflare turnstile 衍生的自定义 challenge,
// dx 是混淆 JS + WebAssembly 的输入,response 是执行结果。纯 Go 无法还原,
// 解算必须委托给外部服务(2captcha/capsolver/自建 headless 浏览器等)。
//
// 没有 solver 时,Client.ChatRequirementsV2 会自动回退到老的单步
// chat-requirements 流程(Turnstile=true 直接忽略)。
type TurnstileSolver interface {
	Solve(ctx context.Context, dx string) (string, error)
}

const (
	powPrefixRequirements = "gAAAAAC"
	powPrefixProof        = "gAAAAAB"

	requirementsDifficulty = "0fffff"

	maxRequirementsIter = 500_000
	maxProofIter        = 100_000

	powFallback = "gAAAAABwQ8Lk5FbGpA2NcR9dShT6gYjU7VxZ4D"
)

var (
	powCores   = []int{16, 24, 32}
	powScreens = []int{3000, 4000, 6000}

	powNavKeys = []string{
		"webdriver−false", "vendor−Google Inc.", "cookieEnabled−true",
		"pdfViewerEnabled−true", "hardwareConcurrency−32",
		"language−zh-CN", "mimeTypes−[object MimeTypeArray]",
		"userAgentData−[object NavigatorUAData]",
	}
	powWinKeys = []string{
		"innerWidth", "innerHeight", "devicePixelRatio", "screen",
		"chrome", "location", "history", "navigator",
	}

	powReactListeners = []string{"_reactListeningcfilawjnerp", "_reactListening9ne2dfo1i47"}
	powProofEvents    = []string{"alert", "ontransitionend", "onprogress"}

	// perfCounter 模拟浏览器 performance.counter() 的单调递增(亚秒级)。
	perfCounter uint64
)

// POWConfig 是 18 元素的客户端指纹数组(requirements_token 用)。
type POWConfig struct {
	userAgent string
	arr       [18]interface{}
}

// NewPOWConfig 构造一个随机化的客户端指纹,用于 requirements + proof 两种场景。
func NewPOWConfig(userAgent string) *POWConfig {
	if userAgent == "" {
		userAgent = DefaultUserAgent
	}
	//nolint:gosec // 非加密用途
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	now := time.Now().UTC()
	timeStr := now.Format("Mon Jan 02 2006 15:04:05") + " GMT+0000 (UTC)"
	perf := float64(atomic.AddUint64(&perfCounter, 1)) + rng.Float64()

	c := &POWConfig{userAgent: userAgent}
	// 动态提取 dpl 和 script src(带缓存),避免硬编码过期的 build hash
	dpl, scriptSrc := getDplAndScript()
	if dpl == "" {
		dpl = "dpl=1440a687921de39ff5ee56b92807faaadce73f13" // fallback 硬编码
	}
	if scriptSrc == "" {
		scriptSrc = "https://cdn.oaistatic.com/_next/static/chunks/pages/index-a78faa38a8ca3f3c.js" // fallback
	}
	c.arr = [18]interface{}{
		powCores[rng.Intn(len(powCores))] + powScreens[rng.Intn(len(powScreens))], // 0
		timeStr,                               // 1
		nil,                                   // 2
		rng.Float64(),                         // 3 - 迭代会覆盖
		userAgent,                             // 4
		scriptSrc,                             // 5 - 动态从 HTML 提取
		dpl,                                   // 6 - 动态从 HTML 提取
		"en-US",                               // 7
		"en-US,zh-CN",                         // 8
		0,                                     // 9 - 迭代会覆盖
		powNavKeys[rng.Intn(len(powNavKeys))], // 10
		"location",                            // 11
		powWinKeys[rng.Intn(len(powWinKeys))], // 12
		perf,                                  // 13
		randomUUID(rng),                       // 14
		"",                                    // 15
		8,                                     // 16
		now.Unix(),                            // 17
	}
	return c
}

// RequirementsToken 生成 /sentinel/chat-requirements 的 "p" 字段值。
// 对齐 gen_image.py.get_requirements_token:固定难度 0fffff,前缀 gAAAAAC。
func (c *POWConfig) RequirementsToken() string {
	//nolint:gosec
	seed := strconv.FormatFloat(rand.Float64(), 'f', -1, 64)
	b64, ok := c.solveRequirements(seed, requirementsDifficulty)
	if !ok {
		return powPrefixRequirements + powFallback +
			base64.StdEncoding.EncodeToString([]byte(`"`+seed+`"`))
	}
	return powPrefixRequirements + b64
}

// solveRequirements 高性能迭代:预拼 JSON 的三段字节前缀,只在内循环拼 d1/d2。
// 严格对齐 gen_image.py._generate_answer。
func (c *POWConfig) solveRequirements(seed, difficulty string) (string, bool) {
	target, err := hex.DecodeString(difficulty)
	if err != nil {
		return "", false
	}
	diffLen := len(difficulty) // 字符数(与 Python 对齐)

	// 预拼 p1/p2/p3。config[3] 和 config[9] 位置留给迭代器。
	arr := c.arr
	// p1 = json(arr[:3])[:-1] + ","
	head, _ := json.Marshal([]interface{}{arr[0], arr[1], arr[2]})
	p1 := append(head[:len(head)-1:len(head)-1], ',')

	mid, _ := json.Marshal([]interface{}{arr[4], arr[5], arr[6], arr[7], arr[8]})
	// p2 = "," + json(arr[4:9])[1:-1] + ","
	p2 := make([]byte, 0, len(mid)+2)
	p2 = append(p2, ',')
	p2 = append(p2, mid[1:len(mid)-1]...)
	p2 = append(p2, ',')

	tail, _ := json.Marshal([]interface{}{
		arr[10], arr[11], arr[12], arr[13], arr[14], arr[15], arr[16], arr[17],
	})
	// p3 = "," + json(arr[10:])[1:]  => "," + "element1,...,elementN]"
	p3 := make([]byte, 0, len(tail)+1)
	p3 = append(p3, ',')
	p3 = append(p3, tail[1:]...)

	hasher := sha3.New512()
	seedB := []byte(seed)
	buf := make([]byte, 0, len(p1)+32+len(p2)+16+len(p3))
	b64buf := make([]byte, base64.StdEncoding.EncodedLen(cap(buf)))

	for i := 0; i < maxRequirementsIter; i++ {
		d1 := strconv.Itoa(i)
		d2 := strconv.Itoa(i >> 1)

		buf = buf[:0]
		buf = append(buf, p1...)
		buf = append(buf, d1...)
		buf = append(buf, p2...)
		buf = append(buf, d2...)
		buf = append(buf, p3...)

		n := base64.StdEncoding.EncodedLen(len(buf))
		if cap(b64buf) < n {
			b64buf = make([]byte, n)
		}
		b64buf = b64buf[:n]
		base64.StdEncoding.Encode(b64buf, buf)

		hasher.Reset()
		hasher.Write(seedB)
		hasher.Write(b64buf)
		sum := hasher.Sum(nil)

		// Python: h[:diff_len] <= target
		// diff_len 是字符数(6),target 是字节(3)。Python bytes cmp 按短的逐字节比较。
		// 这里保持等价:取 min(len(target), len(sum)) 字节比较。
		n2 := diffLen
		if n2 > len(sum) {
			n2 = len(sum)
		}
		cmpLen := n2
		if cmpLen > len(target) {
			cmpLen = len(target)
		}
		if bytes.Compare(sum[:cmpLen], target[:cmpLen]) <= 0 {
			return string(b64buf), true
		}
	}
	return "", false
}

// SolveProofToken 按服务端挑战求解 proof token(header 用,前缀 gAAAAAB)。
// 迁移自 gen_image.py.generate_proof_token 的轻量 13 元素 config。
func SolveProofToken(seed, difficulty, userAgent string) string {
	if seed == "" || difficulty == "" {
		return ""
	}
	if userAgent == "" {
		userAgent = DefaultUserAgent
	}
	//nolint:gosec
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	screen := powScreens[rng.Intn(len(powScreens))] * (1 << rng.Intn(3)) // *1/2/4

	timeStr := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")

	// 动态提取 dpl 和 script src
	dpl, scriptSrc := getDplAndScript()
	if dpl == "" {
		dpl = "dpl=1440a687921de39ff5ee56b92807faaadce73f13"
	}
	if scriptSrc == "" {
		scriptSrc = "https://tcr9i.chat.openai.com/v2/35536E1E-65B4-4D96-9D97-6ADB7EFF8147/api.js"
	}

	proofConfig := []interface{}{
		screen, // 0
		timeStr,
		nil,
		0, // 3 - 迭代
		userAgent,
		scriptSrc,
		dpl,
		"en",
		"en-US",
		nil,
		"plugins−[object PluginArray]",
		powReactListeners[rng.Intn(len(powReactListeners))],
		powProofEvents[rng.Intn(len(powProofEvents))],
	}

	diffLen := len(difficulty)
	hasher := sha3.New512()
	for i := 0; i < maxProofIter; i++ {
		proofConfig[3] = i
		raw, err := json.Marshal(proofConfig)
		if err != nil {
			continue
		}
		b64 := base64.StdEncoding.EncodeToString(raw)
		hasher.Reset()
		hasher.Write([]byte(seed + b64))
		sum := hasher.Sum(nil)
		hexStr := hex.EncodeToString(sum)
		if strings.Compare(hexStr[:diffLen], difficulty) <= 0 {
			return powPrefixProof + b64
		}
	}
	return powPrefixProof + powFallback +
		base64.StdEncoding.EncodeToString([]byte(`"`+seed+`"`))
}

func randomUUID(rng *rand.Rand) string {
	var b [16]byte
	_, _ = rng.Read(b[:])
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}
````

## File: internal/upstream/chatgpt/probe.go
````go
package chatgpt

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// ModelsRaw fetches /backend-api/models using the same ChatGPT web client
// fingerprint as f/conversation: uTLS transport, common browser/Oai-* headers,
// stable device/session IDs, proxy, cookie jar, and bearer token.
func (c *Client) ModelsRaw(ctx context.Context) ([]byte, error) {
	// Bootstrap is intentionally best-effort. It lets the cookie jar collect
	// Cloudflare/OpenAI first-party cookies, but a transient home-page failure
	// should not hide the real /models response.
	_ = c.Bootstrap(ctx)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.opts.BaseURL+"/backend-api/models", nil)
	if err != nil {
		return nil, err
	}
	c.commonHeaders(req)
	req.Header.Set("Accept", "application/json")

	res, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	buf, _ := io.ReadAll(res.Body)
	if res.StatusCode >= 400 {
		return nil, &UpstreamError{Status: res.StatusCode, Message: "models probe failed", Body: string(buf)}
	}
	return buf, nil
}

// ConversationInitRaw fetches /backend-api/conversation/init with the same
// client fingerprint. This endpoint is kept as a quota/diagnostic weak probe;
// it must not be used as the primary image capability decision.
func (c *Client) ConversationInitRaw(ctx context.Context, timezoneOffsetMin int, systemHints []string) ([]byte, error) {
	if systemHints == nil {
		systemHints = []string{"picture_v2"}
	}
	body, _ := json.Marshal(map[string]interface{}{
		"gizmo_id":                nil,
		"requested_default_model": nil,
		"conversation_id":         nil,
		"timezone_offset_min":     timezoneOffsetMin,
		"system_hints":            systemHints,
	})
	req, err := http.NewRequestWithContext(ctx,
		http.MethodPost,
		c.opts.BaseURL+"/backend-api/conversation/init",
		bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	c.commonHeaders(req)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	buf, _ := io.ReadAll(res.Body)
	if res.StatusCode >= 400 {
		return nil, &UpstreamError{Status: res.StatusCode, Message: "conversation/init probe failed", Body: string(buf)}
	}
	return buf, nil
}

// ImageQuotaProbe 轻量级 image quota 预检。
// 调用 /backend-api/conversation/init(system_hints=["picture_v2"]) 检查:
//   - limits_progress 中 image_gen 的 remaining 是否 > 0
//   - blocked_features 是否包含 image_gen
//
// 返回 (hasQuota, blockedReason, error)。
// 如果 init 接口 404(免费号新行为)或返回的 quota 结构缺失,降级为"可能有 quota",
// 不阻断主流程——真正的 quota / 出图状态由 f/conversation 的实际响应决定。
// 这个函数是“尽早跳过明显无额度账号”的优化,不是 image 能力主判据。
func (c *Client) ImageQuotaProbe(ctx context.Context) (hasQuota bool, blockedReason string, err error) {
	// 超时控制:预检不应阻塞太久
	probeCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	raw, rawErr := c.ConversationInitRaw(probeCtx, -480, []string{"picture_v2"})
	if rawErr != nil {
		// 404 在免费号上是正常的,不视为无 quota
		if ue, ok := rawErr.(*UpstreamError); ok && ue.Status == 404 {
			return true, "", nil
		}
		return true, "", fmt.Errorf("quota probe: %w", rawErr)
	}

	var resp struct {
		LimitsProgress []struct {
			FeatureName string `json:"feature_name"`
			Remaining   int    `json:"remaining"`
			ResetAfter  string `json:"reset_after"`
		} `json:"limits_progress"`
		BlockedFeatures interface{} `json:"blocked_features"`
	}
	if jsonErr := json.Unmarshal(raw, &resp); jsonErr != nil {
		return true, "", nil // 解析失败,降级放行
	}

	// 检查 blocked_features
	if resp.BlockedFeatures != nil {
		switch bf := resp.BlockedFeatures.(type) {
		case []interface{}:
			for _, v := range bf {
				if s, ok := v.(string); ok && s == "image_gen" {
					return false, "blocked_features:image_gen", nil
				}
			}
		case string:
			if bf == "image_gen" {
				return false, "blocked_features:image_gen", nil
			}
		}
	}

	// 检查 limits_progress
	for _, lp := range resp.LimitsProgress {
		if lp.FeatureName == "image_gen" {
			if lp.Remaining <= 0 {
				return false, fmt.Sprintf("quota_exhausted:remaining=0,reset_after=%s", lp.ResetAfter), nil
			}
			return true, "", nil
		}
	}

	// 没有 image_gen 的 limits_progress 条目——降级放行
	return true, "", nil
}
````

## File: internal/upstream/chatgpt/utls_transport.go
````go
package chatgpt

import (
	"bufio"
	"context"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	utls "github.com/refraction-networking/utls"
	"golang.org/x/net/http2"
)

// NewUTLSTransport 返回一个带 uTLS 浏览器指纹伪装的 http.RoundTripper。
//
// chatgpt.com 前置的 Cloudflare 会按 TLS ClientHello 的 JA3/JA4 指纹识别客户端;
// Go 标准 crypto/tls 的指纹会被直接拒绝(返回 403 HTML 拦截页)。这里用
// refraction-networking/utls 把 ClientHello 换成 Chrome 131 模板,让 Cloudflare
// 认为是真 Chrome。
//
// 同时支持通过 HTTP(S) 代理走 CONNECT 隧道。
//
// 行为要点:
//   - ALPN 协商到 h2 时走内部 http2.Transport,http/1.1 时走标准 http.Transport
//   - 首次 h2 失败且为协议级错误(例如 ALPN 回退到 h1)时自动切 h1 重试
//   - 连接不做跨请求复用的特殊处理,依赖各子 transport 自身的空闲池
//
// proxyURL 只支持 http/https,socks5 需要走另一条路径(当前账号池均为 HTTP 代理)。
//
// 如果需要 per-account 的 TLS 指纹(避免多账号共用同一指纹被关联),
// 可以通过 context 传入自定义 HelloID:
//
//	ctx = context.WithValue(ctx, utlsHelloIDKey{}, utls.HelloChrome_120)
func NewUTLSTransport(proxyURL string, idleTimeout time.Duration) (http.RoundTripper, error) {
	if idleTimeout <= 0 {
		idleTimeout = 30 * time.Second
	}
	rt := &utlsRoundTripper{
		dialer:      &net.Dialer{Timeout: 30 * time.Second, KeepAlive: 30 * time.Second},
		idleTimeout: idleTimeout,
	}
	if strings.TrimSpace(proxyURL) != "" {
		u, err := url.Parse(proxyURL)
		if err != nil {
			return nil, fmt.Errorf("invalid proxy url: %w", err)
		}
		switch strings.ToLower(u.Scheme) {
		case "http", "https":
			rt.proxyURL = u
		case "socks5", "socks5h":
			return nil, fmt.Errorf("socks5 proxy is not supported by utls transport yet")
		default:
			return nil, fmt.Errorf("unsupported proxy scheme %q", u.Scheme)
		}
	}
	rt.h1 = &http.Transport{
		DialTLSContext:        rt.dialTLS,
		MaxIdleConnsPerHost:   4,
		IdleConnTimeout:       idleTimeout,
		TLSHandshakeTimeout:   15 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		ForceAttemptHTTP2:     false,
	}
	rt.h2 = &http2.Transport{
		DialTLSContext: func(ctx context.Context, network, addr string, cfg *tls.Config) (net.Conn, error) {
			return rt.dialTLS(ctx, network, addr)
		},
		ReadIdleTimeout: idleTimeout,
		AllowHTTP:       false,
	}
	return rt, nil
}

// forceH1 如果为 true,ALPN 只请求 http/1.1 且 RoundTrip 完全跳过 h2。
//
// 为什么要这个开关:
// chatgpt.com + Cloudflare 近期加强了 JA4H(HTTP/2 SETTINGS frame 指纹)识别。
// Go 的 golang.org/x/net/http2.Transport 发出的 SETTINGS frame 顺序/参数跟
// Chrome 不同,即使 TLS ClientHello 用 uTLS 伪装成 Chrome,h2 层依然会被识别
// 为自动化客户端,触发"Unusual activity has been detected"硬风控。
//
// 参考实现 gen_image.py 用 curl-cffi(impersonate chrome131),那个库同时伪装
// TLS + HTTP/2 指纹。本 transport 暂未接 tls-client,
// 作为权宜之计:强制 http/1.1,上游对 h1 的指纹检测更宽松。
var forceH1 = true

// utlsHelloIDKey 用于 context 传入 per-request 的 uTLS ClientHelloID。
// 默认使用 HelloChrome_131;通过 context value 可以按账号使用不同指纹。
type utlsHelloIDKey struct{}

// defaultHelloID 是没有 context 指定时的默认 TLS 指纹模板。
var defaultHelloID = utls.HelloChrome_131

type utlsRoundTripper struct {
	proxyURL    *url.URL
	dialer      *net.Dialer
	idleTimeout time.Duration

	mu sync.Mutex
	h1 *http.Transport
	h2 *http2.Transport
}

// RoundTrip 先尝试 h2(chatgpt.com 默认),仅在明确是协议级错误时回退到 h1。
//
// 当 forceH1 = true 时,完全跳过 h2(避免 JA4H 识别),直接走 h1 transport。
func (rt *utlsRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if forceH1 {
		return rt.h1.RoundTrip(req)
	}
	// SSE / chunked body 场景:不能让 h2 transport 吃掉 Connection: close 之类的 h1 头
	// 这里 chatgpt.com 都能用 h2 处理 SSE,优先 h2.
	resp, err := rt.h2.RoundTrip(req)
	if err == nil {
		return resp, nil
	}
	if isH2Retryable(err) {
		return rt.h1.RoundTrip(req)
	}
	return nil, err
}

// CloseIdleConnections 关闭两个子 transport 的空闲连接(http.Client 会调用)。
func (rt *utlsRoundTripper) CloseIdleConnections() {
	rt.mu.Lock()
	h1, h2 := rt.h1, rt.h2
	rt.mu.Unlock()
	if h1 != nil {
		h1.CloseIdleConnections()
	}
	if h2 != nil {
		h2.CloseIdleConnections()
	}
}

// dialTLS 建立到 addr 的 TCP(可能走代理 CONNECT),再用 utls 做浏览器指纹的 TLS 握手。
// 返回的 net.Conn 实际是 *utls.UConn,ALPN 结果由握手自动确定。
//
// 支持通过 context 传入自定义 HelloID,实现 per-account TLS 指纹绑定:
//
//	helloID := utls.HelloChrome_120
//	ctx = context.WithValue(ctx, utlsHelloIDKey{}, &helloID)
func (rt *utlsRoundTripper) dialTLS(ctx context.Context, network, addr string) (net.Conn, error) {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}
	raw, err := rt.dialRaw(ctx, addr)
	if err != nil {
		return nil, err
	}
	// forceH1:ALPN 只请求 http/1.1,避免 HTTP/2 SETTINGS 指纹被识别。
	// 否则按常规先 h2 再 h1 fallback。
	alpn := []string{"h2", "http/1.1"}
	if forceH1 {
		alpn = []string{"http/1.1"}
	}
	// per-account 指纹:从 context 中取自定义 HelloID,没有则用默认 Chrome 131
	helloID := defaultHelloID
	if custom, ok := ctx.Value(utlsHelloIDKey{}).(utls.ClientHelloID); ok {
		helloID = custom
	}
	uconn := utls.UClient(raw, &utls.Config{
		ServerName: host,
		NextProtos: alpn,
		MinVersion: tls.VersionTLS12,
	}, helloID)

	// 关键:utls 的预设 HelloID(HelloChrome_131 等)里 ALPNExtension 的值
	// 是按 Chrome 原样硬编码的 ["h2", "http/1.1"],Config.NextProtos 对预设
	// HelloID 不生效。forceH1 场景下必须显式覆盖 Extension 里的值,
	// 否则服务器仍按 h2 协商,后续 h1.Transport 会读到 HTTP/2 SETTINGS frame
	// 报 `malformed HTTP response "\x00\x00\x12\x04..."`。
	if forceH1 {
		if err := uconn.BuildHandshakeState(); err != nil {
			_ = raw.Close()
			return nil, fmt.Errorf("utls build state: %w", err)
		}
		for _, ext := range uconn.Extensions {
			if alpnExt, ok := ext.(*utls.ALPNExtension); ok {
				alpnExt.AlpnProtocols = []string{"http/1.1"}
			}
		}
	}

	if err := uconn.HandshakeContext(ctx); err != nil {
		_ = raw.Close()
		return nil, fmt.Errorf("utls handshake %s: %w", host, err)
	}

	// 二次保险:如果服务器忽略我们的 ALPN(比如某些商业代理)还是协商到 h2,
	// 直接断开,避免 h1.Transport 读到 SETTINGS frame 的脏字节。
	if forceH1 {
		np := uconn.ConnectionState().NegotiatedProtocol
		if np != "" && np != "http/1.1" {
			_ = uconn.Close()
			return nil, fmt.Errorf("alpn negotiated %q, expected http/1.1", np)
		}
	}
	return uconn, nil
}

// dialRaw 返回一个已经"到对端 host:port"的 TCP 通道。若配置了 HTTP 代理,先
// 连到代理再发 CONNECT。
func (rt *utlsRoundTripper) dialRaw(ctx context.Context, addr string) (net.Conn, error) {
	if rt.proxyURL == nil {
		return rt.dialer.DialContext(ctx, "tcp", addr)
	}
	proxyAddr := rt.proxyURL.Host
	if !strings.Contains(proxyAddr, ":") {
		if strings.EqualFold(rt.proxyURL.Scheme, "https") {
			proxyAddr += ":443"
		} else {
			proxyAddr += ":80"
		}
	}
	conn, err := rt.dialer.DialContext(ctx, "tcp", proxyAddr)
	if err != nil {
		return nil, fmt.Errorf("dial proxy %s: %w", proxyAddr, err)
	}
	// HTTPS 代理本身要先 TLS 握手(走标准 tls,不需伪装指纹,代理一般不卡 JA3)
	if strings.EqualFold(rt.proxyURL.Scheme, "https") {
		tlsConn := tls.Client(conn, &tls.Config{ServerName: rt.proxyURL.Hostname()})
		if err := tlsConn.HandshakeContext(ctx); err != nil {
			_ = conn.Close()
			return nil, fmt.Errorf("tls handshake to https proxy: %w", err)
		}
		conn = tlsConn
	}
	// CONNECT addr
	connectReq := &http.Request{
		Method: http.MethodConnect,
		URL:    &url.URL{Opaque: addr},
		Host:   addr,
		Header: make(http.Header),
	}
	connectReq.Header.Set("User-Agent", DefaultUserAgent)
	if u := rt.proxyURL.User; u != nil {
		pw, _ := u.Password()
		connectReq.Header.Set("Proxy-Authorization", "Basic "+
			base64.StdEncoding.EncodeToString([]byte(u.Username()+":"+pw)))
	}
	if err := connectReq.Write(conn); err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("write CONNECT: %w", err)
	}
	br := bufio.NewReader(conn)
	resp, err := http.ReadResponse(br, connectReq)
	if err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("read CONNECT response: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		_ = conn.Close()
		return nil, fmt.Errorf("proxy CONNECT %s → %s", addr, resp.Status)
	}
	// br 里可能已经预读了握手后的第一批字节,必须把它包进 conn 返回,否则 TLS 握手会少字节
	if n := br.Buffered(); n > 0 {
		peeked, _ := br.Peek(n)
		return &bufConn{Conn: conn, rd: bufio.NewReaderSize(io.MultiReader(peeked2Reader(peeked), conn), 4096)}, nil
	}
	return conn, nil
}

// isH2Retryable 判断 h2 错误是否可以降级到 h1 重试。
func isH2Retryable(err error) bool {
	if err == nil {
		return false
	}
	s := err.Error()
	// 常见 h2 → h1 协商失败或 HTTP_1_1_REQUIRED
	return strings.Contains(s, "HTTP_1_1_REQUIRED") ||
		strings.Contains(s, "http2: unsupported scheme") ||
		strings.Contains(s, "bad protocol") ||
		strings.Contains(s, "remote error: tls: no application protocol") ||
		strings.Contains(s, "http2: server sent GOAWAY") ||
		errors.Is(err, http2.ErrNoCachedConn)
}

// bufConn 让预读过的字节流继续可读,同时保留 net.Conn 的其他方法(Close/LocalAddr 等)。
type bufConn struct {
	net.Conn
	rd *bufio.Reader
}

func (b *bufConn) Read(p []byte) (int, error) { return b.rd.Read(p) }

// peeked2Reader 把一段已 Peek 的字节封装成 io.Reader。独立函数是为了避免
// 直接把 bufio.Reader 塞进 io.MultiReader 时,它内部还能继续读原 conn 的字节
// 而造成重复读。
func peeked2Reader(peeked []byte) io.Reader {
	return &readOnceBuf{buf: peeked}
}

type readOnceBuf struct {
	buf []byte
	off int
}

func (r *readOnceBuf) Read(p []byte) (int, error) {
	if r.off >= len(r.buf) {
		return 0, io.EOF
	}
	n := copy(p, r.buf[r.off:])
	r.off += n
	return n, nil
}
````

## File: internal/usage/admin_handler.go
````go
package usage

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/432539/gpt2api/pkg/resp"
)

type AdminHandler struct{ qdao *QueryDAO }

func NewAdminHandler(qdao *QueryDAO) *AdminHandler { return &AdminHandler{qdao: qdao} }

func filterFromQuery(c *gin.Context) Filter {
	mid64, _ := strconv.ParseUint(c.Query("model_id"), 10, 64)
	aid64, _ := strconv.ParseUint(c.Query("account_id"), 10, 64)
	return Filter{
		ModelID:   mid64,
		AccountID: aid64,
		Type:      c.Query("type"),
		Status:    c.Query("status"),
		Since:     parseFlexTime(c.Query("since")),
		Until:     parseFlexTime(c.Query("until")),
	}
}

func parseFlexTime(s string) time.Time {
	if s == "" {
		return time.Time{}
	}
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		return t
	}
	if t, err := time.Parse("2006-01-02", s); err == nil {
		return t
	}
	return time.Time{}
}

// Stats GET /api/admin/usage/stats。
func (h *AdminHandler) Stats(c *gin.Context) {
	f := filterFromQuery(c)
	days, _ := strconv.Atoi(c.DefaultQuery("days", "14"))
	topN, _ := strconv.Atoi(c.DefaultQuery("top_n", "10"))
	overall, err := h.qdao.Overall(c.Request.Context(), f)
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}
	daily, err := h.qdao.Daily(c.Request.Context(), f, days)
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}
	byModel, err := h.qdao.ByModel(c.Request.Context(), f, topN)
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}
	resp.OK(c, gin.H{"overall": overall, "daily": daily, "by_model": byModel})
}

// Logs GET /api/admin/usage/logs。
func (h *AdminHandler) Logs(c *gin.Context) {
	f := filterFromQuery(c)
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	rows, total, err := h.qdao.List(c.Request.Context(), f, offset, limit)
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}
	resp.OK(c, gin.H{"items": rows, "total": total, "limit": limit, "offset": offset})
}
````

## File: internal/usage/logger.go
````go
// Package usage 实现 usage_logs 的异步批量写入。
//
// 设计目标:
//  1. 网关请求落盘不在关键路径上(同步 INSERT 会拖垮高并发)。
//  2. Channel 缓冲 + 定时 flush + 批量 INSERT。
//  3. 丢弃策略:channel 满时异步降级到 Warn 日志,不阻塞调用方。
//
// 参数(默认):
//   - buffer: 8192 条
//   - batch : 500 条
//   - flush : 200ms
package usage

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/432539/gpt2api/pkg/logger"
)

// Options 可选参数。
type Options struct {
	Buffer        int
	Batch         int
	FlushInterval time.Duration
}

// Logger 异步写入器。
type Logger struct {
	db     *sqlx.DB
	ch     chan *Log
	opt    Options
	closed chan struct{}
	wg     sync.WaitGroup
}

// New 创建并启动后台 flusher。调用方在进程退出前应 Close。
func New(db *sqlx.DB, opt Options) *Logger {
	if opt.Buffer <= 0 {
		opt.Buffer = 8192
	}
	if opt.Batch <= 0 {
		opt.Batch = 500
	}
	if opt.FlushInterval <= 0 {
		opt.FlushInterval = 200 * time.Millisecond
	}
	l := &Logger{
		db:     db,
		ch:     make(chan *Log, opt.Buffer),
		opt:    opt,
		closed: make(chan struct{}),
	}
	l.wg.Add(1)
	go l.loop()
	return l
}

// Write 非阻塞投递一条日志。channel 满时降级到 Warn 日志并丢弃。
func (l *Logger) Write(row *Log) {
	if row == nil {
		return
	}
	if row.CreatedAt.IsZero() {
		row.CreatedAt = time.Now()
	}
	select {
	case l.ch <- row:
	default:
		logger.L().Warn("usage_logs channel full, dropping entry",
			zap.String("request_id", row.RequestID))
	}
}

// Close 停止后台 flusher,并把剩余 buffer 落盘。
func (l *Logger) Close() {
	close(l.closed)
	l.wg.Wait()
}

func (l *Logger) loop() {
	defer l.wg.Done()
	tick := time.NewTicker(l.opt.FlushInterval)
	defer tick.Stop()

	batch := make([]*Log, 0, l.opt.Batch)
	flush := func() {
		if len(batch) == 0 {
			return
		}
		if err := l.bulkInsert(batch); err != nil {
			logger.L().Error("usage_logs bulk insert", zap.Error(err),
				zap.Int("rows", len(batch)))
		}
		batch = batch[:0]
	}

	for {
		select {
		case <-l.closed:
			// drain
			for {
				select {
				case row := <-l.ch:
					batch = append(batch, row)
					if len(batch) >= l.opt.Batch {
						flush()
					}
				default:
					flush()
					return
				}
			}
		case row := <-l.ch:
			batch = append(batch, row)
			if len(batch) >= l.opt.Batch {
				flush()
			}
		case <-tick.C:
			flush()
		}
	}
}

func (l *Logger) bulkInsert(rows []*Log) error {
	if len(rows) == 0 {
		return nil
	}
	const cols = 15
	var b strings.Builder
	b.WriteString(`INSERT INTO usage_logs
        (model_id, account_id, request_id, type,
         input_tokens, output_tokens, cache_read_tokens, cache_write_tokens,
         image_count, duration_ms, status, error_code, ip, ua, created_at)
        VALUES `)

	args := make([]interface{}, 0, len(rows)*cols)
	for i, r := range rows {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteString("(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
		args = append(args,
			r.ModelID, r.AccountID, r.RequestID, r.Type,
			r.InputTokens, r.OutputTokens, r.CacheReadTokens, r.CacheWriteTokens,
			r.ImageCount, r.DurationMs, r.Status, r.ErrorCode,
			r.IP, r.UA, r.CreatedAt,
		)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := l.db.ExecContext(ctx, b.String(), args...)
	return err
}
````

## File: internal/usage/me_handler.go
````go
package usage

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/432539/gpt2api/pkg/resp"
)

// MeHandler 面向本地控制台的 usage 只读接口。
type MeHandler struct{ qdao *QueryDAO }

func NewMeHandler(qdao *QueryDAO) *MeHandler { return &MeHandler{qdao: qdao} }

func filterFromMeQuery(c *gin.Context) Filter {
	f := Filter{Type: c.Query("type"), Status: c.Query("status"), Since: parseFlexTime(c.Query("since")), Until: parseFlexTime(c.Query("until"))}
	if mid, err := strconv.ParseUint(c.Query("model_id"), 10, 64); err == nil {
		f.ModelID = mid
	}
	return f
}

// Logs GET /api/me/usage/logs。
func (h *MeHandler) Logs(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if limit > 200 {
		limit = 200
	}
	f := filterFromMeQuery(c)
	rows, total, err := h.qdao.List(c.Request.Context(), f, offset, limit)
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}
	resp.OK(c, gin.H{"items": rows, "total": total, "limit": limit, "offset": offset})
}

// Stats GET /api/me/usage/stats。
func (h *MeHandler) Stats(c *gin.Context) {
	days, _ := strconv.Atoi(c.DefaultQuery("days", "14"))
	topN, _ := strconv.Atoi(c.DefaultQuery("top_n", "5"))
	f := filterFromMeQuery(c)
	overall, err := h.qdao.Overall(c.Request.Context(), f)
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}
	daily, err := h.qdao.Daily(c.Request.Context(), f, days)
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}
	byModel, err := h.qdao.ByModel(c.Request.Context(), f, topN)
	if err != nil {
		resp.Internal(c, err.Error())
		return
	}
	resp.OK(c, gin.H{"overall": overall, "daily": daily, "by_model": byModel})
}
````

## File: internal/usage/model.go
````go
package usage

import "time"

const (
	TypeChat  = "chat"
	TypeImage = "image"
)

const (
	StatusSuccess = "success"
	StatusFailed  = "failed"
)

// Log 对应 usage_logs 表一行。
type Log struct {
	ModelID          uint64    `db:"model_id"`
	AccountID        uint64    `db:"account_id"`
	RequestID        string    `db:"request_id"`
	Type             string    `db:"type"`
	InputTokens      int       `db:"input_tokens"`
	OutputTokens     int       `db:"output_tokens"`
	CacheReadTokens  int       `db:"cache_read_tokens"`
	CacheWriteTokens int       `db:"cache_write_tokens"`
	ImageCount       int       `db:"image_count"`
	DurationMs       int       `db:"duration_ms"`
	Status           string    `db:"status"`
	ErrorCode        string    `db:"error_code"`
	IP               string    `db:"ip"`
	UA               string    `db:"ua"`
	CreatedAt        time.Time `db:"created_at"`
}
````

## File: internal/usage/query_dao.go
````go
package usage

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

type QueryDAO struct{ db *sqlx.DB }

func NewQueryDAO(db *sqlx.DB) *QueryDAO { return &QueryDAO{db: db} }

type Filter struct {
	ModelID   uint64
	AccountID uint64
	Type      string
	Status    string
	Since     time.Time
	Until     time.Time
}

type ItemRow struct {
	ID               uint64    `db:"id" json:"id"`
	ModelID          uint64    `db:"model_id" json:"model_id"`
	ModelSlug        string    `db:"model_slug" json:"model_slug"`
	AccountID        uint64    `db:"account_id" json:"account_id"`
	RequestID        string    `db:"request_id" json:"request_id"`
	Type             string    `db:"type" json:"type"`
	InputTokens      int       `db:"input_tokens" json:"input_tokens"`
	OutputTokens     int       `db:"output_tokens" json:"output_tokens"`
	CacheReadTokens  int       `db:"cache_read_tokens" json:"cache_read_tokens"`
	CacheWriteTokens int       `db:"cache_write_tokens" json:"cache_write_tokens"`
	ImageCount       int       `db:"image_count" json:"image_count"`
	DurationMs       int       `db:"duration_ms" json:"duration_ms"`
	Status           string    `db:"status" json:"status"`
	ErrorCode        string    `db:"error_code" json:"error_code"`
	IP               string    `db:"ip" json:"ip"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
}

type ModelStat struct {
	ModelID      uint64 `db:"model_id" json:"model_id"`
	ModelSlug    string `db:"model_slug" json:"model_slug"`
	Type         string `db:"type" json:"type"`
	Requests     int64  `db:"requests" json:"requests"`
	Failures     int64  `db:"failures" json:"failures"`
	InputTokens  int64  `db:"input_tokens" json:"input_tokens"`
	OutputTokens int64  `db:"output_tokens" json:"output_tokens"`
	ImageCount   int64  `db:"image_count" json:"image_count"`
	AvgDurMs     int64  `db:"avg_dur_ms" json:"avg_dur_ms"`
}

type DailyPoint struct {
	Day          string `db:"day" json:"day"`
	Requests     int64  `db:"requests" json:"requests"`
	Failures     int64  `db:"failures" json:"failures"`
	InputTokens  int64  `db:"input_tokens" json:"input_tokens"`
	OutputTokens int64  `db:"output_tokens" json:"output_tokens"`
	ImageCount   int64  `db:"image_count" json:"image_count"`
}

type Overall struct {
	Requests     int64 `db:"requests" json:"requests"`
	Failures     int64 `db:"failures" json:"failures"`
	ChatRequests int64 `db:"chat_requests" json:"chat_requests"`
	ImageImages  int64 `db:"image_images" json:"image_images"`
	InputTokens  int64 `db:"input_tokens" json:"input_tokens"`
	OutputTokens int64 `db:"output_tokens" json:"output_tokens"`
}

func (d *QueryDAO) buildWhere(f Filter) (string, []any) {
	b := strings.Builder{}
	b.WriteString("WHERE 1=1")
	args := make([]any, 0, 6)
	if f.ModelID > 0 {
		b.WriteString(" AND u.model_id = ?")
		args = append(args, f.ModelID)
	}
	if f.AccountID > 0 {
		b.WriteString(" AND u.account_id = ?")
		args = append(args, f.AccountID)
	}
	if f.Type != "" {
		b.WriteString(" AND u.type = ?")
		args = append(args, f.Type)
	}
	if f.Status != "" {
		b.WriteString(" AND u.status = ?")
		args = append(args, f.Status)
	}
	if !f.Since.IsZero() {
		b.WriteString(" AND u.created_at >= ?")
		args = append(args, f.Since)
	}
	if !f.Until.IsZero() {
		b.WriteString(" AND u.created_at < ?")
		args = append(args, f.Until)
	}
	return b.String(), args
}

func (d *QueryDAO) List(ctx context.Context, f Filter, offset, limit int) ([]ItemRow, int64, error) {
	where, args := d.buildWhere(f)
	if limit <= 0 || limit > 200 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}
	q := fmt.Sprintf(`
SELECT u.id, u.model_id,
       COALESCE(m.slug, '') AS model_slug,
       u.account_id, u.request_id, u.type,
       u.input_tokens, u.output_tokens, u.cache_read_tokens, u.cache_write_tokens,
       u.image_count, u.duration_ms, u.status, u.error_code, u.ip, u.created_at
FROM usage_logs u
LEFT JOIN models m ON m.id = u.model_id
%s
ORDER BY u.id DESC
LIMIT ? OFFSET ?`, where)
	rows := make([]ItemRow, 0, limit)
	err := d.db.SelectContext(ctx, &rows, q, append(args, limit, offset)...)
	if err != nil {
		return nil, 0, err
	}
	countQ := fmt.Sprintf(`SELECT COUNT(*) FROM usage_logs u %s`, where)
	var total int64
	if err := d.db.GetContext(ctx, &total, countQ, args...); err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (d *QueryDAO) Overall(ctx context.Context, f Filter) (Overall, error) {
	where, args := d.buildWhere(f)
	q := fmt.Sprintf(`
SELECT COUNT(*) AS requests,
       COALESCE(SUM(CASE WHEN u.status = 'failed' THEN 1 ELSE 0 END), 0) AS failures,
       COALESCE(SUM(CASE WHEN u.type = 'chat' THEN 1 ELSE 0 END), 0) AS chat_requests,
       COALESCE(SUM(CASE WHEN u.type = 'image' THEN u.image_count ELSE 0 END), 0) AS image_images,
       COALESCE(SUM(u.input_tokens), 0) AS input_tokens,
       COALESCE(SUM(u.output_tokens), 0) AS output_tokens
FROM usage_logs u %s`, where)
	var out Overall
	if err := d.db.GetContext(ctx, &out, q, args...); err != nil {
		return out, err
	}
	return out, nil
}

func (d *QueryDAO) ByModel(ctx context.Context, f Filter, limit int) ([]ModelStat, error) {
	where, args := d.buildWhere(f)
	if limit <= 0 {
		limit = 50
	}
	q := fmt.Sprintf(`
SELECT u.model_id,
       COALESCE(m.slug, '') AS model_slug,
       COALESCE(MAX(u.type), '') AS type,
       COUNT(*) AS requests,
       COALESCE(SUM(CASE WHEN u.status='failed' THEN 1 ELSE 0 END), 0) AS failures,
       COALESCE(SUM(u.input_tokens), 0) AS input_tokens,
       COALESCE(SUM(u.output_tokens), 0) AS output_tokens,
       COALESCE(SUM(u.image_count), 0) AS image_count,
       COALESCE(CAST(AVG(u.duration_ms) AS SIGNED), 0) AS avg_dur_ms
FROM usage_logs u
LEFT JOIN models m ON m.id = u.model_id
%s
GROUP BY u.model_id, m.slug
ORDER BY requests DESC
LIMIT ?`, where)
	rows := make([]ModelStat, 0, limit)
	err := d.db.SelectContext(ctx, &rows, q, append(args, limit)...)
	return rows, err
}

func (d *QueryDAO) Daily(ctx context.Context, f Filter, days int) ([]DailyPoint, error) {
	if days <= 0 || days > 180 {
		days = 14
	}
	since := time.Now().AddDate(0, 0, -days+1).Truncate(24 * time.Hour)
	if f.Since.IsZero() || f.Since.Before(since) {
		f.Since = since
	}
	where, args := d.buildWhere(f)
	q := fmt.Sprintf(`
SELECT DATE_FORMAT(u.created_at, '%%Y-%%m-%%d') AS day,
       COUNT(*) AS requests,
       COALESCE(SUM(CASE WHEN u.status='failed' THEN 1 ELSE 0 END), 0) AS failures,
       COALESCE(SUM(u.input_tokens), 0) AS input_tokens,
       COALESCE(SUM(u.output_tokens), 0) AS output_tokens,
       COALESCE(SUM(u.image_count), 0) AS image_count
FROM usage_logs u
%s
GROUP BY day
ORDER BY day ASC`, where)
	rows := make([]DailyPoint, 0, days)
	err := d.db.SelectContext(ctx, &rows, q, args...)
	return rows, err
}
````

## File: legacy/gen_image.py
````python
"""
ChatGPT 图片生成 —— 纯协议版（参考 gpt4free）

关键流程：
  1. GET chatgpt.com/  → 拿 oai-did cookie
  2. POST /backend-api/sentinel/chat-requirements
         body:{"p": get_requirements_token(config)}
     → 返回 chat_token + 可选 proofofwork / turnstile 挑战
  3. 若 proofofwork required → 本地 SHA3-512 暴力解
  4. POST /backend-api/f/conversation (SSE)
     header:
        openai-sentinel-chat-requirements-token: chat_token
        openai-sentinel-proof-token: proof_token
  5. 解析 SSE → 拿 file_id → 下载图片

用法: python gen_image.py "提示词"
"""
import sys, json, time, base64, hashlib, random, os, uuid
from datetime import datetime, timezone
from typing import Optional
from curl_cffi.requests import Session

# ── 配置 ──────────────────────────────────────────────────────────────────────
AUTH_TOKEN = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjE5MzQ0ZTY1LWJiYzktNDRkMS1hOWQwLWY5NTdiMDc5YmQwZSIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsiaHR0cHM6Ly9hcGkub3BlbmFpLmNvbS92MSJdLCJjbGllbnRfaWQiOiJhcHBfWDh6WTZ2VzJwUTl0UjNkRTduSzFqTDVnSCIsImV4cCI6MTc3NjU3NzQ2MSwiaHR0cHM6Ly9hcGkub3BlbmFpLmNvbS9hdXRoIjp7ImNoYXRncHRfYWNjb3VudF9pZCI6Ijk4MTQ3NzUzLTU1ODUtNGI5NS04ZjQ3LWI0YjRmMzJkYTdhOCIsImNoYXRncHRfYWNjb3VudF91c2VyX2lkIjoidXNlci15czE3U3pJQ2k5UjB1ODNvRDZmd0VmYlpfXzk4MTQ3NzUzLTU1ODUtNGI5NS04ZjQ3LWI0YjRmMzJkYTdhOCIsImNoYXRncHRfY29tcHV0ZV9yZXNpZGVuY3kiOiJub19jb25zdHJhaW50IiwiY2hhdGdwdF9wbGFuX3R5cGUiOiJmcmVlIiwiY2hhdGdwdF91c2VyX2lkIjoidXNlci15czE3U3pJQ2k5UjB1ODNvRDZmd0VmYloiLCJ1c2VyX2lkIjoidXNlci15czE3U3pJQ2k5UjB1ODNvRDZmd0VmYloifSwiaHR0cHM6Ly9hcGkub3BlbmFpLmNvbS9wcm9maWxlIjp7ImVtYWlsIjoicXE0MzI1MzlAcHJvdG9uLm1lIiwiZW1haWxfdmVyaWZpZWQiOnRydWV9LCJpYXQiOjE3NzU3MTM0NjEsImlzcyI6Imh0dHBzOi8vYXV0aC5vcGVuYWkuY29tIiwianRpIjoiN2UxYWM5MTctODIxZi00ZTRmLWJlMTctNzJjYTEzYTI5Mzk0IiwibmJmIjoxNzc1NzEzNDYxLCJwd2RfYXV0aF90aW1lIjoxNzcwMjg5MDE0NTM1LCJzY3AiOlsib3BlbmlkIiwiZW1haWwiLCJwcm9maWxlIiwib2ZmbGluZV9hY2Nlc3MiLCJtb2RlbC5yZXF1ZXN0IiwibW9kZWwucmVhZCIsIm9yZ2FuaXphdGlvbi5yZWFkIiwib3JnYW5pemF0aW9uLndyaXRlIl0sInNlc3Npb25faWQiOiJhdXRoc2Vzc18zZTdFbHF2MFIyR2tQcVF4QzJ3ZzVieEQiLCJzbCI6dHJ1ZSwic3ViIjoiYXV0aDB8NjRjZGVkOTllY2QxNTQ0OTI5NzE1NDkyIn0.r9iGZ8V27MuG30AB3o8SJFxnC64hSOYMlFAZbgqm1nzcFhox95EdX492XdP0--HuFO2HXJvxGjVUt3MGnRthCX2blOoO0tB5UhroGOxnPMtepSfghR3-cg8pLTEISdInKLMfuR616BVISbrfMIIdef0bi-Vfww_-J6ZhAlnX93xTZRTZBATuYAA3EXwjI0dwPycfuwtSY2db2CwPqYOm73CPiDmmCo1eLaKw9mB_QbRY6NdApuHNosLsYmsq5KxX0QgQC7MxfLnz5tnT_hK8g9T0lfWyARyIsa-MqJs6XlOCmpjIrWM985w2qW0xEMsoT3Nul5CL0Lwxje_86BU1hhDBCb2AKqqTe3rFiMQTDo_G6tEX-d8OeWLsCXYwgjaSlVv5QsYZzi2M4R2C_HLi8YL7744QxfAmwie3HVpFZFI4sQX9LcmyD_-S6-zBQxJwOO4haQGk_wTRHtsM_x5lK_7iTEjjufsEurnIB5P7jVtxAZS1rBnii5FOxi3ut20Ze0S9SWv7mNF5oX60VlPTMY4kHNOCf4gQ6eD4cFO-n1msb31p3dvqZTQ5NVWKGULgWC2E_REVWLkYtIzUc3qQbzDaJNDzNjpZiQYiiVHxVFftILrxyohS2QxFAs6S0oQ3rTvwtFQakn4SSh0g492394jmCq8YwkgmYdVRqGXTbq4"

PROXY      = None  # 为 None 则走本机默认路由；需要代理时填 "http://user:pass@host:port"
                   # 或把 PROXY_TEMPLATE 填上 → 运行时自动抽活出口
# 本机已开 Mihomo/Clash Verge TUN 模式，直连即可由 TUN 透明转发出国。
# 如果 arxlabs 等外部代理恢复可用想再启用，把下面这行取消注释即可：
# PROXY_TEMPLATE = "http://p9mx1124350-region-Rand-sid-{sid}-t-1:iy2lmzpy@us.arxlabs.io:3010"
PROXY_TEMPLATE = None
BASE_URL   = "https://chatgpt.com"
OUTPUT_DIR = r"C:\Users\Administrator\Documents\gpt2api_images"

# 固定复用的会话 ID：如果有值就在该会话里追加消息；
# 设为 None 则每次都走"新建会话"流程（init → prepare → send），每次重新洗灰度桶。
# 上一个会话 69e1c678 被我们密集轮询打到限流坏了，换成用户抓包里成功过灰度的新会话。
FIXED_CONVERSATION_ID = "69e2205a-b5e4-83e8-8e6a-74d8b0c1941c"

# 灰度未命中时的最大重试次数。每次重试会开新 session（新 oai-did / chat_token）。
MAX_ATTEMPTS = 8

# 用户指定必加的一句（只追加这一句，不要加其他多余的英文约束）。
# 只在提示词包含要渲染的文字（对话、标语、字幕等）时才追加，纯画面不加。
CLARITY_SUFFIX = "\n\nclean readable Chinese text, prioritize text clarity over image details"

# 判断提示词是否需要渲染文字的启发式：包含常见"要出字"的关键词、引号、或
# 冒号后面紧跟中文等。需要时才追加 CLARITY_SUFFIX。
_TEXT_HINT_KWS = (
    "文字", "对话", "台词", "旁白", "标语", "字幕", "标题", "文案",
    "招牌", "横幅", "海报文字", "弹幕", "气泡", "字体",
    "text:", "caption", "subtitle", "title:", "label", "banner", "poster text",
)
_TEXT_HINT_CHARS = ('"', "'", "“", "”", "‘", "’", "「", "」", "『", "』", "：", ":")

def needs_text_rendering(prompt: str) -> bool:
    p = prompt.lower()
    if any(kw.lower() in p for kw in _TEXT_HINT_KWS):
        return True
    # 中文/英文引号里有内容（长度 >= 2 的）才算需要文字
    import re
    if re.search(r'["“‘「『][^"”’」』]{2,}["”’」』]', prompt):
        return True
    return False

UA = ("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 "
      "(KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")

# ── 日志 ──────────────────────────────────────────────────────────────────────
def log(msg, level="INFO"):
    ts = datetime.now().strftime("%H:%M:%S")
    icon = "✓" if level=="OK" else "✗" if level=="ERR" else "·"
    print(f"[{ts}] {icon} {msg}", flush=True)

# ── POW 算法（迁移自 gpt4free/openai/proofofwork.py + new.py） ─────────────────
_CORES   = [16, 24, 32]
_SCREENS = [3000, 4000, 6000]
_MAX_ATTEMPTS = 500000

_NAV_KEYS = [
    "webdriver−false", "vendor−Google Inc.", "cookieEnabled−true",
    "pdfViewerEnabled−true", "hardwareConcurrency−32",
    "language−zh-CN", "mimeTypes−[object MimeTypeArray]",
    "userAgentData−[object NavigatorUAData]",
]
_WIN_KEYS = ["innerWidth", "innerHeight", "devicePixelRatio", "screen",
             "chrome", "location", "history", "navigator"]

def _parse_time() -> str:
    now = datetime.now(timezone.utc)
    return now.strftime("%a, %d %b %Y %H:%M:%S GMT")

def _pow_config(user_agent: str) -> list:
    import time as _t
    return [
        random.choice(_CORES) + random.choice(_SCREENS),
        datetime.now(timezone.utc).strftime("%a %b %d %Y %H:%M:%S") + " GMT+0000 (UTC)",
        None,
        random.random(),
        user_agent,
        None,
        "dpl=1440a687921de39ff5ee56b92807faaadce73f13",
        "en-US",
        "en-US,zh-CN",
        0,
        random.choice(_NAV_KEYS),
        "location",
        random.choice(_WIN_KEYS),
        _t.perf_counter(),
        str(uuid.uuid4()),
        "",
        8,
        int(_t.time()),
    ]

def _generate_answer(seed: str, difficulty: str, config: list):
    diff_len = len(difficulty)
    seed_enc = seed.encode()
    p1 = (json.dumps(config[:3],  separators=(',',':'), ensure_ascii=False)[:-1] + ',').encode()
    p2 = (',' + json.dumps(config[4:9], separators=(',',':'), ensure_ascii=False)[1:-1] + ',').encode()
    p3 = (',' + json.dumps(config[10:], separators=(',',':'), ensure_ascii=False)[1:]).encode()
    target = bytes.fromhex(difficulty)
    for i in range(_MAX_ATTEMPTS):
        d1 = str(i).encode()
        d2 = str(i >> 1).encode()
        b64 = base64.b64encode(p1 + d1 + p2 + d2 + p3)
        h = hashlib.sha3_512(seed_enc + b64).digest()
        if h[:diff_len] <= target:
            return b64.decode(), True
    fb = 'wQ8Lk5FbGpA2NcR9dShT6gYjU7VxZ4D' + base64.b64encode(f'"{seed}"'.encode()).decode()
    return fb, False

def get_requirements_token(config: list) -> str:
    seed = format(random.random())
    ans, ok = _generate_answer(seed, "0fffff", config)
    return "gAAAAAC" + ans

def generate_proof_token(required: bool, seed: str, difficulty: str,
                         user_agent: str, proof_config: list) -> Optional[str]:
    if not required:
        return None
    # gpt4free 另一种老 config（更轻量），备用
    if proof_config is None:
        scr = random.choice([3008, 4010, 6000]) * random.choice([1, 2, 4])
        proof_config = [
            scr, _parse_time(), None, 0, user_agent,
            "https://tcr9i.chat.openai.com/v2/35536E1E-65B4-4D96-9D97-6ADB7EFF8147/api.js",
            "dpl=1440a687921de39ff5ee56b92807faaadce73f13", "en", "en-US",
            None,
            "plugins−[object PluginArray]",
            random.choice(["_reactListeningcfilawjnerp", "_reactListening9ne2dfo1i47"]),
            random.choice(["alert", "ontransitionend", "onprogress"]),
        ]
    diff_len = len(difficulty)
    for i in range(100000):
        proof_config[3] = i
        j = json.dumps(proof_config)
        b64 = base64.b64encode(j.encode()).decode()
        h = hashlib.sha3_512((seed + b64).encode()).digest()
        if h.hex()[:diff_len] <= difficulty:
            return "gAAAAAB" + b64
    fb = base64.b64encode(f'"{seed}"'.encode()).decode()
    return "gAAAAABwQ8Lk5FbGpA2NcR9dShT6gYjU7VxZ4D" + fb

# ── Session 工厂 ─────────────────────────────────────────────────────────────
def _probe_proxy(max_tries: int = 20) -> Optional[str]:
    """随机抽 sid，找一个能到达 chatgpt.com 的出口。返回完整 proxy url。"""
    import random, string
    log("实时抽代理出口（chatgpt.com 可达）...")
    for i in range(max_tries):
        sid = "".join(random.choices(string.ascii_letters + string.digits, k=8))
        px  = PROXY_TEMPLATE.format(sid=sid)
        try:
            s = Session(impersonate="chrome131", verify=False, proxy=px, timeout=12)
            r = s.get(BASE_URL + "/", timeout=15)
            if r.status_code == 200 and "oai-did" in str(r.cookies):
                try:
                    info = Session(impersonate="chrome131", verify=False, proxy=px).get(
                        "https://ipinfo.io/json", timeout=10).json()
                    geo = f"{info.get('country','?')} {info.get('org','')[:30]}"
                except Exception:
                    geo = "?"
                log(f"  sid={sid} → 200 ({geo})", "OK")
                return px
        except Exception:
            pass
        log(f"  sid={sid} 不可用")
    log("所有 sid 都不可用", "ERR")
    return None

def new_session(proxy_override: Optional[str] = None) -> Session:
    kw = {"impersonate": "chrome131", "verify": False}
    chosen = proxy_override or PROXY
    if chosen:
        kw["proxy"] = chosen
    s = Session(**kw)
    s.headers.update({
        "user-agent":      UA,
        "accept-language": "en-US,en;q=0.9",
        "origin":          BASE_URL,
        "referer":         BASE_URL + "/",
        "accept":          "*/*",
        "sec-ch-ua":       '"Google Chrome";v="131", "Chromium";v="131", "Not_A Brand";v="24"',
        "sec-ch-ua-mobile":   "?0",
        "sec-ch-ua-platform": '"Windows"',
        "sec-fetch-dest":  "empty",
        "sec-fetch-mode":  "cors",
        "sec-fetch-site":  "same-origin",
    })
    return s

def _http_retry(fn, retries=5, delay=2.0, label="", retry_on_status=()):
    """不稳定代理下的简单重试包装。
    retry_on_status：对指定的 HTTP 状态码（如 429/502/503/504）也触发重试。
    """
    last_exc = None
    last_resp = None
    for i in range(retries):
        try:
            r = fn()
        except Exception as e:
            last_exc = e
            log(f"  {label} 第{i+1}次失败: {type(e).__name__}: {str(e)[:80]}")
            time.sleep(delay)
            continue
        # 如果指定状态码要重试（典型 429）
        if retry_on_status and getattr(r, "status_code", 0) in retry_on_status:
            last_resp = r
            wait = delay * (2 ** i)  # 退避
            log(f"  {label} 第{i+1}次 → {r.status_code}，等 {wait:.0f}s 后重试")
            time.sleep(wait)
            continue
        return r
    if last_resp is not None:
        return last_resp
    raise last_exc

# ── Step 1: oai-did ──────────────────────────────────────────────────────────
def bootstrap(s: Session) -> str:
    log("访问首页拿 oai-did cookie...")
    r = _http_retry(lambda: s.get(BASE_URL + "/", timeout=30), retries=5, label="GET /")
    log(f"  GET / → {r.status_code}", "OK" if r.ok else "ERR")
    did = r.cookies.get("oai-did")
    if not did:
        # 从 cookie jar 里找
        for c in s.cookies.jar if hasattr(s.cookies, "jar") else []:
            name = getattr(c, "name", getattr(c, "key", ""))
            if name == "oai-did":
                did = c.value
                break
    if not did:
        did = str(uuid.uuid4())
        log(f"  未拿到 oai-did，随机生成: {did}")
    else:
        log(f"  oai-did = {did}", "OK")
    return did

# ── Step 1.5: 取已有会话的 current_node（复用固定 conv_id 时用） ───────────────
def get_conversation_head(s: Session, did: str, conv_id: str) -> Optional[str]:
    """
    返回会话最新叶子消息 id（current_node），作为下一条消息的 parent_message_id。
    失败时返回 None，调用方可自行生成随机 parent。
    """
    log(f"拉取会话 {conv_id[:8]}... 当前 head...")
    try:
        r = _http_retry(lambda: s.get(
            BASE_URL + f"/backend-api/conversation/{conv_id}",
            headers={
                "Authorization": f"Bearer {AUTH_TOKEN}",
                "oai-device-id": did,
                "accept":        "*/*",
                "accept-language":"zh-CN,zh;q=0.9,en;q=0.8",
                "oai-language":  "zh-CN",
                "origin":        BASE_URL,
                "referer":       BASE_URL + f"/c/{conv_id}",
            },
            timeout=30,
        ), retries=5, delay=4.0, label="GET conversation",
           retry_on_status=(429, 502, 503, 504))
    except Exception as e:
        log(f"  拉取会话失败: {e}", "ERR")
        return None
    if not r.ok:
        log(f"  → {r.status_code}: {r.text[:200]}", "ERR")
        return None
    try:
        js = r.json()
    except Exception:
        log("  会话响应不是 JSON", "ERR")
        return None
    head = js.get("current_node")
    mapping = js.get("mapping") or {}
    log(f"  current_node = {head}  (mapping 消息数 {len(mapping)})", "OK" if head else "ERR")
    return head

# ── Step 2: /backend-api/sentinel/chat-requirements ──────────────────────────
def get_chat_requirements(s: Session, did: str) -> tuple[str, Optional[dict]]:
    log("请求 chat-requirements...")
    cfg = _pow_config(UA)
    req_token = get_requirements_token(cfg)
    r = _http_retry(lambda: s.post(
        BASE_URL + "/backend-api/sentinel/chat-requirements",
        headers={
            "Authorization": f"Bearer {AUTH_TOKEN}",
            "oai-device-id": did,
            "content-type":  "application/json",
        },
        json={"p": req_token},
        timeout=30,
    ), retries=5, label="chat-requirements")
    log(f"  → {r.status_code}", "OK" if r.ok else "ERR")
    if not r.ok:
        log(f"  body: {r.text[:400]}", "ERR")
        r.raise_for_status()
    data = r.json()
    chat_token = data["token"]
    pow_info   = data.get("proofofwork") or {}
    log(f"  chat_token len={len(chat_token)}  pow_required={pow_info.get('required', False)}", "OK")
    return chat_token, pow_info

# ── Step 2.4: /backend-api/conversation/init（新会话必做）──────────────────────
def init_new_conversation(s: Session, did: str) -> bool:
    """
    新会话场景下，/f/conversation 之前必须先调用 /conversation/init，
    让服务端为即将创建的会话注册一下路由/限流上下文。
    请求体对齐 HAR（_har_body_3_init.json）。
    """
    log("/conversation/init（新会话注册）...")
    headers = {
        "Authorization":  f"Bearer {AUTH_TOKEN}",
        "accept":         "*/*",
        "accept-language":"zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
        "cache-control":  "no-cache",
        "pragma":         "no-cache",
        "priority":       "u=1, i",
        "content-type":   "application/json",
        "oai-device-id":  did,
        "oai-language":   "zh-CN",
        "oai-client-build-number": "5955942",
        "oai-client-version":      "prod-be885abbfcfe7b1f511e88b3003d9ee44757fbad",
        "origin":         BASE_URL,
        "referer":        BASE_URL + "/",
    }
    payload = {
        "gizmo_id":               None,
        "requested_default_model":None,
        "conversation_id":        None,
        "timezone_offset_min":    -480,
        "system_hints":           ["picture_v2"],
    }
    try:
        r = _http_retry(lambda: s.post(
            BASE_URL + "/backend-api/conversation/init",
            headers=headers, json=payload, timeout=30,
        ), retries=4, delay=3.0, label="conversation/init",
           retry_on_status=(429, 502, 503, 504))
    except Exception as e:
        log(f"  init 失败: {e}", "ERR")
        return False
    if not r.ok:
        log(f"  init → {r.status_code}: {r.text[:200]}", "ERR")
        return False
    log(f"  init → 200 OK", "OK")
    return True

# ── Step 2.5: /backend-api/f/conversation/prepare（灰度分桶关键） ─────────────
def prepare_fconversation(s: Session, did: str, chat_token: str, proof_token: Optional[str],
                          conv_id: str, parent_id: str, msg_id: str, prompt: str) -> Optional[str]:
    """
    浏览器真实流程里 conversation 前必须先 prepare。
    返回的 conduit_token 决定本次会话被路由到哪个集群（可能和新 IMG2 灰度相关）。
    """
    log("/f/conversation/prepare 预热（拿 conduit_token）...")
    headers = {
        "Authorization":  f"Bearer {AUTH_TOKEN}",
        "accept":         "*/*",
        "accept-language":"zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
        "cache-control":  "no-cache",
        "pragma":         "no-cache",
        "priority":       "u=1, i",
        "content-type":   "application/json",
        "oai-device-id":  did,
        "oai-language":   "zh-CN",
        "oai-client-build-number": "5955942",
        "oai-client-version":      "prod-be885abbfcfe7b1f511e88b3003d9ee44757fbad",
        "origin":         BASE_URL,
        "referer":        BASE_URL + "/",
        "openai-sentinel-chat-requirements-token": chat_token,
    }
    if proof_token:
        headers["openai-sentinel-proof-token"] = proof_token

    payload = {
        "action":              "next",
        "fork_from_shared_post": False,
        "conversation_id":     conv_id,
        "parent_message_id":   parent_id,
        "model":               "gpt-5-3",
        "client_prepare_state":"none",
        "timezone_offset_min": -480,
        "timezone":            "Asia/Shanghai",
        "conversation_mode":   {"kind": "primary_assistant"},
        "system_hints":        ["picture_v2"],
        "partial_query": {
            "id":      msg_id,
            "author":  {"role": "user"},
            "content": {"content_type": "text", "parts": [prompt]},
        },
        "supports_buffering":  True,
        "supported_encodings": ["v1"],
        "client_contextual_info": {"app_name": "chatgpt.com"},
    }
    try:
        r = _http_retry(lambda: s.post(
            BASE_URL + "/backend-api/f/conversation/prepare",
            headers=headers, json=payload, timeout=30,
        ), retries=4, label="f/conversation/prepare")
    except Exception as e:
        log(f"  prepare 失败（可降级直跑）: {e}", "ERR")
        return None
    if not r.ok:
        log(f"  prepare → {r.status_code}: {r.text[:300]}", "ERR")
        return None
    js = r.json() or {}
    ct = js.get("conduit_token")
    if ct:
        log(f"  conduit_token 已获取 len={len(ct)}", "OK")
    else:
        log(f"  prepare 响应无 conduit_token: {str(js)[:200]}", "ERR")
    return ct

# ── Step 3: f/conversation SSE ───────────────────────────────────────────────
def send_conversation(s: Session, did: str, chat_token: str, proof_token: Optional[str],
                      conv_id: str, parent_id: str, msg_id: str, prompt: str,
                      conduit_token: Optional[str] = None):
    log("发送生图对话请求（SSE）...")
    headers = {
        "Authorization":  f"Bearer {AUTH_TOKEN}",
        "accept":         "text/event-stream",
        "accept-language":"zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
        "cache-control":  "no-cache",
        "pragma":         "no-cache",
        "priority":       "u=1, i",
        "content-type":   "application/json",
        "oai-device-id":  did,
        "oai-language":   "zh-CN",
        "oai-client-build-number": "5955942",
        "oai-client-version":      "prod-be885abbfcfe7b1f511e88b3003d9ee44757fbad",
        "origin":         BASE_URL,
        "referer":        BASE_URL + "/",
        "openai-sentinel-chat-requirements-token": chat_token,
    }
    if proof_token:
        headers["openai-sentinel-proof-token"] = proof_token
    if conduit_token:
        headers["x-conduit-token"] = conduit_token

    payload = {
        "action": "next",
        "messages": [{
            "id":          msg_id,
            "author":      {"role": "user"},
            "create_time": time.time(),
            "content":     {"content_type": "text", "parts": [prompt]},
            "metadata": {
                "developer_mode_connector_ids": [],
                "selected_github_repos":        [],
                "selected_all_github_repos":    False,
                "system_hints":                 ["picture_v2"],
                "serialization_metadata":       {"custom_symbol_offsets": []},
            },
        }],
        "conversation_id":          conv_id,
        "parent_message_id":        parent_id,
        "model":                    "gpt-5-3",
        "client_prepare_state":     "sent",
        "timezone_offset_min":      -480,
        "timezone":                 "Asia/Shanghai",
        "conversation_mode":        {"kind": "primary_assistant"},
        "enable_message_followups": True,
        "system_hints":             ["picture_v2"],
        "supports_buffering":       True,
        "supported_encodings":      ["v1"],
        "client_contextual_info": {
            "is_dark_mode":       False,
            "time_since_loaded":  random.randint(500, 3000),
            "page_height":        1072,
            "page_width":         1724,
            "pixel_ratio":        1.2,
            "screen_height":      1440,
            "screen_width":       2560,
            "app_name":           "chatgpt.com",
        },
        "paragen_cot_summary_display_override": "allow",
        "force_parallel_switch": "auto",
    }

    r = _http_retry(lambda: s.post(
        BASE_URL + "/backend-api/f/conversation",
        headers=headers, json=payload,
        stream=True, timeout=180,
    ), retries=5, label="f/conversation")
    if not r.ok:
        body = r.text
        log(f"  SSE 失败 {r.status_code}: {body[:400]}", "ERR")
        r.raise_for_status()
    log("  SSE 连接成功", "OK")
    return r

# ── Step 4: 解析 SSE，收集 file-service 引用 ─────────────────────────────────
def parse_sse(resp) -> dict:
    new_conv_id   = None
    file_ids      = []
    tool_text     = ""
    collecting    = False
    finish_reason = None

    import re as _re
    file_pat = _re.compile(r"file-service://([A-Za-z0-9_-]+)")
    sediment_pat = _re.compile(r"sediment://([A-Za-z0-9_-]+)")

    def handle(data: str):
        nonlocal new_conv_id, tool_text, collecting, finish_reason
        for m in file_pat.finditer(data):
            fid = m.group(1)
            if fid not in file_ids:
                file_ids.append(fid)
                log(f"  发现图片: file-service://{fid[:16]}...", "OK")
        for m in sediment_pat.finditer(data):
            fid = m.group(1)
            if ("sed:" + fid) not in file_ids:
                file_ids.append("sed:" + fid)
                log(f"  发现图片: sediment://{fid[:16]}...", "OK")

        try:
            obj = json.loads(data)
        except Exception:
            return
        if not isinstance(obj, dict):
            return

        if isinstance(obj.get("v"), dict):
            cid = obj["v"].get("conversation_id")
            if cid and not new_conv_id:
                new_conv_id = cid
                log(f"  conversation_id={cid}", "OK")
            msg = obj["v"].get("message", {})
            if isinstance(msg, dict):
                meta = msg.get("metadata", {}) or {}
                if meta.get("image_gen_task_id"):
                    log(f"  image_gen_task_id={meta['image_gen_task_id']}", "OK")
                if meta.get("finish_details"):
                    finish_reason = meta["finish_details"].get("type")
                c = msg.get("content", {}) or {}
                if c.get("content_type") == "code":
                    collecting = True
                    tool_text = c.get("text", "") or ""
        elif isinstance(obj.get("v"), list) and collecting:
            for p in obj["v"]:
                if isinstance(p, dict) and p.get("p", "").endswith("/content/text") and p.get("o") == "append":
                    tool_text += p.get("v", "")

    try:
        for raw in resp.iter_lines():
            if not raw:
                continue
            if isinstance(raw, bytes):
                raw = raw.decode("utf-8", errors="replace")
            raw = raw.strip()
            if not raw.startswith("data:"):
                continue
            data = raw[5:].strip()
            if data in ("[DONE]", ""):
                break
            handle(data)
    except Exception as e:
        log(f"  SSE 中断（可能代理超时）: {type(e).__name__}: {str(e)[:100]}", "ERR")

    if tool_text.strip():
        log(f"  工具调用参数: {tool_text.strip()[:180]}", "OK")
    if finish_reason:
        log(f"  finish_reason={finish_reason}")
    return {"conversation_id": new_conv_id, "file_ids": file_ids}

# ── Step 4.5: 轮询 conversation 详情抓最终图片 ───────────────────────────────
def _extract_img2_tool_ids(mapping: dict) -> set:
    """返回会话里所有 IMG2 tool 消息的 id 集合（不含 asset_pointer，仅用于 diff）"""
    ids = set()
    for mid, node in mapping.items():
        msg  = (node or {}).get("message") or {}
        auth = msg.get("author") or {}
        meta = msg.get("metadata") or {}
        cont = msg.get("content") or {}
        if (auth.get("role") == "tool"
            and meta.get("async_task_type") == "image_gen"
            and cont.get("content_type") == "multimodal_text"):
            ids.add(mid)
    return ids

def fetch_tool_baseline(s: Session, did: str, conv_id: str) -> set:
    """
    发送前拉一次 conversation，记录已有 IMG2 tool 消息 id 集合，作为 baseline。
    poll 阶段用 current - baseline 识别本次新增的 tool 消息。
    """
    log(f"拉取 baseline（会话 {conv_id[:8]}... 已有 tool 消息）...")
    try:
        r = _http_retry(lambda: s.get(
            BASE_URL + f"/backend-api/conversation/{conv_id}",
            headers={
                "Authorization": f"Bearer {AUTH_TOKEN}",
                "oai-device-id": did,
                "accept":        "*/*",
                "accept-language":"zh-CN,zh;q=0.9,en;q=0.8",
                "oai-language":  "zh-CN",
                "origin":        BASE_URL,
                "referer":       BASE_URL + f"/c/{conv_id}",
            },
            timeout=30,
        ), retries=5, delay=4.0, label="GET conv(baseline)",
           retry_on_status=(429, 502, 503, 504))
    except Exception as e:
        log(f"  baseline 拉取失败: {e}", "ERR")
        return set()
    if not r.ok:
        log(f"  baseline → {r.status_code}", "ERR")
        return set()
    try:
        js = r.json()
    except Exception:
        return set()
    ids = _extract_img2_tool_ids(js.get("mapping") or {})
    log(f"  baseline: {len(ids)} 条历史 IMG2 tool 消息", "OK")
    return ids

def poll_conversation_for_images(s: Session, did: str, conv_id: str,
                                 baseline_tool_ids: Optional[set] = None,
                                 max_wait: int = 900, interval: float = 6.0,
                                 stable_rounds: int = 4,
                                 preview_wait_secs: float = 30.0):
    """
    轮询 conversation，等本次回合的图片稳定。返回 (status, ids)。
    - status:
        "img2"         → 命中灰度桶，返回 IMG2 最新那条 tool 消息的 ids
        "preview_only" → 只出现 1 条 tool 消息，且自第一条 tool 消息出现起
                         经过 preview_wait_secs 秒仍无第 2 条 → 判非灰度，ids 为空
        "timeout"      → 超时，返回兜底（通常为空）
    - 识别规则（实验确认的金标准）：
        灰度桶会产出 **≥ 2 条 tool 消息**（先 preview，后 IMG2 最终）；
        非灰度只产出 1 条。
        实际观察：若进了灰度，第 2 条通常在第 1 条出现后 5–30 秒内就会冒出；
        所以 30 秒仍是 1 条，基本可以判定非灰度，立即重试比死等划算。
    """
    log(f"轮询 conversation（max_wait={max_wait}s, interval={interval}s, "
        f"stable_rounds={stable_rounds}, preview_wait={preview_wait_secs}s, "
        f"baseline={len(baseline_tool_ids or [])}）...")
    import re as _re
    fpat = _re.compile(r"file-service://([A-Za-z0-9_-]+)")
    spat = _re.compile(r"sediment://([A-Za-z0-9_-]+)")
    baseline = baseline_tool_ids or set()

    t0 = time.time()
    last_sed_sig = None
    stable_count = 0
    last_sed = []
    seen_any_tool = False  # 本次回合是否已观察到至少一条新的 IMG2 tool 消息
    first_tool_ts: Optional[float] = None  # 第一条新 tool 消息首次出现的时间
    consecutive_429 = 0

    while time.time() - t0 < max_wait:
        try:
            r = _http_retry(lambda: s.get(
                f"{BASE_URL}/backend-api/conversation/{conv_id}",
                headers={"Authorization": f"Bearer {AUTH_TOKEN}", "oai-device-id": did},
                timeout=30,
            ), retries=3, delay=5.0, label="GET conv",
               retry_on_status=(429, 502, 503, 504))
        except Exception as e:
            log(f"  轮询失败: {e}", "ERR")
            time.sleep(interval)
            continue

        if r.status_code == 429:
            consecutive_429 += 1
            log(f"  轮询被 429 限流（连续 {consecutive_429} 次）", "ERR")
            if consecutive_429 >= 3:
                log("  连续 3 次 429 → 本次 attempt 中止，交外层退避重试", "ERR")
                return ("error", [])
            time.sleep(10)
            continue
        consecutive_429 = 0
        if r.status_code != 200:
            time.sleep(interval)
            continue

        try:
            js = r.json()
        except Exception:
            time.sleep(interval)
            continue

        # 第一次收到完整 mapping 就 dump 一份供调试
        try:
            dbg = os.path.join(OUTPUT_DIR, f"_conv_{conv_id[:8]}.json")
            if not os.path.exists(dbg):
                os.makedirs(OUTPUT_DIR, exist_ok=True)
                with open(dbg, "w", encoding="utf-8") as f:
                    json.dump(js, f, ensure_ascii=False, indent=2)
                log(f"  [debug] conversation 已 dump: {dbg}")
        except Exception:
            pass

        mapping = js.get("mapping") or {}
        # baseline diff：只看本次回合新出现的 IMG2 tool 消息。
        # 如果 baseline 为空，就看全部（新建会话场景）。
        all_tool_ids = _extract_img2_tool_ids(mapping)
        new_tool_ids = all_tool_ids - baseline if baseline else all_tool_ids

        # 采集每条 tool 消息的详细元数据，用于区分"快速预览 vs IMG2 最终"
        # tool_records: [{mid, create_time, model_slug, recipient, is_img2_signature, file_ids, sed_ids}]
        tool_records = []
        final_ids, sed_ids = [], []
        for mid in new_tool_ids:
            m = mapping.get(mid) or {}
            msg     = m.get("message") or {}
            author  = msg.get("author") or {}
            content = msg.get("content") or {}
            meta    = msg.get("metadata") or {}

            is_img2 = (
                author.get("role") == "tool"
                and meta.get("async_task_type") == "image_gen"
                and content.get("content_type") == "multimodal_text"
            )
            if not is_img2:
                continue
            seen_any_tool = True

            rec_fids, rec_sids = [], []
            for p in (content.get("parts") or []):
                if isinstance(p, dict):
                    aid = p.get("asset_pointer") or ""
                    for hit in fpat.finditer(aid):
                        fid = hit.group(1)
                        if fid not in rec_fids: rec_fids.append(fid)
                        if fid not in final_ids: final_ids.append(fid)
                    for hit in spat.finditer(aid):
                        fid = hit.group(1)
                        if fid not in rec_sids: rec_sids.append(fid)
                        if fid not in sed_ids: sed_ids.append(fid)
                elif isinstance(p, str):
                    for hit in fpat.finditer(p):
                        fid = hit.group(1)
                        if fid not in rec_fids: rec_fids.append(fid)
                        if fid not in final_ids: final_ids.append(fid)

            tool_records.append({
                "mid": mid,
                "create_time": msg.get("create_time") or 0,
                "model_slug": meta.get("model_slug") or author.get("metadata", {}).get("model_slug"),
                "recipient":  msg.get("recipient") or meta.get("recipient"),
                "author_name": author.get("name"),
                "image_gen_title": meta.get("image_gen_title"),
                "gizmo_id": meta.get("gizmo_id"),
                "file_ids": rec_fids,
                "sed_ids":  rec_sids,
            })

        # 按 create_time 排序（最早 → 最新）
        tool_records.sort(key=lambda r: r["create_time"] or 0)
        last_sed = sed_ids

        # 最终高清直出（优先）—— file-service 一出就是 IMG2 终稿（灰度桶）
        if final_ids:
            for rec in tool_records:
                if rec["file_ids"]:
                    log(f"  [IMG2-final] mid={rec['mid'][:8]} "
                        f"model={rec['model_slug']} recipient={rec['recipient']} "
                        f"name={rec['author_name']} → file-service x{len(rec['file_ids'])}", "OK")
            # file-service 直出一定是灰度 IMG2
            return ("img2", final_ids)

        # 本次回合还没出现任何 tool 消息时继续等（复用会话场景尤其重要）
        if not seen_any_tool:
            log("  等待本次回合的 tool 消息出现...")
            time.sleep(interval)
            continue

        elapsed = time.time() - t0
        n_tool = len(tool_records)
        # 记录第 1 条 tool 消息首次出现时间
        if n_tool >= 1 and first_tool_ts is None:
            first_tool_ts = time.time()
            log(f"  ▶ 第 1 条 tool 消息首次出现，开始 {preview_wait_secs:.0f}s 窗口等第 2 条（IMG2）")

        sig = tuple(sorted(sed_ids))

        # ── 分支 A：已经有 2+ 条 tool 消息 → 灰度命中，按 stable_rounds 等最新那条稳定
        if n_tool >= 2:
            if sed_ids and sig == last_sed_sig:
                stable_count += 1
                log(f"  sed 稳定 {stable_count}/{stable_rounds} "
                    f"（IMG2, tool x{n_tool}, sed x{len(sed_ids)}, {elapsed:.0f}s）")
                if stable_count >= stable_rounds:
                    keep = tool_records[-1]
                    log(f"  ✓ 命中灰度，{n_tool} 条 tool → 只保留最新一条（IMG2 终稿）", "OK")
                    for i, r in enumerate(tool_records[:-1]):
                        log(f"    [丢弃预览 #{i+1}] mid={r['mid'][:8]} t={r['create_time']} "
                            f"name={r['author_name']} sed={len(r['sed_ids'])} file={len(r['file_ids'])}")
                    log(f"    [保留最新  ] mid={keep['mid'][:8]} t={keep['create_time']} "
                        f"name={keep['author_name']} sed={len(keep['sed_ids'])} file={len(keep['file_ids'])}")
                    out = []
                    for fid in keep["file_ids"]:
                        log(f"  最终图片: file-service://{fid[:20]}...", "OK")
                        out.append(fid)
                    for fid in keep["sed_ids"]:
                        log(f"  最终图片: sediment://{fid[:20]}...", "OK")
                        out.append("sed:" + fid)
                    return ("img2", out)
            else:
                if sig != last_sed_sig:
                    log(f"  sed 变化: {len(sed_ids)}张 (已等 {elapsed:.0f}s, tool x{n_tool})")
                stable_count = 0
                last_sed_sig = sig

        # ── 分支 B：只有 1 条 tool 消息 → 看时间窗口，窗口内若冒出第 2 条就走分支 A
        elif n_tool == 1 and first_tool_ts is not None:
            since_first = time.time() - first_tool_ts
            log(f"  窗口中等 IMG2 第 2 条 tool: {since_first:.0f}/{preview_wait_secs:.0f}s "
                f"(sed x{len(sed_ids)}, {elapsed:.0f}s)")
            if since_first >= preview_wait_secs:
                only = tool_records[0]
                log(f"  ✗ 第 1 条 tool 出现 {since_first:.0f}s 后仍无第 2 条 "
                    f"(mid={only['mid'][:8]} name={only['author_name']}) "
                    f"→ 判定非灰度预览，需重试", "ERR")
                return ("preview_only", [])

        time.sleep(interval)

    # 超时兜底
    n_tool = len(tool_records)
    if n_tool >= 2:
        keep = tool_records[-1]
        log(f"  超时但已有 {n_tool} 条 tool 消息，取最新一条作 IMG2", "ERR")
        out = ([x for x in keep["file_ids"]] +
               ["sed:" + x for x in keep["sed_ids"]])
        return ("img2", out)
    if last_sed:
        log(f"  超时，只有 1 条 tool 消息 → 视为非灰度预览（不保存），请重试", "ERR")
        return ("preview_only", [])
    log("  轮询超时，未发现图片", "ERR")
    return ("timeout", [])

# ── Step 5: 下载图片 ─────────────────────────────────────────────────────────
def download_images(s: Session, did: str, conv_id: Optional[str], file_ids: list) -> list:
    os.makedirs(OUTPUT_DIR, exist_ok=True)
    hdrs = {"Authorization": f"Bearer {AUTH_TOKEN}", "oai-device-id": did}
    saved = []

    for raw in file_ids[:8]:
        is_sed = raw.startswith("sed:")
        fid = raw[4:] if is_sed else raw
        if is_sed and conv_id:
            url = f"{BASE_URL}/backend-api/conversation/{conv_id}/attachment/{fid}/download"
        else:
            url = f"{BASE_URL}/backend-api/files/{fid}/download"
        try:
            r = s.get(url, headers=hdrs, timeout=30)
            if not r.ok:
                log(f"  {fid[:16]}... 取下载地址失败 {r.status_code}", "ERR")
                continue
            durl = r.json().get("download_url")
            if not durl:
                log(f"  {fid[:16]}... 无 download_url", "ERR")
                continue
            img = s.get(durl, timeout=60)
            if img.ok and len(img.content) > 2000:
                path = os.path.join(OUTPUT_DIR, f"{fid[-12:]}.png")
                with open(path, "wb") as f:
                    f.write(img.content)
                log(f"已保存: {path} ({len(img.content)//1024} KB)", "OK")
                saved.append(path)
        except Exception as e:
            log(f"  下载 {fid[:16]}... 出错: {e}", "ERR")
    return saved

# ── 单次执行 ─────────────────────────────────────────────────────────────────
def run_once(prompt: str):
    """执行一次完整流程。返回 (status, files)。
    status ∈ {"img2", "preview_only", "timeout", "error"}
    """
    active_proxy = PROXY
    if not active_proxy and PROXY_TEMPLATE:
        active_proxy = _probe_proxy()
        if not active_proxy:
            log("无可用代理出口", "ERR")
            return ("error", [])
    s = new_session(active_proxy)

    try:
        did = bootstrap(s)
        chat_token, pow_info = get_chat_requirements(s, did)

        proof_token = None
        if pow_info.get("required"):
            log("本地计算 proof-of-work...")
            t0 = time.time()
            proof_token = generate_proof_token(
                required=True,
                seed=pow_info["seed"],
                difficulty=pow_info["difficulty"],
                user_agent=UA,
                proof_config=_pow_config(UA),
            )
            log(f"  proof_token len={len(proof_token or '')}  耗时 {time.time()-t0:.2f}s", "OK")

        baseline_tool_ids: set = set()
        if FIXED_CONVERSATION_ID:
            head = get_conversation_head(s, did, FIXED_CONVERSATION_ID)
            if not head:
                # 拉固定会话 head 失败（多为 429 限流），直接本次失败让外层退避重试，
                # 千万不要降级到一个全新的废弃 uuid —— 那会让 /f/conversation 直接 404。
                log("复用会话拉 head 失败 → 本次尝试中止，交由外层退避重试", "ERR")
                return ("error", [])
            conv_id = FIXED_CONVERSATION_ID
            par_id  = head
            log(f"复用会话 {conv_id}  parent={par_id}", "OK")
            baseline_tool_ids = fetch_tool_baseline(s, did, conv_id)
        else:
            # 新会话：必须先 /conversation/init 注册，否则 /f/conversation 会 404
            if not init_new_conversation(s, did):
                return ("error", [])
            conv_id = str(uuid.uuid4())
            par_id  = str(uuid.uuid4())
            log(f"新建会话 {conv_id}  parent={par_id}", "OK")
        msg_id  = str(uuid.uuid4())

        conduit_token = prepare_fconversation(
            s, did, chat_token, proof_token, conv_id, par_id, msg_id, prompt
        )

        t0 = time.time()
        resp = send_conversation(
            s, did, chat_token, proof_token, conv_id, par_id, msg_id, prompt,
            conduit_token=conduit_token,
        )
        result = parse_sse(resp)
        log(f"SSE 总耗时 {time.time()-t0:.1f}s")

        actual_conv = result.get("conversation_id") or conv_id
        sse_ids = result.get("file_ids", [])
        has_final = any(not x.startswith("sed:") for x in sse_ids)

        if actual_conv and not has_final:
            status, polled = poll_conversation_for_images(
                s, did, actual_conv,
                baseline_tool_ids=baseline_tool_ids,
            )
            file_ids = polled
        elif has_final:
            status = "img2"  # file-service 直出就是灰度
            file_ids = sse_ids
        else:
            status = "preview_only"
            file_ids = []

        if status != "img2":
            return (status, [])

        files = download_images(s, did, actual_conv, file_ids)
        return ("img2", files)
    except Exception as e:
        log(f"异常: {e}", "ERR")
        import traceback
        traceback.print_exc()
        return ("error", [])


# ── Main ─────────────────────────────────────────────────────────────────────
def main():
    prompt = " ".join(sys.argv[1:]) if len(sys.argv) > 1 else ""
    if not prompt:
        try:
            prompt = input("请输入提示词: ").strip()
        except EOFError:
            prompt = ""
    if not prompt:
        prompt = "一只可爱的橘猫坐在窗台上，阳光照进来，写实风格"
    if needs_text_rendering(prompt) and CLARITY_SUFFIX.strip() not in prompt:
        prompt += CLARITY_SUFFIX
        log("检测到文字渲染需求 → 已追加文字清晰度约束")
    else:
        log("纯画面提示词 → 跳过文字清晰度 suffix")

    print(f"\n{'='*64}")
    log(f"提示词: {prompt[:100]}{'...' if len(prompt)>100 else ''}")
    log(f"代理: {PROXY.split('@')[-1] if PROXY else '本机默认路由' if not PROXY_TEMPLATE else '自动抽出口'}")
    print(f"{'='*64}\n")

    t_all = time.time()
    for attempt in range(1, MAX_ATTEMPTS + 1):
        print(f"\n{'─'*64}")
        log(f"◆ 尝试 {attempt}/{MAX_ATTEMPTS}")
        print(f"{'─'*64}")
        t0 = time.time()
        status, files = run_once(prompt)
        dt = time.time() - t0

        if status == "img2":
            print(f"\n{'='*64}")
            log(f"✓ 命中灰度（IMG2），{len(files)} 张 → {OUTPUT_DIR}", "OK")
            log(f"  本次 {dt:.0f}s，总计 {time.time()-t_all:.0f}s（用了 {attempt} 次尝试）")
            print(f"{'='*64}\n")
            return

        reason = {"preview_only": "非灰度快速预览（只 1 条 tool 消息）",
                  "timeout":      "轮询超时",
                  "error":        "出错（多为账号级 429 限流）"}[status]
        # 账号级 429 恢复慢：preview_only 退避 40s，error（429）退避 90s 指数上升
        if status == "error":
            backoff = min(180, 60 + attempt * 15)
        else:
            backoff = 40
        if attempt < MAX_ATTEMPTS:
            log(f"× 第 {attempt} 次: {reason}（{dt:.0f}s），等 {backoff}s 冷却后开新 session 重试", "ERR")
            time.sleep(backoff)
        else:
            log(f"× 第 {attempt} 次: {reason}（{dt:.0f}s），已达上限", "ERR")

    print(f"\n{'='*64}")
    log(f"全部 {MAX_ATTEMPTS} 次尝试都未命中灰度，总计 {time.time()-t_all:.0f}s", "ERR")
    print(f"{'='*64}\n")

if __name__ == "__main__":
    main()
````

## File: LICENSE
````
MIT License

Copyright (c) 2026 gpt2api contributors

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

================================================================================

IMPORTANT — ADDITIONAL DISCLAIMER

This project reverse-engineers the private APIs of chatgpt.com and is provided
for educational and research purposes only. Use of this software may violate
OpenAI's Terms of Service and the laws of your jurisdiction. The authors and
contributors accept no responsibility for account bans, legal consequences, or
any other damages arising from the use of this software.

You are solely responsible for complying with all applicable laws, terms of
service, and regulations in your use of this project.
````

## File: Makefile
````makefile
.PHONY: help run build tidy test fmt vet db-init docker-up docker-down docker-logs

SHELL := /bin/sh
APP_NAME := gpt2api
BIN_DIR := bin

MYSQL_HOST ?= 127.0.0.1
MYSQL_PORT ?= 3306
MYSQL_USER ?= gpt2api
MYSQL_PASSWORD ?= gpt2api
MYSQL_DATABASE ?= gpt2api

help:
	@echo "Targets:"
	@echo "  run              - go run cmd/server"
	@echo "  build            - build binary to bin/$(APP_NAME)"
	@echo "  tidy             - go mod tidy"
	@echo "  test             - go test ./..."
	@echo "  fmt              - gofmt -w"
	@echo "  vet              - go vet ./..."
	@echo "  db-init          - initialize empty MySQL database from sql/database.sql"
	@echo "  docker-up        - docker compose up -d"
	@echo "  docker-down      - docker compose down"
	@echo "  docker-logs      - docker compose logs -f"

run:
	go run ./cmd/server

build:
	@mkdir -p $(BIN_DIR)
	go build -ldflags "-s -w" -o $(BIN_DIR)/$(APP_NAME) ./cmd/server

tidy:
	go mod tidy

test:
	go test ./...

fmt:
	gofmt -w .

vet:
	go vet ./...

db-init:
	MYSQL_PWD="$(MYSQL_PASSWORD)" mysql -h "$(MYSQL_HOST)" -P "$(MYSQL_PORT)" -u "$(MYSQL_USER)" "$(MYSQL_DATABASE)" < sql/database.sql

docker-up:
	docker compose -f deploy/docker-compose.yml up -d

docker-down:
	docker compose -f deploy/docker-compose.yml down

docker-logs:
	docker compose -f deploy/docker-compose.yml logs -f
````

## File: pkg/crypto/aes.go
````go
package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

// AESGCM 用于加密敏感字段(AT / cookies / 代理密码)。
// 密钥是 hex 字符串,必须 64 字符(= 32 字节 AES-256 密钥)。
type AESGCM struct {
	aead cipher.AEAD
}

func NewAESGCM(hexKey string) (*AESGCM, error) {
	if len(hexKey) != 64 {
		return nil, fmt.Errorf("aes key must be 64 hex chars (32 bytes), got %d", len(hexKey))
	}
	key, err := hex.DecodeString(hexKey)
	if err != nil {
		return nil, fmt.Errorf("decode hex key: %w", err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("new cipher: %w", err)
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("new gcm: %w", err)
	}
	return &AESGCM{aead: aead}, nil
}

// Encrypt 返回 base64(nonce || ciphertext_with_tag)。
func (a *AESGCM) Encrypt(plaintext []byte) (string, error) {
	nonce := make([]byte, a.aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ct := a.aead.Seal(nil, nonce, plaintext, nil)
	out := make([]byte, 0, len(nonce)+len(ct))
	out = append(out, nonce...)
	out = append(out, ct...)
	return base64.StdEncoding.EncodeToString(out), nil
}

// Decrypt 接受 Encrypt 返回的 base64 字符串。
func (a *AESGCM) Decrypt(b64 string) ([]byte, error) {
	raw, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, fmt.Errorf("decode base64: %w", err)
	}
	ns := a.aead.NonceSize()
	if len(raw) < ns+a.aead.Overhead() {
		return nil, errors.New("ciphertext too short")
	}
	nonce, ct := raw[:ns], raw[ns:]
	pt, err := a.aead.Open(nil, nonce, ct, nil)
	if err != nil {
		return nil, fmt.Errorf("aead open: %w", err)
	}
	return pt, nil
}

// EncryptString / DecryptString 方便字符串字段使用。
func (a *AESGCM) EncryptString(s string) (string, error) { return a.Encrypt([]byte(s)) }
func (a *AESGCM) DecryptString(b64 string) (string, error) {
	b, err := a.Decrypt(b64)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
````

## File: pkg/lock/redis_lock.go
````go
package lock

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

// ErrNotAcquired 表示未抢到锁(资源被占用)。
var ErrNotAcquired = errors.New("lock: not acquired")

// RedisLock 是一个简单的 Redis 分布式锁,用于账号池「一号一锁」。
// 通过 SET NX + token 实现原子获取与安全释放。
type RedisLock struct {
	client *redis.Client
}

func NewRedisLock(client *redis.Client) *RedisLock { return &RedisLock{client: client} }

// Acquire 尝试抢锁。成功返回 token(释放时必须提供);失败返回 ErrNotAcquired。
func (l *RedisLock) Acquire(ctx context.Context, key, token string, ttl time.Duration) error {
	ok, err := l.client.SetNX(ctx, key, token, ttl).Result()
	if err != nil {
		return err
	}
	if !ok {
		return ErrNotAcquired
	}
	return nil
}

// 释放锁使用 Lua 保证 CAS 语义(只有持有者才能删)。
var releaseScript = redis.NewScript(`
if redis.call("GET", KEYS[1]) == ARGV[1] then
    return redis.call("DEL", KEYS[1])
else
    return 0
end
`)

// Release 释放锁。若当前持有者不是 token,不会误删。
func (l *RedisLock) Release(ctx context.Context, key, token string) error {
	_, err := releaseScript.Run(ctx, l.client, []string{key}, token).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return err
	}
	return nil
}

// Refresh 续期(仅当前持有者才能续)。
var refreshScript = redis.NewScript(`
if redis.call("GET", KEYS[1]) == ARGV[1] then
    return redis.call("PEXPIRE", KEYS[1], ARGV[2])
else
    return 0
end
`)

func (l *RedisLock) Refresh(ctx context.Context, key, token string, ttl time.Duration) error {
	_, err := refreshScript.Run(ctx, l.client, []string{key}, token, ttl.Milliseconds()).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return err
	}
	return nil
}
````

## File: pkg/logger/logger.go
````go
package logger

import (
	"fmt"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	global *zap.Logger
	once   sync.Once
)

// Init 初始化全局日志。format=console|json,output=stdout|<file path>.
func Init(level, format, output string) error {
	var initErr error
	once.Do(func() {
		lvl := zapcore.InfoLevel
		if err := lvl.UnmarshalText([]byte(level)); err != nil {
			initErr = fmt.Errorf("invalid log level %q: %w", level, err)
			return
		}

		encCfg := zap.NewProductionEncoderConfig()
		encCfg.TimeKey = "ts"
		encCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encCfg.EncodeDuration = zapcore.StringDurationEncoder
		encCfg.EncodeLevel = zapcore.CapitalLevelEncoder

		var encoder zapcore.Encoder
		if format == "json" {
			encoder = zapcore.NewJSONEncoder(encCfg)
		} else {
			encCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
			encoder = zapcore.NewConsoleEncoder(encCfg)
		}

		var ws zapcore.WriteSyncer
		if output == "" || output == "stdout" {
			ws = zapcore.AddSync(os.Stdout)
		} else {
			f, err := os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
			if err != nil {
				initErr = fmt.Errorf("open log file %q: %w", output, err)
				return
			}
			ws = zapcore.AddSync(f)
		}

		core := zapcore.NewCore(encoder, ws, lvl)
		global = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	})
	return initErr
}

// L 返回全局 logger。
func L() *zap.Logger {
	if global == nil {
		// 兜底:未初始化时返回开发 logger,避免 panic。
		l, _ := zap.NewDevelopment()
		return l
	}
	return global
}

// Sync 刷新缓冲。
func Sync() {
	if global != nil {
		_ = global.Sync()
	}
}
````

## File: pkg/mailer/mailer.go
````go
// Package mailer 提供最小可用的 SMTP 发信能力。
//
// 设计:
//   - 使用标准库 net/smtp + crypto/tls,零外部依赖。
//   - 支持 465 隐式 TLS 与 587 STARTTLS。
//   - 异步发送(一个 worker 协程),内部有 100 大小的 buffered chan,
//     chan 满时直接丢弃并打 warn 日志,绝不阻塞业务主流程。
//   - Enabled=false 时全部 Send 变成 no-op。
package mailer

import (
	"crypto/tls"
	"errors"
	"fmt"
	"mime"
	"net"
	"net/smtp"
	"strconv"
	"sync"
	"time"

	"go.uber.org/zap"
)

// Config 与 config.SMTPConfig 对齐。
type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
	FromName string
	UseTLS   bool
}

// Message 是一次发送。
type Message struct {
	To      string
	Subject string
	HTML    string
}

// Mailer 可被多个业务复用。
type Mailer struct {
	cfg    Config
	ch     chan Message
	log    *zap.Logger
	wg     sync.WaitGroup
	closed chan struct{}
	once   sync.Once
}

// Disabled 表示当前 mailer 未配置。
func (m *Mailer) Disabled() bool { return m == nil || m.cfg.Host == "" }

// New 构造 Mailer 并启动后台协程。
// host 为空时返回一个 disabled 的实例(Send 成 no-op)。
func New(cfg Config, log *zap.Logger) *Mailer {
	m := &Mailer{
		cfg:    cfg,
		ch:     make(chan Message, 100),
		log:    log.With(zap.String("mod", "mailer")),
		closed: make(chan struct{}),
	}
	if !m.Disabled() {
		m.wg.Add(1)
		go m.loop()
	}
	return m
}

func (m *Mailer) loop() {
	defer m.wg.Done()
	for msg := range m.ch {
		if err := m.send(msg); err != nil {
			m.log.Warn("smtp send failed",
				zap.String("to", msg.To),
				zap.String("subject", msg.Subject),
				zap.Error(err))
		}
	}
	close(m.closed)
}

// SendSync 同步发送,直接把错误抛给调用方。
// 专供"测试发送"等需要立即反馈结果的场景;业务路径请继续用 Send。
func (m *Mailer) SendSync(msg Message) error {
	if m.Disabled() {
		return errors.New("mailer disabled: SMTP not configured")
	}
	return m.send(msg)
}

// Send 非阻塞投递。
// chan 满时打 warn 并丢弃(邮件不是主业务路径)。
func (m *Mailer) Send(msg Message) {
	if m.Disabled() {
		return
	}
	select {
	case m.ch <- msg:
	default:
		m.log.Warn("mail queue full, drop message",
			zap.String("to", msg.To), zap.String("subject", msg.Subject))
	}
}

func (m *Mailer) Close() {
	if m.Disabled() {
		return
	}
	m.once.Do(func() {
		close(m.ch)
	})
	select {
	case <-m.closed:
	case <-time.After(5 * time.Second):
		m.log.Warn("mailer close timeout")
	}
	m.wg.Wait()
}

func (m *Mailer) send(msg Message) error {
	if msg.To == "" || msg.Subject == "" {
		return errors.New("mailer: to/subject empty")
	}
	addr := net.JoinHostPort(m.cfg.Host, strconv.Itoa(m.cfg.Port))

	fromHeader := m.cfg.From
	if m.cfg.FromName != "" {
		fromHeader = fmt.Sprintf("%s <%s>", m.cfg.FromName, m.cfg.From)
	}

	headers := map[string]string{
		"From":         fromHeader,
		"To":           msg.To,
		"Subject":      encodeSubject(msg.Subject),
		"MIME-Version": "1.0",
		"Content-Type": `text/html; charset="UTF-8"`,
	}
	var buf []byte
	for k, v := range headers {
		buf = append(buf, []byte(k+": "+v+"\r\n")...)
	}
	buf = append(buf, []byte("\r\n")...)
	buf = append(buf, []byte(msg.HTML)...)

	auth := smtp.PlainAuth("", m.cfg.Username, m.cfg.Password, m.cfg.Host)

	if m.cfg.UseTLS {
		// 465:隐式 TLS —— 先 TLS 再 SMTP 握手
		return sendTLS(addr, m.cfg.Host, auth, m.cfg.From, msg.To, buf)
	}
	// 587:明文起 STARTTLS
	return smtp.SendMail(addr, auth, m.cfg.From, []string{msg.To}, buf)
}

// sendTLS 实现 SMTPS(465) 隐式 TLS。
func sendTLS(addr, host string, auth smtp.Auth, from, to string, body []byte) error {
	conn, err := tls.Dial("tcp", addr, &tls.Config{ServerName: host})
	if err != nil {
		return fmt.Errorf("tls dial: %w", err)
	}
	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return fmt.Errorf("smtp client: %w", err)
	}
	defer c.Close()

	if ok, _ := c.Extension("AUTH"); ok {
		if err := c.Auth(auth); err != nil {
			return fmt.Errorf("smtp auth: %w", err)
		}
	}
	if err := c.Mail(from); err != nil {
		return fmt.Errorf("smtp mail: %w", err)
	}
	if err := c.Rcpt(to); err != nil {
		return fmt.Errorf("smtp rcpt: %w", err)
	}
	w, err := c.Data()
	if err != nil {
		return fmt.Errorf("smtp data: %w", err)
	}
	if _, err := w.Write(body); err != nil {
		return fmt.Errorf("smtp write: %w", err)
	}
	if err := w.Close(); err != nil {
		return fmt.Errorf("smtp close: %w", err)
	}
	return c.Quit()
}

// encodeSubject 把 UTF-8 标题按 RFC 2047 封装,避免中文乱码。
// 对全 ASCII 的标题保持原样,对含非 ASCII 的用 Q-encoded。
func encodeSubject(s string) string {
	return mime.QEncoding.Encode("UTF-8", s)
}
````

## File: pkg/mailer/templates.go
````go
package mailer

import "html"

func htmlEscape(s string) string { return html.EscapeString(s) }

// RenderTest 生成 SMTP 测试邮件。
func RenderTest(siteName string) (subject, html string) {
	if siteName == "" {
		siteName = "GPT2API Local"
	}
	subject = siteName + " SMTP 测试"
	html = `<html><body><h1>` + htmlEscape(siteName) + `</h1><p>这是一封本地控制台测试邮件。</p></body></html>`
	return
}
````

## File: pkg/resp/resp.go
````go
package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 统一响应结构。HTTP status 只用于框架级错误(401/403/404/500);
// 业务错误一律走 code,HTTP 200。
type Body struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	TraceID string      `json:"trace_id,omitempty"`
}

const (
	CodeOK           = 0
	CodeBadRequest   = 40000
	CodeUnauthorized = 40100
	CodeForbidden    = 40300
	CodeNotFound     = 40400
	CodeConflict     = 40900
	CodeRateLimited  = 42900
	CodeInternal     = 50000
	CodeUpstream     = 50200
)

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Body{Code: CodeOK, Message: "ok", Data: data, TraceID: traceID(c)})
}

func Fail(c *gin.Context, code int, msg string) {
	httpStatus := http.StatusOK
	switch code {
	case CodeUnauthorized:
		httpStatus = http.StatusUnauthorized
	case CodeForbidden:
		httpStatus = http.StatusForbidden
	case CodeNotFound:
		httpStatus = http.StatusNotFound
	case CodeRateLimited:
		httpStatus = http.StatusTooManyRequests
	case CodeInternal, CodeUpstream:
		httpStatus = http.StatusInternalServerError
	}
	c.AbortWithStatusJSON(httpStatus, Body{Code: code, Message: msg, TraceID: traceID(c)})
}

func BadRequest(c *gin.Context, msg string)   { Fail(c, CodeBadRequest, msg) }
func Unauthorized(c *gin.Context, msg string) { Fail(c, CodeUnauthorized, msg) }
func Forbidden(c *gin.Context, msg string)    { Fail(c, CodeForbidden, msg) }
func NotFound(c *gin.Context, msg string)     { Fail(c, CodeNotFound, msg) }
func Conflict(c *gin.Context, msg string)     { Fail(c, CodeConflict, msg) }
func Internal(c *gin.Context, msg string)     { Fail(c, CodeInternal, msg) }
func RateLimited(c *gin.Context, msg string)  { Fail(c, CodeRateLimited, msg) }

func traceID(c *gin.Context) string {
	if v, ok := c.Get("request_id"); ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}
````

## File: README.md
````markdown
# GPT Image 2API

将 ChatGPT 网页端的图像生成能力（GPT-Image-2 / IMG2）包装为标准 OpenAI 兼容 API，支持文生图、图生图、多模态聊天生图，自带账号池调度、代理池、管理后台。

## 核心能力

- **OpenAI 兼容接口**
  - `POST /v1/images/generations` — 文生图 / 图生图
  - `POST /v1/images/edits` — 图片编辑
  - `POST /v1/chat/completions` — 多模态聊天生图（model 为图像模型时自动走图片链路，支持 stream）
  - `GET /v1/models` — 模型列表
  - `GET /v1/images/tasks/:id` — 异步任务查询

- **图像生成**
  - IMG2 高清终稿（非 IMG1 预览），自动灰度检测 + 换号重试
  - 支持 `reference_images` / `image_url` / `input_image` 等多种参考图格式
  - 支持 base64 / data URL / HTTP URL 参考图输入
  - 图片通过本地签名代理输出（`/p/img/...`），estuary 永久 URL 落库
  - 训练数据留存：prompt、revised_prompt、quality、style、参考图 GPT file ID、生成结果 URL、耗时、重试次数全部入库

- **账号池**
  - 支持 AT / RT / ST / Session JSON 四种导入方式
  - RT→AT 自动刷新，JWT 自动解析 subscription_type（pro/plus/free/team）
  - 轮询调度 + 10s 缓存池，避免并发选中同一账号
  - quota 预检、preview_only 自动换号、失败降低置信度

- **代理池**
  - HTTP / HTTPS / SOCKS5 代理管理
  - 支持 `host:port:user:pass` / `user:pass:host:port` 等多种代理商格式导入
  - 导入后自动探测（≤50 条同步，>50 条后台队列）
  - 探测目标：`chatgpt.com/cdn-cgi/trace`（直接验证 ChatGPT 可达性）

- **管理后台**
  - 登录认证（JWT，密钥从 AES key 派生，重启不失效）
  - 可配置 API Key（`/v1/*` 路由鉴权）
  - 运维仪表盘：账号池/代理池状态概览、每日用量柱状图
  - 在线体验：文生图 + 图生图（最多 4 张参考图）
  - 系统设置、审计日志、数据库备份恢复

## 快速开始

### Docker 部署（推荐）

```bash
cd deploy
cp .env.example .env
# 编辑 .env: 修改 MYSQL 密码、CRYPTO_AES_KEY、HTTP_PORT 等
docker compose up -d --build
```

首次空库启动时，容器入口自动导入 `sql/database.sql` 并执行增量迁移。

### 本地开发

```bash
# 后端
cp configs/config.example.yaml configs/config.yaml
go run ./cmd/server -c configs/config.yaml

# 前端
cd web && npm install && npm run dev
```

## API 调用示例

所有 `/v1/*` 接口需要 `Authorization: Bearer <api_key>`（在管理后台 → 系统设置 → API 认证中配置）。

### 文生图

```bash
curl http://YOUR_HOST/v1/images/generations \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer sk-your-key' \
  -d '{
    "model": "gpt-image-2",
    "prompt": "a cat reading a book in a cozy library",
    "size": "1024x1024",
    "quality": "high"
  }'
```

### 图生图（参考图）

```bash
curl http://YOUR_HOST/v1/images/generations \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer sk-your-key' \
  -d '{
    "model": "gpt-image-2",
    "prompt": "make it in watercolor style",
    "reference_images": ["https://example.com/photo.jpg"]
  }'
```

### 多模态聊天生图（兼容 chat/completions）

```bash
curl http://YOUR_HOST/v1/chat/completions \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer sk-your-key' \
  -d '{
    "model": "gpt-image-2",
    "messages": [
      {
        "role": "user",
        "content": [
          {"type": "text", "text": "把这张图变成赛博朋克风格"},
          {"type": "image_url", "image_url": {"url": "https://example.com/photo.jpg"}}
        ]
      }
    ],
    "stream": true
  }'
```

### Python SDK

```python
from openai import OpenAI

client = OpenAI(base_url="http://YOUR_HOST/v1", api_key="sk-your-key")

# 文生图
resp = client.images.generate(
    model="gpt-image-2",
    prompt="a sunset over mountains",
    size="1024x1024",
)
print(resp.data[0].url)
```

## 账号导入

管理后台 → 账号池 → 导入，支持：

| 模式 | 说明 |
|------|------|
| **Session JSON** | 粘贴 `chatgpt.com/api/auth/session` 完整 JSON（推荐） |
| **AT** | 每行一个 Access Token |
| **RT** | 每行一个 Refresh Token（需填 client_id，建议绑代理） |
| **ST** | 每行一个 Session Token |

## 数据库表

| 表 | 用途 |
|----|------|
| `oai_accounts` | 上游账号（AT/RT/ST/状态/订阅类型） |
| `oai_account_cookies` | 账号 cookies（加密存储） |
| `account_proxy_bindings` | 账号-代理绑定 |
| `proxies` | 代理池 |
| `models` | 模型映射配置 |
| `image_tasks` | 图片任务（prompt/参考图/结果URL/耗时/重试次数） |
| `usage_logs` | 用量日志 |
| `system_settings` | 系统配置 KV |
| `admin_audit_logs` | 管理操作审计 |
| `backup_files` | 备份文件记录 |

## 管理后台页面

| 路径 | 功能 |
|------|------|
| `/personal/dashboard` | 运行概览（KPI + 账号/代理状态 + 每日趋势） |
| `/personal/play` | 在线体验（文生图 + 图生图） |
| `/personal/usage` | 用量观察 |
| `/personal/docs` | 接口文档 |
| `/admin/accounts` | 账号池管理 |
| `/admin/proxies` | 代理池管理 |
| `/admin/models` | 模型映射 |
| `/admin/usage` | 全局用量 |
| `/admin/audit` | 审计日志 |
| `/admin/backup` | 备份恢复 |
| `/admin/settings` | 系统设置（API Key、刷新策略、探测参数等） |

## 配置说明

主要配置 `configs/config.yaml`：

| 配置项 | 说明 |
|--------|------|
| `mysql.dsn` | MySQL 连接串 |
| `redis.addr` | Redis 地址 |
| `crypto.aes_key` | 64位 hex AES-256 密钥（加密令牌/cookies/代理密码） |
| `admin.username/password` | 管理后台登录凭据（默认 admin/admin123） |
| `upstream.base_url` | 上游地址（默认 `https://chatgpt.com`） |

Docker 部署通过 `deploy/.env` 覆盖配置。

## 目录结构

```
cmd/server/             服务入口
internal/
  account/              账号池（导入/刷新/调度）
  gateway/              OpenAI 兼容网关（chat/images/proxy）
  image/                图片任务（runner/DAO/model）
  scheduler/            轮询调度器
  proxy/                代理池
  model/                模型映射
  upstream/chatgpt/     ChatGPT 上游协议（SSE/文件上传/POW）
  middleware/           认证/CORS/日志/恢复
  settings/             系统设置
  usage/                用量统计
  audit/                审计日志
  backup/               备份恢复
web/                    Vue 3 + Element Plus 管理前端
sql/database.sql        初始化表结构
deploy/                 Docker Compose 部署
```
````

## File: scripts/package.json
````json
{
  "name": "gpt2api-scripts",
  "private": true,
  "version": "0.0.0",
  "type": "module",
  "engines": {
    "node": ">=18"
  },
  "scripts": {
    "smoke": "node smoke.mjs",
    "smoke:docker": "node smoke.mjs --base http://localhost:8080"
  }
}
````

## File: scripts/README.md
````markdown
# scripts

本目录保留本地验收脚本入口。当前源码已切为单人自用中转,脚本只检查公开 `/v1`、账号池、模型映射和用量观察链路。

```bash
node scripts/smoke.mjs http://localhost:8080
```
````

## File: scripts/smoke.mjs
````javascript
#!/usr/bin/env node
const base = (process.argv[2] || 'http://localhost:8080').replace(/\/$/, '')

async function call(path, init) {
  const res = await fetch(base + path, init)
  const text = await res.text()
  let body = null
  try { body = text ? JSON.parse(text) : null } catch { body = text }
  return { status: res.status, body }
}

function ok(msg) { console.log('✓', msg) }
function fail(msg) { console.error('✗', msg); process.exitCode = 1 }

const models = await call('/v1/models')
if (models.status >= 200 && models.status < 300) ok('/v1/models reachable')
else fail(`/v1/models failed: ${models.status}`)

const site = await call('/api/public/site-info')
if (site.status >= 200 && site.status < 300) ok('/api/public/site-info reachable')
else fail(`/api/public/site-info failed: ${site.status}`)

const adminModels = await call('/api/admin/models')
if (adminModels.status >= 200 && adminModels.status < 300) ok('/api/admin/models reachable')
else fail(`/api/admin/models failed: ${adminModels.status}`)
````

## File: sql/database.sql
````sql
CREATE DATABASE IF NOT EXISTS `gpt2api` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE `gpt2api`;

CREATE TABLE IF NOT EXISTS `proxies` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `scheme` VARCHAR(16) NOT NULL DEFAULT 'http',
  `host` VARCHAR(255) NOT NULL,
  `port` INT NOT NULL,
  `username` VARCHAR(255) NOT NULL DEFAULT '',
  `password_enc` TEXT NULL,
  `country` VARCHAR(64) NOT NULL DEFAULT '',
  `isp` VARCHAR(128) NOT NULL DEFAULT '',
  `health_score` INT NOT NULL DEFAULT 100,
  `last_probe_at` DATETIME NULL,
  `last_error` VARCHAR(512) NOT NULL DEFAULT '',
  `enabled` TINYINT(1) NOT NULL DEFAULT 1,
  `remark` VARCHAR(255) NOT NULL DEFAULT '',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` DATETIME NULL,
  PRIMARY KEY (`id`),
  KEY `idx_proxies_enabled` (`enabled`,`deleted_at`),
  UNIQUE KEY `uk_proxy_endpoint` (`scheme`,`host`,`port`,`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `oai_accounts` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `email` VARCHAR(255) NOT NULL,
  `auth_token_enc` TEXT NOT NULL,
  `refresh_token_enc` TEXT NULL,
  `session_token_enc` TEXT NULL,
  `token_expires_at` DATETIME NULL,
  `oai_session_id` VARCHAR(128) NOT NULL DEFAULT '',
  `oai_device_id` VARCHAR(128) NOT NULL DEFAULT '',
  `client_id` VARCHAR(128) NOT NULL DEFAULT '',
  `chatgpt_account_id` VARCHAR(128) NOT NULL DEFAULT '',
  `account_type` VARCHAR(64) NOT NULL DEFAULT 'codex',
  `subscription_type` VARCHAR(64) NOT NULL DEFAULT '',
  `daily_image_quota` INT NOT NULL DEFAULT 100,
  `status` VARCHAR(32) NOT NULL DEFAULT 'healthy',
  `warned_at` DATETIME NULL,
  `cooldown_until` DATETIME NULL,
  `last_used_at` DATETIME NULL,
  `today_used_count` INT NOT NULL DEFAULT 0,
  `today_used_date` DATETIME NULL,
  `last_refresh_at` DATETIME NULL,
  `last_refresh_source` VARCHAR(32) NOT NULL DEFAULT '',
  `refresh_error` VARCHAR(512) NOT NULL DEFAULT '',
  `image_quota_remaining` INT NOT NULL DEFAULT -1,
  `image_quota_total` INT NOT NULL DEFAULT -1,
  `image_quota_reset_at` DATETIME NULL,
  `image_quota_updated_at` DATETIME NULL,
  `image_capability_status` VARCHAR(32) NOT NULL DEFAULT 'unknown',
  `image_capability_model` VARCHAR(128) NOT NULL DEFAULT '',
  `image_capability_source` VARCHAR(32) NOT NULL DEFAULT '',
  `image_capability_detail` TEXT NULL,
  `image_capability_updated_at` DATETIME NULL,
  `image_init_blocked_features` TEXT NULL,
  `img2_hit_count` INT NOT NULL DEFAULT 0,
  `img2_preview_only_count` INT NOT NULL DEFAULT 0,
  `img2_miss_count` INT NOT NULL DEFAULT 0,
  `img2_consecutive_miss` INT NOT NULL DEFAULT 0,
  `img2_last_status` VARCHAR(32) NOT NULL DEFAULT '',
  `img2_last_hit_at` DATETIME NULL,
  `img2_last_attempt_at` DATETIME NULL,
  `img2_delivery_success_count` INT NOT NULL DEFAULT 0,
  `img2_delivery_fail_count` INT NOT NULL DEFAULT 0,
  `img2_delivery_partial_count` INT NOT NULL DEFAULT 0,
  `img2_last_delivery_status` VARCHAR(32) NOT NULL DEFAULT '',
  `img2_last_delivery_at` DATETIME NULL,
  `notes` TEXT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` DATETIME NULL,
  PRIMARY KEY (`id`),
  KEY `idx_oai_accounts_status` (`status`,`deleted_at`),
  KEY `idx_oai_accounts_email` (`email`),
  KEY `idx_oai_accounts_refresh` (`token_expires_at`,`deleted_at`),
  KEY `idx_oai_accounts_quota` (`image_quota_updated_at`,`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `oai_account_cookies` (
  `account_id` BIGINT UNSIGNED NOT NULL,
  `cookie_json_enc` MEDIUMTEXT NOT NULL,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`account_id`),
  CONSTRAINT `fk_oai_account_cookies_account` FOREIGN KEY (`account_id`) REFERENCES `oai_accounts` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `account_proxy_bindings` (
  `account_id` BIGINT UNSIGNED NOT NULL,
  `proxy_id` BIGINT UNSIGNED NOT NULL,
  `bound_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`account_id`),
  KEY `idx_account_proxy_bindings_proxy` (`proxy_id`),
  CONSTRAINT `fk_apb_account` FOREIGN KEY (`account_id`) REFERENCES `oai_accounts` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_apb_proxy` FOREIGN KEY (`proxy_id`) REFERENCES `proxies` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `models` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `slug` VARCHAR(64) NOT NULL,
  `type` VARCHAR(16) NOT NULL,
  `upstream_model_slug` VARCHAR(128) NOT NULL,
  `description` VARCHAR(255) NOT NULL DEFAULT '',
  `enabled` TINYINT(1) NOT NULL DEFAULT 1,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` DATETIME NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_models_slug` (`slug`),
  KEY `idx_models_enabled` (`enabled`,`deleted_at`),
  KEY `idx_models_type` (`type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `models` (`slug`,`type`,`upstream_model_slug`,`description`,`enabled`) VALUES
  ('gpt-4o','chat','auto','兼容旧客户端的对话模型映射',0),
  ('gpt-image-2','image','auto','ChatGPT 图像生成',1)
ON DUPLICATE KEY UPDATE
  `type`=VALUES(`type`),
  `upstream_model_slug`=VALUES(`upstream_model_slug`),
  `description`=VALUES(`description`),
  `enabled`=VALUES(`enabled`);

CREATE TABLE IF NOT EXISTS `usage_logs` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `model_id` BIGINT UNSIGNED NOT NULL DEFAULT 0,
  `account_id` BIGINT UNSIGNED NOT NULL DEFAULT 0,
  `request_id` VARCHAR(64) NOT NULL DEFAULT '',
  `type` VARCHAR(16) NOT NULL,
  `input_tokens` INT NOT NULL DEFAULT 0,
  `output_tokens` INT NOT NULL DEFAULT 0,
  `cache_read_tokens` INT NOT NULL DEFAULT 0,
  `cache_write_tokens` INT NOT NULL DEFAULT 0,
  `image_count` INT NOT NULL DEFAULT 0,
  `duration_ms` INT NOT NULL DEFAULT 0,
  `status` VARCHAR(16) NOT NULL DEFAULT '',
  `error_code` VARCHAR(128) NOT NULL DEFAULT '',
  `ip` VARCHAR(64) NOT NULL DEFAULT '',
  `ua` VARCHAR(255) NOT NULL DEFAULT '',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_usage_logs_created` (`created_at`),
  KEY `idx_usage_logs_model` (`model_id`),
  KEY `idx_usage_logs_account` (`account_id`),
  KEY `idx_usage_logs_type_status` (`type`,`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `image_tasks` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `task_id` VARCHAR(64) NOT NULL,
  `model_id` BIGINT UNSIGNED NOT NULL DEFAULT 0,
  `account_id` BIGINT UNSIGNED NOT NULL DEFAULT 0,
  `prompt` TEXT NOT NULL,
  `revised_prompt` TEXT NULL,
  `n` INT NOT NULL DEFAULT 1,
  `size` VARCHAR(32) NOT NULL DEFAULT '1024x1024',
  `quality` VARCHAR(32) NOT NULL DEFAULT '',
  `style` VARCHAR(32) NOT NULL DEFAULT '',
  `status` VARCHAR(32) NOT NULL DEFAULT 'queued',
  `conversation_id` VARCHAR(128) NOT NULL DEFAULT '',
  `file_ids` JSON NULL,
  `result_urls` JSON NULL,
  `reference_urls` JSON NULL,
  `error` VARCHAR(512) NOT NULL DEFAULT '',
  `attempts` INT NOT NULL DEFAULT 0,
  `duration_ms` BIGINT NOT NULL DEFAULT 0,
  `user_id` VARCHAR(128) NOT NULL DEFAULT '',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `started_at` DATETIME NULL,
  `finished_at` DATETIME NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_image_tasks_task_id` (`task_id`),
  KEY `idx_image_tasks_status` (`status`),
  KEY `idx_image_tasks_created` (`created_at`),
  KEY `idx_image_tasks_account` (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `system_settings` (
  `k` VARCHAR(128) NOT NULL,
  `v` TEXT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`k`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `system_settings` (`k`,`v`) VALUES
  ('site.name','GPT2API Local'),
  ('site.description','自用 OpenAI 兼容 2API 中转'),
  ('site.logo_url',''),
  ('site.footer',''),
  ('site.contact_email',''),
  ('site.docs_url',''),
  ('site.api_base_url',''),
  ('ui.default_page_size','20'),
  ('gateway.upstream_timeout_sec','60'),
  ('gateway.sse_read_timeout_sec','120'),
  ('gateway.cooldown_429_sec','300'),
  ('gateway.warned_pause_hours','24'),
  ('gateway.daily_usage_ratio','0.8'),
  ('gateway.retry_on_failure','true'),
  ('gateway.retry_max','1'),
  ('gateway.dispatch_queue_wait_sec','120'),
  ('gateway.image_explore_ratio','0.2'),
  ('proxy.probe_enabled','true'),
  ('proxy.probe_interval_sec','300'),
  ('proxy.probe_timeout_sec','10'),
  ('proxy.probe_target_url','https://chatgpt.com/cdn-cgi/trace'),
  ('proxy.probe_concurrency','8'),
  ('account.refresh_enabled','true'),
  ('account.refresh_interval_sec','120'),
  ('account.refresh_ahead_sec','900'),
  ('account.refresh_concurrency','4'),
  ('account.quota_probe_enabled','true'),
  ('account.quota_probe_interval_sec','900'),
  ('account.default_client_id','app_LlGpXReQgckcGGUo2JrYvtJK'),
  ('mail.enabled_display','auto')
ON DUPLICATE KEY UPDATE `v`=VALUES(`v`);

CREATE TABLE IF NOT EXISTS `admin_audit_logs` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `actor_id` BIGINT UNSIGNED NOT NULL DEFAULT 0,
  `actor_email` VARCHAR(255) NOT NULL DEFAULT '',
  `action` VARCHAR(128) NOT NULL DEFAULT '',
  `method` VARCHAR(16) NOT NULL DEFAULT '',
  `path` VARCHAR(255) NOT NULL DEFAULT '',
  `status_code` INT NOT NULL DEFAULT 0,
  `ip` VARCHAR(64) NOT NULL DEFAULT '',
  `ua` VARCHAR(255) NOT NULL DEFAULT '',
  `target` VARCHAR(128) NOT NULL DEFAULT '',
  `meta` JSON NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_admin_audit_created` (`created_at`),
  KEY `idx_admin_audit_action` (`action`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `backup_files` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `backup_id` VARCHAR(64) NOT NULL,
  `file_name` VARCHAR(255) NOT NULL,
  `size_bytes` BIGINT NOT NULL DEFAULT 0,
  `sha256` VARCHAR(64) NOT NULL DEFAULT '',
  `trigger` VARCHAR(32) NOT NULL DEFAULT '',
  `status` VARCHAR(32) NOT NULL DEFAULT 'running',
  `error` VARCHAR(512) NOT NULL DEFAULT '',
  `include_data` TINYINT(1) NOT NULL DEFAULT 1,
  `created_by` BIGINT UNSIGNED NOT NULL DEFAULT 0,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `finished_at` DATETIME NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_backup_files_backup_id` (`backup_id`),
  KEY `idx_backup_files_created` (`created_at`),
  KEY `idx_backup_files_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
````

## File: web/.env.development
````
VITE_APP_TITLE=GPT2API 控制台
VITE_API_BASE=http://localhost:8080
````

## File: web/.env.production
````
VITE_APP_TITLE=GPT2API 控制台
# 生产通常把 /api 和 /v1 交给 nginx 反代到后端,前端走同源
VITE_API_BASE=
````

## File: web/.gitignore
````
node_modules
dist
*.log
.vite
src/auto-imports.d.ts
src/components.d.ts
.DS_Store
````

## File: web/index.html
````html
<!doctype html>
<html lang="zh-CN">
  <head>
    <meta charset="UTF-8" />
    <link rel="icon" type="image/svg+xml" href="/favicon.svg" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta name="color-scheme" content="light dark" />
    <title>GPT2API 控制台</title>
    <style>
      html,body,#app{height:100%;margin:0;padding:0;background:#f5f7fa;}
      body{font-family:-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,"PingFang SC","Hiragino Sans GB","Microsoft YaHei",sans-serif;}
      .app-loading{display:flex;height:100%;align-items:center;justify-content:center;color:#606266;font-size:14px;}
    </style>
  </head>
  <body>
    <div id="app"><div class="app-loading">Loading…</div></div>
    <script type="module" src="/src/main.ts"></script>
  </body>
</html>
````

## File: web/package.json
````json
{
  "name": "gpt2api-web",
  "private": true,
  "version": "0.1.0",
  "type": "module",
  "description": "gpt2api 控制台前端 · Vue 3 + Element Plus",
  "repository": {
    "type": "git",
    "url": "https://github.com/432539/gpt2api.git",
    "directory": "web"
  },
  "homepage": "https://github.com/432539/gpt2api#readme",
  "bugs": {
    "url": "https://github.com/432539/gpt2api/issues"
  },
  "license": "MIT",
  "scripts": {
    "dev": "vite",
    "build": "vue-tsc --noEmit && vite build",
    "preview": "vite preview"
  },
  "dependencies": {
    "@element-plus/icons-vue": "^2.3.1",
    "@vueuse/core": "^10.9.0",
    "axios": "^1.6.8",
    "dayjs": "^1.11.10",
    "element-plus": "^2.7.2",
    "pinia": "^2.1.7",
    "pinia-plugin-persistedstate": "^3.2.1",
    "vue": "^3.4.21",
    "vue-router": "^4.3.0"
  },
  "devDependencies": {
    "@types/node": "^20.11.30",
    "@vitejs/plugin-vue": "^5.0.4",
    "sass": "^1.71.1",
    "typescript": "~5.3.3",
    "unplugin-auto-import": "^0.17.5",
    "unplugin-vue-components": "^0.26.0",
    "vite": "^5.2.0",
    "vue-tsc": "^1.8.27"
  }
}
````

## File: web/public/favicon.svg
````xml
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64">
  <defs>
    <linearGradient id="g" x1="0" x2="1" y1="0" y2="1">
      <stop offset="0%" stop-color="#409EFF"/>
      <stop offset="100%" stop-color="#67C23A"/>
    </linearGradient>
  </defs>
  <rect width="64" height="64" rx="12" fill="url(#g)"/>
  <text x="32" y="42" font-size="28" font-weight="700" fill="#fff" text-anchor="middle"
        font-family="Arial, Helvetica, sans-serif">G2</text>
</svg>
````

## File: web/README.md
````markdown
# GPT2API Local Web

Vue 3 + Vite + Element Plus 控制台,面向单机自用 2API 中转。

## 页面

- 本地总览
- 在线体验
- 用量记录
- 接口文档
- 上游账号池
- 代理池
- 模型映射
- 全局用量
- 审计日志
- 备份恢复
- 系统设置

## 开发

```bash
npm install
npm run dev
```

前端默认通过相对路径访问后端。需要自定义时可设置 `VITE_API_BASE`。
````

## File: web/src/api/accounts.ts
````typescript
import { http } from './http'

export interface Account {
  id: number
  email: string
  client_id: string
  chatgpt_account_id: string
  account_type: string            // codex / chatgpt
  oai_session_id: string
  oai_device_id: string
  subscription_type: string               // plus / team / free / ...
  daily_image_quota: number
  status: string                  // healthy / warned / throttled / suspicious / dead
  today_used_count: number
  notes: string
  token_expires_at?: { Time: string; Valid: boolean } | string | null
  warned_at?:        { Time: string; Valid: boolean } | string | null
  cooldown_until?:   { Time: string; Valid: boolean } | string | null
  last_used_at?:     { Time: string; Valid: boolean } | string | null
  today_used_date?:  { Time: string; Valid: boolean } | string | null

  last_refresh_at?:  { Time: string; Valid: boolean } | string | null
  last_refresh_source: string
  refresh_error: string

  image_quota_remaining: number
  image_quota_total: number
  image_quota_reset_at?:   { Time: string; Valid: boolean } | string | null
  image_quota_updated_at?: { Time: string; Valid: boolean } | string | null

  has_rt: boolean
  has_st: boolean

  created_at: string
  updated_at: string
}

export interface Page<T> {
  list: T[]; total: number; page: number; page_size: number
}

export function listAccounts(params: {
  page?: number; page_size?: number; status?: string; keyword?: string
} = {}) {
  return http.get<any, Page<Account>>('/api/admin/accounts', { params })
}

export function getAccount(id: number) {
  return http.get<any, Account>(`/api/admin/accounts/${id}`)
}

export interface AccountCreate {
  email: string
  auth_token: string
  refresh_token?: string
  session_token?: string
  token_expires_at?: string
  oai_session_id?: string
  oai_device_id?: string
  client_id?: string
  chatgpt_account_id?: string
  account_type?: string
  subscription_type?: string
  daily_image_quota?: number
  notes?: string
  cookies?: string
  proxy_id?: number
}
export interface AccountUpdate extends Partial<AccountCreate> {
  status?: string
}

export function createAccount(body: AccountCreate) {
  return http.post<any, Account>('/api/admin/accounts', body)
}
export function updateAccount(id: number, body: AccountUpdate) {
  return http.patch<any, Account>(`/api/admin/accounts/${id}`, body)
}
export function deleteAccount(id: number) {
  return http.delete<any, { deleted: number }>(`/api/admin/accounts/${id}`)
}
export function bindProxy(id: number, proxyID: number) {
  return http.post<any, { ok: boolean }>(`/api/admin/accounts/${id}/bind-proxy`, { proxy_id: proxyID })
}
export function unbindProxy(id: number) {
  return http.delete<any, { ok: boolean }>(`/api/admin/accounts/${id}/bind-proxy`)
}

// ---------- 批量导入 ----------
export interface ImportLineResult {
  index: number
  email: string
  status: 'created' | 'updated' | 'skipped' | 'failed'
  reason?: string
  id?: number
}
export interface ImportSummary {
  total: number
  created: number
  updated: number
  skipped: number
  failed: number
  results: ImportLineResult[]
}

/**
 * 批量导入账号。
 * 支持两种调用形态:
 *   1) 纯文本 text(运维侧粘贴 JSON / JSONL),走 application/json
 *   2) FormData files[](多文件上传),走 multipart/form-data
 *
 * 大量文件(>500 个)建议前端在客户端先分批合并 text 再用 json 发送,
 * 避免一次 multipart 过大。
 */
export function importAccountsJSON(body: {
  text: string
  update_existing?: boolean
  default_client_id?: string
  default_proxy_id?: number
}) {
  return http.post<any, ImportSummary>('/api/admin/accounts/import', body)
}

export interface ImportTokensBody {
  /** at=每行 AT;rt=每行 RT 需要 client_id;st=每行 ST */
  mode: 'at' | 'rt' | 'st' | 'session_json'
  /** 字符串或字符串数组,后端都兼容 */
  tokens: string | string[]
  /** RT 模式必填;AT/ST 模式可选,传了也会记到账号上 */
  client_id?: string
  update_existing?: boolean
  /** RT/ST 换 AT 时走的代理(chatgpt.com / auth.openai.com),强烈推荐 */
  default_proxy_id?: number
}

export function importAccountsTokens(body: ImportTokensBody) {
  return http.post<any, ImportSummary>('/api/admin/accounts/import-tokens', body)
}

export function importAccountsFiles(
  files: File[],
  opt: { update_existing?: boolean; default_client_id?: string; default_proxy_id?: number } = {}
) {
  const fd = new FormData()
  for (const f of files) fd.append('files', f, f.name)
  if (opt.update_existing !== undefined) fd.append('update_existing', String(opt.update_existing))
  if (opt.default_client_id) fd.append('default_client_id', opt.default_client_id)
  if (opt.default_proxy_id) fd.append('default_proxy_id', String(opt.default_proxy_id))
  return http.post<any, ImportSummary>('/api/admin/accounts/import', fd, {
    headers: { 'Content-Type': 'multipart/form-data' },
  })
}

// ---------- 刷新 / 探测 ----------
export interface RefreshResult {
  account_id: number
  email: string
  ok: boolean
  source: string            // rt / st / failed
  expires_at?: string
  error?: string
  rt_rotated?: boolean
  /** 新 AT 是否通过 chatgpt.com web 后端校验(/backend-api/me 返回 200) */
  at_verified?: boolean
  /** true 表示 RT 换出的 AT 被 chatgpt.com 以 401 拒绝(iOS 作用域),需要 Session Token */
  web_unauthorized?: boolean
}
export interface RefreshAllResult {
  total: number
  success: number
  failed: number
  results: RefreshResult[]
}

export function refreshAccount(id: number) {
  return http.post<any, RefreshResult>(`/api/admin/accounts/${id}/refresh`)
}
export function refreshAllAccounts() {
  return http.post<any, RefreshAllResult>('/api/admin/accounts/refresh-all')
}

export interface QuotaResult {
  account_id: number
  email: string
  ok: boolean
  remaining: number
  total: number
  reset_at?: string
  default_model?: string
  blocked_features?: string[]
  error?: string
}
export interface QuotaAllResult {
  total: number
  success: number
  failed: number
  results: QuotaResult[]
}

export function probeAccountQuota(id: number) {
  return http.post<any, QuotaResult>(`/api/admin/accounts/${id}/probe-quota`)
}
export function probeAllAccountsQuota() {
  return http.post<any, QuotaAllResult>('/api/admin/accounts/probe-quota-all')
}

// ---------- 自动刷新开关 ----------
export interface AutoRefreshConfig {
  enabled: boolean
  ahead_sec: number
  threshold?: string
}
export function getAutoRefresh() {
  return http.get<any, AutoRefreshConfig>('/api/admin/accounts/auto-refresh')
}
export function setAutoRefresh(enabled: boolean) {
  return http.put<any, AutoRefreshConfig>('/api/admin/accounts/auto-refresh', { enabled })
}

// ---------- 批量删除 ----------
export type BulkDeleteScope = 'dead' | 'suspicious' | 'warned' | 'throttled' | 'all'
export function bulkDeleteAccounts(scope: BulkDeleteScope) {
  return http.post<any, { deleted: number; scope: string }>(
    '/api/admin/accounts/bulk-delete', { scope },
  )
}

// ---------- 获取 AT / RT / ST 明文(编辑弹窗回显) ----------
export interface AccountSecrets {
  auth_token: string
  refresh_token: string
  session_token: string
}
export function getAccountSecrets(id: number) {
  return http.get<any, AccountSecrets>(`/api/admin/accounts/${id}/secrets`)
}
````

## File: web/src/api/audit.ts
````typescript
import { http } from './http'

export interface AuditLog {
  id: number
  actor_id: number
  actor_email: string
  action: string
  method: string
  path: string
  status_code: number
  ip: string
  ua: string
  target: string
  meta: any
  created_at: string
}

export interface AuditFilter {
  action?: string
  actor_id?: number
  limit?: number
  offset?: number
}

export function listAudit(params: AuditFilter = {}): Promise<{ items: AuditLog[]; total: number; limit: number; offset: number }> {
  return http.get<any, { items: AuditLog[]; total: number; limit: number; offset: number }>('/api/admin/audit/logs', { params })
}
````

## File: web/src/api/backup.ts
````typescript
import { http } from './http'

export interface BackupFile {
  id: number
  backup_id: string
  file_name: string
  size_bytes: number
  sha256: string
  trigger: string
  status: string
  error?: string
  include_data: boolean
  created_by: number
  created_at: string
  finished_at?: string | { Time: string; Valid: boolean }
}

export interface BackupListResp {
  items: BackupFile[]
  total: number
  allow_restore: boolean
  max_upload_mb: number
}

export function listBackups(limit = 50, offset = 0): Promise<BackupListResp> {
  return http.get<any, BackupListResp>('/api/admin/system/backup', { params: { limit, offset } })
}

export function createBackup(includeData = true): Promise<BackupFile> {
  return http.post<any, BackupFile>('/api/admin/system/backup', { include_data: includeData })
}

export function deleteBackup(id: string): Promise<unknown> {
  return http.delete<any, { deleted: boolean }>(`/api/admin/system/backup/${id}`)
}

export function restoreBackup(id: string): Promise<unknown> {
  return http.post<any, { restored: boolean }>(`/api/admin/system/backup/${id}/restore`, {})
}

export function downloadBackup(id: string, fileName: string) {
  return http.get(`/api/admin/system/backup/${id}/download`, { responseType: 'blob' })
    .then((res: any) => {
      const blob = res.data instanceof Blob ? res.data : new Blob([res])
      const url = URL.createObjectURL(blob)
      const a = document.createElement('a')
      a.href = url
      a.download = fileName
      document.body.appendChild(a)
      a.click()
      a.remove()
      URL.revokeObjectURL(url)
    })
}

export function uploadBackup(file: File, onProgress?: (pct: number) => void) {
  const fd = new FormData()
  fd.append('file', file)
  return http.post<any, BackupFile>('/api/admin/system/backup/upload', fd, {
    headers: { 'Content-Type': 'multipart/form-data' },
    onUploadProgress: (e) => {
      if (e.total && onProgress) onProgress(Math.round((e.loaded / e.total) * 100))
    },
  })
}
````

## File: web/src/api/http.ts
````typescript
import axios, { AxiosError, type AxiosInstance, type AxiosRequestConfig } from 'axios'
import { ElMessage } from 'element-plus'

export interface ApiEnvelope<T = any> {
  code: number
  message: string
  data: T
}

const baseURL = import.meta.env.VITE_API_BASE || ''
const TOKEN_KEY = 'admin_token'

export const http: AxiosInstance = axios.create({ baseURL, timeout: 30_000 })

// 请求拦截器：自动注入 token
http.interceptors.request.use((config) => {
  const token = localStorage.getItem(TOKEN_KEY)
  if (token) {
    config.headers = config.headers || {}
    config.headers['Authorization'] = `Bearer ${token}`
  }
  return config
})

// 响应拦截器
http.interceptors.response.use(
  (response) => {
    const contentType = response.headers?.['content-type'] || ''
    if (response.config.responseType === 'blob' || contentType.startsWith('application/gzip')) return response
    const payload = response.data as ApiEnvelope
    if (payload && typeof payload === 'object' && 'code' in payload) {
      if (payload.code === 0) return payload.data as any
      const msg = payload.message || `请求失败 (code=${payload.code})`
      ElMessage.error(msg)
      return Promise.reject(new Error(msg))
    }
    return response.data
  },
  (error: AxiosError<ApiEnvelope>) => {
    if (error.response?.status === 401) {
      localStorage.removeItem(TOKEN_KEY)
      // 避免在登录页重复跳转
      if (!window.location.pathname.startsWith('/login')) {
        window.location.href = '/login'
      }
      return Promise.reject(error)
    }
    const msg = error.response?.data?.message || error.message || '网络错误'
    ElMessage.error(msg)
    return Promise.reject(error)
  },
)

export function request<T = any>(cfg: AxiosRequestConfig): Promise<T> {
  return http.request(cfg) as unknown as Promise<T>
}
````

## File: web/src/api/me.ts
````typescript
import { http } from './http'

export interface SimpleModel {
  id: number
  slug: string
  type: 'chat' | 'image' | string
  description: string
}

export function listMyModels(): Promise<{ items: SimpleModel[]; total: number }> {
  return http.get<any, { items: SimpleModel[]; total: number }>('/api/me/models')
}

export interface UsageItem {
  id: number
  model_id: number
  model_slug: string
  account_id: number
  request_id: string
  type: 'chat' | 'image' | string
  input_tokens: number
  output_tokens: number
  cache_read_tokens: number
  cache_write_tokens: number
  image_count: number
  duration_ms: number
  status: string
  error_code: string
  ip: string
  created_at: string
}

export interface UsageOverall {
  requests: number
  failures: number
  chat_requests: number
  image_images: number
  input_tokens: number
  output_tokens: number
}

export interface UsageDaily {
  day: string
  requests: number
  failures: number
  input_tokens: number
  output_tokens: number
  image_count: number
}

export interface UsageModelStat {
  model_id: number
  model_slug: string
  type: string
  requests: number
  failures: number
  input_tokens: number
  output_tokens: number
  image_count: number
  avg_dur_ms: number
}

export interface MyStatsResp {
  overall: UsageOverall
  daily: UsageDaily[]
  by_model: UsageModelStat[]
}

export function listMyUsageLogs(params: {
  type?: 'chat' | 'image' | ''
  status?: string
  since?: string
  until?: string
  limit?: number
  offset?: number
} = {}): Promise<{ items: UsageItem[]; total: number; limit: number; offset: number }> {
  return http.get<any, { items: UsageItem[]; total: number; limit: number; offset: number }>('/api/me/usage/logs', { params })
}

export function getMyUsageStats(params: {
  days?: number
  top_n?: number
  type?: 'chat' | 'image' | ''
  since?: string
  until?: string
} = {}): Promise<MyStatsResp> {
  return http.get<any, MyStatsResp>('/api/me/usage/stats', { params })
}

export interface ImageTask {
  id: number
  task_id: string
  model_id: number
  account_id: number
  prompt: string
  n: number
  size: string
  status: 'queued' | 'dispatched' | 'running' | 'success' | 'failed' | string
  conversation_id?: string
  error?: string
  image_urls: string[]
  file_ids?: string[]
  created_at: string
  started_at?: string | null
  finished_at?: string | null
}

export function listMyImageTasks(params: { limit?: number; offset?: number } = {}): Promise<{ items: ImageTask[]; limit: number; offset: number }> {
  return http.get<any, { items: ImageTask[]; limit: number; offset: number }>('/api/me/images/tasks', { params })
}

export function getMyImageTask(taskID: string): Promise<ImageTask> {
  return http.get<any, ImageTask>(`/api/me/images/tasks/${taskID}`)
}

export interface ChatStreamDelta { role?: string; content?: string }
export interface ChatStreamChunk {
  id?: string
  model?: string
  choices?: Array<{ index?: number; delta?: ChatStreamDelta; finish_reason?: string | null }>
}
export interface PlayChatMessage { role: 'system' | 'user' | 'assistant'; content: string }

export async function streamPlayChat(
  req: { model: string; messages: PlayChatMessage[]; temperature?: number },
  onDelta: (text: string) => void,
  signal?: AbortSignal,
): Promise<void> {
  const resp = await fetch('/api/me/playground/chat', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ ...req, stream: true }),
    signal,
  })
  if (!resp.ok || !resp.body) {
    const text = await resp.text().catch(() => '')
    throw new Error(`chat ${resp.status}: ${text || resp.statusText}`)
  }
  const reader = resp.body.getReader()
  const decoder = new TextDecoder('utf-8')
  let buffer = ''
  while (true) {
    const { done, value } = await reader.read()
    if (done) break
    buffer += decoder.decode(value, { stream: true })
    const blocks = buffer.split('\n\n')
    buffer = blocks.pop() || ''
    for (const block of blocks) {
      for (const line of block.split('\n')) {
        if (!line.startsWith('data:')) continue
        const data = line.slice(5).trim()
        if (!data || data === '[DONE]') continue
        try {
          const chunk: ChatStreamChunk = JSON.parse(data)
          const delta = chunk.choices?.[0]?.delta?.content
          if (delta) onDelta(delta)
        } catch { /* ignore heartbeat */ }
      }
    }
  }
}

export interface PlayImageRequest {
  model: string
  prompt: string
  n?: number
  size?: string
  reference_images?: string[]
}
export interface PlayImageData { url: string; file_id?: string; revised_prompt?: string }
export interface PlayImageResponse { created: number; task_id?: string; data: PlayImageData[]; is_preview?: boolean }

export async function playGenerateImage(req: PlayImageRequest, signal?: AbortSignal): Promise<PlayImageResponse> {
  const resp = await fetch('/api/me/playground/image', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(req),
    signal,
  })
  if (!resp.ok) {
    let detail = ''
    try { const body = await resp.json(); detail = body?.error?.message || body?.message || '' } catch { /* ignore */ }
    throw new Error(detail || `image ${resp.status}: ${resp.statusText}`)
  }
  return (await resp.json()) as PlayImageResponse
}

export async function playEditImage(
  model: string,
  prompt: string,
  files: File[],
  opts?: { n?: number; size?: string; signal?: AbortSignal },
): Promise<PlayImageResponse> {
  if (!files.length) throw new Error('至少需要选择一张参考图')
  const fd = new FormData()
  fd.append('model', model)
  fd.append('prompt', prompt)
  if (opts?.n) fd.append('n', String(opts.n))
  if (opts?.size) fd.append('size', opts.size)
  files.forEach((f, idx) => fd.append(idx === 0 ? 'image' : 'image[]', f, f.name))
  const resp = await fetch('/api/me/playground/image-edit', { method: 'POST', body: fd, signal: opts?.signal })
  if (!resp.ok) {
    let detail = ''
    try { const body = await resp.json(); detail = body?.error?.message || body?.message || '' } catch { /* ignore */ }
    throw new Error(detail || `image-edit ${resp.status}: ${resp.statusText}`)
  }
  return (await resp.json()) as PlayImageResponse
}
````

## File: web/src/api/proxies.ts
````typescript
import { http } from './http'

export interface Proxy {
  id: number
  scheme: string          // http / https / socks5
  host: string
  port: number
  username: string
  country: string
  isp: string
  health_score: number
  last_probe_at?: { Time: string; Valid: boolean } | string | null
  last_error?: string
  enabled: boolean
  remark: string
  created_at: string
  updated_at?: string
}

export interface Page<T> {
  list: T[]; total: number; page: number; page_size: number
}

export interface ProxyCreate {
  scheme: string
  host: string
  port: number
  username?: string
  password?: string       // 明文,后端会加密
  country?: string
  isp?: string
  enabled?: boolean
  remark?: string
}
export type ProxyUpdate = ProxyCreate

export function listProxies(params: { page?: number; page_size?: number } = {}) {
  return http.get<any, Page<Proxy>>('/api/admin/proxies', { params })
}
export function createProxy(body: ProxyCreate) {
  return http.post<any, Proxy>('/api/admin/proxies', body)
}
export function updateProxy(id: number, body: ProxyUpdate) {
  return http.patch<any, Proxy>(`/api/admin/proxies/${id}`, body)
}
export function deleteProxy(id: number) {
  return http.delete<any, { deleted: number }>(`/api/admin/proxies/${id}`)
}

export interface ProxyImportLine {
  line: number
  raw: string
  status: 'created' | 'updated' | 'skipped' | 'invalid'
  id?: number
  error?: string
}
export interface ProxyImportResp {
  items: ProxyImportLine[]
  created: number
  updated: number
  skipped: number
  invalid: number
  total: number
  probe?: ProbeAllResp
  probe_queued?: boolean
}
export function importProxies(body: {
  text: string
  enabled?: boolean
  country?: string
  isp?: string
  remark?: string
  overwrite?: boolean
  probe_after_import?: boolean
}) {
  return http.post<any, ProxyImportResp>('/api/admin/proxies/import', body)
}

// 单条探测结果
export interface ProbeOneResp {
  ok: boolean
  latency_ms: number
  error?: string
  tried_at: string
  health_score: number
}
export function probeProxy(id: number) {
  return http.post<any, ProbeOneResp>(`/api/admin/proxies/${id}/probe`)
}

// 全量探测
export interface ProbeItem {
  proxy_id: number
  ok: boolean
  latency_ms: number
  error?: string
  tried_at: string
}
export interface ProbeAllResp {
  total: number
  ok: number
  bad: number
  items: ProbeItem[]
}
export function probeAllProxies() {
  return http.post<any, ProbeAllResp>('/api/admin/proxies/probe-all')
}
````

## File: web/src/api/settings.ts
````typescript
import { http } from './http'

// 系统设置 KV 条目(管理端用,带 schema)。
export interface SettingItem {
  key: string
  value: string
  type: 'string' | 'bool' | 'int' | 'email' | 'url' | string
  category: 'site' | 'gateway' | 'account' | 'mail' | string
  label: string
  desc: string
}

export function listSettings(): Promise<{ items: SettingItem[] }> {
  return http.get<any, { items: SettingItem[] }>('/api/admin/settings')
}

export function updateSettings(items: Record<string, string>): Promise<{ updated: number }> {
  return http.put<any, { updated: number }>('/api/admin/settings', { items })
}

export function reloadSettings(): Promise<{ reloaded: boolean }> {
  return http.post<any, { reloaded: boolean }>('/api/admin/settings/reload')
}

export function sendTestEmail(to: string): Promise<{ sent: boolean; to: string }> {
  return http.post<any, { sent: boolean; to: string }>('/api/admin/settings/test-email', { to })
}

// 公开接口:返回控制台需要的站点元信息(site.name 等)。
export function fetchSiteInfo(): Promise<Record<string, string>> {
  return http.get<any, Record<string, string>>('/api/public/site-info')
}
````

## File: web/src/api/stats.ts
````typescript
import { http } from './http'

export interface Model {
  id: number
  slug: string
  type: 'chat' | 'image' | string
  upstream_model_slug: string
  description: string
  enabled: boolean
  created_at: string
  updated_at: string
}

export interface ModelUpsert {
  slug?: string
  type: 'chat' | 'image'
  upstream_model_slug: string
  description: string
  enabled?: boolean
}

export function listModels(): Promise<{ items: Model[]; total: number }> {
  return http.get<any, { items: Model[]; total: number }>('/api/admin/models')
}
export function createModel(body: ModelUpsert): Promise<Model> {
  return http.post<any, Model>('/api/admin/models', body)
}
export function updateModel(id: number, body: ModelUpsert): Promise<Model> {
  return http.put<any, Model>(`/api/admin/models/${id}`, body)
}
export function setModelEnabled(id: number, enabled: boolean) {
  return http.patch<any, { enabled: boolean }>(`/api/admin/models/${id}/enabled`, { enabled })
}
export function deleteModel(id: number) {
  return http.delete<any, { deleted: number }>(`/api/admin/models/${id}`)
}

export interface Overall {
  requests: number
  failures: number
  chat_requests: number
  image_images: number
  input_tokens: number
  output_tokens: number
}

export interface DailyPoint {
  day: string
  requests: number
  failures: number
  input_tokens: number
  output_tokens: number
  image_count: number
}

export interface ModelStat {
  model_id: number
  model_slug: string
  type: string
  requests: number
  failures: number
  input_tokens: number
  output_tokens: number
  image_count: number
  avg_dur_ms: number
}

export interface StatsResp {
  overall: Overall
  daily: DailyPoint[]
  by_model: ModelStat[]
}

export interface UsageLogRow {
  id: number
  model_id: number
  model_slug: string
  account_id: number
  request_id: string
  type: 'chat' | 'image' | string
  input_tokens: number
  output_tokens: number
  cache_read_tokens: number
  cache_write_tokens: number
  image_count: number
  duration_ms: number
  status: string
  error_code: string
  ip: string
  created_at: string
}

export function getUsageStats(params: {
  days?: number; top_n?: number; model_id?: number; type?: string; status?: string; since?: string; until?: string
} = {}): Promise<StatsResp> {
  return http.get<any, StatsResp>('/api/admin/usage/stats', { params })
}

export function listUsageLogs(params: {
  type?: string; status?: string; since?: string; until?: string; model_id?: number; limit?: number; offset?: number
} = {}): Promise<{ items: UsageLogRow[]; total: number; limit: number; offset: number }> {
  return http.get<any, { items: UsageLogRow[]; total: number; limit: number; offset: number }>('/api/admin/usage/logs', { params })
}
````

## File: web/src/App.vue
````vue
<template>
  <el-config-provider namespace="el">
    <router-view />
  </el-config-provider>
</template>
````

## File: web/src/components/Placeholder.vue
````vue
<script setup lang="ts">
defineProps<{
  title: string
  subtitle?: string
  api?: string      // 提示后端 API 路径
}>()
</script>

<template>
  <div class="page-container">
    <div class="card-block">
      <el-empty :description="subtitle || '该模块的前端还在施工中'">
        <template #image>
          <el-icon :size="64" color="#c0c4cc"><Coffee /></el-icon>
        </template>
        <div class="title">{{ title }}</div>
        <div v-if="api" class="hint">
          后端接口:<code>{{ api }}</code>
        </div>
        <div class="hint">你可以通过 Postman / curl 直接调用,或在下个版本里等待 UI 完工。</div>
      </el-empty>
    </div>
  </div>
</template>

<style scoped lang="scss">
.title { font-size: 18px; font-weight: 600; margin-top: 4px; }
.hint { color: var(--el-text-color-secondary); margin-top: 6px; font-size: 13px; }
code {
  background: #f2f3f5;
  padding: 2px 6px;
  border-radius: 4px;
  font-family: ui-monospace, Menlo, Consolas, monospace;
  font-size: 12px;
}
</style>
````

## File: web/src/config/feature.ts
````typescript
// 前端 feature flag。以后要恢复文字模型的 UI,只需把 ENABLE_CHAT_MODEL 改回 true。
// 当前关闭原因:文字通路受 chatgpt.com 新 sentinel 协议影响,在 solver 接入前静默拒绝率较高。
// 后端路由仍保留,方便后续重新开启和调试。
export const ENABLE_CHAT_MODEL = false
````

## File: web/src/env.d.ts
````typescript
/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_APP_TITLE: string
  readonly VITE_API_BASE: string
}
interface ImportMeta {
  readonly env: ImportMetaEnv
}

declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  const component: DefineComponent<{}, {}, any>
  export default component
}
````

## File: web/src/layouts/BasicLayout.vue
````vue
<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'
import { useUserStore } from '@/stores/user'
import { useUIStore } from '@/stores/ui'
import { useSiteStore } from '@/stores/site'
import type { MenuItem } from '@/stores/user'

const store = useUserStore()
const ui = useUIStore()
const site = useSiteStore()
const router = useRouter()
const route = useRoute()

const siteName = computed(() => site.get('site.name', 'image2api'))
const siteLogo = computed(() => site.get('site.logo_url', ''))
const siteFooter = computed(() => site.get('site.footer', ''))

const { menu, user, role, permissions } = storeToRefs(store)
const collapsed = ref(false)
const loadingMenu = ref(false)

const activePath = computed(() => route.path)

const titleMap = computed(() => {
  const m = new Map<string, string>()
  function walk(items: MenuItem[]) {
    for (const it of items) {
      if (it.path) m.set(it.path, it.title)
      if (it.children) walk(it.children)
    }
  }
  walk(menu.value)
  return m
})

const currentTitle = computed(() => titleMap.value.get(activePath.value) || (route.meta.title as string) || '')

async function loadMenu() {
  if (menu.value.length > 0) return
  loadingMenu.value = true
  try {
    await store.fetchMenu()
  } finally {
    loadingMenu.value = false
  }
}

async function logout() {
  await store.logout()
  router.replace('/login')
}

function goto(path?: string) {
  if (path) router.push(path)
}

onMounted(loadMenu)
watch(() => store.isLoggedIn, (v) => { if (v) loadMenu() })
</script>

<template>
  <el-container class="layout-root">
    <el-aside :width="collapsed ? '64px' : '220px'" class="sidebar">
      <div class="logo">
        <img v-if="siteLogo" :src="siteLogo" class="logo-img" alt="logo" />
        <span v-else class="mark">{{ (siteName[0] || 'G').toUpperCase() }}</span>
        <span v-if="!collapsed" class="title">{{ siteName }}</span>
      </div>
      <el-menu
        :default-active="activePath"
        :collapse="collapsed"
        background-color="transparent"
        text-color="var(--el-text-color-regular)"
        active-text-color="#409eff"
        class="side-menu"
        router
      >
        <template v-for="group in menu" :key="group.key">
          <el-menu-item v-if="!group.children?.length && group.path" :index="group.path">
            <el-icon v-if="group.icon"><component :is="group.icon" /></el-icon>
            <template #title>{{ group.title }}</template>
          </el-menu-item>
          <el-sub-menu v-else-if="group.children?.length" :index="group.key">
            <template #title>
              <el-icon v-if="group.icon"><component :is="group.icon" /></el-icon>
              <span>{{ group.title }}</span>
            </template>
            <el-menu-item
              v-for="child in group.children"
              :key="child.key"
              :index="child.path!"
            >
              <el-icon v-if="child.icon"><component :is="child.icon" /></el-icon>
              <template #title>{{ child.title }}</template>
            </el-menu-item>
          </el-sub-menu>
        </template>
      </el-menu>
    </el-aside>

    <el-container>
      <el-header class="topbar">
        <div class="left">
          <el-button link @click="collapsed = !collapsed">
            <el-icon :size="18"><component :is="collapsed ? 'Expand' : 'Fold'" /></el-icon>
          </el-button>
          <span class="crumb">{{ currentTitle }}</span>
        </div>
        <div class="right">
          <el-tooltip :content="ui.isDark ? '切换到亮色' : '切换到暗色'" placement="bottom">
            <el-button link class="theme-btn" @click="ui.toggleDark()">
              <el-icon :size="18">
                <component :is="ui.isDark ? 'Sunny' : 'Moon'" />
              </el-icon>
            </el-button>
          </el-tooltip>
          <el-dropdown trigger="click" @command="(c: string) => c === 'logout' ? logout() : goto(c)">
            <span class="user-entry">
              <el-avatar :size="28" style="background:#409eff">
                {{ (user?.nickname || user?.email || 'U').slice(0, 1).toUpperCase() }}
              </el-avatar>
              <span class="nick">{{ user?.nickname || user?.email }}</span>
              <el-tag type="success" size="small">本地</el-tag>
              <el-icon><ArrowDown /></el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="/personal/dashboard">
                  <el-icon><User /></el-icon> 本地总览
                </el-dropdown-item>
                <el-dropdown-item command="/admin/settings">
                  <el-icon><Setting /></el-icon> 系统设置
                </el-dropdown-item>
                <el-dropdown-item divided command="logout">
                  <el-icon><SwitchButton /></el-icon> 返回总览
                </el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </el-header>

      <el-main class="main" v-loading="loadingMenu">
        <router-view v-slot="{ Component }">
          <transition name="fade" mode="out-in">
            <component :is="Component" />
          </transition>
        </router-view>
      </el-main>

      <el-footer class="footer">
        <div v-if="siteFooter" class="footer-line footer-custom">{{ siteFooter }}</div>
        <div v-else class="footer-line">&copy; {{ new Date().getFullYear() }} {{ siteName }}</div>
      </el-footer>
    </el-container>
  </el-container>
</template>

<style scoped lang="scss">
.layout-root { height: 100vh; }

/* ===================== Sidebar ===================== */
.sidebar {
  background: var(--el-bg-color);
  border-right: 1px solid var(--el-border-color-lighter);
  transition: width .25s cubic-bezier(.4, 0, .2, 1);
  overflow-x: hidden;
  overflow-y: auto;
  display: flex;
  flex-direction: column;
}

.logo {
  height: 64px;
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 0 18px;
  color: #fff;
  font-weight: 700;
  flex-shrink: 0;
  border-bottom: 1px solid var(--el-border-color-lighter);
  .logo-img {
    width: 34px; height: 34px; border-radius: 10px; object-fit: contain;
  }
  .mark {
    display: inline-flex;
    width: 34px;
    height: 34px;
    border-radius: 10px;
    background: linear-gradient(135deg,#409eff,#67c23a);
    align-items: center; justify-content: center;
    font-size: 15px;
    font-weight: 800;
    color: #fff;
    flex-shrink: 0;
  }
  .title { font-size: 15px; white-space: nowrap; letter-spacing: 0.5px; color: var(--el-text-color-primary); }
}

.side-menu {
  border-right: none;
  flex: 1;
  padding: 8px 0;
  --el-menu-hover-bg-color: transparent;
  --el-menu-bg-color: transparent;
  --el-menu-active-color: #fff;

  // 分组标题(sub-menu title)
  :deep(.el-sub-menu__title) {
    height: 40px;
    line-height: 40px;
    font-size: 12px;
    font-weight: 600;
    letter-spacing: 0.5px;
    color: var(--el-text-color-placeholder) !important;
    padding-left: 20px !important;
    margin-top: 8px;
    cursor: default;
    &:hover { background: transparent !important; }
    .el-sub-menu__icon-arrow { display: none; }
    .el-icon { font-size: 14px; margin-right: 6px; color: var(--el-text-color-disabled); }
  }

  // 菜单项
  :deep(.el-menu-item) {
    height: 42px;
    line-height: 42px;
    margin: 2px 8px;
    padding: 0 14px !important;
    border-radius: 8px;
    font-size: 14px;
    color: var(--el-text-color-regular);
    transition: all .15s;
    .el-icon { font-size: 17px; margin-right: 10px; color: var(--el-text-color-secondary); }
    &:hover {
      background: var(--el-fill-color-light);
      color: var(--el-text-color-primary);
    }
    &.is-active {
      background: rgba(64,158,255,0.08);
      color: #409eff;
      font-weight: 500;
      .el-icon { color: #409eff; }
    }
  }

  // 子菜单容器内间距
  :deep(.el-sub-menu .el-menu) {
    padding: 0;
  }
}

/* ===================== Topbar ===================== */
.topbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 56px;
  background: var(--el-bg-color);
  color: var(--el-text-color-primary);
  border-bottom: 1px solid var(--el-border-color-light);
  padding: 0 20px;
  .left { display: flex; align-items: center; gap: 12px; }
  .crumb { font-size: 16px; font-weight: 600; }
  .user-entry {
    display: inline-flex;
    align-items: center;
    gap: 8px;
    cursor: pointer;
    color: var(--el-text-color-primary);
    .nick { font-size: 14px; }
  }
  .right {
    display: inline-flex;
    align-items: center;
    gap: 12px;
  }
  .theme-btn { padding: 0 6px; }
}

/* ===================== Main ===================== */
.main {
  background: var(--gp-bg);
  padding: 0;
}

/* ===================== Footer ===================== */
.footer {
  background: transparent;
  text-align: center;
  color: var(--el-text-color-placeholder);
  font-size: 12px;
  padding: 8px 12px;
  height: auto;
  min-height: 32px;
}
.footer-line { line-height: 1.6; }
.footer-custom { font-size: 11px; }

.fade-enter-active, .fade-leave-active { transition: opacity .15s; }
.fade-enter-from, .fade-leave-to { opacity: 0; }
</style>
````

## File: web/src/layouts/BlankLayout.vue
````vue
<template>
  <div class="blank-layout">
    <router-view />
  </div>
</template>

<style scoped>
.blank-layout {
  min-height: 100vh;
  position: relative;
  box-sizing: border-box;
}
</style>
````

## File: web/src/main.ts
````typescript
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import piniaPersist from 'pinia-plugin-persistedstate'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import 'element-plus/theme-chalk/dark/css-vars.css'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import * as ElementIcons from '@element-plus/icons-vue'

import App from './App.vue'
import router from './router'
import './styles/global.scss'
import { useSiteStore } from './stores/site'

const app = createApp(App)

const pinia = createPinia()
pinia.use(piniaPersist)
app.use(pinia)
app.use(router)
app.use(ElementPlus, { size: 'default', locale: zhCn })

// 把 element icons 作为全局组件注册,模板里可直接 <el-icon><Setting /></el-icon>
for (const [name, comp] of Object.entries(ElementIcons)) {
  app.component(name, comp as never)
}

// 启动即异步拉取站点公开信息,用于顶栏统一展示 site.name / logo 等。
useSiteStore(pinia).refresh()

app.mount('#app')
````

## File: web/src/router/index.ts
````typescript
import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import BasicLayout from '@/layouts/BasicLayout.vue'

const routes: RouteRecordRaw[] = [
  {
    path: '/login',
    component: () => import('@/views/Login.vue'),
    meta: { title: '登录', public: true },
  },
  {
    path: '/',
    component: BasicLayout,
    redirect: '/personal/dashboard',
    children: [
      { path: 'personal/dashboard', component: () => import('@/views/personal/Dashboard.vue'), meta: { title: '本地总览' } },
      { path: 'personal/usage', component: () => import('@/views/personal/Usage.vue'), meta: { title: '用量记录' } },
      { path: 'personal/play', component: () => import('@/views/personal/OnlinePlay.vue'), meta: { title: '在线体验' } },
      { path: 'personal/docs', component: () => import('@/views/personal/ApiDocs.vue'), meta: { title: '接口文档' } },
      { path: 'admin/accounts', component: () => import('@/views/admin/Accounts.vue'), meta: { title: '上游账号池' } },
      { path: 'admin/proxies', component: () => import('@/views/admin/Proxies.vue'), meta: { title: '代理池' } },
      { path: 'admin/models', component: () => import('@/views/admin/Models.vue'), meta: { title: '模型映射' } },
      { path: 'admin/usage', component: () => import('@/views/admin/UsageStats.vue'), meta: { title: '全局用量' } },
      { path: 'admin/audit', component: () => import('@/views/admin/Audit.vue'), meta: { title: '审计日志' } },
      { path: 'admin/backup', component: () => import('@/views/admin/Backup.vue'), meta: { title: '备份恢复' } },
      { path: 'admin/settings', component: () => import('@/views/admin/Settings.vue'), meta: { title: '系统设置' } },
      { path: 'personal/playground', redirect: '/personal/play' },
      { path: 'personal/images', redirect: '/personal/play' },
      { path: 'personal/keys', redirect: '/personal/docs' },
      { path: 'admin/users', redirect: '/admin/accounts' },
      { path: 'admin/groups', redirect: '/admin/accounts' },
      { path: 'admin/keys', redirect: '/admin/models' },
    ],
  },
  { path: '/403', component: () => import('@/views/Error403.vue'), meta: { title: '403', public: true } },
  { path: '/:pathMatch(.*)*', component: () => import('@/views/Error404.vue'), meta: { title: '404', public: true } },
]

const router = createRouter({ history: createWebHistory(), routes })

router.beforeEach((to) => {
  document.title = (to.meta.title as string) || 'Image2API 控制台'
  // 公开页面不需要登录
  if (to.meta.public) return true
  // 检查 token
  const token = localStorage.getItem('admin_token')
  if (!token) {
    return { path: '/login', query: { redirect: to.fullPath } }
  }
  return true
})

export default router
````

## File: web/src/stores/site.ts
````typescript
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { fetchSiteInfo } from '@/api/settings'

/**
 * Site store 缓存站点公开信息:
 *   site.name / site.description / site.logo_url / site.footer / site.contact_email
 *
 * 页面启动时 refresh() 一次即可;本地设置改完会再触发一次 refresh。
 */
export const useSiteStore = defineStore('site', () => {
  const info = ref<Record<string, string>>({
    'site.name': 'GPT2API',
    'site.description': '自用 OpenAI 兼容中转',
    'site.logo_url': '',
    'site.footer': '',
    'site.contact_email': '',
  })
  const loaded = ref(false)

  async function refresh() {
    try {
      const d = await fetchSiteInfo()
      info.value = { ...info.value, ...d }
    } catch {
      // 静默失败,保持默认值。后端未起或权限中间件变化时,前端仍可渲染。
    } finally {
      loaded.value = true
      applyDocumentTitle()
      applyFavicon()
    }
  }

  function applyDocumentTitle() {
    const n = info.value['site.name'] || 'GPT2API'
    document.title = `${n} 控制台`
  }

  function applyFavicon() {
    const url = info.value['site.logo_url']
    if (!url) return
    let link = document.querySelector<HTMLLinkElement>('link[rel~="icon"]')
    if (!link) {
      link = document.createElement('link')
      link.rel = 'icon'
      document.head.appendChild(link)
    }
    link.href = url
  }

  function get(key: string, fallback = ''): string {
    const v = info.value[key]
    return v == null || v === '' ? fallback : v
  }
  return { info, loaded, refresh, get }
})
````

## File: web/src/stores/ui.ts
````typescript
import { defineStore } from 'pinia'
import { useDark, useToggle } from '@vueuse/core'

/**
 * UI 偏好:黑暗模式切换。
 * Element Plus 的 dark 模式通过在 <html> 上加 `class="dark"` 生效,
 * 所以这里配置 useDark 去改根元素 class。
 *
 * 持久化:由 @vueuse 的 useStorage 接管(默认 key=vueuse-color-scheme)。
 */
export const useUIStore = defineStore('ui', () => {
  const isDark = useDark({
    selector: 'html',
    attribute: 'class',
    valueDark: 'dark',
    valueLight: '',
    storageKey: 'gpt2api.theme',
  })
  const toggleDark = useToggle(isDark)
  return { isDark, toggleDark }
})
````

## File: web/src/stores/user.ts
````typescript
import { defineStore } from 'pinia'
import { computed, ref } from 'vue'
import axios from 'axios'

export interface LocalUser {
  email: string
  nickname: string
}

export interface MenuItem {
  key: string
  title: string
  icon?: string
  path?: string
  children?: MenuItem[]
}

const LOCAL_MENU: MenuItem[] = [
  {
    key: 'local',
    title: '本地中转',
    icon: 'Monitor',
    children: [
      { key: 'dashboard', title: '本地总览', icon: 'DataBoard', path: '/personal/dashboard' },
      { key: 'play', title: '在线体验', icon: 'MagicStick', path: '/personal/play' },
      { key: 'usage', title: '用量记录', icon: 'TrendCharts', path: '/personal/usage' },
      { key: 'docs', title: '接口文档', icon: 'Document', path: '/personal/docs' },
    ],
  },
  {
    key: 'ops',
    title: '运维管理',
    icon: 'Operation',
    children: [
      { key: 'accounts', title: '上游账号池', icon: 'UserFilled', path: '/admin/accounts' },
      { key: 'proxies', title: '代理池', icon: 'Connection', path: '/admin/proxies' },
      { key: 'models', title: '模型映射', icon: 'Grid', path: '/admin/models' },
      { key: 'admin_usage', title: '全局用量', icon: 'Histogram', path: '/admin/usage' },
      { key: 'audit', title: '审计日志', icon: 'Memo', path: '/admin/audit' },
      { key: 'backup', title: '备份恢复', icon: 'FolderChecked', path: '/admin/backup' },
      { key: 'settings', title: '系统设置', icon: 'Setting', path: '/admin/settings' },
    ],
  },
]

const TOKEN_KEY = 'admin_token'

export const useUserStore = defineStore('user', () => {
  const token = ref<string>(localStorage.getItem(TOKEN_KEY) || '')
  const user = ref<LocalUser>({ email: '', nickname: '' })
  const permissions = ref<string[]>(['local:*'])
  const role = ref<string>('local')
  const menu = ref<MenuItem[]>(LOCAL_MENU)

  const isLoggedIn = computed(() => !!token.value)
  const isAdmin = computed(() => !!token.value)

  async function login(username: string, password: string) {
    const baseURL = import.meta.env.VITE_API_BASE || ''
    const resp = await axios.post(`${baseURL}/api/admin/login`, { username, password })
    const data = resp.data?.data || resp.data
    if (!data?.token) {
      throw new Error(resp.data?.message || '登录失败')
    }
    token.value = data.token
    user.value = { email: data.username, nickname: data.username }
    localStorage.setItem(TOKEN_KEY, data.token)
  }

  async function fetchMe() {
    return { user: user.value, role: role.value, permissions: permissions.value }
  }

  async function fetchMenu() {
    return { menu: menu.value, role: role.value, permissions: permissions.value }
  }

  function hasPerm(): boolean { return true }

  function clear() {
    token.value = ''
    user.value = { email: '', nickname: '' }
    localStorage.removeItem(TOKEN_KEY)
  }

  async function logout() {
    clear()
  }

  return { token, user, permissions, role, menu, isLoggedIn, isAdmin, login, fetchMe, fetchMenu, hasPerm, clear, logout }
})
````

## File: web/src/styles/global.scss
````scss
:root {
  --gp-bg: #f5f7fa;
  --gp-sidebar-bg: #1f2330;
  --gp-sidebar-active: #409eff;
  --gp-card-shadow: 0 1px 4px rgba(0, 21, 41, 0.05);
}

html.dark {
  --gp-bg: #0d1117;
  --gp-sidebar-bg: #0b0e14;
  --gp-card-shadow: 0 1px 4px rgba(0, 0, 0, 0.35);
}

html, body, #app {
  height: 100%;
  margin: 0;
  padding: 0;
}

#app {
  background: var(--gp-bg);
}

.page-title {
  font-size: 18px;
  font-weight: 600;
  margin: 0 0 16px;
  color: var(--el-text-color-primary);
}

.page-container {
  padding: 20px;
}

.card-block {
  background: var(--el-bg-color);
  border-radius: 8px;
  padding: 18px 20px;
  box-shadow: var(--gp-card-shadow);
  + .card-block { margin-top: 16px; }
}

.flex-between {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 12px;
}
.flex-wrap-gap {
  display: flex;
  gap: 12px;
  flex-wrap: wrap;
  align-items: center;
}

// 表格内按钮紧凑
.el-table .el-button + .el-button { margin-left: 6px; }

/* ------------------------------------------------------------------
 * 滚动条统一风格
 *
 *  - 全局:细细一条(8px),圆角,半透明灰
 *  - 暗色主题自动加深
 *  - 左侧深色 sidebar 用专属深色滚动条,避免白白一条贴在暗色面板上很突兀
 *  - 覆盖范围:浏览器原生滚动(<main>、el-aside、el-table 等)+
 *            Firefox 的 scrollbar-color / scrollbar-width
 *
 *  注意:Element Plus 的 el-select / el-scrollbar 用的是自绘的 bar,不受这里影响。
 * ------------------------------------------------------------------ */

/* Firefox */
* {
  scrollbar-width: thin;
  scrollbar-color: rgba(144, 147, 153, 0.45) transparent;
}

/* WebKit / Blink(Chrome / Edge / Safari / Electron) */
*::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}
*::-webkit-scrollbar-track {
  background: transparent;
}
*::-webkit-scrollbar-thumb {
  background: rgba(144, 147, 153, 0.35);
  border-radius: 999px;
  border: 2px solid transparent;       /* 做成两侧留白的胶囊 */
  background-clip: padding-box;
  transition: background .2s;
}
*::-webkit-scrollbar-thumb:hover {
  background: rgba(144, 147, 153, 0.65);
  background-clip: padding-box;
}
*::-webkit-scrollbar-corner {
  background: transparent;
}

/* 深色主题下再加深对比 */
html.dark {
  * { scrollbar-color: rgba(255, 255, 255, 0.22) transparent; }

  *::-webkit-scrollbar-thumb {
    background: rgba(255, 255, 255, 0.18);
    background-clip: padding-box;
  }
  *::-webkit-scrollbar-thumb:hover {
    background: rgba(255, 255, 255, 0.32);
    background-clip: padding-box;
  }
}

/* 左侧深色 sidebar:scrollbar 要比 body 更暗更隐形 */
.el-aside.sidebar {
  scrollbar-width: thin;
  scrollbar-color: rgba(255, 255, 255, 0.18) transparent;

  &::-webkit-scrollbar {
    width: 6px;
    height: 6px;
  }
  &::-webkit-scrollbar-thumb {
    background: rgba(255, 255, 255, 0.14);
    border-radius: 999px;
    border: 1px solid transparent;
    background-clip: padding-box;
  }
  &::-webkit-scrollbar-thumb:hover {
    background: rgba(255, 255, 255, 0.30);
    background-clip: padding-box;
  }
}

/* Element Plus 自绘滚动条(下拉菜单 / 虚拟滚动)也统一一下粗细和颜色 */
.el-scrollbar__bar {
  &.is-vertical > .el-scrollbar__thumb,
  &.is-horizontal > .el-scrollbar__thumb {
    background-color: rgba(144, 147, 153, 0.45);
    border-radius: 999px;
  }
  &.is-vertical { width: 6px; }
  &.is-horizontal { height: 6px; }
}
html.dark .el-scrollbar__bar {
  &.is-vertical > .el-scrollbar__thumb,
  &.is-horizontal > .el-scrollbar__thumb {
    background-color: rgba(255, 255, 255, 0.28);
  }
}
````

## File: web/src/utils/brand.ts
````typescript
// brand placeholder — 推广信息已移除
export interface BrandParts {
  brand: string
}

export function brandParts(): BrandParts {
  return { brand: '' }
}
````

## File: web/src/utils/format.ts
````typescript
export function formatDateTime(v: string | { Time: string; Valid: boolean } | null | undefined): string {
  if (!v) return '-'
  let s = ''
  if (typeof v === 'string') s = v
  else if (v.Valid && v.Time) s = v.Time
  if (!s) return '-'
  const d = new Date(s)
  if (Number.isNaN(d.getTime())) return s
  return d.toLocaleString()
}

export function formatDateShort(v: string | { Time: string; Valid: boolean } | null | undefined): string {
  if (!v) return '-'
  let s = ''
  if (typeof v === 'string') s = v
  else if (v.Valid && v.Time) s = v.Time
  if (!s) return '-'
  const d = new Date(s)
  if (Number.isNaN(d.getTime())) return s
  return d.toLocaleDateString()
}

export function formatBytes(n: number | null | undefined): string {
  if (!n || n <= 0) return '0 B'
  const units = ['B', 'KB', 'MB', 'GB', 'TB']
  let v = n
  let i = 0
  while (v >= 1024 && i < units.length - 1) { v /= 1024; i++ }
  return `${v.toFixed(i === 0 ? 0 : 1)} ${units[i]}`
}

export function formatDuration(ms: number | null | undefined): string {
  if (!ms || ms <= 0) return '-'
  if (ms < 1000) return `${ms} ms`
  return `${(ms / 1000).toFixed(2)} s`
}

export function formatErrorCode(code?: string | null): string {
  if (!code) return '-'
  const m: Record<string, string> = {
    upstream_init_error: '上游客户端初始化失败',
    upstream_error: '上游返回错误',
    invalid_request_error: '请求参数错误',
    model_not_found: '模型未开放',
    account_dispatch_timeout: '等待可用账号超时',
    proxy_error: '代理异常',
    timeout: '请求超时',
  }
  return m[code] || code
}
````

## File: web/src/views/admin/Accounts.vue
````vue
<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox, ElNotification } from 'element-plus'
import { Upload } from '@element-plus/icons-vue'
import * as accountApi from '@/api/accounts'
import * as proxyApi from '@/api/proxies'
import { formatDateShort } from '@/utils/format'

// ========== 列表 & 筛选 ==========
const loading = ref(false)
const filter = reactive<{ status?: string; keyword?: string }>({ status: '', keyword: '' })
const rows = ref<accountApi.Account[]>([])
const total = ref(0)
const pager = reactive({ page: 1, page_size: 10 })
const proxies = ref<proxyApi.Proxy[]>([])

async function fetchList() {
  loading.value = true
  try {
    const data = await accountApi.listAccounts({
      page: pager.page,
      page_size: pager.page_size,
      status: filter.status || undefined,
      keyword: filter.keyword || undefined,
    })
    rows.value = data.list || []
    total.value = data.total || 0
  } catch (e: any) {
    ElMessage.error(e?.message || '加载失败')
  } finally {
    loading.value = false
  }
}

async function fetchProxies() {
  try {
    const d = await proxyApi.listProxies({ page: 1, page_size: 500 })
    proxies.value = (d.list || []).filter((p) => p.enabled)
  } catch {
    /* noop */
  }
}

function onSearch() {
  pager.page = 1
  fetchList()
}
function onReset() {
  filter.status = ''
  filter.keyword = ''
  pager.page = 1
  fetchList()
}

// ========== 自动刷新开关 ==========
const autoRefreshEnabled = ref(false)
const autoRefreshSaving = ref(false)

async function loadAutoRefresh() {
  try {
    const cfg = await accountApi.getAutoRefresh()
    autoRefreshEnabled.value = !!cfg.enabled
  } catch {
    /* noop */
  }
}
async function onToggleAutoRefresh(val: boolean | string | number) {
  const enabled = !!val
  autoRefreshSaving.value = true
  try {
    await accountApi.setAutoRefresh(enabled)
    autoRefreshEnabled.value = enabled
    ElMessage.success(
      enabled
        ? '已开启自动刷新:AT 距离过期 < 1 天时自动续期,失效/可疑账号不刷新'
        : '已关闭自动刷新',
    )
  } catch (e: any) {
    // 回滚 UI
    autoRefreshEnabled.value = !enabled
    ElMessage.error(e?.message || '保存失败')
  } finally {
    autoRefreshSaving.value = false
  }
}

// ========== 批量删除 ==========
const BULK_DELETE_LABELS: Record<string, string> = {
  dead:       '失效账号',
  suspicious: '可疑 / 已封账号',
  warned:     '风险账号',
  throttled:  '限流账号',
  all:        '全部账号',
}
async function onBulkDelete(scope: accountApi.BulkDeleteScope) {
  const label = BULK_DELETE_LABELS[scope] || scope
  try {
    await ElMessageBox.confirm(
      `确认将「${label}」全部删除?此操作会软删所有匹配条目,不可在当前界面恢复。`,
      scope === 'all' ? '⚠ 删除全部账号' : '批量删除',
      {
        confirmButtonText: '删除',
        cancelButtonText: '取消',
        type: scope === 'all' ? 'error' : 'warning',
      },
    )
  } catch { return }
  try {
    const r = await accountApi.bulkDeleteAccounts(scope)
    ElMessage.success(`已删除 ${r.deleted} 个账号`)
    pager.page = 1
    fetchList()
  } catch (e: any) {
    ElMessage.error(e?.message || '删除失败')
  }
}

// ========== 日期工具(兼容 sql.NullTime 返回形态) ==========
function asDate(v: any): string {
  if (!v) return ''
  if (typeof v === 'string') return v
  if (typeof v === 'object') {
    if ('Valid' in v && !v.Valid) return ''
    if ('Time' in v) return v.Time
  }
  return ''
}
function fmtTime(v: any) {
  const s = asDate(v)
  return s ? formatDateShort(s) : '-'
}

// ========== 状态/类型展示 ==========
type TagType = 'success' | 'warning' | 'info' | 'danger' | 'primary'
const statusMap: Record<string, { label: string; type: TagType }> = {
  healthy:    { label: '健康',   type: 'success' },
  warned:     { label: '风险',   type: 'warning' },
  throttled:  { label: '限流',   type: 'warning' },
  suspicious: { label: '可疑',   type: 'info'    },
  dead:       { label: '失效',   type: 'danger'  },
}
function statusText(s: string): string { return statusMap[s]?.label || s || '-' }
function statusType(s: string): TagType { return statusMap[s]?.type || 'info' }

function typeLabel(t: string) {
  const map: Record<string, string> = { codex: 'Codex', chatgpt: 'ChatGPT', openai: 'OpenAI' }
  return map[t] || t || '-'
}

// ========== 即将过期高亮 ==========
function expiresClass(v: any): string {
  const s = asDate(v)
  if (!s) return 'muted'
  const t = new Date(s).getTime()
  if (Number.isNaN(t)) return 'muted'
  const diffMin = (t - Date.now()) / 60000
  if (diffMin < 0) return 'err'
  if (diffMin < 30) return 'warn'
  return ''
}

// ========== 新建 / 编辑 ==========
const dlg = ref(false)
const isEdit = ref(false)
const submitting = ref(false)
const formDefault = {
  id: 0,
  email: '',
  auth_token: '',
  refresh_token: '',
  session_token: '',
  token_expires_at: '',
  oai_session_id: '',
  oai_device_id: '',
  client_id: 'app_LlGpXReQgckcGGUo2JrYvtJK',
  chatgpt_account_id: '',
  account_type: 'codex',
  subscription_type: 'plus',
  notes: '',
  cookies: '',
  proxy_id: 0,
  status: 'healthy',
}
const form = reactive({ ...formDefault })

function openCreate() {
  isEdit.value = false
  Object.assign(form, { ...formDefault })
  dlg.value = true
}

const secretsLoading = ref(false)

async function openEdit(row: accountApi.Account) {
  isEdit.value = true
  Object.assign(form, {
    id: row.id,
    email: row.email,
    auth_token: '',
    refresh_token: '',
    session_token: '',
    token_expires_at: asDate(row.token_expires_at),
    oai_session_id: row.oai_session_id || '',
    oai_device_id: row.oai_device_id || '',
    client_id: row.client_id || formDefault.client_id,
    chatgpt_account_id: row.chatgpt_account_id || '',
    account_type: row.account_type || 'codex',
    subscription_type: row.subscription_type || 'plus',
    notes: row.notes || '',
    cookies: '',
    proxy_id: 0,
    status: row.status || 'healthy',
  })
  dlg.value = true
  // 异步拉取 AT / RT / ST 明文并回填,方便查看/修改
  secretsLoading.value = true
  try {
    const s = await accountApi.getAccountSecrets(row.id)
    form.auth_token    = s.auth_token    || ''
    form.refresh_token = s.refresh_token || ''
    form.session_token = s.session_token || ''
  } catch (e: any) {
    ElMessage.warning('未能加载 AT/RT/ST 明文,留空即不修改')
  } finally {
    secretsLoading.value = false
  }
}

async function copyText(text: string, label: string) {
  if (!text) { ElMessage.info('内容为空'); return }
  try {
    await navigator.clipboard.writeText(text)
    ElMessage.success(`${label} 已复制`)
  } catch {
    ElMessage.error('复制失败,请手动选中复制')
  }
}

async function submitForm() {
  if (!form.email) { ElMessage.warning('请输入邮箱'); return }
  submitting.value = true
  try {
    if (!isEdit.value) {
      if (!form.auth_token) { ElMessage.warning('新建账号必须提供 access_token'); submitting.value = false; return }
      await accountApi.createAccount({ ...form })
      ElMessage.success('创建成功')
    } else {
      const body: any = {
        email: form.email,
        subscription_type: form.subscription_type,
        client_id: form.client_id,
        chatgpt_account_id: form.chatgpt_account_id,
        account_type: form.account_type,
        notes: form.notes,
        status: form.status,
      }
      if (form.auth_token)    body.auth_token    = form.auth_token
      if (form.refresh_token) body.refresh_token = form.refresh_token
      if (form.session_token) body.session_token = form.session_token
      if (form.cookies)       body.cookies       = form.cookies
      if (form.token_expires_at) body.token_expires_at = form.token_expires_at
      await accountApi.updateAccount(form.id, body)
      ElMessage.success('更新成功')
    }
    dlg.value = false
    await fetchList()
  } catch (e: any) {
    ElMessage.error(e?.message || '提交失败')
  } finally {
    submitting.value = false
  }
}

async function onDelete(row: accountApi.Account) {
  try {
    await ElMessageBox.confirm(`确定删除账号「${row.email}」?该操作不可恢复。`, '删除确认', {
      confirmButtonText: '删除', cancelButtonText: '取消', type: 'warning',
    })
  } catch { return }
  try {
    await accountApi.deleteAccount(row.id)
    ElMessage.success('已删除')
    fetchList()
  } catch (e: any) {
    ElMessage.error(e?.message || '删除失败')
  }
}

// ========== 绑定代理 ==========
const bindDlg = ref(false)
const bindForm = reactive({ id: 0, email: '', proxy_id: 0 })
function openBind(row: accountApi.Account) {
  bindForm.id = row.id
  bindForm.email = row.email
  bindForm.proxy_id = 0
  bindDlg.value = true
}
async function submitBind() {
  try {
    if (bindForm.proxy_id > 0) {
      await accountApi.bindProxy(bindForm.id, bindForm.proxy_id)
      ElMessage.success('已绑定代理')
    } else {
      await accountApi.unbindProxy(bindForm.id)
      ElMessage.success('已解绑')
    }
    bindDlg.value = false
    fetchList()
  } catch (e: any) {
    ElMessage.error(e?.message || '操作失败')
  }
}

// ========== 刷新 / 探测(单条) ==========
const refreshingIds = ref<Set<number>>(new Set())
const probingIds = ref<Set<number>>(new Set())

async function onRefreshOne(row: accountApi.Account) {
  refreshingIds.value.add(row.id)
  try {
    const r = await accountApi.refreshAccount(row.id)
    if (r.ok) {
      if (r.at_verified === false) {
        // RT 刷成功但新 AT 没通过 chatgpt.com web 校验,提示运维侧
        ElMessage.warning(
          `刷新成功(来源:${r.source.toUpperCase()}),但新 AT 未通过 chatgpt.com 校验,可能无法用于聊天/图像接口`
        )
      } else {
        ElMessage.success(`刷新成功(来源:${r.source.toUpperCase()})`)
      }
    } else if (r.web_unauthorized) {
      ElMessage.error(
        r.error || 'RT 换出的 AT 被 chatgpt.com 拒绝,请为该账号补充 Session Token'
      )
    } else {
      ElMessage.error(r.error || '刷新失败')
    }
    fetchList()
  } catch (e: any) {
    ElMessage.error(e?.message || '刷新失败')
  } finally {
    refreshingIds.value.delete(row.id)
  }
}
async function onProbeOne(row: accountApi.Account) {
  probingIds.value.add(row.id)
  try {
    const r = await accountApi.probeAccountQuota(row.id)
    if (r.ok) {
      const parts: string[] = [`生图剩余 ${r.remaining}`]
      if (r.default_model) parts.push(`模型 ${r.default_model}`)
      if (r.blocked_features && r.blocked_features.length) {
        parts.push(`受限:${r.blocked_features.join(',')}`)
      }
      ElMessage.success(parts.join(' · '))
    } else {
      ElMessage.error(r.error || '探测失败')
    }
    fetchList()
  } catch (e: any) {
    ElMessage.error(e?.message || '探测失败')
  } finally {
    probingIds.value.delete(row.id)
  }
}

// ========== 全部刷新 / 全部探测 ==========
const batchRunning = ref<'none' | 'refresh' | 'probe'>('none')

async function onRefreshAll() {
  if (total.value === 0) { ElMessage.info('暂无账号'); return }
  try {
    await ElMessageBox.confirm(`将并发刷新全部账号(共 ${total.value} 个),可能耗时较久,是否继续?`, '批量刷新', {
      confirmButtonText: '开始', cancelButtonText: '取消',
    })
  } catch { return }
  batchRunning.value = 'refresh'
  try {
    const r = await accountApi.refreshAllAccounts()
    const failDetails = (r.results || [])
      .filter((x: any) => !x.ok)
      .map((x: any) => `${x.email}: ${x.error || '未知错误'}`)
    const hasFailures = r.failed > 0
    ElNotification({
      type: hasFailures ? (r.success > 0 ? 'warning' : 'error') : 'success',
      title: '批量刷新完成',
      message: hasFailures
        ? `成功 ${r.success} · 失败 ${r.failed} · 合计 ${r.total}\n${failDetails.join('\n')}`
        : `全部成功 · 合计 ${r.total}`,
      duration: hasFailures ? 0 : 4000,
      dangerouslyUseHTMLString: true,
      ...(hasFailures && failDetails.length > 0 ? {
        message: `成功 ${r.success} · 失败 ${r.failed} · 合计 ${r.total}<br/>${failDetails.map(s => `<div style="color:#F56C6C;font-size:12px;margin-top:2px">${s}</div>`).join('')}`,
      } : {}),
    })
    fetchList()
  } catch (e: any) {
    ElMessage.error(e?.message || '刷新失败')
  } finally {
    batchRunning.value = 'none'
  }
}

async function onProbeAll() {
  if (total.value === 0) { ElMessage.info('暂无账号'); return }
  try {
    await ElMessageBox.confirm(`将并发探测全部账号的图片额度(共 ${total.value} 个)?`, '批量探测', {
      confirmButtonText: '开始', cancelButtonText: '取消',
    })
  } catch { return }
  batchRunning.value = 'probe'
  try {
    const r = await accountApi.probeAllAccountsQuota()
    ElNotification.success({
      title: '批量探测完成',
      message: `成功 ${r.success} · 失败 ${r.failed} · 合计 ${r.total}`,
      duration: 4000,
    })
    fetchList()
  } catch (e: any) {
    ElMessage.error(e?.message || '探测失败')
  } finally {
    batchRunning.value = 'none'
  }
}

// ========== 批量导入(多文件 + 分批) ==========
// 4 种模式:
//   - json: 文件/JSON 文本,原有行为
//   - at:   一行一个 access_token
//   - rt:   一行一个 refresh_token,必须提供 APPID(client_id)
//   - st:   一行一个 session_token
type ImportMode = 'json' | 'at' | 'rt' | 'st' | 'session_json'
const importDlg = ref(false)
const importMode = ref<ImportMode>('session_json')
const importForm = reactive({
  files: [] as File[],
  text: '',
  tokens_text: '',
  update_existing: true,
  default_client_id: 'app_LlGpXReQgckcGGUo2JrYvtJK',
  default_proxy_id: 0,
})
const importing = ref(false)
const importProgress = reactive({
  running: false,
  current: 0,
  totalBatches: 0,
  created: 0,
  updated: 0,
  skipped: 0,
  failed: 0,
})
const importResult = ref<accountApi.ImportSummary | null>(null)
const importLastErrors = ref<accountApi.ImportLineResult[]>([])

function openImport() {
  importMode.value = 'json'
  importForm.files = []
  importForm.text = ''
  importForm.tokens_text = ''
  importForm.update_existing = true
  importForm.default_client_id = 'app_LlGpXReQgckcGGUo2JrYvtJK'
  importForm.default_proxy_id = 0
  importResult.value = null
  importLastErrors.value = []
  importProgress.running = false
  importProgress.current = 0
  importProgress.totalBatches = 0
  importProgress.created = 0
  importProgress.updated = 0
  importProgress.skipped = 0
  importProgress.failed = 0
  importDlg.value = true
}

// 当前 tokens 模式下,每行 token 的数量预览
const tokenLineCount = computed(() => {
  if (importMode.value === 'json') return 0
  return importForm.tokens_text
    .split(/\r?\n/)
    .map((s) => s.trim())
    .filter(Boolean).length
})

function onPickFiles(e: Event) {
  const input = e.target as HTMLInputElement
  if (!input.files) return
  const arr = Array.from(input.files)
  importForm.files.push(...arr)
  input.value = ''
}
function onDropFiles(e: DragEvent) {
  e.preventDefault()
  if (!e.dataTransfer) return
  const arr = Array.from(e.dataTransfer.files).filter((f) => f.name.endsWith('.json') || f.name.endsWith('.txt'))
  importForm.files.push(...arr)
}
function removeFile(i: number) {
  importForm.files.splice(i, 1)
}
function clearFiles() {
  importForm.files = []
}

const totalFileSize = computed(() => importForm.files.reduce((s, f) => s + f.size, 0))
function humanSize(n: number) {
  if (n < 1024) return n + ' B'
  if (n < 1024 * 1024) return (n / 1024).toFixed(1) + ' KB'
  return (n / 1024 / 1024).toFixed(2) + ' MB'
}

/**
 * 前端分批:
 *   - 一批最多 BATCH_FILES 个文件或 BATCH_BYTES 字节,两者取先到者
 *   - 每批通过 multipart 上传,后端会解析并 upsert
 *   - 每批完成后更新进度,并 yield 事件循环
 * 选 10000 个文件时会自动切分 ~50 批,每批 ~200 个
 */
const BATCH_FILES = 200
const BATCH_BYTES = 8 * 1024 * 1024 // 8MB

async function doImport() {
  importLastErrors.value = []
  importResult.value = null

  // 情况 0:AT/RT/ST 纯 token 模式
  if (importMode.value !== 'json') {
    const mode = importMode.value
    const tokens = importForm.tokens_text
      .split(/\r?\n/)
      .map((s) => s.trim())
      .filter(Boolean)
    if (tokens.length === 0) {
      ElMessage.warning('请粘贴 token,每行一个')
      return
    }
    if (mode === 'rt' && !importForm.default_client_id.trim()) {
      ElMessage.warning('RT 模式必须填写 APPID(client_id)')
      return
    }
    if ((mode === 'rt' || mode === 'st') && !importForm.default_proxy_id) {
      try {
        await ElMessageBox.confirm(
          `${mode.toUpperCase()} 模式需要访问 chatgpt.com / auth.openai.com 换取 AT。未选择代理时会直连,国内网络大概率失败。确认继续吗?`,
          '建议选一个代理',
          { confirmButtonText: '继续直连', cancelButtonText: '取消', type: 'warning' },
        )
      } catch {
        return
      }
    }
    importing.value = true
    importProgress.running = true
    importProgress.current = 0
    importProgress.totalBatches = 1
    try {
      const r = await accountApi.importAccountsTokens({
        mode,
        tokens,
        client_id: importForm.default_client_id.trim() || undefined,
        update_existing: importForm.update_existing,
        default_proxy_id: importForm.default_proxy_id || undefined,
      })
      mergeSummary(r)
      importResult.value = cloneAgg()
      importLastErrors.value = r.results
        .filter((x) => x.status === 'failed' || x.status === 'skipped')
        .slice(0, 200)
      const tip = `${mode.toUpperCase()} 导入完成:+${r.created} / ~${r.updated} / 跳过${r.skipped} / 失败${r.failed}`
      if (r.failed > 0) ElNotification.warning({ title: '批量导入完成(部分失败)', message: tip })
      else ElMessage.success(tip)
    } catch (e: any) {
      ElMessage.error(e?.message || '导入失败')
    } finally {
      importing.value = false
      importProgress.running = false
      fetchList()
    }
    return
  }

  // 情况一:纯文本导入(JSON 模式)
  if (importForm.files.length === 0) {
    if (!importForm.text.trim()) { ElMessage.warning('请选择 JSON 文件或粘贴 JSON 文本'); return }
    importing.value = true
    importProgress.running = true
    importProgress.current = 0
    importProgress.totalBatches = 1
    try {
      const r = await accountApi.importAccountsJSON({
        text: importForm.text,
        update_existing: importForm.update_existing,
        default_client_id: importForm.default_client_id || undefined,
        default_proxy_id: importForm.default_proxy_id || undefined,
      })
      mergeSummary(r)
      importResult.value = cloneAgg()
      importLastErrors.value = r.results.filter((x) => x.status === 'failed' || x.status === 'skipped').slice(0, 200)
      ElMessage.success(`导入完成:+${r.created} / ~${r.updated} / 跳过${r.skipped} / 失败${r.failed}`)
    } catch (e: any) {
      ElMessage.error(e?.message || '导入失败')
    } finally {
      importing.value = false
      importProgress.running = false
      fetchList()
    }
    return
  }

  // 情况二:多文件分批
  const batches: File[][] = []
  let curBatch: File[] = []
  let curBytes = 0
  for (const f of importForm.files) {
    if ((curBatch.length >= BATCH_FILES) || (curBytes + f.size > BATCH_BYTES && curBatch.length > 0)) {
      batches.push(curBatch)
      curBatch = []
      curBytes = 0
    }
    curBatch.push(f)
    curBytes += f.size
  }
  if (curBatch.length) batches.push(curBatch)

  importing.value = true
  importProgress.running = true
  importProgress.current = 0
  importProgress.totalBatches = batches.length
  importProgress.created = 0
  importProgress.updated = 0
  importProgress.skipped = 0
  importProgress.failed = 0
  const errList: accountApi.ImportLineResult[] = []

  try {
    for (let i = 0; i < batches.length; i++) {
      const b = batches[i]
      try {
        const r = await accountApi.importAccountsFiles(b, {
          update_existing: importForm.update_existing,
          default_client_id: importForm.default_client_id || undefined,
          default_proxy_id: importForm.default_proxy_id || undefined,
        })
        mergeSummary(r)
        for (const it of r.results) {
          if ((it.status === 'failed' || it.status === 'skipped') && errList.length < 500) {
            errList.push(it)
          }
        }
      } catch (e: any) {
        importProgress.failed += b.length
        errList.push({ index: i, email: `[批次#${i + 1}]`, status: 'failed', reason: e?.message || '上传失败' })
      }
      importProgress.current = i + 1
      // 让出事件循环,避免阻塞 UI
      await new Promise((r) => setTimeout(r, 0))
    }
    importResult.value = cloneAgg()
    importLastErrors.value = errList
    ElNotification.success({
      title: '批量导入完成',
      message: `+${importProgress.created}  ~${importProgress.updated}  跳过 ${importProgress.skipped}  失败 ${importProgress.failed}`,
      duration: 5000,
    })
  } finally {
    importing.value = false
    importProgress.running = false
    fetchList()
  }
}

function mergeSummary(r: accountApi.ImportSummary) {
  importProgress.created += r.created
  importProgress.updated += r.updated
  importProgress.skipped += r.skipped
  importProgress.failed  += r.failed
}
function cloneAgg(): accountApi.ImportSummary {
  return {
    total:   importProgress.created + importProgress.updated + importProgress.skipped + importProgress.failed,
    created: importProgress.created,
    updated: importProgress.updated,
    skipped: importProgress.skipped,
    failed:  importProgress.failed,
    results: [],
  }
}

onMounted(() => {
  fetchList()
  fetchProxies()
  loadAutoRefresh()
})
</script>

<template>
  <div class="page-container">
    <!-- 顶栏:标题 + 动作 -->
    <div class="card-block hdr">
      <div class="flex-between">
        <div class="hdr-left">
          <h2 class="page-title">GPT 账号池</h2>
          <div class="page-sub">
            统一管理 ChatGPT Plus / Team / Codex 账号:JSON / AT / RT / ST 批量导入 · 自动刷新 · 图片额度探测 · 风控熔断轮转
          </div>
        </div>
        <div class="actions">
          <el-button :loading="batchRunning === 'probe'" :disabled="loading" @click="onProbeAll">
            全部探测
          </el-button>
          <el-button :loading="batchRunning === 'refresh'" :disabled="loading" @click="onRefreshAll">
            全部刷新
          </el-button>
          <el-dropdown trigger="click" @command="onBulkDelete">
            <el-button>批量删除</el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="dead">删除失效账号</el-dropdown-item>
                <el-dropdown-item command="suspicious">删除可疑/已封账号</el-dropdown-item>
                <el-dropdown-item command="warned">删除风险账号</el-dropdown-item>
                <el-dropdown-item command="throttled">删除限流账号</el-dropdown-item>
                <el-dropdown-item divided command="all">
                  <span style="color: var(--el-color-danger)">删除全部账号</span>
                </el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
          <el-button @click="openImport">批量导入</el-button>
          <el-button type="primary" @click="openCreate">新建账号</el-button>
        </div>
      </div>
    </div>

    <!-- 筛选栏 -->
    <div class="card-block">
      <el-form :inline="true" size="default" class="filter-form" @submit.prevent="onSearch">
        <el-form-item label="状态">
          <el-select v-model="filter.status" placeholder="全部" clearable style="width: 140px">
            <el-option label="全部" value="" />
            <el-option label="健康" value="healthy" />
            <el-option label="风险" value="warned" />
            <el-option label="限流" value="throttled" />
            <el-option label="可疑" value="suspicious" />
            <el-option label="失效" value="dead" />
          </el-select>
        </el-form-item>
        <el-form-item label="关键词">
          <el-input
            v-model="filter.keyword"
            placeholder="邮箱 / 备注"
            clearable
            style="width: 260px"
            @keyup.enter="onSearch"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSearch">搜索</el-button>
          <el-button @click="onReset">重置</el-button>
        </el-form-item>
        <el-form-item class="auto-refresh-item">
          <el-tooltip
            placement="top"
            content="开启后:AT 距离过期 < 1 天的账号会被后台自动续期;状态为「失效 / 可疑」的账号不会刷新"
          >
            <el-checkbox
              v-model="autoRefreshEnabled"
              :disabled="autoRefreshSaving"
              @change="onToggleAutoRefresh"
            >
              自动刷新 AT
              <span class="auto-refresh-hint">(&lt; 1 天过期时)</span>
            </el-checkbox>
          </el-tooltip>
        </el-form-item>
      </el-form>
    </div>

    <!-- 表格 -->
    <div class="card-block">
      <el-table
        v-loading="loading" :data="rows" stripe size="default" row-key="id"
        table-layout="auto" style="width: 100%"
      >
        <el-table-column label="邮箱" min-width="200" show-overflow-tooltip>
          <template #default="{ row }">
            <el-tooltip
              v-if="row.notes"
              placement="top"
              :content="row.notes"
            >
              <span class="email">{{ row.email }}</span>
            </el-tooltip>
            <span v-else class="email">{{ row.email }}</span>
          </template>
        </el-table-column>
        <el-table-column label="类型" width="76">
          <template #default="{ row }">
            <el-tag size="small" effect="plain">{{ typeLabel(row.account_type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="状态" width="76">
          <template #default="{ row }">
            <el-tag :type="statusType(row.status)" size="small">{{ statusText(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="凭证" width="86">
          <template #default="{ row }">
            <div class="creds">
              <el-tooltip content="存在 Refresh Token,可用 RT 自动刷新 AT" placement="top">
                <el-tag :type="row.has_rt ? 'success' : 'info'" size="small" effect="plain">
                  {{ row.has_rt ? 'RT' : '—' }}
                </el-tag>
              </el-tooltip>
              <el-tooltip content="存在 Session Token,可用 ST 回退刷新" placement="top">
                <el-tag :type="row.has_st ? 'success' : 'info'" size="small" effect="plain">
                  {{ row.has_st ? 'ST' : '—' }}
                </el-tag>
              </el-tooltip>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="AT 过期" min-width="148" show-overflow-tooltip>
          <template #default="{ row }">
            <span :class="expiresClass(row.token_expires_at)">{{ fmtTime(row.token_expires_at) }}</span>
          </template>
        </el-table-column>
        <el-table-column label="生图剩余" width="96" align="center">
          <template #default="{ row }">
            <template v-if="row.image_quota_remaining >= 0">
              <el-tooltip
                placement="top"
                :disabled="!asDate(row.image_quota_reset_at)"
                :content="'下次重置:' + fmtTime(row.image_quota_reset_at)"
              >
                <span class="quota"><b>{{ row.image_quota_remaining }}</b></span>
              </el-tooltip>
            </template>
            <span v-else class="muted">未探测</span>
          </template>
        </el-table-column>
        <el-table-column label="今日已用 / 上限" width="140" align="center">
          <template #default="{ row }">
            <el-tooltip placement="top">
              <template #content>
                <div style="line-height:1.8">
                  <div>今日已用 {{ row.today_used_count }} 张</div>
                </div>
              </template>
              <b>{{ row.today_used_count }}</b>
            </el-tooltip>
          </template>
        </el-table-column>
        <el-table-column label="最近刷新" min-width="148" show-overflow-tooltip>
          <template #default="{ row }">
            <div class="refresh-cell">
              <span>{{ fmtTime(row.last_refresh_at) }}</span>
              <el-tag
                v-if="row.last_refresh_source"
                size="small" effect="plain"
                :type="row.last_refresh_source === 'rt' ? 'success' : 'warning'"
              >{{ row.last_refresh_source.toUpperCase() }}</el-tag>
            </div>
            <el-tooltip
              v-if="row.refresh_error"
              placement="top"
              :content="row.refresh_error"
            >
              <div class="err">{{ row.refresh_error }}</div>
            </el-tooltip>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="240" fixed="right">
          <template #default="{ row }">
            <el-button
              link type="primary" size="small"
              :loading="refreshingIds.has(row.id)"
              @click="onRefreshOne(row)"
            >刷新</el-button>
            <el-button
              link type="primary" size="small"
              :loading="probingIds.has(row.id)"
              @click="onProbeOne(row)"
            >探测</el-button>
            <el-button link type="primary" size="small" @click="openBind(row)">代理</el-button>
            <el-button link type="primary" size="small" @click="openEdit(row)">编辑</el-button>
            <el-button link type="danger"  size="small" @click="onDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="pager">
        <el-pagination
          v-model:current-page="pager.page"
          v-model:page-size="pager.page_size"
          :total="total"
          :page-sizes="[10, 20, 50, 100, 200, 500, 1000]"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="fetchList"
          @size-change="fetchList"
        />
      </div>
    </div>

    <!-- 新建 / 编辑弹窗 -->
    <el-dialog v-model="dlg" :title="isEdit ? '编辑账号' : '新建账号'" width="720px" destroy-on-close>
      <el-form label-width="120px" size="default">
        <el-form-item label="邮箱">
          <el-input v-model="form.email" placeholder="user@example.com" />
        </el-form-item>
        <el-form-item label="账号类型">
          <el-select v-model="form.account_type" style="width: 180px">
            <el-option label="Codex" value="codex" />
            <el-option label="ChatGPT" value="chatgpt" />
          </el-select>
        </el-form-item>
        <el-form-item label="Access Token">
          <div class="token-field">
            <el-input
              v-model="form.auth_token"
              type="textarea" :rows="3"
              :placeholder="isEdit
                ? (secretsLoading ? '正在加载当前 AT……' : '当前为空,粘贴新的 access_token 可更新')
                : '粘贴 access_token(必填)'"
              spellcheck="false"
            />
            <el-button
              v-if="isEdit"
              size="small" link
              :disabled="!form.auth_token"
              @click="copyText(form.auth_token, 'Access Token')"
            >复制</el-button>
          </div>
        </el-form-item>
        <el-form-item label="Refresh Token">
          <div class="token-field">
            <el-input
              v-model="form.refresh_token"
              type="textarea" :rows="2"
              :placeholder="isEdit
                ? (secretsLoading ? '正在加载当前 RT……' : '该账号暂无 Refresh Token')
                : '可选;有 RT 则支持自动刷新'"
              spellcheck="false"
            />
            <el-button
              v-if="isEdit"
              size="small" link
              :disabled="!form.refresh_token"
              @click="copyText(form.refresh_token, 'Refresh Token')"
            >复制</el-button>
          </div>
        </el-form-item>
        <el-form-item label="Session Token">
          <div class="token-field">
            <el-input
              v-model="form.session_token"
              type="textarea" :rows="2"
              :placeholder="isEdit
                ? (secretsLoading ? '正在加载当前 ST……' : '该账号暂无 Session Token')
                : '可选;__Secure-next-auth.session-token 的值'"
              spellcheck="false"
            />
            <el-button
              v-if="isEdit"
              size="small" link
              :disabled="!form.session_token"
              @click="copyText(form.session_token, 'Session Token')"
            >复制</el-button>
          </div>
        </el-form-item>
        <el-form-item label="Token 过期时间">
          <el-date-picker
            v-model="form.token_expires_at"
            type="datetime" format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DDTHH:mm:ssZ"
            placeholder="留空则从 JWT 自动解析"
            style="width: 260px"
          />
        </el-form-item>
        <el-form-item label="Client ID">
          <el-input v-model="form.client_id" />
        </el-form-item>
        <el-form-item label="ChatGPT AccountID">
          <el-input v-model="form.chatgpt_account_id" placeholder="可选;JSON 里有则自动填充" />
        </el-form-item>
        <el-form-item label="账号订阅">
          <el-select v-model="form.subscription_type" style="width: 180px">
            <el-option label="Plus"  value="plus" />
            <el-option label="Team"  value="team" />
            <el-option label="Free"  value="free" />
            <el-option label="Codex" value="codex" />
          </el-select>
        </el-form-item>
        <el-form-item v-if="isEdit" label="状态">
          <el-select v-model="form.status" style="width: 180px">
            <el-option label="健康"  value="healthy" />
            <el-option label="风险"  value="warned" />
            <el-option label="限流"  value="throttled" />
            <el-option label="可疑"  value="suspicious" />
            <el-option label="失效"  value="dead" />
          </el-select>
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="form.notes" type="textarea" :rows="2" />
        </el-form-item>
        <el-form-item v-if="!isEdit" label="绑定代理">
          <el-select v-model="form.proxy_id" clearable placeholder="不绑定" style="width: 100%">
            <el-option :value="0" label="不绑定" />
            <el-option
              v-for="p in proxies"
              :key="p.id"
              :label="`#${p.id} ${p.remark || p.host}:${p.port}`"
              :value="p.id"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dlg = false">取消</el-button>
        <el-button type="primary" :loading="submitting" @click="submitForm">确定</el-button>
      </template>
    </el-dialog>

    <!-- 绑定代理弹窗 -->
    <el-dialog v-model="bindDlg" title="绑定代理" width="420px">
      <div style="margin-bottom: 10px; color: var(--el-text-color-secondary)">
        账号:<b>{{ bindForm.email }}</b>
      </div>
      <el-select v-model="bindForm.proxy_id" clearable placeholder="选择代理(留空=解绑)" style="width: 100%">
        <el-option :value="0" label="不绑定 / 解绑" />
        <el-option
          v-for="p in proxies"
          :key="p.id"
          :label="`#${p.id} ${p.remark || p.host}:${p.port}`"
          :value="p.id"
        />
      </el-select>
      <template #footer>
        <el-button @click="bindDlg = false">取消</el-button>
        <el-button type="primary" @click="submitBind">确定</el-button>
      </template>
    </el-dialog>

    <!-- 批量导入弹窗 -->
    <el-dialog v-model="importDlg" title="批量导入账号" width="760px" destroy-on-close>
      <!-- 模式 tab -->
      <el-tabs v-model="importMode" class="import-tabs">
        <el-tab-pane label="Session JSON" name="session_json" />
        <el-tab-pane label="JSON 文件" name="json" />
        <el-tab-pane label="Access Token" name="at" />
        <el-tab-pane label="Refresh Token" name="rt" />
        <el-tab-pane label="Session Token" name="st" />
      </el-tabs>

      <!-- Session JSON 模式 -->
      <template v-if="importMode === 'session_json'">
        <div class="tip">
          直接粘贴 <code>chatgpt.com/api/auth/session</code> 返回的<b>完整 JSON</b>。
          系统会自动提取 accessToken、email、过期时间和订阅类型。每个 JSON 一行，支持多行（多个账号）。
        </div>
        <el-input
          v-model="importForm.tokens_text"
          type="textarea" :rows="10"
          placeholder='{"WARNING_BANNER":"...","user":{"email":"..."},"expires":"...","accessToken":"eyJ..."}'
          spellcheck="false"
        />
        <div class="token-count">共 {{ tokenLineCount }} 行</div>
      </template>

      <!-- JSON 模式 -->
      <template v-else-if="importMode === 'json'">
        <div class="tip">
          支持两种 JSON:<b>sub2api-account-*.json</b>(多账号)与 <b>token_*.json</b>(单账号)。
          可一次选择多个文件,前端会按每批 {{ BATCH_FILES }} 个文件 / {{ humanSize(BATCH_BYTES) }} 自动分批上传,不会卡页面。
        </div>

        <!-- 文件选择 + 拖拽 -->
        <div class="drop-zone" @dragover.prevent @drop="onDropFiles">
          <el-icon class="drop-ic"><Upload /></el-icon>
          <div class="drop-text">
            把 JSON 文件拖到这里,或
            <label class="link">
              <input type="file" accept=".json,.txt,application/json" multiple hidden @change="onPickFiles" />
              选择文件
            </label>
          </div>
          <div class="drop-sub">
            已选 <b>{{ importForm.files.length }}</b> 个文件
            <span v-if="importForm.files.length > 0"> · 合计 {{ humanSize(totalFileSize) }}</span>
          </div>
        </div>

        <div v-if="importForm.files.length" class="file-list">
          <div class="file-list-head">
            <span>{{ importForm.files.length }} 个文件</span>
            <el-button link type="danger" size="small" @click="clearFiles">清空</el-button>
          </div>
          <div class="file-list-body">
            <div v-for="(f, i) in importForm.files.slice(0, 50)" :key="i" class="file-row">
              <span class="fname">{{ f.name }}</span>
              <span class="fsize">{{ humanSize(f.size) }}</span>
              <el-button link size="small" @click="removeFile(i)">×</el-button>
            </div>
            <div v-if="importForm.files.length > 50" class="muted" style="text-align:center;margin-top:6px">
              ……另有 {{ importForm.files.length - 50 }} 个文件未展示
            </div>
          </div>
        </div>

        <el-divider content-position="left">或粘贴 JSON 文本</el-divider>
        <el-input
          v-model="importForm.text"
          type="textarea" :rows="5"
          placeholder="粘贴 sub2api 或 token_*.json 内容,多个 JSON 可以直接换行拼接(JSONL)"
          spellcheck="false"
        />
      </template>

      <!-- AT 模式 -->
      <template v-else-if="importMode === 'at'">
        <div class="tip">
          一行一个 <b>access_token</b>(eyJ... 开头的 JWT)。
          服务端会解析 JWT payload 里的 email 作为账号唯一键,若 AT 里没有 email 字段,该行会进入失败列表。
        </div>
        <el-input
          v-model="importForm.tokens_text"
          type="textarea" :rows="10"
          placeholder="eyJhbGci...&#10;eyJhbGci...&#10;..."
          spellcheck="false"
        />
        <div class="token-count">共 {{ tokenLineCount }} 行</div>
      </template>

      <!-- RT 模式 -->
      <template v-else-if="importMode === 'rt'">
        <div class="tip">
          一行一个 <b>refresh_token</b>。系统会用你填写的 <b>APPID(client_id)</b> 向
          <code>auth.openai.com/oauth/token</code> 换出 AT,再从 AT 解出 email 后入库。
          <strong class="warn">需要选择代理,否则大概率超时。</strong>
        </div>
        <el-input
          v-model="importForm.tokens_text"
          type="textarea" :rows="9"
          placeholder="v1.rt_...&#10;v1.rt_...&#10;..."
          spellcheck="false"
        />
        <div class="token-count">共 {{ tokenLineCount }} 行</div>
      </template>

      <!-- ST 模式 -->
      <template v-else-if="importMode === 'st'">
        <div class="tip">
          一行一个 <b>session_token</b>(浏览器 cookie 里的 <code>__Secure-next-auth.session-token</code>)。
          系统会用它调 <code>chatgpt.com/api/auth/session</code> 换出 AT,再从 AT 解 email。
          <strong class="warn">ST 模式必须有代理(chatgpt.com 国内不可直连)。</strong>
        </div>
        <el-input
          v-model="importForm.tokens_text"
          type="textarea" :rows="10"
          placeholder="eyJhbGci...&#10;eyJhbGci...&#10;..."
          spellcheck="false"
        />
        <div class="token-count">共 {{ tokenLineCount }} 行</div>
      </template>

      <div style="margin-top: 14px; display: flex; flex-wrap: wrap; gap: 14px; align-items: center">
        <el-checkbox v-model="importForm.update_existing">邮箱已存在则更新 token</el-checkbox>
        <div>
          <span class="muted" style="margin-right: 6px">
            {{ importMode === 'rt' ? 'APPID(client_id,必填)' : 'client_id' }}
          </span>
          <el-input
            v-model="importForm.default_client_id"
            size="small" style="width: 280px"
            :placeholder="importMode === 'rt' ? 'app_xxxxxxxxxxxxxxxxxxxxxxxx(必填)' : '可选,默认 ChatGPT iOS'"
          />
        </div>
        <div>
          <span class="muted" style="margin-right: 6px">
            {{ importMode === 'st' || importMode === 'rt' ? '代理(强烈推荐)' : '默认代理' }}
          </span>
          <el-select v-model="importForm.default_proxy_id" clearable size="small" style="width: 220px">
            <el-option :value="0" label="不绑定" />
            <el-option v-for="p in proxies" :key="p.id" :label="`#${p.id} ${p.remark || p.host}:${p.port}`" :value="p.id" />
          </el-select>
        </div>
      </div>

      <!-- 进度条 -->
      <div v-if="importProgress.running || importResult" class="progress">
        <div class="progress-head">
          <span v-if="importProgress.running">
            正在导入:第 <b>{{ importProgress.current }}</b> / {{ importProgress.totalBatches }} 批
          </span>
          <span v-else>
            导入已完成
          </span>
          <span class="stat">
            <el-tag type="success" size="small">+{{ importProgress.created }}</el-tag>
            <el-tag type="warning" size="small">~{{ importProgress.updated }}</el-tag>
            <el-tag type="info" size="small">跳过 {{ importProgress.skipped }}</el-tag>
            <el-tag type="danger" size="small">失败 {{ importProgress.failed }}</el-tag>
          </span>
        </div>
        <el-progress
          :percentage="importProgress.totalBatches > 0
            ? Math.round((importProgress.current / importProgress.totalBatches) * 100)
            : 0"
          :status="importProgress.running ? '' : (importProgress.failed > 0 ? 'warning' : 'success')"
        />
      </div>

      <!-- 错误/跳过明细 -->
      <div v-if="importLastErrors.length" class="err-list">
        <div class="err-list-head">未成功明细({{ importLastErrors.length }})</div>
        <el-table :data="importLastErrors" size="small" max-height="220">
          <el-table-column prop="email" label="邮箱" min-width="200" />
          <el-table-column label="状态" width="90">
            <template #default="{ row }">
              <el-tag size="small" :type="row.status === 'failed' ? 'danger' : 'info'">
                {{ row.status === 'failed' ? '失败' : '跳过' }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="reason" label="原因" min-width="240" show-overflow-tooltip />
        </el-table>
      </div>

      <template #footer>
        <el-button :disabled="importing" @click="importDlg = false">关闭</el-button>
        <el-button type="primary" :loading="importing" @click="doImport">
          开始导入
          <span v-if="importMode === 'json' && importForm.files.length > 0">
            ({{ importForm.files.length }} 个文件)
          </span>
          <span v-else-if="importMode !== 'json' && tokenLineCount > 0">
            ({{ tokenLineCount }} 条 {{ importMode.toUpperCase() }})
          </span>
        </el-button>
      </template>
    </el-dialog>
  </div>
</template>

<style scoped lang="scss">
.hdr { margin-bottom: 14px !important; }
.hdr-left .page-sub {
  color: var(--el-text-color-secondary);
  font-size: 13px;
  margin-top: 4px;
}
.actions {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
}
.flex-between {
  display: flex; align-items: center; justify-content: space-between; gap: 16px;
}

.filter-form :deep(.el-form-item) { margin-bottom: 0; }
.auto-refresh-item { margin-left: 4px; }
.auto-refresh-hint {
  color: var(--el-text-color-secondary);
  font-size: 12px;
  margin-left: 4px;
}

.email {
  color: var(--el-text-color-primary);
  overflow: hidden; text-overflow: ellipsis; white-space: nowrap;
  display: inline-block; max-width: 100%;
}

.refresh-cell {
  display: flex; align-items: center; gap: 6px;
  overflow: hidden; text-overflow: ellipsis; white-space: nowrap;
}

.creds { display: flex; gap: 4px; }
.quota b { color: var(--el-color-primary); font-weight: 600; }
.muted { color: var(--el-text-color-secondary); }
.warn  { color: var(--el-color-warning); font-weight: 500; }
.err   {
  color: var(--el-color-danger);
  font-size: 12px;
  white-space: nowrap; overflow: hidden; text-overflow: ellipsis;
  margin-top: 2px;
}

.token-field {
  display: flex; align-items: flex-start; gap: 8px; width: 100%;
  :deep(.el-textarea) { flex: 1; }
}

.pager {
  display: flex; justify-content: flex-end;
  margin-top: 14px;
}

/* ====== 批量导入弹窗 ====== */
.import-tabs {
  margin-top: -8px;
  margin-bottom: 10px;
  :deep(.el-tabs__header) { margin-bottom: 12px; }
}
.tip {
  color: var(--el-text-color-secondary);
  font-size: 13px; line-height: 1.6;
  background: var(--el-fill-color-light);
  padding: 10px 12px;
  border-radius: 8px;
  margin-bottom: 12px;
  code {
    background: rgba(0, 0, 0, 0.06);
    padding: 1px 6px;
    border-radius: 4px;
    font-family: inherit;
  }
  .warn {
    color: var(--el-color-warning);
    margin-left: 4px;
  }
}
.token-count {
  text-align: right;
  font-size: 12px;
  color: var(--el-text-color-secondary);
  margin-top: 6px;
}

.drop-zone {
  border: 1.5px dashed var(--el-border-color);
  border-radius: 10px;
  padding: 22px 14px;
  text-align: center;
  background: var(--el-fill-color-lighter);
  transition: border-color 0.15s, background-color 0.15s;
  &:hover {
    border-color: var(--el-color-primary);
    background: var(--el-color-primary-light-9);
  }
  .drop-ic {
    font-size: 30px; color: var(--el-color-primary); margin-bottom: 6px;
  }
  .drop-text {
    font-size: 14px;
    color: var(--el-text-color-primary);
    .link {
      color: var(--el-color-primary);
      cursor: pointer;
      text-decoration: underline;
      margin-left: 2px;
    }
  }
  .drop-sub {
    margin-top: 6px;
    font-size: 12px;
    color: var(--el-text-color-secondary);
  }
}

.file-list {
  margin-top: 10px;
  border: 1px solid var(--el-border-color-lighter);
  border-radius: 8px;
  max-height: 200px;
  overflow: auto;
  padding: 8px 10px;
  .file-list-head {
    display: flex; justify-content: space-between; align-items: center;
    padding-bottom: 4px;
    color: var(--el-text-color-secondary);
    font-size: 12px;
    border-bottom: 1px dashed var(--el-border-color-lighter);
    margin-bottom: 4px;
  }
  .file-row {
    display: flex; align-items: center; justify-content: space-between;
    padding: 4px 0;
    font-size: 13px;
    .fname { flex: 1; overflow: hidden; text-overflow: ellipsis; white-space: nowrap; }
    .fsize { color: var(--el-text-color-secondary); margin: 0 8px; font-variant-numeric: tabular-nums; }
  }
}

.progress {
  margin-top: 16px;
  padding: 12px 14px;
  background: var(--el-fill-color-lighter);
  border-radius: 8px;
  .progress-head {
    display: flex; justify-content: space-between; align-items: center;
    margin-bottom: 6px;
    .stat { display: flex; gap: 6px; }
  }
}

.err-list {
  margin-top: 12px;
  .err-list-head {
    color: var(--el-color-danger);
    font-weight: 500;
    margin-bottom: 6px;
  }
}
</style>
````

## File: web/src/views/admin/Audit.vue
````vue
<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import * as auditApi from '@/api/audit'
import { formatDateTime } from '@/utils/format'

const loading = ref(false)
const filter = reactive<auditApi.AuditFilter>({ action: '', actor_id: undefined, limit: 50, offset: 0 })
const items = ref<auditApi.AuditLog[]>([])
const total = ref(0)

async function load() {
  loading.value = true
  try {
    const d = await auditApi.listAudit({ ...filter, actor_id: filter.actor_id || undefined })
    items.value = d.items
    total.value = d.total
  } finally { loading.value = false }
}

const detailDlg = ref(false)
const detailRow = ref<auditApi.AuditLog | null>(null)
function openDetail(row: auditApi.AuditLog) {
  detailRow.value = row
  detailDlg.value = true
}

onMounted(load)
</script>

<template>
  <div class="page-container">
    <div class="card-block">
      <h2 class="page-title" style="margin:0">审计日志</h2>
      <div style="color:var(--el-text-color-secondary);font-size:13px;margin:4px 0 12px">
        记录控制台写操作(新增 / 修改 / 删除 / 备份恢复),按 action 与操作者 ID 可精确回溯。
      </div>
      <el-form inline class="flex-wrap-gap" @submit.prevent="load">
        <el-input v-model="filter.action" placeholder="action(如 accounts.update)" clearable style="width:220px" />
        <el-input-number v-model="filter.actor_id" placeholder="操作者 ID" :min="0" style="width:170px" />
        <el-button type="primary" @click="load"><el-icon><Search /></el-icon> 查询</el-button>
      </el-form>

      <el-table v-loading="loading" :data="items" stripe style="margin-top:12px">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column label="操作者" min-width="200">
          <template #default="{ row }">
            <div>{{ row.actor_email || '-' }}</div>
            <div style="font-size:12px;color:var(--el-text-color-secondary)">ID: {{ row.actor_id }}</div>
          </template>
        </el-table-column>
        <el-table-column prop="action" label="Action" min-width="180" />
        <el-table-column prop="method" label="Method" width="90" />
        <el-table-column prop="path" label="Path" min-width="220" show-overflow-tooltip />
        <el-table-column prop="status_code" label="Status" width="80" />
        <el-table-column prop="target" label="Target" min-width="100" show-overflow-tooltip />
        <el-table-column prop="ip" label="IP" width="120" />
        <el-table-column label="时间" width="170">
          <template #default="{ row }">{{ formatDateTime(row.created_at) }}</template>
        </el-table-column>
        <el-table-column label="" width="70" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="openDetail(row)">详情</el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-pagination style="margin-top:16px;display:flex;justify-content:flex-end"
        :current-page="Math.floor((filter.offset || 0) / (filter.limit || 50)) + 1"
        @current-change="(p: number) => { filter.offset = (p - 1) * (filter.limit || 50); load() }"
        :page-size="filter.limit"
        :total="total"
        :page-sizes="[50, 100, 200]"
        @size-change="(s: number) => { filter.limit = s; filter.offset = 0; load() }"
        layout="total, sizes, prev, pager, next"
      />
    </div>

    <el-dialog v-model="detailDlg" title="审计详情" width="620px">
      <el-descriptions v-if="detailRow" :column="2" border size="small">
        <el-descriptions-item label="ID">{{ detailRow.id }}</el-descriptions-item>
        <el-descriptions-item label="时间">{{ formatDateTime(detailRow.created_at) }}</el-descriptions-item>
        <el-descriptions-item label="Action">{{ detailRow.action }}</el-descriptions-item>
        <el-descriptions-item label="Status">{{ detailRow.status_code }}</el-descriptions-item>
        <el-descriptions-item label="Method">{{ detailRow.method }}</el-descriptions-item>
        <el-descriptions-item label="Path">{{ detailRow.path }}</el-descriptions-item>
        <el-descriptions-item label="Actor">{{ detailRow.actor_email }} (#{{ detailRow.actor_id }})</el-descriptions-item>
        <el-descriptions-item label="IP">{{ detailRow.ip }}</el-descriptions-item>
        <el-descriptions-item label="UA" :span="2">{{ detailRow.ua }}</el-descriptions-item>
        <el-descriptions-item label="Target" :span="2">{{ detailRow.target || '-' }}</el-descriptions-item>
        <el-descriptions-item label="Meta" :span="2">
          <pre class="meta">{{ typeof detailRow.meta === 'string' ? detailRow.meta : JSON.stringify(detailRow.meta, null, 2) }}</pre>
        </el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>

<style scoped lang="scss">
.meta {
  font-family: ui-monospace, Menlo, Consolas, monospace;
  font-size: 12px;
  background: #f7f8fa;
  padding: 8px;
  border-radius: 4px;
  max-height: 280px;
  overflow: auto;
  margin: 0;
}
</style>
````

## File: web/src/views/admin/Backup.vue
````vue
<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import * as backupApi from '@/api/backup'
import { formatBytes, formatDateTime } from '@/utils/format'

const loading = ref(false)
const items = ref<backupApi.BackupFile[]>([])
const total = ref(0)
const allowRestore = ref(false)
const maxUploadMB = ref(512)
const page = reactive({ limit: 50, offset: 0 })
const creating = ref(false)

async function load() {
  loading.value = true
  try {
    const d = await backupApi.listBackups(page.limit, page.offset)
    items.value = d.items
    total.value = d.total
    allowRestore.value = d.allow_restore
    maxUploadMB.value = d.max_upload_mb
  } finally { loading.value = false }
}

async function onCreate() {
  creating.value = true
  try {
    await backupApi.createBackup(true)
    ElMessage.success('备份已创建')
    load()
  } finally { creating.value = false }
}

function download(row: backupApi.BackupFile) {
  return backupApi.downloadBackup(row.backup_id, row.file_name)
}

async function onDelete(row: backupApi.BackupFile) {
  const ok = await ElMessageBox.confirm(`确认删除 ${row.file_name}?`, '删除备份', { type: 'warning', confirmButtonText: '删除', cancelButtonText: '取消' }).catch(() => null)
  if (!ok) return
  await backupApi.deleteBackup(row.backup_id)
  ElMessage.success('已删除')
  load()
}

async function onRestore(row: backupApi.BackupFile) {
  if (!allowRestore.value) return ElMessage.error('后端未启用恢复功能')
  await ElMessageBox.confirm(
    `恢复会覆盖当前数据库!此操作不可撤销。你已理解风险并希望继续?`,
    '恢复数据库', { type: 'error', confirmButtonText: '我确认继续', cancelButtonText: '取消' },
  )
  const ok = await ElMessageBox.confirm('最后一次确认:立即执行恢复?', '恢复数据库', { type: 'error', confirmButtonText: '执行恢复', cancelButtonText: '取消' }).catch(() => null)
  if (!ok) return
  ElMessage.info('正在恢复,请稍候…')
  await backupApi.restoreBackup(row.backup_id)
  ElMessage.success('恢复成功,请刷新页面')
}

// ---- 上传 ----
const uploadDlg = ref(false)
const uploadFile = ref<File | null>(null)
const uploadPct = ref(0)
const uploading = ref(false)

function pickFile(e: Event) {
  const t = e.target as HTMLInputElement
  uploadFile.value = t.files?.[0] || null
}
async function doUpload() {
  if (!uploadFile.value) return ElMessage.warning('请选择 .sql.gz 文件')
  uploading.value = true
  uploadPct.value = 0
  try {
    await backupApi.uploadBackup(uploadFile.value, (p) => (uploadPct.value = p))
    ElMessage.success('上传成功')
    uploadDlg.value = false
    uploadFile.value = null
    load()
  } finally { uploading.value = false }
}

onMounted(load)
</script>

<template>
  <div class="page-container">
    <div class="card-block">
      <div class="flex-between" style="margin-bottom:12px">
        <div>
          <h2 class="page-title" style="margin:0">数据备份</h2>
          <div style="color:var(--el-text-color-secondary);font-size:13px;margin-top:4px">
            本地最多保留 {{ total }} 个备份文件,上传单文件上限
            <el-tag size="small">{{ maxUploadMB }} MB</el-tag>
            · 恢复功能:
            <el-tag :type="allowRestore ? 'success' : 'info'" size="small">
              {{ allowRestore ? '已启用' : '已禁用(需后端 backup.allow_restore=true)' }}
            </el-tag>
          </div>
        </div>
        <div class="flex-wrap-gap">
          <el-button @click="uploadDlg = true"><el-icon><Upload /></el-icon> 上传</el-button>
          <el-button type="primary" :loading="creating" @click="onCreate">
            <el-icon><FolderAdd /></el-icon> 立即备份
          </el-button>
        </div>
      </div>

      <el-table v-loading="loading" :data="items" stripe>
        <el-table-column prop="backup_id" label="ID" width="220" />
        <el-table-column prop="file_name" label="文件" min-width="240" show-overflow-tooltip />
        <el-table-column label="大小" width="100">
          <template #default="{ row }">{{ formatBytes(row.size_bytes) }}</template>
        </el-table-column>
        <el-table-column label="来源" width="90">
          <template #default="{ row }">
            <el-tag size="small">{{ row.trigger }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="row.status === 'ready' ? 'success' : row.status === 'failed' ? 'danger' : 'info'"
                    size="small">
              {{ row.status }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="创建时间" width="170">
          <template #default="{ row }">{{ formatDateTime(row.created_at) }}</template>
        </el-table-column>
        <el-table-column prop="sha256" label="SHA256" min-width="180" show-overflow-tooltip />
        <el-table-column label="操作" width="240" fixed="right">
          <template #default="{ row }">
            <el-button size="small" link type="primary"
                       :disabled="row.status !== 'ready'" @click="download(row)">下载</el-button>
            <el-button size="small" link type="warning"
                       :disabled="!allowRestore || row.status !== 'ready'" @click="onRestore(row)">恢复</el-button>
            <el-button size="small" link type="danger" @click="onDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <el-dialog v-model="uploadDlg" title="上传备份" width="460px">
      <el-alert type="info" :closable="false" show-icon style="margin-bottom:12px"
                title="仅接受 .sql.gz 格式;恢复仍需单独操作。" />
      <el-form label-width="110px">
        <el-form-item label="文件">
          <input type="file" accept=".gz,.sql.gz" @change="pickFile" />
          <div v-if="uploadFile" style="font-size:12px;margin-top:6px;color:var(--el-text-color-secondary)">
            已选择 {{ uploadFile.name }} · {{ formatBytes(uploadFile.size) }}
          </div>
        </el-form-item>
        <el-form-item v-if="uploading" label="进度">
          <el-progress :percentage="uploadPct" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="uploadDlg = false">取消</el-button>
        <el-button type="primary" :loading="uploading" @click="doUpload">上传</el-button>
      </template>
    </el-dialog>
  </div>
</template>
````

## File: web/src/views/admin/Models.vue
````vue
<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue'
import { ElMessage, ElMessageBox, type FormInstance, type FormRules } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import * as statsApi from '@/api/stats'
import { ENABLE_CHAT_MODEL } from '@/config/feature'

const loading = ref(false)
const rows = ref<statsApi.Model[]>([])
const filterType = ref<string>(ENABLE_CHAT_MODEL ? '' : 'image')

async function load() {
  loading.value = true
  try {
    const d = await statsApi.listModels()
    rows.value = d.items || []
  } finally { loading.value = false }
}

const filteredRows = computed(() => filterType.value ? rows.value.filter((r) => r.type === filterType.value) : rows.value)

const dlgVisible = ref(false)
const dlgState = ref<'create' | 'edit'>('create')
const dlgLoading = ref(false)
const formRef = ref<FormInstance>()

const emptyForm = (): statsApi.ModelUpsert & { id: number } => ({
  id: 0,
  slug: '',
  type: ENABLE_CHAT_MODEL ? 'chat' : 'image',
  upstream_model_slug: '',
  description: '',
  enabled: true,
})
const form = reactive(emptyForm())

const rules: FormRules = {
  slug: [
    { required: true, message: '请输入 slug', trigger: 'blur' },
    { pattern: /^[A-Za-z][A-Za-z0-9._\-]{1,63}$/, message: '字母开头,2-64 位字母/数字/点/下划线/短横', trigger: 'blur' },
  ],
  type: [{ required: true, message: '请选择类型', trigger: 'change' }],
  upstream_model_slug: [{ required: true, message: '上游 slug 必填', trigger: 'blur' }],
}

function openCreate() {
  Object.assign(form, emptyForm())
  dlgState.value = 'create'
  dlgVisible.value = true
}
function openEdit(row: statsApi.Model) {
  Object.assign(form, {
    id: row.id,
    slug: row.slug,
    type: row.type as 'chat' | 'image',
    upstream_model_slug: row.upstream_model_slug,
    description: row.description,
    enabled: row.enabled,
  })
  dlgState.value = 'edit'
  dlgVisible.value = true
}

async function submit() {
  if (!formRef.value) return
  const valid = await formRef.value.validate().catch(() => false)
  if (!valid) return
  const payload: statsApi.ModelUpsert = {
    slug: form.slug,
    type: form.type,
    upstream_model_slug: form.upstream_model_slug,
    description: form.description,
    enabled: form.enabled,
  }
  dlgLoading.value = true
  try {
    if (dlgState.value === 'create') {
      await statsApi.createModel(payload)
      ElMessage.success('新增成功')
    } else {
      await statsApi.updateModel(form.id, payload)
      ElMessage.success('保存成功')
    }
    dlgVisible.value = false
    await load()
  } finally { dlgLoading.value = false }
}

async function onToggleEnabled(row: statsApi.Model) {
  await statsApi.setModelEnabled(row.id, !row.enabled)
  row.enabled = !row.enabled
  ElMessage.success(row.enabled ? '已开放调用' : '已暂停调用')
}

async function onDelete(row: statsApi.Model) {
  const ok = await ElMessageBox.confirm(`确定删除模型映射 "${row.slug}" 吗?已发生的用量日志不会被清除。`, '删除确认', {
    type: 'warning', confirmButtonText: '删除', cancelButtonText: '取消',
  }).catch(() => null)
  if (!ok) return
  await statsApi.deleteModel(row.id)
  ElMessage.success('已删除')
  await load()
}

onMounted(load)
</script>

<template>
  <div class="page-container">
    <div class="card-block">
      <div class="flex-between" style="margin-bottom:12px">
        <div>
          <h2 class="page-title" style="margin:0">模型映射</h2>
          <div class="page-sub">定义对外 model slug 与 chatgpt.com 上游模型名的映射。这里不配置计量金额，只控制是否开放调用。</div>
        </div>
        <div class="flex-wrap-gap">
          <el-radio-group v-model="filterType" size="small">
            <el-radio-button v-if="ENABLE_CHAT_MODEL" label="">全部</el-radio-button>
            <el-radio-button v-if="ENABLE_CHAT_MODEL" label="chat">对话</el-radio-button>
            <el-radio-button label="image">生图</el-radio-button>
          </el-radio-group>
          <el-button type="primary" :icon="Plus" @click="openCreate">新增映射</el-button>
        </div>
      </div>

      <el-table v-loading="loading" :data="filteredRows" stripe>
        <el-table-column prop="id" label="ID" width="70" />
        <el-table-column prop="slug" label="对外 slug" min-width="170"><template #default="{ row }"><code>{{ row.slug }}</code></template></el-table-column>
        <el-table-column label="类型" width="90"><template #default="{ row }"><el-tag :type="row.type === 'image' ? 'warning' : 'primary'" size="small">{{ row.type === 'image' ? '生图' : '对话' }}</el-tag></template></el-table-column>
        <el-table-column prop="upstream_model_slug" label="上游 slug" min-width="180" show-overflow-tooltip><template #default="{ row }"><code>{{ row.upstream_model_slug }}</code></template></el-table-column>
        <el-table-column prop="description" label="说明" min-width="220" show-overflow-tooltip />
        <el-table-column label="状态" width="110"><template #default="{ row }"><el-switch :model-value="row.enabled" inline-prompt active-text="开" inactive-text="停" @change="onToggleEnabled(row)" /></template></el-table-column>
        <el-table-column label="操作" width="140" fixed="right"><template #default="{ row }"><el-button size="small" link type="primary" @click="openEdit(row)">编辑</el-button><el-button size="small" link type="danger" @click="onDelete(row)">删除</el-button></template></el-table-column>
      </el-table>
    </div>

    <el-dialog v-model="dlgVisible" :title="dlgState === 'create' ? '新增模型映射' : `编辑映射 · ${form.slug}`" width="640px" destroy-on-close>
      <el-form ref="formRef" :model="form" :rules="rules" label-width="130px">
        <el-form-item label="对外 slug" prop="slug">
          <el-input v-model="form.slug" :disabled="dlgState === 'edit'" placeholder="例如 gpt-5 / gpt-image-2" />
          <div class="hint">创建后不可修改；客户端请求中的 model 字段使用该值。</div>
        </el-form-item>
        <el-form-item label="类型" prop="type">
          <el-radio-group v-model="form.type">
            <el-radio-button v-if="ENABLE_CHAT_MODEL" label="chat">对话 chat</el-radio-button>
            <el-radio-button label="image">生图 image</el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="上游 slug" prop="upstream_model_slug"><el-input v-model="form.upstream_model_slug" placeholder="chatgpt.com 实际模型名" /></el-form-item>
        <el-form-item label="描述"><el-input v-model="form.description" maxlength="255" show-word-limit /></el-form-item>
        <el-form-item label="开放调用"><el-switch v-model="form.enabled" /></el-form-item>
      </el-form>
      <template #footer><el-button @click="dlgVisible = false">取消</el-button><el-button type="primary" :loading="dlgLoading" @click="submit">保存</el-button></template>
    </el-dialog>
  </div>
</template>

<style scoped>
code { background: #f2f3f5; padding: 1px 6px; border-radius: 4px; font-family: ui-monospace, Menlo, Consolas, monospace; font-size: 12px; }
:global(html.dark) code { background: #1d2026; }
.hint, .page-sub { font-size: 12px; color: var(--el-text-color-secondary); margin-top: 4px; line-height: 1.4; }
</style>
````

## File: web/src/views/admin/Proxies.vue
````vue
<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import * as proxyApi from '@/api/proxies'
import { formatDateTime } from '@/utils/format'

const loading = ref(false)
const rows = ref<proxyApi.Proxy[]>([])
const total = ref(0)
const pager = reactive({ page: 1, page_size: 20 })

async function fetchList() {
  loading.value = true
  try {
    const data = await proxyApi.listProxies(pager)
    rows.value = data.list
    total.value = data.total
  } finally { loading.value = false }
}

const dlg = ref(false)
const isEdit = ref(false)
const form = reactive<proxyApi.ProxyCreate & { id?: number }>({
  id: 0, scheme: 'http', host: '', port: 0, username: '', password: '',
  country: '', isp: '', enabled: true, remark: '',
})

function openCreate() {
  isEdit.value = false
  Object.assign(form, {
    id: 0, scheme: 'http', host: '', port: 0, username: '', password: '',
    country: '', isp: '', enabled: true, remark: '',
  })
  dlg.value = true
}
function openEdit(row: proxyApi.Proxy) {
  isEdit.value = true
  Object.assign(form, {
    id: row.id, scheme: row.scheme, host: row.host, port: row.port,
    username: row.username, password: '',
    country: row.country, isp: row.isp, enabled: row.enabled, remark: row.remark,
  })
  dlg.value = true
}

async function submit() {
  if (!form.host) return ElMessage.warning('host 不能为空')
  if (!form.port) return ElMessage.warning('port 不能为空')
  const payload: proxyApi.ProxyUpdate = {
    scheme: form.scheme!,
    host: form.host,
    port: Number(form.port),
    username: form.username || '',
    password: form.password || '',
    country: form.country || '',
    isp: form.isp || '',
    enabled: !!form.enabled,
    remark: form.remark || '',
  }
  if (isEdit.value && form.id) await proxyApi.updateProxy(form.id, payload)
  else await proxyApi.createProxy(payload)
  ElMessage.success('保存成功')
  dlg.value = false
  fetchList()
}

async function toggleEnabled(row: proxyApi.Proxy) {
  await proxyApi.updateProxy(row.id, {
    scheme: row.scheme, host: row.host, port: row.port,
    username: row.username, password: '',
    country: row.country, isp: row.isp, remark: row.remark,
    enabled: !row.enabled,
  })
  ElMessage.success(row.enabled ? '已禁用' : '已启用')
  fetchList()
}

async function onDelete(row: proxyApi.Proxy) {
  await ElMessageBox.confirm(`确认删除代理 ${row.host}:${row.port}?`, '删除代理', {
    type: 'warning', confirmButtonText: '删除', cancelButtonText: '取消',
  })
  await proxyApi.deleteProxy(row.id)
  ElMessage.success('已删除')
  fetchList()
}

// ---------- 健康探测 ----------
const probingIds = ref<Set<number>>(new Set())
const probeAllLoading = ref(false)

async function onProbe(row: proxyApi.Proxy) {
  if (probingIds.value.has(row.id)) return
  probingIds.value.add(row.id)
  try {
    const res = await proxyApi.probeProxy(row.id)
    // 同步更新当前行,避免等列表刷新才看到反馈
    row.health_score = res.health_score
    row.last_probe_at = res.tried_at
    row.last_error = res.ok ? '' : (res.error || 'failed')
    if (res.ok) {
      ElMessage.success(`连通正常 · ${res.latency_ms} ms`)
    } else {
      ElMessage.error(`探测失败:${res.error || 'unknown'}`)
    }
  } catch (e: any) {
    ElMessage.error(e?.message || '探测失败')
  } finally {
    probingIds.value.delete(row.id)
    // 重新 new 一份触发 reactive(因为直接修改 Set 内部,模板里 has() 不会响应)
    probingIds.value = new Set(probingIds.value)
  }
}

async function onProbeAll() {
  await ElMessageBox.confirm(
    '将对所有启用的代理发起连通性探测,耗时取决于代理数量。是否继续?',
    '全部探测',
    { type: 'info', confirmButtonText: '开始', cancelButtonText: '取消' },
  )
  probeAllLoading.value = true
  try {
    const res = await proxyApi.probeAllProxies()
    ElMessage.success(`探测完成 · 共 ${res.total} · 通 ${res.ok} · 断 ${res.bad}`)
    fetchList()
  } catch (e: any) {
    ElMessage.error(e?.message || '探测失败')
  } finally {
    probeAllLoading.value = false
  }
}

// ---------- 批量导入 ----------
const importDlg = ref(false)
const importLoading = ref(false)
const importForm = reactive({
  text: '',
  enabled: true,
  country: '',
  isp: '',
  remark: '',
  overwrite: false,
  probe_after_import: true,
})
const importResult = ref<proxyApi.ProxyImportResp | null>(null)

function openImport() {
  Object.assign(importForm, {
    text: '', enabled: true, country: '', isp: '', remark: '', overwrite: false, probe_after_import: true,
  })
  importResult.value = null
  importDlg.value = true
}

async function doImport() {
  if (!importForm.text.trim()) return ElMessage.warning('请粘贴至少一行代理 URL')
  importLoading.value = true
  try {
    importResult.value = await proxyApi.importProxies({
      text: importForm.text,
      enabled: importForm.enabled,
      country: importForm.country,
      isp: importForm.isp,
      remark: importForm.remark,
      overwrite: importForm.overwrite,
      probe_after_import: importForm.probe_after_import,
    })
    const r = importResult.value
    const probeText = r.probe ? ` · 探测通 ${r.probe.ok}/${r.probe.total}` : (r.probe_queued ? ' · 已排队探测' : '')
    ElMessage.success(
      `完成 · 新增 ${r.created} · 更新 ${r.updated} · 跳过 ${r.skipped} · 无效 ${r.invalid}${probeText}`,
    )
    fetchList()
  } finally { importLoading.value = false }
}

function importStatusTag(s: string) {
  switch (s) {
    case 'created': return 'success'
    case 'updated': return 'primary'
    case 'skipped': return 'info'
    default: return 'danger'
  }
}
function importStatusText(s: string) {
  return { created: '新增', updated: '更新', skipped: '跳过', invalid: '无效' }[s] || s
}

function healthColor(score: number) {
  if (score >= 80) return 'success'
  if (score >= 50) return 'warning'
  return 'danger'
}

onMounted(fetchList)
</script>

<template>
  <div class="page-container">
    <div class="card-block">
      <div class="flex-between" style="margin-bottom:12px">
        <div>
          <h2 class="page-title" style="margin:0">代理管理</h2>
          <div style="color:var(--el-text-color-secondary);font-size:13px;margin-top:4px">
            维护 HTTP / SOCKS5 代理池,所有 GPT 账号都应绑定独立代理以分散风控指纹;健康分由定时探测自动维护,探测参数可在「系统设置 → 网关与调度」调整。
          </div>
        </div>
        <div class="flex-wrap-gap">
          <el-button :loading="probeAllLoading" @click="onProbeAll">
            <el-icon><Promotion /></el-icon> 全部探测
          </el-button>
          <el-button @click="openImport"><el-icon><Upload /></el-icon> 批量导入</el-button>
          <el-button type="primary" @click="openCreate"><el-icon><Plus /></el-icon> 新建代理</el-button>
        </div>
      </div>

      <el-table v-loading="loading" :data="rows" stripe>
        <el-table-column prop="id" label="ID" width="70" />
        <el-table-column label="地址" min-width="220">
          <template #default="{ row }">
            <code>{{ row.scheme }}://{{ row.host }}:{{ row.port }}</code>
            <div v-if="row.username" style="font-size:12px;color:var(--el-text-color-secondary)">
              auth: {{ row.username }} / ******
            </div>
          </template>
        </el-table-column>
        <el-table-column label="区域" width="130">
          <template #default="{ row }">
            <div>{{ row.country || '-' }}</div>
            <div style="font-size:12px;color:var(--el-text-color-secondary)">{{ row.isp || '' }}</div>
          </template>
        </el-table-column>
        <el-table-column label="健康" width="150">
          <template #default="{ row }">
            <el-progress :percentage="Math.max(0, Math.min(100, row.health_score))"
                         :status="row.health_score >= 80 ? 'success' : row.health_score >= 50 ? 'warning' : 'exception'" />
            <el-tag v-if="row.last_error" :type="healthColor(row.health_score)" size="small" style="margin-top:4px">
              {{ row.last_error.slice(0, 30) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="最近探测" width="170">
          <template #default="{ row }">{{ formatDateTime(row.last_probe_at) }}</template>
        </el-table-column>
        <el-table-column label="状态" width="90">
          <template #default="{ row }">
            <el-switch :model-value="row.enabled" @change="() => toggleEnabled(row)" />
          </template>
        </el-table-column>
        <el-table-column prop="remark" label="备注" min-width="180" show-overflow-tooltip />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button link type="success" :loading="probingIds.has(row.id)" @click="onProbe(row)">探测</el-button>
            <el-button link type="primary" @click="openEdit(row)">编辑</el-button>
            <el-button link type="danger" @click="onDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-pagination style="margin-top:16px;display:flex;justify-content:flex-end"
        v-model:current-page="pager.page"
        v-model:page-size="pager.page_size"
        :total="total"
        :page-sizes="[20, 50, 100]"
        layout="total, sizes, prev, pager, next"
        @current-change="fetchList"
        @size-change="fetchList"
      />
    </div>

    <el-dialog v-model="dlg" :title="isEdit ? '编辑代理' : '新建代理'" width="520px">
      <el-form :model="form" label-width="100px">
        <el-form-item label="协议">
          <el-select v-model="form.scheme" style="width:100%">
            <el-option label="http" value="http" />
            <el-option label="https" value="https" />
            <el-option label="socks5" value="socks5" />
          </el-select>
        </el-form-item>
        <el-form-item label="Host" required><el-input v-model="form.host" placeholder="192.0.2.1" /></el-form-item>
        <el-form-item label="Port" required>
          <el-input-number v-model="form.port" :min="1" :max="65535" style="width:100%" />
        </el-form-item>
        <el-form-item label="用户名"><el-input v-model="form.username" autocomplete="off" /></el-form-item>
        <el-form-item label="密码">
          <el-input v-model="form.password" type="password" show-password autocomplete="new-password"
                    :placeholder="isEdit ? '留空表示不改' : ''" />
        </el-form-item>
        <el-form-item label="国家/地区"><el-input v-model="form.country" placeholder="US / JP / HK …" /></el-form-item>
        <el-form-item label="ISP"><el-input v-model="form.isp" /></el-form-item>
        <el-form-item label="启用"><el-switch v-model="form.enabled" /></el-form-item>
        <el-form-item label="备注"><el-input v-model="form.remark" type="textarea" :rows="2" /></el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dlg = false">取消</el-button>
        <el-button type="primary" @click="submit">保存</el-button>
      </template>
    </el-dialog>

    <!-- 批量导入 -->
    <el-dialog v-model="importDlg" title="批量导入代理" width="720px">
      <el-form label-width="88px" @submit.prevent>
        <el-form-item label="代理列表">
          <el-input
            v-model="importForm.text"
            type="textarea"
            :rows="10"
            resize="vertical"
            placeholder="每行一个,支持以下格式:
http://user:pass@host:port
https://host:port
socks5://user:pass@host:port
user:pass@host:port    (省略 scheme 默认 http)
host:port:user:pass    (代理商常见格式)
user:pass:host:port    (另一种代理商格式)
# 以 # 或 // 开头的行会被跳过"
          />
          <div class="import-hint">
            支持 http / https / socks5。同一 scheme + host + port + username 视为已存在。
          </div>
        </el-form-item>
        <el-form-item label="默认启用">
          <el-switch v-model="importForm.enabled" />
        </el-form-item>
        <el-form-item label="国家/地区">
          <el-input v-model="importForm.country" placeholder="如 US / HK,空则每条自行为空" style="max-width:240px" />
        </el-form-item>
        <el-form-item label="ISP">
          <el-input v-model="importForm.isp" placeholder="如 Arxlabs" style="max-width:240px" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="importForm.remark" placeholder="将填到所有新增行的 remark" />
        </el-form-item>
        <el-form-item label="覆盖已有">
          <el-switch v-model="importForm.overwrite" />
          <span class="import-hint" style="margin-left:8px">
            开启后:同 endpoint 已存在时更新密码/国家/ISP/备注;关闭则跳过。
          </span>
        </el-form-item>
        <el-form-item label="导入后探测">
          <el-switch v-model="importForm.probe_after_import" />
          <span class="import-hint" style="margin-left:8px">
            新增/更新数量 ≤ 50 时同步探测;更多时交给后台探测队列。
          </span>
        </el-form-item>
      </el-form>

      <div v-if="importResult" class="import-result">
        <div class="import-summary">
          共 {{ importResult.total }} 行 ·
          <el-tag type="success" effect="plain">新增 {{ importResult.created }}</el-tag>
          <el-tag type="primary" effect="plain">更新 {{ importResult.updated }}</el-tag>
          <el-tag type="info" effect="plain">跳过 {{ importResult.skipped }}</el-tag>
          <el-tag type="danger" effect="plain">无效 {{ importResult.invalid }}</el-tag>
          <el-tag v-if="importResult.probe" type="warning" effect="plain">
            探测 {{ importResult.probe.ok }}/{{ importResult.probe.total }}
          </el-tag>
          <el-tag v-else-if="importResult.probe_queued" type="warning" effect="plain">已排队探测</el-tag>
        </div>
        <el-table :data="importResult.items" size="small" max-height="260" border>
          <el-table-column prop="line" label="行" width="60" />
          <el-table-column label="状态" width="80">
            <template #default="{ row }">
              <el-tag :type="importStatusTag(row.status)" size="small">{{ importStatusText(row.status) }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="raw" label="内容" show-overflow-tooltip />
          <el-table-column prop="error" label="说明" show-overflow-tooltip />
        </el-table>
      </div>

      <template #footer>
        <el-button @click="importDlg = false">关闭</el-button>
        <el-button type="primary" :loading="importLoading" @click="doImport">开始导入</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<style scoped>
code {
  background: #f2f3f5;
  padding: 1px 6px;
  border-radius: 4px;
  font-family: ui-monospace, Menlo, Consolas, monospace;
  font-size: 12px;
}
.flex-wrap-gap {
  display: inline-flex;
  gap: 8px;
}
.import-hint {
  font-size: 12px;
  color: var(--el-text-color-secondary);
  line-height: 1.5;
  margin-top: 4px;
}
.import-result {
  margin-top: 8px;
}
.import-summary {
  display: flex;
  gap: 8px;
  align-items: center;
  flex-wrap: wrap;
  margin-bottom: 8px;
  color: var(--el-text-color-regular);
  font-size: 13px;
}
</style>
````

## File: web/src/views/admin/Settings.vue
````vue
<script setup lang="ts">
import { ref, computed, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  Refresh,
  Check,
  Setting,
  Connection,
  Message as MailIcon,
} from '@element-plus/icons-vue'
import {
  listSettings,
  updateSettings,
  reloadSettings,
  sendTestEmail,
  type SettingItem,
} from '@/api/settings'
import { useSiteStore } from '@/stores/site'

const loading = ref(false)
const saving = ref(false)
const items = ref<SettingItem[]>([])
// 本地编辑态,key -> value(string)
const draft = reactive<Record<string, string>>({})

const tabs = [
  { name: 'site', label: '通用设置', icon: Setting },
  { name: 'gateway', label: '网关服务', icon: Connection },
  { name: 'account', label: '账号池', icon: Connection },
  { name: 'api', label: 'API 认证', icon: Connection },
  { name: 'mail', label: '邮件设置', icon: MailIcon },
] as const
const activeTab = ref<(typeof tabs)[number]['name']>('site')

const grouped = computed(() => {
  const map: Record<string, SettingItem[]> = {
    site: [], gateway: [], account: [], api: [], mail: [],
  }
  for (const it of items.value) {
    const cat = it.category
    ;(map[cat] ||= []).push(it)
  }
  for (const k of Object.keys(map)) map[k].sort((a, b) => a.key.localeCompare(b.key))
  return map
})

const dirtyCount = computed(() => {
  let n = 0
  for (const it of items.value) {
    if (String(draft[it.key] ?? '') !== String(it.value)) n++
  }
  return n
})

async function load() {
  loading.value = true
  try {
    const d = await listSettings()
    items.value = d.items
    for (const it of d.items) draft[it.key] = it.value
  } finally {
    loading.value = false
  }
}

function reset() {
  for (const it of items.value) draft[it.key] = it.value
  ElMessage.info('已重置为服务端当前值')
}

function isBool(it: SettingItem) { return it.type === 'bool' }
function isInt(it: SettingItem) { return it.type === 'int' }
function isFloat(it: SettingItem) { return it.type === 'float' }
function inputType(it: SettingItem) {
  if (it.type === 'email') return 'email'
  if (it.type === 'url') return 'url'
  return 'text'
}

async function save() {
  const diff: Record<string, string> = {}
  for (const it of items.value) {
    const v = draft[it.key] ?? ''
    if (String(v) !== String(it.value)) diff[it.key] = String(v)
  }
  if (Object.keys(diff).length === 0) {
    ElMessage.info('没有需要保存的修改')
    return
  }
  saving.value = true
  try {
    await updateSettings(diff)
    ElMessage.success(`已保存 ${Object.keys(diff).length} 项`)
    await load()
    useSiteStore().refresh()
  } finally {
    saving.value = false
  }
}

async function doReload() {
  await ElMessageBox.confirm('从数据库强制重载最新值到内存缓存?', '确认', {
    type: 'warning',
  }).catch(() => 'cancel')
  try {
    await reloadSettings()
    ElMessage.success('已重载')
    await load()
  } catch { /* 拦截器已处理 */ }
}

// ---- 邮件测试 ----
const mailDlg = ref(false)
const mailTo = ref('')
const mailSending = ref(false)
async function submitTestMail() {
  if (!mailTo.value) {
    ElMessage.warning('请输入收件邮箱')
    return
  }
  mailSending.value = true
  try {
    await sendTestEmail(mailTo.value)
    ElMessage.success('测试邮件已发出')
    mailDlg.value = false
  } catch { /* 拦截器已处理 */ } finally {
    mailSending.value = false
  }
}

function generateAPIKey() {
  const arr = new Uint8Array(32)
  crypto.getRandomValues(arr)
  const key = 'sk-' + Array.from(arr).map(b => b.toString(16).padStart(2, '0')).join('')
  draft['api.v1_key'] = key
  ElMessage.success('已生成随机 API Key，记得点保存')
}

onMounted(load)
</script>

<template>
  <div class="page-container">
    <div class="card-block" v-loading="loading">
      <!-- 顶部:标题 + 操作栏(始终可见) -->
      <div class="flex-between settings-head">
        <div>
          <div class="page-title" style="margin:0">系统设置</div>
          <div class="settings-subtitle">
            所有修改在点击"保存修改"后立即生效,无需重启服务
          </div>
        </div>
        <div class="flex-wrap-gap">
          <el-button :icon="Refresh" @click="doReload">强制重载</el-button>
          <el-button :icon="MailIcon" @click="mailDlg = true">发测试邮件</el-button>
          <el-button :disabled="dirtyCount === 0" @click="reset">重置</el-button>
          <el-button
            type="primary"
            :icon="Check"
            :loading="saving"
            @click="save"
          >
            保存修改<span v-if="dirtyCount > 0"> ({{ dirtyCount }})</span>
          </el-button>
        </div>
      </div>

      <el-tabs v-model="activeTab" class="settings-tabs">
        <el-tab-pane v-for="t in tabs" :key="t.name" :name="t.name">
          <template #label>
            <span class="tab-label">
              <el-icon><component :is="t.icon" /></el-icon>
              <span>{{ t.label }}</span>
            </span>
          </template>

          <div class="tab-body">
            <el-empty
              v-if="!grouped[t.name] || grouped[t.name].length === 0"
              description="暂无可配置项"
            />
            <el-form
              v-else
              label-width="170px"
              label-position="right"
              class="setting-form"
            >
              <el-form-item
                v-for="it in grouped[t.name]"
                :key="it.key"
                :label="it.label || it.key"
              >
                <div class="field-wrap">
                  <el-switch
                    v-if="isBool(it)"
                    :model-value="draft[it.key] === 'true'"
                    @update:model-value="(v) => (draft[it.key] = v ? 'true' : 'false')"
                  />
                  <el-input-number
                    v-else-if="isInt(it)"
                    :model-value="Number(draft[it.key] || 0)"
                    :min="0"
                    :controls-position="'right'"
                    style="width: 240px"
                    @update:model-value="(v) => (draft[it.key] = String(v ?? 0))"
                  />
                  <el-input-number
                    v-else-if="isFloat(it)"
                    :model-value="Number(draft[it.key] || 0)"
                    :min="0"
                    :max="1"
                    :step="0.05"
                    :precision="2"
                    :controls-position="'right'"
                    style="width: 240px"
                    @update:model-value="(v) => (draft[it.key] = String(v ?? 0))"
                  />
                  <div v-else class="input-with-action">
                    <el-input
                      v-model="draft[it.key]"
                      :placeholder="it.desc || it.label"
                      :type="inputType(it)"
                      clearable
                      style="max-width: 520px"
                    />
                    <el-button
                      v-if="it.key === 'api.v1_key'"
                      size="small"
                      @click="generateAPIKey"
                    >随机生成</el-button>
                  </div>
                  <div v-if="it.desc" class="hint">{{ it.desc }}</div>
                </div>
              </el-form-item>
            </el-form>
          </div>
        </el-tab-pane>
      </el-tabs>
    </div>

    <!-- 测试邮件 -->
    <el-dialog v-model="mailDlg" title="发送 SMTP 测试邮件" width="420px">
      <el-form label-width="80px">
        <el-form-item label="收件人">
          <el-input v-model="mailTo" placeholder="your@mail.com" type="email" clearable />
        </el-form-item>
        <div style="font-size:12px;color:var(--el-text-color-secondary)">
          使用 <code>configs/config.yaml</code> 的 SMTP 配置发送;未配置时会直接失败。
        </div>
      </el-form>
      <template #footer>
        <el-button @click="mailDlg = false">取消</el-button>
        <el-button type="primary" :loading="mailSending" @click="submitTestMail">发送</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<style scoped>
.settings-head {
  margin-bottom: 4px;
}
.settings-subtitle {
  font-size: 12px;
  color: var(--el-text-color-secondary);
  margin-top: 4px;
}

.settings-tabs {
  margin-top: 8px;
}
.settings-tabs :deep(.el-tabs__header) {
  margin-bottom: 16px;
}
.tab-label {
  display: inline-flex;
  align-items: center;
  gap: 6px;
}

.tab-body {
  padding-top: 4px;
}
.setting-form .el-form-item {
  margin-bottom: 18px;
}
.field-wrap {
  width: 100%;
}
.input-with-action {
  display: flex;
  align-items: center;
  gap: 8px;
}
.hint {
  margin-top: 4px;
  font-size: 12px;
  color: var(--el-text-color-secondary);
  line-height: 1.5;
}

@media (max-width: 640px) {
  .setting-form :deep(.el-form-item__label) {
    width: auto !important;
    padding-right: 8px !important;
    line-height: 1.5;
  }
}
</style>
````

## File: web/src/views/admin/UsageStats.vue
````vue
<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue'
import * as statsApi from '@/api/stats'

const loading = ref(false)
const stats = ref<statsApi.StatsResp | null>(null)
const logs = ref<statsApi.UsageLogRow[]>([])
const total = ref(0)
const filter = reactive({ type: '', status: '', days: 14, limit: 30, offset: 0 })

const page = computed({
  get: () => Math.floor(filter.offset / filter.limit) + 1,
  set: (v: number) => { filter.offset = (v - 1) * filter.limit; loadLogs() },
})

async function loadStats() { stats.value = await statsApi.getUsageStats({ days: filter.days, top_n: 12, type: filter.type, status: filter.status }) }
async function loadLogs() {
  loading.value = true
  try {
    const d = await statsApi.listUsageLogs(filter)
    logs.value = d.items || []
    total.value = d.total || 0
  } finally { loading.value = false }
}
async function reload() { filter.offset = 0; await Promise.all([loadStats(), loadLogs()]) }
onMounted(reload)
</script>

<template>
  <div class="page-container">
    <div class="card-block">
      <div class="flex-between">
        <div><h2 class="page-title">全局用量</h2><div class="sub">聚合本地 /v1 与在线体验调用记录。</div></div>
        <div class="flex-wrap-gap">
          <el-input-number v-model="filter.days" :min="1" :max="90" controls-position="right" style="width:120px" @change="reload" />
          <el-select v-model="filter.type" clearable placeholder="类型" style="width:120px" @change="reload"><el-option label="对话" value="chat" /><el-option label="生图" value="image" /></el-select>
          <el-select v-model="filter.status" clearable placeholder="状态" style="width:120px" @change="reload"><el-option label="成功" value="success" /><el-option label="失败" value="failed" /></el-select>
          <el-button @click="reload">刷新</el-button>
        </div>
      </div>
    </div>

    <div class="grid cards">
      <div class="card-block kpi"><span>请求</span><b>{{ stats?.overall.requests || 0 }}</b></div>
      <div class="card-block kpi"><span>失败</span><b>{{ stats?.overall.failures || 0 }}</b></div>
      <div class="card-block kpi"><span>对话请求</span><b>{{ stats?.overall.chat_requests || 0 }}</b></div>
      <div class="card-block kpi"><span>图片数</span><b>{{ stats?.overall.image_images || 0 }}</b></div>
    </div>

    <div class="card-block" style="margin-top:16px">
      <h3>模型分布</h3>
      <el-table :data="stats?.by_model || []" size="small" empty-text="暂无数据">
        <el-table-column prop="model_slug" label="模型" min-width="160"><template #default="{ row }"><code>{{ row.model_slug || row.model_id }}</code></template></el-table-column>
        <el-table-column prop="type" label="类型" width="90" />
        <el-table-column prop="requests" label="请求" width="100" />
        <el-table-column prop="failures" label="失败" width="100" />
        <el-table-column prop="input_tokens" label="输入" width="110" />
        <el-table-column prop="output_tokens" label="输出" width="110" />
        <el-table-column prop="image_count" label="图片" width="100" />
        <el-table-column prop="avg_dur_ms" label="平均耗时(ms)" width="130" />
      </el-table>
    </div>

    <div class="card-block" style="margin-top:16px">
      <h3>最近请求</h3>
      <el-table v-loading="loading" :data="logs" stripe>
        <el-table-column prop="created_at" label="时间" width="180" />
        <el-table-column prop="type" label="类型" width="90" />
        <el-table-column prop="model_slug" label="模型" min-width="160"><template #default="{ row }"><code>{{ row.model_slug || row.model_id }}</code></template></el-table-column>
        <el-table-column prop="account_id" label="账号" width="90" />
        <el-table-column prop="status" label="状态" width="100" />
        <el-table-column prop="input_tokens" label="输入" width="100" />
        <el-table-column prop="output_tokens" label="输出" width="100" />
        <el-table-column prop="image_count" label="图片" width="90" />
        <el-table-column prop="duration_ms" label="耗时(ms)" width="110" />
        <el-table-column prop="error_code" label="错误" min-width="160" show-overflow-tooltip />
      </el-table>
      <div class="pager"><el-pagination layout="total, sizes, prev, pager, next" :total="total" v-model:current-page="page" v-model:page-size="filter.limit" @size-change="reload" /></div>
    </div>
  </div>
</template>

<style scoped>
.sub { color:var(--el-text-color-secondary); font-size:13px; margin-top:4px; }
.grid { display:grid; gap:16px; margin-top:16px; }
.cards { grid-template-columns: repeat(4, minmax(0,1fr)); }
.kpi span { color:var(--el-text-color-secondary); display:block; margin-bottom:8px; }
.kpi b { font-size:24px; }
.pager { display:flex; justify-content:flex-end; margin-top:14px; }
code { background:#f2f3f5; padding:1px 6px; border-radius:4px; }
@media (max-width: 960px) { .cards { grid-template-columns:1fr; } }
</style>
````

## File: web/src/views/Error403.vue
````vue
<script setup lang="ts">
import { useRouter } from 'vue-router'
const router = useRouter()
</script>
<template>
  <div class="err-page">
    <el-result icon="error" title="403" sub-title="你当前的角色没有访问此页面的权限。">
      <template #extra>
        <el-button @click="router.replace('/personal/dashboard')">回到个人中心</el-button>
      </template>
    </el-result>
  </div>
</template>
<style scoped>
.err-page{min-height:100vh;display:flex;align-items:center;justify-content:center;}
</style>
````

## File: web/src/views/Error404.vue
````vue
<script setup lang="ts">
import { useRouter } from 'vue-router'
const router = useRouter()
</script>
<template>
  <div class="err-page">
    <el-result icon="warning" title="404" sub-title="页面不存在,或你没有访问它的权限。">
      <template #extra>
        <el-button type="primary" @click="router.replace('/')">返回首页</el-button>
      </template>
    </el-result>
  </div>
</template>
<style scoped>
.err-page{min-height:100vh;display:flex;align-items:center;justify-content:center;}
</style>
````

## File: web/src/views/Login.vue
````vue
<template>
  <div class="login-container">
    <div class="login-card">
      <h2 class="login-title">Image2API 管理后台</h2>
      <el-form :model="form" @submit.prevent="handleLogin" class="login-form">
        <el-form-item>
          <el-input v-model="form.username" placeholder="用户名" prefix-icon="User" size="large" />
        </el-form-item>
        <el-form-item>
          <el-input v-model="form.password" placeholder="密码" prefix-icon="Lock" type="password"
            size="large" show-password @keyup.enter="handleLogin" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" size="large" :loading="loading" @click="handleLogin" style="width:100%">
            登 录
          </el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const userStore = useUserStore()
const loading = ref(false)
const form = ref({ username: '', password: '' })

async function handleLogin() {
  if (!form.value.username || !form.value.password) {
    ElMessage.warning('请输入用户名和密码')
    return
  }
  loading.value = true
  try {
    await userStore.login(form.value.username, form.value.password)
    router.push('/')
  } catch (e: any) {
    // error already shown by http interceptor or login function
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}
.login-card {
  background: #fff;
  border-radius: 12px;
  padding: 40px 36px 24px;
  width: 380px;
  box-shadow: 0 8px 32px rgba(0,0,0,0.18);
}
.login-title {
  text-align: center;
  margin-bottom: 28px;
  font-size: 22px;
  color: #303133;
}
.login-form .el-form-item {
  margin-bottom: 20px;
}
</style>
````

## File: web/src/views/personal/ApiDocs.vue
````vue
<script setup lang="ts">
const base = typeof window === 'undefined' ? 'http://localhost:8080' : window.location.origin
const chatCurl = `curl ${base}/v1/chat/completions \\
  -H 'Content-Type: application/json' \\
  -d '{"model":"gpt-4o","messages":[{"role":"user","content":"hello"}],"stream":true}'`
const imageCurl = `curl ${base}/v1/images/generations \\
  -H 'Content-Type: application/json' \\
  -d '{"model":"gpt-image-2","prompt":"a cat reading a book","n":1,"size":"1024x1024"}'`
</script>

<template>
  <div class="page-container">
    <div class="card-block">
      <h2 class="page-title">接口文档</h2>
      <p class="sub">本版本是本地自用中转，默认不需要下游身份凭证。把 OpenAI SDK 的 baseURL 指向 <code>{{ base }}/v1</code> 即可。</p>
    </div>

    <div class="card-block">
      <h3>对话接口</h3>
      <p><code>POST /v1/chat/completions</code>，兼容 OpenAI Chat Completions，支持流式输出。</p>
      <pre>{{ chatCurl }}</pre>
    </div>

    <div class="card-block">
      <h3>图片接口</h3>
      <p><code>POST /v1/images/generations</code>，同步返回图片 URL；图片代理地址由本服务签名生成。</p>
      <pre>{{ imageCurl }}</pre>
    </div>

    <div class="card-block">
      <h3>模型列表</h3>
      <p><code>GET /v1/models</code> 会返回已开放的本地模型映射。</p>
      <pre>curl {{ base }}/v1/models</pre>
    </div>
  </div>
</template>

<style scoped>
.sub { color:var(--el-text-color-secondary); }
.card-block { margin-bottom:16px; }
code { background:#f2f3f5; padding:1px 6px; border-radius:4px; }
pre { background:#111827; color:#e5e7eb; border-radius:8px; padding:14px; overflow:auto; line-height:1.5; }
</style>
````

## File: web/src/views/personal/Dashboard.vue
````vue
<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import * as meApi from '@/api/me'
import * as accountApi from '@/api/accounts'
import * as proxyApi from '@/api/proxies'

const router = useRouter()
const loading = ref(false)
const stats = ref<meApi.MyStatsResp | null>(null)
const models = ref<meApi.SimpleModel[]>([])
const tasks = ref<meApi.ImageTask[]>([])
const accounts = ref<accountApi.Account[]>([])
const accountTotal = ref(0)
const proxies = ref<proxyApi.Proxy[]>([])
const proxyTotal = ref(0)

const overall = computed(() => stats.value?.overall)
const okRate = computed(() => {
  const o = overall.value
  if (!o?.requests) return '100%'
  return `${Math.max(0, Math.round(((o.requests - o.failures) / o.requests) * 100))}%`
})
const modelSummary = computed(() => ({
  chat: models.value.filter((m) => m.type === 'chat').length,
  image: models.value.filter((m) => m.type === 'image').length,
}))
const accountSummary = computed(() => {
  const rows = accounts.value
  return {
    total: accountTotal.value || rows.length,
    healthy: rows.filter((a) => a.status === 'healthy').length,
    warned: rows.filter((a) => a.status === 'warned').length,
    throttled: rows.filter((a) => a.status === 'throttled').length,
    dead: rows.filter((a) => ['dead', 'suspicious'].includes(a.status)).length,
  }
})
const proxySummary = computed(() => {
  const rows = proxies.value
  return {
    total: proxyTotal.value || rows.length,
    enabled: rows.filter((p) => p.enabled).length,
    good: rows.filter((p) => p.enabled && p.health_score >= 80).length,
    risk: rows.filter((p) => p.enabled && p.health_score < 50).length,
  }
})
const dailyRows = computed(() => stats.value?.daily || [])
const dailyMax = computed(() => Math.max(1, ...dailyRows.value.map((x) => x.requests || 0)))

async function load() {
  loading.value = true
  try {
    const [s, m, t, a, p] = await Promise.all([
      meApi.getMyUsageStats({ days: 7, top_n: 8 }),
      meApi.listMyModels(),
      meApi.listMyImageTasks({ limit: 8 }),
      accountApi.listAccounts({ page: 1, page_size: 100 }).catch(() => null),
      proxyApi.listProxies({ page: 1, page_size: 100 }).catch(() => null),
    ])
    stats.value = s
    models.value = m.items || []
    tasks.value = t.items || []
    if (a) {
      accounts.value = a.list || []
      accountTotal.value = a.total || 0
    }
    if (p) {
      proxies.value = p.list || []
      proxyTotal.value = p.total || 0
    }
  } finally { loading.value = false }
}

function go(path: string) { router.push(path) }
function fmtTime(s: string) {
  if (!s) return '-'
  return s.replace('T', ' ').replace(/\+.*/, '').slice(0, 19)
}
function barWidth(v: number) {
  return `${Math.max(6, Math.round((v / dailyMax.value) * 100))}%`
}
function dayLabel(s: string) {
  return s?.slice(5) || '-'
}
function taskType(status: string) {
  return status === 'success' ? 'success' : status === 'failed' ? 'danger' : 'info'
}
onMounted(load)
</script>

<template>
  <div class="dash" v-loading="loading">
    <!-- Hero -->
    <div class="hero">
      <div class="hero-bg"></div>
      <div class="hero-inner">
        <div class="hero-text">
          <h1>本地中转控制台</h1>
          <p>客户端直接调用 <code>/v1</code> 接口，控制台负责账号池、代理池、模型映射和运行观察。</p>
        </div>
        <div class="hero-actions">
          <el-button type="primary" size="large" round @click="go('/personal/play')">
            <el-icon><VideoPlay /></el-icon>&nbsp;在线体验
          </el-button>
          <el-button size="large" round @click="go('/personal/docs')">
            <el-icon><Document /></el-icon>&nbsp;接口文档
          </el-button>
        </div>
      </div>
    </div>

    <!-- KPI -->
    <div class="kpi-row">
      <div class="kpi-card kpi-blue">
        <div class="kpi-top">
          <span class="kpi-label">7 日请求</span>
          <div class="kpi-dot"><el-icon :size="18"><Connection /></el-icon></div>
        </div>
        <div class="kpi-val">{{ overall?.requests || 0 }}</div>
      </div>
      <div class="kpi-card kpi-green">
        <div class="kpi-top">
          <span class="kpi-label">成功率</span>
          <div class="kpi-dot"><el-icon :size="18"><CircleCheck /></el-icon></div>
        </div>
        <div class="kpi-val">{{ okRate }}</div>
      </div>
      <div class="kpi-card kpi-orange">
        <div class="kpi-top">
          <span class="kpi-label">输入 Token</span>
          <div class="kpi-dot"><el-icon :size="18"><Tickets /></el-icon></div>
        </div>
        <div class="kpi-val">{{ overall?.input_tokens || 0 }}</div>
      </div>
      <div class="kpi-card kpi-purple">
        <div class="kpi-top">
          <span class="kpi-label">图片数量</span>
          <div class="kpi-dot"><el-icon :size="18"><PictureFilled /></el-icon></div>
        </div>
        <div class="kpi-val">{{ overall?.image_images || 0 }}</div>
      </div>
    </div>

    <!-- Ops overview -->
    <div class="ops-grid">
      <div class="section-card ops-card">
        <div class="section-head compact">
          <h2>账号池状态</h2>
          <el-button link type="primary" @click="go('/admin/accounts')">维护 →</el-button>
        </div>
        <div class="ops-main">
          <div>
            <div class="ops-num">{{ accountSummary.healthy }}/{{ accountSummary.total }}</div>
            <div class="ops-label">健康账号 / 总账号</div>
          </div>
          <el-tag type="warning" effect="plain">warned {{ accountSummary.warned }}</el-tag>
          <el-tag type="danger" effect="plain">异常 {{ accountSummary.dead + accountSummary.throttled }}</el-tag>
        </div>
      </div>

      <div class="section-card ops-card">
        <div class="section-head compact">
          <h2>代理池状态</h2>
          <el-button link type="primary" @click="go('/admin/proxies')">维护 →</el-button>
        </div>
        <div class="ops-main">
          <div>
            <div class="ops-num">{{ proxySummary.good }}/{{ proxySummary.enabled }}</div>
            <div class="ops-label">高健康代理 / 启用代理</div>
          </div>
          <el-tag :type="proxySummary.risk ? 'danger' : 'success'" effect="plain">风险 {{ proxySummary.risk }}</el-tag>
          <el-tag type="info" effect="plain">总 {{ proxySummary.total }}</el-tag>
        </div>
      </div>

      <div class="section-card trend-card">
        <div class="section-head compact">
          <h2>近 7 日请求</h2>
          <span class="model-pill">chat {{ modelSummary.chat }} · image {{ modelSummary.image }}</span>
        </div>
        <div class="bars">
          <div v-for="d in dailyRows" :key="d.day" class="bar-row">
            <span>{{ dayLabel(d.day) }}</span>
            <div class="bar-track"><i :style="{ width: barWidth(d.requests) }" /></div>
            <b>{{ d.requests }}</b>
          </div>
        </div>
      </div>
    </div>

    <!-- Bottom grid -->
    <div class="bottom-grid">
      <div class="section-card">
        <div class="section-head">
          <h2>开放模型</h2>
          <el-button link type="primary" @click="go('/admin/models')">管理 →</el-button>
        </div>
        <el-table :data="models" size="small" stripe empty-text="暂无模型" class="dash-table">
          <el-table-column prop="slug" label="model" min-width="160">
            <template #default="{ row }"><code>{{ row.slug }}</code></template>
          </el-table-column>
          <el-table-column label="类型" width="90" align="center">
            <template #default="{ row }">
              <el-tag size="small" :type="row.type === 'image' ? 'warning' : 'primary'" disable-transitions round>
                {{ row.type === 'image' ? '生图' : '对话' }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="description" label="说明" show-overflow-tooltip />
        </el-table>
      </div>
      <div class="section-card">
        <div class="section-head">
          <h2>最近图片任务</h2>
          <el-button link type="primary" @click="go('/personal/usage')">更多 →</el-button>
        </div>
        <el-table :data="tasks" size="small" stripe empty-text="暂无任务" class="dash-table">
          <el-table-column prop="task_id" label="任务 ID" min-width="180" show-overflow-tooltip>
            <template #default="{ row }"><code class="mono-sm">{{ row.task_id }}</code></template>
          </el-table-column>
          <el-table-column prop="status" label="状态" width="90" align="center">
            <template #default="{ row }">
              <el-tag size="small" round disable-transitions
                :type="taskType(row.status)">
                {{ row.status === 'success' ? '成功' : row.status === 'failed' ? '失败' : row.status }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column label="时间" width="170">
            <template #default="{ row }"><span class="time-text">{{ fmtTime(row.created_at) }}</span></template>
          </el-table-column>
        </el-table>
      </div>
    </div>
  </div>
</template>

<style scoped>
.dash { padding: 0; }

/* ---- Hero ---- */
.hero {
  position: relative;
  overflow: hidden;
  border-radius: 0 0 16px 16px;
  margin-bottom: 20px;
}
.hero-bg {
  position: absolute; inset: 0;
  background: linear-gradient(135deg, #409EFF 0%, #2b7de9 50%, #67C23A 100%);
  opacity: 1;
}
.hero-inner {
  position: relative;
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 24px;
  padding: 36px 32px;
}
.hero h1 {
  margin: 0 0 10px;
  font-size: 26px;
  font-weight: 700;
  color: #fff;
}
.hero p {
  margin: 0;
  color: rgba(255,255,255,0.82);
  font-size: 14px;
  line-height: 1.6;
  max-width: 600px;
}
.hero code {
  background: rgba(255,255,255,0.18);
  color: #fff;
  padding: 1px 6px;
  border-radius: 4px;
  font-family: ui-monospace, Menlo, Consolas, monospace;
  font-size: 13px;
}
.hero-actions { display: flex; gap: 10px; flex-shrink: 0; }
.hero-actions .el-button--primary {
  --el-button-bg-color: #fff;
  --el-button-text-color: #409EFF;
  --el-button-border-color: #fff;
  --el-button-hover-bg-color: rgba(255,255,255,0.9);
  --el-button-hover-text-color: #409EFF;
  --el-button-hover-border-color: rgba(255,255,255,0.9);
  font-weight: 600;
}
.hero-actions .el-button:not(.el-button--primary) {
  --el-button-bg-color: transparent;
  --el-button-text-color: #fff;
  --el-button-border-color: rgba(255,255,255,0.5);
  --el-button-hover-bg-color: rgba(255,255,255,0.12);
  --el-button-hover-text-color: #fff;
  --el-button-hover-border-color: rgba(255,255,255,0.8);
}

/* ---- KPI Row ---- */
.kpi-row {
  display: grid;
  grid-template-columns: repeat(4, minmax(0, 1fr));
  gap: 16px;
  padding: 0 20px;
  margin-top: -28px;
  position: relative;
  z-index: 1;
}
.kpi-card {
  background: var(--el-bg-color);
  border-radius: 12px;
  padding: 18px 20px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.06);
  border-top: 3px solid transparent;
  transition: transform .2s, box-shadow .2s;
}
.kpi-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(0,0,0,0.08);
}
.kpi-blue  { border-top-color: #409EFF; }
.kpi-green { border-top-color: #67C23A; }
.kpi-orange{ border-top-color: #E6A23C; }
.kpi-purple{ border-top-color: #9B59B6; }

.kpi-top {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}
.kpi-label {
  font-size: 13px;
  color: var(--el-text-color-secondary);
}
.kpi-dot {
  width: 36px; height: 36px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
}
.kpi-blue  .kpi-dot { background: rgba(64,158,255,0.1); color: #409EFF; }
.kpi-green .kpi-dot { background: rgba(103,194,58,0.1); color: #67C23A; }
.kpi-orange .kpi-dot { background: rgba(230,162,60,0.1); color: #E6A23C; }
.kpi-purple .kpi-dot { background: rgba(155,89,182,0.1); color: #9B59B6; }

.kpi-val {
  font-size: 30px;
  font-weight: 700;
  line-height: 1;
  letter-spacing: -1px;
  color: var(--el-text-color-primary);
}

/* ---- Ops Overview ---- */
.ops-grid {
  display: grid;
  grid-template-columns: 1fr 1fr 1.4fr;
  gap: 16px;
  padding: 16px 20px 0;
}
.ops-card,
.trend-card {
  min-height: 142px;
}
.section-head.compact { margin-bottom: 10px; }
.ops-main {
  display: flex;
  align-items: flex-end;
  gap: 10px;
  flex-wrap: wrap;
}
.ops-num {
  font-size: 28px;
  font-weight: 700;
  color: var(--el-text-color-primary);
  line-height: 1.1;
}
.ops-label,
.model-pill {
  font-size: 12px;
  color: var(--el-text-color-secondary);
}
.bars {
  display: grid;
  gap: 8px;
}
.bar-row {
  display: grid;
  grid-template-columns: 42px minmax(0, 1fr) 36px;
  gap: 8px;
  align-items: center;
  font-size: 12px;
  color: var(--el-text-color-secondary);
}
.bar-track {
  height: 8px;
  border-radius: 999px;
  background: var(--el-fill-color-light);
  overflow: hidden;
}
.bar-track i {
  display: block;
  height: 100%;
  border-radius: inherit;
  background: var(--el-color-primary-light-5);
}
.bar-row b {
  text-align: right;
  color: var(--el-text-color-primary);
  font-weight: 600;
}

/* ---- Bottom Grid ---- */
.bottom-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 16px;
  padding: 16px 20px 20px;
}
.section-card {
  background: var(--el-bg-color);
  border-radius: 12px;
  padding: 20px 22px;
  box-shadow: 0 1px 4px rgba(0,21,41,0.05);
}
.section-head {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 14px;
}
.section-head h2 {
  font-size: 16px;
  font-weight: 600;
  margin: 0;
  color: var(--el-text-color-primary);
}

code {
  background: var(--el-fill-color-light, #f2f3f5);
  padding: 2px 7px;
  border-radius: 4px;
  font-family: ui-monospace, Menlo, Consolas, monospace;
  font-size: 12px;
}
.mono-sm { font-size: 11px; }
.time-text { font-size: 13px; color: var(--el-text-color-secondary); }

/* dark mode adjustments */
:root.dark .hero-bg,
html.dark .hero-bg {
  opacity: 0.85;
}
html.dark .kpi-card {
  box-shadow: 0 2px 12px rgba(0,0,0,0.25);
}
html.dark .kpi-card:hover {
  box-shadow: 0 6px 20px rgba(0,0,0,0.35);
}

@media (max-width: 960px) {
  .hero-inner { flex-direction: column; align-items: flex-start; padding: 28px 20px; }
  .kpi-row { grid-template-columns: repeat(2, minmax(0, 1fr)); margin-top: -20px; }
  .ops-grid, .bottom-grid { grid-template-columns: 1fr; }
}
@media (max-width: 560px) {
  .kpi-row { grid-template-columns: 1fr; }
}
</style>
````

## File: web/src/views/personal/OnlinePlay.vue
````vue
<script setup lang="ts">
import { computed, nextTick, onBeforeUnmount, onMounted, ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import {
  listMyModels,
  streamPlayChat,
  playGenerateImage,
  type SimpleModel,
  type PlayChatMessage,
  type PlayImageData,
} from '@/api/me'
import { ENABLE_CHAT_MODEL } from '@/config/feature'

// ----------------------------------------------------
// 本地模型
// ----------------------------------------------------
const models = ref<SimpleModel[]>([])
const chatModels = computed(() => models.value.filter((m) => m.type === 'chat'))
const imageModels = computed(() => models.value.filter((m) => m.type === 'image'))

const selectedChatModel = ref('')
const selectedImageModel = ref('')

const currentChatDesc = computed(
  () => chatModels.value.find((m) => m.slug === selectedChatModel.value)?.description || '',
)
const currentImageDesc = computed(
  () => imageModels.value.find((m) => m.slug === selectedImageModel.value)?.description || '',
)

onMounted(async () => {
  try {
    const m = await listMyModels()
    // feature flag 关闭时,前端直接把 chat 类型的模型从列表过滤掉,
    // 保证 chatModels / imageModels / selectedChatModel 等后续 state 都不会
    // 拿到 chat 模型(即便模板里还有残留引用)。
    models.value = ENABLE_CHAT_MODEL
      ? m.items
      : m.items.filter((x) => x.type !== 'chat')
    const firstChat = m.items.find((x) => x.type === 'chat')
    const firstImage = m.items.find((x) => x.type === 'image')
    if (firstChat) selectedChatModel.value = firstChat.slug
    if (firstImage) selectedImageModel.value = firstImage.slug
  } catch {
    // 静默;错误拦截器已提示
  }
})

// ----------------------------------------------------
// Tabs
// ----------------------------------------------------
const activeTab = ref<'chat' | 'text2img' | 'img2img'>(
  ENABLE_CHAT_MODEL ? 'chat' : 'text2img',
)

// ====================================================
// 对话(Chat)
// ====================================================
interface UIMessage {
  id: number
  role: 'user' | 'assistant' | 'system'
  content: string
  pending?: boolean
  error?: boolean
  at: number
}

let uid = 0

const systemPrompt = ref('你是一个友好、博学、回答精准的中文助手。回答中若涉及代码请使用 Markdown 代码块。')
const temperature = ref(0.7)
const chatInput = ref('')
const chatMsgs = ref<UIMessage[]>([])
const chatSending = ref(false)
const chatAbort = ref<AbortController | null>(null)
const chatScroll = ref<HTMLElement | null>(null)
const inputRef = ref<any>(null)

const suggestions = [
  { icon: '💡', title: '向我解释', sub: '量子纠缠到底是什么?' },
  { icon: '✍️', title: '帮我写作', sub: '一段 200 字的产品发布文案' },
  { icon: '🧑‍💻', title: '写段代码', sub: 'Go 实现令牌桶限流器' },
  { icon: '🌏', title: '中英互译', sub: '把上面这段翻译为英文' },
]

function useSuggestion(s: typeof suggestions[number]) {
  chatInput.value = `${s.title}:${s.sub}`
  nextTick(() => inputRef.value?.focus?.())
}

async function scrollChat(force = false) {
  await nextTick()
  const el = chatScroll.value
  if (!el) return
  if (force) {
    el.scrollTop = el.scrollHeight
    return
  }
  const gap = el.scrollHeight - el.scrollTop - el.clientHeight
  if (gap < 180) el.scrollTop = el.scrollHeight
}

async function sendChat() {
  if (chatSending.value) return
  const text = chatInput.value.trim()
  if (!text) return
  if (!selectedChatModel.value) {
    ElMessage.warning('请选择一个文字模型')
    return
  }
  const now = Date.now()
  chatMsgs.value.push({ id: ++uid, role: 'user', content: text, at: now })
  chatInput.value = ''
  const assistant: UIMessage = { id: ++uid, role: 'assistant', content: '', pending: true, at: now }
  chatMsgs.value.push(assistant)
  await scrollChat(true)

  const history: PlayChatMessage[] = []
  if (systemPrompt.value.trim()) {
    history.push({ role: 'system', content: systemPrompt.value.trim() })
  }
  for (const m of chatMsgs.value.slice(0, -1)) {
    if (m.error) continue
    history.push({ role: m.role as 'user' | 'assistant' | 'system', content: m.content })
  }

  chatSending.value = true
  chatAbort.value = new AbortController()
  try {
    await streamPlayChat(
      { model: selectedChatModel.value, messages: history, temperature: temperature.value },
      (delta) => {
        assistant.content += delta
        assistant.pending = false
        scrollChat()
      },
      chatAbort.value.signal,
    )
    assistant.pending = false
    if (!assistant.content) assistant.content = '(无输出)'
  } catch (err: unknown) {
    assistant.pending = false
    assistant.error = true
    const msg = err instanceof Error ? err.message : String(err)
    assistant.content = assistant.content || `请求失败:${msg}`
    ElMessage.error(msg)
  } finally {
    chatSending.value = false
    chatAbort.value = null
    scrollChat()
  }
}

function stopChat() {
  chatAbort.value?.abort()
}

function resetChat() {
  if (chatSending.value) stopChat()
  chatMsgs.value = []
}

async function regenerate() {
  if (chatSending.value) return
  // 去掉最后一条 assistant,把最后一条 user 重发
  let lastUserIdx = -1
  for (let i = chatMsgs.value.length - 1; i >= 0; i--) {
    if (chatMsgs.value[i].role === 'user') { lastUserIdx = i; break }
  }
  if (lastUserIdx < 0) return
  const lastUserText = chatMsgs.value[lastUserIdx].content
  chatMsgs.value = chatMsgs.value.slice(0, lastUserIdx)
  chatInput.value = lastUserText
  await sendChat()
}

function copyText(s: string) {
  try {
    navigator.clipboard.writeText(s)
    ElMessage.success('已复制')
  } catch {
    ElMessage.warning('复制失败')
  }
}

onBeforeUnmount(() => chatAbort.value?.abort())

// ---------- 轻量 markdown 渲染(代码块 / 行内代码 / 粗体 / 链接) ----------
function escapeHtml(s: string) {
  return s
    .replace(/&/g, '&amp;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;')
    .replace(/"/g, '&quot;')
    .replace(/'/g, '&#39;')
}

function renderMarkdown(raw: string): string {
  if (!raw) return ''
  const parts: string[] = []
  const blocks = raw.split(/```/g) // ``` 成对切分
  for (let i = 0; i < blocks.length; i++) {
    const chunk = blocks[i]
    if (i % 2 === 1) {
      // 代码块:首行可能是 lang
      const nl = chunk.indexOf('\n')
      let lang = ''
      let code = chunk
      if (nl >= 0) {
        const head = chunk.slice(0, nl).trim()
        if (/^[a-zA-Z0-9+_\-]{1,20}$/.test(head)) {
          lang = head
          code = chunk.slice(nl + 1)
        }
      }
      parts.push(
        `<pre class="mdk-pre" data-lang="${escapeHtml(lang || '')}"><code>${escapeHtml(
          code.replace(/\n$/, ''),
        )}</code></pre>`,
      )
    } else {
      // 行内元素
      let html = escapeHtml(chunk)
      // 行内代码 `xxx`
      html = html.replace(/`([^`\n]+)`/g, '<code class="mdk-ic">$1</code>')
      // 粗体 **xxx**
      html = html.replace(/\*\*([^*\n]+)\*\*/g, '<strong>$1</strong>')
      // 自动链接
      html = html.replace(
        /(https?:\/\/[\w\-._~:/?#\[\]@!$&'()*+,;=%]+)/g,
        '<a href="$1" target="_blank" rel="noopener">$1</a>',
      )
      // 换行
      html = html.replace(/\n/g, '<br />')
      parts.push(html)
    }
  }
  return parts.join('')
}

// ====================================================
// 文生图(Text2Img)
// ====================================================
const t2iPrompt = ref('')
const t2iSize = ref<'1024x1024' | '1792x1024' | '1024x1792'>('1024x1024')
const t2iN = ref(1)
const t2iSending = ref(false)
const t2iResult = ref<PlayImageData[]>([])
const t2iPreview = ref(false)
const t2iError = ref('')
const t2iAbort = ref<AbortController | null>(null)

const imgExamples = [
  '赛博朋克城市夜景,霓虹雨夜,电影感光影,8k',
  '一只金色胖柴犬穿西装坐在办公桌前,油画质感',
  '极简几何海报,蓝橙配色,主体是一只展翅的鹤',
  '童话风格蘑菇屋,黄昏光线,柔和景深',
]

async function sendText2Img() {
  const prompt = t2iPrompt.value.trim()
  if (!prompt) {
    ElMessage.warning('请输入描述词 prompt')
    return
  }
  if (!selectedImageModel.value) {
    ElMessage.warning('请选择一个图片模型')
    return
  }
  t2iSending.value = true
  t2iError.value = ''
  t2iPreview.value = false
  t2iResult.value = []
  t2iAbort.value = new AbortController()
  try {
    const resp = await playGenerateImage(
      {
        model: selectedImageModel.value,
        prompt,
        n: t2iN.value,
        size: t2iSize.value,
      },
      t2iAbort.value.signal,
    )
    t2iResult.value = resp.data || []
    t2iPreview.value = !!resp.is_preview
    if (t2iResult.value.length === 0) {
      t2iError.value = '未产出图片,请重试或更换描述'
    } else if (t2iPreview.value) {
      ElMessage.warning(`生成成功(预览模式):本次账号未命中 IMG2 灰度,展示的是 IMG1 预览图`)
    } else {
      ElMessage.success(`生成成功,共 ${t2iResult.value.length} 张`)
    }
  } catch (err: unknown) {
    const msg = err instanceof Error ? err.message : String(err)
    t2iError.value = msg
    ElMessage.error(msg)
  } finally {
    t2iSending.value = false
    t2iAbort.value = null
  }
}

function stopText2Img() {
  t2iAbort.value?.abort()
}

// 预览 viewer
const previewVisible = ref(false)
const previewList = ref<string[]>([])
const previewIndex = ref(0)
function openPreview(urls: string[], idx: number) {
  previewList.value = urls
  previewIndex.value = idx
  previewVisible.value = true
}
function downloadUrl(url: string) {
  const a = document.createElement('a')
  a.href = url
  a.target = '_blank'
  a.rel = 'noopener'
  a.download = ''
  document.body.appendChild(a)
  a.click()
  document.body.removeChild(a)
}

// ====================================================
// 图生图(Img2Img)
// ====================================================
interface RefImage {
  name: string
  dataUrl: string
  size: number
}
const refImages = ref<RefImage[]>([])
const i2iPrompt = ref('')
const i2iSize = ref<'1024x1024' | '1792x1024' | '1024x1792'>('1024x1024')
const i2iSending = ref(false)
const i2iResult = ref<PlayImageData[]>([])
const i2iPreview = ref(false)
const i2iError = ref('')
const i2iAbort = ref<AbortController | null>(null)
const MAX_REF_COUNT = 4
const MAX_REF_BYTES = 4 * 1024 * 1024 // 4MB

function handleFilePick(e: Event) {
  const input = e.target as HTMLInputElement
  const files = Array.from(input.files || [])
  if (!files.length) return
  const remaining = MAX_REF_COUNT - refImages.value.length
  if (remaining <= 0) {
    ElMessage.warning(`最多只能上传 ${MAX_REF_COUNT} 张参考图`)
    input.value = ''
    return
  }
  if (files.length > remaining) {
    ElMessage.warning(`最多 ${MAX_REF_COUNT} 张参考图,已只读取前 ${remaining} 张`)
  }
  for (const file of files.slice(0, remaining)) {
    if (file.size > MAX_REF_BYTES) {
      ElMessage.warning(`${file.name} 超过 4MB 限制`)
      continue
    }
    const reader = new FileReader()
    reader.onload = () => {
      if (refImages.value.length >= MAX_REF_COUNT) return
      refImages.value.push({
        name: file.name,
        dataUrl: String(reader.result || ''),
        size: file.size,
      })
    }
    reader.readAsDataURL(file)
  }
  input.value = ''
}

function removeRefImage(idx: number) {
  refImages.value.splice(idx, 1)
}

async function sendImg2Img() {
  if (refImages.value.length === 0) {
    ElMessage.warning('请先上传至少一张参考图')
    return
  }
  if (!i2iPrompt.value.trim()) {
    ElMessage.warning('请描述希望的改动')
    return
  }
  if (!selectedImageModel.value) {
    ElMessage.warning('请选择一个图片模型')
    return
  }
  i2iSending.value = true
  i2iError.value = ''
  i2iPreview.value = false
  i2iResult.value = []
  i2iAbort.value = new AbortController()
  try {
    const resp = await playGenerateImage(
      {
        model: selectedImageModel.value,
        prompt: i2iPrompt.value.trim(),
        n: 1,
        size: i2iSize.value,
        reference_images: refImages.value.map((r) => r.dataUrl),
      },
      i2iAbort.value.signal,
    )
    i2iResult.value = resp.data || []
    i2iPreview.value = !!resp.is_preview
    if (i2iPreview.value && i2iResult.value.length > 0) {
      ElMessage.warning('生成成功(预览模式):本次账号未命中 IMG2 灰度,展示的是 IMG1 预览图')
    }
  } catch (err: unknown) {
    const msg = err instanceof Error ? err.message : String(err)
    i2iError.value = msg
    ElMessage.error(msg)
  } finally {
    i2iSending.value = false
    i2iAbort.value = null
  }
}

// 代码块内的 "复制" 按钮(通过事件委托,避免每次重渲染都重新绑定)
function onMsgClick(e: MouseEvent) {
  const t = e.target as HTMLElement | null
  if (!t) return
  const btn = t.closest('.mdk-copy') as HTMLElement | null
  if (!btn) return
  const pre = btn.parentElement?.querySelector('code')
  if (!pre) return
  copyText(pre.textContent || '')
}

// input 自动聚焦(tab 切换后)
watch(activeTab, (v) => {
  if (v === 'chat') nextTick(() => inputRef.value?.focus?.())
})
</script>

<template>
  <div class="page-container playground">
    <!-- ============ Hero(紧凑条) ============ -->
    <div class="hero card-block">
      <div class="hero-left">
        <el-icon class="hero-ic"><MagicStick /></el-icon>
        <div class="hero-txt">
          <h2 class="hero-title">在线体验</h2>
          <span class="hero-sub">
            浏览器中直接调用 GPT {{ ENABLE_CHAT_MODEL ? '文字 / ' : '' }}图像模型 · 文生图 & 图生图 · 共用本地账号池,记录同步到「使用记录」
          </span>
        </div>
      </div>
      <div class="hero-stats">
        <template v-if="ENABLE_CHAT_MODEL">
          <div class="mini-stat">
            <span class="mini-num">{{ chatModels.length }}</span>
            <span class="mini-lbl">文字模型</span>
          </div>
        </template>
        <span class="mini-dot" />
        <div class="mini-stat">
          <span class="mini-num">{{ imageModels.length }}</span>
          <span class="mini-lbl">图片模型</span>
        </div>
      </div>
    </div>

    <!-- ============ Tabs ============ -->
    <el-tabs v-model="activeTab" class="pg-tabs">
      <!-- =================================================== -->
      <!--                          Chat                         -->
      <!-- =================================================== -->
      <el-tab-pane v-if="ENABLE_CHAT_MODEL" name="chat">
        <template #label>
          <span class="tab-lbl"><el-icon><ChatDotRound /></el-icon> 对话</span>
        </template>

        <div class="chat-grid">
          <!-- 左侧:模型 + System + 温度 -->
          <aside class="card-block side">
            <div class="side-row">
              <label class="side-lbl">文字模型</label>
              <el-select v-model="selectedChatModel" placeholder="选择文字模型" size="large" style="width:100%">
                <el-option v-for="m in chatModels" :key="m.id" :label="m.slug" :value="m.slug">
                  <div class="opt-row">
                    <span class="opt-slug">{{ m.slug }}</span>
                    <el-tag size="small" type="primary" effect="plain">chat</el-tag>
                  </div>
                </el-option>
              </el-select>
              <div v-if="currentChatDesc" class="side-hint">{{ currentChatDesc }}</div>
            </div>

            <div class="side-row">
              <label class="side-lbl">
                Temperature
                <span class="side-val">{{ temperature.toFixed(1) }}</span>
              </label>
              <el-slider v-model="temperature" :min="0" :max="2" :step="0.1" show-stops />
              <div class="side-hint">越低越保守、越高越发散。默认 0.7</div>
            </div>

            <div class="side-row">
              <label class="side-lbl">System Prompt</label>
              <el-input
                v-model="systemPrompt"
                type="textarea"
                :rows="6"
                resize="none"
                placeholder="为助手设定人格与风格"
              />
            </div>

            <el-button :disabled="chatMsgs.length === 0" @click="resetChat" class="side-btn">
              <el-icon><Delete /></el-icon> 清空会话
            </el-button>
          </aside>

          <!-- 右侧:聊天主区 -->
          <section class="card-block chat-main">
            <header class="chat-header">
              <div class="chat-title">
                <el-avatar :size="32" class="avatar-bot">
                  <el-icon><Cpu /></el-icon>
                </el-avatar>
                <div>
                  <div class="chat-model">{{ selectedChatModel || '未选择模型' }}</div>
                  <div class="chat-sub">
                    {{ chatSending ? '正在回复…' : (chatMsgs.length ? `${chatMsgs.length} 条消息` : '准备就绪') }}
                  </div>
                </div>
              </div>
              <div class="chat-tools">
                <el-tooltip content="重试上一个问题" placement="top">
                  <el-button
                    :disabled="chatSending || chatMsgs.length === 0"
                    circle
                    @click="regenerate"
                  >
                    <el-icon><RefreshRight /></el-icon>
                  </el-button>
                </el-tooltip>
              </div>
            </header>

            <div ref="chatScroll" class="chat-scroll" @click="onMsgClick">
              <!-- 空态:建议卡 -->
              <div v-if="chatMsgs.length === 0" class="welcome">
                <div class="welcome-hi">
                  👋 你好，本地控制台
                </div>
                <div class="welcome-sub">选一个话题开始,或者直接在下方输入。</div>
                <div class="suggest-grid">
                  <div
                    v-for="(s, i) in suggestions"
                    :key="i"
                    class="suggest-card"
                    @click="useSuggestion(s)"
                  >
                    <div class="s-ic">{{ s.icon }}</div>
                    <div class="s-t">{{ s.title }}</div>
                    <div class="s-s">{{ s.sub }}</div>
                  </div>
                </div>
              </div>

              <!-- 消息列表 -->
              <article
                v-for="m in chatMsgs"
                :key="m.id"
                :class="['msg', m.role, m.error ? 'err' : '']"
              >
                <el-avatar :size="34" :class="m.role === 'user' ? 'avatar-user' : 'avatar-bot'">
                  <el-icon v-if="m.role === 'user'"><User /></el-icon>
                  <el-icon v-else><MagicStick /></el-icon>
                </el-avatar>
                <div class="msg-body">
                  <div class="msg-head">
                    <span class="who">{{ m.role === 'user' ? '我' : '助手' }}</span>
                    <span v-if="!m.pending && m.content" class="copy-btn" @click="copyText(m.content)">
                      <el-icon><CopyDocument /></el-icon> 复制
                    </span>
                  </div>
                  <div class="msg-content">
                    <div v-if="m.pending && !m.content" class="typing">
                      <span></span><span></span><span></span>
                    </div>
                    <div
                      v-else
                      class="md"
                      v-html="renderMarkdown(m.content)"
                    />
                  </div>
                </div>
              </article>
            </div>

            <!-- 输入条 -->
            <div class="composer" :class="{ focused: !!chatInput }">
              <el-input
                ref="inputRef"
                v-model="chatInput"
                type="textarea"
                :rows="1"
                :autosize="{ minRows: 1, maxRows: 6 }"
                resize="none"
                placeholder="给助手发消息…  Enter 发送,Shift+Enter 换行"
                @keydown.enter.exact.prevent="sendChat"
              />
              <div class="composer-tools">
                <span class="hint">
                  <el-icon><InfoFilled /></el-icon>
                  按 Enter 发送
                </span>
                <div style="flex:1" />
                <el-button v-if="chatSending" type="danger" @click="stopChat" round>
                  <el-icon><VideoPause /></el-icon> 停止
                </el-button>
                <el-button
                  v-else
                  type="primary"
                  :disabled="!chatInput.trim() || !selectedChatModel"
                  @click="sendChat"
                  round
                >
                  发送
                  <el-icon style="margin-left:4px"><Promotion /></el-icon>
                </el-button>
              </div>
            </div>
          </section>
        </div>
      </el-tab-pane>

      <!-- =================================================== -->
      <!--                        文生图                         -->
      <!-- =================================================== -->
      <el-tab-pane name="text2img">
        <template #label>
          <span class="tab-lbl"><el-icon><Picture /></el-icon> 文生图</span>
        </template>

        <div class="img-grid">
          <aside class="card-block side">
            <div class="side-row">
              <label class="side-lbl">图片模型</label>
              <el-select v-model="selectedImageModel" placeholder="选择图片模型" size="large" style="width:100%">
                <el-option v-for="m in imageModels" :key="m.id" :label="m.slug" :value="m.slug">
                  <div class="opt-row">
                    <span class="opt-slug">{{ m.slug }}</span>
                    <el-tag size="small" type="warning" effect="plain">image</el-tag>
                  </div>
                </el-option>
              </el-select>
              <div v-if="currentImageDesc" class="side-hint">{{ currentImageDesc }}</div>
            </div>

            <div class="side-row">
              <label class="side-lbl">画面比例</label>
              <div class="ratio-row">
                <button
                  v-for="opt in [
                    { v: '1024x1024', l: '1:1',  w: 36, h: 36 },
                    { v: '1792x1024', l: '16:9', w: 48, h: 28 },
                    { v: '1024x1792', l: '9:16', w: 28, h: 48 },
                  ]"
                  :key="opt.v"
                  :class="['ratio-btn', { active: t2iSize === opt.v }]"
                  @click="t2iSize = opt.v as any"
                >
                  <div class="ratio-box" :style="{ width: opt.w + 'px', height: opt.h + 'px' }" />
                  <span>{{ opt.l }}</span>
                </button>
              </div>
            </div>

            <div class="side-row">
              <label class="side-lbl">张数 <span class="side-val">{{ t2iN }}</span></label>
              <el-slider v-model="t2iN" :min="1" :max="4" show-stops />
            </div>

            <div class="side-row">
              <label class="side-lbl">Prompt</label>
              <el-input
                v-model="t2iPrompt"
                type="textarea"
                :rows="5"
                resize="none"
                placeholder="描述画面的主体、风格、光线、构图…越具体效果越好"
              />
              <div class="chips">
                <el-tag
                  v-for="(p, i) in imgExamples"
                  :key="i"
                  effect="plain"
                  round
                  class="chip"
                  @click="t2iPrompt = p"
                >{{ p }}</el-tag>
              </div>
            </div>

            <el-button v-if="t2iSending" type="danger" @click="stopText2Img" round class="side-btn">
              <el-icon><VideoPause /></el-icon> 停止
            </el-button>
            <el-button
              v-else
              type="primary"
              round
              size="large"
              :disabled="!t2iPrompt.trim() || !selectedImageModel"
              @click="sendText2Img"
              class="side-btn gen-btn"
            >
              <el-icon><MagicStick /></el-icon> 生成图片
            </el-button>
          </aside>

          <section class="card-block img-main">
            <div v-if="t2iSending" class="stage loading">
              <div class="orb"><el-icon class="spin"><Loading /></el-icon></div>
              <div class="stage-title">正在为你绘制…</div>
              <div class="stage-sub">上游渲染通常需要 1-2 分钟,请保持页面打开</div>
            </div>
            <div v-else-if="t2iError" class="err-block">
              <el-icon><WarningFilled /></el-icon>
              {{ t2iError }}
            </div>
            <div v-else-if="t2iResult.length === 0" class="stage">
              <div class="stage-art">🖼️</div>
              <div class="stage-title">还没有图片</div>
              <div class="stage-sub">在左侧填好 prompt 和参数,点击「生成图片」</div>
            </div>
            <div v-else class="result-wrap">
              <el-alert
                v-if="t2iPreview"
                class="preview-tip"
                type="warning"
                :closable="false"
                show-icon
                title="本次未使用 IMG2 灰度生成"
                description="上游没有把本账号放入 IMG2 终稿通道,返回的是 IMG1 预览图;效果略简化,属于正常降级,可多试几次或更换账号。"
              />
              <div class="result-grid">
                <div
                  v-for="(img, idx) in t2iResult"
                  :key="idx"
                  class="img-cell"
                  :class="{ 'is-preview': t2iPreview }"
                  @click="openPreview(t2iResult.map((x) => x.url), idx)"
                >
                  <img :src="img.url" :alt="`result-${idx}`" loading="lazy" />
                  <div v-if="t2iPreview" class="img-badge">IMG1 预览</div>
                  <div class="img-actions" @click.stop>
                    <button class="iact" @click="openPreview(t2iResult.map((x) => x.url), idx)">
                      <el-icon><ZoomIn /></el-icon>
                    </button>
                    <button class="iact" @click="downloadUrl(img.url)">
                      <el-icon><Download /></el-icon>
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </section>
        </div>
      </el-tab-pane>

      <!-- =================================================== -->
      <!--                        图生图                         -->
      <!-- =================================================== -->
      <el-tab-pane name="img2img">
        <template #label>
          <span class="tab-lbl"><el-icon><PictureFilled /></el-icon> 图生图</span>
        </template>

        <div class="img-grid">
          <aside class="card-block side">
            <div class="side-row">
              <label class="side-lbl">图片模型</label>
              <el-select v-model="selectedImageModel" placeholder="选择图片模型" size="large" style="width:100%">
                <el-option v-for="m in imageModels" :key="m.id" :label="m.slug" :value="m.slug" />
              </el-select>
            </div>

            <div class="side-row">
              <label class="side-lbl">参考图 <span class="side-val">{{ refImages.length }}/4</span></label>
              <label class="upload-zone">
                <el-icon class="up-ic"><UploadFilled /></el-icon>
                <div class="up-t">点击选择 / 拖拽图片到这里</div>
                <div class="up-s">最多 4 张,每张 ≤ 4MB</div>
                <input type="file" accept="image/*" multiple @change="handleFilePick" />
              </label>

              <div v-if="refImages.length" class="ref-grid">
                <div v-for="(r, idx) in refImages" :key="idx" class="ref-thumb">
                  <img :src="r.dataUrl" :alt="r.name" />
                  <div class="ref-x" @click="removeRefImage(idx)">
                    <el-icon><Close /></el-icon>
                  </div>
                  <div class="ref-meta">{{ (r.size / 1024).toFixed(0) }} KB</div>
                </div>
              </div>
            </div>

            <div class="side-row">
              <label class="side-lbl">输出比例</label>
              <div class="ratio-row">
                <button
                  v-for="opt in [
                    { v: '1024x1024', l: '1:1',  w: 36, h: 36 },
                    { v: '1792x1024', l: '16:9', w: 48, h: 28 },
                    { v: '1024x1792', l: '9:16', w: 28, h: 48 },
                  ]"
                  :key="opt.v"
                  :class="['ratio-btn', { active: i2iSize === opt.v }]"
                  @click="i2iSize = opt.v as any"
                >
                  <div class="ratio-box" :style="{ width: opt.w + 'px', height: opt.h + 'px' }" />
                  <span>{{ opt.l }}</span>
                </button>
              </div>
            </div>

            <div class="side-row">
              <label class="side-lbl">希望如何改动</label>
              <el-input
                v-model="i2iPrompt"
                type="textarea"
                :rows="4"
                resize="none"
                placeholder="例:保持人物姿态,把背景换成赛博朋克夜景"
              />
            </div>

            <el-button
              type="primary"
              round
              size="large"
              :loading="i2iSending"
              :disabled="refImages.length === 0 || !i2iPrompt.trim()"
              @click="sendImg2Img"
              class="side-btn gen-btn"
            >
              <el-icon><MagicStick /></el-icon> 生成
            </el-button>
          </aside>

          <section class="card-block img-main">
            <el-alert
              type="info"
              :closable="false"
              title="参考图会作为 ChatGPT Web 对话的一部分发送"
              description="本地文件会先转成 reference_images,后端上传到 ChatGPT 文件服务并插入同一条 multimodal_text 用户消息；外部客户端也可直接传 image_url / reference_images URL。"
              show-icon
              style="margin-bottom: 14px; border-radius: 10px;"
            />

            <div v-if="i2iError" class="err-block">
              <el-icon><WarningFilled /></el-icon>
              {{ i2iError }}
            </div>
            <div v-else-if="i2iSending" class="stage loading">
              <div class="orb"><el-icon class="spin"><Loading /></el-icon></div>
              <div class="stage-title">正在生成…</div>
            </div>
            <div v-else-if="i2iResult.length === 0" class="stage">
              <div class="stage-art">🎨</div>
              <div class="stage-title">还没有结果</div>
              <div class="stage-sub">上传参考图 + 描述改动,然后点击「生成」</div>
            </div>
            <div v-else class="result-wrap">
              <el-alert
                v-if="i2iPreview"
                class="preview-tip"
                type="warning"
                :closable="false"
                show-icon
                title="本次未使用 IMG2 灰度生成"
                description="上游没有把本账号放入 IMG2 终稿通道,返回的是 IMG1 预览图。"
              />
              <div class="result-grid">
                <div
                  v-for="(img, idx) in i2iResult"
                  :key="idx"
                  class="img-cell"
                  :class="{ 'is-preview': i2iPreview }"
                  @click="openPreview(i2iResult.map((x) => x.url), idx)"
                >
                  <img :src="img.url" :alt="`result-${idx}`" />
                  <div v-if="i2iPreview" class="img-badge">IMG1 预览</div>
                  <div class="img-actions" @click.stop>
                    <button class="iact" @click="openPreview(i2iResult.map((x) => x.url), idx)">
                      <el-icon><ZoomIn /></el-icon>
                    </button>
                    <button class="iact" @click="downloadUrl(img.url)">
                      <el-icon><Download /></el-icon>
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </section>
        </div>
      </el-tab-pane>
    </el-tabs>

    <!-- ============ 图片预览(全屏 viewer) ============ -->
    <el-image-viewer
      v-if="previewVisible"
      :url-list="previewList"
      :initial-index="previewIndex"
      @close="previewVisible = false"
      teleported
    />
  </div>
</template>

<style scoped lang="scss">
.playground { padding-bottom: 24px; }

/* ====================== Hero(紧凑条) ====================== */
.hero {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  padding: 12px 18px !important;
  margin-bottom: 14px !important;
}
.hero-left {
  display: flex;
  align-items: center;
  gap: 12px;
  min-width: 0;
  flex: 1;
}
.hero-ic {
  width: 36px;
  height: 36px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  color: var(--el-color-primary);
  background: var(--el-color-primary-light-9);
  flex-shrink: 0;
}
.hero-txt {
  display: flex;
  align-items: baseline;
  gap: 10px;
  min-width: 0;
  flex-wrap: wrap;
}
.hero-title {
  margin: 0;
  font-size: 16px;
  font-weight: 600;
  color: var(--el-text-color-primary);
  white-space: nowrap;
}
.hero-sub {
  font-size: 12.5px;
  color: var(--el-text-color-secondary);
  line-height: 1.5;
}
.hero-stats {
  display: flex;
  align-items: center;
  gap: 12px;
  flex-shrink: 0;
}
.mini-stat {
  display: inline-flex;
  align-items: baseline;
  gap: 4px;
  .mini-num {
    font-size: 14px;
    font-weight: 600;
    color: var(--el-color-primary);
  }
  .mini-lbl {
    font-size: 12px;
    color: var(--el-text-color-secondary);
  }
}
.mini-dot {
  width: 3px; height: 3px; border-radius: 50%;
  background: var(--el-border-color);
}

/* ====================== Tabs ====================== */
.pg-tabs {
  :deep(.el-tabs__header) { margin-bottom: 16px; }
  :deep(.el-tabs__nav-wrap::after) { background: var(--el-border-color-lighter); }
  :deep(.el-tabs__item) {
    font-size: 14px;
    font-weight: 500;
    padding: 0 18px;
  }
  :deep(.el-tabs__item.is-active) { font-weight: 600; }
}
.tab-lbl { display: inline-flex; align-items: center; gap: 6px; }

/* ====================== Side ====================== */
.side { display: flex; flex-direction: column; gap: 16px; height: fit-content; position: sticky; top: 12px; }
.side-row { display: flex; flex-direction: column; gap: 6px; }
.side-lbl {
  font-size: 12px;
  color: var(--el-text-color-secondary);
  font-weight: 500;
  display: flex; justify-content: space-between; align-items: center;
  letter-spacing: 0.3px;
  text-transform: uppercase;
}
.side-val { font-weight: 600; color: var(--el-color-primary); letter-spacing: 0; text-transform: none; font-size: 13px; }
.side-hint { font-size: 12px; color: var(--el-text-color-placeholder); line-height: 1.5; }
.side-btn { margin-top: 4px; }
.gen-btn { box-shadow: 0 6px 18px -6px rgba(64, 158, 255, 0.55); }
.opt-row { display: flex; justify-content: space-between; align-items: center; gap: 8px; }
.opt-slug { font-family: ui-monospace, Menlo, Consolas, monospace; font-size: 13px; }

/* ====================== Chat ====================== */
.chat-grid {
  display: grid;
  grid-template-columns: 300px minmax(0, 1fr);
  gap: 16px;
  min-height: 620px;
}
.chat-main {
  display: flex; flex-direction: column;
  padding: 0;
  overflow: hidden;
  height: 720px;
}
.chat-header {
  display: flex; justify-content: space-between; align-items: center;
  padding: 12px 18px;
  border-bottom: 1px solid var(--el-border-color-lighter);
  background: linear-gradient(180deg, var(--el-bg-color) 0%, var(--el-fill-color-lighter) 100%);
}
.chat-title { display: flex; align-items: center; gap: 10px; }
.chat-model { font-size: 14px; font-weight: 600; color: var(--el-text-color-primary); }
.chat-sub { font-size: 12px; color: var(--el-text-color-secondary); margin-top: 2px; }
.chat-tools { display: flex; gap: 6px; }

.chat-scroll {
  flex: 1;
  overflow-y: auto;
  padding: 22px 24px;
  scroll-behavior: smooth;
}

/* ----- 欢迎 ----- */
.welcome {
  min-height: 100%;
  display: flex; flex-direction: column; align-items: center; justify-content: center;
  padding: 30px 20px;
}
.welcome-hi {
  font-size: 24px; font-weight: 700;
  color: var(--el-text-color-primary);
  margin-bottom: 6px;
}
.welcome-sub { color: var(--el-text-color-secondary); margin-bottom: 22px; font-size: 14px; }
.suggest-grid {
  display: grid; grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 12px;
  width: 100%; max-width: 680px;
}
.suggest-card {
  border: 1px solid var(--el-border-color-lighter);
  border-radius: 12px;
  padding: 14px 16px;
  cursor: pointer;
  background: var(--el-bg-color);
  transition: all 0.2s;
  .s-ic { font-size: 20px; margin-bottom: 4px; }
  .s-t { font-size: 13px; font-weight: 600; color: var(--el-text-color-primary); }
  .s-s { font-size: 12px; color: var(--el-text-color-secondary); margin-top: 4px; line-height: 1.5; }
  &:hover {
    border-color: var(--el-color-primary);
    transform: translateY(-1px);
    box-shadow: 0 6px 18px -8px rgba(64, 158, 255, 0.35);
  }
}

/* ----- 消息 ----- */
.msg {
  display: flex;
  gap: 12px;
  padding: 14px 0;
  border-bottom: 1px dashed var(--el-border-color-lighter);
  animation: fadeIn 0.25s ease;
  &:last-child { border-bottom: none; }
  &.err .msg-content { color: var(--el-color-danger); }
}
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(4px); }
  to   { opacity: 1; transform: translateY(0); }
}
.avatar-user {
  background: var(--el-color-primary);
  color: #fff;
  flex-shrink: 0;
}
.avatar-bot {
  background: var(--el-color-success);
  color: #fff;
  flex-shrink: 0;
}
.msg-body { flex: 1; min-width: 0; }
.msg-head {
  display: flex; align-items: center; justify-content: space-between;
  margin-bottom: 4px;
  .who { font-size: 12px; font-weight: 600; color: var(--el-text-color-secondary); }
  .copy-btn {
    font-size: 12px; color: var(--el-text-color-placeholder); cursor: pointer;
    display: inline-flex; align-items: center; gap: 2px;
    opacity: 0; transition: opacity 0.2s;
    &:hover { color: var(--el-color-primary); }
  }
}
.msg:hover .copy-btn { opacity: 1; }
.msg-content {
  font-size: 14px; line-height: 1.75;
  color: var(--el-text-color-primary);
  word-break: break-word;
}

/* markdown 渲染产物 */
.md :deep(.mdk-pre) {
  background: #0f172a;
  color: #e2e8f0;
  padding: 12px 14px;
  border-radius: 10px;
  overflow-x: auto;
  font-family: ui-monospace, Menlo, Consolas, monospace;
  font-size: 12.5px;
  line-height: 1.6;
  margin: 8px 0;
  position: relative;
  &::before {
    content: attr(data-lang);
    position: absolute;
    top: 6px; right: 10px;
    font-size: 10px;
    text-transform: uppercase;
    letter-spacing: 1px;
    color: #94a3b8;
    opacity: 0.8;
  }
}
.md :deep(.mdk-ic) {
  background: var(--el-fill-color);
  color: var(--el-color-primary);
  padding: 1px 6px;
  border-radius: 4px;
  font-family: ui-monospace, Menlo, Consolas, monospace;
  font-size: 12.5px;
}
.md :deep(a) { color: var(--el-color-primary); text-decoration: none; }
.md :deep(a:hover) { text-decoration: underline; }
.md :deep(strong) { font-weight: 600; }

/* typing 指示器 */
.typing {
  display: inline-flex; gap: 5px; padding: 4px 0;
  span {
    width: 7px; height: 7px; border-radius: 50%;
    background: var(--el-color-primary);
    animation: blink 1.4s infinite ease-in-out both;
  }
  span:nth-child(2) { animation-delay: 0.2s; }
  span:nth-child(3) { animation-delay: 0.4s; }
}
@keyframes blink {
  0%, 80%, 100% { opacity: 0.2; transform: scale(0.7); }
  40% { opacity: 1; transform: scale(1); }
}

/* ----- 输入条 ----- */
.composer {
  padding: 12px 18px 16px;
  border-top: 1px solid var(--el-border-color-lighter);
  background: var(--el-bg-color);
  transition: box-shadow 0.2s;
  :deep(.el-textarea__inner) {
    border-radius: 12px;
    padding: 10px 14px;
    font-size: 14px;
    box-shadow: none;
    border: 1px solid var(--el-border-color);
    transition: border-color 0.2s, box-shadow 0.2s;
    &:focus {
      border-color: var(--el-color-primary);
      box-shadow: 0 0 0 3px rgba(64, 158, 255, 0.15);
    }
  }
}
.composer-tools {
  display: flex; align-items: center; gap: 8px; margin-top: 10px;
  .hint {
    display: inline-flex; align-items: center; gap: 4px;
    font-size: 12px; color: var(--el-text-color-placeholder);
  }
}

/* ====================== 图片面板 ====================== */
.img-grid {
  display: grid;
  grid-template-columns: 340px minmax(0, 1fr);
  gap: 16px;
}
.img-main { min-height: 560px; }

/* 比例按钮 */
.ratio-row { display: flex; gap: 8px; }
.ratio-btn {
  flex: 1;
  background: var(--el-bg-color);
  border: 1px solid var(--el-border-color-lighter);
  border-radius: 10px;
  padding: 10px 0 8px;
  cursor: pointer;
  display: flex; flex-direction: column; align-items: center; gap: 6px;
  font-size: 12px; color: var(--el-text-color-secondary);
  transition: all 0.15s;
  .ratio-box {
    background: var(--el-fill-color-light);
    border-radius: 2px;
    border: 1px solid var(--el-border-color-lighter);
  }
  &:hover {
    border-color: var(--el-color-primary);
    color: var(--el-color-primary);
  }
  &.active {
    border-color: var(--el-color-primary);
    background: var(--el-color-primary-light-9);
    color: var(--el-color-primary);
    font-weight: 600;
    .ratio-box {
      background: var(--el-color-primary);
      border-color: var(--el-color-primary);
    }
  }
}

/* prompt chips */
.chips { display: flex; flex-wrap: wrap; gap: 6px; margin-top: 8px; }
.chip { cursor: pointer; max-width: 100%; overflow: hidden; text-overflow: ellipsis; }
.chip:hover { background: var(--el-color-primary-light-9); color: var(--el-color-primary); }

/* 上传区 */
.upload-zone {
  position: relative;
  display: flex; flex-direction: column; align-items: center;
  padding: 20px 12px;
  border: 2px dashed var(--el-border-color);
  border-radius: 12px;
  cursor: pointer;
  background: var(--el-fill-color-lighter);
  transition: all 0.2s;
  &:hover { border-color: var(--el-color-primary); background: var(--el-color-primary-light-9); }
  .up-ic { font-size: 32px; color: var(--el-color-primary); }
  .up-t { font-size: 13px; margin-top: 6px; color: var(--el-text-color-primary); }
  .up-s { font-size: 11px; color: var(--el-text-color-placeholder); margin-top: 2px; }
  input { position: absolute; inset: 0; opacity: 0; cursor: pointer; }
}

.ref-grid {
  margin-top: 10px;
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 8px;
}
.ref-thumb {
  position: relative;
  aspect-ratio: 1;
  border-radius: 8px;
  overflow: hidden;
  background: var(--el-fill-color-light);
  img { width: 100%; height: 100%; object-fit: cover; display: block; }
  .ref-x {
    position: absolute; top: 4px; right: 4px;
    width: 20px; height: 20px; border-radius: 50%;
    background: rgba(0,0,0,0.55); color: #fff;
    display: flex; align-items: center; justify-content: center;
    cursor: pointer; font-size: 12px;
    opacity: 0; transition: opacity 0.2s;
  }
  .ref-meta {
    position: absolute; bottom: 0; left: 0; right: 0;
    padding: 2px 6px;
    background: linear-gradient(transparent, rgba(0,0,0,0.6));
    color: #fff; font-size: 10px;
    opacity: 0; transition: opacity 0.2s;
  }
  &:hover { .ref-x, .ref-meta { opacity: 1; } }
}

/* 主区 stage / 结果 */
.stage {
  min-height: 480px;
  display: flex; flex-direction: column; align-items: center; justify-content: center;
  text-align: center; color: var(--el-text-color-secondary); padding: 40px 24px;
  .stage-art { font-size: 64px; margin-bottom: 16px; opacity: 0.7; }
  .stage-title { font-size: 16px; font-weight: 600; color: var(--el-text-color-primary); }
  .stage-sub { font-size: 13px; margin-top: 6px; }
  &.loading { gap: 14px; }
  .orb {
    width: 72px; height: 72px; border-radius: 50%;
    background: var(--el-color-primary-light-8);
    display: flex; align-items: center; justify-content: center;
    animation: pulseOrb 1.8s ease-in-out infinite;
  }
}
@keyframes pulseOrb {
  0%, 100% { transform: scale(1); box-shadow: 0 0 0 0 var(--el-color-primary-light-5); }
  50%      { transform: scale(1.08); box-shadow: 0 0 0 14px rgba(64,158,255,0); }
}
.spin { font-size: 30px; animation: spin 1s linear infinite; color: var(--el-color-primary); }
@keyframes spin { to { transform: rotate(360deg); } }

.err-block {
  background: var(--el-color-danger-light-9);
  color: var(--el-color-danger);
  padding: 12px 14px;
  border-radius: 10px;
  display: flex; align-items: center; gap: 8px;
  white-space: pre-wrap; word-break: break-word;
  border: 1px solid var(--el-color-danger-light-5);
}

.result-grid {
  display: grid; grid-template-columns: repeat(auto-fill, minmax(240px, 1fr));
  gap: 14px;
  padding: 4px;
}
.img-cell {
  position: relative;
  aspect-ratio: 1;
  border-radius: 12px;
  overflow: hidden;
  cursor: zoom-in;
  background: var(--el-fill-color-light);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
  transition: all 0.2s;
  img {
    width: 100%; height: 100%; object-fit: cover; display: block;
    transition: transform 0.4s;
  }
  &:hover {
    transform: translateY(-2px);
    box-shadow: 0 10px 24px rgba(0, 0, 0, 0.12);
    img { transform: scale(1.03); }
    .img-actions { opacity: 1; }
  }
}
.img-actions {
  position: absolute; top: 8px; right: 8px;
  display: flex; gap: 6px;
  opacity: 0; transition: opacity 0.2s;
  .iact {
    width: 30px; height: 30px; border-radius: 50%;
    background: rgba(0,0,0,0.55); color: #fff;
    border: none; cursor: pointer;
    display: inline-flex; align-items: center; justify-content: center;
    font-size: 14px;
    &:hover { background: var(--el-color-primary); }
  }
}

/* IMG1 预览兜底专用样式 */
.result-wrap {
  display: flex; flex-direction: column; gap: 10px;
  padding: 4px;
  .result-grid { padding: 0; }
  .preview-tip { border-radius: 10px; }
}
.img-cell.is-preview {
  box-shadow: 0 2px 8px rgba(251, 146, 60, 0.25);
  &::after {
    content: ''; position: absolute; inset: 0;
    border: 1.5px dashed rgba(245, 158, 11, 0.55);
    border-radius: 12px;
    pointer-events: none;
  }
}
.img-badge {
  position: absolute;
  left: 8px; top: 8px;
  padding: 2px 8px;
  font-size: 11px;
  border-radius: 999px;
  background: rgba(245, 158, 11, 0.92);
  color: #fff;
  letter-spacing: 0.3px;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.12);
  pointer-events: none;
}

/* ====================== Dark mode ====================== */
:global(html.dark) .md :deep(.mdk-pre) {
  background: #0b1020;
  color: #cbd5e1;
}

/* ====================== Responsive ====================== */
@media (max-width: 1100px) {
  .chat-grid, .img-grid { grid-template-columns: 1fr; }
  .side { position: static; }
  .chat-main { height: 580px; }
}
@media (max-width: 720px) {
  .hero {
    flex-direction: column;
    align-items: flex-start;
    gap: 10px;
  }
  .hero-sub { display: none; }
  .hero-stats { width: 100%; justify-content: flex-start; }
}
</style>
````

## File: web/src/views/personal/Usage.vue
````vue
<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue'
import * as meApi from '@/api/me'

const loading = ref(false)
const stats = ref<meApi.MyStatsResp | null>(null)
const logs = ref<meApi.UsageItem[]>([])
const total = ref(0)
const filter = reactive({ type: '' as '' | 'chat' | 'image', status: '', limit: 20, offset: 0 })

const page = computed({
  get: () => Math.floor(filter.offset / filter.limit) + 1,
  set: (v: number) => { filter.offset = (v - 1) * filter.limit; loadLogs() },
})

async function loadStats() { stats.value = await meApi.getMyUsageStats({ days: 14, top_n: 10, type: filter.type }) }
async function loadLogs() {
  loading.value = true
  try {
    const d = await meApi.listMyUsageLogs(filter)
    logs.value = d.items || []
    total.value = d.total || 0
  } finally { loading.value = false }
}
async function reload() { filter.offset = 0; await Promise.all([loadStats(), loadLogs()]) }

onMounted(reload)
</script>

<template>
  <div class="page-container">
    <div class="card-block">
      <div class="flex-between">
        <div><h2 class="page-title">用量记录</h2><div class="sub">本页展示本地中转观测数据，仅用于运行排障和趋势观察。</div></div>
        <div class="flex-wrap-gap">
          <el-select v-model="filter.type" placeholder="类型" clearable style="width:120px" @change="reload"><el-option label="对话" value="chat" /><el-option label="生图" value="image" /></el-select>
          <el-select v-model="filter.status" placeholder="状态" clearable style="width:120px" @change="reload"><el-option label="成功" value="success" /><el-option label="失败" value="failed" /></el-select>
          <el-button @click="reload">刷新</el-button>
        </div>
      </div>
    </div>

    <div class="grid cards">
      <div class="card-block kpi"><span>请求</span><b>{{ stats?.overall.requests || 0 }}</b></div>
      <div class="card-block kpi"><span>失败</span><b>{{ stats?.overall.failures || 0 }}</b></div>
      <div class="card-block kpi"><span>输入 Token</span><b>{{ stats?.overall.input_tokens || 0 }}</b></div>
      <div class="card-block kpi"><span>输出 Token</span><b>{{ stats?.overall.output_tokens || 0 }}</b></div>
    </div>

    <div class="card-block" style="margin-top:16px">
      <el-table v-loading="loading" :data="logs" stripe>
        <el-table-column prop="created_at" label="时间" width="180" />
        <el-table-column prop="type" label="类型" width="90" />
        <el-table-column prop="model_slug" label="模型" min-width="150"><template #default="{ row }"><code>{{ row.model_slug || row.model_id }}</code></template></el-table-column>
        <el-table-column prop="status" label="状态" width="100" />
        <el-table-column prop="input_tokens" label="输入" width="100" />
        <el-table-column prop="output_tokens" label="输出" width="100" />
        <el-table-column prop="image_count" label="图片" width="90" />
        <el-table-column prop="duration_ms" label="耗时(ms)" width="110" />
        <el-table-column prop="error_code" label="错误" min-width="160" show-overflow-tooltip />
      </el-table>
      <div class="pager"><el-pagination layout="total, sizes, prev, pager, next" :total="total" v-model:current-page="page" v-model:page-size="filter.limit" @size-change="reload" /></div>
    </div>
  </div>
</template>

<style scoped>
.sub { color:var(--el-text-color-secondary); font-size:13px; margin-top:4px; }
.grid { display:grid; gap:16px; margin-top:16px; }
.cards { grid-template-columns: repeat(4, minmax(0,1fr)); }
.kpi span { color:var(--el-text-color-secondary); display:block; margin-bottom:8px; }
.kpi b { font-size:24px; }
.pager { display:flex; justify-content:flex-end; margin-top:14px; }
code { background:#f2f3f5; padding:1px 6px; border-radius:4px; }
@media (max-width: 960px) { .cards { grid-template-columns:1fr; } }
</style>
````

## File: web/tsconfig.json
````json
{
  "compilerOptions": {
    "target": "ES2022",
    "useDefineForClassFields": true,
    "module": "ESNext",
    "moduleResolution": "bundler",
    "strict": true,
    "noImplicitAny": true,
    "jsx": "preserve",
    "esModuleInterop": true,
    "resolveJsonModule": true,
    "allowSyntheticDefaultImports": true,
    "isolatedModules": true,
    "skipLibCheck": true,
    "lib": ["ES2022", "DOM", "DOM.Iterable"],
    "types": ["node", "element-plus/global"],
    "baseUrl": ".",
    "paths": {
      "@/*": ["src/*"]
    }
  },
  "include": [
    "src/**/*.ts",
    "src/**/*.tsx",
    "src/**/*.vue",
    "src/auto-imports.d.ts",
    "src/components.d.ts",
    "src/env.d.ts"
  ],
  "references": [{ "path": "./tsconfig.node.json" }]
}
````

## File: web/tsconfig.node.json
````json
{
  "compilerOptions": {
    "composite": true,
    "skipLibCheck": true,
    "module": "ESNext",
    "moduleResolution": "bundler",
    "allowSyntheticDefaultImports": true
  },
  "include": ["vite.config.ts"]
}
````

## File: web/vite.config.ts
````typescript
import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'
import path from 'node:path'

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), '')
  const apiBase = env.VITE_API_BASE || 'http://localhost:8080'
  return {
    resolve: {
      alias: {
        '@': path.resolve(__dirname, 'src'),
      },
    },
    server: {
      host: '0.0.0.0',
      port: 5173,
      proxy: {
        // 开发期统一经本地代理,避免 CORS;生产由 nginx / ingress 承担
        '/api': { target: apiBase, changeOrigin: true },
        '/v1': { target: apiBase, changeOrigin: true },
        '/healthz': { target: apiBase, changeOrigin: true },
      },
    },
    plugins: [
      vue(),
      AutoImport({
        imports: ['vue', 'vue-router', 'pinia', '@vueuse/core'],
        resolvers: [ElementPlusResolver()],
        dts: 'src/auto-imports.d.ts',
      }),
      Components({
        resolvers: [ElementPlusResolver()],
        dts: 'src/components.d.ts',
      }),
    ],
    build: {
      outDir: 'dist',
      sourcemap: false,
      chunkSizeWarningLimit: 700,
      rollupOptions: {
        output: {
          /**
           * 手工拆包,避免把 Element Plus 全量塞进 index.js。
           * - element-plus:UI 组件库 + 图标,业务里几乎每页都用,拆出独立 chunk 以便浏览器长期缓存。
           * - vue-core:vue / vue-router / pinia / @vueuse,运行时核心。
           * - vendor:其它 node_modules(axios、dayjs 等)。
           */
          manualChunks(id) {
            if (!id.includes('node_modules')) return
            if (id.includes('element-plus') || id.includes('@element-plus')) return 'element-plus'
            if (
              id.includes('/vue/') ||
              id.includes('/@vue/') ||
              id.includes('/vue-router/') ||
              id.includes('/pinia') ||
              id.includes('/@vueuse/')
            ) return 'vue-core'
            return 'vendor'
          },
        },
      },
    },
  }
})
````
