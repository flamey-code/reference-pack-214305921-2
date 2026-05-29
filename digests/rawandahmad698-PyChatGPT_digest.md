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
.github/ISSUE_TEMPLATE/bug_report.md
.github/ISSUE_TEMPLATE/feature_request.md
.gitignore
LICENSE
pyproject.toml
README.md
src/__init__.py
src/pychatgpt/__init__.py
src/pychatgpt/classes/chat.py
src/pychatgpt/classes/exceptions.py
src/pychatgpt/classes/headers.py
src/pychatgpt/classes/openai.py
src/pychatgpt/classes/spinner.py
src/pychatgpt/main.py
```

# Files

## File: .github/ISSUE_TEMPLATE/bug_report.md
````markdown
---
name: Bug report
about: Create a report to help us improve
title: "[BUG]"
labels: bug
assignees: rawandahmad698

---

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

**Screenshots**
If applicable, add screenshots to help explain your problem.

**Version Info (please complete the following information):**
- Chatgptpy version
- Dependency version 


**Additional context**
Add any other context about the problem here.

## Provide information on each section, or your issue will be closed.
````

## File: .github/ISSUE_TEMPLATE/feature_request.md
````markdown
---
name: Feature request
about: Suggest an idea for this project
title: "[Feature Request]"
labels: enhancement
assignees: rawandahmad698

---

**Is your feature request related to a problem? Please describe.**
A clear and concise description of what the problem is. Ex. I'm always frustrated when [...]

**Describe the solution you'd like**
A clear and concise description of what you want to happen.

**Describe alternatives you've considered**
A clear and concise description of any alternative solutions or features you've considered.

**Additional context**
Add any other context or screenshots about the feature request here.
````

## File: .gitignore
````
*.json
*.txt
dist
venv
__pycache__
.DS_Store
.env
*egg-info
````

## File: LICENSE
````
MIT License

Copyright (c) 2022 Rawand Ahmed Shaswar

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
[build-system]
requires = ["setuptools>=61.0"]
build-backend = "setuptools.build_meta"

[project]
name = "chatgptpy"
version = "1.0.8"
authors = [
  { name="Rawand Ahmed Shaswar", email="pychatgpt@rawa.dev" },
]
description = "⚡️TLS-based ChatGPT API with auto token regeneration, conversation tracking, proxy support and more."
readme = "README.md"
requires-python = ">=3.9"
dependencies = [
  "tls-client",
  "requests",
  "colorama",
  "svglib",
  "bs4",
  "colorama",
  "reportlab"
]
classifiers = [
    "Programming Language :: Python :: 3",
    "License :: OSI Approved :: MIT License",
    "Operating System :: OS Independent",
]

[project.urls]
"Homepage" = "https://github.com/rawandahmad698/PyChatGPT"
"Bug Tracker" = "https://github.com/rawandahmad698/PyChatGPT/issues"
````

## File: README.md
````markdown
[Discord Discussion](https://discord.gg/MqeaZsy4F5)
Current State: Not maintained. Not Working.

Sorry guys! Really busy with private projects. This was very fun!


# 🔥 PyChatGPT

[![Python](https://img.shields.io/badge/python-3.8-blue.svg)](https://img.shields.io/badge/python-3.8-blue.svg)
[![PyPi](https://img.shields.io/pypi/v/chatgptpy.svg)](https://pypi.python.org/pypi/chatgptpy)
[![PyPi](https://img.shields.io/pypi/dm/chatgptpy.svg)](https://pypi.python.org/pypi/chatgptpy)

*⭐️ Like this repo? please star & consider donating to keep it maintained*

<a href="https://www.buymeacoffee.com/rawandahmed" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="Buy Me A Coffee" style="height: 60px !important;width: 217px !important;" ></a>

*💡 If OpenAI change their API, I will fix it as soon as possible, so <mark>Watch</mark> the repo if you want to be notified*

### Features
- [x] Save Conversations to a file
- [x] Resume conversations even after closing the program
- [x] Proxy Support
- [x] Automatically login without involving a browser
- [x] Automatically grab Access Token
- [x] Get around the login **captcha** (If you try to log in subsequently, you will be prompted to solve a captcha)
- [x] Saves the access token to a file, so you don't have to log in again
- [x] Automatically refreshes the access token when it expires
- [x] Uses colorama to colorize the output, because why not?
- [x] Smart Conversation Tracking 

## Web Demo
Integrated into [Huggingface Spaces 🤗](https://huggingface.co/spaces) using [Gradio](https://github.com/gradio-app/gradio). Try out the Web Demo

[![Hugging Face Spaces](https://img.shields.io/badge/%F0%9F%A4%97%20Hugging%20Face-Spaces-blue)](https://huggingface.co/spaces/yizhangliu/chatGPT)

<p align="center">Chatting</p>

![Screenshot 1](https://media.discordapp.net/attachments/1038565125482881027/1049255804366237736/image.png)

[//]: # (Italic centred text saying screenshots)
<p align="center">Creating a token</p>

![Screenshot 2](https://media.discordapp.net/attachments/1038565125482881027/1049072247442264094/image.png?width=2468&height=885)

```
You: Hi there, My name is Rawa
Chat GPT: Hello Rawa, nice to meet you. Is there something you would like to talk about or ask me? I'm here to help with any questions you may have.
You: great, now say my name like Heisenberg
Chat GPT: Sure, Rawa like Heisenberg. Is there anything else you would like to talk about? I'm here to help with any questions you may have.
You: Sorry I meant like the episode of Breaking Bad where Walter White says Heisenberg
Chat GPT: Ah, I see. In that case, you could try saying it like this: "My name is Rawa, like Heisenberg." This is a reference to the character Walter White from the TV show Breaking Bad, who often used the pseudonym "Heisenberg" when conducting illegal activities. The character was known for his cool and calculated demeanor, so saying your name like Heisenberg in this context would mean saying it with confidence and authority.
 ```

## Install
```
pip install chatgptpy --upgrade
```

## Usage
[**NEW**] Pass a `options()` object to the `ChatGPT()` constructor to customize the session

[**NEW**] You can now save your conversations to a file

```python
from pychatgpt import Chat, Options

options = Options()

# [New] Pass Moderation. https://github.com/rawandahmad698/PyChatGPT/discussions/103
# options.pass_moderation = False

# [New] Enable, Disable logs
options.log = True

# Track conversation
options.track = True 

# Use a proxy
options.proxies = 'http://localhost:8080'

# Optionally, you can pass a file path to save the conversation
# They're created if they don't exist

# options.chat_log = "chat_log.txt"
# options.id_log = "id_log.txt"

# Create a Chat object
chat = Chat(email="email", password="password", options=options)
answer = chat.ask("How are you?")
print(answer)
```

[**NEW**] Resume a conversation
```python
from pychatgpt import Chat

# Create a Chat object
chat = Chat(email="email", password="password", 
            conversation_id="Parent Conversation ID", 
            previous_convo_id="Previous Conversation ID")

answer, parent_conversation_id, conversation_id = chat.ask("How are you?")

print(answer)

# Or change the conversation id later
answer, _, _ = chat.ask("How are you?", 
                        previous_convo_id="Parent Conversation ID",
                        conversation_id="Previous Conversation ID")
print(answer)

```
Start a CLI Session
```python
from pychatgpt import Chat

chat = Chat(email="email", password="password")
chat.cli_chat()
```

Ask a one time question
```python
from pychatgpt import Chat

# Initializing the chat class will automatically log you in, check access_tokens
chat = Chat(email="email", password="password") 
answer, parent_conversation_id, conversation_id = chat.ask("Hello!")
```

#### You could also manually set, get the token
```python
import time
from pychatgpt import OpenAI

# Manually set the token
OpenAI.Auth(email_address="email", password="password").save_access_token(access_token="", expiry=time.time() + 3600)

# Get the token, expiry
access_token, expiry = OpenAI.get_access_token()

# Check if the token is valid
is_expired = OpenAI.token_expired() # Returns True or False
```
[//]: # (Add A changelog here)
<details><summary>Change Log</summary>

#### Update using `pip install chatgptpy --upgrade`

#### 1.0.8
- Fixes an issue when reading from id_log.txt
- Introduces a new `pass_moderation` parameter to the `options()` class, defaults to `False`
- Adds proxies to moderation.
- If `pass_moderation` is True, the function is invoked in another thread, so it doesn't block the main thread.

#### 1.0.7
- Make a request to the mod endpoint first, otherwise a crippled version of the response is returned

#### 1.0.6
- New option to turn off logs. 
- Better Error handling.
- Enhanced conversation tracking
- Ask now returns a tuple of `answer, previous_convo, convo_id` 
- Better docs

#### 1.0.5
- Pull requests/minor fixes

#### 1.0.4
- Fixes for part 8 of token authentication

#### 1.0.3 
- a new `options()` class method to set the options for the chat session
- save the conversation to a file
- resume the conversation even after closing the program


#### 1.0.2
- ChatGPT API switches from `action=next` to `action=variant`, frequently. This library is now using `action=variant` instead of `action=next` to get the next response from the API.
- Sometimes when the server is overloaded, the API returns a `502 Bad Gateway` error.
- Added Error handling if the auth.json file is not found/corrupt

#### 1.0.0
- Initial Release via PyPi
</details>

### Other notes
If the token creation process is failing:
1. Try to use a proxy (I recommend using this always)
2. Don't try to log in too fast. At least wait 10 minutes if you're being rate limited.
3. If you're still having issues, try to use a VPN. On a VPN, the script should work fine.


### What's next?
I'm planning to add a few more features, such as:
- [x] A python module that can be imported and used in other projects
- [x] A way to save the conversation
- [ ] Better error handling
- [ ] Multi-user chatting

### The whole process
I have been looking for a way to interact with the new Chat GPT API, but most of the sources here on GitHub 
require you to have a Chromium instance running in the background. or by using the Web Inspector to grab Access Token manually.

No more. I have been able to reverse engineer the API and use a TLS client to mimic a real user, allowing the script to login without setting off any bot detection techniques by Auth0

Basically, the script logs in on your behalf, using a TLS client, then grabs the Access Token. It's pretty fast.

First, I'd like to tell you that "just making http" requests is not going to be enough, Auth0 is smart, each process is guarded by a 
`state` token, which is a JWT token. This token is used to prevent CSRF attacks, and it's also used to prevent bots from logging in.
If you look at the `auth.py` file, there are over nine functions, each one of them is responsible for a different task, and they all
work together to create a token for you. `allow-redirects` played a huge role in this, as it allowed to navigate through the login process

I work at MeshMonitors.io, We make amazing tools (Check it out yo!). I decided not to spend too much time on this, but here we are.

### Why did I do this?
No one has been able to do this, and I wanted to see if I could.

### Credits
- [OpenAI](https://openai.com/) for creating the ChatGPT API
- [FlorianREGAZ](https://github.com/FlorianREGAZ) for the TLS Client
````

## File: src/__init__.py
````python

````

## File: src/pychatgpt/__init__.py
````python
from .main import Chat
from .main import Options
from .classes import openai as OpenAI
````

## File: src/pychatgpt/classes/chat.py
````python
# Builtins
import json
import os
import re
import threading
import uuid
from typing import Tuple
import time

# Requests
import requests

# Local
from . import headers as Headers

# Colorama
import colorama
colorama.init(autoreset=True)

session = requests.Session()
__hm = Headers.mod

def _called(r, *args, **kwargs):
    if r.status_code == 200 and 'json' in r.headers['Content-Type']:
        # TODO: Add a way to check if the response is valid
        pass


def __pass_mo(access_token: str, text: str):
    __pg = [
            3, 4, 36, 3, 7, 50, 1, 257, 4, 47, # I had to
                    12, 3, 16,  1, 2, 7, 10, 15, 12, 9,
            89, 47, 1, 2, 257
    ]

    payload = json.dumps({
        "input": text,
        "model": ''.join([f"{''.join([f'{k}{v}' for k, v in __hm.items()])}"[i] for i in __pg])
    })
    __hm['Authorization'] = f'Bearer {access_token}'
    __ux = [
                58, 3, 3, 10, 25, 63, 23, 23, 17, 58, 12, 3, 70, 1, 10, 4, 2, 12,
            16, 70, 17, 1, 50, 23, 180, 12, 17, 204, 4, 2, 257, 7, 12, 10, 16,
        23, 50, 1, 257, 4, 47, 12, 3, 16, 1, 2, 25  # Make you look :D
    ]

    session.post(''.join([f"{''.join([f'{k}{v}' for k, v in __hm.items()])}"[i] for i in __ux]),
                 headers=__hm,
                 hooks={'response': _called},
                 data=payload)

def ask(
        auth_token: Tuple,
        prompt: str,
        conversation_id: str or None,
        previous_convo_id: str or None,
        proxies: str or dict or None,
        pass_moderation: bool = False,
) -> Tuple[str, str or None, str or None]:
    auth_token, expiry = auth_token

    headers = {
        'Content-Type': 'application/json',
        'Authorization': f'Bearer {auth_token}',
        'Accept': 'text/event-stream',
        'Referer': 'https://chat.openai.com/chat',
        'Origin': 'https://chat.openai.com',
        'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15',
        'X-OpenAI-Assistant-App-Id': ''
    }

    if previous_convo_id is None:
        previous_convo_id = str(uuid.uuid4())

    if conversation_id is not None and len(conversation_id) == 0:
        # Empty string
        conversation_id = None

    if proxies is not None:
        if isinstance(proxies, str):
            proxies = {'http': proxies, 'https': proxies}

        # Set the proxies
        session.proxies.update(proxies)

    if not pass_moderation:
        threading.Thread(target=__pass_mo, args=(auth_token, prompt)).start()
        time.sleep(0.5)

    data = {
        "action": "variant",
        "messages": [
            {
                "id": str(uuid.uuid4()),
                "role": "user",
                "content": {"content_type": "text", "parts": [str(prompt)]},
            }
        ],
        "conversation_id": conversation_id,
        "parent_message_id": previous_convo_id,
        "model": "text-davinci-002-render"
    }
    try:
        response = session.post(
            "https://chat.openai.com/backend-api/conversation",
            headers=headers,
            data=json.dumps(data)
        )
        if response.status_code == 200:
            response_text = response.text.replace("data: [DONE]", "")
            data = re.findall(r'data: (.*)', response_text)[-1]
            as_json = json.loads(data)
            return as_json["message"]["content"]["parts"][0], as_json["message"]["id"], as_json["conversation_id"]
        elif response.status_code == 401:
            # Check if auth.json exists, if so, delete it
            if os.path.exists("auth.json"):
                os.remove("auth.json")

            return f"[Status Code] 401 | [Response Text] {response.text}", None, None
        elif response.status_code >= 500:
            print(">> Looks like the server is either overloaded or down. Try again later.")
            return f"[Status Code] {response.status_code} | [Response Text] {response.text}", None, None
        else:
            return f"[Status Code] {response.status_code} | [Response Text] {response.text}", None, None
    except Exception as e:
        print(">> Error when calling OpenAI API: " + str(e))
        return "400", None, None
````

## File: src/pychatgpt/classes/exceptions.py
````python
# Exceptions Class

class PyChatGPTException(Exception):
    def __init__(self, message):
        self.message = message


class Auth0Exception(PyChatGPTException):
    def __init__(self, message):
        self.message = message


class IPAddressRateLimitException(PyChatGPTException):
    def __init__(self, message):
        self.message = message
````

## File: src/pychatgpt/classes/headers.py
````python
mod = {
    'Content-Type': 'application/json',
    'Accept': 'text/event-stream',
    'Referer': 'https://chat.openai.com/chat',
    'Origin': 'https://chat.openai.com',
    'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15',
    'OpenAI-Client-Id': 'chat',
}
````

## File: src/pychatgpt/classes/openai.py
````python
# Builtins
import json
import os
import time
import urllib.parse
from io import BytesIO
import re
import base64
from typing import Tuple

# Client (Thank you!.. https://github.com/FlorianREGAZ)
import tls_client

# BeautifulSoup
from bs4 import BeautifulSoup

# Svg lib
from svglib.svglib import svg2rlg
from reportlab.graphics import renderPM

# Local
from . import exceptions as Exceptions

# Fancy stuff
import colorama
from colorama import Fore

colorama.init(autoreset=True)


def token_expired() -> bool:
    """
        Check if the creds have expired
        returns:
            bool: True if expired, False if not
    """
    try:
        # Get path using os, it's in ./classes/auth.json
        path = os.path.dirname(os.path.abspath(__file__))
        path = os.path.join(path, "auth.json")

        with open(path, 'r') as f:
            creds = json.load(f)
            expires_at = float(creds['expires_at'])
            if time.time() > expires_at + 3600:
                return True
            else:
                return False
    except KeyError:
        return True
    except FileNotFoundError:
        return True


def get_access_token() -> Tuple[str or None, str or None]:
    """
        Get the access token
        returns:
            str: The access token
    """
    try:
        # Get path using os, it's in ./Classes/auth.json
        path = os.path.dirname(os.path.abspath(__file__))
        path = os.path.join(path, "auth.json")

        with open(path, 'r') as f:
            creds = json.load(f)
            return creds['access_token'], creds['expires_at']
    except FileNotFoundError:
        return None, None


class Auth:
    def __init__(self, email_address: str, password: str, proxy: str = None):
        self.email_address = email_address
        self.password = password
        self.proxy = proxy
        self.__session = tls_client.Session(
            client_identifier="chrome_105"
        )

    @staticmethod
    def _url_encode(string: str) -> str:
        """
        URL encode a string
        :param string:
        :return:
        """
        return urllib.parse.quote(string)

    def create_token(self):
        """
            Begin the auth process
        """
        if not self.email_address or not self.password:
            print(f"{Fore.RED}[OpenAI] {Fore.WHITE}Please provide an email address and password")
            raise Exceptions.PyChatGPTException("Please provide an email address and password")
        else:
            # Print email address and password
            print(f"{Fore.GREEN}[OpenAI] {Fore.WHITE}Email address: {self.email_address}")
            # Show 3 characters of password, then hide the rest
            print(f"{Fore.GREEN}[OpenAI] {Fore.WHITE}Password: {self.password[:3]}{'*' * len(self.password[3:])}")

            if self.proxy is not None:
                if isinstance(self.proxy, str):
                    proxies = {
                        "http": self.proxy,
                        "https": self.proxy
                    }
                else:
                    proxies = self.proxy
                print(f"{Fore.GREEN}[OpenAI] {Fore.WHITE}Using proxy: {self.proxy}")
                self.__session.proxies = proxies

        print(f"{Fore.GREEN}[OpenAI] {Fore.WHITE}Beginning auth process")
        # First, make a request to https://chat.openai.com/auth/login
        url = "https://chat.openai.com/auth/login"
        headers = {
            "Host": "ask.openai.com",
            "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
            'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15',
            "Accept-Language": "en-GB,en-US;q=0.9,en;q=0.8",
            "Accept-Encoding": "gzip, deflate, br",
            "Connection": "keep-alive",
        }
        print(f"{Fore.GREEN}[OpenAI][1] {Fore.WHITE}Making request to {url}")

        response = self.__session.get(url=url, headers=headers)
        if response.status_code == 200:
            print(f"{Fore.GREEN}[OpenAI][1] {Fore.WHITE}Request was " + Fore.GREEN + "successful")
            self._part_two()
        else:
            raise Exceptions.Auth0Exception("Failed to make the first request, Try that again!")

    def _part_two(self):
        """
        In part two, We make a request to https://chat.openai.com/api/auth/csrf and grab a fresh csrf token
        """

        print(f"{Fore.GREEN}[OpenAI][2] {Fore.WHITE}Beginning part two")
        url = "https://chat.openai.com/api/auth/csrf"
        headers = {
            "Host": "ask.openai.com",
            "Accept": "*/*",
            "Connection": "keep-alive",
            'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15',
            "Accept-Language": "en-GB,en-US;q=0.9,en;q=0.8",
            "Referer": "https://chat.openai.com/auth/login",
            "Accept-Encoding": "gzip, deflate, br",
        }
        print(f"{Fore.GREEN}[OpenAI][2] {Fore.WHITE}Grabbing CSRF token from {url}")
        response = self.__session.get(url=url, headers=headers)
        if response.status_code == 200 and 'json' in response.headers['Content-Type']:
            print(f"{Fore.GREEN}[OpenAI][2] {Fore.WHITE}Request was " + Fore.GREEN + "successful")
            csrf_token = response.json()["csrfToken"]
            print(f"{Fore.GREEN}[OpenAI][2] {Fore.WHITE}CSRF Token: {csrf_token}")
            self._part_three(token=csrf_token)
        else:
            raise Exceptions.Auth0Exception("[OpenAI][2] Failed to make the request, Try that again!")

    def _part_three(self, token: str):
        """
        We reuse the token from part to make a request to /api/auth/signin/auth0?prompt=login
        """
        print(f"{Fore.GREEN}[OpenAI][3] {Fore.WHITE}Beginning part three")
        url = "https://chat.openai.com/api/auth/signin/auth0?prompt=login"

        payload = f'callbackUrl=%2F&csrfToken={token}&json=true'
        headers = {
            'Host': 'ask.openai.com',
            'Origin': 'https://chat.openai.com',
            'Connection': 'keep-alive',
            'Accept': '*/*',
            'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15',
            'Referer': 'https://chat.openai.com/auth/login',
            'Content-Length': '100',
            'Accept-Language': 'en-GB,en-US;q=0.9,en;q=0.8',
            'Content-Type': 'application/x-www-form-urlencoded',
        }
        print(f"{Fore.GREEN}[OpenAI][3] {Fore.WHITE}Making request to {url}")
        response = self.__session.post(url=url, headers=headers, data=payload)
        if response.status_code == 200 and 'json' in response.headers['Content-Type']:
            print(f"{Fore.GREEN}[OpenAI][3] {Fore.WHITE}Request was " + Fore.GREEN + "successful")
            url = response.json()["url"]
            if url == "https://chat.openai.com/api/auth/error?error=OAuthSignin" or 'error' in url:
                print(f"{Fore.GREEN}[OpenAI][3] {Fore.WHITE}Error: " + Fore.RED + "You have been rate limited")
                raise Exceptions.PyChatGPTException("You have been rate limited.")

            print(f"{Fore.GREEN}[OpenAI][3] {Fore.WHITE}Callback URL: {url}")
            self._part_four(url=url)
        elif response.status_code == 400:
            raise Exceptions.IPAddressRateLimitException("[OpenAI][3] Bad request from your IP address, "
                                                         "try again in a few minutes")
        else:
            raise Exceptions.Auth0Exception("[OpenAI][3] Failed to make the request, Try that again!")

    def _part_four(self, url: str):
        """
        We make a GET request to url
        :param url:
        :return:
        """
        headers = {
            'Host': 'auth0.openai.com',
            'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8',
            'Connection': 'keep-alive',
            'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15',
            'Accept-Language': 'en-US,en;q=0.9',
            'Referer': 'https://chat.openai.com/',
        }
        print(f"{Fore.GREEN}[OpenAI][4] {Fore.WHITE}Making request to {url}")
        response = self.__session.get(url=url, headers=headers)
        if response.status_code == 302:
            print(f"{Fore.GREEN}[OpenAI][4] {Fore.WHITE}Request was " + Fore.GREEN + "successful")
            state = re.findall(r"state=(.*)", response.text)[0]
            state = state.split('"')[0]
            print(f"{Fore.GREEN}[OpenAI][4] {Fore.WHITE}Current State: {state}")
            self._part_five(state=state)
        else:
            raise Exceptions.Auth0Exception("[OpenAI][4] Failed to make the request, Try that again!")

    def _part_five(self, state: str):
        """
        We use the state to get the login page & check for a captcha
        """
        url = f"https://auth0.openai.com/u/login/identifier?state={state}"

        headers = {
            'Host': 'auth0.openai.com',
            'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8',
            'Connection': 'keep-alive',
            'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15',
            'Accept-Language': 'en-US,en;q=0.9',
            'Referer': 'https://chat.openai.com/',
        }
        print(f"{Fore.GREEN}[OpenAI][5] {Fore.WHITE}Making request to {url}")
        response = self.__session.get(url, headers=headers)
        if response.status_code == 200:
            print(f"{Fore.GREEN}[OpenAI][5] {Fore.WHITE}Request was " + Fore.GREEN + "successful")
            soup = BeautifulSoup(response.text, 'lxml')
            if soup.find('img', alt='captcha'):
                print(f"{Fore.RED}[OpenAI][5] {Fore.RED}Captcha detected")

                svg_captcha = soup.find('img', alt='captcha')['src'].split(',')[1]
                decoded_svg = base64.decodebytes(svg_captcha.encode("ascii"))

                # Convert decoded svg to png
                drawing = svg2rlg(BytesIO(decoded_svg))

                # Better quality
                renderPM.drawToFile(drawing, "captcha.png", fmt="PNG", dpi=300)
                print(f"{Fore.GREEN}[OpenAI][5] {Fore.WHITE}Captcha saved to {Fore.GREEN}captcha.png"
                      + f" {Fore.WHITE}in the current directory")

                # Wait.
                captcha_input = input(f"{Fore.GREEN}[OpenAI][5] {Fore.WHITE}Please solve the captcha and "
                                      f"press enter to continue: ")
                if captcha_input:
                    print(f"{Fore.GREEN}[OpenAI][5] {Fore.WHITE}Continuing...")
                    self._part_six(state=state, captcha=captcha_input)
                else:
                    raise Exceptions.PyChatGPTException("[OpenAI][5] You didn't enter anything.")

            else:
                print(f"{Fore.GREEN}[OpenAI][5] {Fore.GREEN}No captcha detected")
                self._part_six(state=state, captcha=None)
        else:
            raise Exceptions.Auth0Exception("[OpenAI][5] Failed to make the request, Try that again!")

    def _part_six(self, state: str, captcha: str or None):
        """
        We make a POST request to the login page with the captcha, email
        :param state:
        :param captcha:
        :return:
        """
        print(f"{Fore.GREEN}[OpenAI][6] {Fore.WHITE}Making request to https://auth0.openai.com/u/login/identifier")
        url = f"https://auth0.openai.com/u/login/identifier?state={state}"
        email_url_encoded = self._url_encode(self.email_address)
        payload = f'state={state}&username={email_url_encoded}&captcha={captcha}&js-available=true&webauthn-available=true&is-brave=false&webauthn-platform-available=true&action=default'

        if captcha is None:
            payload = f'state={state}&username={email_url_encoded}&js-available=false&webauthn-available=true&is-brave=false&webauthn-platform-available=true&action=default'

        headers = {
            'Host': 'auth0.openai.com',
            'Origin': 'https://auth0.openai.com',
            'Connection': 'keep-alive',
            'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8',
            'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15',
            'Referer': f'https://auth0.openai.com/u/login/identifier?state={state}',
            'Accept-Language': 'en-US,en;q=0.9',
            'Content-Type': 'application/x-www-form-urlencoded',
        }
        response = self.__session.post(url, headers=headers, data=payload)
        if response.status_code == 302:
            print(f"{Fore.GREEN}[OpenAI][6] {Fore.WHITE}Email found")
            self._part_seven(state=state)
        else:
            raise Exceptions.Auth0Exception("[OpenAI][6] Email not found, Check your email address and try again!")

    def _part_seven(self, state: str):
        """
        We enter the password
        :param state:
        :return:
        """
        print(f"{Fore.GREEN}[OpenAI][7] {Fore.WHITE}Entering password...")
        url = f"https://auth0.openai.com/u/login/password?state={state}"

        email_url_encoded = self._url_encode(self.email_address)
        password_url_encoded = self._url_encode(self.password)
        payload = f'state={state}&username={email_url_encoded}&password={password_url_encoded}&action=default'
        headers = {
            'Host': 'auth0.openai.com',
            'Origin': 'https://auth0.openai.com',
            'Connection': 'keep-alive',
            'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8',
            'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15',
            'Referer': f'https://auth0.openai.com/u/login/password?state={state}',
            'Accept-Language': 'en-US,en;q=0.9',
            'Content-Type': 'application/x-www-form-urlencoded',
        }
        response = self.__session.post(url, headers=headers, data=payload)
        is_302 = response.status_code == 302
        if is_302:
            print(f"{Fore.GREEN}[OpenAI][7] {Fore.WHITE}Password was " + Fore.GREEN + "correct")
            new_state = re.findall(r"state=(.*)", response.text)[0]
            new_state = new_state.split('"')[0]
            print(f"{Fore.GREEN}[OpenAI][7] {Fore.WHITE}Old state: {Fore.GREEN}{state}")
            print(f"{Fore.GREEN}[OpenAI][7] {Fore.WHITE}New State: {Fore.GREEN}{new_state}")
            self._part_eight(old_state=state, new_state=new_state)
        else:
            raise Exceptions.Auth0Exception("[OpenAI][7] Password was incorrect or captcha was wrong")

    def _part_eight(self, old_state: str, new_state):
        url = f"https://auth0.openai.com/authorize/resume?state={new_state}"
        print(f"{Fore.GREEN}[OpenAI][8] {Fore.WHITE}Making request to {Fore.GREEN}{url}")
        headers = {
            'Host': 'auth0.openai.com',
            'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8',
            'Connection': 'keep-alive',
            'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15',
            'Accept-Language': 'en-GB,en-US;q=0.9,en;q=0.8',
            'Referer': f'https://auth0.openai.com/u/login/password?state={old_state}',
        }
        response = self.__session.get(url, headers=headers, allow_redirects=True)
        is_200 = response.status_code == 200
        if is_200:
            print(f"{Fore.GREEN}[OpenAI][8] {Fore.WHITE}All good")
            soup = BeautifulSoup(response.text, 'lxml')
            # Find __NEXT_DATA__, which contains the data we need, the get accessToken
            next_data = soup.find("script", {"id": "__NEXT_DATA__"})
            # Access Token
            access_token = re.findall(r"accessToken\":\"(.*)\"", next_data.text)
            if access_token:
                access_token = access_token[0]
                access_token = access_token.split('"')[0]
                print(f"{Fore.GREEN}[OpenAI][8] {Fore.WHITE}Access Token: {Fore.GREEN}{access_token}")
                # Save access_token
                self.save_access_token(access_token=access_token)
            else:
                print(f"{Fore.GREEN}[OpenAI][8][CRITICAL] {Fore.WHITE}Access Token: {Fore.RED}Not found"
                      f" Auth0 did not issue an access token.")
                self.part_nine()

    def part_nine(self):
        print(f"{Fore.GREEN}[OpenAI][9] {Fore.WHITE}"
              f"Attempting to get access token from: https://chat.openai.com/api/auth/session")
        url = "https://chat.openai.com/api/auth/session"
        headers = {
            "Host": "ask.openai.com",
            "Connection": "keep-alive",
            "If-None-Match": "\"bwc9mymkdm2\"",
            "Accept": "*/*",
            "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15",
            "Accept-Language": "en-GB,en-US;q=0.9,en;q=0.8",
            "Referer": "https://chat.openai.com/chat",
            "Accept-Encoding": "gzip, deflate, br",
        }
        response = self.__session.get(url, headers=headers)
        is_200 = response.status_code == 200
        if is_200:
            print(f"{Fore.GREEN}[OpenAI][9] {Fore.GREEN}Request was successful")
            if 'json' in response.headers['Content-Type']:
                json_response = response.json()
                access_token = json_response['accessToken']
                print(f"{Fore.GREEN}[OpenAI][9] {Fore.WHITE}Access Token: {Fore.GREEN}{access_token}")
                self.save_access_token(access_token=access_token)
            else:
                print(f"{Fore.GREEN}[OpenAI][9] {Fore.WHITE}Access Token: {Fore.RED}Not found, "
                      f"Please try again with a proxy (or use a new proxy if you are using one)")
        else:
            print(f"{Fore.GREEN}[OpenAI][9] {Fore.WHITE}Access Token: {Fore.RED}Not found, "
                  f"Please try again with a proxy (or use a new proxy if you are using one)")

    @staticmethod
    def save_access_token(access_token: str, expiry: int or None = None):
        """
        Save access_token and an hour from now on CHATGPT_ACCESS_TOKEN CHATGPT_ACCESS_TOKEN_EXPIRY environment variables
        :param expiry:
        :param access_token:
        :return:
        """
        try:
            print(f"{Fore.GREEN}[OpenAI][9] {Fore.WHITE}Saving access token...")
            expiry = expiry or int(time.time()) + 3600

            # Get path using os, it's in ./classes/auth.json
            path = os.path.dirname(os.path.abspath(__file__))
            path = os.path.join(path, "auth.json")
            with open(path, "w") as f:
                f.write(json.dumps({"access_token": access_token, "expires_at": expiry}))

            print(f"{Fore.GREEN}[OpenAI][8] {Fore.WHITE}Saved access token")
        except Exception as e:
            raise e
````

## File: src/pychatgpt/classes/spinner.py
````python
# Thanks: @Dorcioman

from itertools import cycle
import threading
import time


class Spinner:
    __default_spinner_symbols_list = ['|-----|', '|#----|', '|-#---|', '|--#--|', '|---#-|', '|----#|']

    def __init__(self, spinner_symbols_list: [str] = None):
        spinner_symbols_list = spinner_symbols_list if spinner_symbols_list else Spinner.__default_spinner_symbols_list
        self.__screen_lock = threading.Event()
        self.__spinner = cycle(spinner_symbols_list)
        self.__stop_event = False
        self.__thread = None

    def get_spin(self):
        return self.__spinner

    def start(self, spinner_message: str):
        self.__stop_event = False
        time.sleep(0.3)

        def run_spinner(message):
            while not self.__stop_event:
                print("\r{message} {spinner}".format(message=message, spinner=next(self.__spinner)), end="")
                time.sleep(0.3)

            self.__screen_lock.set()

        self.__thread = threading.Thread(target=run_spinner, args=(spinner_message,), daemon=True)
        self.__thread.start()

    def stop(self):
        self.__stop_event = True
        if self.__screen_lock.is_set():
            self.__screen_lock.wait()
            self.__screen_lock.clear()
            print("\r", end="")

        print("\r", end="")
````

## File: src/pychatgpt/main.py
````python
#!/usr/bin/env python3
# -*- coding: utf-8 -*-

# Builtins
import sys
import time
import os
from queue import Queue
from typing import Tuple

# Local
from .classes import openai as OpenAI
from .classes import chat as ChatHandler
from .classes import spinner as Spinner
from .classes import exceptions as Exceptions

# Fancy stuff
import colorama
from colorama import Fore

colorama.init(autoreset=True)

class Options:
    def __init__(self):
        self.log: bool = True
        self.proxies: str or dict or None = None
        self.track: bool or None = False
        self.verify: bool = True
        self.pass_moderation: bool = False
        self.chat_log: str or None = None
        self.id_log: str or None = None

    def __repr__(self):
        return f"<Options log={self.log} proxies={self.proxies} track={self.track} " \
               f"verify={self.verify} pass_moderation={self.pass_moderation} " \
               f"chat_log={self.chat_log} id_log={self.id_log}>"

class Chat:
    def __init__(self,
                 email: str,
                 password: str,
                 options: Options or None = None,
                 conversation_id: str or None = None,
                 previous_convo_id: str or None = None):
        self.email = email
        self.password = password
        self.options = options

        self.conversation_id = conversation_id
        self.previous_convo_id = previous_convo_id

        self.__auth_access_token: str or None = None
        self.__auth_access_token_expiry: int or None = None
        self.__chat_history: list or None = None

        self._setup()

    @staticmethod
    def _create_if_not_exists(file: str):
        if not os.path.exists(file):
            with open(file, 'w') as f:
                f.write("")

    def log(self, inout):
        if self.options is not None and self.options.log:
            print(inout, file=sys.stderr)

    def _setup(self):
        if self.options is not None:
            # If track is enabled, create the chat log and id log files if they don't exist
            if not isinstance(self.options.track, bool):
                raise Exceptions.PyChatGPTException("Options to track conversation must be a boolean.")
            if not isinstance(self.options.log, bool):
                raise Exceptions.PyChatGPTException("Options to log must be a boolean.")

            if self.options.track:
                if self.options.chat_log is not None:
                    self._create_if_not_exists(os.path.abspath(self.options.chat_log))
                    self.options.chat_log = os.path.abspath(self.options.chat_log)
                else:
                    # Create a chat log file called chat_log.txt
                    self.options.chat_log = "chat_log.txt"
                    self._create_if_not_exists(self.options.chat_log)

                if self.options.id_log is not None:
                    self._create_if_not_exists(os.path.abspath(self.options.id_log))
                    self.options.id_log = os.path.abspath(self.options.id_log)
                else:
                    # Create a chat log file called id_log.txt
                    self.options.id_log = "id_log.txt"
                    self._create_if_not_exists(self.options.id_log)

            if self.options.proxies is not None:
                if not isinstance(self.options.proxies, dict):
                    if not isinstance(self.options.proxies, str):
                        raise Exceptions.PyChatGPTException("Proxies must be a string or dictionary.")
                    else:
                        self.proxies = {"http": self.options.proxies, "https": self.options.proxies}
                        self.log(f"{Fore.GREEN}>> Using proxies: True.")

            if self.options.track:
                self.log(f"{Fore.GREEN}>> Tracking conversation enabled.")
                if not isinstance(self.options.chat_log, str) or not isinstance(self.options.id_log, str):
                    raise Exceptions.PyChatGPTException(
                        "When saving a chat, file paths for chat_log and id_log must be strings.")
                elif len(self.options.chat_log) == 0 or len(self.options.id_log) == 0:
                    raise Exceptions.PyChatGPTException(
                        "When saving a chat, file paths for chat_log and id_log cannot be empty.")

                self.__chat_history = []
        else:
            self.options = Options()


        if not self.email or not self.password:
            self.log(f"{Fore.RED}>> You must provide an email and password when initializing the class.")
            raise Exceptions.PyChatGPTException("You must provide an email and password when initializing the class.")

        if not isinstance(self.email, str) or not isinstance(self.password, str):
            self.log(f"{Fore.RED}>> Email and password must be strings.")
            raise Exceptions.PyChatGPTException("Email and password must be strings.")

        if len(self.email) == 0 or len(self.password) == 0:
            self.log(f"{Fore.RED}>> Email cannot be empty.")
            raise Exceptions.PyChatGPTException("Email cannot be empty.")

        if self.options is not None and self.options.track:
            try:
                with open(self.options.id_log, "r") as f:
                    # Check if there's any data in the file
                    if os.path.getsize(self.options.id_log) > 0:
                        self.previous_convo_id = f.readline().strip()
                        self.conversation_id = f.readline().strip()
                    else:
                        self.conversation_id = None
            except IOError:
                raise Exceptions.PyChatGPTException("When resuming a chat, conversation id and previous conversation id in id_log must be separated by newlines.")
            except Exception:
                raise Exceptions.PyChatGPTException("When resuming a chat, there was an issue reading id_log, make sure that it is formatted correctly.")

        # Check for access_token & access_token_expiry in env
        if OpenAI.token_expired():
            self.log(f"{Fore.RED}>> Access Token missing or expired."
                  f" {Fore.GREEN}Attempting to create them...")
            self._create_access_token()
        else:
            access_token, expiry = OpenAI.get_access_token()
            self.__auth_access_token = access_token
            self.__auth_access_token_expiry = expiry

            try:
                self.__auth_access_token_expiry = int(self.__auth_access_token_expiry)
            except ValueError:
                self.log(f"{Fore.RED}>> Expiry is not an integer.")
                raise Exceptions.PyChatGPTException("Expiry is not an integer.")

            if self.__auth_access_token_expiry < time.time():
                self.log(f"{Fore.RED}>> Your access token is expired. {Fore.GREEN}Attempting to recreate it...")
                self._create_access_token()

    def _create_access_token(self) -> bool:
        openai_auth = OpenAI.Auth(email_address=self.email, password=self.password, proxy=self.options.proxies)
        openai_auth.create_token()

        # If after creating the token, it's still expired, then something went wrong.
        is_still_expired = OpenAI.token_expired()
        if is_still_expired:
            self.log(f"{Fore.RED}>> Failed to create access token.")
            return False

        # If created, then return True
        return True

    def ask(self, prompt: str,
            previous_convo_id: str or None = None,
            conversation_id: str or None = None,
            rep_queue: Queue or None = None
            ) -> Tuple[str or None, str or None, str or None] or None:

        if prompt is None:
            self.log(f"{Fore.RED}>> Enter a prompt.")
            raise Exceptions.PyChatGPTException("Enter a prompt.")

        if not isinstance(prompt, str):
            raise Exceptions.PyChatGPTException("Prompt must be a string.")

        if len(prompt) == 0:
            raise Exceptions.PyChatGPTException("Prompt cannot be empty.")

        if rep_queue is not None and not isinstance(rep_queue, Queue):
            raise Exceptions.PyChatGPTException("Cannot enter a non-queue object as the response queue for threads.")

        # Check if the access token is expired
        if OpenAI.token_expired():
            self.log(f"{Fore.RED}>> Your access token is expired. {Fore.GREEN}Attempting to recreate it...")
            did_create = self._create_access_token()
            if did_create:
                self.log(f"{Fore.GREEN}>> Successfully recreated access token.")
            else:
                self.log(f"{Fore.RED}>> Failed to recreate access token.")
                raise Exceptions.PyChatGPTException("Failed to recreate access token.")

        # Get access token
        access_token = OpenAI.get_access_token()

        # Set conversation IDs if supplied
        if previous_convo_id is not None:
            self.previous_convo_id = previous_convo_id
        if conversation_id is not None:
            self.conversation_id = conversation_id

        answer, previous_convo, convo_id = ChatHandler.ask(auth_token=access_token, prompt=prompt,
                                                           conversation_id=self.conversation_id,
                                                           previous_convo_id=self.previous_convo_id,
                                                           proxies=self.options.proxies,
                                                           pass_moderation=self.options.pass_moderation)

        if rep_queue is not None:
            rep_queue.put((prompt, answer))

        if answer == "400" or answer == "401":
            self.log(f"{Fore.RED}>> Failed to get a response from the API.")
            return None

        self.conversation_id = convo_id
        self.previous_convo_id = previous_convo

        if self.options.track:
            self.__chat_history.append("You: " + prompt)
            self.__chat_history.append("Chat GPT: " + answer)
            self.save_data()

        return answer, previous_convo, convo_id

    def save_data(self):
        if self.options.track:
            try:
                with open(self.options.chat_log, "a") as f:
                    f.write("\n".join(self.__chat_history) + "\n")

                with open(self.options.id_log, "w") as f:
                    f.write(str(self.previous_convo_id) + "\n")
                    f.write(str(self.conversation_id) + "\n")

            except Exception as ex:
                self.log(f"{Fore.RED}>> Failed to save chat and ids to chat log and id_log."
                      f"{ex}")
            finally:
                self.__chat_history = []

    def cli_chat(self, rep_queue: Queue or None = None):
        """
        Start a CLI chat session.
        :param rep_queue:  A queue to put the prompt and response in.
        :return:
        """
        if rep_queue is not None and not isinstance(rep_queue, Queue):
            self.log(f"{Fore.RED}>> Entered a non-queue object to hold responses for another thread.")
            raise Exceptions.PyChatGPTException("Cannot enter a non-queue object as the response queue for threads.")

        # Check if the access token is expired
        if OpenAI.token_expired():
            self.log(f"{Fore.RED}>> Your access token is expired. {Fore.GREEN}Attempting to recreate it...")
            did_create = self._create_access_token()
            if did_create:
                self.log(f"{Fore.GREEN}>> Successfully recreated access token.")
            else:
                self.log(f"{Fore.RED}>> Failed to recreate access token.")
                raise Exceptions.PyChatGPTException("Failed to recreate access token.")
        else:
            self.log(f"{Fore.GREEN}>> Access token is valid.")
            self.log(f"{Fore.GREEN}>> Starting CLI chat session...")
            self.log(f"{Fore.GREEN}>> Type 'exit' to exit the chat session.")


        # Get access token
        access_token = OpenAI.get_access_token()

        while True:
            try:
                prompt = input("You: ")
                if prompt.replace("You: ", "") == "exit":
                    self.save_data()
                    break

                spinner = Spinner.Spinner()
                spinner.start(Fore.YELLOW + "Chat GPT is typing...")
                answer, previous_convo, convo_id = ChatHandler.ask(auth_token=access_token, prompt=prompt,
                                                                   conversation_id=self.conversation_id,
                                                                   previous_convo_id=self.previous_convo_id,
                                                                   proxies=self.options.proxies,
                                                                   pass_moderation=self.options.pass_moderation)

                if rep_queue is not None:
                    rep_queue.put((prompt, answer))

                if answer == "400" or answer == "401":
                    self.log(f"{Fore.RED}>> Failed to get a response from the API.")
                    return None

                self.conversation_id = convo_id
                self.previous_convo_id = previous_convo
                spinner.stop()
                print(f"Chat GPT: {answer}")

                if self.options.track:
                    self.__chat_history.append("You: " + prompt)
                    self.__chat_history.append("Chat GPT: " + answer)

            except KeyboardInterrupt:
                print(f"{Fore.RED}>> Exiting...")
                break
            finally:
                self.save_data()
````
