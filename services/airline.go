package services

import "FLightening/models"

func FindAirline(origin, dest, page int) []models.Airline {
	return models.FindAirlineByOriginAndDest(origin, dest, page)
}

func FindNearestNShifts(n uint, airline int) []models.Shift {
	return models.FindNearestNShifts(n, airline)
}

func GetAllCities() []models.City {
	return models.GetAllCities()
}
