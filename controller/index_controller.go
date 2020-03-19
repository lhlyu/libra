package controller

import (
	"github.com/lhlyu/libra/service"
	"github.com/lhlyu/yutil/v2"
)

type IndexController struct {
}

type HelloParam struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// http://localhost:8080/index?name=tom&age=12
func (c *IndexController) Hello(ctx *Context) {
	param := &HelloParam{}
	if !ctx.GetParams(param, false) {
		return
	}
	// 打印日志
	ctx.Info("param:", yutil.Json.Marshal(param))
	svc := service.NewIndexService(ctx)
	ctx.JSON(svc.Hello(param.Name, param.Age))
}
