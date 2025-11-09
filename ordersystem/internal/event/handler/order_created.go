package handler

import (
	"context"
	"fmt"

	"github.com/devfullcycle/20-CleanArch/internal/domain"
	"github.com/devfullcycle/20-CleanArch/internal/repository"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
)

type OrderCreatedHandler struct {
	OrderRepository repository.OrderRepository
}

func (h *OrderCreatedHandler) Handle(ctx context.Context, event events.Event) error {
	orderID, ok := event.Data["order_id"].(string)
	if !ok {
		return fmt.Errorf("invalid order_id in event data")
	}

	order, err := h.OrderRepository.FindByID(ctx, orderID)
	if err != nil {
		return fmt.Errorf("could not find order: %w", err)
	}

	// Process the order creation event (e.g., send a notification, update a status, etc.)
	fmt.Printf("Order created: %+v\n", order)

	return nil
}