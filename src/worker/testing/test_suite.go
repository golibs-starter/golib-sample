package testing

import (
	"github.com/gin-gonic/gin"
	"github.com/jarcoal/httpmock"
	"gitlab.com/golibs-starter/golib-message-bus"
	"gitlab.com/golibs-starter/golib-message-bus/testutil"
	"gitlab.com/golibs-starter/golib-sample-worker/bootstrap"
	"gitlab.com/golibs-starter/golib-test"
	"gitlab.com/golibs-starter/golib/log"
	"net/http"
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
	s.Decorate(func(httpClient *http.Client) *http.Client {
		httpmock.ActivateNonDefault(httpClient)
		log.Info("http mock activated")
		return httpClient
	})
	s.Option(golibmsg.KafkaAdminOpt())
	s.Option(golibmsgTestUtil.ResetKafkaConsumerGroupOpt())
	s.Option(golibmsgTestUtil.MessageCollectorOpt())
	s.Options(bootstrap.All())
	s.Option(golibmsg.KafkaProducerOpt())
	s.Option(golibmsg.KafkaConsumerReadyWaitOpt())
	s.SetupApp()
	log.Info("Test App is initialized")
}

func (s *TestSuite) TearDownSuite() {
	httpmock.DeactivateAndReset()
}
