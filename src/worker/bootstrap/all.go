package bootstrap

import (
	"github.com/golibs-starter/golib"
	golibcron "github.com/golibs-starter/golib-cron"
	golibgin "github.com/golibs-starter/golib-gin"
	golibmsg "github.com/golibs-starter/golib-message-bus"
	"github.com/golibs-starter/golib-sample-adapter/http/client"
	adapterProps "github.com/golibs-starter/golib-sample-adapter/properties"
	"github.com/golibs-starter/golib-sample-adapter/service"
	"github.com/golibs-starter/golib-sample-core/usecase"
	"github.com/golibs-starter/golib-sample-worker/handler"
	"github.com/golibs-starter/golib-sample-worker/job"
	"github.com/golibs-starter/golib-sample-worker/router"
	golibsec "github.com/golibs-starter/golib-security"
	"go.uber.org/fx"
)

func All() fx.Option {
	return fx.Options(
		golib.AppOpt(),
		golib.PropertiesOpt(),
		golib.LoggingOpt(),
		golib.EventOpt(),
		golib.BuildInfoOpt(Version, CommitHash, BuildTime),
		golib.ActuatorEndpointOpt(),
		golib.HttpRequestLogOpt(),

		// Provide datasource, message queue auto config
		golibmsg.KafkaCommonOpt(),
		golibmsg.KafkaConsumerOpt(),
		golibmsg.KafkaProducerOpt(),

		// Provide http client auto config with contextual http client by default,
		// Besides, provide an additional wrapper to easy to control security.
		golib.HttpClientOpt(),
		golibsec.SecuredHttpClientOpt(),

		// Provide all application properties
		golib.ProvideProps(adapterProps.NewDeliveryServiceProperties),

		// Provide port's implements
		fx.Provide(client.NewDeliveryServiceAdapter),

		// Provide use cases
		fx.Provide(usecase.NewSendOrderToDeliveryProviderUseCase),

		// Provide services
		fx.Provide(service.NewOrderDeliveryService),

		// Provide handlers
		golibmsg.ProvideConsumer(handler.NewSendOrderToDeliveryProviderHandler),

		// Provide cron jobs
		golibcron.Opt(),
		golibcron.ProvideJob(job.NewYourFirstCronJob),
		golibcron.ProvideJob(job.NewYourSecondCronJob),

		// Provide gin engine, register core handlers,
		// actuator endpoints and application routers
		golibgin.GinHttpServerOpt(),
		fx.Invoke(router.RegisterGinRouters),

		// Graceful shutdown.
		// OnStop hooks will run in reverse order.
		golibgin.OnStopHttpServerOpt(),
		golibmsg.OnStopProducerOpt(),
		golibmsg.OnStopConsumerOpt(),
		golibcron.OnStopHookOpt(),
	)
}
