package model

import "fmt"

type Country struct {
	Id           int
	Name         string
	ContinentId  int
	Population   float32
	Gdp          float32
	GdpPerCapita float32
	Top5Cities   []string
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

func (c *Country) String() string {
	return fmt.Sprintf("Country{%d, %s, %d, %.2f,%.2f,%.2f,%s}", c.Id, c.Name, c.ContinentId, c.Population, c.Gdp, c.GdpPerCapita, c.Top5Cities)
}
