package api

import (
	"github.com/stretchr/testify/suite"
	"gitlab.com/golibs-starter/golib-sample-core/exception"
	"gitlab.com/golibs-starter/golib-sample-internal/testing/base"
	"gitlab.com/golibs-starter/golib-test"
	"net/http"
	"testing"
)

type GetStatusControllerTest struct {
	*base.TestSuite
	authorization string
}

func TestGetStatusControllerTest(t *testing.T) {
	// internal_service:secret
	s := GetStatusControllerTest{authorization: "aW50ZXJuYWxfc2VydmljZTpzZWNyZXQ="}
	s.TestSuite = base.NewTestSuite(
		golibtest.WithTestingDir(".."),
	)
	suite.Run(t, &s)
}

func (s GetStatusControllerTest) TestWhenWrongStatus_ShouldReturnBadRequest() {
	resp := golibtest.NewDefaultHttpClient(s.T()).Do(
		golibtest.NewRequestBuilder(s.T()).
			WithBasicAuthorization(s.authorization).
			Get(s.URL("/v1/statuses/wrong_status")),
	)
	defer resp.Body.Close()
	golibtest.NewRestAssured(s.T(), resp).
		Status(http.StatusBadRequest).
		Body("meta.code", exception.StatusInvalid.Code()).
		Body("meta.message", exception.StatusInvalid.Error())
}

func (s GetStatusControllerTest) TestWhenCorrectStatus_ShouldReturnSuccess() {
	resp := golibtest.NewDefaultHttpClient(s.T()).Do(
		golibtest.NewRequestBuilder(s.T()).
			WithBasicAuthorization(s.authorization).
			Get(s.URL("/v1/statuses/200")),
	)
	defer resp.Body.Close()
	golibtest.NewRestAssured(s.T(), resp).
		Status(http.StatusOK).
		Body("meta.code", 200).
		Body("meta.message", "Ok").
		Body("data.http_code", 200)
}
