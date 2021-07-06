package routers

import (
	"FLightening/services"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int
	Result     gin.H
}

func mountAuthRouters(router *gin.RouterGroup) {
	userRouter := router.Group("/auth")
	userRouter.POST("/signin", auth)
	userRouter.POST("/signup", register)
}

func auth(c *gin.Context) {
	dto := LoginDTO{}
	c.BindJSON(&dto)

	statusCode, ret := services.Login(dto.Username, dto.Password)
	c.JSON(statusCode, ret)
}

func register(c *gin.Context) {
	dto := RegisterDTO{}
	c.BindJSON(&dto)

	statusCode, ret := services.Register(dto.Username, dto.Password, dto.Phone, dto.Email)
	c.JSON(statusCode, ret)
}
