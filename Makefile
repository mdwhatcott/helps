#!/usr/bin/make -f

test:
	go mod tidy
	go fmt ./...
	go build ./...
	go test -v -short ./...

install:
	go install ./cmd/...
