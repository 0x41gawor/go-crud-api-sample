package main

import "fmt"

func main() {
	continentRepository, err := NewContinentRepository()

	if err != nil {
		panic(err)
	}
	item, err := continentRepository.Read(5)
	fmt.Println(item)
	fmt.Println(err)
}
