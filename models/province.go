package models

import "FLightening/sqlconn"

type Province struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func FetchAllProvince() []Province {
	rows, err := sqlconn.FetchAllProvince()
	if err != nil {
		return make([]Province, 0)
	}

	ps := make([]Province, 0)

	for rows.Next() {
		p := Province{}
		rows.Scan(&p.Id, &p.Name)
		ps = append(ps, p)
	}
	return ps
}
