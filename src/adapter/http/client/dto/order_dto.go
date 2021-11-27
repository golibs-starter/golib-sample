package dto

import (
	"gitlab.com/golibs-starter/golib-sample-core/entity"
	"gitlab.com/golibs-starter/golib/web/response"
	"time"
)

type OrderResponseDto struct {
	Meta response.Meta `json:"meta"`
	Data *OrderDto     `json:"data"`
}

type OrderDto struct {
	Id          int   `json:"id"`
	TotalAmount int64 `json:"total_amount"`
	CreatedAt   int64 `json:"created_at"`
}

func (o OrderDto) ToEntity() *entity.Order {
	return &entity.Order{
		Id:          o.Id,
		TotalAmount: o.TotalAmount,
		CreatedAt:   time.Unix(o.CreatedAt, 0),
	}
}
