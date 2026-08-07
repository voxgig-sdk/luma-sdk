# Luma Lua SDK



The Lua SDK for the Luma API — an entity-oriented client using Lua conventions.

It exposes the API as capitalised, semantic **Entities** — e.g. `client:Calendar()` — each with the same small set of operations (`list`, `load`, `create`, `update`, `remove`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to LuaRocks. Install it from the
GitHub release tag (`lua/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/luma-sdk/releases)),
or add the source directory to your `LUA_PATH`:

```bash
export LUA_PATH="path/to/lua/?.lua;path/to/lua/?/init.lua;;"
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```lua
local sdk = require("luma_sdk")

local client = sdk.new({
  apikey = os.getenv("LUMA_APIKEY"),
})
```

### 3. Load a calendar

```lua
local calendar, err = client:Calendar():load({ id = "example_id" })
if err then error(err) end
print(calendar)
```

### 4. Create, update, and remove

```lua
-- Update
client:Calendar():update({ id = "example_id", avatar_url = "example_avatar_url", calendar_id = "example_calendar_id" })

```


## Error handling

Entity operations return `(value, err)`. Check `err` before using
the value:

```lua
local eventcoupons, err = client:EventCoupon():list()
if err then error(err) end
```

`direct` follows the same `(value, err)` convention:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example_id" },
})
if err then error(err) end
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
if err then error(err) end

if result["ok"] then
  print(result["status"])  -- 200
  print(result["data"])    -- response body
end
```

### Prepare a request without sending it

```lua
local fetchdef, err = client:prepare({
  path = "/api/resource/{id}",
  method = "DELETE",
  params = { id = "example" },
})
if err then error(err) end

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```lua
local client = sdk.test()

local result, err = client:EventCoupon():list()
-- result is the returned data; err is set on failure
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```lua
local function mock_fetch(url, init)
  return {
    status = 200,
    statusText = "OK",
    headers = {},
    json = function()
      return { id = "mock01" }
    end,
  }, nil
end

local client = sdk.new({
  base = "http://localhost:8080",
  system = {
    fetch = mock_fetch,
  },
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
cd lua && busted test/
```


## Reference

### LumaSDK

```lua
local sdk = require("luma_sdk")
local client = sdk.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `table` | Feature activation flags. |
| `extend` | `table` | Additional Feature instances to load. |
| `system` | `table` | System overrides (e.g. custom `fetch` function). |

### test

```lua
local client = sdk.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### LumaSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> table` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> table, err` | Build an HTTP request definition without sending. |
| `direct` | `(fetchargs) -> table, err` | Build and send an HTTP request. |
| `Calendar` | `(data) -> CalendarEntity` | Create a Calendar entity instance. |
| `CalendarAdmin` | `(data) -> CalendarAdminEntity` | Create a CalendarAdmin entity instance. |
| `CalendarCoupon` | `(data) -> CalendarCouponEntity` | Create a CalendarCoupon entity instance. |
| `CalendarEvent` | `(data) -> CalendarEventEntity` | Create a CalendarEvent entity instance. |
| `CalendarEventApproval` | `(data) -> CalendarEventApprovalEntity` | Create a CalendarEventApproval entity instance. |
| `CalendarEventRejection` | `(data) -> CalendarEventRejectionEntity` | Create a CalendarEventRejection entity instance. |
| `Contact` | `(data) -> ContactEntity` | Create a Contact entity instance. |
| `ContactBlock` | `(data) -> ContactBlockEntity` | Create a ContactBlock entity instance. |
| `ContactRestore` | `(data) -> ContactRestoreEntity` | Create a ContactRestore entity instance. |
| `ContactTag` | `(data) -> ContactTagEntity` | Create a ContactTag entity instance. |
| `ContactTagAssignment` | `(data) -> ContactTagAssignmentEntity` | Create a ContactTagAssignment entity instance. |
| `EntityLookup` | `(data) -> EntityLookupEntity` | Create an EntityLookup entity instance. |
| `Event` | `(data) -> EventEntity` | Create an Event entity instance. |
| `EventCancelRequest` | `(data) -> EventCancelRequestEntity` | Create an EventCancelRequest entity instance. |
| `EventCoupon` | `(data) -> EventCouponEntity` | Create an EventCoupon entity instance. |
| `EventTag` | `(data) -> EventTagEntity` | Create an EventTag entity instance. |
| `EventTagAssignment` | `(data) -> EventTagAssignmentEntity` | Create an EventTagAssignment entity instance. |
| `Guest` | `(data) -> GuestEntity` | Create a Guest entity instance. |
| `GuestInvite` | `(data) -> GuestInviteEntity` | Create a GuestInvite entity instance. |
| `GuestTicket` | `(data) -> GuestTicketEntity` | Create a GuestTicket entity instance. |
| `Host` | `(data) -> HostEntity` | Create a Host entity instance. |
| `ImageUpload` | `(data) -> ImageUploadEntity` | Create an ImageUpload entity instance. |
| `Member` | `(data) -> MemberEntity` | Create a Member entity instance. |
| `MembershipTier` | `(data) -> MembershipTierEntity` | Create a MembershipTier entity instance. |
| `OrganizationAdmin` | `(data) -> OrganizationAdminEntity` | Create an OrganizationAdmin entity instance. |
| `OrganizationCalendar` | `(data) -> OrganizationCalendarEntity` | Create an OrganizationCalendar entity instance. |
| `OrganizationEvent` | `(data) -> OrganizationEventEntity` | Create an OrganizationEvent entity instance. |
| `OrganizationEventTransfer` | `(data) -> OrganizationEventTransferEntity` | Create an OrganizationEventTransfer entity instance. |
| `TicketType` | `(data) -> TicketTypeEntity` | Create a TicketType entity instance. |
| `User` | `(data) -> UserEntity` | Create an User entity instance. |
| `Webhook` | `(data) -> WebhookEntity` | Create a Webhook entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `(reqmatch, ctrl) -> any, err` | Load a single entity by match criteria. |
| `list` | `(reqmatch, ctrl) -> any, err` | List entities matching the criteria. |
| `create` | `(reqdata, ctrl) -> any, err` | Create a new entity. |
| `update` | `(reqdata, ctrl) -> any, err` | Update an existing entity. |
| `remove` | `(reqmatch, ctrl) -> any, err` | Remove an entity. |
| `data_get` | `() -> table` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> table` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> string` | Return the entity name. |

### Result shape

Entity operations return `(value, err)`. The `value` is the operation's
data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `load` / `create` / `update` / `remove` | the entity record (a `table`) |
| `list` | an array (`table`) of entity records |

Check `err` first (it is non-`nil` on failure), then use `value`:

    local calendar, err = client:Calendar():load({ id = "example_id" })
    if err then error(err) end
    -- calendar is the loaded record

Only `direct()` returns a response envelope — a `table` with `ok`,
`status`, `headers`, and `data` keys.

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

Operations: Load, Update.

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

Operations: List.

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

Operations: Create, List, Update.

API path: `/v1/calendars/coupons/create`

#### CalendarEvent

| Field | Description |
| --- | --- |
| `id` |  |
| `status` |  |
| `submitted_by` |  |
| `tag` |  |

Operations: Create, List, Load.

API path: `/v1/calendars/events/add`

#### CalendarEventApproval

| Field | Description |
| --- | --- |
| `calendar_event_id` |  |

Operations: Create.

API path: `/v1/calendars/events/approve`

#### CalendarEventRejection

| Field | Description |
| --- | --- |
| `calendar_event_id` |  |
| `message` |  |

Operations: Create.

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

Operations: Create, List, Remove.

API path: `/v1/calendars/contacts/import`

#### ContactBlock

| Field | Description |
| --- | --- |
| `contact_id` |  |
| `email` |  |

Operations: Create.

API path: `/v1/calendars/contacts/block`

#### ContactRestore

| Field | Description |
| --- | --- |
| `contact_id` |  |
| `email` |  |

Operations: Create.

API path: `/v1/calendars/contacts/restore`

#### ContactTag

| Field | Description |
| --- | --- |
| `color` |  |
| `id` |  |
| `name` |  |
| `tag_id` |  |

Operations: Create, List, Remove, Update.

API path: `/v1/calendars/contact-tags/create`

#### ContactTagAssignment

| Field | Description |
| --- | --- |
| `applied_count` |  |
| `email` |  |
| `skipped_count` |  |
| `tag` |  |
| `user_id` |  |

Operations: Create, Remove.

API path: `/v1/calendars/contact-tags/apply`

#### EntityLookup

| Field | Description |
| --- | --- |

Operations: Load.

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

Operations: Create, Load, Remove, Update.

API path: `/v1/events/create`

#### EventCancelRequest

| Field | Description |
| --- | --- |
| `cancellation_token` |  |
| `event_id` |  |
| `guest_count` |  |
| `is_paid` |  |

Operations: Create.

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

Operations: Create, List, Update.

API path: `/v1/events/coupons/create`

#### EventTag

| Field | Description |
| --- | --- |
| `color` |  |
| `id` |  |
| `name` |  |
| `tag_id` |  |

Operations: Create, List, Remove, Update.

API path: `/v1/calendars/event-tags/create`

#### EventTagAssignment

| Field | Description |
| --- | --- |
| `applied_count` |  |
| `event_id` |  |
| `skipped_count` |  |
| `tag` |  |

Operations: Create, Remove.

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

Operations: Create, List, Load, Update.

API path: `/v1/events/guests/add`

#### GuestInvite

| Field | Description |
| --- | --- |
| `event_id` |  |
| `guest` |  |
| `message` |  |

Operations: Create.

API path: `/v1/events/guests/send-invites`

#### GuestTicket

| Field | Description |
| --- | --- |
| `event_id` |  |
| `guest_id` |  |
| `send_email` |  |
| `ticket_ids_to_remove` |  |
| `tickets_to_add` |  |

Operations: Update.

API path: `/v1/events/guests/update-tickets`

#### Host

| Field | Description |
| --- | --- |
| `access_level` |  |
| `email` |  |
| `event_id` |  |
| `is_visible` |  |
| `name` |  |

Operations: Create, Remove, Update.

API path: `/v1/events/hosts/add`

#### ImageUpload

| Field | Description |
| --- | --- |
| `content_type` |  |
| `file_url` |  |
| `upload_url` |  |

Operations: Create.

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

Operations: Create, Update.

API path: `/v1/memberships/members/add`

#### MembershipTier

| Field | Description |
| --- | --- |
| `access_info` |  |
| `description` |  |
| `id` |  |
| `name` |  |
| `tint_color` |  |

Operations: List.

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

Operations: List.

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

Operations: Create, List.

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

Operations: List.

API path: `/v1/organizations/events/list`

#### OrganizationEventTransfer

| Field | Description |
| --- | --- |
| `calendar_id` |  |
| `event_id` |  |

Operations: Create.

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

Operations: Create, List, Load, Remove, Update.

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

Operations: Load.

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

Operations: Create, List, Load, Remove, Update.

API path: `/v2/webhooks/create`



## Entities


### Calendar

Create an instance: `local calendar = client:Calendar(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avatar_url` | `any` |  |
| `calendar_id` | `string` |  |
| `coordinate` | `any` |  |
| `cover_image_url` | `any` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `instagram_handle` | `any` |  |
| `is_personal` | `boolean` |  |
| `location` | `any` |  |
| `name` | `string` |  |
| `slug` | `string` |  |
| `social_image_url` | `any` |  |
| `tint_color` | `string` |  |
| `twitter_handle` | `any` |  |
| `url` | `string` |  |
| `website` | `any` |  |
| `youtube_handle` | `any` |  |

#### Example: Load

```lua
local calendar, err = client:Calendar():load({ id = "calendar_id" })
```


### CalendarAdmin

Create an instance: `local calendar_admin = client:CalendarAdmin(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avatar_url` | `string` |  |
| `email` | `string` |  |
| `first_name` | `any` |  |
| `id` | `string` |  |
| `last_name` | `any` |  |
| `name` | `string` |  |

#### Example: List

```lua
local calendar_admins, err = client:CalendarAdmin():list()
```


### CalendarCoupon

Create an instance: `local calendar_coupon = client:CalendarCoupon(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `cents_off` | `any` |  |
| `code` | `string` |  |
| `currency` | `any` |  |
| `discount` | `any` |  |
| `event_ticket_type_id` | `string` |  |
| `id` | `string` |  |
| `percent_off` | `any` |  |
| `remaining_count` | `number` |  |
| `valid_end_at` | `any` |  |
| `valid_start_at` | `any` |  |

#### Example: List

```lua
local calendar_coupons, err = client:CalendarCoupon():list()
```

#### Example: Create

```lua
local calendar_coupon, err = client:CalendarCoupon():create({
  cents_off = "example_cents_off", -- any
  code = "example_code", -- string
  currency = "example_currency", -- any
  discount = "example_discount", -- any
  id = "example_id", -- string
  percent_off = "example_percent_off", -- any
  remaining_count = 1, -- number
  valid_end_at = "example_valid_end_at", -- any
  valid_start_at = "example_valid_start_at", -- any
})
```


### CalendarEvent

Create an instance: `local calendar_event = client:CalendarEvent(nil)`

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
| `submitted_by` | `any` |  |
| `tag` | `table` |  |

#### Example: Load

```lua
local calendar_event, err = client:CalendarEvent():load({ id = "calendar_event_id" })
```

#### Example: List

```lua
local calendar_events, err = client:CalendarEvent():list()
```

#### Example: Create

```lua
local calendar_event, err = client:CalendarEvent():create({
  id = "example_id", -- string
  status = "example_status", -- string
  submitted_by = "example_submitted_by", -- any
  tag = {}, -- table
})
```


### CalendarEventApproval

Create an instance: `local calendar_event_approval = client:CalendarEventApproval(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `calendar_event_id` | `string` |  |

#### Example: Create

```lua
local calendar_event_approval, err = client:CalendarEventApproval():create({
  calendar_event_id = "example_calendar_event_id", -- string
})
```


### CalendarEventRejection

Create an instance: `local calendar_event_rejection = client:CalendarEventRejection(nil)`

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

```lua
local calendar_event_rejection, err = client:CalendarEventRejection():create({
  calendar_event_id = "example_calendar_event_id", -- string
})
```


### Contact

Create an instance: `local contact = client:Contact(nil)`

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
| `contact` | `table` |  |
| `created_at` | `string` |  |
| `email` | `string` |  |
| `event_approved_count` | `number` |  |
| `event_checked_in_count` | `number` |  |
| `first_name` | `any` |  |
| `id` | `string` |  |
| `last_name` | `any` |  |
| `membership` | `any` |  |
| `name` | `string` |  |
| `revenue_usd_cent` | `number` |  |
| `tag` | `any` |  |
| `user_id` | `string` |  |

#### Example: List

```lua
local contacts, err = client:Contact():list()
```

#### Example: Create

```lua
local contact, err = client:Contact():create({
  avatar_url = "example_avatar_url", -- string
  contact = {}, -- table
  created_at = "example_created_at", -- string
  email = "example_email", -- string
  event_approved_count = 1, -- number
  event_checked_in_count = 1, -- number
  first_name = "example_first_name", -- any
  id = "example_id", -- string
  last_name = "example_last_name", -- any
  membership = "example_membership", -- any
  name = "example_name", -- string
  revenue_usd_cent = 1, -- number
  user_id = "example_user_id", -- string
})
```


### ContactBlock

Create an instance: `local contact_block = client:ContactBlock(nil)`

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

```lua
local contact_block, err = client:ContactBlock():create({
})
```


### ContactRestore

Create an instance: `local contact_restore = client:ContactRestore(nil)`

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

```lua
local contact_restore, err = client:ContactRestore():create({
})
```


### ContactTag

Create an instance: `local contact_tag = client:ContactTag(nil)`

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
| `color` | `any` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `tag_id` | `string` |  |

#### Example: List

```lua
local contact_tags, err = client:ContactTag():list()
```

#### Example: Create

```lua
local contact_tag, err = client:ContactTag():create({
  id = "example_id", -- string
  name = "example_name", -- string
  tag_id = "example_tag_id", -- string
})
```


### ContactTagAssignment

Create an instance: `local contact_tag_assignment = client:ContactTagAssignment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `applied_count` | `number` |  |
| `email` | `table` |  |
| `skipped_count` | `number` |  |
| `tag` | `string` |  |
| `user_id` | `table` |  |

#### Example: Create

```lua
local contact_tag_assignment, err = client:ContactTagAssignment():create({
  applied_count = 1, -- number
  skipped_count = 1, -- number
  tag = "example_tag", -- string
})
```


### EntityLookup

Create an instance: `local entity_lookup = client:EntityLookup(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```lua
local entity_lookup, err = client:EntityLookup():load()
```


### Event

Create an instance: `local event = client:Event(nil)`

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
| `coordinate` | `any` |  |
| `cover_url` | `string` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `description_md` | `string` |  |
| `display_price` | `any` |  |
| `duration_interval` | `string` |  |
| `end_at` | `string` |  |
| `event_id` | `string` |  |
| `feedback_email` | `table` |  |
| `geo_address_json` | `any` |  |
| `guest_count` | `table` |  |
| `host` | `table` |  |
| `id` | `string` |  |
| `location_type` | `string` |  |
| `location_visibility` | `string` |  |
| `max_capacity` | `any` |  |
| `meeting_url` | `any` |  |
| `name` | `string` |  |
| `name_requirement` | `string` |  |
| `phone_number_requirement` | `any` |  |
| `platform` | `string` |  |
| `registration_open` | `boolean` |  |
| `registration_question` | `table` |  |
| `reminders_disabled` | `boolean` |  |
| `require_approval` | `boolean` |  |
| `show_guest_list` | `boolean` |  |
| `slug` | `string` |  |
| `spots_remaining` | `any` |  |
| `start_at` | `string` |  |
| `suppress_notification` | `boolean` |  |
| `timezone` | `string` |  |
| `tint_color` | `string` |  |
| `url` | `string` |  |
| `user_id` | `string` |  |
| `visibility` | `string` |  |
| `waitlist_status` | `string` |  |

#### Example: Load

```lua
local event, err = client:Event():load({ id = "event_id" })
```

#### Example: Create

```lua
local event, err = client:Event():create({
  access = "example_access", -- string
  calendar_id = "example_calendar_id", -- string
  coordinate = "example_coordinate", -- any
  cover_url = "example_cover_url", -- string
  created_at = "example_created_at", -- string
  description = "example_description", -- string
  description_md = "example_description_md", -- string
  display_price = "example_display_price", -- any
  duration_interval = "example_duration_interval", -- string
  end_at = "example_end_at", -- string
  event_id = "example_event_id", -- string
  feedback_email = {}, -- table
  geo_address_json = "example_geo_address_json", -- any
  guest_count = {}, -- table
  host = {}, -- table
  id = "example_id", -- string
  location_type = "example_location_type", -- string
  location_visibility = "example_location_visibility", -- string
  meeting_url = "example_meeting_url", -- any
  name = "example_name", -- string
  platform = "example_platform", -- string
  registration_open = true, -- boolean
  require_approval = true, -- boolean
  spots_remaining = "example_spots_remaining", -- any
  start_at = "example_start_at", -- string
  timezone = "example_timezone", -- string
  url = "example_url", -- string
  user_id = "example_user_id", -- string
  visibility = "example_visibility", -- string
  waitlist_status = "example_waitlist_status", -- string
})
```


### EventCancelRequest

Create an instance: `local event_cancel_request = client:EventCancelRequest(nil)`

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

```lua
local event_cancel_request, err = client:EventCancelRequest():create({
  cancellation_token = "example_cancellation_token", -- string
  event_id = "example_event_id", -- string
  guest_count = 1, -- number
  is_paid = true, -- boolean
})
```


### EventCoupon

Create an instance: `local event_coupon = client:EventCoupon(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `cents_off` | `any` |  |
| `code` | `string` |  |
| `currency` | `any` |  |
| `discount` | `any` |  |
| `event_id` | `string` |  |
| `event_ticket_type_id` | `string` |  |
| `id` | `string` |  |
| `percent_off` | `any` |  |
| `remaining_count` | `number` |  |
| `valid_end_at` | `any` |  |
| `valid_start_at` | `any` |  |

#### Example: List

```lua
local event_coupons, err = client:EventCoupon():list()
```

#### Example: Create

```lua
local event_coupon, err = client:EventCoupon():create({
  cents_off = "example_cents_off", -- any
  code = "example_code", -- string
  currency = "example_currency", -- any
  discount = "example_discount", -- any
  event_id = "example_event_id", -- string
  id = "example_id", -- string
  percent_off = "example_percent_off", -- any
  remaining_count = 1, -- number
  valid_end_at = "example_valid_end_at", -- any
  valid_start_at = "example_valid_start_at", -- any
})
```


### EventTag

Create an instance: `local event_tag = client:EventTag(nil)`

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
| `color` | `any` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `tag_id` | `string` |  |

#### Example: List

```lua
local event_tags, err = client:EventTag():list()
```

#### Example: Create

```lua
local event_tag, err = client:EventTag():create({
  id = "example_id", -- string
  name = "example_name", -- string
  tag_id = "example_tag_id", -- string
})
```


### EventTagAssignment

Create an instance: `local event_tag_assignment = client:EventTagAssignment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `applied_count` | `number` |  |
| `event_id` | `table` |  |
| `skipped_count` | `number` |  |
| `tag` | `string` |  |

#### Example: Create

```lua
local event_tag_assignment, err = client:EventTagAssignment():create({
  applied_count = 1, -- number
  event_id = {}, -- table
  skipped_count = 1, -- number
  tag = "example_tag", -- string
})
```


### Guest

Create an instance: `local guest = client:Guest(nil)`

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
| `eth_address` | `any` |  |
| `event_id` | `string` |  |
| `event_ticket` | `table` |  |
| `event_ticket_order` | `table` |  |
| `guest` | `table` |  |
| `guest_id` | `string` |  |
| `id` | `string` |  |
| `invited_at` | `any` |  |
| `joined_at` | `any` |  |
| `message` | `any` |  |
| `phone_number` | `number` |  |
| `registered_at` | `any` |  |
| `registration_answer` | `any` |  |
| `send_email` | `any` |  |
| `should_refund` | `boolean` |  |
| `solana_address` | `any` |  |
| `status` | `string` |  |
| `ticket` | `any` |  |
| `user_email` | `string` |  |
| `user_first_name` | `any` |  |
| `user_id` | `string` |  |
| `user_last_name` | `any` |  |
| `user_name` | `any` |  |
| `utm_source` | `any` |  |

#### Example: Load

```lua
local guest, err = client:Guest():load({ id = "guest_id" })
```

#### Example: List

```lua
local guests, err = client:Guest():list()
```

#### Example: Create

```lua
local guest, err = client:Guest():create({
  approval_status = "example_approval_status", -- string
  check_in_qr_code = "example_check_in_qr_code", -- string
  eth_address = "example_eth_address", -- any
  event_id = "example_event_id", -- string
  event_ticket = {}, -- table
  event_ticket_order = {}, -- table
  guest = {}, -- table
  guest_id = "example_guest_id", -- string
  id = "example_id", -- string
  invited_at = "example_invited_at", -- any
  joined_at = "example_joined_at", -- any
  phone_number = 1, -- number
  registered_at = "example_registered_at", -- any
  registration_answer = "example_registration_answer", -- any
  solana_address = "example_solana_address", -- any
  status = "example_status", -- string
  user_email = "example_user_email", -- string
  user_first_name = "example_user_first_name", -- any
  user_id = "example_user_id", -- string
  user_last_name = "example_user_last_name", -- any
  user_name = "example_user_name", -- any
  utm_source = "example_utm_source", -- any
})
```


### GuestInvite

Create an instance: `local guest_invite = client:GuestInvite(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `event_id` | `string` |  |
| `guest` | `table` |  |
| `message` | `any` |  |

#### Example: Create

```lua
local guest_invite, err = client:GuestInvite():create({
  event_id = "example_event_id", -- string
  guest = {}, -- table
})
```


### GuestTicket

Create an instance: `local guest_ticket = client:GuestTicket(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `event_id` | `string` |  |
| `guest_id` | `string` |  |
| `send_email` | `any` |  |
| `ticket_ids_to_remove` | `table` |  |
| `tickets_to_add` | `table` |  |


### Host

Create an instance: `local host = client:Host(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `access_level` | `any` |  |
| `email` | `string` |  |
| `event_id` | `string` |  |
| `is_visible` | `boolean` |  |
| `name` | `string` |  |

#### Example: Create

```lua
local host, err = client:Host():create({
  email = "example_email", -- string
  event_id = "example_event_id", -- string
})
```


### ImageUpload

Create an instance: `local image_upload = client:ImageUpload(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `content_type` | `any` |  |
| `file_url` | `string` |  |
| `upload_url` | `string` |  |

#### Example: Create

```lua
local image_upload, err = client:ImageUpload():create({
  file_url = "example_file_url", -- string
  upload_url = "example_upload_url", -- string
})
```


### Member

Create an instance: `local member = client:Member(nil)`

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
| `registration_answer` | `table` |  |
| `skip_payment` | `boolean` |  |
| `status` | `string` |  |
| `user_id` | `string` |  |

#### Example: Create

```lua
local member, err = client:Member():create({
  email = "example_email", -- string
  membership_id = "example_membership_id", -- string
  membership_tier_id = "example_membership_tier_id", -- string
  status = "example_status", -- string
  user_id = "example_user_id", -- string
})
```


### MembershipTier

Create an instance: `local membership_tier = client:MembershipTier(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `access_info` | `any` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `tint_color` | `string` |  |

#### Example: List

```lua
local membership_tiers, err = client:MembershipTier():list()
```


### OrganizationAdmin

Create an instance: `local organization_admin = client:OrganizationAdmin(nil)`

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
| `first_name` | `any` |  |
| `id` | `string` |  |
| `last_name` | `any` |  |
| `name` | `string` |  |

#### Example: List

```lua
local organization_admins, err = client:OrganizationAdmin():list()
```


### OrganizationCalendar

Create an instance: `local organization_calendar = client:OrganizationCalendar(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avatar_url` | `any` |  |
| `coordinate` | `any` |  |
| `cover_image_url` | `any` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `instagram_handle` | `any` |  |
| `is_personal` | `boolean` |  |
| `location` | `any` |  |
| `name` | `string` |  |
| `slug` | `string` |  |
| `social_image_url` | `any` |  |
| `tint_color` | `string` |  |
| `twitter_handle` | `any` |  |
| `url` | `string` |  |
| `website` | `any` |  |
| `youtube_handle` | `any` |  |

#### Example: List

```lua
local organization_calendars, err = client:OrganizationCalendar():list()
```

#### Example: Create

```lua
local organization_calendar, err = client:OrganizationCalendar():create({
  avatar_url = "example_avatar_url", -- any
  coordinate = "example_coordinate", -- any
  cover_image_url = "example_cover_image_url", -- any
  description = "example_description", -- string
  id = "example_id", -- string
  instagram_handle = "example_instagram_handle", -- any
  is_personal = true, -- boolean
  location = "example_location", -- any
  name = "example_name", -- string
  slug = "example_slug", -- string
  social_image_url = "example_social_image_url", -- any
  twitter_handle = "example_twitter_handle", -- any
  url = "example_url", -- string
  website = "example_website", -- any
  youtube_handle = "example_youtube_handle", -- any
})
```


### OrganizationEvent

Create an instance: `local organization_event = client:OrganizationEvent(nil)`

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
| `coordinate` | `any` |  |
| `cover_url` | `string` |  |
| `created_at` | `string` |  |
| `display_price` | `any` |  |
| `duration_interval` | `string` |  |
| `end_at` | `string` |  |
| `feedback_email` | `table` |  |
| `geo_address_json` | `any` |  |
| `geo_latitude` | `any` |  |
| `geo_longitude` | `any` |  |
| `id` | `string` |  |
| `location_type` | `string` |  |
| `location_visibility` | `string` |  |
| `managing_calendar` | `table` |  |
| `meeting_url` | `any` |  |
| `name` | `string` |  |
| `platform` | `string` |  |
| `registration_open` | `boolean` |  |
| `registration_question` | `table` |  |
| `require_approval` | `boolean` |  |
| `spots_remaining` | `any` |  |
| `start_at` | `string` |  |
| `timezone` | `string` |  |
| `url` | `string` |  |
| `user_api_id` | `string` |  |
| `user_id` | `string` |  |
| `visibility` | `string` |  |
| `waitlist_status` | `string` |  |
| `zoom_meeting_url` | `any` |  |

#### Example: List

```lua
local organization_events, err = client:OrganizationEvent():list()
```


### OrganizationEventTransfer

Create an instance: `local organization_event_transfer = client:OrganizationEventTransfer(nil)`

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

```lua
local organization_event_transfer, err = client:OrganizationEventTransfer():create({
  calendar_id = "example_calendar_id", -- string
  event_id = "example_event_id", -- string
})
```


### TicketType

Create an instance: `local ticket_type = client:TicketType(nil)`

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
| `cent` | `any` |  |
| `currency` | `any` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `is_flexible` | `boolean` |  |
| `is_hidden` | `boolean` |  |
| `max_capacity` | `any` |  |
| `min_cent` | `any` |  |
| `name` | `string` |  |
| `require_approval` | `boolean` |  |
| `type` | `string` |  |
| `valid_end_at` | `any` |  |
| `valid_start_at` | `any` |  |

#### Example: Load

```lua
local ticket_type, err = client:TicketType():load({ id = "ticket_type_id" })
```

#### Example: List

```lua
local ticket_types, err = client:TicketType():list()
```

#### Example: Create

```lua
local ticket_type, err = client:TicketType():create({
  id = "example_id", -- string
  name = "example_name", -- string
  type = "example_type", -- string
})
```


### User

Create an instance: `local user = client:User(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avatar_url` | `string` |  |
| `email` | `string` |  |
| `first_name` | `any` |  |
| `id` | `string` |  |
| `last_name` | `any` |  |
| `name` | `string` |  |

#### Example: Load

```lua
local user, err = client:User():load({ id = "user_id" })
```


### Webhook

Create an instance: `local webhook = client:Webhook(nil)`

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
| `event_type` | `table` |  |
| `id` | `string` |  |
| `secret` | `string` |  |
| `status` | `string` |  |
| `url` | `string` |  |

#### Example: Load

```lua
local webhook, err = client:Webhook():load({ id = "webhook_id" })
```

#### Example: List

```lua
local webhooks, err = client:Webhook():list()
```

#### Example: Create

```lua
local webhook, err = client:Webhook():create({
  created_at = "example_created_at", -- string
  event_type = {}, -- table
  id = "example_id", -- string
  secret = "example_secret", -- string
  status = "example_status", -- string
  url = "example_url", -- string
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

Features are the extension mechanism. A feature is a Lua table
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as tables

The Lua SDK uses plain Lua tables throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a table.

### Module structure

```
lua/
├── luma_sdk.lua    -- Main SDK module
├── config.lua               -- Configuration
├── features.lua             -- Feature factory
├── core/                    -- Core types and context
├── entity/                  -- Entity implementations
├── feature/                 -- Built-in features (Base, Test, Log)
├── utility/                 -- Utility functions and struct library
└── test/                    -- Test suites
```

The main module (`luma_sdk`) exports the SDK constructor
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```lua
local eventcoupon = client:EventCoupon()
eventcoupon:list()

-- eventcoupon:data_get() now returns the eventcoupon data from the last list
-- eventcoupon:match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
