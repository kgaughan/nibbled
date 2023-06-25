SOURCE:=$(wildcard *.go)

build: go.mod nibbled

tidy: go.mod

nibbled: $(SOURCE)
	CGO_ENABLED=0 go build -tags netgo -trimpath -ldflags '-s -w'

update:
	go get -u ./...
	go mod tidy

go.mod: $(SOURCE)
	go mod tidy

test:
	go test

.DEFAULT: build

.PHONY: build test tidy update
