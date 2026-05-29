# Packages Digest

This file contains the packed representation, metadata, and directory trees of all downloaded reverse-engineering package archives and their extracted source folders.

<file_summary>
<purpose>
Provides a consolidated view of all downloaded reverse-engineering libraries in PyPI and npm, mapping their structure and tracking versions.
</purpose>
</file_summary>

## Directory Structure

```text
- sources/packages/pypi/Pandora_ChatGPT-1.3.5 [Directory]
- sources/packages/pypi/Pandora_ChatGPT-1.3.5-py3-none-any.whl [Archive]
- sources/packages/pypi/chatgpt_wrapper-0.0.1 [Directory]
- sources/packages/pypi/chatgpt_wrapper-0.0.1-py3-none-any.whl [Archive]
- sources/packages/pypi/gpt4free-0.0.0 [Directory]
- sources/packages/pypi/gpt4free-0.0.0-py3-none-any.whl [Archive]
- sources/packages/pypi/re_gpt-4.0.0 [Directory]
- sources/packages/pypi/re_gpt-4.0.0-py3-none-any.whl [Archive]
- sources/packages/pypi/re_gpt-4.0.0.tar.gz [Archive]
- sources/packages/pypi/revChatGPT-6.8.6 [Directory]
- sources/packages/pypi/revChatGPT-6.8.6-py3-none-any.whl [Archive]
- sources/packages/pypi/revChatGPT-6.8.6.tar.gz [Archive]
- sources/packages/npm/chatgpt [Directory]
- sources/packages/npm/chatgpt-5.2.5 [Directory]
- sources/packages/npm/chatgpt-5.2.5.tgz [Archive]
```

## Extracted Packages Details


### PyPI Package: `Pandora_ChatGPT-1.3.5`
* **Type:** Extracted Source Directory
* **Path:** `sources/packages/pypi/Pandora_ChatGPT-1.3.5`
* **Directory Tree Structure:**
```text
Pandora_ChatGPT-1.3.5/
├── Pandora_ChatGPT-1.3.5.dist-info
│   ├── LICENSE
│   ├── METADATA
│   ├── RECORD
│   ├── WHEEL
│   ├── entry_points.txt
│   └── top_level.txt
└── pandora
    ├── __init__.py
    ├── __main__.py
    ├── bots
    │   ├── __init__.py
    │   ├── legacy.py
    │   └── server.py
    ├── cloud_launcher.py
    ├── exts
    │   ├── __init__.py
    │   ├── config.py
    │   ├── hooks.py
    │   └── token.py
    ├── flask
    │   ├── static
    │   │   ├── _next
    │   │   │   └── static
    │   │   │       ├── chunks
    │   │   │       │   ├── 113-23682f80a24dd00d.js
    │   │   │       │   ├── 14-0cb0d20affbd720d.js
    │   │   │       │   ├── 174-bd28069f281ef76f.js
    │   │   │       │   ├── 1f110208-44a6f43ddc5e9011.js
    │   │   │       │   ├── 264-13e92c51b0315184.js
    │   │   │       │   ├── 360-442b869f1ba4bb1b.js
    │   │   │       │   ├── 424-d1d3bfe6a3ca6c4a.js
    │   │   │       │   ├── 554.9b8bfd0762461d74.js
    │   │   │       │   ├── 68a27ff6-1185184b61bc22d0.js
    │   │   │       │   ├── 762-222df1028c0c1555.js
    │   │   │       │   ├── 949.1a6eb804b5e91f61.js
    │   │   │       │   ├── bd26816a-981e1ddc27b37cc6.js
    │   │   │       │   ├── framework-7a789ee31d2a7534.js
    │   │   │       │   ├── main-149b337e061b4d04.js
    │   │   │       │   ├── pages
    │   │   │       │   │   ├── _app-aaa11de1926dcafe.js
    │   │   │       │   │   ├── _error-786d27d84962122a.js
    │   │   │       │   │   └── chat
    │   │   │       │   │       └── [[...chatId]]-76751174916fa3f8.js
    │   │   │       │   ├── polyfills-c67a75d1b6f99dc8.js
    │   │   │       │   └── webpack-c9a868e8e0796ec6.js
    │   │   │       ├── css
    │   │   │       │   └── ac221fb2caad35a6.css
    │   │   │       └── olf4sv64FWIcQ_zCGl90t
    │   │   │           ├── _buildManifest.js
    │   │   │           └── _ssgManifest.js
    │   │   ├── apple-touch-icon.png
    │   │   ├── favicon-16x16.png
    │   │   ├── favicon-32x32.png
    │   │   ├── fonts
    │   │   │   ├── KaTeX_Caligraphic-Bold.woff
    │   │   │   ├── KaTeX_Caligraphic-Regular.woff
    │   │   │   ├── KaTeX_Fraktur-Bold.woff
    │   │   │   ├── KaTeX_Fraktur-Regular.woff
    │   │   │   ├── KaTeX_Main-Bold.woff
    │   │   │   ├── KaTeX_Main-BoldItalic.woff
    │   │   │   ├── KaTeX_Main-Italic.woff
    │   │   │   ├── KaTeX_Main-Regular.woff
    │   │   │   ├── KaTeX_Math-BoldItalic.woff
    │   │   │   ├── KaTeX_Math-Italic.woff
    │   │   │   ├── KaTeX_SansSerif-Bold.woff
    │   │   │   ├── KaTeX_SansSerif-Italic.woff
    │   │   │   ├── KaTeX_SansSerif-Regular.woff
    │   │   │   ├── KaTeX_Script-Regular.woff
    │   │   │   ├── KaTeX_Size1-Regular.woff
    │   │   │   ├── KaTeX_Size2-Regular.woff
    │   │   │   ├── KaTeX_Size3-Regular.woff
    │   │   │   ├── KaTeX_Size4-Regular.woff
    │   │   │   ├── KaTeX_Typewriter-Regular.woff
    │   │   │   ├── Signifier-Regular.otf
    │   │   │   ├── Sohne-Buch.otf
    │   │   │   ├── Sohne-Halbfett.otf
    │   │   │   ├── SohneMono-Buch.otf
    │   │   │   └── SohneMono-Halbfett.otf
    │   │   └── images
    │   │       └── 2022
    │   │           └── 11
    │   │               └── ChatGPT.jpg
    │   └── templates
    │       └── chat.html
    ├── launcher.py
    ├── migrations
    │   ├── __init__.py
    │   ├── database.py
    │   ├── migrate.py
    │   ├── models.py
    │   └── scripts
    │       └── 20230308_01_7ctOr.sql
    ├── openai
    │   ├── __init__.py
    │   ├── api.py
    │   ├── auth.py
    │   ├── token.py
    │   └── utils.py
    ├── py.typed
    └── turbo
        ├── __init__.py
        ├── base.py
        └── chat.py
```


### PyPI Package: `chatgpt_wrapper-0.0.1`
* **Type:** Extracted Source Directory
* **Path:** `sources/packages/pypi/chatgpt_wrapper-0.0.1`
* **Directory Tree Structure:**
```text
chatgpt_wrapper-0.0.1/
├── chatgpt
│   ├── __init__.py
│   └── chatgpt.py
└── chatgpt_wrapper-0.0.1.dist-info
    ├── LICENSE
    ├── METADATA
    ├── RECORD
    ├── WHEEL
    └── top_level.txt
```


### PyPI Package: `gpt4free-0.0.0`
* **Type:** Extracted Source Directory
* **Path:** `sources/packages/pypi/gpt4free-0.0.0`
* **Directory Tree Structure:**
```text
gpt4free-0.0.0/
├── g4f
│   ├── Provider
│   │   ├── AItianhu.py
│   │   ├── AiAsk.py
│   │   ├── AiChatOnline.py
│   │   ├── Aura.py
│   │   ├── Bestim.py
│   │   ├── Bing.py
│   │   ├── ChatAnywhere.py
│   │   ├── ChatBase.py
│   │   ├── ChatForAi.py
│   │   ├── Chatgpt4Online.py
│   │   ├── ChatgptAi.py
│   │   ├── ChatgptDemo.py
│   │   ├── ChatgptDemoAi.py
│   │   ├── ChatgptFree.py
│   │   ├── ChatgptLogin.py
│   │   ├── ChatgptNext.py
│   │   ├── ChatgptX.py
│   │   ├── Chatxyz.py
│   │   ├── DeepInfra.py
│   │   ├── FakeGpt.py
│   │   ├── FreeChatgpt.py
│   │   ├── FreeGpt.py
│   │   ├── GPTalk.py
│   │   ├── GeekGpt.py
│   │   ├── GeminiProChat.py
│   │   ├── Gpt6.py
│   │   ├── GptChatly.py
│   │   ├── GptForLove.py
│   │   ├── GptGo.py
│   │   └── GptGod.py
│   │   └── ... (truncated)
│   ├── __init__.py
│   ├── api
│   │   ├── __init__.py
│   │   ├── _logging.py
│   │   ├── _tokenizer.py
│   │   └── run.py
│   ├── base_provider.py
│   ├── cli.py
│   ├── debug.py
│   ├── defaults.py
│   ├── errors.py
│   ├── gui
│   │   ├── __init__.py
│   │   ├── client
│   │   │   ├── css
│   │   │   │   └── style.css
│   │   │   ├── html
│   │   │   │   └── index.html
│   │   │   ├── img
│   │   │   │   ├── android-chrome-192x192.png
│   │   │   │   ├── android-chrome-512x512.png
│   │   │   │   ├── apple-touch-icon.png
│   │   │   │   ├── favicon-16x16.png
│   │   │   │   ├── favicon-32x32.png
│   │   │   │   ├── gpt.png
│   │   │   │   ├── site.webmanifest
│   │   │   │   └── user.png
│   │   │   └── js
│   │   │       ├── chat.v1.js
│   │   │       ├── highlight.min.js
│   │   │       ├── highlightjs-copy.min.js
│   │   │       └── icons.js
│   │   ├── run.py
│   │   └── server
│   │       ├── app.py
│   │       ├── backend.py
│   │       ├── config.py
│   │       ├── internet.py
│   │       └── website.py
│   ├── image.py
│   ├── models.py
│   ├── requests.py
│   ├── requests_aiohttp.py
│   ├── requests_curl_cffi.py
│   ├── typing.py
│   ├── version.py
│   └── webdriver.py
└── gpt4free-0.0.0.dist-info
    ├── LICENSE
    ├── METADATA
    ├── RECORD
    ├── WHEEL
    ├── entry_points.txt
    └── top_level.txt
```


### PyPI Package: `re_gpt-4.0.0`
* **Type:** Extracted Source Directory
* **Path:** `sources/packages/pypi/re_gpt-4.0.0`
* **Directory Tree Structure:**
```text
re_gpt-4.0.0/
├── LICENSE
├── PKG-INFO
├── README.md
├── re_gpt
│   ├── __init__.py
│   ├── async_chatgpt.py
│   ├── errors.py
│   ├── sync_chatgpt.py
│   └── utils.py
├── re_gpt-4.0.0
│   ├── LICENSE
│   ├── PKG-INFO
│   ├── README.md
│   ├── re_gpt
│   │   ├── __init__.py
│   │   ├── async_chatgpt.py
│   │   ├── errors.py
│   │   ├── sync_chatgpt.py
│   │   └── utils.py
│   ├── setup.cfg
│   └── setup.py
├── setup.cfg
└── setup.py
```


### PyPI Package: `revChatGPT-6.8.6`
* **Type:** Extracted Source Directory
* **Path:** `sources/packages/pypi/revChatGPT-6.8.6`
* **Directory Tree Structure:**
```text
revChatGPT-6.8.6/
├── LICENSE
├── PKG-INFO
├── README.md
├── revChatGPT-6.8.6
│   ├── LICENSE
│   ├── PKG-INFO
│   ├── README.md
│   ├── setup.cfg
│   ├── setup.py
│   ├── src
│   │   └── revChatGPT
│   │       ├── V1.py
│   │       ├── V3.py
│   │       ├── __init__.py
│   │       ├── __main__.py
│   │       ├── config
│   │       │   └── enable_internet.json
│   │       ├── typings.py
│   │       ├── utils.py
│   │       └── version.py
│   └── tests
│       └── test_recipient.py
├── setup.cfg
├── setup.py
├── src
│   └── revChatGPT
│       ├── V1.py
│       ├── V3.py
│       ├── __init__.py
│       ├── __main__.py
│       ├── config
│       │   └── enable_internet.json
│       ├── typings.py
│       ├── utils.py
│       └── version.py
└── tests
    └── test_recipient.py
```


### npm Package: `chatgpt`
* **Type:** Extracted Source Directory
* **Path:** `sources/packages/npm/chatgpt`
* **Directory Tree Structure:**
```text
chatgpt/
├── bin
│   └── cli.js
├── build
│   ├── index.d.ts
│   ├── index.js
│   └── index.js.map
├── license
├── package.json
└── readme.md
```


### npm Package: `chatgpt-5.2.5`
* **Type:** Extracted Source Directory
* **Path:** `sources/packages/npm/chatgpt-5.2.5`
* **Directory Tree Structure:**
```text
chatgpt-5.2.5/
└── package
    ├── bin
    │   └── cli.js
    ├── build
    │   ├── index.d.ts
    │   ├── index.js
    │   └── index.js.map
    ├── license
    ├── package.json
    └── readme.md
```

