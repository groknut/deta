FROM golang:1.25-alpine

WORKDIR /app

ENV CGO_ENABLED=0

CMD go build -v -o /dist/deta ./main.go
