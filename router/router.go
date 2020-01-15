package router

import (
	"github.com/kataras/iris/v12"
	"github.com/lhlyu/libra/logger"
)

func SetRouter(app *iris.Application) {
	//app.AllowMethods(iris.MethodOptions)
	//
	//ctr := &controller.Controller{}
	//
	//
	//app.Party("/api")

	app.Get("/", func(ctx iris.Context) {
		name := ctx.URLParam("name")
		logger.Log(ctx).Infoln("hello", name)
		logger.Log(ctx).Infoln("hi", name)
		ctx.Text("hello %s", name)
	})

}
