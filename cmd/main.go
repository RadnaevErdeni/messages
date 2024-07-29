package main

import (
	ms "messageService"
	"messageService/db/migrations"
	"messageService/handler"
	"messageService/kafka"
	"messageService/repository"
	"messageService/service"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Starting application")
	err := migrations.StartDBmain()
	if err != nil {
		logrus.Fatalf("Error initializing database 1: %v", err)
	}
	err = godotenv.Load("connect.env")
	if err != nil {
		logrus.Fatalf("Failed to load .env file: %v", err)
	}
	logrus.Info("Loaded connect.env file successfully")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSL_MODE")
	kafkaBrokers := os.Getenv("KAFKA_BROKERS")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	cnPort := os.Getenv("CON_PORT")
	cnHost := os.Getenv("CON_HOST")

	db, err := repository.DBC(repository.Conf{
		Host:     dbHost,
		Port:     dbPort,
		Username: dbUsername,
		Password: dbPassword,
		BDname:   dbName,
		SSLMode:  dbSSLMode,
	})
	if err != nil {
		logrus.Fatalf("Failed to initialize the database: %v", err)
	}

	repos := repository.NewRepository(db)
	producer := kafka.NewProducer([]string{kafkaBrokers}, kafkaTopic)
	consumer := kafka.NewConsumer([]string{kafkaBrokers}, "message-group", kafkaTopic, repos.Kafka)

	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(ms.Server)
	if err := srv.Run(cnHost, cnPort, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Launch error HTTP server: %v", err)
	}
	logrus.Info("Application started successfully")
}
