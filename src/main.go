package main

import (
	"database/sql"
	"log"

	"github.com/jacobkania/mmbt/src/configuration"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.Println("--MMBT Starting--")

	config := configuration.LoadConfig()

	router := mux.NewRouter()

	db, err := sql.Open("sqlite3", "./mmbt.db")
	if err != nil {
		log.Fatalf("Database failed to open")
	}

	server := configuration.Server{
		Config: config,
		Router: router,
		DB:     db,
	}

	log.Fatal(server.Run())
}
