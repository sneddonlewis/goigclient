package client

import (
	"net/http"

	"github.com/sneddonlewis/goigclient/internal/rest"
)

// MarketNavigation retrieves all top-level market categories in the market navigation hierarchy.
//
// This method uses **version 1** of the IG API.
//
// Returns a pointer to MarketNavigationResponse containing market hierarchy nodes and market data, or
// an error if the request fails.
func (c *IGClient) MarketNavigation() (*MarketNavigationResponse, error) {
	return rest.NewRequest[MarketNavigationResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodGet,
		"market-navigation",
	).
		Execute()
}

// NavigateMarketNode retrieves all sub-nodes of the given node in the market navigation hierarchy.
//
// This method uses **version 1** of the IG API.
//
// Parameters:
//   - nodeId: The identifier of the node to browse.
//
// Returns a pointer to MarketNavigationResponse containing market hierarchy nodes and market data, or
// an error if the request fails.
func (c *IGClient) NavigateMarketNode(nodeId string) (*MarketNavigationResponse, error) {

	return rest.NewRequest[MarketNavigationResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodGet,
		"market-navigation",
	).
		WithParams(nodeId).
		Execute()
}

// Markets retrieves the details of the given markets.
//
// This method uses **version 2** of the IG API.
//
// Parameters:
//   - epics: A comma-separated list of market epics to retrieve (max: 50).
//   - filter: Optional filter for market details (default: ALL).
//
// Returns a pointer to MarketsResponse containing detailed market information, or
// an error if the request fails.
func (c *IGClient) Markets(epics string, filter *string) (*MarketsResponse, error) {
	params := map[string]string{"epics": epics}
	if filter != nil {
		params["filter"] = *filter
	}

	return rest.NewRequest[MarketsResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v2,
		http.MethodGet,
		"markets",
	).
		WithQueryParams(params).
		Execute()
}

// Market retrieves the details of a single market based on the provided epic.
//
// This method uses **version 3** of the IG API.
//
// Parameters:
//   - epic: The unique identifier of the market.
//
// Returns:
//   - A pointer to MarketResponse containing detailed market information.
//   - An error if the request fails.
func (c *IGClient) Market(epic string) (*MarketResponse, error) {
	return rest.NewRequest[MarketResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v3,
		http.MethodGet,
		"markets",
	).
		WithParams(epic).
		Execute()
}
