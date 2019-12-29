package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/lhlyu/libra/common"
	"github.com/lhlyu/libra/util"
	"time"
)

func Log() context.Handler {
	return func(ctx iris.Context) {
		// 加入唯一ID
		traceId := util.GetGID()
		ctx.Values().Set(common.X_TRACE, traceId)
		now := time.Now()
		ctx.Values().Set(common.X_TIME, now)
		common.Ylog.Log(2, "debug", traceId, "middleware", ctx.String())
		ctx.Next()
	}
}
