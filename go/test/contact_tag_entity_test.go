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

func TestContactTagEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.ContactTag(nil)
		if ent == nil {
			t.Fatal("expected non-nil ContactTagEntity")
		}
	})

	// Feature #4: the entity Stream(action, ...) method runs the op pipeline and
	// returns a channel over result items. With the streaming feature active it
	// yields the feature's incremental output; otherwise it falls back to the
	// materialised list so Stream always yields.
	t.Run("stream", func(t *testing.T) {
		seed := map[string]any{
			"entity": map[string]any{
				"contact_tag": map[string]any{
					"s1": map[string]any{"id": "s1"},
					"s2": map[string]any{"id": "s2"},
					"s3": map[string]any{"id": "s3"},
				},
			},
		}

		// Fallback: streaming inactive -> yields the materialised list items.
		base := sdk.TestSDK(seed, nil)
		var seen []any
		for item := range base.ContactTag(nil).Stream("list", nil, nil) {
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
			for item := range streamSdk.ContactTag(nil).Stream("list", nil, nil) {
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
		setup := contact_tagBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "update", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "contact_tag." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set LUMA_TEST_CONTACT_TAG_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		contactTagRef01Ent := client.ContactTag(nil)
		contactTagRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "contact_tag"}, setup.data), "contact_tag_ref01"))

		contactTagRef01DataResult, err := contactTagRef01Ent.Create(contactTagRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		contactTagRef01Data = core.ToMapAny(contactTagRef01DataResult)
		if contactTagRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if contactTagRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		contactTagRef01Match := map[string]any{}

		contactTagRef01ListResult, err := contactTagRef01Ent.List(contactTagRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		contactTagRef01List, contactTagRef01ListOk := contactTagRef01ListResult.([]any)
		if !contactTagRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", contactTagRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(contactTagRef01List), map[string]any{"id": contactTagRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// UPDATE
		contactTagRef01DataUp0Up := map[string]any{
			"id": contactTagRef01Data["id"],
		}

		contactTagRef01MarkdefUp0Name := "name"
		contactTagRef01MarkdefUp0Value := fmt.Sprintf("Mark01-contact_tag_ref01_%d", setup.now)
		contactTagRef01DataUp0Up[contactTagRef01MarkdefUp0Name] = contactTagRef01MarkdefUp0Value

		contactTagRef01ResdataUp0Result, err := contactTagRef01Ent.Update(contactTagRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		contactTagRef01ResdataUp0 := core.ToMapAny(contactTagRef01ResdataUp0Result)
		if contactTagRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if contactTagRef01ResdataUp0["id"] != contactTagRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if contactTagRef01ResdataUp0[contactTagRef01MarkdefUp0Name] != contactTagRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", contactTagRef01MarkdefUp0Name, contactTagRef01ResdataUp0[contactTagRef01MarkdefUp0Name])
		}

		// REMOVE
		contactTagRef01MatchRm0 := map[string]any{
			"id": contactTagRef01Data["id"],
		}
		_, err = contactTagRef01Ent.Remove(contactTagRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		contactTagRef01MatchRt0 := map[string]any{}

		contactTagRef01ListRt0Result, err := contactTagRef01Ent.List(contactTagRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		contactTagRef01ListRt0, contactTagRef01ListRt0Ok := contactTagRef01ListRt0Result.([]any)
		if !contactTagRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", contactTagRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(contactTagRef01ListRt0), map[string]any{"id": contactTagRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func contact_tagBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "contact_tag", "ContactTagTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read contact_tag test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse contact_tag test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"contact_tag01", "contact_tag02", "contact_tag03"},
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
	entidEnvRaw := os.Getenv("LUMA_TEST_CONTACT_TAG_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"LUMA_TEST_CONTACT_TAG_ENTID": idmap,
		"LUMA_TEST_LIVE":      "FALSE",
		"LUMA_TEST_EXPLAIN":   "FALSE",
		"LUMA_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["LUMA_TEST_CONTACT_TAG_ENTID"])
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
