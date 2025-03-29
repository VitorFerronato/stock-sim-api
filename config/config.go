package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	DB              *sql.DB
	AlphavantageKey string
}

func NewConfig() *Config {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	alphavantageKey := os.Getenv("ALPHA_VANTAGE_API_KEY")
	if alphavantageKey == "" {
		panic("ALPHA_VANTAGE_API_KEY not set in environment")
	}

	return &Config{
		DB:              db,
		AlphavantageKey: alphavantageKey,
	}
}
