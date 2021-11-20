package publisher

import (
	"gitlab.id.vin/vincart/golib-sample-core/port"
	"gitlab.id.vin/vincart/golib/pubsub"
)

type EventPublisherAdapter struct {
}

func NewEventPublisherAdapter() port.EventPublisher {
	return &EventPublisherAdapter{}
}

func (e EventPublisherAdapter) Publish(event pubsub.Event) {
	pubsub.Publish(event)
}
