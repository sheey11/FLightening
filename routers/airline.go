package routers

import (
	"FLightening/models"
	"FLightening/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func mountAirlineRouter(router *gin.RouterGroup) {
	a := router.Group("airline")
	a.GET("/search", searchAirline)
	a.GET("/cities", getAllCities)
}

func searchAirline(c *gin.Context) {
	origin, _ := strconv.Atoi(c.Query("origin"))
	dest, _ := strconv.Atoi(c.Query("dest"))
	page, _ := strconv.Atoi(c.Query("page"))

	if origin == 0 || dest == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"result": nil,
		})
		return
	}

	if page < 1 {
		page = 1
	}

	result := services.FindAirline(origin, dest, page)
	fetchShifts(&result)

	if len(result) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code":   -1,
			"result": result,
			"msg":    "未找到相关航线",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":   200,
			"result": result,
		})
	}
}

func fetchShifts(al *[]models.Airline) {
	for i := range *al {
		(*al)[i].Shifts = services.FindNearestNShifts(1, (*al)[i].GetId())
	}
}

func getAllCities(c *gin.Context) {
	result := services.GetAllCities()
	if len(result) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code":   -1,
			"result": result,
			"msg":    "没有城市",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":   200,
			"result": result,
		})
	}
}
