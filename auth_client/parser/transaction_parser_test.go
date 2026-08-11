package parser

import (
	"testing"

	"github.com/pmurley/go-fantrax/models"
)

// Fixtures below are verbatim draftPickDisplayParts values captured from a
// live getTransactionDetailsHistory TRADE-view response (rosterbot-uc3).
func TestParseTransactionRow_DraftPickStillProjected(t *testing.T) {
	row := models.TransactionRow{
		TxSetID: "kxabor9vmsl9fitp",
		DraftPickDisplayParts: &models.DraftPickDisplayParts{
			Year:      "<b>2027</b> Draft Pick",
			RoundInfo: "Round <b>3</b> (Houston Swang and Bang)",
		},
		Cells: []models.TableCell{
			{Key: "from", Content: "Houston Swang and Bang", TeamID: "a"},
			{Key: "to", Content: "Finding Nimmo", TeamID: "b"},
		},
	}
	tx, err := parseTransactionRow(row, "")
	if err != nil {
		t.Fatalf("parseTransactionRow: %v", err)
	}
	if !tx.IsDraftPick {
		t.Fatal("IsDraftPick = false, want true")
	}
	if tx.DraftPickYear != 2027 {
		t.Errorf("DraftPickYear = %d, want 2027", tx.DraftPickYear)
	}
	if tx.DraftPickRound != 3 {
		t.Errorf("DraftPickRound = %d, want 3", tx.DraftPickRound)
	}
	if tx.DraftPickNumber != 0 {
		t.Errorf("DraftPickNumber = %d, want 0 (slot not yet resolved)", tx.DraftPickNumber)
	}
	if tx.DraftPickOriginalTeam != "Houston Swang and Bang" {
		t.Errorf("DraftPickOriginalTeam = %q, want %q", tx.DraftPickOriginalTeam, "Houston Swang and Bang")
	}
	// Still a trade: from/to must parse exactly as they do for a player row.
	if tx.Type != "TRADE" || tx.FromTeamName != "Houston Swang and Bang" || tx.ToTeamName != "Finding Nimmo" {
		t.Errorf("trade fields not preserved: %+v", tx)
	}
	if tx.PlayerName != "" {
		t.Errorf("PlayerName = %q, want empty -- callers key off this for the pre-existing pick signal", tx.PlayerName)
	}
}

func TestParseTransactionRow_DraftPickResolvedSlot(t *testing.T) {
	row := models.TransactionRow{
		TxSetID: "joljvghdmkrju32n",
		DraftPickDisplayParts: &models.DraftPickDisplayParts{
			Year:      "<b>2026</b> Draft Pick",
			RoundInfo: "Round <b>1</b> Pick <b>6</b>",
		},
		Cells: []models.TableCell{
			{Key: "from", Content: "Houston Swang and Bang", TeamID: "a"},
			{Key: "to", Content: "Voradakis", TeamID: "b"},
		},
	}
	tx, err := parseTransactionRow(row, "")
	if err != nil {
		t.Fatalf("parseTransactionRow: %v", err)
	}
	if tx.DraftPickYear != 2026 || tx.DraftPickRound != 1 || tx.DraftPickNumber != 6 {
		t.Errorf("got year=%d round=%d number=%d, want 2026/1/6",
			tx.DraftPickYear, tx.DraftPickRound, tx.DraftPickNumber)
	}
	if tx.DraftPickOriginalTeam != "" {
		t.Errorf("DraftPickOriginalTeam = %q, want empty once the slot is resolved", tx.DraftPickOriginalTeam)
	}
}

// A player row must not be mistaken for a pick: IsDraftPick only follows the
// presence of DraftPickDisplayParts, never an empty PlayerName by itself.
func TestParseTransactionRow_PlayerRowIsNotADraftPick(t *testing.T) {
	row := models.TransactionRow{
		TxSetID: "x",
		Scorer:  models.TransactionPlayer{Name: "Mike Trout"},
		Cells: []models.TableCell{
			{Key: "from", Content: "A", TeamID: "a"},
			{Key: "to", Content: "B", TeamID: "b"},
		},
	}
	tx, err := parseTransactionRow(row, "")
	if err != nil {
		t.Fatalf("parseTransactionRow: %v", err)
	}
	if tx.IsDraftPick {
		t.Error("IsDraftPick = true for a player row, want false")
	}
	if tx.PlayerName != "Mike Trout" {
		t.Errorf("PlayerName = %q, want %q", tx.PlayerName, "Mike Trout")
	}
}
