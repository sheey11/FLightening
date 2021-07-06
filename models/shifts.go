package models

import (
	"FLightening/sqlconn"
	"time"
)

type Shift struct {
	TakeOff       time.Time   `json:"takeoff"`
	Landing       time.Time   `json:"landing"`
	ActualTakeOff time.Time   `json:"actual_takeoff"`
	ActualLanding time.Time   `json:"actual_landing"`
	Status        ShiftStatus `json:"status"`
	EcoPrice      float32     `json:"economic_price"`
	PrePrice      float32     `json:"premium_price"`
	BusPrice      float32     `json:"business_price"`
	FirPrice      float32     `json:"first_price"`
}

func FindNearestNShifts(n uint, airline int) []Shift {
	rows, err := sqlconn.FindNearestNShifts(n, airline)
	defer rows.Close()

	if err != nil {
		return nil
	}

	shifts := make([]Shift, 0)

	for rows.Next() {
		shift := Shift{}
		rows.Scan(
			&shift.TakeOff,
			&shift.Landing,
			&shift.ActualTakeOff,
			&shift.ActualLanding,
			&shift.Status,
			&shift.EcoPrice,
			&shift.PrePrice,
			&shift.BusPrice,
			&shift.FirPrice,
		)
		shifts = append(shifts, shift)
	}

	return shifts
}
