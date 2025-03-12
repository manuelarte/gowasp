ARG VERSION=1.23.7

FROM golang:${VERSION}-alpine AS builder
RUN apk --no-cache add make git gcc libtool musl-dev ca-certificates dumb-init

# Set destination for COPY
WORKDIR /app

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY ./ ./

RUN go mod download && go mod tidy

WORKDIR /app/cmd/gowasp
# Build
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-w -s" -o /go/bin/gowasp

EXPOSE 8080

ENV MIGRATION_SOURCE_URL="file:///app/resources/migrations" \
    TEMPLATES_PATH="/app/web/templates/**/*"

# Run
ENTRYPOINT ["/go/bin/gowasp"]