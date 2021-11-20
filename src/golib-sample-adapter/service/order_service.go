package service

import (
	"context"
	"gitlab.id.vin/vincart/golib-sample-core/entity"
	"gitlab.id.vin/vincart/golib-sample-core/usecase"
)

type OrderService struct {
	getOrderUseCase *usecase.GetOrderUseCase
}

func NewOrderService(getOrderUseCase *usecase.GetOrderUseCase) *OrderService {
	return &OrderService{getOrderUseCase: getOrderUseCase}
}

func (g OrderService) GetById(ctx context.Context, id int64) (*entity.Order, error) {
	return g.getOrderUseCase.GetById(ctx, id)
}
