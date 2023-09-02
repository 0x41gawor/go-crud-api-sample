package model

import "fmt"

type Country struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	ContinentId  int      `json:"continent_id"`
	Population   float32  `json:"population"`
	Gdp          float32  `json:"gdp"`
	GdpPerCapita float32  `json:"gdp_per_capita"`
	Top5Cities   []string `json:"top_5_cities"`
}

func NewCountry(name string, continentId int, population, gdp, gdpPerCapita float32, top5Cities []string) *Country {
	return &Country{
		Id:           0,
		Name:         name,
		ContinentId:  continentId,
		Population:   population,
		Gdp:          gdp,
		GdpPerCapita: gdpPerCapita,
		Top5Cities:   top5Cities,
	}
}

func (m *Country) String() string {
	return fmt.Sprintf("Country{%d, %s, %d, %.2f,%.2f,%.2f,%s}", m.Id, m.Name, m.ContinentId, m.Population, m.Gdp, m.GdpPerCapita, m.Top5Cities)
}
