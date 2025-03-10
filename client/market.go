package client

func (c *IGClient) OneMarket(epic string) (*MarketPricesResponse, error) {
	return getRequestWithParams[MarketPricesResponse](c, v3, "markets", epic)
}

type MarketSnapshot struct {
	Bid    float64 `json:"bid"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Status string  `json:"marketStatus"`
}

type MarketPricesResponse struct {
	Snapshot MarketSnapshot `json:"snapshot"`
}
