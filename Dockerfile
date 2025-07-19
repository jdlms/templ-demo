# Multi-stage build
FROM node:22-alpine AS frontend-builder

# Update Alpine packages
RUN apk update && apk upgrade

WORKDIR /app

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy web source files
COPY web/ ./web/

# Build frontend assets
RUN npm run build-css && npm run build

# Go build stage
FROM golang:1.23-alpine AS go-builder

# Update Alpine packages
RUN apk update && apk upgrade

# Install templ
RUN go install github.com/a-h/templ/cmd/templ@latest

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Generate templ files
RUN templ generate

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/server/main.go

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copy the binary
COPY --from=go-builder /app/main .

# Copy static assets from frontend build
COPY --from=frontend-builder /app/web/static ./web/static

# Copy other web assets that might be needed
COPY web/static ./web/static

EXPOSE 3000

CMD ["./main"]