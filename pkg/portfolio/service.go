package portfolio

import "errors"

// Service to provide portfolio functionality
type Service interface {
	GetPortfolio(i int) (Portfolio, error)
	CreatePortfolio(Portfolio) (Portfolio, error)
	RebalancePortfolio(i int) (Portfolio, error)
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
func (s *service) CreatePortfolio(p Portfolio) (Portfolio, error) {

	return Portfolio{}, nil
}

// GetPortfolio retrieves the portfolio values
func (s *service) GetPortfolio(i int) (Portfolio, error) {
	return s.repo.GetPortfolio(i)
}

// RebalancePortfolio calls the repo to get asset pricing
// and redistributes based on asset shares
func (s *service) RebalancePortfolio(i int) (Portfolio, error) {
	// Get the portfolio from the repo using the id passed in

	// loop through assets, update the prices

	// figure out total value of the portfolio

	// loop through the assets, update the value they should be adjusted to
	// reduce the difference by exchange modifier

	// Update the portfolio using the repo

	return Portfolio{}, nil
}

// UpdatePrice calls the repo to update the price of 1 coin in fiat
func (s *service) UpdatePrice(a *Asset) error {
	var err error
	a.Price, err = s.repo.GetAssetPrice(a.Name)
	return err
}
