package sqlconn

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

func FetchAllProvince() (*sql.Rows, error) {
	sql, _, _ := dialect.Select(
		goqu.I("province.id").As("id"),
		goqu.I("province.name").As("name"),
	).From(
		"province",
	).ToSQL()

	rows, err := db.Query(sql)
	return rows, err
}

func AddProvince(name string) error {
	_sql, _, _ := dialect.Insert("province").Rows(
		goqu.Record{
			"name": name,
		},
	).ToSQL()

	_, err := db.Exec(_sql)
	return err
}

func UpdateProvince(id int, name string) error {
	_sql, _, _ := dialect.Update("cities").Set(
		goqu.Record{
			"name": name,
		},
	).Where(goqu.Ex{
		"id": id,
	}).ToSQL()

	_, err := db.Exec(_sql)
	return err
}
