# Simple backend on Go + YDB

Отдельный пример такого же backend API (`POST/GET /api/favorites`), но с хранением в YDB.

## Переменные окружения

- `YDB_DSN` - DSN до базы YDB (обязательно)
- `YDB_KEY_FILE` - путь до service account key json (обязательно)
- `YDB_TABLE` - имя таблицы (опционально, по умолчанию `favorites`)
- `PORT` - порт для локального запуска (опционально, по умолчанию `8080`)

## Локальный запуск

```bash
go run .
```

## Схема таблицы

Ожидается таблица с именем из `YDB_TABLE` и колонкой `city` (тип UTF8, PK).
