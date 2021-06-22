package routers

import (
	"FLightening/services"

	"github.com/gin-gonic/gin"
)

var authChan = make(chan *gin.Context)
var authResultChan = make(chan *Response)
var registerChan = make(chan *gin.Context)
var registerResultChan = make(chan *Response)

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
		authResultChan <- &Response{statusCode, ret}
	}
}

func registerHandle() {
	for {
		c := <-registerChan
		dto := RegisterDTO{}
		c.BindJSON(&dto)

		statusCode, ret := services.Register(dto.Username, dto.Password, dto.Phone, dto.Email)
		registerResultChan <- &Response{statusCode, ret}
	}
}

func auth(c *gin.Context) {
	authChan <- c
	res := <-authResultChan
	c.JSON(res.StatusCode, res.Result)
}

func register(c *gin.Context) {
	registerChan <- c
	res := <-registerResultChan
	c.JSON(res.StatusCode, res.Result)
}
