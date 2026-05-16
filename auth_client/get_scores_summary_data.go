package auth_client

import (
	"encoding/json"
	"fmt"

	"github.com/pmurley/go-fantrax/models"
)

// GetScoresSummaryData returns the real-world scoreboard ticker that the
// live-scoring page shows along the bottom: every game today with home/away
// score, status (e.g. "Bot 7th, 0 Out, 0-0"), and a pregame start time.
func (c *Client) GetScoresSummaryData() (*models.ScoresSummary, error) {
	body, err := c.callFXPA("getScoresSummaryData", nil, c.fxpaRefLeagueHome())
	if err != nil {
		return nil, err
	}

	var resp models.ScoresSummaryResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal getScoresSummaryData: %w", err)
	}
	if len(resp.Responses) == 0 {
		return nil, fmt.Errorf("getScoresSummaryData: no responses")
	}
	return &resp.Responses[0].Data, nil
}
