package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gitlab.id.vin/vincart/golib"
	"gitlab.id.vin/vincart/golib-data"
	"gitlab.id.vin/vincart/golib-message-bus"
	"gitlab.id.vin/vincart/golib-sample-adapter/publisher"
	"gitlab.id.vin/vincart/golib-sample-adapter/repository/mysql"
	"gitlab.id.vin/vincart/golib-sample-adapter/service"
	"gitlab.id.vin/vincart/golib-sample-core/usecase"
	"gitlab.id.vin/vincart/golib-sample-public/controller"
	"gitlab.id.vin/vincart/golib-sample-public/properties"
	"gitlab.id.vin/vincart/golib-sample-public/router"
	"gitlab.id.vin/vincart/golib-security"
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

		// Provide gin engine, register core handlers,
		// actuator endpoints and application routers
		fx.Provide(gin.New),
		fx.Invoke(router.RegisterHandlers),
		fx.Invoke(router.RegisterGinRouters),
	}
}
