GOPATH=$(shell pwd)

all:
	env GOPATH=${GOPATH} go get datera-api/dapi
	env GOPATH=${GOPATH} go build datera-api/dapi

clean:
	rm -f -- datera-api
	rm -f -- datera_api.log
	rm -rf -- bin
	rm -rf -- pkg
	rm -rf -- src/github.com
	rm -rf -- src/golang.com

test:
	env GOPATH=${GOPATH} go get datera-api/dapi
	env GOPATH=${GOPATH} go build datera-api/dapi
	env GOPATH=${GOPATH} go get github.com/stretchr/testify/assert
	env GOPATH=${GOPATH} go get github.com/stretchr/testify/mock
	env GOPATH=${GOPATH} go test -v datera-api/test

fmt:
	env GOPATH=${GOPATH} go fmt datera-api
	env GOPATH=${GOPATH} go fmt datera-api-test
