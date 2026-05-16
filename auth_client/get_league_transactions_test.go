package auth_client

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/pmurley/go-fantrax/models"
)

func TestParseLeagueTransactions(t *testing.T) {
	body, err := os.ReadFile(filepath.Join("testdata", "getTransactions.json"))
	if err != nil {
		t.Fatal(err)
	}
	var resp models.LeagueTransactionsResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatal(err)
	}
	data := resp.Responses[0].Data
	if len(data.Transactions) == 0 {
		t.Error("expected Transactions raw JSON to be non-empty")
	}
	if len(data.DisplayedLists) == 0 {
		t.Error("expected DisplayedLists raw JSON to be non-empty")
	}
}
