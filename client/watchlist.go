package client

import (
	"net/http"

	"github.com/sneddonlewis/goigclient/internal/rest"
)

const watchlists = "watchlists"

// Watchlists retrieves all watchlists belonging to the active account.
//
// This method uses **version 1** of the IG API.
//
// Returns a pointer to WatchlistsResponse containing all watchlists, or
// an error if the request fails.
func (c *IGClient) Watchlists() (*Watchlists, error) {
	return rest.NewRequest[Watchlists](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodGet,
		watchlists,
	).
		Execute()
}

// Watchlist retrieves the details of a specific watchlist by its identifier.
//
// This method uses **version 1** of the IG API.
//
// Parameters:
//   - watchlistId: The ID of the watchlist to retrieve.
//
// Returns a pointer to WatchlistResponse containing the watchlist details, or
// an error if the request fails.
func (c *IGClient) Watchlist(watchlistId string) (*WatchlistContent, error) {
	return rest.NewRequest[WatchlistContent](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodGet,
		watchlists,
	).
		WithParams(watchlistId).
		Execute()
}

// CreateWatchlist creates a new watchlist with the specified name and associated market epics.
//
// This method uses **version 1** of the IG API.
//
// Parameters:
//   - request: CreateWatchlistRequest containing the watchlist name and associated epics.
//
// Returns a pointer to CreateWatchlistResponse containing the result of the operation, or
// an error if the request fails.
func (c *IGClient) CreateWatchlist(request CreateWatchlistRequest) (*CreateWatchlistResponse, error) {
	return rest.NewRequest[CreateWatchlistResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodPost,
		watchlists,
	).
		WithBody(request).
		Execute()
}

// DeleteWatchlist deletes a watchlist by its identifier.
//
// This method uses **version 1** of the IG API.
//
// Parameters:
//   - watchlistId: The ID of the watchlist to be deleted.
//
// Returns a pointer to DeleteWatchlistResponse indicating the result of the operation, or
// an error if the request fails.
func (c *IGClient) DeleteWatchlist(watchlistId string) (*OperationResponse, error) {
	return rest.NewRequest[OperationResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodDelete,
		watchlists,
	).
		WithParams(watchlistId).
		Execute()
}

// WatchlistAddMarket adds a market to a watchlist by its identifier.
//
// This method uses **version 1** of the IG API.
//
// Parameters:
//   - watchlistId: The ID of the watchlist to update.
//   - request: WatchlistAddMarketRequest containing the market epic to add.
//
// Returns a pointer to WatchlistAddMarketResponse indicating the result of the operation, or
// an error if the request fails.
func (c *IGClient) WatchlistAddMarket(watchlistId, epic string) (*OperationResponse, error) {
	request := WatchlistAddMarketRequest{
		epic,
	}
	return rest.NewRequest[OperationResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodPut,
		watchlists,
	).
		WithBody(request).
		WithParams(watchlistId).
		Execute()
}

// WatchlistRemoveMarket removes a market from a watchlist by its identifier.
//
// This method uses **version 1** of the IG API.
//
// Parameters:
//   - watchlistId: The ID of the watchlist to update.
//   - epic: The market epic to remove from the watchlist.
//
// Returns a pointer to WatchlistRemoveMarketResponse indicating the result of the operation, or
// an error if the request fails.
func (c *IGClient) WatchlistRemoveMarket(watchlistId, epic string) (*OperationResponse, error) {
	return rest.NewRequest[OperationResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodDelete,
		watchlists,
	).
		WithParams(watchlistId, epic).
		Execute()
}
