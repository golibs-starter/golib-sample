package resource

import "github.com/golibs-starter/golib-sample-core/entity"

type Status struct {
	HttpCode int `json:"http_code"`
}

func NewStatus(entity *entity.Status) *Status {
	return &Status{
		HttpCode: entity.HttpCode,
	}
}
