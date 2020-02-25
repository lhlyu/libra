package controller

import (
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

// http://localhost:8080/index?name=tom&age=12
func (c *IndexController) Hello(ctx iris.Context) {
	param := &HelloParam{}
	if !c.getParams(ctx, param, false) {
		return
	}
	svc := service.NewIndexService(ctx)
	ctx.JSON(svc.Hello(param.Name, param.Age))
}
