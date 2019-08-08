package portfolio_test

import (
	"testing"

	"github.com/hisshoes/crypto-rebalancer/pkg/portfolio"

	"github.com/golang/mock/gomock"
	"github.com/hisshoes/crypto-rebalancer/pkg/mocks"
)

//TestGetPortfolio - unit test portfolio.service.GetPortfolio
func TestGetPortfolio(t *testing.T) {
	//setup mocking objects
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mocks.NewMockRepository(mockCtrl)

	//setup test portfolio to use in mocking
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

	//Portfolio doesn't exist: Test to check when error is returned from repo
	// error is returned from service
	mockRepo.EXPECT().GetPortfolio(1).Return(portfolio.Portfolio{}, portfolio.ErrMissing).Times(1)
	s := portfolio.NewService(mockRepo)
	p, err := s.GetPortfolio(1)
	if err != portfolio.ErrMissing {
		t.Errorf("Portfolio doesn't exist: Error missing not returned")
	}
	if p.ID != 0 {
		t.Errorf("Portfolio doesn't exist: zero ID not returned ")
	}

	//Portfolio Exists: Test to check when portfolio returned from repo
	// it's returned from the service
	mockRepo.EXPECT().GetPortfolio(1).Return(testPortfolio, nil).Times(1)
	p, err = s.GetPortfolio(1)
	if err != nil {
		t.Errorf("Portfolio Exists: Error returned instead of portfolio")
	}
	if p.ID != 1 {
		t.Errorf("Portfolio Exists: portfolio with ID 1 not returned")
	}

}
