package port

import (
	"context"
    "gitlab.com/golibs-starter/golib-sample-core/entity"
    "gitlab.com/golibs-starter/golib-sample-core/entity/request"
)

type OrderRepository interface {
	FindById(ctx context.Context, id int64) (*entity.Order, error)

	CreateOrder(ctx context.Context, req *request.CreateOrderRequest) (*entity.Order, error)
}
