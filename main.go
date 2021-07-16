package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.id.vin/vincart/golib"
	"gitlab.id.vin/vincart/golib-gin"
	"gitlab.id.vin/vincart/golib-sample/event"
	"gitlab.id.vin/vincart/golib-security"
	"gitlab.id.vin/vincart/golib/config"
	"gitlab.id.vin/vincart/golib/pubsub"
	"gitlab.id.vin/vincart/golib/web/log"
)

func main() {
	app := golib.New(
		golib.WithConfigLoader(config.Option{}),
		golib.WithLogger(),
		golib.WithEventBus(map[pubsub.Event][]pubsub.Subscriber{}),
		golib.WithHttpClientAutoConfig(golibsec.SecuredHttpClientWrapper()),
	)

	r := gin.New()
	r.Use(golibgin.WrapAll(app.Middleware())...)

	r.GET("/200", func(context *gin.Context) {
		pubsub.Publish(event.NewOrderCreatedEvent(context, event.OrderCreatedPayload{
			Code:        "VMM1234",
			TotalAmount: 15000,
		}))
		_, err := app.HttpClient.Get("https://api-qc.vinid.dev/vmm-order/v1/orders", nil)
		if err != nil {
			log.Error(context, "cannot request to vmm order with error [%v]", err)
			return
		}
		log.Info(context, "Test log success")
	})

	r.GET("/400", func(context *gin.Context) {
		log.Error(context, "Test log error")
		context.JSON(400, nil)
	})

	// Start HTTP Server
	_ = r.Run(fmt.Sprintf(":%d", 8080))
}
