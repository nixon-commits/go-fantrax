package auth_client

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/pmurley/go-fantrax/models"
)

func TestParseMatchups(t *testing.T) {
	body, err := os.ReadFile(filepath.Join("testdata", "getMatchups.json"))
	if err != nil {
		t.Fatal(err)
	}
	var resp models.MatchupsListResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatal(err)
	}
	data := resp.Responses[0].Data
	if data.RegularSeasonEndPeriod <= 0 {
		t.Errorf("expected RegularSeasonEndPeriod > 0, got %d", data.RegularSeasonEndPeriod)
	}
	if len(data.FantasyTeams) < 8 {
		t.Errorf("expected ≥8 fantasyTeams, got %d", len(data.FantasyTeams))
	}
	if len(data.Periods) < 1 {
		t.Fatal("expected ≥1 period")
	}
	p1 := data.Periods[0]
	if p1.Number != 1 {
		t.Errorf("expected first period Number=1, got %d", p1.Number)
	}
	if len(p1.Matchups) < 1 {
		t.Fatal("expected ≥1 matchup in first period")
	}
	m := p1.Matchups[0]
	if m.HomeTeam.ID == "" || m.AwayTeam.ID == "" {
		t.Errorf("expected matchup teams to have IDs, got %+v / %+v", m.HomeTeam, m.AwayTeam)
	}
}
