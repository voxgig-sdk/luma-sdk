# Luma PHP SDK Reference

Complete API reference for the Luma PHP SDK.


## LumaSDK

### Constructor

```php
require_once __DIR__ . '/luma_sdk.php';

$client = new LumaSDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["apikey"]` | `string` | API key for authentication. |
| `$options["base"]` | `string` | Base URL for API requests. |
| `$options["prefix"]` | `string` | URL prefix appended after base. |
| `$options["suffix"]` | `string` | URL suffix appended after path. |
| `$options["headers"]` | `array` | Custom headers for all requests. |
| `$options["feature"]` | `array` | Feature configuration. |
| `$options["system"]` | `array` | System overrides (e.g. custom fetch). |


### Static Methods

#### `LumaSDK::test($testopts = null, $sdkopts = null)`

Create a test client with mock features active. Both arguments may be `null`.

```php
$client = LumaSDK::test();
```


### Instance Methods

#### `Calendar($data = null)`

Create a new `CalendarEntity` instance. Pass `null` for no initial data.

#### `CalendarAdmin($data = null)`

Create a new `CalendarAdminEntity` instance. Pass `null` for no initial data.

#### `CalendarCoupon($data = null)`

Create a new `CalendarCouponEntity` instance. Pass `null` for no initial data.

#### `CalendarEvent($data = null)`

Create a new `CalendarEventEntity` instance. Pass `null` for no initial data.

#### `CalendarEventApproval($data = null)`

Create a new `CalendarEventApprovalEntity` instance. Pass `null` for no initial data.

#### `CalendarEventRejection($data = null)`

Create a new `CalendarEventRejectionEntity` instance. Pass `null` for no initial data.

#### `Contact($data = null)`

Create a new `ContactEntity` instance. Pass `null` for no initial data.

#### `ContactBlock($data = null)`

Create a new `ContactBlockEntity` instance. Pass `null` for no initial data.

#### `ContactRestore($data = null)`

Create a new `ContactRestoreEntity` instance. Pass `null` for no initial data.

#### `ContactTag($data = null)`

Create a new `ContactTagEntity` instance. Pass `null` for no initial data.

#### `ContactTagAssignment($data = null)`

Create a new `ContactTagAssignmentEntity` instance. Pass `null` for no initial data.

#### `EntityLookup($data = null)`

Create a new `EntityLookupEntity` instance. Pass `null` for no initial data.

#### `Event($data = null)`

Create a new `EventEntity` instance. Pass `null` for no initial data.

#### `EventCancelRequest($data = null)`

Create a new `EventCancelRequestEntity` instance. Pass `null` for no initial data.

#### `EventCoupon($data = null)`

Create a new `EventCouponEntity` instance. Pass `null` for no initial data.

#### `EventTag($data = null)`

Create a new `EventTagEntity` instance. Pass `null` for no initial data.

#### `EventTagAssignment($data = null)`

Create a new `EventTagAssignmentEntity` instance. Pass `null` for no initial data.

#### `Guest($data = null)`

Create a new `GuestEntity` instance. Pass `null` for no initial data.

#### `GuestInvite($data = null)`

Create a new `GuestInviteEntity` instance. Pass `null` for no initial data.

#### `GuestTicket($data = null)`

Create a new `GuestTicketEntity` instance. Pass `null` for no initial data.

#### `Host($data = null)`

Create a new `HostEntity` instance. Pass `null` for no initial data.

#### `ImageUpload($data = null)`

Create a new `ImageUploadEntity` instance. Pass `null` for no initial data.

#### `Member($data = null)`

Create a new `MemberEntity` instance. Pass `null` for no initial data.

#### `MembershipTier($data = null)`

Create a new `MembershipTierEntity` instance. Pass `null` for no initial data.

#### `OrganizationAdmin($data = null)`

Create a new `OrganizationAdminEntity` instance. Pass `null` for no initial data.

#### `OrganizationCalendar($data = null)`

Create a new `OrganizationCalendarEntity` instance. Pass `null` for no initial data.

#### `OrganizationEvent($data = null)`

Create a new `OrganizationEventEntity` instance. Pass `null` for no initial data.

#### `OrganizationEventTransfer($data = null)`

Create a new `OrganizationEventTransferEntity` instance. Pass `null` for no initial data.

#### `TicketType($data = null)`

Create a new `TicketTypeEntity` instance. Pass `null` for no initial data.

#### `User($data = null)`

Create a new `UserEntity` instance. Pass `null` for no initial data.

#### `Webhook($data = null)`

Create a new `WebhookEntity` instance. Pass `null` for no initial data.

#### `options_map(): array`

Return a deep copy of the current SDK options.

#### `get_utility(): LumaUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. This is the raw-HTTP escape
hatch: it does **not** throw. It returns a result array
`["ok" => bool, "status" => int, "headers" => array, "data" => mixed]`, or
`["ok" => false, "err" => \Exception]` on failure. Branch on `$result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `$fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `$fetchargs["params"]` | `array` | Path parameter values for `{param}` substitution. |
| `$fetchargs["query"]` | `array` | Query string parameters. |
| `$fetchargs["headers"]` | `array` | Request headers (merged with defaults). |
| `$fetchargs["body"]` | `mixed` | Request body (arrays are JSON-serialized). |
| `$fetchargs["ctrl"]` | `array` | Control options. |

**Returns:** `array` — the result dict (see above); never throws.

#### `prepare(array $fetchargs = []): mixed`

Prepare a fetch definition without sending the request. Returns the
`$fetchdef` array. Throws on error.


---

## CalendarEntity

```php
$calendar = $client->Calendar();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avatar_url` | `mixed` | Yes |  |
| `calendar_id` | `string` | Yes |  |
| `coordinate` | `mixed` | Yes |  |
| `cover_image_url` | `mixed` | Yes |  |
| `description` | `string` | Yes |  |
| `id` | `string` | Yes |  |
| `instagram_handle` | `mixed` | Yes |  |
| `is_personal` | `bool` | Yes |  |
| `location` | `mixed` | Yes |  |
| `name` | `string` | Yes |  |
| `slug` | `string` | Yes |  |
| `social_image_url` | `mixed` | Yes |  |
| `tint_color` | `string` | No |  |
| `twitter_handle` | `mixed` | Yes |  |
| `url` | `string` | Yes |  |
| `website` | `mixed` | Yes |  |
| `youtube_handle` | `mixed` | Yes |  |

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

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Calendar()->load(["id" => "calendar_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Calendar()->update([
  "id" => "calendar_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CalendarEntity`

Create a new `CalendarEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## CalendarAdminEntity

```php
$calendar_admin = $client->CalendarAdmin();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avatar_url` | `string` | Yes |  |
| `email` | `string` | Yes |  |
| `first_name` | `mixed` | Yes |  |
| `id` | `string` | Yes |  |
| `last_name` | `mixed` | Yes |  |
| `name` | `string` | Yes |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->CalendarAdmin()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CalendarAdminEntity`

Create a new `CalendarAdminEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## CalendarCouponEntity

```php
$calendar_coupon = $client->CalendarCoupon();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `cents_off` | `mixed` | Yes |  |
| `code` | `string` | Yes |  |
| `currency` | `mixed` | Yes |  |
| `discount` | `mixed` | Yes |  |
| `event_ticket_type_id` | `string` | No |  |
| `id` | `string` | Yes |  |
| `percent_off` | `mixed` | Yes |  |
| `remaining_count` | `int` | Yes |  |
| `valid_end_at` | `mixed` | Yes |  |
| `valid_start_at` | `mixed` | Yes |  |

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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->CalendarCoupon()->create([
  "cents_off" => null, // mixed
  "code" => null, // string
  "currency" => null, // mixed
  "discount" => null, // mixed
  "id" => null, // string
  "percent_off" => null, // mixed
  "remaining_count" => null, // int
  "valid_end_at" => null, // mixed
  "valid_start_at" => null, // mixed
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->CalendarCoupon()->list();
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->CalendarCoupon()->update([
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CalendarCouponEntity`

Create a new `CalendarCouponEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## CalendarEventEntity

```php
$calendar_event = $client->CalendarEvent();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `string` | Yes |  |
| `status` | `string` | Yes |  |
| `submitted_by` | `mixed` | Yes |  |
| `tag` | `array` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->CalendarEvent()->create([
  "id" => null, // string
  "status" => null, // string
  "submitted_by" => null, // mixed
  "tag" => null, // array
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->CalendarEvent()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->CalendarEvent()->load(["id" => "calendar_event_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CalendarEventEntity`

Create a new `CalendarEventEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## CalendarEventApprovalEntity

```php
$calendar_event_approval = $client->CalendarEventApproval();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `calendar_event_id` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->CalendarEventApproval()->create([
  "calendar_event_id" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CalendarEventApprovalEntity`

Create a new `CalendarEventApprovalEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## CalendarEventRejectionEntity

```php
$calendar_event_rejection = $client->CalendarEventRejection();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `calendar_event_id` | `string` | Yes |  |
| `message` | `string` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->CalendarEventRejection()->create([
  "calendar_event_id" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CalendarEventRejectionEntity`

Create a new `CalendarEventRejectionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ContactEntity

```php
$contact = $client->Contact();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avatar_url` | `string` | Yes |  |
| `contact` | `array` | Yes |  |
| `created_at` | `string` | Yes |  |
| `email` | `string` | Yes |  |
| `event_approved_count` | `float` | Yes |  |
| `event_checked_in_count` | `float` | Yes |  |
| `first_name` | `mixed` | Yes |  |
| `id` | `string` | Yes |  |
| `last_name` | `mixed` | Yes |  |
| `membership` | `mixed` | Yes |  |
| `name` | `string` | Yes |  |
| `revenue_usd_cent` | `float` | Yes |  |
| `tag` | `mixed` | No |  |
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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Contact()->create([
  "avatar_url" => null, // string
  "contact" => null, // array
  "created_at" => null, // string
  "email" => null, // string
  "event_approved_count" => null, // float
  "event_checked_in_count" => null, // float
  "first_name" => null, // mixed
  "id" => null, // string
  "last_name" => null, // mixed
  "membership" => null, // mixed
  "name" => null, // string
  "revenue_usd_cent" => null, // float
  "user_id" => null, // string
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Contact()->list();
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Contact()->remove(["id" => "id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ContactEntity`

Create a new `ContactEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ContactBlockEntity

```php
$contact_block = $client->ContactBlock();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `contact_id` | `string` | No |  |
| `email` | `string` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->ContactBlock()->create([
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ContactBlockEntity`

Create a new `ContactBlockEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ContactRestoreEntity

```php
$contact_restore = $client->ContactRestore();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `contact_id` | `string` | No |  |
| `email` | `string` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->ContactRestore()->create([
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ContactRestoreEntity`

Create a new `ContactRestoreEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ContactTagEntity

```php
$contact_tag = $client->ContactTag();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `color` | `mixed` | No |  |
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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->ContactTag()->create([
  "id" => null, // string
  "name" => null, // string
  "tag_id" => null, // string
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->ContactTag()->list();
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->ContactTag()->remove(["id" => "id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->ContactTag()->update([
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ContactTagEntity`

Create a new `ContactTagEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ContactTagAssignmentEntity

```php
$contact_tag_assignment = $client->ContactTagAssignment();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `applied_count` | `float` | Yes |  |
| `email` | `array` | No |  |
| `skipped_count` | `float` | Yes |  |
| `tag` | `string` | Yes |  |
| `user_id` | `array` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->ContactTagAssignment()->create([
  "applied_count" => null, // float
  "skipped_count" => null, // float
  "tag" => null, // string
]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->ContactTagAssignment()->remove();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ContactTagAssignmentEntity`

Create a new `ContactTagAssignmentEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EntityLookupEntity

```php
$entity_lookup = $client->EntityLookup();
```

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->EntityLookup()->load();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EntityLookupEntity`

Create a new `EntityLookupEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EventEntity

```php
$event = $client->Event();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `access` | `string` | Yes |  |
| `calendar_id` | `string` | Yes |  |
| `can_register_for_multiple_ticket` | `bool` | No |  |
| `coordinate` | `mixed` | Yes |  |
| `cover_url` | `string` | Yes |  |
| `created_at` | `string` | Yes |  |
| `description` | `string` | Yes |  |
| `description_md` | `string` | Yes |  |
| `display_price` | `mixed` | Yes |  |
| `duration_interval` | `string` | Yes |  |
| `end_at` | `string` | Yes |  |
| `event_id` | `string` | Yes |  |
| `feedback_email` | `array` | Yes |  |
| `geo_address_json` | `mixed` | Yes |  |
| `guest_count` | `array` | Yes |  |
| `host` | `array` | Yes |  |
| `id` | `string` | Yes |  |
| `location_type` | `string` | Yes |  |
| `location_visibility` | `string` | Yes |  |
| `max_capacity` | `mixed` | No |  |
| `meeting_url` | `mixed` | Yes |  |
| `name` | `string` | Yes |  |
| `name_requirement` | `string` | No |  |
| `phone_number_requirement` | `mixed` | No |  |
| `platform` | `string` | Yes |  |
| `registration_open` | `bool` | Yes |  |
| `registration_question` | `array` | No |  |
| `reminders_disabled` | `bool` | No |  |
| `require_approval` | `bool` | Yes |  |
| `show_guest_list` | `bool` | No |  |
| `slug` | `string` | No |  |
| `spots_remaining` | `mixed` | Yes |  |
| `start_at` | `string` | Yes |  |
| `suppress_notification` | `bool` | No |  |
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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Event()->create([
  "access" => null, // string
  "calendar_id" => null, // string
  "coordinate" => null, // mixed
  "cover_url" => null, // string
  "created_at" => null, // string
  "description" => null, // string
  "description_md" => null, // string
  "display_price" => null, // mixed
  "duration_interval" => null, // string
  "end_at" => null, // string
  "event_id" => null, // string
  "feedback_email" => null, // array
  "geo_address_json" => null, // mixed
  "guest_count" => null, // array
  "host" => null, // array
  "id" => null, // string
  "location_type" => null, // string
  "location_visibility" => null, // string
  "meeting_url" => null, // mixed
  "name" => null, // string
  "platform" => null, // string
  "registration_open" => null, // bool
  "require_approval" => null, // bool
  "spots_remaining" => null, // mixed
  "start_at" => null, // string
  "timezone" => null, // string
  "url" => null, // string
  "user_id" => null, // string
  "visibility" => null, // string
  "waitlist_status" => null, // string
]);
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Event()->load(["id" => "event_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Event()->remove(["id" => "event_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Event()->update([
  "id" => "event_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EventEntity`

Create a new `EventEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EventCancelRequestEntity

```php
$event_cancel_request = $client->EventCancelRequest();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `cancellation_token` | `string` | Yes |  |
| `event_id` | `string` | Yes |  |
| `guest_count` | `float` | Yes |  |
| `is_paid` | `bool` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->EventCancelRequest()->create([
  "cancellation_token" => null, // string
  "event_id" => null, // string
  "guest_count" => null, // float
  "is_paid" => null, // bool
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EventCancelRequestEntity`

Create a new `EventCancelRequestEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EventCouponEntity

```php
$event_coupon = $client->EventCoupon();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `cents_off` | `mixed` | Yes |  |
| `code` | `string` | Yes |  |
| `currency` | `mixed` | Yes |  |
| `discount` | `mixed` | Yes |  |
| `event_id` | `string` | Yes |  |
| `event_ticket_type_id` | `string` | No |  |
| `id` | `string` | Yes |  |
| `percent_off` | `mixed` | Yes |  |
| `remaining_count` | `int` | Yes |  |
| `valid_end_at` | `mixed` | Yes |  |
| `valid_start_at` | `mixed` | Yes |  |

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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->EventCoupon()->create([
  "cents_off" => null, // mixed
  "code" => null, // string
  "currency" => null, // mixed
  "discount" => null, // mixed
  "event_id" => null, // string
  "id" => null, // string
  "percent_off" => null, // mixed
  "remaining_count" => null, // int
  "valid_end_at" => null, // mixed
  "valid_start_at" => null, // mixed
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->EventCoupon()->list();
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->EventCoupon()->update([
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EventCouponEntity`

Create a new `EventCouponEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EventTagEntity

```php
$event_tag = $client->EventTag();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `color` | `mixed` | No |  |
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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->EventTag()->create([
  "id" => null, // string
  "name" => null, // string
  "tag_id" => null, // string
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->EventTag()->list();
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->EventTag()->remove(["id" => "id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->EventTag()->update([
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EventTagEntity`

Create a new `EventTagEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EventTagAssignmentEntity

```php
$event_tag_assignment = $client->EventTagAssignment();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `applied_count` | `float` | Yes |  |
| `event_id` | `array` | Yes |  |
| `skipped_count` | `float` | Yes |  |
| `tag` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->EventTagAssignment()->create([
  "applied_count" => null, // float
  "event_id" => null, // array
  "skipped_count" => null, // float
  "tag" => null, // string
]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->EventTagAssignment()->remove();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EventTagAssignmentEntity`

Create a new `EventTagAssignmentEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## GuestEntity

```php
$guest = $client->Guest();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_status` | `string` | Yes |  |
| `check_in_qr_code` | `string` | Yes |  |
| `eth_address` | `mixed` | Yes |  |
| `event_id` | `string` | Yes |  |
| `event_ticket` | `array` | Yes |  |
| `event_ticket_order` | `array` | Yes |  |
| `guest` | `array` | Yes |  |
| `guest_id` | `string` | Yes |  |
| `id` | `string` | Yes |  |
| `invited_at` | `mixed` | Yes |  |
| `joined_at` | `mixed` | Yes |  |
| `message` | `mixed` | No |  |
| `phone_number` | `int` | Yes |  |
| `registered_at` | `mixed` | Yes |  |
| `registration_answer` | `mixed` | Yes |  |
| `send_email` | `mixed` | No |  |
| `should_refund` | `bool` | No |  |
| `solana_address` | `mixed` | Yes |  |
| `status` | `string` | Yes |  |
| `ticket` | `mixed` | No |  |
| `user_email` | `string` | Yes |  |
| `user_first_name` | `mixed` | Yes |  |
| `user_id` | `string` | Yes |  |
| `user_last_name` | `mixed` | Yes |  |
| `user_name` | `mixed` | Yes |  |
| `utm_source` | `mixed` | Yes |  |

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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Guest()->create([
  "approval_status" => null, // string
  "check_in_qr_code" => null, // string
  "eth_address" => null, // mixed
  "event_id" => null, // string
  "event_ticket" => null, // array
  "event_ticket_order" => null, // array
  "guest" => null, // array
  "guest_id" => null, // string
  "id" => null, // string
  "invited_at" => null, // mixed
  "joined_at" => null, // mixed
  "phone_number" => null, // int
  "registered_at" => null, // mixed
  "registration_answer" => null, // mixed
  "solana_address" => null, // mixed
  "status" => null, // string
  "user_email" => null, // string
  "user_first_name" => null, // mixed
  "user_id" => null, // string
  "user_last_name" => null, // mixed
  "user_name" => null, // mixed
  "utm_source" => null, // mixed
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Guest()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Guest()->load(["id" => "guest_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Guest()->update([
  "id" => "guest_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): GuestEntity`

Create a new `GuestEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## GuestInviteEntity

```php
$guest_invite = $client->GuestInvite();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `event_id` | `string` | Yes |  |
| `guest` | `array` | Yes |  |
| `message` | `mixed` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->GuestInvite()->create([
  "event_id" => null, // string
  "guest" => null, // array
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): GuestInviteEntity`

Create a new `GuestInviteEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## GuestTicketEntity

```php
$guest_ticket = $client->GuestTicket();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `event_id` | `string` | Yes |  |
| `guest_id` | `string` | Yes |  |
| `send_email` | `mixed` | No |  |
| `ticket_ids_to_remove` | `array` | No |  |
| `tickets_to_add` | `array` | No |  |

### Operations

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->GuestTicket()->update([
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): GuestTicketEntity`

Create a new `GuestTicketEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## HostEntity

```php
$host = $client->Host();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `access_level` | `mixed` | No |  |
| `email` | `string` | Yes |  |
| `event_id` | `string` | Yes |  |
| `is_visible` | `bool` | No |  |
| `name` | `string` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Host()->create([
  "email" => null, // string
  "event_id" => null, // string
]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Host()->remove();
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Host()->update([
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): HostEntity`

Create a new `HostEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ImageUploadEntity

```php
$image_upload = $client->ImageUpload();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `content_type` | `mixed` | No |  |
| `file_url` | `string` | Yes |  |
| `upload_url` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->ImageUpload()->create([
  "file_url" => null, // string
  "upload_url" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ImageUploadEntity`

Create a new `ImageUploadEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## MemberEntity

```php
$member = $client->Member();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `email` | `string` | Yes |  |
| `membership_id` | `string` | Yes |  |
| `membership_tier_id` | `string` | Yes |  |
| `registration_answer` | `array` | No |  |
| `skip_payment` | `bool` | No |  |
| `status` | `string` | Yes |  |
| `user_id` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Member()->create([
  "email" => null, // string
  "membership_id" => null, // string
  "membership_tier_id" => null, // string
  "status" => null, // string
  "user_id" => null, // string
]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Member()->update([
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): MemberEntity`

Create a new `MemberEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## MembershipTierEntity

```php
$membership_tier = $client->MembershipTier();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `access_info` | `mixed` | Yes |  |
| `description` | `string` | Yes |  |
| `id` | `string` | Yes |  |
| `name` | `string` | Yes |  |
| `tint_color` | `string` | Yes |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->MembershipTier()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): MembershipTierEntity`

Create a new `MembershipTierEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## OrganizationAdminEntity

```php
$organization_admin = $client->OrganizationAdmin();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_id` | `string` | Yes |  |
| `avatar_url` | `string` | Yes |  |
| `email` | `string` | Yes |  |
| `first_name` | `mixed` | Yes |  |
| `id` | `string` | Yes |  |
| `last_name` | `mixed` | Yes |  |
| `name` | `string` | Yes |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->OrganizationAdmin()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): OrganizationAdminEntity`

Create a new `OrganizationAdminEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## OrganizationCalendarEntity

```php
$organization_calendar = $client->OrganizationCalendar();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avatar_url` | `mixed` | Yes |  |
| `coordinate` | `mixed` | Yes |  |
| `cover_image_url` | `mixed` | Yes |  |
| `description` | `string` | Yes |  |
| `id` | `string` | Yes |  |
| `instagram_handle` | `mixed` | Yes |  |
| `is_personal` | `bool` | Yes |  |
| `location` | `mixed` | Yes |  |
| `name` | `string` | Yes |  |
| `slug` | `string` | Yes |  |
| `social_image_url` | `mixed` | Yes |  |
| `tint_color` | `string` | No |  |
| `twitter_handle` | `mixed` | Yes |  |
| `url` | `string` | Yes |  |
| `website` | `mixed` | Yes |  |
| `youtube_handle` | `mixed` | Yes |  |

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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->OrganizationCalendar()->create([
  "avatar_url" => null, // mixed
  "coordinate" => null, // mixed
  "cover_image_url" => null, // mixed
  "description" => null, // string
  "id" => null, // string
  "instagram_handle" => null, // mixed
  "is_personal" => null, // bool
  "location" => null, // mixed
  "name" => null, // string
  "slug" => null, // string
  "social_image_url" => null, // mixed
  "twitter_handle" => null, // mixed
  "url" => null, // string
  "website" => null, // mixed
  "youtube_handle" => null, // mixed
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->OrganizationCalendar()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): OrganizationCalendarEntity`

Create a new `OrganizationCalendarEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## OrganizationEventEntity

```php
$organization_event = $client->OrganizationEvent();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_id` | `string` | Yes |  |
| `calendar_api_id` | `string` | Yes |  |
| `calendar_id` | `string` | Yes |  |
| `coordinate` | `mixed` | Yes |  |
| `cover_url` | `string` | Yes |  |
| `created_at` | `string` | Yes |  |
| `display_price` | `mixed` | Yes |  |
| `duration_interval` | `string` | Yes |  |
| `end_at` | `string` | Yes |  |
| `feedback_email` | `array` | Yes |  |
| `geo_address_json` | `mixed` | Yes |  |
| `geo_latitude` | `mixed` | Yes |  |
| `geo_longitude` | `mixed` | Yes |  |
| `id` | `string` | Yes |  |
| `location_type` | `string` | Yes |  |
| `location_visibility` | `string` | Yes |  |
| `managing_calendar` | `array` | Yes |  |
| `meeting_url` | `mixed` | Yes |  |
| `name` | `string` | Yes |  |
| `platform` | `string` | Yes |  |
| `registration_open` | `bool` | Yes |  |
| `registration_question` | `array` | No |  |
| `require_approval` | `bool` | Yes |  |
| `spots_remaining` | `mixed` | Yes |  |
| `start_at` | `string` | Yes |  |
| `timezone` | `string` | Yes |  |
| `url` | `string` | Yes |  |
| `user_api_id` | `string` | Yes |  |
| `user_id` | `string` | Yes |  |
| `visibility` | `string` | Yes |  |
| `waitlist_status` | `string` | Yes |  |
| `zoom_meeting_url` | `mixed` | Yes |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->OrganizationEvent()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): OrganizationEventEntity`

Create a new `OrganizationEventEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## OrganizationEventTransferEntity

```php
$organization_event_transfer = $client->OrganizationEventTransfer();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `calendar_id` | `string` | Yes |  |
| `event_id` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->OrganizationEventTransfer()->create([
  "calendar_id" => null, // string
  "event_id" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): OrganizationEventTransferEntity`

Create a new `OrganizationEventTransferEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## TicketTypeEntity

```php
$ticket_type = $client->TicketType();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `cent` | `mixed` | No |  |
| `currency` | `mixed` | No |  |
| `description` | `string` | No |  |
| `id` | `string` | Yes |  |
| `is_flexible` | `bool` | No |  |
| `is_hidden` | `bool` | No |  |
| `max_capacity` | `mixed` | No |  |
| `min_cent` | `mixed` | No |  |
| `name` | `string` | Yes |  |
| `require_approval` | `bool` | No |  |
| `type` | `string` | Yes |  |
| `valid_end_at` | `mixed` | No |  |
| `valid_start_at` | `mixed` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->TicketType()->create([
  "id" => null, // string
  "name" => null, // string
  "type" => null, // string
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->TicketType()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->TicketType()->load(["id" => "ticket_type_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->TicketType()->remove(["id" => "ticket_type_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->TicketType()->update([
  "id" => "ticket_type_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): TicketTypeEntity`

Create a new `TicketTypeEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## UserEntity

```php
$user = $client->User();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avatar_url` | `string` | Yes |  |
| `email` | `string` | Yes |  |
| `first_name` | `mixed` | Yes |  |
| `id` | `string` | Yes |  |
| `last_name` | `mixed` | Yes |  |
| `name` | `string` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->User()->load(["id" => "user_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): UserEntity`

Create a new `UserEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## WebhookEntity

```php
$webhook = $client->Webhook();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | Yes |  |
| `event_type` | `array` | Yes |  |
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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Webhook()->create([
  "created_at" => null, // string
  "event_type" => null, // array
  "id" => null, // string
  "secret" => null, // string
  "status" => null, // string
  "url" => null, // string
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Webhook()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Webhook()->load(["id" => "webhook_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Webhook()->remove(["id" => "webhook_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Webhook()->update([
  "id" => "webhook_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): WebhookEntity`

Create a new `WebhookEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```php
$client = new LumaSDK([
  "feature" => [
    "test" => ["active" => true],
  ],
]);
```

