## Environment:

.PHONY: local-up
local-up: docker-up migrate-schema ## Bring up the development environment

.PHONY: docker-up
docker-up:
	@docker-compose up -d --wait

.PHONY: migrate-schema
migrate-schema:
	@echo "TODO: add migrations for posgres" 

.PHONY: local-down
local-down: ## Bring down the development environment
	@docker-compose down