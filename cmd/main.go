package main

import (
	"fmt"

	"github.com/0x41gawor/go-crud-api-sample/pkg/model"
	"github.com/0x41gawor/go-crud-api-sample/pkg/repo"
)

func main() {
	// server := api.NewServer(":3000")
	// server.Run()
	user := model.NewUser(0, "admin", "admin")
	fmt.Println(user)
	db, _ := repo.NewDatabaseConnectionHelper()
	repo := repo.NewUserRepository(db.DB)
	id, _ := repo.Create(user)
	println(id)
	user, err := repo.FindByLogin("admin")
	fmt.Println(user, err)
}
