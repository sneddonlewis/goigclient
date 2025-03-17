package client

import (
	"net/http"

	"github.com/sneddonlewis/goigclient/internal/rest"
)

func (c *IGClient) Positions() (*OpenPositionsResponse, error) {
	return rest.NewRequest[OpenPositionsResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v2,
		http.MethodGet,
		"positions",
	).Execute()
}

func (c *IGClient) Position(dealID string) (*OpenPositionResponse, error) {
	return rest.NewRequest[OpenPositionResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v2,
		http.MethodGet,
		"positions",
	).
		WithParams(dealID).
		Execute()
}

func (c *IGClient) CreateOTCPosition(request *CreatePositionRequest) (*DealResponse, error) {
	return rest.NewRequest[DealResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v2,
		http.MethodPost,
		"positions/otc",
	).
		WithBody(request).
		Execute()
}

func (c *IGClient) ChangeOTCPosition(request *UpdatePositionRequest) (*DealResponse, error) {
	return rest.NewRequest[DealResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v2,
		http.MethodPut,
		"positions/otc",
	).
		WithBody(request).
		Execute()
}

func (c *IGClient) CloseOTCPosition(request *ClosePositionRequest) (*DealResponse, error) {
	return rest.NewRequest[DealResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v2,
		http.MethodDelete,
		"positions/otc",
	).
		WithBody(request).
		Execute()
}

func (c *IGClient) WorkingOrders() (*WorkingOrdersResponse, error) {
	return rest.NewRequest[WorkingOrdersResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v2,
		http.MethodGet,
		"workingorders",
	).Execute()
}

func (c *IGClient) CreateOTCWorkingOrder(request *CreateWorkingOrderRequest) (*DealResponse, error) {
	return rest.NewRequest[DealResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v2,
		http.MethodPost,
		"workingorders/otc",
	).
		WithBody(request).
		Execute()
}

func (c *IGClient) ChangeOTCWorkingOrder(
	request *UpdateWorkingOrderRequest,
	dealID string,
) (*DealResponse, error) {
	return rest.NewRequest[DealResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v2,
		http.MethodPut,
		"workingorders/otc",
	).
		WithBody(request).
		WithParams(dealID).
		Execute()
}

func (c *IGClient) CloseOTCWorkingOrder(dealID string) (*DealResponse, error) {
	return rest.NewRequest[DealResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v2,
		http.MethodDelete,
		"workingorders/otc",
	).
		WithParams(dealID).
		Execute()
}

// DealConfirmation retrieves a deal confirmation for the given deal reference.
//
// This method uses **version 1** of the IG API.
//
// Parameters:
//   - dealReference: The reference of the deal to confirm.
//
// Returns a pointer to DealConfirmationResponse containing deal details, or
// an error if the request fails.
func (c *IGClient) DealConfirmation(dealReference string) (*DealConfirmationResponse, error) {

	return rest.NewRequest[DealConfirmationResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodGet,
		"confirms",
	).
		WithParams(dealReference).
		Execute()
}
