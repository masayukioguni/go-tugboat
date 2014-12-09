BIN = go-tugboat

all: clean build test

setup:
	go get github.com/tools/godep
	go get github.com/masayukioguni/go-digitalocean/digitalocean
	go get github.com/golang/lint/golint

test: deps
	go test $(TESTFLAGS) ./...

build: deps
	go build -o build/$(BIN)

run: build
	./build/$(BIN)

deps:
	godep get 

clean:
	rm -f build/$(BIN)
	go clean

lint:
	golint ./...

vet:
	go vet ./...

coverage:
	bash coverage
	rm api/profile.coverprofile
	rm config/profile.coverprofile
	rm gover.coverprofile

.PHONY: setup test build run deps clean lint vet coverage