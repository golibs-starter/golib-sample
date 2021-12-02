package resource

import "gitlab.com/golibs-starter/golib-sample-core/entity/request"

type CreateOrderRequest struct {
	TotalAmount int64 `json:"total_amount"`
}

func (c CreateOrderRequest) ToEntity() *request.CreateOrderRequest {
	return &request.CreateOrderRequest{
		TotalAmount: c.TotalAmount,
	}
}
