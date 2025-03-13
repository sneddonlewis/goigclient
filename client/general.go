package client

import (
	"net/http"

	"github.com/sneddonlewis/goigclient/internal/rest"
)

// Applications retrieves a list of applications associated with the account.
//
// This method uses version 1 of the IG API.
//
// Returns a pointer to ApplicationsResponse containing the list of applications, or
// an error if the request fails.
func (c *IGClient) Applications() (*ApplicationsResponse, error) {
	return rest.NewRequest[ApplicationsResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodGet,
		"operations/applications",
	).
		Execute()
}

// UpdateApplication alters the details of a given user application.
//
// This method uses version 1 of the IG API.
//
// Parameters:
// - request: UpdateApplicationRequest containing the new application settings.
//
// Returns a pointer to UpdateApplicationResponse containing the updated application details,
// or an error if the request fails.
func (c *IGClient) UpdateApplication(request UpdateApplicationRequest) (*UpdateApplicationResponse, error) {
	return rest.NewRequest[UpdateApplicationResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodPut,
		"operations/applications",
	).
		WithBody(request).
		Execute()
}

// DisableApplication disables the current application key from processing further requests.
//
// This method uses version 1 of the IG API.
//
// Returns a pointer to DisableApplicationResponse containing the updated application details,
// or an error if the request fails.
func (c *IGClient) DisableApplication() (*DisableApplicationResponse, error) {
	return rest.NewRequest[DisableApplicationResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodPut,
		"operations/application/disable",
	).
		Execute()
}
