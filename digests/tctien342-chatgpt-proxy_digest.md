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
.gitignore
docker-compose.yml
Dockerfile
env.ts
index.ts
lab.ts
LICENSE
package.json
README.md
src/agent.ts
src/fetch.ts
src/openai.ts
src/tools.ts
src/types.d.ts
tsconfig.json
</directory_structure>

<files>
This section contains the contents of the repository's files.

<file path=".env.example">
AGENT_ROLL_INTERVAL=60000
MAX_SESSION_TRIES=3
API_TOKEN=
</file>

<file path="docker-compose.yml">
version: "3"
services:
  openai-proxy:
    build: .
    container_name: openai-proxy
    restart: always
    ports:
      - "3000:3000"
    environment:
      - API_TOKEN=
      - AGENT_ROLL_INTERVAL=60000
</file>

<file path="Dockerfile">
# use the official Bun image
# see all versions at https://hub.docker.com/r/oven/bun/tags
FROM oven/bun:slim as base
WORKDIR /usr/src/app

# install dependencies into temp directory
# this will cache them and speed up future builds
FROM base AS install
RUN mkdir -p /temp/dev
COPY package.json bun.lockb /temp/dev/
RUN cd /temp/dev && bun install --frozen-lockfile

# install with --production (exclude devDependencies)
RUN mkdir -p /temp/prod
COPY package.json bun.lockb /temp/prod/
RUN cd /temp/prod && bun install --frozen-lockfile --production

# copy production dependencies and source code into final image
FROM base AS release
COPY --from=install /temp/prod/node_modules node_modules
COPY . .

# run the app
USER bun
ENTRYPOINT [ "bun", "run", "index.ts" ]
</file>

<file path="env.ts">
import z from "zod";
import { randomUUID } from "node:crypto";

const ENVSchema = z.object({
  BASE_URL: z.string().default("https://chat.openai.com"),
  APP_PORT: z
    .string()
    .transform((v) => Number(v))
    .default("3000"),
  MAX_SESSION_TRIES: z
    .string()
    .transform((v) => Number(v))
    .default("3"),
  API_TOKEN: z
    .string()
    .default("")
    .transform((val) => {
      if (!val) {
        return "sk-" + randomUUID().replaceAll("-", "");
      }
      return val;
    }),
  /**
   * Auto generate new token and fetch agent for openAI
   * @default 1 minute
   */
  AGENT_ROLL_INTERVAL: z
    .string()
    .transform((v) => Number(v))
    .default("60000"),
});

export const ENV = ENVSchema.parse(process.env);
</file>

<file path="index.ts">
import cors from "@elysiajs/cors";
import { Elysia } from "elysia";
import Stream from "@elysiajs/stream";
import { handleChatCompletion } from "./src/openai";
import bearer from "@elysiajs/bearer";
import { ENV } from "./env";
import { AppLogger } from "./src/tools";
import { AgentManager } from "./src/agent";

AgentManager.getInstance();

const app = new Elysia()
  .use(
    cors({
      origin: "*",
      allowedHeaders: ["*"],
      methods: ["GET", "POST", "OPTIONS"],
    })
  )
  .use(bearer())
  .onBeforeHandle(({ bearer, set, request }) => {
    if (bearer !== ENV.API_TOKEN) {
      AppLogger.w(
        "handle",
        `Incomming ${request.method.toUpperCase()} request have been denied`,
        {
          url: request.url,
          reason: "Unauthorized",
        }
      );
      set.status = 401;
      return {
        status: 401,
        body: "Unauthorized",
      };
    }
    AppLogger.i("handle", `Incomming ${request.method.toUpperCase()} request`, {
      url: request.url,
    });
  })
  .options("*", () => "OK")
  .post("/v1/chat/completions", async ({ body }) => {
    const req = body as any;
    const isStream = req.stream ?? false;
    if (!isStream) {
      return handleChatCompletion({
        messages: req.messages,
      });
    }
    return new Stream(async (stream) => {
      return handleChatCompletion({
        messages: req.messages,
        streamRes: stream,
      });
    });
  })
  .all("*", () => "Not Found")
  .listen(3000);

AppLogger.i("START", "Server running on port", ENV.APP_PORT);
AppLogger.i("START", "Server running with token", ENV.API_TOKEN);
AppLogger.i("START", "Server running with endpoints", ["/v1/chat/completions"]);
</file>

<file path="lab.ts">
import { getNewSessionId } from "./src/openai";

const device = await getNewSessionId();
</file>

<file path="LICENSE">
MIT License

Copyright (c) 2024 Saintno

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

<file path="package.json">
{
  "name": "openai-proxy",
  "version": "1.0.50",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1",
    "dev": "bun run --watch index.ts",
    "start": "bun run index.ts"
  },
  "dependencies": {
    "@elysiajs/bearer": "^1.0.2",
    "@elysiajs/cors": "^1.0.2",
    "@elysiajs/stream": "^1.0.2",
    "@saintno/needed-tools": "^0.3.4",
    "elysia": "latest",
    "gpt-3-encoder": "^1.1.4",
    "js-sha3": "^0.9.3",
    "user-agents": "^1.1.177",
    "zod": "^3.22.4"
  },
  "devDependencies": {
    "@types/user-agents": "^1.0.4",
    "bun-types": "latest"
  },
  "module": "src/index.js"
}
</file>

<file path="src/agent.ts">
import UserAgent from "user-agents";
import { getNewSession } from "./openai";
import { AppLogger, randomIntTargetOffset } from "./tools";
import { ENV } from "../env";
import { sleep } from "bun";

export class AgentManager {
  static instance: AgentManager;

  private userAgent: UserAgent;
  private session?: Session;

  private constructor() {
    this.userAgent = new UserAgent();
  }

  /**
   * Rolls the agent by generating a new session ID and updating the token and device ID.
   */
  async roll(tries = 0): Promise<Session> {
    AppLogger.i("AgentManager", "Trigger rolling agent");
    this.session = undefined;
    if (tries === ENV.MAX_SESSION_TRIES) {
      throw new Error(
        `Failed to get session ID after ${ENV.MAX_SESSION_TRIES} retries`
      );
    }
    return await getNewSession()
      .then((session) => {
        this.session = session;
        return session;
      })
      .catch(async (e) => {
        AppLogger.w(
          "AgentManager",
          "Failed to get session ID, retry after 3s",
          e
        );
        await sleep(randomIntTargetOffset(3000, 500));
        return await this.roll(tries + 1);
      });
  }

  get userAgentString() {
    return this.userAgent.toString();
  }

  get openAiHeaders(): HeadersInit {
    return {
      "oai-device-id": this.session?.deviceId ?? "",
      "openai-sentinel-chat-requirements-token": this.session?.token ?? "",
    };
  }

  static getInstance(): AgentManager {
    if (!AgentManager.instance) {
      AgentManager.instance = new AgentManager();
    }
    return AgentManager.instance;
  }

  get crrSession() {
    return this.session;
  }
}
</file>

<file path="src/fetch.ts">
import { ENV } from "../env";

import https from "https";
import { APIQueueItem, CustomFetch, QueueManager } from "@saintno/needed-tools";
import { AgentManager } from "./agent";

const cFetch = new CustomFetch();
const fAgent = new https.Agent({
  rejectUnauthorized: false,
});

const Fetcher = APIQueueItem.createInstance(
  cFetch,
  new QueueManager("Fetcher")
);

/**
 * Generates the headers object for a fetch request.
 *
 * @param headers - Optional headers to be included in the generated headers object.
 * @returns The generated headers object.
 */
const generateFetchHeaders = (headers?: HeadersInit) => {
  return {
    ...headers,
    accept: "*/*",
    "accept-language": "en-US,en;q=0.9",
    "cache-control": "no-cache",
    "content-type": "application/json",
    "oai-language": "en-US",
    origin: ENV.BASE_URL,
    pragma: "no-cache",
    referer: ENV.BASE_URL,
    "sec-ch-ua":
      '"Google Chrome";v="123", "Not:A-Brand";v="8", "Chromium";v="123"',
    "sec-ch-ua-mobile": "?0",
    "sec-ch-ua-platform": '"Windows"',
    "sec-fetch-dest": "empty",
    "sec-fetch-mode": "cors",
    "sec-fetch-site": "same-origin",
    "user-agent": AgentManager.getInstance().userAgentString,
  };
};

cFetch.setBeforeCall(async (url, config) => {
  config = {
    ...config,
    headers: generateFetchHeaders(config.headers),
    agent: fAgent,
  } as FetchRequestInit;
  return { url, config };
});

cFetch.setOnParse(async (response) => {
  return response;
});

export { Fetcher, generateFetchHeaders };
</file>

<file path="src/openai.ts">
/**
 * This file included ported funtion from @PawanOsman
 */
import { Fetcher } from "./fetch";
import { ENV } from "../env";

import { randomUUID, randomInt, createHash } from "node:crypto";
import { encode } from "gpt-3-encoder";
import Stream from "@elysiajs/stream";
import { AgentManager } from "./agent";
import { AppLogger } from "./tools";
import { sha3_512 } from "js-sha3";

const apiUrl = `${ENV.BASE_URL}/backend-anon/conversation`;

/**
 * Generates a unique completion ID with the given prefix.
 * @param prefix The prefix to use for the completion ID. Defaults to "cmpl-".
 * @returns The generated completion ID.
 */
function generateCompletionId(prefix: string = "cmpl-") {
  const characters =
    "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
  const length = 28;
  for (let i = 0; i < length; i++) {
    prefix += characters.charAt(Math.floor(Math.random() * characters.length));
  }
  return prefix;
}

/**
 * Converts chunks of data into lines.
 * @param chunksAsync An async iterable that yields chunks of data.
 * @returns An async generator that yields lines of data.
 */
async function* chunksToLines(chunksAsync: any) {
  let previous = "";
  for await (const chunk of chunksAsync) {
    const bufferChunk = Buffer.isBuffer(chunk) ? chunk : Buffer.from(chunk);
    previous += bufferChunk;
    let eolIndex: number;
    while ((eolIndex = previous.indexOf("\n")) >= 0) {
      // line includes the EOL
      const line = previous.slice(0, eolIndex + 1).trimEnd();
      if (line === "data: [DONE]") break;
      if (line.startsWith("data: ")) yield line;
      previous = previous.slice(eolIndex + 1);
    }
  }
}

/**
 * Converts lines from an asynchronous iterator to messages.
 * @param linesAsync - An asynchronous iterator that yields lines.
 * @returns An asynchronous generator that yields messages.
 */
async function* linesToMessages(linesAsync: any) {
  for await (const line of linesAsync) {
    const message = line.substring("data :".length);
    yield message;
  }
}

/**
 * Asynchronously generates a stream of completions from the provided data.
 *
 * @param data - The data to process for completions.
 * @returns An asynchronous generator that yields completions.
 */
async function* streamCompletion(data: any) {
  yield* linesToMessages(chunksToLines(data));
}

/**
 * Generates a new session ID and token for the OpenAI API.
 * @returns An object containing the new device ID and token.
 */
async function getNewSession(): Promise<Session> {
  let newDeviceId = randomUUID();
  const session = await new Fetcher(
    `${ENV.BASE_URL}/backend-anon/sentinel/chat-requirements`
  )
    .json()
    .post<Session>(
      {},
      {
        headers: { "oai-device-id": newDeviceId },
      }
    );
  AppLogger.i(
    "getNewSession",
    `System: Successfully refreshed session ID and token. ${
      !session.token ? "(Now it's ready to process requests)" : ""
    }`
  );
  session.deviceId = newDeviceId;
  return session;
}

/**
 * Generates a proof token for the OpenAI API.
 */
function GenerateProofToken(
  seed: string,
  diff: string,
  userAgent: string
): string {
  const cores: number[] = [8, 12, 16, 24];
  const screens: number[] = [3000, 4000, 6000];

  const core = cores[randomInt(0, cores.length)];
  const screen = screens[randomInt(0, screens.length)];

  const now = new Date(Date.now() - 8 * 3600 * 1000);
  const parseTime = now.toUTCString().replace("GMT", "GMT-0500 (Eastern Time)");

  const config = [core + screen, parseTime, 4294705152, 0, userAgent];

  const diffLen = diff.length / 2;

  for (let i = 0; i < 100000; i++) {
    config[3] = i;
    const jsonData = JSON.stringify(config);
    const base = Buffer.from(jsonData).toString("base64");
    const hashValue = sha3_512.create().update(seed + base);

    if (hashValue.hex().substring(0, diffLen) <= diff) {
      const result = "gAAAAAB" + base;
      return result;
    }
  }

  const fallbackBase = Buffer.from(`"${seed}"`).toString("base64");
  return "gAAAAABwQ8Lk5FbGpA2NcR9dShT6gYjU7VxZ4D" + fallbackBase;
}

/**
 * Handles chat completion by sending messages to the OpenAI API and processing the response.
 */
async function handleChatCompletion({
  messages,
  streamRes,
}: {
  /**
   * The messages to process for completion.
   */
  messages: { role: string; content: string }[];
  /**
   * The stream response to send the completion to.
   */
  streamRes?: Stream<string | number | boolean | object>;
}) {
  const session = await AgentManager.getInstance()
    .roll()
    .catch((e) => {
      AppLogger.w("handleChatCompletion", "Failed to get a new session", e);
      return null;
    });

  if (!session) {
    const resp = {
      status: false,
      error: {
        message: `Error getting a new session, please try again later, if the issue persists, please open an issue on the GitHub repository, https://github.com/PawanOsman/ChatGPT`,
        type: "invalid_request_error",
      },
      support: "https://discord.pawan.krd",
    };
    if (streamRes) {
      streamRes.send(JSON.stringify(resp));
      streamRes.close();
    } else {
      return resp;
    }
    return;
  }

  let proofToken = GenerateProofToken(
    session.proofofwork.seed,
    session.proofofwork.difficulty,
    AgentManager.getInstance().userAgentString
  );

  let promptTokens = 0;
  let completionTokens = 0;
  const body = {
    action: "next",
    messages: messages.map((message: { role: string; content: string }) => ({
      author: { role: message.role },
      content: { content_type: "text", parts: [message.content] },
    })),
    parent_message_id: randomUUID(),
    model: "text-davinci-002-render-sha",
    timezone_offset_min: -180,
    suggestions: [],
    history_and_training_disabled: true,
    conversation_mode: { kind: "primary_assistant" },
    websocket_request_id: randomUUID(),
  };
  for (let message of messages) {
    promptTokens += encode(message.content).length;
  }
  try {
    let fullContent = "";
    let finish_reason = null;
    let created = Math.floor(Date.now() / 1000); // Unix timestamp in seconds
    let requestId = generateCompletionId("chatcmpl-");

    const response = await new Fetcher(apiUrl).post<Response>(body, {
      headers: {
        ...AgentManager.getInstance().openAiHeaders,
        "openai-sentinel-proof-token": proofToken,
      },
    });
    if (!response.ok || !response.body) {
      throw new Error("An error occurred while processing the request.");
    }
    for await (const message of streamCompletion(response.body as any)) {
      if (message.match(/^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}.\d{6}$/)) {
        // Skip heartbeat detection
        continue;
      }

      const parsed = JSON.parse(message);
      let content = parsed?.message?.content?.parts[0] ?? "";
      let status = parsed?.message?.status ?? "";
      for (let message of messages) {
        if (message.content === content) {
          content = "";
          break;
        }
      }

      switch (status) {
        case "in_progress":
          finish_reason = null;
          break;
        case "finished_successfully":
          let finish_reason_data =
            parsed?.message?.metadata?.finish_details?.type ?? null;
          switch (finish_reason_data) {
            case "max_tokens":
              finish_reason = "length";
              break;
            case "stop":
            default:
              finish_reason = "stop";
          }
          break;
        default:
          finish_reason = null;
      }

      if (content === "") continue;
      let completionChunk = content.replace(fullContent, "");
      completionTokens += encode(completionChunk).length;
      if (streamRes) {
        let response = {
          id: requestId,
          created: created,
          object: "chat.completion.chunk",
          model: "gpt-3.5-turbo",
          choices: [
            {
              delta: {
                content: completionChunk,
              },
              index: 0,
              finish_reason: finish_reason,
            },
          ],
        };
        streamRes.send(JSON.stringify(response));
      }
      fullContent = content.length > fullContent.length ? content : fullContent;
    }
    if (streamRes) {
      streamRes.send(
        JSON.stringify({
          id: requestId,
          created: created,
          object: "chat.completion.chunk",
          model: "gpt-3.5-turbo",
          choices: [
            {
              delta: {
                content: "",
              },
              index: 0,
              finish_reason: finish_reason,
            },
          ],
        })
      );
      streamRes.close();
    } else {
      return {
        id: requestId,
        created: created,
        model: "gpt-3.5-turbo",
        object: "chat.completion",
        choices: [
          {
            finish_reason: finish_reason,
            index: 0,
            message: {
              content: fullContent,
              role: "assistant",
            },
          },
        ],
        usage: {
          prompt_tokens: promptTokens,
          completion_tokens: completionTokens,
          total_tokens: promptTokens + completionTokens,
        },
      };
    }
  } catch (error) {
    const content = {
      status: false,
      error: {
        message:
          "An error occurred. Please check the server console to confirm it is ready and free of errors. Additionally, ensure that your request complies with OpenAI's policy.",
        type: "invalid_request_error",
      },
      support: "https://discord.pawan.krd",
    };
    if (streamRes) {
      streamRes.send(JSON.stringify(content));
      streamRes.close();
    } else {
      return content;
    }
  }
}

export {
  generateCompletionId,
  chunksToLines,
  linesToMessages,
  streamCompletion,
  getNewSession,
  handleChatCompletion,
};
</file>

<file path="src/tools.ts">
import { Logger } from "@saintno/needed-tools";

export const AppLogger = new Logger("AppLogger");

export const handleProcessExit = (callback: () => void) => {
  process.on("SIGINT", callback);
  process.on("SIGTERM", callback);
};

export const sleep = (ms: number) =>
  new Promise((resolve) => setTimeout(resolve, ms));

export const randomIntRange = (min: number, max: number) => {
  return Math.floor(Math.random() * (max - min + 1)) + min;
};

export const randomIntTargetOffset = (target: number, offset: number) => {
  return target + randomIntRange(-offset, offset);
};
</file>

<file path="src/types.d.ts">
type Session = {
  deviceId: string;
  persona: string;
  arkose: {
    required: boolean;
    dx: any;
  };
  turnstile: {
    required: boolean;
  };
  proofofwork: {
    required: boolean;
    seed: string;
    difficulty: string;
  };
  token: string;
};
</file>

<file path="tsconfig.json">
{
  "compilerOptions": {
    /* Visit https://aka.ms/tsconfig to read more about this file */

    /* Projects */
    // "incremental": true,                              /* Save .tsbuildinfo files to allow for incremental compilation of projects. */
    // "composite": true,                                /* Enable constraints that allow a TypeScript project to be used with project references. */
    // "tsBuildInfoFile": "./.tsbuildinfo",              /* Specify the path to .tsbuildinfo incremental compilation file. */
    // "disableSourceOfProjectReferenceRedirect": true,  /* Disable preferring source files instead of declaration files when referencing composite projects. */
    // "disableSolutionSearching": true,                 /* Opt a project out of multi-project reference checking when editing. */
    // "disableReferencedProjectLoad": true,             /* Reduce the number of projects loaded automatically by TypeScript. */

    /* Language and Environment */
    "target": "ES2021",                                  /* Set the JavaScript language version for emitted JavaScript and include compatible library declarations. */
    // "lib": [],                                        /* Specify a set of bundled library declaration files that describe the target runtime environment. */
    // "jsx": "preserve",                                /* Specify what JSX code is generated. */
    // "experimentalDecorators": true,                   /* Enable experimental support for TC39 stage 2 draft decorators. */
    // "emitDecoratorMetadata": true,                    /* Emit design-type metadata for decorated declarations in source files. */
    // "jsxFactory": "",                                 /* Specify the JSX factory function used when targeting React JSX emit, e.g. 'React.createElement' or 'h'. */
    // "jsxFragmentFactory": "",                         /* Specify the JSX Fragment reference used for fragments when targeting React JSX emit e.g. 'React.Fragment' or 'Fragment'. */
    // "jsxImportSource": "",                            /* Specify module specifier used to import the JSX factory functions when using 'jsx: react-jsx*'. */
    // "reactNamespace": "",                             /* Specify the object invoked for 'createElement'. This only applies when targeting 'react' JSX emit. */
    // "noLib": true,                                    /* Disable including any library files, including the default lib.d.ts. */
    // "useDefineForClassFields": true,                  /* Emit ECMAScript-standard-compliant class fields. */
    // "moduleDetection": "auto",                        /* Control what method is used to detect module-format JS files. */

    /* Modules */
    "module": "ES2022",                                /* Specify what module code is generated. */
    // "rootDir": "./",                                  /* Specify the root folder within your source files. */
    "moduleResolution": "node",                       /* Specify how TypeScript looks up a file from a given module specifier. */
    // "baseUrl": "./",                                  /* Specify the base directory to resolve non-relative module names. */
    // "paths": {},                                      /* Specify a set of entries that re-map imports to additional lookup locations. */
    // "rootDirs": [],                                   /* Allow multiple folders to be treated as one when resolving modules. */
    // "typeRoots": [],                                  /* Specify multiple folders that act like './node_modules/@types'. */
    "types": ["bun-types"],                                      /* Specify type package names to be included without being referenced in a source file. */
    // "allowUmdGlobalAccess": true,                     /* Allow accessing UMD globals from modules. */
    // "moduleSuffixes": [],                             /* List of file name suffixes to search when resolving a module. */
    // "resolveJsonModule": true,                        /* Enable importing .json files. */
    // "noResolve": true,                                /* Disallow 'import's, 'require's or '<reference>'s from expanding the number of files TypeScript should add to a project. */

    /* JavaScript Support */
    // "allowJs": true,                                  /* Allow JavaScript files to be a part of your program. Use the 'checkJS' option to get errors from these files. */
    // "checkJs": true,                                  /* Enable error reporting in type-checked JavaScript files. */
    // "maxNodeModuleJsDepth": 1,                        /* Specify the maximum folder depth used for checking JavaScript files from 'node_modules'. Only applicable with 'allowJs'. */

    /* Emit */
    // "declaration": true,                              /* Generate .d.ts files from TypeScript and JavaScript files in your project. */
    // "declarationMap": true,                           /* Create sourcemaps for d.ts files. */
    // "emitDeclarationOnly": true,                      /* Only output d.ts files and not JavaScript files. */
    // "sourceMap": true,                                /* Create source map files for emitted JavaScript files. */
    // "outFile": "./",                                  /* Specify a file that bundles all outputs into one JavaScript file. If 'declaration' is true, also designates a file that bundles all .d.ts output. */
    // "outDir": "./",                                   /* Specify an output folder for all emitted files. */
    // "removeComments": true,                           /* Disable emitting comments. */
    // "noEmit": true,                                   /* Disable emitting files from a compilation. */
    // "importHelpers": true,                            /* Allow importing helper functions from tslib once per project, instead of including them per-file. */
    // "importsNotUsedAsValues": "remove",               /* Specify emit/checking behavior for imports that are only used for types. */
    // "downlevelIteration": true,                       /* Emit more compliant, but verbose and less performant JavaScript for iteration. */
    // "sourceRoot": "",                                 /* Specify the root path for debuggers to find the reference source code. */
    // "mapRoot": "",                                    /* Specify the location where debugger should locate map files instead of generated locations. */
    // "inlineSourceMap": true,                          /* Include sourcemap files inside the emitted JavaScript. */
    // "inlineSources": true,                            /* Include source code in the sourcemaps inside the emitted JavaScript. */
    // "emitBOM": true,                                  /* Emit a UTF-8 Byte Order Mark (BOM) in the beginning of output files. */
    // "newLine": "crlf",                                /* Set the newline character for emitting files. */
    // "stripInternal": true,                            /* Disable emitting declarations that have '@internal' in their JSDoc comments. */
    // "noEmitHelpers": true,                            /* Disable generating custom helper functions like '__extends' in compiled output. */
    // "noEmitOnError": true,                            /* Disable emitting files if any type checking errors are reported. */
    // "preserveConstEnums": true,                       /* Disable erasing 'const enum' declarations in generated code. */
    // "declarationDir": "./",                           /* Specify the output directory for generated declaration files. */
    // "preserveValueImports": true,                     /* Preserve unused imported values in the JavaScript output that would otherwise be removed. */

    /* Interop Constraints */
    // "isolatedModules": true,                          /* Ensure that each file can be safely transpiled without relying on other imports. */
    // "allowSyntheticDefaultImports": true,             /* Allow 'import x from y' when a module doesn't have a default export. */
    "esModuleInterop": true,                             /* Emit additional JavaScript to ease support for importing CommonJS modules. This enables 'allowSyntheticDefaultImports' for type compatibility. */
    // "preserveSymlinks": true,                         /* Disable resolving symlinks to their realpath. This correlates to the same flag in node. */
    "forceConsistentCasingInFileNames": true,            /* Ensure that casing is correct in imports. */

    /* Type Checking */
    "strict": true,                                      /* Enable all strict type-checking options. */
    // "noImplicitAny": true,                            /* Enable error reporting for expressions and declarations with an implied 'any' type. */
    // "strictNullChecks": true,                         /* When type checking, take into account 'null' and 'undefined'. */
    // "strictFunctionTypes": true,                      /* When assigning functions, check to ensure parameters and the return values are subtype-compatible. */
    // "strictBindCallApply": true,                      /* Check that the arguments for 'bind', 'call', and 'apply' methods match the original function. */
    // "strictPropertyInitialization": true,             /* Check for class properties that are declared but not set in the constructor. */
    // "noImplicitThis": true,                           /* Enable error reporting when 'this' is given the type 'any'. */
    // "useUnknownInCatchVariables": true,               /* Default catch clause variables as 'unknown' instead of 'any'. */
    // "alwaysStrict": true,                             /* Ensure 'use strict' is always emitted. */
    // "noUnusedLocals": true,                           /* Enable error reporting when local variables aren't read. */
    // "noUnusedParameters": true,                       /* Raise an error when a function parameter isn't read. */
    // "exactOptionalPropertyTypes": true,               /* Interpret optional property types as written, rather than adding 'undefined'. */
    // "noImplicitReturns": true,                        /* Enable error reporting for codepaths that do not explicitly return in a function. */
    // "noFallthroughCasesInSwitch": true,               /* Enable error reporting for fallthrough cases in switch statements. */
    // "noUncheckedIndexedAccess": true,                 /* Add 'undefined' to a type when accessed using an index. */
    // "noImplicitOverride": true,                       /* Ensure overriding members in derived classes are marked with an override modifier. */
    // "noPropertyAccessFromIndexSignature": true,       /* Enforces using indexed accessors for keys declared using an indexed type. */
    // "allowUnusedLabels": true,                        /* Disable error reporting for unused labels. */
    // "allowUnreachableCode": true,                     /* Disable error reporting for unreachable code. */

    /* Completeness */
    // "skipDefaultLibCheck": true,                      /* Skip type checking .d.ts files that are included with TypeScript. */
    "skipLibCheck": true                                 /* Skip type checking all .d.ts files. */
  }
}
</file>

<file path=".gitignore">
# See https://help.github.com/articles/ignoring-files/ for more about ignoring files.

# dependencies
/node_modules
/.pnp
.pnp.js

# testing
/coverage

# next.js
/.next/
/out/

# production
/build

# misc
.DS_Store
*.pem

# debug
npm-debug.log*
yarn-debug.log*
yarn-error.log*

# local env files
.env.local
.env.development.local
.env.test.local
.env.production.local

# vercel
.vercel

**/*.trace
**/*.zip
**/*.tar.gz
**/*.tgz
**/*.log
package-lock.json
**/*.bun
.cache
</file>

<file path="README.md">
# OpenAI Proxy

This project serves as a proxy for the OpenAI API. Using `bun` as runtime.

## Getting Started

To get started with this project, you'll need to install the necessary dependencies. You can do this by running:

```sh
bun install
```

## Development

To start the development server, run:

```sh
bun dev
```

This will start the server and watch for changes in the [`src/index.ts`]("src/index.ts") file.

## Docker Compose

This project includes a `docker-compose.yml` file for running the application in a Docker container. This can be useful for development and testing, as well as for deploying the application in a production environment.

To use Docker Compose, you'll first need to install Docker and Docker Compose on your machine. Once you've done that, you can start the application with the following command:

```sh
docker-compose up
```

This will build the Docker image for the application (if it hasn't been built already) and start a new container. The application will be accessible at `http://localhost:3000`.

## Environment Variables

This project uses environment variables for configuration. These are stored in the [`.env`](".env") file. You'll need to provide your own `API_TOKEN` in this file.

In the Docker environment, these variables are set in the `docker-compose.yml` file:

```yml
version: "3"
services:
  openai-proxy:
    build: .
    container_name: openai-proxy
    restart: always
    ports:
      - "3000:3000"
    environment:
      - API_TOKEN=sk-1234567890abcdef
      - AGENT_ROLL_INTERVAL=60000
```

Please replace `sk-1234567890abcdef` with your middleware API token. Default will be a random token each time container started if `API_TOKEN` empty.

## Code Structure

The main entry point for the application is [`src/index.ts`]("src/index.ts"). This project also includes several utility functions and classes:

- [`src/openai.ts`]("src/openai.ts"): Contains various utility functions for interacting with the OpenAI API.
- [`src/fetch.ts`]("src/fetch.ts"): Defines the `Fetcher` class for making HTTP requests.
- [`src/agent.ts`]("src/agent.ts"): Defines the `AgentManager` class for managing user agents.
- [`src/tools.ts`]("src/tools.ts"): Contains the `AppLogger` class for logging.

## Testing

Currently, this project does not have any tests defined.

## Contributing

Contributions are welcome! Please feel free to submit a pull request.

## License

This project is licensed under the terms of the MIT license.

## Acknowledgements

Special thanks to @PawanOsman for his [wonderful project](https://github.com/PawanOsman/ChatGPT), which served as inspiration for this project.
</file>

</files>
