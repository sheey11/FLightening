package routers

import (
	"FLightening/services"

	"github.com/gin-gonic/gin"
)

var authChan = make(chan *gin.Context)
var registerChan = make(chan *gin.Context)

type Response struct {
	StatusCode int
	Result     gin.H
}

func init() {
	go authHandle()
	go registerHandle()
}

func mountAuthRouters(router *gin.RouterGroup) {
	userRouter := router.Group("/auth")
	userRouter.POST("/signin", auth)
	userRouter.POST("/signup", register)
}

func authHandle() {
	for {
		c := <-authChan
		dto := LoginDTO{}
		c.BindJSON(&dto)

		statusCode, ret := services.Login(dto.Username, dto.Password)
		c.JSON(statusCode, ret)
	}
}

func registerHandle() {
	for {
		c := <-registerChan
		dto := RegisterDTO{}
		c.BindJSON(&dto)

		statusCode, ret := services.Register(dto.Username, dto.Password, dto.Phone, dto.Email)
		c.JSON(statusCode, ret)
	}
}

func auth(c *gin.Context) {
	authChan <- c
}

func register(c *gin.Context) {
	registerChan <- c
}
