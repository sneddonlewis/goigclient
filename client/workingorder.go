package client

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
