.PHONY: test
run:
	go run ./cmd/main.go
sql: 
	go run ./cmd/main.go ./test/test_file/file.sql
csv:
	go run ./cmd/main.go ./test/test_file/file.csv