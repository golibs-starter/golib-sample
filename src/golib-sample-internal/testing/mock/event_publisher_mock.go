package mock

import (
	"gitlab.id.vin/vincart/golib-sample-core/port"
	"gitlab.id.vin/vincart/golib/pubsub"
)

type EventPublisherMock struct {
	receivedEvents []pubsub.Event
}

func NewEventPublisherMock() port.EventPublisher {
	return &EventPublisherMock{
		receivedEvents: make([]pubsub.Event, 0),
	}
}

func (e *EventPublisherMock) Publish(event pubsub.Event) {
	e.receivedEvents = append(e.receivedEvents, event)
}

func (e *EventPublisherMock) ReceivedEvents() []pubsub.Event {
	return e.receivedEvents
}

func (e *EventPublisherMock) Reset() {
	e.receivedEvents = make([]pubsub.Event, 0)
}
