# Build Golang binary
FROM golang:1.20.2 AS build-golang

WORKDIR /app
COPY ./services/core ./services/core
COPY ./pkg ./pkg
COPY ./proto ./proto
COPY ./migrations ./migrations

RUN go work init && go work use ./services/core ./pkg ./proto

RUN go build -o ./bin/core ./services/core/cmd/core/main.go

# Runtime container
FROM debian:stable-slim

WORKDIR /app
COPY --from=build-golang /app/bin .
COPY --from=build-golang /app/migrations ./migrations
COPY --from=build-golang /app/services/core/Makefile .

RUN apt-get update && apt-get install -y curl && apt-get install make
RUN curl -fsSL https://raw.githubusercontent.com/pressly/goose/master/install.sh | GOOSE_INSTALL=/app/goose sh -s v3.5.0

CMD bash