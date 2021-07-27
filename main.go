package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.id.vin/vincart/golib"
	"gitlab.id.vin/vincart/golib-gin"
	"gitlab.id.vin/vincart/golib-sample/event"
	"gitlab.id.vin/vincart/golib-security"
	"gitlab.id.vin/vincart/golib/pubsub"
	"gitlab.id.vin/vincart/golib/web/log"
)

func main() {
	app := golib.New(
		golib.WithConfigProperties(),
		golib.WithLoggerAutoConfig(),
		golib.WithEventAutoConfig(map[pubsub.Event][]pubsub.Subscriber{}),
		golib.WithHttpClientAutoConfig(golibsec.SecuredHttpClientWrapper()),
		golibsec.WithHttpSecurityAutoConfig(
			golibsec.WithJwtAuth(),
		),
	)

	r := gin.New()
	r.Use(golibgin.WrapAll(app.Middleware())...)

	r.GET("/200", func(c *gin.Context) {
		pubsub.Publish(event.NewOrderCreatedEvent(c, event.OrderCreatedPayload{
			Code:        "VMM1234",
			TotalAmount: 15000,
		}))
		log.Info(c, "Test log success")
		c.JSON(200, map[string]interface{}{
			"message": "successful",
		})
	})

	r.GET("/400", func(c *gin.Context) {
		_, err := app.HttpClient.Get(c, "https://api-qc.vinid.dev/vmm-order/v1/orders", nil)
		if err != nil {
			log.Error(c, "cannot request to vmm order with error [%v]", err)
			c.JSON(400, map[string]interface{}{
				"message": "not found",
			})
			return
		}
		log.Error(c, "Test log error")
		c.JSON(400, nil)
	})

	// Start HTTP Server
	_ = r.Run(fmt.Sprintf(":%d", 8080))
}
