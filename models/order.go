package models

import (
	"FLightening/sqlconn"
	"fmt"
	"time"
)

type Order struct {
	id     int
	shift  int
	user   int
	Price  float32     `json:"price"`
	Status OrderStatus `json:"status"`
	Time   time.Time   `json:"time"`
}

func (o *Order) GetUniqueID() string {
	return fmt.Sprintf(
		"%04d%02d%02d%02d%02d%02d%04d%05d%05d",
		o.Time.Year(),
		o.Time.Month(),
		o.Time.Day(),
		o.Time.Hour(),
		o.Time.Minute(),
		o.Time.Second(),
		o.Time.Nanosecond()/1000,
		o.id,
		o.user,
	)
}

func FindOrderById(id int) Order {
	o := sqlconn.FindOrderById(id)
	return Order{
		id:     o.Id,
		shift:  o.Shift,
		user:   o.User,
		Price:  o.Price,
		Status: o.Status,
		Time:   o.Time,
	}
}
