package auth_client

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/pmurley/go-fantrax/models"
)

func TestParseScoringCategoryGlossary(t *testing.T) {
	body, err := os.ReadFile(filepath.Join("testdata", "getScoringCategoryGlossary.json"))
	if err != nil {
		t.Fatal(err)
	}
	var resp models.ScoringCategoryGlossaryResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatal(err)
	}
	data := resp.Responses[0].Data
	if data.SportName == "" {
		t.Error("expected SportName to be set")
	}
	if len(data.IndivCategories) == 0 {
		t.Error("expected IndivCategories raw JSON to be non-empty")
	}
	if len(data.TeamCategories) == 0 {
		t.Error("expected TeamCategories raw JSON to be non-empty")
	}
}
