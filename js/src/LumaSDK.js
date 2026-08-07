// Luma Js SDK

const { CalendarEntity } = require('./entity/CalendarEntity')
const { CalendarAdminEntity } = require('./entity/CalendarAdminEntity')
const { CalendarCouponEntity } = require('./entity/CalendarCouponEntity')
const { CalendarEventEntity } = require('./entity/CalendarEventEntity')
const { CalendarEventApprovalEntity } = require('./entity/CalendarEventApprovalEntity')
const { CalendarEventRejectionEntity } = require('./entity/CalendarEventRejectionEntity')
const { ContactEntity } = require('./entity/ContactEntity')
const { ContactBlockEntity } = require('./entity/ContactBlockEntity')
const { ContactRestoreEntity } = require('./entity/ContactRestoreEntity')
const { ContactTagEntity } = require('./entity/ContactTagEntity')
const { ContactTagAssignmentEntity } = require('./entity/ContactTagAssignmentEntity')
const { EntityLookupEntity } = require('./entity/EntityLookupEntity')
const { EventEntity } = require('./entity/EventEntity')
const { EventCancelRequestEntity } = require('./entity/EventCancelRequestEntity')
const { EventCouponEntity } = require('./entity/EventCouponEntity')
const { EventTagEntity } = require('./entity/EventTagEntity')
const { EventTagAssignmentEntity } = require('./entity/EventTagAssignmentEntity')
const { GuestEntity } = require('./entity/GuestEntity')
const { GuestInviteEntity } = require('./entity/GuestInviteEntity')
const { GuestTicketEntity } = require('./entity/GuestTicketEntity')
const { HostEntity } = require('./entity/HostEntity')
const { ImageUploadEntity } = require('./entity/ImageUploadEntity')
const { MemberEntity } = require('./entity/MemberEntity')
const { MembershipTierEntity } = require('./entity/MembershipTierEntity')
const { OrganizationAdminEntity } = require('./entity/OrganizationAdminEntity')
const { OrganizationCalendarEntity } = require('./entity/OrganizationCalendarEntity')
const { OrganizationEventEntity } = require('./entity/OrganizationEventEntity')
const { OrganizationEventTransferEntity } = require('./entity/OrganizationEventTransferEntity')
const { TicketTypeEntity } = require('./entity/TicketTypeEntity')
const { UserEntity } = require('./entity/UserEntity')
const { WebhookEntity } = require('./entity/WebhookEntity')


const { inspect } = require('node:util')

const { config } = require('./Config')
const { Utility } = require('./utility/Utility')
const { LumaEntityBase } = require('./LumaEntityBase')


const { BaseFeature } = require('./feature/base/BaseFeature')


const stdutil = new Utility()


class LumaSDK {
  _mode = 'live'
  _options
  _utility = new Utility()
  _features
  _rootctx

  constructor(options) {

    this._rootctx = this._utility.makeContext({
      client: this,
      utility: this._utility,
      config,
      options,
      shared: new WeakMap()
    })

    this._options = this._utility.makeOptions(this._rootctx)

    const struct = this._utility.struct
    const getpath = struct.getpath

    if (true === getpath(this._options.feature, 'test.active')) {
      this._mode = 'test'
    }

    this._rootctx.options = this._options

    this._features = []

    const featureAdd = this._utility.featureAdd
    const featureInit = this._utility.featureInit

    // Add features in the resolved order (makeOptions puts an explicit
    // array order first, else defaults to test-first). Ordering matters:
    // the `test` feature installs the base mock transport and the transport
    // features (retry/cache/netsim/proxy/ratelimit) wrap whatever is current,
    // so `test` must be added before them to sit at the base of the chain.
    const featureorder = getpath(this._options, '__derived__.featureorder') || []
    for (const fname of featureorder) {
      const fopts = this._options.feature[fname] || {}
      if (fopts.active) {
        featureAdd(this._rootctx, this._rootctx.config.makeFeature(fname))
      }
    }

    if (null != this._options.extend) {
      for (let f of this._options.extend) {
        featureAdd(this._rootctx, f)
      }
    }

    for (let f of this._features) {
      featureInit(this._rootctx, f)
    }

    const featureHook = this._utility.featureHook
    featureHook(this._rootctx, 'PostConstruct')
  }


  options() {
    return this._utility.struct.clone(this._options)
  }


  utility() {
    return this._utility.struct.clone(this._utility)
  }


  async prepare(fetchargs) {
    const utility = this._utility
    const struct = utility.struct
    const clone = struct.clone

    const {
      makeContext,
      makeFetchDef,
      prepareHeaders,
      prepareAuth,
    } = utility

    fetchargs = fetchargs || {}

    let ctx = makeContext({
      opname: 'prepare',
      ctrl: fetchargs.ctrl || {},
    }, this._rootctx)

    const options = this._options

    // Build spec directly from SDK options + user-provided fetch args.
    const spec = {
      base: options.base,
      prefix: options.prefix,
      suffix: options.suffix,
      path: fetchargs.path || '',
      method: fetchargs.method || 'GET',
      params: fetchargs.params || {},
      query: fetchargs.query || {},
      headers: prepareHeaders(ctx),
      body: fetchargs.body,
      step: 'start',
    }

    ctx.spec = spec

    // Merge user-provided headers over SDK defaults.
    if (fetchargs.headers) {
      const uheaders = fetchargs.headers
      for (let key in uheaders) {
        spec.headers[key] = uheaders[key]
      }
    }

    // Apply SDK auth (apikey, auth prefix, etc.)
    const authResult = prepareAuth(ctx)
    if (authResult instanceof Error) {
      return authResult
    }

    return makeFetchDef(ctx)
  }


  async direct(fetchargs) {
    const utility = this._utility
    const fetcher = utility.fetcher
    const makeContext = utility.makeContext

    const fetchdef = await this.prepare(fetchargs)
    if (fetchdef instanceof Error) {
      return fetchdef
    }

    let ctx = makeContext({
      opname: 'direct',
      ctrl: (fetchargs || {}).ctrl || {},
    }, this._rootctx)

    try {
      const fetched = await fetcher(ctx, fetchdef.url, fetchdef)

      if (null == fetched) {
        return { ok: false, err: ctx.error('direct_no_response', 'response: undefined') }
      }
      else if (fetched instanceof Error) {
        return { ok: false, err: fetched }
      }

      const status = fetched.status
      const json = 'function' === typeof fetched.json ? await fetched.json() : fetched.json

      return {
        ok: status >= 200 && status < 300,
        status,
        headers: fetched.headers,
        data: json,
      }
    }
    catch (err) {
      return { ok: false, err }
    }
  }



  // Entity access: `client.Calendar().list()` / `client.Calendar().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Calendar(entopts) {
    const self = this
    return new CalendarEntity(self, entopts)
  }


  // Entity access: `client.CalendarAdmin().list()` / `client.CalendarAdmin().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  CalendarAdmin(entopts) {
    const self = this
    return new CalendarAdminEntity(self, entopts)
  }


  // Entity access: `client.CalendarCoupon().list()` / `client.CalendarCoupon().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  CalendarCoupon(entopts) {
    const self = this
    return new CalendarCouponEntity(self, entopts)
  }


  // Entity access: `client.CalendarEvent().list()` / `client.CalendarEvent().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  CalendarEvent(entopts) {
    const self = this
    return new CalendarEventEntity(self, entopts)
  }


  // Entity access: `client.CalendarEventApproval().list()` / `client.CalendarEventApproval().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  CalendarEventApproval(entopts) {
    const self = this
    return new CalendarEventApprovalEntity(self, entopts)
  }


  // Entity access: `client.CalendarEventRejection().list()` / `client.CalendarEventRejection().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  CalendarEventRejection(entopts) {
    const self = this
    return new CalendarEventRejectionEntity(self, entopts)
  }


  // Entity access: `client.Contact().list()` / `client.Contact().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Contact(entopts) {
    const self = this
    return new ContactEntity(self, entopts)
  }


  // Entity access: `client.ContactBlock().list()` / `client.ContactBlock().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ContactBlock(entopts) {
    const self = this
    return new ContactBlockEntity(self, entopts)
  }


  // Entity access: `client.ContactRestore().list()` / `client.ContactRestore().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ContactRestore(entopts) {
    const self = this
    return new ContactRestoreEntity(self, entopts)
  }


  // Entity access: `client.ContactTag().list()` / `client.ContactTag().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ContactTag(entopts) {
    const self = this
    return new ContactTagEntity(self, entopts)
  }


  // Entity access: `client.ContactTagAssignment().list()` / `client.ContactTagAssignment().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ContactTagAssignment(entopts) {
    const self = this
    return new ContactTagAssignmentEntity(self, entopts)
  }


  // Entity access: `client.EntityLookup().list()` / `client.EntityLookup().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EntityLookup(entopts) {
    const self = this
    return new EntityLookupEntity(self, entopts)
  }


  // Entity access: `client.Event().list()` / `client.Event().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Event(entopts) {
    const self = this
    return new EventEntity(self, entopts)
  }


  // Entity access: `client.EventCancelRequest().list()` / `client.EventCancelRequest().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EventCancelRequest(entopts) {
    const self = this
    return new EventCancelRequestEntity(self, entopts)
  }


  // Entity access: `client.EventCoupon().list()` / `client.EventCoupon().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EventCoupon(entopts) {
    const self = this
    return new EventCouponEntity(self, entopts)
  }


  // Entity access: `client.EventTag().list()` / `client.EventTag().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EventTag(entopts) {
    const self = this
    return new EventTagEntity(self, entopts)
  }


  // Entity access: `client.EventTagAssignment().list()` / `client.EventTagAssignment().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  EventTagAssignment(entopts) {
    const self = this
    return new EventTagAssignmentEntity(self, entopts)
  }


  // Entity access: `client.Guest().list()` / `client.Guest().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Guest(entopts) {
    const self = this
    return new GuestEntity(self, entopts)
  }


  // Entity access: `client.GuestInvite().list()` / `client.GuestInvite().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  GuestInvite(entopts) {
    const self = this
    return new GuestInviteEntity(self, entopts)
  }


  // Entity access: `client.GuestTicket().list()` / `client.GuestTicket().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  GuestTicket(entopts) {
    const self = this
    return new GuestTicketEntity(self, entopts)
  }


  // Entity access: `client.Host().list()` / `client.Host().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Host(entopts) {
    const self = this
    return new HostEntity(self, entopts)
  }


  // Entity access: `client.ImageUpload().list()` / `client.ImageUpload().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ImageUpload(entopts) {
    const self = this
    return new ImageUploadEntity(self, entopts)
  }


  // Entity access: `client.Member().list()` / `client.Member().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Member(entopts) {
    const self = this
    return new MemberEntity(self, entopts)
  }


  // Entity access: `client.MembershipTier().list()` / `client.MembershipTier().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  MembershipTier(entopts) {
    const self = this
    return new MembershipTierEntity(self, entopts)
  }


  // Entity access: `client.OrganizationAdmin().list()` / `client.OrganizationAdmin().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  OrganizationAdmin(entopts) {
    const self = this
    return new OrganizationAdminEntity(self, entopts)
  }


  // Entity access: `client.OrganizationCalendar().list()` / `client.OrganizationCalendar().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  OrganizationCalendar(entopts) {
    const self = this
    return new OrganizationCalendarEntity(self, entopts)
  }


  // Entity access: `client.OrganizationEvent().list()` / `client.OrganizationEvent().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  OrganizationEvent(entopts) {
    const self = this
    return new OrganizationEventEntity(self, entopts)
  }


  // Entity access: `client.OrganizationEventTransfer().list()` / `client.OrganizationEventTransfer().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  OrganizationEventTransfer(entopts) {
    const self = this
    return new OrganizationEventTransferEntity(self, entopts)
  }


  // Entity access: `client.TicketType().list()` / `client.TicketType().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  TicketType(entopts) {
    const self = this
    return new TicketTypeEntity(self, entopts)
  }


  // Entity access: `client.User().list()` / `client.User().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  User(entopts) {
    const self = this
    return new UserEntity(self, entopts)
  }


  // Entity access: `client.Webhook().list()` / `client.Webhook().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Webhook(entopts) {
    const self = this
    return new WebhookEntity(self, entopts)
  }




  static test(testoptsarg, sdkoptsarg) {
    const struct = stdutil.struct
    const setpath = struct.setpath
    const getdef = struct.getdef
    const clone = struct.clone
    const setprop = struct.setprop

    const sdkopts = getdef(clone(sdkoptsarg), {})
    const testopts = getdef(clone(testoptsarg), {})
    setprop(testopts, 'active', true)
    setpath(sdkopts, 'feature.test', testopts)

    const testsdk = new LumaSDK(sdkopts)
    testsdk._mode = 'test'

    return testsdk
  }


  tester(testopts, sdkopts) {
    return LumaSDK.test(testopts, sdkopts)
  }


  toJSON() {
    return { name: 'Luma' }
  }

  toString() {
    return 'Luma ' + this._utility.struct.jsonify(this.toJSON())
  }

  [inspect.custom]() {
    return this.toString()
  }

}




const SDK = LumaSDK


module.exports = {
  stdutil,
  config,

  BaseFeature,
  LumaEntityBase,

  LumaSDK,
  SDK,
}

