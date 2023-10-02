package usecase

import (
	"context"
	"github.com/golibs-starter/golib-sample-core/entity"
	"github.com/golibs-starter/golib-sample-core/entity/request"
	"github.com/golibs-starter/golib-sample-core/event"
	"github.com/golibs-starter/golib-sample-core/port"
	"github.com/pkg/errors"
)

type CreateOrderUseCase struct {
	orderRepo      port.OrderRepository
	eventPublisher port.EventPublisher
}

func NewCreateOrderUseCase(orderRepo port.OrderRepository, eventPublisher port.EventPublisher) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		orderRepo:      orderRepo,
		eventPublisher: eventPublisher,
	}
}

func (c CreateOrderUseCase) Create(ctx context.Context, req *request.CreateOrderRequest) (*entity.Order, error) {
	order, err := c.orderRepo.CreateOrder(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "create order failed")
	}
	c.eventPublisher.Publish(event.NewOrderCreatedEvent(ctx, order))
	return order, nil
}
