package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func createRequest[T any](c *IGClient, version, verb, url string, body interface{}) (*T, error) {
	url = fmt.Sprintf("%s/%s", c.BaseURL, url)

	var response T

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return &response, err
	}
	requestBody := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(verb, url, requestBody)
	if err != nil {
		return &response, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	req.Header.Set("X-IG-API-KEY", c.APIKey)
	req.Header.Set("IG-ACCOUNT-ID", c.AccountID)

	req.Header.Set("Version", version)

	req.Header.Set("Authorization", "Bearer "+c.AccessToken)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return &response, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return &response, fmt.Errorf("%d: %s request failed for endpoint %s: %s",
			resp.StatusCode,
			verb,
			url,
			string(body),
		)
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return &response, err
	}

	return &response, nil
}

func getRequest[T any](c *IGClient, version, url string) (*T, error) {
	return createRequest[T](
		c,
		version,
		http.MethodGet,
		url,
		nil,
	)
}

func getRequestWithParams[T any](c *IGClient, version, url string, params ...string) (*T, error) {
	var urlBuilder strings.Builder

	urlBuilder.WriteString(url)
	for _, param := range params {
		urlBuilder.WriteString("/")
		urlBuilder.WriteString(param)
	}
	return createRequest[T](
		c,
		version,
		http.MethodGet,
		urlBuilder.String(),
		nil,
	)
}
