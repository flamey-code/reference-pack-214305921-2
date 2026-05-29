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
.github/workflows/go.yml
.gitignore
cmd/time.go
docker-compose.yml
Dockerfile
go.mod
LICENSE
README.md
</directory_structure>

<files>
This section contains the contents of the repository's files.

<file path=".github/workflows/go.yml">
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

<file path="cmd/time.go">
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("%d", time.Now().UnixNano()/1000000000)
}
</file>

<file path="docker-compose.yml">
version: '3.3'
services:
    chatgpt-proxy:
        image: chatgpt-proxy-v4
        container_name: chatgpt-proxy-v4
        ports:
            - '8080:8080'
        env_file:
        - .env
</file>

<file path="Dockerfile">
FROM golang
RUN go install github.com/acheong08/ChatGPTProxy@latest
CMD [ "ChatGPTProxy" ]
</file>

<file path="go.mod">
module github.com/acheong08/ChatGPTProxy

go 1.20

require (
	github.com/acheong08/OpenAIAuth v0.0.0-20230625142757-7b01ccd04f63
	github.com/acheong08/endless v0.0.0-20230615162514-90545c7793fd
	github.com/acheong08/funcaptcha v1.9.2
	github.com/bogdanfinn/fhttp v0.5.23
	github.com/bogdanfinn/tls-client v1.4.0
	github.com/gin-gonic/gin v1.9.1
)

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/bogdanfinn/utls v1.5.16 // indirect
	github.com/bytedance/sonic v1.9.2 // indirect
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
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/crypto v0.10.0 // indirect
	golang.org/x/net v0.11.0 // indirect
	golang.org/x/sys v0.9.0 // indirect
	golang.org/x/text v0.10.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
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
main.go
ChatGPT-Proxy-V4
funcaptcha
</file>

<file path="README.md">
# ChatGPT Proxy

Gets around cloudflare via TLS spoofing

## Notes
There is an IP based rate limit. Set a PUID environment variable to get around it
`export PUID="user-..."`
This requires a ChatGPT Plus account

## Building and running
`go build`
`./ChatGPT-Proxy-V4`

## Limitations
This cannot get around an outright IP ban by OpenAI
</file>

</files>
