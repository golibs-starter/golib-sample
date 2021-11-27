package main

import (
	"context"
	"gitlab.com/golibs-starter/golib"
	"gitlab.com/golibs-starter/golib-data"
	"gitlab.com/golibs-starter/golib-migrate"
	"gitlab.com/golibs-starter/golib/log"
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
