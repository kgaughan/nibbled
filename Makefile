VERSION:=0.1.0

SOURCE:=$(wildcard *.go)

build: go.mod nibbled

tidy: go.mod

nibbled: $(SOURCE)
	CGO_ENABLED=0 go build -trimpath -ldflags '-s -w -X main.Version=$(VERSION)'

go.mod: $(SOURCE)
	go mod tidy

test:
	go test

.DEFAULT: build

.PHONY: build test tidy
