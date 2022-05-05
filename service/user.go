package service

import (
	"encoding/json"
	"identification_email/config"
	"identification_email/constants"
	"identification_email/dbAccess"
	"identification_email/models"
	"identification_email/utility"
	"identification_email/utility/logger"
	"net/http"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	logger.I("SignUp invoked")
	defer logger.I("SignUp returned")

	creds := &models.Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)
	if err != nil {
		logger.E("GenerateFromPassword Failed for ", creds.Email, creds.Password, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := dbAccess.GetUserByEmail(creds.Email)
	if err != nil {
		logger.E("GetUserByEmail Failed for ", creds.Email, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if user != nil {
		logger.E("User already exist ", creds.Email, user.ID)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	_, err = mail.ParseAddress(creds.Email)
	if err != nil {
		logger.E("ParseAddress Failed for ", creds.Email, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	creds.Password = string(hashedPassword)
	err = dbAccess.CreatUser(creds)
	if err != nil {
		logger.E("CreatUser Failed for ", creds.Email, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utility.GenrateResponse(w, nil, err)

}

func Login(w http.ResponseWriter, r *http.Request) {
	logger.I("Login invoked")
	defer logger.I("Login returned")

	creds := &models.Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		logger.E("NewDecoder Failed for ", r.Body, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := dbAccess.GetUserByEmail(creds.Email)
	if err != nil {
		logger.E("GetUserByEmail Failed for ", creds.Email, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if user == nil {
		logger.E("User Not found ", creds.Email)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	if config.IsUserInSession(user.ID) {
		logger.E("User already logged in ", user.ID)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		logger.E("CompareHashAndPassword Failed for ", user.Password, creds.Password, creds.Email, err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err := utility.CreateToken(user.ID)
	if err != nil {
		logger.E("CreateToken Failed for ", user.ID, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	config.AddUserInSession(user.ID, token)

	data := make(map[string]interface{})
	data[constants.ConstToken] = token

	utility.GenrateResponse(w, data, err)

}

func Logout(w http.ResponseWriter, r *http.Request) {
	logger.I("Logout invoked")
	defer logger.I("Logout returned")

	userID, err := utility.GetUserFromToken(r)
	if err != nil {
		logger.E("GetUserFromToken Failed for ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	config.DeleteUserFromSession(userID)

	utility.GenrateResponse(w, nil, nil)

}

func Test(w http.ResponseWriter, r *http.Request) {
	logger.I("Test invoked")
	defer logger.I("Test returned")

	utility.GenrateResponse(w, nil, nil)

}
