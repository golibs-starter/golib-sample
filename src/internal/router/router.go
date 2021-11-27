package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gitlab.com/golibs-starter/golib"
	"gitlab.com/golibs-starter/golib-gin"
	"gitlab.com/golibs-starter/golib-sample-internal/controller"
	"gitlab.com/golibs-starter/golib-sample-internal/docs"
	"gitlab.com/golibs-starter/golib-sample-internal/properties"
	"gitlab.com/golibs-starter/golib/web/actuator"
	"go.uber.org/fx"
)

type RegisterRoutersIn struct {
	fx.In
	App          *golib.App
	Engine       *gin.Engine
	SwaggerProps *properties.SwaggerProperties

	Actuator         *actuator.Endpoint
	StatusController *controller.StatusController
	OrderController  *controller.OrderController
}

func RegisterHandlers(app *golib.App, engine *gin.Engine) {
	engine.Use(golibgin.InitContext())
	engine.Use(golibgin.WrapAll(app.Handlers())...)
}

func RegisterGinRouters(p RegisterRoutersIn) {
	group := p.Engine.Group(p.App.Path())
	group.GET("/actuator/health", gin.WrapF(p.Actuator.Health))
	group.GET("/actuator/info", gin.WrapF(p.Actuator.Info))

	if p.SwaggerProps.Enabled {
		docs.SwaggerInfo.BasePath = p.App.Path()
		group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	group.GET("/v1/statuses/:code", p.StatusController.ReturnStatus)
	group.GET("/v1/orders/:id", p.OrderController.GetOrder)
	group.POST("/v1/orders", p.OrderController.CreateOrder)
}
