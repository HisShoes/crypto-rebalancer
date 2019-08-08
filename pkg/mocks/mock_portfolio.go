// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hisshoes/crypto-rebalancer/pkg/portfolio (interfaces: Repository)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	portfolio "github.com/hisshoes/crypto-rebalancer/pkg/portfolio"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreatePortfolio mocks base method
func (m *MockRepository) CreatePortfolio(arg0 portfolio.Portfolio) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePortfolio", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePortfolio indicates an expected call of CreatePortfolio
func (mr *MockRepositoryMockRecorder) CreatePortfolio(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePortfolio", reflect.TypeOf((*MockRepository)(nil).CreatePortfolio), arg0)
}

// GenerateID mocks base method
func (m *MockRepository) GenerateID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GenerateID indicates an expected call of GenerateID
func (mr *MockRepositoryMockRecorder) GenerateID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateID", reflect.TypeOf((*MockRepository)(nil).GenerateID))
}

// GetAssetPrice mocks base method
func (m *MockRepository) GetAssetPrice(arg0 string) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAssetPrice", arg0)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAssetPrice indicates an expected call of GetAssetPrice
func (mr *MockRepositoryMockRecorder) GetAssetPrice(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAssetPrice", reflect.TypeOf((*MockRepository)(nil).GetAssetPrice), arg0)
}

// GetPortfolioByID mocks base method
func (m *MockRepository) GetPortfolioByID(arg0 string) (portfolio.Portfolio, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPortfolioByID", arg0)
	ret0, _ := ret[0].(portfolio.Portfolio)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPortfolioByID indicates an expected call of GetPortfolioByID
func (mr *MockRepositoryMockRecorder) GetPortfolioByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPortfolioByID", reflect.TypeOf((*MockRepository)(nil).GetPortfolioByID), arg0)
}

// GetPortfolios mocks base method
func (m *MockRepository) GetPortfolios() ([]portfolio.Portfolio, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPortfolios")
	ret0, _ := ret[0].([]portfolio.Portfolio)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPortfolios indicates an expected call of GetPortfolios
func (mr *MockRepositoryMockRecorder) GetPortfolios() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPortfolios", reflect.TypeOf((*MockRepository)(nil).GetPortfolios))
}

// UpdatePortfolio mocks base method
func (m *MockRepository) UpdatePortfolio(arg0 portfolio.Portfolio) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePortfolio", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePortfolio indicates an expected call of UpdatePortfolio
func (mr *MockRepositoryMockRecorder) UpdatePortfolio(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePortfolio", reflect.TypeOf((*MockRepository)(nil).UpdatePortfolio), arg0)
}
