package middlewares

import (
	"FLightening/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt, ok := services.GetJwtByContext(c)
		if !ok || !services.VerifyJwt(jwt) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "请先登录",
			})
			c.Abort()
		}
		c.Next()
	}
}
