package auth_client

import (
	"encoding/json"
	"fmt"

	"github.com/pmurley/go-fantrax/models"
)

// GetLeagueTransactions calls the newer getTransactions fxpa method, which
// returns a league-wide transaction list (claims, drops, trades) with
// display metadata. Distinct from GetTransactionDetailsHistory: that wraps
// the older per-team paged endpoint and the two are not interchangeable.
//
// The transactions field is returned as raw JSON for now — the shape varies
// by tx type and consumers can model whichever subset they care about.
func (c *Client) GetLeagueTransactions() (*models.LeagueTransactionsData, error) {
	body, err := c.callFXPA("getTransactions", nil, c.fxpaRefLeagueHome())
	if err != nil {
		return nil, err
	}

	var resp models.LeagueTransactionsResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal getTransactions: %w", err)
	}
	if len(resp.Responses) == 0 {
		return nil, fmt.Errorf("getTransactions: no responses")
	}
	return &resp.Responses[0].Data, nil
}
