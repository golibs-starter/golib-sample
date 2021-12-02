package api

import (
	"github.com/stretchr/testify/suite"
	"gitlab.id.vin/vincart/golib-sample-public/testing/base"
	"gitlab.id.vin/vincart/golib-test"
	"gitlab.id.vin/vincart/golib/config"
	"go.uber.org/fx"
	"net/http"
	"testing"
)

type ActuatorTest struct {
	*base.TestSuite
	appProps *config.AppProperties
}

func TestActuatorTest(t *testing.T) {
	s := ActuatorTest{}
	s.TestSuite = base.NewTestSuite(
		golibtest.WithTestingDir(".."),
		golibtest.WithFxOption(fx.Invoke(func(appProps *config.AppProperties) {
			s.appProps = appProps
		})),
	)
	suite.Run(t, &s)
}

func (s ActuatorTest) TestActuatorInfo_ShouldReturnSuccess() {
	resp := golibtest.NewDefaultHttpClient(s.T()).Do(
		golibtest.NewRequestBuilder(s.T()).Get(s.URL("/actuator/info")),
	)
	defer resp.Body.Close()
	golibtest.NewRestAssured(s.T(), resp).
		Status(http.StatusOK).
		Body("meta.code", 200).
		Body("data.service_name", s.appProps.Name)
}

func (s ActuatorTest) TestActuatorHealth_ShouldReturnSuccess() {
	resp := golibtest.NewDefaultHttpClient(s.T()).Do(
		golibtest.NewRequestBuilder(s.T()).Get(s.URL("/actuator/health")),
	)
	defer resp.Body.Close()
	golibtest.NewRestAssured(s.T(), resp).
		Status(http.StatusOK).
		Body("meta.code", 200).
		Body("data.status", "UP").
		Body("data.components.redis.status", "UP").
		Body("data.components.datasource.status", "UP")
}
