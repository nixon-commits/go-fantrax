package models

// MatchupPairing is a single home/away matchup within a scoring period.
type MatchupPairing struct {
	HomeTeam FantasyTeam `json:"homeTeam"`
	AwayTeam FantasyTeam `json:"awayTeam"`
}

// MatchupPeriod is the set of pairings for one scoring period.
type MatchupPeriod struct {
	Number   int              `json:"number"`
	Matchups []MatchupPairing `json:"matchups"`
}

// MatchupsList is the response data for getMatchups. It is a complete map of
// the season's matchups — every period, every pairing — and is lighter than
// the standings overload of getAllMatchups.
type MatchupsList struct {
	RegularSeasonEndPeriod int             `json:"regularSeasonEndPeriod"`
	FantasyTeams           []FantasyTeam   `json:"fantasyTeams"`
	Periods                []MatchupPeriod `json:"periods"`
}

// MatchupsListResponse is the full fxpa envelope.
type MatchupsListResponse struct {
	Responses []struct {
		Data MatchupsList `json:"data"`
	} `json:"responses"`
}
