package v1

import (
	"fmt"
	"math/rand"
	"net/http"
)

// IndexHandler `/api/v1/`
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "This is a test with a random number: %v", rand.Int())
}
