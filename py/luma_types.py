# Typed models for the Luma SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class CalendarRequired(TypedDict):
    avatar_url: Any
    calendar_id: str
    coordinate: Any
    cover_image_url: Any
    description: str
    id: str
    instagram_handle: Any
    is_personal: bool
    location: Any
    name: str
    slug: str
    social_image_url: Any
    twitter_handle: Any
    url: str
    website: Any
    youtube_handle: Any


class Calendar(CalendarRequired, total=False):
    tint_color: str


class CalendarLoadMatchRequired(TypedDict):
    id: str


class CalendarLoadMatch(CalendarLoadMatchRequired, total=False):
    avatar_url: Any
    calendar_id: str
    coordinate: Any
    cover_image_url: Any
    description: str
    instagram_handle: Any
    is_personal: bool
    location: Any
    name: str
    slug: str
    social_image_url: Any
    tint_color: str
    twitter_handle: Any
    url: str
    website: Any
    youtube_handle: Any


class CalendarUpdateData(TypedDict, total=False):
    avatar_url: Any
    calendar_id: str
    coordinate: Any
    cover_image_url: Any
    description: str
    id: str
    instagram_handle: Any
    is_personal: bool
    location: Any
    name: str
    slug: str
    social_image_url: Any
    tint_color: str
    twitter_handle: Any
    url: str
    website: Any
    youtube_handle: Any


class CalendarAdmin(TypedDict):
    avatar_url: str
    email: str
    first_name: Any
    id: str
    last_name: Any
    name: str


class CalendarAdminListMatch(TypedDict, total=False):
    avatar_url: str
    email: str
    first_name: Any
    id: str
    last_name: Any
    name: str


class CalendarCouponRequired(TypedDict):
    cents_off: Any
    code: str
    currency: Any
    discount: Any
    id: str
    percent_off: Any
    remaining_count: int
    valid_end_at: Any
    valid_start_at: Any


class CalendarCoupon(CalendarCouponRequired, total=False):
    event_ticket_type_id: str


class CalendarCouponListMatch(TypedDict, total=False):
    cents_off: Any
    code: str
    currency: Any
    discount: Any
    event_ticket_type_id: str
    id: str
    percent_off: Any
    remaining_count: int
    valid_end_at: Any
    valid_start_at: Any


class CalendarCouponCreateDataRequired(TypedDict):
    cents_off: Any
    code: str
    currency: Any
    discount: Any
    id: str
    percent_off: Any
    remaining_count: int
    valid_end_at: Any
    valid_start_at: Any


class CalendarCouponCreateData(CalendarCouponCreateDataRequired, total=False):
    event_ticket_type_id: str


class CalendarCouponUpdateData(TypedDict, total=False):
    cents_off: Any
    code: str
    currency: Any
    discount: Any
    event_ticket_type_id: str
    id: str
    percent_off: Any
    remaining_count: int
    valid_end_at: Any
    valid_start_at: Any


class CalendarEvent(TypedDict):
    id: str
    status: str
    submitted_by: Any
    tag: list


class CalendarEventLoadMatchRequired(TypedDict):
    id: str


class CalendarEventLoadMatch(CalendarEventLoadMatchRequired, total=False):
    status: str
    submitted_by: Any
    tag: list


class CalendarEventListMatch(TypedDict, total=False):
    id: str
    status: str
    submitted_by: Any
    tag: list


class CalendarEventCreateData(TypedDict):
    id: str
    status: str
    submitted_by: Any
    tag: list


class CalendarEventApproval(TypedDict):
    calendar_event_id: str


class CalendarEventApprovalCreateData(TypedDict):
    calendar_event_id: str


class CalendarEventRejectionRequired(TypedDict):
    calendar_event_id: str


class CalendarEventRejection(CalendarEventRejectionRequired, total=False):
    message: str


class CalendarEventRejectionCreateDataRequired(TypedDict):
    calendar_event_id: str


class CalendarEventRejectionCreateData(CalendarEventRejectionCreateDataRequired, total=False):
    message: str


class ContactRequired(TypedDict):
    avatar_url: str
    contact: list
    created_at: str
    email: str
    event_approved_count: float
    event_checked_in_count: float
    first_name: Any
    id: str
    last_name: Any
    membership: Any
    name: str
    revenue_usd_cent: float
    user_id: str


class Contact(ContactRequired, total=False):
    tag: Any


class ContactListMatch(TypedDict, total=False):
    avatar_url: str
    contact: list
    created_at: str
    email: str
    event_approved_count: float
    event_checked_in_count: float
    first_name: Any
    id: str
    last_name: Any
    membership: Any
    name: str
    revenue_usd_cent: float
    tag: Any
    user_id: str


class ContactCreateDataRequired(TypedDict):
    avatar_url: str
    contact: list
    created_at: str
    email: str
    event_approved_count: float
    event_checked_in_count: float
    first_name: Any
    id: str
    last_name: Any
    membership: Any
    name: str
    revenue_usd_cent: float
    user_id: str


class ContactCreateData(ContactCreateDataRequired, total=False):
    tag: Any


class ContactRemoveMatchRequired(TypedDict):
    id: str


class ContactRemoveMatch(ContactRemoveMatchRequired, total=False):
    avatar_url: str
    contact: list
    created_at: str
    email: str
    event_approved_count: float
    event_checked_in_count: float
    first_name: Any
    last_name: Any
    membership: Any
    name: str
    revenue_usd_cent: float
    tag: Any
    user_id: str


class ContactBlock(TypedDict, total=False):
    contact_id: str
    email: str


class ContactBlockCreateData(TypedDict, total=False):
    contact_id: str
    email: str


class ContactRestore(TypedDict, total=False):
    contact_id: str
    email: str


class ContactRestoreCreateData(TypedDict, total=False):
    contact_id: str
    email: str


class ContactTagRequired(TypedDict):
    id: str
    name: str
    tag_id: str


class ContactTag(ContactTagRequired, total=False):
    color: Any


class ContactTagListMatch(TypedDict, total=False):
    color: Any
    id: str
    name: str
    tag_id: str


class ContactTagCreateDataRequired(TypedDict):
    id: str
    name: str
    tag_id: str


class ContactTagCreateData(ContactTagCreateDataRequired, total=False):
    color: Any


class ContactTagUpdateData(TypedDict, total=False):
    color: Any
    id: str
    name: str
    tag_id: str


class ContactTagRemoveMatchRequired(TypedDict):
    id: str


class ContactTagRemoveMatch(ContactTagRemoveMatchRequired, total=False):
    color: Any
    name: str
    tag_id: str


class ContactTagAssignmentRequired(TypedDict):
    applied_count: float
    skipped_count: float
    tag: str


class ContactTagAssignment(ContactTagAssignmentRequired, total=False):
    email: list
    user_id: list


class ContactTagAssignmentCreateDataRequired(TypedDict):
    applied_count: float
    skipped_count: float
    tag: str


class ContactTagAssignmentCreateData(ContactTagAssignmentCreateDataRequired, total=False):
    email: list
    user_id: list


class ContactTagAssignmentRemoveMatch(TypedDict, total=False):
    applied_count: float
    email: list
    skipped_count: float
    tag: str
    user_id: list


class EntityLookup(TypedDict):
    pass


class EntityLookupLoadMatch(TypedDict):
    pass


class EventRequired(TypedDict):
    access: str
    calendar_id: str
    coordinate: Any
    cover_url: str
    created_at: str
    description: str
    description_md: str
    display_price: Any
    duration_interval: str
    end_at: str
    event_id: str
    feedback_email: dict
    geo_address_json: Any
    guest_count: dict
    host: list
    id: str
    location_type: str
    location_visibility: str
    meeting_url: Any
    name: str
    platform: str
    registration_open: bool
    require_approval: bool
    spots_remaining: Any
    start_at: str
    timezone: str
    url: str
    user_id: str
    visibility: str
    waitlist_status: str


class Event(EventRequired, total=False):
    can_register_for_multiple_ticket: bool
    max_capacity: Any
    name_requirement: str
    phone_number_requirement: Any
    registration_question: list
    reminders_disabled: bool
    show_guest_list: bool
    slug: str
    suppress_notification: bool
    tint_color: str


class EventLoadMatchRequired(TypedDict):
    id: str


class EventLoadMatch(EventLoadMatchRequired, total=False):
    access: str
    calendar_id: str
    can_register_for_multiple_ticket: bool
    coordinate: Any
    cover_url: str
    created_at: str
    description: str
    description_md: str
    display_price: Any
    duration_interval: str
    end_at: str
    event_id: str
    feedback_email: dict
    geo_address_json: Any
    guest_count: dict
    host: list
    location_type: str
    location_visibility: str
    max_capacity: Any
    meeting_url: Any
    name: str
    name_requirement: str
    phone_number_requirement: Any
    platform: str
    registration_open: bool
    registration_question: list
    reminders_disabled: bool
    require_approval: bool
    show_guest_list: bool
    slug: str
    spots_remaining: Any
    start_at: str
    suppress_notification: bool
    timezone: str
    tint_color: str
    url: str
    user_id: str
    visibility: str
    waitlist_status: str


class EventCreateDataRequired(TypedDict):
    access: str
    calendar_id: str
    coordinate: Any
    cover_url: str
    created_at: str
    description: str
    description_md: str
    display_price: Any
    duration_interval: str
    end_at: str
    event_id: str
    feedback_email: dict
    geo_address_json: Any
    guest_count: dict
    host: list
    id: str
    location_type: str
    location_visibility: str
    meeting_url: Any
    name: str
    platform: str
    registration_open: bool
    require_approval: bool
    spots_remaining: Any
    start_at: str
    timezone: str
    url: str
    user_id: str
    visibility: str
    waitlist_status: str


class EventCreateData(EventCreateDataRequired, total=False):
    can_register_for_multiple_ticket: bool
    max_capacity: Any
    name_requirement: str
    phone_number_requirement: Any
    registration_question: list
    reminders_disabled: bool
    show_guest_list: bool
    slug: str
    suppress_notification: bool
    tint_color: str


class EventUpdateData(TypedDict, total=False):
    access: str
    calendar_id: str
    can_register_for_multiple_ticket: bool
    coordinate: Any
    cover_url: str
    created_at: str
    description: str
    description_md: str
    display_price: Any
    duration_interval: str
    end_at: str
    event_id: str
    feedback_email: dict
    geo_address_json: Any
    guest_count: dict
    host: list
    id: str
    location_type: str
    location_visibility: str
    max_capacity: Any
    meeting_url: Any
    name: str
    name_requirement: str
    phone_number_requirement: Any
    platform: str
    registration_open: bool
    registration_question: list
    reminders_disabled: bool
    require_approval: bool
    show_guest_list: bool
    slug: str
    spots_remaining: Any
    start_at: str
    suppress_notification: bool
    timezone: str
    tint_color: str
    url: str
    user_id: str
    visibility: str
    waitlist_status: str


class EventRemoveMatchRequired(TypedDict):
    id: str


class EventRemoveMatch(EventRemoveMatchRequired, total=False):
    access: str
    calendar_id: str
    can_register_for_multiple_ticket: bool
    coordinate: Any
    cover_url: str
    created_at: str
    description: str
    description_md: str
    display_price: Any
    duration_interval: str
    end_at: str
    event_id: str
    feedback_email: dict
    geo_address_json: Any
    guest_count: dict
    host: list
    location_type: str
    location_visibility: str
    max_capacity: Any
    meeting_url: Any
    name: str
    name_requirement: str
    phone_number_requirement: Any
    platform: str
    registration_open: bool
    registration_question: list
    reminders_disabled: bool
    require_approval: bool
    show_guest_list: bool
    slug: str
    spots_remaining: Any
    start_at: str
    suppress_notification: bool
    timezone: str
    tint_color: str
    url: str
    user_id: str
    visibility: str
    waitlist_status: str


class EventCancelRequest(TypedDict):
    cancellation_token: str
    event_id: str
    guest_count: float
    is_paid: bool


class EventCancelRequestCreateData(TypedDict):
    cancellation_token: str
    event_id: str
    guest_count: float
    is_paid: bool


class EventCouponRequired(TypedDict):
    cents_off: Any
    code: str
    currency: Any
    discount: Any
    event_id: str
    id: str
    percent_off: Any
    remaining_count: int
    valid_end_at: Any
    valid_start_at: Any


class EventCoupon(EventCouponRequired, total=False):
    event_ticket_type_id: str


class EventCouponListMatch(TypedDict, total=False):
    cents_off: Any
    code: str
    currency: Any
    discount: Any
    event_id: str
    event_ticket_type_id: str
    id: str
    percent_off: Any
    remaining_count: int
    valid_end_at: Any
    valid_start_at: Any


class EventCouponCreateDataRequired(TypedDict):
    cents_off: Any
    code: str
    currency: Any
    discount: Any
    event_id: str
    id: str
    percent_off: Any
    remaining_count: int
    valid_end_at: Any
    valid_start_at: Any


class EventCouponCreateData(EventCouponCreateDataRequired, total=False):
    event_ticket_type_id: str


class EventCouponUpdateData(TypedDict, total=False):
    cents_off: Any
    code: str
    currency: Any
    discount: Any
    event_id: str
    event_ticket_type_id: str
    id: str
    percent_off: Any
    remaining_count: int
    valid_end_at: Any
    valid_start_at: Any


class EventTagRequired(TypedDict):
    id: str
    name: str
    tag_id: str


class EventTag(EventTagRequired, total=False):
    color: Any


class EventTagListMatch(TypedDict, total=False):
    color: Any
    id: str
    name: str
    tag_id: str


class EventTagCreateDataRequired(TypedDict):
    id: str
    name: str
    tag_id: str


class EventTagCreateData(EventTagCreateDataRequired, total=False):
    color: Any


class EventTagUpdateData(TypedDict, total=False):
    color: Any
    id: str
    name: str
    tag_id: str


class EventTagRemoveMatchRequired(TypedDict):
    id: str


class EventTagRemoveMatch(EventTagRemoveMatchRequired, total=False):
    color: Any
    name: str
    tag_id: str


class EventTagAssignment(TypedDict):
    applied_count: float
    event_id: list
    skipped_count: float
    tag: str


class EventTagAssignmentCreateData(TypedDict):
    applied_count: float
    event_id: list
    skipped_count: float
    tag: str


class EventTagAssignmentRemoveMatch(TypedDict, total=False):
    applied_count: float
    event_id: list
    skipped_count: float
    tag: str


class GuestRequired(TypedDict):
    approval_status: str
    check_in_qr_code: str
    eth_address: Any
    event_id: str
    event_ticket: list
    event_ticket_order: list
    guest: list
    guest_id: str
    id: str
    invited_at: Any
    joined_at: Any
    phone_number: int
    registered_at: Any
    registration_answer: Any
    solana_address: Any
    status: str
    user_email: str
    user_first_name: Any
    user_id: str
    user_last_name: Any
    user_name: Any
    utm_source: Any


class Guest(GuestRequired, total=False):
    message: Any
    send_email: Any
    should_refund: bool
    ticket: Any


class GuestLoadMatchRequired(TypedDict):
    id: str


class GuestLoadMatch(GuestLoadMatchRequired, total=False):
    approval_status: str
    check_in_qr_code: str
    eth_address: Any
    event_id: str
    event_ticket: list
    event_ticket_order: list
    guest: list
    guest_id: str
    invited_at: Any
    joined_at: Any
    message: Any
    phone_number: int
    registered_at: Any
    registration_answer: Any
    send_email: Any
    should_refund: bool
    solana_address: Any
    status: str
    ticket: Any
    user_email: str
    user_first_name: Any
    user_id: str
    user_last_name: Any
    user_name: Any
    utm_source: Any


class GuestListMatch(TypedDict, total=False):
    approval_status: str
    check_in_qr_code: str
    eth_address: Any
    event_id: str
    event_ticket: list
    event_ticket_order: list
    guest: list
    guest_id: str
    id: str
    invited_at: Any
    joined_at: Any
    message: Any
    phone_number: int
    registered_at: Any
    registration_answer: Any
    send_email: Any
    should_refund: bool
    solana_address: Any
    status: str
    ticket: Any
    user_email: str
    user_first_name: Any
    user_id: str
    user_last_name: Any
    user_name: Any
    utm_source: Any


class GuestCreateDataRequired(TypedDict):
    approval_status: str
    check_in_qr_code: str
    eth_address: Any
    event_id: str
    event_ticket: list
    event_ticket_order: list
    guest: list
    guest_id: str
    id: str
    invited_at: Any
    joined_at: Any
    phone_number: int
    registered_at: Any
    registration_answer: Any
    solana_address: Any
    status: str
    user_email: str
    user_first_name: Any
    user_id: str
    user_last_name: Any
    user_name: Any
    utm_source: Any


class GuestCreateData(GuestCreateDataRequired, total=False):
    message: Any
    send_email: Any
    should_refund: bool
    ticket: Any


class GuestUpdateData(TypedDict, total=False):
    approval_status: str
    check_in_qr_code: str
    eth_address: Any
    event_id: str
    event_ticket: list
    event_ticket_order: list
    guest: list
    guest_id: str
    id: str
    invited_at: Any
    joined_at: Any
    message: Any
    phone_number: int
    registered_at: Any
    registration_answer: Any
    send_email: Any
    should_refund: bool
    solana_address: Any
    status: str
    ticket: Any
    user_email: str
    user_first_name: Any
    user_id: str
    user_last_name: Any
    user_name: Any
    utm_source: Any


class GuestInviteRequired(TypedDict):
    event_id: str
    guest: list


class GuestInvite(GuestInviteRequired, total=False):
    message: Any


class GuestInviteCreateDataRequired(TypedDict):
    event_id: str
    guest: list


class GuestInviteCreateData(GuestInviteCreateDataRequired, total=False):
    message: Any


class GuestTicketRequired(TypedDict):
    event_id: str
    guest_id: str


class GuestTicket(GuestTicketRequired, total=False):
    send_email: Any
    ticket_ids_to_remove: list
    tickets_to_add: list


class GuestTicketUpdateData(TypedDict, total=False):
    event_id: str
    guest_id: str
    send_email: Any
    ticket_ids_to_remove: list
    tickets_to_add: list


class HostRequired(TypedDict):
    email: str
    event_id: str


class Host(HostRequired, total=False):
    access_level: Any
    is_visible: bool
    name: str


class HostCreateDataRequired(TypedDict):
    email: str
    event_id: str


class HostCreateData(HostCreateDataRequired, total=False):
    access_level: Any
    is_visible: bool
    name: str


class HostUpdateData(TypedDict, total=False):
    access_level: Any
    email: str
    event_id: str
    is_visible: bool
    name: str


class HostRemoveMatch(TypedDict, total=False):
    access_level: Any
    email: str
    event_id: str
    is_visible: bool
    name: str


class ImageUploadRequired(TypedDict):
    file_url: str
    upload_url: str


class ImageUpload(ImageUploadRequired, total=False):
    content_type: Any


class ImageUploadCreateDataRequired(TypedDict):
    file_url: str
    upload_url: str


class ImageUploadCreateData(ImageUploadCreateDataRequired, total=False):
    content_type: Any


class MemberRequired(TypedDict):
    email: str
    membership_id: str
    membership_tier_id: str
    status: str
    user_id: str


class Member(MemberRequired, total=False):
    registration_answer: list
    skip_payment: bool


class MemberCreateDataRequired(TypedDict):
    email: str
    membership_id: str
    membership_tier_id: str
    status: str
    user_id: str


class MemberCreateData(MemberCreateDataRequired, total=False):
    registration_answer: list
    skip_payment: bool


class MemberUpdateData(TypedDict, total=False):
    email: str
    membership_id: str
    membership_tier_id: str
    registration_answer: list
    skip_payment: bool
    status: str
    user_id: str


class MembershipTier(TypedDict):
    access_info: Any
    description: str
    id: str
    name: str
    tint_color: str


class MembershipTierListMatch(TypedDict, total=False):
    access_info: Any
    description: str
    id: str
    name: str
    tint_color: str


class OrganizationAdmin(TypedDict):
    api_id: str
    avatar_url: str
    email: str
    first_name: Any
    id: str
    last_name: Any
    name: str


class OrganizationAdminListMatch(TypedDict, total=False):
    api_id: str
    avatar_url: str
    email: str
    first_name: Any
    id: str
    last_name: Any
    name: str


class OrganizationCalendarRequired(TypedDict):
    avatar_url: Any
    coordinate: Any
    cover_image_url: Any
    description: str
    id: str
    instagram_handle: Any
    is_personal: bool
    location: Any
    name: str
    slug: str
    social_image_url: Any
    twitter_handle: Any
    url: str
    website: Any
    youtube_handle: Any


class OrganizationCalendar(OrganizationCalendarRequired, total=False):
    tint_color: str


class OrganizationCalendarListMatch(TypedDict, total=False):
    avatar_url: Any
    coordinate: Any
    cover_image_url: Any
    description: str
    id: str
    instagram_handle: Any
    is_personal: bool
    location: Any
    name: str
    slug: str
    social_image_url: Any
    tint_color: str
    twitter_handle: Any
    url: str
    website: Any
    youtube_handle: Any


class OrganizationCalendarCreateDataRequired(TypedDict):
    avatar_url: Any
    coordinate: Any
    cover_image_url: Any
    description: str
    id: str
    instagram_handle: Any
    is_personal: bool
    location: Any
    name: str
    slug: str
    social_image_url: Any
    twitter_handle: Any
    url: str
    website: Any
    youtube_handle: Any


class OrganizationCalendarCreateData(OrganizationCalendarCreateDataRequired, total=False):
    tint_color: str


class OrganizationEventRequired(TypedDict):
    api_id: str
    calendar_api_id: str
    calendar_id: str
    coordinate: Any
    cover_url: str
    created_at: str
    display_price: Any
    duration_interval: str
    end_at: str
    feedback_email: dict
    geo_address_json: Any
    geo_latitude: Any
    geo_longitude: Any
    id: str
    location_type: str
    location_visibility: str
    managing_calendar: list
    meeting_url: Any
    name: str
    platform: str
    registration_open: bool
    require_approval: bool
    spots_remaining: Any
    start_at: str
    timezone: str
    url: str
    user_api_id: str
    user_id: str
    visibility: str
    waitlist_status: str
    zoom_meeting_url: Any


class OrganizationEvent(OrganizationEventRequired, total=False):
    registration_question: list


class OrganizationEventListMatch(TypedDict, total=False):
    api_id: str
    calendar_api_id: str
    calendar_id: str
    coordinate: Any
    cover_url: str
    created_at: str
    display_price: Any
    duration_interval: str
    end_at: str
    feedback_email: dict
    geo_address_json: Any
    geo_latitude: Any
    geo_longitude: Any
    id: str
    location_type: str
    location_visibility: str
    managing_calendar: list
    meeting_url: Any
    name: str
    platform: str
    registration_open: bool
    registration_question: list
    require_approval: bool
    spots_remaining: Any
    start_at: str
    timezone: str
    url: str
    user_api_id: str
    user_id: str
    visibility: str
    waitlist_status: str
    zoom_meeting_url: Any


class OrganizationEventTransfer(TypedDict):
    calendar_id: str
    event_id: str


class OrganizationEventTransferCreateData(TypedDict):
    calendar_id: str
    event_id: str


class TicketTypeRequired(TypedDict):
    id: str
    name: str
    type: str


class TicketType(TicketTypeRequired, total=False):
    cent: Any
    currency: Any
    description: str
    is_flexible: bool
    is_hidden: bool
    max_capacity: Any
    min_cent: Any
    require_approval: bool
    valid_end_at: Any
    valid_start_at: Any


class TicketTypeLoadMatchRequired(TypedDict):
    id: str


class TicketTypeLoadMatch(TicketTypeLoadMatchRequired, total=False):
    cent: Any
    currency: Any
    description: str
    is_flexible: bool
    is_hidden: bool
    max_capacity: Any
    min_cent: Any
    name: str
    require_approval: bool
    type: str
    valid_end_at: Any
    valid_start_at: Any


class TicketTypeListMatch(TypedDict, total=False):
    cent: Any
    currency: Any
    description: str
    id: str
    is_flexible: bool
    is_hidden: bool
    max_capacity: Any
    min_cent: Any
    name: str
    require_approval: bool
    type: str
    valid_end_at: Any
    valid_start_at: Any


class TicketTypeCreateDataRequired(TypedDict):
    id: str
    name: str
    type: str


class TicketTypeCreateData(TicketTypeCreateDataRequired, total=False):
    cent: Any
    currency: Any
    description: str
    is_flexible: bool
    is_hidden: bool
    max_capacity: Any
    min_cent: Any
    require_approval: bool
    valid_end_at: Any
    valid_start_at: Any


class TicketTypeUpdateData(TypedDict, total=False):
    cent: Any
    currency: Any
    description: str
    id: str
    is_flexible: bool
    is_hidden: bool
    max_capacity: Any
    min_cent: Any
    name: str
    require_approval: bool
    type: str
    valid_end_at: Any
    valid_start_at: Any


class TicketTypeRemoveMatchRequired(TypedDict):
    id: str


class TicketTypeRemoveMatch(TicketTypeRemoveMatchRequired, total=False):
    cent: Any
    currency: Any
    description: str
    is_flexible: bool
    is_hidden: bool
    max_capacity: Any
    min_cent: Any
    name: str
    require_approval: bool
    type: str
    valid_end_at: Any
    valid_start_at: Any


class User(TypedDict):
    avatar_url: str
    email: str
    first_name: Any
    id: str
    last_name: Any
    name: str


class UserLoadMatchRequired(TypedDict):
    id: str


class UserLoadMatch(UserLoadMatchRequired, total=False):
    avatar_url: str
    email: str
    first_name: Any
    last_name: Any
    name: str


class Webhook(TypedDict):
    created_at: str
    event_type: list
    id: str
    secret: str
    status: str
    url: str


class WebhookLoadMatchRequired(TypedDict):
    id: str


class WebhookLoadMatch(WebhookLoadMatchRequired, total=False):
    created_at: str
    event_type: list
    secret: str
    status: str
    url: str


class WebhookListMatch(TypedDict, total=False):
    created_at: str
    event_type: list
    id: str
    secret: str
    status: str
    url: str


class WebhookCreateData(TypedDict):
    created_at: str
    event_type: list
    id: str
    secret: str
    status: str
    url: str


class WebhookUpdateData(TypedDict, total=False):
    created_at: str
    event_type: list
    id: str
    secret: str
    status: str
    url: str


class WebhookRemoveMatchRequired(TypedDict):
    id: str


class WebhookRemoveMatch(WebhookRemoveMatchRequired, total=False):
    created_at: str
    event_type: list
    secret: str
    status: str
    url: str
