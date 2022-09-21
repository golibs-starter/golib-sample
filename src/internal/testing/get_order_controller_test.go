package testing

import (
	"fmt"
	assert "github.com/stretchr/testify/require"
	"gitlab.com/golibs-starter/golib-sample-adapter/repository/mysql/model"
	"gitlab.com/golibs-starter/golib-sample-core/exception"
	"gitlab.com/golibs-starter/golib-test"
	"net/http"
	"testing"
	"time"
)

func TestGetOrderById_GiveValidId_WhenRepoResponseSuccess_ShouldReturnSuccess(t *testing.T) {
	expectedOrder := model.Order{
		TotalAmount: 100000,
		CreatedAt:   time.Now(),
	}
	assert.NoError(t, db.Create(&expectedOrder).Error)
	assert.Greater(t, expectedOrder.Id, 0)

	golibtest.NewRestAssured(t).
		When().
		Get(fmt.Sprintf("/v1/orders/%d", expectedOrder.Id)).
		BasicAuth("internal_service", "secret").
		Then().
		Status(http.StatusOK).
		Body("data.id", expectedOrder.Id).
		Body("data.total_amount", expectedOrder.TotalAmount).
		BodyCb("data.created_at", func(value interface{}) {
			assert.InDelta(t, expectedOrder.CreatedAt.Unix(), value, 1)
		})
}

func TestGetOrderById_GiveInvalidId_ShouldReturnSuccess(t *testing.T) {
	golibtest.NewRestAssured(t).
		When().
		Get("/v1/orders/xxx").
		BasicAuth("internal_service", "secret").
		Then().
		Status(http.StatusBadRequest).
		Body("meta.code", exception.OrderIdInvalid.Code()).
		Body("meta.message", exception.OrderIdInvalid.Error())
}
