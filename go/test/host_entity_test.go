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

func TestHostEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Host(nil)
		if ent == nil {
			t.Fatal("expected non-nil HostEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := hostBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "update", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "host." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set LUMA_TEST_HOST_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		hostRef01Ent := client.Host(nil)
		hostRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "host"}, setup.data), "host_ref01"))

		hostRef01DataResult, err := hostRef01Ent.Create(hostRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		hostRef01Data = core.ToMapAny(hostRef01DataResult)
		if hostRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}

		// UPDATE
		hostRef01DataUp0Up := map[string]any{
		}

		hostRef01MarkdefUp0Name := "email"
		hostRef01MarkdefUp0Value := fmt.Sprintf("Mark01-host_ref01_%d", setup.now)
		hostRef01DataUp0Up[hostRef01MarkdefUp0Name] = hostRef01MarkdefUp0Value

		hostRef01ResdataUp0Result, err := hostRef01Ent.Update(hostRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		hostRef01ResdataUp0 := core.ToMapAny(hostRef01ResdataUp0Result)
		if hostRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if hostRef01ResdataUp0[hostRef01MarkdefUp0Name] != hostRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", hostRef01MarkdefUp0Name, hostRef01ResdataUp0[hostRef01MarkdefUp0Name])
		}


	})
}

func hostBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "host", "HostTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read host test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse host test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"host01", "host02", "host03"},
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
	entidEnvRaw := os.Getenv("LUMA_TEST_HOST_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"LUMA_TEST_HOST_ENTID": idmap,
		"LUMA_TEST_LIVE":      "FALSE",
		"LUMA_TEST_EXPLAIN":   "FALSE",
		"LUMA_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["LUMA_TEST_HOST_ENTID"])
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
