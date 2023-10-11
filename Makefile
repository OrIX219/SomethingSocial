.PHONY: openapi
openapi:
		@./scripts/openapi.sh auth internal/auth/ports ports
		@./scripts/openapi.sh users internal/users/ports ports
		@./scripts/openapi.sh posts internal/posts/ports ports

.PHONY: proto
proto:
		@./scripts/proto.sh users
		@./scripts/proto.sh posts
		
