package api

import (
	"context"
	assert "github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	golibmsg "gitlab.id.vin/vincart/golib-message-bus"
	"gitlab.id.vin/vincart/golib-sample-adapter/repository/mysql/model"
	"gitlab.id.vin/vincart/golib-sample-core/event"
	"gitlab.id.vin/vincart/golib-sample-internal/testing/base"
	"gitlab.id.vin/vincart/golib-sample-internal/testing/handler"
	"gitlab.id.vin/vincart/golib-test"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"net/http"
	"testing"
	"time"
)

type CreateOrderControllerTest struct {
	*base.TestSuite
	authorization  string
	db             *gorm.DB
	dummyCollector *handler.OrderEventDummyCollector
}

func TestCreateOrderControllerTest(t *testing.T) {
	// internal_service:secret
	s := CreateOrderControllerTest{authorization: "aW50ZXJuYWxfc2VydmljZTpzZWNyZXQ="}
	s.TestSuite = base.NewTestSuite(
		golibtest.WithTestingDir(".."),
		golibtest.WithFxOption(golibmsg.KafkaConsumerOpt()),
		golibtest.WithFxOption(golibmsg.ProvideConsumer(handler.NewOrderCreatedEventDummyHandler)),
		golibtest.WithFxOption(fx.Provide(handler.NewOrderEventDummyCollector)),
		golibtest.WithFxPopulate(&s.db, &s.dummyCollector),
		golibtest.WithFxOption(fx.Invoke(func(db *gorm.DB) {
			db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Order{})
		})),
	)
	suite.Run(t, &s)
}

func (s CreateOrderControllerTest) TestCreateOrder_GiveValidBody_WhenRepoResponseSuccess_ShouldReturnSuccess() {
	resp := golibtest.NewDefaultHttpClient(s.T()).Do(
		golibtest.NewRequestBuilder(s.T()).
			WithBasicAuthorization(s.authorization).
			WithBodyString(`{"total_amount": 85000}`).
			Post(s.URL("/v1/orders")),
	)
	defer resp.Body.Close()

	var actualOrder model.Order
	assert.NoError(s.T(), s.db.First(&actualOrder).Error)
	assert.NotNil(s.T(), actualOrder)
	assert.EqualValues(s.T(), 85000, actualOrder.TotalAmount)

	golibtest.NewRestAssured(s.T(), resp).
		Status(http.StatusOK).
		Body("data.id", actualOrder.Id).
		Body("data.total_amount", actualOrder.TotalAmount).
		BodyCb("data.created_at", func(value interface{}) {
			assert.InDelta(s.T(), actualOrder.CreatedAt.Unix(), value, 1)
		})

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	for {
		if len(s.dummyCollector.CreatedEvents()) >= 1 || ctx.Err() != nil {
			break
		}
	}
	assert.Len(s.T(), s.dummyCollector.CreatedEvents(), 1)
	actualEvent := s.dummyCollector.CreatedEvents()[0]
	assert.Equal(s.T(), "OrderCreatedEvent", actualEvent.Name())
	assert.IsType(s.T(), &event.OrderMessage{}, actualEvent.Payload())
	actualEventPayload := actualEvent.Payload().(*event.OrderMessage)
	assert.EqualValues(s.T(), actualOrder.Id, actualEventPayload.Id)
	assert.EqualValues(s.T(), actualOrder.TotalAmount, actualEventPayload.TotalAmount)
	assert.InDelta(s.T(), actualOrder.CreatedAt.Unix(), actualEventPayload.CreatedAt, 1)
}
