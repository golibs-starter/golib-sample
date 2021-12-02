package handler

import "gitlab.id.vin/vincart/golib-sample-core/event"

type OrderEventDummyCollector struct {
	createdEvents []event.OrderCreatedEvent
}

func NewOrderEventDummyCollector() *OrderEventDummyCollector {
	return &OrderEventDummyCollector{
		createdEvents: make([]event.OrderCreatedEvent, 0),
	}
}

func (o *OrderEventDummyCollector) CreatedEvents() []event.OrderCreatedEvent {
	return o.createdEvents
}
