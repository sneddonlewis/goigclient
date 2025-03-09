package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type OAuthToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

type AuthResponse struct {
	AccountID             string     `json:"accountId"`
	ClientID              string     `json:"clientId"`
	LightstreamerEndpoint string     `json:"lightstreamerEndpoint"`
	OAuth                 OAuthToken `json:"oauthToken"`
	TimezoneOffset        int        `json:"timezoneOffset"`
}

func (c *IGClient) Authenticate() error {
	authData := map[string]string{
		"identifier": c.Username,
		"password":   c.Password,
	}
	payload, _ := json.Marshal(authData)

	req, err := http.NewRequest("POST", c.BaseURL+"/session", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-IG-API-KEY", c.APIKey)
	req.Header.Set("Version", "3")

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
		return errors.New(fmt.Sprintf("Failed to authenticate: %s", string(body)))
	}

	var authResp AuthResponse
	err = json.Unmarshal(body, &authResp)
	if err != nil {
		return err
	}

	c.AccessToken = authResp.OAuth.AccessToken
	c.RefreshToken = authResp.OAuth.RefreshToken

	c.AccountID = authResp.AccountID
	return nil
}
