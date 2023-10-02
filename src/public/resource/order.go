package resource

import "github.com/golibs-starter/golib-sample-core/entity"

type Order struct {
	Id          int    `json:"id"`
	UserId      string `json:"user_id"`
	TotalAmount int64  `json:"total_amount"`
	CreatedAt   int64  `json:"created_at"`
}

func NewOrder(entity *entity.Order) *Order {
	return &Order{
		Id:          entity.Id,
		UserId:      entity.UserId,
		TotalAmount: entity.TotalAmount,
		CreatedAt:   entity.CreatedAt.Unix(),
	}
}
