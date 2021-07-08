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

func mountAdminRouters(router *gin.RouterGroup) {
	adminRouter := router.Group("/admin")

	adminRouter.POST("/signin", adminAuth)
	adminRouter.Use(middlewares.AdminAuthRequired())
	adminRouter.GET("/orders", fetchAllOrders)
	adminRouter.GET("/users", fetchAllUsers)
	adminRouter.GET("/shifts", fetchAllShifts)
	adminRouter.GET("/airlines", fetchAllAirlines)
	adminRouter.GET("/province", fetchAllProvince)

	adminRouter.GET("/orders/:id", orderDetail)
	adminRouter.GET("/orders/filter/user/:id", orderDetail)

	adminRouter.GET("/users/:id", getUserInfo)
	adminRouter.POST("/users/:id/block", blockUser)
	adminRouter.POST("/users/:id/unblock", unblockUser)

	adminRouter.POST("/city", addCity)
	adminRouter.POST("/province", addProvince)

	adminRouter.PUT("/city", updateCity)
	adminRouter.PUT("/province", updateProvince)
}

func adminAuth(c *gin.Context) {
	dto := LoginDTO{}
	c.BindJSON(&dto)

	statusCode, ret := services.AdminLogin(dto.Username, dto.Password)
	c.JSON(statusCode, ret)
}

func fetchAllProvince(c *gin.Context) {
	ps := models.FetchAllProvince()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": ps,
	})
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

func addCity(c *gin.Context) {
	ct := CityDTO{}
	e := c.BindJSON(&ct)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "请求参数错误",
		})
		return
	}
	e = sqlconn.AddCity(ct.Name, *ct.Province, ct.Code)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "添加失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功",
	})
}

func addProvince(c *gin.Context) {
	ct := ProvinceDTO{}
	c.BindJSON(&ct)
	e := sqlconn.AddProvince(ct.Name)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "检查请求",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功",
	})
}

func updateCity(c *gin.Context) {
	ct := CityUpdateDTO{}
	err := c.BindJSON(&ct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "请检查请求",
		})
		return
	}

	err = sqlconn.UpdateCity(ct.Id, ct.Name, ct.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -2,
			"msg":  "更新失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "好了",
	})
}

func updateProvince(c *gin.Context) {
	ct := ProvinceUpdateDTO{}
	err := c.BindJSON(&ct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "请检查请求",
		})
		return
	}

	err = sqlconn.UpdateProvince(ct.Id, ct.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -2,
			"msg":  "更新失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "好了",
	})
}

func blockUser(c *gin.Context) {
	uidStr := c.Param("id")
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "检查请求",
			"err":  err,
		})
		return
	}
	err = sqlconn.BlockUser(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -2,
			"msg":  "更新失败",
			"err":  err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功",
	})
}

func unblockUser(c *gin.Context) {
	uidStr := c.Param("id")
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "检查请求",
			"err":  err,
		})
		return
	}
	err = sqlconn.UnblockUser(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -2,
			"msg":  "更新失败",
			"err":  err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功",
	})
}
