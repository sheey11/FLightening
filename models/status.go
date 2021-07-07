package models

type SeatStatus = int

const (
	SeatVacancy = iota
	SeatOcuppied
	SeatBrought
)

type ShiftStatus = int

const (
	ShiftScheduled = iota
	ShiftAirborne
	ShiftLanded
	ShiftUnknown
)

type OrderStatus = int

const (
	OrderToBePaied = iota
	OrderPaid
	OrderCanceled = -1
)
