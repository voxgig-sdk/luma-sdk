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

func TestCalendarEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Calendar(nil)
		if ent == nil {
			t.Fatal("expected non-nil CalendarEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := calendarBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"update", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "calendar." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set LUMA_TEST_CALENDAR_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		calendarRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.calendar", setup.data)))
		var calendarRef01Data map[string]any
		if len(calendarRef01DataRaw) > 0 {
			calendarRef01Data = core.ToMapAny(calendarRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = calendarRef01Data

		// UPDATE
		calendarRef01Ent := client.Calendar(nil)
		calendarRef01DataUp0Up := map[string]any{
			"id": calendarRef01Data["id"],
		}

		calendarRef01MarkdefUp0Name := "calendar_id"
		calendarRef01MarkdefUp0Value := fmt.Sprintf("Mark01-calendar_ref01_%d", setup.now)
		calendarRef01DataUp0Up[calendarRef01MarkdefUp0Name] = calendarRef01MarkdefUp0Value

		calendarRef01ResdataUp0Result, err := calendarRef01Ent.Update(calendarRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		calendarRef01ResdataUp0 := core.ToMapAny(calendarRef01ResdataUp0Result)
		if calendarRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if calendarRef01ResdataUp0["id"] != calendarRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if calendarRef01ResdataUp0[calendarRef01MarkdefUp0Name] != calendarRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", calendarRef01MarkdefUp0Name, calendarRef01ResdataUp0[calendarRef01MarkdefUp0Name])
		}

		// LOAD
		calendarRef01MatchDt0 := map[string]any{
			"id": calendarRef01Data["id"],
		}
		calendarRef01DataDt0Loaded, err := calendarRef01Ent.Load(calendarRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		calendarRef01DataDt0LoadResult := core.ToMapAny(calendarRef01DataDt0Loaded)
		if calendarRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if calendarRef01DataDt0LoadResult["id"] != calendarRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

	})
}

func calendarBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "calendar", "CalendarTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read calendar test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse calendar test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"calendar01", "calendar02", "calendar03"},
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
	entidEnvRaw := os.Getenv("LUMA_TEST_CALENDAR_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"LUMA_TEST_CALENDAR_ENTID": idmap,
		"LUMA_TEST_LIVE":      "FALSE",
		"LUMA_TEST_EXPLAIN":   "FALSE",
		"LUMA_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["LUMA_TEST_CALENDAR_ENTID"])
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
