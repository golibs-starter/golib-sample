package handler

import (
	"gitlab.com/golibs-starter/golib-message-bus/kafka/core"
	"gitlab.com/golibs-starter/golib-message-bus/kafka/relayer"
	"gitlab.com/golibs-starter/golib-sample-adapter/service"
	"gitlab.com/golibs-starter/golib-sample-core/event"
	"gitlab.com/golibs-starter/golib/log"
)

type SendOrderToDeliveryProviderHandler struct {
	orderDeliveryService *service.OrderDeliveryService
	eventConverter       relayer.EventConverter
}

func NewSendOrderToDeliveryProviderHandler(
	orderDeliveryService *service.OrderDeliveryService,
	eventConverter relayer.EventConverter,
) core.ConsumerHandler {
	return &SendOrderToDeliveryProviderHandler{
		orderDeliveryService: orderDeliveryService,
		eventConverter:       eventConverter,
	}
}

func (c *SendOrderToDeliveryProviderHandler) HandlerFunc(msg *core.ConsumerMessage) {
	var e event.OrderCreatedEvent
	if err := c.eventConverter.Restore(msg, &e); err != nil {
		log.Error("[SendOrderToDeliveryProviderHandler] Error when unmarshal event message, detail: ", err)
		return
	}
	logger := log.WithCtx(e.Context())
	logger.Info("[SendOrderToDeliveryProviderHandler] Success to unmarshal event message")
	payload := e.Payload().(*event.OrderMessage)
	err := c.orderDeliveryService.Send(e.Context(), event.OrderMessageToEntity(payload))
	if err != nil {
		logger.WithError(err).
			Error("[SendOrderToDeliveryProviderHandler] Error when send order [%d] to delivery service", payload.Id)
		return
	}
	logger.Infof("[SendOrderToDeliveryProviderHandler] Success to send order [%d] to delivery service", payload.Id)
}

func (c *SendOrderToDeliveryProviderHandler) Close() {
	//
}
