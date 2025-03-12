package client

type Position struct {
	ContractSize         float64 `json:"contractSize "`
	ControlledRisk       bool    `json:"controlledRisk"`
	CreatedDate          string  `json:"createdDate"`
	CreatedDateUTC       string  `json:"createdDateUTC"`
	Currency             string  `json:"currency"`
	DealID               string  `json:"dealId"`
	DealReference        string  `json:"dealReference"`
	Direction            string  `json:"direction"`
	Level                float64 `json:"level"`
	LimitLevel           float64 `json:"limitLevel"`
	LimitedRiskPremium   float64 `json:"limitedRiskPremium"`
	Size                 float64 `json:"size"`
	StopLevel            float64 `json:"stopLevel"`
	TrailingStep         float64 `json:"trailingStep"`
	TrailingStopDistance float64 `json:"trailingStopDistance"`
}

type Market struct {
	Epic           string `json:"epic"`
	InstrumentName string `json:"instrumentName"`
}

type OpenPositionResponse struct {
	OpenPosition `json:"position"`
}

type OpenPosition struct {
	Position Position `json:"position"`
	Market   Market   `json:"market"`
}

type OpenPositionsResponse struct {
	Positions []OpenPosition `json:"positions"`
}

type CreatePositionRequest struct {
	CurrencyCode          string   `json:"currencyCode"`
	DealReference         *string  `json:"dealReference,omitempty"`
	Direction             string   `json:"direction"`
	Epic                  string   `json:"epic"`
	Expiry                string   `json:"expiry"`
	ForceOpen             bool     `json:"forceOpen"`
	GuaranteedStop        bool     `json:"guaranteedStop"`
	Level                 *float64 `json:"level,omitempty"`
	LimitDistance         *float64 `json:"limitDistance,omitempty"`
	LimitLevel            *float64 `json:"limitLevel,omitempty"`
	OrderType             string   `json:"orderType"`
	QuoteId               *string  `json:"quoteId,omitempty"`
	Size                  float64  `json:"size"`
	StopDistance          *float64 `json:"stopDistance,omitempty"`
	StopLevel             *float64 `json:"stopLevel,omitempty"`
	TimeInForce           *string  `json:"timeInForce,omitempty"`
	TrailingStop          *bool    `json:"trailingStop,omitempty"`
	TrailingStopIncrement *float64 `json:"trailingStopIncrement,omitempty"`
}

type ClosePositionRequest struct {
	DealID      *string  `json:"dealId,omitempty"`
	Epic        *string  `json:"epic,omitempty"`
	Expiry      *string  `json:"expiry,omitempty"`
	Direction   string   `json:"direction"`
	Level       *float64 `json:"level,omitempty"`
	OrderType   string   `json:"orderType"`
	QuoteID     *string  `json:"quoteId,omitempty"`
	Size        float64  `json:"size"`
	TimeInForce *string  `json:"timeInForce,omitempty"`
}

type UpdatePositionRequest struct {
	GuaranteedStop        bool     `json:"guaranteedStop"`
	LimitLevel            *float64 `json:"limitLevel,omitempty"`
	StopLevel             *float64 `json:"stopLevel,omitempty"`
	TrailingStop          *bool    `json:"trailingStop,omitempty"`
	TrailingStopDistance  *float64 `json:"trailingStopDistance,omitempty"`
	TrailingStopIncrement *float64 `json:"trailingStopIncrement,omitempty"`
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

type DealConfirmationResponse struct {
	AffectedDeals  []AffectedDeal `json:"affectedDeals"`
	Date           string         `json:"date"`
	DealID         string         `json:"dealId"`
	DealReference  string         `json:"dealReference"`
	DealStatus     string         `json:"dealStatus"`
	Direction      string         `json:"direction"`
	Epic           string         `json:"epic"`
	Expiry         string         `json:"expiry"`
	GuaranteedStop bool           `json:"guaranteedStop"`
	Level          float64        `json:"level"`
	LimitDistance  float64        `json:"limitDistance"`
	LimitLevel     float64        `json:"limitLevel"`
	Profit         float64        `json:"profit"`
	ProfitCurrency string         `json:"profitCurrency"`
	Reason         string         `json:"reason"`
	Size           float64        `json:"size"`
	Status         string         `json:"status"`
	StopDistance   float64        `json:"stopDistance"`
	StopLevel      float64        `json:"stopLevel"`
	TrailingStop   bool           `json:"trailingStop"`
}

type AffectedDeal struct {
	DealID string `json:"dealId"`
	Status string `json:"status"`
}
