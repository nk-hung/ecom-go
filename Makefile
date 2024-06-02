build:
	@go build -o bin/ecom-go cmd/server/main.go

test:
	@go test -v ./...
	
run: build
	@./bin/ecom-go
