# syntax=docker/dockerfile:1
FROM golang:1.26.1-bookworm AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/server ./cmd/main.go

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates curl \
    && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=build /out/server /usr/bin/server
COPY internals/postgres/migrations ./migrations
EXPOSE 3000
USER 10001:10001
CMD ["/usr/bin/server"]
