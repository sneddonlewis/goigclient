package client

import (
	"net/http"

	"github.com/sneddonlewis/goigclient/internal/rest"
)

// ClientSentiment retrieves the client sentiment for given market identifiers.
//
// This method uses **version 1** of the IG API.
//
// Parameters:
//   - marketIds: A comma-separated list of market identifiers (optional).
//     [Constraint: Pattern(regexp="^(?>(?:[A-Za-z0-9_\\-\\ ]){1,30},?\\s?){0,500}$")]
//
// Returns a pointer to ClientSentimentResponse containing client sentiment data, or
// an error if the request fails.
func (c *IGClient) ClientSentiments(marketIds string) (*ClientSentimentResponse, error) {
	params := map[string]string{}
	if marketIds != "" {
		params["marketIds"] = marketIds
	}

	return rest.NewRequest[ClientSentimentResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodGet,
		"client-sentiment",
	).
		WithQueryParams(params).
		Execute()
}

// ClientSentiment retrieves the client sentiment for given market identifier.
//
// This method uses **version 1** of the IG API.
//
// Parameters:
//   - marketId: A market identifiers.
//
// Returns a pointer to ClientSentiment containing client sentiment data, or
// an error if the request fails.
func (c *IGClient) ClientSentiment(marketId string) (*ClientSentiment, error) {
	return rest.NewRequest[ClientSentiment](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodGet,
		"client-sentiment",
	).
		WithParams(marketId).
		Execute()
}

// RelatedClientSentiment retrieves a list of related market sentiment for the given instrument's market.
//
// This method uses **version 1** of the IG API.
//
// Parameters:
//   - marketId: A market identifiers.
//
// Returns a pointer to ClientSentimentResponse, or
// an error if the request fails.
func (c *IGClient) RelatedClientSentiment(marketId string) (*ClientSentimentResponse, error) {
	return rest.NewRequest[ClientSentimentResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodGet,
		"client-sentiment/related",
	).
		WithParams(marketId).
		Execute()
}
