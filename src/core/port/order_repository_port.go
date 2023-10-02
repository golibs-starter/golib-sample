package port

import (
	"context"
	"github.com/golibs-starter/golib-sample-core/entity"
	"github.com/golibs-starter/golib-sample-core/entity/request"
)

type OrderRepository interface {
	FindById(ctx context.Context, id int64) (*entity.Order, error)

	CreateOrder(ctx context.Context, req *request.CreateOrderRequest) (*entity.Order, error)
}
