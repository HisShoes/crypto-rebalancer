package portfolio

//go:generate mockgen -destination=../mocks/mock_portfolio.go -package=mocks github.com/hisshoes/crypto/rebalancer/portfolio Repository

// Repository interface to allow the portfolio to update
type Repository interface {
	GetAssetPrice(n string) (float64, error)
	GetPortfolio(i int) (Portfolio, error)
	CreatePortfolio(p Portfolio) error
	UpdatePortfolio(p Portfolio) error
}
