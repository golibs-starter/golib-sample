package usecase

import (
	"context"
	"gitlab.com/golibs-starter/golib-sample-core/entity"
	"gitlab.com/golibs-starter/golib-sample-core/exception"
	"gitlab.com/golibs-starter/golib-sample-core/port"
	"gitlab.com/golibs-starter/golib/web/log"
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
