# Stage 1: Build
FROM golang:1.21 AS builder
WORKDIR /app

# Copy go.mod and go.sum for dependency caching
COPY go.mod go.sum ./
RUN go mod download

# Copy all source files
COPY . .

# Build server and client binaries
RUN go build -o chat-server ./server.go
RUN go build -o chat-client ./client.go

# Stage 2: Runtime
FROM debian:bullseye-slim
WORKDIR /app

# Copy binaries from builder
COPY --from=builder /app/chat-server .
COPY --from=builder /app/chat-client .

# Expose server port
EXPOSE 8080

# Default command (can override in Compose)
CMD ["./chat-server"]
