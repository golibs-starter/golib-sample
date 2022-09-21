package testing

import (
	golibmsg "gitlab.com/golibs-starter/golib-message-bus"
	"gitlab.com/golibs-starter/golib-sample-worker/bootstrap"
	"gitlab.com/golibs-starter/golib-test"
	"gitlab.com/golibs-starter/golib/log"
)

type TestSuite struct {
	golibtest.FxTestSuite
}

func (s *TestSuite) SetupSuite() {
	log.Info("Test App is initializing")
	s.Profile("testing")
	s.ProfilePath("../config/", "./config/")
	s.Options(bootstrap.All())
	s.Option(golibmsg.KafkaConsumerReadyWaitOpt())
	s.SetupApp()
	log.Info("Test App is initialized")
}
