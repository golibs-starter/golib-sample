package usecase

import (
	"context"
	"gitlab.id.vin/vincart/golib-sample-core/entity"
	"gitlab.id.vin/vincart/golib-sample-core/entity/request"
	"gitlab.id.vin/vincart/golib-sample-core/event"
	"gitlab.id.vin/vincart/golib-sample-core/port"
	"gitlab.id.vin/vincart/golib/web/log"
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
		log.Error(ctx, "Cannot create order due by error [%v]", err)
		return nil, err
	}
	c.eventPublisher.Publish(event.NewOrderCreatedEvent(ctx, order))
	return order, nil
}
