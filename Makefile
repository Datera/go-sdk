GOPATH=$(shell pwd)

all:
	env GOPATH=${GOPATH} go get dsdk
	env GOPATH=${GOPATH} go build dsdk
	env GOPATH=${GOPATH} go vet dsdk

clean:
	rm -f -- datera_api.log
	rm -rf -- bin
	rm -rf -- pkg
	rm -rf -- src/github.com
	rm -rf -- src/golang.com

test:
	env GOPATH=${GOPATH} go get dsdk
	env GOPATH=${GOPATH} go build dsdk
	env GOPATH=${GOPATH} go test -v dsdk/test

fmt:
	env GOPATH=${GOPATH} go fmt dsdk

doc:
	env GOPATH=${GOPATH} godoc -html dsdk
