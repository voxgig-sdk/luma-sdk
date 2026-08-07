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

func TestEventTagEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.EventTag(nil)
		if ent == nil {
			t.Fatal("expected non-nil EventTagEntity")
		}
	})

	// Feature #4: the entity Stream(action, ...) method runs the op pipeline and
	// returns a channel over result items. With the streaming feature active it
	// yields the feature's incremental output; otherwise it falls back to the
	// materialised list so Stream always yields.
	t.Run("stream", func(t *testing.T) {
		seed := map[string]any{
			"entity": map[string]any{
				"event_tag": map[string]any{
					"s1": map[string]any{"id": "s1"},
					"s2": map[string]any{"id": "s2"},
					"s3": map[string]any{"id": "s3"},
				},
			},
		}

		// Fallback: streaming inactive -> yields the materialised list items.
		base := sdk.TestSDK(seed, nil)
		var seen []any
		for item := range base.EventTag(nil).Stream("list", nil, nil) {
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
			for item := range streamSdk.EventTag(nil).Stream("list", nil, nil) {
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
		setup := event_tagBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "update", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "event_tag." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set LUMA_TEST_EVENT_TAG_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		eventTagRef01Ent := client.EventTag(nil)
		eventTagRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "event_tag"}, setup.data), "event_tag_ref01"))

		eventTagRef01DataResult, err := eventTagRef01Ent.Create(eventTagRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		eventTagRef01Data = core.ToMapAny(eventTagRef01DataResult)
		if eventTagRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if eventTagRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		eventTagRef01Match := map[string]any{}

		eventTagRef01ListResult, err := eventTagRef01Ent.List(eventTagRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		eventTagRef01List, eventTagRef01ListOk := eventTagRef01ListResult.([]any)
		if !eventTagRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", eventTagRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(eventTagRef01List), map[string]any{"id": eventTagRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// UPDATE
		eventTagRef01DataUp0Up := map[string]any{
			"id": eventTagRef01Data["id"],
		}

		eventTagRef01MarkdefUp0Name := "name"
		eventTagRef01MarkdefUp0Value := fmt.Sprintf("Mark01-event_tag_ref01_%d", setup.now)
		eventTagRef01DataUp0Up[eventTagRef01MarkdefUp0Name] = eventTagRef01MarkdefUp0Value

		eventTagRef01ResdataUp0Result, err := eventTagRef01Ent.Update(eventTagRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		eventTagRef01ResdataUp0 := core.ToMapAny(eventTagRef01ResdataUp0Result)
		if eventTagRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if eventTagRef01ResdataUp0["id"] != eventTagRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if eventTagRef01ResdataUp0[eventTagRef01MarkdefUp0Name] != eventTagRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", eventTagRef01MarkdefUp0Name, eventTagRef01ResdataUp0[eventTagRef01MarkdefUp0Name])
		}

		// REMOVE
		eventTagRef01MatchRm0 := map[string]any{
			"id": eventTagRef01Data["id"],
		}
		_, err = eventTagRef01Ent.Remove(eventTagRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		eventTagRef01MatchRt0 := map[string]any{}

		eventTagRef01ListRt0Result, err := eventTagRef01Ent.List(eventTagRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		eventTagRef01ListRt0, eventTagRef01ListRt0Ok := eventTagRef01ListRt0Result.([]any)
		if !eventTagRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", eventTagRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(eventTagRef01ListRt0), map[string]any{"id": eventTagRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func event_tagBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "event_tag", "EventTagTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read event_tag test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse event_tag test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"event_tag01", "event_tag02", "event_tag03"},
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
	entidEnvRaw := os.Getenv("LUMA_TEST_EVENT_TAG_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"LUMA_TEST_EVENT_TAG_ENTID": idmap,
		"LUMA_TEST_LIVE":      "FALSE",
		"LUMA_TEST_EXPLAIN":   "FALSE",
		"LUMA_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["LUMA_TEST_EVENT_TAG_ENTID"])
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
