package models

import (
	"FLightening/sqlconn"
	"database/sql"
	"time"
)

type Shift struct {
	Id                 int          `json:"id"`
	TakeOff            time.Time    `json:"takeoff"`
	Landing            time.Time    `json:"landing"`
	ActualTakeOff_Null sql.NullTime `json:"-"`
	ActualLanding_Null sql.NullTime `json:"-"`
	ActualTakeOff      *time.Time   `json:"actual_takeoff"`
	ActualLanding      *time.Time   `json:"actual_landing"`
	Status             ShiftStatus  `json:"status"`
	EcoPrice           float32      `json:"economic_price"`
	PrePrice           float32      `json:"premium_price"`
	BusPrice           float32      `json:"business_price"`
	FirPrice           float32      `json:"first_price"`
	Vacancy            int          `json:"vacancy"`
	Airline            int          `json:"affiliate_airline"`
	AirlineID          string       `json:"airline_id"`
}

func FindNearestNShifts(n uint, airline int) []Shift {
	rows, err := sqlconn.FindNearestNShifts(n, airline)

	if err != nil {
		return make([]Shift, 0)
	} else {
		defer rows.Close()
	}

	shifts := make([]Shift, 0)

	for rows.Next() {
		shift := Shift{}
		rows.Scan(
			&shift.Id,
			&shift.TakeOff,
			&shift.Landing,
			&shift.ActualTakeOff,
			&shift.ActualLanding,
			&shift.Status,
			&shift.EcoPrice,
			&shift.PrePrice,
			&shift.BusPrice,
			&shift.FirPrice,
			&shift.Vacancy,
			&shift.Airline,
		)
		if shift.ActualLanding_Null.Valid {
			shift.ActualLanding = &shift.ActualLanding_Null.Time
		}
		if shift.ActualTakeOff_Null.Valid {
			shift.ActualTakeOff = &shift.ActualTakeOff_Null.Time
		}
		shifts = append(shifts, shift)
	}

	return shifts
}

func FindShiftById(id int) Shift {
	row := sqlconn.FindShiftById(id)

	shift := Shift{}
	row.Scan(
		&shift.Id,
		&shift.TakeOff,
		&shift.Landing,
		&shift.ActualTakeOff,
		&shift.ActualLanding,
		&shift.Status,
		&shift.EcoPrice,
		&shift.PrePrice,
		&shift.BusPrice,
		&shift.FirPrice,
		&shift.Vacancy,
		&shift.Airline,
	)
	if shift.ActualLanding_Null.Valid {
		shift.ActualLanding = &shift.ActualLanding_Null.Time
	}
	if shift.ActualTakeOff_Null.Valid {
		shift.ActualTakeOff = &shift.ActualTakeOff_Null.Time
	}

	return shift
}

func FetchAllShifts(page uint) []Shift {
	if page < 1 {
		page = 1
	}
	rows, err := sqlconn.FetchAllShifts(page)

	if err != nil {
		return make([]Shift, 0)
	} else {
		defer rows.Close()
	}

	shifts := make([]Shift, 0)

	for rows.Next() {
		shift := Shift{}
		rows.Scan(
			&shift.Id,
			&shift.TakeOff,
			&shift.Landing,
			&shift.ActualTakeOff,
			&shift.ActualLanding,
			&shift.Status,
			&shift.EcoPrice,
			&shift.PrePrice,
			&shift.BusPrice,
			&shift.FirPrice,
			&shift.Vacancy,
			&shift.Airline,
			&shift.AirlineID,
		)
		if shift.ActualLanding_Null.Valid {
			shift.ActualLanding = &shift.ActualLanding_Null.Time
		}
		if shift.ActualTakeOff_Null.Valid {
			shift.ActualTakeOff = &shift.ActualTakeOff_Null.Time
		}
		shifts = append(shifts, shift)
	}

	return shifts
}
