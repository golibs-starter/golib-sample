package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gitlab.com/golibs-starter/golib"
	"gitlab.com/golibs-starter/golib-sample-public/controller"
	"gitlab.com/golibs-starter/golib-sample-public/docs"
	"gitlab.com/golibs-starter/golib-sample-public/properties"
	"gitlab.com/golibs-starter/golib/web/actuator"
	"go.uber.org/fx"
)

type RegisterRoutersIn struct {
	fx.In
	App          *golib.App
	Engine       *gin.Engine
	SwaggerProps *properties.SwaggerProperties

	Actuator        *actuator.Endpoint
	OrderController *controller.OrderController
}

func RegisterGinRouters(p RegisterRoutersIn) {
	group := p.Engine.Group(p.App.Path())
	group.GET("/actuator/health", gin.WrapF(p.Actuator.Health))
	group.GET("/actuator/info", gin.WrapF(p.Actuator.Info))

	if p.SwaggerProps.Enabled {
		docs.SwaggerInfo.BasePath = p.App.Path()
		group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	group.GET("/v1/orders/:id", p.OrderController.GetOrder)
	group.POST("/v1/orders", p.OrderController.CreateOrder)
}
