package testing

import (
	"fmt"
	"github.com/golibs-starter/golib-sample-adapter/repository/mysql/model"
	"github.com/golibs-starter/golib-sample-core/exception"
	golibtest "github.com/golibs-starter/golib-test"
	assert "github.com/stretchr/testify/require"
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

func TestGetOrderById_GiveInvalidId_ShouldReturnBadRequest(t *testing.T) {
	golibtest.NewRestAssured(t).
		When().
		Get("/v1/orders/xxx").
		BasicAuth("internal_service", "secret").
		Then().
		Status(http.StatusBadRequest).
		Body("meta.code", exception.OrderIdInvalid.Code()).
		Body("meta.message", exception.OrderIdInvalid.Error())
}

func TestGetOrderById_GiveNotFoundId_ShouldReturnNotFound(t *testing.T) {
	golibtest.NewRestAssured(t).
		When().
		Get("/v1/orders/10000").
		BasicAuth("internal_service", "secret").
		Then().
		Status(http.StatusNotFound).
		Body("meta.code", exception.ResourceNotFound.Code()).
		Body("meta.message", exception.ResourceNotFound.Error())
}
