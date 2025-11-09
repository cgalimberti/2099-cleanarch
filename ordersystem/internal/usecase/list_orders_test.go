package usecase

import (
	"testing"

	"github.com/cgalimberti/2099-CleanArch/internal/domain"
	"github.com/cgalimberti/2099-CleanArch/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
)

func TestListOrders(t *testing.T) {
	repo := new(mocks.OrderRepository)
	usecase := NewListOrdersUseCase(repo)

	orders := []domain.Order{
		{ID: "1", Product: "Product 1", Quantity: 2},
		{ID: "2", Product: "Product 2", Quantity: 1},
	}

	repo.On("List").Return(orders, nil)

	result, err := usecase.Execute()

	assert.NoError(t, err)
	assert.Equal(t, orders, result)
	repo.AssertExpectations(t)
}

func TestListOrders_Error(t *testing.T) {
	repo := new(mocks.OrderRepository)
	usecase := NewListOrdersUseCase(repo)

	repo.On("List").Return(nil, assert.AnError)

	result, err := usecase.Execute()

	assert.Error(t, err)
	assert.Nil(t, result)
	repo.AssertExpectations(t)
}
