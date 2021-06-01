package helpers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// InitPostgres initiate the database connection based on .env
func InitPostgres() *sql.DB {
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	sslMode := os.Getenv("POSTGRES_SSL_MODE")
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName, sslMode)
	conn, err := sql.Open(`postgres`, connStr)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	conn.SetMaxOpenConns(50)
	conn.SetMaxIdleConns(50)

	return conn
}
