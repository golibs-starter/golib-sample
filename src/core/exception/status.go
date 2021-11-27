package exception

import "gitlab.com/golibs-starter/golib/exception"

var (
	StatusInvalid = exception.New(40008000, "Status code is invalid")
)
