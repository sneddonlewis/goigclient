package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/sneddonlewis/goigclient/internal/rest"
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

// LoginEncrypted creates a trading session with encrypted credentials.
// When encrypting password, first call LoginKey method to retrieve the
// base64 encryption key. If not encrypting password, it's simpler to
// use the Login method.
//
// This method uses **version 2** of the IG API.
//
// Parameters:
//   - request: The authentication request containing login credentials.
//
// Returns:
//   - A pointer to LoginEncryptedResponse containing account information.
//   - An error if the request fails.
func (c *IGClient) LoginEncrypted(request LoginEncryptedRequest) (*LoginEncryptedResponse, error) {
	return rest.NewRequest[LoginEncryptedResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v2,
		http.MethodPost,
		"session",
	).
		WithBody(request).
		Execute()
}

// Logout logs out of the current session.
//
// This method uses **version 1** of the IG API.
// Returns an error if the request fails.
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

	if resp.StatusCode != 204 {
		return errors.New(fmt.Sprintf("Failed to log out: %s", string(body)))
	}

	return nil
}

// Session retrieves the user's session details and optionally tokens.
//
// This method uses **version 1** of the IG API.
//
// Parameters:
//   - fetchSessionTokens: Whether to fetch session token headers (default: false, optional).
//
// Returns a pointer to SessionResponse containing the session details, or
// an error if the request fails.
func (c *IGClient) Session(fetchSessionTokens *bool) (*SessionResponse, error) {
	params := make(map[string]string)
	if fetchSessionTokens != nil {
		params["fetchSessionTokens"] = fmt.Sprintf("%t", *fetchSessionTokens)
	}

	return rest.NewRequest[SessionResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodGet,
		"session",
	).
		WithQueryParams(params).
		Execute()
}

// SwitchAccount switches the active account and optionally sets the default account.
//
// This method uses **version 1** of the IG API.
//
// Parameters:
//   - request: The account switch request containing the account ID and optional default setting.
//
// Returns a pointer to SwitchAccountResponse containing the response details, or
// an error if the request fails.
func (c *IGClient) SwitchAccount(request SwitchAccountRequest) (*SwitchAccountResponse, error) {
	return rest.NewRequest[SwitchAccountResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodPut,
		"session",
	).
		WithBody(request).
		Execute()
}

// LoginWithKey retrieves the encryption key for secure password transmission.
// The retrieved encryption key can then be used to encrypt the user password.
//
// This method uses **version 1** of the IG API.
//
// Returns a pointer to LoginWithKeyResponse containing the encryption key and timestamp, or
// an error if the request fails.
func (c *IGClient) LoginKey() (*LoginKeyResponse, error) {
	return rest.NewRequest[LoginKeyResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodGet,
		"session/encryptionKey",
	).
		Execute()
}

// RefreshSession refreshes a trading session, obtaining new session tokens.
//
// This method uses **version 1** of the IG API.
//
// Parameters:
//   - request: The refresh session request containing the refresh token.
//
// Returns a pointer to RefreshSessionResponse containing the new session tokens, or
// an error if the request fails.
func (c *IGClient) RefreshSession(request RefreshSessionRequest) (*RefreshSessionResponse, error) {
	return rest.NewRequest[RefreshSessionResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodPost,
		"session/refresh-token",
	).
		WithBody(request).
		Execute()
}
