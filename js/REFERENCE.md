# Luma JavaScript SDK Reference

Complete API reference for the Luma JavaScript SDK.


## LumaSDK

### Constructor

```ts
new LumaSDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `LumaSDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = LumaSDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `LumaSDK` instance in test mode.


### Instance Methods

#### `Calendar(data?: object)`

Create a new `Calendar` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CalendarEntity` instance.

#### `CalendarAdmin(data?: object)`

Create a new `CalendarAdmin` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CalendarAdminEntity` instance.

#### `CalendarCoupon(data?: object)`

Create a new `CalendarCoupon` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CalendarCouponEntity` instance.

#### `CalendarEvent(data?: object)`

Create a new `CalendarEvent` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CalendarEventEntity` instance.

#### `CalendarEventApproval(data?: object)`

Create a new `CalendarEventApproval` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CalendarEventApprovalEntity` instance.

#### `CalendarEventRejection(data?: object)`

Create a new `CalendarEventRejection` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CalendarEventRejectionEntity` instance.

#### `Contact(data?: object)`

Create a new `Contact` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ContactEntity` instance.

#### `ContactBlock(data?: object)`

Create a new `ContactBlock` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ContactBlockEntity` instance.

#### `ContactRestore(data?: object)`

Create a new `ContactRestore` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ContactRestoreEntity` instance.

#### `ContactTag(data?: object)`

Create a new `ContactTag` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ContactTagEntity` instance.

#### `ContactTagAssignment(data?: object)`

Create a new `ContactTagAssignment` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ContactTagAssignmentEntity` instance.

#### `EntityLookup(data?: object)`

Create a new `EntityLookup` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EntityLookupEntity` instance.

#### `Event(data?: object)`

Create a new `Event` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EventEntity` instance.

#### `EventCancelRequest(data?: object)`

Create a new `EventCancelRequest` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EventCancelRequestEntity` instance.

#### `EventCoupon(data?: object)`

Create a new `EventCoupon` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EventCouponEntity` instance.

#### `EventTag(data?: object)`

Create a new `EventTag` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EventTagEntity` instance.

#### `EventTagAssignment(data?: object)`

Create a new `EventTagAssignment` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EventTagAssignmentEntity` instance.

#### `Guest(data?: object)`

Create a new `Guest` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `GuestEntity` instance.

#### `GuestInvite(data?: object)`

Create a new `GuestInvite` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `GuestInviteEntity` instance.

#### `GuestTicket(data?: object)`

Create a new `GuestTicket` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `GuestTicketEntity` instance.

#### `Host(data?: object)`

Create a new `Host` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `HostEntity` instance.

#### `ImageUpload(data?: object)`

Create a new `ImageUpload` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ImageUploadEntity` instance.

#### `Member(data?: object)`

Create a new `Member` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `MemberEntity` instance.

#### `MembershipTier(data?: object)`

Create a new `MembershipTier` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `MembershipTierEntity` instance.

#### `OrganizationAdmin(data?: object)`

Create a new `OrganizationAdmin` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `OrganizationAdminEntity` instance.

#### `OrganizationCalendar(data?: object)`

Create a new `OrganizationCalendar` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `OrganizationCalendarEntity` instance.

#### `OrganizationEvent(data?: object)`

Create a new `OrganizationEvent` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `OrganizationEventEntity` instance.

#### `OrganizationEventTransfer(data?: object)`

Create a new `OrganizationEventTransfer` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `OrganizationEventTransferEntity` instance.

#### `TicketType(data?: object)`

Create a new `TicketType` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `TicketTypeEntity` instance.

#### `User(data?: object)`

Create a new `User` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `UserEntity` instance.

#### `Webhook(data?: object)`

Create a new `Webhook` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `WebhookEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |

**Returns:** `Promise<{ ok, status, headers, data } | Error>`

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `LumaSDK.test()`.

**Returns:** `LumaSDK` instance in test mode.


---

## CalendarEntity

```ts
const calendar = client.Calendar()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avatar_url` | `*` | Yes |  |
| `calendar_id` | `string` | Yes |  |
| `coordinate` | `*` | Yes |  |
| `cover_image_url` | `*` | Yes |  |
| `description` | `string` | Yes |  |
| `id` | `string` | Yes |  |
| `instagram_handle` | `*` | Yes |  |
| `is_personal` | `boolean` | Yes |  |
| `location` | `*` | Yes |  |
| `name` | `string` | Yes |  |
| `slug` | `string` | Yes |  |
| `social_image_url` | `*` | Yes |  |
| `tint_color` | `string` | No |  |
| `twitter_handle` | `*` | Yes |  |
| `url` | `string` | Yes |  |
| `website` | `*` | Yes |  |
| `youtube_handle` | `*` | Yes |  |

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

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Calendar().load({ id: 'calendar_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Calendar().update({
  id: 'calendar_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CalendarEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CalendarAdminEntity

```ts
const calendar_admin = client.CalendarAdmin()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avatar_url` | `string` | Yes |  |
| `email` | `string` | Yes |  |
| `first_name` | `*` | Yes |  |
| `id` | `string` | Yes |  |
| `last_name` | `*` | Yes |  |
| `name` | `string` | Yes |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.CalendarAdmin().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CalendarAdminEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CalendarCouponEntity

```ts
const calendar_coupon = client.CalendarCoupon()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `cents_off` | `*` | Yes |  |
| `code` | `string` | Yes |  |
| `currency` | `*` | Yes |  |
| `discount` | `*` | Yes |  |
| `event_ticket_type_id` | `string` | No |  |
| `id` | `string` | Yes |  |
| `percent_off` | `*` | Yes |  |
| `remaining_count` | `number` | Yes |  |
| `valid_end_at` | `*` | Yes |  |
| `valid_start_at` | `*` | Yes |  |

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

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.CalendarCoupon().create({
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.CalendarCoupon().list()
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.CalendarCoupon().update({
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CalendarCouponEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CalendarEventEntity

```ts
const calendar_event = client.CalendarEvent()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `string` | Yes |  |
| `status` | `string` | Yes |  |
| `submitted_by` | `*` | Yes |  |
| `tag` | `Array` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.CalendarEvent().create({
  id: 'example_id',
  status: 'example_status',
  submitted_by: 'example_submitted_by',
  tag: [],
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.CalendarEvent().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.CalendarEvent().load({ id: 'calendar_event_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CalendarEventEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CalendarEventApprovalEntity

```ts
const calendar_event_approval = client.CalendarEventApproval()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `calendar_event_id` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.CalendarEventApproval().create({
  calendar_event_id: 'example_calendar_event_id',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CalendarEventApprovalEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CalendarEventRejectionEntity

```ts
const calendar_event_rejection = client.CalendarEventRejection()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `calendar_event_id` | `string` | Yes |  |
| `message` | `string` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.CalendarEventRejection().create({
  calendar_event_id: 'example_calendar_event_id',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CalendarEventRejectionEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ContactEntity

```ts
const contact = client.Contact()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avatar_url` | `string` | Yes |  |
| `contact` | `Array` | Yes |  |
| `created_at` | `string` | Yes |  |
| `email` | `string` | Yes |  |
| `event_approved_count` | `number` | Yes |  |
| `event_checked_in_count` | `number` | Yes |  |
| `first_name` | `*` | Yes |  |
| `id` | `string` | Yes |  |
| `last_name` | `*` | Yes |  |
| `membership` | `*` | Yes |  |
| `name` | `string` | Yes |  |
| `revenue_usd_cent` | `number` | Yes |  |
| `tag` | `*` | No |  |
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

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Contact().create({
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Contact().list()
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Contact().remove({ id: 'id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ContactEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ContactBlockEntity

```ts
const contact_block = client.ContactBlock()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `contact_id` | `string` | No |  |
| `email` | `string` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ContactBlock().create({
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ContactBlockEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ContactRestoreEntity

```ts
const contact_restore = client.ContactRestore()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `contact_id` | `string` | No |  |
| `email` | `string` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ContactRestore().create({
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ContactRestoreEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ContactTagEntity

```ts
const contact_tag = client.ContactTag()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `color` | `*` | No |  |
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

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ContactTag().create({
  id: 'example_id',
  name: 'example_name',
  tag_id: 'example_tag_id',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.ContactTag().list()
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.ContactTag().remove({ id: 'id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.ContactTag().update({
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ContactTagEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ContactTagAssignmentEntity

```ts
const contact_tag_assignment = client.ContactTagAssignment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `applied_count` | `number` | Yes |  |
| `email` | `Array` | No |  |
| `skipped_count` | `number` | Yes |  |
| `tag` | `string` | Yes |  |
| `user_id` | `Array` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ContactTagAssignment().create({
  applied_count: 1,
  skipped_count: 1,
  tag: 'example_tag',
})
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.ContactTagAssignment().remove()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ContactTagAssignmentEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EntityLookupEntity

```ts
const entity_lookup = client.EntityLookup()
```

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.EntityLookup().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EntityLookupEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EventEntity

```ts
const event = client.Event()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `access` | `string` | Yes |  |
| `calendar_id` | `string` | Yes |  |
| `can_register_for_multiple_ticket` | `boolean` | No |  |
| `coordinate` | `*` | Yes |  |
| `cover_url` | `string` | Yes |  |
| `created_at` | `string` | Yes |  |
| `description` | `string` | Yes |  |
| `description_md` | `string` | Yes |  |
| `display_price` | `*` | Yes |  |
| `duration_interval` | `string` | Yes |  |
| `end_at` | `string` | Yes |  |
| `event_id` | `string` | Yes |  |
| `feedback_email` | `Object` | Yes |  |
| `geo_address_json` | `*` | Yes |  |
| `guest_count` | `Object` | Yes |  |
| `host` | `Array` | Yes |  |
| `id` | `string` | Yes |  |
| `location_type` | `string` | Yes |  |
| `location_visibility` | `string` | Yes |  |
| `max_capacity` | `*` | No |  |
| `meeting_url` | `*` | Yes |  |
| `name` | `string` | Yes |  |
| `name_requirement` | `string` | No |  |
| `phone_number_requirement` | `*` | No |  |
| `platform` | `string` | Yes |  |
| `registration_open` | `boolean` | Yes |  |
| `registration_question` | `Array` | No |  |
| `reminders_disabled` | `boolean` | No |  |
| `require_approval` | `boolean` | Yes |  |
| `show_guest_list` | `boolean` | No |  |
| `slug` | `string` | No |  |
| `spots_remaining` | `*` | Yes |  |
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

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Event().create({
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

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Event().load({ id: 'event_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Event().remove({ id: 'event_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Event().update({
  id: 'event_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EventEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EventCancelRequestEntity

```ts
const event_cancel_request = client.EventCancelRequest()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `cancellation_token` | `string` | Yes |  |
| `event_id` | `string` | Yes |  |
| `guest_count` | `number` | Yes |  |
| `is_paid` | `boolean` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.EventCancelRequest().create({
  cancellation_token: 'example_cancellation_token',
  event_id: 'example_event_id',
  guest_count: 1,
  is_paid: true,
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EventCancelRequestEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EventCouponEntity

```ts
const event_coupon = client.EventCoupon()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `cents_off` | `*` | Yes |  |
| `code` | `string` | Yes |  |
| `currency` | `*` | Yes |  |
| `discount` | `*` | Yes |  |
| `event_id` | `string` | Yes |  |
| `event_ticket_type_id` | `string` | No |  |
| `id` | `string` | Yes |  |
| `percent_off` | `*` | Yes |  |
| `remaining_count` | `number` | Yes |  |
| `valid_end_at` | `*` | Yes |  |
| `valid_start_at` | `*` | Yes |  |

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

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.EventCoupon().create({
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.EventCoupon().list()
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.EventCoupon().update({
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EventCouponEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EventTagEntity

```ts
const event_tag = client.EventTag()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `color` | `*` | No |  |
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

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.EventTag().create({
  id: 'example_id',
  name: 'example_name',
  tag_id: 'example_tag_id',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.EventTag().list()
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.EventTag().remove({ id: 'id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.EventTag().update({
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EventTagEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EventTagAssignmentEntity

```ts
const event_tag_assignment = client.EventTagAssignment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `applied_count` | `number` | Yes |  |
| `event_id` | `Array` | Yes |  |
| `skipped_count` | `number` | Yes |  |
| `tag` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.EventTagAssignment().create({
  applied_count: 1,
  event_id: [],
  skipped_count: 1,
  tag: 'example_tag',
})
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.EventTagAssignment().remove()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EventTagAssignmentEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## GuestEntity

```ts
const guest = client.Guest()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_status` | `string` | Yes |  |
| `check_in_qr_code` | `string` | Yes |  |
| `eth_address` | `*` | Yes |  |
| `event_id` | `string` | Yes |  |
| `event_ticket` | `Array` | Yes |  |
| `event_ticket_order` | `Array` | Yes |  |
| `guest` | `Array` | Yes |  |
| `guest_id` | `string` | Yes |  |
| `id` | `string` | Yes |  |
| `invited_at` | `*` | Yes |  |
| `joined_at` | `*` | Yes |  |
| `message` | `*` | No |  |
| `phone_number` | `number` | Yes |  |
| `registered_at` | `*` | Yes |  |
| `registration_answer` | `*` | Yes |  |
| `send_email` | `*` | No |  |
| `should_refund` | `boolean` | No |  |
| `solana_address` | `*` | Yes |  |
| `status` | `string` | Yes |  |
| `ticket` | `*` | No |  |
| `user_email` | `string` | Yes |  |
| `user_first_name` | `*` | Yes |  |
| `user_id` | `string` | Yes |  |
| `user_last_name` | `*` | Yes |  |
| `user_name` | `*` | Yes |  |
| `utm_source` | `*` | Yes |  |

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

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Guest().create({
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Guest().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Guest().load({ id: 'guest_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Guest().update({
  id: 'guest_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `GuestEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## GuestInviteEntity

```ts
const guest_invite = client.GuestInvite()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `event_id` | `string` | Yes |  |
| `guest` | `Array` | Yes |  |
| `message` | `*` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.GuestInvite().create({
  event_id: 'example_event_id',
  guest: [],
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `GuestInviteEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## GuestTicketEntity

```ts
const guest_ticket = client.GuestTicket()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `event_id` | `string` | Yes |  |
| `guest_id` | `string` | Yes |  |
| `send_email` | `*` | No |  |
| `ticket_ids_to_remove` | `Array` | No |  |
| `tickets_to_add` | `Array` | No |  |

### Operations

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.GuestTicket().update({
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `GuestTicketEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## HostEntity

```ts
const host = client.Host()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `access_level` | `*` | No |  |
| `email` | `string` | Yes |  |
| `event_id` | `string` | Yes |  |
| `is_visible` | `boolean` | No |  |
| `name` | `string` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Host().create({
  email: 'example_email',
  event_id: 'example_event_id',
})
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Host().remove()
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Host().update({
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `HostEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ImageUploadEntity

```ts
const image_upload = client.ImageUpload()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `content_type` | `*` | No |  |
| `file_url` | `string` | Yes |  |
| `upload_url` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.ImageUpload().create({
  file_url: 'example_file_url',
  upload_url: 'example_upload_url',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ImageUploadEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## MemberEntity

```ts
const member = client.Member()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `email` | `string` | Yes |  |
| `membership_id` | `string` | Yes |  |
| `membership_tier_id` | `string` | Yes |  |
| `registration_answer` | `Array` | No |  |
| `skip_payment` | `boolean` | No |  |
| `status` | `string` | Yes |  |
| `user_id` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Member().create({
  email: 'example_email',
  membership_id: 'example_membership_id',
  membership_tier_id: 'example_membership_tier_id',
  status: 'example_status',
  user_id: 'example_user_id',
})
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Member().update({
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `MemberEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## MembershipTierEntity

```ts
const membership_tier = client.MembershipTier()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `access_info` | `*` | Yes |  |
| `description` | `string` | Yes |  |
| `id` | `string` | Yes |  |
| `name` | `string` | Yes |  |
| `tint_color` | `string` | Yes |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.MembershipTier().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `MembershipTierEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## OrganizationAdminEntity

```ts
const organization_admin = client.OrganizationAdmin()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_id` | `string` | Yes |  |
| `avatar_url` | `string` | Yes |  |
| `email` | `string` | Yes |  |
| `first_name` | `*` | Yes |  |
| `id` | `string` | Yes |  |
| `last_name` | `*` | Yes |  |
| `name` | `string` | Yes |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.OrganizationAdmin().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `OrganizationAdminEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## OrganizationCalendarEntity

```ts
const organization_calendar = client.OrganizationCalendar()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avatar_url` | `*` | Yes |  |
| `coordinate` | `*` | Yes |  |
| `cover_image_url` | `*` | Yes |  |
| `description` | `string` | Yes |  |
| `id` | `string` | Yes |  |
| `instagram_handle` | `*` | Yes |  |
| `is_personal` | `boolean` | Yes |  |
| `location` | `*` | Yes |  |
| `name` | `string` | Yes |  |
| `slug` | `string` | Yes |  |
| `social_image_url` | `*` | Yes |  |
| `tint_color` | `string` | No |  |
| `twitter_handle` | `*` | Yes |  |
| `url` | `string` | Yes |  |
| `website` | `*` | Yes |  |
| `youtube_handle` | `*` | Yes |  |

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

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.OrganizationCalendar().create({
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.OrganizationCalendar().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `OrganizationCalendarEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## OrganizationEventEntity

```ts
const organization_event = client.OrganizationEvent()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_id` | `string` | Yes |  |
| `calendar_api_id` | `string` | Yes |  |
| `calendar_id` | `string` | Yes |  |
| `coordinate` | `*` | Yes |  |
| `cover_url` | `string` | Yes |  |
| `created_at` | `string` | Yes |  |
| `display_price` | `*` | Yes |  |
| `duration_interval` | `string` | Yes |  |
| `end_at` | `string` | Yes |  |
| `feedback_email` | `Object` | Yes |  |
| `geo_address_json` | `*` | Yes |  |
| `geo_latitude` | `*` | Yes |  |
| `geo_longitude` | `*` | Yes |  |
| `id` | `string` | Yes |  |
| `location_type` | `string` | Yes |  |
| `location_visibility` | `string` | Yes |  |
| `managing_calendar` | `Array` | Yes |  |
| `meeting_url` | `*` | Yes |  |
| `name` | `string` | Yes |  |
| `platform` | `string` | Yes |  |
| `registration_open` | `boolean` | Yes |  |
| `registration_question` | `Array` | No |  |
| `require_approval` | `boolean` | Yes |  |
| `spots_remaining` | `*` | Yes |  |
| `start_at` | `string` | Yes |  |
| `timezone` | `string` | Yes |  |
| `url` | `string` | Yes |  |
| `user_api_id` | `string` | Yes |  |
| `user_id` | `string` | Yes |  |
| `visibility` | `string` | Yes |  |
| `waitlist_status` | `string` | Yes |  |
| `zoom_meeting_url` | `*` | Yes |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.OrganizationEvent().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `OrganizationEventEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## OrganizationEventTransferEntity

```ts
const organization_event_transfer = client.OrganizationEventTransfer()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `calendar_id` | `string` | Yes |  |
| `event_id` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.OrganizationEventTransfer().create({
  calendar_id: 'example_calendar_id',
  event_id: 'example_event_id',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `OrganizationEventTransferEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## TicketTypeEntity

```ts
const ticket_type = client.TicketType()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `cent` | `*` | No |  |
| `currency` | `*` | No |  |
| `description` | `string` | No |  |
| `id` | `string` | Yes |  |
| `is_flexible` | `boolean` | No |  |
| `is_hidden` | `boolean` | No |  |
| `max_capacity` | `*` | No |  |
| `min_cent` | `*` | No |  |
| `name` | `string` | Yes |  |
| `require_approval` | `boolean` | No |  |
| `type` | `string` | Yes |  |
| `valid_end_at` | `*` | No |  |
| `valid_start_at` | `*` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.TicketType().create({
  id: 'example_id',
  name: 'example_name',
  type: 'example_type',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.TicketType().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.TicketType().load({ id: 'ticket_type_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.TicketType().remove({ id: 'ticket_type_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.TicketType().update({
  id: 'ticket_type_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `TicketTypeEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## UserEntity

```ts
const user = client.User()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avatar_url` | `string` | Yes |  |
| `email` | `string` | Yes |  |
| `first_name` | `*` | Yes |  |
| `id` | `string` | Yes |  |
| `last_name` | `*` | Yes |  |
| `name` | `string` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.User().load({ id: 'user_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `UserEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## WebhookEntity

```ts
const webhook = client.Webhook()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | Yes |  |
| `event_type` | `Array` | Yes |  |
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

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Webhook().create({
  created_at: 'example_created_at',
  event_type: [],
  id: 'example_id',
  secret: 'example_secret',
  status: 'example_status',
  url: 'example_url',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Webhook().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Webhook().load({ id: 'webhook_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Webhook().remove({ id: 'webhook_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Webhook().update({
  id: 'webhook_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `WebhookEntity` instance with the same client and
options.

#### `client()`

Return the parent `LumaSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ts
const client = new LumaSDK({
  feature: {
    test: { active: true },
  }
})
```

