package handler

import (
	"encoding/json"
	"gitlab.com/golibs-starter/golib-message-bus/kafka/core"
	"gitlab.com/golibs-starter/golib-sample-core/event"
	"gitlab.com/golibs-starter/golib/web/log"
)

type OrderCreatedEventDummyHandler struct {
	collector *OrderEventDummyCollector
}

func NewOrderCreatedEventDummyHandler(collector *OrderEventDummyCollector) core.ConsumerHandler {
	log.Infof("OrderCreatedEventDummyHandler init")
	return &OrderCreatedEventDummyHandler{collector: collector}
}

func (c *OrderCreatedEventDummyHandler) HandlerFunc(msg *core.ConsumerMessage) {
	var e event.OrderCreatedEvent
	if err := json.Unmarshal(msg.Value, &e); err != nil {
		log.Errore(e, "[OrderCreatedEventDummyHandler] Error when unmarshal event message, detail: ", err)
		return
	}
	c.collector.createdEvents = append(c.collector.createdEvents, e)
	log.Infoe(e, "[OrderCreatedEventDummyHandler] Success to unmarshal event message: [%+v]", e.Payload())
}

func (c OrderCreatedEventDummyHandler) Close() {
	//
}
