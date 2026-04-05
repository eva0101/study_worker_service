include .env
export

compose-up:
	docker compose up -d --build

compose-down:
	docker compose down

migrate-up:
	@migrate -path migrations -database ${CONN_LOCAL} up

migrate-down:
	@migrate -path migrations -database ${CONN_LOCAL} down	