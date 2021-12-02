package usecase

import (
	"context"
	"gitlab.id.vin/vincart/golib-sample-core/entity"
	"gitlab.id.vin/vincart/golib-sample-core/exception"
	"gitlab.id.vin/vincart/golib-sample-core/port"
	"gitlab.id.vin/vincart/golib/web/log"
)

type GetOrderUseCase struct {
	orderRepo port.OrderRepository
}

func NewGetOrderUseCase(orderRepo port.OrderRepository) *GetOrderUseCase {
	return &GetOrderUseCase{orderRepo: orderRepo}
}

func (g GetOrderUseCase) GetById(ctx context.Context, id int64) (*entity.Order, error) {
	order, err := g.orderRepo.FindById(ctx, id)
	if err != nil {
		log.Error(ctx, "Cannot get order by id [%d] due by error [%v]", id, err)
		return nil, err
	}
	return order, nil
}

func (g GetOrderUseCase) GetByIdAndUser(ctx context.Context, userId string, orderId int64) (*entity.Order, error) {
	order, err := g.GetById(ctx, orderId)
	if err != nil {
		return nil, err
	}
	if order.UserId != userId {
		return nil, exception.OrderNotFound
	}
	return order, nil
}
