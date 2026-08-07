# Luma Golang SDK Reference

Complete API reference for the Luma Golang SDK.


## LumaSDK

### Constructor

```go
func NewLumaSDK(options map[string]any) *LumaSDK
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `map[string]any` | SDK configuration options. |
| `options["apikey"]` | `string` | API key for authentication. |
| `options["base"]` | `string` | Base URL for API requests. |
| `options["prefix"]` | `string` | URL prefix appended after base. |
| `options["suffix"]` | `string` | URL suffix appended after path. |
| `options["headers"]` | `map[string]any` | Custom headers for all requests. |
| `options["feature"]` | `map[string]any` | Feature configuration. |
| `options["system"]` | `map[string]any` | System overrides (e.g. custom fetch). |


### Static Methods

#### `Test() *LumaSDK`

No-arg convenience constructor for the common no-options test case.

```go
client := sdk.Test()
```

#### `TestSDK(testopts, sdkopts map[string]any) *LumaSDK`

Test client with options. Both arguments may be `nil`.

```go
client := sdk.TestSDK(testopts, sdkopts)
```


### Instance Methods

#### `Calendar(data map[string]any) LumaEntity`

Create a new `Calendar` entity instance. Pass `nil` for no initial data.

#### `CalendarAdmin(data map[string]any) LumaEntity`

Create a new `CalendarAdmin` entity instance. Pass `nil` for no initial data.

#### `CalendarCoupon(data map[string]any) LumaEntity`

Create a new `CalendarCoupon` entity instance. Pass `nil` for no initial data.

#### `CalendarEvent(data map[string]any) LumaEntity`

Create a new `CalendarEvent` entity instance. Pass `nil` for no initial data.

#### `CalendarEventApproval(data map[string]any) LumaEntity`

Create a new `CalendarEventApproval` entity instance. Pass `nil` for no initial data.

#### `CalendarEventRejection(data map[string]any) LumaEntity`

Create a new `CalendarEventRejection` entity instance. Pass `nil` for no initial data.

#### `Contact(data map[string]any) LumaEntity`

Create a new `Contact` entity instance. Pass `nil` for no initial data.

#### `ContactBlock(data map[string]any) LumaEntity`

Create a new `ContactBlock` entity instance. Pass `nil` for no initial data.

#### `ContactRestore(data map[string]any) LumaEntity`

Create a new `ContactRestore` entity instance. Pass `nil` for no initial data.

#### `ContactTag(data map[string]any) LumaEntity`

Create a new `ContactTag` entity instance. Pass `nil` for no initial data.

#### `ContactTagAssignment(data map[string]any) LumaEntity`

Create a new `ContactTagAssignment` entity instance. Pass `nil` for no initial data.

#### `EntityLookup(data map[string]any) LumaEntity`

Create a new `EntityLookup` entity instance. Pass `nil` for no initial data.

#### `Event(data map[string]any) LumaEntity`

Create a new `Event` entity instance. Pass `nil` for no initial data.

#### `EventCancelRequest(data map[string]any) LumaEntity`

Create a new `EventCancelRequest` entity instance. Pass `nil` for no initial data.

#### `EventCoupon(data map[string]any) LumaEntity`

Create a new `EventCoupon` entity instance. Pass `nil` for no initial data.

#### `EventTag(data map[string]any) LumaEntity`

Create a new `EventTag` entity instance. Pass `nil` for no initial data.

#### `EventTagAssignment(data map[string]any) LumaEntity`

Create a new `EventTagAssignment` entity instance. Pass `nil` for no initial data.

#### `Guest(data map[string]any) LumaEntity`

Create a new `Guest` entity instance. Pass `nil` for no initial data.

#### `GuestInvite(data map[string]any) LumaEntity`

Create a new `GuestInvite` entity instance. Pass `nil` for no initial data.

#### `GuestTicket(data map[string]any) LumaEntity`

Create a new `GuestTicket` entity instance. Pass `nil` for no initial data.

#### `Host(data map[string]any) LumaEntity`

Create a new `Host` entity instance. Pass `nil` for no initial data.

#### `ImageUpload(data map[string]any) LumaEntity`

Create a new `ImageUpload` entity instance. Pass `nil` for no initial data.

#### `Member(data map[string]any) LumaEntity`

Create a new `Member` entity instance. Pass `nil` for no initial data.

#### `MembershipTier(data map[string]any) LumaEntity`

Create a new `MembershipTier` entity instance. Pass `nil` for no initial data.

#### `OrganizationAdmin(data map[string]any) LumaEntity`

Create a new `OrganizationAdmin` entity instance. Pass `nil` for no initial data.

#### `OrganizationCalendar(data map[string]any) LumaEntity`

Create a new `OrganizationCalendar` entity instance. Pass `nil` for no initial data.

#### `OrganizationEvent(data map[string]any) LumaEntity`

Create a new `OrganizationEvent` entity instance. Pass `nil` for no initial data.

#### `OrganizationEventTransfer(data map[string]any) LumaEntity`

Create a new `OrganizationEventTransfer` entity instance. Pass `nil` for no initial data.

#### `TicketType(data map[string]any) LumaEntity`

Create a new `TicketType` entity instance. Pass `nil` for no initial data.

#### `User(data map[string]any) LumaEntity`

Create a new `User` entity instance. Pass `nil` for no initial data.

#### `Webhook(data map[string]any) LumaEntity`

Create a new `Webhook` entity instance. Pass `nil` for no initial data.

#### `OptionsMap() map[string]any`

Return a deep copy of the current SDK options.

#### `GetUtility() *Utility`

Return a copy of the SDK utility object.

#### `Direct(fetchargs map[string]any) (map[string]any, error)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `map[string]any` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `map[string]any` | Query string parameters. |
| `fetchargs["headers"]` | `map[string]any` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (maps are JSON-serialized). |
| `fetchargs["ctrl"]` | `map[string]any` | Control options (e.g. `map[string]any{"explain": true}`). |

**Returns:** `(map[string]any, error)`

#### `Prepare(fetchargs map[string]any) (map[string]any, error)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `Direct()`.

**Returns:** `(map[string]any, error)`


---

## CalendarEntity

```go
calendar := client.Calendar(nil)
fmt.Println(calendar.GetName()) // "calendar"
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
| `is_personal` | `bool` | Yes |  |
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

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Calendar(nil).Load(map[string]any{"id": "calendar_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Calendar(nil).Update(map[string]any{
    "id": "calendar_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CalendarEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CalendarAdminEntity

```go
calendarAdmin := client.CalendarAdmin(nil)
fmt.Println(calendarAdmin.GetName()) // "calendar_admin"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.CalendarAdmin(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CalendarAdminEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CalendarCouponEntity

```go
calendarCoupon := client.CalendarCoupon(nil)
fmt.Println(calendarCoupon.GetName()) // "calendar_coupon"
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
| `remaining_count` | `int` | Yes |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.CalendarCoupon(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.CalendarCoupon(nil).Create(map[string]any{
    "cents_off": "example_cents_off",
    "code": "example_code",
    "currency": "example_currency",
    "discount": "example_discount",
    "id": "example_id",
    "percent_off": "example_percent_off",
    "remaining_count": 1,
    "valid_end_at": "example_valid_end_at",
    "valid_start_at": "example_valid_start_at",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.CalendarCoupon(nil).Update(map[string]any{
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CalendarCouponEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CalendarEventEntity

```go
calendarEvent := client.CalendarEvent(nil)
fmt.Println(calendarEvent.GetName()) // "calendar_event"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `string` | Yes |  |
| `status` | `string` | Yes |  |
| `submitted_by` | `any` | Yes |  |
| `tag` | `[]any` | Yes |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.CalendarEvent(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.CalendarEvent(nil).Load(map[string]any{"id": "calendar_event_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.CalendarEvent(nil).Create(map[string]any{
    "id": "example_id",
    "status": "example_status",
    "submitted_by": "example_submitted_by",
    "tag": []any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CalendarEventEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CalendarEventApprovalEntity

```go
calendarEventApproval := client.CalendarEventApproval(nil)
fmt.Println(calendarEventApproval.GetName()) // "calendar_event_approval"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `calendar_event_id` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.CalendarEventApproval(nil).Create(map[string]any{
    "calendar_event_id": "example_calendar_event_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CalendarEventApprovalEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CalendarEventRejectionEntity

```go
calendarEventRejection := client.CalendarEventRejection(nil)
fmt.Println(calendarEventRejection.GetName()) // "calendar_event_rejection"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `calendar_event_id` | `string` | Yes |  |
| `message` | `string` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.CalendarEventRejection(nil).Create(map[string]any{
    "calendar_event_id": "example_calendar_event_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CalendarEventRejectionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ContactEntity

```go
contact := client.Contact(nil)
fmt.Println(contact.GetName()) // "contact"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avatar_url` | `string` | Yes |  |
| `contact` | `[]any` | Yes |  |
| `created_at` | `string` | Yes |  |
| `email` | `string` | Yes |  |
| `event_approved_count` | `float64` | Yes |  |
| `event_checked_in_count` | `float64` | Yes |  |
| `first_name` | `any` | Yes |  |
| `id` | `string` | Yes |  |
| `last_name` | `any` | Yes |  |
| `membership` | `any` | Yes |  |
| `name` | `string` | Yes |  |
| `revenue_usd_cent` | `float64` | Yes |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Contact(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Contact(nil).Create(map[string]any{
    "avatar_url": "example_avatar_url",
    "contact": []any{},
    "created_at": "example_created_at",
    "email": "example_email",
    "event_approved_count": 1,
    "event_checked_in_count": 1,
    "first_name": "example_first_name",
    "id": "example_id",
    "last_name": "example_last_name",
    "membership": "example_membership",
    "name": "example_name",
    "revenue_usd_cent": 1,
    "user_id": "example_user_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Contact(nil).Remove(map[string]any{"id": "id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ContactEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ContactBlockEntity

```go
contactBlock := client.ContactBlock(nil)
fmt.Println(contactBlock.GetName()) // "contact_block"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `contact_id` | `string` | No |  |
| `email` | `string` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.ContactBlock(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ContactBlockEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ContactRestoreEntity

```go
contactRestore := client.ContactRestore(nil)
fmt.Println(contactRestore.GetName()) // "contact_restore"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `contact_id` | `string` | No |  |
| `email` | `string` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.ContactRestore(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ContactRestoreEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ContactTagEntity

```go
contactTag := client.ContactTag(nil)
fmt.Println(contactTag.GetName()) // "contact_tag"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.ContactTag(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.ContactTag(nil).Create(map[string]any{
    "id": "example_id",
    "name": "example_name",
    "tag_id": "example_tag_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.ContactTag(nil).Update(map[string]any{
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.ContactTag(nil).Remove(map[string]any{"id": "id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ContactTagEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ContactTagAssignmentEntity

```go
contactTagAssignment := client.ContactTagAssignment(nil)
fmt.Println(contactTagAssignment.GetName()) // "contact_tag_assignment"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `applied_count` | `float64` | Yes |  |
| `email` | `[]any` | No |  |
| `skipped_count` | `float64` | Yes |  |
| `tag` | `string` | Yes |  |
| `user_id` | `[]any` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.ContactTagAssignment(nil).Create(map[string]any{
    "applied_count": 1,
    "skipped_count": 1,
    "tag": "example_tag",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.ContactTagAssignment(nil).Remove(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ContactTagAssignmentEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EntityLookupEntity

```go
entityLookup := client.EntityLookup(nil)
fmt.Println(entityLookup.GetName()) // "entity_lookup"
```

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.EntityLookup(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EntityLookupEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EventEntity

```go
event := client.Event(nil)
fmt.Println(event.GetName()) // "event"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `access` | `string` | Yes |  |
| `calendar_id` | `string` | Yes |  |
| `can_register_for_multiple_ticket` | `bool` | No |  |
| `coordinate` | `any` | Yes |  |
| `cover_url` | `string` | Yes |  |
| `created_at` | `string` | Yes |  |
| `description` | `string` | Yes |  |
| `description_md` | `string` | Yes |  |
| `display_price` | `any` | Yes |  |
| `duration_interval` | `string` | Yes |  |
| `end_at` | `string` | Yes |  |
| `event_id` | `string` | Yes |  |
| `feedback_email` | `map[string]any` | Yes |  |
| `geo_address_json` | `any` | Yes |  |
| `guest_count` | `map[string]any` | Yes |  |
| `host` | `[]any` | Yes |  |
| `id` | `string` | Yes |  |
| `location_type` | `string` | Yes |  |
| `location_visibility` | `string` | Yes |  |
| `max_capacity` | `any` | No |  |
| `meeting_url` | `any` | Yes |  |
| `name` | `string` | Yes |  |
| `name_requirement` | `string` | No |  |
| `phone_number_requirement` | `any` | No |  |
| `platform` | `string` | Yes |  |
| `registration_open` | `bool` | Yes |  |
| `registration_question` | `[]any` | No |  |
| `reminders_disabled` | `bool` | No |  |
| `require_approval` | `bool` | Yes |  |
| `show_guest_list` | `bool` | No |  |
| `slug` | `string` | No |  |
| `spots_remaining` | `any` | Yes |  |
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

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Event(nil).Load(map[string]any{"id": "event_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Event(nil).Create(map[string]any{
    "access": "example_access",
    "calendar_id": "example_calendar_id",
    "coordinate": "example_coordinate",
    "cover_url": "example_cover_url",
    "created_at": "example_created_at",
    "description": "example_description",
    "description_md": "example_description_md",
    "display_price": "example_display_price",
    "duration_interval": "example_duration_interval",
    "end_at": "example_end_at",
    "event_id": "example_event_id",
    "feedback_email": map[string]any{},
    "geo_address_json": "example_geo_address_json",
    "guest_count": map[string]any{},
    "host": []any{},
    "id": "example_id",
    "location_type": "example_location_type",
    "location_visibility": "example_location_visibility",
    "meeting_url": "example_meeting_url",
    "name": "example_name",
    "platform": "example_platform",
    "registration_open": true,
    "require_approval": true,
    "spots_remaining": "example_spots_remaining",
    "start_at": "example_start_at",
    "timezone": "example_timezone",
    "url": "example_url",
    "user_id": "example_user_id",
    "visibility": "example_visibility",
    "waitlist_status": "example_waitlist_status",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Event(nil).Update(map[string]any{
    "id": "event_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Event(nil).Remove(map[string]any{"id": "event_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EventEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EventCancelRequestEntity

```go
eventCancelRequest := client.EventCancelRequest(nil)
fmt.Println(eventCancelRequest.GetName()) // "event_cancel_request"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `cancellation_token` | `string` | Yes |  |
| `event_id` | `string` | Yes |  |
| `guest_count` | `float64` | Yes |  |
| `is_paid` | `bool` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.EventCancelRequest(nil).Create(map[string]any{
    "cancellation_token": "example_cancellation_token",
    "event_id": "example_event_id",
    "guest_count": 1,
    "is_paid": true,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EventCancelRequestEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EventCouponEntity

```go
eventCoupon := client.EventCoupon(nil)
fmt.Println(eventCoupon.GetName()) // "event_coupon"
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
| `remaining_count` | `int` | Yes |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.EventCoupon(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.EventCoupon(nil).Create(map[string]any{
    "cents_off": "example_cents_off",
    "code": "example_code",
    "currency": "example_currency",
    "discount": "example_discount",
    "event_id": "example_event_id",
    "id": "example_id",
    "percent_off": "example_percent_off",
    "remaining_count": 1,
    "valid_end_at": "example_valid_end_at",
    "valid_start_at": "example_valid_start_at",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.EventCoupon(nil).Update(map[string]any{
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EventCouponEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EventTagEntity

```go
eventTag := client.EventTag(nil)
fmt.Println(eventTag.GetName()) // "event_tag"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.EventTag(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.EventTag(nil).Create(map[string]any{
    "id": "example_id",
    "name": "example_name",
    "tag_id": "example_tag_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.EventTag(nil).Update(map[string]any{
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.EventTag(nil).Remove(map[string]any{"id": "id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EventTagEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EventTagAssignmentEntity

```go
eventTagAssignment := client.EventTagAssignment(nil)
fmt.Println(eventTagAssignment.GetName()) // "event_tag_assignment"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `applied_count` | `float64` | Yes |  |
| `event_id` | `[]any` | Yes |  |
| `skipped_count` | `float64` | Yes |  |
| `tag` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.EventTagAssignment(nil).Create(map[string]any{
    "applied_count": 1,
    "event_id": []any{},
    "skipped_count": 1,
    "tag": "example_tag",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.EventTagAssignment(nil).Remove(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EventTagAssignmentEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## GuestEntity

```go
guest := client.Guest(nil)
fmt.Println(guest.GetName()) // "guest"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_status` | `string` | Yes |  |
| `check_in_qr_code` | `string` | Yes |  |
| `eth_address` | `any` | Yes |  |
| `event_id` | `string` | Yes |  |
| `event_ticket` | `[]any` | Yes |  |
| `event_ticket_order` | `[]any` | Yes |  |
| `guest` | `[]any` | Yes |  |
| `guest_id` | `string` | Yes |  |
| `id` | `string` | Yes |  |
| `invited_at` | `any` | Yes |  |
| `joined_at` | `any` | Yes |  |
| `message` | `any` | No |  |
| `phone_number` | `int` | Yes |  |
| `registered_at` | `any` | Yes |  |
| `registration_answer` | `any` | Yes |  |
| `send_email` | `any` | No |  |
| `should_refund` | `bool` | No |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Guest(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Guest(nil).Load(map[string]any{"id": "guest_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Guest(nil).Create(map[string]any{
    "approval_status": "example_approval_status",
    "check_in_qr_code": "example_check_in_qr_code",
    "eth_address": "example_eth_address",
    "event_id": "example_event_id",
    "event_ticket": []any{},
    "event_ticket_order": []any{},
    "guest": []any{},
    "guest_id": "example_guest_id",
    "id": "example_id",
    "invited_at": "example_invited_at",
    "joined_at": "example_joined_at",
    "phone_number": 1,
    "registered_at": "example_registered_at",
    "registration_answer": "example_registration_answer",
    "solana_address": "example_solana_address",
    "status": "example_status",
    "user_email": "example_user_email",
    "user_first_name": "example_user_first_name",
    "user_id": "example_user_id",
    "user_last_name": "example_user_last_name",
    "user_name": "example_user_name",
    "utm_source": "example_utm_source",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Guest(nil).Update(map[string]any{
    "id": "guest_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `GuestEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## GuestInviteEntity

```go
guestInvite := client.GuestInvite(nil)
fmt.Println(guestInvite.GetName()) // "guest_invite"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `event_id` | `string` | Yes |  |
| `guest` | `[]any` | Yes |  |
| `message` | `any` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.GuestInvite(nil).Create(map[string]any{
    "event_id": "example_event_id",
    "guest": []any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `GuestInviteEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## GuestTicketEntity

```go
guestTicket := client.GuestTicket(nil)
fmt.Println(guestTicket.GetName()) // "guest_ticket"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `event_id` | `string` | Yes |  |
| `guest_id` | `string` | Yes |  |
| `send_email` | `any` | No |  |
| `ticket_ids_to_remove` | `[]any` | No |  |
| `tickets_to_add` | `[]any` | No |  |

### Operations

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.GuestTicket(nil).Update(map[string]any{
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `GuestTicketEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## HostEntity

```go
host := client.Host(nil)
fmt.Println(host.GetName()) // "host"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `access_level` | `any` | No |  |
| `email` | `string` | Yes |  |
| `event_id` | `string` | Yes |  |
| `is_visible` | `bool` | No |  |
| `name` | `string` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Host(nil).Create(map[string]any{
    "email": "example_email",
    "event_id": "example_event_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Host(nil).Update(map[string]any{
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Host(nil).Remove(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `HostEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ImageUploadEntity

```go
imageUpload := client.ImageUpload(nil)
fmt.Println(imageUpload.GetName()) // "image_upload"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `content_type` | `any` | No |  |
| `file_url` | `string` | Yes |  |
| `upload_url` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.ImageUpload(nil).Create(map[string]any{
    "file_url": "example_file_url",
    "upload_url": "example_upload_url",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ImageUploadEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## MemberEntity

```go
member := client.Member(nil)
fmt.Println(member.GetName()) // "member"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `email` | `string` | Yes |  |
| `membership_id` | `string` | Yes |  |
| `membership_tier_id` | `string` | Yes |  |
| `registration_answer` | `[]any` | No |  |
| `skip_payment` | `bool` | No |  |
| `status` | `string` | Yes |  |
| `user_id` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Member(nil).Create(map[string]any{
    "email": "example_email",
    "membership_id": "example_membership_id",
    "membership_tier_id": "example_membership_tier_id",
    "status": "example_status",
    "user_id": "example_user_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Member(nil).Update(map[string]any{
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `MemberEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## MembershipTierEntity

```go
membershipTier := client.MembershipTier(nil)
fmt.Println(membershipTier.GetName()) // "membership_tier"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.MembershipTier(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `MembershipTierEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## OrganizationAdminEntity

```go
organizationAdmin := client.OrganizationAdmin(nil)
fmt.Println(organizationAdmin.GetName()) // "organization_admin"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.OrganizationAdmin(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `OrganizationAdminEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## OrganizationCalendarEntity

```go
organizationCalendar := client.OrganizationCalendar(nil)
fmt.Println(organizationCalendar.GetName()) // "organization_calendar"
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
| `is_personal` | `bool` | Yes |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.OrganizationCalendar(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.OrganizationCalendar(nil).Create(map[string]any{
    "avatar_url": "example_avatar_url",
    "coordinate": "example_coordinate",
    "cover_image_url": "example_cover_image_url",
    "description": "example_description",
    "id": "example_id",
    "instagram_handle": "example_instagram_handle",
    "is_personal": true,
    "location": "example_location",
    "name": "example_name",
    "slug": "example_slug",
    "social_image_url": "example_social_image_url",
    "twitter_handle": "example_twitter_handle",
    "url": "example_url",
    "website": "example_website",
    "youtube_handle": "example_youtube_handle",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `OrganizationCalendarEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## OrganizationEventEntity

```go
organizationEvent := client.OrganizationEvent(nil)
fmt.Println(organizationEvent.GetName()) // "organization_event"
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
| `feedback_email` | `map[string]any` | Yes |  |
| `geo_address_json` | `any` | Yes |  |
| `geo_latitude` | `any` | Yes |  |
| `geo_longitude` | `any` | Yes |  |
| `id` | `string` | Yes |  |
| `location_type` | `string` | Yes |  |
| `location_visibility` | `string` | Yes |  |
| `managing_calendar` | `[]any` | Yes |  |
| `meeting_url` | `any` | Yes |  |
| `name` | `string` | Yes |  |
| `platform` | `string` | Yes |  |
| `registration_open` | `bool` | Yes |  |
| `registration_question` | `[]any` | No |  |
| `require_approval` | `bool` | Yes |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.OrganizationEvent(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `OrganizationEventEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## OrganizationEventTransferEntity

```go
organizationEventTransfer := client.OrganizationEventTransfer(nil)
fmt.Println(organizationEventTransfer.GetName()) // "organization_event_transfer"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `calendar_id` | `string` | Yes |  |
| `event_id` | `string` | Yes |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.OrganizationEventTransfer(nil).Create(map[string]any{
    "calendar_id": "example_calendar_id",
    "event_id": "example_event_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `OrganizationEventTransferEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## TicketTypeEntity

```go
ticketType := client.TicketType(nil)
fmt.Println(ticketType.GetName()) // "ticket_type"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `cent` | `any` | No |  |
| `currency` | `any` | No |  |
| `description` | `string` | No |  |
| `id` | `string` | Yes |  |
| `is_flexible` | `bool` | No |  |
| `is_hidden` | `bool` | No |  |
| `max_capacity` | `any` | No |  |
| `min_cent` | `any` | No |  |
| `name` | `string` | Yes |  |
| `require_approval` | `bool` | No |  |
| `type` | `string` | Yes |  |
| `valid_end_at` | `any` | No |  |
| `valid_start_at` | `any` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.TicketType(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.TicketType(nil).Load(map[string]any{"id": "ticket_type_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.TicketType(nil).Create(map[string]any{
    "id": "example_id",
    "name": "example_name",
    "type": "example_type",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.TicketType(nil).Update(map[string]any{
    "id": "ticket_type_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.TicketType(nil).Remove(map[string]any{"id": "ticket_type_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `TicketTypeEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## UserEntity

```go
user := client.User(nil)
fmt.Println(user.GetName()) // "user"
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

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.User(nil).Load(map[string]any{"id": "user_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `UserEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## WebhookEntity

```go
webhook := client.Webhook(nil)
fmt.Println(webhook.GetName()) // "webhook"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | Yes |  |
| `event_type` | `[]any` | Yes |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Webhook(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Webhook(nil).Load(map[string]any{"id": "webhook_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Webhook(nil).Create(map[string]any{
    "created_at": "example_created_at",
    "event_type": []any{},
    "id": "example_id",
    "secret": "example_secret",
    "status": "example_status",
    "url": "example_url",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Webhook(nil).Update(map[string]any{
    "id": "webhook_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Webhook(nil).Remove(map[string]any{"id": "webhook_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `WebhookEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```go
client := sdk.NewLumaSDK(map[string]any{
    "feature": map[string]any{
        "test": map[string]any{"active": true},
    },
})
```

