package testing

import (
	"gitlab.com/golibs-starter/golib"
	"gitlab.com/golibs-starter/golib-migrate"
	"gitlab.com/golibs-starter/golib-sample-internal/bootstrap"
	"gitlab.com/golibs-starter/golib-test"
	"gitlab.com/golibs-starter/golib/log"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"os"
)

var (
	db *gorm.DB
)

func init() {
	log.Info("Test App is initializing")
	_ = os.Setenv("TZ", "UTC")
	_ = golibtest.RequireFxApp(
		golib.ProvidePropsOption(golib.WithPaths([]string{"../config/", "./config/"})),
		golib.ProvidePropsOption(golib.WithActiveProfiles([]string{"testing"})),
		bootstrap.All(),
		golibmigrate.MigrationOpt(),
		fx.Populate(&db),
	)
	log.Info("Test App is initialized")
}
