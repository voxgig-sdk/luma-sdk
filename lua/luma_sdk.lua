-- Luma SDK

local vs = require("utility.struct.struct")
local Utility = require("core.utility_type")
local Spec = require("core.spec")
local helpers = require("core.helpers")

-- Load utility registration (populates Utility._registrar)
require("utility.register")

-- Typed-model annotations (LuaLS ---@class); empty at runtime.
require("luma_types")

-- Load features
local BaseFeature = require("feature.base_feature")
local features_factory = require("features")


local LumaSDK = {}
LumaSDK.__index = LumaSDK


local function _make_feature(name)
  local factory = features_factory[name]
  if factory ~= nil then
    return factory()
  end
  return features_factory.base()
end

LumaSDK._make_feature = _make_feature


function LumaSDK.new(options)
  local self = setmetatable({}, LumaSDK)
  self.mode = "live"
  self.features = {}
  self.options = nil

  local utility = Utility.new()
  self._utility = utility

  local config = require("config")()

  self._rootctx = utility.make_context({
    client = self,
    utility = utility,
    config = config,
    options = options or {},
    shared = {},
  }, nil)

  self.options = utility.make_options(self._rootctx)

  if vs.getpath(self.options, "feature.test.active") == true then
    self.mode = "test"
  end

  self._rootctx.options = self.options

  -- Add features in the resolved order (make_options puts an explicit list
  -- order first, else defaults to test-first). Ordering matters: the `test`
  -- feature installs the base mock transport and the transport features
  -- (retry/cache/netsim/proxy/ratelimit) wrap whatever is current, so `test`
  -- must be added before them to sit at the base of the chain.
  local feature_opts = helpers.to_map(vs.getprop(self.options, "feature"))
  if feature_opts ~= nil then
    local featureorder = vs.getpath(self.options, "__derived__.featureorder")
    if type(featureorder) == "table" then
      for _, fname in ipairs(featureorder) do
        local fopts = helpers.to_map(feature_opts[fname])
        if fopts ~= nil and fopts["active"] == true then
          utility.feature_add(self._rootctx, _make_feature(fname))
        end
      end
    end
  end

  -- Add extension features.
  local extend = vs.getprop(self.options, "extend")
  if type(extend) == "table" then
    for _, f in ipairs(extend) do
      if type(f) == "table" and type(f.get_name) == "function" then
        utility.feature_add(self._rootctx, f)
      end
    end
  end

  -- Initialize features.
  for _, f in ipairs(self.features) do
    utility.feature_init(self._rootctx, f)
  end

  utility.feature_hook(self._rootctx, "PostConstruct")

    -- feature: test


  return self
end


function LumaSDK:options_map()
  local out = vs.clone(self.options)
  if type(out) == "table" then
    return out
  end
  return {}
end


function LumaSDK:get_utility()
  return Utility.copy(self._utility)
end


function LumaSDK:get_root_ctx()
  return self._rootctx
end


function LumaSDK:prepare(fetchargs)
  local utility = self._utility

  fetchargs = fetchargs or {}

  local ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl")) or {}

  local ctx = utility.make_context({
    opname = "prepare",
    ctrl = ctrl,
  }, self._rootctx)

  local options = self.options

  local path = vs.getprop(fetchargs, "path") or ""
  if type(path) ~= "string" then path = "" end

  local method = vs.getprop(fetchargs, "method") or "GET"
  if type(method) ~= "string" then method = "GET" end

  local params = helpers.to_map(vs.getprop(fetchargs, "params")) or {}
  local query = helpers.to_map(vs.getprop(fetchargs, "query")) or {}

  local headers = utility.prepare_headers(ctx)

  local base = vs.getprop(options, "base") or ""
  if type(base) ~= "string" then base = "" end
  local prefix = vs.getprop(options, "prefix") or ""
  if type(prefix) ~= "string" then prefix = "" end
  local suffix = vs.getprop(options, "suffix") or ""
  if type(suffix) ~= "string" then suffix = "" end

  ctx.spec = Spec.new({
    base = base,
    prefix = prefix,
    suffix = suffix,
    path = path,
    method = method,
    params = params,
    query = query,
    headers = headers,
    body = vs.getprop(fetchargs, "body"),
    step = "start",
  })

  -- Merge user-provided headers.
  local uh = vs.getprop(fetchargs, "headers")
  if type(uh) == "table" then
    for k, v in pairs(uh) do
      ctx.spec.headers[k] = v
    end
  end

  local _, err = utility.prepare_auth(ctx)
  if err ~= nil then
    return nil, err
  end

  return utility.make_fetch_def(ctx)
end


function LumaSDK:direct(fetchargs)
  local utility = self._utility

  local fetchdef, err = self:prepare(fetchargs)
  if err ~= nil then
    return { ok = false, err = err }, nil
  end

  fetchargs = fetchargs or {}
  local ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl")) or {}

  local ctx = utility.make_context({
    opname = "direct",
    ctrl = ctrl,
  }, self._rootctx)

  local url = fetchdef["url"] or ""
  local fetched, fetch_err = utility.fetcher(ctx, url, fetchdef)

  if fetch_err ~= nil then
    return { ok = false, err = fetch_err }, nil
  end

  if fetched == nil then
    return {
      ok = false,
      err = ctx:make_error("direct_no_response", "response: undefined"),
    }, nil
  end

  if type(fetched) == "table" then
    local status = helpers.to_int(vs.getprop(fetched, "status"))
    local headers = vs.getprop(fetched, "headers") or {}

    -- No-body responses (204, 304) and explicit zero content-length
    -- must skip JSON parsing — calling json() on an empty body errors.
    local content_length = nil
    if type(headers) == "table" then
      content_length = headers["content-length"]
    end
    local no_body = status == 204 or status == 304 or tostring(content_length) == "0"

    local json_data = nil
    if not no_body then
      local jf = vs.getprop(fetched, "json")
      if type(jf) == "function" then
        local ok, result = pcall(jf)
        if ok then
          json_data = result
        end
        -- Non-JSON body: json_data stays nil, status/headers preserved.
      end
    end

    return {
      ok = status >= 200 and status < 300,
      status = status,
      headers = headers,
      data = json_data,
    }, nil
  end

  return {
    ok = false,
    err = ctx:make_error("direct_invalid", "invalid response type"),
  }, nil
end



-- Idiomatic facade: client:Calendar():list() / client:Calendar():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:Calendar(data)
  local EntityMod = require("entity.calendar_entity")
  if data == nil then
    if self._calendar == nil then
      self._calendar = EntityMod.new(self, nil)
    end
    return self._calendar
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:CalendarAdmin():list() / client:CalendarAdmin():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:CalendarAdmin(data)
  local EntityMod = require("entity.calendar_admin_entity")
  if data == nil then
    if self._calendar_admin == nil then
      self._calendar_admin = EntityMod.new(self, nil)
    end
    return self._calendar_admin
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:CalendarCoupon():list() / client:CalendarCoupon():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:CalendarCoupon(data)
  local EntityMod = require("entity.calendar_coupon_entity")
  if data == nil then
    if self._calendar_coupon == nil then
      self._calendar_coupon = EntityMod.new(self, nil)
    end
    return self._calendar_coupon
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:CalendarEvent():list() / client:CalendarEvent():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:CalendarEvent(data)
  local EntityMod = require("entity.calendar_event_entity")
  if data == nil then
    if self._calendar_event == nil then
      self._calendar_event = EntityMod.new(self, nil)
    end
    return self._calendar_event
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:CalendarEventApproval():list() / client:CalendarEventApproval():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:CalendarEventApproval(data)
  local EntityMod = require("entity.calendar_event_approval_entity")
  if data == nil then
    if self._calendar_event_approval == nil then
      self._calendar_event_approval = EntityMod.new(self, nil)
    end
    return self._calendar_event_approval
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:CalendarEventRejection():list() / client:CalendarEventRejection():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:CalendarEventRejection(data)
  local EntityMod = require("entity.calendar_event_rejection_entity")
  if data == nil then
    if self._calendar_event_rejection == nil then
      self._calendar_event_rejection = EntityMod.new(self, nil)
    end
    return self._calendar_event_rejection
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Contact():list() / client:Contact():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:Contact(data)
  local EntityMod = require("entity.contact_entity")
  if data == nil then
    if self._contact == nil then
      self._contact = EntityMod.new(self, nil)
    end
    return self._contact
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ContactBlock():list() / client:ContactBlock():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:ContactBlock(data)
  local EntityMod = require("entity.contact_block_entity")
  if data == nil then
    if self._contact_block == nil then
      self._contact_block = EntityMod.new(self, nil)
    end
    return self._contact_block
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ContactRestore():list() / client:ContactRestore():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:ContactRestore(data)
  local EntityMod = require("entity.contact_restore_entity")
  if data == nil then
    if self._contact_restore == nil then
      self._contact_restore = EntityMod.new(self, nil)
    end
    return self._contact_restore
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ContactTag():list() / client:ContactTag():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:ContactTag(data)
  local EntityMod = require("entity.contact_tag_entity")
  if data == nil then
    if self._contact_tag == nil then
      self._contact_tag = EntityMod.new(self, nil)
    end
    return self._contact_tag
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ContactTagAssignment():list() / client:ContactTagAssignment():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:ContactTagAssignment(data)
  local EntityMod = require("entity.contact_tag_assignment_entity")
  if data == nil then
    if self._contact_tag_assignment == nil then
      self._contact_tag_assignment = EntityMod.new(self, nil)
    end
    return self._contact_tag_assignment
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EntityLookup():list() / client:EntityLookup():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:EntityLookup(data)
  local EntityMod = require("entity.entity_lookup_entity")
  if data == nil then
    if self._entity_lookup == nil then
      self._entity_lookup = EntityMod.new(self, nil)
    end
    return self._entity_lookup
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Event():list() / client:Event():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:Event(data)
  local EntityMod = require("entity.event_entity")
  if data == nil then
    if self._event == nil then
      self._event = EntityMod.new(self, nil)
    end
    return self._event
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EventCancelRequest():list() / client:EventCancelRequest():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:EventCancelRequest(data)
  local EntityMod = require("entity.event_cancel_request_entity")
  if data == nil then
    if self._event_cancel_request == nil then
      self._event_cancel_request = EntityMod.new(self, nil)
    end
    return self._event_cancel_request
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EventCoupon():list() / client:EventCoupon():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:EventCoupon(data)
  local EntityMod = require("entity.event_coupon_entity")
  if data == nil then
    if self._event_coupon == nil then
      self._event_coupon = EntityMod.new(self, nil)
    end
    return self._event_coupon
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EventTag():list() / client:EventTag():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:EventTag(data)
  local EntityMod = require("entity.event_tag_entity")
  if data == nil then
    if self._event_tag == nil then
      self._event_tag = EntityMod.new(self, nil)
    end
    return self._event_tag
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:EventTagAssignment():list() / client:EventTagAssignment():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:EventTagAssignment(data)
  local EntityMod = require("entity.event_tag_assignment_entity")
  if data == nil then
    if self._event_tag_assignment == nil then
      self._event_tag_assignment = EntityMod.new(self, nil)
    end
    return self._event_tag_assignment
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Guest():list() / client:Guest():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:Guest(data)
  local EntityMod = require("entity.guest_entity")
  if data == nil then
    if self._guest == nil then
      self._guest = EntityMod.new(self, nil)
    end
    return self._guest
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:GuestInvite():list() / client:GuestInvite():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:GuestInvite(data)
  local EntityMod = require("entity.guest_invite_entity")
  if data == nil then
    if self._guest_invite == nil then
      self._guest_invite = EntityMod.new(self, nil)
    end
    return self._guest_invite
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:GuestTicket():list() / client:GuestTicket():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:GuestTicket(data)
  local EntityMod = require("entity.guest_ticket_entity")
  if data == nil then
    if self._guest_ticket == nil then
      self._guest_ticket = EntityMod.new(self, nil)
    end
    return self._guest_ticket
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Host():list() / client:Host():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:Host(data)
  local EntityMod = require("entity.host_entity")
  if data == nil then
    if self._host == nil then
      self._host = EntityMod.new(self, nil)
    end
    return self._host
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ImageUpload():list() / client:ImageUpload():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:ImageUpload(data)
  local EntityMod = require("entity.image_upload_entity")
  if data == nil then
    if self._image_upload == nil then
      self._image_upload = EntityMod.new(self, nil)
    end
    return self._image_upload
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Member():list() / client:Member():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:Member(data)
  local EntityMod = require("entity.member_entity")
  if data == nil then
    if self._member == nil then
      self._member = EntityMod.new(self, nil)
    end
    return self._member
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:MembershipTier():list() / client:MembershipTier():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:MembershipTier(data)
  local EntityMod = require("entity.membership_tier_entity")
  if data == nil then
    if self._membership_tier == nil then
      self._membership_tier = EntityMod.new(self, nil)
    end
    return self._membership_tier
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:OrganizationAdmin():list() / client:OrganizationAdmin():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:OrganizationAdmin(data)
  local EntityMod = require("entity.organization_admin_entity")
  if data == nil then
    if self._organization_admin == nil then
      self._organization_admin = EntityMod.new(self, nil)
    end
    return self._organization_admin
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:OrganizationCalendar():list() / client:OrganizationCalendar():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:OrganizationCalendar(data)
  local EntityMod = require("entity.organization_calendar_entity")
  if data == nil then
    if self._organization_calendar == nil then
      self._organization_calendar = EntityMod.new(self, nil)
    end
    return self._organization_calendar
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:OrganizationEvent():list() / client:OrganizationEvent():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:OrganizationEvent(data)
  local EntityMod = require("entity.organization_event_entity")
  if data == nil then
    if self._organization_event == nil then
      self._organization_event = EntityMod.new(self, nil)
    end
    return self._organization_event
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:OrganizationEventTransfer():list() / client:OrganizationEventTransfer():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:OrganizationEventTransfer(data)
  local EntityMod = require("entity.organization_event_transfer_entity")
  if data == nil then
    if self._organization_event_transfer == nil then
      self._organization_event_transfer = EntityMod.new(self, nil)
    end
    return self._organization_event_transfer
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:TicketType():list() / client:TicketType():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:TicketType(data)
  local EntityMod = require("entity.ticket_type_entity")
  if data == nil then
    if self._ticket_type == nil then
      self._ticket_type = EntityMod.new(self, nil)
    end
    return self._ticket_type
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:User():list() / client:User():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:User(data)
  local EntityMod = require("entity.user_entity")
  if data == nil then
    if self._user == nil then
      self._user = EntityMod.new(self, nil)
    end
    return self._user
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Webhook():list() / client:Webhook():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function LumaSDK:Webhook(data)
  local EntityMod = require("entity.webhook_entity")
  if data == nil then
    if self._webhook == nil then
      self._webhook = EntityMod.new(self, nil)
    end
    return self._webhook
  end
  return EntityMod.new(self, data)
end




function LumaSDK.test(testopts, sdkopts)
  sdkopts = sdkopts or {}
  sdkopts = vs.clone(sdkopts)
  if type(sdkopts) ~= "table" then
    sdkopts = {}
  end

  testopts = testopts or {}
  testopts = vs.clone(testopts)
  if type(testopts) ~= "table" then
    testopts = {}
  end
  testopts["active"] = true

  vs.setpath(sdkopts, "feature.test", testopts)

  local sdk = LumaSDK.new(sdkopts)
  sdk.mode = "test"

  return sdk
end


return LumaSDK
