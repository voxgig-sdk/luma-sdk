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

func TestGuestTicketEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.GuestTicket(nil)
		if ent == nil {
			t.Fatal("expected non-nil GuestTicketEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := guest_ticketBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"update"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "guest_ticket." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set LUMA_TEST_GUEST_TICKET_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		guestTicketRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.guest_ticket", setup.data)))
		var guestTicketRef01Data map[string]any
		if len(guestTicketRef01DataRaw) > 0 {
			guestTicketRef01Data = core.ToMapAny(guestTicketRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = guestTicketRef01Data

		// UPDATE
		guestTicketRef01Ent := client.GuestTicket(nil)
		guestTicketRef01DataUp0Up := map[string]any{
		}

		guestTicketRef01MarkdefUp0Name := "event_id"
		guestTicketRef01MarkdefUp0Value := fmt.Sprintf("Mark01-guest_ticket_ref01_%d", setup.now)
		guestTicketRef01DataUp0Up[guestTicketRef01MarkdefUp0Name] = guestTicketRef01MarkdefUp0Value

		guestTicketRef01ResdataUp0Result, err := guestTicketRef01Ent.Update(guestTicketRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		guestTicketRef01ResdataUp0 := core.ToMapAny(guestTicketRef01ResdataUp0Result)
		if guestTicketRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if guestTicketRef01ResdataUp0[guestTicketRef01MarkdefUp0Name] != guestTicketRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", guestTicketRef01MarkdefUp0Name, guestTicketRef01ResdataUp0[guestTicketRef01MarkdefUp0Name])
		}

	})
}

func guest_ticketBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "guest_ticket", "GuestTicketTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read guest_ticket test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse guest_ticket test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"guest_ticket01", "guest_ticket02", "guest_ticket03"},
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
	entidEnvRaw := os.Getenv("LUMA_TEST_GUEST_TICKET_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"LUMA_TEST_GUEST_TICKET_ENTID": idmap,
		"LUMA_TEST_LIVE":      "FALSE",
		"LUMA_TEST_EXPLAIN":   "FALSE",
		"LUMA_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["LUMA_TEST_GUEST_TICKET_ENTID"])
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
