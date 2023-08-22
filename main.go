package main

import "fmt"

func main() {
	elo, err := NewContinentRepository()
	if err != nil {
		panic(err)
	}
	item, err := elo.List()
	fmt.Println(item[0])
	polska := NewCountry("Poland", 1, 23, 23.3, 23, []string{"Warsaw", "Cracow", "fSf", "fsd", "fs"})
	fmt.Println(polska)
}
