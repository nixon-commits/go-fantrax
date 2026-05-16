package models

// ScoreSummary is one real-world game entry in the scoreboard ticker.
type ScoreSummary struct {
	EventID             string `json:"eventId"`
	HomeTeamID          string `json:"homeTeamId"`
	AwayTeamID          string `json:"awayTeamId"`
	HomeScore           int    `json:"homeScore"`
	AwayScore           int    `json:"awayScore"`
	StatusCode          string `json:"statusCode,omitempty"`         // e.g. "PREGAME"
	StatusDisplay       string `json:"statusDisplay,omitempty"`      // e.g. "Bot 7th, 0 Out, 0-0"
	StatusDisplayBrief  string `json:"statusDisplayBrief,omitempty"` // e.g. "Bot 7th"
	StartTime           int64  `json:"startTime,omitempty"`          // unix-ms for pregame
	PreviewBoxscoreLink string `json:"previewBoxscoreLink,omitempty"`
}

// ScoresSummary is the response data for getScoresSummaryData.
type ScoresSummary struct {
	Scores []ScoreSummary `json:"scores"`
}

// ScoresSummaryResponse is the full fxpa envelope.
type ScoresSummaryResponse struct {
	Responses []struct {
		Data ScoresSummary `json:"data"`
	} `json:"responses"`
}
