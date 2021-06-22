package main

import (
	"FLightening/middlewares"
	"FLightening/routers"
	"FLightening/sqlconn"

	"github.com/gin-gonic/gin"
)

func main() {
	defer sqlconn.Close()

	// setup webserver
	router := gin.Default()

	router.Use(middlewares.Cors())

	v1 := router.Group("/v1")
	routers.MountRouters(v1)

	router.Run()
}
