#!/usr/bin/make -f

test:
	go build ./...
	go test -v -short ./...

install:
	go install ./cmd/...