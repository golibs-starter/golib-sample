package api

import (
	"fmt"
	assert "github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gitlab.id.vin/vincart/golib-sample-adapter/repository/mysql/model"
	"gitlab.id.vin/vincart/golib-sample-core/exception"
	"gitlab.id.vin/vincart/golib-sample-internal/testing/base"
	"gitlab.id.vin/vincart/golib-test"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"net/http"
	"testing"
	"time"
)

type GetOrderControllerTest struct {
	*base.TestSuite
	authorization string
	db            *gorm.DB
}

func TestGetOrderControllerTest(t *testing.T) {
	// internal_service:secret
	s := GetOrderControllerTest{authorization: "aW50ZXJuYWxfc2VydmljZTpzZWNyZXQ="}
	s.TestSuite = base.NewTestSuite(
		golibtest.WithTestingDir(".."),
		golibtest.WithFxPopulate(&s.db),
		golibtest.WithFxOption(fx.Invoke(func(db *gorm.DB) {
			db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Order{})
		})),
	)
	suite.Run(t, &s)
}

func (s GetOrderControllerTest) TestGetOrderById_GiveValidId_WhenRepoResponseSuccess_ShouldReturnSuccess() {
	expectedOrder := model.Order{
		TotalAmount: 100000,
		CreatedAt:   time.Now(),
	}
	assert.NoError(s.T(), s.db.Create(&expectedOrder).Error)
	assert.Greater(s.T(), expectedOrder.Id, 0)

	resp := golibtest.NewDefaultHttpClient(s.T()).Do(
		golibtest.NewRequestBuilder(s.T()).
			WithBasicAuthorization(s.authorization).
			Get(s.URL(fmt.Sprintf("/v1/orders/%d", expectedOrder.Id))),
	)
	defer resp.Body.Close()

	golibtest.NewRestAssured(s.T(), resp).
		Status(http.StatusOK).
		Body("data.id", expectedOrder.Id).
		Body("data.total_amount", expectedOrder.TotalAmount).
		BodyCb("data.created_at", func(value interface{}) {
			assert.InDelta(s.T(), expectedOrder.CreatedAt.Unix(), value, 1)
		})
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
