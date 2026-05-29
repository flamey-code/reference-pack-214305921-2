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
.gitignore
docs/README.md
docs/zh-README.md
examples/async_basic_example.py
examples/async_complex_example.py
examples/basic_example.py
examples/complex_example.py
LICENSE
re_gpt/__init__.py
re_gpt/async_chatgpt.py
re_gpt/errors.py
re_gpt/sync_chatgpt.py
re_gpt/utils.py
README.md
requirements.txt
sampleconfig.ini
setup.py
</directory_structure>

<files>
This section contains the contents of the repository's files.

<file path="docs/README.md">
# Coming soon.... hopefully.
</file>

<file path="docs/zh-README.md">
<div align="center">
  <a href="https://github.com/Zai-Kun/reverse-engineered-chatgpt"></a>

<h1 align="center">Reverse Engineered <a href="https://openai.com/blog/chatgpt">ChatGPT</a> API</h1>

  <p align="center">
    无需API密钥在Python代码中使用OpenAI ChatGPT

[![Stargazers][stars-badge]][stars-url]
[![Forks][forks-badge]][forks-url]
[![Discussions][discussions-badge]][discussions-url]
[![Issues][issues-badge]][issues-url]
[![MIT许可证][license-badge]][license-url]

  [English](../README.md) | 简体中文 
  </p>
    <p align="center">
    <a href="https://github.com/Zai-Kun/reverse-engineered-chatgpt"></a>
    <a href="https://github.com/Zai-Kun/reverse-engineered-chatgpt/issues">报告Bug</a>
    |
    <a href="https://github.com/Zai-Kun/reverse-engineered-chatgpt/discussions">请求新功能</a>
  </p>
</div>

<!-- 目录 -->
<details>
  <summary>目录</summary>
  <ol>
    <li>
      <a href="#关于本项目">关于本项目</a>
      <ul>
        <li><a href="#灵感来源">灵感来源</a></li>
        <li><a href="#工作原理">工作原理</a></li>
        <li><a href="#构建使用">构建使用</a></li>
      </ul>
    </li>
    <li>
      <a href="#开始使用">开始使用</a>
      <ul>
        <li><a href="#前提条件">前提条件</a></li>
        <li><a href="#安装">安装</a></li>
        <li><a href="#获取会话令牌">获取会话令牌</a></li>
      </ul>
    </li>
    <li><a href="#使用方法">使用方法</a>
        <ul>
        <li><a href="#基本示例">基本示例</a></li>
        </ul>
    </li>
    <li><a href="#路线图">路线图</a></li>
    <li><a href="#贡献">贡献</a></li>
    <li><a href="#许可证">许可证</a></li>
    <li><a href="#联系方式">联系方式</a></li>
    <li><a href="#致谢">致谢</a></li>
  </ol>
</details>

## 关于本项目

该项目可用于将OpenAI的ChatGPT服务集成到您的Python代码中。您可以使用这个项目直接从python中提示ChatGPT响应，而无需使用官方API密钥。

如果你想不通过[ChatGPT Plus](https://openai.com/blog/chatgpt-plus)账户使用ChatGPT API，这将非常有用。

### 灵感来源

ChatGPT有一个官方API，可以用于将您的Python代码与之接口，但它需要使用API密钥。这个API密钥只能通过拥有[ChatGPT Plus](https://openai.com/blog/chatgpt-plus)账户获得，这需要20美元/月（截至2023年5月11日）。但是，您可以通过使用[ChatGPT网页界面](https://chat.openai.com/)免费使用ChatGPT。本项目旨在将您的代码与ChatGPT网页版本接口，这样您就可以在不使用API密钥的情况下在Python代码中使用ChatGPT。

### 工作原理

[ChatGPT](https://chat.openai.com/)网页界面的请求已经被反向工程，并直接集成到Python请求中。因此，使用此脚本进行的任何请求都模拟为用户直接在网站上进行的请求。因此，它是免费的，不需要API密钥。

### 构建使用

- [![Python][python-badge]][python-url]

## 开始使用

### 前提条件

- Python >= 3.9

### 安装

```sh
pip install re-gpt
```

## 使用方法

### 简单示例

``` python
from re_gpt import SyncChatGPT

session_token = "__Secure-next-auth.session-token here"
conversation_id = None # 这里填写对话ID


with SyncChatGPT(session_token=session_token) as chatgpt:
    prompt = input("输入你的提示：")

    if conversation_id:
        conversation = chatgpt.get_conversation(conversation_id)
    else:
        conversation = chatgpt.create_new_conversation()

    for message in conversation.chat(prompt):
        print(message["content"], flush=True, end="")

```

### 简单异步示例

``` python
import asyncio
import sys

from re_gpt import AsyncChatGPT

session_token = "__Secure-next-auth.session-token here"
conversation_id = None # 这里填写对话ID

if sys.version_info >= (3, 8) and sys.platform.lower().startswith("win"):
    asyncio.set_event_loop_policy(asyncio.WindowsSelectorEventLoopPolicy())


async def main():
    async with AsyncChatGPT(session_token=session_token) as chatgpt:
        prompt = input("输入你的提示：")

        if conversation_id:
            conversation = chatgpt.get_conversation(conversation_id)
        else:
            conversation = chatgpt.create_new_conversation()

        async for message in conversation.chat(prompt):
            print(message["content"], flush=True, end="")


if __name__ == "__main__":
    asyncio.run(main())
```

## 更多示例

要查看更复杂的示例，请查看存储库中的[examples](/examples)文件夹。

### 获取会话令牌

1. 访问<https://chat.openai.com/chat>并登录或注册。
2. 打开浏览器的开发者工具。
3. 转到`Application`标签页并打开`Cookies`部分。
4. 复制`__Secure-next-auth.session-token`的值并保存。

## 待办事项

- [x] 添加更多示例
- [ ] 添加更好的错误处理
- [x] 实现检索所有ChatGPT聊天的功能
- [ ] 改进文档

## 贡献

贡献是开源社区成为学习、启发和创造的绝佳场所的原因之一。您所做的任何贡献都**非常感激**。

如果您有一个好的建议，可以使这个项目变得更好，请fork本仓库并创建一个拉取请求。
不要忘了给项目加星！再次感谢！

1. Fork项目
2. 创建您的功能分支（`git checkout -b feature/AmazingFeature`）
3. 提交您的更改（`git commit -m 'Add some AmazingFeature'`）
4. 推送到分支（`git push origin feature/AmazingFeature`）
5. 提交拉取请求

## 许可证

根据Apache许可证2.0分发。更多信息请参见[`LICENSE`](https://github.com/Zai-Kun/reverse-engineered-chatgpt/blob/main/LICENSE)。

## 联系/报告Bug

Zai-Kun - [Discord Server](https://discord.gg/ymcqxudVJG)

仓库链接: <https://github.com/Zai-Kun/reverse-engineered-chatgpt>

## 致谢

- [sudoAlphaX (for writing this readme)](https://github.com/sudoAlphaX)

- [yifeikong (curl-cffi module)](https://github.com/yifeikong/curl_cffi)

- [acheong08 (implementation to obtain arkose_token)](https://github.com/acheong08/funcaptcha)

- [pyca (cryptography module)](https://github.com/pyca/cryptography/)

- [Legrandin (pycryptodome module)](https://github.com/Legrandin/pycryptodome/)

- [othneildrew (README Template)](https://github.com/othneildrew)

<!-- MARKDOWN LINKS & IMAGES -->

[forks-badge]: https://img.shields.io/github/forks/Zai-Kun/reverse-engineered-chatgpt
[forks-url]: https://github.com/Zai-Kun/reverse-engineered-chatgpt/network/members
[stars-badge]: https://img.shields.io/github/stars/Zai-Kun/reverse-engineered-chatgpt
[stars-url]: https://github.com/Zai-Kun/reverse-engineered-chatgpt/stargazers
[issues-badge]: https://img.shields.io/github/issues/Zai-Kun/reverse-engineered-chatgpt
[issues-url]: https://github.com/Zai-Kun/reverse-engineered-chatgpt/issues
[discussions-badge]: https://img.shields.io/github/discussions/Zai-Kun/reverse-engineered-chatgpt
[discussions-url]: https://github.com/Zai-Kun/reverse-engineered-chatgpt/discussions
[python-badge]: https://img.shields.io/badge/Python-blue?logo=python&logoColor=yellow
[python-url]: https://www.python.org/
[license-badge]: https://img.shields.io/github/license/Zai-Kun/reverse-engineered-chatgpt
[license-url]: https://github.com/Zai-Kun/reverse-engineered-chatgpt/blob/main/LICENSE
</file>

<file path="examples/async_basic_example.py">
import asyncio
import sys

from re_gpt import AsyncChatGPT

# consts
session_token = "__Secure-next-auth.session-token here"
conversation_id = None  # Set it to the conversation ID if you want to continue an existing chat or None to create a new chat

# If the Python version is 3.8 or higher and the platform is Windows, set the event loop policy
if sys.version_info >= (3, 8) and sys.platform.lower().startswith("win"):
    asyncio.set_event_loop_policy(asyncio.WindowsSelectorEventLoopPolicy())


async def main():
    # Create an asynchronous ChatGPT instance using the session token from the config file
    async with AsyncChatGPT(session_token=session_token) as chatgpt:
        # Get user input for the chat prompt
        prompt = input("Enter your prompt: ")

        # Continue the existing chat using conversation_id or create a new chat if conversation_id is none
        if conversation_id:
            conversation = chatgpt.get_conversation(conversation_id)
        else:
            conversation = chatgpt.create_new_conversation()

        # Iterate through the messages received from the chatgpt and print it
        async for message_chunk in conversation.chat(prompt):
            print(message_chunk["content"], flush=True, end="")


if __name__ == "__main__":
    # Run the asynchronous main function using asyncio.run()
    asyncio.run(main())
</file>

<file path="examples/async_complex_example.py">
import asyncio
import configparser
import sys

from re_gpt import AsyncChatGPT

# Load configuration from 'config.ini'
config = configparser.ConfigParser()
config.read("config.ini")
chat_session = config["session"]

# ANSI color codes for console text formatting
GREEN = "\033[92m"
YELLOW = "\033[93m"
RESET = "\033[0m"

# Required for Windows compatibility
if sys.version_info >= (3, 8) and sys.platform.lower().startswith("win"):
    asyncio.set_event_loop_policy(asyncio.WindowsSelectorEventLoopPolicy())


def print_chat(chat):
    """
    Print formatted chat messages to the console.

    Args:
        chat (dict): The chat data.
    """
    for _, message in chat.get("mapping", {}).items():
        if "message" in message and message["message"]["content"]["parts"][0]:
            role = message["message"]["author"]["role"]
            content = message["message"]["content"]["parts"][0]
            print(f"{GREEN if role == 'user' else YELLOW}{role}: {RESET}{content}\n")


async def main():
    async with AsyncChatGPT(
        # proxies=None,  # Optional proxies for network requests
        session_token=chat_session["token"],  # Use the session token for authentication
    ) as chatgpt:
        if chat_session["conversation_id"]:
            conversation = chatgpt.get_conversation(chat_session["conversation_id"])
        else:
            conversation = chatgpt.create_new_conversation()

        # Fetch and print the existing chat
        fetched_chat = await conversation.fetch_chat()
        print_chat(fetched_chat)

        while True:
            user_input = input(f"{GREEN}user: {RESET}")
            async_chat_stream = conversation.chat(user_input)

            print_header = True
            async for message in async_chat_stream:
                # The 'conversation_id' will be empty if it's a new chat, so we assign the new chat's ID
                if not chat_session["conversation_id"]:
                    chat_session["conversation_id"] = message["conversation_id"]

                # print header for the response
                if print_header:
                    print(f"\n{YELLOW}assistant: {RESET}", end="", flush=True)
                    print_header = False

                # Print the ChatGPT's reply
                print(message["content"], end="", flush=True)

            print("\n")

            # Write the new changes back to the config file
            with open("config.ini", "w") as file:
                config.write(file)


if __name__ == "__main__":
    asyncio.run(main())

# Note: The 'conversation_id' will be found in the chat's url: 'https://chat.openai.com/c/conversation_id'
</file>

<file path="examples/basic_example.py">
from re_gpt import SyncChatGPT

# consts
session_token = "__Secure-next-auth.session-token here"
conversation_id = None  # Set it to the conversation ID if you want to continue an existing chat or None to create a new chat

# Create ChatGPT instance using the session token
with SyncChatGPT(session_token=session_token) as chatgpt:
    prompt = input("Enter your prompt: ")

    # Continue the existing chat using conversation_id or create a new chat if conversation_id is none
    if conversation_id:
        conversation = chatgpt.get_conversation(conversation_id)
    else:
        conversation = chatgpt.create_new_conversation()

    # Iterate through the messages received from the chatgpt and print it
    for message_chunk in conversation.chat(prompt):
        print(message_chunk["content"], flush=True, end="")
</file>

<file path="examples/complex_example.py">
import configparser

from re_gpt import SyncChatGPT

# Load configuration from 'config.ini'
config = configparser.ConfigParser()
config.read("config.ini")
chat_session = config["session"]

# ANSI color codes for console text formatting
GREEN = "\033[92m"
YELLOW = "\033[93m"
RESET = "\033[0m"


def print_chat(chat):
    """
    Print formatted chat messages to the console.

    Args:
        chat (dict): The chat data.
    """
    for _, message in chat.get("mapping", {}).items():
        if "message" in message and message["message"]["content"]["parts"][0]:
            role = message["message"]["author"]["role"]
            content = message["message"]["content"]["parts"][0]
            print(f"{GREEN if role == 'user' else YELLOW}{role}: {RESET}{content}\n")


def main():
    with SyncChatGPT(
        # proxies=None,  # Optional proxies for network requests
        session_token=chat_session["token"],  # Use the session token for authentication
    ) as chatgpt:
        if chat_session["conversation_id"]:
            conversation = chatgpt.get_conversation(chat_session["conversation_id"])
        else:
            conversation = chatgpt.create_new_conversation()

        # Fetch and print the existing chat
        fetched_chat = conversation.fetch_chat()
        print_chat(fetched_chat)

        while True:
            user_input = input(f"{GREEN}user: {RESET}")
            chat_stream = conversation.chat(user_input)

            print_header = True
            for message in chat_stream:
                # The 'conversation_id' will be empty if it's a new chat, so we assign the new chat's ID
                if not chat_session["conversation_id"]:
                    chat_session["conversation_id"] = message["conversation_id"]

                # print header for the response
                if print_header:
                    print(f"\n{YELLOW}assistant: {RESET}", end="", flush=True)
                    print_header = False

                # Print the ChatGPT's reply
                print(message["content"], end="", flush=True)

            print("\n")

            # Write the new changes back to the config file
            with open("config.ini", "w") as file:
                config.write(file)


if __name__ == "__main__":
    main()

# Note: The 'conversation_id' will be found in the chat's url: 'https://chat.openai.com/c/conversation_id'
</file>

<file path="LICENSE">
Apache License
                           Version 2.0, January 2004
                        http://www.apache.org/licenses/

   TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION

   1. Definitions.

      "License" shall mean the terms and conditions for use, reproduction,
      and distribution as defined by Sections 1 through 9 of this document.

      "Licensor" shall mean the copyright owner or entity authorized by
      the copyright owner that is granting the License.

      "Legal Entity" shall mean the union of the acting entity and all
      other entities that control, are controlled by, or are under common
      control with that entity. For the purposes of this definition,
      "control" means (i) the power, direct or indirect, to cause the
      direction or management of such entity, whether by contract or
      otherwise, or (ii) ownership of fifty percent (50%) or more of the
      outstanding shares, or (iii) beneficial ownership of such entity.

      "You" (or "Your") shall mean an individual or Legal Entity
      exercising permissions granted by this License.

      "Source" form shall mean the preferred form for making modifications,
      including but not limited to software source code, documentation
      source, and configuration files.

      "Object" form shall mean any form resulting from mechanical
      transformation or translation of a Source form, including but
      not limited to compiled object code, generated documentation,
      and conversions to other media types.

      "Work" shall mean the work of authorship, whether in Source or
      Object form, made available under the License, as indicated by a
      copyright notice that is included in or attached to the work
      (an example is provided in the Appendix below).

      "Derivative Works" shall mean any work, whether in Source or Object
      form, that is based on (or derived from) the Work and for which the
      editorial revisions, annotations, elaborations, or other modifications
      represent, as a whole, an original work of authorship. For the purposes
      of this License, Derivative Works shall not include works that remain
      separable from, or merely link (or bind by name) to the interfaces of,
      the Work and Derivative Works thereof.

      "Contribution" shall mean any work of authorship, including
      the original version of the Work and any modifications or additions
      to that Work or Derivative Works thereof, that is intentionally
      submitted to Licensor for inclusion in the Work by the copyright owner
      or by an individual or Legal Entity authorized to submit on behalf of
      the copyright owner. For the purposes of this definition, "submitted"
      means any form of electronic, verbal, or written communication sent
      to the Licensor or its representatives, including but not limited to
      communication on electronic mailing lists, source code control systems,
      and issue tracking systems that are managed by, or on behalf of, the
      Licensor for the purpose of discussing and improving the Work, but
      excluding communication that is conspicuously marked or otherwise
      designated in writing by the copyright owner as "Not a Contribution."

      "Contributor" shall mean Licensor and any individual or Legal Entity
      on behalf of whom a Contribution has been received by Licensor and
      subsequently incorporated within the Work.

   2. Grant of Copyright License. Subject to the terms and conditions of
      this License, each Contributor hereby grants to You a perpetual,
      worldwide, non-exclusive, no-charge, royalty-free, irrevocable
      copyright license to reproduce, prepare Derivative Works of,
      publicly display, publicly perform, sublicense, and distribute the
      Work and such Derivative Works in Source or Object form.

   3. Grant of Patent License. Subject to the terms and conditions of
      this License, each Contributor hereby grants to You a perpetual,
      worldwide, non-exclusive, no-charge, royalty-free, irrevocable
      (except as stated in this section) patent license to make, have made,
      use, offer to sell, sell, import, and otherwise transfer the Work,
      where such license applies only to those patent claims licensable
      by such Contributor that are necessarily infringed by their
      Contribution(s) alone or by combination of their Contribution(s)
      with the Work to which such Contribution(s) was submitted. If You
      institute patent litigation against any entity (including a
      cross-claim or counterclaim in a lawsuit) alleging that the Work
      or a Contribution incorporated within the Work constitutes direct
      or contributory patent infringement, then any patent licenses
      granted to You under this License for that Work shall terminate
      as of the date such litigation is filed.

   4. Redistribution. You may reproduce and distribute copies of the
      Work or Derivative Works thereof in any medium, with or without
      modifications, and in Source or Object form, provided that You
      meet the following conditions:

      (a) You must give any other recipients of the Work or
          Derivative Works a copy of this License; and

      (b) You must cause any modified files to carry prominent notices
          stating that You changed the files; and

      (c) You must retain, in the Source form of any Derivative Works
          that You distribute, all copyright, patent, trademark, and
          attribution notices from the Source form of the Work,
          excluding those notices that do not pertain to any part of
          the Derivative Works; and

      (d) If the Work includes a "NOTICE" text file as part of its
          distribution, then any Derivative Works that You distribute must
          include a readable copy of the attribution notices contained
          within such NOTICE file, excluding those notices that do not
          pertain to any part of the Derivative Works, in at least one
          of the following places: within a NOTICE text file distributed
          as part of the Derivative Works; within the Source form or
          documentation, if provided along with the Derivative Works; or,
          within a display generated by the Derivative Works, if and
          wherever such third-party notices normally appear. The contents
          of the NOTICE file are for informational purposes only and
          do not modify the License. You may add Your own attribution
          notices within Derivative Works that You distribute, alongside
          or as an addendum to the NOTICE text from the Work, provided
          that such additional attribution notices cannot be construed
          as modifying the License.

      You may add Your own copyright statement to Your modifications and
      may provide additional or different license terms and conditions
      for use, reproduction, or distribution of Your modifications, or
      for any such Derivative Works as a whole, provided Your use,
      reproduction, and distribution of the Work otherwise complies with
      the conditions stated in this License.

   5. Submission of Contributions. Unless You explicitly state otherwise,
      any Contribution intentionally submitted for inclusion in the Work
      by You to the Licensor shall be under the terms and conditions of
      this License, without any additional terms or conditions.
      Notwithstanding the above, nothing herein shall supersede or modify
      the terms of any separate license agreement you may have executed
      with Licensor regarding such Contributions.

   6. Trademarks. This License does not grant permission to use the trade
      names, trademarks, service marks, or product names of the Licensor,
      except as required for reasonable and customary use in describing the
      origin of the Work and reproducing the content of the NOTICE file.

   7. Disclaimer of Warranty. Unless required by applicable law or
      agreed to in writing, Licensor provides the Work (and each
      Contributor provides its Contributions) on an "AS IS" BASIS,
      WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
      implied, including, without limitation, any warranties or conditions
      of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A
      PARTICULAR PURPOSE. You are solely responsible for determining the
      appropriateness of using or redistributing the Work and assume any
      risks associated with Your exercise of permissions under this License.

   8. Limitation of Liability. In no event and under no legal theory,
      whether in tort (including negligence), contract, or otherwise,
      unless required by applicable law (such as deliberate and grossly
      negligent acts) or agreed to in writing, shall any Contributor be
      liable to You for damages, including any direct, indirect, special,
      incidental, or consequential damages of any character arising as a
      result of this License or out of the use or inability to use the
      Work (including but not limited to damages for loss of goodwill,
      work stoppage, computer failure or malfunction, or any and all
      other commercial damages or losses), even if such Contributor
      has been advised of the possibility of such damages.

   9. Accepting Warranty or Additional Liability. While redistributing
      the Work or Derivative Works thereof, You may choose to offer,
      and charge a fee for, acceptance of support, warranty, indemnity,
      or other liability obligations and/or rights consistent with this
      License. However, in accepting such obligations, You may act only
      on Your own behalf and on Your sole responsibility, not on behalf
      of any other Contributor, and only if You agree to indemnify,
      defend, and hold each Contributor harmless for any liability
      incurred by, or claims asserted against, such Contributor by reason
      of your accepting any such warranty or additional liability.

   END OF TERMS AND CONDITIONS

   APPENDIX: How to apply the Apache License to your work.

      To apply the Apache License to your work, attach the following
      boilerplate notice, with the fields enclosed by brackets "[]"
      replaced with your own identifying information. (Don't include
      the brackets!)  The text should be enclosed in the appropriate
      comment syntax for the file format. We also recommend that a
      file or class name and description of purpose be included on the
      same "printed page" as the copyright notice for easier
      identification within third-party archives.

   Copyright [yyyy] [name of copyright owner]

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
</file>

<file path="re_gpt/__init__.py">
from .async_chatgpt import AsyncChatGPT
from .sync_chatgpt import SyncChatGPT
</file>

<file path="re_gpt/async_chatgpt.py">
import asyncio
import base64
import ctypes
import inspect
import json
import re
import uuid
from typing import AsyncGenerator, Callable, Optional

import websockets
from curl_cffi.requests import AsyncSession

from .errors import (
    BackendError,
    InvalidModelName,
    InvalidSessionToken,
    RetryError,
    TokenNotProvided,
    UnexpectedResponseError,
)
from .utils import async_get_binary_path, get_model_slug

# Constants
USER_AGENT = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"
CHATGPT_API = "https://chat.openai.com/backend-api/{}"
CHATGPT_FREE_API = "https://chat.openai.com/backend-anon/{}"
BACKUP_ARKOSE_TOKEN_GENERATOR = "https://arkose-token-generator.zaieem.repl.co/token"
WS_REGISTER_URL = CHATGPT_API.format("register-websocket")

MODELS = {
    "gpt-4": {"slug": "gpt-4", "needs_arkose_token": True},
    "gpt-3.5": {"slug": "text-davinci-002-render-sha", "needs_arkose_token": False},
}


class AsyncConversation:
    def __init__(self, chatgpt, conversation_id=None, model=None):
        self.chatgpt = chatgpt
        self.conversation_id = conversation_id
        self.parent_id = None
        self.model = model

    async def fetch_chat(self) -> dict:
        """
        Fetches the chat of the conversation from the API.

        Returns:
            dict: The JSON response from the API containing the chat if the conversation_id is not none, else returns an empty dict.

        Raises:
            UnexpectedResponseError: If the response is not a valid JSON object or if the response json is not in the expected format
        """
        if not self.conversation_id:
            return {}

        url = CHATGPT_API.format(f"conversation/{self.conversation_id}")
        response = await self.chatgpt.session.get(
            url=url, headers=self.chatgpt.build_request_headers()
        )

        error = None
        try:
            chat = response.json()
            self.parent_id = list(chat.get("mapping", {}))[-1]
            model_slug = get_model_slug(chat)
            self.model = [
                key for key, value in MODELS.items() if value["slug"] == model_slug
            ][0]
        except Exception as e:
            error = e
        if error is not None:
            raise UnexpectedResponseError(error, response.text)

        return chat

    async def chat(self, user_input: str) -> AsyncGenerator[dict, None]:
        """
        As the name implies, chat with ChatGPT.

        Args:
            user_input (str): The user's input message.

        Yields:
            dict: A dictionary representing assistant responses.

        Returns:
            AsyncGenerator[dict, None]: An asynchronous generator object that yields assistant responses.

        Raises:
            UnexpectedResponseError: If the response is not a valid JSON object or if the response json is not in the expected format
        """

        payload = await self.build_message_payload(user_input)

        server_response = (
            ""  # To store what the server returned for debugging in case of an error
        )
        error = None
        try:
            full_message = None
            while True:
                response = (
                    self.send_message(payload=payload)
                    if not self.chatgpt.websocket_mode
                    else self.send_websocket_message(payload=payload)
                )
                async for chunk in response:
                    decoded_chunk = (
                        chunk.decode() if isinstance(chunk, bytes) else chunk
                    )

                    server_response += decoded_chunk
                    for line in decoded_chunk.splitlines():
                        if not line.startswith("data: "):
                            continue

                        raw_json_data = line[6:]
                        if not (decoded_json := self.decode_raw_json(raw_json_data)):
                            continue

                        if (
                            "message" in decoded_json
                            and decoded_json["message"]["author"]["role"] == "assistant"
                        ):
                            processed_response = self.filter_response(decoded_json)
                            if full_message:
                                prev_resp_len = len(
                                    full_message["message"]["content"]["parts"][0]
                                )
                                processed_response["content"] = processed_response[
                                    "content"
                                ][prev_resp_len::]

                            yield processed_response
                            full_message = decoded_json
                self.conversation_id = full_message["conversation_id"]
                self.parent_id = full_message["message"]["id"]
                if (
                    full_message["message"]["metadata"]["finish_details"]["type"]
                    == "max_tokens"
                ):
                    payload = await self.build_message_continuation_payload()
                else:
                    break
        except Exception as e:
            error = e

        # raising the error outside the 'except' block to prevent the 'During handling of the above exception, another exception occurred' error
        if error is not None:
            raise UnexpectedResponseError(error, server_response)

    async def send_message(self, payload: dict) -> AsyncGenerator[bytes, None]:
        """
        Send a message payload to the server and receive the response.

        Args:
            payload (dict): Payload containing message information.

        Yields:
            bytes: Chunk of data received as a response.
        """
        response_queue = asyncio.Queue()

        async def perform_request():
            def content_callback(chunk):
                response_queue.put_nowait(chunk)

            url = CHATGPT_API.format("conversation")

            headers = self.chatgpt.build_request_headers()
            # Add Chat Requirements Token
            chat_requriments_token = await self.chatgpt.create_chat_requirements_token()
            if chat_requriments_token:
                headers[
                    "openai-sentinel-chat-requirements-token"
                ] = chat_requriments_token

            await self.chatgpt.session.post(
                url=url,
                headers=headers,
                json=payload,
                content_callback=content_callback,
            )
            await response_queue.put(None)

        asyncio.create_task(perform_request())

        while True:
            chunk = await response_queue.get()
            if chunk is None:
                break
            yield chunk

    async def send_websocket_message(self, payload: dict) -> AsyncGenerator[str, None]:
        """
        Send a message payload via WebSocket and receive the response.

        Args:
            payload (dict): Payload containing message information.

        Yields:
            str: Chunk of data received as a response.
        """
        await self.chatgpt.ensure_websocket()

        response_queue = asyncio.Queue()
        websocket_request_id = None

        async def perform_request():
            nonlocal websocket_request_id

            url = CHATGPT_API.format("conversation")
            headers = self.chatgpt.build_request_headers()
            # Add Chat Requirements Token
            chat_requriments_token = await self.chatgpt.create_chat_requirements_token()
            if chat_requriments_token:
                headers[
                    "openai-sentinel-chat-requirements-token"
                ] = chat_requriments_token

            response = (
                await self.chatgpt.session.post(
                    url=url,
                    headers=headers,
                    json=payload,
                )
            ).json()

            websocket_request_id = response.get("websocket_request_id")

            if websocket_request_id is None:
                raise UnexpectedResponseError(
                    "WebSocket request ID not found in response", response
                )

            if websocket_request_id not in self.chatgpt.ws_conversation_map:
                self.chatgpt.ws_conversation_map[websocket_request_id] = response_queue

        asyncio.create_task(perform_request())

        while True:
            chunk = await response_queue.get()
            if chunk is None:
                break
            yield chunk

        del self.chatgpt.ws_conversation_map[websocket_request_id]

    async def build_message_payload(self, user_input: str) -> dict:
        """
        Build a payload for sending a user message.

        Returns:
            dict: Payload containing message information.
        """
        if self.conversation_id and (self.parent_id is None or self.model is None):
            await self.fetch_chat()  # it will automatically fetch the chat and set the parent id

        payload = {
            "conversation_mode": {"conversation_mode": {"kind": "primary_assistant"}},
            "conversation_id": self.conversation_id,
            "action": "next",
            "arkose_token": await self.arkose_token_generator()
            if self.chatgpt.generate_arkose_token
            or MODELS[self.model]["needs_arkose_token"]
            else None,
            "force_paragen": False,
            "history_and_training_disabled": False,
            "messages": [
                {
                    "author": {"role": "user"},
                    "content": {"content_type": "text", "parts": [user_input]},
                    "id": str(uuid.uuid4()),
                    "metadata": {},
                }
            ],
            "model": MODELS[self.model]["slug"],
            "parent_message_id": str(uuid.uuid4())
            if not self.parent_id
            else self.parent_id,
            "websocket_request_id": str(uuid.uuid4())
            if self.chatgpt.websocket_mode
            else None,
        }

        return payload

    async def build_message_continuation_payload(self) -> dict:
        """
        Build a payload for continuing ChatGPT's cut off response.

        Returns:
            dict: Payload containing message information for continuation.
        """
        payload = {
            "conversation_mode": {"conversation_mode": {"kind": "primary_assistant"}},
            "action": "continue",
            "arkose_token": await self.arkose_token_generator()
            if self.chatgpt.generate_arkose_token
            or MODELS[self.model]["needs_arkose_token"]
            else None,
            "conversation_id": self.conversation_id,
            "force_paragen": False,
            "history_and_training_disabled": False,
            "model": MODELS[self.model]["slug"],
            "parent_message_id": self.parent_id,
            "timezone_offset_min": -300,
        }

        return payload

    async def arkose_token_generator(self) -> str:
        """
        Generate an Arkose token.

        Returns:
            str: Arkose token.
        """
        if not self.chatgpt.tried_downloading_binary:
            self.chatgpt.binary_path = await async_get_binary_path(self.chatgpt.session)

            if self.chatgpt.binary_path:
                self.chatgpt.arkose = ctypes.CDLL(self.chatgpt.binary_path)
                self.chatgpt.arkose.GetToken.restype = ctypes.c_char_p

            self.chatgpt.tried_downloading_binary = True

        if self.chatgpt.binary_path:
            try:
                result = self.chatgpt.arkose.GetToken()
                return ctypes.string_at(result).decode("utf-8")
            except:
                pass

        for _ in range(5):
            response = await self.chatgpt.session.get(BACKUP_ARKOSE_TOKEN_GENERATOR)
            if response.text == "null":
                raise BackendError(error_code=505)
            try:
                return response.json()["token"]
            except:
                await asyncio.sleep(0.7)

        raise RetryError(website=BACKUP_ARKOSE_TOKEN_GENERATOR)

    async def delete(self) -> None:
        """
        Deletes the conversation.
        """
        if self.conversation_id:
            await self.chatgpt.delete_conversation(self.conversation_id)

            self.conversation_id = None
            self.parent_id = None

    @staticmethod
    def decode_raw_json(raw_json_data: str) -> dict or bool:
        """
        Decode JSON.

        Args:
            raw_json_data (str): JSON as a string.

        Returns:
            dict: Decoded JSON.
        """
        try:
            decoded_json = json.loads(raw_json_data.strip())
            return decoded_json
        except:
            return False

    @staticmethod
    def filter_response(response):
        processed_response = {
            "content": response["message"]["content"]["parts"][0],
            "message_id": response["message"]["id"],
            "parent_id": response["message"]["metadata"]["parent_id"],
            "conversation_id": response["conversation_id"],
        }

        return processed_response


class AsyncChatGPT:
    def __init__(
        self,
        proxies: Optional[dict] = None,
        session_token: Optional[str] = None,
        exit_callback_function: Optional[Callable] = None,
        auth_token: Optional[str] = None,
        generate_arkose_token: Optional[bool] = False,
        websocket_mode: Optional[bool] = False,
    ):
        """
        Initializes an instance of the class.

        Args:
            proxies (Optional[dict]): A dictionary of proxy settings. Defaults to None.
            session_token (Optional[str]): A session token. Defaults to None.
            exit_callback_function (Optional[callable]): A function to be called on exit. Defaults to None.
            auth_token (Optional[str]): An authentication token. Defaults to None.
            generate_arkose_token (Optional[bool]): Toggle whether to generate and send arkose-token in the payload. Defaults to False.
            websocket_mode (Optional[bool]): Toggle whether to use WebSocket for chat. Defaults to False.
        """
        self.proxies = proxies
        self.exit_callback_function = exit_callback_function

        self.arkose = None
        self.binary_path = None
        self.tried_downloading_binary = False
        self.generate_arkose_token = generate_arkose_token

        self.session_token = session_token
        self.auth_token = auth_token
        self.session = None

        self.websocket_mode = websocket_mode
        self.ws_loop = None
        self.ws_conversation_map = {}
        
        # do not need session mode
        self.free_mode = True if self.session_token is None else False
        self.auth_cookie = None
        self.devive_id = str(uuid.uuid4())

    async def __aenter__(self):
        self.session = AsyncSession(
            impersonate="chrome110", timeout=99999, proxies=self.proxies
        )
        if self.generate_arkose_token:
            self.binary_path = await async_get_binary_path(self.session)

            if self.binary_path:
                self.arkose = ctypes.CDLL(self.binary_path)
                self.arkose.GetToken.restype = ctypes.c_char_p

            self.tried_downloading_binary = True

        if not self.auth_token:
            if not self.free_mode:
                if self.session_token is None:
                    raise TokenNotProvided
                self.auth_token = await self.fetch_auth_token()
            else:
                self.auth_cookie = await self.fetch_free_mode_cookies()

        if not self.websocket_mode:
            self.websocket_mode = await self.check_websocket_availability()

        if self.websocket_mode:
            await self.ensure_websocket()

        return self

    async def __aexit__(self, *_):
        try:
            if self.exit_callback_function and callable(self.exit_callback_function):
                if not inspect.iscoroutinefunction(self.exit_callback_function):
                    self.exit_callback_function(self)
        finally:
            if inspect.iscoroutinefunction(self.session.close):
                await self.session.close()
            else:
                self.session.close()

    def build_request_headers(self) -> dict:
        """
        Build headers for HTTP requests.

        Returns:
            dict: Request headers.
        """
        headers = {
            "User-Agent": USER_AGENT,
            "Accept": "text/event-stream",
            "Accept-Language": "en-US",
            "Accept-Encoding": "gzip, deflate, br",
            "Content-Type": "application/json",
            "Origin": "https://chat.openai.com",
            "Alt-Used": "chat.openai.com",
            "Connection": "keep-alive",
            "Oai-device-id": self.devive_id,
        }

        if self.free_mode:
            headers["Cookie"] = ';'.join([f"{key}={value}" for key, value in self.auth_cookie.items()])
        else:
            headers["Authorization"] = f"Bearer {self.auth_token}"

        return headers

    def get_conversation(self, conversation_id: str) -> AsyncConversation:
        """
        Makes an instance of class Conversation and return it.

        Args:
            conversation_id (str): The ID of the conversation to fetch.

        Returns:
            Conversation: Conversation object.
        """

        return AsyncConversation(self, conversation_id)

    def create_new_conversation(
        self, model: Optional[str] = "gpt-3.5"
    ) -> AsyncConversation:
        if model not in MODELS:
            raise InvalidModelName(model, MODELS)
        return AsyncConversation(self, model=model)

    async def delete_conversation(self, conversation_id: str) -> dict:
        """
        Delete a conversation.

        Args:
            conversation_id (str): Unique identifier for the conversation.

        Returns:
            dict: Server response json.
        """
        url = CHATGPT_API.format(f"conversation/{conversation_id}")
        response = await self.session.patch(
            url=url, headers=self.build_request_headers(), json={"is_visible": False}
        )

        return response.json()

    async def fetch_auth_token(self) -> str:
        """
        Fetch the authentication token for the session.

        Raises:
            InvalidSessionToken: If the session token is invalid.

        Returns: authentication token.
        """
        url = "https://chat.openai.com/api/auth/session"
        cookies = {"__Secure-next-auth.session-token": self.session_token}

        headers = {
            "User-Agent": USER_AGENT,
            "Accept": "*/*",
            "Accept-Language": "en-US,en;q=0.5",
            "Alt-Used": "chat.openai.com",
            "Connection": "keep-alive",
            "Sec-Fetch-Dest": "empty",
            "Sec-Fetch-Mode": "cors",
            "Sec-Fetch-Site": "same-origin",
            "Sec-GPC": "1",
            "Cookie": "; ".join(
                [
                    f"{cookie_key}={cookie_value}"
                    for cookie_key, cookie_value in cookies.items()
                ]
            ),
        }

        response = await self.session.get(url=url, headers=headers)
        response_json = response.json()

        if "accessToken" in response_json:
            return response_json["accessToken"]

        raise InvalidSessionToken

    async def set_custom_instructions(
        self,
        about_user: Optional[str] = "",
        about_model: Optional[str] = "",
        enable_for_new_chats: Optional[bool] = True,
    ) -> dict:
        """
        Set cuteom instructions for ChatGPT.

        Args:
            about_user (str): What would you like ChatGPT to know about you to provide better responses?
            about_model (str): How would you like ChatGPT to respond?
            enable_for_new_chats (bool): Enable for new chats.
        Returns:
            dict: Server response json.
        """
        data = {
            "about_user_message": about_user,
            "about_model_message": about_model,
            "enabled": enable_for_new_chats,
        }
        url = CHATGPT_API.format("user_system_messages")
        response = await self.session.post(
            url=url, headers=self.build_request_headers(), json=data
        )

        return response.json()

    async def retrieve_chats(
        self, offset: Optional[int] = 0, limit: Optional[int] = 28
    ) -> dict:
        params = {
            "offset": offset,
            "limit": limit,
            "order": "updated",
        }
        url = CHATGPT_API.format("conversations")
        response = await self.session.get(
            url=url, params=params, headers=self.build_request_headers()
        )

        return response.json()

    async def check_websocket_availability(self) -> bool:
        """
        Check if WebSocket is available.

        Returns:
            bool: True if WebSocket is available, otherwise False.
        """
        url = CHATGPT_API.format("accounts/check/v4-2023-04-27")

        headers = self.build_request_headers()
        
        raw_response = (await self.session.get(
            url=url, headers=headers
        ))
        try:
            response = raw_response.json()
            if 'account_ordering' in response and 'accounts' in response:
                account_id = response['account_ordering'][0]
                if account_id in response['accounts']:
                    return 'shared_websocket' in response['accounts'][account_id]['features']
        except:
            raise UnexpectedResponseError('Could not enable ws_mode', raw_response.text)

        return False

    async def ensure_websocket(self):
        if not self.ws_loop:
            ws_url_rsp = (
                await self.session.post(
                    WS_REGISTER_URL, headers=self.build_request_headers()
                )
            ).json()
            ws_url = ws_url_rsp["wss_url"]
            access_token = self.extract_access_token(ws_url)
            self.ws_loop = asyncio.create_task(
                self.listen_to_websocket(ws_url, access_token)
            )

    def extract_access_token(self, url):
        match = re.search(r"access_token=([^&]*)", url)
        if match:
            return match.group(1)
        else:
            return None

    async def listen_to_websocket(self, ws_url: str, access_token: str):
        headers = {"Authorization": f"Bearer {access_token}"}
        async with websockets.connect(ws_url, extra_headers=headers) as websocket:
            while True:
                message = await websocket.recv()
                message_data = json.loads(message)
                body_encoded = message_data.get("body", "")
                ws_id = message_data.get("websocket_request_id", "")
                decoded_body = base64.b64decode(body_encoded).decode("utf-8")
                response_queue = self.ws_conversation_map.get(ws_id)
                if response_queue is None:
                    continue
                if "title_generation" in decoded_body:
                    # skip
                    continue
                response_queue.put_nowait(decoded_body)
                if "[DONE]" in decoded_body or "[ERROR]" in decoded_body:
                    await response_queue.put(None)
                    continue

    async def create_chat_requirements_token(self):
        """
        Get a chat requirements token from chatgpt server

        Returns:
            str: chat requirements token
        """
        url = CHATGPT_API.format("sentinel/chat-requirements")

        if self.free_mode:
            url = CHATGPT_FREE_API.format("sentinel/chat-requirements")

        response = await self.session.post(
            url=url, headers=self.build_request_headers()
        )
        body = response.json()
        token = body.get("token", None)
        return token

    async def fetch_free_mode_cookies(self):
        home_url = "https://chat.openai.com/"
        headers = {
            "User-Agent": USER_AGENT,
            "Accept": "*/*",
            "Accept-Language": "en-US,en;q=0.5",
            "Alt-Used": "chat.openai.com",
            "Connection": "keep-alive",
            "Sec-Fetch-Dest": "empty",
            "Sec-Fetch-Mode": "cors",
            "Sec-Fetch-Site": "same-origin",
            "Sec-GPC": "1",
        }

        response = await self.session.get(url=home_url, headers=headers)
        response_cookies = response.cookies
        self.devive_id = response_cookies.get("oai-did")
        
        return response_cookies
</file>

<file path="re_gpt/errors.py">
class TokenNotProvided(Exception):
    def __init__(self):
        self.message = "Token not provided. Please pass your '__Secure-next-auth.session-token' as an argument (e.g., ChatGPT.init(session_token=YOUR_TOKEN))."
        super().__init__(self.message)


class InvalidSessionToken(Exception):
    def __init__(self):
        self.message = "Invalid session token provided."
        super().__init__(self.message)


class RetryError(Exception):
    def __init__(self, website, message="Exceeded maximum retries"):
        self.website = website
        self.message = f"{message} for website: {website}"
        super().__init__(self.message)


class BackendError(Exception):
    def __init__(self, error_code):
        self.error_code = error_code
        self.message = (
            f"An error occurred on the backend. Error code: {self.error_code}"
        )
        super().__init__(self.message)


class UnexpectedResponseError(Exception):
    def __init__(self, original_exception, server_response):
        self.original_exception = original_exception
        self.server_response = server_response
        self.message = f"An unexpected error occurred. Error message: {self.original_exception}.\nThis is what the server returned: {self.server_response}."
        super().__init__(self.message)


class InvalidModelName(Exception):
    def __init__(self, model, avalible_models):
        self.model = model
        self.avalible_models = avalible_models
        self.message = f'"{model}" is not a valid model. Avalible models: {[model for model in avalible_models]}'
        super().__init__(self.message)
</file>

<file path="re_gpt/sync_chatgpt.py">
import ctypes
import inspect
import time
import uuid
import websockets
from websockets.exceptions import ConnectionClosed
import json
import base64
import asyncio
from queue import Queue
from threading import Thread
from typing import Callable, Generator, Optional

from curl_cffi.requests import Session

from .async_chatgpt import (
    BACKUP_ARKOSE_TOKEN_GENERATOR,
    CHATGPT_API,
    USER_AGENT,
    AsyncChatGPT,
    AsyncConversation,
    MODELS,
    WS_REGISTER_URL,
    CHATGPT_FREE_API,
)
from .errors import (
    BackendError,
    InvalidSessionToken,
    RetryError,
    TokenNotProvided,
    UnexpectedResponseError,
    InvalidModelName,
)
from .utils import sync_get_binary_path, get_model_slug


class SyncConversation(AsyncConversation):
    def __init__(self, chatgpt, conversation_id: Optional[str] = None, model=None):
        super().__init__(chatgpt, conversation_id, model)

    def fetch_chat(self) -> dict:
        """
        Fetches the chat of the conversation from the API.

        Returns:
            dict: The JSON response from the API containing the chat if the conversation_id is not none, else returns an empty dict.

        Raises:
            UnexpectedResponseError: If the response is not a valid JSON object or if the response json is not in the expected format
        """
        if not self.conversation_id:
            return {}

        url = CHATGPT_API.format(f"conversation/{self.conversation_id}")
        response = self.chatgpt.session.get(
            url=url, headers=self.chatgpt.build_request_headers()
        )

        error = None
        try:
            chat = response.json()
            self.parent_id = list(chat.get("mapping", {}))[-1]
            model_slug = get_model_slug(chat)
            self.model = [
                key for key, value in MODELS.items() if value["slug"] == model_slug
            ][0]
        except Exception as e:
            error = e
        if error is not None:
            raise UnexpectedResponseError(error, response.text)

        return chat

    def chat(self, user_input: str) -> Generator[dict, None, None]:
        """
        As the name implies, chat with ChatGPT.

        Args:
            user_input (str): The user's input message.

        Yields:
            dict: A dictionary representing assistant responses.

        Returns:
            Generator[dict, None]: A generator object that yields assistant responses.

        Raises:
            UnexpectedResponseError: If the response is not a valid JSON object or if the response json is not in the expected format
        """

        payload = self.build_message_payload(user_input)

        server_response = (
            ""  # To store what the server returned for debugging in case of an error
        )
        error = None
        try:
            full_message = None
            while True:
                response = self.send_message(payload=payload) if not self.chatgpt.websocket_mode else self.send_websocket_message(payload=payload)
                for chunk in response:
                    decoded_chunk = chunk.decode() if not self.chatgpt.websocket_mode else chunk

                    server_response += decoded_chunk
                    for line in decoded_chunk.splitlines():
                        if not line.startswith("data: "):
                            continue

                        raw_json_data = line[6:]
                        if not (decoded_json := self.decode_raw_json(raw_json_data)):
                            continue

                        if (
                            "message" in decoded_json
                            and decoded_json["message"]["author"]["role"] == "assistant"
                        ):
                            processed_response = self.filter_response(decoded_json)
                            if full_message:
                                prev_resp_len = len(
                                    full_message["message"]["content"]["parts"][0]
                                )
                                processed_response["content"] = processed_response[
                                    "content"
                                ][prev_resp_len::]

                            yield processed_response
                            full_message = decoded_json
                self.conversation_id = full_message["conversation_id"]
                self.parent_id = full_message["message"]["id"]
                if (
                    full_message["message"]["metadata"]["finish_details"]["type"]
                    == "max_tokens"
                ):
                    payload = self.build_message_continuation_payload()
                else:
                    break
        except Exception as e:
            error = e

        # raising the error outside the 'except' block to prevent the 'During handling of the above exception, another exception occurred' error
        if error is not None:
            raise UnexpectedResponseError(error, server_response)

    def send_message(self, payload: dict) -> Generator[bytes, None, None]:
        """
        Send a message payload to the server and receive the response.

        Args:
            payload (dict): Payload containing message information.

        Yields:
            bytes: Chunk of data received as a response.
        """
        response_queue = Queue()

        def perform_request():
            def content_callback(chunk):
                response_queue.put(chunk)

            url = CHATGPT_API.format("conversation")
            headers = self.chatgpt.build_request_headers()
            # Add Chat Requirements Token
            chat_requriments_token = self.chatgpt.create_chat_requirements_token()
            if chat_requriments_token:
                headers["openai-sentinel-chat-requirements-token"] = chat_requriments_token

            response = self.chatgpt.session.post(
                url=url,
                headers=headers,
                json=payload,
                content_callback=content_callback,
            )
            response_queue.put(None)

        Thread(target=perform_request).start()

        while True:
            chunk = response_queue.get()
            if chunk is None:
                break
            yield chunk
    
    def send_websocket_message(self, payload: dict) -> Generator[str, None, None]:
        """
        Send a message payload via WebSocket and receive the response.

        Args:
            payload (dict): Payload containing message information.

        Yields:
            str: Chunk of data received as a response.
        """

        response_queue = Queue()
        websocket_request_id = None

        def perform_request():
            nonlocal websocket_request_id
            
            url = CHATGPT_API.format("conversation")
            headers = self.chatgpt.build_request_headers()
            # Add Chat Requirements Token
            chat_requriments_token = self.chatgpt.create_chat_requirements_token()
            if chat_requriments_token:
                headers["openai-sentinel-chat-requirements-token"] = chat_requriments_token

            response = (self.chatgpt.session.post(
                url=url,
                headers=headers,
                json=payload,
            )).json()

            websocket_request_id = response.get("websocket_request_id")
            
            if websocket_request_id is None:
                raise UnexpectedResponseError("WebSocket request ID not found in response", response)

            if websocket_request_id not in self.chatgpt.ws_conversation_map:
                self.chatgpt.ws_conversation_map[websocket_request_id] = response_queue
            
        Thread(target=perform_request).start()

        while True:
            chunk = response_queue.get()
            if chunk is None:
                break
            yield chunk

        del self.chatgpt.ws_conversation_map[websocket_request_id]

    def build_message_payload(self, user_input: str) -> dict:
        """
        Build a payload for sending a user message.

        Returns:
            dict: Payload containing message information.
        """
        if self.conversation_id and (self.parent_id is None or self.model is None):
            self.fetch_chat()  # it will automatically fetch the chat and set the parent id

        payload = {
            "conversation_mode": {"conversation_mode": {"kind": "primary_assistant"}},
            "conversation_id": self.conversation_id,
            "action": "next",
            "arkose_token": self.arkose_token_generator()
            if self.chatgpt.generate_arkose_token
            or MODELS[self.model]["needs_arkose_token"]
            else None,
            "force_paragen": False,
            "history_and_training_disabled": False,
            "messages": [
                {
                    "author": {"role": "user"},
                    "content": {"content_type": "text", "parts": [user_input]},
                    "id": str(uuid.uuid4()),
                    "metadata": {},
                }
            ],
            "model": MODELS[self.model]["slug"],
            "parent_message_id": str(uuid.uuid4())
            if not self.parent_id
            else self.parent_id,
            "websocket_request_id": str(uuid.uuid4())
            if self.chatgpt.websocket_mode
            else None,
        }

        return payload

    def build_message_continuation_payload(self) -> dict:
        """
        Build a payload for continuing ChatGPT's cut off response.

        Returns:
            dict: Payload containing message information for continuation.
        """
        payload = {
            "conversation_mode": {"conversation_mode": {"kind": "primary_assistant"}},
            "action": "continue",
            "arkose_token": self.arkose_token_generator()
            if self.chatgpt.generate_arkose_token
            or MODELS[self.model]["needs_arkose_token"]
            else None,
            "conversation_id": self.conversation_id,
            "force_paragen": False,
            "history_and_training_disabled": False,
            "model": MODELS[self.model]["slug"],
            "parent_message_id": self.parent_id,
            "timezone_offset_min": -300,
        }

        return payload

    def arkose_token_generator(self) -> str:
        """
        Generate an Arkose token.

        Returns:
            str: Arkose token.
        """
        if not self.chatgpt.tried_downloading_binary:
            self.chatgpt.binary_path = sync_get_binary_path(self.chatgpt.session)

            if self.chatgpt.binary_path:
                self.chatgpt.arkose = ctypes.CDLL(self.chatgpt.binary_path)
                self.chatgpt.arkose.GetToken.restype = ctypes.c_char_p

            self.chatgpt.tried_downloading_binary = True

        if self.chatgpt.binary_path:
            try:
                result = self.chatgpt.arkose.GetToken()
                return ctypes.string_at(result).decode("utf-8")
            except:
                pass

        for _ in range(5):
            response = self.chatgpt.session.get(BACKUP_ARKOSE_TOKEN_GENERATOR)
            if response.text == "null":
                raise BackendError(error_code=505)
            try:
                return response.json()["token"]
            except:
                time.sleep(0.7)

        raise RetryError(website=BACKUP_ARKOSE_TOKEN_GENERATOR)

    def delete(self) -> None:
        """
        Deletes the conversation.
        """
        if self.conversation_id:
            self.chatgpt.delete_conversation(self.conversation_id)

            self.conversation_id = None
            self.parent_id = None


class SyncChatGPT(AsyncChatGPT):
    def __init__(
        self,
        proxies: Optional[dict] = None,
        session_token: Optional[str] = None,
        exit_callback_function: Optional[Callable] = None,
        auth_token: Optional[str] = None,
        websocket_mode: Optional[bool] = False,
    ):
        """
        Initializes an instance of the class.

        Args:
            proxies (Optional[dict]): A dictionary of proxy settings. Defaults to None.
            session_token (Optional[str]): A session token. Defaults to None.
            exit_callback_function (Optional[callable]): A function to be called on exit. Defaults to None.
            auth_token (Optional[str]): An authentication token. Defaults to None.
            websocket_mode (Optional[bool]): Toggle whether to use WebSocket for chat. Defaults to False.
        """
        super().__init__(
            proxies=proxies,
            session_token=session_token,
            exit_callback_function=exit_callback_function,
            auth_token=auth_token,
            websocket_mode=websocket_mode,
        )

        self.stop_websocket_flag = False
        self.stop_websocket = None

    def __enter__(self):
        self.session = Session(
            impersonate="chrome110", timeout=99999, proxies=self.proxies
        )

        if self.generate_arkose_token:
            self.binary_path = sync_get_binary_path(self.session)

            if self.binary_path:
                self.arkose = ctypes.CDLL(self.binary_path)
                self.arkose.GetToken.restype = ctypes.c_char_p

            self.tried_downloading_binary = True

        if not self.auth_token:
            if not self.free_mode:
                if self.session_token is None:
                    raise TokenNotProvided
                self.auth_token = self.fetch_auth_token()
            else:
                self.auth_cookie = self.fetch_free_mode_cookies()
            
        # automaticly check the status of websocket_mode
        if not self.websocket_mode:
            self.websocket_mode = self.check_websocket_availability()
            
        if self.websocket_mode:
            def run_websocket():
                asyncio.run(self.ensure_websocket())
            self.ws_loop = Thread(target=run_websocket)
            self.ws_loop.start()

        return self

    def __exit__(self, *args):
        try:
            if self.exit_callback_function and callable(self.exit_callback_function):
                if not inspect.iscoroutinefunction(self.exit_callback_function):
                    self.exit_callback_function(self)
        finally:
            self.session.close()

        if self.websocket_mode:
            self.stop_websocket_flag = True
            self.ws_loop.join()

    def get_conversation(self, conversation_id: str) -> SyncConversation:
        """
        Makes an instance of class Conversation and return it.

        Args:
            conversation_id (str): The ID of the conversation to fetch.

        Returns:
            Conversation: Conversation object.
        """

        return SyncConversation(self, conversation_id)

    def create_new_conversation(
        self, model: Optional[str] = "gpt-3.5"
    ) -> SyncConversation:
        if model not in MODELS:
            raise InvalidModelName(model, MODELS)
        return SyncConversation(self, model=model)

    def delete_conversation(self, conversation_id: str) -> dict:
        """
        Delete a conversation.

        Args:
            conversation_id (str): Unique identifier for the conversation.

        Returns:
            dict: Server response json.
        """
        url = CHATGPT_API.format(f"conversation/{conversation_id}")
        response = self.session.patch(
            url=url, headers=self.build_request_headers(), json={"is_visible": False}
        )

        return response.json()

    def fetch_auth_token(self) -> str:
        """
        Fetch the authentication token for the session.

        Raises:
            InvalidSessionToken: If the session token is invalid.

        Returns: authentication token.
        """
        url = "https://chat.openai.com/api/auth/session"
        cookies = {"__Secure-next-auth.session-token": self.session_token}

        headers = {
            "User-Agent": USER_AGENT,
            "Accept": "*/*",
            "Accept-Language": "en-US,en;q=0.5",
            "Alt-Used": "chat.openai.com",
            "Connection": "keep-alive",
            "Sec-Fetch-Dest": "empty",
            "Sec-Fetch-Mode": "cors",
            "Sec-Fetch-Site": "same-origin",
            "Sec-GPC": "1",
            "Cookie": "; ".join(
                [
                    f"{cookie_key}={cookie_value}"
                    for cookie_key, cookie_value in cookies.items()
                ]
            ),
        }

        response = self.session.get(url=url, headers=headers)
        response_json = response.json()

        if "accessToken" in response_json:
            return response_json["accessToken"]

        raise InvalidSessionToken

    def set_custom_instructions(
        self,
        about_user: Optional[str] = "",
        about_model: Optional[str] = "",
        enable_for_new_chats: Optional[bool] = True,
    ) -> dict:
        """
        Set cuteom instructions for ChatGPT.

        Args:
            about_user (str): What would you like ChatGPT to know about you to provide better responses?
            about_model (str): How would you like ChatGPT to respond?
            enable_for_new_chats (bool): Enable for new chats.
        Returns:
            dict: Server response json.
        """
        data = {
            "about_user_message": about_user,
            "about_model_message": about_model,
            "enabled": enable_for_new_chats,
        }
        url = CHATGPT_API.format("user_system_messages")
        response = self.session.post(
            url=url, headers=self.build_request_headers(), json=data
        )

        return response.json()

    def retrieve_chats(
        self, offset: Optional[int] = 0, limit: Optional[int] = 28
    ) -> dict:
        params = {
            "offset": offset,
            "limit": limit,
            "order": "updated",
        }
        url = CHATGPT_API.format("conversations")
        response = self.session.get(
            url=url, params=params, headers=self.build_request_headers()
        )

        return response.json()
    
    def check_websocket_availability(self) -> bool:
        """
        Check if WebSocket is available.

        Returns:
            bool: True if WebSocket is available, otherwise False.
        """
        url = CHATGPT_API.format("accounts/check/v4-2023-04-27")

        headers = self.build_request_headers()
        
        raw_response = self.session.get(
            url=url, headers=headers
        )
        try:
            response = raw_response.json()
            if 'account_ordering' in response and 'accounts' in response:
                account_id = response['account_ordering'][0]
                if account_id in response['accounts']:
                    return 'shared_websocket' in response['accounts'][account_id]['features']
        except:
            raise UnexpectedResponseError('Could not enable ws_mode', raw_response.text)
    
    async def ensure_websocket(self):
        ws_url_rsp = self.session.post(WS_REGISTER_URL, headers=self.build_request_headers()).json()
        ws_url = ws_url_rsp['wss_url']
        access_token = self.extract_access_token(ws_url)
        asyncio.create_task(self.ensure_close_websocket())
        await self.listen_to_websocket(ws_url, access_token)
        
    async def ensure_close_websocket(self):
        while True:
            if self.stop_websocket_flag:
                break
            await asyncio.sleep(1)
        await self.stop_websocket()

    async def listen_to_websocket(self, ws_url: str, access_token: str):
        headers = {'Authorization': f'Bearer {access_token}'}
        async with websockets.connect(ws_url, extra_headers=headers) as websocket:
            async def stop_websocket():
                await websocket.close()
            self.stop_websocket = stop_websocket

            while True:
                message = None
                try:
                    message = await websocket.recv()
                except ConnectionClosed:
                    break
                message_data = json.loads(message)
                body_encoded = message_data.get("body", "")
                ws_id = message_data.get("websocket_request_id", "")
                decoded_body = base64.b64decode(body_encoded).decode('utf-8')
                response_queue = self.ws_conversation_map.get(ws_id)
                if response_queue is None:
                    continue
                response_queue.put_nowait(decoded_body)
                if '[DONE]' in decoded_body or '[ERROR]' in decoded_body:
                    response_queue.put(None)
                    continue

    def create_chat_requirements_token(self):
        """
        Get a chat requirements token from chatgpt server

        Returns:
            str: chat requirements token
        """
        url = CHATGPT_API.format("sentinel/chat-requirements")
        
        if self.free_mode:
            url = CHATGPT_FREE_API.format("sentinel/chat-requirements")
        
        response = self.session.post(
            url=url, headers=self.build_request_headers()
        )
        body = response.json()
        token = body.get("token", None)
        return token

    def fetch_free_mode_cookies(self):
        home_url = "https://chat.openai.com/"
        headers = {
            "User-Agent": USER_AGENT,
            "Accept": "*/*",
            "Accept-Language": "en-US,en;q=0.5",
            "Alt-Used": "chat.openai.com",
            "Connection": "keep-alive",
            "Sec-Fetch-Dest": "empty",
            "Sec-Fetch-Mode": "cors",
            "Sec-Fetch-Site": "same-origin",
            "Sec-GPC": "1",
        }

        response = self.session.get(url=home_url, headers=headers)
        response_cookies = response.cookies
        self.devive_id = response_cookies.get("oai-did")

        return response_cookies
</file>

<file path="re_gpt/utils.py">
import hashlib
import os
import platform

current_os = platform.system()
current_file_directory = "/".join(
    __file__.split("\\" if current_os == "Windows" else "/")[0:-1]
)

funcaptcha_bin_folder_path = f"{current_file_directory}/funcaptcha_bin"
latest_release_url = (
    "https://api.github.com/repos/Zai-Kun/reverse-engineered-chatgpt/releases"
)

binary_file_name = {"Windows": "windows_arkose.dll", "Linux": "linux_arkose.so"}.get(
    current_os
)

binary_path = {
    "Windows": f"{funcaptcha_bin_folder_path}/{binary_file_name}",
    "Linux": f"{funcaptcha_bin_folder_path}/{binary_file_name}",
}.get(current_os)


def calculate_file_md5(file_path):
    with open(file_path, "rb") as file:
        file_content = file.read()
        md5_hash = hashlib.md5(file_content).hexdigest()
        return md5_hash


def get_file_url(json_data):
    for release in json_data:
        if release["tag_name"].startswith("funcaptcha_bin"):
            file_url = next(
                asset["browser_download_url"]
                for asset in release["assets"]
                if asset["name"] == binary_file_name
            )
            return file_url


# async
async def async_download_binary(session, output_path, file_url):
    with open(output_path, "wb") as output_file:
        response = await session.get(
            url=file_url, content_callback=lambda chunk: output_file.write(chunk)
        )


async def async_get_binary_path(session):
    if binary_path is None:
        return None

    if not os.path.exists(funcaptcha_bin_folder_path) or not os.path.isdir(
        funcaptcha_bin_folder_path
    ):
        os.mkdir(funcaptcha_bin_folder_path)

    if os.path.isfile(binary_path):
        try:
            local_binary_hash = calculate_file_md5(binary_path)
            response = await session.get(latest_release_url)
            json_data = response.json()

            for line in json_data["body"].splitlines():
                if line.startswith(current_os):
                    latest_binary_hash = line.split("=")[-1]
                    break

            if local_binary_hash != latest_binary_hash:
                file_url = get_file_url(json_data)

                await async_download_binary(session, binary_path, file_url)
        except:
            return binary_path
    else:
        response = await session.get(latest_release_url)
        json_data = response.json()
        file_url = get_file_url(json_data)

        await async_download_binary(session, binary_path, file_url)

    return binary_path


# sync
def sync_download_binary(session, output_path, file_url):
    with open(output_path, "wb") as output_file:
        response = session.get(
            url=file_url, content_callback=lambda chunk: output_file.write(chunk)
        )


def sync_get_binary_path(session):
    if binary_path is None:
        return None

    if not os.path.exists(funcaptcha_bin_folder_path) or not os.path.isdir(
        funcaptcha_bin_folder_path
    ):
        os.mkdir(funcaptcha_bin_folder_path)

    if os.path.isfile(binary_path):
        try:
            local_binary_hash = calculate_file_md5(binary_path)
            response = session.get(latest_release_url)
            json_data = response.json()

            for line in json_data["body"].splitlines():
                if line.startswith(current_os):
                    latest_binary_hash = line.split("=")[-1]
                    break

            if local_binary_hash != latest_binary_hash:
                file_url = get_file_url(json_data)

                sync_download_binary(session, binary_path, file_url)
        except:
            return binary_path
    else:
        response = session.get(latest_release_url)
        json_data = response.json()
        file_url = get_file_url(json_data)

        sync_download_binary(session, binary_path, file_url)

    return binary_path


def get_model_slug(chat):
    for _, message in chat.get("mapping", {}).items():
        if "message" in message:
            if message["message"]:
                role = message["message"]["author"]["role"]
                if role == "assistant":
                    return message["message"]["metadata"]["model_slug"]
</file>

<file path="requirements.txt">
curl_cffi==0.5.9
websockets==12.0certifi==2024.2.2
cffi==1.16.0
curl_cffi==0.6.2
pycparser==2.22
websockets==12.0
</file>

<file path="sampleconfig.ini">
[session]
token = __Secure-next-auth.session-token here
conversation_id =
</file>

<file path="setup.py">
from setuptools import find_packages, setup

setup(
    name="re_gpt",
    version="4.0.0",
    author="Zai-Kun",
    description="Unofficial reverse-engineered ChatGPT API in Python.",
    long_description=open("README.md").read(),
    long_description_content_type="text/markdown",
    url="https://github.com/Zai-Kun/reverse-engineered-chatgpt",
    project_urls={
        "Bug Tracker": "https://github.com/Zai-Kun/reverse-engineered-chatgpt/issues",
    },
    packages=find_packages(),
    install_requires=["curl_cffi==0.5.9", "websockets==12.0"],
)
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
cover/

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
.pybuilder/
target/

# Jupyter Notebook
.ipynb_checkpoints

# IPython
profile_default/
ipython_config.py

# pyenv
#   For a library or package, you might want to ignore these files since the code is
#   intended to run in multiple environments; otherwise, check them in:
# .python-version

# pipenv
#   According to pypa/pipenv#598, it is recommended to include Pipfile.lock in version control.
#   However, in case of collaboration, if having platform-specific dependencies or dependencies
#   having no cross-platform support, pipenv may install dependencies that don't work, or not
#   install all needed dependencies.
#Pipfile.lock

# poetry
#   Similar to Pipfile.lock, it is generally recommended to include poetry.lock in version control.
#   This is especially recommended for binary packages to ensure reproducibility, and is more
#   commonly ignored for libraries.
#   https://python-poetry.org/docs/basic-usage/#commit-your-poetrylock-file-to-version-control
#poetry.lock

# pdm
#   Similar to Pipfile.lock, it is generally recommended to include pdm.lock in version control.
#pdm.lock
#   pdm stores project-wide configurations in .pdm.toml, but it is recommended to not include it
#   in version control.
#   https://pdm.fming.dev/#use-with-ide
.pdm.toml

# PEP 582; used by e.g. github.com/David-OConnor/pyflow and github.com/pdm-project/pdm
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

# pytype static type analyzer
.pytype/

# Cython debug symbols
cython_debug/

# PyCharm
#  JetBrains specific template is maintained in a separate JetBrains.gitignore that can
#  be found at https://github.com/github/gitignore/blob/main/Global/JetBrains.gitignore
#  and can be added to the global gitignore or merged into this file.  For a more nuclear
#  option (not recommended) you can uncomment the following to ignore the entire idea folder.
#.idea/

data
funcaptcha_bin
.vscode
test/
config.ini
test.py
*.sublime-project
*.sublime-workspace
</file>

<file path="README.md">
<div align="center">
  <a href="https://github.com/Zai-Kun/reverse-engineered-chatgpt">  </a>

<h1 align="center">Reverse Engineered <a href="https://openai.com/blog/chatgpt">ChatGPT</a> API</h1>

  <p align="center">
    Use OpenAI ChatGPT in your Python code without an API key

[![Stargazers][stars-badge]][stars-url]
[![Forks][forks-badge]][forks-url]
[![Discussions][discussions-badge]][discussions-url]
[![Issues][issues-badge]][issues-url]
[![MIT License][license-badge]][license-url]

  English | [简体中文](./docs/zh-README.md)

  </p>
    <p align="center">
    <a href="https://github.com/Zai-Kun/reverse-engineered-chatgpt"></a>
    <a href="https://github.com/Zai-Kun/reverse-engineered-chatgpt/issues">Report Bug</a>
    |
    <a href="https://github.com/Zai-Kun/reverse-engineered-chatgpt/discussions">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#inspiration">Inspiration</a></li>
        <li><a href="#how-it-works">How it works</a></li>
        <li><a href="#built-using">Built using</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
        <li><a href="#obtaining-session-token">Obtaining Session Token</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a>
        <ul>
        <li><a href="#basic-example">Example Usage</a></li>
        </ul>
    </li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

## About The Project

This project can be used to integrate OpenAI's ChatGPT services into your python code. You can use this project to prompt ChatGPT for responses directly from python, without using an official API key.

This can be useful if you want to use ChatGPT API without a [ChatGPT Plus](https://openai.com/blog/chatgpt-plus) account.

### Inspiration

ChatGPT has an official API which can be used to interface your Python code to it, but it needs to be used with an API key. This API key can only be obtained if you have a [ChatGPT Plus](https://openai.com/blog/chatgpt-plus) account, which requires $20/month (as of 05/11/2023). But you can use ChatGPT for free, using the [ChatGPT web interface](https://chat.openai.com/). This project aims to interface your code to ChatGPT web version so you can use ChatGPT in your Python code without using an API key.

### How it works

[ChatGPT](https://chat.openai.com/) web interface's requests have been reverse engineered, and directly integrated into Python requests. Hence, any requests made using this script is a simulated as a request made by a user directly on the website. Hence, it is free and needs no API key.

### Built Using

- [![Python][python-badge]][python-url]

## Getting Started

### Prerequisites

- Python >= 3.9

### Installation

```sh
pip install re-gpt
```

## Usage

### Basic example

```python
from re_gpt import SyncChatGPT

session_token = "__Secure-next-auth.session-token here"
conversation_id = None # conversation ID here


with SyncChatGPT(session_token=session_token) as chatgpt:
    prompt = input("Enter your prompt: ")

    if conversation_id:
        conversation = chatgpt.get_conversation(conversation_id)
    else:
        conversation = chatgpt.create_new_conversation()

    for message in conversation.chat(prompt):
        print(message["content"], flush=True, end="")

```

### Basic async example

```python
import asyncio
import sys

from re_gpt import AsyncChatGPT

session_token = "__Secure-next-auth.session-token here"
conversation_id = conversation_id = None # conversation ID here

if sys.version_info >= (3, 8) and sys.platform.lower().startswith("win"):
    asyncio.set_event_loop_policy(asyncio.WindowsSelectorEventLoopPolicy())


async def main():
    async with AsyncChatGPT(session_token=session_token) as chatgpt:
        prompt = input("Enter your prompt: ")

        if conversation_id:
            conversation = chatgpt.get_conversation(conversation_id)
        else:
            conversation = chatgpt.create_new_conversation()

        async for message in conversation.chat(prompt):
            print(message["content"], flush=True, end="")


if __name__ == "__main__":
    asyncio.run(main())
```

## More Examples

For a more complex example, check out the [examples](/examples) folder in the repository.

### Obtaining The Session Token

1. Go to <https://chat.openai.com/chat> and log in or sign up.
2. Open the developer tools in your browser.
3. Go to the `Application` tab and open the `Cookies` section.
4. Copy the value for `__Secure-next-auth.session-token` and save it.

## TODO

- [x] Add more examples
- [ ] Add better error handling
- [x] Implement a function to retrieve all ChatGPT chats
- [ ] Improve documentation

## Contributing

Contributions are what makes the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request.
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the Apache License 2.0. See [`LICENSE`](https://github.com/Zai-Kun/reverse-engineered-chatgpt/blob/main/LICENSE) for more information.

## Contact/Bug report

Zai-Kun - [Discord Server](https://discord.gg/ymcqxudVJG)

Repo Link: <https://github.com/Zai-Kun/reverse-engineered-chatgpt>

## Acknowledgments

- [sudoAlphaX (for writing this readme)](https://github.com/sudoAlphaX)

- [yifeikong (curl-cffi module)](https://github.com/yifeikong/curl_cffi)

- [acheong08 (implementation to obtain arkose_token)](https://github.com/acheong08/funcaptcha)

- [pyca (cryptography module)](https://github.com/pyca/cryptography/)

- [Legrandin (pycryptodome module)](https://github.com/Legrandin/pycryptodome/)

- [othneildrew (README Template)](https://github.com/othneildrew)

<!-- MARKDOWN LINKS & IMAGES -->

[forks-badge]: https://img.shields.io/github/forks/Zai-Kun/reverse-engineered-chatgpt
[forks-url]: https://github.com/Zai-Kun/reverse-engineered-chatgpt/network/members
[stars-badge]: https://img.shields.io/github/stars/Zai-Kun/reverse-engineered-chatgpt
[stars-url]: https://github.com/Zai-Kun/reverse-engineered-chatgpt/stargazers
[issues-badge]: https://img.shields.io/github/issues/Zai-Kun/reverse-engineered-chatgpt
[issues-url]: https://github.com/Zai-Kun/reverse-engineered-chatgpt/issues
[discussions-badge]: https://img.shields.io/github/discussions/Zai-Kun/reverse-engineered-chatgpt
[discussions-url]: https://github.com/Zai-Kun/reverse-engineered-chatgpt/discussions
[python-badge]: https://img.shields.io/badge/Python-blue?logo=python&logoColor=yellow
[python-url]: https://www.python.org/
[license-badge]: https://img.shields.io/github/license/Zai-Kun/reverse-engineered-chatgpt
[license-url]: https://github.com/Zai-Kun/reverse-engineered-chatgpt/blob/main/LICENSE
</file>

</files>
