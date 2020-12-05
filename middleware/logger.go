package middleware

import (
	"github.com/flxxyz/ono/conf"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	if conf.DefaultAppConfig().Debug {
		return gin.Logger()
	} else {
		return func(context *gin.Context) {
			context.Next()
		}
	}
}
