package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.id.vin/vincart/golib-sample-adapter/service"
	"gitlab.id.vin/vincart/golib-sample-internal/resource"
	"gitlab.id.vin/vincart/golib/web/response"
)

type StatusController struct {
	statusService *service.StatusService
}

func NewStatusController(statusService *service.StatusService) *StatusController {
	return &StatusController{statusService: statusService}
}

// ReturnStatus
// @Summary API return status code
// @Tags StatusController
// @Accept  json
// @Produce  json
// @Security BasicAuth
// @Param    	code			path	string		true 	"status code"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /v1/statuses/{code} [get]
func (s StatusController) ReturnStatus(c *gin.Context) {
	code := c.Param("code")
	entity, err := s.statusService.Get(c, code)
	if err != nil {
		response.WriteError(c.Writer, err)
		return
	}
	resp := response.New(entity.HttpCode, "Ok", resource.NewStatus(entity))
	response.Write(c.Writer, resp)
}
