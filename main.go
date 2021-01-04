package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"mmbt/configuration"
	"mmbt/db"
	"mmbt/handlers"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("--   MMBT Starting   --")
	log.Println("--  (c) Jacob Kania  --")
	log.Println("-----------------------")
	log.Println("MIT License: free to use and redistribute")
	log.Println("See LICENSE file on https://github.com/jacobkania/mmbt")
	log.Println("-----------------------")

	configuration.LoadConfig()
	dbConn := db.InitializeDB(configuration.Config.DbURL)

	router := mux.NewRouter()
	handlers.SetRoutes(router)

	endSig := make(chan os.Signal, 1)
	signal.Notify(endSig, syscall.SIGTERM, syscall.SIGINT)

	httpServer := getHTTPServer(router)
	go httpServer.ListenAndServe()

	log.Printf("Server started on HTTP: %v\n", configuration.Config.HTTPPort)

	<-endSig

	log.Println("--   MMBT Stopping   --")
	httpServer.Shutdown(context.Background())
	dbConn.Close()
}

func getHTTPServer(r *mux.Router) *http.Server {
	serverAddress := ":" + strconv.Itoa(configuration.Config.HTTPPort)

	return &http.Server{
		Addr:         serverAddress,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
}
