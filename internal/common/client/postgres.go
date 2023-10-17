package client

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgres(addr string) (*sqlx.DB, error) {
	config, err := getConfig(addr)
	if err != nil {
		return nil, err
	}
	configStr :=
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			config.Host, config.Port, config.Username, config.Password,
			config.DBName, config.SSLMode)
	db, err := sqlx.Connect("postgres", configStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getConfig(addr string) (*PostgresConfig, error) {
	addrSplit := strings.Split(addr, ":")
	host := addrSplit[0]
	if host == "" {
		return nil, errors.New("Empty postgres host")
	}
	port := addrSplit[1]
	if port == "" {
		return nil, errors.New("Empty postgres port")
	}
	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		return nil, errors.New("Empty POSTGRES_USER")
	}
	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		return nil, errors.New("Empty POSTGRES_PASSWORD")
	}
	db := os.Getenv("POSTGRES_DB")
	if db == "" {
		return nil, errors.New("Empty POSTGRES_DB")
	}
	sslMode := os.Getenv("POSTGRES_SSLMODE")
	if sslMode == "" {
		return nil, errors.New("Empty POSTGRES_SSLMODE")
	}

	return &PostgresConfig{
		Host:     host,
		Port:     port,
		Username: user,
		Password: password,
		DBName:   db,
		SSLMode:  sslMode,
	}, nil
}
