# Luma JavaScript SDK



The JavaScript SDK for the Luma API — an entity-oriented client with full async/await support.

The API is exposed as capitalised, semantic **Entities** — e.g.
`client.Calendar()` — each with a small set of operations (`list`, `load`, `create`, `update`, `remove`)
instead of raw URL paths and query parameters. This keeps the surface
predictable and low-friction for both humans and AI agents.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
```js
npm install luma
```
## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.


### Create a Client

```js
const { LumaSDK } = require('@voxgig-sdk/luma-js')

const client = new LumaSDK({
  apikey: process.env.LUMA_APIKEY,
})
```

### Load a Calendar

```js
const calendar = await client.Calendar().load({ id: 'calendar_id' })
console.log(calendar)
```

### Update a Calendar

```js
const updated = await client.Calendar().update({
  id: 'calendar_id',
  avatar_url: 'example_avatar_url',
  calendar_id: 'example_calendar_id',
})
console.log(updated)
```

### Direct API Access

Use `client.direct()` to call any API endpoint directly:

```js
const result = await client.direct({
  path: '/custom/endpoint/{id}',
  method: 'GET',
  params: { id: 'abc123' },
})

if (result.ok) {
  console.log(result.data)
}
```


## Error handling

Entity operations reject on failure, so wrap them in `try` / `catch`:

```ts
try {
  const eventcoupons = await client.EventCoupon().list()
  console.log(eventcoupons)
} catch (err) {
  console.error('list failed:', err)
}
```

The low-level `direct()` method does **not** throw — it returns the
value or an `Error`, so check the result before using it:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example_id' },
})

if (result instanceof Error) {
  throw result
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```js
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})

if (result instanceof Error) {
  throw result
}
if (result.ok) {
  console.log(result.status)  // 200
  console.log(result.data)    // response body
}
```

### Prepare a request without sending it

```js
const fetchdef = await client.prepare({
  path: '/api/resource/{id}',
  method: 'DELETE',
  params: { id: 'example' },
})

// Inspect before sending
console.log(fetchdef.url)
console.log(fetchdef.method)
console.log(fetchdef.headers)
```

### Use test mode

Create a mock client for unit testing — no server required:

```js
const client = LumaSDK.test()

const eventcoupon = await client.EventCoupon().list()
// eventcoupon is a bare entity populated with mock response data
console.log(eventcoupon)
```

You can also use the instance method:

```js
const client = new LumaSDK({ apikey: '...' })
const testClient = client.tester()
```

### Retain entity state across calls

Entity instances remember their last match and data:

```js
const entity = client.EventCoupon()

// First call runs the operation and stores its result
await entity.list()

// Subsequent calls reuse the stored state
const data = entity.data()
console.log(data.id)
```

### Add custom middleware

Pass features via the `extend` option:

```js
const logger = {
  hooks: {
    PreRequest: (ctx) => {
      console.log('Requesting:', ctx.spec.method, ctx.spec.path)
    },
    PreResponse: (ctx) => {
      console.log('Status:', ctx.out.request?.status)
    },
  },
}

const client = new LumaSDK({
  apikey: '...',
  extend: [logger],
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
LUMA_TEST_LIVE=TRUE
LUMA_APIKEY=<your-key>
```

Then run:

```bash
cd js && npm test
```


## Reference

### LumaSDK

#### Constructor

```js
new LumaSDK(options?)
```

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `object` | Feature activation flags (e.g. `{ test: { active: true } }`). |
| `extend` | `Feature[]` | Additional feature instances to load. |

#### Methods

| Method | Returns | Description |
| --- | --- | --- |
| `options()` | `object` | Deep copy of current SDK options. |
| `utility()` | `Utility` | Deep copy of the SDK utility object. |
| `prepare(fetchargs?)` | `Promise<FetchDef>` | Build an HTTP request definition without sending it. |
| `direct(fetchargs?)` | `Promise<DirectResult>` | Build and send an HTTP request. |
| `Calendar(data?)` | `CalendarEntity` | Create a Calendar entity instance. |
| `CalendarAdmin(data?)` | `CalendarAdminEntity` | Create a CalendarAdmin entity instance. |
| `CalendarCoupon(data?)` | `CalendarCouponEntity` | Create a CalendarCoupon entity instance. |
| `CalendarEvent(data?)` | `CalendarEventEntity` | Create a CalendarEvent entity instance. |
| `CalendarEventApproval(data?)` | `CalendarEventApprovalEntity` | Create a CalendarEventApproval entity instance. |
| `CalendarEventRejection(data?)` | `CalendarEventRejectionEntity` | Create a CalendarEventRejection entity instance. |
| `Contact(data?)` | `ContactEntity` | Create a Contact entity instance. |
| `ContactBlock(data?)` | `ContactBlockEntity` | Create a ContactBlock entity instance. |
| `ContactRestore(data?)` | `ContactRestoreEntity` | Create a ContactRestore entity instance. |
| `ContactTag(data?)` | `ContactTagEntity` | Create a ContactTag entity instance. |
| `ContactTagAssignment(data?)` | `ContactTagAssignmentEntity` | Create a ContactTagAssignment entity instance. |
| `EntityLookup(data?)` | `EntityLookupEntity` | Create an EntityLookup entity instance. |
| `Event(data?)` | `EventEntity` | Create an Event entity instance. |
| `EventCancelRequest(data?)` | `EventCancelRequestEntity` | Create an EventCancelRequest entity instance. |
| `EventCoupon(data?)` | `EventCouponEntity` | Create an EventCoupon entity instance. |
| `EventTag(data?)` | `EventTagEntity` | Create an EventTag entity instance. |
| `EventTagAssignment(data?)` | `EventTagAssignmentEntity` | Create an EventTagAssignment entity instance. |
| `Guest(data?)` | `GuestEntity` | Create a Guest entity instance. |
| `GuestInvite(data?)` | `GuestInviteEntity` | Create a GuestInvite entity instance. |
| `GuestTicket(data?)` | `GuestTicketEntity` | Create a GuestTicket entity instance. |
| `Host(data?)` | `HostEntity` | Create a Host entity instance. |
| `ImageUpload(data?)` | `ImageUploadEntity` | Create an ImageUpload entity instance. |
| `Member(data?)` | `MemberEntity` | Create a Member entity instance. |
| `MembershipTier(data?)` | `MembershipTierEntity` | Create a MembershipTier entity instance. |
| `OrganizationAdmin(data?)` | `OrganizationAdminEntity` | Create an OrganizationAdmin entity instance. |
| `OrganizationCalendar(data?)` | `OrganizationCalendarEntity` | Create an OrganizationCalendar entity instance. |
| `OrganizationEvent(data?)` | `OrganizationEventEntity` | Create an OrganizationEvent entity instance. |
| `OrganizationEventTransfer(data?)` | `OrganizationEventTransferEntity` | Create an OrganizationEventTransfer entity instance. |
| `TicketType(data?)` | `TicketTypeEntity` | Create a TicketType entity instance. |
| `User(data?)` | `UserEntity` | Create an User entity instance. |
| `Webhook(data?)` | `WebhookEntity` | Create a Webhook entity instance. |
| `tester(testopts?, sdkopts?)` | `LumaSDK` | Create a test-mode client instance. |

#### Static methods

| Method | Returns | Description |
| --- | --- | --- |
| `LumaSDK.test(testopts?, sdkopts?)` | `LumaSDK` | Create a test-mode client. |

### Entity interface

All entities share the same interface.

#### Methods

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `load(reqmatch?, ctrl?): Promise<Entity>` | Load a single entity by match criteria. |
| `list` | `list(reqmatch?, ctrl?): Promise<Entity[]>` | List entities matching the criteria. |
| `create` | `create(reqdata?, ctrl?): Promise<Entity>` | Create a new entity. |
| `update` | `update(reqdata?, ctrl?): Promise<Entity>` | Update an existing entity. |
| `remove` | `remove(reqmatch?, ctrl?): Promise<void>` | Remove an entity. |
| `data` | `data(data?: Partial<Entity>): Entity` | Get or set entity data. |
| `match` | `match(match?: Partial<Entity>): Partial<Entity>` | Get or set entity match criteria. |
| `make` | `make(): Entity` | Create a new instance with the same options. |
| `client` | `client(): LumaSDK` | Return the parent SDK client. |
| `entopts` | `entopts(): object` | Return a copy of the entity options. |

#### Return values

Entity operations resolve to the entity data directly — there is no
result envelope:

- `load`, `create` and `update` resolve to a single entity object.
- `list` resolves to an **array** of entity objects (iterate it directly;
  there is no `.data` and no `.ok`).
- `remove` resolves to `undefined`.

On a failed request these methods **throw**, so wrap calls in
`try`/`catch` to handle errors. Only `direct()` returns the result
envelope described below.

### DirectResult shape

The `direct()` method returns:

```js
{
  ok: true,
  status: 200,
  headers: {},
  data: {}
}
```

On error, `ok` is `false` and an `err` property contains the error.

### FetchDef shape

The `prepare()` method returns:

```js
{
  url: 'string',
  method: 'string',
  headers: {},
  body: undefined
}
```

### Entities

#### Calendar

| Field | Description |
| --- | --- |
| `avatar_url` |  |
| `calendar_id` |  |
| `coordinate` |  |
| `cover_image_url` |  |
| `description` |  |
| `id` |  |
| `instagram_handle` |  |
| `is_personal` |  |
| `location` |  |
| `name` |  |
| `slug` |  |
| `social_image_url` |  |
| `tint_color` |  |
| `twitter_handle` |  |
| `url` |  |
| `website` |  |
| `youtube_handle` |  |

Operations: load, update.

API path: `/v1/calendars/get`

#### CalendarAdmin

| Field | Description |
| --- | --- |
| `avatar_url` |  |
| `email` |  |
| `first_name` |  |
| `id` |  |
| `last_name` |  |
| `name` |  |

Operations: list.

API path: `/v1/calendars/admins/list`

#### CalendarCoupon

| Field | Description |
| --- | --- |
| `cents_off` |  |
| `code` |  |
| `currency` |  |
| `discount` |  |
| `event_ticket_type_id` |  |
| `id` |  |
| `percent_off` |  |
| `remaining_count` |  |
| `valid_end_at` |  |
| `valid_start_at` |  |

Operations: create, list, update.

API path: `/v1/calendars/coupons/create`

#### CalendarEvent

| Field | Description |
| --- | --- |
| `id` |  |
| `status` |  |
| `submitted_by` |  |
| `tag` |  |

Operations: create, list, load.

API path: `/v1/calendars/events/add`

#### CalendarEventApproval

| Field | Description |
| --- | --- |
| `calendar_event_id` |  |

Operations: create.

API path: `/v1/calendars/events/approve`

#### CalendarEventRejection

| Field | Description |
| --- | --- |
| `calendar_event_id` |  |
| `message` |  |

Operations: create.

API path: `/v1/calendars/events/reject`

#### Contact

| Field | Description |
| --- | --- |
| `avatar_url` |  |
| `contact` |  |
| `created_at` |  |
| `email` |  |
| `event_approved_count` |  |
| `event_checked_in_count` |  |
| `first_name` |  |
| `id` |  |
| `last_name` |  |
| `membership` |  |
| `name` |  |
| `revenue_usd_cent` |  |
| `tag` |  |
| `user_id` |  |

Operations: create, list, remove.

API path: `/v1/calendars/contacts/import`

#### ContactBlock

| Field | Description |
| --- | --- |
| `contact_id` |  |
| `email` |  |

Operations: create.

API path: `/v1/calendars/contacts/block`

#### ContactRestore

| Field | Description |
| --- | --- |
| `contact_id` |  |
| `email` |  |

Operations: create.

API path: `/v1/calendars/contacts/restore`

#### ContactTag

| Field | Description |
| --- | --- |
| `color` |  |
| `id` |  |
| `name` |  |
| `tag_id` |  |

Operations: create, list, remove, update.

API path: `/v1/calendars/contact-tags/create`

#### ContactTagAssignment

| Field | Description |
| --- | --- |
| `applied_count` |  |
| `email` |  |
| `skipped_count` |  |
| `tag` |  |
| `user_id` |  |

Operations: create, remove.

API path: `/v1/calendars/contact-tags/apply`

#### EntityLookup

| Field | Description |
| --- | --- |

Operations: load.

API path: `/v1/entities/lookup`

#### Event

| Field | Description |
| --- | --- |
| `access` |  |
| `calendar_id` |  |
| `can_register_for_multiple_ticket` |  |
| `coordinate` |  |
| `cover_url` |  |
| `created_at` |  |
| `description` |  |
| `description_md` |  |
| `display_price` |  |
| `duration_interval` |  |
| `end_at` |  |
| `event_id` |  |
| `feedback_email` |  |
| `geo_address_json` |  |
| `guest_count` |  |
| `host` |  |
| `id` |  |
| `location_type` |  |
| `location_visibility` |  |
| `max_capacity` |  |
| `meeting_url` |  |
| `name` |  |
| `name_requirement` |  |
| `phone_number_requirement` |  |
| `platform` |  |
| `registration_open` |  |
| `registration_question` |  |
| `reminders_disabled` |  |
| `require_approval` |  |
| `show_guest_list` |  |
| `slug` |  |
| `spots_remaining` |  |
| `start_at` |  |
| `suppress_notification` |  |
| `timezone` |  |
| `tint_color` |  |
| `url` |  |
| `user_id` |  |
| `visibility` |  |
| `waitlist_status` |  |

Operations: create, load, remove, update.

API path: `/v1/events/create`

#### EventCancelRequest

| Field | Description |
| --- | --- |
| `cancellation_token` |  |
| `event_id` |  |
| `guest_count` |  |
| `is_paid` |  |

Operations: create.

API path: `/v1/events/cancel/request`

#### EventCoupon

| Field | Description |
| --- | --- |
| `cents_off` |  |
| `code` |  |
| `currency` |  |
| `discount` |  |
| `event_id` |  |
| `event_ticket_type_id` |  |
| `id` |  |
| `percent_off` |  |
| `remaining_count` |  |
| `valid_end_at` |  |
| `valid_start_at` |  |

Operations: create, list, update.

API path: `/v1/events/coupons/create`

#### EventTag

| Field | Description |
| --- | --- |
| `color` |  |
| `id` |  |
| `name` |  |
| `tag_id` |  |

Operations: create, list, remove, update.

API path: `/v1/calendars/event-tags/create`

#### EventTagAssignment

| Field | Description |
| --- | --- |
| `applied_count` |  |
| `event_id` |  |
| `skipped_count` |  |
| `tag` |  |

Operations: create, remove.

API path: `/v1/calendars/event-tags/apply`

#### Guest

| Field | Description |
| --- | --- |
| `approval_status` |  |
| `check_in_qr_code` |  |
| `eth_address` |  |
| `event_id` |  |
| `event_ticket` |  |
| `event_ticket_order` |  |
| `guest` |  |
| `guest_id` |  |
| `id` |  |
| `invited_at` |  |
| `joined_at` |  |
| `message` |  |
| `phone_number` |  |
| `registered_at` |  |
| `registration_answer` |  |
| `send_email` |  |
| `should_refund` |  |
| `solana_address` |  |
| `status` |  |
| `ticket` |  |
| `user_email` |  |
| `user_first_name` |  |
| `user_id` |  |
| `user_last_name` |  |
| `user_name` |  |
| `utm_source` |  |

Operations: create, list, load, update.

API path: `/v1/events/guests/add`

#### GuestInvite

| Field | Description |
| --- | --- |
| `event_id` |  |
| `guest` |  |
| `message` |  |

Operations: create.

API path: `/v1/events/guests/send-invites`

#### GuestTicket

| Field | Description |
| --- | --- |
| `event_id` |  |
| `guest_id` |  |
| `send_email` |  |
| `ticket_ids_to_remove` |  |
| `tickets_to_add` |  |

Operations: update.

API path: `/v1/events/guests/update-tickets`

#### Host

| Field | Description |
| --- | --- |
| `access_level` |  |
| `email` |  |
| `event_id` |  |
| `is_visible` |  |
| `name` |  |

Operations: create, remove, update.

API path: `/v1/events/hosts/add`

#### ImageUpload

| Field | Description |
| --- | --- |
| `content_type` |  |
| `file_url` |  |
| `upload_url` |  |

Operations: create.

API path: `/v1/images/create-upload-url`

#### Member

| Field | Description |
| --- | --- |
| `email` |  |
| `membership_id` |  |
| `membership_tier_id` |  |
| `registration_answer` |  |
| `skip_payment` |  |
| `status` |  |
| `user_id` |  |

Operations: create, update.

API path: `/v1/memberships/members/add`

#### MembershipTier

| Field | Description |
| --- | --- |
| `access_info` |  |
| `description` |  |
| `id` |  |
| `name` |  |
| `tint_color` |  |

Operations: list.

API path: `/v1/memberships/tiers/list`

#### OrganizationAdmin

| Field | Description |
| --- | --- |
| `api_id` |  |
| `avatar_url` |  |
| `email` |  |
| `first_name` |  |
| `id` |  |
| `last_name` |  |
| `name` |  |

Operations: list.

API path: `/v1/organizations/admins/list`

#### OrganizationCalendar

| Field | Description |
| --- | --- |
| `avatar_url` |  |
| `coordinate` |  |
| `cover_image_url` |  |
| `description` |  |
| `id` |  |
| `instagram_handle` |  |
| `is_personal` |  |
| `location` |  |
| `name` |  |
| `slug` |  |
| `social_image_url` |  |
| `tint_color` |  |
| `twitter_handle` |  |
| `url` |  |
| `website` |  |
| `youtube_handle` |  |

Operations: create, list.

API path: `/v2/organizations/calendars/create`

#### OrganizationEvent

| Field | Description |
| --- | --- |
| `api_id` |  |
| `calendar_api_id` |  |
| `calendar_id` |  |
| `coordinate` |  |
| `cover_url` |  |
| `created_at` |  |
| `display_price` |  |
| `duration_interval` |  |
| `end_at` |  |
| `feedback_email` |  |
| `geo_address_json` |  |
| `geo_latitude` |  |
| `geo_longitude` |  |
| `id` |  |
| `location_type` |  |
| `location_visibility` |  |
| `managing_calendar` |  |
| `meeting_url` |  |
| `name` |  |
| `platform` |  |
| `registration_open` |  |
| `registration_question` |  |
| `require_approval` |  |
| `spots_remaining` |  |
| `start_at` |  |
| `timezone` |  |
| `url` |  |
| `user_api_id` |  |
| `user_id` |  |
| `visibility` |  |
| `waitlist_status` |  |
| `zoom_meeting_url` |  |

Operations: list.

API path: `/v1/organizations/events/list`

#### OrganizationEventTransfer

| Field | Description |
| --- | --- |
| `calendar_id` |  |
| `event_id` |  |

Operations: create.

API path: `/v1/organizations/events/transfer-calendar`

#### TicketType

| Field | Description |
| --- | --- |
| `cent` |  |
| `currency` |  |
| `description` |  |
| `id` |  |
| `is_flexible` |  |
| `is_hidden` |  |
| `max_capacity` |  |
| `min_cent` |  |
| `name` |  |
| `require_approval` |  |
| `type` |  |
| `valid_end_at` |  |
| `valid_start_at` |  |

Operations: create, list, load, remove, update.

API path: `/v1/events/ticket-types/create`

#### User

| Field | Description |
| --- | --- |
| `avatar_url` |  |
| `email` |  |
| `first_name` |  |
| `id` |  |
| `last_name` |  |
| `name` |  |

Operations: load.

API path: `/v1/users/get-self`

#### Webhook

| Field | Description |
| --- | --- |
| `created_at` |  |
| `event_type` |  |
| `id` |  |
| `secret` |  |
| `status` |  |
| `url` |  |

Operations: create, list, load, remove, update.

API path: `/v2/webhooks/create`



## Entities


### Calendar

Create an instance: `const calendar = client.Calendar()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avatar_url` | `*` |  |
| `calendar_id` | `string` |  |
| `coordinate` | `*` |  |
| `cover_image_url` | `*` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `instagram_handle` | `*` |  |
| `is_personal` | `boolean` |  |
| `location` | `*` |  |
| `name` | `string` |  |
| `slug` | `string` |  |
| `social_image_url` | `*` |  |
| `tint_color` | `string` |  |
| `twitter_handle` | `*` |  |
| `url` | `string` |  |
| `website` | `*` |  |
| `youtube_handle` | `*` |  |

#### Example: Load

```ts
const calendar = await client.Calendar().load({ id: 'calendar_id' })
```


### CalendarAdmin

Create an instance: `const calendar_admin = client.CalendarAdmin()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avatar_url` | `string` |  |
| `email` | `string` |  |
| `first_name` | `*` |  |
| `id` | `string` |  |
| `last_name` | `*` |  |
| `name` | `string` |  |

#### Example: List

```ts
const calendar_admins = await client.CalendarAdmin().list()
```


### CalendarCoupon

Create an instance: `const calendar_coupon = client.CalendarCoupon()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `cents_off` | `*` |  |
| `code` | `string` |  |
| `currency` | `*` |  |
| `discount` | `*` |  |
| `event_ticket_type_id` | `string` |  |
| `id` | `string` |  |
| `percent_off` | `*` |  |
| `remaining_count` | `number` |  |
| `valid_end_at` | `*` |  |
| `valid_start_at` | `*` |  |

#### Example: List

```ts
const calendar_coupons = await client.CalendarCoupon().list()
```

#### Example: Create

```ts
const calendar_coupon = await client.CalendarCoupon().create({
  cents_off: 'example_cents_off',
  code: 'example_code',
  currency: 'example_currency',
  discount: 'example_discount',
  id: 'example_id',
  percent_off: 'example_percent_off',
  remaining_count: 1,
  valid_end_at: 'example_valid_end_at',
  valid_start_at: 'example_valid_start_at',
})
```


### CalendarEvent

Create an instance: `const calendar_event = client.CalendarEvent()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `string` |  |
| `status` | `string` |  |
| `submitted_by` | `*` |  |
| `tag` | `Array` |  |

#### Example: Load

```ts
const calendar_event = await client.CalendarEvent().load({ id: 'calendar_event_id' })
```

#### Example: List

```ts
const calendar_events = await client.CalendarEvent().list()
```

#### Example: Create

```ts
const calendar_event = await client.CalendarEvent().create({
  id: 'example_id',
  status: 'example_status',
  submitted_by: 'example_submitted_by',
  tag: [],
})
```


### CalendarEventApproval

Create an instance: `const calendar_event_approval = client.CalendarEventApproval()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `calendar_event_id` | `string` |  |

#### Example: Create

```ts
const calendar_event_approval = await client.CalendarEventApproval().create({
  calendar_event_id: 'example_calendar_event_id',
})
```


### CalendarEventRejection

Create an instance: `const calendar_event_rejection = client.CalendarEventRejection()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `calendar_event_id` | `string` |  |
| `message` | `string` |  |

#### Example: Create

```ts
const calendar_event_rejection = await client.CalendarEventRejection().create({
  calendar_event_id: 'example_calendar_event_id',
})
```


### Contact

Create an instance: `const contact = client.Contact()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avatar_url` | `string` |  |
| `contact` | `Array` |  |
| `created_at` | `string` |  |
| `email` | `string` |  |
| `event_approved_count` | `number` |  |
| `event_checked_in_count` | `number` |  |
| `first_name` | `*` |  |
| `id` | `string` |  |
| `last_name` | `*` |  |
| `membership` | `*` |  |
| `name` | `string` |  |
| `revenue_usd_cent` | `number` |  |
| `tag` | `*` |  |
| `user_id` | `string` |  |

#### Example: List

```ts
const contacts = await client.Contact().list()
```

#### Example: Create

```ts
const contact = await client.Contact().create({
  avatar_url: 'example_avatar_url',
  contact: [],
  created_at: 'example_created_at',
  email: 'example_email',
  event_approved_count: 1,
  event_checked_in_count: 1,
  first_name: 'example_first_name',
  id: 'example_id',
  last_name: 'example_last_name',
  membership: 'example_membership',
  name: 'example_name',
  revenue_usd_cent: 1,
  user_id: 'example_user_id',
})
```


### ContactBlock

Create an instance: `const contact_block = client.ContactBlock()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `contact_id` | `string` |  |
| `email` | `string` |  |

#### Example: Create

```ts
const contact_block = await client.ContactBlock().create({
})
```


### ContactRestore

Create an instance: `const contact_restore = client.ContactRestore()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `contact_id` | `string` |  |
| `email` | `string` |  |

#### Example: Create

```ts
const contact_restore = await client.ContactRestore().create({
})
```


### ContactTag

Create an instance: `const contact_tag = client.ContactTag()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `color` | `*` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `tag_id` | `string` |  |

#### Example: List

```ts
const contact_tags = await client.ContactTag().list()
```

#### Example: Create

```ts
const contact_tag = await client.ContactTag().create({
  id: 'example_id',
  name: 'example_name',
  tag_id: 'example_tag_id',
})
```


### ContactTagAssignment

Create an instance: `const contact_tag_assignment = client.ContactTagAssignment()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `applied_count` | `number` |  |
| `email` | `Array` |  |
| `skipped_count` | `number` |  |
| `tag` | `string` |  |
| `user_id` | `Array` |  |

#### Example: Create

```ts
const contact_tag_assignment = await client.ContactTagAssignment().create({
  applied_count: 1,
  skipped_count: 1,
  tag: 'example_tag',
})
```


### EntityLookup

Create an instance: `const entity_lookup = client.EntityLookup()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```ts
const entity_lookup = await client.EntityLookup().load()
```


### Event

Create an instance: `const event = client.Event()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `access` | `string` |  |
| `calendar_id` | `string` |  |
| `can_register_for_multiple_ticket` | `boolean` |  |
| `coordinate` | `*` |  |
| `cover_url` | `string` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `description_md` | `string` |  |
| `display_price` | `*` |  |
| `duration_interval` | `string` |  |
| `end_at` | `string` |  |
| `event_id` | `string` |  |
| `feedback_email` | `Object` |  |
| `geo_address_json` | `*` |  |
| `guest_count` | `Object` |  |
| `host` | `Array` |  |
| `id` | `string` |  |
| `location_type` | `string` |  |
| `location_visibility` | `string` |  |
| `max_capacity` | `*` |  |
| `meeting_url` | `*` |  |
| `name` | `string` |  |
| `name_requirement` | `string` |  |
| `phone_number_requirement` | `*` |  |
| `platform` | `string` |  |
| `registration_open` | `boolean` |  |
| `registration_question` | `Array` |  |
| `reminders_disabled` | `boolean` |  |
| `require_approval` | `boolean` |  |
| `show_guest_list` | `boolean` |  |
| `slug` | `string` |  |
| `spots_remaining` | `*` |  |
| `start_at` | `string` |  |
| `suppress_notification` | `boolean` |  |
| `timezone` | `string` |  |
| `tint_color` | `string` |  |
| `url` | `string` |  |
| `user_id` | `string` |  |
| `visibility` | `string` |  |
| `waitlist_status` | `string` |  |

#### Example: Load

```ts
const event = await client.Event().load({ id: 'event_id' })
```

#### Example: Create

```ts
const event = await client.Event().create({
  access: 'example_access',
  calendar_id: 'example_calendar_id',
  coordinate: 'example_coordinate',
  cover_url: 'example_cover_url',
  created_at: 'example_created_at',
  description: 'example_description',
  description_md: 'example_description_md',
  display_price: 'example_display_price',
  duration_interval: 'example_duration_interval',
  end_at: 'example_end_at',
  event_id: 'example_event_id',
  feedback_email: {},
  geo_address_json: 'example_geo_address_json',
  guest_count: {},
  host: [],
  id: 'example_id',
  location_type: 'example_location_type',
  location_visibility: 'example_location_visibility',
  meeting_url: 'example_meeting_url',
  name: 'example_name',
  platform: 'example_platform',
  registration_open: true,
  require_approval: true,
  spots_remaining: 'example_spots_remaining',
  start_at: 'example_start_at',
  timezone: 'example_timezone',
  url: 'example_url',
  user_id: 'example_user_id',
  visibility: 'example_visibility',
  waitlist_status: 'example_waitlist_status',
})
```


### EventCancelRequest

Create an instance: `const event_cancel_request = client.EventCancelRequest()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `cancellation_token` | `string` |  |
| `event_id` | `string` |  |
| `guest_count` | `number` |  |
| `is_paid` | `boolean` |  |

#### Example: Create

```ts
const event_cancel_request = await client.EventCancelRequest().create({
  cancellation_token: 'example_cancellation_token',
  event_id: 'example_event_id',
  guest_count: 1,
  is_paid: true,
})
```


### EventCoupon

Create an instance: `const event_coupon = client.EventCoupon()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `cents_off` | `*` |  |
| `code` | `string` |  |
| `currency` | `*` |  |
| `discount` | `*` |  |
| `event_id` | `string` |  |
| `event_ticket_type_id` | `string` |  |
| `id` | `string` |  |
| `percent_off` | `*` |  |
| `remaining_count` | `number` |  |
| `valid_end_at` | `*` |  |
| `valid_start_at` | `*` |  |

#### Example: List

```ts
const event_coupons = await client.EventCoupon().list()
```

#### Example: Create

```ts
const event_coupon = await client.EventCoupon().create({
  cents_off: 'example_cents_off',
  code: 'example_code',
  currency: 'example_currency',
  discount: 'example_discount',
  event_id: 'example_event_id',
  id: 'example_id',
  percent_off: 'example_percent_off',
  remaining_count: 1,
  valid_end_at: 'example_valid_end_at',
  valid_start_at: 'example_valid_start_at',
})
```


### EventTag

Create an instance: `const event_tag = client.EventTag()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `color` | `*` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `tag_id` | `string` |  |

#### Example: List

```ts
const event_tags = await client.EventTag().list()
```

#### Example: Create

```ts
const event_tag = await client.EventTag().create({
  id: 'example_id',
  name: 'example_name',
  tag_id: 'example_tag_id',
})
```


### EventTagAssignment

Create an instance: `const event_tag_assignment = client.EventTagAssignment()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `applied_count` | `number` |  |
| `event_id` | `Array` |  |
| `skipped_count` | `number` |  |
| `tag` | `string` |  |

#### Example: Create

```ts
const event_tag_assignment = await client.EventTagAssignment().create({
  applied_count: 1,
  event_id: [],
  skipped_count: 1,
  tag: 'example_tag',
})
```


### Guest

Create an instance: `const guest = client.Guest()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approval_status` | `string` |  |
| `check_in_qr_code` | `string` |  |
| `eth_address` | `*` |  |
| `event_id` | `string` |  |
| `event_ticket` | `Array` |  |
| `event_ticket_order` | `Array` |  |
| `guest` | `Array` |  |
| `guest_id` | `string` |  |
| `id` | `string` |  |
| `invited_at` | `*` |  |
| `joined_at` | `*` |  |
| `message` | `*` |  |
| `phone_number` | `number` |  |
| `registered_at` | `*` |  |
| `registration_answer` | `*` |  |
| `send_email` | `*` |  |
| `should_refund` | `boolean` |  |
| `solana_address` | `*` |  |
| `status` | `string` |  |
| `ticket` | `*` |  |
| `user_email` | `string` |  |
| `user_first_name` | `*` |  |
| `user_id` | `string` |  |
| `user_last_name` | `*` |  |
| `user_name` | `*` |  |
| `utm_source` | `*` |  |

#### Example: Load

```ts
const guest = await client.Guest().load({ id: 'guest_id' })
```

#### Example: List

```ts
const guests = await client.Guest().list()
```

#### Example: Create

```ts
const guest = await client.Guest().create({
  approval_status: 'example_approval_status',
  check_in_qr_code: 'example_check_in_qr_code',
  eth_address: 'example_eth_address',
  event_id: 'example_event_id',
  event_ticket: [],
  event_ticket_order: [],
  guest: [],
  guest_id: 'example_guest_id',
  id: 'example_id',
  invited_at: 'example_invited_at',
  joined_at: 'example_joined_at',
  phone_number: 1,
  registered_at: 'example_registered_at',
  registration_answer: 'example_registration_answer',
  solana_address: 'example_solana_address',
  status: 'example_status',
  user_email: 'example_user_email',
  user_first_name: 'example_user_first_name',
  user_id: 'example_user_id',
  user_last_name: 'example_user_last_name',
  user_name: 'example_user_name',
  utm_source: 'example_utm_source',
})
```


### GuestInvite

Create an instance: `const guest_invite = client.GuestInvite()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `event_id` | `string` |  |
| `guest` | `Array` |  |
| `message` | `*` |  |

#### Example: Create

```ts
const guest_invite = await client.GuestInvite().create({
  event_id: 'example_event_id',
  guest: [],
})
```


### GuestTicket

Create an instance: `const guest_ticket = client.GuestTicket()`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `event_id` | `string` |  |
| `guest_id` | `string` |  |
| `send_email` | `*` |  |
| `ticket_ids_to_remove` | `Array` |  |
| `tickets_to_add` | `Array` |  |


### Host

Create an instance: `const host = client.Host()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `access_level` | `*` |  |
| `email` | `string` |  |
| `event_id` | `string` |  |
| `is_visible` | `boolean` |  |
| `name` | `string` |  |

#### Example: Create

```ts
const host = await client.Host().create({
  email: 'example_email',
  event_id: 'example_event_id',
})
```


### ImageUpload

Create an instance: `const image_upload = client.ImageUpload()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `content_type` | `*` |  |
| `file_url` | `string` |  |
| `upload_url` | `string` |  |

#### Example: Create

```ts
const image_upload = await client.ImageUpload().create({
  file_url: 'example_file_url',
  upload_url: 'example_upload_url',
})
```


### Member

Create an instance: `const member = client.Member()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `email` | `string` |  |
| `membership_id` | `string` |  |
| `membership_tier_id` | `string` |  |
| `registration_answer` | `Array` |  |
| `skip_payment` | `boolean` |  |
| `status` | `string` |  |
| `user_id` | `string` |  |

#### Example: Create

```ts
const member = await client.Member().create({
  email: 'example_email',
  membership_id: 'example_membership_id',
  membership_tier_id: 'example_membership_tier_id',
  status: 'example_status',
  user_id: 'example_user_id',
})
```


### MembershipTier

Create an instance: `const membership_tier = client.MembershipTier()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `access_info` | `*` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `tint_color` | `string` |  |

#### Example: List

```ts
const membership_tiers = await client.MembershipTier().list()
```


### OrganizationAdmin

Create an instance: `const organization_admin = client.OrganizationAdmin()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_id` | `string` |  |
| `avatar_url` | `string` |  |
| `email` | `string` |  |
| `first_name` | `*` |  |
| `id` | `string` |  |
| `last_name` | `*` |  |
| `name` | `string` |  |

#### Example: List

```ts
const organization_admins = await client.OrganizationAdmin().list()
```


### OrganizationCalendar

Create an instance: `const organization_calendar = client.OrganizationCalendar()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avatar_url` | `*` |  |
| `coordinate` | `*` |  |
| `cover_image_url` | `*` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `instagram_handle` | `*` |  |
| `is_personal` | `boolean` |  |
| `location` | `*` |  |
| `name` | `string` |  |
| `slug` | `string` |  |
| `social_image_url` | `*` |  |
| `tint_color` | `string` |  |
| `twitter_handle` | `*` |  |
| `url` | `string` |  |
| `website` | `*` |  |
| `youtube_handle` | `*` |  |

#### Example: List

```ts
const organization_calendars = await client.OrganizationCalendar().list()
```

#### Example: Create

```ts
const organization_calendar = await client.OrganizationCalendar().create({
  avatar_url: 'example_avatar_url',
  coordinate: 'example_coordinate',
  cover_image_url: 'example_cover_image_url',
  description: 'example_description',
  id: 'example_id',
  instagram_handle: 'example_instagram_handle',
  is_personal: true,
  location: 'example_location',
  name: 'example_name',
  slug: 'example_slug',
  social_image_url: 'example_social_image_url',
  twitter_handle: 'example_twitter_handle',
  url: 'example_url',
  website: 'example_website',
  youtube_handle: 'example_youtube_handle',
})
```


### OrganizationEvent

Create an instance: `const organization_event = client.OrganizationEvent()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_id` | `string` |  |
| `calendar_api_id` | `string` |  |
| `calendar_id` | `string` |  |
| `coordinate` | `*` |  |
| `cover_url` | `string` |  |
| `created_at` | `string` |  |
| `display_price` | `*` |  |
| `duration_interval` | `string` |  |
| `end_at` | `string` |  |
| `feedback_email` | `Object` |  |
| `geo_address_json` | `*` |  |
| `geo_latitude` | `*` |  |
| `geo_longitude` | `*` |  |
| `id` | `string` |  |
| `location_type` | `string` |  |
| `location_visibility` | `string` |  |
| `managing_calendar` | `Array` |  |
| `meeting_url` | `*` |  |
| `name` | `string` |  |
| `platform` | `string` |  |
| `registration_open` | `boolean` |  |
| `registration_question` | `Array` |  |
| `require_approval` | `boolean` |  |
| `spots_remaining` | `*` |  |
| `start_at` | `string` |  |
| `timezone` | `string` |  |
| `url` | `string` |  |
| `user_api_id` | `string` |  |
| `user_id` | `string` |  |
| `visibility` | `string` |  |
| `waitlist_status` | `string` |  |
| `zoom_meeting_url` | `*` |  |

#### Example: List

```ts
const organization_events = await client.OrganizationEvent().list()
```


### OrganizationEventTransfer

Create an instance: `const organization_event_transfer = client.OrganizationEventTransfer()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `calendar_id` | `string` |  |
| `event_id` | `string` |  |

#### Example: Create

```ts
const organization_event_transfer = await client.OrganizationEventTransfer().create({
  calendar_id: 'example_calendar_id',
  event_id: 'example_event_id',
})
```


### TicketType

Create an instance: `const ticket_type = client.TicketType()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `cent` | `*` |  |
| `currency` | `*` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `is_flexible` | `boolean` |  |
| `is_hidden` | `boolean` |  |
| `max_capacity` | `*` |  |
| `min_cent` | `*` |  |
| `name` | `string` |  |
| `require_approval` | `boolean` |  |
| `type` | `string` |  |
| `valid_end_at` | `*` |  |
| `valid_start_at` | `*` |  |

#### Example: Load

```ts
const ticket_type = await client.TicketType().load({ id: 'ticket_type_id' })
```

#### Example: List

```ts
const ticket_types = await client.TicketType().list()
```

#### Example: Create

```ts
const ticket_type = await client.TicketType().create({
  id: 'example_id',
  name: 'example_name',
  type: 'example_type',
})
```


### User

Create an instance: `const user = client.User()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avatar_url` | `string` |  |
| `email` | `string` |  |
| `first_name` | `*` |  |
| `id` | `string` |  |
| `last_name` | `*` |  |
| `name` | `string` |  |

#### Example: Load

```ts
const user = await client.User().load({ id: 'user_id' })
```


### Webhook

Create an instance: `const webhook = client.Webhook()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `event_type` | `Array` |  |
| `id` | `string` |  |
| `secret` | `string` |  |
| `status` | `string` |  |
| `url` | `string` |  |

#### Example: Load

```ts
const webhook = await client.Webhook().load({ id: 'webhook_id' })
```

#### Example: List

```ts
const webhooks = await client.Webhook().list()
```

#### Example: Create

```ts
const webhook = await client.Webhook().create({
  created_at: 'example_created_at',
  event_type: [],
  id: 'example_id',
  secret: 'example_secret',
  status: 'example_status',
  url: 'example_url',
})
```


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature is an object with a
`hooks` map. Each hook key is a pipeline stage name, and the value is
a function that receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Module structure

```
luma/
├── src/
│   ├── LumaSDK.js        # Main SDK class
│   ├── entity/             # Entity implementations
│   ├── feature/            # Built-in features (Base, Test, Log)
│   └── utility/            # Utility functions
└── test/                   # Test suites
```

Import the SDK from the package root:

```js
const { LumaSDK } = require('@voxgig-sdk/luma-js')
```

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally. Subsequent
calls on the same instance can rely on this state.

```ts
const eventcoupon = client.EventCoupon()
await eventcoupon.list()

// eventcoupon.data() now returns the eventcoupon data from the last `list`
// eventcoupon.match() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

The `direct` method gives full control over the HTTP request. Use it
for non-standard endpoints, bulk operations, or any path not modelled
as an entity. The `prepare` method is useful for debugging — it
shows exactly what `direct` would send.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
