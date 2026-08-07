package sdktest

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/luma-sdk/go"
	"github.com/voxgig-sdk/luma-sdk/go/core"

	vs "github.com/voxgig-sdk/luma-sdk/go/utility/struct"
)

func TestCalendarCouponEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.CalendarCoupon(nil)
		if ent == nil {
			t.Fatal("expected non-nil CalendarCouponEntity")
		}
	})

	// Feature #4: the entity Stream(action, ...) method runs the op pipeline and
	// returns a channel over result items. With the streaming feature active it
	// yields the feature's incremental output; otherwise it falls back to the
	// materialised list so Stream always yields.
	t.Run("stream", func(t *testing.T) {
		seed := map[string]any{
			"entity": map[string]any{
				"calendar_coupon": map[string]any{
					"s1": map[string]any{"id": "s1"},
					"s2": map[string]any{"id": "s2"},
					"s3": map[string]any{"id": "s3"},
				},
			},
		}

		// Fallback: streaming inactive -> yields the materialised list items.
		base := sdk.TestSDK(seed, nil)
		var seen []any
		for item := range base.CalendarCoupon(nil).Stream("list", nil, nil) {
			seen = append(seen, item)
		}
		if len(seen) != 3 {
			t.Fatalf("expected 3 streamed items, got %d", len(seen))
		}

		// Inbound: streaming active -> yields each item from the feature iterator.
		hasStreaming := false
		if fm, ok := core.MakeConfig()["feature"].(map[string]any); ok {
			_, hasStreaming = fm["streaming"]
		}
		if hasStreaming {
			streamSdk := sdk.TestSDK(seed, map[string]any{
				"feature": map[string]any{"streaming": map[string]any{"active": true}},
			})
			var got []any
			for item := range streamSdk.CalendarCoupon(nil).Stream("list", nil, nil) {
				if sub, ok := item.([]any); ok {
					got = append(got, sub...)
				} else {
					got = append(got, item)
				}
			}
			if len(got) != 3 {
				t.Fatalf("expected 3 items via streaming feature, got %d", len(got))
			}
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := calendar_couponBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "update"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "calendar_coupon." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		// The basic flow consumes synthetic IDs from the fixture. In live mode
		// without an *_ENTID env override, those IDs hit the live API and 4xx.
		if setup.syntheticOnly {
			t.Skip("live entity test uses synthetic IDs from fixture — set LUMA_TEST_CALENDAR_COUPON_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		calendarCouponRef01Ent := client.CalendarCoupon(nil)
		calendarCouponRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "calendar_coupon"}, setup.data), "calendar_coupon_ref01"))

		calendarCouponRef01DataResult, err := calendarCouponRef01Ent.Create(calendarCouponRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		calendarCouponRef01Data = core.ToMapAny(calendarCouponRef01DataResult)
		if calendarCouponRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if calendarCouponRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		calendarCouponRef01Match := map[string]any{}

		calendarCouponRef01ListResult, err := calendarCouponRef01Ent.List(calendarCouponRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		calendarCouponRef01List, calendarCouponRef01ListOk := calendarCouponRef01ListResult.([]any)
		if !calendarCouponRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", calendarCouponRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(calendarCouponRef01List), map[string]any{"id": calendarCouponRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// UPDATE
		calendarCouponRef01DataUp0Up := map[string]any{
			"id": calendarCouponRef01Data["id"],
		}

		calendarCouponRef01MarkdefUp0Name := "code"
		calendarCouponRef01MarkdefUp0Value := fmt.Sprintf("Mark01-calendar_coupon_ref01_%d", setup.now)
		calendarCouponRef01DataUp0Up[calendarCouponRef01MarkdefUp0Name] = calendarCouponRef01MarkdefUp0Value

		calendarCouponRef01ResdataUp0Result, err := calendarCouponRef01Ent.Update(calendarCouponRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		calendarCouponRef01ResdataUp0 := core.ToMapAny(calendarCouponRef01ResdataUp0Result)
		if calendarCouponRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if calendarCouponRef01ResdataUp0["id"] != calendarCouponRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if calendarCouponRef01ResdataUp0[calendarCouponRef01MarkdefUp0Name] != calendarCouponRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", calendarCouponRef01MarkdefUp0Name, calendarCouponRef01ResdataUp0[calendarCouponRef01MarkdefUp0Name])
		}

	})
}

func calendar_couponBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "calendar_coupon", "CalendarCouponTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read calendar_coupon test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse calendar_coupon test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"calendar_coupon01", "calendar_coupon02", "calendar_coupon03"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Detect ENTID env override before envOverride consumes it. When live
	// mode is on without a real override, the basic test runs against synthetic
	// IDs from the fixture and 4xx's. Surface this so the test can skip.
	entidEnvRaw := os.Getenv("LUMA_TEST_CALENDAR_COUPON_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"LUMA_TEST_CALENDAR_COUPON_ENTID": idmap,
		"LUMA_TEST_LIVE":      "FALSE",
		"LUMA_TEST_EXPLAIN":   "FALSE",
		"LUMA_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["LUMA_TEST_CALENDAR_COUPON_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["LUMA_TEST_LIVE"] == "TRUE" {
		mergedOpts := vs.Merge([]any{
			map[string]any{
				"apikey": env["LUMA_APIKEY"],
			},
			extra,
		})
		client = sdk.NewLumaSDK(core.ToMapAny(mergedOpts))
	}

	live := env["LUMA_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["LUMA_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}
