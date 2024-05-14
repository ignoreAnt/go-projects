package client_test

import (
	"github.com/stretchr/testify/mock"
	"go-projects/inventory"
)

// MockSupplierService is a mock implementation of the inventory.SupplierService interface
type MockSupplierService struct {
	mock.Mock
}

// Implement PlaceOrder with a mock implementation (replace with desired behavior)
func (m *MockSupplierService) PlaceOrder(o *inventory.Order) error {
	args := m.Called(o)
	return args.Error(0)
}

// Implement GetStatus with a mock implementation (replace with desired behavior)
func (m *MockSupplierService) GetStatus(o *inventory.Order) error {
	args := m.Called(o)
	return args.Error(0)
}
