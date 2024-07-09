include .env

BUILD_DIR = $(PWD)/build

.PHONY: dev
dev:
	./bin/air server --port $(APP_PORT)

.PHONY: update-deps
update-deps:
	go get -u && go mod tidy

.PHONY: db-migrate
db-migrate:
	migrate -path src/database/migrations -database "postgresql://$(DB_USERNAME):$(DB_PASSWORD)@localhost:5432/$(DB_DATABASE)?sslmode=disable" -verbose up

.PHONY: sqlc-docker
sqlc-docker:
	docker pull sqlc/sqlc

.PHONY: sqlc-generate
sqlc-generate: sqlc-docker
	docker run --rm -v $(PWD):/src -w /src sqlc/sqlc generate
