# Setup
First step is to make a Makefile:
```makefile
build:
	@go build -o bin/go-crud-api-sample
run: build
	@./bin/go-crud-api-sample

test:
	@go test -v ./...
```
then main.go
```golang
package main

import "fmt"

func main() {
	fmt.Println("Hello")
}
```
then run a command:
```bash
go mod init github.com/0x41gawor/go-crud-api-sample
```
And at the end you can go with:
```bash
make run
```

# Database setup
Na VMce uruchomiłem:
```bash
sudo docker run --name mysql-ejek -p 3306:3306 -e MYSQL_ROOT_PASSWORD=ejek -d mysql:latest
```
> dodalem na koniec `--mount source=my-sql-vol`
> uprzednio tworzac volume wg. tego https://docs.docker.com/storage/volumes/ 
A potem na Win10 użyłem MySQL workbench</br>
i tam dałem paramki na new connection
```
root@192.168.56.106:3306 pwd: ejek
```

no i krok po kroczku leciałem z `db_init.sql`


