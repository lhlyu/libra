package middleware

import (
	"github.com/didip/tollbooth"
	"github.com/iris-contrib/middleware/tollboothic"
	"github.com/kataras/iris/v12/context"
	"github.com/lhlyu/libra/common"
)

// 限制每秒请求数量
func Limiter() context.Handler {
	limiter := tollbooth.NewLimiter(common.Cfg.GetFloat64("server.limit"), nil)
	return tollboothic.LimitHandler(limiter)
}
