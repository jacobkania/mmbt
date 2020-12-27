package configuration

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

// InitializeDB runs migrations and opens database connection
func InitializeDB(config *Config) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), config.DbURL)
	if err != nil {
		log.Fatal("Connection to database failed")
	}

	return conn
}
