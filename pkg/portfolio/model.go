package portfolio

import "time"

// Asset represents one crypto in the portfolio
type Asset struct {
	// name of the asset e.g. "BTC"
	Name string `json:"name,omitempty"`

	// current amount of the coin in this portfolio
	Value float64 `json:"value,omitempty"`

	// share of the portfolio that this asset should be in
	Share float64 `json:"share,omitempty"`

	// price of one coin in some fiat currency
	Price float64 `json:"price,omitempty"`
}

// Portfolio defines the properties stored in a portfolio
// to track crypto assets for rebalancing
type Portfolio struct {
	ID         string    `json:"id,omitempty"`
	Assets     []Asset   `json:"assets,omitempty"`
	UpdateTime time.Time `json:"updateTime,omitempty"`
}
