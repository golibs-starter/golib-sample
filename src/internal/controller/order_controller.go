package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/golibs-starter/golib-sample-adapter/service"
	"gitlab.com/golibs-starter/golib-sample-core/exception"
	"gitlab.com/golibs-starter/golib-sample-internal/resource"
	"gitlab.com/golibs-starter/golib/log"
	"gitlab.com/golibs-starter/golib/web/response"
	"strconv"
)

type OrderController struct {
	orderService *service.OrderService
	loggerIns    log.Logger
}

func NewOrderController(orderService *service.OrderService, loggerIns log.Logger) *OrderController {
	return &OrderController{orderService: orderService, loggerIns: loggerIns}
}

// GetOrder
// @Summary Get order by ID
// @Tags OrderController
// @Accept  json
// @Produce  json
// @Security BasicAuth
// @Param    	id			path	int		true 	"order id"
// @Success 200 {object} response.Response{data=resource.Order}
// @Failure 500 {object} response.Response
// @Router /v1/orders/{id} [get]
func (s OrderController) GetOrder(c *gin.Context) {
	contextualLogger := log.WithCtx(c)
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		contextualLogger.WithError(err).Errorf("Cannot parse id")
		response.WriteError(c.Writer, exception.OrderIdInvalid)
		return
	}

	logger := contextualLogger.WithAny("order_id", id)
	entity, err := s.orderService.GetById(c, id)
	if err != nil {
		logger.WithError(err).Errorf("Get order by id failed")
		response.WriteError(c.Writer, err)
		return
	}
	logger.Infof("Success to get order by id")
	response.Write(c.Writer, response.Ok(resource.NewOrder(entity)))
}
