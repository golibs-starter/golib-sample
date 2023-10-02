package service

import (
	"context"
	"github.com/golibs-starter/golib-sample-core/entity"
	"github.com/golibs-starter/golib-sample-core/usecase"
)

type OrderDeliveryService struct {
	sendOrderToDeliveryProviderUseCase *usecase.SendOrderToDeliveryProviderUseCase
}

func NewOrderDeliveryService(sendOrderToDeliveryProviderUseCase *usecase.SendOrderToDeliveryProviderUseCase) *OrderDeliveryService {
	return &OrderDeliveryService{sendOrderToDeliveryProviderUseCase: sendOrderToDeliveryProviderUseCase}
}

func (o OrderDeliveryService) Send(ctx context.Context, order *entity.Order) error {
	return o.sendOrderToDeliveryProviderUseCase.Send(ctx, order)
}
