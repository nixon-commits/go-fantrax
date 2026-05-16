package auth_client

import (
	"encoding/json"
	"fmt"

	"github.com/pmurley/go-fantrax/models"
)

// GetScoringCategoryGlossary returns the full glossary of scoring
// categories for the league's sport — every short name, full name, and
// description Fantrax shows in the scoring-settings reference panel.
func (c *Client) GetScoringCategoryGlossary() (*models.ScoringCategoryGlossary, error) {
	body, err := c.callFXPA("getScoringCategoryGlossary", nil, c.fxpaRefLeagueHome())
	if err != nil {
		return nil, err
	}

	var resp models.ScoringCategoryGlossaryResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal getScoringCategoryGlossary: %w", err)
	}
	if len(resp.Responses) == 0 {
		return nil, fmt.Errorf("getScoringCategoryGlossary: no responses")
	}
	return &resp.Responses[0].Data, nil
}
