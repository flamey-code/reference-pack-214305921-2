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
.github/workflows/docker-push.yaml
.gitignore
api/chatgpt/api.go
api/chatgpt/constant.go
api/chatgpt/health_check.go
api/chatgpt/login.go
api/chatgpt/typings.go
api/common.go
api/imitate/api.go
api/imitate/convert.go
api/imitate/request.go
api/imitate/response.go
api/platform/access_token.go
api/platform/api.go
api/platform/constant.go
api/platform/login.go
api/platform/typings.go
compose.yaml
Dockerfile
env/env.go
example/chatgpt.http
example/imitate.http
example/ios.http
example/platform.http
go.mod
LICENSE
main.go
middleware/authorization.go
middleware/cors.go
README.md
render.yaml
```

# Files

## File: .env.example
````
PORT=8080
TZ=Asia/Shanghai
PROXY=
OPENAI_EMAIL=
OPENAI_PASSWORD=
CONTINUE_SIGNAL=
ENABLE_HISTORY=
IMITATE_ACCESS_TOKEN=
````

## File: .github/workflows/docker-push.yaml
````yaml
name: Docker build and push

on:
  push:
    branches:
      - 'main'

env:
  PLATFORMS: ${{ vars.PLATFORMS || 'linux/amd64,linux/arm64' }}

jobs:
  docker-build-push:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v3

      - name: Set up QEMU
        uses: docker/setup-qemu-action@v2

      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v2

      - name: Cache Docker layers
        uses: actions/cache@v2
        with:
          path: /tmp/.buildx-cache
          key: ${{ runner.os }}-buildx-${{ github.sha }}
          restore-keys: |
            ${{ runner.os }}-buildx-

      - name: Login to DockerHub
        uses: docker/login-action@v2
        with:
          username: ${{ github.actor }}
          password: ${{ secrets.DOCKER_HUB_TOKEN }}

      - name: Build and push
        uses: docker/build-push-action@v4
        with:
          context: .
          platforms: ${{ env.PLATFORMS }}
          push: true
          tags: ${{ github.actor }}/go-chatgpt-api
          cache-from: type=local,src=/tmp/.buildx-cache
          cache-to: type=local,dest=/tmp/.buildx-cache,mode=max

      - name: Log into ghcr
        uses: docker/login-action@v2
        if: ${{ vars.USE_GHCR == '1' }}
        with:
          registry: ghcr.io
          username: ${{ github.actor }}
          password: ${{ secrets.GITHUB_TOKEN }}

      - name: Build and push to ghcr
        uses: docker/build-push-action@v4
        if: ${{ vars.USE_GHCR == '1' }}
        with:
          context: .
          platforms: ${{ env.PLATFORMS }}
          push: true
          tags: ghcr.io/${{ github.actor }}/go-chatgpt-api
          cache-from: type=local,src=/tmp/.buildx-cache
          cache-to: type=local,dest=/tmp/.buildx-cache,mode=max
````

## File: .gitignore
````
.env
go-chatgpt-api
__debug*
chat.openai.com.har
.idea
````

## File: api/chatgpt/api.go
````go
package chatgpt

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	http "github.com/bogdanfinn/fhttp"
	"github.com/gin-gonic/gin"

	"github.com/linweiyuan/go-chatgpt-api/api"
	"github.com/linweiyuan/go-logger/logger"
)

func CreateConversation(c *gin.Context) {
	var request CreateConversationRequest
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.ReturnMessage(parseJsonErrorMessage))
		return
	}

	if request.ConversationID == nil || *request.ConversationID == "" {
		request.ConversationID = nil
	}

	if len(request.Messages) != 0 {
		message := request.Messages[0]
		if message.Author.Role == "" {
			message.Author.Role = defaultRole
		}

		if message.Metadata == nil {
			message.Metadata = map[string]string{}
		}

		request.Messages[0] = message
	}

	if strings.HasPrefix(request.Model, gpt4Model) && request.ArkoseToken == "" {
		arkoseToken, err := api.GetArkoseToken()
		if err != nil || arkoseToken == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, api.ReturnMessage(err.Error()))
			return
		}

		request.ArkoseToken = arkoseToken
	}

	resp, done := sendConversationRequest(c, request)
	if done {
		return
	}

	handleConversationResponse(c, resp, request)
}

func sendConversationRequest(c *gin.Context, request CreateConversationRequest) (*http.Response, bool) {
	jsonBytes, _ := json.Marshal(request)
	req, _ := http.NewRequest(http.MethodPost, api.ChatGPTApiUrlPrefix+"/backend-api/conversation", bytes.NewBuffer(jsonBytes))
	req.Header.Set("User-Agent", api.UserAgent)
	req.Header.Set(api.AuthorizationHeader, api.GetAccessToken(c))
	req.Header.Set("Accept", "text/event-stream")
	if api.PUID != "" {
		req.Header.Set("Cookie", "_puid="+api.PUID)
	}
	resp, err := api.Client.Do(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.ReturnMessage(err.Error()))
		return nil, true
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusUnauthorized {
			logger.Error(fmt.Sprintf(api.AccountDeactivatedErrorMessage, c.GetString(api.EmailKey)))
			responseMap := make(map[string]interface{})
			json.NewDecoder(resp.Body).Decode(&responseMap)
			c.AbortWithStatusJSON(resp.StatusCode, responseMap)
			return nil, true
		}

		req, _ := http.NewRequest(http.MethodGet, api.ChatGPTApiUrlPrefix+"/backend-api/models?history_and_training_disabled=false", nil)
		req.Header.Set("User-Agent", api.UserAgent)
		req.Header.Set(api.AuthorizationHeader, api.GetAccessToken(c))
		response, err := api.Client.Do(req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, api.ReturnMessage(err.Error()))
			return nil, true
		}

		defer response.Body.Close()
		modelAvailable := false
		var getModelsResponse GetModelsResponse
		json.NewDecoder(response.Body).Decode(&getModelsResponse)
		for _, model := range getModelsResponse.Models {
			if model.Slug == request.Model {
				modelAvailable = true
				break
			}
		}
		if !modelAvailable {
			c.AbortWithStatusJSON(http.StatusForbidden, api.ReturnMessage(noModelPermissionErrorMessage))
			return nil, true
		}

		data, _ := io.ReadAll(resp.Body)
		logger.Warn(string(data))

		responseMap := make(map[string]interface{})
		json.NewDecoder(resp.Body).Decode(&responseMap)
		c.AbortWithStatusJSON(resp.StatusCode, responseMap)
		return nil, true
	}

	return resp, false
}

func handleConversationResponse(c *gin.Context, resp *http.Response, request CreateConversationRequest) {
	c.Writer.Header().Set("Content-Type", "text/event-stream; charset=utf-8")

	isMaxTokens := false
	continueParentMessageID := ""
	continueConversationID := ""

	defer resp.Body.Close()
	reader := bufio.NewReader(resp.Body)
	for {
		if c.Request.Context().Err() != nil {
			break
		}

		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "event") ||
			strings.HasPrefix(line, "data: 20") ||
			strings.HasPrefix(line, `data: {"conversation_id"`) ||
			line == "" {
			continue
		}

		responseJson := line[6:]
		if strings.HasPrefix(responseJson, "[DONE]") && isMaxTokens && request.AutoContinue {
			continue
		}

		// no need to unmarshal every time, but if response content has this "max_tokens", need to further check
		if strings.TrimSpace(responseJson) != "" && strings.Contains(responseJson, responseTypeMaxTokens) {
			var createConversationResponse CreateConversationResponse
			json.Unmarshal([]byte(responseJson), &createConversationResponse)
			message := createConversationResponse.Message
			if message.Metadata.FinishDetails.Type == responseTypeMaxTokens && createConversationResponse.Message.Status == responseStatusFinishedSuccessfully {
				isMaxTokens = true
				continueParentMessageID = message.ID
				continueConversationID = createConversationResponse.ConversationID
			}
		}

		c.Writer.Write([]byte(line + "\n\n"))
		c.Writer.Flush()
	}

	if isMaxTokens && request.AutoContinue {
		continueConversationRequest := CreateConversationRequest{
			ArkoseToken:                request.ArkoseToken,
			HistoryAndTrainingDisabled: request.HistoryAndTrainingDisabled,
			Model:                      request.Model,
			TimezoneOffsetMin:          request.TimezoneOffsetMin,

			Action:          actionContinue,
			ParentMessageID: continueParentMessageID,
			ConversationID:  &continueConversationID,
		}
		resp, done := sendConversationRequest(c, continueConversationRequest)
		if done {
			return
		}

		handleConversationResponse(c, resp, continueConversationRequest)
	}
}
````

## File: api/chatgpt/constant.go
````go
package chatgpt

const (
	defaultRole           = "user"
	parseJsonErrorMessage = "failed to parse json request body"

	gpt4Model                          = "gpt-4"
	actionContinue                     = "continue"
	responseTypeMaxTokens              = "max_tokens"
	responseStatusFinishedSuccessfully = "finished_successfully"
	noModelPermissionErrorMessage      = "you have no permission to use this model"
)
````

## File: api/chatgpt/health_check.go
````go
package chatgpt

import (
	"fmt"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
	http "github.com/bogdanfinn/fhttp"

	"github.com/linweiyuan/go-chatgpt-api/api"
	"github.com/linweiyuan/go-logger/logger"
)

const (
	healthCheckUrl         = "https://chat.openai.com/backend-api/accounts/check"
	errorHintBlock         = "looks like you have bean blocked by OpenAI, please change to a new IP or have a try with WARP"
	errorHintFailedToStart = "failed to start, please try again later: %s"
	sleepHours             = 8760 // 365 days
)

func init() {
	proxyUrl := os.Getenv("PROXY")
	if proxyUrl != "" {
		logger.Info("PROXY: " + proxyUrl)
		api.Client.SetProxy(proxyUrl)

		for {
			resp, err := healthCheck()
			if err != nil {
				// wait for proxy to be ready
				time.Sleep(time.Second)
				continue
			}

			checkHealthCheckStatus(resp)
			break
		}
	} else {
		resp, err := healthCheck()
		if err != nil {
			logger.Error("failed to health check: " + err.Error())
			os.Exit(1)
		}

		checkHealthCheckStatus(resp)
	}
}

func healthCheck() (resp *http.Response, err error) {
	req, _ := http.NewRequest(http.MethodGet, healthCheckUrl, nil)
	req.Header.Set("User-Agent", api.UserAgent)
	resp, err = api.Client.Do(req)
	return
}

func checkHealthCheckStatus(resp *http.Response) {
	if resp != nil {
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusUnauthorized {
			logger.Info(api.ReadyHint)
		} else {
			doc, _ := goquery.NewDocumentFromReader(resp.Body)
			alert := doc.Find(".message").Text()
			if alert != "" {
				logger.Error(errorHintBlock)
			} else {
				logger.Error(fmt.Sprintf(errorHintFailedToStart, resp.Status))
			}
			time.Sleep(time.Hour * sleepHours)
			os.Exit(1)
		}
	}
}
````

## File: api/chatgpt/login.go
````go
package chatgpt

import (
	http "github.com/bogdanfinn/fhttp"
	"github.com/gin-gonic/gin"
	"github.com/xqdoo00o/OpenAIAuth/auth"

	"github.com/linweiyuan/go-chatgpt-api/api"
)

func Login(c *gin.Context) {
	var loginInfo api.LoginInfo
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.ReturnMessage(api.ParseUserInfoErrorMessage))
		return
	}

	authenticator := auth.NewAuthenticator(loginInfo.Username, loginInfo.Password, api.ProxyUrl)
	if err := authenticator.Begin(); err != nil {
		c.AbortWithStatusJSON(err.StatusCode, api.ReturnMessage(err.Details))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken": authenticator.GetAccessToken(),
	})
}
````

## File: api/chatgpt/typings.go
````go
package chatgpt

import (
	"github.com/google/uuid"
)

type CreateConversationRequest struct {
	Action                     string    `json:"action"`
	Messages                   []Message `json:"messages"`
	Model                      string    `json:"model"`
	ParentMessageID            string    `json:"parent_message_id"`
	ConversationID             *string   `json:"conversation_id"`
	PluginIDs                  []string  `json:"plugin_ids"`
	TimezoneOffsetMin          int       `json:"timezone_offset_min"`
	ArkoseToken                string    `json:"arkose_token"`
	HistoryAndTrainingDisabled bool      `json:"history_and_training_disabled"`
	AutoContinue               bool      `json:"auto_continue"`
	Suggestions                []string  `json:"suggestions"`
}

func (c *CreateConversationRequest) AddMessage(role string, content string) {
	c.Messages = append(c.Messages, Message{
		ID:       uuid.New().String(),
		Author:   Author{Role: role},
		Content:  Content{ContentType: "text", Parts: []interface{}{content}},
		Metadata: map[string]string{},
	})
}

type Message struct {
	Author   Author      `json:"author"`
	Content  Content     `json:"content"`
	ID       string      `json:"id"`
	Metadata interface{} `json:"metadata"`
}

type Author struct {
	Role string `json:"role"`
}

type Content struct {
	ContentType string        `json:"content_type"`
	Parts       []interface{} `json:"parts"`
}

type CreateConversationResponse struct {
	Message struct {
		ID     string `json:"id"`
		Author struct {
			Role     string      `json:"role"`
			Name     interface{} `json:"name"`
			Metadata struct {
			} `json:"metadata"`
		} `json:"author"`
		CreateTime float64     `json:"create_time"`
		UpdateTime interface{} `json:"update_time"`
		Content    struct {
			ContentType string   `json:"content_type"`
			Parts       []string `json:"parts"`
		} `json:"content"`
		Status   string  `json:"status"`
		EndTurn  bool    `json:"end_turn"`
		Weight   float64 `json:"weight"`
		Metadata struct {
			MessageType   string `json:"message_type"`
			ModelSlug     string `json:"model_slug"`
			FinishDetails struct {
				Type string `json:"type"`
			} `json:"finish_details"`
		} `json:"metadata"`
		Recipient string `json:"recipient"`
	} `json:"message"`
	ConversationID string      `json:"conversation_id"`
	Error          interface{} `json:"error"`
}

type GetModelsResponse struct {
	Models []struct {
		Slug         string   `json:"slug"`
		MaxTokens    int      `json:"max_tokens"`
		Title        string   `json:"title"`
		Description  string   `json:"description"`
		Tags         []string `json:"tags"`
		Capabilities struct {
		} `json:"capabilities"`
		EnabledTools []string `json:"enabled_tools,omitempty"`
	} `json:"models"`
	Categories []struct {
		Category             string `json:"category"`
		HumanCategoryName    string `json:"human_category_name"`
		SubscriptionLevel    string `json:"subscription_level"`
		DefaultModel         string `json:"default_model"`
		CodeInterpreterModel string `json:"code_interpreter_model"`
		PluginsModel         string `json:"plugins_model"`
	} `json:"categories"`
}
````

## File: api/common.go
````go
package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
	"github.com/gin-gonic/gin"
	"github.com/xqdoo00o/OpenAIAuth/auth"
	"github.com/xqdoo00o/funcaptcha"

	"github.com/linweiyuan/go-logger/logger"
)

const (
	ChatGPTApiPrefix    = "/chatgpt"
	ImitateApiPrefix    = "/imitate/v1"
	ChatGPTApiUrlPrefix = "https://chat.openai.com"

	PlatformApiPrefix    = "/platform"
	PlatformApiUrlPrefix = "https://api.openai.com"

	defaultErrorMessageKey             = "errorMessage"
	AuthorizationHeader                = "Authorization"
	XAuthorizationHeader               = "X-Authorization"
	ContentType                        = "application/x-www-form-urlencoded"
	UserAgent                          = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36"
	Auth0Url                           = "https://auth0.openai.com"
	LoginUsernameUrl                   = Auth0Url + "/u/login/identifier?state="
	LoginPasswordUrl                   = Auth0Url + "/u/login/password?state="
	ParseUserInfoErrorMessage          = "failed to parse user login info"
	GetAuthorizedUrlErrorMessage       = "failed to get authorized url"
	EmailInvalidErrorMessage           = "email is not valid"
	EmailOrPasswordInvalidErrorMessage = "email or password is not correct"
	GetAccessTokenErrorMessage         = "failed to get access token"
	defaultTimeoutSeconds              = 600 // 10 minutes

	EmailKey                       = "email"
	AccountDeactivatedErrorMessage = "account %s is deactivated"

	ReadyHint = "service go-chatgpt-api is ready"

	refreshPuidErrorMessage = "failed to refresh PUID"
)

var (
	Client       tls_client.HttpClient
	ArkoseClient tls_client.HttpClient
	PUID         string
	ProxyUrl     string
)

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthLogin interface {
	GetAuthorizedUrl(csrfToken string) (string, int, error)
	GetState(authorizedUrl string) (string, int, error)
	CheckUsername(state string, username string) (int, error)
	CheckPassword(state string, username string, password string) (string, int, error)
	GetAccessToken(code string) (string, int, error)
	GetAccessTokenFromHeader(c *gin.Context) (string, int, error)
}

func init() {
	Client, _ = tls_client.NewHttpClient(tls_client.NewNoopLogger(), []tls_client.HttpClientOption{
		tls_client.WithCookieJar(tls_client.NewCookieJar()),
		tls_client.WithTimeoutSeconds(defaultTimeoutSeconds),
		tls_client.WithClientProfile(profiles.Okhttp4Android13),
	}...)
	ArkoseClient = getHttpClient()

	setupPUID()
}

func NewHttpClient() tls_client.HttpClient {
	client := getHttpClient()

	ProxyUrl = os.Getenv("PROXY")
	if ProxyUrl != "" {
		client.SetProxy(ProxyUrl)
	}

	return client
}

func getHttpClient() tls_client.HttpClient {
	client, _ := tls_client.NewHttpClient(tls_client.NewNoopLogger(), []tls_client.HttpClientOption{
		tls_client.WithCookieJar(tls_client.NewCookieJar()),
		tls_client.WithClientProfile(profiles.Okhttp4Android13),
	}...)
	return client
}

func Proxy(c *gin.Context) {
	url := c.Request.URL.Path
	if strings.Contains(url, ChatGPTApiPrefix) {
		url = strings.ReplaceAll(url, ChatGPTApiPrefix, ChatGPTApiUrlPrefix)
	} else if strings.Contains(url, ImitateApiPrefix) {
		url = strings.ReplaceAll(url, ImitateApiPrefix, ChatGPTApiUrlPrefix+"/backend-api")
	} else {
		url = strings.ReplaceAll(url, PlatformApiPrefix, PlatformApiUrlPrefix)
	}

	method := c.Request.Method
	queryParams := c.Request.URL.Query().Encode()
	if queryParams != "" {
		url += "?" + queryParams
	}

	// if not set, will return 404
	c.Status(http.StatusOK)

	var req *http.Request
	if method == http.MethodGet {
		req, _ = http.NewRequest(http.MethodGet, url, nil)
	} else {
		body, _ := io.ReadAll(c.Request.Body)
		req, _ = http.NewRequest(method, url, bytes.NewReader(body))
	}
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set(AuthorizationHeader, GetAccessToken(c))
	resp, err := Client.Do(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ReturnMessage(err.Error()))
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusUnauthorized {
			logger.Error(fmt.Sprintf(AccountDeactivatedErrorMessage, c.GetString(EmailKey)))
		}

		responseMap := make(map[string]interface{})
		json.NewDecoder(resp.Body).Decode(&responseMap)
		c.AbortWithStatusJSON(resp.StatusCode, responseMap)
		return
	}

	io.Copy(c.Writer, resp.Body)
}

func ReturnMessage(msg string) gin.H {
	logger.Warn(msg)

	return gin.H{
		defaultErrorMessageKey: msg,
	}
}

func GetAccessToken(c *gin.Context) string {
	accessToken := c.GetString(AuthorizationHeader)
	if !strings.HasPrefix(accessToken, "Bearer") {
		return "Bearer " + accessToken
	}

	return accessToken
}

func GetArkoseToken() (string, error) {
	return funcaptcha.GetOpenAIToken(PUID, ProxyUrl)
}

func setupPUID() {
	username := os.Getenv("OPENAI_EMAIL")
	password := os.Getenv("OPENAI_PASSWORD")
	if username != "" && password != "" {
		go func() {
			for {
				authenticator := auth.NewAuthenticator(username, password, ProxyUrl)
				if err := authenticator.Begin(); err != nil {
					logger.Warn(fmt.Sprintf("%s: %s", refreshPuidErrorMessage, err.Details))
					return
				}

				accessToken := authenticator.GetAccessToken()
				if accessToken == "" {
					logger.Error(refreshPuidErrorMessage)
					return
				}

				puid, err := authenticator.GetPUID()
				if err != nil {
					logger.Error(refreshPuidErrorMessage)
					return
				}

				PUID = puid

				time.Sleep(time.Hour * 24 * 7)
			}
		}()
	}
}
````

## File: api/imitate/api.go
````go
package imitate

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	http "github.com/bogdanfinn/fhttp"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/linweiyuan/go-chatgpt-api/api"
	"github.com/linweiyuan/go-chatgpt-api/api/chatgpt"
	"github.com/linweiyuan/go-logger/logger"
)

var (
	reg *regexp.Regexp
)

func init() {
	reg, _ = regexp.Compile("[^a-zA-Z0-9]+")
}

func CreateChatCompletions(c *gin.Context) {
	var originalRequest APIRequest
	err := c.BindJSON(&originalRequest)
	if err != nil {
		c.JSON(400, gin.H{"error": gin.H{
			"message": "Request must be proper JSON",
			"type":    "invalid_request_error",
			"param":   nil,
			"code":    err.Error(),
		}})
		return
	}

	authHeader := c.GetHeader(api.AuthorizationHeader)
	token := os.Getenv("IMITATE_ACCESS_TOKEN")
	if authHeader != "" {
		customAccessToken := strings.Replace(authHeader, "Bearer ", "", 1)
		// Check if customAccessToken starts with sk-
		if strings.HasPrefix(customAccessToken, "eyJhbGciOiJSUzI1NiI") {
			token = customAccessToken
		}
	}

	// 将聊天请求转换为ChatGPT请求。
	translatedRequest, model := convertAPIRequest(originalRequest)

	response, done := sendConversationRequest(c, translatedRequest, token)
	if done {
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	if HandleRequestError(c, response) {
		return
	}

	var fullResponse string

	id := generateId()

	for i := 3; i > 0; i-- {
		var continueInfo *ContinueInfo
		var responsePart string
		var continueSignal string
		responsePart, continueInfo = Handler(c, response, originalRequest.Stream, id, model)
		fullResponse += responsePart
		continueSignal = os.Getenv("CONTINUE_SIGNAL")
		if continueInfo == nil || continueSignal == "" {
			break
		}
		println("Continuing conversation")
		translatedRequest.Messages = nil
		translatedRequest.Action = "continue"
		translatedRequest.ConversationID = &continueInfo.ConversationID
		translatedRequest.ParentMessageID = continueInfo.ParentID
		response, done = sendConversationRequest(c, translatedRequest, token)

		if done {
			return
		}

		// 以下修复代码来自ChatGPT
		// 在循环内部创建一个局部作用域，并将资源的引用传递给匿名函数，保证资源将在每次迭代结束时被正确释放
		func() {
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					return
				}
			}(response.Body)
		}()

		if HandleRequestError(c, response) {
			return
		}
	}

	if !originalRequest.Stream {
		c.JSON(200, newChatCompletion(fullResponse, model, id))
	} else {
		c.String(200, "data: [DONE]\n\n")
	}
}

func generateId() string {
	id := uuid.NewString()
	id = strings.ReplaceAll(id, "-", "")
	id = base64.StdEncoding.EncodeToString([]byte(id))
	id = reg.ReplaceAllString(id, "")
	return "chatcmpl-" + id
}

func convertAPIRequest(apiRequest APIRequest) (chatgpt.CreateConversationRequest, string) {
	chatgptRequest := NewChatGPTRequest()

	var model = "gpt-3.5-turbo-0613"

	if strings.HasPrefix(apiRequest.Model, "gpt-3.5") {
		chatgptRequest.Model = "text-davinci-002-render-sha"
	}

	if strings.HasPrefix(apiRequest.Model, "gpt-4") {
		arkoseToken, err := api.GetArkoseToken()
		if err == nil {
			chatgptRequest.ArkoseToken = arkoseToken
		} else {
			fmt.Println("Error getting Arkose token: ", err)
		}
		chatgptRequest.Model = apiRequest.Model
		model = "gpt-4-0613"
	}

	if apiRequest.PluginIDs != nil {
		chatgptRequest.PluginIDs = apiRequest.PluginIDs
		chatgptRequest.Model = "gpt-4-plugins"
	}

	for _, apiMessage := range apiRequest.Messages {
		if apiMessage.Role == "system" {
			apiMessage.Role = "critic"
		}
		chatgptRequest.AddMessage(apiMessage.Role, apiMessage.Content)
	}

	return chatgptRequest, model
}

func NewChatGPTRequest() chatgpt.CreateConversationRequest {
	enableHistory := os.Getenv("ENABLE_HISTORY") == ""
	return chatgpt.CreateConversationRequest{
		Action:                     "next",
		ParentMessageID:            uuid.NewString(),
		Model:                      "text-davinci-002-render-sha",
		HistoryAndTrainingDisabled: !enableHistory,
	}
}

func sendConversationRequest(c *gin.Context, request chatgpt.CreateConversationRequest, accessToken string) (*http.Response, bool) {
	jsonBytes, _ := json.Marshal(request)
	req, _ := http.NewRequest(http.MethodPost, api.ChatGPTApiUrlPrefix+"/backend-api/conversation", bytes.NewBuffer(jsonBytes))
	req.Header.Set("User-Agent", api.UserAgent)
	req.Header.Set(api.AuthorizationHeader, accessToken)
	req.Header.Set("Accept", "text/event-stream")
	if api.PUID != "" {
		req.Header.Set("Cookie", "_puid="+api.PUID)
	}
	resp, err := api.Client.Do(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.ReturnMessage(err.Error()))
		return nil, true
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusUnauthorized {
			logger.Error(fmt.Sprintf(api.AccountDeactivatedErrorMessage, c.GetString(api.EmailKey)))
		}

		responseMap := make(map[string]interface{})
		json.NewDecoder(resp.Body).Decode(&responseMap)
		c.AbortWithStatusJSON(resp.StatusCode, responseMap)
		return nil, true
	}

	return resp, false
}

func Handler(c *gin.Context, response *http.Response, stream bool, id string, model string) (string, *ContinueInfo) {
	maxTokens := false

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
	var finishReason string
	var previousText StringStruct
	var originalResponse ChatGPTResponse
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

			err = json.Unmarshal([]byte(line), &originalResponse)
			if err != nil {
				continue
			}
			if originalResponse.Error != nil {
				c.JSON(500, gin.H{"error": originalResponse.Error})
				return "", nil
			}
			if originalResponse.Message.Author.Role != "assistant" || originalResponse.Message.Content.Parts == nil {
				continue
			}
			if originalResponse.Message.Metadata.MessageType != "next" && originalResponse.Message.Metadata.MessageType != "continue" || originalResponse.Message.EndTurn != nil {
				continue
			}
			if (len(originalResponse.Message.Content.Parts) == 0 || originalResponse.Message.Content.Parts[0] == "") && !isRole {
				continue
			}
			responseString := ConvertToString(&originalResponse, &previousText, isRole, id, model)
			isRole = false
			if stream {
				_, err = c.Writer.WriteString(responseString)
				if err != nil {
					return "", nil
				}
			}
			// Flush the response writer buffer to ensure that the client receives each line as it's written
			c.Writer.Flush()

			if originalResponse.Message.Metadata.FinishDetails != nil {
				if originalResponse.Message.Metadata.FinishDetails.Type == "max_tokens" {
					maxTokens = true
				}
				finishReason = originalResponse.Message.Metadata.FinishDetails.Type
			}

		} else {
			if stream {
				if finishReason == "" {
					finishReason = "stop"
				}
				finalLine := StopChunk(finishReason, id, model)
				_, err := c.Writer.WriteString("data: " + finalLine.String() + "\n\n")
				if err != nil {
					return "", nil
				}
			}
		}
	}
	if !maxTokens {
		return previousText.Text, nil
	}
	return previousText.Text, &ContinueInfo{
		ConversationID: originalResponse.ConversationID,
		ParentID:       originalResponse.Message.ID,
	}
}
````

## File: api/imitate/convert.go
````go
package imitate

import (
	"fmt"
	"strings"
)

func ConvertToString(chatgptResponse *ChatGPTResponse, previousText *StringStruct, role bool, id string, model string) string {
	var text string

	if len(chatgptResponse.Message.Content.Parts) == 1 {
		if part, ok := chatgptResponse.Message.Content.Parts[0].(string); ok {
			text = strings.ReplaceAll(part, previousText.Text, "")
			previousText.Text = part
		} else {
			text = fmt.Sprintf("%v", chatgptResponse.Message.Content.Parts[0])
		}
	} else {
		// When using GPT-4 messages with images (multimodal_text), the length of 'parts' might be 2.
		// Since the chatgpt API currently does not support multimodal content
		// and there is no official format for multimodal content,
		// the content is temporarily returned as is.
		var parts []string
		for _, part := range chatgptResponse.Message.Content.Parts {
			parts = append(parts, fmt.Sprintf("%v", part))
		}
		text = strings.Join(parts, ", ")
	}

	translatedResponse := NewChatCompletionChunk(text, id, model)
	if role {
		translatedResponse.Choices[0].Delta.Role = chatgptResponse.Message.Author.Role
	}

	return "data: " + translatedResponse.String() + "\n\n"
}
````

## File: api/imitate/request.go
````go
package imitate

import (
	"encoding/json"
	"io"

	http "github.com/bogdanfinn/fhttp"
	"github.com/gin-gonic/gin"
)

type ContinueInfo struct {
	ConversationID string `json:"conversation_id"`
	ParentID       string `json:"parent_id"`
}

type APIRequest struct {
	Messages  []ApiMessage `json:"messages"`
	Stream    bool         `json:"stream"`
	Model     string       `json:"model"`
	PluginIDs []string     `json:"plugin_ids"`
}

type ApiMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func HandleRequestError(c *gin.Context, response *http.Response) bool {
	if response.StatusCode != 200 {
		// Try read response body as JSON
		var errorResponse map[string]interface{}
		err := json.NewDecoder(response.Body).Decode(&errorResponse)
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
			"message": errorResponse["detail"],
			"type":    response.Status,
			"param":   nil,
			"code":    "error",
		}})
		return true
	}
	return false
}
````

## File: api/imitate/response.go
````go
package imitate

import (
	"encoding/json"
	"time"

	"github.com/linweiyuan/go-chatgpt-api/api/chatgpt"
)

type ChatCompletionChunk struct {
	ID      string    `json:"id"`
	Object  string    `json:"object"`
	Created int64     `json:"created"`
	Model   string    `json:"model"`
	Choices []Choices `json:"choices"`
}

func (chunk *ChatCompletionChunk) String() string {
	resp, _ := json.Marshal(chunk)
	return string(resp)
}

type Choices struct {
	Delta        Delta       `json:"delta"`
	Index        int         `json:"index"`
	FinishReason interface{} `json:"finish_reason"`
}

type Delta struct {
	Content string `json:"content,omitempty"`
	Role    string `json:"role,omitempty"`
}

func NewChatCompletionChunk(text string, id string, model string) ChatCompletionChunk {
	return ChatCompletionChunk{
		ID:      id,
		Object:  "chat.completion.chunk",
		Created: time.Now().Unix(),
		Model:   model,
		Choices: []Choices{
			{
				Index: 0,
				Delta: Delta{
					Content: text,
				},
				FinishReason: nil,
			},
		},
	}
}

func StopChunk(reason string, id string, model string) ChatCompletionChunk {
	return ChatCompletionChunk{
		ID:      id,
		Object:  "chat.completion.chunk",
		Created: time.Now().Unix(),
		Model:   model,
		Choices: []Choices{
			{
				Index:        0,
				FinishReason: reason,
			},
		},
	}
}

type ChatCompletion struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Usage   usage    `json:"usage"`
	Choices []Choice `json:"choices"`
}
type Msg struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type Choice struct {
	Index        int         `json:"index"`
	Message      Msg         `json:"message"`
	FinishReason interface{} `json:"finish_reason"`
}
type usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ChatGPTResponse struct {
	Message        Message     `json:"message"`
	ConversationID string      `json:"conversation_id"`
	Error          interface{} `json:"error"`
}

type Message struct {
	ID         string          `json:"id"`
	Author     chatgpt.Author  `json:"author"`
	CreateTime float64         `json:"create_time"`
	UpdateTime interface{}     `json:"update_time"`
	Content    chatgpt.Content `json:"content"`
	EndTurn    interface{}     `json:"end_turn"`
	Weight     float64         `json:"weight"`
	Metadata   Metadata        `json:"metadata"`
	Recipient  string          `json:"recipient"`
}

type Metadata struct {
	Timestamp     string         `json:"timestamp_"`
	MessageType   string         `json:"message_type"`
	FinishDetails *FinishDetails `json:"finish_details"`
	ModelSlug     string         `json:"model_slug"`
	Recipient     string         `json:"recipient"`
}

type FinishDetails struct {
	Type string `json:"type"`
	Stop string `json:"stop"`
}

type StringStruct struct {
	Text string `json:"text"`
}

func newChatCompletion(fullTest, model string, id string) ChatCompletion {
	return ChatCompletion{
		ID:      id,
		Object:  "chat.completion",
		Created: time.Now().Unix(),
		Model:   model,
		Usage: usage{
			PromptTokens:     0,
			CompletionTokens: 0,
			TotalTokens:      0,
		},
		Choices: []Choice{
			{
				Message: Msg{
					Content: fullTest,
					Role:    "assistant",
				},
				Index: 0,
			},
		},
	}
}
````

## File: api/platform/access_token.go
````go
package platform

import (
	"encoding/json"
	"errors"
	"io"
	"net/url"
	"strings"

	http "github.com/bogdanfinn/fhttp"

	"github.com/linweiyuan/go-chatgpt-api/api"
)

func (userLogin *UserLogin) GetAuthorizedUrl(csrfToken string) (string, int, error) {
	urlParams := url.Values{
		"client_id":     {platformAuthClientID},
		"audience":      {platformAuthAudience},
		"redirect_uri":  {platformAuthRedirectURL},
		"scope":         {platformAuthScope},
		"response_type": {platformAuthResponseType},
	}
	req, _ := http.NewRequest(http.MethodGet, platformAuth0Url+urlParams.Encode(), nil)
	req.Header.Set("Content-Type", api.ContentType)
	req.Header.Set("User-Agent", api.UserAgent)
	resp, err := userLogin.client.Do(req)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", resp.StatusCode, errors.New(api.GetAuthorizedUrlErrorMessage)
	}

	return resp.Request.URL.String(), http.StatusOK, nil
}

func (userLogin *UserLogin) GetState(authorizedUrl string) (string, int, error) {
	split := strings.Split(authorizedUrl, "=")
	return split[1], http.StatusOK, nil
}

func (userLogin *UserLogin) CheckUsername(state string, username string) (int, error) {
	formParams := url.Values{
		"state":                       {state},
		"username":                    {username},
		"js-available":                {"true"},
		"webauthn-available":          {"true"},
		"is-brave":                    {"false"},
		"webauthn-platform-available": {"false"},
		"action":                      {"default"},
	}
	req, _ := http.NewRequest(http.MethodPost, api.LoginUsernameUrl+state, strings.NewReader(formParams.Encode()))
	req.Header.Set("Content-Type", api.ContentType)
	req.Header.Set("User-Agent", api.UserAgent)
	resp, err := userLogin.client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, errors.New(api.EmailInvalidErrorMessage)
	}

	return http.StatusOK, nil
}

func (userLogin *UserLogin) CheckPassword(state string, username string, password string) (string, int, error) {
	formParams := url.Values{
		"state":    {state},
		"username": {username},
		"password": {password},
		"action":   {"default"},
	}
	req, _ := http.NewRequest(http.MethodPost, api.LoginPasswordUrl+state, strings.NewReader(formParams.Encode()))
	req.Header.Set("Content-Type", api.ContentType)
	req.Header.Set("User-Agent", api.UserAgent)
	resp, err := userLogin.client.Do(req)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", resp.StatusCode, errors.New(api.EmailOrPasswordInvalidErrorMessage)
	}

	return resp.Request.URL.Query().Get("code"), http.StatusOK, nil
}

func (userLogin *UserLogin) GetAccessToken(code string) (string, int, error) {
	jsonBytes, _ := json.Marshal(GetAccessTokenRequest{
		ClientID:    platformAuthClientID,
		Code:        code,
		GrantType:   platformAuthGrantType,
		RedirectURI: platformAuthRedirectURL,
	})
	req, _ := http.NewRequest(http.MethodPost, getTokenUrl, strings.NewReader(string(jsonBytes)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", api.UserAgent)
	resp, err := userLogin.client.Do(req)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", resp.StatusCode, errors.New(api.GetAccessTokenErrorMessage)
	}

	data, _ := io.ReadAll(resp.Body)
	return string(data), http.StatusOK, nil
}
````

## File: api/platform/api.go
````go
package platform

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	http "github.com/bogdanfinn/fhttp"
	"github.com/gin-gonic/gin"

	"github.com/linweiyuan/go-chatgpt-api/api"
	"github.com/linweiyuan/go-logger/logger"
)

func CreateChatCompletions(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	var request struct {
		Stream bool `json:"stream"`
	}
	json.Unmarshal(body, &request)

	url := c.Request.URL.Path
	if strings.Contains(url, "/chat") {
		url = apiCreateChatCompletions
	} else {
		url = apiCreateCompletions
	}

	resp, err := handlePost(c, url, body, request.Stream)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusUnauthorized {
		logger.Error(fmt.Sprintf(api.AccountDeactivatedErrorMessage, c.GetString(api.EmailKey)))
		responseMap := make(map[string]interface{})
		json.NewDecoder(resp.Body).Decode(&responseMap)
		c.AbortWithStatusJSON(resp.StatusCode, responseMap)
		return
	}

	if request.Stream {
		handleCompletionsResponse(c, resp)
	} else {
		io.Copy(c.Writer, resp.Body)
	}
}

func CreateCompletions(c *gin.Context) {
	CreateChatCompletions(c)
}

func handleCompletionsResponse(c *gin.Context, resp *http.Response) {
	c.Writer.Header().Set("Content-Type", "text/event-stream; charset=utf-8")

	reader := bufio.NewReader(resp.Body)
	for {
		if c.Request.Context().Err() != nil {
			break
		}

		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "event") ||
			strings.HasPrefix(line, "data: 20") ||
			line == "" {
			continue
		}

		c.Writer.Write([]byte(line + "\n\n"))
		c.Writer.Flush()
	}
}

func handlePost(c *gin.Context, url string, data []byte, stream bool) (*http.Response, error) {
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	req.Header.Set(api.AuthorizationHeader, api.GetAccessToken(c))
	if stream {
		req.Header.Set("Accept", "text/event-stream")
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := api.Client.Do(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.ReturnMessage(err.Error()))
		return nil, err
	}

	return resp, nil
}
````

## File: api/platform/constant.go
````go
package platform

import "github.com/linweiyuan/go-chatgpt-api/api"

const (
	apiCreateChatCompletions = api.PlatformApiUrlPrefix + "/v1/chat/completions"
	apiCreateCompletions     = api.PlatformApiUrlPrefix + "/v1/completions"

	platformAuthClientID      = "DRivsnm2Mu42T3KOpqdtwB3NYviHYzwD"
	platformAuthAudience      = "https://api.openai.com/v1"
	platformAuthRedirectURL   = "https://platform.openai.com/auth/callback"
	platformAuthScope         = "openid profile email offline_access"
	platformAuthResponseType  = "code"
	platformAuthGrantType     = "authorization_code"
	platformAuth0Url          = api.Auth0Url + "/authorize?"
	getTokenUrl               = api.Auth0Url + "/oauth/token"
	auth0Client               = "eyJuYW1lIjoiYXV0aDAtc3BhLWpzIiwidmVyc2lvbiI6IjEuMjEuMCJ9" // '{"name":"auth0-spa-js","version":"1.21.0"}'
	auth0LogoutUrl            = api.Auth0Url + "/v2/logout?returnTo=https%3A%2F%2Fplatform.openai.com%2Floggedout&client_id=" + platformAuthClientID + "&auth0Client=" + auth0Client
	dashboardLoginUrl         = "https://api.openai.com/dashboard/onboarding/login"
	getSessionKeyErrorMessage = "failed to get session key"
)
````

## File: api/platform/login.go
````go
package platform

import (
	"encoding/json"
	"io"
	"strings"

	http "github.com/bogdanfinn/fhttp"
	"github.com/gin-gonic/gin"

	"github.com/linweiyuan/go-chatgpt-api/api"
)

func Login(c *gin.Context) {
	var loginInfo api.LoginInfo
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.ReturnMessage(api.ParseUserInfoErrorMessage))
		return
	}

	userLogin := UserLogin{
		client: api.NewHttpClient(),
	}

	// hard refresh cookies
	resp, _ := userLogin.client.Get(auth0LogoutUrl)
	defer resp.Body.Close()

	// get authorized url
	authorizedUrl, statusCode, err := userLogin.GetAuthorizedUrl("")
	if err != nil {
		c.AbortWithStatusJSON(statusCode, api.ReturnMessage(err.Error()))
		return
	}

	// get state
	state, _, _ := userLogin.GetState(authorizedUrl)

	// check username
	statusCode, err = userLogin.CheckUsername(state, loginInfo.Username)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, api.ReturnMessage(err.Error()))
		return
	}

	// check password
	code, statusCode, err := userLogin.CheckPassword(state, loginInfo.Username, loginInfo.Password)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, api.ReturnMessage(err.Error()))
		return
	}

	// get access token
	accessToken, statusCode, err := userLogin.GetAccessToken(code)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, api.ReturnMessage(err.Error()))
		return
	}

	// get session key
	var getAccessTokenResponse GetAccessTokenResponse
	json.Unmarshal([]byte(accessToken), &getAccessTokenResponse)
	req, _ := http.NewRequest(http.MethodPost, dashboardLoginUrl, strings.NewReader("{}"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", api.UserAgent)
	req.Header.Set(api.AuthorizationHeader, "Bearer "+getAccessTokenResponse.AccessToken)
	resp, err = userLogin.client.Do(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.ReturnMessage(err.Error()))
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		c.AbortWithStatusJSON(resp.StatusCode, api.ReturnMessage(getSessionKeyErrorMessage))
		return
	}

	io.Copy(c.Writer, resp.Body)
}
````

## File: api/platform/typings.go
````go
package platform

import tls_client "github.com/bogdanfinn/tls-client"

type UserLogin struct {
	client tls_client.HttpClient
}

type GetAccessTokenRequest struct {
	ClientID    string `json:"client_id"`
	GrantType   string `json:"grant_type"`
	Code        string `json:"code"`
	RedirectURI string `json:"redirect_uri"`
}

type GetAccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}
````

## File: compose.yaml
````yaml
services:
  go-chatgpt-api:
    build: .
    container_name: go-chatgpt-api
    image: linweiyuan/go-chatgpt-api
    ports:
      - 8080:8080
    environment:
      - PORT=
      - TZ=Asia/Shanghai
      - PROXY=
      - OPENAI_EMAIL=
      - OPENAI_PASSWORD=
      - CONTINUE_SIGNAL=
      - ENABLE_HISTORY=
      - IMITATE_ACCESS_TOKEN=
    volumes:
      - ./chat.openai.com.har:/app/chat.openai.com.har
    restart: unless-stopped
````

## File: Dockerfile
````dockerfile
FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -ldflags="-w -s" -o go-chatgpt-api main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/go-chatgpt-api .
RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai
EXPOSE 8080
CMD ["/app/go-chatgpt-api"]
````

## File: env/env.go
````go
package env

import (
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}
````

## File: example/chatgpt.http
````
### login
POST {{baseUrl}}/chatgpt/login
Content-Type: application/json

{
  "username": "{{username}}",
  "password": "{{password}}"
}

### get conversations
GET {{baseUrl}}/chatgpt/backend-api/conversations?offset=0&limit=3&order=updated
Authorization: Bearer {{accessToken}}

### get conversation
GET {{baseUrl}}/chatgpt/backend-api/conversation/id
Authorization: Bearer {{accessToken}}

### create conversation
POST {{baseUrl}}/chatgpt/backend-api/conversation
Authorization: Bearer {{accessToken}}
Content-Type: application/json
Accept: text/event-stream

{
  "action": "next",
  "messages": [
    {
      "id": "{{$guid}}",
      "author": {
        "role": "user"
      },
      "content": {
        "content_type": "text",
        "parts": [
          "hello"
        ]
      },
      "metadata": {}
    }
  ],
  "model": "gpt-4",
  "timezone_offset_min": -480,
  "history_and_training_disabled": false
}

### get models
GET {{baseUrl}}/chatgpt/backend-api/models?history_and_training_disabled=false
Authorization: Bearer {{accessToken}}

### check account
GET {{baseUrl}}/chatgpt/backend-api/accounts/check
Authorization: Bearer {{accessToken}}

### check account v4
GET {{baseUrl}}/chatgpt/backend-api/accounts/check/v4-2023-04-27
Authorization: Bearer {{accessToken}}

### get settings beta features
GET {{baseUrl}}/chatgpt/backend-api/settings/beta_features
Authorization: Bearer {{accessToken}}

### get conversation limit (no need to pass access token)
GET {{baseUrl}}/chatgpt/public-api/conversation_limit
Authorization: Bearer {{accessToken}}

### get models with pandora enabled
GET {{baseUrl}}/api/models?history_and_training_disabled=false
Authorization: Bearer {{accessToken}}

### share link to chat
POST {{baseUrl}}/chatgpt/backend-api/share/create
Authorization: Bearer {{accessToken}}
Content-Type: application/json

{
	"current_node_id": "9020711b-3dcf-4705-82ac-46b5af30fc7b",
	"conversation_id": "74c406dd-a2e8-477a-b420-90ed57a55bf9",
	"is_anonymous": false
}

### copy link
PATCH {{baseUrl}}/chatgpt/backend-api/share/{share_id}
Authorization: Bearer {{accessToken}}
Content-Type: application/json

{
	"share_id": "49cd2432-d084-4ab7-8549-4ee18046812b",
	"highlighted_message_id": null,
	"title": "Summarize Request and Response 11122",
	"is_public": false,
	"is_visible": false,
	"is_anonymous": true
}

### continue shared conversation
POST {{baseUrl}}/chatgpt/backend-api/conversation
Authorization: Bearer {{accessToken}}
Content-Type: application/json

{
	"action": "next",
	"messages": [{
		"id": "{{$guid}}",
		"author": {
			"role": "user"
		},
		"content": {
			"content_type": "text",
			"parts": [
        "hello again"
      ]
		},
		"metadata": {}
	}],
	"continue_from_shared_conversation_id": "this is the share_id",
	"parent_message_id": "this is the current_node_id",
	"model": "text-davinci-002-render-sha",
	"timezone_offset_min": -480,
	"history_and_training_disabled": false,
	"arkose_token": null
}

### get plugins
GET {{baseUrl}}/chatgpt/backend-api/aip/p?offset=0&limit=250&statuses=approved
Authorization: Bearer {{accessToken}}

### get payment url
GET {{baseUrl}}/chatgpt/backend-api/payments/customer_portal
Authorization: Bearer {{accessToken}}
````

## File: example/imitate.http
````
### Create chat completion
POST {{baseUrl}}/imitate/v1/chat/completions
Content-Type: application/json
Authorization: Bearer {{accessToken}}

{
  "model": "gpt-3.5-turbo",
  "messages": [
    {
      "role": "system",
      "content": "You are a helpful assistant."
    },
    {
      "role": "user",
      "content": "Hello!"
    }
  ],
  "stream": true
}
````

## File: example/ios.http
````
### create conversation
POST https://ios.chat.openai.com/backend-api/conversation
Authorization: Bearer {{accessToken}}
Cookie: _devicecheck=user-xxx
Accept: text/event-stream
Content-Type: application/json

{
  "action": "next",
  "messages": [
    {
      "id": "{{$guid}}",
      "author": {
        "role": "user"
      },
      "content": {
        "content_type": "text",
        "parts": [
          "hello"
        ]
      },
      "metadata": {}
    }
  ],
  "model": "gpt-4",
  "timezone_offset_min": -480,
  "history_and_training_disabled": false,
  "supports_modapi": true
}

### get conversations
GET https://ios.chat.openai.com/backend-api/conversations?offset=0&limit=3&order=updated
Authorization: Bearer {{accessToken}}
Cookie: _devicecheck=user-xxx

### device check
POST https://ios.chat.openai.com/backend-api/devicecheck
Authorization: Bearer {{accessToken}}
Cookie: _preauth_devicecheck=xxx
Content-Type: application/json

{
	"bundle_id": "com.openai.chat",
	"device_token": "ad"
}

### me
GET https://ios.chat.openai.com/backend-api/me?include_groups=true
Authorization: Bearer {{accessToken}}
Cookie: _devicecheck=user-xxx
````

## File: example/platform.http
````
### login
POST {{baseUrl}}/platform/login
Content-Type: application/json

{
  "username": "{{username}}",
  "password": "{{password}}"
}

### get models
GET {{baseUrl}}/platform/v1/models
Authorization: Bearer {{apiKey}}

### get model
GET {{baseUrl}}/platform/v1/models/gpt-3.5-turbo-16k-0613
Authorization: Bearer {{apiKey}}

### Create chat completion
POST {{baseUrl}}/platform/v1/chat/completions
Content-Type: application/json
Authorization: Bearer {{apiKey}}

{
  "model": "gpt-3.5-turbo",
  "messages": [
    {
      "role": "system",
      "content": "You are a helpful assistant."
    },
    {
      "role": "user",
      "content": "Hello!"
    }
  ],
  "stream": true
}

### Create completion
POST {{baseUrl}}/platform/v1/completions
Content-Type: application/json
Authorization: Bearer {{apiKey}}

{
  "model": "text-davinci-003",
  "prompt": "Say this is a test",
  "max_tokens": 7,
  "temperature": 0,
  "stream": true
}

### get user api_keys
GET {{baseUrl}}/platform/dashboard/user/api_keys
Authorization: Bearer {{apiKey}}

### get billing credit_grants
GET {{baseUrl}}/platform/dashboard/billing/credit_grants
Authorization: Bearer {{apiKey}}

### get billing subscription
GET {{baseUrl}}/platform/dashboard/billing/subscription
Authorization: Bearer {{apiKey}}

### get billing usage
GET {{baseUrl}}/platform/dashboard/billing/usage?end_date=2023-07-01&start_date=2023-06-01
Authorization: Bearer {{apiKey}}
````

## File: go.mod
````
module github.com/linweiyuan/go-chatgpt-api

go 1.21

require (
	github.com/PuerkitoBio/goquery v1.8.1
	github.com/bogdanfinn/fhttp v0.5.24
	github.com/bogdanfinn/tls-client v1.6.1
	github.com/gin-gonic/gin v1.9.1
	github.com/google/uuid v1.3.1
	github.com/joho/godotenv v1.5.1
	github.com/linweiyuan/go-logger v0.0.0-20230709142852-da1f090a7d4c
	github.com/xqdoo00o/OpenAIAuth v0.0.0-20230928031215-356afd0d7a6b
	github.com/xqdoo00o/funcaptcha v0.0.0-20230928030317-87dbaf7079cf
)

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/andybalholm/cascadia v1.3.2 // indirect
	github.com/bogdanfinn/utls v1.5.16 // indirect
	github.com/bytedance/sonic v1.9.1 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.14.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/tam7t/hpkp v0.0.0-20160821193359-2b70b4024ed5 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
````

## File: LICENSE
````
MIT License

Copyright (c) 2023 linweiyuan

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

## File: main.go
````go
package main

import (
	"log"
	"os"
	"strings"

	http "github.com/bogdanfinn/fhttp"
	"github.com/gin-gonic/gin"

	"github.com/linweiyuan/go-chatgpt-api/api"
	"github.com/linweiyuan/go-chatgpt-api/api/chatgpt"
	"github.com/linweiyuan/go-chatgpt-api/api/imitate"
	"github.com/linweiyuan/go-chatgpt-api/api/platform"
	_ "github.com/linweiyuan/go-chatgpt-api/env"
	"github.com/linweiyuan/go-chatgpt-api/middleware"
)

func init() {
	gin.ForceConsoleColor()
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	router := gin.Default()

	router.Use(middleware.CORS())
	router.Use(middleware.Authorization())

	setupChatGPTAPIs(router)
	setupPlatformAPIs(router)
	setupPandoraAPIs(router)
	setupImitateAPIs(router)
	router.NoRoute(api.Proxy)

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, api.ReadyHint)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal("failed to start server: " + err.Error())
	}
}

func setupChatGPTAPIs(router *gin.Engine) {
	chatgptGroup := router.Group("/chatgpt")
	{
		chatgptGroup.POST("/login", chatgpt.Login)
		chatgptGroup.POST("/backend-api/login", chatgpt.Login) // add support for other projects

		conversationGroup := chatgptGroup.Group("/backend-api/conversation")
		{
			conversationGroup.POST("", chatgpt.CreateConversation)
		}
	}
}

func setupPlatformAPIs(router *gin.Engine) {
	platformGroup := router.Group("/platform")
	{
		platformGroup.POST("/login", platform.Login)
		platformGroup.POST("/v1/login", platform.Login)

		apiGroup := platformGroup.Group("/v1")
		{
			apiGroup.POST("/chat/completions", platform.CreateChatCompletions)
			apiGroup.POST("/completions", platform.CreateCompletions)
		}
	}
}

func setupPandoraAPIs(router *gin.Engine) {
	router.Any("/api/*path", func(c *gin.Context) {
		c.Request.URL.Path = strings.ReplaceAll(c.Request.URL.Path, "/api", "/chatgpt/backend-api")
		router.HandleContext(c)
	})
}

func setupImitateAPIs(router *gin.Engine) {
	imitateGroup := router.Group("/imitate")
	{
		imitateGroup.POST("/login", chatgpt.Login)

		apiGroup := imitateGroup.Group("/v1")
		{
			apiGroup.POST("/chat/completions", imitate.CreateChatCompletions)
		}
	}
}
````

## File: middleware/authorization.go
````go
package middleware

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/linweiyuan/go-chatgpt-api/api"
)

const (
	emptyAccessTokenErrorMessage      = "please provide a valid access token or api key in 'Authorization' header"
	accessTokenHasExpiredErrorMessage = "the accessToken for account %s has expired"
)

type AccessToken struct {
	HTTPSAPIOpenaiComProfile struct {
		Email         string `json:"email"`
		EmailVerified bool   `json:"email_verified"`
	} `json:"https://api.openai.com/profile"`
	HTTPSAPIOpenaiComAuth struct {
		UserID string `json:"user_id"`
	} `json:"https://api.openai.com/auth"`
	Iss   string   `json:"iss"`
	Sub   string   `json:"sub"`
	Aud   []string `json:"aud"`
	Iat   int      `json:"iat"`
	Exp   int      `json:"exp"`
	Azp   string   `json:"azp"`
	Scope string   `json:"scope"`
}

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader(api.AuthorizationHeader)
		if authorization == "" {
			authorization = c.GetHeader(api.XAuthorizationHeader)
		}

		if authorization == "" {
			if c.Request.URL.Path == "/" {
				c.Header("Content-Type", "text/plain")
			} else if strings.HasSuffix(c.Request.URL.Path, "/login") ||
				strings.HasPrefix(c.Request.URL.Path, "/chatgpt/public-api") ||
				(strings.HasPrefix(c.Request.URL.Path, "/imitate") && os.Getenv("IMITATE_ACCESS_TOKEN") != "") {
				c.Header("Content-Type", "application/json")
			} else if c.Request.URL.Path == "/favicon.ico" {
				c.Abort()
				return
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, api.ReturnMessage(emptyAccessTokenErrorMessage))
				return
			}

			c.Next()
		} else {
			if expired := isExpired(c); expired {
				c.AbortWithStatusJSON(http.StatusUnauthorized, api.ReturnMessage(fmt.Sprintf(accessTokenHasExpiredErrorMessage, c.GetString(api.EmailKey))))
				return
			}

			c.Set(api.AuthorizationHeader, authorization)
		}
	}
}

func isExpired(c *gin.Context) bool {
	accessToken := c.GetHeader(api.AuthorizationHeader)
	split := strings.Split(accessToken, ".")
	if len(split) == 3 {
		rawDecodedText, _ := base64.RawStdEncoding.DecodeString(split[1])
		var accessToken AccessToken
		json.Unmarshal(rawDecodedText, &accessToken)

		c.Set(api.EmailKey, accessToken.HTTPSAPIOpenaiComProfile.Email)

		exp := int64(accessToken.Exp)
		expTime := time.Unix(exp, 0)
		now := time.Now()

		return now.After(expTime)
	}

	// apiKey
	return false
}
````

## File: middleware/cors.go
````go
package middleware

import (
	http "github.com/bogdanfinn/fhttp"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
````

## File: README.md
````markdown
# go-chatgpt-api

<details>

<summary>这个项目已经死了，请让它安息吧</summary>

## 一个尝试绕过 `Cloudflare` 来使用 `ChatGPT` 接口的程序

---

# 以下文档内容已过期，不一定支持（2023-10-24）

### 支持接口

- https://chat.openai.com/auth/login 登录返回 `accessToken`（谷歌和微软账号暂不支持登录，但可正常使用其他接口）
- 模型和插件查询
- `GPT-3.5` 和 `GPT-4` 对话增删改查及分享
- https://platform.openai.com/playground 登录返回 `apiKey`
- `apiKey` 余额查询
- 等等 ...
- 支持 `ChatGPT` 转 `API`，接口 `/imitate/v1/chat/completions`，利用 `accessToken` 模拟 `apiKey`，实现伪免费使用 `API`，从而支持集成仅支持 `apiKey` 调用的第三方客户端项目，分享一个好用的脚本测试 `web-to-api` (https://github.com/linweiyuan/go-chatgpt-api/issues/251)

```python
import openai

openai.api_key = "这里填 access token，不是 api key"
openai.api_base = "http://127.0.0.1:8080/imitate/v1"

while True:
    text = input("请输入问题：")
    response = openai.ChatCompletion.create(
        model='gpt-3.5-turbo',
        messages=[
            {'role': 'user', 'content': text},
        ],
        stream=True
    )

    for chunk in response:
        print(chunk.choices[0].delta.get("content", ""), end="", flush=True)
    print("\n")
```

范例（URL 和参数基本保持着和官网一致，部分接口有些许改动），部分例子，不是全部，**理论上**全部基于文本传输的接口都支持

https://github.com/linweiyuan/go-chatgpt-api/tree/main/example

---

### 使用的过程中遇到问题应该如何解决

汇总贴：https://github.com/linweiyuan/go-chatgpt-api/issues/74

如果有疑问而不是什么程序出错其实可以在 [Discussions](https://github.com/linweiyuan/go-chatgpt-api/discussions) 里发而不是新增 Issue

群聊：https://github.com/linweiyuan/go-chatgpt-api/discussions/197

再说一遍，不要来 `Issues` 提你的疑问（再提不回复直接关闭），有讨论区，有群，不要提脑残问题，反面教材：https://github.com/linweiyuan/go-chatgpt-api/issues/255

---

### 配置

如需设置代理，可以设置环境变量 `PROXY`，比如 `PROXY=http://127.0.0.1:20171` 或者 `PROXY=socks5://127.0.0.1:20170`，注释掉或者留空则不启用

如果代理需账号密码验证，则 `http://username:password@ip:port` 或者 `socks5://username:password@ip:port`

如需配合 `warp` 使用：`PROXY=socks5://chatgpt-proxy-server-warp:65535`，因为需要设置 `warp` 的场景已经默认可以直接访问 `ChatGPT` 官网，因此共用一个变量不冲突（国内 `VPS` 不在讨论范围内，请自行配置网络环境，`warp` 服务在魔法环境下才能正常工作）

家庭网络无需跑 `warp` 服务，跑了也没用，会报错，仅在服务器需要

`CONTINUE_SIGNAL=1`，开启 `/imitate` 接口自动继续会话功能，留空关闭，默认关闭

---

`GPT-4` 相关模型目前需要验证 `arkose_token`，社区已经有很多解决方案，请自行查找，其中一个能用的：https://github.com/linweiyuan/go-chatgpt-api/issues/252

参考配置视频（拉到文章最下面点开视频，需要自己有一定的动手能力，根据你的环境不同自行微调配置）：[如何生成 GPT-4 arkose_token](https://linweiyuan.github.io/2023/06/24/%E5%A6%82%E4%BD%95%E7%94%9F%E6%88%90-GPT-4-arkose-token.html)

---

根据你的网络环境不同，可以展开查看对应配置，下面例子是基本参数，更多参数查看 [compose.yaml](https://github.com/linweiyuan/go-chatgpt-api/blob/main/compose.yaml)

<details>

<summary>直接利用现成的服务</summary>

服务器不定时维护，不保证高可用，利用这些服务导致的账号安全问题，与本项目无关

- https://go-chatgpt-api.linweiyuan.com

</details>

<details>

<summary>网络在直连或者通过代理的情况下可以正常访问 ChatGPT</summary>

```yaml
services:
  go-chatgpt-api:
    container_name: go-chatgpt-api
    image: linweiyuan/go-chatgpt-api
    ports:
      - 8080:8080
    environment:
      - TZ=Asia/Shanghai
    restart: unless-stopped
```

</details>

<details>

<summary>服务器无法正常访问 ChatGPT</summary>

```yaml
services:
  go-chatgpt-api:
    container_name: go-chatgpt-api
    image: linweiyuan/go-chatgpt-api
    ports:
      - 8080:8080
    environment:
      - TZ=Asia/Shanghai
      - PROXY=socks5://chatgpt-proxy-server-warp:65535
    depends_on:
      - chatgpt-proxy-server-warp
    restart: unless-stopped

  chatgpt-proxy-server-warp:
    container_name: chatgpt-proxy-server-warp
    image: linweiyuan/chatgpt-proxy-server-warp
    restart: unless-stopped
```

</details>

---

目前 `warp` 容器检测到流量超过 1G 会自动重启，如果你知道什么是 `teams-enroll-token` （不知道就跳过），可以通过环境变量 `TEAMS_ENROLL_TOKEN` 设置它的值，然后利用这条命令来检查是否生效

`docker-compose exec chatgpt-proxy-server-warp warp-cli --accept-tos account | awk 'NR==1'`

```
Account type: Free （没有生效）

Account type: Team （设置正常）
```

---

### Render 部署

点击下面的按钮一键部署，缺点是免费版本冷启动比较慢

[![Deploy to Render](https://render.com/images/deploy-to-render-button.svg)](https://render.com/deploy?repo=https://github.com/linweiyuan/go-chatgpt-api)

---

### 如何集成其他第三方客户端（下面的内容不一定是最新，有问题请去各自项目查看）

- [moeakwak/chatgpt-web-share](https://github.com/moeakwak/chatgpt-web-share)

环境变量

```
CHATGPT_BASE_URL=http://go-chatgpt-api:8080/chatgpt/backend-api/
```

- [lss233/chatgpt-mirai-qq-bot](https://github.com/lss233/chatgpt-mirai-qq-bot)

`config.cfg`

```
[openai]
browserless_endpoint = "http://go-chatgpt-api:8080/chatgpt/backend-api/"
```

- [Kerwin1202/chatgpt-web](https://github.com/Kerwin1202/chatgpt-web) | [Chanzhaoyu/chatgpt-web](https://github.com/Chanzhaoyu/chatgpt-web)

环境变量

```
API_REVERSE_PROXY=http://go-chatgpt-api:8080/chatgpt/backend-api/conversation
```

- [pengzhile/pandora](https://github.com/pengzhile/pandora)（不完全兼容）

环境变量

```
CHATGPT_API_PREFIX=http://go-chatgpt-api:8080
```

---

- [1130600015/feishu-chatgpt](https://github.com/1130600015/feishu-chatgpt)

`application.yaml`

```yaml
proxy:
  url: http://go-chatgpt-api:8080
```

---

- [Yidadaa/ChatGPT-Next-Web](https://github.com/Yidadaa/ChatGPT-Next-Web)

环境变量

```
BASE_URL=http://go-chatgpt-api:8080/imitate
```

---

### 相关博客（程序更新很多次，文章的内容可能和现在的不一样，仅供参考）：[ChatGPT](https://linweiyuan.github.io/categories/ChatGPT/)

- [如何生成 GPT-4 arkose_token](https://linweiyuan.github.io/2023/06/24/%E5%A6%82%E4%BD%95%E7%94%9F%E6%88%90-GPT-4-arkose-token.html)
- [利用 HTTP Client 来调试 go-chatgpt-api](https://linweiyuan.github.io/2023/06/18/%E5%88%A9%E7%94%A8-HTTP-Client-%E6%9D%A5%E8%B0%83%E8%AF%95-go-chatgpt-api.html)
- [一种解决 ChatGPT Access denied 的方法](https://linweiyuan.github.io/2023/04/15/%E4%B8%80%E7%A7%8D%E8%A7%A3%E5%86%B3-ChatGPT-Access-denied-%E7%9A%84%E6%96%B9%E6%B3%95.html)
- [ChatGPT 如何自建代理](https://linweiyuan.github.io/2023/04/08/ChatGPT-%E5%A6%82%E4%BD%95%E8%87%AA%E5%BB%BA%E4%BB%A3%E7%90%86.html)
- [一种取巧的方式绕过 Cloudflare v2 验证](https://linweiyuan.github.io/2023/03/14/%E4%B8%80%E7%A7%8D%E5%8F%96%E5%B7%A7%E7%9A%84%E6%96%B9%E5%BC%8F%E7%BB%95%E8%BF%87-Cloudflare-v2-%E9%AA%8C%E8%AF%81.html)

---

### 最后感谢各位同学

<a href="https://github.com/linweiyuan/go-chatgpt-api/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=linweiyuan/go-chatgpt-api"  alt=""/>
</a>

</details>
````

## File: render.yaml
````yaml
services:
  - type: web
    name: go-chatgpt-api
    runtime: go
    plan: free
    buildCommand: go build -ldflags="-w -s" -o go-chatgpt-api main.go
    startCommand: ./go-chatgpt-api
````
