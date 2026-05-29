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