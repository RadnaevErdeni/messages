# Используйте официальный образ Go
FROM golang:1.21 AS builder

# Установите рабочую директорию
WORKDIR /app

# Скопируйте go.mod и go.sum
COPY go.mod go.sum ./

# Загрузите зависимости
RUN go mod download

# Скопируйте исходный код
COPY . .

# Проверьте содержимое директории
RUN ls -la /app

# Установите миграционный инструмент
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Сборка приложения для Linux
RUN GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

# Установите права на выполнение
RUN chmod +x ./main

# Запустите приложение
CMD ["./main"]
