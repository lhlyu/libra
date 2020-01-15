package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/lhlyu/libra/logger"
)

func Log() context.Handler {
	return func(ctx iris.Context) {
		context := logger.WithLogger(ctx.Request().Context())
		req := ctx.Request().WithContext(context)
		ctx.ResetRequest(req)
		ctx.Next()
	}
}
