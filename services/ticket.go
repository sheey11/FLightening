package services

import (
	"FLightening/sqlconn"
	"errors"
)

func BookTicket(cabin, user, shift int, pass []sqlconn.Passenger) (int, error) {
	if len(pass) == 0 {
		return -1, errors.New("请至少添加一个乘客")
	}

	return sqlconn.BookTicket(cabin, shift, pass, user)
}
