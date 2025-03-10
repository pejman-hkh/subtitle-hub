FROM golang:1.24 AS builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/api /app/.env.production ./

CMD ["./api", "-p"]