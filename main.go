package main

import (
	"FLightening/logging"
	"FLightening/middlewares"
	"FLightening/routers"
	"FLightening/sqlconn"

	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/gin-gonic/gin"
)

func main() {
	logging.Info.Printf("FLightening is launching at %s mode...", gin.Mode())
	defer sqlconn.Close()

	// setup webserver
	router := gin.Default()

	router.Use(middlewares.Cors())
	router.Use(middlewares.Logging())

	logging.Info.Println("Mounting routers...")
	v1 := router.Group("/v1")
	routers.MountRouters(v1)

	router.Run()
	logging.Info.Println("FLightening is ready.")
}
