package mysql

import (
	"github.com/golibs-starter/golib-sample-core/exception"
	"gorm.io/gorm"
)

type base struct {
	db *gorm.DB
}

func (b base) handleError(err error) error {
	if err == gorm.ErrRecordNotFound {
		return exception.ResourceNotFound
	}
	return err
}
