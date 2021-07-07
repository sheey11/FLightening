package sqlconn

import (
	"errors"
	"time"

	"github.com/doug-martin/goqu/v9"
)

type OrderDO struct {
	Id     int
	Shift  int
	User   int
	Price  float32
	Status int
	Time   time.Time
}

func FindOrderById(id int) OrderDO {
	o := OrderDO{}

	_sql, _, _ := dialect.Select("*").From("orders").Where(goqu.Ex{
		"id": id,
	}).ToSQL()

	r := db.QueryRow(_sql)
	r.Scan(&o.Id, &o.Shift, &o.User, &o.Price, &o.Status, &o.Time)
	return o
}

func FetchOrders(uid int, page uint) ([]OrderDO, error) {
	if page == 0 {
		page = 1
	}

	_sql, _, _ := dialect.Select("*").From("orders").Where(goqu.Ex{
		"user": uid,
	}).Limit(10).Offset(uint(page-1) * 10).ToSQL()

	rows, err := db.Query(_sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ret := make([]OrderDO, 0)
	for rows.Next() {
		o := OrderDO{}
		rows.Scan(&o.Id, &o.Shift, &o.User, &o.Price, &o.Status, &o.Time)
		ret = append(ret, o)
	}
	return ret, nil
}

func FetchAllOrders(page uint) ([]OrderDO, error) {
	if page == 0 {
		page = 1
	}

	_sql, _, _ := dialect.Select("*").From("orders").Limit(10).Offset(uint(page-1) * 10).ToSQL()

	rows, err := db.Query(_sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ret := make([]OrderDO, 0)
	for rows.Next() {
		o := OrderDO{}
		rows.Scan(&o.Id, &o.Shift, &o.User, &o.Price, &o.Status, &o.Time)
		ret = append(ret, o)
	}
	return ret, nil
}

func MarkStatus(oid, uid, status int) error {
	_sql, _, _ := dialect.From("orders").Where(goqu.Ex{"id": oid, "user": uid}).Update().Set(
		goqu.Record{"status": status},
	).ToSQL()
	r, e := db.Exec(_sql)
	if ra, _ := r.RowsAffected(); ra != 1 {
		return errors.New("订单不存在")
	}
	return e
}

func RestoreReleventSeatVacancy(oid, uid int) bool {
	_sql, _, _ := dialect.From("orders").Where(goqu.Ex{"id": oid, "user": uid}).Select("shift").ToSQL()
	shift := 0

	tx, _ := db.Begin()
	r := tx.QueryRow(_sql)
	r.Scan(&shift)

	_sql, _, _ = dialect.From("seats").Where(goqu.Ex{"affiliate_order": oid}).Select(goqu.COUNT("id")).ToSQL()
	seats := 0
	r = tx.QueryRow(_sql)
	r.Scan(&seats)

	_sql, _, _ = dialect.From("shifts").Where(goqu.Ex{"id": shift}).Select("remaining_seat").ToSQL()
	r = tx.QueryRow(_sql)
	remainingSeat := 0
	r.Scan(&remainingSeat)

	_sql, _, _ = dialect.From("shifts").Where(goqu.Ex{"id": shift}).Update().Set(
		goqu.Record{"remaining_seat": remainingSeat + seats},
	).ToSQL()
	_, err := tx.Exec(_sql)

	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return err == nil
}
