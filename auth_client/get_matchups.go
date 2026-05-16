package auth_client

import (
	"encoding/json"
	"fmt"

	"github.com/pmurley/go-fantrax/models"
)

// GetMatchups returns the full season matchup schedule (every period, every
// home/away pairing). This is the dedicated fxpa method the matchup-preview
// UI uses; lighter than GetAllMatchups which overloads getStandings.
func (c *Client) GetMatchups() (*models.MatchupsList, error) {
	body, err := c.callFXPA("getMatchups", nil, c.fxpaRefLeagueHome())
	if err != nil {
		return nil, err
	}

	var resp models.MatchupsListResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal getMatchups: %w", err)
	}
	if len(resp.Responses) == 0 {
		return nil, fmt.Errorf("getMatchups: no responses")
	}
	return &resp.Responses[0].Data, nil
}
