package memory

import (
	"time"

	"github.com/segmentio/ksuid"

	"github.com/hisshoes/crypto-rebalancer/pkg/portfolio"
)

//Storage - track portfolios and assets in memory
type Storage struct {
	portfolios []portfolio.Portfolio
	assets     []portfolio.Asset
}

//GetAssetPrice get the price related to an asset
func (s *Storage) GetAssetPrice(n string) (float64, error) {
	for _, a := range s.assets {
		if a.Name == n {
			return a.Price, nil
		}
	}

	return 0, portfolio.ErrMissing
}

//Portfolio return a portfolio relating to an id
func (s *Storage) Portfolio(id string) (portfolio.Portfolio, error) {
	for _, p := range s.portfolios {
		if p.ID == id {
			return p, nil
		}
	}
	return portfolio.Portfolio{}, portfolio.ErrMissing
}

//ListPortfolios return all the portfolios
func (s *Storage) ListPortfolios() ([]portfolio.Portfolio, error) {
	return s.portfolios, nil
}

//CreatePortfolio create a portfolio and append to the slice
func (s *Storage) CreatePortfolio(p portfolio.Portfolio) (string, error) {

	//setup non-user set values
	p.ID = generateID()
	p.UpdateTime = time.Now()

	//append to slice and return the ID
	s.portfolios = append(s.portfolios, p)
	return p.ID, nil
}

//UpdatePortfolio update a specific portfolio
func (s *Storage) UpdatePortfolio(p portfolio.Portfolio) error {
	for _, cp := range s.portfolios {
		if cp.ID == p.ID {
			cp.Assets = p.Assets
			cp.UpdateTime = time.Now()
			return nil
		}
	}

	return portfolio.ErrMissing
}

func generateID() string {
	return ksuid.New().String()
}
