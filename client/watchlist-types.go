package client

type Watchlists struct {
	Watchlists []Watchlist `json:"watchlists"`
}

type Watchlist struct {
	ID                     string `json:"id"`
	Name                   string `json:"name"`
	DefaultSystemWatchlist bool   `json:"defaultSystemWatchlist"`
	Deletable              bool   `json:"deleteable"`
	Editable               bool   `json:"editable"`
}

type WatchlistContent struct {
	Markets []MarketDataDetailed `json:"markets"`
}

type CreateWatchlistRequest struct {
	Name  string   `json:"name"`
	Epics []string `json:"epics"`
}

type CreateWatchlistResponse struct {
	Status      string `json:"status"`
	WatchlistID string `json:"watchlistId"`
}

type WatchlistAddMarketRequest struct {
	Epic string `json:"epic"`
}
