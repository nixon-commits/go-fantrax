package models

// FantasyTeamsData is the per-msg response body for getFantasyTeams.
type FantasyTeamsData struct {
	FantasyTeams []FantasyTeam `json:"fantasyTeams"`
}

// FantasyTeamsResponse is the full fxpa envelope for getFantasyTeams.
type FantasyTeamsResponse struct {
	Responses []struct {
		Data FantasyTeamsData `json:"data"`
	} `json:"responses"`
}
