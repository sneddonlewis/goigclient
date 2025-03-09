package client

const (
	v1 = "1"
	v2 = "2"
	v3 = "3"
)

// Dealing

func (c *IGClient) AllOpenPositions() (*OpenPositionsResponse, error) {
	return getRequest[OpenPositionsResponse](c, v2, "positions")
}

func (c *IGClient) OneOpenPosition(dealID string) (*OpenPositionResponse, error) {
	return getRequestWithParams[OpenPositionResponse](c, v2, "positions", dealID)
}

func (c *IGClient) CreateOTCPosition() error {
	return nil
}

func (c *IGClient) ChangeOTCPosition() error {
	return nil
}

func (c *IGClient) CloseOTCPosition() error {
	return nil
}

func (c *IGClient) AllWorkingOrders() (*WorkingOrdersResponse, error) {
	return getRequest[WorkingOrdersResponse](c, v2, "workingorders")
}

func (c *IGClient) OneWorkingOrder() error {
	return nil
}

func (c *IGClient) CreateOTCWorkingOrder() error {
	return nil
}

func (c *IGClient) ChangeOTCWorkingOrder() error {
	return nil
}

func (c *IGClient) CloseOTCWorkingOrder() error {
	return nil
}

// Markets

func (c *IGClient) OneMarket(epic string) (*MarketPricesResponse, error) {
	return getRequestWithParams[MarketPricesResponse](c, v3, "markets", epic)
}

// Account

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

func (c *IGClient) AllAccounts() (*AccountDetailsResponse, error) {
	return getRequest[AccountDetailsResponse](c, v1, "accounts")
}

func (c *IGClient) AccountSettings() (*AccountSettingsResponse, error) {
	return getRequest[AccountSettingsResponse](c, v1, "accounts/preferences")
}
