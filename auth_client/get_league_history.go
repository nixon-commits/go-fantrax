package auth_client

import (
	"encoding/json"
	"fmt"

	"github.com/pmurley/go-fantrax/models"
)

// GetLeagueHistory returns the league's all-time history summary: past
// seasons, career win%, championships, and the displayed-tab metadata
// Fantrax's history page renders.
func (c *Client) GetLeagueHistory() (*models.LeagueHistory, error) {
	body, err := c.callFXPA("getLeagueHistory", nil, c.fxpaRefLeagueHome())
	if err != nil {
		return nil, err
	}

	var resp models.LeagueHistoryResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal getLeagueHistory: %w", err)
	}
	if len(resp.Responses) == 0 {
		return nil, fmt.Errorf("getLeagueHistory: no responses")
	}
	return &resp.Responses[0].Data, nil
}
