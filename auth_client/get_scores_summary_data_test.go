package auth_client

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/pmurley/go-fantrax/models"
)

func TestParseScoresSummaryData(t *testing.T) {
	body, err := os.ReadFile(filepath.Join("testdata", "getScoresSummaryData.json"))
	if err != nil {
		t.Fatal(err)
	}
	var resp models.ScoresSummaryResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatal(err)
	}
	scores := resp.Responses[0].Data.Scores
	if len(scores) < 1 {
		t.Fatal("expected ≥1 score in fixture")
	}

	// Ensure both pregame (StartTime set, no StatusDisplay) and live
	// (StatusDisplay set) entries parse cleanly.
	var sawLive, sawPregame bool
	for _, s := range scores {
		if s.EventID == "" || s.HomeTeamID == "" || s.AwayTeamID == "" {
			t.Errorf("malformed score entry: %+v", s)
		}
		if s.StatusDisplay != "" {
			sawLive = true
		}
		if s.StatusCode == "PREGAME" && s.StartTime > 0 {
			sawPregame = true
		}
	}
	if !sawLive {
		t.Error("expected at least one live score (StatusDisplay set)")
	}
	if !sawPregame {
		t.Error("expected at least one pregame score (StatusCode=PREGAME, StartTime set)")
	}
}
