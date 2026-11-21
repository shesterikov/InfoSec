FROM golang:1.25-bookworm AS build
WORKDIR /src
COPY . .
RUN go mod tidy && go test ./... && go vet ./... && go build -o /server ./cmd/server
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=build /server /app/server
USER 65532:65532
ENV HTTP_ADDR=0.0.0.0:8080
EXPOSE 8080
ENTRYPOINT ["/app/server"]
