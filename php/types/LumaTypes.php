<?php
declare(strict_types=1);

// Typed models for the Luma SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** Calendar entity data model. */
class Calendar
{
    public mixed $avatar_url;
    public string $calendar_id;
    public mixed $coordinate;
    public mixed $cover_image_url;
    public string $description;
    public string $id;
    public mixed $instagram_handle;
    public bool $is_personal;
    public mixed $location;
    public string $name;
    public string $slug;
    public mixed $social_image_url;
    public ?string $tint_color = null;
    public mixed $twitter_handle;
    public string $url;
    public mixed $website;
    public mixed $youtube_handle;
}

/** Request payload for Calendar#load. */
class CalendarLoadMatch
{
    public mixed $avatar_url = null;
    public ?string $calendar_id = null;
    public mixed $coordinate = null;
    public mixed $cover_image_url = null;
    public ?string $description = null;
    public string $id;
    public mixed $instagram_handle = null;
    public ?bool $is_personal = null;
    public mixed $location = null;
    public ?string $name = null;
    public ?string $slug = null;
    public mixed $social_image_url = null;
    public ?string $tint_color = null;
    public mixed $twitter_handle = null;
    public ?string $url = null;
    public mixed $website = null;
    public mixed $youtube_handle = null;
}

/** Request payload for Calendar#update. */
class CalendarUpdateData
{
    public mixed $avatar_url = null;
    public ?string $calendar_id = null;
    public mixed $coordinate = null;
    public mixed $cover_image_url = null;
    public ?string $description = null;
    public ?string $id = null;
    public mixed $instagram_handle = null;
    public ?bool $is_personal = null;
    public mixed $location = null;
    public ?string $name = null;
    public ?string $slug = null;
    public mixed $social_image_url = null;
    public ?string $tint_color = null;
    public mixed $twitter_handle = null;
    public ?string $url = null;
    public mixed $website = null;
    public mixed $youtube_handle = null;
}

/** CalendarAdmin entity data model. */
class CalendarAdmin
{
    public string $avatar_url;
    public string $email;
    public mixed $first_name;
    public string $id;
    public mixed $last_name;
    public string $name;
}

/** Request payload for CalendarAdmin#list. */
class CalendarAdminListMatch
{
    public ?string $avatar_url = null;
    public ?string $email = null;
    public mixed $first_name = null;
    public ?string $id = null;
    public mixed $last_name = null;
    public ?string $name = null;
}

/** CalendarCoupon entity data model. */
class CalendarCoupon
{
    public mixed $cents_off;
    public string $code;
    public mixed $currency;
    public mixed $discount;
    public ?string $event_ticket_type_id = null;
    public string $id;
    public mixed $percent_off;
    public int $remaining_count;
    public mixed $valid_end_at;
    public mixed $valid_start_at;
}

/** Request payload for CalendarCoupon#list. */
class CalendarCouponListMatch
{
    public mixed $cents_off = null;
    public ?string $code = null;
    public mixed $currency = null;
    public mixed $discount = null;
    public ?string $event_ticket_type_id = null;
    public ?string $id = null;
    public mixed $percent_off = null;
    public ?int $remaining_count = null;
    public mixed $valid_end_at = null;
    public mixed $valid_start_at = null;
}

/** Request payload for CalendarCoupon#create. */
class CalendarCouponCreateData
{
    public mixed $cents_off;
    public string $code;
    public mixed $currency;
    public mixed $discount;
    public ?string $event_ticket_type_id = null;
    public string $id;
    public mixed $percent_off;
    public int $remaining_count;
    public mixed $valid_end_at;
    public mixed $valid_start_at;
}

/** Request payload for CalendarCoupon#update. */
class CalendarCouponUpdateData
{
    public mixed $cents_off = null;
    public ?string $code = null;
    public mixed $currency = null;
    public mixed $discount = null;
    public ?string $event_ticket_type_id = null;
    public ?string $id = null;
    public mixed $percent_off = null;
    public ?int $remaining_count = null;
    public mixed $valid_end_at = null;
    public mixed $valid_start_at = null;
}

/** CalendarEvent entity data model. */
class CalendarEvent
{
    public string $id;
    public string $status;
    public mixed $submitted_by;
    public array $tag;
}

/** Request payload for CalendarEvent#load. */
class CalendarEventLoadMatch
{
    public string $id;
    public ?string $status = null;
    public mixed $submitted_by = null;
    public ?array $tag = null;
}

/** Request payload for CalendarEvent#list. */
class CalendarEventListMatch
{
    public ?string $id = null;
    public ?string $status = null;
    public mixed $submitted_by = null;
    public ?array $tag = null;
}

/** Request payload for CalendarEvent#create. */
class CalendarEventCreateData
{
    public string $id;
    public string $status;
    public mixed $submitted_by;
    public array $tag;
}

/** CalendarEventApproval entity data model. */
class CalendarEventApproval
{
    public string $calendar_event_id;
}

/** Request payload for CalendarEventApproval#create. */
class CalendarEventApprovalCreateData
{
    public string $calendar_event_id;
}

/** CalendarEventRejection entity data model. */
class CalendarEventRejection
{
    public string $calendar_event_id;
    public ?string $message = null;
}

/** Request payload for CalendarEventRejection#create. */
class CalendarEventRejectionCreateData
{
    public string $calendar_event_id;
    public ?string $message = null;
}

/** Contact entity data model. */
class Contact
{
    public string $avatar_url;
    public array $contact;
    public string $created_at;
    public string $email;
    public float $event_approved_count;
    public float $event_checked_in_count;
    public mixed $first_name;
    public string $id;
    public mixed $last_name;
    public mixed $membership;
    public string $name;
    public float $revenue_usd_cent;
    public mixed $tag = null;
    public string $user_id;
}

/** Request payload for Contact#list. */
class ContactListMatch
{
    public ?string $avatar_url = null;
    public ?array $contact = null;
    public ?string $created_at = null;
    public ?string $email = null;
    public ?float $event_approved_count = null;
    public ?float $event_checked_in_count = null;
    public mixed $first_name = null;
    public ?string $id = null;
    public mixed $last_name = null;
    public mixed $membership = null;
    public ?string $name = null;
    public ?float $revenue_usd_cent = null;
    public mixed $tag = null;
    public ?string $user_id = null;
}

/** Request payload for Contact#create. */
class ContactCreateData
{
    public string $avatar_url;
    public array $contact;
    public string $created_at;
    public string $email;
    public float $event_approved_count;
    public float $event_checked_in_count;
    public mixed $first_name;
    public string $id;
    public mixed $last_name;
    public mixed $membership;
    public string $name;
    public float $revenue_usd_cent;
    public mixed $tag = null;
    public string $user_id;
}

/** Request payload for Contact#remove. */
class ContactRemoveMatch
{
    public ?string $avatar_url = null;
    public ?array $contact = null;
    public ?string $created_at = null;
    public ?string $email = null;
    public ?float $event_approved_count = null;
    public ?float $event_checked_in_count = null;
    public mixed $first_name = null;
    public string $id;
    public mixed $last_name = null;
    public mixed $membership = null;
    public ?string $name = null;
    public ?float $revenue_usd_cent = null;
    public mixed $tag = null;
    public ?string $user_id = null;
}

/** ContactBlock entity data model. */
class ContactBlock
{
    public ?string $contact_id = null;
    public ?string $email = null;
}

/** Request payload for ContactBlock#create. */
class ContactBlockCreateData
{
    public ?string $contact_id = null;
    public ?string $email = null;
}

/** ContactRestore entity data model. */
class ContactRestore
{
    public ?string $contact_id = null;
    public ?string $email = null;
}

/** Request payload for ContactRestore#create. */
class ContactRestoreCreateData
{
    public ?string $contact_id = null;
    public ?string $email = null;
}

/** ContactTag entity data model. */
class ContactTag
{
    public mixed $color = null;
    public string $id;
    public string $name;
    public string $tag_id;
}

/** Request payload for ContactTag#list. */
class ContactTagListMatch
{
    public mixed $color = null;
    public ?string $id = null;
    public ?string $name = null;
    public ?string $tag_id = null;
}

/** Request payload for ContactTag#create. */
class ContactTagCreateData
{
    public mixed $color = null;
    public string $id;
    public string $name;
    public string $tag_id;
}

/** Request payload for ContactTag#update. */
class ContactTagUpdateData
{
    public mixed $color = null;
    public ?string $id = null;
    public ?string $name = null;
    public ?string $tag_id = null;
}

/** Request payload for ContactTag#remove. */
class ContactTagRemoveMatch
{
    public mixed $color = null;
    public string $id;
    public ?string $name = null;
    public ?string $tag_id = null;
}

/** ContactTagAssignment entity data model. */
class ContactTagAssignment
{
    public float $applied_count;
    public ?array $email = null;
    public float $skipped_count;
    public string $tag;
    public ?array $user_id = null;
}

/** Request payload for ContactTagAssignment#create. */
class ContactTagAssignmentCreateData
{
    public float $applied_count;
    public ?array $email = null;
    public float $skipped_count;
    public string $tag;
    public ?array $user_id = null;
}

/** Request payload for ContactTagAssignment#remove. */
class ContactTagAssignmentRemoveMatch
{
    public ?float $applied_count = null;
    public ?array $email = null;
    public ?float $skipped_count = null;
    public ?string $tag = null;
    public ?array $user_id = null;
}

/** EntityLookup entity data model. */
class EntityLookup
{
}

/** Request payload for EntityLookup#load. */
class EntityLookupLoadMatch
{
}

/** Event entity data model. */
class Event
{
    public string $access;
    public string $calendar_id;
    public ?bool $can_register_for_multiple_ticket = null;
    public mixed $coordinate;
    public string $cover_url;
    public string $created_at;
    public string $description;
    public string $description_md;
    public mixed $display_price;
    public string $duration_interval;
    public string $end_at;
    public string $event_id;
    public array $feedback_email;
    public mixed $geo_address_json;
    public array $guest_count;
    public array $host;
    public string $id;
    public string $location_type;
    public string $location_visibility;
    public mixed $max_capacity = null;
    public mixed $meeting_url;
    public string $name;
    public ?string $name_requirement = null;
    public mixed $phone_number_requirement = null;
    public string $platform;
    public bool $registration_open;
    public ?array $registration_question = null;
    public ?bool $reminders_disabled = null;
    public bool $require_approval;
    public ?bool $show_guest_list = null;
    public ?string $slug = null;
    public mixed $spots_remaining;
    public string $start_at;
    public ?bool $suppress_notification = null;
    public string $timezone;
    public ?string $tint_color = null;
    public string $url;
    public string $user_id;
    public string $visibility;
    public string $waitlist_status;
}

/** Request payload for Event#load. */
class EventLoadMatch
{
    public ?string $access = null;
    public ?string $calendar_id = null;
    public ?bool $can_register_for_multiple_ticket = null;
    public mixed $coordinate = null;
    public ?string $cover_url = null;
    public ?string $created_at = null;
    public ?string $description = null;
    public ?string $description_md = null;
    public mixed $display_price = null;
    public ?string $duration_interval = null;
    public ?string $end_at = null;
    public ?string $event_id = null;
    public ?array $feedback_email = null;
    public mixed $geo_address_json = null;
    public ?array $guest_count = null;
    public ?array $host = null;
    public string $id;
    public ?string $location_type = null;
    public ?string $location_visibility = null;
    public mixed $max_capacity = null;
    public mixed $meeting_url = null;
    public ?string $name = null;
    public ?string $name_requirement = null;
    public mixed $phone_number_requirement = null;
    public ?string $platform = null;
    public ?bool $registration_open = null;
    public ?array $registration_question = null;
    public ?bool $reminders_disabled = null;
    public ?bool $require_approval = null;
    public ?bool $show_guest_list = null;
    public ?string $slug = null;
    public mixed $spots_remaining = null;
    public ?string $start_at = null;
    public ?bool $suppress_notification = null;
    public ?string $timezone = null;
    public ?string $tint_color = null;
    public ?string $url = null;
    public ?string $user_id = null;
    public ?string $visibility = null;
    public ?string $waitlist_status = null;
}

/** Request payload for Event#create. */
class EventCreateData
{
    public string $access;
    public string $calendar_id;
    public ?bool $can_register_for_multiple_ticket = null;
    public mixed $coordinate;
    public string $cover_url;
    public string $created_at;
    public string $description;
    public string $description_md;
    public mixed $display_price;
    public string $duration_interval;
    public string $end_at;
    public string $event_id;
    public array $feedback_email;
    public mixed $geo_address_json;
    public array $guest_count;
    public array $host;
    public string $id;
    public string $location_type;
    public string $location_visibility;
    public mixed $max_capacity = null;
    public mixed $meeting_url;
    public string $name;
    public ?string $name_requirement = null;
    public mixed $phone_number_requirement = null;
    public string $platform;
    public bool $registration_open;
    public ?array $registration_question = null;
    public ?bool $reminders_disabled = null;
    public bool $require_approval;
    public ?bool $show_guest_list = null;
    public ?string $slug = null;
    public mixed $spots_remaining;
    public string $start_at;
    public ?bool $suppress_notification = null;
    public string $timezone;
    public ?string $tint_color = null;
    public string $url;
    public string $user_id;
    public string $visibility;
    public string $waitlist_status;
}

/** Request payload for Event#update. */
class EventUpdateData
{
    public ?string $access = null;
    public ?string $calendar_id = null;
    public ?bool $can_register_for_multiple_ticket = null;
    public mixed $coordinate = null;
    public ?string $cover_url = null;
    public ?string $created_at = null;
    public ?string $description = null;
    public ?string $description_md = null;
    public mixed $display_price = null;
    public ?string $duration_interval = null;
    public ?string $end_at = null;
    public ?string $event_id = null;
    public ?array $feedback_email = null;
    public mixed $geo_address_json = null;
    public ?array $guest_count = null;
    public ?array $host = null;
    public ?string $id = null;
    public ?string $location_type = null;
    public ?string $location_visibility = null;
    public mixed $max_capacity = null;
    public mixed $meeting_url = null;
    public ?string $name = null;
    public ?string $name_requirement = null;
    public mixed $phone_number_requirement = null;
    public ?string $platform = null;
    public ?bool $registration_open = null;
    public ?array $registration_question = null;
    public ?bool $reminders_disabled = null;
    public ?bool $require_approval = null;
    public ?bool $show_guest_list = null;
    public ?string $slug = null;
    public mixed $spots_remaining = null;
    public ?string $start_at = null;
    public ?bool $suppress_notification = null;
    public ?string $timezone = null;
    public ?string $tint_color = null;
    public ?string $url = null;
    public ?string $user_id = null;
    public ?string $visibility = null;
    public ?string $waitlist_status = null;
}

/** Request payload for Event#remove. */
class EventRemoveMatch
{
    public ?string $access = null;
    public ?string $calendar_id = null;
    public ?bool $can_register_for_multiple_ticket = null;
    public mixed $coordinate = null;
    public ?string $cover_url = null;
    public ?string $created_at = null;
    public ?string $description = null;
    public ?string $description_md = null;
    public mixed $display_price = null;
    public ?string $duration_interval = null;
    public ?string $end_at = null;
    public ?string $event_id = null;
    public ?array $feedback_email = null;
    public mixed $geo_address_json = null;
    public ?array $guest_count = null;
    public ?array $host = null;
    public string $id;
    public ?string $location_type = null;
    public ?string $location_visibility = null;
    public mixed $max_capacity = null;
    public mixed $meeting_url = null;
    public ?string $name = null;
    public ?string $name_requirement = null;
    public mixed $phone_number_requirement = null;
    public ?string $platform = null;
    public ?bool $registration_open = null;
    public ?array $registration_question = null;
    public ?bool $reminders_disabled = null;
    public ?bool $require_approval = null;
    public ?bool $show_guest_list = null;
    public ?string $slug = null;
    public mixed $spots_remaining = null;
    public ?string $start_at = null;
    public ?bool $suppress_notification = null;
    public ?string $timezone = null;
    public ?string $tint_color = null;
    public ?string $url = null;
    public ?string $user_id = null;
    public ?string $visibility = null;
    public ?string $waitlist_status = null;
}

/** EventCancelRequest entity data model. */
class EventCancelRequest
{
    public string $cancellation_token;
    public string $event_id;
    public float $guest_count;
    public bool $is_paid;
}

/** Request payload for EventCancelRequest#create. */
class EventCancelRequestCreateData
{
    public string $cancellation_token;
    public string $event_id;
    public float $guest_count;
    public bool $is_paid;
}

/** EventCoupon entity data model. */
class EventCoupon
{
    public mixed $cents_off;
    public string $code;
    public mixed $currency;
    public mixed $discount;
    public string $event_id;
    public ?string $event_ticket_type_id = null;
    public string $id;
    public mixed $percent_off;
    public int $remaining_count;
    public mixed $valid_end_at;
    public mixed $valid_start_at;
}

/** Request payload for EventCoupon#list. */
class EventCouponListMatch
{
    public mixed $cents_off = null;
    public ?string $code = null;
    public mixed $currency = null;
    public mixed $discount = null;
    public ?string $event_id = null;
    public ?string $event_ticket_type_id = null;
    public ?string $id = null;
    public mixed $percent_off = null;
    public ?int $remaining_count = null;
    public mixed $valid_end_at = null;
    public mixed $valid_start_at = null;
}

/** Request payload for EventCoupon#create. */
class EventCouponCreateData
{
    public mixed $cents_off;
    public string $code;
    public mixed $currency;
    public mixed $discount;
    public string $event_id;
    public ?string $event_ticket_type_id = null;
    public string $id;
    public mixed $percent_off;
    public int $remaining_count;
    public mixed $valid_end_at;
    public mixed $valid_start_at;
}

/** Request payload for EventCoupon#update. */
class EventCouponUpdateData
{
    public mixed $cents_off = null;
    public ?string $code = null;
    public mixed $currency = null;
    public mixed $discount = null;
    public ?string $event_id = null;
    public ?string $event_ticket_type_id = null;
    public ?string $id = null;
    public mixed $percent_off = null;
    public ?int $remaining_count = null;
    public mixed $valid_end_at = null;
    public mixed $valid_start_at = null;
}

/** EventTag entity data model. */
class EventTag
{
    public mixed $color = null;
    public string $id;
    public string $name;
    public string $tag_id;
}

/** Request payload for EventTag#list. */
class EventTagListMatch
{
    public mixed $color = null;
    public ?string $id = null;
    public ?string $name = null;
    public ?string $tag_id = null;
}

/** Request payload for EventTag#create. */
class EventTagCreateData
{
    public mixed $color = null;
    public string $id;
    public string $name;
    public string $tag_id;
}

/** Request payload for EventTag#update. */
class EventTagUpdateData
{
    public mixed $color = null;
    public ?string $id = null;
    public ?string $name = null;
    public ?string $tag_id = null;
}

/** Request payload for EventTag#remove. */
class EventTagRemoveMatch
{
    public mixed $color = null;
    public string $id;
    public ?string $name = null;
    public ?string $tag_id = null;
}

/** EventTagAssignment entity data model. */
class EventTagAssignment
{
    public float $applied_count;
    public array $event_id;
    public float $skipped_count;
    public string $tag;
}

/** Request payload for EventTagAssignment#create. */
class EventTagAssignmentCreateData
{
    public float $applied_count;
    public array $event_id;
    public float $skipped_count;
    public string $tag;
}

/** Request payload for EventTagAssignment#remove. */
class EventTagAssignmentRemoveMatch
{
    public ?float $applied_count = null;
    public ?array $event_id = null;
    public ?float $skipped_count = null;
    public ?string $tag = null;
}

/** Guest entity data model. */
class Guest
{
    public string $approval_status;
    public string $check_in_qr_code;
    public mixed $eth_address;
    public string $event_id;
    public array $event_ticket;
    public array $event_ticket_order;
    public array $guest;
    public string $guest_id;
    public string $id;
    public mixed $invited_at;
    public mixed $joined_at;
    public mixed $message = null;
    public int $phone_number;
    public mixed $registered_at;
    public mixed $registration_answer;
    public mixed $send_email = null;
    public ?bool $should_refund = null;
    public mixed $solana_address;
    public string $status;
    public mixed $ticket = null;
    public string $user_email;
    public mixed $user_first_name;
    public string $user_id;
    public mixed $user_last_name;
    public mixed $user_name;
    public mixed $utm_source;
}

/** Request payload for Guest#load. */
class GuestLoadMatch
{
    public ?string $approval_status = null;
    public ?string $check_in_qr_code = null;
    public mixed $eth_address = null;
    public ?string $event_id = null;
    public ?array $event_ticket = null;
    public ?array $event_ticket_order = null;
    public ?array $guest = null;
    public ?string $guest_id = null;
    public string $id;
    public mixed $invited_at = null;
    public mixed $joined_at = null;
    public mixed $message = null;
    public ?int $phone_number = null;
    public mixed $registered_at = null;
    public mixed $registration_answer = null;
    public mixed $send_email = null;
    public ?bool $should_refund = null;
    public mixed $solana_address = null;
    public ?string $status = null;
    public mixed $ticket = null;
    public ?string $user_email = null;
    public mixed $user_first_name = null;
    public ?string $user_id = null;
    public mixed $user_last_name = null;
    public mixed $user_name = null;
    public mixed $utm_source = null;
}

/** Request payload for Guest#list. */
class GuestListMatch
{
    public ?string $approval_status = null;
    public ?string $check_in_qr_code = null;
    public mixed $eth_address = null;
    public ?string $event_id = null;
    public ?array $event_ticket = null;
    public ?array $event_ticket_order = null;
    public ?array $guest = null;
    public ?string $guest_id = null;
    public ?string $id = null;
    public mixed $invited_at = null;
    public mixed $joined_at = null;
    public mixed $message = null;
    public ?int $phone_number = null;
    public mixed $registered_at = null;
    public mixed $registration_answer = null;
    public mixed $send_email = null;
    public ?bool $should_refund = null;
    public mixed $solana_address = null;
    public ?string $status = null;
    public mixed $ticket = null;
    public ?string $user_email = null;
    public mixed $user_first_name = null;
    public ?string $user_id = null;
    public mixed $user_last_name = null;
    public mixed $user_name = null;
    public mixed $utm_source = null;
}

/** Request payload for Guest#create. */
class GuestCreateData
{
    public string $approval_status;
    public string $check_in_qr_code;
    public mixed $eth_address;
    public string $event_id;
    public array $event_ticket;
    public array $event_ticket_order;
    public array $guest;
    public string $guest_id;
    public string $id;
    public mixed $invited_at;
    public mixed $joined_at;
    public mixed $message = null;
    public int $phone_number;
    public mixed $registered_at;
    public mixed $registration_answer;
    public mixed $send_email = null;
    public ?bool $should_refund = null;
    public mixed $solana_address;
    public string $status;
    public mixed $ticket = null;
    public string $user_email;
    public mixed $user_first_name;
    public string $user_id;
    public mixed $user_last_name;
    public mixed $user_name;
    public mixed $utm_source;
}

/** Request payload for Guest#update. */
class GuestUpdateData
{
    public ?string $approval_status = null;
    public ?string $check_in_qr_code = null;
    public mixed $eth_address = null;
    public ?string $event_id = null;
    public ?array $event_ticket = null;
    public ?array $event_ticket_order = null;
    public ?array $guest = null;
    public ?string $guest_id = null;
    public ?string $id = null;
    public mixed $invited_at = null;
    public mixed $joined_at = null;
    public mixed $message = null;
    public ?int $phone_number = null;
    public mixed $registered_at = null;
    public mixed $registration_answer = null;
    public mixed $send_email = null;
    public ?bool $should_refund = null;
    public mixed $solana_address = null;
    public ?string $status = null;
    public mixed $ticket = null;
    public ?string $user_email = null;
    public mixed $user_first_name = null;
    public ?string $user_id = null;
    public mixed $user_last_name = null;
    public mixed $user_name = null;
    public mixed $utm_source = null;
}

/** GuestInvite entity data model. */
class GuestInvite
{
    public string $event_id;
    public array $guest;
    public mixed $message = null;
}

/** Request payload for GuestInvite#create. */
class GuestInviteCreateData
{
    public string $event_id;
    public array $guest;
    public mixed $message = null;
}

/** GuestTicket entity data model. */
class GuestTicket
{
    public string $event_id;
    public string $guest_id;
    public mixed $send_email = null;
    public ?array $ticket_ids_to_remove = null;
    public ?array $tickets_to_add = null;
}

/** Request payload for GuestTicket#update. */
class GuestTicketUpdateData
{
    public ?string $event_id = null;
    public ?string $guest_id = null;
    public mixed $send_email = null;
    public ?array $ticket_ids_to_remove = null;
    public ?array $tickets_to_add = null;
}

/** Host entity data model. */
class Host
{
    public mixed $access_level = null;
    public string $email;
    public string $event_id;
    public ?bool $is_visible = null;
    public ?string $name = null;
}

/** Request payload for Host#create. */
class HostCreateData
{
    public mixed $access_level = null;
    public string $email;
    public string $event_id;
    public ?bool $is_visible = null;
    public ?string $name = null;
}

/** Request payload for Host#update. */
class HostUpdateData
{
    public mixed $access_level = null;
    public ?string $email = null;
    public ?string $event_id = null;
    public ?bool $is_visible = null;
    public ?string $name = null;
}

/** Request payload for Host#remove. */
class HostRemoveMatch
{
    public mixed $access_level = null;
    public ?string $email = null;
    public ?string $event_id = null;
    public ?bool $is_visible = null;
    public ?string $name = null;
}

/** ImageUpload entity data model. */
class ImageUpload
{
    public mixed $content_type = null;
    public string $file_url;
    public string $upload_url;
}

/** Request payload for ImageUpload#create. */
class ImageUploadCreateData
{
    public mixed $content_type = null;
    public string $file_url;
    public string $upload_url;
}

/** Member entity data model. */
class Member
{
    public string $email;
    public string $membership_id;
    public string $membership_tier_id;
    public ?array $registration_answer = null;
    public ?bool $skip_payment = null;
    public string $status;
    public string $user_id;
}

/** Request payload for Member#create. */
class MemberCreateData
{
    public string $email;
    public string $membership_id;
    public string $membership_tier_id;
    public ?array $registration_answer = null;
    public ?bool $skip_payment = null;
    public string $status;
    public string $user_id;
}

/** Request payload for Member#update. */
class MemberUpdateData
{
    public ?string $email = null;
    public ?string $membership_id = null;
    public ?string $membership_tier_id = null;
    public ?array $registration_answer = null;
    public ?bool $skip_payment = null;
    public ?string $status = null;
    public ?string $user_id = null;
}

/** MembershipTier entity data model. */
class MembershipTier
{
    public mixed $access_info;
    public string $description;
    public string $id;
    public string $name;
    public string $tint_color;
}

/** Request payload for MembershipTier#list. */
class MembershipTierListMatch
{
    public mixed $access_info = null;
    public ?string $description = null;
    public ?string $id = null;
    public ?string $name = null;
    public ?string $tint_color = null;
}

/** OrganizationAdmin entity data model. */
class OrganizationAdmin
{
    public string $api_id;
    public string $avatar_url;
    public string $email;
    public mixed $first_name;
    public string $id;
    public mixed $last_name;
    public string $name;
}

/** Request payload for OrganizationAdmin#list. */
class OrganizationAdminListMatch
{
    public ?string $api_id = null;
    public ?string $avatar_url = null;
    public ?string $email = null;
    public mixed $first_name = null;
    public ?string $id = null;
    public mixed $last_name = null;
    public ?string $name = null;
}

/** OrganizationCalendar entity data model. */
class OrganizationCalendar
{
    public mixed $avatar_url;
    public mixed $coordinate;
    public mixed $cover_image_url;
    public string $description;
    public string $id;
    public mixed $instagram_handle;
    public bool $is_personal;
    public mixed $location;
    public string $name;
    public string $slug;
    public mixed $social_image_url;
    public ?string $tint_color = null;
    public mixed $twitter_handle;
    public string $url;
    public mixed $website;
    public mixed $youtube_handle;
}

/** Request payload for OrganizationCalendar#list. */
class OrganizationCalendarListMatch
{
    public mixed $avatar_url = null;
    public mixed $coordinate = null;
    public mixed $cover_image_url = null;
    public ?string $description = null;
    public ?string $id = null;
    public mixed $instagram_handle = null;
    public ?bool $is_personal = null;
    public mixed $location = null;
    public ?string $name = null;
    public ?string $slug = null;
    public mixed $social_image_url = null;
    public ?string $tint_color = null;
    public mixed $twitter_handle = null;
    public ?string $url = null;
    public mixed $website = null;
    public mixed $youtube_handle = null;
}

/** Request payload for OrganizationCalendar#create. */
class OrganizationCalendarCreateData
{
    public mixed $avatar_url;
    public mixed $coordinate;
    public mixed $cover_image_url;
    public string $description;
    public string $id;
    public mixed $instagram_handle;
    public bool $is_personal;
    public mixed $location;
    public string $name;
    public string $slug;
    public mixed $social_image_url;
    public ?string $tint_color = null;
    public mixed $twitter_handle;
    public string $url;
    public mixed $website;
    public mixed $youtube_handle;
}

/** OrganizationEvent entity data model. */
class OrganizationEvent
{
    public string $api_id;
    public string $calendar_api_id;
    public string $calendar_id;
    public mixed $coordinate;
    public string $cover_url;
    public string $created_at;
    public mixed $display_price;
    public string $duration_interval;
    public string $end_at;
    public array $feedback_email;
    public mixed $geo_address_json;
    public mixed $geo_latitude;
    public mixed $geo_longitude;
    public string $id;
    public string $location_type;
    public string $location_visibility;
    public array $managing_calendar;
    public mixed $meeting_url;
    public string $name;
    public string $platform;
    public bool $registration_open;
    public ?array $registration_question = null;
    public bool $require_approval;
    public mixed $spots_remaining;
    public string $start_at;
    public string $timezone;
    public string $url;
    public string $user_api_id;
    public string $user_id;
    public string $visibility;
    public string $waitlist_status;
    public mixed $zoom_meeting_url;
}

/** Request payload for OrganizationEvent#list. */
class OrganizationEventListMatch
{
    public ?string $api_id = null;
    public ?string $calendar_api_id = null;
    public ?string $calendar_id = null;
    public mixed $coordinate = null;
    public ?string $cover_url = null;
    public ?string $created_at = null;
    public mixed $display_price = null;
    public ?string $duration_interval = null;
    public ?string $end_at = null;
    public ?array $feedback_email = null;
    public mixed $geo_address_json = null;
    public mixed $geo_latitude = null;
    public mixed $geo_longitude = null;
    public ?string $id = null;
    public ?string $location_type = null;
    public ?string $location_visibility = null;
    public ?array $managing_calendar = null;
    public mixed $meeting_url = null;
    public ?string $name = null;
    public ?string $platform = null;
    public ?bool $registration_open = null;
    public ?array $registration_question = null;
    public ?bool $require_approval = null;
    public mixed $spots_remaining = null;
    public ?string $start_at = null;
    public ?string $timezone = null;
    public ?string $url = null;
    public ?string $user_api_id = null;
    public ?string $user_id = null;
    public ?string $visibility = null;
    public ?string $waitlist_status = null;
    public mixed $zoom_meeting_url = null;
}

/** OrganizationEventTransfer entity data model. */
class OrganizationEventTransfer
{
    public string $calendar_id;
    public string $event_id;
}

/** Request payload for OrganizationEventTransfer#create. */
class OrganizationEventTransferCreateData
{
    public string $calendar_id;
    public string $event_id;
}

/** TicketType entity data model. */
class TicketType
{
    public mixed $cent = null;
    public mixed $currency = null;
    public ?string $description = null;
    public string $id;
    public ?bool $is_flexible = null;
    public ?bool $is_hidden = null;
    public mixed $max_capacity = null;
    public mixed $min_cent = null;
    public string $name;
    public ?bool $require_approval = null;
    public string $type;
    public mixed $valid_end_at = null;
    public mixed $valid_start_at = null;
}

/** Request payload for TicketType#load. */
class TicketTypeLoadMatch
{
    public mixed $cent = null;
    public mixed $currency = null;
    public ?string $description = null;
    public string $id;
    public ?bool $is_flexible = null;
    public ?bool $is_hidden = null;
    public mixed $max_capacity = null;
    public mixed $min_cent = null;
    public ?string $name = null;
    public ?bool $require_approval = null;
    public ?string $type = null;
    public mixed $valid_end_at = null;
    public mixed $valid_start_at = null;
}

/** Request payload for TicketType#list. */
class TicketTypeListMatch
{
    public mixed $cent = null;
    public mixed $currency = null;
    public ?string $description = null;
    public ?string $id = null;
    public ?bool $is_flexible = null;
    public ?bool $is_hidden = null;
    public mixed $max_capacity = null;
    public mixed $min_cent = null;
    public ?string $name = null;
    public ?bool $require_approval = null;
    public ?string $type = null;
    public mixed $valid_end_at = null;
    public mixed $valid_start_at = null;
}

/** Request payload for TicketType#create. */
class TicketTypeCreateData
{
    public mixed $cent = null;
    public mixed $currency = null;
    public ?string $description = null;
    public string $id;
    public ?bool $is_flexible = null;
    public ?bool $is_hidden = null;
    public mixed $max_capacity = null;
    public mixed $min_cent = null;
    public string $name;
    public ?bool $require_approval = null;
    public string $type;
    public mixed $valid_end_at = null;
    public mixed $valid_start_at = null;
}

/** Request payload for TicketType#update. */
class TicketTypeUpdateData
{
    public mixed $cent = null;
    public mixed $currency = null;
    public ?string $description = null;
    public ?string $id = null;
    public ?bool $is_flexible = null;
    public ?bool $is_hidden = null;
    public mixed $max_capacity = null;
    public mixed $min_cent = null;
    public ?string $name = null;
    public ?bool $require_approval = null;
    public ?string $type = null;
    public mixed $valid_end_at = null;
    public mixed $valid_start_at = null;
}

/** Request payload for TicketType#remove. */
class TicketTypeRemoveMatch
{
    public mixed $cent = null;
    public mixed $currency = null;
    public ?string $description = null;
    public string $id;
    public ?bool $is_flexible = null;
    public ?bool $is_hidden = null;
    public mixed $max_capacity = null;
    public mixed $min_cent = null;
    public ?string $name = null;
    public ?bool $require_approval = null;
    public ?string $type = null;
    public mixed $valid_end_at = null;
    public mixed $valid_start_at = null;
}

/** User entity data model. */
class User
{
    public string $avatar_url;
    public string $email;
    public mixed $first_name;
    public string $id;
    public mixed $last_name;
    public string $name;
}

/** Request payload for User#load. */
class UserLoadMatch
{
    public ?string $avatar_url = null;
    public ?string $email = null;
    public mixed $first_name = null;
    public string $id;
    public mixed $last_name = null;
    public ?string $name = null;
}

/** Webhook entity data model. */
class Webhook
{
    public string $created_at;
    public array $event_type;
    public string $id;
    public string $secret;
    public string $status;
    public string $url;
}

/** Request payload for Webhook#load. */
class WebhookLoadMatch
{
    public ?string $created_at = null;
    public ?array $event_type = null;
    public string $id;
    public ?string $secret = null;
    public ?string $status = null;
    public ?string $url = null;
}

/** Request payload for Webhook#list. */
class WebhookListMatch
{
    public ?string $created_at = null;
    public ?array $event_type = null;
    public ?string $id = null;
    public ?string $secret = null;
    public ?string $status = null;
    public ?string $url = null;
}

/** Request payload for Webhook#create. */
class WebhookCreateData
{
    public string $created_at;
    public array $event_type;
    public string $id;
    public string $secret;
    public string $status;
    public string $url;
}

/** Request payload for Webhook#update. */
class WebhookUpdateData
{
    public ?string $created_at = null;
    public ?array $event_type = null;
    public ?string $id = null;
    public ?string $secret = null;
    public ?string $status = null;
    public ?string $url = null;
}

/** Request payload for Webhook#remove. */
class WebhookRemoveMatch
{
    public ?string $created_at = null;
    public ?array $event_type = null;
    public string $id;
    public ?string $secret = null;
    public ?string $status = null;
    public ?string $url = null;
}

