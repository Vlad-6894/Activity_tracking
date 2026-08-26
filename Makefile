include .env
export

export PROJECT_ROOT=$(shell pwd)

GO ?= go

# Если пустой TELEGRAM_WEBAPP_URL, то задает его исходя из NGROK_DOMAIN
ifeq ($(strip $(TELEGRAM_WEBAPP_URL)),)
    ifneq ($(strip $(NGROK_DOMAIN)),)
        TELEGRAM_WEBAPP_URL := https://$(NGROK_DOMAIN)
    endif
endif

export TELEGRAM_WEBAPP_URL

# Указываем для makefile что является таргетом,а не файлом для избежания дальнейших конфликтов
.PHONY: help setup run dev tidy check-env \
	env-up env-down env-cleanup env-port-open env-port-close \
	ngrok-up ngrok-down migrate-create migrate-up migrate-down migrate-action

help:
	@echo "setup             .env + postgres + migrations"
	@echo "dev               postgres, migrations, port-forward-open, ngrok - для локальной разработки"
	@echo "run               запускает только приложение, остальное все должно быть запущено"

setup:
	$(GO) mod download
	$(MAKE) env-up
	$(MAKE) migrate-up
	@echo "окружение готово — make dev"

check-env:
	@if [ -z "$(TELEGRAM_BOT_TOKEN)" ]; then echo "TELEGRAM_BOT_TOKEN пустой в .env"; exit 1; fi
	@if [ -z "$(TELEGRAM_WEBAPP_URL)" ]; then \
		echo "укажите NGROK_DOMAIN (или TELEGRAM_WEBAPP_URL) в .env"; exit 1; \
	fi

run: check-env
	@echo "webapp url -> $(TELEGRAM_WEBAPP_URL)"
	$(GO) run ./cmd/activity_tracking

dev:
	$(MAKE) env-up
	$(MAKE) migrate-up
	$(MAKE) env-port-open
	$(MAKE) ngrok-up
	$(MAKE) run

tidy:
	$(GO) mod tidy

env-up:
	docker compose up -d activity-tracking-postgres

env-down:
	docker compose down activity-tracking-postgres activity-tracking-port-forwarder activity-tracking-ngrok

ngrok-up:
	@docker compose up -d activity-tracking-ngrok

ngrok-down:
	@docker compose down activity-tracking-ngrok


env-cleanup:
	@read -p "Очистить все volume файлы окружения? Опасность утери данных. [y.N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down -v activity-tracking-postgres && \
		echo "Файлы окружения очищены"; \
	else \
		echo "Очистка окружения отменена"; \
	fi

env-port-open:
	docker compose up -d activity-tracking-port-forwarder

env-port-close:
	docker compose down activity-tracking-port-forwarder

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутсвует необходимый параметр seq. Пример: make migrate-create seq=init"; \
		exit 1; \
	fi; \
	docker compose run --rm activity-tracking-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	$(MAKE) migrate-action action=up

migrate-down:
	$(MAKE) migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Отсутсвует необходимый параметр action. Пример: make migrate-action action=up"; \
		exit 1; \
	fi; \
	docker compose run --rm activity-tracking-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@activity-tracking-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"
