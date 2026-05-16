package auth_client

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/pmurley/go-fantrax/models"
)

func TestParseFantasyLeagueInfo(t *testing.T) {
	body, err := os.ReadFile(filepath.Join("testdata", "getFantasyLeagueInfo.json"))
	if err != nil {
		t.Fatal(err)
	}
	var resp models.FantasyLeagueInfoResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatal(err)
	}
	data := resp.Responses[0].Data
	if data.FantasySettings.LeagueName == "" {
		t.Error("expected LeagueName to be set")
	}
	if data.FantasySettings.Sport == "" {
		t.Error("expected Sport to be set")
	}
	if !data.FantasySettings.HeadToHead {
		t.Error("fixture league is head-to-head")
	}
	if len(data.FantasySettings.MyTeamIDs) == 0 {
		t.Error("expected MyTeamIDs to be populated")
	}
	if data.FantasySettings.Season.ID == "" {
		t.Error("expected Season.ID to be set")
	}
	if data.FantasySettings.Season.StartDate == 0 || data.FantasySettings.Season.EndDate == 0 {
		t.Error("expected season start/end dates set")
	}

	// PositionMap: every entry must have an ID matching its key and a short name.
	if len(data.PositionMap) < 5 {
		t.Errorf("expected ≥5 positions in map, got %d", len(data.PositionMap))
	}
	for key, pos := range data.PositionMap {
		if pos.ID != key {
			t.Errorf("position map key %q does not match ID %q", key, pos.ID)
		}
		if pos.ShortName == "" {
			t.Errorf("position %q missing ShortName", key)
		}
	}

	// Catcher slot ID is "001" by convention.
	if c, ok := data.PositionMap["001"]; ok {
		if c.ShortName != "C" {
			t.Errorf("expected position 001 ShortName=C, got %q", c.ShortName)
		}
	} else {
		t.Error("expected position 001 (Catcher) in map")
	}
}
