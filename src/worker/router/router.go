package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.id.vin/vincart/golib"
	"gitlab.id.vin/vincart/golib-gin"
	"gitlab.id.vin/vincart/golib/web/actuator"
	"go.uber.org/fx"
)

type RegisterRoutersIn struct {
	fx.In
	App    *golib.App
	Engine *gin.Engine

	Actuator *actuator.Endpoint
}

func RegisterHandlers(app *golib.App, engine *gin.Engine) {
	engine.Use(golibgin.InitContext())
	engine.Use(golibgin.WrapAll(app.Handlers())...)
}

func RegisterGinRouters(p RegisterRoutersIn) {
	group := p.Engine.Group(p.App.Path())
	group.GET("/actuator/health", gin.WrapF(p.Actuator.Health))
	group.GET("/actuator/info", gin.WrapF(p.Actuator.Info))
}
