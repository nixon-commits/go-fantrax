package models

// StandingsRow is a single row in the league-home standings widget.
// Numeric fields come over the wire as strings (Fantrax preserves trailing
// zeros for display); callers parse as needed.
type StandingsRow struct {
	Rank          int    `json:"rank"`
	TeamID        string `json:"teamId"`
	Team          string `json:"team"`
	Score         string `json:"score"`         // e.g. "12-0-0"
	WinPercentage string `json:"winPercentage"` // e.g. "1.000"
	Points        string `json:"points"`        // e.g. "9806"
	GamesBack     string `json:"gamesBack"`     // e.g. "0"
	Commish       bool   `json:"commish,omitempty"`
}

// LeagueHomeStandings is the response data for getLeagueHomeStandings.
type LeagueHomeStandings struct {
	Standings struct {
		// StatsTable holds standings grouped by view. The most useful key is
		// "COMBINED"; other keys appear for split standings (e.g. division).
		StatsTable []map[string][]StandingsRow `json:"statsTable"`
	} `json:"standings"`
}

// LeagueHomeStandingsResponse is the full fxpa envelope.
type LeagueHomeStandingsResponse struct {
	Responses []struct {
		Data LeagueHomeStandings `json:"data"`
	} `json:"responses"`
}

// CombinedStandings returns the COMBINED group rows in rank order, or nil if
// the response has no COMBINED group.
func (l LeagueHomeStandings) CombinedStandings() []StandingsRow {
	for _, m := range l.Standings.StatsTable {
		if rows, ok := m["COMBINED"]; ok {
			return rows
		}
	}
	return nil
}
