package port

import (
	"context"
	"gitlab.id.vin/vincart/golib-sample-core/entity"
)

type DeliveryService interface {
	CreateOrder(ctx context.Context, order *entity.Order) (string, error)
}
