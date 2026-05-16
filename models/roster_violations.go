package models

// RosterViolationPeriod is one scoring period where a team had a violation.
// Two kinds: `IsMinMaxViolated` (position min/max not met) and
// `IsIllegalRoster` (slot ineligibility).
type RosterViolationPeriod struct {
	Number           int    `json:"number"`
	Display          string `json:"display"` // e.g. "27 (Mon Apr 20)"
	IsMinMaxViolated bool   `json:"isMinMaxViolated,omitempty"`
	IsIllegalRoster  bool   `json:"isIllegalRoster,omitempty"`
	StartDate        int64  `json:"startDate"`
	EndData          int64  `json:"endData"` // Fantrax misspells "endDate" as "endData"
}

// TeamRosterViolations is one team's set of violated periods.
type TeamRosterViolations struct {
	TeamID  string                  `json:"teamId"`
	Periods []RosterViolationPeriod `json:"periods"`
}

// RosterViolations is the response data for getRosterViolations.
type RosterViolations struct {
	Violations       []TeamRosterViolations `json:"violations"`
	NoIllegalRosters bool                   `json:"noIllegalRosters"`
	DisplayedLists   map[string]interface{} `json:"displayedLists,omitempty"`
}

// RosterViolationsResponse is the full fxpa envelope.
type RosterViolationsResponse struct {
	Responses []struct {
		Data RosterViolations `json:"data"`
	} `json:"responses"`
}
