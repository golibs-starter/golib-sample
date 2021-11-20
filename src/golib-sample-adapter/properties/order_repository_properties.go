package properties

import "gitlab.id.vin/vincart/golib/config"

type OrderRepositoryProperties struct {
	BaseUrl          string
	GetOrderByIdPath string
}

func NewOrderRepositoryProperties(loader config.Loader) (*OrderRepositoryProperties, error) {
	props := OrderRepositoryProperties{}
	err := loader.Bind(&props)
	return &props, err
}

func (o OrderRepositoryProperties) Prefix() string {
	return "app.services.order"
}
