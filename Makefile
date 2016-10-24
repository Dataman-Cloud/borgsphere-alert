PACKAGES = $(shell go list ./src/...)

.PHONY: build fmt lint run test vet

## OS checking
OS := $(shell uname)
ifeq ($(OS),Darwin)
	BUILD_OPTS=
else
	BUILD_OPTS=CGO_ENABLED=0 GOOS=linux GOARCH=amd64
endif

# Prepend our vendor directory to the system GOPATH
# so that import path resolution will prioritize
# our third party snapshots.
export GO15VENDOREXPERIMENT=1
#GOPATH := ${PWD}/vendor:${GOPATH}
# export GOPATH

# Used to populate version variable in main package.
VERSION=$(shell git describe --always --tags)
BUILD_TIME=$(shell date -u +%Y-%m-%d:%H-%M-%S)
GO_LDFLAGS=-ldflags "-X `go list ./src/version`.Version=$(VERSION) -X `go list ./src/version`.BuildTime=$(BUILD_TIME)"

default: build

build: fmt
	@echo "🐳 $@"
	 ${BUILD_OPTS} go build ${GO_LDFLAGS} -v -o ./bin/alert ./src/

rel: fmt
	@echo "🐳 $@"
	${BUILD_OPTS} go build -v -o ../rel/alert ./src/

# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	@echo "🐳 $@"
	go fmt ./src/...

# https://github.com/golang/lint
# go get github.com/golang/lint/golint
lint:
	@echo "🐳  $@"
	@test -z "$$(golint ./src/... | tee /dev/stderr)"

run: build
	@echo "🐳 $@"
	./bin/alert

test:
	@echo "🐳 $@"
	go test -cover=true ./src/...

# http://godoc.org/code.google.com/p/go.tools/cmd/vet
# go get code.google.com/p/go.tools/cmd/vet
vet:
#	go vet ./src/...
#

clean:
	@echo "🐳 $@"
	rm -rf ../bin/* ../rel/*
