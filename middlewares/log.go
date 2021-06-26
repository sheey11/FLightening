package middlewares

import (
	"FLightening/logging"

	"github.com/gin-gonic/gin"
)

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		logging.Request.Printf("%s\t%s\t%s", c.Request.Method, c.Request.URL, c.Request.Host)
		c.Next()
	}
}
