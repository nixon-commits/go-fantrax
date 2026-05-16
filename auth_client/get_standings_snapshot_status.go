package auth_client

import (
	"encoding/json"
	"fmt"

	"github.com/pmurley/go-fantrax/models"
)

// GetStandingsSnapshotStatus returns the league-wide "is the standings
// snapshot for the most recently completed period ready?" flag. Useful as a
// trigger for weekly post-period jobs (recap, awards, etc.).
func (c *Client) GetStandingsSnapshotStatus() (*models.StandingsSnapshotStatus, error) {
	body, err := c.callFXPA("getStandingsSnapshotStatus", nil, c.fxpaRefLeagueHome())
	if err != nil {
		return nil, err
	}

	var resp models.StandingsSnapshotStatusResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal getStandingsSnapshotStatus: %w", err)
	}
	if len(resp.Responses) == 0 {
		return nil, fmt.Errorf("getStandingsSnapshotStatus: no responses")
	}
	return &resp.Responses[0].Data, nil
}
