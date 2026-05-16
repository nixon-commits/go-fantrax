package models

import "encoding/json"

// FantasyLeagueSettings is the slice of league config exposed via
// getFantasyLeagueInfo. The shape is broader than what's typed here — extra
// fields are dropped silently by Go's json package, which is what we want
// since Fantrax adds fields over time. The fields below are the stable ones
// most callers reach for.
type FantasyLeagueSettings struct {
	LeagueName      string              `json:"leagueName"`
	LeagueHistoryID string              `json:"leagueHistoryId"`
	TeamName        string              `json:"teamName"`
	MyDefaultTeamID string              `json:"myDefaultTeamId"`
	Sport           string              `json:"sport"`        // e.g. "MLB"
	SportID         string              `json:"sportId"`      // e.g. "001"
	GeneralSport    string              `json:"generalSport"` // e.g. "BASEBALL"
	GeneralSportID  string              `json:"generalSportId"`
	Subtitle        string              `json:"subtitle"` // e.g. "2026 MLB"
	HeadToHead      bool                `json:"headToHead"`
	Football        bool                `json:"football"`
	Commissioner    bool                `json:"commissioner"`
	IsMember        bool                `json:"isMember"`
	DraftStatus     string              `json:"draftStatus"`
	ChatEnabled     bool                `json:"chatEnabled"`
	HideLeagueLogo  bool                `json:"hideLeagueLogo"`
	TeamLogo        string              `json:"teamLogo,omitempty"`
	MyTeamIDs       []string            `json:"myTeamIds"`
	Season          FantasyLeagueSeason `json:"season"`
}

// FantasyLeagueSeason describes the league's current season window.
type FantasyLeagueSeason struct {
	ID               string `json:"id"`
	DisplayName      string `json:"displayName"`
	DisplayNameShort string `json:"displayNameShort"`
	DisplayYear      string `json:"displayYear"`
	StartDate        int64  `json:"startDate"`
	EndDate          int64  `json:"endDate"`
	SeasonType       struct {
		Code string `json:"code"` // e.g. "REGULAR"
	} `json:"seasonType"`
}

// PositionInfo is one position-slot definition: short/long name plus the
// numeric ID Fantrax uses everywhere else in the API.
type PositionInfo struct {
	ID        string `json:"id"`        // "001"
	Name      string `json:"name"`      // "Catcher"
	ShortName string `json:"shortName"` // "C"
	SortOrder int    `json:"sortOrder"`
	ColorKey  string `json:"colorKey,omitempty"`
}

// FantasyLeagueInfo is the response data for getFantasyLeagueInfo.
type FantasyLeagueInfo struct {
	FantasySettings FantasyLeagueSettings   `json:"fantasySettings"`
	PositionMap     map[string]PositionInfo `json:"positionMap"`
	// Less-stable nested config kept raw.
	Raw json.RawMessage `json:"-"`
}

// FantasyLeagueInfoResponse is the full fxpa envelope.
type FantasyLeagueInfoResponse struct {
	Responses []struct {
		Data FantasyLeagueInfo `json:"data"`
	} `json:"responses"`
}
