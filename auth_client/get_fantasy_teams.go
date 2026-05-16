package auth_client

import (
	"encoding/json"
	"fmt"

	"github.com/pmurley/go-fantrax/models"
)

// GetFantasyTeams returns the lightweight list of all fantasy teams in the
// league. Faster and cleaner than scraping team IDs out of getStandings.
func (c *Client) GetFantasyTeams() ([]models.FantasyTeam, error) {
	body, err := c.callFXPA("getFantasyTeams", nil, c.fxpaRefLeagueHome())
	if err != nil {
		return nil, err
	}

	var resp models.FantasyTeamsResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal getFantasyTeams: %w", err)
	}
	if len(resp.Responses) == 0 {
		return nil, fmt.Errorf("getFantasyTeams: no responses")
	}
	return resp.Responses[0].Data.FantasyTeams, nil
}
