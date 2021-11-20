package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.id.vin/vincart/golib-sample-adapter/service"
	"gitlab.id.vin/vincart/golib-sample-core/exception"
	"gitlab.id.vin/vincart/golib-sample-internal/resource"
	baseEx "gitlab.id.vin/vincart/golib/exception"
	"gitlab.id.vin/vincart/golib/web/log"
	"gitlab.id.vin/vincart/golib/web/response"
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
// @Security BasicAuth
// @Param    	id			path	int		true 	"order id"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /v1/orders/{id} [get]
func (s OrderController) GetOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.WriteError(c.Writer, exception.OrderIdInvalid)
		return
	}
	entity, err := s.orderService.GetById(c, id)
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
// @Security BasicAuth
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /v1/orders [post]
func (s OrderController) CreateOrder(c *gin.Context) {
	var body resource.CreateOrderRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Warn(c, "Cannot bind request body, err [%s]", err)
		response.WriteError(c.Writer, baseEx.BadRequest)
		return
	}
	entity, err := s.orderService.Create(c, body.ToEntity())
	if err != nil {
		response.WriteError(c.Writer, err)
		return
	}
	response.Write(c.Writer, response.Ok(resource.NewOrder(entity)))
}
