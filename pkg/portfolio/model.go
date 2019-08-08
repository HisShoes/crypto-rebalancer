package portfolio

import "time"

// Asset represents one crypto in the portfolio
type Asset struct {
	// name of the asset e.g. "BTC"
	Name string

	// current value held in the coin for this portfolio
	Value float64

	// share of the portfolio that this asset should be in
	Share float64

	// price of one coin in some fiat currency
	Price float64
}

// Portfolio defines the properties stored in a portfolio
// to track crypto assets for rebalancing
type Portfolio struct {
	ID      int
	Assets  []Asset
	Updated time.Time
}
