nibbled: main.go
	CGO_ENABLED=0 go build -trimpath -ldflags '-s -w'

test:
	go test

.PHONY: test
