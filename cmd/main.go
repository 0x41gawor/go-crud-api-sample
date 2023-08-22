package main

import "fmt"

import "github.com/0x41gawor/go-crud-api-sample/pkg/repo"
import "github.com/0x41gawor/go-crud-api-sample/pkg/types"

func main() {

	elo, err := repo.NewContinentRepository()
	if err != nil {
		err.Error()
	}

	oceania := types.NewContinent("Chujoza", 12, 12, 12)

	elo.Create(oceania)

	items, err := elo.List()
	fmt.Println(items)
}
