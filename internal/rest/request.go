package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func createRequest[T any](
	client *http.Client,
	baseURL string,
	apiKey string,
	accountID string,
	accessToken string,
	version string,
	verb string,
	url string,
	body interface{},
) (*T, error) {
	url = fmt.Sprintf("%s/%s", baseURL, url)

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

	req.Header.Set("X-IG-API-KEY", apiKey)
	req.Header.Set("IG-ACCOUNT-ID", accountID)

	req.Header.Set("Version", version)

	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
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

type RequestBuilder[T any] struct {
	client      *http.Client
	baseURL     string
	apiKey      string
	accountID   string
	accessToken string
	version     string
	verb        string
	url         string
	params      []string
	body        interface{}
}

func NewRequest[T any](
	client *http.Client,
	baseURL string,
	apiKey string,
	accountID string,
	accessToken string,
	version string,
	verb string,
	url string,
) *RequestBuilder[T] {
	return &RequestBuilder[T]{
		client:      client,
		baseURL:     baseURL,
		apiKey:      apiKey,
		accountID:   accountID,
		accessToken: accessToken,
		version:     version,
		verb:        verb,
		url:         url,
	}
}

func (rb *RequestBuilder[T]) WithBody(body interface{}) *RequestBuilder[T] {
	rb.body = body
	return rb
}

func (rb *RequestBuilder[T]) WithParams(params ...string) *RequestBuilder[T] {
	rb.params = params
	return rb
}

func (r *RequestBuilder[T]) WithQueryParams(params map[string]string) *RequestBuilder[T] {
	if len(params) > 0 {
		var queryString string
		for key, value := range params {
			if queryString == "" {
				queryString = fmt.Sprintf("?%s=%s", key, value)
			} else {
				queryString += fmt.Sprintf("&%s=%s", key, value)
			}
		}
		r.url += queryString
	}
	return r
}

func (rb *RequestBuilder[T]) Execute() (*T, error) {
	var urlBuilder strings.Builder
	urlBuilder.WriteString(rb.url)
	for _, param := range rb.params {
		urlBuilder.WriteString("/")
		urlBuilder.WriteString(param)
	}
	fullURL := urlBuilder.String()

	var requestBody io.Reader
	if rb.body != nil {
		jsonBody, err := json.Marshal(rb.body)
		if err != nil {
			var empty T
			return &empty, err
		}
		requestBody = bytes.NewReader(jsonBody)
	}

	return createRequest[T](
		rb.client,
		rb.baseURL,
		rb.apiKey,
		rb.accountID,
		rb.accessToken,
		rb.version,
		rb.verb,
		fullURL,
		requestBody,
	)
}
