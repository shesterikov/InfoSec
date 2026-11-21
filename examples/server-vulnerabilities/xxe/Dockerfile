FROM golang:1.25-bookworm AS build
WORKDIR /src
RUN apt-get update && apt-get install -y --no-install-recommends libxml2-dev pkg-config && rm -rf /var/lib/apt/lists/*
COPY . .
RUN go mod tidy && go test ./... && go vet ./... && go build -o /server ./cmd/server
FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y --no-install-recommends libxml2 && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=build /server /app/server
COPY examples/secret.txt /app/data/secret.txt
USER 65532:65532
ENV HTTP_ADDR=0.0.0.0:8080
EXPOSE 8080
ENTRYPOINT ["/app/server"]
