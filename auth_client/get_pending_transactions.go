package auth_client

import (
	"encoding/json"
	"fmt"

	"github.com/pmurley/go-fantrax/models"
)

// GetPendingTransactions returns the league's pending trades (and other
// pending tx types). Each TradeInfoList entry has the participating teams,
// player moves, and any illegal-roster warnings Fantrax has flagged for
// the trade — useful for in-flight trade monitoring.
func (c *Client) GetPendingTransactions() (*models.PendingTransactions, error) {
	body, err := c.callFXPA("getPendingTransactions", nil, c.fxpaRefLeagueHome())
	if err != nil {
		return nil, err
	}

	var resp models.PendingTransactionsResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal getPendingTransactions: %w", err)
	}
	if len(resp.Responses) == 0 {
		return nil, fmt.Errorf("getPendingTransactions: no responses")
	}
	return &resp.Responses[0].Data, nil
}
