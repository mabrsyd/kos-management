# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

# Final stage
FROM alpine:3.18

WORKDIR /app

# Install bash
RUN apk add --no-cache bash

COPY --from=builder /app/main .

# Tambahkan konfigurasi keamanan
ENV GIN_MODE=release
ENV ALLOWED_ORIGINS="https://kos-management.com,https://localhost:3443"

# Jalankan binary langsung, gak perlu bash/sh
CMD ["./main"]
