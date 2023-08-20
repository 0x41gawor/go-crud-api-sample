package main

import "fmt"

func main() {
	elo, err := NewContinentRepository()
	if err != nil {
		panic(err)
	}
	item, err := elo.List()
	fmt.Println(item[0])
}
