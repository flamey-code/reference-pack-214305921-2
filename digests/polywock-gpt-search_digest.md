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
.babelrc
.gitignore
advancedSearch.md
package.json
postcss.config.js
PRIVACY_POLICY.md
README.md
src/background/index.ts
src/background/init.ts
src/background/requests.ts
src/common/searchStyle.css
src/comps/svgs.tsx
src/declare.d.ts
src/defaults.ts
src/globalVar.ts
src/helper.ts
src/hooks/useStorage.tsx
src/main.ts
src/mainLoader.ts
src/options/App.tsx
src/options/comps/ColorWell.css
src/options/comps/ColorWell.tsx
src/options/comps/FloatTooltip.css
src/options/comps/FloatTooltip.tsx
src/options/comps/NumericInput.css
src/options/comps/NumericInput.tsx
src/options/comps/Option.css
src/options/comps/Option.tsx
src/options/comps/Reset.css
src/options/comps/Reset.tsx
src/options/comps/Toggle.css
src/options/comps/Toggle.tsx
src/options/index.tsx
src/options/styles.css
src/options/utils.ts
src/preamble/index.ts
src/preamble/style.css
src/raccoon/App.tsx
src/raccoon/comps/CleverDiv.tsx
src/raccoon/comps/LoadMore.tsx
src/raccoon/comps/ResultItem.tsx
src/raccoon/hooks/useAutoBar.tsx
src/raccoon/hooks/useClickBlur.tsx
src/raccoon/hooks/useResize.tsx
src/raccoon/index.tsx
src/raccoon/searchChats/core.ts
src/raccoon/searchChats/extractContext.ts
src/raccoon/searchChats/extractOpts.ts
src/raccoon/searchChats/filterChats.ts
src/raccoon/searchChats/Grabby.ts
src/raccoon/searchChats/index.ts
src/raccoon/searchChats/preFilter.ts
src/raccoon/style.css
src/raccoon/types.ts
src/raccoon/utils/bump.tsx
src/raccoon/utils/extractChats.ts
src/raccoon/utils/fetchChats.ts
src/raccoon/utils/getElapsed.tsx
src/raccoon/utils/getTtl.ts
src/raccoon/utils/gizmo.ts
src/raccoon/utils/misc.tsx
src/raccoon/utils/rawTypes.ts
src/types.ts
src/utils/browser.ts
src/utils/getKnown.ts
src/utils/gsm.ts
src/utils/GsmType.ts
src/utils/state.ts
static/_locales/en/messages.json
static/_locales/es/messages.json
static/_locales/it/messages.json
static/_locales/ja/messages.json
static/_locales/ko/messages.json
static/_locales/pt_BR/messages.json
static/_locales/ru/messages.json
static/_locales/tr/messages.json
static/_locales/vi/messages.json
static/_locales/zh_CN/messages.json
static/_locales/zh_TW/messages.json
static/128.png
static/locales/en.json
static/locales/es.json
static/locales/it.json
static/locales/ja.json
static/locales/ko.json
static/locales/pt_BR.json
static/locales/ru.json
static/locales/tr.json
static/locales/vi.json
static/locales/zh_CN.json
static/locales/zh_TW.json
static/options.html
staticCh/manifest.json
staticFf/manifest.json
tools/generateGsmType.js
tools/validateLocale.js
tsconfig.json
webpack.config.js
</directory_structure>

<files>
This section contains the contents of the repository's files.

<file path=".babelrc">
{
  "presets": [
      "@babel/env",
      "@babel/typescript",
      ["@babel/react", {
        "runtime": "automatic"
      }]
  ]
}
</file>

<file path="advancedSearch.md">
## Advanced search for searching chat history.

### Flags 
**+dalle**: Filters for chats that used DALL-E.

**+python**: Filters for chats that used data analysis.  

**+browse**: Filters for chats that used web browsing.   

**+gizmo**: Filters for chats that used a custom GPT.  

**+gizmos**: Filters for chats that used multiple GPTs.   

**+gpt4**: Filter for chats that used GPT-4. 

**+archived**: Filter for archived chats. 

Adding a minus sign (-) before any of these flags, e.g., -dalle, filters out chats that used the specified tool or feature.

#### Example  
Search for chats with "spatula" that used DALL-E, but not a custom GPT. 
```text
spatula +dalle -gizmo  
```

Search for hammer, but in archives. 
```text
hammer +archived 
```

### Other Flags
**+turns \<num\>**: Filter for chats with more than \<num\> messages.   

**-turns \<num\>**: Filter for chats with less than \<num\> messages. 

#### Example  
Search for long chats with word "docker".
```text
docker +turns 20
```

Search for short chats with word ketchup. 
```text
ketchup -turns 5 
```


### Date Flags:
**+c** YYYY/MM/DD: Filters for chats created after the specified date.  

**-c** YYYY/MM/DD: Filters for chats created before the specified date.  

**+u** YYYY/MM/DD: Filters for chats updated after the specified date.  

**-u** YYYY/MM/DD: Filters for chats updated before the specified date.  

#### Example  
Search for chats with "javascript" created in 2023. 
```text
javascript +c 2023/01/01 -c 2024/01/01 
```

### Source flags 
**+title**: Check matches only from chat's title.  

**+body**: Check matches only from chat's body. You can also specify **+ast** to only include ChatGPT's replies, and **+user** to only include your replies.    

+**gpt**: Search only GPT titles. If you just installed GPT Search, this might not work as no GPTs have been cached. 

Using these flags with a minus sign (-) will exclude them them instead. 

Search for chats with "spatula" but exclude ChatGPT's replies. 
```text
spatula -ast  
```

Search for chats with "africa" in title. 

```
+title africa 
```

Search Consensus gpt. 

```
+gpt consensus 
```

### Search chats by custom GPT.

Find all chats that used a specific GPT. This feature might not work if you recently installed GPT Search as no GPTs might be cached yet. 
```
+g write for me 
```

Find all chats that used a specific GPT created before a specific date.
```
+g write for me -c 2024/1/31
```


### Exclude keywords  
You can exclude keywords by including a minus sign. 

Search for chats that contain "rabbit", but not "lion".

```
rabbit -lion 
```

### Chaining
You can filter search results by another query with the && operator. 

Search for chats with "king" in title and "troll" in the body. 
```text
+title king && +body troll
```


Search for chats that used Consensus GPT, but only those that have 'troll' in the title. 
```text
+g consensus && +title troll
```
</file>

<file path="package.json">
{
  "name": "gbar",
  "version": "1.0.0",
  "description": "aa",
  "main": "index.js",
  "sideEffects": [
    "./src/background/*.ts"
  ],
  "browserslist": [
    "chrome >= 112",
    "firefox >= 112"
  ],
  "scripts": {
    "build:common": "  rm -rf build   &&              webpack --config webpack.config.js && cp -r static/. staticCh/. build/unpacked   && find build   -name '.DS_Store' -type f -delete",
    "build:commonFf": "rm -rf buildFf && FIREFOX=true webpack --config webpack.config.js && cp -r static/. staticFf/. buildFf/unpacked && find buildFf -name '.DS_Store' -type f -delete",
    "build:dev": "    export NODE_ENV=development && npm run build:common",
    "build:devFf": "  export NODE_ENV=development && npm run build:commonFf",
    "build:prod": "   export NODE_ENV=production && npm run build:common   && cd build/unpacked   && zip -r ../packed.zip .",
    "build:prodFf": " export NODE_ENV=production && npm run build:commonFf && cd buildFf/unpacked && zip -r ../packed.zip .",
    "build:prodFff": "npm run build:prodFf && rm -rf zed && mkdir -p zed/unpacked && cp -r src static staticCh staticFf tools .babelrc package-lock.json package.json postcss.config.js tsconfig.json webpack.config.js zed/unpacked && (cd zed/unpacked && zip -r ../packed.zip .) && mv zed/packed.zip buildFf/source.zip && rm -rf zed"
  },
  "author": "",
  "dependencies": {
    "@babel/core": "^7.22.10",
    "@babel/preset-env": "^7.23.9",
    "@babel/preset-react": "^7.23.3",
    "@babel/preset-typescript": "^7.22.5",
    "@types/chrome": "^0.0.243",
    "@types/lodash.debounce": "^4.0.9",
    "@types/node": "^20.11.17",
    "@types/react": "^18.2.55",
    "@types/react-dom": "^18.2.19",
    "babel-loader": "^9.1.3",
    "clsx": "^2.1.0",
    "css-loader": "^6.10.0",
    "cssnano": "^6.0.3",
    "escape-string-regexp": "^5.0.0",
    "immer": "^10.0.3",
    "lodash.debounce": "^4.0.8",
    "postcss": "^8.4.35",
    "postcss-loader": "^8.1.0",
    "react": "^18.2.0",
    "react-dom": "^18.2.0",
    "react-icons": "^5.0.1",
    "style-loader": "^3.3.4",
    "terser-webpack-plugin": "^5.3.10",
    "typescript": "^5.1.6",
    "webpack": "^5.88.2",
    "webpack-cli": "^5.1.4"
  }
}
</file>

<file path="postcss.config.js">
module.exports = {
    plugins: [
      require('cssnano')({
        preset: 'default',
    }),
    ]
  }
</file>

<file path="PRIVACY_POLICY.md">
"GPT Search" does not collect any personal information, including chat data. This document may later change.

### Data safety

Chats are cached locally to allow for fast searches. If another person has access to your computer, they might be able to look at your cached chats even if you're not signed in to ChatGPT. 

To remove cached chats, uninstall the extension or "Reset cache" in the options page. If you're using a shared computer, it's recommended to enable 'Session-only caching' in the options page.
</file>

<file path="src/background/index.ts">
import "./init"
import "./requests"

import { loadGsm } from "../utils/gsm"
import { ensureMigrated } from "./init"
import { AnyDict } from "../types"


// Prep is for preamble stuff.
async function getPrep() {
    const ph = await getSearchPlaceholder()
    return {ph: ph || null}
}

async function getSearchPlaceholder(): Promise<string> {
    const v = (await chrome.storage.local.get('o:ph'))['o:ph']
    if (v) return v;
    return (await loadGsm()).search
}


chrome.action?.onClicked.addListener(tab => {
    chrome.tabs.create({url: chrome.runtime.getURL('options.html')})
})


chrome.runtime.onMessage.addListener((msg, sender, reply) => {
    if (msg.type === "LOAD_RACCOON") {
        chrome.scripting.executeScript({
            target: {tabId: sender.tab.id, allFrames: false},
            files: ["./raccoon.js"]
        })
        reply(true)
    } else if (msg.type === "REQUEST_GSM") {
        loadGsm().then(gsm => {
            chrome.storage.local.set({'o:ph': gsm.search})
            reply(gsm)
        }, err => reply(null))
        return true 
    } else if (msg.type === "PREP") {
        getPrep().then(prep => {
            reply(prep)
        })
        return true 
    } else if (msg.type === "OPEN_LINK") {
        chrome.tabs.create({url: msg.url, active: msg.active, index: sender.tab.index + 1})
    } else if (msg.type === "RESET") {
        reset().then(() => reply({}), error => reply({error}))
        return true 
    } else if (msg.type === "REQUEST_CREATE_TAB") {
        chrome.tabs.create({
            url: msg.url,
            index: sender.tab.index + 1,
            active: msg.active 
        })
        reply(true)
    } else if (msg.type === "GET_SESSION_ITEM") {
        chrome.storage.session.get(msg.keys as chrome.storage.StorageGet).then(ok => {
            reply({ok})
        }, error => reply({error}))
        return true 
    } else if (msg.type === "SET_SESSION_ITEM") {
        chrome.storage.session.set(msg.items as AnyDict).then(() => {
            reply({})
        }, error => reply({error}))
        return true 
    } 
}) 

async function reset() {
    await chrome.storage.local.clear()
    await chrome.storage.session.clear()
    await ensureMigrated()
}
</file>

<file path="src/background/init.ts">
import { CURRENT_VERSION, generateConfig } from "../defaults"
import { CONFIG_KEYS, Config } from "../types"

chrome.runtime.onInstalled.addListener(async () => {
    chrome.storage.session.setAccessLevel?.({accessLevel: 'TRUSTED_AND_UNTRUSTED_CONTEXTS'})
    ensureMigrated()
})

export async function ensureMigrated() {
    let config = await chrome.storage.local.get(CONFIG_KEYS)
    config = await migrateConfig(config)
    await chrome.storage.local.set(config)
}

async function migrateConfig(config: Partial<Config>) {
    const version = config["g:version"]
    if (version === 1) {
        config = config 
    }
    if (config["g:version"] !== CURRENT_VERSION) {
        return generateConfig()
    }
    return config 
}

async function migrateOneToTwo(config: Partial<Config>): Promise<Partial<Config>> {
    return config 
}
</file>

<file path="src/background/requests.ts">
import { Gizmo, TempState } from "../types"
import { getLocalItem, setLocal } from "../utils/getKnown"

chrome.webRequest.onBeforeSendHeaders.addListener(deets => {
    processGizmoUrl(deets.url)

    const auth = deets.requestHeaders?.find(r => r.name === "Authorization")?.value
    if (!auth) return 
    chrome.storage.session.set({'s:auth': auth})
    getLocalItem('g:sessionOnly').then(sessionOnly => {
        if (!sessionOnly) setLocal({'o:auth': auth}) 
    })
}, {urls: [
    'https://chatgpt.com/backend-api/*'
], types: ['xmlhttprequest']}, ['requestHeaders'])


const gizmoRegex = /gizmos\/g\-([a-zA-Z0-9]+)/
let gizmos: TempState["o:gizmos"]

async function processGizmoUrl(url: string) {
    const gizmoId = gizmoRegex.exec(url)?.[1]
    if (!gizmoId) return 

    const stk = `s:g:${gizmoId}`
    if ((await chrome.storage.session.get(stk))[stk]) return 
    chrome.storage.session.set({[stk]: true})
    const res = await fetch(`https://chatgpt.com/public-api/gizmos/g-${gizmoId}`)
    if (!res.ok) return 
    const json = await res.json()
    const displayName = json?.gizmo?.display?.name 
    if (!displayName) return 
    const imageUrl = json.gizmo.display.profile_picture_url 
    if (!imageUrl) return 

    if (!gizmos) gizmos = await getLocalItem('o:gizmos') ?? {}
    gizmos[gizmoId] = {
        id: gizmoId,
        name: displayName,
        added: Date.now(),
        imageUrl
    }
    
    setLocal({'o:gizmos': gizmos})
}
</file>

<file path="src/common/searchStyle.css">
.wrapper {
    margin-top: 10px;
    margin-right: 9px;
    margin-left: 10px;
    
    &.new {
        margin: 10px;
        margin-top: 15px;
        margin-bottom: 0px;
    }
}

.search {
    font-size: 16px;
    display: grid;
    position: relative;
    align-items: center;
    background-color: transparent;
    color: var(--text-color-primary);
    position: relative;
    -webkit-font-smoothing: antialiased;

    & > .searchIcon {
        position: absolute;
        left: 10px;
        pointer-events: none;
    }

    & > input {
        width: 100%;
        box-sizing: border-box;
        -webkit-font-smoothing: antialiased;
        background-color: var(--bg-color);
        /* font-family: "Segoe UI", "Avenir", Courier, monospace; */
        padding: 8px;
        display: block;
        padding-left: 34px;
        padding-right: 24px;
        border-radius: 10px;
        border: 1px solid var(--border-medium);
        font-size: inherit;
        background-color: inherit;
        color: var(--text-color-primary);
        cursor: text;
        
        &:focus {
            outline: none;
            /* border: 1px solid var(--context-color); */
            border: 1px solid var(--border-xheavy);
        }
    }
}
</file>

<file path="src/comps/svgs.tsx">
import * as React from "react"

type SvgPropsBase = {
    width?: React.SVGAttributes<SVGElement>["width"],
    height?: React.SVGAttributes<SVGElement>["height"],
    style?: React.SVGAttributes<SVGElement>["style"],
    className?: React.SVGAttributes<SVGElement>["className"],
    color?: React.SVGAttributes<SVGElement>["color"]
}

export type SvgProps = SvgPropsBase & {
    size?: number | string,
    onClick?: (e: React.MouseEvent) => void 
}

function prepareProps(props: SvgProps) {
    props = { ...(props ?? {}) }
    props.width = props.width ?? props.size ?? "1em"
    props.height = props.height ?? props.size ?? "1em"

    delete props.size
    return props as SvgPropsBase
}


export function Gear(props: SvgProps) {
    return (
        <svg
            xmlns="http://www.w3.org/2000/svg"
            height="1em"
            fill="currentColor"
            stroke="currentColor"
            strokeWidth={0}
            viewBox="0 0 14 16"
            {...(prepareProps(props))}
        >
            <path
                fillRule="evenodd"
                stroke="none"
                d="M14 8.77v-1.6l-1.94-.64-.45-1.09.88-1.84-1.13-1.13-1.81.91-1.09-.45-.69-1.92h-1.6l-.63 1.94-1.11.45-1.84-.88-1.13 1.13.91 1.81-.45 1.09L0 7.23v1.59l1.94.64.45 1.09-.88 1.84 1.13 1.13 1.81-.91 1.09.45.69 1.92h1.59l.63-1.94 1.11-.45 1.84.88 1.13-1.13-.92-1.81.47-1.09L14 8.75v.02zM7 11c-1.66 0-3-1.34-3-3s1.34-3 3-3 3 1.34 3 3-1.34 3-3 3z"
            />
        </svg>
    )
}

export function Github(props: SvgProps) {
    return (
        <svg
            xmlns="http://www.w3.org/2000/svg"
            height="1em"
            fill="currentColor"
            stroke="currentColor"
            strokeWidth={0}
            viewBox="0 0 16 16"
            {...props}
        >
            <path
                fillRule="evenodd"
                stroke="none"
                d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0 0 16 8c0-4.42-3.58-8-8-8z"
            />
        </svg>
    )
}

export function Star(props: SvgProps) {
    return (
        <svg
            xmlns="http://www.w3.org/2000/svg"
            height="1em"
            fill="currentColor"
            stroke="currentColor"
            strokeWidth={0}
            viewBox="0 0 576 512"
            {...props}
        >
            <path
                stroke="none"
                d="M259.3 17.8 194 150.2 47.9 171.5c-26.2 3.8-36.7 36.1-17.7 54.6l105.7 103-25 145.5c-4.5 26.3 23.2 46 46.4 33.7L288 439.6l130.7 68.7c23.2 12.2 50.9-7.4 46.4-33.7l-25-145.5 105.7-103c19-18.5 8.5-50.8-17.7-54.6L382 150.2 316.7 17.8c-11.7-23.6-45.6-23.9-57.4 0z"
            />
        </svg>
    )
}

export function Heart(props: SvgProps) {
    return (
        <svg
            xmlns="http://www.w3.org/2000/svg"
            height="1em"
            fill="currentColor"
            stroke="currentColor"
            strokeWidth={0}
            viewBox="0 0 512 512"
            {...props}
        >
            <path
                stroke="none"
                d="M462.3 62.6C407.5 15.9 326 24.3 275.7 76.2L256 96.5l-19.7-20.3C186.1 24.3 104.5 15.9 49.7 62.6c-62.8 53.6-66.1 149.8-9.9 207.9l193.5 199.8c12.5 12.9 32.8 12.9 45.3 0l193.5-199.8c56.3-58.1 53-154.3-9.8-207.9z"
            />
        </svg>
    )
}

export function Pin(props: SvgProps) {
    return (
        <svg
            xmlns="http://www.w3.org/2000/svg"
            width="1em"
            height="1em"
            fill="currentColor"
            stroke="currentColor"
            strokeWidth={0}
            viewBox="0 0 16 16"
            {...props}
        >
            <path
            stroke="none"
            d="M4.146.146A.5.5 0 0 1 4.5 0h7a.5.5 0 0 1 .5.5c0 .68-.342 1.174-.646 1.479-.126.125-.25.224-.354.298v4.431l.078.048c.203.127.476.314.751.555C12.36 7.775 13 8.527 13 9.5a.5.5 0 0 1-.5.5h-4v4.5c0 .276-.224 1.5-.5 1.5s-.5-1.224-.5-1.5V10h-4a.5.5 0 0 1-.5-.5c0-.973.64-1.725 1.17-2.189A5.921 5.921 0 0 1 5 6.708V2.277a2.77 2.77 0 0 1-.354-.298C4.342 1.674 4 1.179 4 .5a.5.5 0 0 1 .146-.354z"
            />
        </svg>
    )
}

export function Diamond(props: SvgProps) {
    return (
        <svg
            xmlns="http://www.w3.org/2000/svg"
            width="1em"
            height="1em"
            fill="currentColor"
            stroke="currentColor"
            strokeWidth={0}
            viewBox="0 0 512 512"
            {...props}
        >
            <path
            stroke="none"
            d="M284.3 11.7c-15.6-15.6-40.9-15.6-56.6 0l-216 216c-15.6 15.6-15.6 40.9 0 56.6l216 216c15.6 15.6 40.9 15.6 56.6 0l216-216c15.6-15.6 15.6-40.9 0-56.6l-216-216z"
            />
        </svg>
    )
}


export function Close(props: SvgProps) {
    return (
        <svg
            xmlns="http://www.w3.org/2000/svg"
            width="1em"
            height="1em"
            fill="currentColor"
            stroke="currentColor"
            strokeWidth={0}
            viewBox="0 0 24 24"
            {...props}
        >
            <path fill="none" stroke="none" d="M0 0h24v24H0z" />
            <path
            stroke="none"
            d="M19 6.41 17.59 5 12 10.59 6.41 5 5 6.41 10.59 12 5 17.59 6.41 19 12 13.41 17.59 19 19 17.59 13.41 12z"
            />
        </svg>
    )
}

export function SearchSvg(props: SvgProps) {
    return (
        <svg
            xmlns="http://www.w3.org/2000/svg"
            width='1em'
            height='1em'
            fill="none"
            stroke="currentColor"
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth={2}
            viewBox="0 0 24 24"
            {...props}
        >
            <circle cx={11} cy={11} r={8} />
            <path d="m21 21-4.35-4.35" />
        </svg>
    )
}

export function ResetSvg(props: SvgProps) {
    return (
        <svg
            xmlns="http://www.w3.org/2000/svg"
            width="1em"
            height="1em"
            fill="currentColor"
            stroke="currentColor"
            strokeWidth={0}
            className="reset active"
            viewBox="0 0 512 512"
            {...props}
        >
            <path
            stroke="none"
            d="M248.91 50a205.9 205.9 0 0 1 35.857 3.13c85.207 15.025 152.077 81.895 167.102 167.102 15.023 85.208-24.944 170.917-99.874 214.178-32.782 18.927-69.254 27.996-105.463 27.553-46.555-.57-92.675-16.865-129.957-48.15l30.855-36.768a157.846 157.846 0 0 0 180.566 15.797 157.846 157.846 0 0 0 76.603-164.274A157.848 157.848 0 0 0 276.429 100.4a157.84 157.84 0 0 0-139.17 43.862L185 192H57V64l46.34 46.342C141.758 71.962 194.17 50.03 248.91 50z"
            />
        </svg>
    )
}
</file>

<file path="src/declare.d.ts">
declare module '*.css?raw' {
    const content: string;
    export default content;
}

declare namespace chrome.storage {
    export type StorageChanges = {[key: string]: chrome.storage.StorageChange}
    
    export type StorageKeysArgument = string | string[] | {[key: string]: any} | null | undefined

    export type StorageGet = string | string[] | { [key: string]: any }; 
}
</file>

<file path="src/defaults.ts">
import { Config } from "./types";
import { loadGsm } from "./utils/gsm";

export const CURRENT_VERSION = 1 


export async function generateConfig() {
    gvar.gsm = await loadGsm()
    return ({
        'g:version': CURRENT_VERSION
    }) satisfies Partial<Config>
}
</file>

<file path="src/globalVar.ts">
export const gvar = ((globalThis.document ?? globalThis) as any).gvar ?? {}
;((globalThis.document ?? globalThis) as any).gvar  = gvar
</file>

<file path="src/helper.ts">
export function timeout(ms: number): Promise<void> {
    return new Promise((res, rej) => setTimeout(() => res(), ms))
}

export class QueueLock {
  public promiseChain = Promise.resolve()
  private chainLength = 0 

  wait(ms: number) {
    const oldChain = this.promiseChain
    this.promiseChain = oldChain.then(() => timeout(ms))
    
    if (this.chainLength++ > 100) {
      this.promiseChain = Promise.resolve()
      this.chainLength = 0 
    }
    return oldChain
  }
}

export function round(value: number, precision: number): number {
	const scalar = 10 ** precision
	return Math.round(value * scalar) / scalar
}

export function clamp(min: number, max: number, value: number) {
    let clamped = value 
    if (min != null) {
      clamped = Math.max(min, clamped)
    }
    if (max != null) {
      clamped = Math.min(max, clamped)
    }
    return clamped 
}

export function randomId() {
  return Math.ceil(Math.random() * 1E10).toString()
}

export function shuffle<T>(arr: T[]): T {
  return arr.at(Math.floor(Math.random() * arr.length))
}

export function assertType<T>(value: any): asserts value is T { }

export function createElement(text: string) {
  const temp = document.createElement('div')
  temp.innerHTML = text
  return temp.firstElementChild
}
</file>

<file path="src/hooks/useStorage.tsx">
import { useCallback, useEffect, useRef, useState } from "react";
import { SubscribeStorageKeys } from "../utils/state";
import { LocalState, AnyDict } from "../types";

type Env = {
    client?: SubscribeStorageKeys
}

export function useStorage(keys: string[], wait?: number, maxWait?: number) {
    const [items, _setItems] = useState(null as AnyDict)
    const env = useRef({} as Env).current

    useEffect(() => {
        env.client = new SubscribeStorageKeys(keys, true, _setItems, wait, maxWait)

        return () => {
            env.client?.release()
            delete env.client
        }
    }, [])

    const setItems = useCallback(async (view: AnyDict) => {
        return env.client?.push(view)
    }, [])


    return [items, setItems]
}

export function useKnownKeys(keys: (keyof LocalState)[], wait?: number, maxWait?: number): [Partial<LocalState>, (v: Partial<LocalState>) => void] {
    return useStorage(keys, wait, maxWait) as [Partial<LocalState>, (v: Partial<LocalState>) => void]
}
</file>

<file path="src/main.ts">
let blockScrollId: number 

window.addEventListener('busbusab', (e: CustomEvent) => {
    const deets = JSON.parse(e.detail)
    if (deets.type === 'NAV') {
        if ((window as any).next?.router?.push) {
            (window as any).next.router.push(deets.path)
        } else if ((window as any).__reactRouterDataRouter?.navigate) {
            (window as any).__reactRouterDataRouter.navigate(deets.path)
        } else if ((window as any).__remixRouter?.navigate) {
            (window as any).__remixRouter.navigate(deets.path)
        } else {
            window.dispatchEvent(new CustomEvent('rusrusar', {detail: JSON.stringify({type: 'NO_PUSH', path: deets.path}), bubbles: false}))
        }
    }  else if (deets.type === "BLOCK_SCROLL") {
    }
    e.stopImmediatePropagation()
}, {capture: true})


function shimScroll2() {
    let originalDesc = Object.getOwnPropertyDescriptor(Element.prototype, "scrollTop")
    Object.defineProperty(Element.prototype, "scrollTop", {set: function(value: number) {
        window.dispatchEvent(new CustomEvent("rusrusar", {detail: JSON.stringify({type: "SCROLL_SET"})}))
        return originalDesc.set.call(this, value)
    }, get: originalDesc.get, configurable: true})
}

shimScroll2()
</file>

<file path="src/mainLoader.ts">
function main() {
    const s = document.createElement("script")
    s.type = "text/javascript"
    s.async = true 
    s.src = chrome.runtime.getURL('main.js')
    document.documentElement.appendChild(s) 
}

main()
</file>

<file path="src/options/App.tsx">
import { Toggle } from "./comps/Toggle"
import { LOCALE_MAP } from "../utils/gsm"
import { Option } from "./comps/Option"
import { NumericInput } from "./comps/NumericInput"
import { ColorWell } from "./comps/ColorWell"
import { useKnownKeys } from "../hooks/useStorage"
import { openLink } from "../utils/browser"
import { CONFIG_KEYS } from "../types"
import "./styles.css"
import { isFirefox } from "./utils"

export function App() {
    const [items, setItems] = useKnownKeys(["g:lang", "g:autoClear", "g:context", "g:highlightColorDark", "g:highlightColorLight", "g:highlightBold", "o:lastTheme", "g:sessionOnly", "g:showImage", "g:orderByDate", "g:strictSearch", "g:scrollTop"])
    if (!items) return 

    return <div className="Options">
        <div className="section">
            <div className="title">{gvar.gsm.general}</div>
            {<Option label={gvar.gsm.language}>
                <select value={items["g:lang"]} onChange={async e => {
                    await chrome.storage.local.set({'g:lang': e.target.value})
                    location.reload()
                }}>
                    {Object.entries(LOCALE_MAP).map(([k, v]) => (
                        <option key={k} value={k} title={v.title}>{v.display}</option>
                    ))}
                </select>
            </Option>}
            {<Option label={gvar.gsm.autoClear} tooltip={gvar.gsm.autoClearTooltip}>
                <Toggle value={items["g:autoClear"]} onChange={v => {
                    setItems({"g:autoClear": v})
                }}/>
            </Option>}
            {items['o:lastTheme'] === 'dark' && (
                <Option label={gvar.gsm.highlightColor} tooltip={gvar.gsm.highlightColorTooltip}>
                    <ColorWell isActive={!!items["g:highlightColorDark"]} color={items["g:highlightColorDark"] || "#6dffd8"} onChange={v => {
                        setItems({'g:highlightColorDark': v})
                    }} onReset={() => {
                        chrome.storage.local.remove('g:highlightColorDark')
                    }}/>
                </Option>
            )}
            {items['o:lastTheme'] === 'light' && (
                <Option label={gvar.gsm.highlightColor} tooltip={gvar.gsm.highlightColorTooltip}>
                    <ColorWell isActive={!!items["g:highlightColorLight"]} color={items["g:highlightColorLight"] || "#008de5"} onChange={v => {
                        setItems({'g:highlightColorLight': v})
                    }} onReset={() => {
                        chrome.storage.local.remove('g:highlightColorLight')
                    }}/>
                </Option>
            )}
            {<Option label={gvar.gsm.highlightBold} tooltip={gvar.gsm.highlightBoldTooltip}>
                <Toggle value={items["g:highlightBold"]} onChange={v => {
                    setItems({"g:highlightBold": v})
                }}/>
            </Option>}
            <Option label={gvar.gsm.showImage} tooltip={gvar.gsm.showImageTooltip}>
                <Toggle value={items["g:showImage"]} onChange={v => {
                    setItems({"g:showImage": v})
                }}/>
            </Option>
            <Option label={gvar.gsm.scrollTop} tooltip={gvar.gsm.scrollTopTooltip}>
                <Toggle value={items["g:scrollTop"]} onChange={v => {
                    setItems({"g:scrollTop": v})
                }}/>
            </Option>
        </div>
        <div className="section">
            <div className="title">{gvar.gsm.search}</div>
            <Option label={gvar.gsm.context} tooltip={gvar.gsm.contextTooltip}>
                <NumericInput rounding={0} min={2} max={5} value={items["g:context"] ?? 2} onChange={v => {
                    setItems({'g:context': v})
                }}/>
            </Option>
            {gvar.gsm._morpho || <Option label={gvar.gsm.strictSearch} tooltip={gvar.gsm.strictSearchTooltip}>
                <Toggle value={items["g:strictSearch"]} onChange={v => {
                    setItems({"g:strictSearch": v})
                }}/>
            </Option>}
            <Option label={gvar.gsm.sortDate} tooltip={gvar.gsm.sortDateTooltip}>
                <Toggle value={items["g:orderByDate"]} onChange={v => {
                    setItems({"g:orderByDate": v})
                }}/>
            </Option>
        </div>
        <div className="section">
            <div className="title">{gvar.gsm.data}</div>
            {<Option label={gvar.gsm.sessionOnly} tooltip={gvar.gsm.sessionOnlyTooltip}>
                <Toggle value={items["g:sessionOnly"]} onChange={v => {
                    if (v) {
                        if (confirm(gvar.gsm.areYouSure)) {
                            removeCachedChats()
                            setItems({"g:sessionOnly": v})
                        }
                    } else {
                        setItems({"g:sessionOnly": v})
                    }
                }}/>
            </Option>}
        </div>
        <div className="section">
            <div className="title">{gvar.gsm.otherProjects.header}</div>
            <div className="promo">
                <a href="https://youtu.be/-tlSHOrf4ro">Ask Screenshot for ChatGPT: </a>
                <span dangerouslySetInnerHTML={{__html: gvar.gsm.otherProjects.askScreenshot}}/>
            </div>
            <div className="promo">
                <a href="https://github.com/polywock/globalSpeed">Global Speed: </a>
                <span>{gvar.gsm.otherProjects.globalSpeed}</span>
            </div>
        </div>
        <div className="section hasTitle help">
            <div className="title">{gvar.gsm.help}</div>
            <br />
            <div className="card">{`${gvar.gsm.issuePrompt} `}<a href="https://github.com/polywock/gpt-search">{gvar.gsm.issueDirective}</a></div>
            <br />
            <div className="buttons">
                <button className="RedButton" onClick={async e => {
                    if (confirm(gvar.gsm.areYouSure)) {
                        chrome.runtime.sendMessage({type: 'RESET'}, () => {
                            location.reload()
                        })
                    }
                }}>{gvar.gsm.reset}</button>
                <button className="RedButton" onClick={async e => {
                    if (confirm(gvar.gsm.areYouSure)) {
                        removeAllCache().then(() => {
                            location.reload()
                        })
                    }
                }}>{gvar.gsm.clearCache}</button>
                <div></div>
                <button onClick={async e => {
                    openLink(`https://github.com/polywock/gptSearch/blob/main/advancedSearch.md`)
                }}>{gvar.gsm.advancedSearch}</button>
            </div>
        </div>
    </div>
}


async function removeCachedChats() {
    const items = await chrome.storage.local.get()
    let keysToDelete: string[] = ['o:auth']
    for (let key in items) {
        if (key.startsWith("o:c:")) keysToDelete.push(key)
    }
    chrome.storage.local.remove(keysToDelete)
}

async function removeAllCache() {
    const items = await chrome.storage.local.get(CONFIG_KEYS)
    await Promise.all([
        chrome.storage.local.clear(),
        chrome.storage.session.clear()
    ])
    await chrome.storage.local.set(items)
}
</file>

<file path="src/options/comps/ColorWell.css">
.ColorWell {

    input[type="color"] {
        width: 32px;
        margin-left: 10px;
    }
    input[type="color"]::-webkit-color-swatch-wrapper {
        padding: 0;
        border: none;
    }
    input[type="color"]::-webkit-color-swatch {
        border: none;
    }
}
</file>

<file path="src/options/comps/ColorWell.tsx">
import { Reset } from "./Reset"
import "./ColorWell.css"

type ColorWellProps = {
    color: string,
    onChange: (newColor: string) => void,
    isActive?: boolean,
    onReset?: () => void 
}

export function ColorWell(props: ColorWellProps) {
    return <div className="ColorWell">
        <Reset active={props.isActive} onClick={() => {
            props.onReset?.()
        }}/>
        <input value={props.color} type="color" onChange={e => {
            props.onChange(e.target.value)
        }}/>
    </div>
}
</file>

<file path="src/options/comps/FloatTooltip.css">
.FloatTooltip {
    position: absolute;
    left: -80px; 
    right: -80px; 
    bottom: 2.85rem;
    font-size: 0.9em;
    display: grid;
    justify-content: center;
    
    & > div {
      display: inline-block; 
      padding: 0.357rem;
      background-color: red;
      color: white;
    }
  }
</file>

<file path="src/options/comps/FloatTooltip.tsx">
import "./FloatTooltip.css"


type FloatTooltipProps = {
  value: string
}

export const FloatTooltip = (props: FloatTooltipProps) => {
  return <div className="FloatTooltip">
    <div>
      {props.value}
    </div>
  </div>
}
</file>

<file path="src/options/comps/NumericInput.css">
.NumericInput input {
    width: 60px;
}
</file>

<file path="src/options/comps/NumericInput.tsx">
import { useState, useEffect, ChangeEvent } from "react"
import { FloatTooltip } from "./FloatTooltip"
import { round } from "../../helper"
import "./NumericInput.css"


const NUMERIC_REGEX = /^-?(?=[\d\.])\d*(\.\d+)?$/

type NumericInputProps = {
  value: number,
  onChange: (newValue: number) => any,
  min?: number,
  max?: number,
  rounding?: number,
  disabled?: boolean,
  className?: string
}


export const NumericInput = (props: NumericInputProps) => {
  const [ghostValue, setGhostValue] = useState("")
  const [problem, setProblem] = useState(null as string)

  useEffect(() => {
    setProblem(null)
    if (props.value == null) {
      ghostValue !== "" && setGhostValue("")
    } else {
      let parsedGhostValue = parseFloat(ghostValue)
      if (parsedGhostValue !== props.value) {
        setGhostValue(`${round(props.value, props.rounding ?? 4)}`)
      }
    }
  }, [props.value])
  
  
  const handleOnChange = (e: ChangeEvent<HTMLInputElement>) => {
    setGhostValue(e.target.value)
    const value = e.target.value.trim()
    
    const parsed = round(parseFloat(value), props.rounding ?? 4)

    if (!isNaN(parsed) && NUMERIC_REGEX.test(value)) {
      let min = props.min
      let max = props.max 

      if (min != null && parsed < min) {
        setProblem(`>= ${min}`)
        return 
      }
      if (max != null && parsed > max) {
        setProblem(`<= ${max}`)
        return
      }

      if (parsed !== round(props.value, props.rounding ?? 4)) {
        props.onChange(parsed)
      }
      setProblem(null)
    } else {
      setProblem(`NaN`) 
    }

  }

  return (
    <div className={`NumericInput ${props.className || ""}`} style={{position: "relative"}}>
      <input 
        disabled={props.disabled ?? false}
        onBlur={e => {
          setProblem(null)
          setGhostValue(props.value == null ? "" : `${round(props.value, props.rounding ?? 4)}`)
        }}
        className={problem ? "error" : ""}
        type="text" 
        onChange={handleOnChange} value={ghostValue}
      />
      {problem && (
        <FloatTooltip value={problem}/>
      )}
    </div>
  )
}
</file>

<file path="src/options/comps/Option.css">
.Option {
    display: grid;
    grid-template-columns: 1fr max-content;
    align-items: center;
    margin-bottom: 20px;
    column-gap: 40px;

    &:last-child {
        margin-bottom: 0px;
    }

    & > .display {
        & > .context {
            font-size: 0.9em;
            opacity: 0.8;
        }
    }
}
</file>

<file path="src/options/comps/Option.tsx">
import { ReactElement } from 'react'
import './Option.css'

type OptionProps = {
    label: string,
    tooltip?: string,
    children?: ReactElement
}

export function Option(props: OptionProps) {
    return <div className="Option">
        <div className="display">
            <div className="label">{props.label}</div>
            {props.tooltip && (
                <div className="context">{props.tooltip}</div>
            )} 
        </div>
        {props.children}
    </div>
}
</file>

<file path="src/options/comps/Reset.css">
.Reset {
    color: var(--text-color);
    border: 1px solid var(--text-color);
    padding: 2px;
    box-sizing: content-box;
    opacity: 0.5;
    user-select: none;
    border-radius: 5px;
  
    &.active {
      opacity: 1;
    }
}
</file>

<file path="src/options/comps/Reset.tsx">
import { GiAnticlockwiseRotation } from "react-icons/gi";
import "./Reset.css"

type ResetProps = {
  onClick?: () => void,
  active?: boolean
}

export function Reset(props: ResetProps) {
  return <GiAnticlockwiseRotation size={"1.07rem"} className={`Reset ${props.active ? "active" : ""}`} onClick={props.onClick}/>
}
</file>

<file path="src/options/comps/Toggle.css">
.Toggle {
    display: inline-block;
    width: 45px; 
    /* border-radius: 50px; */
    border-radius: 5px;
    background-color: #a4a4a4;
    border: 2px solid  #a4a4a4;
    line-height: 0;
    cursor: pointer;
  
    &::after {
      pointer-events: none;
      content: "";
      display: inline-block;
      background-color: white;
      border: 1px solid #aaa;
      border-radius: inherit;
      box-sizing: border-box;
      width: 18px;
      height: 18px;
      transition: transform 0.05s linear;
      transform: translateX(0px);
    }
  
    &.active {
      background-color: var(--accent);
      border-color: var(--accent);
  
      &::after {
        transform: translateX(26px);
      }
    }
  }
</file>

<file path="src/options/comps/Toggle.tsx">
import './Toggle.css'

type ToggleProps = {
    value: boolean,
    onChange: (newValue: boolean) => void 
  }
  
export function Toggle(props: ToggleProps) {
    return <div tabIndex={0} onKeyDown={e => {
        if (e.key === "Enter") {
        props.onChange(!props.value)
        }
    }} onClick={e => {
        props.onChange(!props.value)
    }} className={`Toggle ${props.value ? "active" : ""}`}/>
}
</file>

<file path="src/options/index.tsx">
import { createRoot } from "react-dom/client";
import { loadGsm } from "../utils/gsm";
import { App } from "./App";

loadGsm().then(gsm => {
    gvar.gsm = gsm 
    if (gvar.gsm) main()
})

async function main() {
    const root = createRoot(document.querySelector('#root'))
    root.render(<App/>)
}
</file>

<file path="src/options/styles.css">
:root {
    --accent: #4e00fd;
    --accent-text-color: white;


    --text-color: white;
    --text-color-alt: #eaeaea;

    --true-bg: black;
    --bg: #181818;
    --bg-secondary: #222;
    --link-color: #c7a4f9;
    
    --border-1: #ffffff44;
    --border: #ffffff66;
    --border1: #ffffff88;

    --menu-color: #422d73;
    --menu-color-secondary: #594094;
    --section-width: 700px;
}

body {
    margin: 0;
    background-color: var(--true-bg);
    font-family: "Segoe UI", "Avenir", Courier, monospace;
}

.Options {
    color: var(--text-color-alt);
    font-size: 16px;
    display: grid;
    justify-content: center;
    margin-top: 80px;

    & > .section {
        background-color: var(--bg);
        width: var(--section-width);
        padding: 15px;
        padding-bottom: 40px;
        margin-bottom: 50px;
        font-size: 0.95rem;

        & > .title {
            background-color: var(--accent);
            color: var(--accent-text-color);
            padding: 3px 12px;
            display: inline-block;
            margin-bottom: 20px;
            font-size: 1.1rem;
        }

        .raw {
            white-space: word-break;
        }

        .colorWithReset {
            display: grid;
            align-items: center;
            grid-auto-flow: column;
            column-gap: 5px;
        }

        & > .promo {
            margin-bottom: 10px;
            line-height: 1.5;
        }
    }
}

select, input[type="text"], button {
    font-family: inherit;
    font-size: inherit;
    background-color: inherit;
    color: inherit;
    padding: 8px;
    
    &:disabled {
        cursor: not-allowed;
    }
}

/* Need explicit for Edge. */
select {
    background-color: var(--bg);
    color: var(--text-color);
}

button {
    border: 1px solid var(--border-1);
}


select, input[type="text"] {
    --border: var(--border-1);
    border: none;
    border-radius: 0px;
    border-bottom: 2px solid var(--border);

    &:focus {
        border: none;
        outline: none;
        border-bottom: 2px solid var(--border1);
    }
}



select, button {
    cursor: pointer;
    &:hover {
        opacity: 0.8;
    }
}



button {
    font-size: 1.1em;
}

select, input[type="text"] {
    text-align: center;
}

.card {
    padding: 10px;
    border: 1px solid var(--border-1);
    display: inline-block;
    font-size: 1.1em;
    white-space: pre;
}

a:any-link {
    color: var(--link-color);
    text-decoration: none;
}

a:hover {
    text-decoration: underline;
}

.section.help {
    .buttons {
        display: grid;
        margin-top: 40px;
        font-size: 1rem;
        grid-template-columns: max-content max-content 1fr max-content;
        column-gap: 10px;
        align-items: center;
    }
    
    .RedButton {
        background-color: #9b3232;
        color: white;
        border: none;
    }
}
</file>

<file path="src/options/utils.ts">
import debounce from "lodash.debounce"

export function _requestTemplateSync() {
    chrome.runtime.sendMessage({type: 'SYNC_TEMPLATES'})
}

export const requestTemplateSync = debounce(_requestTemplateSync, 1000, {leading: true})

let isFirefoxResult: boolean
export function isFirefox() {
  isFirefoxResult = isFirefoxResult ?? navigator.userAgent.includes("Firefox/")
  return isFirefoxResult
}
</file>

<file path="src/preamble/index.ts">
import rawStyle from '../common/searchStyle.css?raw'
import { createElement, timeout } from '../helper'
import preStyle from './style.css?raw'

declare global {
    interface GlobalVar {
        askedForRaccoon: boolean,
        preambleHost: HTMLDivElement,
        preambleProxy: HTMLElement,
        preambleStub: HTMLInputElement,
        auth?: string,
        prep?: {
            ph?: string
        },
        mo?: MutationObserver
        lastNav: HTMLElement
    }
    
    var gvar: GlobalVar
}


function loadScaffold() {
    const proxy = document.createElement('div')
    const host = createElement(`<div style='z-index: 999'></div>`) as HTMLDivElement
    proxy.appendChild(host)
    const shadow = host.attachShadow({mode: 'closed'})

    const style = document.createElement('style')
    style.textContent = `${preStyle}\n${rawStyle}`

    shadow.appendChild(style)

    const searchWrapper = createElement(`<div class='wrapper'></div>`)
    if (gvar.lastNav.className.includes("_sidebar")) {
        searchWrapper.classList.add('new')
    }

    const search = createElement(`<div class='search'></div>`) as HTMLDivElement
    const searchIcon = createElement(`<svg class='searchIcon' stroke="currentColor" fill="none" stroke-width="2" viewBox="0 0 24 24" stroke-linecap="round" stroke-linejoin="round" class="pointer-none absolute left-3 top-0 mr-2 h-full text-token-text-tertiary left-6" height="16" width="16" xmlns="http://www.w3.org/2000/svg"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line></svg>`) as SVGElement
    const searchInput = createElement(`<input></input>`) as HTMLInputElement
    searchInput.placeholder = gvar.prep.ph
    search.appendChild(searchIcon)
    search.appendChild(searchInput)
    searchWrapper.appendChild(search)
    shadow.appendChild(searchWrapper)
    
    searchInput.addEventListener('pointerdown', handleStubClick, {capture: true, once: true})
    
    gvar.preambleProxy = proxy
    gvar.preambleStub = searchInput
    gvar.preambleHost = host
}


function insertPageStyle() {
    const s = document.createElement("style")
    s.textContent = `div[role="presentation"]:focus {
        outline: none; 
    }`
    document.documentElement.appendChild(s) 
}

function handleStubClick(e: KeyboardEvent) {
    loadRaccoon()
}
  
function loadRaccoon() {
    if (gvar.askedForRaccoon) return 
    chrome.runtime.sendMessage({type: 'LOAD_RACCOON'})
    gvar.askedForRaccoon = true 
}

function handleMut(muts: MutationRecord[]) {
    if (gvar.lastNav?.isConnected) return 
    for (let mut of muts) {
        for (let added of mut.addedNodes) {
            if (added.nodeType !== Node.ELEMENT_NODE) continue 

            if (checkIfNav(added as Element)) {
                onNewNav(added as HTMLElement)
                return 
            }

            let nav = getNav(added as Element)
            nav && onNewNav(nav)

        }
    }
}

function checkIfNav(elem: Element) {
    if (elem?.tagName === "NAV" && elem.ariaLabel && elem.classList.contains("h-full")) return true 
    if (elem?.tagName === "NAV" && elem.className.includes("_sidebar")) return true 
}

function onNewNav(nav: HTMLElement) {
    gvar.lastNav = nav 
    gvar.preambleProxy ?? loadScaffold()
    nav.insertAdjacentElement('afterbegin', gvar.preambleProxy)
}

function getNav(root?: Element) {
    for (let nav of (root ?? document.body).getElementsByTagName("nav")) {
        if (nav.ariaLabel && nav.classList.contains("h-full")) {
            return nav 
        }
        if (nav.className.includes("_sidebar")) {
            return nav 
        }
    }
}


async function onLoaded() {
    await timeout(500)
    const nav = getNav()
    nav && onNewNav(nav)

    gvar.mo = new MutationObserver(handleMut)
    gvar.mo.observe(document, {subtree: true, childList: true})
}

function main() {
    gvar.prep = {}
    chrome.runtime.sendMessage({type: 'PREP'}).then(v => {
        gvar.prep = v ?? gvar.prep
    }).finally(() => {
        if (document.readyState === "loading")  {
            document.addEventListener("DOMContentLoaded", onLoaded, {capture: true, once: true})
        } else {
            onLoaded() 
        }
    })
    
    insertPageStyle()
}


main()
</file>

<file path="src/preamble/style.css">
.search {
    
}
</file>

<file path="src/raccoon/App.tsx">
import { useEffect, useRef, useState } from "react"
import { SearchChats } from "./searchChats"
import { ResultItem } from "./comps/ResultItem"
import { Close, Gear, Github, Heart, Pin, SearchSvg, Star } from "../comps/svgs"
import { Status } from "./types"
import { LoadMore } from "./comps/LoadMore"
import { useAutoBar } from "./hooks/useAutoBar"
import { useClickBlur } from "./hooks/useClickBlur"
import clsx from "clsx"
import { Config } from "../types"
import { useResize } from "./hooks/useResize"

export const App = (props: {dark: boolean, top: number, left: number, config: Config, isNewDesign: boolean}) => {
    const { config } = props 


    const [blur, setBlur] = useState(false)
    const [pinned, setPinned] = useState(false)
    const mainRef = useRef<HTMLDivElement>(null)
    const searchRef = useRef<HTMLInputElement>(null)
    const [query, setQuery] = useState("")
    const [status, setStatus] = useState<Status>(null)
    const smartBlur = useRef((value: boolean) => {
        if (value && config["g:autoClear"]) {
            setQuery("")
        }
        setBlur(value)
    })
    gvar.raccoonSearch = searchRef.current
    
    useAutoBar(blur, searchRef)
    useClickBlur(blur, !!query && pinned, smartBlur.current)
    const [scale, windowSize] = useResize(mainRef)

    useEffect(() => {
        if (!query) {
            setStatus(null)
            return 
        } 
        setStatus({results: []})
        mainRef.current.scrollTop = 0
        const searchChats = new SearchChats(query, setStatus, mainRef)
        return () => {
            searchChats?.release()
        }
    }, [query])

    const colorOverride = config[props.dark ? "g:highlightColorDark" : "g:highlightColorLight"]

    return <div id="App" className={clsx({
        'peacock': !blur,
        dark: props.dark,
        bold: config["g:highlightBold"],
        wrapper: blur,
        new: props.isNewDesign
    })} style={{
        top: props.top,
        left: props.left,
        width: blur ? undefined : `${windowSize.w}px`
    }}>
        <div className="search">
            <SearchSvg className="searchIcon"/>
            <input onKeyDown={e => {
                if (e.key === "Escape") {
                    setQuery('')
                    setBlur(true)
                }
            }} ref={searchRef} onFocusCapture={e => setBlur(false)} autoFocus={true} type="text" placeholder={gvar.gsm.search} value={(blur && query) ? "..." : query} onChange={e => setQuery(e.target.value)}/>
            {query && <Close className="closeIcon" onClick={e => {
                setQuery('')
                setBlur(true)
            }}/>}
        </div>
        {!blur && (
            <>
                <div className="main" ref={mainRef} style={{
                    maxHeight: `${windowSize.h}vh`,
                    ...(colorOverride ? ({'--needle-color': colorOverride} as any) : {})
                }}>
                    {!!(status?.results.length) && status.results.map(v => (
                        <ResultItem scrollTop={config["g:scrollTop"]} result={v} key={v.id}></ResultItem>
                    ))} 
                    {status && !status.finished && <LoadMore/>}
                    {status && status.finished && !status.results?.length && <div className="NonFound">{gvar.gsm.notFound}</div>}
                    
                </div>
                <div className="footer">
                    <button className="svgButton" onClick={e => {
                        chrome.runtime.sendMessage({type: 'OPEN_LINK', url: chrome.runtime.getURL(`options.html`), active: true})
                    }}><Gear style={{transform: 'scale(1.2)'}}/></button>
                    <button className="svgButton" onClick={e => {
                        chrome.runtime.sendMessage({type: 'OPEN_LINK', url: `https://github.com/polywock/gpt-search`, active: true})
                    }}>
                        <Github/>
                    </button>
                    <button className={`svgButton toggable pin ${pinned ? 'active' : ''}`} onClick={e => {
                        setPinned(!pinned)
                    }}><Pin/></button>
                </div>
                {scale}
            </>
        )}
    </div>
}

window.addEventListener("keypress", e => {
    const shadow = (e.target as any)?.shadowRoot as ShadowRoot
    if (shadow && shadow.activeElement?.tagName === "INPUT") e.stopImmediatePropagation()
}, true)
</file>

<file path="src/raccoon/comps/CleverDiv.tsx">
type CleverDivProps = React.HTMLProps<HTMLDivElement> & {
    onCleverClick?: (e: React.KeyboardEvent | React.PointerEvent) => void 
}

export function CleverDiv(props: CleverDivProps) {
    let p = {...props}
    const cleverClick = p.onCleverClick
    const ogPointerUp = p.onPointerUp 
    const ogKeyUp = p.onKeyUp 
    
    delete p.onCleverClick

    if (cleverClick) {
        delete p.onPointerUp 
        delete p.onKeyUp
    }

    return <div {...p} {...(cleverClick ? {
        onPointerUp: e => {
            if (e.button === 0) cleverClick(e)
            ogPointerUp?.(e)
        },
        onKeyDown: e => {
            if (e.key === "Enter") {
                cleverClick(e)
            }
            ogKeyUp?.(e)
        }
    } : {})}/>
}
</file>

<file path="src/raccoon/comps/LoadMore.tsx">
import { useEffect, useRef } from "react"
import { SearchChats } from "../searchChats"


let latestWasIntersected = false 

export function LoadMore(props: {}) {
    const ref = useRef(null as HTMLDivElement)
    useEffect(() => {
        const obs = new IntersectionObserver(entries => {
            latestWasIntersected = entries.at(-1).isIntersecting
            if (latestWasIntersected) {
                SearchChats.ref?.loadSafely()
            }
        }, {
            threshold: 0.75,
            root: ref.current.parentElement
        })
        obs.observe(ref.current)

        const intervalId = setInterval(() => {
            if (latestWasIntersected) {
                SearchChats.ref?.loadSafely()
            }
        }, 100)

        return () => {
            obs.disconnect()
            clearInterval(intervalId)
        }
    }, [ref.current, ref.current?.parentElement])

    return  <div ref={ref} className="LoadMore"></div> 
}
</file>

<file path="src/raccoon/comps/ResultItem.tsx">
import { memo, useRef, useState } from "react"
import { Result } from "../types"
import {  softLink } from "../utils/misc"
import { timeout } from "../../helper"
import { CleverDiv } from "./CleverDiv"
import { openLink } from "../../utils/browser"
import { isFirefox } from "../../options/utils"

const _ResultItem = (props: { result: Result, scrollTop: boolean}) => {
    const [max, setMax] = useState(4)
    const env = useRef({lastWasChild: false}).current
    const r = props.result
    const titleResult = r.headerResult

    const open = (messageId?: string, mode?: "fg" | "bg") => {
        let url = `/c/${r.id}`
        if (r.isGpt) {
            url = `/g/g-${r.id}`
        }
        if (mode) {
            openLink(`https://chatgpt.com${url}`, mode === "fg")
        } else {
            if (location.pathname !== url) softLink(url)
            clearTimeoutInfo()

            let scrollTo: {id: String, isCurrent?: boolean}

            if (messageId) {
                scrollTo = {id: messageId}
            } else if (r.currentNodeId) {
                scrollTo = {id: r.currentNodeId, isCurrent: true}
            }
            env.lastWasChild = !!messageId

            if (scrollTo) tryScrollIntoView(`div[data-message-id="${scrollTo.id}"]`, scrollTo.isCurrent, props.scrollTop) 
        }

    }

    const handleTitleClick = (e: React.PointerEvent | React.KeyboardEvent) => {
        open(null)
    }

    const handleMessageClick = (messageId: string, e: React.PointerEvent | React.KeyboardEvent) => {
        open(messageId)
    }

    const handleItemPointerUp= (e: React.PointerEvent) => {
        if (e.button !== 1) return 
        open(null, e.shiftKey ? "fg" : "bg")
    }

    let metaChunks: JSX.Element[] = []
    if (!r.isGpt && r.gizmoName) {
        r.gizmoImg && metaChunks.push(<img key="gizmoImg" src={r.gizmoImg} width="15px" height="15px"/>)
        metaChunks.push(<span key="gizmoName">{r.gizmoName}</span>)
    }
    if (r.elapsed) metaChunks.unshift(<span key="elapsed">{`${r.elapsed}${metaChunks.length === 1 ? ' · ' : ''}`}</span>)

    return <div className="ResultItem">
        <CleverDiv className="header" onCleverClick={handleTitleClick} onPointerUp={handleItemPointerUp} tabIndex={0}>
            <div className="title">
                {r.isGpt && r.gizmoImg && (
                    <img src={r.gizmoImg} width="20px" height="20px"/>
                )}
                <span className="normal">{titleResult ? titleResult.prefix : r.title}</span>
                {titleResult && <>
                    <span className="needle">{titleResult.needle}</span>
                    <span className="normal">{titleResult.suffix}</span>
                </>}
            </div>
            {!!metaChunks.length && (
                <div className="meta">{metaChunks.map(c => c)}</div>
            )}
        </CleverDiv>
        {!!r.parts.length && (
            <div className="conti">
                <>
                    {r.parts.slice(0, max).map(part => (
                        <CleverDiv tabIndex={0} onPointerUp={handleItemPointerUp} onCleverClick={handleMessageClick.bind(this, part.messageId)} key={part.messageId} className="context">
                            <span className="normal">{part.prefix}</span>
                            <span className="needle">{part.needle}</span>
                            <span className="normal">{part.suffix}</span>
                        </CleverDiv>
                    ))}
                </>
                {r.parts.length > max && (
                    <CleverDiv tabIndex={0} className="context" onCleverClick={() => {
                        setMax(max + 10)
                    }}>{`${gvar.gsm.showMore} (${Math.min(r.parts.length - max, 10)})`}</CleverDiv>
                )} 
            </div>
        )}
    </div>
}

export const ResultItem = memo(_ResultItem)

let latestSymbol: Symbol 
let lastScrollPath: string 
let SUPPORTS_SCROLL_INTO_VIEW = "scrollIntoViewIfNeeded" in Element.prototype
let recentScroll: {
    target: string,
    subtle?: boolean,
    scrollTop?: boolean,
    time: number 
}



async function tryScrollIntoView(target: string, subtle?: boolean, scrollTop?: boolean, isFake?: boolean, n = 60, delay = 100) {
    gvar.scrollSetCbs.add(handleScrollSet)
    recentScroll = null 
    const mySymbol = Symbol()
    latestSymbol = mySymbol 
    for (let i = 0; i < n; i++) {
        i > 0 && await timeout(delay)
        if (latestSymbol !== mySymbol) return 
        const elem = document.querySelector(target)
        if (elem) {
            const isFirst = lastScrollPath !== location.pathname
            lastScrollPath = location.pathname
            // if (isFirst) await timeout(isFirefox() ? 500 : 500)
            if (isFirst && !isFake) recentScroll = {target, subtle, scrollTop, time: Date.now()}
            
            if (subtle) {
                const parent = getScrollableParent(elem)
                if (parent) {
                    parent.scrollTo({top: scrollTop ? 0 : 999999999, behavior: "instant"})
                    return 
                }
            } 
            !subtle && activateFor(elem)
            SUPPORTS_SCROLL_INTO_VIEW ? (elem as any).scrollIntoViewIfNeeded() : (elem as any).scrollIntoView()
            return 
        }
    }
}


let timeoutInfo: {
    target: Element,
    timeoutId: number 
}

function clearTimeoutInfo() {
    if (timeoutInfo) {
        clearTimeout(timeoutInfo.timeoutId)
        timeoutInfo.target.removeAttribute("bornagain")
        timeoutInfo = null 
    }
}

function activateFor(target: Element) {
    ensureStyleElement()
    clearTimeoutInfo()
    
    target.setAttribute("bornagain", "")
    timeoutInfo = {
        target,
        timeoutId: window.setTimeout(clearTimeoutInfo, 2500)
    }
}


let focusStyle: HTMLStyleElement
function ensureStyleElement() {
    if (!focusStyle) {
        focusStyle = document.createElement('style')
        focusStyle.textContent = ":is([bornagain], #bornagainnnn > #woo > #barked) { outline: 1px solid red !important; }"
    } 
    if (!focusStyle.isConnected) document.documentElement.appendChild(focusStyle)
}



function getScrollableParent(element: Element) {
    let parent = element.parentNode as Element
  
    while (parent) {
      const overflowY = window.getComputedStyle(parent).overflowY
      if ((overflowY === 'auto' || overflowY === 'scroll') && parent.scrollHeight > parent.clientHeight) {
        return parent
      }
      parent = parent.parentNode as Element
    }
}



async function handleScrollSet() {
    let originalRecentScroll = recentScroll
    if (!(recentScroll && Date.now() - recentScroll.time < 10_000)) return 
    await timeout(1000)
    if (originalRecentScroll !== recentScroll) return 
    tryScrollIntoView(recentScroll.target, recentScroll.subtle, recentScroll.scrollTop, true)
}
</file>

<file path="src/raccoon/hooks/useAutoBar.tsx">
import { useLayoutEffect } from "react"
import { isFirefox } from "../../options/utils"

export function useAutoBar(blur: boolean, searchRef: React.MutableRefObject<HTMLInputElement>) {
    useLayoutEffect(() => {
        if (blur) {
            enterSidebar()
        } else {
            exitSidebar()
            searchRef.current.focus()
            isFirefox() && requestAnimationFrame(() => {
                searchRef.current.focus()
            })
        }
    }, [blur, searchRef])
}

function enterSidebar() {
    gvar.preambleProxy.insertAdjacentElement('afterbegin', gvar.raccoonHost)
    gvar.preambleHost.remove()
}

function exitSidebar() {
    gvar.preambleProxy.insertAdjacentElement('afterbegin', gvar.preambleHost)
    document.documentElement.appendChild(gvar.raccoonHost)
}
</file>

<file path="src/raccoon/hooks/useClickBlur.tsx">
import { useEffect } from "react"

export function useClickBlur(blur: boolean, pinned: boolean, setBlur: (newValue: boolean) => any) {
    useEffect(() => {
        if (blur || pinned) return 

        const handleClick = (e: PointerEvent) => {
            if (e.target !== gvar.preambleHost && e.target !== gvar.raccoonHost) {
                setBlur(true)
            }
        }
        
        window.addEventListener('pointerdown', handleClick, true)

        return () => {
            window.removeEventListener('pointerdown', handleClick, true)
        }
    }, [blur, pinned])
}
</file>

<file path="src/raccoon/hooks/useResize.tsx">
import { useEffect, useState } from "react"
import { clamp } from "../../helper"

type DragRef = {
    x: number,
    y: number,
    w: number,
    h: number 
}

type Bounds = {
    w: number,
    h: number 
}


export function useResize(mainRef: React.MutableRefObject<React.ElementRef<'div'>>) {
    const [dragRef, setDragRef] = useState(null as DragRef)
    const [windowSize, setWindowSize] = useState({w: 350, h: 60} as Bounds)

    useEffect(() => {
        if (!dragRef) return 
        const handleMove = (e: PointerEvent) => {
            const deltaX = e.clientX - dragRef.x
            const x = clamp(300, 600, dragRef.w + deltaX)

            const deltaY = (e.clientY - dragRef.y) / window.innerHeight * 100
            let y = clamp(50, 80, Math.min(dragRef.h + deltaY, mainRef.current.scrollHeight / window.innerHeight * 100))

            setWindowSize({w: x, h: y})
        }
        const handleUp = (e: PointerEvent) => {
            handleMove(e)
            setDragRef(null)
        }
        window.addEventListener('pointermove', handleMove, {capture: true})
        window.addEventListener('pointerup', handleUp, {capture: true})

        return () => {
            window.removeEventListener('pointermove', handleMove, {capture: true})
            window.removeEventListener('pointerup', handleUp, {capture: true})
        }
    }, [dragRef])

    return [
        <div className="scale" onPointerDownCapture={e => {
            setDragRef({x: e.clientX, y: e.clientY, w: windowSize.w, h: windowSize.h})
        }}/>,
        windowSize
    ] as [React.ReactElement<HTMLDivElement>,Bounds]
}
</file>

<file path="src/raccoon/index.tsx">
import { createRoot } from "react-dom/client"
import rawStyle from './style.css?raw'
import rawSearchStyle from '../common/searchStyle.css?raw'
import { App } from './App'
import { requestGsm } from "../utils/gsm"
import { CONFIG_KEYS, Config } from "../types"
import { timeout } from "../helper"
import { localGet, localSet, sessionGetFallback, sessionSetFallback } from "../utils/browser"

declare global {
    interface GlobalVar {
        raccoonHost: HTMLDivElement,
        raccoonSearch?: HTMLInputElement,
        config?: Config,
        getItems?: typeof localGet,
        setItems?: typeof localSet,
        scrollSetCbs?: Set<() => void>
    }
}
gvar.scrollSetCbs = gvar.scrollSetCbs || new Set()

async function main() {
    const b = gvar.preambleStub?.getBoundingClientRect()
    if (b.left == null || b.top == null) return 
    
    const [config, auth] = await Promise.all([
        chrome.storage.local.get(CONFIG_KEYS) as Promise<Config>,
        getAuth()
    ])
    if (auth) gvar.auth = auth
    if (!gvar.auth) {
        await timeout(1000)
        gvar.auth = await chrome.runtime.sendMessage({type: 'GET_AUTH'}) 
    }
    if (!gvar.auth) console.error('No auth found')

    if (config["g:sessionOnly"]) {
        gvar.getItems = sessionGetFallback
        gvar.setItems = sessionSetFallback
    } else {
        gvar.getItems = localGet
        gvar.setItems = localSet
    }

    gvar.config = config 

    gvar.raccoonHost = document.createElement('div')
    
    const shadow = gvar.raccoonHost.attachShadow({mode: 'open'})
    const rootBase = document.createElement('div')
    const style = document.createElement('style')
    style.textContent = rawSearchStyle.concat(rawStyle)

    shadow.appendChild(rootBase)
    shadow.appendChild(style)

    document.documentElement.appendChild(gvar.raccoonHost)

    const root = createRoot(rootBase)
    chrome.storage.local.set({'o:lastTheme': document.documentElement.classList.contains('dark') ? 'dark' : 'light'})

    const dark = document.documentElement.classList.contains('dark')

    root.render(<App config={config} dark={dark} top={b.top} left={b.left} isNewDesign={gvar.lastNav.className.includes("_sidebar")}/>)
}


requestGsm().then(gsm => {
    gvar.gsm = gsm 
    if (gvar.gsm) main()
})


window.addEventListener('rusrusar', (e: CustomEvent) => {
    const deets = JSON.parse(e.detail)
    if (deets.type === 'NO_PUSH') {
        chrome.runtime.sendMessage({type: "OPEN_LINK", url: `${location.origin}${deets.path}`, active: false})
    } else if (deets.type === "SCROLL_SET") {
        gvar.scrollSetCbs?.forEach(cb => cb())
    }
    e.stopImmediatePropagation()
}, {capture: true})

async function getAuth() {
    let auth = (await sessionGetFallback(['s:auth']))['s:auth']
    if (auth) return auth 
    auth = (await chrome.storage.local.get('o:auth'))['o:auth']
    if (auth) return auth 
}
</file>

<file path="src/raccoon/searchChats/core.ts">
import escapeStringRegexp from 'escape-string-regexp';

type Context = {
    opts: Options,
    exclude: string[],
    terms: Term[]
}


type Options = {
    ordered?: boolean,
    threshold?: number 
}

type Solt = {lb: number, rb: number, idx: number, score: number}

type Term = {
    needle: string,
    needleSize: number,
    groups: {match: string | RegExp, points: number}[][]
}

// First check if contains 
// Next check if contains strictly with delims. 

export function search(items: string[], query: string, opts: Options): Solt[] {
    const { needles, exclude } = getNeedlesAndExclude(query || "")

    if (needles.length === 0) {
        if (exclude.length) return getNonExcluded(items, exclude)
        return []
    }


    const env: Context = {
        exclude, 
        opts,
        terms: needles.map(needle => {
            let escaped = escapeStringRegexp(needle)?.trim()
            const term = {
                needle,
                needleSize: needle.length,
                groups: [
                    [{match: needle.toLocaleLowerCase(), points: 20}]
                ]
            } as Term

            try {
                term.groups.push(
                    [
                        {match: new RegExp(String.raw`${escaped}(?!\p{Letter})`, 'ui'), points: 20},
                        {match: new RegExp(String.raw`(?<!\p{Letter})${escaped}`, 'ui'), points: 30}
                    ],  
                    [{match: new RegExp(String.raw`(?<!\p{Letter})${escaped}(?!\p{Letter})`, 'ui'), points: 20}]  
                )
            } catch (err) { }

            return term 
        })
    }

    let newItems: Solt[] = []
    for (let i = 0; i < items.length; i++) {
        const item = items[i].toLocaleLowerCase()
        const solt = scoreItem(item, env)
        if (solt?.score >= (opts.threshold ?? 20)) newItems.push({...solt, idx: i} as Solt)
    }

    if (opts.ordered) {
        newItems.sort((a, b) => {
            return b.score - a.score 
        })
    }

    return newItems
}

function getNeedlesAndExclude(query: string) {
    let needles = []
    let current: string[] = []
    let exclude: string[] = []
    for (let token of query.split(' ')) {
        token = token.trim()
        const startsWithMinus = token.startsWith("-")
        if (startsWithMinus || token.startsWith("+")) {
            if (current.length) {
                needles.push(current.join(" ").trim())
                current = []
            }
            token.length > 1 && (startsWithMinus ? exclude : needles).push(token.slice(1))
            continue 
        } 
        token.length && current.push(token)
    }
    if (current.length) needles.push(current.join(" ").trim())
    
    return {needles, exclude}
}

function scoreWithTerm(v: string, term: Term, env: Context): Partial<Solt>  {
    let score = 0
    let lb: number 
    let rb: number 

    for (let group of term.groups) {
        let matched = false 
        for (let {match, points} of group) {
            let idx = typeof match === "object" ? v.search(match) : v.indexOf(match)
            if (idx < 0) continue 
            matched = true 
            score += points 
            lb = idx 
            rb = lb + term.needleSize
        }
        if (!matched) break 
    }

    return {score, lb, rb} 
}

function scoreItem(v: string, env: Context): Partial<Solt> {
    let highestRes: Partial<Solt> 

    for (let term of env.terms) {
        const res = scoreWithTerm(v, term, env)
        if (res && res.score > (highestRes?.score ?? -1)) {
            highestRes = res
        }
    }

    if (checkExcluded(v, env.exclude)) return 

    return highestRes
}


function checkExcluded(v: string, exclude: string[]) {
    for (let token of exclude) {
        if (v.includes(token)) return true 
    }
}

function getNonExcluded(items: string[], exclude: string[]) {
    let solts: Solt[] = []

    for (let i = 0; i < items.length; i++) {
        const item = items[i]
        if (checkExcluded(item, exclude)) continue 
        solts.push({idx: i, lb: 0, rb: 0, score: 20})
    }
    return solts 
}
</file>

<file path="src/raccoon/searchChats/extractContext.ts">
const FULL_STOP = /[?!\.。！？۔।॥\|¡¿\n\t]/

export type Context = {prefix: string, needle: string, suffix: string}

export function getContext(content: string, start: number, end: number, context = 2) {
    if (!end || start === end) return {prefix: content.slice(0, context * 200), suffix: '', needle: ''}

    content = content.trim() // replace(/\n+/, '\n')
    const preContext = Math.ceil(context / 2)
    const postContext = context - preContext

    const needle = content.slice(start, end)

    let prefix = ''
    if (start) {
        const tempPrefix = content.slice(0, start).split('').reverse().join('')
        prefix = collectSearch(tempPrefix, preContext, preContext * 100, true).split('').reverse().join('').trimStart()
    }

    const tempSuffix = content.slice(end)
    let suffix = collectSearch(tempSuffix, postContext, postContext * 100).trimEnd()
    return {prefix, needle, suffix} as Context
}


export function collectSearch(content: string, n: number, max: number, minus?: boolean) {
    let collected = 0 
    for (let i = 0; i < n; i++) {
        const idx = content.slice(collected).search(FULL_STOP)
        if (idx === -1) return content.slice(0, max)

        collected += idx + 1 
    }
    return content.slice(0, Math.min(minus ? collected - 1 : collected, max))
}
</file>

<file path="src/raccoon/searchChats/extractOpts.ts">
import { getGizmosSync } from "../utils/gizmo"
import { search } from "./core"

const optBaseFlags = ['dalle', 'browse', 'python', 'gizmo', 'gizmos', 'title', 'gpt', 'body', 'ast', 'user', 'gpt4', 'archived', 'archive', 'g']
const optSourceMapping = [
    ['+c', 'createdAfter', 'date'], ['-c', 'createdBefore', 'date'], 
    ['+u', 'updatedAfter', 'date'], ['-u', 'updatedBefore', 'date'],
    ['+turns', 'turnsPlus', 'int'], ['-turns', 'turnsMinus', 'int']
]

export function extractOpts(query: string): {query: string, opts: PreFilter} {
    let opts: PreFilter = {}
    let tokens = query.toLowerCase().split(' ')
    let tokensB: string[] = []

    tokens.forEach(token => {
        const suffix = token.slice(1)
        let matched = false 
        if (optBaseFlags.includes(suffix)) {
            if (token[0] === '+')  {
                (opts as any)[suffix] = true 
                matched = true
            } else if (token[0] === '-')  {
                (opts as any)[suffix] = false 
                matched = true
            } 
        }
        if (!matched) {
            tokensB.push(token.trim())
        }
    })

    let matched = false 
    tokens = []
    for (let [index, token] of tokensB.entries()) {
        if (matched) {
            matched = false 
            continue 
        }

        for (let [flag, propertyName, type] of optSourceMapping) {
            if (token === flag) {
                matched = true;
                const parsed = parseAs(tokensB[index + 1], type as any)
                if (parsed) (opts as any)[propertyName] = parsed
            }
        }
        if (!matched) {
            tokens.push(token.trim())
        }
    }
    // pre-processing 
    let exclusionMode = ![opts.ast, opts.body, opts.title, opts.user, opts.gpt].some(v => v)
    opts.title = opts.title || (exclusionMode && opts.title !== false)
    opts.gpt = opts.gpt || (exclusionMode && opts.gpt !== false)

    // False first to preserve specificity 
    opts.ast = opts.ast === false ? false : ( opts.ast || opts.body || (exclusionMode && opts.body !== false) )
    opts.user = opts.user === false ? false : ( opts.user || opts.body || (exclusionMode && opts.body !== false) )

    delete opts.body

    let _query = tokens.join(' ')

    if (opts.g) {
        opts.gizmoIds = findGizmoIdByTitle(_query)
        opts.title = true 
        opts.ast = opts.user = opts.gpt = false 
        return {query: '', opts}
    }

    return {query: _query, opts}
}

function parseAs(v: string, type: 'date' | 'int') {
    if (type === "date") {
        const d = parseDateAsYY(v)
        if (d) {
            return d.getTime() 
        } 
    } else if (type === "int") {
        const d = parseInt(v)
        if (d && !isNaN(d)) {
            return d 
        }
    }
}


const YY_REGEX = /^\d{4}[-\/:,\._]\d{1,2}[-\/:,\._]\d{1,2}$/;
function parseDateAsYY(v: string) {
    if (!v) return
    if (!YY_REGEX.test(v)) return
    return new Date(v)
}

export type PreFilter = {
    title?: boolean,
    body?: boolean,
    user?: boolean,
    ast?: boolean,

    dalle?: boolean
    browse?: boolean,
    python?: boolean,
    gizmo?: boolean,
    gizmos?: boolean,
    gpt4?: boolean,
    archived?: boolean,
    gpt?: boolean,
    g?: boolean,
    gg?: boolean,

    createdAfter?: number,
    createdBefore?: number,
    updatedAfter?: number,
    updatedBefore?: number 

    turnsPlus?: number,
    turnsMinus?: number,

    gizmoIds?: Set<string>
}


function findGizmoIdByTitle(query: string): Set<string> {
    const gizmos = getGizmosSync()
    return new Set((search(gizmos.map(g => g.name), query, {}) ?? []).map(s => `g-${gizmos[s.idx].id}`))
}
</file>

<file path="src/raccoon/searchChats/filterChats.ts">
import { Chat, ChatPart, MessagePart, PartResult, Result } from "../types"
import { extractOpts } from "./extractOpts"
import { getElapsed } from "../utils/getElapsed"
import { preFilterChats } from "./preFilter"
import { Context, getContext } from "./extractContext"
import { getGizmoByIdSync } from "../utils/gizmo"
import { search } from "./core"

const DELIM = /\s*&&\s*/

export function multiFilterChats(chats: Chat[], _query: string, contextLevel: number): Result[] {
    // filter chats  
    for (let query of _query.split(DELIM)) {
        chats = filterChats(chats, query, contextLevel)
    }

    return chats.map(c => chatToResult(c)) 
}



function filterChats(chats: Chat[], _query: string, contextLevel: number): Chat[] {
    if (!chats?.length) return []

    let {query, opts} = extractOpts(_query)
    const parts = preFilterChats(chats, opts)

    // after pre-filtering, there might not be any query left.
    query = query.trim()
    if (query === "") return rebuildChats(parts) 


    const solts = search(parts.map(c => c.content), query, {
        threshold: (gvar.config["g:strictSearch"] && !gvar.gsm._morpho) ? 40 : 20,
        ordered: !gvar.config["g:orderByDate"]
    })
    if (!solts.length) return []
    
    let newParts: ChatPart[] = []

    try {
        for (let solt of solts) {
            const part = parts[solt.idx]
            if (!(solt.rb || part.type === "title")) continue 
            let ctx: Context 
            if (contextLevel) {
                ctx = getContext(part.content, solt.lb, solt.rb, contextLevel)
            }
            newParts.push(part)
    
            part.result = {
                messageId: (part as MessagePart).messageId,
                ...(ctx ?? {})
            }
        }
    } catch (err) {
        console.error(err)
    }

    return rebuildChats(newParts)
}

function chatToResult(c: Chat): Result {

    let headerResult: PartResult
    const headerIdx = c.parts.findIndex(p => p.type === "title" || p.type === "gpt")
    if (headerIdx >= 0) {
        headerResult = c.parts.splice(headerIdx, 1)[0].result
    }

    let gizmoName: string 
    let gizmoImg: string 
    let gizmoId: string 

    if (c._gizmo) {
        gizmoName = c._gizmo.name,
        gizmoImg = c._gizmo.imageUrl,
        gizmoId = `g-${c._gizmo.id}`
    } else if (c.gizmoId) {
        let info = getGizmoByIdSync(c.gizmoId.slice(2))
        if (info) {
            gizmoName = info.name
            if (gvar.config["g:showImage"]) gizmoImg = info.imageUrl
        }
    }

    return {
        id: c.id,
        title: c._gizmo?.name ?? c.title,
        elapsed: c.updateTime ? getElapsed(c.updateTime, gvar.gsm?._lang ?? 'en') : null,
        parts: c.parts.map(c => c.result).filter(v => v),
        headerResult,
        gizmoId: gizmoId ?? c.gizmoId,
        isGpt: c._gizmo ? true : false,
        gizmoName,
        gizmoImg,
        currentNodeId: c.currentNodeId
    }
}

function rebuildChats(parts: ChatPart[]) {
    const chatSet: Set<Chat> = new Set() 
    const chats: Chat[] = []

    parts.forEach(part => {
        const chat = part.chat 
        if (!chatSet.has(chat)) {
            chat.parts = []
            chatSet.add(chat)
            chats.push(chat)
        } 

        // // if no result with +gg. 
        // if (part.type === 'message' && !part.result) {
        //     part.result = {
        //         messageId: part.messageId,
        //         needle: '',
        //         prefix: part.content.slice(0, 35).concat('...'),
        //         suffix: ''
        //     }
        // }
        
        chat.parts.push(part)
    })
    
    return chats
}
</file>

<file path="src/raccoon/searchChats/Grabby.ts">
import { QueueLock } from "../../helper"
import { Chats } from "../types"
import { fetchChats } from "../utils/fetchChats"
import { getTtl } from "../utils/getTtl"
import { getGizmosAsChats } from "../utils/gizmo"



export class Grabby {
    static lastFetch = new QueueLock()
    static lastFetchActive = new QueueLock()
    static maxPages: number = Infinity

    static fetchMap: {[key: string]: Promise<void>} = {}
    static bgFetchMap: {[key: string]: Promise<void>} = {}

    static requestFetch = async (page: number, oldChats?: Chats) => {
        const key = (oldChats ? "bgFetchMap" : "fetchMap") satisfies keyof typeof Grabby

        let prom = Grabby[key][page]
        const now = Date.now() 
        if (prom && prom.time > (now - 60_000 * 10)) {
            return prom 
        }
        Grabby[key][page] = this.fetchChatsAndSave(page, oldChats)
        Grabby[key][page].time = now 
        prom = Grabby[key][page]
        
        return prom 
    }
    static fetchChatsAndSave = async (page: number, oldChats?: Chats) => {
        if (oldChats) {
            await Grabby.lastFetch.wait(3_000)
        } else {
            await Grabby.lastFetchActive.wait(3_000)
        }
        
        const chats = await fetchChats(page, gvar.auth)
        Grabby.maxPages = chats.maxPages
        
        await gvar.setItems({[`o:c:${page.toFixed(0)}`]: chats})
        
        if (oldChats) await this.backgroundShift(chats, oldChats)
    }
    static backgroundShift = async (chats: Chats, oldChats: Chats) => {
        const current = new Set(chats.chats.map(c => c.id))
        const ghosts = oldChats.chats.filter(c => !current.has(c.id))
        const key = `o:c:${(chats.page + 1).toFixed(0)}`
        const nextPage = (await gvar.getItems(key))[key] as Chats 
        if (!nextPage) return 
        nextPage.chats = [...nextPage.chats, ...ghosts]
        await gvar.setItems({[key]: nextPage})
    }   
    static getCached = async (page: number) => {
        const key = `o:c:${page.toFixed(0)}`
        const cached = (await gvar.getItems([key]))[key] as Chats
        if (cached) {
            if (cached.indexed < Date.now() - getTtl(page)) Grabby.requestFetch(page, cached)
            return cached 
        }
    }
    static getPage: (page: number) => Promise<Chats> = async page => {
        if (page === -1) {
            return {chats: await getGizmosAsChats(), hasMore: true} as Chats
        }

        // cached 
        let cached = await Grabby.getCached(page)
        if (cached) return cached 

        await Grabby.requestFetch(page)
        cached = await Grabby.getCached(page)
        if (cached) return cached 

        return {page, chats: [], hasMore: false, indexed: Date.now()} as Chats 
    }
}
</file>

<file path="src/raccoon/searchChats/index.ts">
import { Result, Status } from "../types"
import { multiFilterChats } from "./filterChats"
import { produce } from "immer"
import { Grabby } from "./Grabby"
import debounce from "lodash.debounce"

declare global {
    interface Promise<T> {
      time?: number
    }
}


export class SearchChats {
    released = false 
    page = -2
    latestResults: Result[] = []
    locked = false 
    errorCount = 0
    throttleTimeoutId: number
    context = 2
    processed: Set<string> = new Set()
    initAwait: Promise<void> 

    constructor(private query: string, private setStatus: (status: Status) => void, private mainRef: React.MutableRefObject<HTMLDivElement>) {
        SearchChats.ref = this
        this.initAwait = this.start()
    }
    start = async () => {
        if (this.released) return
        this.context = gvar.config["g:context"] ?? 2
        this.loadSafely()
    }  
    next = async (replace?: boolean) => {
        this.page++
        const data = await Grabby.getPage(this.page)
        if (this.released) return

        if (replace) {
            this.latestResults = []
            this.mainRef.current?.scrollTo({left: 0, top: 0, behavior: 'instant'})  
        } else {
            this.setStatus?.({results: this.latestResults})
        }


        // ignore duplicates 
        data.chats = data.chats.filter(c => {
            if (this.processed.has(c.id)) return false 
            this.processed.add(c.id)
            return true 
        })

        let res: Result[]
        try {
            const chats = [...data.chats]
            res = multiFilterChats(chats, this.query, this.context) ?? []
        } catch (err) {
            console.error('ERROR', err)
            throw err
        }
        

        this.latestResults = produce(this.latestResults, d => {
            d.push(...res)
        })
 
        this.setStatus?.({
            results: this.latestResults,
            finished: !data.hasMore
        })
    }
    loadSafely = debounce(async () => {
        await this.initAwait
        if (this.released || this.locked || (this.page + 1 >= Grabby.maxPages) || this.errorCount > 10) return 

        this.locked = true 
        this.next().then(() => {
            this.errorCount = 0
            this.locked = false
        }, err => {
            this.errorCount++
            this.locked = false
        })
    }, 50, {trailing: true, leading: true, maxWait: 50})
    release = () => {
        if (this.released) return 
        delete this.setStatus
        this.released = true 
        delete SearchChats.ref
        clearTimeout(this.throttleTimeoutId)
    }
    static ref: SearchChats
}
</file>

<file path="src/raccoon/searchChats/preFilter.ts">
import { Chat, ChatPart } from "../types"
import { PreFilter } from "./extractOpts"

export function preFilterChats(chats: Chat[], opts: PreFilter): ChatPart[] {

    if ((opts as any).archive != null) {
        opts.archived = (opts as any).archive
    }

    chats = chats.filter(c => {
        if (opts.dalle != null && opts.dalle != !!c.usedDalle) return false 
        if (opts.browse != null && opts.browse != !!c.usedBrowser) return false 
        if (opts.python != null && opts.python != !!c.usedPython) return false 
        if (opts.gizmo != null && opts.gizmo != !!c.gizmoId) return false 
        if (opts.gizmos != null && opts.gizmos != !!c.multiGizmo) return false 
        if (opts.gpt4 != null && opts.gpt4 != !!c.usedGPT4) return false 
        if (opts.archived != null && opts.archived != !!c.isArchived) return false 

        if (opts.createdAfter && (c.createTime < opts.createdAfter || c.createTime == null)) return false
        if (opts.createdBefore && (c.createTime > opts.createdBefore || c.createTime == null)) return false
        if (opts.updatedAfter && (c.updateTime < opts.updatedAfter || c.updateTime == null)) return false
        if (opts.updatedBefore && (c.updateTime > opts.updatedBefore || c.updateTime == null)) return false

        if (opts.turnsPlus && (c.childCount < opts.turnsPlus || c.childCount == null)) return false
        if (opts.turnsMinus && (c.childCount > opts.turnsMinus || c.childCount == null)) return false
        if (opts.gizmoIds && !opts.gizmoIds.has(c.gizmoId)) return false 
        
        return true 
    })

    const messageParts: ChatPart[] = []

    
    chats.forEach(c => {
        if (opts.gpt && c._gizmo) {
            messageParts.push({
                type: 'gpt',
                content: c._gizmo.name,
                chat: c
            }) 
        } else if (opts.title && !c._gizmo) {
            messageParts.push({
                type: 'title',
                content: c.title,
                chat: c
            }) 
        } 

        c.userChilds.forEach(v => v.chat = c)
        c.astChilds.forEach(v => v.chat = c)
        
        if (opts.user) messageParts.push(...c.userChilds)
        if (opts.ast) messageParts.push(...c.astChilds)
    })

    return messageParts
}
</file>

<file path="src/raccoon/style.css">
#App {

    --bg: var(--sidebar-surface-primary, #f9f9f9);
    --bg-secondary: var(--sidebar-surface-secondary, #ececec);

    &.new {
        --bg: var(--sidebar-surface-secondary, #ececec);
        --bg-secondary: var(--sidebar-surface-tertiary, #e3e3e3);
    }

    --border-light: rgba(0,0,0,.1);
    --border-medium: rgba(0,0,0,.15);
    --border-heavy: rgba(0,0,0,.2);
    --border-xheavy: rgba(0,0,0,.25);

    --text-color-primary: #161616;
    --text-color-primary-off: #303030;
    --text-color-secondary: #535353;
    --text-color-tertiary: #626262;
    --text-color-quaternary: #868686;
    --context-color: #008de5; 
    
    --heart: #ff0000;
    --star: #0031ff;
    --loading: red;
    
    --needle-color: var(--context-color);

    &.dark {
        --bg: var(--sidebar-surface-primary, #0d0d0d);
        --bg-secondary: var(--sidebar-surface-secondary, #262626);
        
        &.new {
            --bg: var(--sidebar-surface-secondary, #262626);
            --bg-secondary: var(--sidebar-surface-tertiary, #333);
        }

        --border-light: rgba(255, 255, 255,.1);
        --border-medium: rgba(255, 255, 255,.15);
        --border-heavy: rgba(255, 255, 255,.2);
        --border-xheavy: rgba(255, 255, 255,.25);
    
        --text-color-primary: #fbfbfb;
        --text-color-primary-off: #f2f2f2;
        --text-color-secondary: #cbcbcb;
        --text-color-tertiary: #a8a8a8;
        --text-color-quaternary: #8f8f8f;
        --context-color: #6dffd8; 
        
        --heart: #ff0000;
        --star: #ffd900;
    }

    -webkit-font-smoothing: antialiased;
    user-select: none;

    &.bold .needle {
        font-weight: bold;
    }
    
    &.peacock {
        border-radius: 10px;
        top: 60px;
        left: 12px;
        z-index: 999999;
        width: 350px;
        position: fixed;
        background-color: var(--bg);
        outline: 1px solid var(--border-medium);  
        
        .search {
            margin-top: 0px;
            margin-bottom: 10px;
        }        
    }

    & > .search {

        & > .closeIcon {
            position: absolute;
            right: 10px;
            cursor: pointer;
        }
    }

    & > .main {
        overflow-y: auto;
        max-height: 60vh;
        font-size: 0.95rem;
        font-weight: 400;
        color: var(--text-color);

        & > *:last-child {
            margin-bottom: 40px;
        }
    }
    
    & > .footer {
        padding: 10px 10px;
        font-size: 0.9em;
        font-weight: 300;

        display: grid;
        grid-template-columns: max-content 1fr max-content;
        align-items: center;
        justify-items: start;
        column-gap: 10px;
        color: var(--text-color-secondary);

        & > .svgButton {
            font-size: 1.2em;
            
            

            &.pin {
                transform: scale(1.15) translate(-1px, 1px) rotate(30deg);
            }
            
            &.heart, &.star {
                transition: color 0.08s ease-in;

                &.heart:hover {
                    color: var(--heart);
                }
                &.star:hover {
                    color: var(--star);
                }
            }

        }
    }

    & > .scale {
        width: 20px;
        height: 20px;
        background-color: transparent;
        position: absolute;
        right: -7px;
        bottom: -5px;
        cursor: grab;
        border-radius: 50%;
    }
}

.svgButton {
    background-color: transparent;
    padding: 0px;
    border: none;
    cursor: pointer;
    color: var(--text-color-secondary);

    &:hover {
        /* opacity: 0.75; */
        color: var(--text-color-primary);
    }

    &.toggable {
        opacity: 0.75;
    }

    &.active {
        opacity: 1;
        color: var(--text-color-primary);
    }
}

.ResultItem {
    cursor: pointer;


    & > .header {
        padding: 10px;
        padding-top: 20px;
        padding-bottom: 5px;

        &:first-child {
            padding-top: 10px;
        }

        &:hover > .meta > img {
            /* opacity: 1; */
        }

        & > .meta {
            font-weight: 300;
            font-size: 0.9em;

            overflow-x: hidden;
            text-overflow: ellipsis;
            text-wrap: nowrap;

            & > span {
                opacity: 0.75;
            }

            & > img {
                border-radius: 25%;
                margin-right: 3px;
                margin-left: 8px;
            }
        }
        

        & > .title > img {
            margin-right: 10px;
            border-radius: 50%;
        }
        
        & img {
            vertical-align: middle;
            border: 1px solid var(--border-xheavy);
        }
    }

    & > .header, & > .conti .context {

        &:hover {
            background-color: var(--bg-secondary);
        }
    }


    span.needle {
        color: var(--needle-color);
    }

    & > .conti {
        padding-left: 10px;

        & > .context {
            padding: 8px;
            white-space: pre;
            font-size: 0.9em;
            text-wrap: wrap;
            color: var(--text-color-primary-off);
            border-top: 1px solid var(--border-medium);

            &:first-child {
                border-top: none;
            }
    
            & > span {
                word-wrap: break-word;
            }
        }
    }


}


@keyframes oscillateSize {
    0%, 100% { transform: scale(0.8); }
    50% { transform: scale(1.1); }
  }

.LoadMore {
    display: grid;
    align-items: center;
    justify-content: center;
    padding: 30px 0;

    &::after {
        content: '';
        width: 5em;
        height: 0.2em;
        background-color: var(--text-color-primary);
        animation: oscillateSize 2s infinite ease-in-out;
    }
}

.NonFound {
    display: grid;
    justify-items: center;
    justify-content: center;
    padding: 30px 0;

    &::after {
        content: '';
        width: 10em;
        height: 0.1em;
        background-color: var(--loading);
        animation: oscillateSize 50s infinite ease-in-out;
    }
}




::-webkit-scrollbar {
    width: 6px; 
}

::-webkit-scrollbar-track {
    background: var(--bg);
}

::-webkit-scrollbar-thumb {
    background: var(--bg-secondary);
    border-radius: 2px;
}
</file>

<file path="src/raccoon/types.ts">
import { Gizmo } from "../types"

export type Status = {
    results: Result[],
    finished?: boolean

}

export type Chats = {
    page: number,
    chats: Chat[],
    indexed: number,
    hasMore: boolean,
    maxPages: number
}

export type Chat = {
    id: string,
    gizmoId?: string,
    gizmoIds?: string[],
    multiGizmo?: boolean,
    usedDalle?: boolean,
    usedBrowser?: boolean,
    usedPython?: boolean,
    usedGPT4?: boolean,
    childCount?: number, // ast + user 

    createTime: number,
    updateTime: number,
    title: string,
    isArchived: boolean,
    userChilds: MessagePart[],
    astChilds: MessagePart[],
    parts?: ChatPart[],
    currentNodeId?: string,
    
    _gizmo?: Gizmo
}

export type MessagePart = {
    type: 'message',
    messageId: string,
    byAst?: boolean,
    content?: string,
    chat?: Chat,
    result?: PartResult
}

export type TitlePart = {
    type: 'title',
    content?: string,
    chat?: Chat,
    result?: PartResult
}

export type GptPart = {
    type: 'gpt',
    content?: string 
    chat?: Chat,
    result?: PartResult,
    gptId?: string 
}

export type ChatPart = MessagePart | TitlePart | GptPart


export type LoadRequest = {
    page?: number,
    cached?: Chats 
}


export type Result = {
    title: string,
    id: string,
    elapsed?: string,
    parts: PartResult[],
    headerResult?: PartResult,
    isGpt?: boolean, 
    currentNodeId?: string,

    gizmoId?: string,
    gizmoImg?: string
    gizmoName?: string 
}

export type PartResult = {
    messageId?: string,
    prefix?: string
    suffix?: string
    needle?: string 
}
</file>

<file path="src/raccoon/utils/bump.tsx">
export async function bump(chatId: string, auth: string) {
    const title = await getTitle(chatId, auth);
    if (!title) throw 'No title';

    const res = await fetch(`https://chatgpt.com/backend-api/conversation/${chatId}`, {
        method: 'PATCH',
        body: JSON.stringify({ title: title }),
        headers: {
            'Authorization': auth,
            'Content-Type': 'application/json'
        }
    });
    if (!res.ok) throw 'Bump not OK';
}

async function getTitle(chatId: string, auth: string) {
    const json = await (await fetch(`https://chatgpt.com/backend-api/conversation/${chatId}`, {
        method: 'GET',
        headers: {
            'Authorization': auth
        }
    })).json()
    return json.title 
}
</file>

<file path="src/raccoon/utils/extractChats.ts">
import { Chat } from "../types"
import { ConversationInterface } from "./rawTypes"

export function extractChat(json: ConversationInterface) {
    let chat: Partial<Chat> = {}
    chat.id = json.id ?? json.conversation_id
    chat.gizmoId = json.gizmo_id
    chat.createTime = json.create_time ? new Date(json.create_time).getTime() : null 
    chat.updateTime = json.update_time ? new Date(json.update_time).getTime() : null 
    chat.title = json.title
    chat.isArchived = json.is_archived
    chat.currentNodeId = typeof json.current_node === "string" ? json.current_node : null 
    
    chat.astChilds = []
    chat.userChilds = []

    const gizmoIds = new Set<string>()

    for (let [_, mapping] of Object.entries(json.mapping ?? {})) {
        const m = mapping.message
        if (!m) continue 

        if (m.metadata) {
            if (m.metadata.gizmo_id) gizmoIds.add(m.metadata.gizmo_id)
            if (m.metadata.model_slug?.startsWith('gpt-4')) chat.usedGPT4 = true 
        }

        if (m.author?.role === 'tool') {
            if (m.author.name === 'dalle.text2im') chat.usedDalle = true 
            if (m.author.name === 'python') chat.usedPython = true 
            if (m.author.name === 'browser') chat.usedBrowser = true 
        }

        // text extraction 
        if (!(m.id && m.author && m.content?.parts) || m.metadata?.is_visually_hidden_from_conversation) continue 
        if (!(m.content.content_type === "text" || m.content.content_type === "multimodal_text")) continue 
        if (!(m.author.role === "user" || m.author.role === "assistant")) continue 

        const isUser = m.author.role === "user" 
        const texts: string[] = []
        m.content.parts.forEach(part => {
            if (typeof part === "string") texts.push(part)
        })
        const text = texts.join('\n\n').trim()

        text.length && (
            chat[isUser ? 'userChilds' : 'astChilds'].push({
                type: 'message',
                content: text,
                messageId: m.id,
                byAst: m.author.role === "assistant"
            })
        )
    }

    chat.gizmoIds = [...gizmoIds]

    if (chat.gizmoIds.length === 1) {
        chat.gizmoId = chat.gizmoIds[0]
    } else if (chat.gizmoIds.length > 1) {
        chat.multiGizmo = true 
    }

    chat.childCount = chat.astChilds.length + chat.userChilds.length

    return chat as Chat 
}
</file>

<file path="src/raccoon/utils/fetchChats.ts">
import { Chat, Chats } from "../types"
import { extractChat } from "./extractChats"


export async function fetchChats(page: number, auth: string) {
    const res = await fetch(`https://chatgpt.com/backend-api/conversations?offset=${page * 100}&limit=100&order=updated&expand=true`, {
        headers: {
            'Authorization': auth
        }
    })
    if (!res.ok) throw Error('Failed')
    const json = await res.json()
    const chats: Chat[] = []
    for (let item of json.items) {
        try {
            const chat = extractChat(item)
            if (chat) chats.push(chat)
        } catch (err) { }
    }

    return {
        chats,
        indexed: Date.now(),
        page,
        hasMore: (page + 1) * 100 < json.total,
        maxPages:  Math.ceil(json.total / 100)
    } satisfies Chats
}
</file>

<file path="src/raccoon/utils/getElapsed.tsx">
export function getElapsed(rb: number, locale: string) {
    const d = new Date(rb);
    const sameYear = new Date().getFullYear() === d.getFullYear();

    const r = new Intl.RelativeTimeFormat(locale, { numeric: 'auto' });
    const seconds = (Date.now() - rb) / 1000;
    const minutes = seconds / 60;
    const hours = minutes / 60;
    const days = hours / 24;
    const years = days / 365;
    const weeks = years * 52;

    if (days < 14) {
        return r.format(-Math.round(days), 'days');
    } else if (weeks < 4) {
        return r.format(-Math.round(weeks), 'weeks');
    } else if (days < 60) {
        return d.toLocaleDateString(locale, { month: 'long', day: 'numeric' });
    } else if (sameYear) {
        return d.toLocaleDateString(locale, { month: 'long' });
    } else {
        return d.toLocaleDateString(locale, { month: 'long', year: 'numeric' });
    }
}
</file>

<file path="src/raccoon/utils/getTtl.ts">
export function getTtl(page: number) {
    if (page < 1) return 60_000 * 10
    else if (page < 2) return 60_000 * 60 * 6 
    else if (page < 4) return 60_000 * 60 * 24 * 2
    else if (page < 12) return 60_000 * 60 * 24 * 6
    else if (page < 24) return 60_000 * 60 * 24 * 15
    else if (page < 48) return 60_000 * 60 * 24 * 30
}
</file>

<file path="src/raccoon/utils/gizmo.ts">
import { Gizmo } from "../../types"
import { getLocalItem } from "../../utils/getKnown"
import { Chat } from "../types"

let gizmos: {[key: string]: Gizmo}

export async function getGizmoById(id: string) {
    await ensureLoaded()
    if (!gizmos) return 
    return gizmos[id]  
}

export function getGizmoByIdSync(id: string) {
    return gizmos[id]  
}

export function getGizmosSync() {
    return Object.entries(gizmos ?? {}).map(v => v[1])
}

async function ensureLoaded() {
    if (!gizmos) {
        gizmos = await getLocalItem('o:gizmos') || {}
    }
}

export async function getGizmosAsChats(): Promise<Chat[]> {
    await ensureLoaded()
    if (!gizmos) return []

    return Object.entries(gizmos).map(([k, g]) => ({
        astChilds: [],
        userChilds: [],
        title: g.name,
        id: g.id,
        gizmoId: `g-${g.id}`,
        _gizmo: g 
    })) as Chat[]
}
</file>

<file path="src/raccoon/utils/misc.tsx">
export function softLink(path: string, blockScroll?: number) {
    window.dispatchEvent(new CustomEvent('busbusab', {detail: JSON.stringify({type: 'NAV', path, blockScroll}), bubbles: false}))
}

export function blockScroll(blockScroll: number) {
    window.dispatchEvent(new CustomEvent('busbusab', {detail: JSON.stringify({type: 'BLOCK_SCROLL', blockScroll}), bubbles: false}))
}
</file>

<file path="src/raccoon/utils/rawTypes.ts">
export type ConversationInterface = {
    title: string,
    create_time: TimeSeconds, 
    update_time: TimeSeconds,
    mapping: {[key: Id]: Mapping},
    current_node: Id,
    id?: Id,
    conversation_id?: Id,
    conversation_template_id?: string,
    gizmo_id?: string,
    is_archived?: boolean
}

type Id = string 
type TimeSeconds = number 

type Mapping = {
    id: Id, 
    message: Message,
    parent: Id,
    children: Id[]
}

type Message = {
    id: Id, 
    author: Author,
    create_time: null,
    update_time: null,
    content: Content,
    status: "finished_successfully" | string,
    end_turn: boolean,
    weight: number,
    metadata: MessageMetadata,
    recipient: "all" | string 
}

type Author = {
    role: "system" | "user" | "assistant" | "tool"
    name: string,
    metadata: {}
}

type Content = TextContent | MultiModalContent

type TextContent = {
    content_type: "text",
    parts: string[]
}

type MultiModalContent = {
    content_type: "multimodal_text",
    parts: (MultiModalContentPart | string)[]
}

type MultiModalContentPart = {
    content_type: "image_asset_pointer" | string,
    asset_pointer: string,
    size_bytes: number,
    width?: number,
    height?: number
}

type MessageMetadata = {
    is_visually_hidden_from_conversation: boolean,
    gizmo_id?: string,
    model_slug?: string 
}
</file>

<file path="src/types.ts">
export type Config = {
    'g:version': number,
    'g:lang': string,
    'g:context': number,
    'g:autoClear': boolean,
    'g:highlightColorDark': string,
    'g:highlightColorLight': string,
    'g:highlightUnderline': boolean,
    'g:highlightBold': boolean,
    'g:showImage': boolean,
    'g:sessionOnly': boolean,
    'g:orderByDate': boolean,
    'g:scrollTop': boolean,
    'g:strictSearch': boolean,
}

export type TempState = {
    'o:lastTheme': string,
    'o:auth': string,
    'o:ph': string,
    'o:changeId': string,
    'o:gizmos': {
        [key: string]: Gizmo 
    }
}

export type SessionState = {
    's:auth': string
}

export type Gizmo = {
    id: string,
    name: string,
    added: number,
    imageUrl: string 
}

export type LocalState = Config & TempState

export const CONFIG_KEYS = ['g:version', 'g:lang', 'g:context', 'g:autoClear', 'g:highlightColorDark', 'g:highlightColorLight', 'g:highlightUnderline', 'g:highlightBold', 'g:showImage', 'g:sessionOnly', 'g:orderByDate', 'g:scrollTop', 'g:strictSearch'] as const

export const TEMP_KEYS = ['o:lastTheme', 'o:auth', 'o:ph', 'o:changeId', 'o:gizmos'] as const 

export const KNOWN_KEYS = [...CONFIG_KEYS, ...TEMP_KEYS] as const 


export type AnyDict = {[key: string]: any}

export type StringDict = {[key: string]: string}
</file>

<file path="src/utils/browser.ts">
import { AnyDict } from "../types"

export function openLink(url: string, active = true) {
    chrome.runtime.sendMessage({type: "REQUEST_CREATE_TAB", url, active})
}

export function localGet(keys: chrome.storage.StorageGet) {
    return chrome.storage.local.get(keys)
}

export function localSet(items: AnyDict) {
    return chrome.storage.local.set(items)
}

export function sessionGet(keys: chrome.storage.StorageGet) {
    return chrome.storage.session.get(keys)
}

export function sessionSet(items: AnyDict) {
    return chrome.storage.session.set(items)
}

async function requestSessionGet(keys: chrome.storage.StorageGet) {
    const res = await chrome.runtime.sendMessage({type: 'GET_SESSION_ITEM', keys})
    if (res?.error) throw res.error 
    return res?.ok as AnyDict
}

async function requestSessionSet(items: AnyDict) {
    const res = await chrome.runtime.sendMessage({type: 'SET_SESSION_ITEM', items})
    if (res?.error) throw res.error 
    return res?.ok as void 
}

export const sessionGetFallback = chrome.storage.session?.setAccessLevel ? sessionGet : requestSessionGet
export const sessionSetFallback = chrome.storage.session?.setAccessLevel ? sessionSet : requestSessionSet
</file>

<file path="src/utils/getKnown.ts">
import { LocalState, SessionState } from "../types";

export async function getLocal(keys: (keyof LocalState)[]) {
    return (await chrome.storage.local.get(keys)) as LocalState 
}

export async function getLocalItem<T extends keyof LocalState>(key: T) {
    return (await chrome.storage.local.get(key))[key] as LocalState[T]
}

export async function setLocal(override: Partial<LocalState>) {
    await chrome.storage.local.set(override)
}

export async function setSessionKnown(override: Partial<SessionState>) {
    await chrome.storage.session.set(override)
}

export async function getSession(keys: (keyof SessionState)[]) {
    return (await chrome.storage.session.get(keys)) as SessionState 
}

export async function getSessionItem<T extends keyof SessionState>(key: T) {
    return (await chrome.storage.session.get(key))[key] as SessionState[T]
}
</file>

<file path="src/utils/gsm.ts">
import { Gsm } from "./GsmType";

declare global {
  interface GlobalVar {
    gsm?: Gsm
}
}

export async function loadGsm(): Promise<Gsm> {
  const language = (await chrome.storage.local.get("g:lang"))["g:lang"]
  return readLocaleFile(getValidLocale(language))
}

export async function requestGsm(): Promise<Gsm> {
  return chrome.runtime.sendMessage({type: "REQUEST_GSM"})
}

export async function readLocaleFile(locale: string): Promise<Gsm> {
  const fetched = await fetch(chrome.runtime.getURL(`locales/${locale}.json`))
  const json = await fetched.json() as Gsm 
  json._lang = locale.replace("_", "-") 
  return json 
}

function getValidLocale(overrideLang?: string) {
  if (overrideLang && AVAILABLE_LOCALES.has(overrideLang)) return overrideLang
  const languages = new Set(navigator.languages.map(l => l.replace("-", "_")))
  languages.forEach(l => {
    if (l.includes("_")) {
      const langPart = l.split("_")[0]
      languages.add(langPart)
    }
  })
  languages.add("en")
  return [...languages].find(l => AVAILABLE_LOCALES.has(l))
}

export function replaceArgs(raw: string, args: string[]) {
  let idx = 0
  for (let arg of args) {
    raw = raw.replaceAll(`$${++idx}`, arg)
  }
  return raw 
}

export const LOCALE_MAP: {
  [key: string]: {
    display: string,
    title: string
  }
} = {
  "detect": {display: "Auto", title: "Try to find a match using browser language settings, system language settings, or fallback to English."},
  "en": { display: "English", title: "English" },
  "es": { display: "Español", title: "Spanish" },
  "it": { display: "Italiano", title: "Italian" },
  "ja": { display: "日本語", title: "Japanese" },
  "ko": { display: "한국어", title: "Korean" },
  "pt_BR": { display: "Português", title: "Portuguese" },
  "ru": { display: "Русский", title: "Russian" },
  "tr": { display: "Türkçe", title: "Turkish" },
  "vi": { display: "Tiếng Việt", title: "Vietnamese" },
  "zh_CN": { display: "中文 (简体)", title: "Chinese (Simplified)" },
  "zh_TW": { display: "中文 (繁體)", title: "Chinese (Traditional)" }
}


const AVAILABLE_LOCALES = new Set(["en", "es", "it", "ja", "ko", "pt_BR", "ru", "tr", "vi", "zh_CN", "zh_TW"])
</file>

<file path="src/utils/GsmType.ts">
export type Gsm = {
  _lang?: string,
  _morpho?: boolean,
  general: string,
  help: string,
  strictSearch: string,
  strictSearchTooltip: string,
  sortDate: string,
  sortDateTooltip: string,
  scrollTop: string,
  scrollTopTooltip: string,
  issuePrompt: string,
  issueDirective: string,
  areYouSure: string,
  showMore: string,
  notFound: string,
  data: string,
  sessionOnly: string,
  sessionOnlyTooltip: string,
  showImage: string,
  showImageTooltip: string,
  searchChats: string,
  search: string,
  language: string,
  context: string,
  contextTooltip: string,
  reset: string,
  clearCache: string,
  autoClear: string,
  autoClearTooltip: string,
  enableShortcut: string,
  enableShortcutTooltip: string,
  highlightColor: string,
  highlightColorTooltip: string,
  highlightBold: string,
  highlightBoldTooltip: string,
  advancedSearch: string,
  otherProjects: {
    header: string,
    askScreenshot: string,
    globalSpeed: string
  }
}
</file>

<file path="src/utils/state.ts">
import debounce from "lodash.debounce"
import { AnyDict } from "../types"
import type { DebouncedFunc } from "lodash"
import { randomId } from "../helper"

export type SubStorageCallback = (view: AnyDict, forOnLaunch?: boolean) => void

export class SubscribeStorageKeys {
    keys: Set<string>
    cbs: Set<SubStorageCallback> = new Set()
    private rawMap?: AnyDict
    latestRaw?: AnyDict
    released = false 
    processedChangeIds: Set<string> = new Set() 

    constructor(_keys: string[], private onLaunch?: boolean, cb?: SubStorageCallback, public wait?: number, public maxWait?: number) {
        this.triggerCbs = this.wait ? (
            debounce(this._triggerCbs, this.wait ?? 0, {trailing: true, leading: true, ...(this.maxWait == null ? {} : {maxWait: this.maxWait})})
        ) : this._triggerCbs

        this.keys = new Set(_keys)
        cb && this.cbs.add(cb)
        this.start()
    }
    start = async () => {
        chrome.storage.local.onChanged.addListener(this.handleChange)
        if (this.onLaunch) {
            await this.handleChange(null)
        }
    }
    release = () => {
        if (this.released) return 
        this.released = true 
        ;(this.triggerCbs as any).cancel?.()
        delete this.triggerCbs
        chrome.storage.local.onChanged.removeListener(this.handleChange)
        this.cbs.clear()
        delete this.cbs, delete this.rawMap, 
        delete this.latestRaw, delete this.keys, delete this.rawMap
    }
    handleChange = async (changes: chrome.storage.StorageChanges) => {
        changes = changes ?? {} 
        const changeId = changes["o:changeId"]?.newValue as string
        if (changeId) {
            if (this.processedChangeIds.has(changeId)) {
                this.processedChangeIds.delete(changeId)
                return
            }  else {
                this.processedChangeIds.add(changeId)
            }
        }

        let hadChanges = false 
        if (!this.rawMap) {
            this.rawMap = await chrome.storage.local.get([...this.keys])
            hadChanges = true 
        }
        for (let key in changes) {
            if (!this.keys.has(key)) continue 
            this.rawMap[key] = changes[key].newValue
            if (this.rawMap[key] === undefined) delete this.rawMap[key]
            hadChanges = true 
        }
        if (!hadChanges) return 

        this.latestRaw = structuredClone(this.rawMap)
        this.triggerCbs()
    }
    _triggerCbs = () => {
        this.cbs.forEach(cb => cb(this.latestRaw))
    }
    triggerCbs: typeof this._triggerCbs | DebouncedFunc<typeof this._triggerCbs>
    
    push = (_override: AnyDict) => {
        const override = structuredClone(_override)
        override["o:changeId"] = randomId()

        const changes = {} as chrome.storage.StorageChanges
        for (let key in override) {
            if (override[key] === undefined) continue 
            changes[key] = {newValue: override[key]} 
        }

        return Promise.all([
            this.handleChange(changes),
            chrome.storage.local.set(override)
        ])
    }
}
</file>

<file path="static/_locales/en/messages.json">
{
  "appName": {
    "message": "GPT Search: Chat History"
  },
  "appDesc": {
    "message": "Search your ChatGPT conversation history."
  }
}
</file>

<file path="static/_locales/es/messages.json">
{
  "appName": {
    "message": "GPT Search: Buscar chats"
  },
  "appDesc": {
    "message": "Busca en tu historial de conversaciones de GPT."
  }
}
</file>

<file path="static/_locales/it/messages.json">
{
  "appName": {
    "message": "GPT Search: Cerca chat"
  },
  "appDesc": {
    "message": "Cerca nel tuo storico conversazioni GPT."
  }
}
</file>

<file path="static/_locales/ja/messages.json">
{
  "appName": {
    "message": "GPT Search: チャットを検索"
  },
  "appDesc": {
    "message": "あなたのGPT会話履歴を検索します。"
  }
}
</file>

<file path="static/_locales/ko/messages.json">
{
  "appName": {
    "message": "GPT Search: 채팅 검색"
  },
  "appDesc": {
    "message": "당신의 GPT 대화 기록을 검색하세요."
  }
}
</file>

<file path="static/_locales/pt_BR/messages.json">
{
  "appName": {
    "message": "GPT Search: Pesquisar chats"
  },
  "appDesc": {
    "message": "Pesquise no seu histórico de conversas GPT."
  }
}
</file>

<file path="static/_locales/ru/messages.json">
{
  "appName": {
    "message": "GPT Search: Поиск чатов"
  },
  "appDesc": {
    "message": "Искать в истории ваших бесед GPT."
  }
}
</file>

<file path="static/_locales/tr/messages.json">
{
  "appName": {
    "message": "GPT Search: Sohbetleri Ara"
  },
  "appDesc": {
    "message": "GPT sohbet geçmişinizi arayın."
  }
}
</file>

<file path="static/_locales/vi/messages.json">
{
  "appName": {
    "message": "GPT Search: Tìm kiếm cuộc trò chuyện"
  },
  "appDesc": {
    "message": "Tìm kiếm lịch sử trò chuyện GPT của bạn."
  }
}
</file>

<file path="static/_locales/zh_CN/messages.json">
{
  "appName": {
    "message": "GPT Search: 搜索聊天"
  },
  "appDesc": {
    "message": "搜索您的GPT对话历史。"
  }
}
</file>

<file path="static/_locales/zh_TW/messages.json">
{
  "appName": {
    "message": "GPT Search: 搜尋聊天"
  },
  "appDesc": {
    "message": "搜尋您的GPT對話歷史。"
  }
}
</file>

<file path="static/locales/en.json">
{
  "_lang": "en",
  "_morpho": false,
  "general": "General",
  "help": "Help",
  "strictSearch": "Strict search",
  "strictSearchTooltip": "Increase threshold for matching results.",
  "sortDate": "Sort by date",
  "sortDateTooltip": "Sort results by date.",
  "scrollTop": "Scroll to top",
  "scrollTopTooltip": "Scroll to the top when selecting a conversation.",
  
  "issuePrompt": "Have issues or a suggestion?",
  "issueDirective": "Provide feedback on Github.",
  "areYouSure": "Are you sure?",
  "showMore": "Show more",
  "notFound": "Not found",

  "data": "Data",
  "sessionOnly": "Session-only caching",
  "sessionOnlyTooltip": "Clear cached chat history on browser close. This increases search times and data usage, so use only if your device is shared.",
  "showImage": "Show icon",
  "showImageTooltip": "If a conversation used a custom GPT, show the icon if available.",
  "searchChats": "Search chats...",
  "search": "Search",
  "language": "Language",
  "context": "Context",
  "contextTooltip": "Adjust how much surrounding information is displayed with each search result.",
  "reset": "Reset",
  "clearCache": "Clear cache",
  "autoClear": "Auto-clear search",
  "autoClearTooltip": "Clear search field when you click away.",
  "enableShortcut": "Enable shortcut",
  "enableShortcutTooltip": "Press '/' to quickly focus on the search box.",

  
  "highlightColor": "Highlight color",
  "highlightColorTooltip": "Adjust the appearance of highlighted text.",

  "highlightBold": "Bold highlight",
  "highlightBoldTooltip": "Makes highlighted text bold for better visibility.",
  "advancedSearch": "Advanced search",

  "otherProjects": {
    "header": "Other Projects",
    "askScreenshot": "Take a screenshot on any page and automatically open it with ChatGPT. <i>Also available for <a href=\"https://youtu.be/0g_EQbp5q4g\">Claude</a> and <a href=\"https://youtu.be/I-VX31SoGZ0\">Gemini</a>.</i>",
    "globalSpeed": "Set a default speed for video and audio!"
  }
}
</file>

<file path="static/locales/es.json">
{
  "general": "General",
  "help": "Ayuda",
  "strictSearch": "Búsqueda estricta",
  "strictSearchTooltip": "Aumentar el umbral para los resultados coincidentes.",
  "sortDate": "Ordenar por fecha",
  "sortDateTooltip": "Ordenar los resultados por fecha.",
  "scrollTop": "Desplazarse hacia arriba",
  "scrollTopTooltip": "Desplácese hacia arriba al seleccionar una conversación.",
  "issuePrompt": "¿Tienes problemas o una sugerencia?",
  "issueDirective": "Proporciona comentarios en Github.",
  "areYouSure": "¿Estás seguro?",
  "showMore": "Mostrar más",
  "notFound": "No encontrado",
  "data": "Datos",
  "sessionOnly": "Caché solo durante la sesión",
  "sessionOnlyTooltip": "Borrar el historial de chat almacenado en caché al cerrar el navegador. Esto aumenta los tiempos de búsqueda y el uso de datos, por lo que se debe usar solo si el dispositivo es compartido.",
  "showImage": "Mostrar icono",
  "showImageTooltip": "Si una conversación utilizó un GPT personalizado, mostrar el icono si está disponible.",
  "searchChats": "Buscar chats...",
  "search": "Buscar",
  "language": "Idioma",
  "context": "Contexto",
  "contextTooltip": "Ajusta cuánta información circundante se muestra con cada resultado de búsqueda.",
  "reset": "Restablecer",
  "clearCache": "Borrar caché",
  "autoClear": "Borrado automático de búsqueda",
  "autoClearTooltip": "Borra el campo de búsqueda cuando haces clic fuera.",
  "enableShortcut": "Habilitar atajo",
  "enableShortcutTooltip": "Presiona '/' para enfocarte en el cuadro de búsqueda.",
  "highlightColor": "Color de resaltado",
  "highlightColorTooltip": "Ajusta la apariencia del texto resaltado.",
  "highlightBold": "Resaltado en negrita",
  "highlightBoldTooltip": "Hace que el texto resaltado sea en negrita para una mejor visibilidad.",
  "advancedSearch": "Búsqueda avanzada",
  "otherProjects": {
    "header": "Otros Proyectos",
    "askScreenshot": "Toma una captura de pantalla en cualquier página y ábrela automáticamente con ChatGPT. <i>También disponible para <a href=\"https://youtu.be/0g_EQbp5q4g\">Claude</a> y <a href=\"https://youtu.be/I-VX31SoGZ0\">Gemini</a>.</i>",
    "globalSpeed": "¡Establece una velocidad predeterminada para vídeo y audio!"
  }
}
</file>

<file path="static/locales/it.json">
{
  "general": "Generale",
  "help": "Aiuto",
  "strictSearch": "Ricerca rigorosa",
  "strictSearchTooltip": "Aumenta la soglia per i risultati corrispondenti.",
  "sortDate": "Ordina per data",
  "sortDateTooltip": "Ordina i risultati per data.",
  "scrollTop": "Scorri verso l'alto",
  "scrollTopTooltip": "Scorri verso l'alto quando selezioni una conversazione.",
  "issuePrompt": "Hai problemi o un suggerimento?",
  "issueDirective": "Fornisci feedback su Github.",
  "areYouSure": "Sei sicuro?",
  "showMore": "Mostra altro",
  "notFound": "Non trovato",
  "data": "Dati",
  "sessionOnly": "Caching solo per la sessione",
  "sessionOnlyTooltip": "Cancella la cronologia delle chat memorizzate nella cache alla chiusura del browser. Questo aumenta i tempi di ricerca e l'utilizzo dei dati, quindi usare solo se il dispositivo è condiviso.",
  "showImage": "Mostra icona",
  "showImageTooltip": "Se una conversazione ha utilizzato un GPT personalizzato, mostra l'icona se disponibile.",
  "searchChats": "Cerca chat...",
  "search": "Cerca",
  "language": "Lingua",
  "context": "Contesto",
  "contextTooltip": "Regola quanto contesto circostante viene mostrato con ogni risultato di ricerca.",
  "reset": "Reimposta",
  "clearCache": "Cancella cache",
  "autoClear": "Cancellazione automatica ricerca",
  "autoClearTooltip": "Cancella il campo di ricerca quando clicchi fuori.",
  "enableShortcut": "Abilita scorciatoia",
  "enableShortcutTooltip": "Premi '/' per focalizzarti sulla casella di ricerca.",
  "highlightColor": "Colore evidenziatore",
  "highlightColorTooltip": "Regola l'aspetto del testo evidenziato.",
  "highlightBold": "Evidenziatura in grassetto",
  "highlightBoldTooltip": "Rende il testo evidenziato in grassetto per una migliore visibilità.",
  "advancedSearch": "Ricerca avanzata",
  "otherProjects": {
    "header": "Altri Progetti",
    "askScreenshot": "Fai uno screenshot su qualsiasi pagina e aprilo automaticamente con ChatGPT. <i>Disponibile anche per <a href=\"https://youtu.be/0g_EQbp5q4g\">Claude</a> e <a href=\"https://youtu.be/I-VX31SoGZ0\">Gemini</a>.</i>",
    "globalSpeed": "Imposta una velocità predefinita per video e audio!"
  }
}
</file>

<file path="static/locales/ja.json">
{
  "_morpho": true,
  "general": "一般",
  "help": "ヘルプ",
  "strictSearch": "厳格な検索",
  "strictSearchTooltip": "一致する結果の閾値を上げる。",
  "sortDate": "日付で並べ替え",
  "sortDateTooltip": "結果を日付順に並べ替える。",
  "scrollTop": "トップへスクロール",
  "scrollTopTooltip": "会話を選択するとトップにスクロールします。",  
  "issuePrompt": "問題または提案がありますか？",
  "issueDirective": "Githubでフィードバックを提供してください。",
  "areYouSure": "よろしいですか？",
  "showMore": "もっと見る",
  "notFound": "見つかりません",
  "data": "データ",
  "sessionOnly": "セッションのみのキャッシング",
  "sessionOnlyTooltip": "ブラウザのクローズ時にキャッシュされたチャット履歴をクリアします。これにより検索時間とデータ使用量が増加するため、デバイスが共有されている場合のみ使用してください。",
  "showImage": "アイコンを表示",
  "showImageTooltip": "カスタムGPTを使用した会話の場合、利用可能であればアイコンを表示します。",
  "searchChats": "チャットを検索...",
  "search": "検索",
  "language": "言語",
  "context": "コンテキスト",
  "contextTooltip": "各検索結果に表示される周囲の情報の量を調整します。",
  "reset": "リセット",
  "clearCache": "キャッシュをクリア",
  "autoClear": "検索自動クリア",
  "autoClearTooltip": "離れるときに検索フィールドをクリアします。",
  "enableShortcut": "ショートカットを有効にする",
  "enableShortcutTooltip": "'/'を押して検索ボックスにフォーカスします。",
  "highlightColor": "ハイライト色",
  "highlightColorTooltip": "ハイライトされたテキストの外観を調整します。",
  "highlightBold": "太字ハイライト",
  "highlightBoldTooltip": "より良い可視性のために、ハイライトされたテキストを太字にします。",
  "advancedSearch": "高度な検索",
  "otherProjects": {
    "header": "その他のプロジェクト",
    "askScreenshot": "任意のページでスクリーンショットを取り、自動的にChatGPTで開きます。<i><a href=\"https://youtu.be/0g_EQbp5q4g\">Claude</a>や<a href=\"https://youtu.be/I-VX31SoGZ0\">Gemini</a>にも利用可能です。</i>",
    "globalSpeed": "ビデオとオーディオのデフォルト速度を設定！"
  }
}
</file>

<file path="static/locales/ko.json">
{
  "general": "일반",
  "help": "도움말",
  "strictSearch": "엄격한 검색",
  "strictSearchTooltip": "일치하는 결과의 임계값을 높입니다.",
  "sortDate": "날짜별 정렬",
  "sortDateTooltip": "결과를 날짜순으로 정렬합니다.",
  "scrollTop": "맨 위로 스크롤",
  "scrollTopTooltip": "대화를 선택할 때 맨 위로 스크롤합니다.",
  "issuePrompt": "문제나 제안이 있나요?",
  "issueDirective": "Github에서 피드백 제공하기.",
  "areYouSure": "확실합니까?",
  "showMore": "더 보기",
  "notFound": "찾을 수 없음",
  "data": "데이터",
  "sessionOnly": "세션-만 캐싱",
  "sessionOnlyTooltip": "브라우저 닫을 때 캐시된 채팅 기록을 지웁니다. 이는 검색 시간과 데이터 사용량을 증가시키므로, 기기가 공유되는 경우에만 사용하세요.",
  "showImage": "아이콘 표시",
  "showImageTooltip": "대화에서 사용자 지정 GPT를 사용한 경우, 가능하다면 아이콘을 표시합니다.",
  "searchChats": "채팅 검색...",
  "search": "검색",
  "language": "언어",
  "context": "컨텍스트",
  "contextTooltip": "각 검색 결과와 함께 표시되는 주변 정보의 양을 조정합니다.",
  "reset": "초기화",
  "clearCache": "캐시 지우기",
  "autoClear": "자동 지우기 검색",
  "autoClearTooltip": "다른 곳을 클릭할 때 검색 필드를 자동으로 지웁니다.",
  "enableShortcut": "단축키 활성화",
  "enableShortcutTooltip": "'/'를 눌러 검색 상자에 초점을 맞춥니다.",
  "highlightColor": "하이라이트 색상",
  "highlightColorTooltip": "하이라이트된 텍스트의 외관을 조정합니다.",
  "highlightBold": "볼드 하이라이트",
  "highlightBoldTooltip": "더 나은 가시성을 위해 하이라이트된 텍스트를 굵게 만듭니다.",
  "advancedSearch": "고급 검색",
  "otherProjects": {
    "header": "기타 프로젝트",
    "askScreenshot": "아무 페이지에서 스크린샷을 찍고 자동으로 ChatGPT로 엽니다. <i><a href=\"https://youtu.be/0g_EQbp5q4g\">Claude</a>와 <a href=\"https://youtu.be/I-VX31SoGZ0\">Gemini</a>에서도 사용 가능합니다.</i>",
    "globalSpeed": "비디오 및 오디오에 대한 기본 속도 설정!"
  }
}
</file>

<file path="static/locales/pt_BR.json">
{
  "general": "Geral",
  "help": "Ajuda",
  "strictSearch": "Pesquisa rigorosa",
  "strictSearchTooltip": "Aumentar o limiar para resultados correspondentes.",
  "sortDate": "Ordenar por data",
  "sortDateTooltip": "Ordenar os resultados por data.",
  "scrollTop": "Rolar para o topo",
  "scrollTopTooltip": "Role para o topo ao selecionar uma conversa.",
  "issuePrompt": "Tem problemas ou uma sugestão?",
  "issueDirective": "Forneça feedback no Github.",
  "areYouSure": "Você tem certeza?",
  "showMore": "Mostrar mais",
  "notFound": "Não encontrado",
  "data": "Dados",
  "sessionOnly": "Caching apenas para a sessão",
  "sessionOnlyTooltip": "Limpar o histórico de chat armazenado em cache ao fechar o navegador. Isso aumenta os tempos de busca e o uso de dados, então use apenas se o dispositivo for compartilhado.",
  "showImage": "Mostrar ícone",
  "showImageTooltip": "Se uma conversa usou um GPT personalizado, mostrar o ícone se disponível.",
  "searchChats": "Pesquisar chats...",
  "search": "Pesquisar",
  "language": "Idioma",
  "context": "Contexto",
  "contextTooltip": "Ajuste a quantidade de informações circundantes exibidas com cada resultado de pesquisa.",
  "reset": "Redefinir",
  "clearCache": "Limpar cache",
  "autoClear": "Limpeza automática de pesquisa",
  "autoClearTooltip": "Limpa o campo de pesquisa quando você clica fora.",
  "enableShortcut": "Habilitar atalho",
  "enableShortcutTooltip": "Pressione '/' para focar na caixa de pesquisa.",
  "highlightColor": "Cor de destaque",
  "highlightColorTooltip": "Ajuste a aparência do texto destacado.",
  "highlightBold": "Destaque em negrito",
  "highlightBoldTooltip": "Torna o texto destacado em negrito para melhor visibilidade.",
  "advancedSearch": "Pesquisa avançada",
  "otherProjects": {
    "header": "Outros Projetos",
    "askScreenshot": "Tire uma captura de tela em qualquer página e abra automaticamente com o ChatGPT. <i>Também disponível para <a href=\"https://youtu.be/0g_EQbp5q4g\">Claude</a> e <a href=\"https://youtu.be/I-VX31SoGZ0\">Gemini</a>.</i>",
    "globalSpeed": "Defina uma velocidade padrão para vídeo e áudio!"
  }
}
</file>

<file path="static/locales/ru.json">
{
  "general": "Общее",
  "help": "Помощь",
  "strictSearch": "Строгий поиск",
  "strictSearchTooltip": "Увеличить порог для соответствующих результатов.",
  "sortDate": "Сортировать по дате",
  "sortDateTooltip": "Сортировать результаты по дате.",
  "scrollTop": "Прокрутить вверх",
  "scrollTopTooltip": "Прокрутите вверх при выборе разговора.",
  "issuePrompt": "Есть проблемы или предложение?",
  "issueDirective": "Оставьте отзыв на Github.",
  "areYouSure": "Вы уверены?",
  "showMore": "Показать больше",
  "notFound": "Не найдено",
  "data": "Данные",
  "sessionOnly": "Кэширование только для сессии",
  "sessionOnlyTooltip": "Очистка кэшированной истории чата при закрытии браузера. Это увеличивает время поиска и использование данных, поэтому используйте только если ваше устройство используется совместно.",
  "showImage": "Показать иконку",
  "showImageTooltip": "Если в разговоре использовался настраиваемый GPT, показать иконку, если она доступна.",
  "searchChats": "Поиск чатов...",
  "search": "Поиск",
  "language": "Язык",
  "context": "Контекст",
  "contextTooltip": "Настройте, сколько окружающей информации отображается с каждым результатом поиска.",
  "reset": "Сброс",
  "clearCache": "Очистить кэш",
  "autoClear": "Автоочистка поиска",
  "autoClearTooltip": "Очищает поле поиска, когда вы кликаете вне его.",
  "enableShortcut": "Включить быструю клавишу",
  "enableShortcutTooltip": "Нажмите '/', чтобы сфокусироваться на поле поиска.",
  "highlightColor": "Цвет выделения",
  "highlightColorTooltip": "Настройте внешний вид выделенного текста.",
  "highlightBold": "Жирное выделение",
  "highlightBoldTooltip": "Делает выделенный текст жирным для лучшей видимости.",
  "advancedSearch": "Расширенный поиск",
  "otherProjects": {
    "header": "Другие Проекты",
    "askScreenshot": "Сделайте скриншот на любой странице и автоматически откройте его через ChatGPT. <i>Также доступно для <a href=\"https://youtu.be/0g_EQbp5q4g\">Claude</a> и <a href=\"https://youtu.be/I-VX31SoGZ0\">Gemini</a>.</i>",
    "globalSpeed": "Установите стандартную скорость для видео и аудио!"
  }
}
</file>

<file path="static/locales/tr.json">
{
  "general": "Genel",
  "help": "Yardım",
  "strictSearch": "Katı Arama",
  "strictSearchTooltip": "Eşleşen sonuçlar için eşiği artırın.",
  "sortDate": "Tarihe Göre Sırala",
  "sortDateTooltip": "Sonuçları tarihe göre sıralayın.",
  "scrollTop": "Yukarı kaydır",
  "scrollTopTooltip": "Bir konuşma seçerken yukarı kaydırın.",
  "issuePrompt": "Bir sorununuz veya öneriniz mi var?",
  "issueDirective": "Github üzerinden geri bildirimde bulunun.",
  "areYouSure": "Emin misiniz?",
  "showMore": "Daha fazla göster",
  "notFound": "Bulunamadı",
  "data": "Veri",
  "sessionOnly": "Yalnızca oturum için önbelleğe alma",
  "sessionOnlyTooltip": "Tarayıcı kapatıldığında önbelleğe alınmış sohbet geçmişini temizle. Bu, arama sürelerini ve veri kullanımını artırır, bu yüzden cihaz paylaşılıyorsa sadece kullanın.",
  "showImage": "Simgeyi göster",
  "showImageTooltip": "Bir konuşma özel bir GPT kullandıysa, simgeyi mevcutsa göster.",
  "searchChats": "Sohbetleri ara...",
  "search": "Ara",
  "language": "Dil",
  "context": "Bağlam",
  "contextTooltip": "Her arama sonucuyla birlikte gösterilen çevreleyen bilgi miktarını ayarlayın.",
  "reset": "Sıfırla",
  "clearCache": "Önbelleği temizle",
  "autoClear": "Otomatik temizleme araması",
  "autoClearTooltip": "Başka bir yere tıkladığınızda arama alanını temizler.",
  "enableShortcut": "Kısayolu etkinleştir",
  "enableShortcutTooltip": "'/' tuşuna basarak arama kutusuna odaklanın.",
  "highlightColor": "Vurgu rengi",
  "highlightColorTooltip": "Vurgulanan metnin görünümünü ayarlayın.",
  "highlightBold": "Kalın vurgu",
  "highlightBoldTooltip": "Daha iyi görünürlük için vurgulanan metni kalın yapar.",
  "advancedSearch": "Gelişmiş arama",
  "otherProjects": {
    "header": "Diğer Projeler",
    "askScreenshot": "Herhangi bir sayfada ekran görüntüsü alın ve otomatik olarak ChatGPT ile açın. <i><a href=\"https://youtu.be/0g_EQbp5q4g\">Claude</a> ve <a href=\"https://youtu.be/I-VX31SoGZ0\">Gemini</a> için de mevcuttur.</i>",
    "globalSpeed": "Video ve ses için varsayılan hızı ayarlayın!"
  }
}
</file>

<file path="static/locales/vi.json">
{
  "_morpho": true,
  "general": "Chung",
  "help": "Trợ giúp",
  "strictSearch": "Tìm kiếm chặt chẽ",
  "strictSearchTooltip": "Tăng ngưỡng cho kết quả phù hợp.",
  "sortDate": "Sắp xếp theo ngày",
  "sortDateTooltip": "Sắp xếp kết quả theo ngày.",
  "scrollTop": "Cuộn lên trên",
  "scrollTopTooltip": "Cuộn lên trên khi chọn một cuộc trò chuyện.",
  "issuePrompt": "Có vấn đề hoặc gợi ý?",
  "issueDirective": "Cung cấp phản hồi trên Github.",
  "areYouSure": "Bạn có chắc không?",
  "showMore": "Xem thêm",
  "notFound": "Không tìm thấy",
  "data": "Dữ liệu",
  "sessionOnly": "Bộ nhớ đệm chỉ trong phiên",
  "sessionOnlyTooltip": "Xóa lịch sử trò chuyện được lưu trong bộ nhớ đệm khi đóng trình duyệt. Điều này làm tăng thời gian tìm kiếm và sử dụng dữ liệu, vì vậy chỉ sử dụng nếu thiết bị của bạn được chia sẻ.",
  "showImage": "Hiển thị biểu tượng",
  "showImageTooltip": "Nếu một cuộc trò chuyện sử dụng GPT tùy chỉnh, hiển thị biểu tượng nếu có.",
  "searchChats": "Tìm kiếm cuộc trò chuyện...",
  "search": "Tìm kiếm",
  "language": "Ngôn ngữ",
  "context": "Bối cảnh",
  "contextTooltip": "Điều chỉnh lượng thông tin xung quanh hiển thị với mỗi kết quả tìm kiếm.",
  "reset": "Đặt lại",
  "clearCache": "Xóa bộ nhớ cache",
  "autoClear": "Tự động xóa tìm kiếm",
  "autoClearTooltip": "Xóa trường tìm kiếm khi bạn nhấp ra ngoài.",
  "enableShortcut": "Kích hoạt phím tắt",
  "enableShortcutTooltip": "Nhấn '/' để tập trung vào hộp tìm kiếm.",
  "highlightColor": "Màu nổi bật",
  "highlightColorTooltip": "Điều chỉnh vẻ ngoài của văn bản được làm nổi bật.",
  "highlightBold": "Nổi bật đậm",
  "highlightBoldTooltip": "Làm cho văn bản nổi bật được in đậm để dễ nhìn hơn.",
  "advancedSearch": "Tìm kiếm nâng cao",
  "otherProjects": {
    "header": "Các Dự Án Khác",
    "askScreenshot": "Chụp ảnh màn hình trên bất kỳ trang nào và tự động mở nó với ChatGPT. <i>Cũng có sẵn cho <a href=\"https://youtu.be/0g_EQbp5q4g\">Claude</a> và <a href=\"https://youtu.be/I-VX31SoGZ0\">Gemini</a>.</i>",
    "globalSpeed": "Thiết lập tốc độ mặc định cho video và âm thanh!"
  }
}
</file>

<file path="static/locales/zh_CN.json">
{
  "_morpho": true,
  "general": "通用",
  "help": "帮助",
  "strictSearch": "严格搜索",
  "strictSearchTooltip": "提高匹配结果的阈值。",
  "sortDate": "按日期排序",
  "sortDateTooltip": "按日期对结果进行排序。",
  "scrollTop": "滚动到顶部",
  "scrollTopTooltip": "选择对话时滚动到顶部。",
  "issuePrompt": "有问题或建议吗？",
  "issueDirective": "在Github上提供反馈。",
  "areYouSure": "您确定吗？",
  "showMore": "显示更多",
  "notFound": "未找到",
  "data": "数据",
  "sessionOnly": "仅限会话缓存",
  "sessionOnlyTooltip": "关闭浏览器时清除缓存的聊天历史记录。这将增加搜索时间和数据使用量，因此仅在您的设备被共享时使用。",
  "showImage": "显示图标",
  "showImageTooltip": "如果对话使用了自定义GPT，可用时显示图标。",
  "searchChats": "搜索聊天...",
  "search": "搜索",
  "language": "语言",
  "context": "上下文",
  "contextTooltip": "调整每个搜索结果显示的周围信息量。",
  "reset": "重置",
  "clearCache": "清除缓存",
  "autoClear": "自动清除搜索",
  "autoClearTooltip": "当您点击其他地方时清除搜索框。",
  "enableShortcut": "启用快捷键",
  "enableShortcutTooltip": "按'/'聚焦搜索框。",
  "highlightColor": "高亮颜色",
  "highlightColorTooltip": "调整高亮文本的外观。",
  "highlightBold": "加粗高亮",
  "highlightBoldTooltip": "使高亮文本加粗以便更好的可见性。",
  "advancedSearch": "高级搜索",
  "otherProjects": {
    "header": "其他项目",
    "askScreenshot": "在任何页面上截图，并自动用ChatGPT打开。<i>也适用于<a href=\"https://youtu.be/0g_EQbp5q4g\">Claude</a>和<a href=\"https://youtu.be/I-VX31SoGZ0\">Gemini</a>。</i>",
    "globalSpeed": "为视频和音频设置默认速度！"
  }
}
</file>

<file path="static/locales/zh_TW.json">
{
  "_morpho": true,
  "general": "一般",
  "help": "幫助",
  "strictSearch": "嚴格搜尋",
  "strictSearchTooltip": "提高匹配結果的閾值。",
  "sortDate": "按日期排序",
  "sortDateTooltip": "按日期對結果進行排序。",
  "scrollTop": "滾動到頂部",
  "scrollTopTooltip": "選擇對話時滾動到頂部。",
  "issuePrompt": "有問題或建議嗎？",
  "issueDirective": "在Github上提供反饋。",
  "areYouSure": "您確定嗎？",
  "showMore": "顯示更多",
  "notFound": "未找到",
  "data": "資料",
  "sessionOnly": "僅限階段快取",
  "sessionOnlyTooltip": "關閉瀏覽器時清除快取的聊天歷史。這會增加搜尋時間和資料使用量，所以只有在您的裝置被共享時才使用。",
  "showImage": "顯示圖標",
  "showImageTooltip": "如果對話中使用了自訂GPT，則顯示圖標（如果可用）。",
  "searchChats": "搜尋聊天...",
  "search": "搜尋",
  "language": "語言",
  "context": "上下文",
  "contextTooltip": "調整每個搜尋結果顯示的周圍信息量。",
  "reset": "重設",
  "clearCache": "清除快取",
  "autoClear": "自動清除搜尋",
  "autoClearTooltip": "當您點擊其他地方時清除搜尋框。",
  "enableShortcut": "啟用快捷鍵",
  "enableShortcutTooltip": "按'/'聚焦搜尋框。",
  "highlightColor": "高亮顏色",
  "highlightColorTooltip": "調整高亮文本的外觀。",
  "highlightBold": "粗體高亮",
  "highlightBoldTooltip": "使高亮文本粗體以便更好的可見性。",
  "advancedSearch": "進階搜尋",
  "otherProjects": {
    "header": "其他專案",
    "askScreenshot": "在任何頁面上截圖，並自動用ChatGPT開啟。<i>也適用於<a href=\"https://youtu.be/0g_EQbp5q4g\">Claude</a>和<a href=\"https://youtu.be/I-VX31SoGZ0\">Gemini</a>。</i>",
    "globalSpeed": "為視頻和音頻設置默認速度！"
  }
}
</file>

<file path="static/options.html">
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<body>
    <div id="root"></div>
    <script src="options.js"></script>
</body>
</html>
</file>

<file path="staticCh/manifest.json">
{
  "name": "__MSG_appName__",
  "short_name": "GPT Search",
  "version": "0.0.985",
  "description": "__MSG_appDesc__",
  "manifest_version": 3,
  "default_locale": "en",
  "host_permissions": ["https://*.chatgpt.com/*"],
  "permissions": ["storage", "unlimitedStorage", "webRequest", "scripting"],
  "icons": { 
    "128": "128.png" 
  },
  "background": {
    "service_worker": "background.js",
    "type": "module"
  },
  "action": {},
  "options_ui": {
    "open_in_tab": true,
    "page": "options.html"
  },
  "content_scripts": [
    {
      "matches": ["https://chatgpt.com/*"],
      "js": ["preamble.js"],
      "run_at": "document_start"
    },
    {
      "matches": ["https://chatgpt.com/*"],
      "js": ["main.js"],
      "world": "MAIN",
      "run_at": "document_start"
    }
  ]
}
</file>

<file path="staticFf/manifest.json">
{
  "name": "__MSG_appName__",
  "short_name": "GPT Search",
  "version": "0.0.985",
  "description": "__MSG_appDesc__",
  "manifest_version": 3,
  "browser_specific_settings": {
    "gecko": {
      "id": "{1e7006c8-e1f9-4d0d-a451-93d8ce23b365}",
      "strict_min_version": "113.0"
    }
  },
  "web_accessible_resources": [
    {"resources": ["main.js"], "matches": ["https://chatgpt.com/*"]}
  ],
  "default_locale": "en",
  "host_permissions": ["https://*.chatgpt.com/*"],
  "permissions": ["storage", "unlimitedStorage", "webRequest", "scripting"],
  "icons": { 
    "128": "128.png" 
  },
  "background": {
    "scripts": ["background.js"]
  },
  "content_scripts": [
    {
      "matches": ["https://chatgpt.com/*"],
      "js": ["preamble.js", "mainLoader.js"],
      "run_at": "document_start"
    }
  ]
}
</file>

<file path="tools/generateGsmType.js">
// /// <reference types="@types/node" />

const { access, constants, writeFile, readFile } = require("fs").promises
const { join } = require("path")

const EN_PATH = join("static", "locales", "en.json")
const GSM_PATH = join("src", "utils", "GsmType.ts")


let newData = ""
async function main() {
    if (!await pathExists(EN_PATH)) return console.error("en.json does not exist")
    const data = JSON.parse( await readFile(EN_PATH, {encoding: "utf8"}))
    walk(data)
    writeFile(GSM_PATH, newData, {encoding: "utf8"})
}

function walk(d, level = 0) {
    if (level === 0) newData =  "\nexport type Gsm = {"
    const e = Object.entries(d)
    for (let i = 0; i < e.length; i++) {
        let postfix = (i === e.length - 1) ? "" : ","
        let l = level + 1 
        let p = "\n".concat(" ".repeat(l * 2))
        let isOptional = e[i][0].startsWith("_")

        const type = typeof e[i][1]
        if (type !== "object") {
            newData = newData.concat(p, e[i][0], isOptional ? "?" : "",  `: ${type}`, postfix)
        } else if (Array.isArray(e[i][1])) {
            newData = newData.concat(p, e[i][0], ": {")
            walk(e[i][1][0], l)
            newData = newData.concat(p, "}[]", postfix)
        } else {
            newData = newData.concat(p, e[i][0], ": {")
            walk(e[i][1], l)
            newData = newData.concat(p, "}", postfix)
        }
    }
    if (level === 0) newData = newData.concat("\n}")
}

async function pathExists(path) {
    try {
        await access(path, constants.W_OK)
        return true 
    } catch (err) {
        return false 
    }
}

main()
</file>

<file path="tools/validateLocale.js">
// /// <reference types="@types/node" />

// Test to make sure all locales have the required strings.

const { readFileSync } = require("fs")
const { exit } = require("process")

const locales = ["en", "it", "es", "ja", "ko", "pt_BR", "ru", "tr", "zh_CN", "zh_TW"]

let targetLeaves;

for (let locale of locales) {
  let leaves; 
  try {
    leaves = getLeafs(JSON.parse(readFileSync(`./static/locales/${locale}.json`, {encoding: "utf8"})))
  } catch (err) {
    console.log("Could not parse", locale, leaves)
    exit()
  }

  if (!targetLeaves) {
    targetLeaves = leaves;
    continue 
  }


  const omitted = new Set(targetLeaves.filter(v => !leaves.includes(v)))
  const extra = new Set(leaves.filter(v => !targetLeaves.includes(v)))

  omitted.forEach(o => o.startsWith("_") && omitted.delete(o))
  extra.forEach(o => o.startsWith("_") && extra.delete(o))

  if (omitted.size) {
    console.log("OMITTED", "\n=========")
    omitted.forEach(v => console.log(v))
  }

  if (extra.size) {
    console.log("\nEXTRA", "\n=========")
    extra.forEach(v => console.log(v))
  }

  if (omitted.size + extra.size ) {
    console.log("\nFIX", locale)
    exit()
  }
}

console.log("ALL GOOD!")

function getLeafs(obj, ctx = []) {
  const leafs = []
  for (let [k, v] of Object.entries(obj)) {
    if (typeof v === "object") {
      leafs.push(...getLeafs(v, [...ctx, k]))
    } else {
      leafs.push([...ctx, k].join('.'))
    }
  }
  return leafs
}
</file>

<file path="tsconfig.json">
{
  "compilerOptions": {
    "target": "es2022",
    "useDefineForClassFields": true,
    "module": "ESNext",
    "lib": ["DOM", "DOM.iterable", "ESNext"],
    "types": ["@types/chrome", "@types/react", "@types/react-dom"],
    "moduleResolution": "bundler",
    "allowImportingTsExtensions": true,
    "resolveJsonModule": true,
    "isolatedModules": true,
    "noEmit": true,

    "noImplicitAny": true,
    "allowJs": true,
    "sourceMap": false,
    "jsx": "react-jsx"
  },
  "exclude": ["./node_modules", "./build"]
}
</file>

<file path="webpack.config.js">
const { resolve } = require("path")
const { env } = require("process")
const webpack = require('webpack')
const TerserPlugin = require('terser-webpack-plugin')

const tsx = {
  test: /\.tsx?$/,
  exclude: /node_modules/,
  resourceQuery: { not: [/sfx/] },
  use: "babel-loader"
}

const entry = {
  raccoon: "./src/raccoon/index.tsx",
  main: "./src/main.ts",
  preamble: "./src/preamble/index.ts",
  background: "./src/background/index.ts",
  options: "./src/options/index.tsx"
}

if (env.FIREFOX) {
  entry["mainLoader"] = "./src/mainLoader.ts"
}

const common = {
  target: "browserslist",
  entry,
  output: {
    path: resolve(__dirname, env.FIREFOX ? "buildFf": "build", "unpacked")
  },
  module: {
    rules: [
      tsx,
      {...tsx, resourceQuery: /sfx/, sideEffects: true},
      {
        sideEffects: true,
        test: /\.css$/,
        exclude: /node_modules/,
        resourceQuery: { not: [/raw/] },
        use: [
            "style-loader", 
            {
              loader: "css-loader",
              options: {
                import: true,
              }
            },
            "postcss-loader"
        ],
      },
      {
        test: /\.css$/,
        resourceQuery: /raw/,
        exclude: [/node_modules/],
        type: 'asset/source',
        use: [
          "postcss-loader"
        ]
      }
    ]
  },
  resolve: {
    extensions: [".tsx", '.ts', '.js']
  },
  plugins: [
    new webpack.ProvidePlugin({
      gvar: [resolve(__dirname, "src", "globalVar.ts"), "gvar"]
    })
  ]
}

if (env.NODE_ENV === "production") {
  module.exports = {
    ...common,
    mode: "production",
    optimization: {
      minimize: true,
      minimizer: [
        new TerserPlugin({
          terserOptions: {
            format: {
              comments: false,
            }
          },
          extractComments: false,
        })
      ]
    }
  }
} else {
  module.exports = {
    ...common,
    mode: "development",
    devtool: false
  }
}
</file>

<file path=".gitignore">
# Logs
logs
*.log
npm-debug.log*
yarn-debug.log*
yarn-error.log*
pnpm-debug.log*
lerna-debug.log*
build/
buildFf
.DS_STORE 

node_modules
build
dist-ssr
*.local

# Editor directories and files
.vscode/*
!.vscode/extensions.json
.idea
.DS_Store
*.suo
*.ntvs*
*.njsproj
*.sln
*.sw?
</file>

<file path="README.md">
# GPT Search: Chat History

## Install on [Chrome](https://chromewebstore.google.com/detail/gpt-search/glhkbfoibolghhfikadjikgfmaknpelb), [Firefox](https://addons.mozilla.org/en-US/firefox/addon/gpt-search), or [Edge](https://microsoftedge.microsoft.com/addons/detail/gpt-search/hcnfioacjbamffbgigbjpdlflnlpaole). 

## Main features
1. Search through your conversation history. 
2. Beautifully integrated into ChatGPT's UI. 
3. Blazingly fast after initial caching. 
4. [Advanced search](./advancedSearch.md)

![screenshot1](https://github.com/user-attachments/assets/60350b14-d7b8-4f9e-8a3d-9e6816387c0c)
</file>

</files>
