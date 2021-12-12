package main

import (
    "gitlab.com/golibs-starter/golib-gin"
    "gitlab.com/golibs-starter/golib-sample-worker/bootstrap"
    "go.uber.org/fx"
)

func main() {
    fx.New(fx.Options(bootstrap.All()...), golibgin.StartOpt()).Run()
}
