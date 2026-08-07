# Luma Golang SDK



The Golang SDK for the Luma API — an entity-oriented client using standard Go conventions. No generics required; data flows as `map[string]any`.

It exposes the API as capitalised, semantic **Entities** — e.g. `client.Calendar(nil)` — each with the same small set of operations (`List`, `Load`, `Create`, `Update`, `Remove`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
```bash
go get github.com/voxgig-sdk/luma-sdk/go@latest
```

The Go module proxy resolves the version from the `go/vX.Y.Z` GitHub
release tag — see [Releases](https://github.com/voxgig-sdk/luma-sdk/releases) for the available versions.

To vendor from a local checkout instead, clone this repo alongside your
project and add a `replace` directive pointing at the checked-out
`go/` directory:

```bash
go mod edit -replace github.com/voxgig-sdk/luma-sdk/go=../luma-sdk/go
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### Quickstart

A complete program: create a client, then call the entity operations.
Each operation returns `(value, error)` — the value is the data itself
(there is no `{ok, data}` wrapper), so check `err` and use the value
directly.

```go
package main

import (
    "fmt"
    "os"
    sdk "github.com/voxgig-sdk/luma-sdk/go"
)

func main() {
    client := sdk.NewLumaSDK(map[string]any{
        "apikey": os.Getenv("LUMA_APIKEY"),
    })

    // Load a single calendar — the value is the loaded record.
    calendar, err := client.Calendar(nil).Load(map[string]any{"id": "example_id"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(calendar)

    // Update a calendar.
    updated, err := client.Calendar(nil).Update(map[string]any{"id": "example_id", "avatar_url": "example_avatar_url", "calendar_id": "example_calendar_id"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(updated)
}
```


## Error handling

Every entity operation returns `(value, error)`. Check `err` before
using the value — there is no exception to catch:

```go
eventcoupons, err := client.EventCoupon(nil).List(nil, nil)
if err != nil {
    // handle err
    return
}
_ = eventcoupons
```

`Direct` follows the same `(value, error)` convention:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example_id"},
})
if err != nil {
    // handle err
}
_ = result
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

if result["ok"] == true {
    fmt.Println(result["status"]) // 200
    fmt.Println(result["data"])   // response body
}
```

### Prepare a request without sending it

```go
fetchdef, err := client.Prepare(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "DELETE",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

fmt.Println(fetchdef["url"])
fmt.Println(fetchdef["method"])
fmt.Println(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```go
client := sdk.Test()

eventCoupon, err := client.EventCoupon(nil).List(
    nil, nil,
)
if err != nil {
    panic(err)
}
fmt.Println(eventCoupon) // the returned mock data
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```go
mockFetch := func(url string, init map[string]any) (map[string]any, error) {
    return map[string]any{
        "status":     200,
        "statusText": "OK",
        "headers":    map[string]any{},
        "json": (func() any)(func() any {
            return map[string]any{"id": "mock01"}
        }),
    }, nil
}

client := sdk.NewLumaSDK(map[string]any{
    "base": "http://localhost:8080",
    "system": map[string]any{
        "fetch": (func(string, map[string]any) (map[string]any, error))(mockFetch),
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
cd go && go test ./test/...
```


## Reference

### NewLumaSDK

```go
func NewLumaSDK(options map[string]any) *LumaSDK
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `"apikey"` | `string` | API key for authentication. |
| `"base"` | `string` | Base URL of the API server. |
| `"prefix"` | `string` | URL path prefix prepended to all requests. |
| `"suffix"` | `string` | URL path suffix appended to all requests. |
| `"feature"` | `map[string]any` | Feature activation flags. |
| `"extend"` | `[]any` | Additional Feature instances to load. |
| `"system"` | `map[string]any` | System overrides (e.g. custom `"fetch"` function). |

### TestSDK

```go
func TestSDK(testopts map[string]any, sdkopts map[string]any) *LumaSDK
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### LumaSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `OptionsMap` | `() map[string]any` | Deep copy of current SDK options. |
| `GetUtility` | `() *Utility` | Copy of the SDK utility object. |
| `Prepare` | `(fetchargs map[string]any) (map[string]any, error)` | Build an HTTP request definition without sending. |
| `Direct` | `(fetchargs map[string]any) (map[string]any, error)` | Build and send an HTTP request. |
| `Calendar` | `(data map[string]any) LumaEntity` | Create a Calendar entity instance. |
| `CalendarAdmin` | `(data map[string]any) LumaEntity` | Create a CalendarAdmin entity instance. |
| `CalendarCoupon` | `(data map[string]any) LumaEntity` | Create a CalendarCoupon entity instance. |
| `CalendarEvent` | `(data map[string]any) LumaEntity` | Create a CalendarEvent entity instance. |
| `CalendarEventApproval` | `(data map[string]any) LumaEntity` | Create a CalendarEventApproval entity instance. |
| `CalendarEventRejection` | `(data map[string]any) LumaEntity` | Create a CalendarEventRejection entity instance. |
| `Contact` | `(data map[string]any) LumaEntity` | Create a Contact entity instance. |
| `ContactBlock` | `(data map[string]any) LumaEntity` | Create a ContactBlock entity instance. |
| `ContactRestore` | `(data map[string]any) LumaEntity` | Create a ContactRestore entity instance. |
| `ContactTag` | `(data map[string]any) LumaEntity` | Create a ContactTag entity instance. |
| `ContactTagAssignment` | `(data map[string]any) LumaEntity` | Create a ContactTagAssignment entity instance. |
| `EntityLookup` | `(data map[string]any) LumaEntity` | Create an EntityLookup entity instance. |
| `Event` | `(data map[string]any) LumaEntity` | Create an Event entity instance. |
| `EventCancelRequest` | `(data map[string]any) LumaEntity` | Create an EventCancelRequest entity instance. |
| `EventCoupon` | `(data map[string]any) LumaEntity` | Create an EventCoupon entity instance. |
| `EventTag` | `(data map[string]any) LumaEntity` | Create an EventTag entity instance. |
| `EventTagAssignment` | `(data map[string]any) LumaEntity` | Create an EventTagAssignment entity instance. |
| `Guest` | `(data map[string]any) LumaEntity` | Create a Guest entity instance. |
| `GuestInvite` | `(data map[string]any) LumaEntity` | Create a GuestInvite entity instance. |
| `GuestTicket` | `(data map[string]any) LumaEntity` | Create a GuestTicket entity instance. |
| `Host` | `(data map[string]any) LumaEntity` | Create a Host entity instance. |
| `ImageUpload` | `(data map[string]any) LumaEntity` | Create an ImageUpload entity instance. |
| `Member` | `(data map[string]any) LumaEntity` | Create a Member entity instance. |
| `MembershipTier` | `(data map[string]any) LumaEntity` | Create a MembershipTier entity instance. |
| `OrganizationAdmin` | `(data map[string]any) LumaEntity` | Create an OrganizationAdmin entity instance. |
| `OrganizationCalendar` | `(data map[string]any) LumaEntity` | Create an OrganizationCalendar entity instance. |
| `OrganizationEvent` | `(data map[string]any) LumaEntity` | Create an OrganizationEvent entity instance. |
| `OrganizationEventTransfer` | `(data map[string]any) LumaEntity` | Create an OrganizationEventTransfer entity instance. |
| `TicketType` | `(data map[string]any) LumaEntity` | Create a TicketType entity instance. |
| `User` | `(data map[string]any) LumaEntity` | Create an User entity instance. |
| `Webhook` | `(data map[string]any) LumaEntity` | Create a Webhook entity instance. |

### Entity interface (LumaEntity)

All entities implement the `LumaEntity` interface.

| Method | Signature | Description |
| --- | --- | --- |
| `Load` | `(reqmatch, ctrl map[string]any) (any, error)` | Load a single entity by match criteria. |
| `List` | `(reqmatch, ctrl map[string]any) (any, error)` | List entities matching the criteria. |
| `Create` | `(reqdata, ctrl map[string]any) (any, error)` | Create a new entity. |
| `Update` | `(reqdata, ctrl map[string]any) (any, error)` | Update an existing entity. |
| `Remove` | `(reqmatch, ctrl map[string]any) (any, error)` | Remove an entity. |
| `Data` | `(args ...any) any` | Get or set entity data. |
| `Match` | `(args ...any) any` | Get or set entity match criteria. |
| `Make` | `() Entity` | Create a new instance with the same options. |
| `GetName` | `() string` | Return the entity name. |

### Result shape

Entity operations return `(value, error)`. The `value` is the
operation's data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `Load` / `Create` / `Update` / `Remove` | the entity record (`map[string]any`) |
| `List` | a `[]any` of entity records |

Check `err` first, then use the value directly (or the typed
`...Typed` variants, which return the entity's model struct and a typed
slice):

    calendar, err := client.Calendar(nil).Load(map[string]any{"id": "example_id"}, nil)
    if err != nil { /* handle */ }
    // calendar is the returned record

Only `Direct()` returns a response envelope — a `map[string]any` with
`"ok"`, `"status"`, `"headers"`, and `"data"` keys.

### Entities

#### Calendar

| Field | Description |
| --- | --- |
| `"avatar_url"` |  |
| `"calendar_id"` |  |
| `"coordinate"` |  |
| `"cover_image_url"` |  |
| `"description"` |  |
| `"id"` |  |
| `"instagram_handle"` |  |
| `"is_personal"` |  |
| `"location"` |  |
| `"name"` |  |
| `"slug"` |  |
| `"social_image_url"` |  |
| `"tint_color"` |  |
| `"twitter_handle"` |  |
| `"url"` |  |
| `"website"` |  |
| `"youtube_handle"` |  |

Operations: Load, Update.

API path: `/v1/calendars/get`

#### CalendarAdmin

| Field | Description |
| --- | --- |
| `"avatar_url"` |  |
| `"email"` |  |
| `"first_name"` |  |
| `"id"` |  |
| `"last_name"` |  |
| `"name"` |  |

Operations: List.

API path: `/v1/calendars/admins/list`

#### CalendarCoupon

| Field | Description |
| --- | --- |
| `"cents_off"` |  |
| `"code"` |  |
| `"currency"` |  |
| `"discount"` |  |
| `"event_ticket_type_id"` |  |
| `"id"` |  |
| `"percent_off"` |  |
| `"remaining_count"` |  |
| `"valid_end_at"` |  |
| `"valid_start_at"` |  |

Operations: Create, List, Update.

API path: `/v1/calendars/coupons/create`

#### CalendarEvent

| Field | Description |
| --- | --- |
| `"id"` |  |
| `"status"` |  |
| `"submitted_by"` |  |
| `"tag"` |  |

Operations: Create, List, Load.

API path: `/v1/calendars/events/add`

#### CalendarEventApproval

| Field | Description |
| --- | --- |
| `"calendar_event_id"` |  |

Operations: Create.

API path: `/v1/calendars/events/approve`

#### CalendarEventRejection

| Field | Description |
| --- | --- |
| `"calendar_event_id"` |  |
| `"message"` |  |

Operations: Create.

API path: `/v1/calendars/events/reject`

#### Contact

| Field | Description |
| --- | --- |
| `"avatar_url"` |  |
| `"contact"` |  |
| `"created_at"` |  |
| `"email"` |  |
| `"event_approved_count"` |  |
| `"event_checked_in_count"` |  |
| `"first_name"` |  |
| `"id"` |  |
| `"last_name"` |  |
| `"membership"` |  |
| `"name"` |  |
| `"revenue_usd_cent"` |  |
| `"tag"` |  |
| `"user_id"` |  |

Operations: Create, List, Remove.

API path: `/v1/calendars/contacts/import`

#### ContactBlock

| Field | Description |
| --- | --- |
| `"contact_id"` |  |
| `"email"` |  |

Operations: Create.

API path: `/v1/calendars/contacts/block`

#### ContactRestore

| Field | Description |
| --- | --- |
| `"contact_id"` |  |
| `"email"` |  |

Operations: Create.

API path: `/v1/calendars/contacts/restore`

#### ContactTag

| Field | Description |
| --- | --- |
| `"color"` |  |
| `"id"` |  |
| `"name"` |  |
| `"tag_id"` |  |

Operations: Create, List, Remove, Update.

API path: `/v1/calendars/contact-tags/create`

#### ContactTagAssignment

| Field | Description |
| --- | --- |
| `"applied_count"` |  |
| `"email"` |  |
| `"skipped_count"` |  |
| `"tag"` |  |
| `"user_id"` |  |

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
| `"access"` |  |
| `"calendar_id"` |  |
| `"can_register_for_multiple_ticket"` |  |
| `"coordinate"` |  |
| `"cover_url"` |  |
| `"created_at"` |  |
| `"description"` |  |
| `"description_md"` |  |
| `"display_price"` |  |
| `"duration_interval"` |  |
| `"end_at"` |  |
| `"event_id"` |  |
| `"feedback_email"` |  |
| `"geo_address_json"` |  |
| `"guest_count"` |  |
| `"host"` |  |
| `"id"` |  |
| `"location_type"` |  |
| `"location_visibility"` |  |
| `"max_capacity"` |  |
| `"meeting_url"` |  |
| `"name"` |  |
| `"name_requirement"` |  |
| `"phone_number_requirement"` |  |
| `"platform"` |  |
| `"registration_open"` |  |
| `"registration_question"` |  |
| `"reminders_disabled"` |  |
| `"require_approval"` |  |
| `"show_guest_list"` |  |
| `"slug"` |  |
| `"spots_remaining"` |  |
| `"start_at"` |  |
| `"suppress_notification"` |  |
| `"timezone"` |  |
| `"tint_color"` |  |
| `"url"` |  |
| `"user_id"` |  |
| `"visibility"` |  |
| `"waitlist_status"` |  |

Operations: Create, Load, Remove, Update.

API path: `/v1/events/create`

#### EventCancelRequest

| Field | Description |
| --- | --- |
| `"cancellation_token"` |  |
| `"event_id"` |  |
| `"guest_count"` |  |
| `"is_paid"` |  |

Operations: Create.

API path: `/v1/events/cancel/request`

#### EventCoupon

| Field | Description |
| --- | --- |
| `"cents_off"` |  |
| `"code"` |  |
| `"currency"` |  |
| `"discount"` |  |
| `"event_id"` |  |
| `"event_ticket_type_id"` |  |
| `"id"` |  |
| `"percent_off"` |  |
| `"remaining_count"` |  |
| `"valid_end_at"` |  |
| `"valid_start_at"` |  |

Operations: Create, List, Update.

API path: `/v1/events/coupons/create`

#### EventTag

| Field | Description |
| --- | --- |
| `"color"` |  |
| `"id"` |  |
| `"name"` |  |
| `"tag_id"` |  |

Operations: Create, List, Remove, Update.

API path: `/v1/calendars/event-tags/create`

#### EventTagAssignment

| Field | Description |
| --- | --- |
| `"applied_count"` |  |
| `"event_id"` |  |
| `"skipped_count"` |  |
| `"tag"` |  |

Operations: Create, Remove.

API path: `/v1/calendars/event-tags/apply`

#### Guest

| Field | Description |
| --- | --- |
| `"approval_status"` |  |
| `"check_in_qr_code"` |  |
| `"eth_address"` |  |
| `"event_id"` |  |
| `"event_ticket"` |  |
| `"event_ticket_order"` |  |
| `"guest"` |  |
| `"guest_id"` |  |
| `"id"` |  |
| `"invited_at"` |  |
| `"joined_at"` |  |
| `"message"` |  |
| `"phone_number"` |  |
| `"registered_at"` |  |
| `"registration_answer"` |  |
| `"send_email"` |  |
| `"should_refund"` |  |
| `"solana_address"` |  |
| `"status"` |  |
| `"ticket"` |  |
| `"user_email"` |  |
| `"user_first_name"` |  |
| `"user_id"` |  |
| `"user_last_name"` |  |
| `"user_name"` |  |
| `"utm_source"` |  |

Operations: Create, List, Load, Update.

API path: `/v1/events/guests/add`

#### GuestInvite

| Field | Description |
| --- | --- |
| `"event_id"` |  |
| `"guest"` |  |
| `"message"` |  |

Operations: Create.

API path: `/v1/events/guests/send-invites`

#### GuestTicket

| Field | Description |
| --- | --- |
| `"event_id"` |  |
| `"guest_id"` |  |
| `"send_email"` |  |
| `"ticket_ids_to_remove"` |  |
| `"tickets_to_add"` |  |

Operations: Update.

API path: `/v1/events/guests/update-tickets`

#### Host

| Field | Description |
| --- | --- |
| `"access_level"` |  |
| `"email"` |  |
| `"event_id"` |  |
| `"is_visible"` |  |
| `"name"` |  |

Operations: Create, Remove, Update.

API path: `/v1/events/hosts/add`

#### ImageUpload

| Field | Description |
| --- | --- |
| `"content_type"` |  |
| `"file_url"` |  |
| `"upload_url"` |  |

Operations: Create.

API path: `/v1/images/create-upload-url`

#### Member

| Field | Description |
| --- | --- |
| `"email"` |  |
| `"membership_id"` |  |
| `"membership_tier_id"` |  |
| `"registration_answer"` |  |
| `"skip_payment"` |  |
| `"status"` |  |
| `"user_id"` |  |

Operations: Create, Update.

API path: `/v1/memberships/members/add`

#### MembershipTier

| Field | Description |
| --- | --- |
| `"access_info"` |  |
| `"description"` |  |
| `"id"` |  |
| `"name"` |  |
| `"tint_color"` |  |

Operations: List.

API path: `/v1/memberships/tiers/list`

#### OrganizationAdmin

| Field | Description |
| --- | --- |
| `"api_id"` |  |
| `"avatar_url"` |  |
| `"email"` |  |
| `"first_name"` |  |
| `"id"` |  |
| `"last_name"` |  |
| `"name"` |  |

Operations: List.

API path: `/v1/organizations/admins/list`

#### OrganizationCalendar

| Field | Description |
| --- | --- |
| `"avatar_url"` |  |
| `"coordinate"` |  |
| `"cover_image_url"` |  |
| `"description"` |  |
| `"id"` |  |
| `"instagram_handle"` |  |
| `"is_personal"` |  |
| `"location"` |  |
| `"name"` |  |
| `"slug"` |  |
| `"social_image_url"` |  |
| `"tint_color"` |  |
| `"twitter_handle"` |  |
| `"url"` |  |
| `"website"` |  |
| `"youtube_handle"` |  |

Operations: Create, List.

API path: `/v2/organizations/calendars/create`

#### OrganizationEvent

| Field | Description |
| --- | --- |
| `"api_id"` |  |
| `"calendar_api_id"` |  |
| `"calendar_id"` |  |
| `"coordinate"` |  |
| `"cover_url"` |  |
| `"created_at"` |  |
| `"display_price"` |  |
| `"duration_interval"` |  |
| `"end_at"` |  |
| `"feedback_email"` |  |
| `"geo_address_json"` |  |
| `"geo_latitude"` |  |
| `"geo_longitude"` |  |
| `"id"` |  |
| `"location_type"` |  |
| `"location_visibility"` |  |
| `"managing_calendar"` |  |
| `"meeting_url"` |  |
| `"name"` |  |
| `"platform"` |  |
| `"registration_open"` |  |
| `"registration_question"` |  |
| `"require_approval"` |  |
| `"spots_remaining"` |  |
| `"start_at"` |  |
| `"timezone"` |  |
| `"url"` |  |
| `"user_api_id"` |  |
| `"user_id"` |  |
| `"visibility"` |  |
| `"waitlist_status"` |  |
| `"zoom_meeting_url"` |  |

Operations: List.

API path: `/v1/organizations/events/list`

#### OrganizationEventTransfer

| Field | Description |
| --- | --- |
| `"calendar_id"` |  |
| `"event_id"` |  |

Operations: Create.

API path: `/v1/organizations/events/transfer-calendar`

#### TicketType

| Field | Description |
| --- | --- |
| `"cent"` |  |
| `"currency"` |  |
| `"description"` |  |
| `"id"` |  |
| `"is_flexible"` |  |
| `"is_hidden"` |  |
| `"max_capacity"` |  |
| `"min_cent"` |  |
| `"name"` |  |
| `"require_approval"` |  |
| `"type"` |  |
| `"valid_end_at"` |  |
| `"valid_start_at"` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/v1/events/ticket-types/create`

#### User

| Field | Description |
| --- | --- |
| `"avatar_url"` |  |
| `"email"` |  |
| `"first_name"` |  |
| `"id"` |  |
| `"last_name"` |  |
| `"name"` |  |

Operations: Load.

API path: `/v1/users/get-self`

#### Webhook

| Field | Description |
| --- | --- |
| `"created_at"` |  |
| `"event_type"` |  |
| `"id"` |  |
| `"secret"` |  |
| `"status"` |  |
| `"url"` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/v2/webhooks/create`



## Entities


### Calendar

Create an instance: `calendar := client.Calendar(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Update(data, ctrl)` | Update an existing entity. |

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
| `is_personal` | `bool` |  |
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

```go
calendar, err := client.Calendar(nil).Load(map[string]any{"id": "calendar_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(calendar) // the loaded record
```


### CalendarAdmin

Create an instance: `calendarAdmin := client.CalendarAdmin(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

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

```go
calendarAdmins, err := client.CalendarAdmin(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(calendarAdmins) // the array of records
```


### CalendarCoupon

Create an instance: `calendarCoupon := client.CalendarCoupon(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |

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
| `remaining_count` | `int` |  |
| `valid_end_at` | `any` |  |
| `valid_start_at` | `any` |  |

#### Example: List

```go
calendarCoupons, err := client.CalendarCoupon(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(calendarCoupons) // the array of records
```

#### Example: Create

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


### CalendarEvent

Create an instance: `calendarEvent := client.CalendarEvent(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `string` |  |
| `status` | `string` |  |
| `submitted_by` | `any` |  |
| `tag` | `[]any` |  |

#### Example: Load

```go
calendarEvent, err := client.CalendarEvent(nil).Load(map[string]any{"id": "calendar_event_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(calendarEvent) // the loaded record
```

#### Example: List

```go
calendarEvents, err := client.CalendarEvent(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(calendarEvents) // the array of records
```

#### Example: Create

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


### CalendarEventApproval

Create an instance: `calendarEventApproval := client.CalendarEventApproval(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `calendar_event_id` | `string` |  |

#### Example: Create

```go
result, err := client.CalendarEventApproval(nil).Create(map[string]any{
    "calendar_event_id": "example_calendar_event_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### CalendarEventRejection

Create an instance: `calendarEventRejection := client.CalendarEventRejection(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `calendar_event_id` | `string` |  |
| `message` | `string` |  |

#### Example: Create

```go
result, err := client.CalendarEventRejection(nil).Create(map[string]any{
    "calendar_event_id": "example_calendar_event_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### Contact

Create an instance: `contact := client.Contact(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avatar_url` | `string` |  |
| `contact` | `[]any` |  |
| `created_at` | `string` |  |
| `email` | `string` |  |
| `event_approved_count` | `float64` |  |
| `event_checked_in_count` | `float64` |  |
| `first_name` | `any` |  |
| `id` | `string` |  |
| `last_name` | `any` |  |
| `membership` | `any` |  |
| `name` | `string` |  |
| `revenue_usd_cent` | `float64` |  |
| `tag` | `any` |  |
| `user_id` | `string` |  |

#### Example: List

```go
contacts, err := client.Contact(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(contacts) // the array of records
```

#### Example: Create

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


### ContactBlock

Create an instance: `contactBlock := client.ContactBlock(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `contact_id` | `string` |  |
| `email` | `string` |  |

#### Example: Create

```go
result, err := client.ContactBlock(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### ContactRestore

Create an instance: `contactRestore := client.ContactRestore(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `contact_id` | `string` |  |
| `email` | `string` |  |

#### Example: Create

```go
result, err := client.ContactRestore(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### ContactTag

Create an instance: `contactTag := client.ContactTag(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `color` | `any` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `tag_id` | `string` |  |

#### Example: List

```go
contactTags, err := client.ContactTag(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(contactTags) // the array of records
```

#### Example: Create

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


### ContactTagAssignment

Create an instance: `contactTagAssignment := client.ContactTagAssignment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `applied_count` | `float64` |  |
| `email` | `[]any` |  |
| `skipped_count` | `float64` |  |
| `tag` | `string` |  |
| `user_id` | `[]any` |  |

#### Example: Create

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


### EntityLookup

Create an instance: `entityLookup := client.EntityLookup(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Example: Load

```go
entityLookup, err := client.EntityLookup(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(entityLookup) // the loaded record
```


### Event

Create an instance: `event := client.Event(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `access` | `string` |  |
| `calendar_id` | `string` |  |
| `can_register_for_multiple_ticket` | `bool` |  |
| `coordinate` | `any` |  |
| `cover_url` | `string` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `description_md` | `string` |  |
| `display_price` | `any` |  |
| `duration_interval` | `string` |  |
| `end_at` | `string` |  |
| `event_id` | `string` |  |
| `feedback_email` | `map[string]any` |  |
| `geo_address_json` | `any` |  |
| `guest_count` | `map[string]any` |  |
| `host` | `[]any` |  |
| `id` | `string` |  |
| `location_type` | `string` |  |
| `location_visibility` | `string` |  |
| `max_capacity` | `any` |  |
| `meeting_url` | `any` |  |
| `name` | `string` |  |
| `name_requirement` | `string` |  |
| `phone_number_requirement` | `any` |  |
| `platform` | `string` |  |
| `registration_open` | `bool` |  |
| `registration_question` | `[]any` |  |
| `reminders_disabled` | `bool` |  |
| `require_approval` | `bool` |  |
| `show_guest_list` | `bool` |  |
| `slug` | `string` |  |
| `spots_remaining` | `any` |  |
| `start_at` | `string` |  |
| `suppress_notification` | `bool` |  |
| `timezone` | `string` |  |
| `tint_color` | `string` |  |
| `url` | `string` |  |
| `user_id` | `string` |  |
| `visibility` | `string` |  |
| `waitlist_status` | `string` |  |

#### Example: Load

```go
event, err := client.Event(nil).Load(map[string]any{"id": "event_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(event) // the loaded record
```

#### Example: Create

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


### EventCancelRequest

Create an instance: `eventCancelRequest := client.EventCancelRequest(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `cancellation_token` | `string` |  |
| `event_id` | `string` |  |
| `guest_count` | `float64` |  |
| `is_paid` | `bool` |  |

#### Example: Create

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


### EventCoupon

Create an instance: `eventCoupon := client.EventCoupon(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |

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
| `remaining_count` | `int` |  |
| `valid_end_at` | `any` |  |
| `valid_start_at` | `any` |  |

#### Example: List

```go
eventCoupons, err := client.EventCoupon(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(eventCoupons) // the array of records
```

#### Example: Create

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


### EventTag

Create an instance: `eventTag := client.EventTag(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `color` | `any` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `tag_id` | `string` |  |

#### Example: List

```go
eventTags, err := client.EventTag(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(eventTags) // the array of records
```

#### Example: Create

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


### EventTagAssignment

Create an instance: `eventTagAssignment := client.EventTagAssignment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `applied_count` | `float64` |  |
| `event_id` | `[]any` |  |
| `skipped_count` | `float64` |  |
| `tag` | `string` |  |

#### Example: Create

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


### Guest

Create an instance: `guest := client.Guest(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approval_status` | `string` |  |
| `check_in_qr_code` | `string` |  |
| `eth_address` | `any` |  |
| `event_id` | `string` |  |
| `event_ticket` | `[]any` |  |
| `event_ticket_order` | `[]any` |  |
| `guest` | `[]any` |  |
| `guest_id` | `string` |  |
| `id` | `string` |  |
| `invited_at` | `any` |  |
| `joined_at` | `any` |  |
| `message` | `any` |  |
| `phone_number` | `int` |  |
| `registered_at` | `any` |  |
| `registration_answer` | `any` |  |
| `send_email` | `any` |  |
| `should_refund` | `bool` |  |
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

```go
guest, err := client.Guest(nil).Load(map[string]any{"id": "guest_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(guest) // the loaded record
```

#### Example: List

```go
guests, err := client.Guest(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(guests) // the array of records
```

#### Example: Create

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


### GuestInvite

Create an instance: `guestInvite := client.GuestInvite(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `event_id` | `string` |  |
| `guest` | `[]any` |  |
| `message` | `any` |  |

#### Example: Create

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


### GuestTicket

Create an instance: `guestTicket := client.GuestTicket(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `event_id` | `string` |  |
| `guest_id` | `string` |  |
| `send_email` | `any` |  |
| `ticket_ids_to_remove` | `[]any` |  |
| `tickets_to_add` | `[]any` |  |


### Host

Create an instance: `host := client.Host(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `access_level` | `any` |  |
| `email` | `string` |  |
| `event_id` | `string` |  |
| `is_visible` | `bool` |  |
| `name` | `string` |  |

#### Example: Create

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


### ImageUpload

Create an instance: `imageUpload := client.ImageUpload(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `content_type` | `any` |  |
| `file_url` | `string` |  |
| `upload_url` | `string` |  |

#### Example: Create

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


### Member

Create an instance: `member := client.Member(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `email` | `string` |  |
| `membership_id` | `string` |  |
| `membership_tier_id` | `string` |  |
| `registration_answer` | `[]any` |  |
| `skip_payment` | `bool` |  |
| `status` | `string` |  |
| `user_id` | `string` |  |

#### Example: Create

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


### MembershipTier

Create an instance: `membershipTier := client.MembershipTier(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `access_info` | `any` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `tint_color` | `string` |  |

#### Example: List

```go
membershipTiers, err := client.MembershipTier(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(membershipTiers) // the array of records
```


### OrganizationAdmin

Create an instance: `organizationAdmin := client.OrganizationAdmin(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

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

```go
organizationAdmins, err := client.OrganizationAdmin(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(organizationAdmins) // the array of records
```


### OrganizationCalendar

Create an instance: `organizationCalendar := client.OrganizationCalendar(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `avatar_url` | `any` |  |
| `coordinate` | `any` |  |
| `cover_image_url` | `any` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `instagram_handle` | `any` |  |
| `is_personal` | `bool` |  |
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

```go
organizationCalendars, err := client.OrganizationCalendar(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(organizationCalendars) // the array of records
```

#### Example: Create

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


### OrganizationEvent

Create an instance: `organizationEvent := client.OrganizationEvent(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

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
| `feedback_email` | `map[string]any` |  |
| `geo_address_json` | `any` |  |
| `geo_latitude` | `any` |  |
| `geo_longitude` | `any` |  |
| `id` | `string` |  |
| `location_type` | `string` |  |
| `location_visibility` | `string` |  |
| `managing_calendar` | `[]any` |  |
| `meeting_url` | `any` |  |
| `name` | `string` |  |
| `platform` | `string` |  |
| `registration_open` | `bool` |  |
| `registration_question` | `[]any` |  |
| `require_approval` | `bool` |  |
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

```go
organizationEvents, err := client.OrganizationEvent(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(organizationEvents) // the array of records
```


### OrganizationEventTransfer

Create an instance: `organizationEventTransfer := client.OrganizationEventTransfer(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `calendar_id` | `string` |  |
| `event_id` | `string` |  |

#### Example: Create

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


### TicketType

Create an instance: `ticketType := client.TicketType(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `cent` | `any` |  |
| `currency` | `any` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `is_flexible` | `bool` |  |
| `is_hidden` | `bool` |  |
| `max_capacity` | `any` |  |
| `min_cent` | `any` |  |
| `name` | `string` |  |
| `require_approval` | `bool` |  |
| `type` | `string` |  |
| `valid_end_at` | `any` |  |
| `valid_start_at` | `any` |  |

#### Example: Load

```go
ticketType, err := client.TicketType(nil).Load(map[string]any{"id": "ticket_type_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(ticketType) // the loaded record
```

#### Example: List

```go
ticketTypes, err := client.TicketType(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(ticketTypes) // the array of records
```

#### Example: Create

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


### User

Create an instance: `user := client.User(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

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

```go
user, err := client.User(nil).Load(map[string]any{"id": "user_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(user) // the loaded record
```


### Webhook

Create an instance: `webhook := client.Webhook(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `event_type` | `[]any` |  |
| `id` | `string` |  |
| `secret` | `string` |  |
| `status` | `string` |  |
| `url` | `string` |  |

#### Example: Load

```go
webhook, err := client.Webhook(nil).Load(map[string]any{"id": "webhook_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(webhook) // the loaded record
```

#### Example: List

```go
webhooks, err := client.Webhook(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(webhooks) // the array of records
```

#### Example: Create

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

Features are the extension mechanism. A feature implements the
`Feature` interface and provides hooks — functions keyed by pipeline
stage names.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as maps

The Go SDK uses `map[string]any` throughout rather than typed structs.
This mirrors the dynamic nature of the API and keeps the SDK
flexible — no code generation is needed when the API schema changes.

Use `core.ToMapAny()` to safely cast results and nested data.

### Package structure

```
github.com/voxgig-sdk/luma-sdk/go/
├── luma.go        # Root package — type aliases and constructors
├── core/               # SDK core — client, types, pipeline
├── entity/             # Entity implementations
├── feature/            # Built-in features (Base, Test, Log)
├── utility/            # Utility functions and struct library
└── test/               # Test suites
```

The root package (`github.com/voxgig-sdk/luma-sdk/go`) re-exports everything needed
for normal use. Import sub-packages only when you need specific types
like `core.ToMapAny`.

### Entity state

Entity instances are stateful. After a successful `List`, the entity
stores the returned data and match criteria internally.

```go
eventcoupon := client.EventCoupon(nil)
eventcoupon.List(nil, nil)

// eventcoupon.Data() now returns the eventcoupon data from the last list
// eventcoupon.Match() returns the last match criteria
```

Call `Make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`Direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `Prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
