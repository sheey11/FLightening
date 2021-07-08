package services

import (
	"FLightening/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func genAdminJwt(u models.Admin) (string, error) {
	payload := map[string]interface{}{
		"username": u.Username,
		"uid":      u.Id,
	}
	return CreateAdminJwt(payload)
}

func AdminLogin(username string, password string) (int, gin.H) {
	u, err := models.FindAdminByUsername(username)
	if err != nil {
		return http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		}
	}

	auth := models.CheckAdminPassword(password, u)

	if auth {
		jwt, _ := genAdminJwt(u)
		return 200, gin.H{
			"code": 0,
			"msg":  "登陆成功",
			"jwt":  jwt,
		}
	}

	return http.StatusUnauthorized, gin.H{
		"code": -2,
		"msg":  "密码错误",
	}
}
