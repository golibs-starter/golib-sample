package main

import (
	"gitlab.id.vin/vincart/golib-gin"
	"gitlab.id.vin/vincart/golib-sample-worker/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(fx.Options(bootstrap.All()...), golibgin.StartOpt()).Run()
}
