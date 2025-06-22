# Build stage
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server cmd/main.go

# Final stage
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/server .
COPY .env .env

EXPOSE 8080

CMD ["./server"]
