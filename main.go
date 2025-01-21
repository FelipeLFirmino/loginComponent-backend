package main

import (
	"desafio-backend/handlers"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Middleware that allows the CORS to happen, using the exact url of the local host of the frontend
// CORS means  (Cross-Origin Resource Sharing)  and it blocks the sharing of two different servers
func simpleCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200") // Permitir todas as origens
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		next.ServeHTTP(w, r)
	})
}

func main() {
	//initializing a router
	r := httprouter.New()

	//we are using the built in handlerFunc that comes from httprouter
	// this function receives 3 parameters, the first is the http method that we are "listening" for, the second one is the path endpoint, the third is the function
	// which is going to be executed, this function is a handler and gets the request and returns a response
	r.HandlerFunc("POST", "/login", handlers.LoginHandler)

	r.HandlerFunc("POST", "/signup", handlers.SignUpHandler)

	//this line of code listens  for erros in the server,
	//in the port 8080 using the router we initialized
	log.Fatal(http.ListenAndServe(":8080", simpleCORS(r)))
}
