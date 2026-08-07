# Knifes

CRUD-приложение для каталога ножей. Go (Fiber) + PostgreSQL + React (Vite).

## Требования

Через Docker (рекомендуется):
- Docker + Docker Compose
- Node.js 18+ (только для фронтенда, он в Docker не поднимается)

Без Docker:
- Go 1.26+
- Node.js 18+
- PostgreSQL 16+
- MinIO (или любое S3-совместимое хранилище) — для фото ножей

## Быстрый старт

### Через Docker Compose

```bash
docker compose up --build
```

Поднимает разом Postgres, MinIO и сам бэкенд, всё в одной docker-сети. Таблицы (`migrations/001_init.sql`) применяются автоматически при первом старте контейнера Postgres — вручную ничего создавать не нужно. Бэкенд ждёт, пока Postgres и MinIO пройдут healthcheck, и только потом стартует.

- Бэкенд: `http://localhost:8080`
- MinIO консоль: `http://localhost:9001` (`minioadmin` / `minioadmin`)

Фронтенд в compose не входит, поднимается отдельно:

```bash
cd web
npm install
npm run dev
```

Откроется на `http://localhost:5173`.

### Без Docker

**1. База данных**

```sql
CREATE DATABASE knives;
```

Таблицы — вручную, тем же SQL, что и в `migrations/001_init.sql`:

```sql
CREATE TABLE knives (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    description TEXT,
    price INTEGER NOT NULL DEFAULT 0,
    material TEXT,
    blade_length NUMERIC,
    handle TEXT,
    brand TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE TABLE knife_photos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    knife_id UUID NOT NULL REFERENCES knives(id),
    s3_key TEXT NOT NULL,
    filename TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);
```

**2. MinIO** — поставить и запустить самостоятельно (например, скачать бинарник с minio.io), либо поднять только его контейнер: `docker compose up minio`.

**3. Бэкенд**

```bash
go mod download
go run cmd/server/main.go
```

Сервер запустится на `http://localhost:8080`.

Настройки берутся из переменных окружения (или `.env` в корне проекта):

| Переменная | По умолчанию | Описание |
|---|---|---|
| `PORT` | `8080` | Порт HTTP-сервера |
| `DATABASE_URL` | `postgres://postgres:postgres@localhost:5432/knives?sslmode=disable` | Строка подключения к Postgres |
| `MINIO_ENDPOINT` | `localhost:9000` | Адрес MinIO |
| `MINIO_ACCESS_KEY` | `minioadmin` | Access key MinIO |
| `MINIO_SECRET_KEY` | `minioadmin` | Secret key MinIO |
| `MINIO_BUCKET` | `knife-photos` | Бакет для фото ножей |

**4. Фронтенд**

```bash
cd web
npm install
npm run dev
```

Откроется на `http://localhost:5173`.

## API

| Метод | Путь | Описание |
|-------|------|----------|
| GET | `/api/v1/knives` | Получить все ножи |
| GET | `/api/v1/knives/:id` | Получить нож по ID |
| POST | `/api/v1/knives` | Создать нож |
| PATCH | `/api/v1/knives/:id` | Обновить нож |
| DELETE | `/api/v1/knives/:id` | Удалить нож (мягкое) |
| POST | `/api/v1/knives/:id/photos` | Загрузить фото ножа |
| GET | `/api/v1/knives/:id/photos` | Список фото ножа |
| GET | `/api/v1/knives/:id/photos/:photoId/file` | Получить файл фото |
| DELETE | `/api/v1/knives/:id/photos/:photoId` | Удалить фото |

## Стек

**Бэкенд:** Go, Fiber, pgx, PostgreSQL, MinIO

**Фронтенд:** React, Vite

## Архитектура бэкенда

Слоистая архитектура с явным разделением портов и адаптеров:

```
cmd/server            точка сборки: создаёт адаптеры и внедряет их в service/handler

internal/domain        доменные сущности (Knife, KnifePhoto), без внешних зависимостей
internal/repository     порты (интерфейсы) доступа к данным
internal/repository/postgres   адаптер: реализация портов поверх pgx/Postgres
internal/service        бизнес-логика и валидация, знает только домен и порты
internal/handler        HTTP-слой (Fiber)
internal/handler/dto    DTO запросов/ответов + мапинг в/из domain
internal/storage        порт файлового хранилища
internal/storage/minio   адаптер: реализация порта поверх MinIO
```
