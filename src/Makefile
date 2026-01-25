APP_NAME=tic-tac-toe
MAIN_PATH=./cmd/app/main.go

run:
	@echo "Запуск..."
	go run $(MAIN_PATH)

build:
	@echo "Сборка проекта..."
	go build -o $(APP_NAME) $(MAIN_PATH)

clean:
	@echo "Очистка..."
	rm -f $(APP_NAME)

help:
	@echo "Доступные команды:"
	@echo "  make run   - запустить приложение"
	@echo "  make build - собрать исполняемый файл"
	@echo "  make clean - удалить исполняемый файл"

.PHONY: run build clean help
