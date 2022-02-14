.PHONY: test
test:

	go test -race ./...

.PHONY: build
build:

	CGO_ENABLED=0 go build -ldflags "-w -s" -o bin/invader-sim cmd/main.go

all: test build