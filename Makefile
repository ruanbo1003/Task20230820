
.PHONY: build clean tool lint help

all: build

build:
	go build -o ./bin/ ./cmd/profile
	go build -o ./bin/ ./cmd/consultation

lint:
	golint ./...

clean:
	rm -rf ./bin/*
	go clean -i .
