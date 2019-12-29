package middleware

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12/context"
	"time"
)

func Cors() context.Handler {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Content-Type"},
		AllowedMethods: []string{"GET", "POST"},
		ExposedHeaders: []string{"X-Header"},
		MaxAge:         int((24 * time.Hour).Seconds()),
		// Debug:          true,
	})
}
