package main

import (
	"gitlab.id.vin/vincart/golib-gin"
	"gitlab.id.vin/vincart/golib-sample-public/bootstrap"
	"go.uber.org/fx"
)

// @title Sample Public API
// @version 1.0.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	fx.New(fx.Options(bootstrap.All()...), golibgin.StartOpt()).Run()
}
