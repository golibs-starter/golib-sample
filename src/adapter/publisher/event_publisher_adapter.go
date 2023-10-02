package publisher

import (
	"github.com/golibs-starter/golib-sample-core/port"
	"github.com/golibs-starter/golib/pubsub"
)

type EventPublisherAdapter struct {
}

func NewEventPublisherAdapter() port.EventPublisher {
	return &EventPublisherAdapter{}
}

func (e EventPublisherAdapter) Publish(event pubsub.Event) {
	pubsub.Publish(event)
}
