FROM golang:1.23-alpine3.21 AS builder

# Setup base software for building an app.
RUN apk add --no-cache ca-certificates git gcc g++ libc-dev binutils

WORKDIR /opt

# Download dependencies.
COPY solution/go.mod solution/go.sum ./
RUN go mod download && go mod verify

# Copy application source.
COPY ./solution .

# Build the application.
RUN go build -o bin/application .

# Prepare executor image.
FROM alpine:3.21 AS runner

RUN apk add --no-cache ca-certificates libc6-compat openssh bash curl

WORKDIR /opt

COPY --from=builder /opt/bin/application ./

# Run the application.
CMD ["./application"]