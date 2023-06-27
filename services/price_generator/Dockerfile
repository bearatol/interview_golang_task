# Build Golang binary
FROM golang:1.20.2 AS build-golang

WORKDIR /app
COPY ./services/price_generator ./services/price_generator
COPY ./pkg ./pkg
COPY ./proto ./proto

RUN go work init && go work use ./services/price_generator ./pkg ./proto

RUN go build -o ./bin/price_generator ./services/price_generator/cmd/price_generator/main.go

# Runtime container
FROM debian:stable-slim

WORKDIR /app
COPY --from=build-golang /app/bin .

CMD ["./price_generator"]