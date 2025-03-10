package client

func (c *IGClient) AllOpenPositions() (*OpenPositionsResponse, error) {
	return getRequest[OpenPositionsResponse](c, v2, "positions")
}

func (c *IGClient) OneOpenPosition(dealID string) (*OpenPositionResponse, error) {
	return getRequestWithParams[OpenPositionResponse](c, v2, "positions", dealID)
}

func (c *IGClient) CreateOTCPosition() error {
	return nil
}

func (c *IGClient) ChangeOTCPosition() error {
	return nil
}

func (c *IGClient) CloseOTCPosition() error {
	return nil
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
