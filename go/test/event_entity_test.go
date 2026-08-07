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

func TestEventEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Event(nil)
		if ent == nil {
			t.Fatal("expected non-nil EventEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := eventBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "update", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "event." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set LUMA_TEST_EVENT_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		eventRef01Ent := client.Event(nil)
		eventRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "event"}, setup.data), "event_ref01"))

		eventRef01DataResult, err := eventRef01Ent.Create(eventRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		eventRef01Data = core.ToMapAny(eventRef01DataResult)
		if eventRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if eventRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// UPDATE
		eventRef01DataUp0Up := map[string]any{
			"id": eventRef01Data["id"],
		}

		eventRef01MarkdefUp0Name := "access"
		eventRef01MarkdefUp0Value := fmt.Sprintf("Mark01-event_ref01_%d", setup.now)
		eventRef01DataUp0Up[eventRef01MarkdefUp0Name] = eventRef01MarkdefUp0Value

		eventRef01ResdataUp0Result, err := eventRef01Ent.Update(eventRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		eventRef01ResdataUp0 := core.ToMapAny(eventRef01ResdataUp0Result)
		if eventRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if eventRef01ResdataUp0["id"] != eventRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if eventRef01ResdataUp0[eventRef01MarkdefUp0Name] != eventRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", eventRef01MarkdefUp0Name, eventRef01ResdataUp0[eventRef01MarkdefUp0Name])
		}

		// LOAD
		eventRef01MatchDt0 := map[string]any{
			"id": eventRef01Data["id"],
		}
		eventRef01DataDt0Loaded, err := eventRef01Ent.Load(eventRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		eventRef01DataDt0LoadResult := core.ToMapAny(eventRef01DataDt0Loaded)
		if eventRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if eventRef01DataDt0LoadResult["id"] != eventRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		eventRef01MatchRm0 := map[string]any{
			"id": eventRef01Data["id"],
		}
		_, err = eventRef01Ent.Remove(eventRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

	})
}

func eventBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "event", "EventTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read event test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse event test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"event01", "event02", "event03"},
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
	entidEnvRaw := os.Getenv("LUMA_TEST_EVENT_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"LUMA_TEST_EVENT_ENTID": idmap,
		"LUMA_TEST_LIVE":      "FALSE",
		"LUMA_TEST_EXPLAIN":   "FALSE",
		"LUMA_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["LUMA_TEST_EVENT_ENTID"])
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
