package auth_client

import (
	"encoding/json"
	"fmt"

	"github.com/pmurley/go-fantrax/models"
)

// GetFantasyLeagueInfo returns the league's full configuration: sport,
// season window, draft status, head-to-head flag, and the complete
// position map (every slot ID Fantrax uses for the sport). Slow-changing
// data — cache aggressively.
func (c *Client) GetFantasyLeagueInfo() (*models.FantasyLeagueInfo, error) {
	body, err := c.callFXPA("getFantasyLeagueInfo", nil, c.fxpaRefLeagueHome())
	if err != nil {
		return nil, err
	}

	var resp models.FantasyLeagueInfoResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal getFantasyLeagueInfo: %w", err)
	}
	if len(resp.Responses) == 0 {
		return nil, fmt.Errorf("getFantasyLeagueInfo: no responses")
	}
	return &resp.Responses[0].Data, nil
}
