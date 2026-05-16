package auth_client

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/pmurley/go-fantrax/models"
)

func TestParseLiveScoringStats(t *testing.T) {
	body, err := os.ReadFile(filepath.Join("testdata", "getLiveScoringStats.json"))
	if err != nil {
		t.Fatal(err)
	}
	var resp models.LiveScoringStatsResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatal(err)
	}
	data := resp.Responses[0].Data
	if data.DisplayPeriod == "" {
		t.Error("expected DisplayPeriod to be set")
	}
	if len(data.FantasyTeams) < 8 {
		t.Errorf("expected ≥8 fantasyTeams, got %d", len(data.FantasyTeams))
	}
	if len(data.OwnTeamIDs) == 0 {
		t.Error("expected OwnTeamIDs populated")
	}
	if len(data.StatsPerTeam.AllTeamsStats) < 1 {
		t.Fatal("expected StatsPerTeam.AllTeamsStats to have at least one team")
	}

	// Verify per-team buckets carry the WP-relevant fields.
	for teamID, buckets := range data.StatsPerTeam.AllTeamsStats {
		if buckets.Active == nil {
			t.Errorf("team %s: expected ACTIVE bucket", teamID)
			continue
		}
		if buckets.Active.TotalFpts2 == nil && buckets.Active.TotalFpts == nil {
			t.Errorf("team %s: expected at least one of TotalFpts/TotalFpts2", teamID)
		}
		if buckets.Active.TotalFpts != nil && *buckets.Active.TotalFpts < 0 {
			t.Errorf("team %s: TotalFpts should not be negative", teamID)
		}
		// playerGameInfo / projected totals are sometimes empty for
		// teams with no players to score on a given day — don't assert
		// presence, just confirm parse didn't blow up.
		_ = buckets.Active.PlayerGameInfo
		_ = buckets.Active.CalculatedProjectedTotalsMap
	}
}
