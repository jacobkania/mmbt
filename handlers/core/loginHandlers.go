package core

import (
	"encoding/json"
	"fmt"
	"net/http"

	"mmbt/configuration"
	"mmbt/db"
	"mmbt/models"

	"golang.org/x/crypto/bcrypt"
)

// LoginHandler creates an access token and returns it
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

type registerAccountParams struct {
	FullName     string
	PrimaryEmail string
	Passw        string
}

// RegisterAccountHandler creates a user
func RegisterAccountHandler(w http.ResponseWriter, r *http.Request) {
	newUserParams := &registerAccountParams{}

	json.NewDecoder(r.Body).Decode(newUserParams)

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(newUserParams.Passw), configuration.Config.BCryptCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Password encoding failure.")
		return
	}

	newUser := &models.User{}

	newUser.FullName = newUserParams.FullName
	newUser.PrimaryEmail = newUserParams.PrimaryEmail
	newUser.Passw = string(hashedPass)

	_, err = db.DB.Model(newUser).Insert()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}
