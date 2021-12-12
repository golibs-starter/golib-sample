package bootstrap

import (
    "github.com/gin-gonic/gin"
    "gitlab.com/golibs-starter/golib"
    "gitlab.com/golibs-starter/golib-sample-adapter/http/client"
    adapterProps "gitlab.com/golibs-starter/golib-sample-adapter/properties"
    "gitlab.com/golibs-starter/golib-sample-adapter/service"
    "gitlab.com/golibs-starter/golib-sample-core/usecase"
    "gitlab.com/golibs-starter/golib-sample-worker/handler"
    "gitlab.com/golibs-starter/golib-sample-worker/router"
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
