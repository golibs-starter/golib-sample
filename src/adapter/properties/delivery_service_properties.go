package properties

import "github.com/golibs-starter/golib/config"

type DeliveryServiceProperties struct {
	BaseUrl         string
	CreateOrderPath string
}

func NewDeliveryServiceProperties(loader config.Loader) (*DeliveryServiceProperties, error) {
	props := DeliveryServiceProperties{}
	err := loader.Bind(&props)
	return &props, err
}

func (o DeliveryServiceProperties) Prefix() string {
	return "app.services.delivery"
}
