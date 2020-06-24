gribbled: gribbled.go
	CGO_ENABLED=0 go build -ldflags '-s -w'

test:
	go test

.PHONY: test
