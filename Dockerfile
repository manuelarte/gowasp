ARG VERSION=1.24.0

# Build Stage
FROM golang:${VERSION}-alpine AS builder
# hadolint ignore=DL3018
RUN apk --no-cache add make git gcc libtool musl-dev ca-certificates dumb-init

WORKDIR /app

# Copy the source code
COPY ./ ./

RUN go mod download && go mod tidy

WORKDIR /app/cmd/gowasp

# Build the binary
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-w -s" -o /gowasp

# Final Stage
FROM alpine:3
# hadolint ignore=DL3018
RUN apk --no-cache add ca-certificates dumb-init

# Copy the binary from builder stage
COPY --from=builder /app/resources/migrations /app/migrations
COPY --from=builder /app/web /app/web

# Copy the binary from builder stage
COPY --from=builder /gowasp /usr/local/bin/gowasp

EXPOSE 8083

ENV MIGRATION_SOURCE_URL="file:///app/migrations" \
  WEB_PATH="/app/web"

# Run
ENTRYPOINT ["/usr/local/bin/gowasp"]
