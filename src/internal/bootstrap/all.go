package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/golibs-starter/golib"
	"gitlab.com/golibs-starter/golib-data"
	"gitlab.com/golibs-starter/golib-message-bus"
	adapterProps "gitlab.com/golibs-starter/golib-sample-adapter/properties"
	"gitlab.com/golibs-starter/golib-sample-adapter/publisher"
	"gitlab.com/golibs-starter/golib-sample-adapter/repository/mysql"
	"gitlab.com/golibs-starter/golib-sample-adapter/service"
	"gitlab.com/golibs-starter/golib-sample-core/usecase"
	"gitlab.com/golibs-starter/golib-sample-internal/controller"
	"gitlab.com/golibs-starter/golib-sample-internal/properties"
	"gitlab.com/golibs-starter/golib-sample-internal/router"
	"gitlab.com/golibs-starter/golib-security"
	"go.uber.org/fx"
)

func All() []fx.Option {
	return []fx.Option{
		golib.AppOpt(),
		golib.PropertiesOpt(),
		golib.LoggingOpt(),
		golib.EventOpt(),
		golib.BuildInfoOpt(Version, CommitHash, BuildTime),
		golib.ActuatorEndpointOpt(),

		// Http security auto config and authentication filters
		golibsec.HttpSecurityOpt(),
		golibsec.BasicAuthOpt(),

		// Provide datasource auto config
		golibdata.RedisOpt(),
		golibdata.DatasourceOpt(),
		golibmsg.KafkaCommonOpt(),
		golibmsg.KafkaAdminOpt(),
		golibmsg.KafkaProducerOpt(),

		// Provide http client auto config with contextual http client by default,
		// Besides, provide an additional wrapper to easy to control security.
		golib.HttpClientOpt(),
		golibsec.SecuredHttpClientOpt(),

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
		fx.Provide(controller.NewStatusController),
		fx.Provide(controller.NewOrderController),

		// Provide gin engine, register core handlers,
		// actuator endpoints and application routers
		fx.Provide(gin.New),
		fx.Invoke(router.RegisterHandlers),
		fx.Invoke(router.RegisterGinRouters),
	}
}
