package service

import (
	"context"
	"github.com/lhlyu/libra/logger"
)

type BaseService struct {
	Ctx context.Context
}

func (s BaseService) Error(err error) {
	if err == nil {
		return
	}
	logger.GetLogger(s.Ctx, 1).WithField("error", err.Error()).Error()
}

func (s BaseService) Info(args ...interface{}) {
	if len(args) == 0 {
		return
	}
	logger.GetLogger(s.Ctx, 1).Info(args...)
}
