FROM golang:1.26-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./cmd
COPY internal ./internal
RUN CGO_ENABLED=0 go build -trimpath -o /chat ./cmd/chat

FROM alpine:3.23
RUN adduser -D -u 10001 student
COPY --from=build /chat /usr/local/bin/chat
USER student
EXPOSE 8080
CMD ["chat"]
