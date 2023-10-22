.PHONY: openapi
openapi:
		@./scripts/openapi.sh auth internal/auth/ports ports
		@./scripts/openapi.sh users internal/users/ports ports
		@./scripts/openapi.sh posts internal/posts/ports ports

.PHONY: proto
proto:
		@./scripts/proto.sh users
		@./scripts/proto.sh posts
		
.env:
	@cp .env.example .env
	@sed -i 's/AUTH_SECRET=/echo &`openssl rand -hex 32`/e' .env
	@sed -i 's/AUTH_SALT=/echo &`openssl rand -hex 16`/e' .env

include .env

.PHONY: tidy
tidy:
	@cd internal/auth && go mod tidy
	@cd internal/users && go mod tidy
	@cd internal/posts && go mod tidy

.PHONY: docker
docker:
	docker compose up -d

.PHONY: migrate
migrate:
	goose -dir migrations/auth postgres "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:3100/postgres?sslmode=disable" up
	goose -dir migrations/users postgres "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:3200/postgres?sslmode=disable" up
	goose -dir migrations/posts postgres "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:3300/postgres?sslmode=disable" up

.PHONY: up
up:
	make .env
	make tidy
	make docker
	make migrate
