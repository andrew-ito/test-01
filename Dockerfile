FROM golang:1.12-alpine AS build

WORKDIR /app

COPY *.go .

RUN go build -o test .

FROM alpine:3.16.2
WORKDIR /

COPY --from=build /app/test ./test

EXPOSE 8080

CMD ["./test"]