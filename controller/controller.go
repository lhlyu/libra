package controller

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/lhlyu/libra/common"
	"github.com/lhlyu/libra/result"
	"github.com/lhlyu/libra/trace"
	"gopkg.in/go-playground/validator.v9"
	"strings"
	"sync"
	"time"
)

// 自定义上下文
type Context struct {
	iris.Context
	trace.BaseTracker
}

func H(h func(*Context)) iris.Handler {
	return func(original iris.Context) {
		ctx := acquire(original)
		h(ctx)
		release(ctx)
	}
}

var contextPool = sync.Pool{New: func() interface{} {
	return &Context{}
}}

func acquire(original iris.Context) *Context {
	ctx := contextPool.Get().(*Context)
	ctx.Context = original
	ctx.BaseTracker = trace.NewBaseTracker(ctx.Values().Get(trace.TRACKER).(*trace.Tracker))
	return ctx
}

func release(ctx *Context) {
	contextPool.Put(ctx)
}

// 校验
var validate = validator.New()

// v - 负责接收参数的对象
// check - 是否校验
func (ctx *Context) GetParams(v interface{}, check bool) bool {
	tracker := ctx.GetTracker()
	// 根据方法获取参数
	// GET  -   query params
	// POST/PUT/DELETE  - body param
	method := ctx.Method()
	switch method {
	case "GET":
		if err := ctx.ReadQuery(v); err != nil {
			tracker.Error(err)
			ctx.JSON(result.IllegalParam)
			return false
		}
	case "POST", "PUT", "DELETE":
		contentType := ctx.GetHeader("Content-Type")
		if strings.Contains(contentType, "application/json") {
			if err := ctx.ReadJSON(v); err != nil {
				tracker.Error(err)
				ctx.JSON(result.IllegalParam)
				return false
			}
		} else if strings.Contains(contentType, "application/x-www-form-urlencoded") {
			if err := ctx.ReadForm(v); err != nil {
				tracker.Error(err)
				ctx.JSON(result.IllegalParam)
				return false
			}
		}
	}
	if check {
		if err := validate.Struct(v); err != nil {
			tracker.Error(err)
			ctx.JSON(result.IllegalParam)
			return false
		}
	}
	return true
}

/**
jwt 通用
iss: 签发者
sub: 面向的用户
aud: 接收方
exp: 过期时间
nbf: 生效时间
iat: 签发时间
jti: 唯一身份标识
*/
func (ctx *Context) GetToken(m map[string]interface{}) string {
	itv := common.Cfg.GetInt("jwt.itv") // 时间间隔
	now := time.Now()
	m["iat"] = now.Unix()
	m["nbf"] = now.Unix()
	m["exp"] = now.Add(time.Second * time.Duration(itv)).Unix()
	m["iss"] = common.Cfg.GetString("author")
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(m))
	tokenString, _ := token.SignedString([]byte(common.Cfg.GetString("jwt.secret")))
	return tokenString
}

func (ctx *Context) GetTracker() trace.ITracker {
	return ctx.Values().Get(trace.TRACKER).(*trace.Tracker)
}
