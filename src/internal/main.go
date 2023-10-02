package main

import (
	"github.com/golibs-starter/golib-sample-internal/bootstrap"
	"go.uber.org/fx"
)

// @title Sample internal API
// @version 1.0.0
// @securityDefinitions.basic BasicAuth
// @in header
// @name Authorization
func main() {
	fx.New(bootstrap.All()).Run()
}
