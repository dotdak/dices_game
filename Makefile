.PHONY: lint build
lint:
	golangci-lint run --timeout 10m -v ./...
build:
	go build main.go
run:
	go build main.go && ./main
