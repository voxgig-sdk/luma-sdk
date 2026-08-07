# Luma SDK

from utility.voxgig_struct import voxgig_struct as vs
from core.utility_type import LumaUtility
from core.spec import LumaSpec
from core import helpers

# Load utility registration (populates Utility._registrar)
from utility import register

# Load features
from feature.base_feature import LumaBaseFeature
from features import _make_feature


class LumaSDK:

    def __init__(self, options=None):
        self.mode = "live"
        self.features = []
        self.options = None

        utility = LumaUtility()
        self._utility = utility

        from config import make_config
        config = make_config()

        self._rootctx = utility.make_context({
            "client": self,
            "utility": utility,
            "config": config,
            "options": options if options is not None else {},
            "shared": {},
        }, None)

        self.options = utility.make_options(self._rootctx)

        if vs.getpath(self.options, "feature.test.active") is True:
            self.mode = "test"

        self._rootctx.options = self.options

        # Add features in the resolved order (make_options puts an explicit
        # list order first, else defaults to test-first). Ordering matters: the
        # `test` feature installs the base mock transport and the transport
        # features (retry/cache/netsim/proxy/ratelimit) wrap whatever is
        # current, so `test` must be added before them to sit at the base.
        feature_opts = helpers.to_map(vs.getprop(self.options, "feature"))
        if feature_opts is not None:
            featureorder = vs.getpath(self.options, "__derived__.featureorder")
            if isinstance(featureorder, list):
                for fname in featureorder:
                    fopts = helpers.to_map(feature_opts.get(fname))
                    if fopts is not None and fopts.get("active") is True:
                        utility.feature_add(self._rootctx, _make_feature(fname))

        # Add extension features.
        extend = vs.getprop(self.options, "extend")
        if isinstance(extend, list):
            for f in extend:
                if isinstance(f, dict) or (hasattr(f, "get_name") and callable(f.get_name)):
                    utility.feature_add(self._rootctx, f)

        # Initialize features.
        for f in self.features:
            utility.feature_init(self._rootctx, f)

        utility.feature_hook(self._rootctx, "PostConstruct")

        # #BuildFeatures

    def options_map(self):
        out = vs.clone(self.options)
        if isinstance(out, dict):
            return out
        return {}

    def get_utility(self):
        return LumaUtility.copy(self._utility)

    def get_root_ctx(self):
        return self._rootctx

    def prepare(self, fetchargs=None):
        utility = self._utility

        if fetchargs is None:
            fetchargs = {}

        ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl"))
        if ctrl is None:
            ctrl = {}

        ctx = utility.make_context({
            "opname": "prepare",
            "ctrl": ctrl,
        }, self._rootctx)

        options = self.options

        path = vs.getprop(fetchargs, "path") or ""
        if not isinstance(path, str):
            path = ""

        method = vs.getprop(fetchargs, "method") or "GET"
        if not isinstance(method, str):
            method = "GET"

        params = helpers.to_map(vs.getprop(fetchargs, "params"))
        if params is None:
            params = {}
        query = helpers.to_map(vs.getprop(fetchargs, "query"))
        if query is None:
            query = {}

        headers = utility.prepare_headers(ctx)

        base = vs.getprop(options, "base") or ""
        if not isinstance(base, str):
            base = ""
        prefix = vs.getprop(options, "prefix") or ""
        if not isinstance(prefix, str):
            prefix = ""
        suffix = vs.getprop(options, "suffix") or ""
        if not isinstance(suffix, str):
            suffix = ""

        ctx.spec = LumaSpec({
            "base": base,
            "prefix": prefix,
            "suffix": suffix,
            "path": path,
            "method": method,
            "params": params,
            "query": query,
            "headers": headers,
            "body": vs.getprop(fetchargs, "body"),
            "step": "start",
        })

        # Merge user-provided headers.
        uh = vs.getprop(fetchargs, "headers")
        if isinstance(uh, dict):
            for k, v in uh.items():
                ctx.spec.headers[k] = v

        _, err = utility.prepare_auth(ctx)
        if err is not None:
            raise err

        fetchdef, err = utility.make_fetch_def(ctx)
        if err is not None:
            raise err

        return fetchdef

    def direct(self, fetchargs=None):
        utility = self._utility

        try:
            fetchdef = self.prepare(fetchargs)
        except Exception as err:
            # direct() is the raw-HTTP escape hatch: it never raises, it
            # returns a result object callers branch on via result["ok"].
            return {"ok": False, "err": err}

        if fetchargs is None:
            fetchargs = {}
        ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl"))
        if ctrl is None:
            ctrl = {}

        ctx = utility.make_context({
            "opname": "direct",
            "ctrl": ctrl,
        }, self._rootctx)

        url = fetchdef.get("url", "")
        fetched, fetch_err = utility.fetcher(ctx, url, fetchdef)

        if fetch_err is not None:
            return {"ok": False, "err": fetch_err}

        if fetched is None:
            return {
                "ok": False,
                "err": ctx.make_error("direct_no_response", "response: undefined"),
            }

        if isinstance(fetched, dict):
            status = helpers.to_int(vs.getprop(fetched, "status"))
            headers = vs.getprop(fetched, "headers") or {}

            # No-body responses (204, 304) and explicit zero content-length
            # must skip JSON parsing — calling json() on an empty body raises.
            content_length = None
            if isinstance(headers, dict):
                content_length = headers.get("content-length")
            no_body = status in (204, 304) or str(content_length) == "0"

            json_data = None
            if not no_body:
                jf = vs.getprop(fetched, "json")
                if callable(jf):
                    try:
                        json_data = jf()
                    except Exception:
                        # Non-JSON body (e.g. text/plain, text/html). Surface
                        # status + headers but leave data as None.
                        json_data = None

            return {
                "ok": status >= 200 and status < 300,
                "status": status,
                "headers": headers,
                "data": json_data,
            }

        return {
            "ok": False,
            "err": ctx.make_error("direct_invalid", "invalid response type"),
        }


    def Calendar(self, data=None) -> "CalendarEntity":
        """Entity factory: client.Calendar().list() / client.Calendar().load({"id": ...})."""
        from entity.calendar_entity import CalendarEntity
        return CalendarEntity(self, data)


    def CalendarAdmin(self, data=None) -> "CalendarAdminEntity":
        """Entity factory: client.CalendarAdmin().list() / client.CalendarAdmin().load({"id": ...})."""
        from entity.calendar_admin_entity import CalendarAdminEntity
        return CalendarAdminEntity(self, data)


    def CalendarCoupon(self, data=None) -> "CalendarCouponEntity":
        """Entity factory: client.CalendarCoupon().list() / client.CalendarCoupon().load({"id": ...})."""
        from entity.calendar_coupon_entity import CalendarCouponEntity
        return CalendarCouponEntity(self, data)


    def CalendarEvent(self, data=None) -> "CalendarEventEntity":
        """Entity factory: client.CalendarEvent().list() / client.CalendarEvent().load({"id": ...})."""
        from entity.calendar_event_entity import CalendarEventEntity
        return CalendarEventEntity(self, data)


    def CalendarEventApproval(self, data=None) -> "CalendarEventApprovalEntity":
        """Entity factory: client.CalendarEventApproval().list() / client.CalendarEventApproval().load({"id": ...})."""
        from entity.calendar_event_approval_entity import CalendarEventApprovalEntity
        return CalendarEventApprovalEntity(self, data)


    def CalendarEventRejection(self, data=None) -> "CalendarEventRejectionEntity":
        """Entity factory: client.CalendarEventRejection().list() / client.CalendarEventRejection().load({"id": ...})."""
        from entity.calendar_event_rejection_entity import CalendarEventRejectionEntity
        return CalendarEventRejectionEntity(self, data)


    def Contact(self, data=None) -> "ContactEntity":
        """Entity factory: client.Contact().list() / client.Contact().load({"id": ...})."""
        from entity.contact_entity import ContactEntity
        return ContactEntity(self, data)


    def ContactBlock(self, data=None) -> "ContactBlockEntity":
        """Entity factory: client.ContactBlock().list() / client.ContactBlock().load({"id": ...})."""
        from entity.contact_block_entity import ContactBlockEntity
        return ContactBlockEntity(self, data)


    def ContactRestore(self, data=None) -> "ContactRestoreEntity":
        """Entity factory: client.ContactRestore().list() / client.ContactRestore().load({"id": ...})."""
        from entity.contact_restore_entity import ContactRestoreEntity
        return ContactRestoreEntity(self, data)


    def ContactTag(self, data=None) -> "ContactTagEntity":
        """Entity factory: client.ContactTag().list() / client.ContactTag().load({"id": ...})."""
        from entity.contact_tag_entity import ContactTagEntity
        return ContactTagEntity(self, data)


    def ContactTagAssignment(self, data=None) -> "ContactTagAssignmentEntity":
        """Entity factory: client.ContactTagAssignment().list() / client.ContactTagAssignment().load({"id": ...})."""
        from entity.contact_tag_assignment_entity import ContactTagAssignmentEntity
        return ContactTagAssignmentEntity(self, data)


    def EntityLookup(self, data=None) -> "EntityLookupEntity":
        """Entity factory: client.EntityLookup().list() / client.EntityLookup().load({"id": ...})."""
        from entity.entity_lookup_entity import EntityLookupEntity
        return EntityLookupEntity(self, data)


    def Event(self, data=None) -> "EventEntity":
        """Entity factory: client.Event().list() / client.Event().load({"id": ...})."""
        from entity.event_entity import EventEntity
        return EventEntity(self, data)


    def EventCancelRequest(self, data=None) -> "EventCancelRequestEntity":
        """Entity factory: client.EventCancelRequest().list() / client.EventCancelRequest().load({"id": ...})."""
        from entity.event_cancel_request_entity import EventCancelRequestEntity
        return EventCancelRequestEntity(self, data)


    def EventCoupon(self, data=None) -> "EventCouponEntity":
        """Entity factory: client.EventCoupon().list() / client.EventCoupon().load({"id": ...})."""
        from entity.event_coupon_entity import EventCouponEntity
        return EventCouponEntity(self, data)


    def EventTag(self, data=None) -> "EventTagEntity":
        """Entity factory: client.EventTag().list() / client.EventTag().load({"id": ...})."""
        from entity.event_tag_entity import EventTagEntity
        return EventTagEntity(self, data)


    def EventTagAssignment(self, data=None) -> "EventTagAssignmentEntity":
        """Entity factory: client.EventTagAssignment().list() / client.EventTagAssignment().load({"id": ...})."""
        from entity.event_tag_assignment_entity import EventTagAssignmentEntity
        return EventTagAssignmentEntity(self, data)


    def Guest(self, data=None) -> "GuestEntity":
        """Entity factory: client.Guest().list() / client.Guest().load({"id": ...})."""
        from entity.guest_entity import GuestEntity
        return GuestEntity(self, data)


    def GuestInvite(self, data=None) -> "GuestInviteEntity":
        """Entity factory: client.GuestInvite().list() / client.GuestInvite().load({"id": ...})."""
        from entity.guest_invite_entity import GuestInviteEntity
        return GuestInviteEntity(self, data)


    def GuestTicket(self, data=None) -> "GuestTicketEntity":
        """Entity factory: client.GuestTicket().list() / client.GuestTicket().load({"id": ...})."""
        from entity.guest_ticket_entity import GuestTicketEntity
        return GuestTicketEntity(self, data)


    def Host(self, data=None) -> "HostEntity":
        """Entity factory: client.Host().list() / client.Host().load({"id": ...})."""
        from entity.host_entity import HostEntity
        return HostEntity(self, data)


    def ImageUpload(self, data=None) -> "ImageUploadEntity":
        """Entity factory: client.ImageUpload().list() / client.ImageUpload().load({"id": ...})."""
        from entity.image_upload_entity import ImageUploadEntity
        return ImageUploadEntity(self, data)


    def Member(self, data=None) -> "MemberEntity":
        """Entity factory: client.Member().list() / client.Member().load({"id": ...})."""
        from entity.member_entity import MemberEntity
        return MemberEntity(self, data)


    def MembershipTier(self, data=None) -> "MembershipTierEntity":
        """Entity factory: client.MembershipTier().list() / client.MembershipTier().load({"id": ...})."""
        from entity.membership_tier_entity import MembershipTierEntity
        return MembershipTierEntity(self, data)


    def OrganizationAdmin(self, data=None) -> "OrganizationAdminEntity":
        """Entity factory: client.OrganizationAdmin().list() / client.OrganizationAdmin().load({"id": ...})."""
        from entity.organization_admin_entity import OrganizationAdminEntity
        return OrganizationAdminEntity(self, data)


    def OrganizationCalendar(self, data=None) -> "OrganizationCalendarEntity":
        """Entity factory: client.OrganizationCalendar().list() / client.OrganizationCalendar().load({"id": ...})."""
        from entity.organization_calendar_entity import OrganizationCalendarEntity
        return OrganizationCalendarEntity(self, data)


    def OrganizationEvent(self, data=None) -> "OrganizationEventEntity":
        """Entity factory: client.OrganizationEvent().list() / client.OrganizationEvent().load({"id": ...})."""
        from entity.organization_event_entity import OrganizationEventEntity
        return OrganizationEventEntity(self, data)


    def OrganizationEventTransfer(self, data=None) -> "OrganizationEventTransferEntity":
        """Entity factory: client.OrganizationEventTransfer().list() / client.OrganizationEventTransfer().load({"id": ...})."""
        from entity.organization_event_transfer_entity import OrganizationEventTransferEntity
        return OrganizationEventTransferEntity(self, data)


    def TicketType(self, data=None) -> "TicketTypeEntity":
        """Entity factory: client.TicketType().list() / client.TicketType().load({"id": ...})."""
        from entity.ticket_type_entity import TicketTypeEntity
        return TicketTypeEntity(self, data)


    def User(self, data=None) -> "UserEntity":
        """Entity factory: client.User().list() / client.User().load({"id": ...})."""
        from entity.user_entity import UserEntity
        return UserEntity(self, data)


    def Webhook(self, data=None) -> "WebhookEntity":
        """Entity factory: client.Webhook().list() / client.Webhook().load({"id": ...})."""
        from entity.webhook_entity import WebhookEntity
        return WebhookEntity(self, data)



    @classmethod
    def test(cls, testopts=None, sdkopts=None) -> "LumaSDK":
        if sdkopts is None:
            sdkopts = {}
        sdkopts = vs.clone(sdkopts)
        if not isinstance(sdkopts, dict):
            sdkopts = {}

        if testopts is None:
            testopts = {}
        testopts = vs.clone(testopts)
        if not isinstance(testopts, dict):
            testopts = {}
        testopts["active"] = True

        vs.setpath(sdkopts, "feature.test", testopts)

        sdk = cls(sdkopts)
        sdk.mode = "test"

        return sdk


from typing import TYPE_CHECKING

if TYPE_CHECKING:
    from entity.calendar_entity import CalendarEntity
    from entity.calendar_admin_entity import CalendarAdminEntity
    from entity.calendar_coupon_entity import CalendarCouponEntity
    from entity.calendar_event_entity import CalendarEventEntity
    from entity.calendar_event_approval_entity import CalendarEventApprovalEntity
    from entity.calendar_event_rejection_entity import CalendarEventRejectionEntity
    from entity.contact_entity import ContactEntity
    from entity.contact_block_entity import ContactBlockEntity
    from entity.contact_restore_entity import ContactRestoreEntity
    from entity.contact_tag_entity import ContactTagEntity
    from entity.contact_tag_assignment_entity import ContactTagAssignmentEntity
    from entity.entity_lookup_entity import EntityLookupEntity
    from entity.event_entity import EventEntity
    from entity.event_cancel_request_entity import EventCancelRequestEntity
    from entity.event_coupon_entity import EventCouponEntity
    from entity.event_tag_entity import EventTagEntity
    from entity.event_tag_assignment_entity import EventTagAssignmentEntity
    from entity.guest_entity import GuestEntity
    from entity.guest_invite_entity import GuestInviteEntity
    from entity.guest_ticket_entity import GuestTicketEntity
    from entity.host_entity import HostEntity
    from entity.image_upload_entity import ImageUploadEntity
    from entity.member_entity import MemberEntity
    from entity.membership_tier_entity import MembershipTierEntity
    from entity.organization_admin_entity import OrganizationAdminEntity
    from entity.organization_calendar_entity import OrganizationCalendarEntity
    from entity.organization_event_entity import OrganizationEventEntity
    from entity.organization_event_transfer_entity import OrganizationEventTransferEntity
    from entity.ticket_type_entity import TicketTypeEntity
    from entity.user_entity import UserEntity
    from entity.webhook_entity import WebhookEntity
