package routers

import (
	"FLightening/services"

	"github.com/gin-gonic/gin"
)

var airlineChan = make(chan *gin.Context)

func init() {
	go searchAirlineHandle()
}

func mountAirlineRouter(router *gin.RouterGroup) {
	g := router.Group("airlne")
	g.GET("/search", searchAirline)
}

func searchAirlineHandle() {
	for {
		c := <-airlineChan
		dto := AirlineSearchDTO{}
		c.ShouldBindJSON(&dto)

		result := services.FindAirline(dto.Origin, dto.Destination, dto.Page)
		c.JSON(200, gin.H{
			"code":   200,
			"result": result,
		})
	}
}

func searchAirline(c *gin.Context) {
	airlineChan <- c
}
