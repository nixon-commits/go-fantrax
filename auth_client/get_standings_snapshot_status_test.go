package auth_client

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/pmurley/go-fantrax/models"
)

func TestParseStandingsSnapshotStatus(t *testing.T) {
	body, err := os.ReadFile(filepath.Join("testdata", "getStandingsSnapshotStatus.json"))
	if err != nil {
		t.Fatal(err)
	}
	var resp models.StandingsSnapshotStatusResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatal(err)
	}
	if len(resp.Responses) != 1 {
		t.Fatalf("expected 1 response, got %d", len(resp.Responses))
	}
	if !resp.Responses[0].Data.Ready {
		t.Error("expected Ready=true in fixture")
	}
}
