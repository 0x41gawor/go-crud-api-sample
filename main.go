package main

import "fmt"

func main() {
	continentRepository, err := NewContinentRepository()

	if err != nil {
		panic(err)
	}

	europe := NewContinent("dgdf", 23, 23, 23)

	id, err := continentRepository.Create(europe)

	fmt.Printf("Inserted id: %d \n", id)

	if err != nil {
		panic(err)
	}
	item, err := continentRepository.List()
	fmt.Println(item)
}
