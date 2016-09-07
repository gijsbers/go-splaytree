.PHONY: all

all:
	@go build
	@golint
	@go test
