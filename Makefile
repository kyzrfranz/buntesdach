# generate all the make targets to build cmd/server.go
GOROOT=$(shell go env GOROOT)
GO=$(GOROOT)/bin/go

build:
	$(GO) build -o ./bin/server ./cmd/server.go

run: build
	./bin/server

dev:
	$(GO) run ./cmd/server.go