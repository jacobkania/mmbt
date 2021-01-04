package handlers

import (
	v1 "mmbt/handlers/v1"

	"github.com/gorilla/mux"
)

// RouterV1 handles sub-routing for API v1
func RouterV1(apiV1 *mux.Router) {
	apiV1.StrictSlash(true)
	apiV1.HandleFunc("", v1.IndexHandler)
}
