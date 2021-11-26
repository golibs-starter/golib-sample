package event

import (
	"context"
	"gitlab.id.vin/vincart/golib-sample-core/entity"
	"gitlab.id.vin/vincart/golib/web/event"
)

func NewOrderCreatedEvent(ctx context.Context, order *entity.Order) *OrderCreatedEvent {
	return &OrderCreatedEvent{
		AbstractEvent: event.NewAbstractEvent(ctx, "OrderCreatedEvent"),
		PayloadData:   OrderEntityToMessage(order),
	}
}

type OrderCreatedEvent struct {
	*event.AbstractEvent
	PayloadData *OrderMessage `json:"payload"`
}

func (a OrderCreatedEvent) Payload() interface{} {
	return a.PayloadData
}

func (a OrderCreatedEvent) String() string {
	return a.ToString(a)
}

type OrderMessage struct {
	Id          int   `json:"id"`
	TotalAmount int64 `json:"total_amount"`
	CreatedAt   int64 `json:"created_at"`
}

func OrderEntityToMessage(order *entity.Order) *OrderMessage {
	return &OrderMessage{
		Id:          order.Id,
		TotalAmount: order.TotalAmount,
		CreatedAt:   order.CreatedAt.Unix(),
	}
}