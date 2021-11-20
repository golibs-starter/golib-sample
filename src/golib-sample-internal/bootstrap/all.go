package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gitlab.id.vin/vincart/golib"
	"gitlab.id.vin/vincart/golib-sample-adapter/http/client"
	adapterProps "gitlab.id.vin/vincart/golib-sample-adapter/properties"
	"gitlab.id.vin/vincart/golib-sample-adapter/service"
	"gitlab.id.vin/vincart/golib-sample-core/usecase"
	"gitlab.id.vin/vincart/golib-sample-internal/controller"
	"gitlab.id.vin/vincart/golib-sample-internal/properties"
	"gitlab.id.vin/vincart/golib-sample-internal/router"
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
		golibsec.BasicAuthOpt(),

		// Provide datasource auto config
		//golibdata.RedisOpt(),

		// Provide http client auto config with contextual http client by default,
		// Besides, provide an additional wrapper to easy to control security.
		golib.HttpClientOpt(),
		golibsec.SecuredHttpClientOpt(),

		// Provide all application properties
		golib.ProvideProps(properties.NewSwaggerProperties),
		golib.ProvideProps(adapterProps.NewOrderRepositoryProperties),

		// Provide port's implements
		fx.Provide(client.NewOrderRepositoryAdapter),

		// Provide use cases
		fx.Provide(usecase.NewGetStatusUseCase),
		fx.Provide(usecase.NewGetOrderUseCase),

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
