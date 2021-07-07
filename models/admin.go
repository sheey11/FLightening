package models

import (
	"FLightening/sqlconn"
	"errors"
)

type Admin struct {
	Id        int
	Username  string
	password  string
	salt      string
	Phone     string
	Email     string
	Validated BitBool
	Blocked   BitBool
}

func CheckAdminCredential(username string, password string, user Admin) bool {
	return user.Username == username && CheckAdminPassword(password, user)
}

func CheckAdminPassword(pass string, user Admin) bool {
	password := EncryptPassword(pass, user.salt)
	return user.password == password
}

func FindAdminByUsername(username string) (Admin, error) {
	if username == "" {
		return Admin{}, errors.New("请输入用户名")
	}
	q, err := sqlconn.FindAdminByUsername(username)

	if err != nil {
		return Admin{}, err
	}

	var u Admin
	err = q.Scan(&u.Id, &u.Username, &u.password, &u.salt, &u.Phone, &u.Email, &u.Validated, &u.Blocked)

	return u, err
}

func FindAdminById(id int) (Admin, error) {
	q, err := sqlconn.FindAdminById(id)

	if err != nil {
		return Admin{}, err
	}

	var u Admin
	err = q.Scan(&u.Id, &u.Username, &u.password, &u.salt, &u.Phone, &u.Email, &u.Validated, &u.Blocked)

	return u, err
}

func FetchAllUsers(page uint) []User {
	rows, err := sqlconn.FetchAllUsers(page)

	if err != nil {
		return make([]User, 0)
	}
	defer rows.Close()

	us := make([]User, 0)
	for rows.Next() {
		u := User{}
		rows.Scan(&u.Id, &u.Username, &u.password, &u.salt, &u.Phone, &u.Email, &u.Validated, &u.Blocked)
	}

	return us
}
