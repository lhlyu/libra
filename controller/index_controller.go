package controller

import (
    "github.com/kataras/iris/v12"
)

type IndexController struct {
    controller
}

func (c *IndexController) Hello(ctx iris.Context){
    ctx.Text("test")
}
