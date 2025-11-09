package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/domain"
	"github.com/devfullcycle/20-CleanArch/internal/repository"
)

type ListOrdersUseCase struct {
	OrderRepository repository.OrderRepository
}

func NewListOrdersUseCase(orderRepository repository.OrderRepository) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: orderRepository,
	}
}

func (u *ListOrdersUseCase) Execute() ([]domain.Order, error) {
	return u.OrderRepository.ListOrders()
}