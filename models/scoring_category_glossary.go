package models

import "encoding/json"

// ScoringCategoryGroup is a glossary section. Two groups appear for most
// sports: indiv (per-player) and team. Within each, "stats" lists every
// scoring category with its abbreviation, full name, and description.
type ScoringCategoryGroup struct {
	ID    string                 `json:"id"`
	Name  string                 `json:"name,omitempty"`
	Stats []ScoringCategoryEntry `json:"stats,omitempty"`
	// Other fields vary by sport — preserved as raw for forward-compat.
	Raw json.RawMessage `json:"-"`
}

// ScoringCategoryEntry is one row in the glossary.
type ScoringCategoryEntry struct {
	ScipID      string `json:"scipId"`
	Name        string `json:"name"`
	ShortName   string `json:"shortName,omitempty"`
	Description string `json:"description,omitempty"`
}

// ScoringCategoryGlossary is the response data for getScoringCategoryGlossary.
type ScoringCategoryGlossary struct {
	SportName       string          `json:"sportName"`
	IsSupport       bool            `json:"isSupport"`
	TeamCategories  json.RawMessage `json:"teamCategories"`
	IndivCategories json.RawMessage `json:"indivCategories"`
}

// ScoringCategoryGlossaryResponse is the full fxpa envelope.
type ScoringCategoryGlossaryResponse struct {
	Responses []struct {
		Data ScoringCategoryGlossary `json:"data"`
	} `json:"responses"`
}
