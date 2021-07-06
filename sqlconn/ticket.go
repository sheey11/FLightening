package sqlconn

import (
	"errors"
	"regexp"

	"github.com/doug-martin/goqu/v9"
)

type Passenger struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func validatePassengers(pass []Passenger) bool {
	for _, p := range pass {
		namePass := len([]rune(p.Name)) < 5 && len([]rune(p.Name)) > 2
		idPass, _ := regexp.MatchString("\\d{17}[X\\d]", p.ID)

		if !namePass || !idPass {
			return false
		}
	}
	return true
}

func CalcPrice(cabin, shift, nPassenger int) int {
	return LookUpPrice(cabin, shift) * nPassenger
}

func BookTicket(cabin, shift int, pass []Passenger, user int) (int, error) {
	// validate
	if !validatePassengers(pass) {
		return 0, errors.New("请检查乘客信息")
	}

	// find shift
	if !ShiftExists(shift) {
		return 0, errors.New("航班不存在")
	}

	// prepare order
	price := CalcPrice(cabin, shift, len(pass))
	orderSql, _, _ := dialect.From("orders").Insert().Rows(
		goqu.Record{"shift": shift, "user": user, "price": price, "time": "NOW()"},
	).ToSQL()

	tx, _ := db.Begin()
	_, e := tx.Exec(orderSql)
	if e != nil {
		tx.Rollback()
		return 0, e
	}

	r := tx.QueryRow("select last_insert_id();")
	oid := 0
	r.Scan(&oid)

	for _, p := range pass {
		_sql, _, _ := dialect.From("seats").Insert().Rows(
			goqu.Record{"shift": shift, "affiliate_order": oid, "passenger_name": p.Name, "passenger_id": p.ID},
		).ToSQL()
		_, e = tx.Exec(_sql)
		if e != nil {
			tx.Rollback()
			return 0, e
		}
	}
	tx.Commit()
	return oid, nil
}
