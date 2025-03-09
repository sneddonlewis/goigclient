package client

import (
	"net/http"
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

func NewLive(apiKey, username, password string) *IGClient {
	return &IGClient{
		APIKey:     apiKey,
		Username:   username,
		Password:   password,
		BaseURL:    "https://api.ig.com/gateway/deal",
		HTTPClient: &http.Client{},
	}
}

func NewDemo(apiKey, username, password string) *IGClient {
	return &IGClient{
		APIKey:     apiKey,
		Username:   username,
		Password:   password,
		BaseURL:    "https://demo-api.ig.com/gateway/deal",
		HTTPClient: &http.Client{},
	}
}
