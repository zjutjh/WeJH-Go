build:
	go build -o wejh

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o wejh

.PHONY: build build-linux