package mysql

import (
	"context"
	"gitlab.id.vin/vincart/golib-sample-adapter/repository/mysql/mapper"
	"gitlab.id.vin/vincart/golib-sample-adapter/repository/mysql/model"
	"gitlab.id.vin/vincart/golib-sample-core/entity"
	"gitlab.id.vin/vincart/golib-sample-core/entity/request"
	"gitlab.id.vin/vincart/golib-sample-core/port"
	"gorm.io/gorm"
)

type OrderRepositoryAdapter struct {
	base
}

func NewOrderRepositoryAdapter(db *gorm.DB) port.OrderRepository {
	return &OrderRepositoryAdapter{base: base{db: db}}
}

func (o OrderRepositoryAdapter) FindById(ctx context.Context, id int64) (*entity.Order, error) {
	var orderModel model.Order
	result := o.db.First(&orderModel, id)
	if result.Error != nil {
		return nil, o.handleError(result.Error)
	}
	return mapper.ModelToOrder(&orderModel), nil
}

func (o OrderRepositoryAdapter) CreateOrder(ctx context.Context, req *request.CreateOrderRequest) (*entity.Order, error) {
	orderModel := model.Order{
		TotalAmount: req.TotalAmount,
	}
	o.db.Create(&orderModel)
	return mapper.ModelToOrder(&orderModel), nil
}
