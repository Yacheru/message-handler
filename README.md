### Тестовое задание для Messaggio

*Разработать микросервис на Go, который будет принимать сообщения через HTTP API, сохранять их в PostgreSQL, 
а затем отправлять в Kafka для дальнейшей обработки. Обработанные сообщения должны помечаться. Сервис должен также 
предоставлять API для получения статистики по обработанным сообщениям.*

### Ручки:

- POST   /messaggio/               Добавить сообщение
- GET    /messaggio/               Получить все сообщения
- GET    /messaggio/stats          Получить статистику по сообщениям
- GET    /messaggio/:id            Получить сообщение по его UUID 
- DELETE /messaggio/:id            Удалить сообщение по его UUID
- PATCH  /messaggio/:id            Изменить сообщение по его UUID

### Запуск локально...

1. docker-compose -f ./deploy/docker-compose.yml --env-file ./configs/.env up -d --remove-orphans --build

2. Ожидаем запуска...

3. [goose](https://pressly.github.io/goose/installation/) -dir schema postgres 'postgresql://Messaggio:somestrongpassword@localhost:5432/db_messages' up

### Переменные окружения

| Название               | Описание                            | Дефолтное значение                                                                    |
|------------------------|-------------------------------------|---------------------------------------------------------------------------------------|
| `API_PORT`             | Порт веб-сервера                    | `8080`                                                                                |
| `API_DEBUG`            | Работать ли серверу в DEBUG-режиме  | `de_dust2`                                                                            |
| `KAFKA_CONSUMER_GROUP` | Название консьюмер группы           | `consumergroup`                                                                       |
| `KAFKA_BROKER`         | Адрес брокера сообщений             | `kafka-1:9092`                                                                        |
| `KAFKA_TOPIC`          | Название топика сообщений           | `messages_topic`                                                                      |
| `POSTGRES_PASSWORD`    | Пароль для пользователя базы данных | `somestrongpassword`                                                                  |
| `POSTGRES_USER`        | Имя пользователя базы данных        | `Messaggio`                                                                           |
| `POSTGRES_DB`          | Название базы данных                | `db_messages`                                                                         |
| `POSTGRES_DSN`         | DSN для подключения к базе          | `postgresql://Messaggio:somestrongpassword@postgres:5432/db_messages?sslmode=disable` |