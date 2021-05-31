FROM golang:1.16.4 as builder
WORKDIR /opt/app

# Install migration tool
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

# Copy code from host and compile
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd ./cmd
COPY gen ./gen
COPY internal ./internal
COPY migrations ./migrations
COPY static ./static

ENTRYPOINT go run cmd/api.go
