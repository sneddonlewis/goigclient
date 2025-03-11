package client

import (
	"net/http"

	"github.com/sneddonlewis/goigclient/internal/rest"
)

func (c *IGClient) AllAccounts() (*AccountDetailsResponse, error) {
	return rest.NewRequest[AccountDetailsResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodGet,
		"accounts",
	).Execute()
}

func (c *IGClient) AccountSettings() (*AccountSettingsResponse, error) {
	return rest.NewRequest[AccountSettingsResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodGet,
		"accounts/preferences",
	).
		Execute()
}

type AccountDetailsResponse struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	AccountID       string  `json:"accountId"`
	AccountName     string  `json:"accountName"`
	AccountAlias    *string `json:"accountAlias"`
	Status          string  `json:"status"`
	AccountType     string  `json:"accountType"`
	Preferred       bool    `json:"preferred"`
	Balance         Balance `json:"balance"`
	Currency        string  `json:"currency"`
	CanTransferFrom bool    `json:"canTransferFrom"`
	CanTransferTo   bool    `json:"canTransferTo"`
}

type Balance struct {
	Balance    float64 `json:"balance"`
	Deposit    float64 `json:"deposit"`
	ProfitLoss float64 `json:"profitLoss"`
	Available  float64 `json:"available"`
}

type AccountSettingsResponse struct {
	TrailingStopsEnabled bool `json:"trailingStopsEnabled"`
}
