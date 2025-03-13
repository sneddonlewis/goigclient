package client

type ApplicationsResponse struct {
	Applications []Application `json:"applications"`
}

// Application represents an individual application associated with the account.
type Application struct {
	APIKey                 string `json:"apiKey"`
	AllowedAPIKey          bool   `json:"allowed"`
	ConcurrentRequests     int    `json:"concurrentRequests"`
	CreatedDate            string `json:"createdDate"`
	Identifier             string `json:"identifier"`
	LastLogin              string `json:"lastLogin"`
	Status                 string `json:"status"`
	ThrottlingLimit        int    `json:"throttlingLimit"`
	ThrottlingRemaining    int    `json:"throttlingRemaining"`
	ThrottlingResetSeconds int    `json:"throttlingResetSeconds"`
}

type UpdateApplicationRequest struct {
	AllowanceAccountOverall int    `json:"allowanceAccountOverall"`
	AllowanceAccountTrading int    `json:"allowanceAccountTrading"`
	APIKey                  string `json:"apiKey"`
	Status                  string `json:"status"` // DISABLED, ENABLED, REVOKED
}

type UpdateApplicationResponse struct {
	AllowEquities                  bool   `json:"allowEquities"`
	AllowQuoteOrders               bool   `json:"allowQuoteOrders"`
	AllowanceAccountHistoricalData int    `json:"allowanceAccountHistoricalData"`
	AllowanceAccountOverall        int    `json:"allowanceAccountOverall"`
	AllowanceAccountTrading        int    `json:"allowanceAccountTrading"`
	AllowanceApplicationOverall    int    `json:"allowanceApplicationOverall"`
	APIKey                         string `json:"apiKey"`
	ConcurrentSubscriptionsLimit   int    `json:"concurrentSubscriptionsLimit"`
	CreatedDate                    string `json:"createdDate"`
	Name                           string `json:"name"`
	Status                         string `json:"status"` // DISABLED, ENABLED, REVOKED
}

type DisableApplicationResponse struct {
	AllowEquities                  bool   `json:"allowEquities"`
	AllowQuoteOrders               bool   `json:"allowQuoteOrders"`
	AllowanceAccountHistoricalData int    `json:"allowanceAccountHistoricalData"`
	AllowanceAccountOverall        int    `json:"allowanceAccountOverall"`
	AllowanceAccountTrading        int    `json:"allowanceAccountTrading"`
	AllowanceApplicationOverall    int    `json:"allowanceApplicationOverall"`
	APIKey                         string `json:"apiKey"`
	ConcurrentSubscriptionsLimit   int    `json:"concurrentSubscriptionsLimit"`
	CreatedDate                    string `json:"createdDate"`
	Name                           string `json:"name"`
	Status                         string `json:"status"` // DISABLED, ENABLED, REVOKED
}
