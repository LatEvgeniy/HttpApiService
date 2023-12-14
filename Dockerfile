# Используем базовый образ, например, с вашим Golang приложением
FROM golang:latest

# Рабочая директория
WORKDIR /app

# Копируем go.mod и go.sum для эффективного кэширования зависимостей
COPY go.mod .
COPY go.sum .

# Загружаем зависимости
RUN go mod download

# Копируем исходники в контейнер
COPY . .

# Экспонируем порт, на котором будет слушать приложение
EXPOSE 8080

# Устанавливаем переменные окружения, например, адрес RabbitMQ
ENV RABBITMQ_HOST=rabbitmq
ENV RABBITMQ_PORT=5672

# Команда для запуска приложения
CMD ["go", "run", "main.go"]