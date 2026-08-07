<?php
declare(strict_types=1);

// Member entity test

require_once __DIR__ . '/../luma_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class MemberEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = LumaSDK::test(null, null);
        $ent = $testsdk->Member(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = member_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["create", "update"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "member." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set LUMA_TEST_MEMBER_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // CREATE
        $member_ref01_ent = $client->Member(null);
        $member_ref01_data = Helpers::to_map(Vs::getprop(
            Vs::getpath($setup["data"], "new.member"), "member_ref01"));

        $member_ref01_data_result = $member_ref01_ent->create($member_ref01_data, null);
        $member_ref01_data = Helpers::to_map($member_ref01_data_result);
        $this->assertNotNull($member_ref01_data);

        // UPDATE
        $member_ref01_data_up0_up = [
        ];

        $member_ref01_markdef_up0_name = "email";
        $member_ref01_markdef_up0_value = "Mark01-member_ref01_" . $setup["now"];
        $member_ref01_data_up0_up[$member_ref01_markdef_up0_name] = $member_ref01_markdef_up0_value;

        $member_ref01_resdata_up0_result = $member_ref01_ent->update($member_ref01_data_up0_up, null);
        $member_ref01_resdata_up0 = Helpers::to_map($member_ref01_resdata_up0_result);
        $this->assertNotNull($member_ref01_resdata_up0);
        $this->assertEquals($member_ref01_resdata_up0[$member_ref01_markdef_up0_name], $member_ref01_markdef_up0_value);

    }
}

function member_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/member/MemberTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = LumaSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["member01", "member02", "member03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("LUMA_TEST_MEMBER_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "LUMA_TEST_MEMBER_ENTID" => $idmap,
        "LUMA_TEST_LIVE" => "FALSE",
        "LUMA_TEST_EXPLAIN" => "FALSE",
        "LUMA_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["LUMA_TEST_MEMBER_ENTID"]);
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
