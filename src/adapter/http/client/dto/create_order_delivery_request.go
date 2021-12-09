package dto

import (
	"gitlab.id.vin/vincart/golib-sample-core/entity"
	"strconv"
)

type CreateOrderDeliveryRequest struct {
	CustomerId  string `json:"customer_id"`
	InvoiceNo   string `json:"invoice_no"`
	TotalAmount int64  `json:"total_amount"`
}

func NewCreateOrderDeliveryRequest(order *entity.Order) *CreateOrderDeliveryRequest {
	return &CreateOrderDeliveryRequest{
		CustomerId:  order.UserId,
		InvoiceNo:   strconv.Itoa(order.Id),
		TotalAmount: order.TotalAmount,
	}
}
