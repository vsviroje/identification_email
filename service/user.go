package service

import (
	"encoding/json"
	"identification_email/dbAccess"
	"identification_email/models"
	"identification_email/utility"
	"net/http"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	creds := &models.Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := dbAccess.GetUserByEmail(creds.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if user != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	_, err = mail.ParseAddress(creds.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	creds.Password = string(hashedPassword)
	err = dbAccess.CreatUser(creds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utility.GenrateResponse(w, nil, err)

}

func Login(w http.ResponseWriter, r *http.Request) {

	creds := &models.Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := dbAccess.GetUserByEmail(creds.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if user == nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	utility.GenrateResponse(w, nil, err)

}

func LogOut(w http.ResponseWriter, r *http.Request) {

}
