package client

func (c *IGClient) AllOpenPositions() (*OpenPositionsResponse, error) {
	return getRequest[OpenPositionsResponse](c, v2, "positions")
}

func (c *IGClient) OneOpenPosition(dealID string) (*OpenPositionResponse, error) {
	return getRequestWithParams[OpenPositionResponse](c, v2, "positions", dealID)
}

func (c *IGClient) CreateOTCPosition(request *CreatePositionRequest) (*AmendPositionResponse, error) {
	return nil, nil
}

func (c *IGClient) ChangeOTCPosition(request *UpdatePositionRequest) (*AmendPositionResponse, error) {
	return nil, nil
}

func (c *IGClient) CloseOTCPosition(request *ClosePositionRequest) (*AmendPositionResponse, error) {
	return nil, nil
}

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

type AmendPositionResponse struct {
	DealReference string `json:"dealReference"`
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
