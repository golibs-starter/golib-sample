package testing

import (
	"gitlab.com/golibs-starter/golib-sample-core/exception"
	"gitlab.com/golibs-starter/golib-test"
	"net/http"
	"testing"
)

func TestWhenWrongStatus_ShouldReturnBadRequest(t *testing.T) {
	golibtest.NewRestAssured(t).
		When().
		Get("/v1/statuses/wrong_status").
		BasicAuth("internal_service", "secret").
		Then().
		Status(http.StatusBadRequest).
		Body("meta.code", exception.StatusInvalid.Code()).
		Body("meta.message", exception.StatusInvalid.Error())
}

func TestWhenCorrectStatus_ShouldReturnSuccess(t *testing.T) {
	golibtest.NewRestAssured(t).
		When().
		Get("/v1/statuses/200").
		BasicAuth("internal_service", "secret").
		Then().
		Status(http.StatusOK).
		Body("meta.code", 200).
		Body("meta.message", "Ok").
		Body("data.http_code", 200)
}
