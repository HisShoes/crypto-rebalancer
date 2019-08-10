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
		ID: "1",
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
	mockRepo.EXPECT().Portfolio("1").Return(portfolio.Portfolio{}, portfolio.ErrMissing).Times(1)
	p, err := s.Portfolio("1")
	if err != portfolio.ErrMissing {
		t.Errorf("Portfolio doesn't exist: Error missing not returned")
	}
	if p.ID != "" {
		t.Errorf("Portfolio doesn't exist: zero ID not returned ")
	}

	// Portfolio Exists: Test to check when portfolio returned from repo
	// it's returned from the service
	mockRepo.EXPECT().Portfolio("1").Return(testPortfolio, nil).Times(1)
	p, err = s.Portfolio("1")
	if err != nil {
		t.Errorf("Portfolio Exists: Error returned instead of portfolio")
	}
	if p.ID != "1" {
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
			ID: "1",
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
			ID: "2",
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
	mockRepo.EXPECT().ListPortfolios().Return([]portfolio.Portfolio{}, errors.New("Some error")).Times(1)
	ps, err := s.ListPortfolios()
	if err == nil {
		t.Errorf("Portfolio List Error: Error not returned")
	}
	if len(ps) != 0 {
		t.Errorf("Portfolio List Error: zero ID not returned ")
	}

	// Portfolio List: Test to check when portfolios returned from repo
	// it's returned from the service
	mockRepo.EXPECT().ListPortfolios().Return(testPortfolios, nil).Times(1)
	ps, err = s.ListPortfolios()
	if err != nil {
		t.Errorf("Portfolio List: Error returned instead of portfolios")
	}
	if len(ps) != 2 {
		t.Errorf("Portfolio List: Wrong number of records returned")
	}
}

// TestCreatePortfolio - test creating a portfolio
func TestCreatePortfolio(t *testing.T) {
	// setup mocking objects
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mocks.NewMockRepository(mockCtrl)

	// setup test portfolio to use in call to create
	testPortfolio := portfolio.Portfolio{
		Assets: []portfolio.Asset{
			portfolio.Asset{
				Name:  "BTC",
				Value: 100,
				Share: 50,
				Price: 8000,
			},
		},
	}

	testID := "FIRSTID"

	// setup the service using mock repo
	s := portfolio.NewService(mockRepo)

	// Portfolio Create: Check portfolio is created properly
	mockRepo.EXPECT().CreatePortfolio(testPortfolio).Return(testID, nil).Times(1)
	ID, err := s.CreatePortfolio(testPortfolio)
	if ID != testID {
		t.Errorf("Portfolio Create: Wrong ID returned")
	}
	if err != nil {
		t.Errorf("Portfolio Create: Error returned but not thrown")
	}

	// Portfolio Create Error: Check portfolio is created properly
	mockRepo.EXPECT().CreatePortfolio(testPortfolio).Return("", errors.New("Some error")).Times(1)
	ID, err = s.CreatePortfolio(testPortfolio)
	if ID != "" {
		t.Errorf("Portfolio Create Error: Non zero ID returned")
	}
	if err == nil {
		t.Errorf("Portfolio Create: Error thrown but not returned")
	}

}

func TestUpdatePrice(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	// mockRepo := mocks.NewMockRepository(mockCtrl)
}
