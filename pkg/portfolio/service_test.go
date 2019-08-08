package portfolio_test

import (
	"testing"

	"github.com/hisshoes/crypto/rebalancer/pkg/portfolio"

	"github.com/golang/mock/gomock"
	"github.com/hisshoes/crypto/rebalancer/pkg/mocks"
)

func TestGetPortfolio(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockRepository(mockCtrl)

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

	mockRepo.EXPECT().GetPortfolio(1).Return(portfolio.Portfolio{}, portfolio.ErrMissing).Times(1)

	s := portfolio.NewService(mockRepo)

	p, err := s.GetPortfolio(1)

	if err != portfolio.ErrMissing {
		t.Errorf("Portfolio doesn't exist: Error missing not returned")
	}

	if p.ID != 0 {
		t.Errorf("Portfolio doesn't exist: zero ID not returned ")
	}

	mockRepo.EXPECT().GetPortfolio(1).Return(testPortfolio, nil).Times(1)

	p, err = s.GetPortfolio(1)

	if err != nil {
		t.Errorf("Portfolio Exists: Error returned instead of portfolio")
	}

	if p.ID != 1 {
		t.Errorf("Portfolio Exists: portfolio with ID 1 not returned")
	}

}
