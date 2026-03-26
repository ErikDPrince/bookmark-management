.PHONY: run swagger dev-run test

run:
	go run cmd/api/main.go

swagger:
	swag init -g cmd/api/main.go --output docs

dev-run: swagger run

test:
	go test ./...