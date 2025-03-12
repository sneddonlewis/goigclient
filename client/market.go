package client

import (
	"fmt"
	"net/http"
	"time"

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
// Returns a pointer to MarketResponse containing detailed market information, or
// an error if the request fails.
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

// SearchMarkets retrieves all markets matching the given search term.
//
// This method uses **version 1** of the IG API.
//
// Parameters:
//   - searchTerm: The term to be used in the search (optional).
//
// Returns:
//   - A pointer to SearchMarketsResponse containing matching market data.
//   - An error if the request fails.
func (c *IGClient) SearchMarkets(searchTerm *string) (*SearchMarketsResponse, error) {
	params := map[string]string{}
	if searchTerm != nil {
		params["searchTerm"] = *searchTerm
	}

	return rest.NewRequest[SearchMarketsResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodGet,
		"markets",
	).
		WithQueryParams(params).
		Execute()
}

// Prices retrieves historical prices for a given instrument epic.
//
// This method uses **version 3** of the IG API.
//
// Parameters:
//   - epic: The instrument epic.
//   - resolution: The price resolution (optional, default: MINUTE).
//   - from: Start date-time in ISO format (optional).
//   - to: End date-time in ISO format (optional).
//   - max: Limits the number of price points (optional, default: 10).
//   - pageSize: Number of results per page (optional, default: 20, set 0 to disable paging).
//   - pageNumber: Page number (optional, default: 1).
//
// Returns a pointer to PricesResponse containing historical price data, or
// an error if the request fails.
func (c *IGClient) Prices(
	epic string,
	resolution *string,
	from, to *time.Time,
	max, pageSize, pageNumber *int,
) (*PricesResponse, error) {
	params := map[string]string{}
	if resolution != nil {
		params["resolution"] = *resolution
	}
	if from != nil {
		params["from"] = from.Format(time.RFC3339)
	}
	if to != nil {
		params["to"] = to.Format(time.RFC3339)
	}
	if max != nil {
		params["max"] = fmt.Sprintf("%d", *max)
	}
	if pageSize != nil {
		params["pageSize"] = fmt.Sprintf("%d", *pageSize)
	}
	if pageNumber != nil {
		params["pageNumber"] = fmt.Sprintf("%d", *pageNumber)
	}

	return rest.NewRequest[PricesResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v3,
		http.MethodGet,
		"prices",
	).
		WithParams(epic).
		WithQueryParams(params).
		Execute()
}
