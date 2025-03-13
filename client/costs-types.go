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

type OpeningCostsRequest struct {
	Ask              float64  `json:"ask"`
	Bid              float64  `json:"bid"`
	DealCurrencyCode string   `json:"dealCurrencyCode"`
	DealReference    string   `json:"dealReference"`
	Direction        string   `json:"direction"`
	Epic             *string  `json:"epic,omitempty"`
	GuaranteedStop   *bool    `json:"guaranteedStop,omitempty"`
	InstrumentID     *string  `json:"instrumentId,omitempty"`
	KnockoutPremium  *float64 `json:"knockoutPremium,omitempty"`
	PriceLevel       *float64 `json:"priceLevel,omitempty"`
	Size             float64  `json:"size"`
	StopLevel        *float64 `json:"stopLevel,omitempty"`
}

type OpeningCostsResponse struct {
	BorrowingCharge             *float64 `json:"borrowingCharge,omitempty"`
	ClosingCommission           *float64 `json:"closingCommission,omitempty"`
	ClosingFxFee                *float64 `json:"closingFxFee,omitempty"`
	ClosingIFTT                 *float64 `json:"closingIFTT,omitempty"`
	ClosingSpread               *float64 `json:"closingSpread,omitempty"`
	CurrencyCodeISO             string   `json:"currencyCodeISO"`
	DailyRunningFxFee           *float64 `json:"dailyRunningFxFee,omitempty"`
	EtpEntryCost                *float64 `json:"etpEntryCost,omitempty"`
	EtpExitCost                 *float64 `json:"etpExitCost,omitempty"`
	EtpOngoingCost              *float64 `json:"etpOngoingCost,omitempty"`
	GuaranteedStopDeposit       *float64 `json:"guaranteedStopDeposit,omitempty"`
	GuaranteedStopReturn        *float64 `json:"guaranteedStopReturn,omitempty"`
	IndicativeQuoteReference    string   `json:"indicativeQuoteReference"`
	Inducements                 *float64 `json:"inducements,omitempty"`
	KnockoutPremiumDeposit      *float64 `json:"knockoutPremiumDeposit,omitempty"`
	KnockoutPremiumReturn       *float64 `json:"knockoutPremiumReturn,omitempty"`
	NotionalValue               float64  `json:"notionalValue"`
	NotionalValueInUserCurrency float64  `json:"notionalValueInUserCurrency"`
	OpeningCommission           *float64 `json:"openingCommission,omitempty"`
	OpeningFxFee                *float64 `json:"openingFxFee,omitempty"`
	OpeningIFTT                 *float64 `json:"openingIFTT,omitempty"`
	OpeningSpread               *float64 `json:"openingSpread,omitempty"`
	OvernightFundingFee         *float64 `json:"overnightFundingFee,omitempty"`
}

type EditingCostsRequest struct {
	Ask              float64  `json:"ask"`
	Bid              float64  `json:"bid"`
	DealCurrencyCode string   `json:"dealCurrencyCode"`
	DealReference    string   `json:"dealReference"`
	Direction        string   `json:"direction"`
	EditType         string   `json:"editType,omitempty"`
	Epic             string   `json:"epic,omitempty"`
	GuaranteedStop   *bool    `json:"guaranteedStop,omitempty"`
	InstrumentID     string   `json:"instrumentId,omitempty"`
	KnockoutPremium  *float64 `json:"knockoutPremium,omitempty"`
	LimitLevel       *float64 `json:"limitLevel,omitempty"`
	OpeningLevel     float64  `json:"openingLevel"`
	PriceLevel       *float64 `json:"priceLevel,omitempty"`
	Size             float64  `json:"size"`
	StopLevel        *float64 `json:"stopLevel,omitempty"`
}

type EditingCostsResponse struct {
	CurrencyCodeISO string                `json:"currencyCodeISO"`
	Limit           ClosingExAnteResponse `json:"limit"`
	Stop            ClosingExAnteResponse `json:"stop"`
}

type ClosingExAnteResponse struct {
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
}

type CostHistoryResponse struct {
	CostsAndChargesHistory []CostHistoryItem `json:"costsAndChargesHistory"`
	Pagination             Pagination        `json:"pagination"`
}

type CostHistoryItem struct {
	CreatedTimestamp         string `json:"createdTimestamp"`
	Direction                string `json:"direction"`
	IndicativeQuoteReference string `json:"indicativeQuoteReference"`
	InstrumentName           string `json:"instrumentName"`
	Type                     string `json:"type"`
}

type Pagination struct {
	PageNumber    int `json:"pageNumber"`
	PageSize      int `json:"pageSize"`
	TotalElements int `json:"totalElements"`
	TotalPages    int `json:"totalPages"`
}
