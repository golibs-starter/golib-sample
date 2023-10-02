package port

import "github.com/golibs-starter/golib/pubsub"

type EventPublisher interface {
	Publish(e pubsub.Event)
}
