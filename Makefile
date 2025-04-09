.PHONY: build, install

build:
	go build -ldflags "-s -w" -o dist/debrix

install:
	go build -o "$$GOPATH/bin/debrix"