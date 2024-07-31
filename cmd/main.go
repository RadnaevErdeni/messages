package main

import (
	"log"
	ms "messageService"
	"messageService/handler"
	"messageService/repository"
	"messageService/service"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Starting application")
	err := godotenv.Load("connect.env")
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

	conf := repository.Conf{
		Host:     dbHost,
		Port:     dbPort,
		Username: dbUsername,
		BDname:   dbName,
		Password: dbPassword,
		SSLMode:  dbSSLMode,
	}

	err = repository.WaitForDB(conf, 60*time.Second)
	if err != nil {
		log.Fatalf("Error waiting for database: %v", err)
	}

	err = repository.CreateDatabase(conf)
	if err != nil {
		log.Fatalf("Error creating database: %v", err)
	}

	db, err := repository.DBC(conf, conf.BDname)
	if err != nil {
		log.Fatalf("Error connecting to the new database: %v", err)
	}
	err = repository.Migrations(conf)
	if err != nil {
		log.Fatalf("Error initializing database 1: %v", err)
	}
	defer db.Close()

	repos := repository.NewRepository(db)
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaBrokers},
		Topic:    kafkaTopic,
		Balancer: &kafka.LeastBytes{},
	})
	defer kafkaWriter.Close()

	services := service.NewService(repos, kafkaWriter)
	handlers := handler.NewHandler(services)

	srv := new(ms.Server)
	if err := srv.Run(cnHost, cnPort, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Launch error HTTP server: %v", err)
	}
	logrus.Info("Application started successfully")
}
