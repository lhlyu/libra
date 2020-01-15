package controller

import (
    "github.com/iris-contrib/middleware/jwt"
    "github.com/kataras/iris/v12"
    "github.com/lhlyu/libra/common"
    "github.com/lhlyu/libra/response"
    "time"
)

type controller struct {

}

func (c controller) getParams(ctx iris.Context, v interface{}) bool {
	// 根据方法获取参数
	// GET  -   query params
	// POST/PUT/DELETE  - body param
	method := ctx.Method()
	switch method {
	case "GET":
		if err := ctx.ReadQuery(v); err != nil {
			ctx.JSON(response.IllegalParam)
			return false
		}
	case "POST", "PUT", "DELETE":
		contentType := ctx.GetHeader("Content-Type")
		switch contentType {
		case "application/json":
			if err := ctx.ReadJSON(v); err != nil {
				ctx.JSON(response.IllegalParam)
				return false
			}
		case "application/x-www-form-urlencoded":
			if err := ctx.ReadForm(v); err != nil {
				ctx.JSON(response.IllegalParam)
				return false
			}
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
func (c controller) getToken(ctx iris.Context, m map[string]interface{}) string {
	itv := common.Cfg.GetInt("jwt.itv")  // 时间间隔
	now := time.Now()
	m["iat"] = now.Unix()
	m["nbf"] = now.Unix()
	m["exp"] = now.Add(time.Second * time.Duration(itv)).Unix()
	m["iss"] = common.Cfg.GetString("author")
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(m))
	tokenString, _ := token.SignedString([]byte(common.Cfg.GetString("jwt.secret")))
	return tokenString
}
