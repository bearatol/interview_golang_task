# Build Golang binary
FROM golang:1.20.2 AS build-golang

WORKDIR /app
COPY . .
RUN go build -o ./bin/price_generator ./cmd/price_generator/main.go

# Runtime container
FROM debian:stable-slim

WORKDIR /app
COPY --from=build-golang /app/bin .

CMD ["price_generator"]