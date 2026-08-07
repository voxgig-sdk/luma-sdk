# Luma Python SDK



The Python SDK for the Luma API — an entity-oriented client following Pythonic conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.Calendar()` — each
carrying a small, uniform set of operations (`list`, `load`, `create`, `update`, `remove`) instead of raw URL
paths and query strings. You work with named resources and verbs, which
keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to PyPI. Install it from the GitHub
release tag (`py/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/luma-sdk/releases)) or
from a source checkout:

```bash
pip install -e .
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```python
import os
from luma_sdk import LumaSDK

client = LumaSDK({
    "apikey": os.environ.get("LUMA_APIKEY"),
})
```

### 3. Load a calendar

`load()` returns the bare record (a `dict`) and raises on error.

```python
try:
    calendar = client.Calendar().load({"id": "example_id"})
    print(calendar)
except Exception as err:
    print(f"load failed: {err}")
```

### 4. Create, update, and remove

```python
# Update
client.Calendar().update({"id": "example_id", "avatar_url": "example_avatar_url", "calendar_id": "example_calendar_id"})

```


## Error handling

Entity operations raise on failure, so wrap them in `try` / `except`:

```python
try:
    eventcoupons = client.EventCoupon().list()
    print(eventcoupons)
except Exception as err:
    print(f"list failed: {err}")
```

`direct()` does **not** raise — it returns the result envelope. Branch
on `ok`; on failure `status` holds the HTTP status (for error responses)
and `err` holds a transport error, so read both defensively:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example_id"},
})

if not result["ok"]:
    print("request failed:", result.get("status"), result.get("err"))
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})

if result["ok"]:
    print(result["status"])  # 200
    print(result["data"])    # response body
else:
    # A non-2xx response carries status + data (the error body); a
    # transport-level failure carries err instead. Only one is present, so
    # read both with .get() rather than indexing a key that may be absent.
    print(result.get("status"), result.get("err"))
```

### Prepare a request without sending it

```python
# prepare() returns the fetch definition and raises on error.
fetchdef = client.prepare({
    "path": "/api/resource/{id}",
    "method": "DELETE",
    "params": {"id": "example"},
})

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```python
client = LumaSDK.test()

# Entity ops return the bare record and raise on error.
eventcoupon = client.EventCoupon().list()
# eventcoupon contains the mock response record
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```python
def mock_fetch(url, init):
    return {
        "status": 200,
        "statusText": "OK",
        "headers": {},
        "json": lambda: {"id": "mock01"},
    }, None

client = LumaSDK({
    "base": "http://localhost:8080",
    "system": {
        "fetch": mock_fetch,
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
cd py && pytest test/
```


## Reference

### LumaSDK

```python
from luma_sdk import LumaSDK

client = LumaSDK(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `str` | API key for authentication. |
| `base` | `str` | Base URL of the API server. |
| `prefix` | `str` | URL path prefix prepended to all requests. |
| `suffix` | `str` | URL path suffix appended to all requests. |
| `feature` | `dict` | Feature activation flags. |
| `extend` | `list` | Additional Feature instances to load. |
| `system` | `dict` | System overrides (e.g. custom `fetch` function). |

### test

```python
client = LumaSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `None`.

### LumaSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> dict` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> dict` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> dict` | Build and send an HTTP request. Returns a result dict (branch on `ok`). |
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
| `load` | `(reqmatch, ctrl) -> any` | Load a single entity by match criteria. Raises on error. |
| `list` | `(reqmatch, ctrl) -> list` | List entities matching the criteria. Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity. Raises on error. |
| `update` | `(reqdata, ctrl) -> any` | Update an existing entity. Raises on error. |
| `remove` | `(reqmatch, ctrl) -> any` | Remove an entity. Raises on error. |
| `data_get` | `() -> dict` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> dict` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> str` | Return the entity name. |

### Result shape

Entity operations return the bare result data (a `dict` for single-entity
ops, a `list` for `list`) and raise on error. Wrap calls in
`try`/`except` to handle failures.

The `direct()` escape hatch never raises — it returns a result `dict`
you branch on via `result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `True` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `dict` | Response headers. |
| `data` | `any` | Parsed JSON response body. |

On error, `ok` is `False` and `err` contains the error value.

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

Create an instance: `calendar = client.Calendar()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avatar_url` | `Any` |  |
| `calendar_id` | `str` |  |
| `coordinate` | `Any` |  |
| `cover_image_url` | `Any` |  |
| `description` | `str` |  |
| `id` | `str` |  |
| `instagram_handle` | `Any` |  |
| `is_personal` | `bool` |  |
| `location` | `Any` |  |
| `name` | `str` |  |
| `slug` | `str` |  |
| `social_image_url` | `Any` |  |
| `tint_color` | `str` |  |
| `twitter_handle` | `Any` |  |
| `url` | `str` |  |
| `website` | `Any` |  |
| `youtube_handle` | `Any` |  |

#### Example: Load

```python
calendar = client.Calendar().load({"id": "calendar_id"})
```


### CalendarAdmin

Create an instance: `calendar_admin = client.CalendarAdmin()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avatar_url` | `str` |  |
| `email` | `str` |  |
| `first_name` | `Any` |  |
| `id` | `str` |  |
| `last_name` | `Any` |  |
| `name` | `str` |  |

#### Example: List

```python
calendar_admins = client.CalendarAdmin().list()
```


### CalendarCoupon

Create an instance: `calendar_coupon = client.CalendarCoupon()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `cents_off` | `Any` |  |
| `code` | `str` |  |
| `currency` | `Any` |  |
| `discount` | `Any` |  |
| `event_ticket_type_id` | `str` |  |
| `id` | `str` |  |
| `percent_off` | `Any` |  |
| `remaining_count` | `int` |  |
| `valid_end_at` | `Any` |  |
| `valid_start_at` | `Any` |  |

#### Example: List

```python
calendar_coupons = client.CalendarCoupon().list()
```

#### Example: Create

```python
calendar_coupon = client.CalendarCoupon().create({
    "cents_off": "example_cents_off",  # Any
    "code": "example_code",  # str
    "currency": "example_currency",  # Any
    "discount": "example_discount",  # Any
    "id": "example_id",  # str
    "percent_off": "example_percent_off",  # Any
    "remaining_count": 1,  # int
    "valid_end_at": "example_valid_end_at",  # Any
    "valid_start_at": "example_valid_start_at",  # Any
})
```


### CalendarEvent

Create an instance: `calendar_event = client.CalendarEvent()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `str` |  |
| `status` | `str` |  |
| `submitted_by` | `Any` |  |
| `tag` | `list` |  |

#### Example: Load

```python
calendar_event = client.CalendarEvent().load({"id": "calendar_event_id"})
```

#### Example: List

```python
calendar_events = client.CalendarEvent().list()
```

#### Example: Create

```python
calendar_event = client.CalendarEvent().create({
    "id": "example_id",  # str
    "status": "example_status",  # str
    "submitted_by": "example_submitted_by",  # Any
    "tag": [],  # list
})
```


### CalendarEventApproval

Create an instance: `calendar_event_approval = client.CalendarEventApproval()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `calendar_event_id` | `str` |  |

#### Example: Create

```python
calendar_event_approval = client.CalendarEventApproval().create({
    "calendar_event_id": "example_calendar_event_id",  # str
})
```


### CalendarEventRejection

Create an instance: `calendar_event_rejection = client.CalendarEventRejection()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `calendar_event_id` | `str` |  |
| `message` | `str` |  |

#### Example: Create

```python
calendar_event_rejection = client.CalendarEventRejection().create({
    "calendar_event_id": "example_calendar_event_id",  # str
})
```


### Contact

Create an instance: `contact = client.Contact()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avatar_url` | `str` |  |
| `contact` | `list` |  |
| `created_at` | `str` |  |
| `email` | `str` |  |
| `event_approved_count` | `float` |  |
| `event_checked_in_count` | `float` |  |
| `first_name` | `Any` |  |
| `id` | `str` |  |
| `last_name` | `Any` |  |
| `membership` | `Any` |  |
| `name` | `str` |  |
| `revenue_usd_cent` | `float` |  |
| `tag` | `Any` |  |
| `user_id` | `str` |  |

#### Example: List

```python
contacts = client.Contact().list()
```

#### Example: Create

```python
contact = client.Contact().create({
    "avatar_url": "example_avatar_url",  # str
    "contact": [],  # list
    "created_at": "example_created_at",  # str
    "email": "example_email",  # str
    "event_approved_count": 1,  # float
    "event_checked_in_count": 1,  # float
    "first_name": "example_first_name",  # Any
    "id": "example_id",  # str
    "last_name": "example_last_name",  # Any
    "membership": "example_membership",  # Any
    "name": "example_name",  # str
    "revenue_usd_cent": 1,  # float
    "user_id": "example_user_id",  # str
})
```


### ContactBlock

Create an instance: `contact_block = client.ContactBlock()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `contact_id` | `str` |  |
| `email` | `str` |  |

#### Example: Create

```python
contact_block = client.ContactBlock().create({
})
```


### ContactRestore

Create an instance: `contact_restore = client.ContactRestore()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `contact_id` | `str` |  |
| `email` | `str` |  |

#### Example: Create

```python
contact_restore = client.ContactRestore().create({
})
```


### ContactTag

Create an instance: `contact_tag = client.ContactTag()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `color` | `Any` |  |
| `id` | `str` |  |
| `name` | `str` |  |
| `tag_id` | `str` |  |

#### Example: List

```python
contact_tags = client.ContactTag().list()
```

#### Example: Create

```python
contact_tag = client.ContactTag().create({
    "id": "example_id",  # str
    "name": "example_name",  # str
    "tag_id": "example_tag_id",  # str
})
```


### ContactTagAssignment

Create an instance: `contact_tag_assignment = client.ContactTagAssignment()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `applied_count` | `float` |  |
| `email` | `list` |  |
| `skipped_count` | `float` |  |
| `tag` | `str` |  |
| `user_id` | `list` |  |

#### Example: Create

```python
contact_tag_assignment = client.ContactTagAssignment().create({
    "applied_count": 1,  # float
    "skipped_count": 1,  # float
    "tag": "example_tag",  # str
})
```


### EntityLookup

Create an instance: `entity_lookup = client.EntityLookup()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```python
entity_lookup = client.EntityLookup().load()
```


### Event

Create an instance: `event = client.Event()`

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
| `access` | `str` |  |
| `calendar_id` | `str` |  |
| `can_register_for_multiple_ticket` | `bool` |  |
| `coordinate` | `Any` |  |
| `cover_url` | `str` |  |
| `created_at` | `str` |  |
| `description` | `str` |  |
| `description_md` | `str` |  |
| `display_price` | `Any` |  |
| `duration_interval` | `str` |  |
| `end_at` | `str` |  |
| `event_id` | `str` |  |
| `feedback_email` | `dict` |  |
| `geo_address_json` | `Any` |  |
| `guest_count` | `dict` |  |
| `host` | `list` |  |
| `id` | `str` |  |
| `location_type` | `str` |  |
| `location_visibility` | `str` |  |
| `max_capacity` | `Any` |  |
| `meeting_url` | `Any` |  |
| `name` | `str` |  |
| `name_requirement` | `str` |  |
| `phone_number_requirement` | `Any` |  |
| `platform` | `str` |  |
| `registration_open` | `bool` |  |
| `registration_question` | `list` |  |
| `reminders_disabled` | `bool` |  |
| `require_approval` | `bool` |  |
| `show_guest_list` | `bool` |  |
| `slug` | `str` |  |
| `spots_remaining` | `Any` |  |
| `start_at` | `str` |  |
| `suppress_notification` | `bool` |  |
| `timezone` | `str` |  |
| `tint_color` | `str` |  |
| `url` | `str` |  |
| `user_id` | `str` |  |
| `visibility` | `str` |  |
| `waitlist_status` | `str` |  |

#### Example: Load

```python
event = client.Event().load({"id": "event_id"})
```

#### Example: Create

```python
event = client.Event().create({
    "access": "example_access",  # str
    "calendar_id": "example_calendar_id",  # str
    "coordinate": "example_coordinate",  # Any
    "cover_url": "example_cover_url",  # str
    "created_at": "example_created_at",  # str
    "description": "example_description",  # str
    "description_md": "example_description_md",  # str
    "display_price": "example_display_price",  # Any
    "duration_interval": "example_duration_interval",  # str
    "end_at": "example_end_at",  # str
    "event_id": "example_event_id",  # str
    "feedback_email": {},  # dict
    "geo_address_json": "example_geo_address_json",  # Any
    "guest_count": {},  # dict
    "host": [],  # list
    "id": "example_id",  # str
    "location_type": "example_location_type",  # str
    "location_visibility": "example_location_visibility",  # str
    "meeting_url": "example_meeting_url",  # Any
    "name": "example_name",  # str
    "platform": "example_platform",  # str
    "registration_open": True,  # bool
    "require_approval": True,  # bool
    "spots_remaining": "example_spots_remaining",  # Any
    "start_at": "example_start_at",  # str
    "timezone": "example_timezone",  # str
    "url": "example_url",  # str
    "user_id": "example_user_id",  # str
    "visibility": "example_visibility",  # str
    "waitlist_status": "example_waitlist_status",  # str
})
```


### EventCancelRequest

Create an instance: `event_cancel_request = client.EventCancelRequest()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `cancellation_token` | `str` |  |
| `event_id` | `str` |  |
| `guest_count` | `float` |  |
| `is_paid` | `bool` |  |

#### Example: Create

```python
event_cancel_request = client.EventCancelRequest().create({
    "cancellation_token": "example_cancellation_token",  # str
    "event_id": "example_event_id",  # str
    "guest_count": 1,  # float
    "is_paid": True,  # bool
})
```


### EventCoupon

Create an instance: `event_coupon = client.EventCoupon()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `cents_off` | `Any` |  |
| `code` | `str` |  |
| `currency` | `Any` |  |
| `discount` | `Any` |  |
| `event_id` | `str` |  |
| `event_ticket_type_id` | `str` |  |
| `id` | `str` |  |
| `percent_off` | `Any` |  |
| `remaining_count` | `int` |  |
| `valid_end_at` | `Any` |  |
| `valid_start_at` | `Any` |  |

#### Example: List

```python
event_coupons = client.EventCoupon().list()
```

#### Example: Create

```python
event_coupon = client.EventCoupon().create({
    "cents_off": "example_cents_off",  # Any
    "code": "example_code",  # str
    "currency": "example_currency",  # Any
    "discount": "example_discount",  # Any
    "event_id": "example_event_id",  # str
    "id": "example_id",  # str
    "percent_off": "example_percent_off",  # Any
    "remaining_count": 1,  # int
    "valid_end_at": "example_valid_end_at",  # Any
    "valid_start_at": "example_valid_start_at",  # Any
})
```


### EventTag

Create an instance: `event_tag = client.EventTag()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `color` | `Any` |  |
| `id` | `str` |  |
| `name` | `str` |  |
| `tag_id` | `str` |  |

#### Example: List

```python
event_tags = client.EventTag().list()
```

#### Example: Create

```python
event_tag = client.EventTag().create({
    "id": "example_id",  # str
    "name": "example_name",  # str
    "tag_id": "example_tag_id",  # str
})
```


### EventTagAssignment

Create an instance: `event_tag_assignment = client.EventTagAssignment()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `applied_count` | `float` |  |
| `event_id` | `list` |  |
| `skipped_count` | `float` |  |
| `tag` | `str` |  |

#### Example: Create

```python
event_tag_assignment = client.EventTagAssignment().create({
    "applied_count": 1,  # float
    "event_id": [],  # list
    "skipped_count": 1,  # float
    "tag": "example_tag",  # str
})
```


### Guest

Create an instance: `guest = client.Guest()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approval_status` | `str` |  |
| `check_in_qr_code` | `str` |  |
| `eth_address` | `Any` |  |
| `event_id` | `str` |  |
| `event_ticket` | `list` |  |
| `event_ticket_order` | `list` |  |
| `guest` | `list` |  |
| `guest_id` | `str` |  |
| `id` | `str` |  |
| `invited_at` | `Any` |  |
| `joined_at` | `Any` |  |
| `message` | `Any` |  |
| `phone_number` | `int` |  |
| `registered_at` | `Any` |  |
| `registration_answer` | `Any` |  |
| `send_email` | `Any` |  |
| `should_refund` | `bool` |  |
| `solana_address` | `Any` |  |
| `status` | `str` |  |
| `ticket` | `Any` |  |
| `user_email` | `str` |  |
| `user_first_name` | `Any` |  |
| `user_id` | `str` |  |
| `user_last_name` | `Any` |  |
| `user_name` | `Any` |  |
| `utm_source` | `Any` |  |

#### Example: Load

```python
guest = client.Guest().load({"id": "guest_id"})
```

#### Example: List

```python
guests = client.Guest().list()
```

#### Example: Create

```python
guest = client.Guest().create({
    "approval_status": "example_approval_status",  # str
    "check_in_qr_code": "example_check_in_qr_code",  # str
    "eth_address": "example_eth_address",  # Any
    "event_id": "example_event_id",  # str
    "event_ticket": [],  # list
    "event_ticket_order": [],  # list
    "guest": [],  # list
    "guest_id": "example_guest_id",  # str
    "id": "example_id",  # str
    "invited_at": "example_invited_at",  # Any
    "joined_at": "example_joined_at",  # Any
    "phone_number": 1,  # int
    "registered_at": "example_registered_at",  # Any
    "registration_answer": "example_registration_answer",  # Any
    "solana_address": "example_solana_address",  # Any
    "status": "example_status",  # str
    "user_email": "example_user_email",  # str
    "user_first_name": "example_user_first_name",  # Any
    "user_id": "example_user_id",  # str
    "user_last_name": "example_user_last_name",  # Any
    "user_name": "example_user_name",  # Any
    "utm_source": "example_utm_source",  # Any
})
```


### GuestInvite

Create an instance: `guest_invite = client.GuestInvite()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `event_id` | `str` |  |
| `guest` | `list` |  |
| `message` | `Any` |  |

#### Example: Create

```python
guest_invite = client.GuestInvite().create({
    "event_id": "example_event_id",  # str
    "guest": [],  # list
})
```


### GuestTicket

Create an instance: `guest_ticket = client.GuestTicket()`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `event_id` | `str` |  |
| `guest_id` | `str` |  |
| `send_email` | `Any` |  |
| `ticket_ids_to_remove` | `list` |  |
| `tickets_to_add` | `list` |  |


### Host

Create an instance: `host = client.Host()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `access_level` | `Any` |  |
| `email` | `str` |  |
| `event_id` | `str` |  |
| `is_visible` | `bool` |  |
| `name` | `str` |  |

#### Example: Create

```python
host = client.Host().create({
    "email": "example_email",  # str
    "event_id": "example_event_id",  # str
})
```


### ImageUpload

Create an instance: `image_upload = client.ImageUpload()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `content_type` | `Any` |  |
| `file_url` | `str` |  |
| `upload_url` | `str` |  |

#### Example: Create

```python
image_upload = client.ImageUpload().create({
    "file_url": "example_file_url",  # str
    "upload_url": "example_upload_url",  # str
})
```


### Member

Create an instance: `member = client.Member()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `email` | `str` |  |
| `membership_id` | `str` |  |
| `membership_tier_id` | `str` |  |
| `registration_answer` | `list` |  |
| `skip_payment` | `bool` |  |
| `status` | `str` |  |
| `user_id` | `str` |  |

#### Example: Create

```python
member = client.Member().create({
    "email": "example_email",  # str
    "membership_id": "example_membership_id",  # str
    "membership_tier_id": "example_membership_tier_id",  # str
    "status": "example_status",  # str
    "user_id": "example_user_id",  # str
})
```


### MembershipTier

Create an instance: `membership_tier = client.MembershipTier()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `access_info` | `Any` |  |
| `description` | `str` |  |
| `id` | `str` |  |
| `name` | `str` |  |
| `tint_color` | `str` |  |

#### Example: List

```python
membership_tiers = client.MembershipTier().list()
```


### OrganizationAdmin

Create an instance: `organization_admin = client.OrganizationAdmin()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_id` | `str` |  |
| `avatar_url` | `str` |  |
| `email` | `str` |  |
| `first_name` | `Any` |  |
| `id` | `str` |  |
| `last_name` | `Any` |  |
| `name` | `str` |  |

#### Example: List

```python
organization_admins = client.OrganizationAdmin().list()
```


### OrganizationCalendar

Create an instance: `organization_calendar = client.OrganizationCalendar()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avatar_url` | `Any` |  |
| `coordinate` | `Any` |  |
| `cover_image_url` | `Any` |  |
| `description` | `str` |  |
| `id` | `str` |  |
| `instagram_handle` | `Any` |  |
| `is_personal` | `bool` |  |
| `location` | `Any` |  |
| `name` | `str` |  |
| `slug` | `str` |  |
| `social_image_url` | `Any` |  |
| `tint_color` | `str` |  |
| `twitter_handle` | `Any` |  |
| `url` | `str` |  |
| `website` | `Any` |  |
| `youtube_handle` | `Any` |  |

#### Example: List

```python
organization_calendars = client.OrganizationCalendar().list()
```

#### Example: Create

```python
organization_calendar = client.OrganizationCalendar().create({
    "avatar_url": "example_avatar_url",  # Any
    "coordinate": "example_coordinate",  # Any
    "cover_image_url": "example_cover_image_url",  # Any
    "description": "example_description",  # str
    "id": "example_id",  # str
    "instagram_handle": "example_instagram_handle",  # Any
    "is_personal": True,  # bool
    "location": "example_location",  # Any
    "name": "example_name",  # str
    "slug": "example_slug",  # str
    "social_image_url": "example_social_image_url",  # Any
    "twitter_handle": "example_twitter_handle",  # Any
    "url": "example_url",  # str
    "website": "example_website",  # Any
    "youtube_handle": "example_youtube_handle",  # Any
})
```


### OrganizationEvent

Create an instance: `organization_event = client.OrganizationEvent()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `api_id` | `str` |  |
| `calendar_api_id` | `str` |  |
| `calendar_id` | `str` |  |
| `coordinate` | `Any` |  |
| `cover_url` | `str` |  |
| `created_at` | `str` |  |
| `display_price` | `Any` |  |
| `duration_interval` | `str` |  |
| `end_at` | `str` |  |
| `feedback_email` | `dict` |  |
| `geo_address_json` | `Any` |  |
| `geo_latitude` | `Any` |  |
| `geo_longitude` | `Any` |  |
| `id` | `str` |  |
| `location_type` | `str` |  |
| `location_visibility` | `str` |  |
| `managing_calendar` | `list` |  |
| `meeting_url` | `Any` |  |
| `name` | `str` |  |
| `platform` | `str` |  |
| `registration_open` | `bool` |  |
| `registration_question` | `list` |  |
| `require_approval` | `bool` |  |
| `spots_remaining` | `Any` |  |
| `start_at` | `str` |  |
| `timezone` | `str` |  |
| `url` | `str` |  |
| `user_api_id` | `str` |  |
| `user_id` | `str` |  |
| `visibility` | `str` |  |
| `waitlist_status` | `str` |  |
| `zoom_meeting_url` | `Any` |  |

#### Example: List

```python
organization_events = client.OrganizationEvent().list()
```


### OrganizationEventTransfer

Create an instance: `organization_event_transfer = client.OrganizationEventTransfer()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `calendar_id` | `str` |  |
| `event_id` | `str` |  |

#### Example: Create

```python
organization_event_transfer = client.OrganizationEventTransfer().create({
    "calendar_id": "example_calendar_id",  # str
    "event_id": "example_event_id",  # str
})
```


### TicketType

Create an instance: `ticket_type = client.TicketType()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `cent` | `Any` |  |
| `currency` | `Any` |  |
| `description` | `str` |  |
| `id` | `str` |  |
| `is_flexible` | `bool` |  |
| `is_hidden` | `bool` |  |
| `max_capacity` | `Any` |  |
| `min_cent` | `Any` |  |
| `name` | `str` |  |
| `require_approval` | `bool` |  |
| `type` | `str` |  |
| `valid_end_at` | `Any` |  |
| `valid_start_at` | `Any` |  |

#### Example: Load

```python
ticket_type = client.TicketType().load({"id": "ticket_type_id"})
```

#### Example: List

```python
ticket_types = client.TicketType().list()
```

#### Example: Create

```python
ticket_type = client.TicketType().create({
    "id": "example_id",  # str
    "name": "example_name",  # str
    "type": "example_type",  # str
})
```


### User

Create an instance: `user = client.User()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avatar_url` | `str` |  |
| `email` | `str` |  |
| `first_name` | `Any` |  |
| `id` | `str` |  |
| `last_name` | `Any` |  |
| `name` | `str` |  |

#### Example: Load

```python
user = client.User().load({"id": "user_id"})
```


### Webhook

Create an instance: `webhook = client.Webhook()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `str` |  |
| `event_type` | `list` |  |
| `id` | `str` |  |
| `secret` | `str` |  |
| `status` | `str` |  |
| `url` | `str` |  |

#### Example: Load

```python
webhook = client.Webhook().load({"id": "webhook_id"})
```

#### Example: List

```python
webhooks = client.Webhook().list()
```

#### Example: Create

```python
webhook = client.Webhook().create({
    "created_at": "example_created_at",  # str
    "event_type": [],  # list
    "id": "example_id",  # str
    "secret": "example_secret",  # str
    "status": "example_status",  # str
    "url": "example_url",  # str
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

Features are the extension mechanism. A feature is a Python class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as dicts

The Python SDK uses plain dicts throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a dict.

### Module structure

```
py/
├── luma_sdk.py         -- Main SDK module
├── config.py                    -- Configuration
├── features.py                  -- Feature factory
├── core/                        -- Core types and context
├── entity/                      -- Entity implementations
├── feature/                     -- Built-in features (Base, Test, Log)
├── utility/                     -- Utility functions and struct library
└── test/                        -- Test suites
```

The main module (`luma_sdk`) exports the SDK class.
Import entity or utility modules directly only when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```python
eventcoupon = client.EventCoupon()
eventcoupon.list()

# eventcoupon.data_get() now returns the eventcoupon data from the last list
# eventcoupon.match_get() returns the last match criteria
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
