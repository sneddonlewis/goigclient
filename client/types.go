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

type MarketSnapshot struct {
	Bid    float64 `json:"bid"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Status string  `json:"marketStatus"`
}

type MarketPricesResponse struct {
	Snapshot MarketSnapshot `json:"snapshot"`
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
