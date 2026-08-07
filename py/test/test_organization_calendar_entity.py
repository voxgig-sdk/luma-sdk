# OrganizationCalendar entity test

import json
import os
import time

import pytest

from utility.voxgig_struct import voxgig_struct as vs
from luma_sdk import LumaSDK
from core import helpers

_TEST_DIR = os.path.dirname(os.path.abspath(__file__))
from test import runner


class TestOrganizationCalendarEntity:

    def test_should_create_instance(self):
        testsdk = LumaSDK.test(None, None)
        ent = testsdk.OrganizationCalendar(None)
        assert ent is not None

    def test_should_stream(self):
        # Feature #4: the entity stream(action, ...) method runs the op
        # pipeline and yields result items. With the streaming feature active
        # it yields the feature's incremental output; otherwise it falls back
        # to the materialised list so stream always yields.
        seed = {
            "entity": {
                "organization_calendar": {
                    "s1": {"id": "s1"},
                    "s2": {"id": "s2"},
                    "s3": {"id": "s3"},
                }
            }
        }

        # Fallback: streaming inactive -> yields the materialised list items.
        base = LumaSDK.test(seed, None)
        seen = list(base.OrganizationCalendar(None).stream("list", None, None))
        assert len(seen) == 3

        # Inbound: streaming active -> yields each item from the feature.
        from config import make_config
        cfg = make_config()
        if isinstance(cfg.get("feature"), dict) and "streaming" in cfg["feature"]:
            sdk = LumaSDK.test(
                seed, {"feature": {"streaming": {"active": True}}})
            got = []
            for item in sdk.OrganizationCalendar(None).stream("list", None, None):
                if isinstance(item, list):
                    got.extend(item)
                else:
                    got.append(item)
            assert len(got) == 3

    def test_should_run_basic_flow(self):
        setup = _organization_calendar_basic_setup(None)
        # Per-op sdk-test-control.json skip — basic test exercises a flow with
        # multiple ops; skipping any one skips the whole flow (steps depend
        # on each other).
        _live = setup.get("live", False)
        for _op in ["create", "list"]:
            _skip, _reason = runner.is_control_skipped("entityOp", "organization_calendar." + _op, "live" if _live else "unit")
            if _skip:
                pytest.skip(_reason or "skipped via sdk-test-control.json")
                return
        # The basic flow consumes synthetic IDs from the fixture. In live mode
        # without an *_ENTID env override, those IDs hit the live API and 4xx.
        if setup.get("synthetic_only"):
            pytest.skip("live entity test uses synthetic IDs from fixture — "
                        "set LUMA_TEST_ORGANIZATION_CALENDAR_ENTID JSON to run live")
        client = setup["client"]

        # CREATE
        organization_calendar_ref01_ent = client.OrganizationCalendar(None)
        organization_calendar_ref01_data = helpers.to_map(vs.getprop(
            vs.getpath(setup["data"], "new.organization_calendar"), "organization_calendar_ref01"))

        organization_calendar_ref01_data = helpers.to_map(organization_calendar_ref01_ent.create(organization_calendar_ref01_data, None))
        assert organization_calendar_ref01_data is not None
        assert organization_calendar_ref01_data["id"] is not None

        # LIST
        organization_calendar_ref01_match = {}

        organization_calendar_ref01_list_result = organization_calendar_ref01_ent.list(organization_calendar_ref01_match, None)
        assert isinstance(organization_calendar_ref01_list_result, list)

        found_item = vs.select(
            runner.entity_list_to_data(organization_calendar_ref01_list_result),
            {"id": organization_calendar_ref01_data["id"]})
        assert not vs.isempty(found_item)



def _organization_calendar_basic_setup(extra):
    runner.load_env_local()

    entity_data_file = os.path.join(_TEST_DIR, "../../.sdk/test/entity/organization_calendar/OrganizationCalendarTestData.json")
    with open(entity_data_file, "r") as f:
        entity_data_source = f.read()

    entity_data = json.loads(entity_data_source)

    options = {}
    options["entity"] = entity_data.get("existing")

    client = LumaSDK.test(options, extra)

    # Generate idmap via transform.
    idmap = vs.transform(
        ["organization_calendar01", "organization_calendar02", "organization_calendar03"],
        {
            "`$PACK`": ["", {
                "`$KEY`": "`$COPY`",
                "`$VAL`": ["`$FORMAT`", "upper", "`$COPY`"],
            }],
        }
    )

    # Detect ENTID env override before envOverride consumes it. When live
    # mode is on without a real override, the basic test runs against synthetic
    # IDs from the fixture and 4xx's. We surface this so the test can skip.
    _entid_env_raw = os.environ.get(
        "LUMA_TEST_ORGANIZATION_CALENDAR_ENTID")
    _idmap_overridden = _entid_env_raw is not None and _entid_env_raw.strip().startswith("{")

    env = runner.env_override({
        "LUMA_TEST_ORGANIZATION_CALENDAR_ENTID": idmap,
        "LUMA_TEST_LIVE": "FALSE",
        "LUMA_TEST_EXPLAIN": "FALSE",
        "LUMA_APIKEY": "NONE",
    })

    idmap_resolved = helpers.to_map(
        env.get("LUMA_TEST_ORGANIZATION_CALENDAR_ENTID"))
    if idmap_resolved is None:
        idmap_resolved = helpers.to_map(idmap)

    if env.get("LUMA_TEST_LIVE") == "TRUE":
        merged_opts = vs.merge([
            {
                "apikey": env.get("LUMA_APIKEY"),
            },
            extra or {},
        ])
        client = LumaSDK(helpers.to_map(merged_opts))

    _live = env.get("LUMA_TEST_LIVE") == "TRUE"
    return {
        "client": client,
        "data": entity_data,
        "idmap": idmap_resolved,
        "env": env,
        "explain": env.get("LUMA_TEST_EXPLAIN") == "TRUE",
        "live": _live,
        "synthetic_only": _live and not _idmap_overridden,
        "now": int(time.time() * 1000),
    }
