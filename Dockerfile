# Stage 1: Builder stage
FROM golang:1.23 AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для загрузки зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем остальной код
COPY . .

# Собираем бинарный файл
RUN go build -o todo-list

# Stage 2: Final stage
FROM debian:bookworm-slim

# Устанавливаем необходимые зависимости (если потребуется)
#RUN apt-get update && apt-get install -y netcat && rm -rf /var/lib/apt/lists/*
RUN apt-get update && \
    apt-get install -y libc6 && \
    rm -rf /var/lib/apt/lists/*

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем бинарный файл из первого этапа
COPY --from=builder /app/todo-list .
COPY .env .

# Экспонируем порт 8080
EXPOSE 8080

# Запускаем приложение
CMD ["./todo-list"]





# Stage 1: Builder stage
# Stage 1: Builder stage
#FROM golang:1.23 AS builder
#
#WORKDIR /app
#
#COPY go.mod go.sum ./
#RUN go mod download
#COPY . .
#
#RUN go build -o todo-list
#
## Stage 2: Final stage
#FROM debian:bookworm-slim
#
#WORKDIR /app
#COPY --from=builder /app/todo-list .
#COPY .env .
#
## Make sure the binary has execute permissions
#RUN chmod +x ./todo-list
#
#EXPOSE 8080
#CMD ["./todo-list"]
