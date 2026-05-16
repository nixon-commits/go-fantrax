package auth_client

import (
	"encoding/json"
	"fmt"

	"github.com/pmurley/go-fantrax/models"
)

// GetRosterViolations returns league-wide roster violations bucketed by
// team and period. Each violation is either a position min/max miss
// (`IsMinMaxViolated`) or a slot ineligibility (`IsIllegalRoster`).
//
// Complements GetIllegalRosterOverview (HTML scrape of the commissioner
// page) — this is the JSON API that backs the same data.
func (c *Client) GetRosterViolations() (*models.RosterViolations, error) {
	body, err := c.callFXPA("getRosterViolations", nil, c.fxpaRefLeagueHome())
	if err != nil {
		return nil, err
	}

	var resp models.RosterViolationsResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal getRosterViolations: %w", err)
	}
	if len(resp.Responses) == 0 {
		return nil, fmt.Errorf("getRosterViolations: no responses")
	}
	return &resp.Responses[0].Data, nil
}
