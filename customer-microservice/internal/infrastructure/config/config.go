package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

type AppServer struct {
	Hostname string
	Port     string
}

type DBServer struct {
	DbUrl string
}

type AppConfig struct {
	AppServer
	DBServer
}

func (app *AppConfig) LoadConfig() {
	app.Hostname = getEnv("APP_HOSTNAME", "localhost")
	app.Port = getEnv("APP1_HOSTPORT", "8081")
	app.DbUrl = getEnv("CUSTOMER_DB_URL", "postgres://go-app1-user:go-app1-pass@localhost:5433/goapp1db?sslmode=disable")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}

func (appServer *AppConfig) GetAppServer() string {
	return fmt.Sprintf("%s:%s", appServer.Hostname, appServer.Port)
}

func (dbServer *AppConfig) GetDBDns() (*sql.DB, error) {
	sqlDb, err := sql.Open("postgres", dbServer.DbUrl)
	if err != nil {
		return nil, err
	}

	if err := sqlDb.Ping(); err != nil {
		return nil, err
	}

	sqlDb.SetMaxIdleConns(50)
	sqlDb.SetConnMaxIdleTime(50 * time.Minute)
	sqlDb.SetConnMaxLifetime(50 * time.Minute)

	return sqlDb, nil
}
