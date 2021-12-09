package client

import (
	"context"
	"fmt"
	"gitlab.id.vin/vincart/golib-sample-adapter/http/client/dto"
	"gitlab.id.vin/vincart/golib-sample-adapter/properties"
	"gitlab.id.vin/vincart/golib-sample-core/entity"
	"gitlab.id.vin/vincart/golib-sample-core/port"
	baseEx "gitlab.id.vin/vincart/golib/exception"
	"gitlab.id.vin/vincart/golib/web/client"
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
