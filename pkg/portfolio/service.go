package portfolio

import (
	"errors"
)

// Service to provide portfolio functionality
type Service interface {
	Portfolio(id string) (Portfolio, error)
	CreatePortfolio(Portfolio) (string, error)
	ListPortfolios() ([]Portfolio, error)
	RebalancePortfolio(id string) (Portfolio, error)
}

type service struct {
	repo Repository
}

// ErrDuplicate is used when a portfolio already exists.
var ErrDuplicate = errors.New("portfolio already exists")

// ErrMissing is used when a portfolio already exists.
var ErrMissing = errors.New("portfolio not found")

// NewService creates a new portfolio service with dependencies
func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

// CreatePortfolio Creates a new portfolio using the repo
// takes in an initial portfolio object
func (s *service) CreatePortfolio(p Portfolio) (string, error) {
	return s.repo.CreatePortfolio(p)
}

// Portfolio retrieves the portfolio matching the id passed in
func (s *service) Portfolio(id string) (Portfolio, error) {
	return s.repo.Portfolio(id)
}

// ListPortfolios retrieves all portfolios
func (s *service) ListPortfolios() ([]Portfolio, error) {
	return s.repo.ListPortfolios()
}

// RebalancePortfolio calls the repo to get asset pricing
// and redistributes based on asset shares
func (s *service) RebalancePortfolio(id string) (Portfolio, error) {
	// Get the portfolio from the repo using the id passed in
	p, err := s.repo.Portfolio(id)
	if err != nil {
		return Portfolio{}, err
	}

	// loop through assets, get the prices and update the asset value
	for _, a := range p.Assets {
		a.Price, err = s.repo.GetAssetPrice(a.Name)
		if err != nil {
			return Portfolio{}, err
		}
	}

	// figure out total value of the portfolio

	// loop through the assets, update the value they should be adjusted to
	// reduce the difference by exchange modifier

	// Update the portfolio using the repo
	err = s.repo.UpdatePortfolio(p)
	if err != nil {
		return Portfolio{}, err
	}

	return Portfolio{}, nil
}
