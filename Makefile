all: gen swag test build

build:
	go vet ./...
	go fmt ./...
	go build -o bin/server ./cmd/server/main.go

test:
	go test -v ./...

run: gen swag test build
	./bin/server

gen:
	oapi-codegen -config configs/oapi-codegen-gin.yaml api/spec.yaml

swag:
	swag init -g cmd/server/main.go -o ./internal/docs
	@echo "Swagger docs generated at: http://localhost:8080/swagger/index.html"