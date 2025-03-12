package client

type ClientSentimentResponse struct {
	ClientSentiments []ClientSentiment `json:"clientSentiments"`
}

type ClientSentiment struct {
	MarketID                string  `json:"marketId"`
	LongPositionPercentage  float64 `json:"longPositionPercentage"`
	ShortPositionPercentage float64 `json:"shortPositionPercentage"`
}
