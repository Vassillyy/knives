# Knifes

CRUD-приложение для каталога ножей. Go (Fiber) + PostgreSQL + React (Vite).

## Требования

- Go 1.26+
- Node.js 18+
- PostgreSQL 16+

## Быстрый старт

### 1. База данных

```sql
CREATE DATABASE knifes;
```

Таблица создаётся автоматически при первом запуске (или вручную):

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
```

### 2. Бэкенд

```bash
go mod download
go run cmd/server/main.go
```

Сервер запустится на `http://localhost:8080`.

### 3. Фронтенд

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

## Структура проекта

```
knifes/
├── cmd/server/main.go        # Точка входа
├── internal/
│   ├── config/               # Конфигурация
│   ├── handler/              # HTTP-хендлеры
│   ├── models/               # Модели данных
│   ├── repository/           # Работа с БД
│   │   └── queries/          # SQL-запросы
│   └── service/              # Бизнес-логика
└── web/                      # React-фронтенд
```

## Стек

**Бэкенд:** Go, Fiber, pgx, PostgreSQL

**Фронтенд:** React, Vite
