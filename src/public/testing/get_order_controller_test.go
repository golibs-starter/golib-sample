package testing

import (
	"fmt"
	assert "github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gitlab.com/golibs-starter/golib-sample-adapter/repository/mysql/model"
	"gitlab.com/golibs-starter/golib-sample-core/exception"
	"gitlab.com/golibs-starter/golib-test"
	"gorm.io/gorm"
	"net/http"
	"testing"
	"time"
)

type GetOrderControllerTest struct {
	TestSuite
	db *gorm.DB
}

func TestGetOrderControllerTest(t *testing.T) {
	s := new(GetOrderControllerTest)
	s.Populate(&s.db)
	s.Invoke(TruncateTablesInvoker("orders"))
	suite.Run(t, s)
}

func (s GetOrderControllerTest) TestGetOrderById_WhenOrderIdIsValid_ShouldReturnSuccess() {
	expectedOrder := model.Order{
		UserId:      "10",
		TotalAmount: 100000,
		CreatedAt:   time.Now(),
	}
	assert.NoError(s.T(), s.db.Create(&expectedOrder).Error)
	assert.Greater(s.T(), expectedOrder.Id, 0)

	golibtest.NewRestAssured(s.T()).
		When().
		Get(fmt.Sprintf("/v1/orders/%d", expectedOrder.Id)).
		BearerToken(s.CreateJwtToken("10")).
		Then().
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

	golibtest.NewRestAssured(s.T()).
		When().
		Get(fmt.Sprintf("/v1/orders/%d", expectedOrder.Id)).
		BearerToken(s.CreateJwtToken("10")).
		Then().
		Status(http.StatusNotFound).
		Body("meta.code", exception.OrderNotFound.Code()).
		Body("meta.message", exception.OrderNotFound.Error())
}

func (s GetOrderControllerTest) TestGetOrderById_WhenOrderIdIsInvalid_ShouldError() {
	golibtest.NewRestAssured(s.T()).
		When().
		Get("/v1/orders/xxx").
		BearerToken(s.CreateJwtToken("10")).
		Then().
		Status(http.StatusBadRequest).
		Body("meta.code", exception.OrderIdInvalid.Code()).
		Body("meta.message", exception.OrderIdInvalid.Error())
}
