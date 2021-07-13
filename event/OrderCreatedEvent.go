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
	orderEvent.AbstractEvent = event.NewAbstractEvent(ctx, orderEvent.GetName(), payload)
	return &orderEvent
}

func (r OrderCreatedEvent) GetName() string {
	return "OrderCreatedEvent"
}

func (r OrderCreatedEvent) GetPayload() interface{} {
	return r.Payload
}
