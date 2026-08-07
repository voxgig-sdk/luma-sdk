package core

import (
	"fmt"

	vs "github.com/voxgig-sdk/luma-sdk/go/utility/struct"
)

type LumaSDK struct {
	Mode     string
	options  map[string]any
	utility  *Utility
	Features []Feature
	rootctx  *Context
}

func NewLumaSDK(options map[string]any) *LumaSDK {
	sdk := &LumaSDK{
		Mode:     "live",
		Features: []Feature{},
	}

	sdk.utility = NewUtility()

	config := MakeConfig()

	sdk.rootctx = sdk.utility.MakeContext(map[string]any{
		"client":  sdk,
		"utility": sdk.utility,
		"config":  config,
		"options": options,
		"shared":  map[string]any{},
	}, nil)

	sdk.options = sdk.utility.MakeOptions(sdk.rootctx)

	if vs.GetPath([]any{"feature", "test", "active"}, sdk.options) == true {
		sdk.Mode = "test"
	}

	sdk.rootctx.Options = sdk.options

	// Add features in the resolved order (MakeOptions puts an explicit array
	// order first, else defaults to test-first). Ordering matters: the `test`
	// feature installs the base mock transport and the transport features
	// (retry/cache/netsim/proxy/ratelimit) wrap whatever is current, so `test`
	// must be added before them to sit at the base of the chain.
	featureOpts := ToMapAny(vs.GetProp(sdk.options, "feature"))
	if featureOpts != nil {
		if fo, ok := vs.GetPath([]any{"__derived__", "featureorder"}, sdk.options).([]any); ok {
			for _, n := range fo {
				fname, _ := n.(string)
				fopts := ToMapAny(featureOpts[fname])
				if fopts != nil {
					if active, ok := fopts["active"]; ok {
						if ab, ok := active.(bool); ok && ab {
							sdk.utility.FeatureAdd(sdk.rootctx, makeFeature(fname))
						}
					}
				}
			}
		}
	}

	// Add extension features.
	if extend := vs.GetProp(sdk.options, "extend"); extend != nil {
		if extList, ok := extend.([]any); ok {
			for _, f := range extList {
				if feat, ok := f.(Feature); ok {
					sdk.utility.FeatureAdd(sdk.rootctx, feat)
				}
			}
		}
	}

	// Initialize features.
	for _, f := range sdk.Features {
		sdk.utility.FeatureInit(sdk.rootctx, f)
	}

	sdk.utility.FeatureHook(sdk.rootctx, "PostConstruct")

	return sdk
}

func (sdk *LumaSDK) OptionsMap() map[string]any {
	out := vs.Clone(sdk.options)
	if om, ok := out.(map[string]any); ok {
		return om
	}
	return map[string]any{}
}

func (sdk *LumaSDK) GetUtility() *Utility {
	return CopyUtility(sdk.utility)
}

func (sdk *LumaSDK) GetRootCtx() *Context {
	return sdk.rootctx
}

func (sdk *LumaSDK) Prepare(fetchargs map[string]any) (map[string]any, error) {
	utility := sdk.utility

	if fetchargs == nil {
		fetchargs = map[string]any{}
	}

	var ctrl map[string]any
	if c := vs.GetProp(fetchargs, "ctrl"); c != nil {
		if cm, ok := c.(map[string]any); ok {
			ctrl = cm
		}
	}
	if ctrl == nil {
		ctrl = map[string]any{}
	}

	ctx := utility.MakeContext(map[string]any{
		"opname": "prepare",
		"ctrl":   ctrl,
	}, sdk.rootctx)

	options := sdk.options

	path, _ := vs.GetProp(fetchargs, "path").(string)
	method, _ := vs.GetProp(fetchargs, "method").(string)
	if method == "" {
		method = "GET"
	}

	params := ToMapAny(vs.GetProp(fetchargs, "params"))
	if params == nil {
		params = map[string]any{}
	}
	query := ToMapAny(vs.GetProp(fetchargs, "query"))
	if query == nil {
		query = map[string]any{}
	}

	headers := utility.PrepareHeaders(ctx)

	base, _ := vs.GetProp(options, "base").(string)
	prefix, _ := vs.GetProp(options, "prefix").(string)
	suffix, _ := vs.GetProp(options, "suffix").(string)

	ctx.Spec = NewSpec(map[string]any{
		"base":    base,
		"prefix":  prefix,
		"suffix":  suffix,
		"path":    path,
		"method":  method,
		"params":  params,
		"query":   query,
		"headers": headers,
		"body":    vs.GetProp(fetchargs, "body"),
		"step":    "start",
	})

	// Merge user-provided headers.
	if uh := vs.GetProp(fetchargs, "headers"); uh != nil {
		if uhm, ok := uh.(map[string]any); ok {
			for k, v := range uhm {
				ctx.Spec.Headers[k] = v
			}
		}
	}

	_, err := utility.PrepareAuth(ctx)
	if err != nil {
		return nil, err
	}

	return utility.MakeFetchDef(ctx)
}

func (sdk *LumaSDK) Direct(fetchargs map[string]any) (map[string]any, error) {
	utility := sdk.utility

	fetchdef, err := sdk.Prepare(fetchargs)
	if err != nil {
		return map[string]any{"ok": false, "err": err}, nil
	}

	if fetchargs == nil {
		fetchargs = map[string]any{}
	}

	var ctrl map[string]any
	if c := vs.GetProp(fetchargs, "ctrl"); c != nil {
		if cm, ok := c.(map[string]any); ok {
			ctrl = cm
		}
	}
	if ctrl == nil {
		ctrl = map[string]any{}
	}

	ctx := utility.MakeContext(map[string]any{
		"opname": "direct",
		"ctrl":   ctrl,
	}, sdk.rootctx)

	url, _ := fetchdef["url"].(string)
	fetched, fetchErr := utility.Fetcher(ctx, url, fetchdef)

	if fetchErr != nil {
		return map[string]any{"ok": false, "err": fetchErr}, nil
	}

	if fetched == nil {
		return map[string]any{
			"ok":  false,
			"err": ctx.MakeError("direct_no_response", "response: undefined"),
		}, nil
	}

	if fm, ok := fetched.(map[string]any); ok {
		status := ToInt(vs.GetProp(fm, "status"))
		headers := vs.GetProp(fm, "headers")

		// No-body responses (204, 304) and explicit zero content-length
		// must skip JSON parsing — calling json() on an empty body errors.
		var contentLength string
		if hm, ok := headers.(map[string]any); ok {
			if cl, ok := hm["content-length"]; ok {
				contentLength = fmt.Sprintf("%v", cl)
			}
		}
		noBody := status == 204 || status == 304 || contentLength == "0"

		var jsonData any
		if !noBody {
			if jf := vs.GetProp(fm, "json"); jf != nil {
				if f, ok := jf.(func() any); ok {
					// f() returns nil on parse error in our fetcher.
					jsonData = f()
				}
			}
		}

		return map[string]any{
			"ok":      status >= 200 && status < 300,
			"status":  status,
			"headers": headers,
			"data":    jsonData,
		}, nil
	}

	return map[string]any{"ok": false, "err": ctx.MakeError("direct_invalid", "invalid response type")}, nil
}


// Calendar returns a Calendar entity bound to this client.
// Idiomatic usage: client.Calendar(nil).List(nil, nil) or
// client.Calendar(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) Calendar(data map[string]any) LumaEntity {
	return NewCalendarEntityFunc(sdk, data)
}


// CalendarAdmin returns a CalendarAdmin entity bound to this client.
// Idiomatic usage: client.CalendarAdmin(nil).List(nil, nil) or
// client.CalendarAdmin(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) CalendarAdmin(data map[string]any) LumaEntity {
	return NewCalendarAdminEntityFunc(sdk, data)
}


// CalendarCoupon returns a CalendarCoupon entity bound to this client.
// Idiomatic usage: client.CalendarCoupon(nil).List(nil, nil) or
// client.CalendarCoupon(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) CalendarCoupon(data map[string]any) LumaEntity {
	return NewCalendarCouponEntityFunc(sdk, data)
}


// CalendarEvent returns a CalendarEvent entity bound to this client.
// Idiomatic usage: client.CalendarEvent(nil).List(nil, nil) or
// client.CalendarEvent(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) CalendarEvent(data map[string]any) LumaEntity {
	return NewCalendarEventEntityFunc(sdk, data)
}


// CalendarEventApproval returns a CalendarEventApproval entity bound to this client.
// Idiomatic usage: client.CalendarEventApproval(nil).List(nil, nil) or
// client.CalendarEventApproval(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) CalendarEventApproval(data map[string]any) LumaEntity {
	return NewCalendarEventApprovalEntityFunc(sdk, data)
}


// CalendarEventRejection returns a CalendarEventRejection entity bound to this client.
// Idiomatic usage: client.CalendarEventRejection(nil).List(nil, nil) or
// client.CalendarEventRejection(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) CalendarEventRejection(data map[string]any) LumaEntity {
	return NewCalendarEventRejectionEntityFunc(sdk, data)
}


// Contact returns a Contact entity bound to this client.
// Idiomatic usage: client.Contact(nil).List(nil, nil) or
// client.Contact(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) Contact(data map[string]any) LumaEntity {
	return NewContactEntityFunc(sdk, data)
}


// ContactBlock returns a ContactBlock entity bound to this client.
// Idiomatic usage: client.ContactBlock(nil).List(nil, nil) or
// client.ContactBlock(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) ContactBlock(data map[string]any) LumaEntity {
	return NewContactBlockEntityFunc(sdk, data)
}


// ContactRestore returns a ContactRestore entity bound to this client.
// Idiomatic usage: client.ContactRestore(nil).List(nil, nil) or
// client.ContactRestore(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) ContactRestore(data map[string]any) LumaEntity {
	return NewContactRestoreEntityFunc(sdk, data)
}


// ContactTag returns a ContactTag entity bound to this client.
// Idiomatic usage: client.ContactTag(nil).List(nil, nil) or
// client.ContactTag(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) ContactTag(data map[string]any) LumaEntity {
	return NewContactTagEntityFunc(sdk, data)
}


// ContactTagAssignment returns a ContactTagAssignment entity bound to this client.
// Idiomatic usage: client.ContactTagAssignment(nil).List(nil, nil) or
// client.ContactTagAssignment(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) ContactTagAssignment(data map[string]any) LumaEntity {
	return NewContactTagAssignmentEntityFunc(sdk, data)
}


// EntityLookup returns a EntityLookup entity bound to this client.
// Idiomatic usage: client.EntityLookup(nil).List(nil, nil) or
// client.EntityLookup(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) EntityLookup(data map[string]any) LumaEntity {
	return NewEntityLookupEntityFunc(sdk, data)
}


// Event returns a Event entity bound to this client.
// Idiomatic usage: client.Event(nil).List(nil, nil) or
// client.Event(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) Event(data map[string]any) LumaEntity {
	return NewEventEntityFunc(sdk, data)
}


// EventCancelRequest returns a EventCancelRequest entity bound to this client.
// Idiomatic usage: client.EventCancelRequest(nil).List(nil, nil) or
// client.EventCancelRequest(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) EventCancelRequest(data map[string]any) LumaEntity {
	return NewEventCancelRequestEntityFunc(sdk, data)
}


// EventCoupon returns a EventCoupon entity bound to this client.
// Idiomatic usage: client.EventCoupon(nil).List(nil, nil) or
// client.EventCoupon(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) EventCoupon(data map[string]any) LumaEntity {
	return NewEventCouponEntityFunc(sdk, data)
}


// EventTag returns a EventTag entity bound to this client.
// Idiomatic usage: client.EventTag(nil).List(nil, nil) or
// client.EventTag(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) EventTag(data map[string]any) LumaEntity {
	return NewEventTagEntityFunc(sdk, data)
}


// EventTagAssignment returns a EventTagAssignment entity bound to this client.
// Idiomatic usage: client.EventTagAssignment(nil).List(nil, nil) or
// client.EventTagAssignment(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) EventTagAssignment(data map[string]any) LumaEntity {
	return NewEventTagAssignmentEntityFunc(sdk, data)
}


// Guest returns a Guest entity bound to this client.
// Idiomatic usage: client.Guest(nil).List(nil, nil) or
// client.Guest(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) Guest(data map[string]any) LumaEntity {
	return NewGuestEntityFunc(sdk, data)
}


// GuestInvite returns a GuestInvite entity bound to this client.
// Idiomatic usage: client.GuestInvite(nil).List(nil, nil) or
// client.GuestInvite(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) GuestInvite(data map[string]any) LumaEntity {
	return NewGuestInviteEntityFunc(sdk, data)
}


// GuestTicket returns a GuestTicket entity bound to this client.
// Idiomatic usage: client.GuestTicket(nil).List(nil, nil) or
// client.GuestTicket(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) GuestTicket(data map[string]any) LumaEntity {
	return NewGuestTicketEntityFunc(sdk, data)
}


// Host returns a Host entity bound to this client.
// Idiomatic usage: client.Host(nil).List(nil, nil) or
// client.Host(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) Host(data map[string]any) LumaEntity {
	return NewHostEntityFunc(sdk, data)
}


// ImageUpload returns a ImageUpload entity bound to this client.
// Idiomatic usage: client.ImageUpload(nil).List(nil, nil) or
// client.ImageUpload(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) ImageUpload(data map[string]any) LumaEntity {
	return NewImageUploadEntityFunc(sdk, data)
}


// Member returns a Member entity bound to this client.
// Idiomatic usage: client.Member(nil).List(nil, nil) or
// client.Member(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) Member(data map[string]any) LumaEntity {
	return NewMemberEntityFunc(sdk, data)
}


// MembershipTier returns a MembershipTier entity bound to this client.
// Idiomatic usage: client.MembershipTier(nil).List(nil, nil) or
// client.MembershipTier(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) MembershipTier(data map[string]any) LumaEntity {
	return NewMembershipTierEntityFunc(sdk, data)
}


// OrganizationAdmin returns a OrganizationAdmin entity bound to this client.
// Idiomatic usage: client.OrganizationAdmin(nil).List(nil, nil) or
// client.OrganizationAdmin(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) OrganizationAdmin(data map[string]any) LumaEntity {
	return NewOrganizationAdminEntityFunc(sdk, data)
}


// OrganizationCalendar returns a OrganizationCalendar entity bound to this client.
// Idiomatic usage: client.OrganizationCalendar(nil).List(nil, nil) or
// client.OrganizationCalendar(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) OrganizationCalendar(data map[string]any) LumaEntity {
	return NewOrganizationCalendarEntityFunc(sdk, data)
}


// OrganizationEvent returns a OrganizationEvent entity bound to this client.
// Idiomatic usage: client.OrganizationEvent(nil).List(nil, nil) or
// client.OrganizationEvent(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) OrganizationEvent(data map[string]any) LumaEntity {
	return NewOrganizationEventEntityFunc(sdk, data)
}


// OrganizationEventTransfer returns a OrganizationEventTransfer entity bound to this client.
// Idiomatic usage: client.OrganizationEventTransfer(nil).List(nil, nil) or
// client.OrganizationEventTransfer(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) OrganizationEventTransfer(data map[string]any) LumaEntity {
	return NewOrganizationEventTransferEntityFunc(sdk, data)
}


// TicketType returns a TicketType entity bound to this client.
// Idiomatic usage: client.TicketType(nil).List(nil, nil) or
// client.TicketType(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) TicketType(data map[string]any) LumaEntity {
	return NewTicketTypeEntityFunc(sdk, data)
}


// User returns a User entity bound to this client.
// Idiomatic usage: client.User(nil).List(nil, nil) or
// client.User(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) User(data map[string]any) LumaEntity {
	return NewUserEntityFunc(sdk, data)
}


// Webhook returns a Webhook entity bound to this client.
// Idiomatic usage: client.Webhook(nil).List(nil, nil) or
// client.Webhook(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *LumaSDK) Webhook(data map[string]any) LumaEntity {
	return NewWebhookEntityFunc(sdk, data)
}



func TestSDK(testopts map[string]any, sdkopts map[string]any) *LumaSDK {
	if sdkopts == nil {
		sdkopts = map[string]any{}
	}
	sdkopts = vs.Clone(sdkopts).(map[string]any)

	if testopts == nil {
		testopts = map[string]any{}
	}
	testopts = vs.Clone(testopts).(map[string]any)
	testopts["active"] = true

	vs.SetPath(sdkopts, []any{"feature", "test"}, testopts)

	sdk := NewLumaSDK(sdkopts)
	sdk.Mode = "test"

	return sdk
}
