FROM golang:1.26.1-bookworm

WORKDIR /app

COPY . .

RUN go mod download

CMD ["go", "run", "./cmd/main.go"]

