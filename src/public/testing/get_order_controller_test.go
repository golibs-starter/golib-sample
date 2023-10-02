package testing

import (
	"fmt"
	golibdataTestUtil "github.com/golibs-starter/golib-data/testutil"
	"github.com/golibs-starter/golib-sample-adapter/repository/mysql/model"
	"github.com/golibs-starter/golib-sample-core/exception"
	golibtest "github.com/golibs-starter/golib-test"
	"github.com/stretchr/testify/suite"
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
	s.Option(golibdataTestUtil.TruncateTablesOpt("orders"))
	suite.Run(t, s)
}

func (s *GetOrderControllerTest) TestGetOrderById_WhenOrderIdIsValid_ShouldReturnSuccess() {
	expectedOrder := model.Order{
		UserId:      "10",
		TotalAmount: 100000,
		CreatedAt:   time.Now(),
	}
	s.NoError(s.db.Create(&expectedOrder).Error)
	s.Greater(expectedOrder.Id, 0)

	golibtest.NewRestAssuredSuite(s).
		When().
		Get(fmt.Sprintf("/v1/orders/%d", expectedOrder.Id)).
		BearerToken(s.jwtTestUtil.CreateJwtToken("10")).
		Then().
		Status(http.StatusOK).
		Body("data.id", expectedOrder.Id).
		Body("data.user_id", expectedOrder.UserId).
		Body("data.total_amount", expectedOrder.TotalAmount).
		BodyCb("data.created_at", func(value interface{}) {
			s.InDelta(expectedOrder.CreatedAt.Unix(), value, 1)
		})
}

func (s *GetOrderControllerTest) TestGetOrderById_WhenGetOrderOfOtherUser_ShouldReturnOrderNotFound() {
	expectedOrder := model.Order{
		UserId:      "11",
		TotalAmount: 100000,
		CreatedAt:   time.Now(),
	}
	s.NoError(s.db.Create(&expectedOrder).Error)
	s.Greater(expectedOrder.Id, 0)

	golibtest.NewRestAssuredSuite(s).
		When().
		Get(fmt.Sprintf("/v1/orders/%d", expectedOrder.Id)).
		BearerToken(s.jwtTestUtil.CreateJwtToken("10")).
		Then().
		Status(http.StatusNotFound).
		Body("meta.code", exception.OrderNotFound.Code()).
		Body("meta.message", exception.OrderNotFound.Error())
}

func (s *GetOrderControllerTest) TestGetOrderById_WhenOrderIdIsInvalid_ShouldError() {
	golibtest.NewRestAssuredSuite(s).
		When().
		Get("/v1/orders/xxx").
		BearerToken(s.jwtTestUtil.CreateJwtToken("10")).
		Then().
		Status(http.StatusBadRequest).
		Body("meta.code", exception.OrderIdInvalid.Code()).
		Body("meta.message", exception.OrderIdInvalid.Error())
}
