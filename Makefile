.PHONY: test
test:

	go test -race ./...

.PHONY: build
build:

all: test build