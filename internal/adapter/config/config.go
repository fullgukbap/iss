package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Container struct {
	DB   *DB
	HTTP *HTTP
}

type DB struct {
	Username    string
	Password    string
	Host        string
	Port        string
	Database    string
	IsMigration string
}

type HTTP struct {
	Port string
}

func New() (*Container, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	HTTP := &HTTP{
		Port: os.Getenv("HTTP_PORT"),
	}

	DB := &DB{
		Username:    os.Getenv("DB_USERNAME"),
		Password:    os.Getenv("DB_PASSWORD"),
		Host:        os.Getenv("DB_HOST"),
		Port:        os.Getenv("DB_PORT"),
		Database:    os.Getenv("DB_DATABASE"),
		IsMigration: os.Getenv("DB_IS_MIGRATION"),
	}

	return &Container{
		HTTP: HTTP,
		DB:   DB,
	}, nil
}
