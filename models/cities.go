package models

import "FLightening/sqlconn"

type City struct {
	Name     string `json:"name"`
	Province string `json:"province"`
	Id       int    `json:"id"`
	Code     string `json:"code"`
}

func GetAllCities() []City {
	rows, err := sqlconn.GetAllCities()
	defer rows.Close()

	if err != nil {
		return nil
	}

	ret := make([]City, 0)

	for rows.Next() {
		city := City{}
		rows.Scan(&city.Id, &city.Name, &city.Province, &city.Code)
		ret = append(ret, city)
	}

	return ret
}
