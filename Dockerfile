FROM golang:1.17 as builder

WORKDIR /build

COPY go.mod .

COPY cmd ./cmd

RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o webserver cmd/main.go

FROM bash

WORKDIR /app

COPY --from=builder /build/webserver /app/webserver

EXPOSE 8080/tcp

CMD ["./webserver"]