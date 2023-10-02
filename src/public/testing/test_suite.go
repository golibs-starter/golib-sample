package testing

import (
	"github.com/gin-gonic/gin"
	golibdataTestUtil "github.com/golibs-starter/golib-data/testutil"
	golibmigrate "github.com/golibs-starter/golib-migrate"
	"github.com/golibs-starter/golib-sample-public/bootstrap"
	golibsecTestUtil "github.com/golibs-starter/golib-security/testutil"
	golibtest "github.com/golibs-starter/golib-test"
	"github.com/golibs-starter/golib/log"
	"os"
)

type TestSuite struct {
	golibtest.FxTestSuite
	jwtTestUtil *golibsecTestUtil.JwtTestUtil
}

func (s *TestSuite) SetupSuite() {
	log.Info("Test App is initializing")
	_ = os.Setenv("TZ", "UTC")
	gin.DefaultWriter = log.NewTestingWriter(s.T())
	s.Profile("testing")
	s.ProfilePath("../config/", "./config/")
	s.Option(golibmigrate.MigrationOpt())
	s.Option(golibsecTestUtil.JwtTestUtilOpt())
	s.Option(golibdataTestUtil.EnableDatabaseTestUtilOpt())
	s.Option(golibtest.EnableWebTestUtil())
	s.Option(golibdataTestUtil.TruncateTablesOpt("orders"))
	s.Populate(&s.jwtTestUtil)
	s.Option(bootstrap.All())
	s.StartApp()
	log.Info("Test App is initialized")
}
