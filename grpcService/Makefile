-include .env

.PHONY: init
init:
	@cp .env.dist .env

.PHONY: start
start:
	@docker-compose up -d

.PHONY: down-v
down-v:
	@docker-compose down -v --remove-orphans
	@docker-compose ps