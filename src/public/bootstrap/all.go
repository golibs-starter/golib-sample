package bootstrap

import (
	"github.com/golibs-starter/golib"
	golibdata "github.com/golibs-starter/golib-data"
	golibgin "github.com/golibs-starter/golib-gin"
	golibmsg "github.com/golibs-starter/golib-message-bus"
	"github.com/golibs-starter/golib-sample-adapter/publisher"
	"github.com/golibs-starter/golib-sample-adapter/repository/mysql"
	"github.com/golibs-starter/golib-sample-adapter/service"
	"github.com/golibs-starter/golib-sample-core/usecase"
	"github.com/golibs-starter/golib-sample-public/controller"
	"github.com/golibs-starter/golib-sample-public/properties"
	"github.com/golibs-starter/golib-sample-public/router"
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

		// Http security auto config and authentication filters
		golibsec.HttpSecurityOpt(),
		golibsec.JwtAuthFilterOpt(),

		// Provide datasource auto config
		golibdata.RedisOpt(),
		golibdata.DatasourceOpt(),
		golibmsg.KafkaCommonOpt(),
		golibmsg.KafkaAdminOpt(),
		golibmsg.KafkaProducerOpt(),

		// Provide all application properties
		golib.ProvideProps(properties.NewSwaggerProperties),

		// Provide port's implements
		fx.Provide(publisher.NewEventPublisherAdapter),
		fx.Provide(mysql.NewOrderRepositoryAdapter),

		// Provide use cases
		fx.Provide(usecase.NewGetStatusUseCase),
		fx.Provide(usecase.NewGetOrderUseCase),
		fx.Provide(usecase.NewCreateOrderUseCase),

		// Provide services
		fx.Provide(service.NewStatusService),
		fx.Provide(service.NewOrderService),

		// Provide controllers, these controllers will be used
		// when register router was invoked
		fx.Provide(controller.NewOrderController),

		// Provide gin http server auto config,
		// actuator endpoints and application routers
		golibgin.GinHttpServerOpt(),
		fx.Invoke(router.RegisterGinRouters),

		// Graceful shutdown.
		// OnStop hooks will run in reverse order.
		golibgin.OnStopHttpServerOpt(),
		golibmsg.OnStopProducerOpt(),
	)
}
