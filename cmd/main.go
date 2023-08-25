package main

import (
	"fmt"

	"github.com/0x41gawor/go-crud-api-sample/pkg/model"
	"github.com/0x41gawor/go-crud-api-sample/pkg/repo"
)

func main() {
	// server := api.NewServer(":3000")
	// server.Run()
	dbConnectionHelper, err := repo.NewDatabaseConnectionHelper()
	if err != nil {
		fmt.Println(err.Error())
	}

	countryRepo := repo.NewCountryRepository(dbConnectionHelper.DB)

	fmt.Println(countryRepo.List())

	polska := model.NewCountry("Polska", 1, 38.4, 4, 4, []string{"Wawa", "Burdno", "Posen", "alko", "plock"})
	_, err = countryRepo.Create(polska)
	if err != nil {
		fmt.Println(err.Error())
	}

}
