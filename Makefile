NAME := readme-txt
.DEFAULT_GOAL := build

VERS ?= $(shell git describe --dirty --long --always)

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

check: vet lint
.PHONY:check

build:
	go build -ldflags "-X main.commitVersion=$(VERS)" -o build/$(NAME)
.PHONY: build

build-linux-x64:
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.commitVersion=$(VERS)" -o build/$(NAME)-linux-amd64
.PHONY: build-linux-x64

build-all:
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.commitVersion=$(VERS)" -o build/$(NAME)-linux-amd64
	GOOS=linux GOARCH=arm64 go build -ldflags "-X main.commitVersion=$(VERS)" -o build/$(NAME)-linux-arm64
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.commitVersion=$(VERS)" -o build/$(NAME)-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build -ldflags "-X main.commitVersion=$(VERS)" -o build/$(NAME)-darwin-arm64
	GOOS=windows GOARCH=amd64 go build -ldflags "-X main.commitVersion=$(VERS)" -o build/$(NAME)-windows-amd64
.PHONY: build-all

clean:
	rm -f build/*
.PHONY:clean

tidy:
	go mod tidy
.PHONY:tidy
