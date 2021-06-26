package routers

import "github.com/gin-gonic/gin"

func MountRouters(router *gin.RouterGroup) {
	mountAuthRouters(router)
	mountUserRouters(router)
	mountAirlineRouter(router)
}
