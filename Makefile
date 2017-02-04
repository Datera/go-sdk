GOPATH=$(shell pwd)

all:
	# test -e src/datera_api/dapi/types.go || echo "Please run the schema-parser.py script"; exit 1
	env GOPATH=${GOPATH} go get dsdk
	env GOPATH=${GOPATH} go build dsdk

clean:
	rm -f -- datera_api.log
	rm -rf -- bin
	rm -rf -- pkg
	rm -rf -- src/github.com
	rm -rf -- src/golang.com

test:
	env GOPATH=${GOPATH} go get dsdk
	env GOPATH=${GOPATH} go build dsdk
	env GOPATH=${GOPATH} go get github.com/stretchr/testify/assert
	env GOPATH=${GOPATH} go get github.com/stretchr/testify/mock
	env GOPATH=${GOPATH} go test -v dsdk/test

fmt:
	env GOPATH=${GOPATH} go fmt dsdk
