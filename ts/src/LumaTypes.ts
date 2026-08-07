// Typed models for the Luma SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface Calendar {
  avatar_url: any
  calendar_id: string
  coordinate: any
  cover_image_url: any
  description: string
  id: string
  instagram_handle: any
  is_personal: boolean
  location: any
  name: string
  slug: string
  social_image_url: any
  tint_color?: string
  twitter_handle: any
  url: string
  website: any
  youtube_handle: any
}

export interface CalendarLoadMatch {
  avatar_url?: any
  calendar_id?: string
  coordinate?: any
  cover_image_url?: any
  description?: string
  id: string
  instagram_handle?: any
  is_personal?: boolean
  location?: any
  name?: string
  slug?: string
  social_image_url?: any
  tint_color?: string
  twitter_handle?: any
  url?: string
  website?: any
  youtube_handle?: any
}

export interface CalendarUpdateData {
  avatar_url?: any
  calendar_id?: string
  coordinate?: any
  cover_image_url?: any
  description?: string
  id?: string
  instagram_handle?: any
  is_personal?: boolean
  location?: any
  name?: string
  slug?: string
  social_image_url?: any
  tint_color?: string
  twitter_handle?: any
  url?: string
  website?: any
  youtube_handle?: any
}

export interface CalendarAdmin {
  avatar_url: string
  email: string
  first_name: any
  id: string
  last_name: any
  name: string
}

export interface CalendarAdminListMatch {
  avatar_url?: string
  email?: string
  first_name?: any
  id?: string
  last_name?: any
  name?: string
}

export interface CalendarCoupon {
  cents_off: any
  code: string
  currency: any
  discount: any
  event_ticket_type_id?: string
  id: string
  percent_off: any
  remaining_count: number
  valid_end_at: any
  valid_start_at: any
}

export interface CalendarCouponListMatch {
  cents_off?: any
  code?: string
  currency?: any
  discount?: any
  event_ticket_type_id?: string
  id?: string
  percent_off?: any
  remaining_count?: number
  valid_end_at?: any
  valid_start_at?: any
}

export interface CalendarCouponCreateData {
  cents_off: any
  code: string
  currency: any
  discount: any
  event_ticket_type_id?: string
  id: string
  percent_off: any
  remaining_count: number
  valid_end_at: any
  valid_start_at: any
}

export interface CalendarCouponUpdateData {
  cents_off?: any
  code?: string
  currency?: any
  discount?: any
  event_ticket_type_id?: string
  id?: string
  percent_off?: any
  remaining_count?: number
  valid_end_at?: any
  valid_start_at?: any
}

export interface CalendarEvent {
  id: string
  status: string
  submitted_by: any
  tag: any[]
}

export interface CalendarEventLoadMatch {
  id: string
  status?: string
  submitted_by?: any
  tag?: any[]
}

export interface CalendarEventListMatch {
  id?: string
  status?: string
  submitted_by?: any
  tag?: any[]
}

export interface CalendarEventCreateData {
  id: string
  status: string
  submitted_by: any
  tag: any[]
}

export interface CalendarEventApproval {
  calendar_event_id: string
}

export interface CalendarEventApprovalCreateData {
  calendar_event_id: string
}

export interface CalendarEventRejection {
  calendar_event_id: string
  message?: string
}

export interface CalendarEventRejectionCreateData {
  calendar_event_id: string
  message?: string
}

export interface Contact {
  avatar_url: string
  contact: any[]
  created_at: string
  email: string
  event_approved_count: number
  event_checked_in_count: number
  first_name: any
  id: string
  last_name: any
  membership: any
  name: string
  revenue_usd_cent: number
  tag?: any
  user_id: string
}

export interface ContactListMatch {
  avatar_url?: string
  contact?: any[]
  created_at?: string
  email?: string
  event_approved_count?: number
  event_checked_in_count?: number
  first_name?: any
  id?: string
  last_name?: any
  membership?: any
  name?: string
  revenue_usd_cent?: number
  tag?: any
  user_id?: string
}

export interface ContactCreateData {
  avatar_url: string
  contact: any[]
  created_at: string
  email: string
  event_approved_count: number
  event_checked_in_count: number
  first_name: any
  id: string
  last_name: any
  membership: any
  name: string
  revenue_usd_cent: number
  tag?: any
  user_id: string
}

export interface ContactRemoveMatch {
  avatar_url?: string
  contact?: any[]
  created_at?: string
  email?: string
  event_approved_count?: number
  event_checked_in_count?: number
  first_name?: any
  id: string
  last_name?: any
  membership?: any
  name?: string
  revenue_usd_cent?: number
  tag?: any
  user_id?: string
}

export interface ContactBlock {
  contact_id?: string
  email?: string
}

export interface ContactBlockCreateData {
  contact_id?: string
  email?: string
}

export interface ContactRestore {
  contact_id?: string
  email?: string
}

export interface ContactRestoreCreateData {
  contact_id?: string
  email?: string
}

export interface ContactTag {
  color?: any
  id: string
  name: string
  tag_id: string
}

export interface ContactTagListMatch {
  color?: any
  id?: string
  name?: string
  tag_id?: string
}

export interface ContactTagCreateData {
  color?: any
  id: string
  name: string
  tag_id: string
}

export interface ContactTagUpdateData {
  color?: any
  id?: string
  name?: string
  tag_id?: string
}

export interface ContactTagRemoveMatch {
  color?: any
  id: string
  name?: string
  tag_id?: string
}

export interface ContactTagAssignment {
  applied_count: number
  email?: any[]
  skipped_count: number
  tag: string
  user_id?: any[]
}

export interface ContactTagAssignmentCreateData {
  applied_count: number
  email?: any[]
  skipped_count: number
  tag: string
  user_id?: any[]
}

export interface ContactTagAssignmentRemoveMatch {
  applied_count?: number
  email?: any[]
  skipped_count?: number
  tag?: string
  user_id?: any[]
}

export interface EntityLookup {
}

export interface EntityLookupLoadMatch {
}

export interface Event {
  access: string
  calendar_id: string
  can_register_for_multiple_ticket?: boolean
  coordinate: any
  cover_url: string
  created_at: string
  description: string
  description_md: string
  display_price: any
  duration_interval: string
  end_at: string
  event_id: string
  feedback_email: Record<string, any>
  geo_address_json: any
  guest_count: Record<string, any>
  host: any[]
  id: string
  location_type: string
  location_visibility: string
  max_capacity?: any
  meeting_url: any
  name: string
  name_requirement?: string
  phone_number_requirement?: any
  platform: string
  registration_open: boolean
  registration_question?: any[]
  reminders_disabled?: boolean
  require_approval: boolean
  show_guest_list?: boolean
  slug?: string
  spots_remaining: any
  start_at: string
  suppress_notification?: boolean
  timezone: string
  tint_color?: string
  url: string
  user_id: string
  visibility: string
  waitlist_status: string
}

export interface EventLoadMatch {
  access?: string
  calendar_id?: string
  can_register_for_multiple_ticket?: boolean
  coordinate?: any
  cover_url?: string
  created_at?: string
  description?: string
  description_md?: string
  display_price?: any
  duration_interval?: string
  end_at?: string
  event_id?: string
  feedback_email?: Record<string, any>
  geo_address_json?: any
  guest_count?: Record<string, any>
  host?: any[]
  id: string
  location_type?: string
  location_visibility?: string
  max_capacity?: any
  meeting_url?: any
  name?: string
  name_requirement?: string
  phone_number_requirement?: any
  platform?: string
  registration_open?: boolean
  registration_question?: any[]
  reminders_disabled?: boolean
  require_approval?: boolean
  show_guest_list?: boolean
  slug?: string
  spots_remaining?: any
  start_at?: string
  suppress_notification?: boolean
  timezone?: string
  tint_color?: string
  url?: string
  user_id?: string
  visibility?: string
  waitlist_status?: string
}

export interface EventCreateData {
  access: string
  calendar_id: string
  can_register_for_multiple_ticket?: boolean
  coordinate: any
  cover_url: string
  created_at: string
  description: string
  description_md: string
  display_price: any
  duration_interval: string
  end_at: string
  event_id: string
  feedback_email: Record<string, any>
  geo_address_json: any
  guest_count: Record<string, any>
  host: any[]
  id: string
  location_type: string
  location_visibility: string
  max_capacity?: any
  meeting_url: any
  name: string
  name_requirement?: string
  phone_number_requirement?: any
  platform: string
  registration_open: boolean
  registration_question?: any[]
  reminders_disabled?: boolean
  require_approval: boolean
  show_guest_list?: boolean
  slug?: string
  spots_remaining: any
  start_at: string
  suppress_notification?: boolean
  timezone: string
  tint_color?: string
  url: string
  user_id: string
  visibility: string
  waitlist_status: string
}

export interface EventUpdateData {
  access?: string
  calendar_id?: string
  can_register_for_multiple_ticket?: boolean
  coordinate?: any
  cover_url?: string
  created_at?: string
  description?: string
  description_md?: string
  display_price?: any
  duration_interval?: string
  end_at?: string
  event_id?: string
  feedback_email?: Record<string, any>
  geo_address_json?: any
  guest_count?: Record<string, any>
  host?: any[]
  id?: string
  location_type?: string
  location_visibility?: string
  max_capacity?: any
  meeting_url?: any
  name?: string
  name_requirement?: string
  phone_number_requirement?: any
  platform?: string
  registration_open?: boolean
  registration_question?: any[]
  reminders_disabled?: boolean
  require_approval?: boolean
  show_guest_list?: boolean
  slug?: string
  spots_remaining?: any
  start_at?: string
  suppress_notification?: boolean
  timezone?: string
  tint_color?: string
  url?: string
  user_id?: string
  visibility?: string
  waitlist_status?: string
}

export interface EventRemoveMatch {
  access?: string
  calendar_id?: string
  can_register_for_multiple_ticket?: boolean
  coordinate?: any
  cover_url?: string
  created_at?: string
  description?: string
  description_md?: string
  display_price?: any
  duration_interval?: string
  end_at?: string
  event_id?: string
  feedback_email?: Record<string, any>
  geo_address_json?: any
  guest_count?: Record<string, any>
  host?: any[]
  id: string
  location_type?: string
  location_visibility?: string
  max_capacity?: any
  meeting_url?: any
  name?: string
  name_requirement?: string
  phone_number_requirement?: any
  platform?: string
  registration_open?: boolean
  registration_question?: any[]
  reminders_disabled?: boolean
  require_approval?: boolean
  show_guest_list?: boolean
  slug?: string
  spots_remaining?: any
  start_at?: string
  suppress_notification?: boolean
  timezone?: string
  tint_color?: string
  url?: string
  user_id?: string
  visibility?: string
  waitlist_status?: string
}

export interface EventCancelRequest {
  cancellation_token: string
  event_id: string
  guest_count: number
  is_paid: boolean
}

export interface EventCancelRequestCreateData {
  cancellation_token: string
  event_id: string
  guest_count: number
  is_paid: boolean
}

export interface EventCoupon {
  cents_off: any
  code: string
  currency: any
  discount: any
  event_id: string
  event_ticket_type_id?: string
  id: string
  percent_off: any
  remaining_count: number
  valid_end_at: any
  valid_start_at: any
}

export interface EventCouponListMatch {
  cents_off?: any
  code?: string
  currency?: any
  discount?: any
  event_id?: string
  event_ticket_type_id?: string
  id?: string
  percent_off?: any
  remaining_count?: number
  valid_end_at?: any
  valid_start_at?: any
}

export interface EventCouponCreateData {
  cents_off: any
  code: string
  currency: any
  discount: any
  event_id: string
  event_ticket_type_id?: string
  id: string
  percent_off: any
  remaining_count: number
  valid_end_at: any
  valid_start_at: any
}

export interface EventCouponUpdateData {
  cents_off?: any
  code?: string
  currency?: any
  discount?: any
  event_id?: string
  event_ticket_type_id?: string
  id?: string
  percent_off?: any
  remaining_count?: number
  valid_end_at?: any
  valid_start_at?: any
}

export interface EventTag {
  color?: any
  id: string
  name: string
  tag_id: string
}

export interface EventTagListMatch {
  color?: any
  id?: string
  name?: string
  tag_id?: string
}

export interface EventTagCreateData {
  color?: any
  id: string
  name: string
  tag_id: string
}

export interface EventTagUpdateData {
  color?: any
  id?: string
  name?: string
  tag_id?: string
}

export interface EventTagRemoveMatch {
  color?: any
  id: string
  name?: string
  tag_id?: string
}

export interface EventTagAssignment {
  applied_count: number
  event_id: any[]
  skipped_count: number
  tag: string
}

export interface EventTagAssignmentCreateData {
  applied_count: number
  event_id: any[]
  skipped_count: number
  tag: string
}

export interface EventTagAssignmentRemoveMatch {
  applied_count?: number
  event_id?: any[]
  skipped_count?: number
  tag?: string
}

export interface Guest {
  approval_status: string
  check_in_qr_code: string
  eth_address: any
  event_id: string
  event_ticket: any[]
  event_ticket_order: any[]
  guest: any[]
  guest_id: string
  id: string
  invited_at: any
  joined_at: any
  message?: any
  phone_number: number
  registered_at: any
  registration_answer: any
  send_email?: any
  should_refund?: boolean
  solana_address: any
  status: string
  ticket?: any
  user_email: string
  user_first_name: any
  user_id: string
  user_last_name: any
  user_name: any
  utm_source: any
}

export interface GuestLoadMatch {
  approval_status?: string
  check_in_qr_code?: string
  eth_address?: any
  event_id?: string
  event_ticket?: any[]
  event_ticket_order?: any[]
  guest?: any[]
  guest_id?: string
  id: string
  invited_at?: any
  joined_at?: any
  message?: any
  phone_number?: number
  registered_at?: any
  registration_answer?: any
  send_email?: any
  should_refund?: boolean
  solana_address?: any
  status?: string
  ticket?: any
  user_email?: string
  user_first_name?: any
  user_id?: string
  user_last_name?: any
  user_name?: any
  utm_source?: any
}

export interface GuestListMatch {
  approval_status?: string
  check_in_qr_code?: string
  eth_address?: any
  event_id?: string
  event_ticket?: any[]
  event_ticket_order?: any[]
  guest?: any[]
  guest_id?: string
  id?: string
  invited_at?: any
  joined_at?: any
  message?: any
  phone_number?: number
  registered_at?: any
  registration_answer?: any
  send_email?: any
  should_refund?: boolean
  solana_address?: any
  status?: string
  ticket?: any
  user_email?: string
  user_first_name?: any
  user_id?: string
  user_last_name?: any
  user_name?: any
  utm_source?: any
}

export interface GuestCreateData {
  approval_status: string
  check_in_qr_code: string
  eth_address: any
  event_id: string
  event_ticket: any[]
  event_ticket_order: any[]
  guest: any[]
  guest_id: string
  id: string
  invited_at: any
  joined_at: any
  message?: any
  phone_number: number
  registered_at: any
  registration_answer: any
  send_email?: any
  should_refund?: boolean
  solana_address: any
  status: string
  ticket?: any
  user_email: string
  user_first_name: any
  user_id: string
  user_last_name: any
  user_name: any
  utm_source: any
}

export interface GuestUpdateData {
  approval_status?: string
  check_in_qr_code?: string
  eth_address?: any
  event_id?: string
  event_ticket?: any[]
  event_ticket_order?: any[]
  guest?: any[]
  guest_id?: string
  id?: string
  invited_at?: any
  joined_at?: any
  message?: any
  phone_number?: number
  registered_at?: any
  registration_answer?: any
  send_email?: any
  should_refund?: boolean
  solana_address?: any
  status?: string
  ticket?: any
  user_email?: string
  user_first_name?: any
  user_id?: string
  user_last_name?: any
  user_name?: any
  utm_source?: any
}

export interface GuestInvite {
  event_id: string
  guest: any[]
  message?: any
}

export interface GuestInviteCreateData {
  event_id: string
  guest: any[]
  message?: any
}

export interface GuestTicket {
  event_id: string
  guest_id: string
  send_email?: any
  ticket_ids_to_remove?: any[]
  tickets_to_add?: any[]
}

export interface GuestTicketUpdateData {
  event_id?: string
  guest_id?: string
  send_email?: any
  ticket_ids_to_remove?: any[]
  tickets_to_add?: any[]
}

export interface Host {
  access_level?: any
  email: string
  event_id: string
  is_visible?: boolean
  name?: string
}

export interface HostCreateData {
  access_level?: any
  email: string
  event_id: string
  is_visible?: boolean
  name?: string
}

export interface HostUpdateData {
  access_level?: any
  email?: string
  event_id?: string
  is_visible?: boolean
  name?: string
}

export interface HostRemoveMatch {
  access_level?: any
  email?: string
  event_id?: string
  is_visible?: boolean
  name?: string
}

export interface ImageUpload {
  content_type?: any
  file_url: string
  upload_url: string
}

export interface ImageUploadCreateData {
  content_type?: any
  file_url: string
  upload_url: string
}

export interface Member {
  email: string
  membership_id: string
  membership_tier_id: string
  registration_answer?: any[]
  skip_payment?: boolean
  status: string
  user_id: string
}

export interface MemberCreateData {
  email: string
  membership_id: string
  membership_tier_id: string
  registration_answer?: any[]
  skip_payment?: boolean
  status: string
  user_id: string
}

export interface MemberUpdateData {
  email?: string
  membership_id?: string
  membership_tier_id?: string
  registration_answer?: any[]
  skip_payment?: boolean
  status?: string
  user_id?: string
}

export interface MembershipTier {
  access_info: any
  description: string
  id: string
  name: string
  tint_color: string
}

export interface MembershipTierListMatch {
  access_info?: any
  description?: string
  id?: string
  name?: string
  tint_color?: string
}

export interface OrganizationAdmin {
  api_id: string
  avatar_url: string
  email: string
  first_name: any
  id: string
  last_name: any
  name: string
}

export interface OrganizationAdminListMatch {
  api_id?: string
  avatar_url?: string
  email?: string
  first_name?: any
  id?: string
  last_name?: any
  name?: string
}

export interface OrganizationCalendar {
  avatar_url: any
  coordinate: any
  cover_image_url: any
  description: string
  id: string
  instagram_handle: any
  is_personal: boolean
  location: any
  name: string
  slug: string
  social_image_url: any
  tint_color?: string
  twitter_handle: any
  url: string
  website: any
  youtube_handle: any
}

export interface OrganizationCalendarListMatch {
  avatar_url?: any
  coordinate?: any
  cover_image_url?: any
  description?: string
  id?: string
  instagram_handle?: any
  is_personal?: boolean
  location?: any
  name?: string
  slug?: string
  social_image_url?: any
  tint_color?: string
  twitter_handle?: any
  url?: string
  website?: any
  youtube_handle?: any
}

export interface OrganizationCalendarCreateData {
  avatar_url: any
  coordinate: any
  cover_image_url: any
  description: string
  id: string
  instagram_handle: any
  is_personal: boolean
  location: any
  name: string
  slug: string
  social_image_url: any
  tint_color?: string
  twitter_handle: any
  url: string
  website: any
  youtube_handle: any
}

export interface OrganizationEvent {
  api_id: string
  calendar_api_id: string
  calendar_id: string
  coordinate: any
  cover_url: string
  created_at: string
  display_price: any
  duration_interval: string
  end_at: string
  feedback_email: Record<string, any>
  geo_address_json: any
  geo_latitude: any
  geo_longitude: any
  id: string
  location_type: string
  location_visibility: string
  managing_calendar: any[]
  meeting_url: any
  name: string
  platform: string
  registration_open: boolean
  registration_question?: any[]
  require_approval: boolean
  spots_remaining: any
  start_at: string
  timezone: string
  url: string
  user_api_id: string
  user_id: string
  visibility: string
  waitlist_status: string
  zoom_meeting_url: any
}

export interface OrganizationEventListMatch {
  api_id?: string
  calendar_api_id?: string
  calendar_id?: string
  coordinate?: any
  cover_url?: string
  created_at?: string
  display_price?: any
  duration_interval?: string
  end_at?: string
  feedback_email?: Record<string, any>
  geo_address_json?: any
  geo_latitude?: any
  geo_longitude?: any
  id?: string
  location_type?: string
  location_visibility?: string
  managing_calendar?: any[]
  meeting_url?: any
  name?: string
  platform?: string
  registration_open?: boolean
  registration_question?: any[]
  require_approval?: boolean
  spots_remaining?: any
  start_at?: string
  timezone?: string
  url?: string
  user_api_id?: string
  user_id?: string
  visibility?: string
  waitlist_status?: string
  zoom_meeting_url?: any
}

export interface OrganizationEventTransfer {
  calendar_id: string
  event_id: string
}

export interface OrganizationEventTransferCreateData {
  calendar_id: string
  event_id: string
}

export interface TicketType {
  cent?: any
  currency?: any
  description?: string
  id: string
  is_flexible?: boolean
  is_hidden?: boolean
  max_capacity?: any
  min_cent?: any
  name: string
  require_approval?: boolean
  type: string
  valid_end_at?: any
  valid_start_at?: any
}

export interface TicketTypeLoadMatch {
  cent?: any
  currency?: any
  description?: string
  id: string
  is_flexible?: boolean
  is_hidden?: boolean
  max_capacity?: any
  min_cent?: any
  name?: string
  require_approval?: boolean
  type?: string
  valid_end_at?: any
  valid_start_at?: any
}

export interface TicketTypeListMatch {
  cent?: any
  currency?: any
  description?: string
  id?: string
  is_flexible?: boolean
  is_hidden?: boolean
  max_capacity?: any
  min_cent?: any
  name?: string
  require_approval?: boolean
  type?: string
  valid_end_at?: any
  valid_start_at?: any
}

export interface TicketTypeCreateData {
  cent?: any
  currency?: any
  description?: string
  id: string
  is_flexible?: boolean
  is_hidden?: boolean
  max_capacity?: any
  min_cent?: any
  name: string
  require_approval?: boolean
  type: string
  valid_end_at?: any
  valid_start_at?: any
}

export interface TicketTypeUpdateData {
  cent?: any
  currency?: any
  description?: string
  id?: string
  is_flexible?: boolean
  is_hidden?: boolean
  max_capacity?: any
  min_cent?: any
  name?: string
  require_approval?: boolean
  type?: string
  valid_end_at?: any
  valid_start_at?: any
}

export interface TicketTypeRemoveMatch {
  cent?: any
  currency?: any
  description?: string
  id: string
  is_flexible?: boolean
  is_hidden?: boolean
  max_capacity?: any
  min_cent?: any
  name?: string
  require_approval?: boolean
  type?: string
  valid_end_at?: any
  valid_start_at?: any
}

export interface User {
  avatar_url: string
  email: string
  first_name: any
  id: string
  last_name: any
  name: string
}

export interface UserLoadMatch {
  avatar_url?: string
  email?: string
  first_name?: any
  id: string
  last_name?: any
  name?: string
}

export interface Webhook {
  created_at: string
  event_type: any[]
  id: string
  secret: string
  status: string
  url: string
}

export interface WebhookLoadMatch {
  created_at?: string
  event_type?: any[]
  id: string
  secret?: string
  status?: string
  url?: string
}

export interface WebhookListMatch {
  created_at?: string
  event_type?: any[]
  id?: string
  secret?: string
  status?: string
  url?: string
}

export interface WebhookCreateData {
  created_at: string
  event_type: any[]
  id: string
  secret: string
  status: string
  url: string
}

export interface WebhookUpdateData {
  created_at?: string
  event_type?: any[]
  id?: string
  secret?: string
  status?: string
  url?: string
}

export interface WebhookRemoveMatch {
  created_at?: string
  event_type?: any[]
  id: string
  secret?: string
  status?: string
  url?: string
}

