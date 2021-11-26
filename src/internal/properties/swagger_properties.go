package properties

import (
	"gitlab.id.vin/vincart/golib/config"
)

func NewSwaggerProperties(loader config.Loader) (*SwaggerProperties, error) {
	props := SwaggerProperties{}
	err := loader.Bind(&props)
	return &props, err
}

type SwaggerProperties struct {
	Enabled bool `default:"false"`
}

func (t SwaggerProperties) Prefix() string {
	return "vinid.swagger"
}