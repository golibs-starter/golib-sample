package testing

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/golibs-starter/golib-message-bus"
	"gitlab.com/golibs-starter/golib-message-bus/testutil"
	"gitlab.com/golibs-starter/golib-sample-worker/bootstrap"
	"gitlab.com/golibs-starter/golib-test"
	"gitlab.com/golibs-starter/golib/log"
	"os"
)

type TestSuite struct {
	golibtest.FxTestSuite
}

func (s *TestSuite) SetupSuite() {
	log.Info("Test App is initializing")
	_ = os.Setenv("TZ", "UTC")
	gin.DefaultWriter = log.NewTestingWriter(s.T())
	s.Profile("testing")
	s.ProfilePath("../config/", "./config/")
	s.Option(golibmsg.KafkaAdminOpt())
	s.Option(golibmsgTestUtil.ResetKafkaConsumerGroupOpt())
	s.Option(golibmsgTestUtil.MessageCollectorOpt())
	s.Options(bootstrap.All())
	s.Option(golibmsg.KafkaProducerOpt())
	s.Option(golibmsg.KafkaConsumerReadyWaitOpt())
	s.SetupApp()
	log.Info("Test App is initialized")
}
