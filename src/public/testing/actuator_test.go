package testing

import (
	golibtest "github.com/golibs-starter/golib-test"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type ActuatorTest struct {
	TestSuite
}

func TestActuatorTest(t *testing.T) {
	suite.Run(t, new(ActuatorTest))
}

func (s *ActuatorTest) TestActuatorInfo_ShouldReturnSuccess() {
	golibtest.NewRestAssuredSuite(s).
		When().
		Get("/actuator/info").
		Then().
		Status(http.StatusOK).
		Body("meta.code", 200).
		Body("data.service_name", "Sample Public API")
}

func (s *ActuatorTest) TestActuatorHealth_ShouldReturnSuccess() {
	golibtest.NewRestAssuredSuite(s).
		When().
		Get("/actuator/health").
		Then().
		Status(http.StatusOK).
		Body("meta.code", 200).
		Body("data.status", "UP").
		Body("data.components.redis.status", "UP").
		Body("data.components.datasource.status", "UP")
}
