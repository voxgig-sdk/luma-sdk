package voxgiglumasdk

import (
	"github.com/voxgig-sdk/luma-sdk/go/core"
	"github.com/voxgig-sdk/luma-sdk/go/entity"
	"github.com/voxgig-sdk/luma-sdk/go/feature"
	_ "github.com/voxgig-sdk/luma-sdk/go/utility"
)

// Type aliases preserve external API.
type LumaSDK = core.LumaSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type LumaEntity = core.LumaEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type LumaError = core.LumaError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewCalendarEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewCalendarEntity(client, entopts)
	}
	core.NewCalendarAdminEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewCalendarAdminEntity(client, entopts)
	}
	core.NewCalendarCouponEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewCalendarCouponEntity(client, entopts)
	}
	core.NewCalendarEventEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewCalendarEventEntity(client, entopts)
	}
	core.NewCalendarEventApprovalEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewCalendarEventApprovalEntity(client, entopts)
	}
	core.NewCalendarEventRejectionEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewCalendarEventRejectionEntity(client, entopts)
	}
	core.NewContactEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewContactEntity(client, entopts)
	}
	core.NewContactBlockEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewContactBlockEntity(client, entopts)
	}
	core.NewContactRestoreEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewContactRestoreEntity(client, entopts)
	}
	core.NewContactTagEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewContactTagEntity(client, entopts)
	}
	core.NewContactTagAssignmentEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewContactTagAssignmentEntity(client, entopts)
	}
	core.NewEntityLookupEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewEntityLookupEntity(client, entopts)
	}
	core.NewEventEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewEventEntity(client, entopts)
	}
	core.NewEventCancelRequestEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewEventCancelRequestEntity(client, entopts)
	}
	core.NewEventCouponEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewEventCouponEntity(client, entopts)
	}
	core.NewEventTagEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewEventTagEntity(client, entopts)
	}
	core.NewEventTagAssignmentEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewEventTagAssignmentEntity(client, entopts)
	}
	core.NewGuestEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewGuestEntity(client, entopts)
	}
	core.NewGuestInviteEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewGuestInviteEntity(client, entopts)
	}
	core.NewGuestTicketEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewGuestTicketEntity(client, entopts)
	}
	core.NewHostEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewHostEntity(client, entopts)
	}
	core.NewImageUploadEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewImageUploadEntity(client, entopts)
	}
	core.NewMemberEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewMemberEntity(client, entopts)
	}
	core.NewMembershipTierEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewMembershipTierEntity(client, entopts)
	}
	core.NewOrganizationAdminEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewOrganizationAdminEntity(client, entopts)
	}
	core.NewOrganizationCalendarEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewOrganizationCalendarEntity(client, entopts)
	}
	core.NewOrganizationEventEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewOrganizationEventEntity(client, entopts)
	}
	core.NewOrganizationEventTransferEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewOrganizationEventTransferEntity(client, entopts)
	}
	core.NewTicketTypeEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewTicketTypeEntity(client, entopts)
	}
	core.NewUserEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewUserEntity(client, entopts)
	}
	core.NewWebhookEntityFunc = func(client *core.LumaSDK, entopts map[string]any) core.LumaEntity {
		return entity.NewWebhookEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewLumaSDK = core.NewLumaSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig

// No-arg convenience constructors. Go has no default-argument syntax,
// so these aliases let callers write `sdk.New()` / `sdk.Test()`
// instead of `sdk.NewLumaSDK(nil)` / `sdk.TestSDK(nil, nil)`
// for the common no-options case.
func New() *LumaSDK  { return NewLumaSDK(nil) }
func Test() *LumaSDK { return TestSDK(nil, nil) }
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature
