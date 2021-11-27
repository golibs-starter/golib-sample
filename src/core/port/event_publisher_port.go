package port

import "gitlab.com/golibs-starter/golib/pubsub"

type EventPublisher interface {
	Publish(e pubsub.Event)
}
