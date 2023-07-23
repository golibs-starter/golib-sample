package usecase

import (
	"context"
	"gitlab.com/golibs-starter/golib-sample-core/entity"
	"gitlab.com/golibs-starter/golib-sample-core/port"
	"gitlab.com/golibs-starter/golib/log"
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
		log.WithCtx(ctx).WithError(err).Errorf("Cannot send order [%d] to delivery service", order.Id)
		return err
	}
	log.WithCtx(ctx).Infof("Send order [%d] to delivery service success, trackingId [%s]", order.Id, trackingId)
	return nil
}
