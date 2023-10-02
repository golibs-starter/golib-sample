package client

import (
	"context"
	"fmt"
	"github.com/golibs-starter/golib-sample-adapter/http/client/dto"
	"github.com/golibs-starter/golib-sample-adapter/properties"
	"github.com/golibs-starter/golib-sample-core/entity"
	"github.com/golibs-starter/golib-sample-core/port"
	baseEx "github.com/golibs-starter/golib/exception"
	"github.com/golibs-starter/golib/web/client"
	"net/http"
)

type DeliveryServiceAdapter struct {
	httpClient client.ContextualHttpClient
	properties *properties.DeliveryServiceProperties
}

func NewDeliveryServiceAdapter(
	httpClient client.ContextualHttpClient,
	properties *properties.DeliveryServiceProperties,
) port.DeliveryService {
	return &DeliveryServiceAdapter{httpClient: httpClient, properties: properties}
}

func (o DeliveryServiceAdapter) CreateOrder(ctx context.Context, order *entity.Order) (string, error) {
	var orderResp dto.OrderDeliveryResponseDto
	var requestBody = dto.NewCreateOrderDeliveryRequest(order)
	resp, err := o.httpClient.Post(ctx, o.properties.BaseUrl+o.properties.CreateOrderPath, requestBody, &orderResp)
	if err != nil {
		return "", err
	}
	if resp.StatusCode == http.StatusCreated {
		if orderResp.Data == nil {
			return "", baseEx.NewWithCause(baseEx.SystemError, "Delivery order data empty")
		}
		return orderResp.Data.Id, nil
	}
	return "", baseEx.NewWithCause(baseEx.SystemError,
		fmt.Sprintf("Unexpected order delivery response, http code [%d]", resp.StatusCode))
}
