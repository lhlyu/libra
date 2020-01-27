package controller

import (
	"errors"
	"github.com/kataras/iris/v12"
	"github.com/lhlyu/libra/service"
)

type IndexController struct {
	BaseController
}

type HelloParam struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (c *IndexController) Hello(ctx iris.Context) {
	param := &HelloParam{}
	if !c.getParams(ctx, param, false) {
		return
	}
	c.Error(ctx, errors.New("hello world"))
	svc := service.NewIndexService(ctx.Request().Context())
	ctx.JSON(svc.Hello(param.Name, param.Age))
}
