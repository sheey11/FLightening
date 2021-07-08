package models

import (
	"FLightening/sqlconn"
	"crypto/sha1"
	"database/sql/driver"
	"encoding/hex"
	"errors"
	"math/rand"
)

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	password  string
	salt      string
	Phone     string  `json:"phone"`
	Email     string  `json:"email"`
	Validated BitBool `json:"validated"`
	Blocked   BitBool `json:"blocked"`
}

type BitBool bool

func (b BitBool) Value() (driver.Value, error) {
	if b {
		return []byte{1}, nil
	} else {
		return []byte{0}, nil
	}
}
func (b *BitBool) Scan(src interface{}) error {
	v, ok := src.([]byte)
	if !ok {
		return errors.New("bad []byte type assertion")
	}
	*b = v[0] == 1
	return nil
}

func GenSalt() string {
	const dict = "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz1234567890123456789012345678901234567890"
	salt := ""
	for i := 0; i < 10; i++ {
		randIndex := rand.Intn(len(dict))
		salt += string(dict[randIndex])
	}
	return salt
}

func EncryptPassword(pass string, salt string) string {
	hash := sha1.New().Sum([]byte(pass + salt))
	str := hex.EncodeToString(hash)
	return str
}

func CheckCredential(username string, password string, user User) bool {
	return user.Username == username && CheckPassword(password, user)
}

func CheckPassword(pass string, user User) bool {
	password := EncryptPassword(pass, user.salt)
	return user.password == password
}

func (u User) New(username string, password string, phone string, email string) User {
	salt := GenSalt()
	return User{
		Username:  username,
		salt:      salt,
		password:  EncryptPassword(password, salt),
		Phone:     phone,
		Email:     email,
		Validated: false,
		Blocked:   false,
	}
}

func FindUserByUsername(username string) (User, error) {
	if username == "" {
		return User{}, errors.New("请输入用户名")
	}
	q, err := sqlconn.FindUserByUsername(username)

	if err != nil {
		return User{}, err
	}

	var u User
	err = q.Scan(&u.Id, &u.Username, &u.password, &u.salt, &u.Phone, &u.Email, &u.Validated, &u.Blocked)

	return u, err
}

func FindUserById(id int) (User, error) {
	q, err := sqlconn.FindUserById(id)

	if err != nil {
		return User{}, err
	}

	var u User
	err = q.Scan(&u.Id, &u.Username, &u.password, &u.salt, &u.Phone, &u.Email, &u.Validated, &u.Blocked)

	return u, err
}

func CreateUser(username, rawPassword, phone, email string) error {
	salt := GenSalt()
	password := EncryptPassword(rawPassword, salt)
	return sqlconn.AddUser(username, password, salt, phone, email)
}
