This file is a merged representation of the entire Gist collection, combined into a single document.

<file_summary>
This section contains a summary of this file.

<purpose>
This file contains a packed representation of all collected 0xdevalias reverse-engineering gists.
It is designed to be easily consumable by AI systems for analysis, code review, or other automated processes.
</purpose>

<file_format>
The Gist collection is structured under the directories:
- sources/gists/0xdevalias-frontend-endpoint-mapping/
- sources/gists/devalias-chatgpt-export-history/
- sources/gists/devalias-chatgpt-frontend-diff/
</file_format>

<usage_guidelines>
- This file should be treated as read-only. Any changes should be made to the original Gists, not this packed version.
- Use the file path blocks to distinguish between different Gist files.
</usage_guidelines>
</file_summary>

<directory_structure>
- sources/gists/0xdevalias-frontend-endpoint-mapping/chatgpt-reverse-engineering-and-deep-dive-code-exploration.md
- sources/gists/devalias-chatgpt-export-history/chatgpt-api-export.md
- sources/gists/devalias-chatgpt-frontend-diff/chatgpt-chunk-9087.diff
</directory_structure>

<files>
=== 
File: sources/gists/0xdevalias-frontend-endpoint-mapping/chatgpt-reverse-engineering-and-deep-dive-code-exploration.md
===
# Reverse engineering ChatGPT's frontend web app + deep dive explorations of the code

Figured it would make sense to create a single gist collating my previous deep dive explorations and notes.. so here it is!

## Table of Contents

<!-- TOC start (generated with https://derlin.github.io/bitdowntoc/) -->
- [My Deep Dives (and others past feature leaks)](#my-deep-dives-and-others-past-feature-leaks)
   - [Jun 14th, 2023 (and all future updates)](#jun-14th-2023-and-all-future-updates)
   - [Jun 13th, 2023](#jun-13th-2023)
   - [May 16th, 2023](#may-16th-2023)
   - [March 25th, 2023](#march-25th-2023)
   - [February 10th, 2023](#february-10th-2023)
   - [February 5th, 2023](#february-5th-2023)
   - [January 19th, 2023](#january-19th-2023)
   - [January 1st, 2023](#january-1st-2023)
- [Everything-ChatGPT](#everything-chatgpt)
- [See Also](#see-also)
   - [My Other Related Deepdive Gist's and Projects](#my-other-related-deepdive-gists-and-projects)
<!-- TOC end -->

## My Deep Dives (and others past feature leaks)

### Jun 14th, 2023 (and all future updates)

You'll find my latest deep dives and continuing research within the following repo:

- https://github.com/0xdevalias/chatgpt-source-watch : Analyzing the evolution of ChatGPT's codebase through time with curated archives and scripts

For the equivalent ongoing content that used to be included on this Gist, you probably want to look at the `CHANGELOG.md`:

- https://github.com/0xdevalias/chatgpt-source-watch/blob/main/CHANGELOG.md

You can also check out the announcement posts on Twitter, Reddit, and HackerNews:

- https://twitter.com/_devalias/status/1672108774432321536
- https://www.reddit.com/r/ChatGPT/comments/14gp7gr/ive_just_released_chatgpt_source_watch_dedicated/
- https://news.ycombinator.com/item?id=36442959

### Jun 13th, 2023

I haven't looked at the ChatGPT source for a while, so I don't have my own personal discoveries for this one, but documenting the leaked features found by others here:

- [[YouTube] ChatGPT Leaked Features! - File Upload, Workspaces, Profiles | BusinessGPT on Horizon?](https://www.youtube.com/watch?v=MfGX_kp8F-w)
  - https://www.reddit.com/r/ChatGPT/comments/144cfzg/upcoming_chatgpt_features_file_uploading_profiles/
    - https://webcache.googleusercontent.com/search?q=cache:https://www.reddit.com/r/ChatGPT/comments/144cfzg/upcoming_chatgpt_features_file_uploading_profiles/
      - Posted: `Friday, June 9, 2023 at 1:34:10 AM GMT+10`
  - https://www.reddit.com/r/ChatGPT/comments/13m5n9z/chatgpt_is_adding_chat_sharing/
    - https://webcache.googleusercontent.com/search?q=cache:https://www.reddit.com/r/ChatGPT/comments/13m5n9z/chatgpt_is_adding_chat_sharing/
      - Posted: `Saturday, May 20, 2023 at 5:42:18 AM GMT+10` (so I beat this leak by 4 days.. see below)

### May 16th, 2023

https://twitter.com/_devalias/status/1658104103090651140

```
Another day, another ChatGPT unreleased 'INTERNAL ONLY' feature: Share Chat

https://chat.openai.com/_next/static/chunks/pages/_app-ab949dad0ea9d6e3.js

createShareLink: POST /share/create
updateShareLink: PATCH /share/:share_id
deleteShareLink: DELETE /share/:share_id

https://chat.openai.com/_next/static/chunks/734-d34f3efd388e555d.js
```

Cross-posted to HackerNews: https://news.ycombinator.com/item?id=35948098

### March 25th, 2023

https://twitter.com/_devalias/status/1639377155347083264

```
Given yesterday's announcement of ChatGPT plugin support in waitlisted beta; I figured it might be a good time to look at the frontend source again and see what I can turn up..

You can find my previous explorations here:
- 2: https://twitter.com/_devalias/status/1623962204252033026
- 1: https://twitter.com/_devalias/status/1622102312994099200
```

### February 10th, 2023

https://twitter.com/_devalias/status/1623962204252033026

```
So I figured I would have a little look into the minified ChatGPT frontend source again today and see what I stumbled upon..

You can see my last exploration in the attached thread from the other day.
```

### February 5th, 2023

https://twitter.com/_devalias/status/1622102312994099200

```
What's that? ChatGPT [private alpha - confidential]? But i'm not even part of ChatGPT Plus.. Is it still confidential if I am just reading publicly available minified source code?🤔

I wouldn't mind using a 'Web search' or 'Calculator' though..
```

### January 19th, 2023

https://gist.github.com/0xdevalias/4e54bb28a02db5357ea4fa3a872fc5fc

> Exploring ChatGPT API's for exporting all history as markdown, etc.

### January 1st, 2023

The following wasn't found by me, but adding here as a historical archive:

- https://www.reddit.com/r/GPT3/comments/1002dym/upcoming_potential_chatgpt_features_not_released/
  - > I made a list of some of the hidden ChatGPT features here https://twitter.com/eeeziii/status/1609069324643471363
    > Also made an extension for anyone who wants to access some of those features and a couple of other features now: https://chrome.google.com/webstore/detail/superpower-chatgpt/amhmeenmapldpjdedekalnfifgnpfnkc
  - Crossposts
    - Crosspost: https://www.reddit.com/r/ChatGPT/comments/1003u4i/upcoming_potential_chatgpt_features_not_released/
    - Semi-Crosspost: https://www.reddit.com/r/ChatGPT/comments/1001rh9/upcoming_features_of_chatgpt_not_released_yet/
    - Semi-Crosspost: https://www.reddit.com/r/ChatGPT/comments/1018uh8/chatgpt_likely_upcoming_features/
  - https://twitter.com/eeeziii/status/1609069324643471363
    - Pause Completions
    - Copy thread to clipboard
    - Palette
    - Palette Features (link, snippet, promptys, editPrompt)
    - Add text from link
    - Completions
    - Artifacts (/artifacts)

## Everything-ChatGPT

I also recently shared a number of these references in a GitHub issue to the 'everything-chatgpt' project, which seems to be similarly exploring the features and capabilities of the ChatGPT frontend app + API's:

- https://github.com/terminalcommandnewsletter/everything-chatgpt/issues/6

As well as noting a few other bits of information that seemed outdated/missing:

- https://github.com/terminalcommandnewsletter/everything-chatgpt/issues/7
  - > Feature not covered: API: `GET /backend-api/settings/beta_features`
- https://github.com/terminalcommandnewsletter/everything-chatgpt/issues/8
  - > Section out-of-date: Models - `GET /backend-api/models`
- https://github.com/terminalcommandnewsletter/everything-chatgpt/issues/9
  - > Section out-of-date: User Data - `GET /backend-api/accounts/check`
- https://github.com/terminalcommandnewsletter/everything-chatgpt/issues/10
  - > Feature not covered: Plugins - `GET /backend-api/aip/p`

## See Also

### My Other Related Deepdive Gist's and Projects

- https://gist.github.com/0xdevalias
- https://github.com/0xdevalias/chatgpt-source-watch : Analyzing the evolution of ChatGPT's codebase through time with curated archives and scripts.
- [Deobfuscating / Unminifying Obfuscated Web App Code (0xdevalias gist)](https://gist.github.com/0xdevalias/d8b743efb82c0e9406fc69da0d6c6581#deobfuscating--unminifying-obfuscated-web-app-code)
- [Reverse Engineering Webpack Apps (0xdevalias gist)](https://gist.github.com/0xdevalias/8c621c5d09d780b1d321bfdb86d67cdd#reverse-engineering-webpack-apps)
- [React Server Components, Next.js v13+, and Webpack: Notes on Streaming Wire Format (`__next_f`, etc) (0xdevalias' gist))](https://gist.github.com/0xdevalias/ac465fb2f7e6fded183c2a4273d21e61#react-server-components-nextjs-v13-and-webpack-notes-on-streaming-wire-format-__next_f-etc)
- [Fingerprinting Minified JavaScript Libraries / AST Fingerprinting / Source Code Similarity / Etc (0xdevalias' gist)](https://gist.github.com/0xdevalias/31c6574891db3e36f15069b859065267#fingerprinting-minified-javascript-libraries--ast-fingerprinting--source-code-similarity--etc)
  - [JavaScript Web App Reverse Engineering - Module Identification (0xdevalias' gist)](https://gist.github.com/0xdevalias/28c18edfc17606f09cf413f97e404a60#javascript-web-app-reverse-engineering---module-identification)
  - [Reverse Engineered Webpack Tailwind-Styled-Component (0xdevalias' gist)](https://gist.github.com/0xdevalias/916e4ababd3cb5e3470b07a024cf3125#reverse-engineered-webpack-tailwind-styled-component)
- [Bypassing Cloudflare, Akamai, etc (0xdevalias gist)](https://gist.github.com/0xdevalias/b34feb567bd50b37161293694066dd53#bypassing-cloudflare-akamai-etc)
- [Debugging Electron Apps (and related memory issues) (0xdevalias gist)](https://gist.github.com/0xdevalias/428e56a146e3c09ec129ee58584583ba#debugging-electron-apps-and-related-memory-issues)
- [devalias' Beeper CSS Hacks (0xdevalias gist)](https://gist.github.com/0xdevalias/3d2f5a861335cc1277b21a29d1285cfe#beeper-custom-theme-styles)
- [Reverse Engineering Golang (0xdevalias' gist)](https://gist.github.com/0xdevalias/4e430914124c3fd2c51cb7ac2801acba#reverse-engineering-golang)
- [Reverse Engineering on macOS (0xdevalias' gist)](https://gist.github.com/0xdevalias/256a8018473839695e8684e37da92c25#reverse-engineering-on-macos)

=== 
File: sources/gists/devalias-chatgpt-export-history/chatgpt-api-export.md
===
_(Edit: I've started collating all of my ChatGPT deep dives and explorations at the following new gist: https://gist.github.com/0xdevalias/4ac297ee3f794c17d0997b4673a2f160)_

---

Exploring ChatGPT API's for exporting all history as markdown, etc.

## `/backend-api/conversations`

List the history of past conversations:

`/backend-api/conversations?offset=0&limit=100`:

```js
const result = await fetch("https://chat.openai.com/backend-api/conversations?offset=0&limit=100", {
  "headers": {
    "accept": "*/*",
    "authorization": "Bearer REDACTEDREDACTEDREDACTEDREDACTED",
    "content-type": "application/json",
  },
  "method": "GET",
  "mode": "cors",
  "credentials": "include"
});

const resultJson = await result.json();
```

The `resultJson` format looks like this (note that despite `total` saying `93`, I think there were only `92` entries in `items`):

```js
{
  items: [{…}, /*..snip..*/, {…}],
  limit: 100,
  offset: 0,
  total: 93
}
```

`resultJson.items` looks like this:

```js
[
  {
    create_time: "2023-01-18T20:43:46.689880",
    id: "11111111-222-3333-4444-555555555556",
    title: "REDACTED\n"
  },
  /* ..snip.. */
  {
    create_time: "2023-01-18T20:43:46.689880",
    id: "11111111-222-3333-4444-555555555555",
    title: "REDACTED\n"
  }
]
```

Limit must be `<=` `100`, otherwise you'll get an error like this:

```json
{
    "detail": [
        {
            "loc": [
                "query",
                "limit"
            ],
            "msg": "ensure this value is less than or equal to 100",
            "type": "value_error.number.not_le",
            "ctx": {
                "limit_value": 100
            }
        },
        {
            "loc": [
                "query",
                "limit"
            ],
            "msg": "ensure this value is less than or equal to 100",
            "type": "value_error.number.not_le",
            "ctx": {
                "limit_value": 100
            }
        }
    ]
}
```

## `/backend-api/conversation/:conversationId`

Retrieve the details of a specific conversation:

```js
const result = await fetch("https://chat.openai.com/backend-api/conversation/11111111-222-3333-4444-555555555555", {
  "headers": {
    "authorization": "Bearer REDACTEDREDACTEDREDACTEDREDACTED",
    "content-type": "application/json",
  },
  "method": "GET",
  "mode": "cors",
  "credentials": "include"
});

const resultJson = await result.json();
```

The `resultJson` format looks like this:

```js
{
  create_time: 1673323045.953257,
  current_node: "AAAAAAAA-222-3333-4444-555555555555",
  mapping: {
    AAAAAAAA-222-3333-4444-555555555555: { /* ..snip.. */ },
    BBBBBBBB-222-3333-4444-555555555555: {
      id: "BBBBBBBB-222-3333-4444-555555555555",
      message: {
        content: {
          content_type: "text",
          parts: ["REDACTED"]
        },
        create_time: 1673323754.300144,
        end_turn: null,
        metadata: {
          timestamp_: "absolute"
        },
        recipient: "all",
        role: "user",
        update_time: null,
        user: null,
        weight: 1,
      },
      children: ["CCCCCCCC-222-3333-4444-555555555555"],
      parent: "DDDDDDDD-222-3333-4444-555555555555",
    },
    /* ..snip.. */
  },
  moderation_results: [],
  title: "REDACTED"
}
```

=== 
File: sources/gists/devalias-chatgpt-frontend-diff/chatgpt-chunk-9087.diff
===
--- Unsaved view (531)
+++ Unsaved view (521)
@@ -1,10 +1,15 @@
-// OLD unpacked/_next/static/chunks/9087.js JSON
+// NEW unpacked/_next/static/chunks/9087.js JSON
 {
     "AccountPaymentModal.accountErrorWarning":
     {
         "defaultMessage": "The account management page encountered an error. Please try again. If the problem continues, please visit help.openai.com.",
         "description": "Error toast when account page has an error"
     },
+    "AccountPaymentModal.haveExistingPlan":
+    {
+        "defaultMessage": "Have an existing plan?",
+        "description": "Have an existing plan question message"
+    },
     "AccountPaymentModal.modalFooterCapabilities":
     {
         "defaultMessage": "Need more capabilities? See <link> ChatGPT Enterprise </link>",
@@ -25,11 +30,6 @@
         "defaultMessage": "I need help with a billing issue",
         "description": "Need help with billing message"
     },
-    "AccountPaymentModel.haveExistingPlan":
-    {
-        "defaultMessage": "Have an existing plan?",
-        "description": "Have an existing plan question message"
-    },
     "AgeVerificationInterstitial.description0":
     {
         "defaultMessage": "To continue using ChatGPT, you need to complete a brief age verification check",
@@ -90,12 +90,17 @@
         "defaultMessage": "{name} logo",
         "description": "The alt text for the app logo image"
     },
-    "AutPage.privacy":
+    "AuthPage.cookies":
+    {
+        "defaultMessage": "Manage cookies",
+        "description": "Open manage cookies modal"
+    },
+    "AuthPage.privacy":
     {
         "defaultMessage": "Privacy policy",
         "description": "Privacy policy link label"
     },
-    "AutPage.terms":
+    "AuthPage.terms":
     {
         "defaultMessage": "Terms of use",
         "description": "Terms of use link label"
@@ -220,6 +225,16 @@
         "defaultMessage": "Shrink",
         "description": "Label for the Shrink button in the DevtoolToolbar"
     },
+    "ConfirmModal.cancelLabel":
+    {
+        "defaultMessage": "Cancel",
+        "description": "Cancel button label"
+    },
+    "ConfirmModal.confirmLabel":
+    {
+        "defaultMessage": "Confirm",
+        "description": "Confirm button label"
+    },
     "ContextConnectorPicker.attachFiles":
     {
         "defaultMessage": "Attach files",
@@ -260,11 +275,46 @@
         "defaultMessage": "Upload from {connector_name}",
         "description": "Dropdown label to upload from a given connector"
     },
+    "ConversationPrivacyIndicator.memoryOff":
+    {
+        "defaultMessage": "Memory Off",
+        "description": "Label for Memory Off"
+    },
+    "ConversationPrivacyIndicator.memoryOffTooltip":
+    {
+        "defaultMessage": "ChatGPT won't remember anything you talk about in this conversation.",
+        "description": "Tooltip for Memory Off"
+    },
+    "ConversationPrivacyIndicator.temporaryChat":
+    {
+        "defaultMessage": "Temporary Chat",
+        "description": "Label for Temporary Chat"
+    },
+    "ConversationPrivacyIndicator.temporaryChatTooltip":
+    {
+        "defaultMessage": "Temporary Chats won't appear in your history, and ChatGPT won't remember anything you talk about.",
+        "description": "Tooltip for Temporary Chat"
+    },
     "ConversationTurn.anonymousName":
     {
         "defaultMessage": "Anonymous",
         "description": "Name on a conversation turn when the original author is unknown or anonymous"
     },
+    "ConversationTurn.badResponseTooltip":
+    {
+        "defaultMessage": "Bad response",
+        "description": "Tooltip on thumbs down message icon button"
+    },
+    "ConversationTurn.goodResponseTooltip":
+    {
+        "defaultMessage": "Good response",
+        "description": "Tooltip on thumbs up message icon button"
+    },
+    "ConversationTurn.regenerateTooltip":
+    {
+        "defaultMessage": "Regenerate",
+        "description": "Tooltip on regenerate response icon button"
+    },
     "ConversationTurn.you":
     {
         "defaultMessage": "You",
@@ -330,6 +380,11 @@
         "defaultMessage": "Accept all",
         "description": "Consent to analytics cookies button"
     },
+    "CookieConsentBanner.authedDescription.0":
+    {
+        "defaultMessage": "We use cookies and similar technologies to deliver, maintain, improve our services and for security purposes. Check our <cookiePolicy>cookie policy</cookiePolicy> for details. Click 'Accept all' to let OpenAI and partners use cookies for these purposes. Click 'Reject all' to say no to cookies, except those that are strictly necessary. Choose <manageLink>manage cookies</manageLink> to change your preferences.",
+        "description": "Banner displayed describing cookie usage and request for consent"
+    },
     "CookieConsentBanner.reject":
     {
         "defaultMessage": "Reject all",
@@ -338,7 +393,12 @@
     "CookieConsentBanner.title":
     {
         "defaultMessage": "We use cookies to provide, improve, and protect our services. Visit our <privacyPolicyLink>privacy policy</privacyPolicyLink> to learn more. You can manage your cookie preferences in your <settingsPanel>settings panel</settingsPanel>. <learnMoreLink>Learn more</learnMoreLink>.",
-        "description": "Explanation of why ChatGPT needs to verify age"
+        "description": "Banner displayed describing cookie usage and request for consent"
+    },
+    "CookieConsentBanner.unauthedDescription.0":
+    {
+        "defaultMessage": "We use cookies and similar technologies to deliver, maintain, improve our services and for security purposes. Check our <cookiePolicy>cookie policy</cookiePolicy> for details. Click 'Accept all' to let OpenAI and partners use cookies for these purposes. Click 'Reject all' to say no to cookies, except those that are strictly necessary. Choose <manageLink>manage cookies</manageLink> to change your preferences.",
+        "description": "Banner displayed describing cookie usage and request for consent"
     },
     "CookieConsentBanner.updateFailure":
     {
@@ -360,6 +420,11 @@
         "defaultMessage": "Copy",
         "description": "Text displayed when the content can be copied"
     },
+    "CopyButton.copyTooltip":
+    {
+        "defaultMessage": "Copy",
+        "description": "Tooltip on copy message icon button"
+    },
     "DebugSidebar.closeSidebar":
     {
         "defaultMessage": "Close sidebar",
@@ -575,6 +640,11 @@
         "defaultMessage": "Browse files",
         "description": "Option to select a file via browser's file picker"
     },
+    "FreeColumn.haveExistingPlan":
+    {
+        "defaultMessage": "Have an existing plan? See <link>billing help</link>",
+        "description": "Have an existing plan question message"
+    },
     "FreeTrialMenuItem.daysOfBenefit":
     {
         "defaultMessage": "{referralTrialDurationDays, plural, one {one day free} other {# days free}}",
@@ -585,6 +655,21 @@
         "defaultMessage": "{referralTrialDurationMonths, plural, one {one month free} other {# months free}}",
         "description": "Duration of the referral trial benefits in months"
     },
+    "GizmoConversationHeader.approveGPTButton":
+    {
+        "defaultMessage": "Approve for workspace",
+        "description": "A button to approve the GPT for the workspace"
+    },
+    "GizmoConversationHeader.newChat":
+    {
+        "defaultMessage": "New Chat",
+        "description": "Button to start a new chat"
+    },
+    "GizmoConversationHeader.usingAsOwner":
+    {
+        "defaultMessage": "Only workspace owners can use this GPT.",
+        "description": "A notice informing the user that only workspace owners can use this GPT"
+    },
     "GizmoConversationOptionsDropdown.debugChat":
     {
         "defaultMessage": "Debug Chat",
@@ -610,11 +695,36 @@
         "defaultMessage": "Web browsing is disabled for your workspace. Contact your admin to enable it.",
         "description": "Web browsing disable tooltip"
     },
+    "GizmoFeedbackEmailModal.includeFrom":
+    {
+        "defaultMessage": "Include my email address {email}",
+        "description": "Include from checkbox label"
+    },
+    "GizmoFeedbackEmailModal.sendTo":
+    {
+        "defaultMessage": "Send To {name}",
+        "description": "Send button text"
+    },
+    "GizmoFeedbackEmailModal.successToast":
+    {
+        "defaultMessage": "Feedback sent",
+        "description": "Success toast message"
+    },
+    "GizmoFeedbackEmailModal.title":
+    {
+        "defaultMessage": "Feedback about {name}",
+        "description": "Title of the feedback email modal"
+    },
     "GizmoInformation.about":
     {
         "defaultMessage": "About",
         "description": "Label for the model details button"
     },
+    "GizmoInformation.archiveChat":
+    {
+        "defaultMessage": "Archive chat",
+        "description": "Archive chat button"
+    },
     "GizmoInformation.copiedURL":
     {
         "defaultMessage": "Copied URL",
@@ -645,6 +755,11 @@
         "defaultMessage": "Favorite",
         "description": "Label for the favorite button"
     },
+    "GizmoInformation.feedbackEmail":
+    {
+        "defaultMessage": "Send feedback",
+        "description": "Label for the send support email button"
+    },
     "GizmoInformation.gpt3_5":
     {
         "defaultMessage": "GPT-3.5",
@@ -675,6 +790,11 @@
         "defaultMessage": "Privacy settings",
         "description": "Label for the model's privacy settings button"
     },
+    "GizmoInformation.removeGPT":
+    {
+        "defaultMessage": "Remove GPT from workspace",
+        "description": "Label for the remove gpt button"
+    },
     "GizmoInformation.renameChat":
     {
         "defaultMessage": "Rename",
@@ -695,6 +815,11 @@
         "defaultMessage": "Share",
         "description": "Share chat button"
     },
+    "GizmoInformation.shareChatLong":
+    {
+        "defaultMessage": "Share chat",
+        "description": "Share chat button"
+    },
     "GizmoInformation.viewDetails":
     {
         "defaultMessage": "View details",
@@ -710,9 +835,9 @@
         "defaultMessage": "Temporary Chat",
         "description": "Temporary chat title"
     },
-    "GizmoLanding.temporaryChatDescription":
-    {
-        "defaultMessage": "Your GPT won't remember anything you talk about, and this chat won't show up in your history or be used to train our models.",
+    "GizmoLanding.temporaryChatDescription.3":
+    {
+        "defaultMessage": "This chat won't appear in history, use memory, or be used to train our models. For safety purposes, we may keep a copy for up to 30 days.",
         "description": "Temporary chat description"
     },
     "GizmoPrivacySettings.actions":
@@ -795,6 +920,16 @@
         "defaultMessage": "Upgrade to Plus to create your own <bolded>GPT</bolded>",
         "description": "Upgrade CTA for free users"
     },
+    "GizmoSidebarList.showLess":
+    {
+        "defaultMessage": "See less",
+        "description": "Button text to collapse sidebar items"
+    },
+    "GizmoSidebarList.showMoreItems":
+    {
+        "defaultMessage": "{numMore} more",
+        "description": "More button text to show more sidebar items"
+    },
     "GizmoSocialRow.addDomain":
     {
         "defaultMessage": "Verify new domain",
@@ -815,6 +950,26 @@
         "defaultMessage": "Verify",
         "description": "Label for social verification button"
     },
+    "HistoryGizmoItem.archiveChatOnboarding":
+    {
+        "defaultMessage": "You can view archived chats in Settings",
+        "description": "Archived chats onboarding text"
+    },
+    "HistoryGizmoItem.archiveTooltip":
+    {
+        "defaultMessage": "Archive",
+        "description": "Archive button tooltip"
+    },
+    "HistoryGizmoItem.leaveTempChatConfirm":
+    {
+        "defaultMessage": "Leave Temporary Chat",
+        "description": "Confirm button text for confirmation modal when leaving Temporary Chat"
+    },
+    "HistoryGizmoItem.moreTooltip":
+    {
+        "defaultMessage": "More",
+        "description": "More button tooltip"
+    },
     "InstallPluginModal.cancel":
     {
         "defaultMessage": "Cancel",
@@ -1070,9 +1225,9 @@
         "defaultMessage": "OpenAI uses cookies to improve your experience and analyze site traffic. For more information, read our <cookiePolicy>cookie policy</cookiePolicy>.",
         "description": "Reasons why openai uses cookies and how they use them"
     },
-    "ManageCookiesModal.preference1desc":
-    {
-        "defaultMessage": "Cookies that are necessary to operate the service, such as logging you in and keeping your account secure.",
+    "ManageCookiesModal.preference1desc.0":
+    {
+        "defaultMessage": "These cookies are required to operate our Services. For example, they allow us to authenticate users or enable specific features within the Services, including for security purposes.",
         "description": "Essential cookies description"
     },
     "ManageCookiesModal.preference1title":
@@ -1085,9 +1240,9 @@
         "defaultMessage": "Allow essential cookies",
         "description": "Essential cookies toggle label"
     },
-    "ManageCookiesModal.preference2desc":
-    {
-        "defaultMessage": "Cookies that are necessary are used by OpenAI and our partners to help us understand how our websites are used in order to improve our services.",
+    "ManageCookiesModal.preference2desc.0":
+    {
+        "defaultMessage": "These cookies help us analyze and understand how our Services perform and are used, such as the number of users, how they interact with our Services, and time spent using the Services.",
         "description": "Analytics cookies description"
     },
     "ManageCookiesModal.preference2title":
@@ -1120,6 +1275,76 @@
         "defaultMessage": "Your cookie preferences were updated successfully",
         "description": "Success toast message for updating cookie preferences"
     },
+    "MemoriesModal.delete":
+    {
+        "defaultMessage": "Delete",
+        "description": "Delete button label"
+    },
+    "MemoriesModal.deleteFailed":
+    {
+        "defaultMessage": "Deleting failed",
+        "description": "Toast message when deleting memory fails"
+    },
+    "MemoriesModal.deleteMemory":
+    {
+        "defaultMessage": "Delete memory",
+        "description": "Label for delete memory"
+    },
+    "MemoriesModal.deleteMemoryDescription":
+    {
+        "defaultMessage": "This will delete the memory: {title}",
+        "description": "Delete memory modal description"
+    },
+    "MemoriesModal.loading":
+    {
+        "defaultMessage": "Loading...",
+        "description": "Loading label"
+    },
+    "MemoriesModal.memoryColumn":
+    {
+        "defaultMessage": "Memory",
+        "description": "Memory column name"
+    },
+    "MemoriesModal.noMemories":
+    {
+        "defaultMessage": "As you chat with ChatGPT, the details and preferences it remembers will be shown here.",
+        "description": "No memories message"
+    },
+    "MemoriesModal.resetModalCancel":
+    {
+        "defaultMessage": "Cancel",
+        "description": "Cancel button for the reset memory modal"
+    },
+    "MemoriesModal.resetModalConfirm":
+    {
+        "defaultMessage": "Clear memory",
+        "description": "Confirm button for the reset memory modal"
+    },
+    "MemoriesModal.resetModalDescription":
+    {
+        "defaultMessage": "ChatGPT will forget everything it has remembered from your chats. This cannot be undone.",
+        "description": "Description for the reset memory modal"
+    },
+    "MemoriesModal.resetModalTitle":
+    {
+        "defaultMessage": "Are you sure?",
+        "description": "Title for the reset memory modal"
+    },
+    "MemoriesModal.retry":
+    {
+        "defaultMessage": "Retry",
+        "description": "Retry button label"
+    },
+    "MemoriesModal.somethingWentWrong":
+    {
+        "defaultMessage": "Something went wrong...",
+        "description": "Error message"
+    },
+    "MemoriesModal.title":
+    {
+        "defaultMessage": "My memories",
+        "description": "Memories modal title"
+    },
     "MemoryOnboardingModal.bulletDescription1":
     {
         "defaultMessage": "Your GPT will carry what it learns between chats, allowing it to provide more relevant responses.",
@@ -1130,9 +1355,9 @@
         "defaultMessage": "As you chat your GPT will become more helpful, remembering details and preferences.",
         "description": "Description for the seconed bullet"
     },
-    "MemoryOnboardingModal.bulletDescription3":
-    {
-        "defaultMessage": "To modify what your GPT knows, just send it a message. You can reset your GPT’s memory or turn this feature off in settings.",
+    "MemoryOnboardingModal.bulletDescription3.1":
+    {
+        "defaultMessage": "Your GPT has been designed to follow your instructions in chats. You can reset your GPT's memory or turn this feature off in settings.",
         "description": "Description for the third bullet"
     },
     "MemoryOnboardingModal.bulletTitle1":
@@ -1140,9 +1365,9 @@
         "defaultMessage": "Keep the conversation going",
         "description": "Title for the first bullet"
     },
-    "MemoryOnboardingModal.bulletTitle2":
-    {
-        "defaultMessage": "Improves over time",
+    "MemoryOnboardingModal.bulletTitle2.1":
+    {
+        "defaultMessage": "More helpful over time",
         "description": "Title for the second bullet"
     },
     "MemoryOnboardingModal.bulletTitle3":
@@ -1167,6 +1392,11 @@
     },
     "ModelSwitcher.gpt4ShortExplainer":
     {
+        "defaultMessage": "With DALL\\xb7E, browsing and analysis",
+        "description": "Short description of GPT-4"
+    },
+    "ModelSwitcher.gpt4ShortExplainerWithoutBrowse":
+    {
         "defaultMessage": "With DALL\\xb7E and analysis",
         "description": "Short description of GPT-4"
     },
@@ -1175,10 +1405,20 @@
         "defaultMessage": "Our smartest and most capable model. Includes DALL\\xb7E, browsing and more.",
         "description": "Description of what new capabilities GPT4 providers"
     },
+    "ModelSwitcher.shareChat":
+    {
+        "defaultMessage": "Share chat",
+        "description": "Share chat option in model switcher"
+    },
     "ModelSwitcher.temporaryChat":
     {
-        "defaultMessage": "Temporary Chat",
+        "defaultMessage": "Temporary chat",
         "description": "Temporary chat option in model switcher"
+    },
+    "ModelSwithcer.signupUpgradeButton2":
+    {
+        "defaultMessage": "Upgrade to Plus",
+        "description": "Sign up CTA for logged out users"
     },
     "ModelSwithcer.upgradeButton":
     {
@@ -1220,6 +1460,26 @@
         "defaultMessage": "Learn more",
         "description": "Learn more link text"
     },
+    "NavigationContent.signInToChat":
+    {
+        "defaultMessage": "Sign in to ChatGPT",
+        "description": "Sign in button label"
+    },
+    "NavigationContent.signInToValueProp":
+    {
+        "defaultMessage": "Sign in to save your conversations with ChatGPT",
+        "description": "Label to save your conversation"
+    },
+    "NavigationContent.unauthLoginCta":
+    {
+        "defaultMessage": "Log in",
+        "description": "CTA for unauthenticated user to login"
+    },
+    "NavigationContent.unauthSignupCta":
+    {
+        "defaultMessage": "Sign up",
+        "description": "CTA for unauthenticated user to sign up"
+    },
     "NodeEditor.cancel":
     {
         "defaultMessage": "Cancel",
@@ -1240,6 +1500,11 @@
         "defaultMessage": "Prev",
         "description": "Label for the previous page button in the pagination controls"
     },
+    "PaymentMenuItems.createATeamWorkspace":
+    {
+        "defaultMessage": "Add Team workspace",
+        "description": "Create a team workspace menu item"
+    },
     "PaymentMenuItems.freeTrialCta":
     {
         "defaultMessage": "Get {duration}!",
@@ -1270,9 +1535,9 @@
         "defaultMessage": "Renew Plus",
         "description": "Renew Plus menu item"
     },
-    "PaymentMenuItems.upgradeToPlus.0":
-    {
-        "defaultMessage": "Upgrade",
+    "PaymentMenuItems.upgradePlan":
+    {
+        "defaultMessage": "Upgrade plan",
         "description": "Upgrade to plan menu item"
     },
     "PaymentMenuItems.upgradeToPlusUpsell.0":
@@ -1300,76 +1565,16 @@
         "defaultMessage": "You've been upgraded to ChatGPT Plus",
         "description": "Message displayed when the payment is successful"
     },
+    "PaymentsSuccess.standard":
+    {
+        "defaultMessage": "Payment successful",
+        "description": "Message displayed when the payment is successful"
+    },
     "PaymentsSuccess.trial":
     {
         "defaultMessage": "Trial started",
         "description": "Message displayed when the payment is successful"
     },
-    "Placeholder.allowsFollowUpCorrections":
-    {
-        "defaultMessage": "Allows user to provide follow-up corrections",
-        "description": "Capability description for allowing follow-up corrections"
-    },
-    "Placeholder.birthdayIdeasExample":
-    {
-        "defaultMessage": "Got any creative ideas for a 10 year old's birthday?",
-        "description": "Example text for creative birthday ideas"
-    },
-    "Placeholder.capabilitiesHeader":
-    {
-        "defaultMessage": "Capabilities",
-        "description": "Header text for the capabilities list"
-    },
-    "Placeholder.declinesInappropriateRequests":
-    {
-        "defaultMessage": "Trained to decline inappropriate requests",
-        "description": "Capability description for declining inappropriate requests"
-    },
-    "Placeholder.examplesHeader":
-    {
-        "defaultMessage": "Examples",
-        "description": "Header text for the examples list"
-    },
-    "Placeholder.httpRequestExample":
-    {
-        "defaultMessage": "How do I make an HTTP request in Javascript?",
-        "description": "Example text for making an HTTP request in Javascript"
-    },
-    "Placeholder.limitationsHeader":
-    {
-        "defaultMessage": "Limitations",
-        "description": "Header text for the limitations list"
-    },
-    "Placeholder.limitedKnowledgeAfter2021":
-    {
-        "defaultMessage": "Limited knowledge of world and events after 2021",
-        "description": "Limitation description for limited knowledge of world and events after 2021"
-    },
-    "Placeholder.mayGenerateIncorrectInfo":
-    {
-        "defaultMessage": "May occasionally generate incorrect information",
-        "description": "Limitation description for occasionally generating incorrect information"
-    },
-    "Placeholder.mayProduceHarmfulInstructions":
-    {
-        "defaultMessage": "May occasionally produce harmful instructions or biased content",
-        "description": "Limitation description for occasionally producing harmful instructions or biased content"
-    },
-    "Placeholder.quantumComputingExample":
-    {
-        "defaultMessage": "Explain quantum computing in simple terms",
-        "description": "Example text for quantum computing explanation"
-    },
-    "Placeholder.remembersEarlierConversation":
-    {
-        "defaultMessage": "Remembers what user said earlier in the conversation",
-        "description": "Capability description for remembering earlier conversation"
-    },
-    "Placeholder.title":
-    {
-        "defaultMessage": "ChatGPT",
-        "description": "Title text for the Placeholder component"
-    },
     "PluginChooser.enabledPluginsLimit":
     {
         "defaultMessage": "{enabledPlugins}/{maxEnabledPlugins} Enabled",
@@ -1495,9 +1700,9 @@
         "defaultMessage": "Agree",
         "description": "Agree/close age verification ui Banner"
     },
-    "PrivacyPolicyUpdateBanner.content.eea":
-    {
-        "defaultMessage": "We've updated our <termsLink>Terms of Use</termsLink>, effective December 14, 2023. By continuing to use our services, you agree to these updated terms. <faqLink>Learn more</faqLink>.",
+    "PrivacyPolicyUpdateBanner.content.eea.0":
+    {
+        "defaultMessage": "We've updated our <termsLink>Terms of Use</termsLink> and <privacyPolicyLink>Privacy Policy</privacyPolicyLink>. By continuing to use our services, you agree to our updated <termsLink>Terms of Use</termsLink>. If you disagree with the updated <termsLink>Terms of Use</termsLink>, you can delete your account. These updates will become effective February 15, 2024. <faqLink>Learn more</faqLink>.",
         "description": "Links to terms of use and privacy policy due to update"
     },
     "PrivacyPolicyUpdateBanner.content.exeea":
@@ -1510,16 +1715,16 @@
         "defaultMessage": "Disagree",
         "description": "Disagree/close age verification ui Banner and navigate to deactivate account"
     },
-    "PrivacyPolicyUpdateBanner.title.eea":
-    {
-        "defaultMessage": "We've updated our Terms of Use",
-        "description": "terms of use notification banner title"
-    },
-    "PrivacyPolicyUpdateBanner.title.exeea":
+    "PrivacyPolicyUpdateBanner.title.eea.0":
     {
         "defaultMessage": "We've updated our Terms of Use and Privacy Policy",
         "description": "terms of use notification banner title"
     },
+    "PrivacyPolicyUpdateBanner.title.exeea":
+    {
+        "defaultMessage": "We've updated our Terms of Use and Privacy Policy",
+        "description": "terms of use notification banner title"
+    },
     "PromptFilePicker.attachFiles":
     {
         "defaultMessage": "Attach files",
@@ -1555,11 +1760,21 @@
         "defaultMessage": "Select a response to continue",
         "description": "Text that shows input is disabled due to mandatory feedback"
     },
+    "PromptTextarea.disallowedByWorkspaceReason":
+    {
+        "defaultMessage": "Your workspace owner hasn’t approved this GPT.",
+        "description": "Reason given when a user goes to a gizmo but their workspace disallows using it."
+    },
     "PromptTextarea.errorGeneratingResponse":
     {
         "defaultMessage": "There was an error generating a response",
         "description": "Error message shown when the response generation fails"
     },
+    "PromptTextarea.gizmoKnowledgeWarning":
+    {
+        "defaultMessage": "Files uploaded here may be included in conversations with your GPT. Files can be downloaded if Code Interpreter is enabled.",
+        "description": "Warning shown when uploading files to a gizmo"
+    },
     "PromptTextarea.guidedRegenerateResponse":
     {
         "defaultMessage": "Improve",
@@ -1895,6 +2110,16 @@
         "defaultMessage": "Enable two-factor authentication",
         "description": "Title for the modal to enable two-factor authentication"
     },
+    "ResetMemoriesButton.resetFailed":
+    {
+        "defaultMessage": "Failed to reset your GPT's memory.",
+        "description": "Error message for the reset memory modal"
+    },
+    "ResetMemoriesButton.resetSuccessful":
+    {
+        "defaultMessage": "Your GPT's memory has been reset.",
+        "description": "Success message for the reset memory modal"
+    },
     "ResultsSection.error":
     {
         "defaultMessage": "Error",
@@ -1975,29 +2200,44 @@
         "defaultMessage": "Unverified",
         "description": "Text shown inside the UnapprovedTagWithText component"
     },
-    "TemporaryChatOnboardingModal.bulletDescription1":
-    {
-        "defaultMessage": "Temporary Chats won’t appear in your history, and your GPT won’t remember anything you talk about.",
+    "TaggingDropdown.dropdownCreateLabel":
+    {
+        "defaultMessage": "Create a GPT",
+        "description": "Label for no results in inline gizmo search"
+    },
+    "TaggingDropdown.searchPlaceholder":
+    {
+        "defaultMessage": "Search recent and pinned GPTs",
+        "description": "Placeholder for inline gizmo search"
+    },
+    "TaggingDropdown.talkingTo":
+    {
+        "defaultMessage": "Talking to",
+        "description": "Label prefix for inline gizmo pill"
+    },
+    "TemporaryChatOnboardingModal.bulletDescription1.1":
+    {
+        "defaultMessage": "Temporary chats won’t appear in your history. For safety purposes, we may keep a copy of your chat for up to 30 days.",
         "description": "Description for the first bullet"
     },
-    "TemporaryChatOnboardingModal.bulletDescription2":
-    {
-        "defaultMessage": "Your GPT won’t be aware of previous conversations or your custom instructions.",
+    "TemporaryChatOnboardingModal.bulletDescription2.1":
+    {
+        "defaultMessage": "ChatGPT won’t use or create memories in Temporary Chats. If you have Custom Instructions, they’ll still be followed.",
         "description": "Description for the seconed bullet"
     },
     "TemporaryChatOnboardingModal.bulletDescription3":
     {
-        "defaultMessage": "Temporary Chats won’t be used to improve our models.",
+        "defaultMessage": "Temporary Chats won't be used to improve our models.",
         "description": "Description for the third bullet"
     },
     "TemporaryChatOnboardingModal.bulletTitle1":
     {
-        "defaultMessage": "Off the record",
+        "defaultMessage": "Not in history",
         "description": "Title for the first bullet"
     },
-    "TemporaryChatOnboardingModal.bulletTitle2":
-    {
-        "defaultMessage": "Blank slate",
+    "TemporaryChatOnboardingModal.bulletTitle2.1":
+    {
+        "defaultMessage": "No memory",
         "description": "Title for the second bullet"
     },
     "TemporaryChatOnboardingModal.bulletTitle3":
@@ -2055,6 +2295,21 @@
         "defaultMessage": "Sorry, conversations created when Chat History is off expire after 6 hours of inactivity. Please start a new conversation to continue using ChatGPT.",
         "description": "History disabled conversation not found message"
     },
+    "UpgradeInviteModal.inviteValidUntilFooter":
+    {
+        "defaultMessage": "Your invite is valid until {date}",
+        "description": "Footer text notifying the user how long their invite is valid until"
+    },
+    "UpgradeInviteModal.title":
+    {
+        "defaultMessage": "You're invited to Plus",
+        "description": "Title for the Upgrade Invite modal"
+    },
+    "UpgradeInviteModal.upgradeButton":
+    {
+        "defaultMessage": "Upgrade",
+        "description": "Upgrade button text"
+    },
     "UpgradedToTeamSuccess.completeDescription":
     {
         "defaultMessage": "You have now created a team workspace. You can invite members to join the workspace. You can do so at any time from the workspace settings page.",
@@ -2265,6 +2520,11 @@
         "defaultMessage": "There was an error reactivating your subscription.",
         "description": "Error warning when trying to reactivate a subscription"
     },
+    "adminPage.gpts":
+    {
+        "defaultMessage": "GPTs",
+        "description": "GPTs sidebar item"
+    },
     "adminPage.identity.0":
     {
         "defaultMessage": "Identity & provisioning",
@@ -2360,3596 +2620,4176 @@
         "defaultMessage": "Please verify your age in the next {deadline, plural, =0 {# days} one {# day} other {# days}}. We are required to check if you're old enough to use ChatGPT. <learnMoreLink>Learn more</learnMoreLink>.",
         "description": "Explanation of why ChatGPT needs to verify age"
     },
+    "ageVerificationModal.actionButton":
+    {
+        "defaultMessage": "I meet OpenAI's age requirements",
+        "description": "Continue button on the age verification modal"
+    },
+    "ageVerificationModal.description.1":
+    {
+        "defaultMessage": "Thank you for using ChatGPT. To comply with legal requirements in Korea, we need to confirm that you’re old enough to use our service. To continue on ChatGPT, please confirm that you are 18+ or are 14+ and have your parent or guardian’s permission to use ChatGPT. <deleteAccountLink>If you do not meet our age requirements, please delete your account.</deleteAccountLink>",
+        "description": "Description of the age verification modal"
+    },
+    "ageVerificationModal.logOut":
+    {
+        "defaultMessage": "Log out",
+        "description": "Log out button on the age verification modal"
+    },
+    "ageVerificationModal.title":
+    {
+        "defaultMessage": "Please confirm you meet our age requirements",
+        "description": "Title of the age verification modal"
+    },
     "announcementTooltip.new":
     {
         "defaultMessage": "New",
         "description": "New badge text"
     },
-    "badge.enterprisePlanName":
+    "approveGPTModal.actionsTitle":
+    {
+        "defaultMessage": "Actions",
+        "description": "Actions title"
+    },
+    "approveGPTModal.approveButtonText":
+    {
+        "defaultMessage": "Approve",
+        "description": "Approve button text"
+    },
+    "approveGPTModal.approveGPTTitle":
+    {
+        "defaultMessage": "Approve a third-party GPT",
+        "description": "Title for deleting a workspace gpt modal"
+    },
+    "approveGPTModal.cancelButtonText":
+    {
+        "defaultMessage": "Cancel",
+        "description": "Cancel button text"
+    },
+    "approveGPTModal.capabilitiesTitle":
+    {
+        "defaultMessage": "Capabilities",
+        "description": "Capabilities title"
+    },
+    "approveGPTModal.gptByLine":
+    {
+        "defaultMessage": "By {authorName}",
+        "description": "By line for a gpt"
+    },
+    "approveGPTModal.noCapabilitiesEnabled":
+    {
+        "defaultMessage": "No capabilities enabled",
+        "description": "No capabilities enabled"
+    },
+    "approveGPTModal.noCustomActions":
+    {
+        "defaultMessage": "No custom actions",
+        "description": "No custom actions"
+    },
+    "archivedConversationsModal.dateCreated":
+    {
+        "defaultMessage": "Date created",
+        "description": "Table column header"
+    },
+    "archivedConversationsModal.deleteConversation":
+    {
+        "defaultMessage": "Delete conversation",
+        "description": "Label for delete shared link icon"
+    },
+    "archivedConversationsModal.deleteFailed":
+    {
+        "defaultMessage": "Deleting failed",
+        "description": "Toaster message when deleting conversation fails"
+    },
+    "archivedConversationsModal.loading":
+    {
+        "defaultMessage": "Loading...",
+        "description": "Loading message"
+    },
+    "archivedConversationsModal.name":
+    {
+        "defaultMessage": "Name",
+        "description": "Table column header"
+    },
+    "archivedConversationsModal.noArchivedConversations":
+    {
+        "defaultMessage": "You have no archived conversations.",
+        "description": "No archived conversations message"
+    },
+    "archivedConversationsModal.retry":
+    {
+        "defaultMessage": "Retry",
+        "description": "Retry button text"
+    },
+    "archivedConversationsModal.somethingWentWrong":
+    {
+        "defaultMessage": "Something went wrong...",
+        "description": "Error message"
+    },
+    "archivedConversationsModal.title":
+    {
+        "defaultMessage": "Archived Chats",
+        "description": "Archived chats modal title"
+    },
+    "archivedConversationsModal.unarchiveConversation":
+    {
+        "defaultMessage": "Unarchive conversation",
+        "description": "Label for unarchive conversation"
+    },
+    "browsingMessage.browsingFailed":
+    {
+        "defaultMessage": "Error browsing",
+        "description": "Status message when browsing failed"
+    },
+    "browsingMessage.browsingStopped":
+    {
+        "defaultMessage": "Stopped browsing",
+        "description": "Status message when browsing was stopped"
+    },
+    "browsingMessage.readingDocument":
+    {
+        "defaultMessage": "Reading {filename}",
+        "description": "Status message when reading a document"
+    },
+    "browsingMessage.retrievalFailed":
+    {
+        "defaultMessage": "Error reading documents",
+        "description": "Status message when document retrieval failed"
+    },
+    "browsingMessage.retrievalStopped":
+    {
+        "defaultMessage": "Stopped reading documents",
+        "description": "Status message when document retrieval was stopped"
+    },
+    "browsingMessage.searching":
+    {
+        "defaultMessage": "Searching Bing",
+        "description": "Status message when searching Bing"
+    },
+    "browsingMessage.searchingForQuery":
+    {
+        "defaultMessage": "Searching “{query}”",
+        "description": "Status message when searching for a query"
+    },
+    "browsingMessage.searchingKnowledge":
+    {
+        "defaultMessage": "Searching my knowledge",
+        "description": "Status message when a GPT is searching its knowledge base"
+    },
+    "browsingMessage.searchingKnowledgeFailed":
+    {
+        "defaultMessage": "Error searching knowledge",
+        "description": "Status message when a GPT failed to search its knowledge base"
+    },
+    "browsingMessage.searchingKnowledgeStopped":
+    {
+        "defaultMessage": "Stopped searching knowledge",
+        "description": "Status message when a GPT stopped searching its knowledge base"
+    },
+    "browsingMessage.startingRetrieval":
+    {
+        "defaultMessage": "Reading documents",
+        "description": "Status message when document retrieval is starting"
+    },
+    "browsingMessage.startingV3":
+    {
+        "defaultMessage": "Doing research with Bing",
+        "description": "Status message when browsing is starting"
+    },
+    "browsingMessage.visiting":
+    {
+        "defaultMessage": "Visiting {url}",
+        "description": "Status message when visiting a webpage"
+    },
+    "cancelTeamPlanModal.areYouSure":
+    {
+        "defaultMessage": "Are you sure?",
+        "description": "Title for confirmation prompt"
+    },
+    "cancelTeamPlanModal.cancelPlan":
+    {
+        "defaultMessage": "Cancel subscription",
+        "description": "Cancel your team plan modal title"
+    },
+    "cancelTeamPlanModal.cancelSubscriptionButton":
+    {
+        "defaultMessage": "Cancel subscription",
+        "description": "Label for the cancel subscription button"
+    },
+    "cancelTeamPlanModal.dataUnavailable":
+    {
+        "defaultMessage": "All workspace data including chat history and settings will not be available.",
+        "description": "Data will not be available message"
+    },
+    "cancelTeamPlanModal.doneButton":
+    {
+        "defaultMessage": "Done",
+        "description": "Label for the done button"
+    },
+    "cancelTeamPlanModal.emailPlaceholder":
+    {
+        "defaultMessage": "abcd@acme.com",
+        "description": "Placeholder text for email input"
+    },
+    "cancelTeamPlanModal.enterEmailToConfirm":
+    {
+        "defaultMessage": "Enter your Email address to confirm",
+        "description": "Prompt to enter email for confirmation"
+    },
+    "cancelTeamPlanModal.errorCancellingSubscription":
+    {
+        "defaultMessage": "There was a problem cancelling your subscription.",
+        "description": "Error cancelling subscription toast text"
+    },
+    "cancelTeamPlanModal.keepSubscriptionButton":
+    {
+        "defaultMessage": "Keep subscription",
+        "description": "Label for the keep subscription button"
+    },
+    "cancelTeamPlanModal.retainEmailAccess":
+    {
+        "defaultMessage": "You will still be able to use ChatGPT with the other workspaces associated with this email address.",
+        "description": "Retain access to other workspaces with your email"
+    },
+    "cancelTeamPlanModal.retainUntil":
+    {
+        "defaultMessage": "You will retain access to the workspace until the end of your billing cycle on {expiryDate, date, long}.",
+        "description": "License expiry date"
+    },
+    "cancelTeamPlanModal.successfullyCanceled":
+    {
+        "defaultMessage": "You have successfully canceled your subscription.",
+        "description": "Confirmation on account canceled text"
+    },
+    "changeAccessModal.accessModalPrompt":
+    {
+        "defaultMessage": "Who can access {gptName}",
+        "description": "Prompt for changing access to a gpt"
+    },
+    "changeAccessModal.accessModalTitle":
+    {
+        "defaultMessage": "Change who has access",
+        "description": "Title for changing access to a gpt"
+    },
+    "changeAccessModal.cancelButtonText":
+    {
+        "defaultMessage": "Cancel",
+        "description": "Cancel button text"
+    },
+    "changeAccessModal.linkShareOption":
+    {
+        "defaultMessage": "Anyone with the link",
+        "description": "Option for link sharing"
+    },
+    "changeAccessModal.marketplaceShareOption":
+    {
+        "defaultMessage": "Public",
+        "description": "Option for marketplace sharing"
+    },
+    "changeAccessModal.privateShareOption":
+    {
+        "defaultMessage": "Only its builder",
+        "description": "Option for private sharing"
+    },
+    "changeAccessModal.saveButtonText":
+    {
+        "defaultMessage": "Save",
+        "description": "Save button text"
+    },
+    "changeAccessModal.workspaceShareOption":
+    {
+        "defaultMessage": "Anyone at {workspaceName}",
+        "description": "Option for workspace sharing"
+    },
+    "changeOwnerModal.cancelButtonText":
+    {
+        "defaultMessage": "Cancel",
+        "description": "Cancel button text"
+    },
+    "changeOwnerModal.changeModalPrompt":
+    {
+        "defaultMessage": "New owner of {gptName}",
+        "description": "Prompt for change the owner of a gpt"
+    },
+    "changeOwnerModal.changeModalTitle":
+    {
+        "defaultMessage": "Change Owner",
+        "description": "Title for changing the owner of a workspace gpt modal"
+    },
+    "changeOwnerModal.changeOwnerButtonText":
+    {
+        "defaultMessage": "Change owner",
+        "description": "Change Owner button text"
+    },
+    "changeOwnerModal.invalidEmail":
+    {
+        "defaultMessage": "Email is invalid or does not belong to a member in this workspace",
+        "description": "Invalid email text"
+    },
+    "citations.invalid":
+    {
+        "defaultMessage": "Invalid citation",
+        "description": "Text when citation is invalid"
+    },
+    "citations.viewAnalysis":
+    {
+        "defaultMessage": "View analysis",
+        "description": "Tooltip text for a citation link to analysis"
+    },
+    "codeInterpreterMessage.errorV2":
+    {
+        "defaultMessage": "<expander>Error analyzing</expander>",
+        "description": "Status message when code interpreter ran into an error"
+    },
+    "codeInterpreterMessage.finished":
+    {
+        "defaultMessage": "<expander>Finished analyzing</expander>",
+        "description": "Status message when code interpreter is finished"
+    },
+    "codeInterpreterMessage.resultLabel":
+    {
+        "defaultMessage": "Result",
+        "description": "Label shown with the code execution result output"
+    },
+    "codeInterpreterMessage.runningV2":
+    {
+        "defaultMessage": "<expander>Analyzing</expander>",
+        "description": "Status message when code interpreter is running"
+    },
+    "codeInterpreterMessage.stoppedV2":
+    {
+        "defaultMessage": "<expander>Stopped analyzing</expander>",
+        "description": "Status message when code interpreter was stopped by the user"
+    },
+    "components.business.NumSeats.description":
+    {
+        "defaultMessage": "{num} seats in use",
+        "description": "number of seats in use description"
+    },
+    "components.business.NumSeats.dividerTooltip":
+    {
+        "defaultMessage": "Your team has {count, plural, =0 {no seats} one {# seat} other {# seats}} purchased",
+        "description": "Tooltip explaining allocated seats limit"
+    },
+    "connectorSettings.connect":
+    {
+        "defaultMessage": "Connect",
+        "description": "Label for the button to connect an app"
+    },
+    "connectorSettings.connectorsTitle":
+    {
+        "defaultMessage": "Connect apps to access their information in ChatGPT.",
+        "description": "Title row for the connected apps settings tab"
+    },
+    "connectorSettings.googleDriveDesc":
+    {
+        "defaultMessage": "Attach Google Docs, Sheets, and Slides to your messages or add them as context to your conversations.",
+        "description": "Description for the Google Drive connected app"
+    },
+    "connectorSettings.googleDriveIconAlt":
+    {
+        "defaultMessage": "Icon for Google Drive",
+        "description": "Alt text for the Google Drive icon"
+    },
+    "connectorSettings.noConnectorSettings":
+    {
+        "defaultMessage": "Unable to get connected apps",
+        "description": "Text when there are no connected apps available"
+    },
+    "connectorSettings.o365Desc":
+    {
+        "defaultMessage": "Attach Microsoft Word, Excel, and Powerpoint files to your messages or add them as context to your conversations.",
+        "description": "Description for the Microsoft 365 connected app"
+    },
+    "connectorSettings.o365IconAlt":
+    {
+        "defaultMessage": "Icon for Microsoft 365",
+        "description": "Alt text for the Microsoft 365 icon"
+    },
+    "createWorkspace.cancel":
+    {
+        "defaultMessage": "Cancel",
+        "description": "Cancel button text"
+    },
+    "createWorkspace.continueToBillingButton":
+    {
+        "defaultMessage": "Continue to billing",
+        "description": "Continue to billing button text"
+    },
+    "createWorkspace.exampleTeamWorkspaceName":
+    {
+        "defaultMessage": "Acme Inc.",
+        "description": "Example Team Workspace name"
+    },
+    "createWorkspace.paymentErrorWarning":
+    {
+        "defaultMessage": "The payments page encountered an error. Please try again. If the problem continues, please visit help.openai.com.",
+        "description": "Error toast when payment page has an error"
+    },
+    "createWorkspace.selectBillingOption":
+    {
+        "defaultMessage": "Select billing options",
+        "description": "Button text for selecting team plan"
+    },
+    "createWorkspace.selectTeamPlan":
+    {
+        "defaultMessage": "Select Team Plan",
+        "description": "Button text for selecting team plan"
+    },
+    "createWorkspace.selectTeamPlanModalTitle":
+    {
+        "defaultMessage": "Select your team plan",
+        "description": "Select team plan modal title"
+    },
+    "createWorkspace.title":
+    {
+        "defaultMessage": "Create workspace",
+        "description": "Title for the create workspace modal"
+    },
+    "createWorkspace.workspaceNameDescription":
+    {
+        "defaultMessage": "Set a workspace name for your team. The name can be changed at any time.",
+        "description": "Description text below the workspace name label"
+    },
+    "createWorkspace.workspaceNameLabel":
+    {
+        "defaultMessage": "Workspace name",
+        "description": "Label for the workspace name input field"
+    },
+    "dalleMessage.creatingImagesV2":
+    {
+        "defaultMessage": "Creating image",
+        "description": "Status message when DALL\\xb7E is creating an image"
+    },
+    "dalleMessage.errorCreatingV2":
+    {
+        "defaultMessage": "Error creating image",
+        "description": "Status message when DALL\\xb7E failed to create an image"
+    },
+    "dalleMessage.generatedImageAltText":
+    {
+        "defaultMessage": "Generated by DALL\\xb7E",
+        "description": "Alt text for images generated by DALL\\xb7E"
+    },
+    "dalleMessage.imageLoadError":
+    {
+        "defaultMessage": "Error loading image",
+        "description": "Error message when an image fails to load"
+    },
+    "dalleMessage.imageViewerMetadataCopyButton":
+    {
+        "defaultMessage": "Copy",
+        "description": "Copy button for the prompt metadata in the image viewer"
+    },
+    "dalleMessage.imageViewerMetadataCopyButtonCopied":
+    {
+        "defaultMessage": "Copied!",
+        "description": "Copy button for the prompt metadata in the image viewer when the prompt is copied"
+    },
+    "dalleMessage.imageViewerMetadataTitle":
+    {
+        "defaultMessage": "Prompt",
+        "description": "Title for the prompt metadata in the image viewer"
+    },
+    "dalleMessage.stoppedV3":
+    {
+        "defaultMessage": "Stopped creating image",
+        "description": "Status message when DALL\\xb7E was stopped by the user"
+    },
+    "deactivatedWorkspaceModal.chatHistoryUnavailable":
+    {
+        "defaultMessage": "Your chat history and settings will not be available.",
+        "description": "Description about chat history unavailability"
+    },
+    "deactivatedWorkspaceModal.createPersonalWorkspace":
+    {
+        "defaultMessage": "Create a personal workspace to continue",
+        "description": "Prompt to create a personal workspace"
+    },
+    "deactivatedWorkspaceModal.createPersonalWorkspaceButton":
+    {
+        "defaultMessage": "Create a personal workspace",
+        "description": "Button label to create a personal workspace"
+    },
+    "deactivatedWorkspaceModal.deactivatedWorkspaceReason":
+    {
+        "defaultMessage": "Because your workspace has been deactivated, you need to create a personal workspace to continue using ChatGPT.",
+        "description": "Explanation for the need to create a personal workspace"
+    },
+    "deactivatedWorkspaceModal.otherWorkspacesAvailable":
+    {
+        "defaultMessage": "You will still be able to use ChatGPT with the other workspaces associated with this email address.",
+        "description": "Description about availability of other workspaces"
+    },
+    "deactivatedWorkspaceModal.profileAlt":
+    {
+        "defaultMessage": "Profile",
+        "description": "Alt text for the profile image"
+    },
+    "deactivatedWorkspaceModal.selectWorkspace":
+    {
+        "defaultMessage": "Select a workspace to continue",
+        "description": "Prompt to select another workspace"
+    },
+    "deactivatedWorkspaceModal.workspaceDeactivated":
+    {
+        "defaultMessage": "Your workspace has been deactivated",
+        "description": "Title indicating workspace deactivation"
+    },
+    "deactivatedWorkspaceModal.workspaceDeactivatedDesc":
+    {
+        "defaultMessage": "Your workspace has been deactivated.",
+        "description": "Description of workspace deactivation"
+    },
+    "deleteGPTModal.accessModalPrompt":
+    {
+        "defaultMessage": "Are you sure you would like to delete {gptName}? This action can not be undone.",
+        "description": "Prompt for deleting a gpt"
+    },
+    "deleteGPTModal.accessModalTitle":
+    {
+        "defaultMessage": "Delete GPT",
+        "description": "Title for deleting a workspace gpt modal"
+    },
+    "deleteGPTModal.cancelButtonText":
+    {
+        "defaultMessage": "Cancel",
+        "description": "Cancel button text"
+    },
+    "deleteGPTModal.deleteButtonText":
+    {
+        "defaultMessage": "Delete",
+        "description": "Delete button text"
+    },
+    "domainModal.cancel":
+    {
+        "defaultMessage": "Cancel",
+        "description": "The label for the cancel button."
+    },
+    "domainModal.check":
+    {
+        "defaultMessage": "Check",
+        "description": "The label for the check button."
+    },
+    "domainModal.copiedTXTRecordToClipboard":
+    {
+        "defaultMessage": "Copied DNS TXT record to clipboard",
+        "description": "Message informing the user their TXT record has been copied."
+    },
+    "domainModal.domainCheckError":
+    {
+        "defaultMessage": "Your domain could not be verified: {message}",
+        "description": "Error message when domain cannot be successfully verified."
+    },
+    "domainModal.domainInput.0":
+    {
+        "defaultMessage": "Add a new domain",
+        "description": "The label for the domain input."
+    },
+    "domainModal.done":
+    {
+        "defaultMessage": "Done",
+        "description": "The label for the done button."
+    },
+    "domainModal.editTitle":
+    {
+        "defaultMessage": "Manage Domain",
+        "description": "The title for the domain modal in edit mode."
+    },
+    "domainModal.newTitle.0":
+    {
+        "defaultMessage": "Verify a new domain",
+        "description": "The title for the domain modal in new mode."
+    },
+    "domainModal.placeholder":
+    {
+        "defaultMessage": "openai.com",
+        "description": "The placeholder text domain input."
+    },
+    "domainModal.submit":
+    {
+        "defaultMessage": "Submit",
+        "description": "The label for the submit button."
+    },
+    "domainModal.successfulVerification.0":
+    {
+        "defaultMessage": "Your domain, \"{hostname}\" has been successfully verified",
+        "description": "Message showing the user their domain verification is complete.."
+    },
+    "domainModal.unverifiedDomainsTableHeader":
+    {
+        "defaultMessage": "Your unverified domains",
+        "description": "The header for the unverified domains table."
+    },
+    "domainModal.verificationTokenFooter":
+    {
+        "defaultMessage": "Then, check the record to complete the verification.",
+        "description": "Message that informs the user how to verify their domain."
+    },
+    "domainModal.verificationTokenMessage":
+    {
+        "defaultMessage": "To verify ownership of {hostname}, navigate to your DNS provider and add a TXT record with this value:",
+        "description": "Message that informs the user how to verify their domain."
+    },
+    "domainModal.verifyDomainButton":
+    {
+        "defaultMessage": "Verify",
+        "description": "The label for the verify button."
+    },
+    "emailsTextarea.clearAllEntries":
+    {
+        "defaultMessage": "Clear all",
+        "description": "Clear all entries in the list of members to be added"
+    },
+    "emailsTextarea.membersAdded":
+    {
+        "defaultMessage": "+{count} {count, plural, one {member} other {members}}",
+        "description": "Current number of members that will be added to the workspace"
+    },
+    "emailsTextarea.placeholder":
+    {
+        "defaultMessage": "Type an email and press enter...",
+        "description": "Placeholder for the insert emails textarea"
+    },
+    "emailsTextarea.removeMember":
+    {
+        "defaultMessage": "Remove {member}",
+        "description": "Remove a member from the list of members to be added"
+    },
+    "emailsTextarea.tooltipInvalidEmail":
+    {
+        "defaultMessage": "\"{email}\" may not be a valid email",
+        "description": "Tooltip for invalid email addresses"
+    },
+    "exportModal.chatDownloadLabel":
+    {
+        "defaultMessage": "The data will be sent to your registered email in a downloadable file. The download link will expire 24 hours after you receive it.",
+        "description": "Label to explain how to download chat history from the export"
+    },
+    "exportModal.chatHistoryLabel":
+    {
+        "defaultMessage": "Your chat history will be included in the export.",
+        "description": "Label to explain chat history will be in the export"
+    },
+    "exportModal.confirm":
+    {
+        "defaultMessage": "Export and delete workspace",
+        "description": "Confirm button text for exporting and then deleting account"
+    },
+    "exportModal.confirmTitle":
+    {
+        "defaultMessage": "Deleting chat history is permanent and can't be undone.",
+        "description": "Title to explain delete is permanent"
+    },
+    "exportModal.deleteLabel":
+    {
+        "defaultMessage": "Your GPTs, Plugins and Custom Instructions will be deleted.",
+        "description": "Label to explain GPTs, Plugins, and Custom Instructions will be deleted"
+    },
+    "exportModal.proceed":
+    {
+        "defaultMessage": "To proceed, click \"Export and delete workspace\" below.",
+        "description": "Prompt to proceed with the export and delete account choice"
+    },
+    "feedbackModal.continueWithChosenAnswer":
+    {
+        "defaultMessage": "The conversation will continue with the answer you choose.",
+        "description": "Information text for user during comparison"
+    },
+    "feedbackModal.copyrightContent":
+    {
+        "defaultMessage": "This content violates copyright law",
+        "description": "Label for Copyrighted Content checkbox"
+    },
+    "feedbackModal.dontLikeThis":
+    {
+        "defaultMessage": "I don't like this",
+        "description": "Label for I Don't Like This checkbox"
+    },
+    "feedbackModal.employeeConsent":
+    {
+        "defaultMessage": "Allow this content to be used for model evals",
+        "description": "Open AI employee is consenting to allow this content to be used in evals"
+    },
+    "feedbackModal.employeeConsentExplanation":
+    {
+        "defaultMessage": "Allow your feedback and conversation to be used to in model evals. Please verify there is no confidential data in the conversation.",
+        "description": "Explanation for employee consent checkbox"
+    },
+    "feedbackModal.harmfulOffensive":
+    {
+        "defaultMessage": "This content is harmful or offensive",
+        "description": "Label for harmful/offensive checkbox"
+    },
+    "feedbackModal.harmfulUnsafe":
+    {
+        "defaultMessage": "This is harmful / unsafe",
+        "description": "Label for harmful/unsafe checkbox"
+    },
+    "feedbackModal.moderationAccept":
+    {
+        "defaultMessage": "Allow Content",
+        "description": "Button text for accepting the share link and allowing it to be viewed"
+    },
+    "feedbackModal.moderationReject":
+    {
+        "defaultMessage": "Block Content",
+        "description": "Button text for rejecting the share link and blocking it from being viewed"
+    },
+    "feedbackModal.neitherAnswerBetter":
+    {
+        "defaultMessage": "Neither answer is better",
+        "description": "Button text for choosing neither answer during comparison"
+    },
+    "feedbackModal.newAnswer":
+    {
+        "defaultMessage": "New Answer",
+        "description": "Title for the new answer during comparison"
+    },
+    "feedbackModal.newAnswerBetter":
+    {
+        "defaultMessage": "New answer is better",
+        "description": "Button text for choosing new answer during comparison"
+    },
+    "feedbackModal.notHelpful":
+    {
+        "defaultMessage": "This isn't helpful",
+        "description": "Label for not helpful checkbox"
+    },
+    "feedbackModal.notTrue":
+    {
+        "defaultMessage": "This isn't true",
+        "description": "Label for not true checkbox"
+    },
+    "feedbackModal.originalAnswer":
+    {
+        "defaultMessage": "Original Answer",
+        "description": "Title for the original answer during comparison"
+    },
+    "feedbackModal.originalAnswerBetter":
+    {
+        "defaultMessage": "Original answer is better",
+        "description": "Button text for choosing original answer during comparison"
+    },
+    "feedbackModal.pickBestAnswer":
+    {
+        "defaultMessage": "Pick the best answer to improve the model",
+        "description": "Title for the compare feedback modal"
+    },
+    "feedbackModal.provideAdditionalFeedback":
+    {
+        "defaultMessage": "Provide additional feedback",
+        "description": "Title for the critique feedback modal"
+    },
+    "feedbackModal.provideReportModalTitle":
+    {
+        "defaultMessage": "Report This Content",
+        "description": "Title for the 'report' feedback modal"
+    },
+    "feedbackModal.reportContentExplanationPlaceholder":
+    {
+        "defaultMessage": "What is wrong with the response? What about this response is harmful? Please be as specific as possible, and add any details that are not present in the checkboxes below.",
+        "description": "Placeholder for textarea input when user chooses to report a shared chat"
+    },
+    "feedbackModal.reportOtherContent":
+    {
+        "defaultMessage": "I don't like this for some other reason (please describe)",
+        "description": "Label for Report Other Content checkbox"
+    },
+    "feedbackModal.sexualAbuse":
+    {
+        "defaultMessage": "This content contains sexual abuse",
+        "description": "Label for Sexual Abuse checkbox"
+    },
+    "feedbackModal.skipStep":
+    {
+        "defaultMessage": "Skip this step",
+        "description": "Button text for skipping comparison step"
+    },
+    "feedbackModal.submitFeedback":
+    {
+        "defaultMessage": "Submit feedback",
+        "description": "Button text for submitting the feedback"
+    },
+    "feedbackModal.submitReport":
+    {
+        "defaultMessage": "Submit report",
+        "description": "Button text for submitting a content-moderation report"
+    },
+    "feedbackModal.thumbsDownPlaceholder":
+    {
+        "defaultMessage": "What was the issue with the response? How could it be improved?",
+        "description": "Placeholder for textarea input when user chooses thumbs down"
+    },
+    "feedbackModal.thumbsUpPlaceholder":
+    {
+        "defaultMessage": "What do you like about the response?",
+        "description": "Placeholder for textarea input when user chooses thumbs up"
+    },
+    "fileUpload.codeInterpreterSessionTimeout":
+    {
+        "defaultMessage": "Code interpreter session expired",
+        "description": "Error message when code interpreter session expired"
+    },
+    "fileUpload.defaultCreateEntryError":
+    {
+        "defaultMessage": "Unable to upload {fileName}",
+        "description": "Error message when file upload fails"
+    },
+    "fileUpload.defaultDownloadLinkError":
+    {
+        "defaultMessage": "Failed to get upload status for {fileName}",
+        "description": "Error message when file download link fails"
+    },
+    "fileUpload.fileCorrupted":
+    {
+        "defaultMessage": "This file is corrupted. Please ensure the file is not corrupted and try again.",
+        "description": "Error message when an uploaded file that is corrupted."
+    },
+    "fileUpload.fileEmpty":
+    {
+        "defaultMessage": "No text could be extracted from this file.",
+        "description": "Error message when an uploaded file does not contain parsable text"
+    },
+    "fileUpload.fileEncrypted":
+    {
+        "defaultMessage": "This file is encrypted/requires a password to access. Please try again with an unencrypted file.",
+        "description": "Error message when an uploaded file that is encrypted."
+    },
+    "fileUpload.fileNotFound":
+    {
+        "defaultMessage": "File not found",
+        "description": "Error message when file was not found"
+    },
+    "fileUpload.fileTimedOut":
+    {
+        "defaultMessage": "File upload timed out. Please try again.",
+        "description": "Error message when file upload timed out"
+    },
+    "fileUpload.fileTooLarge":
+    {
+        "defaultMessage": "File is too large",
+        "description": "Error message when file is too large to upload"
+    },
+    "fileUpload.fileTooManyTokens":
+    {
+        "defaultMessage": "This file contains too much text content. Please try again with a smaller file.",
+        "description": "Error message when an uploaded file contains too many tokens/is too large."
+    },
+    "fileUpload.fileZeroBytes":
+    {
+        "defaultMessage": "File is empty",
+        "description": "Error message when file is zero bytes"
+    },
+    "fileUpload.overUserQuota":
+    {
+        "defaultMessage": "User quota exceeded",
+        "description": "Error message when user storage space (quote) has been exceeded"
+    },
+    "fileUpload.permissionError":
+    {
+        "defaultMessage": "Missing permission to access file",
+        "description": "Error message when user doesn't have permission to access a file"
+    },
+    "fileUpload.unknownError":
+    {
+        "defaultMessage": "Unknown error occurred",
+        "description": "Error message when file upload fails"
+    },
+    "filesModal.fileDownloadFailed":
+    {
+        "defaultMessage": "File download failed. Please try again.",
+        "description": "Error message when file download fails"
+    },
+    "gizmo.actionNeedsPrivacyPolicyURL":
+    {
+        "defaultMessage": "Public actions require valid privacy policy URLs. Click <fixlink>here</fixlink> to update.",
+        "description": "Error message when trying to publish action"
+    },
+    "gizmo.actions.actionsGptHelp":
+    {
+        "defaultMessage": "Get help from ActionsGPT",
+        "description": "button text for external link to ActionsGPT to assist generating schema"
+    },
+    "gizmo.actions.blankExampleTitle":
+    {
+        "defaultMessage": "Blank Template",
+        "description": "Dropdown label for the blank template OpenAPI spec"
+    },
+    "gizmo.actions.examples":
+    {
+        "defaultMessage": "Examples",
+        "description": "Label of examples in GPT actions editor"
+    },
+    "gizmo.actions.petStoreExampleTitle":
+    {
+        "defaultMessage": "Pet Store (YAML)",
+        "description": "Dropdown label for the pet store OpenAPI example"
+    },
+    "gizmo.actions.weatherExampleTitle":
+    {
+        "defaultMessage": "Weather (JSON)",
+        "description": "Dropdown label for the weather OpenAPI example"
+    },
+    "gizmo.anyoneWithLink":
+    {
+        "defaultMessage": "Anyone with a link",
+        "description": "Privacy option for anyone with a link"
+    },
+    "gizmo.appealButton":
+    {
+        "defaultMessage": "Appeal",
+        "description": "Button label for appealing a moderation decision"
+    },
+    "gizmo.appealNeededLabel":
+    {
+        "defaultMessage": "Because this GPT previously may have violated our policies, you will have to submit an appeal to make it available at this level.",
+        "description": "Label for appeal needed message"
+    },
+    "gizmo.categoryError":
+    {
+        "defaultMessage": "Error generating category",
+        "description": "Error message when category generation fails"
+    },
+    "gizmo.categoryLabel":
+    {
+        "defaultMessage": "Category",
+        "description": "Label for category"
+    },
+    "gizmo.categoryLabelTooltip":
+    {
+        "defaultMessage": "Your GPT may appear in this category on the Explore page",
+        "description": "Tooltip for category label"
+    },
+    "gizmo.clearChat":
+    {
+        "defaultMessage": "Clear chat",
+        "description": "Clear chat button label"
+    },
+    "gizmo.confirmPublish":
+    {
+        "defaultMessage": "Confirm",
+        "description": "Message prompting you to confirm publication of an action"
+    },
+    "gizmo.copyLink":
+    {
+        "defaultMessage": "Copy link",
+        "description": "Menu item for copying link to GPT"
+    },
+    "gizmo.createActionLabel":
+    {
+        "defaultMessage": "Create new action",
+        "description": "Label for button to create a new action"
+    },
+    "gizmo.delete":
+    {
+        "defaultMessage": "Delete GPT",
+        "description": "Button label for deleting a GPT"
+    },
+    "gizmo.descriptionTooLong":
+    {
+        "defaultMessage": "GPT descriptions cannot be longer than {length} characters.",
+        "description": "Error message when description is too long"
+    },
+    "gizmo.disabledCustomActionsTooltip":
+    {
+        "defaultMessage": "Custom actions are disabled for your workspace. Contact your admin to enable them.",
+        "description": "Tooltip label when custom actions are are disabled"
+    },
+    "gizmo.discovery.browsingAsOwner":
+    {
+        "defaultMessage": "You're viewing as a <b>workspace member</b>",
+        "description": "A notice informing the user what workspace role they're using"
+    },
+    "gizmo.discovery.browsingAsOwnerTooltip":
+    {
+        "defaultMessage": "Owners can access all third-party GPTs, even if they haven’t been approved.",
+        "description": "Tooltip explaining that users with the owner role can use all third-party GPTs"
+    },
+    "gizmo.discovery.browsingMemberRole":
+    {
+        "defaultMessage": "workspace member",
+        "description": "The standard member role of the workspace"
+    },
+    "gizmo.discovery.browsingOwnerRole":
+    {
+        "defaultMessage": "workspace owner",
+        "description": "The user's role as an owner of their workspace/account"
+    },
+    "gizmo.discovery.createGPT":
+    {
+        "defaultMessage": "Create",
+        "description": "Label for create GPT button"
+    },
+    "gizmo.discovery.createdAgoLabel":
+    {
+        "defaultMessage": "Created {createdAgo}",
+        "description": "Label for created ago in search result"
+    },
+    "gizmo.discovery.empty":
+    {
+        "defaultMessage": "Nothing to discover",
+        "description": "Label for empty discovery page"
+    },
+    "gizmo.discovery.error":
+    {
+        "defaultMessage": "Error loading GPTs",
+        "description": "Label for error loading GPTs"
+    },
+    "gizmo.discovery.globalViewOff":
+    {
+        "defaultMessage": "Global View Off",
+        "description": "Label for global view off button"
+    },
+    "gizmo.discovery.globalViewOn":
+    {
+        "defaultMessage": "Global View On",
+        "description": "Label for global view on button"
+    },
+    "gizmo.discovery.loadMore":
+    {
+        "defaultMessage": "See more",
+        "description": "Button label for loading more GPTs"
+    },
+    "gizmo.discovery.loadMoreError":
+    {
+        "defaultMessage": "Error loading more",
+        "description": "Button label for error while loading more GPTs"
+    },
+    "gizmo.discovery.myGPTs":
+    {
+        "defaultMessage": "My GPTs",
+        "description": "Label for my GPTs button"
+    },
+    "gizmo.discovery.search":
+    {
+        "defaultMessage": "Search public GPTs",
+        "description": "Placeholder for search input"
+    },
+    "gizmo.discovery.search.empty":
+    {
+        "defaultMessage": "No results found",
+        "description": "Label for no search results"
+    },
+    "gizmo.discovery.search.recentlyUsed":
+    {
+        "defaultMessage": "Recently Used",
+        "description": "Label for recently used section in search"
+    },
+    "gizmo.discovery.search.recentlyUsedEmpty":
+    {
+        "defaultMessage": "No recently used GPTs",
+        "description": "Label for recently used section in search"
+    },
+    "gizmo.discovery.search.results":
+    {
+        "defaultMessage": "Search Results",
+        "description": "Label for search results section in search"
+    },
+    "gizmo.discovery.searchSmall":
+    {
+        "defaultMessage": "Search",
+        "description": "Placeholder for search input, when screen is small"
+    },
+    "gizmo.discovery.splashDescription":
+    {
+        "defaultMessage": "Discover and create custom versions of ChatGPT that combine instructions, extra knowledge, and any combination of skills.",
+        "description": "Description for the discovery page"
+    },
+    "gizmo.discovery.splashHeader":
+    {
+        "defaultMessage": "GPTs",
+        "description": "Header for the discovery page"
+    },
+    "gizmo.discovery.thirdPartyGPTsDisabled":
+    {
+        "defaultMessage": "Your admin has blocked GPTs created outside {workspaceName}.",
+        "description": "Description for third party GPTs disabled"
+    },
+    "gizmo.displayNameRequiredHint":
+    {
+        "defaultMessage": "To publish your GPTs to the Explore GPTs page, set up your builder profile.",
+        "description": "Message hinting that you have to setup a builder profile to publish"
+    },
+    "gizmo.draft":
+    {
+        "defaultMessage": "Draft",
+        "description": "Label for draft status"
+    },
+    "gizmo.errorSavingDraft":
+    {
+        "defaultMessage": "Error saving draft",
+        "description": "Error message when saving a draft fails"
+    },
+    "gizmo.explore":
+    {
+        "defaultMessage": "Explore",
+        "description": "Button that allows you to explore more GPTs"
+    },
+    "gizmo.exploreStoreEnabled":
+    {
+        "defaultMessage": "Explore GPTs",
+        "description": "Button that allows you to explore more GPTs"
+    },
+    "gizmo.gpt.grid.conversationCountStrLabel":
+    {
+        "defaultMessage": "{numConvos} Chats",
+        "description": "Label for the number of conversations a GPT has"
+    },
+    "gizmo.hideFromSidebar":
+    {
+        "defaultMessage": "Hide from sidebar",
+        "description": "Whether to hide a gpt from sidebar"
+    },
+    "gizmo.instructionsTooLong":
+    {
+        "defaultMessage": "GPT instructions cannot be longer than {length} characters.",
+        "description": "Error message when instructions are too long"
+    },
+    "gizmo.keepInSidebar":
+    {
+        "defaultMessage": "Keep in sidebar",
+        "description": "Whether to keep a gpt in sidebar"
+    },
+    "gizmo.knowledgeExplanation":
+    {
+        "defaultMessage": "Additional files for this GPT to reference.",
+        "description": "Explainer text around what happens when your files upload."
+    },
+    "gizmo.knowledgeWarning":
+    {
+        "defaultMessage": "If you upload files under Knowledge, conversations with your GPT may include file contents. Files can be downloaded when Code Interpreter is enabled",
+        "description": "Explanation text for what happens when files are uploaded"
+    },
+    "gizmo.maxActionsReached":
+    {
+        "defaultMessage": "GPTs can have a maximum of {number} actions",
+        "description": "Message when maximum number of actions has been reached"
+    },
+    "gizmo.mygpts.loadMore":
+    {
+        "defaultMessage": "Load more",
+        "description": "Button label for loading more GPTs"
+    },
+    "gizmo.nameTooLong":
+    {
+        "defaultMessage": "GPT names cannot be longer than {length} characters.",
+        "description": "Error message when name is too long"
+    },
+    "gizmo.newChat":
+    {
+        "defaultMessage": "New chat",
+        "description": "New chat tooltip"
+    },
+    "gizmo.newGPT":
+    {
+        "defaultMessage": "New GPT",
+        "description": "Placeholder for new GPT name"
+    },
+    "gizmo.onlyMe":
+    {
+        "defaultMessage": "Only me",
+        "description": "Privacy option for only me"
+    },
+    "gizmo.privacyAnyoneWithLink":
+    {
+        "defaultMessage": "Anyone with a link",
+        "description": "Description for a privacy setting of anyone with link"
+    },
+    "gizmo.privacyMarketplace":
+    {
+        "defaultMessage": "Everyone",
+        "description": "Description for a privacy setting of public"
+    },
+    "gizmo.privacyOnlyMe":
+    {
+        "defaultMessage": "Only me",
+        "description": "Description for a privacy setting of fully private"
+    },
+    "gizmo.privacyWorkspace":
+    {
+        "defaultMessage": "Anyone at {workspaceName}",
+        "description": "Description for a privacy setting of workspace"
+    },
+    "gizmo.public":
+    {
+        "defaultMessage": "Everyone",
+        "description": "Privacy option for public"
+    },
+    "gizmo.publicSharingHint":
+    {
+        "defaultMessage": "This GPT may appear in the GPT Store (coming soon)",
+        "description": "Message hinting at what public privacy settings do"
+    },
+    "gizmo.publicSharingHintStoreEnabled":
+    {
+        "defaultMessage": "Your GPT will appear in the Explore GPTs page",
+        "description": "Message hinting at what public privacy settings do"
+    },
+    "gizmo.publishChanges":
+    {
+        "defaultMessage": "Update",
+        "description": "Button label for publishing changes"
+    },
+    "gizmo.publishTo":
+    {
+        "defaultMessage": "Publish to",
+        "description": "Label for above publish options"
+    },
+    "gizmo.published":
+    {
+        "defaultMessage": "Published",
+        "description": "Label for published status"
+    },
+    "gizmo.publishedTitle":
+    {
+        "defaultMessage": "Published!",
+        "description": "Popout title when GPT is published"
+    },
+    "gizmo.recipientBlocked":
+    {
+        "defaultMessage": "Because this GPT previously may have violated our policies, you cannot publish it at this level.",
+        "description": "Label explaining that GPT cannot be shared at the selected level"
+    },
+    "gizmo.revertConfirm":
+    {
+        "defaultMessage": "Are you sure you want to revert to the last saved version?",
+        "description": "Confirmation message for reverting to last saved version"
+    },
+    "gizmo.revertMenuItem":
+    {
+        "defaultMessage": "Revert...",
+        "description": "Menu item for reverting to last saved version"
+    },
+    "gizmo.save":
+    {
+        "defaultMessage": "Save",
+        "description": "Button label for save (publishing)"
+    },
+    "gizmo.viewGPT":
+    {
+        "defaultMessage": "View GPT",
+        "description": "Button label for viewing GPT"
+    },
+    "gizmo.welcomeMessageTooLong":
+    {
+        "defaultMessage": "GPT welcome messages cannot be longer than {length} characters.",
+        "description": "Error message when welcome message is too long"
+    },
+    "gizmo.workspaceDisabledHint":
+    {
+        "defaultMessage": "Your workspace administrator has disabled this setting",
+        "description": "Message saying your workspace administrator isn't allowing something"
+    },
+    "globalToasts.conversationInaccessible":
+    {
+        "defaultMessage": "Conversation inaccessible or not found",
+        "description": "Warning toast message when conversation cannot load"
+    },
+    "globalToasts.gizmoNotFound":
+    {
+        "defaultMessage": "GPT inaccessible or not found",
+        "description": "Warning toast message when GPT cannot load"
+    },
+    "globalToasts.noAccess":
+    {
+        "defaultMessage": "You do not currently have access to this feature",
+        "description": "Warning when you try to access a feature you don't have access to"
+    },
+    "globalToasts.oauthSuccess":
+    {
+        "defaultMessage": "You have successfully signed in via OAuth",
+        "description": "Success message when a user has logged into an app with oauth."
+    },
+    "gptExploreGPTs.description":
+    {
+        "defaultMessage": "Now you can discover GPTs created by the community",
+        "description": "Description of modal announcing Explore GPTs"
+    },
+    "gptExploreGPTs.title":
+    {
+        "defaultMessage": "Explore GPTs",
+        "description": "Title of modal announcing Explore GPTs"
+    },
+    "history.bucket.lastSeven":
+    {
+        "defaultMessage": "Previous 7 Days",
+        "description": "Label for the history bucket of the previous 7 days"
+    },
+    "history.bucket.lastThirty":
+    {
+        "defaultMessage": "Previous 30 Days",
+        "description": "Label for the history bucket of the previous 30 days"
+    },
+    "history.bucket.today":
+    {
+        "defaultMessage": "Today",
+        "description": "Label for today's history bucket"
+    },
+    "history.bucket.yesterday":
+    {
+        "defaultMessage": "Yesterday",
+        "description": "Label for yesterday's history bucket"
+    },
+    "history.deleteModalBody":
+    {
+        "defaultMessage": "This will delete {title}.",
+        "description": "Body of the modal to confirm deleting a conversation"
+    },
+    "history.deleteModalCancel":
+    {
+        "defaultMessage": "Cancel",
+        "description": "Button to cancel deleting a conversation"
+    },
+    "history.deleteModalConfirm":
+    {
+        "defaultMessage": "Delete",
+        "description": "Button to confirm deleting a conversation"
+    },
+    "history.deleteModalTitle":
+    {
+        "defaultMessage": "Delete chat?",
+        "description": "Title of the modal to confirm deleting a conversation"
+    },
+    "history.retryButton":
+    {
+        "defaultMessage": "Retry",
+        "description": "Button to retry loading history"
+    },
+    "history.showMoreButton":
+    {
+        "defaultMessage": "Show more",
+        "description": "Button to show more history items"
+    },
+    "history.unableToLoad":
+    {
+        "defaultMessage": "Unable to load history",
+        "description": "Error message when history fails to load"
+    },
+    "imageViewer.closeModal":
+    {
+        "defaultMessage": "Close Modal",
+        "description": "Button to close the modal"
+    },
+    "imageViewer.downloadImage":
+    {
+        "defaultMessage": "Download Image",
+        "description": "Button to download the image"
+    },
+    "imageViewer.nextImage":
+    {
+        "defaultMessage": "Next Image",
+        "description": "Button to go to the next image"
+    },
+    "imageViewer.previousImage":
+    {
+        "defaultMessage": "Previous Image",
+        "description": "Button to go to the previous image"
+    },
+    "imageViewer.showImage":
+    {
+        "defaultMessage": "Show Image",
+        "description": "Button to show the image in a modal"
+    },
+    "imageViewer.toggleSidebar":
+    {
+        "defaultMessage": "Toggle Sidebar",
+        "description": "Button to toggle the sidebar"
+    },
+    "initialModal.billing":
+    {
+        "defaultMessage": "If you have a paid plan, it will be canceled upon completing this step.",
+        "description": "Description for what happens to billing"
+    },
+    "initialModal.confirm":
+    {
+        "defaultMessage": "Continue",
+        "description": "Confirm button text to continue to next step"
+    },
+    "initialModal.dataDelete":
+    {
+        "defaultMessage": "All Plugins and Custom Instructions in the existing workspace will be deleted.",
+        "description": "Description for what will be deleted"
+    },
+    "initialModal.explanation":
+    {
+        "defaultMessage": "What should we do with your existing workspace?",
+        "description": "Subtitle for initial modal for account transfer"
+    },
+    "initialModal.exportDetail":
+    {
+        "defaultMessage": "Delete the existing workspace",
+        "description": "Description for export and delete workspace option"
+    },
+    "initialModal.exportLabel":
+    {
+        "defaultMessage": "Export and delete existing chat history",
+        "description": "Label for export and delete chat history"
+    },
+    "initialModal.mustChoose":
+    {
+        "defaultMessage": "In order to join this new workspace, you must choose to either migrate or export your existing chat history.",
+        "description": "Description for why must migrate or export"
+    },
+    "initialModal.title":
+    {
+        "defaultMessage": "You've been added to a ChatGPT Enterprise workspace",
+        "description": "Initial modal title for account transfer"
+    },
+    "initialModal.transferDetail":
+    {
+        "defaultMessage": "Merge to a new Enterprise workspace",
+        "description": "Description for transfer to team option"
+    },
+    "initialModal.transferLabel":
+    {
+        "defaultMessage": "Transfer existing chat history and GPTs",
+        "description": "Label for transfer chat history to workspace"
+    },
+    "jitPluginMessage.allow":
+    {
+        "defaultMessage": "Allow",
+        "description": "Button text for the user to allow a custom action"
+    },
+    "jitPluginMessage.alwaysAllow":
+    {
+        "defaultMessage": "Always Allow",
+        "description": "Button text for the user to always allow an action domain"
+    },
+    "jitPluginMessage.confirm":
+    {
+        "defaultMessage": "Confirm",
+        "description": "Button text for the user to confirm a consequential custom action"
+    },
+    "jitPluginMessage.confirmParamsTitleV2":
+    {
+        "defaultMessage": "{gizmoName} needs to send this info to {domain}",
+        "description": "Title describing data that will be sent to the external website"
+    },
+    "jitPluginMessage.confirmingV3":
+    {
+        "defaultMessage": "<params><title>{gizmoName} wants to talk to {domain}</title><subtitle>Only allow sites you trust</subtitle></params>",
+        "description": "Status message when a custom action is showing a user confirmation. Has a separate title and subtitle."
+    },
+    "jitPluginMessage.decline":
+    {
+        "defaultMessage": "Decline",
+        "description": "Button text for the user to decline a custom action"
+    },
+    "jitPluginMessage.declined":
+    {
+        "defaultMessage": "You declined this action",
+        "description": "Status message when a custom action was declined by the user"
+    },
+    "jitPluginMessage.deny":
+    {
+        "defaultMessage": "Deny",
+        "description": "Button text for the user to deny a consequential custom action"
+    },
+    "jitPluginMessage.errorV5":
+    {
+        "defaultMessage": "Error talking to {domain}",
+        "description": "Status message when a custom action ran into an error"
+    },
+    "jitPluginMessage.finishedV3":
+    {
+        "defaultMessage": "<params>Talked to {domain}</params>",
+        "description": "Status message when a custom action is finished"
+    },
+    "jitPluginMessage.privacyPolicyLinkV2":
+    {
+        "defaultMessage": "Privacy policy",
+        "description": "Text for the privacy policy link"
+    },
+    "jitPluginMessage.ranTest":
+    {
+        "defaultMessage": "Tested {operationName}",
+        "description": "Status message when the user launched a test action"
+    },
+    "jitPluginMessage.runningV4":
+    {
+        "defaultMessage": "Talking to {domain}",
+        "description": "Status message when a custom action is running"
+    },
+    "jitPluginMessage.sentParamsTitleV2":
+    {
+        "defaultMessage": "{gizmoName} sent this info to {domain}",
+        "description": "Title describing data that was sent to the external website"
+    },
+    "jitPluginMessage.signInButton":
+    {
+        "defaultMessage": "Sign in with {domain}",
+        "description": "Button text for the user to sign in with an external website"
+    },
+    "jitPluginMessage.starting":
+    {
+        "defaultMessage": "Starting action",
+        "description": "Status message when a custom action is starting"
+    },
+    "jitPluginMessage.stoppedV4":
+    {
+        "defaultMessage": "Stopped talking to {domain}",
+        "description": "Status message when a custom action was stopped by the user"
+    },
+    "keyboardActions.copyLastCodeBlock":
+    {
+        "defaultMessage": "Copy last code block",
+        "description": "Keyboard shortcut to copy the last code block in the chat"
+    },
+    "keyboardActions.copyLastResponse":
+    {
+        "defaultMessage": "Copy last response",
+        "description": "Keyboard shortcut to copy the last response in the chat"
+    },
+    "keyboardActions.deleteChat":
+    {
+        "defaultMessage": "Delete chat",
+        "description": "Keyboard shortcut to delete chat"
+    },
+    "keyboardActions.focusPromptTextarea":
+    {
+        "defaultMessage": "Focus chat input",
+        "description": "Keyboard shortcut to focus the chat input"
+    },
+    "keyboardActions.navigationToggle":
+    {
+        "defaultMessage": "Toggle sidebar",
+        "description": "Keyboard shortcut to toggle navigation"
+    },
+    "keyboardActions.newChat":
+    {
+        "defaultMessage": "Open new chat",
+        "description": "Keyboard shortcut to open a new chat"
+    },
+    "keyboardActions.toggleCustomInstructions":
+    {
+        "defaultMessage": "Set custom instructions",
+        "description": "Keyboard shortcut to toggle custom instructions"
+    },
+    "keyboardActions.toggleKeyboardActions":
+    {
+        "defaultMessage": "Show shortcuts",
+        "description": "Keyboard shortcut to toggle keyboard actions"
+    },
+    "leaveWorkspaceModal.cancel":
+    {
+        "defaultMessage": "Cancel",
+        "description": "Cancel button to close leave-workspace dialog"
+    },
+    "leaveWorkspaceModal.cantLeaveWorkspace":
+    {
+        "defaultMessage": "Couldn't leave the {workspaceName} workspace",
+        "description": "User tried to leave workspace but they're the last owner"
+    },
+    "leaveWorkspaceModal.done":
+    {
+        "defaultMessage": "Done",
+        "description": "Done button to close leave-workspace dialog"
+    },
+    "leaveWorkspaceModal.enterYourEmail":
+    {
+        "defaultMessage": "Enter your email address to confirm",
+        "description": "Enter your email"
+    },
+    "leaveWorkspaceModal.lastOwnerWarning":
+    {
+        "defaultMessage": "Because you're the only owner in the {workspaceName} workspace, assign the owner role to another member before leaving.",
+        "description": "Explaining why the user can't leave the workspace"
+    },
+    "leaveWorkspaceModal.leaveAreYouSure":
+    {
+        "defaultMessage": "Are you sure?",
+        "description": "Help text section title for leave workspace dialog"
+    },
+    "leaveWorkspaceModal.leaveButton":
+    {
+        "defaultMessage": "Leave workspace",
+        "description": "Button to confirm leaving the workspace"
+    },
+    "leaveWorkspaceModal.leaveFailed":
+    {
+        "defaultMessage": "Failed to leave workspace",
+        "description": "Error message when user fails to leave workspace"
+    },
+    "leaveWorkspaceModal.leaveWorkspace":
+    {
+        "defaultMessage": "Leave the {workspaceName} workspace",
+        "description": "Leave workspace dialog title, containing name of workspace that user will leave."
+    },
+    "leaveWorkspaceModal.leaveWorkspaceWarning1":
+    {
+        "defaultMessage": "This will remove you from your workspace and you won't be able to access all data, including profile, settings, and chat history.",
+        "description": "Help text for leaving workspace, first bullet point."
+    },
+    "leaveWorkspaceModal.leaveWorkspaceWarning2":
+    {
+        "defaultMessage": "You will lose access to all channels and messages in this workspace.",
+        "description": "Help text for leaving workspace, second bullet point."
+    },
+    "leaveWorkspaceModal.leftWorkspaceDescription":
+    {
+        "defaultMessage": "You have successfully left the {workspaceName} workspace.",
+        "description": "Description of the leave-workspace dialog after the user has left the workspace"
+    },
+    "leaveWorkspaceModal.leftWorkspaceDescriptionNoOtherWorkspaces":
+    {
+        "defaultMessage": "You have successfully left the {workspaceName} workspace. This will create your personal workspace automatically.",
+        "description": "Description of the leave-workspace dialog after the user has left the workspace"
+    },
+    "leaveWorkspaceModal.leftWorkspaceTitle":
+    {
+        "defaultMessage": "Successfully left the {workspaceName} workspace",
+        "description": "Title of the leave-workspace dialog after the user has left the workspace"
+    },
+    "leaveWorkspaceModal.memberCount":
+    {
+        "defaultMessage": "{memberCount, plural, one {1 member} other {{memberCount} members} }",
+        "description": "Number of members in workspace"
+    },
+    "leaveWorkspaceModal.ok":
+    {
+        "defaultMessage": "OK",
+        "description": "OK button to close leave-workspace dialog"
+    },
+    "leaveWorkspaceModal.startPersonalAccount":
+    {
+        "defaultMessage": "Start using ChatGPT for free",
+        "description": "Button to let the user know that they can start using a free personal account."
+    },
+    "mergerModal.confirm":
+    {
+        "defaultMessage": "Confirm merge",
+        "description": "Confirm button text for merging account"
+    },
+    "message.gizmo.failed":
+    {
+        "defaultMessage": "Failed to generate profile picture.",
+        "description": "Message displayed when the GPT editor failed to generate a profile picture"
+    },
+    "message.gizmo.generatingProfilePic":
+    {
+        "defaultMessage": "Generating profile picture...",
+        "description": "Message displayed when the GPT editor is generating a profile picture"
+    },
+    "message.gizmo.updating":
+    {
+        "defaultMessage": "Updating GPT...",
+        "description": "Message displayed when the GPT editor is updating"
+    },
+    "modelCapMessaging.shortLimitDays":
+    {
+        "defaultMessage": "Limit {numerator, plural, =0 {# message} one {# message} other {# messages}} / {denominator, plural, =0 {# day} one {# day} other {# days}}",
+        "description": "Short message limit"
+    },
+    "modelCapMessaging.shortLimitHours":
+    {
+        "defaultMessage": "Limit {numerator, plural, =0 {# message} one {# message} other {# messages}} / {denominator, plural, =0 {# hour} one {# hour} other {# hours}}",
+        "description": "Short message limit"
+    },
+    "modelCapMessaging.shortLimitMinutes":
+    {
+        "defaultMessage": "Limit {numerator, plural, =0 {# message} one {# message} other {# messages}} / {denominator, plural, =0 {# minute} one {# minute} other {# minutes}}",
+        "description": "Short message limit"
+    },
+    "navigation.accountSwitcherTitle":
+    {
+        "defaultMessage": "Workspaces",
+        "description": "Account switcher title"
+    },
+    "navigation.addWorkspaceTooltip":
+    {
+        "defaultMessage": "Create a Team workspace",
+        "description": "Tooltip for add workspace button"
+    },
+    "navigation.closeSidebar":
+    {
+        "defaultMessage": "Close sidebar",
+        "description": "Close sidebar button label"
+    },
+    "navigation.disabledWorkspaceTooltip":
+    {
+        "defaultMessage": "This workspace has been deactivated",
+        "description": "Tooltip for disabled workspace"
+    },
+    "navigation.helpAndFaq":
+    {
+        "defaultMessage": "Help & FAQ",
+        "description": "Help & FAQ menu item"
+    },
+    "navigation.leaveWorkspace":
+    {
+        "defaultMessage": "Leave workspace",
+        "description": "Menu item in workplace switcher to leave a deactivated workspace"
+    },
+    "navigation.logOut":
+    {
+        "defaultMessage": "Log out",
+        "description": "Log out menu item"
+    },
+    "navigation.openSidebar":
+    {
+        "defaultMessage": "Open sidebar",
+        "description": "Open sidebar button label"
+    },
+    "navigation.reactivateWorkspace":
+    {
+        "defaultMessage": "Reactivate workspace",
+        "description": "Menu item in workplace switcher to reactivate a deactivated workspace"
+    },
+    "navigation.settings":
+    {
+        "defaultMessage": "Settings",
+        "description": "Settings menu item"
+    },
+    "navigation.settingsPlus":
+    {
+        "defaultMessage": "Settings & Beta",
+        "description": "Settings menu item for Plus users"
+    },
+    "navigation.survey.takeSurveyButton":
+    {
+        "defaultMessage": "Take survey",
+        "description": "Survey offer call to action"
+    },
+    "navigation.surveyDescription":
+    {
+        "defaultMessage": "Shape the future of ChatGPT.",
+        "description": "Survey offer description"
+    },
+    "navigation.surveyDismiss":
+    {
+        "defaultMessage": "Dismiss survey",
+        "description": "Survey offer dismiss button"
+    },
+    "navigation.surveyTitle":
+    {
+        "defaultMessage": "We’d love to hear from you!",
+        "description": "Survey offer title"
+    },
+    "onboarding.accuracy":
+    {
+        "defaultMessage": "Check your facts",
+        "description": "Title for the warning about ChatGPT inaccuracy"
+    },
+    "onboarding.accuracyBody":
+    {
+        "defaultMessage": "While we have safeguards, ChatGPT may give you inaccurate information. It’s not intended to give advice.",
+        "description": "Body copy for the warning about ChatGPT inaccuracy"
+    },
+    "onboarding.askAway":
+    {
+        "defaultMessage": "Ask away",
+        "description": "Title for the tip about what ChatGPT can do"
+    },
+    "onboarding.askAwayBody":
+    {
+        "defaultMessage": "ChatGPT can answer questions, help you learn, write code, brainstorm together, and much more.",
+        "description": "Body copy for the tip about what ChatGPT can do"
+    },
+    "onboarding.chatgptTitle.0":
+    {
+        "defaultMessage": "Welcome to the {workspaceName} workspace",
+        "description": "Title for the initial onboarding modal"
+    },
+    "onboarding.continueButton":
+    {
+        "defaultMessage": "Continue",
+        "description": "Label for the continue button"
+    },
+    "onboarding.departments.administrative":
+    {
+        "defaultMessage": "Administrative Assistant",
+        "description": "Department option for Administrative Assistant"
+    },
+    "onboarding.departments.analytics":
+    {
+        "defaultMessage": "Data or Analytics",
+        "description": "Department option for Data or Analytics"
+    },
+    "onboarding.departments.comms":
+    {
+        "defaultMessage": "Communications",
+        "description": "Department option for Communication"
+    },
+    "onboarding.departments.customer_experience":
+    {
+        "defaultMessage": "Customer Experience",
+        "description": "Department option for Customer Experience"
+    },
+    "onboarding.departments.design":
+    {
+        "defaultMessage": "Design",
+        "description": "Department option for Design"
+    },
+    "onboarding.departments.education_professional":
+    {
+        "defaultMessage": "Education Professional",
+        "description": "Department option for Education Professional"
+    },
+    "onboarding.departments.engineering":
+    {
+        "defaultMessage": "Engineering",
+        "description": "Department option for Engineering"
+    },
+    "onboarding.departments.finance":
+    {
+        "defaultMessage": "Finance or Accounting",
+        "description": "Department option for Finance or Accounting"
+    },
+    "onboarding.departments.healthcare":
+    {
+        "defaultMessage": "Healthcare Professional",
+        "description": "Department option for Healthcare Professional"
+    },
+    "onboarding.departments.human_resources":
+    {
+        "defaultMessage": "Human Resources",
+        "description": "Department option for Human Resources"
+    },
+    "onboarding.departments.it":
+    {
+        "defaultMessage": "Information Technology (IT)",
+        "description": "Department option for Information Technology (IT)"
+    },
+    "onboarding.departments.legal":
+    {
+        "defaultMessage": "Legal",
+        "description": "Department option for Legal"
+    },
+    "onboarding.departments.marketing":
+    {
+        "defaultMessage": "Marketing",
+        "description": "Department option for Marketing"
+    },
+    "onboarding.departments.ops":
+    {
+        "defaultMessage": "Operations",
+        "description": "Department option for Operations"
+    },
+    "onboarding.departments.other":
+    {
+        "defaultMessage": "Other",
+        "description": "Department option for Other"
+    },
+    "onboarding.departments.partnerships":
+    {
+        "defaultMessage": "Partnerships",
+        "description": "Department option for Partnerships"
+    },
+    "onboarding.departments.product":
+    {
+        "defaultMessage": "Product Management",
+        "description": "Department option for Product Management"
+    },
+    "onboarding.departments.project_management":
+    {
+        "defaultMessage": "Project or Program Management",
+        "description": "Department option for Project or Program Management"
+    },
+    "onboarding.departments.research":
+    {
+        "defaultMessage": "Research & Development",
+        "description": "Department option for Research & Development"
+    },
+    "onboarding.departments.sales":
+    {
+        "defaultMessage": "Sales",
+        "description": "Department option for Sales"
+    },
+    "onboarding.gettingStartedButton":
+    {
+        "defaultMessage": "Okay, let’s go",
+        "description": "Button to accept the getting started modal"
+    },
+    "onboarding.primaryRoleTitle":
+    {
+        "defaultMessage": "What's your primary role?",
+        "description": "Question asking the user about their primary role"
+    },
+    "onboarding.role.business_owner":
+    {
+        "defaultMessage": "Business Owner",
+        "description": "Label for the role: Business Owner"
+    },
+    "onboarding.role.director":
+    {
+        "defaultMessage": "Director",
+        "description": "Label for the role: Director"
+    },
+    "onboarding.role.executive":
+    {
+        "defaultMessage": "Executive",
+        "description": "Label for the role: Executive"
+    },
+    "onboarding.role.freelancer":
+    {
+        "defaultMessage": "Freelancer",
+        "description": "Label for the role: Freelancer"
+    },
+    "onboarding.role.manager":
+    {
+        "defaultMessage": "Manager",
+        "description": "Label for the role: Manager"
+    },
+    "onboarding.role.other":
+    {
+        "defaultMessage": "Other",
+        "description": "Label for the role: Other"
+    },
+    "onboarding.role.student":
+    {
+        "defaultMessage": "Student",
+        "description": "Label for the role: Student"
+    },
+    "onboarding.role.team_member":
+    {
+        "defaultMessage": "Team Member/ Individual Contributor",
+        "description": "Label for the role: Team Member/ Individual Contributor"
+    },
+    "onboarding.selectAll":
+    {
+        "defaultMessage": "Select all that apply",
+        "description": "Instruction for multi-select options"
+    },
+    "onboarding.skipButton":
+    {
+        "defaultMessage": "Skip",
+        "description": "Label for the skip button"
+    },
+    "onboarding.tailorChatGPT":
+    {
+        "defaultMessage": "This will help us tailor ChatGPT for you.",
+        "description": "Description explaining the reason for the questions"
+    },
+    "onboarding.warning":
+    {
+        "defaultMessage": "Don’t share sensitive info",
+        "description": "Title for the warning about ChatGPT traning"
+    },
+    "onboarding.warningBody":
+    {
+        "defaultMessage": "Chat history may be reviewed or used to improve our services. Learn more about your choices in our <article>Help Center</article>.",
+        "description": "Body copy for the warning about ChatGPT traning"
+    },
+    "onboarding.workTypeTitle":
+    {
+        "defaultMessage": "What kind of work do you do?",
+        "description": "Question asking the user about the kind of work they do"
+    },
+    "onboarding.workUse":
+    {
+        "defaultMessage": "Made for use at work",
+        "description": "Title for the warning about ChatGPT business workspace use"
+    },
+    "onboarding.workUseBody":
+    {
+        "defaultMessage": "By default, chats in this workspace are not used to train our AI models.",
+        "description": "Body copy for the warning about ChatGPT business workspace use"
+    },
+    "onboarding.workspaceWelcome":
+    {
+        "defaultMessage": "Welcome to the {workspaceName} Workspace",
+        "description": "Introduction welcome for workspace onboarding"
+    },
+    "onboarding.workspaceWelcomeBody":
+    {
+        "defaultMessage": "Here you can use our latest models, with more capabilities, and fewer limits.",
+        "description": "Introduction welcome body for workspace onboarding"
+    },
+    "onboarding.workspaceWelcomeNoName":
+    {
+        "defaultMessage": "Welcome to the your Workspace",
+        "description": "Introduction welcome for workspace onboarding when no Workspace name is available"
+    },
+    "organizationBillingInfo.activeEnterpriseLicense":
+    {
+        "defaultMessage": "Enterprise License",
+        "description": "Active enterprise license name"
+    },
+    "organizationBillingInfo.activeTeamLicense":
+    {
+        "defaultMessage": "Team License",
+        "description": "Active team license name"
+    },
+    "organizationBillingInfo.billingLearnMore":
+    {
+        "defaultMessage": "Learn more",
+        "description": "Link to billing help article"
+    },
+    "organizationBillingInfo.cancelSubscriptionBtn":
+    {
+        "defaultMessage": "Cancel subscription",
+        "description": "Cancel subscription button"
+    },
+    "organizationBillingInfo.deactivateDate":
+    {
+        "defaultMessage": "Deactivates on {expiryDate, date, long}",
+        "description": "Subscription renewal date"
+    },
+    "organizationBillingInfo.defaultInvoiceName":
+    {
+        "defaultMessage": "Invoice",
+        "description": "invoice name when we are missing dates"
+    },
+    "organizationBillingInfo.inactiveLicense":
+    {
+        "defaultMessage": "Inactive License",
+        "description": "Inactive license name"
+    },
+    "organizationBillingInfo.invoiceName":
+    {
+        "defaultMessage": "Invoice: {createdDate, date, long}",
+        "description": "invoice name"
+    },
+    "organizationBillingInfo.invoicesLoadError.0":
+    {
+        "defaultMessage": "Failed to load invoices. Contact support@openai.com if error persists.",
+        "description": "Error message when invoices fail to load"
+    },
+    "organizationBillingInfo.invoicesTitle":
+    {
+        "defaultMessage": "Invoices",
+        "description": "Title for the organization invoices"
+    },
+    "organizationBillingInfo.licenseExpiry":
+    {
+        "defaultMessage": "Active until {expiryDate, date, long}",
+        "description": "License expiry date"
+    },
+    "organizationBillingInfo.managePaymentMethodBtn":
+    {
+        "defaultMessage": "Manage payment method",
+        "description": "Manage payment method button"
+    },
+    "organizationBillingInfo.manageSubscription":
+    {
+        "defaultMessage": "Manage subscription",
+        "description": "Manage subscription header"
+    },
+    "organizationBillingInfo.noInvoices":
+    {
+        "defaultMessage": "No invoices found",
+        "description": "Error message when no invoices are found"
+    },
+    "organizationBillingInfo.paidInvoice":
+    {
+        "defaultMessage": "Paid: {invoiceDate, date, long}",
+        "description": "Paid invoice date"
+    },
+    "organizationBillingInfo.planTitle":
+    {
+        "defaultMessage": "Plan",
+        "description": "Title for the organization billing plan"
+    },
+    "organizationBillingInfo.renewalDate":
+    {
+        "defaultMessage": "Renews on {expiryDate, date, long}",
+        "description": "Subscription renewal date"
+    },
+    "organizationBillingInfo.seatsInUse":
+    {
+        "defaultMessage": "{numSeats} in use ({numSeatsPct})",
+        "description": "Number of seats in use"
+    },
+    "organizationBillingInfo.seatsTitle":
+    {
+        "defaultMessage": "Seats",
+        "description": "Title for the organization billing seats"
+    },
+    "organizationBillingInfo.stripeErrorWarning":
+    {
+        "defaultMessage": "Error loading account information",
+        "description": "Error message for loading stripe account information"
+    },
+    "organizationBillingInfo.subtitle":
+    {
+        "defaultMessage": "Only workspace owners can view and change these settings",
+        "description": "Subtitle for the organization billing info page"
+    },
+    "organizationBillingInfo.teamsAutochargeMessage":
+    {
+        "defaultMessage": "Your additional seats will be included on your next invoice.",
+        "description": "Message to inform owners they will be charged for additional seats"
+    },
+    "organizationBillingInfo.title.1":
+    {
+        "defaultMessage": "Billing",
+        "description": "Title for the organization billing info page"
+    },
+    "organizationBillingInfo.upcomingInvoice":
+    {
+        "defaultMessage": "Due: {invoiceDate, date, long}",
+        "description": "Upcoming invoice date"
+    },
+    "parallelBrowsingMessage.almostDone":
+    {
+        "defaultMessage": "Almost done",
+        "description": "Status message when browsing is almost done visiting sites"
+    },
+    "parallelBrowsingMessage.cancelledV2":
+    {
+        "defaultMessage": "Stopped doing research",
+        "description": "Status message when browsing was cancelled"
+    },
+    "parallelBrowsingMessage.creatingPlanV2":
+    {
+        "defaultMessage": "Making a research plan",
+        "description": "Status message when browsing is being planned"
+    },
+    "parallelBrowsingMessage.running_2":
+    {
+        "defaultMessage": "Visiting {numTasks, plural, one {# site} other {# sites}}",
+        "description": "Status message when browsing is visiting sites"
+    },
+    "personalizationSettings.customInstructions":
+    {
+        "defaultMessage": "Custom instructions",
+        "description": "Custom instructions settings label"
+    },
+    "personalizationSettings.manageMemoriesButton":
+    {
+        "defaultMessage": "Manage",
+        "description": "Manage Memories button label"
+    },
+    "personalizationSettings.memories":
+    {
+        "defaultMessage": "Memory",
+        "description": "Memories settings label"
+    },
+    "personalizationSettings.off":
+    {
+        "defaultMessage": "Off",
+        "description": "Off"
+    },
+    "personalizationSettings.on":
+    {
+        "defaultMessage": "On",
+        "description": "On"
+    },
+    "pluginDisplayParams.generatedImage":
+    {
+        "defaultMessage": "Generated by plugin",
+        "description": "Description text for an image that was generated by a plugin"
+    },
+    "popoverNavigation.chatPreferences":
+    {
+        "defaultMessage": "Custom instructions",
+        "description": "Custom instructions menu item"
+    },
+    "popoverNavigation.myGpts":
+    {
+        "defaultMessage": "My GPTs",
+        "description": "My GPTs menu item"
+    },
+    "popoverNavigation.myPlan":
+    {
+        "defaultMessage": "My plan",
+        "description": "My plan menu item"
+    },
+    "pricingPlanConstants.business.callToAction":
+    {
+        "defaultMessage": "Buy for my team",
+        "description": "Business purchase call to action"
+    },
+    "pricingPlanConstants.free.callToAction":
+    {
+        "defaultMessage": "Your current plan",
+        "description": "Call to action for free plan"
+    },
+    "pricingPlanConstants.free.costInDollars":
+    {
+        "defaultMessage": "USD $0/month",
+        "description": "Cost for free plan"
+    },
+    "pricingPlanConstants.free.demandAccess":
+    {
+        "defaultMessage": "Access to our GPT-3.5 model",
+        "description": "Access rights for free plan"
+    },
+    "pricingPlanConstants.free.freeAdvertisedFeatures0":
+    {
+        "defaultMessage": "Unlimited messages, interactions, and history",
+        "description": "Free plan feature message"
+    },
+    "pricingPlanConstants.free.freeAdvertisedFeatures1":
+    {
+        "defaultMessage": "Access to our GPT-3.5 model",
+        "description": "Free plan feature message"
+    },
+    "pricingPlanConstants.free.freeAdvertisedFeatures2":
+    {
+        "defaultMessage": "Access on Web, iOS, and Android",
+        "description": "Free plan feature message"
+    },
+    "pricingPlanConstants.free.freePlanForLine":
+    {
+        "defaultMessage": "For people just getting started with ChatGPT",
+        "description": "Free plan subtitle message"
+    },
+    "pricingPlanConstants.free.modelFeatures":
+    {
+        "defaultMessage": "Regular model updates",
+        "description": "Model features for free plan"
+    },
+    "pricingPlanConstants.free.name":
+    {
+        "defaultMessage": "Free",
+        "description": "Name of the free pricing plan"
+    },
+    "pricingPlanConstants.free.responseSpeed":
+    {
+        "defaultMessage": "Standard response speed",
+        "description": "Response speed for free plan"
+    },
+    "pricingPlanConstants.getHelp.callToAction":
+    {
+        "defaultMessage": "I need help with a billing issue",
+        "description": "Help for billing issues"
+    },
+    "pricingPlanConstants.highDemandDisabledText":
+    {
+        "defaultMessage": "Due to high demand, we've temporarily paused upgrades.",
+        "description": "Message shown when demand is too high and payments are disabled"
+    },
+    "pricingPlanConstants.manageSubscriptionAndroid.callToAction":
+    {
+        "defaultMessage": "Manage my subscription in the ChatGPT Android app",
+        "description": "Android subscription management"
+    },
+    "pricingPlanConstants.manageSubscriptionIos.callToAction":
+    {
+        "defaultMessage": "Manage my subscription in the ChatGPT iOS app",
+        "description": "iOS subscription management"
+    },
+    "pricingPlanConstants.manageSubscriptionWeb.callToAction":
+    {
+        "defaultMessage": "Manage my subscription",
+        "description": "Web subscription management"
+    },
+    "pricingPlanConstants.plus.callToAction.active":
+    {
+        "defaultMessage": "Your current plan",
+        "description": "Active call to action for plus plan"
+    },
+    "pricingPlanConstants.plus.callToAction.inactivePayment":
+    {
+        "defaultMessage": "Upgrade to Plus",
+        "description": "Inactive payment call to action for plus plan"
+    },
+    "pricingPlanConstants.plus.costInDollars":
+    {
+        "defaultMessage": "USD $20/month",
+        "description": "Cost for plus plan"
+    },
+    "pricingPlanConstants.plus.demandAccess":
+    {
+        "defaultMessage": "Access to GPT-4, our most capable model",
+        "description": "Access rights for plus plan"
+    },
+    "pricingPlanConstants.plus.forLine":
+    {
+        "defaultMessage": "Everything in Free, and:",
+        "description": "Plus plan for line"
+    },
+    "pricingPlanConstants.plus.modelFeatures":
+    {
+        "defaultMessage": "Access to beta features like browsing, plugins, and advanced data analysis",
+        "description": "Model features for plus plan"
+    },
+    "pricingPlanConstants.plus.name":
+    {
+        "defaultMessage": "Plus",
+        "description": "Name of the ChatGPT Plus pricing plan"
+    },
+    "pricingPlanConstants.plus.plusAdvertisedFeatures0":
+    {
+        "defaultMessage": "Access to GPT-4, our most capable model",
+        "description": "Plus plan feature message"
+    },
+    "pricingPlanConstants.plus.plusAdvertisedFeatures1":
+    {
+        "defaultMessage": "Browse, create, and use GPTs",
+        "description": "Plus plan feature message"
+    },
+    "pricingPlanConstants.plus.plusAdvertisedFeatures2":
+    {
+        "defaultMessage": "Access to additional tools like DALL\\xb7E, Browsing, Advanced Data Analysis and more",
+        "description": "Plus plan feature message"
+    },
+    "pricingPlanConstants.plus.plusPricingYearlyDisclaimer":
+    {
+        "defaultMessage": "* Price billed annually",
+        "description": "Yearly Plus plan disclaimer message"
+    },
+    "pricingPlanConstants.plus.responseSpeed":
+    {
+        "defaultMessage": "Faster response speed",
+        "description": "Response speed for plus plan"
+    },
+    "pricingPlanConstants.plusWaitlistSignupSuccess":
+    {
+        "defaultMessage": "You've been added to the waitlist to upgrade to Plus",
+        "description": "Plus waitlist signup success message"
+    },
+    "pricingPlanConstants.plusYearly.costInDollars":
+    {
+        "defaultMessage": "USD $199.99/year",
+        "description": "Cost for plus yearly plan"
+    },
+    "pricingPlanConstants.signUpForWaitlistActive":
+    {
+        "defaultMessage": "Sign up for waitlist",
+        "description": "Sign up for waitlist call to action"
+    },
+    "pricingPlanConstants.signUpForWaitlistInactive":
+    {
+        "defaultMessage": "Signed up for waitlist",
+        "description": "Signed up for waitlist call to action"
+    },
+    "pricingPlanConstants.teamWaitlistSignupSuccess":
+    {
+        "defaultMessage": "You've been added to the waitlist to upgrade to Team",
+        "description": "Team waitlist signup success message"
+    },
+    "pricingPlanConstants.teams.teamPlanActive":
+    {
+        "defaultMessage": "Your current plan",
+        "description": "Upgrade to team message"
+    },
+    "pricingPlanConstants.teams.teamPlanContext":
+    {
+        "defaultMessage": "4x longer context lets you work with larger material",
+        "description": "Team plan feature message"
+    },
+    "pricingPlanConstants.teams.teamPlanCreate":
+    {
+        "defaultMessage": "Add Team workspace",
+        "description": "Add Team workspace"
+    },
+    "pricingPlanConstants.teams.teamPlanForLine":
+    {
+        "defaultMessage": "Everything in Plus, and:",
+        "description": "Team plan subtitle message"
+    },
+    "pricingPlanConstants.teams.teamPlanInactive":
+    {
+        "defaultMessage": "Upgrade to Team",
+        "description": "Upgrade to team message"
+    },
+    "pricingPlanConstants.teams.teamPlanName":
+    {
+        "defaultMessage": "Team",
+        "description": "Team plan title message"
+    },
+    "pricingPlanConstants.teams.teamPlanSubTitle":
+    {
+        "defaultMessage": "USD $25 per person/month*",
+        "description": "Team plan subtitle message"
+    },
+    "pricingPlanConstants.teams.teamPlanUsageRates":
+    {
+        "defaultMessage": "Unlimited high-speed GPT-4",
+        "description": "Team plan feature message"
+    },
+    "pricingPlanConstants.teams.teamPricingDisclaimer":
+    {
+        "defaultMessage": "* Price billed annually, minimum 2 users",
+        "description": "Team plan disclaimer message"
+    },
+    "pricingPlanConstants.teams.teamsAdvertisedFeatures0":
+    {
+        "defaultMessage": "Higher message caps on GPT-4 and tools like DALL\\xb7E, Browsing, Advanced Data Analysis, and more",
+        "description": "Team plan feature message"
+    },
+    "pricingPlanConstants.teams.teamsAdvertisedFeatures1":
+    {
+        "defaultMessage": "Create and share GPTs with your workspace",
+        "description": "Team plan feature message"
+    },
+    "pricingPlanConstants.teams.teamsAdvertisedFeatures2":
+    {
+        "defaultMessage": "Admin console for workspace management",
+        "description": "Team plan feature message"
+    },
+    "pricingPlanConstants.teams.teamsAdvertisedFeatures3":
+    {
+        "defaultMessage": "Team data excluded from training by default. <link>Learn more</link>",
+        "description": "Team plan feature message"
+    },
+    "rating.instructions":
+    {
+        "defaultMessage": "Is this conversation helpful so far?",
+        "description": "Ask the user for their rating of the conversation so far"
+    },
+    "rating.thanks":
+    {
+        "defaultMessage": "Thanks for your feedback!",
+        "description": "Thank the user for their rating"
+    },
+    "removeThirdPartyGPTModal.accessModalPrompt":
+    {
+        "defaultMessage": "Are you sure you want to revoke approval of {gptName} for your workspace? Members won't be able to chat with this GPT anymore.",
+        "description": "Prompt for removing a gpt"
+    },
+    "removeThirdPartyGPTModal.accessModalTitle":
+    {
+        "defaultMessage": "Remove GPT from workspace",
+        "description": "Title for removing a workspace gpt modal"
+    },
+    "removeThirdPartyGPTModal.cancelButtonText":
+    {
+        "defaultMessage": "Cancel",
+        "description": "Cancel button text"
+    },
+    "removeThirdPartyGPTModal.removeButtonText":
+    {
+        "defaultMessage": "Remove GPT",
+        "description": "Remove button text"
+    },
+    "selectedModal.back":
+    {
+        "defaultMessage": "Back",
+        "description": "Back button text to return to previous step"
+    },
+    "selectedModal.title":
+    {
+        "defaultMessage": "Are you sure?",
+        "description": "Title for selected modal to confirm"
+    },
+    "settingsModal.addDomain":
+    {
+        "defaultMessage": "Verify new domain",
+        "description": "Add domain button text"
+    },
+    "settingsModal.apiAccessDeletionWarning-2":
+    {
+        "defaultMessage": "Deletion will prevent you from accessing OpenAI services, including ChatGPT, API, and DALL\\xb7E.",
+        "description": "Warning message about API access being deleted."
+    },
+    "settingsModal.archiveChatsButton":
+    {
+        "defaultMessage": "Archive all",
+        "description": "Archive all chat button"
+    },
+    "settingsModal.archiveChatsLabel":
+    {
+        "defaultMessage": "Archive all chats",
+        "description": "Label for the archive all chats button"
+    },
+    "settingsModal.archiveHistoryModalCancel":
+    {
+        "defaultMessage": "Cancel",
+        "description": "Cancel button for the archive history modal"
+    },
+    "settingsModal.archiveHistoryModalConfirm":
+    {
+        "defaultMessage": "Confirm archive",
+        "description": "Confirm button for the archive history modal"
+    },
+    "settingsModal.archiveHistoryModalTitle":
+    {
+        "defaultMessage": "Archive your chat history - are you sure?",
+        "description": "Title for the archive history modal"
+    },
+    "settingsModal.archiveHistorySuccess":
+    {
+        "defaultMessage": "Successfully archived chats. You can view your archived chats in Settings.",
+        "description": "Message shown when a chat is archived"
+    },
+    "settingsModal.archivedConversations":
+    {
+        "defaultMessage": "Archived chats",
+        "description": "Label for the Archived chats setting"
+    },
+    "settingsModal.betaAdvancedDataAnalysisToggleDescription":
+    {
+        "defaultMessage": "Try a version of ChatGPT that knows how to write and execute python code, and can work with file uploads. Try asking for help with data analysis, image conversions, or editing a code file. Note: files will not persist beyond a single session.",
+        "description": "Description for the Advanced data analysis beta toggle."
+    },
+    "settingsModal.betaAdvancedDataAnalysisToggleLabel":
+    {
+        "defaultMessage": "Advanced data analysis",
+        "description": "Label for the Advanced data analysis beta toggle."
+    },
+    "settingsModal.betaBrowsingToggleDescription":
+    {
+        "defaultMessage": "Try a version of ChatGPT that knows when and how to browse the internet to answer questions about recent topics and events.",
+        "description": "Description for the Browsing beta toggle."
+    },
+    "settingsModal.betaBrowsingToggleLabel":
+    {
+        "defaultMessage": "Browse with Bing",
+        "description": "Label for the Browse with Bing beta toggle."
+    },
+    "settingsModal.betaIntro":
+    {
+        "defaultMessage": "As a Plus user, enjoy early access to experimental new features, which may change during development.",
+        "description": "Introduction for the beta features tab"
+    },
+    "settingsModal.betaPluginToggleDescription":
+    {
+        "defaultMessage": "Try a version of ChatGPT that knows when and how to use third-party plugins that you enable.",
+        "description": "Description for the Plugins beta toggle."
+    },
+    "settingsModal.betaPluginToggleLabel":
+    {
+        "defaultMessage": "Plugins",
+        "description": "Label for the Plugins beta toggle."
+    },
+    "settingsModal.betaSettingsUpdateFailed":
+    {
+        "defaultMessage": "Failed to update your beta setting",
+        "description": "Message shown when there's an error updating beta settings"
+    },
+    "settingsModal.betaTab":
+    {
+        "defaultMessage": "Beta features",
+        "description": "Label for the Beta Features tab"
+    },
+    "settingsModal.byName":
+    {
+        "defaultMessage": "by {name}",
+        "description": "By creator name under GPT"
+    },
+    "settingsModal.chatHistoryDescription":
+    {
+        "defaultMessage": "Save new chats on this browser to your history and allow them to be used to improve our models. Unsaved chats will be deleted from our systems within 30 days. This setting does not sync across browsers or devices. <link>Learn more</link>",
+        "description": "Description for the chat history setting"
+    },
+    "settingsModal.chatHistoryLabel":
+    {
+        "defaultMessage": "Chat history",
+        "description": "Label for the chat history toggle."
+    },
+    "settingsModal.chatHistoryOnlyDescription":
+    {
+        "defaultMessage": "Save new chats on this browser to your history. Unsaved chats will be deleted from our systems within 30 days. This setting does not sync across browsers or devices. <link>Learn more</link>",
+        "description": "Description for the chat history setting"
+    },
+    "settingsModal.chatHistoryToggleLabel":
+    {
+        "defaultMessage": "Chat history & training",
+        "description": "Label for the chat history toggle."
+    },
+    "settingsModal.chatTrainingEnterpriseDescription":
+    {
+        "defaultMessage": "This workspace is private and opted out of training.",
+        "description": "Description for the disabled chat training toggle."
+    },
+    "settingsModal.chatTrainingEnterpriseTooltip":
+    {
+        "defaultMessage": "ChatGPT Enterprise automatically disables training.",
+        "description": "Tooltip for the disabled chat training toggle."
+    },
+    "settingsModal.chatTrainingLabel":
+    {
+        "defaultMessage": "Chat training",
+        "description": "Label for the chat training toggle."
+    },
+    "settingsModal.chatTrainingTeamsTooltip":
+    {
+        "defaultMessage": "ChatGPT Team automatically disables training.",
+        "description": "Tooltip for the disabled chat training toggle."
+    },
+    "settingsModal.connectorsTab":
+    {
+        "defaultMessage": "Connected apps",
+        "description": "Label for the Connected Apps tab"
+    },
+    "settingsModal.cookieManagement":
+    {
+        "defaultMessage": "Cookie preferences",
+        "description": "Label for the cookie preferences button"
+    },
+    "settingsModal.cookieManagementButton":
+    {
+        "defaultMessage": "Manage",
+        "description": "Manage cookie management button"
+    },
+    "settingsModal.cookies":
+    {
+        "defaultMessage": "Cookie Preferences",
+        "description": "Label for the cookies tab"
+    },
+    "settingsModal.createrProfileUnverifiedDisclaimer":
+    {
+        "defaultMessage": "Complete verification to publish GPTs to everyone.",
+        "description": "Disclaimer for unverified users on the builder profile tab"
+    },
+    "settingsModal.creatorProfileBusinessUnverifiedDisclaimer":
+    {
+        "defaultMessage": "Contact your workspace administrator to set up a builder profile",
+        "description": "Additional disclaimer for unverified business users on the builder profile tab"
+    },
+    "settingsModal.creatorProfileDescription":
+    {
+        "defaultMessage": "Personalize your builder profile to connect with users of your GPTs. These settings apply to publicly shared GPTs.",
+        "description": "Description for the builder profile tab"
+    },
+    "settingsModal.creatorProfileLinkHeader.0":
+    {
+        "defaultMessage": "Website",
+        "description": "Links header on builder Profile tab"
+    },
+    "settingsModal.creatorProfileNameLabel":
+    {
+        "defaultMessage": "Name",
+        "description": "Label for the Name field on builder Profile tab."
+    },
+    "settingsModal.creatorProfilePersonalUnverifiedDisclaimer":
+    {
+        "defaultMessage": "Verify your identity by adding billing details or verifying ownership of a public domain name.",
+        "description": "Additional disclaimer for unverified personal users on the builder profile tab"
+    },
+    "settingsModal.dark":
+    {
+        "defaultMessage": "Dark",
+        "description": "Option for the dark theme"
+    },
+    "settingsModal.dataControls":
+    {
+        "defaultMessage": "Data controls",
+        "description": "Label for the data controls tab"
+    },
+    "settingsModal.dataExportFailed":
+    {
+        "defaultMessage": "We were unable to process your export at this time. Please try again later.",
+        "description": "Message shown when a data export request fails"
+    },
+    "settingsModal.dataExportModalCancel":
+    {
+        "defaultMessage": "Cancel",
+        "description": "Cancel button for the data export modal"
+    },
+    "settingsModal.dataExportModalConfirm":
+    {
+        "defaultMessage": "Confirm export",
+        "description": "Confirm button for the data export modal"
+    },
+    "settingsModal.dataExportModalDescription1":
+    {
+        "defaultMessage": "Your account details and chats will be included in the export.",
+        "description": "Description for the data export modal"
+    },
+    "settingsModal.dataExportModalDescription2":
+    {
+        "defaultMessage": "The data will be sent to your registered email in a downloadable file.",
+        "description": "Description for the data export modal"
+    },
+    "settingsModal.dataExportModalDescription3":
+    {
+        "defaultMessage": "The download link will expire 24 hours after you receive it.",
+        "description": "Description for the data export modal"
+    },
+    "settingsModal.dataExportModalDescription4":
+    {
+        "defaultMessage": "Processing may take some time. You'll be notified when it's ready.",
+        "description": "Description for the data export modal"
+    },
+    "settingsModal.dataExportModalDescription5":
+    {
+        "defaultMessage": "To proceed, click \"Confirm export\" below.",
+        "description": "Description for the data export modal"
+    },
+    "settingsModal.dataExportModalTitle":
+    {
+        "defaultMessage": "Request data export - are you sure?",
+        "description": "Title for the data export modal"
+    },
+    "settingsModal.dataExportRequested":
+    {
+        "defaultMessage": "Successfully exported data. You should receive an email shortly with your data.",
+        "description": "Message shown when a data export request is received"
+    },
+    "settingsModal.dataRemovalWarning-2":
+    {
+        "defaultMessage": "Your data will be deleted within 30 days, except we may retain a limited set of data for longer where required or permitted by law.",
+        "description": "Warning message about data removal after account deletion."
+    },
+    "settingsModal.deleteAccount":
+    {
+        "defaultMessage": "Delete account",
+        "description": "Label for the delete account button"
+    },
+    "settingsModal.deleteAccountButtonLabel":
+    {
+        "defaultMessage": "Permanently delete my account",
+        "description": "Label for the button to confirm account deletion."
+    },
+    "settingsModal.deleteAccountFailed":
+    {
+        "defaultMessage": "Failed to delete account. Please try again later.",
+        "description": "Message shown when there's an error deleting the user's account."
+    },
+    "settingsModal.deleteAccountSessionTooOld":
+    {
+        "defaultMessage": "Your login session is too old. Please log in again before deleting your account.",
+        "description": "Message shown when the user's login session is too old to delete their account."
+    },
+    "settingsModal.deleteAccountTitle":
+    {
+        "defaultMessage": "Delete account - are you sure?",
+        "description": "Title for the delete account confirmation modal."
+    },
+    "settingsModal.deleteAccountWarning":
+    {
+        "defaultMessage": "Deleting your account is permanent and cannot be undone.",
+        "description": "Warning message about account deletion being permanent."
+    },
+    "settingsModal.deleteButton":
+    {
+        "defaultMessage": "Delete",
+        "description": "Delete account button"
+    },
+    "settingsModal.deleteChatButton":
+    {
+        "defaultMessage": "Delete all",
+        "description": "Delete all chat button"
+    },
+    "settingsModal.deleteChatLabel":
+    {
+        "defaultMessage": "Delete all chats",
+        "description": "Label for the delete all chats button"
+    },
+    "settingsModal.deleteHelpCenter":
+    {
+        "defaultMessage": "Read our <article>help center article</article> for more information.",
+        "description": "Line that directs users to the help article for more information."
+    },
+    "settingsModal.deleteHistoryModalCancel":
+    {
+        "defaultMessage": "Cancel",
+        "description": "Cancel button for the delete history modal"
+    },
+    "settingsModal.deleteHistoryModalConfirm":
+    {
+        "defaultMessage": "Confirm deletion",
+        "description": "Confirm button for the delete history modal"
+    },
+    "settingsModal.deleteHistoryModalTitle":
+    {
+        "defaultMessage": "Clear your chat history - are you sure?",
+        "description": "Title for the delete history modal"
+    },
+    "settingsModal.disable":
+    {
+        "defaultMessage": "Disable",
+        "description": "Disable 2FA button"
+    },
+    "settingsModal.disable2fa":
+    {
+        "defaultMessage": "Disable two factor authentication",
+        "description": "Label for the mfa remove button."
+    },
+    "settingsModal.enable":
+    {
+        "defaultMessage": "Enable",
+        "description": "Enable 2FA button"
+    },
+    "settingsModal.enable2fa":
+    {
+        "defaultMessage": "Enable two-factor authentication",
+        "description": "Label for the enable 2FA button"
+    },
+    "settingsModal.exampleDescription1.1":
+    {
+        "defaultMessage": "ChatGPT will become more helpful as you chat, picking up on details and preferences to tailor its responses to you. <link>Learn more</link>",
+        "description": "Description line 1 for the MyChatgpt example"
+    },
+    "settingsModal.exampleDescription2.2":
+    {
+        "defaultMessage": "To understand what ChatGPT remembers or teach it something new, just chat with it:",
+        "description": "Description line 2 for the MyChatgpt example"
+    },
+    "settingsModal.exampleMessage1.1":
+    {
+        "defaultMessage": "Remember that I like concise responses.",
+        "description": "Message 1 for the MyChatgpt example"
+    },
+    "settingsModal.exampleMessage2.1":
+    {
+        "defaultMessage": "I just got a puppy!",
+        "description": "Message 2 for the MyChatgpt example"
+    },
+    "settingsModal.exampleMessage3.2":
+    {
+        "defaultMessage": "What do you remember about me?",
+        "description": "Message 3 for the MyChatgpt example"
+    },
+    "settingsModal.exampleMessage4":
+    {
+        "defaultMessage": "Where did we leave off on my last project?",
+        "description": "Message 4 for the MyChatgpt example"
+    },
+    "settingsModal.exportButton":
+    {
+        "defaultMessage": "Export",
+        "description": "Export data button"
+    },
+    "settingsModal.exportData":
+    {
+        "defaultMessage": "Export data",
+        "description": "Label for the export data button"
+    },
+    "settingsModal.feedbackEmailTooltip":
+    {
+        "defaultMessage": "Allow users of your GPTs to send feedback to your ChatGPT login email. Your email will never be shared publicly.",
+        "description": "Tooltip for feedback email checkbox"
+    },
+    "settingsModal.generalTab":
+    {
+        "defaultMessage": "General",
+        "description": "Label for the general tab"
+    },
+    "settingsModal.gizmoTab":
+    {
+        "defaultMessage": "Builder profile",
+        "description": "Label for the builder profile tab"
+    },
+    "settingsModal.hideNameToggle":
+    {
+        "defaultMessage": "Hide your name in your builder profile",
+        "description": "Toggle label for hiding name"
+    },
+    "settingsModal.hideWebsiteToggle":
+    {
+        "defaultMessage": "Hide your website in your builder profile",
+        "description": "Toggle label for hiding website"
+    },
+    "settingsModal.iapSubscriptionWarning":
+    {
+        "defaultMessage": "You will need to cancel your in-app purchase subscription in the Apple App Store. We cannot cancel your subscription for you.",
+        "description": "Warning message about cancelling in-app subscriptions."
+    },
+    "settingsModal.light":
+    {
+        "defaultMessage": "Light",
+        "description": "Option for the light theme"
+    },
+    "settingsModal.linkDisabledTooltip":
+    {
+        "defaultMessage": "You must have a verified domain to enable displaying a link",
+        "description": "Tooltip for disabled link toggle"
+    },
+    "settingsModal.localeAuto":
+    {
+        "defaultMessage": "Auto-detect",
+        "description": "Label for the auto-detect locale setting"
+    },
+    "settingsModal.localeDev":
+    {
+        "defaultMessage": "⋆✩★DEV★✩⋆",
+        "description": "Label for the dev locale setting"
+    },
+    "settingsModal.locale_alpha":
+    {
+        "defaultMessage": "Locale (Alpha)",
+        "description": "Label for the locale setting"
+    },
+    "settingsModal.lockedButtonLabel":
+    {
+        "defaultMessage": "Locked",
+        "description": "Label for the locked button when deleting an account."
+    },
+    "settingsModal.manageButton":
+    {
+        "defaultMessage": "Manage",
+        "description": "Manage shared/archived links/conversations button"
+    },
+    "settingsModal.myChagtGptToggleLabel.1":
+    {
+        "defaultMessage": "Improve responses with your chats",
+        "description": "Label for the chat history toggle."
+    },
+    "settingsModal.nameDisabledTooltip":
+    {
+        "defaultMessage": "You must have a verified name to enable displaying a name",
+        "description": "Tooltip for disabled name toggle"
+    },
+    "settingsModal.nameSourceReason":
+    {
+        "defaultMessage": "Name is populated from your billing details",
+        "description": "Reason for name source"
+    },
+    "settingsModal.noCreatorProfile":
+    {
+        "defaultMessage": "Unable to retrieve builder profile",
+        "description": "No builder profile error message"
+    },
+    "settingsModal.openPluginDevtools":
+    {
+        "defaultMessage": "Open plugin devtools",
+        "description": "Label for the open plugin devtools setting"
+    },
+    "settingsModal.personalization":
+    {
+        "defaultMessage": "Personalization",
+        "description": "Label for the Personalization tab"
+    },
+    "settingsModal.placeholderGPT":
+    {
+        "defaultMessage": "PlaceholderGPT",
+        "description": "Placeholder for the GPT preview on builder Profile tab."
+    },
+    "settingsModal.playStoreSubscriptionWarning":
+    {
+        "defaultMessage": "You will need to cancel your in-app purchase subscription in the Google Play Store. We cannot cancel your subscription for you.",
+        "description": "Warning message about cancelling in-app subscriptions."
+    },
+    "settingsModal.preview":
+    {
+        "defaultMessage": "Preview",
+        "description": "Preview tag in builder profile GPT preview"
+    },
+    "settingsModal.receiveFeedbackEmails":
+    {
+        "defaultMessage": "Receive feedback emails ({email})",
+        "description": "Toggle label for receiving feedback emails"
+    },
+    "settingsModal.recentLoginMessage":
+    {
+        "defaultMessage": "You may only delete your account if you have logged in within the last 10 minutes. Please log in again, then return here to continue.",
+        "description": "Message shown when the user needs to log in again to delete their account."
+    },
+    "settingsModal.refreshLoginButtonLabel":
+    {
+        "defaultMessage": "Refresh login",
+        "description": "Label for the button to refresh login."
+    },
+    "settingsModal.reset":
+    {
+        "defaultMessage": "Reset memories",
+        "description": "Label for the Reset button in MyChatGPT"
+    },
+    "settingsModal.resetFailed":
+    {
+        "defaultMessage": "Failed to reset your GPT's memory.",
+        "description": "Error message for the reset memory modal"
+    },
+    "settingsModal.resetModalCancel":
+    {
+        "defaultMessage": "Cancel",
+        "description": "Cancel button for the reset memory modal"
+    },
+    "settingsModal.resetModalConfirm":
+    {
+        "defaultMessage": "Confirm reset",
+        "description": "Confirm button for the reset memory modal"
+    },
+    "settingsModal.resetModalDescription":
+    {
+        "defaultMessage": "Your GPT will forget what it has learned from your previous chats. This can't be undone.",
+        "description": "Description for the reset memory modal"
+    },
+    "settingsModal.resetModalTitle":
+    {
+        "defaultMessage": "Are you sure?",
+        "description": "Title for the reset memory modal"
+    },
+    "settingsModal.resetSuccessful":
+    {
+        "defaultMessage": "Your GPT's memory has been reset.",
+        "description": "Success message for the reset memory modal"
+    },
+    "settingsModal.reuseEmailPhoneWarning-2":
+    {
+        "defaultMessage": "You cannot create a new account using the same email address.",
+        "description": "Warning message about not being able to reuse email for a new account."
+    },
+    "settingsModal.settings":
+    {
+        "defaultMessage": "Settings",
+        "description": "Title for the settings modal"
+    },
+    "settingsModal.sharedConversations":
+    {
+        "defaultMessage": "Shared links",
+        "description": "Label for the shared chat/link button"
+    },
+    "settingsModal.showNameToggle":
+    {
+        "defaultMessage": "Show your name in your builder profile",
+        "description": "Toggle label for showing name"
+    },
+    "settingsModal.showWebsiteToggle":
+    {
+        "defaultMessage": "Show your website in your builder profile",
+        "description": "Toggle label for showing website"
+    },
+    "settingsModal.system":
+    {
+        "defaultMessage": "System",
+        "description": "Option for the system theme"
+    },
+    "settingsModal.theme":
+    {
+        "defaultMessage": "Theme",
+        "description": "Label for the theme setting"
+    },
+    "settingsModal.trainingAllowedDescription":
+    {
+        "defaultMessage": "Allow your chats to be used to improve and train our models, which makes ChatGPT better for you and everyone who uses it. We take steps to protect your privacy. <link>Learn more</link>",
+        "description": "Description for the training-allowed setting"
+    },
+    "settingsModal.trainingAllowedToggleLabel":
+    {
+        "defaultMessage": "Improve the model for everyone",
+        "description": "Label for the training allowed toggle."
+    },
+    "settingsModal.typeDeleteInputLabel":
+    {
+        "defaultMessage": "To proceed, type \"DELETE\" in the input field below.",
+        "description": "Label for DELETE input field when deleting an account."
+    },
+    "settingsModal.typeEmailLabel":
+    {
+        "defaultMessage": "Please type your account email.",
+        "description": "Label for email input field when deleting an account."
+    },
+    "settingsModal.websiteLinkTitle":
+    {
+        "defaultMessage": "Website",
+        "description": "Website link title on builder Profile tab"
+    },
+    "sharedConversation.advancedDataAnalysisSupportDisclaimer":
+    {
+        "defaultMessage": "This chat contains files or images produced by Advanced Data Analysis which are not yet visible in Shared Chats.",
+        "description": "Disclaimer about our lack of support for Advanced Data Analysis inline images and file downloads with shared links"
+    },
+    "sharedConversation.personalizedDataDisclaimer":
+    {
+        "defaultMessage": "This conversation may reflect the link creator’s personalized data, which isn’t shared and can meaningfully change how the model responds.",
+        "description": "Disclaimer about the creator's personalized data (ex: custom instructions, memory) not being part of the shared conversation"
+    },
+    "sharedConversationModal.dateShared":
+    {
+        "defaultMessage": "Date shared",
+        "description": "Table column header"
+    },
+    "sharedConversationModal.deleteSharedAllConversations":
+    {
+        "defaultMessage": "Delete all shared links",
+        "description": "Menu item for deleting all shared links"
+    },
+    "sharedConversationModal.deleteSharedAllConversationsConfirm":
+    {
+        "defaultMessage": "Are you sure you want to delete all your shared links?",
+        "description": "Confirmation message for deleting share links"
+    },
+    "sharedConversationModal.deleteSharedAllConversationsFailed":
+    {
+        "defaultMessage": "Deleting shared links failed",
+        "description": "Toaster message when deleting all share links fails"
+    },
+    "sharedConversationModal.deleteSharedLink":
+    {
+        "defaultMessage": "Delete shared link",
+        "description": "Label for delete shared link icon"
+    },
+    "sharedConversationModal.deleteSharedLinkFailed":
+    {
+        "defaultMessage": "Deleting shared link failed",
+        "description": "Toaster message when deleting share link fails"
+    },
+    "sharedConversationModal.goToOriginConversation":
+    {
+        "defaultMessage": "View source chat",
+        "description": "Label for conversation icon"
+    },
+    "sharedConversationModal.loading":
+    {
+        "defaultMessage": "Loading...",
+        "description": "Loading message"
+    },
+    "sharedConversationModal.name":
+    {
+        "defaultMessage": "Name",
+        "description": "Table column header"
+    },
+    "sharedConversationModal.noSharedConversations":
+    {
+        "defaultMessage": "You have no shared links.",
+        "description": "No shared links message"
+    },
+    "sharedConversationModal.retry":
+    {
+        "defaultMessage": "Retry",
+        "description": "Retry button text"
+    },
+    "sharedConversationModal.somethingWentWrong":
+    {
+        "defaultMessage": "Something went wrong...",
+        "description": "Error message"
+    },
+    "sharedConversationModal.title":
+    {
+        "defaultMessage": "Shared Links",
+        "description": "Shared links modal title"
+    },
+    "sharingModal.advancedDataAnalysisSupportDisclaimer":
+    {
+        "defaultMessage": "Recipients won’t be able to view Advanced Data Analysis images or download files.",
+        "description": "Disclaimer about our lack of support for Advanced Data Analysis inline images and file downloads with shared links"
+    },
+    "sharingModal.bizDescription.1":
+    {
+        "defaultMessage": "Only members of your workspace with the URL will see the latest messages sent in this conversation. Files you attach to the conversation will not be shared, but any file contents referenced in messages will continue to be visible.",
+        "description": "Description of sharing feature in the first paragraph of the sharing modal"
+    },
+    "sharingModal.confirmDeleteLink":
+    {
+        "defaultMessage": "Are you sure you want to delete the share link?",
+        "description": "Confirmation message when deleting share link"
+    },
+    "sharingModal.deleteLink":
+    {
+        "defaultMessage": "Delete Link",
+        "description": "Button text to delete the share link"
+    },
+    "sharingModal.description":
+    {
+        "defaultMessage": "Messages you send after creating your link won't be shared. Anyone with the URL will be able to view the shared chat.",
+        "description": "Description of sharing feature in the first paragraph of the sharing modal"
+    },
+    "sharingModal.exisitingDescription":
+    {
+        "defaultMessage": "You have shared this chat <existingLink>before</existingLink>. If you want to update the shared chat content, <deleteLink>delete this link</deleteLink> and create a new shared link.",
+        "description": "Description in sharing modal when viewing an existing link"
+    },
+    "sharingModal.moderationBlocked":
+    {
+        "defaultMessage": "This shared link has been disabled by moderation.",
+        "description": "Error message in sharing modal when shared link has been moderated."
+    },
+    "sharingModal.personalizedDataDisclaimer":
+    {
+        "defaultMessage": "Any personalized data not present in the conversation won’t be shared with viewers (ex: custom instructions).",
+        "description": "Disclaimer about our policy to not share personalized data (ex: custom instructions, memory)"
+    },
+    "sharingModal.shareAnonymously":
+    {
+        "defaultMessage": "Share anonymously",
+        "description": "Button text to change sharing to be anonymous"
+    },
+    "sharingModal.shareYourName":
+    {
+        "defaultMessage": "Share your name",
+        "description": "Button text to change sharing to show the user's name"
+    },
+    "ssoModal.cancelButton":
+    {
+        "defaultMessage": "Cancel",
+        "description": "The text for the cancel button on the SSO modal"
+    },
+    "ssoModal.certLabel":
+    {
+        "defaultMessage": "X.509 Signing Certificate",
+        "description": "The label for the X.509 signing certificate input"
+    },
+    "ssoModal.createInstructionsManual":
+    {
+        "defaultMessage": "Copy your Sign-in endpoint (SSO URL) and the public X.509 certificate from your Identity Provider.",
+        "description": "Instructions for users to create an SSO connection from their IDP"
+    },
+    "ssoModal.createInstructionsXML":
+    {
+        "defaultMessage": "If your Identity Provider offers a metadata URL or an XML file, add it here for the quickest setup option.",
+        "description": "Instructions for users to create an SSO connection from an XML configuration that the IDP provides"
+    },
+    "ssoModal.createTitle":
+    {
+        "defaultMessage": "Create SAML SSO Configuration",
+        "description": "The title for the create version of the SSO modal"
+    },
+    "ssoModal.deleteButton":
+    {
+        "defaultMessage": "Delete configuration",
+        "description": "The text for the delete button on the SSO modal"
+    },
+    "ssoModal.doneButton":
+    {
+        "defaultMessage": "Done",
+        "description": "The text for the done button on the SSO modal"
+    },
+    "ssoModal.editTitle":
+    {
+        "defaultMessage": "Edit SAML SSO Configuration",
+        "description": "The title for the edit version of the SSO modal"
+    },
+    "ssoModal.nextButton":
+    {
+        "defaultMessage": "Next",
+        "description": "The text for the next button on the SSO modal"
+    },
+    "ssoModal.signInLabel":
+    {
+        "defaultMessage": "SSO URL",
+        "description": "The label for the sign in endpoint input"
+    },
+    "ssoModal.ssoMetadataLabel":
+    {
+        "defaultMessage": "App Federation Metadata URL",
+        "description": "The label for the IDP XML metadata URL input"
+    },
+    "targetedReply.replyTooltip":
+    {
+        "defaultMessage": "Reply",
+        "description": "Tooltip text for the targeted reply button"
+    },
+    "teamAccountTransferModal.confirm":
+    {
+        "defaultMessage": "Continue",
+        "description": "Confirm button text to continue to next step"
+    },
+    "teamAccountTransferModal.createDetail":
+    {
+        "defaultMessage": "Keep your personal account separate. If you’re on Plus, you will maintain your subscription until you cancel",
+        "description": "Description for keeping personal account separate"
+    },
+    "teamAccountTransferModal.createLabel":
+    {
+        "defaultMessage": "Start as empty workspace",
+        "description": "Label for create a new empty workspace option"
+    },
+    "teamAccountTransferModal.explanation":
+    {
+        "defaultMessage": "Do you want to transfer your chat history and GPTs to the Team workspace?",
+        "description": "Subtitle for initial modal for team account transfer"
+    },
+    "teamAccountTransferModal.title":
+    {
+        "defaultMessage": "Your ChatGPT Team workspace is ready",
+        "description": "Initial modal title for team account transfer"
+    },
+    "teamAccountTransferModal.transferDetail":
+    {
+        "defaultMessage": "Carry over existing data from your personal workspace",
+        "description": "Description for transfer to team option"
+    },
+    "teamAccountTransferModal.transferLabel":
+    {
+        "defaultMessage": "Transfer chat history and GPTs",
+        "description": "Label for transfer chat history to workspace"
+    },
+    "teamBilling.annualPlan.addUsersWithRenewal":
+    {
+        "defaultMessage": "Add users as needed, remove users only when you renew the contract",
+        "description": "Feature about adding/removing users in the Annual plan"
+    },
+    "teamBilling.annualPlan.billedAnnually":
+    {
+        "defaultMessage": "Annual price billed annually",
+        "description": "Feature indicating the Annual plan is billed once a year"
+    },
+    "teamBilling.annualPlan.cost":
+    {
+        "defaultMessage": "USD $25 <s>$30</s>",
+        "description": "Cost for the Annual billing plan"
+    },
+    "teamBilling.annualPlan.minBill":
+    {
+        "defaultMessage": "The minimum bill is for 2 users, USD $50/month billed annually",
+        "description": "Minimum bill details for the Annual plan"
+    },
+    "teamBilling.annualPlan.name":
+    {
+        "defaultMessage": "Annual plan",
+        "description": "Name for the Annual billing plan"
+    },
+    "teamBilling.annualPlanBilled":
+    {
+        "defaultMessage": "Price billed annually, starting {date}",
+        "description": "Summary of billing of selected plan"
+    },
+    "teamBilling.annualPlanSelected":
+    {
+        "defaultMessage": "ChatGPT Team annual plan",
+        "description": "Summary title of selected plan"
+    },
+    "teamBilling.annualPlanTotal":
+    {
+        "defaultMessage": "USD ${totalCost} per year total",
+        "description": "Summary total of selected plan"
+    },
+    "teamBilling.annualSavingsPercentage":
+    {
+        "defaultMessage": "16.7% off",
+        "description": "The savings percentage annual has on flexible"
+    },
+    "teamBilling.flexiblePlan.addRemoveUsers":
+    {
+        "defaultMessage": "Add or remove users as needed",
+        "description": "Feature indicating users can be added or removed in the Flexible plan"
+    },
+    "teamBilling.flexiblePlan.billedMonthly":
+    {
+        "defaultMessage": "Price billed monthly",
+        "description": "Feature indicating the Flexible plan is billed monthly"
+    },
+    "teamBilling.flexiblePlan.cost":
+    {
+        "defaultMessage": "USD $30",
+        "description": "Cost for the Flexible billing plan"
+    },
+    "teamBilling.flexiblePlan.minBill":
+    {
+        "defaultMessage": "The minimum monthly bill is for 2 users, USD $60/month",
+        "description": "Minimum bill details for the Flexible plan"
+    },
+    "teamBilling.flexiblePlan.name":
+    {
+        "defaultMessage": "Flexible plan",
+        "description": "Name for the Flexible billing plan"
+    },
+    "teamBilling.flexiblePlanBilled":
+    {
+        "defaultMessage": "Price billed monthly, starting {date}",
+        "description": "Summary billing of selected plan"
+    },
+    "teamBilling.flexiblePlanSelected":
+    {
+        "defaultMessage": "ChatGPT Team flexible plan",
+        "description": "Summary title of selected plan"
+    },
+    "teamBilling.flexiblePlanTotal":
+    {
+        "defaultMessage": "USD ${totalCost} per month total",
+        "description": "Summary total of selected plan"
+    },
+    "teamBilling.seatsTitle":
+    {
+        "defaultMessage": "Seats",
+        "description": "Seats section title"
+    },
+    "teamBilling.summaryTitle":
+    {
+        "defaultMessage": "Summary",
+        "description": "Summary section title"
+    },
+    "teamBilling.teamsCostStructure":
+    {
+        "defaultMessage": "per person/month",
+        "description": "The cost structure for teams plan"
+    },
+    "teamTransferModal.cancelPlus":
+    {
+        "defaultMessage": "If you have a ChatGPT Plus subscription, it will be canceled and refunded upon merging.",
+        "description": "Description to explain existing Plus subscription will be cancelled"
+    },
+    "teamTransferModal.confirm":
+    {
+        "defaultMessage": "Confirm transfer",
+        "description": "Confirm button text for transferring account"
+    },
+    "teamTransferModal.confirmTitle":
+    {
+        "defaultMessage": "Transfer of existing data is permanent and can't be undone.",
+        "description": "Title to explain transfer is permanent"
+    },
+    "teamTransferModal.deleteData":
+    {
+        "defaultMessage": "Your Plugins and custom instructions in your personal account will be deleted.",
+        "description": "Description to explain what data will be deleted"
+    },
+    "teamTransferModal.migrateData":
+    {
+        "defaultMessage": "Your chat history and GPTs in your personal account will be migrated to the Team workspace.",
+        "description": "Description to explain what data will be transferred"
+    },
+    "teamTransferModal.personalAccount":
+    {
+        "defaultMessage": "Personal account",
+        "description": "Label for personal account"
+    },
+    "teamTransferModal.warning":
+    {
+        "defaultMessage": "You will lose access to your data if you leave or are removed from the Team workspace, or if the workspace is deactivated. <link>Learn more</link>",
+        "description": "Warning for user will lose data if they leave workspace"
+    },
+    "textMessage.errorLoadingImage":
+    {
+        "defaultMessage": "Could not load image",
+        "description": "Text that describes an image that failed to load"
+    },
+    "textMessage.imageAltText":
+    {
+        "defaultMessage": "Uploaded image",
+        "description": "Alt text for image asset"
+    },
+    "textMessage.loadingImage":
+    {
+        "defaultMessage": "Loading...",
+        "description": "Text that describes a loading image"
+    },
+    "textMessage.targetedReply":
+    {
+        "defaultMessage": "Replying to:",
+        "description": "Header shown above a targeted reply"
+    },
+    "thread.archivedConversationDescription":
+    {
+        "defaultMessage": "This conversation is archived. To continue, please unarchive it first.",
+        "description": "Description when viewing an Archived conversation"
+    },
+    "thread.businessDisclaimer-oct-30":
+    {
+        "defaultMessage": "{workspaceName} workspace chats aren't used to train our models. ChatGPT can make mistakes.",
+        "description": "Business disclaimer with protected data assurance"
+    },
+    "thread.businessDisclaimerNoName-oct-30":
+    {
+        "defaultMessage": "Your workspace chats aren'ts used to train our models. ChatGPT can make mistakes.",
+        "description": "Business disclaimer with protected data assurance when no Workspace name is available"
+    },
+    "thread.chatgptMayProduceInaccurateInformation-oct-30":
+    {
+        "defaultMessage": "ChatGPT can make mistakes. Consider checking important information.",
+        "description": "ChatGPT disclaimer for producing inaccurate information"
+    },
+    "thread.cookieManager":
+    {
+        "defaultMessage": "Manage cookies",
+        "description": "Cookie manager footer link text"
+    },
+    "thread.helpAndFaq":
+    {
+        "defaultMessage": "Help & FAQ",
+        "description": "Help & FAQ menu item"
+    },
+    "thread.keyboardShortcutsMenu":
+    {
+        "defaultMessage": "Keyboard shortcuts",
+        "description": "Keyboard shortcuts menu item"
+    },
+    "thread.latencyButton":
+    {
+        "defaultMessage": "Latency",
+        "description": "Button to open the latency menu"
+    },
+    "thread.modal.onboarding.title":
+    {
+        "defaultMessage": "Do not share sensitive materials with this application",
+        "description": "Title for the onboarding warning modal"
+    },
+    "thread.modal.reportModalThankYou.description":
+    {
+        "defaultMessage": "Thank you for your report.",
+        "description": "Description for the post-report thank-you modal"
+    },
+    "thread.modal.reportModalThankYou.dismissButton":
+    {
+        "defaultMessage": "Close",
+        "description": "Close button for the post-report thank-you modal"
+    },
+    "thread.modal.reportModalThankYou.title":
+    {
+        "defaultMessage": "Thank you for your report!",
+        "description": "Title for the post-report thank-you modal"
+    },
+    "thread.modal.unrecoverableError.description":
+    {
+        "defaultMessage": "We're sorry, but something went wrong. Please try again later.",
+        "description": "Description for the UnrecoverableErrorModal"
+    },
+    "thread.modal.unrecoverableError.resetThread":
+    {
+        "defaultMessage": "Reset thread",
+        "description": "Reset thread button text"
+    },
+    "thread.modal.unrecoverableError.title":
+    {
+        "defaultMessage": "Something went wrong",
+        "description": "Title for the UnrecoverableErrorModal"
+    },
+    "thread.outdatedGptDisclaimer.0":
+    {
+        "defaultMessage": "<bold>New version of GPT available</bold> - Continue chatting to use the old version, or start a <link>new chat</link> for the latest version.",
+        "description": "Outdated GPT disclaimer"
+    },
+    "thread.privacyPolicy":
+    {
+        "defaultMessage": "Privacy policy",
+        "description": "Privacy policy footer link text"
+    },
+    "thread.releaseNotes":
+    {
+        "defaultMessage": "Release notes",
+        "description": "Release notes menu item"
+    },
+    "thread.reportSharedConversation":
+    {
+        "defaultMessage": "Report content",
+        "description": "Report shared chat footer link text"
+    },
+    "thread.sharedConversation.continue":
+    {
+        "defaultMessage": "Continue this conversation",
+        "description": "Button for shared links to allow user to continue conversation in their own history"
+    },
+    "thread.sharedConversation.moderate":
+    {
+        "defaultMessage": "Moderate conversation",
+        "description": "Button for shared links to moderate a chat for legal, safety, or other reasons"
+    },
+    "thread.sharedConversation.report":
+    {
+        "defaultMessage": "Report conversation",
+        "description": "Button for shared links to report chat for legal, safety, or other reasons"
+    },
+    "thread.sharingModal.confirmCloseWithChanges":
+    {
+        "defaultMessage": "You have unsaved changes. Do you want to continue?",
+        "description": "Confirmation message when closing share modal with changes"
+    },
+    "thread.sharingModal.copied":
+    {
+        "defaultMessage": "Copied!",
+        "description": "Status message after successfully copying the shared link"
+    },
+    "thread.sharingModal.copiedSharedConversationURL":
+    {
+        "defaultMessage": "Copied shared conversation URL to clipboard!",
+        "description": "Success message when shared conversation URL is copied"
+    },
+    "thread.sharingModal.copyLink":
+    {
+        "defaultMessage": "Copy Link",
+        "description": "Button text to copy the shared link"
+    },
+    "thread.sharingModal.copying":
+    {
+        "defaultMessage": "Copying...",
+        "description": "Status message while copying the shared link"
+    },
+    "thread.sharingModal.failedToCopyLink":
+    {
+        "defaultMessage": "Failed to copy link to clipboard",
+        "description": "Error message when failing to copy link to clipboard"
+    },
+    "thread.sharingModal.failedToDeleteSharedLink":
+    {
+        "defaultMessage": "Failed to delete shared link",
+        "description": "Error message when failing to delete shared link"
+    },
+    "thread.sharingModal.moreInfo":
+    {
+        "defaultMessage": "More Info",
+        "description": "Link to a helpdesk article with more information about the sharing modal"
+    },
+    "thread.sharingModal.title":
+    {
+        "defaultMessage": "Share link to Chat",
+        "description": "Title of sharing feature in the title of the sharing modal"
+    },
+    "thread.sharingModal.updateAndCopyLink":
+    {
+        "defaultMessage": "Update and Copy Link",
+        "description": "Button text to update and copy the shared link"
+    },
+    "thread.termsAndPolicies":
+    {
+        "defaultMessage": "Terms & policies",
+        "description": "Terms & Policies menu item"
+    },
+    "thread.termsOfUse":
+    {
+        "defaultMessage": "Terms of use",
+        "description": "Terms of use footer link text"
+    },
+    "thread.unarchiveButton":
+    {
+        "defaultMessage": "Unarchive",
+        "description": "Button to unarchive a conversation"
+    },
+    "toolsUtils.altBrowsingSearchLink1":
+    {
+        "defaultMessage": "For more information, check out these [search results]({searchLink}).",
+        "description": "Text added to the message to link to a Bing search result"
+    },
+    "toolsUtils.altBrowsingSearchLink2":
+    {
+        "defaultMessage": "To dig in deeper, check out these [search results]({searchLink}).",
+        "description": "Text added to the message to link to a Bing search result"
+    },
+    "toolsUtils.altBrowsingSearchLink3":
+    {
+        "defaultMessage": "For further details, check out these [search results]({searchLink}).",
+        "description": "Text added to the message to link to a Bing search result"
+    },
+    "toolsUtils.altBrowsingSearchLink4":
+    {
+        "defaultMessage": "To keep exploring, check out these [search results]({searchLink}).",
+        "description": "Text added to the message to link to a Bing search result"
+    },
+    "toolsUtils.altBrowsingSearchLink5":
+    {
+        "defaultMessage": "Check out these [search results]({searchLink}) for more information.",
+        "description": "Text added to the message to link to a Bing search result"
+    },
+    "toolsUtils.altBrowsingSearchLink6":
+    {
+        "defaultMessage": "Check out these [search results]({searchLink}) to dig in deeper.",
+        "description": "Text added to the message to link to a Bing search result"
+    },
+    "toolsUtils.altBrowsingSearchLink7":
+    {
+        "defaultMessage": "Check out these [search results]({searchLink}) for more details.",
+        "description": "Text added to the message to link to a Bing search result"
+    },
+    "toolsUtils.altBrowsingSearchLink8":
+    {
+        "defaultMessage": "Check out these [search results]({searchLink}) to keep exploring.",
+        "description": "Text added to the message to link to a Bing search result"
+    },
+    "toolsUtils.browsingSearchLinkPrefix1":
+    {
+        "defaultMessage": "Based on a [quick search]({searchLink}), here's what I found.",
+        "description": "Prefix added to the message to link to a Bing search result"
+    },
+    "toolsUtils.browsingSearchLinkPrefix10":
+    {
+        "defaultMessage": "I did a [quick search]({searchLink}) for more information and here's what I discovered.",
+        "description": "Prefix added to the message to link to a Bing search result"
+    },
+    "toolsUtils.browsingSearchLinkPrefix2":
+    {
+        "defaultMessage": "After a [quick search]({searchLink}), here's what I found.",
+        "description": "Prefix added to the message to link to a Bing search result"
+    },
+    "toolsUtils.browsingSearchLinkPrefix3":
+    {
+        "defaultMessage": "From a [quick search]({searchLink}), here's what I found.",
+        "description": "Prefix added to the message to link to a Bing search result"
+    },
+    "toolsUtils.browsingSearchLinkPrefix4":
+    {
+        "defaultMessage": "I did a [quick search]({searchLink}) and here's what I found.",
+        "description": "Prefix added to the message to link to a Bing search result"
+    },
+    "toolsUtils.browsingSearchLinkPrefix5":
+    {
+        "defaultMessage": "I did a [quick search]({searchLink}) for more information and here's what I found.",
+        "description": "Prefix added to the message to link to a Bing search result"
+    },
+    "toolsUtils.browsingSearchLinkPrefix6":
+    {
+        "defaultMessage": "Based on a [quick search]({searchLink}), here's what I discovered.",
+        "description": "Prefix added to the message to link to a Bing search result"
+    },
+    "toolsUtils.browsingSearchLinkPrefix7":
+    {
+        "defaultMessage": "After a [quick search]({searchLink}), here's what I discovered.",
+        "description": "Prefix added to the message to link to a Bing search result"
+    },
+    "toolsUtils.browsingSearchLinkPrefix8":
+    {
+        "defaultMessage": "From a [quick search]({searchLink}), here's what I discovered.",
+        "description": "Prefix added to the message to link to a Bing search result"
+    },
+    "toolsUtils.browsingSearchLinkPrefix9":
+    {
+        "defaultMessage": "I did a [quick search]({searchLink}) and here's what I discovered.",
+        "description": "Prefix added to the message to link to a Bing search result"
+    },
+    "transferModal.confirm":
+    {
+        "defaultMessage": "Confirm transfer",
+        "description": "Confirm button text for transferring account"
+    },
+    "transferModal.confirmLabel":
+    {
+        "defaultMessage": "Your chat history and GPTs will be transferred to the new workspace and your existing workspace will be deleted.",
+        "description": "Label to explain what will be transferred and the workspace will be deleted"
+    },
+    "transferModal.confirmTitle":
+    {
+        "defaultMessage": "Transferring chat history is permanent and can't be undone.",
+        "description": "Title to explain transfer is permanent"
+    },
+    "transferModal.proceed":
+    {
+        "defaultMessage": "To proceed, click \"Confirm transfer\" below.",
+        "description": "Prompt to proceed with the transfer account choice"
+    },
+    "useFilePickerState.maxUploadsAtATime":
+    {
+        "defaultMessage": "Unable to upload {fileName}. Max {maxUploads} uploads at a time",
+        "description": "Error message when user tries to upload more than the max number of files"
+    },
+    "useFilePickerState.retrievalSkippedFile":
+    {
+        "defaultMessage": "Unable to extract text from \"{fileName}\"",
+        "description": "Error message when user uploads a file that we are unable to parse and extract text from."
+    },
+    "useModelSwitcherModels.gpt3_5.disclaimer":
+    {
+        "defaultMessage": "Available to Free and Plus users",
+        "description": "Disclaimer message for GPT-3.5, indicating availability to Free and Plus users"
+    },
+    "useModelSwitcherModels.gpt4.disclaimer":
+    {
+        "defaultMessage": "Available exclusively to Plus users",
+        "description": "Disclaimer message for GPT-4, indicating exclusive availability to Plus users"
+    },
+    "useSubscriptionData.subscriptionLoadError":
+    {
+        "defaultMessage": "Failed to load subscription: {error}. Contact support@openai.com if error persists.",
+        "description": "Error message when subscription fails to load"
+    },
+    "useWorkspaces.adminRoleName":
+    {
+        "defaultMessage": "Admin",
+        "description": "Role name for an admin user"
+    },
+    "useWorkspaces.defaultWorkspaceTitle":
+    {
+        "defaultMessage": "Untitled Workspace",
+        "description": "title for workspace without a name"
+    },
+    "useWorkspaces.enterprisePlanName":
     {
         "defaultMessage": "Enterprise",
         "description": "label for enterprise tier account"
     },
-    "badge.teamPlanName":
+    "useWorkspaces.ownerRoleName":
+    {
+        "defaultMessage": "Owner",
+        "description": "Role name for an owner user"
+    },
+    "useWorkspaces.personalPlanName":
+    {
+        "defaultMessage": "Personal",
+        "description": "label for personal tier account"
+    },
+    "useWorkspaces.personalWorkspaceTitle":
+    {
+        "defaultMessage": "Personal account",
+        "description": "title for personal workspace"
+    },
+    "useWorkspaces.teamPlanName":
     {
         "defaultMessage": "Team",
         "description": "label for team tier account"
     },
-    "browsingMessage.browsingFailed":
-    {
-        "defaultMessage": "Error browsing",
-        "description": "Status message when browsing failed"
-    },
-    "browsingMessage.browsingStopped":
-    {
-        "defaultMessage": "Stopped browsing",
-        "description": "Status message when browsing was stopped"
-    },
-    "browsingMessage.readingDocument":
-    {
-        "defaultMessage": "Reading {filename}",
-        "description": "Status message when reading a document"
-    },
-    "browsingMessage.retrievalFailed":
-    {
-        "defaultMessage": "Error reading documents",
-        "description": "Status message when document retrieval failed"
-    },
-    "browsingMessage.retrievalStopped":
-    {
-        "defaultMessage": "Stopped reading documents",
-        "description": "Status message when document retrieval was stopped"
-    },
-    "browsingMessage.searching":
-    {
-        "defaultMessage": "Searching Bing",
-        "description": "Status message when searching Bing"
-    },
-    "browsingMessage.searchingForQuery":
-    {
-        "defaultMessage": "Searching “{query}”",
-        "description": "Status message when searching for a query"
-    },
-    "browsingMessage.searchingKnowledge":
-    {
-        "defaultMessage": "Searching my knowledge",
-        "description": "Status message when a GPT is searching its knowledge base"
-    },
-    "browsingMessage.searchingKnowledgeFailed":
-    {
-        "defaultMessage": "Error searching knowledge",
-        "description": "Status message when a GPT failed to search its knowledge base"
-    },
-    "browsingMessage.searchingKnowledgeStopped":
-    {
-        "defaultMessage": "Stopped searching knowledge",
-        "description": "Status message when a GPT stopped searching its knowledge base"
-    },
-    "browsingMessage.startingRetrieval":
-    {
-        "defaultMessage": "Reading documents",
-        "description": "Status message when document retrieval is starting"
-    },
-    "browsingMessage.startingV3":
-    {
-        "defaultMessage": "Doing research with Bing",
-        "description": "Status message when browsing is starting"
-    },
-    "browsingMessage.visiting":
-    {
-        "defaultMessage": "Visiting {url}",
-        "description": "Status message when visiting a webpage"
-    },
-    "cancelTeamPlanModal.areYouSure":
-    {
-        "defaultMessage": "Are you sure?",
-        "description": "Title for confirmation prompt"
-    },
-    "cancelTeamPlanModal.cancelPlan":
-    {
-        "defaultMessage": "Cancel subscription",
-        "description": "Cancel your team plan modal title"
-    },
-    "cancelTeamPlanModal.cancelSubscriptionButton":
-    {
-        "defaultMessage": "Cancel subscription",
-        "description": "Label for the cancel subscription button"
-    },
-    "cancelTeamPlanModal.dataUnavailable":
-    {
-        "defaultMessage": "All workspace data including chat history and settings will not be available.",
-        "description": "Data will not be available message"
-    },
-    "cancelTeamPlanModal.doneButton":
-    {
-        "defaultMessage": "Done",
-        "description": "Label for the done button"
-    },
-    "cancelTeamPlanModal.emailPlaceholder":
-    {
-        "defaultMessage": "abcd@acme.com",
-        "description": "Placeholder text for email input"
-    },
-    "cancelTeamPlanModal.enterEmailToConfirm":
-    {
-        "defaultMessage": "Enter your Email address to confirm",
-        "description": "Prompt to enter email for confirmation"
-    },
-    "cancelTeamPlanModal.errorCancellingSubscription":
-    {
-        "defaultMessage": "There was a problem cancelling your subscription.",
-        "description": "Error cancelling subscription toast text"
-    },
-    "cancelTeamPlanModal.keepSubscriptionButton":
-    {
-        "defaultMessage": "Keep subscription",
-        "description": "Label for the keep subscription button"
-    },
-    "cancelTeamPlanModal.retainUntil":
-    {
-        "defaultMessage": "You will still be able to use ChatGPT with the other workspaces associated with this email address.",
-        "description": "Retain access to other workspaces with your email"
-    },
-    "cancelTeamPlanModal.successfullyCanceled":
-    {
-        "defaultMessage": "You have successfully canceled your subscription.",
-        "description": "Confirmation on account canceled text"
-    },
-    "citations.invalid":
-    {
-        "defaultMessage": "Invalid citation",
-        "description": "Text when citation is invalid"
-    },
-    "citations.viewAnalysis":
-    {
-        "defaultMessage": "View analysis",
-        "description": "Tooltip text for a citation link to analysis"
-    },
-    "codeInterpreterMessage.errorV2":
-    {
-        "defaultMessage": "<expander>Error analyzing</expander>",
-        "description": "Status message when code interpreter ran into an error"
-    },
-    "codeInterpreterMessage.finished":
-    {
-        "defaultMessage": "<expander>Finished analyzing</expander>",
-        "description": "Status message when code interpreter is finished"
-    },
-    "codeInterpreterMessage.resultLabel":
-    {
-        "defaultMessage": "Result",
-        "description": "Label shown with the code execution result output"
-    },
-    "codeInterpreterMessage.runningV2":
-    {
-        "defaultMessage": "<expander>Analyzing</expander>",
-        "description": "Status message when code interpreter is running"
-    },
-    "codeInterpreterMessage.stoppedV2":
-    {
-        "defaultMessage": "<expander>Stopped analyzing</expander>",
-        "description": "Status message when code interpreter was stopped by the user"
-    },
-    "components.business.NumSeats.description":
-    {
-        "defaultMessage": "{num} seats in use",
-        "description": "number of seats in use description"
-    },
-    "components.business.NumSeats.dividerTooltip":
-    {
-        "defaultMessage": "Your team has {count, plural, =0 {no seats} one {# seat} other {# seats}} purchased",
-        "description": "Tooltip explaining allocated seats limit"
-    },
-    "connectorSettings.connect":
-    {
-        "defaultMessage": "Connect",
-        "description": "Label for the button to connect an app"
-    },
-    "connectorSettings.connectorsTitle":
-    {
-        "defaultMessage": "Connect apps to access their information in ChatGPT.",
-        "description": "Title row for the connected apps settings tab"
-    },
-    "connectorSettings.googleDriveDesc":
-    {
-        "defaultMessage": "Attach Google Docs, Sheets, and Slides to your messages or add them as context to your conversations.",
-        "description": "Description for the Google Drive connected app"
-    },
-    "connectorSettings.googleDriveIconAlt":
-    {
-        "defaultMessage": "Icon for Google Drive",
-        "description": "Alt text for the Google Drive icon"
-    },
-    "connectorSettings.noConnectorSettings":
-    {
-        "defaultMessage": "Unable to get connected apps",
-        "description": "Text when there are no connected apps available"
-    },
-    "connectorSettings.o365Desc":
-    {
-        "defaultMessage": "Attach Microsoft Word, Excel, and Powerpoint files to your messages or add them as context to your conversations.",
-        "description": "Description for the Microsoft 365 connected app"
-    },
-    "connectorSettings.o365IconAlt":
-    {
-        "defaultMessage": "Icon for Microsoft 365",
-        "description": "Alt text for the Microsoft 365 icon"
-    },
-    "createWorkspace.cancel":
+    "useWorkspacews.standardRoleName":
+    {
+        "defaultMessage": "Member",
+        "description": "Role name for a standard user"
+    },
+    "userContextModal.aboutUserTip1":
+    {
+        "defaultMessage": "Where are you based?",
+        "description": "tips for Custom instructions about you"
+    },
+    "userContextModal.aboutUserTip2":
+    {
+        "defaultMessage": "What do you do for work?",
+        "description": "tips for Custom instructions about you"
+    },
+    "userContextModal.aboutUserTip3":
+    {
+        "defaultMessage": "What are your hobbies and interests?",
+        "description": "tips for Custom instructions about you"
+    },
+    "userContextModal.aboutUserTip4":
+    {
+        "defaultMessage": "What subjects can you talk about for hours?",
+        "description": "tips for Custom instructions about you"
+    },
+    "userContextModal.aboutUserTip5":
+    {
+        "defaultMessage": "What are some goals you have?",
+        "description": "tips for Custom instructions about you"
+    },
+    "userContextModal.aboutYouHelpText":
+    {
+        "defaultMessage": "What would you like ChatGPT to know about you to provide better responses?",
+        "description": "help text for about you section of Custom instructions"
+    },
+    "userContextModal.cancel":
     {
         "defaultMessage": "Cancel",
-        "description": "Cancel button text"
-    },
-    "createWorkspace.continueToBillingButton":
-    {
-        "defaultMessage": "Continue to billing",
-        "description": "Continue to billing button text"
-    },
-    "createWorkspace.exampleTeamWorkspaceName":
-    {
-        "defaultMessage": "Acme Inc.",
-        "description": "Example Team Workspace name"
-    },
-    "createWorkspace.paymentErrorWarning":
-    {
-        "defaultMessage": "The payments page encountered an error. Please try again. If the problem continues, please visit help.openai.com.",
-        "description": "Error toast when payment page has an error"
-    },
-    "createWorkspace.selectBillingOption":
-    {
-        "defaultMessage": "Select billing options",
-        "description": "Button text for selecting team plan"
-    },
-    "createWorkspace.selectTeamPlan":
-    {
-        "defaultMessage": "Select Team Plan",
-        "description": "Button text for selecting team plan"
-    },
-    "createWorkspace.selectTeamPlanModalTitle":
-    {
-        "defaultMessage": "Select your team plan",
-        "description": "Select team plan modal title"
-    },
-    "createWorkspace.title":
-    {
-        "defaultMessage": "Create workspace",
-        "description": "Title for the create workspace modal"
-    },
-    "createWorkspace.workspaceNameDescription":
-    {
-        "defaultMessage": "Set a workspace name for your team. The name can be changed at any time.",
-        "description": "Description text below the workspace name label"
-    },
-    "createWorkspace.workspaceNameLabel":
+        "description": "Cancel button for Custom instructions modal"
+    },
+    "userContextModal.chatPreferencesEnable":
+    {
+        "defaultMessage": "Enable for new chats",
+        "description": "chat preferences is enabled"
+    },
+    "userContextModal.confirmCloseBody":
+    {
+        "defaultMessage": "Are you sure you want to exit? Any changes you made will be permanently lost.",
+        "description": "confirm close modal"
+    },
+    "userContextModal.confirmCloseCancel":
+    {
+        "defaultMessage": "Back",
+        "description": "cancel button for confirm close modal"
+    },
+    "userContextModal.confirmCloseOk":
+    {
+        "defaultMessage": "Exit",
+        "description": "ok button for confirm close modal"
+    },
+    "userContextModal.confirmCloseTitle":
+    {
+        "defaultMessage": "You have unsaved changes.",
+        "description": "title for confirm close modal"
+    },
+    "userContextModal.disableToggleLabel":
+    {
+        "defaultMessage": "Disable chat preferences",
+        "description": "disable chat preferences toggle label"
+    },
+    "userContextModal.enableToggleLabel":
+    {
+        "defaultMessage": "Enable chat preferences",
+        "description": "enable chat preferences toggle label"
+    },
+    "userContextModal.hideTips":
+    {
+        "defaultMessage": "Hide tips",
+        "description": "hide tips button for Custom instructions modal"
+    },
+    "userContextModal.messageLimitError":
+    {
+        "defaultMessage": "Please limit your responses to {limit} characters or less.",
+        "description": "error message for Custom instructions modal"
+    },
+    "userContextModal.modApiVoilation":
+    {
+        "defaultMessage": "This content may violate our <policyLink>content policy</policyLink>. If you believe this to be in error, please <feedbackLink>submit your feedback</feedbackLink> — your input will aid our research in this area.",
+        "description": "error message for mod api voilation"
+    },
+    "userContextModal.modelHelpText":
+    {
+        "defaultMessage": "How would you like ChatGPT to respond?",
+        "description": "help text for about you section of Custom instructions"
+    },
+    "userContextModal.modelTip1":
+    {
+        "defaultMessage": "How formal or casual should ChatGPT be?",
+        "description": "tips for Custom instructions about model"
+    },
+    "userContextModal.modelTip2":
+    {
+        "defaultMessage": "How long or short should responses generally be?",
+        "description": "tips for Custom instructions about model"
+    },
+    "userContextModal.modelTip3":
+    {
+        "defaultMessage": "How do you want to be addressed?",
+        "description": "tips for Custom instructions about model"
+    },
+    "userContextModal.modelTip4":
+    {
+        "defaultMessage": "Should ChatGPT have opinions on topics or remain neutral?",
+        "description": "tips for Custom instructions about model"
+    },
+    "userContextModal.save":
+    {
+        "defaultMessage": "Save",
+        "description": "save button for my profile modal"
+    },
+    "userContextModal.showTips":
+    {
+        "defaultMessage": "Show tips",
+        "description": "show tips button for Custom instructions modal"
+    },
+    "userContextModal.subhead":
+    {
+        "defaultMessage": "<article>Learn more</article> about Custom instructions and how they’re used to help ChatGPT provide better responses.",
+        "description": "subhead for Custom instructions modal"
+    },
+    "userContextModal.tipsHeader":
+    {
+        "defaultMessage": "Thought starters",
+        "description": "header for Custom instructions tips"
+    },
+    "userContextModal.title":
+    {
+        "defaultMessage": "Custom instructions",
+        "description": "title for Custom instructions modal"
+    },
+    "workspaceAnalytics.activeUsersChartLabel":
+    {
+        "defaultMessage": "Active users",
+        "description": "Label for active users chart"
+    },
+    "workspaceAnalytics.activeUsersLastWeek":
+    {
+        "defaultMessage": "Active users last week",
+        "description": "Text under number active users last week"
+    },
+    "workspaceAnalytics.conversationsLabel":
+    {
+        "defaultMessage": "Conversations",
+        "description": "Label of total conversations count"
+    },
+    "workspaceAnalytics.messagesChartLabel":
+    {
+        "defaultMessage": "Messages",
+        "description": "Label for messages chart"
+    },
+    "workspaceAnalytics.messagesLabel":
+    {
+        "defaultMessage": "Messages",
+        "description": "Label of total messages count"
+    },
+    "workspaceAnalytics.messagesLast30Days":
+    {
+        "defaultMessage": "{count, plural, =0 {No messages} one {# message} other {# messages}} processed over last 30 days",
+        "description": "Analytics text about messages"
+    },
+    "workspaceAnalytics.messagesLastWeek":
+    {
+        "defaultMessage": "Messages processed last week",
+        "description": "Text under number messages last week"
+    },
+    "workspaceAnalytics.title":
+    {
+        "defaultMessage": "Workspace Analytics",
+        "description": "Title of analytics page"
+    },
+    "workspaceAnalytics.totals":
+    {
+        "defaultMessage": "Totals over last 30 days",
+        "description": "Section header of total analytics"
+    },
+    "workspaceAnalytics.understand":
+    {
+        "defaultMessage": "Understand how your workspace is using ChatGPT",
+        "description": "Subtext of workspace analytics page"
+    },
+    "workspaceAnalytics.usageTitle":
+    {
+        "defaultMessage": "Usage",
+        "description": "Title of usage section"
+    },
+    "workspaceAnalytics.usersLabel":
+    {
+        "defaultMessage": "Users",
+        "description": "Label of total user count"
+    },
+    "workspaceAnalytics.usersLast30Days":
+    {
+        "defaultMessage": "{count, plural, =0 {No active users} one {# active user} other {# active user}} over last 30 days",
+        "description": "Analytics text about active users"
+    },
+    "workspaceAnalytics.weeklyUsersTitle":
+    {
+        "defaultMessage": "Weekly users",
+        "description": "Title of weekly users section"
+    },
+    "workspaceAppearanceModal.cancel":
+    {
+        "defaultMessage": "Cancel",
+        "description": "Label for the cancel button"
+    },
+    "workspaceAppearanceModal.propagationWarning":
+    {
+        "defaultMessage": "Changes to the workspace name and image may take some time to take effect.",
+        "description": "Warning message about changes to the workspace profile taking a while to show up."
+    },
+    "workspaceAppearanceModal.saveError":
+    {
+        "defaultMessage": "Failed to save workspace appearance",
+        "description": "Error message when saving workspace appearance fails"
+    },
+    "workspaceAppearanceModal.submit":
+    {
+        "defaultMessage": "Save",
+        "description": "Label for the submit button"
+    },
+    "workspaceAppearanceModal.title":
+    {
+        "defaultMessage": "Workspace appearance",
+        "description": "Title for the workspace appearance modal"
+    },
+    "workspaceAppearanceModal.workspaceAvatar":
+    {
+        "defaultMessage": "Workspace image",
+        "description": "Label for the workspace image upload field"
+    },
+    "workspaceAppearanceModal.workspaceAvatarDescription":
+    {
+        "defaultMessage": "Upload a JPEG or PNG workspace image for your team. (Minimum {size}\\xd7{size}px recommended.)",
+        "description": "Help text for the workspace image upload field"
+    },
+    "workspaceAppearanceModal.workspaceName":
     {
         "defaultMessage": "Workspace name",
         "description": "Label for the workspace name input field"
     },
-    "dalleMessage.creatingImagesV2":
-    {
-        "defaultMessage": "Creating image",
-        "description": "Status message when DALL\\xb7E is creating an image"
-    },
-    "dalleMessage.errorCreatingV2":
-    {
-        "defaultMessage": "Error creating image",
-        "description": "Status message when DALL\\xb7E failed to create an image"
-    },
-    "dalleMessage.generatedImageAltText":
-    {
-        "defaultMessage": "Generated by DALL\\xb7E",
-        "description": "Alt text for images generated by DALL\\xb7E"
-    },
-    "dalleMessage.imageLoadError":
-    {
-        "defaultMessage": "Error loading image",
-        "description": "Error message when an image fails to load"
-    },
-    "dalleMessage.imageViewerMetadataCopyButton":
-    {
-        "defaultMessage": "Copy",
-        "description": "Copy button for the prompt metadata in the image viewer"
-    },
-    "dalleMessage.imageViewerMetadataCopyButtonCopied":
-    {
-        "defaultMessage": "Copied!",
-        "description": "Copy button for the prompt metadata in the image viewer when the prompt is copied"
-    },
-    "dalleMessage.imageViewerMetadataTitle":
-    {
-        "defaultMessage": "Prompt",
-        "description": "Title for the prompt metadata in the image viewer"
-    },
-    "dalleMessage.stoppedV3":
-    {
-        "defaultMessage": "Stopped creating image",
-        "description": "Status message when DALL\\xb7E was stopped by the user"
-    },
-    "deactivatedWorkspaceModal.chatHistoryUnavailable":
-    {
-        "defaultMessage": "Your chat history and settings will not be available.",
-        "description": "Description about chat history unavailability"
-    },
-    "deactivatedWorkspaceModal.createPersonalWorkspace":
-    {
-        "defaultMessage": "Create a personal workspace to continue",
-        "description": "Prompt to create a personal workspace"
-    },
-    "deactivatedWorkspaceModal.createPersonalWorkspaceButton":
-    {
-        "defaultMessage": "Create a personal workspace",
-        "description": "Button label to create a personal workspace"
-    },
-    "deactivatedWorkspaceModal.deactivatedWorkspaceReason":
-    {
-        "defaultMessage": "Because your workspace has been deactivated, you need to create a personal workspace to continue using ChatGPT.",
-        "description": "Explanation for the need to create a personal workspace"
-    },
-    "deactivatedWorkspaceModal.otherWorkspacesAvailable":
-    {
-        "defaultMessage": "You will still be able to use ChatGPT with the other workspaces associated with this email address.",
-        "description": "Description about availability of other workspaces"
-    },
-    "deactivatedWorkspaceModal.profileAlt":
-    {
-        "defaultMessage": "Profile",
-        "description": "Alt text for the profile image"
-    },
-    "deactivatedWorkspaceModal.selectWorkspace":
-    {
-        "defaultMessage": "Select a workspace to continue",
-        "description": "Prompt to select another workspace"
-    },
-    "deactivatedWorkspaceModal.workspaceDeactivated":
-    {
-        "defaultMessage": "Your workspace has been deactivated",
-        "description": "Title indicating workspace deactivation"
-    },
-    "deactivatedWorkspaceModal.workspaceDeactivatedDesc":
-    {
-        "defaultMessage": "Your workspace has been deactivated.",
-        "description": "Description of workspace deactivation"
-    },
-    "domainModal.cancel":
-    {
-        "defaultMessage": "Cancel",
-        "description": "The label for the cancel button."
-    },
-    "domainModal.check":
-    {
-        "defaultMessage": "Check",
-        "description": "The label for the check button."
-    },
-    "domainModal.copiedTXTRecordToClipboard":
-    {
-        "defaultMessage": "Copied DNS TXT record to clipboard",
-        "description": "Message informing the user their TXT record has been copied."
-    },
-    "domainModal.domainCheckError":
-    {
-        "defaultMessage": "Your domain could not be verified. Make sure you've correctly set the TXT record above. DNS records may take up to 30 minutes to propagate.",
-        "description": "Error message when domain cannot be successfully verified."
-    },
-    "domainModal.domainInput.0":
-    {
-        "defaultMessage": "Add a new domain",
-        "description": "The label for the domain input."
-    },
-    "domainModal.done":
-    {
-        "defaultMessage": "Done",
-        "description": "The label for the done button."
-    },
-    "domainModal.editTitle":
-    {
-        "defaultMessage": "Manage Domain",
-        "description": "The title for the domain modal in edit mode."
-    },
-    "domainModal.newTitle.0":
-    {
-        "defaultMessage": "Verify a new domain",
-        "description": "The title for the domain modal in new mode."
-    },
-    "domainModal.placeholder":
-    {
-        "defaultMessage": "openai.com",
-        "description": "The placeholder text domain input."
-    },
-    "domainModal.submit":
-    {
-        "defaultMessage": "Submit",
-        "description": "The label for the submit button."
-    },
-    "domainModal.successfulVerification.0":
-    {
-        "defaultMessage": "Your domain, \"{hostname}\" has been successfully verified",
-        "description": "Message showing the user their domain verification is complete.."
-    },
-    "domainModal.unverifiedDomainsTableHeader":
-    {
-        "defaultMessage": "Your unverified domains",
-        "description": "The header for the unverified domains table."
-    },
-    "domainModal.verificationTokenFooter":
-    {
-        "defaultMessage": "Then, check the record to complete the verification.",
-        "description": "Message that informs the user how to verify their domain."
-    },
-    "domainModal.verificationTokenMessage":
-    {
-        "defaultMessage": "To verify ownership of {hostname}, navigate to your DNS provider and add a TXT record with this value:",
-        "description": "Message that informs the user how to verify their domain."
-    },
-    "domainModal.verifyDomainButton":
+    "workspaceAppearanceModal.workspaceNameDescription":
+    {
+        "defaultMessage": "Update the name of the workspace.",
+        "description": "Help text for the workspace name input field"
+    },
+    "workspaceIdentity.acsURLLabel":
+    {
+        "defaultMessage": "Assertion Consumer Service (ACS) URL",
+        "description": "Label for the ACS URL configuration text"
+    },
+    "workspaceIdentity.addDomainButton.0":
+    {
+        "defaultMessage": "Add domain",
+        "description": "Label for add domain button"
+    },
+    "workspaceIdentity.autoProvisionBody":
+    {
+        "defaultMessage": "Automatically create ChatGPT Enterprise accounts for new users who log in with a verified domain.",
+        "description": "Label for toggling automatic provisioning"
+    },
+    "workspaceIdentity.autoProvisionDisabledToast":
+    {
+        "defaultMessage": "Auto provisioning disabled for this workspace",
+        "description": "Toast message for disabling auto provisioning"
+    },
+    "workspaceIdentity.autoProvisionEnabledToast":
+    {
+        "defaultMessage": "Auto provisioning enabled for this workspace",
+        "description": "Toast message for enabling auto provisioning"
+    },
+    "workspaceIdentity.autoProvisionSubtitle":
+    {
+        "defaultMessage": "Automatic account creation",
+        "description": "Subtitle for auto provisioning"
+    },
+    "workspaceIdentity.certExtractFailed":
+    {
+        "defaultMessage": "Unable to find X.509 Certificate in provided XML",
+        "description": "Notice to the user that the given XML did not contain a certificate"
+    },
+    "workspaceIdentity.copiedACSToClipboard":
+    {
+        "defaultMessage": "Copied ACS URL to clipboard",
+        "description": "Message for success toast on copying post-back url"
+    },
+    "workspaceIdentity.copiedEntityToClipboard":
+    {
+        "defaultMessage": "Copied entity ID to clipboard",
+        "description": "Message for success toast on copying entity ID"
+    },
+    "workspaceIdentity.copiedIDPUrlToClipboard":
+    {
+        "defaultMessage": "Copied IDP URL to clipboard",
+        "description": "Message for success toast on copying IDP URL"
+    },
+    "workspaceIdentity.domainSubtitle":
+    {
+        "defaultMessage": "Domain management",
+        "description": "Section header of domain management"
+    },
+    "workspaceIdentity.domainTableHeader":
+    {
+        "defaultMessage": "Domain",
+        "description": "Label for domain column on domains table"
+    },
+    "workspaceIdentity.enforceSSOBody":
+    {
+        "defaultMessage": "When SSO is active, users will no longer be able to use social logins and will be redirected to your SSO portal.",
+        "description": "Message telling users what enforcing SSO will do"
+    },
+    "workspaceIdentity.enforceSSOLabel":
+    {
+        "defaultMessage": "Toggle to enable or disable SSO enforcement",
+        "description": "Label for toggle for SSO enforcement"
+    },
+    "workspaceIdentity.enforceSSOTitle":
+    {
+        "defaultMessage": "Enforce SSO log in",
+        "description": "Section title for toggling SSO enforcement"
+    },
+    "workspaceIdentity.entityIDLabel":
+    {
+        "defaultMessage": "Entity ID",
+        "description": "Label for the Entity ID configuration text"
+    },
+    "workspaceIdentity.idpSignInURL":
+    {
+        "defaultMessage": "IDP Tile URL",
+        "description": "Label for the IDP Tile URL configuration text"
+    },
+    "workspaceIdentity.learnMoreLink":
+    {
+        "defaultMessage": "Learn more",
+        "description": "Label for link to learn more"
+    },
+    "workspaceIdentity.orDivider":
+    {
+        "defaultMessage": "or",
+        "description": "Separator between the top and bottom portions of the form"
+    },
+    "workspaceIdentity.setupSSOContent":
+    {
+        "defaultMessage": "Anyone using email addresses with a verified domain can log in via SAML SSO.",
+        "description": "Content describing how SAML SSO will work when set up"
+    },
+    "workspaceIdentity.ssoAddButton":
+    {
+        "defaultMessage": "Add SAML SSO",
+        "description": "Label for add SSO button"
+    },
+    "workspaceIdentity.ssoDropdownRemoveButtonText":
+    {
+        "defaultMessage": "Remove Domain",
+        "description": "Dropdown choice for remove domain"
+    },
+    "workspaceIdentity.ssoDropdownVerifyButtonText":
     {
         "defaultMessage": "Verify",
-        "description": "The label for the verify button."
-    },
-    "emailsTextarea.clearAllEntries":
-    {
-        "defaultMessage": "Clear all",
-        "description": "Clear all entries in the list of members to be added"
-    },
-    "emailsTextarea.membersAdded":
-    {
-        "defaultMessage": "+{count} {count, plural, one {member} other {members}}",
-        "description": "Current number of members that will be added to the workspace"
-    },
-    "emailsTextarea.placeholder":
-    {
-        "defaultMessage": "Type an email and press enter...",
-        "description": "Placeholder for the insert emails textarea"
-    },
-    "emailsTextarea.removeMember":
-    {
-        "defaultMessage": "Remove {member}",
-        "description": "Remove a member from the list of members to be added"
-    },
-    "emailsTextarea.tooltipInvalidEmail":
-    {
-        "defaultMessage": "\"{email}\" may not be a valid email",
-        "description": "Tooltip for invalid email addresses"
-    },
-    "feedbackModal.continueWithChosenAnswer":
-    {
-        "defaultMessage": "The conversation will continue with the answer you choose.",
-        "description": "Information text for user during comparison"
-    },
-    "feedbackModal.copyrightContent":
-    {
-        "defaultMessage": "This content violates copyright law",
-        "description": "Label for Copyrighted Content checkbox"
-    },
-    "feedbackModal.dontLikeThis":
-    {
-        "defaultMessage": "I don't like this",
-        "description": "Label for I Don't Like This checkbox"
-    },
-    "feedbackModal.employeeConsent":
-    {
-        "defaultMessage": "Allow this content to be used for model evals",
-        "description": "Open AI employee is consenting to allow this content to be used in evals"
-    },
-    "feedbackModal.employeeConsentExplanation":
-    {
-        "defaultMessage": "Allow your feedback and conversation to be used to in model evals. Please verify there is no confidential data in the conversation.",
-        "description": "Explanation for employee consent checkbox"
-    },
-    "feedbackModal.harmfulOffensive":
-    {
-        "defaultMessage": "This content is harmful or offensive",
-        "description": "Label for harmful/offensive checkbox"
-    },
-    "feedbackModal.harmfulUnsafe":
-    {
-        "defaultMessage": "This is harmful / unsafe",
-        "description": "Label for harmful/unsafe checkbox"
-    },
-    "feedbackModal.moderationAccept":
-    {
-        "defaultMessage": "Allow Content",
-        "description": "Button text for accepting the share link and allowing it to be viewed"
-    },
-    "feedbackModal.moderationReject":
-    {
-        "defaultMessage": "Block Content",
-        "description": "Button text for rejecting the share link and blocking it from being viewed"
-    },
-    "feedbackModal.neitherAnswerBetter":
-    {
-        "defaultMessage": "Neither answer is better",
-        "description": "Button text for choosing neither answer during comparison"
-    },
-    "feedbackModal.newAnswer":
-    {
-        "defaultMessage": "New Answer",
-        "description": "Title for the new answer during comparison"
-    },
-    "feedbackModal.newAnswerBetter":
-    {
-        "defaultMessage": "New answer is better",
-        "description": "Button text for choosing new answer during comparison"
-    },
-    "feedbackModal.notHelpful":
-    {
-        "defaultMessage": "This isn't helpful",
-        "description": "Label for not helpful checkbox"
-    },
-    "feedbackModal.notTrue":
-    {
-        "defaultMessage": "This isn't true",
-        "description": "Label for not true checkbox"
-    },
-    "feedbackModal.originalAnswer":
-    {
-        "defaultMessage": "Original Answer",
-        "description": "Title for the original answer during comparison"
-    },
-    "feedbackModal.originalAnswerBetter":
-    {
-        "defaultMessage": "Original answer is better",
-        "description": "Button text for choosing original answer during comparison"
-    },
-    "feedbackModal.pickBestAnswer":
-    {
-        "defaultMessage": "Pick the best answer to improve the model",
-        "description": "Title for the compare feedback modal"
-    },
-    "feedbackModal.provideAdditionalFeedback":
-    {
-        "defaultMessage": "Provide additional feedback",
-        "description": "Title for the critique feedback modal"
-    },
-    "feedbackModal.provideReportModalTitle":
-    {
-        "defaultMessage": "Report This Content",
-        "description": "Title for the 'report' feedback modal"
-    },
-    "feedbackModal.reportContentExplanationPlaceholder":
-    {
-        "defaultMessage": "What is wrong with the response? What about this response is harmful? Please be as specific as possible, and add any details that are not present in the checkboxes below.",
-        "description": "Placeholder for textarea input when user chooses to report a shared chat"
-    },
-    "feedbackModal.reportOtherContent":
-    {
-        "defaultMessage": "I don't like this for some other reason (please describe)",
-        "description": "Label for Report Other Content checkbox"
-    },
-    "feedbackModal.sexualAbuse":
-    {
-        "defaultMessage": "This content contains sexual abuse",
-        "description": "Label for Sexual Abuse checkbox"
-    },
-    "feedbackModal.skipStep":
-    {
-        "defaultMessage": "Skip this step",
-        "description": "Button text for skipping comparison step"
-    },
-    "feedbackModal.submitFeedback":
-    {
-        "defaultMessage": "Submit feedback",
-        "description": "Button text for submitting the feedback"
-    },
-    "feedbackModal.submitReport":
-    {
-        "defaultMessage": "Submit report",
-        "description": "Button text for submitting a content-moderation report"
-    },
-    "feedbackModal.thumbsDownPlaceholder":
-    {
-        "defaultMessage": "What was the issue with the response? How could it be improved?",
-        "description": "Placeholder for textarea input when user chooses thumbs down"
-    },
-    "feedbackModal.thumbsUpPlaceholder":
-    {
-        "defaultMessage": "What do you like about the response?",
-        "description": "Placeholder for textarea input when user chooses thumbs up"
-    },
-    "fileUpload.codeInterpreterSessionTimeout":
-    {
-        "defaultMessage": "Code interpreter session expired",
-        "description": "Error message when code interpreter session expired"
-    },
-    "fileUpload.defaultCreateEntryError":
-    {
-        "defaultMessage": "Unable to upload {fileName}",
-        "description": "Error message when file upload fails"
-    },
-    "fileUpload.defaultDownloadLinkError":
-    {
-        "defaultMessage": "Failed to get upload status for {fileName}",
-        "description": "Error message when file download link fails"
-    },
-    "fileUpload.fileCorrupted":
-    {
-        "defaultMessage": "This file is corrupted. Please ensure the file is not corrupted and try again.",
-        "description": "Error message when an uploaded file that is corrupted."
-    },
-    "fileUpload.fileEmpty":
-    {
-        "defaultMessage": "No text could be extracted from this file.",
-        "description": "Error message when an uploaded file does not contain parsable text"
-    },
-    "fileUpload.fileEncrypted":
-    {
-        "defaultMessage": "This file is encrypted/requires a password to access. Please try again with an unencrypted file.",
-        "description": "Error message when an uploaded file that is encrypted."
-    },
-    "fileUpload.fileNotFound":
-    {
-        "defaultMessage": "File not found",
-        "description": "Error message when file was not found"
-    },
-    "fileUpload.fileTimedOut":
-    {
-        "defaultMessage": "File upload timed out. Please try again.",
-        "description": "Error message when file upload timed out"
-    },
-    "fileUpload.fileTooLarge":
-    {
-        "defaultMessage": "File is too large",
-        "description": "Error message when file is too large to upload"
-    },
-    "fileUpload.fileTooManyTokens":
-    {
-        "defaultMessage": "This file contains too much text content. Please try again with a smaller file.",
-        "description": "Error message when an uploaded file contains too many tokens/is too large."
-    },
-    "fileUpload.fileZeroBytes":
-    {
-        "defaultMessage": "File is empty",
-        "description": "Error message when file is zero bytes"
-    },
-    "fileUpload.overUserQuota":
-    {
-        "defaultMessage": "User quota exceeded",
-        "description": "Error message when user storage space (quote) has been exceeded"
-    },
-    "fileUpload.permissionError":
-    {
-        "defaultMessage": "Missing permission to access file",
-        "description": "Error message when user doesn't have permission to access a file"
-    },
-    "fileUpload.unknownError":
-    {
-        "defaultMessage": "Unknown error occurred",
-        "description": "Error message when file upload fails"
-    },
-    "filesModal.fileDownloadFailed":
-    {
-        "defaultMessage": "File download failed. Please try again.",
-        "description": "Error message when file download fails"
-    },
-    "gizmo.actionNeedsPrivacyPolicyURL":
-    {
-        "defaultMessage": "Public actions require valid privacy policy URLs. Click <fixlink>here</fixlink> to update.",
-        "description": "Error message when trying to publish action"
-    },
-    "gizmo.actions.blankExampleTitle":
-    {
-        "defaultMessage": "Blank Template",
-        "description": "Dropdown label for the blank template OpenAPI spec"
-    },
-    "gizmo.actions.examples":
-    {
-        "defaultMessage": "Examples",
-        "description": "Label of examples in GPT actions editor"
-    },
-    "gizmo.actions.petStoreExampleTitle":
-    {
-        "defaultMessage": "Pet Store (YAML)",
-        "description": "Dropdown label for the pet store OpenAPI example"
-    },
-    "gizmo.actions.weatherExampleTitle":
-    {
-        "defaultMessage": "Weather (JSON)",
-        "description": "Dropdown label for the weather OpenAPI example"
-    },
-    "gizmo.anyoneWithLink":
-    {
-        "defaultMessage": "Anyone with a link",
-        "description": "Privacy option for anyone with a link"
-    },
-    "gizmo.clearChat":
-    {
-        "defaultMessage": "Clear chat",
-        "description": "Clear chat button label"
-    },
-    "gizmo.confirmPublish":
-    {
-        "defaultMessage": "Confirm",
-        "description": "Message prompting you to confirm publication of an action"
-    },
-    "gizmo.copyLink":
-    {
-        "defaultMessage": "Copy link",
-        "description": "Menu item for copying link to GPT"
-    },
-    "gizmo.createActionLabel":
-    {
-        "defaultMessage": "Create new action",
-        "description": "Label for button to create a new action"
-    },
-    "gizmo.delete":
-    {
-        "defaultMessage": "Delete GPT",
-        "description": "Button label for deleting a GPT"
-    },
-    "gizmo.descriptionTooLong":
-    {
-        "defaultMessage": "GPT descriptions cannot be longer than {length} characters.",
-        "description": "Error message when description is too long"
-    },
-    "gizmo.disabledCustomActionsTooltip":
-    {
-        "defaultMessage": "Custom actions are disabled for your workspace. Contact your admin to enable them.",
-        "description": "Tooltip label when custom actions are are disabled"
-    },
-    "gizmo.discovery.createGPT":
-    {
-        "defaultMessage": "Create a GPT",
-        "description": "Label for create GPT button"
-    },
-    "gizmo.discovery.createGPTShort":
-    {
-        "defaultMessage": "Create",
-        "description": "Label for create GPT button on small screens"
-    },
-    "gizmo.discovery.empty":
-    {
-        "defaultMessage": "Nothing to discover",
-        "description": "Label for empty discovery page"
-    },
-    "gizmo.discovery.featured.description":
-    {
-        "defaultMessage": "Curated top picks from this week",
-        "description": "Description for featured section"
-    },
-    "gizmo.discovery.featured.title":
-    {
-        "defaultMessage": "Featured",
-        "description": "Title for featured section"
-    },
-    "gizmo.discovery.loadMore":
-    {
-        "defaultMessage": "Load more",
-        "description": "Button label for loading more GPTs"
-    },
-    "gizmo.discovery.madeByOpenAi.description":
-    {
-        "defaultMessage": "Most popular GPTs created by OpenAI",
-        "description": "Description for made by OpenAI section"
-    },
-    "gizmo.discovery.myGPTs":
-    {
-        "defaultMessage": "My GPTs",
-        "description": "Label for my GPTs button"
-    },
-    "gizmo.discovery.search":
-    {
-        "defaultMessage": "Search public GPTs",
-        "description": "Placeholder for search input"
-    },
-    "gizmo.discovery.search.empty":
-    {
-        "defaultMessage": "No results found",
-        "description": "Label for no search results"
-    },
-    "gizmo.discovery.search.recentlyUsed":
-    {
-        "defaultMessage": "Recently Used",
-        "description": "Label for recently used section in search"
-    },
-    "gizmo.discovery.search.results":
-    {
-        "defaultMessage": "Search Results",
-        "description": "Label for search results section in search"
-    },
-    "gizmo.discovery.thirdPartyGPTsDisabled":
-    {
-        "defaultMessage": "Your admin has blocked GPTs created outside {workspaceName}.",
-        "description": "Description for third party GPTs disabled"
-    },
-    "gizmo.discovery.trending.description":
-    {
-        "defaultMessage": "Most popular GPTs by our community",
-        "description": "Description for trending section"
-    },
-    "gizmo.discovery.trending.title":
-    {
-        "defaultMessage": "Trending",
-        "description": "Title for trending section"
-    },
-    "gizmo.discovery.workspace.description":
-    {
-        "defaultMessage": "GPTs created at {workspaceName}",
-        "description": "Description for workspace section"
-    },
-    "gizmo.displayNameRequiredHint":
-    {
-        "defaultMessage": "To make your GPT public, set up your builder profile.",
-        "description": "Message hinting that you have to setup a builder profile to publish"
-    },
-    "gizmo.draft":
-    {
-        "defaultMessage": "Draft",
-        "description": "Label for draft status"
-    },
-    "gizmo.errorSavingDraft":
-    {
-        "defaultMessage": "Error saving draft",
-        "description": "Error message when saving a draft fails"
-    },
-    "gizmo.explore":
-    {
-        "defaultMessage": "Explore",
-        "description": "Button that allows you to explore more GPTs"
-    },
-    "gizmo.hideFromSidebar":
-    {
-        "defaultMessage": "Hide from sidebar",
-        "description": "Whether to hide a gpt from sidebar"
-    },
-    "gizmo.instructionsTooLong":
-    {
-        "defaultMessage": "GPT instructions cannot be longer than {length} characters.",
-        "description": "Error message when instructions are too long"
-    },
-    "gizmo.keepInSidebar":
-    {
-        "defaultMessage": "Keep in sidebar",
-        "description": "Whether to keep a gpt in sidebar"
-    },
-    "gizmo.knowledgeExplanation":
-    {
-        "defaultMessage": "Additional files for this GPT to reference.",
-        "description": "Explainer text around what happens when your files upload."
-    },
-    "gizmo.knowledgeWarning":
-    {
-        "defaultMessage": "Conversations with your GPT may include file contents. Files can be downloaded when code interpreter is enabled.",
-        "description": "Explanation text for what happens when files are uploaded"
-    },
-    "gizmo.maxActionsReached":
-    {
-        "defaultMessage": "GPTs can have a maximum of {number} actions",
-        "description": "Message when maximum number of actions has been reached"
-    },
-    "gizmo.nameTooLong":
-    {
-        "defaultMessage": "GPT names cannot be longer than {length} characters.",
-        "description": "Error message when name is too long"
-    },
-    "gizmo.newChat":
-    {
-        "defaultMessage": "New chat",
-        "description": "New chat tooltip"
-    },
-    "gizmo.newGPT":
-    {
-        "defaultMessage": "New GPT",
-        "description": "Placeholder for new GPT name"
-    },
-    "gizmo.noGizmosFound":
-    {
-        "defaultMessage": "No GPTs found for this user",
-        "description": "Description when nothing found for user"
-    },
-    "gizmo.onlyMe":
-    {
-        "defaultMessage": "Only me",
-        "description": "Privacy option for only me"
-    },
-    "gizmo.privacyAnyoneWithLink":
-    {
-        "defaultMessage": "Only people with a link",
-        "description": "Description for a privacy setting of anyone with link"
-    },
-    "gizmo.privacyMarketplace":
-    {
-        "defaultMessage": "Public",
-        "description": "Description for a privacy setting of public"
-    },
-    "gizmo.privacyOnlyMe":
-    {
-        "defaultMessage": "Only me",
-        "description": "Description for a privacy setting of fully private"
-    },
-    "gizmo.privacyWorkspace":
-    {
-        "defaultMessage": "Anyone at {workspaceName}",
-        "description": "Description for a privacy setting of workspace"
-    },
-    "gizmo.public":
-    {
-        "defaultMessage": "Public",
-        "description": "Privacy option for public"
-    },
-    "gizmo.publicSharingHint":
-    {
-        "defaultMessage": "This GPT may appear in the GPT Store (coming soon)",
-        "description": "Message hinting at what public privacy settings do"
-    },
-    "gizmo.publishChanges":
-    {
-        "defaultMessage": "Update",
-        "description": "Button label for publishing changes"
-    },
-    "gizmo.publishTo":
-    {
-        "defaultMessage": "Publish to",
-        "description": "Label for above publish options"
-    },
-    "gizmo.published":
-    {
-        "defaultMessage": "Published",
-        "description": "Label for published status"
-    },
-    "gizmo.publishedTitle":
-    {
-        "defaultMessage": "Published!",
-        "description": "Popout title when GPT is published"
-    },
-    "gizmo.revertConfirm":
-    {
-        "defaultMessage": "Are you sure you want to revert to the last saved version?",
-        "description": "Confirmation message for reverting to last saved version"
-    },
-    "gizmo.revertMenuItem":
-    {
-        "defaultMessage": "Revert...",
-        "description": "Menu item for reverting to last saved version"
-    },
-    "gizmo.save":
-    {
-        "defaultMessage": "Save",
-        "description": "Button label for save (publishing)"
-    },
-    "gizmo.viewGPT":
-    {
-        "defaultMessage": "View GPT",
-        "description": "Button label for viewing GPT"
-    },
-    "gizmo.welcomeMessageTooLong":
-    {
-        "defaultMessage": "GPT welcome messages cannot be longer than {length} characters.",
-        "description": "Error message when welcome message is too long"
-    },
-    "gizmo.workspaceDisabledHint":
-    {
-        "defaultMessage": "Your workspace administrator has disabled this setting",
-        "description": "Message saying your workspace administrator isn't allowing something"
-    },
-    "gizmoGPTsOnboarding.headline":
-    {
-        "defaultMessage": "Introducing GPTs",
-        "description": "Headline for the GPTs onboarding modal"
-    },
-    "gizmoGPTsOnboarding.subheading":
-    {
-        "defaultMessage": "Custom versions of ChatGPT that combine instructions, extra knowledge and capabilities for a specific purpose.",
-        "description": "Subheading for the GPTs onboarding modal"
-    },
-    "gizmoGPTsOnboardingBiz.headline":
-    {
-        "defaultMessage": "Your templates are now GPTs, customized versions of ChatGPT.",
-        "description": "Headline for the biz GPTs onboarding modal"
-    },
-    "gizmoGPTsOnboardingBiz.subheading":
-    {
-        "defaultMessage": "GPTs use custom instructions, data, and capabilities to tailor ChatGPT to your needs.",
-        "description": "Subheading for the biz GPTs onboarding modal"
-    },
-    "gizmoOnboarding.headline":
-    {
-        "defaultMessage": "ChatGPT can now browse the web, analyze data, and generate images.",
-        "description": "Headline for the gizmo onboarding modal"
-    },
-    "gizmoOnboarding.learnMore":
+        "description": "Dropdown choice to bring up the verification modal"
+    },
+    "workspaceIdentity.ssoEditButton":
+    {
+        "defaultMessage": "Edit SAML SSO",
+        "description": "Label for edit SSO button"
+    },
+    "workspaceIdentity.ssoLearnMoreLink":
     {
         "defaultMessage": "Learn more",
-        "description": "Button to learn more about GPTs"
-    },
-    "gizmoOnboarding.next":
-    {
-        "defaultMessage": "Next",
-        "description": "Next button for the biz gizmo onboarding modal"
-    },
-    "gizmoOnboarding.start":
-    {
-        "defaultMessage": "Get started",
-        "description": "Start button for the gizmo onboarding modal"
-    },
-    "gizmoOnboarding.subheading":
-    {
-        "defaultMessage": "These capabilities are now built into GPT-4, no need to select them. Available for Plus users.",
-        "description": "Subheading for the gizmo onboarding modal"
-    },
-    "gizmoOnboarding.subheading.biz":
-    {
-        "defaultMessage": "These capabilities are now built into GPT-4, no need to select them. Available for Enterprise users.",
-        "description": "Subheading for the gizmo onboarding modal for Biz users"
-    },
-    "globalToasts.conversationInaccessible":
-    {
-        "defaultMessage": "Conversation inaccessible or not found",
-        "description": "Warning toast message when conversation cannot load"
-    },
-    "globalToasts.gizmoNotFound":
-    {
-        "defaultMessage": "GPT inaccessible or not found",
-        "description": "Warning toast message when GPT cannot load"
-    },
-    "globalToasts.noAccess":
-    {
-        "defaultMessage": "You do not currently have access to this feature",
-        "description": "Warning when you try to access a feature you don't have access to"
-    },
-    "globalToasts.oauthSuccess":
-    {
-        "defaultMessage": "You have successfully signed in via OAuth",
-        "description": "Success message when a user has logged into an app with oauth."
-    },
-    "history.bucket.lastSeven":
-    {
-        "defaultMessage": "Previous 7 Days",
-        "description": "Label for the history bucket of the previous 7 days"
-    },
-    "history.bucket.lastThirty":
-    {
-        "defaultMessage": "Previous 30 Days",
-        "description": "Label for the history bucket of the previous 30 days"
-    },
-    "history.bucket.today":
-    {
-        "defaultMessage": "Today",
-        "description": "Label for today's history bucket"
-    },
-    "history.bucket.yesterday":
-    {
-        "defaultMessage": "Yesterday",
-        "description": "Label for yesterday's history bucket"
-    },
-    "history.deleteModalBody":
-    {
-        "defaultMessage": "This will delete {title}.",
-        "description": "Body of the modal to confirm deleting a conversation"
-    },
-    "history.deleteModalCancel":
-    {
-        "defaultMessage": "Cancel",
-        "description": "Button to cancel deleting a conversation"
-    },
-    "history.deleteModalConfirm":
-    {
-        "defaultMessage": "Delete",
-        "description": "Button to confirm deleting a conversation"
-    },
-    "history.deleteModalTitle":
-    {
-        "defaultMessage": "Delete chat?",
-        "description": "Title of the modal to confirm deleting a conversation"
-    },
-    "history.retryButton":
-    {
-        "defaultMessage": "Retry",
-        "description": "Button to retry loading history"
-    },
-    "history.showMoreButton":
-    {
-        "defaultMessage": "Show more",
-        "description": "Button to show more history items"
-    },
-    "history.unableToLoad":
-    {
-        "defaultMessage": "Unable to load history",
-        "description": "Error message when history fails to load"
-    },
-    "imageViewer.closeModal":
-    {
-        "defaultMessage": "Close Modal",
-        "description": "Button to close the modal"
-    },
-    "imageViewer.downloadImage":
-    {
-        "defaultMessage": "Download Image",
-        "description": "Button to download the image"
-    },
-    "imageViewer.nextImage":
-    {
-        "defaultMessage": "Next Image",
-        "description": "Button to go to the next image"
-    },
-    "imageViewer.previousImage":
-    {
-        "defaultMessage": "Previous Image",
-        "description": "Button to go to the previous image"
-    },
-    "imageViewer.showImage":
-    {
-        "defaultMessage": "Show Image",
-        "description": "Button to show the image in a modal"
-    },
-    "imageViewer.toggleSidebar":
-    {
-        "defaultMessage": "Toggle Sidebar",
-        "description": "Button to toggle the sidebar"
-    },
-    "jitPluginMessage.allow":
-    {
-        "defaultMessage": "Allow",
-        "description": "Button text for the user to allow the custom action"
-    },
-    "jitPluginMessage.alwaysAllow":
-    {
-        "defaultMessage": "Always allow",
-        "description": "Button text for the user to always allow the custom action"
-    },
-    "jitPluginMessage.confirmParamsTitleV2":
-    {
-        "defaultMessage": "{gizmoName} needs to send this info to {domain}",
-        "description": "Title describing data that will be sent to the external website"
-    },
-    "jitPluginMessage.confirmingV2":
-    {
-        "defaultMessage": "<params>Some info will be sent to {domain}, only do this for sites you trust</params>",
-        "description": "Status message when a custom action is showing a user confirmation"
-    },
-    "jitPluginMessage.decline":
-    {
-        "defaultMessage": "Decline",
-        "description": "Button text for the user to decline the custom action"
-    },
-    "jitPluginMessage.declined":
-    {
-        "defaultMessage": "You declined this action",
-        "description": "Status message when a custom action was declined by the user"
-    },
-    "jitPluginMessage.errorV5":
-    {
-        "defaultMessage": "Error talking to {domain}",
-        "description": "Status message when a custom action ran into an error"
-    },
-    "jitPluginMessage.finishedV3":
-    {
-        "defaultMessage": "<params>Talked to {domain}</params>",
-        "description": "Status message when a custom action is finished"
-    },
-    "jitPluginMessage.privacyPolicyLinkV2":
-    {
-        "defaultMessage": "Privacy policy",
-        "description": "Text for the privacy policy link"
-    },
-    "jitPluginMessage.ranTest":
-    {
-        "defaultMessage": "Tested {operationName}",
-        "description": "Status message when the user launched a test action"
-    },
-    "jitPluginMessage.runningV4":
-    {
-        "defaultMessage": "Talking to {domain}",
-        "description": "Status message when a custom action is running"
-    },
-    "jitPluginMessage.sentParamsTitleV2":
-    {
-        "defaultMessage": "{gizmoName} sent this info to {domain}",
-        "description": "Title describing data that was sent to the external website"
-    },
-    "jitPluginMessage.signInButton":
-    {
-        "defaultMessage": "Sign in with {domain}",
-        "description": "Button text for the user to sign in with an external website"
-    },
-    "jitPluginMessage.signedIn":
-    {
-        "defaultMessage": "Signed in with {domain}",
-        "description": "Status text when the user successfully signed in with an external website"
-    },
-    "jitPluginMessage.starting":
-    {
-        "defaultMessage": "Starting action",
-        "description": "Status message when a custom action is starting"
-    },
-    "jitPluginMessage.stoppedV4":
-    {
-        "defaultMessage": "Stopped talking to {domain}",
-        "description": "Status message when a custom action was stopped by the user"
-    },
-    "keyboardActions.copyLastCodeBlock":
-    {
-        "defaultMessage": "Copy last code block",
-        "description": "Keyboard shortcut to copy the last code block in the chat"
-    },
-    "keyboardActions.copyLastResponse":
-    {
-        "defaultMessage": "Copy last response",
-        "description": "Keyboard shortcut to copy the last response in the chat"
-    },
-    "keyboardActions.deleteChat":
-    {
-        "defaultMessage": "Delete chat",
-        "description": "Keyboard shortcut to delete chat"
-    },
-    "keyboardActions.focusPromptTextarea":
-    {
-        "defaultMessage": "Focus chat input",
-        "description": "Keyboard shortcut to focus the chat input"
-    },
-    "keyboardActions.navigationToggle":
-    {
-        "defaultMessage": "Toggle sidebar",
-        "description": "Keyboard shortcut to toggle navigation"
-    },
-    "keyboardActions.newChat":
-    {
-        "defaultMessage": "Open new chat",
-        "description": "Keyboard shortcut to open a new chat"
-    },
-    "keyboardActions.toggleCustomInstructions":
-    {
-        "defaultMessage": "Set custom instructions",
-        "description": "Keyboard shortcut to toggle custom instructions"
-    },
-    "keyboardActions.toggleKeyboardActions":
-    {
-        "defaultMessage": "Show shortcuts",
-        "description": "Keyboard shortcut to toggle keyboard actions"
-    },
-    "leaveWorkspaceModal.cancel":
-    {
-        "defaultMessage": "Cancel",
-        "description": "Cancel button to close leave-workspace dialog"
-    },
-    "leaveWorkspaceModal.cantLeaveWorkspace":
-    {
-        "defaultMessage": "Couldn't leave the {workspaceName} workspace",
-        "description": "User tried to leave workspace but they're the last owner"
-    },
-    "leaveWorkspaceModal.done":
-    {
-        "defaultMessage": "Done",
-        "description": "Done button to close leave-workspace dialog"
-    },
-    "leaveWorkspaceModal.enterYourEmail":
-    {
-        "defaultMessage": "Enter your email address to confirm",
-        "description": "Enter your email"
-    },
-    "leaveWorkspaceModal.lastOwnerWarning":
-    {
-        "defaultMessage": "Because you're the only owner in the {workspaceName} workspace, assign the owner role to another member before leaving.",
-        "description": "Explaining why the user can't leave the workspace"
-    },
-    "leaveWorkspaceModal.leaveAreYouSure":
-    {
-        "defaultMessage": "Are you sure?",
-        "description": "Help text section title for leave workspace dialog"
-    },
-    "leaveWorkspaceModal.leaveButton":
-    {
-        "defaultMessage": "Leave workspace",
-        "description": "Button to confirm leaving the workspace"
-    },
-    "leaveWorkspaceModal.leaveFailed":
-    {
-        "defaultMessage": "Failed to leave workspace",
-        "description": "Error message when user fails to leave workspace"
-    },
-    "leaveWorkspaceModal.leaveWorkspace":
-    {
-        "defaultMessage": "Leave the {workspaceName} workspace",
-        "description": "Leave workspace dialog title, containing name of workspace that user will leave."
-    },
-    "leaveWorkspaceModal.leaveWorkspaceWarning1":
-    {
-        "defaultMessage": "This will remove you from your workspace and you won't be able to access all data, including profile, settings, and chat history.",
-        "description": "Help text for leaving workspace, first bullet point."
-    },
-    "leaveWorkspaceModal.leaveWorkspaceWarning2":
-    {
-        "defaultMessage": "You will lose access to all channels and messages in this workspace.",
-        "description": "Help text for leaving workspace, second bullet point."
-    },
-    "leaveWorkspaceModal.leftWorkspaceDescription":
-    {
-        "defaultMessage": "You have successfully left the {workspaceName} workspace.",
-        "description": "Description of the leave-workspace dialog after the user has left the workspace"
-    },
-    "leaveWorkspaceModal.leftWorkspaceDescriptionNoOtherWorkspaces":
-    {
-        "defaultMessage": "You have successfully left the {workspaceName} workspace. This will create your personal workspace automatically.",
-        "description": "Description of the leave-workspace dialog after the user has left the workspace"
-    },
-    "leaveWorkspaceModal.leftWorkspaceTitle":
-    {
-        "defaultMessage": "Successfully left the {workspaceName} workspace",
-        "description": "Title of the leave-workspace dialog after the user has left the workspace"
-    },
-    "leaveWorkspaceModal.memberCount":
-    {
-        "defaultMessage": "{memberCount, plural, one {1 member} other {{memberCount} members} }",
-        "description": "Number of members in workspace"
-    },
-    "leaveWorkspaceModal.ok":
-    {
-        "defaultMessage": "OK",
-        "description": "OK button to close leave-workspace dialog"
-    },
-    "leaveWorkspaceModal.startPersonalAccount":
-    {
-        "defaultMessage": "Start using ChatGPT for free",
-        "description": "Button to let the user know that they can start using a free personal account."
-    },
-    "message.gizmo.failed":
-    {
-        "defaultMessage": "Failed to generate profile picture.",
-        "description": "Message displayed when the GPT editor failed to generate a profile picture"
-    },
-    "message.gizmo.finished":
-    {
-        "defaultMessage": "GPT Behavior updated.",
-        "description": "Message displayed when the GPT editor is finished updating"
-    },
-    "message.gizmo.generatedProfilePic":
-    {
-        "defaultMessage": "Generated profile picture prompt.",
-        "description": "Message displayed when the GPT editor is finished generating a profile picture"
-    },
-    "message.gizmo.generatingProfilePic":
-    {
-        "defaultMessage": "Generating profile picture...",
-        "description": "Message displayed when the GPT editor is generating a profile picture"
-    },
-    "message.gizmo.updating":
-    {
-        "defaultMessage": "Updating GPT...",
-        "description": "Message displayed when the GPT editor is updating"
-    },
-    "modal.okayButton":
-    {
-        "defaultMessage": "Okay",
-        "description": "Confirmation button label in modal"
-    },
-    "modelCapMessaging.shortLimitDays":
-    {
-        "defaultMessage": "Limit {numerator, plural, =0 {# message} one {# message} other {# messages}} / {denominator, plural, =0 {# day} one {# day} other {# days}}",
-        "description": "Short message limit"
-    },
-    "navigation.accountSwitcherTitle":
-    {
-        "defaultMessage": "Workspaces",
-        "description": "Account switcher title"
-    },
-    "navigation.addWorkspaceTooltip":
-    {
-        "defaultMessage": "Create a Team workspace",
-        "description": "Tooltip for add workspace button"
-    },
-    "navigation.closeSidebar":
-    {
-        "defaultMessage": "Close sidebar",
-        "description": "Close sidebar button label"
-    },
-    "navigation.disabledWorkspaceTooltip":
-    {
-        "defaultMessage": "This workspace has been deactivated",
-        "description": "Tooltip for disabled workspace"
-    },
-    "navigation.enableChatHistory":
-    {
-        "defaultMessage": "Enable chat history",
-        "description": "Enable chat history button label"
-    },
-    "navigation.helpAndFaq":
-    {
-        "defaultMessage": "Help & FAQ",
-        "description": "Help & FAQ menu item"
-    },
-    "navigation.leaveWorkspace":
-    {
-        "defaultMessage": "Leave workspace",
-        "description": "Menu item in workplace switcher to leave a deactivated workspace"
-    },
-    "navigation.logOut":
-    {
-        "defaultMessage": "Log out",
-        "description": "Log out menu item"
-    },
-    "navigation.newChat":
-    {
-        "defaultMessage": "New chat",
-        "description": "New chat header title"
-    },
-    "navigation.openSidebar":
-    {
-        "defaultMessage": "Open sidebar",
-        "description": "Open sidebar button label"
-    },
-    "navigation.settings":
-    {
-        "defaultMessage": "Settings",
-        "description": "Settings menu item"
-    },
-    "navigation.settingsPlus":
-    {
-        "defaultMessage": "Settings & Beta",
-        "description": "Settings menu item for Plus users"
-    },
-    "navigation.survey.takeSurveyButton":
-    {
-        "defaultMessage": "Take survey",
-        "description": "Survey offer call to action"
-    },
-    "navigation.surveyDescription":
-    {
-        "defaultMessage": "Shape the future of ChatGPT.",
-        "description": "Survey offer description"
-    },
-    "navigation.surveyDismiss":
-    {
-        "defaultMessage": "Dismiss survey",
-        "description": "Survey offer dismiss button"
-    },
-    "navigation.surveyTitle":
-    {
-        "defaultMessage": "We’d love to hear from you!",
-        "description": "Survey offer title"
-    },
-    "onboarding.accuracy":
-    {
-        "defaultMessage": "Check your facts",
-        "description": "Title for the warning about ChatGPT inaccuracy"
-    },
-    "onboarding.accuracyBody":
-    {
-        "defaultMessage": "While we have safeguards, ChatGPT may give you inaccurate information. It’s not intended to give advice.",
-        "description": "Body copy for the warning about ChatGPT inaccuracy"
-    },
-    "onboarding.askAway":
-    {
-        "defaultMessage": "Ask away",
-        "description": "Title for the tip about what ChatGPT can do"
-    },
-    "onboarding.askAwayBody":
-    {
-        "defaultMessage": "ChatGPT can answer questions, help you learn, write code, brainstorm together, and much more.",
-        "description": "Body copy for the tip about what ChatGPT can do"
-    },
-    "onboarding.chatgptTitle.0":
-    {
-        "defaultMessage": "Welcome to the {workspaceName} workspace",
-        "description": "Title for the initial onboarding modal"
-    },
-    "onboarding.continueButton":
-    {
-        "defaultMessage": "Continue",
-        "description": "Label for the continue button"
-    },
-    "onboarding.departments.administrative":
-    {
-        "defaultMessage": "Administrative Assistant",
-        "description": "Department option for Administrative Assistant"
-    },
-    "onboarding.departments.analytics":
-    {
-        "defaultMessage": "Data or Analytics",
-        "description": "Department option for Data or Analytics"
-    },
-    "onboarding.departments.comms":
-    {
-        "defaultMessage": "Communications",
-        "description": "Department option for Communication"
-    },
-    "onboarding.departments.customer_experience":
-    {
-        "defaultMessage": "Customer Experience",
-        "description": "Department option for Customer Experience"
-    },
-    "onboarding.departments.design":
-    {
-        "defaultMessage": "Design",
-        "description": "Department option for Design"
-    },
-    "onboarding.departments.education_professional":
-    {
-        "defaultMessage": "Education Professional",
-        "description": "Department option for Education Professional"
-    },
-    "onboarding.departments.engineering":
-    {
-        "defaultMessage": "Engineering",
-        "description": "Department option for Engineering"
-    },
-    "onboarding.departments.finance":
-    {
-        "defaultMessage": "Finance or Accounting",
-        "description": "Department option for Finance or Accounting"
-    },
-    "onboarding.departments.healthcare":
-    {
-        "defaultMessage": "Healthcare Professional",
-        "description": "Department option for Healthcare Professional"
-    },
-    "onboarding.departments.human_resources":
-    {
-        "defaultMessage": "Human Resources",
-        "description": "Department option for Human Resources"
-    },
-    "onboarding.departments.it":
-    {
-        "defaultMessage": "Information Technology (IT)",
-        "description": "Department option for Information Technology (IT)"
-    },
-    "onboarding.departments.legal":
-    {
-        "defaultMessage": "Legal",
-        "description": "Department option for Legal"
-    },
-    "onboarding.departments.marketing":
-    {
-        "defaultMessage": "Marketing",
-        "description": "Department option for Marketing"
-    },
-    "onboarding.departments.ops":
-    {
-        "defaultMessage": "Operations",
-        "description": "Department option for Operations"
-    },
-    "onboarding.departments.other":
-    {
-        "defaultMessage": "Other",
-        "description": "Department option for Other"
-    },
-    "onboarding.departments.partnerships":
-    {
-        "defaultMessage": "Partnerships",
-        "description": "Department option for Partnerships"
-    },
-    "onboarding.departments.product":
-    {
-        "defaultMessage": "Product Management",
-        "description": "Department option for Product Management"
-    },
-    "onboarding.departments.project_management":
-    {
-        "defaultMessage": "Project or Program Management",
-        "description": "Department option for Project or Program Management"
-    },
-    "onboarding.departments.research":
-    {
-        "defaultMessage": "Research & Development",
-        "description": "Department option for Research & Development"
-    },
-    "onboarding.departments.sales":
-    {
-        "defaultMessage": "Sales",
-        "description": "Department option for Sales"
-    },
-    "onboarding.gettingStartedButton":
-    {
-        "defaultMessage": "Okay, let’s go",
-        "description": "Button to accept the getting started modal"
-    },
-    "onboarding.primaryRoleTitle":
-    {
-        "defaultMessage": "What's your primary role?",
-        "description": "Question asking the user about their primary role"
-    },
-    "onboarding.role.business_owner":
-    {
-        "defaultMessage": "Business Owner",
-        "description": "Label for the role: Business Owner"
-    },
-    "onboarding.role.director":
-    {
-        "defaultMessage": "Director",
-        "description": "Label for the role: Director"
-    },
-    "onboarding.role.executive":
-    {
-        "defaultMessage": "Executive",
-        "description": "Label for the role: Executive"
-    },
-    "onboarding.role.freelancer":
-    {
-        "defaultMessage": "Freelancer",
-        "description": "Label for the role: Freelancer"
-    },
-    "onboarding.role.manager":
-    {
-        "defaultMessage": "Manager",
-        "description": "Label for the role: Manager"
-    },
-    "onboarding.role.other":
-    {
-        "defaultMessage": "Other",
-        "description": "Label for the role: Other"
-    },
-    "onboarding.role.student":
-    {
-        "defaultMessage": "Student",
-        "description": "Label for the role: Student"
-    },
-    "onboarding.role.team_member":
-    {
-        "defaultMessage": "Team Member/ Individual Contributor",
-        "description": "Label for the role: Team Member/ Individual Contributor"
-    },
-    "onboarding.selectAll":
-    {
-        "defaultMessage": "Select all that apply",
-        "description": "Instruction for multi-select options"
-    },
-    "onboarding.skipButton":
-    {
-        "defaultMessage": "Skip",
-        "description": "Label for the skip button"
-    },
-    "onboarding.tailorChatGPT":
-    {
-        "defaultMessage": "This will help us tailor ChatGPT for you.",
-        "description": "Description explaining the reason for the questions"
-    },
-    "onboarding.warning":
-    {
-        "defaultMessage": "Don’t share sensitive info",
-        "description": "Title for the warning about ChatGPT traning"
-    },
-    "onboarding.warningBody":
-    {
-        "defaultMessage": "Chat history may be reviewed or used to improve our services. Learn more about your choices in our <article>Help Center</article>.",
-        "description": "Body copy for the warning about ChatGPT traning"
-    },
-    "onboarding.workTypeTitle":
-    {
-        "defaultMessage": "What kind of work do you do?",
-        "description": "Question asking the user about the kind of work they do"
-    },
-    "onboarding.workUse":
-    {
-        "defaultMessage": "Made for use at work",
-        "description": "Title for the warning about ChatGPT business workspace use"
-    },
-    "onboarding.workUseBody":
-    {
-        "defaultMessage": "By default, chats in this workspace are not used to train our AI models.",
-        "description": "Body copy for the warning about ChatGPT business workspace use"
-    },
-    "onboarding.workspaceWelcome":
-    {
-        "defaultMessage": "Welcome to the {workspaceName} Workspace",
-        "description": "Introduction welcome for workspace onboarding"
-    },
-    "onboarding.workspaceWelcomeBody":
-    {
-        "defaultMessage": "Here you can use our latest models, with more capabilities, and fewer limits.",
-        "description": "Introduction welcome body for workspace onboarding"
-    },
-    "onboarding.workspaceWelcomeNoName":
-    {
-        "defaultMessage": "Welcome to the your Workspace",
-        "description": "Introduction welcome for workspace onboarding when no Workspace name is available"
-    },
-    "organizationBillingInfo.activeLicense":
-    {
-        "defaultMessage": "Team License",
-        "description": "Active team license name"
-    },
-    "organizationBillingInfo.cancelSubscriptionBtn":
-    {
-        "defaultMessage": "Cancel subscription",
-        "description": "Cancel subscription button"
-    },
-    "organizationBillingInfo.defaultInvoiceName":
-    {
-        "defaultMessage": "Invoice",
-        "description": "invoice name when we are missing dates"
-    },
-    "organizationBillingInfo.inactiveLicense":
-    {
-        "defaultMessage": "Inactive License",
-        "description": "Inactive license name"
-    },
-    "organizationBillingInfo.invoiceName":
-    {
-        "defaultMessage": "Invoice: {createdDate, date, long}",
-        "description": "invoice name"
-    },
-    "organizationBillingInfo.invoicesLoadError.0":
-    {
-        "defaultMessage": "Failed to load invoices. Contact support@openai.com if error persists.",
-        "description": "Error message when invoices fail to load"
-    },
-    "organizationBillingInfo.invoicesTitle":
-    {
-        "defaultMessage": "Invoices",
-        "description": "Title for the organization invoices"
-    },
-    "organizationBillingInfo.licenseExpiry":
-    {
-        "defaultMessage": "Active until {expiryDate, date, long}",
-        "description": "License expiry date"
-    },
-    "organizationBillingInfo.managePaymentMethodBtn":
-    {
-        "defaultMessage": "Manage payment method",
-        "description": "Manage payment method button"
-    },
-    "organizationBillingInfo.manageSubscription":
-    {
-        "defaultMessage": "Manage subscription",
-        "description": "Manage subscription header"
-    },
-    "organizationBillingInfo.noInvoices":
-    {
-        "defaultMessage": "No invoices found",
-        "description": "Error message when no invoices are found"
-    },
-    "organizationBillingInfo.paidInvoice":
-    {
-        "defaultMessage": "Paid: {invoiceDate, date, long}",
-        "description": "Paid invoice date"
-    },
-    "organizationBillingInfo.planTitle":
-    {
-        "defaultMessage": "Plan",
-        "description": "Title for the organization billing plan"
-    },
-    "organizationBillingInfo.renewalDate":
-    {
-        "defaultMessage": "Deactivates on {expiryDate, date, long}",
-        "description": "Subscription renewal date"
-    },
-    "organizationBillingInfo.seatsInUse":
-    {
-        "defaultMessage": "{numSeats} in use ({numSeatsPct})",
-        "description": "Number of seats in use"
-    },
-    "organizationBillingInfo.seatsTitle":
-    {
-        "defaultMessage": "Seats",
-        "description": "Title for the organization billing seats"
-    },
-    "organizationBillingInfo.stripeErrorWarning":
-    {
-        "defaultMessage": "Error loading account information",
-        "description": "Error message for loading stripe account information"
-    },
-    "organizationBillingInfo.subtitle":
-    {
-        "defaultMessage": "Only workspace owners can view and change these settings",
-        "description": "Subtitle for the organization billing info page"
-    },
-    "organizationBillingInfo.teamsAutochargeMessage":
-    {
-        "defaultMessage": "Your additional seats will be included on your next invoice.",
-        "description": "Message to inform owners they will be charged for additional seats"
-    },
-    "organizationBillingInfo.title.1":
-    {
-        "defaultMessage": "Billing",
-        "description": "Title for the organization billing info page"
-    },
-    "organizationBillingInfo.upcomingInvoice":
-    {
-        "defaultMessage": "Due: {invoiceDate, date, long}",
-        "description": "Upcoming invoice date"
-    },
-    "parallelBrowsingMessage.almostDone":
-    {
-        "defaultMessage": "Almost done",
-        "description": "Status message when browsing is almost done visiting sites"
-    },
-    "parallelBrowsingMessage.cancelledV2":
-    {
-        "defaultMessage": "Stopped doing research",
-        "description": "Status message when browsing was cancelled"
-    },
-    "parallelBrowsingMessage.creatingPlanV2":
-    {
-        "defaultMessage": "Making a research plan",
-        "description": "Status message when browsing is being planned"
-    },
-    "parallelBrowsingMessage.running_2":
-    {
-        "defaultMessage": "Visiting {numTasks, plural, one {# site} other {# sites}}",
-        "description": "Status message when browsing is visiting sites"
-    },
-    "personalizationSettings.customInstructions":
-    {
-        "defaultMessage": "Custom instructions",
-        "description": "Custom instructions settings label"
-    },
-    "personalizationSettings.off":
-    {
-        "defaultMessage": "Off",
-        "description": "Off"
-    },
-    "personalizationSettings.on":
-    {
-        "defaultMessage": "On",
-        "description": "On"
-    },
-    "pluginDisplayParams.generatedImage":
-    {
-        "defaultMessage": "Generated by plugin",
-        "description": "Description text for an image that was generated by a plugin"
-    },
-    "popoverNavigation.chatPreferences":
+        "description": "Label for link to learn more"
+    },
+    "workspaceIdentity.ssoTitle":
+    {
+        "defaultMessage": "Single Sign On",
+        "description": "Section title of SSO"
+    },
+    "workspaceIdentity.ssoUrlExtractFailed":
+    {
+        "defaultMessage": "Unable to find Entity ID (sign-in endpoint) in provided XML",
+        "description": "Notice to the user that the given XML did not contain an entity ID"
+    },
+    "workspaceIdentity.statusTableHeader":
+    {
+        "defaultMessage": "Status",
+        "description": "Label for status column on domains table"
+    },
+    "workspaceIdentity.title":
+    {
+        "defaultMessage": "Identity & Provisioning",
+        "description": "Title of identity page"
+    },
+    "workspaceIdentity.toggleAutoProvisionLabel":
+    {
+        "defaultMessage": "Toggle to enable or disable auto provisioning",
+        "description": "Label for auto provisioning toggle"
+    },
+    "workspaceIdentity.verifyDomainInfoText":
+    {
+        "defaultMessage": "Verify ownership of an email domain to access advanced security features including single sign-on.",
+        "description": "Informational text for domain verification section"
+    },
+    "workspacePopoverNavigation.chatPreferences":
     {
         "defaultMessage": "Custom instructions",
         "description": "Custom instructions menu item"
     },
-    "popoverNavigation.myGpts":
+    "workspacePopoverNavigation.myGpts":
     {
         "defaultMessage": "My GPTs",
         "description": "My GPTs menu item"
     },
-    "popoverNavigation.myPlan":
-    {
-        "defaultMessage": "My plan",
-        "description": "My plan menu item"
-    },
-    "pricingPlanConstants.business.callToAction":
-    {
-        "defaultMessage": "Buy for my team",
-        "description": "Business purchase call to action"
-    },
-    "pricingPlanConstants.free.callToAction":
-    {
-        "defaultMessage": "Your current plan",
-        "description": "Call to action for free plan"
-    },
-    "pricingPlanConstants.free.costInDollars":
-    {
-        "defaultMessage": "USD $0/month",
-        "description": "Cost for free plan"
-    },
-    "pricingPlanConstants.free.demandAccess":
-    {
-        "defaultMessage": "Access to our GPT-3.5 model",
-        "description": "Access rights for free plan"
-    },
-    "pricingPlanConstants.free.freeAdvertisedFeatures0":
-    {
-        "defaultMessage": "Unlimited messages, interactions, and history",
-        "description": "Free plan feature message"
-    },
-    "pricingPlanConstants.free.freeAdvertisedFeatures1":
-    {
-        "defaultMessage": "Access to our GPT-3.5 model",
-        "description": "Free plan feature message"
-    },
-    "pricingPlanConstants.free.freeAdvertisedFeatures2":
-    {
-        "defaultMessage": "Access on Web, iOS, and Android",
-        "description": "Free plan feature message"
-    },
-    "pricingPlanConstants.free.freePlanForLine":
-    {
-        "defaultMessage": "For people just getting started with ChatGPT",
-        "description": "Free plan subtitle message"
-    },
-    "pricingPlanConstants.free.modelFeatures":
-    {
-        "defaultMessage": "Regular model updates",
-        "description": "Model features for free plan"
-    },
-    "pricingPlanConstants.free.name":
-    {
-        "defaultMessage": "Free",
-        "description": "Name of the free pricing plan"
-    },
-    "pricingPlanConstants.free.responseSpeed":
-    {
-        "defaultMessage": "Standard response speed",
-        "description": "Response speed for free plan"
-    },
-    "pricingPlanConstants.getHelp.callToAction":
-    {
-        "defaultMessage": "I need help with a billing issue",
-        "description": "Help for billing issues"
-    },
-    "pricingPlanConstants.highDemandDisabledText":
-    {
-        "defaultMessage": "Due to high demand, we've temporarily paused upgrades.",
-        "description": "Message shown when demand is too high and payments are disabled"
-    },
-    "pricingPlanConstants.manageSubscriptionAndroid.callToAction":
-    {
-        "defaultMessage": "Manage my subscription in the ChatGPT Android app",
-        "description": "Android subscription management"
-    },
-    "pricingPlanConstants.manageSubscriptionIos.callToAction":
-    {
-        "defaultMessage": "Manage my subscription in the ChatGPT iOS app",
-        "description": "iOS subscription management"
-    },
-    "pricingPlanConstants.manageSubscriptionWeb.callToAction":
-    {
-        "defaultMessage": "Manage my subscription",
-        "description": "Web subscription management"
-    },
-    "pricingPlanConstants.plus.callToAction.active":
-    {
-        "defaultMessage": "Your current plan",
-        "description": "Active call to action for plus plan"
-    },
-    "pricingPlanConstants.plus.callToAction.inactivePayment":
-    {
-        "defaultMessage": "Upgrade to Plus",
-        "description": "Inactive payment call to action for plus plan"
-    },
-    "pricingPlanConstants.plus.costInDollars":
-    {
-        "defaultMessage": "USD $20/month",
-        "description": "Cost for plus plan"
-    },
-    "pricingPlanConstants.plus.demandAccess":
-    {
-        "defaultMessage": "Access to GPT-4, our most capable model",
-        "description": "Access rights for plus plan"
-    },
-    "pricingPlanConstants.plus.forLine":
-    {
-        "defaultMessage": "Everything in Free, and:",
-        "description": "Plus plan for line"
-    },
-    "pricingPlanConstants.plus.modelFeatures":
-    {
-        "defaultMessage": "Access to beta features like browsing, plugins, and advanced data analysis",
-        "description": "Model features for plus plan"
-    },
-    "pricingPlanConstants.plus.name":
-    {
-        "defaultMessage": "Plus",
-        "description": "Name of the ChatGPT Plus pricing plan"
-    },
-    "pricingPlanConstants.plus.plusAdvertisedFeatures0":
-    {
-        "defaultMessage": "Access to GPT-4, our most capable model",
-        "description": "Plus plan feature message"
-    },
-    "pricingPlanConstants.plus.plusAdvertisedFeatures1":
-    {
-        "defaultMessage": "Browse, create, and use GPTs",
-        "description": "Plus plan feature message"
-    },
-    "pricingPlanConstants.plus.plusAdvertisedFeatures2":
-    {
-        "defaultMessage": "Access to additional tools like DALL\\xb7E, Browsing, Advanced Data Analysis and more",
-        "description": "Plus plan feature message"
-    },
-    "pricingPlanConstants.plus.responseSpeed":
-    {
-        "defaultMessage": "Faster response speed",
-        "description": "Response speed for plus plan"
-    },
-    "pricingPlanConstants.plusWaitlistSignupSuccess":
-    {
-        "defaultMessage": "You've been added to the waitlist to upgrade to Plus",
-        "description": "Plus waitlist signup success message"
-    },
-    "pricingPlanConstants.signUpForWaitlistActive":
-    {
-        "defaultMessage": "Sign up for waitlist",
-        "description": "Sign up for waitlist call to action"
-    },
-    "pricingPlanConstants.signUpForWaitlistInactive":
-    {
-        "defaultMessage": "Signed up for waitlist",
-        "description": "Signed up for waitlist call to action"
-    },
-    "pricingPlanConstants.teamWaitlistSignupSuccess":
-    {
-        "defaultMessage": "You've been added to the waitlist to upgrade to Team",
-        "description": "Team waitlist signup success message"
-    },
-    "pricingPlanConstants.teams.teamPlanActive":
-    {
-        "defaultMessage": "Your current plan",
-        "description": "Upgrade to team message"
-    },
-    "pricingPlanConstants.teams.teamPlanContext":
-    {
-        "defaultMessage": "4x longer context lets you work with larger material",
-        "description": "Team plan feature message"
-    },
-    "pricingPlanConstants.teams.teamPlanCreate":
-    {
-        "defaultMessage": "Create Team",
-        "description": "Create a team message"
-    },
-    "pricingPlanConstants.teams.teamPlanForLine":
-    {
-        "defaultMessage": "Everything in Plus, and:",
-        "description": "Team plan subtitle message"
-    },
-    "pricingPlanConstants.teams.teamPlanInactive":
-    {
-        "defaultMessage": "Upgrade to Team",
-        "description": "Upgrade to team message"
-    },
-    "pricingPlanConstants.teams.teamPlanName":
-    {
-        "defaultMessage": "Team",
-        "description": "Team plan title message"
-    },
-    "pricingPlanConstants.teams.teamPlanSubTitle":
-    {
-        "defaultMessage": "USD $25 per person/month*",
-        "description": "Team plan subtitle message"
-    },
-    "pricingPlanConstants.teams.teamPlanUsageRates":
-    {
-        "defaultMessage": "Unlimited high-speed GPT-4",
-        "description": "Team plan feature message"
-    },
-    "pricingPlanConstants.teams.teamPricingDisclaimer":
-    {
-        "defaultMessage": "* Price billed annually, minimum 2 users",
-        "description": "Team plan disclaimer message"
-    },
-    "pricingPlanConstants.teams.teamsAdvertisedFeatures0":
-    {
-        "defaultMessage": "Expanded access to GPT-4 and tools like DALL\\xb7E, Browsing, Advanced Data Analysis, and more",
-        "description": "Team plan feature message"
-    },
-    "pricingPlanConstants.teams.teamsAdvertisedFeatures1":
-    {
-        "defaultMessage": "Create and share GPTs with your workspace",
-        "description": "Team plan feature message"
-    },
-    "pricingPlanConstants.teams.teamsAdvertisedFeatures2":
-    {
-        "defaultMessage": "Admin console for workspace management",
-        "description": "Team plan feature message"
-    },
-    "pricingPlanConstants.teams.teamsAdvertisedFeatures3":
-    {
-        "defaultMessage": "No training on your data",
-        "description": "Team plan feature message"
-    },
-    "rating.instructions":
-    {
-        "defaultMessage": "Is this conversation helpful so far?",
-        "description": "Ask the user for their rating of the conversation so far"
-    },
-    "rating.thanks":
-    {
-        "defaultMessage": "Thanks for your feedback!",
-        "description": "Thank the user for their rating"
-    },
-    "settingsModal.addDomain":
-    {
-        "defaultMessage": "Verify new domain",
-        "description": "Add domain button text"
-    },
-    "settingsModal.apiAccessDeletionWarning-2":
-    {
-        "defaultMessage": "Deletion will prevent you from accessing OpenAI services, including ChatGPT, API, and DALL\\xb7E.",
-        "description": "Warning message about API access being deleted."
-    },
-    "settingsModal.betaAdvancedDataAnalysisToggleDescription":
-    {
-        "defaultMessage": "Try a version of ChatGPT that knows how to write and execute python code, and can work with file uploads. Try asking for help with data analysis, image conversions, or editing a code file. Note: files will not persist beyond a single session.",
-        "description": "Description for the Advanced data analysis beta toggle."
-    },
-    "settingsModal.betaAdvancedDataAnalysisToggleLabel":
-    {
-        "defaultMessage": "Advanced data analysis",
-        "description": "Label for the Advanced data analysis beta toggle."
-    },
-    "settingsModal.betaBrowsingToggleDescription":
-    {
-        "defaultMessage": "Try a version of ChatGPT that knows when and how to browse the internet to answer questions about recent topics and events.",
-        "description": "Description for the Browsing beta toggle."
-    },
-    "settingsModal.betaBrowsingToggleLabel":
-    {
-        "defaultMessage": "Browse with Bing",
-        "description": "Label for the Browse with Bing beta toggle."
-    },
-    "settingsModal.betaIntro":
-    {
-        "defaultMessage": "As a Plus user, enjoy early access to experimental new features, which may change during development.",
-        "description": "Introduction for the beta features tab"
-    },
-    "settingsModal.betaPluginToggleDescription":
-    {
-        "defaultMessage": "Try a version of ChatGPT that knows when and how to use third-party plugins that you enable.",
-        "description": "Description for the Plugins beta toggle."
-    },
-    "settingsModal.betaPluginToggleLabel":
-    {
-        "defaultMessage": "Plugins",
-        "description": "Label for the Plugins beta toggle."
-    },
-    "settingsModal.betaSettingsUpdateFailed":
-    {
-        "defaultMessage": "Failed to update your beta setting",
-        "description": "Message shown when there's an error updating beta settings"
-    },
-    "settingsModal.betaTab":
-    {
-        "defaultMessage": "Beta features",
-        "description": "Label for the Beta Features tab"
-    },
-    "settingsModal.byName":
-    {
-        "defaultMessage": "by {name}",
-        "description": "By creator name under GPT"
-    },
-    "settingsModal.chatHistoryDescription":
-    {
-        "defaultMessage": "Save new chats on this browser to your history and allow them to be used to improve our models. Unsaved chats will be deleted from our systems within 30 days. This setting does not sync across browsers or devices. <link>Learn more</link>",
-        "description": "Description for the chat history setting"
-    },
-    "settingsModal.chatHistoryLabel":
-    {
-        "defaultMessage": "Chat history",
-        "description": "Label for the chat history toggle."
-    },
-    "settingsModal.chatHistoryOnlyDescription":
-    {
-        "defaultMessage": "Save new chats on this browser to your history. Unsaved chats will be deleted from our systems within 30 days. This setting does not sync across browsers or devices. <link>Learn more</link>",
-        "description": "Description for the chat history setting"
-    },
-    "settingsModal.chatHistoryToggleLabel":
-    {
-        "defaultMessage": "Chat history & training",
-        "description": "Label for the chat history toggle."
-    },
-    "settingsModal.chatTrainingEnterpriseDescription":
-    {
-        "defaultMessage": "This workspace is private and opted out of training.",
-        "description": "Description for the disabled chat training toggle."
-    },
-    "settingsModal.chatTrainingEnterpriseTooltip":
-    {
-        "defaultMessage": "ChatGPT Team automatically disables training.",
-        "description": "Tooltip for the disabled chat training toggle."
-    },
-    "settingsModal.chatTrainingLabel":
-    {
-        "defaultMessage": "Chat training",
-        "description": "Label for the chat training toggle."
-    },
-    "settingsModal.clearChatButton":
-    {
-        "defaultMessage": "Clear",
-        "description": "Clear chat button"
-    },
-    "settingsModal.clearChatLabel":
-    {
-        "defaultMessage": "Clear all chats",
-        "description": "Label for the clear chat button"
-    },
-    "settingsModal.connectorsTab":
-    {
-        "defaultMessage": "Connected apps",
-        "description": "Label for the Connected Apps tab"
-    },
-    "settingsModal.cookieManagement":
-    {
-        "defaultMessage": "Cookie preferences",
-        "description": "Label for the cookie preferences button"
-    },
-    "settingsModal.cookieManagementButton":
-    {
-        "defaultMessage": "Manage",
-        "description": "Manage cookie management button"
-    },
-    "settingsModal.cookies":
-    {
-        "defaultMessage": "Cookie Preferences",
-        "description": "Label for the cookies tab"
-    },
-    "settingsModal.creatorProfileDescription":
-    {
-        "defaultMessage": "Personalize your builder profile to connect with users of your GPTs. These settings apply to publicly shared GPTs.",
-        "description": "Description for the builder  Profile tab"
-    },
-    "settingsModal.creatorProfileLinkHeader.0":
-    {
-        "defaultMessage": "Website",
-        "description": "Links header on builder Profile tab"
-    },
-    "settingsModal.creatorProfileNameLabel":
-    {
-        "defaultMessage": "Name",
-        "description": "Label for the Name field on builder Profile tab."
-    },
-    "settingsModal.dark":
-    {
-        "defaultMessage": "Dark",
-        "description": "Option for the dark theme"
-    },
-    "settingsModal.dataControls":
-    {
-        "defaultMessage": "Data controls",
-        "description": "Label for the data controls tab"
-    },
-    "settingsModal.dataExportFailed":
-    {
-        "defaultMessage": "We were unable to process your export at this time. Please try again later.",
-        "description": "Message shown when a data export request fails"
-    },
-    "settingsModal.dataExportModalCancel":
-    {
-        "defaultMessage": "Cancel",
-        "description": "Cancel button for the data export modal"
-    },
-    "settingsModal.dataExportModalConfirm":
-    {
-        "defaultMessage": "Confirm export",
-        "description": "Confirm button for the data export modal"
-    },
-    "settingsModal.dataExportModalDescription1":
-    {
-        "defaultMessage": "Your account details and conversations will be included in the export.",
-        "description": "Description for the data export modal"
-    },
-    "settingsModal.dataExportModalDescription2":
-    {
-        "defaultMessage": "The data will be sent to your registered email in a downloadable file.",
-        "description": "Description for the data export modal"
-    },
-    "settingsModal.dataExportModalDescription3":
-    {
-        "defaultMessage": "The download link will expire 24 hours after you receive it.",
-        "description": "Description for the data export modal"
-    },
-    "settingsModal.dataExportModalDescription4":
-    {
-        "defaultMessage": "Processing may take some time. You'll be notified when it's ready.",
-        "description": "Description for the data export modal"
-    },
-    "settingsModal.dataExportModalDescription5":
-    {
-        "defaultMessage": "To proceed, click \"Confirm export\" below.",
-        "description": "Description for the data export modal"
-    },
-    "settingsModal.dataExportModalTitle":
-    {
-        "defaultMessage": "Request data export - are you sure?",
-        "description": "Title for the data export modal"
-    },
-    "settingsModal.dataExportRequested":
-    {
-        "defaultMessage": "Successfully exported data. You should receive an email shortly with your data.",
-        "description": "Message shown when a data export request is received"
-    },
-    "settingsModal.dataRemovalWarning-2":
-    {
-        "defaultMessage": "Your data will be deleted within 30 days, except we may retain a limited set of data for longer where required or permitted by law.",
-        "description": "Warning message about data removal after account deletion."
-    },
-    "settingsModal.deleteAccount":
-    {
-        "defaultMessage": "Delete account",
-        "description": "Label for the delete account button"
-    },
-    "settingsModal.deleteAccountButtonLabel":
-    {
-        "defaultMessage": "Permanently delete my account",
-        "description": "Label for the button to confirm account deletion."
-    },
-    "settingsModal.deleteAccountFailed":
-    {
-        "defaultMessage": "Failed to delete account. Please try again later.",
-        "description": "Message shown when there's an error deleting the user's account."
-    },
-    "settingsModal.deleteAccountSessionTooOld":
-    {
-        "defaultMessage": "Your login session is too old. Please log in again before deleting your account.",
-        "description": "Message shown when the user's login session is too old to delete their account."
-    },
-    "settingsModal.deleteAccountTitle":
-    {
-        "defaultMessage": "Delete account - are you sure?",
-        "description": "Title for the delete account confirmation modal."
-    },
-    "settingsModal.deleteAccountWarning":
-    {
-        "defaultMessage": "Deleting your account is permanent and cannot be undone.",
-        "description": "Warning message about account deletion being permanent."
-    },
-    "settingsModal.deleteButton":
-    {
-        "defaultMessage": "Delete",
-        "description": "Delete account button"
-    },
-    "settingsModal.deleteHelpCenter":
-    {
-        "defaultMessage": "Read our <article>help center article</article> for more information.",
-        "description": "Line that directs users to the help article for more information."
-    },
-    "settingsModal.deleteHistoryModalCancel":
-    {
-        "defaultMessage": "Cancel",
-        "description": "Cancel button for the delete history modal"
-    },
-    "settingsModal.deleteHistoryModalConfirm":
-    {
-        "defaultMessage": "Confirm deletion",
-        "description": "Confirm button for the delete history modal"
-    },
-    "settingsModal.deleteHistoryModalTitle":
-    {
-        "defaultMessage": "Clear your conversation history - are you sure?",
-        "description": "Title for the delete history modal"
-    },
-    "settingsModal.disable":
-    {
-        "defaultMessage": "Disable",
-        "description": "Disable 2FA button"
-    },
-    "settingsModal.disable2fa":
-    {
-        "defaultMessage": "Disable two factor authentication",
-        "description": "Label for the mfa remove button."
-    },
-    "settingsModal.enable":
-    {
-        "defaultMessage": "Enable",
-        "description": "Enable 2FA button"
-    },
-    "settingsModal.enable2fa":
-    {
-        "defaultMessage": "Enable two-factor authentication",
-        "description": "Label for the enable 2FA button"
-    },
-    "settingsModal.exampleDescription1":
-    {
-        "defaultMessage": "Your primary GPT will continually improve as you chat, picking up on details and preferences to tailor its responses to you. <link>Learn more</link>",
-        "description": "Description line 1 for the MyChatgpt example"
-    },
-    "settingsModal.exampleDescription2":
-    {
-        "defaultMessage": "To modify what your GPT knows or teach it something new, just tell it:",
-        "description": "Description line 2 for the MyChatgpt example"
-    },
-    "settingsModal.exampleMessage1":
-    {
-        "defaultMessage": "What do you know about me?",
-        "description": "Message 1 for the MyChatgpt example"
-    },
-    "settingsModal.exampleMessage2":
-    {
-        "defaultMessage": "Where did we leave off on my last project?",
-        "description": "Message 2 for the MyChatgpt example"
-    },
-    "settingsModal.exampleMessage3":
-    {
-        "defaultMessage": "Forget about my travel plans.",
-        "description": "Message 3 for the MyChatgpt example"
-    },
-    "settingsModal.exportButton":
-    {
-        "defaultMessage": "Export",
-        "description": "Export data button"
-    },
-    "settingsModal.exportData":
-    {
-        "defaultMessage": "Export data",
-        "description": "Label for the export data button"
-    },
-    "settingsModal.generalTab":
-    {
-        "defaultMessage": "General",
-        "description": "Label for the general tab"
-    },
-    "settingsModal.gizmoTab":
-    {
-        "defaultMessage": "Builder profile",
-        "description": "Label for the builder profile tab"
-    },
-    "settingsModal.hideNameToggle":
-    {
-        "defaultMessage": "Hide your name in your builder profile",
-        "description": "Toggle label for hiding name"
-    },
-    "settingsModal.hideWebsiteToggle":
-    {
-        "defaultMessage": "Hide your website in your builder profile",
-        "description": "Toggle label for hiding website"
-    },
-    "settingsModal.iapSubscriptionWarning":
-    {
-        "defaultMessage": "You will need to cancel your in-app purchase subscription in the Apple App Store. We cannot cancel your subscription for you.",
-        "description": "Warning message about cancelling in-app subscriptions."
-    },
-    "settingsModal.light":
-    {
-        "defaultMessage": "Light",
-        "description": "Option for the light theme"
-    },
-    "settingsModal.linkDisabledTooltip":
-    {
-        "defaultMessage": "You must have a verified domain to enable displaying a link",
-        "description": "Tooltip for disabled link toggle"
-    },
-    "settingsModal.localeAuto":
-    {
-        "defaultMessage": "Auto-detect",
-        "description": "Label for the auto-detect locale setting"
-    },
-    "settingsModal.localeDev":
-    {
-        "defaultMessage": "⋆✩★DEV★✩⋆",
-        "description": "Label for the dev locale setting"
-    },
-    "settingsModal.locale_alpha":
-    {
-        "defaultMessage": "Locale (Alpha)",
-        "description": "Label for the locale setting"
-    },
-    "settingsModal.lockedButtonLabel":
-    {
-        "defaultMessage": "Locked",
-        "description": "Label for the locked button when deleting an account."
-    },
-    "settingsModal.myChagtGptToggleLabel":
-    {
-        "defaultMessage": "Improve responses with your chats",
-        "description": "Label for the chat history toggle."
-    },
-    "settingsModal.nameDisabledTooltip":
-    {
-        "defaultMessage": "You must have a verified name to enable displaying a name",
-        "description": "Tooltip for disabled name toggle"
-    },
-    "settingsModal.nameSourceReason":
-    {
-        "defaultMessage": "Name is populated from your billing details",
-        "description": "Reason for name source"
-    },
-    "settingsModal.noCreatorProfile":
-    {
-        "defaultMessage": "Unable to retrieve builder profile",
-        "description": "No builder profile error message"
-    },
-    "settingsModal.noVerifiedName":
-    {
-        "defaultMessage": "No verified name",
-        "description": "Text when user does not have a verified name to display"
-    },
-    "settingsModal.openPluginDevtools":
-    {
-        "defaultMessage": "Open plugin devtools",
-        "description": "Label for the open plugin devtools setting"
-    },
-    "settingsModal.personalization":
-    {
-        "defaultMessage": "Personalization",
-        "description": "Label for the Personalization tab"
-    },
-    "settingsModal.placeholderGPT":
-    {
-        "defaultMessage": "PlaceholderGPT",
-        "description": "Placeholder for the GPT preview on builder Profile tab."
-    },
-    "settingsModal.playStoreSubscriptionWarning":
-    {
-        "defaultMessage": "You will need to cancel your in-app purchase subscription in the Google Play Store. We cannot cancel your subscription for you.",
-        "description": "Warning message about cancelling in-app subscriptions."
-    },
-    "settingsModal.preview":
-    {
-        "defaultMessage": "Preview",
-        "description": "Preview tag in builder profile GPT preview"
-    },
-    "settingsModal.recentLoginMessage":
-    {
-        "defaultMessage": "You may only delete your account if you have logged in within the last 10 minutes. Please log in again, then return here to continue.",
-        "description": "Message shown when the user needs to log in again to delete their account."
-    },
-    "settingsModal.refreshLoginButtonLabel":
-    {
-        "defaultMessage": "Refresh login",
-        "description": "Label for the button to refresh login."
-    },
-    "settingsModal.reset":
-    {
-        "defaultMessage": "Reset",
-        "description": "Label for the Reset button in MyChatGPT"
-    },
-    "settingsModal.resetModalCancel":
-    {
-        "defaultMessage": "Cancel",
-        "description": "Cancel button for the reset memory modal"
-    },
-    "settingsModal.resetModalConfirm":
-    {
-        "defaultMessage": "Confirm reset",
-        "description": "Confirm button for the reset memory modal"
-    },
-    "settingsModal.resetModalDescription":
-    {
-        "defaultMessage": "Your primary GPT will forget what it has learned from your previous chats. This can't be undone.",
-        "description": "Description for the reset memory modal"
-    },
-    "settingsModal.resetModalTitle":
-    {
-        "defaultMessage": "Are you sure?",
-        "description": "Title for the reset memory modal"
-    },
-    "settingsModal.reuseEmailPhoneWarning-2":
-    {
-        "defaultMessage": "You cannot create a new account using the same email address.",
-        "description": "Warning message about not being able to reuse email for a new account."
-    },
-    "settingsModal.settings":
-    {
-        "defaultMessage": "Settings",
-        "description": "Title for the settings modal"
-    },
-    "settingsModal.sharedConversations":
-    {
-        "defaultMessage": "Shared links",
-        "description": "Label for the shared chat/link button"
-    },
-    "settingsModal.sharedConversationsButton":
-    {
-        "defaultMessage": "Manage",
-        "description": "Manage shared links/conversations button"
-    },
-    "settingsModal.showNameToggle":
-    {
-        "defaultMessage": "Show your name in your builder profile",
-        "description": "Toggle label for showing name"
-    },
-    "settingsModal.showWebsiteToggle":
-    {
-        "defaultMessage": "Show your website in your builder profile",
-        "description": "Toggle label for showing website"
-    },
-    "settingsModal.system":
-    {
-        "defaultMessage": "System",
-        "description": "Option for the system theme"
-    },
-    "settingsModal.theme":
-    {
-        "defaultMessage": "Theme",
-        "description": "Label for the theme setting"
-    },
-    "settingsModal.tryItOut":
-    {
-        "defaultMessage": "Message ChatGPT",
-        "description": "Placeholder text in the MyChatgpt example button"
-    },
-    "settingsModal.typeDeleteInputLabel":
-    {
-        "defaultMessage": "To proceed, type \"DELETE\" in the input field below.",
-        "description": "Label for DELETE input field when deleting an account."
-    },
-    "settingsModal.typeEmailLabel":
-    {
-        "defaultMessage": "Please type your account email.",
-        "description": "Label for email input field when deleting an account."
-    },
-    "settingsModal.websiteLinkTitle":
-    {
-        "defaultMessage": "Website",
-        "description": "Website link title on builder Profile tab"
-    },
-    "shared.pagination.back":
-    {
-        "defaultMessage": "Back",
-        "description": "Text for the \"Back\" button in Pagination component."
-    },
-    "shared.pagination.done":
-    {
-        "defaultMessage": "Done",
-        "description": "Text for the \"Done\" button in Pagination component."
-    },
-    "shared.pagination.next":
-    {
-        "defaultMessage": "Next",
-        "description": "Text for the \"Next\" button in Pagination component."
-    },
-    "sharedConversation.advancedDataAnalysisSupportDisclaimer":
-    {
-        "defaultMessage": "This chat contains files or images produced by Advanced Data Analysis which are not yet visible in Shared Chats.",
-        "description": "Disclaimer about our lack of support for Advanced Data Analysis inline images and file downloads with shared links"
-    },
-    "sharedConversation.personalizedDataDisclaimer":
-    {
-        "defaultMessage": "This conversation may reflect the link creator’s personalized data, which isn’t shared and can meaningfully change how the model responds.",
-        "description": "Disclaimer about the creator's personalized data (ex: custom instructions, memory) not being part of the shared conversation"
-    },
-    "sharedConversationModal.dateShared":
-    {
-        "defaultMessage": "Date shared",
-        "description": "Table column header"
-    },
-    "sharedConversationModal.deleteSharedAllConversations":
-    {
-        "defaultMessage": "Delete all shared links",
-        "description": "Menu item for deleting all shared links"
-    },
-    "sharedConversationModal.deleteSharedAllConversationsConfirm":
-    {
-        "defaultMessage": "Are you sure you want to delete all your shared links?",
-        "description": "Confirmation message for deleting share links"
-    },
-    "sharedConversationModal.deleteSharedAllConversationsFailed":
-    {
-        "defaultMessage": "Deleting shared links failed",
-        "description": "Toaster message when deleting all share links fails"
-    },
-    "sharedConversationModal.deleteSharedLink":
-    {
-        "defaultMessage": "Delete shared link",
-        "description": "Label for delete shared link icon"
-    },
-    "sharedConversationModal.deleteSharedLinkFailed":
-    {
-        "defaultMessage": "Deleting shared link failed",
-        "description": "Toaster message when deleting share link fails"
-    },
-    "sharedConversationModal.goToOriginConversation":
-    {
-        "defaultMessage": "View source chat",
-        "description": "Label for conversation icon"
-    },
-    "sharedConversationModal.loading":
-    {
-        "defaultMessage": "Loading...",
-        "description": "Loading message"
-    },
-    "sharedConversationModal.name":
-    {
-        "defaultMessage": "Name",
-        "description": "Table column header"
-    },
-    "sharedConversationModal.noSharedConversations":
-    {
-        "defaultMessage": "You have no shared links.",
-        "description": "No shared links message"
-    },
-    "sharedConversationModal.retry":
-    {
-        "defaultMessage": "Retry",
-        "description": "Retry button text"
-    },
-    "sharedConversationModal.somethingWentWrong":
-    {
-        "defaultMessage": "Something went wrong...",
-        "description": "Error message"
-    },
-    "sharedConversationModal.title":
-    {
-        "defaultMessage": "Shared Links",
-        "description": "Shared links modal title"
-    },
-    "sharingModal.advancedDataAnalysisSupportDisclaimer":
-    {
-        "defaultMessage": "Recipients won’t be able to view Advanced Data Analysis images or download files.",
-        "description": "Disclaimer about our lack of support for Advanced Data Analysis inline images and file downloads with shared links"
-    },
-    "sharingModal.bizDescription.1":
-    {
-        "defaultMessage": "Only members of your workspace with the URL will see the latest messages sent in this conversation. Files you attach to the conversation will not be shared, but any file contents referenced in messages will continue to be visible.",
-        "description": "Description of sharing feature in the first paragraph of the sharing modal"
-    },
-    "sharingModal.confirmDeleteLink":
-    {
-        "defaultMessage": "Are you sure you want to delete the share link?",
-        "description": "Confirmation message when deleting share link"
-    },
-    "sharingModal.deleteLink":
-    {
-        "defaultMessage": "Delete Link",
-        "description": "Button text to delete the share link"
-    },
-    "sharingModal.description":
-    {
-        "defaultMessage": "Messages you send after creating your link won't be shared. Anyone with the URL will be able to view the shared chat.",
-        "description": "Description of sharing feature in the first paragraph of the sharing modal"
-    },
-    "sharingModal.exisitingDescription":
-    {
-        "defaultMessage": "You have shared this chat <existingLink>before</existingLink>. If you want to update the shared chat content, <deleteLink>delete this link</deleteLink> and create a new shared link.",
-        "description": "Description in sharing modal when viewing an existing link"
-    },
-    "sharingModal.moderationBlocked":
-    {
-        "defaultMessage": "This shared link has been disabled by moderation.",
-        "description": "Error message in sharing modal when shared link has been moderated."
-    },
-    "sharingModal.personalizedDataDisclaimer":
-    {
-        "defaultMessage": "Any personalized data not present in the conversation won’t be shared with viewers (ex: custom instructions).",
-        "description": "Disclaimer about our policy to not share personalized data (ex: custom instructions, memory)"
-    },
-    "sharingModal.shareAnonymously":
-    {
-        "defaultMessage": "Share anonymously",
-        "description": "Button text to change sharing to be anonymous"
-    },
-    "sharingModal.shareYourName":
-    {
-        "defaultMessage": "Share your name",
-        "description": "Button text to change sharing to show the user's name"
-    },
-    "ssoModal.cancelButton":
-    {
-        "defaultMessage": "Cancel",
-        "description": "The text for the cancel button on the SSO modal"
-    },
-    "ssoModal.certLabel":
-    {
-        "defaultMessage": "X.509 Signing Certificate",
-        "description": "The label for the X.509 signing certificate input"
-    },
-    "ssoModal.createInstructionsManual":
-    {
-        "defaultMessage": "Copy your Sign-in endpoint (SSO URL) and the public X.509 certificate from your Identity Provider.",
-        "description": "Instructions for users to create an SSO connection from their IDP"
-    },
-    "ssoModal.createInstructionsXML":
-    {
-        "defaultMessage": "If your Identity Provider offers a metadata URL or an XML file, add it here for the quickest setup option.",
-        "description": "Instructions for users to create an SSO connection from an XML configuration that the IDP provides"
-    },
-    "ssoModal.createTitle":
-    {
-        "defaultMessage": "Create SAML SSO Configuration",
-        "description": "The title for the create version of the SSO modal"
-    },
-    "ssoModal.deleteButton":
-    {
-        "defaultMessage": "Delete configuration",
-        "description": "The text for the delete button on the SSO modal"
-    },
-    "ssoModal.doneButton":
-    {
-        "defaultMessage": "Done",
-        "description": "The text for the done button on the SSO modal"
-    },
-    "ssoModal.editTitle":
-    {
-        "defaultMessage": "Edit SAML SSO Configuration",
-        "description": "The title for the edit version of the SSO modal"
-    },
-    "ssoModal.nextButton":
-    {
-        "defaultMessage": "Next",
-        "description": "The text for the next button on the SSO modal"
-    },
-    "ssoModal.signInLabel":
-    {
-        "defaultMessage": "SSO URL",
-        "description": "The label for the sign in endpoint input"
-    },
-    "ssoModal.ssoMetadataLabel":
-    {
-        "defaultMessage": "App Federation Metadata URL",
-        "description": "The label for the IDP XML metadata URL input"
-    },
-    "targetedReply.replyTooltip":
-    {
-        "defaultMessage": "Reply",
-        "description": "Tooltip text for the targeted reply button"
-    },
-    "teamBilling.annualPlan.addUsersWithRenewal":
-    {
-        "defaultMessage": "Add users as needed, remove users only when you renew the contract",
-        "description": "Feature about adding/removing users in the Annual plan"
-    },
-    "teamBilling.annualPlan.billedAnnually":
-    {
-        "defaultMessage": "Annual price billed annually",
-        "description": "Feature indicating the Annual plan is billed once a year"
-    },
-    "teamBilling.annualPlan.cost":
-    {
-        "defaultMessage": "USD $25 <s>$30</s>",
-        "description": "Cost for the Annual billing plan"
-    },
-    "teamBilling.annualPlan.minBill":
-    {
-        "defaultMessage": "The minimum bill is for 2 users, USD $50/month billed annually",
-        "description": "Minimum bill details for the Annual plan"
-    },
-    "teamBilling.annualPlan.name":
-    {
-        "defaultMessage": "Annual plan",
-        "description": "Name for the Annual billing plan"
-    },
-    "teamBilling.annualPlanBilled":
-    {
-        "defaultMessage": "Price billed annually, starting {date}",
-        "description": "Summary of billing of selected plan"
-    },
-    "teamBilling.annualPlanSelected":
-    {
-        "defaultMessage": "ChatGPT Team annual plan",
-        "description": "Summary title of selected plan"
-    },
-    "teamBilling.annualPlanTotal":
-    {
-        "defaultMessage": "USD ${totalCost} per year total",
-        "description": "Summary total of selected plan"
-    },
-    "teamBilling.annualSavingsPercentage":
-    {
-        "defaultMessage": "16.7% off",
-        "description": "The savings percentage annual has on flexible"
-    },
-    "teamBilling.flexiblePlan.addRemoveUsers":
-    {
-        "defaultMessage": "Add or remove users as needed",
-        "description": "Feature indicating users can be added or removed in the Flexible plan"
-    },
-    "teamBilling.flexiblePlan.billedMonthly":
-    {
-        "defaultMessage": "Price billed monthly",
-        "description": "Feature indicating the Flexible plan is billed monthly"
-    },
-    "teamBilling.flexiblePlan.cost":
-    {
-        "defaultMessage": "USD $30",
-        "description": "Cost for the Flexible billing plan"
-    },
-    "teamBilling.flexiblePlan.minBill":
-    {
-        "defaultMessage": "The minimum monthly bill is for 2 users, USD $60/month",
-        "description": "Minimum bill details for the Flexible plan"
-    },
-    "teamBilling.flexiblePlan.name":
-    {
-        "defaultMessage": "Flexible plan",
-        "description": "Name for the Flexible billing plan"
-    },
-    "teamBilling.flexiblePlanBilled":
-    {
-        "defaultMessage": "Price billed monthly, starting {date}",
-        "description": "Summary billing of selected plan"
-    },
-    "teamBilling.flexiblePlanSelected":
-    {
-        "defaultMessage": "ChatGPT Team flexible plan",
-        "description": "Summary title of selected plan"
-    },
-    "teamBilling.flexiblePlanTotal":
-    {
-        "defaultMessage": "USD ${totalCost} per month total",
-        "description": "Summary total of selected plan"
-    },
-    "teamBilling.seatsTitle":
-    {
-        "defaultMessage": "Seats",
-        "description": "Seats section title"
-    },
-    "teamBilling.summaryTitle":
-    {
-        "defaultMessage": "Summary",
-        "description": "Summary section title"
-    },
-    "teamBilling.teamsCostStructure":
-    {
-        "defaultMessage": "per person/month",
-        "description": "The cost structure for teams plan"
-    },
-    "textMessage.errorLoadingImage":
-    {
-        "defaultMessage": "Could not load image",
-        "description": "Text that describes an image that failed to load"
-    },
-    "textMessage.imageAltText":
-    {
-        "defaultMessage": "Uploaded image",
-        "description": "Alt text for image asset"
-    },
-    "textMessage.loadingImage":
-    {
-        "defaultMessage": "Loading...",
-        "description": "Text that describes a loading image"
-    },
-    "textMessage.targetedReply":
-    {
-        "defaultMessage": "Replying to:",
-        "description": "Header shown above a targeted reply"
-    },
-    "thread.businessDisclaimer-oct-30":
-    {
-        "defaultMessage": "{workspaceName} workspace chats aren't used to train our models. ChatGPT can make mistakes.",
-        "description": "Business disclaimer with protected data assurance"
-    },
-    "thread.businessDisclaimerNoName-oct-30":
-    {
-        "defaultMessage": "Your workspace chats aren'ts used to train our models. ChatGPT can make mistakes.",
-        "description": "Business disclaimer with protected data assurance when no Workspace name is available"
-    },
-    "thread.chatgptMayProduceInaccurateInformation-oct-30":
-    {
-        "defaultMessage": "ChatGPT can make mistakes. Consider checking important information.",
-        "description": "ChatGPT disclaimer for producing inaccurate information"
-    },
-    "thread.helpAndFaq":
-    {
-        "defaultMessage": "Help & FAQ",
-        "description": "Help & FAQ menu item"
-    },
-    "thread.keyboardShortcutsMenu":
-    {
-        "defaultMessage": "Keyboard shortcuts",
-        "description": "Keyboard shortcuts menu item"
-    },
-    "thread.latencyButton":
-    {
-        "defaultMessage": "Latency",
-        "description": "Button to open the latency menu"
-    },
-    "thread.modal.onboarding.title":
-    {
-        "defaultMessage": "Do not share sensitive materials with this application",
-        "description": "Title for the onboarding warning modal"
-    },
-    "thread.modal.reportModalThankYou.description":
-    {
-        "defaultMessage": "Thank you for your report.",
-        "description": "Description for the post-report thank-you modal"
-    },
-    "thread.modal.reportModalThankYou.dismissButton":
-    {
-        "defaultMessage": "Close",
-        "description": "Close button for the post-report thank-you modal"
-    },
-    "thread.modal.reportModalThankYou.title":
-    {
-        "defaultMessage": "Thank you for your report!",
-        "description": "Title for the post-report thank-you modal"
-    },
-    "thread.modal.unrecoverableError.description":
-    {
-        "defaultMessage": "We're sorry, but something went wrong. Please try again later.",
-        "description": "Description for the UnrecoverableErrorModal"
-    },
-    "thread.modal.unrecoverableError.resetThread":
-    {
-        "defaultMessage": "Reset thread",
-        "description": "Reset thread button text"
-    },
-    "thread.modal.unrecoverableError.title":
-    {
-        "defaultMessage": "Something went wrong",
-        "description": "Title for the UnrecoverableErrorModal"
-    },
-    "thread.outdatedGptDisclaimer.0":
-    {
-        "defaultMessage": "<bold>New version of GPT available</bold> - Continue chatting to use the old version, or start a <link>new chat</link> for the latest version.",
-        "description": "Outdated GPT disclaimer"
-    },
-    "thread.outdatedTemplateDisclaimer":
-    {
-        "defaultMessage": "<bold>New version of template available</bold> - Continue chatting to use the old version, or start a <link>new chat</link> for the latest version.",
-        "description": "Outdated template disclaimer"
-    },
-    "thread.privacyPolicy":
-    {
-        "defaultMessage": "Privacy policy",
-        "description": "Privacy policy footer link text"
-    },
-    "thread.releaseNotes":
-    {
-        "defaultMessage": "Release notes",
-        "description": "Release notes menu item"
-    },
-    "thread.reportSharedConversation":
-    {
-        "defaultMessage": "Report content",
-        "description": "Report shared chat footer link text"
-    },
-    "thread.sharedConversation.continue":
-    {
-        "defaultMessage": "Continue this conversation",
-        "description": "Button for shared links to allow user to continue conversation in their own history"
-    },
-    "thread.sharedConversation.moderate":
-    {
-        "defaultMessage": "Moderate conversation",
-        "description": "Button for shared links to moderate a chat for legal, safety, or other reasons"
-    },
-    "thread.sharedConversation.report":
-    {
-        "defaultMessage": "Report conversation",
-        "description": "Button for shared links to report chat for legal, safety, or other reasons"
-    },
-    "thread.sharingModal.confirmCloseWithChanges":
-    {
-        "defaultMessage": "You have unsaved changes. Do you want to continue?",
-        "description": "Confirmation message when closing share modal with changes"
-    },
-    "thread.sharingModal.copied":
-    {
-        "defaultMessage": "Copied!",
-        "description": "Status message after successfully copying the shared link"
-    },
-    "thread.sharingModal.copiedSharedConversationURL":
-    {
-        "defaultMessage": "Copied shared conversation URL to clipboard!",
-        "description": "Success message when shared conversation URL is copied"
-    },
-    "thread.sharingModal.copyLink":
-    {
-        "defaultMessage": "Copy Link",
-        "description": "Button text to copy the shared link"
-    },
-    "thread.sharingModal.copying":
-    {
-        "defaultMessage": "Copying...",
-        "description": "Status message while copying the shared link"
-    },
-    "thread.sharingModal.failedToCopyLink":
-    {
-        "defaultMessage": "Failed to copy link to clipboard",
-        "description": "Error message when failing to copy link to clipboard"
-    },
-    "thread.sharingModal.failedToDeleteSharedLink":
-    {
-        "defaultMessage": "Failed to delete shared link",
-        "description": "Error message when failing to delete shared link"
-    },
-    "thread.sharingModal.moreInfo":
-    {
-        "defaultMessage": "More Info",
-        "description": "Link to a helpdesk article with more information about the sharing modal"
-    },
-    "thread.sharingModal.title":
-    {
-        "defaultMessage": "Share link to Chat",
-        "description": "Title of sharing feature in the title of the sharing modal"
-    },
-    "thread.sharingModal.updateAndCopyLink":
-    {
-        "defaultMessage": "Update and Copy Link",
-        "description": "Button text to update and copy the shared link"
-    },
-    "thread.termsAndPolicies":
-    {
-        "defaultMessage": "Terms & policies",
-        "description": "Terms & Policies menu item"
-    },
-    "thread.termsOfUse":
-    {
-        "defaultMessage": "Terms of use",
-        "description": "Terms of use footer link text"
-    },
-    "toolsUtils.browsingSearchLinkPrefix1":
-    {
-        "defaultMessage": "Based on a [quick search]({searchLink}), here's what I found.",
-        "description": "Prefix added to the message to link to a Bing search result"
-    },
-    "toolsUtils.browsingSearchLinkPrefix10":
-    {
-        "defaultMessage": "I did a [quick search]({searchLink}) for more information and here's what I discovered.",
-        "description": "Prefix added to the message to link to a Bing search result"
-    },
-    "toolsUtils.browsingSearchLinkPrefix2":
-    {
-        "defaultMessage": "After a [quick search]({searchLink}), here's what I found.",
-        "description": "Prefix added to the message to link to a Bing search result"
-    },
-    "toolsUtils.browsingSearchLinkPrefix3":
-    {
-        "defaultMessage": "From a [quick search]({searchLink}), here's what I found.",
-        "description": "Prefix added to the message to link to a Bing search result"
-    },
-    "toolsUtils.browsingSearchLinkPrefix4":
-    {
-        "defaultMessage": "I did a [quick search]({searchLink}) and here's what I found.",
-        "description": "Prefix added to the message to link to a Bing search result"
-    },
-    "toolsUtils.browsingSearchLinkPrefix5":
-    {
-        "defaultMessage": "I did a [quick search]({searchLink}) for more information and here's what I found.",
-        "description": "Prefix added to the message to link to a Bing search result"
-    },
-    "toolsUtils.browsingSearchLinkPrefix6":
-    {
-        "defaultMessage": "Based on a [quick search]({searchLink}), here's what I discovered.",
-        "description": "Prefix added to the message to link to a Bing search result"
-    },
-    "toolsUtils.browsingSearchLinkPrefix7":
-    {
-        "defaultMessage": "After a [quick search]({searchLink}), here's what I discovered.",
-        "description": "Prefix added to the message to link to a Bing search result"
-    },
-    "toolsUtils.browsingSearchLinkPrefix8":
-    {
-        "defaultMessage": "From a [quick search]({searchLink}), here's what I discovered.",
-        "description": "Prefix added to the message to link to a Bing search result"
-    },
-    "toolsUtils.browsingSearchLinkPrefix9":
-    {
-        "defaultMessage": "I did a [quick search]({searchLink}) and here's what I discovered.",
-        "description": "Prefix added to the message to link to a Bing search result"
-    },
-    "useFilePickerState.maxUploadsAtATime":
-    {
-        "defaultMessage": "Unable to upload {fileName}. Max {maxUploads} uploads at a time",
-        "description": "Error message when user tries to upload more than the max number of files"
-    },
-    "useFilePickerState.retrievalSkippedFile":
-    {
-        "defaultMessage": "Unable to extract text from \"{fileName}\"",
-        "description": "Error message when user uploads a file that we are unable to parse and extract text from."
-    },
-    "useModelSwitcherModels.gpt3_5.disclaimer":
-    {
-        "defaultMessage": "Available to Free and Plus users",
-        "description": "Disclaimer message for GPT-3.5, indicating availability to Free and Plus users"
-    },
-    "useModelSwitcherModels.gpt4.disclaimer":
-    {
-        "defaultMessage": "Available exclusively to Plus users",
-        "description": "Disclaimer message for GPT-4, indicating exclusive availability to Plus users"
-    },
-    "useSubscriptionData.subscriptionLoadError":
-    {
-        "defaultMessage": "Failed to load subscription: {error}. Contact support@openai.com if error persists.",
-        "description": "Error message when subscription fails to load"
-    },
-    "useWorkspaces.adminRoleName":
-    {
-        "defaultMessage": "Admin",
-        "description": "Role name for an admin user"
-    },
-    "useWorkspaces.defaultWorkspaceTitle":
-    {
-        "defaultMessage": "Untitled Workspace",
-        "description": "title for workspace without a name"
-    },
-    "useWorkspaces.enterprisePlanName":
-    {
-        "defaultMessage": "Enterprise",
-        "description": "label for enterprise tier account"
-    },
-    "useWorkspaces.ownerRoleName":
-    {
-        "defaultMessage": "Owner",
-        "description": "Role name for an owner user"
-    },
-    "useWorkspaces.personalPlanName":
-    {
-        "defaultMessage": "Personal",
-        "description": "label for personal tier account"
-    },
-    "useWorkspaces.personalWorkspaceTitle":
-    {
-        "defaultMessage": "Personal account",
-        "description": "title for personal workspace"
-    },
-    "useWorkspaces.teamPlanName":
-    {
-        "defaultMessage": "Team",
-        "description": "label for team tier account"
-    },
-    "useWorkspacews.standardRoleName":
-    {
-        "defaultMessage": "Member",
-        "description": "Role name for a standard user"
-    },
-    "userContextModal.aboutUserTip1":
-    {
-        "defaultMessage": "Where are you based?",
-        "description": "tips for Custom instructions about you"
-    },
-    "userContextModal.aboutUserTip2":
-    {
-        "defaultMessage": "What do you do for work?",
-        "description": "tips for Custom instructions about you"
-    },
-    "userContextModal.aboutUserTip3":
-    {
-        "defaultMessage": "What are your hobbies and interests?",
-        "description": "tips for Custom instructions about you"
-    },
-    "userContextModal.aboutUserTip4":
-    {
-        "defaultMessage": "What subjects can you talk about for hours?",
-        "description": "tips for Custom instructions about you"
-    },
-    "userContextModal.aboutUserTip5":
-    {
-        "defaultMessage": "What are some goals you have?",
-        "description": "tips for Custom instructions about you"
-    },
-    "userContextModal.aboutYouHelpText":
-    {
-        "defaultMessage": "What would you like ChatGPT to know about you to provide better responses?",
-        "description": "help text for about you section of Custom instructions"
-    },
-    "userContextModal.cancel":
-    {
-        "defaultMessage": "Cancel",
-        "description": "Cancel button for Custom instructions modal"
-    },
-    "userContextModal.chatPreferencesEnable":
-    {
-        "defaultMessage": "Enable for new chats",
-        "description": "chat preferences is enabled"
-    },
-    "userContextModal.confirmCloseBody":
-    {
-        "defaultMessage": "Are you sure you want to exit? Any changes you made will be permanently lost.",
-        "description": "confirm close modal"
-    },
-    "userContextModal.confirmCloseCancel":
-    {
-        "defaultMessage": "Back",
-        "description": "cancel button for confirm close modal"
-    },
-    "userContextModal.confirmCloseOk":
-    {
-        "defaultMessage": "Exit",
-        "description": "ok button for confirm close modal"
-    },
-    "userContextModal.confirmCloseTitle":
-    {
-        "defaultMessage": "You have unsaved changes.",
-        "description": "title for confirm close modal"
-    },
-    "userContextModal.disableToggleLabel":
-    {
-        "defaultMessage": "Disable chat preferences",
-        "description": "disable chat preferences toggle label"
-    },
-    "userContextModal.enableToggleLabel":
-    {
-        "defaultMessage": "Enable chat preferences",
-        "description": "enable chat preferences toggle label"
-    },
-    "userContextModal.hideTips":
-    {
-        "defaultMessage": "Hide tips",
-        "description": "hide tips button for Custom instructions modal"
-    },
-    "userContextModal.messageLimitError":
-    {
-        "defaultMessage": "Please limit your responses to {limit} characters or less.",
-        "description": "error message for Custom instructions modal"
-    },
-    "userContextModal.modApiVoilation":
-    {
-        "defaultMessage": "This content may violate our <policyLink>content policy</policyLink>. If you believe this to be in error, please <feedbackLink>submit your feedback</feedbackLink> — your input will aid our research in this area.",
-        "description": "error message for mod api voilation"
-    },
-    "userContextModal.modelHelpText":
-    {
-        "defaultMessage": "How would you like ChatGPT to respond?",
-        "description": "help text for about you section of Custom instructions"
-    },
-    "userContextModal.modelTip1":
-    {
-        "defaultMessage": "How formal or casual should ChatGPT be?",
-        "description": "tips for Custom instructions about model"
-    },
-    "userContextModal.modelTip2":
-    {
-        "defaultMessage": "How long or short should responses generally be?",
-        "description": "tips for Custom instructions about model"
-    },
-    "userContextModal.modelTip3":
-    {
-        "defaultMessage": "How do you want to be addressed?",
-        "description": "tips for Custom instructions about model"
-    },
-    "userContextModal.modelTip4":
-    {
-        "defaultMessage": "Should ChatGPT have opinions on topics or remain neutral?",
-        "description": "tips for Custom instructions about model"
-    },
-    "userContextModal.save":
-    {
-        "defaultMessage": "Save",
-        "description": "save button for my profile modal"
-    },
-    "userContextModal.showTips":
-    {
-        "defaultMessage": "Show tips",
-        "description": "show tips button for Custom instructions modal"
-    },
-    "userContextModal.subhead":
-    {
-        "defaultMessage": "<article>Learn more</article> about Custom instructions and how they’re used to help ChatGPT provide better responses.",
-        "description": "subhead for Custom instructions modal"
-    },
-    "userContextModal.tipsHeader":
-    {
-        "defaultMessage": "Thought starters",
-        "description": "header for Custom instructions tips"
-    },
-    "userContextModal.title":
-    {
-        "defaultMessage": "Custom instructions",
-        "description": "title for Custom instructions modal"
-    },
-    "workspaceAnalytics.activeUsersChartLabel":
-    {
-        "defaultMessage": "Active users",
-        "description": "Label for active users chart"
-    },
-    "workspaceAnalytics.activeUsersLastWeek":
-    {
-        "defaultMessage": "Active users last week",
-        "description": "Text under number active users last week"
-    },
-    "workspaceAnalytics.conversationsLabel":
-    {
-        "defaultMessage": "Conversations",
-        "description": "Label of total conversations count"
-    },
-    "workspaceAnalytics.messagesChartLabel":
-    {
-        "defaultMessage": "Messages",
-        "description": "Label for messages chart"
-    },
-    "workspaceAnalytics.messagesLabel":
-    {
-        "defaultMessage": "Messages",
-        "description": "Label of total messages count"
-    },
-    "workspaceAnalytics.messagesLast30Days":
-    {
-        "defaultMessage": "{count, plural, =0 {No messages} one {# message} other {# messages}} processed over last 30 days",
-        "description": "Analytics text about messages"
-    },
-    "workspaceAnalytics.messagesLastWeek":
-    {
-        "defaultMessage": "Messages processed last week",
-        "description": "Text under number messages last week"
-    },
-    "workspaceAnalytics.title":
-    {
-        "defaultMessage": "Workspace Analytics",
-        "description": "Title of analytics page"
-    },
-    "workspaceAnalytics.totals":
-    {
-        "defaultMessage": "Totals over last 30 days",
-        "description": "Section header of total analytics"
-    },
-    "workspaceAnalytics.understand":
-    {
-        "defaultMessage": "Understand how your workspace is using ChatGPT",
-        "description": "Subtext of workspace analytics page"
-    },
-    "workspaceAnalytics.usageTitle":
-    {
-        "defaultMessage": "Usage",
-        "description": "Title of usage section"
-    },
-    "workspaceAnalytics.usersLabel":
-    {
-        "defaultMessage": "Users",
-        "description": "Label of total user count"
-    },
-    "workspaceAnalytics.usersLast30Days":
-    {
-        "defaultMessage": "{count, plural, =0 {No active users} one {# active user} other {# active user}} over last 30 days",
-        "description": "Analytics text about active users"
-    },
-    "workspaceAnalytics.weeklyUsersTitle":
-    {
-        "defaultMessage": "Weekly users",
-        "description": "Title of weekly users section"
-    },
-    "workspaceAppearanceModal.cancel":
-    {
-        "defaultMessage": "Cancel",
-        "description": "Label for the cancel button"
-    },
-    "workspaceAppearanceModal.propagationWarning":
-    {
-        "defaultMessage": "Changes to the workspace name and image may take some time to take effect.",
-        "description": "Warning message about changes to the workspace profile taking a while to show up."
-    },
-    "workspaceAppearanceModal.saveError":
-    {
-        "defaultMessage": "Failed to save workspace appearance",
-        "description": "Error message when saving workspace appearance fails"
-    },
-    "workspaceAppearanceModal.submit":
-    {
-        "defaultMessage": "Save",
-        "description": "Label for the submit button"
-    },
-    "workspaceAppearanceModal.title":
-    {
-        "defaultMessage": "Workspace appearance",
-        "description": "Title for the workspace appearance modal"
-    },
-    "workspaceAppearanceModal.workspaceAvatar":
-    {
-        "defaultMessage": "Workspace image",
-        "description": "Label for the workspace image upload field"
-    },
-    "workspaceAppearanceModal.workspaceAvatarDescription":
-    {
-        "defaultMessage": "Upload a JPEG or PNG workspace image for your team. (Minimum {size}\\xd7{size}px recommended.)",
-        "description": "Help text for the workspace image upload field"
-    },
-    "workspaceAppearanceModal.workspaceName":
-    {
-        "defaultMessage": "Workspace name",
-        "description": "Label for the workspace name input field"
-    },
-    "workspaceAppearanceModal.workspaceNameDescription":
-    {
-        "defaultMessage": "Update the name of the workspace.",
-        "description": "Help text for the workspace name input field"
-    },
-    "workspaceIdentity.acsURLLabel":
-    {
-        "defaultMessage": "Assertion Consumer Service (ACS) URL",
-        "description": "Label for the ACS URL configuration text"
-    },
-    "workspaceIdentity.addDomainButton.0":
-    {
-        "defaultMessage": "Add domain",
-        "description": "Label for add domain button"
-    },
-    "workspaceIdentity.autoProvisionBody":
-    {
-        "defaultMessage": "Automatically create ChatGPT Enterprise accounts for new users who log in with a verified domain.",
-        "description": "Label for toggling automatic provisioning"
-    },
-    "workspaceIdentity.autoProvisionDisabledToast":
-    {
-        "defaultMessage": "Auto provisioning disabled for this workspace",
-        "description": "Toast message for disabling auto provisioning"
-    },
-    "workspaceIdentity.autoProvisionEnabledToast":
-    {
-        "defaultMessage": "Auto provisioning enabled for this workspace",
-        "description": "Toast message for enabling auto provisioning"
-    },
-    "workspaceIdentity.autoProvisionSubtitle":
-    {
-        "defaultMessage": "Automatic account creation",
-        "description": "Subtitle for auto provisioning"
-    },
-    "workspaceIdentity.certExtractFailed":
-    {
-        "defaultMessage": "Unable to find X.509 Certificate in provided XML",
-        "description": "Notice to the user that the given XML did not contain a certificate"
-    },
-    "workspaceIdentity.copiedACSToClipboard":
-    {
-        "defaultMessage": "Copied ACS URL to clipboard",
-        "description": "Message for success toast on copying post-back url"
-    },
-    "workspaceIdentity.copiedEntityToClipboard":
-    {
-        "defaultMessage": "Copied entity ID to clipboard",
-        "description": "Message for success toast on copying entity ID"
-    },
-    "workspaceIdentity.copiedIDPUrlToClipboard":
-    {
-        "defaultMessage": "Copied IDP URL to clipboard",
-        "description": "Message for success toast on copying IDP URL"
-    },
-    "workspaceIdentity.domainSubtitle":
-    {
-        "defaultMessage": "Domain management",
-        "description": "Section header of domain management"
-    },
-    "workspaceIdentity.domainTableHeader":
-    {
-        "defaultMessage": "Domain",
-        "description": "Label for domain column on domains table"
-    },
-    "workspaceIdentity.enforceSSOBody":
-    {
-        "defaultMessage": "When SSO is active, users will no longer be able to use social logins and will be redirected to your SSO portal.",
-        "description": "Message telling users what enforcing SSO will do"
-    },
-    "workspaceIdentity.enforceSSOLabel":
-    {
-        "defaultMessage": "Toggle to enable or disable SSO enforcement",
-        "description": "Label for toggle for SSO enforcement"
-    },
-    "workspaceIdentity.enforceSSOTitle":
-    {
-        "defaultMessage": "Enforce SSO log in",
-        "description": "Section title for toggling SSO enforcement"
-    },
-    "workspaceIdentity.entityIDLabel":
-    {
-        "defaultMessage": "Entity ID",
-        "description": "Label for the Entity ID configuration text"
-    },
-    "workspaceIdentity.idpSignInURL":
-    {
-        "defaultMessage": "IDP Tile URL",
-        "description": "Label for the IDP Tile URL configuration text"
-    },
-    "workspaceIdentity.learnMoreLink":
-    {
-        "defaultMessage": "Learn more",
-        "description": "Label for link to learn more"
-    },
-    "workspaceIdentity.orDivider":
-    {
-        "defaultMessage": "or",
-        "description": "Separator between the top and bottom portions of the form"
-    },
-    "workspaceIdentity.setupSSOContent":
-    {
-        "defaultMessage": "Anyone using email addresses with a verified domain can log in via SAML SSO.",
-        "description": "Content describing how SAML SSO will work when set up"
-    },
-    "workspaceIdentity.ssoAddButton":
-    {
-        "defaultMessage": "Add SAML SSO",
-        "description": "Label for add SSO button"
-    },
-    "workspaceIdentity.ssoDropdownRemoveButtonText":
-    {
-        "defaultMessage": "Remove Domain",
-        "description": "Dropdown choice for remove domain"
-    },
-    "workspaceIdentity.ssoDropdownVerifyButtonText":
-    {
-        "defaultMessage": "Verify",
-        "description": "Dropdown choice to bring up the verification modal"
-    },
-    "workspaceIdentity.ssoEditButton":
-    {
-        "defaultMessage": "Edit SAML SSO",
-        "description": "Label for edit SSO button"
-    },
-    "workspaceIdentity.ssoLearnMoreLink":
-    {
-        "defaultMessage": "Learn more",
-        "description": "Label for link to learn more"
-    },
-    "workspaceIdentity.ssoTitle":
-    {
-        "defaultMessage": "Single Sign On",
-        "description": "Section title of SSO"
-    },
-    "workspaceIdentity.ssoUrlExtractFailed":
-    {
-        "defaultMessage": "Unable to find Entity ID (sign-in endpoint) in provided XML",
-        "description": "Notice to the user that the given XML did not contain an entity ID"
-    },
-    "workspaceIdentity.statusTableHeader":
-    {
-        "defaultMessage": "Status",
-        "description": "Label for status column on domains table"
-    },
-    "workspaceIdentity.title":
-    {
-        "defaultMessage": "Identity & Provisioning",
-        "description": "Title of identity page"
-    },
-    "workspaceIdentity.toggleAutoProvisionLabel":
-    {
-        "defaultMessage": "Toggle to enable or disable auto provisioning",
-        "description": "Label for auto provisioning toggle"
-    },
-    "workspaceIdentity.verifyDomainInfoText":
-    {
-        "defaultMessage": "Verify ownership of an email domain to access advanced security features including single sign-on.",
-        "description": "Informational text for domain verification section"
-    },
-    "workspacePopoverNavigation.chatPreferences":
-    {
-        "defaultMessage": "Custom instructions",
-        "description": "Custom instructions menu item"
-    },
-    "workspacePopoverNavigation.myGpts":
-    {
-        "defaultMessage": "My GPTs",
-        "description": "My GPTs menu item"
-    },
     "workspacePopoverNavigation.myWorkspaceSettings":
     {
         "defaultMessage": "Manage workspace",
@@ -5970,6 +6810,16 @@
         "defaultMessage": "Appearance",
         "description": "Title for appearance settings"
     },
+    "workspaceSettings.approveAGPTButton":
+    {
+        "defaultMessage": "Approve a GPT",
+        "description": "Text for the approve a GPT button"
+    },
+    "workspaceSettings.approvedCellTitle":
+    {
+        "defaultMessage": "Approved",
+        "description": "Approved cell title"
+    },
     "workspaceSettings.browseOptionTitle":
     {
         "defaultMessage": "Browsing with Bing",
@@ -5980,21 +6830,61 @@
         "defaultMessage": "Allow ChatGPT and GPTs to browse the internet to answer questions about recent topics and events.",
         "description": "Tooltip message for browse setting toggle"
     },
+    "workspaceSettings.builderCellTitle":
+    {
+        "defaultMessage": "Builder",
+        "description": "Builder cell title"
+    },
+    "workspaceSettings.capabilitiesTitle":
+    {
+        "defaultMessage": "Capabilities",
+        "description": "Title for the capabilities section in the tooltip"
+    },
+    "workspaceSettings.changeAccess":
+    {
+        "defaultMessage": "Change who has access",
+        "description": "Change who has access button text"
+    },
+    "workspaceSettings.changeOwner":
+    {
+        "defaultMessage": "Change Owner",
+        "description": "Change owner button text"
+    },
     "workspaceSettings.chatSharingTitle":
     {
         "defaultMessage": "Chats can be shared with...",
         "description": "Title for chat sharing select"
     },
+    "workspaceSettings.chatsCellTitle":
+    {
+        "defaultMessage": "Chats",
+        "description": "Chats cell title"
+    },
     "workspaceSettings.conversationRetention":
     {
         "defaultMessage": "Chat retention",
         "description": "Conversation retention setting label"
     },
+    "workspaceSettings.createdCellTitle":
+    {
+        "defaultMessage": "Created",
+        "description": "Created cell title"
+    },
+    "workspaceSettings.customActionsCellTitle":
+    {
+        "defaultMessage": "Custom Actions",
+        "description": "Custom Actions cell title"
+    },
     "workspaceSettings.customActionsTooltip":
     {
         "defaultMessage": "Allow GPTs to use third-party services for tasks such as finding flights. Actions are defined by GPT builders.",
         "description": "Tooltip message for custom actions toggle"
     },
+    "workspaceSettings.deleteGPT":
+    {
+        "defaultMessage": "Delete GPT",
+        "description": "Delete GPT button text"
+    },
     "workspaceSettings.dontAllowOption":
     {
         "defaultMessage": "Don't allow",
@@ -6015,16 +6905,46 @@
         "defaultMessage": "Failed to update feature setting",
         "description": "Message shown when there's an error updating feature settings"
     },
+    "workspaceSettings.filterByNameOrBuilderPlaceholder":
+    {
+        "defaultMessage": "Filter by name or builder...",
+        "description": "Placeholder text for the search input"
+    },
     "workspaceSettings.gptSharingTitle":
     {
         "defaultMessage": "GPTs can be shared with...",
         "description": "Title for gpt sharing select"
     },
+    "workspaceSettings.nameCellTitle":
+    {
+        "defaultMessage": "Name",
+        "description": "GPT Name cell title"
+    },
     "workspaceSettings.noOneSelect":
     {
         "defaultMessage": "No one",
         "description": "Select value for no one option for sharing"
     },
+    "workspaceSettings.noWorkspaceApprovedGPTs":
+    {
+        "defaultMessage": "You havent approved any third-party GPTs",
+        "description": "Message for empty GPT table state"
+    },
+    "workspaceSettings.noWorkspaceApprovedGPTsFound":
+    {
+        "defaultMessage": "No approved third-party GPTs found",
+        "description": "Message for empty filtered third-party GPT table state"
+    },
+    "workspaceSettings.noWorkspaceGPTs":
+    {
+        "defaultMessage": "This workspace does not have any GPTs",
+        "description": "Message for empty GPT table state"
+    },
+    "workspaceSettings.noWorkspaceGPTsFound":
+    {
+        "defaultMessage": "No workspace GPTs found",
+        "description": "Message for empty filtered GPT table state"
+    },
     "workspaceSettings.pluginsOptionTitle":
     {
         "defaultMessage": "Plugins",
@@ -6040,6 +6960,11 @@
         "defaultMessage": "Allow members to install plugins so ChatGPT can use third-party services for tasks such as finding flights.",
         "description": "Tooltip message for plugin setting toggle"
     },
+    "workspaceSettings.removeFromWorkspace":
+    {
+        "defaultMessage": "Remove from workspace",
+        "description": "Remove from workspace button text"
+    },
     "workspaceSettings.retentionChange":
     {
         "defaultMessage": "Contact your account manager to change this setting.",
@@ -6075,6 +7000,21 @@
         "defaultMessage": "Contact your account manager to change this setting",
         "description": "Tooltip for conversation retention setting"
     },
+    "workspaceSettings.shareRecipientTitleLink":
+    {
+        "defaultMessage": "Link",
+        "description": "Tooltip for enterprise plan required sections"
+    },
+    "workspaceSettings.shareRecipientTitlePrivate":
+    {
+        "defaultMessage": "Private",
+        "description": "Tooltip for enterprise plan required sections"
+    },
+    "workspaceSettings.shareRecipientTitlePublic":
+    {
+        "defaultMessage": "Public",
+        "description": "Tooltip for enterprise plan required sections"
+    },
     "workspaceSettings.sharingSettingsTitle":
     {
         "defaultMessage": "Sharing",
@@ -6087,14 +7027,24 @@
     },
     "workspaceSettings.thirdPartyGPTsTitle":
     {
-        "defaultMessage": "Third-party GPTs",
+        "defaultMessage": "Third-party",
         "description": "Third Party GPTs setting title"
+    },
+    "workspaceSettings.thirdPartyTab":
+    {
+        "defaultMessage": "Third-party",
+        "description": "Tab title for the third-party section"
     },
     "workspaceSettings.title":
     {
         "defaultMessage": "Settings",
         "description": "Title of settings page"
     },
+    "workspaceSettings.whoHasAccessTitle":
+    {
+        "defaultMessage": "Who has access",
+        "description": "Title for the who has access section in the tooltip"
+    },
     "workspaceSettings.workspaceGPTsDescription":
     {
         "defaultMessage": "Manage which capabilities are available for standard ChatGPT and custom GPTs created in your workspace. These settings do not apply to third-party GPTs.",
@@ -6106,6 +7056,186 @@
         "description": "Title for workspace gpts settings"
     },
     "workspaceSettings.workspaceMembersOnlySelect":
+    {
+        "defaultMessage": "Workspace members only",
+        "description": "Select value for workspace members only option for sharing"
+    },
+    "workspaceSettings.workspaceTab":
+    {
+        "defaultMessage": "Workspace",
+        "description": "Tab title for the workspace section"
+    },
+    "workspaceSettingsOld.allowAllOption":
+    {
+        "defaultMessage": "Allow all",
+        "description": "Allow all option for third party gpts"
+    },
+    "workspaceSettingsOld.anyOneSelect":
+    {
+        "defaultMessage": "Anyone",
+        "description": "Select value for any one option for sharing"
+    },
+    "workspaceSettingsOld.appearanceTitle":
+    {
+        "defaultMessage": "Appearance",
+        "description": "Title for appearance settings"
+    },
+    "workspaceSettingsOld.browseOptionTitle":
+    {
+        "defaultMessage": "Browsing with Bing",
+        "description": "Browse setting title"
+    },
+    "workspaceSettingsOld.browseTooltip":
+    {
+        "defaultMessage": "Allow ChatGPT and GPTs to browse the internet to answer questions about recent topics and events.",
+        "description": "Tooltip message for browse setting toggle"
+    },
+    "workspaceSettingsOld.chatSharingTitle":
+    {
+        "defaultMessage": "Chats can be shared with...",
+        "description": "Title for chat sharing select"
+    },
+    "workspaceSettingsOld.conversationRetention":
+    {
+        "defaultMessage": "Chat retention",
+        "description": "Conversation retention setting label"
+    },
+    "workspaceSettingsOld.customActionsTooltip":
+    {
+        "defaultMessage": "Allow GPTs to use third-party services for tasks such as finding flights. Actions are defined by GPT builders.",
+        "description": "Tooltip message for custom actions toggle"
+    },
+    "workspaceSettingsOld.dontAllowOption":
+    {
+        "defaultMessage": "Don't allow",
+        "description": "Don't allow option for third party gpts"
+    },
+    "workspaceSettingsOld.enterpriseRequiredMessage":
+    {
+        "defaultMessage": "Only workspaces with the Enterprise plan can change these settings",
+        "description": "Tooltip for enterprise plan required sections"
+    },
+    "workspaceSettingsOld.enterpriseUpsellPill":
+    {
+        "defaultMessage": "Enterprise",
+        "description": "Pill contents for enterprise upsell"
+    },
+    "workspaceSettingsOld.featureSettingsUpdateFailed":
+    {
+        "defaultMessage": "Failed to update feature setting",
+        "description": "Message shown when there's an error updating feature settings"
+    },
+    "workspaceSettingsOld.gptSharingTitle":
+    {
+        "defaultMessage": "GPTs can be shared with...",
+        "description": "Title for gpt sharing select"
+    },
+    "workspaceSettingsOld.noOneSelect":
+    {
+        "defaultMessage": "No one",
+        "description": "Select value for no one option for sharing"
+    },
+    "workspaceSettingsOld.ownerApprovedOnlyOption":
+    {
+        "defaultMessage": "Owner-approved only",
+        "description": "Owner-approved only option for third party gpts"
+    },
+    "workspaceSettingsOld.pluginsOptionTitle":
+    {
+        "defaultMessage": "Plugins",
+        "description": "Plugins setting title"
+    },
+    "workspaceSettingsOld.pluginsOptionTitle.0":
+    {
+        "defaultMessage": "Custom actions",
+        "description": "Custom Actions setting title"
+    },
+    "workspaceSettingsOld.pluginsTooltip":
+    {
+        "defaultMessage": "Allow members to install plugins so ChatGPT can use third-party services for tasks such as finding flights.",
+        "description": "Tooltip message for plugin setting toggle"
+    },
+    "workspaceSettingsOld.retentionChange":
+    {
+        "defaultMessage": "Contact your account manager to change this setting.",
+        "description": "Text underneath conversation retention setting"
+    },
+    "workspaceSettingsOld.retentionCustomDays":
+    {
+        "defaultMessage": "{num} days",
+        "description": "Conversation retention value"
+    },
+    "workspaceSettingsOld.retentionInfinite":
+    {
+        "defaultMessage": "Infinite",
+        "description": "Conversation retention value"
+    },
+    "workspaceSettingsOld.retentionNinetyDays":
+    {
+        "defaultMessage": "90 days",
+        "description": "Conversation retention value"
+    },
+    "workspaceSettingsOld.retentionOneYear":
+    {
+        "defaultMessage": "1 year",
+        "description": "Conversation retention value"
+    },
+    "workspaceSettingsOld.retentionPolicyTitle":
+    {
+        "defaultMessage": "Retention policy",
+        "description": "Title of Retention policy settings section"
+    },
+    "workspaceSettingsOld.retentionSettingTooltip":
+    {
+        "defaultMessage": "Contact your account manager to change this setting",
+        "description": "Tooltip for conversation retention setting"
+    },
+    "workspaceSettingsOld.shareRecipientTitleLink":
+    {
+        "defaultMessage": "Link",
+        "description": "Tooltip for enterprise plan required sections"
+    },
+    "workspaceSettingsOld.shareRecipientTitlePrivate":
+    {
+        "defaultMessage": "Private",
+        "description": "Tooltip for enterprise plan required sections"
+    },
+    "workspaceSettingsOld.shareRecipientTitlePublic":
+    {
+        "defaultMessage": "Public",
+        "description": "Tooltip for enterprise plan required sections"
+    },
+    "workspaceSettingsOld.sharingSettingsTitle":
+    {
+        "defaultMessage": "Sharing",
+        "description": "Title for workspace sharing toggles"
+    },
+    "workspaceSettingsOld.sidebarTitle":
+    {
+        "defaultMessage": "GPTs",
+        "description": "Title of settings page in the sidebar"
+    },
+    "workspaceSettingsOld.thirdPartyGPTsDescription":
+    {
+        "defaultMessage": "Manage whether members can use GPTs created outside your workspace.",
+        "description": "Description of third GPT settings"
+    },
+    "workspaceSettingsOld.thirdPartyGPTsTitle":
+    {
+        "defaultMessage": "Third-party",
+        "description": "Third Party GPTs setting title"
+    },
+    "workspaceSettingsOld.workspaceGPTsDescription":
+    {
+        "defaultMessage": "Manage which capabilities are available for GPTs built in your workspace or by ChatGPT. This does not apply to third-party GPTs.",
+        "description": "Workspace GPTS settings description"
+    },
+    "workspaceSettingsOld.workspaceGPTsTitle":
+    {
+        "defaultMessage": "Workspace",
+        "description": "Title for workspace gpts settings"
+    },
+    "workspaceSettingsOld.workspaceMembersOnlySelect":
     {
         "defaultMessage": "Workspace members only",
         "description": "Select value for workspace members only option for sharing"


</files>
