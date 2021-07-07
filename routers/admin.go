package routers

import (
	"FLightening/middlewares"
	"FLightening/models"
	"FLightening/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func mountAdminRouters(router *gin.RouterGroup) {
	adminRouter := router.Group("/admin")

	adminRouter.POST("/signin", adminAuth)
	adminRouter.Use(middlewares.AdminAuthRequired())
	adminRouter.GET("/orders", fetchAllOrders)
	adminRouter.GET("/users", fetchAllUsers)
	adminRouter.GET("/shifts", fetchAllShifts)

	adminRouter.GET("/orders/:id", orderDetail)
	adminRouter.GET("/users/:id", getUserInfo)
}

func adminAuth(c *gin.Context) {
	dto := LoginDTO{}
	c.BindJSON(&dto)

	statusCode, ret := services.AdminLogin(dto.Username, dto.Password)
	c.JSON(statusCode, ret)
}

func fetchAllOrders(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	orders, e := models.FetchAllOrders(uint(page))
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "获取订单时出现错误",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": orders,
		})
	}
}

func fetchAllUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	users := models.FetchAllUsers(uint(page))
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": users,
	})
}

func fetchAllShifts(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	shifts := models.FetchAllShifts(uint(page))
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": shifts,
	})
}

func fetchAllAirlines(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	airlines := models.FetchAllAirlines(uint(page))
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": airlines,
	})
}

func getUserInfo(c *gin.Context) {
	uidStr := c.Param("id")
	uid, _ := strconv.Atoi(uidStr)
	u, e := models.FindUserById(uid)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "无法获取用户",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"username":  u.Username,
			"email":     u.Email,
			"phone":     u.Phone,
			"validated": u.Validated,
			"blocked":   u.Blocked,
		},
	})
}
