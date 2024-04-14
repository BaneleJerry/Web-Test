include backend/.env
.PHONY: start build run env schema-up schema-down

start:
	@cd frontend/web-test && npm start

build:
	@cd backend && go build -o bin/web-test

run: build
	@cd backend && ./bin/web-test

sqlc-generate:
	@cd backend && sqlc generate

env:
	@echo $(DB_URL)
