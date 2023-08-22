package main

import "fmt"

type Continent struct {
	id           int
	name         string
	population   float32
	gdp          float32
	gdpPerCapita float32
}

type Country struct {
	id           int
	name         string
	continentId  int
	population   float32
	gdp          float32
	gdpPerCapita float32
	top5Cities   []string
}

func NewContinent(name string, population, gdp, gdpPerCapita float32) *Continent {
	return &Continent{
		id:           0,
		name:         name,
		population:   population,
		gdp:          gdp,
		gdpPerCapita: gdpPerCapita,
	}
}

func (c *Continent) String() string {
	// "mordo, to jest backed" ladna prezentacje danych to sobie zrobisz na froncie, tu chodzi tylko o to, zeby w logach byloa widac ze jakies dane sa
	return fmt.Sprintf("Continent{%d, %s, %.2f, %.2f, %.2f}", c.id, c.name, c.population, c.gdp, c.gdpPerCapita)
}

func NewCountry(name string, continentId int, population, gdp, gdpPerCapita float32, top5Cities []string) *Country {
	return &Country{
		id:           0,
		name:         name,
		continentId:  continentId,
		population:   population,
		gdp:          gdp,
		gdpPerCapita: gdpPerCapita,
		top5Cities:   top5Cities,
	}
}

func (c *Country) String() string {
	return fmt.Sprintf("Country{%d, %s, %d, %.2f,%.2f,%.2f,%s}", c.id, c.name, c.continentId, c.population, c.gdp, c.gdpPerCapita, c.top5Cities)
}
