package service

import (
	"context"
	"gitlab.id.vin/vincart/golib-sample-core/entity"
	"gitlab.id.vin/vincart/golib-sample-core/usecase"
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
