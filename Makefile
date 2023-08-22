
build:
	@cd cmd; \
	go build -o bin/go-crud-api-sample
run: build
	@./bin/go-crud-api-sample

test:
	@go test -v ./...
