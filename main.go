package main

import (
	"database/sql"
	"log"

	"mmbt/configuration"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.Println("--   MMBT Starting   --")
	log.Println("--  (c) Jacob Kania  --")
	log.Println("-----------------------")
	log.Println("MIT License: free to use and redistribute")
	log.Println("See LICENSE file on https://github.com/jacobkania/mmbt")
	log.Println("-----------------------")

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
