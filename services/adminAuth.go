package services

import (
	"FLightening/models"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

var adminSecretKey = []byte(If(os.Getenv("ADMIN_JWT_SECRET") == "", "BNMnlj6G6DXK1ZqPF3", os.Getenv("ADMIN_JWT_SECRET")).(string))

func CreateAdminJwt(payload map[string]interface{}) (string, error) {
	token := jwt.New()

	token.Set(jwt.IssuerKey, `sheey`)
	token.Set(jwt.AudienceKey, `FLighteningAdmin`)
	token.Set(jwt.IssuedAtKey, time.Now().Unix())
	token.Set(jwt.ExpirationKey, time.Now().AddDate(0, 0, 30).Unix())

	for s := range payload {
		token.Set(s, payload[s])
	}

	jwtStr, err := jwt.Sign(token, jwa.HS256, secretKey)
	return string(jwtStr), err
}

func VerifyAdminJwt(token string) bool {
	j, e := jwt.ParseString(
		token,
		jwt.WithValidate(true),
		jwt.WithVerify(jwa.HS256, secretKey),
	)
	if e != nil {
		return false
	}
	if j.Issuer() != "sheey" {
		return false
	}
	if len(j.Audience()) > 0 && j.Audience()[0] != "FLighteningAdmin" {
		return false
	}
	if j.Expiration().Before(time.Now()) {
		return false
	}

	return true
}

func GetAdminByJwt(token string) (*models.Admin, bool) {
	j, e := jwt.ParseString(
		token,
		jwt.WithValidate(true),
		jwt.WithVerify(jwa.HS256, secretKey),
	)
	if e != nil {
		return nil, false
	}
	uid, ok := j.Get("uid")
	if !ok {
		return nil, false
	}

	u, err := models.FindAdminById(int(uid.(float64)))
	if err != nil {
		return nil, false
	}

	return &u, true
}

func GetAdminJwtByContext(c *gin.Context) (string, bool) {
	token := c.GetHeader("Authorization")

	if len(token) > 7 && token[:6] == "Bearer" {
		jwt := token[strings.Index(token, " ")+1:]
		return jwt, true
	}

	return "", false
}

func GetAdminByContext(c *gin.Context) (*models.Admin, bool) {
	jwt, ok := GetAdminJwtByContext(c)
	if !ok {
		return nil, false
	}
	return GetAdminByJwt(jwt)
}
