# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install templ CLI
RUN go install github.com/a-h/templ/cmd/templ@latest

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Generate templ files
RUN templ generate

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/api ./cmd/api

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/bin/api .

# Copy static files (CSS, JS, images)
COPY --from=builder /app/static ./static

# Copy templates (if you need the original .templ files for debugging)
# COPY --from=builder /app/templates ./templates

EXPOSE 8080

CMD ["./api"]