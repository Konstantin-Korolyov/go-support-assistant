FROM golang:1.25-alpine AS builder

WORKDIR /app

# Копируем go.mod и go.sum из папки ai-service
COPY ai-service/go.mod ai-service/go.sum ./

# Копируем папку proto в /app/proto (соответствует replace ./proto)
COPY proto /proto

# Загружаем зависимости
RUN go mod download

# Копируем остальной код ai-service
COPY ai-service/ ./

# Приводим go.mod в порядок (обновляет его в соответствии с импортами)
RUN go mod tidy

# Собираем бинарник
RUN go build -o ai-service ./cmd/ai

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/ai-service .

EXPOSE 50051

CMD ["./ai-service"]