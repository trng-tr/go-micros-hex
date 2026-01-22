package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type ApiServer struct {
	Hostname string
	Port     string
}

type DBServer struct {
	DB_URL string
}

type RemoteApiUrl struct {
	CustomerBaseUrl string
	ProductBaseUrl  string
}

type AppConfig struct {
	ApiServer
	DBServer
	RemoteApiUrl
}

func (app *AppConfig) LoadConfig() {
	app.Hostname = getEnv("APP_HOSTNAME", "localhost")
	app.Port = getEnv("APP2_HOSTPORT", "8083")
	app.DB_URL = getEnv("ORDER_DB_URL", "postgres://go-app3-user:go-app3-pass@localhost:5435/goapp3db?sslmode=disable")
	app.CustomerBaseUrl = getEnv("CUSTOMER_BASE_URL", "http://localhost:8081/api/v1")
	app.ProductBaseUrl = getEnv("PRODUCT_BASE_URL", "http://localhost:8082/api/v1")
}

func getEnv(key, defalutValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defalutValue
}

func (appS *AppConfig) GetAppServer() string {
	return fmt.Sprintf("%s:%s", appS.Hostname, appS.Port)
}

func (appDb *AppConfig) GetDbServer() (*sql.DB, error) {
	sql, err := sql.Open("postgres", appDb.DB_URL)
	if err != nil {
		return nil, err
	}
	if err := sql.Ping(); err != nil {
		return nil, err
	}
	var t = 50 * time.Minute
	sql.SetConnMaxIdleTime(t)
	sql.SetMaxIdleConns(50)
	sql.SetConnMaxLifetime(t)

	return sql, nil
}
