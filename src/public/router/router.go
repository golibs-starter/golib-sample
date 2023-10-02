package router

import (
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib"
	"github.com/golibs-starter/golib-sample-public/controller"
	"github.com/golibs-starter/golib-sample-public/docs"
	"github.com/golibs-starter/golib-sample-public/properties"
	"github.com/golibs-starter/golib/web/actuator"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
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
