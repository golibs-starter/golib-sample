package exception

import "gitlab.id.vin/vincart/golib/exception"

var (
	OrderIdInvalid = exception.New(40008001, "Order id is invalid")
	OrderNotFound  = exception.New(40408001, "Order not found")
)
