# Build stage
FROM golang:1.24.1-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./main.go

# Runtime stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main /app/main

EXPOSE 3000
CMD ["/app/main"]