## REST API на Go для управления записями о покупках


## 📖 Описание

Этот сервис предоставляет HTTP-эндпоинты для работы с сущностью **Purchase**:

* Получение списка покупок
* Получение одной покупки по ID
* Создание новой покупки
* Обновление существующей
* Удаление записи

Под капотом:

* Echo - Фреймворк HTTP
* СУБД PostgreSQL (поднимается через Docker Compose)
* Swagger-документация


## 🚀 Быстрый старт

1. **Клонируйте репозиторий**

   ```bash
   git clone https://github.com/shoksin/go-REST-API-purchases.git
   cd go-REST-API-purchases
   ```
2. \*\*Создайте файл \*\***`.env`** в корне и задайте параметры подключения к БД:

   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=secret
   DB_NAME=purchases_db
   PORT=8080
   ```
3. **Запустите базу данных и сервис через Docker Compose**

   ```bash
   docker-compose up -d
   ```

После старта API будет доступно по адресу `http://localhost:8080`.


## 💠 Доступные команды Makefile

* `make run` — компиляция и запуск сервера
* `make build` — сборка бинарника
* `make migrate` — применение миграций к БД
* `make swagger` — генерация/обновление Swagger-документации
* `make docker-build` — сборка Docker-образа
* `make docker-up` — поднять контейнеры через Docker Compose


## 🔌 Эндпоинты

| Метод  | Путь              | Описание                      |
| ------ | ----------------- | ----------------------------- |
| GET    | `/purchases`      | Список всех покупок           |
| GET    | `/purchases/{id}` | Детали одной покупки по ID    |
| POST   | `/purchases`      | Создать новую покупку         |
| PUT    | `/purchases/{id}` | Обновить данные покупки по ID |
| DELETE | `/purchases/{id}` | Удалить покупку по ID         |



## 📁 Swagger-документация

После запуска доступна по адресу:

```
http://localhost:8080/swagger/index.html
```


## 📁 Структура проекта

```
.
├── cmd/
│   └── api/            # Точка входа: main.go
├── app/                # HTTP-хендлеры и маршруты
├── config/             # Конфигурация, миграции
├── internal/           # Бизнес-логика (models, services)
├── middleware/         # Middleware (авторизация)
├── pkg/utils/          # Утилиты и вспомогательные функции
├── swagger/            # Swagger-сборка, схемы
├── docker-compose.yaml # Сервис БД и иные контейнеры
├── Makefile            # Скрипты сборки и запуска
├── go.mod
└── go.sum
```
