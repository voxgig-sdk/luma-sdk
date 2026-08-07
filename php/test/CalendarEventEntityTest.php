<?php
declare(strict_types=1);

// CalendarEvent entity test

require_once __DIR__ . '/../luma_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class CalendarEventEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = LumaSDK::test(null, null);
        $ent = $testsdk->CalendarEvent(null);
        $this->assertNotNull($ent);
    }

    // Feature #4: the entity stream(action, ...) method runs the op pipeline
    // and yields result items. With the streaming feature active it yields the
    // feature's incremental output; otherwise it falls back to the materialised
    // list so stream always yields.
    public function test_stream(): void
    {
        $seed = [
            "entity" => [
                "calendar_event" => [
                    "s1" => ["id" => "s1"],
                    "s2" => ["id" => "s2"],
                    "s3" => ["id" => "s3"],
                ],
            ],
        ];

        // Fallback: streaming inactive -> yields the materialised list items.
        $base = LumaSDK::test($seed, null);
        $seen = iterator_to_array($base->CalendarEvent(null)->stream("list", null, null), false);
        $this->assertCount(3, $seen);

        // Inbound: streaming active -> yields each item from the feature.
        $cfg = LumaConfig::make_config();
        if (isset($cfg["feature"]) && is_array($cfg["feature"]) && isset($cfg["feature"]["streaming"])) {
            $sdk = LumaSDK::test($seed, ["feature" => ["streaming" => ["active" => true]]]);
            $got = [];
            foreach ($sdk->CalendarEvent(null)->stream("list", null, null) as $item) {
                if (is_array($item) && array_is_list($item)) {
                    foreach ($item as $sub) {
                        $got[] = $sub;
                    }
                } else {
                    $got[] = $item;
                }
            }
            $this->assertCount(3, $got);
        }
    }

    public function test_basic_flow(): void
    {
        $setup = calendar_event_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["create", "list", "load"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "calendar_event." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set LUMA_TEST_CALENDAR_EVENT_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // CREATE
        $calendar_event_ref01_ent = $client->CalendarEvent(null);
        $calendar_event_ref01_data = Helpers::to_map(Vs::getprop(
            Vs::getpath($setup["data"], "new.calendar_event"), "calendar_event_ref01"));

        $calendar_event_ref01_data_result = $calendar_event_ref01_ent->create($calendar_event_ref01_data, null);
        $calendar_event_ref01_data = Helpers::to_map($calendar_event_ref01_data_result);
        $this->assertNotNull($calendar_event_ref01_data);
        $this->assertNotNull($calendar_event_ref01_data["id"]);

        // LIST
        $calendar_event_ref01_match = [];

        $calendar_event_ref01_list_result = $calendar_event_ref01_ent->list($calendar_event_ref01_match, null);
        $this->assertIsArray($calendar_event_ref01_list_result);

        $found_item = sdk_select(
            Runner::entity_list_to_data($calendar_event_ref01_list_result),
            ["id" => $calendar_event_ref01_data["id"]]);
        $this->assertNotEmpty($found_item);

        // LOAD
        $calendar_event_ref01_match_dt0 = [
            "id" => $calendar_event_ref01_data["id"],
        ];
        $calendar_event_ref01_data_dt0_loaded = $calendar_event_ref01_ent->load($calendar_event_ref01_match_dt0, null);
        $calendar_event_ref01_data_dt0_load_result = Helpers::to_map($calendar_event_ref01_data_dt0_loaded);
        $this->assertNotNull($calendar_event_ref01_data_dt0_load_result);
        $this->assertEquals($calendar_event_ref01_data_dt0_load_result["id"], $calendar_event_ref01_data["id"]);

    }
}

function calendar_event_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/calendar_event/CalendarEventTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = LumaSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["calendar_event01", "calendar_event02", "calendar_event03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("LUMA_TEST_CALENDAR_EVENT_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "LUMA_TEST_CALENDAR_EVENT_ENTID" => $idmap,
        "LUMA_TEST_LIVE" => "FALSE",
        "LUMA_TEST_EXPLAIN" => "FALSE",
        "LUMA_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["LUMA_TEST_CALENDAR_EVENT_ENTID"]);
    if ($idmap_resolved === null) {
        $idmap_resolved = Helpers::to_map($idmap);
    }

    if ($env["LUMA_TEST_LIVE"] === "TRUE") {
        $merged_opts = Vs::merge([
            [
                "apikey" => $env["LUMA_APIKEY"],
            ],
            $extra ?? [],
        ]);
        $client = new LumaSDK(Helpers::to_map($merged_opts));
    }

    $live = $env["LUMA_TEST_LIVE"] === "TRUE";
    return [
        "client" => $client,
        "data" => $entity_data,
        "idmap" => $idmap_resolved,
        "env" => $env,
        "explain" => $env["LUMA_TEST_EXPLAIN"] === "TRUE",
        "live" => $live,
        "synthetic_only" => $live && !$idmap_overridden,
        "now" => (int)(microtime(true) * 1000),
    ];
}
