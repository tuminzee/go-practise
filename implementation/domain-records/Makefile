BINARY_NAME=domain-records
OS_NAME := $(shell uname -s)
DOMAIN ?= default.com
USERNAME ?=

.PHONY: build run prod-run clean hello

build:
	GOARCH=amd64 GOOS=darwin go build -o ./bin/${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ./bin/${BINARY_NAME}-windows main.go

run:
	go run *.go $(DOMAIN)

prod: build
ifeq ($(OS_NAME),Linux)
	./bin/${BINARY_NAME}-linux $(DOMAIN)
endif
ifeq ($(OS_NAME),Darwin)
	./bin/${BINARY_NAME}-darwin $(DOMAIN)
endif
ifeq ($(OS_NAME),MINGW64_NT)
	./bin/${BINARY_NAME}-windows $(DOMAIN)
endif

clean:
	go clean
	rm -f ${BINARY_NAME}-darwin
	rm -f ${BINARY_NAME}-linux
	rm -f ${BINARY_NAME}-windows

os:
	@echo "Operating System: $(OS_NAME)"
