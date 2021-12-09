package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gitlab.id.vin/vincart/golib"
	"gitlab.id.vin/vincart/golib-message-bus"
	"gitlab.id.vin/vincart/golib-sample-adapter/http/client"
	adapterProps "gitlab.id.vin/vincart/golib-sample-adapter/properties"
	"gitlab.id.vin/vincart/golib-sample-adapter/service"
	"gitlab.id.vin/vincart/golib-sample-core/usecase"
	"gitlab.id.vin/vincart/golib-sample-worker/handler"
	"gitlab.id.vin/vincart/golib-sample-worker/router"
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

		// Provide datasource, message queue auto config
		golibmsg.KafkaCommonOpt(),
		golibmsg.KafkaConsumerOpt(),

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

		// Provide gin engine, register core handlers,
		// actuator endpoints and application routers
		fx.Provide(gin.New),
		fx.Invoke(router.RegisterHandlers),
		fx.Invoke(router.RegisterGinRouters),
	}
}
