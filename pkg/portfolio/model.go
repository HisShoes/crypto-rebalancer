package portfolio

import "time"

// Asset represents one crypto in the portfolio
type Asset struct {
	// name of the asset e.g. "BTC"
	Name string `bson:"name" json:"name"`

	// current amount of the coin in this portfolio
	Value float64 `bson:"value,omitempty" json:"value,omitempty"`

	// share of the portfolio that this asset should be in
	Share float64 `bson:"share,omitempty" json:"share,omitempty"`

	// price of one coin in some fiat currency
	Price float64 `bson:"price,omitempty" json:"price,omitempty"`
}

// Portfolio defines the properties stored in a portfolio
// to track crypto assets for rebalancing
type Portfolio struct {
	ID         string    `bson:"id,omitempty" json:"id,omitempty"`
	Assets     []Asset   `bson:"assets,omitempty" json:"assets,omitempty"`
	UpdateTime time.Time `bson:"updateTime,omitempty" json:"updateTime,omitempty"`
}
