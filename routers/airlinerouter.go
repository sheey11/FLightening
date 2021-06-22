package routers

import "github.com/gin-gonic/gin"

var airlineChan = make(chan *gin.Context)
var airlineResultChan = make(chan *gin.Context)

func init() {
	go searchAirlineHandle()
}

func mountAirlineRouter(router *gin.RouterGroup) {
	g := router.Group("airlne")
	g.GET("/search", searchAirline)
}

func searchAirlineHandle() {
	for {
		//c := <-airlineChan

	}
}

func searchAirline(c *gin.Context) {
	airlineChan <- c

}
