package sqlconn

import (
	"database/sql"
	"errors"

	"github.com/doug-martin/goqu/v9"
)

func FindUserByUsername(username string) (*sql.Row, error) {
	if db == nil {
		return nil, errors.New("database is not connected")
	}
	if !UserExists(goqu.Ex{"username": username}) {
		return nil, errors.New("用户不存在")
	}

	_sql, _, _ := dialect.From("users").Where(goqu.C("username").Eq(username)).Select("*").ToSQL()
	return db.QueryRow(_sql), nil
}

func FindUserById(id int) (*sql.Row, error) {
	if db == nil {
		return nil, errors.New("database is not connected")
	}
	if !UserExists(goqu.Ex{"id": id}) {
		return nil, errors.New("用户不存在")
	}

	_sql, _, _ := dialect.From("users").Where(goqu.C("id").Eq(id)).Select("*").ToSQL()
	return db.QueryRow(_sql), nil
}

func UserExists(condition goqu.Ex) bool {
	if db == nil {
		return false
	}

	_sql, _, _ := dialect.From("users").Select(goqu.COUNT("*")).Where(condition).ToSQL()

	q := db.QueryRow(_sql)

	var count int
	q.Scan(&count)

	// fmt.Println("count", count, _sql)

	if count == 0 {
		return false
	} else {
		return true
	}
}

func UsernameExists(username string) bool {
	return UserExists(goqu.Ex{
		"username": username,
	})
}

func PhoneExists(phone string) bool {
	return UserExists(goqu.Ex{
		"phone": phone,
	})
}
func EmailExists(email string) bool {
	return UserExists(goqu.Ex{
		"email": email,
	})
}

type dbuser struct {
	Username  string `db:"username"`
	Password  string `db:"password"`
	Salt      string `db:"salt"`
	Phone     string `db:"phone"`
	Email     string `db:"email"`
	Validated int    `db:"validated"`
	Blocked   int    `db:"blocked"`
}

func AddUser(username, password, salt, phone, email string) error {
	if db == nil {
		return errors.New("database is not connected")
	}
	if UsernameExists(username) {
		return errors.New("用户名已被占用")
	}
	if PhoneExists(phone) {
		return errors.New("手机号码已被注册")
	}
	if EmailExists(email) {
		return errors.New("邮箱已被注册")
	}

	sql, _, _ := dialect.From("users").Insert().Rows(
		dbuser{username, password, salt, phone, email, 0, 0},
	).ToSQL()

	_, err := db.Exec(sql)
	return err
}
