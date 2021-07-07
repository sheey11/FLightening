package routers

import (
	"FLightening/middlewares"
	"FLightening/models"
	"FLightening/services"
	"FLightening/sqlconn"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func mountOrderRouter(router *gin.RouterGroup) {
	g := router.Group("order")
	g.Use(middlewares.AuthRequired())
	g.POST("/book", bookTicket)
	g.GET("/mine", fetchOrders)
	g.POST("/complete", markAsComplete)
	g.POST("/cancel", markAsCanceled)
	g.GET("/:id", orderDetail)
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
				"id":     o.GetUniqueID(),
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

func markAsComplete(c *gin.Context) {
	uidStr := c.Query("uid")
	if len(uidStr) != 28 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "订单号错误",
		})
		return
	}
	uid, err := strconv.Atoi(uidStr[23:])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "订单号错误",
		})
		return
	}

	u, _ := services.GetUserByContext(c)
	err = models.MarkAsComplete(uid, u.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -2,
			"msg":  "更新订单时出现错误",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功",
	})
}

func markAsCanceled(c *gin.Context) {
	oidStr := c.Query("uid")
	if len(oidStr) != 28 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "订单号错误",
		})
		return
	}
	oid, err := strconv.Atoi(oidStr[23:])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "订单号错误",
		})
		return
	}

	u, _ := services.GetUserByContext(c)
	err = models.MarkAsCanceled(oid, u.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -2,
			"msg":  "更新订单时出现错误",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功",
	})
}

func orderDetail(c *gin.Context) {
	oidStr := c.Param("id")
	if len(oidStr) != 28 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "订单号错误",
		})
		return
	}
	oid, err := strconv.Atoi(oidStr[23:])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "订单号错误",
		})
		return
	}

	r, err := sqlconn.FindPassenger(oid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "查找乘客时出现错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": r,
	})
}
