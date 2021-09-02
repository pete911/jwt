VERSION ?= dev

.DEFAULT_GOAL := build

test:
	go test ./...
.PHONY:test

build: test
	go fmt ./...
	go vet ./...
	go mod tidy
	go mod vendor
	go build -ldflags "-X 'github.com/pete911/jwt/cmd.Version=${VERSION}'"
.PHONY:build

install: test
	go install -ldflags "-X 'github.com/pete911/jwt/cmd.Version=${VERSION}'"
.PHONY:install