package event

import (
	"context"
	"gitlab.id.vin/vincart/golib-sample-core/entity"
	"gitlab.id.vin/vincart/golib/web/event"
)

type OrderCreatedEvent struct {
	*event.AbstractEvent
}

func NewOrderCreatedEvent(ctx context.Context, payload *entity.Order) *OrderCreatedEvent {
	return &OrderCreatedEvent{AbstractEvent: event.NewAbstractEvent(ctx, "OrderCreatedEvent", payload)}
}

func (a OrderCreatedEvent) String() string {
	return a.ToString(a)
}
