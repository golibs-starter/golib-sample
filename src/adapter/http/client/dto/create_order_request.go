package dto

import (
	"gitlab.id.vin/vincart/golib-sample-core/entity/request"
)

type CreateOrderRequest struct {
	TotalAmount int64 `json:"total_amount"`
}

func NewCreateOrderRequest(req *request.CreateOrderRequest) *CreateOrderRequest {
	return &CreateOrderRequest{TotalAmount: req.TotalAmount}
}
