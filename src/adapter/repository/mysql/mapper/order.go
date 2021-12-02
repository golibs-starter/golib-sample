package mapper

import (
	"gitlab.id.vin/vincart/golib-sample-adapter/repository/mysql/model"
	"gitlab.id.vin/vincart/golib-sample-core/entity"
)

func ModelToOrderEntity(order *model.Order) *entity.Order {
	return &entity.Order{
		Id:          order.Id,
		UserId:      order.UserId,
		TotalAmount: order.TotalAmount,
		CreatedAt:   order.CreatedAt,
	}
}
