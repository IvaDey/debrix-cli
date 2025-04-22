.PHONY: build, install

build:
	go build -ldflags "-s -w -X cmd.version=$(git describe --tags)" -o dist/debrix

install:
	go build -ldflags "-X github.com/ivadey/debrix-cli/cmd.version=$(git describe --tags)-dev" -o "$$GOPATH/bin/debrix"