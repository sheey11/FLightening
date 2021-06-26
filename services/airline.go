package services

import "FLightening/models"

func FindAirline(origin, dest, page int) []models.Airline {
	return models.FindAirlineByOriginAndDest(origin, dest, page)
}
