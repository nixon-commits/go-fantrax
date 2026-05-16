package models

import "encoding/json"

// PendingTransactionMove is one side of a pending trade or claim — a player
// changing teams and/or status.
type PendingTransactionMove struct {
	EligibleStatusIDs   []string        `json:"eligibleStatusesIds,omitempty"`
	EligiblePositionIDs []string        `json:"eligiblePositionIds,omitempty"`
	ToStatus            json.RawMessage `json:"toStatus,omitempty"`
	FromStatus          json.RawMessage `json:"fromStatus,omitempty"`
	Scorer              json.RawMessage `json:"scorer,omitempty"`
	// Many additional optional fields (priority, dropDate, salary, etc.) are
	// preserved on Raw for callers that need them without re-fetching.
}

// IllegalRosterMessage is the per-period roster-validity warning attached to
// a pending trade for one of the participating teams.
type IllegalRosterMessage struct {
	TeamID            string `json:"teamId"`
	MessagesPerPeriod []struct {
		Period   string   `json:"period"`   // e.g. "54 (May 17, 2026)"
		Messages []string `json:"messages"` // HTML-tagged warnings
	} `json:"messagesPerPeriod"`
}

// PendingTradeInfo is one pending trade in the league.
type PendingTradeInfo struct {
	TxSetID               string                   `json:"txSetId"`
	LeagueID              string                   `json:"leagueId"`
	CreatorTeamID         string                   `json:"creatorTeamId"`
	Pending               bool                     `json:"pending"`
	AllowEditSalaryTraded bool                     `json:"allowEditSalaryTraded"`
	AllowEditBudgetTraded bool                     `json:"allowEditBudgetTraded"`
	IllegalRosterMessages []IllegalRosterMessage   `json:"illegalRosterMsgs,omitempty"`
	Moves                 []PendingTransactionMove `json:"moves"`
}

// PendingTransactions is the response data for getPendingTransactions.
type PendingTransactions struct {
	MyTeamIDs      []string           `json:"myTeamIds"`
	TeamIDs        []string           `json:"teamIds"`
	SelectedTeamID string             `json:"selectedTeamId"`
	SelectedTxType string             `json:"selectedTxType"` // e.g. "TRADE"
	TimeZoneLabel  string             `json:"timeZoneLabel"`
	TeamID         string             `json:"teamId"`
	TradeInfoList  []PendingTradeInfo `json:"tradeInfoList"`
	MiscData       json.RawMessage    `json:"miscData,omitempty"`
}

// PendingTransactionsResponse is the full fxpa envelope.
type PendingTransactionsResponse struct {
	Responses []struct {
		Data PendingTransactions `json:"data"`
	} `json:"responses"`
}
