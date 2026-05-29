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
.github/workflows/publish.yml
.gitignore
.npmignore
package.json
README.md
src/index.ts
src/utils/utils.ts
tests/answer.test.ts
tests/rotate.test.ts
tests/stream.test.ts
tsconfig.json
tsup.config.ts
vitest.config.ts
</directory_structure>

<files>
This section contains the contents of the repository's files.

<file path=".github/workflows/publish.yml">
name: Build & Publish to npm

on:
  push:
    tags:
      - "v*"

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout repo
        uses: actions/checkout@v4

      - name: Set up Node
        uses: actions/setup-node@v4
        with:
          node-version: 24
          registry-url: "https://registry.npmjs.org/"

      - name: Install dependencies
        run: npm install

      - name: Run build
        run: npm run build

      - name: Run tests
        run: npm run test

      - name: Publish to npm
        run: npm publish --access public
        env:
          NODE_AUTH_TOKEN: ${{secrets.NPM_TOKEN}}
</file>

<file path=".npmignore">
/src
/node_modules

.env

tsconfig.json
tsup.config.ts
vitest.config.ts
/tests
/.github
</file>

<file path="package.json">
{
  "name": "chatgptreversed",
  "version": "0.2.5",
  "main": "./dist/index.js",
  "module": "./dist/index.mjs",
  "types": "./dist/index.d.ts",
  "scripts": {
    "build": "tsup --minify",
    "publish": "npm run build && npm publish",
    "test": "vitest --disable-console-intercept"
  },
  "author": "gin337",
  "license": "MIT",
  "homepage": "https://github.com/gin337/ChatGPTReversed",
  "repository": {
    "url": "https://github.com/gin337/ChatGPTReversed"
  },
  "description": "Free ChatGPT API reversed and simplified",
  "keywords": [
    "chatgpt",
    "chatgpt api",
    "chatgpt proxy",
    "openai"
  ],
  "devDependencies": {
    "@types/node": "^20.14.2",
    "tsup": "^8.1.0",
    "typescript": "^5.4.5",
    "vitest": "^3.0.4"
  }
}
</file>

<file path="src/index.ts">
import {randomUUID} from "crypto";
import {generateFakeSentinelToken, simulateBypassHeaders, solveSentinelChallenge} from "./utils/utils";

interface CompleteOptionsProfile {
  stream?: boolean;
}

export class ChatGPTReversed {
  public static csrfToken: string | undefined = undefined;
  private static initialized: boolean = false;

  constructor() {
    if (ChatGPTReversed.initialized) throw new Error("ChatGPTReversed has already been initialized.");

    this.initialize();
  }

  private async initialize(): Promise<void> {
    ChatGPTReversed.initialized = true;
  }

  public async rotateSessionData(): Promise<{
    uuid: string;
    csrf: string;
    sentinel: {
      token: string;
      proof: string;
      oaiSc: string;
    };
  }> {
    const uuid = randomUUID();
    const csrfToken = await this.getCSRFToken(uuid);
    const sentinelToken = await this.getSentinelToken(uuid, csrfToken);

    ChatGPTReversed.csrfToken = csrfToken;

    return {
      uuid,
      csrf: csrfToken,
      sentinel: sentinelToken,
    };
  }

  private async getCSRFToken(uuid: string): Promise<string> {
    if (ChatGPTReversed.csrfToken !== undefined) {
      return ChatGPTReversed.csrfToken;
    }

    const headers = await simulateBypassHeaders({
      spoofAddress: true,
      preOaiUUID: uuid,
      accept: "application/json",
    });

    const response = await fetch("https://chatgpt.com/api/auth/csrf", {
      method: "GET",
      headers: headers,
    });

    const data = await response.json();

    if (data.csrfToken === undefined) {
      throw new Error("Failed to fetch required CSRF token");
    }

    return data.csrfToken;
  }

  private async getSentinelToken(
    uuid: string,
    csrf: string
  ): Promise<{
    token: string;
    proof: string;
    oaiSc: string;
  }> {
    const headers = await simulateBypassHeaders({
      spoofAddress: true,
      preOaiUUID: uuid,
      accept: "application/json",
    });

    const test = await generateFakeSentinelToken();

    const response = await fetch("https://chatgpt.com/backend-anon/sentinel/chat-requirements", {
      body: JSON.stringify({
        p: test,
      }),
      headers: {
        ...headers,
        Cookie: `__Host-next-auth.csrf-token=${csrf}; oai-did=${uuid}; oai-nav-state=1;`,
      },
      method: "POST",
    });

    const data = await response.json();

    if (data.token === undefined || data.proofofwork === undefined) {
      throw new Error("Failed to fetch required required sentinel token");
    }

    const oaiSc = response.headers.get("set-cookie")?.split("oai-sc=")[1]?.split(";")[0] || "";

    if (!oaiSc) {
      throw new Error("Failed to fetch required oai-sc token");
    }

    const challengeToken = await solveSentinelChallenge(data.proofofwork.seed, data.proofofwork.difficulty);

    return {
      token: data.token,
      proof: challengeToken,
      oaiSc: oaiSc,
    };
  }

  public async complete(
    message: string,
    options: {stream: true}
  ): Promise<AsyncGenerator<{text: string; metadata: any}>>;

  public async complete(message: string, options?: {stream?: false}): Promise<string>;

  public async complete(
    message: string,
    options?: CompleteOptionsProfile
  ): Promise<string | AsyncGenerator<{text: string; metadata: any}>> {
    const sessionData = await this.rotateSessionData();

    if (!ChatGPTReversed.initialized) {
      throw new Error(
        "ChatGPTReversed has not been initialized. Please initialize the instance before calling this method."
      );
    }

    const headers = await simulateBypassHeaders({
      accept: "text/event-stream",
      spoofAddress: true,
      preOaiUUID: sessionData.uuid,
    });

    const messageID = randomUUID();

    const response = await fetch("https://chatgpt.com/backend-anon/conversation", {
      headers: {
        ...headers,
        Cookie: `__Host-next-auth.csrf-token=${sessionData.csrf}; oai-did=${sessionData.uuid}; oai-nav-state=1; oai-sc=${sessionData.sentinel.oaiSc};`,
        "openai-sentinel-chat-requirements-token": sessionData.sentinel.token,
        "openai-sentinel-proof-token": sessionData.sentinel.proof,
      },
      body: JSON.stringify({
        action: "next",
        messages: [
          {
            id: messageID,
            author: {
              role: "user",
            },
            create_time: Date.now(),
            content: {
              content_type: "text",
              parts: [message],
            },
            metadata: {
              selected_all_github_repos: false,
              selected_github_repos: [],
              serialization_metadata: {
                custom_symbol_offsets: [],
              },
              dictation: false,
            },
          },
        ],
        paragen_cot_summary_display_override: "allow",
        parent_message_id: "client-created-root",
        model: "auto",
        timezone_offset_min: -60,
        timezone: "Europe/Berlin",
        suggestions: [],
        history_and_training_disabled: true,
        conversation_mode: {
          kind: "primary_assistant",
        },
        system_hints: [],
        supports_buffering: true,
        supported_encodings: ["v1"],
        client_contextual_info: {
          is_dark_mode: true,
          time_since_loaded: 7,
          page_height: 911,
          page_width: 1080,
          pixel_ratio: 1,
          screen_height: 1080,
          screen_width: 1920,
          app_name: "chatgpt.com",
        },
      }),
      method: "POST",
    });

    if (!response.ok) {
      throw new Error(`Request failed with status ${response.status}: ${response.statusText}`);
    }

    if (response.body === null) {
      throw new Error("Failed to receive response body. Please check your sessionToken and try again.");
    }

    if (options?.stream) {
      return this.streamResponse(response);
    }

    return this.collectFullResponse(response);
  }

  private async collectFullResponse(response: Response): Promise<string> {
    const reader = response.body!.getReader();
    const decoder = new TextDecoder();

    let result = "";
    let buffer = "";
    let finished = false;

    while (true) {
      const {done, value} = await reader.read();
      if (done) break;

      buffer += decoder.decode(value, {stream: true});
      let lines = buffer.split("\n");
      buffer = lines.pop() || "";

      for (const line of lines) {
        if (!line.startsWith("data:")) continue;
        const dataStr = line.replace("data:", "").trim();
        if (!dataStr || dataStr === "[DONE]") continue;

        try {
          const json = JSON.parse(dataStr);

          if (json.message) {
            if (json.message.content && json.message.content.parts) {
              result = json.message.content.parts[0];
            }
            if (json.message.status === "finished_successfully") {
              finished = true;
              break;
            }
          } else if (json.o === "append" && json.p === "/message/content/parts/0") {
            result += json.v;
          } else if (Array.isArray(json.v)) {
            for (const op of json.v) {
              if (op.o === "append" && op.p === "/message/content/parts/0") {
                result += op.v;
              }
              if (op.p === "/message/status" && op.o === "replace" && op.v === "finished_successfully") {
                finished = true;
              }
            }
          }
        } catch {
          continue;
        }
      }

      if (finished) break;
    }

    if (!finished && buffer.startsWith("data:")) {
      try {
        const json = JSON.parse(buffer.replace("data:", "").trim());
        if (json.message && json.message.content && json.message.content.parts) {
          result = json.message.content.parts[0];
        }
      } catch {
      }
    }

    return result;
  }

  private async *streamResponse(response: Response): AsyncGenerator<{text: string; metadata: any}> {
    const reader = response.body!.getReader();
    const decoder = new TextDecoder();

    let buffer = "";
    let fullText = "";
    let finished = false;

    while (!finished) {
      const {done, value} = await reader.read();
      if (done) break;

      buffer += decoder.decode(value, {stream: true});
      let lines = buffer.split("\n");
      buffer = lines.pop() || "";

      for (const line of lines) {
        if (!line.startsWith("data:")) continue;

        const dataStr = line.slice("data:".length).trim();
        if (!dataStr || dataStr === "[DONE]") continue;

        let json: any;
        try {
          json = JSON.parse(dataStr);
        } catch {
          continue;
        }

        let deltaText = "";
        let metadata: any = undefined;

        if (json.message) {
          const parts = json.message.content?.parts;
          metadata = json.message.metadata ?? json.metadata;

          if (Array.isArray(parts) && typeof parts[0] === "string") {
            const current = parts[0];
            if (current.startsWith(fullText)) {
              deltaText = current.slice(fullText.length);
            } else {
              deltaText = current;
            }
            fullText = current;
          }

          if (json.message.status === "finished_successfully") {
            finished = true;
          }
        }

        if (json.o === "append" && json.p === "/message/content/parts/0") {
          deltaText += json.v;
          fullText += json.v;
        }

        if (Array.isArray(json.v)) {
          for (const op of json.v) {
            if (op.o === "append" && op.p === "/message/content/parts/0") {
              deltaText += op.v;
              fullText += op.v;
            }
            if (op.p === "/message/status" && op.o === "replace" && op.v === "finished_successfully") {
              finished = true;
            }
          }
        }

        if (json.type === "message_stream_complete") {
          finished = true;
        }

        if (deltaText) {
          yield {
            text: deltaText,
            metadata: metadata ?? json.metadata ?? {},
          };
        }

        if (finished) break;
      }
    }
  }
}
</file>

<file path="src/utils/utils.ts">
import {randomUUID, randomInt, createHash} from "crypto";

export const randomIP = async (): Promise<string> =>
  Array.from({length: 4}, () => Math.floor(Math.random() * 256)).join(".");

export const _randomUUID = (): string => randomUUID().toString();

const simulated = {
  agent:
    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36",
  platform: "Windows",
  mobile: "?0",
  ua: 'Not A(Brand";v="8", "Chromium";v="132", "Google Chrome";v="132',
};

export async function simulateBypassHeaders({
  accept,
  spoofAddress = false,
  preOaiUUID,
}: {
  accept: string;
  spoofAddress?: boolean;
  preOaiUUID?: string;
}): Promise<Record<string, string>> {
  const ip = await randomIP();
  const uuid = _randomUUID();

  return {
    accept: accept,
    "Content-Type": "application/json",
    "cache-control": "no-cache",
    Referer: "https://chatgpt.com/",
    "Referrer-Policy": "strict-origin-when-cross-origin",
    "oai-device-id": preOaiUUID || uuid,
    "oai-language": "en",
    "User-Agent": simulated.agent,
    pragma: "no-cache",
    priority: "u=1, i",
    "sec-ch-ua": `"${simulated.ua}"`,
    "sec-ch-ua-mobile": simulated.mobile,
    "sec-ch-ua-platform": `"${simulated.platform}"`,
    "sec-fetch-site": "same-origin",
    "sec-fetch-mode": "cors",
    ...(spoofAddress
      ? {
          "X-Forwarded-For": ip,
          "X-Originating-IP": ip,
          "X-Remote-IP": ip,
          "X-Remote-Addr": ip,
          "X-Host": ip,
          "X-Forwarded-Host": ip,
          Forwarded: `for=${ip}`,
          "True-Client-IP": ip,
          "X-Real-IP": ip,
        }
      : {}),
  };
}

export async function solveSentinelChallenge(seed: string, difficulty: string): Promise<string> {
  const cores = [8, 12, 16, 24];
  const screens = [3000, 4000, 6000];

  const core = cores[randomInt(0, cores.length)];
  const screen = screens[randomInt(0, screens.length)];

  const now = new Date(Date.now() - 8 * 3600 * 1000);
  const parseTime = now.toUTCString().replace("GMT", "GMT+0100 (Central European Time)");

  const config = [core + screen, parseTime, 4294705152, 0, simulated.agent];

  const diffLen = difficulty.length / 2;

  for (let i = 0; i < 100000; i++) {
    config[3] = i;
    const jsonData = JSON.stringify(config);
    const base = Buffer.from(jsonData).toString("base64");
    const hashValue = createHash("sha3-512")
      .update(seed + base)
      .digest();

    if (hashValue.toString("hex").substring(0, diffLen) <= difficulty) {
      const result = "gAAAAAB" + base;
      return result;
    }
  }

  const fallbackBase = Buffer.from(`"${seed}"`).toString("base64");
  return "gAAAAABwQ8Lk5FbGpA2NcR9dShT6gYjU7VxZ4D" + fallbackBase;
}

export async function generateFakeSentinelToken() {
  const prefix = "gAAAAAC";

  const config = [
    randomInt(3000, 6000),
    new Date().toUTCString().replace("GMT", "GMT+0100 (Central European Time)"),
    4294705152,
    0,
    simulated.agent,
    "de",
    "de",
    401,
    "mediaSession",
    "location",
    "scrollX",
    randomFloat(1000, 5000),
    crypto.randomUUID(),
    "",
    12,
    Date.now(),
  ];

  const base64 = Buffer.from(JSON.stringify(config)).toString("base64");

  return prefix + base64;
}

function randomFloat(min: number, max: number) {
  return (Math.random() * (max - min) + min).toFixed(4);
}
</file>

<file path="tests/answer.test.ts">
import {expect, test} from "vitest";
import {ChatGPTReversed} from "../src/index";

test("Retrieves answer from oai", async () => {
  const chatgpt = new ChatGPTReversed();

  const result = await chatgpt.complete("Hey, how are you?");
  console.log("Result: ", result);
  expect(typeof result).toBe("string");
});
</file>

<file path="tests/rotate.test.ts">
import {expect, test} from "vitest";
import {ChatGPTReversed} from "../src/index";

test("Rotates all needed values", async () => {
  const chatgpt = new ChatGPTReversed();

  const result = await chatgpt.rotateSessionData();

  console.log("Result: ", result);

  expect(result).toEqual({
    uuid: expect.any(String),
    csrf: expect.any(String),
    sentinel: {
      token: expect.any(String),
      proof: expect.any(String),
      oaiSc: expect.any(String),
    },
  });
});
</file>

<file path="tests/stream.test.ts">
import {expect, test} from "vitest";
import {ChatGPTReversed} from "../src/index";

test("Retrieves stream answer from oai", async () => {
  const chatgpt = new ChatGPTReversed();

  const result = await chatgpt.complete("Hey, how are you?", {stream: true});

  let streamData = "";
  for await (const chunk of result) {
    streamData += chunk.text;
    console.log("Chunk: ", chunk.text);
  }
  console.log("Streamed Result: ", streamData);

  expect(streamData).toBeTypeOf("string");
});
</file>

<file path="tsconfig.json">
{
    "compilerOptions": {
        "strict": true,
        "noImplicitAny": true,
        "esModuleInterop": true,
        "strictNullChecks": true,
        "target": "ES2022",
        "moduleResolution": "Node10",
        "module": "CommonJS",
        "declaration": true,
        "isolatedModules": true,
        "noEmit": true,
        "outDir": "dist"
    },
    "include": [
        "src"
    ],
    "exclude": [
        "node_modules"
    ]
}
</file>

<file path="tsup.config.ts">
import {defineConfig} from "tsup";

export default defineConfig({
  format: ["cjs", "esm"],
  entry: ["./src/index.ts"],
  dts: true,
  shims: true,
  skipNodeModulesBundle: true,
  clean: true,
});
</file>

<file path="vitest.config.ts">
import {defineConfig} from "vitest/config";

export default defineConfig({
  test: {
    globals: true,
    environment: "node",
  },
});
</file>

<file path=".gitignore">
/node_modules
/dist

.package-lock.json


.env
</file>

<file path="README.md">
# ChatGPTReversed - Educational project

## **Update 02.02.2025**: The `ChatGPTReversed` instance doesn't need any params anymore. Scroll to the End for documentation

## **Update 30.12.2025**: In recent weeks, OpenAI changed their frontend flow. I will update the project in the next months once I have time

Lets keep it simple, this is a educational project to learn to reverse complex API's and understand how they communicate with the frontend.
In this case we take a look at the ChatGPT frontend and reverse engineer the API used to communicate with the LLM.

OpenAi uses several techniques to prevent malicious use of their API, eg. rate limiting, token expiration, hashing, proof of work, continious calls, proxies, captchas, etc.

We take a look at the ChatGPT webapp as starting point and only use the chromium devtools to understand the process.

![step1](https://i.imgur.com/FbvbKML.png)
We start by opening the ChatGPT webapp and open the devtools to see the network requests.

![step2](https://i.imgur.com/SSXA50s.png)
`https://chatgpt.com/backend-api/conversation` endpoint is called with a POST request, we can see the payload and type of response. (In this case its a EventStream)

![step3](https://i.imgur.com/52qYDXC.png)
![step4](https://i.imgur.com/p3eRbQ8.png)
Several Identifing headers and cookies are used to prevent abuse, in this case:
`Authorization(JWT Token)`, `csrf-token(CSRF protection)`, `session-token (Same as the JWT token)`, `Requirements-Token & Proof Token`

![step5](https://i.imgur.com/CwCzpnV.png)
`https://chatgpt.com/backend-api/sentinel/chat-requirements` endpoint is called before the conversation starts, it passes in the token x and returns the token y.

```json
{
  "persona": "chatgpt-freeaccount",
  "token": "y",
  "arkose": {},
  "turnstile": {},
  "proofofwork": {
    "required": true,
    "seed": "0.81186133b2821174",
    "difficulty": "073682"
  }
}
```

To find out how x is retrieved we need to take a look at the minified source code of the frontend.

![step6](https://i.imgur.com/XuWosqk.png)
Token x in this case is variable e which is passed as callback from variable n which uses the function `getRequirementsToken` to retrieve the token.
![step7](https://i.imgur.com/hJfvHKS.png)
The function `getRequirementsToken` in this case returns the token x by checking if the value is already in a map called `answers`, if not it calls the function \_generateAnswer which returns the token x by using a hash function provided by the hashing library `hash-wasm`.

![step8](https://i.imgur.com/Ld0al4b.png)
We place a breakpoint right after the `getRequirementsToken` function is called and check the returned value which is the token x.

So we have the token x (Requirements token), we need to pass to the endpoint, we also have the token y (Required Requirements Token) which is returned by the endpoint. The last thing we need is the token z (Proof token) which as we find out is also generated with `_generateAnswer`.

`_generateAnswer` function is called with the seed and difficulty returned by the endpoint, it uses the seed and also multiple parameters retrieved by the `getConfig` such as screen size, timezone, cpu cores, etc. to generate a hash and satisfy the difficulty condition. If no hash is found it will increment the step and try again. It falls back to a specified value after multiple steps.

In this case the function `_generateAnswer` is called with the seed and difficulty returned by the endpoint and that returns the token z.

So we have all the required tokens to call the conversation endpoint and start a conversation with the LLM.
To recap:

- Session Token (JWT Token) is returned by `https://chatgpt.com/api/auth/session` in field `accessToken`

- CSRF Token is returned by `https://chatgpt.com/api/auth/csrf` in field `csrfToken`

- Requirements Token (Token x) is returned by `getRequirementsToken` function

- Required Requirements Token (Token y) is returned by the endpoint `https://chatgpt.com/backend-api/sentinel/chat-requirements`

- Proof Token (Token z) is returned by `_generateAnswer` function with the seed and difficulty returned by the `https://chatgpt.com/backend-api/sentinel/chat-requirements` endpoint

The rest is basic web communication knowledge.

## Documentation

```typescript
import {ChatGPTReversed} from "chatgptreversed"; // const {ChatGPTReversed} = require("chatgptreversed");

const chatgpt = new ChatGPTReversed();

const result = await chatgpt.complete("Hello, how are you?");
console.log(result);

// Output: Hello! I'm here and ready to assist you. How can I help you today?
```

```typescript
import {ChatGPTReversed} from "chatgptreversed"; // const {ChatGPTReversed} = require("chatgptreversed");

const chatgpt = new ChatGPTReversed();

async function main() {
  const result = await chatgpt.complete("Hello, how are you?");
  console.log(result);
}

main();

// Output: Hello! I'm here and ready to assist you. How can I help you today?
```

Like this project? Leave a star! 💫⭐
</file>

</files>
