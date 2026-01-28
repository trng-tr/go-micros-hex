package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type ServerConfig struct {
	Host string
	Port string
}

type DatabaseConfig struct {
	DatabaseURL string
}

/*
	AppConfig: veut dire la struct AppConfig extends

les deux austres structres, equivaut à mettre tout dans
la meme struct AppConfig
*/
type AppConfig struct {
	ServerConfig
	DatabaseConfig
}

func (app *AppConfig) LoadConfigurations() {
	app.Host = getEnv("APP_HOSTNAME", "localhost")
	app.Port = getEnv("APP2_HOSTPORT", "8082")
	app.DatabaseURL = getEnv("PRODUCT_DB_URL", "postgres://go-app2-user:go-app2-pass@localhost:5434/goapp2db?sslmode=disable")
}

func (appServer *AppConfig) GetServerAddress() string {
	return fmt.Sprintf("%s:%s", appServer.Host, appServer.Port)
}

func (dbServer *AppConfig) GetDbDNS() (*sql.DB, error) {
	sqlBD, err := sql.Open("postgres", dbServer.DatabaseURL)
	if err != nil {
		return nil, err
	}
	if err := sqlBD.Ping(); err != nil {
		return nil, err
	}
	sqlBD.SetMaxIdleConns(50)
	sqlBD.SetConnMaxIdleTime(50 * time.Minute)
	sqlBD.SetConnMaxLifetime(50 * time.Minute)
	return sqlBD, nil
}

func getEnv(key string, defaultV string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultV
	}
	return value
}
