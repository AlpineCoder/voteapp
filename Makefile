all: build

build:
	go vet ./...
	go fmt ./...
	go build -o bin/server ./cmd/main.go

test:
	go test -v ./...

run: gen test build
	./bin/server

gen:
	oapi-codegen -config configs/oapi-codegen-gin.yaml api/spec.yaml