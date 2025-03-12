package client

type MarketNavigationResponse struct {
	Markets []MarketDataDetailed `json:"markets"`
	Nodes   []MarketNode         `json:"nodes"`
}

type MarketDataDetailed struct {
	Bid                      float64 `json:"bid"`
	DelayTime                int     `json:"delayTime"`
	Epic                     string  `json:"epic"`
	Expiry                   string  `json:"expiry"`
	High                     float64 `json:"high"`
	InstrumentName           string  `json:"instrumentName"`
	InstrumentType           string  `json:"instrumentType"`
	LotSize                  float64 `json:"lotSize"`
	Low                      float64 `json:"low"`
	MarketStatus             string  `json:"marketStatus"`
	NetChange                float64 `json:"netChange"`
	Offer                    float64 `json:"offer"`
	OtcTradeable             bool    `json:"otcTradeable"`
	PercentageChange         float64 `json:"percentageChange"`
	ScalingFactor            float64 `json:"scalingFactor"`
	StreamingPricesAvailable bool    `json:"streamingPricesAvailable"`
	UpdateTime               string  `json:"updateTime"`
	UpdateTimeUTC            string  `json:"updateTimeUTC"`
}

type MarketNode struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type MarketsResponse struct {
	MarketDetails []MarketDetail `json:"marketDetails"`
}

type MarketDetail struct {
	Instrument     Instrument     `json:"instrument"`
	DealingRules   DealingRules   `json:"dealingRules"`
	MarketSnapshot MarketSnapshot `json:"snapshot"`
}

type InstrumentDetailed struct {
	ChartCode                string        `json:"chartCode"`
	ContractSize             string        `json:"contractSize"`
	ControlledRiskAllowed    bool          `json:"controlledRiskAllowed"`
	Country                  string        `json:"country"`
	Currencies               []Currency    `json:"currencies"`
	Epic                     string        `json:"epic"`
	Expiry                   string        `json:"expiry"`
	ExpiryDetails            ExpiryDetails `json:"expiryDetails"`
	ForceOpenAllowed         bool          `json:"forceOpenAllowed"`
	LotSize                  float64       `json:"lotSize"`
	MarketID                 string        `json:"marketId"`
	Name                     string        `json:"name"`
	NewsCode                 string        `json:"newsCode"`
	OpeningHours             OpeningHours  `json:"openingHours"`
	StopsLimitsAllowed       bool          `json:"stopsLimitsAllowed"`
	StreamingPricesAvailable bool          `json:"streamingPricesAvailable"`
}

type Currency struct {
	BaseExchangeRate float64 `json:"baseExchangeRate"`
	Code             string  `json:"code"`
	ExchangeRate     float64 `json:"exchangeRate"`
	IsDefault        bool    `json:"isDefault"`
	Symbol           string  `json:"symbol"`
}

type ExpiryDetails struct {
	LastDealingDate string `json:"lastDealingDate"`
	SettlementInfo  string `json:"settlementInfo"`
}

type OpeningHours struct {
	MarketTimes []MarketTime `json:"marketTimes"`
}

type MarketTime struct {
	OpenTime  string `json:"openTime"`
	CloseTime string `json:"closeTime"`
}

type Instrument struct {
	Epic                     string  `json:"epic"`
	InstrumentName           string  `json:"name"`
	Expiry                   string  `json:"expiry"`
	LotSize                  float64 `json:"lotSize"`
	Type                     string  `json:"type"`
	ControlledRiskAllowed    bool    `json:"controlledRiskAllowed"`
	StreamingPricesAvailable bool    `json:"streamingPricesAvailable"`
}

type DealingRules struct {
	MinDealSize                   Rule `json:"minDealSize"`
	MaxStopOrLimitDistance        Rule `json:"maxStopOrLimitDistance"`
	MinControlledRiskStopDistance Rule `json:"minControlledRiskStopDistance"`
	MinNormalStopOrLimitDistance  Rule `json:"minNormalStopOrLimitDistance"`
	MinStepDistance               Rule `json:"minStepDistance"`
}

type Rule struct {
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
}

type MarketSnapshot struct {
	Bid              float64 `json:"bid"`
	Offer            float64 `json:"offer"`
	High             float64 `json:"high"`
	Low              float64 `json:"low"`
	NetChange        float64 `json:"netChange"`
	PercentageChange float64 `json:"percentageChange"`
	MarketStatus     string  `json:"marketStatus"`
	UpdateTime       string  `json:"updateTime"`
}

type MarketResponse struct {
	Instrument     InstrumentDetailed `json:"instrument"`
	DealingRules   DealingRules       `json:"dealingRules"`
	MarketSnapshot MarketSnapshot     `json:"snapshot"`
}

type SearchMarketsResponse struct {
	Markets []MarketData `json:"markets"`
}

type MarketData struct {
	Bid                      float64 `json:"bid"`
	DelayTime                int     `json:"delayTime"`
	Epic                     string  `json:"epic"`
	Expiry                   string  `json:"expiry"`
	High                     float64 `json:"high"`
	InstrumentName           string  `json:"instrumentName"`
	InstrumentType           string  `json:"instrumentType"`
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

type PricesResponse struct {
	InstrumentType string       `json:"instrumentType"`
	Metadata       Metadata     `json:"metadata"`
	Allowance      Allowance    `json:"allowance"`
	Prices         []PriceEntry `json:"prices"`
}

type PageData struct {
	PageNumber int `json:"pageNumber"`
	PageSize   int `json:"pageSize"`
	TotalPages int `json:"totalPages"`
	Size       int `json:"size"`
}

type Allowance struct {
	AllowanceExpiry    int `json:"allowanceExpiry"`
	RemainingAllowance int `json:"remainingAllowance"`
	TotalAllowance     int `json:"totalAllowance"`
}

type PriceEntry struct {
	ClosePrice       Price  `json:"closePrice"`
	HighPrice        Price  `json:"highPrice"`
	LowPrice         Price  `json:"lowPrice"`
	OpenPrice        Price  `json:"openPrice"`
	LastTradedVolume *int   `json:"lastTradedVolume"`
	SnapshotTime     string `json:"snapshotTime"`
	SnapshotTimeUTC  string `json:"snapshotTimeUTC"`
}

type Price struct {
	Ask        float64  `json:"ask"`
	Bid        float64  `json:"bid"`
	LastTraded *float64 `json:"lastTraded"`
}
