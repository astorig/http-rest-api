.PHONY: build
.PHONY: migrate

DB_URL ?= postgres://restapi:restapi@127.0.0.1:63341/restapi_dev?sslmode=disable
DB_TEST_URL ?= postgres://restapi:restapi@127.0.0.1:63343/restapi_test?sslmode=disable
MIGRATIONS_DIR := $(CURDIR)/migrations

build:
	docker compose -f docker-compose-local.yml build
	go build -v -o ./bin/apiserver ./cmd/apiserver
	docker compose -f docker-compose-local.yml up -d

test:
	docker compose -f docker-compose-test.yml up -d
	sleep 10s
	docker run --rm --network host -v "$(MIGRATIONS_DIR):/migrations" migrate/migrate \
    		-path=/migrations \
    		-database "$(DB_TEST_URL)" up
	go test -v -race -timeout 30s ./...
	docker compose -f docker-compose-test.yml down

migrate:
	docker run --rm --network host -v "$(MIGRATIONS_DIR):/migrations" migrate/migrate \
		-path=/migrations \
		-database "$(DB_URL)" up

migrate_down:
	docker run --rm --network host -v "$(MIGRATIONS_DIR):/migrations" migrate/migrate \
    		-path=/migrations \
    		-database "$(DB_URL)" down

.DEFAULT_GOAL := build

