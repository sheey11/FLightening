package sqlconn

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

func FindNearestNShifts(n uint, airline int) (*sql.Rows, error) {
	if n == 0 {
		n = 1
	}

	_sql, _, _ := dialect.Select(
		"scheduled_takeoff",
		"scheduled_landing",
		"actual_takeoff",
		"actual_landing",
		"status",
		"economy_price",
		"premium_price",
		"business_price",
		"first_price",
	).
		From("shifts").
		Limit(n).
		Where(goqu.Ex{
			"airline": airline,
		}).
		ToSQL()
	return db.Query(_sql)
}

func ShiftExists(id int) bool {
	_sql, _, _ := dialect.Select("COUNT(*)").From("shifts").Where(goqu.Ex{
		"id": id,
	}).ToSQL()

	r := db.QueryRow(_sql)

	c := 0
	r.Scan(&c)
	return c == 1
}

func LookUpPrice(cabin int, id int) int {
	var sqls *goqu.SelectDataset

	switch cabin {
	case 0:
		sqls = dialect.Select("economy_price")
	case 1:
		sqls = dialect.Select("premium_price")
	case 2:
		sqls = dialect.Select("business_price")
	default:
		sqls = dialect.Select("remaining_seat")
	}

	_sql, _, _ := sqls.From("shifts").Where(goqu.Ex{
		"id": id,
	}).ToSQL()

	r := db.QueryRow(_sql)

	price := 0
	r.Scan(&price)
	return price
}
