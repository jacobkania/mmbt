package configuration

import (
	"crypto/tls"
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	mux "github.com/gorilla/mux"
)

// Server holds configuration for the server
type Server struct {
	Config      *Config
	Router      *mux.Router
	DB          *sql.DB
	httpsServer *http.Server
	httpServer  *http.Server
}

// Run initializes and sets the server
// Before calling this method on a Server, you must first set the Config, Router, and Db dependencies.
func (s *Server) Run() error {
	setRoutes(s.Router)

	httpURL := ":" + strconv.Itoa(s.Config.HTTPPort)
	httpsURL := ":" + strconv.Itoa(s.Config.HTTPSPort)
	if s.Config.HTTPOnly {
		s.newInsecureServer(httpURL)

		log.Printf("*** WARNING: HTTPS IS DISABLED ***")
		log.Printf("*** SERVER WILL BE INSECURE ***")
		log.Printf("*** DO NOT USE IN PRODUCTION ***")
		log.Printf("Server starting on HTTP: %v", s.Config.HTTPPort)
		return s.httpServer.ListenAndServe()
	}

	s.newServer(httpsURL, httpURL)

	log.Printf("Server starting on HTTPS: %v", s.Config.HTTPSPort)

	go s.httpServer.ListenAndServe()
	return s.httpsServer.ListenAndServeTLS(s.Config.CertFilePath, s.Config.KeyFilePath)
}

// Initial setup for the http and https servers.
func (s *Server) newServer(tlsServerAddress string, redirectServerAddress string) {
	s.httpsServer = &http.Server{
		Addr:         tlsServerAddress,
		Handler:      s.Router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig: &tls.Config{
			PreferServerCipherSuites: true,
			CurvePreferences:         []tls.CurveID{tls.CurveP256, tls.X25519},
			MinVersion:               tls.VersionTLS12,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			},
		},
	}

	s.httpServer = &http.Server{
		Addr:         redirectServerAddress,
		Handler:      redirectToHTTPS(tlsServerAddress),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
}

func (s *Server) newInsecureServer(serverAddress string) {
	s.httpServer = &http.Server{
		Addr:         serverAddress,
		Handler:      s.Router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
}

func redirectToHTTPS(url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, url+r.RequestURI, http.StatusMovedPermanently)
	}
}
