# ChatGPT Internal Backend API (`backend-api`) Documentation

This is a comprehensive list of all the undocumented, internal `backend-api` endpoints used by the ChatGPT web client, reverse-engineered from the `acheong08`, `gin337`, and `Siamsnus` repositories. 

These endpoints are located at `https://chatgpt.com/backend-api/...` and generally require an `Authorization: Bearer <token>` header, and in some cases, proof-of-work tokens.

## 1. Chat & Conversation Management
The core endpoints for retrieving and interacting with chat histories.

*   `GET /conversations?offset={int}&limit={int}&order=updated`
    *   **Purpose:** Fetches the standard list of your chat history.
*   `GET /conversations?is_archived=true&offset={int}&limit={int}`
    *   **Purpose:** Fetches archived conversations.
*   `GET /conversation/{conversation_id}`
    *   **Purpose:** Fetches the full JSON tree (messages, nodes, metadata) for a specific conversation.
*   `POST /conversation`
    *   **Purpose:** Sends a prompt to the LLM and receives a Server-Sent Event (SSE) stream of the response. Requires complex Proof Tokens (Token x, y, z) mapped by `gin337/ChatGPTReversed`.
*   `POST /conversation/gen_title/{conversation_id}`
    *   **Purpose:** Triggers the LLM to auto-generate a title for the conversation.
*   `POST /conversations/batch`
    *   **Purpose:** Perform batch operations on multiple conversations (e.g., bulk delete or archive).

## 2. Projects & Custom GPTs (Gizmos)
ChatGPT treats "Projects" and "Custom GPTs" under the internal umbrella term **Gizmos**.

*   `GET /projects`
    *   **Purpose:** Lists all the Team/Enterprise Workspaces and Projects the user belongs to.
*   `GET /gizmos/discovery/mine`
    *   **Purpose:** Lists all custom GPTs created or pinned by the user.
*   `GET /gizmos/{project_or_gizmo_id}`
    *   **Purpose:** Gets the specific configuration, instructions, and metadata of a Project or Custom GPT.
*   `GET /gizmos/{project_or_gizmo_id}/conversations?cursor={string}`
    *   **Purpose:** **The Golden Endpoint.** Fetches the chat history *specifically* isolated within a particular Project. Unlike standard chats, it uses a `cursor` instead of an `offset` for pagination.

## 3. Memory & Personalization
Endpoints related to ChatGPT's persistent memory feature.

*   `GET /memories?include_memory_entries=true`
    *   **Purpose:** Retrieves all the facts and details ChatGPT has autonomously memorized about the user across all chats.
*   `GET /user_system_messages`
    *   **Purpose:** Retrieves the user's "Custom Instructions" (What would you like ChatGPT to know about you / How would you like it to respond).

## 4. Shared Links
Endpoints for interacting with publicly shared ChatGPT links.

*   `GET /shared_conversations?order=updated&limit=100&offset={int}`
    *   **Purpose:** Lists all the conversations the user has generated public sharing links for.
*   `GET /share/{share_id}`
    *   **Purpose:** Retrieves the contents of a publicly shared conversation using its share ID.

## 5. Security & Anti-Bot
The endpoints mapped primarily by `gin337/ChatGPTReversed` to bypass Cloudflare and Turnstile.

*   `POST /sentinel/chat-requirements`
    *   **Purpose:** Called immediately before sending a prompt. The server provides a `seed` and `difficulty`. The client must compute a Proof of Work hash (Token z) to prove it is not an automated bot.
*   `POST /backend-anon/sentinel/chat-requirements`
    *   **Purpose:** Same as above, but for unauthenticated (logged-out) users.

## 6. Authentication (`/api/auth`) & Auth0 Flow
These endpoints exist outside the `/backend-api` scope and are used to manage the session tokens that the backend requires. The web client uses a standard Auth0 OAuth flow to get the initial credentials.

*   `GET https://chat.openai.com/auth/login`
    *   **Purpose:** The entry point for web authentication. Redirects the user to the Auth0 provider.
*   `GET https://auth0.openai.com/u/login/identifier`
    *   **Purpose:** The Auth0 endpoint where the user actually submits their email/password or authenticates via Google/Microsoft SSO.
*   `GET /api/auth/csrf`
    *   **Purpose:** Retrieves the CSRF token needed to establish a session after Auth0 login.
*   `GET /api/auth/session`
    *   **Purpose:** The critical endpoint that retrieves your actual `accessToken` (Bearer token) used for all `backend-api` requests.

## 7. Account & Settings
Endpoints for managing the user profile and feature flags.

*   `GET /accounts/check/v4-2023-04-27`
    *   **Purpose:** Validates the user's session, returns account tier (Plus, Team, Free), and capabilities.
*   `GET /settings`
    *   **Purpose:** General user settings configuration.
*   `GET /settings/beta_features`
    *   **Purpose:** Lists which beta features the user has access to, and whether they are toggled on or off (e.g., Video Screen Sharing, Advanced Voice).
*   `POST /settings/beta_features?feature={name}&value={bool}`
    *   **Purpose:** Toggles a specific beta feature on or off.
*   `GET /models`
    *   **Purpose:** Lists the specific LLM models available to the current user (e.g., `gpt-4o`, `gpt-4`, `text-davinci-002-render-sha`).
*   `GET /compliance`
    *   **Purpose:** Account compliance standing and region locks.
*   `GET /codex/usage`
    *   **Purpose:** Retrieves usage statistics specifically for the Codex models (mostly legacy or specialized developer environments).

## 8. Official Public APIs (Found in reverse-engineered proxies)
Some of the repositories (like `ChatGPT-to-API`) actually create proxies that mimic the official OpenAI API structure, or they call the official APIs as fallbacks.
*   `POST https://api.openai.com/v1/chat/completions`
    *   **Purpose:** The standard, paid, official OpenAI API endpoint for generating chat completions.
*   `POST /chat/completions?api-version=2023-05-15`
    *   **Purpose:** The standard Azure OpenAI endpoint format (often used as a backend fallback or proxy target in these repos).

## 9. Local Proxy Endpoints (`/api/ask`, `/api/connections`)
These are **not** OpenAI endpoints. They are local REST API endpoints generated by the `acheong08/ChatGPT-API-server` repository to host the reverse-engineered bot locally.
*   `POST /api/ask`
    *   **Purpose:** A local endpoint you can hit on `localhost:8080` to interact with the Python backend.
*   `GET /api/connections`
    *   **Purpose:** A local endpoint used to check the status of active proxied connections.

## 10. Frontend Routing & Domains
The scripts also extract standard frontend routes and base domains used by the web client.
*   `https://chatgpt.com` & `https://chat.openai.com`
    *   **Note:** Both domains are used interchangeably. OpenAI transitioned from `chat.openai.com` to `chatgpt.com`, and the reverse-engineering scripts check for both.
*   `GET /c/{convo_id}` or `/chat/{convo_id}`
    *   **Purpose:** The literal browser URL you see in your address bar when viewing a specific chat. Used by scrapers to extract the UUID from the URL.

---

# Non-Endpoint Intelligence
Endpoints are only half the battle. The reverse-engineered repositories also reveal the exact internal mechanisms ChatGPT uses for authentication, anti-bot protection, and real-time streaming.

## 1. Required Headers & Workspace Switching
To interact with these endpoints, specific headers are strictly required.
*   **`Authorization: Bearer <accessToken>`**: Retrieved from the `/api/auth/session` endpoint.
*   **`OAI-Device-Id`**: A UUID generated by the client to track the device session.
*   **`Chatgpt-Account-Id` (CRITICAL)**: If this header is omitted, the API assumes you are operating in your personal workspace. To access a Project or Team workspace, you **must** pass the specific workspace/project UUID in this header.

## 2. Proof of Work (PoW) & Anti-Bot (Turnstile)
OpenAI aggressively protects the `/backend-api/conversation` endpoint using a Proof of Work algorithm (detailed in `gin337/ChatGPTReversed`).
1.  **Seed & Difficulty:** The client calls `/sentinel/chat-requirements` and receives a random `seed` string and a `difficulty` hex string.
2.  **Token Generation (Token z):** The client must use a SHA3 hashing algorithm to generate a hash combining the seed, difficulty, and the `OAI-Device-Id` until it finds a hash that meets the difficulty criteria.
3.  **Submission:** This computed hash is then submitted as a specific header (`OpenAI-Sentinel-Proof-Token`) when making the POST request to generate a message.

## 3. SSE (Server-Sent Events) & WebSockets
ChatGPT does not return standard JSON responses for chat generation.
*   **Text Generation:** Uses `text/event-stream` (SSE). The stream sends chunks prefixed with `data: ` containing JSON objects with the delta of the text, ending with a final `data: [DONE]`.
*   **Real-Time Audio/Canvas:** Newer implementations use WebSockets (`wss://chatgpt.com/backend-api/ws`) to maintain a persistent bi-directional connection for streaming audio chunks and canvas edits.

## 4. Conversation Data Schema (Nodes)
ChatGPT's chat history is not a linear array of messages; it is a **Node Tree**.
*   Every message has a `node_id` and a `parent_id`.
*   This structure allows for branching (e.g., when you edit a prompt and regenerate, it creates a new branch on the node tree).
*   Messages denote the author using `author: { role: 'user' | 'assistant' | 'tool' }`.

---

# Practical Applications & Exploitation
With this complete mapping, you essentially have "God Mode" access to the ChatGPT web infrastructure. You no longer have to rely on the official (and expensive) OpenAI API, nor are you restricted by the limitations of the web browser UI. Here is exactly what you can build and how to use it:

## 1. Zero-Cost Autonomous Agents
**What it is:** The ability to write headless scripts (Python/Go) that interact with `gpt-4o` without paying for API credits.
**How it works:**
1. Retrieve your `accessToken` from `/api/auth/session`.
2. Generate an `OAI-Device-Id` (UUID).
3. Call `/sentinel/chat-requirements` to get the PoW seed and difficulty.
4. Calculate the SHA3 hash (Token z).
5. Send a POST request to `/conversation` with the `accessToken` and the PoW hash.
**What it's for:** Building automated bots, coding assistants, or data pipelines that utilize your existing Plus/Team subscription limits rather than a pay-per-token API. 

## 2. The "Free" Unauthenticated API (`/backend-anon/`)
**What it is:** ChatGPT allows users to chat without an account on the web interface. Reverse-engineers mapped this to the `/backend-anon/` structure.
**How it works:**
Instead of passing an `Authorization: Bearer <token>`, you rely entirely on the `OAI-Device-Id` header to establish an anonymous session. You still must solve the `/backend-anon/sentinel/chat-requirements` Proof-of-Work to prove you aren't a simple bot.
**What makes it act as "free":** Because it requires absolutely no account or login, you can technically rotate IPs (proxies) and `OAI-Device-Id` UUIDs to run massive scraping bots or parallel AI queries completely anonymously. There is no account to ban, and you bypass account-level rate limits by simply rotating your fingerprint.

## 3. Total Data Extraction & Local RAG
**What it is:** Automating the downloading of your entire intellectual history across all Workspaces and Projects.
**How it works:**
1. A script authenticates and hits `/projects` to get a list of all your Workspaces.
2. It sets the `Chatgpt-Account-Id` header to "teleport" into each project context.
3. For each project, it hits `/gizmos/{id}/conversations?cursor=0` to get the list of chats.
4. It extracts the `conversation_id`, and then hits `/conversation/{id}` to download the full JSON node tree.
**What it's for:** Building a highly structured local Retrieval-Augmented Generation (RAG) database, bypassing ChatGPT's slow manual export feature, and perfectly capturing the exact branching logic of your chats for local OSINT or LLM fine-tuning.

## 4. Custom Interface Development
**What it is:** Building completely custom frontends (like CLI tools or customized web apps) that act as ChatGPT clients.
**How it works:** By leveraging the WebSockets (`wss://chatgpt.com/backend-api/ws`) and parsing the Server-Sent Events (SSE) from the `data: ` stream, you can pipe real-time text or audio directly into a terminal, bypassing the heavy Chrome UI entirely.
