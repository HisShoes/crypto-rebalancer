package portfolio

//go:generate $GOPATH/bin/mockgen -destination=../mocks/mock_portfolio.go -package=mocks github.com/hisshoes/crypto-rebalancer/pkg/portfolio Repository

// Repository interface to allow the portfolio to update
type Repository interface {
	GetAssetPrice(n string) (float64, error)
	GetPortfolioByID(i int) (Portfolio, error)
	GetPortfolios() ([]Portfolio, error)
	CreatePortfolio(p Portfolio) error
	UpdatePortfolio(p Portfolio) error
}
