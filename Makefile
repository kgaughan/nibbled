REPO=github.com/kgaughan/gribbled

build:
	go build -o ${GOPATH}/bin/gribbled $(REPO)

test:
	go test $(REPO)

.PHONY: build test
