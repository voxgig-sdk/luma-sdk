package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewCalendarEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewCalendarAdminEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewCalendarCouponEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewCalendarEventEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewCalendarEventApprovalEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewCalendarEventRejectionEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewContactEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewContactBlockEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewContactRestoreEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewContactTagEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewContactTagAssignmentEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewEntityLookupEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewEventEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewEventCancelRequestEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewEventCouponEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewEventTagEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewEventTagAssignmentEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewGuestEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewGuestInviteEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewGuestTicketEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewHostEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewImageUploadEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewMemberEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewMembershipTierEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewOrganizationAdminEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewOrganizationCalendarEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewOrganizationEventEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewOrganizationEventTransferEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewTicketTypeEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewUserEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

var NewWebhookEntityFunc func(client *LumaSDK, entopts map[string]any) LumaEntity

