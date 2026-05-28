
.DEFAULT_GOAL := help
.PHONY: test

help:
	@echo "deta - data viewer"
	@echo "  run: start program with golang"
	@echo "  test: test program with TEST_FILE"
	@echo "  b-linux: build for linux"
	@echo "  b-windows: build for windows"
	@echo "  b-mac-int: build for mac with intel"
	@echo "  b-mac: build for arm64 mac"



TEST_FILE ?= test/test_file/file.json

run:
	go run ./main.go

test:
	go run ./main.go $(TEST_FILE)

test-all:
	go test -v ./test/

b-linux:
	GOOS=linux GOARCH=amd64 go build -o deta ./main.go

b-windows:
	GOOS=windows GOARCH=amd64 go build -o deta.exe ./main.go

b-mac-int:
	GOOS=darwin GOARCH=amd64 go build -o deta

b-mac:
	GOOS=darwin GOARCH=arm64 go build -o deta
