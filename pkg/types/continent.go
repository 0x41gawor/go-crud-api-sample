package types

import "fmt"

type Continent struct {
	Id           int
	Name         string
	Population   float32
	Gdp          float32
	GdpPerCapita float32
}

func NewContinent(name string, population, gdp, gdpPerCapita float32) *Continent {
	return &Continent{
		Id:           0,
		Name:         name,
		Population:   population,
		Gdp:          gdp,
		GdpPerCapita: gdpPerCapita,
	}
}

func (c *Continent) String() string {
	// "mordo, to jest backed" ladna prezentacje danych to sobie zrobisz na froncie, tu chodzi tylko o to, zeby w logach byloa widac ze jakies dane sa
	return fmt.Sprintf("Continent{%d, %s, %.2f, %.2f, %.2f}", c.Id, c.Name, c.Population, c.Gdp, c.GdpPerCapita)
}
