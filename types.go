package main

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
