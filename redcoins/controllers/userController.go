package controllers

import (
	"encoding/json"
	"net/http"

	"redcoins/models"
	u "redcoins/utils"
)

var CreateUser = func(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		u.Respond(w, u.Message(false, "Request Invalido"))
		return
	}

	resp := user.Create()
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Request invalido"))
		return
	}

	resp := models.Login(user.Email, user.Password)
	u.Respond(w, resp)
}

var RecoverPassword = func(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Request invalido"))
		return
	}

	resp := models.RecoverPassword(user.Email, user.Password)
	u.Respond(w, resp)
}
