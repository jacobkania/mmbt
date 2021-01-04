package db

import (
	"github.com/go-pg/pg/v10"
)

// DB is the application database
var DB *pg.DB

// InitializeDB runs migrations and opens database connection
func InitializeDB(dbURL string) *pg.DB {
	conn := pg.Connect(&pg.Options{Addr: dbURL})

	DB = conn
	return conn
}
