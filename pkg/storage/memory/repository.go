package memory

import (
	"fmt"
	"time"

	"github.com/hisshoes/crypto-rebalancer/pkg/portfolio"

	uuid "github.com/nu7hatch/gouuid"
)

//Storage - track portfolios and assets in memory
type Storage struct {
	portfolios []portfolio.Portfolio
	assets     []portfolio.Asset
}

//GetAssetPrice get the price related to an asset
func (s *Storage) GetAssetPrice(n string) (float64, error) {
	return 0, nil
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
	//generate new uuid
	u, err := uuid.NewV4()
	if err != nil {
		fmt.Println("error:", err)
		return "", err
	}

	//setup non-user set values
	p.ID = u.String()
	p.Updated = time.Now()

	//append to slice and return the ID
	s.portfolios = append(s.portfolios, p)
	return p.ID, nil
}

//UpdatePortfolio update a specific portfolio
func (s *Storage) UpdatePortfolio(p portfolio.Portfolio) error {
	return nil
}
