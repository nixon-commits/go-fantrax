package auth_client

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/pmurley/go-fantrax/models"
)

func TestParseLeagueHomeStandings(t *testing.T) {
	body, err := os.ReadFile(filepath.Join("testdata", "getLeagueHomeStandings.json"))
	if err != nil {
		t.Fatal(err)
	}
	var resp models.LeagueHomeStandingsResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatal(err)
	}

	combined := resp.Responses[0].Data.CombinedStandings()
	if len(combined) < 8 {
		t.Fatalf("expected ≥8 combined rows, got %d", len(combined))
	}

	first := combined[0]
	if first.Rank != 1 {
		t.Errorf("expected first row Rank=1, got %d", first.Rank)
	}
	if first.Team == "" {
		t.Error("expected first row Team non-empty")
	}
	if first.TeamID == "" {
		t.Error("expected first row TeamID non-empty")
	}
	if first.WinPercentage == "" {
		t.Error("expected first row WinPercentage non-empty")
	}
}
