package routers

import (
	"FLightening/middlewares"
	"FLightening/models"
	"FLightening/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func mountOrderRouter(router *gin.RouterGroup) {
	g := router.Group("order")
	g.Use(middlewares.AuthRequired())
	g.POST("/book", bookTicket)
	g.GET("/mine", fetchOrders)
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
			"msg":  "订票成功",
			"data": gin.H{
				"price":  o.Price,
				"status": o.Status,
				"time":   o.Time,
				"uid":    o.GetUniqueID(),
			},
		})
	}
}

func fetchOrders(c *gin.Context) {
	u, _ := services.GetUserByContext(c)
	pagestr := c.Query("page")
	var page uint
	if pagestr == "" {
		page = 1
	} else {
		p, err := strconv.ParseUint(pagestr, 10, 16)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": -1,
				"msg":  "页号有误",
			})
			return
		}
		page = uint(p)
	}

	orders, err := models.FetchOrders(u.Id, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  -2,
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": orders,
	})
}
