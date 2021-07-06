package routers

import (
	"FLightening/middlewares"
	"FLightening/models"
	"FLightening/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func mountOrderRouter(router *gin.RouterGroup) {
	g := router.Group("order")
	g.Use(middlewares.AuthRequired())
	g.POST("/book", bookTicket)
}

func bookTicket(c *gin.Context) {
	u, _ := services.GetUserByContext(c)

	dto := BookingDTO{}
	e := c.ShouldBindJSON(&dto)

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "请求参数有问题",
		})
	}

	oid, e := services.BookTicket(dto.Cabin, u.Id, dto.Shift, dto.Passenger)

	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  e.Error(),
		})
	} else {
		o := models.FindOrderById(oid)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "芜湖",
			"data": gin.H{
				"price":  o.Price,
				"status": o.Status,
				"time":   o.Time,
				"uid":    o.GetUniqueID(),
			},
		})
	}
}
