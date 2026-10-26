GO           ?= go
COMPOSE      ?= docker compose
DB_HOST_PORT ?= 5433

.PHONY: help run up db-up down tidy

help: ## Отобразить список доступных команд
	@echo "Использование: make [цель]"
	@echo ""
	@echo "Доступные команды:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)


db-up: ## Поднять только PostgreSQL (localhost:5433) для локального запуска
	@echo "==> Поднимаю PostgreSQL..."
	$(COMPOSE) up -d db
	@echo "==> БД готова: postgres://user:password@localhost:$(DB_HOST_PORT)/tictactoe_db — пропишите в src/.env"

run: ## Запустить приложение локально (нужен src/.env и запущенная БД)
	@echo "==> Запуск: http://localhost:8080"
	cd src && $(GO) run ./cmd/app
	

up: ## Поднять весь стек в docker: БД + приложение на http://localhost:8080
	@echo "==> Сборка и запуск стека..."
	$(COMPOSE) up -d --build
	@echo "==> Приложение: http://localhost:8080"

down: ## Остановить контейнеры (данные БД в volume сохраняются)
	@echo "==> Останавливаю контейнеры..."
	$(COMPOSE) down


tidy: ## Привести в порядок зависимости модуля
	cd src && $(GO) mod tidy
