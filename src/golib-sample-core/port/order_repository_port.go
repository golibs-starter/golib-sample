package port

import (
	"context"
	"gitlab.id.vin/vincart/golib-sample-core/entity"
	"gitlab.id.vin/vincart/golib-sample-core/entity/request"
)

type OrderRepository interface {
	FindById(ctx context.Context, id int64) (*entity.Order, error)

	CreateOrder(ctx context.Context, req *request.CreateOrderRequest) (*entity.Order, error)
}
