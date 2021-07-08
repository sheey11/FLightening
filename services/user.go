package services

import (
	"FLightening/models"
	"FLightening/sqlconn"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func genJwt(u models.User) (string, error) {
	payload := map[string]interface{}{
		"username": u.Username,
		"uid":      u.Id,
	}
	return CreateJwt(payload)
}

func Login(username string, password string) (int, gin.H) {
	u, err := models.FindUserByUsername(username)
	if err != nil {
		return http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		}
	}

	auth := models.CheckPassword(password, u)

	if auth && !bool(u.Blocked) {
		jwt, _ := genJwt(u)
		return 200, gin.H{
			"code": 0,
			"msg":  "登陆成功",
			"jwt":  jwt,
		}
	} else if bool(u.Blocked) {
		return http.StatusUnauthorized, gin.H{
			"code": -2,
			"msg":  "用户已被封禁",
		}
	}

	return http.StatusUnauthorized, gin.H{
		"code": -3,
		"msg":  "密码错误",
	}
}

func Register(username string, password string, phone string, email string) (int, gin.H) {
	phoneMatch, _ := regexp.MatchString("^\\d{11}$", phone)
	if !phoneMatch {
		return http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "手机号格式不正确",
		}
	}

	emailMatch, _ := regexp.MatchString("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)\\])", email)
	if !emailMatch {
		return http.StatusBadRequest, gin.H{
			"code": -2,
			"msg":  "邮箱格式不正确",
		}
	}

	usernameMatch, _ := regexp.MatchString("^[a-zA-Z][a-zA-Z\\d]{5,11}$", username)
	if !usernameMatch {
		return http.StatusBadRequest, gin.H{
			"code": -3,
			"msg":  "用户名格式不正确",
		}
	}

	passwordMatch, _ := regexp.MatchString("^[a-zA-Z\\d]{8,32}$", password)
	if !passwordMatch {
		return http.StatusBadRequest, gin.H{
			"code": -4,
			"msg":  "密码格式不正确",
		}
	}

	err := models.CreateUser(username, password, phone, email)
	if err != nil {
		errStr := err.Error()
		if strings.Contains(errStr, "SQL") {
			errStr = "内部错误"
		}
		return http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  errStr,
		}
	}

	u, _ := models.FindUserByUsername(username)

	jwt, err := genJwt(u)

	if err != nil {
		return 500, gin.H{
			"code": -1,
			"msg":  err.Error(),
		}
	}

	return 200, gin.H{
		"code": 200,
		"msg":  "注册成功",
		"jwt":  jwt,
	}
}

func UpdateInfo(info map[string]interface{}, id int) (int, gin.H) {
	email, o := info["email"]
	if !o {
		email = ""
	}
	phone, o := info["phone"]
	if !o {
		phone = ""
	}
	err := sqlconn.UpdateInfo(email.(string), phone.(string), id)
	if err == nil {
		return http.StatusOK, gin.H{
			"code": 200,
			"msg":  "已保存",
		}
	} else {
		if strings.Contains(err.Error(), "SQL") {
			return http.StatusBadRequest, gin.H{
				"code": -2,
				"msg":  "未知错误",
			}
		} else {
			return http.StatusBadRequest, gin.H{
				"code": -1,
				"msg":  err.Error(),
			}
		}
	}
}
