.PHONY: test run build
test:
	go test -race ./...
run:
	go run .
build:
	go build -o bin/api .
