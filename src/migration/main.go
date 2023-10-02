package main

import (
	"context"
	"github.com/golibs-starter/golib"
	golibdata "github.com/golibs-starter/golib-data"
	golibmigrate "github.com/golibs-starter/golib-migrate"
	"github.com/golibs-starter/golib/log"
	"go.uber.org/fx"
)

func main() {
	if err := fx.New(
		golib.AppOpt(),
		golib.PropertiesOpt(),
		golib.LoggingOpt(),
		golibdata.DatasourceOpt(),
		golibmigrate.MigrationOpt(),
	).Start(context.Background()); err != nil {
		log.Fatal("Error when migrate database: ", err)
	}
}
