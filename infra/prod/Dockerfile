FROM golang:1.16.4 as builder
WORKDIR /opt/app

# Copy code from host and compile
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd ./cmd
COPY gen ./gen
COPY internal ./internal

RUN go build -o /bin/api cmd/api.go

# Copy binary to debian and run
FROM debian:jessie-slim

# Need ca-certificates to make https requests from container
RUN apt-get update
RUN apt-get install -y ca-certificates

COPY --from=builder /bin/api /bin/api

# Start API
ENTRYPOINT /bin/api
