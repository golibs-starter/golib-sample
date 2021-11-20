package port

import "gitlab.id.vin/vincart/golib/pubsub"

type EventPublisher interface {
	Publish(e pubsub.Event)
}
