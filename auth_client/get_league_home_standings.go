package auth_client

import (
	"encoding/json"
	"fmt"

	"github.com/pmurley/go-fantrax/models"
)

// GetLeagueHomeStandings returns the trimmed standings widget that the
// league-home page renders. Lighter than getStandings — useful when you
// just need the rank/record/points columns.
func (c *Client) GetLeagueHomeStandings() (*models.LeagueHomeStandings, error) {
	body, err := c.callFXPA("getLeagueHomeStandings", nil, c.fxpaRefLeagueHome())
	if err != nil {
		return nil, err
	}

	var resp models.LeagueHomeStandingsResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal getLeagueHomeStandings: %w", err)
	}
	if len(resp.Responses) == 0 {
		return nil, fmt.Errorf("getLeagueHomeStandings: no responses")
	}
	return &resp.Responses[0].Data, nil
}
