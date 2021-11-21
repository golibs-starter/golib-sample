package exception

import "gitlab.id.vin/vincart/golib/exception"

var (
	StatusInvalid = exception.New(40008000, "Status code is invalid")
)
