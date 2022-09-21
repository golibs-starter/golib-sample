package testing

import (
	"gitlab.com/golibs-starter/golib-test"
	"net/http"
	"testing"
)

func TestActuatorInfo_ShouldSuccess(t *testing.T) {
	golibtest.NewRestAssured(t).
		When().
		Get("/actuator/info").
		Then().
		Status(http.StatusOK).
		Body("meta.code", 200).
		Body("data.service_name", "Sample Internal API")
}

func TestActuatorHealth_ShouldSuccess(t *testing.T) {
	golibtest.NewRestAssured(t).
		When().
		Get("/actuator/health").
		Then().
		Status(http.StatusOK).
		Body("meta.code", 200).
		Body("data.status", "UP").
		Body("data.components.redis.status", "UP").
		Body("data.components.datasource.status", "UP")
}
