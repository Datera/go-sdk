VERSION ?= v0.1.0
NAME=dsdk

compile:
	@echo "==> Building the Datera Golang SDK"
	@env go get -d ./...
	@env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./pkg/v2/dsdk
	@env go vet ./...

local:
	@echo "==> Building the Datera Golang SDK locally"
	@env CGO_ENABLED=0 GOARCH=amd64 go build ./pkg/v2/dsdk
	@env go vet ./...

clean:
	@echo "==> Cleaning artifacts"
	@GOOS=linux go clean -i -x ./...

test:
	@echo "==> Testing the Datera Golang SDK"
	@env go test ./tests/*.go

testv:
	@echo "==> Testing the Datera Golang SDK (verbose)"
	@env go test -v ./tests/*.go

testc:
	@echo "==> Testing the Datera Golang SDK (compile only)"
	@env go test -c ./tests/*.go
