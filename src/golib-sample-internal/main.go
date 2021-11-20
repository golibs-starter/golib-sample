package main

import (
	"gitlab.id.vin/vincart/golib-gin"
	"gitlab.id.vin/vincart/golib-sample-internal/bootstrap"
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
