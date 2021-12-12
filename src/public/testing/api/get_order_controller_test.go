package api

import (
	"fmt"
	assert "github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gitlab.com/golibs-starter/golib-sample-adapter/properties"
	"gitlab.com/golibs-starter/golib-sample-adapter/repository/mysql/model"
	"gitlab.com/golibs-starter/golib-sample-core/exception"
	"gitlab.com/golibs-starter/golib-sample-public/testing/base"
    "gitlab.com/golibs-starter/golib-test"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"net/http"
	"testing"
	"time"
)

type GetOrderControllerTest struct {
	*base.TestSuite
	db *gorm.DB
}

func TestGetOrderControllerTest(t *testing.T) {
	s := GetOrderControllerTest{}
	s.TestSuite = base.NewTestSuite(
		golibtest.WithTestingDir(".."),
		golibtest.WithFxPopulate(&s.db),
		golibtest.WithFxOption(fx.Invoke(func(db *gorm.DB) {
			db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Order{})
		})),
	)
	suite.Run(t, &s)
}

func (s GetOrderControllerTest) TestGetOrderById_WhenOrderIdIsValid_ShouldReturnSuccess() {
	expectedOrder := model.Order{
		UserId:      "10",
		TotalAmount: 100000,
		CreatedAt:   time.Now(),
	}
	assert.NoError(s.T(), s.db.Create(&expectedOrder).Error)
	assert.Greater(s.T(), expectedOrder.Id, 0)

	resp := golibtest.NewDefaultHttpClient(s.T()).Do(
		golibtest.NewRequestBuilder(s.T()).
			WithBearerAuthorization(s.CreateJwtToken("10")).
			Get(s.URL(fmt.Sprintf("/v1/orders/%d", expectedOrder.Id))),
	)
	defer resp.Body.Close()

	golibtest.NewRestAssured(s.T(), resp).
		Status(http.StatusOK).
		Body("data.id", expectedOrder.Id).
		Body("data.user_id", expectedOrder.UserId).
		Body("data.total_amount", expectedOrder.TotalAmount).
		BodyCb("data.created_at", func(value interface{}) {
			assert.InDelta(s.T(), expectedOrder.CreatedAt.Unix(), value, 1)
		})
}

func (s GetOrderControllerTest) TestGetOrderById_WhenGetOrderOfOtherUser_ShouldReturnOrderNotFound() {
	expectedOrder := model.Order{
		UserId:      "11",
		TotalAmount: 100000,
		CreatedAt:   time.Now(),
	}
	assert.NoError(s.T(), s.db.Create(&expectedOrder).Error)
	assert.Greater(s.T(), expectedOrder.Id, 0)

	resp := golibtest.NewDefaultHttpClient(s.T()).Do(
		golibtest.NewRequestBuilder(s.T()).
			WithBearerAuthorization(s.CreateJwtToken("10")).
			Get(s.URL(fmt.Sprintf("/v1/orders/%d", expectedOrder.Id))),
	)
	defer resp.Body.Close()

	golibtest.NewRestAssured(s.T(), resp).
		Status(http.StatusNotFound).
		Body("meta.code", exception.OrderNotFound.Code()).
		Body("meta.message", exception.OrderNotFound.Error())
}

func (s GetOrderControllerTest) TestGetOrderById_WhenOrderIdIsInvalid_ShouldError() {
	resp := golibtest.NewDefaultHttpClient(s.T()).Do(
		golibtest.NewRequestBuilder(s.T()).
			WithBearerAuthorization(s.CreateJwtToken("10")).
			Get(s.URL("/v1/orders/xxx")),
	)
	defer resp.Body.Close()

	golibtest.NewRestAssured(s.T(), resp).
		Status(http.StatusBadRequest).
		Body("meta.code", exception.OrderIdInvalid.Code()).
		Body("meta.message", exception.OrderIdInvalid.Error())
}
