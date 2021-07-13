package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.id.vin/vincart/golib"
	"gitlab.id.vin/vincart/golib-gin"
	"gitlab.id.vin/vincart/golib-sample/event"
	"gitlab.id.vin/vincart/golib/pubsub"
	"gitlab.id.vin/vincart/golib/web/logging/logc"
)

func main() {
	app := golib.Init(golib.Options{})

	r := gin.New()
	r.Use(golibgin.WrapAll(app.Middleware)...)

	r.GET("/200", func(context *gin.Context) {
		pubsub.Publish(event.NewOrderCreatedEvent(context, event.OrderCreatedPayload{
			Code:        "VMM1234",
			TotalAmount: 15000,
		}))
		logc.Info(context, "Test log success")
	})

	r.GET("/400", func(context *gin.Context) {
		logc.Error(context, "Test log error")
		context.JSON(400, nil)
	})

	// Start HTTP Server
	_ = r.Run(fmt.Sprintf(":%d", 8080))
}
