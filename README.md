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

## Quickstart in other languages

### Python

```python
import os
from luma_sdk import LumaSDK

client = LumaSDK({
    "apikey": os.environ.get("LUMA_APIKEY"),
})


# Load a specific calendar (returns the record, raises on error)
calendar = client.Calendar().load({"id": "example_id"})
print(calendar)
```

### PHP

```php
<?php
require_once 'luma_sdk.php';

$client = new LumaSDK([
    "apikey" => getenv("LUMA_APIKEY"),
]);


// Load a specific calendar (returns the bare record; throws on error)
$calendar = $client->Calendar()->load(["id" => "example_id"]);
print_r($calendar);
```

### Golang

```go
import sdk "github.com/voxgig-sdk/luma-sdk/go"

client := sdk.NewLumaSDK(map[string]any{
    "apikey": os.Getenv("LUMA_APIKEY"),
})

// Load calendar data
calendar, err := client.Calendar(nil).Load(map[string]any{"id": "example_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(calendar)
```

### Lua

```lua
local sdk = require("luma_sdk")

local client = sdk.new({
  apikey = os.getenv("LUMA_APIKEY"),
})


-- Load a specific calendar
local calendar, err = client:Calendar():load({ id = "example_id" })
print(calendar)
```

### JavaScript

```js
const { LumaSDK } = require('@voxgig-sdk/luma-js')

const client = new LumaSDK({
  apikey: process.env.LUMA_APIKEY,
})

```

## Direct and prepare

For endpoints the entity model doesn't cover, use the low-level methods:

- **`direct(fetchargs)`** — build and send an HTTP request in one step.
- **`prepare(fetchargs)`** — build the request without sending it.

Both accept a map with `path`, `method`, `params`, `query`,
`headers`, and `body`. See the [How-to guides](#how-to-guides) below.

## How-to guides

### Make a direct API call

When the entity interface does not cover an endpoint, use `direct`:

**TypeScript:**
```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})
if (result instanceof Error) {
  throw result
}
console.log(result.data)
```

**Python:**
```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})
```

**PHP:**
```php
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);
```

**Go:**
```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}
fmt.Println(result)
```

**Lua:**
```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
```

**JavaScript:**
```js
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})
if (result instanceof Error) {
  throw result
}
console.log(result.data)
```

## Entities

The API exposes 31 entities:

| Entity | Description | API path |
| --- | --- | --- |
| **Calendar** | The Calendar entity (load, update). | `/v1/calendars/get` |
| **CalendarAdmin** | The CalendarAdmin entity (list). | `/v1/calendars/admins/list` |
| **CalendarCoupon** | The CalendarCoupon entity (create, list, update). | `/v1/calendars/coupons/create` |
| **CalendarEvent** | The CalendarEvent entity (create, list, load). | `/v1/calendars/events/add` |
| **CalendarEventApproval** | The CalendarEventApproval entity (create). | `/v1/calendars/events/approve` |
| **CalendarEventRejection** | The CalendarEventRejection entity (create). | `/v1/calendars/events/reject` |
| **Contact** | The Contact entity (create, list, remove). | `/v1/calendars/contacts/import` |
| **ContactBlock** | The ContactBlock entity (create). | `/v1/calendars/contacts/block` |
| **ContactRestore** | The ContactRestore entity (create). | `/v1/calendars/contacts/restore` |
| **ContactTag** | The ContactTag entity (create, list, remove, update). | `/v1/calendars/contact-tags/create` |
| **ContactTagAssignment** | The ContactTagAssignment entity (create, remove). | `/v1/calendars/contact-tags/apply` |
| **EntityLookup** | The EntityLookup entity (load). | `/v1/entities/lookup` |
| **Event** | The Event entity (create, load, remove, update). | `/v1/events/create` |
| **EventCancelRequest** | The EventCancelRequest entity (create). | `/v1/events/cancel/request` |
| **EventCoupon** | The EventCoupon entity (create, list, update). | `/v1/events/coupons/create` |
| **EventTag** | The EventTag entity (create, list, remove, update). | `/v1/calendars/event-tags/create` |
| **EventTagAssignment** | The EventTagAssignment entity (create, remove). | `/v1/calendars/event-tags/apply` |
| **Guest** | The Guest entity (create, list, load, update). | `/v1/events/guests/add` |
| **GuestInvite** | The GuestInvite entity (create). | `/v1/events/guests/send-invites` |
| **GuestTicket** | The GuestTicket entity (update). | `/v1/events/guests/update-tickets` |
| **Host** | The Host entity (create, remove, update). | `/v1/events/hosts/add` |
| **ImageUpload** | The ImageUpload entity (create). | `/v1/images/create-upload-url` |
| **Member** | The Member entity (create, update). | `/v1/memberships/members/add` |
| **MembershipTier** | The MembershipTier entity (list). | `/v1/memberships/tiers/list` |
| **OrganizationAdmin** | The OrganizationAdmin entity (list). | `/v1/organizations/admins/list` |
| **OrganizationCalendar** | The OrganizationCalendar entity (create, list). | `/v2/organizations/calendars/create` |
| **OrganizationEvent** | The OrganizationEvent entity (list). | `/v1/organizations/events/list` |
| **OrganizationEventTransfer** | The OrganizationEventTransfer entity (create). | `/v1/organizations/events/transfer-calendar` |
| **TicketType** | The TicketType entity (create, list, load, remove, update). | `/v1/events/ticket-types/create` |
| **User** | The User entity (load). | `/v1/users/get-self` |
| **Webhook** | The Webhook entity (create, list, load, remove, update). | `/v2/webhooks/create` |

The operations available across these entities are **load**, **list**, **create**, **update**, **remove** — see each entity's
own list above for exactly which it supports.

## How it works

> Everyday use only needs the sections above. This explains the internals
> behind every call — relevant when writing custom features.

Every SDK call runs the same five-stage pipeline:

1. **Point** — resolve the API endpoint from the operation definition.
2. **Spec** — build the HTTP specification (URL, method, headers, body).
3. **Request** — send the HTTP request.
4. **Response** — receive and parse the response.
5. **Result** — extract the result data for the caller.

A feature hook fires at each stage (e.g. `PrePoint`, `PreSpec`,
`PreRequest`), so features can inspect or modify the pipeline without
forking the SDK.

### Features

| Feature | Purpose |
| --- | --- |
| **TestFeature** | In-memory mock transport for testing without a live server |

Pass custom features via the `extend` option at construction time.

## Per-language documentation

- [TypeScript](ts/README.md)
- [JavaScript](js/README.md)
- [Python](py/README.md)
- [PHP](php/README.md)
- [Go](go/README.md)
- [Lua](lua/README.md)

## Upstream API

Generated from Luma's public OpenAPI spec. The spec is vendored, unmodified, in
[`.sdk/def/`](.sdk/def/) with its source URL, fetch date, and checksum.

- Luma API docs: [https://docs.luma.com](https://docs.luma.com)
- Spec: [https://public-api.luma.com/openapi.json](https://public-api.luma.com/openapi.json)

## Security

Please report security issues to security@voxgig.com. See [SECURITY.md](SECURITY.md).
Do not open public issues for suspected vulnerabilities.

---

Generated by the [Voxgig SDK Generator](https://voxgig.com/sdk), MIT-licensed. Browse 600+
generated SDKs at [github.com/voxgig-sdk](https://github.com/voxgig-sdk).

Questions, or want these production-grade? Email richard@voxgig.com.

If you are from Luma and would like this repository removed, or transferred to your own GitHub
organisation, email richard@voxgig.com and it will be done within two business days, no
questions asked.
