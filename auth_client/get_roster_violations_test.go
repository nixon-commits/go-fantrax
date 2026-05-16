package auth_client

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/pmurley/go-fantrax/models"
)

func TestParseRosterViolations(t *testing.T) {
	body, err := os.ReadFile(filepath.Join("testdata", "getRosterViolations.json"))
	if err != nil {
		t.Fatal(err)
	}
	var resp models.RosterViolationsResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatal(err)
	}
	data := resp.Responses[0].Data
	if len(data.Violations) == 0 {
		t.Fatal("expected ≥1 violations in fixture")
	}

	// Look for at least one team with a MinMax violation and one with
	// an IllegalRoster violation — both kinds appear in the fixture.
	var sawMinMax, sawIllegal bool
	for _, v := range data.Violations {
		for _, p := range v.Periods {
			if p.IsMinMaxViolated {
				sawMinMax = true
			}
			if p.IsIllegalRoster {
				sawIllegal = true
			}
			if p.Number == 0 {
				t.Error("period number should be non-zero")
			}
			if p.Display == "" {
				t.Error("period display should be non-empty")
			}
		}
	}
	if !sawMinMax {
		t.Error("expected fixture to include a MinMax violation")
	}
	if !sawIllegal {
		t.Error("expected fixture to include an IllegalRoster violation")
	}
}
