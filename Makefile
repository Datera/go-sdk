VERSION ?= v0.1.0
NAME=dsdk

compile:
	@echo "==> Building the Datera Golang SDK"
	@env go get -d ./...
	@env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${NAME} main.go
	@env go vet ./...

compile-local:
	@echo "==> Building the Datera Golang SDK locally"
	@env CGO_ENABLED=0 GOARCH=amd64 go build -o ${NAME} main.go
	@env go vet ./...

clean:
	@echo "==> Cleaning artifacts"
	@GOOS=linux go clean -i -x ./...
