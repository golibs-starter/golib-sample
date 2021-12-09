package dummy

import (
	"encoding/json"
	"gitlab.id.vin/vincart/golib-message-bus/kafka/core"
	"gitlab.id.vin/vincart/golib-sample-core/event"
	"gitlab.id.vin/vincart/golib/web/log"
)

type OrderCreatedEventDummyHandler struct {
	collector *OrderEventDummyCollector
}

func NewOrderCreatedEventDummyHandler(collector *OrderEventDummyCollector) core.ConsumerHandler {
	return &OrderCreatedEventDummyHandler{collector: collector}
}

func (c *OrderCreatedEventDummyHandler) HandlerFunc(msg *core.ConsumerMessage) {
	var e event.OrderCreatedEvent
	if err := json.Unmarshal(msg.Value, &e); err != nil {
		log.Errore(e, "[OrderCreatedEventDummyHandler] Error when unmarshal event message, detail: ", err)
		return
	}
	c.collector.createdEvents = append(c.collector.createdEvents, e)
	log.Infoe(e, "[OrderCreatedEventDummyHandler] Success to unmarshal event message")
}

func (c OrderCreatedEventDummyHandler) Close() {
	//
}
