package client

import (
	"net/http"
)

const (
	v1 = "1"
	v2 = "2"
	v3 = "3"
)

type IGClient struct {
	APIKey       string
	Username     string
	Password     string
	BaseURL      string
	HTTPClient   *http.Client
	AccessToken  string
	RefreshToken string
	AccountID    string
}

func NewIGClient(apiKey, username, password string, isLive bool) *IGClient {
	var baseUrl string
	if isLive {
		baseUrl = "https://api.ig.com/gateway/deal"
	} else {
		baseUrl = "https://demo-api.ig.com/gateway/deal"
	}

	return &IGClient{
		APIKey:     apiKey,
		Username:   username,
		Password:   password,
		BaseURL:    baseUrl,
		HTTPClient: &http.Client{},
	}
}
