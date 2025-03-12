package client

import (
	"fmt"
	"net/http"
	"time"

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

// TransactionHistory returns the transaction history for the specified
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

// GetActivityHistory retrieves the account activity history.
//
// This method uses **version 3** of the IG API.
// It allows filtering by date range, deal ID, FIQL filter, page size, and additional details.
//
// Parameters:
//   - from: Start date (optional).
//   - to: End date (default is current time, optional).
//   - detailed: Whether to retrieve additional details (default: false, optional).
//   - dealId: Filter by deal ID (optional).
//   - filter: FIQL filter (supported operators: == | != | , | ; , optional).
//   - pageSize: Number of results per page (default: 50, min: 10, max: 500, optional).
//
// Returns:
//   - A pointer to ActivityHistoryResponse containing the activity records.
//   - An error if the request fails.
func (c *IGClient) GetActivityHistory(
	from, to *time.Time,
	detailed *bool,
	dealId, filter *string,
	pageSize *int,
) (*ActivityHistoryResponse, error) {
	params := make(map[string]string)
	if from != nil {
		params["from"] = from.Format(time.RFC3339)
	}
	if to != nil {
		params["to"] = to.Format(time.RFC3339)
	}
	if detailed != nil {
		params["detailed"] = fmt.Sprintf("%t", *detailed)
	}
	if dealId != nil {
		params["dealId"] = *dealId
	}
	if filter != nil {
		params["filter"] = *filter
	}
	if pageSize != nil {
		params["pageSize"] = fmt.Sprintf("%d", *pageSize)
	}

	return rest.NewRequest[ActivityHistoryResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v3,
		http.MethodGet,
		"history/activity",
	).
		WithQueryParams(params).
		Execute()
}

// GetActivityHistoryDateRange retrieves the account activity history for a specific date range.
//
// This method uses **version 1** of the IG API.
//
// Parameters:
//   - fromDate: Start date in format dd-mm-yyyy.
//   - toDate: End date in format dd-mm-yyyy.
//
// Returns:
//   - A pointer to ActivityHistoryDateRangeResponse containing the activity records.
//   - An error if the request fails.
func (c *IGClient) GetActivityHistoryDateRange(fromDate, toDate string) (*ActivityHistoryDateRangeResponse, error) {
	return rest.NewRequest[ActivityHistoryDateRangeResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodGet,
		"history/activity",
	).
		WithParams(fromDate, toDate).
		Execute()
}

// TransactionHistoryQuery retrieves the transaction history. Returns the minute
// prices within the last 10 minutes by default.
//
// This method uses **version 2** of the IG API.
// It allows filtering by transaction type, date range, and pagination options.
//
// Parameters:
//   - transactionType: Transaction type (default: ALL, optional).
//   - from: Start date (optional).
//   - to: End date (date without time refers to end of day, optional).
//   - maxSpanSeconds: Limits timespan in seconds through to current time (default: 600, optional).
//   - pageSize: Number of results per page (default: 20, optional, disable paging = 0).
//   - pageNumber: Page number (default: 1, optional).
//
// Returns:
//   - A pointer to TransactionHistoryResponse containing the transaction records.
//   - An error if the request fails.
func (c *IGClient) TransactionHistoryQuery(
	transactionType *string,
	from, to *string,
	maxSpanSeconds, pageSize, pageNumber *int,
) (*TransactionHistoryResponse, error) {
	params := make(map[string]string)
	if transactionType != nil {
		params["type"] = *transactionType
	}
	if from != nil {
		params["from"] = *from
	}
	if to != nil {
		params["to"] = *to
	}
	if maxSpanSeconds != nil {
		params["maxSpanSeconds"] = fmt.Sprintf("%d", *maxSpanSeconds)
	}
	if pageSize != nil {
		params["pageSize"] = fmt.Sprintf("%d", *pageSize)
	}
	if pageNumber != nil {
		params["pageNumber"] = fmt.Sprintf("%d", *pageNumber)
	}

	return rest.NewRequest[TransactionHistoryResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v2,
		http.MethodGet,
		"history/transactions",
	).
		WithQueryParams(params).
		Execute()
}

// TransactionByPeriod retrieves the transaction history for a specified transaction type and period.
//
// This method uses **version 1** of the IG API.
//
// Parameters:
//   - transactionType: Transaction type (ALL, ALL_DEAL, DEPOSIT, WITHDRAWAL).
//   - lastPeriod: Interval in milliseconds.
//
// Returns:
//   - A pointer to TransactionByPeriodResponse containing the transaction records.
//   - An error if the request fails.
func (c *IGClient) TransactionByPeriod(transactionType string, lastPeriod int64) (*TransactionByPeriodResponse, error) {

	return rest.NewRequest[TransactionByPeriodResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodGet,
		"history/transactions",
	).
		WithParams(transactionType, fmt.Sprint(lastPeriod)).
		Execute()
}
