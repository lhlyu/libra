package controller

import (
    "github.com/iris-contrib/middleware/jwt"
    "github.com/kataras/iris/v12"
    "github.com/lhlyu/libra/common"
    "github.com/lhlyu/libra/errcode"
    "gopkg.in/go-playground/validator.v9"
    "strings"
    "time"
)

var validate = validator.New()

type controller struct {
	common.BaseController
}

func (c controller) getParams(ctx iris.Context, v interface{}, check bool) bool {
	// 根据方法获取参数
	// GET  -   query params
	// POST/PUT/DELETE  - body param
	method := ctx.Method()
	switch method {
	case "GET":
		if err := ctx.ReadQuery(v); err != nil {
			c.Error(c.GetTraceId(ctx), err)
			ctx.JSON(errcode.IllegalParam)
			return false
		}
	case "POST", "PUT", "DELETE":
		contentType := ctx.GetHeader("Content-Type")
		switch contentType {
		case "application/json":
			if err := ctx.ReadJSON(v); err != nil {
				c.Error(c.GetTraceId(ctx), err)
				ctx.JSON(errcode.IllegalParam)
				return false
			}
		case "application/x-www-form-urlencoded":
			if err := ctx.ReadForm(v); err != nil {
				c.Error(c.GetTraceId(ctx), err)
				ctx.JSON(errcode.IllegalParam)
				return false
			}
		}
	}
	if !check {
		return true
	}
	if err := validate.Struct(v); err != nil {
		c.Error(c.GetTraceId(ctx), err)
		ctx.JSON(errcode.IllegalParam)
		return false
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
func (c controller) getToken(ctx iris.Context, user *common.XUser) string {
	//guid := ctx.Values().Get(common.X_TRACE).(string)

	itv := common.Cfg.GetInt("jwt.itv")
	if itv == 0 {
		itv = common.ITV
	}
	m := make(map[string]interface{})
	m[common.X_ID] = user.Id
	m[common.X_ROLE] = user.Role
	now := time.Now()
	m["iat"] = now.Unix()
	m["nbf"] = now.Unix()
	m["exp"] = now.Add(time.Second * time.Duration(itv)).Unix()
	m["iss"] = common.Cfg.GetString("author")
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(m))
	tokenString, _ := token.SignedString([]byte(common.Cfg.GetString("jwt.secret")))
	return tokenString
}

// 检查参数是否大于0
func (controller) checkUInt(v int) *errcode.ErrCode {
	if v <= 0 {
		return errcode.IllegalParam
	}
	return nil
}

// 检查参数是否非空字符串
func (controller) checkEmpty(v string) *errcode.ErrCode {
	if strings.TrimSpace(v) == "" {
		return errcode.IllegalParam
	}
	return nil
}

// 获取用户信息
func (controller) GetUser(ctx iris.Context) *common.XUser {
	return ctx.Values().Get(common.X_USER).(*common.XUser)
}

// 获取日志追踪ID
func (controller) GetTraceId(ctx iris.Context) string {
	return ctx.Values().Get(common.X_TRACE).(string)
}

// 获取日志追踪器
func (c controller) GetTracker(ctx iris.Context) *common.Tracker {
	user := c.GetUser(ctx)
	startTime := ctx.Values().Get(common.X_TIME).(time.Time)
	traceId := c.GetTraceId(ctx)
	agent := ctx.GetHeader(common.USER_AGENT)
	return common.NewTracker(user, startTime, traceId, agent)
}

// 获取agent
func (controller) GetAgent(ctx iris.Context) string {
	return ctx.GetHeader(common.USER_AGENT)
}

// 是否是管理员
func (c controller) IsAdmin(ctx iris.Context) bool {
	user := c.GetUser(ctx)
	return user.Role >= common.PERMISSION
}

// 是否是走管理路由
func (c controller) IsAdminRouter(ctx iris.Context) bool {
	admin, ok := ctx.Values().Get(common.ADMIN).(bool)
	if !ok {
		return false
	}
	return admin
}

type Controller struct {

}
