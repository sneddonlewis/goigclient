package client

type TransactionsHistoryResponse struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	CashTransaction bool   `json:"cashTransaction"`
	CloseLevel      string `json:"closeLevel"`
	Currency        string `json:"currency"`
	Date            string `json:"date"`
	InstrumentName  string `json:"instrumentName"`
	OpenLevel       string `json:"openLevel"`
	Period          string `json:"period"`
	ProfitAndLoss   string `json:"profitAndLoss"`
	Reference       string `json:"reference"`
	Size            string `json:"size"`
	TransactionType string `json:"transactionType"`
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

type AccountPreferences struct {
	TrailingStopsEnabled bool `json:"trailingStopsEnabled"`
}

type UpdateAccountPreferencesResponse struct {
	Status string `json:"status"`
}
