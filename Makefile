BINARY := gogenapi
SOURCES := $(shell find . -name '*.go' -type f | grep -v _examples)

LDFLAGS := -ldflags="-s -w"

.DEFAULT_GOAL := bin/$(BINARY)

bin/$(BINARY): deps $(SOURCES)
	go generate
	go build $(LDFLAGS) -o bin/$(BINARY)

.PHONY: clean
clean:
	rm -fr bin/*
	rm -fr vendor/*

.PHONY: deps
deps:
	go get github.com/jteeuwen/go-bindata/...

.PHONY: install
install:
	go generate
	go install $(LDFLAGS)

.PHONY: test
test:
	go generate
	go test -cover -v ./gogenapi ./command

.PHONY: generation-test
generation-test: bin/$(BINARY)
	script/generation_test.sh
