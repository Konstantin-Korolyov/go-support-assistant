# go-support-assistant

AI-ассистент службы поддержки на Go с использованием DDD, Kafka, gRPC и LangChain.

## Технологии

- Go (Gin, gRPC, langchaingo)
- PostgreSQL
- Redis
- Kafka
- Docker
- Prometheus + Grafana (планируется)

## Архитектура

- **api-service** — публичное REST API (Gin), работа с данными, отправка событий в Kafka.
- **ai-service** — gRPC-сервис, использующий LangChain для классификации и генерации ответов.
- Взаимодействие между сервисами через gRPC.
- Асинхронность через Kafka.
- Хранение: PostgreSQL (основные данные), Redis (кэш).

## Запуск

```bash
docker-compose up -d