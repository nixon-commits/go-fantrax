package models

// LeagueHistorySeason summarises a single past season of a league. Most
// fields are pass-throughs from Fantrax — keep as map[string]interface{} for
// forward-compat until specific consumers want them typed.
type LeagueHistorySeason = map[string]interface{}

// LeagueHistoryTab describes a tab in the league-history UI (Seasons, Records, Teams).
type LeagueHistoryTab struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// LeagueHistoryDisplayedSelections mirrors the view-state Fantrax echoes back.
type LeagueHistoryDisplayedSelections struct {
	View                  string `json:"view"`
	ShowTabs              bool   `json:"showTabs"`
	IncludePlayoffPeriods bool   `json:"includePlayoffPeriods"`
	IncludeMergedPeriods  bool   `json:"includeMergedPeriods"`
}

// LeagueHistory is the response data for getLeagueHistory.
type LeagueHistory struct {
	IsHeadToHead        bool                             `json:"isHeadToHead"`
	Seasons             []LeagueHistorySeason            `json:"seasons"`
	AnyPlayoffsUsed     bool                             `json:"anyPlayoffsUsed"`
	LeagueHistoryID     string                           `json:"leagueHistoryId"`
	FantasyPeriodName   string                           `json:"fantasyPeriodName"`
	DisplayedSelections LeagueHistoryDisplayedSelections `json:"displayedSelections"`
	DisplayedLists      struct {
		Tabs []LeagueHistoryTab `json:"tabs"`
	} `json:"displayedLists"`
}

// LeagueHistoryResponse is the full fxpa envelope.
type LeagueHistoryResponse struct {
	Responses []struct {
		Data LeagueHistory `json:"data"`
	} `json:"responses"`
}
