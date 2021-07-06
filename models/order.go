package models

import (
	"fmt"
	"time"
)

type Order struct {
	id     int
	Shift  Shift
	User   User
	Price  float32
	Status OrderStatus
	Time   time.Time
}

func (o *Order) GetUniqueID() string {
	return fmt.Sprintf(
		"%04d%02d%02d%02d%02d%02d%04d%05d%5d",
		o.Time.Year(),
		o.Time.Month(),
		o.Time.Day(),
		o.Time.Hour(),
		o.Time.Minute(),
		o.Time.Second(),
		o.Time.Nanosecond()/1000,
		o.id,
		o.User.Id,
	)
}
