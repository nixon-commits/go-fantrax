package auth_client

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/pmurley/go-fantrax/models"
)

func TestParseLeagueHistory(t *testing.T) {
	body, err := os.ReadFile(filepath.Join("testdata", "getLeagueHistory.json"))
	if err != nil {
		t.Fatal(err)
	}
	var resp models.LeagueHistoryResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatal(err)
	}
	data := resp.Responses[0].Data
	if data.LeagueHistoryID == "" {
		t.Error("expected non-empty LeagueHistoryID")
	}
	if data.FantasyPeriodName != "Week" {
		t.Errorf("expected FantasyPeriodName=Week, got %q", data.FantasyPeriodName)
	}
	if !data.AnyPlayoffsUsed {
		t.Error("fixture has AnyPlayoffsUsed=true")
	}
	if len(data.DisplayedLists.Tabs) < 1 {
		t.Error("expected at least one tab")
	}
}
