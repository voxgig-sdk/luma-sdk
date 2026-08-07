// Typed models for the Luma SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import "encoding/json"

// Calendar is the typed data model for the calendar entity.
type Calendar struct {
	AvatarUrl any `json:"avatar_url"`
	CalendarId string `json:"calendar_id"`
	Coordinate any `json:"coordinate"`
	CoverImageUrl any `json:"cover_image_url"`
	Description string `json:"description"`
	Id string `json:"id"`
	InstagramHandle any `json:"instagram_handle"`
	IsPersonal bool `json:"is_personal"`
	Location any `json:"location"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	SocialImageUrl any `json:"social_image_url"`
	TintColor *string `json:"tint_color,omitempty"`
	TwitterHandle any `json:"twitter_handle"`
	Url string `json:"url"`
	Website any `json:"website"`
	YoutubeHandle any `json:"youtube_handle"`
}

// CalendarLoadMatch is the typed request payload for Calendar.LoadTyped.
type CalendarLoadMatch struct {
	AvatarUrl *any `json:"avatar_url,omitempty"`
	CalendarId *string `json:"calendar_id,omitempty"`
	Coordinate *any `json:"coordinate,omitempty"`
	CoverImageUrl *any `json:"cover_image_url,omitempty"`
	Description *string `json:"description,omitempty"`
	Id string `json:"id"`
	InstagramHandle *any `json:"instagram_handle,omitempty"`
	IsPersonal *bool `json:"is_personal,omitempty"`
	Location *any `json:"location,omitempty"`
	Name *string `json:"name,omitempty"`
	Slug *string `json:"slug,omitempty"`
	SocialImageUrl *any `json:"social_image_url,omitempty"`
	TintColor *string `json:"tint_color,omitempty"`
	TwitterHandle *any `json:"twitter_handle,omitempty"`
	Url *string `json:"url,omitempty"`
	Website *any `json:"website,omitempty"`
	YoutubeHandle *any `json:"youtube_handle,omitempty"`
}

// CalendarUpdateData is the typed request payload for Calendar.UpdateTyped.
type CalendarUpdateData struct {
	AvatarUrl *any `json:"avatar_url,omitempty"`
	CalendarId *string `json:"calendar_id,omitempty"`
	Coordinate *any `json:"coordinate,omitempty"`
	CoverImageUrl *any `json:"cover_image_url,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	InstagramHandle *any `json:"instagram_handle,omitempty"`
	IsPersonal *bool `json:"is_personal,omitempty"`
	Location *any `json:"location,omitempty"`
	Name *string `json:"name,omitempty"`
	Slug *string `json:"slug,omitempty"`
	SocialImageUrl *any `json:"social_image_url,omitempty"`
	TintColor *string `json:"tint_color,omitempty"`
	TwitterHandle *any `json:"twitter_handle,omitempty"`
	Url *string `json:"url,omitempty"`
	Website *any `json:"website,omitempty"`
	YoutubeHandle *any `json:"youtube_handle,omitempty"`
}

// CalendarAdmin is the typed data model for the calendar_admin entity.
type CalendarAdmin struct {
	AvatarUrl string `json:"avatar_url"`
	Email string `json:"email"`
	FirstName any `json:"first_name"`
	Id string `json:"id"`
	LastName any `json:"last_name"`
	Name string `json:"name"`
}

// CalendarAdminListMatch is the typed request payload for CalendarAdmin.ListTyped.
type CalendarAdminListMatch struct {
	AvatarUrl *string `json:"avatar_url,omitempty"`
	Email *string `json:"email,omitempty"`
	FirstName *any `json:"first_name,omitempty"`
	Id *string `json:"id,omitempty"`
	LastName *any `json:"last_name,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CalendarCoupon is the typed data model for the calendar_coupon entity.
type CalendarCoupon struct {
	CentsOff any `json:"cents_off"`
	Code string `json:"code"`
	Currency any `json:"currency"`
	Discount any `json:"discount"`
	EventTicketTypeId *string `json:"event_ticket_type_id,omitempty"`
	Id string `json:"id"`
	PercentOff any `json:"percent_off"`
	RemainingCount int `json:"remaining_count"`
	ValidEndAt any `json:"valid_end_at"`
	ValidStartAt any `json:"valid_start_at"`
}

// CalendarCouponListMatch is the typed request payload for CalendarCoupon.ListTyped.
type CalendarCouponListMatch struct {
	CentsOff *any `json:"cents_off,omitempty"`
	Code *string `json:"code,omitempty"`
	Currency *any `json:"currency,omitempty"`
	Discount *any `json:"discount,omitempty"`
	EventTicketTypeId *string `json:"event_ticket_type_id,omitempty"`
	Id *string `json:"id,omitempty"`
	PercentOff *any `json:"percent_off,omitempty"`
	RemainingCount *int `json:"remaining_count,omitempty"`
	ValidEndAt *any `json:"valid_end_at,omitempty"`
	ValidStartAt *any `json:"valid_start_at,omitempty"`
}

// CalendarCouponCreateData is the typed request payload for CalendarCoupon.CreateTyped.
type CalendarCouponCreateData struct {
	CentsOff any `json:"cents_off"`
	Code string `json:"code"`
	Currency any `json:"currency"`
	Discount any `json:"discount"`
	EventTicketTypeId *string `json:"event_ticket_type_id,omitempty"`
	Id string `json:"id"`
	PercentOff any `json:"percent_off"`
	RemainingCount int `json:"remaining_count"`
	ValidEndAt any `json:"valid_end_at"`
	ValidStartAt any `json:"valid_start_at"`
}

// CalendarCouponUpdateData is the typed request payload for CalendarCoupon.UpdateTyped.
type CalendarCouponUpdateData struct {
	CentsOff *any `json:"cents_off,omitempty"`
	Code *string `json:"code,omitempty"`
	Currency *any `json:"currency,omitempty"`
	Discount *any `json:"discount,omitempty"`
	EventTicketTypeId *string `json:"event_ticket_type_id,omitempty"`
	Id *string `json:"id,omitempty"`
	PercentOff *any `json:"percent_off,omitempty"`
	RemainingCount *int `json:"remaining_count,omitempty"`
	ValidEndAt *any `json:"valid_end_at,omitempty"`
	ValidStartAt *any `json:"valid_start_at,omitempty"`
}

// CalendarEvent is the typed data model for the calendar_event entity.
type CalendarEvent struct {
	Id string `json:"id"`
	Status string `json:"status"`
	SubmittedBy any `json:"submitted_by"`
	Tag []any `json:"tag"`
}

// CalendarEventLoadMatch is the typed request payload for CalendarEvent.LoadTyped.
type CalendarEventLoadMatch struct {
	Id string `json:"id"`
	Status *string `json:"status,omitempty"`
	SubmittedBy *any `json:"submitted_by,omitempty"`
	Tag *[]any `json:"tag,omitempty"`
}

// CalendarEventListMatch is the typed request payload for CalendarEvent.ListTyped.
type CalendarEventListMatch struct {
	Id *string `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
	SubmittedBy *any `json:"submitted_by,omitempty"`
	Tag *[]any `json:"tag,omitempty"`
}

// CalendarEventCreateData is the typed request payload for CalendarEvent.CreateTyped.
type CalendarEventCreateData struct {
	Id string `json:"id"`
	Status string `json:"status"`
	SubmittedBy any `json:"submitted_by"`
	Tag []any `json:"tag"`
}

// CalendarEventApproval is the typed data model for the calendar_event_approval entity.
type CalendarEventApproval struct {
	CalendarEventId string `json:"calendar_event_id"`
}

// CalendarEventApprovalCreateData is the typed request payload for CalendarEventApproval.CreateTyped.
type CalendarEventApprovalCreateData struct {
	CalendarEventId string `json:"calendar_event_id"`
}

// CalendarEventRejection is the typed data model for the calendar_event_rejection entity.
type CalendarEventRejection struct {
	CalendarEventId string `json:"calendar_event_id"`
	Message *string `json:"message,omitempty"`
}

// CalendarEventRejectionCreateData is the typed request payload for CalendarEventRejection.CreateTyped.
type CalendarEventRejectionCreateData struct {
	CalendarEventId string `json:"calendar_event_id"`
	Message *string `json:"message,omitempty"`
}

// Contact is the typed data model for the contact entity.
type Contact struct {
	AvatarUrl string `json:"avatar_url"`
	Contact []any `json:"contact"`
	CreatedAt string `json:"created_at"`
	Email string `json:"email"`
	EventApprovedCount float64 `json:"event_approved_count"`
	EventCheckedInCount float64 `json:"event_checked_in_count"`
	FirstName any `json:"first_name"`
	Id string `json:"id"`
	LastName any `json:"last_name"`
	Membership any `json:"membership"`
	Name string `json:"name"`
	RevenueUsdCent float64 `json:"revenue_usd_cent"`
	Tag *any `json:"tag,omitempty"`
	UserId string `json:"user_id"`
}

// ContactListMatch is the typed request payload for Contact.ListTyped.
type ContactListMatch struct {
	AvatarUrl *string `json:"avatar_url,omitempty"`
	Contact *[]any `json:"contact,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Email *string `json:"email,omitempty"`
	EventApprovedCount *float64 `json:"event_approved_count,omitempty"`
	EventCheckedInCount *float64 `json:"event_checked_in_count,omitempty"`
	FirstName *any `json:"first_name,omitempty"`
	Id *string `json:"id,omitempty"`
	LastName *any `json:"last_name,omitempty"`
	Membership *any `json:"membership,omitempty"`
	Name *string `json:"name,omitempty"`
	RevenueUsdCent *float64 `json:"revenue_usd_cent,omitempty"`
	Tag *any `json:"tag,omitempty"`
	UserId *string `json:"user_id,omitempty"`
}

// ContactCreateData is the typed request payload for Contact.CreateTyped.
type ContactCreateData struct {
	AvatarUrl string `json:"avatar_url"`
	Contact []any `json:"contact"`
	CreatedAt string `json:"created_at"`
	Email string `json:"email"`
	EventApprovedCount float64 `json:"event_approved_count"`
	EventCheckedInCount float64 `json:"event_checked_in_count"`
	FirstName any `json:"first_name"`
	Id string `json:"id"`
	LastName any `json:"last_name"`
	Membership any `json:"membership"`
	Name string `json:"name"`
	RevenueUsdCent float64 `json:"revenue_usd_cent"`
	Tag *any `json:"tag,omitempty"`
	UserId string `json:"user_id"`
}

// ContactRemoveMatch is the typed request payload for Contact.RemoveTyped.
type ContactRemoveMatch struct {
	AvatarUrl *string `json:"avatar_url,omitempty"`
	Contact *[]any `json:"contact,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Email *string `json:"email,omitempty"`
	EventApprovedCount *float64 `json:"event_approved_count,omitempty"`
	EventCheckedInCount *float64 `json:"event_checked_in_count,omitempty"`
	FirstName *any `json:"first_name,omitempty"`
	Id string `json:"id"`
	LastName *any `json:"last_name,omitempty"`
	Membership *any `json:"membership,omitempty"`
	Name *string `json:"name,omitempty"`
	RevenueUsdCent *float64 `json:"revenue_usd_cent,omitempty"`
	Tag *any `json:"tag,omitempty"`
	UserId *string `json:"user_id,omitempty"`
}

// ContactBlock is the typed data model for the contact_block entity.
type ContactBlock struct {
	ContactId *string `json:"contact_id,omitempty"`
	Email *string `json:"email,omitempty"`
}

// ContactBlockCreateData is the typed request payload for ContactBlock.CreateTyped.
type ContactBlockCreateData struct {
	ContactId *string `json:"contact_id,omitempty"`
	Email *string `json:"email,omitempty"`
}

// ContactRestore is the typed data model for the contact_restore entity.
type ContactRestore struct {
	ContactId *string `json:"contact_id,omitempty"`
	Email *string `json:"email,omitempty"`
}

// ContactRestoreCreateData is the typed request payload for ContactRestore.CreateTyped.
type ContactRestoreCreateData struct {
	ContactId *string `json:"contact_id,omitempty"`
	Email *string `json:"email,omitempty"`
}

// ContactTag is the typed data model for the contact_tag entity.
type ContactTag struct {
	Color *any `json:"color,omitempty"`
	Id string `json:"id"`
	Name string `json:"name"`
	TagId string `json:"tag_id"`
}

// ContactTagListMatch is the typed request payload for ContactTag.ListTyped.
type ContactTagListMatch struct {
	Color *any `json:"color,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	TagId *string `json:"tag_id,omitempty"`
}

// ContactTagCreateData is the typed request payload for ContactTag.CreateTyped.
type ContactTagCreateData struct {
	Color *any `json:"color,omitempty"`
	Id string `json:"id"`
	Name string `json:"name"`
	TagId string `json:"tag_id"`
}

// ContactTagUpdateData is the typed request payload for ContactTag.UpdateTyped.
type ContactTagUpdateData struct {
	Color *any `json:"color,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	TagId *string `json:"tag_id,omitempty"`
}

// ContactTagRemoveMatch is the typed request payload for ContactTag.RemoveTyped.
type ContactTagRemoveMatch struct {
	Color *any `json:"color,omitempty"`
	Id string `json:"id"`
	Name *string `json:"name,omitempty"`
	TagId *string `json:"tag_id,omitempty"`
}

// ContactTagAssignment is the typed data model for the contact_tag_assignment entity.
type ContactTagAssignment struct {
	AppliedCount float64 `json:"applied_count"`
	Email *[]any `json:"email,omitempty"`
	SkippedCount float64 `json:"skipped_count"`
	Tag string `json:"tag"`
	UserId *[]any `json:"user_id,omitempty"`
}

// ContactTagAssignmentCreateData is the typed request payload for ContactTagAssignment.CreateTyped.
type ContactTagAssignmentCreateData struct {
	AppliedCount float64 `json:"applied_count"`
	Email *[]any `json:"email,omitempty"`
	SkippedCount float64 `json:"skipped_count"`
	Tag string `json:"tag"`
	UserId *[]any `json:"user_id,omitempty"`
}

// ContactTagAssignmentRemoveMatch is the typed request payload for ContactTagAssignment.RemoveTyped.
type ContactTagAssignmentRemoveMatch struct {
	AppliedCount *float64 `json:"applied_count,omitempty"`
	Email *[]any `json:"email,omitempty"`
	SkippedCount *float64 `json:"skipped_count,omitempty"`
	Tag *string `json:"tag,omitempty"`
	UserId *[]any `json:"user_id,omitempty"`
}

// EntityLookup is the typed data model for the entity_lookup entity.
type EntityLookup struct {
}

// EntityLookupLoadMatch is the typed request payload for EntityLookup.LoadTyped.
type EntityLookupLoadMatch struct {
}

// Event is the typed data model for the event entity.
type Event struct {
	Access string `json:"access"`
	CalendarId string `json:"calendar_id"`
	CanRegisterForMultipleTicket *bool `json:"can_register_for_multiple_ticket,omitempty"`
	Coordinate any `json:"coordinate"`
	CoverUrl string `json:"cover_url"`
	CreatedAt string `json:"created_at"`
	Description string `json:"description"`
	DescriptionMd string `json:"description_md"`
	DisplayPrice any `json:"display_price"`
	DurationInterval string `json:"duration_interval"`
	EndAt string `json:"end_at"`
	EventId string `json:"event_id"`
	FeedbackEmail map[string]any `json:"feedback_email"`
	GeoAddressJson any `json:"geo_address_json"`
	GuestCount map[string]any `json:"guest_count"`
	Host []any `json:"host"`
	Id string `json:"id"`
	LocationType string `json:"location_type"`
	LocationVisibility string `json:"location_visibility"`
	MaxCapacity *any `json:"max_capacity,omitempty"`
	MeetingUrl any `json:"meeting_url"`
	Name string `json:"name"`
	NameRequirement *string `json:"name_requirement,omitempty"`
	PhoneNumberRequirement *any `json:"phone_number_requirement,omitempty"`
	Platform string `json:"platform"`
	RegistrationOpen bool `json:"registration_open"`
	RegistrationQuestion *[]any `json:"registration_question,omitempty"`
	RemindersDisabled *bool `json:"reminders_disabled,omitempty"`
	RequireApproval bool `json:"require_approval"`
	ShowGuestList *bool `json:"show_guest_list,omitempty"`
	Slug *string `json:"slug,omitempty"`
	SpotsRemaining any `json:"spots_remaining"`
	StartAt string `json:"start_at"`
	SuppressNotification *bool `json:"suppress_notification,omitempty"`
	Timezone string `json:"timezone"`
	TintColor *string `json:"tint_color,omitempty"`
	Url string `json:"url"`
	UserId string `json:"user_id"`
	Visibility string `json:"visibility"`
	WaitlistStatus string `json:"waitlist_status"`
}

// EventLoadMatch is the typed request payload for Event.LoadTyped.
type EventLoadMatch struct {
	Access *string `json:"access,omitempty"`
	CalendarId *string `json:"calendar_id,omitempty"`
	CanRegisterForMultipleTicket *bool `json:"can_register_for_multiple_ticket,omitempty"`
	Coordinate *any `json:"coordinate,omitempty"`
	CoverUrl *string `json:"cover_url,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Description *string `json:"description,omitempty"`
	DescriptionMd *string `json:"description_md,omitempty"`
	DisplayPrice *any `json:"display_price,omitempty"`
	DurationInterval *string `json:"duration_interval,omitempty"`
	EndAt *string `json:"end_at,omitempty"`
	EventId *string `json:"event_id,omitempty"`
	FeedbackEmail *map[string]any `json:"feedback_email,omitempty"`
	GeoAddressJson *any `json:"geo_address_json,omitempty"`
	GuestCount *map[string]any `json:"guest_count,omitempty"`
	Host *[]any `json:"host,omitempty"`
	Id string `json:"id"`
	LocationType *string `json:"location_type,omitempty"`
	LocationVisibility *string `json:"location_visibility,omitempty"`
	MaxCapacity *any `json:"max_capacity,omitempty"`
	MeetingUrl *any `json:"meeting_url,omitempty"`
	Name *string `json:"name,omitempty"`
	NameRequirement *string `json:"name_requirement,omitempty"`
	PhoneNumberRequirement *any `json:"phone_number_requirement,omitempty"`
	Platform *string `json:"platform,omitempty"`
	RegistrationOpen *bool `json:"registration_open,omitempty"`
	RegistrationQuestion *[]any `json:"registration_question,omitempty"`
	RemindersDisabled *bool `json:"reminders_disabled,omitempty"`
	RequireApproval *bool `json:"require_approval,omitempty"`
	ShowGuestList *bool `json:"show_guest_list,omitempty"`
	Slug *string `json:"slug,omitempty"`
	SpotsRemaining *any `json:"spots_remaining,omitempty"`
	StartAt *string `json:"start_at,omitempty"`
	SuppressNotification *bool `json:"suppress_notification,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
	TintColor *string `json:"tint_color,omitempty"`
	Url *string `json:"url,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
	WaitlistStatus *string `json:"waitlist_status,omitempty"`
}

// EventCreateData is the typed request payload for Event.CreateTyped.
type EventCreateData struct {
	Access string `json:"access"`
	CalendarId string `json:"calendar_id"`
	CanRegisterForMultipleTicket *bool `json:"can_register_for_multiple_ticket,omitempty"`
	Coordinate any `json:"coordinate"`
	CoverUrl string `json:"cover_url"`
	CreatedAt string `json:"created_at"`
	Description string `json:"description"`
	DescriptionMd string `json:"description_md"`
	DisplayPrice any `json:"display_price"`
	DurationInterval string `json:"duration_interval"`
	EndAt string `json:"end_at"`
	EventId string `json:"event_id"`
	FeedbackEmail map[string]any `json:"feedback_email"`
	GeoAddressJson any `json:"geo_address_json"`
	GuestCount map[string]any `json:"guest_count"`
	Host []any `json:"host"`
	Id string `json:"id"`
	LocationType string `json:"location_type"`
	LocationVisibility string `json:"location_visibility"`
	MaxCapacity *any `json:"max_capacity,omitempty"`
	MeetingUrl any `json:"meeting_url"`
	Name string `json:"name"`
	NameRequirement *string `json:"name_requirement,omitempty"`
	PhoneNumberRequirement *any `json:"phone_number_requirement,omitempty"`
	Platform string `json:"platform"`
	RegistrationOpen bool `json:"registration_open"`
	RegistrationQuestion *[]any `json:"registration_question,omitempty"`
	RemindersDisabled *bool `json:"reminders_disabled,omitempty"`
	RequireApproval bool `json:"require_approval"`
	ShowGuestList *bool `json:"show_guest_list,omitempty"`
	Slug *string `json:"slug,omitempty"`
	SpotsRemaining any `json:"spots_remaining"`
	StartAt string `json:"start_at"`
	SuppressNotification *bool `json:"suppress_notification,omitempty"`
	Timezone string `json:"timezone"`
	TintColor *string `json:"tint_color,omitempty"`
	Url string `json:"url"`
	UserId string `json:"user_id"`
	Visibility string `json:"visibility"`
	WaitlistStatus string `json:"waitlist_status"`
}

// EventUpdateData is the typed request payload for Event.UpdateTyped.
type EventUpdateData struct {
	Access *string `json:"access,omitempty"`
	CalendarId *string `json:"calendar_id,omitempty"`
	CanRegisterForMultipleTicket *bool `json:"can_register_for_multiple_ticket,omitempty"`
	Coordinate *any `json:"coordinate,omitempty"`
	CoverUrl *string `json:"cover_url,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Description *string `json:"description,omitempty"`
	DescriptionMd *string `json:"description_md,omitempty"`
	DisplayPrice *any `json:"display_price,omitempty"`
	DurationInterval *string `json:"duration_interval,omitempty"`
	EndAt *string `json:"end_at,omitempty"`
	EventId *string `json:"event_id,omitempty"`
	FeedbackEmail *map[string]any `json:"feedback_email,omitempty"`
	GeoAddressJson *any `json:"geo_address_json,omitempty"`
	GuestCount *map[string]any `json:"guest_count,omitempty"`
	Host *[]any `json:"host,omitempty"`
	Id *string `json:"id,omitempty"`
	LocationType *string `json:"location_type,omitempty"`
	LocationVisibility *string `json:"location_visibility,omitempty"`
	MaxCapacity *any `json:"max_capacity,omitempty"`
	MeetingUrl *any `json:"meeting_url,omitempty"`
	Name *string `json:"name,omitempty"`
	NameRequirement *string `json:"name_requirement,omitempty"`
	PhoneNumberRequirement *any `json:"phone_number_requirement,omitempty"`
	Platform *string `json:"platform,omitempty"`
	RegistrationOpen *bool `json:"registration_open,omitempty"`
	RegistrationQuestion *[]any `json:"registration_question,omitempty"`
	RemindersDisabled *bool `json:"reminders_disabled,omitempty"`
	RequireApproval *bool `json:"require_approval,omitempty"`
	ShowGuestList *bool `json:"show_guest_list,omitempty"`
	Slug *string `json:"slug,omitempty"`
	SpotsRemaining *any `json:"spots_remaining,omitempty"`
	StartAt *string `json:"start_at,omitempty"`
	SuppressNotification *bool `json:"suppress_notification,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
	TintColor *string `json:"tint_color,omitempty"`
	Url *string `json:"url,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
	WaitlistStatus *string `json:"waitlist_status,omitempty"`
}

// EventRemoveMatch is the typed request payload for Event.RemoveTyped.
type EventRemoveMatch struct {
	Access *string `json:"access,omitempty"`
	CalendarId *string `json:"calendar_id,omitempty"`
	CanRegisterForMultipleTicket *bool `json:"can_register_for_multiple_ticket,omitempty"`
	Coordinate *any `json:"coordinate,omitempty"`
	CoverUrl *string `json:"cover_url,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Description *string `json:"description,omitempty"`
	DescriptionMd *string `json:"description_md,omitempty"`
	DisplayPrice *any `json:"display_price,omitempty"`
	DurationInterval *string `json:"duration_interval,omitempty"`
	EndAt *string `json:"end_at,omitempty"`
	EventId *string `json:"event_id,omitempty"`
	FeedbackEmail *map[string]any `json:"feedback_email,omitempty"`
	GeoAddressJson *any `json:"geo_address_json,omitempty"`
	GuestCount *map[string]any `json:"guest_count,omitempty"`
	Host *[]any `json:"host,omitempty"`
	Id string `json:"id"`
	LocationType *string `json:"location_type,omitempty"`
	LocationVisibility *string `json:"location_visibility,omitempty"`
	MaxCapacity *any `json:"max_capacity,omitempty"`
	MeetingUrl *any `json:"meeting_url,omitempty"`
	Name *string `json:"name,omitempty"`
	NameRequirement *string `json:"name_requirement,omitempty"`
	PhoneNumberRequirement *any `json:"phone_number_requirement,omitempty"`
	Platform *string `json:"platform,omitempty"`
	RegistrationOpen *bool `json:"registration_open,omitempty"`
	RegistrationQuestion *[]any `json:"registration_question,omitempty"`
	RemindersDisabled *bool `json:"reminders_disabled,omitempty"`
	RequireApproval *bool `json:"require_approval,omitempty"`
	ShowGuestList *bool `json:"show_guest_list,omitempty"`
	Slug *string `json:"slug,omitempty"`
	SpotsRemaining *any `json:"spots_remaining,omitempty"`
	StartAt *string `json:"start_at,omitempty"`
	SuppressNotification *bool `json:"suppress_notification,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
	TintColor *string `json:"tint_color,omitempty"`
	Url *string `json:"url,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
	WaitlistStatus *string `json:"waitlist_status,omitempty"`
}

// EventCancelRequest is the typed data model for the event_cancel_request entity.
type EventCancelRequest struct {
	CancellationToken string `json:"cancellation_token"`
	EventId string `json:"event_id"`
	GuestCount float64 `json:"guest_count"`
	IsPaid bool `json:"is_paid"`
}

// EventCancelRequestCreateData is the typed request payload for EventCancelRequest.CreateTyped.
type EventCancelRequestCreateData struct {
	CancellationToken string `json:"cancellation_token"`
	EventId string `json:"event_id"`
	GuestCount float64 `json:"guest_count"`
	IsPaid bool `json:"is_paid"`
}

// EventCoupon is the typed data model for the event_coupon entity.
type EventCoupon struct {
	CentsOff any `json:"cents_off"`
	Code string `json:"code"`
	Currency any `json:"currency"`
	Discount any `json:"discount"`
	EventId string `json:"event_id"`
	EventTicketTypeId *string `json:"event_ticket_type_id,omitempty"`
	Id string `json:"id"`
	PercentOff any `json:"percent_off"`
	RemainingCount int `json:"remaining_count"`
	ValidEndAt any `json:"valid_end_at"`
	ValidStartAt any `json:"valid_start_at"`
}

// EventCouponListMatch is the typed request payload for EventCoupon.ListTyped.
type EventCouponListMatch struct {
	CentsOff *any `json:"cents_off,omitempty"`
	Code *string `json:"code,omitempty"`
	Currency *any `json:"currency,omitempty"`
	Discount *any `json:"discount,omitempty"`
	EventId *string `json:"event_id,omitempty"`
	EventTicketTypeId *string `json:"event_ticket_type_id,omitempty"`
	Id *string `json:"id,omitempty"`
	PercentOff *any `json:"percent_off,omitempty"`
	RemainingCount *int `json:"remaining_count,omitempty"`
	ValidEndAt *any `json:"valid_end_at,omitempty"`
	ValidStartAt *any `json:"valid_start_at,omitempty"`
}

// EventCouponCreateData is the typed request payload for EventCoupon.CreateTyped.
type EventCouponCreateData struct {
	CentsOff any `json:"cents_off"`
	Code string `json:"code"`
	Currency any `json:"currency"`
	Discount any `json:"discount"`
	EventId string `json:"event_id"`
	EventTicketTypeId *string `json:"event_ticket_type_id,omitempty"`
	Id string `json:"id"`
	PercentOff any `json:"percent_off"`
	RemainingCount int `json:"remaining_count"`
	ValidEndAt any `json:"valid_end_at"`
	ValidStartAt any `json:"valid_start_at"`
}

// EventCouponUpdateData is the typed request payload for EventCoupon.UpdateTyped.
type EventCouponUpdateData struct {
	CentsOff *any `json:"cents_off,omitempty"`
	Code *string `json:"code,omitempty"`
	Currency *any `json:"currency,omitempty"`
	Discount *any `json:"discount,omitempty"`
	EventId *string `json:"event_id,omitempty"`
	EventTicketTypeId *string `json:"event_ticket_type_id,omitempty"`
	Id *string `json:"id,omitempty"`
	PercentOff *any `json:"percent_off,omitempty"`
	RemainingCount *int `json:"remaining_count,omitempty"`
	ValidEndAt *any `json:"valid_end_at,omitempty"`
	ValidStartAt *any `json:"valid_start_at,omitempty"`
}

// EventTag is the typed data model for the event_tag entity.
type EventTag struct {
	Color *any `json:"color,omitempty"`
	Id string `json:"id"`
	Name string `json:"name"`
	TagId string `json:"tag_id"`
}

// EventTagListMatch is the typed request payload for EventTag.ListTyped.
type EventTagListMatch struct {
	Color *any `json:"color,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	TagId *string `json:"tag_id,omitempty"`
}

// EventTagCreateData is the typed request payload for EventTag.CreateTyped.
type EventTagCreateData struct {
	Color *any `json:"color,omitempty"`
	Id string `json:"id"`
	Name string `json:"name"`
	TagId string `json:"tag_id"`
}

// EventTagUpdateData is the typed request payload for EventTag.UpdateTyped.
type EventTagUpdateData struct {
	Color *any `json:"color,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	TagId *string `json:"tag_id,omitempty"`
}

// EventTagRemoveMatch is the typed request payload for EventTag.RemoveTyped.
type EventTagRemoveMatch struct {
	Color *any `json:"color,omitempty"`
	Id string `json:"id"`
	Name *string `json:"name,omitempty"`
	TagId *string `json:"tag_id,omitempty"`
}

// EventTagAssignment is the typed data model for the event_tag_assignment entity.
type EventTagAssignment struct {
	AppliedCount float64 `json:"applied_count"`
	EventId []any `json:"event_id"`
	SkippedCount float64 `json:"skipped_count"`
	Tag string `json:"tag"`
}

// EventTagAssignmentCreateData is the typed request payload for EventTagAssignment.CreateTyped.
type EventTagAssignmentCreateData struct {
	AppliedCount float64 `json:"applied_count"`
	EventId []any `json:"event_id"`
	SkippedCount float64 `json:"skipped_count"`
	Tag string `json:"tag"`
}

// EventTagAssignmentRemoveMatch is the typed request payload for EventTagAssignment.RemoveTyped.
type EventTagAssignmentRemoveMatch struct {
	AppliedCount *float64 `json:"applied_count,omitempty"`
	EventId *[]any `json:"event_id,omitempty"`
	SkippedCount *float64 `json:"skipped_count,omitempty"`
	Tag *string `json:"tag,omitempty"`
}

// Guest is the typed data model for the guest entity.
type Guest struct {
	ApprovalStatus string `json:"approval_status"`
	CheckInQrCode string `json:"check_in_qr_code"`
	EthAddress any `json:"eth_address"`
	EventId string `json:"event_id"`
	EventTicket []any `json:"event_ticket"`
	EventTicketOrder []any `json:"event_ticket_order"`
	Guest []any `json:"guest"`
	GuestId string `json:"guest_id"`
	Id string `json:"id"`
	InvitedAt any `json:"invited_at"`
	JoinedAt any `json:"joined_at"`
	Message *any `json:"message,omitempty"`
	PhoneNumber int `json:"phone_number"`
	RegisteredAt any `json:"registered_at"`
	RegistrationAnswer any `json:"registration_answer"`
	SendEmail *any `json:"send_email,omitempty"`
	ShouldRefund *bool `json:"should_refund,omitempty"`
	SolanaAddress any `json:"solana_address"`
	Status string `json:"status"`
	Ticket *any `json:"ticket,omitempty"`
	UserEmail string `json:"user_email"`
	UserFirstName any `json:"user_first_name"`
	UserId string `json:"user_id"`
	UserLastName any `json:"user_last_name"`
	UserName any `json:"user_name"`
	UtmSource any `json:"utm_source"`
}

// GuestLoadMatch is the typed request payload for Guest.LoadTyped.
type GuestLoadMatch struct {
	ApprovalStatus *string `json:"approval_status,omitempty"`
	CheckInQrCode *string `json:"check_in_qr_code,omitempty"`
	EthAddress *any `json:"eth_address,omitempty"`
	EventId *string `json:"event_id,omitempty"`
	EventTicket *[]any `json:"event_ticket,omitempty"`
	EventTicketOrder *[]any `json:"event_ticket_order,omitempty"`
	Guest *[]any `json:"guest,omitempty"`
	GuestId *string `json:"guest_id,omitempty"`
	Id string `json:"id"`
	InvitedAt *any `json:"invited_at,omitempty"`
	JoinedAt *any `json:"joined_at,omitempty"`
	Message *any `json:"message,omitempty"`
	PhoneNumber *int `json:"phone_number,omitempty"`
	RegisteredAt *any `json:"registered_at,omitempty"`
	RegistrationAnswer *any `json:"registration_answer,omitempty"`
	SendEmail *any `json:"send_email,omitempty"`
	ShouldRefund *bool `json:"should_refund,omitempty"`
	SolanaAddress *any `json:"solana_address,omitempty"`
	Status *string `json:"status,omitempty"`
	Ticket *any `json:"ticket,omitempty"`
	UserEmail *string `json:"user_email,omitempty"`
	UserFirstName *any `json:"user_first_name,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	UserLastName *any `json:"user_last_name,omitempty"`
	UserName *any `json:"user_name,omitempty"`
	UtmSource *any `json:"utm_source,omitempty"`
}

// GuestListMatch is the typed request payload for Guest.ListTyped.
type GuestListMatch struct {
	ApprovalStatus *string `json:"approval_status,omitempty"`
	CheckInQrCode *string `json:"check_in_qr_code,omitempty"`
	EthAddress *any `json:"eth_address,omitempty"`
	EventId *string `json:"event_id,omitempty"`
	EventTicket *[]any `json:"event_ticket,omitempty"`
	EventTicketOrder *[]any `json:"event_ticket_order,omitempty"`
	Guest *[]any `json:"guest,omitempty"`
	GuestId *string `json:"guest_id,omitempty"`
	Id *string `json:"id,omitempty"`
	InvitedAt *any `json:"invited_at,omitempty"`
	JoinedAt *any `json:"joined_at,omitempty"`
	Message *any `json:"message,omitempty"`
	PhoneNumber *int `json:"phone_number,omitempty"`
	RegisteredAt *any `json:"registered_at,omitempty"`
	RegistrationAnswer *any `json:"registration_answer,omitempty"`
	SendEmail *any `json:"send_email,omitempty"`
	ShouldRefund *bool `json:"should_refund,omitempty"`
	SolanaAddress *any `json:"solana_address,omitempty"`
	Status *string `json:"status,omitempty"`
	Ticket *any `json:"ticket,omitempty"`
	UserEmail *string `json:"user_email,omitempty"`
	UserFirstName *any `json:"user_first_name,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	UserLastName *any `json:"user_last_name,omitempty"`
	UserName *any `json:"user_name,omitempty"`
	UtmSource *any `json:"utm_source,omitempty"`
}

// GuestCreateData is the typed request payload for Guest.CreateTyped.
type GuestCreateData struct {
	ApprovalStatus string `json:"approval_status"`
	CheckInQrCode string `json:"check_in_qr_code"`
	EthAddress any `json:"eth_address"`
	EventId string `json:"event_id"`
	EventTicket []any `json:"event_ticket"`
	EventTicketOrder []any `json:"event_ticket_order"`
	Guest []any `json:"guest"`
	GuestId string `json:"guest_id"`
	Id string `json:"id"`
	InvitedAt any `json:"invited_at"`
	JoinedAt any `json:"joined_at"`
	Message *any `json:"message,omitempty"`
	PhoneNumber int `json:"phone_number"`
	RegisteredAt any `json:"registered_at"`
	RegistrationAnswer any `json:"registration_answer"`
	SendEmail *any `json:"send_email,omitempty"`
	ShouldRefund *bool `json:"should_refund,omitempty"`
	SolanaAddress any `json:"solana_address"`
	Status string `json:"status"`
	Ticket *any `json:"ticket,omitempty"`
	UserEmail string `json:"user_email"`
	UserFirstName any `json:"user_first_name"`
	UserId string `json:"user_id"`
	UserLastName any `json:"user_last_name"`
	UserName any `json:"user_name"`
	UtmSource any `json:"utm_source"`
}

// GuestUpdateData is the typed request payload for Guest.UpdateTyped.
type GuestUpdateData struct {
	ApprovalStatus *string `json:"approval_status,omitempty"`
	CheckInQrCode *string `json:"check_in_qr_code,omitempty"`
	EthAddress *any `json:"eth_address,omitempty"`
	EventId *string `json:"event_id,omitempty"`
	EventTicket *[]any `json:"event_ticket,omitempty"`
	EventTicketOrder *[]any `json:"event_ticket_order,omitempty"`
	Guest *[]any `json:"guest,omitempty"`
	GuestId *string `json:"guest_id,omitempty"`
	Id *string `json:"id,omitempty"`
	InvitedAt *any `json:"invited_at,omitempty"`
	JoinedAt *any `json:"joined_at,omitempty"`
	Message *any `json:"message,omitempty"`
	PhoneNumber *int `json:"phone_number,omitempty"`
	RegisteredAt *any `json:"registered_at,omitempty"`
	RegistrationAnswer *any `json:"registration_answer,omitempty"`
	SendEmail *any `json:"send_email,omitempty"`
	ShouldRefund *bool `json:"should_refund,omitempty"`
	SolanaAddress *any `json:"solana_address,omitempty"`
	Status *string `json:"status,omitempty"`
	Ticket *any `json:"ticket,omitempty"`
	UserEmail *string `json:"user_email,omitempty"`
	UserFirstName *any `json:"user_first_name,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	UserLastName *any `json:"user_last_name,omitempty"`
	UserName *any `json:"user_name,omitempty"`
	UtmSource *any `json:"utm_source,omitempty"`
}

// GuestInvite is the typed data model for the guest_invite entity.
type GuestInvite struct {
	EventId string `json:"event_id"`
	Guest []any `json:"guest"`
	Message *any `json:"message,omitempty"`
}

// GuestInviteCreateData is the typed request payload for GuestInvite.CreateTyped.
type GuestInviteCreateData struct {
	EventId string `json:"event_id"`
	Guest []any `json:"guest"`
	Message *any `json:"message,omitempty"`
}

// GuestTicket is the typed data model for the guest_ticket entity.
type GuestTicket struct {
	EventId string `json:"event_id"`
	GuestId string `json:"guest_id"`
	SendEmail *any `json:"send_email,omitempty"`
	TicketIdsToRemove *[]any `json:"ticket_ids_to_remove,omitempty"`
	TicketsToAdd *[]any `json:"tickets_to_add,omitempty"`
}

// GuestTicketUpdateData is the typed request payload for GuestTicket.UpdateTyped.
type GuestTicketUpdateData struct {
	EventId *string `json:"event_id,omitempty"`
	GuestId *string `json:"guest_id,omitempty"`
	SendEmail *any `json:"send_email,omitempty"`
	TicketIdsToRemove *[]any `json:"ticket_ids_to_remove,omitempty"`
	TicketsToAdd *[]any `json:"tickets_to_add,omitempty"`
}

// Host is the typed data model for the host entity.
type Host struct {
	AccessLevel *any `json:"access_level,omitempty"`
	Email string `json:"email"`
	EventId string `json:"event_id"`
	IsVisible *bool `json:"is_visible,omitempty"`
	Name *string `json:"name,omitempty"`
}

// HostCreateData is the typed request payload for Host.CreateTyped.
type HostCreateData struct {
	AccessLevel *any `json:"access_level,omitempty"`
	Email string `json:"email"`
	EventId string `json:"event_id"`
	IsVisible *bool `json:"is_visible,omitempty"`
	Name *string `json:"name,omitempty"`
}

// HostUpdateData is the typed request payload for Host.UpdateTyped.
type HostUpdateData struct {
	AccessLevel *any `json:"access_level,omitempty"`
	Email *string `json:"email,omitempty"`
	EventId *string `json:"event_id,omitempty"`
	IsVisible *bool `json:"is_visible,omitempty"`
	Name *string `json:"name,omitempty"`
}

// HostRemoveMatch is the typed request payload for Host.RemoveTyped.
type HostRemoveMatch struct {
	AccessLevel *any `json:"access_level,omitempty"`
	Email *string `json:"email,omitempty"`
	EventId *string `json:"event_id,omitempty"`
	IsVisible *bool `json:"is_visible,omitempty"`
	Name *string `json:"name,omitempty"`
}

// ImageUpload is the typed data model for the image_upload entity.
type ImageUpload struct {
	ContentType *any `json:"content_type,omitempty"`
	FileUrl string `json:"file_url"`
	UploadUrl string `json:"upload_url"`
}

// ImageUploadCreateData is the typed request payload for ImageUpload.CreateTyped.
type ImageUploadCreateData struct {
	ContentType *any `json:"content_type,omitempty"`
	FileUrl string `json:"file_url"`
	UploadUrl string `json:"upload_url"`
}

// Member is the typed data model for the member entity.
type Member struct {
	Email string `json:"email"`
	MembershipId string `json:"membership_id"`
	MembershipTierId string `json:"membership_tier_id"`
	RegistrationAnswer *[]any `json:"registration_answer,omitempty"`
	SkipPayment *bool `json:"skip_payment,omitempty"`
	Status string `json:"status"`
	UserId string `json:"user_id"`
}

// MemberCreateData is the typed request payload for Member.CreateTyped.
type MemberCreateData struct {
	Email string `json:"email"`
	MembershipId string `json:"membership_id"`
	MembershipTierId string `json:"membership_tier_id"`
	RegistrationAnswer *[]any `json:"registration_answer,omitempty"`
	SkipPayment *bool `json:"skip_payment,omitempty"`
	Status string `json:"status"`
	UserId string `json:"user_id"`
}

// MemberUpdateData is the typed request payload for Member.UpdateTyped.
type MemberUpdateData struct {
	Email *string `json:"email,omitempty"`
	MembershipId *string `json:"membership_id,omitempty"`
	MembershipTierId *string `json:"membership_tier_id,omitempty"`
	RegistrationAnswer *[]any `json:"registration_answer,omitempty"`
	SkipPayment *bool `json:"skip_payment,omitempty"`
	Status *string `json:"status,omitempty"`
	UserId *string `json:"user_id,omitempty"`
}

// MembershipTier is the typed data model for the membership_tier entity.
type MembershipTier struct {
	AccessInfo any `json:"access_info"`
	Description string `json:"description"`
	Id string `json:"id"`
	Name string `json:"name"`
	TintColor string `json:"tint_color"`
}

// MembershipTierListMatch is the typed request payload for MembershipTier.ListTyped.
type MembershipTierListMatch struct {
	AccessInfo *any `json:"access_info,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	TintColor *string `json:"tint_color,omitempty"`
}

// OrganizationAdmin is the typed data model for the organization_admin entity.
type OrganizationAdmin struct {
	ApiId string `json:"api_id"`
	AvatarUrl string `json:"avatar_url"`
	Email string `json:"email"`
	FirstName any `json:"first_name"`
	Id string `json:"id"`
	LastName any `json:"last_name"`
	Name string `json:"name"`
}

// OrganizationAdminListMatch is the typed request payload for OrganizationAdmin.ListTyped.
type OrganizationAdminListMatch struct {
	ApiId *string `json:"api_id,omitempty"`
	AvatarUrl *string `json:"avatar_url,omitempty"`
	Email *string `json:"email,omitempty"`
	FirstName *any `json:"first_name,omitempty"`
	Id *string `json:"id,omitempty"`
	LastName *any `json:"last_name,omitempty"`
	Name *string `json:"name,omitempty"`
}

// OrganizationCalendar is the typed data model for the organization_calendar entity.
type OrganizationCalendar struct {
	AvatarUrl any `json:"avatar_url"`
	Coordinate any `json:"coordinate"`
	CoverImageUrl any `json:"cover_image_url"`
	Description string `json:"description"`
	Id string `json:"id"`
	InstagramHandle any `json:"instagram_handle"`
	IsPersonal bool `json:"is_personal"`
	Location any `json:"location"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	SocialImageUrl any `json:"social_image_url"`
	TintColor *string `json:"tint_color,omitempty"`
	TwitterHandle any `json:"twitter_handle"`
	Url string `json:"url"`
	Website any `json:"website"`
	YoutubeHandle any `json:"youtube_handle"`
}

// OrganizationCalendarListMatch is the typed request payload for OrganizationCalendar.ListTyped.
type OrganizationCalendarListMatch struct {
	AvatarUrl *any `json:"avatar_url,omitempty"`
	Coordinate *any `json:"coordinate,omitempty"`
	CoverImageUrl *any `json:"cover_image_url,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	InstagramHandle *any `json:"instagram_handle,omitempty"`
	IsPersonal *bool `json:"is_personal,omitempty"`
	Location *any `json:"location,omitempty"`
	Name *string `json:"name,omitempty"`
	Slug *string `json:"slug,omitempty"`
	SocialImageUrl *any `json:"social_image_url,omitempty"`
	TintColor *string `json:"tint_color,omitempty"`
	TwitterHandle *any `json:"twitter_handle,omitempty"`
	Url *string `json:"url,omitempty"`
	Website *any `json:"website,omitempty"`
	YoutubeHandle *any `json:"youtube_handle,omitempty"`
}

// OrganizationCalendarCreateData is the typed request payload for OrganizationCalendar.CreateTyped.
type OrganizationCalendarCreateData struct {
	AvatarUrl any `json:"avatar_url"`
	Coordinate any `json:"coordinate"`
	CoverImageUrl any `json:"cover_image_url"`
	Description string `json:"description"`
	Id string `json:"id"`
	InstagramHandle any `json:"instagram_handle"`
	IsPersonal bool `json:"is_personal"`
	Location any `json:"location"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	SocialImageUrl any `json:"social_image_url"`
	TintColor *string `json:"tint_color,omitempty"`
	TwitterHandle any `json:"twitter_handle"`
	Url string `json:"url"`
	Website any `json:"website"`
	YoutubeHandle any `json:"youtube_handle"`
}

// OrganizationEvent is the typed data model for the organization_event entity.
type OrganizationEvent struct {
	ApiId string `json:"api_id"`
	CalendarApiId string `json:"calendar_api_id"`
	CalendarId string `json:"calendar_id"`
	Coordinate any `json:"coordinate"`
	CoverUrl string `json:"cover_url"`
	CreatedAt string `json:"created_at"`
	DisplayPrice any `json:"display_price"`
	DurationInterval string `json:"duration_interval"`
	EndAt string `json:"end_at"`
	FeedbackEmail map[string]any `json:"feedback_email"`
	GeoAddressJson any `json:"geo_address_json"`
	GeoLatitude any `json:"geo_latitude"`
	GeoLongitude any `json:"geo_longitude"`
	Id string `json:"id"`
	LocationType string `json:"location_type"`
	LocationVisibility string `json:"location_visibility"`
	ManagingCalendar []any `json:"managing_calendar"`
	MeetingUrl any `json:"meeting_url"`
	Name string `json:"name"`
	Platform string `json:"platform"`
	RegistrationOpen bool `json:"registration_open"`
	RegistrationQuestion *[]any `json:"registration_question,omitempty"`
	RequireApproval bool `json:"require_approval"`
	SpotsRemaining any `json:"spots_remaining"`
	StartAt string `json:"start_at"`
	Timezone string `json:"timezone"`
	Url string `json:"url"`
	UserApiId string `json:"user_api_id"`
	UserId string `json:"user_id"`
	Visibility string `json:"visibility"`
	WaitlistStatus string `json:"waitlist_status"`
	ZoomMeetingUrl any `json:"zoom_meeting_url"`
}

// OrganizationEventListMatch is the typed request payload for OrganizationEvent.ListTyped.
type OrganizationEventListMatch struct {
	ApiId *string `json:"api_id,omitempty"`
	CalendarApiId *string `json:"calendar_api_id,omitempty"`
	CalendarId *string `json:"calendar_id,omitempty"`
	Coordinate *any `json:"coordinate,omitempty"`
	CoverUrl *string `json:"cover_url,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	DisplayPrice *any `json:"display_price,omitempty"`
	DurationInterval *string `json:"duration_interval,omitempty"`
	EndAt *string `json:"end_at,omitempty"`
	FeedbackEmail *map[string]any `json:"feedback_email,omitempty"`
	GeoAddressJson *any `json:"geo_address_json,omitempty"`
	GeoLatitude *any `json:"geo_latitude,omitempty"`
	GeoLongitude *any `json:"geo_longitude,omitempty"`
	Id *string `json:"id,omitempty"`
	LocationType *string `json:"location_type,omitempty"`
	LocationVisibility *string `json:"location_visibility,omitempty"`
	ManagingCalendar *[]any `json:"managing_calendar,omitempty"`
	MeetingUrl *any `json:"meeting_url,omitempty"`
	Name *string `json:"name,omitempty"`
	Platform *string `json:"platform,omitempty"`
	RegistrationOpen *bool `json:"registration_open,omitempty"`
	RegistrationQuestion *[]any `json:"registration_question,omitempty"`
	RequireApproval *bool `json:"require_approval,omitempty"`
	SpotsRemaining *any `json:"spots_remaining,omitempty"`
	StartAt *string `json:"start_at,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
	Url *string `json:"url,omitempty"`
	UserApiId *string `json:"user_api_id,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
	WaitlistStatus *string `json:"waitlist_status,omitempty"`
	ZoomMeetingUrl *any `json:"zoom_meeting_url,omitempty"`
}

// OrganizationEventTransfer is the typed data model for the organization_event_transfer entity.
type OrganizationEventTransfer struct {
	CalendarId string `json:"calendar_id"`
	EventId string `json:"event_id"`
}

// OrganizationEventTransferCreateData is the typed request payload for OrganizationEventTransfer.CreateTyped.
type OrganizationEventTransferCreateData struct {
	CalendarId string `json:"calendar_id"`
	EventId string `json:"event_id"`
}

// TicketType is the typed data model for the ticket_type entity.
type TicketType struct {
	Cent *any `json:"cent,omitempty"`
	Currency *any `json:"currency,omitempty"`
	Description *string `json:"description,omitempty"`
	Id string `json:"id"`
	IsFlexible *bool `json:"is_flexible,omitempty"`
	IsHidden *bool `json:"is_hidden,omitempty"`
	MaxCapacity *any `json:"max_capacity,omitempty"`
	MinCent *any `json:"min_cent,omitempty"`
	Name string `json:"name"`
	RequireApproval *bool `json:"require_approval,omitempty"`
	Type string `json:"type"`
	ValidEndAt *any `json:"valid_end_at,omitempty"`
	ValidStartAt *any `json:"valid_start_at,omitempty"`
}

// TicketTypeLoadMatch is the typed request payload for TicketType.LoadTyped.
type TicketTypeLoadMatch struct {
	Cent *any `json:"cent,omitempty"`
	Currency *any `json:"currency,omitempty"`
	Description *string `json:"description,omitempty"`
	Id string `json:"id"`
	IsFlexible *bool `json:"is_flexible,omitempty"`
	IsHidden *bool `json:"is_hidden,omitempty"`
	MaxCapacity *any `json:"max_capacity,omitempty"`
	MinCent *any `json:"min_cent,omitempty"`
	Name *string `json:"name,omitempty"`
	RequireApproval *bool `json:"require_approval,omitempty"`
	Type *string `json:"type,omitempty"`
	ValidEndAt *any `json:"valid_end_at,omitempty"`
	ValidStartAt *any `json:"valid_start_at,omitempty"`
}

// TicketTypeListMatch is the typed request payload for TicketType.ListTyped.
type TicketTypeListMatch struct {
	Cent *any `json:"cent,omitempty"`
	Currency *any `json:"currency,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	IsFlexible *bool `json:"is_flexible,omitempty"`
	IsHidden *bool `json:"is_hidden,omitempty"`
	MaxCapacity *any `json:"max_capacity,omitempty"`
	MinCent *any `json:"min_cent,omitempty"`
	Name *string `json:"name,omitempty"`
	RequireApproval *bool `json:"require_approval,omitempty"`
	Type *string `json:"type,omitempty"`
	ValidEndAt *any `json:"valid_end_at,omitempty"`
	ValidStartAt *any `json:"valid_start_at,omitempty"`
}

// TicketTypeCreateData is the typed request payload for TicketType.CreateTyped.
type TicketTypeCreateData struct {
	Cent *any `json:"cent,omitempty"`
	Currency *any `json:"currency,omitempty"`
	Description *string `json:"description,omitempty"`
	Id string `json:"id"`
	IsFlexible *bool `json:"is_flexible,omitempty"`
	IsHidden *bool `json:"is_hidden,omitempty"`
	MaxCapacity *any `json:"max_capacity,omitempty"`
	MinCent *any `json:"min_cent,omitempty"`
	Name string `json:"name"`
	RequireApproval *bool `json:"require_approval,omitempty"`
	Type string `json:"type"`
	ValidEndAt *any `json:"valid_end_at,omitempty"`
	ValidStartAt *any `json:"valid_start_at,omitempty"`
}

// TicketTypeUpdateData is the typed request payload for TicketType.UpdateTyped.
type TicketTypeUpdateData struct {
	Cent *any `json:"cent,omitempty"`
	Currency *any `json:"currency,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	IsFlexible *bool `json:"is_flexible,omitempty"`
	IsHidden *bool `json:"is_hidden,omitempty"`
	MaxCapacity *any `json:"max_capacity,omitempty"`
	MinCent *any `json:"min_cent,omitempty"`
	Name *string `json:"name,omitempty"`
	RequireApproval *bool `json:"require_approval,omitempty"`
	Type *string `json:"type,omitempty"`
	ValidEndAt *any `json:"valid_end_at,omitempty"`
	ValidStartAt *any `json:"valid_start_at,omitempty"`
}

// TicketTypeRemoveMatch is the typed request payload for TicketType.RemoveTyped.
type TicketTypeRemoveMatch struct {
	Cent *any `json:"cent,omitempty"`
	Currency *any `json:"currency,omitempty"`
	Description *string `json:"description,omitempty"`
	Id string `json:"id"`
	IsFlexible *bool `json:"is_flexible,omitempty"`
	IsHidden *bool `json:"is_hidden,omitempty"`
	MaxCapacity *any `json:"max_capacity,omitempty"`
	MinCent *any `json:"min_cent,omitempty"`
	Name *string `json:"name,omitempty"`
	RequireApproval *bool `json:"require_approval,omitempty"`
	Type *string `json:"type,omitempty"`
	ValidEndAt *any `json:"valid_end_at,omitempty"`
	ValidStartAt *any `json:"valid_start_at,omitempty"`
}

// User is the typed data model for the user entity.
type User struct {
	AvatarUrl string `json:"avatar_url"`
	Email string `json:"email"`
	FirstName any `json:"first_name"`
	Id string `json:"id"`
	LastName any `json:"last_name"`
	Name string `json:"name"`
}

// UserLoadMatch is the typed request payload for User.LoadTyped.
type UserLoadMatch struct {
	AvatarUrl *string `json:"avatar_url,omitempty"`
	Email *string `json:"email,omitempty"`
	FirstName *any `json:"first_name,omitempty"`
	Id string `json:"id"`
	LastName *any `json:"last_name,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Webhook is the typed data model for the webhook entity.
type Webhook struct {
	CreatedAt string `json:"created_at"`
	EventType []any `json:"event_type"`
	Id string `json:"id"`
	Secret string `json:"secret"`
	Status string `json:"status"`
	Url string `json:"url"`
}

// WebhookLoadMatch is the typed request payload for Webhook.LoadTyped.
type WebhookLoadMatch struct {
	CreatedAt *string `json:"created_at,omitempty"`
	EventType *[]any `json:"event_type,omitempty"`
	Id string `json:"id"`
	Secret *string `json:"secret,omitempty"`
	Status *string `json:"status,omitempty"`
	Url *string `json:"url,omitempty"`
}

// WebhookListMatch is the typed request payload for Webhook.ListTyped.
type WebhookListMatch struct {
	CreatedAt *string `json:"created_at,omitempty"`
	EventType *[]any `json:"event_type,omitempty"`
	Id *string `json:"id,omitempty"`
	Secret *string `json:"secret,omitempty"`
	Status *string `json:"status,omitempty"`
	Url *string `json:"url,omitempty"`
}

// WebhookCreateData is the typed request payload for Webhook.CreateTyped.
type WebhookCreateData struct {
	CreatedAt string `json:"created_at"`
	EventType []any `json:"event_type"`
	Id string `json:"id"`
	Secret string `json:"secret"`
	Status string `json:"status"`
	Url string `json:"url"`
}

// WebhookUpdateData is the typed request payload for Webhook.UpdateTyped.
type WebhookUpdateData struct {
	CreatedAt *string `json:"created_at,omitempty"`
	EventType *[]any `json:"event_type,omitempty"`
	Id *string `json:"id,omitempty"`
	Secret *string `json:"secret,omitempty"`
	Status *string `json:"status,omitempty"`
	Url *string `json:"url,omitempty"`
}

// WebhookRemoveMatch is the typed request payload for Webhook.RemoveTyped.
type WebhookRemoveMatch struct {
	CreatedAt *string `json:"created_at,omitempty"`
	EventType *[]any `json:"event_type,omitempty"`
	Id string `json:"id"`
	Secret *string `json:"secret,omitempty"`
	Status *string `json:"status,omitempty"`
	Url *string `json:"url,omitempty"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedFrom decodes a runtime value (a map[string]any produced by the op
// pipeline) into a typed model T via a JSON round-trip. On any error it
// returns the zero value of T; the op's own (value, error) tuple carries the
// real error.
func typedFrom[T any](v any) T {
	var out T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value ([]any of maps) into a typed
// slice []T via a JSON round-trip, for list ops.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
