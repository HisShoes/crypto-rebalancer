package portfolio_test

import (
	"errors"
	"testing"

	"github.com/hisshoes/crypto-rebalancer/pkg/portfolio"

	"github.com/golang/mock/gomock"
	"github.com/hisshoes/crypto-rebalancer/pkg/mocks"
)

// TestGetPortfolioByID - unit test portfolio.service.GetPortfolio
func TestGetPortfolioByID(t *testing.T) {
	// setup mocking objects
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mocks.NewMockRepository(mockCtrl)

	// setup test portfolio to use in mocking
	testPortfolio := portfolio.Portfolio{
		ID: 1,
		Assets: []portfolio.Asset{
			portfolio.Asset{
				Name:  "BTC",
				Value: 100,
				Share: 50,
				Price: 8000,
			},
		},
	}

	//setup the service using mock repo
	s := portfolio.NewService(mockRepo)

	// Portfolio doesn't exist: Test to check when error is returned from repo
	// error is returned from service
	mockRepo.EXPECT().GetPortfolioByID(1).Return(portfolio.Portfolio{}, portfolio.ErrMissing).Times(1)
	p, err := s.GetPortfolioByID(1)
	if err != portfolio.ErrMissing {
		t.Errorf("Portfolio doesn't exist: Error missing not returned")
	}
	if p.ID != 0 {
		t.Errorf("Portfolio doesn't exist: zero ID not returned ")
	}

	// Portfolio Exists: Test to check when portfolio returned from repo
	// it's returned from the service
	mockRepo.EXPECT().GetPortfolioByID(1).Return(testPortfolio, nil).Times(1)
	p, err = s.GetPortfolioByID(1)
	if err != nil {
		t.Errorf("Portfolio Exists: Error returned instead of portfolio")
	}
	if p.ID != 1 {
		t.Errorf("Portfolio Exists: portfolio with ID 1 not returned")
	}
}

// TestGetPortfolio - unit test portfolio.service.GetPortfolio
func TestGetPortfolios(t *testing.T) {
	// setup mocking objects
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mocks.NewMockRepository(mockCtrl)

	// setup test portfolio to use in mocking
	testPortfolios := []portfolio.Portfolio{
		portfolio.Portfolio{
			ID: 1,
			Assets: []portfolio.Asset{
				portfolio.Asset{
					Name:  "BTC",
					Value: 100,
					Share: 50,
					Price: 8000,
				},
			},
		},
		portfolio.Portfolio{
			ID: 2,
			Assets: []portfolio.Asset{
				portfolio.Asset{
					Name:  "ETH",
					Value: 100,
					Share: 50,
					Price: 8000,
				},
			},
		},
	}

	//setup the service using mock repo
	s := portfolio.NewService(mockRepo)

	// Portfolio List: Test to check when error is returned from repo
	// error is returned from service
	mockRepo.EXPECT().GetPortfolios().Return([]portfolio.Portfolio{}, errors.New("Some error")).Times(1)
	ps, err := s.GetAllPortfolios()
	if err == nil {
		t.Errorf("Portfolio doesn't exist: Error missing not returned")
	}
	if len(ps) != 0 {
		t.Errorf("Portfolio doesn't exist: zero ID not returned ")
	}

	// Portfolio List: Test to check when portfolios returned from repo
	// it's returned from the service
	mockRepo.EXPECT().GetPortfolios().Return(testPortfolios, nil).Times(1)
	ps, err = s.GetAllPortfolios()
	if err != nil {
		t.Errorf("Portfolio List: Error returned instead of portfolios")
	}
	if len(ps) != 2 {
		t.Errorf("Portfolio List: Wrong number of records returned")
	}
}
