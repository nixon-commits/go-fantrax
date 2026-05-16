package models

import "encoding/json"

// LeagueTransactionsData is the response data for the getTransactions fxpa
// method (distinct from the older getTransactionDetailsHistory). It is a
// lighter, league-wide transaction list with display metadata.
type LeagueTransactionsData struct {
	Transactions        json.RawMessage `json:"transactions"`
	DisplayedSelections json.RawMessage `json:"displayedSelections,omitempty"`
	DisplayedLists      json.RawMessage `json:"displayedLists,omitempty"`
	MiscData            json.RawMessage `json:"miscData,omitempty"`
}

// LeagueTransactionsResponse is the full fxpa envelope.
type LeagueTransactionsResponse struct {
	Responses []struct {
		Data LeagueTransactionsData `json:"data"`
	} `json:"responses"`
}
