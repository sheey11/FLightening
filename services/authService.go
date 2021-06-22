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

var secretKey = []byte(If(os.Getenv("JWT_SECRET") == "", "8yMDA1LzA1L2lkZW50", os.Getenv("JWT_SECRET")).(string))

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	} else {
		return falseVal
	}
}

func CreateJwt(payload map[string]interface{}) (string, error) {
	token := jwt.New()

	token.Set(jwt.IssuerKey, `sheey`)
	token.Set(jwt.AudienceKey, `FLighteningUser`)
	token.Set(jwt.IssuedAtKey, time.Now().Unix())
	token.Set(jwt.ExpirationKey, time.Now().AddDate(0, 0, 30).Unix())

	for s := range payload {
		token.Set(s, payload[s])
	}

	jwtStr, err := jwt.Sign(token, jwa.HS256, secretKey)
	return string(jwtStr), err
}

func VerifyJwt(token string) bool {
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
	if len(j.Audience()) > 0 && j.Audience()[0] != "FLighteningUser" {
		return false
	}
	if j.Expiration().Before(time.Now()) {
		return false
	}

	return true
}

func GetUserByJwt(token string) (*models.User, bool) {
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

	u, err := models.FindUserById(int(uid.(float64)))
	if err != nil {
		return nil, false
	}

	return &u, true
}

func GetJwtByContext(c *gin.Context) (string, bool) {
	token := c.GetHeader("Authorization")

	if len(token) > 7 && token[:6] == "Bearer" {
		jwt := token[strings.Index(token, " ")+1:]
		return jwt, true
	}

	return "", false
}
