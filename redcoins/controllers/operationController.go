package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"redcoins/models"
	u "redcoins/utils"

	"github.com/gorilla/mux"
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

var GetOperationByUser = func(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["user_id"]
	idUInt, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		u.Respond(w, u.Message(false, "Request invalido"))
		return
	}

	response, operations := models.GetOperationByUser(idUInt)

	if operations == nil {
		u.Respond(w, response)
		return
	}

	if len(operations) == 0 {
		u.Respond(w, u.Message(false, "Nenhuma transacao encontrada para esse usuario"))
		return
	}

	u.Respond(w, response)
}

var GetOperationByDate = func(w http.ResponseWriter, r *http.Request) {

	date := mux.Vars(r)["date"]

	response, operations := models.GetOperationByDate(date)

	if operations == nil {
		u.Respond(w, response)
		return
	}

	if len(operations) == 0 {
		u.Respond(w, u.Message(false, "Nenhuma transacao encontrada para essa data"))
		return
	}

	u.Respond(w, response)
}
