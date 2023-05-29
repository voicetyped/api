
SERVICE		?= $(shell basename `go list`)
VERSION		?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || cat $(PWD)/.version 2> /dev/null || echo v0)
PACKAGE		?= $(shell go list)
PACKAGES	?= $(shell go list ./...)
FILES		?= $(shell find . -type f -name '*.go' -not -path "./vendor/*")



default: help

help:   ## show this help
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

tidy: ##tidy go lang dependencies
	go mod tidy && go mod vendor

clean:  ## go clean
	go clean

fmt:    ## format the go source files
	go fmt ./...

vet:    ## run go vet on the source files
	go vet -all ./...

doc:    ## generate godocs and start a local documentation webserver on port 8085
	godoc -http=:8085 -index

gen:
	protoc --proto_path=../apis --proto_path=./v1 --go_out=./ --validate_out=lang=go:. transcriber.proto;
	protoc --proto_path=../apis --proto_path=./v1  transcriber.proto --go-grpc_out=./
	mockgen -source=transcriber_grpc.pb.go -self_package=github.com/antinvestor/service-transcriber-api -package=transcriberv1 -destination=transcriber_grpc_mock.go


build: clean fmt vet ## run all preliminary steps and tests the setup

