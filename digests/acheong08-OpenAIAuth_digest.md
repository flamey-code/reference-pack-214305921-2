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
.github/workflows/pypi.yml
.gitignore
auth/OpenAiAuth.go
go.mod
LICENSE
main.go
Makefile
README.md
requirements.txt
setup.py
src/OpenAIAuth.py
</directory_structure>

<files>
This section contains the contents of the repository's files.

<file path=".github/workflows/pypi.yml">
name: Upload Python Package

on:
  release:
    types: [published]
  push:
    branches: [main]
  pull_request:
    types: [opened, reopened, synchronize]

permissions:
  contents: read

jobs:
  ci:
    if: github.event_name != 'release'
    runs-on: ${{ matrix.os }}
    strategy:
      fail-fast: false
      matrix:
        python-version: ["3.9", "3.11"]
        os: [ubuntu-latest, windows-latest]
    steps:
      - uses: actions/checkout@v3
      - name: Set up Python ${{ matrix.python-version }}
        uses: actions/setup-python@v4
        with:
          python-version: ${{ matrix.python-version }}
      - name: Install dependencies
        run: make
      - name: Build CI
        run: make build
      - name: Syntax CI
        run: make ci

  deploy:
    runs-on: ubuntu-latest
    if: github.event_name == 'release'
    steps:
    - uses: actions/checkout@v3
    - name: Set up Python
      uses: actions/setup-python@v4
      with:
        python-version: '3.11.0'
    - name: Install dependencies
      run: make
    - name: Build package
      run: make build
    - name: Publish package
      uses: pypa/gh-action-pypi-publish@v1.8.5
      with:
        user: __token__
        password: ${{ secrets.PYPI_API_TOKEN }}
</file>

<file path="auth/OpenAiAuth.go">
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

<file path="go.mod">
module github.com/acheong08/OpenAIAuth

go 1.20

require (
	github.com/bogdanfinn/fhttp v0.5.19
	github.com/bogdanfinn/tls-client v1.3.8
	github.com/nirasan/go-oauth-pkce-code-verifier v0.0.0-20220510032225-4f9f17eaec4c
)

require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/bogdanfinn/utls v1.5.15 // indirect
	github.com/klauspost/compress v1.15.12 // indirect
	github.com/tam7t/hpkp v0.0.0-20160821193359-2b70b4024ed5 // indirect
	golang.org/x/crypto v0.1.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
)
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
	"encoding/json"
	"fmt"
	"os"

	"github.com/acheong08/OpenAIAuth/auth"
)

func main() {
	auth := auth.NewAuthenticator(os.Getenv("OPENAI_EMAIL"), os.Getenv("OPENAI_PASSWORD"), os.Getenv("PROXY"))
	err := auth.Begin()
	if err != nil {
		println("Error: " + err.Details)
		println("Location: " + err.Location)
		println("Status code: " + fmt.Sprint(err.StatusCode))
		println("Embedded error: " + err.Error.Error())
		return
	}
	// if os.Getenv("PROXY") != "" {
	_, err = auth.GetPUID()
	if err != nil {
		println("Error: " + err.Details)
		println("Location: " + err.Location)
		println("Status code: " + fmt.Sprint(err.StatusCode))
		println("Embedded error: " + err.Error.Error())
		return
	}
	// }
	// JSON encode auth.GetAuthResult()
	result := auth.GetAuthResult()
	result_json, _ := json.Marshal(result)
	println(string(result_json))
}
</file>

<file path="Makefile">
.PHONY: docs
init:
	python -m pip install --upgrade pip
	python -m pip install -r ./requirements.txt --upgrade
	python -m pip install build setuptools wheel flake8 --upgrade
build:
	python -m build
ci:
	python -m flake8 src --count --exit-zero --max-complexity=10 --max-line-length=127 --statistics
	python -m flake8 src --count --select=E9,F63,F7,F82 --show-source --statistics
	python setup.py install
</file>

<file path="requirements.txt">
requests
</file>

<file path="setup.py">
from setuptools import find_packages
from setuptools import setup

setup(
    name="OpenAIAuth",
    version="3.0.0",
    license="MIT",
    author="pengzhile",
    author_email="acheong@student.dalat.org",
    description="OpenAI Authentication Reverse Engineered",
    packages=find_packages("src"),
    package_dir={"": "src"},
    py_modules=["OpenAIAuth"],
    url="https://github.com/acheong08/OpenAIAuth",
    project_urls={"Bug Report": "https://github.com/acheong08/OpenAIAuth/issues/new"},
    install_requires=[
        "tls_client",
    ],
    classifiers=[
        "License :: OSI Approved :: MIT License",
        "Intended Audience :: Developers",
        "Natural Language :: English",
        "Topic :: Software Development :: Libraries :: Python Modules",
        "Programming Language :: Python :: 3 :: Only",
        "Programming Language :: Python :: 3.7",
        "Programming Language :: Python :: 3.8",
        "Programming Language :: Python :: 3.9",
        "Programming Language :: Python :: 3.10",
        "Programming Language :: Python :: 3.11",
    ],
    long_description=open("README.md", "rt", encoding="utf-8").read(),
    long_description_content_type="text/markdown",
)
</file>

<file path="src/OpenAIAuth.py">
# Credits to github.com/rawandahmad698/PyChatGPT
import re
import os
import urllib

import tls_client as requests


class Error(Exception):
    """
    Base error class
    """

    location: str
    status_code: int
    details: str

    def __init__(self, location: str, status_code: int, details: str):
        self.location = location
        self.status_code = status_code
        self.details = details


class Auth0:
    """
    OpenAI Authentication Reverse Engineered
    """

    def __init__(
        self,
        email_address: str,
        password: str,
        puid: str = None,
        proxy: str = None,
    ):
        puid = puid or os.environ.get("PUID")
        # if not puid:
        #     raise ValueError("PUID is required")
        self.email_address = email_address
        self.password = password
        self.proxy = proxy
        self.session = requests.Session(
            client_identifier="chrome112",
            random_tls_extension_order=True,
        )
        proxies = {
            "http": self.proxy,
            "https": self.proxy,
        }
        self.session.proxies.update(proxies)
        self.access_token: str = None
        self.user_agent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"
        self.session.cookies.set("_puid", puid)

    @staticmethod
    def url_encode(string: str) -> str:
        """
        URL encode a string
        :param string:
        :return:
        """
        return urllib.parse.quote(string)

    def begin(self) -> None:
        """
        In part two, We make a request to https://chat.openai.com/api/auth/csrf and grab a fresh csrf token
        """
        url = "https://chat.openai.com/api/auth/csrf"
        headers = {
            "Host": "chat.openai.com",
            "Accept": "*/*",
            "Connection": "keep-alive",
            "User-Agent": self.user_agent,
            "Accept-Language": "en-GB,en-US;q=0.9,en;q=0.8",
            "Referer": "https://chat.openai.com/auth/login",
            "Accept-Encoding": "gzip, deflate, br",
        }
        response = self.session.get(
            url=url,
            headers=headers,
        )
        if response.status_code == 200 and "json" in response.headers["Content-Type"]:
            csrf_token = response.json()["csrfToken"]
            # self.session.cookies.set("__Host-next-auth.csrf-token", csrf_token)
            self.__part_one(token=csrf_token)
        else:
            error = Error(
                location="begin",
                status_code=response.status_code,
                details=response.text,
            )
            print(error.details)
            raise error

    def __part_one(self, token: str) -> None:
        """
        We reuse the token from part to make a request to /api/auth/signin/auth0?prompt=login
        """
        url = "https://chat.openai.com/api/auth/signin/auth0?prompt=login"
        payload = f"callbackUrl=%2F&csrfToken={token}&json=true"
        headers = {
            "Host": "chat.openai.com",
            "User-Agent": self.user_agent,
            "Content-Type": "application/x-www-form-urlencoded",
            "Accept": "*/*",
            "Sec-Gpc": "1",
            "Accept-Language": "en-US,en;q=0.8",
            "Origin": "https://chat.openai.com",
            "Sec-Fetch-Site": "same-origin",
            "Sec-Fetch-Mode": "cors",
            "Sec-Fetch-Dest": "empty",
            "Referer": "https://chat.openai.com/auth/login",
            "Accept-Encoding": "gzip, deflate",
            #
        }
        response = self.session.post(url=url, headers=headers, data=payload)
        if response.status_code == 200 and "json" in response.headers["Content-Type"]:
            url = response.json()["url"]
            if (
                url == "https://chat.openai.com/api/auth/error?error=OAuthSignin"
                or "error" in url
            ):
                error = Error(
                    location="__part_one",
                    status_code=response.status_code,
                    details="You have been rate limited. Please try again later.",
                )
                raise error
            self.__part_two(url=url)
        else:
            error = Error(
                location="__part_one",
                status_code=response.status_code,
                details=response.text,
            )
            raise error

    def __part_two(self, url: str) -> None:
        """
        We make a GET request to url
        :param url:
        :return:
        """
        headers = {
            "Host": "auth0.openai.com",
            "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
            "Connection": "keep-alive",
            "User-Agent": self.user_agent,
            "Accept-Language": "en-US,en;q=0.9",
            "Referer": "https://chat.openai.com/",
        }
        response = self.session.get(
            url=url,
            headers=headers,
        )
        if response.status_code == 302 or response.status_code == 200:
            state = re.findall(r"state=(.*)", response.text)[0]
            state = state.split('"')[0]
            self.__part_three(state=state)
        else:
            error = Error(
                location="__part_two",
                status_code=response.status_code,
                details=response.text,
            )
            raise error

    def __part_three(self, state: str) -> None:
        """
        We use the state to get the login page
        """
        url = f"https://auth0.openai.com/u/login/identifier?state={state}"

        headers = {
            "Host": "auth0.openai.com",
            "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
            "Connection": "keep-alive",
            "User-Agent": self.user_agent,
            "Accept-Language": "en-US,en;q=0.9",
            "Referer": "https://chat.openai.com/",
        }
        response = self.session.get(url, headers=headers)
        if response.status_code == 200:
            self.__part_four(state=state)
        else:
            error = Error(
                location="__part_three",
                status_code=response.status_code,
                details=response.text,
            )
            raise error

    def __part_four(self, state: str) -> None:
        """
        We make a POST request to the login page with the captcha, email
        :param state:
        :return:
        """
        url = f"https://auth0.openai.com/u/login/identifier?state={state}"
        email_url_encoded = self.url_encode(self.email_address)

        payload = (
            f"state={state}&username={email_url_encoded}&js-available=false&webauthn-available=true&is"
            f"-brave=false&webauthn-platform-available=true&action=default "
        )

        headers = {
            "Host": "auth0.openai.com",
            "Origin": "https://auth0.openai.com",
            "Connection": "keep-alive",
            "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
            "User-Agent": self.user_agent,
            "Referer": f"https://auth0.openai.com/u/login/identifier?state={state}",
            "Accept-Language": "en-US,en;q=0.9",
            "Content-Type": "application/x-www-form-urlencoded",
        }
        response = self.session.post(
            url,
            headers=headers,
            data=payload,
        )
        if response.status_code == 302 or response.status_code == 200:
            self.__part_five(state=state)
        else:
            error = Error(
                location="__part_four",
                status_code=response.status_code,
                details="Your email address is invalid.",
            )
            raise error

    def __part_five(self, state: str) -> None:
        """
        We enter the password
        :param state:
        :return:
        """

        email_url_encoded = self.url_encode(self.email_address)
        password_url_encoded = self.url_encode(self.password)
        payload = f"state={state}&username={email_url_encoded}&password={password_url_encoded}&action=default"
        url = f"https://auth0.openai.com/u/login/password?state={state}"
        headers = {
            "Host": "auth0.openai.com",
            "Origin": "https://auth0.openai.com",
            "Connection": "keep-alive",
            "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
            "User-Agent": self.user_agent,
            "Referer": f"https://auth0.openai.com/u/login/password?state={state}",
            "Accept-Language": "en-US,en;q=0.9",
            "Content-Type": "application/x-www-form-urlencoded",
        }
        response = self.session.post(
            url,
            headers=headers,
            allow_redirects=False,
            data=payload,
        )
        if response.status_code == 302:
            redirect_url = response.headers.get("Location")
            self.__part_six(old_state=state, redirect_url=redirect_url)
        else:
            error = Error(
                location="__part_five",
                status_code=response.status_code,
                details="Your credentials are invalid.",
            )
            raise error

    def __part_six(self, old_state: str, redirect_url) -> None:
        url = "https://auth0.openai.com" + redirect_url
        headers = {
            "Host": "auth0.openai.com",
            "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
            "Connection": "keep-alive",
            "User-Agent": self.user_agent,
            "Accept-Language": "en-GB,en-US;q=0.9,en;q=0.8",
            "Referer": f"https://auth0.openai.com/u/login/password?state={old_state}",
        }
        response = self.session.get(url, headers=headers, allow_redirects=False)
        if response.status_code == 302:
            redirect_url = response.headers.get("Location")
            self.__part_seven(redirect_url=redirect_url, previous_url=url)
        else:
            error = Error(
                location="__part_six",
                status_code=response.status_code,
                details=response.text,
            )
            raise error

    def __part_seven(self, redirect_url: str, previous_url: str) -> None:
        url = redirect_url
        headers = {
            "Host": "chat.openai.com",
            "Accept": "application/json",
            "Connection": "keep-alive",
            "User-Agent": self.user_agent,
            "Accept-Language": "en-GB,en-US;q=0.9,en;q=0.8",
            "Referer": previous_url,
        }
        response = self.session.get(url, headers=headers, allow_redirects=True)
        if response.status_code == 200:
            return
        else:
            error = Error(
                location="__part_seven",
                status_code=response.status_code,
                details=response.text,
            )
            raise error

    def get_access_token(self):
        """
        Gets access token
        """
        self.begin()
        response = self.session.get(
            "https://chat.openai.com/api/auth/session",
        )
        if response.status_code == 200:
            self.access_token = response.json().get("accessToken")
            return self.access_token
        else:
            error = Error(
                location="get_access_token",
                status_code=response.status_code,
                details=response.text,
            )
            raise error

    def get_puid(self) -> str:
        url = os.getenv("OPENAI_MODELS_URL", "https://bypass.churchless.tech/models")
        headers = {
            "Authorization": "Bearer " + self.access_token,
        }
        resp = self.session.get(url, headers=headers)
        if resp.status_code == 200:
            # Get _puid cookie
            puid = resp.headers.get("Set-Cookie", "")
            if not puid:
                raise Exception("Get _puid cookie failed.")
            self.puid = puid.split("_puid=")[1].split(";")[0]
            return self.puid
        else:
            raise Exception(resp.text)


if __name__ == "__main__":
    import os

    email = os.getenv("OPENAI_EMAIL")
    password = os.getenv("OPENAI_PASSWORD")
    openai = Auth0(email, password)
    print(openai.get_access_token())

    print(openai.get_puid())
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
#   According to pypa/pipenv#598, it is recommended to include Pipfile.lock in version control.
#   However, in case of collaboration, if having platform-specific dependencies or dependencies
#   having no cross-platform support, pipenv may install dependencies that don't work, or not
#   install all needed dependencies.
#Pipfile.lock

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
test.py

.idea
</file>

<file path="README.md">
# OpenAIAuth
Fetch access tokens for chat.openai.com

## Python version
```py
from OpenAIAuth import Auth0
auth = Auth0(email_address="example@example.com", password="example_password")
access_token = auth.get_access_token()
```

## Go version
```go
package main

import (
	"fmt"
	"os"

	"github.com/acheong08/OpenAIAuth/auth"
)

func main() {
	auth := auth.NewAuthenticator(os.Getenv("OPENAI_EMAIL"), os.Getenv("OPENAI_PASSWORD"), os.Getenv("PROXY"))
	err := auth.Begin()
	if err.Error != nil {
		println("Error: " + err.Details)
		println("Location: " + err.Location)
		println("Status code: " + fmt.Sprint(err.StatusCode))
		println("Embedded error: " + err.Error.Error())
		return
	}
	token, err := auth.GetAccessToken()
	if err.Error != nil {
		println("Error: " + err.Details)
		println("Location: " + err.Location)
		println("Status code: " + fmt.Sprint(err.StatusCode))
		println("Embedded error: " + err.Error.Error())
		return
	}
	fmt.Println(token)
}
```

## Credits
- @linweiyuan
- @rawandahmad698
- @pengzhile
</file>

</files>
