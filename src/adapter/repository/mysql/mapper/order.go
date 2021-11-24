package mapper

import (
	"gitlab.id.vin/vincart/golib-sample-adapter/repository/mysql/model"
	"gitlab.id.vin/vincart/golib-sample-core/entity"
)

func ModelToOrder(order *model.Order) *entity.Order {
	return &entity.Order{
		Id:          order.Id,
		TotalAmount: order.TotalAmount,
		CreatedAt:   order.CreatedAt,
	}
}
