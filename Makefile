
all: gen fmt test install

deps:
	go install golang.org/x/tools/cmd/goimports@latest
	go install golang.org/x/lint/golint@latest

gen:
	go generate ./...

fmt:
	find . -name '*.go' -exec goimports -l -w {} \;

clean:
	go clean ./...

install:
	go install ./...

test:
	go test ./...

lint:
	golint ./...

.PHONY: all deps gen fmt clean install test lint

