package sqlconn

import (
	"database/sql"
	"errors"

	"github.com/doug-martin/goqu/v9"
)

func FindAdminByUsername(username string) (*sql.Row, error) {
	if db == nil {
		return nil, errors.New("database is not connected")
	}
	if !UserExists(goqu.Ex{"username": username}) {
		return nil, errors.New("用户不存在")
	}

	_sql, _, _ := dialect.From("admin").Where(goqu.C("username").Eq(username)).Select("*").ToSQL()
	return db.QueryRow(_sql), nil
}

func FindAdminById(id int) (*sql.Row, error) {
	if db == nil {
		return nil, errors.New("database is not connected")
	}
	if !UserExists(goqu.Ex{"id": id}) {
		return nil, errors.New("用户不存在")
	}

	_sql, _, _ := dialect.From("admin").Where(goqu.C("id").Eq(id)).Select("*").ToSQL()
	return db.QueryRow(_sql), nil
}
