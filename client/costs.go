package client

import (
	"fmt"
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

// CostsPDF retrieves a previously generated indicative costs and charges quote as a PDF.
//
// This method uses **version 1** of the IG API.
//
// Parameters:
//   - indicativeQuoteReference: The reference ID for the indicative costs and charges quote.
//
// Returns a byte slice containing the PDF data, or
// an error if the request fails.
func (c *IGClient) CostsPDF(indicativeQuoteReference string) (*[]byte, error) {
	return rest.NewRequest[[]byte](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodGet,
		"indicativecostsandcharges/durablemedium",
	).
		WithParams(indicativeQuoteReference).
		Execute()
}

// EditingCosts returns indicative costs and charges for editing an order.
//
// This method uses **version 1** of the IG API.
//
// Parameters:
//   - request: The request model containing details of the order to be edited.
//
// Returns a pointer to EditingCostsResponse containing indicative costs and charges, or
// an error if the request fails.
func (c *IGClient) EditingCosts(request EditingCostsRequest) (*EditingCostsResponse, error) {
	return rest.NewRequest[EditingCostsResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodPost,
		"indicativecostsandcharges/edit",
	).
		WithBody(request).
		Execute()
}

// CostHistory retrieves indicative costs and charges history for the given date range.
//
// This method uses **version 1** of the IG API.
//
// Parameters:
//   - from: Start date in format (yyyy-MM-dd'T'HH:mm:ss.SSSXXX).
//   - to: End date in format (yyyy-MM-dd'T'HH:mm:ss.SSSXXX).
//   - pageSize: Number of results per page (optional).
//   - pageNumber: Page number (optional).
//   - type: Type filter (optional).
//
// Returns a pointer to CostHistoryResponse containing the historical costs and charges data, or
// an error if the request fails.
func (c *IGClient) CostHistory(
	from,
	to string,
	pageSize,
	pageNumber *int,
	costType string,
) (*CostHistoryResponse, error) {
	params := make(map[string]string)
	if pageSize != nil {
		params["pageSize"] = fmt.Sprintf("%d", *pageSize)
	}
	if pageNumber != nil {
		params["pageNumber"] = fmt.Sprintf("%d", *pageNumber)
	}
	if costType != "" {
		params["type"] = costType
	}

	// non-standard as has variables in between fixed "/to/"
	requestURL := fmt.Sprintf("indicativecostsandcharges/history/from/%s/to/%s", from, to)

	return rest.NewRequest[CostHistoryResponse](
		c.HTTPClient,
		c.BaseURL,
		c.APIKey,
		c.AccountID,
		c.AccessToken,
		v1,
		http.MethodGet,
		requestURL,
	).
		WithQueryParams(params).
		Execute()
}
