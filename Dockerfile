FROM golang:1.23

# Set destination for COPY
WORKDIR /app

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY ./ ./

RUN go mod download
RUN go mod tidy

WORKDIR /app/cmd/gowasp
# Build
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-w -s" -o /go/bin/gowasp

EXPOSE 8080

ENV MIGRATION_SOURCE_URL="file:///app/resources/migrations"
ENV TEMPLATES_PATH="/app/web/templates/**/*"

# Run
ENTRYPOINT ["/go/bin/gowasp"]