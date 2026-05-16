package auth_client

import (
	"encoding/json"
	"fmt"

	"github.com/pmurley/go-fantrax/models"
)

// GetLiveScoringStatsRequest are the args for getLiveScoringStats.
// SppID is the scoring period ID — pass "-1" for the current/most-recent
// period. NewView=true is what the UI sends for first-load of the page.
type GetLiveScoringStatsRequest struct {
	SppID   string `json:"sppId"`
	NewView bool   `json:"newView"`
}

// GetLiveScoringStats returns the period's live-scoring data — per-team
// totals, projected totals, and game-info arrays the live-scoring UI
// renders. The response is large (>800KB unfiltered); the model exposes
// only the fields needed for win-probability and matchup summaries, with
// the rest preserved on `FullRaw` for advanced callers.
//
// Pass sppId="-1" for the current period.
func (c *Client) GetLiveScoringStats(sppId string) (*models.LiveScoringStats, error) {
	req := GetLiveScoringStatsRequest{SppID: sppId, NewView: true}
	if sppId == "" {
		req.SppID = "-1"
	}
	body, err := c.callFXPA("getLiveScoringStats", req,
		fmt.Sprintf("https://www.fantrax.com/fantasy/league/%s/livescoring", c.LeagueID))
	if err != nil {
		return nil, err
	}

	var resp models.LiveScoringStatsResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal getLiveScoringStats: %w", err)
	}
	if len(resp.Responses) == 0 {
		return nil, fmt.Errorf("getLiveScoringStats: no responses")
	}

	out := resp.Responses[0].Data
	// Stash the full raw per-msg data for callers that need fields outside
	// our deliberately-narrow typed view.
	var envelope struct {
		Responses []struct {
			Data json.RawMessage `json:"data"`
		} `json:"responses"`
	}
	if err := json.Unmarshal(body, &envelope); err == nil && len(envelope.Responses) > 0 {
		out.FullRaw = envelope.Responses[0].Data
	}
	return &out, nil
}
