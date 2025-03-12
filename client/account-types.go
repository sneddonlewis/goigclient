package client

type TransactionsHistoryResponse struct {
	Transactions []Transaction `json:"transactions"`
}

type TransactionDetailed struct {
	CashTransaction bool   `json:"cashTransaction"`
	CloseLevel      string `json:"closeLevel"`
	Currency        string `json:"currency"`
	Date            string `json:"date"`
	DateUTC         string `json:"dateUtc"`
	InstrumentName  string `json:"instrumentName"`
	OpenDateUTC     string `json:"openDateUtc"`
	OpenLevel       string `json:"openLevel"`
	Period          string `json:"period"`
	ProfitAndLoss   string `json:"profitAndLoss"`
	Reference       string `json:"reference"`
	Size            string `json:"size"`
	TransactionType string `json:"transactionType"`
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

type ActivityHistoryResponse struct {
	Activities []Activity `json:"activities"`
	Metadata   Metadata   `json:"metadata"`
}

type Activity struct {
	Channel     string  `json:"channel"`
	Date        string  `json:"date"`
	DealID      string  `json:"dealId"`
	Description string  `json:"description"`
	Details     Details `json:"details"`
	Type        string  `json:"type"`
}

type Details struct {
	Actions []Action `json:"actions"`
}

type Action struct {
	ActionType           string  `json:"actionType"`
	AffectedDealID       string  `json:"affectedDealId"`
	Currency             string  `json:"currency"`
	DealReference        string  `json:"dealReference"`
	Direction            string  `json:"direction"`
	GoodTillDate         string  `json:"goodTillDate"`
	GuaranteedStop       bool    `json:"guaranteedStop"`
	Level                float64 `json:"level"`
	LimitDistance        float64 `json:"limitDistance"`
	LimitLevel           float64 `json:"limitLevel"`
	MarketName           string  `json:"marketName"`
	Size                 float64 `json:"size"`
	StopDistance         float64 `json:"stopDistance"`
	StopLevel            float64 `json:"stopLevel"`
	TrailingStep         float64 `json:"trailingStep"`
	TrailingStopDistance float64 `json:"trailingStopDistance"`
	Epic                 string  `json:"epic"`
	Period               string  `json:"period"`
	Status               string  `json:"status"`
}

type Metadata struct {
	Paging Paging `json:"paging"`
}

type Paging struct {
	Next string `json:"next"`
	Size int    `json:"size"`
}

type ActivityHistoryDateRangeResponse struct {
	Activities []ActivityDateRange `json:"activities"`
}

type ActivityDateRange struct {
	ActionStatus      string `json:"actionStatus"`
	Activity          string `json:"activity"`
	ActivityHistoryID string `json:"activityHistoryId"`
	Channel           string `json:"channel"`
	Currency          string `json:"currency"`
	Date              string `json:"date"`
	DealID            string `json:"dealId"`
	Epic              string `json:"epic"`
	Level             string `json:"level"`
	Limit             string `json:"limit"`
	MarketName        string `json:"marketName"`
	Period            string `json:"period"`
	Result            string `json:"result"`
	Size              string `json:"size"`
	Stop              string `json:"stop"`
	StopType          string `json:"stopType"`
	Time              string `json:"time"`
}

type TransactionHistoryResponse struct {
	Transactions []TransactionDetailed `json:"transactions"`
	Metadata     Metadata              `json:"metadata"`
}

type TransactionByPeriodResponse struct {
	Transactions []Transaction `json:"transactions"`
}
