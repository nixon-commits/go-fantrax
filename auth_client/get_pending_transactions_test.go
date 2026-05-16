package auth_client

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/pmurley/go-fantrax/models"
)

func TestParsePendingTransactions(t *testing.T) {
	body, err := os.ReadFile(filepath.Join("testdata", "getPendingTransactions.json"))
	if err != nil {
		t.Fatal(err)
	}
	var resp models.PendingTransactionsResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatal(err)
	}
	data := resp.Responses[0].Data
	if data.SelectedTxType == "" {
		t.Error("expected SelectedTxType to be set")
	}
	if len(data.MyTeamIDs) == 0 {
		t.Error("expected MyTeamIDs to be populated")
	}
	if len(data.TradeInfoList) < 1 {
		t.Fatal("expected at least one pending trade in fixture")
	}
	t0 := data.TradeInfoList[0]
	if t0.TxSetID == "" {
		t.Error("expected TxSetID to be set")
	}
	if t0.CreatorTeamID == "" {
		t.Error("expected CreatorTeamID to be set")
	}
	if len(t0.Moves) < 1 {
		t.Error("expected at least one move")
	}
	// Fixture has illegalRosterMsgs — verify they parse.
	if len(t0.IllegalRosterMessages) > 0 {
		msg := t0.IllegalRosterMessages[0]
		if msg.TeamID == "" {
			t.Error("expected IllegalRosterMessage.TeamID set")
		}
		if len(msg.MessagesPerPeriod) > 0 && msg.MessagesPerPeriod[0].Period == "" {
			t.Error("expected per-period message Period set")
		}
	}
}
