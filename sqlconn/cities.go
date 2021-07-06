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
