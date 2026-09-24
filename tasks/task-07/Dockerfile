FROM golang:1.21-alpine3.22

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    PORT=8080

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /build/app

RUN mkdir -p /build

ENTRYPOINT ["sh", "-c", "go install github.com/githubnemo/CompileDaemon@latest && CompileDaemon -build='go build -o /build/app' -command='/build/app'"]