package middleware

import (
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/context"
)

func Log() context.Handler {
	return func(ctx iris.Context) {
		ctx.Next()
	}
}
