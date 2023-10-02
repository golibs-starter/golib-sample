package testing

import (
	"github.com/golibs-starter/golib"
	golibmsg "github.com/golibs-starter/golib-message-bus"
	golibmsgTestUtil "github.com/golibs-starter/golib-message-bus/testutil"
	"github.com/golibs-starter/golib-sample-worker/bootstrap"
	golibtest "github.com/golibs-starter/golib-test"
	"github.com/golibs-starter/golib/log"
	"github.com/jarcoal/httpmock"
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
		golibmsgTestUtil.ResetKafkaConsumerGroupOpt(),
		golibmsg.KafkaProducerOpt(),
		golibmsgTestUtil.MessageCollectorOpt(),
		fx.Populate(&messageCollector),
		fx.Decorate(func(httpClient *http.Client) *http.Client {
			httpmock.ActivateNonDefault(httpClient)
			return httpClient
		}),
		golibtest.EnableWebTestUtil(),
		bootstrap.All(),
		golibmsg.KafkaConsumerReadyWaitOpt(),
	)
	log.Info("Test App is initialized")
}
