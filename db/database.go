package db

import (
	"log"

	"github.com/go-pg/pg/v10"
)

// DB is the application database
var DB *pg.DB

// InitializeDB runs migrations and opens database connection
func InitializeDB(dbURL string) *pg.DB {
	opt, err := pg.ParseURL(dbURL)
	if err != nil {
		log.Fatal(err)
	}

	conn := pg.Connect(opt)

	DB = conn
	return conn
}
