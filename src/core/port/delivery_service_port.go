package port

import (
    "context"
    "gitlab.com/golibs-starter/golib-sample-core/entity"
)

type DeliveryService interface {
    CreateOrder(ctx context.Context, order *entity.Order) (string, error)
}
