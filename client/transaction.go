package client

import (
	"net/http"

	"github.com/sneddonlewis/goigclient/internal/rest"
)

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
