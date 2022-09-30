package testing

import (
	"github.com/stretchr/testify/suite"
	"gitlab.com/golibs-starter/golib-test"
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
	golibtest.NewRestAssured(s.T()).
		When().
		Get("/actuator/info").
		Then().
		Status(http.StatusOK).
		Body("meta.code", 200).
		Body("data.service_name", "Sample Worker")
}

func (s *ActuatorTest) TestActuatorHealth_ShouldReturnSuccess() {
	golibtest.NewRestAssured(s.T()).
		When().
		Get("/actuator/health").
		Then().
		Status(http.StatusOK).
		Body("meta.code", 200).
		Body("data.status", "UP")
}