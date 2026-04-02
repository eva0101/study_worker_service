include .env
export

service-run:
	@go run main.go

run-http-app:
	docker run -p 5050:5050 employees

migrate-up:
	@migrate -path migrations -database ${CONN_STRING} up

migrate-down:
	@migrate -path migrations -database ${CONN_STRING} down