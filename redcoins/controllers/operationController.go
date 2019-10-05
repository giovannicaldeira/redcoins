package controllers

import (
	"encoding/json"
	"net/http"

	"redcoins/models"
	u "redcoins/utils"
)

var CreateOperation = func(w http.ResponseWriter, r *http.Request) {

	operation := &models.Operation{}
	err := json.NewDecoder(r.Body).Decode(operation)

	if err != nil {
		u.Respond(w, u.Message(false, "Request Invalido"))
		return
	}

	resp := operation.Create()
	u.Respond(w, resp)
}
