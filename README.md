# go-support-assistant

AI-ассистент службы поддержки на Go с использованием DDD, Kafka, gRPC и LangChain.

## Описание проекта

Проект представляет собой микросервисную систему для автоматизации обработки обращений в техподдержку.  
Она включает:

- **API-сервис** (REST на Gin) — приём обращений от пользователей, управление тикетами, взаимодействие с БД и Kafka.
- **AI-сервис** (gRPC) — классификация сообщений, генерация ответов с помощью LangChain.
- **PostgreSQL** — хранение тикетов, сообщений, пользователей и аудита.
- **Redis** — кэширование и временные данные.
- **Kafka** — асинхронная обработка событий (новые сообщения, смена статуса).
- **Docker** — контейнеризация всех сервисов.
- **Prometheus + Grafana** (планируется) — мониторинг и метрики.

## Технологический стек

- Go 1.25
- Gin
- gRPC + Protocol Buffers
- PostgreSQL
- Redis
- Kafka
- Docker / Docker Compose
- LangChain (через библиотеку langchaingo)

## Архитектура

Система состоит из двух основных микросервисов: `api-service` (REST API) и `ai-service` (gRPC).  
`api-service` обрабатывает HTTP-запросы клиентов, сохраняет данные в PostgreSQL, кэширует в Redis и публикует события в Kafka.  
`ai-service` получает сообщения по gRPC от `api-service` (или через Kafka), классифицирует их и генерирует ответ с помощью LangChain.  
Все сервисы запускаются в Docker-контейнерах и взаимодействуют через общую сеть.

## Текущий статус

- ✅ Инфраструктура (PostgreSQL, Redis, Kafka, Zookeeper) поднимается через docker-compose.
- ✅ api-service и ai-service успешно собираются и работают.
- ✅ Базовый gRPC-контракт определён и сгенерирован.
- ⏳ Подключение PostgreSQL к api-service (в процессе).
- ⏳ Реализация DDD-модели.
- ⏳ Интеграция Kafka.
- ⏳ Интеграция LangChain.

## Запуск проекта

1. Убедитесь, что установлены Docker и Docker Compose.

2. Клонируйте репозиторий:
   git clone https://github.com/Konstantin-Korolyov/go-support-assistant.git
   cd go-support-assistant

3. Запустите все сервисы:
docker-compose up -d --build

Проверьте логи:
docker logs support-api
docker logs support-ai

Остановка:
docker-compose down

API (в разработке)
Планируемые эндпоинты:
POST /api/tickets — создать новый тикет.
POST /api/tickets/{id}/messages — добавить сообщение в тикет.
GET /api/tickets/{id} — получить тикет с историей.

Подробная документация будет добавлена после реализации.

Контакты
Автор: Константин Королёв
GitHub: Konstantin-Korolyov