.PHONY: test
run:
	go run ./cmd/main.go
sql: 
	go run ./cmd/main.go ./test/test_file/file.sql
csv:
	go run ./cmd/main.go ./test/test_file/file.csv

json: 
	go run ./cmd/main.go ./test/test_file/file.json
b-linux:
	GOOS=linux GOARCH=amd64 go build -o deta ./cmd/main.go

b-windows:
	GOOS=windows GOARCH=amd64 go build -o deta.exe ./cmd/main.go

b-mac-int:
	GOOS=darwin GOARCH=amd64 go build -o deta

b-mac:
	GOOS=darwin GOARCH=arm64 go build -o deta