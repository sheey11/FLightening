package routers

import "FLightening/sqlconn"

type LoginDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type AirlineSearchDTO struct {
	Origin      int `json:"origin" binding:"required"`
	Destination int `json:"dest" binding:"required"`
	Page        int `json:"page"`
}

type BookingDTO struct {
	Shift     int                 `json:"shift" binding:"required"`
	Cabin     *int                `json:"cabin" binding:"required"`
	Passenger []sqlconn.Passenger `json:"passenger"`
}

type CityDTO struct {
	Name     string `json:"name" binding:"required"`
	Province *int   `json:"province" binding:"required"`
	Code     string `json:"code" binding:"required"`
}

type CityUpdateDTO struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}

type ProvinceDTO struct {
	Name string `json:"name" binding:"required"`
}
type ProvinceUpdateDTO struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
