package client_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sneddonlewis/goigclient/client"
)

func TestGetWorkingOrders(t *testing.T) {
	testAPIKey := "testAPIKey"

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/workingorders" {
			t.Errorf("Unexpected URL path: %s", r.URL.Path)
		}
		if r.Header.Get("X-IG-API-KEY") != testAPIKey {
			t.Errorf("Expected API key header, got %s", r.Header.Get("X-IG-API-KEY"))
		}
		if r.Header.Get("Version") != "2" {
			t.Errorf("Expected Version header to be 2, got %s", r.Header.Get("Version"))
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"working-orders": []}`))
	}))
	defer mockServer.Close()

	apiClient := client.NewIGClient(testAPIKey, "", "", true)
	apiClient.BaseURL = mockServer.URL

	workingOrders, err := apiClient.AllWorkingOrders()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if workingOrders == nil {
		t.Errorf("Expected response deserialised, got nil")
	}
}
