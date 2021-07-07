package sqlconn

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

func GetAllCities() (*sql.Rows, error) {
	sql, _, _ := dialect.Select(
		goqu.I("cities.id").As("id"),
		goqu.I("cities.name").As("name"),
		goqu.I("province.name").As("province"),
		goqu.I("cities.code").As("code"),
	).From(
		"cities",
		"province",
	).Where(goqu.Ex{
		"cities.province": goqu.I("province.id"),
	}).ToSQL()

	rows, err := db.Query(sql)
	return rows, err
}

func AddCity(name string, province int, code string) error {
	_sql, _, _ := dialect.Insert("cities").Rows(
		goqu.Record{
			"name":     name,
			"province": province,
			"code":     code,
		},
	).ToSQL()

	_, err := db.Exec(_sql)
	return err
}

func UpdateCity(id int, name string, province int, code string) error {
	_sql, _, _ := dialect.Update("cities").Set(
		goqu.Record{
			"name":     name,
			"province": province,
			"code":     code,
		},
	).Where(goqu.Ex{
		"id": id,
	}).ToSQL()

	_, err := db.Exec(_sql)
	return err
}
