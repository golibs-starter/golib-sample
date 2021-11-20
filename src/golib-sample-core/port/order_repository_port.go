package port

import (
	"context"
	"gitlab.id.vin/vincart/golib-sample-core/entity"
)

type OrderRepositoryPort interface {
	FindById(ctx context.Context, id int64) (*entity.Order, error)
}
