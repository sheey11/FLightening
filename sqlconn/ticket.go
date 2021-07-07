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
	regex := regexp.MustCompile(`\d{17}[X\d]`)

	for _, p := range pass {
		namePass := len([]rune(p.Name)) <= 5 && len([]rune(p.Name)) >= 2
		idPass := regex.MatchString(p.ID)

		if !namePass || !idPass {
			return false
		}
	}
	return true
}

func CalcPrice(cabin, shift, nPassenger int) float32 {
	return LookUpPrice(cabin, shift) * float32(nPassenger)
}

func BookTicket(cabin, shift int, pass []Passenger, user int) (int, error) {
	// validate
	if !validatePassengers(pass) {
		return 0, errors.New("请检查乘客信息")
	}

	// find shift
	shiftStatus := ShiftExistsAndVacancy(shift)
	if shiftStatus == -1 {
		return 0, errors.New("航班不存在")
	}
	if shiftStatus == -2 {
		return 0, errors.New("航班没有空位了")
	}

	// prepare order
	price := CalcPrice(cabin, shift, len(pass))
	orderSql, _, _ := dialect.From("orders").Insert().Rows(
		goqu.Record{"shift": shift, "user": user, "price": price, "time": goqu.Func("NOW")},
	).ToSQL()

	// begin tranaction
	tx, _ := db.Begin()
	// insert order
	_, e := tx.Exec(orderSql)
	if e != nil {
		tx.Rollback()
		return 0, e
	}

	// get order id
	r := tx.QueryRow("select last_insert_id();")
	oid := 0
	r.Scan(&oid)

	// insert seats for passengers
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

	// query for remaining_seat
	remainingSql, _, _ := dialect.From("shifts").Where(goqu.Ex{"id": shift}).Select("remaining_seat").ToSQL()
	r = tx.QueryRow(remainingSql)
	remainingSeat := 0
	r.Scan(&remainingSeat)

	// write back new remaining_seat
	shiftSql, _, _ := dialect.From("shifts").Update().Set(
		goqu.Record{"remaining_seat": remainingSeat - 1},
	).ToSQL()
	_, err := tx.Exec(shiftSql)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()
	return oid, nil
}

func FindPassenger(order int) ([]Passenger, error) {
	_sql, _, _ := dialect.From("seats").Where(
		goqu.Ex{"affiliate_order": order},
	).Select("passenger_name", "passenger_id").ToSQL()

	rows, err := db.Query(_sql)
	if err != nil {
		return make([]Passenger, 0), err
	}

	pass := make([]Passenger, 0)
	for rows.Next() {
		p := Passenger{}
		rows.Scan(&p.Name, &p.ID)
		p.ID = p.ID[:6]
	}
	return pass, nil
}
