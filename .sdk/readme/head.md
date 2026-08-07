# Luma SDK

Luma API clients in TypeScript, JavaScript, Go, Python, PHP, and Lua, plus a CLI and an MCP
server for AI agents. All generated from Luma's public OpenAPI spec, so every surface stays in
sync with the API.

> **Unofficial.** This is an unofficial SDK for the Luma public API, built by
> [Voxgig](https://voxgig.com/sdk). It is not affiliated with, endorsed by, or sponsored by
> Luma.

**Why this exists:** Voxgig builds public SDK and MCP examples for APIs we think are
interesting. This is one of those. MIT-licensed, take whatever's useful.

Luma publishes an OpenAPI spec at `https://public-api.luma.com/openapi.json`, and its docs say
you can use it "to generate client libraries, import into tools like Postman, or build
integrations with AI agents". That is exactly what this repo is.

All 66 operations in the spec are covered, in every language.

## Try it (TypeScript)

```bash
git clone https://github.com/voxgig-sdk/luma-sdk
cd luma-sdk/ts
npm install && npm run build && npm test
```

The test suite runs fully offline. Every SDK ships a test mode that swaps the HTTP transport
for an in-memory mock, so you can try it without credentials or a network.

## Quickstart

Every Luma endpoint authenticates with an `x-luma-api-key` header. Generate a key at
`luma.com/calendar/manage/api-keys`. API access needs a Luma Plus subscription.

```ts
import { LumaSDK } from '@voxgig-sdk/luma'

const client = new LumaSDK({
  apikey: process.env.LUMA_APIKEY,
})

// Events on your calendar  ->  GET /v1/calendars/events/list
const events = await client.CalendarEvent().list({})
for (const event of events) {
  console.log(event.id, event.status)
}

// Guests on one event  ->  GET /v1/events/guests/list
const guests = await client.Guest().list({ event_id: 'evt-XXXXXXXX' })
for (const guest of guests) {
  console.log(guest.user_email, guest.approval_status)
}
```

## What's in the box

| Surface | Use it for | Where |
|---|---|---|
| SDK, 6 languages | App integration | `/ts` `/js` `/go` `/py` `/php` `/lua` |
| CLI | Scripts, CI, exploration (interactive REPL mode included) | `/go-cli` |
| MCP server | AI agents: Claude, Cursor, and friends | `/go-mcp` |
| Agent guide | Points coding agents at all of the above | `AGENTS.md` |

All of it comes from one spec. Change the spec, regenerate, and every surface updates
together. None of them drift.

## Using the MCP server

```bash
cd go-mcp && go build -o luma-mcp .
```

```json
{
  "mcpServers": {
    "luma": {
      "command": "/path/to/luma-sdk/go-mcp/luma-mcp",
      "args": ["-transport", "stdio"],
      "env": { "LUMA_APIKEY": "your-luma-api-key" }
    }
  }
}
```

The server registers as `luma` and exposes two entity-parameterized tools, `luma_list` and
`luma_load`, each taking an `entity` name plus an optional `query`. Both are read-only. Writes
go through the SDKs or the CLI.

Your customers' AI agents can call the Luma API through this today.

## Honest state

Generated from Luma's public OpenAPI spec on 2026-08-07. Not production-tuned. Use it as a
starting point or a reference.

Known rough edges, all traceable to the spec rather than to any one language:

- **About a quarter of the generated fields come through as `any`.** Luma's spec declares
  nullable fields as `anyOf: [{...}, {"type": "null"}]`, 599 times, and defines every schema
  inline rather than in `components.schemas`, so there is nothing named to generate a type
  from. Of 365 model fields, 202 land as `string` and 96 land untyped. `Calendar.name` is a
  `string`; `Calendar.avatar_url` is `any`.
- **`load()` takes the generator's `id` field, but Luma's endpoints name their own
  parameter.** `Event().load()` wants `{ id, event_id }`, not just `{ id }`, because
  `/v1/events/get` declares its query parameter as `event_id`. The entity table below lists
  each endpoint's real path; pass the parameter the spec names alongside `id`.
- **The entity map in `.sdk/model/guide/guide.aontu` was written by hand.** Luma's API puts
  the verb in the path (`/v1/events/create`) and spreads all 66 operations across only six tag
  values, so grouping by tag collapses unrelated endpoints together. The guide maps each path
  to an entity explicitly. Regenerating picks up spec changes to existing paths; brand new
  paths need a line added to the guide.

Every one of the 66 operations was checked end to end, in all six languages, against a local
recording server: correct HTTP method, correct path, every time. Test suites are green
(TypeScript 278, JavaScript 270, Python 279, PHP 330, Go ok).

When teams want SDKs like these production-grade, idiomatic per language, tested, documented,
and released through a real pipeline, Voxgig does that work as a consulting engagement. The
toolkit also generates Java and C# if your customers need them.

