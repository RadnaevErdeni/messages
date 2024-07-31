package repository

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	messageTable = "messages"
)

type Conf struct {
	Host     string
	Port     string
	Username string
	BDname   string
	Password string
	SSLMode  string
}

func DBC(c Conf, dbname string) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.Username, c.Password, dbname, c.SSLMode)

	dbcon, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = dbcon.Ping()
	if err != nil {
		return nil, err
	}

	return dbcon, nil
}

func CreateDatabase(conf Conf) error {
	exists, err := databaseExists(conf, conf.BDname)
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	db, err := DBC(conf, "postgres")
	if err != nil {
		return fmt.Errorf("failed to connect to postgres database: %v", err)
	}
	defer db.Close()

	createDBSQL := fmt.Sprintf("CREATE DATABASE %s", conf.BDname)
	_, err = db.Exec(createDBSQL)
	if err != nil {
		return fmt.Errorf("failed to create database: %v", err)
	}

	return nil
}

func Migrations(conf Conf) error {

	dbURL := "postgres://" + conf.Username + ":" + conf.Password + "@" + conf.Host + ":" + conf.Port + "/" + conf.BDname + "?sslmode=" + conf.SSLMode
	cmd := exec.Command("migrate", "-path", "db/migrations", "-database", dbURL, "up")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error during migration: %v", err)
	}
	return nil
}

func WaitForDB(c Conf, timeout time.Duration) error {
	start := time.Now()
	for {
		db, err := DBC(c, "postgres")
		if err == nil {
			db.Close()
			return nil
		}

		if time.Since(start) > timeout {
			return fmt.Errorf("timed out waiting for database to be ready: %v", err)
		}

		time.Sleep(2 * time.Second)
	}
}
func databaseExists(c Conf, dbName string) (bool, error) {
	db, err := DBC(c, "postgres")
	if err != nil {
		return false, fmt.Errorf("failed to connect to postgres database: %v", err)
	}
	defer db.Close()

	var exists bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = '%s')", dbName)
	err = db.QueryRow(query).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check if database exists: %v", err)
	}

	return exists, nil
}
