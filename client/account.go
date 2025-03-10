package client

func (c *IGClient) AllAccounts() (*AccountDetailsResponse, error) {
	return getRequest[AccountDetailsResponse](c, v1, "accounts")
}

func (c *IGClient) AccountSettings() (*AccountSettingsResponse, error) {
	return getRequest[AccountSettingsResponse](c, v1, "accounts/preferences")
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
