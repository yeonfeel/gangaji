BINARY = gangaji
VERSION=$(shell git describe --tags)

GOCMD = go
GOFLAGS ?= $(GOFLAGS:)
LDFLAGS =

default: build

mod: ## ensure dependencies are met
	GO111MODULE=on go mod vendor -v

install: ## install into $GOPATH/bin
	"$(GOCMD)" install -v

build:
	"$(GOCMD)" build -v ${GOFLAGS} ${LDFLAGS} -o "${BINARY}"

test:
	"$(GOCMD)" test -race -v $(shell go list ./... | grep -v '/vendor/')

vet:
	go vet ./...
