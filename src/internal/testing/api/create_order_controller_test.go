package api

import (
	assert "github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gitlab.id.vin/vincart/golib-sample-adapter/publisher"
	"gitlab.id.vin/vincart/golib-sample-adapter/repository/mysql/model"
	"gitlab.id.vin/vincart/golib-sample-core/entity"
	"gitlab.id.vin/vincart/golib-sample-core/port"
	"gitlab.id.vin/vincart/golib-sample-internal/testing/base"
	"gitlab.id.vin/vincart/golib-sample-internal/testing/mock"
	"gitlab.id.vin/vincart/golib-test"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"net/http"
	"testing"
)

type CreateOrderControllerTest struct {
	*base.TestSuite
	authorization  string
	db             *gorm.DB
	eventPublisher *mock.EventPublisherMock
}

func TestCreateOrderControllerTest(t *testing.T) {
	// internal_service:secret
	s := CreateOrderControllerTest{authorization: "aW50ZXJuYWxfc2VydmljZTpzZWNyZXQ="}
	s.TestSuite = base.NewTestSuite(
		golibtest.WithTestingDir(".."),
		golibtest.ReplaceFxOption(
			fx.Provide(publisher.NewEventPublisherAdapter), fx.Provide(mock.NewEventPublisherMock),
		),
		golibtest.WithFxPopulate(&s.db),
		golibtest.WithFxOption(fx.Invoke(func(eventPublisher port.EventPublisher) {
			s.eventPublisher = eventPublisher.(*mock.EventPublisherMock)
		})),
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

	assert.Len(s.T(), s.eventPublisher.ReceivedEvents(), 1)
	actualEvent := s.eventPublisher.ReceivedEvents()[0]
	assert.Equal(s.T(), "OrderCreatedEvent", actualEvent.Name())
	assert.IsType(s.T(), &entity.Order{}, actualEvent.Payload())
	assert.EqualValues(s.T(), actualOrder.Id, actualEvent.Payload().(*entity.Order).Id)
	assert.EqualValues(s.T(), actualOrder.TotalAmount, actualEvent.Payload().(*entity.Order).TotalAmount)
	assert.InDelta(s.T(), actualOrder.CreatedAt.Unix(), actualEvent.Payload().(*entity.Order).CreatedAt.Unix(), 1)
}
