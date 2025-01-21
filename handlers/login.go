package handlers

import (
	"encoding/json"
	"net/http"

	"desafio-backend/models"
)

// LoginHandler is the function that is going to be called when a request is made to the endpoint we set up on main.go
// this function
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	//initialize a variable with the "type" loginData that comes from the models package
	var loginData models.LoginData
	//this variable is to get the Json sent by the http request and decode it
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginData)
	if err != nil {
		// if an error occurs when decodind it will return error 400  (Bad Request)
		http.Error(w, "Erro ao processar o JSON", http.StatusBadRequest)
		return
	}

	// this step verifies if the logindata is valid
	if loginData.Username == "felipe" && loginData.Password == "123" {
		// if it is valid it returns a sucess state, writting it on the header of the response
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Login bem-sucedido!"})
	} else {
		// if the data is invalid it will send an error of unauthorized also written in the header of the response
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"message": "Credenciais inválidas"})
	}
}
