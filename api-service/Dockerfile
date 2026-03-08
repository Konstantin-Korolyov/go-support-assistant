FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod* ./
RUN go mod download

COPY . .

RUN go build -o api-service ./cmd/api

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/api-service .

EXPOSE 8080

CMD ["./api-service"]