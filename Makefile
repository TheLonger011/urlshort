.PHONY: build run clean fmt test

BINARY=url-shortener

build:
	go build -o $(BINARY) ./cmd/url-shortener/main.go

run:
	go run ./cmd/url-shortener/main.go

clean:
	rm -f $(BINARY)

fmt:
	go fmt ./...

test:
	go test -v ./...