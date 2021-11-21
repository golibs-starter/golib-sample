package api

import (
	assert "github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gitlab.id.vin/vincart/golib-sample-adapter/properties"
	"gitlab.id.vin/vincart/golib-sample-adapter/publisher"
	"gitlab.id.vin/vincart/golib-sample-core/entity"
	"gitlab.id.vin/vincart/golib-sample-core/port"
	"gitlab.id.vin/vincart/golib-sample-internal/testing/base"
	"gitlab.id.vin/vincart/golib-sample-internal/testing/mock"
	golibtest "gitlab.id.vin/vincart/golib-test"
	baseEx "gitlab.id.vin/vincart/golib/exception"
	"go.uber.org/fx"
	"net/http"
	"testing"
)

type CreateOrderControllerTest struct {
	*base.TestSuite
	authorization  string
	props          *properties.OrderRepositoryProperties
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
		golibtest.WithFxPopulate(&s.props),
		golibtest.WithFxOption(fx.Invoke(func(eventPublisher port.EventPublisher) {
			s.eventPublisher = eventPublisher.(*mock.EventPublisherMock)
		})),
	)
	suite.Run(t, &s)
}

func (s CreateOrderControllerTest) TestCreateOrder_GiveValidBody_WhenRepoResponseSuccess_ShouldReturnSuccess() {
	ts := golibtest.NewHttpTestServer(http.StatusCreated, `{
  "meta": {
    "code": 201,
    "message": "Successful"
  },
  "data": {
    "id": 20,
    "total_amount": 85000,
    "created_at": 1637415974
  }
}`)
	s.props.BaseUrl = ts.URL

	resp := golibtest.NewDefaultHttpClient(s.T()).Do(
		golibtest.NewRequestBuilder(s.T()).
			WithBasicAuthorization(s.authorization).
			WithBodyString(`{"total_amount": 85000}`).
			Post(s.URL("/v1/orders")),
	)
	defer resp.Body.Close()

	golibtest.NewRestAssured(s.T(), resp).
		Status(http.StatusOK).
		Body("data.id", 20).
		Body("data.total_amount", 85000).
		Body("data.created_at", 1637415974)

	assert.Len(s.T(), s.eventPublisher.ReceivedEvents(), 1)
	actualEvent := s.eventPublisher.ReceivedEvents()[0]
	assert.Equal(s.T(), "OrderCreatedEvent", actualEvent.Name())
	assert.IsType(s.T(), &entity.Order{}, actualEvent.Payload())
	assert.EqualValues(s.T(), 20, actualEvent.Payload().(*entity.Order).Id)
	assert.EqualValues(s.T(), 85000, actualEvent.Payload().(*entity.Order).TotalAmount)
	assert.EqualValues(s.T(), 1637415974, actualEvent.Payload().(*entity.Order).CreatedAt.Unix())
}

func (s CreateOrderControllerTest) TestCreateOrder_GiveValidBody_WhenRepoResponseError_ShouldReturnSuccess() {
	ts := golibtest.NewHttpTestServer(http.StatusBadRequest, `{
  "meta": {
    "code": 400,
    "message": "Bad request"
  }
}`)
	s.props.BaseUrl = ts.URL

	resp := golibtest.NewDefaultHttpClient(s.T()).Do(
		golibtest.NewRequestBuilder(s.T()).
			WithBasicAuthorization(s.authorization).
			WithBodyString(`{"total_amount": 85000}`).
			Post(s.URL("/v1/orders")),
	)
	defer resp.Body.Close()

	golibtest.NewRestAssured(s.T(), resp).
		Status(http.StatusInternalServerError).
		Body("meta.code", baseEx.SystemError.Code()).
		Body("meta.message", baseEx.SystemError.Error())
}
