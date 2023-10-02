package testing

import (
	"github.com/golibs-starter/golib"
	golibmigrate "github.com/golibs-starter/golib-migrate"
	"github.com/golibs-starter/golib-sample-internal/bootstrap"
	golibtest "github.com/golibs-starter/golib-test"
	"github.com/golibs-starter/golib/log"
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
