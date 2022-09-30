package testing

import (
	"github.com/jarcoal/httpmock"
	"gitlab.com/golibs-starter/golib"
	"gitlab.com/golibs-starter/golib-message-bus"
	"gitlab.com/golibs-starter/golib-message-bus/testutil"
	"gitlab.com/golibs-starter/golib-sample-worker/bootstrap"
	"gitlab.com/golibs-starter/golib-test"
	"gitlab.com/golibs-starter/golib/log"
	"go.uber.org/fx"
	"net/http"
	"os"
)

var (
	messageCollector *golibmsgTestUtil.MessageCollector
)

func init() {
	log.Info("Test App is initializing")
	_ = os.Setenv("TZ", "UTC")
	golibtest.RequireFxApp(
		golib.ProvidePropsOption(golib.WithActiveProfiles([]string{"testing"})),
		golib.ProvidePropsOption(golib.WithPaths([]string{"../config/", "./config/"})),
		golibmsg.KafkaAdminOpt(),
		golibmsg.KafkaProducerOpt(),
		golibmsgTestUtil.ResetKafkaConsumerGroupOpt(),
		golibmsgTestUtil.MessageCollectorOpt(),
		bootstrap.All(),
		golibmsg.KafkaConsumerReadyWaitOpt(),
		fx.Decorate(func(httpClient *http.Client) *http.Client {
			httpmock.ActivateNonDefault(httpClient)
			return httpClient
		}),
		fx.Populate(&messageCollector),
	)
	log.Info("Test App is initialized")
}
