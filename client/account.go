package client

import (
	"net/http"

	"github.com/sneddonlewis/goigclient/internal/rest"
)

// Accounts returns a list of the logged-in client's accounts.
//
// This method uses **version 1** of the IG API.
// It returns a pointer to AccountDetailsResponse containing the account details,
// or an error if the request fails.
func (c *IGClient) Accounts() (*AccountDetailsResponse, error) {
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

// AccountPreferences returns all account related preferences.
//
// This method uses **version 1** of the IG API.
// It returns a pointer to AccountPreferences, or an error if the request fails.
func (c *IGClient) AccountPreferences() (*AccountPreferences, error) {
	return rest.NewRequest[AccountPreferences](
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

// UpdateAccountPreferences.
//
// This method uses **version 1** of the IG API.
// It returns a pointer to UpdateAccountPreferencesResponse, or an error if the request fails.
func (c *IGClient) UpdateAccountPreferences(request AccountPreferences) (*UpdateAccountPreferencesResponse, error) {
	return rest.NewRequest[UpdateAccountPreferencesResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodPut,
		"accounts/preferences",
	).
		WithBody(request).
		Execute()
}

// TransactionHistory returns the transaction history for the specified.
// transaction type and given date range.
//
// This method uses **version 1** of the IG API.
// It returns a pointer to TransactionHistoryResponse, or an error if the request fails.
func (c *IGClient) TransactionHistory(
	trxType,
	fromDate,
	toDate string,
) (*TransactionsHistoryResponse, error) {
	return rest.NewRequest[TransactionsHistoryResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v2,
		http.MethodGet,
		"history/transactions",
	).
		WithParams(trxType, fromDate, toDate).
		Execute()
}
