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


