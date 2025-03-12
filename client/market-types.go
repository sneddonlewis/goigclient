package client

type MarketNavigationResponse struct {
	Markets []MarketData `json:"markets"`
	Nodes   []MarketNode `json:"nodes"`
}

type MarketData struct {
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
	Instrument     Instrument     `json:"instrument"`
	DealingRules   DealingRules   `json:"dealingRules"`
	MarketSnapshot MarketSnapshot `json:"snapshot"`
}

type Instrument struct {
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
