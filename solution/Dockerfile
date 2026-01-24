FROM golang:1.25-alpine AS builder

# Setup base software for building an app.
RUN apk add --no-cache ca-certificates git gcc g++ libc-dev binutils

WORKDIR /opt

# Download dependencies.
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy application source.
COPY cmd ./cmd/
COPY internal ./internal/
COPY pkg ./pkg/
COPY platform ./platform/
COPY migrations ./migrations/

# Debug: Check if directories exist
RUN ls -la ./cmd && ls -la ./migrations

RUN go build -o bin/application ./cmd/main.go

# Prepare executor image.
FROM alpine:3.21 AS runner

RUN apk add --no-cache ca-certificates libc6-compat openssh bash curl

WORKDIR /opt

COPY --from=builder /opt/bin/application ./
COPY --from=builder /opt/migrations ./migrations

# Run the application.
CMD ["./application"]
