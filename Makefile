.PHONY: openapi
openapi:
		@./scripts/openapi.sh auth internal/auth/ports ports
		@./scripts/openapi.sh users internal/users/ports ports
		@./scripts/openapi.sh posts internal/posts/ports ports

.PHONY: proto
proto:
		@./scripts/proto.sh users
		@./scripts/proto.sh posts
		
.PHONY: env
env:
	@$(eval SHELL:=/bin/bash)
	@cp .env.example .env
	@sed -i 's/AUTH_SECRET=/echo &`openssl rand -hex 32`/e' .env
	@sed -i 's/AUTH_SALT=/echo &`openssl rand -hex 16`/e' .env
