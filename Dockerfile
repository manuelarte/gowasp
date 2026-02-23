ARG VERSION=1.25

# Build Stage
FROM golang:${VERSION}-alpine AS builder
# hadolint ignore=DL3018
RUN apk --no-cache add ca-certificates dumb-init make git gcc libtool musl-dev nodejs npm \
    && npm install -g pnpm@10.19.0

WORKDIR /app

# Copy the source code
COPY ./ ./

# build backend and frontend
RUN pnpm -C web install && pnpm -C web build \
    && go mod download && go mod tidy

WORKDIR /app/cmd/gowasp

# Build the binary
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-w -s" -o /gowasp

# Final Stage
FROM alpine:3

# hadolint ignore=DL3018
RUN apk --no-cache add ca-certificates dumb-init

# Copy the binary from builder stage
COPY --from=builder /gowasp /usr/local/bin/gowasp

# Run
ENTRYPOINT ["/usr/local/bin/gowasp"]
