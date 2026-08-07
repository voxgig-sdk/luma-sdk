<?php
declare(strict_types=1);

// Calendar entity test

require_once __DIR__ . '/../luma_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class CalendarEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = LumaSDK::test(null, null);
        $ent = $testsdk->Calendar(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = calendar_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["update", "load"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "calendar." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set LUMA_TEST_CALENDAR_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // Bootstrap entity data from existing test data.
        $calendar_ref01_data_raw = Vs::items(Helpers::to_map(
            Vs::getpath($setup["data"], "existing.calendar")));
        $calendar_ref01_data = null;
        if (count($calendar_ref01_data_raw) > 0) {
            $calendar_ref01_data = Helpers::to_map($calendar_ref01_data_raw[0][1]);
        }

        // UPDATE
        $calendar_ref01_ent = $client->Calendar(null);
        $calendar_ref01_data_up0_up = [
            "id" => $calendar_ref01_data["id"],
        ];

        $calendar_ref01_markdef_up0_name = "calendar_id";
        $calendar_ref01_markdef_up0_value = "Mark01-calendar_ref01_" . $setup["now"];
        $calendar_ref01_data_up0_up[$calendar_ref01_markdef_up0_name] = $calendar_ref01_markdef_up0_value;

        $calendar_ref01_resdata_up0_result = $calendar_ref01_ent->update($calendar_ref01_data_up0_up, null);
        $calendar_ref01_resdata_up0 = Helpers::to_map($calendar_ref01_resdata_up0_result);
        $this->assertNotNull($calendar_ref01_resdata_up0);
        $this->assertEquals($calendar_ref01_resdata_up0["id"], $calendar_ref01_data_up0_up["id"]);
        $this->assertEquals($calendar_ref01_resdata_up0[$calendar_ref01_markdef_up0_name], $calendar_ref01_markdef_up0_value);

        // LOAD
        $calendar_ref01_match_dt0 = [
            "id" => $calendar_ref01_data["id"],
        ];
        $calendar_ref01_data_dt0_loaded = $calendar_ref01_ent->load($calendar_ref01_match_dt0, null);
        $calendar_ref01_data_dt0_load_result = Helpers::to_map($calendar_ref01_data_dt0_loaded);
        $this->assertNotNull($calendar_ref01_data_dt0_load_result);
        $this->assertEquals($calendar_ref01_data_dt0_load_result["id"], $calendar_ref01_data["id"]);

    }
}

function calendar_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/calendar/CalendarTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = LumaSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["calendar01", "calendar02", "calendar03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("LUMA_TEST_CALENDAR_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "LUMA_TEST_CALENDAR_ENTID" => $idmap,
        "LUMA_TEST_LIVE" => "FALSE",
        "LUMA_TEST_EXPLAIN" => "FALSE",
        "LUMA_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["LUMA_TEST_CALENDAR_ENTID"]);
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
