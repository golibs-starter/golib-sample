package resource

import "gitlab.id.vin/vincart/golib-sample-core/entity"

type Order struct {
	Id          int   `json:"id"`
	TotalAmount int64 `json:"total_amount"`
	CreatedAt   int64 `json:"created_at"`
}

func NewOrder(entity *entity.Order) *Order {
	return &Order{
		Id:          entity.Id,
		TotalAmount: entity.TotalAmount,
		CreatedAt:   entity.CreatedAt.Unix(),
	}
}
