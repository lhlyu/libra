package middleware

import (
	"fmt"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/lhlyu/libra/common"
	"github.com/lhlyu/yutil"
)

/**
Authorization:bearer xxxxxxxxxxx
*/
func Jwt() iris.Handler {
	return func(ctx iris.Context) {
		ip := yutil.ClientIp(ctx.Request())
		var err error
		if err = jwt.New(jwt.Config{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte(common.Cfg.GetString("jwt.secret")), nil
			},
			SigningMethod: jwt.SigningMethodHS256,
			Expiration:    true,
		}).CheckJWT(ctx); err == nil {
			tokens, _ := ctx.Values().Get("jwt").(*jwt.Token)
			tokenMap := tokens.Claims.(jwt.MapClaims)
			fmt.Println(ip, tokenMap)
		}
		ctx.Next()
	}
}
