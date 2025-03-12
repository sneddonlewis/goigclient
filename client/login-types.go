package client

type OAuthToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

type AuthResponse struct {
	AccountID             string     `json:"accountId"`
	ClientID              string     `json:"clientId"`
	LightstreamerEndpoint string     `json:"lightstreamerEndpoint"`
	OAuth                 OAuthToken `json:"oauthToken"`
	TimezoneOffset        int        `json:"timezoneOffset"`
}

type SessionResponse struct {
	AccountID             string `json:"accountId"`
	ClientID              string `json:"clientId"`
	Currency              string `json:"currency"`
	LightstreamerEndpoint string `json:"lightstreamerEndpoint"`
	Locale                string `json:"locale"`
	TimezoneOffset        int    `json:"timezoneOffset"`
}

type SwitchAccountRequest struct {
	AccountID      string `json:"accountId"`
	DefaultAccount *bool  `json:"defaultAccount,omitempty"`
}

type SwitchAccountResponse struct {
	DealingEnabled        bool `json:"dealingEnabled"`
	HasActiveDemoAccounts bool `json:"hasActiveDemoAccounts"`
	HasActiveLiveAccounts bool `json:"hasActiveLiveAccounts"`
	TrailingStopsEnabled  bool `json:"trailingStopsEnabled"`
}

type LoginKeyResponse struct {
	EncryptionKey string `json:"encryptionKey"`
	TimeStamp     int64  `json:"timeStamp"`
}

type LoginEncryptedRequest struct {
	EncryptedPassword bool   `json:"encryptedPassword"`
	Identifier        string `json:"identifier"`
	Password          string `json:"password"`
}

type LoginEncryptedResponse struct {
	AccountInfo           AccountFinancialData `json:"accountInfo"`
	Accounts              []AccountDetails     `json:"accounts"`
	ClientID              string               `json:"clientId"`
	CurrencyIsoCode       string               `json:"currencyIsoCode"`
	CurrencySymbol        string               `json:"currencySymbol"`
	CurrentAccountID      string               `json:"currentAccountId"`
	DealingEnabled        bool                 `json:"dealingEnabled"`
	HasActiveDemoAccounts bool                 `json:"hasActiveDemoAccounts"`
	HasActiveLiveAccounts bool                 `json:"hasActiveLiveAccounts"`
	LightstreamerEndpoint string               `json:"lightstreamerEndpoint"`
	ReroutingEnvironment  string               `json:"reroutingEnvironment"`
	TimezoneOffset        int                  `json:"timezoneOffset"`
	TrailingStopsEnabled  bool                 `json:"trailingStopsEnabled"`
}

type AccountFinancialData struct {
	Available  float64 `json:"available"`
	Balance    float64 `json:"balance"`
	Deposit    float64 `json:"deposit"`
	ProfitLoss float64 `json:"profitLoss"`
}

type AccountDetails struct {
	AccountID   string `json:"accountId"`
	AccountName string `json:"accountName"`
	AccountType string `json:"accountType"`
	Preferred   bool   `json:"preferred"`
}

type RefreshSessionRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshSessionResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}
