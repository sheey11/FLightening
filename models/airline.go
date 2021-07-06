package models

import "FLightening/sqlconn"

type Airline struct {
	id          int
	Name        string  `json:"id"`
	Affiliate   string  `json:"affiliate"`
	CompanyLogo string  `json:"logo"`
	Model       string  `json:"model"`
	Origin      string  `json:"origin"`
	Destination string  `json:"dest"`
	Shifts      []Shift `json:"shifts"`
}

func (a Airline) GetId() int {
	return a.id
}

func FindAirlineByOriginAndDest(origin, dest, page int) []Airline {
	rows, err := sqlconn.FindAirlineByOriginAndDest(origin, dest, page)
	defer rows.Close()

	if err != nil {
		return nil
	}

	ret := make([]Airline, 0)

	for rows.Next() {
		airline := Airline{}
		rows.Scan(&airline.id, &airline.Name, &airline.Model, &airline.Origin, &airline.Destination, &airline.Affiliate, &airline.CompanyLogo)
		ret = append(ret, airline)
	}

	return ret
}
