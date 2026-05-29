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
background.js
LICENSE
manifest.json
popup/options.css
popup/options.html
popup/options.js
popup/popup.html
popup/popup.js
README.md
src/content.js
src/login.js
</directory_structure>

<files>
This section contains the contents of the repository's files.

<file path="background.js">
// Add an event listener for the browser action icon click event
browser.browserAction.onClicked.addListener(() => {
  // Get credentials from storage
  browser.storage.local.get("credentials").then((results) => {
    for (let i = 0; i < results.credentials.length; i++) {
      // Open a new tab for each credential
      new_tab(results.credentials[i].email, results.credentials[i].password);
    }
    console.log("credentials: ", results);
    console.log("credentials.length: ", results.credentials.length);
  });
});

browser.tabs.onRemoved.addListener((tabID) => {
  // Delete the container with the tabID as the name OpenAI (tabID)
  browser.contextualIdentities.query({}).then((containers) => {
    containers.forEach((container) => {
      if (container.name == "OpenAI (" + tabID + ")") {
        browser.contextualIdentities.remove(container.cookieStoreId);
      }
    });
  });
});

function new_tab(email, password) {
  // Create a new container with tabID
  browser.contextualIdentities
    .create({
      name: "OpenAI (tmp)",
      color: "blue",
      icon: "circle",
    })
    .then((container) => {
      // Save email and password to container's cookies
      browser.cookies.set({
        url: "https://auth0.openai.com",
        name: "email",
        value: email,
        storeId: container.cookieStoreId,
      });
      browser.cookies.set({
        url: "https://auth0.openai.com",
        name: "password",
        value: password,
        storeId: container.cookieStoreId,
      });
      // Create a new tab in the container with the tabID and URL
      browser.tabs
        .create({
          cookieStoreId: container.cookieStoreId,
          url: "https://chat.openai.com",
        })
        .then((tabID) => {
          // Add the tabID to the container
          browser.contextualIdentities.update(container.cookieStoreId, {
            name: "OpenAI (" + tabID.id + ")",
          });
        });
    });
}
</file>

<file path="LICENSE">
MIT License

Copyright (c) 2022 ChatGPT-Hackers

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

<file path="manifest.json">
{
  "name": "ChatGPT API Client",
  "description": "A module for interacting with the ChatGPT API",
  "version": "1.2",
  "manifest_version": 2,
  "homepage_url": "https://github.com/ChatGPT-Hackers/firefox-client",
  "background": {
    "scripts": ["background.js"]
  },
  "content_scripts": [
    {
      "matches": [
        "https://chat.openai.com/chat",
        "https://chat.openai.com/chat/*"
      ],
      "js": ["src/content.js"],
      "run_at": "document_end"
    },
    {
      "matches": ["https://chat.openai.com/auth/login", "https://auth0.openai.com/u/login/*"],
      "js": ["src/login.js"],
      "run_at": "document_idle"
    }
  ],
  "permissions": [
    "cookies",
    "webRequest",
    "storage",
    "tabs",
    "contextualIdentities",
    "<all_urls>"
  ],
  "browser_action": {
    "default_title": "Click here to open ChatGPT"
  },
  "options_ui": {
    "page": "popup/options.html",
    "open_in_tab": true
  },
  "browser_specific_settings": {
    "gecko": {
      "id": "acheong@student.dalat.org"
    }
  }
}
</file>

<file path="popup/options.css">
/* options.css */

/* Style the "Add Credential" button */
button#add-credential {
  background-color: #4caf50;
  color: white;
  padding: 12px 20px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  margin-top: 20px;
}

button#add-credential:hover {
  background-color: #45a049;
}

/* Style the credential popup */
.popup {
  display: none; /* Hide the popup by default */
  position: fixed; /* Stay in place */
  z-index: 1; /* Sit on top */
  left: 0;
  top: 0;
  width: 100%; /* Full width */
  height: 100%; /* Full height */
  overflow: auto; /* Enable scroll if needed */
  background-color: rgb(0, 0, 0); /* Fallback color */
  background-color: rgba(0, 0, 0, 0.4); /* Black w/ opacity */
}

/* Style the form in the popup */
.popup form {
  background-color: #fefefe;
  margin: 15% auto; /* 15% from the top and centered */
  padding: 20px;
  border: 1px solid #888;
  width: 80%;
}

/* Style the "Save" button in the popup */
.popup button[type="submit"] {
  background-color: #4caf50;
  color: white;
  padding: 12px 20px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.popup button[type="submit"]:hover {
  background-color: #45a049;
}

/* Style the "Save" button in the form */
button[type="submit"] {
  background-color: #4caf50;
  color: white;
  padding: 12px 20px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  margin-top: 20px;
}

button[type="submit"]:hover {
  background-color: #45a049;
}

/* Style the credential list */
.credential-list {
  margin-top: 20px;
}

table {
  /* Table border */
  border: 1px solid rgb(0, 0, 0);
  border-collapse: collapse;
  width: 80%;
  margin: 1rem
}

td{
  border: 1px solid rgb(0, 0, 0);
  padding: 5px;
}
</file>

<file path="popup/options.html">
<!-- options.html -->
<html>
  <head>
    <script src="options.js" defer></script>
    <!-- options.css -->
    <link rel="stylesheet" href="options.css" />
  </head>
  <body>
    <form>
      <label for="endpoint">Endpoint:</label>
      <input type="text" id="endpoint" name="endpoint" />
    </form>
    <!-- Allow the user to add credentials -->
    <div>
      <button id="add-credential">Add Credential</button>
    </div>
    <!-- Popup to enter email and password -->
    <div class="popup" id="credential-popup">
      <form>
        <label for="email">Email:</label>
        <input type="text" id="email" name="email" />
        <br />
        <label for="password">Password:</label>
        <input type="password" id="password" name="password" />
        <br />
        <button type="submit" id="submit-credential">Save</button>
      </form>
    </div>
    <table>
      <thead>
        <tr>
          <th>Email</th>
          <th>Password</th>
        </tr>
      </thead>
      <tbody id="credential-list"></tbody>
    </table>
    <button type="submit" id="save">Save</button>
    <button type="submit" id="clear">Clear</button>
  </body>
</html>
</file>

<file path="popup/options.js">
// options.js
let credentials = [];

// Get references to the elements we will be interacting with
const addCredentialButton = document.getElementById("add-credential");
const credentialPopup = document.getElementById("credential-popup");
const credentialList = document.getElementById("credential-list");
const submitCredentialButton = document.getElementById("submit-credential");
let save = document.querySelector("#save");
let clear = document.querySelector("#clear");

// Get the endpoint string from the extension's storage (if it exists)
browser.storage.local.get("endpoint").then((result) => {
  // If the endpoint string exists, set the value of the form's endpoint input
  if (result.endpoint) {
    // If endpoint is empty, set it to the default value
    if (result.endpoint == "") {
      result.endpoint = "localhost:8080";
    }
    document.querySelector("#endpoint").value = result.endpoint;
  }
});

// Get the credentials array from the extension's storage (if it exists)
browser.storage.local.get("credentials").then((result) => {
  // If the credentials array exists, add each credential to the list
  if (result.credentials) {
    credentials = result.credentials;
    for (let credential of result.credentials) {
      const tableRow = document.createElement("tr");
      const emailCell = document.createElement("td");
      emailCell.textContent = credential.email;
      const passwordCell = document.createElement("td");
      passwordCell.textContent = credential.password;
      tableRow.appendChild(emailCell);
      tableRow.appendChild(passwordCell);
      credentialList.appendChild(tableRow);
    }
  }
});

// Add an event listener to the "Add Credential" button
addCredentialButton.addEventListener("click", () => {
  // Show the credential popup
  credentialPopup.style.display = "block";
});

// Add an event listener to the "Save" button in the popup
submitCredentialButton.addEventListener("click", (event) => {
  event.preventDefault(); // prevent the form from being submitted

  // Get the email and password from the form
  const email = document.getElementById("email").value;
  const password = document.getElementById("password").value;

  // Add the credentials to the list
  const tableRow = document.createElement("tr");
  const emailCell = document.createElement("td");
  emailCell.textContent = email;
  const passwordCell = document.createElement("td");
  passwordCell.textContent = password;
  tableRow.appendChild(emailCell);
  tableRow.appendChild(passwordCell);
  credentialList.appendChild(tableRow);

  // Save the credentials to credentials array
  credentials.push({ email: email, password: password });

  // Clear the form and close the popup
  document.getElementById("email").value = "";
  document.getElementById("password").value = "";
  credentialPopup.style.display = "none";

  // Save the credentials to the extension's storage
  browser.storage.local.set({ credentials: credentials });
});

// Add an event listener for the submit event
save.addEventListener("click", function (event) {
  // Prevent the form from being submitted
  event.preventDefault();

  // Get the endpoint string from the form
  const endpoint = document.querySelector("#endpoint").value;

  // Save the endpoint string to the extension's storage
  browser.storage.local.set({ endpoint: endpoint });
});

clear.addEventListener("click", function (event) {
  // Prevent the form from being submitted
  event.preventDefault();

  // Clear the endpoint string from the extension's storage
  browser.storage.local.remove("endpoint");

  // Clear the endpoint string from the form
  document.querySelector("#endpoint").value = "";

  // Clear credentials 
  credentials = [];
  browser.storage.local.remove("credentials");
  // Refresh the page
  location.reload();
});
</file>

<file path="popup/popup.html">

</file>

<file path="popup/popup.js">

</file>

<file path="src/content.js">
const BASE_URL = "https://chat.openai.com";
const CHAT_URL = `${BASE_URL}/chat`;
const BACKEND_URL = `${BASE_URL}/backend-api/conversation`;
const SESSION_URL = `${BASE_URL}/api/auth/session`;

browser.storage.local.get("endpoint").then((result) => {
  const endpoint = result.endpoint || "localhost:8080";
  const wsRoute = "ws://" + endpoint + "/client/register";
  const ws = new WebSocket(wsRoute);
  console.info("Connecting to " + wsRoute);

  window.onunload = function () {
    console.info("Connection closed");
    ws.close();
  };

  ws.onerror = function (error) {
    console.error("An error occured");
    console.error(error);
    console.info("Connection closed");
    ws.close();
  };

  ws.onopen = function () {
    console.info("Connection opened");
    ws.onmessage = function (event) {
      const data = JSON.parse(event.data);

      console.log(`data ${JSON.stringify(data)}`);
      switch (data.message) {
        case "Connection id":
          handleConnectionId(ws, data);
          break;
        case "ping":
          handlePing(ws, data);
          break;
        case "ChatGptRequest":
          handleChatGptRequest(ws, data);
          break;
        default:
          console.error("Unknown message: " + data.message);
          break;
      }
    };
  };

  ws.onclose = function () {
    console.info("Connection closed");
    delete connectionId;
  };
});

function handleConnectionId(ws, data) {
  // Get connection id from cookies (if it exists)
  const cookies = document.cookie.split(";");
  let storedConnectionId = "";
  for (let i = 0; i < cookies.length; i++) {
    const cookie = cookies[i];
    if (cookie.includes("connectionId")) {
      storedConnectionId = cookie.split("=")[1];
    }
  }
  console.debug(`storedConnectionId ${storedConnectionId}`);
  // If it exists, send it to the server
  if (storedConnectionId) {
    sendWebSocketMessage(ws, storedConnectionId, "Connection id", "");
  } else {
    sendWebSocketMessage(ws, data.id, "Connection id", "");
    // Store connectionId in cookie
    document.cookie = "connectionId=" + data.id;
  }
}

function handlePing(ws, data) {
  sendWebSocketMessage(ws, data.id, "pong", "");
}

async function handleChatGptRequest(ws, data) {
  // Construct API request
  const requestData = JSON.parse(data.data);
  // If conversation_id is "", make it undefined
  if (requestData.conversation_id == "") {
    requestData.conversation_id = undefined;
  }
  // Payload
  const payload = {
    action: "next",
    messages: [
      {
        id: requestData.message_id,
        role: "user",
        content: { content_type: "text", parts: [requestData.content] },
      },
    ],
    parent_message_id: requestData.parent_id,
    conversation_id: requestData.conversation_id,
    model: "text-davinci-002-render",
  };

  try {
    const accessToken = await getAccessToken();
    const conversationResponse = await sendChatRequest(accessToken, payload);
    const responseData = createResponseData(conversationResponse);
    sendWebSocketMessage(ws, data.id, "ChatGptResponse", responseData);
  } catch (error) {
    console.error(error);
    sendWebSocketMessage(ws, data.id, "error", "Unknown error", error);
  }
}

async function getAccessToken() {
  try {
    const sessionResponse = await fetchData(SESSION_URL);
    const sessionResponseJson = await sessionResponse.json();
    return sessionResponseJson.accessToken;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

async function sendChatRequest(accessToken, payload) {
  try {
    const response = await fetchData(BACKEND_URL, {
      method: "POST",
      headers: {
        Accept: "text/event-stream",
        Authorization: `Bearer ${accessToken}`,
        "Content-Type": "application/json",
        "X-Openai-Assistant-App-Id": "",
        Connection: "close",
        Referer: CHAT_URL,
      },
      body: JSON.stringify(payload),
    });
    const conversationResponse = await response.text();
    try {
      // Check if conversationResponse can be parsed as JSON
      const respJson = JSON.parse(conversationResponse);
      if (respJson.detail) {
        console.error(`Error: ${respJson.detail}`);
        throw new Error(`Error: ${respJson.detail}`);
      }
    } catch (e) {
      console.log("Not JSON");
    }
    // Split data on "data: " prefix
    const dataArray = conversationResponse.split("data: ");
    // Get the second last element of the array
    const lastElement = JSON.parse(dataArray[dataArray.length - 2]);
    console.log(lastElement);
    return lastElement;
  } catch (error) {
    console.error(error);
    throw error;
  }
}

function createResponseData(conversationResponse) {
  return JSON.stringify({
    response_id: conversationResponse.message.id,
    conversation_id: conversationResponse.conversation_id,
    content: conversationResponse.message.content.parts[0],
  });
}

/**
 * @param {WebSocket} ws The websocket server
 * @param {string} id The data id
 * @param {string} message The message type
 * @param {string} data The data to send
 * @param {string} error If there is an error, the error message
 */
function sendWebSocketMessage(ws, id, message, data, error) {
  try {
    const wsMessage = {
      id,
      message,
      data,
      error,
    };
    ws.send(JSON.stringify(wsMessage));
  } catch (error) {
    console.error("Error sending websocket message");
    console.error(error);
  }
}

/**
 * Fetch data from an url with optional options.
 * It will throw an error if the response is not ok, or if the status code is not 200.
 *
 * @async
 * @param {String} url The url to fetch from
 * @param {RequestInit} options The options to pass to fetch
 * @returns The response
 */
async function fetchData(url, options = {}) {
  const response = await window.fetch(url, options);
  if (!response.ok) {
    throw new Error(`Request failed with status code: ${response.status}`);
  }

  if (response.status !== 200) {
    // This isn't the best way to handle the status code check, but it works for now
    throw new Error(`Wrong response code: ${response.status}`);
  }

  return response;
}
</file>

<file path="src/login.js">
/// login.js
console.log("login.js loaded");

// Find the button with the text "Log in"
let button = document.querySelector(
  ".btn.flex.justify-center.gap-2.btn-primary"
);

if (button && button.innerText === "Log in") {
  button.click();
} else {
  console.error("button not found");
}

// Get credentials from cookies
const cookies = document.cookie.split("; ");
// Get email and password from cookies
let email = cookies.find((cookie) => cookie.startsWith("email=")).split("=")[1];
let password = cookies
  .find((cookie) => cookie.startsWith("password="))
  .split("=")[1];

function setUsername() {
  let emailForm = document.getElementById("username");
  if (emailForm) {
    emailForm.value = email;
    console.log("email: ", email);
  } else {
    setTimeout(setUsername, 1000);
  }
}

setUsername();

function setPassword() {
  let passwordForm = document.getElementById("password");
  if (passwordForm) {
    passwordForm.value = password;
    console.log("password: ", password);
  } else {
    setTimeout(setPassword, 1000);
  }
}
setPassword();

// c8fca5323 cb6b7c993 cee1c07cc c850d9a60 _button-login-password
button = document.querySelector("._button-login-password");
if (button) {
  button.click();
} else {
  console.error("button not found");
}
</file>

<file path="README.md">
# ChatGPT API Agent (Firefox version)

# Setup
1. Download from [Mozilla Addons](https://addons.mozilla.org/en-US/firefox/addon/chatgpt-api-client/)

# Running
1. Go to extension preferences
![image](https://user-images.githubusercontent.com/36258159/209443449-73ca41c3-39ad-4429-b1b7-028b508dddff.png)
![image](https://user-images.githubusercontent.com/36258159/209443463-7ca046e3-758b-4541-8b9d-f0f5eeebbc58.png)
2. Configure endpoint from the [server](https://github.com/ChatGPT-Hackers/ChatGPT-API-server)
3. Add emails/passwords to the preferences
4. Press save
![image](https://user-images.githubusercontent.com/36258159/209443551-ce03ce90-d1de-4e42-8b35-df46bb70c62b.png)
5. Click on the extension
![image](https://user-images.githubusercontent.com/36258159/209443565-6bb9866a-99d2-4947-96e9-2934c93db80c.png)
This will spawn the same number of tabs as there are emails/passwords
6. Wait for it to load
7. Complete the captcha and press continue (email/password autofills)
![image](https://user-images.githubusercontent.com/36258159/209443617-d96ee8d2-a016-4fa1-85da-f815a38e0087.png)
8. After that, it will autofill the password and continue to the chat site.

Done. It connects to the endpoint and you can leave it open.

<br>
<br>
<br>

# Firefox Docker (optional)

```yaml
version: '3.3'
services:
    firefox:
        container_name: firefox
        ports:
            - '5800:5800'
        volumes:
            - '<host folder path>:/config:rw'
        image: jlesage/firefox

```
1. create a folder that will contain the app data for firefox
2. access container via `<ip-address>:5800` and finish the firefox setup
3. procceed to follow <a href="#top">step 1</a> in Setup section
4. now follow steps in <a href="#top">Running section</a> 

# Contributing
In order to develop locally you need to use guide provided by Mozilla: https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/Your_first_WebExtension#trying_it_out and follow next steps:
1. Clone this repository
2. Go to `about:debugging` in Firefox
3. Specify this directory as a temporary extension
4. It will be loaded on top of the existing extension if you have one
5. You can debug your new feature.
</file>

</files>
