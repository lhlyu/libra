package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/lhlyu/libra/result"
)

type IndexController struct {
	BaseController
}

type HelloParam struct {
	Name string `json:"name"`
}

func (c *IndexController) Hello(ctx iris.Context) {
	param := &HelloParam{}
	if !c.getParams(ctx, param, false) {
		return
	}
	ctx.JSON(result.Success.WithData(param))
}
