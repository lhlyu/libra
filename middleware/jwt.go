package middleware

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/lhlyu/libra/common"
	"github.com/lhlyu/yutil"
)

/**
Authorization:bearer xxxxxxxxxxx
*/
func Jwt() context.Handler {
	return func(ctx iris.Context) {
		user := &common.XUser{}
		user.Ip = yutil.ClientIp(ctx.Request())
		var err error
		if err = jwt.New(jwt.Config{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte(common.Cfg.GetString("jwt.secret")), nil
			},
			SigningMethod: jwt.SigningMethodHS256,
			Expiration:    true,
		}).CheckJWT(ctx); err == nil {
			//token, _ := jwt.FromAuthHeader(ctx)
			tokens, _ := ctx.Values().Get("jwt").(*jwt.Token)
			tokenMap := tokens.Claims.(jwt.MapClaims)
			user.Id = int(tokenMap[common.X_ID].(float64))
			user.Role = int(tokenMap[common.X_ROLE].(float64))
		}
		ctx.Values().Set(common.X_USER, user)
		ctx.Next()
	}
}
