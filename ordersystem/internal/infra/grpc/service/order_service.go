package service

import (
	"context"

	"github.com/cgalimberti/2099-CleanArch/internal/infra/grpc/pb"
	"github.com/cgalimberti/2099-CleanArch/internal/usecase"
)

type OrderService struct {
	usecase.ListOrdersUseCase
}

func NewOrderService(listOrdersUseCase usecase.ListOrdersUseCase) *OrderService {
	return &OrderService{
		ListOrdersUseCase: listOrdersUseCase,
	}
}

func (s *OrderService) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	orders, err := s.ListOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var orderResponses []*pb.Order
	for _, order := range orders {
		orderResponses = append(orderResponses, &pb.Order{
			Id:     order.ID,
			Amount: order.Amount,
			Status: order.Status,
		})
	}

	return &pb.ListOrdersResponse{Orders: orderResponses}, nil
}
