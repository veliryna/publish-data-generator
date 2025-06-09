FROM golang:1.24 AS builder

WORKDIR /app

# Copy Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Build server and client binaries
RUN go build -o /producer ./producer/producer.go
RUN go build -o /client ./client/client.go

FROM debian:bookworm-slim

WORKDIR /app

# Copy built binaries
COPY --from=builder /producer ./producer
COPY --from=builder /client ./client

EXPOSE 50051
EXPOSE 8080