package handler

import (
	"encoding/json"
	"gitlab.id.vin/vincart/golib-message-bus/kafka/core"
	"gitlab.id.vin/vincart/golib-sample-adapter/service"
	"gitlab.id.vin/vincart/golib-sample-core/event"
	"gitlab.id.vin/vincart/golib/web/log"
)

type SendOrderToDeliveryProviderHandler struct {
	orderDeliveryService *service.OrderDeliveryService
}

func NewSendOrderToDeliveryProviderHandler(orderDeliveryService *service.OrderDeliveryService) core.ConsumerHandler {
	return &SendOrderToDeliveryProviderHandler{orderDeliveryService: orderDeliveryService}
}

func (c *SendOrderToDeliveryProviderHandler) HandlerFunc(msg *core.ConsumerMessage) {
	var e event.OrderCreatedEvent
	if err := json.Unmarshal(msg.Value, &e); err != nil {
		log.Errore(e, "[SendOrderToDeliveryProviderHandler] Error when unmarshal event message, detail: ", err)
		return
	}
	log.Infoe(e, "[SendOrderToDeliveryProviderHandler] Success to unmarshal event message")
	err := c.orderDeliveryService.Send(e.Context(), event.OrderMessageToEntity(e.PayloadData))
	if err != nil {
		log.Error(e.Context(), "[SendOrderToDeliveryProviderHandler] Error when send order [%d] to delivery service [%v]",
			e.PayloadData.Id, err)
		return
	}
	log.Info(e.Context(), "[SendOrderToDeliveryProviderHandler] Success to send order [%d] to delivery service", e.PayloadData.Id)
}

func (c SendOrderToDeliveryProviderHandler) Close() {
	//
}
