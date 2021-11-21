package api

import (
	"github.com/stretchr/testify/suite"
	"gitlab.id.vin/vincart/golib-sample-adapter/properties"
	"gitlab.id.vin/vincart/golib-sample-core/exception"
	"gitlab.id.vin/vincart/golib-sample-internal/testing/base"
	golibtest "gitlab.id.vin/vincart/golib-test"
	baseEx "gitlab.id.vin/vincart/golib/exception"
	"net/http"
	"testing"
)

type GetOrderControllerTest struct {
	*base.TestSuite
	authorization string
	props         *properties.OrderRepositoryProperties
}

func TestGetOrderControllerTest(t *testing.T) {
	// internal_service:secret
	s := GetOrderControllerTest{authorization: "aW50ZXJuYWxfc2VydmljZTpzZWNyZXQ="}
	s.TestSuite = base.NewTestSuite(
		golibtest.WithTestingDir(".."),
		golibtest.WithFxPopulate(&s.props),
	)
	suite.Run(t, &s)
}

func (s GetOrderControllerTest) TestGetOrderById_GiveValidId_WhenRepoResponseSuccess_ShouldReturnSuccess() {
	ts := golibtest.NewHttpTestServer(http.StatusOK, `{
  "meta": {
    "code": 200,
    "message": "Successful"
  },
  "data": {
    "id": 19,
    "total_amount": 100000,
    "created_at": 1637415974
  }
}`)
	s.props.BaseUrl = ts.URL

	resp := golibtest.NewDefaultHttpClient(s.T()).Do(
		golibtest.NewRequestBuilder(s.T()).
			WithBasicAuthorization(s.authorization).
			Get(s.URL("/v1/orders/19")),
	)
	defer resp.Body.Close()

	golibtest.NewRestAssured(s.T(), resp).
		Status(http.StatusOK).
		Body("data.id", 19).
		Body("data.total_amount", 100000).
		Body("data.created_at", 1637415974)
}

func (s GetOrderControllerTest) TestGetOrderById_GiveValidId_WhenRepoResponseError_ShouldReturnSuccess() {
	ts := golibtest.NewHttpTestServer(http.StatusInternalServerError, `{
  "meta": {
    "code": 500,
    "message": "Unexpected error"
  }
}`)
	s.props.BaseUrl = ts.URL

	resp := golibtest.NewDefaultHttpClient(s.T()).Do(
		golibtest.NewRequestBuilder(s.T()).
			WithBasicAuthorization(s.authorization).
			Get(s.URL("/v1/orders/19")),
	)
	defer resp.Body.Close()

	golibtest.NewRestAssured(s.T(), resp).
		Status(http.StatusInternalServerError).
		Body("meta.code", baseEx.SystemError.Code()).
		Body("meta.message", baseEx.SystemError.Message())
}

func (s GetOrderControllerTest) TestGetOrderById_GiveValidId_WhenRepoResponseNotFound_ShouldReturnSuccess() {
	ts := golibtest.NewHttpTestServer(http.StatusNotFound, `{
  "meta": {
    "code": 404,
    "message": "Order not found"
  }
}`)
	s.props.BaseUrl = ts.URL

	resp := golibtest.NewDefaultHttpClient(s.T()).Do(
		golibtest.NewRequestBuilder(s.T()).
			WithBasicAuthorization(s.authorization).
			Get(s.URL("/v1/orders/19")),
	)
	defer resp.Body.Close()

	golibtest.NewRestAssured(s.T(), resp).
		Status(http.StatusNotFound).
		Body("meta.code", exception.OrderNotFound.Code()).
		Body("meta.message", exception.OrderNotFound.Error())
}

func (s GetOrderControllerTest) TestGetOrderById_GiveInvalidId_ShouldReturnSuccess() {
	resp := golibtest.NewDefaultHttpClient(s.T()).Do(
		golibtest.NewRequestBuilder(s.T()).
			WithBasicAuthorization(s.authorization).
			Get(s.URL("/v1/orders/xxx")),
	)
	defer resp.Body.Close()

	golibtest.NewRestAssured(s.T(), resp).
		Status(http.StatusBadRequest).
		Body("meta.code", exception.OrderIdInvalid.Code()).
		Body("meta.message", exception.OrderIdInvalid.Error())
}
