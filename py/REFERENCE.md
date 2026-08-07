# Luma Python SDK Reference

Complete API reference for the Luma Python SDK.


## LumaSDK

### Constructor

```python
from luma_sdk import LumaSDK

client = LumaSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["apikey"]` | `str` | API key for authentication. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `LumaSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = LumaSDK.test()
```


### Instance Methods

#### `Calendar(data=None)`

Create a new `CalendarEntity` instance. Pass `None` for no initial data.

#### `CalendarAdmin(data=None)`

Create a new `CalendarAdminEntity` instance. Pass `None` for no initial data.

#### `CalendarCoupon(data=None)`

Create a new `CalendarCouponEntity` instance. Pass `None` for no initial data.

#### `CalendarEvent(data=None)`

Create a new `CalendarEventEntity` instance. Pass `None` for no initial data.

#### `CalendarEventApproval(data=None)`

Create a new `CalendarEventApprovalEntity` instance. Pass `None` for no initial data.

#### `CalendarEventRejection(data=None)`

Create a new `CalendarEventRejectionEntity` instance. Pass `None` for no initial data.

#### `Contact(data=None)`

Create a new `ContactEntity` instance. Pass `None` for no initial data.

#### `ContactBlock(data=None)`

Create a new `ContactBlockEntity` instance. Pass `None` for no initial data.

#### `ContactRestore(data=None)`

Create a new `ContactRestoreEntity` instance. Pass `None` for no initial data.

#### `ContactTag(data=None)`

Create a new `ContactTagEntity` instance. Pass `None` for no initial data.

#### `ContactTagAssignment(data=None)`

Create a new `ContactTagAssignmentEntity` instance. Pass `None` for no initial data.

#### `EntityLookup(data=None)`

Create a new `EntityLookupEntity` instance. Pass `None` for no initial data.

#### `Event(data=None)`

Create a new `EventEntity` instance. Pass `None` for no initial data.

#### `EventCancelRequest(data=None)`

Create a new `EventCancelRequestEntity` instance. Pass `None` for no initial data.

#### `EventCoupon(data=None)`

Create a new `EventCouponEntity` instance. Pass `None` for no initial data.

#### `EventTag(data=None)`

Create a new `EventTagEntity` instance. Pass `None` for no initial data.

#### `EventTagAssignment(data=None)`

Create a new `EventTagAssignmentEntity` instance. Pass `None` for no initial data.

#### `Guest(data=None)`

Create a new `GuestEntity` instance. Pass `None` for no initial data.

#### `GuestInvite(data=None)`

Create a new `GuestInviteEntity` instance. Pass `None` for no initial data.

#### `GuestTicket(data=None)`

Create a new `GuestTicketEntity` instance. Pass `None` for no initial data.

#### `Host(data=None)`

Create a new `HostEntity` instance. Pass `None` for no initial data.

#### `ImageUpload(data=None)`

Create a new `ImageUploadEntity` instance. Pass `None` for no initial data.

#### `Member(data=None)`

Create a new `MemberEntity` instance. Pass `None` for no initial data.

#### `MembershipTier(data=None)`

Create a new `MembershipTierEntity` instance. Pass `None` for no initial data.

#### `OrganizationAdmin(data=None)`

Create a new `OrganizationAdminEntity` instance. Pass `None` for no initial data.

#### `OrganizationCalendar(data=None)`

Create a new `OrganizationCalendarEntity` instance. Pass `None` for no initial data.

#### `OrganizationEvent(data=None)`

Create a new `OrganizationEventEntity` instance. Pass `None` for no initial data.

#### `OrganizationEventTransfer(data=None)`

Create a new `OrganizationEventTransferEntity` instance. Pass `None` for no initial data.

#### `TicketType(data=None)`

Create a new `TicketTypeEntity` instance. Pass `None` for no initial data.

#### `User(data=None)`

Create a new `UserEntity` instance. Pass `None` for no initial data.

#### `Webhook(data=None)`

Create a new `WebhookEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## CalendarEntity

```python
calendar = client.Calendar()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avatar_url` | `Any` | Yes |  |
| `calendar_id` | `str` | Yes |  |
| `coordinate` | `Any` | Yes |  |
| `cover_image_url` | `Any` | Yes |  |
| `description` | `str` | Yes |  |
| `id` | `str` | Yes |  |
| `instagram_handle` | `Any` | Yes |  |
| `is_personal` | `bool` | Yes |  |
| `location` | `Any` | Yes |  |
| `name` | `str` | Yes |  |
| `slug` | `str` | Yes |  |
| `social_image_url` | `Any` | Yes |  |
| `tint_color` | `str` | No |  |
| `twitter_handle` | `Any` | Yes |  |
| `url` | `str` | Yes |  |
| `website` | `Any` | Yes |  |
| `youtube_handle` | `Any` | Yes |  |

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

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Calendar().load({"id": "calendar_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Calendar().update({
    "id": "calendar_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CalendarEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CalendarAdminEntity

```python
calendar_admin = client.CalendarAdmin()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avatar_url` | `str` | Yes |  |
| `email` | `str` | Yes |  |
| `first_name` | `Any` | Yes |  |
| `id` | `str` | Yes |  |
| `last_name` | `Any` | Yes |  |
| `name` | `str` | Yes |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.CalendarAdmin().list()
for calendar_admin in results:
    print(calendar_admin)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CalendarAdminEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CalendarCouponEntity

```python
calendar_coupon = client.CalendarCoupon()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `cents_off` | `Any` | Yes |  |
| `code` | `str` | Yes |  |
| `currency` | `Any` | Yes |  |
| `discount` | `Any` | Yes |  |
| `event_ticket_type_id` | `str` | No |  |
| `id` | `str` | Yes |  |
| `percent_off` | `Any` | Yes |  |
| `remaining_count` | `int` | Yes |  |
| `valid_end_at` | `Any` | Yes |  |
| `valid_start_at` | `Any` | Yes |  |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.CalendarCoupon().create({
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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.CalendarCoupon().list()
for calendar_coupon in results:
    print(calendar_coupon)
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.CalendarCoupon().update({
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CalendarCouponEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CalendarEventEntity

```python
calendar_event = client.CalendarEvent()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `str` | Yes |  |
| `status` | `str` | Yes |  |
| `submitted_by` | `Any` | Yes |  |
| `tag` | `list` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.CalendarEvent().create({
    "id": "example_id",  # str
    "status": "example_status",  # str
    "submitted_by": "example_submitted_by",  # Any
    "tag": [],  # list
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.CalendarEvent().list()
for calendar_event in results:
    print(calendar_event)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.CalendarEvent().load({"id": "calendar_event_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CalendarEventEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CalendarEventApprovalEntity

```python
calendar_event_approval = client.CalendarEventApproval()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `calendar_event_id` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.CalendarEventApproval().create({
    "calendar_event_id": "example_calendar_event_id",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CalendarEventApprovalEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CalendarEventRejectionEntity

```python
calendar_event_rejection = client.CalendarEventRejection()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `calendar_event_id` | `str` | Yes |  |
| `message` | `str` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.CalendarEventRejection().create({
    "calendar_event_id": "example_calendar_event_id",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CalendarEventRejectionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ContactEntity

```python
contact = client.Contact()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avatar_url` | `str` | Yes |  |
| `contact` | `list` | Yes |  |
| `created_at` | `str` | Yes |  |
| `email` | `str` | Yes |  |
| `event_approved_count` | `float` | Yes |  |
| `event_checked_in_count` | `float` | Yes |  |
| `first_name` | `Any` | Yes |  |
| `id` | `str` | Yes |  |
| `last_name` | `Any` | Yes |  |
| `membership` | `Any` | Yes |  |
| `name` | `str` | Yes |  |
| `revenue_usd_cent` | `float` | Yes |  |
| `tag` | `Any` | No |  |
| `user_id` | `str` | Yes |  |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Contact().create({
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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Contact().list()
for contact in results:
    print(contact)
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Contact().remove({"id": "id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ContactEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ContactBlockEntity

```python
contact_block = client.ContactBlock()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `contact_id` | `str` | No |  |
| `email` | `str` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.ContactBlock().create({
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ContactBlockEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ContactRestoreEntity

```python
contact_restore = client.ContactRestore()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `contact_id` | `str` | No |  |
| `email` | `str` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.ContactRestore().create({
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ContactRestoreEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ContactTagEntity

```python
contact_tag = client.ContactTag()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `color` | `Any` | No |  |
| `id` | `str` | Yes |  |
| `name` | `str` | Yes |  |
| `tag_id` | `str` | Yes |  |

### Field Usage by Operation

| Field | list | create | update | remove |
| --- | --- | --- | --- | --- |
| `color` | Yes | - | - | - |
| `id` | - | - | - | - |
| `name` | - | - | Yes | - |
| `tag_id` | - | - | - | - |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.ContactTag().create({
    "id": "example_id",  # str
    "name": "example_name",  # str
    "tag_id": "example_tag_id",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.ContactTag().list()
for contact_tag in results:
    print(contact_tag)
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.ContactTag().remove({"id": "id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.ContactTag().update({
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ContactTagEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ContactTagAssignmentEntity

```python
contact_tag_assignment = client.ContactTagAssignment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `applied_count` | `float` | Yes |  |
| `email` | `list` | No |  |
| `skipped_count` | `float` | Yes |  |
| `tag` | `str` | Yes |  |
| `user_id` | `list` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.ContactTagAssignment().create({
    "applied_count": 1,  # float
    "skipped_count": 1,  # float
    "tag": "example_tag",  # str
})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.ContactTagAssignment().remove()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ContactTagAssignmentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EntityLookupEntity

```python
entity_lookup = client.EntityLookup()
```

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.EntityLookup().load()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EntityLookupEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EventEntity

```python
event = client.Event()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `access` | `str` | Yes |  |
| `calendar_id` | `str` | Yes |  |
| `can_register_for_multiple_ticket` | `bool` | No |  |
| `coordinate` | `Any` | Yes |  |
| `cover_url` | `str` | Yes |  |
| `created_at` | `str` | Yes |  |
| `description` | `str` | Yes |  |
| `description_md` | `str` | Yes |  |
| `display_price` | `Any` | Yes |  |
| `duration_interval` | `str` | Yes |  |
| `end_at` | `str` | Yes |  |
| `event_id` | `str` | Yes |  |
| `feedback_email` | `dict` | Yes |  |
| `geo_address_json` | `Any` | Yes |  |
| `guest_count` | `dict` | Yes |  |
| `host` | `list` | Yes |  |
| `id` | `str` | Yes |  |
| `location_type` | `str` | Yes |  |
| `location_visibility` | `str` | Yes |  |
| `max_capacity` | `Any` | No |  |
| `meeting_url` | `Any` | Yes |  |
| `name` | `str` | Yes |  |
| `name_requirement` | `str` | No |  |
| `phone_number_requirement` | `Any` | No |  |
| `platform` | `str` | Yes |  |
| `registration_open` | `bool` | Yes |  |
| `registration_question` | `list` | No |  |
| `reminders_disabled` | `bool` | No |  |
| `require_approval` | `bool` | Yes |  |
| `show_guest_list` | `bool` | No |  |
| `slug` | `str` | No |  |
| `spots_remaining` | `Any` | Yes |  |
| `start_at` | `str` | Yes |  |
| `suppress_notification` | `bool` | No |  |
| `timezone` | `str` | Yes |  |
| `tint_color` | `str` | No |  |
| `url` | `str` | Yes |  |
| `user_id` | `str` | Yes |  |
| `visibility` | `str` | Yes |  |
| `waitlist_status` | `str` | Yes |  |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Event().create({
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

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Event().load({"id": "event_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Event().remove({"id": "event_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Event().update({
    "id": "event_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EventEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EventCancelRequestEntity

```python
event_cancel_request = client.EventCancelRequest()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `cancellation_token` | `str` | Yes |  |
| `event_id` | `str` | Yes |  |
| `guest_count` | `float` | Yes |  |
| `is_paid` | `bool` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.EventCancelRequest().create({
    "cancellation_token": "example_cancellation_token",  # str
    "event_id": "example_event_id",  # str
    "guest_count": 1,  # float
    "is_paid": True,  # bool
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EventCancelRequestEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EventCouponEntity

```python
event_coupon = client.EventCoupon()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `cents_off` | `Any` | Yes |  |
| `code` | `str` | Yes |  |
| `currency` | `Any` | Yes |  |
| `discount` | `Any` | Yes |  |
| `event_id` | `str` | Yes |  |
| `event_ticket_type_id` | `str` | No |  |
| `id` | `str` | Yes |  |
| `percent_off` | `Any` | Yes |  |
| `remaining_count` | `int` | Yes |  |
| `valid_end_at` | `Any` | Yes |  |
| `valid_start_at` | `Any` | Yes |  |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.EventCoupon().create({
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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.EventCoupon().list()
for event_coupon in results:
    print(event_coupon)
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.EventCoupon().update({
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EventCouponEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EventTagEntity

```python
event_tag = client.EventTag()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `color` | `Any` | No |  |
| `id` | `str` | Yes |  |
| `name` | `str` | Yes |  |
| `tag_id` | `str` | Yes |  |

### Field Usage by Operation

| Field | list | create | update | remove |
| --- | --- | --- | --- | --- |
| `color` | Yes | - | - | - |
| `id` | - | - | - | - |
| `name` | - | - | Yes | - |
| `tag_id` | - | - | - | - |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.EventTag().create({
    "id": "example_id",  # str
    "name": "example_name",  # str
    "tag_id": "example_tag_id",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.EventTag().list()
for event_tag in results:
    print(event_tag)
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.EventTag().remove({"id": "id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.EventTag().update({
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EventTagEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EventTagAssignmentEntity

```python
event_tag_assignment = client.EventTagAssignment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `applied_count` | `float` | Yes |  |
| `event_id` | `list` | Yes |  |
| `skipped_count` | `float` | Yes |  |
| `tag` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.EventTagAssignment().create({
    "applied_count": 1,  # float
    "event_id": [],  # list
    "skipped_count": 1,  # float
    "tag": "example_tag",  # str
})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.EventTagAssignment().remove()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EventTagAssignmentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## GuestEntity

```python
guest = client.Guest()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_status` | `str` | Yes |  |
| `check_in_qr_code` | `str` | Yes |  |
| `eth_address` | `Any` | Yes |  |
| `event_id` | `str` | Yes |  |
| `event_ticket` | `list` | Yes |  |
| `event_ticket_order` | `list` | Yes |  |
| `guest` | `list` | Yes |  |
| `guest_id` | `str` | Yes |  |
| `id` | `str` | Yes |  |
| `invited_at` | `Any` | Yes |  |
| `joined_at` | `Any` | Yes |  |
| `message` | `Any` | No |  |
| `phone_number` | `int` | Yes |  |
| `registered_at` | `Any` | Yes |  |
| `registration_answer` | `Any` | Yes |  |
| `send_email` | `Any` | No |  |
| `should_refund` | `bool` | No |  |
| `solana_address` | `Any` | Yes |  |
| `status` | `str` | Yes |  |
| `ticket` | `Any` | No |  |
| `user_email` | `str` | Yes |  |
| `user_first_name` | `Any` | Yes |  |
| `user_id` | `str` | Yes |  |
| `user_last_name` | `Any` | Yes |  |
| `user_name` | `Any` | Yes |  |
| `utm_source` | `Any` | Yes |  |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Guest().create({
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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Guest().list()
for guest in results:
    print(guest)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Guest().load({"id": "guest_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Guest().update({
    "id": "guest_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `GuestEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## GuestInviteEntity

```python
guest_invite = client.GuestInvite()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `event_id` | `str` | Yes |  |
| `guest` | `list` | Yes |  |
| `message` | `Any` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.GuestInvite().create({
    "event_id": "example_event_id",  # str
    "guest": [],  # list
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `GuestInviteEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## GuestTicketEntity

```python
guest_ticket = client.GuestTicket()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `event_id` | `str` | Yes |  |
| `guest_id` | `str` | Yes |  |
| `send_email` | `Any` | No |  |
| `ticket_ids_to_remove` | `list` | No |  |
| `tickets_to_add` | `list` | No |  |

### Operations

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.GuestTicket().update({
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `GuestTicketEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## HostEntity

```python
host = client.Host()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `access_level` | `Any` | No |  |
| `email` | `str` | Yes |  |
| `event_id` | `str` | Yes |  |
| `is_visible` | `bool` | No |  |
| `name` | `str` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Host().create({
    "email": "example_email",  # str
    "event_id": "example_event_id",  # str
})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Host().remove()
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Host().update({
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `HostEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ImageUploadEntity

```python
image_upload = client.ImageUpload()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `content_type` | `Any` | No |  |
| `file_url` | `str` | Yes |  |
| `upload_url` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.ImageUpload().create({
    "file_url": "example_file_url",  # str
    "upload_url": "example_upload_url",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ImageUploadEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## MemberEntity

```python
member = client.Member()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `email` | `str` | Yes |  |
| `membership_id` | `str` | Yes |  |
| `membership_tier_id` | `str` | Yes |  |
| `registration_answer` | `list` | No |  |
| `skip_payment` | `bool` | No |  |
| `status` | `str` | Yes |  |
| `user_id` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Member().create({
    "email": "example_email",  # str
    "membership_id": "example_membership_id",  # str
    "membership_tier_id": "example_membership_tier_id",  # str
    "status": "example_status",  # str
    "user_id": "example_user_id",  # str
})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Member().update({
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MemberEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## MembershipTierEntity

```python
membership_tier = client.MembershipTier()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `access_info` | `Any` | Yes |  |
| `description` | `str` | Yes |  |
| `id` | `str` | Yes |  |
| `name` | `str` | Yes |  |
| `tint_color` | `str` | Yes |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.MembershipTier().list()
for membership_tier in results:
    print(membership_tier)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MembershipTierEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## OrganizationAdminEntity

```python
organization_admin = client.OrganizationAdmin()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_id` | `str` | Yes |  |
| `avatar_url` | `str` | Yes |  |
| `email` | `str` | Yes |  |
| `first_name` | `Any` | Yes |  |
| `id` | `str` | Yes |  |
| `last_name` | `Any` | Yes |  |
| `name` | `str` | Yes |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.OrganizationAdmin().list()
for organization_admin in results:
    print(organization_admin)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `OrganizationAdminEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## OrganizationCalendarEntity

```python
organization_calendar = client.OrganizationCalendar()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avatar_url` | `Any` | Yes |  |
| `coordinate` | `Any` | Yes |  |
| `cover_image_url` | `Any` | Yes |  |
| `description` | `str` | Yes |  |
| `id` | `str` | Yes |  |
| `instagram_handle` | `Any` | Yes |  |
| `is_personal` | `bool` | Yes |  |
| `location` | `Any` | Yes |  |
| `name` | `str` | Yes |  |
| `slug` | `str` | Yes |  |
| `social_image_url` | `Any` | Yes |  |
| `tint_color` | `str` | No |  |
| `twitter_handle` | `Any` | Yes |  |
| `url` | `str` | Yes |  |
| `website` | `Any` | Yes |  |
| `youtube_handle` | `Any` | Yes |  |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.OrganizationCalendar().create({
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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.OrganizationCalendar().list()
for organization_calendar in results:
    print(organization_calendar)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `OrganizationCalendarEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## OrganizationEventEntity

```python
organization_event = client.OrganizationEvent()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_id` | `str` | Yes |  |
| `calendar_api_id` | `str` | Yes |  |
| `calendar_id` | `str` | Yes |  |
| `coordinate` | `Any` | Yes |  |
| `cover_url` | `str` | Yes |  |
| `created_at` | `str` | Yes |  |
| `display_price` | `Any` | Yes |  |
| `duration_interval` | `str` | Yes |  |
| `end_at` | `str` | Yes |  |
| `feedback_email` | `dict` | Yes |  |
| `geo_address_json` | `Any` | Yes |  |
| `geo_latitude` | `Any` | Yes |  |
| `geo_longitude` | `Any` | Yes |  |
| `id` | `str` | Yes |  |
| `location_type` | `str` | Yes |  |
| `location_visibility` | `str` | Yes |  |
| `managing_calendar` | `list` | Yes |  |
| `meeting_url` | `Any` | Yes |  |
| `name` | `str` | Yes |  |
| `platform` | `str` | Yes |  |
| `registration_open` | `bool` | Yes |  |
| `registration_question` | `list` | No |  |
| `require_approval` | `bool` | Yes |  |
| `spots_remaining` | `Any` | Yes |  |
| `start_at` | `str` | Yes |  |
| `timezone` | `str` | Yes |  |
| `url` | `str` | Yes |  |
| `user_api_id` | `str` | Yes |  |
| `user_id` | `str` | Yes |  |
| `visibility` | `str` | Yes |  |
| `waitlist_status` | `str` | Yes |  |
| `zoom_meeting_url` | `Any` | Yes |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.OrganizationEvent().list()
for organization_event in results:
    print(organization_event)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `OrganizationEventEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## OrganizationEventTransferEntity

```python
organization_event_transfer = client.OrganizationEventTransfer()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `calendar_id` | `str` | Yes |  |
| `event_id` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.OrganizationEventTransfer().create({
    "calendar_id": "example_calendar_id",  # str
    "event_id": "example_event_id",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `OrganizationEventTransferEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## TicketTypeEntity

```python
ticket_type = client.TicketType()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `cent` | `Any` | No |  |
| `currency` | `Any` | No |  |
| `description` | `str` | No |  |
| `id` | `str` | Yes |  |
| `is_flexible` | `bool` | No |  |
| `is_hidden` | `bool` | No |  |
| `max_capacity` | `Any` | No |  |
| `min_cent` | `Any` | No |  |
| `name` | `str` | Yes |  |
| `require_approval` | `bool` | No |  |
| `type` | `str` | Yes |  |
| `valid_end_at` | `Any` | No |  |
| `valid_start_at` | `Any` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.TicketType().create({
    "id": "example_id",  # str
    "name": "example_name",  # str
    "type": "example_type",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.TicketType().list()
for ticket_type in results:
    print(ticket_type)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.TicketType().load({"id": "ticket_type_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.TicketType().remove({"id": "ticket_type_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.TicketType().update({
    "id": "ticket_type_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TicketTypeEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## UserEntity

```python
user = client.User()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avatar_url` | `str` | Yes |  |
| `email` | `str` | Yes |  |
| `first_name` | `Any` | Yes |  |
| `id` | `str` | Yes |  |
| `last_name` | `Any` | Yes |  |
| `name` | `str` | Yes |  |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.User().load({"id": "user_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UserEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## WebhookEntity

```python
webhook = client.Webhook()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `str` | Yes |  |
| `event_type` | `list` | Yes |  |
| `id` | `str` | Yes |  |
| `secret` | `str` | Yes |  |
| `status` | `str` | Yes |  |
| `url` | `str` | Yes |  |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Webhook().create({
    "created_at": "example_created_at",  # str
    "event_type": [],  # list
    "id": "example_id",  # str
    "secret": "example_secret",  # str
    "status": "example_status",  # str
    "url": "example_url",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Webhook().list()
for webhook in results:
    print(webhook)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Webhook().load({"id": "webhook_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Webhook().remove({"id": "webhook_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Webhook().update({
    "id": "webhook_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `WebhookEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```python
client = LumaSDK({
    "feature": {
        "test": {"active": True},
    },
})
```

