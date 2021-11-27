package publisher

import (
	"gitlab.com/golibs-starter/golib-sample-core/port"
	"gitlab.com/golibs-starter/golib/pubsub"
)

type EventPublisherAdapter struct {
}

func NewEventPublisherAdapter() port.EventPublisher {
	return &EventPublisherAdapter{}
}

func (e EventPublisherAdapter) Publish(event pubsub.Event) {
	pubsub.Publish(event)
}
