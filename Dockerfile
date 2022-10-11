FROM golang:1.12-alpine AS build

WORKDIR /app

COPY *.go .

RUN go build -o test .

FROM alpine:3.16.2
WORKDIR /
ENV GOLANG_PORT $GOLANG_PORT
COPY --from=build /app/test ./test

EXPOSE $GOLANG_PORT

CMD ["./test"]