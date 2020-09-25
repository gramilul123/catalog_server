package controllers

import (
	"catalog_server/models"
	"catalog_server/utils"
	"encoding/json"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}

	if !bodyAccountDecoder(w, r, account, "Invalid request") {
		return
	}

	resp := account.Create()
	utils.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}

	if !bodyAccountDecoder(w, r, account, "Invalid request") {
		return
	}

	resp := models.Login(account.Email, account.Password)
	utils.Respond(w, resp)
}

func bodyAccountDecoder(w http.ResponseWriter, r *http.Request, account *models.Account, error string) bool {

	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		utils.Respond(w, utils.Message(false, error))

		return false
	}

	return true
}
