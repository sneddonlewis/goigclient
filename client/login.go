package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Login creates a trading session, obtaining session tokens for
// subsequent API access. Please note, region-specific login restrictions may apply.
//
// This method uses **version 3** of the IG API.
// It returns a pointer to AuthResponse containing the account ID and tokens,
// or an error if the request fails.
func (c *IGClient) Login() (*AuthResponse, error) {
	authData := map[string]string{
		"identifier": c.Username,
		"password":   c.Password,
	}
	payload, _ := json.Marshal(authData)

	req, err := http.NewRequest("POST", c.BaseURL+"/session", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-IG-API-KEY", c.APIKey)
	req.Header.Set("Version", "3")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return nil, errors.New(fmt.Sprintf("Failed to authenticate: %s", string(body)))
	}

	var authResp AuthResponse
	err = json.Unmarshal(body, &authResp)
	if err != nil {
		return nil, err
	}

	c.AccessToken = authResp.OAuth.AccessToken
	c.RefreshToken = authResp.OAuth.RefreshToken

	c.AccountID = authResp.AccountID
	return &authResp, nil
}

// Logout logs out of the current session.
//
// This method uses **version 1** of the IG API.
// It returns an error if the request fails.
func (c *IGClient) Logout() error {
	req, err := http.NewRequest(http.MethodDelete, c.BaseURL+"/session", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	req.Header.Set("X-IG-API-KEY", c.APIKey)
	req.Header.Set("IG-ACCOUNT-ID", c.AccountID)
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)

	req.Header.Set("Version", v1)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return errors.New(fmt.Sprintf("Failed to log out: %s", string(body)))
	}

	return nil
}
