
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)


AppName ?= qcloud-cdn-flusher

build: tidy
	CGO_ENABLED=0 go build -o ./out/$(AppName)-$(GOOS)-$(GOARCH)

buildx:
	GOOS=linux GOARCH=amd64 make build
	GOOS=linux GOARCH=arm64 make build
	GOOS=darwin GOARCH=amd64 make build
	GOOS=darwin GOARCH=arm64 make build

tidy:
	go mod tidy

clear:
	rm -rf ./out
