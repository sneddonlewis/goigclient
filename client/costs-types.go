package client

type ClosingCostsRequest struct {
	Ask              float64 `json:"ask"`
	Bid              float64 `json:"bid"`
	DealCurrencyCode string  `json:"dealCurrencyCode"`
	DealReference    string  `json:"dealReference"`
	Direction        string  `json:"direction"`
	Epic             string  `json:"epic,omitempty"`
	GuaranteedStop   bool    `json:"guaranteedStop,omitempty"`
	InstrumentId     string  `json:"instrumentId,omitempty"`
	KnockoutPremium  float64 `json:"knockoutPremium,omitempty"`
	OpeningLevel     float64 `json:"openingLevel"`
	PriceLevel       float64 `json:"priceLevel,omitempty"`
	Size             float64 `json:"size"`
	StopLevel        float64 `json:"stopLevel,omitempty"`
}

type ClosingCostsResponse struct {
	ClosingExAnteResponse struct {
		ClosingCommission           float64 `json:"closingCommission"`
		ClosingFxFee                float64 `json:"closingFxFee"`
		ClosingIFTT                 float64 `json:"closingIFTT"`
		ClosingSpread               float64 `json:"closingSpread"`
		ETPExitCost                 float64 `json:"etpExitCost"`
		GuaranteedStopReturn        float64 `json:"guaranteedStopReturn"`
		IndicativeQuoteReference    string  `json:"indicativeQuoteReference"`
		KnockoutPremiumReturn       float64 `json:"knockoutPremiumReturn"`
		NotionalValue               float64 `json:"notionalValue"`
		NotionalValueInUserCurrency float64 `json:"notionalValueInUserCurrency"`
		CurrencyCodeISO             string  `json:"currencyCodeISO"`
	} `json:"close"`
}
