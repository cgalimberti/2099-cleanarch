package usecase

import (
	"github.com/cgalimberti/2099-CleanArch/internal/domain"
	"github.com/cgalimberti/2099-CleanArch/internal/repository"
)

type ListOrdersUseCase struct {
	orderRepository repository.OrderRepository
}

func NewListOrdersUseCase(orderRepository repository.OrderRepository) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		orderRepository: orderRepository,
	}
}

func (u *ListOrdersUseCase) Execute() ([]domain.Order, error) {
	return u.orderRepository.ListOrders()
}
