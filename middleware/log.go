package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/lhlyu/libra/trace"
)

func Log() iris.Handler {
	return func(ctx iris.Context) {
		// 加入日志追踪者
		ctx.Values().Set(trace.TRACKER, trace.NewTracker())
		ctx.Next()
	}
}
