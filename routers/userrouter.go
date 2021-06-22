package routers

import (
	"FLightening/middlewares"
	"FLightening/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func mountUserRouters(router *gin.RouterGroup) {
	userRouter := router.Group("/user")
	userRouter.Use(middlewares.AuthRequired())
	userRouter.POST("/me", getInfo)
}

func getInfo(c *gin.Context) {
	jwt, _ := services.GetJwtByContext(c)
	u, ok := services.GetUserByJwt(jwt)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": -1,
			"msg":  "未知错误，无法获取用户",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"username":  u.Username,
			"email":     u.Email,
			"phone":     u.Phone,
			"validated": u.Validated,
			"blocked":   u.Blocked,
		},
	})
}
