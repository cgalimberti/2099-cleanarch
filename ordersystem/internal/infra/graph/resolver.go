package graph

import (
	"context"

	"github.com/devfullcycle/20-CleanArch/internal/usecase"
	"github.com/devfullcycle/20-CleanArch/internal/domain"
)

type Resolver struct {
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrdersUseCase
}

func (r *Resolver) ListOrders(ctx context.Context) ([]*domain.Order, error) {
	orders, err := r.ListOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}
	return orders, nil
}