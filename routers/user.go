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
	userRouter.GET("/me", getInfo)
	userRouter.PUT("/me", updateInfo)
}

func getInfo(c *gin.Context) {
	u, ok := services.GetUserByContext(c)

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

func updateInfo(c *gin.Context) {
	u, ok := services.GetUserByContext(c)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "无法找到用户",
		})
		return
	}

	body := make(map[string]interface{})
	c.BindJSON(&body)

	code, res := services.UpdateInfo(body, u.Id)
	c.JSON(code, res)
}
