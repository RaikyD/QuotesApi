# Quotes API

Простой REST‑сервис на Go для хранения цитат.

## Требования

* Go 1.22 или новее
* Утилита `swag` для генерации Swagger‑описания:

  ```bash
  go install github.com/swaggo/swag/cmd/swag@latest
  ```

## Запуск

```bash
git clone https://github.com/RaikyD/QuotesApi

# установка зависимостей
go mod tidy        

swag init -g main.go   
go run ./...
```

Сервис будет доступен на `http://localhost:8080`. Документация Swagger — `http://localhost:8080/swagger/index.html`.

## Тесты

```bash
go test ./...
```

## Структура проекта

```
.
├── cmd/            # файлы с функцией main
├── internal/
│   ├── delivery/   # HTTP‑слой (handlers)
│   ├── usecases/   # бизнес‑логика
│   ├── repositories/ # хранилище данных
│   └── entity/     # доменные сущности
└── docs/           # swagger‑файлы (генерируются)
```

## API

| Метод  | Путь           | Описание                                          |
| ------ | -------------- | ------------------------------------------------- |
| POST   | /quotes        | Создать цитату                                    |
| GET    | /quotes        | Получить все цитаты или отфильтровать по `author` |
| GET    | /quotes/random | Получить случайную цитату                         |
| DELETE | /quotes/{id}   | Удалить цитату по ID                              |

В ответ на POST `/quotes` сервис возвращает созданную цитату с сгенерированным полем `id`. Сделано для удобства при тестировании с curl
