NAME := awesome-cli
VERSION := 0.1.0
REVISION := $(shell git describe --always)
LDFLAGS := -ldflags="-s -w -X \"main.Name=$(NAME)\" -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""

.PHONY: clean
clean:
	rm -rf bin/*
	rm -rf vendor/*

.PHONY: deps
deps:
	go get -d -t -v .

.PHONY: build
build: deps
	go build $(LDFLAGS) -o bin/$(NAME)

.PHONY: test-all
test-all: test test-race vet lint

.PHONY: test
test:
	go test -v -covermode=count -coverprofile=cover.out -timeout=30s -parallel=4 ./...

.PHONY: test-race
test-race:
	go test -race .

.PHONY: vet
vet:
	go vet *.go

.PHONY: lint
lint:
	@go get github.com/golang/lint/golint
	golint ./...
