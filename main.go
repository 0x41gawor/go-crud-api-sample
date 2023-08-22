package main

import "fmt"

func main() {
	continentRepository, err := NewContinentRepository()

	america := NewContinent("Oceania", 324, 342, 4324)

	if err != nil {
		panic(err)
	}
	item, err := continentRepository.Read(4)
	fmt.Println(item)
	fmt.Println(err)
	err = continentRepository.Update(4, america)
	fmt.Println(err)
	item, err = continentRepository.Read(4)
	fmt.Println(item)
	fmt.Println(err)
}
