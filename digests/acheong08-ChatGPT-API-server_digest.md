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
.github/workflows/docker-build.yaml
.github/workflows/release.yml
.gitignore
Data/placeholder
docker-compose.yml
Dockerfile
go.mod
handlers/admin.go
handlers/api.go
handlers/client.go
handlers/init.go
LICENSE
main.go
README.md
types/types.go
utils/auth.go
utils/utils.go
</directory_structure>

<files>
This section contains the contents of the repository's files.

<file path=".github/workflows/docker-build.yaml">
name: docker build

on:
  push:
    branches:
      - "**"
    tags:
      - "**"
  pull_request_target:

permissions:
  contents: read
  packages: write

jobs:
  docker:
    name: Publish Docker image
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v2
      - name: Docker meta
        id: docker_meta
        uses: docker/metadata-action@v3.2.0
        with:
          images: ghcr.io/${{ github.repository }}
          flavor: |
            latest=false
          tags: |
            type=ref,event=pr
            type=ref,event=branch
            type=sha,prefix=,format=long,event=branch
            type=ref,event=tag
            type=sha,prefix=,format=long,event=tag
            type=raw,value=latest,enable=${{ endsWith(github.ref, github.event.repository.default_branch) }}
      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v1
      - name: Login to GitHub Container Registry
        uses: docker/login-action@v1
        with:
          registry: ghcr.io
          username: ${{ github.repository_owner }}
          password: ${{ secrets.GITHUB_TOKEN }}
      - name: Cache Docker layers
        uses: actions/cache@v2.1.5
        with:
          path: /tmp/.buildx-cache
          key: ${{ runner.os }}-buildx-${{ github.sha }}
          restore-keys: |
            ${{ runner.os }}-buildx-
      - name: Build and push
        uses: docker/build-push-action@v2
        with:
          context: .
          push: true
          tags: ${{ steps.docker_meta.outputs.tags }}
          labels: ${{ steps.docker_meta.outputs.labels }}
          cache-from: type=local,src=/tmp/.buildx-cache
          cache-to: type=local,dest=/tmp/.buildx-cache,mode=max
</file>

<file path=".github/workflows/release.yml">
name: Release Go Binaries

on: 
  release:
    types: [created]
  workflow_dispatch:
env:
  CMD_PATH: ./


jobs:
  releases-matrix:
    name: Release Matrix
    runs-on: ubuntu-latest
    strategy:
      matrix:
        goos: [linux, windows, darwin]
        goarch: ["386", amd64,arm64]
        exclude:
          - goarch: "386"
            goos: darwin 
    steps:
    - uses: actions/checkout@v2

    - name: Set APP_VERSION env
      run: echo APP_VERSION=$(echo ${GITHUB_REF} | rev | cut -d'/' -f 1 | rev ) >> ${GITHUB_ENV}
    - name: Set BUILD_TIME env
      run: echo BUILD_TIME=$(date) >> ${GITHUB_ENV}
    - name: Environment Printer
      uses: managedkaos/print-env@v1.0

    - uses: wangyoucao577/go-release-action@v1.34
      with:
        github_token: ${{ secrets.ACTIONS_TOKEN }}
        goos: ${{ matrix.goos }}
        goarch: ${{ matrix.goarch }}
        project_path: "${{ env.CMD_PATH }}"
        build_flags: -v
        ldflags: -X "main.appVersion=${{ env.APP_VERSION }}" -X "main.buildTime=${{ env.BUILD_TIME }}" -X main.gitCommit=${{ github.sha }} -X main.gitRef=${{ github.ref }}
</file>

<file path="Data/placeholder">

</file>

<file path="docker-compose.yml">
version: "3"

services:
  chatgpt-api-server:
    build: .
    ports:
      - "8080:8080"
    command: ["ChatGPT-API-server", "8080", "<API_KEY>", "-listen", "0.0.0.0"]
    networks:
      - chatgpt-api-server_default
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8080/health"]
      interval: 30s
      timeout: 10s
      retries: 3

networks:
  chatgpt-api-server_default:
</file>

<file path="Dockerfile">
FROM golang:1.19-alpine AS build

RUN apk add --no-cache git
RUN apk add gcc build-base

RUN mkdir -p /go/src/ChatGPT-API-server
WORKDIR /go/src/ChatGPT-API-server

RUN git clone https://github.com/ChatGPT-Hackers/ChatGPT-API-server/ .
RUN go install .

FROM alpine:latest
COPY --from=build /go/bin/ChatGPT-API-server /usr/local/bin/

RUN apk add --no-cache curl
</file>

<file path="go.mod">
module github.com/ChatGPT-Hackers/ChatGPT-API-server

go 1.19

require (
	github.com/gin-gonic/gin v1.8.1
	github.com/google/uuid v1.3.0
	github.com/gorilla/websocket v1.5.0
	github.com/mattn/go-sqlite3 v1.14.16
)

require (
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.0 // indirect
	github.com/goccy/go-json v0.9.7 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.1 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/sys v0.0.0-20210806184541-e5e7981a1069 // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
</file>

<file path="handlers/admin.go">
package handlers

import (
	"os"

	"github.com/ChatGPT-Hackers/ChatGPT-API-server/utils"
	_ "github.com/ChatGPT-Hackers/ChatGPT-API-server/utils"
	"github.com/gin-gonic/gin"
)

type Request struct {
	AdminKey string `json:"admin_key"`
	UserID   string `json:"user_id"`
}

func Admin_userAdd(c *gin.Context) {
	// Get admin key from request body
	var request Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// Check if admin key is valid
	if !utils.VerifyAdminKey(request.AdminKey) {
		c.JSON(401, gin.H{
			"error": "Invalid admin key",
		})
		return
	}

	// Generate user_id and token
	user_id := utils.GenerateId()
	token := utils.GenerateId()

	// Insert user_id and token into database
	err := utils.DatabaseInsert(user_id, token)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to insert user_id and token into database",
		})
		return
	}

	// Return user_id and token
	c.JSON(200, gin.H{
		"user_id": user_id,
		"token":   token,
	})
}

// POST request to delete a user
func Admin_userDel(c *gin.Context) {
	// Get admin key from request body
	var request Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// Check if admin key is valid
	if !utils.VerifyAdminKey(request.AdminKey) {
		c.JSON(401, gin.H{
			"error": "Invalid admin key",
		})
		return
	}

	// Delete user from database
	err := utils.DatabaseDelete(request.UserID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to delete user from database",
		})
		return
	}

	// Return success
	c.JSON(200, gin.H{
		"message": "User deleted",
	})
}

func Admin_usersGet(c *gin.Context) {
	// Get admin key from GET parameter
	AdminKey := c.Query("admin_key")

	// Check if admin key is valid
	if !utils.VerifyAdminKey(AdminKey) {
		c.JSON(401, gin.H{
			"error":   "Invalid admin key",
			"key":     AdminKey,
			"correct": os.Args[2],
		})
		return
	}

	// Get users from database
	users, err := utils.DatabaseSelectAll()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to get users from database",
			"error":   err.Error(),
		})
		return
	}

	// Return users
	c.JSON(200, gin.H{
		"users": users,
	})
}
</file>

<file path="handlers/api.go">
package handlers

import (
	"encoding/json"
	"time"

	"github.com/ChatGPT-Hackers/ChatGPT-API-server/types"
	"github.com/ChatGPT-Hackers/ChatGPT-API-server/utils"
	"github.com/gin-gonic/gin"
)

// // # API routes
func API_ask(c *gin.Context) {
	// Get request
	var request types.ChatGptRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request",
		})
		return
	}
	// Check if "Authorization" in Header
	if c.Request.Header["Authorization"] == nil {
		c.JSON(401, gin.H{
			"error": "API key not provided",
		})
		return
	}
	// Check if API key is valid
	verified, err := utils.VerifyToken(c.Request.Header["Authorization"][0])
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to verify API key",
		})
		return
	}
	if !verified {
		c.JSON(401, gin.H{
			"error": "Invalid API key",
		})
		return
	}
	// If Id is not set, generate a new one
	if request.MessageId == "" {
		request.MessageId = utils.GenerateId()
	}
	// If parent id is not set, generate a new one
	if request.ParentId == "" {
		request.ParentId = utils.GenerateId()
	}
	jsonRequest, err := json.Marshal(request)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to convert request to json",
		})
		return
	}
	var connection *types.Connection
	// Check conversation id
	connectionPool.Mu.RLock()
	// Check number of connections
	if len(connectionPool.Connections) == 0 {
		c.JSON(503, gin.H{
			"error": "No available clients",
		})
		return
	}
	connectionPool.Mu.RUnlock()
	if request.ConversationId == "" {
		// Retry 3 times before giving up
		var succeeded bool = false
		for i := 0; i < 3; i++ {
			// Find connection with the lowest load and where heartbeat is after last message time
			connectionPool.Mu.RLock()
			for _, conn := range connectionPool.Connections {
				if connection == nil || conn.LastMessageTime.Before(connection.LastMessageTime) {
					if conn.Heartbeat.After(conn.LastMessageTime) {
						connection = conn
					}
				}
			}
			connectionPool.Mu.RUnlock()
			// Check if connection was found
			if connection == nil {
				c.JSON(503, gin.H{
					"error": "No available clients",
				})
				return
			}
			// Ping before sending request
			var pingSucceeded bool = ping(connection.Id)
			if !pingSucceeded {
				// Ping failed. Try again
				connectionPool.Delete(connection.Id)
				succeeded = false
				connection = nil
				continue
			} else {
				succeeded = true
				break
			}
		}
		if !succeeded {
			// Delete connection
			c.JSON(503, gin.H{
				"error": "Ping failed",
			})
			return
		}
	} else {
		// Check if conversation exists
		conversation, ok := conversationPool.Get(request.ConversationId)
		if !ok {
			// Error
			c.JSON(500, gin.H{
				"error": "Conversation doesn't exists",
			})
			return
		} else {
			// Get connectionId of the conversation
			connectionId := conversation.ConnectionId
			// Check if connection exists
			connection, ok = connectionPool.Get(connectionId)
			if !ok {
				// Error
				c.JSON(500, gin.H{
					"error": "Connection no longer exists",
				})
				return
			}
		}
		// Ping before sending request
		if !ping(connection.Id) {
			c.JSON(503, gin.H{
				"error": "Ping failed",
			})
			return
		}
	}
	message := types.Message{
		Id:      utils.GenerateId(),
		Message: "ChatGptRequest",
		// Convert request to json
		Data: string(jsonRequest),
	}
	err = connection.Ws.WriteJSON(message)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to send request to the client",
		})
		// Delete connection
		connectionPool.Delete(connection.Id)
		return
	}
	// Set last message time
	connection.LastMessageTime = time.Now()
	// Wait for response with a timeout
	for {
		// Read message
		var receive types.Message
		connection.Ws.SetReadDeadline(time.Now().Add(120 * time.Second))
		err = connection.Ws.ReadJSON(&receive)
		if err != nil {
			c.JSON(500, gin.H{
				"error": "Failed to read response from the client",
				"err":   err.Error(),
			})
			// Delete connection
			connectionPool.Delete(connection.Id)
			return
		}
		// Check if the message is the response
		if receive.Id == message.Id {
			// Convert response to ChatGptResponse
			var response types.ChatGptResponse
			err = json.Unmarshal([]byte(receive.Data), &response)
			if err != nil {
				c.JSON(500, gin.H{
					"error":    "Failed to convert response to ChatGptResponse",
					"response": receive,
				})
				return
			}
			// Add conversation to pool
			conversation := &types.Conversation{
				Id:           response.ConversationId,
				ConnectionId: connection.Id,
			}
			conversationPool.Set(conversation)
			// Send response
			c.JSON(200, response)
			// Heartbeat
			connection.Heartbeat = time.Now()
			return
		} else {
			// Error
			c.JSON(500, gin.H{
				"error": "Failed to find response from the client",
			})
			return
		}

	}

}

func API_getConnections(c *gin.Context) {
	// Get connections
	var connections []*types.Connection
	connectionPool.Mu.RLock()
	for _, connection := range connectionPool.Connections {
		connections = append(connections, connection)
	}
	connectionPool.Mu.RUnlock()
	// Send connections
	c.JSON(200, gin.H{
		"connections": connections,
	})
}

func ping(connection_id string) bool {
	// Get connection
	connection, ok := connectionPool.Get(connection_id)
	// Send "ping" to the connection
	if ok {
		id := utils.GenerateId()
		send := types.Message{
			Id:      id,
			Message: "ping",
		}
		connection.Ws.SetReadDeadline(time.Now().Add(5 * time.Second))
		err := connection.Ws.WriteJSON(send)
		if err != nil {
			return false
		}
		// Wait for response with a timeout
		for {
			// Read message
			var receive types.Message
			err = connection.Ws.ReadJSON(&receive)
			if err != nil {
				return false
			}
			// Check if the message is the response
			if receive.Id == send.Id {
				return true
			} else {
				// Error
				return false
			}
		}
	}
	return false
}
</file>

<file path="handlers/client.go">
package handlers

import (
	"time"

	// Import local packages
	"github.com/ChatGPT-Hackers/ChatGPT-API-server/types"
	"github.com/ChatGPT-Hackers/ChatGPT-API-server/utils"

	"github.com/gin-gonic/gin"
)

// // # Client routes
func Client_register(c *gin.Context) {
	// Make websocket connection
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	// Generate connection id
	id := utils.GenerateId()
	// Send connection id
	err = ws.WriteJSON(types.Message{
		Id:      id,
		Message: "Connection id",
	})
	if err != nil {
		return
	}
	// Wait for client to send connection id
	for {
		// Read message
		var message types.Message
		err = ws.ReadJSON(&message)
		if err != nil {
			return
		}
		// Check if the message is the connection id
		if message.Id == id {
			break
		} else {
			// This is probably a reconnect
			// Check if the connection id is in the pool
			connection, ok := connectionPool.Get(message.Id)
			if ok {
				// Close the old connection
				connection.Ws.Close()
				// Remove the connection from the pool
				connectionPool.Delete(message.Id)
			}
			id = message.Id
			break
		}
	}
	// Add connection to the pool
	connection := &types.Connection{
		Id: id,
		Ws: ws,
		// Set last message time to the beginning of time
		LastMessageTime: time.Time{},
		Heartbeat:       time.Now(),
	}
	connectionPool.Set(connection)
	// Debug
	println("New connection:", connection.Id)
}
</file>

<file path="handlers/init.go">
package handlers

import (
	"net/http"

	"github.com/ChatGPT-Hackers/ChatGPT-API-server/types"
	"github.com/gorilla/websocket"
)

var (
	// The websocket upgrader.
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

var connectionPool = types.NewConnectionPool()
var conversationPool = types.NewConversationPool()
</file>

<file path="LICENSE">
MIT License

Copyright (c) 2022 Antonio Cheong

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

<file path="main.go">
package main

import (

	// Import local packages

	"net/http"

	"os"

	"github.com/ChatGPT-Hackers/ChatGPT-API-server/handlers"
	"github.com/ChatGPT-Hackers/ChatGPT-API-server/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// get arg server port and admin key
	if len(os.Args) < 3 {
		println("Usage: ./ChatGPT-API-server <port> <admin key>")
		return
	}
	println(os.Args[1], os.Args[2])

	// Make database
	err := utils.DatabaseCreate()
	if err != nil {
		println("Failed to create database:", err.Error())
		return
	}

	router := gin.Default()

	//// # Headers
	// Allow CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	})
	//// # Add routes
	// Register new client connection
	router.GET("/client/register", handlers.Client_register)
	router.POST("/api/ask", handlers.API_ask)
	router.GET("/api/connections", handlers.API_getConnections)
	router.POST("/admin/users/add", handlers.Admin_userAdd)
	router.POST("/admin/users/delete", handlers.Admin_userDel)
	router.GET("/admin/users", handlers.Admin_usersGet)

	// Add a health endpoint
	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// Start server
	router.Run(":" + os.Args[1])
}
</file>

<file path="types/types.go">
package types

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Message struct {
	Id      string `json:"id"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type ChatGptResponse struct {
	Id             string `json:"id"`
	ResponseId     string `json:"response_id"`
	ConversationId string `json:"conversation_id"`
	Content        string `json:"content"`
	Error          string `json:"error"`
}

type ChatGptRequest struct {
	MessageId      string `json:"message_id"`
	ConversationId string `json:"conversation_id"`
	ParentId       string `json:"parent_id"`
	Content        string `json:"content"`
}

type Connection struct {
	// The websocket connection.
	Ws *websocket.Conn
	// The connecton id.
	Id string
	// Last heartbeat time.
	Heartbeat time.Time
	// Last message time.
	LastMessageTime time.Time
}

type ConnectionPool struct {
	Connections map[string]*Connection
	Mu          sync.RWMutex
}

func (p *ConnectionPool) Get(id string) (*Connection, bool) {
	p.Mu.RLock()
	defer p.Mu.RUnlock()
	conn, ok := p.Connections[id]
	if conn == nil {
		ok = false
	}
	return conn, ok
}

func (p *ConnectionPool) Set(conn *Connection) {
	p.Mu.Lock()
	defer p.Mu.Unlock()
	p.Connections[conn.Id] = conn
}

func (p *ConnectionPool) Delete(id string) error {
	p.Mu.Lock()
	defer p.Mu.Unlock()
	delete(p.Connections, id)
	return nil
}

func NewConnectionPool() *ConnectionPool {
	return &ConnectionPool{
		Connections: make(map[string]*Connection),
	}
}

type Conversation struct {
	Id           string `json:"id"`
	ConnectionId string `json:"connection_id"`
}

type ConversationPool struct {
	Conversations map[string]*Conversation
	Mu            sync.RWMutex
}

func (p *ConversationPool) Get(id string) (*Conversation, bool) {
	p.Mu.RLock()
	defer p.Mu.RUnlock()
	conversation, ok := p.Conversations[id]
	return conversation, ok
}

func (p *ConversationPool) Set(conversation *Conversation) {
	p.Mu.Lock()
	defer p.Mu.Unlock()
	p.Conversations[conversation.Id] = conversation
}

func (p *ConversationPool) Delete(id string) {
	p.Mu.Lock()
	defer p.Mu.Unlock()
	delete(p.Conversations, id)
}

func NewConversationPool() *ConversationPool {
	return &ConversationPool{
		Conversations: make(map[string]*Conversation),
	}
}
</file>

<file path="utils/auth.go">
package utils

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func DatabaseCreate() error {
	// Create Data directory if it doesn't already exist
	if _, err := os.Stat("./Data"); os.IsNotExist(err) {
		err = os.Mkdir("./Data", 0755)
		if err != nil {
			return fmt.Errorf("error creating Data directory: %v", err)
		}
	}
	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "./Data/auth.db")
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}
	defer db.Close()

	// Create the table if it doesn't already exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tokens (user_id TEXT PRIMARY KEY, token TEXT UNIQUE)`)
	if err != nil {
		return fmt.Errorf("error creating table: %v", err)
	}

	return nil
}

func DatabaseInsert(user_id string, token string) error {
	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "./Data/auth.db")
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}
	defer db.Close()

	// Insert the token into the database
	_, err = db.Exec(`INSERT INTO tokens (user_id, token) VALUES (?, ?)`, user_id, token)
	if err != nil {
		return fmt.Errorf("error inserting token: %v", err)
	}

	return nil
}

func DatabaseDelete(user_id string) error {
	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "./Data/auth.db")
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}
	defer db.Close()

	// Delete the token from the database
	_, err = db.Exec(`DELETE FROM tokens WHERE user_id = ?`, user_id)
	if err != nil {
		return fmt.Errorf("error deleting token: %v", err)
	}

	return nil
}

type User struct {
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}

func DatabaseSelectAll() ([]User, error) {
	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "./Data/auth.db")
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}
	defer db.Close()

	// Select the token from the database
	rows, err := db.Query(`SELECT * FROM tokens`)
	if err != nil {
		return nil, fmt.Errorf("error selecting token: %v", err)
	}
	defer rows.Close()

	// user map ({"users": ["user_id": "...", "token": "..."], ...})
	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.UserID, &user.Token)
		if err != nil {
			return nil, fmt.Errorf("error scanning rows: %v", err)
		}
		users = append(users, user)
	}

	return users, nil

}

// Verify admin key
func VerifyAdminKey(key string) bool {
	return key == os.Args[2]
}

func VerifyToken(token string) (bool, error) {
	// Check if token is admin key
	if VerifyAdminKey(token) {
		return true, nil
	}
	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "./Data/auth.db")
	if err != nil {
		return false, fmt.Errorf("error opening database: %v", err)
	}
	defer db.Close()

	// Select the token from the database
	rows, err := db.Query(`SELECT * FROM tokens WHERE token = ?`, token)
	if err != nil {
		return false, fmt.Errorf("error selecting token: %v", err)
	}
	defer rows.Close()

	// Check if the token exists
	if rows.Next() {
		return true, nil
	} else {
		return false, nil
	}
}
</file>

<file path="utils/utils.go">
package utils

import (
	"github.com/google/uuid"
)

func GenerateId() string {
	return uuid.New().String()
}
</file>

<file path=".gitignore">
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with `go test -c`
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/
Data/auth.db
</file>

<file path="README.md">
> # Official API released by OpenAI. Please use that instead. The model name is `text-chat-davinci-002-20230126`

# ChatGPT API Server
[![Release Go Binaries](https://github.com/acheong08/ChatGPT-API-server/actions/workflows/release.yml/badge.svg)](https://github.com/acheong08/ChatGPT-API-server/actions/workflows/release.yml)
# Quickstart

## Setup

1. Install Go
2. `go install github.com/acheong08/ChatGPT-API-server@latest`

If the latest commit fails, try using one of the release binaries

# Build

1. `git clone https://github.com/acheong08/ChatGPT-API-server/`
2. `cd ChatGPT-API-server`
3. `go install .`

# Usage

`ChatGPT-API-server <port> <API_KEY>`

The admin key can be anything you want. It's just for authenticating yourself.

# Connect agents

Take note of your IP address or domain name. This could be `localhost` or a remote IP address. The default port is `8080`

Check out our [firefox agent](https://github.com/acheong08/ChatGPT-API-agent). More versions in the works.

There is also a [Python based client](https://github.com/ahmetkca/chatgpt-unofficial-api-docker/tree/ChatGPT-API-agent) by @ahmetkca (WIP)

# Usage

## Quickstart

(After connecting agents)

```bash
 $ curl "http://localhost:8080/api/ask" -X POST --header 'Authorization: <API_KEY>' -d '{"content": "Hello world", "conversation_id": "<optional>", "parent_id": "<optional>"}'
```
Note: if you want to use `conversation_id`, you also need to use `parent_id`!

## Routes

```go
	router.GET("/client/register", handlers.Client_register) // Used by agent
	router.POST("/api/ask", handlers.API_ask) // For making ChatGPT requests
	router.GET("/api/connections", handlers.API_getConnections) // For debugging
	router.POST("/admin/users/add", handlers.Admin_userAdd) // Adds an API token
	router.POST("/admin/users/delete", handlers.Admin_userDel) // Invalidates a token (based on user_id)
	router.GET("/admin/users", handlers.Admin_usersGet) // Get all users
```

### Parameters for each route

#### /client/register (GET)

N/A. Used for websocket

#### /api/ask (POST)

Headers: `Authorization: <USER_TOKEN>`

_The user token can be set by the admin via /admin/users/add. You can also use the api key as the token. Both work by default_

Data:

```json
{
  "content": "Hello world",
  "conversation_id": "<optional>",
  "parent_id": "<optional>"
}
```

Do not enter conversation or parent id if not available.
If you want to use either of these, you need to specify both! i.e. `request.parent_id=response.response_id` and `request.conversation_id=response.conversation_id`

Response:

```json
{
  "id": "",
  "response_id": "<UUID>",
  "conversation_id": "<UUID>",
  "content": "<string>",
  "error": ""
}
```

#### /api/connections (GET)

Headers: None

Data: None

Response:

```json
{
  "connections": [
    {
      "Ws": {},
      "Id": "<UUID>",
      "Heartbeat": "<Time string>",
      "LastMessageTime": "<Time string>"
    }
  ]
}
```

#### /admin/users/add (POST)

Headers: None

Data:

```json
{
  "admin_key": "<string>"
}
```

Response:

```json
{
  "user_id": "<UUID>",
  "token": "<UUID>"
}
```

#### /admin/users/delete (POST)

Headers: None

Data:

```json
{
  "admin_key": "<string>",
  "user_id": "<UUID>"
}
```

Response:

```json
{ "message": "User deleted" }
```

#### /admin/users (GET)

Parameters: `?admin_key=<string>`

Example usage: `curl "http://localhost:8080/admin/users?admin_key=some_random_key"`

Response:

```json
{
  "users": [
    {
      "user_id": "<UUID>",
      "token": "<UUID>"
    },
    {
      "user_id": "<UUID>",
      "token": "<UUID>"
    },
    {
      "user_id": "<UUID>",
      "token": "<UUID>"
    },
    {
      "user_id": "<UUID>",
      "token": "<UUID>"
    },
    ...
  ]
}
```

# Docker

open `docker-compose.yml` and add your own custom api-key in `<API_KEY>` section

```yaml
version: "3"

services:
  chatgpt-api-server:
    build: .
    ports:
      - "8080:8080"
    command: ["ChatGPT-API-server", "8080", "<API_KEY>", "-listen", "0.0.0.0"]
```

then run:

`docker-compose up` or `docker-compose up -d` (if you want a persistent instance)
</file>

</files>
