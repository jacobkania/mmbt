package core

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"mmbt/configuration"
	"mmbt/db"
	"mmbt/models"
	"mmbt/services/auth"

	"golang.org/x/crypto/bcrypt"
)

type loginParams struct {
	Username string
	Passw    string
}

// LoginHandler creates an access token and returns it
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	reqTimeEqualizer := time.After(3 * time.Second)
	loginParams := &loginParams{}

	json.NewDecoder(r.Body).Decode(loginParams)

	if len(loginParams.Passw) < 10 {
		<-reqTimeEqualizer
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.Error{Error: "Incorrect login information"})
		return
	}

	user := &models.User{}

	db.DB.Model(user).Where("username = ?", loginParams.Username).Select()

	if user.ID == 0 {
		<-reqTimeEqualizer
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.Error{Error: "Incorrect login information"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Passw), []byte(loginParams.Passw))
	if err != nil {
		<-reqTimeEqualizer
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.Error{Error: "Incorrect login information"})
		return
	}

	// v v v Below here, user is authenticated v v v

	tokenSvc := &auth.TokenService{DB: db.DB}

	token, err := tokenSvc.CreateToken(user)

	if err != nil {
		<-reqTimeEqualizer
		log.Printf("ERROR: handlers/core/loginHandlers.go:LoginHandler : %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Error{Error: "Incorrect login information"})
		return
	}

	<-reqTimeEqualizer
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(token)
}

type registerAccountParams struct {
	FullName string
	Username string
	Passw    string
}

// RegisterAccountHandler creates a user
func RegisterAccountHandler(w http.ResponseWriter, r *http.Request) {
	newUserParams := &registerAccountParams{}

	json.NewDecoder(r.Body).Decode(newUserParams)

	if len(newUserParams.Passw) < 10 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(models.Error{Error: "Password must be at least 10 characters long"})
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(newUserParams.Passw), configuration.Config.BCryptCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Error{Error: "Password encoding failure."})
		return
	}

	newUser := &models.User{}

	newUser.FullName = newUserParams.FullName
	newUser.Username = newUserParams.Username
	newUser.Passw = string(hashedPass)

	_, err = db.DB.Model(newUser).Insert()
	if err != nil {
		log.Printf("ERROR: handlers/core/loginHandlers.go:LoginHandler : %v", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(models.Error{Error: "Could not save user account information"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}
