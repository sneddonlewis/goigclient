package client

import (
	"net/http"

	"github.com/sneddonlewis/goigclient/internal/rest"
)

// ClosingCosts retrieves indicative costs and charges at closing.
//
// This method uses **version 1** of the IG API.
//
// Parameters:
//   - request: ClosingCostsRequest containing details for calculating indicative closing costs.
//
// Returns a pointer to ClosingCostsResponse containing the estimated costs, or
// an error if the request fails.
func (c *IGClient) ClosingCosts(request ClosingCostsRequest) (*ClosingCostsResponse, error) {
	return rest.NewRequest[ClosingCostsResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodPost,
		"indicativecostsandcharges/close",
	).
		WithBody(request).
		Execute()
}
