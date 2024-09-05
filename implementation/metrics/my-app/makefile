BINARY_NAME=metrics
OS_NAME := $(shell uname -s)

.PHONY: build run prod-run clean hello

build:
	GOARCH=amd64 GOOS=darwin go build -o ./bin/${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ./bin/${BINARY_NAME}-windows main.go

run:
	go run *.go

prod: build
ifeq ($(OS_NAME),Linux)
	./bin/${BINARY_NAME}-linux
endif
ifeq ($(OS_NAME),Darwin)
	./bin/${BINARY_NAME}-darwin
endif
ifeq ($(OS_NAME),MINGW64_NT)
	./bin/${BINARY_NAME}-windows
endif

clean:
	go clean
	rm -f ${BINARY_NAME}-darwin
	rm -f ${BINARY_NAME}-linux
	rm -f ${BINARY_NAME}-windows

os:
	@echo "Operating System: $(OS_NAME)"
