package sdktest

import (
	"encoding/json"
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

func TestGuestInviteEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.GuestInvite(nil)
		if ent == nil {
			t.Fatal("expected non-nil GuestInviteEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := guest_inviteBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "guest_invite." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set LUMA_TEST_GUEST_INVITE_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		guestInviteRef01Ent := client.GuestInvite(nil)
		guestInviteRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "guest_invite"}, setup.data), "guest_invite_ref01"))

		guestInviteRef01DataResult, err := guestInviteRef01Ent.Create(guestInviteRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		guestInviteRef01Data = core.ToMapAny(guestInviteRef01DataResult)
		if guestInviteRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}

	})
}

func guest_inviteBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "guest_invite", "GuestInviteTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read guest_invite test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse guest_invite test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"guest_invite01", "guest_invite02", "guest_invite03"},
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
	entidEnvRaw := os.Getenv("LUMA_TEST_GUEST_INVITE_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"LUMA_TEST_GUEST_INVITE_ENTID": idmap,
		"LUMA_TEST_LIVE":      "FALSE",
		"LUMA_TEST_EXPLAIN":   "FALSE",
		"LUMA_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["LUMA_TEST_GUEST_INVITE_ENTID"])
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
