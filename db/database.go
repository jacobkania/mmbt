package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn

// InitializeDB runs migrations and opens database connection
func InitializeDB(dbURL string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatal("Connection to database failed")
	}

	db = conn
	return conn
}

// Query does a query on db
func Query(s string) (pgx.Rows, error) {
	return db.Query(context.Background(), s)
}
