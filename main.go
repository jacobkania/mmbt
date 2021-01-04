package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"mmbt/configuration"
	"mmbt/db"

	"github.com/gorilla/mux"
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

	dbConn := db.InitializeDB(config.DbURL)

	endSig := make(chan os.Signal, 1)
	signal.Notify(endSig, syscall.SIGTERM, syscall.SIGINT)

	server := configuration.Server{
		Config: config,
		Router: router,
	}

	httpServer := server.HTTPServer()
	go httpServer.ListenAndServe()

	log.Printf("Server started on HTTP: %v\n", config.HTTPPort)

	<-endSig

	log.Println("--   MMBT Stopping   --")
	httpServer.Shutdown(context.Background())
	dbConn.Close(context.Background())
}
