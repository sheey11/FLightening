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
		"id",
		"scheduled_takeoff",
		"scheduled_landing",
		"actual_takeoff",
		"actual_landing",
		"status",
		"economy_price",
		"premium_price",
		"business_price",
		"first_price",
		"remaining_seat",
		"airline",
	).
		From("shifts").
		Limit(n).
		Where(goqu.Ex{
			"airline": airline,
			"scheduled_takeoff": goqu.Op{
				"gt": goqu.Func("NOW"),
			},
		}).
		ToSQL()
	return db.Query(_sql)
}

func FindShiftById(id int) *sql.Row {
	_sql, _, _ := dialect.Select(
		"id",
		"scheduled_takeoff",
		"scheduled_landing",
		"actual_takeoff",
		"actual_landing",
		"status",
		"economy_price",
		"premium_price",
		"business_price",
		"first_price",
		"remaining_seat",
		"airline",
	).
		From("shifts").
		Where(goqu.Ex{
			"id": id,
		}).
		ToSQL()
	return db.QueryRow(_sql)
}

func FetchAllShifts(page uint) (*sql.Rows, error) {
	_sql, _, _ := dialect.Select(
		"shifts.id",
		"shifts.scheduled_takeoff",
		"shifts.scheduled_landing",
		"shifts.actual_takeoff",
		"shifts.actual_landing",
		"shifts.status",
		"shifts.economy_price",
		"shifts.premium_price",
		"shifts.business_price",
		"shifts.first_price",
		"shifts.remaining_seat",
		"shifts.airline",
		"airline.name",
	).
		From("shifts", "airline").
		Where(goqu.Ex{"shifts.airline": goqu.I("airline.id")}).
		Limit(10).
		Offset(uint(page-1) * 10).
		ToSQL()
	return db.Query(_sql)
}

func ShiftExistsAndVacancy(id int) int {
	// ensure exist
	_sql, _, _ := dialect.Select(goqu.COUNT("id")).From("shifts").Where(goqu.Ex{
		"id": id,
	}).ToSQL()

	r := db.QueryRow(_sql)

	c := 0
	r.Scan(&c)
	if c == 0 {
		return -1
	}

	// ensure vacancy
	_sql, _, _ = dialect.Select("remaining_seat").From("shifts").Where(goqu.Ex{
		"id": id,
	}).ToSQL()

	r = db.QueryRow(_sql)
	r.Scan(&c)
	if c == 0 {
		return -2
	}

	return 0
}

func LookUpPrice(cabin int, id int) float32 {
	var sqls *goqu.SelectDataset

	switch cabin {
	case 0:
		sqls = dialect.Select("economy_price")
	case 1:
		sqls = dialect.Select("premium_price")
	case 2:
		sqls = dialect.Select("business_price")
	default:
		sqls = dialect.Select("first_price")
	}

	_sql, _, _ := sqls.From("shifts").Where(goqu.Ex{
		"id": id,
	}).ToSQL()

	r := db.QueryRow(_sql)

	var price float32
	r.Scan(&price)
	return price
}
