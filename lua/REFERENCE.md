# Luma Lua SDK Reference

Complete API reference for the Luma Lua SDK.


## LumaSDK

### Constructor

```lua
local sdk = require("luma_sdk")
local client = sdk.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `table` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `table` | Custom headers for all requests. |
| `options.feature` | `table` | Feature configuration. |
| `options.system` | `table` | System overrides (e.g. custom fetch). |


### Static Methods

#### `sdk.test(testopts?, sdkopts?)`

Create a test client with mock features active. Both arguments are optional.

```lua
local client = sdk.test()
```


### Instance Methods

#### `Calendar(data)`

Create a new `Calendar` entity instance. Pass `nil` for no initial data.

#### `CalendarAdmin(data)`

Create a new `CalendarAdmin` entity instance. Pass `nil` for no initial data.

#### `CalendarCoupon(data)`

Create a new `CalendarCoupon` entity instance. Pass `nil` for no initial data.

#### `CalendarEvent(data)`

Create a new `CalendarEvent` entity instance. Pass `nil` for no initial data.

#### `CalendarEventApproval(data)`

Create a new `CalendarEventApproval` entity instance. Pass `nil` for no initial data.

#### `CalendarEventRejection(data)`

Create a new `CalendarEventRejection` entity instance. Pass `nil` for no initial data.

#### `Contact(data)`

Create a new `Contact` entity instance. Pass `nil` for no initial data.

#### `ContactBlock(data)`

Create a new `ContactBlock` entity instance. Pass `nil` for no initial data.

#### `ContactRestore(data)`

Create a new `ContactRestore` entity instance. Pass `nil` for no initial data.

#### `ContactTag(data)`

Create a new `ContactTag` entity instance. Pass `nil` for no initial data.

#### `ContactTagAssignment(data)`

Create a new `ContactTagAssignment` entity instance. Pass `nil` for no initial data.

#### `EntityLookup(data)`

Create a new `EntityLookup` entity instance. Pass `nil` for no initial data.

#### `Event(data)`

Create a new `Event` entity instance. Pass `nil` for no initial data.

#### `EventCancelRequest(data)`

Create a new `EventCancelRequest` entity instance. Pass `nil` for no initial data.

#### `EventCoupon(data)`

Create a new `EventCoupon` entity instance. Pass `nil` for no initial data.

#### `EventTag(data)`

Create a new `EventTag` entity instance. Pass `nil` for no initial data.

#### `EventTagAssignment(data)`

Create a new `EventTagAssignment` entity instance. Pass `nil` for no initial data.

#### `Guest(data)`

Create a new `Guest` entity instance. Pass `nil` for no initial data.

#### `GuestInvite(data)`

Create a new `GuestInvite` entity instance. Pass `nil` for no initial data.

#### `GuestTicket(data)`

Create a new `GuestTicket` entity instance. Pass `nil` for no initial data.

#### `Host(data)`

Create a new `Host` entity instance. Pass `nil` for no initial data.

#### `ImageUpload(data)`

Create a new `ImageUpload` entity instance. Pass `nil` for no initial data.

#### `Member(data)`

Create a new `Member` entity instance. Pass `nil` for no initial data.

#### `MembershipTier(data)`

Create a new `MembershipTier` entity instance. Pass `nil` for no initial data.

#### `OrganizationAdmin(data)`

Create a new `OrganizationAdmin` entity instance. Pass `nil` for no initial data.

#### `OrganizationCalendar(data)`

Create a new `OrganizationCalendar` entity instance. Pass `nil` for no initial data.

#### `OrganizationEvent(data)`

Create a new `OrganizationEvent` entity instance. Pass `nil` for no initial data.

#### `OrganizationEventTransfer(data)`

Create a new `OrganizationEventTransfer` entity instance. Pass `nil` for no initial data.

#### `TicketType(data)`

Create a new `TicketType` entity instance. Pass `nil` for no initial data.

#### `User(data)`

Create a new `User` entity instance. Pass `nil` for no initial data.

#### `Webhook(data)`

Create a new `Webhook` entity instance. Pass `nil` for no initial data.

#### `options_map() -> table`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs) -> table, err`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs.params` | `table` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `table` | Query string parameters. |
| `fetchargs.headers` | `table` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (tables are JSON-serialized). |
| `fetchargs.ctrl` | `table` | Control options (e.g. `{ explain = true }`). |

**Returns:** `table, err`

#### `prepare(fetchargs) -> table, err`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `table, err`


---

## CalendarEntity

```lua
local calendar = client:Calendar(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avatar_url` | `any` | Yes |  |
| `calendar_id` | `string` | Yes |  |
| `coordinate` | `any` | Yes |  |
| `cover_image_url` | `any` | Yes |  |
| `description` | `string` | Yes |  |
| `id` | `string` | Yes |  |
| `instagram_handle` | `any` | Yes |  |
| `is_personal` | `boolean` | Yes |  |
| `location` | `any` | Yes |  |
| `name` | `string` | Yes |  |
| `slug` | `string` | Yes |  |
| `social_image_url` | `any` | Yes |  |
| `tint_color` | `string` | No |  |
| `twitter_handle` | `any` | Yes |  |
| `url` | `string` | Yes |  |
| `website` | `any` | Yes |  |
| `youtube_handle` | `any` | Yes |  |

### Field Usage by Operation

| Field | load | update |
| --- | --- | --- |
| `avatar_url` | - | Yes |
| `calendar_id` | - | - |
| `coordinate` | - | - |
| `cover_image_url` | - | - |
| `description` | - | Yes |
| `id` | - | - |
| `instagram_handle` | - | - |
| `is_personal` | - | - |
| `location` | - | - |
| `name` | - | Yes |
| `slug` | - | Yes |
| `social_image_url` | - | - |
| `tint_color` | - | - |
| `twitter_handle` | - | - |
| `url` | - | - |
| `website` | - | - |
| `youtube_handle` | - | - |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Calendar():load({ id = "calendar_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Calendar():update({
  id = "calendar_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CalendarEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## CalendarAdminEntity

```lua
local calendar_admin = client:CalendarAdmin(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avatar_url` | `string` | Yes |  |
| `email` | `string` | Yes |  |
| `first_name` | `any` | Yes |  |
| `id` | `string` | Yes |  |
| `last_name` | `any` | Yes |  |
| `name` | `string` | Yes |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:CalendarAdmin():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CalendarAdminEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## CalendarCouponEntity

```lua
local calendar_coupon = client:CalendarCoupon(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `cents_off` | `any` | Yes |  |
| `code` | `string` | Yes |  |
| `currency` | `any` | Yes |  |
| `discount` | `any` | Yes |  |
| `event_ticket_type_id` | `string` | No |  |
| `id` | `string` | Yes |  |
| `percent_off` | `any` | Yes |  |
| `remaining_count` | `number` | Yes |  |
| `valid_end_at` | `any` | Yes |  |
| `valid_start_at` | `any` | Yes |  |

### Field Usage by Operation

| Field | list | create | update |
| --- | --- | --- | --- |
| `cents_off` | - | - | - |
| `code` | - | - | - |
| `currency` | - | - | - |
| `discount` | - | - | - |
| `event_ticket_type_id` | - | - | - |
| `id` | - | - | - |
| `percent_off` | - | - | - |
| `remaining_count` | - | Yes | Yes |
| `valid_end_at` | - | Yes | Yes |
| `valid_start_at` | - | Yes | Yes |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:CalendarCoupon():create({
  cents_off = --[[ any ]],
  code = --[[ string ]],
  currency = --[[ any ]],
  discount = --[[ any ]],
  id = --[[ string ]],
  percent_off = --[[ any ]],
  remaining_count = --[[ number ]],
  valid_end_at = --[[ any ]],
  valid_start_at = --[[ any ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:CalendarCoupon():list()
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:CalendarCoupon():update({
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CalendarCouponEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## CalendarEventEntity

```lua
local calendar_event = client:CalendarEvent(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `string` | Yes |  |
| `status` | `string` | Yes |  |
| `submitted_by` | `any` | Yes |  |
| `tag` | `table` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:CalendarEvent():create({
  id = --[[ string ]],
  status = --[[ string ]],
  submitted_by = --[[ any ]],
  tag = --[[ table ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:CalendarEvent():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:CalendarEvent():load({ id = "calendar_event_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CalendarEventEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## CalendarEventApprovalEntity

```lua
local calendar_event_approval = client:CalendarEventApproval(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `calendar_event_id` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:CalendarEventApproval():create({
  calendar_event_id = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CalendarEventApprovalEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## CalendarEventRejectionEntity

```lua
local calendar_event_rejection = client:CalendarEventRejection(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `calendar_event_id` | `string` | Yes |  |
| `message` | `string` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:CalendarEventRejection():create({
  calendar_event_id = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CalendarEventRejectionEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ContactEntity

```lua
local contact = client:Contact(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avatar_url` | `string` | Yes |  |
| `contact` | `table` | Yes |  |
| `created_at` | `string` | Yes |  |
| `email` | `string` | Yes |  |
| `event_approved_count` | `number` | Yes |  |
| `event_checked_in_count` | `number` | Yes |  |
| `first_name` | `any` | Yes |  |
| `id` | `string` | Yes |  |
| `last_name` | `any` | Yes |  |
| `membership` | `any` | Yes |  |
| `name` | `string` | Yes |  |
| `revenue_usd_cent` | `number` | Yes |  |
| `tag` | `any` | No |  |
| `user_id` | `string` | Yes |  |

### Field Usage by Operation

| Field | list | create | remove |
| --- | --- | --- | --- |
| `avatar_url` | - | - | - |
| `contact` | - | - | - |
| `created_at` | - | - | - |
| `email` | - | - | - |
| `event_approved_count` | - | - | - |
| `event_checked_in_count` | - | - | - |
| `first_name` | - | - | - |
| `id` | - | - | - |
| `last_name` | - | - | - |
| `membership` | - | - | - |
| `name` | - | - | - |
| `revenue_usd_cent` | - | - | - |
| `tag` | Yes | - | - |
| `user_id` | - | - | - |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Contact():create({
  avatar_url = --[[ string ]],
  contact = --[[ table ]],
  created_at = --[[ string ]],
  email = --[[ string ]],
  event_approved_count = --[[ number ]],
  event_checked_in_count = --[[ number ]],
  first_name = --[[ any ]],
  id = --[[ string ]],
  last_name = --[[ any ]],
  membership = --[[ any ]],
  name = --[[ string ]],
  revenue_usd_cent = --[[ number ]],
  user_id = --[[ string ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Contact():list()
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:Contact():remove({ id = "id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ContactEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ContactBlockEntity

```lua
local contact_block = client:ContactBlock(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `contact_id` | `string` | No |  |
| `email` | `string` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:ContactBlock():create({
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ContactBlockEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ContactRestoreEntity

```lua
local contact_restore = client:ContactRestore(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `contact_id` | `string` | No |  |
| `email` | `string` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:ContactRestore():create({
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ContactRestoreEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ContactTagEntity

```lua
local contact_tag = client:ContactTag(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `color` | `any` | No |  |
| `id` | `string` | Yes |  |
| `name` | `string` | Yes |  |
| `tag_id` | `string` | Yes |  |

### Field Usage by Operation

| Field | list | create | update | remove |
| --- | --- | --- | --- | --- |
| `color` | Yes | - | - | - |
| `id` | - | - | - | - |
| `name` | - | - | Yes | - |
| `tag_id` | - | - | - | - |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:ContactTag():create({
  id = --[[ string ]],
  name = --[[ string ]],
  tag_id = --[[ string ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:ContactTag():list()
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:ContactTag():remove({ id = "id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:ContactTag():update({
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ContactTagEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ContactTagAssignmentEntity

```lua
local contact_tag_assignment = client:ContactTagAssignment(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `applied_count` | `number` | Yes |  |
| `email` | `table` | No |  |
| `skipped_count` | `number` | Yes |  |
| `tag` | `string` | Yes |  |
| `user_id` | `table` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:ContactTagAssignment():create({
  applied_count = --[[ number ]],
  skipped_count = --[[ number ]],
  tag = --[[ string ]],
})
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:ContactTagAssignment():remove()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ContactTagAssignmentEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EntityLookupEntity

```lua
local entity_lookup = client:EntityLookup(nil)
```

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:EntityLookup():load()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EntityLookupEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EventEntity

```lua
local event = client:Event(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `access` | `string` | Yes |  |
| `calendar_id` | `string` | Yes |  |
| `can_register_for_multiple_ticket` | `boolean` | No |  |
| `coordinate` | `any` | Yes |  |
| `cover_url` | `string` | Yes |  |
| `created_at` | `string` | Yes |  |
| `description` | `string` | Yes |  |
| `description_md` | `string` | Yes |  |
| `display_price` | `any` | Yes |  |
| `duration_interval` | `string` | Yes |  |
| `end_at` | `string` | Yes |  |
| `event_id` | `string` | Yes |  |
| `feedback_email` | `table` | Yes |  |
| `geo_address_json` | `any` | Yes |  |
| `guest_count` | `table` | Yes |  |
| `host` | `table` | Yes |  |
| `id` | `string` | Yes |  |
| `location_type` | `string` | Yes |  |
| `location_visibility` | `string` | Yes |  |
| `max_capacity` | `any` | No |  |
| `meeting_url` | `any` | Yes |  |
| `name` | `string` | Yes |  |
| `name_requirement` | `string` | No |  |
| `phone_number_requirement` | `any` | No |  |
| `platform` | `string` | Yes |  |
| `registration_open` | `boolean` | Yes |  |
| `registration_question` | `table` | No |  |
| `reminders_disabled` | `boolean` | No |  |
| `require_approval` | `boolean` | Yes |  |
| `show_guest_list` | `boolean` | No |  |
| `slug` | `string` | No |  |
| `spots_remaining` | `any` | Yes |  |
| `start_at` | `string` | Yes |  |
| `suppress_notification` | `boolean` | No |  |
| `timezone` | `string` | Yes |  |
| `tint_color` | `string` | No |  |
| `url` | `string` | Yes |  |
| `user_id` | `string` | Yes |  |
| `visibility` | `string` | Yes |  |
| `waitlist_status` | `string` | Yes |  |

### Field Usage by Operation

| Field | load | create | update | remove |
| --- | --- | --- | --- | --- |
| `access` | - | - | - | - |
| `calendar_id` | - | - | - | - |
| `can_register_for_multiple_ticket` | - | - | - | - |
| `coordinate` | - | - | - | - |
| `cover_url` | - | Yes | Yes | - |
| `created_at` | - | - | - | - |
| `description` | - | - | - | - |
| `description_md` | - | Yes | Yes | - |
| `display_price` | - | - | - | - |
| `duration_interval` | - | - | - | - |
| `end_at` | - | Yes | Yes | - |
| `event_id` | - | - | - | - |
| `feedback_email` | - | - | - | - |
| `geo_address_json` | - | Yes | Yes | - |
| `guest_count` | - | - | - | - |
| `host` | - | - | - | - |
| `id` | - | - | - | - |
| `location_type` | - | - | - | - |
| `location_visibility` | - | Yes | Yes | - |
| `max_capacity` | - | - | - | - |
| `meeting_url` | - | Yes | Yes | - |
| `name` | - | - | Yes | - |
| `name_requirement` | - | - | - | - |
| `phone_number_requirement` | - | - | - | - |
| `platform` | - | - | - | - |
| `registration_open` | - | Yes | Yes | - |
| `registration_question` | - | - | - | - |
| `reminders_disabled` | - | - | - | - |
| `require_approval` | - | - | - | - |
| `show_guest_list` | - | - | - | - |
| `slug` | - | - | - | - |
| `spots_remaining` | - | - | - | - |
| `start_at` | - | - | Yes | - |
| `suppress_notification` | - | - | - | - |
| `timezone` | - | - | Yes | - |
| `tint_color` | - | - | - | - |
| `url` | - | - | - | - |
| `user_id` | - | - | - | - |
| `visibility` | - | Yes | Yes | - |
| `waitlist_status` | - | Yes | Yes | - |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Event():create({
  access = --[[ string ]],
  calendar_id = --[[ string ]],
  coordinate = --[[ any ]],
  cover_url = --[[ string ]],
  created_at = --[[ string ]],
  description = --[[ string ]],
  description_md = --[[ string ]],
  display_price = --[[ any ]],
  duration_interval = --[[ string ]],
  end_at = --[[ string ]],
  event_id = --[[ string ]],
  feedback_email = --[[ table ]],
  geo_address_json = --[[ any ]],
  guest_count = --[[ table ]],
  host = --[[ table ]],
  id = --[[ string ]],
  location_type = --[[ string ]],
  location_visibility = --[[ string ]],
  meeting_url = --[[ any ]],
  name = --[[ string ]],
  platform = --[[ string ]],
  registration_open = --[[ boolean ]],
  require_approval = --[[ boolean ]],
  spots_remaining = --[[ any ]],
  start_at = --[[ string ]],
  timezone = --[[ string ]],
  url = --[[ string ]],
  user_id = --[[ string ]],
  visibility = --[[ string ]],
  waitlist_status = --[[ string ]],
})
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Event():load({ id = "event_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:Event():remove({ id = "event_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Event():update({
  id = "event_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EventEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EventCancelRequestEntity

```lua
local event_cancel_request = client:EventCancelRequest(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `cancellation_token` | `string` | Yes |  |
| `event_id` | `string` | Yes |  |
| `guest_count` | `number` | Yes |  |
| `is_paid` | `boolean` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:EventCancelRequest():create({
  cancellation_token = --[[ string ]],
  event_id = --[[ string ]],
  guest_count = --[[ number ]],
  is_paid = --[[ boolean ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EventCancelRequestEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EventCouponEntity

```lua
local event_coupon = client:EventCoupon(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `cents_off` | `any` | Yes |  |
| `code` | `string` | Yes |  |
| `currency` | `any` | Yes |  |
| `discount` | `any` | Yes |  |
| `event_id` | `string` | Yes |  |
| `event_ticket_type_id` | `string` | No |  |
| `id` | `string` | Yes |  |
| `percent_off` | `any` | Yes |  |
| `remaining_count` | `number` | Yes |  |
| `valid_end_at` | `any` | Yes |  |
| `valid_start_at` | `any` | Yes |  |

### Field Usage by Operation

| Field | list | create | update |
| --- | --- | --- | --- |
| `cents_off` | - | - | - |
| `code` | - | - | - |
| `currency` | - | - | - |
| `discount` | - | - | - |
| `event_id` | - | - | - |
| `event_ticket_type_id` | - | - | - |
| `id` | - | - | - |
| `percent_off` | - | - | - |
| `remaining_count` | - | Yes | Yes |
| `valid_end_at` | - | Yes | Yes |
| `valid_start_at` | - | Yes | Yes |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:EventCoupon():create({
  cents_off = --[[ any ]],
  code = --[[ string ]],
  currency = --[[ any ]],
  discount = --[[ any ]],
  event_id = --[[ string ]],
  id = --[[ string ]],
  percent_off = --[[ any ]],
  remaining_count = --[[ number ]],
  valid_end_at = --[[ any ]],
  valid_start_at = --[[ any ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:EventCoupon():list()
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:EventCoupon():update({
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EventCouponEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EventTagEntity

```lua
local event_tag = client:EventTag(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `color` | `any` | No |  |
| `id` | `string` | Yes |  |
| `name` | `string` | Yes |  |
| `tag_id` | `string` | Yes |  |

### Field Usage by Operation

| Field | list | create | update | remove |
| --- | --- | --- | --- | --- |
| `color` | Yes | - | - | - |
| `id` | - | - | - | - |
| `name` | - | - | Yes | - |
| `tag_id` | - | - | - | - |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:EventTag():create({
  id = --[[ string ]],
  name = --[[ string ]],
  tag_id = --[[ string ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:EventTag():list()
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:EventTag():remove({ id = "id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:EventTag():update({
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EventTagEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EventTagAssignmentEntity

```lua
local event_tag_assignment = client:EventTagAssignment(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `applied_count` | `number` | Yes |  |
| `event_id` | `table` | Yes |  |
| `skipped_count` | `number` | Yes |  |
| `tag` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:EventTagAssignment():create({
  applied_count = --[[ number ]],
  event_id = --[[ table ]],
  skipped_count = --[[ number ]],
  tag = --[[ string ]],
})
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:EventTagAssignment():remove()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EventTagAssignmentEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## GuestEntity

```lua
local guest = client:Guest(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_status` | `string` | Yes |  |
| `check_in_qr_code` | `string` | Yes |  |
| `eth_address` | `any` | Yes |  |
| `event_id` | `string` | Yes |  |
| `event_ticket` | `table` | Yes |  |
| `event_ticket_order` | `table` | Yes |  |
| `guest` | `table` | Yes |  |
| `guest_id` | `string` | Yes |  |
| `id` | `string` | Yes |  |
| `invited_at` | `any` | Yes |  |
| `joined_at` | `any` | Yes |  |
| `message` | `any` | No |  |
| `phone_number` | `number` | Yes |  |
| `registered_at` | `any` | Yes |  |
| `registration_answer` | `any` | Yes |  |
| `send_email` | `any` | No |  |
| `should_refund` | `boolean` | No |  |
| `solana_address` | `any` | Yes |  |
| `status` | `string` | Yes |  |
| `ticket` | `any` | No |  |
| `user_email` | `string` | Yes |  |
| `user_first_name` | `any` | Yes |  |
| `user_id` | `string` | Yes |  |
| `user_last_name` | `any` | Yes |  |
| `user_name` | `any` | Yes |  |
| `utm_source` | `any` | Yes |  |

### Field Usage by Operation

| Field | load | list | create | update |
| --- | --- | --- | --- | --- |
| `approval_status` | - | - | Yes | - |
| `check_in_qr_code` | - | - | - | - |
| `eth_address` | - | - | - | - |
| `event_id` | - | - | - | - |
| `event_ticket` | - | - | - | - |
| `event_ticket_order` | - | - | - | - |
| `guest` | - | - | - | - |
| `guest_id` | - | - | - | - |
| `id` | - | - | - | - |
| `invited_at` | - | - | - | - |
| `joined_at` | - | - | - | - |
| `message` | - | - | - | - |
| `phone_number` | - | - | - | - |
| `registered_at` | - | - | - | - |
| `registration_answer` | - | - | - | - |
| `send_email` | - | - | - | - |
| `should_refund` | - | - | - | - |
| `solana_address` | - | - | - | - |
| `status` | - | - | - | - |
| `ticket` | - | - | - | - |
| `user_email` | - | - | - | - |
| `user_first_name` | - | - | - | - |
| `user_id` | - | - | - | - |
| `user_last_name` | - | - | - | - |
| `user_name` | - | - | - | - |
| `utm_source` | - | - | - | - |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Guest():create({
  approval_status = --[[ string ]],
  check_in_qr_code = --[[ string ]],
  eth_address = --[[ any ]],
  event_id = --[[ string ]],
  event_ticket = --[[ table ]],
  event_ticket_order = --[[ table ]],
  guest = --[[ table ]],
  guest_id = --[[ string ]],
  id = --[[ string ]],
  invited_at = --[[ any ]],
  joined_at = --[[ any ]],
  phone_number = --[[ number ]],
  registered_at = --[[ any ]],
  registration_answer = --[[ any ]],
  solana_address = --[[ any ]],
  status = --[[ string ]],
  user_email = --[[ string ]],
  user_first_name = --[[ any ]],
  user_id = --[[ string ]],
  user_last_name = --[[ any ]],
  user_name = --[[ any ]],
  utm_source = --[[ any ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Guest():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Guest():load({ id = "guest_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Guest():update({
  id = "guest_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `GuestEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## GuestInviteEntity

```lua
local guest_invite = client:GuestInvite(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `event_id` | `string` | Yes |  |
| `guest` | `table` | Yes |  |
| `message` | `any` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:GuestInvite():create({
  event_id = --[[ string ]],
  guest = --[[ table ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `GuestInviteEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## GuestTicketEntity

```lua
local guest_ticket = client:GuestTicket(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `event_id` | `string` | Yes |  |
| `guest_id` | `string` | Yes |  |
| `send_email` | `any` | No |  |
| `ticket_ids_to_remove` | `table` | No |  |
| `tickets_to_add` | `table` | No |  |

### Operations

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:GuestTicket():update({
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `GuestTicketEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## HostEntity

```lua
local host = client:Host(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `access_level` | `any` | No |  |
| `email` | `string` | Yes |  |
| `event_id` | `string` | Yes |  |
| `is_visible` | `boolean` | No |  |
| `name` | `string` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Host():create({
  email = --[[ string ]],
  event_id = --[[ string ]],
})
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:Host():remove()
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Host():update({
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `HostEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ImageUploadEntity

```lua
local image_upload = client:ImageUpload(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `content_type` | `any` | No |  |
| `file_url` | `string` | Yes |  |
| `upload_url` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:ImageUpload():create({
  file_url = --[[ string ]],
  upload_url = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ImageUploadEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## MemberEntity

```lua
local member = client:Member(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `email` | `string` | Yes |  |
| `membership_id` | `string` | Yes |  |
| `membership_tier_id` | `string` | Yes |  |
| `registration_answer` | `table` | No |  |
| `skip_payment` | `boolean` | No |  |
| `status` | `string` | Yes |  |
| `user_id` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Member():create({
  email = --[[ string ]],
  membership_id = --[[ string ]],
  membership_tier_id = --[[ string ]],
  status = --[[ string ]],
  user_id = --[[ string ]],
})
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Member():update({
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MemberEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## MembershipTierEntity

```lua
local membership_tier = client:MembershipTier(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `access_info` | `any` | Yes |  |
| `description` | `string` | Yes |  |
| `id` | `string` | Yes |  |
| `name` | `string` | Yes |  |
| `tint_color` | `string` | Yes |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:MembershipTier():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MembershipTierEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## OrganizationAdminEntity

```lua
local organization_admin = client:OrganizationAdmin(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_id` | `string` | Yes |  |
| `avatar_url` | `string` | Yes |  |
| `email` | `string` | Yes |  |
| `first_name` | `any` | Yes |  |
| `id` | `string` | Yes |  |
| `last_name` | `any` | Yes |  |
| `name` | `string` | Yes |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:OrganizationAdmin():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `OrganizationAdminEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## OrganizationCalendarEntity

```lua
local organization_calendar = client:OrganizationCalendar(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avatar_url` | `any` | Yes |  |
| `coordinate` | `any` | Yes |  |
| `cover_image_url` | `any` | Yes |  |
| `description` | `string` | Yes |  |
| `id` | `string` | Yes |  |
| `instagram_handle` | `any` | Yes |  |
| `is_personal` | `boolean` | Yes |  |
| `location` | `any` | Yes |  |
| `name` | `string` | Yes |  |
| `slug` | `string` | Yes |  |
| `social_image_url` | `any` | Yes |  |
| `tint_color` | `string` | No |  |
| `twitter_handle` | `any` | Yes |  |
| `url` | `string` | Yes |  |
| `website` | `any` | Yes |  |
| `youtube_handle` | `any` | Yes |  |

### Field Usage by Operation

| Field | list | create |
| --- | --- | --- |
| `avatar_url` | - | Yes |
| `coordinate` | - | - |
| `cover_image_url` | - | - |
| `description` | - | Yes |
| `id` | - | - |
| `instagram_handle` | - | - |
| `is_personal` | - | - |
| `location` | - | - |
| `name` | - | - |
| `slug` | - | Yes |
| `social_image_url` | - | - |
| `tint_color` | - | - |
| `twitter_handle` | - | - |
| `url` | - | - |
| `website` | - | - |
| `youtube_handle` | - | - |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:OrganizationCalendar():create({
  avatar_url = --[[ any ]],
  coordinate = --[[ any ]],
  cover_image_url = --[[ any ]],
  description = --[[ string ]],
  id = --[[ string ]],
  instagram_handle = --[[ any ]],
  is_personal = --[[ boolean ]],
  location = --[[ any ]],
  name = --[[ string ]],
  slug = --[[ string ]],
  social_image_url = --[[ any ]],
  twitter_handle = --[[ any ]],
  url = --[[ string ]],
  website = --[[ any ]],
  youtube_handle = --[[ any ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:OrganizationCalendar():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `OrganizationCalendarEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## OrganizationEventEntity

```lua
local organization_event = client:OrganizationEvent(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_id` | `string` | Yes |  |
| `calendar_api_id` | `string` | Yes |  |
| `calendar_id` | `string` | Yes |  |
| `coordinate` | `any` | Yes |  |
| `cover_url` | `string` | Yes |  |
| `created_at` | `string` | Yes |  |
| `display_price` | `any` | Yes |  |
| `duration_interval` | `string` | Yes |  |
| `end_at` | `string` | Yes |  |
| `feedback_email` | `table` | Yes |  |
| `geo_address_json` | `any` | Yes |  |
| `geo_latitude` | `any` | Yes |  |
| `geo_longitude` | `any` | Yes |  |
| `id` | `string` | Yes |  |
| `location_type` | `string` | Yes |  |
| `location_visibility` | `string` | Yes |  |
| `managing_calendar` | `table` | Yes |  |
| `meeting_url` | `any` | Yes |  |
| `name` | `string` | Yes |  |
| `platform` | `string` | Yes |  |
| `registration_open` | `boolean` | Yes |  |
| `registration_question` | `table` | No |  |
| `require_approval` | `boolean` | Yes |  |
| `spots_remaining` | `any` | Yes |  |
| `start_at` | `string` | Yes |  |
| `timezone` | `string` | Yes |  |
| `url` | `string` | Yes |  |
| `user_api_id` | `string` | Yes |  |
| `user_id` | `string` | Yes |  |
| `visibility` | `string` | Yes |  |
| `waitlist_status` | `string` | Yes |  |
| `zoom_meeting_url` | `any` | Yes |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:OrganizationEvent():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `OrganizationEventEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## OrganizationEventTransferEntity

```lua
local organization_event_transfer = client:OrganizationEventTransfer(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `calendar_id` | `string` | Yes |  |
| `event_id` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:OrganizationEventTransfer():create({
  calendar_id = --[[ string ]],
  event_id = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `OrganizationEventTransferEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## TicketTypeEntity

```lua
local ticket_type = client:TicketType(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `cent` | `any` | No |  |
| `currency` | `any` | No |  |
| `description` | `string` | No |  |
| `id` | `string` | Yes |  |
| `is_flexible` | `boolean` | No |  |
| `is_hidden` | `boolean` | No |  |
| `max_capacity` | `any` | No |  |
| `min_cent` | `any` | No |  |
| `name` | `string` | Yes |  |
| `require_approval` | `boolean` | No |  |
| `type` | `string` | Yes |  |
| `valid_end_at` | `any` | No |  |
| `valid_start_at` | `any` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:TicketType():create({
  id = --[[ string ]],
  name = --[[ string ]],
  type = --[[ string ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:TicketType():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:TicketType():load({ id = "ticket_type_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:TicketType():remove({ id = "ticket_type_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:TicketType():update({
  id = "ticket_type_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TicketTypeEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## UserEntity

```lua
local user = client:User(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avatar_url` | `string` | Yes |  |
| `email` | `string` | Yes |  |
| `first_name` | `any` | Yes |  |
| `id` | `string` | Yes |  |
| `last_name` | `any` | Yes |  |
| `name` | `string` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:User():load({ id = "user_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UserEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## WebhookEntity

```lua
local webhook = client:Webhook(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | Yes |  |
| `event_type` | `table` | Yes |  |
| `id` | `string` | Yes |  |
| `secret` | `string` | Yes |  |
| `status` | `string` | Yes |  |
| `url` | `string` | Yes |  |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `created_at` | - | - | - | - | - |
| `event_type` | - | - | - | Yes | - |
| `id` | - | - | - | - | - |
| `secret` | - | - | - | - | - |
| `status` | - | - | - | Yes | - |
| `url` | - | - | - | - | - |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Webhook():create({
  created_at = --[[ string ]],
  event_type = --[[ table ]],
  id = --[[ string ]],
  secret = --[[ string ]],
  status = --[[ string ]],
  url = --[[ string ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Webhook():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Webhook():load({ id = "webhook_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:Webhook():remove({ id = "webhook_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Webhook():update({
  id = "webhook_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `WebhookEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```lua
local client = sdk.new({
  feature = {
    test = { active = true },
  },
})
```

