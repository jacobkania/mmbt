package configuration

import (
	"log"
	"mmbt/constants"
	"mmbt/handlers"
	"net/http"
	"strconv"

	mux "github.com/gorilla/mux"
)

// SetRoutes sets all routes on the provided mux router
func setRoutes(server *Server) {
	var r *mux.Router = server.Router

	/* API Routes */
	r.HandleFunc("/api", handlers.IndexHandler)

	/* Frontend Routes */
	if server.Config.Environment == constants.Environment.Development {
		redirectURL := "http://" + server.Config.Host + ":" + strconv.Itoa(server.Config.DevFrontendPort)

		r.HandleFunc("/{rest:.*}", func(w http.ResponseWriter, r *http.Request) {
			log.Printf("redirect to: %v", redirectURL)
			http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
		})
	} else {
		r.Handle("/{rest:.*}", http.FileServer(http.Dir("./js/build")))
	}
}
