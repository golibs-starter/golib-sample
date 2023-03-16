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
	golibtest.RequireFxApp(
		golib.ProvidePropsOption(golib.WithActiveProfiles([]string{"testing"})),
		golib.ProvidePropsOption(golib.WithPaths([]string{"../config/", "./config/"})),
		golibmigrate.MigrationOpt(),
		golibtest.EnableWebTestUtil(),
		fx.Populate(&db),
		bootstrap.All(),
	)
	log.Info("Test App is initialized")
}
