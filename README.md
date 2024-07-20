### Тестовое задание для Messaggio

FRIST STEP: docker-compose -f ./deploy/docker-compose.yml --env-file ./configs/.env up -d --remove-orphans --build

WAITING START...

SECONDS STEP: goose -dir schema postgres 'postgresql://Messaggio:somestrongpassword@localhost:5432/db_messages' up