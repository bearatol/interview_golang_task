# Build Golang binary
FROM golang:1.20.2 AS build-golang

WORKDIR /app
COPY . .
RUN go build -o ./bin/core ./cmd/core/main.go

# Runtime container
FROM debian:stable-slim

WORKDIR /app
COPY --from=build-golang /app/bin .

CMD ["./core"]