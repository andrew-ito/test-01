FROM golang:1.18.7-alpine3.16 as build

WORKDIR /app
COPY *.go ./

EXPOSE 8080
CMD go run ./main.go