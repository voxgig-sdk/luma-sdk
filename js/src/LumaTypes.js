// Typed models for the Luma SDK (JSDoc typedefs).
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
// edit by hand.

/**
 * @typedef {Object} Calendar
 * @property {*} avatar_url
 * @property {string} calendar_id
 * @property {*} coordinate
 * @property {*} cover_image_url
 * @property {string} description
 * @property {string} id
 * @property {*} instagram_handle
 * @property {boolean} is_personal
 * @property {*} location
 * @property {string} name
 * @property {string} slug
 * @property {*} social_image_url
 * @property {string} [tint_color]
 * @property {*} twitter_handle
 * @property {string} url
 * @property {*} website
 * @property {*} youtube_handle
 */

/**
 * @typedef {Object} CalendarLoadMatch
 * @property {*} [avatar_url]
 * @property {string} [calendar_id]
 * @property {*} [coordinate]
 * @property {*} [cover_image_url]
 * @property {string} [description]
 * @property {string} id
 * @property {*} [instagram_handle]
 * @property {boolean} [is_personal]
 * @property {*} [location]
 * @property {string} [name]
 * @property {string} [slug]
 * @property {*} [social_image_url]
 * @property {string} [tint_color]
 * @property {*} [twitter_handle]
 * @property {string} [url]
 * @property {*} [website]
 * @property {*} [youtube_handle]
 */

/**
 * @typedef {Object} CalendarUpdateData
 * @property {*} [avatar_url]
 * @property {string} [calendar_id]
 * @property {*} [coordinate]
 * @property {*} [cover_image_url]
 * @property {string} [description]
 * @property {string} [id]
 * @property {*} [instagram_handle]
 * @property {boolean} [is_personal]
 * @property {*} [location]
 * @property {string} [name]
 * @property {string} [slug]
 * @property {*} [social_image_url]
 * @property {string} [tint_color]
 * @property {*} [twitter_handle]
 * @property {string} [url]
 * @property {*} [website]
 * @property {*} [youtube_handle]
 */

/**
 * @typedef {Object} CalendarAdmin
 * @property {string} avatar_url
 * @property {string} email
 * @property {*} first_name
 * @property {string} id
 * @property {*} last_name
 * @property {string} name
 */

/**
 * @typedef {Object} CalendarAdminListMatch
 * @property {string} [avatar_url]
 * @property {string} [email]
 * @property {*} [first_name]
 * @property {string} [id]
 * @property {*} [last_name]
 * @property {string} [name]
 */

/**
 * @typedef {Object} CalendarCoupon
 * @property {*} cents_off
 * @property {string} code
 * @property {*} currency
 * @property {*} discount
 * @property {string} [event_ticket_type_id]
 * @property {string} id
 * @property {*} percent_off
 * @property {number} remaining_count
 * @property {*} valid_end_at
 * @property {*} valid_start_at
 */

/**
 * @typedef {Object} CalendarCouponListMatch
 * @property {*} [cents_off]
 * @property {string} [code]
 * @property {*} [currency]
 * @property {*} [discount]
 * @property {string} [event_ticket_type_id]
 * @property {string} [id]
 * @property {*} [percent_off]
 * @property {number} [remaining_count]
 * @property {*} [valid_end_at]
 * @property {*} [valid_start_at]
 */

/**
 * @typedef {Object} CalendarCouponCreateData
 * @property {*} cents_off
 * @property {string} code
 * @property {*} currency
 * @property {*} discount
 * @property {string} [event_ticket_type_id]
 * @property {string} id
 * @property {*} percent_off
 * @property {number} remaining_count
 * @property {*} valid_end_at
 * @property {*} valid_start_at
 */

/**
 * @typedef {Object} CalendarCouponUpdateData
 * @property {*} [cents_off]
 * @property {string} [code]
 * @property {*} [currency]
 * @property {*} [discount]
 * @property {string} [event_ticket_type_id]
 * @property {string} [id]
 * @property {*} [percent_off]
 * @property {number} [remaining_count]
 * @property {*} [valid_end_at]
 * @property {*} [valid_start_at]
 */

/**
 * @typedef {Object} CalendarEvent
 * @property {string} id
 * @property {string} status
 * @property {*} submitted_by
 * @property {Array} tag
 */

/**
 * @typedef {Object} CalendarEventLoadMatch
 * @property {string} id
 * @property {string} [status]
 * @property {*} [submitted_by]
 * @property {Array} [tag]
 */

/**
 * @typedef {Object} CalendarEventListMatch
 * @property {string} [id]
 * @property {string} [status]
 * @property {*} [submitted_by]
 * @property {Array} [tag]
 */

/**
 * @typedef {Object} CalendarEventCreateData
 * @property {string} id
 * @property {string} status
 * @property {*} submitted_by
 * @property {Array} tag
 */

/**
 * @typedef {Object} CalendarEventApproval
 * @property {string} calendar_event_id
 */

/**
 * @typedef {Object} CalendarEventApprovalCreateData
 * @property {string} calendar_event_id
 */

/**
 * @typedef {Object} CalendarEventRejection
 * @property {string} calendar_event_id
 * @property {string} [message]
 */

/**
 * @typedef {Object} CalendarEventRejectionCreateData
 * @property {string} calendar_event_id
 * @property {string} [message]
 */

/**
 * @typedef {Object} Contact
 * @property {string} avatar_url
 * @property {Array} contact
 * @property {string} created_at
 * @property {string} email
 * @property {number} event_approved_count
 * @property {number} event_checked_in_count
 * @property {*} first_name
 * @property {string} id
 * @property {*} last_name
 * @property {*} membership
 * @property {string} name
 * @property {number} revenue_usd_cent
 * @property {*} [tag]
 * @property {string} user_id
 */

/**
 * @typedef {Object} ContactListMatch
 * @property {string} [avatar_url]
 * @property {Array} [contact]
 * @property {string} [created_at]
 * @property {string} [email]
 * @property {number} [event_approved_count]
 * @property {number} [event_checked_in_count]
 * @property {*} [first_name]
 * @property {string} [id]
 * @property {*} [last_name]
 * @property {*} [membership]
 * @property {string} [name]
 * @property {number} [revenue_usd_cent]
 * @property {*} [tag]
 * @property {string} [user_id]
 */

/**
 * @typedef {Object} ContactCreateData
 * @property {string} avatar_url
 * @property {Array} contact
 * @property {string} created_at
 * @property {string} email
 * @property {number} event_approved_count
 * @property {number} event_checked_in_count
 * @property {*} first_name
 * @property {string} id
 * @property {*} last_name
 * @property {*} membership
 * @property {string} name
 * @property {number} revenue_usd_cent
 * @property {*} [tag]
 * @property {string} user_id
 */

/**
 * @typedef {Object} ContactRemoveMatch
 * @property {string} [avatar_url]
 * @property {Array} [contact]
 * @property {string} [created_at]
 * @property {string} [email]
 * @property {number} [event_approved_count]
 * @property {number} [event_checked_in_count]
 * @property {*} [first_name]
 * @property {string} id
 * @property {*} [last_name]
 * @property {*} [membership]
 * @property {string} [name]
 * @property {number} [revenue_usd_cent]
 * @property {*} [tag]
 * @property {string} [user_id]
 */

/**
 * @typedef {Object} ContactBlock
 * @property {string} [contact_id]
 * @property {string} [email]
 */

/**
 * @typedef {Object} ContactBlockCreateData
 * @property {string} [contact_id]
 * @property {string} [email]
 */

/**
 * @typedef {Object} ContactRestore
 * @property {string} [contact_id]
 * @property {string} [email]
 */

/**
 * @typedef {Object} ContactRestoreCreateData
 * @property {string} [contact_id]
 * @property {string} [email]
 */

/**
 * @typedef {Object} ContactTag
 * @property {*} [color]
 * @property {string} id
 * @property {string} name
 * @property {string} tag_id
 */

/**
 * @typedef {Object} ContactTagListMatch
 * @property {*} [color]
 * @property {string} [id]
 * @property {string} [name]
 * @property {string} [tag_id]
 */

/**
 * @typedef {Object} ContactTagCreateData
 * @property {*} [color]
 * @property {string} id
 * @property {string} name
 * @property {string} tag_id
 */

/**
 * @typedef {Object} ContactTagUpdateData
 * @property {*} [color]
 * @property {string} [id]
 * @property {string} [name]
 * @property {string} [tag_id]
 */

/**
 * @typedef {Object} ContactTagRemoveMatch
 * @property {*} [color]
 * @property {string} id
 * @property {string} [name]
 * @property {string} [tag_id]
 */

/**
 * @typedef {Object} ContactTagAssignment
 * @property {number} applied_count
 * @property {Array} [email]
 * @property {number} skipped_count
 * @property {string} tag
 * @property {Array} [user_id]
 */

/**
 * @typedef {Object} ContactTagAssignmentCreateData
 * @property {number} applied_count
 * @property {Array} [email]
 * @property {number} skipped_count
 * @property {string} tag
 * @property {Array} [user_id]
 */

/**
 * @typedef {Object} ContactTagAssignmentRemoveMatch
 * @property {number} [applied_count]
 * @property {Array} [email]
 * @property {number} [skipped_count]
 * @property {string} [tag]
 * @property {Array} [user_id]
 */

/**
 * @typedef {Object} EntityLookup
 */

/**
 * @typedef {Object} EntityLookupLoadMatch
 */

/**
 * @typedef {Object} Event
 * @property {string} access
 * @property {string} calendar_id
 * @property {boolean} [can_register_for_multiple_ticket]
 * @property {*} coordinate
 * @property {string} cover_url
 * @property {string} created_at
 * @property {string} description
 * @property {string} description_md
 * @property {*} display_price
 * @property {string} duration_interval
 * @property {string} end_at
 * @property {string} event_id
 * @property {Object} feedback_email
 * @property {*} geo_address_json
 * @property {Object} guest_count
 * @property {Array} host
 * @property {string} id
 * @property {string} location_type
 * @property {string} location_visibility
 * @property {*} [max_capacity]
 * @property {*} meeting_url
 * @property {string} name
 * @property {string} [name_requirement]
 * @property {*} [phone_number_requirement]
 * @property {string} platform
 * @property {boolean} registration_open
 * @property {Array} [registration_question]
 * @property {boolean} [reminders_disabled]
 * @property {boolean} require_approval
 * @property {boolean} [show_guest_list]
 * @property {string} [slug]
 * @property {*} spots_remaining
 * @property {string} start_at
 * @property {boolean} [suppress_notification]
 * @property {string} timezone
 * @property {string} [tint_color]
 * @property {string} url
 * @property {string} user_id
 * @property {string} visibility
 * @property {string} waitlist_status
 */

/**
 * @typedef {Object} EventLoadMatch
 * @property {string} [access]
 * @property {string} [calendar_id]
 * @property {boolean} [can_register_for_multiple_ticket]
 * @property {*} [coordinate]
 * @property {string} [cover_url]
 * @property {string} [created_at]
 * @property {string} [description]
 * @property {string} [description_md]
 * @property {*} [display_price]
 * @property {string} [duration_interval]
 * @property {string} [end_at]
 * @property {string} [event_id]
 * @property {Object} [feedback_email]
 * @property {*} [geo_address_json]
 * @property {Object} [guest_count]
 * @property {Array} [host]
 * @property {string} id
 * @property {string} [location_type]
 * @property {string} [location_visibility]
 * @property {*} [max_capacity]
 * @property {*} [meeting_url]
 * @property {string} [name]
 * @property {string} [name_requirement]
 * @property {*} [phone_number_requirement]
 * @property {string} [platform]
 * @property {boolean} [registration_open]
 * @property {Array} [registration_question]
 * @property {boolean} [reminders_disabled]
 * @property {boolean} [require_approval]
 * @property {boolean} [show_guest_list]
 * @property {string} [slug]
 * @property {*} [spots_remaining]
 * @property {string} [start_at]
 * @property {boolean} [suppress_notification]
 * @property {string} [timezone]
 * @property {string} [tint_color]
 * @property {string} [url]
 * @property {string} [user_id]
 * @property {string} [visibility]
 * @property {string} [waitlist_status]
 */

/**
 * @typedef {Object} EventCreateData
 * @property {string} access
 * @property {string} calendar_id
 * @property {boolean} [can_register_for_multiple_ticket]
 * @property {*} coordinate
 * @property {string} cover_url
 * @property {string} created_at
 * @property {string} description
 * @property {string} description_md
 * @property {*} display_price
 * @property {string} duration_interval
 * @property {string} end_at
 * @property {string} event_id
 * @property {Object} feedback_email
 * @property {*} geo_address_json
 * @property {Object} guest_count
 * @property {Array} host
 * @property {string} id
 * @property {string} location_type
 * @property {string} location_visibility
 * @property {*} [max_capacity]
 * @property {*} meeting_url
 * @property {string} name
 * @property {string} [name_requirement]
 * @property {*} [phone_number_requirement]
 * @property {string} platform
 * @property {boolean} registration_open
 * @property {Array} [registration_question]
 * @property {boolean} [reminders_disabled]
 * @property {boolean} require_approval
 * @property {boolean} [show_guest_list]
 * @property {string} [slug]
 * @property {*} spots_remaining
 * @property {string} start_at
 * @property {boolean} [suppress_notification]
 * @property {string} timezone
 * @property {string} [tint_color]
 * @property {string} url
 * @property {string} user_id
 * @property {string} visibility
 * @property {string} waitlist_status
 */

/**
 * @typedef {Object} EventUpdateData
 * @property {string} [access]
 * @property {string} [calendar_id]
 * @property {boolean} [can_register_for_multiple_ticket]
 * @property {*} [coordinate]
 * @property {string} [cover_url]
 * @property {string} [created_at]
 * @property {string} [description]
 * @property {string} [description_md]
 * @property {*} [display_price]
 * @property {string} [duration_interval]
 * @property {string} [end_at]
 * @property {string} [event_id]
 * @property {Object} [feedback_email]
 * @property {*} [geo_address_json]
 * @property {Object} [guest_count]
 * @property {Array} [host]
 * @property {string} [id]
 * @property {string} [location_type]
 * @property {string} [location_visibility]
 * @property {*} [max_capacity]
 * @property {*} [meeting_url]
 * @property {string} [name]
 * @property {string} [name_requirement]
 * @property {*} [phone_number_requirement]
 * @property {string} [platform]
 * @property {boolean} [registration_open]
 * @property {Array} [registration_question]
 * @property {boolean} [reminders_disabled]
 * @property {boolean} [require_approval]
 * @property {boolean} [show_guest_list]
 * @property {string} [slug]
 * @property {*} [spots_remaining]
 * @property {string} [start_at]
 * @property {boolean} [suppress_notification]
 * @property {string} [timezone]
 * @property {string} [tint_color]
 * @property {string} [url]
 * @property {string} [user_id]
 * @property {string} [visibility]
 * @property {string} [waitlist_status]
 */

/**
 * @typedef {Object} EventRemoveMatch
 * @property {string} [access]
 * @property {string} [calendar_id]
 * @property {boolean} [can_register_for_multiple_ticket]
 * @property {*} [coordinate]
 * @property {string} [cover_url]
 * @property {string} [created_at]
 * @property {string} [description]
 * @property {string} [description_md]
 * @property {*} [display_price]
 * @property {string} [duration_interval]
 * @property {string} [end_at]
 * @property {string} [event_id]
 * @property {Object} [feedback_email]
 * @property {*} [geo_address_json]
 * @property {Object} [guest_count]
 * @property {Array} [host]
 * @property {string} id
 * @property {string} [location_type]
 * @property {string} [location_visibility]
 * @property {*} [max_capacity]
 * @property {*} [meeting_url]
 * @property {string} [name]
 * @property {string} [name_requirement]
 * @property {*} [phone_number_requirement]
 * @property {string} [platform]
 * @property {boolean} [registration_open]
 * @property {Array} [registration_question]
 * @property {boolean} [reminders_disabled]
 * @property {boolean} [require_approval]
 * @property {boolean} [show_guest_list]
 * @property {string} [slug]
 * @property {*} [spots_remaining]
 * @property {string} [start_at]
 * @property {boolean} [suppress_notification]
 * @property {string} [timezone]
 * @property {string} [tint_color]
 * @property {string} [url]
 * @property {string} [user_id]
 * @property {string} [visibility]
 * @property {string} [waitlist_status]
 */

/**
 * @typedef {Object} EventCancelRequest
 * @property {string} cancellation_token
 * @property {string} event_id
 * @property {number} guest_count
 * @property {boolean} is_paid
 */

/**
 * @typedef {Object} EventCancelRequestCreateData
 * @property {string} cancellation_token
 * @property {string} event_id
 * @property {number} guest_count
 * @property {boolean} is_paid
 */

/**
 * @typedef {Object} EventCoupon
 * @property {*} cents_off
 * @property {string} code
 * @property {*} currency
 * @property {*} discount
 * @property {string} event_id
 * @property {string} [event_ticket_type_id]
 * @property {string} id
 * @property {*} percent_off
 * @property {number} remaining_count
 * @property {*} valid_end_at
 * @property {*} valid_start_at
 */

/**
 * @typedef {Object} EventCouponListMatch
 * @property {*} [cents_off]
 * @property {string} [code]
 * @property {*} [currency]
 * @property {*} [discount]
 * @property {string} [event_id]
 * @property {string} [event_ticket_type_id]
 * @property {string} [id]
 * @property {*} [percent_off]
 * @property {number} [remaining_count]
 * @property {*} [valid_end_at]
 * @property {*} [valid_start_at]
 */

/**
 * @typedef {Object} EventCouponCreateData
 * @property {*} cents_off
 * @property {string} code
 * @property {*} currency
 * @property {*} discount
 * @property {string} event_id
 * @property {string} [event_ticket_type_id]
 * @property {string} id
 * @property {*} percent_off
 * @property {number} remaining_count
 * @property {*} valid_end_at
 * @property {*} valid_start_at
 */

/**
 * @typedef {Object} EventCouponUpdateData
 * @property {*} [cents_off]
 * @property {string} [code]
 * @property {*} [currency]
 * @property {*} [discount]
 * @property {string} [event_id]
 * @property {string} [event_ticket_type_id]
 * @property {string} [id]
 * @property {*} [percent_off]
 * @property {number} [remaining_count]
 * @property {*} [valid_end_at]
 * @property {*} [valid_start_at]
 */

/**
 * @typedef {Object} EventTag
 * @property {*} [color]
 * @property {string} id
 * @property {string} name
 * @property {string} tag_id
 */

/**
 * @typedef {Object} EventTagListMatch
 * @property {*} [color]
 * @property {string} [id]
 * @property {string} [name]
 * @property {string} [tag_id]
 */

/**
 * @typedef {Object} EventTagCreateData
 * @property {*} [color]
 * @property {string} id
 * @property {string} name
 * @property {string} tag_id
 */

/**
 * @typedef {Object} EventTagUpdateData
 * @property {*} [color]
 * @property {string} [id]
 * @property {string} [name]
 * @property {string} [tag_id]
 */

/**
 * @typedef {Object} EventTagRemoveMatch
 * @property {*} [color]
 * @property {string} id
 * @property {string} [name]
 * @property {string} [tag_id]
 */

/**
 * @typedef {Object} EventTagAssignment
 * @property {number} applied_count
 * @property {Array} event_id
 * @property {number} skipped_count
 * @property {string} tag
 */

/**
 * @typedef {Object} EventTagAssignmentCreateData
 * @property {number} applied_count
 * @property {Array} event_id
 * @property {number} skipped_count
 * @property {string} tag
 */

/**
 * @typedef {Object} EventTagAssignmentRemoveMatch
 * @property {number} [applied_count]
 * @property {Array} [event_id]
 * @property {number} [skipped_count]
 * @property {string} [tag]
 */

/**
 * @typedef {Object} Guest
 * @property {string} approval_status
 * @property {string} check_in_qr_code
 * @property {*} eth_address
 * @property {string} event_id
 * @property {Array} event_ticket
 * @property {Array} event_ticket_order
 * @property {Array} guest
 * @property {string} guest_id
 * @property {string} id
 * @property {*} invited_at
 * @property {*} joined_at
 * @property {*} [message]
 * @property {number} phone_number
 * @property {*} registered_at
 * @property {*} registration_answer
 * @property {*} [send_email]
 * @property {boolean} [should_refund]
 * @property {*} solana_address
 * @property {string} status
 * @property {*} [ticket]
 * @property {string} user_email
 * @property {*} user_first_name
 * @property {string} user_id
 * @property {*} user_last_name
 * @property {*} user_name
 * @property {*} utm_source
 */

/**
 * @typedef {Object} GuestLoadMatch
 * @property {string} [approval_status]
 * @property {string} [check_in_qr_code]
 * @property {*} [eth_address]
 * @property {string} [event_id]
 * @property {Array} [event_ticket]
 * @property {Array} [event_ticket_order]
 * @property {Array} [guest]
 * @property {string} [guest_id]
 * @property {string} id
 * @property {*} [invited_at]
 * @property {*} [joined_at]
 * @property {*} [message]
 * @property {number} [phone_number]
 * @property {*} [registered_at]
 * @property {*} [registration_answer]
 * @property {*} [send_email]
 * @property {boolean} [should_refund]
 * @property {*} [solana_address]
 * @property {string} [status]
 * @property {*} [ticket]
 * @property {string} [user_email]
 * @property {*} [user_first_name]
 * @property {string} [user_id]
 * @property {*} [user_last_name]
 * @property {*} [user_name]
 * @property {*} [utm_source]
 */

/**
 * @typedef {Object} GuestListMatch
 * @property {string} [approval_status]
 * @property {string} [check_in_qr_code]
 * @property {*} [eth_address]
 * @property {string} [event_id]
 * @property {Array} [event_ticket]
 * @property {Array} [event_ticket_order]
 * @property {Array} [guest]
 * @property {string} [guest_id]
 * @property {string} [id]
 * @property {*} [invited_at]
 * @property {*} [joined_at]
 * @property {*} [message]
 * @property {number} [phone_number]
 * @property {*} [registered_at]
 * @property {*} [registration_answer]
 * @property {*} [send_email]
 * @property {boolean} [should_refund]
 * @property {*} [solana_address]
 * @property {string} [status]
 * @property {*} [ticket]
 * @property {string} [user_email]
 * @property {*} [user_first_name]
 * @property {string} [user_id]
 * @property {*} [user_last_name]
 * @property {*} [user_name]
 * @property {*} [utm_source]
 */

/**
 * @typedef {Object} GuestCreateData
 * @property {string} approval_status
 * @property {string} check_in_qr_code
 * @property {*} eth_address
 * @property {string} event_id
 * @property {Array} event_ticket
 * @property {Array} event_ticket_order
 * @property {Array} guest
 * @property {string} guest_id
 * @property {string} id
 * @property {*} invited_at
 * @property {*} joined_at
 * @property {*} [message]
 * @property {number} phone_number
 * @property {*} registered_at
 * @property {*} registration_answer
 * @property {*} [send_email]
 * @property {boolean} [should_refund]
 * @property {*} solana_address
 * @property {string} status
 * @property {*} [ticket]
 * @property {string} user_email
 * @property {*} user_first_name
 * @property {string} user_id
 * @property {*} user_last_name
 * @property {*} user_name
 * @property {*} utm_source
 */

/**
 * @typedef {Object} GuestUpdateData
 * @property {string} [approval_status]
 * @property {string} [check_in_qr_code]
 * @property {*} [eth_address]
 * @property {string} [event_id]
 * @property {Array} [event_ticket]
 * @property {Array} [event_ticket_order]
 * @property {Array} [guest]
 * @property {string} [guest_id]
 * @property {string} [id]
 * @property {*} [invited_at]
 * @property {*} [joined_at]
 * @property {*} [message]
 * @property {number} [phone_number]
 * @property {*} [registered_at]
 * @property {*} [registration_answer]
 * @property {*} [send_email]
 * @property {boolean} [should_refund]
 * @property {*} [solana_address]
 * @property {string} [status]
 * @property {*} [ticket]
 * @property {string} [user_email]
 * @property {*} [user_first_name]
 * @property {string} [user_id]
 * @property {*} [user_last_name]
 * @property {*} [user_name]
 * @property {*} [utm_source]
 */

/**
 * @typedef {Object} GuestInvite
 * @property {string} event_id
 * @property {Array} guest
 * @property {*} [message]
 */

/**
 * @typedef {Object} GuestInviteCreateData
 * @property {string} event_id
 * @property {Array} guest
 * @property {*} [message]
 */

/**
 * @typedef {Object} GuestTicket
 * @property {string} event_id
 * @property {string} guest_id
 * @property {*} [send_email]
 * @property {Array} [ticket_ids_to_remove]
 * @property {Array} [tickets_to_add]
 */

/**
 * @typedef {Object} GuestTicketUpdateData
 * @property {string} [event_id]
 * @property {string} [guest_id]
 * @property {*} [send_email]
 * @property {Array} [ticket_ids_to_remove]
 * @property {Array} [tickets_to_add]
 */

/**
 * @typedef {Object} Host
 * @property {*} [access_level]
 * @property {string} email
 * @property {string} event_id
 * @property {boolean} [is_visible]
 * @property {string} [name]
 */

/**
 * @typedef {Object} HostCreateData
 * @property {*} [access_level]
 * @property {string} email
 * @property {string} event_id
 * @property {boolean} [is_visible]
 * @property {string} [name]
 */

/**
 * @typedef {Object} HostUpdateData
 * @property {*} [access_level]
 * @property {string} [email]
 * @property {string} [event_id]
 * @property {boolean} [is_visible]
 * @property {string} [name]
 */

/**
 * @typedef {Object} HostRemoveMatch
 * @property {*} [access_level]
 * @property {string} [email]
 * @property {string} [event_id]
 * @property {boolean} [is_visible]
 * @property {string} [name]
 */

/**
 * @typedef {Object} ImageUpload
 * @property {*} [content_type]
 * @property {string} file_url
 * @property {string} upload_url
 */

/**
 * @typedef {Object} ImageUploadCreateData
 * @property {*} [content_type]
 * @property {string} file_url
 * @property {string} upload_url
 */

/**
 * @typedef {Object} Member
 * @property {string} email
 * @property {string} membership_id
 * @property {string} membership_tier_id
 * @property {Array} [registration_answer]
 * @property {boolean} [skip_payment]
 * @property {string} status
 * @property {string} user_id
 */

/**
 * @typedef {Object} MemberCreateData
 * @property {string} email
 * @property {string} membership_id
 * @property {string} membership_tier_id
 * @property {Array} [registration_answer]
 * @property {boolean} [skip_payment]
 * @property {string} status
 * @property {string} user_id
 */

/**
 * @typedef {Object} MemberUpdateData
 * @property {string} [email]
 * @property {string} [membership_id]
 * @property {string} [membership_tier_id]
 * @property {Array} [registration_answer]
 * @property {boolean} [skip_payment]
 * @property {string} [status]
 * @property {string} [user_id]
 */

/**
 * @typedef {Object} MembershipTier
 * @property {*} access_info
 * @property {string} description
 * @property {string} id
 * @property {string} name
 * @property {string} tint_color
 */

/**
 * @typedef {Object} MembershipTierListMatch
 * @property {*} [access_info]
 * @property {string} [description]
 * @property {string} [id]
 * @property {string} [name]
 * @property {string} [tint_color]
 */

/**
 * @typedef {Object} OrganizationAdmin
 * @property {string} api_id
 * @property {string} avatar_url
 * @property {string} email
 * @property {*} first_name
 * @property {string} id
 * @property {*} last_name
 * @property {string} name
 */

/**
 * @typedef {Object} OrganizationAdminListMatch
 * @property {string} [api_id]
 * @property {string} [avatar_url]
 * @property {string} [email]
 * @property {*} [first_name]
 * @property {string} [id]
 * @property {*} [last_name]
 * @property {string} [name]
 */

/**
 * @typedef {Object} OrganizationCalendar
 * @property {*} avatar_url
 * @property {*} coordinate
 * @property {*} cover_image_url
 * @property {string} description
 * @property {string} id
 * @property {*} instagram_handle
 * @property {boolean} is_personal
 * @property {*} location
 * @property {string} name
 * @property {string} slug
 * @property {*} social_image_url
 * @property {string} [tint_color]
 * @property {*} twitter_handle
 * @property {string} url
 * @property {*} website
 * @property {*} youtube_handle
 */

/**
 * @typedef {Object} OrganizationCalendarListMatch
 * @property {*} [avatar_url]
 * @property {*} [coordinate]
 * @property {*} [cover_image_url]
 * @property {string} [description]
 * @property {string} [id]
 * @property {*} [instagram_handle]
 * @property {boolean} [is_personal]
 * @property {*} [location]
 * @property {string} [name]
 * @property {string} [slug]
 * @property {*} [social_image_url]
 * @property {string} [tint_color]
 * @property {*} [twitter_handle]
 * @property {string} [url]
 * @property {*} [website]
 * @property {*} [youtube_handle]
 */

/**
 * @typedef {Object} OrganizationCalendarCreateData
 * @property {*} avatar_url
 * @property {*} coordinate
 * @property {*} cover_image_url
 * @property {string} description
 * @property {string} id
 * @property {*} instagram_handle
 * @property {boolean} is_personal
 * @property {*} location
 * @property {string} name
 * @property {string} slug
 * @property {*} social_image_url
 * @property {string} [tint_color]
 * @property {*} twitter_handle
 * @property {string} url
 * @property {*} website
 * @property {*} youtube_handle
 */

/**
 * @typedef {Object} OrganizationEvent
 * @property {string} api_id
 * @property {string} calendar_api_id
 * @property {string} calendar_id
 * @property {*} coordinate
 * @property {string} cover_url
 * @property {string} created_at
 * @property {*} display_price
 * @property {string} duration_interval
 * @property {string} end_at
 * @property {Object} feedback_email
 * @property {*} geo_address_json
 * @property {*} geo_latitude
 * @property {*} geo_longitude
 * @property {string} id
 * @property {string} location_type
 * @property {string} location_visibility
 * @property {Array} managing_calendar
 * @property {*} meeting_url
 * @property {string} name
 * @property {string} platform
 * @property {boolean} registration_open
 * @property {Array} [registration_question]
 * @property {boolean} require_approval
 * @property {*} spots_remaining
 * @property {string} start_at
 * @property {string} timezone
 * @property {string} url
 * @property {string} user_api_id
 * @property {string} user_id
 * @property {string} visibility
 * @property {string} waitlist_status
 * @property {*} zoom_meeting_url
 */

/**
 * @typedef {Object} OrganizationEventListMatch
 * @property {string} [api_id]
 * @property {string} [calendar_api_id]
 * @property {string} [calendar_id]
 * @property {*} [coordinate]
 * @property {string} [cover_url]
 * @property {string} [created_at]
 * @property {*} [display_price]
 * @property {string} [duration_interval]
 * @property {string} [end_at]
 * @property {Object} [feedback_email]
 * @property {*} [geo_address_json]
 * @property {*} [geo_latitude]
 * @property {*} [geo_longitude]
 * @property {string} [id]
 * @property {string} [location_type]
 * @property {string} [location_visibility]
 * @property {Array} [managing_calendar]
 * @property {*} [meeting_url]
 * @property {string} [name]
 * @property {string} [platform]
 * @property {boolean} [registration_open]
 * @property {Array} [registration_question]
 * @property {boolean} [require_approval]
 * @property {*} [spots_remaining]
 * @property {string} [start_at]
 * @property {string} [timezone]
 * @property {string} [url]
 * @property {string} [user_api_id]
 * @property {string} [user_id]
 * @property {string} [visibility]
 * @property {string} [waitlist_status]
 * @property {*} [zoom_meeting_url]
 */

/**
 * @typedef {Object} OrganizationEventTransfer
 * @property {string} calendar_id
 * @property {string} event_id
 */

/**
 * @typedef {Object} OrganizationEventTransferCreateData
 * @property {string} calendar_id
 * @property {string} event_id
 */

/**
 * @typedef {Object} TicketType
 * @property {*} [cent]
 * @property {*} [currency]
 * @property {string} [description]
 * @property {string} id
 * @property {boolean} [is_flexible]
 * @property {boolean} [is_hidden]
 * @property {*} [max_capacity]
 * @property {*} [min_cent]
 * @property {string} name
 * @property {boolean} [require_approval]
 * @property {string} type
 * @property {*} [valid_end_at]
 * @property {*} [valid_start_at]
 */

/**
 * @typedef {Object} TicketTypeLoadMatch
 * @property {*} [cent]
 * @property {*} [currency]
 * @property {string} [description]
 * @property {string} id
 * @property {boolean} [is_flexible]
 * @property {boolean} [is_hidden]
 * @property {*} [max_capacity]
 * @property {*} [min_cent]
 * @property {string} [name]
 * @property {boolean} [require_approval]
 * @property {string} [type]
 * @property {*} [valid_end_at]
 * @property {*} [valid_start_at]
 */

/**
 * @typedef {Object} TicketTypeListMatch
 * @property {*} [cent]
 * @property {*} [currency]
 * @property {string} [description]
 * @property {string} [id]
 * @property {boolean} [is_flexible]
 * @property {boolean} [is_hidden]
 * @property {*} [max_capacity]
 * @property {*} [min_cent]
 * @property {string} [name]
 * @property {boolean} [require_approval]
 * @property {string} [type]
 * @property {*} [valid_end_at]
 * @property {*} [valid_start_at]
 */

/**
 * @typedef {Object} TicketTypeCreateData
 * @property {*} [cent]
 * @property {*} [currency]
 * @property {string} [description]
 * @property {string} id
 * @property {boolean} [is_flexible]
 * @property {boolean} [is_hidden]
 * @property {*} [max_capacity]
 * @property {*} [min_cent]
 * @property {string} name
 * @property {boolean} [require_approval]
 * @property {string} type
 * @property {*} [valid_end_at]
 * @property {*} [valid_start_at]
 */

/**
 * @typedef {Object} TicketTypeUpdateData
 * @property {*} [cent]
 * @property {*} [currency]
 * @property {string} [description]
 * @property {string} [id]
 * @property {boolean} [is_flexible]
 * @property {boolean} [is_hidden]
 * @property {*} [max_capacity]
 * @property {*} [min_cent]
 * @property {string} [name]
 * @property {boolean} [require_approval]
 * @property {string} [type]
 * @property {*} [valid_end_at]
 * @property {*} [valid_start_at]
 */

/**
 * @typedef {Object} TicketTypeRemoveMatch
 * @property {*} [cent]
 * @property {*} [currency]
 * @property {string} [description]
 * @property {string} id
 * @property {boolean} [is_flexible]
 * @property {boolean} [is_hidden]
 * @property {*} [max_capacity]
 * @property {*} [min_cent]
 * @property {string} [name]
 * @property {boolean} [require_approval]
 * @property {string} [type]
 * @property {*} [valid_end_at]
 * @property {*} [valid_start_at]
 */

/**
 * @typedef {Object} User
 * @property {string} avatar_url
 * @property {string} email
 * @property {*} first_name
 * @property {string} id
 * @property {*} last_name
 * @property {string} name
 */

/**
 * @typedef {Object} UserLoadMatch
 * @property {string} [avatar_url]
 * @property {string} [email]
 * @property {*} [first_name]
 * @property {string} id
 * @property {*} [last_name]
 * @property {string} [name]
 */

/**
 * @typedef {Object} Webhook
 * @property {string} created_at
 * @property {Array} event_type
 * @property {string} id
 * @property {string} secret
 * @property {string} status
 * @property {string} url
 */

/**
 * @typedef {Object} WebhookLoadMatch
 * @property {string} [created_at]
 * @property {Array} [event_type]
 * @property {string} id
 * @property {string} [secret]
 * @property {string} [status]
 * @property {string} [url]
 */

/**
 * @typedef {Object} WebhookListMatch
 * @property {string} [created_at]
 * @property {Array} [event_type]
 * @property {string} [id]
 * @property {string} [secret]
 * @property {string} [status]
 * @property {string} [url]
 */

/**
 * @typedef {Object} WebhookCreateData
 * @property {string} created_at
 * @property {Array} event_type
 * @property {string} id
 * @property {string} secret
 * @property {string} status
 * @property {string} url
 */

/**
 * @typedef {Object} WebhookUpdateData
 * @property {string} [created_at]
 * @property {Array} [event_type]
 * @property {string} [id]
 * @property {string} [secret]
 * @property {string} [status]
 * @property {string} [url]
 */

/**
 * @typedef {Object} WebhookRemoveMatch
 * @property {string} [created_at]
 * @property {Array} [event_type]
 * @property {string} id
 * @property {string} [secret]
 * @property {string} [status]
 * @property {string} [url]
 */

