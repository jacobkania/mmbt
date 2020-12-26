package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
)

// IndexHandler `/`
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "This is a test with a random number: %v", rand.Int())
}
