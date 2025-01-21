package handlers

import (
	"encoding/json"
	"net/http"

	"desafio-backend/models"
	"desafio-backend/utils"
)

// LoginHandler is the function that is going to be called when a request is made to the endpoint we set up on main.go
// this function
func SignUpHandler(w http.ResponseWriter, r *http.Request) {

	var newUserData models.User

	//this variable is to get the Json sent by the http request and decode it
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUserData)
	if err != nil {
		// if an error occurs when decodind it will return error 400  (Bad Request)
		http.Error(w, "Erro ao processar o JSON", http.StatusBadRequest)
		return
	}
	// calls the function that checks and saves the user in the archive
	err = utils.CheckAndSaveUser("data/users.json", newUserData)
	if err != nil {
		if err.Error() == "the email is already being used" || err.Error() == "username is already being used" {
			http.Error(w, `{"message": "`+err.Error()+`"}`, http.StatusConflict)
			return
		}
	}
}
