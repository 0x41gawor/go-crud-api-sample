package main

import "fmt"

func main() {
	continentRepository, err := NewContinentRepository()

	if err != nil {
		panic(err)
	}
	item, err := continentRepository.List()
	fmt.Println(item)
	fmt.Println(err)

	err = continentRepository.Delete(5)
	fmt.Println(err)

}
