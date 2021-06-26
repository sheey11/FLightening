package sqlconn

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

func FindAirlineByOriginAndDest(origin, dest, page int) (*sql.Rows, error) {
	sql, _, _ := dialect.Select(
		goqu.I("airline.name").As("airline_id"),
		goqu.I("am.name").As("model"),
		goqu.I("p1.name").As("origin"),
		goqu.I("p2.name").As("dest"),
		goqu.I("ac.name").As("affiliate"),
		goqu.I("ac.logo_url").As("logo"),
	).From(
		"airline",
		goqu.I("airport").As("p1"),
		goqu.I("airport").As("p2"),
		goqu.I("cities").As("c1"),
		goqu.I("cities").As("c2"),
		goqu.I("airlinecompanies").As("ac"),
		goqu.I("airplanemodel").As("am"),
	).Where(goqu.Ex{
		"airline.origin":      origin,
		"airline.destination": dest,
		"p1.id":               origin,
		"p2.id":               dest,
		"airline.affiliate":   goqu.I("ac.id"),
		"airline.model":       goqu.I("am.id"),
	}).Limit(10).Offset(uint(page-1) * 10).Distinct().ToSQL()

	rows, err := db.Query(sql)
	return rows, err
}
