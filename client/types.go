package client

type DealResponse struct {
	DealReference string `json:"dealReference"`
}

type Metadata struct {
	Paging Paging `json:"paging"`
}

type OperationResponse struct {
	Status string `json:"status"`
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
