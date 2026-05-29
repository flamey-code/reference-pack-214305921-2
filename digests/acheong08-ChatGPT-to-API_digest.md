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
.env.example
.github/workflows/build_docker.yml
.github/workflows/build.yml
.github/workflows/release.yml
.gitignore
auth.go
cases/sider/Sider-Config.md
cases/sider/Sider-example.png
conversion/requests/chatgpt/convert.go
conversion/response/chatgpt/convert.go
docker-compose.yml
Dockerfile
docs/admin.md
docs/Docker_CN.md
docs/GUIDE_CN.md
docs/TOKEN_CN.md
go.mod
handlers.go
internal/bard/lib.go
internal/bard/request.go
internal/bard/utilities.go
internal/chatgpt/request.go
internal/tokens/tokens.go
main.go
middleware.go
README_CN.md
README_JA.md
README_ZH.md
README.md
tools/authenticator/auth/auth.go
tools/authenticator/go.mod
tools/authenticator/main.go
tools/authenticator/README.md
tools/authenticator/remove_duplicates.py
tools/authenticator/tojson.sh
tools/plugin_check/check_access.go
tools/proxy_check/proxy_check.go
</directory_structure>

<files>
This section contains the contents of the repository's files.

<file path=".env.example">
ADMIN_PASSWORD=
OPENAI_EMAIL=
OPENAI_PASSWORD=
</file>

<file path=".github/workflows/build_docker.yml">
name: build_docker

on:
  release:
    types: [created] # 表示在创建新的 Release 时触发

jobs:
  build_docker:
    name: Build docker
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v3

      - run: |
          echo "本次构建的版本为：${GITHUB_REF_NAME} (但是这个变量目前上下文中无法获取到)"
          echo 本次构建的版本为：${{ github.ref_name }}
          env

      - name: Set up QEMU
        uses: docker/setup-qemu-action@v2
      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v2
      - name: Login to DockerHub
        uses: docker/login-action@v2
        with:
          username: ${{ secrets.DOCKERHUB_USERNAME }}
          password: ${{ secrets.DOCKERHUB_TOKEN }}
      - name: Build and push
        id: docker_build
        uses: docker/build-push-action@v4
        with:
          context: .
          push: true
          labels: ${{ steps.meta.outputs.labels }}
          platforms: linux/amd64,linux/arm64
          tags: |
            ${{ secrets.DOCKERHUB_USERNAME }}/chatgpt-to-api:${{ github.ref_name }}
            ${{ secrets.DOCKERHUB_USERNAME }}/chatgpt-to-api:latest
</file>

<file path=".github/workflows/build.yml">
name: Go Build

on:
  push:
    branches: [ master ]
  pull_request:
    branches: [ master ]

jobs:
  build:
    name: Build
    runs-on: ubuntu-latest
    strategy:
      fail-fast: false
      matrix:
        include:
          - goos: windows
            goarch: 386
          - goos: windows
            goarch: amd64
          - goos: windows
            goarch: arm64
          - goos: linux
            goarch: 386
          - goos: linux
            goarch: amd64
          - goos: linux
            goarch: arm64
    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Setup Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.20'

      - name: Build
        run: |
          echo "Building for ${{ matrix.goos }} ${{ matrix.goarch }}"
          suffix=""
          if [ "${{ matrix.goos }}" == "windows" ]; then
            suffix=".exe"
          fi
          GOOS=${{ matrix.goos }} GOARCH=${{ matrix.goarch }} go build -o ./build/freechatgpt-${{ matrix.goos }}-${{ matrix.goarch }}$suffix

      - name: Upload artifacts
        uses: actions/upload-artifact@v3
        with:
          name: freechatgpt-${{ matrix.goos }}-${{ matrix.goarch }}
          path: ./build/freechatgpt-${{ matrix.goos }}-${{ matrix.goarch }}*
</file>

<file path=".github/workflows/release.yml">
# This workflow will build a golang project
# For more information see: https://docs.github.com/en/actions/automating-builds-and-tests/building-and-testing-go

name: Release Workflow
on:
  release:
    types:
      - created
permissions:
  contents: write

jobs:

  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3

    - name: Set up Go
      uses: actions/setup-go@v3
      with:
        go-version: 1.19

    - name: Build
      run: go build -o bin/ .
    
    - name: recursively list files
      run: ls -R

    - name: Get existing release body
      id: get_release_body
      run: |
        echo "::set-output name=body::$(curl -s -H 'Authorization: Bearer ${{ secrets.GITHUB_TOKEN }}' https://api.github.com/repos/${{ github.repository }}/releases/tags/${{ github.ref_path }} | jq -r '.body')"

    - name: Upload release artifact
      uses: svenstaro/upload-release-action@v2
      with:
        file: bin/*
        file_glob: true
        tag: ${{ github.ref }}
        body: |
          ${{ steps.get_release_body.outputs.body }}
        repo_token: ${{ secrets.GITHUB_TOKEN }}
</file>

<file path="auth.go">
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"freechatgpt/internal/tokens"

	"github.com/acheong08/OpenAIAuth/auth"
)

var accounts []Account

type Account struct {
	Email    string `json:"username"`
	Password string `json:"password"`
}

// Read accounts.txt and create a list of accounts
func readAccounts() {
	accounts = []Account{}
	// Read accounts.txt and create a list of accounts
	if _, err := os.Stat("accounts.txt"); err == nil {
		// Each line is a proxy, put in proxies array
		file, _ := os.Open("accounts.txt")
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			// Split by :
			line := strings.Split(scanner.Text(), ":")
			// Create an account
			account := Account{
				Email:    line[0],
				Password: line[1],
			}
			// Append to accounts
			accounts = append(accounts, account)
		}
	}
}
func scheduleTokenPUID() {
	// Check if access_tokens.json exists
	if stat, err := os.Stat("access_tokens.json"); os.IsNotExist(err) {
		// Create the file
		file, err := os.Create("access_tokens.json")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		updateToken()
	} else {
		nowTime := time.Now()
		usedTime := nowTime.Sub(stat.ModTime())
		// update access token 7 days after last modify token file
		toExpire := 6.048e14 - usedTime
		if toExpire > 0 {
			file, err := os.Open("access_tokens.json")
			if err != nil {
				panic(err)
			}
			defer file.Close()
			decoder := json.NewDecoder(file)
			var token_list []tokens.Secret
			err = decoder.Decode(&token_list)
			if err != nil {
				updateToken()
				return
			}
			if len(token_list) == 0 {
				updateToken()
			} else {
				ACCESS_TOKENS = tokens.NewAccessToken(token_list, false)
				time.AfterFunc(toExpire, updateToken)
			}
		} else {
			updateToken()
		}
	}
}

func updateToken() {
	token_list := []tokens.Secret{}
	// Loop through each account
	for _, account := range accounts {
		if os.Getenv("CF_PROXY") != "" {
			// exec warp-cli disconnect and connect
			exec.Command("warp-cli", "disconnect").Run()
			exec.Command("warp-cli", "connect").Run()
			time.Sleep(5 * time.Second)
		}
		println("Updating access token for " + account.Email)
		var proxy_url string
		if len(proxies) == 0 {
			proxy_url = ""
		} else {
			proxy_url = proxies[0]
			// Push used proxy to the back of the list
			proxies = append(proxies[1:], proxies[0])
		}
		authenticator := auth.NewAuthenticator(account.Email, account.Password, proxy_url)
		err := authenticator.Begin()
		if err != nil {
			// println("Error: " + err.Details)
			println("Location: " + err.Location)
			println("Status code: " + fmt.Sprint(err.StatusCode))
			println("Details: " + err.Details)
			println("Embedded error: " + err.Error.Error())
			return
		}
		access_token := authenticator.GetAccessToken()
		puid, _ := authenticator.GetPUID()
		token_list = append(token_list, tokens.Secret{access_token, puid})
		println("Success!")
		// Write authenticated account to authenticated_accounts.txt
		f, go_err := os.OpenFile("authenticated_accounts.txt", os.O_APPEND|os.O_WRONLY, 0600)
		if go_err != nil {
			continue
		}
		defer f.Close()
		if _, go_err = f.WriteString(account.Email + ":" + account.Password + "\n"); go_err != nil {
			continue
		}
		// Remove accounts.txt
		os.Remove("accounts.txt")
		// Create accounts.txt
		f, go_err = os.Create("accounts.txt")
		if go_err != nil {
			continue
		}
		defer f.Close()
		// Remove account from accounts
		accounts = accounts[1:]
		// Write unauthenticated accounts to accounts.txt
		for _, acc := range accounts {
			// Check if account is authenticated
			if acc.Email == account.Email {
				continue
			}
			if _, go_err = f.WriteString(acc.Email + ":" + acc.Password + "\n"); go_err != nil {
				continue
			}
		}
	}
	// Append access token to access_tokens.json
	ACCESS_TOKENS = tokens.NewAccessToken(token_list, true)
	time.AfterFunc(6.048e14, updateToken)
}
</file>

<file path="cases/sider/Sider-Config.md">
# 这是一个指导你如何在Sider程序中应用本程序的API
Sider ，一个浏览器插件

https://chrome.google.com/webstore/detail/difoiogjjojoaoomphldepapgpbgkhkb

---
# 设置

打开插件自己的设置页，大概是 chrome-extension://difoiogjjojoaoomphldepapgpbgkhkb/options.html?section=general

找到通用设置-如何访问 ChatGPT 并在任何地方使用它-OpenAI KEY 

在API KEY中填写你的API管理密码

模型默认即可，获取不到的（接口不一样）

URL写你的域名/IP+端口，比如 http://127.0.0.1:8080 ，原先完整URL后面的 /v1 什么的不需要填写

然后直接就能用了

示例图

![Sider-example](./Sider-example.png)


# 不足

- Sider的策略是开启一个聊天，然后始终重复编辑这个对话，使用了本项目API后会导致一直会新建对话，不会删除或编辑旧的对话，所以你需要自行清理。
</file>

<file path="conversion/requests/chatgpt/convert.go">
package chatgpt

import (
	"fmt"
	chatgpt_types "freechatgpt/typings/chatgpt"
	official_types "freechatgpt/typings/official"
	"strings"

	arkose "github.com/acheong08/funcaptcha"
)

func ConvertAPIRequest(api_request official_types.APIRequest, puid string, proxy string) chatgpt_types.ChatGPTRequest {
	chatgpt_request := chatgpt_types.NewChatGPTRequest()
	if strings.HasPrefix(api_request.Model, "gpt-3.5") {
		chatgpt_request.Model = "text-davinci-002-render-sha"
	}
	if strings.HasPrefix(api_request.Model, "gpt-4") {
		token, _, err := arkose.GetOpenAIToken(puid, proxy)
		if err == nil {
			chatgpt_request.ArkoseToken = token
		} else {
			fmt.Println("Error getting Arkose token: ", err)
		}
		chatgpt_request.Model = api_request.Model
		// Cover some models like gpt-4-32k
		if len(api_request.Model) >= 7 && api_request.Model[6] >= 48 && api_request.Model[6] <= 57 {
			chatgpt_request.Model = "gpt-4"
		}
	}
	if api_request.PluginIDs != nil {
		chatgpt_request.PluginIDs = api_request.PluginIDs
		chatgpt_request.Model = "gpt-4-plugins"
	}
	for _, api_message := range api_request.Messages {
		if api_message.Role == "system" {
			api_message.Role = "critic"
		}
		chatgpt_request.AddMessage(api_message.Role, api_message.Content)
	}
	return chatgpt_request
}
</file>

<file path="conversion/response/chatgpt/convert.go">
package chatgpt

import (
	"freechatgpt/typings"
	chatgpt_types "freechatgpt/typings/chatgpt"
	official_types "freechatgpt/typings/official"
	"strings"
)

func ConvertToString(chatgpt_response *chatgpt_types.ChatGPTResponse, previous_text *typings.StringStruct, role bool) string {
	translated_response := official_types.NewChatCompletionChunk(strings.ReplaceAll(chatgpt_response.Message.Content.Parts[0], *&previous_text.Text, ""))
	if role {
		translated_response.Choices[0].Delta.Role = chatgpt_response.Message.Author.Role
	}
	previous_text.Text = chatgpt_response.Message.Content.Parts[0]
	return "data: " + translated_response.String() + "\n\n"

}
</file>

<file path="docker-compose.yml">
version: '3'

services:
  app:
    image: acheong08/chatgpt-to-api # 总是使用latest,更新时重新pull该tag镜像即可
    container_name: chatgpttoapi
    restart: unless-stopped
    ports:
      - '8080:8080'
    environment:
      SERVER_HOST: 0.0.0.0
      SERVER_PORT: 8080
      ADMIN_PASSWORD: TotallySecurePassword
      # If the parameter API_REVERSE_PROXY is empty, the default request URL is https://chat.openai.com/backend-api/conversation, and the PUID is <NOT> equired.
      PUID: xxx
</file>

<file path="Dockerfile">
# Use the official Golang image as the builder
FROM golang:1.20.3-alpine as builder

# Enable CGO to use C libraries (set to 0 to disable it)
# We set it to 0 to build a fully static binary for our final image
ENV CGO_ENABLED=0

# Set the working directory
WORKDIR /app

# Copy the Go Modules manifests (go.mod and go.sum files)
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application and output the binary to /app/ChatGPT-Proxy-V4
RUN go build -o /app/ChatGPT-To-API .

# Use a scratch image as the final distroless image
FROM scratch

# Set the working directory
WORKDIR /app

# Copy the built Go binary from the builder stage
COPY --from=builder /app/ChatGPT-To-API /app/ChatGPT-To-API

# Expose the port where the application is running
EXPOSE 8080

# Start the application
CMD [ "./ChatGPT-To-API" ]
</file>

<file path="docs/admin.md">
# API Documentation:
## openaiHandler:

This API endpoint receives a POST request with a JSON body that contains OpenAI credentials. The API updates the environment variables "OPENAI_EMAIL" and "OPENAI_PASSWORD" with the provided credentials.

HTTP method: PATCH

Endpoint: /openai

Request body:

```json
{
    "OpenAI_Email": "string",
    "OpenAI_Password": "string"
}
```

Response status codes:
- 200 OK: The OpenAI credentials were successfully updated.
- 400 Bad Request: The JSON in the request body is invalid.

## passwordHandler:

This API endpoint receives a POST request with a JSON body that contains a new password. The API updates the global variable "ADMIN_PASSWORD" and the environment variable "ADMIN_PASSWORD" with the new password.

HTTP method: PATCH

Endpoint: /password

Request body:

```json
{
    "password": "string"
}
```

Response status codes:
- 200 OK: The password was successfully updated.
- 400 Bad Request: The password is missing or not provided in the request body.

## puidHandler:

This API endpoint receives a POST request with a JSON body that contains a new PUID (Personal User ID). The API updates the environment variable "PUID" with the new PUID.

HTTP method: PATCH

Endpoint: /puid

Request body:

```json
{
    "puid": "string"
}
```

Response status codes:
- 200 OK: The PUID was successfully updated.
- 400 Bad Request: The PUID is missing or not provided in the request body.

## tokensHandler:

This API endpoint receives a POST request with a JSON body that contains an array of request tokens. The API updates the value of the global variable "ACCESS_TOKENS" with a new access token generated from the request tokens provided in the request body.

HTTP method: PATCH

Endpoint: /tokens

Request body:

```json
[
    "string", "..."
]
```

Response status codes:
- 200 OK: The ACCESS_TOKENS variable was successfully updated.
- 400 Bad Request: The request tokens are missing or not provided in the request body.
</file>

<file path="docs/Docker_CN.md">
# 使用阿里源实现Docker安装

移除旧的

```yum remove -y docker docker-common docker-selinux docker-engine```

安装依赖

```yum install -y yum-utils device-mapper-persistent-data lvm2```

配置Docker安装源（阿里）

```yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo```

检查可用的Docker-CE版本

```yum list docker-ce --showduplicates | sort -r```

安装Docker-CE

```yum -y install docker-ce```

运行Docker（默认不运行）

```systemctl start docker```

配置开机启动Docker

```systemctl enable docker```

# 使用官方二进制包安装Docker-Compase

下载 Docker-Compose 的二进制文件
```sudo curl -L "https://github.com/docker/compose/releases/download/v2.18.1/docker-compose-linux-x86_64" -o /usr/local/bin/docker-compose```

添加可执行权限

```sudo chmod +x /usr/local/bin/docker-compose```

验证 Docker-Compose 是否安装成功

```docker-compose --version```

启动容器

```docker-compose up -d```

关闭容器

```docker-compose down```

查看容器（如果启动了这里没有说明启动失败）

```docker ps```

# ChatGPT-TO-API的Docker-Compase文件

```
    ports:
      - '31480:31480'
    environment:
      SERVER_HOST: 0.0.0.0
      SERVER_PORT: 31480
      ADMIN_PASSWORD: TotallySecurePassword
      # Reverse Proxy - Available on accessToken
      #API_REVERSE_PROXY: https://bypass.churchless.tech/api/conversation
      #API_REVERSE_PROXY: https://ai.fakeopen.com/api/conversation
      PUID: user-7J4tdvHySlcilVgjFIrAtK1k

```

- 这里的ports，左边是外部端口，用于外部访问。右边的Docker端口，需要匹配下面程序设置的监听Port。
- 如果参数`API_REVERSE_PROXY`为空，则默认的请求URL为`https://chat.openai.com/backend-api/conversation`，并且需要提供PUID。PUID的获取参考 [README_CN.md](../README_CN.md)
- 这个密码需要自定义，我们构建请求的时候需要它来鉴权。默认是```TotallySecurePassword```
</file>

<file path="docs/GUIDE_CN.md">
# 中文指导手册

本中文手册由 [@BlueSkyXN](https://github.com/BlueSkyXN) 编写


[中文文档（Chinese Docs）](../README_CN.md)  
 [英文文档（English Docs）](../README.md)

# 基本配置

有关docker的指导请阅读 [DOCKER中文手册](Docker_CN.md)

有关Token的指导请阅读 [TOKEN中文手册](TOKEN_CN.md)

## Docker-Compose配置

```
version: '3'

services:
  app:
    image: acheong08/chatgpt-to-api
    container_name: chatgpttoapi
    restart: unless-stopped
    ports:
      - '10080:10080'
    environment:
      SERVER_HOST: 0.0.0.0
      SERVER_PORT: 10080
      ADMIN_PASSWORD: TotallySecurePassword
      API_REVERSE_PROXY: https://ai.fakeopen.com/api/conversation
      PUID: user-X

```
- ports 左边是外部端口，右边是内部端口，内部端口要和下面环境变量的Server port一致。
- Server host/port：监听配置，默认0000监听某一端口。
- ADMIN_PASSWORD：管理员密码，HTTP请求时候需要验证。
- API_REVERSE_PROXY:接口的反向代理，具体介绍请看下文的后端代理介绍部分。
- PUID: user-X，请看[中文文档（Chinese Docs）](../README_CN.md) 的介绍

其他可以不需要设置，包括预设的AccessToken和代理表、HTTP/S5代理。


# 后端代理
目前使用PUID+官网URL的方式不是很可靠，建议使用第三方程序或者网站绕过这个WAF限制。


## 公共代理
温馨提醒，由于OpenAI用的强力CloudFlareWAF，所以7层转发是无效的（不过4层在浏览器还是可以的）

目前根据几个大项目的介绍，我找到了这个介绍页 https://github.com/transitive-bullshit/chatgpt-api#reverse-proxy
最后得知主要是这两个

| Reverse Proxy URL                              | Author       | Rate Limits                    | Last Checked |
|-----------------------------------------------|--------------|--------------------------------|--------------|
| https://ai.fakeopen.com/api/conversation       | @pengzhile   | 5 req / 10 seconds by IP       | 4/18/2023    |
| https://api.pawan.krd/backend-api/conversation | @PawanOsman  | 50 req / 15 seconds (~3 r/s)   | 3/23/2023    |


## 自建方案

我经过测试，发现Pandora的API不行，原因可能是发起对话后的返回值会一次性返回一堆信息导致提取失败。不过我亲测GO-ChatGPT-API是可以的。

GO-ChatGPT-API项目 https://github.com/linweiyuan/go-chatgpt-api

我是注释掉 ##- GO_CHATGPT_API_PROXY= 的环境变量、换个外部端口后用Docker-Compose启动即可。然后不需要对这个代理接口做其他操作，包括登录。

搭建好之后最好测试下基本调用能不能用，下面是一个示例，你需要根据实际情况修改。


```
curl http://127.0.0.1:8080/chatgpt/backend-api/conversation \
  -H "Content-Type: application/json" \
  -d '{
     "model": "gpt-3.5-turbo",
     "messages": [{"role": "user", "content": "Say this is a test!"}],
     "temperature": 0.7
   }'

```

如果得到缺少认证的提示比如 ```{"errorMessage":"Missing accessToken."}``` 就说明已经正常跑了

# 用例

## 基本提问
```
curl http://127.0.0.1:10080/v1/chat/completions \
  -d '{
     "model": "text-davinci-002-render-sha",
     "messages": [{"role": "user", "content": "你是什么模型，是GPT3.5吗"}]
   }'
```

参考回复如下

```
{"id":"chatcmpl-QXlha2FBbmROaXhpZUFyZUF3XXXXXX","object":"chat.completion","created":0,"model":"gpt-3.5-turbo-0301","usage":{"prompt_tokens":0,"completion_tokens":0,"total_tokens":0},"choices":[{"index":0,"message":{"role":"assistant","content":"是的，我是一个基于GPT-3.5架构的语言模型，被称为ChatGPT。我可以回答各种问题，提供信息和进行对话。尽管我会尽力提供准确和有用的回答，但请记住，我并不是完美的，有时候可能会出现错误或者误导性的答案。"},"finish_reason":null}]}
```

请注意无论什么模型提问都只会显示为模型是GPT3.5T-0301。你在网页版看不到消息记录（可能是删除了），Chat不支持并发提问，你需要token轮询。

## 提交Token
通过文件提交

```
curl -X PATCH \
     -H "Content-Type: application/json" \
     -H "Authorization: TotallySecurePassword" \
     -d "@/root/access_tokens.json" \
     http://127.0.0.1:10080/admin/tokens

```

直接提交

```
curl -X PATCH \
  -H "Content-Type: application/json" \
  -H "Authorization: TotallySecurePassword" \
  -d '["eyJhbXXX"]' \
  http://127.0.0.1:10080/admin/tokens
```

要清理Token直接停用删除Docker容器后重新构建运行容器即可
</file>

<file path="docs/TOKEN_CN.md">
# 获取Token
---
# 参考Pandora项目的作者指导

https://github.com/pengzhile/pandora

获取Token的技术原理 https://zhile.io/2023/05/19/how-to-get-chatgpt-access-token-via-pkce.html

## 第三方接口获取Token
http://ai.fakeopen.com/auth 

你需要在这个新的网站的指导下安装浏览器插件，官方说明的有效期是14天。支持谷歌微软等第三方登录。（我谷歌注册的OpenAI就可以用这个）    

## 官网获取 Token
https://chat.openai.com/api/auth/session

打开后是个JSON，你需要先登录官方的ChatGPT网页版。里面有一个参数就是AccessToken。

# 参考go-chatgpt-api项目的作者指导
https://github.com/linweiyuan/go-chatgpt-api

ChatGPT 登录（返回 accessToken）（目前仅支持 ChatGPT 账号，谷歌或微软账号没有测试）

```POST /chatgpt/login```
</file>

<file path="go.mod">
module freechatgpt

go 1.20

require (
	github.com/acheong08/OpenAIAuth v0.0.0-20230719092354-c8cd4e19491b
	github.com/acheong08/endless v0.0.0-20230615162514-90545c7793fd
	github.com/acheong08/funcaptcha v1.9.3-0.20230803133445-f4d081d60ac7
	github.com/bogdanfinn/fhttp v0.5.23
	github.com/bogdanfinn/tls-client v1.4.0
	github.com/gin-gonic/gin v1.9.1
	github.com/go-resty/resty/v2 v2.7.0
	github.com/google/uuid v1.3.0
	github.com/joho/godotenv v1.5.1
	github.com/tidwall/gjson v1.14.4
	k8s.io/apimachinery v0.27.2
)

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/bogdanfinn/utls v1.5.16 // indirect
	github.com/bytedance/sonic v1.9.1 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.14.1 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.16.6 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nirasan/go-oauth-pkce-code-verifier v0.0.0-20220510032225-4f9f17eaec4c // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/tam7t/hpkp v0.0.0-20160821193359-2b70b4024ed5 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/crypto v0.10.0 // indirect
	golang.org/x/net v0.11.0 // indirect
	golang.org/x/sys v0.9.0 // indirect
	golang.org/x/text v0.10.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
</file>

<file path="handlers.go">
package main

import (
	chatgpt_request_converter "freechatgpt/conversion/requests/chatgpt"
	chatgpt "freechatgpt/internal/chatgpt"
	"freechatgpt/internal/tokens"
	official_types "freechatgpt/typings/official"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func openaiHandler(c *gin.Context) {
	var authorizations struct {
		OpenAI_Email     string `json:"openai_email"`
		OpenAI_Password  string `json:"openai_password"`
		Official_API_Key string `json:"official_api_key"`
	}
	err := c.BindJSON(&authorizations)
	if err != nil {
		c.JSON(400, gin.H{"error": "JSON invalid"})
	}
	if authorizations.OpenAI_Email != "" && authorizations.OpenAI_Password != "" {
		os.Setenv("OPENAI_EMAIL", authorizations.OpenAI_Email)
		os.Setenv("OPENAI_PASSWORD", authorizations.OpenAI_Password)
	}
	if authorizations.Official_API_Key != "" {
		os.Setenv("OFFICIAL_API_KEY", authorizations.Official_API_Key)
	}
	if authorizations.OpenAI_Email == "" && authorizations.OpenAI_Password == "" && authorizations.Official_API_Key == "" {
		c.JSON(400, gin.H{"error": "JSON invalid"})
		return
	}
	c.String(200, "OpenAI credentials updated")
}

func passwordHandler(c *gin.Context) {
	// Get the password from the request (json) and update the password
	type password_struct struct {
		Password string `json:"password"`
	}
	var password password_struct
	err := c.BindJSON(&password)
	if err != nil {
		c.String(400, "password not provided")
		return
	}
	ADMIN_PASSWORD = password.Password
	// Set environment variable
	os.Setenv("ADMIN_PASSWORD", ADMIN_PASSWORD)
	c.String(200, "password updated")
}

func puidHandler(c *gin.Context) {
	// Get the password from the request (json) and update the password
	type puid_struct struct {
		PUID string `json:"puid"`
	}
	var puid puid_struct
	err := c.BindJSON(&puid)
	if err != nil {
		c.String(400, "puid not provided")
		return
	}
	// Set environment variable
	os.Setenv("PUID", puid.PUID)
	c.String(200, "puid updated")
}

func tokensHandler(c *gin.Context) {
	// Get the request_tokens from the request (json) and update the request_tokens
	var request_tokens []tokens.Secret
	err := c.BindJSON(&request_tokens)
	if err != nil {
		c.String(400, "tokens not provided")
		return
	}
	ACCESS_TOKENS = tokens.NewAccessToken(request_tokens, true)
	c.String(200, "tokens updated")
}
func optionsHandler(c *gin.Context) {
	// Set headers for CORS
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST")
	c.Header("Access-Control-Allow-Headers", "*")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func nightmare(c *gin.Context) {
	var original_request official_types.APIRequest
	err := c.BindJSON(&original_request)
	if err != nil {
		c.JSON(400, gin.H{"error": gin.H{
			"message": "Request must be proper JSON",
			"type":    "invalid_request_error",
			"param":   nil,
			"code":    err.Error(),
		}})
		return
	}

	authHeader := c.GetHeader("Authorization")
	token, puid := ACCESS_TOKENS.GetSecret()
	if authHeader != "" {
		customAccessToken := strings.Replace(authHeader, "Bearer ", "", 1)
		// Check if customAccessToken starts with sk-
		if strings.HasPrefix(customAccessToken, "eyJhbGciOiJSUzI1NiI") {
			token = customAccessToken
		}
	}
	var proxy_url string
	if len(proxies) == 0 {
		proxy_url = ""
	} else {
		proxy_url = proxies[0]
		// Push used proxy to the back of the list
		proxies = append(proxies[1:], proxies[0])
	}

	// Convert the chat request to a ChatGPT request
	translated_request := chatgpt_request_converter.ConvertAPIRequest(original_request, puid, proxy_url)

	response, err := chatgpt.POSTconversation(translated_request, token, puid, proxy_url)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "error sending request",
		})
		return
	}
	defer response.Body.Close()
	if chatgpt.Handle_request_error(c, response) {
		return
	}
	var full_response string
	for i := 3; i > 0; i-- {
		var continue_info *chatgpt.ContinueInfo
		var response_part string
		response_part, continue_info = chatgpt.Handler(c, response, token, translated_request, original_request.Stream)
		full_response += response_part
		if continue_info == nil {
			break
		}
		println("Continuing conversation")
		translated_request.Messages = nil
		translated_request.Action = "continue"
		translated_request.ConversationID = continue_info.ConversationID
		translated_request.ParentMessageID = continue_info.ParentID
		response, err = chatgpt.POSTconversation(translated_request, token, puid, proxy_url)
		if err != nil {
			c.JSON(500, gin.H{
				"error": "error sending request",
			})
			return
		}
		defer response.Body.Close()
		if chatgpt.Handle_request_error(c, response) {
			return
		}
	}
	if !original_request.Stream {
		c.JSON(200, official_types.NewChatCompletion(full_response))
	} else {
		c.String(200, "data: [DONE]\n\n")
	}

}

func engines_handler(c *gin.Context) {
	resp, status, err := chatgpt.GETengines()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "error sending request",
		})
		return
	}
	c.JSON(status, resp)
}
</file>

<file path="internal/bard/lib.go">
package bard

// By @mosajjal at https://github.com/mosajjal/bard-cli/blob/main/bard/bard.go
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/tidwall/gjson"
	"k8s.io/apimachinery/pkg/util/rand"

	nhttp "net/http"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	resty "github.com/go-resty/resty/v2"
)

var (
	jar     = tls_client.NewCookieJar()
	options = []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(360),
		tls_client.WithClientProfile(tls_client.Safari_IOS_15_5),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar), // create cookieJar instance and pass it as argument
		// Disable SSL verification
		tls_client.WithInsecureSkipVerify(),
	}
	client, _    = tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	http_proxy   = os.Getenv("http_proxy")
	resty_client = resty.New()
)

var headers map[string]string = map[string]string{
	"Host":          "bard.google.com",
	"X-Same-Domain": "1",
	"User-Agent":    "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.4472.114 Safari/537.36",
	"Content-Type":  "application/x-www-form-urlencoded;charset=UTF-8",
	"Origin":        "https://bard.google.com",
	"Referer":       "https://bard.google.com/",
}

func init() {
	if http_proxy != "" {
		client.SetProxy(http_proxy)
	}
}

const bardURL string = "https://bard.google.com/_/BardChatUi/data/assistant.lamda.BardFrontendService/StreamGenerate"

type Answer struct {
	Content string `json:"content"`
	// FactualityQueries []string `json:"factualityQueries"`
	Choices []string `json:"choices"`
}

// Bard is the main struct for the Bard AI
type Bard struct {
	Cookie              string
	ChoiceID            string
	ConversationID      string
	ResponseID          string
	SNlM0e              string
	LastInteractionTime time.Time
}

// New creates a new Bard AI instance. Cookie is the __Secure-1PSID cookie from Google
func New(cookie string) (*Bard, error) {
	b := &Bard{
		Cookie: cookie,
	}
	err := b.getSNlM0e()
	return b, err
}

func (b *Bard) getSNlM0e() error {
	req, _ := http.NewRequest("GET", "https://bard.google.com/", nil)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	req.AddCookie(&http.Cookie{
		Name:  "__Secure-1PSID",
		Value: b.Cookie,
	})
	// in response text, the value shows. in python:
	r := regexp.MustCompile(`SNlM0e\":\"(.*?)\"`)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	tmpValues := r.FindStringSubmatch(string(body))
	if len(tmpValues) < 2 {
		return fmt.Errorf("failed to find snim0e value. possibly misconfigured cookies?")
	}
	b.SNlM0e = tmpValues[1]
	return nil
}

// Ask generates a Bard AI response and returns it to the user
func (b *Bard) Ask(prompt string) (*Answer, error) {
	b.LastInteractionTime = time.Now()

	// req paramters for the actual request
	reqParams := map[string]string{
		"bl":     "boq_assistant-bard-web-server_20230606.12_p0",
		"_reqid": fmt.Sprintf("%d", rand.IntnRange(100000, 999999)),
		"rt":     "c",
	}

	req := fmt.Sprintf(`[null, "[[\"%s\"], null, [\"%s\", \"%s\", \"%s\"]]"]`,
		//prompt, b.answer.ConversationID, b.answer.ResponseID, b.answer.ChoiceID)
		prompt, b.ConversationID, b.ResponseID, b.ChoiceID)

	reqData := map[string]string{
		"f.req": string(req),
		"at":    b.SNlM0e,
	}
	resty_client.SetHeaders(headers)
	resty_client.SetCookie(&nhttp.Cookie{
		Name:  "__Secure-1PSID",
		Value: b.Cookie,
	})
	resty_client.SetBaseURL(bardURL)
	resty_client.SetTimeout(60 * time.Second)
	resty_client.SetFormData(reqData)
	resty_client.SetQueryParams(reqParams)
	resty_client.SetDoNotParseResponse(true)
	resp, err := resty_client.R().Post("")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		// curl, _ := http2curl.GetCurlCommand(resp.Request.EnableTrace().RawRequest)
		// fmt.Println(curl)
		return nil, fmt.Errorf("status code is not 200: %d", resp.StatusCode())
	}

	// this is the Go version
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(resp.RawResponse.Body)

	respLines := strings.Split(buf.String(), "\n")
	respJSON := respLines[3]

	var fullRes [][]interface{}
	err = json.Unmarshal([]byte(respJSON), &fullRes)
	if err != nil {
		return nil, err
	}

	// get the main answer
	res, ok := fullRes[0][2].(string)
	if !ok {
		return nil, fmt.Errorf("failed to get answer from bard")
	}

	answer := Answer{}

	answer.Content = gjson.Get(res, "0.0").String()
	b.ConversationID = gjson.Get(res, "1.0").String()
	b.ResponseID = gjson.Get(res, "1.1").String()
	choices := gjson.Get(res, "4").Array()
	answer.Choices = make([]string, len(choices))
	for i, choice := range choices {
		answer.Choices[i] = choice.Array()[0].String()
	}
	b.ChoiceID = choices[0].Array()[0].String()

	return &answer, nil
}
</file>

<file path="internal/bard/request.go">
package bard

import "time"

type BardCache struct {
	Bards map[string]*Bard
}

var cache *BardCache

func init() {
	cache = &BardCache{
		Bards: make(map[string]*Bard),
	}
	go func() {
		for {
			GarbageCollectCache(cache)
			time.Sleep(time.Minute)
		}
	}()
}
</file>

<file path="internal/bard/utilities.go">
package bard

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

func HashConversation(conversation []string) string {
	hash := md5.New()
	for _, message := range conversation {
		hash.Write([]byte(message))
	}
	return hex.EncodeToString(hash.Sum(nil))
}

func GarbageCollectCache(cache *BardCache) {
	for k, v := range cache.Bards {
		if time.Since(v.LastInteractionTime) > time.Minute*5 {
			delete(cache.Bards, k)
		}
	}
}

func UpdateBardHash(old_hash, hash string) {
	if _, ok := cache.Bards[old_hash]; ok {
		cache.Bards[hash] = cache.Bards[old_hash]
		delete(cache.Bards, old_hash)
	}
}
</file>

<file path="internal/chatgpt/request.go">
package chatgpt

import (
	"bufio"
	"bytes"
	"encoding/json"
	"freechatgpt/typings"
	chatgpt_types "freechatgpt/typings/chatgpt"
	"io"
	"os"
	"strings"

	arkose "github.com/acheong08/funcaptcha"
	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/gin-gonic/gin"

	chatgpt_response_converter "freechatgpt/conversion/response/chatgpt"

	// chatgpt "freechatgpt/internal/chatgpt"

	official_types "freechatgpt/typings/official"
)

var (
	jar     = tls_client.NewCookieJar()
	options = []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(360),
		tls_client.WithClientProfile(tls_client.Okhttp4Android13),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar), // create cookieJar instance and pass it as argument
		// Disable SSL verification
		tls_client.WithInsecureSkipVerify(),
	}
	client, _         = tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	API_REVERSE_PROXY = os.Getenv("API_REVERSE_PROXY")
)

func init() {
	arkose.SetTLSClient(&client)
}

func POSTconversation(message chatgpt_types.ChatGPTRequest, access_token string, puid string, proxy string) (*http.Response, error) {
	if proxy != "" {
		client.SetProxy(proxy)
	}

	apiUrl := "https://chat.openai.com/backend-api/conversation"
	if API_REVERSE_PROXY != "" {
		apiUrl = API_REVERSE_PROXY
	}

	// JSONify the body and add it to the request
	body_json, err := json.Marshal(message)
	if err != nil {
		return &http.Response{}, err
	}

	request, err := http.NewRequest(http.MethodPost, apiUrl, bytes.NewBuffer(body_json))
	if err != nil {
		return &http.Response{}, err
	}
	// Clear cookies
	if puid != "" {
		request.Header.Set("Cookie", "_puid="+puid+";")
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
	request.Header.Set("Accept", "*/*")
	if access_token != "" {
		request.Header.Set("Authorization", "Bearer "+access_token)
	}
	if err != nil {
		return &http.Response{}, err
	}
	response, err := client.Do(request)
	return response, err
}

// Returns whether an error was handled
func Handle_request_error(c *gin.Context, response *http.Response) bool {
	if response.StatusCode != 200 {
		// Try read response body as JSON
		var error_response map[string]interface{}
		err := json.NewDecoder(response.Body).Decode(&error_response)
		if err != nil {
			// Read response body
			body, _ := io.ReadAll(response.Body)
			c.JSON(500, gin.H{"error": gin.H{
				"message": "Unknown error",
				"type":    "internal_server_error",
				"param":   nil,
				"code":    "500",
				"details": string(body),
			}})
			return true
		}
		c.JSON(response.StatusCode, gin.H{"error": gin.H{
			"message": error_response["detail"],
			"type":    response.Status,
			"param":   nil,
			"code":    "error",
		}})
		return true
	}
	return false
}

type ContinueInfo struct {
	ConversationID string `json:"conversation_id"`
	ParentID       string `json:"parent_id"`
}

func Handler(c *gin.Context, response *http.Response, token string, translated_request chatgpt_types.ChatGPTRequest, stream bool) (string, *ContinueInfo) {
	max_tokens := false

	// Create a bufio.Reader from the response body
	reader := bufio.NewReader(response.Body)

	// Read the response byte by byte until a newline character is encountered
	if stream {
		// Response content type is text/event-stream
		c.Header("Content-Type", "text/event-stream")
	} else {
		// Response content type is application/json
		c.Header("Content-Type", "application/json")
	}
	var finish_reason string
	var previous_text typings.StringStruct
	var original_response chatgpt_types.ChatGPTResponse
	var isRole = true
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", nil
		}
		if len(line) < 6 {
			continue
		}
		// Remove "data: " from the beginning of the line
		line = line[6:]
		// Check if line starts with [DONE]
		if !strings.HasPrefix(line, "[DONE]") {
			// Parse the line as JSON

			err = json.Unmarshal([]byte(line), &original_response)
			if err != nil {
				continue
			}
			if original_response.Error != nil {
				c.JSON(500, gin.H{"error": original_response.Error})
				return "", nil
			}
			if original_response.Message.Author.Role != "assistant" || original_response.Message.Content.Parts == nil {
				continue
			}
			if original_response.Message.Metadata.MessageType != "next" && original_response.Message.Metadata.MessageType != "continue" || original_response.Message.EndTurn != nil {
				continue
			}
			response_string := chatgpt_response_converter.ConvertToString(&original_response, &previous_text, isRole)
			isRole = false
			if stream {
				_, err = c.Writer.WriteString(response_string)
				if err != nil {
					return "", nil
				}
			}
			// Flush the response writer buffer to ensure that the client receives each line as it's written
			c.Writer.Flush()

			if original_response.Message.Metadata.FinishDetails != nil {
				if original_response.Message.Metadata.FinishDetails.Type == "max_tokens" {
					max_tokens = true
				}
				finish_reason = original_response.Message.Metadata.FinishDetails.Type
			}

		} else {
			if stream {
				final_line := official_types.StopChunk(finish_reason)
				c.Writer.WriteString("data: " + final_line.String() + "\n\n")
			}
		}
	}
	if !max_tokens {
		return previous_text.Text, nil
	}
	return previous_text.Text, &ContinueInfo{
		ConversationID: original_response.ConversationID,
		ParentID:       original_response.Message.ID,
	}
}

func GETengines() (interface{}, int, error) {
	url := "https://api.openai.com/v1/models"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+os.Getenv("OFFICIAL_API_KEY"))
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	var result interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result, resp.StatusCode, nil
}
</file>

<file path="internal/tokens/tokens.go">
package tokens

import (
	"encoding/json"
	"os"
	"sync"
)

type Secret struct {
	Token string `json:"token"`
	PUID  string `json:"puid"`
}
type AccessToken struct {
	tokens []Secret
	lock   sync.Mutex
}

func NewAccessToken(tokens []Secret, save bool) AccessToken {
	// Save the tokens to a file
	if _, err := os.Stat("access_tokens.json"); os.IsNotExist(err) {
		// Create the file
		file, err := os.Create("access_tokens.json")
		if err != nil {
			return AccessToken{}
		}
		defer file.Close()
	}
	if save {
		saved := Save(tokens)
		if saved == false {
			return AccessToken{}
		}
	}
	return AccessToken{
		tokens: tokens,
	}
}

func Save(tokens []Secret) bool {
	file, err := os.OpenFile("access_tokens.json", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return false
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(tokens)
	if err != nil {
		return false
	}
	return true
}

func (a *AccessToken) GetSecret() (string, string) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if len(a.tokens) == 0 {
		return "", ""
	}

	secret := a.tokens[0]
	a.tokens = append(a.tokens[1:], secret)
	return secret.Token, secret.PUID
}
</file>

<file path="main.go">
package main

import (
	"bufio"
	"freechatgpt/internal/tokens"
	"os"
	"strings"

	"github.com/acheong08/endless"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var HOST string
var PORT string
var ACCESS_TOKENS tokens.AccessToken
var proxies []string

func checkProxy() {
	// first check for proxies.txt
	proxies = []string{}
	if _, err := os.Stat("proxies.txt"); err == nil {
		// Each line is a proxy, put in proxies array
		file, _ := os.Open("proxies.txt")
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			// Split line by :
			proxy := scanner.Text()
			proxy_parts := strings.Split(proxy, ":")
			if len(proxy_parts) > 1 {
				proxies = append(proxies, proxy)
			} else {
				continue
			}
		}
	}
	// if no proxies, then check env http_proxy
	if len(proxies) == 0 {
		proxy := os.Getenv("http_proxy")
		if proxy != "" {
			proxies = append(proxies, proxy)
		}
	}
}

func init() {
	_ = godotenv.Load(".env")

	HOST = os.Getenv("SERVER_HOST")
	PORT = os.Getenv("SERVER_PORT")
	if PORT == "" {
		PORT = os.Getenv("PORT")
	}
	if HOST == "" {
		HOST = "127.0.0.1"
	}
	if PORT == "" {
		PORT = "8080"
	}
	checkProxy()
	readAccounts()
	scheduleTokenPUID()
}
func main() {
	router := gin.Default()

	router.Use(cors)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	admin_routes := router.Group("/admin")
	admin_routes.Use(adminCheck)

	/// Admin routes
	admin_routes.PATCH("/password", passwordHandler)
	admin_routes.PATCH("/tokens", tokensHandler)
	admin_routes.PATCH("/puid", puidHandler)
	admin_routes.PATCH("/openai", openaiHandler)
	/// Public routes
	router.OPTIONS("/v1/chat/completions", optionsHandler)
	router.POST("/v1/chat/completions", Authorization, nightmare)
	router.GET("/v1/engines", Authorization, engines_handler)
	router.GET("/v1/models", Authorization, engines_handler)
	endless.ListenAndServe(HOST+":"+PORT, router)
}
</file>

<file path="middleware.go">
package main

import (
	"bufio"
	"os"
	"strings"

	gin "github.com/gin-gonic/gin"
)

var ADMIN_PASSWORD string
var API_KEYS map[string]bool

func init() {
	ADMIN_PASSWORD = os.Getenv("ADMIN_PASSWORD")
	if ADMIN_PASSWORD == "" {
		ADMIN_PASSWORD = "TotallySecurePassword"
	}
}

func adminCheck(c *gin.Context) {
	password := c.Request.Header.Get("Authorization")
	if password != ADMIN_PASSWORD {
		c.String(401, "Unauthorized")
		c.Abort()
		return
	}
	c.Next()
}

func cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Next()
}

func Authorization(c *gin.Context) {
	if API_KEYS == nil {
		API_KEYS = make(map[string]bool)
		if _, err := os.Stat("api_keys.txt"); err == nil {
			file, _ := os.Open("api_keys.txt")
			defer file.Close()
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				key := scanner.Text()
				if key != "" {
					API_KEYS["Bearer "+key] = true
				}
			}
		}
	}
	if len(API_KEYS) != 0 && !API_KEYS[c.Request.Header.Get("Authorization")] {
		if c.Request.Header.Get("Authorization") == "" {
			c.JSON(401, gin.H{"error": "No API key provided. Get one at https://discord.gg/9K2BvbXEHT"})
		} else if strings.HasPrefix(c.Request.Header.Get("Authorization"), "Bearer sk-") {
			c.JSON(401, gin.H{"error": "You tried to use the official API key which is not supported."})
		} else if strings.HasPrefix(c.Request.Header.Get("Authorization"), "Bearer eyJhbGciOiJSUzI1NiI") {
			return
		} else {
			c.JSON(401, gin.H{"error": "Invalid API key."})
		}
		c.Abort()
		return
	}
	c.Next()
}
</file>

<file path="README_CN.md">
# ChatGPT-to-API
创建一个模拟API（通过ChatGPT网页版）。使用AccessToken把ChatGPT模拟成OpenAI API，从而在各类应用程序中使用OpenAI的API且不需要为API额外付费，因为模拟成网页版的使用了，和官方API基本互相兼容。

本中文手册由 [@BlueSkyXN](https://github.com/BlueSkyXN) 编写

[英文文档（English Docs）](README.md)

## 认证和各项准备工作

在使用之前，你需要完成一系列准备工作

1. 准备ChatGPT账号，最好的PLUS订阅的，有没有开API不重要
2. 完善的运行环境和网络环境（否则你总是要寻找方法绕过）
3. Access Token和PUID，下面会教你怎么获取
4. 选择一个代理后端或者自行搭建
5. 你可以在 https://github.com/BlueSkyXN/OpenAI-Quick-DEV 项目找到一些常用组件以及一些快速运行的教程或程序。

### 获取PUID

`_puid` cookie.

### 获取Access Token
目前有多种方法和原理，这部分内容可以参考 [TOKEN中文手册](docs/TOKEN_CN.md)

## 安装和运行
  
作者在[英文版介绍](README.md) 通过GO编译来构建二进制程序，但是我猜测这可能需要一个GO编译环境。所以我建议基于作者的Compose配置文件来Docker运行。 

有关docker的指导请阅读 [DOCKER中文手册](docs/Docker_CN.md)

安装好Docker和Docker-Compase后，通过Compase来启动

```docker-compose up -d```

注意，启动之前你需要配置 yml 配置文件，主要是端口和环境变量，各项参数、用法请参考 [中文指导手册](docs/GUIDE_CN.md)

最后的API端点（Endpoint）是

```http://127.0.0.1:8080/v1/chat/completions```

注意域名/IP和端口要改成你自己的

### 环境变量
  - `PUID` - 用户ID
  - `http_proxy` - SOCKS5 或 HTTP 代理 `socks5://HOST:PORT`
  - `SERVER_HOST` - (default)比如 127.0.0.1
  - `SERVER_PORT` - (default)比如 8080 by

### 文件选项
  - `access_tokens.json` - 附带AccessToken的Json文件
  - `proxies.txt` - 代理表 (格式: `USERNAME:PASSWORD:HOST:PORT`)
</file>

<file path="README_JA.md">
# ChatGPT-to-API
ChatGPT のウェブサイトを使って偽 API を作る

> ## 重要
> このリポジトリに対する無償のサポートは受けられません。これは私個人の使用のために作られたもので、ドキュメントは本当に必要ないので、ドキュメントは制限され続けます。貢献者による中国語のドキュメントに、より詳細なドキュメントがあります。

**API エンドポイント: http://127.0.0.1:8080/v1/chat/completions.**

[英語ドキュメント（English Docs）](README.md)
[中国語ドキュメント（Chinese Docs）](https://github.com/acheong08/ChatGPT-to-API/blob/master/README_ZH.md)
## セットアップ

### 認証

アクセストークンの取得は [OpenAIAuth](https://github.com/acheong08/OpenAIAuth/) により、アカウントのメールアドレスとパスワードで自動化されています。

`accounts.txt` - 改行で区切られたアカウントのリスト

フォーマット:
```
email:password
...
```

すべての認証されたアクセストークンは `access_tokens.json` に保存されます

アクセストークンは 7 日後に自動更新されます

注意！認証にはブロックされていない ip を使用してください。可能であれば、まず `https://chat.openai.com/` にログインして ip の可用性を確認してください。

### API認証（オプション）

OpenAI の API と同じような、この偽 API 用のカスタム API キー

`api_keys.txt` - 改行で区切られた API キーのリスト

フォーマット:
```
sk-123456
88888888
...
```

## 準備
```
git clone https://github.com/acheong08/ChatGPT-to-API
cd ChatGPT-to-API
go build
./freechatgpt
```

### 環境変数
  - `PUID` - chat.openai.com の Plus ユーザー向けのクッキーです。これは Cloudflare のレート制限を回避します
  - `SERVER_HOST` - デフォルトで 127.0.0.1 に設定
  - `SERVER_PORT` - デフォルトで 8080 に設定
  - `OPENAI_EMAIL` と `OPENAI_PASSWORD` - PUID が設定されている場合、自動的に更新されます
  - `ENABLE_HISTORY` - デフォルトで true に設定

### ファイル（オプション）
  - `proxies.txt` - 改行で区切られたプロキシのリスト

    ```
    http://127.0.0.1:8888
    ...
    ```
  - `access_tokens.json` - サイクリング用のアクセストークンの JSON 配列（あるいは、[正しいエンドポイント](https://github.com/acheong08/ChatGPT-to-API/blob/master/docs/admin.md)に PATCH リクエストを送る）
    ```
    ["access_token1", "access_token2"...]
    ```

## Admin API ドキュメント
https://github.com/acheong08/ChatGPT-to-API/blob/master/docs/admin.md

## API 使用方法ドキュメント
https://platform.openai.com/docs/api-reference/chat
</file>

<file path="README_ZH.md">
# ChatGPT-to-API
从ChatGPT网站模拟使用API

**模拟API地址: http://127.0.0.1:8080/v1/chat/completions.**

## 使用
    
### 设置

配置账户邮箱和密码，自动生成和更新Access tokens 和 PUID（仅PLUS账户）（使用[OpenAIAuth](https://github.com/acheong08/OpenAIAuth/)）

`accounts.txt` - 存放OpenAI账号邮箱和密码的文件

格式:
```
邮箱:密码
邮箱:密码
...
```

所有登录后的Access tokens和PUID会存放在`access_tokens.json`

每7天自动更新Access tokens和PUID

注意！ 请使用未封锁的ip登录账号，请先打开浏览器登录`https://chat.openai.com/`以检查ip是否可用

### GPT-4 设置（可选）

如果配置PLUS账户并使用GPT-4模型，则需要HAR文件（`chat.openai.com.har`）以完成captcha验证

1. 使用基于chromium的浏览器（Chrome，Edge）或Safari浏览器 登录`https://chat.openai.com/`，然后打开浏览器开发者工具（F12），并切换到网络标签页。

2. 新建聊天并选择GPT-4模型，随意问一个问题，点击网络标签页下的导出HAR按钮，导出文件`chat.openai.com.har`

### API 密钥（可选）

如OpenAI的官方API一样，可给模拟的API添加API密钥认证

`api_keys.txt` - 存放API密钥的文件

格式:
```
sk-123456
88888888
...
```

## 开始
```  
git clone https://github.com/acheong08/ChatGPT-to-API
cd ChatGPT-to-API
go build
./freechatgpt
```

### 环境变量
  - `PUID` - Plus账户可在`chat.openai.com`的cookies里找到，用于绕过cf的频率限制
  - `SERVER_HOST` - 默认127.0.0.1
  - `SERVER_PORT` - 默认8080
  - `ENABLE_HISTORY` - 默认true，允许网页端历史记录

### 可选文件配置
  - `proxies.txt` - 存放代理地址的文件

    ```
    http://127.0.0.1:8888
    socks5://127.0.0.1:9999
    ...
    ```
  - `access_tokens.json` - 一个存放Access tokens 和PUID JSON数组的文件 (可使用 PATCH请求更新Access tokens [correct endpoint](https://github.com/acheong08/ChatGPT-to-API/blob/master/docs/admin.md))
    ```
    [{token:"access_token1", puid:"puid1"}, {token:"access_token2", puid:"puid2"}...]
    ```

## 用户管理文档
https://github.com/acheong08/ChatGPT-to-API/blob/master/docs/admin.md

## API使用说明
https://platform.openai.com/docs/api-reference/chat
</file>

<file path="tools/authenticator/auth/auth.go">
package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"regexp"
	"strings"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	pkce "github.com/nirasan/go-oauth-pkce-code-verifier"
)

type Error struct {
	Location   string
	StatusCode int
	Details    string
	Error      error
}

func NewError(location string, statusCode int, details string, err error) *Error {
	return &Error{
		Location:   location,
		StatusCode: statusCode,
		Details:    details,
		Error:      err,
	}
}

type Authenticator struct {
	EmailAddress       string
	Password           string
	Proxy              string
	Session            tls_client.HttpClient
	UserAgent          string
	State              string
	URL                string
	Verifier_code      string
	Verifier_challenge string
	AuthResult         AuthResult
}

type AuthResult struct {
	AccessToken string `json:"access_token"`
	PUID        string `json:"puid"`
}

func NewAuthenticator(emailAddress, password, proxy string) *Authenticator {
	auth := &Authenticator{
		EmailAddress: emailAddress,
		Password:     password,
		Proxy:        proxy,
		UserAgent:    "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
	}
	jar := tls_client.NewCookieJar()
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(20),
		tls_client.WithClientProfile(tls_client.Okhttp4Android13),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar), // create cookieJar instance and pass it as argument
		// Proxy
		tls_client.WithProxyUrl(proxy),
	}
	auth.Session, _ = tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)

	// PKCE
	verifier, _ := pkce.CreateCodeVerifier()
	auth.Verifier_code = verifier.String()
	auth.Verifier_challenge = verifier.CodeChallengeS256()

	return auth
}

func (auth *Authenticator) URLEncode(str string) string {
	return url.QueryEscape(str)
}

func (auth *Authenticator) Begin() *Error {

	url := "https://chat.openai.com/api/auth/csrf"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return NewError("begin", 0, "", err)
	}

	req.Header.Set("Host", "chat.openai.com")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", auth.UserAgent)
	req.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Set("Referer", "https://chat.openai.com/auth/login")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")

	resp, err := auth.Session.Do(req)
	if err != nil {
		return NewError("begin", 0, "", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return NewError("begin", 0, "", err)
	}

	if resp.StatusCode == 200 && strings.Contains(resp.Header.Get("Content-Type"), "json") {

		var csrfTokenResponse struct {
			CsrfToken string `json:"csrfToken"`
		}
		err = json.Unmarshal(body, &csrfTokenResponse)
		if err != nil {
			return NewError("begin", 0, "", err)
		}

		csrfToken := csrfTokenResponse.CsrfToken
		return auth.partOne(csrfToken)
	} else {
		err := NewError("begin", resp.StatusCode, string(body), fmt.Errorf("error: Check details"))
		return err
	}
}

func (auth *Authenticator) partOne(csrfToken string) *Error {

	auth_url := "https://chat.openai.com/api/auth/signin/auth0?prompt=login"
	headers := map[string]string{
		"Host":            "chat.openai.com",
		"User-Agent":      auth.UserAgent,
		"Content-Type":    "application/x-www-form-urlencoded",
		"Accept":          "*/*",
		"Sec-Gpc":         "1",
		"Accept-Language": "en-US,en;q=0.8",
		"Origin":          "https://chat.openai.com",
		"Sec-Fetch-Site":  "same-origin",
		"Sec-Fetch-Mode":  "cors",
		"Sec-Fetch-Dest":  "empty",
		"Referer":         "https://chat.openai.com/auth/login",
		"Accept-Encoding": "gzip, deflate",
	}

	// Construct payload
	payload := fmt.Sprintf("callbackUrl=%%2F&csrfToken=%s&json=true", csrfToken)
	req, _ := http.NewRequest("POST", auth_url, strings.NewReader(payload))

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := auth.Session.Do(req)
	if err != nil {
		return NewError("part_one", 0, "Failed to send request", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return NewError("part_one", 0, "Failed to read body", err)
	}

	if resp.StatusCode == 200 && strings.Contains(resp.Header.Get("Content-Type"), "json") {
		var urlResponse struct {
			URL string `json:"url"`
		}
		err = json.Unmarshal(body, &urlResponse)
		if err != nil {
			return NewError("part_one", 0, "Failed to decode JSON", err)
		}
		if urlResponse.URL == "https://chat.openai.com/api/auth/error?error=OAuthSignin" || strings.Contains(urlResponse.URL, "error") {
			err := NewError("part_one", resp.StatusCode, "You have been rate limited. Please try again later.", fmt.Errorf("error: Check details"))
			return err
		}
		return auth.partTwo(urlResponse.URL)
	} else {
		return NewError("part_one", resp.StatusCode, string(body), fmt.Errorf("error: Check details"))
	}
}

func (auth *Authenticator) partTwo(url string) *Error {

	headers := map[string]string{
		"Host":            "auth0.openai.com",
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
		"Connection":      "keep-alive",
		"User-Agent":      auth.UserAgent,
		"Accept-Language": "en-US,en;q=0.9",
	}

	req, _ := http.NewRequest("GET", url, nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := auth.Session.Do(req)
	if err != nil {
		return NewError("part_two", 0, "Failed to make request", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode == 302 || resp.StatusCode == 200 {

		stateRegex := regexp.MustCompile(`state=(.*)`)
		stateMatch := stateRegex.FindStringSubmatch(string(body))
		if len(stateMatch) < 2 {
			return NewError("part_two", 0, "Could not find state in response", fmt.Errorf("error: Check details"))
		}

		state := strings.Split(stateMatch[1], `"`)[0]
		return auth.partThree(state)
	} else {
		return NewError("part_two", resp.StatusCode, string(body), fmt.Errorf("error: Check details"))

	}
}
func (auth *Authenticator) partThree(state string) *Error {

	url := fmt.Sprintf("https://auth0.openai.com/u/login/identifier?state=%s", state)
	emailURLEncoded := auth.URLEncode(auth.EmailAddress)

	payload := fmt.Sprintf(
		"state=%s&username=%s&js-available=false&webauthn-available=true&is-brave=false&webauthn-platform-available=true&action=default",
		state, emailURLEncoded,
	)

	headers := map[string]string{
		"Host":            "auth0.openai.com",
		"Origin":          "https://auth0.openai.com",
		"Connection":      "keep-alive",
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
		"User-Agent":      auth.UserAgent,
		"Referer":         fmt.Sprintf("https://auth0.openai.com/u/login/identifier?state=%s", state),
		"Accept-Language": "en-US,en;q=0.9",
		"Content-Type":    "application/x-www-form-urlencoded",
	}

	req, _ := http.NewRequest("POST", url, strings.NewReader(payload))

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := auth.Session.Do(req)
	if err != nil {
		return NewError("part_three", 0, "Failed to send request", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 302 || resp.StatusCode == 200 {
		return auth.partFour(state)
	} else {
		return NewError("part_three", resp.StatusCode, "Your email address is invalid.", fmt.Errorf("error: Check details"))

	}

}
func (auth *Authenticator) partFour(state string) *Error {

	url := fmt.Sprintf("https://auth0.openai.com/u/login/password?state=%s", state)
	emailURLEncoded := auth.URLEncode(auth.EmailAddress)
	passwordURLEncoded := auth.URLEncode(auth.Password)
	payload := fmt.Sprintf("state=%s&username=%s&password=%s&action=default", state, emailURLEncoded, passwordURLEncoded)

	headers := map[string]string{
		"Host":            "auth0.openai.com",
		"Origin":          "https://auth0.openai.com",
		"Connection":      "keep-alive",
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
		"User-Agent":      auth.UserAgent,
		"Referer":         fmt.Sprintf("https://auth0.openai.com/u/login/password?state=%s", state),
		"Accept-Language": "en-US,en;q=0.9",
		"Content-Type":    "application/x-www-form-urlencoded",
	}

	req, _ := http.NewRequest("POST", url, strings.NewReader(payload))

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := auth.Session.Do(req)
	if err != nil {
		return NewError("part_four", 0, "Failed to send request", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 302 {
		redirectURL := resp.Header.Get("Location")
		return auth.partFive(state, redirectURL)
	} else {
		body := bytes.NewBuffer(nil)
		body.ReadFrom(resp.Body)
		return NewError("part_four", resp.StatusCode, body.String(), fmt.Errorf("error: Check details"))

	}

}
func (auth *Authenticator) partFive(oldState string, redirectURL string) *Error {

	url := "https://auth0.openai.com" + redirectURL

	headers := map[string]string{
		"Host":            "auth0.openai.com",
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
		"Connection":      "keep-alive",
		"User-Agent":      auth.UserAgent,
		"Accept-Language": "en-GB,en-US;q=0.9,en;q=0.8",
		"Referer":         fmt.Sprintf("https://auth0.openai.com/u/login/password?state=%s", oldState),
	}

	req, _ := http.NewRequest("GET", url, nil)

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := auth.Session.Do(req)
	if err != nil {
		return NewError("part_five", 0, "Failed to send request", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 302 {
		return auth.partSix(resp.Header.Get("Location"), url)
	} else {
		return NewError("part_five", resp.StatusCode, resp.Status, fmt.Errorf("error: Check details"))

	}

}
func (auth *Authenticator) partSix(url, redirect_url string) *Error {
	req, _ := http.NewRequest("GET", url, nil)
	for k, v := range map[string]string{
		"Host":            "chat.openai.com",
		"Accept":          "application/json",
		"Connection":      "keep-alive",
		"User-Agent":      auth.UserAgent,
		"Accept-Language": "en-GB,en-US;q=0.9,en;q=0.8",
		"Referer":         redirect_url,
	} {
		req.Header.Set(k, v)
	}
	resp, err := auth.Session.Do(req)
	if err != nil {
		return NewError("part_six", 0, "Failed to send request", err)
	}
	defer resp.Body.Close()
	if err != nil {
		return NewError("part_six", 0, "Response was not JSON", err)
	}
	if resp.StatusCode != 302 {
		return NewError("part_six", resp.StatusCode, url, fmt.Errorf("incorrect response code"))
	}
	// Check location header
	if location := resp.Header.Get("Location"); location != "https://chat.openai.com/" {
		return NewError("part_six", resp.StatusCode, location, fmt.Errorf("incorrect redirect"))
	}

	url = "https://chat.openai.com/api/auth/session"

	req, _ = http.NewRequest("GET", url, nil)

	// Set user agent
	req.Header.Set("User-Agent", auth.UserAgent)

	resp, err = auth.Session.Do(req)
	if err != nil {
		return NewError("get_access_token", 0, "Failed to send request", err)
	}

	if resp.StatusCode != 200 {
		return NewError("get_access_token", resp.StatusCode, "Incorrect response code", fmt.Errorf("error: Check details"))
	}
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return NewError("get_access_token", 0, "", err)
	}

	// Check if access token in data
	if _, ok := result["accessToken"]; !ok {
		result_string := fmt.Sprintf("%v", result)
		return NewError("part_six", 0, result_string, fmt.Errorf("missing access token"))
	}
	auth.AuthResult.AccessToken = result["accessToken"].(string)

	return nil
}

func (auth *Authenticator) GetAccessToken() string {
	return auth.AuthResult.AccessToken
}

func (auth *Authenticator) GetPUID() (string, *Error) {
	// Check if user has access token
	if auth.AuthResult.AccessToken == "" {
		return "", NewError("get_puid", 0, "Missing access token", fmt.Errorf("error: Check details"))
	}
	// Make request to https://chat.openai.com/backend-api/models
	req, _ := http.NewRequest("GET", "https://chat.openai.com/backend-api/models", nil)
	// Add headers
	req.Header.Add("Authorization", "Bearer "+auth.AuthResult.AccessToken)
	req.Header.Add("User-Agent", auth.UserAgent)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9")
	req.Header.Add("Referer", "https://chat.openai.com/")
	req.Header.Add("Origin", "https://chat.openai.com")
	req.Header.Add("Connection", "keep-alive")

	resp, err := auth.Session.Do(req)
	if err != nil {
		return "", NewError("get_puid", 0, "Failed to make request", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", NewError("get_puid", resp.StatusCode, "Failed to make request", fmt.Errorf("error: Check details"))
	}
	// Find `_puid` cookie in response
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "_puid" {
			auth.AuthResult.PUID = cookie.Value
			return cookie.Value, nil
		}
	}
	// If cookie not found, return error
	return "", NewError("get_puid", 0, "PUID cookie not found", fmt.Errorf("error: Check details"))
}

func (auth *Authenticator) GetAuthResult() AuthResult {
	return auth.AuthResult
}
</file>

<file path="tools/authenticator/go.mod">
module authenticator

go 1.20

require (
	github.com/bogdanfinn/fhttp v0.5.23
	github.com/bogdanfinn/tls-client v1.5.0
	github.com/nirasan/go-oauth-pkce-code-verifier v0.0.0-20220510032225-4f9f17eaec4c
)

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/bogdanfinn/utls v1.5.16 // indirect
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/tam7t/hpkp v0.0.0-20160821193359-2b70b4024ed5 // indirect
	golang.org/x/crypto v0.11.0 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
)
</file>

<file path="tools/authenticator/main.go">
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	auth "authenticator/auth"
)

type Account struct {
	Email    string `json:"username"`
	Password string `json:"password"`
}
type Proxy struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

func (p Proxy) Socks5URL() string {
	// Returns proxy URL (socks5)
	if p.User == "" && p.Pass == "" {
		return fmt.Sprintf("socks5://%s:%s", p.IP, p.Port)
	}
	return fmt.Sprintf("socks5://%s:%s@%s:%s", p.User, p.Pass, p.IP, p.Port)
}

// Read accounts.txt and create a list of accounts
func readAccounts() []Account {
	accounts := []Account{}
	// Read accounts.txt and create a list of accounts
	file, err := os.Open("accounts.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// Loop through each line in the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split by :
		line := strings.Split(scanner.Text(), ":")
		// Create an account
		account := Account{
			Email:    line[0],
			Password: line[1],
		}
		// Append to accounts
		accounts = append(accounts, account)
	}
	return accounts
}

// Read proxies from proxies.txt and create a list of proxies
func readProxies() []Proxy {
	proxies := []Proxy{}
	// Read proxies.txt and create a list of proxies
	file, err := os.Open("proxies.txt")
	if err != nil {
		return []Proxy{}
	}
	defer file.Close()
	// Loop through each line in the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split by :
		lines := strings.Split(scanner.Text(), ":")
		var proxy Proxy
		if len(lines) == 4 {
			// Create a proxy
			proxy = Proxy{
				IP:   lines[0],
				Port: lines[1],
				User: lines[2],
				Pass: lines[3],
			}
		} else if len(lines) == 2 {
			proxy = Proxy{
				IP:   lines[0],
				Port: lines[1],
			}
		} else {
			continue
		}
		// Append to proxies
		proxies = append(proxies, proxy)
	}
	return proxies
}

func main() {
	// Read accounts and proxies
	accounts := readAccounts()
	proxies := readProxies()

	// Loop through each account
	for _, account := range accounts {
		if os.Getenv("CF_PROXY") != "" {
			// exec warp-cli disconnect and connect
			exec.Command("warp-cli", "disconnect").Run()
			exec.Command("warp-cli", "connect").Run()
			time.Sleep(5 * time.Second)
		}
		println(account.Email)
		println(account.Password)
		var proxy_url string
		if len(proxies) == 0 {
			if os.Getenv("http_proxy") != "" {
				proxy_url = os.Getenv("http_proxy")
			}
		} else {
			proxy_url = proxies[0].Socks5URL()
			// Push used proxy to the back of the list
			proxies = append(proxies[1:], proxies[0])
		}
		println(proxy_url)
		authenticator := auth.NewAuthenticator(account.Email, account.Password, proxy_url)
		err := authenticator.Begin()
		if err != nil {
			// println("Error: " + err.Details)
			println("Location: " + err.Location)
			println("Status code: " + fmt.Sprint(err.StatusCode))
			println("Details: " + err.Details)
			println("Embedded error: " + err.Error.Error())
			// Sleep for 10 seconds
			panic(err)
		}
		access_token := authenticator.GetAccessToken()
		// Append access token to access_tokens.txt
		f, go_err := os.OpenFile("access_tokens.txt", os.O_APPEND|os.O_WRONLY, 0600)
		if go_err != nil {
			continue
		}
		defer f.Close()
		if _, go_err = f.WriteString(access_token + "\n"); go_err != nil {
			continue
		}
		// Write authenticated account to authenticated_accounts.txt
		f, go_err = os.OpenFile("authenticated_accounts.txt", os.O_APPEND|os.O_WRONLY, 0600)
		if go_err != nil {
			continue
		}
		defer f.Close()
		if _, go_err = f.WriteString(account.Email + ":" + account.Password + "\n"); go_err != nil {
			continue
		}
		// Remove accounts.txt
		os.Remove("accounts.txt")
		// Create accounts.txt
		f, go_err = os.Create("accounts.txt")
		if go_err != nil {
			continue
		}
		defer f.Close()
		// Remove account from accounts
		accounts = accounts[1:]
		// Write unauthenticated accounts to accounts.txt
		for _, acc := range accounts {
			// Check if account is authenticated
			if acc.Email == account.Email {
				continue
			}
			if _, go_err = f.WriteString(acc.Email + ":" + acc.Password + "\n"); go_err != nil {
				continue
			}
		}

	}
}
</file>

<file path="tools/authenticator/README.md">
# Automated authentication for ChatGPT
Fetches access tokens from a large number of accounts

## Setup
### `proxies.txt`
Format:
```
IP:Port:User:Password
...
```

### `accounts.txt`
Format:
```
email:password
...
```

Remember to:
`touch access_tokens.txt authenticated_accounts.txt`
</file>

<file path="tools/authenticator/remove_duplicates.py">
# Removes duplicate lines from a file
# Usage: python remove_duplicates.py <file>

import sys
import json


def remove_duplicates(file_lines):
    """
    Removes duplicate lines from a file
    """
    lines_set = set()
    for lin in file_lines:
        #if json.loads(lin)["output"] == "":
        #    continue
        lines_set.add(lin)
    return lines_set


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python remove_duplicates.py <file>")
        sys.exit(1)
    orig_file = open(sys.argv[1], "r", encoding="utf-8").readlines()
    lines = remove_duplicates(orig_file)
    file = open("clean_" + sys.argv[1], "w", encoding="utf-8")
    for line in lines:
        file.write(line)
    file.close()
    # Print difference
    print(len(orig_file) - len(lines))
</file>

<file path="tools/authenticator/tojson.sh">
#!/bin/bash

# Check if a file name is provided as an argument
if [ $# -eq 0 ]; then
    echo "Usage: $0 <file>"
    exit 1
fi

file="$1"
output="$2"
# Declare an empty array
lines=()

# Read the file line by line and add each line to the array
while IFS= read -r line; do
    lines+=("\"$line\"")
done < "$file"

# Join array elements with commas and print the result enclosed in square brackets
result="["
for ((i = 0; i < ${#lines[@]}; i++)); do
    result+="${lines[i]}"
    if ((i < ${#lines[@]} - 1)); then
        result+=","
    fi
done
result+="]"
if [ $# -eq 1 ]; then
	echo "$result"
fi
if [ $# -eq 2 ]; then
	echo "$result" | tee $output
fi
</file>

<file path="tools/plugin_check/check_access.go">
package main

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"
)

func main() {
	var access_tokens []string
	// Read access_tokens.txt and split by new line
	file, err := os.Open("access_tokens.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		access_tokens = append(access_tokens, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	// Go routine to check access for each token (limit to 20 simultaneous)
	sem := make(chan bool, 20)
	for _, token := range access_tokens {
		sem <- true
		go func(token string) {
			defer func() { <-sem }()
			if check_access(token) {
				println(token)
			}
		}(token)
	}
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
}

func check_access(token string) bool {
	print(".")
	req, _ := http.NewRequest("GET", "https://ai.fakeopen.com/api/accounts/check", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		// Parse response body as JSON
		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)
		// Check if "tool1", "tool2", or "tool3" is in the features array
		for _, feature := range result["features"].([]interface{}) {
			if feature == "tools1" || feature == "tools2" || feature == "tools3" {
				return true
			}
		}
		return false
	}
	println(resp.StatusCode)
	return false
}
</file>

<file path="tools/proxy_check/proxy_check.go">
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	tls_client "github.com/bogdanfinn/tls-client"
)

var proxies []string

// Read proxies.txt and check if they work
func init() {
	// Check for proxies.txt
	if _, err := os.Stat("proxies.txt"); err == nil {
		// Each line is a proxy, put in proxies array
		file, _ := os.Open("proxies.txt")
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			// Split line by :
			proxy := scanner.Text()
			proxy_parts := strings.Split(proxy, ":")
			if len(proxy_parts) == 2 {
				proxy = "socks5://" + proxy
			} else if len(proxy_parts) == 4 {
				proxy = "socks5://" + proxy_parts[2] + ":" + proxy_parts[3] + "@" + proxy_parts[0] + ":" + proxy_parts[1]
			} else {
				continue
			}
			proxies = append(proxies, proxy)
		}
	}

}

func main() {
	wg := sync.WaitGroup{}
	for _, proxy := range proxies {
		wg.Add(1)
		go func(proxy string) {
			defer wg.Done()
			jar := tls_client.NewCookieJar()
			options := []tls_client.HttpClientOption{
				tls_client.WithTimeoutSeconds(360),
				tls_client.WithClientProfile(tls_client.Chrome_110),
				tls_client.WithNotFollowRedirects(),
				tls_client.WithCookieJar(jar), // create cookieJar instance and pass it as argument
				// Disable SSL verification
				tls_client.WithInsecureSkipVerify(),
			}
			client, _ := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)

			client.SetProxy(proxy)
			resp, err := client.Get("https://example.com")
			if err != nil {
				fmt.Println("Error: ", err)
				fmt.Println("Proxy: ", proxy)
				return
			}
			if resp.StatusCode != 200 {
				fmt.Println("Error: ", resp.StatusCode)
				fmt.Println("Proxy: ", proxy)
				return
			} else {
				fmt.Println(".")
			}
		}(proxy)
	}
	wg.Wait()
}
</file>

<file path=".gitignore">
tools/authenticator/100-ACCOUNTS_COMPILED.txt
tools/authenticator/accounts.txt
tools/authenticator/proxies.txt
tools/authenticator/authenticated_accounts.txt
tools/authenticator/access_tokens.txt
*.txt
access_tokens.json
freechatgpt
chatgpttoapi
tools/authenticator/.proxies.txt.swp
.env
*.har
</file>

<file path="README.md">
# ChatGPT-to-API
Create a fake API using ChatGPT's website

> ## IMPORTANT
> You will not get free support for this repository. This was made for my own personal use and documentation will continue to be limited as I don't really need documentation. You will find more detailed documentation in the Chinese docs by a contributor.

**API endpoint: http://127.0.0.1:8080/v1/chat/completions.**

[中文文档（Chinese Docs）](https://github.com/acheong08/ChatGPT-to-API/blob/master/README_ZH.md)
## Setup
    
### Authentication

Access token and PUID(only for PLUS account) retrieval has been automated by [OpenAIAuth](https://github.com/acheong08/OpenAIAuth/) with account email & password.

`accounts.txt` - A list of accounts separated by new line 

Format:
```
email:password
...
```

All authenticated access tokens and PUID will store in `access_tokens.json`

Auto renew access tokens and PUID after 7 days

Caution! please use unblocked ip for authentication, first login to `https://chat.openai.com/` to check ip availability if you can.

### GPT-4 Model (Optional)

If you configured a PLUS account and use the GPT-4 model, a HAR file (`chat.openai.com.har`) is required to complete CAPTCHA verification

1. Use a chromium-based browser (Chrome, Edge) or Safari to login to `https://chat.openai.com/`, then open the browser developer tools (F12), and switch to the Network tab.

2. Create a new chat and select the GPT-4 model, ask a question at will, click the Export HAR button under the Network tab, export the file `chat.openai.com.har`

### API Authentication (Optional)

Custom API keys for this fake API, just like OpenAI api

`api_keys.txt` - A list of API keys separated by new line

Format:
```
sk-123456
88888888
...
```

## Getting set up
```  
git clone https://github.com/acheong08/ChatGPT-to-API
cd ChatGPT-to-API
go build
./freechatgpt
```

### Environment variables
  - `PUID` - A cookie found on chat.openai.com for Plus users. This gets around Cloudflare rate limits
  - `SERVER_HOST` - Set to 127.0.0.1 by default
  - `SERVER_PORT` - Set to 8080 by default
  - `ENABLE_HISTORY` - Set to true by default

### Files (Optional)
  - `proxies.txt` - A list of proxies separated by new line

    ```
    http://127.0.0.1:8888
    ...
    ```
  - `access_tokens.json` - A JSON array of access tokens for cycling (Alternatively, send a PATCH request to the [correct endpoint](https://github.com/acheong08/ChatGPT-to-API/blob/master/docs/admin.md))
    ```
    [{token:"access_token1", puid:"puid1"}, {token:"access_token2", puid:"puid2"}...]
    ```

## Admin API docs
https://github.com/acheong08/ChatGPT-to-API/blob/master/docs/admin.md

## API usage docs
https://platform.openai.com/docs/api-reference/chat
</file>

</files>
