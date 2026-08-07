<?php
declare(strict_types=1);

// Luma SDK

require_once __DIR__ . '/utility/struct/Struct.php';
require_once __DIR__ . '/core/UtilityType.php';
require_once __DIR__ . '/core/Spec.php';
require_once __DIR__ . '/core/Helpers.php';

// Load utility registration
require_once __DIR__ . '/utility/Register.php';

// Load config and features
require_once __DIR__ . '/config.php';
require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/features.php';

use Voxgig\Struct\Struct;

// Features record diagnostic state on the client as dynamic properties
// (_retry, _cache, _metrics, ...); allow them explicitly (PHP 8.2+
// deprecates implicit dynamic properties).
#[\AllowDynamicProperties]
class LumaSDK
{
    public string $mode;
    public array $features;
    public ?array $options;

    private $_utility;
    private $_rootctx;

    public function __construct(array $options = [])
    {
        $this->mode = "live";
        $this->features = [];
        $this->options = null;

        $utility = new LumaUtility();
        $this->_utility = $utility;

        $config = LumaConfig::make_config();

        $this->_rootctx = ($utility->make_context)([
            "client" => $this,
            "utility" => $utility,
            "config" => $config,
            "options" => $options ?? [],
            "shared" => [],
        ], null);

        $this->options = ($utility->make_options)($this->_rootctx);

        if (Struct::getpath($this->options, "feature.test.active") === true) {
            $this->mode = "test";
        }

        $this->_rootctx->options = $this->options;

        // Add features in the resolved order (make_options puts an explicit
        // list order first, else defaults to test-first). Ordering matters: the
        // `test` feature installs the base mock transport and the transport
        // features (retry/cache/netsim/proxy/ratelimit) wrap whatever is
        // current, so `test` must be added before them to sit at the base.
        $feature_opts = LumaHelpers::to_map(Struct::getprop($this->options, "feature"));
        if ($feature_opts) {
            $featureorder = Struct::getpath($this->options, "__derived__.featureorder");
            if (is_array($featureorder)) {
                foreach ($featureorder as $fname) {
                    $fopts = LumaHelpers::to_map($feature_opts[$fname] ?? null);
                    if ($fopts && isset($fopts["active"]) && $fopts["active"] === true) {
                        ($utility->feature_add)($this->_rootctx, LumaFeatures::make_feature($fname));
                    }
                }
            }
        }

        // Add extension features.
        $extend_val = Struct::getprop($this->options, "extend");
        if (is_array($extend_val)) {
            foreach ($extend_val as $f) {
                if (is_object($f) && method_exists($f, 'get_name')) {
                    ($utility->feature_add)($this->_rootctx, $f);
                }
            }
        }

        // Initialize features.
        foreach ($this->features as $f) {
            ($utility->feature_init)($this->_rootctx, $f);
        }

        ($utility->feature_hook)($this->_rootctx, "PostConstruct");
    }

    public function options_map(): array
    {
        $out = Struct::clone($this->options);
        return is_array($out) ? $out : [];
    }

    public function get_utility()
    {
        return LumaUtility::copy($this->_utility);
    }

    public function get_root_ctx()
    {
        return $this->_rootctx;
    }

    public function prepare(array $fetchargs = []): mixed
    {
        $utility = $this->_utility;
        $fetchargs = $fetchargs ?? [];

        $ctrl = LumaHelpers::to_map(Struct::getprop($fetchargs, "ctrl")) ?? [];

        $ctx = ($utility->make_context)([
            "opname" => "prepare",
            "ctrl" => $ctrl,
        ], $this->_rootctx);

        $opts = $this->options;
        $path = Struct::getprop($fetchargs, "path") ?? "";
        $path = is_string($path) ? $path : "";
        $method_val = Struct::getprop($fetchargs, "method") ?? "GET";
        $method_val = is_string($method_val) ? $method_val : "GET";
        $params = LumaHelpers::to_map(Struct::getprop($fetchargs, "params")) ?? [];
        $query = LumaHelpers::to_map(Struct::getprop($fetchargs, "query")) ?? [];
        $headers = ($utility->prepare_headers)($ctx);

        $base = Struct::getprop($opts, "base") ?? "";
        $base = is_string($base) ? $base : "";
        $prefix = Struct::getprop($opts, "prefix") ?? "";
        $prefix = is_string($prefix) ? $prefix : "";
        $suffix = Struct::getprop($opts, "suffix") ?? "";
        $suffix = is_string($suffix) ? $suffix : "";

        $ctx->spec = new LumaSpec([
            "base" => $base, "prefix" => $prefix, "suffix" => $suffix,
            "path" => $path, "method" => $method_val,
            "params" => $params, "query" => $query, "headers" => $headers,
            "body" => Struct::getprop($fetchargs, "body"),
            "step" => "start",
        ]);

        // Merge user-provided headers.
        $uh = Struct::getprop($fetchargs, "headers");
        if (is_array($uh)) {
            foreach ($uh as $k => $v) {
                $ctx->spec->headers[$k] = $v;
            }
        }

        [$_, $err] = ($utility->prepare_auth)($ctx);
        if ($err) {
            return ($utility->make_error)($ctx, $err);
        }

        [$fetchdef, $fd_err] = ($utility->make_fetch_def)($ctx);
        if ($fd_err) {
            return ($utility->make_error)($ctx, $fd_err);
        }
        return $fetchdef;
    }

    public function direct(array $fetchargs = []): mixed
    {
        $utility = $this->_utility;

        // direct() is the raw-HTTP escape hatch: it never throws, it returns
        // an {ok, err, ...} dict. prepare() now raises on error, so catch it
        // and surface the failure through the dict instead.
        try {
            $fetchdef = $this->prepare($fetchargs);
        } catch (\Throwable $err) {
            return ["ok" => false, "err" => $err];
        }

        $fetchargs = $fetchargs ?? [];
        $ctrl = LumaHelpers::to_map(Struct::getprop($fetchargs, "ctrl")) ?? [];

        $ctx = ($utility->make_context)([
            "opname" => "direct",
            "ctrl" => $ctrl,
        ], $this->_rootctx);

        $url = $fetchdef["url"] ?? "";
        [$fetched, $fetch_err] = ($utility->fetcher)($ctx, $url, $fetchdef);

        if ($fetch_err) {
            return ["ok" => false, "err" => $fetch_err];
        }

        if ($fetched === null) {
            return [
                "ok" => false,
                "err" => $ctx->make_error("direct_no_response", "response: undefined"),
            ];
        }

        if (is_array($fetched)) {
            $status = LumaHelpers::to_int(Struct::getprop($fetched, "status"));
            $headers = Struct::getprop($fetched, "headers") ?? [];

            // No-body responses (204, 304) and explicit zero content-length
            // must skip JSON parsing — calling json() on an empty body errors.
            $content_length = is_array($headers) ? ($headers["content-length"] ?? null) : null;
            $no_body = $status === 204 || $status === 304 || (string)$content_length === "0";

            $json_data = null;
            if (!$no_body) {
                $jf = Struct::getprop($fetched, "json");
                if (is_callable($jf)) {
                    try {
                        $json_data = $jf();
                    } catch (\Throwable $e) {
                        // Non-JSON body — leave data null but keep status/ok.
                        $json_data = null;
                    }
                }
            }

            return [
                "ok" => $status >= 200 && $status < 300,
                "status" => $status,
                "headers" => Struct::getprop($fetched, "headers"),
                "data" => $json_data,
            ];
        }

        return [
            "ok" => false,
            "err" => $ctx->make_error("direct_invalid", "invalid response type"),
        ];
    }


    private $_calendar = null;

    // Canonical facade: $client->Calendar()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->calendar()
    // resolves here too.
    public function Calendar($data = null)
    {
        require_once __DIR__ . '/entity/calendar_entity.php';
        if ($data === null) {
            if ($this->_calendar === null) {
                $this->_calendar = new CalendarEntity($this, null);
            }
            return $this->_calendar;
        }
        return new CalendarEntity($this, $data);
    }


    private $_calendar_admin = null;

    // Canonical facade: $client->CalendarAdmin()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->calendar_admin()
    // resolves here too.
    public function CalendarAdmin($data = null)
    {
        require_once __DIR__ . '/entity/calendar_admin_entity.php';
        if ($data === null) {
            if ($this->_calendar_admin === null) {
                $this->_calendar_admin = new CalendarAdminEntity($this, null);
            }
            return $this->_calendar_admin;
        }
        return new CalendarAdminEntity($this, $data);
    }


    private $_calendar_coupon = null;

    // Canonical facade: $client->CalendarCoupon()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->calendar_coupon()
    // resolves here too.
    public function CalendarCoupon($data = null)
    {
        require_once __DIR__ . '/entity/calendar_coupon_entity.php';
        if ($data === null) {
            if ($this->_calendar_coupon === null) {
                $this->_calendar_coupon = new CalendarCouponEntity($this, null);
            }
            return $this->_calendar_coupon;
        }
        return new CalendarCouponEntity($this, $data);
    }


    private $_calendar_event = null;

    // Canonical facade: $client->CalendarEvent()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->calendar_event()
    // resolves here too.
    public function CalendarEvent($data = null)
    {
        require_once __DIR__ . '/entity/calendar_event_entity.php';
        if ($data === null) {
            if ($this->_calendar_event === null) {
                $this->_calendar_event = new CalendarEventEntity($this, null);
            }
            return $this->_calendar_event;
        }
        return new CalendarEventEntity($this, $data);
    }


    private $_calendar_event_approval = null;

    // Canonical facade: $client->CalendarEventApproval()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->calendar_event_approval()
    // resolves here too.
    public function CalendarEventApproval($data = null)
    {
        require_once __DIR__ . '/entity/calendar_event_approval_entity.php';
        if ($data === null) {
            if ($this->_calendar_event_approval === null) {
                $this->_calendar_event_approval = new CalendarEventApprovalEntity($this, null);
            }
            return $this->_calendar_event_approval;
        }
        return new CalendarEventApprovalEntity($this, $data);
    }


    private $_calendar_event_rejection = null;

    // Canonical facade: $client->CalendarEventRejection()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->calendar_event_rejection()
    // resolves here too.
    public function CalendarEventRejection($data = null)
    {
        require_once __DIR__ . '/entity/calendar_event_rejection_entity.php';
        if ($data === null) {
            if ($this->_calendar_event_rejection === null) {
                $this->_calendar_event_rejection = new CalendarEventRejectionEntity($this, null);
            }
            return $this->_calendar_event_rejection;
        }
        return new CalendarEventRejectionEntity($this, $data);
    }


    private $_contact = null;

    // Canonical facade: $client->Contact()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->contact()
    // resolves here too.
    public function Contact($data = null)
    {
        require_once __DIR__ . '/entity/contact_entity.php';
        if ($data === null) {
            if ($this->_contact === null) {
                $this->_contact = new ContactEntity($this, null);
            }
            return $this->_contact;
        }
        return new ContactEntity($this, $data);
    }


    private $_contact_block = null;

    // Canonical facade: $client->ContactBlock()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->contact_block()
    // resolves here too.
    public function ContactBlock($data = null)
    {
        require_once __DIR__ . '/entity/contact_block_entity.php';
        if ($data === null) {
            if ($this->_contact_block === null) {
                $this->_contact_block = new ContactBlockEntity($this, null);
            }
            return $this->_contact_block;
        }
        return new ContactBlockEntity($this, $data);
    }


    private $_contact_restore = null;

    // Canonical facade: $client->ContactRestore()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->contact_restore()
    // resolves here too.
    public function ContactRestore($data = null)
    {
        require_once __DIR__ . '/entity/contact_restore_entity.php';
        if ($data === null) {
            if ($this->_contact_restore === null) {
                $this->_contact_restore = new ContactRestoreEntity($this, null);
            }
            return $this->_contact_restore;
        }
        return new ContactRestoreEntity($this, $data);
    }


    private $_contact_tag = null;

    // Canonical facade: $client->ContactTag()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->contact_tag()
    // resolves here too.
    public function ContactTag($data = null)
    {
        require_once __DIR__ . '/entity/contact_tag_entity.php';
        if ($data === null) {
            if ($this->_contact_tag === null) {
                $this->_contact_tag = new ContactTagEntity($this, null);
            }
            return $this->_contact_tag;
        }
        return new ContactTagEntity($this, $data);
    }


    private $_contact_tag_assignment = null;

    // Canonical facade: $client->ContactTagAssignment()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->contact_tag_assignment()
    // resolves here too.
    public function ContactTagAssignment($data = null)
    {
        require_once __DIR__ . '/entity/contact_tag_assignment_entity.php';
        if ($data === null) {
            if ($this->_contact_tag_assignment === null) {
                $this->_contact_tag_assignment = new ContactTagAssignmentEntity($this, null);
            }
            return $this->_contact_tag_assignment;
        }
        return new ContactTagAssignmentEntity($this, $data);
    }


    private $_entity_lookup = null;

    // Canonical facade: $client->EntityLookup()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->entity_lookup()
    // resolves here too.
    public function EntityLookup($data = null)
    {
        require_once __DIR__ . '/entity/entity_lookup_entity.php';
        if ($data === null) {
            if ($this->_entity_lookup === null) {
                $this->_entity_lookup = new EntityLookupEntity($this, null);
            }
            return $this->_entity_lookup;
        }
        return new EntityLookupEntity($this, $data);
    }


    private $_event = null;

    // Canonical facade: $client->Event()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->event()
    // resolves here too.
    public function Event($data = null)
    {
        require_once __DIR__ . '/entity/event_entity.php';
        if ($data === null) {
            if ($this->_event === null) {
                $this->_event = new EventEntity($this, null);
            }
            return $this->_event;
        }
        return new EventEntity($this, $data);
    }


    private $_event_cancel_request = null;

    // Canonical facade: $client->EventCancelRequest()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->event_cancel_request()
    // resolves here too.
    public function EventCancelRequest($data = null)
    {
        require_once __DIR__ . '/entity/event_cancel_request_entity.php';
        if ($data === null) {
            if ($this->_event_cancel_request === null) {
                $this->_event_cancel_request = new EventCancelRequestEntity($this, null);
            }
            return $this->_event_cancel_request;
        }
        return new EventCancelRequestEntity($this, $data);
    }


    private $_event_coupon = null;

    // Canonical facade: $client->EventCoupon()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->event_coupon()
    // resolves here too.
    public function EventCoupon($data = null)
    {
        require_once __DIR__ . '/entity/event_coupon_entity.php';
        if ($data === null) {
            if ($this->_event_coupon === null) {
                $this->_event_coupon = new EventCouponEntity($this, null);
            }
            return $this->_event_coupon;
        }
        return new EventCouponEntity($this, $data);
    }


    private $_event_tag = null;

    // Canonical facade: $client->EventTag()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->event_tag()
    // resolves here too.
    public function EventTag($data = null)
    {
        require_once __DIR__ . '/entity/event_tag_entity.php';
        if ($data === null) {
            if ($this->_event_tag === null) {
                $this->_event_tag = new EventTagEntity($this, null);
            }
            return $this->_event_tag;
        }
        return new EventTagEntity($this, $data);
    }


    private $_event_tag_assignment = null;

    // Canonical facade: $client->EventTagAssignment()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->event_tag_assignment()
    // resolves here too.
    public function EventTagAssignment($data = null)
    {
        require_once __DIR__ . '/entity/event_tag_assignment_entity.php';
        if ($data === null) {
            if ($this->_event_tag_assignment === null) {
                $this->_event_tag_assignment = new EventTagAssignmentEntity($this, null);
            }
            return $this->_event_tag_assignment;
        }
        return new EventTagAssignmentEntity($this, $data);
    }


    private $_guest = null;

    // Canonical facade: $client->Guest()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->guest()
    // resolves here too.
    public function Guest($data = null)
    {
        require_once __DIR__ . '/entity/guest_entity.php';
        if ($data === null) {
            if ($this->_guest === null) {
                $this->_guest = new GuestEntity($this, null);
            }
            return $this->_guest;
        }
        return new GuestEntity($this, $data);
    }


    private $_guest_invite = null;

    // Canonical facade: $client->GuestInvite()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->guest_invite()
    // resolves here too.
    public function GuestInvite($data = null)
    {
        require_once __DIR__ . '/entity/guest_invite_entity.php';
        if ($data === null) {
            if ($this->_guest_invite === null) {
                $this->_guest_invite = new GuestInviteEntity($this, null);
            }
            return $this->_guest_invite;
        }
        return new GuestInviteEntity($this, $data);
    }


    private $_guest_ticket = null;

    // Canonical facade: $client->GuestTicket()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->guest_ticket()
    // resolves here too.
    public function GuestTicket($data = null)
    {
        require_once __DIR__ . '/entity/guest_ticket_entity.php';
        if ($data === null) {
            if ($this->_guest_ticket === null) {
                $this->_guest_ticket = new GuestTicketEntity($this, null);
            }
            return $this->_guest_ticket;
        }
        return new GuestTicketEntity($this, $data);
    }


    private $_host = null;

    // Canonical facade: $client->Host()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->host()
    // resolves here too.
    public function Host($data = null)
    {
        require_once __DIR__ . '/entity/host_entity.php';
        if ($data === null) {
            if ($this->_host === null) {
                $this->_host = new HostEntity($this, null);
            }
            return $this->_host;
        }
        return new HostEntity($this, $data);
    }


    private $_image_upload = null;

    // Canonical facade: $client->ImageUpload()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->image_upload()
    // resolves here too.
    public function ImageUpload($data = null)
    {
        require_once __DIR__ . '/entity/image_upload_entity.php';
        if ($data === null) {
            if ($this->_image_upload === null) {
                $this->_image_upload = new ImageUploadEntity($this, null);
            }
            return $this->_image_upload;
        }
        return new ImageUploadEntity($this, $data);
    }


    private $_member = null;

    // Canonical facade: $client->Member()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->member()
    // resolves here too.
    public function Member($data = null)
    {
        require_once __DIR__ . '/entity/member_entity.php';
        if ($data === null) {
            if ($this->_member === null) {
                $this->_member = new MemberEntity($this, null);
            }
            return $this->_member;
        }
        return new MemberEntity($this, $data);
    }


    private $_membership_tier = null;

    // Canonical facade: $client->MembershipTier()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->membership_tier()
    // resolves here too.
    public function MembershipTier($data = null)
    {
        require_once __DIR__ . '/entity/membership_tier_entity.php';
        if ($data === null) {
            if ($this->_membership_tier === null) {
                $this->_membership_tier = new MembershipTierEntity($this, null);
            }
            return $this->_membership_tier;
        }
        return new MembershipTierEntity($this, $data);
    }


    private $_organization_admin = null;

    // Canonical facade: $client->OrganizationAdmin()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->organization_admin()
    // resolves here too.
    public function OrganizationAdmin($data = null)
    {
        require_once __DIR__ . '/entity/organization_admin_entity.php';
        if ($data === null) {
            if ($this->_organization_admin === null) {
                $this->_organization_admin = new OrganizationAdminEntity($this, null);
            }
            return $this->_organization_admin;
        }
        return new OrganizationAdminEntity($this, $data);
    }


    private $_organization_calendar = null;

    // Canonical facade: $client->OrganizationCalendar()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->organization_calendar()
    // resolves here too.
    public function OrganizationCalendar($data = null)
    {
        require_once __DIR__ . '/entity/organization_calendar_entity.php';
        if ($data === null) {
            if ($this->_organization_calendar === null) {
                $this->_organization_calendar = new OrganizationCalendarEntity($this, null);
            }
            return $this->_organization_calendar;
        }
        return new OrganizationCalendarEntity($this, $data);
    }


    private $_organization_event = null;

    // Canonical facade: $client->OrganizationEvent()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->organization_event()
    // resolves here too.
    public function OrganizationEvent($data = null)
    {
        require_once __DIR__ . '/entity/organization_event_entity.php';
        if ($data === null) {
            if ($this->_organization_event === null) {
                $this->_organization_event = new OrganizationEventEntity($this, null);
            }
            return $this->_organization_event;
        }
        return new OrganizationEventEntity($this, $data);
    }


    private $_organization_event_transfer = null;

    // Canonical facade: $client->OrganizationEventTransfer()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->organization_event_transfer()
    // resolves here too.
    public function OrganizationEventTransfer($data = null)
    {
        require_once __DIR__ . '/entity/organization_event_transfer_entity.php';
        if ($data === null) {
            if ($this->_organization_event_transfer === null) {
                $this->_organization_event_transfer = new OrganizationEventTransferEntity($this, null);
            }
            return $this->_organization_event_transfer;
        }
        return new OrganizationEventTransferEntity($this, $data);
    }


    private $_ticket_type = null;

    // Canonical facade: $client->TicketType()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->ticket_type()
    // resolves here too.
    public function TicketType($data = null)
    {
        require_once __DIR__ . '/entity/ticket_type_entity.php';
        if ($data === null) {
            if ($this->_ticket_type === null) {
                $this->_ticket_type = new TicketTypeEntity($this, null);
            }
            return $this->_ticket_type;
        }
        return new TicketTypeEntity($this, $data);
    }


    private $_user = null;

    // Canonical facade: $client->User()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->user()
    // resolves here too.
    public function User($data = null)
    {
        require_once __DIR__ . '/entity/user_entity.php';
        if ($data === null) {
            if ($this->_user === null) {
                $this->_user = new UserEntity($this, null);
            }
            return $this->_user;
        }
        return new UserEntity($this, $data);
    }


    private $_webhook = null;

    // Canonical facade: $client->Webhook()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->webhook()
    // resolves here too.
    public function Webhook($data = null)
    {
        require_once __DIR__ . '/entity/webhook_entity.php';
        if ($data === null) {
            if ($this->_webhook === null) {
                $this->_webhook = new WebhookEntity($this, null);
            }
            return $this->_webhook;
        }
        return new WebhookEntity($this, $data);
    }



    public static function test(?array $testopts = null, ?array $sdkopts = null): self
    {
        $sdkopts = $sdkopts ?? [];
        $sdkopts = Struct::clone($sdkopts);
        $sdkopts = is_array($sdkopts) ? $sdkopts : [];

        $testopts = $testopts ?? [];
        $testopts = Struct::clone($testopts);
        $testopts = is_array($testopts) ? $testopts : [];
        $testopts["active"] = true;

        if (!isset($sdkopts["feature"])) {
            $sdkopts["feature"] = [];
        }
        $sdkopts["feature"]["test"] = $testopts;

        $sdk = new LumaSDK($sdkopts);
        $sdk->mode = "test";
        return $sdk;
    }
}
