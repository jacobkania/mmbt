package configuration

import (
	"net/http"
	"strconv"
	"time"

	mux "github.com/gorilla/mux"
)

// Server holds configuration for the server
type Server struct {
	Config *Config
	Router *mux.Router
}

// HTTPServer gets the http server
func (s *Server) HTTPServer() *http.Server {
	setRoutes(s)
	serverAddress := ":" + strconv.Itoa(s.Config.HTTPPort)

	return &http.Server{
		Addr:         serverAddress,
		Handler:      s.Router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
}
