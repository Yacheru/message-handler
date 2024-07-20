message-start:
	docker-compose -f ./deploy/docker-compose.yml --env-file ./configs/.env up -d --remove-orphans --build

migration-up:
	goose -dir schema postgres 'postgresql://Messaggio:somestrongpassword@localhost:5432/db_messages' up