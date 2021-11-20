package client

import (
	"context"
	"fmt"
	"gitlab.id.vin/vincart/golib-sample-adapter/http/client/dto"
	"gitlab.id.vin/vincart/golib-sample-adapter/properties"
	"gitlab.id.vin/vincart/golib-sample-core/entity"
	"gitlab.id.vin/vincart/golib-sample-core/exception"
	"gitlab.id.vin/vincart/golib-sample-core/port"
	baseEx "gitlab.id.vin/vincart/golib/exception"
	"gitlab.id.vin/vincart/golib/web/client"
	"net/http"
)

type OrderRepositoryAdapter struct {
	httpClient client.ContextualHttpClient
	properties *properties.OrderRepositoryProperties
}

func NewOrderRepositoryAdapter(
	httpClient client.ContextualHttpClient,
	properties *properties.OrderRepositoryProperties,
) port.OrderRepositoryPort {
	return &OrderRepositoryAdapter{httpClient: httpClient, properties: properties}
}

func (o OrderRepositoryAdapter) FindById(ctx context.Context, id int64) (*entity.Order, error) {
	var orderResp dto.OrderResponseDto
	url := fmt.Sprintf(o.properties.BaseUrl+o.properties.GetOrderByIdPath, id)
	resp, err := o.httpClient.Get(ctx, url, &orderResp)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusOK {
		if orderResp.Data == nil {
			return nil, baseEx.NewWithCause(baseEx.SystemError, "Order data empty")
		}
		return orderResp.Data.ToEntity(), nil
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil, exception.OrderNotFound
	}
	return nil, baseEx.NewWithCause(baseEx.SystemError,
		fmt.Sprintf("Unexpected order response, http code [%d]", resp.StatusCode))
}
