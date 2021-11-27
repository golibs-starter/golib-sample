package client

import (
	"context"
	"fmt"
	"gitlab.com/golibs-starter/golib-sample-adapter/http/client/dto"
	"gitlab.com/golibs-starter/golib-sample-adapter/properties"
	"gitlab.com/golibs-starter/golib-sample-core/entity"
	"gitlab.com/golibs-starter/golib-sample-core/entity/request"
	"gitlab.com/golibs-starter/golib-sample-core/exception"
	"gitlab.com/golibs-starter/golib-sample-core/port"
	baseEx "gitlab.com/golibs-starter/golib/exception"
	"gitlab.com/golibs-starter/golib/web/client"
	"net/http"
)

type OrderRepositoryAdapter struct {
	httpClient client.ContextualHttpClient
	properties *properties.OrderRepositoryProperties
}

func NewOrderRepositoryAdapter(
	httpClient client.ContextualHttpClient,
	properties *properties.OrderRepositoryProperties,
) port.OrderRepository {
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

func (o OrderRepositoryAdapter) CreateOrder(ctx context.Context, req *request.CreateOrderRequest) (*entity.Order, error) {
	var orderResp dto.OrderResponseDto
	var requestBody = dto.NewCreateOrderRequest(req)
	resp, err := o.httpClient.Post(ctx, o.properties.BaseUrl+o.properties.CreateOrderPath, requestBody, &orderResp)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusCreated {
		if orderResp.Data == nil {
			return nil, baseEx.NewWithCause(baseEx.SystemError, "Order data empty")
		}
		return orderResp.Data.ToEntity(), nil
	}
	return nil, baseEx.NewWithCause(baseEx.SystemError,
		fmt.Sprintf("Unexpected order response, http code [%d]", resp.StatusCode))
}
