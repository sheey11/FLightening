package sqlconn

import (
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
