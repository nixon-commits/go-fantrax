package models

import "encoding/json"

// LiveScoringStats is a deliberately narrow view of the getLiveScoringStats
// response. The full payload is huge (>800KB on a typical league) and most
// of it is per-player roster/score detail that the consumer can re-derive
// from getTeamRosterInfo + the matchup metadata. The fields exposed here
// are the ones needed to drive client-side win-probability and matchup
// summaries — everything else is preserved via FullRaw for forward-compat.
type LiveScoringStats struct {
	// DisplayPeriod is a date-range label like "(May 11, 2026 - May 17, 2026)".
	DisplayPeriod     string        `json:"displayPeriod"`
	DisplayDate       string        `json:"displayDate"`
	AllEventsFinished bool          `json:"allEventsFinished"`
	OwnTeamIDs        []string      `json:"ownTeamIds"`
	FantasyTeams      []FantasyTeam `json:"fantasyTeams"`
	StatsPerTeam      StatsPerTeam  `json:"statsPerTeam"`
	// Matchups is the list of matchup keys in the current period, each
	// formatted "<homeTeamID>_<awayTeamID>". Split on "_" to recover
	// the two team IDs.
	Matchups []string `json:"matchups,omitempty"`

	// Anything not typed above stays on FullRaw for advanced callers.
	FullRaw json.RawMessage `json:"-"`
}

// StatsPerTeam wraps the per-team active/reserve/IR stat blob.
type StatsPerTeam struct {
	AllTeamsStats map[string]TeamStatsBuckets `json:"allTeamsStats"`
}

// TeamStatsBuckets keys by roster bucket (ACTIVE, RESERVE, IR, MINORS) — only
// ACTIVE is populated for the live-scoring view.
type TeamStatsBuckets struct {
	Active *TeamActiveStats `json:"ACTIVE,omitempty"`
}

// TeamActiveStats is the per-team running stat snapshot.
// `TotalFpts` is season-to-date, `TotalFpts2` is the matchup-period total
// (preferred for live-scoring display). Both arrive as scalar floats at
// this level; the {total,digit,decimal} object form (FptsTotal) only
// appears in downstream UI-derived structures.
type TeamActiveStats struct {
	TotalFpts        *float64        `json:"totalFpts,omitempty"`
	TotalFpts2       *float64        `json:"totalFpts2,omitempty"`
	PointsAdjustment json.RawMessage `json:"pointsAdjustment,omitempty"`
	// PlayerGameInfo: last element is the team's remaining game count
	// for the current period. Fantrax sometimes serialises these counts
	// as floats (54.0), so the slice is []float64 — cast as needed.
	PlayerGameInfo               []float64          `json:"playerGameInfo,omitempty"`
	ProjectedTotalsMap           map[string]float64 `json:"projectedTotalsMap,omitempty"`
	CalculatedProjectedTotalsMap map[string]float64 `json:"calculatedProjectedTotalsMap,omitempty"`
	RemainingEventPercent        json.RawMessage    `json:"remainingEventPercent,omitempty"`
}

// FptsTotal is a fantasy-points total. Total is the numeric value; Digit/
// Decimal are pre-split display strings ("123" / "45").
type FptsTotal struct {
	Total   float64 `json:"total"`
	Digit   string  `json:"digit,omitempty"`
	Decimal string  `json:"decimal,omitempty"`
}

// LiveScoringStatsResponse is the full fxpa envelope.
type LiveScoringStatsResponse struct {
	Responses []struct {
		Data LiveScoringStats `json:"data"`
	} `json:"responses"`
}
