package middlewares

import (
	"FLightening/logging"

	"github.com/gin-gonic/gin"
)

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		logging.Request.Printf("%s\t|\t%s\t|\t%s\n", c.Request.Host, c.Request.Method, c.Request.URL)
		c.Next()
	}
}
