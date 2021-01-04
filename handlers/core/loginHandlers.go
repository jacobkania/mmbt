package core

import (
	"encoding/json"
	"net/http"
)

// LoginHandler creates an access token and returns it
func LoginHandler(w http.ResponseWriter, r *http.Request) {

}

type registerAccountParams struct {
	FullName     string
	PrimaryEmail string
	Passw        string
}

// RegisterAccountHandler creates a user
func RegisterAccountHandler(w http.ResponseWriter, r *http.Request) {
	newUser := &registerAccountParams{}

	json.NewDecoder(r.Body).Decode(newUser)

	// hashedPass, err := bcrypt.GenerateFromPassword([]byte(newUser.Passw), configuration.Config.BCryptCost)
}
