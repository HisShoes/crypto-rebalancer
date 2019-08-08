package portfolio

import "errors"

//Service to provide portfolio functionality
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

//NewService creates a new portfolio service with dependencies
func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

//CreatePortfolio Creates a new portfolio using the repo
// takes in an initial portfolio object
func (s *service) CreatePortfolio(p Portfolio) (Portfolio, error) {

	return Portfolio{}, nil
}

//GetPortfolio retrieves the portfolio values
func (s *service) GetPortfolio(i int) (Portfolio, error) {
	return s.repo.GetPortfolio(i)
}

//RebalancePortfolio calls the repo to get asset pricing
// and redistributes based on asset shares
func (s *service) RebalancePortfolio(i int) (Portfolio, error) {

	return Portfolio{}, nil
}
