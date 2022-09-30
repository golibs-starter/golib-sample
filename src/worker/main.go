package main

import (
	"gitlab.com/golibs-starter/golib-sample-worker/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(bootstrap.All()).Run()
}
