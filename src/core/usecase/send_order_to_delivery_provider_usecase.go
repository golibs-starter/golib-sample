package usecase

import (
	"context"
	"gitlab.id.vin/vincart/golib-sample-core/entity"
	"gitlab.id.vin/vincart/golib-sample-core/port"
	"gitlab.id.vin/vincart/golib/web/log"
)

type SendOrderToDeliveryProviderUseCase struct {
	deliveryService port.DeliveryService
}

func NewSendOrderToDeliveryProviderUseCase(deliveryService port.DeliveryService) *SendOrderToDeliveryProviderUseCase {
	return &SendOrderToDeliveryProviderUseCase{deliveryService: deliveryService}
}

func (c SendOrderToDeliveryProviderUseCase) Send(ctx context.Context, order *entity.Order) error {
	trackingId, err := c.deliveryService.CreateOrder(ctx, order)
	if err != nil {
		log.Error(ctx, "Cannot send order [%d] to delivery service due by error [%v]", order.Id, err)
		return err
	}
	log.Info(ctx, "Send order [%d] to delivery service success, trackingId [%s]", order.Id, trackingId)
	return nil
}
