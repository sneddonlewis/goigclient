package client

type DealResponse struct {
	DealReference string `json:"dealReference"`
}

type Metadata struct {
	Paging Paging `json:"paging"`
}
