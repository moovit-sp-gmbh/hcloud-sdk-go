.PHONY: build

test:
	go test ./...

run:
	go run -race main.go
