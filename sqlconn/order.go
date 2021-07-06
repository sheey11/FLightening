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
