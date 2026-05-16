package auth_client

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/pmurley/go-fantrax/models"
)

func TestParseFantasyTeams(t *testing.T) {
	body, err := os.ReadFile(filepath.Join("testdata", "getFantasyTeams.json"))
	if err != nil {
		t.Fatal(err)
	}
	var resp models.FantasyTeamsResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatal(err)
	}
	teams := resp.Responses[0].Data.FantasyTeams
	if len(teams) < 8 {
		t.Fatalf("expected ≥8 teams, got %d", len(teams))
	}

	// Look up a specific known team.
	var bigz4 *models.FantasyTeam
	for i := range teams {
		if teams[i].Name == "BigZ4" {
			bigz4 = &teams[i]
			break
		}
	}
	if bigz4 == nil {
		t.Fatal("did not find BigZ4 in fixture")
	}
	if bigz4.ID == "" {
		t.Error("team ID should be non-empty")
	}
	if bigz4.LogoURL128 == "" {
		t.Error("team LogoURL128 should be non-empty")
	}
}
