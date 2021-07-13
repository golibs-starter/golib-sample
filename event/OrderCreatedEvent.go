package event

import (
	"context"
	"gitlab.id.vin/vincart/golib/web/event"
)

type OrderCreatedEvent struct {
	event.AbstractEvent
}

type OrderCreatedPayload struct {
	Code        string `json:"code"`
	TotalAmount int    `json:"total_amount"`
}

func NewOrderCreatedEvent(ctx context.Context, payload OrderCreatedPayload) *OrderCreatedEvent {
	orderEvent := OrderCreatedEvent{}
	orderEvent.AbstractEvent = event.NewAbstractEvent(ctx, orderEvent.Name(), payload)
	return &orderEvent
}

func (r OrderCreatedEvent) Name() string {
	return "OrderCreatedEvent"
}

func (r OrderCreatedEvent) Payload() interface{} {
	return r.Payload
}
