package usecase

import (
	"context"
	"github.com/golibs-starter/golib-sample-core/entity"
	"github.com/golibs-starter/golib-sample-core/exception"
	"github.com/golibs-starter/golib-sample-core/port"
	"github.com/pkg/errors"
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
		return nil, errors.WithMessage(err, "get order failed")
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
