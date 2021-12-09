package dto

import (
	"gitlab.id.vin/vincart/golib/web/response"
)

type OrderDeliveryResponseDto struct {
	Meta response.Meta     `json:"meta"`
	Data *OrderDeliveryDto `json:"data"`
}

type OrderDeliveryDto struct {
	Id        string `json:"id"`
	CreatedAt int64  `json:"created_at"`
}
