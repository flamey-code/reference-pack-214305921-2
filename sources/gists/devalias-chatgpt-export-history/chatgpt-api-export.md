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