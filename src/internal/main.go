package main

import (
	"gitlab.com/golibs-starter/golib-gin"
	"gitlab.com/golibs-starter/golib-sample-internal/bootstrap"
	"go.uber.org/fx"
)

// @title Sample internal API
// @version 1.0.0
// @securityDefinitions.basic BasicAuth
// @in header
// @name Authorization
func main() {
	fx.New(fx.Options(bootstrap.All()...), golibgin.StartOpt()).Run()
}
