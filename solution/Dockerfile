FROM golang:1.25-alpine AS builder

# Setup base software for building an app.
RUN apk add --no-cache ca-certificates git gcc g++ libc-dev binutils

WORKDIR /opt

# Download dependencies.
COPY solution/go.mod solution/go.sum ./
RUN go mod download && go mod verify

# Copy application source.
COPY solution/cmd solution/internal solution/pkg solution/platform solution/migrations ./

# Move migrations to correct location.
RUN mv ./migrations ./migrations

RUN go build -o bin/application ./cmd/main.go

# Prepare executor image.
FROM alpine:3.21 AS runner

RUN apk add --no-cache ca-certificates libc6-compat openssh bash curl

WORKDIR /opt

COPY --from=builder /opt/bin/application ./
COPY --from=builder /opt/migrations ./migrations

# Run the application.
CMD ["./application"]
