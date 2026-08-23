include .env
export

export PROJECT_ROOT=$(shell pwd)
export DOCKER_UID=$(shell id -u) #для Linux
export DOCKER_GID=$(shell id -g) #для Linux

postgres-up:
	@docker compose up -d activity-tracking-postgres

postgres-down:
	@docker compose down activity-tracking-postgres

postgres-cleanup:
	@read -p "Очистить pg_data? Опасность утери данных. [y/N]: " choice; \
	if [ "$$choice" = "y" ] || [ "$$choice" = "Y" ]; then \
		docker compose down activity-tracking-postgres activity-tracking-port-forwarder && \
		sudo rm -rf ${PROJECT_ROOT}/out/pg_data && \
		echo "Очищено"; \
	else \
		echo "Операция отменена"; \
	fi

port-forward-up:
	@docker compose up -d activity-tracking-port-forwarder

port-forward-down:
	@docker compose down activity-tracking-port-forwarder

ngrok-up:
	@docker compose up -d activity-tracking-ngrok

ngrok-down:
	@docker compose down activity-tracking-ngrok

create-migrate:
	@if [ -z "$(seq)" ]; then \
		echo "Нет параметра seq"; \
		exit 1; \
	fi;
	@docker compose run --rm activity-tracking-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down n=$(or $(n),1)

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Нет параметра action"; \
		exit 1; \
	fi;
	@docker compose run --rm activity-tracking-migrate \
		-path /migrations \
		-database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@activity-tracking-postgres:5432/${POSTGRES_DB}?sslmode=disable" \
		"$(action)" $(n)
