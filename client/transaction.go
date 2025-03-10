package client

func (c *IGClient) TransactionHistory(
	trxType,
	fromDate,
	toDate string,
) (*TransactionsHistoryResponse, error) {
	return getRequestWithParams[TransactionsHistoryResponse](
		c,
		v1,
		"history/transactions",
		trxType,
		fromDate,
		toDate,
	)
}

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
