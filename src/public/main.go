package main

import (
	"gitlab.com/golibs-starter/golib-sample-public/bootstrap"
	"go.uber.org/fx"
)

// @title Sample Public API
// @version 1.0.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	fx.New(fx.Options(bootstrap.All()...)).Run()
}
