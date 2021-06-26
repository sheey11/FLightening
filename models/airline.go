package models

import "FLightening/sqlconn"

type Airline struct {
	Name        string
	Affiliate   string
	CompanyLogo string
	Model       string
	Origin      string
	Destination string
}

func FindAirlineByOriginAndDest(origin, dest, page int) []Airline {
	rows, err := sqlconn.FindAirlineByOriginAndDest(origin, dest, page)
	if err != nil {
		return nil
	}

	ret := make([]Airline, 1)

	for rows.Next() {
		airline := Airline{}
		rows.Scan(&airline.Name, &airline.Model, &airline.Origin, &airline.Destination, &airline.Affiliate, &airline.CompanyLogo)
		ret = append(ret, airline)
	}

	return ret
}
