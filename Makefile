REV=$(shell git log --max-count=1 --pretty="format:%h")
GO_VER=$(shell go version|grep "go version"|cut -d' ' -f3|sed "s/[\s\t]*//"|sed "s/^go//")
VER='-X main.Revision=$(REV) -X main.GoVersion=$(GO_VER)'
export GOPATH=$(shell pwd)/../

binary: dep
	go build -ldflags $(VER) -v 

all: binary

debug:
	go build -tags "debug" -ldflags "-X main.EnableDebug=true"

test: 
	go test --cover ./...

dep:
	go get -d -v ./...
