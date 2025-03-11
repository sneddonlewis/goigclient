package client

import (
	"net/http"

	"github.com/sneddonlewis/goigclient/internal/rest"
)

func (c *IGClient) AllWorkingOrders() (*WorkingOrdersResponse, error) {
	return rest.NewRequest[WorkingOrdersResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v2,
		http.MethodGet,
		"workingorders",
	).Execute()
}

func (c *IGClient) CreateOTCWorkingOrder(request *CreateWorkingOrderRequest) (*DealResponse, error) {
	return rest.NewRequest[DealResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v2,
		http.MethodPost,
		"workingorders/otc",
	).
		WithBody(request).
		Execute()
}

func (c *IGClient) ChangeOTCWorkingOrder(
	request *UpdateWorkingOrderRequest,
	dealID string,
) (*DealResponse, error) {
	return rest.NewRequest[DealResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v2,
		http.MethodPut,
		"workingorders/otc",
	).
		WithBody(request).
		WithParams(dealID).
		Execute()
}

func (c *IGClient) CloseOTCWorkingOrder(dealID string) (*DealResponse, error) {
	return rest.NewRequest[DealResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v2,
		http.MethodDelete,
		"workingorders/otc",
	).
		WithParams(dealID).
		Execute()
}

type WorkingOrdersResponse struct {
	WorkingOrders []WorkingOrder `json:"workingOrders"`
}

type WorkingOrder struct {
	MarketData       MarketData       `json:"marketData"`
	WorkingOrderData WorkingOrderData `json:"workingOrderData"`
}

type MarketData struct {
	Bid                      float64 `json:"bid"`
	DelayTime                int     `json:"delayTime"`
	Epic                     string  `json:"epic"`
	ExchangeId               string  `json:"exchangeId"`
	Expiry                   string  `json:"expiry"`
	High                     float64 `json:"high"`
	InstrumentName           string  `json:"instrumentName"`
	InstrumentType           string  `json:"instrumentType"`
	LotSize                  float64 `json:"lotSize"`
	Low                      float64 `json:"low"`
	MarketStatus             string  `json:"marketStatus"`
	NetChange                float64 `json:"netChange"`
	Offer                    float64 `json:"offer"`
	PercentageChange         float64 `json:"percentageChange"`
	ScalingFactor            float64 `json:"scalingFactor"`
	StreamingPricesAvailable bool    `json:"streamingPricesAvailable"`
	UpdateTime               string  `json:"updateTime"`
	UpdateTimeUTC            string  `json:"updateTimeUTC"`
}

type WorkingOrderData struct {
	CreatedDate        string  `json:"createdDate"`
	CreatedDateUTC     string  `json:"createdDateUTC"`
	CurrencyCode       string  `json:"currencyCode"`
	DealId             string  `json:"dealId"`
	Direction          string  `json:"direction"`
	DMA                bool    `json:"dma"`
	Epic               string  `json:"epic"`
	GoodTillDate       string  `json:"goodTillDate"`
	GoodTillDateISO    string  `json:"goodTillDateISO"`
	GuaranteedStop     bool    `json:"guaranteedStop"`
	LimitDistance      float64 `json:"limitDistance"`
	LimitedRiskPremium float64 `json:"limitedRiskPremium"`
	OrderLevel         float64 `json:"orderLevel"`
	OrderSize          float64 `json:"orderSize"`
	OrderType          string  `json:"orderType"`
	StopDistance       float64 `json:"stopDistance"`
	TimeInForce        string  `json:"timeInForce"`
}

type CreateWorkingOrderRequest struct {
	CurrencyCode   string   `json:"currencyCode"`
	DealReference  *string  `json:"dealReference,omitempty"`
	Direction      string   `json:"direction"`
	Epic           string   `json:"epic"`
	Expiry         string   `json:"expiry"`
	ForceOpen      bool     `json:"forceOpen"`
	GoodTillDate   *string  `json:"goodTillDate,omitempty"`
	GuaranteedStop bool     `json:"guaranteedStop"`
	Level          float64  `json:"level"`
	LimitDistance  *float64 `json:"limitDistance,omitempty"`
	LimitLevel     *float64 `json:"limitLevel,omitempty"`
	Size           float64  `json:"size"`
	StopDistance   *float64 `json:"stopDistance,omitempty"`
	StopLevel      *float64 `json:"stopLevel,omitempty"`
	TimeInForce    string   `json:"timeInForce"`
	Type           string   `json:"type"`
}

type UpdateWorkingOrderRequest struct {
	GoodTillDate   *string  `json:"goodTillDate,omitempty"`
	GuaranteedStop bool     `json:"guaranteedStop"`
	Level          float64  `json:"level"`
	LimitDistance  *float64 `json:"limitDistance,omitempty"`
	LimitLevel     *float64 `json:"limitLevel,omitempty"`
	StopDistance   *float64 `json:"stopDistance,omitempty"`
	StopLevel      *float64 `json:"stopLevel,omitempty"`
	TimeInForce    string   `json:"timeInForce"`
	Type           string   `json:"type"`
}
