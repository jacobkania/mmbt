package configuration

import (
	"github.com/jacobkania/mmbt/src/handlers"

	mux "github.com/gorilla/mux"
)

// SetRoutes sets all routes on the provided mux router
func setRoutes(r *mux.Router) {
	r.HandleFunc("/", handlers.IndexHandler)
}
