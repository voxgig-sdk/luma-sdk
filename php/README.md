# Luma PHP SDK



The PHP SDK for the Luma API — an entity-oriented client using PHP conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `$client->Calendar()` — with named operations (`list`/`load`/`create`/`update`/`remove`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to Packagist. Install it from the
GitHub release tag (`php/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/luma-sdk/releases](https://github.com/voxgig-sdk/luma-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```php
<?php
require_once 'luma_sdk.php';

$client = new LumaSDK([
    "apikey" => getenv("LUMA_APIKEY"),
]);
```

### 3. Load a calendar

```php
try {
    // load() returns the bare Calendar record (throws on error).
    $calendar = $client->Calendar()->load(["id" => "example_id"]);
    print_r($calendar);
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

### 4. Create, update, and remove

```php
// Update
$client->Calendar()->update(["id" => "example_id", "avatar_url" => "example_avatar_url", "calendar_id" => "example_calendar_id"]);

```


## Error handling

Entity operations throw a `\Throwable` on failure, so wrap them in
`try` / `catch`:

```php
try {
    $eventcoupons = $client->EventCoupon()->list();
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

`direct()` does **not** throw — it returns the result array. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```php
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example_id"],
]);

if (! $result["ok"]) {
    $err = $result["err"] ?? null;
    echo "request failed: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```php
// direct() is the raw-HTTP escape hatch: it returns a result array
// (it does not throw). Branch on $result["ok"].
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);

if ($result["ok"]) {
    echo $result["status"];  // 200
    print_r($result["data"]);  // response body
} else {
    // On an HTTP error status there is no err (only a transport failure sets
    // it), so fall back to the status code.
    $err = $result["err"] ?? null;
    echo "Error: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```

### Prepare a request without sending it

```php
// prepare() throws on error and returns the fetch definition.
$fetchdef = $client->prepare([
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => ["id" => "example"],
]);

echo $fetchdef["url"];
echo $fetchdef["method"];
print_r($fetchdef["headers"]);
```

### Use test mode

Create a mock client for unit testing — no server required:

```php
$client = LumaSDK::test();

// Entity ops return the bare mock record (throws on error).
$eventcoupon = $client->EventCoupon()->list();
print_r($eventcoupon);
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```php
$mock_fetch = function ($url, $init) {
    return [
        [
            "status" => 200,
            "statusText" => "OK",
            "headers" => [],
            "json" => function () { return ["id" => "mock01"]; },
        ],
        null,
    ];
};

$client = new LumaSDK([
    "base" => "http://localhost:8080",
    "system" => [
        "fetch" => $mock_fetch,
    ],
]);
```

### Run live tests

Create a `.env.local` file at the project root:

```
LUMA_TEST_LIVE=TRUE
LUMA_APIKEY=<your-key>
```

Then run:

```bash
cd php && ./vendor/bin/phpunit test/
```


## Reference

### LumaSDK

```php
require_once 'luma_sdk.php';
$client = new LumaSDK($options);
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `array` | Feature activation flags. |
| `extend` | `array` | Additional Feature instances to load. |
| `system` | `array` | System overrides (e.g. custom `fetch` callable). |

### test

```php
$client = LumaSDK::test($testopts, $sdkopts);
```

Creates a test-mode client with mock transport. Both arguments may be `null`.

### LumaSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `(): array` | Deep copy of current SDK options. |
| `get_utility` | `(): Utility` | Copy of the SDK utility object. |
| `prepare` | `(array $fetchargs): array` | Build an HTTP request definition without sending. |
| `direct` | `(array $fetchargs): array` | Build and send an HTTP request. |
| `Calendar` | `($data): CalendarEntity` | Create a Calendar entity instance. |
| `CalendarAdmin` | `($data): CalendarAdminEntity` | Create a CalendarAdmin entity instance. |
| `CalendarCoupon` | `($data): CalendarCouponEntity` | Create a CalendarCoupon entity instance. |
| `CalendarEvent` | `($data): CalendarEventEntity` | Create a CalendarEvent entity instance. |
| `CalendarEventApproval` | `($data): CalendarEventApprovalEntity` | Create a CalendarEventApproval entity instance. |
| `CalendarEventRejection` | `($data): CalendarEventRejectionEntity` | Create a CalendarEventRejection entity instance. |
| `Contact` | `($data): ContactEntity` | Create a Contact entity instance. |
| `ContactBlock` | `($data): ContactBlockEntity` | Create a ContactBlock entity instance. |
| `ContactRestore` | `($data): ContactRestoreEntity` | Create a ContactRestore entity instance. |
| `ContactTag` | `($data): ContactTagEntity` | Create a ContactTag entity instance. |
| `ContactTagAssignment` | `($data): ContactTagAssignmentEntity` | Create a ContactTagAssignment entity instance. |
| `EntityLookup` | `($data): EntityLookupEntity` | Create an EntityLookup entity instance. |
| `Event` | `($data): EventEntity` | Create an Event entity instance. |
| `EventCancelRequest` | `($data): EventCancelRequestEntity` | Create an EventCancelRequest entity instance. |
| `EventCoupon` | `($data): EventCouponEntity` | Create an EventCoupon entity instance. |
| `EventTag` | `($data): EventTagEntity` | Create an EventTag entity instance. |
| `EventTagAssignment` | `($data): EventTagAssignmentEntity` | Create an EventTagAssignment entity instance. |
| `Guest` | `($data): GuestEntity` | Create a Guest entity instance. |
| `GuestInvite` | `($data): GuestInviteEntity` | Create a GuestInvite entity instance. |
| `GuestTicket` | `($data): GuestTicketEntity` | Create a GuestTicket entity instance. |
| `Host` | `($data): HostEntity` | Create a Host entity instance. |
| `ImageUpload` | `($data): ImageUploadEntity` | Create an ImageUpload entity instance. |
| `Member` | `($data): MemberEntity` | Create a Member entity instance. |
| `MembershipTier` | `($data): MembershipTierEntity` | Create a MembershipTier entity instance. |
| `OrganizationAdmin` | `($data): OrganizationAdminEntity` | Create an OrganizationAdmin entity instance. |
| `OrganizationCalendar` | `($data): OrganizationCalendarEntity` | Create an OrganizationCalendar entity instance. |
| `OrganizationEvent` | `($data): OrganizationEventEntity` | Create an OrganizationEvent entity instance. |
| `OrganizationEventTransfer` | `($data): OrganizationEventTransferEntity` | Create an OrganizationEventTransfer entity instance. |
| `TicketType` | `($data): TicketTypeEntity` | Create a TicketType entity instance. |
| `User` | `($data): UserEntity` | Create an User entity instance. |
| `Webhook` | `($data): WebhookEntity` | Create a Webhook entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `($reqmatch, $ctrl): array` | Load a single entity by match criteria. |
| `list` | `(?array $reqmatch = null, $ctrl): array` | List entities matching the criteria (call with no argument to list all). |
| `create` | `($reqdata, $ctrl): array` | Create a new entity. |
| `update` | `($reqdata, $ctrl): array` | Update an existing entity. |
| `remove` | `($reqmatch, $ctrl): array` | Remove an entity. |
| `data_get` | `(): array` | Get entity data. |
| `data_set` | `($data): void` | Set entity data. |
| `match_get` | `(): array` | Get entity match criteria. |
| `match_set` | `($match): void` | Set entity match criteria. |
| `make` | `(): Entity` | Create a new instance with the same options. |
| `get_name` | `(): string` | Return the entity name. |

### Result shape

Entity operations return the bare result data (an `array` for single-entity
ops, a `list` for `list`) and throw on error. Wrap calls in
`try`/`catch` to handle failures.

The `direct()` escape hatch never throws — it returns a result `array`
you branch on via `$result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `true` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `array` | Response headers. |
| `data` | `mixed` | Parsed JSON response body. |

On error, `ok` is `false` and `$err` contains the error value.

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

Create an instance: `$calendar = $client->Calendar();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avatar_url` | `mixed` |  |
| `calendar_id` | `string` |  |
| `coordinate` | `mixed` |  |
| `cover_image_url` | `mixed` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `instagram_handle` | `mixed` |  |
| `is_personal` | `bool` |  |
| `location` | `mixed` |  |
| `name` | `string` |  |
| `slug` | `string` |  |
| `social_image_url` | `mixed` |  |
| `tint_color` | `string` |  |
| `twitter_handle` | `mixed` |  |
| `url` | `string` |  |
| `website` | `mixed` |  |
| `youtube_handle` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare Calendar record (throws on error).
$calendar = $client->Calendar()->load(["id" => "calendar_id"]);
```


### CalendarAdmin

Create an instance: `$calendar_admin = $client->CalendarAdmin();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avatar_url` | `string` |  |
| `email` | `string` |  |
| `first_name` | `mixed` |  |
| `id` | `string` |  |
| `last_name` | `mixed` |  |
| `name` | `string` |  |

#### Example: List

```php
// list() returns an array of CalendarAdmin records (throws on error).
$calendar_admins = $client->CalendarAdmin()->list();
```


### CalendarCoupon

Create an instance: `$calendar_coupon = $client->CalendarCoupon();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `cents_off` | `mixed` |  |
| `code` | `string` |  |
| `currency` | `mixed` |  |
| `discount` | `mixed` |  |
| `event_ticket_type_id` | `string` |  |
| `id` | `string` |  |
| `percent_off` | `mixed` |  |
| `remaining_count` | `int` |  |
| `valid_end_at` | `mixed` |  |
| `valid_start_at` | `mixed` |  |

#### Example: List

```php
// list() returns an array of CalendarCoupon records (throws on error).
$calendar_coupons = $client->CalendarCoupon()->list();
```

#### Example: Create

```php
$calendar_coupon = $client->CalendarCoupon()->create([
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


### CalendarEvent

Create an instance: `$calendar_event = $client->CalendarEvent();`

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
| `submitted_by` | `mixed` |  |
| `tag` | `array` |  |

#### Example: Load

```php
// load() returns the bare CalendarEvent record (throws on error).
$calendar_event = $client->CalendarEvent()->load(["id" => "calendar_event_id"]);
```

#### Example: List

```php
// list() returns an array of CalendarEvent records (throws on error).
$calendar_events = $client->CalendarEvent()->list();
```

#### Example: Create

```php
$calendar_event = $client->CalendarEvent()->create([
    "id" => null, // string
    "status" => null, // string
    "submitted_by" => null, // mixed
    "tag" => null, // array
]);
```


### CalendarEventApproval

Create an instance: `$calendar_event_approval = $client->CalendarEventApproval();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `calendar_event_id` | `string` |  |

#### Example: Create

```php
$calendar_event_approval = $client->CalendarEventApproval()->create([
    "calendar_event_id" => null, // string
]);
```


### CalendarEventRejection

Create an instance: `$calendar_event_rejection = $client->CalendarEventRejection();`

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

```php
$calendar_event_rejection = $client->CalendarEventRejection()->create([
    "calendar_event_id" => null, // string
]);
```


### Contact

Create an instance: `$contact = $client->Contact();`

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
| `contact` | `array` |  |
| `created_at` | `string` |  |
| `email` | `string` |  |
| `event_approved_count` | `float` |  |
| `event_checked_in_count` | `float` |  |
| `first_name` | `mixed` |  |
| `id` | `string` |  |
| `last_name` | `mixed` |  |
| `membership` | `mixed` |  |
| `name` | `string` |  |
| `revenue_usd_cent` | `float` |  |
| `tag` | `mixed` |  |
| `user_id` | `string` |  |

#### Example: List

```php
// list() returns an array of Contact records (throws on error).
$contacts = $client->Contact()->list();
```

#### Example: Create

```php
$contact = $client->Contact()->create([
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


### ContactBlock

Create an instance: `$contact_block = $client->ContactBlock();`

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

```php
$contact_block = $client->ContactBlock()->create([
]);
```


### ContactRestore

Create an instance: `$contact_restore = $client->ContactRestore();`

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

```php
$contact_restore = $client->ContactRestore()->create([
]);
```


### ContactTag

Create an instance: `$contact_tag = $client->ContactTag();`

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
| `color` | `mixed` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `tag_id` | `string` |  |

#### Example: List

```php
// list() returns an array of ContactTag records (throws on error).
$contact_tags = $client->ContactTag()->list();
```

#### Example: Create

```php
$contact_tag = $client->ContactTag()->create([
    "id" => null, // string
    "name" => null, // string
    "tag_id" => null, // string
]);
```


### ContactTagAssignment

Create an instance: `$contact_tag_assignment = $client->ContactTagAssignment();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `applied_count` | `float` |  |
| `email` | `array` |  |
| `skipped_count` | `float` |  |
| `tag` | `string` |  |
| `user_id` | `array` |  |

#### Example: Create

```php
$contact_tag_assignment = $client->ContactTagAssignment()->create([
    "applied_count" => null, // float
    "skipped_count" => null, // float
    "tag" => null, // string
]);
```


### EntityLookup

Create an instance: `$entity_lookup = $client->EntityLookup();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```php
// load() returns the bare EntityLookup record (throws on error).
$entity_lookup = $client->EntityLookup()->load();
```


### Event

Create an instance: `$event = $client->Event();`

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
| `can_register_for_multiple_ticket` | `bool` |  |
| `coordinate` | `mixed` |  |
| `cover_url` | `string` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `description_md` | `string` |  |
| `display_price` | `mixed` |  |
| `duration_interval` | `string` |  |
| `end_at` | `string` |  |
| `event_id` | `string` |  |
| `feedback_email` | `array` |  |
| `geo_address_json` | `mixed` |  |
| `guest_count` | `array` |  |
| `host` | `array` |  |
| `id` | `string` |  |
| `location_type` | `string` |  |
| `location_visibility` | `string` |  |
| `max_capacity` | `mixed` |  |
| `meeting_url` | `mixed` |  |
| `name` | `string` |  |
| `name_requirement` | `string` |  |
| `phone_number_requirement` | `mixed` |  |
| `platform` | `string` |  |
| `registration_open` | `bool` |  |
| `registration_question` | `array` |  |
| `reminders_disabled` | `bool` |  |
| `require_approval` | `bool` |  |
| `show_guest_list` | `bool` |  |
| `slug` | `string` |  |
| `spots_remaining` | `mixed` |  |
| `start_at` | `string` |  |
| `suppress_notification` | `bool` |  |
| `timezone` | `string` |  |
| `tint_color` | `string` |  |
| `url` | `string` |  |
| `user_id` | `string` |  |
| `visibility` | `string` |  |
| `waitlist_status` | `string` |  |

#### Example: Load

```php
// load() returns the bare Event record (throws on error).
$event = $client->Event()->load(["id" => "event_id"]);
```

#### Example: Create

```php
$event = $client->Event()->create([
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


### EventCancelRequest

Create an instance: `$event_cancel_request = $client->EventCancelRequest();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `cancellation_token` | `string` |  |
| `event_id` | `string` |  |
| `guest_count` | `float` |  |
| `is_paid` | `bool` |  |

#### Example: Create

```php
$event_cancel_request = $client->EventCancelRequest()->create([
    "cancellation_token" => null, // string
    "event_id" => null, // string
    "guest_count" => null, // float
    "is_paid" => null, // bool
]);
```


### EventCoupon

Create an instance: `$event_coupon = $client->EventCoupon();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `cents_off` | `mixed` |  |
| `code` | `string` |  |
| `currency` | `mixed` |  |
| `discount` | `mixed` |  |
| `event_id` | `string` |  |
| `event_ticket_type_id` | `string` |  |
| `id` | `string` |  |
| `percent_off` | `mixed` |  |
| `remaining_count` | `int` |  |
| `valid_end_at` | `mixed` |  |
| `valid_start_at` | `mixed` |  |

#### Example: List

```php
// list() returns an array of EventCoupon records (throws on error).
$event_coupons = $client->EventCoupon()->list();
```

#### Example: Create

```php
$event_coupon = $client->EventCoupon()->create([
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


### EventTag

Create an instance: `$event_tag = $client->EventTag();`

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
| `color` | `mixed` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `tag_id` | `string` |  |

#### Example: List

```php
// list() returns an array of EventTag records (throws on error).
$event_tags = $client->EventTag()->list();
```

#### Example: Create

```php
$event_tag = $client->EventTag()->create([
    "id" => null, // string
    "name" => null, // string
    "tag_id" => null, // string
]);
```


### EventTagAssignment

Create an instance: `$event_tag_assignment = $client->EventTagAssignment();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `applied_count` | `float` |  |
| `event_id` | `array` |  |
| `skipped_count` | `float` |  |
| `tag` | `string` |  |

#### Example: Create

```php
$event_tag_assignment = $client->EventTagAssignment()->create([
    "applied_count" => null, // float
    "event_id" => null, // array
    "skipped_count" => null, // float
    "tag" => null, // string
]);
```


### Guest

Create an instance: `$guest = $client->Guest();`

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
| `eth_address` | `mixed` |  |
| `event_id` | `string` |  |
| `event_ticket` | `array` |  |
| `event_ticket_order` | `array` |  |
| `guest` | `array` |  |
| `guest_id` | `string` |  |
| `id` | `string` |  |
| `invited_at` | `mixed` |  |
| `joined_at` | `mixed` |  |
| `message` | `mixed` |  |
| `phone_number` | `int` |  |
| `registered_at` | `mixed` |  |
| `registration_answer` | `mixed` |  |
| `send_email` | `mixed` |  |
| `should_refund` | `bool` |  |
| `solana_address` | `mixed` |  |
| `status` | `string` |  |
| `ticket` | `mixed` |  |
| `user_email` | `string` |  |
| `user_first_name` | `mixed` |  |
| `user_id` | `string` |  |
| `user_last_name` | `mixed` |  |
| `user_name` | `mixed` |  |
| `utm_source` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare Guest record (throws on error).
$guest = $client->Guest()->load(["id" => "guest_id"]);
```

#### Example: List

```php
// list() returns an array of Guest records (throws on error).
$guests = $client->Guest()->list();
```

#### Example: Create

```php
$guest = $client->Guest()->create([
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


### GuestInvite

Create an instance: `$guest_invite = $client->GuestInvite();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `event_id` | `string` |  |
| `guest` | `array` |  |
| `message` | `mixed` |  |

#### Example: Create

```php
$guest_invite = $client->GuestInvite()->create([
    "event_id" => null, // string
    "guest" => null, // array
]);
```


### GuestTicket

Create an instance: `$guest_ticket = $client->GuestTicket();`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `event_id` | `string` |  |
| `guest_id` | `string` |  |
| `send_email` | `mixed` |  |
| `ticket_ids_to_remove` | `array` |  |
| `tickets_to_add` | `array` |  |


### Host

Create an instance: `$host = $client->Host();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `access_level` | `mixed` |  |
| `email` | `string` |  |
| `event_id` | `string` |  |
| `is_visible` | `bool` |  |
| `name` | `string` |  |

#### Example: Create

```php
$host = $client->Host()->create([
    "email" => null, // string
    "event_id" => null, // string
]);
```


### ImageUpload

Create an instance: `$image_upload = $client->ImageUpload();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `content_type` | `mixed` |  |
| `file_url` | `string` |  |
| `upload_url` | `string` |  |

#### Example: Create

```php
$image_upload = $client->ImageUpload()->create([
    "file_url" => null, // string
    "upload_url" => null, // string
]);
```


### Member

Create an instance: `$member = $client->Member();`

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
| `registration_answer` | `array` |  |
| `skip_payment` | `bool` |  |
| `status` | `string` |  |
| `user_id` | `string` |  |

#### Example: Create

```php
$member = $client->Member()->create([
    "email" => null, // string
    "membership_id" => null, // string
    "membership_tier_id" => null, // string
    "status" => null, // string
    "user_id" => null, // string
]);
```


### MembershipTier

Create an instance: `$membership_tier = $client->MembershipTier();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `access_info` | `mixed` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `tint_color` | `string` |  |

#### Example: List

```php
// list() returns an array of MembershipTier records (throws on error).
$membership_tiers = $client->MembershipTier()->list();
```


### OrganizationAdmin

Create an instance: `$organization_admin = $client->OrganizationAdmin();`

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
| `first_name` | `mixed` |  |
| `id` | `string` |  |
| `last_name` | `mixed` |  |
| `name` | `string` |  |

#### Example: List

```php
// list() returns an array of OrganizationAdmin records (throws on error).
$organization_admins = $client->OrganizationAdmin()->list();
```


### OrganizationCalendar

Create an instance: `$organization_calendar = $client->OrganizationCalendar();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avatar_url` | `mixed` |  |
| `coordinate` | `mixed` |  |
| `cover_image_url` | `mixed` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `instagram_handle` | `mixed` |  |
| `is_personal` | `bool` |  |
| `location` | `mixed` |  |
| `name` | `string` |  |
| `slug` | `string` |  |
| `social_image_url` | `mixed` |  |
| `tint_color` | `string` |  |
| `twitter_handle` | `mixed` |  |
| `url` | `string` |  |
| `website` | `mixed` |  |
| `youtube_handle` | `mixed` |  |

#### Example: List

```php
// list() returns an array of OrganizationCalendar records (throws on error).
$organization_calendars = $client->OrganizationCalendar()->list();
```

#### Example: Create

```php
$organization_calendar = $client->OrganizationCalendar()->create([
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


### OrganizationEvent

Create an instance: `$organization_event = $client->OrganizationEvent();`

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
| `coordinate` | `mixed` |  |
| `cover_url` | `string` |  |
| `created_at` | `string` |  |
| `display_price` | `mixed` |  |
| `duration_interval` | `string` |  |
| `end_at` | `string` |  |
| `feedback_email` | `array` |  |
| `geo_address_json` | `mixed` |  |
| `geo_latitude` | `mixed` |  |
| `geo_longitude` | `mixed` |  |
| `id` | `string` |  |
| `location_type` | `string` |  |
| `location_visibility` | `string` |  |
| `managing_calendar` | `array` |  |
| `meeting_url` | `mixed` |  |
| `name` | `string` |  |
| `platform` | `string` |  |
| `registration_open` | `bool` |  |
| `registration_question` | `array` |  |
| `require_approval` | `bool` |  |
| `spots_remaining` | `mixed` |  |
| `start_at` | `string` |  |
| `timezone` | `string` |  |
| `url` | `string` |  |
| `user_api_id` | `string` |  |
| `user_id` | `string` |  |
| `visibility` | `string` |  |
| `waitlist_status` | `string` |  |
| `zoom_meeting_url` | `mixed` |  |

#### Example: List

```php
// list() returns an array of OrganizationEvent records (throws on error).
$organization_events = $client->OrganizationEvent()->list();
```


### OrganizationEventTransfer

Create an instance: `$organization_event_transfer = $client->OrganizationEventTransfer();`

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

```php
$organization_event_transfer = $client->OrganizationEventTransfer()->create([
    "calendar_id" => null, // string
    "event_id" => null, // string
]);
```


### TicketType

Create an instance: `$ticket_type = $client->TicketType();`

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
| `cent` | `mixed` |  |
| `currency` | `mixed` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `is_flexible` | `bool` |  |
| `is_hidden` | `bool` |  |
| `max_capacity` | `mixed` |  |
| `min_cent` | `mixed` |  |
| `name` | `string` |  |
| `require_approval` | `bool` |  |
| `type` | `string` |  |
| `valid_end_at` | `mixed` |  |
| `valid_start_at` | `mixed` |  |

#### Example: Load

```php
// load() returns the bare TicketType record (throws on error).
$ticket_type = $client->TicketType()->load(["id" => "ticket_type_id"]);
```

#### Example: List

```php
// list() returns an array of TicketType records (throws on error).
$ticket_types = $client->TicketType()->list();
```

#### Example: Create

```php
$ticket_type = $client->TicketType()->create([
    "id" => null, // string
    "name" => null, // string
    "type" => null, // string
]);
```


### User

Create an instance: `$user = $client->User();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avatar_url` | `string` |  |
| `email` | `string` |  |
| `first_name` | `mixed` |  |
| `id` | `string` |  |
| `last_name` | `mixed` |  |
| `name` | `string` |  |

#### Example: Load

```php
// load() returns the bare User record (throws on error).
$user = $client->User()->load(["id" => "user_id"]);
```


### Webhook

Create an instance: `$webhook = $client->Webhook();`

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
| `event_type` | `array` |  |
| `id` | `string` |  |
| `secret` | `string` |  |
| `status` | `string` |  |
| `url` | `string` |  |

#### Example: Load

```php
// load() returns the bare Webhook record (throws on error).
$webhook = $client->Webhook()->load(["id" => "webhook_id"]);
```

#### Example: List

```php
// list() returns an array of Webhook records (throws on error).
$webhooks = $client->Webhook()->list();
```

#### Example: Create

```php
$webhook = $client->Webhook()->create([
    "created_at" => null, // string
    "event_type" => null, // array
    "id" => null, // string
    "secret" => null, // string
    "status" => null, // string
    "url" => null, // string
]);
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

Features are the extension mechanism. A feature is a PHP class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as arrays

The PHP SDK uses plain PHP associative arrays throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers::to_map()` to safely validate that a value is an array.

### Directory structure

```
php/
├── luma_sdk.php          -- Main SDK class
├── config.php                     -- Configuration
├── features.php                   -- Feature factory
├── core/                          -- Core types and context
├── entity/                        -- Entity implementations
├── feature/                       -- Built-in features (Base, Test, Log)
├── utility/                       -- Utility functions and struct library
└── test/                          -- Test suites
```

The main class (`luma_sdk.php`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```php
$eventcoupon = $client->EventCoupon();
$eventcoupon->list();

// $eventcoupon->data_get() now returns the eventcoupon data from the last list
// $eventcoupon->match_get() returns the last match criteria
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
