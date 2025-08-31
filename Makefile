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
	cp api/spec.yaml docs/openapi.yaml
	cp api/spec.yaml swagger-ui/openapi.yaml
	@echo "Swagger docs generated at: http://localhost:8080/docs/index.html"