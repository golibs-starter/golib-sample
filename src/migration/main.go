package main

import (
	"context"
	"gitlab.id.vin/vincart/golib"
	"gitlab.id.vin/vincart/golib-data"
	"gitlab.id.vin/vincart/golib-migrate"
	"gitlab.id.vin/vincart/golib/log"
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
