package service

import (
	"context"
	"fmt"
	"github.com/lhlyu/libra/result"
)

type IndexService struct {
	BaseService
}

func NewIndexService(ctx context.Context) *IndexService {
	return &IndexService{BaseService{
		Ctx: ctx,
	}}
}

func (s *IndexService) Hello(name string, age int) *result.R {
	s.Info(name, age)
	if age < 18 {
		return result.Failure.WithData(fmt.Sprintf("%s is not an adult", name))
	}
	return result.Success.WithData(fmt.Sprintf("%s is an adult", name))
}
