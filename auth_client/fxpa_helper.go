package auth_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// callFXPA POSTs a single-msg request to /fxpa/req?leagueId=... and returns
// the response body. It centralises the boilerplate (envelope fields,
// content-type, status check, body read) that every new fxpa wrapper would
// otherwise duplicate.
//
// `refUrl` should match the screen the message corresponds to — Fantrax uses
// it for telemetry and occasionally for authorization on commissioner-only
// methods. Pass an empty string for endpoints that don't care.
func (c *Client) callFXPA(method string, data interface{}, refUrl string) ([]byte, error) {
	if data == nil {
		data = map[string]interface{}{}
	}
	body := map[string]interface{}{
		"msgs": []FantraxMessage{
			{Method: method, Data: data},
		},
		"uiv":    3,
		"refUrl": refUrl,
		"dt":     0,
		"at":     0,
		"av":     "0.0",
		"tz":     "America/Chicago",
		"v":      "179.0.1",
	}

	jsonStr, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("marshal %s request: %w", method, err)
	}

	req, err := http.NewRequest("POST", "https://www.fantrax.com/fxpa/req?leagueId="+c.LeagueID, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, fmt.Errorf("build %s request: %w", method, err)
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send %s request: %w", method, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s returned status %d", method, resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read %s response: %w", method, err)
	}

	return respBody, nil
}

// fxpaRefLeagueHome returns the default refUrl pointing at the league home —
// works for every read-only fxpa method we've discovered.
func (c *Client) fxpaRefLeagueHome() string {
	return fmt.Sprintf("https://www.fantrax.com/fantasy/league/%s/home", c.LeagueID)
}
