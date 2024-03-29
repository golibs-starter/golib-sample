package mysql

import (
	"context"
	"github.com/golibs-starter/golib-sample-adapter/repository/mysql/mapper"
	"github.com/golibs-starter/golib-sample-adapter/repository/mysql/model"
	"github.com/golibs-starter/golib-sample-core/entity"
	"github.com/golibs-starter/golib-sample-core/entity/request"
	"github.com/golibs-starter/golib-sample-core/port"
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
	return mapper.ModelToOrderEntity(&orderModel), nil
}

func (o OrderRepositoryAdapter) CreateOrder(ctx context.Context, req *request.CreateOrderRequest) (*entity.Order, error) {
	orderModel := model.Order{
		UserId:      req.UserId,
		TotalAmount: req.TotalAmount,
	}
	o.db.Create(&orderModel)
	return mapper.ModelToOrderEntity(&orderModel), nil
}
