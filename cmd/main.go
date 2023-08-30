package main

import "github.com/0x41gawor/go-crud-api-sample/pkg/api"

func main() {
	server := api.NewServer(":3000")
	server.Run()
}
