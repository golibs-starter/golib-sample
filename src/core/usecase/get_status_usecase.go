package usecase

import (
	"context"
	"gitlab.com/golibs-starter/golib-sample-core/entity"
	"gitlab.com/golibs-starter/golib-sample-core/exception"
	"gitlab.com/golibs-starter/golib/log"
	"strconv"
)

type GetStatusUseCase struct {
}

func NewGetStatusUseCase() *GetStatusUseCase {
	return &GetStatusUseCase{}
}

func (g GetStatusUseCase) Get(ctx context.Context, code string) (*entity.Status, error) {
	httpCode, err := strconv.Atoi(code)
	if err != nil {
		log.WithCtx(ctx).WithError(err).Warnf("Status code is invalid [%s]", code)
		return nil, exception.StatusInvalid
	}
	return &entity.Status{HttpCode: httpCode}, nil
}
