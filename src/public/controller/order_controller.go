package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/golibs-starter/golib-sample-adapter/service"
	"gitlab.com/golibs-starter/golib-sample-core/exception"
	"gitlab.com/golibs-starter/golib-sample-public/resource"
	"gitlab.com/golibs-starter/golib-security/web/context"
	baseEx "gitlab.com/golibs-starter/golib/exception"
	"gitlab.com/golibs-starter/golib/log"
	"gitlab.com/golibs-starter/golib/web/response"
	"strconv"
)

type OrderController struct {
	orderService *service.OrderService
}

func NewOrderController(orderService *service.OrderService) *OrderController {
	return &OrderController{orderService: orderService}
}

// GetOrder
// @Summary Get order by ID
// @Tags OrderController
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param    	id			path	int		true 	"order id"
// @Success 200 {object} response.Response{data=resource.Order}
// @Failure 500 {object} response.Response
// @Router /v1/orders/{id} [get]
func (s OrderController) GetOrder(c *gin.Context) {
	idStr := c.Param("id")
	orderId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.WriteError(c.Writer, exception.OrderIdInvalid)
		return
	}
	entity, err := s.orderService.GetByIdAndUser(c, context.GetUserDetails(c.Request).Username(), orderId)
	if err != nil {
		response.WriteError(c.Writer, err)
		return
	}
	response.Write(c.Writer, response.Ok(resource.NewOrder(entity)))
}

// CreateOrder
// @Summary Create order
// @Tags OrderController
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param request body resource.CreateOrderRequest true "Request body"
// @Success 200 {object} response.Response{data=resource.Order}
// @Failure 500 {object} response.Response
// @Router /v1/orders [post]
func (s OrderController) CreateOrder(c *gin.Context) {
	contextualLogger := log.WithCtx(c)
	var body resource.CreateOrderRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		contextualLogger.WithError(err).Info("Cannot bind request body")
		response.WriteError(c.Writer, baseEx.BadRequest)
		return
	}

	packet := body.ToEntity()
	packet.UserId = context.GetUserDetails(c.Request).Username()
	entity, err := s.orderService.Create(c, packet)
	if err != nil {
		contextualLogger.WithError(err).Info("Create order failed")
		response.WriteError(c.Writer, err)
		return
	}
	contextualLogger.WithAny("order_id", entity.Id).Info("Order created")
	response.Write(c.Writer, response.Ok(resource.NewOrder(entity)))
}
