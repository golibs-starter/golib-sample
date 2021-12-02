package service

import (
	"context"
	"gitlab.com/golibs-starter/golib-sample-core/entity"
	"gitlab.com/golibs-starter/golib-sample-core/entity/request"
	"gitlab.com/golibs-starter/golib-sample-core/usecase"
)

type OrderService struct {
	getOrderUseCase    *usecase.GetOrderUseCase
	createOrderUseCase *usecase.CreateOrderUseCase
}

func NewOrderService(
	getOrderUseCase *usecase.GetOrderUseCase,
	createOrderUseCase *usecase.CreateOrderUseCase,
) *OrderService {
	return &OrderService{
		getOrderUseCase:    getOrderUseCase,
		createOrderUseCase: createOrderUseCase,
	}
}

func (g OrderService) GetById(ctx context.Context, orderId int64) (*entity.Order, error) {
	return g.getOrderUseCase.GetById(ctx, orderId)
}

func (g OrderService) GetByIdAndUser(ctx context.Context, userId string, orderId int64) (*entity.Order, error) {
	return g.getOrderUseCase.GetByIdAndUser(ctx, userId, orderId)
}

func (g OrderService) Create(ctx context.Context, req *request.CreateOrderRequest) (*entity.Order, error) {
	return g.createOrderUseCase.Create(ctx, req)
}
