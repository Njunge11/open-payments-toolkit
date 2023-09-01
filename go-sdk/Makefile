.DEFAULT_GOAL := build

.PHONY: fmt vet check-libs build run

fmt:
	@go fmt ./...

vet: fmt
	@go vet ./...

check-libs:
	@go build -o /dev/null ./payment/

build: vet check-libs
	@go build -o go-sdk ./cmd/

run: build
	@./go-sdk
