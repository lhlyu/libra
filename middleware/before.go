package middleware

import (
	"github.com/kataras/iris/v12"
)

// 统一入口
func Before() iris.Handler {
	return func(ctx iris.Context) {
		ctx.Record() // 开启响应body记录
		ctx.Next()
	}
}
