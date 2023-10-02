package service

import (
	"context"
	"github.com/golibs-starter/golib-sample-core/entity"
	"github.com/golibs-starter/golib-sample-core/usecase"
)

type StatusService struct {
	getStatusUseCase *usecase.GetStatusUseCase
}

func NewStatusService(getStatusUseCase *usecase.GetStatusUseCase) *StatusService {
	return &StatusService{getStatusUseCase: getStatusUseCase}
}

func (g StatusService) Get(ctx context.Context, code string) (*entity.Status, error) {
	return g.getStatusUseCase.Get(ctx, code)
}
