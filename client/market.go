package client

import (
	"net/http"

	"github.com/sneddonlewis/goigclient/internal/rest"
)

func (c *IGClient) OneMarket(epic string) (*MarketPricesResponse, error) {
	return rest.NewRequest[MarketPricesResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v3,
		http.MethodGet,
		"markets",
	).
		WithParams(epic).
		Execute()
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
